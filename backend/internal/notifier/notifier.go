package notifier

import (
	"fmt"
	"log"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
)

type Notifier interface {
	SendNotification(r models.ChangeRecord, desc string, userEmail string) error
	FormatNotification(r models.ChangeRecord, desc string) string
}

type BasicNotifier struct {
}

func (b *BasicNotifier) SendNotification(r models.ChangeRecord, desc string) error {
	log.Print(b.FormatNotification(r, desc))
	return nil
}

func (b *BasicNotifier) FormatNotification(r models.ChangeRecord, desc string) string {
	formatted_addLogs := ""
	for i, v := range r.Added {
		formatted_addLogs += fmt.Sprintf("%d. %s\n", i+1, v.String_NameOnly())
	}
	return fmt.Sprintf("🚨 Job Board Updated: %q\n🔗 %q\n📝 Summary: %q\n📚 Detailed View:\n%s", desc, r.URL, r.DiffSummary, formatted_addLogs)
}
