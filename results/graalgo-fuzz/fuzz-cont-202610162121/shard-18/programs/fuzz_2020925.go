package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint16, p1 uint8, p2 bool) int8 {
    if (42 == -842) {
        switch 42 {
        case 32767:
            v5 := []int{127, -1, 0, -9223372036854775808}
            _ = v5
        }
    } else {
        v6 := GenS1{F0: 1, F1: 309, F2: 667}
        v7 := GenS2{F0: (p1 - (p1 * p1))}
        _ = v6
        _ = v7
    }
    fmt.Println(p1)
    var v8 int64 = -1
    var v9 int16 = -520
    var v10 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int8) int8 {
    var v11 uint32 = 361
    for fi12 := 0; fi12 < 17; fi12++ {
        fmt.Println(fi12, v11)
        for fi13 := 0; fi13 < 3; fi13++ {
            v11 -= (v11 % (((v11 >> ((255) & 63))) | 1))
            _ = fi13
        }
        _ = fi12
    }
    _ = p0
    _ = v11
    return p0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 625
    var v3 float64 = -2.25
    var v4 float64 = (v3 * (v3 * -2.25))
    var v14 int16 = 32767
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        v1[-461] = 585
    default:
        for fi15 := 0; fi15 < 9; fi15++ {
            v2 = 0
            for fi16 := 0; fi16 < 10; fi16++ {
                if (120 != 94) {
                    switch 42 {
                    case 222:
                        fmt.Println(v4, fi16)
                    case 9223372036854775807:
                        var v17 int16 = v14
                        _ = v17
                    case 32768:
                        var v18 float64 = 0.0
                        _ = v18
                    default:
                        var v19 uint32 = 578
                        _ = v19
                    }
                } else {
                    v14 = (v14 + -86)
                    v2 = 42
                    fmt.Println(fi15, v3)
                }
                var v20 int8 = 1
                _ = fi16
                _ = v20
            }
            v2 += v2
            _ = fi15
        }
    }
    var v21 int16 = (v14 >> ((0) & 63))
    v22 := GenS1{F0: 42, F1: -127, F2: 1}
    fmt.Println(v21, v4, v4)
    if true {
        v23 := []int{-9223372036854775808, -806, -391, 2147483647}
        var v24 int = -1
        fmt.Println(v2)
        _ = v23
        _ = v24
    } else {
        fmt.Println(v21, v3, v21)
    }
    v25 := GenS2{F0: 1}
    v21 = v21
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
}
