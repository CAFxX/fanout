package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int }
func genFn0(p0 uint8) uint8 {
    p0 = (^(p0))
    for fi7 := 0; fi7 < 9; fi7++ {
        switch fi7 {
        case -657:
            v8 := GenS1{F0: 0, F1: -9223372036854775808, F2: fi7}
            _ = v8
        case 834:
            var v9 uint32 = 2147483648
            _ = v9
        case 640:
            p0 = (p0 + (p0 - 128))
        }
        _ = fi7
    }
    var v10 float64 = (1.5 / (3.14159 * -2.25))
    switch "foo bar" {
    case "\n\t":
        var v11 uint64 = 2
        _ = v11
    case "a":
        fmt.Println(v10)
    }
    _ = p0
    _ = v10
    return p0
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    var v4 int32 = 587
    v5 := []int{32768, -2147483648, -55}
    var v6 uint8 = 145
    v1 = (v1 + (v1 + (v1 + "")))
    v6 = 0
    fmt.Println(v1, v6)
    v12 := GenS1{F0: 0, F1: -9223372036854775808, F2: 903}
    v6 += (v6 % ((v6) | 1))
    if (!((0 <= 386) && true)) {
        v4 = v4
        _ = len(v2)
        if (true && false) {
            v3[65535] = -9223372036854775808
        }
    } else {
        switch v1 {
        case "foo bar":
            if (293 == 42) {
                v4 = (v4 & v4)
            }
        case "":
            var v13 float64 = 0.1
            _ = v13
        case "\n\t":
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 float64 = ((1e10 - (0.1 + -2.25)) / 1.0)
                var v16 uint = 468
                var v17 uint16 = uint16(v16)
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        default:
            v1 = v1
        }
        var v18 int32 = v4
        var v19 uint64 = 4294967295
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: 2, F1: 9223372036854775807, F2: 394}
    v6 -= v6
    var v21 int64 = -975
    v22 := GenS1{F0: 0, F1: (v21 - (v21 * v21)), F2: 600}
    v23 := GenS1{F0: 862, F1: (v21 % (((v21 * (v21 | -1))) | 1)), F2: 1}
    for fi24 := 0; fi24 < 8; fi24++ {
        fmt.Println(v4)
        _ = fi24
    }
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    var v25 int8 = 100
    fmt.Println(v1, v4)
    v26 := v2["hello"]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
