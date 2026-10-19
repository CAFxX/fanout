package main

import (
    "fmt"
)


func main() {
    v1 := []string{"a", ("foo bar" + "a")}
    var v2 int64 = -608
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 14; fi4++ {
        if false {
            v2 = 42
            var v5 bool = true
            _ = v5
        }
        fmt.Println(v2, fi4)
        var v6 uint16 = 248
        _ = fi4
        _ = v6
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 += v2
    v2 -= 0
    _ = len(v1)
    var v7 int64 = v2
    fmt.Println(v7, v7)
    var v8 uint8 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
