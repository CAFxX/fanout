package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint32 }
type GenS2 struct { F0 int }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 2; fi6++ {
        if ((!(83 >= 18446744073709551615)) || (2 <= 18446744073709551615)) {
            fmt.Println(fi6)
            if (!false) {
                v7 := []int{}
                var v8 uint8 = 255
                _ = v7
                _ = v8
            } else {
                v9 := []int{(fi6 - (fi6 * fi6)), 246}
                var v10 int8 = -128
                _ = v10
                _ = v9
            }
            switch ("a" + ("foo bar" + "hello")) {
            case "foo bar":
                var v11 int32 = 2147483647
                _ = v11
            default:
                var v12 float32 = (0.0 - ((0.0 / 1.0) * 0.0))
                _ = v12
            }
        } else {
            v13 := GenS1{F0: 1, F1: -662, F2: 1}
            v14 := []int{2147483647, (fi6 >> ((4294967295) & 63))}
            if (-936 > 42) {
                fmt.Println(fi6)
            }
            _ = v13
            _ = v14
        }
        v15 := GenS1{F0: -162, F1: 42, F2: 0}
        _ = fi6
        _ = v15
    }
    v16 := GenS1{F0: 1, F1: 34, F2: 2}
    _ = v16
    return 1
}

func main() {
    var v1 string = "hello"
    var v2 int16 = -32768
    var v3 float32 = -2.25
    var v4 int8 = 1
    v5 := make(map[uint32]uint32)
    if (1 < 570) {
        v17 := v5[184]
        _ = v17
    } else {
        switch "a" {
        case "hello":
            var v18 string = ("\n\t" + (v1 + v1))
            _ = v18
        case "":
            v1 = v1
        case "\n\t":
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v4)
            case "":
                v19 := []int{127}
                _ = v19
            default:
                var v20 int16 = 0
                _ = v20
            }
        default:
            v21 := GenS1{F0: 0, F1: (v2 ^ v2), F2: 189}
            _ = v21
        }
        if false {
            v4 -= (v4 / ((v4) | 1))
            v3 = 0.0
        }
    }
    v22 := GenS2{F0: 719}
    var v23 uint16 = 454
    if true {
        v24 := []int{42, 4294967295}
        v2 -= v2
        v25 := GenS1{F0: 238, F1: (v2 - v2), F2: 950}
        _ = v24
        _ = v25
    }
    var v26 bool = true
    if false {
        var v27 bool = true
        _ = v27
    } else {
        fmt.Println(v3, v1)
        v4 = -96
    }
    v28 := GenS1{F0: 480, F1: v2, F2: 4294967295}
    fmt.Println(v23, v23)
    var v29 int = -1
    v30 := GenS2{F0: v29}
    var v31 float64 = (-2.25 * (-2.25 * 1.5))
    for fi32 := 0; fi32 < 8; fi32++ {
        if v26 {
            var v33 float32 = v3
            _ = v33
        }
        _ = fi32
    }
    if v26 {
        var v34 bool = (v1 == "")
        switch "hello" {
        case "foo bar":
            fmt.Println(v31, v34, v29)
        case "":
            v23 = v23
        case "\n\t":
            v23 = v23
        }
        fmt.Println(v3, v31, v31)
        _ = v34
    }
    v35 := GenS1{F0: 448, F1: (v2 << ((255) & 63)), F2: 2147483647}
    var v36 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
