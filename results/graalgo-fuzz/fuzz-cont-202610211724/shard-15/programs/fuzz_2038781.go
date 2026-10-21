package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 uint; F2 int64 }
type GenS6 struct { F0 uint; F1 int; F2 uint64 }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 string; F1 int32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 bool; F1 uint }
type GenS11 struct { F0 int32; F1 int16; F2 int32 }
type GenS12 struct { F0 uint64; F1 int32; F2 bool }
type GenS13 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS14 struct { F0 bool; F1 int32; F2 uint }
type GenS15 struct { F0 int16; F1 bool }
type GenS16 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS17 struct { F0 uint8; F1 int64 }
type GenS18 struct { F0 bool; F1 bool }
func genFn0(p0 uint, p1 bool, p2 uint8) int64 {
    var v6 uint = (^(p0))
    v7 := GenS1{F0: 1, F1: 18446744073709551615}
    p0 -= v6
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{(fi8 ^ 42), (fi8 | fi8), fi8}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -278
}
func genFn1(p0 uint64) uint32 {
    var v10 float64 = (float64(1e10) - 1e10)
    var v11 uint8 = 2
    _ = p0
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 uint8 = 2
    v2 := []int{-528}
    var v3 bool = (true && false)
    var v4 float64 = 0.1
    v5 := make(map[uint32]string)
    var v12 uint64 = 1
    v13 := GenS1{F0: 1, F1: (v12 ^ (v12 & (^(v12))))}
    switch ("foo bar" + "\n\t") {
    case "":
        if v3 {
            v1 = (^(v1))
            v3 = (false && v3)
            var v14 uint = 42
            _ = v14
        } else {
            fmt.Println(v1, v1)
            if ((-1 == 1) || true) {
                for fi15 := 0; fi15 < 8; fi15++ {
                    v4 += 3.14159
                    _ = fi15
                }
                v16 := GenS1{F0: 196, F1: (^(v12))}
                v17 := GenS1{F0: -9223372036854775808, F1: (v12 >> ((2) & 63))}
                _ = v16
                _ = v17
            }
        }
    case "\n\t":
        if true {
            v18 := GenS2{F0: -362, F1: 0}
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 int32 = -1
                v21 := GenS3{F0: (((("hello" + ("foo bar" + ("a" + ("foo bar" + "foo bar")))) + (("" + "hello") + "")) + ("foo bar" + ("a" + "foo bar"))) + "\n\t"), F1: 797}
                var v22 int = int(v20)
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            v1 -= v1
            _ = v18
        }
    }
    if ((!(v3 || (!(0 != 65535)))) || v3) {
        switch "foo bar" {
        case "":
            v12 = 42
        case "a":
            if (42 >= -75) {
                var v23 uint32 = 789
                if (v3 || ((v12 + v12) <= (v12 << ((838) & 63)))) {
                    for fi24 := 0; fi24 < 9; fi24++ {
                        for fi25 := 0; fi25 < 12; fi25++ {
                            if (v3 && (1 <= 854)) {
                                v1 = 128
                            } else {
                                if (!true) {
                                    var v26 int64 = -408
                                    if (("a" + "hello") == "a") {
                                        v27 := GenS1{F0: (v26 / (((v26 & v26)) | 1)), F1: (v12 * (v12 ^ (v12 % ((v12) | 1))))}
                                        _ = len(v2)
                                        v28 := GenS1{F0: v26, F1: (v12 % (((v12 ^ (v12 / ((2) | 1)))) | 1))}
                                        if true {
                                            if v3 {
                                                fmt.Println(v23, v26, v4)
                                                var v29 float32 = 1.5
                                                v29 = 1e10
                                                _ = v29
                                            }
                                            v26 = genFn0(880, v3, (v1 / ((0) | 1)))
                                            var v30 float32 = (0.1 / 1e10)
                                            _ = v30
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = (^(fi25))
                                        }
                                        _ = v27
                                        _ = v28
                                    } else {
                                        if v3 {
                                            if v3 {
                                                switch ("\n\t" + "foo bar") {
                                                case "\n\t":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                            } else {
                                                if true {
                                                    var v31 uint = 906
                                                    switch (fi24 - (fi25 & int(v23))) {
                                                    case 668:
                                                        if len(v2) > 0 {
                                                            v2[0] = fi24
                                                        }
                                                    }
                                                    _ = v31
                                                } else {
                                                    if (0 != (^(v1))) {
                                                        v3 = v3
                                                        v32 := GenS4{F0: 56}
                                                        _ = v32
                                                    } else {
                                                        if v3 {
                                                            switch 127 {
                                                            case 2147483647:
                                                                fmt.Println(v12, v12)
                                                            }
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        } else {
                                                            var v33 uint16 = 2
                                                            v3 = true
                                                            _ = v33
                                                        }
                                                    }
                                                }
                                                v34 := GenS2{F0: fi25, F1: 364}
                                                if true {
                                                    v4 = -2.25
                                                    if true {
                                                        var v35 int8 = 97
                                                        fmt.Println(v35, v4, v4)
                                                        _ = v35
                                                    } else {
                                                        fmt.Println(fi25, v12)
                                                    }
                                                    var v36 int8 = 42
                                                    _ = v36
                                                } else {
                                                    v26 = (v26 & (-(v26)))
                                                }
                                                _ = v34
                                            }
                                            v37 := GenS5{F0: 4294967295, F1: 4294967295, F2: int64(v12)}
                                            switch "hello" {
                                            case "":
                                                _ = len(v2)
                                            case "a":
                                                v38 := GenS4{F0: int8(fi25)}
                                                _ = v38
                                            case "foo bar":
                                                delete(v5, (v23 ^ genFn1(v12)))
                                            default:
                                                if true {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        if (!v3) {
                                                            var v39 uint8 = v1
                                                            _ = v39
                                                        } else {
                                                            switch fi25 {
                                                            case -931:
                                                                _ = len(v2)
                                                            default:
                                                                switch 405 {
                                                                case 42:
                                                                    v26 -= (v26 + v26)
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        v1 = (v1 / (((^(v1))) | 1))
                                                    }
                                                    var v40 int64 = 0
                                                    _ = v40
                                                }
                                            }
                                            var v41 bool = (((!v3) && false) || false)
                                            _ = v37
                                            _ = v41
                                        }
                                        v3 = (1e10 < 0.1)
                                        var v42 uint16 = 0
                                        var v43 int32 = 1
                                        if false {
                                            if false {
                                                v12 += v12
                                                var v44 int64 = v26
                                                var v45 uint32 = genFn1((v12 | (v12 ^ uint64(v23))))
                                                _ = v44
                                                _ = v45
                                            } else {
                                                if v3 {
                                                    v46 := GenS5{F0: uint(v1), F1: 372, F2: int64(v43)}
                                                    if (4 <= 100) {
                                                        fmt.Println(fi24, v26, v3)
                                                        switch ("hello" + ((("" + "a") + "a") + (("hello" + ("" + ("hello" + ("foo bar" + ("foo bar" + "a"))))) + "foo bar"))) {
                                                        case "foo bar":
                                                            var v47 uint8 = (v1 / ((v1) | 1))
                                                            _ = v47
                                                        }
                                                        var v48 int8 = 127
                                                        switch -2147483648 {
                                                        case 128:
                                                            switch (fi25 / (((fi25 % ((545) | 1))) | 1)) {
                                                            case 32767:
                                                                fmt.Println(fi25, v42)
                                                                v49 := GenS6{F0: 650, F1: (fi25 + (fi25 | fi25)), F2: 2}
                                                                v1 -= v1
                                                                if (float32(1.5) < 0.1) {
                                                                    if true {
                                                                        v12 += (v12 & (v12 + uint64(v1)))
                                                                    }
                                                                    var v50 uint32 = uint32(v1)
                                                                    _ = v50
                                                                } else {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                }
                                                                fmt.Println(v26, v23)
                                                                v43 += (v43 % ((42) | 1))
                                                                if v3 {
                                                                    if (9223372036854775807 <= 4294967295) {
                                                                        switch fi24 {
                                                                        case 1:
                                                                            v23 -= (v23 << ((4294967295) & 63))
                                                                        default:
                                                                            var v51 string = ""
                                                                            _ = v51
                                                                        }
                                                                        v43 -= (v43 >> ((uint(v42)) & 63))
                                                                        var v52 uint64 = (v12 << ((0) & 63))
                                                                        _ = v52
                                                                    }
                                                                    var v53 uint16 = v42
                                                                    if (18446744073709551615 == 194) {
                                                                        var v54 uint = 4294967295
                                                                        if v3 {
                                                                            switch fi25 {
                                                                            case 32767:
                                                                                v43 = (v43 + v43)
                                                                            }
                                                                            if false {
                                                                                if ((fi24 - 32768) == (-(fi25))) {
                                                                                    fmt.Println(v43, v53)
                                                                                    switch (fi24 + 127) {
                                                                                    case 0:
                                                                                        if v3 {
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                if ((fi24 > (fi24 / ((fi24) | 1))) && (1.5 > 0.0)) {
                                                                                                    var v55 uint32 = uint32(v54)
                                                                                                    _ = len(v2)
                                                                                                    v55 += v23
                                                                                                    switch fi24 {
                                                                                                    case 32767:
                                                                                                        delete(v5, v55)
                                                                                                    case 9223372036854775807:
                                                                                                        fmt.Println(v42, fi25)
                                                                                                    }
                                                                                                    _ = v55
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        v56 := GenS2{F0: -9223372036854775808, F1: v53}
                                                                                        _ = v56
                                                                                    }
                                                                                    var v57 uint64 = 409
                                                                                    _ = len(v2)
                                                                                    switch (fi24 >> (((v54 << (((v54 / ((v54) | 1))) & 63))) & 63)) {
                                                                                    case -2147483648:
                                                                                        fmt.Println(v1, fi25)
                                                                                    case 1:
                                                                                        fmt.Println(v23, fi24)
                                                                                    }
                                                                                    v57 -= v57
                                                                                    switch "a" {
                                                                                    case "a":
                                                                                        if v3 {
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = 1
                                                                                            }
                                                                                            v58 := v5[uint32(v1)]
                                                                                            fmt.Println(v53)
                                                                                            _ = v58
                                                                                        }
                                                                                    case "":
                                                                                        v12 -= 381
                                                                                    default:
                                                                                        v4 -= 1e10
                                                                                        var v59 uint8 = uint8(v12)
                                                                                        fmt.Println(v53, fi24, v59)
                                                                                        _ = v59
                                                                                    }
                                                                                    v60 := GenS7{F0: (fi25 & fi24), F1: (fi25 | fi24)}
                                                                                    fmt.Println(v12)
                                                                                    _ = v57
                                                                                    _ = v60
                                                                                }
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = (fi24 << ((58) & 63))
                                                                                }
                                                                            } else {
                                                                                var v61 int64 = (^(v26))
                                                                                if v3 {
                                                                                    var v62 uint8 = v1
                                                                                    v4 = v4
                                                                                    v23 += (v23 * (v23 >> ((v54) & 63)))
                                                                                    _ = v62
                                                                                }
                                                                                v63 := GenS4{F0: (v48 ^ int8(fi24))}
                                                                                _ = v61
                                                                                _ = v63
                                                                            }
                                                                        }
                                                                        v64 := GenS2{F0: (fi25 * (-(fi25))), F1: (v42 << ((v54) & 63))}
                                                                        _ = v54
                                                                        _ = v64
                                                                    } else {
                                                                        v26 = (v26 / ((v26) | 1))
                                                                        v65 := GenS6{F0: 1, F1: (-(fi24)), F2: 77}
                                                                        _ = v65
                                                                    }
                                                                    switch 42 {
                                                                    case 32768:
                                                                        if false {
                                                                            var v66 int32 = (v43 ^ int32(fi25))
                                                                            delete(v5, (v23 - v23))
                                                                            if (1.5 <= (float32(1e10) - 1.5)) {
                                                                                if true {
                                                                                    var v67 int64 = (v26 << ((2) & 63))
                                                                                    switch "foo bar" {
                                                                                    case "\n\t":
                                                                                        var v68 uint64 = (v12 % ((v12) | 1))
                                                                                        _ = v68
                                                                                    case "":
                                                                                        fmt.Println(v3, v12)
                                                                                    case "foo bar":
                                                                                        v66 = v66
                                                                                    default:
                                                                                        var v69 uint32 = genFn1(42)
                                                                                        _ = v69
                                                                                    }
                                                                                    _ = v67
                                                                                }
                                                                                if true {
                                                                                    var v70 bool = true
                                                                                    v42 = uint16(v12)
                                                                                    v48 = v48
                                                                                    _ = v70
                                                                                }
                                                                                v26 = int64(v1)
                                                                                v66 = v66
                                                                            } else {
                                                                                if v3 {
                                                                                    var v71 float32 = 3.14159
                                                                                    v43 += (^(v43))
                                                                                    _ = v71
                                                                                } else {
                                                                                    var v72 bool = v3
                                                                                    if (!v3) {
                                                                                        if (v3 || false) {
                                                                                            v4 += 0.0
                                                                                            switch (^(fi25)) {
                                                                                            case 9223372036854775807:
                                                                                                v42 -= (^(v42))
                                                                                            }
                                                                                        } else {
                                                                                            v3 = v3
                                                                                        }
                                                                                        fmt.Println(fi24, v72)
                                                                                        if false {
                                                                                            if v72 {
                                                                                                switch 0 {
                                                                                                case 366:
                                                                                                    _ = len(v5)
                                                                                                default:
                                                                                                    var v73 float64 = v4
                                                                                                    _ = v73
                                                                                                }
                                                                                                fmt.Println(v53, v53)
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                            } else {
                                                                                                if v72 {
                                                                                                    var v74 uint64 = (v12 ^ uint64(v42))
                                                                                                    if (false || (0.0 > 0.0)) {
                                                                                                        v75 := GenS8{F0: "", F1: v43}
                                                                                                        var v76 uint16 = (v42 >> ((42) & 63))
                                                                                                        var v77 uint64 = v12
                                                                                                        var v78 string = ("hello" + ("hello" + ("hello" + ("a" + "\n\t"))))
                                                                                                        v79 := GenS7{F0: fi25, F1: fi24}
                                                                                                        _ = v75
                                                                                                        _ = v76
                                                                                                        _ = v77
                                                                                                        _ = v78
                                                                                                        _ = v79
                                                                                                    } else {
                                                                                                        var v80 bool = ((!false) && false)
                                                                                                        _ = v80
                                                                                                    }
                                                                                                    fmt.Println(v4, v12, v26)
                                                                                                    _ = v74
                                                                                                } else {
                                                                                                    _ = len(v2)
                                                                                                }
                                                                                                v81 := GenS9{F0: 105}
                                                                                                _ = v81
                                                                                            }
                                                                                            if ((v26 != 9223372036854775807) && (v3 || v3)) {
                                                                                                v1 = (^(v1))
                                                                                                v82 := GenS6{F0: 65535, F1: (fi25 / ((32768) | 1)), F2: (v12 >> ((2) & 63))}
                                                                                                _ = v82
                                                                                            }
                                                                                            if (18446744073709551615 >= 428) {
                                                                                                var v83 string = ("foo bar" + "")
                                                                                                switch -2147483648 {
                                                                                                case 533:
                                                                                                    fmt.Println(v1, v53, v3)
                                                                                                }
                                                                                                v66 -= int32(fi24)
                                                                                                var v84 int8 = (v48 - (v48 >> ((836) & 63)))
                                                                                                var v85 int64 = v26
                                                                                                v86 := v5[genFn1((v12 | (v12 >> ((18446744073709551615) & 63))))]
                                                                                                if false {
                                                                                                    v2 = append(v2, -9223372036854775808)
                                                                                                }
                                                                                                if (((!v72) || v72) || true) {
                                                                                                    var v87 int = (fi24 + (fi24 | (fi25 + (fi24 ^ 336))))
                                                                                                    v4 = 3.14159
                                                                                                    var v88 int16 = -71
                                                                                                    if len(v2) >= 2 {
                                                                                                        v2 = v2[1:len(v2)]
                                                                                                    }
                                                                                                    _ = v87
                                                                                                    _ = v88
                                                                                                }
                                                                                                v2 = append(v2, (fi24 >> ((819) & 63)))
                                                                                                switch fi24 {
                                                                                                case 127:
                                                                                                    if v3 {
                                                                                                        v89 := GenS6{F0: 4294967295, F1: -2147483648, F2: (v12 % ((0) | 1))}
                                                                                                        _ = v89
                                                                                                    }
                                                                                                case 32767:
                                                                                                    v90 := v5[(v23 & 2147483648)]
                                                                                                    _ = v90
                                                                                                case -9223372036854775808:
                                                                                                    switch (fi25 * fi24) {
                                                                                                    case 32768:
                                                                                                        fmt.Println(v43)
                                                                                                    case 128:
                                                                                                        fmt.Println(v83)
                                                                                                    case -2147483648:
                                                                                                        if v3 {
                                                                                                            var v91 int32 = int32(v84)
                                                                                                            var v92 int64 = v85
                                                                                                            _ = v91
                                                                                                            _ = v92
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                _ = len(v5)
                                                                                                if v3 {
                                                                                                    var v93 float64 = v4
                                                                                                    v94 := v5[(^(v23))]
                                                                                                    if (0 == 225) {
                                                                                                        fmt.Println(v85, v43, v26)
                                                                                                        var v95 int8 = (v48 * v48)
                                                                                                        var v96 string = v83
                                                                                                        _ = v95
                                                                                                        _ = v96
                                                                                                    }
                                                                                                    _ = v93
                                                                                                    _ = v94
                                                                                                }
                                                                                                _ = v83
                                                                                                _ = v84
                                                                                                _ = v85
                                                                                                _ = v86
                                                                                            } else {
                                                                                                v97 := v5[(v23 & (^(v23)))]
                                                                                                if true {
                                                                                                    if v3 {
                                                                                                        v98 := GenS10{F0: ((v97 + "\n\t") > v97), F1: 872}
                                                                                                        var v99 uint = 472
                                                                                                        v1 = v1
                                                                                                        _ = v98
                                                                                                        _ = v99
                                                                                                    } else {
                                                                                                        var v100 int16 = 32767
                                                                                                        v4 -= float64(0.1)
                                                                                                        v23 -= 65535
                                                                                                        _ = v100
                                                                                                    }
                                                                                                    var v101 int32 = v66
                                                                                                    switch (fi24 & fi25) {
                                                                                                    case -606:
                                                                                                        v102 := GenS2{F0: fi25, F1: 0}
                                                                                                        _ = v102
                                                                                                    default:
                                                                                                        v103 := GenS11{F0: v101, F1: 42, F2: (v66 + -2147483648)}
                                                                                                        _ = v103
                                                                                                    }
                                                                                                    _ = v101
                                                                                                } else {
                                                                                                    var v104 int = fi24
                                                                                                    switch (v97 + "a") {
                                                                                                    case "hello":
                                                                                                        switch fi24 {
                                                                                                        case -535:
                                                                                                            v72 = v3
                                                                                                        case -2147483648:
                                                                                                            v105 := GenS6{F0: 112, F1: (-(fi25)), F2: 4294967295}
                                                                                                            _ = v105
                                                                                                        case 32767:
                                                                                                            if v72 {
                                                                                                                var v106 int8 = (^(v48))
                                                                                                                var v107 uint8 = uint8(v12)
                                                                                                                _ = len(v5)
                                                                                                                _ = v106
                                                                                                                _ = v107
                                                                                                            } else {
                                                                                                                if (4294967295 >= uint(v23)) {
                                                                                                                    if ((v66 | v43) <= (v66 >> ((1) & 63))) {
                                                                                                                        switch v97 {
                                                                                                                        case "foo bar":
                                                                                                                            switch "foo bar" {
                                                                                                                            case "":
                                                                                                                                fmt.Println(fi25, v43, v4)
                                                                                                                            case "a":
                                                                                                                                v48 = int8(v1)
                                                                                                                            case "hello":
                                                                                                                                v48 += v48
                                                                                                                            }
                                                                                                                        case "\n\t":
                                                                                                                            v72 = ((!true) && (-1 != genFn0(0, v3, (v1 + v1))))
                                                                                                                        default:
                                                                                                                            var v108 uint = 1
                                                                                                                            _ = v108
                                                                                                                        }
                                                                                                                        var v109 float32 = 0.1
                                                                                                                        _ = v109
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v110 := GenS6{F0: 869, F1: 32768, F2: v12}
                                                                                                                    v72 = (!v3)
                                                                                                                    switch ((("hello" + (v97 + ((v97 + v97) + "\n\t"))) + "") + "a") {
                                                                                                                    case "foo bar":
                                                                                                                        v48 -= (v48 >> ((uint(v12)) & 63))
                                                                                                                    default:
                                                                                                                        v111 := GenS12{F0: (v12 | (v12 | 1)), F1: 776, F2: (v72 || (26 <= v1))}
                                                                                                                        _ = v111
                                                                                                                    }
                                                                                                                    _ = v110
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            v43 -= 621
                                                                                                        }
                                                                                                    default:
                                                                                                        v5[genFn1((^(v12)))] = ("foo bar" + v97)
                                                                                                    }
                                                                                                    v48 = (v48 / ((v48) | 1))
                                                                                                    _ = v104
                                                                                                }
                                                                                                if v72 {
                                                                                                    fmt.Println(v72, v3, v23)
                                                                                                }
                                                                                                _ = v97
                                                                                            }
                                                                                        } else {
                                                                                            var v112 string = ("foo bar" + "a")
                                                                                            v23 += genFn1(v12)
                                                                                            _ = v112
                                                                                        }
                                                                                    }
                                                                                    _ = v72
                                                                                }
                                                                            }
                                                                            _ = v66
                                                                        }
                                                                    }
                                                                    if v3 {
                                                                        var v113 int32 = 2147483647
                                                                        v48 -= 127
                                                                        switch int(v23) {
                                                                        case -256:
                                                                            if (true && v3) {
                                                                                if ((((0 <= -2147483648) || v3) && (275 > 0)) || (!(genFn1((^(v12))) != uint32(v12)))) {
                                                                                    v23 += v23
                                                                                }
                                                                            }
                                                                        default:
                                                                            switch (fi24 << ((4294967295) & 63)) {
                                                                            case -2147483648:
                                                                                fmt.Println(v23, v4, v26)
                                                                            default:
                                                                                var v114 int32 = v43
                                                                                _ = v114
                                                                            }
                                                                        }
                                                                        _ = v113
                                                                    }
                                                                    _ = v53
                                                                }
                                                                v26 = -1
                                                                v115 := v5[v23]
                                                                _ = v115
                                                                _ = v49
                                                            default:
                                                                if ((2 > 302) || (-2.25 >= float32(v4))) {
                                                                    if (v3 || (!true)) {
                                                                        var v116 uint8 = (v1 | v1)
                                                                        var v117 uint64 = (^(v12))
                                                                        _ = v116
                                                                        _ = v117
                                                                    } else {
                                                                        delete(v5, genFn1((^(v12))))
                                                                    }
                                                                    var v118 int = -172
                                                                    if ((v3 || (v3 || ((v12 / ((v12) | 1)) == (v12 * (v12 & (^(v12))))))) && v3) {
                                                                        v119 := GenS13{F0: genFn0(0, (v3 || (!(true && v3))), v1), F1: (v23 - uint32(v42)), F2: 2}
                                                                        switch (fi24 + (fi25 + -2147483648)) {
                                                                        case 2147483647:
                                                                            v23 = uint32(v1)
                                                                        case 0:
                                                                            v118 = fi25
                                                                        case 144:
                                                                            var v120 int16 = 788
                                                                            _ = v120
                                                                        default:
                                                                            _ = len(v5)
                                                                        }
                                                                        _ = len(v2)
                                                                        _ = v119
                                                                    } else {
                                                                        v118 -= v118
                                                                        var v121 bool = ((true || (v48 == -12)) || true)
                                                                        _ = v121
                                                                    }
                                                                    switch (^(v118)) {
                                                                    case 432:
                                                                        switch "" {
                                                                        case "foo bar":
                                                                            fmt.Println(v48)
                                                                        case "\n\t":
                                                                            var v122 int16 = -32768
                                                                            _ = v122
                                                                        case "":
                                                                            var v123 uint16 = 0
                                                                            _ = v123
                                                                        }
                                                                    case 271:
                                                                        if (v3 && ((!((^(v43)) != -947)) && (42 <= 255))) {
                                                                            var v124 int = (fi24 - (fi25 - (^(fi24))))
                                                                            _ = v124
                                                                        }
                                                                    }
                                                                    var v125 int = v118
                                                                    var v126 string = "a"
                                                                    switch fi24 {
                                                                    case 4294967295:
                                                                        if (((!v3) || (!false)) || (!true)) {
                                                                            var v127 int32 = (v43 - (-(v43)))
                                                                            _ = v127
                                                                        }
                                                                    case 32768:
                                                                        var v128 bool = (v48 == v48)
                                                                        _ = v128
                                                                    case 2147483647:
                                                                        var v129 uint64 = (v12 + (v12 * v12))
                                                                        _ = v129
                                                                    }
                                                                    v1 += (v1 | v1)
                                                                    if v3 {
                                                                        v5[(^(v23))] = "hello"
                                                                    } else {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        fmt.Println(v23)
                                                                    }
                                                                    v12 += 2
                                                                    v43 += (v43 & 327)
                                                                    v130 := v5[(v23 * 2)]
                                                                    _ = v118
                                                                    _ = v125
                                                                    _ = v126
                                                                    _ = v130
                                                                } else {
                                                                    var v131 bool = (false || (((v3 || v3) || false) && ((v48 & v48) > (v48 >> ((0) & 63)))))
                                                                    v1 += 255
                                                                    _ = v131
                                                                }
                                                            }
                                                        }
                                                        v43 += v43
                                                        if v3 {
                                                            switch "foo bar" {
                                                            case "":
                                                                if (v43 <= (v43 * v43)) {
                                                                    var v132 uint16 = v42
                                                                    _ = v132
                                                                } else {
                                                                    v133 := GenS14{F0: (v1 < v1), F1: int32(v26), F2: 2}
                                                                    _ = v133
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v23)
                                                                v42 = v42
                                                                fmt.Println(v42)
                                                            }
                                                            v134 := GenS6{F0: 0, F1: fi25, F2: v12}
                                                            var v135 string = "hello"
                                                            switch fi25 {
                                                            case -260:
                                                                switch (v135 + "\n\t") {
                                                                case "hello":
                                                                    v136 := GenS3{F0: v135, F1: (v42 | (v42 / ((v42) | 1)))}
                                                                    _ = v136
                                                                default:
                                                                    v43 -= (v43 << ((296) & 63))
                                                                }
                                                            }
                                                            _ = v134
                                                            _ = v135
                                                        } else {
                                                            v1 += (v1 / ((v1) | 1))
                                                            fmt.Println(v3, v12)
                                                            switch -9223372036854775808 {
                                                            case 824:
                                                                var v137 string = ((("" + ("hello" + "")) + ("hello" + "a")) + ("\n\t" + "hello"))
                                                                _ = v137
                                                            case 4294967295:
                                                                fmt.Println(v42, v43)
                                                            case 128:
                                                                fmt.Println(v1, v1)
                                                            }
                                                        }
                                                        _ = v48
                                                    } else {
                                                        delete(v5, (v23 & v23))
                                                        var v138 float32 = float32(v4)
                                                        _ = v138
                                                    }
                                                    _ = v46
                                                } else {
                                                    v12 -= v12
                                                    var v139 uint16 = v42
                                                    _ = v139
                                                }
                                                if (v4 != v4) {
                                                    fmt.Println(v4, fi24)
                                                    if (false || (v1 < 42)) {
                                                        if (!(1e10 > 1e10)) {
                                                            fmt.Println(v23)
                                                            var v140 uint = 1
                                                            _ = v140
                                                        } else {
                                                            switch int(v23) {
                                                            case -9223372036854775808:
                                                                var v141 int32 = v43
                                                                switch fi25 {
                                                                case 32768:
                                                                    if v3 {
                                                                        v142 := GenS14{F0: v3, F1: (v43 >> ((485) & 63)), F2: 18}
                                                                        fmt.Println(v3, fi25)
                                                                        _ = v142
                                                                    } else {
                                                                        v143 := GenS15{F0: 0, F1: v3}
                                                                        if v3 {
                                                                            var v144 int = fi25
                                                                            var v145 bool = v3
                                                                            _ = v144
                                                                            _ = v145
                                                                        }
                                                                        _ = v143
                                                                    }
                                                                case 4294967295:
                                                                    if (v141 <= (v141 * v43)) {
                                                                        v1 = 241
                                                                        var v146 uint32 = uint32(v12)
                                                                        _ = v146
                                                                    }
                                                                    if ((v26 ^ v26) >= 9223372036854775807) {
                                                                        fmt.Println(v4)
                                                                        switch 42 {
                                                                        case 0:
                                                                            var v147 int32 = (-(v141))
                                                                            _ = v147
                                                                        }
                                                                    } else {
                                                                        var v148 int = 893
                                                                        var v149 int8 = -76
                                                                        _ = v148
                                                                        _ = v149
                                                                    }
                                                                case 2147483647:
                                                                    v150 := GenS11{F0: v141, F1: -32768, F2: (v141 % (((v43 | (v43 - v43))) | 1))}
                                                                    _ = v150
                                                                default:
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        v141 -= (v43 + -2147483648)
                                                                    case "foo bar":
                                                                        if v3 {
                                                                            switch fi24 {
                                                                            case -958:
                                                                                var v151 bool = (42 <= 11)
                                                                                _ = v151
                                                                            case -536:
                                                                                switch (-(fi24)) {
                                                                                case -2147483648:
                                                                                    fmt.Println(v4)
                                                                                case 432:
                                                                                    if ((fi25 - 2147483647) < (fi24 << ((uint(v42)) & 63))) {
                                                                                        v4 += v4
                                                                                        var v152 int16 = -808
                                                                                        switch "" {
                                                                                        case "hello":
                                                                                            v26 += (v26 ^ v26)
                                                                                        default:
                                                                                            v4 = v4
                                                                                            var v153 string = (("" + (("\n\t" + ("foo bar" + "foo bar")) + ("foo bar" + "hello"))) + (("hello" + "") + ""))
                                                                                            if true {
                                                                                                v154 := GenS16{F0: (^(v1)), F1: (v23 << ((255) & 63)), F2: (v152 - v152)}
                                                                                                var v155 uint = 65535
                                                                                                _ = v154
                                                                                                _ = v155
                                                                                            } else {
                                                                                                switch v153 {
                                                                                                case "":
                                                                                                    var v156 int16 = (v152 + int16(fi25))
                                                                                                    _ = v156
                                                                                                default:
                                                                                                    if v3 {
                                                                                                        fmt.Println(v43, v4)
                                                                                                        switch "foo bar" {
                                                                                                        case "a":
                                                                                                            v157 := GenS17{F0: v1, F1: 909}
                                                                                                            _ = v157
                                                                                                        case "":
                                                                                                            var v158 float64 = -2.25
                                                                                                            _ = v158
                                                                                                        case "foo bar":
                                                                                                            var v159 uint = 0
                                                                                                            _ = v159
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v160 float64 = 3.14159
                                                                                                        var v161 uint16 = v42
                                                                                                        _ = v160
                                                                                                        _ = v161
                                                                                                    }
                                                                                                }
                                                                                                var v162 int32 = (v43 & (v43 << ((855) & 63)))
                                                                                                if v3 {
                                                                                                    v23 += 42
                                                                                                    fmt.Println(v152, v23, v42)
                                                                                                } else {
                                                                                                    if true {
                                                                                                        v141 -= 202
                                                                                                        var v163 uint32 = 2
                                                                                                        switch int(v42) {
                                                                                                        case 4294967295:
                                                                                                            v163 += uint32(v1)
                                                                                                        case 1:
                                                                                                            var v164 string = v153
                                                                                                            v12 += v12
                                                                                                            _ = v164
                                                                                                        }
                                                                                                        if v3 {
                                                                                                            var v165 int32 = (v162 - int32(v26))
                                                                                                            if (-1 > -128) {
                                                                                                                var v166 uint64 = v12
                                                                                                                if true {
                                                                                                                    _ = len(v5)
                                                                                                                } else {
                                                                                                                    v4 = (float64(float32((v4 - (v4 + (0.0 - 3.14159))))) - (-2.25 * -2.25))
                                                                                                                    v26 = v26
                                                                                                                    var v167 string = v153
                                                                                                                    _ = v167
                                                                                                                }
                                                                                                                fmt.Println(v141, fi24, v152)
                                                                                                                _ = v166
                                                                                                            }
                                                                                                            _ = v165
                                                                                                        } else {
                                                                                                            var v168 bool = ((^(v43)) == v43)
                                                                                                            _ = v168
                                                                                                        }
                                                                                                        _ = v163
                                                                                                    } else {
                                                                                                        v169 := GenS18{F0: (!((v4 - ((v4 * 1.5) / 1.0)) == v4)), F1: (v3 || ((-(v26)) > 0))}
                                                                                                        var v170 int = (fi24 * fi24)
                                                                                                        _ = v169
                                                                                                        _ = v170
                                                                                                    }
                                                                                                    v141 = v141
                                                                                                }
                                                                                                _ = v162
                                                                                            }
                                                                                            _ = v153
                                                                                        }
                                                                                        _ = v152
                                                                                    } else {
                                                                                        var v171 uint64 = 330
                                                                                        var v172 string = "hello"
                                                                                        _ = v171
                                                                                        _ = v172
                                                                                    }
                                                                                case 2147483647:
                                                                                    v3 = v3
                                                                                default:
                                                                                    v4 -= (v4 + 3.14159)
                                                                                }
                                                                                v1 = 0
                                                                            case -1:
                                                                                fmt.Println(v23, v1)
                                                                            default:
                                                                                if v3 {
                                                                                    fmt.Println(v43, v42)
                                                                                    v173 := GenS11{F0: 2147483647, F1: 42, F2: 1}
                                                                                    var v174 float64 = ((v4 + (v4 * v4)) / ((0.0 / 3.14159) - v4))
                                                                                    switch "hello" {
                                                                                    case "\n\t":
                                                                                        var v175 int = 4294967295
                                                                                        _ = v175
                                                                                    case "foo bar":
                                                                                        v141 = (v141 << ((42) & 63))
                                                                                    case "hello":
                                                                                        v176 := GenS17{F0: 0, F1: (v26 << ((4294967295) & 63))}
                                                                                        _ = v176
                                                                                    default:
                                                                                        v23 = (v23 * 2147483648)
                                                                                    }
                                                                                    var v177 int = fi25
                                                                                    _ = v173
                                                                                    _ = v174
                                                                                    _ = v177
                                                                                } else {
                                                                                    v1 += 255
                                                                                    if false {
                                                                                        fmt.Println(v141, v43)
                                                                                        fmt.Println(fi24)
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = fi25
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            v178 := GenS17{F0: v1, F1: (-(v26))}
                                                                            _ = v178
                                                                        }
                                                                    default:
                                                                        fmt.Println(v3, v23, v43)
                                                                    }
                                                                }
                                                                _ = v141
                                                            }
                                                            fmt.Println(v3)
                                                            var v179 uint16 = v42
                                                            _ = v179
                                                        }
                                                        v180 := GenS11{F0: (^(v43)), F1: 825, F2: (v43 << ((633) & 63))}
                                                        var v181 float64 = float64(-2.25)
                                                        _ = v180
                                                        _ = v181
                                                    } else {
                                                        var v182 int16 = 46
                                                        var v183 uint8 = v1
                                                        v1 = 127
                                                        switch int(v182) {
                                                        case 600:
                                                            v43 -= 2147483647
                                                        case 146:
                                                            v3 = (v43 != 1)
                                                        default:
                                                            v4 = 3.14159
                                                        }
                                                        _ = v182
                                                        _ = v183
                                                    }
                                                    if v3 {
                                                        fmt.Println(v43)
                                                        var v184 int32 = (v43 | (v43 | v43))
                                                        _ = v184
                                                    }
                                                }
                                                switch 128 {
                                                case -181:
                                                    var v185 float32 = -2.25
                                                    _ = v185
                                                case 4294967295:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case 32768:
                                                    var v186 int8 = 100
                                                    _ = v186
                                                }
                                            }
                                            var v187 int32 = -1
                                            _ = v187
                                        }
                                        _ = v42
                                        _ = v43
                                    }
                                    v188 := GenS16{F0: (v1 & (v1 / ((v1) | 1))), F1: genFn1(v12), F2: -1}
                                    _ = v188
                                    _ = v26
                                }
                            }
                            v3 = ((v12 | v12) > 473)
                            _ = fi25
                        }
                        for fi189 := 0; fi189 < 7; fi189++ {
                            v190 := v5[uint32(v1)]
                            _ = fi189
                            _ = v190
                        }
                        if v3 {
                            v191 := GenS11{F0: 0, F1: -1, F2: 0}
                            var v192 int8 = -62
                            v193 := GenS18{F0: ((18446744073709551615 < 0) && (-1 <= 32767)), F1: true}
                            _ = v191
                            _ = v192
                            _ = v193
                        }
                        _ = fi24
                    }
                } else {
                    var v194 uint32 = genFn1((v12 - (v12 / ((v12) | 1))))
                    if len(v2) > 0 {
                        v2[0] = 32767
                    }
                    var v195 uint32 = (v23 ^ 2147483648)
                    _ = v194
                    _ = v195
                }
                var v196 float64 = 0.1
                _ = v196
                _ = v23
            }
        default:
            v197 := GenS4{F0: -1}
            _ = v197
        }
        if (32768 >= 32768) {
            var v198 int16 = 42
            _ = v198
        } else {
            _ = len(v5)
            for fi199 := 0; fi199 < 11; fi199++ {
                var v200 float64 = ((v4 - v4) / 1.0)
                _ = fi199
                _ = v200
            }
            if v3 {
                var v201 int32 = 42
                var v202 bool = (4294967295 != 206)
                if (793 < 42) {
                    v203 := GenS6{F0: 857, F1: 376, F2: v12}
                    var v204 int = -1
                    if v3 {
                        fmt.Println(v12, v204, v202)
                        var v205 int = 311
                        _ = v205
                    }
                    _ = v203
                    _ = v204
                }
                _ = v201
                _ = v202
            } else {
                v3 = (!(!(genFn1(v12) >= 2147483648)))
                v2 = append(v2, 42)
            }
        }
        fmt.Println(v1)
    } else {
        _ = len(v5)
        _ = len(v2)
    }
    v206 := v5[454]
    v207 := GenS11{F0: 932, F1: -1, F2: -1}
    var v208 int32 = 246
    v209 := GenS14{F0: false, F1: (v208 | (v208 << ((42) & 63))), F2: 18446744073709551615}
    var v210 int16 = -1
    v211 := GenS10{F0: ((v210 + v210) >= (v210 ^ 192)), F1: 18446744073709551615}
    for fi212 := 0; fi212 < 20; fi212++ {
        var v213 uint32 = 65535
        _ = fi212
        _ = v213
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v206:", v206)
    fmt.Println("v208:", v208)
    fmt.Println("v210:", v210)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v206
    _ = v207
    _ = v208
    _ = v209
    _ = v210
    _ = v211
    _ = v3
    _ = v4
    _ = v5
}
