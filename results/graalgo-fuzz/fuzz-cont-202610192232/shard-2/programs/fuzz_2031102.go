package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 string }

func main() {
    var v1 float64 = 1e10
    var v2 string = "\n\t"
    var v3 int = -865
    var v4 uint8 = 128
    v5 := []int{(v3 % ((32767) | 1)), (v3 << ((2) & 63))}
    v6 := make(map[int]int)
    var v7 int16 = 32767
    v3 -= v3
    fmt.Println(v1)
    var v8 int64 = 451
    v9 := GenS1{F0: -17}
    delete(v6, (v3 * v3))
    _ = len(v5)
    fmt.Println(v7, v3)
    v10 := GenS1{F0: -79}
    for fi11 := 0; fi11 < 2; fi11++ {
        v7 = (-(v7))
        _ = fi11
    }
    var v12 uint32 = 104
    if len(v5) > 0 {
        v5[0] = v3
    }
    var v13 int16 = (-(v7))
    switch v2 {
    case "foo bar":
        v8 = v8
    case "":
        switch -445 {
        case 641:
            _ = len(v6)
        case -9223372036854775808:
            if ((v4 / ((v4) | 1)) == (^(v4))) {
                fmt.Println(v1)
                delete(v6, -2147483648)
                switch v2 {
                case "foo bar":
                    v13 = (v13 % (((^(v13))) | 1))
                }
            }
        case -1:
            v14 := GenS1{F0: 100}
            _ = v14
        default:
            v13 = (^(v13))
        }
    default:
        v2 = v2
    }
    var v15 int16 = v7
    v16 := GenS2{F0: 0, F1: v2}
    if false {
        var v17 uint = 0
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
