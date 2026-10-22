package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int, p2 uint) uint16 {
    fmt.Println(p0, p2, p0)
    var v7 int8 = 0
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 uint16 = (p0 % ((p0) | 1))
            p2 -= 255
            var v11 uint8 = 42
            _ = fi9
            _ = v10
            _ = v11
        }
        if (84 == 255) {
            v12 := []int{4294967295, p1, 2147483647, int(p2)}
            _ = v12
        } else {
            var v13 uint64 = 834
            var v14 uint64 = (^(v13))
            for fi15 := 0; fi15 < 11; fi15++ {
                v14 += (^(v14))
                v16 := []int{(p1 - (p1 * p1)), (p1 % (((p1 | int(p2))) | 1)), (fi15 & (^(fi8))), (-(p1))}
                fmt.Println(fi8)
                fmt.Println(fi15, v14, fi8)
                v17 := GenS1{F0: (v13 << ((p2) & 63)), F1: 617, F2: -207}
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v13
            _ = v14
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}
func genFn1(p0 uint8, p1 int8) int16 {
    v18 := GenS2{F0: 618}
    if ((3.14159 < -2.25) && false) {
        v19 := GenS1{F0: 4294967295, F1: -1, F2: 9223372036854775807}
        if (42 == 4294967295) {
            p1 += (p1 - p1)
            p1 = -126
            fmt.Println(p0)
        } else {
            v20 := []int{-9223372036854775808, 1, -2147483648, 1}
            switch 42 {
            case 391:
                var v21 uint16 = 255
                _ = v21
            }
            _ = v20
        }
        v22 := []int{9223372036854775807, 32768, -9223372036854775808, 1}
        _ = v19
        _ = v22
    }
    var v23 float64 = 1e10
    _ = p0
    _ = p1
    _ = v18
    _ = v23
    return -737
}

func main() {
    v1 := []string{("\n\t" + "foo bar"), "foo bar"}
    var v2 uint64 = 246
    var v3 uint32 = uint32(v2)
    var v4 uint32 = (^(v3))
    var v5 string = ("" + "a")
    var v6 float32 = (3.14159 - (0.0 + float32(0.0)))
    switch ("" + v5) {
    case "hello":
        var v24 int16 = 1
        _ = v24
    default:
        var v25 float64 = 3.14159
        _ = v25
    }
    fmt.Println(v5, v4, v3)
    if false {
        v26 := GenS1{F0: (^(v2)), F1: 197, F2: -864}
        if (-118 < -1) {
            if (!(((2147483647 > -2147483648) && true) && (!(255 <= 1)))) {
                var v27 int16 = genFn1(42, -20)
                v2 -= v2
                switch "" {
                case "\n\t":
                    var v28 bool = (0 < 60)
                    _ = v28
                case "":
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                case "a":
                    var v29 int = 1
                    _ = v29
                default:
                    fmt.Println(v2, v27, v3)
                }
                _ = v27
            } else {
                v6 = ((v6 * v6) + (0.1 - 0.0))
            }
        }
        var v30 bool = true
        _ = v26
        _ = v30
    }
    var v31 float64 = -2.25
    var v32 bool = (94 <= 128)
    switch (v5 + v5) {
    case "":
        v31 += v31
    case "hello":
        for fi33 := 0; fi33 < 15; fi33++ {
            var v34 string = v5
            _ = fi33
            _ = v34
        }
    }
    if len(v1) > 0 {
        v1[0] = v5
    }
    fmt.Println(v4, v32, v31)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
