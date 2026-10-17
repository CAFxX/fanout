package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int32, p1 int32, p2 uint8) string {
    if ((p0 ^ (p1 + -2147483648)) < p0) {
        p1 = (-(p1))
    } else {
        if (!(9223372036854775807 < -947)) {
            p0 = -1
            switch "a" {
            case "\n\t":
                if (1 == -652) {
                    if ((false || (-9223372036854775808 <= 42)) || false) {
                        fmt.Println(p0, p2, p2)
                        for fi6 := 0; fi6 < 20; fi6++ {
                            for fi7 := 0; fi7 < 11; fi7++ {
                                p1 = (^(p0))
                                _ = fi7
                            }
                            for fi8 := 0; fi8 < 3; fi8++ {
                                v9 := []int{(fi8 ^ fi8)}
                                _ = fi8
                                _ = v9
                            }
                            _ = fi6
                        }
                    } else {
                        for fi10 := 0; fi10 < 13; fi10++ {
                            v11 := []int{0, (fi10 >> ((4294967295) & 63))}
                            for fi12 := 0; fi12 < 2; fi12++ {
                                var v13 float32 = 1e10
                                v13 = (v13 + (float32((float64(((0.0 + (3.14159 - v13)) * (v13 * v13))) + -2.25)) * ((v13 - (((1.5 * 0.0) * (0.1 + v13)) + v13)) + (v13 / v13))))
                                _ = fi12
                                _ = v13
                            }
                            var v14 uint32 = 771
                            _ = fi10
                            _ = v11
                            _ = v14
                        }
                        p0 -= p1
                    }
                } else {
                    fmt.Println(p0, p1)
                }
            case "hello":
                p0 -= -2147483648
            case "foo bar":
                v15 := GenS1{F0: 889}
                _ = v15
            }
        } else {
            fmt.Println(p0)
        }
    }
    v16 := GenS2{F0: -1, F1: (p1 + p1)}
    v17 := []int{-414, 512, 9223372036854775807}
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return "a"
}

func main() {
    v1 := []int{128, 642, 2147483647, 4294967295}
    v2 := make(map[string]uint32)
    var v3 uint8 = 255
    v4 := []uint32{uint32(v3), 0, 4294967295, 4294967295}
    var v5 float32 = -2.25
    if len(v4) > 0 {
        v4[0] = 168
    }
    v18 := GenS3{F0: true, F1: 100}
    v19 := GenS4{F0: v3}
    var v20 uint16 = 516
    var v21 int16 = -1
    var v22 int = 128
    v3 = (^(v3))
    fmt.Println(v21)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
