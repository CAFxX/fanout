package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 int32; F2 int }
type GenS5 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS6 struct { F0 int; F1 int16 }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string; F1 uint64; F2 string }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int; F1 int8; F2 int16 }
type GenS13 struct { F0 uint64; F1 uint32 }
type GenS14 struct { F0 bool; F1 string; F2 int64 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 bool; F1 int64; F2 int }
type GenS17 struct { F0 uint; F1 uint64 }
type GenS18 struct { F0 uint8; F1 uint64 }
type GenS19 struct { F0 int16; F1 uint16 }
type GenS20 struct { F0 bool; F1 uint }
type GenS21 struct { F0 uint32; F1 uint }
type GenS22 struct { F0 int16; F1 int16 }
type GenS23 struct { F0 uint; F1 int32 }
type GenS24 struct { F0 uint8 }
type GenS25 struct { F0 uint8; F1 int8; F2 int }
type GenS26 struct { F0 bool; F1 int64 }
type GenS27 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS28 struct { F0 uint64; F1 bool; F2 bool }
type GenS29 struct { F0 uint }
func genFn0(p0 int32) uint64 {
    v7 := GenS1{F0: 32767, F1: -1, F2: 255}
    var v8 float64 = (((1e10 + 0.1) + float64(3.14159)) + 1e10)
    _ = p0
    _ = v7
    _ = v8
    return 573
}

func main() {
    v1 := []uint32{4294967295}
    var v2 uint = 1
    var v3 float32 = 3.14159
    v4 := []string{"foo bar"}
    v5 := make(map[string]uint32)
    var v6 uint8 = 196
    switch 9223372036854775807 {
    case 4294967295:
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 int8 = 1
            v6 -= (v6 * v6)
            _ = fi9
            _ = v10
        }
    case 9223372036854775807:
        if (0 == 429) {
            var v11 int16 = 1000
            v12 := GenS2{F0: -1, F1: genFn0(-1), F2: (v6 & v6)}
            v11 = v11
            _ = v11
            _ = v12
        }
    case 32767:
        if false {
            switch 769 {
            case 9223372036854775807:
                v6 = (v6 + 2)
            case 2147483647:
                if false {
                    v3 += (float32((3.14159 / 1e10)) / v3)
                    for fi13 := 0; fi13 < 19; fi13++ {
                        v6 += (v6 - v6)
                        switch (("foo bar" + ("hello" + "hello")) + ((("hello" + "a") + "hello") + ("foo bar" + "a"))) {
                        case "hello":
                            var v14 uint8 = (v6 * v6)
                            _ = v14
                        }
                        v3 -= v3
                        _ = fi13
                    }
                    if true {
                        v6 = v6
                    } else {
                        v15 := GenS2{F0: 394, F1: 0, F2: 255}
                        delete(v5, "a")
                        _ = v15
                    }
                } else {
                    for fi16 := 0; fi16 < 11; fi16++ {
                        if (!(-459 <= -44)) {
                            var v17 uint16 = uint16(v2)
                            for fi18 := 0; fi18 < 18; fi18++ {
                                v3 = ((v3 + 1e10) + v3)
                                if (!false) {
                                    v19 := GenS3{F0: 9223372036854775807}
                                    _ = v19
                                } else {
                                    v2 = (v2 | 824)
                                    v20 := GenS1{F0: (^(v17)), F1: -1, F2: v17}
                                    switch "hello" {
                                    case "a":
                                        if true {
                                            var v21 string = (("a" + "hello") + "hello")
                                            if false {
                                                var v22 int8 = -30
                                                fmt.Println(fi16, v2, v3)
                                                _ = v22
                                            } else {
                                                if true {
                                                    v21 = v21
                                                } else {
                                                    v17 = 1
                                                }
                                                var v23 int8 = 42
                                                v6 += (v6 + (v6 * v6))
                                                _ = v23
                                            }
                                            fmt.Println(v2)
                                            v3 = ((float32(3.14159) + 1e10) / ((v3 / (((3.14159 / 1.0) + 1.5) * (3.14159 * v3))) - v3))
                                            if false {
                                                var v24 uint64 = 0
                                                var v25 int8 = 1
                                                _ = v24
                                                _ = v25
                                            }
                                            _ = v21
                                        } else {
                                            var v26 uint = 4294967295
                                            _ = v26
                                        }
                                    case "":
                                        v27 := GenS1{F0: v17, F1: 0, F2: v17}
                                        _ = v27
                                    case "hello":
                                        switch "hello" {
                                        case "hello":
                                            var v28 bool = false
                                            _ = v28
                                        case "\n\t":
                                            v6 = (v6 | (^(v6)))
                                            if true {
                                                switch "foo bar" {
                                                case "":
                                                    if len(v4) > 0 {
                                                        v4[0] = "\n\t"
                                                    }
                                                case "hello":
                                                    switch 42 {
                                                    case 2147483647:
                                                        if false {
                                                            switch (fi18 - fi16) {
                                                            case -814:
                                                                v29 := GenS4{F0: -1, F1: -2147483648, F2: (fi18 % ((317) | 1))}
                                                                _ = v29
                                                            case 42:
                                                                var v30 int8 = 0
                                                                _ = v30
                                                            case -1:
                                                                var v31 int64 = 808
                                                                _ = v31
                                                            default:
                                                                v32 := GenS5{F0: 4294967295, F1: -128, F2: (v6 / (((v6 / (((v6 + (v6 | (v6 % ((160) | 1))))) | 1))) | 1))}
                                                                _ = v32
                                                            }
                                                            v33 := GenS1{F0: (v17 % (((v17 - (v17 >> ((v2) & 63)))) | 1)), F1: -43, F2: v17}
                                                            var v34 bool = false
                                                            _ = v33
                                                            _ = v34
                                                        }
                                                    }
                                                case "a":
                                                    var v35 float64 = 1e10
                                                    _ = v35
                                                }
                                                switch ("hello" + "foo bar") {
                                                case "foo bar":
                                                    fmt.Println(v17, fi16, v17)
                                                case "a":
                                                    switch fi16 {
                                                    case 128:
                                                        _ = len(v4)
                                                        v36 := GenS6{F0: fi16, F1: 1000}
                                                        _ = v36
                                                    case 464:
                                                        if len(v1) > 0 {
                                                            v1[0] = 1
                                                        }
                                                    default:
                                                        if true {
                                                            v2 -= (v2 - v2)
                                                            var v37 bool = true
                                                            var v38 float32 = 3.14159
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                    }
                                                case "":
                                                    v2 = 1000
                                                default:
                                                    switch fi16 {
                                                    case 1:
                                                        var v39 uint32 = 65535
                                                        _ = v39
                                                    }
                                                }
                                                if (-711 != int16(v6)) {
                                                    switch "\n\t" {
                                                    case "":
                                                        v2 = (v2 >> (((^(v2))) & 63))
                                                    }
                                                    v40 := GenS2{F0: -765, F1: uint64(v17), F2: 160}
                                                    v3 = 1.5
                                                    _ = v40
                                                }
                                                switch fi18 {
                                                case 850:
                                                    var v41 uint = 42
                                                    _ = v41
                                                case 9223372036854775807:
                                                    fmt.Println(fi18)
                                                case -1:
                                                    v2 = 65535
                                                default:
                                                    var v42 int8 = 0
                                                    _ = v42
                                                }
                                                if (v2 > v2) {
                                                    v43 := GenS5{F0: uint64(v2), F1: 42, F2: v6}
                                                    _ = v43
                                                } else {
                                                    v44 := GenS1{F0: (v17 & 255), F1: 0, F2: (^(v17))}
                                                    v1 = append(v1, 65535)
                                                    _ = v44
                                                }
                                            } else {
                                                var v45 int = (^(fi18))
                                                _ = v45
                                            }
                                            fmt.Println(fi16, v17, fi18)
                                        case "a":
                                            v6 += 32
                                            switch "foo bar" {
                                            case "hello":
                                                _ = len(v1)
                                                v46 := v5[(("hello" + ("\n\t" + "\n\t")) + ("a" + "\n\t"))]
                                                _ = v46
                                            case "a":
                                                v2 = v2
                                            case "\n\t":
                                                var v47 int64 = -1
                                                _ = v47
                                            }
                                            v48 := GenS7{F0: (v2 % ((v2) | 1)), F1: v17}
                                            _ = v48
                                        default:
                                            switch ("\n\t" + ("hello" + ("a" + ("hello" + "")))) {
                                            case "a":
                                                var v49 float64 = 1.5
                                                _ = v49
                                            default:
                                                v6 -= (v6 - v6)
                                            }
                                        }
                                    default:
                                        fmt.Println(v3, v6, fi18)
                                    }
                                    if false {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if (!false) {
                                            v50 := GenS2{F0: -1, F1: 2, F2: (^(v6))}
                                            _ = v50
                                        }
                                    }
                                    _ = v20
                                }
                                _ = fi18
                            }
                            if (1 >= 0) {
                                fmt.Println(v6)
                            }
                            _ = v17
                        } else {
                            if true {
                                var v51 int16 = 1000
                                switch 732 {
                                case 4294967295:
                                    var v52 string = ("a" + "")
                                    _ = v52
                                case 42:
                                    delete(v5, "a")
                                default:
                                    for fi53 := 0; fi53 < 15; fi53++ {
                                        v54 := v5["hello"]
                                        switch fi16 {
                                        case 1:
                                            if false {
                                                fmt.Println(fi53, fi16, v6)
                                                var v55 int16 = v51
                                                _ = v55
                                            } else {
                                                v56 := GenS8{F0: 42}
                                                var v57 uint8 = 2
                                                _ = v56
                                                _ = v57
                                            }
                                        case -9223372036854775808:
                                            v2 = (v2 / (((v2 | 65535)) | 1))
                                        }
                                        if (true || true) {
                                            if ((1.5 / 1e10) < 0.0) {
                                                if true {
                                                    if ((v6 * v6) > v6) {
                                                        var v58 uint8 = 242
                                                        _ = v58
                                                    } else {
                                                        if true {
                                                            v59 := GenS9{F0: "\n\t", F1: 1, F2: "a"}
                                                            var v60 int64 = 42
                                                            if true {
                                                                v61 := GenS3{F0: fi16}
                                                                var v62 int16 = v51
                                                                _ = v61
                                                                _ = v62
                                                            } else {
                                                                var v63 uint32 = v54
                                                                var v64 int64 = (v60 % ((9223372036854775807) | 1))
                                                                _ = v63
                                                                _ = v64
                                                            }
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                    }
                                                }
                                                v3 = (3.14159 - (v3 + v3))
                                                switch (fi53 | (-(fi16))) {
                                                case 1:
                                                    switch "hello" {
                                                    case "a":
                                                        fmt.Println(v6)
                                                        if true {
                                                            switch "" {
                                                            case "a":
                                                                if ((fi16 ^ 42) < 4294967295) {
                                                                    fmt.Println(v6, v6, v51)
                                                                    v1 = append(v1, uint32(v6))
                                                                }
                                                            default:
                                                                v65 := GenS3{F0: 9223372036854775807}
                                                                if false {
                                                                    var v66 int8 = 1
                                                                    v67 := v5[(("hello" + (("a" + "") + "a")) + "hello")]
                                                                    _ = v66
                                                                    _ = v67
                                                                }
                                                                v54 = 2
                                                                _ = v65
                                                            }
                                                            v6 -= v6
                                                            v5["foo bar"] = v54
                                                            if (!true) {
                                                                v68 := GenS4{F0: -1, F1: 42, F2: fi53}
                                                                _ = v68
                                                            } else {
                                                                v2 = v2
                                                                var v69 uint = 977
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                _ = v69
                                                            }
                                                            switch (fi16 | fi16) {
                                                            case -9223372036854775808:
                                                                var v70 uint64 = 4294967295
                                                                _ = v70
                                                            default:
                                                                var v71 int16 = (v51 | v51)
                                                                _ = v71
                                                            }
                                                        } else {
                                                            var v72 int16 = v51
                                                            v72 += (-(v72))
                                                            _ = v72
                                                        }
                                                        v51 = (v51 & v51)
                                                        if true {
                                                            _ = len(v4)
                                                            var v73 uint32 = (v54 / ((856) | 1))
                                                            if false {
                                                                _ = len(v5)
                                                                v5[("" + (("" + "foo bar") + ""))] = (v54 & (v54 << ((18446744073709551615) & 63)))
                                                            } else {
                                                                var v74 int16 = (v51 ^ (v51 * (v51 ^ v51)))
                                                                v1 = append(v1, 2147483648)
                                                                switch "" {
                                                                case "a":
                                                                    v6 -= 208
                                                                default:
                                                                    _ = len(v5)
                                                                }
                                                                _ = v74
                                                            }
                                                            v6 = 255
                                                            _ = v73
                                                        } else {
                                                            fmt.Println(v6, fi53)
                                                            if true {
                                                                var v75 int8 = -86
                                                                _ = len(v4)
                                                                _ = v75
                                                            }
                                                        }
                                                    case "hello":
                                                        v3 = v3
                                                    default:
                                                        if (42 < 0) {
                                                            v2 = uint(v6)
                                                        } else {
                                                            fmt.Println(v54, v2, v3)
                                                            if true {
                                                                if true {
                                                                    var v76 string = ("\n\t" + "hello")
                                                                    var v77 uint32 = 736
                                                                    if (-875 != 1) {
                                                                        var v78 uint = v2
                                                                        _ = v78
                                                                    } else {
                                                                        fmt.Println(v2, fi53, v51)
                                                                        var v79 uint64 = 4294967295
                                                                        fmt.Println(fi53, v6, fi53)
                                                                        _ = v79
                                                                    }
                                                                    v5[("" + ((v76 + "hello") + ("\n\t" + v76)))] = v77
                                                                    _ = v76
                                                                    _ = v77
                                                                }
                                                                v80 := GenS6{F0: (fi16 * -737), F1: 1000}
                                                                _ = v80
                                                            } else {
                                                                switch "" {
                                                                case "\n\t":
                                                                    v81 := v5["foo bar"]
                                                                    fmt.Println(fi16)
                                                                    _ = v81
                                                                case "foo bar":
                                                                    switch ("" + "\n\t") {
                                                                    case "a":
                                                                        v54 -= v54
                                                                    case "foo bar":
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            v82 := GenS9{F0: "\n\t", F1: 630, F2: "hello"}
                                                                            _ = v82
                                                                        case "\n\t":
                                                                            if true {
                                                                                v83 := GenS4{F0: 2147483647, F1: 811, F2: fi53}
                                                                                if (((v6 * (v6 | v6)) != v6) || ((-1 < 1) || (false && true))) {
                                                                                    v84 := GenS8{F0: 1}
                                                                                    _ = v84
                                                                                }
                                                                                _ = v83
                                                                            }
                                                                        case "hello":
                                                                            v6 += (^(v6))
                                                                        default:
                                                                            v2 = 894
                                                                        }
                                                                        var v85 float32 = v3
                                                                        _ = v85
                                                                    case "":
                                                                        switch "" {
                                                                        case "foo bar":
                                                                            var v86 bool = false
                                                                            _ = v86
                                                                        case "":
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        }
                                                                    default:
                                                                        var v87 uint32 = v54
                                                                        _ = v87
                                                                    }
                                                                }
                                                                var v88 uint = v2
                                                                _ = v88
                                                            }
                                                            var v89 int8 = 42
                                                            _ = v89
                                                        }
                                                    }
                                                case -549:
                                                    if (fi16 <= -9223372036854775808) {
                                                        v90 := GenS10{F0: (0 != -128)}
                                                        _ = v90
                                                    } else {
                                                        var v91 int16 = -405
                                                        var v92 float64 = (0.0 / (1e10 + 3.14159))
                                                        _ = v91
                                                        _ = v92
                                                    }
                                                default:
                                                    if (909 > v2) {
                                                        v6 = 1
                                                    } else {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        v93 := v5["a"]
                                                        var v94 int16 = v51
                                                        _ = v93
                                                        _ = v94
                                                    }
                                                }
                                                fmt.Println(v3, fi53, v54)
                                            } else {
                                                v2 -= v2
                                                v3 -= v3
                                            }
                                            if true {
                                                if true {
                                                    if true {
                                                        fmt.Println(v6, v6, v54)
                                                        if ((v2 >= v2) || ((v51 ^ int16(v54)) > v51)) {
                                                            _ = len(v5)
                                                            v54 -= v54
                                                        } else {
                                                            var v95 int32 = 1
                                                            v95 -= v95
                                                            switch ("hello" + "hello") {
                                                            case "foo bar":
                                                                v96 := GenS2{F0: (v95 / (((v95 & v95)) | 1)), F1: 18446744073709551615, F2: v6}
                                                                _ = v96
                                                            case "a":
                                                                var v97 int16 = -32768
                                                                _ = v97
                                                            default:
                                                                if false {
                                                                    var v98 int = (fi53 - fi16)
                                                                    var v99 float32 = ((0.0 / 3.14159) / ((float32(0.0) / 1.0) + float32(3.14159)))
                                                                    _ = v98
                                                                    _ = v99
                                                                }
                                                            }
                                                            _ = v95
                                                        }
                                                    }
                                                    delete(v5, "hello")
                                                    v100 := GenS4{F0: 1, F1: 0, F2: fi53}
                                                    _ = len(v5)
                                                    _ = v100
                                                }
                                                switch "\n\t" {
                                                case "foo bar":
                                                    switch "\n\t" {
                                                    case "":
                                                        v54 -= v54
                                                    case "a":
                                                        if (false || true) {
                                                            v101 := GenS4{F0: int32(v2), F1: 42, F2: (fi53 & -1)}
                                                            if (18446744073709551615 >= 2) {
                                                                if (0 >= 255) {
                                                                    var v102 uint8 = 127
                                                                    var v103 uint = v2
                                                                    _ = v102
                                                                    _ = v103
                                                                } else {
                                                                    v54 += v54
                                                                    if (!false) {
                                                                        if (((v51 * (v51 ^ (-(v51)))) >= (v51 % ((-32768) | 1))) || (1 < -9223372036854775808)) {
                                                                            if ((!(-2147483648 == -705)) && true) {
                                                                                switch fi53 {
                                                                                case 1:
                                                                                    if (!false) {
                                                                                        _ = len(v1)
                                                                                        if (789 >= 613) {
                                                                                            v54 += 1
                                                                                            var v104 string = ""
                                                                                            v105 := GenS11{F0: -128}
                                                                                            v106 := v5[(v104 + v104)]
                                                                                            if false {
                                                                                                if (!false) {
                                                                                                    v107 := v5[v104]
                                                                                                    var v108 bool = (v6 >= (v6 & (^(v6))))
                                                                                                    _ = v107
                                                                                                    _ = v108
                                                                                                }
                                                                                                var v109 float32 = (1.5 + v3)
                                                                                                _ = v109
                                                                                            } else {
                                                                                                var v110 int16 = -303
                                                                                                if ((v106 >> ((uint(v54)) & 63)) > 991) {
                                                                                                    v110 = v51
                                                                                                    _ = len(v5)
                                                                                                    v104 = ((v104 + "foo bar") + v104)
                                                                                                    if (!true) {
                                                                                                        switch fi53 {
                                                                                                        case 1:
                                                                                                            v6 += v6
                                                                                                        }
                                                                                                    } else {
                                                                                                        if len(v4) >= 2 {
                                                                                                            v4 = v4[1:len(v4)]
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                if (v2 <= 4294967295) {
                                                                                                    v104 = ""
                                                                                                }
                                                                                                _ = v110
                                                                                            }
                                                                                            _ = v104
                                                                                            _ = v105
                                                                                            _ = v106
                                                                                        }
                                                                                        v54 = v54
                                                                                    }
                                                                                case 32767:
                                                                                    if false {
                                                                                        fmt.Println(v2, v6, v54)
                                                                                    } else {
                                                                                        if false {
                                                                                            if ((0.0 - (0.1 - 0.1)) >= (((float64(-2.25) + 3.14159) - 0.0) - float64(((v3 - (1.5 / v3)) * (((v3 / v3) - (v3 / 1e10)) - 0.0))))) {
                                                                                                v2 = v2
                                                                                            } else {
                                                                                                if (0.0 < ((0.0 * (-2.25 - float64(v3))) / 1.0)) {
                                                                                                    delete(v5, (("\n\t" + "a") + "a"))
                                                                                                    fmt.Println(v3)
                                                                                                    if (v6 != (v6 & 128)) {
                                                                                                        switch -2147483648 {
                                                                                                        case 32767:
                                                                                                            var v111 uint16 = 32767
                                                                                                            _ = v111
                                                                                                        case 0:
                                                                                                            switch ("\n\t" + "") {
                                                                                                            case "a":
                                                                                                                fmt.Println(v3, v51)
                                                                                                            case "":
                                                                                                                if (18446744073709551615 != uint64(v2)) {
                                                                                                                    v51 = (v51 / ((v51) | 1))
                                                                                                                    v112 := GenS1{F0: 601, F1: 9223372036854775807, F2: 1}
                                                                                                                    _ = v112
                                                                                                                } else {
                                                                                                                    if false {
                                                                                                                        v3 = (v3 - 3.14159)
                                                                                                                        if true {
                                                                                                                            var v113 int16 = int16(v6)
                                                                                                                            var v114 bool = ((((0.1 + ((0.1 * 1e10) / (1.5 / (float64(0.0) + ((float64(float32(((float64(v3) + -2.25) + (-2.25 / 0.1)))) + (0.1 / float64(v3))) * -2.25))))) < 0.0) && true) && ("" < "\n\t"))
                                                                                                                            v115 := GenS11{F0: 127}
                                                                                                                            _ = v113
                                                                                                                            _ = v114
                                                                                                                            _ = v115
                                                                                                                        }
                                                                                                                        v2 -= v2
                                                                                                                    }
                                                                                                                    v4 = append(v4, ("" + "hello"))
                                                                                                                    _ = len(v5)
                                                                                                                    _ = len(v5)
                                                                                                                    var v116 int16 = (v51 * v51)
                                                                                                                    v117 := GenS12{F0: (^(fi16)), F1: 42, F2: int16(fi53)}
                                                                                                                    _ = v116
                                                                                                                    _ = v117
                                                                                                                }
                                                                                                            default:
                                                                                                                v118 := GenS6{F0: (-(fi53)), F1: v51}
                                                                                                                _ = v118
                                                                                                            }
                                                                                                            v119 := GenS6{F0: fi53, F1: v51}
                                                                                                            v2 += 0
                                                                                                            _ = v119
                                                                                                        case 127:
                                                                                                            if (!false) {
                                                                                                                v51 -= int16(v2)
                                                                                                            }
                                                                                                            _ = len(v5)
                                                                                                        }
                                                                                                        v5[""] = v54
                                                                                                    }
                                                                                                    v120 := GenS13{F0: 0, F1: (^(v54))}
                                                                                                    v54 += v54
                                                                                                    _ = v120
                                                                                                } else {
                                                                                                    var v121 float32 = v3
                                                                                                    v122 := GenS14{F0: (!false), F1: (("hello" + "a") + ""), F2: int64(v51)}
                                                                                                    _ = v121
                                                                                                    _ = v122
                                                                                                }
                                                                                            }
                                                                                            v6 -= 209
                                                                                            switch (fi53 + (^(fi53))) {
                                                                                            case 1:
                                                                                                if false {
                                                                                                    v54 += v54
                                                                                                    if ("a" >= ("foo bar" + "a")) {
                                                                                                        v4 = append(v4, "hello")
                                                                                                        _ = len(v5)
                                                                                                    } else {
                                                                                                        v3 = (v3 + (0.0 / v3))
                                                                                                        var v123 int32 = 0
                                                                                                        if ((^(v54)) == (^(v54))) {
                                                                                                            v124 := GenS1{F0: 427, F1: -1, F2: 0}
                                                                                                            var v125 int16 = v51
                                                                                                            _ = v124
                                                                                                            _ = v125
                                                                                                        }
                                                                                                        _ = v123
                                                                                                    }
                                                                                                    v126 := GenS9{F0: "", F1: 2, F2: (("hello" + "foo bar") + ((((("" + ("hello" + ("foo bar" + "foo bar"))) + "hello") + "a") + "") + ""))}
                                                                                                    _ = v126
                                                                                                }
                                                                                            default:
                                                                                                if true {
                                                                                                    v51 += v51
                                                                                                    if ((2 < 712) || (false || true)) {
                                                                                                        var v127 float32 = 1.5
                                                                                                        _ = v127
                                                                                                    } else {
                                                                                                        v51 -= v51
                                                                                                        v4 = append(v4, "hello")
                                                                                                    }
                                                                                                } else {
                                                                                                    v128 := GenS15{F0: 2147483647}
                                                                                                    _ = v128
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    switch (fi53 * fi16) {
                                                                                    case -1:
                                                                                        v51 = (v51 ^ (v51 | v51))
                                                                                    case -2147483648:
                                                                                        if ((uint32(v2) >= 4294967295) && false) {
                                                                                            var v129 float64 = 1e10
                                                                                            var v130 string = ((("foo bar" + ("foo bar" + "hello")) + ("hello" + "foo bar")) + "a")
                                                                                            v131 := GenS12{F0: fi16, F1: -128, F2: v51}
                                                                                            _ = v129
                                                                                            _ = v130
                                                                                            _ = v131
                                                                                        } else {
                                                                                            v132 := GenS16{F0: false, F1: 42, F2: (-(fi53))}
                                                                                            var v133 int64 = 556
                                                                                            fmt.Println(v54, v2, fi53)
                                                                                            _ = v132
                                                                                            _ = v133
                                                                                        }
                                                                                    case 769:
                                                                                        v2 -= (v2 << (((v2 / ((v2) | 1))) & 63))
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v134 int8 = 127
                                                                                var v135 float32 = v3
                                                                                switch "hello" {
                                                                                case "hello":
                                                                                    v136 := v5[""]
                                                                                    _ = v136
                                                                                case "":
                                                                                    fmt.Println(fi16, v54)
                                                                                default:
                                                                                    if ("a" != "foo bar") {
                                                                                        fmt.Println(v54)
                                                                                    } else {
                                                                                        if (!false) {
                                                                                            v137 := GenS17{F0: v2, F1: genFn0(1)}
                                                                                            var v138 uint8 = 1
                                                                                            _ = v137
                                                                                            _ = v138
                                                                                        } else {
                                                                                            v139 := GenS10{F0: (v2 >= (v2 | (v2 / (((v2 | 921)) | 1))))}
                                                                                            if false {
                                                                                                if false {
                                                                                                    v6 = (v6 & 255)
                                                                                                }
                                                                                            }
                                                                                            var v140 int16 = 1
                                                                                            _ = v139
                                                                                            _ = v140
                                                                                        }
                                                                                        fmt.Println(fi16)
                                                                                        var v141 int64 = 106
                                                                                        _ = v141
                                                                                    }
                                                                                }
                                                                                _ = v134
                                                                                _ = v135
                                                                            }
                                                                        } else {
                                                                            var v142 uint32 = v54
                                                                            v142 = v54
                                                                            _ = v142
                                                                        }
                                                                        v143 := GenS1{F0: 2, F1: 697, F2: 32767}
                                                                        _ = v143
                                                                    } else {
                                                                        fmt.Println(v51, fi16)
                                                                    }
                                                                }
                                                            } else {
                                                                v144 := GenS18{F0: v6, F1: 2}
                                                                _ = v144
                                                            }
                                                            if false {
                                                                var v145 int64 = 42
                                                                _ = v145
                                                            }
                                                            var v146 int32 = 486
                                                            _ = v101
                                                            _ = v146
                                                        }
                                                    }
                                                case "\n\t":
                                                    v3 += float32(((1e10 + 0.1) + -2.25))
                                                case "":
                                                    if true {
                                                        fmt.Println(fi16, v2)
                                                        v147 := GenS19{F0: v51, F1: 838}
                                                        var v148 uint64 = 77
                                                        _ = v147
                                                        _ = v148
                                                    }
                                                }
                                                v3 = 1e10
                                                v149 := GenS20{F0: false, F1: 849}
                                                _ = v149
                                            } else {
                                                switch (("hello" + "\n\t") + "hello") {
                                                case "foo bar":
                                                    fmt.Println(v51)
                                                case "a":
                                                    v150 := GenS14{F0: (!false), F1: "a", F2: 42}
                                                    _ = v150
                                                case "hello":
                                                    if false {
                                                        v54 = v54
                                                        fmt.Println(v3, v6, v51)
                                                    } else {
                                                        if (!false) {
                                                            v6 = v6
                                                            v151 := GenS4{F0: 0, F1: 2147483647, F2: (-(fi53))}
                                                            _ = v151
                                                        }
                                                        v54 = (v54 - uint32(v6))
                                                        if (v3 == v3) {
                                                            v2 -= 334
                                                            var v152 uint = v2
                                                            var v153 int64 = 0
                                                            fmt.Println(fi16)
                                                            _ = v152
                                                            _ = v153
                                                        } else {
                                                            if true {
                                                                v154 := v5["foo bar"]
                                                                v155 := GenS3{F0: fi53}
                                                                _ = v154
                                                                _ = v155
                                                            } else {
                                                                v156 := GenS17{F0: v2, F1: 4294967295}
                                                                _ = v156
                                                            }
                                                            v2 -= v2
                                                        }
                                                    }
                                                default:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                            }
                                        }
                                        _ = fi53
                                        _ = v54
                                    }
                                }
                                v157 := v5["foo bar"]
                                _ = v157
                                _ = v51
                            }
                            v5[""] = 4294967295
                        }
                        v3 = (3.14159 - ((((3.14159 * (0.0 * 3.14159)) / (v3 - float32(((0.0 / -2.25) + float64(v3))))) * v3) + v3))
                        for fi158 := 0; fi158 < 13; fi158++ {
                            var v159 int64 = -1
                            _ = fi158
                            _ = v159
                        }
                        _ = fi16
                    }
                    if true {
                        for fi160 := 0; fi160 < 7; fi160++ {
                            if false {
                                if false {
                                    if ((!(!((-1 >= 398) || (9223372036854775807 > 1)))) && (319 <= 255)) {
                                        v6 = v6
                                    }
                                } else {
                                    v161 := v5["\n\t"]
                                    if true {
                                        for fi162 := 0; fi162 < 7; fi162++ {
                                            fmt.Println(v3, v3, v161)
                                            var v163 uint16 = 65535
                                            _ = fi162
                                            _ = v163
                                        }
                                    }
                                    for fi164 := 0; fi164 < 6; fi164++ {
                                        switch "" {
                                        case "":
                                            if true {
                                                _ = len(v4)
                                            }
                                        case "\n\t":
                                            if true {
                                                switch 127 {
                                                case 54:
                                                    switch ("foo bar" + "foo bar") {
                                                    case "a":
                                                        v3 = ((v3 / 1e10) - (((v3 - v3) - -2.25) + v3))
                                                    case "hello":
                                                        switch (fi160 & (^(fi164))) {
                                                        case 770:
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        case 0:
                                                            v165 := GenS14{F0: true, F1: "\n\t", F2: 9223372036854775807}
                                                            _ = v165
                                                        }
                                                    default:
                                                        if (v6 != (v6 % ((255) | 1))) {
                                                            _ = len(v5)
                                                            switch (-(fi164)) {
                                                            case 0:
                                                                _ = len(v5)
                                                            default:
                                                                v166 := GenS13{F0: 4294967295, F1: (v161 + (^(v161)))}
                                                                _ = v166
                                                            }
                                                        }
                                                    }
                                                case 4294967295:
                                                    if (1 > 240) {
                                                        v2 += v2
                                                        v167 := GenS21{F0: 65535, F1: v2}
                                                        var v168 int16 = int16(fi164)
                                                        _ = v167
                                                        _ = v168
                                                    }
                                                }
                                                v2 = v2
                                            }
                                        default:
                                            var v169 uint32 = v161
                                            _ = v169
                                        }
                                        var v170 float32 = v3
                                        fmt.Println(v3, v161, v170)
                                        _ = fi164
                                        _ = v170
                                    }
                                    _ = v161
                                }
                            } else {
                                if (!(!(!true))) {
                                    v171 := GenS22{F0: -1, F1: 0}
                                    _ = v171
                                }
                                var v172 bool = (true || (1 >= 119))
                                v173 := GenS18{F0: v6, F1: 1}
                                _ = v172
                                _ = v173
                            }
                            for fi174 := 0; fi174 < 11; fi174++ {
                                _ = len(v1)
                                if ((v6 + v6) > v6) {
                                    var v175 uint = (^(v2))
                                    _ = v175
                                } else {
                                    v176 := GenS23{F0: 255, F1: 2147483647}
                                    v177 := GenS19{F0: -372, F1: 85}
                                    _ = v176
                                    _ = v177
                                }
                                _ = fi174
                            }
                            _ = fi160
                        }
                    }
                }
            }
            delete(v5, "\n\t")
            v178 := GenS9{F0: "\n\t", F1: 42, F2: "\n\t"}
            _ = v178
        }
    default:
        if (2147483647 < 2147483647) {
            v179 := GenS24{F0: v6}
            for fi180 := 0; fi180 < 18; fi180++ {
                for fi181 := 0; fi181 < 3; fi181++ {
                    if (32767 <= 0) {
                        var v182 uint8 = (^(v6))
                        var v183 uint64 = 0
                        var v184 float32 = ((v3 + 0.1) + float32(1e10))
                        _ = v182
                        _ = v183
                        _ = v184
                    }
                    fmt.Println(v3, v6, v3)
                    var v185 float32 = float32((1.5 + 1e10))
                    v186 := GenS25{F0: v6, F1: -50, F2: fi180}
                    v3 = ((v3 * v185) + v3)
                    _ = fi181
                    _ = v185
                    _ = v186
                }
                _ = fi180
            }
            if false {
                var v187 int8 = 1
                v188 := GenS26{F0: true, F1: -9223372036854775808}
                _ = v187
                _ = v188
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if false {
                    var v189 float32 = v3
                    _ = v189
                } else {
                    var v190 float64 = 3.14159
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v190
                }
                if false {
                    for fi191 := 0; fi191 < 8; fi191++ {
                        for fi192 := 0; fi192 < 11; fi192++ {
                            v193 := GenS27{F0: (v6 * v6), F1: -1, F2: 125}
                            _ = fi192
                            _ = v193
                        }
                        _ = fi191
                    }
                    v4 = append(v4, "hello")
                    v194 := GenS25{F0: v6, F1: 0, F2: 4294967295}
                    _ = v194
                }
            }
            _ = v179
        } else {
            v2 += (v2 ^ v2)
            v195 := GenS8{F0: -570}
            var v196 uint8 = (v6 ^ v6)
            _ = v195
            _ = v196
        }
    }
    if false {
        for fi197 := 0; fi197 < 8; fi197++ {
            v2 = v2
            _ = fi197
        }
    } else {
        if true {
            v3 = v3
            v6 = (^(v6))
        }
    }
    v3 += 1.5
    v3 = v3
    for fi198 := 0; fi198 < 13; fi198++ {
        v4 = append(v4, "\n\t")
        switch fi198 {
        case -1:
            if true {
                var v199 int16 = 42
                _ = v199
            } else {
                var v200 uint8 = 42
                _ = v200
            }
        }
        v6 += (v6 & 128)
        _ = fi198
    }
    v201 := GenS11{F0: -128}
    v3 += v3
    v202 := GenS28{F0: genFn0(1), F1: true, F2: false}
    var v203 float64 = 1e10
    var v204 int = 127
    switch (v204 << (((v2 ^ (v2 / (((v2 % ((v2) | 1))) | 1)))) & 63)) {
    case 32767:
        var v205 uint8 = v6
        _ = v205
    }
    for fi206 := 0; fi206 < 18; fi206++ {
        var v207 uint = v2
        _ = fi206
        _ = v207
    }
    fmt.Println(v2, v6)
    if false {
        var v208 uint32 = 0
        _ = v208
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if (!false) {
        for fi209 := 0; fi209 < 3; fi209++ {
            for fi210 := 0; fi210 < 10; fi210++ {
                if (0.1 <= v3) {
                    v204 += (fi210 + (-(fi209)))
                    v211 := v5["\n\t"]
                    if false {
                        if false {
                            fmt.Println(v203)
                        }
                        var v212 int = (-(fi210))
                        v213 := GenS8{F0: 1}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v212
                        _ = v213
                    }
                    _ = v211
                } else {
                    v6 -= v6
                    var v214 int16 = -32768
                    var v215 float64 = (0.1 - float64(((float32(v203) * (1.5 - -2.25)) / 1.5)))
                    _ = v214
                    _ = v215
                }
                v6 = (v6 + v6)
                _ = fi210
            }
            if (2 > 42) {
                v6 = 192
            }
            _ = fi209
        }
        var v216 string = ""
        _ = v216
    }
    v1 = append(v1, uint32(v2))
    if (v6 > (v6 << (((v2 * 255)) & 63))) {
        var v217 uint16 = 527
        switch v204 {
        case 32768:
            v4 = append(v4, ("hello" + ""))
        case 128:
            switch (("hello" + (("\n\t" + "foo bar") + ((("foo bar" + "") + "a") + "hello"))) + ((("foo bar" + "\n\t") + "") + "")) {
            case "a":
                for fi218 := 0; fi218 < 4; fi218++ {
                    v219 := GenS29{F0: v2}
                    _ = fi218
                    _ = v219
                }
            case "\n\t":
                for fi220 := 0; fi220 < 17; fi220++ {
                    if len(v1) > 0 {
                        v1[0] = 2
                    }
                    _ = fi220
                }
            case "hello":
                switch 0 {
                case -312:
                    switch ("a" + "\n\t") {
                    case "hello":
                        fmt.Println(v217)
                    case "foo bar":
                        v6 = v6
                    case "":
                        v221 := GenS3{F0: -190}
                        _ = v221
                    default:
                        for fi222 := 0; fi222 < 6; fi222++ {
                            v1 = append(v1, 4294967295)
                            _ = fi222
                        }
                    }
                default:
                    switch v204 {
                    case 398:
                        var v223 int = (v204 | -9223372036854775808)
                        _ = v223
                    case 4294967295:
                        if true {
                            for fi224 := 0; fi224 < 15; fi224++ {
                                v204 += v204
                                v217 = v217
                                _ = fi224
                            }
                            v3 += (1e10 / (v3 + 0.1))
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(v217)
                            default:
                                for fi225 := 0; fi225 < 7; fi225++ {
                                    if true {
                                        var v226 uint = v2
                                        v227 := GenS14{F0: true, F1: "foo bar", F2: 42}
                                        _ = v226
                                        _ = v227
                                    } else {
                                        v217 = (v217 >> (((v2 * 4294967295)) & 63))
                                        v2 = 18446744073709551615
                                    }
                                    if true {
                                        switch 158 {
                                        case 128:
                                            if (2147483647 > 0) {
                                                var v228 uint32 = 286
                                                _ = v228
                                            }
                                        case 4294967295:
                                            var v229 float64 = v203
                                            _ = v229
                                        case 0:
                                            if true {
                                                delete(v5, "a")
                                                v230 := GenS2{F0: 42, F1: 4294967295, F2: v6}
                                                v6 -= (v6 & (^(v6)))
                                                _ = v230
                                            } else {
                                                v231 := GenS16{F0: ((v6 << (((v2 / (((v2 % ((4294967295) | 1))) | 1))) & 63)) <= v6), F1: -1, F2: 0}
                                                fmt.Println(v6, v217, v217)
                                                var v232 int32 = -279
                                                _ = v231
                                                _ = v232
                                            }
                                        }
                                    }
                                    if false {
                                        var v233 int = (fi225 & v204)
                                        _ = v233
                                    } else {
                                        fmt.Println(v3, v6)
                                        var v234 uint = uint(v6)
                                        _ = v234
                                    }
                                    _ = fi225
                                }
                            }
                        }
                    default:
                        for fi235 := 0; fi235 < 18; fi235++ {
                            v236 := v5["a"]
                            _ = fi235
                            _ = v236
                        }
                    }
                }
            default:
                for fi237 := 0; fi237 < 5; fi237++ {
                    v1 = append(v1, 502)
                    _ = fi237
                }
            }
        default:
            var v238 int = (v204 | -446)
            _ = v238
        }
        _ = v217
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v203:", v203)
    fmt.Println("v204:", v204)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v201
    _ = v202
    _ = v203
    _ = v204
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
