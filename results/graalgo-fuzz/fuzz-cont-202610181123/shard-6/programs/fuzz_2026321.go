package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 42
    var v3 uint64 = 595
    var v4 int32 = 896
    v4 = v4
    fmt.Println(v4, v3, v2)
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 uint32 = 1
        for fi7 := 0; fi7 < 14; fi7++ {
            fmt.Println(v3, v6, v2)
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    var v8 int8 = -26
    if false {
        v8 -= v8
        v9 := GenS1{F0: -507, F1: -1, F2: 0}
        _ = v9
    } else {
        _ = len(v1)
        var v10 uint8 = 52
        _ = v10
    }
    var v11 float64 = 1e10
    switch "" {
    case "":
        v3 = 2
    }
    v12 := GenS1{F0: (^(v2)), F1: -217, F2: 65535}
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int64 = 400
        _ = fi13
        _ = v14
    }
    var v15 int64 = -9223372036854775808
    v8 = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
