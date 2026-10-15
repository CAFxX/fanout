package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int32; F1 uint32 }
func genFn0() int64 {
    for fi4 := 0; fi4 < 13; fi4++ {
        if true {
            v5 := []int{fi4, fi4, (-(fi4)), fi4}
            for fi6 := 0; fi6 < 9; fi6++ {
                switch (fi6 + fi6) {
                case -2147483648:
                    v5 = append(v5, -453)
                case 32767:
                    _ = len(v5)
                }
                _ = len(v5)
                switch (fi4 / ((fi4) | 1)) {
                case 285:
                    switch "" {
                    case "hello":
                        var v7 int8 = -128
                        _ = v7
                    default:
                        var v8 int = (fi6 + (fi4 << ((65535) & 63)))
                        _ = v8
                    }
                    var v9 int16 = 847
                    var v10 int8 = 0
                    _ = v10
                    _ = v9
                case -844:
                    v5 = append(v5, (fi6 % ((fi6) | 1)))
                    if (((true || ("hello" > "")) && false) && false) {
                        v11 := GenS1{F0: 42, F1: 42}
                        _ = len(v5)
                        _ = v11
                    } else {
                        var v12 int8 = 63
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v12
                    }
                    var v13 int32 = 2147483647
                    if false {
                        v13 = v13
                    }
                    _ = v13
                default:
                    v5 = append(v5, (fi6 % ((fi6) | 1)))
                }
                _ = fi6
            }
            _ = v5
        } else {
            v14 := GenS1{F0: 127, F1: 292}
            switch (^(fi4)) {
            case 398:
                var v15 uint64 = 18446744073709551615
                _ = v15
            case 127:
                var v16 float32 = 1e10
                _ = v16
            default:
                v17 := GenS1{F0: 100, F1: 0}
                _ = v17
            }
            var v18 uint64 = 0
            _ = v14
            _ = v18
        }
        var v19 uint16 = 65535
        var v20 uint = 230
        _ = fi4
        _ = v19
        _ = v20
    }
    if (!true) {
        var v21 bool = true
        _ = v21
    } else {
        if (2 == 2) {
            var v22 uint64 = 1
            var v23 int64 = 0
            _ = v22
            _ = v23
        }
    }
    var v24 uint = 4294967295
    _ = v24
    return -691
}
func genFn1(p0 int8, p1 string) uint {
    for fi25 := 0; fi25 < 18; fi25++ {
        v26 := GenS2{F0: 42, F1: 42}
        if ((^(fi25)) < 9223372036854775807) {
            p0 += p0
        }
        _ = fi25
        _ = v26
    }
    v27 := []int{2147483647}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v27
    return 877
}

func main() {
    v1 := []int{128, -692, -1}
    v2 := make(map[string]uint32)
    var v3 float32 = 1e10
    v3 -= v3
    v28 := GenS1{F0: 13, F1: 0}
    switch 2147483647 {
    case 0:
        var v29 float32 = float32(float64(1.5))
        _ = v29
    case 9223372036854775807:
        for fi30 := 0; fi30 < 16; fi30++ {
            v1 = append(v1, 302)
            _ = fi30
        }
    }
    v3 -= 3.14159
    if true {
        v1 = append(v1, 616)
    }
    var v31 uint32 = 42
    var v32 uint16 = 0
    if true {
        v32 = (v32 % (((^(v32))) | 1))
    }
    if (0 > -6) {
        var v33 int16 = 1000
        var v34 uint8 = uint8(v31)
        fmt.Println(v34, v3)
        _ = v33
        _ = v34
    }
    v32 = v32
    v32 = v32
    if false {
        fmt.Println(v3, v32)
    } else {
        var v35 int32 = 1
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
}
