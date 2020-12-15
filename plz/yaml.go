package plz

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

type Command struct {
  Id string `yaml:"id"`
  Cmd string `yaml:"cmd"`
  Cmds []string `yaml:"cmds"`
}

type Config struct {
  Commands []Command `commands`
}

func parseYaml(filePath string) (Config, error) {
    var config Config

    // open config file
    file, err := ioutil.ReadFile(filePath)
    if err != nil {
        // Return an empty struct
        return config, err
    }


    yamlErr := yaml.Unmarshal(file, &config)
    if yamlErr != nil {
        return config, yamlErr
    }

    validateErr := validateYaml(config)
    return config, validateErr
}

func validateYaml(config Config) error {
    for _, command := range config.Commands {
        if len(command.Cmd) > 0 && len(command.Cmds) > 0 {
            return fmt.Errorf("Config error, id '%s' has both 'cmd' and 'cmds', must pick one", command.Id)
        }
    }
    return nil
}
