package schemas


var SignupBody struct {
	Email string
	Password string
	FirstName string
	LastName string 
	DOB string
}

var LoginBody struct {
	Email string
	Password string 
}