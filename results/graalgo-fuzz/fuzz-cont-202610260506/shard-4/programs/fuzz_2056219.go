package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint16, p1 int8) int64 {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint32 = 65535
        if false {
            fmt.Println(fi7, v8)
        }
        _ = fi7
        _ = v8
    }
    var v9 int8 = (-(p1))
    var v10 int = 917
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 int64 = 430
    v2 := make(map[int]uint32)
    var v3 uint32 = 165
    var v4 int32 = 42
    var v5 uint64 = 71
    v6 := []int{}
    _ = len(v6)
    fmt.Println(v5)
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 float32 = (1e10 + (-2.25 * (0.1 - -2.25)))
        var v13 uint32 = (v3 + 380)
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            v4 = 0
            switch (fi15 * 2147483647) {
            case 9223372036854775807:
                var v16 bool = (2 < 1)
                _ = v16
            case 1:
                v3 = (v3 >> ((2) & 63))
            default:
                v1 = (v1 >> ((255) & 63))
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v3 = v3
            _ = fi15
        }
        var v17 float64 = 1e10
        _ = fi14
        _ = v17
    }
    if false {
        v1 += genFn0(42, -9)
        v18 := GenS1{F0: v4, F1: ""}
        _ = v18
    } else {
        switch "foo bar" {
        case "\n\t":
            if (0 > 100) {
                var v19 uint16 = 65535
                _ = v19
            }
        default:
            fmt.Println(v5, v3)
        }
    }
    v20 := GenS2{F0: -71}
    v21 := GenS2{F0: 1}
    switch 384 {
    case -9223372036854775808:
        var v22 int8 = 41
        _ = v22
    }
    v23 := GenS3{F0: 1000}
    if true {
        fmt.Println(v3, v1)
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 uint16 = 255
            _ = fi24
            _ = v25
        }
        v3 -= 246
    } else {
        var v26 int = 371
        v1 -= v1
        _ = v26
    }
    v27 := GenS2{F0: 99}
    var v28 uint8 = 232
    var v29 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
