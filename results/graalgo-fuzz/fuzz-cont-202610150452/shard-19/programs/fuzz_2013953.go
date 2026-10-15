package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint = 255
    var v3 int = 9223372036854775807
    v4 := []int{v3, (-(v3)), v3, 249}
    v5 := make(map[uint32]uint32)
    var v6 int32 = -2147483648
    v4 = append(v4, v3)
    if (false || true) {
        v1 = (v1 / ((555) | 1))
        v7 := GenS1{F0: (^(v1)), F1: 2147483647}
        v8 := GenS1{F0: v1, F1: 857}
        _ = v7
        _ = v8
    }
    var v9 int = (v3 + v3)
    var v10 bool = false
    v11 := GenS2{F0: 147, F1: 128, F2: 4294967295}
    for fi12 := 0; fi12 < 15; fi12++ {
        v9 = (^(fi12))
        fmt.Println(v10)
        switch "" {
        case "hello":
            v13 := GenS3{F0: 0}
            _ = v13
        }
        _ = fi12
    }
    v10 = v10
    var v14 uint8 = 55
    fmt.Println(v14)
    v15 := GenS3{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
