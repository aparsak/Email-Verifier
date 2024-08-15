package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/mail"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("email, validFormat, hasMX, validDomain, emailExists\n")

	for scanner.Scan() {
		email := scanner.Text()
		checkEmail(email)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v\n", err)
	}
}

func checkEmail(email string) {
	_, err := mail.ParseAddress(email)
	validFormat := err == nil

	if !validFormat {
		fmt.Printf("%s, %v, , , \n", email, validFormat)
		return
	}

	domain := strings.Split(email, "@")[1]
	hasMX := checkMXRecords(domain)
	validDomain := hasMX

	emailExists := false
	if hasMX {
		emailExists = checkEmailExists(domain)
	}

	fmt.Printf("%s, %v, %v, %v, %v\n", email, validFormat, hasMX, validDomain, emailExists)
}

func checkMXRecords(domain string) bool {
	_, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error looking up MX records for %s: %v\n", domain, err)
		return false
	}
	return true
}

func checkEmailExists(domain string) bool {
	// Simplified email existence check
	// Instead of trying to connect and send SMTP commands, just check if the domain has MX records.
	// Actual email existence check is complex and not always reliable due to server policies.
	return checkMXRecords(domain)
}
