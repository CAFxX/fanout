package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int64 }

func main() {
    var v1 uint16 = 934
    var v2 float32 = 1e10
    var v3 float32 = 1.5
    var v4 uint16 = (^(v1))
    v5 := GenS1{F0: 1, F1: 65535, F2: 9223372036854775807}
    v6 := []int{-342, 0, 366, -1}
    v4 = v1
    var v7 uint8 = 12
    fmt.Println(v1, v3, v1)
    v7 += (^(v7))
    for fi8 := 0; fi8 < 14; fi8++ {
        if (!(!true)) {
            v9 := GenS1{F0: -16, F1: 0, F2: 9223372036854775807}
            _ = v9
        } else {
            v4 = v1
            _ = len(v6)
        }
        _ = fi8
    }
    if len(v6) > 0 {
        v6[0] = -9223372036854775808
    }
    var v10 int64 = 0
    fmt.Println(v3, v2, v1)
    var v11 int = 2147483647
    fmt.Println(v4)
    v10 = v10
    if (40 <= 127) {
        var v12 int = (v11 + v11)
        var v13 int64 = -1
        var v14 uint16 = 2
        _ = v12
        _ = v13
        _ = v14
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
