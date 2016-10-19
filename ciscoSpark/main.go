package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/jbogarin/go-cisco-spark/ciscospark"
)

func checkErr(msg string, err error) {
	if err != nil {
		panic(fmt.Sprintf("%s %v", msg, err))
	}
}

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	sparkClient := ciscospark.NewClient(client)
	token := "MTM4NmY3NGUtMjA5Mi00ZGZkLWEwYzktZjQwZDE4ZDc4NzBhMGY3ODYwNTMtNjEz" // Change to your token
	sparkClient.Authorization = "Bearer " + token

	// POST rooms
	roomRequest := &ciscospark.RoomRequest{
		Title: "Go Test Room",
	}

	newRoom, _, err := sparkClient.Rooms.Post(roomRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST:", newRoom.ID, newRoom.Title, newRoom.IsLocked, newRoom.Created)
	/*

		ROOMS

	*/
	// GET rooms
	var myRoomID string // Change to your testing room
	roomsQueryParams := &ciscospark.RoomQueryParams{
		Max:  2,
		Type: "group",
	}

	rooms, _, err := sparkClient.Rooms.Get(roomsQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	for id, room := range rooms {
		if id == 0 {
			myRoomID = room.ID
		}
		fmt.Println("GET:", id, room.ID, room.IsLocked, room.Title)
	}
	/*

		MESSAGES

	*/
	// myRoomID := "" // Change to your testing room
	//
	// POST messages - Text Message

	message := &ciscospark.MessageRequest{
		Text:   "This is a test message from a JRS script :)",
		RoomID: myRoomID,
	}
	newTextMessage, _, err := sparkClient.Messages.Post(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("POST:", newTextMessage.ID, newTextMessage.Text, newTextMessage.Created)
	//
	// // POST messages - Markdown Message
	//
	// markDownMessage := &ciscospark.MessageRequest{
	// 	MarkDown: "This is a markdown message. *Italic*, **bold** and ***italic/bold***.",
	// 	RoomID:   myRoomID,
	// }
	// newMarkDownMessage, _, err := sparkClient.Messages.Post(markDownMessage)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("POST:", newMarkDownMessage.ID, newMarkDownMessage.MarkDown, newMarkDownMessage.Created)
	//
	// // POST messages - Markdown Message
	//
	// htmlMessage := &ciscospark.MessageRequest{
	// 	MarkDown: "This is a html message with <strong>strong</strong>",
	// 	RoomID:   myRoomID,
	// }
	// newHTMLMessage, _, err := sparkClient.Messages.Post(htmlMessage)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("POST:", newHTMLMessage.ID, newHTMLMessage.MarkDown, newHTMLMessage.Created)
	//
	// // GET messages
	// messageQueryParams := &ciscospark.MessageQueryParams{
	// 	Max:    5,
	// 	RoomID: myRoomID,
	// }
	//
	// messages, _, err := sparkClient.Messages.Get(messageQueryParams)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for id, message := range messages {
	// 	fmt.Println("GET:", id, message.ID, message.Text, message.Created)
	// }
	//
	// // GET messages/<ID>
	//
	// htmlMessageGet, _, err := sparkClient.Messages.GetMessage(newHTMLMessage.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("GET <ID>:", htmlMessageGet.ID, htmlMessageGet.Text, htmlMessageGet.Created)
	//
	// // DELETE messages<ID>
	//
	// resp, err := sparkClient.Messages.DeleteMessage(newTextMessage.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("DELETE:", resp.StatusCode)

}
