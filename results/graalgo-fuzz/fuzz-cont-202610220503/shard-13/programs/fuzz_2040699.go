package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int64 }
func genFn0() uint64 {
    var v6 float32 = 1e10
    switch 42 {
    case 524:
        v7 := []int{1, -2147483648, 127, -2147483648}
        _ = v7
    case 875:
        fmt.Println(v6)
    case 32767:
        var v8 string = ""
        _ = v8
    }
    _ = v6
    return 870
}

func main() {
    var v1 int8 = 65
    var v2 int16 = 32767
    v3 := make(map[uint32]int)
    var v4 int16 = (v2 % ((-356) | 1))
    v5 := make(map[string]string)
    var v9 int64 = 1
    v1 += (v1 << ((255) & 63))
    v10 := GenS1{F0: v1}
    v9 -= -9223372036854775808
    if (!false) {
        var v11 float32 = ((1e10 * 3.14159) - 3.14159)
        fmt.Println(v11, v11, v4)
        _ = v11
    } else {
        v1 = v1
        var v12 int32 = -1
        var v13 uint = 2
        _ = v12
        _ = v13
    }
    var v14 uint64 = 18446744073709551615
    v4 = v2
    _ = len(v5)
    v9 = 899
    switch 584 {
    case -9223372036854775808:
        var v15 int32 = 1
        _ = v15
    case 946:
        var v16 uint64 = (v14 * genFn0())
        _ = v16
    case 127:
        v17 := GenS2{F0: (v9 * 9223372036854775807), F1: int64(v2)}
        _ = v17
    default:
        v9 = v9
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
