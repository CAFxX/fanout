package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int8 }
type GenS2 struct { F0 bool; F1 uint; F2 bool }
func genFn0() uint32 {
    var v5 uint16 = 0
    v5 += 1
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: 128, F1: 100, F2: 1}
        _ = fi6
        _ = v7
    }
    v8 := []int{42}
    _ = v5
    _ = v8
    return 995
}
func genFn1(p0 uint64) int {
    p0 = p0
    var v9 int64 = -9223372036854775808
    if (false || false) {
        v9 = v9
        var v10 bool = false
        _ = v10
    }
    _ = p0
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 bool = (4294967295 != 90)
    var v2 int64 = 9223372036854775807
    var v3 uint64 = 4294967295
    var v4 int16 = -860
    v4 += v4
    switch "a" {
    case "hello":
        var v11 uint16 = 32767
        _ = v11
    default:
        v12 := GenS2{F0: (((0.1 * ((3.14159 / 3.14159) * (1e10 + -2.25))) / -2.25) > (0.1 + 0.1)), F1: 0, F2: true}
        _ = v12
    }
    fmt.Println(v3, v2, v2)
    v4 += 0
    fmt.Println(v4, v2)
    v13 := GenS1{F0: 244, F1: -128, F2: 100}
    v2 = -261
    for fi14 := 0; fi14 < 19; fi14++ {
        fmt.Println(v4)
        _ = fi14
    }
    v3 = 65
    v3 = 18446744073709551615
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 bool = (!true)
        var v17 int = -554
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := []int{}
    if (-1 != 0) {
        if v1 {
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            v18 = append(v18, 2147483647)
        } else {
            if true {
                if len(v18) > 0 {
                    v18[0] = -693
                }
                fmt.Println(v4, v4)
                v2 += (v2 << ((255) & 63))
            }
            if len(v18) > 0 {
                v18[0] = -9223372036854775808
            }
        }
    } else {
        if v1 {
            v3 -= 18446744073709551615
        } else {
            fmt.Println(v1)
        }
        v1 = false
    }
    fmt.Println(v4, v3, v1)
    v1 = (!v1)
    v19 := GenS2{F0: false, F1: 65535, F2: true}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
