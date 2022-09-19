# golang_APIServer
## Overview
- Simple api server with golang
- Using fiber webframework
- Referring layered architecture
## How to run
```
# go to deploy directory
cd deploy

# build api-server
docker-compose build

# run api-server with mysql
docker-compose up
```
## API Specification
### Join
- Method: POST
- Uri: /join
- Request Body
  - {"nickname": "\[NICKNAME\]"} 
- Resonse Body
  - {"ResultCode": 0, "Response": {"Id": \[LAST_INSERT_ID\]}}
### Login
- Method: POST
- Uri: /login
- Request Body
  - {"nickname": "\[NICKNAME\]"} 
- Resonse Body
  - {"ResultCode": 0, "Response": {"Id": \[ACCOUNT_ID\]}}
## Reference
- https://nurcahyaari.medium.com/how-i-implement-clean-code-architecture-on-golang-projects-68be58830621
- https://github.com/gofiber/fiber
- https://github.com/labstack/echo
