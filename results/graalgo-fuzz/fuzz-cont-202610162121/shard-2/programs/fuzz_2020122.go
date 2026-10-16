package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int8, p1 int8, p2 uint16) int64 {
    v4 := []int{-868}
    v5 := GenS1{F0: p1}
    v4 = append(v4, -9223372036854775808)
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            p1 -= (p1 * p0)
            _ = fi7
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -293
}

func main() {
    v1 := []int{}
    v2 := []uint32{2}
    var v3 float64 = 3.14159
    switch 1 {
    case -9223372036854775808:
        fmt.Println(v3)
    default:
        var v8 uint16 = 32767
        _ = v8
    }
    v3 = float64(0.0)
    if false {
        if false {
            for fi9 := 0; fi9 < 17; fi9++ {
                fmt.Println(v3, v3)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi9
            }
        }
        var v10 float32 = 1.5
        _ = v10
    }
    var v11 uint = 0
    v3 -= 1e10
    var v12 uint64 = 18446744073709551615
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v11 += (v11 | 255)
    switch "a" {
    case "":
        v12 = v12
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = 304
        }
    case "a":
        var v13 string = "foo bar"
        _ = v13
    default:
        var v14 float32 = 1e10
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
}
