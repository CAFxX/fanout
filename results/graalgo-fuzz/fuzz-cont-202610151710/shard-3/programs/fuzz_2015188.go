package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int; F2 int }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 bool; F1 string; F2 int16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 uint }
type GenS10 struct { F0 uint; F1 int16; F2 string }
type GenS11 struct { F0 int16; F1 int }
type GenS12 struct { F0 int; F1 uint; F2 uint8 }
type GenS13 struct { F0 uint8; F1 uint64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 bool; F1 uint64 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 int32; F1 uint16; F2 string }
type GenS18 struct { F0 uint8; F1 uint }
func genFn0() int32 {
    v5 := []int{952, -640, -507}
    var v6 float32 = 1e10
    v6 -= v6
    for fi7 := 0; fi7 < 1; fi7++ {
        switch "\n\t" {
        case "":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "\n\t":
            fmt.Println(fi7, v6)
        case "a":
            for fi8 := 0; fi8 < 5; fi8++ {
                _ = len(v5)
                v6 = ((v6 * v6) / v6)
                _ = fi8
            }
        }
        switch fi7 {
        case 0:
            v9 := GenS1{F0: -81}
            _ = v9
        default:
            switch -9223372036854775808 {
            case 127:
                for fi10 := 0; fi10 < 6; fi10++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi10
                }
            case -1:
                var v11 uint8 = 1
                _ = v11
            case 128:
                v12 := GenS1{F0: -20}
                _ = v12
            default:
                if len(v5) > 0 {
                    v5[0] = fi7
                }
            }
        }
        _ = fi7
    }
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 int8, p1 int32, p2 uint16) bool {
    if false {
        var v13 uint64 = 2
        var v14 int64 = 9223372036854775807
        _ = v13
        _ = v14
    } else {
        p2 -= p2
        var v15 bool = false
        _ = v15
    }
    var v16 string = ((("" + "a") + ((("a" + "") + "\n\t") + "\n\t")) + "\n\t")
    switch 9223372036854775807 {
    case 42:
        v17 := []int{42, 42, 32767, -557}
        _ = v17
    case 127:
        var v18 int = 127
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return ((p0 << ((255) & 63)) >= (p0 & (p0 * p0)))
}

func main() {
    var v1 int = 32767
    v2 := make(map[string]bool)
    var v3 uint32 = 154
    var v4 int8 = 1
    var v19 int64 = -1
    for fi20 := 0; fi20 < 1; fi20++ {
        if (!false) {
            if false {
                fmt.Println(v1, v4, fi20)
                fmt.Println(v1, v3)
                for fi21 := 0; fi21 < 7; fi21++ {
                    if (!true) {
                        if (-20 <= (^(v4))) {
                            fmt.Println(fi21)
                        } else {
                            var v22 uint16 = 255
                            v23 := GenS1{F0: (v4 - 0)}
                            if ("\n\t" != "hello") {
                                var v24 int32 = genFn0()
                                v3 -= (^(v3))
                                if (!true) {
                                    v25 := v2[((("a" + "\n\t") + "") + ("foo bar" + "hello"))]
                                    if ((("a" + ((("" + "a") + "a") + "a")) == "foo bar") && true) {
                                        fmt.Println(v3, v25, v19)
                                        switch fi21 {
                                        case 42:
                                            v26 := []int{}
                                            _ = v26
                                        default:
                                            fmt.Println(v25, fi21, v22)
                                        }
                                        switch ("a" + (("foo bar" + (("hello" + "foo bar") + ("foo bar" + (("" + "foo bar") + "foo bar")))) + "a")) {
                                        case "hello":
                                            switch 2147483647 {
                                            case 539:
                                                var v27 int32 = v24
                                                _ = v27
                                            case 32767:
                                                v28 := GenS2{F0: 0}
                                                _ = v28
                                            }
                                        }
                                        v24 += v24
                                        var v29 float32 = float32(1e10)
                                        _ = v29
                                    } else {
                                        if (-32768 >= 42) {
                                            v24 -= v24
                                            fmt.Println(v1, v4)
                                        } else {
                                            var v30 uint = 148
                                            _ = v30
                                        }
                                        v4 -= 18
                                        v31 := GenS1{F0: (-(v4))}
                                        _ = v31
                                    }
                                    _ = v25
                                }
                                v4 -= (v4 << ((1) & 63))
                                v24 += 554
                                _ = v24
                            }
                            _ = v22
                            _ = v23
                        }
                        v4 -= v4
                        switch int(v19) {
                        case 0:
                            delete(v2, "foo bar")
                        default:
                            switch ("\n\t" + "a") {
                            case "a":
                                v32 := GenS3{F0: 0, F1: int(v3), F2: v1}
                                _ = v32
                            case "hello":
                                v33 := GenS2{F0: -1}
                                _ = v33
                            case "\n\t":
                                fmt.Println(v4, v4)
                            }
                        }
                    }
                    v34 := []int{-278, (^(fi20)), (fi20 >> ((42) & 63))}
                    _ = fi21
                    _ = v34
                }
            } else {
                for fi35 := 0; fi35 < 16; fi35++ {
                    v36 := GenS2{F0: 1}
                    v4 += (^(v4))
                    _ = fi35
                    _ = v36
                }
                switch "hello" {
                case "a":
                    v37 := v2[("a" + (("a" + "hello") + "foo bar"))]
                    _ = v37
                case "":
                    v19 += 9223372036854775807
                case "\n\t":
                    v38 := GenS2{F0: 1000}
                    _ = v38
                }
                var v39 int32 = -2147483648
                _ = v39
            }
        } else {
            fmt.Println(v1)
        }
        v40 := []int{v1, 32768}
        if ((v3 & 2) < 85) {
            var v41 uint8 = 0
            for fi42 := 0; fi42 < 13; fi42++ {
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    var v43 bool = (228 >= 32767)
                    _ = v43
                case "a":
                    v4 = (v4 / ((int8(v3)) | 1))
                    if true {
                        var v44 int64 = (v19 - v19)
                        _ = v44
                    }
                default:
                    switch "" {
                    case "":
                        if true {
                            v45 := GenS1{F0: v4}
                            var v46 int8 = (^(v4))
                            v4 = (^(v46))
                            _ = v45
                            _ = v46
                        }
                    case "hello":
                        if (2147483647 <= -85) {
                            if true {
                                switch "foo bar" {
                                case "foo bar":
                                    var v47 bool = false
                                    _ = v47
                                }
                            }
                            var v48 uint32 = (v3 >> ((42) & 63))
                            _ = v48
                        }
                    }
                }
                _ = fi42
            }
            var v49 int = fi20
            _ = v41
            _ = v49
        }
        _ = fi20
        _ = v40
    }
    v50 := GenS4{F0: 351, F1: (v1 & v1)}
    switch (v1 ^ v1) {
    case 593:
        for fi51 := 0; fi51 < 18; fi51++ {
            v2[""] = (!false)
            delete(v2, "\n\t")
            _ = fi51
        }
    case -654:
        if true {
            v52 := GenS5{F0: 42}
            _ = v52
        } else {
            v53 := GenS5{F0: 4294967295}
            _ = v53
        }
    case 4294967295:
        if (!false) {
            v54 := GenS4{F0: 18446744073709551615, F1: v1}
            if true {
                fmt.Println(v3)
                v55 := GenS6{F0: int16(v3), F1: 42}
                _ = v55
            }
            _ = v54
        } else {
            v56 := GenS7{F0: ((v1 % ((-506) | 1)) != v1), F1: ("a" + ""), F2: 42}
            v1 += -9223372036854775808
            _ = v56
        }
    }
    for fi57 := 0; fi57 < 8; fi57++ {
        var v58 int16 = 489
        var v59 int64 = (v19 % ((v19) | 1))
        if false {
            for fi60 := 0; fi60 < 4; fi60++ {
                v61 := GenS7{F0: false, F1: "a", F2: 1}
                v62 := GenS8{F0: 42}
                _ = fi60
                _ = v61
                _ = v62
            }
            v4 -= (-(v4))
        } else {
            var v63 uint32 = (v3 | (v3 ^ v3))
            _ = v63
        }
        _ = fi57
        _ = v58
        _ = v59
    }
    fmt.Println(v19)
    for fi64 := 0; fi64 < 10; fi64++ {
        v65 := []int{(v1 ^ (v1 / (((v1 & 4294967295)) | 1))), fi64, v1}
        _ = fi64
        _ = v65
    }
    fmt.Println(v3)
    v4 -= (v4 - v4)
    _ = len(v2)
    for fi66 := 0; fi66 < 10; fi66++ {
        for fi67 := 0; fi67 < 14; fi67++ {
            v68 := GenS7{F0: false, F1: ("hello" + "hello"), F2: 1}
            _ = fi67
            _ = v68
        }
        for fi69 := 0; fi69 < 19; fi69++ {
            v70 := []int{301, (fi69 + 1), 2147483647}
            v3 = (v3 >> ((18446744073709551615) & 63))
            if false {
                v4 -= -1
            }
            v71 := GenS9{F0: (v19 | v19), F1: 0}
            switch 42 {
            case 32768:
                if false {
                    fmt.Println(v3)
                    v72 := GenS1{F0: 42}
                    _ = v72
                } else {
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v3, v19)
                    case "\n\t":
                        if ((v4 + (v4 ^ 81)) != v4) {
                            v1 = (^(fi69))
                            if false {
                                v73 := GenS10{F0: 1, F1: -1, F2: ""}
                                v74 := GenS2{F0: -1}
                                v70 = append(v70, v1)
                                var v75 int = 488
                                _ = v73
                                _ = v74
                                _ = v75
                            } else {
                                switch (("hello" + ("foo bar" + "a")) + "\n\t") {
                                case "hello":
                                    fmt.Println(v3)
                                default:
                                    if (((!(false && (!false))) || (false || true)) || true) {
                                        switch "foo bar" {
                                        case "foo bar":
                                            if (float32(float64(1e10)) < float32(1.5)) {
                                                if false {
                                                    var v76 float64 = 0.0
                                                    v3 = v3
                                                    var v77 int = fi66
                                                    _ = v76
                                                    _ = v77
                                                }
                                            } else {
                                                if true {
                                                    fmt.Println(fi69, fi66, v4)
                                                    var v78 uint = 18446744073709551615
                                                    v2[("\n\t" + ("hello" + "foo bar"))] = (!(genFn0() > 42))
                                                    _ = v78
                                                } else {
                                                    var v79 int16 = 0
                                                    _ = v79
                                                }
                                            }
                                        case "hello":
                                            v80 := GenS11{F0: 32767, F1: 42}
                                            _ = v80
                                        default:
                                            var v81 uint32 = 1
                                            _ = v81
                                        }
                                        v82 := GenS9{F0: (v19 % ((v19) | 1)), F1: 4294967295}
                                        _ = v82
                                    }
                                }
                            }
                            var v83 int32 = -1
                            if true {
                                v19 -= (v19 % ((v19) | 1))
                                var v84 uint8 = 42
                                fmt.Println(v3)
                                _ = v84
                            }
                            _ = v83
                        }
                        if false {
                            var v85 int = fi69
                            _ = len(v2)
                            _ = v85
                        }
                        var v86 bool = true
                        _ = v86
                    case "hello":
                        var v87 int8 = (v4 & 127)
                        var v88 bool = true
                        v89 := GenS9{F0: v19, F1: 0}
                        _ = v87
                        _ = v88
                        _ = v89
                    default:
                        v19 -= -875
                    }
                    switch "\n\t" {
                    case "foo bar":
                        if false {
                            v2["hello"] = (false || true)
                            if (!(105 >= 42)) {
                                if true {
                                    if len(v70) > 0 {
                                        v70[0] = -690
                                    }
                                    var v90 uint16 = 538
                                    v90 = v90
                                    if (!(!(true && false))) {
                                        var v91 uint16 = 438
                                        _ = v91
                                    }
                                    _ = v90
                                }
                                v4 -= v4
                            }
                            if true {
                                v19 = v19
                            } else {
                                fmt.Println(v3, v19)
                            }
                        }
                    case "\n\t":
                        v92 := v2[("\n\t" + ("foo bar" + "a"))]
                        fmt.Println(fi66)
                        _ = v92
                    default:
                        _ = len(v2)
                    }
                    var v93 int64 = (-(v19))
                    _ = v93
                }
            }
            _ = fi69
            _ = v70
            _ = v71
        }
        _ = fi66
    }
    for fi94 := 0; fi94 < 9; fi94++ {
        if false {
            var v95 uint = 727
            switch ("a" + "a") {
            case "foo bar":
                for fi96 := 0; fi96 < 3; fi96++ {
                    if false {
                        var v97 uint = (v95 / ((v95) | 1))
                        v98 := GenS12{F0: fi96, F1: (v95 + (v95 % ((v97) | 1))), F2: 127}
                        switch "\n\t" {
                        case "hello":
                            v97 -= (^(v97))
                        case "":
                            if true {
                                v4 += (v4 ^ v4)
                            } else {
                                v99 := GenS13{F0: 1, F1: 0}
                                fmt.Println(v3, fi96)
                                var v100 bool = (v4 <= v4)
                                _ = v100
                                _ = v99
                            }
                        default:
                            v1 -= 9223372036854775807
                        }
                        _ = v97
                        _ = v98
                    }
                    _ = fi96
                }
            case "hello":
                delete(v2, "\n\t")
            }
            _ = v95
        } else {
            if false {
                for fi101 := 0; fi101 < 15; fi101++ {
                    v102 := GenS10{F0: 18446744073709551615, F1: 42, F2: ("\n\t" + (("foo bar" + "") + (("" + "a") + "foo bar")))}
                    v19 = (v19 - v19)
                    if false {
                        if ((-(v19)) != v19) {
                            if false {
                                var v103 float32 = ((0.1 * 1.5) + (((0.1 / -2.25) * 0.1) / 0.1))
                                switch "" {
                                case "foo bar":
                                    if true {
                                        if (!(42 >= 128)) {
                                            v3 -= (^(v3))
                                            v104 := GenS14{F0: -1}
                                            _ = v104
                                        }
                                    } else {
                                        v3 -= v3
                                        v105 := GenS15{F0: true, F1: 1}
                                        _ = v105
                                    }
                                case "a":
                                    var v106 int = fi101
                                    _ = v106
                                case "\n\t":
                                    v107 := GenS16{F0: v1}
                                    _ = v107
                                default:
                                    v108 := GenS17{F0: 42, F1: 255, F2: "a"}
                                    _ = v108
                                }
                                _ = v103
                            }
                        } else {
                            if (false && (!("a" <= ("a" + ("a" + ""))))) {
                                delete(v2, "hello")
                            } else {
                                v109 := []int{fi101}
                                if true {
                                    var v110 string = ("" + "")
                                    _ = v110
                                } else {
                                    v3 -= 0
                                    v111 := GenS11{F0: -32768, F1: fi94}
                                    _ = v111
                                }
                                var v112 uint16 = 1
                                _ = v109
                                _ = v112
                            }
                        }
                        v19 = v19
                        v19 = v19
                        v3 += 330
                        fmt.Println(v1)
                    }
                    _ = fi101
                    _ = v102
                }
            } else {
                if false {
                    for fi113 := 0; fi113 < 17; fi113++ {
                        var v114 uint32 = 308
                        v115 := v2["hello"]
                        v116 := GenS18{F0: 217, F1: 4294967295}
                        _ = fi113
                        _ = v114
                        _ = v115
                        _ = v116
                    }
                    v3 = (v3 / (((v3 % ((v3) | 1))) | 1))
                }
                var v117 float32 = (3.14159 / (0.0 - float32(0.1)))
                fmt.Println(v1, v1, v117)
                _ = v117
            }
            v4 -= 42
        }
        delete(v2, "hello")
        _ = fi94
    }
    switch (v1 + v1) {
    case 42:
        var v118 uint = 255
        _ = v118
    }
    for fi119 := 0; fi119 < 2; fi119++ {
        for fi120 := 0; fi120 < 3; fi120++ {
            v121 := v2["\n\t"]
            _ = fi120
            _ = v121
        }
        v122 := GenS15{F0: (1e10 <= ((0.1 * 3.14159) - 1e10)), F1: 594}
        _ = fi119
        _ = v122
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v50
}
