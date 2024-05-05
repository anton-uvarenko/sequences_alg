package core

import "math/rand"

func RandomizeArray[E any](arr []E) {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}
