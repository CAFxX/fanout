package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0(p0 uint8, p1 uint16, p2 int) int8 {
    for fi6 := 0; fi6 < 15; fi6++ {
        p0 += p0
        _ = fi6
    }
    var v7 uint16 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}
func genFn1(p0 uint32, p1 int, p2 int64) int32 {
    var v8 float64 = (-2.25 - (0.0 - 0.1))
    p1 = (p1 & p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint8 = 3
    v3 := make(map[int]string)
    v4 := []int{113, 825, 128}
    v5 := make(map[int]string)
    fmt.Println(v1, v1)
    v1 = 2147483648
    fmt.Println(v1)
    fmt.Println(v1, v2)
    var v9 uint16 = 0
    v10 := GenS1{F0: 2, F1: (v1 * v1)}
    var v11 uint = 255
    v9 = v9
    v1 = (v1 | (v1 | (v1 - uint32(v2))))
    var v12 string = "\n\t"
    var v13 float32 = 0.0
    delete(v5, 9223372036854775807)
    switch 32767 {
    case 4294967295:
        v14 := GenS2{F0: 18446744073709551615, F1: 2}
        _ = v14
    case 9223372036854775807:
        fmt.Println(v11)
    default:
        v11 -= (v11 ^ (^(v11)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
