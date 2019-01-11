package capacitor


func Discharge() {
    // If we don't already have a token, get one.
    if token == "" {
      token = "Token " + getToken(email, password)
    }
}
