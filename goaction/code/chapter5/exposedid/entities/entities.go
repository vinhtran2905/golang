package entities

// type User struct {
// 	name  string //unexported field
// 	Email string //exported field
// }
type user struct {
	Name  string
	Email string
}

//the identitifers from the inner type are promoted to the outer type, those exported fields from user type are know through a value of the outer type
type Admin struct {
	user //user is unexported field
	Wage float64
}
