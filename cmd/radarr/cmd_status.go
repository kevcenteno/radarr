package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/urfave/cli/v2"
)

var statusCommand *cli.Command = &cli.Command{
	Name:   "status",
	Usage:  "Get your Radarr instance status",
	Action: getStatus,
}

func init() {
	app.Commands = append(app.Commands, statusCommand)
}

func getStatus(c *cli.Context) error {
	status, err := radarrClient.SystemStatus.Get()
	if err != nil {
		return err
	}

	// Print as JSON if provided
	if c.Bool("json") {
		data, err := json.Marshal(status)
		if err != nil {
			return err
		}

		fmt.Println(string(data))
		return nil
	}

	v := reflect.ValueOf(*status)
	typeOfStatus := v.Type()

	for i := 0; i < v.NumField(); i++ {
		t.Append([]string{
			fmt.Sprintf("%s:", typeOfStatus.Field(i).Name),
			fmt.Sprintf("%v", v.Field(i).Interface()),
		})
	}

	t.SetCenterSeparator("")
	t.SetColumnSeparator("")
	t.SetRowSeparator("")
	t.Render()
	return nil
}
