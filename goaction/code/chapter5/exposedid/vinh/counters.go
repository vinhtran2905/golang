package vinh

//unexported identifer since it started with the lower case
type alertCounter int

//Newimer ...
//identifer is exported since it defined with the Upper case
type Newimer float64

// New creates and returns values of the unexported
// type alertCounter.
func New(value int) alertCounter {
	return alertCounter(value)
}
