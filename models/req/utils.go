package req

type SendEmailToMe struct {
	Captcha string `binding:"required"`
	Title   string `binding:"required"`
	Email   string `binding:"required"`
	Text    string `binding:"required"`
}
