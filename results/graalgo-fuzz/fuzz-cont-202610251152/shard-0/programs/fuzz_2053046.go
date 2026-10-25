package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS7 struct { F0 uint8; F1 int; F2 uint8 }
type GenS8 struct { F0 string; F1 uint64 }
type GenS9 struct { F0 string }
func genFn0(p0 uint32, p1 uint64) int32 {
    p1 += p1
    for fi7 := 0; fi7 < 1; fi7++ {
        p1 = p1
        _ = fi7
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 2147483647
}
func genFn1(p0 uint32, p1 int, p2 uint) int {
    p0 = p0
    if (4294967295 >= 484) {
        var v8 int16 = 1000
        _ = v8
    } else {
        v9 := GenS1{F0: 0, F1: (p1 ^ (p1 & 32768))}
        switch p1 {
        case 655:
            for fi10 := 0; fi10 < 20; fi10++ {
                for fi11 := 0; fi11 < 20; fi11++ {
                    p1 -= fi11
                    _ = fi11
                }
                fmt.Println(p2)
                v12 := []int{fi10, 32767, -121, p1}
                _ = fi10
                _ = v12
            }
        default:
            fmt.Println(p2, p1, p2)
        }
        v13 := []int{p1, p1}
        _ = v13
        _ = v9
    }
    if (32768 > (p1 ^ (-(p1)))) {
        var v14 int64 = 42
        _ = v14
    }
    if true {
        switch p1 {
        case 4294967295:
            v15 := GenS2{F0: 3, F1: 42}
            _ = v15
        case 42:
            for fi16 := 0; fi16 < 7; fi16++ {
                v17 := []int{p1, (fi16 + (fi16 | p1)), (p1 % (((fi16 + p1)) | 1))}
                _ = fi16
                _ = v17
            }
        default:
            var v18 int8 = -21
            _ = v18
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            for fi20 := 0; fi20 < 19; fi20++ {
                p1 = (fi19 - 865)
                var v21 uint32 = (p0 ^ p0)
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        var v22 int8 = 104
        _ = v22
    }
    v23 := []int{4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return 127
}

func main() {
    var v1 uint = 255
    var v2 uint32 = 42
    var v3 int64 = 9223372036854775807
    v4 := make(map[uint32]uint32)
    var v5 string = (("hello" + "foo bar") + "foo bar")
    var v6 bool = false
    v24 := GenS2{F0: uint8(v2), F1: 42}
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 uint = v1
        _ = fi25
        _ = v26
    }
    var v27 int = 42
    v28 := GenS2{F0: 0, F1: -65}
    v2 = 963
    fmt.Println(v3, v3, v6)
    v29 := GenS1{F0: 128, F1: -9223372036854775808}
    if ("" != v5) {
        if (18446744073709551615 > 18446744073709551615) {
            v27 = (v27 | (v27 + v27))
        }
        for fi30 := 0; fi30 < 5; fi30++ {
            if (!(96 != 128)) {
                for fi31 := 0; fi31 < 4; fi31++ {
                    v32 := GenS3{F0: 1}
                    if ((!(v6 || ((!(v6 && (462 != 861))) && (128 > 0)))) || (-6 < 42)) {
                        v33 := GenS4{F0: 1}
                        var v34 uint16 = 65535
                        _ = v33
                        _ = v34
                    }
                    _ = fi31
                    _ = v32
                }
                v5 = ((v5 + v5) + "")
            } else {
                for fi35 := 0; fi35 < 15; fi35++ {
                    v1 -= (v1 + (v1 * v1))
                    _ = fi35
                }
                var v36 uint8 = 0
                if ((fi30 | v27) < genFn1((^(v2)), fi30, (v1 | (^(v1))))) {
                    if false {
                        v37 := v4[v2]
                        var v38 int64 = (-(v3))
                        _ = v37
                        _ = v38
                    } else {
                        v39 := GenS4{F0: uint16(v1)}
                        _ = v39
                    }
                    fmt.Println(v1, v5)
                }
                _ = v36
            }
            switch v27 {
            case -1:
                fmt.Println(v5, v1)
            case -9223372036854775808:
                v2 = (v2 - (v2 ^ (^(v2))))
            default:
                fmt.Println(v6, v5)
            }
            _ = fi30
        }
    }
    switch 1 {
    case 128:
        v40 := []int{(^(v27)), (v27 ^ (-(v27)))}
        _ = v40
    case 32768:
        v27 = v27
    case 127:
        v41 := GenS2{F0: 255, F1: 100}
        _ = v41
    default:
        v42 := GenS4{F0: 32767}
        _ = v42
    }
    delete(v4, v2)
    v2 = (v2 ^ v2)
    var v43 uint32 = v2
    if (0.1 != (1.5 - (0.0 - (-2.25 / 0.1)))) {
        v5 = "a"
        if (!(1 <= 6)) {
            if (uint16(v2) > 203) {
                var v44 uint8 = 127
                if ((v44 <= 0) && (v5 >= v5)) {
                    if ((^(v27)) < (v27 << ((531) & 63))) {
                        var v45 string = ""
                        if (!(-1 <= int8(v43))) {
                            if ((v2 < v2) && v6) {
                                if ((1e10 + (3.14159 + 1e10)) >= ((1.5 - 0.0) - 1.5)) {
                                    var v46 string = v45
                                    v46 = v46
                                    _ = v46
                                } else {
                                    v47 := GenS4{F0: 0}
                                    fmt.Println(v3, v45)
                                    for fi48 := 0; fi48 < 17; fi48++ {
                                        var v49 float32 = 1e10
                                        v1 = 1
                                        v45 = "foo bar"
                                        _ = fi48
                                        _ = v49
                                    }
                                    _ = v47
                                }
                            }
                            v50 := GenS1{F0: 7, F1: v27}
                            switch (v27 | 149) {
                            case 32768:
                                for fi51 := 0; fi51 < 6; fi51++ {
                                    if (!true) {
                                        fmt.Println(v44)
                                        if ((v44 >> ((134) & 63)) < v44) {
                                            v52 := GenS5{F0: 1}
                                            v5 = "foo bar"
                                            v3 -= (^(v3))
                                            _ = v52
                                        } else {
                                            var v53 int16 = -1
                                            v6 = v6
                                            for fi54 := 0; fi54 < 5; fi54++ {
                                                v53 += v53
                                                _ = fi54
                                            }
                                            _ = v53
                                        }
                                    }
                                    switch (^(v27)) {
                                    case 757:
                                        v2 -= v2
                                    case 9223372036854775807:
                                        v5 = (v45 + "")
                                    case 127:
                                        v2 += (v2 % ((v2) | 1))
                                    }
                                    var v55 float64 = float64(((0.0 + 1.5) * (3.14159 / -2.25)))
                                    _ = fi51
                                    _ = v55
                                }
                            case 840:
                                delete(v4, v43)
                            case 42:
                                var v56 bool = (v6 && (!(71 < 42)))
                                _ = v56
                            }
                            _ = v50
                        } else {
                            v57 := []int{9223372036854775807, v27}
                            var v58 uint64 = 319
                            _ = len(v57)
                            _ = v57
                            _ = v58
                        }
                        switch v27 {
                        case 9223372036854775807:
                            switch v5 {
                            case "":
                                var v59 uint32 = 0
                                _ = v59
                            case "hello":
                                switch (^(v27)) {
                                case 449:
                                    for fi60 := 0; fi60 < 6; fi60++ {
                                        for fi61 := 0; fi61 < 14; fi61++ {
                                            fmt.Println(v2)
                                            v27 += fi61
                                            fmt.Println(v2, fi61, v43)
                                            _ = fi61
                                        }
                                        for fi62 := 0; fi62 < 6; fi62++ {
                                            var v63 int64 = 748
                                            fmt.Println(v63, fi60)
                                            fmt.Println(fi60, v6)
                                            _ = fi62
                                            _ = v63
                                        }
                                        _ = fi60
                                    }
                                case -1:
                                    var v64 int = genFn1(v2, (v27 ^ genFn1(4294967295, 2147483647, v1)), (v1 >> ((v1) & 63)))
                                    _ = v64
                                }
                            case "a":
                                v65 := v4[v2]
                                _ = v65
                            default:
                                v66 := GenS1{F0: 42, F1: v27}
                                _ = v66
                            }
                        case 128:
                            _ = len(v4)
                        case -827:
                            v67 := []int{}
                            _ = v67
                        default:
                            var v68 uint = 18446744073709551615
                            _ = v68
                        }
                        _ = v45
                    } else {
                        v69 := []int{v27}
                        fmt.Println(v44, v2)
                        _ = v69
                    }
                    var v70 string = (v5 + (v5 + v5))
                    if v6 {
                        v71 := v4[(v43 * 4294967295)]
                        _ = v71
                    }
                    _ = v70
                }
                switch "hello" {
                case "a":
                    fmt.Println(v3)
                case "":
                    fmt.Println(v43, v5)
                case "hello":
                    if (!v6) {
                        if false {
                            v72 := GenS5{F0: 65535}
                            _ = v72
                        } else {
                            for fi73 := 0; fi73 < 13; fi73++ {
                                v74 := v4[v43]
                                _ = fi73
                                _ = v74
                            }
                            v75 := GenS6{F0: 42, F1: 65535, F2: (v44 | v44)}
                            fmt.Println(v6, v1)
                            _ = v75
                        }
                        var v76 uint64 = 4294967295
                        _ = v76
                    } else {
                        for fi77 := 0; fi77 < 18; fi77++ {
                            v43 -= 2147483647
                            switch v5 {
                            case "a":
                                for fi78 := 0; fi78 < 16; fi78++ {
                                    var v79 uint32 = (v2 + (^(v43)))
                                    switch genFn1(4294967295, 630, (v1 - v1)) {
                                    case 127:
                                        var v80 int8 = 42
                                        _ = v80
                                    case 835:
                                        var v81 bool = (-306 <= -1)
                                        var v82 uint16 = 42
                                        switch (v5 + "\n\t") {
                                        case "\n\t":
                                            var v83 int32 = 336
                                            _ = v83
                                        case "foo bar":
                                            switch (v27 >> ((255) & 63)) {
                                            case 973:
                                                v84 := v4[v79]
                                                _ = v84
                                            case 9223372036854775807:
                                                v43 += v2
                                            case 127:
                                                switch genFn1((v2 - v43), fi78, v1) {
                                                case 9223372036854775807:
                                                    var v85 bool = (4294967295 < 1)
                                                    _ = v85
                                                case 127:
                                                    v86 := v4[(v43 * v79)]
                                                    _ = v86
                                                case 2147483647:
                                                    v87 := GenS7{F0: v44, F1: (v27 << (((v1 & (v1 & v1))) & 63)), F2: (v44 | (v44 - v44))}
                                                    _ = v87
                                                default:
                                                    switch ("foo bar" + "foo bar") {
                                                    case "hello":
                                                        var v88 int32 = -1
                                                        _ = v88
                                                    case "foo bar":
                                                        v43 = v79
                                                        v89 := GenS8{F0: v5, F1: 4294967295}
                                                        _ = v89
                                                    case "a":
                                                        v27 = 0
                                                        v43 -= (^(v79))
                                                        fmt.Println(v79)
                                                    default:
                                                        v44 += (v44 * 128)
                                                    }
                                                }
                                            }
                                        }
                                        _ = v81
                                        _ = v82
                                    case 128:
                                        v43 = v79
                                    default:
                                        var v90 uint64 = uint64(v2)
                                        _ = v90
                                    }
                                    _ = fi78
                                    _ = v79
                                }
                            case "foo bar":
                                v44 += 1
                            case "":
                                if (v6 && (!true)) {
                                    v1 += (v1 & v1)
                                    for fi91 := 0; fi91 < 7; fi91++ {
                                        if (!v6) {
                                            v92 := GenS3{F0: 18446744073709551615}
                                            switch "foo bar" {
                                            case "":
                                                var v93 float32 = (float32(0.0) + -2.25)
                                                _ = v93
                                            case "hello":
                                                var v94 bool = (v6 || (false && (false && v6)))
                                                _ = v94
                                            case "a":
                                                v95 := v4[v43]
                                                _ = v95
                                            }
                                            _ = v92
                                        }
                                        v96 := []int{93, fi91, (fi77 >> ((0) & 63))}
                                        _ = fi91
                                        _ = v96
                                    }
                                    v97 := []int{v27, v27, v27, v27}
                                    _ = v97
                                } else {
                                    if (!(42 < 23)) {
                                        v98 := []int{int(v3), (v27 - v27), genFn1(130, -269, (v1 * 241))}
                                        _ = v98
                                    }
                                    if (v6 && ((!(fi77 <= 1)) && false)) {
                                        if v6 {
                                            v3 = (v3 * int64(v27))
                                        } else {
                                            v99 := GenS9{F0: v5}
                                            var v100 float64 = (1e10 - 0.0)
                                            v27 += (fi77 >> (((v1 % ((v1) | 1))) & 63))
                                            _ = v100
                                            _ = v99
                                        }
                                    } else {
                                        v101 := []int{}
                                        _ = v101
                                    }
                                    fmt.Println(v27, v43, v5)
                                }
                            default:
                                fmt.Println(fi77, v1, v27)
                            }
                            _ = fi77
                        }
                        if (1e10 <= 3.14159) {
                            switch v27 {
                            case 4294967295:
                                v27 = 2147483647
                            case -2147483648:
                                v102 := v4[v2]
                                _ = v102
                            case 32768:
                                v103 := []int{v27, (v27 - genFn1((v2 - 0), v27, (v1 / ((65535) | 1)))), v27}
                                _ = v103
                            }
                            if (2 > 32767) {
                                var v104 uint64 = 308
                                fmt.Println(v44, v104)
                                var v105 int16 = -32768
                                _ = v104
                                _ = v105
                            } else {
                                v106 := GenS3{F0: 1}
                                var v107 float32 = (1.5 + float32(0.0))
                                _ = v106
                                _ = v107
                            }
                            if ((v44 + (v44 | (v44 << (((^(v1))) & 63)))) <= (v44 | (v44 + uint8(v1)))) {
                                v108 := []int{(v27 % ((v27) | 1)), -867, genFn1(v2, 32768, 18446744073709551615), genFn1(v2, v27, (v1 << ((977) & 63)))}
                                if len(v108) > 0 {
                                    v108[0] = (v27 % ((genFn1(979, (v27 % ((v27) | 1)), v1)) | 1))
                                }
                                _ = v108
                            }
                        } else {
                            v2 += 765
                            var v109 string = v5
                            for fi110 := 0; fi110 < 8; fi110++ {
                                if v6 {
                                    fmt.Println(v3, v2, v27)
                                }
                                _ = fi110
                            }
                            _ = v109
                        }
                    }
                default:
                    fmt.Println(v27)
                }
                _ = v44
            } else {
                v111 := GenS2{F0: uint8(v1), F1: 42}
                v6 = (v5 < v5)
                _ = v111
            }
            v112 := []int{-9223372036854775808, (v27 ^ 2147483647), genFn1((v2 & v2), (v27 + v27), (^(v1))), v27}
            _ = v112
        }
    }
    var v113 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v113
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
