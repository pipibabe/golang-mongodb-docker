package sentence

import (
	// "fmt"
	"math/rand"
)

// function字首大寫為public, 小寫為private
func Random() string {
    // A slice of message formats.
    formats := []string {
        "今天天氣很好",
        "大吉",
        "生日快樂",
        "貓咪大戰爭",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}