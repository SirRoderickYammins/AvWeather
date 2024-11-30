package core

import (
	"encoding/json"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Default DB path
var dbPath = "../db/pilots.json"

type PilotProfilePage struct {
	Grid *tview.Grid
	Node *tview.TreeNode
}

type Pilot struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DisplayName string `json:"displayName"`
	HomeAirport string `json:"homeAirport"`
}

// Checks if db Exists, if not will create one.
// Otherwise, returns array of pilots in db
func loadPilotDB() ([]Pilot, error) {
	// Check if the database exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		// If file doesn't exist, create it.
		err := os.WriteFile(dbPath, []byte("[]"), 0o644)
		if err != nil {
			return nil, err
		}
	}
	// Open the file
	file, err := os.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pilots []Pilot

	err = json.NewDecoder(file).Decode(&pilots)
	if err != nil {
		return nil, err
	}
	return pilots, nil
}

func addNewPilot(pilot *Pilot) error {
	data, err := json.MarshalIndent(pilot, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(dbPath, data, 0o644)
}

func GenPilot(firstName string, lastName string, homeApt string) (pilot *Pilot) {
	displayName := firstName + lastName

	pilot = &Pilot{
		FirstName:   firstName,
		LastName:    lastName,
		HomeAirport: homeApt,
		DisplayName: displayName,
	}

	loadPilotDB()

	return pilot
}

// Generate the pilot form itself
func CreatePilotForm() *tview.Form {
	firstName := ""
	lastName := ""
	homeAirport := ""

	getFormItems := func() {
		GenPilot(firstName, lastName, homeAirport)
	}
	form := tview.NewForm().
		AddInputField("First Name", "", 20, nil, func(text string) {
			firstName = text
		}).
		AddInputField("Last Name", "", 10, nil, func(text string) {
			lastName = text
		}).
		AddInputField("Home Airport", "", 5, nil, func(text string) {
			homeAirport = text
		}).
		AddButton("Create Pilot", getFormItems)

	return form
}

// Create the pilot profiles page
func CreatePilotProfilePage() *PilotProfilePage {
	newpilotForm := CreatePilotForm()

	pilotProfilesTree := tview.NewTreeNode("Pilots").
		SetColor(tcell.ColorGreen)

	tree := tview.NewTreeView().
		SetRoot(pilotProfilesTree).
		SetCurrentNode(pilotProfilesTree)

	grid := tview.NewGrid().
		SetColumns(0, 0).
		SetRows(0, 1)

	controls := tview.NewTextView().
		SetText(`Press Q to return to Main Menu`).
		SetTextAlign(tview.AlignCenter)

	grid.AddItem(tree, 0, 0, 1, 1, 0, 100, true).
		AddItem(newpilotForm, 0, 1, 1, 1, 0, 100, false).
		AddItem(controls, 1, 0, 1, 2, 1, 5, false).
		SetBorders(true)

	return &PilotProfilePage{
		Grid: grid,
		Node: pilotProfilesTree,
	}
}
