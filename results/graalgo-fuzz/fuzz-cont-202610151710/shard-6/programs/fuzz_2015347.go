package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int16, p1 int32, p2 int8) bool {
    v7 := GenS1{F0: 4294967295}
    switch "\n\t" {
    case "":
        p0 -= 42
    case "\n\t":
        v8 := []int{-1, 175, -1}
        _ = v8
    }
    if (0.1 < ((3.14159 + 0.1) - (((0.0 - (0.0 - 1e10)) + 1e10) / 1.0))) {
        for fi9 := 0; fi9 < 15; fi9++ {
            p2 = 19
            v10 := []int{fi9, (fi9 - fi9), fi9, (fi9 * fi9)}
            _ = fi9
            _ = v10
        }
    }
    if false {
        p2 -= p2
        p0 -= p0
        p0 += p0
    }
    var v11 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return (!(!false))
}

func main() {
    var v1 int = 4294967295
    var v2 int32 = 2147483647
    var v3 uint32 = 4294967295
    var v4 int64 = -1
    var v5 int64 = v4
    var v6 int16 = 1000
    fmt.Println(v1)
    v12 := []int{}
    if len(v12) > 0 {
        v12[0] = v1
    }
    if (false && ((false && false) && true)) {
        fmt.Println(v6)
    } else {
        fmt.Println(v2)
        var v13 float32 = (0.0 * float32(0.1))
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        fmt.Println(v4, v1, fi14)
        fmt.Println(fi14, v4, v6)
        _ = fi14
    }
    v1 -= 42
    switch v1 {
    case 32767:
        fmt.Println(v5, v4, v1)
    case 880:
        v5 -= v5
    case 952:
        v15 := GenS1{F0: 661}
        _ = v15
    default:
        v1 = 32767
    }
    var v16 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
