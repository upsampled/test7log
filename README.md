# test7log
7 Level Severity Logger/ Logrus Field Logger for Go testing

This allows for explicit control of logging in `go test` by using `-v`.

`t.Log` does not handle concurrency well so if you need to use typical STDERR logging instead, set the `TEST7LOG_OUT` variable to `STDERR`.

```
TEST7LOG_OUT=STDERR go test
```

You can set the log level with `TEST7LOG_LVL` to any of the values mentioned in `logrus.ParseLevel` to adjust the minimal severity level to log. 


# Example

```go
func HelloWorldLog(l *logrus.FieldLogger){
    l.Info("HelloWorld")
}

func TestHellowWorldLog(t *testing.T){
    lg := NewTestFieldLogger(t)
    HelloWorldLog(lg)
}

```