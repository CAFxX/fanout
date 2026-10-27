package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 string }
func genFn0(p0 int32, p1 int8) string {
    if (9223372036854775807 != 42) {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 int16 = 0
            v9 := GenS1{F0: (fi7 + fi7)}
            var v10 float32 = float32((1e10 + float64((1.5 - (0.0 / (3.14159 / (1e10 + 1e10)))))))
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        var v11 uint = 1
        v11 -= (^(v11))
        _ = v11
    }
    if (0 != 1) {
        p0 += (^(p0))
        v12 := GenS1{F0: 42}
        if (false && false) {
            switch "\n\t" {
            case "":
                v13 := []int{-1, -558}
                _ = v13
            case "hello":
                p0 -= p0
            case "\n\t":
                v14 := []int{-2147483648, 32767, -1}
                _ = v14
            default:
                p0 += -1
            }
        }
        _ = v12
    }
    _ = p0
    _ = p1
    return "foo bar"
}
func genFn1(p0 uint) string {
    v15 := GenS1{F0: 0}
    v16 := []int{32768, 909, 32768, 532}
    _ = p0
    _ = v15
    _ = v16
    return (("\n\t" + "foo bar") + ("foo bar" + "foo bar"))
}

func main() {
    v1 := []uint32{}
    v2 := []int{}
    var v3 float32 = 1.5
    var v4 float32 = ((v3 - v3) * 0.1)
    var v5 uint32 = 2147483648
    var v6 float32 = ((v3 - v3) * 0.1)
    var v17 int16 = -480
    if false {
        fmt.Println(v3)
        fmt.Println(v17, v3, v4)
    }
    var v18 uint = 0
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 uint32 = v5
        v21 := GenS1{F0: (fi19 << ((410) & 63))}
        v22 := GenS1{F0: fi19}
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    v18 = v18
    var v23 uint64 = 838
    switch -1 {
    case -941:
        var v24 int16 = v17
        _ = v24
    case 9223372036854775807:
        v4 = v6
    case 4294967295:
        var v25 uint = (v18 * v18)
        _ = v25
    }
    switch "foo bar" {
    case "":
        var v26 int16 = (v17 % ((42) | 1))
        _ = v26
    case "hello":
        v27 := GenS2{F0: v18, F1: "a"}
        _ = v27
    }
    fmt.Println(v3, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
