package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 bool }
type GenS2 struct { F0 uint32; F1 uint16; F2 int8 }
func genFn0(p0 int, p1 bool, p2 uint8) uint8 {
    if p1 {
        v5 := GenS1{F0: (p2 >> ((2) & 63)), F1: (false && false), F2: p1}
        _ = v5
    }
    v6 := []int{}
    for fi7 := 0; fi7 < 16; fi7++ {
        if (42 > 434) {
            switch "\n\t" {
            case "\n\t":
                if (p1 || (1.5 < 0.0)) {
                    for fi8 := 0; fi8 < 10; fi8++ {
                        v9 := GenS1{F0: p2, F1: (0 > -653), F2: (2147483648 > 2147483648)}
                        _ = fi8
                        _ = v9
                    }
                    var v10 int32 = -2147483648
                    var v11 bool = false
                    _ = v10
                    _ = v11
                } else {
                    if false {
                        v12 := GenS1{F0: p2, F1: p1, F2: (0 == 9223372036854775807)}
                        var v13 uint16 = 792
                        _ = v12
                        _ = v13
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        p2 = 0
                        var v14 bool = p1
                        _ = v14
                    }
                    if (!true) {
                        var v15 float32 = float32(-2.25)
                        _ = v15
                    }
                    p0 -= (p0 / ((p0) | 1))
                }
            case "":
                v6 = append(v6, (fi7 >> ((4294967295) & 63)))
            }
        } else {
            p2 -= p2
        }
        v6 = append(v6, p0)
        _ = fi7
    }
    v6 = append(v6, 112)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}

func main() {
    v1 := []int{32767, 0, 267}
    var v2 uint32 = 107
    v3 := make(map[int]int)
    var v4 int8 = -109
    var v16 uint = uint(v2)
    v17 := GenS2{F0: (^(v2)), F1: 243, F2: (v4 ^ (v4 - 42))}
    var v18 int = int(v4)
    switch "a" {
    case "\n\t":
        v4 = v4
    case "":
        v2 = v2
    case "hello":
        v2 -= v2
    default:
        v1 = append(v1, -1)
    }
    v16 -= v16
    var v19 int8 = v4
    var v20 float64 = (1.5 / 1.5)
    var v21 float32 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
