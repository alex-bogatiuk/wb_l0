package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/nats-io/nats.go"
)

func runExecutableFile(path string, args ...string) {
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing %s: %v\n", path, err)
	}
}
func main() {

	//run nats-server in goroutine
	go runExecutableFile("./nats-server.exe")

	nc, _ := nats.Connect(nats.DefaultURL)

	defer nc.Drain()

	nc.Publish("greet.joe", []byte("hello"))

	sub, _ := nc.SubscribeSync("greet.*")

	msg, _ := sub.NextMsg(10 * time.Millisecond)
	fmt.Println("subscribed after a publish...")
	fmt.Printf("msg is nil? %v\n", msg == nil)

	nc.Publish("greet.joe", []byte("hello"))
	nc.Publish("greet.pam", []byte("hello"))

	msg, _ = sub.NextMsg(10 * time.Millisecond)
	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)

	msg, _ = sub.NextMsg(10 * time.Millisecond)
	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)

	nc.Publish("greet.bob", []byte("hello"))

	msg, _ = sub.NextMsg(10 * time.Millisecond)
	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)
}
