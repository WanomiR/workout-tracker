package main

import (
	"backend/internal/models"
	"errors"
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
	w.Header().Add("Access-Control-Allow-Credentials", "true")

	var payload models.UserCredentials
	err := readJSONPayload(w, r, &payload)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	// validate user against database
	user, err := app.DB.GetUserByEmail(payload.Email)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	// check password
	valid, err := passwordMatches(user.Password, payload.Password)
	if err != nil || !valid {
		writeJSONError(w, errors.New("invalid credentials password"), http.StatusBadRequest)
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

	resp := JSONResponse{
		Error:   false,
		Message: "user successfully authenticated",
		Data:    tokens,
	}

	writeJSONResponse(w, http.StatusAccepted, resp)
}
