package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 int32 }
type GenS3 struct { F0 int; F1 uint64; F2 int64 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int; F1 uint32; F2 int }
type GenS10 struct { F0 uint8; F1 int }
type GenS11 struct { F0 bool; F1 uint16 }
type GenS12 struct { F0 int; F1 uint }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint32; F1 int16 }
type GenS16 struct { F0 uint16; F1 int32; F2 int32 }
type GenS17 struct { F0 uint64; F1 uint16 }
type GenS18 struct { F0 int; F1 string }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 int64; F1 uint8; F2 int }
type GenS21 struct { F0 uint16; F1 bool; F2 uint64 }
func genFn0(p0 int64, p1 bool, p2 uint16) string {
    var v5 float64 = 3.14159
    var v6 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return ("\n\t" + ("a" + "\n\t"))
}
func genFn1(p0 int32) string {
    p0 = 2147483647
    v7 := []int{-2147483648, 874, 9223372036854775807}
    _ = p0
    _ = v7
    return ""
}

func main() {
    var v1 int = -582
    var v2 bool = (!(89 == 128))
    var v3 int64 = -1
    var v4 int = (v1 * v1)
    v2 = (79 < 0)
    if v2 {
        var v8 string = ""
        var v9 float32 = ((0.0 + -2.25) / 0.1)
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: 0}
        _ = v10
    }
    v3 -= (v3 << ((42) & 63))
    var v11 int16 = 1000
    if (float64((float32((0.1 / 1.0)) - 0.0)) <= 0.1) {
        var v12 bool = true
        _ = v12
    }
    fmt.Println(v2)
    v13 := GenS2{F0: 18446744073709551615, F1: 567, F2: 1}
    if (18446744073709551615 < 18446744073709551615) {
        v2 = ((!false) || v2)
        v14 := GenS2{F0: 1, F1: 255, F2: 100}
        var v15 uint = 0
        _ = v14
        _ = v15
    } else {
        if v2 {
            if (float64(-2.25) < (-2.25 - 1e10)) {
                var v16 uint32 = 795
                for fi17 := 0; fi17 < 1; fi17++ {
                    switch v4 {
                    case 32767:
                        for fi18 := 0; fi18 < 8; fi18++ {
                            switch "a" {
                            case "a":
                                var v19 bool = ("foo bar" != "foo bar")
                                _ = v19
                            case "foo bar":
                                if true {
                                    if true {
                                        v11 += 1
                                        var v20 int = (fi18 % ((671) | 1))
                                        _ = v20
                                    }
                                } else {
                                    v16 = (v16 | 924)
                                }
                            default:
                                v2 = (v2 || (false && v2))
                            }
                            v21 := GenS3{F0: int(v16), F1: 652, F2: 42}
                            _ = fi18
                            _ = v21
                        }
                    case 797:
                        v4 += (v4 << ((808) & 63))
                    default:
                        if v2 {
                            var v22 uint16 = 114
                            for fi23 := 0; fi23 < 10; fi23++ {
                                var v24 int16 = 0
                                v25 := GenS3{F0: 9223372036854775807, F1: 42, F2: (v3 >> ((4294967295) & 63))}
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS4{F0: (42 > 2), F1: 18446744073709551615}
                            _ = v22
                            _ = v26
                        }
                    }
                    switch (v4 * 127) {
                    case -2147483648:
                        var v27 float64 = (1e10 * 0.0)
                        _ = v27
                    case 2147483647:
                        fmt.Println(v16, v4)
                    case 42:
                        v1 += v1
                    }
                    _ = fi17
                }
                if (v2 || false) {
                    v28 := []int{(v4 ^ (v1 / (((v1 & 890)) | 1))), v4}
                    var v29 bool = v2
                    _ = v28
                    _ = v29
                }
                _ = v16
            } else {
                v30 := GenS5{F0: 1, F1: 65535, F2: -1}
                var v31 uint16 = 986
                for fi32 := 0; fi32 < 20; fi32++ {
                    for fi33 := 0; fi33 < 2; fi33++ {
                        switch 0 {
                        case 9223372036854775807:
                            switch ("a" + "hello") {
                            case "hello":
                                v31 -= (v31 - (v31 << ((842) & 63)))
                            case "foo bar":
                                v34 := []int{668}
                                _ = v34
                            case "\n\t":
                                switch "" {
                                case "foo bar":
                                    v3 = (v3 + int64(v4))
                                case "\n\t":
                                    if v2 {
                                        if (0 != 0) {
                                            switch "" {
                                            case "":
                                                v35 := GenS1{F0: -179}
                                                _ = v35
                                            default:
                                                var v36 uint8 = 153
                                                _ = v36
                                            }
                                        } else {
                                            switch "foo bar" {
                                            case "hello":
                                                fmt.Println(fi33)
                                                switch (-(v4)) {
                                                case 128:
                                                    v37 := GenS6{F0: 365}
                                                    _ = v37
                                                case 9223372036854775807:
                                                    var v38 int16 = (-(v11))
                                                    _ = v38
                                                case 32768:
                                                    v39 := GenS5{F0: 2, F1: 191, F2: -1}
                                                    _ = v39
                                                default:
                                                    switch "" {
                                                    case "foo bar":
                                                        switch -1 {
                                                        case 9223372036854775807:
                                                            fmt.Println(v3, v4)
                                                        case -441:
                                                            if v2 {
                                                                fmt.Println(v4, v4, v4)
                                                                if (210 >= 2) {
                                                                    v40 := GenS5{F0: 703, F1: 2147483647, F2: -2147483648}
                                                                    _ = v40
                                                                } else {
                                                                    switch (-(v1)) {
                                                                    case 946:
                                                                        v4 = (-(v4))
                                                                        if (!v2) {
                                                                            switch v1 {
                                                                            case 127:
                                                                                v41 := []int{32767}
                                                                                _ = v41
                                                                            case -2147483648:
                                                                                if (v31 == (v31 * (v31 >> ((1) & 63)))) {
                                                                                    switch fi32 {
                                                                                    case -859:
                                                                                        v42 := GenS7{F0: int32(v3)}
                                                                                        _ = v42
                                                                                    case -2147483648:
                                                                                        var v43 uint = 771
                                                                                        _ = v43
                                                                                    }
                                                                                    v31 = (v31 / (((v31 | 2)) | 1))
                                                                                }
                                                                            default:
                                                                                if ((v2 || v2) && v2) {
                                                                                    if true {
                                                                                        var v44 uint32 = 601
                                                                                        v3 += (v3 ^ (v3 << ((255) & 63)))
                                                                                        fmt.Println(fi32, v1)
                                                                                        _ = v44
                                                                                    } else {
                                                                                        switch 0 {
                                                                                        case 2147483647:
                                                                                            if (!v2) {
                                                                                                if v2 {
                                                                                                    v4 = (fi32 % (((v1 ^ fi33)) | 1))
                                                                                                }
                                                                                                v2 = (-2.25 == 0.1)
                                                                                                var v45 bool = (!false)
                                                                                                v11 = (v11 + (v11 + v11))
                                                                                                v46 := GenS8{F0: 255}
                                                                                                _ = v45
                                                                                                _ = v46
                                                                                            } else {
                                                                                                v3 = 42
                                                                                                var v47 uint32 = 42
                                                                                                _ = v47
                                                                                            }
                                                                                        case -2147483648:
                                                                                            var v48 uint = 1
                                                                                            _ = v48
                                                                                        }
                                                                                        switch "hello" {
                                                                                        case "":
                                                                                            v11 -= 42
                                                                                        case "foo bar":
                                                                                            if true {
                                                                                                switch v4 {
                                                                                                case -2147483648:
                                                                                                    v2 = v2
                                                                                                case 42:
                                                                                                    v49 := GenS9{F0: (fi32 * (^(v4))), F1: 2, F2: fi33}
                                                                                                    _ = v49
                                                                                                default:
                                                                                                    var v50 uint32 = 42
                                                                                                    _ = v50
                                                                                                }
                                                                                                v31 = v31
                                                                                                v51 := GenS10{F0: 40, F1: 815}
                                                                                                _ = v51
                                                                                            }
                                                                                            switch "" {
                                                                                            case "a":
                                                                                                var v52 int32 = -2147483648
                                                                                                _ = v52
                                                                                            case "foo bar":
                                                                                                v53 := []int{fi32, (-(v1)), (fi33 / ((v4) | 1)), 128}
                                                                                                _ = v53
                                                                                            case "hello":
                                                                                                v1 = 9223372036854775807
                                                                                            default:
                                                                                                var v54 bool = (false && true)
                                                                                                _ = v54
                                                                                            }
                                                                                            v55 := GenS11{F0: v2, F1: (v31 ^ 951)}
                                                                                            _ = v55
                                                                                        default:
                                                                                            v56 := GenS11{F0: v2, F1: v31}
                                                                                            _ = v56
                                                                                        }
                                                                                        v57 := []int{}
                                                                                        _ = v57
                                                                                    }
                                                                                }
                                                                            }
                                                                            v58 := []int{(fi32 / ((0) | 1)), (fi33 - (-(v4)))}
                                                                            v59 := GenS12{F0: -1, F1: 921}
                                                                            _ = v58
                                                                            _ = v59
                                                                        }
                                                                    case 0:
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            switch 405 {
                                                                            case 4294967295:
                                                                                var v60 int = 32768
                                                                                _ = v60
                                                                            default:
                                                                                v61 := GenS1{F0: -2147483648}
                                                                                _ = v61
                                                                            }
                                                                        case "\n\t":
                                                                            if ((((("\n\t" + ("" + ("hello" + (("foo bar" + "") + (("a" + "foo bar") + "hello"))))) + (("\n\t" + ("\n\t" + "\n\t")) + ("\n\t" + "a"))) + "a") + "") > ("" + (("" + ((("" + (("" + "foo bar") + "a")) + "foo bar") + ("foo bar" + "hello"))) + "foo bar"))) {
                                                                                if (32767 <= v31) {
                                                                                    var v62 float32 = float32(float64(-2.25))
                                                                                    _ = v62
                                                                                } else {
                                                                                    v63 := []int{}
                                                                                    fmt.Println(v31, v31, v11)
                                                                                    _ = v63
                                                                                }
                                                                            } else {
                                                                                v64 := GenS12{F0: (fi32 << ((427) & 63)), F1: 0}
                                                                                if (42 < 0) {
                                                                                    v11 = -296
                                                                                    v65 := GenS13{F0: (false || ((100 == 27) || true))}
                                                                                    v3 += v3
                                                                                    _ = v65
                                                                                }
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    if ((v31 + 9) <= 2) {
                                                                                        switch (-(v4)) {
                                                                                        case 42:
                                                                                            var v66 uint16 = v31
                                                                                            _ = v66
                                                                                        case -227:
                                                                                            if v2 {
                                                                                                if v2 {
                                                                                                    var v67 int8 = 0
                                                                                                    v4 -= fi32
                                                                                                    v3 += -1
                                                                                                    _ = v67
                                                                                                }
                                                                                                v68 := []int{int(v31), 364, (^(v4))}
                                                                                                _ = v68
                                                                                            }
                                                                                        case -128:
                                                                                            if v2 {
                                                                                                var v69 uint32 = 65535
                                                                                                var v70 int8 = 47
                                                                                                _ = v69
                                                                                                _ = v70
                                                                                            }
                                                                                        }
                                                                                        v71 := GenS5{F0: 497, F1: 1, F2: 516}
                                                                                        if false {
                                                                                            var v72 bool = ((^(v31)) <= v31)
                                                                                            if v72 {
                                                                                                v73 := []int{(fi32 & (v4 * fi32)), fi33, 2147483647, (v4 & fi33)}
                                                                                                var v74 int64 = v3
                                                                                                fmt.Println(v31, fi33)
                                                                                                var v75 int32 = -2147483648
                                                                                                v76 := GenS6{F0: 0}
                                                                                                _ = v73
                                                                                                _ = v74
                                                                                                _ = v75
                                                                                                _ = v76
                                                                                            }
                                                                                            v77 := GenS13{F0: true}
                                                                                            var v78 int8 = -128
                                                                                            _ = v72
                                                                                            _ = v77
                                                                                            _ = v78
                                                                                        }
                                                                                        _ = v71
                                                                                    }
                                                                                case "foo bar":
                                                                                    v1 -= 890
                                                                                }
                                                                                _ = v64
                                                                            }
                                                                        }
                                                                    case 42:
                                                                        v79 := []int{}
                                                                        v80 := GenS14{F0: 4294967295}
                                                                        switch v4 {
                                                                        case 9223372036854775807:
                                                                            var v81 uint8 = 56
                                                                            _ = v81
                                                                        case -2147483648:
                                                                            var v82 int8 = 100
                                                                            _ = v82
                                                                        }
                                                                        _ = v79
                                                                        _ = v80
                                                                    default:
                                                                        if (0 < -50) {
                                                                            var v83 int32 = 783
                                                                            _ = v83
                                                                        } else {
                                                                            v84 := []int{fi33}
                                                                            _ = len(v84)
                                                                            _ = v84
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            v3 += (v3 + -9223372036854775808)
                                                        }
                                                    case "":
                                                        var v85 string = "hello"
                                                        _ = v85
                                                    case "a":
                                                        v4 = (fi33 << ((255) & 63))
                                                    default:
                                                        var v86 uint64 = 4294967295
                                                        v87 := []int{}
                                                        v4 = (fi32 * 9223372036854775807)
                                                        fmt.Println(fi32, v31, fi33)
                                                        if len(v87) >= 2 {
                                                            v87 = v87[1:len(v87)]
                                                        }
                                                        v88 := GenS10{F0: 127, F1: 4294967295}
                                                        var v89 uint64 = (v86 >> ((42) & 63))
                                                        _ = v86
                                                        _ = v87
                                                        _ = v88
                                                        _ = v89
                                                    }
                                                }
                                                if v2 {
                                                    if (v2 || (true || v2)) {
                                                        v4 += fi32
                                                        v90 := []int{(fi33 & (^(v1)))}
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            if len(v90) > 0 {
                                                                v90[0] = v4
                                                            }
                                                        case "hello":
                                                            if (128 != 128) {
                                                                if len(v90) > 0 {
                                                                    v90[0] = 9223372036854775807
                                                                }
                                                                if len(v90) > 0 {
                                                                    v90[0] = int(v3)
                                                                }
                                                            } else {
                                                                if (!false) {
                                                                    var v91 float32 = -2.25
                                                                    if len(v90) > 0 {
                                                                        v90[0] = (-(fi33))
                                                                    }
                                                                    if v2 {
                                                                        if (!(0.0 <= -2.25)) {
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                var v92 uint = 255
                                                                                _ = v92
                                                                            }
                                                                            v3 += 0
                                                                            v3 -= v3
                                                                        }
                                                                        v11 -= v11
                                                                        v4 = 9223372036854775807
                                                                    }
                                                                    _ = v91
                                                                }
                                                            }
                                                        case "\n\t":
                                                            switch (fi32 - (fi32 % (((^(fi32))) | 1))) {
                                                            case 128:
                                                                fmt.Println(v31, v3)
                                                            }
                                                        default:
                                                            v1 -= (fi33 | (v1 << ((255) & 63)))
                                                        }
                                                        _ = v90
                                                    } else {
                                                        fmt.Println(v31, fi32)
                                                        v11 -= (v11 * v11)
                                                    }
                                                    v93 := GenS8{F0: 2}
                                                    v31 = v31
                                                    _ = v93
                                                }
                                                var v94 float32 = 1.5
                                                v95 := []int{959, v1}
                                                if len(v95) >= 2 {
                                                    v95 = v95[1:len(v95)]
                                                }
                                                if (713 != 28) {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        if v2 {
                                                            v1 += (v1 >> ((789) & 63))
                                                        } else {
                                                            fmt.Println(v94)
                                                        }
                                                    default:
                                                        switch "a" {
                                                        case "foo bar":
                                                            fmt.Println(v3, v94, v2)
                                                        case "a":
                                                            var v96 int64 = (v3 ^ v3)
                                                            _ = v96
                                                        case "":
                                                            var v97 int64 = v3
                                                            _ = v97
                                                        default:
                                                            v2 = v2
                                                        }
                                                    }
                                                    v11 -= -32768
                                                } else {
                                                    if v2 {
                                                        v4 = (fi33 | (fi33 * v4))
                                                        v94 = v94
                                                        var v98 bool = ((!true) || (!(647 < 1)))
                                                        _ = v98
                                                    } else {
                                                        v94 += (v94 * float32((3.14159 + 3.14159)))
                                                    }
                                                    v31 += v31
                                                }
                                                _ = v94
                                                _ = v95
                                            }
                                        }
                                        v99 := GenS15{F0: 4294967295, F1: (^(v11))}
                                        v2 = ((200 > 1) || v2)
                                        _ = v99
                                    } else {
                                        if false {
                                            v3 = -732
                                            switch "a" {
                                            case "hello":
                                                if v2 {
                                                    var v100 float64 = (3.14159 + float64(0.1))
                                                    if true {
                                                        if (!(734 == 4294967295)) {
                                                            if ((v1 << ((0) & 63)) > v1) {
                                                                v11 = 1
                                                            }
                                                            if (true && false) {
                                                                fmt.Println(fi32, v4)
                                                                if v2 {
                                                                    switch "hello" {
                                                                    case "":
                                                                        switch fi32 {
                                                                        case 0:
                                                                            switch v1 {
                                                                            case -231:
                                                                                v2 = (("hello" > "") && (v3 > (-(v3))))
                                                                            default:
                                                                                v4 = v1
                                                                            }
                                                                        default:
                                                                            v101 := []int{}
                                                                            _ = v101
                                                                        }
                                                                    case "hello":
                                                                        v11 += v11
                                                                    }
                                                                    var v102 int32 = -694
                                                                    v11 += v11
                                                                    _ = v102
                                                                }
                                                                if v2 {
                                                                    var v103 uint32 = 2147483647
                                                                    v104 := []int{}
                                                                    v31 += v31
                                                                    _ = v103
                                                                    _ = v104
                                                                }
                                                            }
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                if v2 {
                                                                    if (!(42 <= v3)) {
                                                                        if false {
                                                                            v31 += 65535
                                                                            v100 += ((v100 - (1e10 + v100)) * (v100 * 0.0))
                                                                            v1 += v4
                                                                        } else {
                                                                            switch fi32 {
                                                                            case -232:
                                                                                v105 := []int{(fi33 ^ (-(v1))), (fi33 % (((fi32 - (fi33 * v4))) | 1)), fi32, (v1 >> ((2) & 63))}
                                                                                _ = v105
                                                                            case -923:
                                                                                if (16 >= -109) {
                                                                                    v31 = (v31 << ((1) & 63))
                                                                                } else {
                                                                                    v106 := GenS15{F0: 171, F1: -32768}
                                                                                    if (18446744073709551615 >= 898) {
                                                                                        var v107 uint = 4294967295
                                                                                        var v108 int64 = v3
                                                                                        _ = v107
                                                                                        _ = v108
                                                                                    } else {
                                                                                        var v109 uint64 = 249
                                                                                        if true {
                                                                                            v3 -= (v3 + (v3 * (v3 + (-(v3)))))
                                                                                            fmt.Println(v1, v1, fi33)
                                                                                            v110 := GenS9{F0: (fi32 % ((32768) | 1)), F1: 0, F2: 9223372036854775807}
                                                                                            _ = v110
                                                                                        } else {
                                                                                            v111 := GenS7{F0: -2147483648}
                                                                                            _ = v111
                                                                                        }
                                                                                        _ = v109
                                                                                    }
                                                                                    v11 = v11
                                                                                    _ = v106
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    v112 := GenS15{F0: 581, F1: -56}
                                                                    var v113 int8 = 42
                                                                    _ = v112
                                                                    _ = v113
                                                                }
                                                            case "a":
                                                                if (!(-128 == -128)) {
                                                                    fmt.Println(v100, fi32, v100)
                                                                } else {
                                                                    v114 := []int{(v4 & v4), v4, int(v3)}
                                                                    var v115 int16 = v11
                                                                    _ = v114
                                                                    _ = v115
                                                                }
                                                            }
                                                        }
                                                        v116 := GenS16{F0: 177, F1: -1, F2: 42}
                                                        _ = v116
                                                    }
                                                    if (-2.25 > (1.5 / -2.25)) {
                                                        v4 += (v4 ^ v1)
                                                        v117 := GenS17{F0: 268, F1: (v31 + v31)}
                                                        _ = v117
                                                    } else {
                                                        v118 := GenS18{F0: fi32, F1: ("hello" + "foo bar")}
                                                        fmt.Println(fi32, v31, v2)
                                                        v100 += v100
                                                        v11 = v11
                                                        v119 := []int{v1, int(v3)}
                                                        var v120 float32 = -2.25
                                                        _ = v118
                                                        _ = v119
                                                        _ = v120
                                                    }
                                                    _ = v100
                                                } else {
                                                    switch 1 {
                                                    case 866:
                                                        var v121 uint16 = v31
                                                        _ = v121
                                                    }
                                                    v4 = fi33
                                                }
                                            case "a":
                                                fmt.Println(v4)
                                            default:
                                                var v122 uint16 = (v31 | (v31 % ((v31) | 1)))
                                                _ = v122
                                            }
                                            var v123 uint32 = 0
                                            _ = v123
                                        } else {
                                            v31 += (v31 / (((v31 >> ((0) & 63))) | 1))
                                        }
                                        fmt.Println(v31, v11)
                                        v3 -= (^(v3))
                                    }
                                    fmt.Println(v31, fi32)
                                case "a":
                                    v124 := []int{fi33}
                                    _ = v124
                                }
                            default:
                                v11 = v11
                            }
                        case 32767:
                            v125 := GenS18{F0: (v4 >> ((42) & 63)), F1: ""}
                            _ = v125
                        case -807:
                            v126 := []int{fi33, (v4 - fi32), (-(v4))}
                            v1 -= (fi33 % ((v1) | 1))
                            _ = v126
                        }
                        _ = fi33
                    }
                    _ = fi32
                }
                _ = v30
                _ = v31
            }
            switch v4 {
            case -2147483648:
                var v127 int = v4
                _ = v127
            case -538:
                for fi128 := 0; fi128 < 20; fi128++ {
                    v129 := GenS19{F0: 1}
                    var v130 int8 = 65
                    _ = fi128
                    _ = v129
                    _ = v130
                }
            default:
                v131 := GenS18{F0: (v4 >> ((42) & 63)), F1: ("hello" + "hello")}
                _ = v131
            }
        }
        var v132 int64 = (v3 | int64(v1))
        var v133 uint16 = 0
        _ = v132
        _ = v133
    }
    v4 += v4
    v134 := GenS8{F0: 0}
    v1 = (^(v1))
    if false {
        switch "foo bar" {
        case "hello":
            var v135 uint64 = 1
            _ = v135
        case "\n\t":
            var v136 int16 = (-(v11))
            _ = v136
        case "":
            v11 -= v11
        default:
            if (1 <= v3) {
                fmt.Println(v1, v3, v3)
                fmt.Println(v1, v11)
            } else {
                if (true || (93 != 94)) {
                    v4 -= v4
                    v1 += 2147483647
                } else {
                    v4 += 1
                }
                if v2 {
                    v137 := []int{2147483647, (v4 << ((424) & 63))}
                    _ = v137
                } else {
                    var v138 float64 = 0.1
                    _ = v138
                }
                v139 := GenS20{F0: 9223372036854775807, F1: 255, F2: 127}
                _ = v139
            }
        }
    }
    v140 := GenS21{F0: 2, F1: (-734 < -797), F2: 274}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v134
    _ = v140
    _ = v2
    _ = v3
    _ = v4
}
