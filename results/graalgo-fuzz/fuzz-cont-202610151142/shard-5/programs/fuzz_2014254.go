package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint64 }
func genFn0() string {
    v7 := GenS1{F0: (-431 > -1), F1: 987}
    if true {
        var v8 int = -2147483648
        v9 := GenS2{F0: 134}
        v8 -= v8
        _ = v8
        _ = v9
    } else {
        v10 := GenS3{F0: 1, F1: 3, F2: 2}
        _ = v10
    }
    v11 := []int{32768, -2147483648, 32768, 396}
    _ = v11
    _ = v7
    return "foo bar"
}
func genFn1(p0 uint16) uint {
    if true {
        v12 := []int{-2147483648}
        if len(v12) > 0 {
            v12[0] = 139
        }
        v12 = append(v12, 127)
        _ = v12
    }
    if true {
        var v13 int32 = -1
        fmt.Println(p0, v13)
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := []int{}
        _ = fi14
        _ = v15
    }
    if ((true && true) && false) {
        for fi16 := 0; fi16 < 2; fi16++ {
            p0 -= 42
            _ = fi16
        }
    }
    _ = p0
    return 0
}

func main() {
    var v1 uint8 = 62
    var v2 uint16 = 2
    var v3 float64 = ((3.14159 - -2.25) * (3.14159 / 1.0))
    var v4 uint8 = v1
    var v5 int8 = 127
    var v6 uint64 = 2
    var v17 uint = genFn1(v2)
    v1 = v4
    v18 := []int{}
    v18 = append(v18, 9223372036854775807)
    fmt.Println(v6, v3)
    if (-2147483648 == 4294967295) {
        fmt.Println(v5, v17, v3)
        if (v3 != 3.14159) {
            if len(v18) > 0 {
                v18[0] = -880
            }
            var v19 bool = (824 != uint32(v2))
            _ = v19
        } else {
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            if ((v3 - -2.25) == v3) {
                _ = len(v18)
                for fi20 := 0; fi20 < 17; fi20++ {
                    v5 -= v5
                    if (!false) {
                        var v21 uint = 42
                        if (!false) {
                            v21 = 0
                        }
                        _ = v21
                    }
                    v6 = v6
                    _ = fi20
                }
                fmt.Println(v2, v2)
            } else {
                var v22 uint8 = 2
                _ = v22
            }
        }
    }
    v6 -= v6
    var v23 uint32 = 115
    for fi24 := 0; fi24 < 8; fi24++ {
        if (1 >= -968) {
            var v25 bool = (1e10 <= (0.0 / 1.5))
            fmt.Println(v1, v6)
            _ = v25
        } else {
            v23 = (v23 - v23)
        }
        _ = fi24
    }
    if (3.14159 <= (-2.25 + 0.0)) {
        v26 := GenS3{F0: (v2 << ((v17) & 63)), F1: uint64(v4), F2: v6}
        _ = v26
    }
    switch ("hello" + "hello") {
    case "foo bar":
        fmt.Println(v1, v6)
    default:
        v27 := GenS2{F0: (v4 - (^(v4)))}
        _ = v27
    }
    v5 += (v5 % ((v5) | 1))
    switch "foo bar" {
    case "":
        _ = len(v18)
    case "hello":
        if len(v18) > 0 {
            v18[0] = 42
        }
    default:
        switch (("\n\t" + "foo bar") + "") {
        case "a":
            var v28 float32 = (1e10 / 1.0)
            _ = v28
        case "\n\t":
            v5 -= v5
        default:
            fmt.Println(v23, v3, v2)
        }
    }
    v29 := GenS1{F0: false, F1: v2}
    fmt.Println(v4, v3)
    v2 = uint16(v4)
    v17 = 92
    switch "" {
    case "foo bar":
        v18 = append(v18, 0)
    case "hello":
        v17 -= v17
    case "a":
        v23 = 0
    default:
        v17 = uint(v23)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
