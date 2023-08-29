# cacheLib
cash library for cashing key, value data

how it works:
```go
func main() {
cacshLib.Set("1", time.Second*5, "1")
cacshLib.Set("2", time.Second*1, "2")
cacshLib.Set("3", time.Second*4, "4")

time.Sleep(time.Second * 2)
fmt.Println(cacshLib.Get("1"))
fmt.Println(cacshLib.Get("2"))
fmt.Println(cacshLib.Get("3"))
}

```