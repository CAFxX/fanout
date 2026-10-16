package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint16; F2 int8 }
func genFn0(p0 uint64, p1 uint64, p2 int8) int64 {
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 bool = (false && true)
            fmt.Println(fi5, p2, fi6)
            if (1 > 42) {
                var v8 int32 = -2147483648
                v7 = (!(1 <= 0))
                switch (fi5 * fi5) {
                case 2147483647:
                    p0 += p0
                case -799:
                    fmt.Println(v7, fi6)
                default:
                    var v9 uint = 4294967295
                    if (1000 != 32767) {
                        v7 = (2147483647 <= (^(fi5)))
                    }
                    _ = v9
                }
                v10 := []int{}
                _ = v10
                _ = v8
            }
            fmt.Println(p0)
            v11 := []int{128, fi5}
            _ = fi6
            _ = v11
            _ = v7
        }
        _ = fi5
    }
    switch -1 {
    case -2147483648:
        v12 := []int{324, -1, -9223372036854775808}
        _ = v12
    case -9223372036854775808:
        var v13 int16 = -169
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 9223372036854775807
}
func genFn1(p0 uint) uint32 {
    var v14 float64 = 0.1
    v14 = (v14 * ((0.0 + ((float64(3.14159) * 0.1) * 1.5)) / (v14 / v14)))
    _ = p0
    _ = v14
    return 1
}

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{}
    v3 := []int{}
    v4 := []uint32{0, 65535, 65535, 318}
    v15 := GenS1{F0: (4294967295 > 0), F1: 32767, F2: 705}
    v1 = float32(1e10)
    if true {
        fmt.Println(v1)
        var v16 int16 = 1
        _ = v16
    } else {
        fmt.Println(v1)
        var v17 int16 = -560
        for fi18 := 0; fi18 < 10; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 uint = 255
                var v21 int8 = -128
                if true {
                    v22 := GenS2{F0: (fi19 % ((fi18) | 1))}
                    v23 := GenS2{F0: (fi18 + fi18)}
                    _ = v22
                    _ = v23
                } else {
                    fmt.Println(v21, v21, v17)
                    v20 += 689
                }
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi18
        }
        _ = v17
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        switch 128 {
        case 0:
            var v25 int16 = 32767
            _ = v25
        case 32768:
            for fi26 := 0; fi26 < 8; fi26++ {
                v1 = v1
                switch (fi26 / ((fi26) | 1)) {
                case 1:
                    fmt.Println(fi26, v1, fi24)
                default:
                    v27 := GenS3{F0: 4294967295, F1: 2, F2: -72}
                    _ = v27
                }
                _ = fi26
            }
        case -9223372036854775808:
            for fi28 := 0; fi28 < 13; fi28++ {
                _ = len(v4)
                v29 := GenS1{F0: false, F1: 144, F2: 42}
                _ = fi28
                _ = v29
            }
        default:
            for fi30 := 0; fi30 < 20; fi30++ {
                v4 = append(v4, 64)
                switch "\n\t" {
                case "\n\t":
                    v1 = 0.1
                }
                if len(v2) > 0 {
                    v2[0] = 280
                }
                _ = fi30
            }
        }
        _ = len(v3)
        _ = fi24
    }
    v31 := GenS2{F0: 617}
    v1 += (v1 * v1)
    var v32 uint32 = 2147483647
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 += v1
    if (0 != -2147483648) {
        v32 = 1
    } else {
        var v33 uint = 255
        _ = v33
    }
    var v34 int8 = 0
    var v35 uint16 = 449
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v4
}
