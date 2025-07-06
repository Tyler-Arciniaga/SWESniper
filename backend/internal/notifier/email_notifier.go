package notifier

import (
	"fmt"
	"os"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailNotifier struct{}

// TODO: eventually buy domain of swesniper.com to make email more legitamate and less likely to be sent to spam
func (e *EmailNotifier) SendNotification(r models.ChangeRecord, desc string) error {
	from := mail.NewEmail("SWE Sniper", "tyarciniaga@gmail.com")
	to := mail.NewEmail("Ty", "tylerarc@umich.edu")
	plainTextContent := e.FormatNotification(r, desc)
	//htmlContent := e.FormatNotification(r, desc)
	subject := fmt.Sprintf("Job Board Updated: %q", desc)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, "")
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	if err != nil {
		return err
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
