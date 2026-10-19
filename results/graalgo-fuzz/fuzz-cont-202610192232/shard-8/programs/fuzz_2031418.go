package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint32) uint32 {
    v4 := []int{989, 1, 492}
    v5 := GenS1{F0: 42}
    v6 := GenS1{F0: 9223372036854775807}
    switch "\n\t" {
    case "a":
        p0 += (^(p0))
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return (p0 - (p0 * p0))
}
func genFn1() uint32 {
    var v7 int8 = 42
    var v8 int64 = 1
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    v9 := []int{4294967295, 2147483647, 691}
    var v10 float64 = float64(-2.25)
    v9 = append(v9, -1)
    v11 := GenS1{F0: 1}
    v10 += v10
    v10 += 1e10
    if (2 >= 172) {
        var v12 uint16 = 42
        if true {
            _ = len(v9)
            v13 := GenS1{F0: 42}
            var v14 int8 = -4
            _ = v13
            _ = v14
        }
        _ = v12
    }
    fmt.Println(v10)
    v15 := GenS2{F0: 4294967295, F1: 135}
    fmt.Println(v10)
    v10 -= 1.5
    var v16 int32 = -2147483648
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 int16 = 284
        var v19 int32 = 679
        v18 -= 32767
        _ = fi17
        _ = v18
        _ = v19
    }
    v10 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
