package funcs

import (
	"fmt"
)

func CheckCollector() {

	output := make(map[string]bool)

	output["rds  "] = len(RDSMetrics()) > 0

	for k, v := range output {
		status := "fail"
		if v {
			status = "ok"
		}
		fmt.Println(k, "...", status)
	}
}
