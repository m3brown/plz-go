package plz

import (
    "fmt"
    "errors"
    "strings"
    "github.com/commander-cli/cmd"
    "github.com/lithammer/dedent"
//    "log"
//    "os"
)

func exec_header(command string) {
    s := `
      ===============================================================================
      Running command: %s
      ===============================================================================
      `
    print_info_dim(dedent.Dedent(fmt.Sprintf(s, command)))
}

func execIndividual(command string, args []string, cwd string) error {
    fullArray := append([]string{command}, args...)
    fullCmd := strings.Join(fullArray[:], " ")
    exec_header(fullCmd)
    var options []func(*cmd.Command)
    options = append(options, cmd.WithStandardStreams)
    if len(cwd) > 0 {
        options = append(options, cmd.WithWorkingDir(cwd))
    }
    c := cmd.NewCommand(fullCmd, options...)

    err := c.Execute()
    if err != nil {
        return err
    } else if c.ExitCode() > 0 {
        return errors.New(fmt.Sprintf("%d", c.ExitCode()))
    }

    return nil
}

func execCommand(command Command, args []string, cwd string) error {
    var err *error = nil

    // Path 1: individual cmd
    if len(command.Cmd) > 0 {
        return execIndividual(command.Cmd, args, cwd)
    }

    // Path 2: array of cmds
    for _, command := range command.Cmds {
        if err == nil {
            loopErr := execIndividual(command, nil, cwd)
            err = &loopErr
        } else {
            s := `
              ===============================================================================
              Skipping command due to previous errors: '%s'
              ===============================================================================
            `
            print_error_dim(dedent.Dedent(fmt.Sprintf(s, command)))
        }
    }
    return *err
}
