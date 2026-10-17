package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() uint64 {
    var v5 uint8 = 255
    v5 -= v5
    _ = v5
    return 799
}
func genFn1(p0 int64) int {
    p0 -= 9223372036854775807
    fmt.Println(p0)
    if (18446744073709551615 > 0) {
        v6 := GenS1{F0: 128}
        var v7 int32 = 920
        _ = v6
        _ = v7
    } else {
        var v8 uint = 878
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: fi9}
            v8 -= 0
            v8 += 4294967295
            _ = fi9
            _ = v10
        }
        v11 := []int{-921}
        _ = v11
        _ = v8
    }
    _ = p0
    return 511
}

func main() {
    v1 := []int{-9223372036854775808, 9223372036854775807}
    v2 := make(map[uint32]string)
    var v3 float32 = 0.1
    v4 := make(map[int]int)
    fmt.Println(v3)
    v3 = (v3 + v3)
    delete(v4, 920)
    v3 += v3
    v2[45] = ("foo bar" + "")
    v3 = 0.1
    var v12 int32 = -224
    v12 -= (v12 & v12)
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 uint64 = 0
        _ = fi13
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
