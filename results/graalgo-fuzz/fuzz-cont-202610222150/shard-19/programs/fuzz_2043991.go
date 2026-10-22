package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int32; F1 int64; F2 bool }
type GenS5 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS6 struct { F0 uint8; F1 int32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS9 struct { F0 int8; F1 uint32; F2 bool }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int32; F1 int }
type GenS13 struct { F0 uint32; F1 string }
type GenS14 struct { F0 string; F1 int16; F2 int8 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int8; F1 int16 }
type GenS18 struct { F0 uint64; F1 int16; F2 int64 }
type GenS19 struct { F0 string; F1 uint32 }
type GenS20 struct { F0 bool; F1 int }
type GenS21 struct { F0 bool; F1 uint8; F2 int }
func genFn0() uint16 {
    if false {
        var v7 string = "hello"
        _ = v7
    }
    var v8 uint32 = 580
    var v9 int32 = 2147483647
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 int64) int64 {
    v10 := GenS1{F0: -2147483648, F1: 0, F2: -2147483648}
    var v11 uint64 = 42
    switch 128 {
    case 32767:
        var v12 bool = true
        _ = v12
    default:
        fmt.Println(p0, v11)
    }
    _ = p0
    _ = v10
    _ = v11
    return int64(v11)
}

func main() {
    var v1 bool = (660 <= 704)
    v2 := []int{599, -2147483648}
    var v3 bool = v1
    v4 := []int{}
    var v5 int8 = 0
    v6 := make(map[int]string)
    if v3 {
        var v13 int16 = 42
        var v14 uint8 = 127
        _ = v13
        _ = v14
    } else {
        var v15 float32 = ((1e10 - (0.1 / 1.0)) / -2.25)
        if ((42 < 150) || (("foo bar" == (("" + "foo bar") + ("foo bar" + ""))) || false)) {
            for fi16 := 0; fi16 < 11; fi16++ {
                fmt.Println(fi16)
                v2 = append(v2, fi16)
                switch "foo bar" {
                case "foo bar":
                    for fi17 := 0; fi17 < 1; fi17++ {
                        var v18 bool = ((v5 < v5) || v3)
                        switch ("\n\t" + ("hello" + (("foo bar" + ("foo bar" + ("a" + "hello"))) + "\n\t"))) {
                        case "foo bar":
                            fmt.Println(fi17, v5, fi16)
                        case "hello":
                            fmt.Println(v1, fi17, v15)
                        default:
                            var v19 uint64 = 42
                            _ = v19
                        }
                        v18 = (1 < 9223372036854775807)
                        _ = fi17
                        _ = v18
                    }
                case "\n\t":
                    for fi20 := 0; fi20 < 3; fi20++ {
                        if v3 {
                            fmt.Println(fi16)
                        }
                        _ = fi20
                    }
                }
                _ = fi16
            }
            delete(v6, 127)
            _ = len(v2)
        } else {
            if (false && (0 >= 855)) {
                _ = len(v6)
            } else {
                fmt.Println(v5)
            }
            var v21 int32 = 625
            v22 := v6[127]
            _ = v21
            _ = v22
        }
        _ = v15
    }
    switch -772 {
    case 4294967295:
        var v23 uint = 0
        _ = v23
    case 127:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -826:
        v3 = (1 > 0)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v24 int64 = 9223372036854775807
    fmt.Println(v24, v1, v3)
    fmt.Println(v24, v24)
    if len(v2) > 0 {
        v2[0] = -424
    }
    v5 = (v5 + (v5 + v5))
    if false {
        var v25 int16 = -657
        switch 2147483647 {
        case 2147483647:
            v3 = (0.1 >= 1.5)
        }
        for fi26 := 0; fi26 < 16; fi26++ {
            if v1 {
                _ = len(v4)
                v27 := GenS1{F0: -782, F1: 295, F2: 1}
                _ = v27
            } else {
                v28 := GenS1{F0: -810, F1: 184, F2: 42}
                fmt.Println(v25)
                if ((-2.25 + ((1.5 * (0.1 - float32(1.5))) - -2.25)) >= (float32(1.5) - float32(((1e10 / 1.0) + 0.1)))) {
                    delete(v6, fi26)
                }
                _ = v28
            }
            var v29 int16 = (v25 | (v25 * v25))
            _ = fi26
            _ = v29
        }
        _ = v25
    } else {
        if (!false) {
            fmt.Println(v1, v3)
            var v30 uint16 = 786
            v31 := GenS1{F0: -2147483648, F1: 478, F2: 934}
            _ = v30
            _ = v31
        } else {
            if v1 {
                for fi32 := 0; fi32 < 1; fi32++ {
                    if false {
                        fmt.Println(v3, v1)
                        _ = len(v2)
                    }
                    switch ("" + ((("hello" + "") + ("" + "")) + "hello")) {
                    case "hello":
                        for fi33 := 0; fi33 < 17; fi33++ {
                            if true {
                                if (0.0 > 1e10) {
                                    if v1 {
                                        v24 = genFn1(genFn1(v24))
                                        switch (fi32 | (fi32 / (((-(fi33))) | 1))) {
                                        case 127:
                                            v24 -= v24
                                            if (!(3.14159 == (0.1 / float32(float64((float32((-2.25 / float64(0.1))) / 1.0)))))) {
                                                v3 = (!(1.5 != (((float32(3.14159) * 0.0) * 3.14159) / 1e10)))
                                                var v34 uint16 = 2
                                                fmt.Println(v3, fi32, fi32)
                                                if v1 {
                                                    var v35 uint16 = (^(v34))
                                                    switch fi32 {
                                                    case 9223372036854775807:
                                                        v34 = genFn0()
                                                        var v36 uint = 18446744073709551615
                                                        _ = v36
                                                    case 32768:
                                                        var v37 int32 = -972
                                                        _ = v37
                                                    case -575:
                                                        v38 := GenS1{F0: 0, F1: 246, F2: -163}
                                                        _ = v38
                                                    default:
                                                        switch "hello" {
                                                        case "a":
                                                            v39 := GenS2{F0: int16(fi32)}
                                                            if (((54 <= 65535) && false) || ((!(45 == (v35 ^ genFn0()))) && v3)) {
                                                                if true {
                                                                    fmt.Println(v34)
                                                                    switch ("foo bar" + "") {
                                                                    case "\n\t":
                                                                        v24 -= v24
                                                                    case "a":
                                                                        _ = len(v2)
                                                                    default:
                                                                        var v40 int16 = 0
                                                                        _ = v40
                                                                    }
                                                                }
                                                                if (fi33 == fi33) {
                                                                    if false {
                                                                        var v41 float64 = 0.0
                                                                        v41 -= ((-2.25 + ((v41 / ((((0.0 - v41) * v41) * float64(float32(-2.25))) / 1.0)) * 3.14159)) - (v41 - float64(1.5)))
                                                                        v42 := GenS3{F0: 4294967295, F1: 2147483647}
                                                                        var v43 uint32 = 2147483648
                                                                        _ = v41
                                                                        _ = v42
                                                                        _ = v43
                                                                    }
                                                                    var v44 uint64 = 4294967295
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        var v45 uint16 = v35
                                                                        _ = v45
                                                                    default:
                                                                        if (341 >= 557) {
                                                                            v2 = append(v2, -1)
                                                                            fmt.Println(v44, v1, v34)
                                                                            switch 4294967295 {
                                                                            case 42:
                                                                                var v46 string = "foo bar"
                                                                                _ = v46
                                                                            case 887:
                                                                                v4 = append(v4, (fi33 % ((-259) | 1)))
                                                                            case 9223372036854775807:
                                                                                if ((fi33 ^ (fi32 - fi33)) < fi32) {
                                                                                    var v47 int64 = v24
                                                                                    if v3 {
                                                                                        if true {
                                                                                            if v1 {
                                                                                                if v1 {
                                                                                                    var v48 int16 = -1
                                                                                                    _ = v48
                                                                                                } else {
                                                                                                    v6[fi32] = "\n\t"
                                                                                                    v49 := GenS2{F0: -977}
                                                                                                    _ = v49
                                                                                                }
                                                                                            }
                                                                                            if ((true || v3) && ((1.5 + (0.0 - (3.14159 - (1e10 / (((0.0 * (0.0 * 1.5)) / -2.25) + float64(float32(0.1))))))) <= 1.5)) {
                                                                                                var v50 int = int(v5)
                                                                                                if v1 {
                                                                                                    if len(v4) > 0 {
                                                                                                        v4[0] = v50
                                                                                                    }
                                                                                                    v6[(fi33 | fi32)] = ""
                                                                                                    if (42 >= 1) {
                                                                                                        v51 := GenS3{F0: 4294967295, F1: 2147483648}
                                                                                                        var v52 bool = false
                                                                                                        _ = v51
                                                                                                        _ = v52
                                                                                                    }
                                                                                                }
                                                                                                _ = v50
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v53 := v6[fi33]
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        var v54 uint32 = 42
                                                                                        _ = v54
                                                                                    }
                                                                                    v55 := GenS4{F0: 0, F1: (v24 + 42), F2: (2147483647 <= 100)}
                                                                                    switch fi33 {
                                                                                    case -2147483648:
                                                                                        switch (fi32 | -136) {
                                                                                        case 0:
                                                                                            switch ((((v53 + v53) + (("hello" + "") + "")) + ("a" + (v53 + v53))) + v53) {
                                                                                            case "hello":
                                                                                                v34 += 111
                                                                                            case "a":
                                                                                                v56 := v6[(fi32 % (((^(fi33))) | 1))]
                                                                                                _ = v56
                                                                                            default:
                                                                                                var v57 uint64 = (v44 >> ((0) & 63))
                                                                                                _ = v57
                                                                                            }
                                                                                        case 127:
                                                                                            if (!(v3 && (!v3))) {
                                                                                                v58 := GenS5{F0: 0, F1: (v44 + (v44 * v44)), F2: uint16(v44)}
                                                                                                v53 = "foo bar"
                                                                                                if true {
                                                                                                    fmt.Println(v24, v53)
                                                                                                    switch ("" + (v53 + ((v53 + "a") + "hello"))) {
                                                                                                    case "":
                                                                                                        v59 := GenS6{F0: 46, F1: 0}
                                                                                                        _ = v59
                                                                                                    }
                                                                                                }
                                                                                                switch fi33 {
                                                                                                case 42:
                                                                                                    v5 = (v5 / (((v5 ^ v5)) | 1))
                                                                                                case 2147483647:
                                                                                                    var v60 float64 = 0.0
                                                                                                    _ = v60
                                                                                                default:
                                                                                                    if true {
                                                                                                        _ = len(v6)
                                                                                                        var v61 int16 = -675
                                                                                                        if v1 {
                                                                                                            var v62 uint32 = 638
                                                                                                            v63 := GenS4{F0: 0, F1: 9223372036854775807, F2: true}
                                                                                                            _ = v62
                                                                                                            _ = v63
                                                                                                        } else {
                                                                                                            if (127 == v5) {
                                                                                                                if v1 {
                                                                                                                    var v64 uint = 65535
                                                                                                                    v2 = append(v2, fi33)
                                                                                                                    var v65 int = int(v64)
                                                                                                                    _ = v64
                                                                                                                    _ = v65
                                                                                                                }
                                                                                                                var v66 float32 = 1e10
                                                                                                                _ = v66
                                                                                                            } else {
                                                                                                                if v3 {
                                                                                                                    if v1 {
                                                                                                                        switch (-(fi32)) {
                                                                                                                        case 4294967295:
                                                                                                                            var v67 int16 = (-(v61))
                                                                                                                            _ = v67
                                                                                                                        case -1:
                                                                                                                            if true {
                                                                                                                                fmt.Println(v53, v53, v5)
                                                                                                                                var v68 int64 = int64(v61)
                                                                                                                                v69 := GenS7{F0: v5}
                                                                                                                                _ = v68
                                                                                                                                _ = v69
                                                                                                                            }
                                                                                                                        case 298:
                                                                                                                            v70 := GenS4{F0: 2147483647, F1: (v24 - (v24 - (v47 / ((v24) | 1)))), F2: v3}
                                                                                                                            _ = v70
                                                                                                                        default:
                                                                                                                            var v71 string = "foo bar"
                                                                                                                            _ = v71
                                                                                                                        }
                                                                                                                        v1 = v1
                                                                                                                        v44 += 483
                                                                                                                    } else {
                                                                                                                        v72 := v6[9223372036854775807]
                                                                                                                        v73 := GenS1{F0: 2147483647, F1: (v44 << ((2) & 63)), F2: 42}
                                                                                                                        if true {
                                                                                                                            var v74 float32 = ((float32(0.1) + 3.14159) / (((1e10 * 1.5) * 0.1) + 1.5))
                                                                                                                            if ((v61 | (v61 >> ((255) & 63))) < v61) {
                                                                                                                                fmt.Println(v61, v44, v53)
                                                                                                                                if (0.0 == (1e10 - -2.25)) {
                                                                                                                                    var v75 uint32 = uint32(v44)
                                                                                                                                    _ = len(v6)
                                                                                                                                    _ = v75
                                                                                                                                } else {
                                                                                                                                    v76 := GenS8{F0: 746, F1: 0, F2: 186}
                                                                                                                                    _ = v76
                                                                                                                                }
                                                                                                                                v72 = ((("hello" + v53) + v72) + v53)
                                                                                                                            }
                                                                                                                            v24 += v47
                                                                                                                            _ = v74
                                                                                                                        } else {
                                                                                                                            if v3 {
                                                                                                                                var v77 uint64 = (^(v44))
                                                                                                                                fmt.Println(v24, v47, v77)
                                                                                                                                _ = v77
                                                                                                                            } else {
                                                                                                                                var v78 uint64 = (v44 - (v44 << ((18446744073709551615) & 63)))
                                                                                                                                switch fi33 {
                                                                                                                                case -9223372036854775808:
                                                                                                                                    if true {
                                                                                                                                        v79 := v6[fi33]
                                                                                                                                        v47 = v47
                                                                                                                                        _ = v79
                                                                                                                                    }
                                                                                                                                case 32768:
                                                                                                                                    v35 += (v35 - 0)
                                                                                                                                }
                                                                                                                                v80 := GenS2{F0: (^(v61))}
                                                                                                                                _ = v78
                                                                                                                                _ = v80
                                                                                                                            }
                                                                                                                            v24 = 178
                                                                                                                        }
                                                                                                                        _ = v72
                                                                                                                        _ = v73
                                                                                                                    }
                                                                                                                }
                                                                                                                var v81 string = ""
                                                                                                                fmt.Println(v53)
                                                                                                                _ = v81
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v61
                                                                                                    } else {
                                                                                                        v35 += genFn0()
                                                                                                        var v82 int = (fi33 & -2147483648)
                                                                                                        _ = len(v6)
                                                                                                        _ = v82
                                                                                                    }
                                                                                                }
                                                                                                if v3 {
                                                                                                    v3 = false
                                                                                                    v34 += v35
                                                                                                    var v83 uint16 = v34
                                                                                                    if (v44 >= (v44 & 1)) {
                                                                                                        var v84 int64 = v24
                                                                                                        var v85 string = ""
                                                                                                        _ = v84
                                                                                                        _ = v85
                                                                                                    }
                                                                                                    _ = v83
                                                                                                }
                                                                                                _ = v58
                                                                                            }
                                                                                            var v86 int8 = -128
                                                                                            var v87 int16 = 32767
                                                                                            _ = v86
                                                                                            _ = v87
                                                                                        case 9223372036854775807:
                                                                                            v5 -= v5
                                                                                        default:
                                                                                            fmt.Println(fi33)
                                                                                        }
                                                                                    case 32768:
                                                                                        v47 += (v24 / ((genFn1((v47 * (v47 % ((v47) | 1))))) | 1))
                                                                                    }
                                                                                    _ = v47
                                                                                    _ = v53
                                                                                    _ = v55
                                                                                }
                                                                            default:
                                                                                v24 -= (v24 - int64(v44))
                                                                            }
                                                                        }
                                                                    }
                                                                    switch fi33 {
                                                                    case 0:
                                                                        var v88 int64 = 0
                                                                        _ = v88
                                                                    }
                                                                    switch 127 {
                                                                    case 9223372036854775807:
                                                                        var v89 uint64 = (v44 * v44)
                                                                        _ = v89
                                                                    case -355:
                                                                        switch fi32 {
                                                                        case 1:
                                                                            v90 := GenS7{F0: (v5 + (v5 | (^(v5))))}
                                                                            _ = v90
                                                                        case -52:
                                                                            v35 = v34
                                                                        case 396:
                                                                            var v91 string = "a"
                                                                            _ = v91
                                                                        }
                                                                    default:
                                                                        if ((!((!(215 > 42)) || v3)) || false) {
                                                                            if ((fi32 != fi32) || v1) {
                                                                                var v92 string = "hello"
                                                                                var v93 float64 = (0.0 * -2.25)
                                                                                var v94 int64 = (v24 >> ((2) & 63))
                                                                                _ = v92
                                                                                _ = v93
                                                                                _ = v94
                                                                            }
                                                                            switch (fi33 + fi33) {
                                                                            case 1:
                                                                                var v95 int64 = v24
                                                                                _ = v95
                                                                            case -2147483648:
                                                                                switch fi32 {
                                                                                case -9223372036854775808:
                                                                                    v3 = v3
                                                                                    v6[(fi33 ^ fi33)] = "a"
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = 42
                                                                                    }
                                                                                case 32767:
                                                                                    v6[fi33] = "foo bar"
                                                                                }
                                                                            case 42:
                                                                                var v96 string = "\n\t"
                                                                                _ = v96
                                                                            default:
                                                                                var v97 int64 = (v24 % ((v24) | 1))
                                                                                _ = v97
                                                                            }
                                                                            v24 = v24
                                                                        }
                                                                    }
                                                                    _ = v44
                                                                }
                                                            }
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            _ = v39
                                                        case "":
                                                            v98 := GenS9{F0: 0, F1: 2, F2: (42 != 255)}
                                                            _ = v98
                                                        }
                                                        v24 += (v24 | v24)
                                                    }
                                                    switch ("foo bar" + ("foo bar" + "foo bar")) {
                                                    case "a":
                                                        v35 += 42
                                                    case "foo bar":
                                                        switch int(v24) {
                                                        case 629:
                                                            var v99 int16 = 32767
                                                            _ = v99
                                                        default:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                    case "":
                                                        switch "hello" {
                                                        case "":
                                                            var v100 uint8 = 255
                                                            _ = v100
                                                        case "hello":
                                                            v5 = (v5 % ((0) | 1))
                                                        case "\n\t":
                                                            var v101 string = "\n\t"
                                                            _ = v101
                                                        }
                                                    }
                                                    var v102 uint = 65535
                                                    fmt.Println(v35, v35, fi33)
                                                    _ = v102
                                                    _ = v35
                                                }
                                                v3 = (v5 != v5)
                                                if v3 {
                                                    v103 := v6[fi33]
                                                    _ = v103
                                                } else {
                                                    if v3 {
                                                        fmt.Println(v1)
                                                        fmt.Println(v24)
                                                        if (-2.25 > 1.5) {
                                                            var v104 int = fi33
                                                            v3 = true
                                                            _ = v104
                                                        } else {
                                                            v34 = (v34 ^ 2)
                                                            if (-2147483648 <= -2147483648) {
                                                                var v105 uint64 = 2
                                                                if false {
                                                                    v106 := GenS8{F0: 18446744073709551615, F1: 2147483648, F2: 133}
                                                                    _ = v106
                                                                } else {
                                                                    fmt.Println(v1, fi33, fi32)
                                                                }
                                                                v5 -= (-(v5))
                                                                _ = v105
                                                            } else {
                                                                switch (fi32 - (fi33 | fi32)) {
                                                                case 32768:
                                                                    var v107 float32 = (float32(((1.5 * (0.1 / 1e10)) * 0.0)) * (3.14159 * 1.5))
                                                                    _ = v107
                                                                case -2147483648:
                                                                    var v108 uint32 = 2147483648
                                                                    _ = v108
                                                                default:
                                                                    fmt.Println(fi32, v34, v1)
                                                                }
                                                                if v3 {
                                                                    v6[(fi33 % (((^(fi32))) | 1))] = "a"
                                                                    v24 -= int64(fi33)
                                                                }
                                                                switch (fi32 ^ (fi32 >> ((527) & 63))) {
                                                                case 171:
                                                                    var v109 uint = 1
                                                                    _ = v109
                                                                case 2147483647:
                                                                    switch (fi32 * fi33) {
                                                                    case 203:
                                                                        if true {
                                                                            var v110 int8 = 0
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                v111 := GenS10{F0: 4294967295}
                                                                                _ = v111
                                                                            default:
                                                                                v112 := GenS5{F0: 255, F1: 380, F2: genFn0()}
                                                                                _ = v112
                                                                            }
                                                                            fmt.Println(v3, v24, v1)
                                                                            if false {
                                                                                var v113 uint64 = 1
                                                                                _ = v113
                                                                            }
                                                                            _ = v110
                                                                        } else {
                                                                            if true {
                                                                                v3 = v1
                                                                                fmt.Println(v24, v3, fi32)
                                                                                v4 = append(v4, 128)
                                                                                var v114 uint16 = (v34 << ((42) & 63))
                                                                                if true {
                                                                                    var v115 uint32 = 1
                                                                                    v114 = (v114 | v34)
                                                                                    var v116 uint = 839
                                                                                    var v117 uint32 = (v115 * v115)
                                                                                    _ = v115
                                                                                    _ = v116
                                                                                    _ = v117
                                                                                } else {
                                                                                    if v3 {
                                                                                        delete(v6, 9223372036854775807)
                                                                                        v6[4294967295] = ""
                                                                                        switch "\n\t" {
                                                                                        case "hello":
                                                                                            _ = len(v6)
                                                                                        case "\n\t":
                                                                                            v3 = (v1 && false)
                                                                                        case "a":
                                                                                            _ = len(v6)
                                                                                        }
                                                                                        fmt.Println(v114, v114)
                                                                                        delete(v6, fi33)
                                                                                    } else {
                                                                                        var v118 int8 = (v5 + (v5 - (v5 * v5)))
                                                                                        if (genFn1(1) > 9223372036854775807) {
                                                                                            switch "\n\t" {
                                                                                            case "foo bar":
                                                                                                var v119 uint = 65535
                                                                                                _ = v119
                                                                                            }
                                                                                        }
                                                                                        var v120 int8 = 0
                                                                                        _ = v118
                                                                                        _ = v120
                                                                                    }
                                                                                    if true {
                                                                                        if v3 {
                                                                                            v1 = v3
                                                                                            var v121 float64 = 0.0
                                                                                            var v122 int8 = v5
                                                                                            var v123 uint8 = 2
                                                                                            _ = v121
                                                                                            _ = v122
                                                                                            _ = v123
                                                                                        }
                                                                                        v114 = genFn0()
                                                                                    }
                                                                                }
                                                                                _ = len(v4)
                                                                                _ = v114
                                                                            }
                                                                            if false {
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    fmt.Println(v24, v34)
                                                                                default:
                                                                                    v4 = append(v4, (-(fi32)))
                                                                                }
                                                                            } else {
                                                                                if v3 {
                                                                                    if (((0.1 > 0.1) && (v3 && (!(v24 >= (-(v24)))))) || (true && (1 > -2147483648))) {
                                                                                        var v124 uint64 = 4294967295
                                                                                        var v125 uint64 = (v124 | v124)
                                                                                        v126 := GenS10{F0: 13}
                                                                                        _ = v124
                                                                                        _ = v125
                                                                                        _ = v126
                                                                                    } else {
                                                                                        v24 -= (v24 % ((genFn1(v24)) | 1))
                                                                                    }
                                                                                    if (!((!(v3 || v1)) || (0 > 1))) {
                                                                                        if v1 {
                                                                                            v127 := GenS1{F0: int32(v24), F1: 4294967295, F2: 42}
                                                                                            _ = v127
                                                                                        }
                                                                                        v128 := GenS11{F0: 0}
                                                                                        v24 += (v24 - v24)
                                                                                        _ = v128
                                                                                    } else {
                                                                                        v129 := GenS12{F0: 243, F1: (fi32 + (^(fi33)))}
                                                                                        _ = v129
                                                                                    }
                                                                                } else {
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                }
                                                                                v130 := GenS13{F0: 65535, F1: ("" + "foo bar")}
                                                                                _ = v130
                                                                            }
                                                                        }
                                                                    case 9223372036854775807:
                                                                        v5 -= v5
                                                                    case -751:
                                                                        if (!v1) {
                                                                            var v131 string = ("\n\t" + "")
                                                                            _ = v131
                                                                        } else {
                                                                            v132 := GenS14{F0: "hello", F1: 341, F2: v5}
                                                                            if (v1 && v3) {
                                                                                v133 := GenS12{F0: 0, F1: int(v34)}
                                                                                fmt.Println(v34, v5)
                                                                                v3 = (4294967295 == 18446744073709551615)
                                                                                _ = v133
                                                                            }
                                                                            _ = v132
                                                                        }
                                                                    default:
                                                                        v24 = v24
                                                                    }
                                                                }
                                                                v3 = (v3 || v1)
                                                                var v134 int = 9223372036854775807
                                                                fmt.Println(fi33, fi33)
                                                                fmt.Println(fi32)
                                                                v135 := GenS15{F0: 0}
                                                                var v136 uint16 = (v34 - 65535)
                                                                _ = v134
                                                                _ = v135
                                                                _ = v136
                                                            }
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                    }
                                                }
                                                v4 = append(v4, (fi33 << ((4294967295) & 63)))
                                                _ = v34
                                            } else {
                                                v6[(^(fi32))] = "foo bar"
                                            }
                                            var v137 uint8 = 1
                                            var v138 int = (fi32 - fi32)
                                            _ = v137
                                            _ = v138
                                        case 32767:
                                            switch -26 {
                                            case 0:
                                                if (((0.1 / 1.0) < 3.14159) && (42 > 1)) {
                                                    if (v1 && (!(-2.25 <= 1.5))) {
                                                        v1 = (0.0 >= (3.14159 + -2.25))
                                                        v139 := GenS16{F0: (!v3)}
                                                        _ = v139
                                                    } else {
                                                        v24 = 919
                                                    }
                                                }
                                            case -418:
                                                v5 = (v5 >> ((2) & 63))
                                            case 2147483647:
                                                switch ("a" + "hello") {
                                                case "\n\t":
                                                    if ((0 != -32768) || v1) {
                                                        if (!((float32(3.14159) / 3.14159) >= (-2.25 * 3.14159))) {
                                                            v1 = (true && (0.1 >= 1.5))
                                                        }
                                                        if v3 {
                                                            var v140 float64 = (-2.25 + (((0.0 + (float64(-2.25) / 1.5)) / -2.25) - 0.1))
                                                            _ = v140
                                                        }
                                                        var v141 uint32 = 682
                                                        _ = v141
                                                    } else {
                                                        switch "\n\t" {
                                                        case "a":
                                                            v5 = v5
                                                        case "":
                                                            v3 = false
                                                        }
                                                        fmt.Println(v24, v24)
                                                        if (2 > 127) {
                                                            var v142 uint16 = 1
                                                            switch (-(fi33)) {
                                                            case -357:
                                                                var v143 int64 = -873
                                                                _ = v143
                                                            }
                                                            v144 := GenS10{F0: 953}
                                                            _ = v142
                                                            _ = v144
                                                        }
                                                    }
                                                default:
                                                    if (128 < 72) {
                                                        delete(v6, -1)
                                                        var v145 int16 = -1
                                                        _ = v145
                                                    }
                                                }
                                            }
                                        case 0:
                                            switch "hello" {
                                            case "foo bar":
                                                v146 := v6[fi32]
                                                _ = v146
                                            case "a":
                                                var v147 int8 = (v5 - -5)
                                                _ = v147
                                            case "hello":
                                                if (!(42 > 255)) {
                                                    v1 = v3
                                                    if v3 {
                                                        v148 := GenS17{F0: (v5 % (((-(v5))) | 1)), F1: -1}
                                                        v6[(fi32 ^ 591)] = "hello"
                                                        _ = v148
                                                    } else {
                                                        if (0 > v5) {
                                                            v5 -= v5
                                                            v1 = true
                                                            v1 = ((3.14159 - 1.5) >= (1e10 + 0.0))
                                                        }
                                                        if (3.14159 < 3.14159) {
                                                            var v149 int64 = (v24 >> ((0) & 63))
                                                            v1 = ((v24 + (v24 % ((v149) | 1))) != v149)
                                                            _ = v149
                                                        }
                                                    }
                                                }
                                                var v150 uint32 = 2
                                                _ = v150
                                            }
                                        }
                                        switch ("a" + (("a" + "hello") + "a")) {
                                        case "\n\t":
                                            v3 = v1
                                        case "foo bar":
                                            var v151 int = fi32
                                            _ = v151
                                        case "hello":
                                            v24 = genFn1(v24)
                                        }
                                    }
                                    switch fi32 {
                                    case -140:
                                        if v3 {
                                            fmt.Println(v1, v24)
                                            if (1000 <= 704) {
                                                v152 := GenS3{F0: 194, F1: 4294967295}
                                                _ = v152
                                            }
                                        } else {
                                            v5 = (v5 / (((-(v5))) | 1))
                                            v153 := GenS18{F0: 552, F1: 1, F2: v24}
                                            v1 = v3
                                            _ = v153
                                        }
                                    case 4294967295:
                                        var v154 float32 = 1e10
                                        _ = v154
                                    }
                                    if (v1 && false) {
                                        v1 = ((127 > (v5 | v5)) && (42 >= 42))
                                        v155 := v6[-784]
                                        var v156 int = fi33
                                        _ = v155
                                        _ = v156
                                    }
                                }
                                var v157 int64 = genFn1((v24 * 9223372036854775807))
                                _ = v157
                            }
                            _ = fi33
                        }
                    case "a":
                        for fi158 := 0; fi158 < 5; fi158++ {
                            v159 := GenS19{F0: "a", F1: 2}
                            fmt.Println(v24, fi32)
                            _ = fi158
                            _ = v159
                        }
                    case "":
                        var v160 bool = v3
                        _ = v160
                    default:
                        var v161 uint8 = 0
                        _ = v161
                    }
                    _ = fi32
                }
                var v162 int8 = -96
                v1 = (v3 || (v1 && v3))
                _ = v162
            }
            var v163 bool = v3
            _ = v163
        }
        v164 := v6[32767]
        _ = v164
    }
    v2 = append(v2, -92)
    var v165 float64 = 1e10
    for fi166 := 0; fi166 < 11; fi166++ {
        switch ("" + "\n\t") {
        case "hello":
            var v167 uint64 = 935
            _ = v167
        default:
            if (fi166 <= fi166) {
                var v168 uint = 0
                _ = v168
            }
        }
        _ = fi166
    }
    v169 := GenS20{F0: true, F1: 42}
    if false {
        var v170 uint16 = genFn0()
        var v171 int8 = (v5 << ((483) & 63))
        _ = v170
        _ = v171
    } else {
        v172 := GenS21{F0: v1, F1: 0, F2: int(v5)}
        _ = v172
    }
    var v173 int8 = (v5 / ((v5) | 1))
    fmt.Println(v173, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v165:", v165)
    fmt.Println("v173:", v173)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v165
    _ = v169
    _ = v173
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
