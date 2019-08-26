package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type command struct {
	name string
	next *command
}

// func print(head *command) {
// 	tmp := head
// 	for tmp != nil {
// 		fmt.Print(tmp.name, " ")
// 		tmp = tmp.next
// 	}
// }

func addLink(head **command, new *command) {
	if *head == nil {
		*head = new
	} else {
		tmp := *head
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = new
	}
}

func scanner(stackA, stackB stack, functions map[string]func(A, B *Stack)) *command {
	var head *command
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if _, ok := functions[input]; !ok {
			fmt.Println("Error")
			os.Exit(1)
		}
		addLink(&head, &command{input, nil})
	}
	return head
}

func dataFromStack(stack *Stack) []float64 {
	data := []float64{}
	tmp := stack.top
	for tmp != nil {
		data = append(data, float64(tmp.num))
		tmp = tmp.below
	}
	return data
}

func visualizer(curent *command, A, B stack, functions map[string]func(A, B stack)) {

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to init termui: %v", err)
	}
	defer ui.Close()

	sA := widgets.NewSparkline()
	sA.Data = A
	sA.LineColor = ui.ColorRed
	sA.TitleStyle.Fg = ui.ColorWhite

	sB := widgets.NewSparkline()
	sB.Data = B
	sB.LineColor = ui.ColorBlue
	sB.TitleStyle.Fg = ui.ColorWhite

	sAg := widgets.NewSparklineGroup(sA)
	sAg.Title = "Stack A"
	sBg := widgets.NewSparklineGroup(sB)
	sBg.Title = "Stack B"

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(ui.NewRow(1.0/2, ui.NewCol(1.0, sAg)), ui.NewRow(1.0/2, ui.NewCol(1.0, sBg)))
	ui.Render(grid)
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Millisecond).C

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		case <-ticker:
			if curent != nil {
				functions[curent.name](A, B)
				curent = curent.next
				sAg.Sparklines[0].Data = A
				sBg.Sparklines[0].Data = B
				ui.Render(grid)
			}
		}
	}
}

func dupChecker() func(int) bool {
	duplicate := map[int]bool{}
	return func(num int) bool {
		if _, ok := duplicate[num]; !ok {
			duplicate[num] = true
			return false
		} else {
			return true
		}
	}
}

func main() {
	functions := map[string]func(A, B stack){
		"sa": sa, "sb": sb, "ss": ss, "pa": pa, "pb": pb, "ra": ra, "rb": rb, "rr": rr, "rra": rra, "rrb": rrb, "rrr": rrr}

	argv := os.Args[1:]
	stackA := stack{}
	stackB := stack{}
	isDup := dupChecker()
	size := len(argv)
	if size == 0 {
		os.Exit(0)
	}
	for _, str := range argv {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Bad input")
			os.Exit(0)
		}
		if isDup(num) {
			fmt.Println(num, "is duplicated, Error")
			os.Exit(1)
		}
		stackA = append(stackA, float64(num))
	}
	commandList := scanner(stackA, stackB, functions) //Make a linked list
	// print(commandList)
	visualizer(commandList, &stackA, &stackB, functions)
	// check(&stackA, &stackB)
}

/*
	Todo:
		Deal with negative numbers
			* get smallest number, and add that |x| to all the numbers
*/
