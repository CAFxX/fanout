package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint8; F2 bool }
func genFn0() string {
    var v4 float64 = 1e10
    var v5 int = 2147483647
    v6 := GenS1{F0: 42, F1: 0}
    _ = v4
    _ = v5
    _ = v6
    return "\n\t"
}
func genFn1(p0 uint8) int8 {
    v7 := []int{128, 128}
    var v8 uint64 = 747
    v7 = append(v7, 1)
    if (2 == p0) {
        v9 := GenS1{F0: uint16(p0), F1: 128}
        _ = v9
    } else {
        var v10 int = 42
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 uint = 669
    var v2 int = 201
    var v3 bool = (1000 <= -1)
    v11 := []int{}
    for fi12 := 0; fi12 < 15; fi12++ {
        v11 = append(v11, 9223372036854775807)
        if len(v11) > 0 {
            v11[0] = fi12
        }
        _ = fi12
    }
    switch (v2 % (((v2 << ((v1) & 63))) | 1)) {
    case -1:
        var v13 string = "foo bar"
        _ = v13
    case 339:
        v1 += (v1 | (^(v1)))
    case -658:
        var v14 string = ("\n\t" + "foo bar")
        _ = v14
    }
    v15 := GenS1{F0: 559, F1: 0}
    if v3 {
        if len(v11) > 0 {
            v11[0] = v2
        }
        v16 := GenS2{F0: (("foo bar" + "a") >= ("" + "\n\t")), F1: 2, F2: (v2 <= v2)}
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 uint8 = 0
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    v3 = (0.0 > 1.5)
    if (!(0.0 > 0.1)) {
        var v19 uint = (v1 | (v1 % (((v1 % (((v1 + (v1 + (^(v1))))) | 1))) | 1)))
        _ = v19
    }
    v1 -= (^(v1))
    var v20 int8 = 127
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v2 = v2
    v1 = 65535
    if v3 {
        v1 -= (^(v1))
        var v21 int = 128
        _ = v21
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 uint32 = 0
            v24 = 65535
            v25 := GenS2{F0: ((((!v3) || v3) || v3) && v3), F1: 0, F2: (float64(float32((-2.25 / 0.1))) >= float64((0.0 * -2.25)))}
            var v26 uint32 = (v24 % ((v24) | 1))
            var v27 int = fi22
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 1; fi28++ {
            if v3 {
                var v29 uint = v1
                _ = v29
            }
            v11 = append(v11, fi22)
            v20 += v20
            _ = fi28
        }
        v20 += (v20 % (((v20 >> ((v1) & 63))) | 1))
        _ = fi22
    }
    if (2147483647 != 42) {
        for fi30 := 0; fi30 < 8; fi30++ {
            if (42 != -558) {
                if v3 {
                    if ((fi30 << (((v1 ^ v1)) & 63)) != -350) {
                        if v3 {
                            fmt.Println(v20, fi30)
                        } else {
                            v1 = v1
                        }
                        v1 = (v1 * 4294967295)
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    }
                    fmt.Println(v1, v1)
                    v1 = v1
                }
            } else {
                if false {
                    fmt.Println(v20, v3)
                    if (false && false) {
                        var v31 uint64 = 256
                        for fi32 := 0; fi32 < 11; fi32++ {
                            fmt.Println(fi30, v2)
                            _ = fi32
                        }
                        _ = v31
                    }
                    v33 := GenS2{F0: false, F1: 1, F2: (32767 >= 0)}
                    _ = v33
                }
            }
            fmt.Println(v1, v3)
            _ = fi30
        }
        v20 -= v20
    } else {
        fmt.Println(v2, v20, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v20
    _ = v3
}
