package main

import (
	"fmt"
	"log"
)

// the number of attendees we need to serve lunch to
const consumerCount = 10

// foodCourses represents the types of resources to pass to the consumers
var foodCourses = []string{
	"Caprese Salad",
	"Spaghetti Carbonara",
	"Vanilla Panna Cotta",
}

// takeLunch is the consumer function for the lunch simulation
// Change the signature of this function as required
func takeLunch(name string, in []chan string, done chan<- struct{}) {
	for _, ch := range in {
		log.Printf("%s eats %s\n", name, <-ch)
	}
	done <- struct{}{}
}

// serveLunch is the producer function for the lunch simulation.
// Change the signature of this function as required
func serveLunch(course string, out chan<- string, done <-chan struct{}) {
	for {
		select {
		case out <- course:
		case <-done:
			return
		}
	}
}

func main() {
	log.Printf("Welcome to the conference lunch! Serving %d attendees.\n",
		consumerCount)
	doneServing := make(chan struct{})
	doneEating := make(chan struct{})
	var courses []chan string

	for _, course := range foodCourses {
		ch := make(chan string)
		courses = append(courses, ch)
		go serveLunch(course, ch, doneServing)
	}
	for i := 0; i < consumerCount; i++ {
		consumer := fmt.Sprintf("Consumer %d", i)
		go takeLunch(consumer, courses, doneEating)
	}
	for i := 0; i < consumerCount; i++ {
		<-doneEating
	}
	close(doneServing)
}
