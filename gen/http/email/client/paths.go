// Code generated by goa v3.20.0, DO NOT EDIT.
//
// HTTP request path constructors for the email service.
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package client

import (
	"fmt"
)

// ListTemplatesEmailPath returns the URL path to the email service list_templates HTTP endpoint.
func ListTemplatesEmailPath() string {
	return "/v1/email/templates"
}

// CreateTemplateEmailPath returns the URL path to the email service create_template HTTP endpoint.
func CreateTemplateEmailPath() string {
	return "/v1/email/templates"
}

// GetTemplateEmailPath returns the URL path to the email service get_template HTTP endpoint.
func GetTemplateEmailPath(id string) string {
	return fmt.Sprintf("/v1/email/templates/%v", id)
}

// GetTemplateByTypeEmailPath returns the URL path to the email service get_template_by_type HTTP endpoint.
func GetTemplateByTypeEmailPath(type_ string) string {
	return fmt.Sprintf("/v1/email/templates/by-type/%v", type_)
}

// UpdateTemplateEmailPath returns the URL path to the email service update_template HTTP endpoint.
func UpdateTemplateEmailPath(id string) string {
	return fmt.Sprintf("/v1/email/templates/%v", id)
}

// DeleteTemplateEmailPath returns the URL path to the email service delete_template HTTP endpoint.
func DeleteTemplateEmailPath(id string) string {
	return fmt.Sprintf("/v1/email/templates/%v", id)
}

// SendEmailPath returns the URL path to the email service send HTTP endpoint.
func SendEmailPath() string {
	return "/v1/email/send"
}

// SendTemplateEmailPath returns the URL path to the email service send_template HTTP endpoint.
func SendTemplateEmailPath() string {
	return "/v1/email/send-template"
}
