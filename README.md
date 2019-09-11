# Go API client for mailslurp

> Requires Go 1.7+. `go get mailslurp/mailslurp-client-go`

For documentation see [developer guide](https://www.mailslurp.com/developers). [Create an account](https://app.mailslurp.com) in the MailSlurp Dashboard to [view your API Key](https://app). For all bugs, feature requests, or help please [see support](https://www.mailslurp.com/support/).

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1-alpha
- Package version: 4.8.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.mailslurp.com](https://www.mailslurp.com)

## Installation

Install the following dependencies:
```
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:
```golang
import "./mailslurp"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.mailslurp.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CommonOperationsApi* | [**CreateNewEmailAddress**](docs/CommonOperationsApi.md#createnewemailaddress) | **Post** /newEmailAddress | Create new email address
*CommonOperationsApi* | [**DeleteEmail**](docs/CommonOperationsApi.md#deleteemail) | **Delete** /deleteEmail | Delete an email
*CommonOperationsApi* | [**DeleteEmailAddress**](docs/CommonOperationsApi.md#deleteemailaddress) | **Delete** /deleteEmailAddress | Delete email address and its emails
*CommonOperationsApi* | [**EmptyInbox**](docs/CommonOperationsApi.md#emptyinbox) | **Delete** /emptyInbox | Delete all emails in an inbox
*CommonOperationsApi* | [**SendEmailSimple**](docs/CommonOperationsApi.md#sendemailsimple) | **Post** /sendEmail | Send an email from a random email address
*CommonOperationsApi* | [**WaitForEmailCount**](docs/CommonOperationsApi.md#waitforemailcount) | **Get** /waitForEmailCount | Wait for and return count number of emails 
*CommonOperationsApi* | [**WaitForLatestEmail**](docs/CommonOperationsApi.md#waitforlatestemail) | **Get** /waitForLatestEmail | Fetch inbox&#39;s latest email or if empty wait for email to arrive
*CommonOperationsApi* | [**WaitForMatchingEmail**](docs/CommonOperationsApi.md#waitformatchingemail) | **Post** /waitForMatchingEmails | Wait or return list of emails that match simple matching patterns
*CommonOperationsApi* | [**WaitForNthEmail**](docs/CommonOperationsApi.md#waitfornthemail) | **Get** /waitForNthEmail | Wait for or fetch the email with a given index in the inbox specified
*ExtraOperationsApi* | [**BulkCreateInboxes**](docs/ExtraOperationsApi.md#bulkcreateinboxes) | **Post** /bulk/inboxes | Bulk create Inboxes (email addresses)
*ExtraOperationsApi* | [**BulkDeleteInboxes**](docs/ExtraOperationsApi.md#bulkdeleteinboxes) | **Delete** /bulk/inboxes | Bulk Delete Inboxes
*ExtraOperationsApi* | [**BulkSendEmails**](docs/ExtraOperationsApi.md#bulksendemails) | **Post** /bulk/send | Bulk Send Emails
*ExtraOperationsApi* | [**CreateInbox**](docs/ExtraOperationsApi.md#createinbox) | **Post** /inboxes | Create an Inbox (email address)
*ExtraOperationsApi* | [**CreateWebhook**](docs/ExtraOperationsApi.md#createwebhook) | **Post** /inboxes/{inboxId}/webhooks | Attach a WebHook URL to an inbox
*ExtraOperationsApi* | [**DeleteEmail1**](docs/ExtraOperationsApi.md#deleteemail1) | **Delete** /emails/{emailId} | Delete Email
*ExtraOperationsApi* | [**DeleteInbox**](docs/ExtraOperationsApi.md#deleteinbox) | **Delete** /inboxes/{inboxId} | Delete Inbox / Email Address
*ExtraOperationsApi* | [**DeleteWebhook**](docs/ExtraOperationsApi.md#deletewebhook) | **Delete** /inboxes/{inboxId}/webhooks/{webhookId} | Delete and disable a WebHook for an Inbox
*ExtraOperationsApi* | [**DownloadAttachment**](docs/ExtraOperationsApi.md#downloadattachment) | **Get** /emails/{emailId}/attachments/{attachmentId} | Get email attachment
*ExtraOperationsApi* | [**GetEmail**](docs/ExtraOperationsApi.md#getemail) | **Get** /emails/{emailId} | Get Email Content
*ExtraOperationsApi* | [**GetEmails**](docs/ExtraOperationsApi.md#getemails) | **Get** /inboxes/{inboxId}/emails | List Emails in an Inbox / EmailAddress
*ExtraOperationsApi* | [**GetInbox**](docs/ExtraOperationsApi.md#getinbox) | **Get** /inboxes/{inboxId} | Get Inbox / EmailAddress
*ExtraOperationsApi* | [**GetInboxes**](docs/ExtraOperationsApi.md#getinboxes) | **Get** /inboxes | List Inboxes / Email Addresses
*ExtraOperationsApi* | [**GetRawEmailContents**](docs/ExtraOperationsApi.md#getrawemailcontents) | **Get** /emails/{emailId}/raw | Get Raw Email Content
*ExtraOperationsApi* | [**GetWebhooks**](docs/ExtraOperationsApi.md#getwebhooks) | **Get** /inboxes/{inboxId}/webhooks | Get all WebHooks for an Inbox
*ExtraOperationsApi* | [**SendEmail**](docs/ExtraOperationsApi.md#sendemail) | **Post** /inboxes/{inboxId} | Send Email
*ExtraOperationsApi* | [**UploadAttachment**](docs/ExtraOperationsApi.md#uploadattachment) | **Post** /attachments | Upload an attachment for sending
*ExtraOperationsApi* | [**UploadMultipartForm**](docs/ExtraOperationsApi.md#uploadmultipartform) | **Post** /attachments/multipart | Upload an attachment for sending using Multipart Form


## Documentation For Models

 - [BasicAuthOptions](docs/BasicAuthOptions.md)
 - [BulkSendEmailOptions](docs/BulkSendEmailOptions.md)
 - [CreateWebhookOptions](docs/CreateWebhookOptions.md)
 - [Email](docs/Email.md)
 - [EmailAnalysis](docs/EmailAnalysis.md)
 - [EmailPreview](docs/EmailPreview.md)
 - [Inbox](docs/Inbox.md)
 - [MatchOption](docs/MatchOption.md)
 - [MatchOptions](docs/MatchOptions.md)
 - [SendEmailOptions](docs/SendEmailOptions.md)
 - [UploadAttachmentOptions](docs/UploadAttachmentOptions.md)
 - [Webhook](docs/Webhook.md)


## Documentation For Authorization

## API_KEY
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

contact@mailslurp.dev

