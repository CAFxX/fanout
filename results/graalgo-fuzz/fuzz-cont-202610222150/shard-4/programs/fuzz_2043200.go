package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint64; F1 int }
func genFn0(p0 string, p1 int16, p2 uint64) uint32 {
    switch 1 {
    case 32768:
        v4 := []int{32768, 970, 0, 2147483647}
        _ = v4
    }
    p2 -= 2
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1() uint8 {
    var v5 uint8 = 2
    fmt.Println(v5)
    _ = v5
    return (v5 ^ 127)
}

func main() {
    var v1 int16 = 315
    v2 := make(map[uint32]uint32)
    var v3 float64 = (-2.25 + 1.5)
    delete(v2, 2)
    var v6 int = -736
    v6 = (^(v6))
    var v7 int64 = 1
    if true {
        for fi8 := 0; fi8 < 3; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                v10 := []int{-639, v6, 32767, (fi9 | 919)}
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v3, v6, v7)
                default:
                    switch "\n\t" {
                    case "a":
                        fmt.Println(v7)
                    default:
                        if (42 == 18446744073709551615) {
                            v6 = fi9
                            if false {
                                v11 := v2[2147483648]
                                fmt.Println(v6, v1, fi8)
                                if true {
                                    var v12 int32 = 2147483647
                                    fmt.Println(fi8, fi9, v11)
                                    v1 += (v1 | v1)
                                    _ = v12
                                } else {
                                    var v13 int64 = v7
                                    v1 += v1
                                    _ = v13
                                }
                                _ = v11
                            } else {
                                if len(v10) > 0 {
                                    v10[0] = fi9
                                }
                            }
                        } else {
                            v1 += (v1 | 1000)
                            var v14 int32 = 1
                            _ = v14
                        }
                    }
                    if (1e10 > 3.14159) {
                        v10 = append(v10, (^(fi9)))
                        var v15 int8 = -128
                        _ = v15
                    } else {
                        var v16 int32 = -2147483648
                        _ = v16
                    }
                }
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    } else {
        for fi17 := 0; fi17 < 5; fi17++ {
            _ = len(v2)
            v18 := GenS1{F0: 742, F1: 9}
            for fi19 := 0; fi19 < 19; fi19++ {
                var v20 float64 = v3
                _ = fi19
                _ = v20
            }
            _ = fi17
            _ = v18
        }
        var v21 uint = 18446744073709551615
        v22 := []int{v6, (v6 * 32768), (v6 % ((v6) | 1)), (v6 * (v6 >> ((v21) & 63)))}
        _ = v21
        _ = v22
    }
    v1 += (v1 >> ((4294967295) & 63))
    v23 := GenS1{F0: 42, F1: 1}
    if (!true) {
        v1 -= v1
        if false {
            var v24 uint = 972
            var v25 float64 = float64(0.0)
            _ = v24
            _ = v25
        } else {
            if true {
                v26 := GenS2{F0: 18446744073709551615, F1: v6}
                _ = v26
            } else {
                var v27 uint32 = 2147483648
                _ = v27
            }
        }
    }
    switch v6 {
    case 709:
        v6 -= 9223372036854775807
    case -9223372036854775808:
        v7 = 232
    case 1:
        switch 128 {
        case -9223372036854775808:
            if (281 > 32767) {
                v6 = v6
                v1 -= v1
                v2[2147483647] = 65535
            } else {
                switch ("foo bar" + "foo bar") {
                case "hello":
                    v7 = v7
                }
                var v28 int16 = (-(v1))
                _ = v28
            }
        case 1:
            v3 = v3
        }
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        fmt.Println(v6)
        _ = fi29
    }
    v30 := GenS1{F0: -2147483648, F1: 1}
    v31 := GenS2{F0: 2, F1: (^(v6))}
    switch -966 {
    case -1:
        switch "foo bar" {
        case "a":
            delete(v2, 2147483648)
        }
    default:
        v32 := GenS2{F0: 2, F1: v6}
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v6
    _ = v7
}
