package main

type System struct{
	Services []ItService //All of the Services
}

//What type of Service We Have
const (
	SERVICE_ACCESS_POINT = 0
	SERVICE_ROUTER = 1
	SERVICE_PRINTER = 2
	SERVICE_SERVER = 3
)

func CreateAccessPoints() []ItService{
	service := []ItService{{Name: "Founder's Room AP", IpAddress: "10.1.0.6", Id: SERVICE_ACCESS_POINT, Status: true},{Name: "Room 2 AP", IpAddress: "10.1.0.7", Id: SERVICE_ACCESS_POINT, Status: true},{Name: "Cave AP", IpAddress: "10.1.0.8", Id: 2, Status: true}}
	return service
}

func CreateRouters() []ItService{
	service := []ItService{{Name: "Big House Router", IpAddress:"10.1.0.1", Id: SERVICE_ROUTER, Status: false},{Name: "Big House Ethernet Switch", IpAddress:"10.1.0.2", Id: SERVICE_ROUTER, Status: false}}
	return service
}

func CreatePrinters() []ItService{
	service := []ItService{{Name: "Founder's Room", IpAddress:"10.1.0.31", Id: SERVICE_PRINTER, Status: true}, {Name: "Cryso Printer", IpAddress:"10.1.0.37", Id: SERVICE_PRINTER, Status: true}}
	return service
}

func CreateServers() []ItService{
	service := []ItService{{Name: "Active Directory Server", IpAddress:"10.1.0.4", Id: SERVICE_SERVER, Status: true}, {Name: "Windows Server", IpAddress:"10.1.0.17", Id: SERVICE_SERVER, Status: true}}
	return service
}

func CreateSystem() System {
	s := System{Services:CreateAccessPoints()}
	s.Services = append(s.Services, CreateRouters()...)
	s.Services = append(s.Services, CreatePrinters()...)
	s.Services = append(s.Services, CreateServers()...)

	return s
}


