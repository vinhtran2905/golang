package vinh

//unexported identifer since it started with the lower case
type alertCounter int

//Timer ...
//identifer is exported since it defined with the Upper case
type Timer float64

// New creates and returns values of the unexported
// type alertCounter.
func New(value int) alertCounter {
	return alertCounter(value)
}
