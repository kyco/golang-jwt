package main

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// generateJWTToken generates a new, secure and unique token
func generateJWTToken(username *string) (tokenString string, err error) {
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["username"] = *username
	claims["expires"] = time.Now().Add(time.Minute * time.Duration(30)).Unix()

	/* Sign the token with our secret */
	tokenString, err = token.SignedString(JWTSecretKey)
	if err != nil {
		return tokenString, err
	}

	/* Finally, return the token */
	return tokenString, nil
}

// validateJWT validates JWT tokens
func validateJWT(jwtToken *string) (resp bool, err error) {

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(*jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Errorf("Unexpected signing method: %+v", token.Header["alg"])
			return false, errors.New("Unexpected signing method")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return JWTSecretKey, nil
	})

	// Check if the decoding of the JWT token
	// failed, if it did, return the error
	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Convert timestamp into an actual time
		ts := time.Unix(int64(claims["expires"].(float64)), 0)

		// Print all values in claims
		log.Debugf("%+v", claimts)

		// Check if the time has passed already
		if ts.Before(time.Now()) == true {
			return false, errors.New("Token has expired")
		}

	} else {
		log.Error(err.Error())
	}

	return true, nil
}

// validateLoginData checks if the correct username and password were supplied
func validateLoginData(username *string, password *string) bool {
	// Default demo data
	var (
		LoginData = LoginDataStruct{Username: "user", Password: "password"}
	)

	// Check if there is a match
	if *username == LoginData.Username && *password == LoginData.Password {
		return true
	}
	return false
}
