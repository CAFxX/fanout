package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint16, p1 uint8, p2 int16) string {
    if false {
        for fi6 := 0; fi6 < 12; fi6++ {
            if false {
                v7 := []int{(fi6 ^ (fi6 / (((fi6 | (^(fi6)))) | 1))), fi6, (fi6 % ((-752) | 1))}
                p2 = p2
                _ = v7
            }
            for fi8 := 0; fi8 < 2; fi8++ {
                fmt.Println(p1, fi6)
                v9 := GenS1{F0: -128}
                if ((false || (!((0 > p1) && (42 > 42)))) && false) {
                    v10 := []int{(fi8 << ((255) & 63))}
                    if len(v10) > 0 {
                        v10[0] = fi8
                    }
                    if true {
                        var v11 uint = uint(p1)
                        var v12 int = fi8
                        switch "" {
                        case "a":
                            v10 = append(v10, v12)
                        }
                        if len(v10) > 0 {
                            v10[0] = v12
                        }
                        fmt.Println(fi8, fi8)
                        v13 := GenS1{F0: 100}
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    _ = v10
                }
                _ = fi8
                _ = v9
            }
            if false {
                p0 -= p0
                var v14 int8 = 72
                fmt.Println(p2)
                _ = v14
            }
            _ = fi6
        }
        v15 := []int{-867}
        _ = v15
    }
    var v16 float64 = 3.14159
    if (0 == 100) {
        v17 := []int{}
        for fi18 := 0; fi18 < 16; fi18++ {
            v16 = (3.14159 + (v16 - 1e10))
            v19 := GenS1{F0: 24}
            fmt.Println(p1, p0)
            _ = fi18
            _ = v19
        }
        switch "foo bar" {
        case "\n\t":
            if ((0 == uint64(p0)) || true) {
                switch 0 {
                case 820:
                    var v20 uint8 = p1
                    _ = v20
                case -2147483648:
                    p2 += 953
                }
                p2 += (p2 | (p2 - 42))
            } else {
                p1 = p1
            }
        case "foo bar":
            v16 -= (v16 / v16)
        case "hello":
            var v21 uint64 = 42
            _ = v21
        default:
            v16 += (0.0 / float64(float32(float64(float32((0.1 - v16))))))
        }
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return (("foo bar" + "") + ("" + "foo bar"))
}
func genFn1(p0 uint16) bool {
    v22 := GenS1{F0: 0}
    p0 += (p0 % ((0) | 1))
    _ = p0
    _ = v22
    return (false && (!true))
}

func main() {
    v1 := []string{}
    var v2 uint16 = 518
    var v3 int16 = 1000
    var v4 float32 = (-2.25 / (((0.0 - 0.1) * (3.14159 - 3.14159)) + 0.1))
    var v5 uint32 = 2147483648
    for fi23 := 0; fi23 < 9; fi23++ {
        for fi24 := 0; fi24 < 20; fi24++ {
            if (false && true) {
                if false {
                    v5 -= v5
                    var v25 uint = 4294967295
                    _ = v25
                }
                var v26 float64 = 3.14159
                var v27 int16 = (v3 | v3)
                _ = v26
                _ = v27
            } else {
                var v28 uint8 = 208
                var v29 uint32 = (v5 ^ v5)
                _ = v28
                _ = v29
            }
            _ = fi24
        }
        _ = fi23
    }
    v4 += (v4 + 3.14159)
    if (v5 <= 2) {
        v4 = ((v4 * float32(-2.25)) * (float32((3.14159 * (float64((((v4 - (v4 / v4)) * v4) - (0.0 / (-2.25 - (1e10 * 0.1))))) / -2.25))) / v4))
        for fi30 := 0; fi30 < 11; fi30++ {
            v3 = (v3 << ((2) & 63))
            v3 = v3
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi30
        }
    }
    var v31 int64 = 0
    var v32 int16 = (v3 << ((4294967295) & 63))
    v33 := GenS1{F0: 0}
    var v34 int32 = 2147483647
    var v35 uint8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
