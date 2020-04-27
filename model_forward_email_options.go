/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: d1659dc1567a5b62f65d0612107a50aace03e085
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mailslurp
// ForwardEmailOptions Options for forwarding an email
type ForwardEmailOptions struct {
	// Optional bcc recipients
	Bcc []string `json:"bcc,omitempty"`
	// Optional cc recipients
	Cc []string `json:"cc,omitempty"`
	// Subject for forwarded email
	Subject string `json:"subject,omitempty"`
	// To recipients for forwarded email
	To []string `json:"to,omitempty"`
}
