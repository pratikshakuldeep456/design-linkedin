package main

import (
	"fmt"
	"pratikshakuldeep456/design-linkedin/pkg/linkedin"
)

func main() {

	u1 := linkedin.CreateUser(1, "Pratiksha", "abc@gmail.com", "password@123")

	u2 := linkedin.CreateUser(2, "Pratiksha", "abc@gmail.com", "password@123")
	fmt.Println(*u2)

	u1.UpdateInfo("Software Engineer", "Experienced Go Developer", &linkedin.Experience{Company: "Google"}, nil, nil)

	fmt.Println(*u1.Profile)
	u2.UpdateInfo("Another headline", "Another summary", &linkedin.Experience{}, &linkedin.Education{}, &linkedin.Skills{})

	fmt.Println(*u2)

	err := u1.SendRequest(u2)
	if err != nil {
		fmt.Println("SendRequest Error:", err)
	}
	fmt.Println("Bob's Pending Requests:", u2.PendingRequests)
	fmt.Println("Alice's Pending Requests:", u1.PendingRequests)

	//u2.DeclineRequest(1)
	u2.AcceptRequest(1)

	fmt.Println("see list", u2.FetchconnectionList())

	fmt.Println("Bob's Pending Requests:", u2.PendingRequests)
}
