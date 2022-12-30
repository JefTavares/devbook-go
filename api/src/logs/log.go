package logs

import (
	"log"
	"os"
)

func Write(message, filepath string) {
	if filepath == "" {
		filepath = "./tmp/my-app.log"
	}

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	log.SetOutput(file)

	// loga data e hora
	log.SetFlags(log.LstdFlags | log.Ldate | log.Ltime | log.Lshortfile)

	log.Println(message)
}
