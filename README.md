# kestra-go

### Summary

This is the unofficial Kestra GoLang support module. 

Currently this module's functionality includes allowing users to log messages with either TRACE, DEBUG, INFO, WARNING, ERROR, or CRTIICAL level and it'll log in a format that is understandable by your Kestra instance.

The roadmap for this package to add better documentation and unit tests in the near future, and also adding the ability to send outputs to your Kestra instance. 

### Installation

`go get github.com/japerry911/kestra-go@v0.1.3`

### Example

```go
import "github.com/japerry911/kestra-go"

k := kestra.NewKestra()

l := k.Logger

l.Trace("Test Trace")
l.Debug("Test Debug")
l.Info("Test Info")
l.Warn("Test Warning")
l.Error("Test Error")
```