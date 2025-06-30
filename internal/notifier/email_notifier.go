package notifier

import (
	"fmt"
	"log"
	"os"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailNotifier struct{}

// TODO: eventually buy domain of swesniper.com to make email more legitamate and less likely to be sent to spam
func (e *EmailNotifier) SendNotification(r models.ChangeRecord, desc string) error {
	from := mail.NewEmail("SWE Sniper", "tyarciniaga@gmail.com")
	to := mail.NewEmail("Ty", "tyarciniaga@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := e.FormatNotification(r, desc)
	subject := fmt.Sprintf("Job Board Updated: %q", desc)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return nil
}

func (e *EmailNotifier) FormatNotification(r models.ChangeRecord, desc string) string {
	formatted_addLogs := ""
	for i, v := range r.Added {
		formatted_addLogs += fmt.Sprintf("%d. %s\n", i+1, v.String_NameOnly())
	}
	return fmt.Sprintf("🚨 Job Board Updated: %s\n🔗 %s\n📝 Summary: %s\n📚 Detailed View:\n%s", desc, r.URL, r.DiffSummary, formatted_addLogs)
}
