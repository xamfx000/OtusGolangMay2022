package pkg

func IntervalsIntersect(firstStart, firstEnd, secondStart, secondEnd int64) bool {
	return (secondStart >= firstStart && secondStart <= firstEnd) ||
		(secondEnd >= firstStart && secondEnd <= firstEnd)
}
