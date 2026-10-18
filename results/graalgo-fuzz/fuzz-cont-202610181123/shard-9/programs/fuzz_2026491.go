package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint16; F1 int; F2 bool }
func genFn0() uint16 {
    v6 := GenS1{F0: 2147483648}
    var v7 string = ("a" + "foo bar")
    _ = v6
    _ = v7
    return 520
}
func genFn1(p0 int8, p1 uint32) uint32 {
    fmt.Println(p0, p1)
    p0 = -128
    _ = p0
    _ = p1
    return 841
}

func main() {
    var v1 int32 = -644
    var v2 uint = 119
    var v3 uint16 = 246
    var v4 float32 = 0.1
    var v5 uint16 = v3
    v8 := GenS1{F0: 804}
    v9 := GenS1{F0: 4294967295}
    v10 := []int{127, -1, -2147483648, 2147483647}
    v11 := GenS1{F0: 2147483647}
    switch 127 {
    case 32768:
        v5 += v5
    case -2147483648:
        if len(v10) > 0 {
            v10[0] = 2147483647
        }
    }
    var v12 uint = v2
    v4 += float32(0.1)
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            if (533 > 767) {
                v15 := GenS2{F0: 32767, F1: v1}
                _ = v15
            }
            v3 = v3
            if false {
                fmt.Println(v5, v1)
                fmt.Println(v1, v3)
            } else {
                switch 128 {
                case 459:
                    v4 -= v4
                case 69:
                    v10 = append(v10, (fi13 >> ((uint(v3)) & 63)))
                default:
                    fmt.Println(v2, v1, v1)
                    v16 := GenS3{F0: 985, F1: (fi13 >> ((816) & 63)), F2: true}
                    var v17 uint = 255
                    _ = v16
                    _ = v17
                }
                v12 -= 728
                if (false && false) {
                    v18 := GenS2{F0: -32768, F1: (v1 >> ((18446744073709551615) & 63))}
                    _ = v18
                }
            }
            _ = fi14
        }
        fmt.Println(v1)
        fmt.Println(v5, v12, v4)
        _ = fi13
    }
    if ((v5 - (^(v3))) < genFn0()) {
        fmt.Println(v4, v1)
        v19 := GenS1{F0: 0}
        var v20 uint32 = 997
        _ = v19
        _ = v20
    } else {
        fmt.Println(v2, v5)
        var v21 uint64 = 0
        _ = v21
    }
    switch "foo bar" {
    case "hello":
        v22 := GenS2{F0: 574, F1: v1}
        _ = v22
    case "foo bar":
        v23 := GenS1{F0: 311}
        _ = v23
    }
    if (167 < 181) {
        var v24 int8 = -86
        _ = v24
    }
    var v25 uint = v12
    fmt.Println(v5, v25, v3)
    v26 := GenS2{F0: 0, F1: v1}
    fmt.Println(v12, v5, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
