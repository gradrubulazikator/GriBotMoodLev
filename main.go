package main

import (
    "log"
    "time"
    "GriBotMoodLev/internal"
)

func moodReminder() {
    messages := []string{
        "Как ваше настроение сегодня?",
        "Помните: важно заботиться о себе.",
        "Если чувствуете усталость, дайте себе немного отдыха.",
    }

    for {
        for _, msg := range messages {
            if err := internal.SendMessage(msg); err != nil {
                log.Printf("Ошибка отправки напоминания: %v", err)
            }
            time.Sleep(internal.RemindDelay)
        }
    }
}

func main() {
    log.Println("Запуск GriBotMoodLev...")
    go moodReminder()
    select {}
}

