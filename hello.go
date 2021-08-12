package main

import "fmt"
import "log"

type MessageResponse struct {
	Message string `json:"message"`
}

// JobResponse models the respose returned by the /job endpoint.
type JobResponse struct {
	MessageResponse
	JobID string `json:"job_id"`
}

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)

    log.Println("start")
    log.Printf("var: %v", 42)

    fmt.Println("Hello, World!")

    resp := JobResponse{JobID: "42"}
    log.Printf("%v", resp.JobID)
    log.Printf("%v", resp.Message)

    log.Println("end")
}
