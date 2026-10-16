package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int16; F1 int }

func main() {
    var v1 float32 = 1.5
    var v2 uint8 = 1
    v3 := make(map[uint32]uint32)
    v4 := []uint32{2147483647, 2147483647}
    v5 := make(map[uint32]string)
    var v6 uint64 = 1
    var v7 int32 = -1
    v2 = (v2 / (((v2 * (v2 + v2))) | 1))
    for fi8 := 0; fi8 < 13; fi8++ {
        switch ("" + "") {
        case "a":
            v6 -= 4294967295
        case "hello":
            if true {
                for fi9 := 0; fi9 < 11; fi9++ {
                    switch 747 {
                    case 32768:
                        var v10 uint8 = (^(v2))
                        _ = v10
                    case 73:
                        v6 += v6
                    }
                    v11 := GenS1{F0: (v7 & (v7 - (v7 ^ (v7 ^ -311)))), F1: (^(fi8)), F2: 42}
                    _ = fi9
                    _ = v11
                }
                v1 -= -2.25
                var v12 uint8 = v2
                _ = v12
            }
        default:
            var v13 float32 = (v1 - ((v1 / ((1.5 - 0.1) + 0.1)) / (0.1 * v1)))
            _ = v13
        }
        if (1 != -1) {
            v14 := v5[2147483648]
            var v15 int8 = -1
            _ = v14
            _ = v15
        }
        _ = fi8
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        var v16 int = 42
        _ = v16
    case 431:
        v17 := GenS1{F0: 42, F1: 184, F2: 1}
        _ = v17
    }
    var v18 uint16 = 2
    v6 += v6
    v19 := GenS2{F0: int8(v2), F1: -960}
    v20 := GenS3{F0: 1, F1: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
