package util

import (
	"CPS406-Assignment-Backend/pkg/jwtM"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"time"
)

func SetTokenAsCookie(w http.ResponseWriter, token string) {
	// Create a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",                   // Cookie name
		Value:    token,                          // JWT token
		Path:     "/",                            // Cookie path
		Expires:  time.Now().Add(24 * time.Hour), // Cookie expiration time
		HttpOnly: true,                           // Accessible only by the web server, helps mitigate XSS
		Secure:   true,                           // Ensure cookie is sent over HTTPS only
		SameSite: http.SameSiteStrictMode,        // Strict mode for CSRF mitigation
	})
}

func GenerateJWT(username string, privileges jwtM.Privileges) (string, error) {
	// Create an instance of CustomClaims
	claims := jwtM.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			// Expires 30 days from now
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			Issuer:    "go_app",
		},
		Username:   username,
		Privileges: privileges,
	}

	// Create a new JWT token using the HS256 signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret key
	secret := os.Getenv("SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		println(err.Error())
		return "", err
	}

	return tokenString, nil
}
func SendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
func SetPrivileges(user jwtM.CustomClaims) jwtM.Privileges {
	privileges := jwtM.Privileges{}
	if user.Privileges.Admin {
		privileges.Admin = true
	}
	if user.Privileges.User {
		privileges.User = true
	}
	if user.Privileges.Coach {
		privileges.Coach = true
	}
	return privileges
}
func JwtMiddlewareUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Cookie
		cookie, err := r.Cookie("auth_token")
		tokenString := cookie.Value
		if err != nil {
			SendJSONError(w, "No token provided", http.StatusUnauthorized)
			return
		}
		if tokenString == "" {
			SendJSONError(w, "No token provided", http.StatusUnauthorized)
			return
		}

		// Parse the token

		token, err := jwt.ParseWithClaims(tokenString, &jwtM.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			SendJSONError(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			SendJSONError(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*jwtM.CustomClaims); ok {
			if !claims.Privileges.User {
				SendJSONError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
func JwtMiddlewareAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		cookie, err := r.Cookie("auth_token")
		tokenString := cookie.Value
		if tokenString == "" {
			SendJSONError(w, "No token provided", http.StatusUnauthorized)
			return
		}

		// Parse the token

		token, err := jwt.ParseWithClaims(tokenString, &jwtM.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			SendJSONError(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			SendJSONError(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*jwtM.CustomClaims); ok {
			if !claims.Privileges.Admin {
				SendJSONError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
func JwtMiddlewareCoach(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		cookie, err := r.Cookie("auth_token")
		tokenString := cookie.Value
		if tokenString == "" {
			SendJSONError(w, "No token provided", http.StatusUnauthorized)
			return
		}

		// Parse the token

		token, err := jwt.ParseWithClaims(tokenString, &jwtM.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			SendJSONError(w, "Invalid token, problem with token"+err.Error()+tokenString, http.StatusUnauthorized)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			SendJSONError(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*jwtM.CustomClaims); ok {
			if !claims.Privileges.Coach {
				SendJSONError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
func CombinedJwtMiddleware(adminMiddleware, coachMiddleware func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Attempt to use the admin middleware
			adminPassed := false
			adminMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				adminPassed = true
				next.ServeHTTP(w, r)
			})).ServeHTTP(w, r)

			// If admin check didn't pass, try the coach middleware
			if !adminPassed {
				coachMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					next.ServeHTTP(w, r)
				})).ServeHTTP(w, r)
			}
		})
	}
}
