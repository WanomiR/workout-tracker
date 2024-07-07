package main

import (
	"backend/internal/models"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		"active", "Workout Tracker is up and running", "0.0.0",
	}

	writeJSONResponse(w, http.StatusOK, payload)
}

// Authenticate godoc
// @Summary authenticate user
// @Description Authenticate user by email address and password.
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.UserCredentials true "User credentials"
// @Success 202 {object} JSONResponse
// @Failure 400,500 {object} JSONResponse
// @Router /authenticate [post]
func (app *App) Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var payload models.UserCredentials
	err := readJSONPayload(w, r, &payload)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	// validate user against database
	user, err := app.DB.GetUserByEmail(payload.Email)
	if err != nil {
		writeJSONError(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// check password
	valid, err := passwordMatches(user.Password, payload.Password)
	if err != nil || !valid {
		writeJSONError(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// create a jwt user
	jwtUser := &JwtUser{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	// generate tokens
	tokens, err := app.Auth.GenerateTokensPair(jwtUser)
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}

	refreshCookie := app.Auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	resp := JSONResponse{
		Error:   false,
		Message: "user authenticated",
		// TODO: remove tokens from response
		Data: tokens,
	}

	writeJSONResponse(w, http.StatusAccepted, resp)
}

// RefreshToken godoc
// @Summary refresh jwt token
// @Description Refresh JWT token.
// @Tags auth
// @Produce json
// @Success 200 {object} JSONResponse
// @Failure 401,500 {object} JSONResponse
// @Router /refresh [get]
func (app *App) RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	for _, cookie := range r.Cookies() {
		if cookie.Name == app.Auth.CookieName {
			claims := &Claims{}
			refreshToken := cookie.Value

			// parse the token to get the claims
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (any, error) {
				return []byte(app.Auth.Secret), nil
			})
			if err != nil {
				writeJSONError(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}

			// find user by email
			user, err := app.DB.GetUserByEmail(claims.Subject)
			if err != nil {
				writeJSONError(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			jwtUser := JwtUser{
				ID:    user.ID,
				Email: user.Email,
				Name:  user.Name,
			}

			tokens, err := app.Auth.GenerateTokensPair(&jwtUser)
			if err != nil {
				writeJSONError(w, errors.New("error generating tokens"), http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, app.Auth.GetRefreshCookie(tokens.RefreshToken))

			resp := JSONResponse{
				Error:   false,
				Message: "token refreshed",
				// TODO: remove tokens from response
				Data: tokens,
			}

			writeJSONResponse(w, http.StatusOK, resp)
		}
	}
}

// Logout godoc
// @Summary logout
// @Security ApiKeyAuth
// @Description Logout and remove refresh token from cookie storage.
// @Tags auth
// @Produce json
// @Success 202 {object} JSONResponse
// @Router /logout [get]
func (app *App) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	http.SetCookie(w, app.Auth.GetExpiredRefreshCookie())

	resp := JSONResponse{
		Error:   false,
		Message: "user logged out",
	}

	writeJSONResponse(w, http.StatusAccepted, resp)
}
