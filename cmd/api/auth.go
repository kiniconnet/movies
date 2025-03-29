package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth struct {
	Issuer        string
	Audience      string
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

type jwtUser struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
}

// Issueing token in pair
type TokenPairs struct {
	Token        string  
	RefreshToken string  
}

type Claims struct {
	jwt.StandardClaims
}


// GenerateTokenPair generates a new token pair
func (a *Auth) GenerateTokenPair(user *jwtUser) (TokenPairs, error) {
	// Create a token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	claims["sub"] = fmt.Sprintf("%d", user.ID)
	claims["iss"] = a.Issuer
	claims["aud"] = a.Audience
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = "JWT"

	// Set the expiry time jwt 
	claims["exp"] = time.Now().Add(a.TokenExpiry).UTC().Unix()

	// create a signed token
	signedAccessToken, err := token.SignedString([]byte(a.Secret))
	if err != nil {
		return TokenPairs{}, err
	} 

	// Create a refresh token and set claims
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = fmt.Sprintf("%d", user.ID)
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()

	// Set the expiry time for the refresh token
	refreshTokenClaims["exp"] = time.Now().Add(a.RefreshExpiry).UTC().Unix()

	// Create a signed refresh token
	signedRefreshToken, err := refreshToken.SignedString([]byte(a.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	// Create a token pair and populate the token and refresh token
	tokenPair := TokenPairs{
		Token: signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	// return the token pair
	return tokenPair, nil
}

func (a *Auth) GetRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     a.CookieName,
		Path:    a.CookiePath,
		Value:   refreshToken,
		Expires: time.Now().Add(a.RefreshExpiry),
		MaxAge: int(a.RefreshExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Secure:   true,
	}
}


func (a *Auth) GetExpiredRefreshCookie() *http.Cookie {
	return &http.Cookie{
		Name:     a.CookieName,
		Path:    a.CookiePath,
		Value:   "",
		Expires: time.Unix(0, 0),
		MaxAge: -1,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Secure:   true,
	}
}