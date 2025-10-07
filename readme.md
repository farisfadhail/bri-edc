migrate -database "mysql://faris:password@tcp(localhost:3306)/bri-edc?charset=utf8mb4&parseTime=True&loc=Local" -path database/migrations up

cd internal/injector && wire

docker-compose build --no-cache
docker-compose up -d
docker-compose down -v
