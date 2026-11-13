package main

import (
    "fmt"
)

func genFn0(p0 uint8) int {
    fmt.Println(p0)
    v7 := []int{42, -9223372036854775808}
    switch "a" {
    case "\n\t":
        p0 -= (p0 >> ((923) & 63))
    default:
        p0 = p0
    }
    var v8 uint32 = 556
    var v9 string = "a"
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -95
}

func main() {
    var v1 float64 = float64(0.1)
    var v2 uint64 = 4294967295
    var v3 uint64 = 740
    v4 := []uint32{65535, 4294967295, 280}
    v5 := []string{"\n\t", "a"}
    var v6 int16 = -1
    fmt.Println(v1, v2, v2)
    var v10 uint32 = 4294967295
    var v11 bool = true
    v10 = (v10 & v10)
    if len(v5) > 0 {
        v5[0] = "a"
    }
    fmt.Println(v6, v6, v3)
    v4 = append(v4, 42)
    var v12 int = genFn0(1)
    var v13 uint64 = (^(v2))
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
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
