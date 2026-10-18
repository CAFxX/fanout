package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint }
type GenS4 struct { F0 uint32; F1 uint; F2 int32 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint; F1 uint }
type GenS10 struct { F0 string; F1 uint8 }
type GenS11 struct { F0 int; F1 bool }
func genFn0() int16 {
    if false {
        v7 := GenS1{F0: 4294967295, F1: (true || true)}
        var v8 bool = false
        _ = v7
        _ = v8
    } else {
        if (245 != 128) {
            var v9 int32 = -521
            _ = v9
        } else {
            if false {
                for fi10 := 0; fi10 < 16; fi10++ {
                    v11 := []int{fi10, 127, fi10}
                    if len(v11) > 0 {
                        v11[0] = 1
                    }
                    _ = fi10
                    _ = v11
                }
                var v12 int = 0
                if false {
                    v12 = 128
                    var v13 int8 = 127
                    v13 += v13
                    _ = v13
                }
                _ = v12
            } else {
                v14 := GenS2{F0: 520, F1: 2}
                _ = v14
            }
        }
        if false {
            var v15 bool = false
            _ = v15
        } else {
            var v16 float32 = (1e10 * 1e10)
            for fi17 := 0; fi17 < 20; fi17++ {
                v18 := []int{}
                v18 = append(v18, -9223372036854775808)
                v19 := GenS1{F0: 2, F1: (18 == 0)}
                _ = fi17
                _ = v18
                _ = v19
            }
            for fi20 := 0; fi20 < 7; fi20++ {
                var v21 uint64 = 154
                if true {
                    v21 += v21
                    switch (("foo bar" + "") + "a") {
                    case "hello":
                        if false {
                            if (2 == 1) {
                                var v22 float64 = (1e10 / (1.5 + 3.14159))
                                v23 := GenS2{F0: 1, F1: 255}
                                _ = v22
                                _ = v23
                            }
                            switch (-(fi20)) {
                            case 203:
                                v16 -= ((((v16 * 1e10) / 1.0) + (0.1 + (3.14159 + 0.0))) * v16)
                            default:
                                v16 = -2.25
                            }
                        }
                    case "":
                        for fi24 := 0; fi24 < 1; fi24++ {
                            v25 := []int{-2147483648, 128}
                            v21 = v21
                            switch ("\n\t" + (("hello" + "hello") + ((("foo bar" + "\n\t") + "foo bar") + "a"))) {
                            case "a":
                                var v26 uint8 = 138
                                _ = v26
                            case "":
                                v27 := GenS1{F0: 42, F1: true}
                                _ = v27
                            case "foo bar":
                                var v28 bool = true
                                _ = v28
                            default:
                                switch (fi20 & 9223372036854775807) {
                                case -747:
                                    switch ("foo bar" + "foo bar") {
                                    case "":
                                        if ("a" <= "\n\t") {
                                            v21 += 359
                                            var v29 uint8 = 7
                                            v16 = 0.1
                                            if len(v25) >= 2 {
                                                v25 = v25[1:len(v25)]
                                            }
                                            _ = v29
                                        } else {
                                            v25 = append(v25, fi24)
                                        }
                                    case "foo bar":
                                        if (72 >= 211) {
                                            var v30 uint8 = 254
                                            if (false && ((^(v30)) == v30)) {
                                                var v31 uint8 = v30
                                                fmt.Println(v21, v21)
                                                switch fi24 {
                                                case -1:
                                                    var v32 float64 = -2.25
                                                    _ = v32
                                                case 0:
                                                    switch ("foo bar" + (("" + "a") + ("\n\t" + "foo bar"))) {
                                                    case "hello":
                                                        if len(v25) > 0 {
                                                            v25[0] = fi24
                                                        }
                                                    }
                                                case 128:
                                                    switch "hello" {
                                                    case "":
                                                        fmt.Println(v21, v16, v16)
                                                    case "\n\t":
                                                        v30 += 113
                                                    case "hello":
                                                        var v33 uint64 = (v21 ^ 18446744073709551615)
                                                        _ = v33
                                                    }
                                                }
                                                var v34 uint16 = 1
                                                _ = v31
                                                _ = v34
                                            }
                                            v21 = (v21 % ((42) | 1))
                                            _ = v30
                                        } else {
                                            _ = len(v25)
                                            v21 = 18446744073709551615
                                            var v35 int16 = int16(fi20)
                                            _ = v35
                                        }
                                    case "hello":
                                        var v36 float32 = v16
                                        _ = v36
                                    }
                                }
                                var v37 string = ("\n\t" + ("" + ("hello" + "a")))
                                _ = v37
                            }
                            _ = fi24
                            _ = v25
                        }
                    }
                    v16 -= (v16 + 1.5)
                }
                if true {
                    v21 += v21
                    v38 := GenS3{F0: 127, F1: 0, F2: 42}
                    v39 := GenS3{F0: 127, F1: 128, F2: 1}
                    _ = v38
                    _ = v39
                } else {
                    var v40 int = fi20
                    var v41 float64 = 0.1
                    _ = v40
                    _ = v41
                }
                _ = fi20
                _ = v21
            }
            _ = v16
        }
        for fi42 := 0; fi42 < 15; fi42++ {
            if ((3.14159 - 0.1) >= 1e10) {
                var v43 int8 = 42
                _ = v43
            } else {
                var v44 string = ""
                for fi45 := 0; fi45 < 16; fi45++ {
                    v46 := []int{819, 127, (-(fi45))}
                    var v47 string = v44
                    var v48 uint64 = 18446744073709551615
                    _ = fi45
                    _ = v46
                    _ = v47
                    _ = v48
                }
                _ = v44
            }
            var v49 string = "\n\t"
            v50 := GenS3{F0: 127, F1: 128, F2: 1}
            _ = fi42
            _ = v49
            _ = v50
        }
    }
    for fi51 := 0; fi51 < 15; fi51++ {
        var v52 int16 = -32768
        v53 := GenS1{F0: 2147483648, F1: false}
        _ = fi51
        _ = v52
        _ = v53
    }
    var v54 uint8 = 1
    if ((false && false) || false) {
        switch "foo bar" {
        case "foo bar":
            for fi55 := 0; fi55 < 16; fi55++ {
                v54 = v54
                if true {
                    if ((true && true) && (9223372036854775807 == (fi55 ^ fi55))) {
                        var v56 float32 = 3.14159
                        for fi57 := 0; fi57 < 6; fi57++ {
                            if (((-128 != 127) || (733 == 65535)) && true) {
                                if (!(18446744073709551615 != 255)) {
                                    var v58 int32 = 0
                                    if ((-(fi55)) < (fi55 & fi57)) {
                                        v59 := []int{fi57, 127, fi55, 2147483647}
                                        if (!(false || false)) {
                                            var v60 int64 = -1
                                            switch "a" {
                                            case "hello":
                                                v56 = v56
                                            }
                                            v56 += (v56 / (v56 / (v56 / (v56 * -2.25))))
                                            _ = v60
                                        }
                                        v61 := GenS2{F0: 42, F1: 255}
                                        if (!true) {
                                            var v62 int64 = -987
                                            _ = v62
                                        } else {
                                            v59 = append(v59, (^(fi55)))
                                            v56 = (v56 + v56)
                                        }
                                        var v63 bool = (9223372036854775807 < -1)
                                        v56 = (v56 + 0.1)
                                        var v64 float64 = 0.1
                                        _ = v59
                                        _ = v61
                                        _ = v63
                                        _ = v64
                                    }
                                    _ = v58
                                } else {
                                    fmt.Println(fi55)
                                    if (float64((v56 * ((v56 - v56) - v56))) < 1.5) {
                                        v65 := GenS4{F0: 4294967295, F1: 42, F2: 1}
                                        if false {
                                            v54 += (^(v54))
                                            fmt.Println(v56, fi57, fi55)
                                            if (65535 >= 2) {
                                                if false {
                                                    if (284 >= 210) {
                                                        var v66 uint64 = 868
                                                        v67 := GenS1{F0: 1, F1: true}
                                                        if true {
                                                            fmt.Println(fi55, v66)
                                                            fmt.Println(v54)
                                                        } else {
                                                            if false {
                                                                fmt.Println(fi55, v54)
                                                            } else {
                                                                var v68 uint8 = (v54 << ((2) & 63))
                                                                v69 := GenS5{F0: 154, F1: (2 == 65535)}
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                        }
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                }
                                                if true {
                                                    var v70 int8 = -1
                                                    var v71 int16 = 954
                                                    _ = v70
                                                    _ = v71
                                                } else {
                                                    v56 = (-2.25 * v56)
                                                    var v72 uint32 = 2
                                                    _ = v72
                                                }
                                            } else {
                                                var v73 uint = 330
                                                var v74 string = ""
                                                var v75 uint = v73
                                                v56 -= (v56 / v56)
                                                v76 := []int{128, fi55, int(v54)}
                                                _ = v73
                                                _ = v74
                                                _ = v75
                                                _ = v76
                                            }
                                            var v77 uint64 = 4294967295
                                            _ = v77
                                        }
                                        if (!false) {
                                            switch (^(fi57)) {
                                            case -1:
                                                if true {
                                                    var v78 string = ("a" + "a")
                                                    _ = v78
                                                } else {
                                                    var v79 int16 = 32767
                                                    _ = v79
                                                }
                                                if true {
                                                    fmt.Println(v56)
                                                }
                                                var v80 int = -112
                                                _ = v80
                                            case 4294967295:
                                                var v81 bool = (0 != 4294967295)
                                                _ = v81
                                            }
                                        }
                                        _ = v65
                                    } else {
                                        if true {
                                            v82 := GenS2{F0: 426, F1: (v54 & (^(v54)))}
                                            fmt.Println(fi55)
                                            if true {
                                                v83 := []int{(^(fi57)), fi55}
                                                fmt.Println(fi55)
                                                v56 = 3.14159
                                                _ = v83
                                            }
                                            _ = v82
                                        }
                                        v84 := GenS3{F0: (v54 ^ 210), F1: v54, F2: 4294967295}
                                        if false {
                                            switch -210 {
                                            case 127:
                                                switch fi55 {
                                                case 127:
                                                    if (!(65535 == 308)) {
                                                        if false {
                                                            fmt.Println(v54, fi55)
                                                            var v85 int64 = -1
                                                            _ = v85
                                                        } else {
                                                            v86 := GenS2{F0: 255, F1: (v54 % ((v54) | 1))}
                                                            v87 := []int{}
                                                            v56 = 0.0
                                                            var v88 int8 = -1
                                                            v54 += (v54 % ((v54) | 1))
                                                            v89 := GenS6{F0: -1}
                                                            var v90 float64 = 1.5
                                                            if ((-2.25 - float32(1.5)) == v56) {
                                                                v56 = v56
                                                                fmt.Println(v54, v54, v90)
                                                                if (true && (0 >= 42)) {
                                                                    if false {
                                                                        var v91 int8 = (v88 ^ v88)
                                                                        _ = v91
                                                                    } else {
                                                                        var v92 int64 = 42
                                                                        v90 = float64((0.0 - v56))
                                                                        v93 := GenS3{F0: v54, F1: 128, F2: 4294967295}
                                                                        _ = v92
                                                                        _ = v93
                                                                    }
                                                                }
                                                            }
                                                            _ = v86
                                                            _ = v87
                                                            _ = v88
                                                            _ = v89
                                                            _ = v90
                                                        }
                                                        v54 = (v54 / ((v54) | 1))
                                                    } else {
                                                        v54 += (v54 & v54)
                                                    }
                                                }
                                            default:
                                                v56 += 3.14159
                                            }
                                            v94 := []int{(fi55 % ((int(v54)) | 1)), (^(fi55)), int(v54)}
                                            _ = v94
                                        } else {
                                            v95 := GenS6{F0: 817}
                                            _ = v95
                                        }
                                        _ = v84
                                    }
                                    switch (("a" + (((("\n\t" + "a") + "hello") + "hello") + ("\n\t" + ""))) + (("\n\t" + "foo bar") + "hello")) {
                                    case "a":
                                        var v96 uint64 = 0
                                        _ = v96
                                    default:
                                        fmt.Println(v56, fi55, v54)
                                    }
                                    fmt.Println(fi55)
                                }
                                v97 := []int{(^(fi57)), fi55, 767, (fi55 << ((65535) & 63))}
                                if false {
                                    v54 = 123
                                }
                                _ = v97
                            }
                            v56 -= (1e10 + v56)
                            v54 += (v54 ^ 1)
                            _ = fi57
                        }
                        v98 := GenS5{F0: 2, F1: true}
                        _ = v56
                        _ = v98
                    }
                    if (1 > -114) {
                        v54 += v54
                    }
                    var v99 uint8 = (v54 >> ((42) & 63))
                    _ = v99
                } else {
                    v100 := []int{(fi55 ^ (fi55 / ((fi55) | 1)))}
                    v54 = v54
                    _ = v100
                }
                v54 -= v54
                _ = fi55
            }
        default:
            v101 := []int{-1}
            _ = v101
        }
        for fi102 := 0; fi102 < 1; fi102++ {
            v103 := []int{(fi102 >> ((0) & 63)), 42, (fi102 | (fi102 >> ((65535) & 63)))}
            _ = fi102
            _ = v103
        }
        v54 = 153
    }
    v54 += (v54 - (v54 ^ v54))
    _ = v54
    return 415
}
func genFn1() int64 {
    var v104 int = 0
    if false {
        v105 := GenS7{F0: v104, F1: 1}
        switch (v104 << ((42) & 63)) {
        case -2147483648:
            v104 += v104
        case 32768:
            v106 := GenS2{F0: 42, F1: 42}
            _ = v106
        default:
            var v107 uint = 255
            _ = v107
        }
        switch "foo bar" {
        case "foo bar":
            v104 -= v104
        }
        _ = v105
    }
    for fi108 := 0; fi108 < 19; fi108++ {
        v109 := GenS1{F0: 606, F1: true}
        v110 := GenS8{F0: 42}
        _ = fi108
        _ = v109
        _ = v110
    }
    fmt.Println(v104)
    _ = v104
    return 9223372036854775807
}

func main() {
    v1 := []uint32{1}
    var v2 uint16 = 2
    v3 := make(map[string]uint32)
    var v4 int16 = 1
    v5 := []int{1}
    v6 := make(map[int]bool)
    v4 = v4
    if (false && (v4 != v4)) {
        fmt.Println(v4, v2)
        v111 := GenS9{F0: 1, F1: 42}
        _ = v111
    } else {
        v112 := GenS10{F0: "hello", F1: 255}
        v113 := GenS1{F0: 42, F1: ((true || (-9223372036854775808 < 874)) && (true && (65535 == 469)))}
        _ = v112
        _ = v113
    }
    for fi114 := 0; fi114 < 20; fi114++ {
        if true {
            fmt.Println(v2, v2, v2)
            switch (fi114 + (fi114 ^ (^(fi114)))) {
            case 32767:
                var v115 int64 = -364
                _ = v115
            case 286:
                switch ("" + "\n\t") {
                case "":
                    _ = len(v5)
                default:
                    if true {
                        var v116 uint16 = (v2 << ((2) & 63))
                        _ = v116
                    }
                }
            case 4294967295:
                v4 -= (-(v4))
            default:
                fmt.Println(v2)
            }
            v2 = (^(v2))
        } else {
            for fi117 := 0; fi117 < 3; fi117++ {
                v118 := GenS4{F0: 2147483648, F1: 4294967295, F2: 440}
                v119 := GenS3{F0: 1, F1: 255, F2: 948}
                _ = fi117
                _ = v118
                _ = v119
            }
        }
        if false {
            fmt.Println(fi114)
        }
        _ = fi114
    }
    switch "foo bar" {
    case "\n\t":
        v120 := GenS1{F0: 2147483647, F1: true}
        _ = v120
    case "foo bar":
        switch 863 {
        case 127:
            v2 += (v2 & (v2 % ((v2) | 1)))
        case 42:
            _ = len(v5)
        }
    case "a":
        v121 := GenS11{F0: -9223372036854775808, F1: ("a" >= "")}
        _ = v121
    default:
        for fi122 := 0; fi122 < 1; fi122++ {
            var v123 string = "hello"
            _ = fi122
            _ = v123
        }
    }
    v3[(("foo bar" + "hello") + "\n\t")] = 42
    v124 := GenS8{F0: -1}
    _ = len(v1)
    fmt.Println(v4, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v124
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
