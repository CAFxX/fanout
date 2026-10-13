package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int64; F2 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 int; F2 bool }
type GenS6 struct { F0 bool; F1 int64; F2 bool }
type GenS7 struct { F0 string; F1 uint64; F2 uint8 }
type GenS8 struct { F0 uint16; F1 int16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16; F1 bool; F2 uint }
type GenS11 struct { F0 int16; F1 uint; F2 uint64 }
type GenS12 struct { F0 uint32; F1 uint8 }
type GenS13 struct { F0 uint64 }

func main() {
    var v1 bool = (!true)
    v2 := []uint32{314, 2147483648}
    var v3 int64 = 9223372036854775807
    var v4 uint8 = 42
    fmt.Println(v1)
    var v5 int16 = -747
    v2 = append(v2, 2)
    var v6 uint16 = 255
    if ((!false) || (!false)) {
        v6 = 42
        switch 1 {
        case -362:
            switch "foo bar" {
            case "":
                switch 42 {
                case 32768:
                    if (v1 || ((v5 + -941) >= v5)) {
                        var v7 int32 = 42
                        _ = v7
                    } else {
                        var v8 uint32 = 2
                        v6 += 0
                        v9 := GenS1{F0: (v4 - v4), F1: (v6 << ((255) & 63))}
                        _ = v8
                        _ = v9
                    }
                case 9223372036854775807:
                    var v10 uint = 369
                    _ = v10
                case -590:
                    switch "" {
                    case "":
                        switch ("\n\t" + ((("foo bar" + (("\n\t" + "hello") + "foo bar")) + "\n\t") + ("a" + ("foo bar" + (("" + "hello") + "foo bar"))))) {
                        case "hello":
                            v11 := GenS1{F0: 42, F1: v6}
                            _ = v11
                        }
                    case "foo bar":
                        var v12 uint64 = 559
                        _ = v12
                    }
                }
            case "hello":
                v13 := GenS2{F0: 496}
                _ = v13
            case "foo bar":
                fmt.Println(v3)
            }
        default:
            if (v1 && (-32 == 0)) {
                v1 = (1e10 >= ((1e10 / 1e10) + (float32(float64(1e10)) + 0.0)))
                v14 := GenS2{F0: 0}
                _ = v14
            } else {
                fmt.Println(v4, v5, v5)
            }
        }
        switch "hello" {
        case "\n\t":
            var v15 uint64 = 68
            _ = v15
        case "foo bar":
            var v16 float64 = ((float64(1.5) - float64(0.0)) * 1.5)
            _ = v16
        default:
            v17 := GenS2{F0: 4294967295}
            _ = v17
        }
    } else {
        if (v1 || v1) {
            if false {
                switch 32767 {
                case -1:
                    if v1 {
                        v18 := GenS2{F0: 962}
                        switch 0 {
                        case 419:
                            v19 := GenS3{F0: (-(v3)), F1: (v3 + v3), F2: int16(v6)}
                            _ = v19
                        case 823:
                            for fi20 := 0; fi20 < 9; fi20++ {
                                var v21 uint = 4294967295
                                for fi22 := 0; fi22 < 5; fi22++ {
                                    fmt.Println(v21, v21, v1)
                                    v1 = (42 <= -100)
                                    _ = len(v2)
                                    _ = fi22
                                }
                                _ = fi20
                                _ = v21
                            }
                        case 349:
                            fmt.Println(v3, v1)
                        default:
                            switch "foo bar" {
                            case "hello":
                                v5 = (^(v5))
                            default:
                                for fi23 := 0; fi23 < 11; fi23++ {
                                    _ = len(v2)
                                    _ = fi23
                                }
                            }
                        }
                        var v24 uint64 = 18446744073709551615
                        _ = v18
                        _ = v24
                    }
                case -9223372036854775808:
                    v4 = 128
                case 128:
                    v4 = (v4 ^ (^(v4)))
                default:
                    switch 2147483647 {
                    case -672:
                        var v25 bool = v1
                        _ = v25
                    case 1:
                        var v26 uint = 255
                        _ = v26
                    case 32768:
                        var v27 float64 = 3.14159
                        _ = v27
                    }
                }
            }
        } else {
            var v28 int32 = -2147483648
            _ = v28
        }
        var v29 uint8 = v4
        _ = v29
    }
    fmt.Println(v3, v4, v3)
    var v30 int64 = 377
    v4 -= v4
    if v1 {
        v3 += v30
        var v31 bool = (1e10 > 3.14159)
        v2 = append(v2, 490)
        _ = v31
    } else {
        var v32 int8 = 100
        _ = v32
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        fmt.Println(v6, v6)
        _ = fi33
    }
    v34 := GenS4{F0: 42}
    for fi35 := 0; fi35 < 10; fi35++ {
        v36 := GenS5{F0: (v3 - (v3 >> ((903) & 63))), F1: (fi35 - 0), F2: false}
        _ = fi35
        _ = v36
    }
    fmt.Println(v3)
    for fi37 := 0; fi37 < 16; fi37++ {
        for fi38 := 0; fi38 < 6; fi38++ {
            _ = len(v2)
            fmt.Println(fi37)
            _ = fi38
        }
        switch (("hello" + "hello") + "foo bar") {
        case "a":
            if true {
                fmt.Println(fi37, v1)
                for fi39 := 0; fi39 < 9; fi39++ {
                    _ = len(v2)
                    fmt.Println(v5)
                    _ = fi39
                }
            }
        case "foo bar":
            for fi40 := 0; fi40 < 9; fi40++ {
                v41 := GenS6{F0: ((v6 / (((v6 - (^(v6)))) | 1)) <= (^(v6))), F1: v30, F2: (!v1)}
                var v42 uint = 65535
                v3 += -1
                _ = fi40
                _ = v41
                _ = v42
            }
        default:
            for fi43 := 0; fi43 < 3; fi43++ {
                if true {
                    v44 := GenS2{F0: 0}
                    var v45 uint32 = 545
                    if (!(1.5 == 3.14159)) {
                        if (int16(v45) > v5) {
                            v4 = 128
                            var v46 string = "a"
                            _ = v46
                        }
                    }
                    var v47 uint32 = v45
                    v48 := GenS3{F0: 661, F1: (v30 ^ v30), F2: v5}
                    _ = v44
                    _ = v45
                    _ = v47
                    _ = v48
                } else {
                    _ = len(v2)
                    v5 += v5
                    if true {
                        v49 := GenS5{F0: 0, F1: int(v5), F2: (0.0 != 3.14159)}
                        if (v4 < (v4 - (^(v4)))) {
                            _ = len(v2)
                            v50 := GenS2{F0: uint32(v4)}
                            v51 := GenS2{F0: 2}
                            _ = v50
                            _ = v51
                        }
                        v52 := GenS2{F0: 2147483648}
                        _ = v49
                        _ = v52
                    } else {
                        fmt.Println(fi37, v5, v30)
                    }
                }
                switch (("a" + "hello") + "a") {
                case "\n\t":
                    v2 = append(v2, 0)
                }
                v53 := GenS7{F0: "a", F1: 2, F2: v4}
                _ = fi43
                _ = v53
            }
        }
        v6 -= v6
        _ = fi37
    }
    if (v6 <= v6) {
        switch -1 {
        case 48:
            var v54 string = ("a" + "\n\t")
            _ = v54
        case -544:
            if false {
                v1 = false
                var v55 bool = (!(v3 != (v30 << ((4294967295) & 63))))
                v4 -= v4
                _ = v55
            } else {
                var v56 float32 = 1e10
                if v1 {
                    switch "foo bar" {
                    case "foo bar":
                        switch 148 {
                        case 2147483647:
                            var v57 uint32 = 65535
                            _ = v57
                        case -221:
                            v58 := GenS8{F0: (v6 << ((593) & 63)), F1: 32767}
                            _ = v58
                        default:
                            var v59 uint64 = 0
                            _ = v59
                        }
                    case "\n\t":
                        _ = len(v2)
                    case "a":
                        if ("hello" <= "\n\t") {
                            v6 = v6
                        } else {
                            var v60 int16 = -32768
                            v30 -= v30
                            var v61 int32 = -1
                            _ = v60
                            _ = v61
                        }
                    }
                    var v62 uint64 = 1
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v62
                } else {
                    v1 = false
                }
                v3 += (v30 << ((0) & 63))
                _ = v56
            }
        case 327:
            for fi63 := 0; fi63 < 4; fi63++ {
                if v1 {
                    var v64 uint8 = v4
                    var v65 bool = false
                    _ = v64
                    _ = v65
                } else {
                    v6 = (^(v6))
                    for fi66 := 0; fi66 < 13; fi66++ {
                        v67 := GenS2{F0: 1}
                        v6 = v6
                        _ = fi66
                        _ = v67
                    }
                    for fi68 := 0; fi68 < 18; fi68++ {
                        var v69 string = (("\n\t" + "hello") + "")
                        v4 = (^(v4))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        switch 9223372036854775807 {
                        case 4294967295:
                            var v70 int8 = -128
                            _ = v70
                        case -898:
                            v30 += v3
                            fmt.Println(v4, v30, v30)
                            if ((fi63 - (fi68 << ((58) & 63))) > -528) {
                                var v71 int64 = int64(v6)
                                if v1 {
                                    var v72 uint8 = (v4 + 127)
                                    _ = v72
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = len(v2)
                                }
                                _ = v71
                            }
                            v73 := GenS5{F0: (v3 * (v30 << ((974) & 63))), F1: 1, F2: (!v1)}
                            _ = v73
                        }
                        var v74 float32 = 0.0
                        v2 = append(v2, 4294967295)
                        _ = fi68
                        _ = v69
                        _ = v74
                    }
                }
                v75 := GenS5{F0: (v3 / ((v3) | 1)), F1: 32768, F2: v1}
                _ = fi63
                _ = v75
            }
        default:
            var v76 int = 272
            _ = v76
        }
    } else {
        if v1 {
            v6 = (v6 & (v6 + (v6 - 468)))
            var v77 uint = 65535
            v78 := GenS3{F0: -9223372036854775808, F1: v30, F2: v5}
            _ = v77
            _ = v78
        } else {
            var v79 int8 = 1
            _ = v79
        }
        _ = len(v2)
        if v1 {
            switch -201 {
            case 704:
                v30 = 42
            case 127:
                var v80 float32 = -2.25
                _ = v80
            case 9223372036854775807:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                var v81 string = "foo bar"
                _ = v81
            }
            if false {
                if v1 {
                    v4 = v4
                    for fi82 := 0; fi82 < 17; fi82++ {
                        switch fi82 {
                        case 127:
                            var v83 uint8 = (v4 + (v4 << ((0) & 63)))
                            _ = v83
                        case 4294967295:
                            v84 := GenS4{F0: 1}
                            _ = v84
                        case -119:
                            switch ("foo bar" + ("hello" + "")) {
                            case "hello":
                                v2 = append(v2, 780)
                            }
                        default:
                            v85 := GenS7{F0: "hello", F1: 18446744073709551615, F2: (v4 >> ((65535) & 63))}
                            _ = v85
                        }
                        var v86 uint = 573
                        v87 := GenS9{F0: ""}
                        _ = fi82
                        _ = v86
                        _ = v87
                    }
                    var v88 int32 = 2147483647
                    _ = v88
                } else {
                    if len(v2) > 0 {
                        v2[0] = uint32(v6)
                    }
                    var v89 float32 = 1.5
                    _ = v89
                }
            } else {
                v6 = 97
                if (-728 >= -2147483648) {
                    var v90 string = (("hello" + ("" + "foo bar")) + "foo bar")
                    switch "a" {
                    case "foo bar":
                        v1 = false
                    default:
                        for fi91 := 0; fi91 < 4; fi91++ {
                            for fi92 := 0; fi92 < 9; fi92++ {
                                if v1 {
                                    v93 := GenS10{F0: v5, F1: (1.5 >= 3.14159), F2: 42}
                                    _ = v93
                                } else {
                                    _ = len(v2)
                                    var v94 int16 = v5
                                    v5 = int16(v6)
                                    _ = v94
                                }
                                _ = fi92
                            }
                            v95 := GenS1{F0: 66, F1: (v6 >> ((uint(v4)) & 63))}
                            for fi96 := 0; fi96 < 12; fi96++ {
                                if v1 {
                                    if len(v2) > 0 {
                                        v2[0] = 0
                                    }
                                    if v1 {
                                        switch ("hello" + ((v90 + v90) + "foo bar")) {
                                        case "hello":
                                            if true {
                                                var v97 int8 = 42
                                                switch v90 {
                                                case "":
                                                    v90 = ("a" + (("\n\t" + "a") + "hello"))
                                                }
                                                _ = v97
                                            }
                                            if ((v5 ^ (v5 - (v5 >> ((2) & 63)))) > -32768) {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            if (-2.25 < (0.1 - float32((((1.5 + (1.5 * float64(3.14159))) + 3.14159) / 0.1)))) {
                                                var v98 float32 = 1.5
                                                _ = v98
                                            }
                                            var v99 uint = 110
                                            _ = v99
                                        }
                                    } else {
                                        switch v90 {
                                        case "":
                                            fmt.Println(v1, v1, fi96)
                                        case "foo bar":
                                            v5 -= (v5 ^ 1)
                                        }
                                        var v100 int32 = -1
                                        v101 := GenS4{F0: 621}
                                        _ = v100
                                        _ = v101
                                    }
                                } else {
                                    v1 = (18446744073709551615 <= 882)
                                }
                                var v102 float32 = 1.5
                                var v103 uint8 = v4
                                _ = fi96
                                _ = v102
                                _ = v103
                            }
                            _ = fi91
                            _ = v95
                        }
                    }
                    _ = v90
                } else {
                    for fi104 := 0; fi104 < 1; fi104++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if true {
                            for fi105 := 0; fi105 < 9; fi105++ {
                                var v106 int32 = -2147483648
                                var v107 uint = 18446744073709551615
                                _ = fi105
                                _ = v106
                                _ = v107
                            }
                            var v108 bool = true
                            for fi109 := 0; fi109 < 6; fi109++ {
                                v1 = ((0.1 / 1e10) > (0.0 * (float64(1.5) + 0.1)))
                                _ = len(v2)
                                v5 -= -688
                                _ = fi109
                            }
                            _ = v108
                        } else {
                            fmt.Println(v1, v5)
                            v2 = append(v2, 1)
                            for fi110 := 0; fi110 < 1; fi110++ {
                                var v111 int64 = v30
                                if len(v2) > 0 {
                                    v2[0] = 2
                                }
                                _ = fi110
                                _ = v111
                            }
                        }
                        v30 += int64(v5)
                        _ = fi104
                    }
                }
            }
            switch "hello" {
            case "":
                v112 := GenS11{F0: v5, F1: 42, F2: 18446744073709551615}
                _ = v112
            case "hello":
                if (!(127 <= 127)) {
                    var v113 uint = 42
                    v5 -= v5
                    var v114 float32 = 0.1
                    _ = v113
                    _ = v114
                }
            case "\n\t":
                v1 = (127 > 581)
            default:
                v115 := GenS10{F0: -32768, F1: ((v30 ^ v30) == (v3 & v3)), F2: 937}
                _ = v115
            }
        } else {
            if (-2.25 <= -2.25) {
                var v116 uint64 = 18446744073709551615
                v116 += (v116 + 42)
                _ = v116
            } else {
                for fi117 := 0; fi117 < 3; fi117++ {
                    v118 := GenS9{F0: "hello"}
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                    _ = fi117
                    _ = v118
                }
            }
        }
    }
    if (!v1) {
        v5 += (-(v5))
        if v1 {
            var v119 uint = 18446744073709551615
            for fi120 := 0; fi120 < 10; fi120++ {
                fmt.Println(v30)
                if v1 {
                    v121 := GenS12{F0: 2147483647, F1: (v4 ^ (v4 | v4))}
                    _ = v121
                }
                _ = fi120
            }
            v6 += (^(v6))
            _ = v119
        } else {
            if false {
                v1 = v1
                fmt.Println(v1, v4)
            } else {
                var v122 uint64 = 991
                var v123 bool = true
                _ = v122
                _ = v123
            }
            fmt.Println(v3, v3)
            fmt.Println(v6, v30)
        }
        v4 -= (v4 / ((v4) | 1))
    } else {
        var v124 int64 = v30
        v125 := GenS8{F0: v6, F1: v5}
        v126 := GenS13{F0: 1}
        _ = v124
        _ = v125
        _ = v126
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
