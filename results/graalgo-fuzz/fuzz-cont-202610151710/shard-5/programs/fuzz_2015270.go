package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 int; F2 uint }
type GenS7 struct { F0 uint8; F1 int8; F2 bool }
type GenS8 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS9 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS10 struct { F0 int64; F1 uint }
type GenS11 struct { F0 string; F1 uint64; F2 int16 }
type GenS12 struct { F0 bool; F1 string }
type GenS13 struct { F0 uint; F1 int32; F2 uint16 }
func genFn0(p0 int, p1 uint32, p2 int8) bool {
    v4 := GenS1{F0: (p0 * p0)}
    var v5 bool = true
    fmt.Println(v5, v5, v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return false
}

func main() {
    var v1 float64 = 0.1
    var v2 uint8 = 38
    var v3 uint = 857
    v3 -= v3
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint32 = 42
        var v8 uint = v3
        var v9 bool = true
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v3 += uint(v2)
    v10 := GenS1{F0: 32767}
    v11 := []int{}
    if (-27 != 1) {
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 float64 = float64((3.14159 - ((1e10 * 3.14159) + 0.1)))
            if true {
                v14 := GenS2{F0: 729}
                fmt.Println(fi12, v13)
                if true {
                    var v15 uint64 = 0
                    if true {
                        _ = len(v11)
                    } else {
                        fmt.Println(fi12)
                        var v16 uint32 = 65535
                        _ = v16
                    }
                    v11 = append(v11, (fi12 | 42))
                    _ = v15
                } else {
                    v17 := GenS3{F0: 42}
                    if (32767 > -962) {
                        var v18 bool = ((1 != 837) && (65535 == uint16(v3)))
                        _ = v18
                    } else {
                        v3 = v3
                        if (v1 == float64(0.0)) {
                            fmt.Println(fi12, v1)
                            v19 := GenS4{F0: (v2 << (((v3 + v3)) & 63)), F1: 0}
                            _ = v19
                        } else {
                            if (!(!((false && (float32(0.1) >= -2.25)) && true))) {
                                v3 = v3
                                if len(v11) > 0 {
                                    v11[0] = (fi12 - (-(fi12)))
                                }
                            } else {
                                if true {
                                    v20 := GenS5{F0: 589}
                                    v21 := GenS6{F0: 2147483648, F1: (fi12 % ((fi12) | 1)), F2: (^(v3))}
                                    for fi22 := 0; fi22 < 2; fi22++ {
                                        fmt.Println(v1, v13)
                                        v13 = v13
                                        _ = fi22
                                    }
                                    _ = v20
                                    _ = v21
                                } else {
                                    for fi23 := 0; fi23 < 20; fi23++ {
                                        var v24 int32 = 42
                                        if (!false) {
                                            v3 -= 42
                                            v25 := GenS5{F0: 18446744073709551615}
                                            _ = v25
                                        }
                                        _ = fi23
                                        _ = v24
                                    }
                                    var v26 uint32 = 65535
                                    _ = v26
                                }
                                v1 = (1.5 * v1)
                            }
                            fmt.Println(v1, v3)
                            var v27 string = "hello"
                            _ = v27
                        }
                        v28 := GenS6{F0: 4294967295, F1: (fi12 % ((fi12) | 1)), F2: (^(v3))}
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 17; fi29++ {
                        var v30 float64 = v1
                        v1 += (-2.25 / v13)
                        _ = fi29
                        _ = v30
                    }
                    _ = v17
                }
                _ = v14
            }
            fmt.Println(v1, v1, v3)
            _ = fi12
            _ = v13
        }
        var v31 int64 = -576
        _ = v31
    } else {
        var v32 int32 = 770
        var v33 int32 = 2147483647
        v34 := GenS7{F0: v2, F1: -100, F2: true}
        _ = v32
        _ = v33
        _ = v34
    }
    if len(v11) > 0 {
        v11[0] = -9223372036854775808
    }
    _ = len(v11)
    v1 = 1e10
    fmt.Println(v2)
    switch 0 {
    case 1:
        var v35 int16 = 42
        _ = v35
    case 128:
        if false {
            v1 = v1
            switch -2147483648 {
            case 821:
                v3 += (v3 ^ 2)
            default:
                if false {
                    switch "foo bar" {
                    case "":
                        for fi36 := 0; fi36 < 2; fi36++ {
                            if false {
                                v37 := GenS4{F0: v2, F1: 4294967295}
                                switch "\n\t" {
                                case "foo bar":
                                    fmt.Println(v3)
                                case "a":
                                    for fi38 := 0; fi38 < 14; fi38++ {
                                        v39 := GenS8{F0: 846, F1: 663, F2: 342}
                                        var v40 int32 = -1
                                        _ = fi38
                                        _ = v39
                                        _ = v40
                                    }
                                case "":
                                    v1 -= v1
                                default:
                                    fmt.Println(v3, fi36)
                                }
                                _ = v37
                            }
                            v41 := GenS9{F0: 892, F1: 2147483647, F2: 838}
                            for fi42 := 0; fi42 < 14; fi42++ {
                                var v43 bool = (!(-917 == fi42))
                                _ = fi42
                                _ = v43
                            }
                            _ = fi36
                            _ = v41
                        }
                    case "hello":
                        var v44 string = ""
                        _ = v44
                    case "a":
                        var v45 int16 = -659
                        _ = v45
                    }
                    var v46 string = "hello"
                    _ = v46
                }
            }
            fmt.Println(v3, v3)
        } else {
            v11 = append(v11, 32767)
            var v47 bool = false
            _ = v47
        }
    default:
        switch 32768 {
        case -9223372036854775808:
            for fi48 := 0; fi48 < 20; fi48++ {
                v49 := GenS10{F0: 614, F1: (v3 ^ (^(v3)))}
                v50 := GenS11{F0: "hello", F1: 2, F2: 1000}
                _ = fi48
                _ = v49
                _ = v50
            }
        case -1:
            v3 -= 1
        case -2147483648:
            if (v1 < v1) {
                if (0 > -128) {
                    switch 4294967295 {
                    case 599:
                        var v51 uint8 = v2
                        _ = v51
                    case 0:
                        if (2147483648 >= 4294967295) {
                            for fi52 := 0; fi52 < 3; fi52++ {
                                v53 := GenS12{F0: false, F1: ((("hello" + "hello") + "foo bar") + "a")}
                                switch (fi52 - (fi52 * fi52)) {
                                case 170:
                                    fmt.Println(fi52)
                                case 439:
                                    v54 := GenS5{F0: 0}
                                    _ = v54
                                case 467:
                                    v3 = 65535
                                default:
                                    var v55 int8 = 127
                                    _ = v55
                                }
                                _ = fi52
                                _ = v53
                            }
                        }
                    case 302:
                        var v56 uint8 = v2
                        _ = v56
                    default:
                        var v57 int16 = 42
                        _ = v57
                    }
                    if false {
                        v2 -= (^(v2))
                        v1 -= 0.0
                        for fi58 := 0; fi58 < 1; fi58++ {
                            v59 := GenS5{F0: 0}
                            switch "hello" {
                            case "a":
                                for fi60 := 0; fi60 < 19; fi60++ {
                                    switch (fi60 - fi58) {
                                    case 171:
                                        var v61 string = "\n\t"
                                        _ = v61
                                    case 0:
                                        v62 := GenS13{F0: 583, F1: -1, F2: 1}
                                        _ = v62
                                    case -9223372036854775808:
                                        v3 = 18446744073709551615
                                    default:
                                        v2 += v2
                                    }
                                    fmt.Println(fi60, v1, fi58)
                                    _ = fi60
                                }
                            case "\n\t":
                                v63 := GenS10{F0: 1, F1: v3}
                                _ = v63
                            case "hello":
                                fmt.Println(fi58, v2)
                            default:
                                fmt.Println(v2, v2, v3)
                            }
                            v64 := GenS1{F0: fi58}
                            _ = fi58
                            _ = v59
                            _ = v64
                        }
                    }
                } else {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    if (-2147483648 != 42) {
                        for fi65 := 0; fi65 < 1; fi65++ {
                            var v66 int16 = 953
                            switch -474 {
                            case -2147483648:
                                v66 -= 927
                            default:
                                v2 = 255
                            }
                            _ = fi65
                            _ = v66
                        }
                    } else {
                        var v67 uint32 = 2147483648
                        _ = v67
                    }
                }
            } else {
                var v68 uint8 = 128
                if (1000 < 1000) {
                    for fi69 := 0; fi69 < 1; fi69++ {
                        var v70 uint16 = 255
                        v68 -= (^(v2))
                        _ = fi69
                        _ = v70
                    }
                } else {
                    switch "" {
                    case "foo bar":
                        fmt.Println(v3)
                    case "":
                        fmt.Println(v2, v2, v1)
                    }
                    v11 = append(v11, -762)
                    var v71 int = -753
                    _ = v71
                }
                var v72 int = 4294967295
                _ = v68
                _ = v72
            }
        }
    }
    for fi73 := 0; fi73 < 4; fi73++ {
        v11 = append(v11, fi73)
        v3 = (v3 + 1)
        _ = fi73
    }
    for fi74 := 0; fi74 < 12; fi74++ {
        switch ("" + "foo bar") {
        case "foo bar":
            v75 := GenS13{F0: v3, F1: 1, F2: uint16(v2)}
            _ = v75
        default:
            var v76 int8 = 42
            _ = v76
        }
        v2 = v2
        var v77 int64 = -1
        _ = fi74
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
}
