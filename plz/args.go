package plz

import (
    "flag"
    "os"
)

func getArgs() (string, []string) {
    flag.Parse()
    if flag.NArg() == 0 {
        flag.Usage()
        os.Exit(1)
    }
    cmd := flag.Args()[0]
    rest := flag.Args()[1:]

    return cmd, rest
}
