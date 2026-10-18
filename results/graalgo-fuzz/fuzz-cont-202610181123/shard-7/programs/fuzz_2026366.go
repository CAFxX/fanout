package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int; F1 uint16; F2 bool }
func genFn0(p0 int32, p1 string) string {
    v5 := GenS1{F0: 65535, F1: 418}
    p0 -= (p0 ^ p0)
    if false {
        fmt.Println(p0, p0)
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := []int{-126, fi6, 128}
            fmt.Println(p0)
            _ = fi6
            _ = v7
        }
        fmt.Println(p1, p0)
    } else {
        var v8 string = "hello"
        _ = v8
    }
    v9 := []int{127}
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return p1
}
func genFn1(p0 uint32, p1 uint32) string {
    v10 := GenS1{F0: 0, F1: (p0 << ((692) & 63))}
    p1 -= 2147483648
    p0 = p1
    for fi11 := 0; fi11 < 18; fi11++ {
        if (false && false) {
            fmt.Println(p1, fi11)
            p1 = 2147483647
            p1 += (p0 * p1)
        } else {
            var v12 float64 = 3.14159
            _ = v12
        }
        p0 = 1
        var v13 uint = 18446744073709551615
        _ = fi11
        _ = v13
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v10
    return "hello"
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = ((1e10 * ((-2.25 * 3.14159) + 3.14159)) * 1.5)
    v3 := make(map[string]int)
    v4 := []uint32{2147483648, 0, 4294967295, 477}
    v2 = v2
    v2 = v2
    var v14 int64 = 0
    var v15 uint16 = 2
    if false {
        v16 := GenS2{F0: -825, F1: (v15 * 364), F2: true}
        v14 = (v14 + (v14 % ((-1) | 1)))
        _ = v16
    }
    var v17 uint16 = 15
    v3["foo bar"] = 4294967295
    _ = len(v3)
    delete(v3, ("a" + ""))
    var v18 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
