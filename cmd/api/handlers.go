package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kiniconnet/react-go-tutorial/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}

	err := app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (app *application) Authenticate(w http.ResponseWriter, r *http.Request) {
	// Step 1: Read JSON payload
	var requestPayload struct {
		Email    string `json:"email" bson:"email"`
		Password string `json:"password" bson:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		fmt.Printf("Error reading JSON payload: %v\n", err)
		app.errorJSON(w, "Unable to read payload", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received payload: %+v\n", requestPayload)

	// Step 2: Validate input data
	if requestPayload.Email == "" || requestPayload.Password == "" {
		app.errorJSON(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Step 3: Validate the user against the database
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		fmt.Printf("Error fetching user from database: %v\n", err)
		app.errorJSON(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if user == nil {
		fmt.Println("No user found with the provided email")
		app.errorJSON(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	fmt.Printf("Fetched user from database: %+v\n", user)

	// Step 4: Check the password
	valid, err := user.PasswordMatch(requestPayload.Password)
	if err != nil {
		fmt.Printf("Error comparing passwords: %v\n", err)
		app.errorJSON(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if !valid {
		fmt.Println("Password does not match")
		app.errorJSON(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	fmt.Println("Password matched successfully")

	// Step 5: Create a JWT user
	u := jwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	fmt.Printf("Created JWT user: %+v\n", u)

	// Step 6: Generate token pair
	tokenPair, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		fmt.Printf("Error generating token pair: %v\n", err)
		app.errorJSON(w, "Unable to generate token pair", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Generated token pair: %+v\n", tokenPair)

	// Step 7: Set the refresh token as a secure cookie
	refreshCookie := app.auth.GetRefreshCookie(tokenPair.RefreshToken)
	http.SetCookie(w, refreshCookie)

	fmt.Println("Set refresh token cookie")

	// Step 8: Respond with success and include the access token in the response body
	response := JSONResponse{
		Error:   "false",
		Message: "Authentication successful",
		Data: map[string]string{
			"access_token":  tokenPair.Token,
			"refresh_token": tokenPair.RefreshToken,
		},
	}

	err = app.writeJSON(w, http.StatusOK, response)
	if err != nil {
		fmt.Printf("Error writing JSON response: %v\n", err)
		app.errorJSON(w, "Unable to write JSON response", http.StatusInternalServerError)
	}
}


func (app *application) Signup(w http.ResponseWriter, r *http.Request) {
	// Step 1: Read JSON payload
	var requestPayload struct {
		FirstName string `json:"first_name" bson:"first_name"`
		LastName  string `json:"last_name" bson:"last_name"`
		Email     string `json:"email" bson:"email"`
		Password  string `json:"password" bson:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		fmt.Printf("Error reading JSON payload: %v\n", err)
		app.errorJSON(w, "Unable to read payload", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received payload: %+v\n", requestPayload)

	// Step 2: Validate input data
	if requestPayload.FirstName == "" || requestPayload.LastName == "" || requestPayload.Email == "" || requestPayload.Password == "" {
		app.errorJSON(w, "All fields are required", http.StatusBadRequest)
		return
	}

	if !isValidEmail(requestPayload.Email) {
		app.errorJSON(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	if len(requestPayload.Password) < 8 {
		app.errorJSON(w, "Password must be at least 8 characters long", http.StatusBadRequest)
		return
	}

	// Step 3: Check if the email is already registered
	existingUser, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		fmt.Printf("Error checking email: %v\n", err)
		app.errorJSON(w, "Unable to check email", http.StatusInternalServerError)
		return
	}

	if existingUser != nil {
		app.errorJSON(w, "Email is already registered", http.StatusConflict)
		return
	}

	// Step 4: Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestPayload.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error hashing password: %v\n", err)
		app.errorJSON(w, "Unable to hash password", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Hashed password: %s\n", string(hashedPassword))

	// Step 5: Create a new user
	newUser := models.User{
		ID:        primitive.NewObjectID(),
		FirstName: requestPayload.FirstName,
		LastName:  requestPayload.LastName,
		Email:     requestPayload.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Insert the new user into the database
	err = app.DB.InsertUser(newUser)
	if err != nil {
		fmt.Printf("Error inserting user: %v\n", err)
		app.errorJSON(w, "Unable to create user", http.StatusInternalServerError)
		return
	}

	// Step 6: Respond with success
	response := JSONResponse{
		Error:   "false",
		Message: "User successfully created",
	}

	err = app.writeJSON(w, http.StatusCreated, response)
	if err != nil {
		fmt.Printf("Error writing JSON response: %v\n", err)
		app.errorJSON(w, "Unable to write JSON response", http.StatusInternalServerError)
	}
}
