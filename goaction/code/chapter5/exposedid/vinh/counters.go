package vinh

//unexported identifer since it started with the lower case
type alertCounter int

// NewAlertCounter creates and returns values of the unexported
// type alertCounter.
func New(value int) alertCounter {
	return alertCounter(value)
}

//interalTimer ...
//identifer is exported since it defined with the Upper case
type interalTimer int

//Newtimer ...
func Newtimer(value int) interalTimer {
	return interalTimer(value)
}
