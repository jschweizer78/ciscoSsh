package main

/*

The goal of this app is to launch an API server that when triggered will read from Firebase and Google sheets/docs
for information and then SSH to cisco devices to login and apply configuration changes. The UI for this API server will be an Angular2
web app that will be hosted at Firebase. This executeable will launch a browser window to the hosted web app. The server Should
launch the API service and the web page and wait for a kill signal.

*/

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jschweizer78/ciscoSsh/resource"
	"github.com/jschweizer78/ciscoSsh/storage/tiedot"
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go-adapter/gingonic"
	"github.com/manyminds/api2go/examples/model"
)

func main() {
	r := gin.Default()

	api := api2go.NewAPIWithRouting(
		"api",
		api2go.NewStaticResolver("/"),
		gingonic.New(r),
	)

	// url := "https://project-1263673411523801245.firebaseio.com/people"
	// // Can also be your Firebase secret:
	// //authToken := "AIzaSyDLVe2jtGSarILB04EHXFUZ9uSegUAz9_A"
	// fireRef := firebase.NewReference(url).Export(false)
	// ds := service.NewService(r, fireRef)
	// addRoutes(ds)
	// p := models.NewPerson("Jason", "Schweizer")
	// err := ds.Firebase.Value(p)
	// if err != nil {
	// 	panic(err)
	// }
	userStorage := tiedot.NewUserStorage()
	defer tiedot.GetDB().Close()
	api.AddResource(model.User{}, resource.UserResource{UserStorage: userStorage})
	fmt.Println(p.DisplayName())
	r.Run("localhost:8080")
}
