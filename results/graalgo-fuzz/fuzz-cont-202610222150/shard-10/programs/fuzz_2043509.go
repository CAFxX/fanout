package main

import (
    "fmt"
)

func genFn0(p0 int16, p1 int, p2 int32) int8 {
    p0 -= p0
    p2 -= (p2 & p2)
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 bool, p1 bool, p2 uint16) int {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 bool = ((!(1 < -141)) || p0)
        for fi7 := 0; fi7 < 17; fi7++ {
            fmt.Println(p0, p0)
            v8 := []int{(fi7 | fi5), 42}
            var v9 float32 = 1e10
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    v10 := []int{591, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 650
    var v3 int8 = 0
    var v4 float64 = 1.5
    var v11 int = -9223372036854775808
    var v12 uint32 = 299
    v2 += (-(v2))
    v13 := []int{v11, 32768, (v11 % (((^(v11))) | 1)), (v11 | genFn1((!true), ((v3 / ((12) | 1)) >= (v3 >> ((uint(v12)) & 63))), 0))}
    var v14 uint16 = 117
    var v15 int16 = 800
    v1[v12] = (v11 & 878)
    v15 += v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
