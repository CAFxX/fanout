package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint8 }
type GenS2 struct { F0 int; F1 uint8; F2 bool }
type GenS3 struct { F0 int; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := []int{4294967295, 799, -9223372036854775808, 0}
    var v2 int = -1
    var v3 int64 = -644
    v4 := make(map[string]uint32)
    v5 := make(map[int]bool)
    var v6 float64 = 0.0
    v6 = v6
    v7 := GenS1{F0: -175, F1: v2, F2: 255}
    for fi8 := 0; fi8 < 2; fi8++ {
        v6 = (float64(0.0) * v6)
        switch "foo bar" {
        case "":
            if (true || true) {
                v9 := GenS2{F0: (v2 | 2147483647), F1: 128, F2: false}
                var v10 string = ("a" + "hello")
                for fi11 := 0; fi11 < 11; fi11++ {
                    v10 = (v10 + "foo bar")
                    _ = fi11
                }
                _ = v10
                _ = v9
            }
        default:
            _ = len(v1)
        }
        fmt.Println(v6, v2, fi8)
        _ = fi8
    }
    fmt.Println(v3)
    _ = len(v1)
    v12 := GenS3{F0: (v2 | (v2 & 124)), F1: 0, F2: 1}
    v2 = (-(v2))
    var v13 bool = (((3.14159 + 0.1) >= float32((v6 - (float64(0.1) / float64(3.14159))))) || (342 != 18446744073709551615))
    v14 := v5[v2]
    for fi15 := 0; fi15 < 3; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            v3 -= (v3 | v3)
            v17 := GenS4{F0: 42, F1: 104, F2: 32767}
            v18 := GenS5{F0: 2}
            _ = fi16
            _ = v17
            _ = v18
        }
        fmt.Println(v2, fi15, fi15)
        v3 = v3
        _ = fi15
    }
    var v19 float32 = 1.5
    if v14 {
        var v20 uint = 1
        var v21 uint8 = 0
        fmt.Println(v14)
        _ = v20
        _ = v21
    } else {
        fmt.Println(v13)
        switch "" {
        case "foo bar":
            var v22 int = (v2 / ((v2) | 1))
            _ = v22
        default:
            for fi23 := 0; fi23 < 7; fi23++ {
                var v24 float32 = 3.14159
                v24 = (0.0 * -2.25)
                _ = fi23
                _ = v24
            }
        }
    }
    var v25 int32 = -1
    fmt.Println(v14)
    v19 -= v19
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
