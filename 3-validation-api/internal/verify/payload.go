package verify

type EmailVerification struct {
	Email string `json:"email" validate:"required,email"`
}

type EmailVerificationItems map[string]string
