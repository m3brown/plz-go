package plz

import (
    "fmt"
    "os"
    "path"
    "os/exec"
    "strings"
    "github.com/lithammer/dedent"
)

var configFilename string = ".plz.yaml"
var docURL string = "https://github.com/m3brown/plz"


func gitRoot() (string, error) {
	path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", err
	}
    fmt.Println(strings.TrimSpace(string(path)))
	return strings.TrimSpace(string(path)), nil
}

func fileExists(name string) bool {
    _, err := os.Stat(name)
    if os.IsNotExist(err) {
        return false
    }
    return err == nil
}


func printConfigError() {
    print_error(fmt.Sprintf(dedent.Dedent(`
        plz must be run from:
          a) a directory that has a valid .plz.yaml file or
          b) within a git repo that contains a .plz.yaml in the repo root path
        For more information, visit %s
    `), docURL))
}

func configFilePath() (string, error) {
    fmt.Println(fileExists(configFilename))
    if fileExists(configFilename) {
        return "", nil
    } else {
        root, gitErr := gitRoot()
        if gitErr == nil && fileExists(path.Join(root, configFilename)) {
            return root, nil
        }
    }
    printConfigError()
    return "", fmt.Errorf("Unable to locate config file")

}

func plzConfig() (Config, error) {
    filePath, err := configFilePath()

    if err != nil {
        var config Config
        return config, err
    }

    return parseYaml(filePath + configFilename)
}
