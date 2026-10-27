package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 string }

func main() {
    var v1 int32 = 834
    v2 := make(map[int]bool)
    var v3 uint8 = 42
    var v4 uint8 = 31
    var v5 int8 = 0
    v4 += 236
    var v6 uint8 = (v3 << ((849) & 63))
    v7 := []int{}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 float32 = 0.1
        var v10 float64 = -2.25
        for fi11 := 0; fi11 < 15; fi11++ {
            if (-678 > 42) {
                var v12 int16 = 1000
                _ = v12
            }
            fmt.Println(fi8)
            v1 = (v1 % (((v1 | (v1 / ((-710) | 1)))) | 1))
            _ = fi11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    if len(v7) > 0 {
        v7[0] = -262
    }
    if (-454 >= 101) {
        v13 := GenS1{F0: 519, F1: 65535, F2: "a"}
        if false {
            var v14 uint16 = 2
            v2[-1] = true
            _ = v14
        } else {
            if false {
                v7 = append(v7, 4294967295)
            }
            fmt.Println(v1, v6)
        }
        _ = v13
    } else {
        fmt.Println(v1, v5)
        switch 0 {
        case -2147483648:
            var v15 float32 = -2.25
            _ = v15
        case -1:
            if false {
                v16 := GenS1{F0: 709, F1: 696, F2: ""}
                _ = v16
            }
        }
        v6 = v6
    }
    if ("hello" < "foo bar") {
        v3 = 79
    } else {
        var v17 int64 = 42
        for fi18 := 0; fi18 < 4; fi18++ {
            v4 -= v6
            if true {
                v5 = int8(fi18)
            } else {
                _ = len(v7)
                var v19 int8 = v5
                v3 = 159
                _ = v19
            }
            _ = fi18
        }
        v20 := GenS1{F0: 0, F1: 42, F2: "hello"}
        _ = v17
        _ = v20
    }
    v6 += (v4 ^ (v4 - 100))
    var v21 uint32 = 42
    if false {
        var v22 int64 = 9223372036854775807
        _ = v22
    }
    v5 += 127
    var v23 uint32 = v21
    var v24 uint64 = 42
    if (!false) {
        fmt.Println(v6)
        var v25 float64 = ((-2.25 * 0.0) / 1.5)
        _ = v25
    }
    fmt.Println(v21, v5)
    var v26 uint16 = 0
    for fi27 := 0; fi27 < 7; fi27++ {
        switch "\n\t" {
        case "\n\t":
            var v28 int8 = -85
            _ = v28
        case "a":
            if len(v7) > 0 {
                v7[0] = fi27
            }
        default:
            v29 := v2[(^(fi27))]
            _ = v29
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
