package plz

import (
    "fmt"
//    "errors"
    "github.com/logrusorgru/aurora"
)

func print_text(text string, color_code aurora.Color) {
    fmt.Println(aurora.Colorize(text, color_code))
}

func print_error(text string) {
    print_text(text, aurora.RedFg)
}

func print_error_dim(text string) {
    print_text(text, aurora.RedFg|aurora.FaintFm)
}

func print_info_dim(text string) {
    print_text(text, aurora.CyanFg|aurora.FaintFm)
}
