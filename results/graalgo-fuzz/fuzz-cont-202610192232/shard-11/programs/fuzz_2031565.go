package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int; F1 uint64; F2 int16 }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint64; F1 uint64; F2 string }
type GenS10 struct { F0 bool; F1 bool }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint; F1 int }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int; F1 int; F2 uint8 }
type GenS15 struct { F0 uint; F1 string }
type GenS16 struct { F0 uint }

func main() {
    var v1 int16 = 0
    var v2 float64 = 0.0
    var v3 uint32 = 2147483648
    v4 := []int{-1, 128, 32768}
    var v5 float32 = (1.5 * 3.14159)
    if len(v4) > 0 {
        v4[0] = 128
    }
    var v6 string = "\n\t"
    v1 += (v1 | v1)
    if (!((false || true) || true)) {
        v4 = append(v4, 0)
        var v7 bool = true
        _ = v7
    } else {
        v1 = v1
        var v8 float64 = ((1.5 * (0.1 + (3.14159 - float64(v5)))) / v2)
        switch (v6 + ("\n\t" + "")) {
        case "hello":
            switch -312 {
            case 32768:
                var v9 int16 = v1
                _ = v9
            }
        case "\n\t":
            var v10 float32 = 3.14159
            _ = v10
        }
        _ = v8
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v3 += 2147483647
        _ = fi11
    }
    var v12 uint = 18446744073709551615
    var v13 uint8 = 1
    if (((!true) && (uint32(v12) < (v3 - uint32(v12)))) || ((^(v1)) == (v1 ^ (v1 / (((v1 / (((^(v1))) | 1))) | 1))))) {
        var v14 int8 = 42
        for fi15 := 0; fi15 < 10; fi15++ {
            if (0 < uint16(v13)) {
                var v16 uint8 = v13
                _ = v16
            }
            if (89 < (^(v3))) {
                for fi17 := 0; fi17 < 8; fi17++ {
                    _ = len(v4)
                    v18 := GenS1{F0: 18446744073709551615, F1: (v13 >> (((^(v12))) & 63)), F2: 858}
                    if true {
                        var v19 uint8 = (v13 * (v13 | 128))
                        _ = v19
                    } else {
                        switch (fi17 - (fi17 * 127)) {
                        case -2147483648:
                            switch v6 {
                            case "hello":
                                v20 := GenS2{F0: ((v6 + (v6 + "a")) + "\n\t")}
                                _ = v20
                            case "foo bar":
                                var v21 int16 = 0
                                _ = v21
                            }
                        }
                    }
                    _ = fi17
                    _ = v18
                }
                switch (fi15 & (fi15 - fi15)) {
                case -568:
                    var v22 string = "\n\t"
                    _ = v22
                case 9223372036854775807:
                    for fi23 := 0; fi23 < 15; fi23++ {
                        var v24 int8 = (v14 | v14)
                        var v25 float32 = v5
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                case 541:
                    var v26 int32 = 294
                    _ = v26
                default:
                    fmt.Println(v5)
                }
                if (float32(0.0) != v5) {
                    v27 := GenS2{F0: "hello"}
                    var v28 string = (v6 + v6)
                    _ = v27
                    _ = v28
                }
            } else {
                for fi29 := 0; fi29 < 19; fi29++ {
                    _ = len(v4)
                    v4 = append(v4, fi29)
                    _ = fi29
                }
                fmt.Println(v12)
            }
            _ = fi15
        }
        var v30 uint64 = 4294967295
        _ = v14
        _ = v30
    }
    var v31 int = 780
    v31 -= v31
    var v32 uint64 = 666
    for fi33 := 0; fi33 < 20; fi33++ {
        if false {
            switch fi33 {
            case 1:
                var v34 float64 = -2.25
                _ = v34
            case 9223372036854775807:
                for fi35 := 0; fi35 < 1; fi35++ {
                    v4 = append(v4, 1)
                    _ = fi35
                }
            default:
                for fi36 := 0; fi36 < 11; fi36++ {
                    v32 -= uint64(v3)
                    fmt.Println(v3)
                    fmt.Println(v31)
                    if false {
                        var v37 uint16 = 0
                        switch "\n\t" {
                        case "hello":
                            var v38 int32 = 1
                            _ = v38
                        case "\n\t":
                            v39 := GenS1{F0: uint64(v13), F1: v13, F2: uint64(v13)}
                            _ = v39
                        case "":
                            switch "a" {
                            case "\n\t":
                                fmt.Println(v12, v6)
                            case "foo bar":
                                v40 := GenS1{F0: v32, F1: 127, F2: (v32 - (v32 >> ((v12) & 63)))}
                                _ = v40
                            default:
                                if false {
                                    switch (fi33 << ((v12) & 63)) {
                                    case 127:
                                        if true {
                                            switch fi33 {
                                            case 32768:
                                                v37 = uint16(v12)
                                            case 0:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            default:
                                                var v41 bool = false
                                                v12 = 192
                                                _ = v41
                                            }
                                            _ = len(v4)
                                            switch (-(v31)) {
                                            case -905:
                                                v42 := GenS2{F0: (v6 + ((v6 + ("hello" + (v6 + (v6 + (v6 + v6))))) + (((((v6 + v6) + v6) + ("a" + v6)) + (("a" + (((v6 + "a") + v6) + v6)) + v6)) + (v6 + "foo bar"))))}
                                                _ = v42
                                            }
                                        } else {
                                            v6 = ((v6 + "") + ((v6 + v6) + v6))
                                            if len(v4) > 0 {
                                                v4[0] = fi36
                                            }
                                            fmt.Println(v12, fi36, v3)
                                            v4 = append(v4, fi33)
                                            var v43 bool = (("\n\t" + "\n\t") > "foo bar")
                                            _ = v43
                                        }
                                    case 1:
                                        if len(v4) > 0 {
                                            v4[0] = fi33
                                        }
                                    default:
                                        if (!false) {
                                            if (-1 != 42) {
                                                var v44 bool = (!((v37 | (v37 * (^(v37)))) >= v37))
                                                v45 := GenS3{F0: 18446744073709551615, F1: (fi33 * v31)}
                                                _ = v44
                                                _ = v45
                                            } else {
                                                if (true && ((v3 - (v3 >> ((42) & 63))) <= uint32(v13))) {
                                                    if ((v37 < (v37 - uint16(v3))) && false) {
                                                        v32 = (v32 / (((v32 + v32)) | 1))
                                                        v46 := GenS1{F0: (v32 * v32), F1: (v13 - v13), F2: v32}
                                                        fmt.Println(fi33)
                                                        _ = v46
                                                    } else {
                                                        if (float64((float32(((1.5 / (((float64(0.1) + v2) + float64(0.0)) / v2)) - (3.14159 / v2))) - ((-2.25 * 0.0) - v5))) <= v2) {
                                                            v47 := GenS4{F0: 0, F1: (fi36 + (^(fi36)))}
                                                            if false {
                                                                switch ((("foo bar" + v6) + "foo bar") + v6) {
                                                                case "\n\t":
                                                                    if ((v2 + (float64((((((v5 - v5) + v5) + 3.14159) + v5) * float32((1.5 + (v2 - (v2 + v2)))))) - (v2 + 1.5))) < 1e10) {
                                                                        if len(v4) > 0 {
                                                                            v4[0] = v31
                                                                        }
                                                                        var v48 uint64 = 18446744073709551615
                                                                        var v49 int32 = -1
                                                                        _ = v48
                                                                        _ = v49
                                                                    }
                                                                case "foo bar":
                                                                    v50 := GenS5{F0: (^(fi36))}
                                                                    _ = v50
                                                                case "a":
                                                                    fmt.Println(fi33, v6)
                                                                default:
                                                                    fmt.Println(v32, fi33)
                                                                }
                                                                var v51 int8 = 100
                                                                if (!true) {
                                                                    v1 -= -552
                                                                    v51 = (v51 / ((4) | 1))
                                                                    v52 := GenS6{F0: (fi36 + v31), F1: 0, F2: (v1 << ((v12) & 63))}
                                                                    _ = v52
                                                                }
                                                                _ = v51
                                                            } else {
                                                                var v53 uint = 1
                                                                var v54 float32 = v5
                                                                switch "foo bar" {
                                                                case "a":
                                                                    switch (fi33 + fi36) {
                                                                    case -1:
                                                                        var v55 float32 = ((v5 + v54) / v54)
                                                                        _ = v55
                                                                    case 293:
                                                                        v56 := GenS7{F0: false, F1: fi33}
                                                                        _ = v56
                                                                    case -2147483648:
                                                                        var v57 string = v6
                                                                        _ = v57
                                                                    default:
                                                                        if true {
                                                                            if true {
                                                                                if true {
                                                                                    v54 += v54
                                                                                    var v58 uint64 = (^(v32))
                                                                                    var v59 uint8 = (v13 / ((v13) | 1))
                                                                                    _ = v58
                                                                                    _ = v59
                                                                                } else {
                                                                                    switch (-(v31)) {
                                                                                    case 2147483647:
                                                                                        if (true || false) {
                                                                                            var v60 uint8 = v13
                                                                                            if false {
                                                                                                if false {
                                                                                                    fmt.Println(v5, v53, fi36)
                                                                                                    var v61 float32 = v5
                                                                                                    fmt.Println(fi33, v2, v54)
                                                                                                    _ = v61
                                                                                                } else {
                                                                                                    if true {
                                                                                                        switch 4294967295 {
                                                                                                        case 297:
                                                                                                            v3 -= 842
                                                                                                        case 913:
                                                                                                            if false {
                                                                                                                _ = len(v4)
                                                                                                            }
                                                                                                        default:
                                                                                                            var v62 int8 = int8(v37)
                                                                                                            _ = v62
                                                                                                        }
                                                                                                        v31 += int(v1)
                                                                                                        var v63 int8 = 81
                                                                                                        _ = v63
                                                                                                    }
                                                                                                    if len(v4) > 0 {
                                                                                                        v4[0] = v31
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            _ = v60
                                                                                        } else {
                                                                                            v53 -= v53
                                                                                            if ((822 <= v31) && true) {
                                                                                                var v64 uint = 18446744073709551615
                                                                                                _ = v64
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        if false {
                                                                                            fmt.Println(v31, v13)
                                                                                            if false {
                                                                                                switch (v31 ^ v31) {
                                                                                                case 9223372036854775807:
                                                                                                    fmt.Println(fi36, v53)
                                                                                                default:
                                                                                                    if true {
                                                                                                        if len(v4) > 0 {
                                                                                                            v4[0] = 394
                                                                                                        }
                                                                                                        v65 := GenS3{F0: (v53 - (^(v12))), F1: fi33}
                                                                                                        _ = v65
                                                                                                    } else {
                                                                                                        v2 = v2
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                var v66 uint32 = v3
                                                                                                if false {
                                                                                                    var v67 int32 = 0
                                                                                                    _ = v67
                                                                                                } else {
                                                                                                    var v68 int16 = (v1 + (v1 - (v1 + -32768)))
                                                                                                    _ = v68
                                                                                                }
                                                                                                var v69 uint32 = (v66 / ((uint32(v13)) | 1))
                                                                                                _ = v66
                                                                                                _ = v69
                                                                                            }
                                                                                            v70 := GenS8{F0: v32}
                                                                                            if false {
                                                                                                v71 := GenS2{F0: v6}
                                                                                                _ = v71
                                                                                            } else {
                                                                                                v13 = v13
                                                                                                var v72 int = 9223372036854775807
                                                                                                _ = v72
                                                                                            }
                                                                                            _ = v70
                                                                                        } else {
                                                                                            var v73 uint32 = v3
                                                                                            v74 := GenS9{F0: 4294967295, F1: (v32 & (v32 << ((2) & 63))), F2: v6}
                                                                                            _ = v73
                                                                                            _ = v74
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v31, v37, v5)
                                                                                    var v75 uint64 = (v32 | v32)
                                                                                    var v76 float32 = (v54 * v5)
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = v31
                                                                                    }
                                                                                    var v77 int64 = 42
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                    _ = v77
                                                                                }
                                                                                var v78 uint8 = uint8(v3)
                                                                                if false {
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = 127
                                                                                    }
                                                                                    v79 := GenS10{F0: true, F1: false}
                                                                                    v31 = -372
                                                                                    v1 -= (v1 & -266)
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = -508
                                                                                    }
                                                                                    v80 := GenS9{F0: (v32 + v32), F1: v32, F2: ("hello" + v6)}
                                                                                    if (v53 > v12) {
                                                                                        var v81 uint64 = (^(v32))
                                                                                        v82 := GenS4{F0: 1, F1: 127}
                                                                                        if ((float32((-2.25 - -2.25)) - 0.0) >= (1.5 / v54)) {
                                                                                            fmt.Println(v5, v6, v12)
                                                                                            v83 := GenS2{F0: "a"}
                                                                                            v13 = 42
                                                                                            _ = v83
                                                                                        } else {
                                                                                            var v84 uint8 = uint8(v32)
                                                                                            _ = v84
                                                                                        }
                                                                                        _ = v81
                                                                                        _ = v82
                                                                                    } else {
                                                                                        if false {
                                                                                            v85 := GenS11{F0: 42}
                                                                                            if true {
                                                                                                if (42 > -2147483648) {
                                                                                                    v31 = fi33
                                                                                                } else {
                                                                                                    v3 += v3
                                                                                                    v86 := GenS12{F0: v53, F1: (fi36 ^ (fi33 | (fi36 & (fi36 ^ 259))))}
                                                                                                    if false {
                                                                                                        fmt.Println(v13)
                                                                                                        v87 := GenS13{F0: 100}
                                                                                                        _ = v87
                                                                                                    } else {
                                                                                                        if len(v4) > 0 {
                                                                                                            v4[0] = (fi33 + -9223372036854775808)
                                                                                                        }
                                                                                                        var v88 uint16 = (v37 - uint16(v3))
                                                                                                        if (!false) {
                                                                                                            v88 += v88
                                                                                                            switch (v6 + ((v6 + "foo bar") + (v6 + v6))) {
                                                                                                            case "":
                                                                                                                if len(v4) >= 2 {
                                                                                                                    v4 = v4[1:len(v4)]
                                                                                                                }
                                                                                                            case "a":
                                                                                                                v89 := GenS9{F0: 478, F1: uint64(v3), F2: v6}
                                                                                                                _ = v89
                                                                                                            default:
                                                                                                                fmt.Println(v78, fi33)
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v90 int16 = 32767
                                                                                                            _ = v90
                                                                                                        }
                                                                                                        _ = v88
                                                                                                    }
                                                                                                    v91 := GenS14{F0: fi36, F1: (^(v31)), F2: v78}
                                                                                                    _ = v86
                                                                                                    _ = v91
                                                                                                }
                                                                                            } else {
                                                                                                var v92 uint = 627
                                                                                                var v93 int64 = 42
                                                                                                switch ("\n\t" + (v6 + v6)) {
                                                                                                case "":
                                                                                                    v2 -= ((1.5 + 1.5) + v2)
                                                                                                case "hello":
                                                                                                    if (true && ((v12 % ((uint(v37)) | 1)) <= v53)) {
                                                                                                        switch (fi36 / ((fi33) | 1)) {
                                                                                                        case -703:
                                                                                                            v31 += 32768
                                                                                                        }
                                                                                                    } else {
                                                                                                        v94 := GenS14{F0: (fi33 & 32767), F1: 127, F2: (v13 ^ (v78 % ((v13) | 1)))}
                                                                                                        _ = v94
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    switch fi33 {
                                                                                                    case 127:
                                                                                                        v78 += (v13 | (v13 ^ v13))
                                                                                                        var v95 int32 = int32(v3)
                                                                                                        _ = v95
                                                                                                    case 693:
                                                                                                        fmt.Println(v12, v32, v13)
                                                                                                    case -1:
                                                                                                        var v96 uint = v53
                                                                                                        _ = v96
                                                                                                    default:
                                                                                                        v31 = int(v12)
                                                                                                    }
                                                                                                }
                                                                                                _ = v92
                                                                                                _ = v93
                                                                                            }
                                                                                            _ = v85
                                                                                        } else {
                                                                                            v97 := GenS7{F0: (!false), F1: fi33}
                                                                                            _ = v97
                                                                                        }
                                                                                        _ = len(v4)
                                                                                        v12 = v53
                                                                                    }
                                                                                    v1 = (v1 >> ((4294967295) & 63))
                                                                                    var v98 bool = false
                                                                                    _ = v79
                                                                                    _ = v80
                                                                                    _ = v98
                                                                                } else {
                                                                                    v32 += (v32 * uint64(v12))
                                                                                }
                                                                                _ = v78
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                _ = v53
                                                                _ = v54
                                                            }
                                                            _ = v47
                                                        } else {
                                                            v99 := GenS10{F0: ((v2 * ((v2 - v2) + (((v2 * 1.5) * v2) * ((v2 + (v2 + 3.14159)) + float64(0.0))))) >= v2), F1: (false && true)}
                                                            switch (v6 + "foo bar") {
                                                            case "hello":
                                                                v100 := GenS6{F0: fi33, F1: uint64(v3), F2: (v1 & v1)}
                                                                _ = v100
                                                            }
                                                            if true {
                                                                fmt.Println(v32)
                                                            }
                                                            _ = v99
                                                        }
                                                        v5 += (1.5 + float32((v2 / v2)))
                                                        var v101 uint = v12
                                                        _ = v101
                                                    }
                                                    v37 += uint16(v3)
                                                    var v102 string = (v6 + v6)
                                                    _ = v102
                                                }
                                            }
                                            var v103 uint16 = (v37 | (v37 * v37))
                                            v104 := GenS15{F0: (v12 | v12), F1: v6}
                                            v37 = uint16(v32)
                                            if false {
                                                if (true || (!false)) {
                                                    var v105 uint = v12
                                                    if ((^(v103)) != v37) {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        fmt.Println(v31, v1)
                                                    }
                                                    switch (v6 + v6) {
                                                    case "":
                                                        var v106 string = (v6 + "foo bar")
                                                        _ = v106
                                                    }
                                                    _ = v105
                                                } else {
                                                    var v107 int16 = (-(v1))
                                                    v108 := GenS15{F0: (^(v12)), F1: v6}
                                                    v109 := GenS16{F0: (v12 << ((v12) & 63))}
                                                    v110 := GenS10{F0: ((v12 ^ v12) < v12), F1: (((^(v32)) == (v32 ^ v32)) || true)}
                                                    if true {
                                                        fmt.Println(v13)
                                                        var v111 int = fi33
                                                        _ = v111
                                                    }
                                                    _ = v107
                                                    _ = v108
                                                    _ = v109
                                                    _ = v110
                                                }
                                            } else {
                                                if false {
                                                    v4 = append(v4, (-(fi33)))
                                                    v12 = (v12 / (((v12 % ((v12) | 1))) | 1))
                                                } else {
                                                    var v112 uint32 = v3
                                                    _ = v112
                                                }
                                                var v113 bool = false
                                                v4 = append(v4, fi33)
                                                _ = v113
                                            }
                                            var v114 float32 = -2.25
                                            switch v31 {
                                            case 9223372036854775807:
                                                v3 = (v3 ^ 4294967295)
                                            case 813:
                                                v5 = v114
                                                if (42 <= (^(v37))) {
                                                    var v115 int = -897
                                                    v116 := GenS16{F0: 484}
                                                    v114 = float32(1.5)
                                                    v1 -= 1000
                                                    _ = v115
                                                    _ = v116
                                                }
                                            default:
                                                if (!false) {
                                                    v6 = v6
                                                    fmt.Println(v2)
                                                    v31 += (fi33 & 56)
                                                }
                                            }
                                            if ((v12 == (^(v12))) && (!((fi36 ^ 261) <= v31))) {
                                                if false {
                                                    var v117 int16 = v1
                                                    var v118 int64 = 1
                                                    _ = v117
                                                    _ = v118
                                                } else {
                                                    v6 = ("a" + v6)
                                                    v119 := GenS14{F0: 128, F1: fi36, F2: v13}
                                                    _ = v119
                                                }
                                                if true {
                                                    _ = len(v4)
                                                    fmt.Println(v103, v1)
                                                    v37 = v37
                                                } else {
                                                    v3 = (^(v3))
                                                }
                                            }
                                            v5 = v5
                                            v31 = 42
                                            v37 = (v37 + (v103 << ((uint(v37)) & 63)))
                                            _ = v103
                                            _ = v104
                                            _ = v114
                                        } else {
                                            v120 := GenS6{F0: fi33, F1: (^(v32)), F2: -883}
                                            var v121 int8 = -128
                                            _ = v120
                                            _ = v121
                                        }
                                    }
                                    v1 = (v1 << ((2) & 63))
                                }
                            }
                        }
                        v32 += v32
                        _ = v37
                    }
                    _ = fi36
                }
            }
            for fi122 := 0; fi122 < 10; fi122++ {
                var v123 int16 = -1
                var v124 string = v6
                if false {
                    if ((fi122 << ((42) & 63)) <= int(v32)) {
                        var v125 uint32 = v3
                        var v126 int8 = -1
                        var v127 uint16 = 617
                        _ = v125
                        _ = v126
                        _ = v127
                    }
                    var v128 uint32 = v3
                    v123 = (v1 ^ v123)
                    _ = v128
                } else {
                    v129 := GenS8{F0: (v32 << (((v12 / ((v12) | 1))) & 63))}
                    v12 += 304
                    _ = v129
                }
                _ = fi122
                _ = v123
                _ = v124
            }
            v6 = "a"
        } else {
            v3 -= 0
            v12 = v12
            var v130 uint64 = 239
            _ = v130
        }
        _ = fi33
    }
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
