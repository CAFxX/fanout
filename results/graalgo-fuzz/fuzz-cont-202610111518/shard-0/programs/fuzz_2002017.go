package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint }
type GenS2 struct { F0 int32; F1 int32; F2 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 uint32; F1 uint64; F2 string }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int32; F1 bool; F2 int }
type GenS11 struct { F0 uint; F1 bool; F2 int32 }
type GenS12 struct { F0 string; F1 uint; F2 int32 }
type GenS13 struct { F0 uint64; F1 uint; F2 int64 }
type GenS14 struct { F0 string; F1 int16 }
type GenS15 struct { F0 int64; F1 uint8 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint8; F1 int }
type GenS18 struct { F0 string; F1 string }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 uint; F1 uint16 }
type GenS21 struct { F0 uint16 }
type GenS22 struct { F0 uint32; F1 int16; F2 int8 }
func genFn0(p0 uint8, p1 int16) string {
    v6 := []int{127, 0, 32768, 32767}
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint = 255
        if (!(1e10 > -2.25)) {
            _ = len(v6)
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 int32 = 0
                var v11 uint16 = 255
                p0 += 0
                _ = fi9
                _ = v10
                _ = v11
            }
            for fi12 := 0; fi12 < 13; fi12++ {
                if ((!((^(p0)) < 2)) && false) {
                    if true {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if len(v6) > 0 {
                            v6[0] = 632
                        }
                        v8 = (v8 * 65535)
                        v8 += (^(v8))
                        if false {
                            if true {
                                v13 := GenS1{F0: (fi7 * fi7), F1: 0, F2: (^(v8))}
                                var v14 uint8 = p0
                                _ = v13
                                _ = v14
                            }
                            p1 -= (p1 + 32767)
                        }
                        var v15 uint16 = 42
                        if (!(((0.0 - (((1.5 + 3.14159) / 1e10) - 3.14159)) / 1.5) < 0.0)) {
                            if ((("a" + ((("foo bar" + "foo bar") + "foo bar") + "foo bar")) + ("\n\t" + ("foo bar" + (("hello" + "") + "\n\t")))) == "foo bar") {
                                if false {
                                    switch "hello" {
                                    case "foo bar":
                                        if true {
                                            v16 := GenS2{F0: -731, F1: -2147483648, F2: (9223372036854775807 < 9223372036854775807)}
                                            fmt.Println(v15)
                                            if false {
                                                v8 = 672
                                                if (1.5 < (((1.5 + 0.1) / -2.25) / 1.5)) {
                                                    var v17 int16 = (p1 / ((p1) | 1))
                                                    var v18 float64 = (0.0 / float64((0.0 + 1.5)))
                                                    fmt.Println(v17)
                                                    _ = len(v6)
                                                    var v19 uint = v8
                                                    _ = v17
                                                    _ = v18
                                                    _ = v19
                                                }
                                                p0 -= (p0 + (p0 << ((1) & 63)))
                                            } else {
                                                if (0 < 0) {
                                                    v8 += v8
                                                    v20 := GenS2{F0: 297, F1: 2147483647, F2: (("" + ("hello" + "\n\t")) != "foo bar")}
                                                    v21 := GenS2{F0: -2147483648, F1: 42, F2: false}
                                                    _ = v20
                                                    _ = v21
                                                }
                                                var v22 int16 = 1
                                                v23 := GenS3{F0: "a"}
                                                _ = v22
                                                _ = v23
                                            }
                                            _ = v16
                                        } else {
                                            var v24 int = (fi12 | (fi12 >> (((v8 >> (((^(v8))) & 63))) & 63)))
                                            var v25 int16 = (p1 | 32767)
                                            _ = v24
                                            _ = v25
                                        }
                                    case "\n\t":
                                        v26 := GenS2{F0: -298, F1: -851, F2: (!false)}
                                        _ = v26
                                    case "":
                                        v8 = v8
                                    default:
                                        if true {
                                            v6 = append(v6, fi7)
                                        } else {
                                            p1 = int16(v8)
                                            var v27 int32 = 367
                                            _ = v27
                                        }
                                    }
                                }
                            } else {
                                var v28 int8 = 127
                                var v29 int64 = -1
                                if (!false) {
                                    if true {
                                        if true {
                                            var v30 bool = true
                                            v29 += v29
                                            _ = v30
                                        } else {
                                            v31 := GenS1{F0: fi7, F1: (v28 ^ v28), F2: v8}
                                            p0 = p0
                                            var v32 int = fi7
                                            _ = v31
                                            _ = v32
                                        }
                                    }
                                    var v33 float32 = 1.5
                                    _ = v33
                                } else {
                                    v34 := GenS4{F0: (v28 * (v28 ^ int8(v15))), F1: p1}
                                    v8 = v8
                                    if true {
                                        var v35 int32 = 2147483647
                                        v36 := GenS5{F0: 2147483647, F1: 516, F2: "hello"}
                                        var v37 int32 = (v35 | v35)
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                    _ = v34
                                }
                                _ = v28
                                _ = v29
                            }
                            switch ("\n\t" + "hello") {
                            case "hello":
                                v8 -= 1
                            default:
                                if len(v6) > 0 {
                                    v6[0] = fi7
                                }
                            }
                            if (42 >= 42) {
                                v8 = v8
                                v15 = v15
                            }
                        } else {
                            var v38 uint64 = 18446744073709551615
                            _ = v38
                        }
                        v39 := GenS2{F0: 895, F1: 42, F2: (0 == 1)}
                        if (fi12 == -9223372036854775808) {
                            if len(v6) > 0 {
                                v6[0] = fi12
                            }
                            v40 := GenS3{F0: "\n\t"}
                            v8 = (v8 ^ (^(v8)))
                            _ = v40
                        } else {
                            v6 = append(v6, 127)
                        }
                        if true {
                            if true {
                                p0 -= p0
                                var v41 uint32 = 65535
                                _ = v41
                            } else {
                                if len(v6) > 0 {
                                    v6[0] = (fi12 + (fi12 * 127))
                                }
                                var v42 uint64 = 42
                                _ = v42
                            }
                            var v43 float64 = ((1.5 + 0.0) * ((0.1 / (float64((3.14159 / 1.5)) / ((1e10 * float64(1.5)) + (1e10 * (3.14159 + 1.5))))) - (0.0 - 0.1)))
                            v6 = append(v6, (fi7 + (^(fi12))))
                            if (!true) {
                                if (((p0 % ((132) | 1)) >= (p0 & p0)) && ((2147483648 <= 4294967295) && false)) {
                                    _ = len(v6)
                                    switch "hello" {
                                    case "\n\t":
                                        var v44 float32 = 1.5
                                        _ = v44
                                    default:
                                        v45 := GenS2{F0: 0, F1: 42, F2: (!(!false))}
                                        if len(v6) > 0 {
                                            v6[0] = fi7
                                        }
                                        if len(v6) > 0 {
                                            v6[0] = fi12
                                        }
                                        _ = v45
                                    }
                                    var v46 int64 = 0
                                    fmt.Println(fi12, v15, fi7)
                                    _ = v46
                                } else {
                                    var v47 int16 = p1
                                    _ = v47
                                }
                                v48 := GenS6{F0: 0, F1: (v15 / ((1) | 1))}
                                _ = v48
                            } else {
                                v15 = 252
                                if false {
                                    v15 = (v15 & (v15 << (((v8 + uint(v15))) & 63)))
                                    if true {
                                        v49 := GenS3{F0: ""}
                                        _ = v49
                                    } else {
                                        var v50 uint32 = 4294967295
                                        switch "a" {
                                        case "foo bar":
                                            switch fi7 {
                                            case 32767:
                                                fmt.Println(v50, v50, v43)
                                                var v51 uint16 = uint16(v8)
                                                if true {
                                                    var v52 int16 = 0
                                                    var v53 int = 45
                                                    p0 = (p0 ^ (^(p0)))
                                                    _ = v52
                                                    _ = v53
                                                } else {
                                                    var v54 uint32 = 800
                                                    fmt.Println(v51)
                                                    var v55 bool = (false && false)
                                                    _ = v54
                                                    _ = v55
                                                }
                                                if false {
                                                    var v56 bool = true
                                                    _ = v56
                                                }
                                                _ = v51
                                            }
                                        case "hello":
                                            switch fi7 {
                                            case 191:
                                                p1 += 72
                                            default:
                                                if ((v8 * (v8 ^ (v8 / ((4294967295) | 1)))) != (^(v8))) {
                                                    var v57 uint8 = 255
                                                    var v58 uint8 = (p0 >> ((v8) & 63))
                                                    _ = v57
                                                    _ = v58
                                                } else {
                                                    _ = len(v6)
                                                    var v59 string = "hello"
                                                    var v60 int = (fi7 / (((fi12 | (fi7 >> (((v8 >> ((uint(p0)) & 63))) & 63)))) | 1))
                                                    v6 = append(v6, (fi12 / ((128) | 1)))
                                                    var v61 string = ((v59 + (v59 + v59)) + ("\n\t" + v59))
                                                    fmt.Println(fi12, fi7)
                                                    var v62 uint = 400
                                                    var v63 int8 = -128
                                                    _ = v59
                                                    _ = v60
                                                    _ = v61
                                                    _ = v62
                                                    _ = v63
                                                }
                                            }
                                        }
                                        _ = v50
                                    }
                                    v15 = v15
                                } else {
                                    v64 := GenS4{F0: -1, F1: -32768}
                                    p1 = p1
                                    _ = len(v6)
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = v64
                                }
                                v65 := GenS3{F0: "\n\t"}
                                _ = v65
                            }
                            _ = v43
                        } else {
                            var v66 uint16 = (v15 + 733)
                            switch fi7 {
                            case 360:
                                switch fi7 {
                                case 32767:
                                    if (!((("hello" != ("\n\t" + ("hello" + ("" + "foo bar")))) || (-736 >= 52)) && false)) {
                                        switch fi12 {
                                        case -57:
                                            var v67 int64 = -1
                                            _ = v67
                                        case 1:
                                            switch fi7 {
                                            case 32768:
                                                if (!true) {
                                                    var v68 int = (-(fi7))
                                                    _ = v68
                                                } else {
                                                    if true {
                                                        if ("" <= ("foo bar" + "\n\t")) {
                                                            if false {
                                                                var v69 uint64 = 1
                                                                v15 += v15
                                                                if (((!(!(v8 != v8))) || false) || (true && (v15 > 42))) {
                                                                    v6 = append(v6, -9223372036854775808)
                                                                    var v70 int16 = p1
                                                                    _ = v70
                                                                }
                                                                _ = v69
                                                            }
                                                            v71 := GenS5{F0: 234, F1: 42, F2: ("foo bar" + "hello")}
                                                            _ = v71
                                                        }
                                                        v6 = append(v6, fi12)
                                                    }
                                                }
                                            case 42:
                                                fmt.Println(p0)
                                            default:
                                                v72 := GenS3{F0: ("\n\t" + ((("a" + "\n\t") + "") + "foo bar"))}
                                                _ = v72
                                            }
                                        case -619:
                                            var v73 int8 = 127
                                            _ = v73
                                        }
                                        v8 += (^(v8))
                                        v6 = append(v6, (fi7 - (^(fi12))))
                                        if false {
                                            if true {
                                                if true {
                                                    p1 = (p1 - (p1 * (p1 | p1)))
                                                    switch (-(fi12)) {
                                                    case 1:
                                                        var v74 uint32 = 4294967295
                                                        _ = v74
                                                    case -662:
                                                        if true {
                                                            if false {
                                                                var v75 uint8 = uint8(v15)
                                                                fmt.Println(v66)
                                                                _ = v75
                                                            }
                                                        }
                                                    default:
                                                        p0 = (p0 % ((p0) | 1))
                                                    }
                                                    v76 := GenS7{F0: ("" + "foo bar")}
                                                    _ = v76
                                                } else {
                                                    v8 += 255
                                                    v8 = v8
                                                    v8 -= (v8 / (((^(v8))) | 1))
                                                }
                                                v66 -= v66
                                                v77 := GenS4{F0: 42, F1: (^(p1))}
                                                _ = v77
                                            } else {
                                                if false {
                                                    v8 -= v8
                                                    if (-658 != 9223372036854775807) {
                                                        var v78 uint = 42
                                                        if true {
                                                            var v79 float64 = -2.25
                                                            v6 = append(v6, 597)
                                                            _ = v79
                                                        } else {
                                                            fmt.Println(fi7)
                                                            var v80 float32 = (3.14159 - (1.5 * (0.1 * ((1e10 / 0.1) * 3.14159))))
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            _ = v80
                                                        }
                                                        if (!true) {
                                                            var v81 uint64 = 42
                                                            v82 := GenS7{F0: "hello"}
                                                            var v83 float64 = (0.0 / (3.14159 - 1e10))
                                                            _ = v81
                                                            _ = v82
                                                            _ = v83
                                                        } else {
                                                            if true {
                                                                if (!(v15 <= (v15 * 573))) {
                                                                    var v84 int16 = p1
                                                                    _ = v84
                                                                } else {
                                                                    var v85 float32 = -2.25
                                                                    v86 := GenS8{F0: 1, F1: p0, F2: 1}
                                                                    _ = v85
                                                                    _ = v86
                                                                }
                                                                _ = len(v6)
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    fmt.Println(v8)
                                                                default:
                                                                    v6 = append(v6, fi7)
                                                                }
                                                            }
                                                            switch 2147483647 {
                                                            case 42:
                                                                v78 = (v78 + v8)
                                                            case -993:
                                                                p1 -= (p1 - (p1 - p1))
                                                            case 4294967295:
                                                                p1 = 1
                                                            default:
                                                                v87 := GenS3{F0: ("hello" + "\n\t")}
                                                                _ = v87
                                                            }
                                                        }
                                                        _ = v78
                                                    }
                                                    if true {
                                                        if ((!false) || true) {
                                                            v88 := GenS3{F0: "foo bar"}
                                                            _ = v88
                                                        } else {
                                                            var v89 float64 = 1.5
                                                            var v90 string = (("a" + ("" + "a")) + "hello")
                                                            if true {
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                if (v15 >= uint16(v8)) {
                                                                    v89 = (v89 / (-2.25 - (float64(0.0) * 0.0)))
                                                                } else {
                                                                    fmt.Println(v15, v66, v66)
                                                                }
                                                            }
                                                            var v91 uint64 = 1
                                                            _ = v89
                                                            _ = v90
                                                            _ = v91
                                                        }
                                                        var v92 int8 = int8(p0)
                                                        _ = v92
                                                    } else {
                                                        v6 = append(v6, (-(fi7)))
                                                        fmt.Println(fi12, p0)
                                                        v93 := GenS2{F0: 34, F1: -773, F2: true}
                                                        _ = v93
                                                    }
                                                } else {
                                                    switch ("a" + ("a" + "a")) {
                                                    case "foo bar":
                                                        fmt.Println(v66, v8, p1)
                                                    case "":
                                                        p0 -= (p0 & 2)
                                                        v66 = (v66 / (((v15 | 32767)) | 1))
                                                    case "a":
                                                        _ = len(v6)
                                                    default:
                                                        v94 := GenS4{F0: 127, F1: (p1 / ((p1) | 1))}
                                                        _ = v94
                                                    }
                                                }
                                            }
                                            v95 := GenS9{F0: (0 <= 42)}
                                            v8 = 18446744073709551615
                                            if len(v6) > 0 {
                                                v6[0] = -2147483648
                                            }
                                            var v96 bool = false
                                            _ = v95
                                            _ = v96
                                        } else {
                                            p1 = 523
                                            switch (fi7 % (((fi12 & fi7)) | 1)) {
                                            case 0:
                                                p0 = 63
                                                fmt.Println(fi12, fi12)
                                                if (1 > 0) {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    v97 := GenS10{F0: 2147483647, F1: true, F2: (fi7 ^ int(p0))}
                                                    var v98 uint16 = (v66 << ((42) & 63))
                                                    _ = v97
                                                    _ = v98
                                                }
                                            default:
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            }
                                            if true {
                                                var v99 uint16 = 0
                                                _ = v99
                                            } else {
                                                _ = len(v6)
                                            }
                                        }
                                    } else {
                                        switch (fi7 + fi12) {
                                        case -2147483648:
                                            var v100 uint16 = v66
                                            _ = v100
                                        case 42:
                                            if ((fi12 / ((682) | 1)) == (fi7 >> (((v8 << ((4294967295) & 63))) & 63))) {
                                                if len(v6) > 0 {
                                                    v6[0] = fi12
                                                }
                                                switch "hello" {
                                                case "foo bar":
                                                    var v101 uint32 = 721
                                                    var v102 float64 = -2.25
                                                    if ((((float64(0.0) + float64(0.1)) + 0.0) < v102) || true) {
                                                        var v103 uint32 = v101
                                                        v104 := GenS1{F0: (fi7 * (fi7 << ((v8) & 63))), F1: 127, F2: v8}
                                                        v101 = 188
                                                        _ = v103
                                                        _ = v104
                                                    } else {
                                                        switch "a" {
                                                        case "foo bar":
                                                            v105 := GenS11{F0: 0, F1: true, F2: 2147483647}
                                                            _ = v105
                                                        default:
                                                            if len(v6) > 0 {
                                                                v6[0] = fi7
                                                            }
                                                        }
                                                        p1 -= p1
                                                    }
                                                    _ = v101
                                                    _ = v102
                                                }
                                                fmt.Println(fi7, p0, v15)
                                            } else {
                                                if false {
                                                    var v106 uint16 = 32767
                                                    p0 -= 2
                                                    v15 -= (v106 >> (((v8 << ((335) & 63))) & 63))
                                                    _ = v106
                                                } else {
                                                    if false {
                                                        var v107 int = (fi7 * fi12)
                                                        var v108 uint64 = 386
                                                        _ = v107
                                                        _ = v108
                                                    }
                                                }
                                                v109 := GenS12{F0: "", F1: v8, F2: -1}
                                                _ = v109
                                            }
                                        }
                                        v110 := GenS1{F0: fi7, F1: -128, F2: uint(p0)}
                                        v111 := GenS11{F0: 538, F1: true, F2: 42}
                                        _ = v110
                                        _ = v111
                                    }
                                    var v112 float32 = (-2.25 * ((float32((((float64(-2.25) * (float64((((((0.1 * 3.14159) / 1.0) / 0.1) * ((1e10 * -2.25) + (0.1 * -2.25))) / 3.14159)) + 3.14159)) / (((float64((0.1 * 3.14159)) / (((((1.5 - 3.14159) + 1e10) / -2.25) / 1e10) * (0.1 - ((3.14159 / 0.1) / 0.1)))) * 1.5) / -2.25)) * (float64((-2.25 + -2.25)) * -2.25))) / float32(float64(3.14159))) + 0.0))
                                    fmt.Println(fi12)
                                    _ = v112
                                case 577:
                                    p1 -= (p1 ^ (p1 % ((p1) | 1)))
                                case -119:
                                    v15 = (^(v66))
                                }
                            default:
                                switch (fi7 + (^(fi12))) {
                                case -750:
                                    var v113 int32 = 42
                                    _ = v113
                                case 1:
                                    if true {
                                        v114 := GenS13{F0: 0, F1: v8, F2: 403}
                                        fmt.Println(v15, fi7)
                                        _ = v114
                                    }
                                default:
                                    fmt.Println(p1, v15, fi12)
                                }
                            }
                            var v115 bool = false
                            if false {
                                fmt.Println(fi12)
                                v8 += 255
                            } else {
                                v116 := GenS14{F0: "a", F1: 32767}
                                v117 := GenS15{F0: -1, F1: 2}
                                var v118 uint = 42
                                _ = v116
                                _ = v117
                                _ = v118
                            }
                            _ = len(v6)
                            _ = v115
                            _ = v66
                        }
                        v119 := GenS16{F0: 1}
                        _ = v119
                        _ = v15
                        _ = v39
                    }
                    var v120 uint64 = 4294967295
                    _ = v120
                }
                v121 := GenS13{F0: 1, F1: (^(v8)), F2: -1}
                switch "foo bar" {
                case "foo bar":
                    if true {
                        var v122 uint8 = p0
                        if (!false) {
                            var v123 float64 = 0.1
                            var v124 int32 = 2147483647
                            if true {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                var v125 uint16 = 2
                                if true {
                                    v125 -= uint16(v122)
                                    if (!(-9223372036854775808 >= -9223372036854775808)) {
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v126 uint16 = (v125 + (v125 - (v125 - v125)))
                                            _ = v126
                                        case "hello":
                                            if (!false) {
                                                v127 := GenS17{F0: 0, F1: fi7}
                                                _ = v127
                                            }
                                        default:
                                            if len(v6) > 0 {
                                                v6[0] = fi12
                                            }
                                        }
                                    }
                                }
                                _ = v125
                            } else {
                                v124 = (^(v124))
                                switch ("\n\t" + "foo bar") {
                                case "a":
                                    var v128 bool = false
                                    _ = v128
                                default:
                                    if true {
                                        v129 := GenS18{F0: (("" + "a") + "foo bar"), F1: "foo bar"}
                                        var v130 uint64 = 4294967295
                                        _ = v129
                                        _ = v130
                                    }
                                }
                                v8 = 436
                            }
                            _ = v123
                            _ = v124
                        }
                        switch "\n\t" {
                        case "a":
                            var v131 uint32 = 876
                            _ = v131
                        case "\n\t":
                            v6 = append(v6, (fi12 & -905))
                        default:
                            var v132 uint32 = 42
                            _ = v132
                        }
                        _ = v122
                    }
                }
                _ = fi12
                _ = v121
            }
        } else {
            if (!(true || (!false))) {
                switch "" {
                case "\n\t":
                    fmt.Println(fi7)
                default:
                    p1 += p1
                }
            } else {
                if len(v6) > 0 {
                    v6[0] = -9223372036854775808
                }
                p0 -= p0
            }
            switch fi7 {
            case 1:
                p0 -= (p0 | p0)
            default:
                p0 = 58
            }
        }
        var v133 int8 = -128
        _ = fi7
        _ = v133
        _ = v8
    }
    v134 := GenS3{F0: (("foo bar" + "a") + "hello")}
    _ = p0
    _ = p1
    _ = v134
    _ = v6
    return "a"
}

func main() {
    var v1 uint32 = 824
    var v2 int = 2147483647
    v3 := make(map[uint32]bool)
    var v4 int64 = 42
    v5 := make(map[uint32]int)
    var v135 float32 = -2.25
    if true {
        var v136 uint32 = v1
        _ = v136
    } else {
        v2 = -839
        v137 := GenS19{F0: 113}
        _ = v137
    }
    fmt.Println(v1, v135)
    for fi138 := 0; fi138 < 3; fi138++ {
        if (true && (v4 < (v4 % ((v4) | 1)))) {
            if true {
                var v139 uint32 = v1
                var v140 int32 = 42
                _ = v139
                _ = v140
            }
            v2 -= 1
        }
        v135 = (0.1 - 1.5)
        _ = fi138
    }
    v141 := GenS17{F0: 2, F1: v2}
    delete(v5, 1)
    v142 := GenS12{F0: "a", F1: 18446744073709551615, F2: 1}
    for fi143 := 0; fi143 < 4; fi143++ {
        if false {
            var v144 uint8 = 2
            switch v2 {
            case 591:
                v145 := GenS20{F0: 839, F1: 65535}
                _ = v145
            case -1:
                fmt.Println(v1, v135, v144)
            }
            var v146 string = ("foo bar" + "foo bar")
            _ = v144
            _ = v146
        } else {
            var v147 bool = ((("hello" + "foo bar") == "\n\t") || (42 < 0))
            for fi148 := 0; fi148 < 18; fi148++ {
                v149 := []int{-537, fi143, fi143, 1}
                _ = fi148
                _ = v149
            }
            _ = v147
        }
        v1 -= (v1 & v1)
        _ = fi143
    }
    var v150 uint = 65535
    v4 = int64(v150)
    if (1 != (v4 / (((^(v4))) | 1))) {
        v135 = 0.1
        for fi151 := 0; fi151 < 18; fi151++ {
            v4 += v4
            switch ("hello" + "") {
            case "hello":
                for fi152 := 0; fi152 < 14; fi152++ {
                    switch int(v150) {
                    case -424:
                        if (0.0 <= 1e10) {
                            if false {
                                v135 += v135
                                if true {
                                    v2 = (fi151 / ((fi151) | 1))
                                } else {
                                    fmt.Println(fi152)
                                }
                                v153 := []int{(fi152 | fi151), 1, fi151}
                                v154 := GenS11{F0: 753, F1: ((v1 * (v1 & 165)) == 2), F2: 42}
                                _ = v153
                                _ = v154
                            } else {
                                switch 525 {
                                case 966:
                                    v150 += v150
                                case 9223372036854775807:
                                    v4 += v4
                                }
                                v3[1] = (104 >= 127)
                                switch (fi151 | (^(v2))) {
                                case 128:
                                    v150 = uint(v1)
                                }
                            }
                            var v155 float32 = (v135 * v135)
                            if (true && (!(32767 > 255))) {
                                v1 = 2147483648
                            } else {
                                v156 := GenS19{F0: 2}
                                v157 := GenS14{F0: "\n\t", F1: -32768}
                                _ = v156
                                _ = v157
                            }
                            _ = v155
                        } else {
                            if (1 == 0) {
                                if ((1.5 > ((0.1 - 3.14159) * 0.1)) || false) {
                                    v4 = 42
                                    var v158 int32 = -2147483648
                                    var v159 uint8 = 1
                                    _ = v158
                                    _ = v159
                                }
                                v135 += (v135 - v135)
                                var v160 uint64 = 999
                                switch "" {
                                case "a":
                                    v161 := GenS12{F0: "\n\t", F1: (v150 % (((v150 % ((v150) | 1))) | 1)), F2: -2147483648}
                                    switch v2 {
                                    case -9223372036854775808:
                                        v162 := v3[v1]
                                        _ = v162
                                    case -227:
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v163 uint = v150
                                            _ = v163
                                        default:
                                            switch fi152 {
                                            case 9223372036854775807:
                                                fmt.Println(v150)
                                            case -779:
                                                v4 += v4
                                                v164 := GenS1{F0: -2147483648, F1: -128, F2: v150}
                                                v165 := GenS21{F0: 255}
                                                _ = v164
                                                _ = v165
                                            }
                                        }
                                    case 9223372036854775807:
                                        var v166 float32 = 0.1
                                        _ = v166
                                    }
                                    v167 := []int{}
                                    _ = v161
                                    _ = v167
                                case "":
                                    v168 := GenS22{F0: 0, F1: 0, F2: -1}
                                    _ = v168
                                default:
                                    v160 += v160
                                }
                                var v169 int16 = 1
                                v170 := GenS12{F0: "", F1: v150, F2: -127}
                                _ = v160
                                _ = v169
                                _ = v170
                            }
                            v171 := GenS20{F0: v150, F1: 1}
                            var v172 int8 = 119
                            _ = v171
                            _ = v172
                        }
                    }
                    _ = fi152
                }
            case "\n\t":
                var v173 int32 = 1
                _ = v173
            default:
                v174 := []int{(v2 >> ((v150) & 63))}
                _ = v174
            }
            _ = fi151
        }
        switch "a" {
        case "":
            for fi175 := 0; fi175 < 14; fi175++ {
                if ((v4 >> (((v150 >> ((42) & 63))) & 63)) > 372) {
                    v176 := GenS22{F0: v1, F1: 32767, F2: 127}
                    _ = v176
                }
                var v177 int8 = 0
                fmt.Println(v135, fi175, v4)
                _ = fi175
                _ = v177
            }
        case "hello":
            v178 := GenS8{F0: (v1 << ((v150) & 63)), F1: 128, F2: 946}
            _ = v178
        }
    } else {
        var v179 uint16 = 2
        _ = v179
    }
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("v150:", v150)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v135
    _ = v141
    _ = v142
    _ = v150
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
