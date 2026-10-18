package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint, p1 int32, p2 int8) string {
    var v6 uint8 = 127
    var v7 uint = (p0 / (((p0 >> ((13) & 63))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "hello"
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 int32 = 2147483647
    var v4 float64 = 3.14159
    var v5 bool = false
    if (((18446744073709551615 <= 65535) || v5) || v5) {
        if v5 {
            if (2 >= 4294967295) {
                v8 := GenS1{F0: 723}
                var v9 int64 = 190
                _ = v8
                _ = v9
            } else {
                var v10 int = 32767
                if (-30 < -621) {
                    if (v5 && true) {
                        v5 = (!((0.0 + 0.1) < 3.14159))
                        var v11 int64 = -882
                        var v12 float64 = v4
                        _ = v11
                        _ = v12
                    } else {
                        v13 := []int{-655}
                        _ = v13
                    }
                    if v5 {
                        for fi14 := 0; fi14 < 8; fi14++ {
                            if ((v3 + v3) >= v3) {
                                v10 += -2147483648
                                fmt.Println(v3)
                                v15 := v1[v10]
                                _ = v15
                            }
                            _ = fi14
                        }
                        fmt.Println(v4, v10)
                        v16 := GenS2{F0: 1, F1: 2, F2: 4294967295}
                        _ = v16
                    }
                    var v17 int16 = 323
                    _ = v17
                }
                fmt.Println(v5)
                _ = v10
            }
        }
    } else {
        v1[-9223372036854775808] = 4294967295
    }
    v3 = (v3 | (v3 | v3))
    var v18 uint32 = 2147483648
    var v19 int64 = -9223372036854775808
    switch 4294967295 {
    case -572:
        v20 := GenS3{F0: 2, F1: 0}
        _ = v20
    }
    fmt.Println(v19, v5)
    v21 := GenS1{F0: (v18 | v18)}
    for fi22 := 0; fi22 < 5; fi22++ {
        for fi23 := 0; fi23 < 12; fi23++ {
            v24 := []int{}
            v18 = (v18 ^ 486)
            v24 = append(v24, fi22)
            _ = fi23
            _ = v24
        }
        v25 := []int{fi22, (fi22 * -73)}
        _ = len(v1)
        _ = fi22
        _ = v25
    }
    v5 = (1 != 0)
    v26 := GenS4{F0: 99}
    var v27 uint32 = (v18 + (v18 / (((v18 >> ((2) & 63))) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
