package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int; F2 uint }
func genFn0(p0 int16) uint {
    v5 := GenS1{F0: 486}
    v6 := []int{-1}
    p0 += p0
    _ = p0
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    var v1 int64 = 693
    v2 := []string{("" + ("" + "\n\t")), "foo bar"}
    var v3 int = 1
    var v4 int8 = 127
    var v7 uint16 = 974
    var v8 int16 = 0
    v9 := GenS2{F0: (^(v4)), F1: v3, F2: 255}
    fmt.Println(v7)
    v1 -= v1
    v3 = (v3 >> ((4294967295) & 63))
    var v10 float64 = (((-2.25 + (1.5 - 0.0)) - 0.0) / -2.25)
    for fi11 := 0; fi11 < 19; fi11++ {
        fmt.Println(v8, v1, v10)
        fmt.Println(v8)
        _ = fi11
    }
    if true {
        fmt.Println(v8)
        switch -474 {
        case 539:
            v8 = -1
        case 4294967295:
            var v12 bool = true
            _ = v12
        default:
            for fi13 := 0; fi13 < 15; fi13++ {
                var v14 int = v3
                _ = fi13
                _ = v14
            }
        }
        v3 = (v3 * v3)
    } else {
        var v15 uint8 = 0
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
