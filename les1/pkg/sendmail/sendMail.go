package sendmail

import (
	"awesomeProject/GoArchitecture/gb-go-architecture-master/lesson-2/shop/models"
	"fmt"
	"net/smtp"
)

// smtpServer data to smtp server.
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server.
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

type SendMailOrder interface {
	SendOrderNotification(order *models.Order) error
}

func SendMail() {
	// Sender data.
	from := "ya@yandex.ru"
	password := "pass"

	// Receiver email address.
	to := []string{"ya@yandex.ru"}

	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.yandex.ru", port: "25"}

	// Message.
	id := "Test"
	message := []byte(fmt.Sprintf("new order %d\n", id))

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)

	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
