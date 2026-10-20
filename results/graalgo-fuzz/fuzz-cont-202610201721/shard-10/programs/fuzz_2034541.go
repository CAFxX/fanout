package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int8; F1 int64; F2 int8 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 bool; F1 int16; F2 int32 }
type GenS5 struct { F0 uint; F1 int16; F2 string }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int64; F1 int8; F2 bool }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint8; F1 int32 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int8; F1 int64 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint8; F1 uint64; F2 uint32 }
func genFn0(p0 string, p1 int) uint64 {
    fmt.Println(p0, p0)
    var v7 bool = false
    v8 := []int{p1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 783
}
func genFn1() uint32 {
    switch 128 {
    case 9223372036854775807:
        var v9 uint8 = 128
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint32 = 4294967295
        v12 := []int{fi10, (fi10 + 32767), fi10}
        switch (("\n\t" + (("\n\t" + "foo bar") + ("" + ""))) + "a") {
        case "hello":
            var v13 uint = 65535
            _ = v13
        case "\n\t":
            v14 := GenS1{F0: -1, F1: (fi10 << ((1) & 63))}
            _ = v14
        case "":
            v11 = (v11 % ((v11) | 1))
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    v15 := []int{-804}
    if (false || false) {
        if true {
            var v16 int16 = 1
            var v17 int16 = 1000
            _ = v16
            _ = v17
        }
        if false {
            if false {
                for fi18 := 0; fi18 < 1; fi18++ {
                    if true {
                        for fi19 := 0; fi19 < 1; fi19++ {
                            var v20 uint = 18446744073709551615
                            if (1 < 0) {
                                switch ("foo bar" + "") {
                                case "foo bar":
                                    if true {
                                        _ = len(v15)
                                        v21 := GenS2{F0: -128, F1: -1, F2: 100}
                                        _ = v21
                                    }
                                case "hello":
                                    v15 = append(v15, 397)
                                case "a":
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                default:
                                    v20 = v20
                                }
                                switch (fi19 - (-(fi18))) {
                                case 42:
                                    v20 = (^(v20))
                                case -9223372036854775808:
                                    v22 := GenS2{F0: 100, F1: 992, F2: 0}
                                    switch "hello" {
                                    case "a":
                                        v20 = (v20 * 42)
                                        if (-1 <= 1) {
                                            if (32767 == 32767) {
                                                if ((42 > v20) || ((!false) || (true || ((!((fi19 != (^(fi18))) || true)) && (855 > 42))))) {
                                                    fmt.Println(fi18, fi19)
                                                }
                                                switch "" {
                                                case "":
                                                    if true {
                                                        switch ("" + "\n\t") {
                                                        case "foo bar":
                                                            fmt.Println(v20, v20)
                                                        default:
                                                            if len(v15) >= 2 {
                                                                v15 = v15[1:len(v15)]
                                                            }
                                                        }
                                                        if true {
                                                            if (float64((-2.25 / 3.14159)) > 1.5) {
                                                                var v23 uint16 = 2
                                                                v23 = (v23 - v23)
                                                                _ = v23
                                                            } else {
                                                                v24 := GenS2{F0: 0, F1: 744, F2: 1}
                                                                if (0.0 > 3.14159) {
                                                                    v25 := GenS3{F0: (v20 * 4294967295), F1: (-2.25 < (1e10 * float32(float64(1.5))))}
                                                                    _ = v25
                                                                } else {
                                                                    var v26 int32 = 2147483647
                                                                    _ = v26
                                                                }
                                                                var v27 uint8 = 1
                                                                if true {
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        var v28 uint = (v20 / ((v20) | 1))
                                                                        _ = v28
                                                                    case "foo bar":
                                                                        switch "a" {
                                                                        case "hello":
                                                                            if (65535 <= 832) {
                                                                                switch fi19 {
                                                                                case 128:
                                                                                    v20 -= v20
                                                                                }
                                                                            } else {
                                                                                v20 = 2
                                                                            }
                                                                        default:
                                                                            v15 = append(v15, fi19)
                                                                            v27 -= v27
                                                                            if (!false) {
                                                                                v20 = (v20 ^ uint(v27))
                                                                                v29 := GenS4{F0: true, F1: -1, F2: 2147483647}
                                                                                v27 += (v27 & (^(v27)))
                                                                                _ = v29
                                                                            } else {
                                                                                v27 = v27
                                                                                if false {
                                                                                    v30 := GenS2{F0: 42, F1: 42, F2: -86}
                                                                                    if (42 >= 0) {
                                                                                        if true {
                                                                                            if (-55 >= -2147483648) {
                                                                                                fmt.Println(v20, v27, fi18)
                                                                                                if len(v15) >= 2 {
                                                                                                    v15 = v15[1:len(v15)]
                                                                                                }
                                                                                                fmt.Println(fi18, v20)
                                                                                            }
                                                                                        }
                                                                                        _ = len(v15)
                                                                                        v31 := GenS5{F0: v20, F1: 32767, F2: ("" + ("a" + "\n\t"))}
                                                                                        _ = v31
                                                                                    } else {
                                                                                        v32 := GenS6{F0: (fi18 / ((fi19) | 1)), F1: 389}
                                                                                        var v33 uint8 = v27
                                                                                        v33 -= v27
                                                                                        var v34 bool = false
                                                                                        _ = v32
                                                                                        _ = v33
                                                                                        _ = v34
                                                                                    }
                                                                                    _ = v30
                                                                                }
                                                                                v35 := GenS6{F0: (fi18 % ((fi18) | 1)), F1: 0}
                                                                                _ = v35
                                                                            }
                                                                        }
                                                                    default:
                                                                        if (3.14159 < 3.14159) {
                                                                            if true {
                                                                                v36 := GenS2{F0: 127, F1: 1, F2: 0}
                                                                                var v37 int16 = 876
                                                                                v20 = (^(v20))
                                                                                fmt.Println(v37, v37)
                                                                                var v38 string = "a"
                                                                                _ = v36
                                                                                _ = v37
                                                                                _ = v38
                                                                            } else {
                                                                                v27 = 93
                                                                            }
                                                                            v27 = (v27 ^ 0)
                                                                            var v39 float64 = 0.0
                                                                            _ = v39
                                                                        }
                                                                    }
                                                                    var v40 int32 = -230
                                                                    if (-9223372036854775808 >= -9223372036854775808) {
                                                                        v15 = append(v15, (fi19 >> (((^(v20))) & 63)))
                                                                        switch (("hello" + "") + "hello") {
                                                                        case "\n\t":
                                                                            _ = len(v15)
                                                                        case "foo bar":
                                                                            v27 = (v27 / (((^(v27))) | 1))
                                                                        case "a":
                                                                            v41 := GenS7{F0: (^(v20))}
                                                                            _ = v41
                                                                        default:
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                if true {
                                                                                    v42 := GenS8{F0: 494}
                                                                                    _ = v42
                                                                                }
                                                                            case "hello":
                                                                                var v43 int = fi18
                                                                                _ = v43
                                                                            case "foo bar":
                                                                                fmt.Println(fi18)
                                                                            default:
                                                                                v40 = (v40 >> ((v20) & 63))
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v40
                                                                } else {
                                                                    v20 += v20
                                                                    if true {
                                                                        fmt.Println(fi19, fi18)
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            v44 := GenS5{F0: v20, F1: 1, F2: "foo bar"}
                                                                            _ = v44
                                                                        default:
                                                                            v27 = 42
                                                                        }
                                                                    } else {
                                                                        v27 += uint8(v20)
                                                                        var v45 uint64 = 307
                                                                        _ = v45
                                                                    }
                                                                }
                                                                _ = v24
                                                                _ = v27
                                                            }
                                                        }
                                                        if false {
                                                            var v46 uint64 = 2
                                                            var v47 bool = true
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                        var v48 int16 = 1000
                                                        if (2 < 413) {
                                                            var v49 uint64 = genFn0("", fi18)
                                                            if false {
                                                                switch "a" {
                                                                case "hello":
                                                                    var v50 uint64 = 42
                                                                    _ = v50
                                                                case "":
                                                                    _ = len(v15)
                                                                }
                                                                v48 = v48
                                                                if len(v15) > 0 {
                                                                    v15[0] = 1
                                                                }
                                                            }
                                                            if (2147483647 >= int32(fi18)) {
                                                                var v51 float64 = 0.1
                                                                if len(v15) >= 2 {
                                                                    v15 = v15[1:len(v15)]
                                                                }
                                                                if true {
                                                                    if true {
                                                                        fmt.Println(fi18, fi19)
                                                                        v52 := GenS4{F0: ((0.1 / (((-2.25 + 3.14159) / 1.0) * (1.5 + (1e10 - (-2.25 - ((3.14159 * 1e10) * float32(v51))))))) != ((0.0 / float32(1.5)) + (float32((((-2.25 - v51) / 3.14159) / v51)) + 1e10))), F1: v48, F2: -2147483648}
                                                                        _ = v52
                                                                    }
                                                                    v20 += 2
                                                                }
                                                                var v53 uint8 = 42
                                                                _ = v51
                                                                _ = v53
                                                            } else {
                                                                if ((("a" + "hello") + "\n\t") < "hello") {
                                                                    var v54 int32 = 484
                                                                    fmt.Println(fi19, v20, v49)
                                                                    _ = v54
                                                                } else {
                                                                    v49 = v49
                                                                    var v55 int64 = -1
                                                                    _ = v55
                                                                }
                                                                if true {
                                                                    if (false || (((((v49 / (((v49 << ((0) & 63))) | 1)) <= genFn0("foo bar", fi19)) && (42 != 2)) && (!false)) || (1.5 == float32(0.0)))) {
                                                                        v48 = v48
                                                                        var v56 uint16 = 684
                                                                        _ = v56
                                                                    } else {
                                                                        v57 := GenS6{F0: -9223372036854775808, F1: 1}
                                                                        switch "a" {
                                                                        case "a":
                                                                            var v58 int64 = 9223372036854775807
                                                                            _ = v58
                                                                        default:
                                                                            v48 = v48
                                                                        }
                                                                        _ = v57
                                                                    }
                                                                    var v59 uint8 = 1
                                                                    var v60 float64 = 1e10
                                                                    _ = v59
                                                                    _ = v60
                                                                } else {
                                                                    v61 := GenS9{F0: 127}
                                                                    v62 := GenS2{F0: -128, F1: 0, F2: -128}
                                                                    if false {
                                                                        fmt.Println(fi18, fi18, fi18)
                                                                        v63 := GenS10{F0: 0, F1: 121, F2: true}
                                                                        _ = v63
                                                                    } else {
                                                                        var v64 bool = (1 > 42)
                                                                        _ = v64
                                                                    }
                                                                    _ = v61
                                                                    _ = v62
                                                                }
                                                                if len(v15) >= 2 {
                                                                    v15 = v15[1:len(v15)]
                                                                }
                                                            }
                                                            _ = v49
                                                        } else {
                                                            fmt.Println(v48)
                                                            v20 = v20
                                                        }
                                                        _ = v48
                                                    } else {
                                                        v20 = v20
                                                    }
                                                case "a":
                                                    v20 = v20
                                                }
                                                var v65 int64 = -9223372036854775808
                                                _ = v65
                                            } else {
                                                var v66 uint8 = 253
                                                _ = v66
                                            }
                                            var v67 uint32 = 2147483647
                                            v20 = v20
                                            _ = v67
                                        } else {
                                            if len(v15) >= 2 {
                                                v15 = v15[1:len(v15)]
                                            }
                                        }
                                    case "\n\t":
                                        v20 += v20
                                    }
                                    _ = v22
                                }
                            }
                            _ = fi19
                            _ = v20
                        }
                        for fi68 := 0; fi68 < 19; fi68++ {
                            var v69 uint64 = 18446744073709551615
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            switch int(v69) {
                            case -865:
                                if (29 > 65535) {
                                    var v70 uint16 = 42
                                    v70 = v70
                                    _ = v70
                                }
                            case 128:
                                switch (("foo bar" + "hello") + "a") {
                                case "hello":
                                    v69 = (^(v69))
                                case "\n\t":
                                    if len(v15) > 0 {
                                        v15[0] = fi68
                                    }
                                    if (0.0 <= (-2.25 / (((-2.25 / 0.1) / (0.1 / ((1.5 * 1e10) / 1e10))) / 0.1))) {
                                        if false {
                                            var v71 uint = 1
                                            if false {
                                                fmt.Println(fi18, fi18)
                                                v72 := GenS7{F0: 255}
                                                _ = v72
                                            } else {
                                                var v73 string = ("foo bar" + (("a" + "") + "a"))
                                                fmt.Println(v71)
                                                if (!true) {
                                                    switch (fi18 << ((uint(v69)) & 63)) {
                                                    case -560:
                                                        v69 -= v69
                                                    case 9223372036854775807:
                                                        var v74 string = "a"
                                                        _ = v74
                                                    case -2147483648:
                                                        v69 = genFn0("hello", 9223372036854775807)
                                                    default:
                                                        v71 = (v71 >> ((v71) & 63))
                                                    }
                                                } else {
                                                    v73 = (("" + v73) + ((v73 + ("a" + ((v73 + v73) + "foo bar"))) + v73))
                                                    if false {
                                                        v71 += v71
                                                        if (3.14159 == 0.0) {
                                                            if len(v15) >= 2 {
                                                                v15 = v15[1:len(v15)]
                                                            }
                                                            if false {
                                                                v71 += 4294967295
                                                                v71 = v71
                                                                _ = len(v15)
                                                            }
                                                            if true {
                                                                switch -9223372036854775808 {
                                                                case -9223372036854775808:
                                                                    var v75 float32 = 1.5
                                                                    _ = v75
                                                                case 32768:
                                                                    if (665 <= (v69 % (((v69 + genFn0("a", (fi68 ^ (fi68 * fi18))))) | 1))) {
                                                                        v69 -= genFn0("foo bar", (fi68 ^ fi18))
                                                                    }
                                                                default:
                                                                    if false {
                                                                        v69 += (v69 ^ genFn0(v73, fi18))
                                                                        switch ((v73 + "foo bar") + ("hello" + "hello")) {
                                                                        case "foo bar":
                                                                            if (!true) {
                                                                                if ("" <= (v73 + v73)) {
                                                                                    fmt.Println(fi18)
                                                                                    if (0.1 >= 3.14159) {
                                                                                        v71 += 4294967295
                                                                                        var v76 int8 = -128
                                                                                        if ((((!(32767 != 0)) || true) || ((v76 % (((v76 ^ (-(v76)))) | 1)) == (v76 ^ 42))) && false) {
                                                                                            switch v73 {
                                                                                            case "\n\t":
                                                                                                fmt.Println(v69, v76)
                                                                                            case "hello":
                                                                                                var v77 bool = false
                                                                                                _ = v77
                                                                                            default:
                                                                                                v78 := GenS4{F0: true, F1: 0, F2: 42}
                                                                                                _ = v78
                                                                                            }
                                                                                            v69 -= (v69 / (((v69 * v69)) | 1))
                                                                                            var v79 int = (fi18 >> ((1) & 63))
                                                                                            _ = v79
                                                                                        } else {
                                                                                            fmt.Println(v76, fi18)
                                                                                            v76 += (v76 * -1)
                                                                                        }
                                                                                        _ = v76
                                                                                    }
                                                                                } else {
                                                                                    switch v73 {
                                                                                    case "":
                                                                                        var v80 bool = ((^(v69)) <= v69)
                                                                                        _ = v80
                                                                                    case "\n\t":
                                                                                        if true {
                                                                                            switch (fi68 / ((fi68) | 1)) {
                                                                                            case 1:
                                                                                                v81 := GenS6{F0: 1, F1: 0}
                                                                                                _ = v81
                                                                                            }
                                                                                            var v82 uint32 = uint32(v69)
                                                                                            _ = v82
                                                                                        } else {
                                                                                            if false {
                                                                                                v69 = 686
                                                                                            } else {
                                                                                                v83 := GenS11{F0: 189}
                                                                                                _ = v83
                                                                                            }
                                                                                        }
                                                                                        v84 := GenS8{F0: 303}
                                                                                        _ = v84
                                                                                    case "a":
                                                                                        v71 = v71
                                                                                    }
                                                                                    v85 := GenS12{F0: 2, F1: 0}
                                                                                    var v86 uint = (v71 ^ 4294967295)
                                                                                    _ = v85
                                                                                    _ = v86
                                                                                }
                                                                                v73 = "a"
                                                                                if false {
                                                                                    var v87 uint = 504
                                                                                    if false {
                                                                                        fmt.Println(v69, fi18, v71)
                                                                                        switch fi18 {
                                                                                        case 32767:
                                                                                            var v88 uint16 = 42
                                                                                            _ = v88
                                                                                        }
                                                                                        var v89 uint16 = 65535
                                                                                        _ = v89
                                                                                    }
                                                                                    _ = len(v15)
                                                                                    _ = v87
                                                                                }
                                                                            } else {
                                                                                var v90 uint = v71
                                                                                switch v73 {
                                                                                case "foo bar":
                                                                                    if true {
                                                                                        var v91 uint16 = 32767
                                                                                        _ = len(v15)
                                                                                        if len(v15) >= 2 {
                                                                                            v15 = v15[1:len(v15)]
                                                                                        }
                                                                                        var v92 int = -83
                                                                                        var v93 string = v73
                                                                                        _ = v91
                                                                                        _ = v92
                                                                                        _ = v93
                                                                                    }
                                                                                default:
                                                                                    if (-1 >= 834) {
                                                                                        var v94 float32 = 3.14159
                                                                                        v73 = v73
                                                                                        fmt.Println(fi68)
                                                                                        _ = v94
                                                                                    }
                                                                                }
                                                                                _ = v90
                                                                            }
                                                                        case "hello":
                                                                            var v95 uint16 = 2
                                                                            _ = v95
                                                                        case "a":
                                                                            if (!(!(!((v71 + v71) >= (v71 + (v71 / (((v71 * (^(v71)))) | 1))))))) {
                                                                                var v96 float32 = 1.5
                                                                                _ = v96
                                                                            }
                                                                        }
                                                                        fmt.Println(v73)
                                                                    }
                                                                }
                                                                v97 := GenS13{F0: "hello"}
                                                                v71 += v71
                                                                v73 = (((v73 + v73) + v73) + "foo bar")
                                                                _ = v97
                                                            } else {
                                                                if (((!false) && (2 <= 0)) && false) {
                                                                    switch (v73 + "a") {
                                                                    case "":
                                                                        var v98 uint = 936
                                                                        _ = v98
                                                                    case "hello":
                                                                        var v99 float64 = (-2.25 / 3.14159)
                                                                        _ = v99
                                                                    default:
                                                                        v71 += (v71 << ((459) & 63))
                                                                    }
                                                                } else {
                                                                    v100 := GenS12{F0: 1, F1: -986}
                                                                    if (!((v71 * (^(v71))) != (v71 - v71))) {
                                                                        if (true || true) {
                                                                            var v101 int8 = -29
                                                                            var v102 float64 = 0.1
                                                                            v103 := GenS11{F0: 7}
                                                                            var v104 int = fi18
                                                                            var v105 float64 = -2.25
                                                                            var v106 int = (fi68 ^ fi68)
                                                                            _ = v101
                                                                            _ = v102
                                                                            _ = v103
                                                                            _ = v104
                                                                            _ = v105
                                                                            _ = v106
                                                                        } else {
                                                                            v71 = 0
                                                                        }
                                                                        if len(v15) > 0 {
                                                                            v15[0] = fi18
                                                                        }
                                                                        fmt.Println(v69)
                                                                        if len(v15) >= 2 {
                                                                            v15 = v15[1:len(v15)]
                                                                        }
                                                                        v69 = 4294967295
                                                                        if false {
                                                                            v73 = v73
                                                                            if len(v15) >= 2 {
                                                                                v15 = v15[1:len(v15)]
                                                                            }
                                                                            v69 = (v69 % (((^(v69))) | 1))
                                                                        } else {
                                                                            fmt.Println(fi68)
                                                                        }
                                                                    }
                                                                    var v107 int = (fi18 / (((fi18 - fi18)) | 1))
                                                                    _ = v100
                                                                    _ = v107
                                                                }
                                                                v71 = (^(v71))
                                                            }
                                                            switch v73 {
                                                            case "":
                                                                fmt.Println(v73, fi68, v71)
                                                            default:
                                                                v73 = v73
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v73
                                            }
                                            _ = v71
                                        }
                                    }
                                    v69 = 2
                                    v69 += (v69 << ((155) & 63))
                                    v69 -= (v69 * (v69 >> ((711) & 63)))
                                case "foo bar":
                                    var v108 int = fi18
                                    _ = v108
                                }
                            case -1:
                                fmt.Println(fi18, fi18)
                            default:
                                _ = len(v15)
                            }
                            _ = fi68
                            _ = v69
                        }
                    }
                    var v109 bool = (-1 < 119)
                    v109 = v109
                    _ = fi18
                    _ = v109
                }
                switch 1 {
                case 127:
                    var v110 uint64 = 550
                    _ = v110
                case 32768:
                    var v111 int32 = 42
                    _ = v111
                default:
                    var v112 uint64 = 0
                    _ = v112
                }
                var v113 bool = (32767 >= 42)
                _ = v113
            } else {
                var v114 uint8 = 255
                v114 += (^(v114))
                for fi115 := 0; fi115 < 8; fi115++ {
                    fmt.Println(fi115)
                    if len(v15) > 0 {
                        v15[0] = (fi115 >> ((4294967295) & 63))
                    }
                    switch (fi115 / (((fi115 & (fi115 << ((1) & 63)))) | 1)) {
                    case 32767:
                        var v116 uint64 = 762
                        _ = v116
                    }
                    _ = fi115
                }
                _ = v114
            }
            v117 := GenS6{F0: 791, F1: 0}
            if false {
                var v118 float32 = (0.0 * 0.1)
                _ = v118
            } else {
                if (true && false) {
                    if (478 > 617) {
                        var v119 bool = true
                        var v120 int16 = 42
                        v120 -= (v120 / (((v120 & v120)) | 1))
                        _ = v119
                        _ = v120
                    } else {
                        var v121 int = -9223372036854775808
                        var v122 float32 = (1e10 + 0.1)
                        _ = v121
                        _ = v122
                    }
                    var v123 uint8 = 0
                    var v124 float32 = 1.5
                    _ = v123
                    _ = v124
                } else {
                    var v125 int16 = 42
                    _ = len(v15)
                    for fi126 := 0; fi126 < 15; fi126++ {
                        v125 -= v125
                        _ = fi126
                    }
                    _ = v125
                }
            }
            _ = v117
        }
        var v127 uint64 = 2
        _ = v127
    } else {
        var v128 int8 = -42
        v128 -= 1
        fmt.Println(v128)
        _ = v128
    }
    _ = v15
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 32767
    v3 := make(map[int]uint32)
    var v4 uint = 42
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]uint32)
    var v129 int = 218
    delete(v6, 2)
    var v130 uint = (v4 << ((v4) & 63))
    var v131 uint8 = 128
    v129 += (v129 & v129)
    switch "hello" {
    case "foo bar":
        var v132 uint16 = v2
        _ = v132
    case "\n\t":
        for fi133 := 0; fi133 < 1; fi133++ {
            for fi134 := 0; fi134 < 18; fi134++ {
                v129 += (fi133 << (((^(v4))) & 63))
                _ = fi134
            }
            _ = fi133
        }
    case "hello":
        switch ("" + "") {
        case "":
            for fi135 := 0; fi135 < 17; fi135++ {
                switch fi135 {
                case -2147483648:
                    switch (^(fi135)) {
                    case 713:
                        fmt.Println(v2, v130)
                    }
                case 128:
                    v4 -= 87
                case 377:
                    _ = len(v5)
                }
                if true {
                    if true {
                        var v136 int16 = -365
                        for fi137 := 0; fi137 < 4; fi137++ {
                            var v138 float64 = 0.1
                            v136 += (v136 << (((v130 * 18446744073709551615)) & 63))
                            v129 -= fi135
                            _ = fi137
                            _ = v138
                        }
                        _ = v136
                    } else {
                        v139 := GenS12{F0: 2, F1: -692}
                        for fi140 := 0; fi140 < 11; fi140++ {
                            v130 = (^(v130))
                            _ = fi140
                        }
                        for fi141 := 0; fi141 < 17; fi141++ {
                            v142 := GenS12{F0: v131, F1: -245}
                            _ = fi141
                            _ = v142
                        }
                        _ = v139
                    }
                    for fi143 := 0; fi143 < 14; fi143++ {
                        var v144 float64 = 3.14159
                        if false {
                            v6[0] = 33
                        }
                        _ = fi143
                        _ = v144
                    }
                    v145 := GenS2{F0: 0, F1: -1, F2: 127}
                    _ = v145
                } else {
                    v129 += 1
                    v131 += (^(v131))
                }
                v146 := GenS14{F0: -1, F1: 9223372036854775807}
                _ = fi135
                _ = v146
            }
        case "\n\t":
            var v147 uint = v4
            _ = v147
        }
    default:
        var v148 float64 = 3.14159
        _ = v148
    }
    v131 -= v131
    v149 := GenS15{F0: "a"}
    v150 := GenS16{F0: v131, F1: 0, F2: 0}
    if (false || (93 == v131)) {
        v131 = uint8(v4)
        var v151 uint8 = v131
        _ = v151
    }
    if false {
        v3[497] = 279
        if true {
            switch (("\n\t" + (("a" + "\n\t") + "foo bar")) + ("\n\t" + ("hello" + ((("hello" + "\n\t") + "a") + ("\n\t" + ""))))) {
            case "\n\t":
                v131 -= 240
            default:
                var v152 uint16 = (v2 & uint16(v130))
                _ = v152
            }
            _ = len(v3)
        }
    }
    var v153 uint64 = 0
    for fi154 := 0; fi154 < 1; fi154++ {
        var v155 uint16 = v2
        var v156 bool = (false && false)
        _ = fi154
        _ = v155
        _ = v156
    }
    if false {
        if false {
            var v157 int16 = 1
            var v158 uint = 4294967295
            for fi159 := 0; fi159 < 4; fi159++ {
                v158 -= v130
                v160 := GenS8{F0: 440}
                _ = fi159
                _ = v160
            }
            _ = v157
            _ = v158
        } else {
            for fi161 := 0; fi161 < 8; fi161++ {
                v129 += 32768
                switch "foo bar" {
                case "\n\t":
                    var v162 int8 = 0
                    _ = v162
                }
                _ = fi161
            }
            v163 := []int{(v129 / ((v129) | 1)), v129, -292}
            _ = v163
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v129:", v129)
    fmt.Println("v130:", v130)
    fmt.Println("v131:", v131)
    fmt.Println("v153:", v153)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v129
    _ = v130
    _ = v131
    _ = v149
    _ = v150
    _ = v153
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
