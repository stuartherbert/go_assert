# Assert (c) Blake Mizerany and Keith Rarick -- MIT LICENCE

## Assertions for Go packages

Forked from https://github.com/bmizerany/assert and modified to work outside of the 'go test' environment.

If you're looking for asserts to use in your tests, you should definitely use their package rather than my fork!

## Install

    $ go get github.com/stuartherbert/go_assert

## Use

    import(
        "github.com/stuartherbert/go_assert"
    )

    func DoSomething(
        expected := "foo"
        actual := "bar"

        assert.Equal(expected, actual)
    )

## Docs