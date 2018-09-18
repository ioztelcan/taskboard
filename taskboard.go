package main

import (
	ui "github.com/gizak/termui"
	//"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

type Task struct {
	id int
	status string
	description string
	tags []string
	project string
	uuid string
}

var g_tasks []Task

var tasks []string
var rows [][]string

func reset_rows() {
	rows = nil
	rows = append(rows, []string{"Tasks"})
}

func create_empty_task() Task {
	return Task {
		id: 0,
		status: "new",
		description: "empty",
		project: "dev",
		uuid: "0"}
}

func update_tasks() {
	file, err := os.Open("/home/ioztelcan/.task/pending.data")
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()

	var lines []string
	task_counter := 1

	g_tasks = nil // This extra bad way of avoiding duplicates list will change.

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, val := range lines {
		new_task := create_empty_task()
		new_task.id = task_counter
		split_line := strings.Split(val, "\"")

		for i, e := range split_line {
			switch stripped := strings.Trim(e, "[]: "); stripped {
			case "description":
				new_task.description = split_line[i + 1]

			case "status":
				new_task.status = split_line[i + 1]

			case "tags":
				t := split_line[i + 1]
				tags := strings.Split(t, ",")
				for _, tag := range tags {
					new_task.tags = append(new_task.tags, tag)
				}
			case "project":
				new_task.project = split_line[i + 1]

			case "uuid":
				new_task.uuid = split_line[i + 1]
			}
		}
		g_tasks = append(g_tasks, new_task)
		task_counter++
	}

    if scanner.Err() != nil {
        log.Fatal(err)
	}
}

func create_items_list(project string)(items []string) {
	for _, val := range g_tasks {
		if val.project == project && val.status == "pending" {
			task := strings.Join([]string{strconv.Itoa(val.id), ": ", val.description}, "")
			items = append(items, task)
		}
	}
	return items
}

func main() {

	reset_rows()
	update_tasks()

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	// List

	column_left := ui.NewList()
	column_mid := ui.NewList()
	column_right := ui.NewList()


	column_left.Items = create_items_list("dev")
	column_left.BorderLabel = "Development"
	column_left.Height = 50

	column_mid.Items = create_items_list("Test_Framework")
	column_mid.BorderLabel = "Test_Framework"
	column_mid.Height = 50

	column_right.Items = create_items_list("Other")
	column_right.BorderLabel = "Other"
	column_right.Height = 50

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(4, 0, column_left),
			ui.NewCol(4, 0, column_mid),
			ui.NewCol(4, 0, column_right)))

	ui.Body.Align()

	ui.Render(column_left, column_mid, column_right)

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})

	// Tasks are updated every 1 second, this can be much better if it is based file change.
	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)
		// t is a EvtTimer
		if t.Count%2 ==0 {
			update_tasks()
			column_left.Items = create_items_list("dev")
			column_mid.Items = create_items_list("Test_Framework")
			column_right.Items = create_items_list("Other")
			ui.Render(column_left, column_mid, column_right)
		}
	})

	ui.Loop()
}
