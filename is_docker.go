package osenv

const DOCKERENV_FILE string = "/.dockerenv"

func IsInDocker() bool {
	return PathExist(DOCKERENV_FILE)
}
