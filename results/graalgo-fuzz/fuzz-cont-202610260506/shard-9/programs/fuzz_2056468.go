package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int }
type GenS3 struct { F0 int8; F1 bool; F2 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 int64; F1 int64 }
type GenS8 struct { F0 uint64; F1 int32 }
type GenS9 struct { F0 uint8; F1 int }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int8; F1 string }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint16; F1 uint32 }
type GenS16 struct { F0 uint32; F1 int64; F2 int32 }
type GenS17 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 string }
type GenS20 struct { F0 uint8; F1 uint; F2 int }
func genFn0(p0 uint64, p1 int64) uint16 {
    p0 = 2
    v7 := GenS1{F0: 1, F1: 1000}
    v8 := GenS1{F0: -1, F1: -637}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 uint8 = 244
    var v2 int = -9223372036854775808
    var v3 int16 = -407
    v4 := make(map[uint32]uint32)
    var v5 uint32 = uint32(v1)
    v6 := make(map[int]int)
    for fi9 := 0; fi9 < 14; fi9++ {
        v5 = (v5 * 0)
        _ = fi9
    }
    var v10 int64 = int64(v1)
    if (1.5 >= 0.0) {
        v11 := GenS2{F0: (v1 * 1), F1: 1, F2: v2}
        switch int(v10) {
        case 792:
            var v12 uint64 = 401
            _ = v12
        default:
            switch 283 {
            case 1:
                var v13 float32 = 1.5
                _ = v13
            }
        }
        _ = v11
    }
    switch 4294967295 {
    case 981:
        v6[v2] = (v2 >> ((0) & 63))
    case 9223372036854775807:
        var v14 int = -222
        _ = v14
    case 1:
        fmt.Println(v1)
    default:
        var v15 int32 = -753
        _ = v15
    }
    v16 := GenS2{F0: v1, F1: v5, F2: v2}
    v1 += 153
    v1 -= 110
    v3 -= 0
    for fi17 := 0; fi17 < 6; fi17++ {
        switch "a" {
        case "hello":
            if false {
                v18 := []int{(fi17 % ((fi17) | 1)), 957, (v2 * v2)}
                v19 := GenS2{F0: 128, F1: 42, F2: (v2 & 0)}
                _ = v18
                _ = v19
            }
        case "":
            var v20 int8 = 1
            _ = v20
        default:
            var v21 uint8 = uint8(v5)
            _ = v21
        }
        _ = len(v6)
        _ = fi17
    }
    var v22 int64 = 9223372036854775807
    if true {
        for fi23 := 0; fi23 < 5; fi23++ {
            switch 127 {
            case -451:
                if false {
                    if (("foo bar" + "foo bar") >= "") {
                        var v24 bool = (!((v2 * (fi23 & (fi23 | -868))) >= 127))
                        for fi25 := 0; fi25 < 2; fi25++ {
                            v10 = (v22 - v10)
                            _ = fi25
                        }
                        v5 -= 1
                        _ = v24
                    } else {
                        v26 := GenS3{F0: 0, F1: ((false && false) && (false && false)), F2: v5}
                        var v27 uint16 = 950
                        _ = v26
                        _ = v27
                    }
                    delete(v6, (fi23 / ((v2) | 1)))
                    var v28 bool = false
                    _ = v28
                } else {
                    switch v2 {
                    case 128:
                        var v29 bool = ((0.1 + 3.14159) >= 3.14159)
                        _ = v29
                    case -928:
                        v30 := GenS4{F0: true}
                        _ = v30
                    }
                    var v31 uint8 = v1
                    _ = v31
                }
            default:
                if false {
                    v5 -= v5
                } else {
                    v32 := GenS5{F0: (^(v1)), F1: v3}
                    _ = v32
                }
            }
            for fi33 := 0; fi33 < 17; fi33++ {
                var v34 float32 = 0.1
                _ = fi33
                _ = v34
            }
            for fi35 := 0; fi35 < 19; fi35++ {
                v36 := v6[(^(fi35))]
                _ = fi35
                _ = v36
            }
            _ = fi23
        }
    }
    if (v10 == v10) {
        if true {
            var v37 uint = 65535
            for fi38 := 0; fi38 < 6; fi38++ {
                if (42 == (^(v3))) {
                    var v39 string = ""
                    if true {
                        if ((255 < (v1 / (((v1 + 128)) | 1))) && (729 == 42)) {
                            var v40 float32 = 1e10
                            _ = v40
                        }
                    } else {
                        fmt.Println(v2, v3)
                        var v41 uint32 = 2
                        _ = v41
                    }
                    _ = v39
                } else {
                    v42 := GenS6{F0: (v2 / (((fi38 / (((v2 >> ((533) & 63))) | 1))) | 1)), F1: v10}
                    v3 = (v3 + v3)
                    if true {
                        for fi43 := 0; fi43 < 19; fi43++ {
                            var v44 float64 = float64((0.0 + (0.1 + 3.14159)))
                            fmt.Println(fi38, v1)
                            _ = fi43
                            _ = v44
                        }
                        for fi45 := 0; fi45 < 2; fi45++ {
                            v3 -= v3
                            var v46 int = 2147483647
                            fmt.Println(v46)
                            _ = fi45
                            _ = v46
                        }
                    } else {
                        var v47 int8 = 42
                        v48 := GenS4{F0: (false && (770 == 1))}
                        v49 := GenS1{F0: (^(v47)), F1: 32767}
                        _ = v47
                        _ = v48
                        _ = v49
                    }
                    _ = v42
                }
                _ = fi38
            }
            v22 = v22
            _ = v37
        } else {
            v5 += (v5 % ((v5) | 1))
            var v50 float64 = 3.14159
            var v51 uint64 = 4294967295
            _ = v50
            _ = v51
        }
        if ((v22 <= (v10 >> ((uint(v1)) & 63))) || true) {
            if true {
                v4[133] = (^(v5))
                if true {
                    var v52 uint8 = (v1 + 231)
                    _ = v52
                }
            } else {
                v3 += 32767
                if true {
                    v53 := v6[(v2 % ((42) | 1))]
                    v1 = (v1 & v1)
                    v54 := GenS3{F0: 100, F1: (int32(v1) >= -2147483648), F2: v5}
                    _ = v53
                    _ = v54
                } else {
                    switch ("hello" + "hello") {
                    case "foo bar":
                        switch v2 {
                        case -190:
                            var v55 string = "hello"
                            _ = v55
                        case 944:
                            fmt.Println(v3)
                        case -751:
                            v6[(^(v2))] = 4294967295
                        default:
                            var v56 float32 = ((((3.14159 * 1.5) / (float32(float64(0.0)) - 3.14159)) * -2.25) * -2.25)
                            _ = v56
                        }
                    case "\n\t":
                        for fi57 := 0; fi57 < 15; fi57++ {
                            if false {
                                v58 := GenS3{F0: 0, F1: true, F2: (v5 << ((255) & 63))}
                                _ = v58
                            } else {
                                v3 = 32767
                                v1 = v1
                            }
                            _ = len(v6)
                            v3 += v3
                            _ = fi57
                        }
                    default:
                        v59 := v4[0]
                        _ = v59
                    }
                    var v60 int = v2
                    _ = v60
                }
            }
            v22 = (v22 + (v10 % ((v10) | 1)))
            if (1e10 < 1e10) {
                if ((v3 | v3) >= (v3 - v3)) {
                    v10 -= v10
                    if true {
                        for fi61 := 0; fi61 < 11; fi61++ {
                            var v62 int = fi61
                            v1 = uint8(v5)
                            v63 := GenS7{F0: -514, F1: 503}
                            _ = fi61
                            _ = v62
                            _ = v63
                        }
                    }
                    v1 += v1
                } else {
                    var v64 int8 = -1
                    _ = v64
                }
                v22 += (^(v10))
                if true {
                    if ((int8(v2) >= 97) || (false || ((^(v1)) < (v1 << ((2) & 63))))) {
                        v22 = v22
                        if true {
                            for fi65 := 0; fi65 < 2; fi65++ {
                                v66 := []int{(fi65 / (((fi65 + 32767)) | 1))}
                                var v67 uint64 = 555
                                if len(v66) > 0 {
                                    v66[0] = (fi65 << ((914) & 63))
                                }
                                _ = fi65
                                _ = v66
                                _ = v67
                            }
                        } else {
                            v22 -= int64(v1)
                            switch (v2 / (((v2 << ((65535) & 63))) | 1)) {
                            case 128:
                                var v68 bool = false
                                _ = v68
                            default:
                                v10 = v10
                            }
                        }
                    }
                }
            }
        }
    } else {
        switch ("\n\t" + ("foo bar" + ("\n\t" + ("foo bar" + "")))) {
        case "foo bar":
            if true {
                for fi69 := 0; fi69 < 1; fi69++ {
                    fmt.Println(fi69, v5)
                    if (!(!(("a" > ("a" + "a")) && (482 != genFn0(1, (v10 % (((-(v22))) | 1))))))) {
                        v1 -= v1
                    } else {
                        v70 := []int{2147483647, fi69}
                        _ = v70
                    }
                    _ = fi69
                }
            } else {
                var v71 int32 = int32(v22)
                v72 := GenS8{F0: uint64(v5), F1: (v71 - v71)}
                for fi73 := 0; fi73 < 5; fi73++ {
                    delete(v4, 0)
                    v10 = v10
                    switch "a" {
                    case "a":
                        if (243 >= v1) {
                            for fi74 := 0; fi74 < 7; fi74++ {
                                v75 := GenS9{F0: v1, F1: 32768}
                                switch 997 {
                                case 887:
                                    fmt.Println(v71, v10, fi74)
                                    var v76 int8 = 0
                                    _ = v76
                                case 32767:
                                    v71 = v71
                                case -9223372036854775808:
                                    v22 = v10
                                default:
                                    fmt.Println(v71, v2, fi74)
                                }
                                _ = fi74
                                _ = v75
                            }
                            if true {
                                var v77 uint8 = v1
                                var v78 int64 = (v22 % ((9223372036854775807) | 1))
                                _ = v77
                                _ = v78
                            } else {
                                v22 -= (v22 >> ((42) & 63))
                                switch fi73 {
                                case 32768:
                                    if ((1.5 - 0.1) != float64((1.5 - (1e10 / 1.0)))) {
                                        v79 := GenS10{F0: (v10 ^ 258)}
                                        v22 -= (-(v10))
                                        _ = v79
                                    } else {
                                        if true {
                                            if false {
                                                if false {
                                                    v80 := []int{-2147483648, (-(v2))}
                                                    _ = v80
                                                } else {
                                                    var v81 int = -2147483648
                                                    _ = v81
                                                }
                                            }
                                        }
                                        if false {
                                            for fi82 := 0; fi82 < 18; fi82++ {
                                                if (false && true) {
                                                    fmt.Println(v22)
                                                    if false {
                                                        v83 := GenS5{F0: 2, F1: (v3 >> ((216) & 63))}
                                                        _ = v83
                                                    } else {
                                                        var v84 float64 = 1.5
                                                        fmt.Println(v5)
                                                        var v85 int32 = (v71 - -1)
                                                        fmt.Println(v5)
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                    if (18446744073709551615 > 42) {
                                                        if true {
                                                            switch "" {
                                                            case "foo bar":
                                                                if (255 != 436) {
                                                                    v22 -= v22
                                                                    var v86 int16 = (v3 << ((540) & 63))
                                                                    fmt.Println(fi73)
                                                                    var v87 int = 701
                                                                    fmt.Println(v10, v87)
                                                                    _ = v86
                                                                    _ = v87
                                                                } else {
                                                                    var v88 float32 = 1.5
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        v89 := GenS11{F0: fi82}
                                                                        _ = v89
                                                                    }
                                                                    _ = v88
                                                                }
                                                            case "":
                                                                var v90 uint16 = 42
                                                                _ = v90
                                                            case "a":
                                                                if true {
                                                                    var v91 bool = (0.0 != 3.14159)
                                                                    v10 = v10
                                                                    _ = v91
                                                                }
                                                            default:
                                                                fmt.Println(v5)
                                                            }
                                                            v71 = (v71 ^ 42)
                                                            if false {
                                                                v1 += v1
                                                            } else {
                                                                if true {
                                                                    var v92 int32 = v71
                                                                    v3 = v3
                                                                    if (!(0 >= (v3 & (v3 * v3)))) {
                                                                        fmt.Println(v71)
                                                                    }
                                                                    if true {
                                                                        var v93 uint64 = 0
                                                                        var v94 uint64 = v93
                                                                        v95 := GenS12{F0: "foo bar"}
                                                                        _ = v93
                                                                        _ = v94
                                                                        _ = v95
                                                                    } else {
                                                                        if true {
                                                                            v96 := GenS13{F0: 1, F1: "hello"}
                                                                            v22 -= v10
                                                                            _ = v96
                                                                        } else {
                                                                            switch v2 {
                                                                            case -604:
                                                                                if false {
                                                                                    var v97 int16 = 32767
                                                                                    _ = v97
                                                                                }
                                                                            case 169:
                                                                                if (((v5 >> ((2) & 63)) > 0) && true) {
                                                                                    fmt.Println(v1)
                                                                                    if (!(1e10 != 1e10)) {
                                                                                        var v98 uint32 = v5
                                                                                        _ = v98
                                                                                    } else {
                                                                                        v99 := GenS4{F0: false}
                                                                                        v100 := GenS14{F0: v5}
                                                                                        fmt.Println(v5, v5)
                                                                                        if false {
                                                                                            v101 := GenS15{F0: 42, F1: v5}
                                                                                            if (9223372036854775807 == v2) {
                                                                                                v5 = (v5 - 471)
                                                                                                var v102 uint = 255
                                                                                                fmt.Println(v22)
                                                                                                _ = v102
                                                                                            } else {
                                                                                                v3 = (v3 >> ((0) & 63))
                                                                                                switch "foo bar" {
                                                                                                case "":
                                                                                                    v103 := GenS16{F0: (v5 * v5), F1: v10, F2: v71}
                                                                                                    _ = v103
                                                                                                case "\n\t":
                                                                                                    v22 = (v10 ^ v10)
                                                                                                case "foo bar":
                                                                                                    v104 := GenS17{F0: 2, F1: 1, F2: v92}
                                                                                                    _ = v104
                                                                                                }
                                                                                                var v105 int16 = v3
                                                                                                _ = v105
                                                                                            }
                                                                                            v106 := []int{(fi82 ^ 128), fi73, (fi82 | fi82), (fi82 % (((v2 << ((4294967295) & 63))) | 1))}
                                                                                            _ = v101
                                                                                            _ = v106
                                                                                        } else {
                                                                                            fmt.Println(fi73)
                                                                                            fmt.Println(fi82)
                                                                                        }
                                                                                        _ = v100
                                                                                        _ = v99
                                                                                    }
                                                                                    v107 := GenS11{F0: 32768}
                                                                                    _ = v107
                                                                                } else {
                                                                                    fmt.Println(v5, v71, v71)
                                                                                }
                                                                            case 9223372036854775807:
                                                                                v2 -= fi82
                                                                            default:
                                                                                if (true && true) {
                                                                                    _ = len(v4)
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    delete(v4, 563)
                                                                    _ = v92
                                                                }
                                                            }
                                                        }
                                                        if true {
                                                            v5 = (v5 * 2147483648)
                                                        }
                                                    } else {
                                                        if (((v71 % ((-1) | 1)) != (v71 - (v71 & v71))) && ((!((65535 >= 255) || (v1 == (v1 >> ((1) & 63))))) && (((("" + "") + ("" + (("" + ("" + "\n\t")) + "hello"))) + "foo bar") != "foo bar"))) {
                                                            fmt.Println(v10)
                                                            var v108 float32 = ((0.0 * ((1.5 + (1e10 + float32(-2.25))) - 0.1)) / 1.0)
                                                            var v109 float64 = ((0.0 * ((0.0 * 0.1) * (1e10 / -2.25))) / -2.25)
                                                            _ = v108
                                                            _ = v109
                                                        }
                                                        var v110 bool = (125 <= 1)
                                                        if true {
                                                            v111 := v6[9223372036854775807]
                                                            _ = v111
                                                        }
                                                        _ = v110
                                                    }
                                                    fmt.Println(v10, v1)
                                                    v10 = int64(v3)
                                                } else {
                                                    var v112 int8 = 4
                                                    delete(v4, v5)
                                                    if false {
                                                        var v113 float32 = ((-2.25 - -2.25) + 0.1)
                                                        _ = v113
                                                    }
                                                    _ = v112
                                                }
                                                _ = fi82
                                            }
                                            v10 += (v10 - (v10 | v10))
                                        }
                                        v114 := v6[(v2 & 4294967295)]
                                        _ = v114
                                    }
                                case 930:
                                    v115 := GenS18{F0: 286}
                                    _ = v115
                                default:
                                    var v116 int = (fi73 << ((42) & 63))
                                    _ = v116
                                }
                                var v117 uint64 = 90
                                _ = v117
                            }
                        } else {
                            var v118 uint = 42
                            v71 = -1
                            v5 -= (v5 / (((v5 >> (((v118 * (v118 & v118))) & 63))) | 1))
                            _ = v118
                        }
                    case "foo bar":
                        var v119 uint = 2
                        _ = v119
                    }
                    _ = fi73
                }
                _ = v71
                _ = v72
            }
        case "":
            v120 := GenS19{F0: ("\n\t" + "a")}
            _ = v120
        case "\n\t":
            v10 += (v22 & v22)
        }
        if (!((v5 * 0) <= (v5 | 2147483647))) {
            var v121 uint64 = 1
            v5 = (v5 >> ((1) & 63))
            if true {
                v122 := GenS20{F0: v1, F1: 278, F2: 767}
                _ = v122
            }
            _ = v121
        } else {
            v123 := []int{747}
            fmt.Println(v5, v2, v5)
            if (false && false) {
                v3 = 42
            }
            _ = v123
        }
    }
    v5 -= (v5 - 42)
    var v124 bool = false
    if (false || v124) {
        _ = len(v4)
    } else {
        for fi125 := 0; fi125 < 16; fi125++ {
            if v124 {
                var v126 float64 = 1e10
                fmt.Println(v1, v22, v3)
                if v124 {
                    v124 = v124
                    var v127 int32 = 42
                    for fi128 := 0; fi128 < 9; fi128++ {
                        v127 = -2147483648
                        v3 -= 1
                        _ = fi128
                    }
                    _ = v127
                }
                _ = v126
            }
            switch (^(v2)) {
            case 4294967295:
                for fi129 := 0; fi129 < 3; fi129++ {
                    fmt.Println(fi125)
                    v130 := GenS13{F0: 84, F1: ("a" + ("\n\t" + ("\n\t" + ("" + "foo bar"))))}
                    v131 := GenS12{F0: "foo bar"}
                    _ = fi129
                    _ = v130
                    _ = v131
                }
            case 2147483647:
                switch 9223372036854775807 {
                case 127:
                    v5 -= (v5 % (((v5 ^ 2)) | 1))
                }
            default:
                v132 := GenS16{F0: uint32(v1), F1: v22, F2: 754}
                _ = v132
            }
            v133 := []int{}
            _ = fi125
            _ = v133
        }
        switch ("\n\t" + "") {
        case "\n\t":
            v10 = v10
        case "foo bar":
            v134 := GenS20{F0: 2, F1: 910, F2: (v2 ^ 9223372036854775807)}
            _ = v134
        }
        var v135 int8 = 100
        _ = v135
    }
    v1 += 0
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v124
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
