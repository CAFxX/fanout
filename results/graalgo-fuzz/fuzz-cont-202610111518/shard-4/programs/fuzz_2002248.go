package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint64; F1 bool }
func genFn0(p0 uint16) bool {
    p0 = (^(p0))
    v5 := []int{-2147483648}
    _ = p0
    _ = v5
    return ((2 >= (p0 * 65535)) && ((-2.25 >= 0.1) && ((9223372036854775807 >= 1) || (((p0 > p0) || true) && false))))
}
func genFn1(p0 string, p1 uint64, p2 string) int16 {
    v6 := GenS1{F0: 4294967295, F1: 500}
    v7 := GenS1{F0: 462, F1: 301}
    p0 = ((("" + "hello") + "\n\t") + p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 997
}

func main() {
    v1 := []uint32{65535}
    var v2 uint32 = 2
    v3 := []int{}
    v4 := make(map[string]string)
    v2 = (v2 | v2)
    fmt.Println(v2)
    for fi8 := 0; fi8 < 11; fi8++ {
        fmt.Println(fi8, v2)
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 float64 = 0.0
            v2 = 429
            fmt.Println(v10, v10)
            _ = fi9
            _ = v10
        }
        if (false && (true && true)) {
            var v11 int8 = 100
            v2 += v2
            _ = v11
        } else {
            for fi12 := 0; fi12 < 18; fi12++ {
                var v13 int = -649
                _ = len(v1)
                v2 = v2
                _ = fi12
                _ = v13
            }
        }
        _ = fi8
    }
    var v14 uint32 = v2
    fmt.Println(v14, v14)
    var v15 uint = uint(v14)
    v16 := GenS2{F0: 2, F1: false}
    var v17 int8 = -1
    var v18 int32 = -1
    for fi19 := 0; fi19 < 5; fi19++ {
        v14 -= (v14 >> (((v15 & v15)) & 63))
        _ = fi19
    }
    v20 := GenS1{F0: v2, F1: (v15 ^ (v15 | v15))}
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
