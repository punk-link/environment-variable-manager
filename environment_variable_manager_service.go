package environmentvariablemanager

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type EnvironmentVariableManagerService struct{}

func init() {
	localVariables := getLocalVariables()
	systemVariables := getSystemVariables()

	totalLength := len(localVariables) + len(systemVariables)

	_environmentVariables = make(map[string]string, totalLength)
	for key, value := range localVariables {
		_environmentVariables[key] = value
	}

	for key, value := range systemVariables {
		_environmentVariables[key] = value
	}
}

func New() EnvironmentVariableManager {
	return &EnvironmentVariableManagerService{}
}

func (t *EnvironmentVariableManagerService) TryGet(variableName string) (bool, string) {
	if result, ok := _environmentVariables[variableName]; ok {
		return true, result
	}

	logger := log.Default()
	logger.Printf("no configuration value has been found for '%s'", variableName)
	return false, ""
}

func getLocalVariables() map[string]string {
	file, err := os.Open(".env")
	if err != nil {
		logger := log.Default()
		logger.Printf("can't read a local configuration file: %s", err.Error())

		return make(map[string]string, 0)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	slice := make([]string, 0)
	for fileScanner.Scan() {
		slice = append(slice, fileScanner.Text())
	}

	file.Close()

	return splitKeysAndValues(slice)
}

func getSystemVariables() map[string]string {
	variables := os.Environ()
	return splitKeysAndValues(variables)
}

func splitKeysAndValues(source []string) map[string]string {
	result := make(map[string]string, len(source))

	for _, value := range source {
		split := strings.SplitN(value, "=", 2)
		result[split[0]] = split[1]
	}

	return result
}

var _environmentVariables map[string]string
