package main

import (
    "fmt"
)


func main() {
    v1 := []string{"foo bar", "hello", (((("foo bar" + ("foo bar" + "foo bar")) + "foo bar") + "a") + "foo bar")}
    var v2 string = "foo bar"
    var v3 string = "a"
    var v4 int16 = 42
    var v5 float64 = (1e10 + (1.5 / -2.25))
    v2 = v3
    v4 = (v4 << ((42) & 63))
    if false {
        v3 = v3
    }
    v5 += v5
    var v6 int64 = 1
    _ = len(v1)
    var v7 float64 = v5
    if true {
        var v8 bool = false
        var v9 uint = 477
        for fi10 := 0; fi10 < 2; fi10++ {
            fmt.Println(v6, v9)
            _ = fi10
        }
        _ = v8
        _ = v9
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
