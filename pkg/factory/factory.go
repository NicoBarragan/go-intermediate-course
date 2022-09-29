package factory

import "fmt"

/* SMS and Email notification msg sending */

// interface that defines how must be composed
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender // func that returns an interface
}

// ISender interface that must be integrated for being INotificationFactory compliant
type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

/* SMS */
// We create the first struct for SMS
type SMSNotification struct{}

// Create the first func in the interface
func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}

// We create the func that returns the struct with the ISender func's
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

// Create the struct that will implement the ISender func's
type SMSNotificationSender struct {
}

// Create on of the ISender func's on the ISender struct
func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

// Create the other ISender func on the ISender struct
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

/* Email */
// Create the Email notification struct
type EmailNotification struct {
}

// Create first INotificationSender func
func (EmailNotification) SendNotification() {
	fmt.Println("Sending notification via email")
}

// Create the struct that will implement the ISender func's
type EmailNotificationSender struct {
}

// Create the func that will return the ISender struct
func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

// Create on of the ISender func's on the ISender struct
func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

// Create the other ISender func on the ISender struct
func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

//////////////////////

// Define func for getting the notification factory type
// Returns the factory or an error
func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("no notification type")
}

// Sends notification of the factory that is passed
func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

// Gets the method of the sender from the factory that is passed
func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

// Declares both SMS and Email factories and executes their methods (public func)
func Factory() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
