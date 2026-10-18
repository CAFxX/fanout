package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 1
    var v3 string = ("" + "")
    var v4 uint64 = (^(v2))
    var v5 uint8 = 1
    var v6 int64 = 42
    if (false && (1 > 1)) {
        var v7 int = int(v5)
        _ = v7
    }
    fmt.Println(v4)
    var v8 float64 = (1.5 - 1e10)
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint64 = 2
        _ = fi9
        _ = v10
    }
    v3 = v3
    v1[1] = (false && (4294967295 < 0))
    v3 = v3
    v3 = "\n\t"
    v8 += -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
