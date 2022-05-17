package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var total int = 0;
	for _,bird := range birdsPerDay {
		total += bird
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekOffset := 7 * (week - 1)
	return TotalBirdCount(birdsPerDay[weekOffset:weekOffset+7])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {	
	for i := range birdsPerDay {
		if i == 0 || i%2 == 0 {
			birdsPerDay[i] += 1 
		}
	}
	return birdsPerDay
}
