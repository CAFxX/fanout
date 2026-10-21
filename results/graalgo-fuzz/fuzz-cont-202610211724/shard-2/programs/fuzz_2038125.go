package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint; F1 uint32 }
type GenS9 struct { F0 int32; F1 bool; F2 uint64 }
type GenS10 struct { F0 int32; F1 int16; F2 int }
type GenS11 struct { F0 string; F1 uint16; F2 uint16 }
type GenS12 struct { F0 uint; F1 string }
type GenS13 struct { F0 string; F1 uint64 }
type GenS14 struct { F0 uint8; F1 uint8; F2 int }
type GenS15 struct { F0 int8; F1 int8; F2 uint64 }
type GenS16 struct { F0 uint; F1 uint8; F2 bool }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int32; F1 string }
type GenS20 struct { F0 int8; F1 uint8; F2 bool }
type GenS21 struct { F0 int16; F1 bool }
type GenS22 struct { F0 int }
type GenS23 struct { F0 int16; F1 int32 }
type GenS24 struct { F0 int; F1 bool; F2 int }
type GenS25 struct { F0 bool; F1 string }
type GenS26 struct { F0 uint }
type GenS27 struct { F0 int64; F1 string }
type GenS28 struct { F0 uint64; F1 string }
type GenS29 struct { F0 uint; F1 int64 }
type GenS30 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS31 struct { F0 int64; F1 uint }
func genFn0(p0 uint, p1 int64) int32 {
    p1 -= 1
    for fi7 := 0; fi7 < 20; fi7++ {
        if false {
            v8 := []int{-265, (fi7 * fi7)}
            for fi9 := 0; fi9 < 1; fi9++ {
                _ = len(v8)
                _ = fi9
            }
            _ = v8
        }
        var v10 string = "hello"
        _ = fi7
        _ = v10
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 string = ""
    var v2 float32 = 0.1
    var v3 string = (v1 + v1)
    v4 := []uint32{65535}
    var v5 int8 = -73
    var v6 int8 = v5
    var v11 bool = true
    var v12 string = v1
    v13 := GenS1{F0: 1, F1: (("a" + v1) + ("foo bar" + ("\n\t" + ("hello" + (v1 + v3)))))}
    _ = len(v4)
    v12 = v1
    if true {
        switch -1 {
        case 32767:
            fmt.Println(v6)
        case 1:
            _ = len(v4)
        case 42:
            v6 += (v5 << ((65535) & 63))
        default:
            fmt.Println(v3)
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            if ((-563 < -1) || false) {
                switch ((((v1 + v1) + "\n\t") + "") + "") {
                case "foo bar":
                    for fi15 := 0; fi15 < 4; fi15++ {
                        var v16 int = (-(fi14))
                        switch -9223372036854775808 {
                        case 2147483647:
                            if true {
                                v17 := GenS1{F0: 32767, F1: v3}
                                if (128 == 1) {
                                    if ((40 >= 0) || (!v11)) {
                                        var v18 int64 = 625
                                        if (true || (!v11)) {
                                            switch v16 {
                                            case 25:
                                                v11 = (v6 >= (v6 ^ v5))
                                            }
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = 4294967295
                                        }
                                        _ = v18
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        switch v12 {
                                        case "foo bar":
                                            var v19 uint32 = 150
                                            _ = v19
                                        }
                                        var v20 uint = 897
                                        _ = v20
                                    }
                                    if (!v11) {
                                        if true {
                                            if ("\n\t" >= ("" + ("foo bar" + ("a" + (v1 + v3))))) {
                                                v16 = (^(fi14))
                                                var v21 uint32 = 2147483648
                                                _ = v21
                                            } else {
                                                switch 2147483647 {
                                                case 0:
                                                    var v22 bool = false
                                                    var v23 int = (v16 << ((42) & 63))
                                                    switch (v1 + "") {
                                                    case "\n\t":
                                                        v24 := GenS2{F0: 538}
                                                        _ = v24
                                                    case "foo bar":
                                                        if (v12 >= v12) {
                                                            switch v16 {
                                                            case 42:
                                                                if v22 {
                                                                    if (1e10 <= 3.14159) {
                                                                        if (494 != 763) {
                                                                            v16 = (-(fi15))
                                                                            v6 -= (v5 & 0)
                                                                            var v25 float32 = v2
                                                                            var v26 float64 = ((3.14159 + 1.5) / 1.0)
                                                                            _ = v25
                                                                            _ = v26
                                                                        } else {
                                                                            if v11 {
                                                                                if (255 != 1) {
                                                                                    v27 := GenS2{F0: 849}
                                                                                    var v28 int8 = -1
                                                                                    _ = v27
                                                                                    _ = v28
                                                                                } else {
                                                                                    var v29 uint64 = 18446744073709551615
                                                                                    _ = v29
                                                                                }
                                                                                switch (-(v16)) {
                                                                                case 4294967295:
                                                                                    var v30 uint = 255
                                                                                    _ = v30
                                                                                case 42:
                                                                                    if v11 {
                                                                                        var v31 int16 = 0
                                                                                        _ = v31
                                                                                    }
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = 850
                                                                                    }
                                                                                    if v22 {
                                                                                        var v32 uint64 = 2
                                                                                        var v33 uint8 = 0
                                                                                        _ = v32
                                                                                        _ = v33
                                                                                    } else {
                                                                                        if true {
                                                                                            v11 = v22
                                                                                            v16 -= fi14
                                                                                            v34 := GenS3{F0: -1}
                                                                                            v16 -= fi15
                                                                                            fmt.Println(v22, v1)
                                                                                            _ = v34
                                                                                        } else {
                                                                                            v35 := GenS1{F0: (fi14 / ((32768) | 1)), F1: "hello"}
                                                                                            if len(v4) >= 2 {
                                                                                                v4 = v4[1:len(v4)]
                                                                                            }
                                                                                            _ = v35
                                                                                        }
                                                                                    }
                                                                                    var v36 uint64 = 444
                                                                                    v6 += v5
                                                                                    _ = v36
                                                                                }
                                                                                v4 = append(v4, 437)
                                                                            }
                                                                        }
                                                                        v4 = append(v4, 2147483647)
                                                                        v37 := GenS4{F0: (v22 && false), F1: 1}
                                                                        _ = v37
                                                                    } else {
                                                                        switch v12 {
                                                                        case "a":
                                                                            var v38 int16 = 1
                                                                            _ = v38
                                                                        case "hello":
                                                                            if v22 {
                                                                                var v39 uint32 = 4294967295
                                                                                v12 = "a"
                                                                                var v40 float64 = (0.0 - (3.14159 - ((0.1 - 0.1) - (-2.25 / 0.1))))
                                                                                _ = v39
                                                                                _ = v40
                                                                            } else {
                                                                                switch fi14 {
                                                                                case 32768:
                                                                                    switch (v12 + v12) {
                                                                                    case "a":
                                                                                        var v41 uint8 = 42
                                                                                        _ = v41
                                                                                    case "":
                                                                                        v22 = v22
                                                                                    }
                                                                                default:
                                                                                    var v42 int32 = -1
                                                                                    _ = v42
                                                                                }
                                                                                fmt.Println(fi14, fi15)
                                                                            }
                                                                        case "foo bar":
                                                                            var v43 uint = 304
                                                                            _ = v43
                                                                        default:
                                                                            v6 -= (v5 ^ (v5 | v5))
                                                                        }
                                                                        v23 = v23
                                                                    }
                                                                }
                                                            default:
                                                                if (!v22) {
                                                                    v44 := GenS5{F0: (((float64(1e10) - 1.5) == -2.25) && (255 >= 412)), F1: 42}
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 1
                                                                    }
                                                                    v5 += (v6 / (((v5 << ((0) & 63))) | 1))
                                                                    _ = v44
                                                                } else {
                                                                    fmt.Println(v3, v1)
                                                                }
                                                            }
                                                            v4 = append(v4, 2147483648)
                                                            if (!v11) {
                                                                switch (-(fi15)) {
                                                                case 167:
                                                                    var v45 int8 = (^(v5))
                                                                    _ = v45
                                                                case -263:
                                                                    switch 42 {
                                                                    case -2147483648:
                                                                        if (-1 <= -1) {
                                                                            v22 = v11
                                                                            var v46 float64 = 0.0
                                                                            switch v12 {
                                                                            case "a":
                                                                                if false {
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = 1
                                                                                    }
                                                                                    var v47 int64 = 42
                                                                                    v22 = (v2 <= 1.5)
                                                                                    v23 = fi14
                                                                                    _ = v47
                                                                                } else {
                                                                                    var v48 float32 = v2
                                                                                    v1 = v1
                                                                                    _ = v48
                                                                                }
                                                                            case "hello":
                                                                                if v22 {
                                                                                    switch v12 {
                                                                                    case "a":
                                                                                        v6 += (v6 * v6)
                                                                                        var v49 int32 = genFn0(4294967295, 1)
                                                                                        v50 := GenS3{F0: -1}
                                                                                        _ = v49
                                                                                        _ = v50
                                                                                    default:
                                                                                        var v51 float32 = 0.0
                                                                                        _ = v51
                                                                                    }
                                                                                }
                                                                            case "\n\t":
                                                                                fmt.Println(v2, v16)
                                                                            default:
                                                                                if v22 {
                                                                                    var v52 uint16 = 675
                                                                                    if (v46 < (float64(1e10) / 1.0)) {
                                                                                        switch ((v1 + (v1 + ((v3 + "a") + ((v12 + ((((v1 + v3) + v12) + v12) + "a")) + v3)))) + v3) {
                                                                                        case "":
                                                                                            v53 := GenS4{F0: v22, F1: -482}
                                                                                            _ = v53
                                                                                        }
                                                                                        v54 := GenS4{F0: true, F1: 42}
                                                                                        _ = v54
                                                                                    }
                                                                                    v4 = append(v4, 2147483647)
                                                                                    v4 = append(v4, 572)
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = 180
                                                                                    }
                                                                                    _ = v52
                                                                                }
                                                                            }
                                                                            var v55 int = v23
                                                                            _ = v46
                                                                            _ = v55
                                                                        }
                                                                        if ((!false) && v11) {
                                                                            v1 = (("" + "foo bar") + v1)
                                                                            if ((v16 >= v16) && ((true && (!v11)) && ("" == (v1 + (v12 + v12))))) {
                                                                                var v56 int8 = v5
                                                                                if (0 > 4294967295) {
                                                                                    var v57 string = (("a" + v1) + v12)
                                                                                    _ = v57
                                                                                } else {
                                                                                    switch v3 {
                                                                                    case "hello":
                                                                                        v58 := GenS3{F0: 556}
                                                                                        _ = len(v4)
                                                                                        var v59 uint = 42
                                                                                        _ = v58
                                                                                        _ = v59
                                                                                    case "a":
                                                                                        var v60 float32 = (float32(1.5) + v2)
                                                                                        _ = v60
                                                                                    case "\n\t":
                                                                                        v56 = v56
                                                                                    default:
                                                                                        _ = len(v4)
                                                                                    }
                                                                                }
                                                                                switch v23 {
                                                                                case 963:
                                                                                    switch int(v5) {
                                                                                    case 1:
                                                                                        v23 -= v16
                                                                                    }
                                                                                case 128:
                                                                                    var v61 int64 = -9223372036854775808
                                                                                    _ = v61
                                                                                case 42:
                                                                                    fmt.Println(fi14, fi14)
                                                                                    _ = len(v4)
                                                                                }
                                                                                _ = v56
                                                                            } else {
                                                                                var v62 uint64 = 0
                                                                                _ = v62
                                                                            }
                                                                            v22 = true
                                                                            if (!(!(179 == 171))) {
                                                                                v22 = ((0.0 <= 1.5) && true)
                                                                                fmt.Println(v5, v22)
                                                                            } else {
                                                                                switch "hello" {
                                                                                case "":
                                                                                    var v63 int = (fi15 - fi15)
                                                                                    _ = v63
                                                                                case "foo bar":
                                                                                    v64 := GenS4{F0: (!v11), F1: 42}
                                                                                    v3 = ("foo bar" + v12)
                                                                                    _ = v64
                                                                                default:
                                                                                    var v65 bool = false
                                                                                    _ = v65
                                                                                }
                                                                                if ((((v3 + ("a" + v1)) + v3) <= v12) || (true || (-534 > 9223372036854775807))) {
                                                                                    fmt.Println(v1, v23, v6)
                                                                                    var v66 int8 = v6
                                                                                    var v67 uint16 = 825
                                                                                    _ = v66
                                                                                    _ = v67
                                                                                } else {
                                                                                    switch v12 {
                                                                                    case "hello":
                                                                                        if (v11 && (952 < 2)) {
                                                                                            v68 := GenS4{F0: false, F1: 9223372036854775807}
                                                                                            if true {
                                                                                                if len(v4) >= 2 {
                                                                                                    v4 = v4[1:len(v4)]
                                                                                                }
                                                                                                switch 1 {
                                                                                                case -2147483648:
                                                                                                    if len(v4) > 0 {
                                                                                                        v4[0] = 65535
                                                                                                    }
                                                                                                case -355:
                                                                                                    v11 = ("foo bar" <= v1)
                                                                                                default:
                                                                                                    v69 := GenS3{F0: 1}
                                                                                                    _ = v69
                                                                                                }
                                                                                                v70 := GenS5{F0: v11, F1: -1}
                                                                                                v12 = (v1 + (v12 + ((v3 + (v1 + "hello")) + ("a" + v1))))
                                                                                                _ = v70
                                                                                            } else {
                                                                                                v71 := GenS3{F0: -802}
                                                                                                if ((18446744073709551615 == 2) || v22) {
                                                                                                    if v22 {
                                                                                                        switch fi14 {
                                                                                                        case -997:
                                                                                                            if (827 >= 42) {
                                                                                                                v72 := GenS6{F0: 0}
                                                                                                                _ = v72
                                                                                                            } else {
                                                                                                                v1 = "foo bar"
                                                                                                                if (-2147483648 <= 1) {
                                                                                                                    v73 := GenS3{F0: 9223372036854775807}
                                                                                                                    var v74 float64 = 1e10
                                                                                                                    _ = v73
                                                                                                                    _ = v74
                                                                                                                }
                                                                                                                fmt.Println(v12)
                                                                                                            }
                                                                                                        case 0:
                                                                                                            switch "\n\t" {
                                                                                                            case "foo bar":
                                                                                                                _ = len(v4)
                                                                                                            case "hello":
                                                                                                                if false {
                                                                                                                    v12 = v12
                                                                                                                    if len(v4) >= 2 {
                                                                                                                        v4 = v4[1:len(v4)]
                                                                                                                    }
                                                                                                                    if (-2.25 != (-2.25 * 0.1)) {
                                                                                                                        fmt.Println(v16, fi15)
                                                                                                                        var v75 float32 = 1.5
                                                                                                                        _ = v75
                                                                                                                    } else {
                                                                                                                        v4 = append(v4, 1)
                                                                                                                        if len(v4) >= 2 {
                                                                                                                            v4 = v4[1:len(v4)]
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                v76 := GenS2{F0: 18446744073709551615}
                                                                                                                _ = v76
                                                                                                            }
                                                                                                        }
                                                                                                        v77 := GenS7{F0: 65535}
                                                                                                        _ = v77
                                                                                                    } else {
                                                                                                        var v78 bool = ((928 < 568) && (1 > 866))
                                                                                                        _ = v78
                                                                                                    }
                                                                                                }
                                                                                                _ = v71
                                                                                            }
                                                                                            var v79 uint32 = 2147483648
                                                                                            var v80 int16 = 42
                                                                                            fmt.Println(v12, v79, v80)
                                                                                            _ = v68
                                                                                            _ = v79
                                                                                            _ = v80
                                                                                        }
                                                                                    case "":
                                                                                        if v22 {
                                                                                            var v81 bool = v22
                                                                                            var v82 bool = (!(!v81))
                                                                                            v5 -= (v5 ^ (-(v6)))
                                                                                            _ = v81
                                                                                            _ = v82
                                                                                        }
                                                                                    case "foo bar":
                                                                                        switch v1 {
                                                                                        case "\n\t":
                                                                                            var v83 uint8 = 24
                                                                                            _ = v83
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v2, v16)
                                                                                    var v84 uint64 = 1
                                                                                    _ = v84
                                                                                }
                                                                                v2 = (v2 * (v2 - float32((1.5 - ((1e10 + 3.14159) / 1e10)))))
                                                                            }
                                                                            v4 = append(v4, 0)
                                                                        }
                                                                    default:
                                                                        fmt.Println(fi15)
                                                                    }
                                                                case -9223372036854775808:
                                                                    v2 += 0.1
                                                                }
                                                                var v85 int16 = 1
                                                                v22 = v22
                                                                switch fi14 {
                                                                case -342:
                                                                    v86 := GenS8{F0: 255, F1: 2147483647}
                                                                    _ = v86
                                                                case 292:
                                                                    if (!false) {
                                                                        fmt.Println(v6, v23, v16)
                                                                    } else {
                                                                        fmt.Println(fi14, v23)
                                                                        if (553 <= (v23 << ((353) & 63))) {
                                                                            v87 := GenS9{F0: 0, F1: v11, F2: 0}
                                                                            var v88 int16 = v85
                                                                            _ = v87
                                                                            _ = v88
                                                                        } else {
                                                                            fmt.Println(v22, v16)
                                                                        }
                                                                        v85 -= v85
                                                                    }
                                                                case 128:
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 1
                                                                    }
                                                                }
                                                                var v89 float32 = 3.14159
                                                                v6 -= (v6 | 1)
                                                                _ = v85
                                                                _ = v89
                                                            }
                                                        }
                                                        if len(v4) > 0 {
                                                            v4[0] = 0
                                                        }
                                                        var v90 int8 = v5
                                                        _ = v90
                                                    case "":
                                                        var v91 uint32 = 171
                                                        _ = v91
                                                    default:
                                                        v92 := GenS1{F0: 1, F1: "foo bar"}
                                                        _ = v92
                                                    }
                                                    v5 += (v6 << ((18446744073709551615) & 63))
                                                    _ = v22
                                                    _ = v23
                                                case -1:
                                                    if (false || ((42 >= 1) || (0 != 536))) {
                                                        if ((-2.25 * 3.14159) < (3.14159 - 3.14159)) {
                                                            fmt.Println(v2)
                                                            v93 := GenS5{F0: ((((false && (fi15 == fi14)) && v11) && v11) && (2147483647 <= 569)), F1: -458}
                                                            if ((4294967295 < 185) && v11) {
                                                                switch -140 {
                                                                case -605:
                                                                    if v11 {
                                                                        var v94 uint16 = 255
                                                                        if v11 {
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                v4 = append(v4, 790)
                                                                            case "a":
                                                                                if v11 {
                                                                                    v1 = v1
                                                                                    v94 += v94
                                                                                    v11 = (2 <= 648)
                                                                                } else {
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = 0
                                                                                    }
                                                                                }
                                                                            default:
                                                                                var v95 int64 = 0
                                                                                _ = v95
                                                                            }
                                                                            switch (fi14 * (v16 << ((255) & 63))) {
                                                                            case 42:
                                                                                v96 := GenS5{F0: v11, F1: -1}
                                                                                _ = v96
                                                                            case 127:
                                                                                if v11 {
                                                                                    if true {
                                                                                        if v11 {
                                                                                            switch v16 {
                                                                                            case -1:
                                                                                                v5 = (v6 << ((612) & 63))
                                                                                            case 4294967295:
                                                                                                var v97 uint32 = 730
                                                                                                if len(v4) >= 2 {
                                                                                                    v4 = v4[1:len(v4)]
                                                                                                }
                                                                                                _ = v97
                                                                                            default:
                                                                                                var v98 uint16 = 2
                                                                                                _ = v98
                                                                                            }
                                                                                            var v99 float64 = -2.25
                                                                                            _ = v99
                                                                                        } else {
                                                                                            var v100 int = v16
                                                                                            var v101 int8 = (v5 & v6)
                                                                                            _ = v100
                                                                                            _ = v101
                                                                                        }
                                                                                        switch (v3 + v12) {
                                                                                        case "hello":
                                                                                            _ = len(v4)
                                                                                        case "a":
                                                                                            var v102 float32 = v2
                                                                                            _ = v102
                                                                                        }
                                                                                        v12 = v12
                                                                                    }
                                                                                    fmt.Println(v6, v94, fi14)
                                                                                    _ = len(v4)
                                                                                    if v11 {
                                                                                        if (true && (v11 && v11)) {
                                                                                            if (161 != 172) {
                                                                                                if len(v4) > 0 {
                                                                                                    v4[0] = 2147483648
                                                                                                }
                                                                                                switch fi15 {
                                                                                                case 2147483647:
                                                                                                    var v103 int16 = 972
                                                                                                    var v104 int16 = 42
                                                                                                    v105 := GenS3{F0: -1}
                                                                                                    var v106 bool = (genFn0(42, 42) != 42)
                                                                                                    v12 = ("" + "\n\t")
                                                                                                    _ = v103
                                                                                                    _ = v104
                                                                                                    _ = v105
                                                                                                    _ = v106
                                                                                                case -404:
                                                                                                    v1 = (v3 + "foo bar")
                                                                                                    if (v16 == -1) {
                                                                                                        var v107 float64 = 0.0
                                                                                                        _ = v107
                                                                                                    } else {
                                                                                                        fmt.Println(v11, v16, v12)
                                                                                                        if len(v4) >= 2 {
                                                                                                            v4 = v4[1:len(v4)]
                                                                                                        }
                                                                                                        switch v12 {
                                                                                                        case "foo bar":
                                                                                                            var v108 float64 = -2.25
                                                                                                            _ = v108
                                                                                                        case "hello":
                                                                                                            v11 = v11
                                                                                                        case "":
                                                                                                            v109 := GenS10{F0: -1, F1: -936, F2: v16}
                                                                                                            _ = v109
                                                                                                        }
                                                                                                    }
                                                                                                case 4294967295:
                                                                                                    switch (^(fi14)) {
                                                                                                    case 9223372036854775807:
                                                                                                        v11 = (42 >= 65535)
                                                                                                    case -504:
                                                                                                        v110 := GenS5{F0: true, F1: 663}
                                                                                                        _ = v110
                                                                                                    }
                                                                                                }
                                                                                                v111 := GenS3{F0: -1}
                                                                                                var v112 uint16 = 255
                                                                                                var v113 int = (fi15 + 944)
                                                                                                switch (fi14 * v113) {
                                                                                                case 32768:
                                                                                                    v114 := GenS5{F0: v11, F1: 9223372036854775807}
                                                                                                    _ = v114
                                                                                                case 0:
                                                                                                    if false {
                                                                                                        v112 = v112
                                                                                                        switch "hello" {
                                                                                                        case "":
                                                                                                            switch v16 {
                                                                                                            case 2147483647:
                                                                                                                var v115 uint8 = 255
                                                                                                                _ = v115
                                                                                                            case 42:
                                                                                                                var v116 uint16 = v112
                                                                                                                _ = v116
                                                                                                            case 0:
                                                                                                                var v117 float32 = v2
                                                                                                                _ = v117
                                                                                                            default:
                                                                                                                switch ("foo bar" + v1) {
                                                                                                                case "hello":
                                                                                                                    v1 = v3
                                                                                                                case "a":
                                                                                                                    v94 = v94
                                                                                                                case "\n\t":
                                                                                                                    switch v113 {
                                                                                                                    case -2147483648:
                                                                                                                        v118 := GenS6{F0: 0}
                                                                                                                        _ = v118
                                                                                                                    case -9223372036854775808:
                                                                                                                        var v119 uint64 = 237
                                                                                                                        v2 += (1e10 + (float32((0.0 - ((3.14159 + float64(v2)) * ((-2.25 / (1.5 - 3.14159)) + ((0.1 / (float64(((((1e10 - -2.25) * 0.1) / (1.5 + v2)) - (v2 + 1e10))) * 3.14159)) / float64(v2)))))) + v2))
                                                                                                                        if len(v4) >= 2 {
                                                                                                                            v4 = v4[1:len(v4)]
                                                                                                                        }
                                                                                                                        if len(v4) >= 2 {
                                                                                                                            v4 = v4[1:len(v4)]
                                                                                                                        }
                                                                                                                        if len(v4) >= 2 {
                                                                                                                            v4 = v4[1:len(v4)]
                                                                                                                        }
                                                                                                                        _ = v119
                                                                                                                    case 4294967295:
                                                                                                                        switch 42 {
                                                                                                                        case -1:
                                                                                                                            v3 = v3
                                                                                                                        case 32768:
                                                                                                                            var v120 int16 = 625
                                                                                                                            _ = v120
                                                                                                                        case -9223372036854775808:
                                                                                                                            if (255 <= 123) {
                                                                                                                                if v11 {
                                                                                                                                    v121 := GenS8{F0: 216, F1: 2147483648}
                                                                                                                                    var v122 int = (v16 << ((65535) & 63))
                                                                                                                                    fmt.Println(v11, v16)
                                                                                                                                    _ = v121
                                                                                                                                    _ = v122
                                                                                                                                } else {
                                                                                                                                    var v123 int = (fi15 << ((611) & 63))
                                                                                                                                    var v124 uint = 42
                                                                                                                                    if v11 {
                                                                                                                                        v4 = append(v4, uint32(v112))
                                                                                                                                    }
                                                                                                                                    if (42 != 42) {
                                                                                                                                        fmt.Println(fi15, v123)
                                                                                                                                    }
                                                                                                                                    _ = v123
                                                                                                                                    _ = v124
                                                                                                                                }
                                                                                                                                switch v1 {
                                                                                                                                case "hello":
                                                                                                                                    var v125 uint = 1
                                                                                                                                    _ = v125
                                                                                                                                case "\n\t":
                                                                                                                                    var v126 uint64 = 557
                                                                                                                                    _ = v126
                                                                                                                                default:
                                                                                                                                    _ = len(v4)
                                                                                                                                }
                                                                                                                                var v127 int64 = -139
                                                                                                                                _ = v127
                                                                                                                            } else {
                                                                                                                                v128 := GenS11{F0: "a", F1: 2, F2: 0}
                                                                                                                                fmt.Println(v16, v1)
                                                                                                                                v129 := GenS12{F0: 2, F1: ""}
                                                                                                                                _ = v128
                                                                                                                                _ = v129
                                                                                                                            }
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        switch 0 {
                                                                                                                        case -979:
                                                                                                                            if v11 {
                                                                                                                                var v130 uint32 = 65535
                                                                                                                                _ = v130
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            switch (v113 << ((4294967295) & 63)) {
                                                                                                            case 409:
                                                                                                                var v131 float64 = 0.1
                                                                                                                _ = v131
                                                                                                            case 32768:
                                                                                                                v4 = append(v4, 2147483648)
                                                                                                            default:
                                                                                                                var v132 string = "a"
                                                                                                                _ = v132
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                case 32767:
                                                                                                    switch (v3 + v12) {
                                                                                                    case "foo bar":
                                                                                                        v3 = ("a" + v1)
                                                                                                    case "a":
                                                                                                        v133 := GenS13{F0: ("" + ("hello" + v1)), F1: 458}
                                                                                                        _ = v133
                                                                                                    }
                                                                                                }
                                                                                                if (!(fi15 != v16)) {
                                                                                                    var v134 uint8 = 255
                                                                                                    v16 += v16
                                                                                                    switch ("hello" + v12) {
                                                                                                    case "":
                                                                                                        v11 = (!(fi15 >= (v16 | (fi15 >> ((42) & 63)))))
                                                                                                        var v135 uint = 1
                                                                                                        _ = v135
                                                                                                    case "foo bar":
                                                                                                        fmt.Println(v5)
                                                                                                    case "\n\t":
                                                                                                        fmt.Println(fi15, v94)
                                                                                                    }
                                                                                                    _ = v134
                                                                                                } else {
                                                                                                    v4 = append(v4, 4294967295)
                                                                                                }
                                                                                                _ = v111
                                                                                                _ = v112
                                                                                                _ = v113
                                                                                            } else {
                                                                                                v136 := GenS14{F0: 1, F1: 2, F2: (-(v16))}
                                                                                                if true {
                                                                                                    if (fi14 > 42) {
                                                                                                        var v137 int64 = 0
                                                                                                        _ = v137
                                                                                                    }
                                                                                                }
                                                                                                v16 += 127
                                                                                                _ = v136
                                                                                            }
                                                                                            var v138 int64 = -9223372036854775808
                                                                                            _ = len(v4)
                                                                                            _ = v138
                                                                                        } else {
                                                                                            v1 = v12
                                                                                            v5 += v5
                                                                                            fmt.Println(fi14)
                                                                                        }
                                                                                        var v139 uint32 = 2147483648
                                                                                        fmt.Println(v139)
                                                                                        switch (v1 + v12) {
                                                                                        case "\n\t":
                                                                                            v94 -= 255
                                                                                        case "foo bar":
                                                                                            switch -9223372036854775808 {
                                                                                            case 42:
                                                                                                _ = len(v4)
                                                                                                var v140 uint = 2
                                                                                                fmt.Println(v6, v16, v6)
                                                                                                fmt.Println(fi15, fi14)
                                                                                                _ = v140
                                                                                            default:
                                                                                                if true {
                                                                                                    if v11 {
                                                                                                        v141 := GenS7{F0: 0}
                                                                                                        switch fi14 {
                                                                                                        case 4294967295:
                                                                                                            if v11 {
                                                                                                                switch (v16 % ((fi15) | 1)) {
                                                                                                                case 32768:
                                                                                                                    if len(v4) >= 2 {
                                                                                                                        v4 = v4[1:len(v4)]
                                                                                                                    }
                                                                                                                default:
                                                                                                                    var v142 int = 127
                                                                                                                    _ = v142
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v141
                                                                                                    } else {
                                                                                                        if (370 > uint64(v94)) {
                                                                                                            switch (v16 ^ fi15) {
                                                                                                            case 4294967295:
                                                                                                                if v11 {
                                                                                                                    if (128 <= 128) {
                                                                                                                        fmt.Println(v12)
                                                                                                                        var v143 uint16 = (^(v94))
                                                                                                                        var v144 uint32 = v139
                                                                                                                        _ = v143
                                                                                                                        _ = v144
                                                                                                                    } else {
                                                                                                                        if v11 {
                                                                                                                            var v145 uint32 = 2147483648
                                                                                                                            _ = v145
                                                                                                                        }
                                                                                                                        v1 = v12
                                                                                                                        var v146 uint64 = 4294967295
                                                                                                                        _ = v146
                                                                                                                    }
                                                                                                                    var v147 bool = true
                                                                                                                    switch fi14 {
                                                                                                                    case -2147483648:
                                                                                                                        fmt.Println(v94)
                                                                                                                    default:
                                                                                                                        if len(v4) >= 2 {
                                                                                                                            v4 = v4[1:len(v4)]
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v147
                                                                                                                } else {
                                                                                                                    var v148 uint32 = v139
                                                                                                                    switch (fi14 | (fi14 >> ((65535) & 63))) {
                                                                                                                    case 0:
                                                                                                                        if (v11 && false) {
                                                                                                                            if (-1 > -32768) {
                                                                                                                                v149 := GenS15{F0: v5, F1: 1, F2: 369}
                                                                                                                                v150 := GenS5{F0: ((!v11) || (777 != 65535)), F1: 1}
                                                                                                                                switch "a" {
                                                                                                                                case "\n\t":
                                                                                                                                    v4 = append(v4, v148)
                                                                                                                                case "a":
                                                                                                                                    v151 := GenS11{F0: "hello", F1: v94, F2: (v94 & 984)}
                                                                                                                                    _ = v151
                                                                                                                                case "":
                                                                                                                                    if len(v4) > 0 {
                                                                                                                                        v4[0] = (v148 << ((255) & 63))
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    v12 = v12
                                                                                                                                }
                                                                                                                                _ = v149
                                                                                                                                _ = v150
                                                                                                                            } else {
                                                                                                                                v4 = append(v4, (v148 << ((18446744073709551615) & 63)))
                                                                                                                                v3 = (("" + "") + "foo bar")
                                                                                                                                var v152 string = v12
                                                                                                                                _ = v152
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            switch -572 {
                                                                                                                            case 0:
                                                                                                                                var v153 int8 = v5
                                                                                                                                _ = v153
                                                                                                                            }
                                                                                                                            if (v148 != v139) {
                                                                                                                                var v154 uint8 = 255
                                                                                                                                v155 := GenS7{F0: 612}
                                                                                                                                if (!v11) {
                                                                                                                                    fmt.Println(fi14, v139, v94)
                                                                                                                                    var v156 uint32 = v139
                                                                                                                                    if len(v4) >= 2 {
                                                                                                                                        v4 = v4[1:len(v4)]
                                                                                                                                    }
                                                                                                                                    _ = v156
                                                                                                                                }
                                                                                                                                _ = v154
                                                                                                                                _ = v155
                                                                                                                            }
                                                                                                                        }
                                                                                                                    case 250:
                                                                                                                        switch (v16 >> ((18446744073709551615) & 63)) {
                                                                                                                        case 42:
                                                                                                                            switch ((v12 + "hello") + "a") {
                                                                                                                            case "hello":
                                                                                                                                fmt.Println(v148)
                                                                                                                            case "\n\t":
                                                                                                                                var v157 bool = (191 < 11)
                                                                                                                                if (int32(v94) == 42) {
                                                                                                                                    if v11 {
                                                                                                                                        var v158 int16 = 690
                                                                                                                                        _ = v158
                                                                                                                                    }
                                                                                                                                    if len(v4) >= 2 {
                                                                                                                                        v4 = v4[1:len(v4)]
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    var v159 uint32 = v148
                                                                                                                                    fmt.Println(v16, v157)
                                                                                                                                    v160 := GenS16{F0: 626, F1: 0, F2: (!(!v11))}
                                                                                                                                    _ = len(v4)
                                                                                                                                    var v161 int32 = genFn0(214, 9223372036854775807)
                                                                                                                                    _ = v159
                                                                                                                                    _ = v160
                                                                                                                                    _ = v161
                                                                                                                                }
                                                                                                                                _ = v157
                                                                                                                            case "":
                                                                                                                                v94 += v94
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            switch -416 {
                                                                                                                            case -472:
                                                                                                                                v5 = 42
                                                                                                                            }
                                                                                                                            switch 127 {
                                                                                                                            case 4294967295:
                                                                                                                                v94 -= 42
                                                                                                                            case -338:
                                                                                                                                fmt.Println(v11)
                                                                                                                            case 32767:
                                                                                                                                fmt.Println(v148)
                                                                                                                            }
                                                                                                                            v16 += v16
                                                                                                                        }
                                                                                                                    case -9223372036854775808:
                                                                                                                        if (42 > 814) {
                                                                                                                            var v162 uint32 = v139
                                                                                                                            fmt.Println(v148)
                                                                                                                            v163 := GenS16{F0: 18446744073709551615, F1: 255, F2: v11}
                                                                                                                            _ = v162
                                                                                                                            _ = v163
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        v164 := GenS6{F0: 230}
                                                                                                                        _ = v164
                                                                                                                    }
                                                                                                                    _ = v148
                                                                                                                }
                                                                                                            case -140:
                                                                                                                fmt.Println(fi15, v94, fi14)
                                                                                                            }
                                                                                                        }
                                                                                                        v165 := GenS12{F0: 42, F1: "hello"}
                                                                                                        _ = v165
                                                                                                    }
                                                                                                    if (127 != 127) {
                                                                                                        if len(v4) > 0 {
                                                                                                            v4[0] = 42
                                                                                                        }
                                                                                                        switch 2147483647 {
                                                                                                        case -2147483648:
                                                                                                            v6 = (v6 << ((42) & 63))
                                                                                                        case 4294967295:
                                                                                                            fmt.Println(v3, v94, fi15)
                                                                                                        default:
                                                                                                            v11 = v11
                                                                                                            v139 = 0
                                                                                                        }
                                                                                                        var v166 float64 = float64(0.0)
                                                                                                        _ = v166
                                                                                                    }
                                                                                                    v3 = (v3 + (v12 + ((v1 + v3) + v12)))
                                                                                                } else {
                                                                                                    var v167 uint32 = v139
                                                                                                    v139 = v167
                                                                                                    v139 += 65535
                                                                                                    _ = v167
                                                                                                }
                                                                                            }
                                                                                        case "hello":
                                                                                            v12 = "hello"
                                                                                        }
                                                                                        _ = v139
                                                                                    }
                                                                                    var v168 uint = 18446744073709551615
                                                                                    _ = v168
                                                                                }
                                                                            default:
                                                                                if v11 {
                                                                                    v169 := GenS17{F0: 2}
                                                                                    _ = len(v4)
                                                                                    v94 = 948
                                                                                    _ = v169
                                                                                } else {
                                                                                    v94 -= (v94 * (^(v94)))
                                                                                    switch 1 {
                                                                                    case 9223372036854775807:
                                                                                        var v170 uint64 = 27
                                                                                        _ = v170
                                                                                    }
                                                                                    v11 = (v11 && v11)
                                                                                }
                                                                            }
                                                                            var v171 int32 = genFn0(1, 9223372036854775807)
                                                                            _ = v171
                                                                        } else {
                                                                            v5 = v6
                                                                        }
                                                                        v172 := GenS8{F0: 0, F1: 2147483648}
                                                                        _ = v172
                                                                        _ = v94
                                                                    } else {
                                                                        fmt.Println(v2, v3, v5)
                                                                        if (!v11) {
                                                                            var v173 int64 = int64(v16)
                                                                            fmt.Println(v173)
                                                                            var v174 int64 = 862
                                                                            v2 -= (((v2 + v2) + ((3.14159 - float32(3.14159)) + 1.5)) / 1.5)
                                                                            _ = v173
                                                                            _ = v174
                                                                        }
                                                                        var v175 int = (fi15 + 2147483647)
                                                                        _ = v175
                                                                    }
                                                                default:
                                                                    v176 := GenS10{F0: 456, F1: 32767, F2: v16}
                                                                    _ = v176
                                                                }
                                                            } else {
                                                                var v177 uint32 = 1
                                                                var v178 float32 = (v2 + 0.0)
                                                                var v179 uint64 = 1
                                                                _ = v177
                                                                _ = v178
                                                                _ = v179
                                                            }
                                                            _ = v93
                                                        } else {
                                                            if len(v4) > 0 {
                                                                v4[0] = 0
                                                            }
                                                            v180 := GenS16{F0: 2, F1: 0, F2: (!(v11 || true))}
                                                            _ = v180
                                                        }
                                                    }
                                                }
                                            }
                                            v181 := GenS17{F0: 44}
                                            var v182 int = v16
                                            _ = v181
                                            _ = v182
                                        } else {
                                            switch fi14 {
                                            case -2147483648:
                                                fmt.Println(v1, v6)
                                            case 1:
                                                v4 = append(v4, 796)
                                            default:
                                                if false {
                                                    fmt.Println(v3, fi14)
                                                    var v183 uint32 = 0
                                                    _ = v183
                                                } else {
                                                    v184 := GenS18{F0: true}
                                                    _ = v184
                                                }
                                                if (-781 == 2147483647) {
                                                    v185 := GenS3{F0: 0}
                                                    v186 := GenS19{F0: genFn0(544, -772), F1: (v1 + "hello")}
                                                    _ = v185
                                                    _ = v186
                                                }
                                                v11 = v11
                                            }
                                        }
                                        var v187 bool = true
                                        _ = v187
                                    }
                                    switch v16 {
                                    case -9223372036854775808:
                                        fmt.Println(v5, v5, v11)
                                    case 2147483647:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    default:
                                        var v188 bool = false
                                        _ = v188
                                    }
                                    if (42 == 65535) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v189 uint32 = 2147483648
                                        v189 += v189
                                        _ = v189
                                    } else {
                                        switch v12 {
                                        case "hello":
                                            if (1e10 < float32(3.14159)) {
                                                v1 = ("a" + v1)
                                                var v190 int16 = 32767
                                                switch v1 {
                                                case "foo bar":
                                                    v6 -= v5
                                                default:
                                                    var v191 int16 = v190
                                                    _ = v191
                                                }
                                                _ = v190
                                            } else {
                                                v6 = v6
                                            }
                                        case "a":
                                            _ = len(v4)
                                            if len(v4) > 0 {
                                                v4[0] = 519
                                            }
                                        case "":
                                            switch v16 {
                                            case 0:
                                                v11 = v11
                                            case 32768:
                                                v192 := GenS18{F0: (v11 && (false || v11))}
                                                _ = v192
                                            }
                                        }
                                    }
                                    if ((((2 <= 849) || (426 != 65535)) || true) && (v11 && (0 > 247))) {
                                        var v193 int = v16
                                        if (false && v11) {
                                            if v11 {
                                                if (42 <= 36) {
                                                    var v194 int64 = -33
                                                    _ = v194
                                                }
                                            } else {
                                                v195 := GenS20{F0: v5, F1: 0, F2: false}
                                                if (4294967295 > 4294967295) {
                                                    var v196 int32 = 347
                                                    _ = v196
                                                }
                                                if v11 {
                                                    var v197 uint64 = 414
                                                    if len(v4) > 0 {
                                                        v4[0] = 2147483648
                                                    }
                                                    _ = v197
                                                }
                                                _ = v195
                                            }
                                        } else {
                                            v3 = v3
                                            if len(v4) > 0 {
                                                v4[0] = 4294967295
                                            }
                                        }
                                        switch v1 {
                                        case "\n\t":
                                            v3 = "foo bar"
                                        default:
                                            if v11 {
                                                v12 = v12
                                            } else {
                                                if ((v6 & -128) <= 21) {
                                                    if (!true) {
                                                        fmt.Println(v12)
                                                        if v11 {
                                                            fmt.Println(fi15, v3, v1)
                                                            _ = len(v4)
                                                        } else {
                                                            var v198 uint16 = 101
                                                            var v199 int32 = 42
                                                            _ = v198
                                                            _ = v199
                                                        }
                                                    }
                                                    if v11 {
                                                        if (((187 > 1) || v11) && (170 != 255)) {
                                                            v11 = v11
                                                        }
                                                        _ = len(v4)
                                                    } else {
                                                        v1 = "hello"
                                                        if len(v4) > 0 {
                                                            v4[0] = 498
                                                        }
                                                        var v200 uint = 0
                                                        _ = v200
                                                    }
                                                } else {
                                                    _ = len(v4)
                                                }
                                                v201 := GenS21{F0: -32768, F1: v11}
                                                if len(v4) > 0 {
                                                    v4[0] = 832
                                                }
                                                _ = v201
                                            }
                                        }
                                        _ = v193
                                    } else {
                                        v2 += -2.25
                                    }
                                } else {
                                    v202 := GenS8{F0: 1, F1: 0}
                                    _ = v202
                                }
                                if true {
                                    _ = len(v4)
                                    v11 = (1 != 0)
                                    switch v3 {
                                    case "a":
                                        v5 = (v5 % ((v6) | 1))
                                        var v203 uint32 = 0
                                        _ = v203
                                    default:
                                        switch 275 {
                                        case -618:
                                            v16 = fi14
                                        case 9223372036854775807:
                                            switch v1 {
                                            case "\n\t":
                                                v204 := GenS22{F0: v16}
                                                _ = v204
                                            default:
                                                fmt.Println(v16, fi15, fi15)
                                            }
                                            var v205 uint = 60
                                            v206 := GenS23{F0: 0, F1: -1}
                                            var v207 int8 = (v6 & v5)
                                            _ = v205
                                            _ = v206
                                            _ = v207
                                        default:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        var v208 string = (v1 + ((v3 + v3) + v3))
                                        var v209 int8 = -108
                                        _ = v208
                                        _ = v209
                                    }
                                } else {
                                    v16 -= fi14
                                    switch "a" {
                                    case "":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    if (((0.1 * (-2.25 * 1e10)) / 1e10) > float64(v2)) {
                                        var v210 string = v3
                                        _ = v210
                                    }
                                }
                                _ = v17
                            }
                        case -1:
                            switch v1 {
                            case "foo bar":
                                v2 = v2
                            case "":
                                v211 := GenS24{F0: fi14, F1: (0.0 < v2), F2: fi14}
                                _ = v211
                            default:
                                v212 := GenS8{F0: 255, F1: 620}
                                _ = v212
                            }
                        case -926:
                            _ = len(v4)
                        default:
                            v213 := GenS16{F0: 18446744073709551615, F1: 127, F2: false}
                            _ = v213
                        }
                        var v214 string = (v3 + (v12 + v1))
                        _ = fi15
                        _ = v16
                        _ = v214
                    }
                }
                for fi215 := 0; fi215 < 8; fi215++ {
                    fmt.Println(v3, v12, v3)
                    _ = fi215
                }
            } else {
                if true {
                    v2 += float32(float64((v2 + -2.25)))
                }
            }
            for fi216 := 0; fi216 < 14; fi216++ {
                if (!v11) {
                    if (255 >= 32767) {
                        v4 = append(v4, 65535)
                    } else {
                        v12 = ""
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v217 uint16 = 1
                        _ = v217
                    }
                    v11 = (782 == 2)
                    var v218 uint32 = 447
                    _ = v218
                }
                var v219 float32 = -2.25
                var v220 uint8 = 255
                _ = fi216
                _ = v219
                _ = v220
            }
            for fi221 := 0; fi221 < 16; fi221++ {
                if true {
                    var v222 uint32 = 65535
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v222
                } else {
                    v223 := GenS25{F0: v11, F1: v1}
                    if len(v4) > 0 {
                        v4[0] = 317
                    }
                    _ = v223
                }
                var v224 float32 = ((3.14159 * (-2.25 + 0.1)) + v2)
                if (42 > 2) {
                    v4 = append(v4, 4294967295)
                    v1 = v3
                    v12 = "hello"
                } else {
                    var v225 int8 = (v5 % (((v5 - (^(v6)))) | 1))
                    switch fi14 {
                    case -1:
                        var v226 string = v1
                        _ = v226
                    case -2147483648:
                        v225 = 100
                    default:
                        switch (fi14 - (fi14 * fi221)) {
                        case 127:
                            fmt.Println(fi221, v5, v224)
                        case 0:
                            v3 = v12
                        case 4294967295:
                            var v227 int64 = 1
                            _ = v227
                        default:
                            v6 = (v6 & -128)
                        }
                    }
                    v2 = float32(-2.25)
                    _ = v225
                }
                _ = fi221
                _ = v224
            }
            _ = fi14
        }
    } else {
        v3 = "a"
    }
    v12 = (v3 + v3)
    var v228 int = 2147483647
    if v11 {
        v229 := GenS26{F0: 255}
        for fi230 := 0; fi230 < 19; fi230++ {
            switch v12 {
            case "\n\t":
                if v11 {
                    if (((v1 + ((((v1 + "hello") + "a") + v1) + "a")) >= (v1 + ((v1 + "") + "hello"))) || v11) {
                        v11 = (255 <= 0)
                        switch v228 {
                        case -641:
                            v231 := GenS5{F0: (100 <= v5), F1: 42}
                            _ = v231
                        case -375:
                            if v11 {
                                if v11 {
                                    if v11 {
                                        for fi232 := 0; fi232 < 9; fi232++ {
                                            v233 := GenS27{F0: 1, F1: "\n\t"}
                                            _ = fi232
                                            _ = v233
                                        }
                                        for fi234 := 0; fi234 < 5; fi234++ {
                                            if ((65535 < 42) && v11) {
                                                fmt.Println(v12, v12)
                                                switch v1 {
                                                case "":
                                                    if v11 {
                                                        v235 := GenS25{F0: false, F1: ("" + ((("a" + (v3 + "\n\t")) + ((v3 + ((("\n\t" + (("hello" + (v1 + v12)) + (v12 + v12))) + ((v3 + "hello") + (v1 + ("" + v12)))) + ("a" + ("foo bar" + v3)))) + "hello")) + "a"))}
                                                        if v11 {
                                                            switch (^(fi230)) {
                                                            case 32768:
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            default:
                                                                v11 = (884 < 0)
                                                            }
                                                        }
                                                        _ = v235
                                                    }
                                                case "a":
                                                    fmt.Println(v5)
                                                default:
                                                    var v236 string = v1
                                                    _ = v236
                                                }
                                                v237 := GenS28{F0: 42, F1: v1}
                                                _ = v237
                                            }
                                            _ = fi234
                                        }
                                        v238 := GenS1{F0: v228, F1: v3}
                                        _ = v238
                                    } else {
                                        v5 = (v6 | (v6 >> ((221) & 63)))
                                        var v239 float32 = (v2 * (-2.25 + 1.5))
                                        var v240 float64 = (1e10 - ((1.5 - 1e10) - (0.1 + float64(v239))))
                                        _ = v239
                                        _ = v240
                                    }
                                } else {
                                    if true {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        fmt.Println(v228)
                                    } else {
                                        v241 := GenS29{F0: 104, F1: -1}
                                        switch fi230 {
                                        case 32768:
                                            fmt.Println(v11, v11)
                                        default:
                                            if false {
                                                switch v3 {
                                                case "a":
                                                    for fi242 := 0; fi242 < 13; fi242++ {
                                                        fmt.Println(v2, fi230, v6)
                                                        _ = fi242
                                                    }
                                                }
                                                v4 = append(v4, 892)
                                            }
                                        }
                                        if ((360 != -1) && v11) {
                                            if (-278 < (v228 << ((0) & 63))) {
                                                _ = len(v4)
                                            }
                                            fmt.Println(v12)
                                        } else {
                                            var v243 uint32 = 2147483647
                                            var v244 int8 = 1
                                            v12 = "a"
                                            _ = v243
                                            _ = v244
                                        }
                                        _ = v241
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = 42
                                    }
                                }
                                v245 := GenS30{F0: 0, F1: 255, F2: v6}
                                _ = v245
                            } else {
                                _ = len(v4)
                                for fi246 := 0; fi246 < 5; fi246++ {
                                    v4 = append(v4, 0)
                                    if v11 {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v6 = v6
                                        v247 := GenS31{F0: 42, F1: 65535}
                                        _ = v247
                                    }
                                    _ = fi246
                                }
                            }
                        case 128:
                            var v248 string = v1
                            _ = v248
                        }
                    } else {
                        var v249 string = v1
                        _ = v249
                    }
                    v1 = (("foo bar" + "foo bar") + (v1 + v1))
                    for fi250 := 0; fi250 < 3; fi250++ {
                        if (!true) {
                            v251 := GenS23{F0: 25, F1: genFn0(1, 575)}
                            v11 = (!false)
                            _ = v251
                        }
                        _ = fi250
                    }
                }
            default:
                v1 = "foo bar"
            }
            var v252 uint16 = 32767
            _ = fi230
            _ = v252
        }
        v12 = v3
        _ = v229
    }
    for fi253 := 0; fi253 < 19; fi253++ {
        switch 564 {
        case 42:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 4294967295:
            var v254 bool = true
            _ = v254
        case 1:
            fmt.Println(v1)
        default:
            for fi255 := 0; fi255 < 6; fi255++ {
                var v256 int = (-(fi253))
                _ = fi255
                _ = v256
            }
        }
        v2 -= v2
        _ = fi253
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v228:", v228)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v228
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
