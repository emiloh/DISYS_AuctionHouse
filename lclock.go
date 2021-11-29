package main

type LamportClock struct {
	time int64
}

// Creates and returns a clock with a time set to 0.
func newClock() *LamportClock {
	clock := &LamportClock{
		time: 0,
	}
	return clock
}
func main() {}
// Returns the time of the lamport clock.
func (clock *LamportClock) now() int64 {
	return clock.time
}

// Compares two lamport clocks.
// It will set the lamport clock to the highest value of the two.
func (this *LamportClock) set(that *LamportClock) {
	if that.time >= this.time {
		this.time = that.time
	}
}
