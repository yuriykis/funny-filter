package filter

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/yuriykis/funny-filter/log"
)

func run(command string) (string, error) {

	log.WithFields(log.Fields{
		"command": command,
	}).Info("Running command")

	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
	} else {
		fmt.Println(string(output))
	}
	return string(output), err
}

func build(params ...string) string {

	log.WithFields(log.Fields{
		"params": params,
	}).Info("Building command")

	var validParams []string
	for _, param := range params {
		if param != "" {
			validParams = append(validParams, param)
		}
	}
	return strings.Join(validParams, " ")
}
