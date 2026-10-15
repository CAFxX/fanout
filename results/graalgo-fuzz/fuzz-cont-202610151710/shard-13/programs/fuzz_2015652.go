package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 string; F2 uint16 }
type GenS4 struct { F0 uint; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8; F1 uint }
type GenS8 struct { F0 int; F1 uint32 }
type GenS9 struct { F0 int16; F1 uint; F2 int64 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 bool; F1 int16; F2 uint }
type GenS13 struct { F0 uint16; F1 int64 }
type GenS14 struct { F0 uint8; F1 int16 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 uint64; F1 uint32 }
type GenS17 struct { F0 int64; F1 uint8; F2 uint64 }

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{2147483647, 2147483648}
    var v3 uint32 = 744
    var v4 float64 = 1e10
    var v5 string = "\n\t"
    if (!(false && false)) {
        v3 = v3
    }
    if ((!((!false) || (3.14159 > 0.0))) && ((float32(0.1) + ((0.0 * 3.14159) + 0.1)) > 3.14159)) {
        v6 := GenS1{F0: -875, F1: 768, F2: 0}
        switch 798 {
        case 186:
            v7 := GenS2{F0: (v5 + v5)}
            _ = v7
        case -9223372036854775808:
            var v8 bool = (v5 != (v5 + v5))
            _ = v8
        default:
            if (0 <= 2147483647) {
                var v9 int = 1
                switch v5 {
                case "a":
                    v10 := v1[("" + (v5 + ((v5 + (v5 + ("\n\t" + v5))) + "\n\t")))]
                    _ = v10
                case "\n\t":
                    var v11 int64 = 9223372036854775807
                    _ = v11
                default:
                    var v12 int16 = 1
                    _ = v12
                }
                v5 = v5
                _ = v9
            } else {
                var v13 int = 2147483647
                _ = v13
            }
        }
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 uint8 = 56
            var v16 int8 = 42
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v6
    }
    if true {
        if false {
            v17 := GenS2{F0: ((v5 + (v5 + "foo bar")) + v5)}
            _ = v17
        } else {
            var v18 uint32 = 1
            switch 128 {
            case 2147483647:
                v18 += (v3 / (((^(v18))) | 1))
            }
            _ = v18
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            for fi20 := 0; fi20 < 20; fi20++ {
                if true {
                    if (v4 == (1e10 + v4)) {
                        fmt.Println(fi19, fi19, fi20)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    if (574 <= 255) {
                        v21 := GenS2{F0: v5}
                        var v22 int16 = 1
                        fmt.Println(fi20, fi20, v4)
                        _ = len(v2)
                        switch ("foo bar" + v5) {
                        case "foo bar":
                            var v23 float64 = 3.14159
                            _ = v23
                        case "hello":
                            v3 = v3
                            var v24 float32 = -2.25
                            fmt.Println(fi20)
                            _ = v24
                        default:
                            if true {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                fmt.Println(fi19, v4, v22)
                            } else {
                                var v25 uint = 65535
                                _ = v25
                            }
                        }
                        _ = v21
                        _ = v22
                    } else {
                        var v26 float64 = v4
                        _ = v26
                    }
                    if false {
                        var v27 uint16 = 42
                        var v28 uint8 = 76
                        switch 4294967295 {
                        case -539:
                            if true {
                                if (254 > 300) {
                                    if true {
                                        v29 := GenS2{F0: (((v5 + (v5 + ((("a" + v5) + v5) + v5))) + (v5 + "a")) + v5)}
                                        _ = v29
                                    } else {
                                        fmt.Println(fi19)
                                        v27 = uint16(v3)
                                        var v30 int16 = 1000
                                        _ = v30
                                    }
                                } else {
                                    switch 2147483647 {
                                    case 128:
                                        var v31 float32 = float32(-2.25)
                                        _ = v31
                                    case 127:
                                        v32 := GenS2{F0: v5}
                                        _ = v32
                                    case -223:
                                        switch (fi19 * fi19) {
                                        case 32768:
                                            var v33 bool = ((v4 - ((3.14159 * 1e10) + float64(1e10))) >= v4)
                                            _ = v33
                                        case 32767:
                                            var v34 int = 128
                                            _ = v34
                                        case 9223372036854775807:
                                            switch v5 {
                                            case "hello":
                                                var v35 uint16 = v27
                                                _ = v35
                                            case "a":
                                                _ = len(v1)
                                            case "":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                        }
                                        var v36 int16 = -704
                                        _ = v36
                                    default:
                                        if (1 >= -1) {
                                            v37 := GenS1{F0: (fi20 * 32767), F1: (fi20 << ((0) & 63)), F2: (v27 + (v27 / (((v27 - (v27 + (^(v27))))) | 1)))}
                                            if true {
                                                if true {
                                                    var v38 int32 = 0
                                                    v27 -= v27
                                                    v3 -= v3
                                                    v39 := GenS1{F0: fi19, F1: fi19, F2: v27}
                                                    fmt.Println(v38)
                                                    if true {
                                                        v40 := GenS3{F0: (!false), F1: (v5 + v5), F2: v27}
                                                        v41 := GenS1{F0: fi20, F1: fi19, F2: (v27 | v27)}
                                                        switch (^(fi19)) {
                                                        case 127:
                                                            fmt.Println(v5, v28)
                                                        case 9223372036854775807:
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                switch v5 {
                                                                case "hello":
                                                                    if (!(!(!false))) {
                                                                        var v42 float32 = (0.1 * 1.5)
                                                                        _ = v42
                                                                    }
                                                                case "":
                                                                    if true {
                                                                        var v43 float32 = 1.5
                                                                        delete(v1, (("\n\t" + "a") + v5))
                                                                        _ = v43
                                                                    } else {
                                                                        _ = len(v1)
                                                                        if (111 <= -1) {
                                                                            v3 = 2147483647
                                                                            v5 = ("foo bar" + "\n\t")
                                                                            var v44 float32 = -2.25
                                                                            switch "hello" {
                                                                            case "a":
                                                                                var v45 float64 = v4
                                                                                _ = v45
                                                                            }
                                                                            var v46 int64 = -1
                                                                            v47 := GenS2{F0: (v5 + v5)}
                                                                            _ = v44
                                                                            _ = v46
                                                                            _ = v47
                                                                        } else {
                                                                            if false {
                                                                                v27 -= (v27 << ((uint(v3)) & 63))
                                                                            } else {
                                                                                v28 -= (^(v28))
                                                                            }
                                                                            switch fi19 {
                                                                            case 128:
                                                                                var v48 uint32 = 230
                                                                                _ = v48
                                                                            default:
                                                                                _ = len(v2)
                                                                            }
                                                                        }
                                                                        var v49 uint16 = (v27 >> ((255) & 63))
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            v50 := v1["\n\t"]
                                                                            _ = v50
                                                                        case "hello":
                                                                            if true {
                                                                                switch (fi20 << ((518) & 63)) {
                                                                                case 2147483647:
                                                                                    v5 = v5
                                                                                default:
                                                                                    switch 1 {
                                                                                    case 4294967295:
                                                                                        if (int16(v49) == 650) {
                                                                                            var v51 bool = false
                                                                                            _ = v51
                                                                                        } else {
                                                                                            if true {
                                                                                                switch 42 {
                                                                                                case 0:
                                                                                                    var v52 float32 = (1.5 - (0.0 / (-2.25 * (1e10 + 1.5))))
                                                                                                    _ = v52
                                                                                                case -9223372036854775808:
                                                                                                    if false {
                                                                                                        v53 := GenS4{F0: 475, F1: (-2147483648 < (v38 ^ v38)), F2: (v3 >> ((245) & 63))}
                                                                                                        fmt.Println(fi20)
                                                                                                        _ = v53
                                                                                                    }
                                                                                                    v38 -= v38
                                                                                                    fmt.Println(v5, v3)
                                                                                                    if (100 <= 0) {
                                                                                                        if len(v2) >= 2 {
                                                                                                            v2 = v2[1:len(v2)]
                                                                                                        }
                                                                                                    } else {
                                                                                                        v1["foo bar"] = true
                                                                                                    }
                                                                                                    switch v5 {
                                                                                                    case "\n\t":
                                                                                                        var v54 bool = true
                                                                                                        _ = v54
                                                                                                    case "a":
                                                                                                        if (0.0 <= float32(1e10)) {
                                                                                                            delete(v1, v5)
                                                                                                            var v55 int64 = -948
                                                                                                            v56 := GenS5{F0: 2, F1: ((float64(((float32(3.14159) - 0.1) / float32(float64((3.14159 + -2.25))))) * float64((0.1 - -2.25))) >= v4)}
                                                                                                            _ = v55
                                                                                                            _ = v56
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                v5 = v5
                                                                                            } else {
                                                                                                fmt.Println(v27, v3, fi19)
                                                                                                v57 := GenS6{F0: 9223372036854775807}
                                                                                                v3 -= v3
                                                                                                _ = v57
                                                                                            }
                                                                                            var v58 uint32 = (v3 << ((4294967295) & 63))
                                                                                            _ = v58
                                                                                        }
                                                                                    case -1:
                                                                                        v5 = ((v5 + ((v5 + "") + v5)) + v5)
                                                                                    }
                                                                                }
                                                                            }
                                                                        case "a":
                                                                            var v59 string = "a"
                                                                            _ = v59
                                                                        default:
                                                                            var v60 bool = (fi20 <= (fi20 - fi19))
                                                                            _ = v60
                                                                        }
                                                                        _ = v49
                                                                    }
                                                                case "foo bar":
                                                                    v61 := GenS1{F0: (fi19 | (fi20 & fi20)), F1: 1, F2: v27}
                                                                    _ = v61
                                                                default:
                                                                    v62 := GenS7{F0: int8(v28), F1: 862}
                                                                    _ = v62
                                                                }
                                                            case "":
                                                                _ = len(v2)
                                                            case "a":
                                                                v3 += (v3 + (v3 | v3))
                                                            }
                                                        case -696:
                                                            switch fi20 {
                                                            case -2147483648:
                                                                var v63 uint16 = v27
                                                                if (v3 == (v3 & (v3 >> ((18446744073709551615) & 63)))) {
                                                                    var v64 int64 = 1
                                                                    _ = len(v1)
                                                                    fmt.Println(fi20, v28)
                                                                    _ = v64
                                                                } else {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    _ = len(v1)
                                                                }
                                                                if false {
                                                                    fmt.Println(v28)
                                                                    if true {
                                                                        switch 32767 {
                                                                        case 127:
                                                                            var v65 float64 = ((((v4 + (-2.25 - (v4 + float64(0.0)))) - v4) / v4) / v4)
                                                                            _ = v65
                                                                        case 400:
                                                                            v66 := GenS1{F0: -332, F1: 0, F2: uint16(v28)}
                                                                            _ = v66
                                                                        case 2147483647:
                                                                            var v67 uint32 = 65535
                                                                            _ = v67
                                                                        }
                                                                    }
                                                                }
                                                                _ = v63
                                                            }
                                                            var v68 float32 = 3.14159
                                                            v69 := GenS8{F0: fi20, F1: 154}
                                                            _ = v68
                                                            _ = v69
                                                        default:
                                                            if false {
                                                                v38 += v38
                                                            }
                                                        }
                                                        _ = v40
                                                        _ = v41
                                                    } else {
                                                        switch fi20 {
                                                        case 0:
                                                            v28 += (^(v28))
                                                        }
                                                        v3 = (v3 & 170)
                                                        var v70 int16 = 32767
                                                        v71 := GenS7{F0: -49, F1: 610}
                                                        var v72 bool = false
                                                        fmt.Println(v38, v3)
                                                        fmt.Println(v38)
                                                        _ = v70
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    fmt.Println(fi19)
                                                    if (!false) {
                                                        switch -855 {
                                                        case 0:
                                                            var v73 uint16 = (v27 % ((v27) | 1))
                                                            _ = v73
                                                        case -2147483648:
                                                            fmt.Println(v27, fi20, v4)
                                                        case 128:
                                                            fmt.Println(v3, v5, v3)
                                                        }
                                                        var v74 int32 = int32(v3)
                                                        _ = len(v2)
                                                        _ = v74
                                                    } else {
                                                        var v75 int32 = (v38 << ((18446744073709551615) & 63))
                                                        v76 := GenS1{F0: (fi19 | fi19), F1: fi19, F2: v27}
                                                        var v77 int8 = 113
                                                        if false {
                                                            v4 += (float64(3.14159) - ((3.14159 / ((1.5 * (float64(1.5) * (-2.25 / v4))) + v4)) - (-2.25 + v4)))
                                                        } else {
                                                            switch v5 {
                                                            case "\n\t":
                                                                switch fi19 {
                                                                case 23:
                                                                    v77 = 54
                                                                case 0:
                                                                    v5 = v5
                                                                case 4294967295:
                                                                    v78 := GenS9{F0: -32768, F1: 2, F2: int64(v3)}
                                                                    _ = v78
                                                                default:
                                                                    var v79 bool = false
                                                                    _ = v79
                                                                }
                                                            case "foo bar":
                                                                switch (fi20 % ((fi19) | 1)) {
                                                                case 42:
                                                                    v80 := GenS2{F0: ("" + v5)}
                                                                    _ = v80
                                                                case 4294967295:
                                                                    switch v5 {
                                                                    case "\n\t":
                                                                        if (!false) {
                                                                            switch (fi19 - fi20) {
                                                                            case -613:
                                                                                if (!false) {
                                                                                    v81 := GenS4{F0: 1, F1: true, F2: (v3 | (v3 / (((v3 % (((v3 * v3)) | 1))) | 1)))}
                                                                                    if (32767 < -299) {
                                                                                        v38 -= v75
                                                                                        if ((v27 / (((v27 * v27)) | 1)) == (v27 << ((430) & 63))) {
                                                                                            v82 := v1[((v5 + v5) + "foo bar")]
                                                                                            _ = v82
                                                                                        } else {
                                                                                            var v83 int = 32767
                                                                                            switch ("" + v5) {
                                                                                            case "a":
                                                                                                if ((18446744073709551615 >= 1) || ((v3 << ((420) & 63)) == v3)) {
                                                                                                    v4 = 3.14159
                                                                                                } else {
                                                                                                    if false {
                                                                                                        v84 := GenS7{F0: v77, F1: 688}
                                                                                                        if (93 <= (v77 * 42)) {
                                                                                                            var v85 int64 = 0
                                                                                                            var v86 uint8 = v28
                                                                                                            if (true && (true && false)) {
                                                                                                                var v87 int16 = -88
                                                                                                                v83 = (fi19 & -524)
                                                                                                                v28 = (v86 * 1)
                                                                                                                _ = v87
                                                                                                            } else {
                                                                                                                v88 := GenS6{F0: v85}
                                                                                                                if ((v75 + 42) <= v38) {
                                                                                                                    var v89 float64 = 0.0
                                                                                                                    v3 -= (v3 % (((^(v3))) | 1))
                                                                                                                    v4 = float64(0.0)
                                                                                                                    _ = v89
                                                                                                                } else {
                                                                                                                    var v90 int8 = 100
                                                                                                                    var v91 float32 = (0.1 * (float32(v4) / 1.5))
                                                                                                                    _ = v90
                                                                                                                    _ = v91
                                                                                                                }
                                                                                                                var v92 int8 = v77
                                                                                                                _ = v88
                                                                                                                _ = v92
                                                                                                            }
                                                                                                            _ = v85
                                                                                                            _ = v86
                                                                                                        }
                                                                                                        _ = v84
                                                                                                    } else {
                                                                                                        fmt.Println(v28, v27, v28)
                                                                                                    }
                                                                                                    v75 = (v38 - v75)
                                                                                                }
                                                                                                var v93 float32 = 0.0
                                                                                                switch (v83 & (fi20 >> ((989) & 63))) {
                                                                                                case -187:
                                                                                                    v77 -= -128
                                                                                                case 9223372036854775807:
                                                                                                    switch (-(fi20)) {
                                                                                                    case 128:
                                                                                                        v4 += 1e10
                                                                                                    case -63:
                                                                                                        if false {
                                                                                                            var v94 int8 = 127
                                                                                                            var v95 uint64 = 1
                                                                                                            var v96 int8 = (v94 + (v77 | (v94 & 42)))
                                                                                                            _ = len(v1)
                                                                                                            v97 := GenS10{F0: v27}
                                                                                                            _ = v94
                                                                                                            _ = v95
                                                                                                            _ = v96
                                                                                                            _ = v97
                                                                                                        } else {
                                                                                                            switch "\n\t" {
                                                                                                            case "":
                                                                                                                if (4294967295 <= uint(v27)) {
                                                                                                                    fmt.Println(v27, v4, fi19)
                                                                                                                    if len(v2) >= 2 {
                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                    }
                                                                                                                    switch ("\n\t" + v5) {
                                                                                                                    case "foo bar":
                                                                                                                        fmt.Println(v38, v77)
                                                                                                                    case "":
                                                                                                                        v98 := v1[((v5 + "\n\t") + v5)]
                                                                                                                        _ = v98
                                                                                                                    default:
                                                                                                                        var v99 bool = true
                                                                                                                        _ = v99
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    var v100 int = (v83 / ((fi20) | 1))
                                                                                                                    v83 = fi19
                                                                                                                    v4 = ((0.1 / 1.5) / float64(v93))
                                                                                                                    if false {
                                                                                                                        var v101 float64 = 3.14159
                                                                                                                        _ = v101
                                                                                                                    }
                                                                                                                    fmt.Println(v3, v100, v75)
                                                                                                                    delete(v1, (v5 + v5))
                                                                                                                    _ = v100
                                                                                                                }
                                                                                                            default:
                                                                                                                v83 += v83
                                                                                                            }
                                                                                                            v2 = append(v2, (v3 >> ((4294967295) & 63)))
                                                                                                            var v102 int64 = 1
                                                                                                            _ = v102
                                                                                                        }
                                                                                                    case 0:
                                                                                                        _ = len(v2)
                                                                                                    default:
                                                                                                        v93 -= v93
                                                                                                    }
                                                                                                default:
                                                                                                    if (!false) {
                                                                                                        v2 = append(v2, (v3 & (v3 << ((765) & 63))))
                                                                                                        var v103 float64 = v4
                                                                                                        _ = v103
                                                                                                    } else {
                                                                                                        var v104 uint16 = 2
                                                                                                        _ = v104
                                                                                                    }
                                                                                                }
                                                                                                _ = v93
                                                                                            case "\n\t":
                                                                                                if true {
                                                                                                    fmt.Println(v27)
                                                                                                    var v105 uint8 = 2
                                                                                                    v106 := GenS11{F0: "foo bar"}
                                                                                                    _ = v105
                                                                                                    _ = v106
                                                                                                }
                                                                                            }
                                                                                            v107 := GenS12{F0: (v77 >= (v77 & -1)), F1: -952, F2: 18446744073709551615}
                                                                                            _ = v107
                                                                                            _ = v83
                                                                                        }
                                                                                        switch ("" + v5) {
                                                                                        case "":
                                                                                            v5 = ("" + (v5 + "foo bar"))
                                                                                        default:
                                                                                            v38 = (v38 / (((v38 + 1)) | 1))
                                                                                        }
                                                                                    }
                                                                                    _ = v81
                                                                                } else {
                                                                                    var v108 int8 = v77
                                                                                    v38 += (v38 - (v75 | v38))
                                                                                    v38 += -142
                                                                                    _ = v108
                                                                                }
                                                                            case 4294967295:
                                                                                v38 -= v75
                                                                            case 2147483647:
                                                                                var v109 uint64 = 728
                                                                                _ = v109
                                                                            default:
                                                                                v4 += ((0.1 - 0.0) / v4)
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        var v110 int = int(v77)
                                                                        _ = v110
                                                                    default:
                                                                        switch (fi19 | fi20) {
                                                                        case -9223372036854775808:
                                                                            v3 = 363
                                                                        case 973:
                                                                            var v111 string = ("hello" + v5)
                                                                            _ = v111
                                                                        case 128:
                                                                            v75 = 42
                                                                        default:
                                                                            if false {
                                                                                var v112 int32 = v75
                                                                                v27 += 65535
                                                                                _ = v112
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    var v113 float64 = ((v4 / ((v4 + v4) / v4)) * v4)
                                                                    _ = v113
                                                                }
                                                            default:
                                                                v114 := GenS6{F0: -58}
                                                                _ = v114
                                                            }
                                                            v115 := GenS10{F0: (v27 / ((v27) | 1))}
                                                            if false {
                                                                v77 += v77
                                                                var v116 float64 = (v4 / (v4 - float64(((0.0 + 0.0) * float32(float64((((-2.25 - 0.0) / 1.0) + (-2.25 + 0.1))))))))
                                                                var v117 int64 = 1
                                                                _ = v116
                                                                _ = v117
                                                            }
                                                            switch 42 {
                                                            case -877:
                                                                if true {
                                                                    v118 := GenS12{F0: false, F1: 1, F2: 1}
                                                                    _ = v118
                                                                }
                                                            case 32768:
                                                                if false {
                                                                    var v119 int32 = 123
                                                                    v120 := GenS13{F0: v27, F1: -9223372036854775808}
                                                                    _ = v119
                                                                    _ = v120
                                                                } else {
                                                                    v121 := GenS8{F0: int(v27), F1: (v3 / (((v3 | (v3 << ((942) & 63)))) | 1))}
                                                                    _ = v121
                                                                }
                                                                if (v75 != (v38 % (((v38 % ((v75) | 1))) | 1))) {
                                                                    switch v5 {
                                                                    case "a":
                                                                        v122 := GenS13{F0: (v27 + v27), F1: 0}
                                                                        _ = v122
                                                                    default:
                                                                        v123 := GenS4{F0: 714, F1: (-362 > fi19), F2: (v3 - 0)}
                                                                        _ = v123
                                                                    }
                                                                    v2 = append(v2, (v3 | (^(v3))))
                                                                    if false {
                                                                        fmt.Println(v3, v4, v3)
                                                                        var v124 int8 = -1
                                                                        _ = v124
                                                                    }
                                                                } else {
                                                                    v125 := v1[""]
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        v126 := v1[v5]
                                                                        if len(v2) > 0 {
                                                                            v2[0] = 808
                                                                        }
                                                                        _ = len(v2)
                                                                        _ = v126
                                                                    case "hello":
                                                                        v38 -= int32(v27)
                                                                    default:
                                                                        var v127 float64 = float64(0.1)
                                                                        _ = v127
                                                                    }
                                                                    _ = v125
                                                                }
                                                                v2 = append(v2, (^(v3)))
                                                            case -830:
                                                                if len(v2) > 0 {
                                                                    v2[0] = (v3 - (v3 & (^(v3))))
                                                                }
                                                            default:
                                                                v3 += (v3 << ((263) & 63))
                                                            }
                                                            var v128 int = fi20
                                                            _ = v115
                                                            _ = v128
                                                        }
                                                        if true {
                                                            v129 := GenS6{F0: 149}
                                                            delete(v1, "")
                                                            fmt.Println(v4, v77, v38)
                                                            _ = v129
                                                        } else {
                                                            if true {
                                                                if (fi19 > fi19) {
                                                                    v2 = append(v2, v3)
                                                                    v2 = append(v2, (v3 - 2147483648))
                                                                    var v130 string = "a"
                                                                    _ = v130
                                                                }
                                                                var v131 int64 = 42
                                                                fmt.Println(v131, fi19)
                                                                var v132 int8 = v77
                                                                _ = v131
                                                                _ = v132
                                                            } else {
                                                                var v133 uint16 = (v27 | 472)
                                                                _ = v133
                                                            }
                                                            v134 := GenS14{F0: (^(v28)), F1: -566}
                                                            var v135 int8 = (-(v77))
                                                            _ = v134
                                                            _ = v135
                                                        }
                                                        _ = v75
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    v28 = (v28 - v28)
                                                }
                                            } else {
                                                switch (fi20 & 318) {
                                                case -773:
                                                    var v136 bool = (false && true)
                                                    v136 = v136
                                                    _ = v136
                                                default:
                                                    _ = len(v1)
                                                }
                                                v1[v5] = (v3 > (v3 | 65535))
                                                var v137 string = "hello"
                                                _ = v137
                                            }
                                            _ = v37
                                        }
                                    }
                                    delete(v1, (v5 + "\n\t"))
                                    if true {
                                        var v138 int8 = 127
                                        _ = v138
                                    }
                                    v139 := GenS9{F0: -760, F1: 255, F2: -1}
                                    var v140 int32 = 0
                                    _ = v139
                                    _ = v140
                                }
                                if false {
                                    _ = len(v1)
                                    v141 := GenS3{F0: false, F1: ("" + ""), F2: (v27 + (v27 / ((v27) | 1)))}
                                    v4 += float64((0.0 - ((3.14159 / 1.0) * 1.5)))
                                    _ = v141
                                } else {
                                    fmt.Println(v27)
                                    v5 = v5
                                }
                            } else {
                                var v142 uint = uint(v28)
                                v3 += v3
                                if ((fi19 > int(v142)) || (!false)) {
                                    v4 -= v4
                                    if false {
                                        delete(v1, ("hello" + "a"))
                                    }
                                }
                                _ = v142
                            }
                        }
                        _ = v27
                        _ = v28
                    } else {
                        var v143 float32 = 1e10
                        if (((v4 / (((float64((1e10 + v143)) * 0.1) * v4) / v4)) < float64(((1e10 + (((v143 + (3.14159 / (v143 * float32(v4)))) * float32(-2.25)) + v143)) * v143))) && (true && (true && (v3 <= (v3 ^ 2147483648))))) {
                            v144 := GenS15{F0: 65535}
                            _ = v144
                        }
                        _ = v143
                    }
                    var v145 int8 = 100
                    _ = v145
                }
                _ = fi20
            }
            for fi146 := 0; fi146 < 16; fi146++ {
                v5 = v5
                v3 = v3
                var v147 int32 = -632
                v148 := GenS12{F0: false, F1: 42, F2: 2}
                if false {
                    v149 := GenS7{F0: 42, F1: 1}
                    _ = v149
                } else {
                    var v150 int = fi146
                    var v151 uint64 = 78
                    v5 = ("\n\t" + ("\n\t" + "a"))
                    var v152 uint8 = 128
                    _ = v150
                    _ = v151
                    _ = v152
                }
                _ = fi146
                _ = v147
                _ = v148
            }
            _ = fi19
        }
        for fi153 := 0; fi153 < 5; fi153++ {
            fmt.Println(v3, v3)
            if (int32(v3) > 0) {
                var v154 float32 = 1.5
                v1[((v5 + v5) + "\n\t")] = true
                var v155 string = v5
                _ = v154
                _ = v155
            } else {
                fmt.Println(v3, v3)
            }
            for fi156 := 0; fi156 < 20; fi156++ {
                v157 := GenS7{F0: 42, F1: 189}
                _ = fi156
                _ = v157
            }
            _ = fi153
        }
    }
    if (42 < 32767) {
        switch "hello" {
        case "hello":
            var v158 string = v5
            _ = v158
        default:
            if false {
                fmt.Println(v3, v5, v5)
                v3 -= v3
            }
        }
        var v159 int64 = 42
        _ = v159
    }
    v2 = append(v2, 2147483648)
    for fi160 := 0; fi160 < 17; fi160++ {
        for fi161 := 0; fi161 < 13; fi161++ {
            if (!(0 < (^(v3)))) {
                var v162 uint64 = 4294967295
                if false {
                    v4 += (((-2.25 - 3.14159) - v4) * v4)
                    switch v5 {
                    case "a":
                        _ = len(v2)
                    case "":
                        switch v5 {
                        case "":
                            var v163 int = (fi161 & fi160)
                            _ = v163
                        case "hello":
                            v162 += 4294967295
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                } else {
                    v164 := v1[((("hello" + v5) + "a") + "foo bar")]
                    v165 := GenS16{F0: v162, F1: (^(v3))}
                    v166 := GenS17{F0: 42, F1: 255, F2: 876}
                    _ = v164
                    _ = v165
                    _ = v166
                }
                v4 += (v4 / v4)
                _ = v162
            } else {
                var v167 float64 = v4
                if false {
                    if false {
                        var v168 int = (fi161 ^ fi161)
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            switch (v5 + ("foo bar" + "\n\t")) {
                            case "a":
                                v2 = append(v2, v3)
                            case "":
                                v169 := GenS14{F0: 2, F1: 323}
                                _ = v169
                            case "foo bar":
                                v170 := GenS2{F0: "foo bar"}
                                _ = v170
                            }
                            if (189 == 0) {
                                switch v168 {
                                case 128:
                                    fmt.Println(v168, v168, v4)
                                }
                                var v171 int64 = 9223372036854775807
                                _ = v171
                            }
                        } else {
                            v167 = v4
                            var v172 int = fi161
                            _ = v172
                        }
                        _ = v168
                    }
                } else {
                    if (v3 != (v3 ^ (^(v3)))) {
                        v173 := GenS5{F0: 841, F1: false}
                        _ = v173
                    } else {
                        var v174 string = v5
                        if true {
                            v175 := GenS9{F0: -1, F1: 1, F2: -1}
                            var v176 int8 = -1
                            var v177 int64 = 1
                            _ = v175
                            _ = v176
                            _ = v177
                        }
                        _ = v174
                    }
                }
                _ = v167
            }
            v3 += 2147483647
            var v178 int8 = 50
            _ = fi161
            _ = v178
        }
        var v179 string = ((v5 + v5) + v5)
        var v180 int64 = -9223372036854775808
        _ = fi160
        _ = v179
        _ = v180
    }
    switch 32768 {
    case -1:
        var v181 int = -682
        _ = v181
    case 32767:
        var v182 float64 = float64(0.1)
        _ = v182
    case 185:
        switch (("\n\t" + "") + "") {
        case "":
            v3 += 421
        }
    default:
        fmt.Println(v3, v4)
    }
    for fi183 := 0; fi183 < 15; fi183++ {
        fmt.Println(v5)
        _ = fi183
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
