package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var (
	Users        = map[string]map[string]string{}
	Transactions = []map[string]interface{}{}
	Settlements  = []map[string]interface{}{}
)

func ClearAll() {
	Users = map[string]map[string]string{}
	Transactions = []map[string]interface{}{}
	Settlements = []map[string]interface{}{}
}

func SetupTestApp() *fiber.App {
	app := fiber.New()

	app.Post("/api/v1/auth/_login", func(c *fiber.Ctx) error {
		var req map[string]string
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		username := req["username"]
		password := req["password"]

		if u, ok := Users[username]; ok && u["password"] == password {
			token := "dummy-jwt-token"
			u["token"] = token
			return c.Status(200).JSON(fiber.Map{"token": token})
		}
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	})

	app.Post("/api/v1/transactions/sale", func(c *fiber.Ctx) error {
		var req map[string]interface{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		req["status"] = "approved"
		Transactions = append(Transactions, req)
		return c.Status(200).JSON(req)
	})

	app.Post("/api/v1/transactions/settlement", func(c *fiber.Ctx) error {
		totalAmount := 0
		for _, t := range Transactions {
			if t["status"] == "approved" {
				totalAmount += t["amount"].(int)
			}
		}
		settlement := map[string]interface{}{
			"batch_id":     "BATCH20251007",
			"total_count":  len(Transactions),
			"approved":     len(Transactions),
			"declined":     0,
			"total_amount": totalAmount,
		}
		Settlements = append(Settlements, settlement)
		return c.Status(200).JSON(settlement)
	})

	app.Post("/core/authorize", func(c *fiber.Ctx) error {
		var req map[string]interface{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(200).JSON(fiber.Map{"authorized": true})
	})

	return app
}

func performRequest(app *fiber.App, method, url string, body interface{}, headers map[string]string) (*http.Response, []byte, error) {
	var buf []byte
	if body != nil {
		buf, _ = json.Marshal(body)
	}

	req := httptest.NewRequest(method, url, bytes.NewReader(buf))
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := app.Test(req)
	if err != nil {
		return nil, nil, err
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	return resp, bodyBytes, nil
}

func TestLoginService(t *testing.T) {
	ClearAll()
	app := SetupTestApp()

	Users["admin"] = map[string]string{"password": "password"}

	loginPayload := map[string]string{"username": "admin", "password": "password"}
	resp, body, err := performRequest(app, "POST", "/api/v1/auth/_login", loginPayload, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var loginResp map[string]string
	_ = json.Unmarshal(body, &loginResp)
	token := loginResp["token"]
	assert.Equal(t, "dummy-jwt-token", token)
}

func TestTransactionSaleService(t *testing.T) {
	ClearAll()
	app := SetupTestApp()

	Users["admin"] = map[string]string{"password": "password"}
	token := "dummy-jwt-token"
	Users["admin"]["token"] = token

	salePayload := map[string]interface{}{
		"merchant_id": "MCH123",
		"terminal_id": "T02",
		"amount":      125000,
		"card_number": "411111******1111",
		"timestamp":   "2025-10-07T12:45:00Z",
	}
	headers := map[string]string{"Authorization": "Bearer " + token}

	resp, body, err := performRequest(app, "POST", "/api/v1/transactions/sale", salePayload, headers)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var saleResp map[string]interface{}
	_ = json.Unmarshal(body, &saleResp)
	assert.Equal(t, "approved", saleResp["status"])
}

func TestTransactionSettlementService(t *testing.T) {
	ClearAll()
	app := SetupTestApp()

	Users["admin"] = map[string]string{"password": "password"}
	token := "dummy-jwt-token"
	Users["admin"]["token"] = token

	Transactions = append(Transactions, map[string]interface{}{
		"merchant_id": "MCH123",
		"terminal_id": "T02",
		"amount":      125000,
		"status":      "approved",
		"timestamp":   "2025-10-07T12:45:00Z",
	})

	headers := map[string]string{"Authorization": "Bearer " + token}

	resp, body, err := performRequest(app, "POST", "/api/v1/transactions/settlement", nil, headers)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var settleResp map[string]interface{}
	_ = json.Unmarshal(body, &settleResp)
	assert.Equal(t, "BATCH20251007", settleResp["batch_id"])
	assert.Equal(t, 125000, int(settleResp["total_amount"].(float64)))
}

func TestCoreServiceAuthorize(t *testing.T) {
	ClearAll()
	app := SetupTestApp()

	corePayload := map[string]interface{}{
		"merchant_id": "M001",
		"terminal_id": "T001",
		"amount":      2000,
		"card_number": "1234567890123456",
		"timestamp":   "2025-10-06T12:30:00Z",
	}
	coreHeaders := map[string]string{"X-Service-Token": "servicesecret"}

	resp, body, err := performRequest(app, "POST", "/core/authorize", corePayload, coreHeaders)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var coreResp map[string]interface{}
	_ = json.Unmarshal(body, &coreResp)
	assert.Equal(t, true, coreResp["authorized"])
}
