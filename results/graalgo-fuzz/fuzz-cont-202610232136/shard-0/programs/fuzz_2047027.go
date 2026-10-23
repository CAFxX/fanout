package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int }
func genFn0(p0 uint64, p1 uint16, p2 int8) int32 {
    fmt.Println(p1, p1, p1)
    switch "foo bar" {
    case "a":
        p0 = p0
    default:
        p2 = -1
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1(p0 int, p1 string) int {
    p1 = "a"
    var v7 uint64 = 2
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := []int{}
        var v10 float32 = 1.5
        var v11 int32 = 2147483647
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return (p0 & p0)
}

func main() {
    v1 := []string{"hello"}
    var v2 float32 = (((float32((-2.25 / (3.14159 / 1e10))) - 1e10) - 3.14159) - (((1.5 - (1e10 + ((float32(0.0) * ((0.0 / 0.1) * float32((0.0 - 0.1)))) + (1e10 * 0.0)))) / -2.25) + 1.5))
    v3 := []string{}
    v4 := make(map[int]int)
    v5 := []uint32{990, 612, 2}
    v6 := make(map[uint32]uint32)
    var v12 bool = true
    v12 = v12
    fmt.Println(v2, v12)
    if true {
        var v13 int = 1
        v12 = (-1 == -2147483648)
        var v14 int8 = -1
        _ = v13
        _ = v14
    } else {
        var v15 int8 = 8
        _ = v15
    }
    v2 = v2
    fmt.Println(v2, v2)
    v4[0] = 9223372036854775807
    v16 := GenS1{F0: 636, F1: 42, F2: 0}
    if v12 {
        fmt.Println(v12)
        v17 := v4[genFn1(-2147483648, ("a" + ("hello" + ((("a" + ("\n\t" + ("a" + ("" + "")))) + ("a" + ("foo bar" + "\n\t"))) + ("a" + "")))))]
        _ = v17
    }
    var v18 float32 = 1e10
    fmt.Println(v18)
    v4[128] = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
