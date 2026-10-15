package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint }
func genFn0(p0 int, p1 int32, p2 int8) int {
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(p1, p0, p2)
        p2 -= (p2 ^ 42)
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    p0 = 127
    _ = p0
    _ = p1
    _ = p2
    return (p0 * (p0 & (p0 & p0)))
}
func genFn1(p0 uint64, p1 uint64, p2 uint) int16 {
    var v9 float64 = 0.0
    v10 := GenS1{F0: -32768, F1: 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 0
}

func main() {
    v1 := []int{9223372036854775807, 66, 9223372036854775807, 758}
    var v2 uint16 = 65535
    var v3 uint8 = 127
    var v4 float64 = 1e10
    v5 := make(map[int]string)
    var v6 float32 = -2.25
    v3 = (v3 / ((v3) | 1))
    v3 -= (v3 - 1)
    var v11 uint = 1
    v2 += (v2 * v2)
    for fi12 := 0; fi12 < 5; fi12++ {
        v2 -= (v2 | (v2 / (((v2 << (((v11 / ((0) | 1))) & 63))) | 1)))
        v1 = append(v1, (fi12 - int(v11)))
        _ = fi12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        for fi14 := 0; fi14 < 14; fi14++ {
            v6 += v6
            switch "" {
            case "hello":
                switch (fi14 ^ fi14) {
                case 42:
                    switch genFn0(genFn0((fi13 | fi14), -1, 20), 42, 127) {
                    case 9223372036854775807:
                        var v15 uint8 = (v3 | (v3 % ((v3) | 1)))
                        _ = v15
                    case 32767:
                        v16 := GenS2{F0: v11}
                        _ = v16
                    }
                case -219:
                    switch "a" {
                    case "\n\t":
                        fmt.Println(v3)
                    case "":
                        var v17 int8 = 100
                        _ = v17
                    }
                case 32767:
                    if true {
                        if true {
                            v4 = (float64((v6 + (((v6 * (0.0 + v6)) - 1e10) - float32(v4)))) - 1e10)
                            if true {
                                v11 -= (^(v11))
                                if (v11 >= v11) {
                                    if (!(true || true)) {
                                        _ = len(v5)
                                        v6 = v6
                                        v6 = (-2.25 / float32((v4 / 1e10)))
                                        _ = len(v5)
                                    }
                                    fmt.Println(v4)
                                } else {
                                    v11 = v11
                                    if true {
                                        _ = len(v5)
                                    } else {
                                        var v18 uint = (v11 * (v11 | 18446744073709551615))
                                        var v19 int8 = 0
                                        _ = v18
                                        _ = v19
                                    }
                                    var v20 uint = v11
                                    _ = v20
                                }
                            } else {
                                var v21 uint8 = (^(v3))
                                var v22 float64 = v4
                                v5[fi14] = "foo bar"
                                _ = v21
                                _ = v22
                            }
                            switch genFn0(fi14, 346, 42) {
                            case 4294967295:
                                fmt.Println(v11, fi13, v4)
                            default:
                                fmt.Println(v11, fi13)
                            }
                            v4 = 0.1
                            var v23 int16 = 0
                            _ = v23
                        } else {
                            v2 -= 1
                        }
                        _ = len(v5)
                    }
                    v11 = 4294967295
                }
            }
            _ = fi14
        }
        var v24 int8 = 87
        v11 = (v11 * 4294967295)
        _ = fi13
        _ = v24
    }
    var v25 float32 = v6
    v6 -= v25
    if (true || (32767 < int16(v11))) {
        var v26 uint8 = (^(v3))
        var v27 int = 1
        v6 -= v6
        _ = v26
        _ = v27
    } else {
        v6 = 1.5
    }
    v11 += (^(v11))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
