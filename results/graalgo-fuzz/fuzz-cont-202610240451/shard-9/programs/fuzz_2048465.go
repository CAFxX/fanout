package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS3 struct { F0 int; F1 uint32; F2 int8 }
func genFn0() int16 {
    var v6 uint64 = 18446744073709551615
    v6 = (v6 % ((v6) | 1))
    switch -701 {
    case 9223372036854775807:
        v7 := GenS1{F0: 671, F1: 42}
        _ = v7
    case 1:
        v8 := []int{32768, 177, 2147483647}
        _ = v8
    case -9223372036854775808:
        var v9 uint32 = 2
        _ = v9
    default:
        v6 = v6
    }
    v6 = v6
    var v10 uint32 = 4294967295
    _ = v10
    _ = v6
    return 344
}

func main() {
    var v1 uint32 = 955
    v2 := make(map[int]bool)
    var v3 bool = (!false)
    var v4 bool = false
    v5 := make(map[string]uint32)
    v4 = v3
    v11 := GenS1{F0: 4294967295, F1: -32768}
    fmt.Println(v4)
    v12 := GenS2{F0: v1, F1: (1 < -59), F2: 2}
    v13 := GenS2{F0: (v1 % ((2147483648) | 1)), F1: ((-32768 <= genFn0()) || v3), F2: 0}
    v14 := GenS3{F0: -2147483648, F1: (v1 | v1), F2: -1}
    var v15 uint32 = v1
    var v16 int16 = -32768
    v15 -= v15
    fmt.Println(v1)
    v17 := []int{9223372036854775807, 4294967295, 255, -2147483648}
    fmt.Println(v4, v1, v3)
    var v18 int32 = -2147483648
    v19 := GenS1{F0: (v1 & (v1 % ((2147483648) | 1))), F1: (v16 / ((v16) | 1))}
    delete(v2, -247)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
