package main

type ItService struct{
	Name string
	IpAddress string
	Id uint
	Status bool
}

func CreateAccessPoints() []ItService{
	service := []ItService{{Name: "Big House AP 1", IpAddress: "10.1.0.6", Id: 1, Status: true},{Name: "Big House AP 2", IpAddress: "10.1.0.7", Id: 2, Status: true}}
	return service
}