package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int32; F1 int; F2 string }
func genFn0(p0 int8, p1 uint32) int16 {
    p0 += (p0 >> ((0) & 63))
    if (false || (false || true)) {
        for fi6 := 0; fi6 < 14; fi6++ {
            fmt.Println(p1, p0, p0)
            p0 += p0
            _ = fi6
        }
    }
    _ = p0
    _ = p1
    return 26
}
func genFn1(p0 int, p1 uint, p2 int8) int16 {
    switch (p0 ^ 42) {
    case 32768:
        if false {
            var v7 uint64 = 0
            _ = v7
        }
    case 2147483647:
        var v8 uint64 = 18446744073709551615
        _ = v8
    case -439:
        p1 = p1
    default:
        fmt.Println(p2, p1, p1)
    }
    v9 := GenS1{F0: 1, F1: "\n\t"}
    v10 := GenS2{F0: -2147483648, F1: p0, F2: ("\n\t" + "\n\t")}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]bool)
    var v3 string = "foo bar"
    v4 := make(map[int]int)
    var v5 uint16 = 42
    v5 -= (v5 * v5)
    delete(v2, -541)
    fmt.Println(v5, v5)
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 int16 = 810
        _ = fi11
        _ = v12
    }
    v1 = ("\n\t" + v1)
    v13 := []int{32768, 579}
    _ = len(v2)
    v14 := v4[0]
    v4[v14] = 151
    if true {
        if false {
            if true {
                if (61 > -1) {
                    if (!(!((float64(3.14159) == ((0.0 - float64(float32(-2.25))) / (1e10 / 1e10))) && false))) {
                        fmt.Println(v1, v14)
                        var v15 int32 = 0
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 3; fi16++ {
                        for fi17 := 0; fi17 < 12; fi17++ {
                            switch (v14 / (((-(fi16))) | 1)) {
                            case 582:
                                var v18 uint64 = 42
                                _ = v18
                            default:
                                delete(v4, 128)
                            }
                            v3 = v1
                            var v19 int64 = -1
                            _ = fi17
                            _ = v19
                        }
                        _ = fi16
                    }
                }
            }
        } else {
            for fi20 := 0; fi20 < 16; fi20++ {
                v21 := GenS2{F0: -2147483648, F1: (fi20 * 707), F2: "\n\t"}
                _ = fi20
                _ = v21
            }
            var v22 float32 = 3.14159
            if len(v13) > 0 {
                v13[0] = v14
            }
            _ = v22
        }
        if (-128 >= 1) {
            var v23 float64 = (((1.5 - -2.25) + -2.25) / 1.5)
            _ = v23
        } else {
            v1 = (("\n\t" + v1) + v3)
        }
        switch v14 {
        case 552:
            var v24 float32 = -2.25
            _ = v24
        case -436:
            v25 := GenS2{F0: 1, F1: v14, F2: (v1 + v3)}
            _ = v25
        }
    }
    if ((489 != 2) || (4294967295 <= 18446744073709551615)) {
        var v26 int = 2147483647
        v14 += v14
        _ = v26
    } else {
        if len(v13) > 0 {
            v13[0] = 727
        }
        v27 := v2[v14]
        fmt.Println(v5, v3, v3)
        _ = v27
    }
    switch (v14 % (((v14 << ((4294967295) & 63))) | 1)) {
    case 128:
        var v28 int64 = 113
        _ = v28
    default:
        v29 := GenS2{F0: -1, F1: (v14 << ((0) & 63)), F2: v1}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
