package main

import "fmt"


	type notification interface {
		Send() string
	}

	//for sms , email , push 
	type sms struct {
		message string
	}

	func (s sms) Send() string {
		return fmt.Sprintf("Sending SMS with message: %s", s.message)
	}

	type email struct {
		name string
	}

	func (e email) Send() string {
		return fmt.Sprintf("Sending Email to: %s", e.name)
	}

	type push struct {
		message string
	}

	func (p push) Send() string {
		return fmt.Sprintf("Sending Push Notification with message: %s", p.message)
	}

func send(n notification) {
	fmt.Println(n.Send())
}

func main () {
	smsNotification := sms{message: "Hello via SMS!"}
	emailNotification := email{name: "John Doe"}
	pushNotification := push{message: "Hello via Push Notification!"}

	send(smsNotification)
	send(emailNotification)
	send(pushNotification)
}