# 📧 Email Validator Tool 

A simple Go application that validates email addresses through:
- Format checks
- MX (Mail Exchange) record verification
- Domain validity checks

## Demo
  
[Email Verifier.webm](https://github.com/user-attachments/assets/02176b4b-0d78-4640-93b3-4ca8a87ceedb)


## ✨ Features 

- **Email Format Validation**: Ensures the email is in a proper format.
- **MX Record Check**: Verifies if the domain has mail servers.
- **Domain Validity**: Confirms if the domain is valid for email.
- **Web Interface**: Provides a user-friendly web interface for email validation results.

## 🛠️ Usage 

1. **Navigate to the project directory:**
    ```bash
    cd email-validator
    ```

2. **Build and run the application:**
    ```bash
    go build -o email-validator
    ./email-validator
    ```

3. **Open your browser and go to** `http://localhost:8081` **to use the web interface.**

4. **Enter an email address in the form and click "Check" to view the validation results.**

## 📄 Web Interface 

- The web interface allows users to input email addresses and view results directly on the page.
- Results include email format validation, MX record checks, and domain validity.

## 🚀 Future Work 

- **Enhanced Validation**: Add more sophisticated email checks, such as syntax validation and disposable email detection.
- **SMTP Checks**: Implement direct SMTP checks to verify if an email address exists.
- **User Interface Improvements**: Further refine the web interface for a better user experience.
- **Internationalization**: Support for multiple languages and regional email formats.

Feel free to contribute or provide feedback!
