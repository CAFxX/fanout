package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int8; F2 int8 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int8; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int, p1 uint, p2 int) int64 {
    p2 = (p0 - -935)
    p2 -= p2
    p1 += 2
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1() int8 {
    var v5 int = 1
    for fi6 := 0; fi6 < 8; fi6++ {
        v5 = fi6
        _ = fi6
    }
    _ = v5
    return 0
}

func main() {
    var v1 uint16 = 1
    v2 := []string{"foo bar", ""}
    var v3 uint32 = 785
    v4 := make(map[uint32]uint32)
    if ((1 != v3) || (true && true)) {
        for fi7 := 0; fi7 < 3; fi7++ {
            v3 = (v3 ^ 2147483647)
            var v8 float32 = (float32(float64(((-2.25 / (1.5 / ((0.0 - 1.5) - 1e10))) + (0.1 + 0.1)))) / 1.5)
            _ = fi7
            _ = v8
        }
    } else {
        fmt.Println(v1)
    }
    v9 := GenS1{F0: (v3 + 2147483648)}
    v1 = 65535
    _ = len(v2)
    v1 -= v1
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS2{F0: 0, F1: 127, F2: 42}
        fmt.Println(v3)
        _ = fi10
        _ = v11
    }
    var v12 uint16 = 32767
    v13 := v4[(v3 * uint32(v1))]
    fmt.Println(v1, v13, v3)
    fmt.Println(v13, v3)
    if (572 != -9223372036854775808) {
        v14 := GenS3{F0: 0, F1: "foo bar"}
        var v15 int8 = -128
        v13 = v13
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 15; fi16++ {
            if false {
                for fi17 := 0; fi17 < 4; fi17++ {
                    v3 = v13
                    _ = fi17
                }
                v18 := GenS4{F0: 127, F1: false, F2: v3}
                v3 = (v13 ^ (v13 << ((4294967295) & 63)))
                _ = v18
            } else {
                delete(v4, (^(v3)))
            }
            fmt.Println(fi16, v12, v1)
            v19 := GenS2{F0: 127, F1: -1, F2: genFn1()}
            _ = fi16
            _ = v19
        }
    }
    v20 := GenS5{F0: v1, F1: v13}
    v21 := GenS4{F0: 127, F1: (-1 > 9223372036854775807), F2: 2}
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
