package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int16 }

func main() {
    var v1 int = -1
    v2 := make(map[string]string)
    var v3 int32 = -2147483648
    v4 := make(map[uint32]string)
    var v5 uint = 887
    v3 -= (v3 % ((v3) | 1))
    delete(v4, 197)
    v3 = -2147483648
    if (0 > 1) {
        v1 -= v1
        var v6 int32 = 1
        _ = v6
    }
    if false {
        if false {
            fmt.Println(v5, v3, v1)
            v7 := GenS1{F0: 0, F1: 42, F2: -38}
            v3 = (v3 * (v3 * v3))
            _ = v7
        } else {
            v5 -= v5
            var v8 int8 = 42
            _ = v8
        }
    } else {
        fmt.Println(v1, v5)
        var v9 uint = v5
        v10 := GenS1{F0: 0, F1: 2, F2: -203}
        _ = v10
        _ = v9
    }
    v2["hello"] = "a"
    var v11 int8 = -1
    v12 := GenS1{F0: 879, F1: 1, F2: -913}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
