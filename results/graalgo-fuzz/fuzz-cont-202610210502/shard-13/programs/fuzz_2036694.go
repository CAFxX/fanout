package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int64 }

func main() {
    var v1 float64 = 3.14159
    var v2 uint8 = 1
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]uint32)
    var v5 uint8 = v2
    var v6 uint64 = 18446744073709551615
    fmt.Println(v1, v6)
    if false {
        fmt.Println(v6, v5)
        v7 := GenS1{F0: v6, F1: 1}
        v6 += (v6 % ((0) | 1))
        _ = v7
    }
    _ = len(v4)
    for fi8 := 0; fi8 < 15; fi8++ {
        v2 -= (^(v5))
        if true {
            fmt.Println(v1, v6, v2)
            v3[449] = false
        } else {
            var v9 uint32 = 0
            v5 = (^(v2))
            v4[1] = v9
            _ = v9
        }
        switch ("foo bar" + "a") {
        case "\n\t":
            v10 := GenS1{F0: 4294967295, F1: -2147483648}
            _ = v10
        case "":
            for fi11 := 0; fi11 < 11; fi11++ {
                if (!true) {
                    v12 := GenS2{F0: -743}
                    v13 := GenS3{F0: 104, F1: 423}
                    _ = v12
                    _ = v13
                } else {
                    fmt.Println(fi11, fi11)
                    var v14 float64 = 1e10
                    v15 := GenS4{F0: 9223372036854775807}
                    _ = v14
                    _ = v15
                }
                _ = fi11
            }
        default:
            v4[65535] = 0
        }
        _ = fi8
    }
    v1 -= float64((0.0 * 3.14159))
    v16 := []int{982, -9223372036854775808}
    fmt.Println(v6, v1)
    for fi17 := 0; fi17 < 20; fi17++ {
        v5 = 127
        v3[1] = true
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
