package server

import (
  "context"
  "ecomm-backend/internal/auth"
  "ecomm-backend/internal/database"
  "net/http"

  "github.com/gin-contrib/sessions"
  "github.com/gin-gonic/gin"
)


func (s *Server) sendSMS(c *gin.Context) {
  _, cancel := context.WithTimeout(context.Background(), appTimeout)
  var payload auth.OTPData
  defer cancel()

  auth.ValidateBody(c, &payload)

  newData := auth.OTPData{
    PhoneNumber: payload.PhoneNumber,
  }
  // for development
  if appEnv == "development" {
    c.JSON(http.StatusAccepted, gin.H{"success": "OTP sent successfully"})
    return
  }

  _, err := auth.TwilioSendOTP(newData.PhoneNumber)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    return
  }


  c.JSON(http.StatusAccepted, gin.H{"success": "OTP sent successfully"})
}

func (s *Server) verifySMS(c *gin.Context) {
  _, cancel := context.WithTimeout(context.Background(), appTimeout)
  var payload auth.VerifyData
  defer cancel()

  // Validate the request's body
  auth.ValidateBody(c, &payload)

  newData := auth.VerifyData{
    PhoneNumber: payload.PhoneNumber,
    Code: payload.Code,
  }
  // for development
  if appEnv == "development" {
    IDToken, err := auth.CreateIDToken("+77776134709")
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusOK, gin.H{"token": IDToken})
    return
  }

  // request twilio to check if phone number verified
  err := auth.TwilioVerifyOTP(newData.PhoneNumber, newData.Code)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    return
  }

  // create ID token that proves the verification
  // this is used further for sign in or sign up
  token, err := auth.CreateIDToken(newData.PhoneNumber)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"token": token})
}

func (s *Server) signIn(c *gin.Context) {
  var user database.User
  session := sessions.Default(c)
  // TODO: avoid creating new session; needs test
  err := auth.RetrieveSession(session, &user)
  if err != nil {
    c.JSON(http.StatusOK, gin.H{"success": "session already exists"})
    return
  }

  if user.ID.String() == "00000000-0000-0000-0000-000000000000" {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "no user found"})
    return
  }

  // verify that the user passed verification
  IDClaims, err := auth.VerifyUser(c)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err})
    return
  }

  // try to sign in and create session. If error, try to sign up
  user.PhoneNumber = IDClaims.PhoneNumber

  ok := auth.SignInUser(&user)
  if !ok {
    c.JSON(http.StatusNotFound, gin.H{"error": "didn't found user. Signup"})
    return
  }

  err = auth.CreateSession(session, &user)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err})
    return
  }
  c.JSON(http.StatusCreated, gin.H{"success": "session created"})
}

func (s *Server) signUp(c *gin.Context) {
  // check if user is verified
  IDClaims, err := auth.VerifyUser(c)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err})
    return
  }

  var user database.User
  user.PhoneNumber = IDClaims.PhoneNumber

  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // try to sign up user
  if err := auth.SignUpUser(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"success": "signin to proceed"})
}

func (s *Server) signOut(c * gin.Context) {
  session := sessions.Default(c)
  auth.DeleteSession(session)

  c.JSON(http.StatusOK, gin.H{"success": "session removed"})
}
