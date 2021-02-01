package main

import "fmt"

func main() {

	var userEmails map[int]string = make(map[int]string)

	userEmails[1] = "user1@gmail.com"
	userEmails[2] = "user2@gmail.com"

	fmt.Println(userEmails)

	usrMails := map[int]string{
		3: "user3@gmail.com",
		4: "user4@gmail.com",
	}

	fmt.Println(usrMails)

	usrMails[1] = "new1@gmail.com"

	fmt.Println(usrMails)

	firstEmail := usrMails[1]
	fmt.Println(firstEmail)

	fifthEmail, ok := usrMails[5]
	fmt.Println(fifthEmail, ok)

	if _, ok := usrMails[5]; ok {
		fmt.Println("email exisits")
	} else {
		fmt.Println("email doesn't exist")
	}

	delete(usrMails, 3)
	fmt.Println(usrMails)
}
