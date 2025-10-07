# bri-edc

## Overview
`bri-edc` is a project designed to manage user transactions efficiently. When a user initiates a transaction, the balance is temporarily stored as transactions are processed in batch during settlement. Batch settlement aims to increase efficiency and reduce costs by grouping multiple approved transactions together for simultaneous processing, optimizing both time and system resources.

## Setup
- Install Docker
- Ensure Docker Compose is available

## Usage
```bash
git clone https://github.com/farisfadhail/bri-edc.git
cd bri-edc
docker-compose up --build

# Run migrations and seeders
docker compose exec api ./migration migrate:up
docker compose exec api ./seeder seed:run
```