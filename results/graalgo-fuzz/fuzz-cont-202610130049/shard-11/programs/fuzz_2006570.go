package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint8; F1 uint64 }
func genFn0() int16 {
    v6 := []int{969, 4294967295, -1}
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 int8 = 42
            switch fi7 {
            case -9223372036854775808:
                v10 := GenS1{F0: -32768, F1: 0}
                var v11 bool = (-1 > 42)
                fmt.Println(fi7, fi8)
                _ = v10
                _ = v11
            default:
                v12 := GenS1{F0: 42, F1: 163}
                _ = v12
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    _ = v6
    return 276
}
func genFn1(p0 int16, p1 string, p2 string) uint64 {
    p1 = (p1 + p2)
    var v13 uint32 = 826
    v14 := []int{}
    var v15 int64 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return 42
}

func main() {
    var v1 bool = (true || (0 > 65535))
    v2 := []int{}
    v3 := make(map[string]string)
    var v4 uint64 = 42
    v5 := []uint32{}
    fmt.Println(v1, v4)
    v16 := GenS2{F0: 128, F1: (v4 % ((4294967295) | 1))}
    _ = len(v5)
    switch 397 {
    case 676:
        v17 := GenS1{F0: 22, F1: 18446744073709551615}
        _ = v17
    case 9223372036854775807:
        v3["foo bar"] = ("\n\t" + "foo bar")
    }
    fmt.Println(v1, v4)
    var v18 int = -340
    var v19 int32 = 0
    var v20 int64 = 42
    fmt.Println(v1, v4, v1)
    var v21 float64 = ((3.14159 - 1.5) * 1e10)
    var v22 uint32 = 4294967295
    fmt.Println(v19, v1)
    v21 += v21
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
