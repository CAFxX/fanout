package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int64 }
func genFn0(p0 int64) string {
    p0 -= p0
    var v4 float64 = float64(1e10)
    v5 := []int{127, -1, 4294967295, 9223372036854775807}
    _ = len(v5)
    _ = p0
    _ = v4
    _ = v5
    return "foo bar"
}
func genFn1(p0 uint32, p1 int64) int64 {
    v6 := []int{-274, 0, 4294967295}
    var v7 int16 = 32767
    if (!true) {
        fmt.Println(p1)
    } else {
        fmt.Println(p1)
        for fi8 := 0; fi8 < 14; fi8++ {
            v7 -= v7
            _ = fi8
        }
        fmt.Println(v7, p1)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p1 / ((p1) | 1))
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    v3 := []string{"hello"}
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v3 = append(v3, ("hello" + "a"))
    var v9 uint64 = 18446744073709551615
    switch (("hello" + ("a" + ("a" + "\n\t"))) + "foo bar") {
    case "":
        var v10 bool = true
        _ = v10
    case "\n\t":
        fmt.Println(v9)
    }
    fmt.Println(v9)
    v9 = 2
    switch "a" {
    case "hello":
        fmt.Println(v9)
    case "\n\t":
        v11 := GenS1{F0: 132, F1: 0, F2: 0}
        _ = v11
    default:
        var v12 uint8 = 127
        _ = v12
    }
    v9 -= 42
    var v13 int16 = 42
    fmt.Println(v9, v13)
    var v14 uint64 = 42
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 bool = true
        _ = fi15
        _ = v16
    }
    var v17 int64 = 9223372036854775807
    v18 := GenS1{F0: 42, F1: -408, F2: v17}
    switch -892 {
    case 813:
        var v19 float32 = 3.14159
        _ = v19
    case 651:
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 float32 = (float32(3.14159) * 1e10)
            switch "foo bar" {
            case "":
                v17 = v17
            case "\n\t":
                var v22 uint8 = 100
                _ = v22
            default:
                if true {
                    v14 -= 18446744073709551615
                }
            }
            _ = fi20
            _ = v21
        }
    default:
        var v23 int64 = v17
        _ = v23
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
