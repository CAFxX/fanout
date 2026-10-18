package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int16; F2 uint }

func main() {
    var v1 float64 = 0.1
    v2 := []string{}
    v3 := make(map[string]uint32)
    var v4 float32 = 1.5
    var v5 uint64 = 950
    var v6 int64 = 1
    var v7 int16 = -511
    var v8 float64 = float64(float32(1.5))
    _ = len(v3)
    var v9 uint8 = 193
    var v10 int = -2147483648
    for fi11 := 0; fi11 < 14; fi11++ {
        v5 -= v5
        v12 := GenS1{F0: v5, F1: 2}
        v13 := GenS1{F0: (v5 ^ (^(v5))), F1: uint32(v9)}
        _ = fi11
        _ = v12
        _ = v13
    }
    v3["a"] = uint32(v5)
    v14 := GenS1{F0: (v5 - (v5 >> ((382) & 63))), F1: 559}
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            switch "a" {
            case "foo bar":
                var v17 int64 = int64(v5)
                _ = v17
            case "a":
                v5 -= v5
            case "hello":
                v18 := GenS1{F0: 42, F1: 139}
                _ = v18
            }
            v7 -= 644
            _ = fi16
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi15
    }
    var v19 uint8 = (v9 * v9)
    v20 := GenS2{F0: (v10 * (v10 & 42))}
    fmt.Println(v5, v4)
    v4 = (v4 * (v4 / v4))
    v7 -= v7
    v21 := GenS3{F0: 1, F1: 1000, F2: 255}
    if ((v1 == v8) && false) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v22 bool = (!false)
        v3["\n\t"] = 65535
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
