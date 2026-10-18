package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint }
type GenS2 struct { F0 uint16; F1 uint8; F2 int16 }

func main() {
    var v1 float64 = (1e10 - (3.14159 - 3.14159))
    v2 := []int{42, 802, -2147483648}
    var v3 uint64 = 4294967295
    var v4 int = 127
    if false {
        v3 += v3
    }
    var v5 uint64 = (v3 << ((0) & 63))
    var v6 uint32 = 65535
    if (("foo bar" + ("" + "\n\t")) != "\n\t") {
        switch (^(v4)) {
        case 4294967295:
            fmt.Println(v1, v1)
        case 42:
            var v7 float64 = 3.14159
            _ = v7
        }
    } else {
        var v8 bool = (true && false)
        _ = v8
    }
    _ = len(v2)
    v4 = v4
    if true {
        fmt.Println(v4, v1)
        v5 = (v3 * v5)
        var v9 bool = ((!(!(861 != 0))) || true)
        _ = v9
    }
    var v10 uint64 = v5
    v6 += (^(v6))
    v5 -= (v10 % ((v5) | 1))
    v11 := GenS1{F0: 137, F1: 42, F2: 42}
    v5 = 4294967295
    if true {
        _ = len(v2)
        switch "a" {
        case "\n\t":
            v10 = 613
        default:
            v12 := GenS2{F0: 895, F1: 1, F2: 42}
            _ = v12
        }
    }
    var v13 int = v4
    for fi14 := 0; fi14 < 14; fi14++ {
        fmt.Println(v5, fi14)
        switch (^(v4)) {
        case 32767:
            v2 = append(v2, v4)
        case -9223372036854775808:
            var v15 int8 = 42
            _ = v15
        }
        v3 = (v3 | v5)
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
