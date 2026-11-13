package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
func genFn0(p0 int32, p1 int, p2 string) int64 {
    var v4 int8 = 1
    if (false || (false && true)) {
        var v5 float32 = 0.1
        _ = v5
    } else {
        for fi6 := 0; fi6 < 1; fi6++ {
            v4 += 100
            _ = fi6
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return int64(p0)
}
func genFn1(p0 uint16, p1 uint32) uint8 {
    v7 := GenS1{F0: -31, F1: ("hello" + "\n\t")}
    var v8 uint = 893
    v9 := GenS1{F0: 85, F1: "foo bar"}
    v10 := []int{448, 1}
    if ((((!false) && false) || false) || ((!((!false) || false)) || true)) {
        v11 := GenS1{F0: 1, F1: ""}
        fmt.Println(p0)
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 128
}

func main() {
    var v1 int = -1
    var v2 uint8 = 127
    var v3 float32 = 0.0
    v1 = 572
    var v12 int8 = 100
    v2 = v2
    v1 -= (v1 & -9223372036854775808)
    v13 := []int{}
    if false {
        fmt.Println(v3, v3, v12)
    } else {
        if ((!(4294967295 >= 18446744073709551615)) && false) {
            var v14 float64 = -2.25
            _ = v14
        } else {
            for fi15 := 0; fi15 < 12; fi15++ {
                switch v1 {
                case 0:
                    var v16 float64 = (1e10 - 1.5)
                    _ = v16
                }
                v12 = v12
                switch "a" {
                case "foo bar":
                    var v17 float32 = float32(-2.25)
                    _ = v17
                case "\n\t":
                    var v18 uint16 = 2
                    _ = v18
                }
                _ = fi15
            }
        }
        v12 -= (v12 ^ (v12 / (((^(v12))) | 1)))
    }
    var v19 float32 = v3
    v1 = (-(v1))
    fmt.Println(v19, v3)
    fmt.Println(v12, v1)
    var v20 int = (v1 << ((2) & 63))
    var v21 uint32 = 475
    switch ("" + "\n\t") {
    case "a":
        v21 += 436
    case "":
        v2 -= (^(v2))
    case "hello":
        v22 := GenS1{F0: (v12 / ((v12) | 1)), F1: "a"}
        _ = v22
    default:
        v1 += 32768
    }
    if len(v13) > 0 {
        v13[0] = -9223372036854775808
    }
    var v23 float32 = (v19 / (v19 - v3))
    var v24 string = "hello"
    var v25 float64 = (3.14159 + 1e10)
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 uint64 = 612
        v13 = append(v13, v1)
        _ = len(v13)
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
