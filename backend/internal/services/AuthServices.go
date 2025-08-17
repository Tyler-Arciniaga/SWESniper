package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
)

type AuthService struct{}

func (s *AuthService) AddNewUser(u *models.SignUpUser) error {
	body, _ := json.Marshal(u)

	endpoint := fmt.Sprintf("https://%s.supabase.co/auth/v1/admin/users", os.Getenv("SUPABASE_PROJECT_REF"))
	r, _ := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(body))
	r.Header.Set("apikey", os.Getenv("SUPABASE_SERVICE_ROLE_KEY"))
	r.Header.Set("Authorization", "Bearer "+os.Getenv("SUPABASE_SERVICE_ROLE_KEY"))
	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error creating user: %v, response: %s", resp.Status, string(b))
	}
	return nil
}
