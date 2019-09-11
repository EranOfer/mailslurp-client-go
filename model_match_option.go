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

type MatchOption struct {
	// The email property to match on. One of SUBJECT, TO, BCC, CC or FROM
	Field string `json:"field,omitempty"`
	// What criteria to apply. CONTAIN or EQUAL. Note CONTAIN is recommended due to some SMTP servers adding new lines
	Should string `json:"should,omitempty"`
	// The value to compare to the field using EQUAL or CONTAIN
	Value string `json:"value,omitempty"`
}
