package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 float32 = (1.5 - (1e10 / (1e10 * 1e10)))
    v2 := []int{1, 317}
    var v3 uint16 = 2
    v4 := make(map[int]string)
    var v5 uint32 = 273
    var v6 int32 = 614
    v7 := GenS1{F0: 422, F1: (v5 ^ v5)}
    var v8 uint = 2
    v3 = (v3 >> ((1) & 63))
    fmt.Println(v3)
    v2 = append(v2, 128)
    var v9 int64 = -9223372036854775808
    v10 := v4[9223372036854775807]
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            switch "foo bar" {
            case "hello":
                var v13 int8 = 50
                _ = v13
            case "a":
                v3 = v3
            default:
                v14 := GenS1{F0: 170, F1: v5}
                _ = v14
            }
            var v15 uint8 = 1
            v16 := GenS1{F0: uint64(v3), F1: 139}
            _ = fi12
            _ = v15
            _ = v16
        }
        _ = fi11
    }
    v17 := GenS1{F0: 42, F1: v5}
    var v18 uint32 = (v5 | (v5 + v5))
    if (!false) {
        fmt.Println(v18, v3, v8)
        var v19 float64 = float64(1.5)
        _ = v19
    } else {
        var v20 uint32 = (v18 & 2147483648)
        if ((v5 & 339) > 191) {
            v10 = ("\n\t" + v10)
            var v21 uint = (v8 & (v8 % (((v8 * v8)) | 1)))
            v3 = (v3 + (^(v3)))
            _ = v21
        }
        for fi22 := 0; fi22 < 9; fi22++ {
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := v4[fi23]
                _ = fi23
                _ = v24
            }
            v2 = append(v2, 0)
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 uint8 = 2
                _ = len(v2)
                v3 = v3
                _ = fi25
                _ = v26
            }
            _ = fi22
        }
        _ = v20
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        v5 = v5
        v28 := GenS1{F0: 4294967295, F1: 256}
        var v29 uint8 = 42
        _ = fi27
        _ = v28
        _ = v29
    }
    v30 := v4[440]
    var v31 uint16 = (v3 * (^(v3)))
    var v32 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
