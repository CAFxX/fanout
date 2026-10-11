package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 bool }
func genFn0(p0 int8, p1 uint64) uint {
    switch -95 {
    case -1:
        if (2 <= 2) {
            p1 = (p1 ^ (p1 % ((p1) | 1)))
            fmt.Println(p1, p0)
            var v7 uint16 = 1
            _ = v7
        }
    case 2147483647:
        p1 = (^(p1))
    default:
        fmt.Println(p1, p1)
    }
    v8 := []int{154}
    _ = p0
    _ = p1
    _ = v8
    return 92
}
func genFn1(p0 int64) uint16 {
    fmt.Println(p0)
    var v9 uint64 = 890
    _ = p0
    _ = v9
    return 255
}

func main() {
    v1 := []uint32{2147483647, 1}
    var v2 float64 = 1.5
    var v3 bool = (!(!(42 >= 132)))
    var v4 float64 = v2
    var v5 bool = (-2147483648 > -1)
    var v6 uint32 = 65535
    v10 := GenS1{F0: 0}
    if (v3 && (v5 || (-2.25 < v4))) {
        v4 -= ((v2 - ((1.5 / v2) - v4)) - ((((v4 * v2) - v2) / 1e10) / 1e10))
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "foo bar" {
    case "a":
        v5 = true
    }
    switch "foo bar" {
    case "a":
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := GenS2{F0: (fi11 - 127), F1: false}
            _ = len(v1)
            _ = fi11
            _ = v12
        }
    case "\n\t":
        var v13 uint32 = v6
        _ = v13
    }
    var v14 int8 = 127
    v1 = append(v1, v6)
    v15 := GenS2{F0: -2147483648, F1: (genFn1(9223372036854775807) < 163)}
    fmt.Println(v2, v14)
    if true {
        var v16 float64 = ((1e10 - (v2 - float64(1e10))) * (-2.25 + (((float64(((0.1 + 1e10) * 0.0)) - float64((0.0 / 0.1))) / v2) / (float64(0.1) / v2))))
        v6 -= 895
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
