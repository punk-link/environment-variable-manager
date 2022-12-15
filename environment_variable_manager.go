package environmentvariablemanager

type EnvironmentVariableManager interface {
	TryGet(variableName string) (string, bool)
}
