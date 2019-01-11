package capacitor

import (
    "net/http"
    "os"
)

var BASE_API = getEnv("BASE_API", "http://localhost:8000/")
var API = BASE_API + "api/v1/"
var client = &http.Client{}
var email string = getEnv("FLUX_EMAIL", "aeon@localhost")
var password string = getEnv("FLUX_PASSWORD", "password123")
var token string = getEnv("FLUX_TOKEN", "")


// Provided an environment variable key, return its value if it exists,
// otherwise return the fallback value.
func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
