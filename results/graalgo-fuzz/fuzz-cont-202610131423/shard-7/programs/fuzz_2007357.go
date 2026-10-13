package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int64 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int8; F1 int64; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS7 struct { F0 uint64; F1 uint8; F2 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint64; F1 int }
type GenS12 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS13 struct { F0 uint32; F1 int16 }
type GenS14 struct { F0 int64; F1 int32 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 uint8; F1 uint32 }
type GenS20 struct { F0 uint32; F1 int32; F2 string }
type GenS21 struct { F0 int8; F1 int }
type GenS22 struct { F0 uint16; F1 int32; F2 int64 }
type GenS23 struct { F0 uint; F1 uint8; F2 int64 }
func genFn0(p0 uint, p1 int) uint8 {
    for fi7 := 0; fi7 < 20; fi7++ {
        p0 += p0
        var v8 int16 = 556
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: 42, F2: 1}
    switch (("a" + "\n\t") + "\n\t") {
    case "\n\t":
        var v10 string = (("hello" + "foo bar") + ((("hello" + "foo bar") + ("a" + "a")) + ""))
        _ = v10
    }
    var v11 int64 = 127
    var v12 float32 = (((0.1 / 3.14159) - 1.5) / (1e10 / 1e10))
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v9
    return 127
}
func genFn1(p0 int32, p1 uint, p2 int32) int64 {
    v13 := []int{2147483647}
    v14 := GenS1{F0: -1, F1: 2147483648, F2: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 0
}

func main() {
    var v1 float64 = (((1e10 * (0.1 + (0.1 * -2.25))) * 0.1) + 3.14159)
    var v2 int32 = 42
    var v3 int8 = 0
    var v4 bool = (v1 == v1)
    var v5 uint8 = 0
    var v6 uint8 = (v5 + (^(v5)))
    var v15 float64 = (1.5 * 3.14159)
    v16 := []int{}
    for fi17 := 0; fi17 < 17; fi17++ {
        switch (fi17 ^ (^(fi17))) {
        case 4294967295:
            var v18 float32 = 0.1
            _ = v18
        case 442:
            for fi19 := 0; fi19 < 17; fi19++ {
                _ = len(v16)
                _ = fi19
            }
        case 2147483647:
            if len(v16) > 0 {
                v16[0] = (fi17 % (((fi17 - fi17)) | 1))
            }
        }
        var v20 int16 = -32768
        _ = fi17
        _ = v20
    }
    if len(v16) > 0 {
        v16[0] = 32767
    }
    var v21 bool = (!true)
    switch (("foo bar" + (("hello" + ("\n\t" + "hello")) + (("hello" + "hello") + "foo bar"))) + "hello") {
    case "\n\t":
        for fi22 := 0; fi22 < 1; fi22++ {
            if true {
                fmt.Println(fi22, v2)
                fmt.Println(v21, v21)
                for fi23 := 0; fi23 < 18; fi23++ {
                    v16 = append(v16, -534)
                    v24 := GenS1{F0: -240, F1: 0, F2: -792}
                    if v4 {
                        fmt.Println(fi22)
                        switch "" {
                        case "a":
                            var v25 int64 = -9223372036854775808
                            _ = v25
                        case "foo bar":
                            var v26 int16 = 438
                            _ = v26
                        case "hello":
                            var v27 float64 = 0.1
                            _ = v27
                        default:
                            v28 := GenS1{F0: 0, F1: 1, F2: 9223372036854775807}
                            _ = v28
                        }
                        fmt.Println(v15, v4)
                        var v29 uint8 = (v5 >> ((2) & 63))
                        _ = v29
                    }
                    _ = fi23
                    _ = v24
                }
            } else {
                if v4 {
                    switch (-(fi22)) {
                    case 857:
                        v30 := GenS1{F0: 625, F1: 4294967295, F2: -9223372036854775808}
                        _ = v30
                    case 0:
                        if len(v16) > 0 {
                            v16[0] = fi22
                        }
                    case -754:
                        fmt.Println(v1, v2, fi22)
                    }
                }
                if (42 < 0) {
                    for fi31 := 0; fi31 < 6; fi31++ {
                        if v21 {
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                        } else {
                            var v32 float64 = v15
                            switch fi31 {
                            case -2147483648:
                                switch (fi22 / (((-(fi31))) | 1)) {
                                case 1:
                                    fmt.Println(fi22)
                                case 562:
                                    v15 = 0.1
                                case 966:
                                    v16 = append(v16, (fi31 >> ((859) & 63)))
                                default:
                                    var v33 int16 = 0
                                    _ = v33
                                }
                            case 128:
                                if (!(v21 || (682 > 801))) {
                                    fmt.Println(v15, v5, v32)
                                    if (!v21) {
                                        switch fi31 {
                                        case 0:
                                            if len(v16) > 0 {
                                                v16[0] = 585
                                            }
                                        case 4294967295:
                                            v4 = true
                                        default:
                                            fmt.Println(v32)
                                            v3 = int8(v2)
                                            var v34 string = ""
                                            _ = v34
                                        }
                                        v32 -= float64((1e10 * 1.5))
                                    } else {
                                        var v35 int8 = int8(v2)
                                        var v36 int16 = 0
                                        _ = len(v16)
                                        _ = v35
                                        _ = v36
                                    }
                                } else {
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                }
                            case 9223372036854775807:
                                var v37 float64 = v32
                                _ = v37
                            }
                            v6 = (v5 * (v5 ^ 255))
                            if (0 >= 1) {
                                v38 := GenS1{F0: 0, F1: 1, F2: -1}
                                if (v21 && ((("" + "a") + ("a" + "")) <= "\n\t")) {
                                    var v39 int16 = -921
                                    var v40 int8 = (v3 * (v3 - v3))
                                    switch (fi31 | fi22) {
                                    case 1:
                                        v39 = v39
                                    case 0:
                                        switch "foo bar" {
                                        case "hello":
                                            if (!(255 > 18446744073709551615)) {
                                                _ = len(v16)
                                            }
                                        case "\n\t":
                                            v1 = 1.5
                                            v41 := GenS1{F0: 42, F1: 1, F2: int64(fi31)}
                                            if (18446744073709551615 <= 65535) {
                                                var v42 uint32 = 374
                                                fmt.Println(v3, v5)
                                                _ = len(v16)
                                                if (!(100 <= v40)) {
                                                    var v43 float32 = (1e10 * ((0.1 - 0.1) - (-2.25 + (float32(1.5) + -2.25))))
                                                    switch (fi22 % (((-(fi22))) | 1)) {
                                                    case 0:
                                                        if (1 > 65535) {
                                                            v43 = (v43 + (v43 / v43))
                                                            var v44 uint = 326
                                                            _ = v44
                                                        } else {
                                                            if (-9223372036854775808 == 0) {
                                                                var v45 uint = 22
                                                                v46 := GenS2{F0: int16(v42), F1: 255}
                                                                _ = v45
                                                                _ = v46
                                                            } else {
                                                                v47 := GenS3{F0: (v40 >> ((18446744073709551615) & 63)), F1: 42, F2: genFn0(485, int(v39))}
                                                                v42 = v42
                                                                _ = v47
                                                            }
                                                        }
                                                    case 128:
                                                        fmt.Println(fi22, v15, v15)
                                                        if (65535 < 0) {
                                                            if (v3 > 0) {
                                                                if (v43 < v43) {
                                                                    if (!v21) {
                                                                        switch fi31 {
                                                                        case 128:
                                                                            switch (("hello" + "") + (("" + "a") + "\n\t")) {
                                                                            case "foo bar":
                                                                                if false {
                                                                                    switch (("foo bar" + "\n\t") + "hello") {
                                                                                    case "foo bar":
                                                                                        switch int(v42) {
                                                                                        case 911:
                                                                                            v48 := GenS4{F0: (v39 * (^(v39))), F1: genFn0(513, (fi31 % (((fi22 % ((fi31) | 1))) | 1)))}
                                                                                            _ = v48
                                                                                        case 419:
                                                                                            v4 = (637 <= 42)
                                                                                            if len(v16) > 0 {
                                                                                                v16[0] = (^(fi22))
                                                                                            }
                                                                                        case 981:
                                                                                            if ((-432 > int64(v42)) || false) {
                                                                                                _ = len(v16)
                                                                                            }
                                                                                        }
                                                                                    case "a":
                                                                                        var v49 uint = 0
                                                                                        _ = v49
                                                                                    default:
                                                                                        switch "\n\t" {
                                                                                        case "a":
                                                                                            v50 := GenS5{F0: 195, F1: 2}
                                                                                            _ = v50
                                                                                        case "\n\t":
                                                                                            if (v21 && false) {
                                                                                                var v51 string = "\n\t"
                                                                                                if ((-(v3)) >= (v3 - 38)) {
                                                                                                    var v52 uint8 = (v5 + (^(v5)))
                                                                                                    if false {
                                                                                                        fmt.Println(v21, v15, v51)
                                                                                                        switch (fi22 | fi31) {
                                                                                                        case -134:
                                                                                                            var v53 int64 = int64(fi31)
                                                                                                            v32 -= (v15 + v32)
                                                                                                            var v54 float32 = v43
                                                                                                            _ = v53
                                                                                                            _ = v54
                                                                                                        case 2147483647:
                                                                                                            if v21 {
                                                                                                                if true {
                                                                                                                    switch v51 {
                                                                                                                    case "\n\t":
                                                                                                                        _ = len(v16)
                                                                                                                    case "foo bar":
                                                                                                                        v55 := GenS6{F0: v42, F1: 2, F2: 4294967295}
                                                                                                                        _ = v55
                                                                                                                    case "":
                                                                                                                        fmt.Println(v32, v4, v21)
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    switch v51 {
                                                                                                                    case "\n\t":
                                                                                                                        v56 := GenS4{F0: (^(v39)), F1: v6}
                                                                                                                        _ = v56
                                                                                                                    case "foo bar":
                                                                                                                        fmt.Println(v32)
                                                                                                                    case "":
                                                                                                                        if ((fi31 + fi31) == int(v39)) {
                                                                                                                            if (!v21) {
                                                                                                                                fmt.Println(v21)
                                                                                                                            } else {
                                                                                                                                v57 := GenS7{F0: 1, F1: (v5 + v52), F2: (4294967295 < 4294967295)}
                                                                                                                                v43 = v43
                                                                                                                                if len(v16) > 0 {
                                                                                                                                    v16[0] = fi22
                                                                                                                                }
                                                                                                                                _ = v57
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            if (70 >= (v3 - 42)) {
                                                                                                                                var v58 uint = 18446744073709551615
                                                                                                                                switch -2147483648 {
                                                                                                                                case -2147483648:
                                                                                                                                    v59 := GenS3{F0: (v3 % ((100) | 1)), F1: 0, F2: v52}
                                                                                                                                    v4 = (!(false || (!v4)))
                                                                                                                                    switch (v51 + (v51 + v51)) {
                                                                                                                                    case "\n\t":
                                                                                                                                        if false {
                                                                                                                                            v2 += v2
                                                                                                                                            switch fi31 {
                                                                                                                                            case 9223372036854775807:
                                                                                                                                                var v60 uint32 = (v42 | uint32(v58))
                                                                                                                                                _ = v60
                                                                                                                                            case 0:
                                                                                                                                                if len(v16) >= 2 {
                                                                                                                                                    v16 = v16[1:len(v16)]
                                                                                                                                                }
                                                                                                                                            default:
                                                                                                                                                v42 = 42
                                                                                                                                            }
                                                                                                                                        } else {
                                                                                                                                            v42 = 0
                                                                                                                                            var v61 int = (-(fi22))
                                                                                                                                            _ = v61
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    _ = v59
                                                                                                                                case 42:
                                                                                                                                    v51 = "a"
                                                                                                                                }
                                                                                                                                var v62 uint32 = 65535
                                                                                                                                _ = v58
                                                                                                                                _ = v62
                                                                                                                            }
                                                                                                                            fmt.Println(v42, v40, v15)
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        v63 := GenS8{F0: 42}
                                                                                                                        _ = v63
                                                                                                                    }
                                                                                                                    if (!(-9223372036854775808 < 1)) {
                                                                                                                        fmt.Println(v43, v5, v15)
                                                                                                                        v2 += (v2 / ((v2) | 1))
                                                                                                                    } else {
                                                                                                                        v21 = (!(!(1 < 1)))
                                                                                                                    }
                                                                                                                    var v64 int8 = v40
                                                                                                                    _ = v64
                                                                                                                }
                                                                                                                v65 := GenS9{F0: ("foo bar" + v51)}
                                                                                                                _ = v65
                                                                                                            } else {
                                                                                                                switch (fi22 >> ((42) & 63)) {
                                                                                                                case 32768:
                                                                                                                    v66 := GenS10{F0: ((v52 * v52) > v6)}
                                                                                                                    var v67 int64 = int64(v39)
                                                                                                                    _ = v66
                                                                                                                    _ = v67
                                                                                                                case -617:
                                                                                                                    v1 += 3.14159
                                                                                                                default:
                                                                                                                    v16 = append(v16, (fi22 / ((fi31) | 1)))
                                                                                                                }
                                                                                                            }
                                                                                                        case -2147483648:
                                                                                                            var v68 uint = 620
                                                                                                            _ = v68
                                                                                                        default:
                                                                                                            v16 = append(v16, (fi31 ^ 32768))
                                                                                                        }
                                                                                                    } else {
                                                                                                        if ((v43 / 1e10) > v43) {
                                                                                                            var v69 uint32 = v42
                                                                                                            _ = len(v16)
                                                                                                            _ = v69
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v52
                                                                                                } else {
                                                                                                    if v4 {
                                                                                                        switch ("" + v51) {
                                                                                                        case "foo bar":
                                                                                                            switch (fi22 / (((fi31 & -902)) | 1)) {
                                                                                                            case -9223372036854775808:
                                                                                                                switch fi31 {
                                                                                                                case 2147483647:
                                                                                                                    v40 = int8(fi22)
                                                                                                                }
                                                                                                            case 42:
                                                                                                                var v70 int32 = -2147483648
                                                                                                                _ = v70
                                                                                                            }
                                                                                                        }
                                                                                                        switch (fi31 >> ((795) & 63)) {
                                                                                                        case 4294967295:
                                                                                                            v39 -= v39
                                                                                                            switch (("" + (v51 + v51)) + v51) {
                                                                                                            case "\n\t":
                                                                                                                switch (v51 + "a") {
                                                                                                                case "a":
                                                                                                                    v32 += 3.14159
                                                                                                                default:
                                                                                                                    v15 -= 0.1
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                if (53 < 105) {
                                                                                                                    var v71 int8 = v40
                                                                                                                    var v72 uint8 = 55
                                                                                                                    var v73 float64 = v15
                                                                                                                    _ = v71
                                                                                                                    _ = v72
                                                                                                                    _ = v73
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        v6 -= (v5 ^ (v5 | genFn0(2, 32768)))
                                                                                                    } else {
                                                                                                        var v74 uint8 = 128
                                                                                                        if false {
                                                                                                            var v75 uint = 182
                                                                                                            if len(v16) >= 2 {
                                                                                                                v16 = v16[1:len(v16)]
                                                                                                            }
                                                                                                            if v4 {
                                                                                                                fmt.Println(v75, v40, v3)
                                                                                                                if (genFn0(65535, (fi22 / (((fi31 ^ -198)) | 1))) >= genFn0(589, fi31)) {
                                                                                                                    v4 = (v3 == (v40 ^ v40))
                                                                                                                    v74 = v6
                                                                                                                    v16 = append(v16, 4294967295)
                                                                                                                } else {
                                                                                                                    var v76 uint16 = 42
                                                                                                                    v6 = genFn0(v75, fi22)
                                                                                                                    _ = v76
                                                                                                                }
                                                                                                                if ((32767 >= 2) || v21) {
                                                                                                                    var v77 int8 = (v3 ^ v40)
                                                                                                                    if ((^(v6)) <= v6) {
                                                                                                                        var v78 int = (fi31 / (((fi22 ^ 724)) | 1))
                                                                                                                        _ = v78
                                                                                                                    }
                                                                                                                    _ = v77
                                                                                                                }
                                                                                                            } else {
                                                                                                                v40 -= (v3 % ((v3) | 1))
                                                                                                                if (-423 < 1) {
                                                                                                                    if (false && v21) {
                                                                                                                        v4 = ((0 == -1) || ((v39 > v39) && ((((v1 * v1) / v1) >= v15) && v21)))
                                                                                                                    }
                                                                                                                    var v79 float32 = (3.14159 - v43)
                                                                                                                    v80 := GenS11{F0: 99, F1: int(v5)}
                                                                                                                    _ = v79
                                                                                                                    _ = v80
                                                                                                                }
                                                                                                                fmt.Println(v3, v6, v3)
                                                                                                                if v4 {
                                                                                                                    v81 := GenS12{F0: 1, F1: 2, F2: (-(v39))}
                                                                                                                    _ = v81
                                                                                                                } else {
                                                                                                                    fmt.Println(v51, v3, v6)
                                                                                                                    v82 := GenS13{F0: v42, F1: 0}
                                                                                                                    var v83 float64 = v1
                                                                                                                    if (!v21) {
                                                                                                                        _ = len(v16)
                                                                                                                        var v84 int8 = v3
                                                                                                                        _ = v84
                                                                                                                    } else {
                                                                                                                        var v85 bool = false
                                                                                                                        v6 -= 128
                                                                                                                        v42 -= (v42 - (v42 << ((975) & 63)))
                                                                                                                        _ = v85
                                                                                                                    }
                                                                                                                    var v86 bool = (2 <= 42)
                                                                                                                    _ = v82
                                                                                                                    _ = v83
                                                                                                                    _ = v86
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v75
                                                                                                        }
                                                                                                        v2 = v2
                                                                                                        _ = v74
                                                                                                    }
                                                                                                }
                                                                                                _ = v51
                                                                                            }
                                                                                        default:
                                                                                            v40 = v3
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    switch (fi31 / ((fi31) | 1)) {
                                                                                    case 32767:
                                                                                        if false {
                                                                                            if v21 {
                                                                                                v3 = v3
                                                                                                v16 = append(v16, fi31)
                                                                                                v39 = (v39 | (v39 | (v39 << ((18446744073709551615) & 63))))
                                                                                                switch "" {
                                                                                                case "":
                                                                                                    switch fi31 {
                                                                                                    case 0:
                                                                                                        if len(v16) >= 2 {
                                                                                                            v16 = v16[1:len(v16)]
                                                                                                        }
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    v87 := GenS11{F0: 2, F1: fi22}
                                                                                                    _ = v87
                                                                                                default:
                                                                                                    fmt.Println(v2, v6, v43)
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            if ("a" != "hello") {
                                                                                                var v88 uint64 = 146
                                                                                                _ = v88
                                                                                            } else {
                                                                                                _ = len(v16)
                                                                                                if (!false) {
                                                                                                    v2 = v2
                                                                                                }
                                                                                                switch int(v40) {
                                                                                                case -9223372036854775808:
                                                                                                    v32 = (3.14159 * v1)
                                                                                                case -2147483648:
                                                                                                    fmt.Println(v4, v43)
                                                                                                    fmt.Println(v4, v21)
                                                                                                    if v21 {
                                                                                                        if v21 {
                                                                                                            v89 := GenS14{F0: 0, F1: 94}
                                                                                                            if v4 {
                                                                                                                var v90 float32 = 1e10
                                                                                                                _ = v90
                                                                                                            } else {
                                                                                                                switch fi22 {
                                                                                                                case -9223372036854775808:
                                                                                                                    v91 := GenS6{F0: v42, F1: 18446744073709551615, F2: 2}
                                                                                                                    _ = v91
                                                                                                                case -320:
                                                                                                                    fmt.Println(v5, v15, v1)
                                                                                                                case 42:
                                                                                                                    if v4 {
                                                                                                                        var v92 string = "\n\t"
                                                                                                                        var v93 float64 = float64(v43)
                                                                                                                        switch fi22 {
                                                                                                                        case 966:
                                                                                                                            v94 := GenS10{F0: ((v3 * 1) > v3)}
                                                                                                                            _ = v94
                                                                                                                        case 42:
                                                                                                                            v5 -= genFn0(455, -693)
                                                                                                                        default:
                                                                                                                            v5 += (^(v5))
                                                                                                                        }
                                                                                                                        if v21 {
                                                                                                                            var v95 bool = (!v21)
                                                                                                                            _ = v95
                                                                                                                        }
                                                                                                                        var v96 uint64 = 1
                                                                                                                        fmt.Println(v21, fi22, v93)
                                                                                                                        if ((v40 <= v3) && v4) {
                                                                                                                            var v97 int = (fi31 + fi22)
                                                                                                                            v97 -= (-(v97))
                                                                                                                            _ = v97
                                                                                                                        } else {
                                                                                                                            if false {
                                                                                                                                var v98 uint64 = 882
                                                                                                                                v15 = (v32 - ((((0.0 / v93) / (1.5 / v32)) * ((v15 / -2.25) + (v93 - (v93 / v1)))) / v32))
                                                                                                                                _ = v98
                                                                                                                            } else {
                                                                                                                                fmt.Println(v4, v21)
                                                                                                                                v93 += (3.14159 - (v15 * (1.5 - float64(0.1))))
                                                                                                                                var v99 int16 = v39
                                                                                                                                _ = v99
                                                                                                                            }
                                                                                                                            v39 += (v39 ^ (v39 - v39))
                                                                                                                            switch (fi31 | 4294967295) {
                                                                                                                            case 127:
                                                                                                                                v43 += v43
                                                                                                                            default:
                                                                                                                                if (18446744073709551615 == 65535) {
                                                                                                                                    v16 = append(v16, fi22)
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v93, v42, fi22)
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                        _ = v92
                                                                                                                        _ = v93
                                                                                                                        _ = v96
                                                                                                                    } else {
                                                                                                                        switch "\n\t" {
                                                                                                                        case "\n\t":
                                                                                                                            if (false && ((true && (4294967295 == 255)) || (!v4))) {
                                                                                                                                var v100 int32 = -1
                                                                                                                                v101 := GenS15{F0: (v39 | (-(v39)))}
                                                                                                                                var v102 float32 = v43
                                                                                                                                _ = v100
                                                                                                                                _ = v101
                                                                                                                                _ = v102
                                                                                                                            }
                                                                                                                        case "hello":
                                                                                                                            if ((^(fi22)) > (fi22 | (fi31 >> ((2) & 63)))) {
                                                                                                                                switch (fi31 / (((fi31 & (fi31 ^ fi31))) | 1)) {
                                                                                                                                case 32767:
                                                                                                                                    if v21 {
                                                                                                                                        var v103 uint16 = 32767
                                                                                                                                        v104 := GenS16{F0: 4294967295}
                                                                                                                                        _ = v103
                                                                                                                                        _ = v104
                                                                                                                                    }
                                                                                                                                case -589:
                                                                                                                                    var v105 int8 = (v3 & v3)
                                                                                                                                    _ = v105
                                                                                                                                case -79:
                                                                                                                                    switch "\n\t" {
                                                                                                                                    case "hello":
                                                                                                                                        var v106 string = "foo bar"
                                                                                                                                        _ = v106
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    var v107 uint8 = 42
                                                                                                                                    _ = v107
                                                                                                                                }
                                                                                                                            }
                                                                                                                        case "foo bar":
                                                                                                                            var v108 bool = false
                                                                                                                            _ = v108
                                                                                                                        default:
                                                                                                                            v15 = (v32 + v1)
                                                                                                                        }
                                                                                                                    }
                                                                                                                default:
                                                                                                                    v42 = (^(v42))
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v89
                                                                                                        }
                                                                                                        var v109 int = 821
                                                                                                        _ = len(v16)
                                                                                                        var v110 bool = v21
                                                                                                        _ = v109
                                                                                                        _ = v110
                                                                                                    }
                                                                                                case 183:
                                                                                                    if v21 {
                                                                                                        v5 += v6
                                                                                                        var v111 int64 = 756
                                                                                                        var v112 uint = 4294967295
                                                                                                        var v113 int = fi31
                                                                                                        _ = v111
                                                                                                        _ = v112
                                                                                                        _ = v113
                                                                                                    }
                                                                                                }
                                                                                                var v114 uint8 = v6
                                                                                                if len(v16) > 0 {
                                                                                                    v16[0] = -807
                                                                                                }
                                                                                                _ = v114
                                                                                            }
                                                                                        }
                                                                                    case 4294967295:
                                                                                        var v115 uint32 = uint32(v5)
                                                                                        _ = v115
                                                                                    default:
                                                                                        var v116 int64 = 1
                                                                                        _ = v116
                                                                                    }
                                                                                    if (!(!(602 > (v2 & (v2 ^ (v2 * v2)))))) {
                                                                                        v117 := GenS7{F0: 1, F1: (v5 >> ((65535) & 63)), F2: true}
                                                                                        v15 = 0.0
                                                                                        _ = v117
                                                                                    }
                                                                                    v5 = v5
                                                                                    var v118 uint32 = (v42 | 955)
                                                                                    _ = v118
                                                                                }
                                                                            case "hello":
                                                                                switch "foo bar" {
                                                                                case "":
                                                                                    switch "" {
                                                                                    case "hello":
                                                                                        var v119 uint32 = v42
                                                                                        _ = v119
                                                                                    case "foo bar":
                                                                                        fmt.Println(v6, v15, v32)
                                                                                        v120 := GenS8{F0: 32767}
                                                                                        var v121 int8 = -1
                                                                                        var v122 int = (fi31 ^ fi22)
                                                                                        v16 = append(v16, (fi31 << ((255) & 63)))
                                                                                        v3 = (-(v121))
                                                                                        var v123 uint64 = 622
                                                                                        fmt.Println(v5, v1, v43)
                                                                                        var v124 string = ("foo bar" + "")
                                                                                        v125 := GenS17{F0: (v6 + v5)}
                                                                                        _ = v120
                                                                                        _ = v121
                                                                                        _ = v122
                                                                                        _ = v123
                                                                                        _ = v124
                                                                                        _ = v125
                                                                                    }
                                                                                case "foo bar":
                                                                                    v16 = append(v16, (fi22 + (-(fi22))))
                                                                                }
                                                                                var v126 float64 = 1e10
                                                                                v127 := GenS18{F0: 146}
                                                                                switch "hello" {
                                                                                case "":
                                                                                    var v128 uint8 = (^(v5))
                                                                                    _ = v128
                                                                                case "\n\t":
                                                                                    fmt.Println(v3, v2)
                                                                                }
                                                                                switch (fi31 >> ((182) & 63)) {
                                                                                case 9223372036854775807:
                                                                                    fmt.Println(v126, v15, fi31)
                                                                                }
                                                                                _ = v126
                                                                                _ = v127
                                                                            default:
                                                                                v42 = v42
                                                                            }
                                                                        case 167:
                                                                            v42 = (^(v42))
                                                                        case 32768:
                                                                            if v21 {
                                                                                if len(v16) >= 2 {
                                                                                    v16 = v16[1:len(v16)]
                                                                                }
                                                                                switch (fi31 + fi22) {
                                                                                case 9223372036854775807:
                                                                                    switch ("" + "") {
                                                                                    case "foo bar":
                                                                                        v5 += v5
                                                                                    case "hello":
                                                                                        v40 = (v40 - (v40 * (v40 - (v3 >> ((156) & 63)))))
                                                                                    case "a":
                                                                                        v15 -= (v32 + v15)
                                                                                    default:
                                                                                        if (!(!v21)) {
                                                                                            v129 := GenS3{F0: 42, F1: 317, F2: v5}
                                                                                            v130 := GenS1{F0: v39, F1: 0, F2: int64(v2)}
                                                                                            _ = v129
                                                                                            _ = v130
                                                                                        } else {
                                                                                            if ((1 < (v5 & v6)) || ((v2 >> ((773) & 63)) == 942)) {
                                                                                                switch (-(fi31)) {
                                                                                                case 42:
                                                                                                    var v131 int32 = (v2 / (((-(v2))) | 1))
                                                                                                    _ = v131
                                                                                                case 0:
                                                                                                    switch fi31 {
                                                                                                    case 931:
                                                                                                        if len(v16) > 0 {
                                                                                                            v16[0] = fi22
                                                                                                        }
                                                                                                    }
                                                                                                    if v4 {
                                                                                                        v132 := GenS9{F0: "\n\t"}
                                                                                                        _ = v132
                                                                                                    } else {
                                                                                                        var v133 float64 = (float64((((1e10 - -2.25) * v43) + 3.14159)) * v32)
                                                                                                        _ = v133
                                                                                                    }
                                                                                                    var v134 bool = v21
                                                                                                    _ = v134
                                                                                                }
                                                                                                var v135 int64 = 42
                                                                                                _ = v135
                                                                                            } else {
                                                                                                fmt.Println(v21, fi31)
                                                                                                if len(v16) > 0 {
                                                                                                    v16[0] = (fi31 / (((fi31 << ((408) & 63))) | 1))
                                                                                                }
                                                                                            }
                                                                                            switch (fi31 << ((444) & 63)) {
                                                                                            case 1:
                                                                                                var v136 uint = 4294967295
                                                                                                _ = v136
                                                                                            case 128:
                                                                                                v6 -= genFn0(0, -1)
                                                                                            case 9223372036854775807:
                                                                                                if (v21 && true) {
                                                                                                    var v137 uint = 4294967295
                                                                                                    _ = v137
                                                                                                }
                                                                                            }
                                                                                            switch 2147483647 {
                                                                                            case 0:
                                                                                                if v4 {
                                                                                                    fmt.Println(fi22, fi31, v32)
                                                                                                } else {
                                                                                                    v40 += 127
                                                                                                    v138 := GenS7{F0: 42, F1: v6, F2: v21}
                                                                                                    _ = v138
                                                                                                }
                                                                                            case 32768:
                                                                                                v16 = append(v16, (^(fi22)))
                                                                                            default:
                                                                                                fmt.Println(v40)
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                case 128:
                                                                                    fmt.Println(v3, v3)
                                                                                case 32767:
                                                                                    v139 := GenS9{F0: ""}
                                                                                    v32 += (0.1 + v15)
                                                                                    var v140 int8 = (v40 << ((210) & 63))
                                                                                    _ = v139
                                                                                    _ = v140
                                                                                default:
                                                                                    v2 = 0
                                                                                }
                                                                                v141 := GenS13{F0: (v42 ^ (^(v42))), F1: v39}
                                                                                _ = v141
                                                                            }
                                                                        default:
                                                                            var v142 int = fi31
                                                                            _ = v142
                                                                        }
                                                                        var v143 string = ""
                                                                        fmt.Println(v3)
                                                                        _ = v143
                                                                    } else {
                                                                        fmt.Println(v1, v21, v3)
                                                                        v16 = append(v16, (fi22 * 816))
                                                                    }
                                                                    var v144 uint16 = 0
                                                                    _ = v144
                                                                } else {
                                                                    var v145 int32 = (-(v2))
                                                                    if v4 {
                                                                        var v146 bool = v4
                                                                        switch fi31 {
                                                                        case 128:
                                                                            fmt.Println(v4, fi22)
                                                                        case 42:
                                                                            v39 = v39
                                                                        }
                                                                        _ = v146
                                                                    } else {
                                                                        switch ("a" + ("hello" + ("hello" + "hello"))) {
                                                                        case "\n\t":
                                                                            var v147 int64 = -9223372036854775808
                                                                            _ = v147
                                                                        case "a":
                                                                            v148 := GenS2{F0: v39, F1: 2}
                                                                            _ = v148
                                                                        case "hello":
                                                                            v40 = v3
                                                                        default:
                                                                            fmt.Println(fi22)
                                                                        }
                                                                        v149 := GenS18{F0: -889}
                                                                        _ = v149
                                                                    }
                                                                    v145 -= int32(v5)
                                                                    _ = v145
                                                                }
                                                                var v150 uint64 = 4294967295
                                                                var v151 uint64 = (v150 + 392)
                                                                var v152 int32 = 42
                                                                v150 = 18446744073709551615
                                                                v152 = (v152 >> ((18446744073709551615) & 63))
                                                                _ = v150
                                                                _ = v151
                                                                _ = v152
                                                            } else {
                                                                v16 = append(v16, (fi22 >> ((65535) & 63)))
                                                                v16 = append(v16, (fi22 >> ((65535) & 63)))
                                                            }
                                                        }
                                                        if v21 {
                                                            if len(v16) >= 2 {
                                                                v16 = v16[1:len(v16)]
                                                            }
                                                            var v153 bool = true
                                                            _ = v153
                                                        }
                                                    case 1:
                                                        fmt.Println(v1)
                                                    default:
                                                        v154 := GenS19{F0: 117, F1: v42}
                                                        _ = v154
                                                    }
                                                    if v4 {
                                                        v39 += -302
                                                        v32 = float64(v43)
                                                        switch (fi31 / (((-(fi31))) | 1)) {
                                                        case 0:
                                                            switch (fi31 + int(v42)) {
                                                            case 769:
                                                                fmt.Println(fi22, v40)
                                                            case 2147483647:
                                                                switch (("a" + (("foo bar" + "a") + "foo bar")) + ("foo bar" + "a")) {
                                                                case "\n\t":
                                                                    v43 -= float32(v1)
                                                                default:
                                                                    v40 -= (-(v40))
                                                                    var v155 bool = false
                                                                    if len(v16) > 0 {
                                                                        v16[0] = fi22
                                                                    }
                                                                    _ = v155
                                                                }
                                                            case 4:
                                                                if len(v16) >= 2 {
                                                                    v16 = v16[1:len(v16)]
                                                                }
                                                            }
                                                        case 2147483647:
                                                            if len(v16) >= 2 {
                                                                v16 = v16[1:len(v16)]
                                                            }
                                                        case -947:
                                                            if len(v16) > 0 {
                                                                v16[0] = -1
                                                            }
                                                        default:
                                                            var v156 int16 = (v39 - v39)
                                                            _ = v156
                                                        }
                                                        v157 := GenS19{F0: (v5 % ((1) | 1)), F1: v42}
                                                        fmt.Println(v6, fi22, v1)
                                                        _ = v157
                                                    } else {
                                                        fmt.Println(v21)
                                                        v43 += (v43 + (v43 - float32(v15)))
                                                        v158 := GenS5{F0: v2, F1: 1}
                                                        var v159 float64 = v1
                                                        v16 = append(v16, (fi22 % (((fi31 * fi22)) | 1)))
                                                        _ = v158
                                                        _ = v159
                                                    }
                                                    _ = v43
                                                } else {
                                                    v6 -= 42
                                                    if v4 {
                                                        if v21 {
                                                            v15 -= v1
                                                            var v160 int16 = (v39 % ((int16(fi22)) | 1))
                                                            _ = v160
                                                        }
                                                        var v161 uint = 4294967295
                                                        _ = v161
                                                    } else {
                                                        if (v21 && false) {
                                                            if v4 {
                                                                v32 -= (float64(0.1) * (v1 - float64(-2.25)))
                                                                switch "a" {
                                                                case "":
                                                                    v162 := GenS15{F0: v39}
                                                                    _ = v162
                                                                case "foo bar":
                                                                    v1 += v32
                                                                }
                                                                fmt.Println(v3, v3)
                                                            } else {
                                                                switch fi31 {
                                                                case 42:
                                                                    if len(v16) >= 2 {
                                                                        v16 = v16[1:len(v16)]
                                                                    }
                                                                }
                                                            }
                                                            if (351 < 18446744073709551615) {
                                                                var v163 float64 = 0.0
                                                                _ = v163
                                                            }
                                                            fmt.Println(v21, v6, v6)
                                                        }
                                                        var v164 int8 = 1
                                                        fmt.Println(v164, v2, v4)
                                                        _ = v164
                                                    }
                                                }
                                                _ = v42
                                            } else {
                                                _ = len(v16)
                                                switch "hello" {
                                                case "\n\t":
                                                    if v21 {
                                                        v2 = 2147483647
                                                    } else {
                                                        v6 -= v5
                                                    }
                                                default:
                                                    v3 = v40
                                                }
                                            }
                                            _ = v41
                                        case "":
                                            fmt.Println(v6)
                                        default:
                                            if true {
                                                if (v40 == (-(v3))) {
                                                    v4 = (v4 || ((v21 && ((v6 <= v5) || (966 != -650))) || (((3.14159 - 1.5) + (3.14159 - (1e10 * (((0.0 * float32((-2.25 / (float64(3.14159) / 1e10)))) * -2.25) - (float32((((v15 * v32) / 1.0) - v1)) - 1.5))))) > float32((-2.25 + v15)))))
                                                } else {
                                                    v165 := GenS17{F0: (v6 / ((genFn0(uint(v6), (-(fi31)))) | 1))}
                                                    _ = v165
                                                }
                                                var v166 int32 = v2
                                                v16 = append(v16, (fi31 & (-(fi31))))
                                                var v167 int32 = v2
                                                v168 := GenS3{F0: (v3 | (v40 - 100)), F1: -1, F2: genFn0(18446744073709551615, 2147483647)}
                                                if ((v4 && false) || (!(1 == 998))) {
                                                    var v169 uint16 = 265
                                                    v40 += (-(v40))
                                                    _ = v169
                                                } else {
                                                    var v170 uint = 101
                                                    var v171 uint = v170
                                                    if (0 > 117) {
                                                        v32 -= 0.1
                                                    }
                                                    var v172 uint = v170
                                                    _ = v170
                                                    _ = v171
                                                    _ = v172
                                                }
                                                v166 -= v166
                                                _ = v166
                                                _ = v167
                                                _ = v168
                                            }
                                        }
                                    default:
                                        if (32767 <= 42) {
                                            v2 += (v2 + v2)
                                            v6 = genFn0(4294967295, (fi22 >> ((65535) & 63)))
                                        } else {
                                            if len(v16) > 0 {
                                                v16[0] = fi22
                                            }
                                            var v173 uint64 = 1
                                            if len(v16) > 0 {
                                                v16[0] = (fi31 & 128)
                                            }
                                            v2 = 42
                                            _ = v173
                                        }
                                    }
                                    var v174 uint32 = 4294967295
                                    _ = v174
                                    _ = v39
                                    _ = v40
                                } else {
                                    _ = len(v16)
                                    v21 = (788 <= 42)
                                }
                                var v175 int8 = v3
                                _ = v175
                                _ = v38
                            } else {
                                fmt.Println(v15, v32, fi22)
                            }
                            var v176 int64 = 9223372036854775807
                            v32 += float64(1e10)
                            _ = v176
                            _ = v32
                        }
                        fmt.Println(v5, v2, v2)
                        if v4 {
                            var v177 uint16 = 981
                            _ = v177
                        } else {
                            v178 := GenS20{F0: 2, F1: (v2 >> ((255) & 63)), F2: ("hello" + ((("a" + "\n\t") + ("a" + ("a" + "\n\t"))) + "foo bar"))}
                            var v179 uint8 = v6
                            if v4 {
                                switch fi22 {
                                case 329:
                                    v15 += v15
                                case 128:
                                    var v180 int8 = v3
                                    _ = v180
                                }
                                fmt.Println(v2, fi31)
                            } else {
                                if (!(v21 && v4)) {
                                    v2 += (v2 / (((-(v2))) | 1))
                                    if v4 {
                                        var v181 bool = true
                                        _ = len(v16)
                                        _ = v181
                                    }
                                } else {
                                    v4 = v21
                                }
                                if v21 {
                                    _ = len(v16)
                                    var v182 bool = false
                                    v183 := GenS4{F0: -32768, F1: genFn0(255, (fi31 & (fi31 | 2147483647)))}
                                    if false {
                                        var v184 int32 = v2
                                        v184 -= 663
                                        v3 += 100
                                        _ = v184
                                    } else {
                                        v16 = append(v16, fi31)
                                    }
                                    if (v2 < -1) {
                                        if v4 {
                                            v16 = append(v16, fi31)
                                            _ = len(v16)
                                            v4 = (v1 > (1.5 / 1e10))
                                        } else {
                                            var v185 uint32 = 317
                                            fmt.Println(v4, v21, v1)
                                            if v182 {
                                                switch "" {
                                                case "hello":
                                                    var v186 bool = v182
                                                    _ = v186
                                                default:
                                                    v187 := GenS21{F0: (^(v3)), F1: (fi31 + -1)}
                                                    _ = v187
                                                }
                                                var v188 string = ("\n\t" + "")
                                                _ = v188
                                            }
                                            _ = v185
                                        }
                                        v189 := GenS3{F0: 42, F1: -372, F2: (v179 + v179)}
                                        _ = v189
                                    } else {
                                        v1 = 3.14159
                                        if v4 {
                                            fmt.Println(v3)
                                        }
                                        v4 = true
                                        v190 := GenS9{F0: ""}
                                        _ = v190
                                    }
                                    _ = v182
                                    _ = v183
                                }
                            }
                            _ = v178
                            _ = v179
                        }
                        _ = fi31
                    }
                } else {
                    var v191 int = fi22
                    _ = v191
                }
            }
            _ = fi22
        }
    case "a":
        if v4 {
            var v192 string = "hello"
            if v21 {
                switch 192 {
                case 128:
                    for fi193 := 0; fi193 < 4; fi193++ {
                        var v194 bool = false
                        fmt.Println(v4, v21)
                        var v195 int16 = 32767
                        _ = fi193
                        _ = v194
                        _ = v195
                    }
                }
                fmt.Println(v6, v5, v192)
            } else {
                v192 = v192
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
            }
            if len(v16) > 0 {
                v16[0] = 127
            }
            _ = v192
        } else {
            var v196 uint8 = genFn0(1, 1)
            _ = v196
        }
    }
    switch -50 {
    case 1:
        v1 -= v15
    }
    if ((!v4) || (true || true)) {
        var v197 float32 = 0.1
        switch 127 {
        case 127:
            var v198 uint16 = 32767
            _ = v198
        case -2147483648:
            fmt.Println(v2, v5)
        }
        v199 := GenS9{F0: "hello"}
        _ = v197
        _ = v199
    } else {
        v200 := GenS22{F0: 495, F1: v2, F2: 9223372036854775807}
        switch "foo bar" {
        case "hello":
            var v201 uint = 289
            _ = v201
        default:
            fmt.Println(v2)
        }
        for fi202 := 0; fi202 < 5; fi202++ {
            var v203 float32 = float32(1.5)
            _ = fi202
            _ = v203
        }
        _ = v200
    }
    if len(v16) > 0 {
        v16[0] = -532
    }
    v2 = v2
    switch -756 {
    case -9223372036854775808:
        for fi204 := 0; fi204 < 17; fi204++ {
            switch "foo bar" {
            case "a":
                var v205 int8 = (v3 + (v3 / ((v3) | 1)))
                _ = v205
            case "hello":
                if (true && v21) {
                    var v206 uint = 4294967295
                    var v207 string = ("" + "foo bar")
                    _ = v206
                    _ = v207
                } else {
                    var v208 uint64 = 4294967295
                    _ = v208
                }
            case "":
                switch "a" {
                case "foo bar":
                    v209 := GenS12{F0: 32767, F1: 433, F2: 185}
                    _ = v209
                case "hello":
                    var v210 int = -2147483648
                    _ = v210
                default:
                    fmt.Println(v1, v6)
                }
            default:
                if (!v4) {
                    if (!v21) {
                        if ((true && (v4 && (0 < 889))) && (830 > 18446744073709551615)) {
                            if (v4 && (v4 || v21)) {
                                v16 = append(v16, (fi204 ^ fi204))
                                fmt.Println(v2)
                                for fi211 := 0; fi211 < 13; fi211++ {
                                    v212 := GenS5{F0: 42, F1: 65535}
                                    _ = fi211
                                    _ = v212
                                }
                            }
                            v6 += (v5 & v5)
                        } else {
                            for fi213 := 0; fi213 < 5; fi213++ {
                                fmt.Println(v5, fi204, v1)
                                v5 = (^(v5))
                                _ = fi213
                            }
                            switch 32767 {
                            case -612:
                                v214 := GenS10{F0: (-2.25 != ((((float32(v1) / (0.1 - ((0.1 * float32(1.5)) * 1e10))) - (((0.1 + -2.25) / 1e10) * (3.14159 - ((((-2.25 / ((0.0 - (1.5 + 3.14159)) * 0.1)) * -2.25) * (((0.0 * 0.0) + 0.0) / -2.25)) / 0.1)))) * float32((v15 / 1.0))) + (3.14159 + 0.1)))}
                                _ = v214
                            case 4294967295:
                                v4 = v4
                            case 32767:
                                var v215 int64 = 42
                                _ = v215
                            default:
                                var v216 uint16 = 1
                                _ = v216
                            }
                        }
                        for fi217 := 0; fi217 < 5; fi217++ {
                            var v218 int8 = (-(v3))
                            v219 := GenS10{F0: false}
                            var v220 int = (fi217 << ((4294967295) & 63))
                            _ = fi217
                            _ = v218
                            _ = v219
                            _ = v220
                        }
                        var v221 int16 = 32767
                        _ = v221
                    }
                    v16 = append(v16, fi204)
                    var v222 uint = 317
                    _ = v222
                } else {
                    var v223 uint = 4294967295
                    _ = v223
                }
            }
            _ = fi204
        }
    default:
        v5 = genFn0(42, -2147483648)
    }
    v224 := GenS23{F0: 255, F1: (v6 % ((v6) | 1)), F2: 0}
    _ = len(v16)
    if (!(406 == 1)) {
        var v225 string = ""
        _ = v225
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v224
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
