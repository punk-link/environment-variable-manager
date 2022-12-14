package environmentvariablemanager

type EnvironmentVariableManager interface {
	TryGet(variableName string) (bool, string)
}
