package database

import (
	"fmt"

	"github.com/aybabtme/rgbterm"
)

func Display(arr []Data) {
	fmt.Println("ID, Memo, Finished, Created_at, Due_at, Priority, Tag")
	for _, data := range arr {
		fmt.Println(data.Id, data.Memo, data.Finished, data.Created_at, data.Due_at, priorityColor(data.Priority), data.Tag)
	}
}

func priorityColor(priority string) string {
	if priority == "Red" {
		return rgbterm.String(priority, 255, 10, 10, 0, 0, 0)
	} else if priority == "Yellow" {
		return rgbterm.String(priority, 255, 255, 0, 0, 0, 0)
	} else if priority == "Green" {
		return rgbterm.String(priority, 0, 255, 0, 0, 0, 0)
	}
	return priority
}
