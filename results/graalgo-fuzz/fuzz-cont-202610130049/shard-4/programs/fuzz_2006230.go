package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int16 }
func genFn0(p0 int8) uint8 {
    p0 = p0
    p0 += p0
    _ = p0
    return 255
}
func genFn1(p0 uint8, p1 uint16, p2 int64) uint32 {
    v7 := GenS1{F0: -1}
    var v8 uint32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 uint8 = 1
    var v2 bool = false
    var v3 int = -516
    var v4 int64 = 1
    v5 := []string{"foo bar"}
    v6 := []int{v3, -9223372036854775808, 2147483647, (-(v3))}
    if v2 {
        v5 = append(v5, "a")
        v1 = (v1 & (v1 * 0))
    } else {
        if v2 {
            v9 := GenS2{F0: v3, F1: -153}
            _ = v9
        }
    }
    var v10 uint64 = 42
    var v11 float64 = 1.5
    fmt.Println(v4, v10)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v12 int16 = -26
    var v13 uint32 = 4294967295
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
