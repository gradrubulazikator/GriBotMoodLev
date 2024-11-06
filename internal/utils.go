package internal

import (
    "fmt"
    "net/http"
    "net/url"
)

const telegramAPI = "https://api.telegram.org/bot"

// SendMessage отправляет сообщение через API Telegram
func SendMessage(message string) error {
    endpoint := fmt.Sprintf("%s%s/sendMessage", telegramAPI, BotToken)
    data := url.Values{}
    data.Set("chat_id", ChatID)
    data.Set("text", message)

    resp, err := http.PostForm(endpoint, data)
    if err != nil {
        return fmt.Errorf("ошибка отправки сообщения: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("ошибка ответа от Telegram: %s", resp.Status)
    }
    return nil
}

