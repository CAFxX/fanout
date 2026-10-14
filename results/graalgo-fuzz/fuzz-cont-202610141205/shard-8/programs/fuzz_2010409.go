package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int16; F1 uint32; F2 uint64 }
func genFn0(p0 uint32, p1 uint32) bool {
    fmt.Println(p1, p1)
    p0 = (p1 | 265)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return true
}
func genFn1() int {
    var v6 uint = 722
    v6 -= (^(v6))
    _ = v6
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = (float32(3.14159) * -2.25)
    v3 := make(map[uint32]uint32)
    var v4 string = (("\n\t" + (("" + ("foo bar" + "\n\t")) + "foo bar")) + "")
    var v5 int = 32768
    var v7 bool = (117 < 17)
    v8 := []int{(v5 % ((v5) | 1))}
    if (!(-1 >= 587)) {
        for fi9 := 0; fi9 < 12; fi9++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            v10 := GenS1{F0: 128, F1: fi9}
            var v11 float32 = (float32((((((float64(1.5) - (3.14159 - 1e10)) - (float64((v2 + float32(((1e10 - 1e10) * (3.14159 - float64((1.5 / v2))))))) * ((3.14159 * float64(v2)) - 1.5))) - 1e10) - ((1.5 + 1.5) + (0.1 - ((1.5 * float64(v2)) - (1.5 * 3.14159))))) + 0.1)) * v2)
            _ = fi9
            _ = v10
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 int8 = -16
            var v14 float64 = 3.14159
            v15 := GenS1{F0: v5, F1: v5}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    }
    fmt.Println(v2, v7, v4)
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := GenS1{F0: genFn1(), F1: (v5 / ((v5) | 1))}
        _ = fi16
        _ = v17
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    if (680 != 65535) {
        v8 = append(v8, (-(v5)))
        for fi18 := 0; fi18 < 15; fi18++ {
            v4 = v4
            _ = fi18
        }
    }
    v4 = ((("a" + "foo bar") + v4) + "\n\t")
    switch -273 {
    case 2147483647:
        switch genFn1() {
        case 234:
            var v19 float32 = 1.5
            _ = v19
        case 0:
            switch -117 {
            case 32767:
                fmt.Println(v5)
            default:
                v7 = true
            }
        case 32768:
            var v20 uint16 = 2
            _ = v20
        default:
            v21 := v3[1]
            _ = v21
        }
    default:
        v5 += (v5 | (v5 / (((v5 >> ((18446744073709551615) & 63))) | 1)))
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 int = genFn1()
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 uint32 = 252
            var v26 float32 = v2
            _ = fi24
            _ = v25
            _ = v26
        }
        if (-1 != 0) {
            var v27 uint16 = 42
            v28 := GenS1{F0: (v23 - fi22), F1: (-(fi22))}
            v5 -= (v23 % (((v5 >> ((65535) & 63))) | 1))
            _ = v27
            _ = v28
        } else {
            var v29 int64 = -9223372036854775808
            _ = v29
        }
        _ = fi22
        _ = v23
    }
    switch genFn1() {
    case 817:
        if len(v8) > 0 {
            v8[0] = v5
        }
    default:
        fmt.Println(v4, v4)
    }
    v30 := GenS2{F0: 388, F1: 4294967295, F2: 0}
    var v31 uint = 4294967295
    fmt.Println(v5, v5, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
