package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 string; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 bool; F1 int32; F2 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32; F1 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 bool; F1 int; F2 uint }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int32; F1 uint32; F2 int }
type GenS12 struct { F0 int; F1 int8; F2 int8 }

func main() {
    var v1 string = "hello"
    var v2 int32 = -819
    v3 := []int{4294967295, 762, 9223372036854775807}
    v4 := make(map[uint32]int)
    fmt.Println(v2, v2)
    var v5 float64 = 0.0
    for fi6 := 0; fi6 < 8; fi6++ {
        if ((v1 > v1) || false) {
            fmt.Println(v1)
            var v7 int8 = -128
            v1 = v1
            _ = v7
        }
        _ = fi6
    }
    if true {
        var v8 int64 = 1
        switch v1 {
        case "foo bar":
            v3 = append(v3, 32768)
        case "hello":
            var v9 uint8 = 2
            _ = v9
        case "":
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 uint32 = 650
                fmt.Println(fi10)
                delete(v4, v11)
                _ = fi10
                _ = v11
            }
        default:
            _ = len(v3)
        }
        _ = v8
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        if (1 != 1) {
            v13 := GenS1{F0: true}
            v14 := GenS2{F0: false, F1: "\n\t", F2: 42}
            _ = v13
            _ = v14
        }
        switch int(v2) {
        case 129:
            v15 := GenS3{F0: 2, F1: 111}
            _ = v15
        case 9223372036854775807:
            for fi16 := 0; fi16 < 4; fi16++ {
                v3 = append(v3, fi16)
                if false {
                    switch (fi16 >> ((2) & 63)) {
                    case 0:
                        switch v1 {
                        case "foo bar":
                            v2 = -154
                        default:
                            fmt.Println(v1, fi16, v1)
                        }
                    case 32767:
                        var v17 int8 = -1
                        _ = v17
                    case 32768:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    default:
                        switch fi16 {
                        case -2147483648:
                            v1 = ""
                        case 42:
                            v3 = append(v3, fi16)
                        case 1:
                            var v18 float64 = (v5 * ((v5 * v5) / (3.14159 * v5)))
                            _ = v18
                        }
                    }
                    v19 := GenS4{F0: true, F1: v2}
                    _ = v19
                }
                _ = fi16
            }
        }
        _ = fi12
    }
    var v20 int8 = -1
    if true {
        if false {
            if false {
                for fi21 := 0; fi21 < 3; fi21++ {
                    fmt.Println(v20)
                    _ = fi21
                }
                for fi22 := 0; fi22 < 18; fi22++ {
                    for fi23 := 0; fi23 < 4; fi23++ {
                        v5 += (v5 * (v5 / v5))
                        _ = len(v3)
                        v2 = (v2 + (-(v2)))
                        _ = fi23
                    }
                    if true {
                        _ = len(v3)
                        var v24 float64 = (v5 + (v5 / (1.5 - v5)))
                        _ = v24
                    } else {
                        var v25 bool = (374 > 4294967295)
                        for fi26 := 0; fi26 < 1; fi26++ {
                            delete(v4, 42)
                            if v25 {
                                if (0 <= v2) {
                                    var v27 int = 225
                                    _ = v27
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v3 = append(v3, (fi22 | (fi26 + 127)))
                                }
                                if len(v3) > 0 {
                                    v3[0] = 4294967295
                                }
                            } else {
                                v5 = v5
                                delete(v4, 65535)
                                var v28 int64 = 1
                                var v29 uint32 = 115
                                if (v25 && (65535 == 255)) {
                                    var v30 int32 = (-(v2))
                                    var v31 int64 = v28
                                    if false {
                                        v32 := GenS4{F0: (true && (42 == 127)), F1: (v30 * 42)}
                                        if true {
                                            fmt.Println(v25, v31)
                                            v1 = ("\n\t" + (("" + "hello") + v1))
                                            v33 := GenS5{F0: ((v1 + "\n\t") < ((v1 + ("hello" + v1)) + "\n\t")), F1: v30, F2: v30}
                                            _ = len(v3)
                                            v29 -= (v29 % (((v29 & v29)) | 1))
                                            _ = v33
                                        }
                                        v34 := GenS6{F0: (!(v31 < (v31 ^ 114)))}
                                        _ = v32
                                        _ = v34
                                    }
                                    _ = v30
                                    _ = v31
                                }
                                v35 := GenS5{F0: true, F1: v2, F2: v2}
                                fmt.Println(fi26, v20, v28)
                                _ = v28
                                _ = v29
                                _ = v35
                            }
                            if true {
                                if v25 {
                                    if (42 > 42) {
                                        var v36 bool = false
                                        v37 := GenS7{F0: 2, F1: (v20 / ((v20) | 1))}
                                        _ = v36
                                        _ = v37
                                    } else {
                                        switch (fi22 * (fi26 * (fi26 >> ((779) & 63)))) {
                                        case 1:
                                            var v38 uint32 = 65535
                                            var v39 string = v1
                                            if false {
                                                v40 := GenS6{F0: false}
                                                if (0 < 926) {
                                                    v41 := GenS8{F0: "a"}
                                                    v42 := v4[(v38 ^ v38)]
                                                    fmt.Println(v5, v5)
                                                    _ = v41
                                                    _ = v42
                                                } else {
                                                    if (-232 == 32767) {
                                                        if v25 {
                                                            var v43 int8 = -128
                                                            _ = v43
                                                        } else {
                                                            v44 := v4[65535]
                                                            switch v1 {
                                                            case "\n\t":
                                                                v2 = v2
                                                            case "foo bar":
                                                                var v45 int64 = 0
                                                                _ = v45
                                                            }
                                                            var v46 bool = true
                                                            _ = v44
                                                            _ = v46
                                                        }
                                                        var v47 uint32 = v38
                                                        if v25 {
                                                            var v48 float64 = v5
                                                            var v49 int8 = v20
                                                            v47 -= (v47 ^ 2147483648)
                                                            _ = v48
                                                            _ = v49
                                                        }
                                                        _ = v47
                                                    }
                                                    if (v25 && (!("hello" != v1))) {
                                                        var v50 bool = true
                                                        var v51 int32 = (v2 | -2147483648)
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                    var v52 int8 = (v20 / (((-(v20))) | 1))
                                                    _ = v52
                                                }
                                                v39 = ((("hello" + "") + "\n\t") + "")
                                                var v53 uint8 = 2
                                                _ = v40
                                                _ = v53
                                            } else {
                                                var v54 uint16 = 0
                                                v20 -= 47
                                                _ = v54
                                            }
                                            _ = v38
                                            _ = v39
                                        case 2147483647:
                                            if true {
                                                v25 = v25
                                            }
                                        }
                                        switch v1 {
                                        case "":
                                            if true {
                                                v20 = int8(fi26)
                                                var v55 int8 = (v20 >> ((625) & 63))
                                                _ = v55
                                            }
                                        case "foo bar":
                                            var v56 int16 = -476
                                            _ = v56
                                        }
                                    }
                                }
                                fmt.Println(v5, v5, v1)
                                v25 = v25
                            }
                            _ = fi26
                        }
                        _ = v25
                    }
                    v57 := GenS3{F0: 412, F1: 128}
                    _ = fi22
                    _ = v57
                }
                v58 := GenS9{F0: (v2 <= 2147483647), F1: 4294967295, F2: 4294967295}
                _ = v58
            }
            for fi59 := 0; fi59 < 20; fi59++ {
                switch fi59 {
                case 4294967295:
                    var v60 int16 = 0
                    _ = v60
                case 1:
                    var v61 uint16 = 411
                    _ = v61
                case -2147483648:
                    fmt.Println(v2, v2)
                default:
                    if true {
                        var v62 uint64 = 0
                        if (2 < 100) {
                            for fi63 := 0; fi63 < 8; fi63++ {
                                fmt.Println(v2, v2)
                                switch v1 {
                                case "hello":
                                    v20 = (v20 & (v20 / (((v20 % ((v20) | 1))) | 1)))
                                    v64 := v4[2147483647]
                                    switch fi59 {
                                    case -9223372036854775808:
                                        var v65 float32 = (0.1 / (1e10 / 0.1))
                                        _ = v65
                                    case 0:
                                        switch fi59 {
                                        case 9223372036854775807:
                                            switch ("foo bar" + (v1 + v1)) {
                                            case "a":
                                                var v66 float32 = 1.5
                                                _ = v66
                                            case "foo bar":
                                                var v67 float32 = 3.14159
                                                _ = v67
                                            }
                                        case 1:
                                            delete(v4, 0)
                                        case 42:
                                            if true {
                                                v64 += fi59
                                                switch v1 {
                                                case "a":
                                                    fmt.Println(fi59)
                                                case "hello":
                                                    v5 += (1.5 - float64((0.1 * (-2.25 * ((0.1 + 3.14159) / 1e10)))))
                                                case "foo bar":
                                                    switch -120 {
                                                    case 128:
                                                        v2 = v2
                                                    }
                                                default:
                                                    if false {
                                                        fmt.Println(v2, v1)
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if true {
                                                            v68 := GenS2{F0: (0 == v20), F1: v1, F2: v20}
                                                            if (v1 >= (v1 + "a")) {
                                                                if (42 != (v62 & (v62 << ((628) & 63)))) {
                                                                    if (215 != 652) {
                                                                        v69 := GenS9{F0: (147 != uint(v62)), F1: fi63, F2: 82}
                                                                        v3 = append(v3, (^(v64)))
                                                                        var v70 int64 = -137
                                                                        _ = v69
                                                                        _ = v70
                                                                    }
                                                                    v2 = -847
                                                                    var v71 float32 = (0.0 / -2.25)
                                                                    _ = v71
                                                                } else {
                                                                    v64 = 2147483647
                                                                }
                                                                fmt.Println(v1, v1, v2)
                                                            }
                                                            var v72 int16 = -787
                                                            _ = v68
                                                            _ = v72
                                                        }
                                                    } else {
                                                        v1 = ("" + "hello")
                                                    }
                                                }
                                                var v73 int8 = v20
                                                _ = v73
                                            }
                                        default:
                                            v1 = ""
                                        }
                                    }
                                    _ = v64
                                default:
                                    var v74 int64 = 9223372036854775807
                                    _ = v74
                                }
                                _ = fi63
                            }
                        } else {
                            if (2 >= 674) {
                                switch fi59 {
                                case 210:
                                    var v75 uint = 2
                                    _ = v75
                                default:
                                    v5 = v5
                                }
                                var v76 uint16 = 330
                                v1 = (("a" + "") + (v1 + v1))
                                _ = v76
                            } else {
                                var v77 uint8 = 1
                                v5 += v5
                                var v78 uint = 65535
                                _ = v77
                                _ = v78
                            }
                            v79 := GenS8{F0: v1}
                            _ = v79
                        }
                        _ = v62
                    } else {
                        for fi80 := 0; fi80 < 14; fi80++ {
                            v81 := GenS10{F0: 18446744073709551615}
                            var v82 int64 = -9223372036854775808
                            _ = fi80
                            _ = v81
                            _ = v82
                        }
                        if true {
                            for fi83 := 0; fi83 < 15; fi83++ {
                                if (true && (false || (788 > 65535))) {
                                    v1 = v1
                                }
                                _ = fi83
                            }
                        } else {
                            if true {
                                if true {
                                    var v84 uint32 = 2147483647
                                    _ = v84
                                }
                            } else {
                                v5 += ((float64((-2.25 - (float32(((v5 / v5) - 3.14159)) / (-2.25 / 1e10)))) - v5) * 0.0)
                            }
                        }
                    }
                }
                switch (fi59 / ((1) | 1)) {
                case 32768:
                    _ = len(v3)
                case -950:
                    _ = len(v4)
                case -464:
                    for fi85 := 0; fi85 < 3; fi85++ {
                        if (false || false) {
                            fmt.Println(fi85)
                            if true {
                                v20 = (v20 / ((v20) | 1))
                                _ = len(v4)
                                v86 := GenS11{F0: (v2 + v2), F1: 144, F2: fi85}
                                if false {
                                    v2 += -462
                                    var v87 uint = 52
                                    if true {
                                        v5 = (1e10 - (float64((((-2.25 - (float32(((v5 - 1e10) * (float64(1e10) + v5))) - (0.1 + 0.1))) * 1e10) / float32((3.14159 / v5)))) + 0.0))
                                        fmt.Println(v20, v2, v5)
                                        _ = len(v3)
                                        if ((3.14159 != 1.5) || ((-(v20)) == v20)) {
                                            if true {
                                                if ((^(v87)) < v87) {
                                                    switch (fi59 % (((^(fi59))) | 1)) {
                                                    case 127:
                                                        var v88 int8 = (v20 / ((42) | 1))
                                                        _ = v88
                                                    case 119:
                                                        v2 = (v2 | -422)
                                                    default:
                                                        v89 := GenS10{F0: 772}
                                                        _ = v89
                                                    }
                                                    var v90 float32 = float32(1.5)
                                                    delete(v4, 0)
                                                    _ = v90
                                                }
                                                var v91 float64 = (v5 / 0.1)
                                                if false {
                                                    fmt.Println(fi85, v20, v1)
                                                    switch fi85 {
                                                    case -23:
                                                        v1 = "foo bar"
                                                    }
                                                } else {
                                                    fmt.Println(fi85, fi59, v5)
                                                }
                                                _ = v91
                                            } else {
                                                var v92 uint8 = 255
                                                if true {
                                                    fmt.Println(v1, v20)
                                                    switch -1 {
                                                    case 162:
                                                        v5 = (v5 * (float64((0.1 + -2.25)) - v5))
                                                    case 32768:
                                                        v20 = (v20 & (-(v20)))
                                                    case 389:
                                                        var v93 float32 = 0.0
                                                        _ = v93
                                                    default:
                                                        delete(v4, 4294967295)
                                                    }
                                                    v94 := GenS6{F0: true}
                                                    _ = v94
                                                }
                                                _ = v92
                                            }
                                            v95 := v4[65535]
                                            _ = v95
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    _ = v87
                                } else {
                                    if false {
                                        v96 := GenS3{F0: 2, F1: 2}
                                        delete(v4, 42)
                                        _ = len(v3)
                                        var v97 string = v1
                                        _ = v96
                                        _ = v97
                                    } else {
                                        v98 := GenS7{F0: 1, F1: v20}
                                        if false {
                                            v20 += v20
                                            v99 := GenS5{F0: true, F1: (v2 & (-(v2))), F2: v2}
                                            if false {
                                                v100 := GenS8{F0: (v1 + v1)}
                                                v3 = append(v3, fi85)
                                                v20 -= -1
                                                var v101 float64 = v5
                                                if ((v20 > (^(v20))) || true) {
                                                    _ = len(v4)
                                                    v102 := GenS1{F0: false}
                                                    _ = v102
                                                } else {
                                                    if (1000 == 566) {
                                                        v1 = "a"
                                                        if true {
                                                            var v103 float32 = (-2.25 - ((float32(3.14159) / 1.5) / 1.5))
                                                            _ = len(v3)
                                                            var v104 int16 = 1000
                                                            _ = v103
                                                            _ = v104
                                                        }
                                                        v20 += (v20 / (((v20 >> ((65535) & 63))) | 1))
                                                    } else {
                                                        v3 = append(v3, (fi59 % ((642) | 1)))
                                                        v101 -= v101
                                                    }
                                                    switch v1 {
                                                    case "foo bar":
                                                        delete(v4, 0)
                                                    case "a":
                                                        v101 += float64(1.5)
                                                    case "hello":
                                                        switch 1 {
                                                        case 32768:
                                                            var v105 int16 = -1
                                                            _ = v105
                                                        case 893:
                                                            if false {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = fi85
                                                                }
                                                                v5 -= v5
                                                            }
                                                        }
                                                    default:
                                                        var v106 int64 = 117
                                                        _ = v106
                                                    }
                                                }
                                                _ = v100
                                                _ = v101
                                            } else {
                                                if ((v5 <= float64(float32((v5 * 1.5)))) || true) {
                                                    v5 = v5
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = (fi59 >> ((18446744073709551615) & 63))
                                                    }
                                                }
                                                v107 := GenS12{F0: (^(fi85)), F1: 0, F2: v20}
                                                v1 = "hello"
                                                _ = v107
                                            }
                                            _ = v99
                                        }
                                        _ = len(v3)
                                        _ = v98
                                    }
                                }
                                _ = v86
                            } else {
                                if (!(2 < 162)) {
                                    if true {
                                        _ = len(v3)
                                        fmt.Println(v1, v5, v2)
                                        var v108 uint16 = 32767
                                        _ = v108
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if (v20 == (v20 - v20)) {
                                            if (-32768 != -1) {
                                                var v109 uint64 = 396
                                                _ = v109
                                            } else {
                                                v110 := GenS6{F0: true}
                                                fmt.Println(v5, v2)
                                                _ = v110
                                            }
                                            v1 = v1
                                            if (0 != 18446744073709551615) {
                                                v1 = "a"
                                                v1 = "\n\t"
                                                v1 = (v1 + (v1 + (v1 + v1)))
                                            }
                                        } else {
                                            var v111 float64 = v5
                                            switch v1 {
                                            case "a":
                                                v3 = append(v3, (fi85 + fi59))
                                            case "foo bar":
                                                v20 = (v20 | v20)
                                            }
                                            var v112 bool = true
                                            switch 9223372036854775807 {
                                            case 2147483647:
                                                var v113 float64 = ((v5 - float64((float32(v5) * 1e10))) / v111)
                                                _ = v113
                                            default:
                                                var v114 float32 = 0.1
                                                _ = v114
                                            }
                                            _ = v111
                                            _ = v112
                                        }
                                        var v115 uint8 = 128
                                        _ = v115
                                    }
                                    fmt.Println(v5, v20, fi59)
                                    if len(v3) > 0 {
                                        v3[0] = 9223372036854775807
                                    }
                                    v1 = "a"
                                    v116 := v4[1]
                                    _ = v116
                                }
                                fmt.Println(v2)
                                var v117 uint16 = 0
                                if true {
                                    var v118 uint32 = 442
                                    fmt.Println(fi59, fi85, v5)
                                    _ = v118
                                } else {
                                    v1 = (v1 + v1)
                                    var v119 uint64 = 124
                                    _ = v119
                                }
                                if false {
                                    var v120 bool = (v20 != v20)
                                    if (2 <= 26) {
                                        v117 = v117
                                        v5 = 3.14159
                                    } else {
                                        v117 += (v117 >> ((327) & 63))
                                        v20 += v20
                                        v121 := GenS10{F0: 2}
                                        _ = v121
                                    }
                                    _ = v120
                                }
                                _ = v117
                            }
                            switch v1 {
                            case "hello":
                                switch -207 {
                                case 64:
                                    var v122 uint = 684
                                    _ = v122
                                }
                            }
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(fi59, fi85, v20)
                        _ = fi85
                    }
                }
                _ = fi59
            }
        } else {
            v123 := GenS3{F0: 797, F1: 34}
            _ = v123
        }
    }
    var v124 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v124
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
