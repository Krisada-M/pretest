package models

import jwt "github.com/dgrijalva/jwt-go"

// SignedDetail is model claims
type SignedDetail struct {
	Data Payload
	jwt.StandardClaims
}

// Payload Data
type Payload struct {
	Username string
}
