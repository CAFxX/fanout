package main

import (
    "fmt"
)


func main() {
    v1 := []string{"foo bar", ("a" + "foo bar")}
    var v2 int = 2147483647
    var v3 int = v2
    var v4 bool = (0 != 1)
    var v5 float32 = float32(3.14159)
    v5 = 1.5
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int8 = 127
        var v8 int64 = 310
        v3 += int(v7)
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = len(v1)
    v3 -= -1
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 += (^(v2))
    if false {
        if v4 {
            fmt.Println(v3, v2)
            _ = len(v1)
            v3 += 32768
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
