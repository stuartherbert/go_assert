// Assert (c) Blake Mizerany and Keith Rarick -- MIT LICENCE
//
// Forked and modified to run outside of the 'go test' environment
package assert

// Testing helpers for doozer.

import (
    "fmt"
    "github.com/kr/pretty"
    "os"
    "reflect"
    "runtime"
)

func getCaller(cd int) (string, string, int) {
    pc, file, line, ok := runtime.Caller(cd + 1)
    if !ok {
        return "unknown", "unknown", 0
    }

    them := runtime.FuncForPC(pc)
    if them == nil {
        return "unnamed", file, line
    }

    return them.Name(), file, line
}

func FailNow() {
    os.Exit(1)
}

func Error(args ...interface{}) {
    fmt.Println(args...)
}

func assert(result bool, f func(), failFn func(), cd int) {
    if !result {
        _, file, line, _ := runtime.Caller(cd + 1)
        fmt.Printf("%s:%d", file, line)
        f()
        failFn()
    }
}

func equal(failFn func(), exp, got interface{}, cd int, args ...interface{}) {
    fn := func() {
        for _, desc := range pretty.Diff(exp, got) {
            Error("!", desc)
        }
        if len(args) > 0 {
            Error("!", " -", fmt.Sprint(args...))
        }
    }
    result := reflect.DeepEqual(exp, got)
    assert(result, fn, failFn, cd+1)
}

func tt(result bool, cd int, args ...interface{}) {
    fn := func() {
        Error("!  Failure")
        if len(args) > 0 {
            Error("!", " -", fmt.Sprint(args...))
        }
    }
    assert(result, fn, FailNow, cd+1)
}

func T(result bool, args ...interface{}) {
    tt(result, 1, args...)
}

func Tf(result bool, format string, args ...interface{}) {
    tt(result, 1, fmt.Sprintf(format, args...))
}

func Equal(failFn func(), exp, got interface{}, args ...interface{}) {
    equal(failFn, exp, got, 1, args...)
}

func Equalf(failFn func(), exp, got interface{}, format string, args ...interface{}) {
    equal(failFn, exp, got, 1, fmt.Sprintf(format, args...))
}

func NotEqual(failFn func(), exp, got interface{}, args ...interface{}) {
    fn := func() {
        Error("!  Unexpected: <%#v>", exp)
        if len(args) > 0 {
            Error("!", " -", fmt.Sprint(args...))
        }
    }
    result := !reflect.DeepEqual(exp, got)
    assert(result, fn, failFn, 1)
}

func Panic(err interface{}, fn func()) {
    defer func() {
        equal(FailNow, err, recover(), 3)
    }()
    fn()
}
