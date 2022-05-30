package core

import (
	"errors"
	"fmt"
	guuid "github.com/google/uuid"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
)

func BOD(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func GenerateCode() string {
	return  strings.ToUpper(guuid.NewString()[:7])
}

func Now() time.Time {

	return time.Now()
}

func ShortMonth(t time.Time) string {
	return t.Format("02 Jan 2006")
}

func Dare(passcode string, otp string)  bool {
	// TODO: check expired
	return passcode == otp
}

func VALIDATENumber( mobile string ) error {
	startsWithZero := strings.HasPrefix(mobile, "0")
	numberLength, err := strconv.Atoi(os.Getenv("NUMBER_LENGTH"))
	if err != nil { return errors.New("invalid number length") }
	validLength := len(mobile) == numberLength
	if !startsWithZero || !validLength {
		return errors.New("invalid mobile number")
	}
	return nil
}

func PrependCode(mobile string) string {
	m := strings.TrimPrefix(mobile, "0")
	return fmt.Sprintf("%v%v", os.Getenv("COUNTRY_CODE"), m)
}

func GenerateOTP() (string, error) {

	now := time.Now()
	secret := os.Getenv("OTP_SECRET")
	otpLength, err := strconv.Atoi(os.Getenv("OTP_LENGTH"))
	if err != nil {log.Error("Problem getting OTP Length"); return "0000", err}
	digits := otp.Digits(otpLength)
	counter := now.Add(100)
	opts := totp.ValidateOpts{Digits: digits}
	code, err := totp.GenerateCodeCustom(secret, counter, opts)
	if err != nil { log.Error("Problem Generating OTP: ", err.Error()); return "", err}
	log.Warn("[GenerateOTP]", code)

	return code, nil
}
