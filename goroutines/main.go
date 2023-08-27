package main

import (
	"fmt"
	_ "sync"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch chan Message
}

func (s *Server) StartandListen() {
	// time.Sleep(40 * time.Millisecond)
	for {
		msg := <-s.msgch
		fmt.Printf("Received message from %s and payload %s\n", msg.From, msg.Payload)
	}
}

func sendMessage(msgch chan Message, payload string) {
	msg := Message{
		From:    "JOEBIDEN",
		Payload: payload,
	}
	msgch <- msg
}

func main() {
	s := &Server{
		msgch: make(chan Message),
	}

	go s.StartandListen()

	sendMessage(s.msgch, "hi everyone")
}

// func main() {
// 	// now := time.Now()
// 	userID := 10
// 	respch := make(chan string, 3)
// 	wg := &sync.WaitGroup{}
// 	wg.Add(3)
// 	go fetchUserData(userID, respch, wg)
// 	// wg.Add(1)
// 	go fetchUserRecommendations(userID, respch, wg)
// 	// wg.Add(1)
// 	go fetchUserLikes(userID, respch, wg)
// 	// wg.Add(1)

// 	wg.Wait()

// 	close(respch)
// 	for rsp := range respch {
// 		fmt.Println(rsp)
// 	}
// 	// fmt.Println(userData)
// 	// fmt.Println(userRecs)
// 	// fmt.Println(userLikes)
// 	// fmt.Println(time.Since(now))
// }

// func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {
// 	// time.Sleep(80 * time.Microsecond)

// 	respch <- "UserData"

// 	wg.Done()
// }

// func fetchUserRecommendations(userID int, respch chan string, wg *sync.WaitGroup) {
// 	// time.Sleep(120 * time.Microsecond)

// 	respch <- "user Recommendations"

// 	wg.Done()
// }

// func fetchUserLikes(userID int, respch chan string, wg *sync.WaitGroup) {
// 	// time.Sleep(50 * time.Microsecond)
// 	respch <- "User Likes"

// 	wg.Done()
// }
