package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int8; F2 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 string; F1 uint16; F2 string }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint32; F1 uint; F2 uint }
type GenS11 struct { F0 string; F1 int16; F2 uint64 }
type GenS12 struct { F0 int8; F1 int }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint8; F1 int8 }
type GenS16 struct { F0 int8; F1 int64 }
type GenS17 struct { F0 int64; F1 int16; F2 int64 }
type GenS18 struct { F0 uint64; F1 int16 }
type GenS19 struct { F0 string; F1 int; F2 uint16 }
type GenS20 struct { F0 uint64 }
type GenS21 struct { F0 uint8; F1 int }
type GenS22 struct { F0 int }
type GenS23 struct { F0 uint64; F1 uint }
type GenS24 struct { F0 uint32; F1 uint8; F2 int }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]bool)
    v5 := make(map[string]string)
    var v6 uint32 = 225
    v3 = -710
    var v7 string = "hello"
    v3 = (v3 * (^(v3)))
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(fi8)
        var v9 float64 = (0.0 + float64(0.0))
        var v10 int16 = -813
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 100}
    v3 = 9223372036854775807
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 uint = 4294967295
        v14 := GenS1{F0: 100}
        if false {
            v15 := GenS2{F0: int8(v3), F1: 56, F2: -1}
            if true {
                if false {
                    if true {
                        var v16 int16 = -32768
                        if false {
                            fmt.Println(fi12)
                        } else {
                            var v17 int64 = (-(v3))
                            var v18 int64 = (v17 & (v17 * v3))
                            _ = v17
                            _ = v18
                        }
                        v3 -= v3
                        _ = v16
                    } else {
                        v13 -= v13
                        var v19 int8 = -122
                        if ((v19 >> ((v13) & 63)) == -95) {
                            for fi20 := 0; fi20 < 12; fi20++ {
                                switch v7 {
                                case "\n\t":
                                    if true {
                                        if (false && (0.0 < ((0.1 + 0.0) / 1e10))) {
                                            v13 = (^(v13))
                                            v21 := GenS2{F0: v19, F1: (-(v19)), F2: v19}
                                            switch v7 {
                                            case "\n\t":
                                                if false {
                                                    fmt.Println(v19, v19, v19)
                                                    v22 := GenS3{F0: (fi12 ^ fi20)}
                                                    _ = len(v5)
                                                    _ = v22
                                                }
                                            case "a":
                                                fmt.Println(v3, v3)
                                            case "foo bar":
                                                var v23 uint16 = 65535
                                                _ = v23
                                            default:
                                                v24 := GenS2{F0: v19, F1: (v19 << (((v13 / (((v13 >> (((v13 ^ v13)) & 63))) | 1))) & 63)), F2: v19}
                                                _ = v24
                                            }
                                            _ = v21
                                        } else {
                                            v3 = (v3 | v3)
                                            v3 -= (v3 ^ -269)
                                            var v25 int32 = -1
                                            _ = v25
                                        }
                                        v13 -= (v13 << ((v13) & 63))
                                    } else {
                                        v13 = v13
                                        if false {
                                            v2[(fi20 + fi12)] = (-1 > -32768)
                                            var v26 int64 = int64(v13)
                                            v27 := GenS4{F0: -630, F1: int32(v13), F2: 1}
                                            _ = v26
                                            _ = v27
                                        } else {
                                            var v28 float32 = 0.0
                                            var v29 uint32 = v6
                                            fmt.Println(v29, fi20, v7)
                                            _ = v28
                                            _ = v29
                                        }
                                        if (v6 == v6) {
                                            var v30 bool = (!(!(985 <= 480)))
                                            switch (("" + "") + "a") {
                                            case "\n\t":
                                                if true {
                                                    var v31 bool = v30
                                                    _ = v31
                                                }
                                            case "hello":
                                                switch "foo bar" {
                                                case "foo bar":
                                                    switch 2147483647 {
                                                    case 298:
                                                        fmt.Println(v6)
                                                    default:
                                                        if (v6 < (v6 + v6)) {
                                                            fmt.Println(v6)
                                                            if (999 != 557) {
                                                                v30 = false
                                                            }
                                                            var v32 bool = (!(true && v30))
                                                            _ = v32
                                                        } else {
                                                            var v33 uint8 = 1
                                                            v6 = (v6 % ((2147483647) | 1))
                                                            v3 -= 42
                                                            var v34 uint8 = 127
                                                            v6 -= (v6 * 2147483647)
                                                            _ = v33
                                                            _ = v34
                                                        }
                                                    }
                                                    v35 := GenS2{F0: -33, F1: (v19 % (((-(v19))) | 1)), F2: v19}
                                                    switch (fi12 / ((fi20) | 1)) {
                                                    case -1:
                                                        if (-2.25 == (0.1 + float64(-2.25))) {
                                                            v36 := GenS1{F0: (v19 << ((v13) & 63))}
                                                            if (42 >= 477) {
                                                                if ((0.1 / ((1e10 - 0.0) * -2.25)) > 1.5) {
                                                                    var v37 float32 = 0.0
                                                                    _ = v37
                                                                }
                                                                if false {
                                                                    if (v30 || (!(v30 || v30))) {
                                                                        v13 = (v13 + (^(v13)))
                                                                        v38 := []int{(fi12 + fi12)}
                                                                        if v30 {
                                                                            fmt.Println(v19, fi12)
                                                                            if v30 {
                                                                                v39 := GenS5{F0: 42, F1: 1}
                                                                                _ = v39
                                                                            } else {
                                                                                if (9223372036854775807 >= 1) {
                                                                                    _ = len(v1)
                                                                                    v3 -= v3
                                                                                    var v40 uint32 = v6
                                                                                    if v30 {
                                                                                        var v41 int16 = 24
                                                                                        _ = v41
                                                                                    }
                                                                                    switch "foo bar" {
                                                                                    case "\n\t":
                                                                                        if len(v38) > 0 {
                                                                                            v38[0] = (fi12 / ((int(v40)) | 1))
                                                                                        }
                                                                                    default:
                                                                                        v13 = (v13 ^ v13)
                                                                                    }
                                                                                    _ = v40
                                                                                }
                                                                                var v42 bool = v30
                                                                                _ = v42
                                                                            }
                                                                        } else {
                                                                            if (fi12 <= (fi12 >> ((v13) & 63))) {
                                                                                v19 -= v19
                                                                                var v43 int16 = int16(fi20)
                                                                                if (v13 == 2) {
                                                                                    v3 = v3
                                                                                    v44 := GenS3{F0: fi12}
                                                                                    if v30 {
                                                                                        v3 = int64(v43)
                                                                                    }
                                                                                    _ = v44
                                                                                }
                                                                                _ = v43
                                                                            } else {
                                                                                v6 = v6
                                                                                v45 := GenS6{F0: fi20, F1: (-(v3))}
                                                                                _ = v45
                                                                            }
                                                                            var v46 uint = uint(v6)
                                                                            if false {
                                                                                var v47 uint16 = 585
                                                                                var v48 bool = v30
                                                                                _ = v47
                                                                                _ = v48
                                                                            }
                                                                            if len(v38) > 0 {
                                                                                v38[0] = 2147483647
                                                                            }
                                                                            _ = v46
                                                                        }
                                                                        fmt.Println(v7, v13, fi12)
                                                                        _ = v38
                                                                    } else {
                                                                        v30 = (!((v6 | v6) < v6))
                                                                        var v49 int8 = v19
                                                                        if (!(1 <= 255)) {
                                                                            switch (("\n\t" + "a") + ("hello" + v7)) {
                                                                            case "foo bar":
                                                                                var v50 int64 = (^(v3))
                                                                                _ = v50
                                                                            case "":
                                                                                v51 := v5[(v7 + "a")]
                                                                                _ = v51
                                                                            }
                                                                        } else {
                                                                            if v30 {
                                                                                var v52 int64 = v3
                                                                                _ = v52
                                                                            }
                                                                            var v53 uint8 = 127
                                                                            _ = len(v5)
                                                                            v13 = (^(v13))
                                                                            switch 0 {
                                                                            case 709:
                                                                                switch (fi12 * -9223372036854775808) {
                                                                                case -1:
                                                                                    v54 := GenS7{F0: (v7 + v7), F1: 806, F2: ""}
                                                                                    _ = v54
                                                                                default:
                                                                                    var v55 int64 = (v3 << (((v13 * 18446744073709551615)) & 63))
                                                                                    _ = v55
                                                                                }
                                                                            case -662:
                                                                                var v56 bool = (!v30)
                                                                                _ = v56
                                                                            }
                                                                            fmt.Println(v3, v53, v3)
                                                                            _ = v53
                                                                        }
                                                                        _ = v49
                                                                    }
                                                                } else {
                                                                    v19 -= (v19 * (v19 >> (((^(v13))) & 63)))
                                                                    switch fi12 {
                                                                    case 32767:
                                                                        v57 := []int{(fi20 * (fi20 ^ int(v3))), int(v13)}
                                                                        _ = v57
                                                                    case -1:
                                                                        v58 := GenS5{F0: 455, F1: 229}
                                                                        delete(v4, fi20)
                                                                        if (!v30) {
                                                                            if (v30 && false) {
                                                                                fmt.Println(v13)
                                                                                v59 := GenS8{F0: 172}
                                                                                var v60 uint32 = 1
                                                                                var v61 int8 = 42
                                                                                _ = v59
                                                                                _ = v60
                                                                                _ = v61
                                                                            }
                                                                        } else {
                                                                            v62 := []int{}
                                                                            switch "a" {
                                                                            case "a":
                                                                                var v63 int32 = 1
                                                                                v13 -= 42
                                                                                v64 := GenS7{F0: v7, F1: 2, F2: (v7 + v7)}
                                                                                _ = v63
                                                                                _ = v64
                                                                            case "":
                                                                                switch v7 {
                                                                                case "a":
                                                                                    switch v7 {
                                                                                    case "\n\t":
                                                                                        v6 += v6
                                                                                        var v65 float64 = (((0.1 / 3.14159) / (0.1 + (0.1 - -2.25))) - (0.0 - 1.5))
                                                                                        _ = v65
                                                                                    }
                                                                                case "foo bar":
                                                                                    var v66 int16 = 32767
                                                                                    _ = v66
                                                                                default:
                                                                                    if true {
                                                                                        var v67 int8 = (v19 | v19)
                                                                                        _ = v67
                                                                                    }
                                                                                }
                                                                            default:
                                                                                if v30 {
                                                                                    switch (fi20 | 42) {
                                                                                    case 32768:
                                                                                        delete(v1, (v7 + ""))
                                                                                    case 536:
                                                                                        v30 = (0 == -2147483648)
                                                                                    default:
                                                                                        v68 := GenS9{F0: (v3 * (v3 | (v3 % (((v3 & v3)) | 1))))}
                                                                                        _ = v68
                                                                                    }
                                                                                    if false {
                                                                                        fmt.Println(v19)
                                                                                    } else {
                                                                                        var v69 uint16 = 2
                                                                                        v70 := GenS10{F0: 2147483648, F1: 174, F2: 0}
                                                                                        if len(v62) > 0 {
                                                                                            v62[0] = (fi12 * (fi12 - (fi20 ^ fi20)))
                                                                                        }
                                                                                        _ = v69
                                                                                        _ = v70
                                                                                    }
                                                                                }
                                                                            }
                                                                            fmt.Println(v19)
                                                                            var v71 uint32 = 77
                                                                            _ = v62
                                                                            _ = v71
                                                                        }
                                                                        _ = v58
                                                                    case 9223372036854775807:
                                                                        v7 = ""
                                                                    }
                                                                }
                                                                v5[""] = (("hello" + v7) + v7)
                                                            } else {
                                                                fmt.Println(v6, fi12)
                                                                v7 = v7
                                                                var v72 int8 = 100
                                                                v72 = (-(v19))
                                                                if (((v3 >= v3) && (!v30)) && ((-2.25 != 1e10) || false)) {
                                                                    if true {
                                                                        if (128 != 128) {
                                                                            var v73 int64 = (v3 / (((v3 * v3)) | 1))
                                                                            switch v7 {
                                                                            case "foo bar":
                                                                                v74 := GenS8{F0: 0}
                                                                                _ = v74
                                                                            case "a":
                                                                                var v75 uint16 = 65535
                                                                                _ = v75
                                                                            case "":
                                                                                if (true || (4294967295 == 42)) {
                                                                                    v76 := GenS11{F0: (("\n\t" + (v7 + v7)) + "hello"), F1: 1, F2: 4294967295}
                                                                                    var v77 int16 = 32767
                                                                                    _ = v76
                                                                                    _ = v77
                                                                                }
                                                                            default:
                                                                                if (("foo bar" > (v7 + v7)) && v30) {
                                                                                    v13 += (v13 ^ (v13 / ((v13) | 1)))
                                                                                }
                                                                            }
                                                                            _ = v73
                                                                        } else {
                                                                            var v78 uint64 = 273
                                                                            _ = v78
                                                                        }
                                                                        v79 := []int{fi20}
                                                                        v13 = v13
                                                                        _ = v79
                                                                    } else {
                                                                        var v80 int8 = (v72 % (((^(v19))) | 1))
                                                                        v81 := GenS12{F0: v80, F1: fi20}
                                                                        v82 := []int{fi20, (fi12 * 42)}
                                                                        _ = v80
                                                                        _ = v81
                                                                        _ = v82
                                                                    }
                                                                    if (v30 && ((v3 + 341) != (v3 * 42))) {
                                                                        v1[v7] = fi20
                                                                        if true {
                                                                            fmt.Println(v7, fi20)
                                                                        } else {
                                                                            _ = len(v1)
                                                                            v3 += v3
                                                                            v72 -= v72
                                                                        }
                                                                        v83 := []int{767, 128, -9223372036854775808}
                                                                        _ = v83
                                                                    }
                                                                } else {
                                                                    var v84 float32 = 0.0
                                                                    var v85 int32 = 42
                                                                    _ = v84
                                                                    _ = v85
                                                                }
                                                                _ = v72
                                                            }
                                                            _ = v36
                                                        } else {
                                                            v19 -= (-(v19))
                                                            v86 := []int{-691, (fi12 - 0)}
                                                            var v87 int64 = (v3 % ((-1) | 1))
                                                            _ = v86
                                                            _ = v87
                                                        }
                                                    case 1:
                                                        if false {
                                                            var v88 int = 1
                                                            switch v88 {
                                                            case 9223372036854775807:
                                                                v89 := GenS7{F0: (v7 + (((v7 + v7) + v7) + v7)), F1: 32767, F2: v7}
                                                                _ = v89
                                                            case 42:
                                                                if v30 {
                                                                    if ((3.14159 - 0.0) <= 1e10) {
                                                                        var v90 int16 = -397
                                                                        if (3.14159 <= 1e10) {
                                                                            if (0.1 > (float32(((0.1 - (-2.25 * float64(float32((3.14159 * ((1.5 / (-2.25 * 1e10)) - 1.5)))))) + 3.14159)) - -2.25)) {
                                                                                v88 = 445
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    if (true && v30) {
                                                                                        v19 -= 127
                                                                                        var v91 int = fi20
                                                                                        v92 := GenS13{F0: "\n\t"}
                                                                                        v93 := GenS2{F0: v19, F1: v19, F2: (v19 * (v19 % ((int8(v90)) | 1)))}
                                                                                        v94 := GenS6{F0: (-(fi12)), F1: -9223372036854775808}
                                                                                        _ = v91
                                                                                        _ = v92
                                                                                        _ = v93
                                                                                        _ = v94
                                                                                    }
                                                                                case "a":
                                                                                    var v95 int = fi20
                                                                                    _ = v95
                                                                                }
                                                                                var v96 uint = v13
                                                                                _ = v96
                                                                            } else {
                                                                                var v97 int16 = (-(v90))
                                                                                _ = v97
                                                                            }
                                                                            fmt.Println(fi12, v3, v3)
                                                                            v6 = v6
                                                                        }
                                                                        v3 -= v3
                                                                        _ = v90
                                                                    }
                                                                    if v30 {
                                                                        var v98 uint16 = 65535
                                                                        var v99 uint64 = 18446744073709551615
                                                                        switch "a" {
                                                                        case "":
                                                                            v99 -= 18446744073709551615
                                                                        }
                                                                        _ = v98
                                                                        _ = v99
                                                                    }
                                                                    switch 32768 {
                                                                    case 0:
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            var v100 int16 = 32767
                                                                            _ = v100
                                                                        case "foo bar":
                                                                            if ((v3 & -9223372036854775808) > -643) {
                                                                                v7 = "\n\t"
                                                                                fmt.Println(v3, v19, v7)
                                                                            }
                                                                        case "a":
                                                                            v13 = v13
                                                                            v7 = v7
                                                                            switch (v88 >> ((306) & 63)) {
                                                                            case 42:
                                                                                v101 := GenS14{F0: 218}
                                                                                _ = v101
                                                                            case 127:
                                                                                if (-291 > int16(v88)) {
                                                                                    v6 -= v6
                                                                                } else {
                                                                                    var v102 string = v7
                                                                                    _ = v102
                                                                                }
                                                                            }
                                                                        }
                                                                    case 127:
                                                                        v103 := GenS3{F0: int(v6)}
                                                                        _ = v103
                                                                    default:
                                                                        switch v7 {
                                                                        case "":
                                                                            v3 = 9223372036854775807
                                                                        case "foo bar":
                                                                            var v104 uint32 = 2147483648
                                                                            _ = v104
                                                                        case "hello":
                                                                            v105 := []int{}
                                                                            _ = v105
                                                                        }
                                                                        fmt.Println(v6, fi12)
                                                                    }
                                                                }
                                                            case 522:
                                                                fmt.Println(v19)
                                                            }
                                                            var v106 uint8 = 150
                                                            _ = v106
                                                            _ = v88
                                                        } else {
                                                            switch ("" + v7) {
                                                            case "a":
                                                                var v107 uint32 = v6
                                                                _ = v107
                                                            case "foo bar":
                                                                _ = len(v1)
                                                            case "hello":
                                                                var v108 uint64 = 728
                                                                _ = v108
                                                            }
                                                        }
                                                        var v109 int8 = (v19 / ((v19) | 1))
                                                        if true {
                                                            delete(v4, fi12)
                                                            fmt.Println(v6, v13)
                                                        } else {
                                                            v30 = true
                                                            v109 = v19
                                                            if true {
                                                                fmt.Println(fi12, fi20)
                                                                var v110 int = fi20
                                                                switch v7 {
                                                                case "\n\t":
                                                                    v111 := []int{}
                                                                    _ = v111
                                                                case "a":
                                                                    var v112 int = (fi20 | 901)
                                                                    _ = v112
                                                                default:
                                                                    v113 := GenS15{F0: 255, F1: 42}
                                                                    _ = v113
                                                                }
                                                                v114 := []int{0}
                                                                if len(v114) >= 2 {
                                                                    v114 = v114[1:len(v114)]
                                                                }
                                                                if v30 {
                                                                    var v115 uint8 = 42
                                                                    switch v7 {
                                                                    case "a":
                                                                        v1["\n\t"] = (fi20 << ((uint(v115)) & 63))
                                                                        v19 = (^(v109))
                                                                        v110 += (fi12 - (v110 ^ v110))
                                                                        v2[-771] = false
                                                                    default:
                                                                        var v116 uint64 = 18446744073709551615
                                                                        _ = v116
                                                                    }
                                                                    _ = v115
                                                                } else {
                                                                    v7 = v7
                                                                    var v117 uint8 = 23
                                                                    v118 := v1["hello"]
                                                                    _ = v117
                                                                    _ = v118
                                                                }
                                                                _ = v110
                                                                _ = v114
                                                            }
                                                        }
                                                        _ = v109
                                                    }
                                                    if (-1 >= 75) {
                                                        switch ("" + (((v7 + v7) + v7) + "")) {
                                                        case "foo bar":
                                                            var v119 uint64 = 18446744073709551615
                                                            var v120 int64 = (v3 & v3)
                                                            var v121 uint = (v13 + (^(v13)))
                                                            _ = v119
                                                            _ = v120
                                                            _ = v121
                                                        default:
                                                            if (!(v19 != -1)) {
                                                                switch "hello" {
                                                                case "hello":
                                                                    v3 += (^(v3))
                                                                }
                                                            }
                                                        }
                                                        if false {
                                                            var v122 bool = (!((true || (((fi12 << (((^(v13))) & 63)) >= -2147483648) && v30)) && v30))
                                                            v19 = v19
                                                            _ = v122
                                                        } else {
                                                            v123 := GenS16{F0: v19, F1: v3}
                                                            _ = v123
                                                        }
                                                        if (-9223372036854775808 < (v3 >> (((v13 + (v13 % ((v13) | 1)))) & 63))) {
                                                            if (fi12 != fi20) {
                                                                v124 := GenS10{F0: 4294967295, F1: (v13 ^ (v13 << (((v13 * 554)) & 63))), F2: v13}
                                                                var v125 int64 = v3
                                                                _ = v124
                                                                _ = v125
                                                            } else {
                                                                var v126 string = "\n\t"
                                                                switch (-(fi20)) {
                                                                case 2147483647:
                                                                    if ((!v30) || v30) {
                                                                        v127 := GenS11{F0: v7, F1: -148, F2: 2}
                                                                        v19 += (-(v19))
                                                                        fmt.Println(v19, v19, v3)
                                                                        _ = v127
                                                                    }
                                                                case 1:
                                                                    v6 = (v6 * 926)
                                                                case 571:
                                                                    if ((v30 && v30) && (true || (42 == 127))) {
                                                                        v126 = ""
                                                                        v7 = v126
                                                                        v2[fi12] = v30
                                                                    }
                                                                }
                                                                _ = v126
                                                            }
                                                        }
                                                    }
                                                    _ = v35
                                                default:
                                                    var v128 uint16 = 42
                                                    _ = v128
                                                }
                                            default:
                                                var v129 int32 = int32(v3)
                                                _ = v129
                                            }
                                            v13 = (v13 << (((v13 | 678)) & 63))
                                            _ = v30
                                        }
                                    }
                                case "a":
                                    if false {
                                        v6 = (^(v6))
                                    }
                                default:
                                    if (42 > 32767) {
                                        var v130 float32 = 0.1
                                        v131 := GenS7{F0: v7, F1: 948, F2: v7}
                                        switch v7 {
                                        case "foo bar":
                                            var v132 uint8 = 0
                                            _ = v132
                                        case "hello":
                                            switch (fi12 << ((v13) & 63)) {
                                            case 9223372036854775807:
                                                if true {
                                                    if (!false) {
                                                        if (!(!false)) {
                                                            var v133 uint = 255
                                                            switch (fi20 << ((255) & 63)) {
                                                            case 32768:
                                                                v13 = v133
                                                            }
                                                            if (false || true) {
                                                                var v134 uint = v13
                                                                var v135 int16 = -1
                                                                _ = v134
                                                                _ = v135
                                                            }
                                                            _ = v133
                                                        } else {
                                                            v136 := []int{}
                                                            _ = v136
                                                        }
                                                    } else {
                                                        v137 := GenS10{F0: 837, F1: 0, F2: v13}
                                                        var v138 string = "hello"
                                                        v130 = float32(1.5)
                                                        _ = v137
                                                        _ = v138
                                                    }
                                                    v139 := []int{(fi20 ^ (-(fi12)))}
                                                    switch (fi20 % ((fi20) | 1)) {
                                                    case 42:
                                                        v140 := v2[-274]
                                                        _ = v140
                                                    case -439:
                                                        v13 = v13
                                                    case 913:
                                                        v4[-9223372036854775808] = false
                                                    default:
                                                        var v141 int32 = 42
                                                        _ = v141
                                                    }
                                                    _ = v139
                                                } else {
                                                    v142 := []int{(fi12 ^ (fi12 + 176)), (fi12 ^ fi20), fi20, 51}
                                                    v6 = v6
                                                    v3 -= (v3 * (-(v3)))
                                                    var v143 uint64 = 18446744073709551615
                                                    v144 := GenS11{F0: v7, F1: -693, F2: (^(v143))}
                                                    _ = v142
                                                    _ = v143
                                                    _ = v144
                                                }
                                            }
                                        case "":
                                            v145 := GenS15{F0: 10, F1: -90}
                                            _ = v145
                                        }
                                        _ = v130
                                        _ = v131
                                    } else {
                                        fmt.Println(fi12)
                                        v146 := GenS17{F0: 0, F1: 1, F2: -9223372036854775808}
                                        v3 -= 1
                                        fmt.Println(fi20, v19, v3)
                                        _ = v146
                                    }
                                }
                                _ = fi20
                            }
                            var v147 uint16 = 2
                            _ = v147
                        } else {
                            if (!((141 >= 434) || (18446744073709551615 < 0))) {
                                v148 := GenS3{F0: fi12}
                                for fi149 := 0; fi149 < 17; fi149++ {
                                    v3 -= v3
                                    v7 = (("a" + v7) + v7)
                                    _ = fi149
                                }
                                for fi150 := 0; fi150 < 6; fi150++ {
                                    v19 -= 118
                                    v6 -= (v6 & uint32(v13))
                                    v19 -= -128
                                    _ = fi150
                                }
                                _ = v148
                            }
                            switch "" {
                            case "\n\t":
                                for fi151 := 0; fi151 < 1; fi151++ {
                                    switch ((v7 + v7) + "foo bar") {
                                    case "hello":
                                        var v152 bool = true
                                        _ = v152
                                    case "foo bar":
                                        switch (fi12 - fi12) {
                                        case -9223372036854775808:
                                            _ = len(v2)
                                        default:
                                            v153 := GenS18{F0: 255, F1: 1000}
                                            _ = v153
                                        }
                                    default:
                                        var v154 int64 = (v3 & v3)
                                        _ = v154
                                    }
                                    fmt.Println(v19)
                                    var v155 int8 = v19
                                    _ = fi151
                                    _ = v155
                                }
                            case "hello":
                                for fi156 := 0; fi156 < 1; fi156++ {
                                    fmt.Println(fi156, v3)
                                    v157 := GenS13{F0: "a"}
                                    switch (-(fi12)) {
                                    case -2147483648:
                                        switch "\n\t" {
                                        case "a":
                                            switch (^(fi12)) {
                                            case -568:
                                                fmt.Println(v7, fi12)
                                            case 410:
                                                v158 := GenS9{F0: v3}
                                                v159 := GenS13{F0: v7}
                                                _ = v158
                                                _ = v159
                                            }
                                        case "":
                                            var v160 uint = (v13 - v13)
                                            _ = v160
                                        case "\n\t":
                                            v161 := []int{}
                                            _ = v161
                                        default:
                                            if ((((v13 >= 2) || true) && false) && ((v3 | v3) == v3)) {
                                                if (!(!(1 <= 2))) {
                                                    v3 -= 315
                                                }
                                            }
                                        }
                                        v162 := GenS19{F0: v7, F1: 9223372036854775807, F2: 1}
                                        _ = v162
                                    case -671:
                                        var v163 int32 = -348
                                        _ = v163
                                    case 128:
                                        var v164 uint64 = 1
                                        _ = v164
                                    default:
                                        if (v19 <= v19) {
                                            v165 := GenS15{F0: 128, F1: (^(v19))}
                                            switch (fi12 ^ int(v13)) {
                                            case 602:
                                                fmt.Println(v13, v7)
                                            default:
                                                var v166 int8 = 127
                                                _ = v166
                                            }
                                            v6 += v6
                                            if (18446744073709551615 == v13) {
                                                var v167 float64 = (1e10 / 0.1)
                                                _ = len(v5)
                                                _ = v167
                                            }
                                            _ = v165
                                        }
                                    }
                                    switch v7 {
                                    case "\n\t":
                                        fmt.Println(v19, fi12)
                                    case "a":
                                        v168 := v1[("hello" + v7)]
                                        _ = v168
                                    case "":
                                        fmt.Println(fi156, v6, fi156)
                                    }
                                    _ = fi156
                                    _ = v157
                                }
                            default:
                                if true {
                                    if ((((fi12 - fi12) == fi12) || (false && true)) || true) {
                                        var v169 float64 = 1e10
                                        v4[(fi12 >> ((0) & 63))] = (4294967295 == 15)
                                        _ = v169
                                    }
                                } else {
                                    var v170 int8 = v19
                                    _ = v170
                                }
                            }
                            for fi171 := 0; fi171 < 20; fi171++ {
                                var v172 int32 = 824
                                _ = fi171
                                _ = v172
                            }
                        }
                        _ = v19
                    }
                    v7 = (v7 + v7)
                    v6 = v6
                } else {
                    for fi173 := 0; fi173 < 11; fi173++ {
                        v174 := GenS7{F0: v7, F1: 172, F2: v7}
                        _ = fi173
                        _ = v174
                    }
                    if (true && true) {
                        if (128 == (fi12 % ((fi12) | 1))) {
                            var v175 int8 = -1
                            fmt.Println(v13, v175)
                            if false {
                                var v176 string = v7
                                _ = v176
                            } else {
                                var v177 float32 = 3.14159
                                for fi178 := 0; fi178 < 5; fi178++ {
                                    var v179 int = (-(fi12))
                                    v180 := GenS2{F0: (v175 | (v175 % ((v175) | 1))), F1: (v175 + v175), F2: 0}
                                    v179 = fi178
                                    _ = fi178
                                    _ = v179
                                    _ = v180
                                }
                                _ = v177
                            }
                            _ = v175
                        } else {
                            v181 := GenS20{F0: 2}
                            _ = v181
                        }
                    }
                }
                v6 -= v6
                if false {
                    var v182 int64 = v3
                    _ = v182
                }
            }
            for fi183 := 0; fi183 < 11; fi183++ {
                v184 := GenS21{F0: 1, F1: fi183}
                var v185 int32 = -2147483648
                _ = fi183
                _ = v184
                _ = v185
            }
            _ = v15
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    v3 += v3
    v3 += v3
    var v186 int16 = 42
    if true {
        v187 := GenS22{F0: 127}
        v188 := []int{-1, 9223372036854775807}
        v7 = "\n\t"
        _ = v187
        _ = v188
    } else {
        if (42 < 537) {
            v189 := []int{-905}
            var v190 string = "foo bar"
            _ = v189
            _ = v190
        } else {
            var v191 int32 = -925
            for fi192 := 0; fi192 < 6; fi192++ {
                var v193 uint64 = 18446744073709551615
                for fi194 := 0; fi194 < 9; fi194++ {
                    if true {
                        v5["hello"] = v7
                    }
                    v195 := GenS23{F0: 107, F1: 41}
                    _ = fi194
                    _ = v195
                }
                v196 := GenS3{F0: int(v6)}
                _ = fi192
                _ = v193
                _ = v196
            }
            v191 = 0
            _ = v191
        }
    }
    v3 = -9223372036854775808
    fmt.Println(v3)
    fmt.Println(v6, v7)
    v186 = -131
    switch (v7 + v7) {
    case "\n\t":
        var v197 int = 0
        _ = v197
    case "a":
        delete(v5, (v7 + (v7 + v7)))
    case "foo bar":
        var v198 int32 = 2147483647
        _ = v198
    default:
        if (!((0 >= -329) && (932 != 622))) {
            var v199 bool = (!true)
            switch ("" + (((v7 + v7) + ("hello" + (v7 + v7))) + (v7 + v7))) {
            case "":
                var v200 int64 = v3
                _ = v200
            case "hello":
                if v199 {
                    for fi201 := 0; fi201 < 5; fi201++ {
                        _ = len(v1)
                        if (((-2.25 + 1e10) + 1e10) <= 1e10) {
                            var v202 float64 = ((3.14159 - ((0.0 + 3.14159) / float64(1.5))) / float64(((0.1 - 0.0) / 3.14159)))
                            v203 := GenS24{F0: (v6 << ((986) & 63)), F1: 2, F2: (fi201 ^ (fi201 & (fi201 >> ((636) & 63))))}
                            _ = v202
                            _ = v203
                        }
                        _ = fi201
                    }
                } else {
                    v7 = "hello"
                }
            }
            for fi204 := 0; fi204 < 10; fi204++ {
                if (!(55 == 128)) {
                    v205 := GenS14{F0: -975}
                    _ = v205
                }
                _ = fi204
            }
            _ = v199
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v186:", v186)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v186
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
