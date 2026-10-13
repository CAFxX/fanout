package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint64; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
func genFn0(p0 int32, p1 uint8) int {
    p1 = p1
    p0 = -2147483648
    p1 = (p1 % ((130) | 1))
    if false {
        fmt.Println(p1)
        v7 := []int{319, 4294967295, -85, -1}
        if len(v7) > 0 {
            v7[0] = -9223372036854775808
        }
        _ = v7
    }
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    var v4 uint64 = 4294967295
    var v5 int8 = -55
    var v6 int16 = 1
    switch 128 {
    case -1:
        var v8 int32 = 42
        _ = v8
    case 42:
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := GenS1{F0: 815, F1: 255}
            if false {
                var v11 float64 = (3.14159 * 1e10)
                _ = v11
            } else {
                if (255 > 59) {
                    var v12 int16 = (v6 << ((4294967295) & 63))
                    _ = v12
                }
                v13 := []int{(fi9 % (((fi9 | fi9)) | 1)), fi9}
                v2[(fi9 + (fi9 * (fi9 | fi9)))] = 370
                _ = v13
            }
            if (65535 >= 919) {
                v14 := GenS2{F0: v6, F1: (v4 % ((366) | 1)), F2: ""}
                _ = v14
            }
            _ = fi9
            _ = v10
        }
    default:
        if true {
            if (false || (266 >= v6)) {
                v6 += -514
                switch "a" {
                case "a":
                    v15 := []int{221, 9223372036854775807}
                    _ = v15
                default:
                    for fi16 := 0; fi16 < 12; fi16++ {
                        v17 := []int{}
                        _ = fi16
                        _ = v17
                    }
                }
            }
            if (988 <= 18446744073709551615) {
                v4 += v4
            } else {
                for fi18 := 0; fi18 < 13; fi18++ {
                    var v19 uint8 = 0
                    v4 = v4
                    _ = fi18
                    _ = v19
                }
            }
            v20 := GenS1{F0: 18446744073709551615, F1: 127}
            _ = v20
        } else {
            for fi21 := 0; fi21 < 10; fi21++ {
                v22 := []int{(fi21 << ((0) & 63)), fi21, genFn0(-2147483648, 128), fi21}
                _ = fi21
                _ = v22
            }
            v6 = (v6 + v6)
            v23 := v3[32767]
            _ = v23
        }
    }
    v6 = (v6 | (v6 ^ v6))
    if ((-2.25 <= (1e10 - -2.25)) || false) {
        for fi24 := 0; fi24 < 16; fi24++ {
            var v25 uint32 = 798
            v26 := []int{(fi24 & (fi24 + (fi24 & (fi24 / (((fi24 << ((4294967295) & 63))) | 1))))), genFn0(42, 2), (fi24 % ((-753) | 1)), genFn0(-470, 127)}
            v27 := GenS3{F0: (v5 + (v5 << ((536) & 63)))}
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
    }
    if false {
        for fi28 := 0; fi28 < 10; fi28++ {
            v6 -= 708
            v5 += (v5 * (-(v5)))
            _ = fi28
        }
        v29 := GenS1{F0: 749, F1: 207}
        v5 = 0
        _ = v29
    }
    v4 = (v4 & v4)
    switch ("\n\t" + (("hello" + "a") + "hello")) {
    case "\n\t":
        if (182 >= 115) {
            v5 -= (v5 % ((v5) | 1))
            if true {
                v6 = (^(v6))
                fmt.Println(v5, v6)
            } else {
                v6 -= 0
            }
            fmt.Println(v5)
        }
    case "a":
        v4 -= v4
    default:
        v30 := []int{520, 32768, -1, 481}
        _ = v30
    }
    fmt.Println(v4, v4)
    var v31 float32 = 0.1
    v31 += (v31 * (v31 * v31))
    v31 = v31
    fmt.Println(v5, v5, v4)
    var v32 uint32 = 1
    v33 := GenS4{F0: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
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
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
