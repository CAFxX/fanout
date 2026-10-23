package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 bool; F1 int64; F2 int16 }

func main() {
    var v1 int32 = 42
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 uint = 18446744073709551615
    v4 -= v4
    v4 = v4
    v4 = v4
    v4 = v4
    if false {
        v5 := GenS1{F0: 1, F1: 42, F2: 169}
        fmt.Println(v4, v1)
        fmt.Println(v1, v4)
        _ = v5
    }
    if false {
        for fi6 := 0; fi6 < 11; fi6++ {
            switch fi6 {
            case 1:
                var v7 float64 = -2.25
                _ = v7
            case -9223372036854775808:
                switch fi6 {
                case 127:
                    var v8 uint64 = 18446744073709551615
                    _ = v8
                case 0:
                    v9 := GenS1{F0: 0, F1: 834, F2: -990}
                    _ = v9
                case -9223372036854775808:
                    if (true || true) {
                        v10 := GenS2{F0: fi6, F1: 0}
                        _ = v10
                    }
                default:
                    v4 += 546
                }
            }
            _ = fi6
        }
        fmt.Println(v4, v4)
    }
    v1 += int32(v4)
    v11 := GenS3{F0: (2147483647 == -9223372036854775808), F1: 1, F2: 1}
    var v12 uint32 = 0
    v13 := GenS2{F0: 32767, F1: 4294967295}
    fmt.Println(v1)
    switch "a" {
    case "\n\t":
        var v14 uint64 = 2
        _ = v14
    }
    v15 := v2[2147483647]
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
