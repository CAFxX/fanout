package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int }
type GenS2 struct { F0 int16; F1 uint; F2 string }
func genFn0() uint16 {
    var v6 int32 = -349
    if (!false) {
        v6 -= (v6 / ((v6) | 1))
        v7 := []int{-277, 138, 2147483647, -2147483648}
        _ = v7
    }
    _ = v6
    return 32767
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    var v4 float64 = 0.0
    var v5 int16 = 42
    v8 := GenS1{F0: 476, F1: 42, F2: 9223372036854775807}
    v1 -= 0
    var v9 int32 = -2147483648
    v4 -= 1e10
    var v10 uint32 = (v1 / ((4294967295) | 1))
    _ = len(v3)
    if false {
        switch -465 {
        case 32768:
            v4 = (v4 * (1.5 + ((float64(3.14159) * (((v4 - v4) * (v4 * v4)) * 0.0)) + 0.0)))
        case 128:
            fmt.Println(v10, v1, v9)
        case -1:
            v1 = 0
        }
        if (!(("\n\t" + "") < "hello")) {
            var v11 int32 = (v9 ^ (v9 << ((2) & 63)))
            if (1 < uint(v10)) {
                delete(v3, 429)
            } else {
                if (!(v4 != v4)) {
                    v12 := v3[-1]
                    _ = v12
                }
                fmt.Println(v10)
            }
            v13 := v2[32768]
            _ = v11
            _ = v13
        } else {
            var v14 int16 = (v5 + 42)
            for fi15 := 0; fi15 < 14; fi15++ {
                v1 = 42
                fmt.Println(v4, v1, v9)
                _ = fi15
            }
            _ = v14
        }
        switch "hello" {
        case "a":
            if true {
                fmt.Println(v1)
                var v16 uint = 127
                v5 += 1
                _ = v16
            } else {
                v17 := GenS2{F0: v5, F1: 18446744073709551615, F2: "\n\t"}
                _ = v17
            }
        case "\n\t":
            v10 = 349
        default:
            if true {
                if (!true) {
                    var v18 string = "foo bar"
                    var v19 float64 = 3.14159
                    var v20 bool = true
                    _ = v18
                    _ = v19
                    _ = v20
                }
                switch "\n\t" {
                case "a":
                    var v21 uint = 1
                    _ = v21
                case "":
                    var v22 float64 = (1.5 / v4)
                    _ = v22
                case "\n\t":
                    v10 = v1
                }
            } else {
                v5 -= (^(v5))
                var v23 int8 = -118
                _ = v23
            }
        }
    } else {
        v4 += 0.1
        var v24 bool = true
        _ = v24
    }
    var v25 int64 = -420
    v26 := GenS2{F0: (-(v5)), F1: 121, F2: "hello"}
    v1 = (v1 | (^(v10)))
    _ = len(v3)
    var v27 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
