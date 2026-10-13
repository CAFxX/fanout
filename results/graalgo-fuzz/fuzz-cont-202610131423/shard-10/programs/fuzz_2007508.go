package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }

func main() {
    var v1 uint8 = 120
    var v2 uint = 2
    var v3 int = 321
    v4 := make(map[int]bool)
    var v5 string = ""
    var v6 uint = 214
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint32 = 134
        var v9 float64 = (-2.25 - (3.14159 * 0.0))
        _ = fi7
        _ = v8
        _ = v9
    }
    v1 = 59
    if (!false) {
        v5 = v5
        v10 := GenS1{F0: -2147483648, F1: 14}
        _ = v10
    } else {
        switch v3 {
        case 778:
            v4[(^(v3))] = false
        case -813:
            if (2 == (v2 ^ 18446744073709551615)) {
                var v11 int = (v3 + v3)
                var v12 int64 = -200
                v13 := GenS1{F0: 1, F1: -67}
                _ = v11
                _ = v12
                _ = v13
            }
        default:
            var v14 int8 = 42
            _ = v14
        }
        v1 = 128
        if true {
            switch "foo bar" {
            case "a":
                v15 := v4[v3]
                _ = v15
            case "hello":
                v4[v3] = true
            }
            var v16 string = ("" + v5)
            v17 := []int{127, (v3 | 127), 191, v3}
            _ = v16
            _ = v17
        }
    }
    v18 := v4[(v3 * v3)]
    if (v3 != -221) {
        var v19 uint64 = 2
        v5 = v5
        var v20 int64 = 42
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 uint = v6
            var v24 int32 = -1
            var v25 uint16 = 255
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        var v26 uint32 = 4294967295
        v3 += v3
        _ = fi21
        _ = v26
    }
    var v27 uint8 = 0
    v28 := []int{v3, -1, (v3 ^ v3)}
    if v18 {
        var v29 int64 = -9223372036854775808
        _ = v29
    }
    v2 += 880
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
