package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint8, p1 uint8, p2 int32) uint16 {
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: 424, F1: true}
            fmt.Println(p2, fi7)
            var v9 uint8 = p1
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    fmt.Println(p0)
    p0 += p1
    v10 := []int{633, -2147483648, 1, 32768}
    v11 := GenS2{F0: (p0 + 0), F1: "hello"}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 255
}
func genFn1() uint64 {
    v12 := GenS3{F0: 0}
    v13 := []int{}
    var v14 uint8 = 127
    _ = v12
    _ = v13
    _ = v14
    return 285
}

func main() {
    var v1 float64 = 0.0
    var v2 uint = 15
    var v3 float64 = (v1 / 1.0)
    v4 := make(map[string]uint32)
    var v5 uint16 = 510
    v15 := []int{2147483647, -433}
    if ("a" == (((("\n\t" + "foo bar") + "") + ("\n\t" + (("\n\t" + "foo bar") + ""))) + "")) {
        switch "a" {
        case "foo bar":
            var v16 uint32 = 702
            _ = v16
        case "":
            for fi17 := 0; fi17 < 19; fi17++ {
                delete(v4, ("foo bar" + ""))
                _ = fi17
            }
        default:
            v2 -= 255
        }
        var v18 uint8 = 127
        v19 := v4[("\n\t" + ("\n\t" + "hello"))]
        _ = v18
        _ = v19
    } else {
        v20 := GenS1{F0: genFn0(1, 42, 2147483647), F1: true}
        v21 := GenS4{F0: 42}
        _ = v20
        _ = v21
    }
    v3 = float64(3.14159)
    fmt.Println(v2)
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    v5 += v5
    var v22 uint = (v2 - uint(v5))
    for fi23 := 0; fi23 < 10; fi23++ {
        fmt.Println(fi23, v1)
        var v24 int8 = 0
        _ = fi23
        _ = v24
    }
    v3 = v3
    v25 := GenS1{F0: uint16(v22), F1: (1 < 42)}
    v5 = (v5 % ((739) | 1))
    var v26 uint64 = 998
    for fi27 := 0; fi27 < 18; fi27++ {
        switch ((("a" + ("a" + "hello")) + "\n\t") + (((("hello" + ("" + "")) + "hello") + "\n\t") + "\n\t")) {
        case "":
            var v28 bool = (!((^(v5)) != genFn0(238, 2, 1)))
            _ = v28
        case "\n\t":
            switch (("hello" + ("a" + "foo bar")) + ("foo bar" + "\n\t")) {
            case "a":
                var v29 uint32 = 2
                _ = v29
            default:
                v1 -= (v3 * 0.0)
            }
        }
        v30 := GenS1{F0: (^(v5)), F1: true}
        fmt.Println(v26)
        _ = fi27
        _ = v30
    }
    v26 = 761
    v1 = v1
    var v31 uint16 = (v5 | (v5 & 255))
    v32 := GenS3{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
