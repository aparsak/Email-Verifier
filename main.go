package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"net/mail"
	"strings"
)

type EmailCheckResult struct {
	Email       string
	ValidFormat string
	MXRecord    string
	EmailExists string
	ShowResult  bool
}

func emailVerifier(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}
	result := EmailCheckResult{
		ShowResult: false, // Initially, do not show results
	}
	err = t.Execute(w, result)
	if err != nil {
		fmt.Println("Error when executing template", err)
		return
	}
}

func emailChecker(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	// Use checkEmail to generate output for the browser
	result := checkEmail(email)

	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}

	// Display the result in the browser
	err = t.Execute(w, result)
	if err != nil {
		fmt.Println("Error when executing template", err)
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/Email":
		emailVerifier(w, r)
	case "/Email-check":
		emailChecker(w, r)
	default:
		fmt.Fprintf(w, "something went wrong")
	}
}

func checkEmail(email string) EmailCheckResult {
	result := EmailCheckResult{
		Email:      email,
		ShowResult: true, // Set to true to show results
	}

	_, err := mail.ParseAddress(email)
	if err == nil {
		result.ValidFormat = "The email address has a valid format."
	} else {
		result.ValidFormat = "The email address has an invalid format."
		return result
	}

	domain := strings.Split(email, "@")[1]
	if checkMXRecords(domain) {
		result.MXRecord = fmt.Sprintf("The domain '%s' has valid MX records.", domain)
	} else {
		result.MXRecord = fmt.Sprintf("The domain '%s' does not have valid MX records.", domain)
		return result
	}

	if checkEmailExists(domain) {
		result.EmailExists = "The email server for the domain exists and is reachable."
	} else {
		result.EmailExists = "The email server for the domain does not seem to exist or is unreachable."
	}

	return result
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
	return checkMXRecords(domain)
}

func main() {
	// Serve static files like CSS
	fs := http.FileServer(http.Dir("."))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
