# Assessment Bank INA

Simple REST API For assesment test

## ![](https://cdn-icons-png.flaticon.com/24/2694/2694997.png) **Stack**

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

## ![](https://cdn-icons-png.flaticon.com/24/4319/4319207.png) **Framework**

[![](https://badgen.net/badge/github/Gin/blue?icon=github)](https://github.com/gin-gonic/gin)
[![](https://badgen.net/badge/github/godotenv/cyan?icon=github)](https://github.com/joho/godotenv)
[![](https://badgen.net/badge/github/JWT/black?icon=github)](https://github.com/golang-jwt/jwt)

## ![](https://cdn-icons-png.flaticon.com/24/610/610363.png) Preparation On Local

if you want to run code on your local machine, follow instruction that explain below :
notes : not need to create or import database if using docker because it's include running migration

1. With Docker
   - Prerequisite : Docker compose was installed in your local machine
   - enter to path repository and run `docker-compose -f docker-compose.yaml up -d --build`
2. Without Docker
   - Prerequisite : Go was installed in your local machine
   - enter to path folder and run **`go mod tidy`**
   - create database with name `assessment_bank_ina`
   - import up sql file on migrations folder or u can using golang migrate and running command `migrate -database "mysql://{username}:{password}@tcp(localhost:3306)/assessment_bank_ina" -path migrations up`
   - create .env file and put code below\
     `PORT=9010`
     `DB_USER=root`
     `DB_PASS=root`
     `DB_HOST=localhost`
     `DB_PORT=3306`
     `DB_NAME=assessment_bank_ina`
     `EXPIRED_TIME_TOKEN=10`in minutes
     `SECRET_KEY=17c11ae94e6859e0c04daae2f55b0073d2c947294ea38b79280ed0dd514c8454`
     `PASS_SECRET="abc&1*~#^2^#s0^=)^^7%b34"`
   - Running app with command `go run main.go local`
3. Import postman collection and environment postman to your local postman, the file was collected with this code project in zip file, every endpoint is isolated with authorization token using jwt , so for the first time if you want to try some endpoint , please execute endpoint post login, before execute this you need to create some user using endpoint post user, and then login using email and password to get token.
