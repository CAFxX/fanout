package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 uint32 = 42
    v2 := []string{}
    var v3 float32 = -2.25
    _ = len(v2)
    if (false || false) {
        var v4 int = 9223372036854775807
        if (false && (1000 != -11)) {
            _ = len(v2)
            fmt.Println(v1)
            var v5 int64 = 1
            _ = v5
        }
        if true {
            if false {
                v6 := GenS1{F0: 2, F1: v1}
                fmt.Println(v4, v1, v4)
                _ = v6
            }
            var v7 uint = 4294967295
            _ = v7
        }
        _ = v4
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(v1, v1)
        if (0.0 != ((1e10 * -2.25) - float64((v3 - (-2.25 * (-2.25 + v3)))))) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if true {
                v3 += v3
                var v9 string = "hello"
                _ = v9
            }
            var v10 uint16 = 42
            _ = v10
        }
        fmt.Println(v3, v1, v1)
        _ = fi8
    }
    fmt.Println(v1, v3)
    var v11 int16 = -1
    if ((641 < (v11 - 0)) && false) {
        v11 -= 42
        if false {
            var v12 uint64 = 157
            var v13 uint8 = 242
            var v14 uint32 = (v1 / ((v1) | 1))
            _ = v12
            _ = v13
            _ = v14
        }
        var v15 int32 = 520
        _ = v15
    }
    _ = len(v2)
    fmt.Println(v3)
    var v16 uint16 = 0
    v16 += v16
    v1 -= (v1 * v1)
    v17 := GenS1{F0: v16, F1: (^(v1))}
    fmt.Println(v16)
    var v18 int8 = 42
    if (0 > (v16 & 32767)) {
        v1 = 42
    } else {
        if ("" <= ("" + "hello")) {
            v1 = (v1 & (v1 - v1))
            var v19 int64 = -1
            _ = v19
        } else {
            v16 -= 32767
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 int64 = 42
                _ = fi20
                _ = v21
            }
            v18 -= -128
        }
        if false {
            switch "" {
            case "":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "a":
                var v22 bool = false
                _ = v22
            case "\n\t":
                _ = len(v2)
            default:
                v1 = 2147483647
            }
        } else {
            if (true && (255 == 127)) {
                fmt.Println(v1, v18)
                v11 += v11
            } else {
                _ = len(v2)
                fmt.Println(v18)
                var v23 int8 = (v18 ^ 77)
                _ = v23
            }
            v2 = append(v2, ("" + ""))
        }
        if (!(!false)) {
            fmt.Println(v18)
            for fi24 := 0; fi24 < 7; fi24++ {
                v25 := GenS1{F0: (v16 / ((v16) | 1)), F1: (v1 << ((0) & 63))}
                var v26 uint = 0
                _ = fi24
                _ = v25
                _ = v26
            }
            v11 = v11
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
