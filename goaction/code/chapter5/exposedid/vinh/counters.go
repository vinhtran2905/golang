package vinh

//unexported identifer since it started with the lower case
type newalertCounter int

// NewAlertCounter creates and returns values of the unexported
// type alertCounter.
func NewAlertCounter(value int) newalertCounter {
	return newalertCounter(value)
}

//Newtimer ...
//identifer is exported since it defined with the Upper case
type Newtimer int
