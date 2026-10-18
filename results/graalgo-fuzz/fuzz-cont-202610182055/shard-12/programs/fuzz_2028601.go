package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 int32; F2 bool }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint16; F1 uint8 }
type GenS10 struct { F0 int64; F1 bool; F2 int32 }
type GenS11 struct { F0 bool; F1 uint }
type GenS12 struct { F0 string; F1 uint16; F2 bool }
type GenS13 struct { F0 string; F1 uint }
func genFn0(p0 int64, p1 uint) int {
    p0 = (^(p0))
    switch "" {
    case "":
        var v4 uint64 = 4294967295
        _ = v4
    case "\n\t":
        v5 := GenS1{F0: p0}
        _ = v5
    }
    switch -905 {
    case -2147483648:
        switch 127 {
        case 127:
            for fi6 := 0; fi6 < 12; fi6++ {
                for fi7 := 0; fi7 < 17; fi7++ {
                    v8 := []int{(-(fi7)), (-(fi6)), (fi7 << ((4294967295) & 63)), (fi6 >> ((p1) & 63))}
                    p0 -= (p0 >> (((p1 << (((p1 ^ (^(p1)))) & 63))) & 63))
                    var v9 uint64 = 74
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                switch "hello" {
                case "a":
                    v10 := GenS2{F0: p0, F1: 2147483648}
                    _ = v10
                }
                for fi11 := 0; fi11 < 19; fi11++ {
                    v12 := GenS1{F0: (p0 & p0)}
                    p0 += (p0 / ((p0) | 1))
                    v13 := GenS3{F0: p0, F1: (p1 / ((p1) | 1))}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = fi6
            }
        case 0:
            v14 := GenS2{F0: p0, F1: 455}
            _ = v14
        case 2147483647:
            p0 = (^(p0))
        }
    case 128:
        var v15 uint = (p1 << ((p1) & 63))
        _ = v15
    case 1:
        v16 := GenS3{F0: (p0 % ((p0) | 1)), F1: p1}
        _ = v16
    default:
        if false {
            var v17 uint16 = 1
            v18 := GenS4{F0: 127, F1: 183}
            switch int(p0) {
            case 981:
                var v19 int = -869
                _ = v19
            case 0:
                p1 += p1
            case 762:
                fmt.Println(v17)
            default:
                p1 -= 934
            }
            _ = v17
            _ = v18
        }
    }
    fmt.Println(p1)
    for fi20 := 0; fi20 < 5; fi20++ {
        v21 := GenS5{F0: p1}
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int32, p1 int) int64 {
    if false {
        v22 := GenS4{F0: -128, F1: 2}
        p1 = 42
        var v23 int = (p1 * (-(p1)))
        _ = v22
        _ = v23
    }
    if false {
        v24 := []int{}
        _ = v24
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 int8 = 100
        p1 += p1
        if false {
            for fi27 := 0; fi27 < 1; fi27++ {
                var v28 bool = false
                var v29 float64 = ((0.1 * -2.25) * 0.0)
                _ = fi27
                _ = v28
                _ = v29
            }
        } else {
            if false {
                p1 -= (-(p1))
                var v30 uint32 = 2147483648
                switch -9223372036854775808 {
                case 128:
                    fmt.Println(v26)
                case 32768:
                    for fi31 := 0; fi31 < 8; fi31++ {
                        switch "foo bar" {
                        case "a":
                            v32 := []int{genFn0(42, 4294967295), genFn0(-9223372036854775808, 65535), (fi31 * 127)}
                            _ = v32
                        case "":
                            v26 += v26
                        case "hello":
                            if true {
                                if false {
                                    v33 := []int{genFn0(42, 18446744073709551615), (p1 * (fi25 | int(p0))), (p1 + (^(fi31))), int(v26)}
                                    p1 = (p1 * (fi31 | (fi25 + (p1 ^ p1))))
                                    _ = v33
                                } else {
                                    switch "" {
                                    case "a":
                                        if true {
                                            var v34 uint = 0
                                            fmt.Println(v34, v34)
                                            v35 := []int{(p1 >> ((243) & 63)), genFn0(173, v34)}
                                            if ("foo bar" != "hello") {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v36 uint = v34
                                                    _ = v36
                                                default:
                                                    _ = len(v35)
                                                }
                                                var v37 uint8 = 255
                                                v38 := GenS4{F0: 42, F1: v37}
                                                _ = v37
                                                _ = v38
                                            }
                                            _ = v34
                                            _ = v35
                                        }
                                    case "hello":
                                        fmt.Println(v26, v30, fi31)
                                    default:
                                        if (255 < 533) {
                                            v30 += v30
                                            var v39 int64 = -9223372036854775808
                                            if ((((float32((((((0.1 / (0.1 / 1e10)) + 1.5) / 1e10) + 1e10) / (((float64(3.14159) + 1e10) + 1.5) - 0.0))) * 1e10) + ((1e10 + (0.0 + 0.1)) + ((((0.1 - 1e10) * (0.1 + float32((0.1 - 3.14159)))) / 1e10) + 1.5))) + (float32(1.5) / 1.0)) > 0.0) {
                                                p0 = 42
                                                v40 := []int{int(v26), genFn0((v39 + (v39 << ((45) & 63))), 255), (p1 << ((652) & 63)), fi25}
                                                _ = v40
                                            } else {
                                                switch (p1 - (fi31 + genFn0(-9223372036854775808, 963))) {
                                                case 32768:
                                                    var v41 uint = 255
                                                    _ = v41
                                                }
                                                switch "\n\t" {
                                                case "":
                                                    fmt.Println(fi31)
                                                default:
                                                    var v42 int = (fi25 & (fi31 | (p1 - (p1 - p1))))
                                                    _ = v42
                                                }
                                                if (!(!((v39 + 42) >= -1))) {
                                                    if (true || ((v26 - int8(p0)) < v26)) {
                                                        var v43 uint8 = 214
                                                        _ = v43
                                                    }
                                                    if (!false) {
                                                        var v44 int16 = -1
                                                        v45 := GenS4{F0: (v26 % (((v26 / ((127) | 1))) | 1)), F1: 1}
                                                        fmt.Println(v39)
                                                        fmt.Println(p0)
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        v30 = v30
                                                        var v46 string = "hello"
                                                        v30 = (v30 % (((v30 * 2)) | 1))
                                                        _ = v46
                                                    }
                                                }
                                            }
                                            fmt.Println(v30)
                                            v39 = (v39 ^ (^(v39)))
                                            _ = v39
                                        } else {
                                            v47 := []int{(-(fi31)), (fi25 >> ((0) & 63)), (fi31 ^ 0), p1}
                                            _ = v47
                                        }
                                    }
                                    if (0.1 == float64(0.1)) {
                                        var v48 float64 = (float64(0.0) + 0.1)
                                        fmt.Println(p1)
                                        v48 += v48
                                        if true {
                                            if ("hello" >= ("a" + "\n\t")) {
                                                fmt.Println(v26, v30)
                                                if false {
                                                    var v49 bool = true
                                                    v50 := GenS6{F0: 4294967295, F1: p0, F2: (1 == 65535)}
                                                    v51 := GenS1{F0: 196}
                                                    _ = v49
                                                    _ = v50
                                                    _ = v51
                                                }
                                                var v52 int16 = 1
                                                _ = v52
                                            }
                                        } else {
                                            var v53 uint32 = v30
                                            var v54 uint32 = (^(v30))
                                            v26 -= 42
                                            _ = v53
                                            _ = v54
                                        }
                                        _ = v48
                                    }
                                    v55 := GenS1{F0: -9223372036854775808}
                                    _ = v55
                                }
                            }
                        }
                        _ = fi31
                    }
                default:
                    for fi56 := 0; fi56 < 18; fi56++ {
                        p1 = p1
                        switch fi56 {
                        case 9223372036854775807:
                            v57 := []int{int(p0), -37, (fi56 ^ (^(fi56))), p1}
                            _ = v57
                        case -2147483648:
                            switch -46 {
                            case -500:
                                if false {
                                    var v58 float32 = 0.0
                                    v59 := GenS7{F0: -9223372036854775808, F1: (v26 % ((v26) | 1))}
                                    _ = v58
                                    _ = v59
                                }
                                v26 += 108
                            default:
                                v26 = 100
                            }
                        default:
                            var v60 int = (^(p1))
                            _ = v60
                        }
                        v61 := []int{}
                        if (!false) {
                            v30 -= (v30 << ((108) & 63))
                            var v62 string = "foo bar"
                            var v63 bool = false
                            _ = v62
                            _ = v63
                        } else {
                            var v64 string = (("" + ("a" + "hello")) + ("" + ("" + "")))
                            var v65 float32 = float32((-2.25 - -2.25))
                            var v66 uint16 = 184
                            _ = v64
                            _ = v65
                            _ = v66
                        }
                        var v67 string = ("foo bar" + "foo bar")
                        if false {
                            p0 = 42
                        }
                        _ = fi56
                        _ = v61
                        _ = v67
                    }
                }
                _ = v30
            } else {
                v68 := GenS4{F0: (v26 ^ v26), F1: 1}
                _ = v68
            }
            v69 := GenS8{F0: false}
            _ = v69
        }
        _ = fi25
        _ = v26
    }
    _ = p0
    _ = p1
    return -730
}

func main() {
    var v1 float32 = 3.14159
    v2 := []int{}
    var v3 float32 = v1
    v70 := GenS6{F0: 2, F1: -2147483648, F2: false}
    v3 = 0.0
    if (true && (-703 <= genFn0(1, 4294967295))) {
        if (42 >= 1) {
            v3 += v1
        }
        switch 717 {
        case -9223372036854775808:
            if (((((float32(-2.25) * -2.25) * 3.14159) + 3.14159) / v3) <= 3.14159) {
                for fi71 := 0; fi71 < 8; fi71++ {
                    v1 = 1.5
                    fmt.Println(fi71, fi71)
                    _ = fi71
                }
            }
        default:
            v1 = float32(3.14159)
        }
        v72 := GenS8{F0: true}
        _ = v72
    } else {
        for fi73 := 0; fi73 < 4; fi73++ {
            v74 := GenS5{F0: 100}
            _ = fi73
            _ = v74
        }
        v3 += (((v3 + v1) - v3) / v1)
        var v75 string = "a"
        _ = v75
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    var v76 uint8 = 81
    v77 := GenS3{F0: -349, F1: 934}
    for fi78 := 0; fi78 < 17; fi78++ {
        fmt.Println(v3, v3)
        if (0 == 0) {
            var v79 int64 = 0
            if (32767 == int16(v76)) {
                if (2147483648 > 262) {
                    v2 = append(v2, fi78)
                    for fi80 := 0; fi80 < 10; fi80++ {
                        if (!true) {
                            if len(v2) > 0 {
                                v2[0] = -997
                            }
                            fmt.Println(v3)
                            v1 = (((-2.25 + float32(0.0)) * (float32(-2.25) - float32(0.1))) + (v3 * v3))
                        }
                        var v81 uint32 = 2147483648
                        _ = fi80
                        _ = v81
                    }
                } else {
                    switch "foo bar" {
                    case "":
                        if true {
                            for fi82 := 0; fi82 < 15; fi82++ {
                                v79 = (v79 >> ((452) & 63))
                                _ = fi82
                            }
                            var v83 string = ("hello" + ("hello" + ""))
                            _ = v83
                        }
                    case "foo bar":
                        var v84 int = (fi78 >> ((18446744073709551615) & 63))
                        _ = v84
                    case "hello":
                        if (!(true && true)) {
                            v1 = ((float32(3.14159) * (v1 * v3)) - v3)
                            if len(v2) > 0 {
                                v2[0] = (^(fi78))
                            }
                        }
                    }
                    v85 := GenS7{F0: v79, F1: 127}
                    _ = v85
                }
                v86 := GenS9{F0: 769, F1: 42}
                for fi87 := 0; fi87 < 3; fi87++ {
                    var v88 uint = 941
                    _ = fi87
                    _ = v88
                }
                _ = v86
            } else {
                if false {
                    v89 := GenS10{F0: (^(v79)), F1: (!(!false)), F2: -2147483648}
                    switch 975 {
                    case 4294967295:
                        if false {
                            v90 := GenS1{F0: genFn1(357, 236)}
                            v3 += (v3 / v1)
                            var v91 float32 = v1
                            _ = v90
                            _ = v91
                        } else {
                            var v92 string = ("hello" + "a")
                            for fi93 := 0; fi93 < 17; fi93++ {
                                _ = len(v2)
                                _ = fi93
                            }
                            _ = v92
                        }
                    case 32768:
                        if ((v76 >> ((1) & 63)) <= (v76 << ((42) & 63))) {
                            v76 = (v76 % ((v76) | 1))
                            var v94 uint = 18446744073709551615
                            _ = v94
                        } else {
                            var v95 int32 = 42
                            _ = v95
                        }
                    default:
                        v79 = v79
                    }
                    _ = v89
                }
                fmt.Println(v3, v76, v76)
            }
            _ = len(v2)
            _ = v79
        } else {
            for fi96 := 0; fi96 < 20; fi96++ {
                v1 = (0.0 / float32(3.14159))
                v2 = append(v2, (fi96 ^ -278))
                _ = fi96
            }
            var v97 uint = 4294967295
            _ = v97
        }
        v76 = v76
        _ = fi78
    }
    v76 -= 42
    if true {
        if false {
            for fi98 := 0; fi98 < 12; fi98++ {
                if false {
                    for fi99 := 0; fi99 < 8; fi99++ {
                        var v100 int16 = 32767
                        var v101 uint16 = 42
                        _ = fi99
                        _ = v100
                        _ = v101
                    }
                    v3 = v1
                    if true {
                        var v102 uint = 18446744073709551615
                        v102 = uint(v76)
                        _ = v102
                    }
                } else {
                    var v103 int64 = -869
                    v104 := GenS11{F0: (1.5 <= ((3.14159 * -2.25) / float64((0.1 / v3)))), F1: 4294967295}
                    fmt.Println(v1)
                    _ = v103
                    _ = v104
                }
                var v105 string = ""
                _ = fi98
                _ = v105
            }
            var v106 int8 = 127
            var v107 bool = (!(!(!(!(65535 != 2147483647)))))
            _ = v106
            _ = v107
        }
        v1 -= float32(0.0)
    }
    var v108 int64 = -1
    fmt.Println(v1)
    switch ("foo bar" + "hello") {
    case "":
        v109 := GenS12{F0: (((("\n\t" + "hello") + "foo bar") + ("" + ("foo bar" + "\n\t"))) + ""), F1: 65535, F2: false}
        _ = v109
    case "foo bar":
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v110 int8 = 55
            switch genFn0((v108 << ((42) & 63)), 4294967295) {
            case 2147483647:
                for fi111 := 0; fi111 < 14; fi111++ {
                    var v112 int64 = (v108 | 9223372036854775807)
                    v113 := GenS13{F0: ("\n\t" + (("" + "foo bar") + "a")), F1: 255}
                    v108 = genFn1(1, int(v108))
                    _ = fi111
                    _ = v112
                    _ = v113
                }
            case 1:
                switch 2147483647 {
                case 1:
                    v114 := GenS7{F0: (-(v108)), F1: (v110 % ((v110) | 1))}
                    _ = v114
                case 2147483647:
                    v2 = append(v2, genFn0(genFn1(-2147483648, -2147483648), 255))
                default:
                    v115 := GenS4{F0: v110, F1: (v76 >> ((1) & 63))}
                    _ = v115
                }
            case -975:
                if false {
                    var v116 int = -2147483648
                    v110 = v110
                    _ = v116
                }
            default:
                var v117 int = 32768
                _ = v117
            }
            _ = v110
        } else {
            if false {
                if true {
                    for fi118 := 0; fi118 < 12; fi118++ {
                        var v119 float32 = (v1 * v1)
                        _ = fi118
                        _ = v119
                    }
                    v2 = append(v2, 0)
                    switch ("a" + "") {
                    case "":
                        v3 += v1
                    case "\n\t":
                        v76 = v76
                    case "a":
                        var v120 uint8 = (^(v76))
                        _ = v120
                    }
                } else {
                    v2 = append(v2, 1)
                }
                var v121 uint64 = 18446744073709551615
                if true {
                    var v122 uint8 = (v76 % ((v76) | 1))
                    _ = v122
                }
                _ = v121
            } else {
                v2 = append(v2, -274)
                for fi123 := 0; fi123 < 1; fi123++ {
                    for fi124 := 0; fi124 < 6; fi124++ {
                        if false {
                            fmt.Println(fi124, v76)
                            fmt.Println(fi124)
                            if (-9223372036854775808 > (fi123 - fi124)) {
                                var v125 int64 = v108
                                _ = v125
                            }
                        } else {
                            fmt.Println(fi124, fi124)
                            _ = len(v2)
                            var v126 uint32 = 1000
                            _ = v126
                        }
                        _ = fi124
                    }
                    _ = fi123
                }
            }
            v2 = append(v2, 49)
            v2 = append(v2, -153)
        }
    case "\n\t":
        var v127 uint32 = 4294967295
        _ = v127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v108
    _ = v2
    _ = v3
    _ = v70
    _ = v76
    _ = v77
}
