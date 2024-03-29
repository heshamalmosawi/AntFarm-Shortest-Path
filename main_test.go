package main

import (
	"lemin/lemin"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestProcessData(t *testing.T) {
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

func TestRemovefrompath(t *testing.T) {
	path := []string{"0", "2", "3", "1"}
	paths := [][]string{path}
	returned := lemin.RemoveFromPath(paths, "2")
	if len(returned) != 0 {
		t.Errorf("Removefrompath return value expected 0 but got %v", returned)
	}
	paths = [][]string{
		{"A", "B", "C", "D"},
		{"X", "Y", "Z"},
		{"M", "N", "O", "P", "Q"},
	}

	// Remove room "B" from all paths
	returned = lemin.RemoveFromPath(paths, "B")

	expected := [][]string{
		{"X", "Y", "Z"},
		{"M", "N", "O", "P", "Q"},
	}

	// Check if the returned value matches the expected value
	if !reflect.DeepEqual(returned, expected) {
		t.Errorf("Removefrompath return value expected %v but got %v", expected, returned)
	}

}

func TestGenerateCombinations(t *testing.T) {
	paths := [][]string{
		{"A", "B", "C", "D"},
		{"X", "A", "Z"},
		{"M", "N", "O", "P", "Q"},
	}
	var result = [][][]string{}
	lemin.GenerateCombinations(paths, &result)
	expected := [][][]string{
		{{"A", "B", "C", "D"}},
		{{"A", "B", "C", "D"}, {"M", "N", "O", "P", "Q"}},
		{{"M", "N", "O", "P", "Q"}},
	}
	if reflect.TypeOf(result) != reflect.TypeOf(expected) {
		t.Errorf("Generate combinations return type expected [][][]string but got %v", reflect.TypeOf(result))
	}
	if reflect.DeepEqual(result, expected) {
		t.Errorf("Generated combination return value expected %v but got %v", expected, result)
	}
}
