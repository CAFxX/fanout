package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 int16; F2 int64 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int64; F1 string }
type GenS9 struct { F0 int16; F1 string }
type GenS10 struct { F0 int16; F1 uint64 }
type GenS11 struct { F0 uint64; F1 int64 }
type GenS12 struct { F0 uint64; F1 uint64 }
func genFn0(p0 bool, p1 int, p2 uint64) uint64 {
    var v7 int64 = 46
    v8 := GenS1{F0: -128, F1: -87}
    var v9 uint32 = 65535
    fmt.Println(p1, v9, v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p2
}
func genFn1(p0 int64, p1 int16, p2 int32) uint8 {
    v10 := GenS1{F0: 1, F1: (p1 / ((235) | 1))}
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 uint32 = 1
        if true {
            p2 -= p2
            v12 = 2
            v13 := GenS2{F0: 2}
            _ = v13
        } else {
            v14 := []int{2147483647}
            var v15 int = -705
            if (-2.25 != float64(float32((1.5 + 0.0)))) {
                if true {
                    v16 := GenS1{F0: -1, F1: 1000}
                    v17 := GenS3{F0: p0}
                    var v18 uint8 = 128
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    v19 := GenS4{F0: 2, F1: p1, F2: (p0 >> ((2) & 63))}
                    v20 := GenS1{F0: 127, F1: (p1 % ((-1) | 1))}
                    _ = v19
                    _ = v20
                }
                for fi21 := 0; fi21 < 19; fi21++ {
                    _ = len(v14)
                    _ = fi21
                }
            } else {
                var v22 int = v15
                if true {
                    var v23 uint16 = 1
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    if true {
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        if (565 != 65535) {
                            p1 = p1
                        }
                    } else {
                        var v24 int = (v15 | (fi11 << ((uint(v23)) & 63)))
                        v25 := GenS4{F0: 885, F1: 1, F2: p0}
                        p0 = p0
                        _ = v24
                        _ = v25
                    }
                    _ = v23
                }
                fmt.Println(v22, v15, v22)
                _ = v22
            }
            _ = v14
            _ = v15
        }
        v26 := GenS1{F0: -128, F1: (p1 / ((p1) | 1))}
        _ = fi11
        _ = v12
        _ = v26
    }
    var v27 int32 = -2147483648
    if (("" + "\n\t") != (("" + "foo bar") + "foo bar")) {
        var v28 int64 = p0
        for fi29 := 0; fi29 < 4; fi29++ {
            var v30 string = "a"
            var v31 int = (-(fi29))
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v27
    return 1
}

func main() {
    var v1 uint16 = 255
    v2 := []uint32{84}
    v3 := []uint32{329, 143, 77}
    v4 := make(map[uint32]bool)
    var v5 uint32 = 1
    v6 := []string{"foo bar"}
    switch 128 {
    case -9223372036854775808:
        var v32 uint8 = 0
        _ = v32
    case 42:
        var v33 uint16 = v1
        _ = v33
    default:
        v34 := GenS5{F0: ("foo bar" + "\n\t"), F1: genFn0(("\n\t" != ("a" + "a")), 42, genFn0((32767 <= 127), 42, 42))}
        _ = v34
    }
    v5 += v5
    for fi35 := 0; fi35 < 11; fi35++ {
        v1 += v1
        var v36 float64 = 0.1
        _ = fi35
        _ = v36
    }
    switch ("hello" + "") {
    case "\n\t":
        if (((3.14159 / float64(1.5)) != (0.1 - 1.5)) || false) {
            if (1e10 > 1.5) {
                v1 += (v1 << ((12) & 63))
                if (((0.1 * 0.1) - -2.25) != 0.0) {
                    var v37 uint = uint(v5)
                    var v38 int8 = 0
                    _ = v37
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 9; fi39++ {
                        switch (fi39 / (((fi39 * (fi39 * fi39))) | 1)) {
                        case 32768:
                            for fi40 := 0; fi40 < 1; fi40++ {
                                if (-1 != -687) {
                                    var v41 int = fi40
                                    var v42 uint32 = (v5 + 0)
                                    var v43 int8 = -1
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                }
                                _ = fi40
                            }
                        case 32767:
                            if (genFn0(((((v5 ^ (v5 / ((42) | 1))) <= 42) && true) && (100 <= 127)), fi39, genFn0(true, fi39, genFn0(true, fi39, genFn0(((-418 == fi39) || false), fi39, 786)))) < genFn0((42 <= 128), fi39, 42)) {
                                v1 -= v1
                                v1 = 2
                            } else {
                                for fi44 := 0; fi44 < 12; fi44++ {
                                    if ((fi39 * 0) > (fi39 * fi39)) {
                                        switch fi39 {
                                        case 42:
                                            v1 = (v1 * 123)
                                        }
                                        v45 := GenS2{F0: (v5 | (v5 / (((^(v5))) | 1)))}
                                        if false {
                                            v5 = 2147483647
                                        } else {
                                            v1 += 2
                                            if false {
                                                var v46 uint64 = 2
                                                var v47 int8 = 20
                                                _ = v46
                                                _ = v47
                                            }
                                            switch "foo bar" {
                                            case "\n\t":
                                                fmt.Println(v1, fi39, v1)
                                            case "":
                                                v4[(v5 - (v5 - v5))] = false
                                            case "hello":
                                                var v48 int64 = -712
                                                _ = v48
                                            default:
                                                v1 = (^(v1))
                                            }
                                        }
                                        _ = v45
                                    }
                                    var v49 uint = 2
                                    if false {
                                        _ = len(v6)
                                        if false {
                                            var v50 int64 = -193
                                            _ = v50
                                        } else {
                                            v5 -= v5
                                            if true {
                                                v1 += (v1 << ((680) & 63))
                                                if (736 > 42) {
                                                    v51 := GenS6{F0: -27, F1: (!(false || true))}
                                                    _ = v51
                                                } else {
                                                    switch "" {
                                                    case "\n\t":
                                                        var v52 uint = (v49 | (^(v49)))
                                                        _ = v52
                                                    default:
                                                        v5 -= (v5 / ((v5) | 1))
                                                    }
                                                    fmt.Println(fi44, v5, v5)
                                                    fmt.Println(v1, fi44)
                                                    v1 = v1
                                                }
                                                v53 := GenS4{F0: 387, F1: 32767, F2: -9223372036854775808}
                                                _ = v53
                                            } else {
                                                v54 := GenS7{F0: 1}
                                                if false {
                                                    if ((-766 != 967) && false) {
                                                        switch 32768 {
                                                        case 1:
                                                            if (true && false) {
                                                                var v55 uint8 = 1
                                                                v56 := GenS8{F0: 0, F1: "a"}
                                                                if false {
                                                                    var v57 uint = (v49 % ((v49) | 1))
                                                                    var v58 int32 = 1
                                                                    if false {
                                                                        v59 := v4[v5]
                                                                        _ = v59
                                                                    }
                                                                    _ = v57
                                                                    _ = v58
                                                                }
                                                                _ = v55
                                                                _ = v56
                                                            }
                                                        case -599:
                                                            if (((v1 | 2) <= v1) && false) {
                                                                v2 = append(v2, (v5 >> (((^(v49))) & 63)))
                                                                fmt.Println(v49, fi44)
                                                                if (!false) {
                                                                    v60 := GenS7{F0: 443}
                                                                    fmt.Println(v49, v1, fi39)
                                                                    fmt.Println(v49)
                                                                    v6 = append(v6, "\n\t")
                                                                    _ = v60
                                                                }
                                                            }
                                                        case -197:
                                                            v1 += 32767
                                                        default:
                                                            var v61 uint64 = 122
                                                            _ = v61
                                                        }
                                                        var v62 string = "foo bar"
                                                        _ = v62
                                                    } else {
                                                        var v63 string = ("foo bar" + (("a" + "") + "foo bar"))
                                                        v49 -= 2
                                                        if false {
                                                            var v64 string = "\n\t"
                                                            v64 = "\n\t"
                                                            switch 1 {
                                                            case 42:
                                                                switch fi44 {
                                                                case -699:
                                                                    if (0 <= (v5 ^ 425)) {
                                                                        v5 += v5
                                                                        if true {
                                                                            v49 = v49
                                                                            v65 := v4[2]
                                                                            _ = v65
                                                                        }
                                                                        delete(v4, (^(v5)))
                                                                        var v66 uint32 = 2147483648
                                                                        _ = v66
                                                                    }
                                                                default:
                                                                    v49 = 0
                                                                }
                                                            }
                                                            v67 := GenS7{F0: 0}
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v64
                                                            _ = v67
                                                        } else {
                                                            var v68 float32 = 0.0
                                                            _ = v68
                                                        }
                                                        if (1 < 0) {
                                                            var v69 float32 = float32(0.1)
                                                            _ = v69
                                                        } else {
                                                            v1 = (v1 | (v1 | (v1 + (v1 | (v1 | 0)))))
                                                            v49 -= (v49 * uint(v1))
                                                            if (!false) {
                                                                var v70 string = "\n\t"
                                                                if (!(1 <= -167)) {
                                                                    var v71 uint64 = 435
                                                                    var v72 string = v70
                                                                    v72 = "foo bar"
                                                                    _ = v71
                                                                    _ = v72
                                                                }
                                                                _ = v70
                                                            }
                                                            var v73 int8 = -128
                                                            _ = v73
                                                        }
                                                        _ = v63
                                                    }
                                                    var v74 float32 = (0.1 + (-2.25 * -2.25))
                                                    _ = v74
                                                }
                                                var v75 uint64 = 789
                                                _ = v54
                                                _ = v75
                                            }
                                            if (!true) {
                                                switch -18 {
                                                case -474:
                                                    _ = len(v4)
                                                case 4294967295:
                                                    var v76 uint32 = 968
                                                    v1 += (v1 * (v1 & 907))
                                                    _ = v76
                                                case -9223372036854775808:
                                                    var v77 string = "\n\t"
                                                    _ = v77
                                                default:
                                                    fmt.Println(fi39)
                                                }
                                                var v78 uint64 = 18446744073709551615
                                                _ = v78
                                            } else {
                                                v1 = 32767
                                                v1 -= v1
                                                var v79 uint = v49
                                                _ = v79
                                            }
                                        }
                                        switch int(v49) {
                                        case 32768:
                                            var v80 int = (fi39 & (fi39 * fi44))
                                            switch "a" {
                                            case "foo bar":
                                                v49 = v49
                                            case "":
                                                v5 = (^(v5))
                                            case "\n\t":
                                                var v81 uint16 = v1
                                                _ = v81
                                            default:
                                                var v82 bool = (-389 < -1)
                                                _ = v82
                                            }
                                            v83 := GenS9{F0: -91, F1: ("a" + "hello")}
                                            _ = v80
                                            _ = v83
                                        }
                                    }
                                    _ = fi44
                                    _ = v49
                                }
                            }
                        default:
                            v5 = v5
                        }
                        delete(v4, v5)
                        _ = len(v4)
                        _ = fi39
                    }
                    v5 -= v5
                    if len(v2) > 0 {
                        v2[0] = (v5 & 4294967295)
                    }
                }
                var v84 int32 = 42
                _ = v84
            } else {
                v85 := GenS6{F0: 127, F1: false}
                if false {
                    v86 := GenS10{F0: 1000, F1: 42}
                    _ = v86
                }
                for fi87 := 0; fi87 < 1; fi87++ {
                    v5 = v5
                    var v88 uint16 = 255
                    _ = fi87
                    _ = v88
                }
                _ = v85
            }
            v1 = (v1 & 255)
        }
    case "foo bar":
        var v89 int32 = -1
        _ = v89
    }
    v90 := GenS11{F0: 18446744073709551615, F1: 1}
    if (false || true) {
        for fi91 := 0; fi91 < 6; fi91++ {
            if len(v2) > 0 {
                v2[0] = v5
            }
            switch "\n\t" {
            case "\n\t":
                switch fi91 {
                case 42:
                    fmt.Println(v1, v5)
                case 971:
                    var v92 float64 = 0.1
                    _ = v92
                case 4294967295:
                    v5 -= 2147483647
                }
            }
            v1 = v1
            _ = fi91
        }
    }
    var v93 uint32 = 394
    switch 824 {
    case 127:
        v94 := GenS4{F0: 346, F1: 42, F2: -9223372036854775808}
        _ = v94
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v95 int64 = 9223372036854775807
    var v96 uint64 = 4294967295
    v93 -= (v5 * (v5 % ((v93) | 1)))
    if len(v3) > 0 {
        v3[0] = (v93 | (v5 << ((255) & 63)))
    }
    v96 = v96
    v95 = v95
    _ = len(v2)
    var v97 uint16 = 195
    v98 := GenS12{F0: (v96 % ((v96) | 1)), F1: genFn0(true, 0, (v96 % (((v96 / ((genFn0((1e10 > 1e10), 32767, (v96 / ((v96) | 1)))) | 1))) | 1)))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v93:", v93)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v90
    _ = v93
    _ = v95
    _ = v96
    _ = v97
    _ = v98
}
