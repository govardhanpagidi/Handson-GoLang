package main

import (
	"fmt"
	"net/http"
	"github.com/RangelReale/osin"
)

//Source : https://github.com/RangelReale/osin

//How to run : run the below command
// go run osintest.go osinstorage.go

func main(){
	
	server := osin.NewServer(osin.NewServerConfig(),NewStorage())
	// 
	//Authorization end point
	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request){
		resp := server.NewResponse()
		defer resp.Close()

		if ar := server.HandleAuthorizeRequest(resp,r); ar != nil{

				//keep the backend validation for login
				ar.Authorized = true;
				server.FinishAuthorizeRequest(resp,r,ar)
			}
			osin.OutputJSON(resp,w,r)
		})

	//Access Token endpoint
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()

	if ar := server.HandleAccessRequest(resp, r); ar != nil {
			ar.Authorized = true
			server.FinishAccessRequest(resp, r, ar)
		}
	osin.OutputJSON(resp, w, r)
	})

	fmt.Println("server running on port 14000..")
	http.ListenAndServe(":14000",nil)
}	
