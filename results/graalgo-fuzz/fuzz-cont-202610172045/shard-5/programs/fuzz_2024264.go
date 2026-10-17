package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint32; F1 string; F2 int32 }
type GenS13 struct { F0 bool; F1 uint64; F2 int32 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 bool; F1 string; F2 bool }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 string; F1 int16; F2 int }
type GenS18 struct { F0 string }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 int64; F1 uint; F2 uint16 }
type GenS21 struct { F0 int32; F1 int; F2 int32 }
func genFn0(p0 uint8, p1 uint) int64 {
    fmt.Println(p1, p0)
    p0 = 76
    _ = p0
    _ = p1
    return 856
}
func genFn1(p0 uint) int {
    p0 += p0
    var v6 int = -2147483648
    _ = p0
    _ = v6
    return (v6 - (v6 % ((v6) | 1)))
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 65535
    var v3 string = "hello"
    v4 := []uint32{uint32(v2)}
    var v5 int8 = 0
    v2 = (^(v2))
    v2 -= (v2 % ((v2) | 1))
    v2 += v2
    v5 -= v5
    var v7 float32 = (1.5 + 1.5)
    v3 = ""
    if true {
        if true {
            switch -2147483648 {
            case -9223372036854775808:
                var v8 int32 = 0
                _ = v8
            case 32767:
                for fi9 := 0; fi9 < 3; fi9++ {
                    var v10 float64 = 3.14159
                    _ = fi9
                    _ = v10
                }
            default:
                fmt.Println(v7)
            }
        }
    } else {
        if false {
            v2 += 194
        } else {
            v11 := GenS1{F0: 148, F1: (((3.14159 / ((0.0 - (1e10 - float64(-2.25))) + 1.5)) * (0.1 / 1.5)) >= (1.5 / float64((v7 - ((v7 / v7) - (v7 + (((v7 + v7) * 3.14159) + -2.25)))))))}
            _ = len(v1)
            _ = v11
        }
        switch (("foo bar" + "foo bar") + v3) {
        case "foo bar":
            v4 = append(v4, 2147483648)
        case "hello":
            v12 := GenS1{F0: uint32(v2), F1: (2147483647 >= 4294967295)}
            _ = v12
        case "a":
            if false {
                for fi13 := 0; fi13 < 12; fi13++ {
                    switch (fi13 >> ((v2) & 63)) {
                    case -9223372036854775808:
                        var v14 uint16 = 255
                        _ = v14
                    case 32768:
                        var v15 int16 = 1
                        _ = v15
                    }
                    var v16 float64 = (float64((v7 + v7)) / (-2.25 - 0.1))
                    for fi17 := 0; fi17 < 2; fi17++ {
                        _ = len(v1)
                        _ = len(v4)
                        var v18 uint64 = 343
                        _ = fi17
                        _ = v18
                    }
                    _ = fi13
                    _ = v16
                }
                v19 := GenS1{F0: 65535, F1: true}
                _ = v19
            }
        default:
            if false {
                var v20 int32 = 42
                v5 = (v5 & (v5 + (v5 >> ((0) & 63))))
                _ = v20
            } else {
                for fi21 := 0; fi21 < 6; fi21++ {
                    if false {
                        switch (v3 + v3) {
                        case "":
                            switch (v3 + v3) {
                            case "foo bar":
                                for fi22 := 0; fi22 < 8; fi22++ {
                                    if true {
                                        fmt.Println(v7)
                                    }
                                    _ = fi22
                                }
                            case "":
                                v2 = (v2 << ((v2) & 63))
                            case "\n\t":
                                v23 := GenS2{F0: int8(fi21)}
                                _ = v23
                            }
                        case "a":
                            v24 := GenS1{F0: 0, F1: true}
                            _ = v24
                        case "foo bar":
                            for fi25 := 0; fi25 < 8; fi25++ {
                                fmt.Println(fi25, v2)
                                _ = fi25
                            }
                        default:
                            var v26 float64 = 1.5
                            _ = v26
                        }
                        if ((v7 / v7) < (v7 * (float32(0.0) / v7))) {
                            v27 := GenS3{F0: v2}
                            var v28 float32 = v7
                            _ = v27
                            _ = v28
                        } else {
                            if true {
                                if true {
                                    for fi29 := 0; fi29 < 20; fi29++ {
                                        v30 := GenS4{F0: 42}
                                        v31 := GenS5{F0: 975}
                                        switch v3 {
                                        case "a":
                                            v3 = v3
                                        case "\n\t":
                                            if len(v4) > 0 {
                                                v4[0] = 453
                                            }
                                        }
                                        _ = fi29
                                        _ = v30
                                        _ = v31
                                    }
                                    v3 = ("\n\t" + (v3 + v3))
                                    var v32 float32 = v7
                                    _ = v32
                                } else {
                                    if (((970 > 18446744073709551615) || true) && (1.5 < v7)) {
                                        if true {
                                            v33 := GenS4{F0: (v2 | (v2 << (((v2 % (((v2 >> ((v2) & 63))) | 1))) & 63)))}
                                            var v34 uint16 = 1
                                            _ = v33
                                            _ = v34
                                        } else {
                                            fmt.Println(v3, v3)
                                        }
                                        v3 = (v3 + (v3 + v3))
                                        v7 += v7
                                    } else {
                                        var v35 string = ("a" + v3)
                                        _ = v35
                                    }
                                }
                            } else {
                                var v36 int32 = 469
                                _ = v36
                            }
                            var v37 uint = (^(v2))
                            var v38 uint32 = 0
                            _ = v37
                            _ = v38
                        }
                        var v39 int32 = -1
                        _ = v39
                    } else {
                        var v40 float32 = v7
                        v2 -= 2
                        _ = v40
                    }
                    fmt.Println(v2, v3)
                    v1[v3] = (18446744073709551615 == 2)
                    _ = fi21
                }
                v41 := GenS6{F0: v5, F1: 0}
                for fi42 := 0; fi42 < 5; fi42++ {
                    v43 := GenS3{F0: v2}
                    v44 := GenS3{F0: (v2 * (v2 | 941))}
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                _ = v41
            }
        }
    }
    switch -90 {
    case 32767:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case -763:
        if true {
            fmt.Println(v3)
            v2 = 1
            v4 = append(v4, 550)
        }
    case -625:
        fmt.Println(v3, v5)
    default:
        for fi45 := 0; fi45 < 5; fi45++ {
            if false {
                for fi46 := 0; fi46 < 17; fi46++ {
                    if true {
                        _ = len(v4)
                        if true {
                            if false {
                                if (!false) {
                                    v47 := v1[v3]
                                    if v47 {
                                        fmt.Println(v2)
                                    } else {
                                        v7 = v7
                                        fmt.Println(fi46, fi46)
                                        if v47 {
                                            var v48 uint32 = 2
                                            var v49 int = fi45
                                            _ = v48
                                            _ = v49
                                        } else {
                                            v3 = v3
                                            if (!(("" + v3) <= "hello")) {
                                                fmt.Println(v47, fi45, v5)
                                            } else {
                                                v7 -= v7
                                            }
                                        }
                                        switch "hello" {
                                        case "":
                                            v50 := GenS7{F0: genFn0(0, v2)}
                                            _ = v50
                                        case "\n\t":
                                            if (2 < 791) {
                                                switch fi46 {
                                                case 127:
                                                    var v51 bool = ((v2 % ((1) | 1)) == v2)
                                                    _ = v51
                                                case 4294967295:
                                                    var v52 uint8 = 151
                                                    _ = v52
                                                case 1:
                                                    v7 += 3.14159
                                                }
                                                if v47 {
                                                    var v53 uint64 = 434
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    switch v3 {
                                                    case "":
                                                        v54 := GenS5{F0: fi45}
                                                        _ = v54
                                                    case "\n\t":
                                                        if (fi46 == (fi46 + fi46)) {
                                                            switch fi45 {
                                                            case 0:
                                                                v55 := GenS7{F0: 9223372036854775807}
                                                                _ = v55
                                                            case 9223372036854775807:
                                                                v5 = v5
                                                            case 777:
                                                                _ = len(v4)
                                                            }
                                                            v47 = v47
                                                            switch v3 {
                                                            case "a":
                                                                delete(v1, "hello")
                                                            case "\n\t":
                                                                if ((1 > 128) || v47) {
                                                                    v56 := GenS8{F0: 42}
                                                                    if false {
                                                                        var v57 float64 = ((3.14159 + (1.5 + ((-2.25 * 0.0) + float64(1e10)))) * ((float64(v7) * (0.1 + 0.1)) + (1e10 / 1.5)))
                                                                        if v47 {
                                                                            v3 = "foo bar"
                                                                        }
                                                                        _ = v57
                                                                    } else {
                                                                        v58 := GenS2{F0: (v5 ^ int8(v53))}
                                                                        var v59 float32 = 3.14159
                                                                        _ = v58
                                                                        _ = v59
                                                                    }
                                                                    var v60 int64 = 1
                                                                    var v61 int16 = 869
                                                                    if v47 {
                                                                        v1[(v3 + "foo bar")] = ((!v47) || (fi46 != (fi45 << (((v2 >> (((^(v2))) & 63))) & 63))))
                                                                    }
                                                                    v7 += v7
                                                                    v62 := GenS4{F0: (v2 / ((v2) | 1))}
                                                                    switch 127 {
                                                                    case -1:
                                                                        v63 := GenS2{F0: (-(v5))}
                                                                        _ = v63
                                                                    case 32768:
                                                                        var v64 float32 = (v7 * 0.1)
                                                                        _ = v64
                                                                    case 0:
                                                                        switch ("a" + (v3 + v3)) {
                                                                        case "foo bar":
                                                                            var v65 int32 = 345
                                                                            _ = v65
                                                                        case "hello":
                                                                            v66 := GenS8{F0: 42}
                                                                            _ = v66
                                                                        case "":
                                                                            switch v3 {
                                                                            case "hello":
                                                                                if (318 != 2147483647) {
                                                                                    if ((v2 * v2) == 821) {
                                                                                        v67 := GenS9{F0: true}
                                                                                        if v47 {
                                                                                            if true {
                                                                                                var v68 uint = v2
                                                                                                if (-9223372036854775808 >= 4294967295) {
                                                                                                    var v69 bool = (!((v47 || (v47 && v47)) || ((^(v53)) < 18446744073709551615)))
                                                                                                    v70 := GenS4{F0: v68}
                                                                                                    _ = v69
                                                                                                    _ = v70
                                                                                                } else {
                                                                                                    var v71 uint32 = 776
                                                                                                    _ = v71
                                                                                                }
                                                                                                var v72 int8 = v5
                                                                                                _ = v68
                                                                                                _ = v72
                                                                                            }
                                                                                            v73 := GenS6{F0: v5, F1: 1}
                                                                                            _ = v73
                                                                                        } else {
                                                                                            v3 = (v3 + v3)
                                                                                            v74 := GenS5{F0: int(v61)}
                                                                                            _ = v74
                                                                                        }
                                                                                        _ = v67
                                                                                    } else {
                                                                                        switch "" {
                                                                                        case "":
                                                                                            if (42 >= genFn1((v2 ^ (v2 / ((1) | 1))))) {
                                                                                                v61 = (v61 & (v61 & v61))
                                                                                            }
                                                                                        case "\n\t":
                                                                                            fmt.Println(v60)
                                                                                        case "a":
                                                                                            v53 += 54
                                                                                        default:
                                                                                            v75 := GenS10{F0: (fi46 % ((fi45) | 1))}
                                                                                            _ = v75
                                                                                        }
                                                                                        var v76 uint = v2
                                                                                        var v77 int64 = (v60 / ((int64(fi45)) | 1))
                                                                                        _ = v76
                                                                                        _ = v77
                                                                                    }
                                                                                    if false {
                                                                                        var v78 int32 = 765
                                                                                        _ = v78
                                                                                    }
                                                                                    var v79 int8 = (-(v5))
                                                                                    _ = v79
                                                                                } else {
                                                                                    var v80 float32 = 3.14159
                                                                                    v47 = (!(v61 < v61))
                                                                                    _ = v80
                                                                                }
                                                                            case "\n\t":
                                                                                var v81 int8 = v5
                                                                                _ = v81
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v56
                                                                    _ = v60
                                                                    _ = v61
                                                                    _ = v62
                                                                } else {
                                                                    v82 := GenS6{F0: (v5 ^ v5), F1: 319}
                                                                    v7 = v7
                                                                    v53 = 120
                                                                    var v83 float32 = v7
                                                                    _ = v82
                                                                    _ = v83
                                                                }
                                                            }
                                                            switch "" {
                                                            case "hello":
                                                                v84 := GenS4{F0: 2}
                                                                _ = v84
                                                            case "foo bar":
                                                                v85 := GenS8{F0: 32767}
                                                                _ = v85
                                                            default:
                                                                fmt.Println(v53)
                                                            }
                                                            if (!(956 >= 556)) {
                                                                v4 = append(v4, 1)
                                                                fmt.Println(v7, fi45)
                                                            } else {
                                                                v3 = (("" + "\n\t") + v3)
                                                            }
                                                        } else {
                                                            v4 = append(v4, 970)
                                                        }
                                                    case "foo bar":
                                                        var v86 uint16 = 716
                                                        _ = v86
                                                    }
                                                    _ = v53
                                                } else {
                                                    v87 := GenS6{F0: 42, F1: 2}
                                                    v2 = v2
                                                    var v88 string = ("\n\t" + "")
                                                    var v89 int32 = 0
                                                    _ = v87
                                                    _ = v88
                                                    _ = v89
                                                }
                                                if (float32(((float64(1.5) / float64(v7)) / (((((((-2.25 / 1.0) + (float64(1.5) - 0.1)) * -2.25) * ((-2.25 / ((1e10 - 0.0) + ((-2.25 / 1.0) * ((0.0 * (3.14159 * float64(v7))) + 3.14159)))) * 0.0)) + 0.1) / (1e10 * 0.1)) + 0.0))) != v7) {
                                                    v1["hello"] = (42 >= -9223372036854775808)
                                                    switch genFn1((v2 << ((v2) & 63))) {
                                                    case 32768:
                                                        v90 := GenS11{F0: 1}
                                                        _ = v90
                                                    case 991:
                                                        var v91 int16 = 32767
                                                        _ = v91
                                                    case 1:
                                                        if v47 {
                                                            v4 = append(v4, 2147483648)
                                                        }
                                                    }
                                                } else {
                                                    v47 = v47
                                                }
                                            } else {
                                                v7 -= v7
                                            }
                                        default:
                                            delete(v1, ("foo bar" + "\n\t"))
                                        }
                                        v92 := GenS10{F0: (-(fi46))}
                                        v93 := GenS12{F0: 0, F1: v3, F2: 973}
                                        _ = v92
                                        _ = v93
                                    }
                                    fmt.Println(v5)
                                    if true {
                                        v94 := GenS10{F0: genFn1(524)}
                                        _ = v94
                                    } else {
                                        switch v3 {
                                        case "hello":
                                            var v95 float32 = 1e10
                                            _ = v95
                                        default:
                                            v96 := GenS10{F0: (fi45 | (fi46 << ((0) & 63)))}
                                            _ = v96
                                        }
                                        var v97 int8 = v5
                                        v3 = v3
                                        _ = v97
                                    }
                                    var v98 int64 = 42
                                    _ = v47
                                    _ = v98
                                } else {
                                    fmt.Println(v7, fi46)
                                    var v99 uint32 = 856
                                    _ = v99
                                }
                                switch (^(fi46)) {
                                case 0:
                                    var v100 uint64 = 99
                                    _ = v100
                                case 42:
                                    if (false || false) {
                                        var v101 int64 = 619
                                        if ((v101 * 0) != v101) {
                                            v4 = append(v4, 65535)
                                            v7 -= v7
                                        } else {
                                            var v102 int32 = 42
                                            v4 = append(v4, 532)
                                            _ = v102
                                        }
                                        v3 = v3
                                        _ = v101
                                    } else {
                                        v103 := GenS13{F0: false, F1: 2, F2: -2147483648}
                                        var v104 uint8 = 255
                                        _ = v103
                                        _ = v104
                                    }
                                }
                                var v105 int = (fi46 << (((v2 & (v2 / (((v2 - v2)) | 1)))) & 63))
                                if true {
                                    var v106 bool = (genFn1(v2) != -705)
                                    _ = v106
                                } else {
                                    v3 = v3
                                }
                                _ = v105
                            } else {
                                if (1 != 381) {
                                    fmt.Println(v2, fi45)
                                }
                                v2 = (v2 * (v2 ^ (v2 + (v2 % (((v2 % ((v2) | 1))) | 1)))))
                                switch ("a" + v3) {
                                case "a":
                                    var v107 int8 = (v5 | v5)
                                    _ = v107
                                case "\n\t":
                                    var v108 float64 = 3.14159
                                    _ = v108
                                case "hello":
                                    var v109 int = genFn1((v2 % ((v2) | 1)))
                                    _ = v109
                                default:
                                    var v110 string = (((v3 + "") + ("\n\t" + "\n\t")) + v3)
                                    _ = v110
                                }
                            }
                        } else {
                            switch "a" {
                            case "\n\t":
                                switch (v3 + (("\n\t" + "hello") + "")) {
                                case "a":
                                    if false {
                                        v111 := GenS1{F0: 974, F1: (true || true)}
                                        _ = v111
                                    }
                                case "foo bar":
                                    if false {
                                        v7 += v7
                                        var v112 int64 = 0
                                        _ = v112
                                    }
                                default:
                                    var v113 int32 = -798
                                    if ((3.14159 <= (v7 + (v7 - ((v7 * v7) / v7)))) && (false || (1.5 < (0.0 + -2.25)))) {
                                        _ = len(v1)
                                    } else {
                                        var v114 int64 = genFn0(234, 2)
                                        var v115 int64 = (v114 - genFn0(181, (v2 << ((v2) & 63))))
                                        var v116 float64 = 3.14159
                                        _ = v114
                                        _ = v115
                                        _ = v116
                                    }
                                    _ = v113
                                }
                            case "hello":
                                v7 += 3.14159
                            default:
                                var v117 uint16 = 511
                                _ = v117
                            }
                        }
                    }
                    _ = fi46
                }
                for fi118 := 0; fi118 < 7; fi118++ {
                    var v119 int32 = 42
                    var v120 uint16 = 42
                    var v121 float32 = v7
                    _ = fi118
                    _ = v119
                    _ = v120
                    _ = v121
                }
            }
            _ = fi45
        }
    }
    switch -322 {
    case 296:
        switch 42 {
        case -1:
            if false {
                var v122 float64 = 1e10
                if false {
                    var v123 uint32 = 1
                    _ = v123
                }
                var v124 uint32 = 65535
                _ = v122
                _ = v124
            } else {
                v7 += v7
                v4 = append(v4, 2147483648)
                v125 := GenS12{F0: 255, F1: (("" + v3) + v3), F2: 42}
                _ = v125
            }
        case 690:
            if (9223372036854775807 != -9223372036854775808) {
                v2 = 4294967295
                fmt.Println(v2, v5, v7)
            } else {
                var v126 bool = (79 != (-(v5)))
                var v127 int = 1
                fmt.Println(v2, v127)
                _ = v126
                _ = v127
            }
        case 32768:
            v2 = v2
        default:
            var v128 uint64 = 42
            _ = v128
        }
    case 127:
        for fi129 := 0; fi129 < 15; fi129++ {
            var v130 uint64 = 1
            _ = fi129
            _ = v130
        }
    }
    for fi131 := 0; fi131 < 5; fi131++ {
        for fi132 := 0; fi132 < 4; fi132++ {
            if true {
                if false {
                    var v133 float64 = (1e10 * 0.1)
                    _ = v133
                }
            } else {
                var v134 uint32 = 4294967295
                v7 += 0.0
                _ = v134
            }
            _ = fi132
        }
        v135 := GenS14{F0: 32767}
        v5 = (v5 / (((v5 * (-(v5)))) | 1))
        _ = fi131
        _ = v135
    }
    var v136 int16 = -32768
    v136 = v136
    v1[v3] = false
    for fi137 := 0; fi137 < 11; fi137++ {
        fmt.Println(v5, v2, v5)
        var v138 uint64 = 125
        _ = fi137
        _ = v138
    }
    for fi139 := 0; fi139 < 17; fi139++ {
        var v140 int32 = -435
        v140 += int32(fi139)
        var v141 uint8 = 1
        _ = fi139
        _ = v140
        _ = v141
    }
    for fi142 := 0; fi142 < 13; fi142++ {
        for fi143 := 0; fi143 < 3; fi143++ {
            if (v5 < -128) {
                fmt.Println(v5)
                fmt.Println(v3, v7, v3)
                v7 = (((v7 / 1.0) / v7) / -2.25)
                if true {
                    fmt.Println(v7)
                    _ = len(v4)
                }
            } else {
                var v144 int64 = 9223372036854775807
                _ = v144
            }
            if (-307 >= genFn0(42, v2)) {
                if false {
                    if true {
                        v7 += (v7 / v7)
                        if (!false) {
                            var v145 uint = (v2 | v2)
                            if true {
                                v146 := GenS15{F0: false, F1: "a", F2: false}
                                var v147 uint = v2
                                if true {
                                    if len(v4) > 0 {
                                        v4[0] = 44
                                    }
                                    var v148 int64 = -783
                                    _ = v148
                                } else {
                                    switch (fi142 * 503) {
                                    case -1:
                                        fmt.Println(v5, v7)
                                    case 273:
                                        v149 := GenS16{F0: -1}
                                        _ = v149
                                    case 42:
                                        if len(v4) > 0 {
                                            v4[0] = 2
                                        }
                                    default:
                                        switch fi142 {
                                        case -985:
                                            if len(v4) > 0 {
                                                v4[0] = 2147483647
                                            }
                                        case -2147483648:
                                            if false {
                                                var v150 uint32 = 65535
                                                var v151 uint = v2
                                                _ = v150
                                                _ = v151
                                            }
                                        }
                                    }
                                    var v152 int64 = 0
                                    _ = v152
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v1["a"] = false
                                fmt.Println(v136, v3)
                                _ = v146
                                _ = v147
                            }
                            _ = v145
                        }
                        var v153 bool = ((v3 + "hello") <= (v3 + v3))
                        if (v3 > "a") {
                            if v153 {
                                var v154 float64 = float64((v7 / float32(1.5)))
                                switch fi142 {
                                case 42:
                                    var v155 int8 = v5
                                    _ = v155
                                default:
                                    switch (fi143 - fi143) {
                                    case 4294967295:
                                        if ((false || (v136 != (v136 + -32768))) || false) {
                                            v154 = v154
                                            switch v3 {
                                            case "\n\t":
                                                switch "hello" {
                                                case "":
                                                    v156 := GenS8{F0: 0}
                                                    _ = v156
                                                case "foo bar":
                                                    fmt.Println(v3, v153)
                                                case "\n\t":
                                                    v157 := GenS17{F0: v3, F1: v136, F2: (fi143 >> ((255) & 63))}
                                                    _ = v157
                                                }
                                            case "":
                                                v3 = "\n\t"
                                            case "hello":
                                                v158 := GenS15{F0: false, F1: "foo bar", F2: ((-(fi143)) > genFn1(479))}
                                                _ = v158
                                            }
                                            fmt.Println(v136, v2, fi142)
                                            v159 := GenS18{F0: v3}
                                            _ = v159
                                        }
                                    case -443:
                                        fmt.Println(v3, v154, fi142)
                                    case 127:
                                        v153 = (236 <= 225)
                                    default:
                                        if len(v4) > 0 {
                                            v4[0] = 42
                                        }
                                    }
                                    var v160 float64 = float64(-2.25)
                                    fmt.Println(v5)
                                    _ = v160
                                }
                                _ = v154
                            }
                            if true {
                                v5 += (v5 << ((v2) & 63))
                                if false {
                                    v7 -= float32(3.14159)
                                } else {
                                    v4 = append(v4, 2147483647)
                                    var v161 float32 = (v7 / (v7 - (v7 / v7)))
                                    _ = v161
                                }
                            }
                        }
                        v162 := GenS19{F0: v136}
                        _ = v153
                        _ = v162
                    } else {
                        v2 += (v2 * (v2 + v2))
                        if (0 != 0) {
                            if (1 >= 2147483647) {
                                var v163 float64 = (1.5 + (1e10 + (-2.25 - 3.14159)))
                                switch (fi143 ^ fi143) {
                                case 619:
                                    var v164 int8 = (v5 % (((v5 & v5)) | 1))
                                    var v165 uint64 = 406
                                    v5 = -106
                                    _ = v164
                                    _ = v165
                                case 32767:
                                    if true {
                                        if (!(179 <= 2147483648)) {
                                            var v166 int64 = 9223372036854775807
                                            v167 := GenS2{F0: (^(v5))}
                                            var v168 uint32 = 1
                                            _ = v166
                                            _ = v167
                                            _ = v168
                                        }
                                        fmt.Println(fi142, fi143)
                                        v169 := GenS20{F0: 233, F1: 1, F2: 42}
                                        _ = v169
                                    }
                                default:
                                    v5 = (-(v5))
                                }
                                _ = v163
                            } else {
                                if false {
                                    v5 += v5
                                    switch (fi142 >> ((v2) & 63)) {
                                    case -2147483648:
                                        var v170 int8 = v5
                                        _ = v170
                                    case 32767:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    default:
                                        if true {
                                            v171 := GenS16{F0: -1}
                                            _ = v171
                                        }
                                    }
                                }
                                v172 := GenS21{F0: 2147483647, F1: (fi143 + (fi143 >> ((v2) & 63))), F2: -1}
                                var v173 uint64 = 2
                                _ = v172
                                _ = v173
                            }
                        }
                        var v174 int16 = v136
                        _ = v174
                    }
                    if true {
                        v175 := GenS1{F0: 0, F1: false}
                        _ = v175
                    }
                } else {
                    var v176 string = ((("\n\t" + v3) + ((v3 + "a") + (v3 + v3))) + "foo bar")
                    _ = v176
                }
            }
            var v177 int32 = 2147483647
            v7 = v7
            v1["foo bar"] = true
            _ = fi143
            _ = v177
        }
        v3 = "foo bar"
        _ = fi142
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
