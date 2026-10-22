package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 bool; F2 int64 }
func genFn0(p0 uint32) uint16 {
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 float32 = 1e10
            if false {
                fmt.Println(v8, p0, p0)
                var v9 uint16 = 1
                var v10 uint16 = v9
                _ = v10
                _ = v9
            } else {
                switch "a" {
                case "foo bar":
                    fmt.Println(p0)
                case "a":
                    var v11 float64 = 0.1
                    _ = v11
                case "":
                    v12 := GenS1{F0: -1}
                    _ = v12
                }
            }
            _ = fi7
            _ = v8
        }
        var v13 bool = true
        _ = fi6
        _ = v13
    }
    p0 -= (p0 % ((p0) | 1))
    _ = p0
    return 940
}
func genFn1(p0 int64) int32 {
    var v14 int64 = p0
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS2{F0: 2147483648}
        _ = fi15
        _ = v16
    }
    fmt.Println(v14, p0)
    _ = p0
    _ = v14
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 42
    var v3 int64 = 9223372036854775807
    var v4 float64 = 0.1
    var v5 int = 32767
    v5 += v5
    v2 -= genFn1(-9223372036854775808)
    v17 := []int{0}
    var v18 int64 = (v3 % ((9223372036854775807) | 1))
    v3 = v18
    switch 32767 {
    case 127:
        switch "\n\t" {
        case "hello":
            for fi19 := 0; fi19 < 8; fi19++ {
                var v20 int8 = 127
                var v21 uint = 1
                _ = fi19
                _ = v20
                _ = v21
            }
        case "a":
            fmt.Println(v4, v18, v18)
        default:
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 int = v5
                if false {
                    if (true && (42 == v3)) {
                        var v24 int = (v23 >> ((639) & 63))
                        _ = v24
                    }
                    switch v23 {
                    case 172:
                        if ((!false) && true) {
                            var v25 uint16 = genFn0(81)
                            v26 := GenS3{F0: 4294967295}
                            v27 := GenS4{F0: 2, F1: (((((1e10 / 0.1) * 0.1) * (1.5 * 0.0)) + -2.25) <= 1e10), F2: v3}
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    default:
                        v23 -= 4294967295
                    }
                }
                v5 = v23
                _ = fi22
                _ = v23
            }
        }
    case -9223372036854775808:
        switch v5 {
        case 0:
            v2 = (v2 + (-(v2)))
        }
    }
    v18 -= (v18 * (^(v3)))
    v18 = v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
