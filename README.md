# TryHard

<p align="left">
  <a href="https://github.com/devit-tel/tryhard"><img alt="GitHub Actions status" src="https://github.com/devit-tel/tryhard/workflows/go-test/badge.svg"></a>
</p>
TryHard is a library for handle retry support key and automatic clear (TTL)

##### Support
- [x] Config maximum retry, backoff value
- [x] Support In-memory Store
- [ ] Support Redis Store
- [x] Support retry with key & ttl


---

### Installation

```shell script
    go get -u github.com/devit-tel/tryhard
```


---
### Usage

simple use case if function return error ```.Do()``` will retry until lap exceed maximum retry
```go
        try := New(inmem.New())
	try.Do(func() error {
                // Do something ...
		return nil
	})
```

retry with key
```go
        try := New(inmem.New())
        try.DoWithKey("key_1", func() error {
            // Do something ...
            return nil
        }
```

init with config
```go
	try := New(inmem.New(), SetMaximumRetry(3), SetBackOff(10))
```

---


### Test
use ```Test()``` for testing this function will create new TryHard with zero value 
```go
        try := Test(inmem.New())
```
