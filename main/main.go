package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/micmonay/keybd_event"
	"io"
	"log"
	"os"
	"runtime"
	"time"
)

const timeSleep = 30
const spamFilePath = "../spam_message/message.txt"

var text = ""

func main() {
	err := getSpam()
	if err != nil {
		log.Print(err)
		time.Sleep(5 * time.Second)
		log.Fatal(err)
	}

	kb, err := initKeyboard()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	for true {
		saveToBuffer()
		enter(kb)
		time.Sleep(time.Microsecond)
		ctrlV(kb)
		time.Sleep(time.Microsecond)
		enter(kb)
		time.Sleep(timeSleep * time.Second)
	}
}

func getSpam() error {
	file, err := os.Open(spamFilePath)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	text = string(b)
	return nil
}

func saveToBuffer() {
	err := walk.Clipboard().Clear()
	err = walk.Clipboard().SetText(text)
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func initKeyboard() (keybd_event.KeyBonding, error) {
	kb, err := keybd_event.NewKeyBonding()
	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	return kb, err
}

func enter(kb keybd_event.KeyBonding) {
	// Select keys to be pressed
	kb.SetKeys(keybd_event.VK_ENTER)
	err := kb.Launching()
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func ctrlV(kb keybd_event.KeyBonding) {
	// Select keys to be pressed
	kb.SetKeys(keybd_event.VK_V)

	kb.HasCTRL(true)
	err := kb.Launching()
	if err != nil {
		fmt.Printf("%v", err)
	}
}
