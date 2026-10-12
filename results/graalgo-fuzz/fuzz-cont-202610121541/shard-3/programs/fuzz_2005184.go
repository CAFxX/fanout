package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
func genFn0() int64 {
    var v6 float64 = 0.0
    var v7 uint16 = 833
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int16) string {
    var v8 int64 = -9223372036854775808
    var v9 int16 = (p0 * p0)
    v10 := []int{42, 1}
    var v11 string = "\n\t"
    var v12 uint32 = 1
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v8
    _ = v9
    return (v11 + (v11 + v11))
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]int)
    var v3 uint64 = 1
    var v4 uint = 2
    v5 := []int{}
    var v13 uint8 = 184
    v14 := GenS1{F0: -128, F1: (^(v13))}
    v15 := GenS1{F0: -118, F1: (v13 >> ((18446744073709551615) & 63))}
    var v16 uint = v4
    fmt.Println(v16, v4)
    v17 := v2[113]
    var v18 string = "a"
    v19 := GenS1{F0: -4, F1: v13}
    fmt.Println(v4)
    var v20 float64 = v1
    var v21 float64 = v1
    for fi22 := 0; fi22 < 4; fi22++ {
        v21 += 3.14159
        v23 := GenS1{F0: -128, F1: v13}
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: 90, F1: v13}
    v5 = append(v5, (v17 % (((v17 << ((v16) & 63))) | 1)))
    delete(v2, (v17 ^ (v17 << (((v4 % (((v4 + (v16 & v16))) | 1))) & 63))))
    v16 -= (v16 / ((v4) | 1))
    fmt.Println(v1)
    var v25 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
