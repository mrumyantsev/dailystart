package dailyt

import "time"

// SleepUntil sleeps until next daily task launch. It is a shorthand
// for calling time.Sleep function with the value from dailyt.Delay
// function.
func SleepUntil(dailyTime string) {
	delay := Delay(dailyTime)

	time.Sleep(delay)
}

// Delay gets a time duration before next daily task launch. A given
// dailyTime string parameter must fit following date format:
// "15:04:05", otherwize it will cause panic.
func Delay(dailyTime string) time.Duration {
	now := time.Now()
	nowYear, nowMonth, nowDay := now.Date()

	there, err := time.Parse(time.TimeOnly, dailyTime)
	if err != nil {
		// Panic, when dailyTime is not fit "15:04:05" format.
		panic(err)
	}

	thereDate := time.Date(
		nowYear,
		nowMonth,
		nowDay,
		there.Hour(),
		there.Minute(),
		there.Second(),
		0, // nsec
		now.Location(),
	)

	delta := now.Sub(thereDate)

	if delta < 0 {
		return -delta
	}

	const twentyFourHours = 86_400_000_000_000

	return twentyFourHours - delta
}
