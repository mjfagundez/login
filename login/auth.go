package login

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type loginData struct {
	nickName   string
	password   string
	authNumber int
}

var (
	user1 = loginData{"Xox", "123qwe", 12}
	user2 = loginData{"Xox2", "114234qwer", 15}
	user3 = loginData{"Lisa", "p", 12}
)

var users = []loginData{user1, user2, user3}

func checkLogin(l loginData, u string, p string) bool {
	//fmt.Printf("User name: '%s'\n", u)
	//fmt.Printf("Password: '%s'\n", p)

	if l.nickName == u && l.password == p {
		return true
	}
	return false
}
func checkAuthNumber(l loginData, a int) bool {
	if l.authNumber == a {
		return true
	}
	return false
}

func LoginInterface() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	fmt.Print("Enter authentication number: ")
	var authNumber int
	fmt.Scanf("%d", &authNumber)

	username = strings.TrimSuffix(username, "\n")
	password = strings.TrimSuffix(password, "\n")

	found := false
	authValid := false
	for _, user := range users {
		if checkLogin(user, username, password) {
			found = true
			if checkAuthNumber(user, authNumber) {
				authValid = true
				break
			}
		}
	}
	if found && authValid {
		fmt.Println("Welcome")
	} else if found && !authValid {
		fmt.Println("Invalid authentication number")
	} else {
		fmt.Println("Invalid Username or Password")
	}

}
