package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
	"io/ioutil"
  	"encoding/json"
	"strings"
	"github.com/gorilla/mux"
)

var pid = -1
type TCPDumpRequest struct {
	Action   string `json:"action"`
	Pcapfile string `json:"pcap_file"`
}

type TCPDumpResponse struct {
	Status   string `json:"status"`
	Pcapfile string `json:"pcap_file"`
}

func createDir(filePath string) error {
	parts := strings.Split(filePath, "/")
	dir := ""
	for s := range parts {
		if s < (len(parts) - 1) {
			fmt.Println("Out:", s, parts[s])
			dir += parts[s] + "/"
		}
	}
	return os.MkdirAll(dir, os.ModePerm)
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "This is TCPDump Server "
	}
	log.Printf("Received request for %s\n", name)
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}

func tcpdumpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[DEBUG] tcpdumpHandler++")
	var msg TCPDumpRequest
	var resp TCPDumpResponse
	resp.Status = "Unknown"

	switch r.Method {

	case "POST":
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Unmarshal		
		err = json.Unmarshal(b, &msg)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		
		log.Printf("Request received %+v \n", msg)
		
		if msg.Action == "start" {
			if pid == -1 {
	                	filePath := os.Getenv("PCAP_DIR")+"/"+ msg.Pcapfile 
				err := createDir(filePath)
				if err != nil {
					resp.Status = "Failed to create dir"
					fmt.Printf("[ERROR]Failed to create dir %s", err)
				}else{
					fmt.Println("[DEBUG] tcpdump start")
					cmd := exec.Command("tcpdump","-i","any","-w", "/dump/trace_%Y-%m-%d_%H%M%S"+filePath+".pcap", "-G", "3600", "-z", "/var/tmp/cleanup.sh")
					err := cmd.Start()
					if err != nil {
						fmt.Println("[ERROR] tcpdump start",err)
					}
					pid = cmd.Process.Pid
					fmt.Println("Started tcpdump pid",cmd.Process.Pid)  
					resp.Status = "Started"
					resp.Pcapfile = msg.Pcapfile
				}
				
			}else{
				resp.Status = "Already running"
				fmt.Println("Already running tcpdump pid",pid)
			}
			
		}else if msg.Action == "stop" {
			fmt.Println("[DEBUG] tcpdump stop ",pid)
			
			process,err := os.FindProcess(pid)
			if err == nil {
				if err = process.Kill(); err == nil {
					resp.Status = "Stopped"
				}else{
					resp.Status = "Failed to stop"					
					fmt.Println("[ERROR] failed to kill process  ",pid)
				}
			}else{
				resp.Status = "Process not found"
				fmt.Println("[ERROR] process not found ",pid)
			}
			
			pid = -1
			resp.Pcapfile = msg.Pcapfile
			
		}
		output, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
		return
		
	default:
		log.Printf("Method not supported %s\n", r.Method)
		http.Error(w, fmt.Sprintf("%sMethod not supported\n",r.Method), 404)
		
		return
	}

}

func main() {
	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", handler)
	r.HandleFunc("/tcpdump", tcpdumpHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}

