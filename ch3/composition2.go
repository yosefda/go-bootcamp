package main

import (
	"log"
	"os"
)

// Job Job data structure
type Job struct {
	Command string
	Logger  *log.Logger
}

// Job2 In this data structure we directly compose
// Logger struct. This way we directly use all its
// methods
type Job2 struct {
	Command     string
	*log.Logger // composing Logger struct
}

func main() {
	job := &Job{
		Command: "demo job1",
		Logger:  log.New(os.Stdout, "Job1: ", log.Ldate),
	}

	job.Logger.Print("test job1")

	job2 := &Job2{
		Command: "demo job2",
		Logger:  log.New(os.Stdout, "Job2: ", log.Ldate),
	}

	job2.Print("test job2")
}
