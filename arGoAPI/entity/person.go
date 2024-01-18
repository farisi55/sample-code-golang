package entity

//Person object for REST(CRUD)
type Person struct {
	ID	int `json:"id"`
	FirstName	string `json:"first_name"`
	LastName	string `json:"last_name"`
	Age	int `json:"age"`
}