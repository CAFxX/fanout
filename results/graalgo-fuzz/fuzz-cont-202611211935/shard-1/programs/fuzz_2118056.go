package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int8 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{2147483647, 42, 45, 42}
    var v3 float32 = -2.25
    var v4 uint8 = 127
    var v5 int8 = 42
    v6 := GenS1{F0: 32767, F1: -1, F2: int8(v4)}
    var v7 string = "foo bar"
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int64 = -562
        v1[2147483647] = false
        _ = fi8
        _ = v9
    }
    v3 = v3
    fmt.Println(v3, v4)
    switch (v7 + "") {
    case "a":
        var v10 bool = true
        _ = v10
    }
    v11 := GenS2{F0: 32767}
    v5 = 59
    _ = len(v1)
    if ("hello" == (v7 + "foo bar")) {
        var v12 float32 = (3.14159 / (-2.25 / (((0.0 / (float32(0.1) / ((v3 / v3) + (-2.25 / 1e10)))) - v3) - v3)))
        _ = v12
    } else {
        var v13 float32 = v3
        var v14 string = v7
        v15 := GenS2{F0: 62}
        _ = v13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        _ = len(v2)
        _ = fi16
    }
    v4 += v4
    var v17 uint8 = (v4 / ((1) | 1))
    switch v7 {
    case "\n\t":
        v17 += 2
    case "a":
        v17 = (^(v4))
    default:
        if false {
            v5 -= v5
        }
    }
    v17 = 127
    var v18 float64 = (1e10 / ((-2.25 - 0.1) - (1e10 - -2.25)))
    v1[439] = false
    var v19 uint8 = v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
