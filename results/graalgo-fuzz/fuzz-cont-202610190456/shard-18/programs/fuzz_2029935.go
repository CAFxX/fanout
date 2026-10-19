package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int8; F1 int; F2 uint }
type GenS4 struct { F0 bool; F1 int32; F2 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int8; F1 bool }
type GenS13 struct { F0 uint8 }

func main() {
    var v1 bool = (2 == 128)
    var v2 bool = v1
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    v5 := make(map[string]uint32)
    var v6 uint32 = 2147483647
    v7 := GenS1{F0: -898}
    v8 := GenS1{F0: 0}
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(fi9, fi9)
        _ = fi9
    }
    switch "foo bar" {
    case "foo bar":
        v10 := GenS2{F0: 65, F1: 1}
        _ = v10
    case "\n\t":
        var v11 int64 = 1
        _ = v11
    }
    v12 := GenS2{F0: -105, F1: 2}
    if false {
        fmt.Println(v1, v6, v2)
        var v13 uint16 = 65535
        for fi14 := 0; fi14 < 11; fi14++ {
            switch "hello" {
            case "":
                switch "" {
                case "":
                    if (0 > uint(v13)) {
                        v15 := GenS2{F0: -83, F1: 18446744073709551615}
                        _ = v15
                    } else {
                        v16 := GenS3{F0: 0, F1: (fi14 | 0), F2: 889}
                        v17 := GenS4{F0: (2 >= 420), F1: -870, F2: 0}
                        _ = v16
                        _ = v17
                    }
                case "foo bar":
                    fmt.Println(v6, v13, fi14)
                }
            case "hello":
                _ = len(v3)
            default:
                fmt.Println(v1)
            }
            var v18 int64 = 0
            v3[fi14] = true
            _ = fi14
            _ = v18
        }
        _ = v13
    } else {
        var v19 bool = (true && (1 <= 42))
        if (!false) {
            switch "hello" {
            case "hello":
                switch 127 {
                case 316:
                    switch -9223372036854775808 {
                    case -251:
                        for fi20 := 0; fi20 < 17; fi20++ {
                            switch 128 {
                            case 4294967295:
                                if (!v19) {
                                    for fi21 := 0; fi21 < 5; fi21++ {
                                        if true {
                                            v22 := GenS1{F0: 42}
                                            var v23 uint8 = 255
                                            v1 = true
                                            _ = v22
                                            _ = v23
                                        }
                                        v6 -= 2
                                        var v24 uint64 = 0
                                        v2 = ((false && true) || v2)
                                        if true {
                                            v2 = v2
                                            fmt.Println(v24, fi21)
                                            var v25 int8 = 127
                                            _ = v25
                                        } else {
                                            v1 = false
                                            v26 := []int{fi20}
                                            _ = v26
                                        }
                                        _ = fi21
                                        _ = v24
                                    }
                                    v27 := GenS2{F0: 24, F1: 18446744073709551615}
                                    var v28 uint8 = 150
                                    _ = v27
                                    _ = v28
                                } else {
                                    v29 := GenS5{F0: fi20}
                                    switch "" {
                                    case "foo bar":
                                        v19 = ((!true) || false)
                                    case "":
                                        v30 := GenS6{F0: 127}
                                        _ = v30
                                    }
                                    var v31 uint = 494
                                    _ = v29
                                    _ = v31
                                }
                            default:
                                for fi32 := 0; fi32 < 9; fi32++ {
                                    v33 := []int{-484, 4294967295, (fi32 + fi20)}
                                    if v2 {
                                        fmt.Println(v2, v6, fi32)
                                        switch "a" {
                                        case "a":
                                            fmt.Println(v2)
                                            if (v1 && true) {
                                                if ((fi32 ^ fi20) != -1) {
                                                    switch fi20 {
                                                    case -336:
                                                        v34 := GenS7{F0: int64(fi32)}
                                                        _ = v34
                                                    case -166:
                                                        v35 := GenS3{F0: 127, F1: (fi32 << ((2) & 63)), F2: 0}
                                                        _ = v35
                                                    default:
                                                        _ = len(v4)
                                                    }
                                                    v19 = v2
                                                    fmt.Println(v19, fi20, fi20)
                                                }
                                            } else {
                                                var v36 int16 = 32767
                                                var v37 bool = v2
                                                if len(v33) >= 2 {
                                                    v33 = v33[1:len(v33)]
                                                }
                                                _ = v36
                                                _ = v37
                                            }
                                            var v38 uint32 = (v6 << ((0) & 63))
                                            _ = v38
                                        default:
                                            if false {
                                                switch fi32 {
                                                case 0:
                                                    if false {
                                                        v39 := GenS8{F0: -32768}
                                                        _ = v39
                                                    } else {
                                                        if (float32(1e10) <= 3.14159) {
                                                            v40 := v5[""]
                                                            var v41 uint16 = 117
                                                            _ = v40
                                                            _ = v41
                                                        } else {
                                                            var v42 uint8 = uint8(v6)
                                                            v2 = false
                                                            switch ("foo bar" + "hello") {
                                                            case "":
                                                                if false {
                                                                    v43 := GenS5{F0: fi20}
                                                                    _ = v43
                                                                } else {
                                                                    var v44 int32 = 0
                                                                    var v45 uint16 = uint16(v6)
                                                                    var v46 string = ("hello" + "")
                                                                    _ = v44
                                                                    _ = v45
                                                                    _ = v46
                                                                }
                                                            case "foo bar":
                                                                if (42 > 1) {
                                                                    var v47 string = ("foo bar" + (((("\n\t" + "a") + "") + "\n\t") + "\n\t"))
                                                                    _ = v47
                                                                } else {
                                                                    var v48 uint = 18446744073709551615
                                                                    var v49 uint = v48
                                                                    switch "a" {
                                                                    case "a":
                                                                        switch (fi32 * (fi32 % ((32768) | 1))) {
                                                                        case 2147483647:
                                                                            v50 := GenS6{F0: 253}
                                                                            _ = v50
                                                                        case -2147483648:
                                                                            delete(v5, "\n\t")
                                                                        case -9223372036854775808:
                                                                            if v19 {
                                                                                var v51 bool = false
                                                                                _ = v51
                                                                            }
                                                                        }
                                                                    case "hello":
                                                                        if (-128 >= int8(fi32)) {
                                                                            if v1 {
                                                                                var v52 int = fi20
                                                                                var v53 float32 = 3.14159
                                                                                fmt.Println(v49, fi32, v52)
                                                                                _ = v52
                                                                                _ = v53
                                                                            }
                                                                            v54 := v5[""]
                                                                            _ = v54
                                                                        } else {
                                                                            v19 = (!false)
                                                                        }
                                                                    default:
                                                                        v55 := v3[fi20]
                                                                        _ = v55
                                                                    }
                                                                    _ = v48
                                                                    _ = v49
                                                                }
                                                            }
                                                            _ = v42
                                                        }
                                                    }
                                                case 32768:
                                                    v19 = v19
                                                }
                                            } else {
                                                var v56 uint8 = 2
                                                var v57 uint64 = 1
                                                _ = v56
                                                _ = v57
                                            }
                                        }
                                        v58 := GenS8{F0: 42}
                                        _ = v58
                                    } else {
                                        v19 = (0 != 104)
                                        _ = len(v33)
                                        switch ("a" + "a") {
                                        case "":
                                            if v1 {
                                                switch (-(fi32)) {
                                                case 165:
                                                    if len(v33) > 0 {
                                                        v33[0] = (fi32 >> ((0) & 63))
                                                    }
                                                case 32768:
                                                    var v59 int = fi32
                                                    _ = v59
                                                default:
                                                    if (661 >= 0) {
                                                        if (0 == 317) {
                                                            if (!v1) {
                                                                fmt.Println(v1, v19, fi32)
                                                                if (-2.25 > 3.14159) {
                                                                    v60 := GenS3{F0: 42, F1: (^(fi20)), F2: 1}
                                                                    v19 = v19
                                                                    v19 = v19
                                                                    _ = v60
                                                                } else {
                                                                    if false {
                                                                        v1 = (!v19)
                                                                        var v61 uint32 = v6
                                                                        v61 -= v6
                                                                        if v19 {
                                                                            if v1 {
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    if len(v33) > 0 {
                                                                                        v33[0] = (^(fi20))
                                                                                    }
                                                                                case "foo bar":
                                                                                    v5["\n\t"] = (v6 << ((42) & 63))
                                                                                case "":
                                                                                    v62 := GenS8{F0: 32767}
                                                                                    _ = v62
                                                                                default:
                                                                                    v19 = (524 < 65535)
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v1, v19, v1)
                                                                                var v63 float32 = float32(0.1)
                                                                                var v64 float64 = 3.14159
                                                                                _ = v63
                                                                                _ = v64
                                                                            }
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                v6 += v61
                                                                            default:
                                                                                v33 = append(v33, fi20)
                                                                            }
                                                                            var v65 float32 = (3.14159 * float32(0.0))
                                                                            if (-9223372036854775808 <= 1) {
                                                                                _ = len(v33)
                                                                                var v66 int = fi20
                                                                                fmt.Println(fi32, v2)
                                                                                _ = v66
                                                                            } else {
                                                                                fmt.Println(fi20, fi20, v19)
                                                                                v67 := v3[(^(fi20))]
                                                                                v68 := GenS2{F0: -128, F1: 0}
                                                                                switch fi32 {
                                                                                case 1:
                                                                                    if ((v2 || v1) && (1 < 0)) {
                                                                                        _ = len(v33)
                                                                                        fmt.Println(v2, v65, v2)
                                                                                        var v69 bool = ((("\n\t" + "\n\t") + "\n\t") <= "")
                                                                                        _ = v69
                                                                                    } else {
                                                                                        fmt.Println(v19, fi20)
                                                                                    }
                                                                                case 9223372036854775807:
                                                                                    fmt.Println(v6, v1)
                                                                                    fmt.Println(fi32, v67)
                                                                                    if v19 {
                                                                                        v70 := GenS7{F0: 0}
                                                                                        v71 := GenS9{F0: 128}
                                                                                        _ = v70
                                                                                        _ = v71
                                                                                    } else {
                                                                                        switch (fi32 + fi32) {
                                                                                        case 1:
                                                                                            switch "foo bar" {
                                                                                            case "":
                                                                                                fmt.Println(v2, v6, v19)
                                                                                            case "a":
                                                                                                var v72 int32 = 70
                                                                                                _ = v72
                                                                                            }
                                                                                        case 42:
                                                                                            fmt.Println(v65, fi20, fi20)
                                                                                        case 127:
                                                                                            var v73 uint8 = 184
                                                                                            _ = v73
                                                                                        default:
                                                                                            v74 := v5["foo bar"]
                                                                                            _ = v74
                                                                                        }
                                                                                        var v75 uint64 = 1
                                                                                        _ = v75
                                                                                    }
                                                                                }
                                                                                v19 = (32767 < 65535)
                                                                                fmt.Println(v6)
                                                                                _ = v67
                                                                                _ = v68
                                                                            }
                                                                            v76 := GenS5{F0: 1}
                                                                            _ = v65
                                                                            _ = v76
                                                                        }
                                                                        var v77 string = "\n\t"
                                                                        var v78 int64 = 0
                                                                        _ = v61
                                                                        _ = v77
                                                                        _ = v78
                                                                    } else {
                                                                        var v79 uint32 = v6
                                                                        if true {
                                                                            if true {
                                                                                v19 = v19
                                                                                if v19 {
                                                                                    _ = len(v3)
                                                                                    if len(v33) >= 2 {
                                                                                        v33 = v33[1:len(v33)]
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(v1)
                                                                                    if (v2 && false) {
                                                                                        var v80 string = "\n\t"
                                                                                        if true {
                                                                                            v79 += v79
                                                                                            fmt.Println(v2, v6, v6)
                                                                                            var v81 int = (fi20 << ((1) & 63))
                                                                                            _ = v81
                                                                                        } else {
                                                                                            v1 = (false || false)
                                                                                            if (!v19) {
                                                                                                switch "hello" {
                                                                                                case "foo bar":
                                                                                                    switch (fi20 & fi20) {
                                                                                                    case 32767:
                                                                                                        var v82 float64 = float64((float32(0.1) * 0.0))
                                                                                                        _ = v82
                                                                                                    case -9223372036854775808:
                                                                                                        v19 = (!v2)
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    v1 = true
                                                                                                }
                                                                                            } else {
                                                                                                v79 += 2
                                                                                                var v83 float64 = 3.14159
                                                                                                var v84 uint16 = 0
                                                                                                _ = v83
                                                                                                _ = v84
                                                                                            }
                                                                                            if (316 > -2147483648) {
                                                                                                var v85 uint32 = (v6 | (v6 + 0))
                                                                                                fmt.Println(v2, v79)
                                                                                                var v86 float64 = 1.5
                                                                                                _ = v85
                                                                                                _ = v86
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(fi32)
                                                                                        _ = v80
                                                                                    }
                                                                                    var v87 uint64 = 18446744073709551615
                                                                                    _ = v87
                                                                                }
                                                                                v1 = (!v19)
                                                                            }
                                                                        } else {
                                                                            var v88 uint8 = 52
                                                                            _ = v88
                                                                        }
                                                                        var v89 uint64 = 4294967295
                                                                        _ = v79
                                                                        _ = v89
                                                                    }
                                                                    var v90 uint64 = 2
                                                                    _ = v90
                                                                }
                                                            } else {
                                                                v2 = (v2 || (451 < 42))
                                                                v91 := GenS10{F0: -2147483648}
                                                                if true {
                                                                    switch fi32 {
                                                                    case 32768:
                                                                        fmt.Println(v6, v6, fi20)
                                                                    }
                                                                    var v92 int8 = -1
                                                                    v92 = (v92 | v92)
                                                                    _ = v92
                                                                } else {
                                                                    if v1 {
                                                                        switch -9223372036854775808 {
                                                                        case 0:
                                                                            if (false || v1) {
                                                                                if (false || v19) {
                                                                                    v6 = (v6 ^ v6)
                                                                                    v6 += v6
                                                                                    var v93 uint16 = 345
                                                                                    v19 = (!(!((!(!(!(int32(v6) > -688)))) && false)))
                                                                                    v93 += (v93 & uint16(v6))
                                                                                    _ = v93
                                                                                } else {
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        var v94 int32 = 665
                                                                                        _ = v94
                                                                                    case "a":
                                                                                        var v95 float32 = 0.0
                                                                                        _ = v95
                                                                                    }
                                                                                    if len(v33) > 0 {
                                                                                        v33[0] = fi32
                                                                                    }
                                                                                }
                                                                                v2 = (619 != 42)
                                                                                v33 = append(v33, fi32)
                                                                            } else {
                                                                                v6 = v6
                                                                                var v96 uint32 = 675
                                                                                v2 = (true || (100 != 1))
                                                                                var v97 float32 = float32(0.0)
                                                                                var v98 uint64 = 1
                                                                                var v99 int64 = 202
                                                                                var v100 bool = (-28 != 42)
                                                                                _ = v100
                                                                                _ = v96
                                                                                _ = v97
                                                                                _ = v98
                                                                                _ = v99
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if (!v2) {
                                                                            v101 := v4["\n\t"]
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                v102 := GenS11{F0: 0}
                                                                                _ = v102
                                                                            case "a":
                                                                                if v19 {
                                                                                    v103 := GenS2{F0: 42, F1: 565}
                                                                                    fmt.Println(v19, fi20, fi32)
                                                                                    var v104 uint16 = 65535
                                                                                    _ = v103
                                                                                    _ = v104
                                                                                }
                                                                            case "\n\t":
                                                                                if v19 {
                                                                                    v105 := GenS8{F0: -838}
                                                                                    if v1 {
                                                                                        if false {
                                                                                            if v1 {
                                                                                                if v1 {
                                                                                                    v19 = v19
                                                                                                    v6 -= v6
                                                                                                    v106 := v3[fi20]
                                                                                                    _ = v106
                                                                                                }
                                                                                                fmt.Println(fi32, fi20, v6)
                                                                                                v19 = v1
                                                                                            }
                                                                                        }
                                                                                        var v107 bool = (float32(1.5) > float32(3.14159))
                                                                                        _ = v107
                                                                                    } else {
                                                                                        v108 := GenS1{F0: -1}
                                                                                        if v1 {
                                                                                            v109 := GenS10{F0: -2147483648}
                                                                                            if ((v1 || ("a" != (("hello" + "hello") + "hello"))) && false) {
                                                                                                v110 := GenS12{F0: 100, F1: v19}
                                                                                                v111 := GenS11{F0: -183}
                                                                                                switch ("foo bar" + (("" + "foo bar") + "foo bar")) {
                                                                                                case "":
                                                                                                    v3[(fi20 ^ v101)] = v1
                                                                                                case "a":
                                                                                                    v6 += v6
                                                                                                default:
                                                                                                    v1 = ((v6 - v6) == (v6 & v6))
                                                                                                }
                                                                                                var v112 int16 = 32767
                                                                                                var v113 int = (v101 - (fi20 | v101))
                                                                                                v19 = v2
                                                                                                _ = v110
                                                                                                _ = v111
                                                                                                _ = v112
                                                                                                _ = v113
                                                                                            } else {
                                                                                                v33 = append(v33, (fi32 >> ((42) & 63)))
                                                                                                v5["hello"] = v6
                                                                                                fmt.Println(v1, fi32, fi32)
                                                                                                var v114 float64 = ((((float64((0.0 + -2.25)) * (3.14159 + 3.14159)) / float64(((-2.25 / (float32((1e10 / ((0.1 * 0.0) - ((1.5 + (0.0 * 3.14159)) + float64((1.5 * 1.5)))))) * 1e10)) + 3.14159))) + 0.1) - 0.1)
                                                                                                _ = v114
                                                                                            }
                                                                                            _ = v109
                                                                                        } else {
                                                                                            var v115 uint = 0
                                                                                            _ = v115
                                                                                        }
                                                                                        fmt.Println(v2)
                                                                                        _ = v108
                                                                                    }
                                                                                    switch fi32 {
                                                                                    case 279:
                                                                                        if ((("hello" + "a") + "hello") != ("hello" + "")) {
                                                                                            v116 := v4["foo bar"]
                                                                                            _ = v116
                                                                                        }
                                                                                    }
                                                                                    _ = v105
                                                                                }
                                                                            }
                                                                            v6 = (v6 - (^(v6)))
                                                                            _ = v101
                                                                        }
                                                                        var v117 int16 = -1
                                                                        _ = v117
                                                                    }
                                                                    v19 = (false && (-24 < 100))
                                                                }
                                                                if len(v33) >= 2 {
                                                                    v33 = v33[1:len(v33)]
                                                                }
                                                                var v118 uint8 = 255
                                                                _ = v118
                                                                _ = v91
                                                            }
                                                            v19 = v1
                                                        } else {
                                                            if (32768 > (fi20 % (((fi20 & (fi32 | fi32))) | 1))) {
                                                                if v2 {
                                                                    fmt.Println(fi32, v2)
                                                                    var v119 float64 = (((1e10 + 1e10) + (0.0 / -2.25)) - 3.14159)
                                                                    _ = v119
                                                                }
                                                                var v120 uint64 = 2
                                                                _ = v120
                                                            }
                                                            var v121 uint8 = 192
                                                            _ = v121
                                                        }
                                                        if len(v33) >= 2 {
                                                            v33 = v33[1:len(v33)]
                                                        }
                                                        fmt.Println(v2, v6)
                                                    } else {
                                                        v122 := GenS11{F0: -278}
                                                        _ = len(v33)
                                                        var v123 uint32 = 564
                                                        _ = v122
                                                        _ = v123
                                                    }
                                                }
                                            } else {
                                                var v124 int64 = -1
                                                var v125 int64 = 1
                                                switch ("a" + (("" + "\n\t") + "")) {
                                                case "":
                                                    v19 = v1
                                                }
                                                _ = v124
                                                _ = v125
                                            }
                                        }
                                    }
                                    v19 = (0 <= 2)
                                    _ = fi32
                                    _ = v33
                                }
                            }
                            var v126 uint32 = (v6 >> ((428) & 63))
                            _ = fi20
                            _ = v126
                        }
                    case 714:
                        for fi127 := 0; fi127 < 13; fi127++ {
                            if v1 {
                                var v128 int8 = -58
                                var v129 float32 = (1.5 - ((1.5 * 1e10) / -2.25))
                                v130 := GenS9{F0: 42}
                                _ = v128
                                _ = v129
                                _ = v130
                            }
                            for fi131 := 0; fi131 < 1; fi131++ {
                                if (v1 && ((42 != -1) || (v1 || ((v19 && v2) || true)))) {
                                    var v132 int32 = 2147483647
                                    v2 = v19
                                    _ = v132
                                }
                                switch "a" {
                                case "\n\t":
                                    v1 = (-1 == 100)
                                default:
                                    v133 := v4[(("hello" + (("foo bar" + (("\n\t" + "hello") + "a")) + "")) + "hello")]
                                    _ = v133
                                }
                                var v134 int = 71
                                fmt.Println(fi127, fi131)
                                v135 := []int{}
                                _ = fi131
                                _ = v134
                                _ = v135
                            }
                            switch fi127 {
                            case 32767:
                                v2 = v2
                            case 172:
                                var v136 int32 = 0
                                _ = v136
                            case -2147483648:
                                v2 = v2
                            default:
                                for fi137 := 0; fi137 < 15; fi137++ {
                                    v138 := GenS5{F0: fi137}
                                    _ = fi137
                                    _ = v138
                                }
                            }
                            _ = fi127
                        }
                    case -753:
                        v6 += (v6 - v6)
                    default:
                        v139 := v3[0]
                        _ = v139
                    }
                case 567:
                    _ = len(v3)
                case 128:
                    var v140 bool = (787 > 651)
                    _ = v140
                }
            }
            v141 := GenS13{F0: 127}
            fmt.Println(v1, v6, v1)
            _ = v141
        } else {
            v6 = 42
            v2 = true
        }
        _ = v19
    }
    switch 128 {
    case -170:
        var v142 string = "foo bar"
        _ = v142
    case 32767:
        v143 := []int{-9223372036854775808}
        _ = v143
    case -2147483648:
        v144 := []int{32768, 32767}
        _ = v144
    default:
        v2 = ((((0.0 * ((1.5 - -2.25) / (0.1 * -2.25))) * 1.5) * 1.5) < -2.25)
    }
    var v145 uint = 42
    var v146 uint16 = 2
    v145 += v145
    fmt.Println(v1, v1, v2)
    var v147 uint32 = 557
    fmt.Println("v1:", v1)
    fmt.Println("v145:", v145)
    fmt.Println("v146:", v146)
    fmt.Println("v147:", v147)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v145
    _ = v146
    _ = v147
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
