package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint8; F1 int32 }
func genFn0(p0 int64, p1 bool, p2 uint) int16 {
    var v6 float64 = 3.14159
    v7 := GenS1{F0: 42}
    for fi8 := 0; fi8 < 1; fi8++ {
        p2 = p2
        v9 := GenS2{F0: 160, F1: p0}
        var v10 int32 = 2147483647
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 uint32) uint64 {
    var v11 string = ""
    v11 = "a"
    v12 := []int{0}
    v11 = (v11 + "")
    fmt.Println(p0, v11)
    _ = p0
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 uint32 = 651
    var v2 uint8 = 61
    var v3 int16 = 1000
    var v4 int16 = v3
    var v5 int16 = v4
    fmt.Println(v2, v4)
    v13 := GenS1{F0: -2147483648}
    fmt.Println(v2)
    for fi14 := 0; fi14 < 6; fi14++ {
        fmt.Println(v4)
        for fi15 := 0; fi15 < 18; fi15++ {
            switch "a" {
            case "a":
                var v16 uint = 129
                _ = v16
            case "\n\t":
                v17 := []int{}
                _ = v17
            case "":
                v3 += v4
            }
            fmt.Println(v2, fi15, v2)
            if (32767 >= 627) {
                v18 := GenS1{F0: int32(v1)}
                _ = v18
            }
            _ = fi15
        }
        switch fi14 {
        case 0:
            fmt.Println(v3, v3)
        case -843:
            if (65535 < 790) {
                fmt.Println(v3)
            }
        }
        _ = fi14
    }
    if (0 >= -1) {
        v4 -= (v3 + genFn0(629, true, 255))
        v2 = (v2 & v2)
    } else {
        v19 := GenS3{F0: v2, F1: -1}
        var v20 int64 = 900
        _ = v19
        _ = v20
    }
    v21 := GenS3{F0: (v2 * v2), F1: 1}
    var v22 uint64 = 172
    v23 := []int{4294967295, 9223372036854775807, 0}
    var v24 float64 = 3.14159
    switch "" {
    case "":
        var v25 int8 = -69
        _ = v25
    case "a":
        if (42 <= 54) {
            v2 += 127
            v1 -= v1
        } else {
            var v26 uint8 = (v2 << ((uint(v22)) & 63))
            _ = v26
        }
    }
    v22 = (v22 ^ v22)
    if ((42 < (^(v22))) || true) {
        var v27 bool = (!(!(!(((1 < v3) || (int32(v2) >= 441)) && false))))
        v28 := GenS3{F0: v2, F1: -2147483648}
        v29 := GenS2{F0: 42, F1: -9223372036854775808}
        _ = v27
        _ = v28
        _ = v29
    }
    if (!true) {
        v23 = append(v23, 0)
        var v30 float64 = (((0.0 - ((v24 - -2.25) - v24)) * (v24 - v24)) / v24)
        var v31 int64 = -1
        _ = v30
        _ = v31
    }
    fmt.Println(v22, v4, v2)
    switch ("" + "") {
    case "\n\t":
        v32 := GenS2{F0: 42, F1: -366}
        _ = v32
    default:
        var v33 uint8 = v2
        _ = v33
    }
    v2 = v2
    var v34 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
