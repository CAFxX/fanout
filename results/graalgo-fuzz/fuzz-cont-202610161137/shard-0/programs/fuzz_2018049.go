package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int16; F1 int64 }
func genFn0(p0 uint64, p1 uint16) int16 {
    switch ("hello" + "\n\t") {
    case "\n\t":
        p0 = p0
    case "":
        var v6 string = "foo bar"
        _ = v6
    case "foo bar":
        var v7 int64 = 678
        _ = v7
    }
    if (!false) {
        p0 = (p0 >> ((578) & 63))
        var v8 uint8 = 76
        fmt.Println(v8, p0)
        _ = v8
    }
    if (2 < 2) {
        v9 := GenS1{F0: 127, F1: 42, F2: 190}
        _ = v9
    }
    var v10 uint8 = 153
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 uint16 = 65535
        fmt.Println(p0)
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return 0
}

func main() {
    var v1 float64 = (0.1 + 0.1)
    var v2 uint8 = 1
    var v3 bool = (32768 > -205)
    var v4 float64 = v1
    var v5 uint = 1
    v2 = (v2 - v2)
    v2 += v2
    fmt.Println(v2, v5)
    if v3 {
        var v13 float64 = 0.0
        _ = v13
    }
    var v14 int64 = 1
    if (v3 && v3) {
        switch "a" {
        case "\n\t":
            switch 326 {
            case -9223372036854775808:
                if (genFn0(42, uint16(v2)) >= 1000) {
                    var v15 uint64 = 279
                    var v16 int16 = 1
                    var v17 string = ""
                    _ = v15
                    _ = v16
                    _ = v17
                }
            default:
                var v18 float32 = (-2.25 - 1e10)
                _ = v18
            }
        }
    }
    v1 += float64((((3.14159 * 0.0) + (0.0 * 1e10)) + 0.0))
    for fi19 := 0; fi19 < 8; fi19++ {
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := GenS1{F0: (v2 << ((4294967295) & 63)), F1: 1, F2: 42}
            var v22 uint = (^(v5))
            var v23 string = "\n\t"
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 bool = v3
            _ = fi24
            _ = v25
        }
        v14 += v14
        _ = fi19
    }
    var v26 float64 = v1
    v27 := []int{9223372036854775807, -534, -100}
    v28 := GenS1{F0: 0, F1: 4294967295, F2: v2}
    fmt.Println(v3, v26)
    fmt.Println(v1, v4, v3)
    v29 := GenS2{F0: -306, F1: (v14 * (v14 / ((1) | 1)))}
    if len(v27) > 0 {
        v27[0] = -9223372036854775808
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
