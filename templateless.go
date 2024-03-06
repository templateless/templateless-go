package templateless

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ObjectId = string

type EmailResponsePreview struct {
	Email   string `json:"email"`
	Preview string `json:"preview"`
}

type EmailResponse struct {
	Emails   []ObjectId              `json:"emails"`
	Previews *[]EmailResponsePreview `json:"previews"`
}

type Templateless struct {
	apiKey string
	domain string
}

func NewTemplateless(apiKey string) *Templateless {
	domain := os.Getenv("TEMPLATELESS_DOMAIN")
	if domain == "" {
		domain = "https://api.templateless.com"
	}

	return &Templateless{
		apiKey: apiKey,
		domain: domain,
	}
}

func (t *Templateless) Domain(domain string) *Templateless {
	t.domain = domain
	return t
}

func (t *Templateless) Send(email Email) ([]ObjectId, *CustomError) {
	return t.SendMany([]Email{email})
}

func (t *Templateless) SendMany(emails []Email) ([]ObjectId, *CustomError) {
	emailSlices := make([]EmailSlice, 0)
	for _, email := range emails {
		emailSlices = append(emailSlices, *email.ToSlice())
	}

	payload, err := json.Marshal(map[string][]EmailSlice{"payload": emailSlices})
	if err != nil {
		return nil, UnknownError()
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/emails", t.domain), bytes.NewBuffer(payload))
	if err != nil {
		return nil, UnknownError()
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.apiKey))
	req.Header.Set("Referer", "templateless-go v0")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, UnknownError()
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, UnknownError()
	}

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return nil, UnauthorizedError()
	case http.StatusForbidden:
		return nil, ForbiddenError()
	case http.StatusUnprocessableEntity:
		type InvalidParameter struct {
			Field string `json:"field"`
		}

		var e InvalidParameter
		if err := json.Unmarshal(body, &e); err == nil {
			return nil, InvalidParameterError(e.Field)
		}

		return nil, UnknownError()
	case http.StatusBadRequest:
		type BadRequest struct {
			Code  BadRequestCode `json:"code"`
			Error string         `json:"error"`
		}

		var e BadRequest
		if err := json.Unmarshal(body, &e); err == nil {
			return nil, BadRequestError(e.Code, e.Error)
		}

		return nil, UnknownError()
	case http.StatusInternalServerError:
		return nil, UnavailableError()
	case http.StatusOK:
		var emailResponse EmailResponse
		if err := json.Unmarshal(body, &emailResponse); err != nil {
			return nil, UnknownError()
		}

		if emailResponse.Previews != nil {
			for _, preview := range *emailResponse.Previews {
				log.Printf("Templateless [TEST MODE]: Emailed %s, preview: https://tmpl.sh/%s\n",
					preview.Email, preview.Preview)
			}
		}

		return emailResponse.Emails, nil
	default:
		return nil, UnknownError()
	}
}
