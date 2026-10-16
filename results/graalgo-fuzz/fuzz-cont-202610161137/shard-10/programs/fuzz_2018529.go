package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 int; F2 uint16 }
type GenS5 struct { F0 int; F1 int64; F2 uint }
type GenS6 struct { F0 uint16 }
func genFn0(p0 int64, p1 int) uint8 {
    var v7 int16 = -480
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint64 = 370
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 2147483647, F1: 2147483648}
    v11 := []int{(^(p1)), (p1 % ((-88) | 1))}
    v12 := GenS1{F0: -2147483648, F1: 2147483648}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v7
    return 168
}
func genFn1(p0 bool) uint64 {
    if (128 < 42) {
        v13 := GenS1{F0: -460, F1: 955}
        _ = v13
    }
    var v14 uint = 2
    p0 = true
    _ = p0
    _ = v14
    return 4294967295
}

func main() {
    var v1 int = 9223372036854775807
    v2 := []int{(v1 ^ v1)}
    var v3 int8 = -1
    v4 := make(map[int]int)
    v5 := make(map[int]string)
    var v6 uint64 = 0
    v3 = (v3 << ((18446744073709551615) & 63))
    for fi15 := 0; fi15 < 18; fi15++ {
        delete(v4, v1)
        _ = fi15
    }
    v16 := GenS2{F0: v3}
    v1 += (v1 ^ 2147483647)
    var v17 string = "a"
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := GenS3{F0: (false && true)}
        if false {
            if false {
                var v20 int16 = 772
                _ = v20
            }
            if (!true) {
                var v21 uint32 = 340
                var v22 int32 = 0
                v21 -= v21
                _ = v21
                _ = v22
            } else {
                if len(v2) > 0 {
                    v2[0] = v1
                }
            }
            v17 = v17
        } else {
            switch v17 {
            case "a":
                switch v1 {
                case -728:
                    v23 := GenS4{F0: 4294967295, F1: 9223372036854775807, F2: 255}
                    _ = v23
                case 1:
                    fmt.Println(v1, v3)
                case 128:
                    for fi24 := 0; fi24 < 12; fi24++ {
                        v6 = v6
                        v25 := GenS5{F0: (fi24 + (fi24 | -9223372036854775808)), F1: -1, F2: 219}
                        if (!false) {
                            _ = len(v5)
                            var v26 uint32 = 4294967295
                            var v27 uint = 2
                            _ = v26
                            _ = v27
                        } else {
                            if ((fi24 ^ fi24) >= v1) {
                                switch fi18 {
                                case 451:
                                    v5[(fi24 + v1)] = v17
                                case 2147483647:
                                    v3 = (^(v3))
                                case 127:
                                    var v28 int = v1
                                    _ = v28
                                default:
                                    _ = len(v2)
                                }
                                if len(v2) > 0 {
                                    v2[0] = fi18
                                }
                                v3 += (v3 ^ (v3 - 0))
                            }
                        }
                        v29 := GenS6{F0: 2}
                        _ = fi24
                        _ = v25
                        _ = v29
                    }
                }
            case "":
                var v30 float64 = (1.5 * float64(3.14159))
                _ = v30
            }
            if true {
                for fi31 := 0; fi31 < 8; fi31++ {
                    if (9223372036854775807 < 9223372036854775807) {
                        fmt.Println(v6, v1)
                        v3 = (v3 / ((-128) | 1))
                        delete(v4, (^(fi31)))
                    } else {
                        v3 = (v3 | 42)
                        v32 := GenS4{F0: 648, F1: v1, F2: 32767}
                        _ = v32
                    }
                    _ = fi31
                }
            }
        }
        _ = fi18
        _ = v19
    }
    v33 := GenS3{F0: (((float32((((0.1 * 1.5) / 1.0) * (3.14159 * (0.1 + 0.1)))) + ((0.1 / float32(0.1)) - (0.1 / (0.1 / -2.25)))) < float32(0.0)) && false)}
    var v34 uint32 = 4294967295
    switch v17 {
    case "\n\t":
        var v35 int32 = 855
        _ = v35
    }
    var v36 float32 = 1e10
    fmt.Println(v34)
    for fi37 := 0; fi37 < 9; fi37++ {
        var v38 int32 = 2147483647
        if true {
            v1 = (v1 / ((v1) | 1))
            var v39 uint = 0
            _ = v39
        } else {
            delete(v5, (-(fi37)))
        }
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
