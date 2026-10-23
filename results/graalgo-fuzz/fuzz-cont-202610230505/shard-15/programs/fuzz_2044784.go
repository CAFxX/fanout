package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint16; F1 bool; F2 int32 }
func genFn0(p0 int, p1 int8) string {
    p1 = -1
    var v5 uint16 = 899
    if ((p0 * p0) < -1) {
        if (1 > 42) {
            var v6 uint8 = 255
            var v7 uint64 = 2
            _ = v6
            _ = v7
        }
        var v8 int = p0
        var v9 int = p0
        _ = v8
        _ = v9
    } else {
        var v10 int16 = 1
        var v11 int16 = 193
        _ = v10
        _ = v11
    }
    v12 := []int{789, int(v5), -488}
    _ = p0
    _ = p1
    _ = v12
    _ = v5
    return "foo bar"
}

func main() {
    var v1 bool = ((0 <= 18446744073709551615) || false)
    v2 := make(map[int]string)
    var v3 uint64 = 915
    var v4 int8 = -128
    var v13 int16 = 911
    if v1 {
        switch 32767 {
        case 0:
            v4 += 39
        case 937:
            var v14 string = "foo bar"
            _ = v14
        case 4294967295:
            var v15 int32 = 2147483647
            _ = v15
        default:
            v1 = ((!((1.5 + 3.14159) < 0.0)) || ((0 >= 485) && ((v4 + (^(v4))) >= 119)))
        }
        var v16 int64 = 366
        _ = v16
    }
    v3 -= v3
    fmt.Println(v13, v4)
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 1; fi18++ {
            fmt.Println(fi18, fi18)
            _ = fi18
        }
        _ = fi17
    }
    v13 = (v13 % ((-1) | 1))
    if false {
        v1 = v1
        var v19 int32 = 2147483647
        _ = v19
    }
    fmt.Println(v4, v1)
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 uint8 = 209
        v22 := GenS1{F0: (^(fi20)), F1: ""}
        _ = fi20
        _ = v21
        _ = v22
    }
    v23 := GenS2{F0: 222, F1: (((1e10 + 0.0) == 1e10) && true), F2: 1}
    if v1 {
        for fi24 := 0; fi24 < 16; fi24++ {
            switch "foo bar" {
            case "\n\t":
                var v25 uint = 687
                _ = v25
            case "hello":
                v26 := []int{32767}
                _ = v26
            case "foo bar":
                fmt.Println(v13, v13)
            default:
                fmt.Println(fi24)
            }
            _ = fi24
        }
        v27 := GenS2{F0: 1, F1: v1, F2: 490}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
