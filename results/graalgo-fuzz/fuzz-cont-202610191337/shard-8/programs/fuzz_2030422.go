package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 bool }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 string; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint64, p1 bool) int16 {
    v6 := []int{960, 32767, -345, 2147483647}
    if len(v6) > 0 {
        v6[0] = -42
    }
    var v7 uint32 = 42
    switch "\n\t" {
    case "":
        var v8 uint = uint(v7)
        _ = v8
    case "\n\t":
        if (!((1000 == 1) && p1)) {
            _ = len(v6)
        }
    default:
        for fi9 := 0; fi9 < 3; fi9++ {
            p1 = (-1 >= 0)
            v10 := GenS1{F0: 42, F1: 29, F2: (p1 || (1 != uint8(p0)))}
            _ = len(v6)
            _ = fi9
            _ = v10
        }
    }
    v7 -= (v7 - v7)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -32768
}

func main() {
    var v1 uint32 = 0
    v2 := []string{"a"}
    var v3 int = 9223372036854775807
    var v4 uint64 = 550
    v5 := make(map[uint32]string)
    v3 -= (v3 | v3)
    var v11 bool = (42 != 42)
    if (42 != (v1 << ((523) & 63))) {
        v4 = v4
        v5[uint32(v4)] = ((("a" + (("a" + "a") + ("a" + "a"))) + "hello") + ("foo bar" + "\n\t"))
    } else {
        var v12 float64 = (((-2.25 - -2.25) + (3.14159 / 1.0)) / float64(0.1))
        _ = len(v5)
        _ = v12
    }
    v13 := GenS1{F0: 42, F1: 2, F2: v11}
    switch "foo bar" {
    case "hello":
        var v14 string = "hello"
        _ = v14
    case "":
        if (100 > -17) {
            v15 := v5[(v1 + v1)]
            _ = v15
        }
    case "\n\t":
        var v16 int = v3
        _ = v16
    default:
        v17 := GenS1{F0: 127, F1: 128, F2: (((false || (255 == 1)) && (615 == 66)) || (!(!true)))}
        _ = v17
    }
    var v18 string = "\n\t"
    if v11 {
        switch -524 {
        case 128:
            var v19 int = v3
            _ = v19
        case 863:
            fmt.Println(v3)
        case -601:
            v20 := GenS2{F0: 2, F1: 42}
            _ = v20
        default:
            switch v3 {
            case 127:
                var v21 uint = 0
                _ = v21
            case 1:
                fmt.Println(v18)
            case 90:
                v11 = v11
            default:
                if v11 {
                    var v22 float64 = ((1e10 + ((-2.25 + (1.5 - 1.5)) - float64((((3.14159 * (0.0 / float32(-2.25))) - 1e10) * -2.25)))) * 0.1)
                    for fi23 := 0; fi23 < 6; fi23++ {
                        var v24 int = (-(fi23))
                        v18 = v18
                        _ = fi23
                        _ = v24
                    }
                    _ = v22
                } else {
                    v25 := GenS1{F0: 1, F1: 127, F2: (v4 > (v4 - (v4 ^ 18446744073709551615)))}
                    _ = v25
                }
            }
        }
        for fi26 := 0; fi26 < 6; fi26++ {
            for fi27 := 0; fi27 < 6; fi27++ {
                v28 := GenS3{F0: v18, F1: 49, F2: 65535}
                _ = fi27
                _ = v28
            }
            if (!((1 != 1) && v11)) {
                v3 = v3
                delete(v5, (v1 >> ((65535) & 63)))
                v1 = (v1 / ((642) | 1))
            }
            _ = fi26
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        v2 = append(v2, v18)
        _ = fi29
    }
    v18 = v18
    if v11 {
        v3 = v3
        v30 := GenS4{F0: (^(v1)), F1: v1}
        var v31 int32 = 1
        _ = v30
        _ = v31
    }
    var v32 bool = v11
    var v33 uint32 = (v1 - (v1 & v1))
    v34 := GenS1{F0: 100, F1: 127, F2: v11}
    v18 = v18
    v5[(v1 ^ v1)] = (v18 + v18)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
