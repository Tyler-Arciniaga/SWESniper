package notifier

import (
	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
)

type DiscordNotifier struct{}

func (d *DiscordNotifier) SendNotification(r models.ChangeRecord, desc string) error {
	return nil
}

func (d *DiscordNotifier) FormatNotification(r models.ChangeRecord, desc string) string {
	return ""
}
