package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int32 }
func genFn0(p0 int) uint8 {
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{128, p0, -9223372036854775808}
        v8 := GenS1{F0: 42}
        if false {
            _ = len(v7)
            fmt.Println(fi6, fi6)
        } else {
            for fi9 := 0; fi9 < 14; fi9++ {
                if false {
                    var v10 float32 = 1.5
                    fmt.Println(fi9, v10, p0)
                    var v11 uint = 922
                    _ = v10
                    _ = v11
                }
                fmt.Println(fi9, fi9)
                if (65535 == 539) {
                    p0 -= (fi6 + (p0 % (((p0 / ((32768) | 1))) | 1)))
                    _ = len(v7)
                }
                p0 += fi6
                v12 := GenS2{F0: -1}
                _ = fi9
                _ = v12
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    p0 += (^(p0))
    v13 := []int{(p0 >> ((1) & 63)), (p0 >> ((2) & 63)), (p0 - 32767), p0}
    if (false || (false || (false || false))) {
        v14 := GenS2{F0: -475}
        _ = v14
    } else {
        v15 := GenS1{F0: 1}
        _ = v15
    }
    _ = p0
    _ = v13
    return 64
}

func main() {
    var v1 int16 = -183
    v2 := make(map[uint32]uint32)
    var v3 uint = 7
    var v4 int16 = (v1 - -827)
    var v5 uint = 18446744073709551615
    v4 = v1
    fmt.Println(v4, v5, v1)
    v5 = v3
    v4 = (v4 >> ((0) & 63))
    fmt.Println(v3, v1, v4)
    v3 = v5
    v16 := GenS3{F0: -740, F1: -2147483648}
    if false {
        var v17 int64 = 1
        _ = v17
    }
    switch "hello" {
    case "a":
        fmt.Println(v1)
    case "\n\t":
        v3 = 18446744073709551615
    default:
        v3 = (v5 << (((v5 - v5)) & 63))
    }
    if (v4 <= v1) {
        v2[2147483647] = 4294967295
    } else {
        var v18 string = "hello"
        v5 -= v3
        if true {
            var v19 int64 = 363
            var v20 float32 = ((((-2.25 - 1e10) + 1e10) - (3.14159 + 3.14159)) * float32(0.0))
            _ = v19
            _ = v20
        }
        _ = v18
    }
    if true {
        var v21 uint = (v5 | v5)
        if (float32(((1.5 - 0.0) * float64(0.1))) <= (-2.25 / ((float32(0.1) * 1e10) - (1.5 * (0.0 + -2.25))))) {
            v5 = (v5 + 839)
            if (v4 == (v4 << (((v3 + (v3 * (v3 % ((v3) | 1))))) & 63))) {
                for fi22 := 0; fi22 < 14; fi22++ {
                    v21 = (v5 << (((v3 + 949)) & 63))
                    _ = fi22
                }
                var v23 uint64 = 18446744073709551615
                _ = v23
            }
        } else {
            if (-1 == -9223372036854775808) {
                for fi24 := 0; fi24 < 13; fi24++ {
                    if true {
                        v3 += 0
                        if false {
                            v4 = (v1 | 0)
                            var v25 int8 = -1
                            v1 += (v1 & v4)
                            _ = v25
                        }
                    }
                    _ = fi24
                }
                if false {
                    fmt.Println(v4, v4, v21)
                }
                var v26 float64 = (float64(1.5) - 0.0)
                _ = v26
            }
        }
        _ = v21
    } else {
        v3 -= v3
        v27 := GenS2{F0: -126}
        switch 742 {
        case 406:
            for fi28 := 0; fi28 < 7; fi28++ {
                v3 += (^(v3))
                _ = fi28
            }
        case 1:
            v29 := GenS2{F0: 478}
            _ = v29
        }
        _ = v27
    }
    v30 := GenS3{F0: 0, F1: 1}
    var v31 uint = 356
    var v32 int = 9223372036854775807
    v4 += v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
