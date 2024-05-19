package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	// Establish Socket, monitor port first step: binding port
	//netListen, err := net.Listen("tcp", "localhost:1024")
	netListen, err := net.Listen("tcp", "localhost:9800")
	CheckError(err)
	// defer delay closes the resource to avoid memory leakage
	defer netListen.Close()
	Log("Waiting for clients")
	//cmd, err := exec.Command("C:\\Program Files\\Java\\jdk-19\\bin\\java.exe", "-Djava.library.path=D:\\Apache\\apache-jena-4.6.1\\lib", "-Xmx2G", "-jar", "D:\\Apache\\apache_idea\\example\\out\\artifacts\\example_jar\\example.jar").Output()
	/*
		cmd_output, err := exec.Command("bash", "-jar", "D:\\Apache\\apache_idea\\example\\out\\artifacts\\example_jar\\example.jar").Output()
	*/
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue // Error exits the current cycle
		}
		Log(conn.RemoteAddr().String(), " tcp connect success")
		//conn.Write([]byte("200"))
		// Handleconnection (conn) // normal connection
		// The front of this code plus a Go, you can make the server concurrently process the request from different clients.
		go handleConnection(conn) // use goroutine to handle user requests
	}
}

// Processing connection
func handleConnection(conn net.Conn) {
	fmt.Println("before defer")
	// Close connection
	defer conn.Close()
	fmt.Println("after defer")
	javaPath := "C:\\Program Files\\Java\\jdk-19\\bin\\java"
	jarPath := "D:\\Apache\\apache_idea\\example\\out\\artifacts\\example_jar\\example.jar"
	data, err := ioutil.ReadFile("D:\\Apache\\apache_idea\\example\\src\\Jena03\\resources\\example_road\\dataset.n3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	cmd := exec.Command(javaPath, "-jar", jarPath)
	fmt.Println("before Start")
	err = cmd.Start()
	fmt.Println("after Start")
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cmd)
}

// log output
func Log(v ...interface{}) {
	log.Println(v...)
}

// Handling Error
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
