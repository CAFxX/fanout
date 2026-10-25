package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int64 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int32 = -389
    v2 := make(map[uint32]bool)
    var v3 int8 = 26
    var v4 string = "foo bar"
    var v5 uint8 = 127
    for fi6 := 0; fi6 < 9; fi6++ {
        v4 = v4
        _ = fi6
    }
    v1 = v1
    v4 = ("a" + "")
    switch v4 {
    case "foo bar":
        v2[2147483648] = true
    default:
        var v7 float64 = 1e10
        _ = v7
    }
    var v8 uint32 = 342
    if false {
        var v9 float32 = -2.25
        _ = v9
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            v4 = (v4 + v4)
            v1 += int32(fi10)
            _ = fi10
        }
        if (true || true) {
            v1 -= v1
        } else {
            v11 := GenS1{F0: 4294967295, F1: uint16(v8), F2: 0}
            v12 := GenS2{F0: v3}
            v3 = v3
            _ = v11
            _ = v12
        }
        var v13 float64 = (float64((0.0 + (0.0 - (float32((1e10 - (-2.25 + ((0.1 * 1.5) / 0.1)))) + 0.1)))) / (0.0 + 3.14159))
        _ = v13
    }
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
