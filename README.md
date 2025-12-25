#auth-service-go

Auth Service is the authentication and authorization microservice for a microservices-based backend system, implemented in Go.
It is responsible for user authentication, JWT generation, refresh token handling, and token validation.

This service is designed to run behind an API Gateway, while other services rely on JWTs for authorization.


---

Overview

In a microservices architecture, authentication should be centralized and isolated from business services.
The Auth Service acts as the single source of truth for:

User credentials

Authentication state

Token lifecycle management


Other services do not store passwords or authentication logic — they only verify tokens.


---

Key Features

User Signup and Login

Secure password hashing (bcrypt)

JWT access token generation

Refresh token support

Token validation endpoint

Logout and token revocation

Database-backed authentication

Clean and scalable Go project structure



---

Architecture

Client
  │
  ▼
API Gateway
  │
  ▼
Auth Service ───► Database
  │
  └──► Issues JWT → Other Services


---

Project Structure

auth-service-go/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── config/
│   ├── handler/
│   │   └── auth_handler.go
│   ├── service/
│   │   └── auth_service.go
│   ├── repository/
│   │   ├── user_repository.go
│   │   └── token_repository.go
│   ├── model/
│   │   ├── user.go
│   │   └── refresh_token.go
│   ├── middleware/
│   │   └── jwt_middleware.go
│   └── utils/
│       ├── jwt.go
│       └── password.go
│
├── migrations/
│   └── 001_create_auth_tables.sql
│
├── go.mod
├── go.sum
└── README.md


---

Authentication Flow

Signup

1. Client sends email and password


2. Password is hashed using bcrypt


3. User is stored in the database



Login

1. Credentials are verified


2. Auth Service generates:

Short-lived Access Token (JWT)

Long-lived Refresh Token



3. Refresh token is stored in the database



Token Validation

Other services validate JWTs locally

Auth Service is not called for every request



---

API Endpoints

Method	Endpoint	Description

POST	/signup	Register a new user
POST	/login	Authenticate user
POST	/refresh	Generate new access token
POST	/logout	Revoke refresh token
GET	/validate	Validate JWT token



---

Security Considerations

Passwords are never stored in plain text

JWTs are signed and time-bound

Refresh tokens are persisted and revocable

Stateless authentication for downstream services



---

Tech Stack

Language: Go

HTTP: net/http / chi / gin

Authentication: JWT

Database: MySQL / PostgreSQL

Hashing: bcrypt



---

Intended Usage

Works behind an API Gateway

Used by internal services for token validation

Scales independently from business services



---

Future Improvements

OAuth2 / Social Login

Role-based access control (RBAC)

gRPC support

Rate limiting

Public / private key JWT signing



---

Author

Shrutik Borikar
Backend • Go • Microservices


---

If you want, I can also:

Generate folder-wise Go code

Add JWT middleware

Create Docker + docker-compose

Align it with your Gateway service


Just tell me what’s next 👌
