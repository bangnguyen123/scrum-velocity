# Well come to the Backend of Velocity Project

## 1. Set up database and env
1. *If use choose to use Docker, you don't need to worry about database installation, let's go ahead with our docker-compose.yml file*
*But, with setup on localhost, we need to do some setup for database*:
    - *install postgres*
    - *start postgress*
2. Create new `.env` file (touch .env)
3. Run: `cp .env.example .env`
4. Change the DATABASE_VARIABLE in .env following you information*

## 2. Use AIR directly on host to start the application
1. Install: Golang on your localhost
2. Initialize and dowload the dependencies: `make install-local`
3. Build: `make build-local`
4. Run app: `make run-local`

## 3. Use docker to start the application
1. Install Docker and docker-compose
2. Build Docker image: `make build`
3. Run: `make run`

## 4. Database Migration
We are using Prisma, this needs to be generated after we have any change in our schema.prisma
### 4.1. For Docker
1. `make prisma-generate` -> generate new update without update database
2. `make prisma-migrate` -> generate new update and apply to database
### 4.2. On local host:
1. `make prisma-generate-local` -> generate new update without update database
2. `make prisma-migrate-local` -> generate new update and apply to database

## 5. Install linter
1. Install: https://golangci-lint.run/usage/install/
2. Run lint: `golangci-lint run`