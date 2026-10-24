package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int16 }
func genFn0(p0 int, p1 uint8) uint32 {
    var v5 float64 = 1e10
    v6 := GenS1{F0: 47, F1: 0, F2: 998}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint64 = 0
    var v3 int16 = -957
    var v4 int64 = 9223372036854775807
    fmt.Println(v4)
    v3 -= (-(v3))
    v2 += (v1 & (v1 << ((0) & 63)))
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint8 = 2
        v9 := []int{(-(fi7))}
        var v10 bool = ((false && false) && (("hello" + ("a" + "")) == "a"))
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := []int{}
    v3 -= v3
    fmt.Println(v3, v3, v4)
    for fi12 := 0; fi12 < 6; fi12++ {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        v2 += v1
        _ = fi12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = fi13
    }
    fmt.Println(v1, v2, v3)
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
