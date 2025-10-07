# ADR: `bri-edc`

## Status
Accepted

## Context
The `bri-edc` project requires a Docker-based multi-service setup to ensure isolation, scalability, and easy management. Docker Compose is used to define and run multiple containers for the project.

## Decision
1. **Use MySQL 8.0**
    - **Rationale**: MySQL 8.0 provides modern features, better performance, and wide compatibility, making it suitable for the project's database needs.

2. **Healthcheck for `db` Service**
    - **Rationale**: Adding a healthcheck ensures that the `db` service is only considered ready once MySQL can accept connections, preventing dependent services from failing.

3. **Use `depends_on` with `condition: service_healthy`**
    - **Rationale**: Ensures dependent services start only after their dependencies are healthy, improving startup reliability.

4. **Environment Variables for Security**
    - **Rationale**: Variables such as `SERVICE_TOKEN` and `JWT_SECRET` allow secure and flexible configuration without exposing sensitive information in the source code.

5. **Port Mapping for Service Access**
    - **Rationale**: Mapping ports like `3307:3306` for MySQL and `8080:8080` for the API enables easy access during development and testing while keeping internal ports secure.

## Consequences
- **Benefits**:
    - Clear service isolation simplifies maintenance and management.
    - Better scalability by separating services by function.
    - Enhanced security through environment variables and proper access control.

- **Risks**:
    - Increased configuration complexity requires understanding Docker and Docker Compose.
    - Potential compatibility issues between MySQL versions and the application if not managed carefully.

## References
- [Docker Compose Health Checks](https://medium.com/@cbaah123/docker-compose-health-checks-made-easy-a-practical-guide-3a340571b88e)
- [Control startup and shutdown order with Compose](https://docs.docker.com/compose/how-tos/startup-order/)
