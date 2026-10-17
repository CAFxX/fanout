package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
func genFn0() int64 {
    var v5 int64 = 904
    var v6 uint = 19
    var v7 uint8 = 128
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 1
    var v3 uint64 = 1
    var v4 float64 = 0.1
    v8 := GenS1{F0: 1, F1: (v2 + (^(v2)))}
    if false {
        var v9 uint32 = 754
        v10 := []int{695, -9223372036854775808, 622, 42}
        _ = v10
        _ = v9
    }
    v4 += (v4 + 1e10)
    v11 := []int{-9223372036854775808, -1, -660}
    for fi12 := 0; fi12 < 6; fi12++ {
        v3 += (v3 ^ (^(v3)))
        _ = fi12
    }
    if ((0 == 2147483647) && (32768 >= -2147483648)) {
        v13 := v1[42]
        switch 127 {
        case 699:
            switch "hello" {
            case "hello":
                fmt.Println(v4)
            }
        default:
            var v14 uint8 = 127
            _ = v14
        }
        _ = v13
    } else {
        if false {
            var v15 int32 = -911
            _ = v15
        }
        _ = len(v11)
        v16 := v1[0]
        _ = v16
    }
    var v17 uint8 = 68
    var v18 int64 = 0
    switch "\n\t" {
    case "\n\t":
        v1[0] = 4294967295
    default:
        var v19 int16 = 32767
        _ = v19
    }
    var v20 uint = 255
    var v21 int = 32767
    v22 := GenS1{F0: 100, F1: 255}
    for fi23 := 0; fi23 < 7; fi23++ {
        fmt.Println(v21, v20, v21)
        _ = fi23
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        switch "hello" {
        case "a":
            var v25 uint8 = 127
            _ = v25
        case "\n\t":
            v26 := GenS1{F0: int8(fi24), F1: v2}
            _ = v26
        default:
            for fi27 := 0; fi27 < 9; fi27++ {
                v4 -= v4
                var v28 uint8 = (v17 & (v17 | v17))
                var v29 int8 = 1
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        var v30 uint = v20
        _ = fi24
        _ = v30
    }
    v20 = 65535
    if false {
        v1[65535] = -607
    }
    var v31 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v4
    _ = v8
}
