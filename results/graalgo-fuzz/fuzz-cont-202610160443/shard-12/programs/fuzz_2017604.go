package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 bool }
type GenS2 struct { F0 string }
func genFn0() uint32 {
    v7 := []int{128}
    if len(v7) > 0 {
        v7[0] = -312
    }
    v8 := GenS1{F0: 2, F1: 4294967295, F2: true}
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint, p1 uint) uint8 {
    v9 := []int{-1, 32768, 4294967295}
    if ((42 <= 2) || true) {
        switch "\n\t" {
        case "\n\t":
            if (42 >= (p0 * p1)) {
                var v10 uint32 = 4294967295
                v11 := GenS1{F0: 812, F1: (v10 << (((p1 / ((uint(v10)) | 1))) & 63)), F2: (!(4294967295 < v10))}
                v12 := GenS2{F0: "\n\t"}
                _ = v10
                _ = v11
                _ = v12
            } else {
                if (0.0 != 0.0) {
                    var v13 int8 = -128
                    _ = v13
                }
            }
        case "foo bar":
            p0 = p0
        case "a":
            var v14 bool = (!(!((255 > 131) && false)))
            _ = v14
        default:
            fmt.Println(p0)
        }
    } else {
        var v15 bool = true
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        v17 := GenS2{F0: "hello"}
        p0 = (^(p1))
        for fi18 := 0; fi18 < 6; fi18++ {
            fmt.Println(p1)
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    if ((true && ((true || false) || false)) || ((1 != 229) && false)) {
        if true {
            p1 = p1
            p0 += 800
            p1 -= p0
        }
        fmt.Println(p0, p1)
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 bool = (!(!(false && false)))
            fmt.Println(p1)
            _ = fi19
            _ = v20
        }
    }
    p0 = p1
    _ = p0
    _ = p1
    _ = v9
    return 2
}

func main() {
    v1 := []int{864, 9223372036854775807}
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    var v4 float32 = 0.0
    var v5 bool = true
    v6 := []int{}
    v4 += -2.25
    switch "\n\t" {
    case "\n\t":
        var v21 uint8 = 128
        _ = v21
    case "foo bar":
        v4 = 0.1
    case "a":
        var v22 int = 128
        _ = v22
    }
    var v23 uint = 82
    var v24 int8 = -128
    fmt.Println(v5)
    var v25 uint32 = 0
    fmt.Println(v5, v5)
    var v26 string = "foo bar"
    var v27 uint16 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
