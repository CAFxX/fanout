package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int64; F1 bool }
func genFn0(p0 string, p1 bool) int {
    fmt.Println(p1, p1)
    v7 := GenS1{F0: 42, F1: 4294967295}
    _ = p0
    _ = p1
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 uint = 1
    var v2 string = ("\n\t" + "hello")
    var v3 int = -794
    v4 := make(map[int]bool)
    var v5 uint16 = 677
    v6 := []uint32{4294967295, 288, 65535}
    var v8 string = v2
    var v9 uint8 = 2
    for fi10 := 0; fi10 < 16; fi10++ {
        _ = len(v4)
        v8 = v2
        if true {
            var v11 int16 = 0
            _ = v11
        }
        _ = fi10
    }
    fmt.Println(v9)
    for fi12 := 0; fi12 < 14; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 float32 = -2.25
            if true {
                fmt.Println(fi12)
                switch v8 {
                case "\n\t":
                    if (false && false) {
                        v2 = (("\n\t" + "hello") + "\n\t")
                        var v15 uint16 = 0
                        _ = v15
                    }
                case "a":
                    var v16 int64 = -1
                    var v17 float64 = 1.5
                    v18 := GenS1{F0: -2147483648, F1: (^(v1))}
                    _ = v16
                    _ = v17
                    _ = v18
                default:
                    var v19 string = (((("a" + "\n\t") + v8) + (v8 + ("" + v2))) + ("a" + ""))
                    _ = v19
                }
            }
            v1 = v1
            var v20 float64 = 3.14159
            _ = fi13
            _ = v14
            _ = v20
        }
        v21 := GenS2{F0: 170, F1: true}
        _ = fi12
        _ = v21
    }
    switch genFn0(v8, (!false)) {
    case -1:
        v5 = (v5 - (v5 % ((v5) | 1)))
    case 29:
        v22 := GenS2{F0: -9223372036854775808, F1: false}
        _ = v22
    default:
        fmt.Println(v8)
    }
    v3 -= genFn0("\n\t", true)
    fmt.Println(v5, v9, v3)
    var v23 int32 = -1
    v6 = append(v6, 585)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
