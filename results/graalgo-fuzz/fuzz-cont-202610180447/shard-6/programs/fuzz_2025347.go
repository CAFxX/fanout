package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int64; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint16 }

func main() {
    var v1 uint16 = 65535
    var v2 float32 = ((0.1 + float32(-2.25)) / ((-2.25 - (3.14159 * 3.14159)) + 1.5))
    v3 := make(map[int]uint32)
    var v4 string = "\n\t"
    var v5 uint = 667
    v6 := GenS1{F0: 127}
    v5 = v5
    v5 += (v5 * v5)
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 int16 = -1
            if false {
                switch v4 {
                case "\n\t":
                    v1 = (v1 ^ 32767)
                }
                v5 -= v5
            }
            _ = fi8
            _ = v9
        }
        var v10 bool = true
        _ = fi7
        _ = v10
    }
    switch "" {
    case "hello":
        var v11 bool = (4294967295 != 2)
        _ = v11
    case "a":
        v4 = ""
    default:
        if false {
            v5 -= (v5 + (v5 >> (((v5 * (^(v5)))) & 63)))
        }
    }
    v12 := GenS1{F0: 144}
    var v13 float32 = v2
    v14 := GenS2{F0: 9223372036854775807, F1: 42, F2: 0}
    if false {
        var v15 uint8 = 42
        var v16 string = v4
        _ = v15
        _ = v16
    } else {
        var v17 bool = false
        for fi18 := 0; fi18 < 3; fi18++ {
            fmt.Println(v4, v2, v13)
            _ = fi18
        }
        _ = v17
    }
    v19 := v3[-155]
    var v20 int32 = -190
    v21 := []int{-1, int(v20), 80}
    for fi22 := 0; fi22 < 10; fi22++ {
        v5 = 370
        var v23 int32 = v20
        for fi24 := 0; fi24 < 8; fi24++ {
            var v25 int8 = 15
            _ = fi24
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        for fi27 := 0; fi27 < 17; fi27++ {
            fmt.Println(v2)
            _ = fi27
        }
        v28 := GenS2{F0: -741, F1: 601, F2: 42}
        var v29 int64 = -9223372036854775808
        _ = fi26
        _ = v28
        _ = v29
    }
    if (!false) {
        var v30 float32 = (1.5 / v13)
        for fi31 := 0; fi31 < 13; fi31++ {
            v2 = (v13 + v2)
            _ = fi31
        }
        _ = v30
    } else {
        v32 := GenS3{F0: 42, F1: (v1 << ((v5) & 63))}
        _ = v32
    }
    v13 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
