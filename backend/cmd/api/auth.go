package main

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
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

type JwtUser struct {
	ID    int
	Email string
	Name  string
}

type TokensPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	jwt.RegisteredClaims
}

func (j *Auth) GenerateTokensPair(user *JwtUser) (TokensPair, error) {
	// create token
	accessToken := jwt.New(jwt.SigningMethodHS256)

	// set the claims
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["sub"] = user.Email
	claims["aud"] = j.Audience
	claims["iss"] = j.Issuer
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = "JWT"
	// set the expiry
	claims["exp"] = time.Now().UTC().Add(j.TokenExpiry).Unix()

	// create a signed accessToken
	signedAccessToken, err := accessToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokensPair{}, err
	}

	// create refresh token and set claims
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = user.Email
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()
	// set the expiry
	refreshTokenClaims["exp"] = time.Now().UTC().Add(j.RefreshExpiry).Unix()

	// create a signed refresh token
	signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokensPair{}, err
	}

	tokensPair := TokensPair{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	return tokensPair, nil
}

func (j *Auth) GetRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     j.CookieName,
		Path:     j.CookiePath,
		Value:    refreshToken,
		Expires:  time.Now().UTC().Add(j.TokenExpiry),
		MaxAge:   int(j.TokenExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   j.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}

func (j *Auth) GetExpiredRefreshCookie() *http.Cookie {
	return &http.Cookie{
		Name:     j.CookieName,
		Path:     j.CookiePath,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		Domain:   j.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}
