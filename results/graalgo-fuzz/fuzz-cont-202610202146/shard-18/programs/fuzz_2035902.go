package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 string; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 int8; F1 uint; F2 uint64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int; F1 int32; F2 uint32 }
type GenS9 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS10 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS11 struct { F0 bool; F1 int16 }
type GenS12 struct { F0 uint64; F1 int16 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 bool; F1 bool; F2 int64 }
func genFn0(p0 int16, p1 uint32) uint8 {
    v7 := []int{int(p0)}
    p0 -= -32768
    p1 += (p1 & (p1 + p1))
    p0 = p0
    for fi8 := 0; fi8 < 4; fi8++ {
        switch 0 {
        case 127:
            var v9 bool = true
            _ = v9
        case 128:
            fmt.Println(p0, p0)
        }
        if false {
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := GenS1{F0: true}
                v12 := GenS1{F0: false}
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        var v13 float64 = (((1.5 + ((((3.14159 + ((1.5 + ((0.1 / 1.0) + 1e10)) + float64((0.1 + 1.5)))) / float64((-2.25 * float32(-2.25)))) / 0.1) * 1.5)) + -2.25) - 1e10)
        _ = fi8
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v7
    return 128
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 32767
    v3 := make(map[string]bool)
    var v4 int32 = -2147483648
    v5 := make(map[uint32]int)
    v6 := []int{v2, (v2 ^ -9223372036854775808)}
    var v14 float32 = 1e10
    v15 := GenS1{F0: true}
    v16 := GenS1{F0: (!false)}
    if (!(1 >= 4294967295)) {
        switch (v2 + 32768) {
        case 738:
            if false {
                if ((true && true) || true) {
                    v4 = int32(v2)
                    v4 = v4
                    v4 += (^(v4))
                }
            } else {
                v2 = v2
            }
        }
        v17 := GenS1{F0: ((v2 != -1) && false)}
        _ = v17
    } else {
        switch (^(v2)) {
        case 4294967295:
            var v18 int32 = v4
            _ = v18
        case 42:
            switch "a" {
            case "foo bar":
                var v19 float32 = (v14 - v14)
                _ = v19
            case "\n\t":
                for fi20 := 0; fi20 < 7; fi20++ {
                    fmt.Println(v14, v2)
                    for fi21 := 0; fi21 < 10; fi21++ {
                        if true {
                            v22 := GenS2{F0: "\n\t", F1: "a", F2: 0}
                            if len(v6) > 0 {
                                v6[0] = 1
                            }
                            var v23 uint8 = 127
                            _ = v22
                            _ = v23
                        } else {
                            var v24 string = ("hello" + "\n\t")
                            var v25 int = fi20
                            _ = v24
                            _ = v25
                        }
                        _ = fi21
                    }
                    switch "foo bar" {
                    case "":
                        if (!((fi20 - v2) < fi20)) {
                            v26 := GenS1{F0: true}
                            _ = v26
                        } else {
                            switch (fi20 >> ((65535) & 63)) {
                            case 4294967295:
                                switch ("\n\t" + "hello") {
                                case "a":
                                    delete(v3, "hello")
                                case "":
                                    v27 := GenS1{F0: true}
                                    _ = v27
                                default:
                                    fmt.Println(fi20, fi20, fi20)
                                }
                            case 127:
                                v28 := GenS1{F0: true}
                                _ = v28
                            }
                            for fi29 := 0; fi29 < 5; fi29++ {
                                var v30 uint64 = 18446744073709551615
                                _ = fi29
                                _ = v30
                            }
                        }
                    case "foo bar":
                        if false {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            for fi31 := 0; fi31 < 16; fi31++ {
                                v32 := GenS2{F0: "a", F1: (("foo bar" + "\n\t") + (((("" + (((("a" + "\n\t") + "a") + (("hello" + "foo bar") + "a")) + "a")) + (("a" + "a") + "hello")) + "\n\t") + ("foo bar" + ("a" + ("hello" + "foo bar"))))), F2: 4294967295}
                                var v33 uint8 = 1
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                        }
                    case "a":
                        for fi34 := 0; fi34 < 10; fi34++ {
                            fmt.Println(v2)
                            var v35 int64 = 853
                            _ = fi34
                            _ = v35
                        }
                    }
                    _ = fi20
                }
            default:
                if true {
                    if (1 <= -1) {
                        var v36 string = ""
                        _ = v36
                    }
                }
            }
        default:
            v37 := GenS3{F0: 0, F1: 2147483647}
            _ = v37
        }
        _ = len(v6)
        for fi38 := 0; fi38 < 19; fi38++ {
            switch v2 {
            case -731:
                var v39 int64 = -9223372036854775808
                _ = v39
            case 2147483647:
                v40 := GenS3{F0: 42, F1: 2147483648}
                _ = v40
            case 1:
                v41 := GenS3{F0: -9223372036854775808, F1: 2}
                _ = v41
            }
            if (!(18 <= 127)) {
                var v42 string = "\n\t"
                _ = v42
            } else {
                if (!true) {
                    var v43 string = ("" + "\n\t")
                    switch v2 {
                    case 0:
                        v44 := GenS4{F0: v43}
                        _ = v44
                    case 128:
                        fmt.Println(v14, fi38)
                    case 32767:
                        var v45 uint64 = 18446744073709551615
                        _ = v45
                    }
                    if (42 == 9223372036854775807) {
                        if true {
                            v4 -= v4
                            fmt.Println(v43, v4, fi38)
                            var v46 int = v2
                            _ = v46
                        } else {
                            v47 := GenS3{F0: 0, F1: 0}
                            _ = v47
                        }
                        _ = len(v1)
                    }
                    _ = v43
                } else {
                    v4 += 870
                }
            }
            _ = fi38
        }
    }
    var v48 uint64 = 0
    fmt.Println(v2, v48)
    var v49 float64 = 1.5
    v50 := GenS3{F0: -1, F1: 4294967295}
    v4 += v4
    if true {
        v49 = float64((v14 + 3.14159))
        v4 = (v4 ^ 42)
    } else {
        v4 += -2147483648
    }
    v2 += -2147483648
    if false {
        var v51 uint64 = v48
        var v52 uint = 917
        if true {
            var v53 bool = ((!(-25 < -841)) && (!(18446744073709551615 <= (v51 * 2))))
            if v53 {
                for fi54 := 0; fi54 < 14; fi54++ {
                    for fi55 := 0; fi55 < 14; fi55++ {
                        v56 := GenS5{F0: 0, F1: (v4 | (v4 << ((v52) & 63)))}
                        var v57 int16 = 42
                        _ = fi55
                        _ = v56
                        _ = v57
                    }
                    if (!(716 > 42)) {
                        v53 = true
                        var v58 uint64 = v51
                        _ = v58
                    } else {
                        _ = len(v5)
                    }
                    _ = fi54
                }
            }
            _ = v53
        } else {
            v59 := GenS1{F0: ((^(v51)) <= uint64(v52))}
            v51 = (v51 << ((v52) & 63))
            _ = v59
        }
        _ = v51
        _ = v52
    } else {
        v60 := GenS3{F0: 636, F1: 2147483647}
        for fi61 := 0; fi61 < 12; fi61++ {
            switch v2 {
            case 4294967295:
                v5[1] = 32768
            case -159:
                if (v49 >= v49) {
                    var v62 uint32 = 42
                    v63 := GenS5{F0: 1, F1: v4}
                    _ = v62
                    _ = v63
                } else {
                    v48 -= (^(v48))
                    if false {
                        v64 := GenS6{F0: 61, F1: 65535, F2: (v48 << ((2) & 63))}
                        if (0 == 244) {
                            switch fi61 {
                            case 2147483647:
                                for fi65 := 0; fi65 < 18; fi65++ {
                                    fmt.Println(fi65, fi61)
                                    var v66 uint32 = 2147483647
                                    v66 = 706
                                    if (("a" != "\n\t") && false) {
                                        if false {
                                            v67 := GenS4{F0: "\n\t"}
                                            delete(v1, 282)
                                            v68 := GenS7{F0: 4294967295}
                                            _ = v67
                                            _ = v68
                                        }
                                    } else {
                                        switch -2147483648 {
                                        case -726:
                                            fmt.Println(fi61)
                                        case -9223372036854775808:
                                            v69 := GenS1{F0: true}
                                            _ = v69
                                        case 4294967295:
                                            v14 = (((v14 - (((3.14159 * (v14 / ((v14 + 1e10) + 3.14159))) / v14) * ((v14 * float32((1.5 + v49))) * v14))) / 1.0) / float32(1.5))
                                        default:
                                            if true {
                                                v70 := GenS7{F0: 1}
                                                var v71 int8 = -1
                                                v72 := GenS8{F0: v2, F1: (v4 >> ((42) & 63)), F2: v66}
                                                _ = v70
                                                _ = v71
                                                _ = v72
                                            }
                                        }
                                        v73 := GenS9{F0: (float32((float64((float32(0.1) + (((v14 - (0.0 / float32((v49 + (((v49 + v49) - v49) + 3.14159))))) + float32(v49)) * v14))) / 3.14159)) != (v14 * v14)), F1: 2, F2: (v66 - (v66 / ((v66) | 1)))}
                                        var v74 int16 = 42
                                        _ = v73
                                        _ = v74
                                    }
                                    _ = fi65
                                    _ = v66
                                }
                            case 1:
                                for fi75 := 0; fi75 < 2; fi75++ {
                                    fmt.Println(v2, v4)
                                    fmt.Println(v14, v2, v2)
                                    _ = fi75
                                }
                            }
                        }
                        v14 -= float32((3.14159 + float64((0.0 * (3.14159 * (v14 - float32((v49 / v49))))))))
                        _ = v64
                    } else {
                        v48 = (v48 * v48)
                        for fi76 := 0; fi76 < 13; fi76++ {
                            v4 = (v4 | (v4 | v4))
                            _ = fi76
                        }
                    }
                    var v77 int16 = 0
                    _ = v77
                }
            case -278:
                var v78 uint16 = 2
                _ = v78
            default:
                fmt.Println(v49)
            }
            switch ("foo bar" + "") {
            case "foo bar":
                v48 -= v48
            case "a":
                if false {
                    switch 9223372036854775807 {
                    case -1:
                        v48 = v48
                    case 616:
                        if (!false) {
                            v49 -= (v49 - v49)
                            if true {
                                if ((true || (true || (float64((v14 / -2.25)) > v49))) && false) {
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch v2 {
                                        case 4294967295:
                                            if (!false) {
                                                fmt.Println(v49)
                                                var v79 int64 = 793
                                                for fi80 := 0; fi80 < 20; fi80++ {
                                                    v49 += (v49 * v49)
                                                    var v81 uint16 = 0
                                                    _ = fi80
                                                    _ = v81
                                                }
                                                _ = v79
                                            } else {
                                                var v82 int32 = v4
                                                _ = v82
                                            }
                                        case -2147483648:
                                            var v83 int = (v2 >> ((1) & 63))
                                            _ = v83
                                        case 2147483647:
                                            var v84 float64 = v49
                                            _ = v84
                                        }
                                    }
                                    v14 += float32(float64(0.0))
                                } else {
                                    var v85 string = ("hello" + "\n\t")
                                    _ = v85
                                }
                                for fi86 := 0; fi86 < 13; fi86++ {
                                    if true {
                                        if false {
                                            var v87 string = "a"
                                            v6 = append(v6, v2)
                                            fmt.Println(fi61)
                                            v4 += (v4 - -346)
                                            _ = v87
                                        } else {
                                            var v88 uint64 = (v48 >> ((0) & 63))
                                            v49 = 0.0
                                            _ = v88
                                        }
                                    } else {
                                        if (false && (v4 == v4)) {
                                            switch ("hello" + (("a" + (("\n\t" + ("" + "")) + (("a" + "") + "foo bar"))) + "a")) {
                                            case "hello":
                                                v14 = (v14 / (v14 - v14))
                                            case "a":
                                                v14 = v14
                                            case "":
                                                v89 := GenS7{F0: uint(v48)}
                                                _ = v89
                                            default:
                                                var v90 int = fi86
                                                _ = v90
                                            }
                                        } else {
                                            var v91 uint = 65535
                                            fmt.Println(v4)
                                            _ = len(v5)
                                            _ = v91
                                        }
                                        if (v4 == v4) {
                                            v2 -= (^(fi86))
                                        } else {
                                            switch (fi86 % (((fi61 ^ (fi86 % (((-(fi86))) | 1)))) | 1)) {
                                            case 689:
                                                var v92 uint = 65535
                                                _ = v92
                                            case 97:
                                                if false {
                                                    switch 464 {
                                                    case -782:
                                                        var v93 int8 = 100
                                                        _ = v93
                                                    default:
                                                        var v94 int8 = 79
                                                        _ = v94
                                                    }
                                                    v49 = v49
                                                } else {
                                                    v49 = 1.5
                                                }
                                                if (-1 > 42) {
                                                    v14 = v14
                                                    var v95 bool = false
                                                    _ = v95
                                                }
                                            case -478:
                                                v4 = (v4 << ((225) & 63))
                                            default:
                                                v49 += v49
                                            }
                                        }
                                    }
                                    _ = fi86
                                }
                                var v96 uint64 = 58
                                _ = v96
                            }
                            if false {
                                if false {
                                    v97 := GenS10{F0: 138, F1: v48, F2: 1}
                                    if (v14 != v14) {
                                        for fi98 := 0; fi98 < 14; fi98++ {
                                            var v99 uint8 = 0
                                            _ = fi98
                                            _ = v99
                                        }
                                    } else {
                                        var v100 int8 = 100
                                        v2 = fi61
                                        for fi101 := 0; fi101 < 1; fi101++ {
                                            v100 = v100
                                            _ = fi101
                                        }
                                        _ = v100
                                    }
                                    if true {
                                        var v102 uint16 = 255
                                        switch "" {
                                        case "hello":
                                            var v103 uint64 = (v48 + (v48 - 0))
                                            _ = v103
                                        case "":
                                            if true {
                                                delete(v3, "hello")
                                                v104 := GenS3{F0: 254, F1: 1}
                                                v14 = v14
                                                _ = v104
                                            }
                                        case "\n\t":
                                            fmt.Println(fi61, v102, v4)
                                        }
                                        _ = v102
                                    }
                                    _ = v97
                                } else {
                                    if false {
                                        for fi105 := 0; fi105 < 4; fi105++ {
                                            if (!(!(fi105 < (v2 ^ 0)))) {
                                                if false {
                                                    if true {
                                                        v4 = (v4 << ((65535) & 63))
                                                    } else {
                                                        v106 := v1[65535]
                                                        v107 := GenS9{F0: false, F1: 0, F2: 2}
                                                        _ = v106
                                                        _ = v107
                                                    }
                                                    v14 -= v14
                                                    if true {
                                                        if len(v6) > 0 {
                                                            v6[0] = (v2 >> ((1) & 63))
                                                        }
                                                        if (83 != -5) {
                                                            var v108 int32 = v4
                                                            _ = v108
                                                        }
                                                        var v109 string = ("" + "a")
                                                        _ = v109
                                                    }
                                                } else {
                                                    v110 := GenS7{F0: 255}
                                                    v2 += (fi61 + (fi61 * 127))
                                                    v111 := GenS9{F0: (!true), F1: 42, F2: 2147483647}
                                                    _ = v110
                                                    _ = v111
                                                }
                                                var v112 int16 = 270
                                                if ((!true) && true) {
                                                    v113 := GenS11{F0: true, F1: (v112 << ((65535) & 63))}
                                                    _ = v113
                                                } else {
                                                    v114 := GenS12{F0: v48, F1: (v112 % (((-(v112))) | 1))}
                                                    _ = len(v6)
                                                    _ = v114
                                                }
                                                _ = v112
                                            } else {
                                                var v115 string = ((("hello" + "\n\t") + "a") + "a")
                                                var v116 float64 = v49
                                                v117 := GenS1{F0: (414 <= 575)}
                                                fmt.Println(fi105)
                                                fmt.Println(fi61, v116, v4)
                                                var v118 int = 0
                                                _ = v115
                                                _ = v116
                                                _ = v117
                                                _ = v118
                                            }
                                            _ = fi105
                                        }
                                        delete(v1, 0)
                                    }
                                    switch (v2 - (-(fi61))) {
                                    case -2147483648:
                                        if (255 != 42) {
                                            v119 := GenS9{F0: false, F1: 939, F2: 0}
                                            if (1 <= 0) {
                                                v49 -= 1.5
                                            } else {
                                                fmt.Println(v14, v49)
                                                v2 = v2
                                            }
                                            for fi120 := 0; fi120 < 7; fi120++ {
                                                v49 = 0.0
                                                _ = fi120
                                            }
                                            _ = v119
                                        } else {
                                            for fi121 := 0; fi121 < 18; fi121++ {
                                                v4 -= -201
                                                _ = fi121
                                            }
                                            v49 -= 0.0
                                        }
                                    case -9223372036854775808:
                                        var v122 uint32 = 4294967295
                                        _ = v122
                                    case -512:
                                        v2 = (v2 / (((fi61 - fi61)) | 1))
                                    default:
                                        for fi123 := 0; fi123 < 1; fi123++ {
                                            if false {
                                                v48 = (v48 / (((v48 & (v48 >> ((65535) & 63)))) | 1))
                                            } else {
                                                v49 = ((v49 * v49) + 0.1)
                                                var v124 uint = 2
                                                v5[65535] = fi61
                                                _ = v124
                                            }
                                            v125 := GenS8{F0: 458, F1: 0, F2: 0}
                                            _ = fi123
                                            _ = v125
                                        }
                                    }
                                }
                                var v126 float64 = 0.0
                                _ = v126
                            } else {
                                v4 -= (^(v4))
                                v2 = fi61
                            }
                        }
                    case 1:
                        var v127 float32 = (v14 - v14)
                        _ = v127
                    default:
                        v128 := GenS13{F0: 42}
                        _ = v128
                    }
                }
            case "\n\t":
                switch v2 {
                case 790:
                    fmt.Println(v4, v2, fi61)
                }
            }
            _ = fi61
        }
        if true {
            fmt.Println(v48, v49)
            if false {
                fmt.Println(v49, v14, v2)
            }
            fmt.Println(v48, v48)
        }
        _ = v60
    }
    if (!(944 > uint(v48))) {
        v2 -= v2
        v129 := GenS14{F0: false, F1: false, F2: -9223372036854775808}
        var v130 string = "a"
        _ = v129
        _ = v130
    } else {
        var v131 int = (v2 * v2)
        switch (v131 % ((588) | 1)) {
        case -9223372036854775808:
            var v132 uint = 65535
            _ = v132
        case 32767:
            var v133 int8 = -1
            _ = v133
        case 1:
            v49 = v49
        default:
            _ = len(v5)
        }
        _ = v131
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
