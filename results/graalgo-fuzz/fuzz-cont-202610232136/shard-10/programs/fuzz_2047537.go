package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int16 }
func genFn0(p0 int64) uint32 {
    fmt.Println(p0)
    v7 := []int{-1, -2147483648, 1}
    _ = p0
    _ = v7
    return 2
}

func main() {
    var v1 bool = (2147483647 != 42)
    v2 := []uint32{636, 2147483648}
    v3 := []uint32{42, 489, 42}
    var v4 uint8 = 42
    var v5 float32 = (((3.14159 - (((0.1 * (-2.25 * 0.0)) - -2.25) + (0.1 * 0.0))) / float32(((-2.25 - 3.14159) / 3.14159))) / (float32(((1.5 / 1e10) * ((3.14159 - 0.1) + (-2.25 + float64(float32(1e10)))))) * 3.14159))
    var v6 float64 = (1.5 * (1e10 / 1.5))
    if false {
        switch 42 {
        case -677:
            v6 = v6
        case 2147483647:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 4294967295:
            v3 = append(v3, 2147483647)
        }
    } else {
        var v8 uint = 65535
        _ = v8
    }
    fmt.Println(v4, v1, v6)
    if (0 >= -1) {
        var v9 uint = 42
        _ = v9
    } else {
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 uint = 0
            v12 := GenS1{F0: 604, F1: ("a" >= ("hello" + ("" + ""))), F2: -619}
            v5 -= (3.14159 + v5)
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    var v13 float32 = (v5 / (v5 / v5))
    var v14 string = ""
    switch 128 {
    case 32768:
        var v15 int64 = 0
        _ = v15
    default:
        fmt.Println(v14, v6, v14)
    }
    v5 = float32(0.1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v16 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
