package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 string; F1 int16 }
func genFn0(p0 int8) bool {
    for fi6 := 0; fi6 < 6; fi6++ {
        switch "foo bar" {
        case "a":
            for fi7 := 0; fi7 < 9; fi7++ {
                p0 = p0
                var v8 int32 = 1
                var v9 int64 = 1
                _ = fi7
                _ = v8
                _ = v9
            }
        case "\n\t":
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 string = "foo bar"
                _ = fi10
                _ = v11
            }
        default:
            var v12 uint8 = 26
            _ = v12
        }
        switch (^(fi6)) {
        case -373:
            var v13 bool = true
            _ = v13
        default:
            v14 := []int{fi6}
            _ = v14
        }
        switch "" {
        case "a":
            var v15 uint8 = 42
            _ = v15
        case "\n\t":
            p0 = p0
        default:
            v16 := []int{32768, fi6, fi6}
            _ = v16
        }
        _ = fi6
    }
    p0 -= -54
    _ = p0
    return false
}
func genFn1(p0 uint64, p1 uint8, p2 int32) int32 {
    for fi17 := 0; fi17 < 15; fi17++ {
        v18 := GenS1{F0: 1, F1: (^(p0))}
        v19 := []int{42, fi17, 517}
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: -290, F1: (p0 * p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return p2
}

func main() {
    var v1 bool = true
    var v2 uint32 = 1
    v3 := []uint32{v2, v2, 82, v2}
    v4 := make(map[int]uint32)
    var v5 string = ("" + "\n\t")
    v2 = v2
    fmt.Println(v5, v2, v2)
    switch "a" {
    case "":
        v2 += v2
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "\n\t":
        var v21 uint16 = 279
        _ = v21
    default:
        _ = len(v4)
    }
    _ = len(v3)
    if len(v3) > 0 {
        v3[0] = v2
    }
    v4[-9223372036854775808] = v2
    if v1 {
        v5 = v5
        v22 := GenS2{F0: (("a" + "foo bar") + v5), F1: 655}
        v1 = (((1.5 - 0.1) / 1e10) < (3.14159 + 3.14159))
        _ = v22
    } else {
        v2 = (^(v2))
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS2{F0: (v5 + ((v5 + (((("" + v5) + ("foo bar" + "hello")) + "foo bar") + (v5 + "a"))) + v5)), F1: 1000}
            _ = fi23
            _ = v24
        }
    }
    fmt.Println(v2, v2, v5)
    fmt.Println(v5)
    var v25 int32 = 2147483647
    switch v5 {
    case "a":
        if (96 != 127) {
            v4[9223372036854775807] = v2
        }
    case "foo bar":
        var v26 bool = v1
        _ = v26
    }
    v27 := GenS1{F0: 358, F1: 601}
    var v28 string = v5
    switch 142 {
    case 2147483647:
        var v29 uint64 = 0
        _ = v29
    default:
        v2 -= v2
    }
    v30 := GenS1{F0: 0, F1: 2}
    var v31 float32 = 3.14159
    var v32 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
