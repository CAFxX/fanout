package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int32; F2 int }
func genFn0(p0 uint32, p1 uint32) uint32 {
    var v7 int = 32767
    v8 := []int{(v7 % ((516) | 1)), 2147483647}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint64) int16 {
    for fi9 := 0; fi9 < 11; fi9++ {
        if true {
            v10 := []int{(fi9 | fi9), fi9, (fi9 | fi9)}
            p0 -= p0
            if len(v10) > 0 {
                v10[0] = fi9
            }
            _ = v10
        }
        switch "\n\t" {
        case "":
            var v11 int16 = 579
            _ = v11
        case "foo bar":
            v12 := []int{fi9, (fi9 >> ((65535) & 63)), (^(fi9)), (fi9 << ((2) & 63))}
            _ = v12
        }
        _ = fi9
    }
    p0 -= p0
    p0 = p0
    _ = p0
    return 1000
}

func main() {
    var v1 uint64 = 42
    v2 := []string{"", "\n\t", "", "hello"}
    var v3 uint8 = 0
    var v4 uint8 = 15
    var v5 uint8 = (v3 / (((^(v4))) | 1))
    var v6 float32 = (float32(float64((1.5 - ((1e10 * 1e10) / 0.1)))) / -2.25)
    fmt.Println(v6, v4)
    v3 -= 250
    fmt.Println(v3)
    fmt.Println(v5, v5, v5)
    switch "" {
    case "foo bar":
        if ((0.1 + (1.5 + float64(float32((0.1 + ((0.0 - (0.0 * 3.14159)) / 1.5)))))) > ((float64(1.5) * 1.5) + ((((1e10 * 1.5) * -2.25) * float64(0.0)) + float64(((1.5 / v6) * v6))))) {
            switch ((("\n\t" + "") + "hello") + "a") {
            case "hello":
                if (!(!((true || true) || (-2.25 <= (1e10 * 0.0))))) {
                    for fi13 := 0; fi13 < 6; fi13++ {
                        v6 = -2.25
                        v14 := GenS1{F0: fi13}
                        _ = fi13
                        _ = v14
                    }
                    fmt.Println(v4, v1)
                    for fi15 := 0; fi15 < 14; fi15++ {
                        for fi16 := 0; fi16 < 18; fi16++ {
                            if (-128 < 1) {
                                fmt.Println(fi15, v1, fi16)
                            }
                            _ = fi16
                        }
                        _ = fi15
                    }
                } else {
                    var v17 int = 32767
                    v18 := GenS1{F0: (v17 + v17)}
                    v19 := GenS1{F0: (v17 % ((v17) | 1))}
                    _ = v17
                    _ = v18
                    _ = v19
                }
            default:
                v5 = 1
            }
            var v20 int = 32767
            _ = v20
        } else {
            for fi21 := 0; fi21 < 15; fi21++ {
                v4 -= (v3 * 255)
                _ = fi21
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v22 int8 = -66
            _ = v22
        }
    case "":
        if false {
            var v23 uint8 = 255
            fmt.Println(v4, v23, v6)
            if (v5 == (v5 * (v3 - v4))) {
                var v24 float32 = 0.1
                _ = v24
            } else {
                var v25 int32 = 42
                _ = v25
            }
            _ = v23
        } else {
            var v26 float64 = (-2.25 + ((0.1 / (0.0 + ((1.5 + float64(v6)) / 1e10))) + 1.5))
            v27 := GenS2{F0: 9223372036854775807, F1: -1, F2: 0}
            _ = v26
            _ = v27
        }
    case "a":
        v6 = v6
    default:
        v1 += v1
    }
    var v28 int32 = -1
    var v29 uint8 = v5
    v30 := GenS2{F0: 9223372036854775807, F1: -479, F2: -439}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
