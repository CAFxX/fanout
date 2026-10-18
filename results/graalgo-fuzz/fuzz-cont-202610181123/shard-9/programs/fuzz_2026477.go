package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 bool, p1 int16, p2 uint32) int8 {
    p1 += (p1 - (p1 | p1))
    p2 += 2
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int16 = 943
    var v2 int = 1
    var v3 int8 = -1
    v4 := []uint32{}
    var v5 uint16 = 42
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 float32 = float32(1e10)
        fmt.Println(v2)
        if (true && false) {
            if (385 == 462) {
                _ = len(v4)
            } else {
                fmt.Println(v2, v7)
                var v8 int64 = -260
                switch "a" {
                case "a":
                    var v9 uint16 = v5
                    _ = v9
                default:
                    fmt.Println(v5, v5)
                }
                _ = v8
            }
            fmt.Println(fi6)
            v10 := GenS1{F0: 1}
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    var v11 int32 = 1
    v5 = (^(v5))
    if (!(42 > 225)) {
        var v12 uint8 = 1
        _ = v12
    } else {
        v5 = v5
        v13 := GenS1{F0: 282}
        v3 += v3
        _ = v13
    }
    if ((v11 * (-(v11))) != 295) {
        var v14 float32 = (float32(0.1) - -2.25)
        _ = v14
    }
    if true {
        if len(v4) > 0 {
            v4[0] = 42
        }
    }
    v11 = -2147483648
    var v15 uint8 = 2
    var v16 float64 = -2.25
    v5 -= (v5 << ((306) & 63))
    var v17 uint16 = (v5 - v5)
    var v18 int32 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
