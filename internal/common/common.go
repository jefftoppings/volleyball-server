package common

import "math/rand"

type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func ShuffleArray[T any](arr []T) {
    n := len(arr)
    for i := n - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        arr[i], arr[j] = arr[j], arr[i]
    }
}