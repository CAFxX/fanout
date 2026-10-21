package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 uint = 42
    var v4 uint64 = 919
    var v5 uint64 = 18446744073709551615
    v5 = 1
    v6 := []int{531, 32768, -2147483648}
    v6 = append(v6, 1)
    v3 = (^(v3))
    var v7 float64 = (0.0 + (float64(0.0) * (-2.25 - float64((3.14159 + ((0.1 - ((-2.25 - 1.5) / ((float32(float64(1.5)) * 3.14159) / float32((-2.25 * 0.1))))) + ((1.5 + (((1.5 - 0.0) + -2.25) - 0.0)) - ((1e10 / 1.0) / ((1e10 - ((3.14159 + 1e10) + 1e10)) - (1.5 / 1e10))))))))))
    v3 = v3
    var v8 int32 = -849
    if len(v6) > 0 {
        v6[0] = 4294967295
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS1{F0: 378}
        v8 = (v8 ^ -2147483648)
        _ = fi9
        _ = v10
    }
    _ = len(v1)
    fmt.Println(v3, v3, v5)
    delete(v1, 1)
    v11 := GenS2{F0: 193, F1: "hello"}
    switch 127 {
    case -9223372036854775808:
        if (!false) {
            var v12 int = 42
            v8 = (v8 & -1)
            v13 := GenS1{F0: 2}
            _ = v12
            _ = v13
        } else {
            var v14 float32 = -2.25
            fmt.Println(v3, v3, v7)
            _ = v14
        }
    case -2147483648:
        if true {
            var v15 int64 = int64(v5)
            for fi16 := 0; fi16 < 16; fi16++ {
                v5 = v5
                var v17 int8 = 91
                v18 := GenS1{F0: 733}
                _ = fi16
                _ = v17
                _ = v18
            }
            v19 := GenS1{F0: 32767}
            _ = v15
            _ = v19
        } else {
            var v20 int = 2147483647
            _ = v20
        }
    default:
        var v21 uint = v3
        _ = v21
    }
    var v22 int = 788
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
