package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int) uint64 {
    v6 := GenS1{F0: 255}
    v7 := GenS1{F0: 2}
    p0 = p0
    _ = p0
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 42
    var v3 uint32 = 4294967295
    var v4 int = -2147483648
    v5 := []int{(v4 & 2147483647), v4, v4}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v2 -= v2
    var v8 int = v4
    v8 -= -1
    v3 += v3
    var v9 int8 = 39
    v2 += (v2 + (v2 / (((^(v2))) | 1)))
    var v10 float32 = 1e10
    if len(v5) > 0 {
        v5[0] = v4
    }
    v10 -= ((v10 - 0.0) / ((v10 + (3.14159 + v10)) * v10))
    v5 = append(v5, (v4 / ((-9223372036854775808) | 1)))
    v10 += v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
