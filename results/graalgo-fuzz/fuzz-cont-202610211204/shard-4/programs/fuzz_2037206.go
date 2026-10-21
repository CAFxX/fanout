package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint }
type GenS2 struct { F0 int; F1 int8; F2 string }

func main() {
    var v1 string = "hello"
    v2 := []int{}
    v3 := []int{}
    v4 := GenS1{F0: 828, F1: -32768, F2: 4294967295}
    if (42 != 0) {
        switch 623 {
        case 128:
            var v5 int8 = -128
            _ = v5
        case 4294967295:
            switch 163 {
            case 431:
                var v6 float64 = 0.0
                _ = v6
            }
        case -861:
            v7 := GenS2{F0: -1, F1: 78, F2: ("" + "foo bar")}
            _ = v7
        default:
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 int16 = 322
                fmt.Println(fi8, v1)
                _ = fi8
                _ = v9
            }
        }
    }
    var v10 float64 = float64(((-2.25 * (float32(3.14159) * (1e10 * ((1e10 - float32(0.1)) / (-2.25 - (0.1 / (-2.25 / 1.5))))))) - (0.1 / 0.1)))
    switch 9223372036854775807 {
    case -1:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 32768:
        var v11 int64 = 0
        _ = v11
    }
    v10 = v10
    var v12 uint8 = 0
    v13 := GenS2{F0: -2147483648, F1: 15, F2: ("\n\t" + (("\n\t" + v1) + v1))}
    v1 = ""
    if false {
        v1 = v1
        var v14 uint = 42
        var v15 uint = v14
        _ = v14
        _ = v15
    }
    v1 = v1
    if (18446744073709551615 > 496) {
        var v16 uint = 18446744073709551615
        for fi17 := 0; fi17 < 10; fi17++ {
            v2 = append(v2, int(v16))
            var v18 uint = 671
            var v19 int64 = 970
            _ = fi17
            _ = v18
            _ = v19
        }
        var v20 float64 = (v10 + (1e10 * float64(0.1)))
        _ = v16
        _ = v20
    }
    if true {
        v10 = (v10 - ((1e10 / (v10 / 3.14159)) - v10))
    }
    v12 += (v12 | (v12 >> ((0) & 63)))
    v10 = v10
    v1 = v1
    fmt.Println(v1)
    if false {
        for fi21 := 0; fi21 < 11; fi21++ {
            var v22 float32 = 1.5
            v1 = v1
            _ = fi21
            _ = v22
        }
        v23 := GenS2{F0: -679, F1: -8, F2: v1}
        _ = v23
    }
    v24 := GenS1{F0: 102, F1: 32767, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
