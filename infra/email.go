package infra

import (
	"net/smtp"
	"os"

	"github.com/nokin-all-of-career/career-web-backend/configs"
)

// Auth is the global value that connect smtp server
var Auth smtp.Auth

// OperationEmail : email for notification
var OperationEmail string = "allofcareer@gmail.com"

// NewSMTPConnection is the func that create SMTP connection
func NewSMTPConnection() {

	serverAccount := configs.Config.ServerAccount
	if serverAccount == "" {
		serverAccount = os.Getenv("GMAIL_SERVER_ACCOUNT")
	}
	serverpassword := configs.Config.ServerAccountPW
	if serverpassword == "" {
		serverpassword = os.Getenv("GMAIL_SERVER_ACCOUNT_PW")
	}
	Auth = smtp.PlainAuth(
		"",
		serverAccount, // foo@gmail.com
		serverpassword,
		"smtp.gmail.com",
	)
}
