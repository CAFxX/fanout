package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int32 }
func genFn0(p0 string) int32 {
    switch p0 {
    case "":
        var v6 uint = 4294967295
        _ = v6
    case "hello":
        if ((3.14159 / 3.14159) >= float64((1e10 - float32(0.0)))) {
            p0 = ""
            v7 := GenS1{F0: -645}
            p0 = "\n\t"
            _ = v7
        }
    default:
        p0 = ""
    }
    switch "a" {
    case "a":
        switch (p0 + p0) {
        case "hello":
            if false {
                switch 724 {
                case 2147483647:
                    var v8 string = p0
                    _ = v8
                case 127:
                    v9 := GenS1{F0: 656}
                    _ = v9
                default:
                    v10 := GenS1{F0: 963}
                    _ = v10
                }
            }
        case "foo bar":
            switch (p0 + "") {
            case "\n\t":
                v11 := GenS1{F0: 9223372036854775807}
                _ = v11
            case "foo bar":
                var v12 uint8 = 147
                _ = v12
            default:
                p0 = "a"
            }
        default:
            p0 = (p0 + (p0 + "a"))
        }
    case "hello":
        v13 := GenS2{F0: (9223372036854775807 < -460), F1: -696}
        _ = v13
    case "":
        v14 := GenS1{F0: 4294967295}
        _ = v14
    }
    p0 = ""
    fmt.Println(p0)
    _ = p0
    return 326
}
func genFn1(p0 int8, p1 uint64) uint {
    var v15 uint16 = 1
    var v16 float32 = float32(-2.25)
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 4294967295
}

func main() {
    var v1 int = 1
    var v2 bool = (!true)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    var v5 uint = 65535
    if false {
        fmt.Println(v5)
    }
    v3[503] = 4294967295
    v17 := []int{}
    if len(v17) > 0 {
        v17[0] = v1
    }
    v5 = (v5 & (v5 * v5))
    v1 = v1
    var v18 uint = (v5 & 414)
    switch v1 {
    case 1:
        var v19 int64 = 1
        _ = v19
    default:
        v18 -= genFn1(-128, 1)
    }
    v5 = (^(v5))
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := GenS1{F0: v1}
        v18 = v18
        _ = fi20
        _ = v21
    }
    v22 := GenS2{F0: v2, F1: -2147483648}
    fmt.Println(v2, v18)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
