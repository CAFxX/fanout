package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 string }
type GenS2 struct { F0 uint16; F1 int64; F2 int32 }
func genFn0(p0 int, p1 string, p2 bool) int8 {
    switch p1 {
    case "foo bar":
        v6 := GenS1{F0: -32768, F1: -1, F2: ""}
        _ = v6
    case "\n\t":
        p0 -= p0
    default:
        var v7 int8 = 42
        _ = v7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        switch p1 {
        case "\n\t":
            fmt.Println(p0, fi8, fi8)
        case "a":
            p0 -= (fi8 ^ 135)
        default:
            v9 := []int{4294967295}
            _ = v9
        }
        v10 := []int{}
        _ = fi8
        _ = v10
    }
    v11 := []int{p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 63
}

func main() {
    v1 := []uint32{1}
    var v2 float64 = 1e10
    var v3 uint32 = 131
    var v4 uint16 = 65535
    var v5 int = 0
    v3 = (v3 + (^(v3)))
    var v12 uint32 = (v3 << ((2) & 63))
    if (!true) {
        v12 = (v12 / (((v3 * (v12 * (^(v3))))) | 1))
    } else {
        for fi13 := 0; fi13 < 3; fi13++ {
            _ = len(v1)
            if true {
                for fi14 := 0; fi14 < 9; fi14++ {
                    if len(v1) > 0 {
                        v1[0] = v3
                    }
                    v15 := GenS2{F0: (v4 & (v4 | (v4 & (v4 << ((2) & 63))))), F1: 876, F2: -2147483648}
                    _ = fi14
                    _ = v15
                }
                _ = len(v1)
                fmt.Println(v5, v12, v12)
            }
            _ = fi13
        }
    }
    v16 := GenS1{F0: 963, F1: -27, F2: ("\n\t" + "foo bar")}
    var v17 int64 = 42
    var v18 int = (v5 << ((65535) & 63))
    if true {
        var v19 bool = false
        v18 -= (v18 << ((969) & 63))
        _ = v19
    } else {
        if true {
            v12 += 0
            fmt.Println(v17, v12, v3)
        } else {
            v12 += v12
        }
    }
    v20 := GenS2{F0: (^(v4)), F1: (v17 + 9223372036854775807), F2: -1}
    v21 := GenS1{F0: 1000, F1: 42, F2: "a"}
    v22 := GenS1{F0: 195, F1: 127, F2: "a"}
    v17 += (v17 << ((370) & 63))
    fmt.Println(v4, v4)
    fmt.Println(v12)
    var v23 uint = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
