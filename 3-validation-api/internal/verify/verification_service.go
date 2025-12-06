package verify

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/ANB98prog/purple-school-homeworks/3-validation-api/configs"
	"github.com/ANB98prog/purple-school-homeworks/3-validation-api/pkg/json"
	"github.com/jordan-wright/email"
	"hash/fnv"
	"net/smtp"
	"os"
)

type VerificationService interface {
	SendVerification(emailAddress string) error
	Verify(hash string) bool
}

const template = `
		<html>
		<body>
			<h1>Подтверждение почты</h1>
			<p>Подтвердите электронную почту перейдя по ссылке: <a href="url">%s</a></p>
		</body>
		</html>
	`

const fileName = "verifications.json"

type EmailVerificationService struct {
	config configs.EmailSenderConfig
}

func NewEmailVerificationService(config configs.EmailSenderConfig) *EmailVerificationService {
	return &EmailVerificationService{config: config}
}

func (sender *EmailVerificationService) SendVerification(emailAddress string) error {
	e := email.NewEmail()
	e.Subject = "Подтверждение электронной почты"
	e.To = []string{emailAddress}
	e.From = sender.config.From
	auth, err := auth(sender.config)
	if err != nil {
		return err
	}

	hash := getHash(emailAddress)

	if err := upsertVerification(hash, emailAddress); err != nil {
		return err
	}

	link := fmt.Sprintf("%s/verify/%s", sender.config.ApiAddress, hash)
	body := fmt.Sprintf(template, link)

	e.HTML = []byte(body)

	return e.SendWithTLS(fmt.Sprintf("%s:%s", sender.config.SmtpAuth.Host, sender.config.SmtpAuth.Port), auth, &tls.Config{
		ServerName: sender.config.SmtpAuth.Host,
	})
}

func auth(config configs.EmailSenderConfig) (smtp.Auth, error) {
	auth := smtp.PlainAuth("", config.SmtpAuth.Login, config.SmtpAuth.Password, config.SmtpAuth.Host)

	return auth, nil
}

func getHash(emailAddress string) string {
	h := fnv.New32a()
	h.Write([]byte(emailAddress))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))[:8]
}

func (sender *EmailVerificationService) Verify(hash string) bool {
	ok := isCorrectVerification(hash)
	if !ok {
		return false
	}

	// если совпал удаляем хэш из списка
	deleteVerification(hash)

	return true
}

func upsertVerification(hash string, emailAddress string) error {
	verifications, err := getVerifications()
	if err != nil {
		return err
	}

	verifications[hash] = emailAddress

	return saveVerifications(verifications)
}

func deleteVerification(hash string) error {
	verifications, err := getVerifications()
	if err != nil {
		return err
	}

	delete(verifications, hash)

	return saveVerifications(verifications)
}

func isCorrectVerification(hash string) bool {
	items, err := getVerifications()
	if err != nil {
		return false
	}

	_, ok := items[hash]
	return ok
}

func getVerifications() (EmailVerificationItems, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return EmailVerificationItems{}, err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return EmailVerificationItems{}, err
	}
	if info.Size() == 0 {
		return EmailVerificationItems{}, nil
	}

	verifications, err := json.Decode[EmailVerificationItems](file)
	if err != nil {
		return nil, fmt.Errorf("verification file has invalid JSON structure. Error: %w", err)
	}

	return verifications, nil
}

func saveVerifications(verifications EmailVerificationItems) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	if err = json.Encode(file, verifications); err != nil {
		return err
	}

	return nil
}
