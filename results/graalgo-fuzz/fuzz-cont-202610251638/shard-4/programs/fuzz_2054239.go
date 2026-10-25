package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int16; F1 int; F2 uint8 }
func genFn0(p0 int32, p1 uint64, p2 uint8) int16 {
    p2 = (p2 + p2)
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(p0)
        _ = fi7
    }
    p2 -= 42
    _ = p0
    _ = p1
    _ = p2
    return -294
}
func genFn1(p0 int8, p1 uint16, p2 int8) uint {
    if (-598 <= -1) {
        switch -41 {
        case 32767:
            var v8 uint64 = 42
            _ = v8
        case -849:
            for fi9 := 0; fi9 < 14; fi9++ {
                v10 := GenS1{F0: 0, F1: (^(p1))}
                var v11 uint8 = 42
                _ = fi9
                _ = v10
                _ = v11
            }
        case -778:
            v12 := []int{32767, 253, 362, -819}
            _ = v12
        default:
            v13 := GenS2{F0: genFn0(733, 562, 1), F1: 42, F2: 224}
            _ = v13
        }
        if (("\n\t" != "foo bar") || false) {
            if (p2 >= (p2 / (((p0 & (p2 >> ((18446744073709551615) & 63)))) | 1))) {
                var v14 bool = (42 <= 0)
                p1 = 805
                var v15 uint32 = 2147483647
                _ = v14
                _ = v15
            }
        }
        var v16 int32 = int32(p1)
        _ = v16
    }
    p1 -= (^(p1))
    v17 := []int{579, 127, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 255
}

func main() {
    var v1 string = ""
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    var v4 int8 = 79
    var v5 uint32 = 1
    var v6 uint8 = 192
    var v18 int8 = 1
    if ((v6 << ((18446744073709551615) & 63)) <= 126) {
        v19 := []int{-2147483648, 859, 127, 42}
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 uint = 255
        _ = fi20
        _ = v21
    }
    delete(v3, 343)
    v22 := GenS1{F0: 883, F1: 32767}
    var v23 int32 = -1
    fmt.Println(v18, v6)
    v24 := GenS2{F0: 390, F1: -1, F2: v6}
    delete(v2, 32768)
    var v25 string = "\n\t"
    v26 := v2[1]
    v27 := []int{-9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
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
