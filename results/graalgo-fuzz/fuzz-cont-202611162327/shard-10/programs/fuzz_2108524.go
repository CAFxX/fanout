package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int64; F1 bool }
func genFn0(p0 uint8) int16 {
    fmt.Println(p0)
    var v7 uint8 = 128
    for fi8 := 0; fi8 < 7; fi8++ {
        v7 = 42
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    var v10 int8 = 0
    v11 := GenS1{F0: -9223372036854775808, F1: v10}
    _ = p0
    _ = v10
    _ = v11
    _ = v7
    return 1000
}
func genFn1(p0 uint32, p1 int16, p2 int8) uint8 {
    v12 := []int{823}
    var v13 int16 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 0.0
    var v3 int8 = -106
    var v4 int = -103
    var v5 float32 = 3.14159
    v6 := make(map[int]bool)
    var v14 uint = 413
    v15 := GenS1{F0: 0, F1: (-(v3))}
    v5 = float32(float64(0.0))
    switch v4 {
    case 883:
        v16 := GenS1{F0: 9223372036854775807, F1: v3}
        _ = v16
    default:
        v17 := GenS2{F0: 0, F1: true}
        _ = v17
    }
    switch "a" {
    case "foo bar":
        var v18 uint = (v14 * v14)
        _ = v18
    case "\n\t":
        v3 += 127
    default:
        switch 42 {
        case 260:
            if (genFn0(1) != 0) {
                fmt.Println(v14, v14)
                for fi19 := 0; fi19 < 20; fi19++ {
                    v4 = v4
                    _ = fi19
                }
                for fi20 := 0; fi20 < 3; fi20++ {
                    v5 += ((3.14159 + 1.5) * v5)
                    var v21 uint = v14
                    _ = fi20
                    _ = v21
                }
            }
        case 32767:
            for fi22 := 0; fi22 < 4; fi22++ {
                fmt.Println(v4, v4)
                fmt.Println(v4, v4, v5)
                for fi23 := 0; fi23 < 5; fi23++ {
                    v24 := GenS1{F0: 0, F1: (v3 + v3)}
                    if (true || false) {
                        delete(v1, 818)
                        if false {
                            v25 := []int{}
                            fmt.Println(v5, fi23, v4)
                            _ = v25
                        } else {
                            var v26 float64 = v2
                            var v27 int8 = v3
                            if (!true) {
                                fmt.Println(v5)
                            }
                            _ = v26
                            _ = v27
                        }
                        var v28 int32 = -2147483648
                        v29 := []int{}
                        _ = v28
                        _ = v29
                    }
                    var v30 bool = false
                    _ = fi23
                    _ = v24
                    _ = v30
                }
                _ = fi22
            }
        default:
            v31 := GenS1{F0: 1, F1: (v3 - (v3 + (v3 / (((v3 - -128)) | 1))))}
            _ = v31
        }
    }
    v6[(v4 ^ 816)] = true
    delete(v6, v4)
    fmt.Println(v14, v5, v2)
    fmt.Println(v5, v2)
    fmt.Println(v3, v2)
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 int16 = 1000
        fmt.Println(v2, v33)
        _ = fi32
        _ = v33
    }
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
