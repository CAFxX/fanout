package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int8 }
type GenS2 struct { F0 uint; F1 uint64; F2 bool }
func genFn0(p0 int64, p1 uint32) int {
    var v4 uint8 = 184
    var v5 int16 = 597
    if true {
        if (4294967295 >= 2) {
            if false {
                if (1 == 1) {
                    v4 = 1
                }
            }
            v6 := GenS1{F0: "\n\t", F1: 473, F2: 0}
            _ = v6
        }
    }
    fmt.Println(p1, p0)
    v7 := GenS2{F0: 42, F1: 2, F2: false}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v7
    return 791
}
func genFn1(p0 uint8) string {
    if (614 == 42) {
        v8 := []int{}
        var v9 uint16 = 389
        fmt.Println(v9, v9)
        _ = v8
        _ = v9
    } else {
        var v10 bool = false
        switch -903 {
        case 2147483647:
            for fi11 := 0; fi11 < 11; fi11++ {
                v12 := []int{(fi11 - fi11), -181, (fi11 / ((42) | 1))}
                _ = fi11
                _ = v12
            }
        case 1:
            var v13 uint8 = 127
            _ = v13
        default:
            v14 := GenS2{F0: 255, F1: 42, F2: true}
            _ = v14
        }
        _ = v10
    }
    if (false && true) {
        if (!(1 >= 1)) {
            fmt.Println(p0)
            p0 = (p0 ^ (p0 ^ (p0 >> ((2) & 63))))
        } else {
            for fi15 := 0; fi15 < 2; fi15++ {
                p0 = p0
                _ = fi15
            }
            p0 -= p0
        }
        fmt.Println(p0)
    } else {
        if false {
            p0 += (^(p0))
        }
        fmt.Println(p0)
    }
    v16 := GenS2{F0: 0, F1: 2, F2: (true && true)}
    _ = p0
    _ = v16
    return "\n\t"
}

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    var v17 int8 = -1
    v17 += 100
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 uint16 = 42
        v20 := GenS1{F0: "", F1: 995, F2: (v17 % (((v17 * 100)) | 1))}
        v19 -= v19
        _ = fi18
        _ = v19
        _ = v20
    }
    v1 = 9223372036854775807
    switch 2147483647 {
    case 870:
        v1 += 9223372036854775807
    }
    switch "\n\t" {
    case "\n\t":
        v21 := []int{-800, 9223372036854775807, -2147483648}
        _ = v21
    case "a":
        if true {
            for fi22 := 0; fi22 < 18; fi22++ {
                if false {
                    var v23 int32 = 0
                    for fi24 := 0; fi24 < 8; fi24++ {
                        switch ("" + "hello") {
                        case "":
                            v23 -= -1
                            if (9223372036854775807 <= (^(v1))) {
                                switch genFn0(-9223372036854775808, 548) {
                                case -1:
                                    v25 := GenS1{F0: "a", F1: 524, F2: (v17 / (((v17 << ((4294967295) & 63))) | 1))}
                                    _ = v25
                                default:
                                    v23 = v23
                                }
                                v2[1] = "a"
                            }
                        case "\n\t":
                            var v26 float64 = 1e10
                            _ = v26
                        case "a":
                            v1 -= (v1 + (v1 - v1))
                        }
                        _ = fi24
                    }
                    _ = v23
                }
                fmt.Println(v17, fi22, fi22)
                var v27 float64 = 1e10
                _ = fi22
                _ = v27
            }
            v28 := []int{genFn0((v1 % ((v1) | 1)), 1)}
            if true {
                var v29 float64 = (0.1 * (1.5 - (3.14159 / 1.0)))
                _ = v29
            }
            _ = v28
        } else {
            for fi30 := 0; fi30 < 11; fi30++ {
                v17 = (v17 - -89)
                _ = fi30
            }
            v1 = v1
        }
    case "":
        v31 := []int{0, 32767}
        _ = v31
    default:
        v1 += 700
    }
    v1 += v1
    v32 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v32
}
