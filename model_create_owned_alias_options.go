/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: d1659dc1567a5b62f65d0612107a50aace03e085
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mailslurp
// CreateOwnedAliasOptions Create email alias options. Email aliases can be used to mask real email addresses behind an ID. You can also attach an inbox to an alias so that any email received by the inbox email address if forwarded to the alias email address.
type CreateOwnedAliasOptions struct {
	// Email address to be hidden behind alias
	EmailAddress string `json:"emailAddress,omitempty"`
	// Optional inbox ID to attach to alias. Emails received by this inbox will be forwarded to the alias email address
	InboxId string `json:"inboxId,omitempty"`
	// Optional name for alias
	Name string `json:"name,omitempty"`
	// Optional proxied flag. When proxied is true alias will forward the incoming emails to the aliased email address via a proxy inbox. A new proxy is created for every new email thread. By replying to the proxy you can correspond with using your email alias without revealing your real email address.
	Proxied bool `json:"proxied,omitempty"`
}
