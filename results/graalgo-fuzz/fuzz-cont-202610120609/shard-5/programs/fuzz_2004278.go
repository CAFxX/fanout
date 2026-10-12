package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int; F1 uint16; F2 bool }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 int; F1 int8; F2 int32 }
func genFn0(p0 uint32, p1 int16, p2 int64) int8 {
    var v4 int64 = -1
    var v5 int = 21
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -1
}

func main() {
    v1 := []uint32{1}
    var v2 int8 = 10
    var v3 uint32 = 2147483648
    if (440 >= 2) {
        var v6 int64 = 42
        var v7 uint = 925
        var v8 uint8 = 133
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 int8 = 44
    var v10 int16 = -676
    var v11 int = 42
    for fi12 := 0; fi12 < 4; fi12++ {
        if (!(v11 < (v11 + v11))) {
            if false {
                if true {
                    v13 := GenS1{F0: v3, F1: 9223372036854775807}
                    _ = v13
                }
            }
        } else {
            if true {
                var v14 float64 = (3.14159 + -2.25)
                _ = v14
            } else {
                switch (fi12 % ((v11) | 1)) {
                case 32767:
                    var v15 int64 = 42
                    _ = v15
                case 2147483647:
                    _ = len(v1)
                case -684:
                    if (42 == 18446744073709551615) {
                        if (false || ((1e10 == 1e10) && false)) {
                            v2 += v2
                        } else {
                            var v16 int16 = (-(v10))
                            _ = v16
                        }
                        switch v11 {
                        case 128:
                            if true {
                                var v17 int32 = 0
                                _ = len(v1)
                                _ = v17
                            }
                        case 9223372036854775807:
                            var v18 int8 = (v9 - (v9 + (v2 | v9)))
                            _ = v18
                        default:
                            v19 := GenS1{F0: (v3 / ((42) | 1)), F1: -9223372036854775808}
                            _ = v19
                        }
                    } else {
                        if false {
                            switch "" {
                            case "a":
                                for fi20 := 0; fi20 < 8; fi20++ {
                                    switch (-(fi20)) {
                                    case 128:
                                        v21 := GenS2{F0: (v10 ^ (v10 | v10)), F1: 9223372036854775807}
                                        _ = v21
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = (v3 & v3)
                                    }
                                    _ = fi20
                                }
                            }
                            v22 := GenS3{F0: (v11 << ((373) & 63)), F1: 1, F2: false}
                            v23 := GenS3{F0: (-(fi12)), F1: 688, F2: true}
                            _ = v22
                            _ = v23
                        }
                        var v24 uint16 = 255
                        _ = v24
                    }
                }
                v10 = (v10 ^ int16(fi12))
            }
            for fi25 := 0; fi25 < 2; fi25++ {
                v10 = (^(v10))
                v26 := GenS2{F0: (v10 << ((1) & 63)), F1: 128}
                switch "" {
                case "a":
                    v27 := GenS1{F0: (v3 - 20), F1: -1}
                    _ = v27
                case "\n\t":
                    v11 += v11
                case "foo bar":
                    if true {
                        var v28 bool = false
                        v29 := GenS4{F0: fi12, F1: 255}
                        v2 = (v2 % ((v9) | 1))
                        _ = v28
                        _ = v29
                    }
                default:
                    if false {
                        v30 := GenS5{F0: (v10 * v10)}
                        _ = v30
                    } else {
                        if (!true) {
                            v9 -= (v2 + (-(v9)))
                        } else {
                            v10 = v10
                        }
                        if true {
                            var v31 int16 = v10
                            if (!false) {
                                v3 = v3
                            }
                            v9 = v2
                            _ = v31
                        }
                    }
                    if true {
                        fmt.Println(v11, fi25, v9)
                        v11 = -2147483648
                    } else {
                        fmt.Println(v9, v2, v3)
                        if (true || (1e10 != (3.14159 / -2.25))) {
                            v10 = (v10 << ((312) & 63))
                            var v32 int32 = 463
                            switch "a" {
                            case "foo bar":
                                var v33 int8 = (^(v9))
                                _ = v33
                            case "a":
                                var v34 uint8 = 176
                                _ = v34
                            case "":
                                if true {
                                    if ((v10 * (v10 << ((42) & 63))) >= v10) {
                                        if false {
                                            v35 := GenS6{F0: (-(v32)), F1: (v32 | v32)}
                                            _ = len(v1)
                                            switch v11 {
                                            case 4294967295:
                                                var v36 uint64 = 353
                                                _ = v36
                                            default:
                                                v9 = v2
                                            }
                                            _ = v35
                                        } else {
                                            var v37 uint32 = (^(v3))
                                            var v38 uint8 = 82
                                            var v39 uint = 532
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                        }
                                        v2 -= genFn0((v3 >> ((0) & 63)), (v10 / (((v10 >> ((1) & 63))) | 1)), -9223372036854775808)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    } else {
                                        _ = len(v1)
                                        var v40 float64 = 1.5
                                        _ = v40
                                    }
                                }
                            default:
                                var v41 uint32 = 716
                                _ = v41
                            }
                            if ((v10 + v10) != v10) {
                                if (166 >= (v10 + (v10 & (v10 + 32767)))) {
                                    v1 = append(v1, v3)
                                    var v42 int8 = v2
                                    _ = v42
                                } else {
                                    var v43 bool = false
                                    v3 = v3
                                    _ = v43
                                }
                                fmt.Println(v9, v3, v32)
                                switch "" {
                                case "hello":
                                    var v44 uint16 = 0
                                    _ = v44
                                case "a":
                                    fmt.Println(v11, fi12)
                                default:
                                    v3 = 154
                                }
                            } else {
                                var v45 float64 = 1e10
                                _ = v45
                            }
                            _ = v32
                        } else {
                            var v46 uint16 = 65535
                            if false {
                                switch (fi25 % (((fi12 * (-(fi25)))) | 1)) {
                                case 73:
                                    v9 -= 0
                                }
                                v1 = append(v1, (v3 + (v3 << ((uint(v3)) & 63))))
                                v1 = append(v1, (v3 * (v3 | 65535)))
                                v47 := GenS4{F0: fi25, F1: v46}
                                v3 -= v3
                                v48 := GenS2{F0: v10, F1: (v11 & (v11 + fi12))}
                                _ = v47
                                _ = v48
                            }
                            if len(v1) > 0 {
                                v1[0] = v3
                            }
                            _ = v46
                        }
                    }
                    switch "" {
                    case "foo bar":
                        var v49 uint16 = 119
                        _ = v49
                    default:
                        var v50 int64 = -1
                        _ = v50
                    }
                    if len(v1) > 0 {
                        v1[0] = v3
                    }
                    if true {
                        if (((((((0.0 * 1.5) * ((1.5 - 3.14159) + 0.0)) + (0.1 * 0.0)) * (((-2.25 - 3.14159) + 1.5) / 1.0)) + 1e10) + 0.0) >= float32((1e10 * (0.0 * float64(3.14159))))) {
                            if (true && true) {
                                fmt.Println(fi12, v3, v10)
                                if true {
                                    var v51 int32 = 2147483647
                                    _ = v51
                                }
                            }
                            var v52 uint16 = 65535
                            fmt.Println(fi12, fi12)
                            v53 := GenS2{F0: v10, F1: 433}
                            _ = v52
                            _ = v53
                        } else {
                            var v54 int16 = (v10 % ((v10) | 1))
                            var v55 uint8 = 127
                            v1 = append(v1, (v3 / (((^(v3))) | 1)))
                            _ = v54
                            _ = v55
                        }
                    } else {
                        var v56 uint16 = 445
                        _ = v56
                    }
                    v2 = (v9 * -59)
                    v57 := GenS7{F0: (v11 | (fi12 - -1)), F1: 1, F2: 1}
                    _ = v57
                }
                _ = fi25
                _ = v26
            }
        }
        fmt.Println(v9)
        _ = fi12
    }
    v1 = append(v1, 2)
    v58 := GenS6{F0: int32(v10), F1: -116}
    var v59 int8 = v2
    for fi60 := 0; fi60 < 6; fi60++ {
        var v61 int16 = (^(v10))
        _ = fi60
        _ = v61
    }
    for fi62 := 0; fi62 < 3; fi62++ {
        var v63 int64 = -195
        _ = fi62
        _ = v63
    }
    if (!(true || true)) {
        v59 = (v2 | v2)
        var v64 bool = true
        v65 := GenS1{F0: v3, F1: 412}
        _ = v64
        _ = v65
    }
    switch 1 {
    case 0:
        var v66 float32 = ((1.5 / 0.1) - (1e10 + 0.0))
        _ = v66
    }
    fmt.Println(v10, v3, v59)
    var v67 int = (v11 + v11)
    var v68 int64 = 9223372036854775807
    for fi69 := 0; fi69 < 14; fi69++ {
        v1 = append(v1, (^(v3)))
        _ = fi69
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v59:", v59)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v58
    _ = v59
    _ = v67
    _ = v68
    _ = v9
}
