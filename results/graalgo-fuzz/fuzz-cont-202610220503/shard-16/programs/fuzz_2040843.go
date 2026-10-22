package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint }
func genFn0(p0 uint64, p1 uint16) uint8 {
    p0 = p0
    v7 := GenS1{F0: p0, F1: 1, F2: 2}
    if (false && false) {
        var v8 bool = true
        var v9 uint32 = 964
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: (p0 ^ p0), F1: 98, F2: 1}
        switch "\n\t" {
        case "":
            v11 := []int{1, -881, 390}
            _ = v11
        default:
            v12 := []int{9223372036854775807, -2147483648, 42}
            _ = v12
        }
        _ = v10
    }
    if false {
        p0 -= p0
        var v13 int8 = -103
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint = 159
        v16 := GenS1{F0: 4294967295, F1: 42, F2: (v15 ^ (v15 << ((v15) & 63)))}
        var v17 uint8 = 42
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v7
    return 245
}

func main() {
    v1 := []int{550, -2147483648, -1}
    v2 := []int{2147483647, 177, 0, 9223372036854775807}
    var v3 int = -700
    var v4 uint16 = 249
    v5 := []uint32{}
    var v6 float64 = (3.14159 - (1e10 - ((0.1 - (((1.5 + (float64((0.1 * 1e10)) + (1.5 * (1.5 + 3.14159)))) - (0.1 * float64(float32((0.1 / (0.0 + 3.14159)))))) / 1.0)) / (float64(1e10) - (1e10 * (float64(3.14159) - (3.14159 - 0.0)))))))
    var v18 uint = 823
    if (("hello" < (("" + ("" + (("a" + ("hello" + "foo bar")) + ("hello" + "")))) + "")) || false) {
        switch v3 {
        case 301:
            v3 += (v3 + v3)
        case 503:
            v18 = v18
        case 4294967295:
            if len(v1) > 0 {
                v1[0] = (v3 * 32767)
            }
        }
    } else {
        v19 := GenS1{F0: 1, F1: uint32(v4), F2: (^(v18))}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v19
    }
    var v20 bool = (-2147483648 < -2147483648)
    v6 -= 3.14159
    for fi21 := 0; fi21 < 17; fi21++ {
        v18 += v18
        for fi22 := 0; fi22 < 11; fi22++ {
            switch v3 {
            case 2147483647:
                v4 += 2
            }
            var v23 uint = v18
            fmt.Println(v4, v20)
            _ = fi22
            _ = v23
        }
        if (((v18 * v18) <= (v18 - 18446744073709551615)) || false) {
            v6 = float64(-2.25)
            switch "" {
            case "\n\t":
                if len(v5) > 0 {
                    v5[0] = 0
                }
            case "hello":
                _ = len(v2)
            default:
                if v20 {
                    if (1 != 0) {
                        var v24 int64 = -424
                        v5 = append(v5, 42)
                        var v25 int16 = 143
                        _ = v24
                        _ = v25
                    } else {
                        fmt.Println(v18, v6, v18)
                        fmt.Println(v6, v6, v20)
                        v18 = (v18 << (((v18 - v18)) & 63))
                    }
                    var v26 uint8 = 2
                    v4 = 32767
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 13; fi27++ {
                        fmt.Println(fi27, v18, v18)
                        fmt.Println(v20, v4, v18)
                        _ = fi27
                    }
                    for fi28 := 0; fi28 < 3; fi28++ {
                        if len(v5) > 0 {
                            v5[0] = 0
                        }
                        _ = fi28
                    }
                }
            }
        }
        _ = fi21
    }
    var v29 float32 = 0.1
    v6 += ((float64(0.0) - (0.0 - v6)) - 3.14159)
    v1 = append(v1, (v3 + v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
