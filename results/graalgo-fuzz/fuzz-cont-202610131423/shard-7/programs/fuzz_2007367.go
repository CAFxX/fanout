package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint64; F1 int16; F2 bool }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 bool = true
        _ = fi6
        _ = v7
    }
    if (!false) {
        v8 := []int{-114, 32767, 0, 9223372036854775807}
        _ = v8
    }
    var v9 uint = 4294967295
    _ = v9
    return 18446744073709551615
}
func genFn1(p0 uint8, p1 uint32) int64 {
    fmt.Println(p1)
    v10 := GenS1{F0: 537, F1: 922}
    _ = p0
    _ = p1
    _ = v10
    return 234
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 103
    var v3 int = 128
    var v4 string = "\n\t"
    v5 := []uint32{v2}
    if true {
        v11 := GenS1{F0: -1, F1: uint(v2)}
        _ = v11
    }
    var v12 uint8 = 3
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v13 bool = (true || (true && (!true)))
    v14 := GenS2{F0: 42, F1: 742, F2: (true || (("a" + "\n\t") < ("" + v4)))}
    var v15 float64 = 1.5
    if v13 {
        v15 -= 0.0
    } else {
        if len(v5) > 0 {
            v5[0] = (^(v2))
        }
    }
    var v16 uint32 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
