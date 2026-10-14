package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int64; F1 int8; F2 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint64; F1 int8; F2 bool }
type GenS6 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8; F1 int8; F2 string }
type GenS9 struct { F0 int32; F1 int8 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int64; F1 uint16; F2 uint32 }
func genFn0(p0 string, p1 uint64) bool {
    v7 := GenS1{F0: 1, F1: -30}
    var v8 float64 = (-2.25 / 0.1)
    switch -1 {
    case -1:
        if ((!false) || true) {
            if (529 < 255) {
                var v9 int = 127
                v8 = (v8 / 1.5)
                _ = v9
            }
        } else {
            if (true || false) {
                var v10 string = "a"
                _ = v10
            }
            if (0 > 255) {
                var v11 uint64 = p1
                _ = v11
            } else {
                p0 = (p0 + "a")
                if false {
                    fmt.Println(p0, p0, p1)
                    v12 := []int{32767, 32768, 9223372036854775807, 1}
                    p1 -= 1
                    _ = v12
                } else {
                    v13 := []int{-2147483648, 42}
                    _ = v13
                }
                var v14 int32 = 0
                _ = v14
            }
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (-318 <= 32767)
}
func genFn1() int {
    var v15 uint = 4294967295
    v16 := GenS1{F0: 1, F1: 100}
    v17 := []int{-9223372036854775808}
    for fi18 := 0; fi18 < 13; fi18++ {
        if (!true) {
            for fi19 := 0; fi19 < 8; fi19++ {
                switch -726 {
                case 32768:
                    switch fi19 {
                    case 0:
                        var v20 float32 = (3.14159 + 1e10)
                        _ = v20
                    }
                case -2147483648:
                    var v21 float64 = -2.25
                    fmt.Println(v21, fi19, fi19)
                    _ = len(v17)
                    _ = v21
                case 42:
                    if len(v17) > 0 {
                        v17[0] = -9223372036854775808
                    }
                }
                v22 := GenS2{F0: 545, F1: 127, F2: ((1 <= -1) && false)}
                var v23 bool = false
                _ = fi19
                _ = v22
                _ = v23
            }
            var v24 int64 = -976
            _ = v24
        } else {
            fmt.Println(fi18)
            if len(v17) > 0 {
                v17[0] = -480
            }
        }
        v25 := GenS3{F0: true}
        var v26 int = fi18
        _ = fi18
        _ = v25
        _ = v26
    }
    _ = v15
    _ = v16
    _ = v17
    return -2147483648
}

func main() {
    var v1 string = ("a" + ("a" + "hello"))
    var v2 uint16 = 32767
    v3 := []int{-22, 42}
    var v4 float32 = 1e10
    var v5 int32 = -1
    v6 := make(map[uint32]string)
    for fi27 := 0; fi27 < 3; fi27++ {
        if (!false) {
            switch (fi27 >> ((65535) & 63)) {
            case 685:
                v5 -= (v5 * (v5 / (((v5 ^ 2147483647)) | 1)))
            case 2147483647:
                v2 = v2
            case 9223372036854775807:
                v28 := GenS4{F0: -1, F1: v5}
                _ = v28
            }
            fmt.Println(fi27)
        }
        _ = fi27
    }
    switch 127 {
    case -925:
        var v29 int32 = 2147483647
        _ = v29
    case 4294967295:
        var v30 int32 = (v5 - int32(v2))
        _ = v30
    case 928:
        v5 -= (v5 * 640)
    default:
        if true {
            for fi31 := 0; fi31 < 8; fi31++ {
                var v32 string = v1
                var v33 string = "foo bar"
                _ = fi31
                _ = v32
                _ = v33
            }
            if true {
                if (true && (32767 <= 1)) {
                    for fi34 := 0; fi34 < 20; fi34++ {
                        switch v1 {
                        case "\n\t":
                            if true {
                                fmt.Println(v5, fi34)
                                if (false && false) {
                                    for fi35 := 0; fi35 < 14; fi35++ {
                                        var v36 uint8 = 16
                                        var v37 string = (v1 + "hello")
                                        var v38 bool = false
                                        _ = fi35
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    }
                                    delete(v6, 65535)
                                } else {
                                    switch (("\n\t" + "\n\t") + ("" + v1)) {
                                    case "\n\t":
                                        var v39 float32 = v4
                                        _ = v39
                                    }
                                    for fi40 := 0; fi40 < 3; fi40++ {
                                        if false {
                                            v41 := v6[65535]
                                            if ((1 == -92) || true) {
                                                var v42 uint32 = 761
                                                _ = v42
                                            } else {
                                                var v43 int64 = 9223372036854775807
                                                v41 = ("" + v1)
                                                v41 = (v41 + "a")
                                                _ = v43
                                            }
                                            _ = v41
                                        }
                                        _ = fi40
                                    }
                                }
                            }
                        case "foo bar":
                            fmt.Println(v5, v5)
                        case "hello":
                            v6[2] = "\n\t"
                        }
                        _ = fi34
                    }
                    switch -2147483648 {
                    case -597:
                        v1 = ""
                    case 32767:
                        v44 := GenS4{F0: 995, F1: 792}
                        _ = v44
                    case -1:
                        if (((float64((v4 * v4)) + (1.5 / float64(-2.25))) - ((0.0 - -2.25) / 3.14159)) <= float64(float32(1.5))) {
                            v1 = ("a" + ((("hello" + v1) + "foo bar") + (("\n\t" + "a") + v1)))
                        } else {
                            v45 := GenS1{F0: 36, F1: 0}
                            if false {
                                for fi46 := 0; fi46 < 12; fi46++ {
                                    var v47 uint8 = 127
                                    _ = fi46
                                    _ = v47
                                }
                                var v48 int16 = 32767
                                v3 = append(v3, 32767)
                                _ = v48
                            } else {
                                switch ("hello" + v1) {
                                case "a":
                                    if true {
                                        v4 = -2.25
                                        var v49 uint16 = (v2 - v2)
                                        var v50 int16 = -98
                                        _ = v49
                                        _ = v50
                                    }
                                case "\n\t":
                                    v1 = v1
                                case "hello":
                                    for fi51 := 0; fi51 < 19; fi51++ {
                                        v2 += (v2 & (v2 + v2))
                                        switch fi51 {
                                        case -2147483648:
                                            fmt.Println(v4, v4, v2)
                                        default:
                                            v4 = v4
                                        }
                                        _ = fi51
                                    }
                                }
                            }
                            _ = v45
                        }
                    default:
                        if ((^(v2)) < 672) {
                            switch (v1 + ((("" + "") + ((v1 + ("a" + v1)) + v1)) + (v1 + v1))) {
                            case "hello":
                                fmt.Println(v4)
                            case "a":
                                var v52 uint64 = 462
                                _ = v52
                            default:
                                if ((v1 + "foo bar") > v1) {
                                    var v53 int = -2147483648
                                    for fi54 := 0; fi54 < 9; fi54++ {
                                        if (!(((4294967295 != 681) && (((v2 / (((v2 & (v2 / ((v2) | 1)))) | 1)) <= (v2 & 0)) && (false && false))) || (v5 > (^(v5))))) {
                                            if (423 <= 622) {
                                                v53 += int(v2)
                                            } else {
                                                fmt.Println(fi54, v2)
                                            }
                                            delete(v6, 1)
                                        } else {
                                            for fi55 := 0; fi55 < 19; fi55++ {
                                                if (!(!false)) {
                                                    if false {
                                                        v4 += (float32(-2.25) / 0.1)
                                                        if (v5 > 32) {
                                                            fmt.Println(v4, fi54, v2)
                                                            if len(v3) > 0 {
                                                                v3[0] = (v53 + (fi55 * genFn1()))
                                                            }
                                                        }
                                                        var v56 uint8 = 128
                                                        _ = v56
                                                    }
                                                    v57 := GenS2{F0: -9223372036854775808, F1: -1, F2: false}
                                                    _ = v57
                                                } else {
                                                    if (!false) {
                                                        switch v1 {
                                                        case "a":
                                                            var v58 uint32 = 4294967295
                                                            _ = v58
                                                        }
                                                    }
                                                    var v59 uint = 469
                                                    var v60 int64 = -23
                                                    _ = v59
                                                    _ = v60
                                                }
                                                var v61 uint = 2
                                                _ = fi55
                                                _ = v61
                                            }
                                            v2 -= (v2 << ((4294967295) & 63))
                                            v4 -= -2.25
                                        }
                                        var v62 float32 = 1.5
                                        for fi63 := 0; fi63 < 2; fi63++ {
                                            var v64 uint32 = 1
                                            _ = fi63
                                            _ = v64
                                        }
                                        _ = fi54
                                        _ = v62
                                    }
                                    _ = v53
                                }
                            }
                        }
                    }
                    var v65 bool = (1000 <= -1)
                    _ = v65
                } else {
                    if (v1 < (v1 + "a")) {
                        if (2 < (v2 % (((v2 << ((4294967295) & 63))) | 1))) {
                            v1 = v1
                            v66 := GenS3{F0: true}
                            _ = v66
                        }
                        for fi67 := 0; fi67 < 10; fi67++ {
                            var v68 int = fi67
                            fmt.Println(v1, v1, fi67)
                            _ = fi67
                            _ = v68
                        }
                        for fi69 := 0; fi69 < 10; fi69++ {
                            for fi70 := 0; fi70 < 18; fi70++ {
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = fi70
                                    }
                                    switch fi70 {
                                    case 310:
                                        v71 := GenS5{F0: 3, F1: -103, F2: (v2 >= v2)}
                                        _ = v71
                                    case 2147483647:
                                        v5 = (v5 >> ((0) & 63))
                                    case 127:
                                        switch fi69 {
                                        case 0:
                                            v72 := GenS6{F0: 127, F1: 42, F2: 18446744073709551615}
                                            _ = v72
                                        case 9223372036854775807:
                                            var v73 uint = 65535
                                            _ = v73
                                        }
                                        v1 = ("foo bar" + (v1 + v1))
                                        fmt.Println(fi70)
                                        fmt.Println(v1, v4, v5)
                                        v5 = (^(v5))
                                    }
                                    v74 := GenS7{F0: 127}
                                    _ = v74
                                }
                                fmt.Println(fi69, fi69, v2)
                                delete(v6, 65535)
                                _ = fi70
                            }
                            _ = fi69
                        }
                    } else {
                        v75 := GenS8{F0: 1, F1: 100, F2: v1}
                        _ = v75
                    }
                }
            }
            if true {
                v4 = v4
                v5 = 547
            }
        }
    }
    v5 = v5
    var v76 uint = 255
    fmt.Println(v5)
    var v77 int64 = -9223372036854775808
    v78 := GenS9{F0: (v5 & (-(v5))), F1: 42}
    switch -1 {
    case 1:
        if (!(!true)) {
            v1 = v1
            for fi79 := 0; fi79 < 17; fi79++ {
                v76 = 255
                var v80 uint = (v76 >> ((v76) & 63))
                _ = fi79
                _ = v80
            }
        }
    case -1:
        v5 = 42
    case 4294967295:
        if false {
            var v81 uint64 = 0
            for fi82 := 0; fi82 < 15; fi82++ {
                var v83 uint = (^(v76))
                v76 = v83
                _ = fi82
                _ = v83
            }
            v77 = (^(v77))
            _ = v81
        } else {
            if ((0.1 - v4) < v4) {
                fmt.Println(v77, v76)
                v84 := GenS10{F0: 726}
                _ = v84
            } else {
                v2 = v2
                v85 := GenS7{F0: 113}
                v1 = "hello"
                _ = v85
            }
        }
    default:
        switch v1 {
        case "foo bar":
            var v86 int8 = 65
            _ = v86
        case "hello":
            _ = len(v6)
        }
    }
    if false {
        var v87 float64 = (float64((v4 * v4)) * 3.14159)
        v77 = (v77 ^ (^(v77)))
        var v88 uint16 = (^(v2))
        _ = v87
        _ = v88
    } else {
        var v89 uint8 = 127
        switch -9223372036854775808 {
        case 128:
            var v90 uint16 = 1
            _ = v90
        case 32767:
            v89 += v89
        case -2147483648:
            if true {
                for fi91 := 0; fi91 < 8; fi91++ {
                    for fi92 := 0; fi92 < 20; fi92++ {
                        v4 = ((((-2.25 - v4) - (1e10 - 1e10)) - v4) - v4)
                        var v93 int = genFn1()
                        v1 = ("hello" + ("a" + (v1 + v1)))
                        _ = fi92
                        _ = v93
                    }
                    _ = fi91
                }
            }
        }
        _ = v89
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi94 := 0; fi94 < 16; fi94++ {
        v95 := GenS11{F0: (v76 / ((62) | 1))}
        _ = fi94
        _ = v95
    }
    fmt.Println(v2, v5)
    v3 = append(v3, genFn1())
    if (false || false) {
        v96 := GenS12{F0: ("hello" + ("" + v1))}
        for fi97 := 0; fi97 < 14; fi97++ {
            for fi98 := 0; fi98 < 16; fi98++ {
                if false {
                    v5 = (v5 ^ (v5 ^ v5))
                    fmt.Println(v77, v76)
                    v4 -= v4
                }
                _ = fi98
            }
            _ = fi97
        }
        _ = v96
    } else {
        if ((v77 - (v77 % (((v77 ^ v77)) | 1))) == int64(v5)) {
            v1 = (v1 + v1)
            for fi99 := 0; fi99 < 7; fi99++ {
                if ((fi99 % (((fi99 * 571)) | 1)) < genFn1()) {
                    var v100 int64 = -1
                    v101 := GenS2{F0: -595, F1: 1, F2: (v76 == 42)}
                    _ = v100
                    _ = v101
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v2 += (v2 - (v2 + 0))
                    if false {
                        var v102 string = "\n\t"
                        _ = v102
                    }
                }
                _ = fi99
            }
        } else {
            switch -739 {
            case 2147483647:
                fmt.Println(v77, v1)
            case 0:
                var v103 float32 = (3.14159 / (3.14159 / 1.0))
                _ = v103
            case 128:
                if false {
                    if true {
                        v77 = (v77 - v77)
                        v4 = (v4 + 1.5)
                    } else {
                        v6[114] = "\n\t"
                        var v104 uint = v76
                        _ = v104
                    }
                    var v105 int32 = int32(v2)
                    _ = v105
                } else {
                    var v106 float64 = -2.25
                    _ = v106
                }
            default:
                fmt.Println(v77)
            }
            v6[uint32(v2)] = (v1 + "\n\t")
            v2 -= (v2 * (v2 / ((65535) | 1)))
        }
    }
    if (true || false) {
        v3 = append(v3, 42)
        if (v5 == -2147483648) {
            var v107 float64 = ((1.5 + float64(v4)) * 0.1)
            for fi108 := 0; fi108 < 7; fi108++ {
                _ = len(v6)
                for fi109 := 0; fi109 < 4; fi109++ {
                    v110 := GenS13{F0: v77, F1: 372, F2: 733}
                    fmt.Println(v76, fi109)
                    _ = fi109
                    _ = v110
                }
                _ = fi108
            }
            v3 = append(v3, 218)
            _ = v107
        } else {
            if (623 < 837) {
                fmt.Println(v76)
                fmt.Println(v5)
                var v111 int16 = -629
                _ = v111
            }
            for fi112 := 0; fi112 < 1; fi112++ {
                var v113 uint64 = uint64(v2)
                var v114 int32 = v5
                for fi115 := 0; fi115 < 4; fi115++ {
                    if true {
                        _ = len(v3)
                        v114 = 1
                    }
                    var v116 int16 = 32767
                    _ = fi115
                    _ = v116
                }
                _ = fi112
                _ = v113
                _ = v114
            }
            if false {
                v76 -= v76
                v117 := GenS4{F0: (v77 + (v77 - v77)), F1: (^(v5))}
                v118 := GenS8{F0: -45, F1: 0, F2: v1}
                _ = v117
                _ = v118
            } else {
                if (!(((v76 * v76) <= (v76 | 18446744073709551615)) && false)) {
                    var v119 int16 = 32767
                    v120 := GenS3{F0: (true || ((3.14159 == 1e10) && (-2.25 > 1.5)))}
                    _ = v119
                    _ = v120
                } else {
                    v121 := GenS1{F0: -128, F1: 42}
                    _ = v121
                }
                _ = len(v6)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v77
    _ = v78
}
