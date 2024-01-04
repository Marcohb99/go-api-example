package release

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	apiExample "github.com/marcohb99/go-api-example/internal"
)

func SampleReleaseCollection(size int) []apiExample.Release {
	var result []apiExample.Release

	for i := 0; i < size; i++ {
		releaseObj, _ := apiExample.NewRelease(
			uuid.NewString(),
			generateRandomString(10),
			generateRandomDate(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2022, 12, 31, 23, 59, 59, 999999999, time.UTC)),
			"https://example.com/" + generateRandomString(10),
			"https://example.com/" + generateRandomString(10),
			strconv.Itoa(rand.Intn(5000)),
		)
		result = append(result, releaseObj)	
	}
	return result
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func generateRandomDate(start, end time.Time) string {
	if start.After(end) {
		panic("start date cannot be after end date")
	}

	// Calculate the duration between start and end dates
	duration := end.Sub(start)

	// Generate a random duration within the specified range
	randomDuration := time.Duration(rand.Int63n(int64(duration)))

	// Add the random duration to the start date to get a random date
	randomDate := start.Add(randomDuration)

	return randomDate.Format("YYYY-mm-dd")
}