package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// https://bcrypt-generator.com/

// The cost factor (in this case, 14) is a measure of how
// computationally expensive the hashing should be. Higher cost
// factors result in slower hash generation but also make it more
// difficult for attackers to perform brute-force or rainbow table attacks.

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	fmt.Println(string(bytes))
	return string(bytes), err
}
