/*
Package lunacy provides an intuitive terminal-based interface for viewing LeftWM keybindings.

Lunacy lets you easily view your LeftWM keybindings defined in your standard config.ron file.
It reads your config.ron file and uses regular expressions to provide a readable
table for you to invoke.

Introduction:

LeftWM is a versatile window manager with the ability to customize keybindings.
As users transition from TOML to RON configurations, the need for a clear view
of keybindings becomes essential. Lunacy addresses this need, providing
a user-friendly way to quickly reference LeftWM keybinds without diving into config files.

Features:

 - Display Keybindings: View all your LeftWM keybindings in a neat table format.
 - Grouping: Keybindings are grouped by similar commands for ease of reference.
 - Custom Descriptions: The user has the capability to add custom descriptions for each keybinding, providing additional context.

Usage:

After building and installing the tool, simply run:

    lunacy

This will display a table of your LeftWM keybindings in the terminal.

Example:

    +--------------------------------+-------------------+-----------------------+
    | Command                        | Modifier          | Key                   |
    +--------------------------------+-------------------+-----------------------+
    | Execute rofi -show drun        | modkey            | space                 |
    ... (and so on for other keybinds)

Future Enhancements:

While the current implementation focuses on viewing keybindings, future versions may include
functionalities such as editing, searching, and exporting keybindings.

Note: This tool is specifically tailored for LeftWM users transitioning from TOML to RON configurations.
While it can be adapted for other use cases, certain functionalities might not directly translate.
*/
package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/rivo/tview"
)

type KeyBind struct {
	Command  string
	Value    string
	Modifier []string
	Key      string
}

var descriptions = map[string]string{
	"Execute":                       "Execute",
	"CloseWindow":                   "Close Window",
	"CloseAllOtherWindows":          "Close All Other Windows",
	"SoftReload":                    "Soft Reload",
	"MoveToLastWorkspace":           "Move to Last Workspace",
	"SwapTags":                      "Swap Tags",
	"HardReload":                    "Hard Reload",
	"ToggleFullScreen":              "Toggle Fullscreen",
	"ToggleMaximized":               "Toggle Maximize",
	"ToggleSticky":                  "Toggle Sticky Window",
	"GotoTag":                       "Go to tag",
	"ReturnToLastTag":               "Return To Last Tag",
	"FloatingToTile":                "Floating To Tile",
	"TileToFloating":                "Tile to Floating",
	"ToggleFloating":                "Toggle Floating",
	"MoveWindowUp":                  "Move Window Up",
	"MoveWindowDown":                "Move Window Down",
	"MoveWindowTop":                 "Move Window Top",
	"SwapWindowTop":                 "Swap Window Top",
	"FocusWindowUp":                 "Focus Window Up",
	"FocusWindowDown":               "Focus Window Down",
	"FocusWindowTop":                "Focus Window Top",
	"FocusWorkspaceNext":            "Focus Workspace Next",
	"FocusWorkspacePrevious":        "Focus Workspace Previous",
	"MoveWindowToNextTag":           "Move Window To Next Tag",
	"MoveWindowToPreviousTag":       "Move Window To Previous Tag",
	"MoveWindowToNextWorkspace":     "Move Window To Next Workspace",
	"MoveWindowToPreviousWorkspace": "Move Window To Next Workspace",
	"NextLayout":                    "Next Layout",
	"PreviousLayout":                "Previous Layout",
	"IncreaseMainSize":              "Increase Main Size",
	"DecreaseMainSize":              "Decrease Main Size",
	"IncreaseMainCount":             "Decrease Main Count",
	"DecreaseMainCount":             "Decrease Main Count",
	"UnloadTheme":                   "Unload Theme",
}

func getDescription(command, value string) string {
	if command == "MoveToTag" {
		return "Move to Tag " + value
	}

	if command == "GotoTag" {
		return "Go to Tag " + value
	}

	if command == "Execute" {
		return "Execute " + value
	}

	if description, ok := descriptions[command]; ok {
		return description
	}
	return command
}

func main() {
	keybinds := parseConfig()
	app := tview.NewApplication()

	table := tview.NewTable().
		SetFixed(1, 1).
		SetBorders(true)
	table.SetTitle("LeftWM Keybinds").SetBorder(true)

	// Column headers
	table.SetCell(0, 0, tview.NewTableCell("Command"))
	table.SetCell(0, 1, tview.NewTableCell("Modifier"))
	table.SetCell(0, 2, tview.NewTableCell("Key"))

	for i, kb := range keybinds {
		row := i + 1
		table.SetCell(row, 0, tview.NewTableCell(getDescription(kb.Command, kb.Value)).SetExpansion(2))
		table.SetCell(row, 1, tview.NewTableCell(strings.Join(kb.Modifier, ", ")).SetExpansion(1))
		table.SetCell(row, 2, tview.NewTableCell(kb.Key).SetExpansion(1))
	}

	app.SetRoot(table, true).SetFocus(table)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func parseConfig() (keybinds []KeyBind) {
	file, err := os.Open(os.Getenv("HOME") + "/.config/leftwm/config.ron")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`\(command: (\w+), value: "(.*?)", modifier: \[([^\]]*)\], key: "(.*?)"\)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		if len(matches) == 5 {
			modifierStrings := strings.Split(matches[3], ",")
			for i, mod := range modifierStrings {
				modifierStrings[i] = strings.TrimSpace(strings.ReplaceAll(mod, "\"", ""))
			}
			keybinds = append(keybinds, KeyBind{
				Command:  matches[1],
				Value:    matches[2],
				Modifier: modifierStrings,
				Key:      matches[4],
			})
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return keybinds
}
