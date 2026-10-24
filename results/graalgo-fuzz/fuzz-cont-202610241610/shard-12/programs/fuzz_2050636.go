package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint8; F2 int64 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 bool; F1 uint }
func genFn0(p0 int16, p1 uint) uint {
    p0 = p0
    v4 := GenS1{F0: -48}
    v5 := []int{9223372036854775807, -270}
    v5 = append(v5, -832)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (p1 + p1)
}
func genFn1(p0 int, p1 uint16) int64 {
    v6 := []int{(p0 & 42)}
    var v7 uint32 = 2
    var v8 float32 = -2.25
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]bool)
    var v3 uint16 = 1
    var v9 uint16 = (^(v3))
    v10 := []int{-1, 9223372036854775807, 1, -1}
    fmt.Println(v3, v1, v9)
    switch ("a" + "\n\t") {
    case "":
        v1 = v1
    case "\n\t":
        v11 := GenS2{F0: -1, F1: 255, F2: 42}
        _ = v11
    case "hello":
        switch "foo bar" {
        case "hello":
            v12 := GenS3{F0: 4294967295, F1: 412}
            _ = v12
        case "\n\t":
            var v13 int64 = genFn1(-1, 167)
            _ = v13
        case "a":
            for fi14 := 0; fi14 < 7; fi14++ {
                switch "" {
                case "\n\t":
                    var v15 uint = 18446744073709551615
                    _ = v15
                case "hello":
                    if (148 >= 42) {
                        var v16 uint64 = 4294967295
                        _ = v16
                    }
                }
                if (!(!v1)) {
                    for fi17 := 0; fi17 < 11; fi17++ {
                        v18 := GenS4{F0: 255, F1: 32767}
                        var v19 uint64 = 1
                        if v1 {
                            var v20 int32 = 590
                            if len(v10) > 0 {
                                v10[0] = fi17
                            }
                            var v21 int32 = 1
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            v22 := GenS5{F0: v1, F1: 255}
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    _ = len(v10)
                } else {
                    v9 += (v9 >> ((0) & 63))
                }
                v3 -= 32767
                _ = fi14
            }
        }
    default:
        var v23 int8 = 0
        _ = v23
    }
    var v24 bool = (genFn0(1, 919) >= 491)
    v10 = append(v10, 117)
    v3 = 2
    var v25 int = -1
    var v26 uint = 18446744073709551615
    if (((((1.5 - ((-2.25 - (1e10 / 1e10)) - 0.0)) + 0.1) - ((3.14159 * 3.14159) / 1.5)) - 3.14159) >= (3.14159 * float64(0.0))) {
        switch 32768 {
        case 4294967295:
            delete(v2, 15)
        default:
            switch 42 {
            case 32767:
                v24 = ((!(42 != 65535)) && (v1 || (!false)))
            case 1:
                v27 := GenS4{F0: (v9 % (((v3 >> ((genFn0(1000, v26)) & 63))) | 1)), F1: 1}
                _ = v27
            default:
                var v28 int32 = 42
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 5; fi29++ {
            v24 = (2 >= 42)
            _ = fi29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v9
}
