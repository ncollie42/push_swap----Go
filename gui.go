package main

import (
	"fmt"
	"log"

	c "github.com/jroimartin/gocui"
	"github.com/pkg/errors"
)

const (
	lw = 20
	ih = 3
)

var listItems = []string{
	"Line 1",
	"Line 2",
	"Line 3",
	"Line 4",
	"Line 5",
}

func runGocui() {
	// Create a new GUI.
	g, err := c.NewGui(c.OutputNormal)
	if err != nil {
		log.Println("Failed to create a GUI:", err)
		return
	}
	defer g.Close()

	// Activate the cursor for the current view.
	g.Cursor = true

	// The GUI object wants to know how to manage the layout.
	// Unlike `termui`, `gocui` does not use
	// a grid layout. Instead, it relies on a custom layout handler function
	// to manage the layout.

	// Here we set the layout manager to a function named `layout`
	// that is defined further down.
	g.SetManagerFunc(layout)

	// Bind the `quit` handler function (also defined further down) to Ctrl-C,
	// so that we can leave the application at any time.
	err = g.SetKeybinding("", c.KeyCtrlC, c.ModNone, quit)
	if err != nil {
		log.Println("Could not set key binding:", err)
		return
	}

	// Now let's define the views.

	// The terminal's width and height are needed for layout calculations.
	tw, th := g.Size()

	// First, create the list view.
	lv, err := g.SetView("list", 0, 0, lw, th-1)
	// ErrUnknownView is not a real error condition.
	// It just says that the view did not exist before and needs initialization.
	if err != nil && err != c.ErrUnknownView {
		log.Println("Failed to create main view:", err)
		return
	}
	lv.Title = "List"
	lv.FgColor = c.ColorCyan

	// Then the output view.
	ov, err := g.SetView("output", lw+1, 0, tw-1, th-ih-1)
	if err != nil && err != c.ErrUnknownView {
		log.Println("Failed to create output view:", err)
		return
	}
	ov.Title = "Output"
	ov.FgColor = c.ColorGreen
	// Let the view scroll if the output exceeds the visible area.
	ov.Autoscroll = true
	_, err = fmt.Fprintln(ov, "Press Ctrl-c to quit")
	if err != nil {
		log.Println("Failed to print into output view:", err)
	}

	// And finally the input view.
	iv, err := g.SetView("input", lw+1, th-ih, tw-1, th-1)
	if err != nil && err != c.ErrUnknownView {
		log.Println("Failed to create input view:", err)
		return
	}
	iv.Title = "Input"
	iv.FgColor = c.ColorYellow
	// The input view shall be editable.
	iv.Editable = true
	err = iv.SetCursor(0, 0)
	if err != nil {
		log.Println("Failed to set cursor:", err)
		return
	}

	// Make the enter key copy the input to the output.
	err = g.SetKeybinding("input", c.KeyEnter, c.ModNone, func(g *c.Gui, iv *c.View) error {
		// We want to read the view's buffer from the beginning.
		iv.Rewind()

		// Get the output view via its name.
		ov, e := g.View("output")
		if e != nil {
			log.Println("Cannot get output view:", e)
			return e
		}
		// Thanks to views being an io.Writer, we can simply Fprint to a view.
		_, e = fmt.Fprint(ov, iv.Buffer())
		if e != nil {
			log.Println("Cannot print to output view:", e)
		}
		// Clear the input view
		iv.Clear()
		// Put the cursor back to the start.
		e = iv.SetCursor(0, 0)
		if e != nil {
			log.Println("Failed to set cursor:", e)
		}
		return e

	})
	if err != nil {
		log.Println("Cannot bind the enter key:", err)
	}

	// Fill the list view.
	for _, s := range listItems {
		// Again, we can simply Fprint to a view.
		_, err = fmt.Fprintln(lv, s)
		if err != nil {
			log.Println("Error writing to the list view:", err)
			return
		}
	}

	// Set the focus to the input view.
	_, err = g.SetCurrentView("input")
	if err != nil {
		log.Println("Cannot set focus to input view:", err)
	}

	// Start the main loop.
	err = g.MainLoop()
	log.Println("Main loop has finished:", err)
}

// The layout handler calculates all sizes depending
// on the current terminal size.
func layout(g *c.Gui) error {
	// Get the current terminal size.
	tw, th := g.Size()

	// Update the views according to the new terminal size.
	_, err := g.SetView("list", 0, 0, lw, th-1)
	if err != nil {
		return errors.Wrap(err, "Cannot update list view")
	}
	_, err = g.SetView("output", lw+1, 0, tw-1, th-ih-1)
	if err != nil {
		return errors.Wrap(err, "Cannot update output view")
	}
	_, err = g.SetView("input", lw+1, th-ih, tw-1, th-1)
	if err != nil {
		return errors.Wrap(err, "Cannot update input view.")
	}
	return nil
}

// `quit` is a handler that gets bound to Ctrl-C.
// It signals the main loop to exit.
func quit(g *c.Gui, v *c.View) error {
	return c.ErrQuit
}
