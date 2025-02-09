package auth

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)


var (
    accountSID  string              = os.Getenv("TWILIO_ACCOUNT_SID")
    authToken   string              = os.Getenv("TWILIO_AUTH_TOKEN")
    servicesID  string              = os.Getenv("TWILIO_SERVICES_ID")
    twilioClient *twilio.RestClient
)

type OTPData struct {
    PhoneNumber string `json:"phone_number,omitempty" validate:"required"`
}

type VerifyData struct {
    PhoneNumber string `json:"phone_number,omitempty" validate:"required"`
    Code string `json:"code,omitempty" validate:"required"`
}


func NewTwilioClient() {
    if twilioClient != nil {
        return
    }
    twilioClient = twilio.NewRestClientWithParams(twilio.ClientParams{
        Username: accountSID,
        Password: authToken,
    })
}

func TwilioSendOTP(phoneNumber string) (string, error) {
    NewTwilioClient()

    params := &twilioApi.CreateVerificationParams{}
    params.SetTo(phoneNumber)
    params.SetChannel("sms")

    resp, err := twilioClient.VerifyV2.CreateVerification(servicesID, params)
    if err != nil {
        return "", err
    }

    return *resp.Sid, nil
}

func TwilioVerifyOTP(phoneNumber string, code string) error {
    NewTwilioClient()

    params := &twilioApi.CreateVerificationCheckParams{}
    params.SetTo(phoneNumber)
    params.SetCode(code)

    resp, err := twilioClient.VerifyV2.CreateVerificationCheck(servicesID, params)
    if err != nil {
        return err
    } else if *resp.Status == "approved" {
        return nil
    }

    return nil
}
