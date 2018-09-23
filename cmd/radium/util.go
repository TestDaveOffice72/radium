package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/shivylp/radium"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

func writeOut(cmd *cobra.Command, v interface{}) {
	ugly, _ := cmd.Flags().GetBool("ugly")
	asJSON, _ := cmd.Flags().GetBool("json")
	if ugly {
		rawDump(v, asJSON)
	} else {
		tryPrettyPrint(v)
	}
}

func tryPrettyPrint(v interface{}) {
	switch v.(type) {
	case radium.Article, *radium.Article:
		fmt.Println((v.(radium.Article)).Content)
	case []radium.Article:
		results := v.([]radium.Article)
		if len(results) == 1 {
			tryPrettyPrint(results[0])
		} else {
			rawDump(results, false)
		}
	case error:
		fmt.Printf("error: %s\n", v)
	default:
		rawDump(v, false)
	}
}

func rawDump(v interface{}, asJSON bool) {
	var data []byte
	if asJSON {
		data, _ = json.MarshalIndent(v, "", "    ")
	} else {
		data, _ = yaml.Marshal(v)
	}
	os.Stdout.Write(data)
	os.Stdout.Write([]byte("\n"))
}
