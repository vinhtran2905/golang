//use go mod init github.com/vintran2905/golang ('root folder of the project where contain the .git')
//from internal package can call another internal package by manually import "github.com/vinhtran2905/golang/<subdirectory>/<internal package name>"
package vinh

//unexported identifer since it started with the lower case
type alertCounter int

// New creates and returns values of the unexported type alertCounter.
//This function is called FACTORY function which is used to create a new instance
//It is a Go convention to give factory function the name of New
func New(value int) alertCounter {
	return alertCounter(value)
}

//interalTimer ...
//identifer is exported since it defined with the Upper case
type interalTimer float64

//Newtimer ...
func Newtimer(value float64) interalTimer {
	return interalTimer(value)
}

//Export identifer
type NewString string
