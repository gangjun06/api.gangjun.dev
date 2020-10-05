package req

type SendEmailToMe struct {
	Captcha string
	title   string
	email   string
	text    string
}
