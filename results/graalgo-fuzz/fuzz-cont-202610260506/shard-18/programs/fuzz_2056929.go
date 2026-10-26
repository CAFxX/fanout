package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int8 }
type GenS2 struct { F0 int; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint; F1 string; F2 int }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int8; F1 bool }
func genFn0(p0 int) uint64 {
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 int = p0
        switch "\n\t" {
        case "":
            v6 := []int{}
            _ = v6
        }
        _ = fi4
        _ = v5
    }
    v7 := GenS1{F0: 484, F1: 0, F2: 45}
    var v8 int = 856
    switch v8 {
    case 4294967295:
        v8 = p0
    case 2147483647:
        var v9 string = "\n\t"
        _ = v9
    case 128:
        fmt.Println(p0, v8)
    default:
        v10 := GenS2{F0: (p0 & 2147483647), F1: 18446744073709551615, F2: 0}
        _ = v10
    }
    var v11 string = "\n\t"
    _ = p0
    _ = v11
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint64 = 653
    v2 := make(map[uint32]bool)
    var v3 int64 = -1
    v1 = (v1 / ((v1) | 1))
    v12 := v2[2147483647]
    switch 4294967295 {
    case -1:
        v1 -= (v1 * v1)
    case -9223372036854775808:
        v3 = v3
    case 213:
        fmt.Println(v3)
    }
    var v13 int32 = 0
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := []int{fi14, fi14, (fi14 * fi14)}
        v16 := v2[617]
        if (!false) {
            if ((!(!(!(!(42 == -26))))) || v12) {
                switch "foo bar" {
                case "\n\t":
                    delete(v2, 1)
                case "hello":
                    _ = len(v15)
                case "":
                    fmt.Println(v12, v13, v1)
                }
            }
            fmt.Println(v13, v12, v16)
            if true {
                if v16 {
                    fmt.Println(v3)
                } else {
                    v17 := GenS3{F0: 65535, F1: "a", F2: fi14}
                    switch fi14 {
                    case -9223372036854775808:
                        fmt.Println(v13, fi14)
                    case 0:
                        fmt.Println(v12, v16, fi14)
                    }
                    _ = v17
                }
            } else {
                if ((v13 * -2147483648) > 0) {
                    fmt.Println(v1)
                    switch (fi14 / (((fi14 / ((0) | 1))) | 1)) {
                    case 127:
                        if v16 {
                            fmt.Println(fi14, v1, v16)
                            var v18 float32 = (((0.1 + 1.5) + (0.1 * (0.0 * (float32(float64(1e10)) / ((((-2.25 + 0.0) * 1e10) / 0.1) - -2.25))))) + (-2.25 + 1.5))
                            var v19 int32 = (v13 - (v13 >> ((401) & 63)))
                            _ = v18
                            _ = v19
                        } else {
                            switch fi14 {
                            case -22:
                                fmt.Println(v12, v16)
                            case 127:
                                switch "hello" {
                                case "hello":
                                    _ = len(v15)
                                case "\n\t":
                                    var v20 int64 = (-(v3))
                                    _ = v20
                                default:
                                    if v16 {
                                        if true {
                                            v1 = 42
                                            var v21 uint16 = 65535
                                            _ = v21
                                        } else {
                                            v12 = (2 != 2)
                                        }
                                        var v22 string = "foo bar"
                                        for fi23 := 0; fi23 < 10; fi23++ {
                                            var v24 int64 = v3
                                            var v25 string = v22
                                            _ = fi23
                                            _ = v24
                                            _ = v25
                                        }
                                        _ = v22
                                    }
                                }
                            }
                            v3 += (v3 / ((-9223372036854775808) | 1))
                            v26 := GenS3{F0: 65535, F1: ("a" + "a"), F2: fi14}
                            _ = v26
                        }
                    case 4294967295:
                        v27 := GenS4{F0: 2, F1: v12}
                        _ = v27
                    default:
                        v28 := GenS1{F0: 448, F1: 1, F2: 0}
                        _ = v28
                    }
                    v15 = append(v15, -1)
                }
                var v29 float64 = (float64(1.5) * 0.0)
                v16 = ((v1 | (v1 ^ (v1 / ((4294967295) | 1)))) <= (v1 << ((0) & 63)))
                _ = v29
            }
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    var v30 uint16 = 255
    switch "" {
    case "a":
        v2[42] = false
    default:
        var v31 int64 = 42
        _ = v31
    }
    var v32 int16 = -800
    var v33 uint64 = genFn0(-49)
    v34 := GenS5{F0: int8(v30), F1: (v13 > v13)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
}
