package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 uint32; F1 int32; F2 int64 }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8; F1 uint; F2 int16 }
type GenS9 struct { F0 int64; F1 int32; F2 int32 }
type GenS10 struct { F0 string; F1 int8; F2 string }
type GenS11 struct { F0 int; F1 uint8 }
type GenS12 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS13 struct { F0 bool; F1 uint16 }
type GenS14 struct { F0 uint8; F1 bool }
func genFn0(p0 string, p1 int32, p2 uint64) int16 {
    var v7 float32 = (-2.25 - 0.1)
    v8 := GenS1{F0: -190, F1: 42, F2: true}
    if true {
        var v9 int8 = 13
        var v10 float64 = 1.5
        _ = v10
        _ = v9
    } else {
        v11 := GenS1{F0: 32768, F1: 2147483647, F2: false}
        p2 = p2
        _ = v11
    }
    p1 = p1
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := GenS1{F0: fi12, F1: 0, F2: (-1 == 42)}
        fmt.Println(v7)
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return int16(p1)
}
func genFn1(p0 uint8, p1 uint8, p2 uint) int16 {
    p0 = (p1 << (((p2 << ((p2) & 63))) & 63))
    if false {
        for fi14 := 0; fi14 < 19; fi14++ {
            v15 := GenS2{F0: 1}
            var v16 float32 = -2.25
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    var v17 uint32 = 4294967295
    switch "hello" {
    case "\n\t":
        var v18 uint8 = (p1 ^ (p0 ^ (p0 & (^(p1)))))
        _ = v18
    case "foo bar":
        var v19 int8 = -97
        _ = v19
    default:
        v20 := []int{-891, 0, 42}
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return -174
}

func main() {
    var v1 uint32 = 1
    v2 := []uint32{v1, 1, v1}
    var v3 string = "a"
    v4 := []int{2147483647, 945, -9223372036854775808, -318}
    var v5 uint64 = 4294967295
    var v6 int = 32767
    v5 += 417
    v4 = append(v4, 9223372036854775807)
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v21 int = v6
        v6 -= v6
        _ = v21
    } else {
        for fi22 := 0; fi22 < 13; fi22++ {
            if ((!false) || (829 > 255)) {
                switch (v6 * (fi22 + v6)) {
                case 9223372036854775807:
                    for fi23 := 0; fi23 < 8; fi23++ {
                        if true {
                            if false {
                                var v24 uint = 255
                                fmt.Println(fi23, v1)
                                _ = v24
                            } else {
                                if true {
                                    v3 = (("a" + v3) + v3)
                                } else {
                                    v6 = -587
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            switch "hello" {
                            case "hello":
                                fmt.Println(v6, v1)
                                if false {
                                    fmt.Println(fi23, v1, fi22)
                                    var v25 uint16 = 1
                                    _ = v25
                                } else {
                                    var v26 int = fi22
                                    if true {
                                        var v27 uint8 = 1
                                        _ = len(v4)
                                        _ = v27
                                    } else {
                                        var v28 string = v3
                                        _ = len(v4)
                                        var v29 bool = (1000 != 176)
                                        _ = v28
                                        _ = v29
                                    }
                                    _ = v26
                                }
                                v5 = (v5 ^ (v5 * v5))
                            case "foo bar":
                                v6 -= fi23
                            case "a":
                                if ((false || (!false)) && ((1e10 < 0.0) || true)) {
                                    var v30 uint16 = 1
                                    var v31 uint64 = 42
                                    _ = v30
                                    _ = v31
                                }
                            default:
                                var v32 int16 = 1
                                _ = v32
                            }
                            if (0 > 2) {
                                v4 = append(v4, int(v1))
                            } else {
                                var v33 int = (fi23 | (fi23 >> ((2) & 63)))
                                if (1 < 1) {
                                    if true {
                                        var v34 float32 = float32(0.0)
                                        var v35 int16 = 1000
                                        fmt.Println(v33, fi22)
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v36 := GenS3{F0: uint64(v1), F1: 42, F2: 373}
                                        fmt.Println(v5, fi22, fi22)
                                        var v37 uint64 = v5
                                        _ = v36
                                        _ = v37
                                    }
                                } else {
                                    switch fi22 {
                                    case 9223372036854775807:
                                        v4 = append(v4, 32767)
                                    case -2147483648:
                                        if false {
                                            fmt.Println(v6)
                                        }
                                    default:
                                        v38 := GenS4{F0: 255, F1: true}
                                        _ = v38
                                    }
                                }
                                _ = v33
                            }
                        } else {
                            var v39 uint = 526
                            _ = v39
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi23
                    }
                case -9223372036854775808:
                    for fi40 := 0; fi40 < 14; fi40++ {
                        if true {
                            var v41 int8 = -15
                            v4 = append(v4, (v6 - 0))
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                            _ = v41
                        } else {
                            _ = len(v2)
                            v4 = append(v4, (v6 - (fi22 + (^(v6)))))
                            var v42 int = 9223372036854775807
                            _ = v42
                        }
                        _ = fi40
                    }
                case 32767:
                    for fi43 := 0; fi43 < 1; fi43++ {
                        var v44 int32 = 0
                        var v45 int32 = 42
                        var v46 string = (("a" + v3) + ((v3 + (v3 + "a")) + v3))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi43
                        _ = v44
                        _ = v45
                        _ = v46
                    }
                }
                var v47 int = (^(fi22))
                _ = v47
            } else {
                v48 := GenS2{F0: 65535}
                v1 += (v1 | (v1 | (^(v1))))
                _ = v48
            }
            _ = fi22
        }
        var v49 float32 = 0.0
        switch v6 {
        case 233:
            var v50 int64 = 255
            _ = v50
        }
        _ = v49
    }
    var v51 int64 = 0
    for fi52 := 0; fi52 < 19; fi52++ {
        if true {
            if false {
                for fi53 := 0; fi53 < 8; fi53++ {
                    if (v6 > fi52) {
                        v54 := GenS5{F0: (v1 ^ (v1 & 46)), F1: -1, F2: v51}
                        if false {
                            v55 := GenS5{F0: 2147483647, F1: 2147483647, F2: (v51 & (v51 % ((v51) | 1)))}
                            var v56 string = (v3 + "hello")
                            _ = v55
                            _ = v56
                        } else {
                            fmt.Println(v3, fi52, v51)
                        }
                        var v57 uint16 = 65535
                        _ = v54
                        _ = v57
                    } else {
                        switch v3 {
                        case "hello":
                            fmt.Println(fi53, v1, v1)
                        case "\n\t":
                            v58 := GenS6{F0: 1, F1: (v1 & (v1 | (v1 * (^(v1)))))}
                            _ = v58
                        default:
                            v6 += (fi53 >> ((255) & 63))
                        }
                    }
                    var v59 float32 = 0.1
                    _ = fi53
                    _ = v59
                }
            }
        }
        v60 := GenS6{F0: v5, F1: (^(v1))}
        for fi61 := 0; fi61 < 10; fi61++ {
            var v62 uint32 = (v1 | (v1 | v1))
            v63 := GenS5{F0: (^(v1)), F1: 1, F2: (v51 * int64(fi52))}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi61
            _ = v62
            _ = v63
        }
        _ = fi52
        _ = v60
    }
    var v64 uint8 = 1
    v5 = v5
    fmt.Println(v51)
    var v65 string = ("a" + "a")
    if (!false) {
        if false {
            v51 += (v51 * v51)
            if ("a" == "hello") {
                v66 := GenS5{F0: 2, F1: 42, F2: (v51 % (((v51 % ((-216) | 1))) | 1))}
                var v67 int32 = -965
                _ = v66
                _ = v67
            }
        }
        for fi68 := 0; fi68 < 11; fi68++ {
            if false {
                v69 := GenS7{F0: (true || false)}
                _ = v69
            } else {
                var v70 int = -853
                _ = v70
            }
            if false {
                v71 := GenS8{F0: -36, F1: 301, F2: 594}
                _ = v71
            } else {
                if ((v64 | v64) >= (v64 + v64)) {
                    if true {
                        if false {
                            for fi72 := 0; fi72 < 1; fi72++ {
                                if true {
                                    v73 := GenS9{F0: v51, F1: -2147483648, F2: 0}
                                    _ = v73
                                }
                                _ = len(v2)
                                switch fi68 {
                                case 4294967295:
                                    if true {
                                        if true {
                                            _ = len(v2)
                                        }
                                    }
                                case -9223372036854775808:
                                    switch v3 {
                                    case "foo bar":
                                        if len(v2) > 0 {
                                            v2[0] = (v1 % ((v1) | 1))
                                        }
                                    }
                                }
                                var v74 int16 = 45
                                _ = fi72
                                _ = v74
                            }
                        }
                        for fi75 := 0; fi75 < 19; fi75++ {
                            if len(v2) > 0 {
                                v2[0] = (^(v1))
                            }
                            _ = fi75
                        }
                    } else {
                        for fi76 := 0; fi76 < 11; fi76++ {
                            var v77 uint8 = (^(v64))
                            switch v65 {
                            case "a":
                                switch "hello" {
                                case "foo bar":
                                    v78 := GenS4{F0: uint(v64), F1: false}
                                    _ = v78
                                case "a":
                                    var v79 int16 = genFn1(113, (^(v64)), 65535)
                                    _ = v79
                                }
                            case "":
                                if true {
                                    _ = len(v4)
                                    var v80 string = ""
                                    if true {
                                        switch v3 {
                                        case "foo bar":
                                            if true {
                                                var v81 string = v80
                                                _ = v81
                                            } else {
                                                v51 -= v51
                                                if true {
                                                    v82 := GenS2{F0: uint16(v64)}
                                                    _ = v82
                                                } else {
                                                    var v83 uint8 = 0
                                                    if (413 != -32768) {
                                                        if false {
                                                            switch (((v3 + "") + v80) + (v80 + "\n\t")) {
                                                            case "a":
                                                                v80 = (v3 + v80)
                                                            case "hello":
                                                                var v84 uint32 = v1
                                                                _ = v84
                                                            case "\n\t":
                                                                var v85 int64 = (^(v51))
                                                                _ = v85
                                                            default:
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            }
                                                        }
                                                        var v86 int16 = 934
                                                        _ = len(v2)
                                                        fmt.Println(fi68)
                                                        switch ((v65 + v80) + v3) {
                                                        case "\n\t":
                                                            v51 = v51
                                                        case "":
                                                            v51 += -735
                                                        case "foo bar":
                                                            if true {
                                                                if true {
                                                                    v64 += (v83 * v83)
                                                                } else {
                                                                    v87 := GenS7{F0: true}
                                                                    _ = v87
                                                                }
                                                                v2 = append(v2, (v1 - v1))
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                v3 = v80
                                                            } else {
                                                                if (-27 > int8(v83)) {
                                                                    switch fi68 {
                                                                    case 32767:
                                                                        var v88 int16 = (v86 ^ (v86 + (v86 / ((-333) | 1))))
                                                                        _ = v88
                                                                    case 9223372036854775807:
                                                                        if ((float32(1.5) + (float32(3.14159) / float32(1.5))) != 0.0) {
                                                                            var v89 int = fi68
                                                                            v51 = v51
                                                                            _ = v89
                                                                        }
                                                                    case 80:
                                                                        fmt.Println(v64, v83, v64)
                                                                        var v90 int8 = -1
                                                                        _ = v90
                                                                    }
                                                                    var v91 uint64 = (v5 | v5)
                                                                    _ = v91
                                                                }
                                                                if false {
                                                                    v2 = append(v2, v1)
                                                                }
                                                            }
                                                        }
                                                        _ = v86
                                                    }
                                                    _ = v83
                                                }
                                            }
                                        case "a":
                                            var v92 float64 = 0.1
                                            switch 334 {
                                            case 2147483647:
                                                fmt.Println(v92, v6, v1)
                                            case -297:
                                                v93 := GenS6{F0: (v5 / (((v5 + (v5 ^ 42))) | 1)), F1: 78}
                                                _ = v93
                                            }
                                            switch v65 {
                                            case "":
                                                v65 = "hello"
                                            case "foo bar":
                                                v77 = (v77 & v64)
                                            default:
                                                var v94 int8 = -1
                                                _ = v94
                                            }
                                            _ = v92
                                        }
                                    } else {
                                        v95 := GenS10{F0: v3, F1: 127, F2: (v65 + v3)}
                                        var v96 float64 = (((1.5 + 0.0) * float64(float32(0.0))) / 1e10)
                                        fmt.Println(v6)
                                        _ = v95
                                        _ = v96
                                    }
                                    _ = v80
                                } else {
                                    if (-1 == -1) {
                                        var v97 int32 = 42
                                        _ = len(v4)
                                        if (true || false) {
                                            var v98 int16 = int16(v1)
                                            switch "hello" {
                                            case "foo bar":
                                                v99 := GenS11{F0: (-(fi68)), F1: v77}
                                                _ = v99
                                            case "a":
                                                if (1 < -517) {
                                                    var v100 int32 = (^(v97))
                                                    _ = v100
                                                }
                                            }
                                            var v101 uint = 298
                                            var v102 float32 = (float32(float64(float32(-2.25))) + float32((3.14159 * float64(((3.14159 * 3.14159) * 0.1)))))
                                            _ = v101
                                            _ = v102
                                            _ = v98
                                        }
                                        _ = v97
                                    }
                                    var v103 uint16 = 0
                                    _ = v103
                                }
                            case "foo bar":
                                v51 -= 9223372036854775807
                            default:
                                v104 := GenS7{F0: true}
                                if true {
                                    fmt.Println(v1, v51)
                                    var v105 int = v6
                                    var v106 uint = 733
                                    _ = v105
                                    _ = v106
                                } else {
                                    v107 := GenS8{F0: 0, F1: 113, F2: 1}
                                    if len(v4) > 0 {
                                        v4[0] = 42
                                    }
                                    _ = v107
                                }
                                _ = v104
                            }
                            _ = fi76
                            _ = v77
                        }
                        v51 -= 622
                    }
                    v3 = v65
                }
            }
            var v108 uint64 = (^(v5))
            _ = fi68
            _ = v108
        }
    } else {
        switch 32768 {
        case 42:
            if (!false) {
                var v109 int = v6
                _ = v109
            }
        case 128:
            switch (v6 << ((99) & 63)) {
            case 128:
                switch (v6 | (v6 << ((42) & 63))) {
                case 127:
                    v5 += (v5 + 168)
                case 42:
                    v110 := GenS8{F0: 127, F1: 0, F2: -1}
                    _ = v110
                default:
                    for fi111 := 0; fi111 < 18; fi111++ {
                        if ((-32768 == 1000) && (255 == 2)) {
                            switch "foo bar" {
                            case "\n\t":
                                v6 = v6
                            default:
                                var v112 float32 = 1.5
                                _ = v112
                            }
                            var v113 float64 = 1.5
                            var v114 uint32 = (v1 | 0)
                            _ = v113
                            _ = v114
                        } else {
                            v51 = -1
                            var v115 bool = true
                            _ = v115
                        }
                        if false {
                            v65 = (v3 + v65)
                            for fi116 := 0; fi116 < 3; fi116++ {
                                v3 = v3
                                switch 568 {
                                case 128:
                                    v117 := GenS4{F0: 255, F1: ((("hello" > v65) || true) && (134 > 255))}
                                    _ = v117
                                default:
                                    var v118 string = ""
                                    _ = v118
                                }
                                _ = fi116
                            }
                        } else {
                            if false {
                                v51 = 890
                                for fi119 := 0; fi119 < 13; fi119++ {
                                    v65 = (("\n\t" + v65) + ("" + v3))
                                    var v120 bool = (!false)
                                    _ = fi119
                                    _ = v120
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            } else {
                                _ = len(v2)
                                _ = len(v4)
                            }
                            for fi121 := 0; fi121 < 16; fi121++ {
                                v1 = (v1 + (v1 << ((739) & 63)))
                                if (806 < -1) {
                                    fmt.Println(fi121, v64)
                                    var v122 int = (^(v6))
                                    v123 := GenS12{F0: genFn0(((v3 + (v65 + ("\n\t" + v3))) + v3), 0, v5), F1: 105, F2: 65535}
                                    _ = v122
                                    _ = v123
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v6 -= int(v64)
                                }
                                _ = fi121
                            }
                        }
                        _ = fi111
                    }
                }
            case 32767:
                v65 = "\n\t"
            }
        default:
            var v124 bool = (!false)
            _ = v124
        }
        if (0 <= -87) {
            fmt.Println(v5)
        }
        for fi125 := 0; fi125 < 18; fi125++ {
            var v126 uint64 = v5
            _ = fi125
            _ = v126
        }
    }
    for fi127 := 0; fi127 < 5; fi127++ {
        var v128 int16 = 42
        switch "" {
        case "":
            fmt.Println(v6, v64)
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v129 := GenS13{F0: true, F1: 255}
        _ = fi127
        _ = v128
        _ = v129
    }
    v130 := GenS14{F0: (^(v64)), F1: true}
    var v131 float32 = (0.0 + 0.0)
    for fi132 := 0; fi132 < 12; fi132++ {
        var v133 string = v3
        _ = fi132
        _ = v133
    }
    if len(v4) > 0 {
        v4[0] = v6
    }
    v3 = "a"
    for fi134 := 0; fi134 < 8; fi134++ {
        var v135 int32 = 777
        var v136 int16 = 446
        _ = fi134
        _ = v135
        _ = v136
    }
    fmt.Println("v1:", v1)
    fmt.Println("v131:", v131)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v130
    _ = v131
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v64
    _ = v65
}
