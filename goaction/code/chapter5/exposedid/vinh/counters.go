package vinh

type alertCounter int

// New creates and returns values of the unexported
// type alertCounter.
func New(value int) alertCounter {
	return alertCounter(value)
}
