package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go/v3"
)

// Getausersaccountinformation : Get a user's account information.
// GET /user/account
func Getausersaccountinformation() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/account", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Retrieveyourcreditbalance : Retrieve your credit balance
// GET /user/credits
func Retrieveyourcreditbalance() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/credits", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Updateyouraccountemailaddress : Update your account email address
// PUT /user/email
func Updateyouraccountemailaddress() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/email", host)
	request.Method = "PUT"
	request.Body = []byte(` {
  "email": "example@example.com"
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Retrieveyouraccountemailaddress : Retrieve your account email address
// GET /user/email
func Retrieveyouraccountemailaddress() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/email", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Updateyourpassword : Update your password
// PUT /user/password
func Updateyourpassword() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/password", host)
	request.Method = "PUT"
	request.Body = []byte(` {
  "new_password": "new_password",
  "old_password": "old_password"
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Updateausersprofile : Update a user's profile
// PATCH /user/profile
func Updateausersprofile() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/profile", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "city": "Orange",
  "first_name": "Example",
  "last_name": "User"
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Getausersprofile : Get a user's profile
// GET /user/profile
func Getausersprofile() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/profile", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Cancelorpauseascheduledsend : Cancel or pause a scheduled send
// POST /user/scheduled_sends
func Cancelorpauseascheduledsend() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/scheduled_sends", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "batch_id": "YOUR_BATCH_ID",
  "status": "pause"
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Retrieveallscheduledsends : Retrieve all scheduled sends
// GET /user/scheduled_sends
func Retrieveallscheduledsends() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/scheduled_sends", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Updateuserscheduledsendinformation : Update user scheduled send information
// PATCH /user/scheduled_sends/{batch_id}
func Updateuserscheduledsendinformation() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/scheduled_sends/{batch_id}", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "status": "pause"
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Retrievescheduledsend : Retrieve scheduled send
// GET /user/scheduled_sends/{batch_id}
func Retrievescheduledsend() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/scheduled_sends/{batch_id}", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Deleteacancellationorpauseofascheduledsend : Delete a cancellation or pause of a scheduled send
// DELETE /user/scheduled_sends/{batch_id}
func Deleteacancellationorpauseofascheduledsend() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/scheduled_sends/{batch_id}", host)
	request.Method = "DELETE"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// UpdateEnforcedTLSsettings : Update Enforced TLS settings
// PATCH /user/settings/enforced_tls
func UpdateEnforcedTLSsettings() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/settings/enforced_tls", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "require_tls": true,
  "require_valid_cert": false
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// RetrievecurrentEnforcedTLSsettings : Retrieve current Enforced TLS settings.
// GET /user/settings/enforced_tls
func RetrievecurrentEnforcedTLSsettings() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/settings/enforced_tls", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Updateyourusername : Update your username
// PUT /user/username
func Updateyourusername() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/username", host)
	request.Method = "PUT"
	request.Body = []byte(` {
  "username": "test_username"
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Retrieveyourusername : Retrieve your username
// GET /user/username
func Retrieveyourusername() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/username", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// UpdateEventNotificationSettings : Update Event Notification Settings
// PATCH /user/webhooks/event/settings
func UpdateEventNotificationSettings() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/webhooks/event/settings", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "bounce": true,
  "click": true,
  "deferred": true,
  "delivered": true,
  "dropped": true,
  "enabled": true,
  "group_resubscribe": true,
  "group_unsubscribe": true,
  "open": true,
  "processed": true,
  "spam_report": true,
  "unsubscribe": true,
  "url": "url"
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// RetrieveEventWebhooksettings : Retrieve Event Webhook settings
// GET /user/webhooks/event/settings
func RetrieveEventWebhooksettings() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/webhooks/event/settings", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// TestEventNotificationSettings : Test Event Notification Settings
// POST /user/webhooks/event/test
func TestEventNotificationSettings() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/webhooks/event/test", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "url": "url"
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Createaparsesetting : Create a parse setting
// POST /user/webhooks/parse/settings
func Createaparsesetting() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/webhooks/parse/settings", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "hostname": "myhostname.com",
  "send_raw": false,
  "spam_check": true,
  "url": "http://email.myhosthame.com"
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Retrieveallparsesettings : Retrieve all parse settings
// GET /user/webhooks/parse/settings
func Retrieveallparsesettings() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/webhooks/parse/settings", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Updateaparsesetting : Update a parse setting
// PATCH /user/webhooks/parse/settings/{hostname}
func Updateaparsesetting() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/webhooks/parse/settings/{hostname}", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "send_raw": true,
  "spam_check": false,
  "url": "http://newdomain.com/parse"
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Retrieveaspecificparsesetting : Retrieve a specific parse setting
// GET /user/webhooks/parse/settings/{hostname}
func Retrieveaspecificparsesetting() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/webhooks/parse/settings/{hostname}", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Deleteaparsesetting : Delete a parse setting
// DELETE /user/webhooks/parse/settings/{hostname}
func Deleteaparsesetting() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/webhooks/parse/settings/{hostname}", host)
	request.Method = "DELETE"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// RetrievesInboundParseWebhookstatistics : Retrieves Inbound Parse Webhook statistics.
// GET /user/webhooks/parse/stats
func RetrievesInboundParseWebhookstatistics() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/webhooks/parse/stats", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "test_string"
	queryParams["start_date"] = "2016-01-01"
	queryParams["end_date"] = "2016-04-01"
	queryParams["offset"] = "test_string"
	request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func main() {
	// add your function calls here
}
