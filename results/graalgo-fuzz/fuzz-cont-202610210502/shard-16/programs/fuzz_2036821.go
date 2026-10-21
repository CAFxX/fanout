package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint; F2 int32 }
type GenS5 struct { F0 bool; F1 int16; F2 uint64 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8; F1 bool; F2 uint32 }
type GenS9 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS10 struct { F0 int8; F1 uint64 }
func genFn0(p0 int8, p1 int32, p2 int) int64 {
    v6 := []int{(-(p2)), -799, int(p0)}
    if false {
        if (734 > 255) {
            var v7 int64 = -831
            v8 := GenS1{F0: 2147483647}
            fmt.Println(p0, p0, p1)
            _ = v7
            _ = v8
        }
        var v9 uint32 = 4294967295
        _ = v9
    } else {
        v10 := GenS1{F0: 2147483647}
        if (-2.25 != (-2.25 * 0.0)) {
            v11 := GenS1{F0: 65535}
            switch "foo bar" {
            case "\n\t":
                p1 -= (p1 ^ p1)
            case "hello":
                var v12 uint = 1
                _ = v12
            case "a":
                for fi13 := 0; fi13 < 9; fi13++ {
                    p2 -= p2
                    for fi14 := 0; fi14 < 5; fi14++ {
                        p1 = p1
                        if (18446744073709551615 != 18446744073709551615) {
                            fmt.Println(fi13)
                            if (!true) {
                                v6 = append(v6, (^(p2)))
                                v6 = append(v6, -85)
                            } else {
                                var v15 int64 = -659
                                fmt.Println(p0, fi13, v15)
                                _ = len(v6)
                                _ = v15
                            }
                        } else {
                            if (100 != (p0 * (p0 << ((2) & 63)))) {
                                var v16 uint16 = 65535
                                _ = v16
                            } else {
                                v17 := GenS2{F0: (p2 >> ((1) & 63))}
                                switch "hello" {
                                case "hello":
                                    var v18 uint8 = 42
                                    _ = v18
                                case "a":
                                    switch "\n\t" {
                                    case "a":
                                        p1 += (p1 + p1)
                                    case "hello":
                                        var v19 uint16 = 0
                                        _ = v19
                                    default:
                                        p1 += (-(p1))
                                    }
                                case "foo bar":
                                    if (127 <= p0) {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        p1 = (p1 ^ p1)
                                        var v20 float32 = (0.1 + ((1.5 + (float32(-2.25) * float32(0.0))) - 1.5))
                                        _ = v20
                                    } else {
                                        var v21 int32 = p1
                                        v22 := GenS1{F0: 0}
                                        _ = v21
                                        _ = v22
                                    }
                                default:
                                    fmt.Println(p2, p0)
                                }
                                if false {
                                    switch 32767 {
                                    case -615:
                                        p1 -= p1
                                    case 32767:
                                        v23 := GenS2{F0: int(p0)}
                                        _ = v23
                                    case 9223372036854775807:
                                        p0 += p0
                                        if ((-(fi14)) == fi13) {
                                            var v24 uint64 = 1
                                            _ = v24
                                        } else {
                                            p0 = p0
                                            if len(v6) > 0 {
                                                v6[0] = (^(fi14))
                                            }
                                        }
                                    default:
                                        if true {
                                            var v25 bool = (false && false)
                                            var v26 float32 = ((float32(1.5) / 1e10) / float32((3.14159 / -2.25)))
                                            _ = v25
                                            _ = v26
                                        }
                                    }
                                    if len(v6) > 0 {
                                        v6[0] = p2
                                    }
                                }
                                _ = v17
                            }
                            fmt.Println(fi14, p0)
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        if true {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        } else {
                            var v27 int32 = (-(p1))
                            _ = v27
                        }
                        _ = fi14
                    }
                    _ = fi13
                }
            default:
                v28 := GenS3{F0: (p0 % (((p0 + 0)) | 1)), F1: 2, F2: 128}
                _ = v28
            }
            _ = v11
        }
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 9223372036854775807
}

func main() {
    v1 := []int{-174, 127, 127}
    var v2 int32 = 262
    var v3 int = 1
    var v4 int16 = -1
    var v5 uint16 = 789
    v29 := GenS4{F0: (-(v4)), F1: 667, F2: (-(v2))}
    var v30 float64 = 3.14159
    v31 := GenS5{F0: (65535 <= 0), F1: v4, F2: 42}
    var v32 uint32 = 574
    v3 -= 29
    v33 := GenS6{F0: (185 <= 0), F1: 1}
    v34 := GenS6{F0: false, F1: 18446744073709551615}
    var v35 bool = false
    var v36 uint16 = (^(v5))
    if ((-95 > 100) && true) {
        for fi37 := 0; fi37 < 14; fi37++ {
            v3 -= fi37
            var v38 int64 = genFn0(100, (^(v2)), v3)
            _ = fi37
            _ = v38
        }
        v39 := GenS7{F0: (v5 | v36)}
        v40 := GenS8{F0: -1, F1: (!((v36 / (((v36 / (((v36 << ((2) & 63))) | 1))) | 1)) >= (^(v5)))), F2: v32}
        _ = v39
        _ = v40
    }
    v1 = append(v1, v3)
    var v41 float64 = v30
    var v42 uint32 = uint32(v36)
    fmt.Println(v2)
    v43 := GenS9{F0: 4294967295, F1: 334, F2: 2}
    var v44 int64 = 42
    for fi45 := 0; fi45 < 1; fi45++ {
        var v46 int32 = v2
        switch "\n\t" {
        case "foo bar":
            if len(v1) > 0 {
                v1[0] = (v3 & -557)
            }
        case "\n\t":
            for fi47 := 0; fi47 < 20; fi47++ {
                v36 = 32767
                v48 := GenS10{F0: -128, F1: 771}
                _ = fi47
                _ = v48
            }
        case "":
            v42 = (v32 << ((294) & 63))
        }
        _ = fi45
        _ = v46
    }
    if v35 {
        v5 = (v5 - (v36 << ((767) & 63)))
        var v49 uint = 0
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
