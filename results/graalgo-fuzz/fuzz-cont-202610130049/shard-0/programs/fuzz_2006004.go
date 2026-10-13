package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int16; F1 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 491
    var v3 float32 = 0.0
    var v4 uint = 138
    var v5 uint8 = 2
    v6 := GenS1{F0: 0, F1: (v2 << (((^(v4))) & 63)), F2: (v2 % (((v2 >> ((v4) & 63))) | 1))}
    if (-128 < 1) {
        fmt.Println(v4, v5, v2)
    }
    v5 += v5
    if ((0.1 - 0.1) != float64(v3)) {
        fmt.Println(v4, v5, v4)
    }
    var v7 float64 = 1.5
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            switch "\n\t" {
            case "foo bar":
                var v10 int16 = int16(v5)
                _ = v10
            case "\n\t":
                v1[65535] = ("\n\t" + ("" + "\n\t"))
            default:
                v11 := []int{fi8, 9223372036854775807, (fi8 + (fi9 / ((fi8) | 1))), fi8}
                _ = v11
            }
            v12 := []int{fi8, 2147483647}
            v4 -= 1
            var v13 uint32 = 2147483648
            _ = fi9
            _ = v12
            _ = v13
        }
        var v14 uint16 = 138
        _ = fi8
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        if true {
            var v16 bool = (false || false)
            v17 := GenS2{F0: 42, F1: 501, F2: 42}
            _ = v16
            _ = v17
        }
        var v18 uint = v4
        _ = fi15
        _ = v18
    }
    v2 = v2
    switch 4294967295 {
    case -2147483648:
        v5 = 42
    default:
        var v19 float32 = ((v3 / (1e10 / (v3 * 1e10))) - v3)
        _ = v19
    }
    v7 -= v7
    var v20 uint = v4
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := GenS3{F0: v20, F1: 0}
        var v23 int8 = -61
        for fi24 := 0; fi24 < 1; fi24++ {
            v20 += v4
            fmt.Println(v7, v2)
            var v25 uint16 = 626
            var v26 uint8 = v5
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    var v27 uint64 = (v2 << (((v20 + (v4 | 42))) & 63))
    if false {
        v28 := GenS4{F0: 42, F1: -1}
        v29 := GenS1{F0: (v20 / ((v20) | 1)), F1: (v2 ^ (v27 | v2)), F2: v27}
        v3 -= (((v3 / v3) * (v3 - ((v3 - v3) * v3))) - -2.25)
        _ = v28
        _ = v29
    } else {
        v30 := []int{}
        var v31 int = 9223372036854775807
        _ = v30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
