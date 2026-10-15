package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int; F1 int8; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 bool; F1 string; F2 int }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint64; F1 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16; F1 int32; F2 string }
type GenS11 struct { F0 string }
type GenS12 struct { F0 bool; F1 int32; F2 int32 }
type GenS13 struct { F0 int64; F1 int; F2 uint64 }
func genFn0(p0 uint, p1 int64, p2 uint32) uint8 {
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 int16 = 1000
        _ = fi4
        _ = v5
    }
    var v6 bool = true
    fmt.Println(v6)
    var v7 int8 = 0
    switch (("\n\t" + ("" + "")) + "\n\t") {
    case "":
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 uint64 = 58
            _ = fi8
            _ = v9
        }
    case "foo bar":
        fmt.Println(p2, p1, v7)
    case "a":
        var v10 uint64 = 18446744073709551615
        _ = v10
    default:
        var v11 string = (("foo bar" + ("hello" + ("foo bar" + ("" + ("\n\t" + ("a" + "hello")))))) + ("foo bar" + "hello"))
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint8) int {
    var v12 uint8 = 0
    fmt.Println(p0)
    if ((860 == 0) || true) {
        var v13 bool = (true && true)
        if false {
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := []int{(fi14 << ((65535) & 63)), int(p0), fi14}
                _ = fi14
                _ = v15
            }
            var v16 int8 = -1
            _ = v16
        } else {
            v13 = (false || (1 <= 45))
            for fi17 := 0; fi17 < 18; fi17++ {
                v18 := []int{}
                var v19 string = ((((("foo bar" + "") + "hello") + ("a" + "hello")) + ("hello" + "hello")) + "foo bar")
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        _ = v13
    }
    _ = p0
    _ = v12
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{}
    var v3 uint = 65535
    delete(v1, 197)
    var v20 int64 = int64(v3)
    v21 := GenS1{F0: 997, F1: v3}
    fmt.Println(v20, v3)
    v22 := GenS1{F0: 1000, F1: (v3 >> ((v3) & 63))}
    if ((false && true) && (true || (!false))) {
        for fi23 := 0; fi23 < 2; fi23++ {
            var v24 uint64 = 990
            delete(v1, 65535)
            _ = fi23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        if false {
            var v26 int32 = -878
            switch (fi25 | genFn1(2)) {
            case -2147483648:
                for fi27 := 0; fi27 < 15; fi27++ {
                    fmt.Println(fi27)
                    v28 := GenS1{F0: -130, F1: v3}
                    if (0 > -128) {
                        var v29 float64 = (0.0 + 3.14159)
                        var v30 int64 = 0
                        if (255 < 2) {
                            v3 += (^(v3))
                            fmt.Println(v30)
                            switch ("a" + "foo bar") {
                            case "\n\t":
                                switch "foo bar" {
                                case "":
                                    v26 = 42
                                }
                            default:
                                var v31 int16 = 70
                                _ = v31
                            }
                            v20 += v20
                        }
                        _ = v29
                        _ = v30
                    } else {
                        var v32 float64 = 1e10
                        _ = v32
                    }
                    _ = fi27
                    _ = v28
                }
            case 9223372036854775807:
                for fi33 := 0; fi33 < 5; fi33++ {
                    if (-745 < v26) {
                        switch fi33 {
                        case 389:
                            if ((!true) || (!true)) {
                                var v34 int = fi25
                                if true {
                                    var v35 int8 = 42
                                    switch "" {
                                    case "hello":
                                        v35 = -128
                                    case "foo bar":
                                        if len(v2) > 0 {
                                            v2[0] = "foo bar"
                                        }
                                    }
                                    _ = v35
                                } else {
                                    v36 := GenS1{F0: -1, F1: 200}
                                    var v37 uint8 = genFn0(v3, v20, 4294967295)
                                    _ = v36
                                    _ = v37
                                }
                                _ = v34
                            } else {
                                if true {
                                    v38 := GenS2{F0: 198, F1: 42, F2: 65535}
                                    v3 = (^(v3))
                                    v39 := v1[65535]
                                    _ = v38
                                    _ = v39
                                } else {
                                    fmt.Println(fi25, v3, v3)
                                    v20 -= (-(v20))
                                }
                            }
                        case 796:
                            if true {
                                var v40 int32 = 0
                                switch fi33 {
                                case -2147483648:
                                    var v41 int16 = int16(v3)
                                    _ = v41
                                case 329:
                                    if false {
                                        v42 := GenS1{F0: -1, F1: (v3 << (((v3 - (v3 + 18446744073709551615))) & 63))}
                                        if false {
                                            if true {
                                                var v43 uint8 = 190
                                                if false {
                                                    v1[37] = 2
                                                    delete(v1, 42)
                                                }
                                                switch "" {
                                                case "foo bar":
                                                    _ = len(v1)
                                                    if false {
                                                        v44 := GenS3{F0: (true && (true && false))}
                                                        fmt.Println(v26, v20)
                                                        if (!(!(((0.0 + float32(((1e10 + (1e10 / 1.5)) * -2.25))) * 0.1) == 1e10))) {
                                                            v45 := GenS4{F0: (v20 ^ 0), F1: "\n\t"}
                                                            _ = v45
                                                        } else {
                                                            v3 -= (v3 - 18446744073709551615)
                                                            fmt.Println(fi33, v3, v3)
                                                        }
                                                        if true {
                                                            _ = len(v2)
                                                        }
                                                        _ = v44
                                                    }
                                                case "\n\t":
                                                    v3 += v3
                                                default:
                                                    v46 := GenS1{F0: -32768, F1: (v3 * v3)}
                                                    var v47 int8 = 127
                                                    _ = v46
                                                    _ = v47
                                                }
                                                _ = v43
                                            } else {
                                                v1[948] = 4294967295
                                                switch (^(fi25)) {
                                                case -2147483648:
                                                    v26 += -530
                                                case 32768:
                                                    v48 := GenS5{F0: (1000 > -813), F1: "\n\t", F2: (^(fi25))}
                                                    _ = v48
                                                default:
                                                    v49 := GenS6{F0: ""}
                                                    _ = v49
                                                }
                                            }
                                            var v50 int = fi33
                                            v2 = append(v2, ("hello" + ""))
                                            if (((^(v3)) <= (v3 / (((v3 + v3)) | 1))) || true) {
                                                if false {
                                                    switch (v50 & v50) {
                                                    case -992:
                                                        v50 = (fi33 | fi25)
                                                    }
                                                    switch "\n\t" {
                                                    case "a":
                                                        v51 := GenS7{F0: "foo bar", F1: false}
                                                        _ = v51
                                                    }
                                                }
                                            }
                                            var v52 int8 = -1
                                            _ = v50
                                            _ = v52
                                        } else {
                                            v40 = v40
                                        }
                                        switch 32767 {
                                        case 128:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        v53 := GenS8{F0: 42, F1: (fi25 < genFn1(255))}
                                        _ = v42
                                        _ = v53
                                    }
                                }
                                _ = v40
                            }
                        case -2147483648:
                            fmt.Println(v26)
                            v26 = v26
                            var v54 bool = true
                            if (v54 || true) {
                                v26 += (^(v26))
                                _ = len(v2)
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                                fmt.Println(fi25)
                                delete(v1, 646)
                            } else {
                                var v55 int8 = 120
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch 1 {
                                case 219:
                                    var v56 int32 = -1
                                    _ = v56
                                case -693:
                                    var v57 uint64 = 823
                                    _ = len(v1)
                                    _ = v57
                                default:
                                    if true {
                                        v58 := GenS6{F0: ("a" + (("foo bar" + (("hello" + "foo bar") + ("foo bar" + "\n\t"))) + ("a" + "")))}
                                        _ = v58
                                    } else {
                                        switch "\n\t" {
                                        case "hello":
                                            fmt.Println(fi33)
                                        }
                                    }
                                }
                                if v54 {
                                    if v54 {
                                        if ((fi33 % ((genFn1(uint8(v3))) | 1)) != fi33) {
                                            v59 := GenS3{F0: (v54 && (!false))}
                                            _ = v59
                                        } else {
                                            v26 = (v26 & v26)
                                            v26 = int32(v55)
                                            var v60 int32 = (v26 & v26)
                                            _ = v60
                                        }
                                    }
                                }
                                _ = v55
                            }
                            var v61 float32 = 3.14159
                            _ = v54
                            _ = v61
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        var v62 uint8 = 255
                        _ = v62
                    } else {
                        if false {
                            v20 += -9223372036854775808
                            fmt.Println(fi25)
                        }
                        if len(v2) > 0 {
                            v2[0] = (("hello" + ((("hello" + "\n\t") + (("\n\t" + "hello") + "hello")) + ((("" + "hello") + (("" + "a") + "")) + "a"))) + "a")
                        }
                        fmt.Println(v20, fi33, v3)
                    }
                    var v63 int32 = v26
                    var v64 uint32 = 480
                    _ = fi33
                    _ = v63
                    _ = v64
                }
            }
            for fi65 := 0; fi65 < 2; fi65++ {
                switch (fi65 & fi65) {
                case -624:
                    v3 = (v3 >> ((4294967295) & 63))
                default:
                    if true {
                        v3 = 0
                    } else {
                        v26 += (v26 | v26)
                        fmt.Println(v20, fi25, v26)
                        delete(v1, 42)
                    }
                }
                fmt.Println(v20, v20)
                _ = fi65
            }
            _ = v26
        }
        for fi66 := 0; fi66 < 10; fi66++ {
            v67 := GenS2{F0: genFn1(42), F1: -81, F2: 42}
            v68 := GenS9{F0: "foo bar"}
            _ = fi66
            _ = v67
            _ = v68
        }
        _ = fi25
    }
    var v69 uint64 = 0
    var v70 float64 = float64(3.14159)
    v71 := GenS4{F0: v20, F1: "a"}
    switch "hello" {
    case "hello":
        var v72 uint = v3
        _ = v72
    case "":
        if true {
            var v73 uint = v3
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v74 int8 = -1
            _ = v73
            _ = v74
        } else {
            v3 = 73
            v69 = 4294967295
            v75 := GenS4{F0: (v20 ^ (v20 + v20)), F1: ("hello" + "a")}
            _ = v75
        }
    }
    var v76 int = -910
    for fi77 := 0; fi77 < 11; fi77++ {
        v78 := GenS9{F0: ""}
        if true {
            if false {
                var v79 uint = (v3 & v3)
                _ = v79
            }
            v69 += (v69 * v69)
        } else {
            for fi80 := 0; fi80 < 18; fi80++ {
                var v81 int64 = (v20 * v20)
                v69 -= (v69 << ((2) & 63))
                v70 -= float64(1.5)
                _ = fi80
                _ = v81
            }
        }
        _ = fi77
        _ = v78
    }
    if ((16 == 255) && true) {
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            var v82 uint16 = 65535
            _ = v82
        }
    } else {
        v83 := GenS5{F0: true, F1: "a", F2: (v76 & (-(v76)))}
        if false {
            var v84 uint8 = 2
            v2 = append(v2, "a")
            _ = v84
        } else {
            fmt.Println(v76, v3, v70)
        }
        _ = v83
    }
    for fi85 := 0; fi85 < 2; fi85++ {
        var v86 uint = 255
        _ = fi85
        _ = v86
    }
    v87 := GenS10{F0: 0, F1: 2147483647, F2: ("a" + ("a" + "foo bar"))}
    if true {
        v3 = v3
        fmt.Println(v70, v3)
    } else {
        fmt.Println(v69, v69)
        for fi88 := 0; fi88 < 4; fi88++ {
            if true {
                v89 := GenS2{F0: (v76 % (((^(v76))) | 1)), F1: 22, F2: 327}
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                _ = v89
            } else {
                var v90 string = ("a" + "\n\t")
                switch ("hello" + "hello") {
                case "hello":
                    if false {
                        if true {
                            var v91 uint8 = 2
                            v92 := GenS1{F0: -332, F1: v3}
                            if false {
                                if ((!true) && false) {
                                    fmt.Println(v69, v20, v70)
                                    var v93 uint32 = 4294967295
                                    _ = v93
                                }
                                for fi94 := 0; fi94 < 9; fi94++ {
                                    var v95 int8 = 0
                                    _ = fi94
                                    _ = v95
                                }
                            } else {
                                v96 := GenS11{F0: ((v90 + (("foo bar" + (v90 + v90)) + "foo bar")) + v90)}
                                var v97 int = v76
                                _ = v96
                                _ = v97
                            }
                            _ = v91
                            _ = v92
                        }
                        var v98 int32 = 1
                        if true {
                            if false {
                                v99 := GenS12{F0: true, F1: v98, F2: v98}
                                v90 = v90
                                v100 := GenS10{F0: 32767, F1: 42, F2: ("" + "foo bar")}
                                _ = v100
                                _ = v99
                            } else {
                                switch (fi88 - 272) {
                                case 1:
                                    var v101 bool = false
                                    _ = v101
                                case -9223372036854775808:
                                    for fi102 := 0; fi102 < 15; fi102++ {
                                        fmt.Println(v70, v90)
                                        _ = fi102
                                    }
                                }
                                v70 = float64(0.1)
                            }
                            switch v76 {
                            case 11:
                                v76 = (^(v76))
                            default:
                                for fi103 := 0; fi103 < 8; fi103++ {
                                    v2 = append(v2, v90)
                                    _ = fi103
                                }
                            }
                            var v104 uint64 = 4294967295
                            _ = v104
                        }
                        _ = v98
                    }
                case "":
                    var v105 uint16 = 42
                    _ = v105
                case "\n\t":
                    v106 := GenS13{F0: v20, F1: (fi88 + 127), F2: 2}
                    _ = v106
                default:
                    var v107 int64 = (v20 * (v20 | v20))
                    _ = v107
                }
                _ = len(v1)
                _ = v90
            }
            v20 += (v20 << (((v3 + 455)) & 63))
            if (((float32(3.14159) / (3.14159 - (((float32(v70) + (float32((v70 * v70)) * 3.14159)) + ((1.5 - (1e10 + -2.25)) * 1e10)) - 1e10))) - 1.5) == 1e10) {
                if (127 >= -128) {
                    var v108 int = int(v20)
                    v109 := GenS5{F0: ("" >= ("\n\t" + "")), F1: "", F2: int(v20)}
                    _ = v108
                    _ = v109
                }
                if false {
                    fmt.Println(fi88)
                }
            }
            _ = fi88
        }
        for fi110 := 0; fi110 < 1; fi110++ {
            if false {
                var v111 bool = ((("hello" >= "hello") || false) || (!false))
                _ = v111
            }
            if (true && (1 == v3)) {
                var v112 uint64 = 1
                var v113 int8 = 127
                _ = v112
                _ = v113
            } else {
                var v114 int16 = 0
                if (true || (fi110 >= genFn1(128))) {
                    var v115 uint32 = 2147483648
                    switch -419 {
                    case 580:
                        switch 32767 {
                        case 128:
                            v3 -= v3
                        default:
                            for fi116 := 0; fi116 < 17; fi116++ {
                                switch v76 {
                                case 4294967295:
                                    v115 = v115
                                case 9223372036854775807:
                                    delete(v1, v115)
                                default:
                                    v3 = v3
                                }
                                switch "\n\t" {
                                case "\n\t":
                                    var v117 float32 = 0.0
                                    _ = v117
                                case "a":
                                    v20 += (v20 | 319)
                                default:
                                    v118 := GenS8{F0: 1, F1: (v115 <= 2147483648)}
                                    _ = v118
                                }
                                _ = fi116
                            }
                        }
                    default:
                        var v119 uint32 = (^(v115))
                        _ = v119
                    }
                    if false {
                        switch "" {
                        case "":
                            if false {
                                v120 := GenS2{F0: (-(fi110)), F1: 19, F2: 2}
                                v121 := GenS4{F0: (v20 << ((42) & 63)), F1: ""}
                                var v122 int64 = 9223372036854775807
                                _ = v120
                                _ = v121
                                _ = v122
                            }
                        }
                        fmt.Println(fi110, v115, fi110)
                        v123 := GenS1{F0: -32768, F1: (v3 | v3)}
                        _ = v123
                    }
                    _ = v115
                } else {
                    if true {
                        var v124 uint32 = 2
                        var v125 uint16 = 35
                        _ = v124
                        _ = v125
                    }
                    v69 = v69
                    var v126 bool = false
                    _ = v126
                }
                v76 = int(v114)
                _ = v114
            }
            _ = fi110
        }
    }
    var v127 float32 = (3.14159 / (0.0 - 3.14159))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v127:", v127)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v127
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v69
    _ = v70
    _ = v71
    _ = v76
    _ = v87
}
