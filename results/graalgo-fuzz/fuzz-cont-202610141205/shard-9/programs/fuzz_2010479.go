package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int16 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int64, p1 uint) uint {
    var v5 float64 = float64(3.14159)
    v6 := []int{}
    v5 = v5
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p1 & (p1 * p1))
}
func genFn1(p0 uint, p1 int16) int16 {
    v7 := []int{-176, -913, 2147483647, 32767}
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int = fi8
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return (p1 * (-(p1)))
}

func main() {
    var v1 int8 = 0
    var v2 int64 = -1
    var v3 uint = 2
    var v4 int16 = -32768
    v10 := GenS1{F0: ((true || false) || true), F1: v1, F2: int16(v3)}
    var v11 float64 = 1e10
    var v12 uint = 2
    v12 -= (v12 << ((4294967295) & 63))
    var v13 uint8 = 127
    v4 -= 32767
    v14 := GenS1{F0: (true && false), F1: (-(v1)), F2: v4}
    if true {
        var v15 uint16 = 606
        _ = v15
    } else {
        v3 -= (v3 - 455)
        for fi16 := 0; fi16 < 18; fi16++ {
            v17 := []int{4294967295, 698, 430}
            var v18 int64 = (-(v2))
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    v19 := []int{-416, 127, 32768}
    v19 = append(v19, 2147483647)
    v11 = (1.5 - 0.1)
    if (genFn0((v2 + (v2 & v2)), v3) < (v12 / (((v3 >> ((uint(v13)) & 63))) | 1))) {
        var v20 int64 = v2
        if true {
            v21 := GenS1{F0: true, F1: (^(v1)), F2: genFn1(v12, -748)}
            _ = v21
        }
        _ = v20
    } else {
        switch -9223372036854775808 {
        case -721:
            if ((false || (true && true)) && false) {
                var v22 uint8 = v13
                _ = v22
            } else {
                v23 := GenS1{F0: false, F1: (v1 | 127), F2: (v4 % (((v4 + int16(v3))) | 1))}
                var v24 int = 1
                _ = v23
                _ = v24
            }
        case 1:
            v25 := GenS2{F0: -1, F1: (^(v1))}
            _ = v25
        case 42:
            if (("foo bar" + ("" + "")) >= "") {
                v26 := GenS1{F0: (true && (true || false)), F1: 109, F2: (v4 << ((uint(v13)) & 63))}
                var v27 float64 = v11
                if (true && (255 <= v3)) {
                    v11 = float64(1.5)
                    v28 := GenS1{F0: (0 != 2147483647), F1: -108, F2: v4}
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 19; fi29++ {
                        if (!(false || (1 >= 42))) {
                            v12 = genFn0((v2 * v2), 4294967295)
                        } else {
                            if (false && (uint8(v3) >= (v13 << (((v3 * (v3 - v3))) & 63)))) {
                                if len(v19) > 0 {
                                    v19[0] = fi29
                                }
                                if len(v19) >= 2 {
                                    v19 = v19[1:len(v19)]
                                }
                                if len(v19) >= 2 {
                                    v19 = v19[1:len(v19)]
                                }
                            }
                        }
                        fmt.Println(v12)
                        if true {
                            if false {
                                fmt.Println(v2, v4, v3)
                            }
                            fmt.Println(fi29)
                            var v30 int32 = 42
                            _ = v30
                        }
                        _ = fi29
                    }
                    v11 = v11
                }
                _ = v26
                _ = v27
            }
        }
    }
    v31 := GenS1{F0: true, F1: int8(v2), F2: (v4 >> (((v12 + genFn0(996, v3))) & 63))}
    for fi32 := 0; fi32 < 4; fi32++ {
        if len(v19) > 0 {
            v19[0] = -2147483648
        }
        _ = fi32
    }
    v33 := GenS3{F0: 388}
    for fi34 := 0; fi34 < 1; fi34++ {
        if true {
            for fi35 := 0; fi35 < 11; fi35++ {
                if true {
                    v13 -= (^(v13))
                } else {
                    var v36 uint64 = 18446744073709551615
                    if true {
                        v4 = v4
                        var v37 uint16 = 464
                        v38 := GenS4{F0: v1}
                        _ = v37
                        _ = v38
                    }
                    _ = v36
                }
                _ = fi35
            }
            if len(v19) > 0 {
                v19[0] = 2147483647
            }
        }
        switch -9223372036854775808 {
        case -1:
            for fi39 := 0; fi39 < 20; fi39++ {
                v40 := GenS4{F0: v1}
                v41 := GenS4{F0: (v1 | int8(fi39))}
                var v42 uint8 = (v13 + v13)
                if (42 != -633) {
                    fmt.Println(v1, fi39, v2)
                } else {
                    v4 += (v4 % (((v4 ^ (v4 % (((-(v4))) | 1)))) | 1))
                    if len(v19) > 0 {
                        v19[0] = (fi34 - 205)
                    }
                    v3 = (v3 + v3)
                }
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
            }
        case 128:
            for fi43 := 0; fi43 < 15; fi43++ {
                var v44 bool = false
                v19 = append(v19, (fi43 + (fi43 << (((v3 * (v12 ^ genFn0(v2, (v3 - v3))))) & 63))))
                _ = fi43
                _ = v44
            }
        default:
            for fi45 := 0; fi45 < 7; fi45++ {
                v12 = v12
                var v46 uint16 = 68
                _ = fi45
                _ = v46
            }
        }
        var v47 uint32 = 0
        _ = fi34
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v4
}
