/*
 * MailSlurp API
 *
 * For documentation see [developer guide](https://www.mailslurp.com/developers). [Create an account](https://app.mailslurp.com) in the MailSlurp Dashboard to [view your API Key](https://app). For all bugs, feature requests, or help please [see support](https://www.mailslurp.com/support/).
 *
 * API version: 0.0.1-alpha
 * Contact: contact@mailslurp.dev
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mailslurp

// Options for sending an email message from an inbox
type SendEmailOptions struct {
	// Optional list of attachment IDs to send with this email
	Attachments []string `json:"attachments,omitempty"`
	// Optional list of bcc destination email addresses
	Bcc []string `json:"bcc,omitempty"`
	// Contents of email
	Body string `json:"body,omitempty"`
	// Optional list of cc destination email addresses
	Cc []string `json:"cc,omitempty"`
	// Optional charset
	Charset string `json:"charset,omitempty"`
	// Optional from address. If not set source inbox address will be used
	From string `json:"from,omitempty"`
	Html bool `json:"html,omitempty"`
	// Optional replyTo header
	ReplyTo string `json:"replyTo,omitempty"`
	// Optional email subject line
	Subject string `json:"subject,omitempty"`
	// List of destination email addresses. Even single recipients must be in array form.
	To []string `json:"to"`
}
