package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint64 = 2
    var v3 float64 = 0.1
    var v4 uint32 = 42
    if (!(!false)) {
        var v5 uint = 331
        _ = v5
    }
    v6 := GenS1{F0: 18, F1: 2147483647}
    if true {
        for fi7 := 0; fi7 < 18; fi7++ {
            switch ("" + "foo bar") {
            case "\n\t":
                fmt.Println(fi7)
            }
            v8 := GenS1{F0: 42, F1: 132}
            v9 := GenS2{F0: 1}
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := []int{-344, 395, -184}
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = v10
    }
    v11 := []int{0, 0}
    if true {
        var v12 int32 = -1
        v3 -= ((((float64((0.0 / 1.0)) + (((v3 / 1.5) + v3) - -2.25)) + 1.5) / (v3 * (1.5 * v3))) - (v3 * ((float64(1e10) / float64(0.1)) - 1.5)))
        _ = v12
    }
    var v13 int64 = -1
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    _ = len(v11)
    if (!(true && false)) {
        var v14 uint8 = uint8(v2)
        fmt.Println(v2)
        _ = v14
    } else {
        for fi15 := 0; fi15 < 8; fi15++ {
            v1 += v1
            fmt.Println(v3, fi15, v4)
            var v16 uint = 0
            _ = fi15
            _ = v16
        }
        v11 = append(v11, 127)
    }
    var v17 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
