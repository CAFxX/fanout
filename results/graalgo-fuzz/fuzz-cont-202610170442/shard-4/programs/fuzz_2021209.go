package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int) uint32 {
    fmt.Println(p0)
    p0 -= p0
    _ = p0
    return 2147483648
}
func genFn1(p0 bool, p1 uint16, p2 uint) int {
    fmt.Println(p0, p0)
    v4 := []int{160, -604, -71}
    fmt.Println(p2)
    for fi5 := 0; fi5 < 10; fi5++ {
        fmt.Println(p0, p1)
        _ = fi5
    }
    if (32767 == p1) {
        fmt.Println(p1, p1, p2)
        fmt.Println(p1, p2, p2)
        v4 = append(v4, 32767)
    } else {
        p1 = 42
        switch 0 {
        case 55:
            switch -1 {
            case -9223372036854775808:
                v6 := GenS1{F0: 32768}
                _ = v6
            case 617:
                _ = len(v4)
            }
        case 833:
            p1 = 639
        case 32768:
            for fi7 := 0; fi7 < 17; fi7++ {
                for fi8 := 0; fi8 < 1; fi8++ {
                    switch fi7 {
                    case -624:
                        if (32767 != 42) {
                            var v9 int = -733
                            _ = v9
                        }
                    case 9223372036854775807:
                        switch (fi8 ^ (fi8 & 541)) {
                        case -9223372036854775808:
                            p0 = (!p0)
                        }
                    default:
                        fmt.Println(p0, p1)
                    }
                    if p0 {
                        if len(v4) > 0 {
                            v4[0] = fi8
                        }
                        if (32767 < 549) {
                            switch -9223372036854775808 {
                            case -2147483648:
                                if (float32((-2.25 + (1e10 * 3.14159))) >= (3.14159 + (((0.1 - -2.25) / 1.5) / (-2.25 * 1.5)))) {
                                    var v10 uint64 = 2
                                    fmt.Println(p0, p0, v10)
                                    if true {
                                        var v11 int64 = -1
                                        _ = v11
                                    }
                                    _ = v10
                                } else {
                                    v12 := GenS1{F0: fi8}
                                    _ = v12
                                }
                            case -9223372036854775808:
                                p2 += (^(p2))
                            case -290:
                                p0 = (!p0)
                            }
                            p2 = p2
                            fmt.Println(p0, p0, p1)
                        }
                    }
                    var v13 uint32 = 1
                    _ = fi8
                    _ = v13
                }
                _ = len(v4)
                v4 = append(v4, fi7)
                _ = fi7
            }
        }
        p1 += (p1 & (p1 / (((p1 | (p1 ^ p1))) | 1)))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    var v3 uint32 = 783
    if false {
        fmt.Println(v3)
    }
    v14 := []int{}
    v15 := GenS2{F0: 1, F1: 4294967295}
    fmt.Println(v3)
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    v3 = (v3 & (v3 ^ v3))
    var v16 int16 = -1
    fmt.Println(v16, v16)
    switch 0 {
    case -9223372036854775808:
        var v17 float64 = (1e10 / -2.25)
        _ = v17
    default:
        var v18 uint16 = 2
        _ = v18
    }
    v3 = 42
    v19 := GenS1{F0: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
