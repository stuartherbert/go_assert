package assert

import (
    "testing"
)

func TestLineNumbers(t *testing.T) {
    Equal(FailNow, "foo", "foo", "msg!")
    //Equal(t, "foo", "bar", "this should blow up")
}

func TestNotEqual(t *testing.T) {
    NotEqual(FailNow, "foo", "bar", "msg!")
    //NotEqual(t, "foo", "foo", "this should blow up")
}
