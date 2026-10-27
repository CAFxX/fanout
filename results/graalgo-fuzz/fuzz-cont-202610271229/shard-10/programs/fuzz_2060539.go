package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int32, p1 int8, p2 int) bool {
    var v4 uint = 18446744073709551615
    if (true || false) {
        v5 := GenS1{F0: 42, F1: p1}
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (((1.5 / 1.0) + (3.14159 - float32(-2.25))) < 0.1)
}
func genFn1(p0 int64, p1 bool) uint32 {
    v6 := []int{32767, 32767, -667, 32768}
    p0 -= p0
    var v7 uint32 = 2147483648
    fmt.Println(p1, p1, v7)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return v7
}

func main() {
    v1 := []int{32768, 2147483647}
    v2 := make(map[string]string)
    var v3 int64 = 498
    switch "" {
    case "":
        switch 32767 {
        case 2147483647:
            v8 := GenS1{F0: 255, F1: 105}
            _ = v8
        default:
            v3 -= (v3 | (v3 | 465))
        }
    default:
        _ = len(v2)
    }
    var v9 bool = true
    var v10 uint8 = 2
    var v11 int32 = 1
    var v12 uint32 = 2147483648
    v1 = append(v1, 2147483647)
    var v13 int = 2147483647
    switch (("\n\t" + "foo bar") + ("hello" + ("\n\t" + ""))) {
    case "\n\t":
        switch (v13 / ((4294967295) | 1)) {
        case -730:
            var v14 uint64 = 2
            _ = v14
        default:
            var v15 uint16 = 65535
            _ = v15
        }
    case "":
        v16 := GenS2{F0: v11, F1: 42}
        _ = v16
    default:
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 uint8 = (^(v10))
            if true {
                var v19 uint64 = 42
                v20 := GenS3{F0: (v18 ^ (v18 % (((^(v18))) | 1)))}
                var v21 int16 = 889
                _ = v19
                _ = v20
                _ = v21
            }
            v22 := GenS1{F0: 234, F1: -87}
            _ = fi17
            _ = v18
            _ = v22
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 float32 = (((3.14159 * 1e10) + 1.5) * float32(1e10))
        switch "\n\t" {
        case "":
            if false {
                v24 = (0.0 + float32(float64(1.5)))
            }
        }
        _ = fi23
        _ = v24
    }
    v25 := GenS1{F0: uint16(v12), F1: -1}
    v26 := GenS2{F0: 987, F1: 18446744073709551615}
    v9 = (127 == (v10 - 127))
    for fi27 := 0; fi27 < 14; fi27++ {
        var v28 uint8 = 2
        var v29 float32 = 1e10
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v9
}
