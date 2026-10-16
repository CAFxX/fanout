package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint) int32 {
    v7 := GenS1{F0: 1, F1: 255}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 2147483647
}
func genFn1(p0 int16) uint8 {
    v8 := GenS2{F0: 0}
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS2{F0: -2147483648}
        _ = fi9
        _ = v10
    }
    var v11 int = -1
    v12 := []int{1, v11, (v11 << ((18446744073709551615) & 63)), (v11 ^ (v11 - v11))}
    _ = p0
    _ = v11
    _ = v12
    _ = v8
    return 42
}

func main() {
    var v1 int64 = 0
    var v2 uint = 753
    v3 := make(map[int]int)
    var v4 float32 = 0.1
    v5 := []string{"\n\t", ("hello" + "")}
    v6 := make(map[int]string)
    v4 = v4
    v13 := GenS1{F0: 1, F1: 1}
    var v14 float64 = 0.1
    v2 += (v2 / (((v2 >> ((v2) & 63))) | 1))
    v2 = v2
    v4 = float32((v14 * v14))
    fmt.Println(v4, v2, v1)
    if (v2 != v2) {
        var v15 int64 = (^(v1))
        var v16 int = 127
        _ = len(v5)
        _ = v15
        _ = v16
    }
    var v17 uint16 = 32767
    v14 = v14
    var v18 uint64 = 192
    fmt.Println(v18)
    for fi19 := 0; fi19 < 1; fi19++ {
        _ = len(v5)
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
