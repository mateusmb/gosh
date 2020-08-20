package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("-----------------------------")
	fmt.Println("------ Welcome to gosh ------")
	fmt.Println("-----------------------------")

	user, err := user.Current()
	if err != nil {
		log.Println(err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Println(err)
	}

	var prompt string

	for {
		prompt = SetPrompt(user.Username, hostname)
		fmt.Print(prompt)
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		command = strings.Replace(command, "\n", "", -1)

		if strings.Compare("exit", command) == 0 {
			break
		} else if strings.HasPrefix(command, "cd") {
			ChangeDirectory(command, user.HomeDir)
		} else if strings.HasPrefix(command, "ls") {
			fmt.Println(ListDirectory(command))
		} else if strings.Compare("pwd", command) == 0 {
			fmt.Println(getCwd())
		} else {
			fmt.Println("Command doesn't exist or is not implemented yet! Sorry :(")
		}

	}

}

func SetPrompt(username string, hostname string) string {
	var path string
	home := "/home/" + username

	cwd := getCwd()

	if strings.HasPrefix(cwd, home) {
		path = "~" + cwd[len(home):]
	} else {
		path = cwd
	}
	prompt := username + "@" + hostname + ":" + path + "$ "
	return prompt
}

func getCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return cwd
}

func ChangeDirectory(command string, userHome string) {
	cdCommand := strings.Split(command, " ")

	if len(cdCommand) == 1 {
		os.Chdir(userHome)
	} else if len(cdCommand) == 2 {
		os.Chdir(cdCommand[1])
	} else {
		fmt.Println("Error! Too much arguments for cd")
	}
}

func ListDirectory(command string) string {
	lsCommand := strings.Split(command, " ")
	app := "/bin/ls"

	if len(lsCommand) == 1 {
		out, err := exec.Command(app).Output()
		if err != nil {
			log.Println(err)
		}
		return string(out[:])
	} else {
		args := lsCommand[1:]
		out, err := exec.Command(app, args...).Output()
		if err != nil {
			log.Println(err)
		}
		return string(out[:])
	}
}
