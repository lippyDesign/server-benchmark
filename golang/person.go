package main

// Person type is will be serialized into json and sent back to the client
type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ID        string `json:"id"`
	Phone     int    `json:"phone"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       int    `json:"zip"`
}
