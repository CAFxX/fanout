package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }

func main() {
    var v1 uint64 = 4294967295
    var v2 int64 = -619
    v3 := []int{}
    var v4 uint = 1
    _ = len(v3)
    v3 = append(v3, int(v2))
    if (true || true) {
        var v5 int8 = 0
        fmt.Println(v1)
        _ = v5
    } else {
        var v6 float32 = (3.14159 + float32(-2.25))
        v6 -= v6
        var v7 int8 = 42
        _ = v6
        _ = v7
    }
    var v8 bool = true
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 int64 = v2
        switch (fi9 & fi9) {
        case 9223372036854775807:
            v11 := GenS1{F0: (^(v4)), F1: 101}
            _ = v11
        default:
            v8 = (v8 || (2147483647 < -1))
        }
        var v12 int = 1
        _ = fi9
        _ = v10
        _ = v12
    }
    var v13 uint8 = 119
    for fi14 := 0; fi14 < 3; fi14++ {
        v13 += v13
        var v15 int8 = -1
        var v16 uint16 = 32767
        _ = fi14
        _ = v15
        _ = v16
    }
    if (4294967295 <= 1) {
        v4 = v4
    }
    var v17 float32 = (0.1 - -2.25)
    fmt.Println(v17, v2, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
