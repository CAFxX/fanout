package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int16, p1 int32) bool {
    var v6 int = 42
    p1 -= p1
    if false {
        var v7 int64 = 1
        _ = v7
    }
    var v8 int8 = 42
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return false
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = float64(0.0)
    v3 := make(map[int]bool)
    var v4 uint64 = 516
    var v5 uint64 = 837
    v9 := GenS1{F0: -2147483648}
    v5 -= (v4 << ((424) & 63))
    var v10 uint16 = 32767
    v11 := []int{}
    fmt.Println(v4, v10)
    v12 := v3[-850]
    v10 = 934
    v5 = (v4 >> ((uint(v4)) & 63))
    v13 := GenS1{F0: 1}
    fmt.Println(v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
