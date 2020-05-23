package act

import (
    "fmt"
    "path/filepath"
    "runtime"
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
        err = fmt.Errorf("%s:\n  %v", caller(2), err)
        panic(err)
    }
}

func caller(skip int) string {

    pc, f, l, _ := runtime.Caller(skip)
    dir := filepath.Base(filepath.Dir(f))
    f    = dir + "/" + filepath.Base(f)
    fn  := runtime.FuncForPC(pc)
    n   := filepath.Base(fn.Name())

    return fmt.Sprintf("%s[%s:%d]", n, f, l)

}