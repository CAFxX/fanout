package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int; F1 uint8; F2 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS10 struct { F0 int32; F1 uint }
type GenS11 struct { F0 string; F1 uint16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 128
    v3 := make(map[uint32]bool)
    var v4 int32 = 0
    v5 := []int{-83}
    var v6 int32 = (v4 * v4)
    if true {
        var v7 uint = 1
        if (1 >= 2) {
            v2 -= (v2 + v2)
        }
        if ((v4 ^ 795) == (v6 << (((^(v7))) & 63))) {
            var v8 int16 = 32767
            switch -307 {
            case 4294967295:
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    if true {
                        for fi9 := 0; fi9 < 5; fi9++ {
                            v7 += (v7 >> (((^(v7))) & 63))
                            _ = fi9
                        }
                        fmt.Println(v2)
                        if (1 >= 2) {
                            v4 -= (v4 << ((v7) & 63))
                            var v10 float64 = 1e10
                            switch 127 {
                            case -9223372036854775808:
                                if true {
                                    for fi11 := 0; fi11 < 11; fi11++ {
                                        if false {
                                            if true {
                                                v12 := GenS1{F0: uint16(v2), F1: 32767}
                                                var v13 float64 = v10
                                                _ = v12
                                                _ = v13
                                            } else {
                                                var v14 uint16 = 222
                                                fmt.Println(v7)
                                                if (v14 < (v14 - (^(v14)))) {
                                                    var v15 int = fi11
                                                    _ = v15
                                                }
                                                _ = v14
                                            }
                                            var v16 uint32 = (v2 / (((v2 % ((1) | 1))) | 1))
                                            delete(v1, (fi11 + fi11))
                                            _ = v16
                                        }
                                        _ = fi11
                                    }
                                    if false {
                                        if true {
                                            v17 := GenS1{F0: 42, F1: 743}
                                            _ = v17
                                        } else {
                                            v18 := GenS2{F0: -9223372036854775808, F1: 1, F2: (v8 * (v8 & (v8 >> ((0) & 63))))}
                                            _ = v18
                                        }
                                    } else {
                                        if false {
                                            var v19 int8 = 42
                                            var v20 uint16 = 32767
                                            v5 = append(v5, 802)
                                            _ = v19
                                            _ = v20
                                        } else {
                                            if ((((((("foo bar" + ("" + ("" + "\n\t"))) + ("foo bar" + "\n\t")) + "foo bar") + "hello") + "") + ("a" + "foo bar")) < "a") {
                                                if true {
                                                    if false {
                                                        v6 = v6
                                                        v10 -= 3.14159
                                                        for fi21 := 0; fi21 < 12; fi21++ {
                                                            if true {
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    switch ("\n\t" + ("\n\t" + ("\n\t" + "foo bar"))) {
                                                                    case "hello":
                                                                        var v22 string = "\n\t"
                                                                        _ = v22
                                                                    case "foo bar":
                                                                        v2 += 1
                                                                    default:
                                                                        if false {
                                                                            var v23 int16 = v8
                                                                            var v24 uint64 = 4294967295
                                                                            _ = len(v1)
                                                                            _ = v23
                                                                            _ = v24
                                                                        } else {
                                                                            v25 := GenS3{F0: (("\n\t" + ("" + "a")) + (("foo bar" + ("hello" + ("hello" + "a"))) + ""))}
                                                                            _ = v25
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    fmt.Println(v6)
                                                                }
                                                                _ = len(v3)
                                                            } else {
                                                                var v26 float64 = v10
                                                                v27 := GenS3{F0: "\n\t"}
                                                                _ = v26
                                                                _ = v27
                                                            }
                                                            var v28 int = (-(fi21))
                                                            var v29 int8 = 1
                                                            _ = fi21
                                                            _ = v28
                                                            _ = v29
                                                        }
                                                    } else {
                                                        var v30 bool = ((v7 * v7) == (^(v7)))
                                                        _ = v30
                                                    }
                                                    v8 += (v8 & v8)
                                                    for fi31 := 0; fi31 < 18; fi31++ {
                                                        v32 := v1[(-(fi31))]
                                                        if false {
                                                            v33 := GenS1{F0: 32767, F1: 0}
                                                            v7 += (v7 | (v7 << ((1) & 63)))
                                                            _ = len(v5)
                                                            _ = v33
                                                        } else {
                                                            var v34 string = (("" + "") + "foo bar")
                                                            _ = v34
                                                        }
                                                        _ = fi31
                                                        _ = v32
                                                    }
                                                } else {
                                                    if (!(!true)) {
                                                        v35 := GenS3{F0: ("foo bar" + "hello")}
                                                        var v36 uint64 = 4294967295
                                                        v4 = -1
                                                        _ = v35
                                                        _ = v36
                                                    } else {
                                                        if true {
                                                            v8 = (v8 + 1)
                                                            for fi37 := 0; fi37 < 20; fi37++ {
                                                                if (v7 != (v7 - v7)) {
                                                                    for fi38 := 0; fi38 < 1; fi38++ {
                                                                        if true {
                                                                            switch (fi38 >> (((^(v7))) & 63)) {
                                                                            case -146:
                                                                                v5 = append(v5, -1)
                                                                            case 128:
                                                                                var v39 bool = (true && true)
                                                                                _ = v39
                                                                            case -188:
                                                                                v10 = (v10 / 3.14159)
                                                                            default:
                                                                                var v40 uint64 = 0
                                                                                _ = v40
                                                                            }
                                                                            fmt.Println(v4, v8)
                                                                        }
                                                                        _ = fi38
                                                                    }
                                                                } else {
                                                                    var v41 bool = (!(v10 >= (v10 * v10)))
                                                                    v4 = v6
                                                                    var v42 uint = 255
                                                                    _ = v41
                                                                    _ = v42
                                                                }
                                                                for fi43 := 0; fi43 < 15; fi43++ {
                                                                    var v44 uint = v7
                                                                    fmt.Println(v6)
                                                                    if true {
                                                                        v45 := GenS3{F0: ""}
                                                                        v46 := GenS3{F0: ("a" + "a")}
                                                                        v47 := GenS3{F0: ""}
                                                                        v48 := GenS3{F0: ("hello" + ("a" + ""))}
                                                                        _ = v45
                                                                        _ = v46
                                                                        _ = v47
                                                                        _ = v48
                                                                    }
                                                                    _ = fi43
                                                                    _ = v44
                                                                }
                                                                _ = fi37
                                                            }
                                                            v10 = (0.0 / 1.0)
                                                        }
                                                    }
                                                }
                                                v49 := v1[-9223372036854775808]
                                                fmt.Println(v6)
                                                _ = v49
                                            } else {
                                                for fi50 := 0; fi50 < 12; fi50++ {
                                                    v3[(v2 | (v2 << (((v7 * (v7 << ((946) & 63)))) & 63)))] = false
                                                    _ = fi50
                                                }
                                                if true {
                                                    switch ("\n\t" + "a") {
                                                    case "foo bar":
                                                        v51 := GenS3{F0: ("\n\t" + "hello")}
                                                        _ = v51
                                                    default:
                                                        v4 += (^(v6))
                                                    }
                                                } else {
                                                    if (!(false && true)) {
                                                        var v52 uint32 = (v2 * 2147483648)
                                                        if true {
                                                            for fi53 := 0; fi53 < 11; fi53++ {
                                                                var v54 uint64 = 0
                                                                _ = fi53
                                                                _ = v54
                                                            }
                                                            var v55 int64 = -9223372036854775808
                                                            for fi56 := 0; fi56 < 18; fi56++ {
                                                                for fi57 := 0; fi57 < 10; fi57++ {
                                                                    var v58 float64 = ((3.14159 * v10) * v10)
                                                                    var v59 int16 = (v8 - -218)
                                                                    switch fi57 {
                                                                    case -2147483648:
                                                                        var v60 int16 = -1
                                                                        _ = v60
                                                                    default:
                                                                        v61 := GenS1{F0: 0, F1: 1}
                                                                        _ = v61
                                                                    }
                                                                    switch fi57 {
                                                                    case 0:
                                                                        v62 := GenS1{F0: 255, F1: 255}
                                                                        _ = v62
                                                                    case -2147483648:
                                                                        var v63 uint16 = 1
                                                                        _ = v63
                                                                    default:
                                                                        fmt.Println(v10, v10, v52)
                                                                    }
                                                                    v58 = (0.0 * v58)
                                                                    fmt.Println(v6, v10)
                                                                    v10 = (((v58 - (((v10 + (float64(-2.25) + v58)) - v58) / (float64(-2.25) / -2.25))) - v58) + float64(1.5))
                                                                    _ = fi57
                                                                    _ = v58
                                                                    _ = v59
                                                                }
                                                                v2 = 4294967295
                                                                _ = fi56
                                                            }
                                                            _ = v55
                                                        } else {
                                                            v64 := GenS4{F0: "a"}
                                                            v10 = (v10 - (v10 - ((v10 - v10) - (v10 - (float64(3.14159) / float64((1e10 / 0.1)))))))
                                                            _ = v64
                                                        }
                                                        var v65 int8 = 93
                                                        _ = v52
                                                        _ = v65
                                                    } else {
                                                        if true {
                                                            v8 += int16(v7)
                                                            switch "\n\t" {
                                                            case "":
                                                                var v66 string = ""
                                                                _ = v66
                                                            case "hello":
                                                                switch 127 {
                                                                case 128:
                                                                    var v67 int64 = 230
                                                                    _ = v67
                                                                default:
                                                                    switch -9223372036854775808 {
                                                                    case -66:
                                                                        v1[-2147483648] = 597
                                                                    case -795:
                                                                        var v68 float32 = (0.1 * 1e10)
                                                                        _ = v68
                                                                    case -1:
                                                                        v8 = (v8 | (-(v8)))
                                                                    }
                                                                }
                                                            default:
                                                                var v69 float32 = (1e10 / 3.14159)
                                                                _ = v69
                                                            }
                                                        } else {
                                                            for fi70 := 0; fi70 < 13; fi70++ {
                                                                v8 = -919
                                                                _ = fi70
                                                            }
                                                        }
                                                        if ((1 < 42) && (-988 > 0)) {
                                                            v4 = (-(v4))
                                                            var v71 int32 = v6
                                                            _ = v71
                                                        } else {
                                                            for fi72 := 0; fi72 < 18; fi72++ {
                                                                if false {
                                                                    v6 = (v6 / ((int32(v8)) | 1))
                                                                } else {
                                                                    var v73 float32 = float32((v10 + (((v10 / v10) * v10) / 1.0)))
                                                                    _ = v73
                                                                }
                                                                var v74 uint8 = 122
                                                                _ = fi72
                                                                _ = v74
                                                            }
                                                            var v75 int64 = -9223372036854775808
                                                            _ = v75
                                                        }
                                                        var v76 uint = v7
                                                        _ = v76
                                                    }
                                                    var v77 float64 = ((v10 - (v10 / 1.5)) - (0.0 - (float64(-2.25) * v10)))
                                                    _ = v77
                                                }
                                                if true {
                                                    var v78 bool = false
                                                    _ = v78
                                                }
                                            }
                                        }
                                    }
                                    var v79 int32 = (v6 % ((-372) | 1))
                                    _ = v79
                                }
                            case -2147483648:
                                fmt.Println(v4)
                            }
                            _ = v10
                        } else {
                            if (255 > 42) {
                                for fi80 := 0; fi80 < 15; fi80++ {
                                    delete(v1, (fi80 & fi80))
                                    _ = fi80
                                }
                                if ((-1 > -113) && (18446744073709551615 <= 459)) {
                                    fmt.Println(v8)
                                }
                            } else {
                                for fi81 := 0; fi81 < 1; fi81++ {
                                    v82 := GenS5{F0: 127, F1: -991}
                                    var v83 float32 = float32(1.5)
                                    delete(v3, 1)
                                    _ = fi81
                                    _ = v82
                                    _ = v83
                                }
                                var v84 uint64 = 2
                                _ = v84
                            }
                            switch "\n\t" {
                            case "hello":
                                v7 = v7
                            }
                        }
                    } else {
                        var v85 float32 = 1e10
                        _ = len(v5)
                        if false {
                            switch 0 {
                            case 42:
                                fmt.Println(v85, v7, v4)
                            case 128:
                                var v86 uint16 = 115
                                _ = v86
                            case 1:
                                v4 += v4
                            }
                        } else {
                            v87 := GenS6{F0: v7}
                            _ = v87
                        }
                        _ = v85
                    }
                default:
                    v4 -= (v4 % ((1) | 1))
                }
            }
            _ = v8
        }
        _ = v7
    } else {
        v88 := v1[2147483647]
        if (-799 == 9223372036854775807) {
            var v89 uint16 = 980
            _ = v89
        } else {
            if (!(true && (665 == int64(v4)))) {
                for fi90 := 0; fi90 < 7; fi90++ {
                    var v91 uint64 = 2
                    for fi92 := 0; fi92 < 7; fi92++ {
                        fmt.Println(v6)
                        v4 = -1
                        _ = fi92
                    }
                    fmt.Println(v4, v6)
                    _ = fi90
                    _ = v91
                }
                v93 := GenS7{F0: 42, F1: 663}
                _ = v93
            }
            var v94 int64 = 1
            switch 2147483647 {
            case -2147483648:
                v4 += (-(v4))
            case 128:
                _ = len(v1)
            case 9223372036854775807:
                switch "a" {
                case "hello":
                    if true {
                        delete(v3, v88)
                        var v95 int32 = -642
                        var v96 float32 = 3.14159
                        _ = v95
                        _ = v96
                    }
                case "a":
                    if true {
                        var v97 int = -9223372036854775808
                        _ = v97
                    }
                case "foo bar":
                    if (!true) {
                        var v98 float64 = 0.0
                        var v99 uint64 = 321
                        switch 1 {
                        case 4294967295:
                            if true {
                                v88 -= (v88 ^ v88)
                            }
                        case 0:
                            for fi100 := 0; fi100 < 12; fi100++ {
                                switch fi100 {
                                case 42:
                                    v98 = ((((v98 + 0.0) - 1.5) * v98) + v98)
                                case -9223372036854775808:
                                    for fi101 := 0; fi101 < 6; fi101++ {
                                        v4 += (-(v4))
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v88 = 2147483648
                                        _ = fi101
                                    }
                                default:
                                    var v102 int16 = 0
                                    _ = v102
                                }
                                var v103 uint8 = 0
                                _ = fi100
                                _ = v103
                            }
                        }
                        _ = v98
                        _ = v99
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v104 uint32 = (v2 << ((0) & 63))
                        _ = v104
                    }
                }
            default:
                if len(v5) > 0 {
                    v5[0] = 579
                }
            }
            _ = v94
        }
        for fi105 := 0; fi105 < 10; fi105++ {
            v106 := GenS8{F0: 1, F1: (0.1 <= 1.5)}
            if true {
                switch fi105 {
                case -921:
                    v107 := GenS9{F0: (v2 % ((368) | 1)), F1: 42, F2: 42}
                    _ = v107
                case -689:
                    v108 := GenS1{F0: 0, F1: 1}
                    _ = v108
                default:
                    v2 -= (v88 << ((187) & 63))
                }
                if true {
                    switch int(v6) {
                    case 42:
                        var v109 int8 = 115
                        _ = v109
                    default:
                        var v110 int8 = 0
                        _ = v110
                    }
                }
                var v111 uint8 = 1
                _ = v111
            }
            _ = fi105
            _ = v106
        }
        _ = v88
    }
    for fi112 := 0; fi112 < 16; fi112++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v113 := GenS10{F0: 1, F1: 4294967295}
        _ = fi112
        _ = v113
    }
    var v114 int64 = 42
    for fi115 := 0; fi115 < 18; fi115++ {
        fmt.Println(v4)
        _ = fi115
    }
    var v116 uint64 = 4294967295
    var v117 uint = 0
    for fi118 := 0; fi118 < 10; fi118++ {
        fmt.Println(v116, v114, v2)
        v4 = (v6 | -2147483648)
        _ = fi118
    }
    v3[2147483648] = ((v6 % ((v4) | 1)) > v4)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v6, v117)
    if false {
        fmt.Println(v117)
        if ((^(v117)) <= (^(v117))) {
            switch -2147483648 {
            case 128:
                v2 = v2
            case 4294967295:
                var v119 uint16 = 535
                _ = v119
            default:
                if false {
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v114, v6)
                    }
                    switch "a" {
                    case "":
                        fmt.Println(v114, v117)
                    case "hello":
                        switch (((("a" + "foo bar") + "foo bar") + "a") + ("a" + "hello")) {
                        case "hello":
                            v116 = v116
                        }
                    case "foo bar":
                        for fi120 := 0; fi120 < 18; fi120++ {
                            delete(v3, (v2 & 65535))
                            if true {
                                if false {
                                    v4 = 0
                                }
                                v121 := GenS1{F0: 2, F1: 0}
                                v122 := GenS6{F0: (v117 & v117)}
                                _ = v121
                                _ = v122
                            }
                            v123 := GenS8{F0: v114, F1: true}
                            _ = fi120
                            _ = v123
                        }
                    }
                }
            }
            v117 = 4294967295
        }
    } else {
        if ((v116 > v116) || (!(("" + "\n\t") == "\n\t"))) {
            if ("" < "foo bar") {
                var v124 int32 = -1
                _ = v124
            } else {
                var v125 uint32 = (v2 & v2)
                v4 += v6
                v126 := GenS11{F0: ((("" + (("foo bar" + "") + "foo bar")) + "foo bar") + ("" + "foo bar")), F1: 42}
                _ = v125
                _ = v126
            }
            fmt.Println(v4, v2)
            var v127 int8 = 100
            _ = v127
        }
        switch 554 {
        case 9223372036854775807:
            v128 := GenS8{F0: (v114 >> ((uint(v2)) & 63)), F1: true}
            _ = v128
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v114:", v114)
    fmt.Println("v116:", v116)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v114
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
