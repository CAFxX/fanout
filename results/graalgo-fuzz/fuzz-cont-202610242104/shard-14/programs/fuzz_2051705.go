package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int32; F1 uint }
type GenS10 struct { F0 uint16; F1 uint }
type GenS11 struct { F0 int8; F1 uint32; F2 int16 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint32; F1 int32 }
type GenS14 struct { F0 int64; F1 int64; F2 uint }
type GenS15 struct { F0 int8; F1 int64; F2 uint64 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS18 struct { F0 int64; F1 int32 }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 uint16 }
type GenS21 struct { F0 string }
func genFn0() int8 {
    var v4 bool = false
    if true {
        v5 := GenS1{F0: 1}
        v4 = v4
        if v4 {
            if (false || v4) {
                v4 = (-128 <= 78)
            } else {
                var v6 uint32 = 283
                var v7 float32 = -2.25
                switch -1 {
                case -9223372036854775808:
                    var v8 int16 = 32767
                    _ = v8
                case 9223372036854775807:
                    switch 772 {
                    case -489:
                        for fi9 := 0; fi9 < 19; fi9++ {
                            var v10 int8 = 127
                            switch (fi9 | (fi9 - fi9)) {
                            case -518:
                                v4 = (65535 == 1)
                            case 838:
                                if true {
                                    v6 -= (v6 / (((v6 + (v6 / (((^(v6))) | 1)))) | 1))
                                    if ("foo bar" >= "hello") {
                                        v4 = v4
                                        for fi11 := 0; fi11 < 7; fi11++ {
                                            v12 := GenS2{F0: 42, F1: -2147483648, F2: 0}
                                            if v4 {
                                                var v13 float32 = 3.14159
                                                var v14 int64 = -9223372036854775808
                                                _ = v13
                                                _ = v14
                                            }
                                            _ = fi11
                                            _ = v12
                                        }
                                        v15 := GenS1{F0: 185}
                                        _ = v15
                                    }
                                } else {
                                    var v16 uint = 65535
                                    _ = v16
                                }
                            }
                            _ = fi9
                            _ = v10
                        }
                    case 95:
                        v17 := []int{1, 0, 1}
                        _ = v17
                    }
                default:
                    v7 += float32(1.5)
                }
                _ = v6
                _ = v7
            }
            v4 = ((2147483647 != 128) && (false || true))
        } else {
            if (v4 && true) {
                v4 = ((0 < 17) && (352 >= 160))
                switch "" {
                case "\n\t":
                    for fi18 := 0; fi18 < 17; fi18++ {
                        var v19 uint8 = 169
                        switch ((("\n\t" + "foo bar") + "a") + "foo bar") {
                        case "\n\t":
                            var v20 string = "hello"
                            _ = v20
                        case "a":
                            v21 := GenS1{F0: 2}
                            _ = v21
                        case "hello":
                            v4 = v4
                        }
                        switch "foo bar" {
                        case "foo bar":
                            var v22 int16 = -1
                            _ = v22
                        case "hello":
                            for fi23 := 0; fi23 < 17; fi23++ {
                                fmt.Println(fi18)
                                switch "foo bar" {
                                case "foo bar":
                                    if (!v4) {
                                        v24 := GenS3{F0: 255, F1: 42}
                                        if (((("foo bar" + "hello") + "hello") + "\n\t") != "") {
                                            if (-625 > 9223372036854775807) {
                                                var v25 float64 = 1.5
                                                var v26 float64 = (((3.14159 + ((v25 - v25) - (float64(float32((3.14159 - v25))) / (1e10 + 3.14159)))) - (v25 * 3.14159)) - v25)
                                                _ = v25
                                                _ = v26
                                            } else {
                                                var v27 uint8 = 152
                                                var v28 string = ("hello" + ("" + ("foo bar" + (("a" + "") + ("hello" + "hello")))))
                                                _ = v27
                                                _ = v28
                                            }
                                            v29 := []int{(fi23 + (fi23 & (fi23 >> ((65535) & 63)))), (fi18 & 171), (fi18 | fi18), (fi18 >> ((533) & 63))}
                                            var v30 int16 = 444
                                            _ = v29
                                            _ = v30
                                        } else {
                                            if v4 {
                                                v4 = true
                                                switch 0 {
                                                case 32768:
                                                    var v31 uint32 = 0
                                                    _ = v31
                                                }
                                            } else {
                                                if v4 {
                                                    v32 := GenS2{F0: 763, F1: 42, F2: (v19 << ((18446744073709551615) & 63))}
                                                    v19 = v19
                                                    _ = v32
                                                }
                                                if v4 {
                                                    if (v4 && v4) {
                                                        fmt.Println(v4, v19, v19)
                                                    }
                                                    var v33 bool = v4
                                                    fmt.Println(v33, v33, v19)
                                                    _ = v33
                                                }
                                            }
                                            if false {
                                                v34 := GenS2{F0: 46, F1: 2147483647, F2: v19}
                                                _ = v34
                                            }
                                        }
                                        _ = v24
                                    }
                                    var v35 float32 = -2.25
                                    if v4 {
                                        if v4 {
                                            var v36 int64 = -979
                                            v36 -= (v36 / ((v36) | 1))
                                            if true {
                                                var v37 bool = v4
                                                v38 := GenS3{F0: 42, F1: v36}
                                                _ = v37
                                                _ = v38
                                            } else {
                                                if (uint16(v19) < 255) {
                                                    v36 = 0
                                                    v35 = (-2.25 + (-2.25 + v35))
                                                    var v39 uint32 = 42
                                                    _ = v39
                                                } else {
                                                    v36 = v36
                                                }
                                                if (1 == v19) {
                                                    var v40 uint64 = 1
                                                    switch "a" {
                                                    case "":
                                                        fmt.Println(v19, fi23)
                                                    default:
                                                        v35 -= (3.14159 * 1e10)
                                                        v41 := []int{fi23, (-(fi18)), (fi23 >> ((1) & 63)), fi18}
                                                        if (32767 != 730) {
                                                            if v4 {
                                                                v42 := GenS4{F0: 0, F1: (-(v36)), F2: 128}
                                                                var v43 int = fi23
                                                                v44 := GenS4{F0: 255, F1: (^(v36)), F2: v19}
                                                                if v4 {
                                                                    if len(v41) > 0 {
                                                                        v41[0] = (fi23 & (v43 + -402))
                                                                    }
                                                                    var v45 uint16 = 306
                                                                    var v46 float64 = (0.1 - 0.1)
                                                                    _ = v45
                                                                    _ = v46
                                                                } else {
                                                                    v41 = append(v41, v43)
                                                                    if (!v4) {
                                                                        v41 = append(v41, (fi23 << ((uint(v19)) & 63)))
                                                                    }
                                                                    if (!(92 >= (v19 & (^(v19))))) {
                                                                        var v47 float64 = (1.5 + 1e10)
                                                                        _ = v47
                                                                    } else {
                                                                        v35 -= -2.25
                                                                    }
                                                                }
                                                                var v48 uint64 = (v40 * v40)
                                                                _ = v42
                                                                _ = v43
                                                                _ = v44
                                                                _ = v48
                                                            } else {
                                                                var v49 int64 = (v36 + (v36 << ((uint(v40)) & 63)))
                                                                var v50 uint32 = 4294967295
                                                                switch ("\n\t" + "\n\t") {
                                                                case "\n\t":
                                                                    var v51 int = (^(fi23))
                                                                    _ = v51
                                                                case "":
                                                                    var v52 bool = v4
                                                                    _ = v52
                                                                case "hello":
                                                                    if v4 {
                                                                        if len(v41) >= 2 {
                                                                            v41 = v41[1:len(v41)]
                                                                        }
                                                                        fmt.Println(v49, v36, v49)
                                                                    }
                                                                }
                                                                _ = v49
                                                                _ = v50
                                                            }
                                                        } else {
                                                            var v53 bool = (((v35 * (v35 * ((v35 / 0.1) / (1.5 / (float32(1e10) + (v35 / 0.1)))))) - ((v35 / v35) + v35)) >= 0.1)
                                                            if (0 != -44) {
                                                                if len(v41) > 0 {
                                                                    v41[0] = (fi18 | (^(fi23)))
                                                                }
                                                                v36 -= v36
                                                                if v53 {
                                                                    if v53 {
                                                                        fmt.Println(fi23, v40, v35)
                                                                        var v54 int = fi18
                                                                        if len(v41) > 0 {
                                                                            v41[0] = fi18
                                                                        }
                                                                        _ = v54
                                                                    } else {
                                                                        if (((((((0.1 + -2.25) * (-2.25 - 1e10)) * (0.0 * 0.0)) - ((float64((v35 + v35)) / 1e10) - (1.5 - 1.5))) * (0.0 - 1e10)) > (1.5 - (((((1.5 + 3.14159) * 0.1) * 3.14159) + -2.25) - 1.5))) || v53) {
                                                                            v4 = true
                                                                        }
                                                                        var v55 uint64 = 4294967295
                                                                        _ = v55
                                                                    }
                                                                    fmt.Println(v36, v4, v35)
                                                                    v56 := GenS5{F0: 19}
                                                                    _ = v56
                                                                } else {
                                                                    v57 := GenS5{F0: 42}
                                                                    v35 = v35
                                                                    v53 = true
                                                                    _ = v57
                                                                }
                                                                fmt.Println(v19, v36, v36)
                                                                v35 = (1.5 + (v35 * v35))
                                                            } else {
                                                                fmt.Println(v36, v35, v4)
                                                                v58 := GenS5{F0: int8(v36)}
                                                                if (v4 && (v40 <= 1)) {
                                                                    v41 = append(v41, (fi23 - 32767))
                                                                    v4 = (v35 == v35)
                                                                    switch fi23 {
                                                                    case 128:
                                                                        if v53 {
                                                                            _ = len(v41)
                                                                            var v59 uint32 = 2
                                                                            v60 := GenS6{F0: fi23}
                                                                            _ = v59
                                                                            _ = v60
                                                                        } else {
                                                                            if (0 >= 33) {
                                                                                if len(v41) >= 2 {
                                                                                    v41 = v41[1:len(v41)]
                                                                                }
                                                                                v4 = v4
                                                                            }
                                                                            var v61 string = "hello"
                                                                            v62 := GenS7{F0: -2147483648}
                                                                            if (v36 > v36) {
                                                                                var v63 uint8 = (v19 / ((uint8(v40)) | 1))
                                                                                if true {
                                                                                    var v64 uint64 = uint64(v19)
                                                                                    fmt.Println(v61)
                                                                                    _ = v64
                                                                                } else {
                                                                                    v65 := GenS8{F0: 4294967295}
                                                                                    v41 = append(v41, 32768)
                                                                                    var v66 uint32 = 0
                                                                                    var v67 uint32 = v66
                                                                                    var v68 uint32 = 2147483647
                                                                                    _ = v65
                                                                                    _ = v66
                                                                                    _ = v67
                                                                                    _ = v68
                                                                                }
                                                                                v35 -= -2.25
                                                                                v69 := GenS9{F0: 0, F1: 88}
                                                                                v35 += v35
                                                                                _ = v63
                                                                                _ = v69
                                                                            } else {
                                                                                if (false || v53) {
                                                                                    switch (fi18 + fi18) {
                                                                                    case -2147483648:
                                                                                        v40 = v40
                                                                                    case -865:
                                                                                        fmt.Println(fi18, fi23)
                                                                                    case 160:
                                                                                        var v70 bool = ((false || ((v36 / (((-(v36))) | 1)) != (^(v36)))) || (!((42 >= int8(v36)) && (v53 && (!((v19 - (v19 % ((v19) | 1))) >= (v19 - (v19 * (v19 | 127)))))))))
                                                                                        _ = v70
                                                                                    }
                                                                                }
                                                                                v71 := GenS8{F0: 208}
                                                                                var v72 int64 = v36
                                                                                _ = v71
                                                                                _ = v72
                                                                            }
                                                                            v36 = v36
                                                                            _ = v61
                                                                            _ = v62
                                                                        }
                                                                    case -2147483648:
                                                                        fmt.Println(v40, v35)
                                                                    case 0:
                                                                        fmt.Println(v40)
                                                                    default:
                                                                        v73 := GenS10{F0: 65535, F1: 65535}
                                                                        v41 = append(v41, (fi18 - fi23))
                                                                        v35 = (0.0 + v35)
                                                                        _ = v73
                                                                    }
                                                                } else {
                                                                    if v4 {
                                                                        if (1.5 > v35) {
                                                                            v74 := GenS9{F0: -1, F1: 4294967295}
                                                                            switch "" {
                                                                            case "":
                                                                                v75 := GenS11{F0: -1, F1: 65535, F2: 1}
                                                                                _ = v75
                                                                            case "\n\t":
                                                                                _ = len(v41)
                                                                            }
                                                                            v4 = v53
                                                                            _ = v74
                                                                        } else {
                                                                            _ = len(v41)
                                                                            v19 = v19
                                                                            if len(v41) > 0 {
                                                                                v41[0] = (^(fi18))
                                                                            }
                                                                            if v4 {
                                                                                var v76 uint8 = (v19 + (v19 ^ (^(v19))))
                                                                                _ = v76
                                                                            } else {
                                                                                var v77 int32 = int32(v36)
                                                                                var v78 uint64 = v40
                                                                                var v79 int64 = (v36 << ((uint(v19)) & 63))
                                                                                _ = v77
                                                                                _ = v78
                                                                                _ = v79
                                                                            }
                                                                            if (v35 < 3.14159) {
                                                                                fmt.Println(v4)
                                                                                var v80 int64 = v36
                                                                                var v81 int16 = -1
                                                                                _ = v80
                                                                                _ = v81
                                                                            } else {
                                                                                var v82 uint16 = 42
                                                                                _ = v82
                                                                            }
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            fmt.Println(v19)
                                                                        case "":
                                                                            switch fi23 {
                                                                            case -2147483648:
                                                                                if true {
                                                                                    var v83 uint16 = 255
                                                                                    var v84 float32 = v35
                                                                                    var v85 int16 = 32767
                                                                                    v36 = int64(fi23)
                                                                                    var v86 int64 = -643
                                                                                    if v4 {
                                                                                        if v4 {
                                                                                            fmt.Println(v84)
                                                                                            v19 += v19
                                                                                        }
                                                                                        v40 = v40
                                                                                    } else {
                                                                                        v36 -= v36
                                                                                        _ = len(v41)
                                                                                    }
                                                                                    if len(v41) >= 2 {
                                                                                        v41 = v41[1:len(v41)]
                                                                                    }
                                                                                    _ = v83
                                                                                    _ = v84
                                                                                    _ = v85
                                                                                    _ = v86
                                                                                }
                                                                            case 32768:
                                                                                v87 := GenS3{F0: 65535, F1: -9223372036854775808}
                                                                                _ = v87
                                                                            default:
                                                                                v36 = v36
                                                                            }
                                                                        case "\n\t":
                                                                            v53 = (710 >= 840)
                                                                        default:
                                                                            v41 = append(v41, fi23)
                                                                        }
                                                                    } else {
                                                                        var v88 int32 = 2147483647
                                                                        if v53 {
                                                                            v88 = -2147483648
                                                                        }
                                                                        _ = v88
                                                                    }
                                                                    var v89 bool = false
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        v36 += -1
                                                                    }
                                                                    _ = v89
                                                                }
                                                                _ = v58
                                                            }
                                                            var v90 float64 = (((-2.25 - ((0.0 * (1.5 * 1e10)) + float64((v35 - ((v35 * v35) * v35))))) * 0.0) - (0.1 / 1.0))
                                                            if false {
                                                                _ = len(v41)
                                                                if true {
                                                                    fmt.Println(v40)
                                                                    var v91 uint8 = (^(v19))
                                                                    fmt.Println(v40, v91)
                                                                    var v92 int16 = -476
                                                                    v53 = v53
                                                                    _ = v91
                                                                    _ = v92
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        v93 := GenS12{F0: (v19 << ((553) & 63))}
                                                                        _ = v93
                                                                    case "a":
                                                                        v36 += v36
                                                                    }
                                                                    v94 := GenS13{F0: 169, F1: -515}
                                                                    var v95 float64 = float64((v35 - (v35 / ((0.1 / (1.5 * v35)) / ((v35 - 1.5) / v35)))))
                                                                    _ = v94
                                                                    _ = v95
                                                                }
                                                                if len(v41) > 0 {
                                                                    v41[0] = (fi23 / ((fi23) | 1))
                                                                }
                                                            } else {
                                                                fmt.Println(v19, fi23, v36)
                                                                var v96 uint8 = (v19 << ((1) & 63))
                                                                _ = v96
                                                            }
                                                            var v97 uint8 = v19
                                                            _ = v53
                                                            _ = v90
                                                            _ = v97
                                                        }
                                                        _ = v41
                                                    }
                                                    var v98 int16 = int16(v40)
                                                    _ = v40
                                                    _ = v98
                                                } else {
                                                    var v99 uint8 = 127
                                                    var v100 uint = 65535
                                                    fmt.Println(v99, v99, v100)
                                                    _ = v100
                                                    _ = v99
                                                }
                                            }
                                            _ = v36
                                        }
                                        v101 := []int{(fi18 % ((fi18) | 1)), 4294967295}
                                        var v102 bool = (-2.25 < 1e10)
                                        _ = v101
                                        _ = v102
                                    } else {
                                        v4 = v4
                                    }
                                    fmt.Println(fi23)
                                    _ = v35
                                case "":
                                    v19 = v19
                                case "\n\t":
                                    v19 -= v19
                                }
                                v103 := []int{fi18, fi18, fi23, 749}
                                _ = fi23
                                _ = v103
                            }
                        case "\n\t":
                            for fi104 := 0; fi104 < 10; fi104++ {
                                v105 := GenS6{F0: fi104}
                                if (v4 || (18446744073709551615 <= 1)) {
                                    v4 = (440 == 831)
                                    switch (-(fi18)) {
                                    case 42:
                                        fmt.Println(fi18, fi104)
                                    case 32768:
                                        if v4 {
                                            v106 := []int{(fi104 * fi18)}
                                            v107 := GenS14{F0: -1, F1: 9223372036854775807, F2: 2}
                                            if (9223372036854775807 != 0) {
                                                var v108 int32 = -2147483648
                                                v19 = v19
                                                v19 -= v19
                                                _ = v108
                                            } else {
                                                v109 := GenS15{F0: 0, F1: 42, F2: 4294967295}
                                                var v110 int = -932
                                                fmt.Println(v110)
                                                _ = v109
                                                _ = v110
                                            }
                                            if false {
                                                v4 = (0.0 == 3.14159)
                                            }
                                            _ = v106
                                            _ = v107
                                        } else {
                                            v111 := []int{}
                                            if len(v111) >= 2 {
                                                v111 = v111[1:len(v111)]
                                            }
                                            if len(v111) > 0 {
                                                v111[0] = 0
                                            }
                                            _ = v111
                                        }
                                    default:
                                        var v112 uint64 = 151
                                        _ = v112
                                    }
                                }
                                v113 := GenS16{F0: fi104}
                                _ = fi104
                                _ = v105
                                _ = v113
                            }
                        }
                        _ = fi18
                        _ = v19
                    }
                }
                for fi114 := 0; fi114 < 8; fi114++ {
                    v115 := []int{(fi114 | (fi114 ^ (fi114 / ((fi114) | 1)))), (fi114 ^ 32767), fi114, fi114}
                    v116 := GenS15{F0: -120, F1: 0, F2: 4294967295}
                    _ = fi114
                    _ = v115
                    _ = v116
                }
            }
            v4 = false
            for fi117 := 0; fi117 < 14; fi117++ {
                switch "a" {
                case "hello":
                    v118 := GenS9{F0: -2147483648, F1: 4294967295}
                    _ = v118
                default:
                    for fi119 := 0; fi119 < 8; fi119++ {
                        v4 = false
                        var v120 uint16 = 726
                        if v4 {
                            if ((v4 && (3.14159 >= 0.0)) || ("foo bar" != "foo bar")) {
                                v121 := []int{-1, (fi117 & -1)}
                                _ = v121
                            }
                            fmt.Println(fi117, v120)
                            var v122 float32 = 1.5
                            v122 -= 3.14159
                            _ = v122
                        } else {
                            if (true && v4) {
                                switch "hello" {
                                case "":
                                    fmt.Println(fi119, fi119, v4)
                                case "a":
                                    v123 := []int{}
                                    _ = v123
                                default:
                                    v4 = true
                                }
                            } else {
                                v120 += v120
                                var v124 string = "foo bar"
                                _ = v124
                            }
                            var v125 int = (^(fi119))
                            _ = v125
                        }
                        _ = fi119
                        _ = v120
                    }
                }
                v4 = false
                _ = fi117
            }
        }
        _ = v5
    }
    _ = v4
    return -1
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 128
    var v3 float64 = (3.14159 + ((1.5 * ((1e10 + (0.1 + ((3.14159 * 3.14159) - float64(float32(3.14159))))) * 1e10)) / 1.5))
    var v126 string = (("a" + "") + "")
    var v127 int16 = -609
    if (!true) {
        var v128 uint16 = 1
        var v129 bool = true
        _ = v128
        _ = v129
    } else {
        var v130 string = ((v126 + v126) + v126)
        var v131 uint8 = (^(v2))
        delete(v1, "hello")
        _ = v130
        _ = v131
    }
    for fi132 := 0; fi132 < 19; fi132++ {
        v3 = v3
        if (65535 <= 42) {
            switch (fi132 | fi132) {
            case -291:
                var v133 uint16 = 255
                _ = v133
            case -28:
                var v134 int64 = 573
                _ = v134
            }
            _ = len(v1)
        } else {
            v2 -= v2
            if (!true) {
                for fi135 := 0; fi135 < 4; fi135++ {
                    if true {
                        if true {
                            v136 := []int{-1, 517, 299}
                            _ = v136
                        }
                        if false {
                            v2 = v2
                        } else {
                            var v137 int64 = 9223372036854775807
                            _ = v137
                        }
                    } else {
                        if (578 < 822) {
                            v138 := GenS17{F0: 0, F1: (v2 << ((255) & 63)), F2: (v2 / ((v2) | 1))}
                            _ = v138
                        } else {
                            if (-9223372036854775808 <= 785) {
                                fmt.Println(v126, v3)
                                var v139 uint64 = 870
                                v127 = (-(v127))
                                _ = v139
                            } else {
                                var v140 float64 = 1.5
                                fmt.Println(v2, v127)
                                _ = v140
                            }
                            if true {
                                v3 = 0.1
                                if false {
                                    _ = len(v1)
                                    var v141 int8 = 42
                                    v127 = (-(v127))
                                    if (!(2147483647 == 4294967295)) {
                                        v126 = v126
                                    }
                                    _ = v141
                                }
                                v1["a"] = v126
                                if true {
                                    var v142 int16 = (v127 / (((v127 / ((v127) | 1))) | 1))
                                    switch fi135 {
                                    case 2147483647:
                                        var v143 bool = ((v142 >> ((629) & 63)) < v127)
                                        _ = v143
                                    case -279:
                                        var v144 int64 = -702
                                        _ = v144
                                    case -1:
                                        switch (v126 + v126) {
                                        case "hello":
                                            if ((true && false) && (1e10 != 1e10)) {
                                                v145 := GenS7{F0: 42}
                                                fmt.Println(v2, v3)
                                                var v146 int8 = -69
                                                v147 := GenS18{F0: 9223372036854775807, F1: int32(v146)}
                                                _ = v145
                                                _ = v146
                                                _ = v147
                                            } else {
                                                delete(v1, v126)
                                                v148 := []int{32767, 1}
                                                if false {
                                                    if (fi135 == fi132) {
                                                        if len(v148) >= 2 {
                                                            v148 = v148[1:len(v148)]
                                                        }
                                                        _ = len(v148)
                                                    } else {
                                                        v149 := v1[""]
                                                        if false {
                                                            v2 = v2
                                                            var v150 uint32 = 42
                                                            var v151 int = (fi135 & (fi135 + (fi135 << ((0) & 63))))
                                                            _ = v150
                                                            _ = v151
                                                        } else {
                                                            v149 = v126
                                                        }
                                                        _ = v149
                                                    }
                                                    v148 = append(v148, 42)
                                                } else {
                                                    v2 = 255
                                                    fmt.Println(v126)
                                                    if len(v148) >= 2 {
                                                        v148 = v148[1:len(v148)]
                                                    }
                                                }
                                                _ = v148
                                            }
                                        default:
                                            fmt.Println(fi132, fi135, v142)
                                        }
                                    }
                                    switch v126 {
                                    case "\n\t":
                                        fmt.Println(v126)
                                    default:
                                        v126 = "\n\t"
                                    }
                                    _ = v142
                                } else {
                                    switch fi135 {
                                    case 127:
                                        v126 = ("a" + (v126 + v126))
                                    default:
                                        var v152 int64 = -9223372036854775808
                                        _ = v152
                                    }
                                    v126 = v126
                                    switch fi135 {
                                    case 4294967295:
                                        v3 -= (0.1 / (0.0 - v3))
                                    case -768:
                                        switch fi132 {
                                        case 2147483647:
                                            var v153 int = 620
                                            _ = v153
                                        case -1:
                                            if false {
                                                delete(v1, "\n\t")
                                                v154 := GenS19{F0: 42}
                                                _ = v154
                                            }
                                            v155 := GenS20{F0: 42}
                                            _ = v155
                                        default:
                                            fmt.Println(v2, fi132, v126)
                                        }
                                    }
                                }
                                if true {
                                    fmt.Println(v2, v126, fi132)
                                }
                            }
                        }
                        v156 := []int{0, -881, (fi132 | fi132)}
                        _ = v156
                    }
                    fmt.Println(fi135, v126, fi132)
                    v157 := GenS4{F0: 720, F1: 9223372036854775807, F2: 40}
                    var v158 int32 = 1
                    _ = fi135
                    _ = v157
                    _ = v158
                }
                v159 := []int{fi132, (^(fi132)), fi132}
                for fi160 := 0; fi160 < 15; fi160++ {
                    v126 = v126
                    _ = fi160
                }
                _ = v159
            }
        }
        _ = fi132
    }
    switch v126 {
    case "hello":
        _ = len(v1)
    case "a":
        for fi161 := 0; fi161 < 15; fi161++ {
            v162 := v1["foo bar"]
            _ = fi161
            _ = v162
        }
    case "\n\t":
        var v163 uint8 = v2
        _ = v163
    }
    if false {
        var v164 uint8 = v2
        switch ("" + v126) {
        case "":
            v126 = ("" + v126)
        case "\n\t":
            if (((v3 / -2.25) - v3) == -2.25) {
                var v165 uint32 = 4294967295
                var v166 float64 = ((v3 / float64(0.1)) - v3)
                var v167 float32 = 0.0
                _ = v165
                _ = v166
                _ = v167
            }
        default:
            for fi168 := 0; fi168 < 16; fi168++ {
                v169 := GenS21{F0: v126}
                var v170 string = v126
                _ = fi168
                _ = v169
                _ = v170
            }
        }
        _ = v164
    }
    v126 = ""
    var v171 int16 = (v127 / ((v127) | 1))
    v127 += v171
    v3 -= v3
    v2 -= (v2 & v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v126:", v126)
    fmt.Println("v127:", v127)
    fmt.Println("v171:", v171)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v126
    _ = v127
    _ = v171
    _ = v2
    _ = v3
}
