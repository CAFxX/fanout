package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint }

func main() {
    var v1 int8 = 100
    v2 := []string{("hello" + "\n\t")}
    v3 := []uint32{1, 587, 4294967295, 2147483648}
    var v4 float64 = -2.25
    v5 := []string{"", "a", "hello", ""}
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 string = "\n\t"
        _ = fi6
        _ = v7
    }
    v2 = append(v2, (("" + "") + "a"))
    if false {
        v1 -= v1
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int32 = -2147483648
        var v10 uint32 = 227
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 = v1
    v11 := GenS1{F0: 2, F1: 42, F2: 255}
    switch "a" {
    case "hello":
        var v12 uint64 = 18446744073709551615
        _ = v12
    case "foo bar":
        v4 = ((((v4 / 1.0) * float64(-2.25)) - v4) * 1e10)
    default:
        if len(v3) > 0 {
            v3[0] = 65535
        }
    }
    fmt.Println(v4, v1)
    v4 -= 1.5
    v1 += (v1 + (v1 * (v1 / ((-2) | 1))))
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
