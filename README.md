# cacheLib
cash library for cashing key, value data

how it works:
```go
func main() {
cacheLib.Set("1", time.Second*5, "1")
cacheLib.Set("2", time.Second*1, "2")
cacheLib.Set("3", time.Second*4, "4")

time.Sleep(time.Second * 2)
fmt.Println(cacheLib.Get("1"))
fmt.Println(cacheLib.Get("2"))
fmt.Println(cacheLib.Get("3"))
}

```