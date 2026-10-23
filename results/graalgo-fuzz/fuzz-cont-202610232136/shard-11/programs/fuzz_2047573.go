package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 int; F2 uint64 }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint; F1 uint64 }
type GenS9 struct { F0 int16; F1 string; F2 uint16 }
type GenS10 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS11 struct { F0 int; F1 int8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 string; F1 int }
type GenS14 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 uint32; F1 int8; F2 int64 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int16; F1 uint64 }
type GenS19 struct { F0 int32; F1 int64; F2 uint16 }
type GenS20 struct { F0 int32; F1 uint64; F2 int64 }
type GenS21 struct { F0 int32; F1 int; F2 uint64 }
type GenS22 struct { F0 int32 }
type GenS23 struct { F0 int16; F1 uint64 }
type GenS24 struct { F0 int8; F1 int16; F2 int }
type GenS25 struct { F0 bool; F1 string; F2 int16 }
type GenS26 struct { F0 int }
type GenS27 struct { F0 uint64; F1 int32 }
type GenS28 struct { F0 uint; F1 string; F2 int32 }
func genFn0(p0 uint, p1 uint16, p2 int8) uint64 {
    switch 0 {
    case -2147483648:
        var v7 int = 298
        _ = v7
    default:
        var v8 int32 = -895
        _ = v8
    }
    if false {
        p0 += (p0 / ((p0) | 1))
        p2 -= (p2 >> ((p0) & 63))
    } else {
        var v9 int8 = int8(p0)
        v10 := []int{}
        for fi11 := 0; fi11 < 17; fi11++ {
            fmt.Println(fi11, fi11, p2)
            var v12 uint64 = 194
            for fi13 := 0; fi13 < 20; fi13++ {
                p1 -= 0
                p2 = (p2 % ((v9) | 1))
                if true {
                    var v14 string = ((("\n\t" + "a") + "foo bar") + "hello")
                    if (!true) {
                        if (2 > uint32(v12)) {
                            if len(v10) > 0 {
                                v10[0] = -1
                            }
                            switch (fi13 & (fi11 ^ fi13)) {
                            case 9223372036854775807:
                                if true {
                                    switch (fi13 + 4294967295) {
                                    case 1:
                                        v15 := GenS1{F0: 0, F1: ((-9223372036854775808 != 184) || (700 <= 32767))}
                                        var v16 uint = 403
                                        _ = v15
                                        _ = v16
                                    case 2147483647:
                                        if false {
                                            p0 = p0
                                            fmt.Println(v14, p1)
                                        }
                                    case -1:
                                        var v17 int8 = (v9 % ((p2) | 1))
                                        _ = v17
                                    default:
                                        var v18 int16 = 541
                                        _ = v18
                                    }
                                    v19 := GenS1{F0: 32767, F1: (-2.25 >= float64(3.14159))}
                                    if ((float64(3.14159) + 1e10) >= 3.14159) {
                                        var v20 bool = (1000 > 42)
                                        _ = v20
                                    } else {
                                        var v21 string = "a"
                                        v9 = 127
                                        var v22 float64 = (0.0 / 1.5)
                                        _ = v21
                                        _ = v22
                                    }
                                    _ = v19
                                }
                            case 42:
                                if false {
                                    var v23 float32 = (-2.25 + 1.5)
                                    _ = v23
                                } else {
                                    p1 = p1
                                    if len(v10) > 0 {
                                        v10[0] = 42
                                    }
                                }
                            case -975:
                                var v24 uint8 = 128
                                _ = v24
                            default:
                                var v25 int8 = (v9 / ((v9) | 1))
                                _ = v25
                            }
                            v26 := GenS1{F0: 1000, F1: false}
                            _ = v26
                        }
                        if true {
                            switch 867 {
                            case 128:
                                switch v14 {
                                case "hello":
                                    var v27 float32 = 1.5
                                    _ = v27
                                case "a":
                                    v12 += v12
                                    v28 := GenS1{F0: 861, F1: (((-2.25 / (1.5 - 0.0)) * 0.0) > (3.14159 - (0.0 + 1e10)))}
                                    v14 = "\n\t"
                                    fmt.Println(v12, p1, v14)
                                    _ = v28
                                case "foo bar":
                                    var v29 float64 = 3.14159
                                    _ = v29
                                default:
                                    p2 -= 15
                                }
                            }
                            if false {
                                fmt.Println(fi11, fi13)
                                if ("hello" == v14) {
                                    var v30 int32 = 1
                                    if true {
                                        if (!(true || (true || true))) {
                                            v31 := GenS1{F0: 1000, F1: (!false)}
                                            _ = v31
                                        } else {
                                            if (!true) {
                                                var v32 int8 = (v9 / ((1) | 1))
                                                _ = v32
                                            } else {
                                                v33 := GenS1{F0: -842, F1: false}
                                                _ = v33
                                            }
                                            fmt.Println(v14, v30, v9)
                                            p0 = (^(p0))
                                        }
                                        v34 := GenS1{F0: 32767, F1: (fi13 <= (fi13 + fi13))}
                                        _ = v34
                                    } else {
                                        v30 -= v30
                                    }
                                    _ = v30
                                } else {
                                    p1 -= p1
                                    if true {
                                        switch v14 {
                                        case "\n\t":
                                            fmt.Println(p2)
                                        default:
                                            v12 = (v12 * 812)
                                        }
                                        if (p1 <= 771) {
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                        } else {
                                            switch (v14 + (v14 + (v14 + v14))) {
                                            case "":
                                                if true {
                                                    if false {
                                                        var v35 float32 = (3.14159 - (-2.25 * ((-2.25 - 1e10) / (1e10 / 1.5))))
                                                        if len(v10) >= 2 {
                                                            v10 = v10[1:len(v10)]
                                                        }
                                                        var v36 int32 = 2147483647
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                    v37 := GenS1{F0: -1, F1: ((!(4294967295 <= fi13)) || (0.0 != 1e10))}
                                                    _ = v37
                                                }
                                            default:
                                                var v38 int64 = 9223372036854775807
                                                _ = v38
                                            }
                                            var v39 uint32 = 0
                                            v9 = (p2 * (v9 << (((p0 % (((^(p0))) | 1))) & 63)))
                                            _ = v39
                                        }
                                    } else {
                                        var v40 uint32 = 2147483647
                                        _ = v40
                                    }
                                }
                            }
                            v41 := GenS2{F0: p0, F1: fi11}
                            if (true || true) {
                                var v42 float64 = (1.5 + 3.14159)
                                _ = v42
                            }
                            var v43 float64 = 1e10
                            _ = v41
                            _ = v43
                        }
                        switch fi13 {
                        case 127:
                            p2 -= (^(v9))
                        default:
                            var v44 bool = false
                            _ = v44
                        }
                        if true {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            switch (fi13 + int(v9)) {
                            case -825:
                                p1 = p1
                            case -9223372036854775808:
                                var v45 bool = (!(669 == 42))
                                _ = v45
                            case 32767:
                                var v46 uint32 = 2147483648
                                _ = v46
                            default:
                                v12 -= 0
                            }
                            var v47 uint16 = (p1 * p1)
                            _ = v47
                        }
                    }
                    var v48 bool = false
                    _ = v14
                    _ = v48
                } else {
                    if ((!("foo bar" > ("hello" + (("foo bar" + ("" + ((("foo bar" + "foo bar") + "a") + (("a" + ((("a" + (("" + "foo bar") + ("a" + ("foo bar" + (((("hello" + "") + "") + "hello") + "foo bar"))))) + "\n\t") + ((("a" + "\n\t") + "foo bar") + "foo bar"))) + "\n\t")))) + ("hello" + ("foo bar" + ("foo bar" + "a"))))))) && (uint8(v12) > uint8(p1))) {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    } else {
                        v9 = int8(fi11)
                    }
                    if (!false) {
                        v49 := GenS1{F0: 32767, F1: ((v9 - (v9 >> ((p0) & 63))) <= 16)}
                        fmt.Println(p1, fi11)
                        v50 := GenS3{F0: (p1 >> ((42) & 63)), F1: 42}
                        p0 = uint(p1)
                        fmt.Println(v12, p0, v9)
                        _ = v49
                        _ = v50
                    }
                    if true {
                        if (false && (!true)) {
                            p0 += p0
                            fmt.Println(fi13, v9, p1)
                        } else {
                            if len(v10) > 0 {
                                v10[0] = fi11
                            }
                        }
                        if ((4294967295 < 2) || false) {
                            if true {
                                switch fi13 {
                                case 42:
                                    p1 -= 65535
                                case 449:
                                    if (0 > 2147483647) {
                                        v51 := GenS4{F0: 282}
                                        if ((9223372036854775807 < -9223372036854775808) || true) {
                                            var v52 int64 = 9223372036854775807
                                            var v53 uint32 = 42
                                            if true {
                                                v12 = v12
                                                var v54 uint32 = (v53 ^ (v53 - (^(v53))))
                                                _ = v54
                                            }
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                            var v55 uint8 = 2
                                            _ = v52
                                            _ = v53
                                            _ = v55
                                        }
                                        _ = v51
                                    }
                                case 128:
                                    if len(v10) > 0 {
                                        v10[0] = (fi11 & fi13)
                                    }
                                }
                                var v56 uint16 = (p1 >> (((p0 >> ((uint(v12)) & 63))) & 63))
                                v57 := GenS3{F0: 32, F1: 666}
                                _ = v56
                                _ = v57
                            } else {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                v58 := GenS4{F0: 1}
                                var v59 uint = (p0 / (((p0 / ((p0) | 1))) | 1))
                                _ = v58
                                _ = v59
                            }
                            var v60 int8 = (p2 & (-(v9)))
                            v61 := GenS5{F0: uint32(p0), F1: (fi13 | fi13), F2: (v12 ^ v12)}
                            v62 := GenS6{F0: 73, F1: -614}
                            switch ("" + "\n\t") {
                            case "hello":
                                p2 = -128
                            default:
                                if (-474 < 1) {
                                    p2 += (^(p2))
                                    fmt.Println(v9)
                                    if true {
                                        if false {
                                            v63 := GenS7{F0: v60}
                                            v64 := GenS8{F0: p0, F1: 1}
                                            v9 = -15
                                            _ = v63
                                            _ = v64
                                        } else {
                                            if ((v12 >= v12) && false) {
                                                var v65 int32 = 2147483647
                                                var v66 int32 = int32(v12)
                                                v67 := GenS9{F0: 236, F1: ("foo bar" + "a"), F2: (p1 ^ (p1 | (p1 >> ((626) & 63))))}
                                                fmt.Println(v66, v60)
                                                switch (("" + "foo bar") + (((("foo bar" + "hello") + "a") + "") + "hello")) {
                                                case "foo bar":
                                                    var v68 float32 = -2.25
                                                    _ = v68
                                                case "":
                                                    fmt.Println(fi11, v60)
                                                }
                                                _ = v65
                                                _ = v66
                                                _ = v67
                                            }
                                            v69 := GenS10{F0: 2, F1: 42, F2: (p1 + (^(p1)))}
                                            switch "" {
                                            case "":
                                                v70 := GenS11{F0: 570, F1: (p2 - v60)}
                                                if true {
                                                    fmt.Println(v9)
                                                    if false {
                                                        p0 += p0
                                                    } else {
                                                        var v71 int32 = -516
                                                        switch fi13 {
                                                        case -2147483648:
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                        case -625:
                                                            switch fi11 {
                                                            case 4294967295:
                                                                var v72 int = (fi13 ^ fi11)
                                                                _ = v72
                                                            case 32767:
                                                                p2 = (v60 - v9)
                                                            case 893:
                                                                v12 += 42
                                                            }
                                                        case -984:
                                                            var v73 uint32 = 4294967295
                                                            _ = v73
                                                        default:
                                                            var v74 int32 = (v71 & (v71 - (v71 << (((p0 % ((42) | 1))) & 63))))
                                                            _ = v74
                                                        }
                                                        _ = v71
                                                    }
                                                    if true {
                                                        v75 := GenS5{F0: 2, F1: 127, F2: 721}
                                                        _ = v75
                                                    } else {
                                                        var v76 float64 = 0.0
                                                        v77 := GenS7{F0: v60}
                                                        if len(v10) >= 2 {
                                                            v10 = v10[1:len(v10)]
                                                        }
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    p1 += p1
                                                    v10 = append(v10, fi11)
                                                    switch "\n\t" {
                                                    case "hello":
                                                        v60 = -1
                                                        var v78 int64 = -1
                                                        _ = v78
                                                    case "a":
                                                        switch fi13 {
                                                        case 1:
                                                            v79 := GenS12{F0: (("a" + "a") + (("\n\t" + (("" + "\n\t") + "")) + "hello"))}
                                                            _ = v79
                                                        case 42:
                                                            v60 = (p2 ^ (v60 | v9))
                                                        }
                                                        if true {
                                                            v80 := GenS13{F0: "hello", F1: (fi13 - 32768)}
                                                            switch "foo bar" {
                                                            case "a":
                                                                if true {
                                                                    v12 += uint64(p1)
                                                                } else {
                                                                    switch fi13 {
                                                                    case -97:
                                                                        fmt.Println(p2)
                                                                    case -986:
                                                                        var v81 float32 = 1e10
                                                                        _ = v81
                                                                    case 345:
                                                                        p2 = (v9 | p2)
                                                                    default:
                                                                        v9 = 103
                                                                    }
                                                                    if len(v10) > 0 {
                                                                        v10[0] = (fi11 << ((p0) & 63))
                                                                    }
                                                                    fmt.Println(v60, v60, p1)
                                                                }
                                                            }
                                                            v10 = append(v10, (^(fi11)))
                                                            _ = v80
                                                        }
                                                        v9 = p2
                                                        var v82 uint64 = v12
                                                        _ = v82
                                                    case "":
                                                        if ((true && (0.1 >= ((1.5 + float32(1.5)) * -2.25))) || true) {
                                                            v12 += 963
                                                            switch ((("hello" + ("hello" + "foo bar")) + ("a" + ("foo bar" + ""))) + ("\n\t" + "")) {
                                                            case "foo bar":
                                                                var v83 uint8 = 128
                                                                _ = v83
                                                            case "a":
                                                                p0 += (p0 * (p0 + (p0 >> ((569) & 63))))
                                                                v12 = 42
                                                            case "hello":
                                                                switch (fi11 >> (((p0 % ((18446744073709551615) | 1))) & 63)) {
                                                                case 128:
                                                                    switch "a" {
                                                                    case "a":
                                                                        if false {
                                                                            _ = len(v10)
                                                                        } else {
                                                                            switch "hello" {
                                                                            case "a":
                                                                                p2 += p2
                                                                            case "":
                                                                                if true {
                                                                                    var v84 int16 = -164
                                                                                    v85 := GenS8{F0: p0, F1: (^(v12))}
                                                                                    p2 = 114
                                                                                    _ = v84
                                                                                    _ = v85
                                                                                } else {
                                                                                    if len(v10) >= 2 {
                                                                                        v10 = v10[1:len(v10)]
                                                                                    }
                                                                                }
                                                                                if ((!false) || false) {
                                                                                    v9 += 42
                                                                                    if true {
                                                                                        v12 = 513
                                                                                    }
                                                                                    p0 -= 42
                                                                                } else {
                                                                                    p0 = uint(v12)
                                                                                }
                                                                                _ = len(v10)
                                                                            case "\n\t":
                                                                                var v86 int = (fi13 ^ (^(fi13)))
                                                                                _ = v86
                                                                            default:
                                                                                if (1 == 1000) {
                                                                                    switch "a" {
                                                                                    case "hello":
                                                                                        if false {
                                                                                            if len(v10) > 0 {
                                                                                                v10[0] = int(v12)
                                                                                            }
                                                                                            fmt.Println(p1, p2, fi11)
                                                                                        }
                                                                                    case "a":
                                                                                        if len(v10) >= 2 {
                                                                                            v10 = v10[1:len(v10)]
                                                                                        }
                                                                                    case "":
                                                                                        p2 -= v60
                                                                                    }
                                                                                    if true {
                                                                                        if true {
                                                                                            fmt.Println(fi11)
                                                                                        } else {
                                                                                            switch fi11 {
                                                                                            case -164:
                                                                                                v87 := GenS14{F0: p0, F1: 34, F2: (v12 * (v12 / ((v12) | 1)))}
                                                                                                _ = v87
                                                                                            case 128:
                                                                                                var v88 uint8 = 127
                                                                                                _ = v88
                                                                                            case 127:
                                                                                                if len(v10) >= 2 {
                                                                                                    v10 = v10[1:len(v10)]
                                                                                                }
                                                                                            }
                                                                                            v12 -= v12
                                                                                            if len(v10) >= 2 {
                                                                                                v10 = v10[1:len(v10)]
                                                                                            }
                                                                                            var v89 float64 = (3.14159 + 3.14159)
                                                                                            _ = v89
                                                                                        }
                                                                                        var v90 uint8 = 1
                                                                                        fmt.Println(p1, v90, fi13)
                                                                                        _ = v90
                                                                                    }
                                                                                } else {
                                                                                    if (int16(p0) > -931) {
                                                                                        var v91 bool = false
                                                                                        if len(v10) > 0 {
                                                                                            v10[0] = fi13
                                                                                        }
                                                                                        switch fi13 {
                                                                                        case -243:
                                                                                            if (184 > 128) {
                                                                                                if len(v10) > 0 {
                                                                                                    v10[0] = -479
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        var v92 int16 = -32768
                                                                                        fmt.Println(v12, p2)
                                                                                        _ = v91
                                                                                        _ = v92
                                                                                    } else {
                                                                                        v93 := GenS10{F0: uint32(v12), F1: 42, F2: (^(p1))}
                                                                                        _ = v93
                                                                                    }
                                                                                    if ((!((((-2.25 * float64(0.0)) / 3.14159) + 1e10) > (float64(1.5) - (1e10 + 0.0)))) || true) {
                                                                                        p2 -= p2
                                                                                        v94 := GenS15{F0: 1}
                                                                                        v60 += (v60 / (((-(p2))) | 1))
                                                                                        _ = v94
                                                                                    } else {
                                                                                        switch (fi11 ^ 257) {
                                                                                        case 0:
                                                                                            switch -916 {
                                                                                            case -616:
                                                                                                if false {
                                                                                                    fmt.Println(p0, p2, fi13)
                                                                                                    var v95 uint16 = p1
                                                                                                    p1 -= v95
                                                                                                    _ = v95
                                                                                                }
                                                                                            case 105:
                                                                                                var v96 uint = p0
                                                                                                _ = v96
                                                                                            }
                                                                                        case 766:
                                                                                            switch (fi13 * 1) {
                                                                                            case -670:
                                                                                                var v97 uint = p0
                                                                                                v98 := GenS2{F0: (v97 << (((p0 % (((p0 ^ (^(v97)))) | 1))) & 63)), F1: fi11}
                                                                                                p1 = 0
                                                                                                _ = v97
                                                                                                _ = v98
                                                                                            default:
                                                                                                var v99 float32 = 0.1
                                                                                                _ = v99
                                                                                            }
                                                                                        case -577:
                                                                                            if true {
                                                                                                v9 -= 1
                                                                                                fmt.Println(p0)
                                                                                            } else {
                                                                                                var v100 bool = true
                                                                                                p0 += p0
                                                                                                v10 = append(v10, 1)
                                                                                                _ = v100
                                                                                            }
                                                                                            fmt.Println(fi13, p2)
                                                                                            fmt.Println(v9, p1)
                                                                                            if false {
                                                                                                if false {
                                                                                                    v101 := GenS5{F0: 65535, F1: fi13, F2: (v12 | (v12 % ((v12) | 1)))}
                                                                                                    var v102 uint32 = uint32(v12)
                                                                                                    if len(v10) > 0 {
                                                                                                        v10[0] = (fi13 ^ (fi11 >> (((p0 / ((562) | 1))) & 63)))
                                                                                                    }
                                                                                                    _ = v101
                                                                                                    _ = v102
                                                                                                } else {
                                                                                                    var v103 int64 = int64(fi13)
                                                                                                    var v104 string = ("\n\t" + "hello")
                                                                                                    if ((^(v103)) < (v103 & v103)) {
                                                                                                        var v105 float64 = 1.5
                                                                                                        _ = v105
                                                                                                    } else {
                                                                                                        if len(v10) > 0 {
                                                                                                            v10[0] = (fi13 | fi11)
                                                                                                        }
                                                                                                        var v106 float32 = (((3.14159 + (0.1 / 0.1)) * ((3.14159 + 1e10) - 0.0)) + 1e10)
                                                                                                        _ = v106
                                                                                                    }
                                                                                                    _ = v103
                                                                                                    _ = v104
                                                                                                }
                                                                                                if len(v10) > 0 {
                                                                                                    v10[0] = (-(fi11))
                                                                                                }
                                                                                                var v107 int8 = (v9 / ((0) | 1))
                                                                                                v108 := GenS2{F0: (p0 - (p0 >> ((p0) & 63))), F1: -924}
                                                                                                _ = len(v10)
                                                                                                _ = v107
                                                                                                _ = v108
                                                                                            }
                                                                                        default:
                                                                                            v10 = append(v10, fi11)
                                                                                        }
                                                                                    }
                                                                                    _ = len(v10)
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        if ((true && (-32768 == 1)) || ("foo bar" > (("\n\t" + ("a" + ("foo bar" + "hello"))) + "a"))) {
                                                                            v109 := GenS16{F0: 2, F1: p2, F2: 9223372036854775807}
                                                                            _ = v109
                                                                        } else {
                                                                            v110 := GenS17{F0: 2}
                                                                            _ = v110
                                                                        }
                                                                    }
                                                                default:
                                                                    _ = len(v10)
                                                                }
                                                            default:
                                                                switch (fi13 + (fi11 ^ -9223372036854775808)) {
                                                                case -9223372036854775808:
                                                                    if len(v10) > 0 {
                                                                        v10[0] = fi13
                                                                    }
                                                                case 1:
                                                                    v12 = v12
                                                                case 393:
                                                                    switch fi11 {
                                                                    case -2147483648:
                                                                        var v111 int32 = -2147483648
                                                                        _ = v111
                                                                    case -1:
                                                                        v112 := GenS15{F0: 111}
                                                                        _ = v112
                                                                    }
                                                                default:
                                                                    switch 128 {
                                                                    case 32768:
                                                                        if (((3.14159 / 1.5) / 1.0) == (-2.25 * 1e10)) {
                                                                            var v113 int8 = v9
                                                                            var v114 int64 = 9223372036854775807
                                                                            p0 += p0
                                                                            v115 := GenS18{F0: 1000, F1: 2}
                                                                            switch ("a" + ("" + ("" + "foo bar"))) {
                                                                            case "a":
                                                                                v116 := GenS7{F0: 1}
                                                                                _ = v116
                                                                            default:
                                                                                v114 += (v114 | -127)
                                                                            }
                                                                            if (42 <= 2) {
                                                                                if (0 >= -2147483648) {
                                                                                    var v117 uint16 = (p1 << ((255) & 63))
                                                                                    _ = v117
                                                                                } else {
                                                                                    fmt.Println(v114)
                                                                                    v10 = append(v10, 520)
                                                                                    v60 += 100
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v113)
                                                                                if len(v10) >= 2 {
                                                                                    v10 = v10[1:len(v10)]
                                                                                }
                                                                                var v118 string = ("hello" + "foo bar")
                                                                                p1 += p1
                                                                                _ = v118
                                                                            }
                                                                            _ = v113
                                                                            _ = v114
                                                                            _ = v115
                                                                        } else {
                                                                            p1 = (p1 / ((p1) | 1))
                                                                            var v119 float32 = 0.1
                                                                            fmt.Println(p2)
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                p2 = int8(p1)
                                                                            case "hello":
                                                                                if true {
                                                                                    fmt.Println(v119)
                                                                                } else {
                                                                                    v120 := GenS19{F0: 0, F1: 761, F2: p1}
                                                                                    if true {
                                                                                        v60 = 42
                                                                                        p0 = (p0 % (((p0 * p0)) | 1))
                                                                                        v121 := GenS20{F0: 0, F1: (v12 | 278), F2: 9223372036854775807}
                                                                                        _ = v121
                                                                                    } else {
                                                                                        var v122 uint8 = 53
                                                                                        p0 -= uint(p1)
                                                                                        _ = len(v10)
                                                                                        _ = v122
                                                                                    }
                                                                                    v123 := GenS1{F0: -342, F1: false}
                                                                                    _ = v120
                                                                                    _ = v123
                                                                                }
                                                                            }
                                                                            if true {
                                                                                var v124 int = (fi13 ^ fi13)
                                                                                v125 := GenS11{F0: 861, F1: (v60 + (v60 - (v60 - v60)))}
                                                                                p1 += 547
                                                                                if len(v10) > 0 {
                                                                                    v10[0] = fi13
                                                                                }
                                                                                v126 := GenS17{F0: 2}
                                                                                fmt.Println(fi11)
                                                                                var v127 int8 = (v60 & v9)
                                                                                _ = v124
                                                                                _ = v125
                                                                                _ = v126
                                                                                _ = v127
                                                                            } else {
                                                                                p0 -= p0
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    v9 = -1
                                                                                case "":
                                                                                    if true {
                                                                                        v119 = ((v119 / v119) * 0.0)
                                                                                        var v128 float32 = (((float32((1.5 / float64((((v119 - ((float32((1.5 - 1.5)) - ((v119 / v119) + (v119 * v119))) * v119)) / float32(-2.25)) + v119)))) + 0.1) + v119) - v119)
                                                                                        fmt.Println(v12, p1, v60)
                                                                                        var v129 uint16 = p1
                                                                                        _ = v128
                                                                                        _ = v129
                                                                                    } else {
                                                                                        v130 := GenS19{F0: 2147483647, F1: 9223372036854775807, F2: p1}
                                                                                        _ = v130
                                                                                    }
                                                                                case "hello":
                                                                                    if (false || false) {
                                                                                        v131 := GenS21{F0: -2147483648, F1: (fi13 + fi11), F2: 4294967295}
                                                                                        if (255 == 0) {
                                                                                            switch fi13 {
                                                                                            case 32767:
                                                                                                if (62 > 233) {
                                                                                                    var v132 uint8 = 127
                                                                                                    v12 = v12
                                                                                                    _ = v132
                                                                                                }
                                                                                            default:
                                                                                                v10 = append(v10, fi13)
                                                                                            }
                                                                                        } else {
                                                                                            p1 = 836
                                                                                            v133 := GenS22{F0: 2147483647}
                                                                                            _ = v133
                                                                                        }
                                                                                        _ = v131
                                                                                    } else {
                                                                                        var v134 uint32 = 2147483648
                                                                                        var v135 uint16 = p1
                                                                                        v135 = (p1 - (p1 | p1))
                                                                                        switch fi11 {
                                                                                        case 0:
                                                                                            var v136 string = "a"
                                                                                            _ = v136
                                                                                        case -4:
                                                                                            var v137 uint64 = v12
                                                                                            _ = v137
                                                                                        case 128:
                                                                                            var v138 int8 = p2
                                                                                            switch fi13 {
                                                                                            case 32767:
                                                                                                if false {
                                                                                                    v119 += float32(3.14159)
                                                                                                    var v139 float64 = (1.5 - (-2.25 - 1.5))
                                                                                                    _ = v139
                                                                                                }
                                                                                            case 464:
                                                                                                var v140 string = ""
                                                                                                _ = v140
                                                                                            case -9223372036854775808:
                                                                                                if (1.5 < float64((1e10 / (v119 / (1.5 * v119))))) {
                                                                                                    switch 127 {
                                                                                                    case 127:
                                                                                                        var v141 int64 = 9223372036854775807
                                                                                                        _ = v141
                                                                                                    case 323:
                                                                                                        var v142 uint16 = (^(p1))
                                                                                                        _ = v142
                                                                                                    default:
                                                                                                        v10 = append(v10, (fi13 >> (((p0 % (((p0 | (p0 | uint(v12)))) | 1))) & 63)))
                                                                                                    }
                                                                                                    v135 += p1
                                                                                                }
                                                                                            default:
                                                                                                var v143 string = "\n\t"
                                                                                                _ = v143
                                                                                            }
                                                                                            var v144 uint64 = (v12 + v12)
                                                                                            if len(v10) >= 2 {
                                                                                                v10 = v10[1:len(v10)]
                                                                                            }
                                                                                            var v145 int32 = 0
                                                                                            _ = v138
                                                                                            _ = v144
                                                                                            _ = v145
                                                                                        }
                                                                                        var v146 uint = (^(p0))
                                                                                        switch "foo bar" {
                                                                                        case "":
                                                                                            var v147 bool = true
                                                                                            _ = v147
                                                                                        case "foo bar":
                                                                                            switch 42 {
                                                                                            case 42:
                                                                                                v134 = (v134 - v134)
                                                                                            case 32767:
                                                                                                var v148 float32 = -2.25
                                                                                                _ = v148
                                                                                            case 1:
                                                                                                v135 = (p1 + (^(v135)))
                                                                                            }
                                                                                        case "hello":
                                                                                            if len(v10) >= 2 {
                                                                                                v10 = v10[1:len(v10)]
                                                                                            }
                                                                                        }
                                                                                        _ = v134
                                                                                        _ = v135
                                                                                        _ = v146
                                                                                    }
                                                                                }
                                                                                v149 := GenS5{F0: 1, F1: 0, F2: (v12 >> (((p0 & p0)) & 63))}
                                                                                _ = v149
                                                                            }
                                                                            _ = v119
                                                                        }
                                                                    case 32767:
                                                                        p2 = (-(v60))
                                                                    default:
                                                                        p0 -= 4294967295
                                                                    }
                                                                }
                                                                var v150 uint16 = (p1 & p1)
                                                                if len(v10) >= 2 {
                                                                    v10 = v10[1:len(v10)]
                                                                }
                                                                _ = v150
                                                            }
                                                            v10 = append(v10, fi11)
                                                            v151 := GenS7{F0: (p2 % ((42) | 1))}
                                                            _ = v151
                                                        }
                                                    default:
                                                        if (0 >= 65535) {
                                                            fmt.Println(fi11, v12)
                                                            _ = len(v10)
                                                            v152 := GenS18{F0: 1, F1: v12}
                                                            _ = v152
                                                        } else {
                                                            v9 = (p2 << ((257) & 63))
                                                        }
                                                    }
                                                }
                                                p0 -= p0
                                                _ = v70
                                            case "foo bar":
                                                fmt.Println(p1, fi13)
                                            default:
                                                if ((p1 * p1) < (p1 * (p1 ^ 65535))) {
                                                    fmt.Println(v12, p0)
                                                }
                                            }
                                            _ = v69
                                        }
                                        fmt.Println(p2, v9, v60)
                                        fmt.Println(fi11)
                                    }
                                }
                            }
                            v12 -= (v12 - 0)
                            p0 = p0
                            if ((v12 >> ((p0) & 63)) == v12) {
                                v10 = append(v10, 832)
                                if ((1.5 <= -2.25) || (false && true)) {
                                    fmt.Println(v9, p0)
                                    var v153 int32 = 2147483647
                                    v154 := GenS5{F0: 715, F1: fi11, F2: 0}
                                    _ = v153
                                    _ = v154
                                }
                                p0 += uint(v12)
                            }
                            var v155 bool = (0 != 2147483647)
                            _ = v155
                            _ = v60
                            _ = v61
                            _ = v62
                        } else {
                            var v156 int = 32768
                            if (p1 >= p1) {
                                switch "" {
                                case "foo bar":
                                    v9 = v9
                                case "\n\t":
                                    if false {
                                        if (235 != 145) {
                                            if (-2.25 < 1.5) {
                                                p1 += p1
                                                v157 := GenS12{F0: (((("a" + "") + "") + (("hello" + ("hello" + "hello")) + "\n\t")) + "\n\t")}
                                                v9 = v9
                                                v158 := GenS1{F0: 970, F1: (!(!true))}
                                                v159 := GenS11{F0: (v156 & fi11), F1: (p2 << ((4294967295) & 63))}
                                                var v160 int32 = int32(p2)
                                                _ = v157
                                                _ = v158
                                                _ = v159
                                                _ = v160
                                            }
                                        } else {
                                            v10 = append(v10, (fi11 << ((627) & 63)))
                                        }
                                        p0 -= 255
                                    }
                                default:
                                    fmt.Println(fi11)
                                }
                                v161 := GenS9{F0: -1, F1: (("" + "hello") + "foo bar"), F2: (p1 + 1)}
                                p2 = v9
                                _ = v161
                            }
                            p1 = p1
                            _ = v156
                        }
                    }
                }
                fmt.Println(v12)
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
        _ = v10
        _ = v9
    }
    v162 := []int{127, -2147483648, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v162
    return 42
}
func genFn1() int16 {
    var v163 int64 = -998
    if false {
        if ("hello" == ("" + "foo bar")) {
            var v164 int64 = 9223372036854775807
            v163 -= (v164 % ((v163) | 1))
            v163 += 9223372036854775807
            _ = v164
        }
        v165 := []int{-9223372036854775808, -850, 127, 128}
        switch "" {
        case "":
            if (42 == 4294967295) {
                for fi166 := 0; fi166 < 17; fi166++ {
                    if true {
                        var v167 uint64 = 30
                        v163 = v163
                        _ = v167
                    }
                    _ = fi166
                }
            }
        case "hello":
            if len(v165) >= 2 {
                v165 = v165[1:len(v165)]
            }
        case "\n\t":
            v163 -= 0
        default:
            v163 += 9223372036854775807
        }
        _ = v165
    }
    _ = v163
    return 0
}

func main() {
    var v1 uint16 = 909
    var v2 float64 = (3.14159 / 1e10)
    var v3 int32 = 0
    var v4 string = "\n\t"
    var v5 float64 = -2.25
    var v6 uint = 2
    var v168 string = "hello"
    var v169 uint32 = 2
    var v170 string = "a"
    if true {
        v171 := GenS10{F0: v169, F1: -9223372036854775808, F2: (v1 / (((v1 - 867)) | 1))}
        var v172 uint = (^(v6))
        v5 -= ((0.1 + v5) - (1e10 / v5))
        _ = v171
        _ = v172
    } else {
        v173 := GenS23{F0: 0, F1: 1}
        if (-2.25 >= 0.0) {
            v168 = "\n\t"
            if (127 >= -19) {
                var v174 bool = (0 < (v1 & v1))
                for fi175 := 0; fi175 < 13; fi175++ {
                    v176 := []int{}
                    _ = fi175
                    _ = v176
                }
                var v177 uint = (v6 / (((v6 + v6)) | 1))
                _ = v174
                _ = v177
            }
            for fi178 := 0; fi178 < 15; fi178++ {
                if false {
                    for fi179 := 0; fi179 < 1; fi179++ {
                        var v180 int8 = -1
                        switch (v4 + v170) {
                        case "foo bar":
                            if false {
                                fmt.Println(v1, v3, v1)
                                v4 = v168
                                v168 = (v168 + (v4 + "a"))
                                v181 := GenS24{F0: 34, F1: genFn1(), F2: fi179}
                                switch v168 {
                                case "hello":
                                    v182 := []int{fi179, (fi178 - 4294967295), fi179, int(v6)}
                                    _ = v182
                                default:
                                    var v183 bool = true
                                    _ = v183
                                }
                                if (1 == 32767) {
                                    v184 := GenS25{F0: (((-9223372036854775808 == -1) || false) && false), F1: v4, F2: 1}
                                    v185 := []int{-2147483648, -9223372036854775808}
                                    if false {
                                        if ((-(fi179)) != fi178) {
                                            v3 = int32(fi179)
                                        }
                                        fmt.Println(v180, v3, v169)
                                    } else {
                                        var v186 uint32 = (v169 % ((2147483647) | 1))
                                        v187 := GenS10{F0: (v186 * v169), F1: -9223372036854775808, F2: 0}
                                        _ = v186
                                        _ = v187
                                    }
                                    _ = v184
                                    _ = v185
                                }
                                var v188 float64 = v2
                                _ = v181
                                _ = v188
                            }
                        case "hello":
                            var v189 int32 = (v3 / ((v3) | 1))
                            _ = v189
                        case "a":
                            var v190 int64 = 0
                            _ = v190
                        }
                        v168 = "\n\t"
                        v170 = v168
                        var v191 uint = v6
                        _ = fi179
                        _ = v180
                        _ = v191
                    }
                } else {
                    v192 := GenS12{F0: "foo bar"}
                    v1 = uint16(v169)
                    _ = v192
                }
                for fi193 := 0; fi193 < 19; fi193++ {
                    var v194 uint = (^(v6))
                    _ = fi193
                    _ = v194
                }
                _ = fi178
            }
        } else {
            v195 := []int{-2147483648, 42, 127}
            if ((v168 + "a") <= (v170 + (("hello" + v4) + v170))) {
                var v196 uint8 = 179
                for fi197 := 0; fi197 < 9; fi197++ {
                    switch 646 {
                    case -2147483648:
                        v198 := GenS6{F0: v1, F1: v3}
                        _ = v198
                    case 2147483647:
                        if (1 == 42) {
                            switch fi197 {
                            case 9223372036854775807:
                                fmt.Println(v196, v196)
                            case 32768:
                                v195 = append(v195, (-(fi197)))
                            case 567:
                                var v199 uint32 = v169
                                _ = v199
                            }
                            v200 := GenS23{F0: -368, F1: 42}
                            _ = v200
                        } else {
                            var v201 float64 = v2
                            if true {
                                v3 -= 2147483647
                                v170 = "hello"
                                var v202 int = (fi197 | fi197)
                                _ = v202
                            }
                            if (v170 > "a") {
                                v168 = ""
                                v2 += v2
                                v201 += v201
                            } else {
                                var v203 string = "hello"
                                v204 := GenS26{F0: (fi197 / (((fi197 / ((617) | 1))) | 1))}
                                _ = v203
                                _ = v204
                            }
                            _ = v201
                        }
                    }
                    _ = fi197
                }
                v1 = (^(v1))
                _ = v196
            }
            _ = v195
        }
        v2 = (v5 / (float64(3.14159) + v5))
        _ = v173
    }
    v6 -= 2
    for fi205 := 0; fi205 < 1; fi205++ {
        if (genFn1() == 42) {
            v5 -= v5
            fmt.Println(v3)
        }
        _ = fi205
    }
    var v206 uint64 = 82
    v206 = v206
    if ((false && false) && (100 >= int8(v3))) {
        for fi207 := 0; fi207 < 13; fi207++ {
            var v208 int = (fi207 * (fi207 << (((v6 << (((v6 - (v6 + 18446744073709551615))) & 63))) & 63)))
            if (42 == 0) {
                switch 0 {
                case 0:
                    v3 = (v3 % (((v3 << ((v6) & 63))) | 1))
                case 2147483647:
                    var v209 int64 = int64(v6)
                    _ = v209
                }
            } else {
                v168 = v168
                var v210 string = v4
                var v211 float64 = 1e10
                _ = v210
                _ = v211
            }
            var v212 uint64 = 18446744073709551615
            _ = fi207
            _ = v208
            _ = v212
        }
        v213 := GenS27{F0: 0, F1: v3}
        for fi214 := 0; fi214 < 6; fi214++ {
            switch "foo bar" {
            case "foo bar":
                v215 := GenS9{F0: 1000, F1: ((v4 + "a") + v170), F2: 300}
                _ = v215
            default:
                v216 := GenS28{F0: (v6 / ((65535) | 1)), F1: v170, F2: 726}
                _ = v216
            }
            var v217 uint = v6
            _ = fi214
            _ = v217
        }
        _ = v213
    } else {
        fmt.Println(v168, v5)
        v218 := []int{1, 42, -1, -928}
        _ = v218
    }
    fmt.Println(v170)
    fmt.Println(v3, v206, v1)
    v219 := GenS10{F0: v169, F1: -893, F2: 997}
    var v220 int64 = 1
    v221 := []int{4294967295, 32767, 32767}
    if len(v221) > 0 {
        v221[0] = 900
    }
    if len(v221) > 0 {
        v221[0] = -132
    }
    v221 = append(v221, 495)
    for fi222 := 0; fi222 < 12; fi222++ {
        _ = len(v221)
        _ = fi222
    }
    fmt.Println("v1:", v1)
    fmt.Println("v168:", v168)
    fmt.Println("v169:", v169)
    fmt.Println("v170:", v170)
    fmt.Println("v2:", v2)
    fmt.Println("v206:", v206)
    fmt.Println("v220:", v220)
    fmt.Println("slicelen v221:", len(v221), v221)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v168
    _ = v169
    _ = v170
    _ = v2
    _ = v206
    _ = v219
    _ = v220
    _ = v221
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
