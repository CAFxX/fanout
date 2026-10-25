package main

import (
    "fmt"
)

func genFn0(p0 int32, p1 int64) int16 {
    fmt.Println(p1)
    p0 -= (p0 % ((p0) | 1))
    _ = p0
    _ = p1
    return 1000
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 18446744073709551615
    var v3 uint64 = 0
    fmt.Println(v2)
    v2 = v2
    v4 := v1[848]
    v3 = (v2 << ((270) & 63))
    v2 = (v3 << ((0) & 63))
    fmt.Println(v2, v2, v2)
    fmt.Println(v4, v2)
    v5 := []int{(v4 >> ((18446744073709551615) & 63)), v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
