package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 string }
type GenS2 struct { F0 uint32; F1 uint32; F2 int64 }
func genFn0(p0 uint, p1 uint16) string {
    v7 := []int{128, 32768, 32768}
    if len(v7) > 0 {
        v7[0] = 856
    }
    if len(v7) > 0 {
        v7[0] = -68
    }
    if (!false) {
        fmt.Println(p1, p0)
    } else {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 bool = false
            _ = fi8
            _ = v9
        }
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    _ = v7
    return "foo bar"
}

func main() {
    var v1 int8 = 42
    var v2 float64 = 0.1
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    v5 := make(map[string]int)
    var v6 float64 = ((((((v2 * float64(0.0)) - -2.25) + v2) / 1.0) * v2) + ((((1.5 - (v2 - (v2 / -2.25))) - v2) - float64(float32((v2 - ((v2 * (float64(3.14159) * 1e10)) + float64(1.5)))))) + v2))
    v10 := GenS1{F0: 719, F1: true, F2: (("" + "hello") + "hello")}
    switch 790 {
    case -1:
        v11 := GenS1{F0: 2147483647, F1: (!false), F2: "\n\t"}
        _ = v11
    case 0:
        if (724 == 42) {
            switch 109 {
            case 9223372036854775807:
                _ = len(v5)
            case 2147483647:
                v1 = v1
            case 261:
                v2 = v6
            }
            v2 -= (3.14159 + (float64(1.5) * (v2 + 3.14159)))
            v2 += 3.14159
        } else {
            var v12 bool = true
            switch ("a" + "") {
            case "foo bar":
                for fi13 := 0; fi13 < 9; fi13++ {
                    var v14 int = fi13
                    _ = fi13
                    _ = v14
                }
            case "":
                if true {
                    fmt.Println(v2)
                    for fi15 := 0; fi15 < 4; fi15++ {
                        fmt.Println(v1)
                        var v16 int32 = -1
                        if (!((v1 ^ v1) < -32)) {
                            v17 := GenS2{F0: 2, F1: 65535, F2: 42}
                            var v18 int = 152
                            for fi19 := 0; fi19 < 18; fi19++ {
                                v18 = -82
                                _ = fi19
                            }
                            _ = v17
                            _ = v18
                        } else {
                            _ = len(v4)
                            v20 := GenS2{F0: 14, F1: 789, F2: 688}
                            var v21 int16 = -1
                            _ = v20
                            _ = v21
                        }
                        _ = fi15
                        _ = v16
                    }
                    v6 = 3.14159
                } else {
                    fmt.Println(v2, v1)
                    _ = len(v3)
                    v2 += 0.0
                }
            default:
                var v22 bool = ((2147483647 < 0) || v12)
                _ = v22
            }
            _ = v12
        }
    case 4294967295:
        fmt.Println(v1)
    default:
        var v23 uint = 18446744073709551615
        _ = v23
    }
    var v24 uint8 = 128
    v25 := GenS2{F0: 292, F1: 759, F2: 157}
    fmt.Println(v1)
    switch "hello" {
    case "a":
        v26 := []int{32767, 128, -9223372036854775808, 118}
        _ = v26
    case "\n\t":
        var v27 int16 = -650
        _ = v27
    }
    var v28 bool = (-1 != (v1 << ((42) & 63)))
    v28 = v28
    var v29 bool = (1 <= 560)
    delete(v3, 128)
    var v30 string = (("a" + ("\n\t" + ("a" + "\n\t"))) + "a")
    v31 := []int{-638, -864}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
