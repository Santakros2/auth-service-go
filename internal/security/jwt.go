package security

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userID, email, role string) (string, error) {
	//1. Define claims (payload)
	claims := jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"role":  role,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	}

	// 2. Creating token with signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3. Sign the token with secret key
	signedToken, err := token.SignedString([]byte("MY_SECRET_KEY"))

	if err != nil {
		return "", err
	}

	// 4. Return JWT string
	return signedToken, nil
}

func GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

func GenerateToken(userID, email, role string) (*TokenPair, error) {
	accessToken, err := GenerateAccessToken(userID, email, role)

	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// func GenerateAccessTokenWithRefreshToken(refresh string) (string, error) {

// 	// Hash the refresh token
// 	hashed := sha256.Sum256([]byte(refresh))

// 	// 2. Look up in the DB for refresh token
// 	rt, err := repo.FindRefreshTokenByHash(hex.EncodeToString(hashed[:]))
// 	if err != nil {
// 		return "", errors.New("invalid refresh token")
// 	}

// 	if rt.Revoked || time.Now().After(rt.ExpireAt) {
// 		return "", errors.New("refresh token expired or revoked")
// 	}

// 	// 4. fetch the user
// 	user, err := repo.FindByID(rt.UserID)
// 	if err != nil {
// 		return "", errors.New("user not found")
// 	}

// 	// 5. Generate new access token
// 	return GenerateAccessToken(user.ID, user.Email, user.Role), nil

// }
