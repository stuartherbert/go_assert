package assert

import (
    "testing"
)

func TestLineNumbers(t *testing.T) {
    Equal("foo", "foo", "msg!")
    //Equal(t, "foo", "bar", "this should blow up")
}

func TestNotEqual(t *testing.T) {
    NotEqual("foo", "bar", "msg!")
    //NotEqual(t, "foo", "foo", "this should blow up")
}
