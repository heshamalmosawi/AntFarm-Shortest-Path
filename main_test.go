package main

import (
	"lemin/lemin"
	"testing"
)

func TestMain(m *testing.M){
	m.Run()
}

func TestProcessData(t *testing.T){
	input := []string{
		"10",
		"##start",
		"start_room 1 2",
		"room1 3 4",
		"##end",
		"end_room 5 6",
		"start_room-room1",
		"room1-end_room",
	}
	lemin.ProcessData(input)
}

func TestValidCoord(t *testing.T) {
    // Prepare test data with duplicate coordinates
    graph := &lemin.Graph{}
	err := graph.Addvertex("room1", "1", "2")
	err1 := graph.Addvertex("room2", "2", "2")
	err2 := graph.Addvertex("room3", "2", "2")
	if err != nil || err1 != nil || err2 != nil {
		t.Error("Error adding vertex")
	}
    // Call the method being tested
    err = graph.ValidCoord()

    // Check if error is returned as expected
    if err == nil {
        t.Error("Expected error but got nil")
    } else {
        expectedErrorMessage := "ERROR: Invalid data. Two or more vertices have matching coordinates \n[room2 and room3]"
        if err.Error() != expectedErrorMessage {
            t.Errorf("Expected error message '%s' but got '%s'", expectedErrorMessage, err.Error())
        }
    }
}
