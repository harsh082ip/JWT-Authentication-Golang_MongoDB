package helpers

import (
	"fmt"
	"strings"

	emailverifier "github.com/AfterShip/email-verifier"
)

var (
	verifier = emailverifier.NewVerifier().DisableCatchAllCheck()
)

func VerifyEmail(email string) (bool, error) {

	ret, err := verifier.Verify(email)
	if err != nil {
		return false, fmt.Errorf("Verification Failedddd: %v", err)
	}

	if !ret.Syntax.Valid {
		return false, fmt.Errorf("Verification failed: %v", "Syntax Invalid")
	}

	// Get Username and Domain
	// username, domain, err := SplitEmail(email)
	// if err != nil {
	// 	return false, err
	// }

	// _, err = verifier.CheckSMTP(domain, username)
	// if err != nil {
	// 	return false, fmt.Errorf("Email Invalid, SMTP check failed: %v", err)
	// }

	return true, nil
}

func SplitEmail(email string) (string, string, error) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("Email Address Invalid: %v", "split error")
	}

	return parts[0], parts[1], nil
}
