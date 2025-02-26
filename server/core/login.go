package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func Login(username string, password string, db DB_Handler) (User, error) {
	u := User{}
	var passwdFromDb string

	query, args := u.GetHashFrom(username)
	row := db.QueryRow(query, args...)
	err := row.Scan(&passwdFromDb)
	if err != nil {
		return User{}, fmt.Errorf("database error: %v", err)
	}

	if CheckPasswordHash(password, passwdFromDb) {
		query, args := u.Verify(username)
		row := db.QueryRow(query, args...)
		err = row.Scan(&u.ID, &u.Name, &u.Role.ID, &u.Role.Name)
		if err != nil {
			Logger.Error(err.Error())
			return User{}, fmt.Errorf("database error: %v", err)
		}
		return u, nil
	}
	return User{}, fmt.Errorf("database error: %v", err)
}

func LoginHandler(w http.ResponseWriter, r *http.Request, db DB_Handler) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := Login(creds.Username, creds.Password, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		Role:     user.Role.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: false,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	Logger.Info(fmt.Sprintf("User %s logged in", creds.Username))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		tokenString := cookie.Value
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func RequireRole(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value("claims").(*Claims)
		if !ok || claims.Role != role {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func AdminTestHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(*Claims)
	w.Write([]byte(fmt.Sprintf("Welcome admin %s!", claims.Username)))
}

func UserTestHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(*Claims)
	w.Write([]byte(fmt.Sprintf("Welcome user %s!", claims.Username)))
}

func EnableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8090")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		next.ServeHTTP(w, r)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0), // Expire immediately
		HttpOnly: false,
	})
	http.Redirect(w, r, "/", http.StatusFound)
}
