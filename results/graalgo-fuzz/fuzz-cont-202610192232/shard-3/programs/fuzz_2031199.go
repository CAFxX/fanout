package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 bool; F1 int8; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint64; F2 string }
func genFn0(p0 int, p1 int8, p2 uint8) int8 {
    if false {
        var v7 float64 = -2.25
        _ = v7
    }
    if true {
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 int64 = 42
    var v2 uint16 = 653
    var v3 int8 = 47
    var v4 uint8 = 128
    var v5 uint8 = v4
    var v6 int16 = 32767
    var v8 uint = 255
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := []int{}
        var v11 int16 = v6
        var v12 int32 = 1
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: 42, F1: 42}
    fmt.Println(v5, v1, v6)
    if (!((v2 | (v2 << ((v8) & 63))) == uint16(v5))) {
        if (!false) {
            switch -9223372036854775808 {
            case -591:
                for fi14 := 0; fi14 < 18; fi14++ {
                    if (int8(fi14) <= (v3 >> ((v8) & 63))) {
                        v6 = 1
                    } else {
                        var v15 uint64 = 192
                        _ = v15
                    }
                    switch fi14 {
                    case 127:
                        v16 := GenS2{F0: false, F1: (v3 / ((1) | 1)), F2: v2}
                        _ = v16
                    case -836:
                        v17 := GenS2{F0: true, F1: genFn0(fi14, 98, v5), F2: (^(v2))}
                        _ = v17
                    default:
                        fmt.Println(v5, v6, v3)
                    }
                    switch "" {
                    case "foo bar":
                        v18 := []int{(fi14 & fi14)}
                        _ = v18
                    case "\n\t":
                        v5 -= (v5 * v5)
                    }
                    _ = fi14
                }
            case -295:
                v19 := GenS3{F0: v1, F1: 18446744073709551615, F2: ""}
                _ = v19
            default:
                var v20 uint64 = 1
                _ = v20
            }
            var v21 int16 = v6
            _ = v21
        }
        v22 := GenS1{F0: (v8 - v8), F1: uint64(v2)}
        _ = v22
    }
    var v23 int64 = v1
    if true {
        v24 := []int{}
        v25 := GenS3{F0: (v1 - (v23 % ((v1) | 1))), F1: 4294967295, F2: "a"}
        for fi26 := 0; fi26 < 8; fi26++ {
            v5 += v5
            _ = fi26
        }
        _ = v24
        _ = v25
    }
    v27 := []int{}
    for fi28 := 0; fi28 < 17; fi28++ {
        if ((false && false) && true) {
            switch ("hello" + "a") {
            case "":
                fmt.Println(v1, v3, v8)
            }
        }
        switch "" {
        case "a":
            fmt.Println(v4)
        }
        v1 += v1
        _ = fi28
    }
    if len(v27) > 0 {
        v27[0] = int(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
