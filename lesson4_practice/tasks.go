package practice

type Notifier interface {
	Send(message string) error
}

type BaseNotifier struct {
	Channel string
}

func (b BaseNotifier) Label(message string) string {
	return "[" + b.Channel + "] " + message
}

type EmailNotifier struct {
	BaseNotifier
	Address  string
	Messages []string
}

type PushNotifier struct {
	BaseNotifier
	DeviceID string
	Messages []string
}

func (n *EmailNotifier) Send(message string) error {
	n.Messages = append(n.Messages, n.Label(message))
	return nil
}

func (n *PushNotifier) Send(message string) error {
	n.Messages = append(n.Messages, n.Label(message))
	return nil
}

func Dispatch(all []Notifier, message string) error {

	for _, notifier := range all {
		if err := notifier.Send(message); err != nil {
			return err
		}
	}

	return nil
}

func EmailAddress(n Notifier) (string, bool) {
	if n == nil {
		return "", false
	}

	return Describe(n), true
}

func Describe(n Notifier) string {
	switch v := n.(type) {
	case *EmailNotifier:
		return v.Address
	case *PushNotifier:
		return v.DeviceID
	default:
		return "unknown"
	}
}
