package vinh

//unexported identifer since it started with the lower case
type alertCounter int

// New creates and returns values of the unexported
// type alertCounter.
func New(value int) alertCounter {
	return alertCounter(value)
}

//Newtimer ...
//identifer is exported since it defined with the Upper case
type Newtimer int
