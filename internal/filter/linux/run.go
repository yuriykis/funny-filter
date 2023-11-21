package linux

import (
	"os/exec"
	"strings"

	"github.com/yuriykis/funny-filter/log"
)

func Run(command string) (string, error) {

	log.WithFields(log.Fields{
		"command": command,
	}).Info("Running command")

	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("Command failed with error: %s: %s", err, string(output))
	} else {
		log.Debugf("Command output: %s", string(output))
	}
	return string(output), err
}

func Build(params ...string) string {

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
