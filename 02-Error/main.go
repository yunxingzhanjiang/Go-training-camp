package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type dao struct {
	url string
}

func (d *dao) queryRow() (string, error) {
	return "", sql.ErrNoRows
}

func service(query string, d *dao) error {

	row, err := d.queryRow()
	if err != nil {
		return errors.Wrap(err, query)
	}

	// Handle row.
	result := fmt.Sprintf("Row is : %v", row)
	fmt.Println(result)

	return nil
}

func main() {
	d := dao{
		url: "db://127.0.0.1:3304",
	}

	err := service("query", &d)
	if err != nil {
		fmt.Printf("Stack trace: %+v", err)
	}
}
