package main

import (
	"errors"
	"fmt"

	homework "lesson5_homework"
)

const message = "заказ готов"

func main() {
	demoSuccess()
	demoEmptyMessage()
	demoMissingConfiguration()
	demoNotifierPanic()
}

func demoSuccess() {
	email := &homework.EmailNotifier{
		Channel: "email",
		Address: "a@a.com",
	}
	push := &homework.PushNotifier{
		Channel:  "push",
		DeviceID: "67676767",
	}
	channels := []homework.Notifier{email, push}

	fmt.Println("1. Успешная отправка")
	if err := homework.Dispatch(channels, message); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Ошибки нет")
	}
	fmt.Printf("Email: %q\nPush: %q\n", email.Messages, push.Messages)

}

func demoEmptyMessage() {
	email := &homework.EmailNotifier{
		Channel: "email",
		Address: "a@a.com",
	}
	push := &homework.PushNotifier{
		Channel:  "push",
		DeviceID: "6767676767",
	}
	channels := []homework.Notifier{email, push}

	fmt.Println("\n2. Пустой текст")
	fmt.Printf("Истории до попытки — Email: %q, Push: %q\n", email.Messages, push.Messages)
	err := homework.Dispatch(channels, "")
	if errors.Is(err, homework.ErrEmptyMessage) {
		fmt.Println("Пустое сообщение")
	} else {
		fmt.Println("Ошибка:", err)
	}
	fmt.Printf("Истории после попытки — Email: %q, Push: %q\n", email.Messages, push.Messages)

}

func demoMissingConfiguration() {
	fmt.Println("\n3. Нет настройки")
	unconfigured := &homework.PushNotifier{
		Channel: "push",
	}
	err := homework.Dispatch([]homework.Notifier{unconfigured}, message)
	var sendErr *homework.SendError
	if errors.As(err, &sendErr) {
		fmt.Printf("%s: %s\n", sendErr.Channel, sendErr.Reason)
	} else {
		fmt.Println("Ошибка:", err)
	}
	fmt.Printf("История push: %q\n", unconfigured.Messages)

}

func demoNotifierPanic() {
	fmt.Println("\n4. Паника одного канала")
	broken := &homework.BrokenNotifier{
		Channel: "broken",
	}
	nextPush := &homework.PushNotifier{
		Channel:  "push",
		DeviceID: "676767676767",
	}
	failures := homework.SafeDispatch([]homework.Notifier{broken, nextPush}, message)
	fmt.Printf("Количество ошибок: %d\n", len(failures))
	for _, err := range failures {
		if errors.Is(err, homework.ErrNotifierPanic) {
			fmt.Println("Паника канала")
		} else {
			fmt.Println("Ошибка:", err)
		}
	}
	fmt.Printf("Следующий push получил: %q\n", nextPush.Messages)
}
