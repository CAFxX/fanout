package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 string; F2 int }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int32; F1 bool; F2 int }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int64; F1 uint64 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint16; F1 int8; F2 int16 }
type GenS16 struct { F0 int32; F1 int16; F2 uint }
func genFn0(p0 int64, p1 uint16) int16 {
    var v4 float32 = 1e10
    var v5 bool = true
    v4 = v4
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint32 = 718
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 480
}
func genFn1(p0 uint32, p1 uint8) int {
    var v8 uint32 = p0
    if ((!false) || (0 != 0)) {
        v9 := []int{454, 0, 617, -512}
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 uint8 = p1
            v8 += p0
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    v12 := []int{}
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 uint = uint(v8)
        if (42 == 147) {
            p0 = p0
            var v15 int = fi13
            _ = v15
        } else {
            v8 -= (^(p0))
            var v16 uint64 = 18446744073709551615
            var v17 float64 = 0.1
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 15; fi18++ {
            v19 := GenS1{F0: 367}
            _ = fi18
            _ = v19
        }
        _ = fi13
        _ = v14
    }
    var v20 int32 = int32(v8)
    _ = p0
    _ = p1
    _ = v12
    _ = v20
    _ = v8
    return 51
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 2
    v3 := make(map[string]int)
    v21 := GenS1{F0: 244}
    var v22 int16 = 634
    if (2 != 18446744073709551615) {
        for fi23 := 0; fi23 < 9; fi23++ {
            switch "foo bar" {
            case "foo bar":
                if (1 >= 1) {
                    v22 = v22
                }
            case "":
                if true {
                    delete(v1, fi23)
                    var v24 int8 = -80
                    v22 += v22
                    _ = v24
                }
            }
            v25 := []int{(fi23 & 32767), (fi23 | -9223372036854775808), fi23, genFn1(199, 168)}
            if (!(fi23 != 4294967295)) {
                fmt.Println(fi23)
                v22 += -1
                v2 -= (v2 % ((v2) | 1))
            }
            _ = fi23
            _ = v25
        }
        fmt.Println(v2, v22)
    } else {
        v26 := GenS1{F0: (v22 ^ 1000)}
        _ = v26
    }
    fmt.Println(v2, v22)
    if true {
        fmt.Println(v2)
        var v27 int8 = -1
        _ = v27
    }
    v22 = (v22 & 42)
    v28 := v3["\n\t"]
    v29 := GenS2{F0: -2147483648, F1: (("a" + ("a" + "\n\t")) + ("a" + "")), F2: 127}
    v30 := []int{}
    v31 := GenS1{F0: genFn0(42, uint16(v2))}
    var v32 int16 = (v22 | -32768)
    switch "foo bar" {
    case "foo bar":
        if true {
            var v33 float64 = 1.5
            var v34 int64 = -1
            switch ("foo bar" + "a") {
            case "hello":
                v35 := GenS3{F0: 9223372036854775807, F1: -128}
                _ = v35
            case "foo bar":
                fmt.Println(v2, v28, v22)
            default:
                v33 += (-2.25 / ((v33 / v33) / ((((v33 / float64(float32(1e10))) * (((1.5 + v33) / 1.0) / 0.1)) + -2.25) / ((1.5 - v33) / float64((float32((3.14159 * (0.0 / v33))) * 1.5))))))
            }
            _ = v33
            _ = v34
        } else {
            for fi36 := 0; fi36 < 10; fi36++ {
                var v37 float32 = 1.5
                v2 = (v2 | v2)
                if ((9223372036854775807 == 9223372036854775807) || (!(2 < 30))) {
                    v28 = (fi36 - (fi36 ^ (-(v28))))
                    v38 := GenS2{F0: 445, F1: "a", F2: v28}
                    var v39 int8 = 71
                    _ = v38
                    _ = v39
                } else {
                    var v40 int8 = 109
                    var v41 int8 = v40
                    for fi42 := 0; fi42 < 9; fi42++ {
                        var v43 int64 = -9223372036854775808
                        v2 = (^(v2))
                        _ = fi42
                        _ = v43
                    }
                    _ = v40
                    _ = v41
                }
                _ = fi36
                _ = v37
            }
            v44 := GenS2{F0: 500, F1: ("\n\t" + "a"), F2: (v28 << ((4294967295) & 63))}
            _ = v44
        }
    case "\n\t":
        switch "foo bar" {
        case "":
            if ((v2 << ((222) & 63)) > v2) {
                var v45 int32 = -266
                _ = v45
            } else {
                for fi46 := 0; fi46 < 2; fi46++ {
                    v47 := GenS2{F0: 0, F1: ("" + (("a" + "") + "\n\t")), F2: fi46}
                    if (false && (506 <= 1)) {
                        if (-2.25 != (0.1 * 0.0)) {
                            v28 -= fi46
                            for fi48 := 0; fi48 < 9; fi48++ {
                                v49 := v1[v28]
                                _ = fi48
                                _ = v49
                            }
                        }
                        v28 = -280
                        delete(v3, "\n\t")
                    } else {
                        if true {
                            v50 := v1[(fi46 - (v28 + (fi46 ^ 2147483647)))]
                            _ = v50
                        }
                    }
                    fmt.Println(fi46, v28, v32)
                    _ = fi46
                    _ = v47
                }
                var v51 uint16 = uint16(v2)
                switch ("\n\t" + ("\n\t" + ("hello" + "a"))) {
                case "\n\t":
                    if false {
                        fmt.Println(v28)
                        for fi52 := 0; fi52 < 17; fi52++ {
                            v2 = 1
                            if true {
                                v32 -= (v22 - genFn0(632, (v51 + v51)))
                                for fi53 := 0; fi53 < 4; fi53++ {
                                    v54 := GenS2{F0: -11, F1: ("a" + "foo bar"), F2: 128}
                                    if (2 < 2) {
                                        switch "a" {
                                        case "hello":
                                            v22 += genFn0(-9223372036854775808, (v51 | (v51 >> ((65535) & 63))))
                                        case "":
                                            v55 := GenS3{F0: 9223372036854775807, F1: 127}
                                            _ = v55
                                        case "foo bar":
                                            var v56 int16 = (v22 * (v22 >> ((18446744073709551615) & 63)))
                                            _ = v56
                                        default:
                                            v51 -= (^(v51))
                                        }
                                        var v57 uint16 = (v51 * 64)
                                        _ = v57
                                    } else {
                                        if true {
                                            v32 += 1000
                                            if false {
                                                var v58 uint = 0
                                                var v59 float64 = (((0.1 + ((float64(0.0) + 0.0) - (3.14159 + float64(3.14159)))) + 3.14159) / 1e10)
                                                _ = v58
                                                _ = v59
                                            } else {
                                                if (!(!false)) {
                                                    var v60 bool = (42 > (^(v2)))
                                                    var v61 int32 = 470
                                                    if false {
                                                        v22 = genFn0(9223372036854775807, (v51 ^ v51))
                                                        var v62 float64 = (0.1 * -2.25)
                                                        var v63 int16 = int16(v61)
                                                        _ = v62
                                                        _ = v63
                                                    } else {
                                                        fmt.Println(v51, v32)
                                                        var v64 float64 = 1e10
                                                        fmt.Println(v28)
                                                        _ = v64
                                                    }
                                                    _ = v60
                                                    _ = v61
                                                } else {
                                                    v65 := GenS2{F0: 2147483647, F1: ("hello" + ""), F2: (fi52 ^ (^(fi53)))}
                                                    switch -392 {
                                                    case 536:
                                                        if false {
                                                            if len(v30) > 0 {
                                                                v30[0] = (^(fi53))
                                                            }
                                                            fmt.Println(v51, v32)
                                                        }
                                                    }
                                                    _ = v65
                                                }
                                                if true {
                                                    if false {
                                                        if false {
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                var v66 int64 = -1
                                                                _ = v66
                                                            case "":
                                                                v1[int(v32)] = genFn1(42, 2)
                                                            default:
                                                                if false {
                                                                    v67 := GenS4{F0: 42}
                                                                    _ = len(v30)
                                                                    _ = v67
                                                                } else {
                                                                    v68 := GenS4{F0: 42}
                                                                    _ = v68
                                                                }
                                                            }
                                                        } else {
                                                            var v69 int32 = 642
                                                            v51 = (^(v51))
                                                            if true {
                                                                _ = len(v30)
                                                            } else {
                                                                v2 += (v2 >> ((4294967295) & 63))
                                                            }
                                                            if true {
                                                                var v70 float64 = 3.14159
                                                                fmt.Println(v32, v2, v22)
                                                                _ = v70
                                                            }
                                                            _ = v69
                                                        }
                                                    }
                                                }
                                                var v71 float64 = ((1.5 + (0.1 + 0.1)) / 0.1)
                                                _ = v71
                                            }
                                            v51 -= v51
                                            _ = len(v3)
                                            if (((v32 / ((0) | 1)) < 42) && (-798 != (v32 / ((genFn0(0, (^(v51)))) | 1)))) {
                                                var v72 bool = true
                                                _ = v72
                                            } else {
                                                var v73 uint8 = 11
                                                _ = v73
                                            }
                                        } else {
                                            if false {
                                                var v74 uint16 = 1
                                                _ = v74
                                            }
                                            switch (v28 ^ (v28 + v28)) {
                                            case -603:
                                                fmt.Println(v32)
                                            case 128:
                                                v32 = v22
                                            }
                                            if false {
                                                if false {
                                                    if ((v32 & 32767) != (^(v32))) {
                                                        if false {
                                                            switch (fi53 & genFn1(2147483647, 0)) {
                                                            case 42:
                                                                v22 += genFn0(-1, v51)
                                                            }
                                                            if (false || false) {
                                                                var v75 bool = true
                                                                var v76 string = "\n\t"
                                                                if (((-2.25 * 1e10) - 1.5) == (0.0 * (float32(-2.25) - ((1.5 * 0.1) - 0.1)))) {
                                                                    if (int64(v32) >= -9223372036854775808) {
                                                                        v77 := GenS5{F0: -165, F1: 42}
                                                                        _ = v77
                                                                    }
                                                                    v32 += (^(v22))
                                                                    v2 += v2
                                                                }
                                                                var v78 int32 = -1
                                                                _ = v75
                                                                _ = v76
                                                                _ = v78
                                                            } else {
                                                                v79 := GenS4{F0: 55}
                                                                _ = v79
                                                            }
                                                        }
                                                    }
                                                    var v80 uint8 = 255
                                                    v2 -= 918
                                                    if true {
                                                        var v81 float32 = ((((float32(1.5) + (1.5 / float32((1.5 + (float64(1e10) + (0.1 + 3.14159)))))) + 0.0) * 1e10) / 0.1)
                                                        v82 := GenS2{F0: -2147483648, F1: "\n\t", F2: v28}
                                                        _ = v81
                                                        _ = v82
                                                    } else {
                                                        v22 = v32
                                                        v83 := GenS1{F0: 1}
                                                        var v84 uint16 = v51
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                    if true {
                                                        if true {
                                                            if ((false && true) || true) {
                                                                v85 := GenS4{F0: 1}
                                                                v51 -= uint16(v2)
                                                                _ = v85
                                                            }
                                                        } else {
                                                            v86 := GenS5{F0: 9223372036854775807, F1: (^(v51))}
                                                            if true {
                                                                v22 += (v32 + v22)
                                                            }
                                                            fmt.Println(v32, fi53)
                                                            _ = v86
                                                        }
                                                        v87 := GenS2{F0: -2147483648, F1: ("hello" + ("hello" + "hello")), F2: genFn1(2147483648, v80)}
                                                        switch 42 {
                                                        case 128:
                                                            if (true || true) {
                                                                var v88 uint = 246
                                                                _ = v88
                                                            } else {
                                                                if false {
                                                                    fmt.Println(v2, v22, fi53)
                                                                    var v89 bool = true
                                                                    var v90 int16 = (v32 >> ((890) & 63))
                                                                    _ = v89
                                                                    _ = v90
                                                                } else {
                                                                    v51 -= 32767
                                                                    var v91 int64 = -1
                                                                    var v92 float32 = ((1.5 * 1.5) + -2.25)
                                                                    switch (^(fi53)) {
                                                                    case 32768:
                                                                        fmt.Println(v28)
                                                                    }
                                                                    _ = v91
                                                                    _ = v92
                                                                }
                                                            }
                                                        case -564:
                                                            v51 -= 45
                                                        case 1:
                                                            var v93 float64 = float64(1e10)
                                                            _ = v93
                                                        }
                                                        v30 = append(v30, fi53)
                                                        if true {
                                                            v2 = 0
                                                            if (v2 == (v2 * v2)) {
                                                                v94 := GenS6{F0: true}
                                                                _ = v94
                                                            } else {
                                                                if (2 < 2) {
                                                                    switch v28 {
                                                                    case 9223372036854775807:
                                                                        if false {
                                                                            switch ("a" + "hello") {
                                                                            case "hello":
                                                                                if true {
                                                                                    var v95 string = ""
                                                                                    var v96 uint16 = (v51 & v51)
                                                                                    _ = len(v30)
                                                                                    _ = v95
                                                                                    _ = v96
                                                                                }
                                                                                if false {
                                                                                    var v97 float64 = (1.5 - (((0.1 + (0.0 / 1.0)) * float64(float32(((3.14159 / (3.14159 / ((1e10 / 1e10) / 1.5))) * (((-2.25 - 0.1) + 3.14159) / float64(((float32(0.1) * (1e10 * 3.14159)) * -2.25))))))) - 0.0))
                                                                                    var v98 int64 = 103
                                                                                    var v99 int32 = 0
                                                                                    _ = v97
                                                                                    _ = v98
                                                                                    _ = v99
                                                                                }
                                                                            case "":
                                                                                if (true && false) {
                                                                                    _ = len(v30)
                                                                                }
                                                                            case "\n\t":
                                                                                var v100 uint64 = 0
                                                                                _ = v100
                                                                            }
                                                                        }
                                                                    case 4294967295:
                                                                        v51 -= v51
                                                                    default:
                                                                        var v101 uint = 255
                                                                        _ = v101
                                                                    }
                                                                    if false {
                                                                        switch (v28 * v28) {
                                                                        case 983:
                                                                            if true {
                                                                                switch fi52 {
                                                                                case 4294967295:
                                                                                    v22 += v22
                                                                                case -1:
                                                                                    var v102 uint = 417
                                                                                    _ = v102
                                                                                case 782:
                                                                                    if (true && true) {
                                                                                        switch fi52 {
                                                                                        case -405:
                                                                                            fmt.Println(fi53, v2, v51)
                                                                                        case -1:
                                                                                            if false {
                                                                                                var v103 int16 = (v32 + int16(fi53))
                                                                                                fmt.Println(v80)
                                                                                                if ((!false) && (1e10 <= 3.14159)) {
                                                                                                    v3["a"] = (fi52 | -440)
                                                                                                    switch (-(v28)) {
                                                                                                    case -354:
                                                                                                        switch "hello" {
                                                                                                        case "":
                                                                                                            fmt.Println(v28)
                                                                                                            v104 := GenS7{F0: -1, F1: (^(v32))}
                                                                                                            v2 = 879
                                                                                                            _ = v104
                                                                                                        default:
                                                                                                            v22 = (v22 / ((v103) | 1))
                                                                                                        }
                                                                                                    case 0:
                                                                                                        fmt.Println(fi52, v22, v103)
                                                                                                    default:
                                                                                                        if len(v30) >= 2 {
                                                                                                            v30 = v30[1:len(v30)]
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                _ = v103
                                                                                            } else {
                                                                                                v51 = v51
                                                                                                if len(v30) > 0 {
                                                                                                    v30[0] = -9223372036854775808
                                                                                                }
                                                                                                var v105 int32 = 1
                                                                                                _ = v105
                                                                                            }
                                                                                        case 2147483647:
                                                                                            v106 := GenS8{F0: 108}
                                                                                            _ = v106
                                                                                        default:
                                                                                            v107 := GenS1{F0: (v32 / (((v32 & -32768)) | 1))}
                                                                                            _ = v107
                                                                                        }
                                                                                    } else {
                                                                                        v51 = (v51 % (((v51 << ((255) & 63))) | 1))
                                                                                    }
                                                                                default:
                                                                                    var v108 uint8 = (v80 / ((54) | 1))
                                                                                    _ = v108
                                                                                }
                                                                                var v109 int8 = 64
                                                                                _ = v109
                                                                            }
                                                                        case 32768:
                                                                            _ = len(v30)
                                                                        }
                                                                    } else {
                                                                        var v110 string = ("\n\t" + ("a" + ("foo bar" + ("foo bar" + "hello"))))
                                                                        v111 := GenS9{F0: (v80 << ((473) & 63))}
                                                                        v32 += (-(v32))
                                                                        _ = v110
                                                                        _ = v111
                                                                    }
                                                                    v80 = (v80 + v80)
                                                                }
                                                            }
                                                            v1[(^(fi53))] = v28
                                                            switch "hello" {
                                                            case "":
                                                                v51 += (^(v51))
                                                            case "\n\t":
                                                                v112 := GenS10{F0: -1, F1: (((v80 % ((v80) | 1)) >= 69) && (true && (float32(1.5) <= ((1e10 + 1e10) / (0.0 + 1e10))))), F2: -1}
                                                                _ = v112
                                                            case "hello":
                                                                var v113 float32 = float32((1.5 - float64(0.1)))
                                                                _ = v113
                                                            default:
                                                                switch 1 {
                                                                case -9223372036854775808:
                                                                    v114 := GenS10{F0: int32(v2), F1: false, F2: fi53}
                                                                    _ = v114
                                                                case 1:
                                                                    var v115 uint16 = v51
                                                                    _ = v115
                                                                case 42:
                                                                    var v116 uint8 = (^(v80))
                                                                    _ = v116
                                                                default:
                                                                    v117 := GenS11{F0: 189}
                                                                    _ = v117
                                                                }
                                                            }
                                                        }
                                                        _ = v87
                                                    }
                                                    _ = v80
                                                } else {
                                                    if ((^(v51)) == 255) {
                                                        if len(v30) > 0 {
                                                            v30[0] = 9223372036854775807
                                                        }
                                                        v118 := v1[128]
                                                        _ = v118
                                                    } else {
                                                        _ = len(v3)
                                                        var v119 int16 = 460
                                                        v2 -= v2
                                                        _ = v119
                                                    }
                                                    v120 := GenS12{F0: -1, F1: (v2 | v2)}
                                                    _ = v120
                                                }
                                                delete(v1, fi53)
                                                v28 -= v28
                                            } else {
                                                switch 32767 {
                                                case 252:
                                                    v32 -= genFn0(0, 42)
                                                case 32768:
                                                    v2 -= (v2 * v2)
                                                    v121 := GenS4{F0: 0}
                                                    var v122 int32 = -2147483648
                                                    _ = v121
                                                    _ = v122
                                                case 789:
                                                    v123 := GenS3{F0: 9223372036854775807, F1: 53}
                                                    _ = v123
                                                }
                                                if true {
                                                    if (genFn0(-1, 2) > (-(v22))) {
                                                        v2 = v2
                                                    } else {
                                                        var v124 float32 = 0.0
                                                        _ = v124
                                                    }
                                                    var v125 uint16 = v51
                                                    _ = v125
                                                }
                                                v126 := GenS13{F0: 244}
                                                var v127 int64 = -131
                                                _ = v126
                                                _ = v127
                                            }
                                        }
                                        v2 += (v2 | (v2 * v2))
                                    }
                                    v51 = (v51 * (^(v51)))
                                    var v128 int16 = v32
                                    _ = fi53
                                    _ = v128
                                    _ = v54
                                }
                            }
                            _ = fi52
                        }
                    } else {
                        fmt.Println(v51, v2, v22)
                        var v129 float64 = ((0.0 / 1.5) * (0.1 + -2.25))
                        _ = v129
                    }
                case "foo bar":
                    var v130 int16 = genFn0(974, v51)
                    _ = v130
                default:
                    v131 := GenS14{F0: v28}
                    _ = v131
                }
                _ = v51
            }
        case "a":
            var v132 float32 = 0.1
            _ = v132
        default:
            v2 += (^(v2))
        }
    case "":
        if true {
            v133 := GenS12{F0: 659, F1: 296}
            v30 = append(v30, (^(v28)))
            _ = v133
        } else {
            v134 := GenS14{F0: v28}
            var v135 int16 = (v22 / (((v32 / ((int16(v2)) | 1))) | 1))
            _ = v134
            _ = v135
        }
    default:
        var v136 int16 = v32
        _ = v136
    }
    v137 := GenS15{F0: 617, F1: 0, F2: v22}
    delete(v1, (v28 - genFn1(42, 127)))
    if true {
        v138 := GenS8{F0: 120}
        v139 := GenS16{F0: -343, F1: (v22 ^ (-(v32))), F2: 1}
        delete(v3, ("a" + ((("\n\t" + "hello") + ("foo bar" + "a")) + ("\n\t" + "\n\t"))))
        _ = v138
        _ = v139
    }
    var v140 uint64 = 18446744073709551615
    var v141 int32 = -2147483648
    var v142 string = (("hello" + "\n\t") + ("foo bar" + "foo bar"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v140:", v140)
    fmt.Println("v141:", v141)
    fmt.Println("v142:", v142)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v137
    _ = v140
    _ = v141
    _ = v142
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
