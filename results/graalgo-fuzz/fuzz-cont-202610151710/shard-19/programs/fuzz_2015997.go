package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint }
type GenS2 struct { F0 string; F1 uint32 }
func genFn0(p0 int16, p1 int, p2 int64) uint64 {
    p0 -= -1
    fmt.Println(p2, p0, p0)
    var v5 uint32 = 689
    var v6 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 523
}
func genFn1(p0 int64, p1 string, p2 string) int {
    var v7 uint = 4294967295
    p0 += (p0 << ((v7) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 4294967295
}

func main() {
    var v1 int = 682
    v2 := make(map[int]int)
    var v3 uint64 = 1
    var v4 int8 = 42
    v1 = (v1 | (v1 + -538))
    v8 := []int{-9223372036854775808}
    fmt.Println(v3, v3, v4)
    var v9 float64 = float64(float32(1.5))
    switch genFn1(-9223372036854775808, "\n\t", "hello") {
    case 32767:
        v1 += (^(v1))
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 int64 = -336
        for fi12 := 0; fi12 < 17; fi12++ {
            v11 = (v11 >> ((255) & 63))
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    v3 = (v3 >> ((18446744073709551615) & 63))
    v13 := GenS1{F0: 2, F1: 65535, F2: 977}
    v1 = 1
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 float64 = (float64((1e10 * -2.25)) + (v9 - (v9 * (-2.25 + v9))))
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 134, F1: 0, F2: 505}
    var v17 float32 = 0.1
    v18 := GenS2{F0: "", F1: 2147483647}
    v4 = v4
    if false {
        var v19 float64 = float64((1.5 * 1.5))
        var v20 int = genFn1(9223372036854775807, ("a" + ("\n\t" + "")), "foo bar")
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
