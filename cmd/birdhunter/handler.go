package main

import (
	"fmt"
	"os"
)

func handler() {
	if len(os.Args) < 2 {
		printDefaults()
	}

	switch os.Args[1] {
	case "popular":
		fmt.Println("This may take some while, popular feed...")
		popularReq()

		popularCommand.Parse(os.Args[2:])

		popular()
	case "tag":
		fmt.Println("This may take some while...")
		tagReq()

		tagCommand.Parse(os.Args[2:])

		tag()
	case "update":
		fmt.Println("Updating the configuration file...")
		updateReq()

		update()
	case "forever":
		foreverReq()

		foreverCommand.Parse(os.Args[2:])

		forever()
	default:
		printDefaults()
	}
}

func popular() {
	NewClient(
		usernamePtr,
		passwordPtr,
		"",
	)
	NewLikes(
		likesPtr,
		minlikesPtr,
	)

	login()

	getPopularIds()
}

func tag() {
	NewClient(
		usernamePtr,
		passwordPtr,
		"",
	)
	NewLikes(
		likesPtr,
		minlikesPtr,
	)

	login()

	getTagIds()
}

func update() {
	updateConfig(
		fmt.Sprintf(
			"%v/update",
			httpAddrPtr,
		),
	)
}

func forever() {
	start(httpAddrPtr)
}
