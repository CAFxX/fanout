package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 string; F2 int64 }
type GenS6 struct { F0 int16; F1 int; F2 int64 }
type GenS7 struct { F0 int8; F1 int8; F2 uint }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int8; F1 uint; F2 uint8 }
type GenS10 struct { F0 int64; F1 uint32; F2 int16 }
type GenS11 struct { F0 int8; F1 uint16; F2 int64 }
func genFn0(p0 uint32) int64 {
    if (true || false) {
        v5 := []int{4294967295, -9223372036854775808}
        for fi6 := 0; fi6 < 15; fi6++ {
            for fi7 := 0; fi7 < 2; fi7++ {
                p0 = 348
                var v8 uint = 1
                _ = fi7
                _ = v8
            }
            var v9 int32 = 2147483647
            _ = fi6
            _ = v9
        }
        _ = v5
    } else {
        if (((0.0 * (0.1 * 0.0)) * 0.0) > 3.14159) {
            var v10 uint8 = 255
            v11 := GenS1{F0: (p0 * (p0 << ((0) & 63)))}
            v12 := GenS2{F0: 17, F1: 65535, F2: ("a" + ("" + "foo bar"))}
            _ = v10
            _ = v11
            _ = v12
        }
    }
    if false {
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := GenS2{F0: 42, F1: 42, F2: "a"}
            switch ("foo bar" + "hello") {
            case "hello":
                for fi15 := 0; fi15 < 18; fi15++ {
                    p0 += 42
                    _ = fi15
                }
            case "":
                switch fi13 {
                case 42:
                    for fi16 := 0; fi16 < 16; fi16++ {
                        v17 := []int{342, -1, (fi13 + (fi13 ^ fi13)), -9223372036854775808}
                        var v18 uint = 601
                        if false {
                            fmt.Println(fi16, p0)
                            if (-127 >= 0) {
                                switch ("\n\t" + ("foo bar" + ("\n\t" + "\n\t"))) {
                                case "":
                                    if (true || true) {
                                        p0 += p0
                                    } else {
                                        v18 = v18
                                        v18 = v18
                                        v19 := GenS3{F0: true}
                                        v20 := GenS3{F0: ((!true) || true)}
                                        if true {
                                            if len(v17) > 0 {
                                                v17[0] = fi13
                                            }
                                            p0 += p0
                                            if (!(0.0 < 0.0)) {
                                                fmt.Println(p0, fi13, v18)
                                            }
                                        } else {
                                            v18 = (v18 ^ v18)
                                        }
                                        v21 := GenS4{F0: 32767}
                                        v18 -= (^(v18))
                                        if true {
                                            var v22 uint8 = 42
                                            var v23 int = (fi16 >> (((^(v18))) & 63))
                                            switch "a" {
                                            case "hello":
                                                p0 = (p0 | 0)
                                            }
                                            _ = v22
                                            _ = v23
                                        }
                                        _ = v19
                                        _ = v20
                                        _ = v21
                                    }
                                }
                                _ = len(v17)
                            }
                            var v24 uint32 = p0
                            _ = v24
                        }
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                case 9223372036854775807:
                    v25 := GenS5{F0: 136, F1: "foo bar", F2: 9223372036854775807}
                    _ = v25
                case -879:
                    var v26 int8 = 42
                    _ = v26
                default:
                    switch fi13 {
                    case -219:
                        var v27 uint = 255
                        _ = v27
                    case 42:
                        v28 := []int{}
                        _ = v28
                    }
                }
            case "\n\t":
                var v29 uint16 = 0
                _ = v29
            default:
                if true {
                    fmt.Println(fi13)
                    var v30 uint64 = 0
                    for fi31 := 0; fi31 < 9; fi31++ {
                        if (1.5 >= 3.14159) {
                            var v32 int64 = 0
                            v33 := []int{(-(fi13))}
                            _ = v32
                            _ = v33
                        } else {
                            switch (fi13 & int(v30)) {
                            case -618:
                                v30 = v30
                            case -781:
                                v30 = 2
                            default:
                                var v34 uint32 = (p0 + p0)
                                _ = v34
                            }
                        }
                        if ("\n\t" > "hello") {
                            var v35 uint = 602
                            var v36 uint16 = 1
                            _ = v35
                            _ = v36
                        }
                        fmt.Println(v30, fi13)
                        _ = fi31
                    }
                    _ = v30
                } else {
                    for fi37 := 0; fi37 < 19; fi37++ {
                        p0 = p0
                        v38 := []int{(fi13 - fi37), fi13, (fi37 << ((65535) & 63))}
                        fmt.Println(fi37, fi13, p0)
                        _ = fi37
                        _ = v38
                    }
                    v39 := []int{fi13, (fi13 % ((fi13) | 1)), (fi13 << ((255) & 63)), (fi13 * (fi13 - (fi13 % ((fi13) | 1))))}
                    _ = v39
                }
            }
            _ = fi13
            _ = v14
        }
    }
    v40 := []int{1, 797, 1}
    if (true || true) {
        if ((!(32767 == 884)) || ("\n\t" < "a")) {
            fmt.Println(p0)
            var v41 int32 = 576
            _ = v41
        } else {
            var v42 int32 = -966
            switch "foo bar" {
            case "\n\t":
                var v43 int32 = -342
                _ = v43
            case "":
                switch 32767 {
                case 42:
                    for fi44 := 0; fi44 < 7; fi44++ {
                        var v45 uint32 = (p0 / ((210) | 1))
                        _ = fi44
                        _ = v45
                    }
                }
            case "hello":
                switch ("hello" + "a") {
                case "a":
                    p0 = (p0 >> ((375) & 63))
                case "hello":
                    p0 -= p0
                case "foo bar":
                    if (1 == 184) {
                        fmt.Println(p0)
                        if len(v40) >= 2 {
                            v40 = v40[1:len(v40)]
                        }
                    }
                }
            default:
                v42 -= (v42 << ((857) & 63))
            }
            switch -358 {
            case 249:
                var v46 uint32 = p0
                _ = v46
            case -9223372036854775808:
                fmt.Println(v42, p0)
            case -954:
                if (4294967295 >= 42) {
                    var v47 uint8 = 108
                    switch "hello" {
                    case "a":
                        var v48 int16 = 42
                        _ = v48
                    case "foo bar":
                        fmt.Println(v42, v42)
                    default:
                        v49 := GenS6{F0: 0, F1: 128, F2: 968}
                        _ = v49
                    }
                    _ = v47
                }
            default:
                var v50 uint16 = 42
                _ = v50
            }
            _ = v42
        }
        if false {
            var v51 uint32 = p0
            _ = v51
        }
    }
    _ = p0
    _ = v40
    return -529
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = ((0.1 * 3.14159) + 1e10)
    var v3 bool = (!true)
    v4 := make(map[string]uint32)
    var v52 uint16 = 255
    v53 := v1[2147483647]
    if ((959 <= -897) && true) {
        v54 := GenS3{F0: (("a" + ("" + ("hello" + ("a" + "a")))) < ((("hello" + "a") + ("" + ("a" + "hello"))) + ""))}
        switch 32767 {
        case -504:
            if false {
                for fi55 := 0; fi55 < 7; fi55++ {
                    fmt.Println(v52, v3, v52)
                    for fi56 := 0; fi56 < 20; fi56++ {
                        switch fi56 {
                        case -1:
                            v52 = 255
                        default:
                            if ("hello" < "foo bar") {
                                var v57 bool = ((!(("" < ("" + "a")) || v3)) || true)
                                v2 += -2.25
                                fmt.Println(fi55, fi56)
                                _ = v57
                            }
                        }
                        v58 := GenS2{F0: 1, F1: 65535, F2: "hello"}
                        if true {
                            var v59 int16 = -255
                            fmt.Println(v3, v53, v52)
                            v59 -= (v59 / ((v59) | 1))
                            _ = v59
                        } else {
                            delete(v4, (("a" + "foo bar") + "hello"))
                            var v60 uint16 = v52
                            if (127 <= -1) {
                                fmt.Println(fi56, v3, v3)
                                var v61 int64 = 807
                                v52 += (v52 ^ v52)
                                _ = v61
                            } else {
                                v62 := GenS6{F0: 42, F1: fi55, F2: 42}
                                if (1 == -32768) {
                                    v52 = 32767
                                    var v63 bool = v53
                                    _ = v63
                                } else {
                                    if (!true) {
                                        if v3 {
                                            v53 = true
                                        } else {
                                            if (2 < 456) {
                                                v64 := GenS2{F0: 98, F1: 454, F2: (("foo bar" + "") + "\n\t")}
                                                switch "" {
                                                case "":
                                                    var v65 float64 = v2
                                                    _ = v65
                                                default:
                                                    v53 = (-1 != 0)
                                                }
                                                var v66 int64 = -1
                                                _ = v64
                                                _ = v66
                                            }
                                        }
                                    } else {
                                        if v53 {
                                            _ = len(v1)
                                        }
                                        switch fi55 {
                                        case 32768:
                                            v52 = v60
                                            var v67 uint64 = 470
                                            var v68 float64 = (v2 / float64(3.14159))
                                            _ = v67
                                            _ = v68
                                        case 565:
                                            v3 = (-886 < -32768)
                                        case 42:
                                            v2 = 0.0
                                        }
                                    }
                                    v69 := GenS5{F0: 99, F1: "\n\t", F2: genFn0(651)}
                                    v53 = (32767 <= v60)
                                    _ = v69
                                }
                                if v3 {
                                    v52 = v60
                                    v70 := []int{(-(fi55)), fi55, fi55, -1}
                                    var v71 uint64 = 18446744073709551615
                                    _ = v70
                                    _ = v71
                                }
                                switch "a" {
                                case "":
                                    var v72 bool = false
                                    _ = v72
                                case "a":
                                    v3 = v53
                                case "\n\t":
                                    fmt.Println(fi55, v52, v52)
                                default:
                                    delete(v4, "foo bar")
                                }
                                var v73 string = "foo bar"
                                _ = v62
                                _ = v73
                            }
                            _ = v60
                        }
                        _ = fi56
                        _ = v58
                    }
                    _ = fi55
                }
                var v74 uint64 = 1
                _ = v74
            }
        }
        _ = v54
    } else {
        if false {
            if (((v2 + (((float64(float32(0.0)) - (0.1 / (v2 + 1e10))) * ((float64((0.1 + 3.14159)) + float64(float32(v2))) - float64(float32(v2)))) * v2)) < v2) || (!(128 > 42))) {
                for fi75 := 0; fi75 < 9; fi75++ {
                    if v53 {
                        var v76 float64 = v2
                        if true {
                            v77 := GenS4{F0: -1}
                            var v78 float64 = v2
                            v79 := []int{127}
                            _ = v77
                            _ = v78
                            _ = v79
                        }
                        if true {
                            for fi80 := 0; fi80 < 8; fi80++ {
                                var v81 string = ""
                                v76 += (v76 - float64(float32(1.5)))
                                _ = fi80
                                _ = v81
                            }
                            var v82 uint = 18446744073709551615
                            _ = v82
                        }
                        _ = v76
                    } else {
                        switch fi75 {
                        case 508:
                            switch fi75 {
                            case 4294967295:
                                fmt.Println(fi75)
                            case -2147483648:
                                _ = len(v4)
                            default:
                                fmt.Println(fi75, v53)
                            }
                        case 9223372036854775807:
                            var v83 string = ("hello" + "hello")
                            _ = v83
                        }
                        switch "" {
                        case "foo bar":
                            for fi84 := 0; fi84 < 6; fi84++ {
                                v3 = v53
                                v85 := GenS2{F0: -1, F1: 255, F2: "foo bar"}
                                v86 := GenS4{F0: -833}
                                switch -840 {
                                case 9223372036854775807:
                                    v87 := GenS3{F0: true}
                                    _ = v87
                                case 155:
                                    var v88 int8 = 127
                                    _ = v88
                                case -759:
                                    switch (fi75 >> ((931) & 63)) {
                                    case -9223372036854775808:
                                        v89 := GenS7{F0: 1, F1: 56, F2: 378}
                                        _ = v89
                                    case 804:
                                        var v90 int64 = -1
                                        _ = v90
                                    case 32767:
                                        v91 := GenS1{F0: 42}
                                        _ = v91
                                    }
                                }
                                _ = fi84
                                _ = v85
                                _ = v86
                            }
                        case "hello":
                            v53 = (-32768 != 42)
                        default:
                            v92 := GenS8{F0: (v52 % ((0) | 1))}
                            _ = v92
                        }
                        v93 := GenS9{F0: -128, F1: 255, F2: 181}
                        _ = v93
                    }
                    for fi94 := 0; fi94 < 1; fi94++ {
                        v53 = (((0 == -953) || (-2.25 != 3.14159)) || v53)
                        var v95 int = fi94
                        var v96 uint64 = 514
                        v95 += fi94
                        var v97 bool = ((-1 < -607) || ((float64((float32(v2) * float32((float64(3.14159) / float64(((1e10 / (float32(v2) - 0.0)) * float32(v2))))))) / v2) <= v2))
                        _ = fi94
                        _ = v95
                        _ = v96
                        _ = v97
                    }
                    v98 := []int{}
                    _ = fi75
                    _ = v98
                }
                var v99 bool = (!(!(v2 > v2)))
                _ = v99
            } else {
                if v3 {
                    v3 = true
                } else {
                    var v100 int32 = 1
                    v101 := GenS2{F0: 100, F1: 1, F2: (("hello" + "\n\t") + ("" + ((((("" + "foo bar") + ("" + "hello")) + "foo bar") + "") + "")))}
                    v100 -= v100
                    _ = v100
                    _ = v101
                }
            }
            v102 := GenS10{F0: 42, F1: 42, F2: -1}
            switch 804 {
            case 32767:
                var v103 bool = true
                _ = v103
            case 1:
                v2 = (v2 / 1.0)
            case 127:
                fmt.Println(v53, v2)
            }
            _ = v102
        } else {
            switch 42 {
            case 861:
                fmt.Println(v52)
            }
            if true {
                var v104 string = "foo bar"
                v105 := GenS3{F0: v3}
                _ = v104
                _ = v105
            } else {
                if false {
                    var v106 uint8 = 1
                    var v107 int64 = 42
                    for fi108 := 0; fi108 < 8; fi108++ {
                        delete(v1, 2147483648)
                        for fi109 := 0; fi109 < 16; fi109++ {
                            switch fi109 {
                            case 0:
                                var v110 string = "hello"
                                _ = v110
                            case -2147483648:
                                v106 += v106
                            default:
                                v3 = ((fi109 >> ((1) & 63)) == (-(fi108)))
                            }
                            v111 := []int{fi108}
                            _ = fi109
                            _ = v111
                        }
                        _ = fi108
                    }
                    _ = v106
                    _ = v107
                }
            }
        }
        var v112 float64 = 3.14159
        _ = v112
    }
    for fi113 := 0; fi113 < 20; fi113++ {
        if (65535 <= 342) {
            for fi114 := 0; fi114 < 13; fi114++ {
                v115 := GenS8{F0: v52}
                _ = fi114
                _ = v115
            }
            var v116 uint8 = 127
            _ = v116
        }
        _ = fi113
    }
    if v3 {
        v3 = (32767 == -1)
        if (!v3) {
            var v117 uint32 = 2147483648
            var v118 uint16 = v52
            _ = v117
            _ = v118
        } else {
            v2 = 1e10
            v119 := GenS11{F0: 42, F1: 0, F2: -1}
            v120 := GenS11{F0: -1, F1: v52, F2: -380}
            _ = v119
            _ = v120
        }
        if v53 {
            switch "hello" {
            case "a":
                fmt.Println(v2)
            case "\n\t":
                _ = len(v1)
            default:
                for fi121 := 0; fi121 < 20; fi121++ {
                    if (("a" + "hello") < ("foo bar" + "a")) {
                        var v122 uint16 = v52
                        if v3 {
                            fmt.Println(v52, fi121, fi121)
                            fmt.Println(v3, v2)
                        }
                        var v123 string = (("\n\t" + "hello") + ("" + "foo bar"))
                        _ = v122
                        _ = v123
                    }
                    v3 = ((1.5 >= 3.14159) || (1 != 127))
                    _ = fi121
                }
            }
        }
    }
    v124 := GenS1{F0: 0}
    if v53 {
        v53 = (271 <= 694)
        var v125 float64 = (((v2 - float64(1.5)) / (1e10 / float64(1.5))) / (v2 / 1.0))
        if (v3 || (false || false)) {
            v53 = true
        } else {
            fmt.Println(v2)
        }
        _ = v125
    } else {
        fmt.Println(v3, v52)
    }
    fmt.Println(v52)
    v126 := []int{}
    v53 = (9223372036854775807 >= 9223372036854775807)
    if len(v126) >= 2 {
        v126 = v126[1:len(v126)]
    }
    for fi127 := 0; fi127 < 7; fi127++ {
        v3 = false
        _ = fi127
    }
    if (!(65535 >= 4294967295)) {
        var v128 float64 = ((0.1 / (1.5 / v2)) - v2)
        v53 = v53
        v3 = v53
        _ = v128
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v126:", len(v126), v126)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v124
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
}
