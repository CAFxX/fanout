package main

import (
    "fmt"
)


func main() {
    var v1 uint64 = 1
    v2 := make(map[string]uint32)
    var v3 int = 32767
    v4 := []string{}
    var v5 uint = 140
    var v6 uint = (v5 % ((v5) | 1))
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    var v7 uint16 = 0
    fmt.Println(v3, v6, v7)
    v6 = v5
    var v8 int64 = 640
    if (1000 == -1) {
        var v9 int64 = v8
        _ = v9
    }
    _ = len(v2)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
