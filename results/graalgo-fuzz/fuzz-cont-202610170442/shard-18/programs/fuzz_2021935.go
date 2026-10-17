package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 string; F2 int }
func genFn0(p0 uint) int32 {
    v6 := GenS1{F0: 1}
    var v7 int = -743
    var v8 int = v7
    p0 -= (^(p0))
    switch (v8 | v8) {
    case 2147483647:
        v9 := GenS1{F0: 65535}
        _ = v9
    case 897:
        v10 := []int{(v7 / ((v8) | 1)), v7}
        _ = v10
    default:
        v11 := GenS2{F0: 510, F1: "hello", F2: 1}
        _ = v11
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint32 = 65535
    var v2 int8 = 1
    v3 := make(map[string]uint32)
    var v4 bool = (!(!false))
    var v5 uint8 = 108
    v2 = v2
    v12 := []int{127}
    v2 = (v2 % (((v2 / (((v2 << ((2) & 63))) | 1))) | 1))
    for fi13 := 0; fi13 < 4; fi13++ {
        v1 -= (v1 * v1)
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        v1 = 2
        _ = fi13
    }
    v1 -= (v1 + (v1 - v1))
    v14 := GenS2{F0: -923, F1: "foo bar", F2: 4294967295}
    if (v4 && v4) {
        v2 = v2
        if v4 {
            var v15 int16 = -1
            var v16 uint64 = 18446744073709551615
            _ = v15
            _ = v16
        }
        v4 = (v4 || ((v4 || v4) && (true && v4)))
    }
    _ = len(v12)
    v17 := v3["foo bar"]
    v12 = append(v12, 9223372036854775807)
    if len(v12) > 0 {
        v12[0] = 128
    }
    fmt.Println(v5)
    v18 := GenS1{F0: 2}
    v19 := v3["a"]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
