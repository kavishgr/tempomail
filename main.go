package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"flag"
)

var cyan = color.New(color.FgCyan)
var cyanBold = cyan.Add(color.Bold)

var m = make(map[int]int) // IDs as key

var path string

func setPath(){
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}

func incrementMap(response []CheckMailTemplate) {
	for _, v := range response {
		if _, found := m[v.Id]; !found {
			m[v.Id] = 0
		}
	}

}

type CheckMailTemplate struct {
	Id          int    `json:"id"`
	From        string `json:"from"`
	Subject     string `json:"subject"`
	Date        string `json:"date"`
	Attachments []struct {
		Filename    string `json:"filename"`
		ContentType string `json:"contentType"`
		Size        int    `json:"size"`
	}
	Body     string `json:"body"`
	Textbody string `json:"textBody"`
	// Htmlbody string `json:"htmlbody,omitempty"`
}

var response []CheckMailTemplate // getMessage

func checkMail(name string, domainOnly string) int {
	s := fmt.Sprintf("https://www.1secmail.com/api/v1/?action=getMessages&login=%v&domain=%v", name, domainOnly)
	resp, _ := http.Get(s)
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		fmt.Println(err)
	}
	resp.Body.Close()

	numOfEmails := len(response)

	switch numOfEmails {
	case 0:
		fmt.Printf("\r%s", "Mailbox is empty")
	case 1:
		fmt.Printf("\r%s", "You received 1 mail in your mailbox")
	default:
		e := fmt.Sprintf("You received %v mails in your mailbox", numOfEmails)
		fmt.Printf("\r%s", e)
	}

	return numOfEmails
}

func saveMail(name, domainOnly string) { //use this function to save emails

	for k, _ := range m {
		if m[k] == 0 {
			s := fmt.Sprintf("https://www.1secmail.com/api/v1/?action=readMessage&login=%v&domain=%v&id=%v", name, domainOnly, k)
			resp, _ := http.Get(s)
			body, _ := ioutil.ReadAll(resp.Body)
			time.Sleep(1 * time.Second)
			resp.Body.Close()
			var recvmail CheckMailTemplate
			err := json.Unmarshal([]byte(body), &recvmail)
			// attachment := recvmail.Attachments[0].Filename
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				receivedMail := fmt.Sprintf("ID: %v\nFrom: %v\nSubject: %v\nDate: %v\nText: %v\n", recvmail.Id, recvmail.From, recvmail.Subject, recvmail.Date, recvmail.Textbody)
				filename := fmt.Sprintf("%v%v", path, recvmail.Id)
				err := ioutil.WriteFile(filename, []byte(receivedMail), 0666)
				if err != nil {
					log.Fatal(err)
				}

				m[k] = 1
			}
		}
	}

}

func deleteMail(name, domainOnly string) error {
	delUrl := "https://www.1secmail.com/mailbox"
	data := url.Values{}
	data.Set("action", "deleteMailbox")
	data.Set("login", name)
	data.Set("domain", domainOnly)
	_, err := http.PostForm(delUrl, data)
	return err

}

func handleInterrupt(name, domainOnly string) {
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		deleteMail(name, domainOnly)
		err := os.RemoveAll(path)
		if err != nil {
			log.Fatal(err)
		}
		clear()
		fmt.Println("Emails Deleted. Exiting.")
		os.Exit(0)
	}()
}

func verifyName() string {
	bannedEmail := []string{"abuse", "webmaster", "contact", "postmaster", "hostmaster", "admin"}
	var name string
	fmt.Print("Input name: ")
	fmt.Scan(&name)
	name = strings.ToLower(name)
	for _, v := range bannedEmail {
		if name == v {
			fmt.Printf("You cannot read messages from these addresses: \n%v", bannedEmail)
			fmt.Println()
			fmt.Println("All other addresses are free to use.")
			os.Exit(1)
		}
	}
	return name
}

func generateEmail(name string) (string, string) {
	domain := "1secmail"
	rootDomain := []string{".com", ".net", ".org"}
	randomIndex := rand.Intn(len(rootDomain))
	value := rootDomain[randomIndex]
	domainOnly := domain + value
	email := name + "@" + domain + value
	return email, domainOnly
}

func createEmail(name string, domainOnly string) {
	info := fmt.Sprintf("https://www.1secmail.com/?login=%v&domain=%v", name, domainOnly)
	resp, err := http.Get(info)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Status Code: %v. Email not created.", resp.StatusCode)
		os.Exit(1)
	}

}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func main() {
	flag.StringVar(&path, "path", "/tmp/1secmails/", "specify directory to store emails")
	flag.Parse()
	setPath()
	name := verifyName()
	email, domainOnly := generateEmail(name)
	createEmail(name, domainOnly)
	clear()
	printName := cyanBold.Sprintf(email)
	fmt.Printf("Your Temporary Email: %v\n", printName)
	fmt.Println("Mailbox content is refreshed automatically every 5 seconds.")
	fmt.Printf("All Emails are saved in %v\n", path)

	for {
		checkMail(name, domainOnly)
		go handleInterrupt(name, domainOnly)
		saveMail(name, domainOnly)
		// handleInterrupt(name, domainOnly) // after receiving mails(not needed)
		time.Sleep(5 * time.Second)
	}

}
