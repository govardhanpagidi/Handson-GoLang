package main

type Event struct {
	CommndId    string `json:string`
	CommandName string `json:string`
	DeviceId    string `json:string`
	MessageId   string `json:string`
}

//
//func main(){
//
//	router := mux.NewRouter()
//
//	router.HandleFunc("/events",GetEvents).Methods("GET")
//
//}

func GetEvents(deviceId string) {
	GetEventsByDeviceId()
}
