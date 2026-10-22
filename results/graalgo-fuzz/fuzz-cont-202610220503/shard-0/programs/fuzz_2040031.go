package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }

func main() {
    v1 := []string{"a", "\n\t", "\n\t"}
    var v2 float32 = ((1e10 * 0.0) / 1e10)
    v3 := []string{"a", "a"}
    var v4 int8 = 42
    var v5 int64 = 0
    v6 := make(map[uint32]uint32)
    v5 = 0
    fmt.Println(v2, v2, v5)
    var v7 bool = ((18446744073709551615 == 547) || ((float32(float64(v2)) + 1.5) != v2))
    fmt.Println(v2, v2)
    v4 = (^(v4))
    if ((v5 / (((v5 | v5)) | 1)) != -9223372036854775808) {
        var v8 int = 1
        fmt.Println(v8, v5)
        for fi9 := 0; fi9 < 11; fi9++ {
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 uint64 = 42
                if v7 {
                    v3 = append(v3, "")
                    v12 := GenS1{F0: 42, F1: 0}
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
            v13 := GenS1{F0: -599, F1: 42}
            _ = fi9
            _ = v13
        }
        _ = v8
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        v4 = v4
        v7 = ((v7 && (22 <= 2147483648)) || (!(!true)))
        _ = fi14
    }
    var v15 uint32 = 65535
    v15 = (v15 & v15)
    fmt.Println(v4)
    v4 -= (v4 * (v4 ^ (v4 & (v4 % (((v4 / ((-128) | 1))) | 1)))))
    var v16 uint8 = 128
    switch 1 {
    case 42:
        var v17 float32 = 1.5
        _ = v17
    default:
        var v18 int8 = v4
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
