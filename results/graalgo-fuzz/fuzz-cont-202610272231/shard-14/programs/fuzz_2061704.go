package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint32, p1 int8) int8 {
    for fi4 := 0; fi4 < 6; fi4++ {
        fmt.Println(fi4, p1)
        v5 := []int{301, fi4}
        var v6 uint = 18446744073709551615
        _ = fi4
        _ = v5
        _ = v6
    }
    p1 -= (p1 % (((p1 * (p1 * (p1 / ((100) | 1))))) | 1))
    v7 := []int{127, -474, 59, -2147483648}
    _ = p0
    _ = p1
    _ = v7
    return p1
}

func main() {
    var v1 bool = false
    var v2 int = -1
    v3 := make(map[int]int)
    var v8 uint32 = 269
    var v9 int64 = 1
    var v10 uint = 200
    var v11 int8 = -80
    fmt.Println(v9, v9, v11)
    v12 := []int{(-(v2))}
    for fi13 := 0; fi13 < 1; fi13++ {
        fmt.Println(v11, v8, v9)
        _ = fi13
    }
    if v1 {
        v14 := GenS1{F0: 0}
        for fi15 := 0; fi15 < 4; fi15++ {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            switch "foo bar" {
            case "":
                var v16 float64 = 3.14159
                _ = v16
            case "\n\t":
                fmt.Println(v1)
            case "foo bar":
                if ((v10 + v10) != 1) {
                    v11 = (v11 % ((100) | 1))
                    fmt.Println(fi15, v1)
                    _ = len(v3)
                }
            }
            _ = fi15
        }
        fmt.Println(v9, v9)
        _ = v14
    } else {
        v12 = append(v12, v2)
        var v17 int64 = (^(v9))
        var v18 uint8 = 128
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: 2, F1: 1}
    v1 = ((v11 - v11) >= (-(v11)))
    switch "\n\t" {
    case "\n\t":
        for fi20 := 0; fi20 < 10; fi20++ {
            v11 = (v11 % ((v11) | 1))
            if len(v12) > 0 {
                v12[0] = fi20
            }
            _ = fi20
        }
    case "foo bar":
        v8 -= (v8 * (^(v8)))
    case "hello":
        var v21 int64 = (v9 ^ v9)
        _ = v21
    default:
        switch "\n\t" {
        case "hello":
            _ = len(v12)
        case "a":
            v22 := GenS1{F0: v9}
            _ = v22
        case "":
            var v23 int = v2
            _ = v23
        }
    }
    if ((v2 ^ v2) != v2) {
        var v24 int = (-(v2))
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
