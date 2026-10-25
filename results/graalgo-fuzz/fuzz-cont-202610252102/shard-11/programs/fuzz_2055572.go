package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint32 }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]int)
    var v3 uint8 = 42
    var v4 uint32 = 2147483648
    var v5 int16 = 32767
    var v6 int16 = (v5 & 1)
    var v7 int64 = -1
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(v1, v3)
        _ = fi8
    }
    v9 := GenS1{F0: 0, F1: 739, F2: (v4 % ((42) | 1))}
    v10 := GenS1{F0: 1, F1: 65535, F2: v4}
    if false {
        var v11 string = ("a" + "hello")
        _ = v11
    } else {
        var v12 bool = false
        _ = v12
    }
    _ = len(v2)
    if ((v7 << ((255) & 63)) <= int64(v1)) {
        var v13 uint64 = 0
        _ = v13
    } else {
        var v14 int32 = 1
        _ = v14
    }
    fmt.Println(v4)
    v5 = v6
    v15 := GenS1{F0: 2147483647, F1: 2, F2: v4}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
