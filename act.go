package act

import (
    "fmt"
)

func Assert(t bool, msg string) {
    if !t {
        panic(msg)
    }
}

func Catch(err *error) {
    r := recover()
    if r != nil {
        *err = fmt.Errorf("%v", r)
    }
}

func Try(err error) {
    if err != nil {
        panic(err)
    }
}
