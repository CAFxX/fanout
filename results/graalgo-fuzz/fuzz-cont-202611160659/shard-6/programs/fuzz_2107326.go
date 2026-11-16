package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 int8; F2 bool }
type GenS7 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS10 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS11 struct { F0 int8; F1 uint64 }
type GenS12 struct { F0 int16; F1 uint; F2 uint64 }
type GenS13 struct { F0 uint32; F1 string }

func main() {
    v1 := make(map[string]string)
    v2 := []int{9223372036854775807, 32768, 4294967295, -9223372036854775808}
    var v3 int8 = 127
    var v4 uint = 620
    var v5 int16 = -32768
    var v6 float32 = -2.25
    v5 += (^(v5))
    v7 := GenS1{F0: -2147483648}
    if (!(2147483648 != 2147483648)) {
        v2 = append(v2, 1)
        if (!true) {
            v8 := GenS1{F0: 9223372036854775807}
            _ = len(v1)
            _ = v8
        } else {
            v9 := GenS1{F0: 241}
            _ = v9
        }
    } else {
        if (("foo bar" + "foo bar") > "hello") {
            switch 42 {
            case 9223372036854775807:
                v10 := GenS2{F0: (("foo bar" + ("" + "a")) + ("a" + "foo bar"))}
                _ = v10
            case 967:
                v2 = append(v2, 127)
            case 0:
                var v11 float64 = 1.5
                _ = v11
            default:
                var v12 float32 = v6
                _ = v12
            }
            var v13 uint32 = 911
            _ = v13
        } else {
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 float32 = v6
                v6 = v15
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 uint = (v4 | v4)
                    var v18 float64 = 3.14159
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = fi14
                _ = v15
            }
        }
    }
    var v19 float64 = float64((0.0 + v6))
    switch 393 {
    case 4294967295:
        for fi20 := 0; fi20 < 7; fi20++ {
            for fi21 := 0; fi21 < 3; fi21++ {
                switch "" {
                case "foo bar":
                    v5 -= 1
                case "a":
                    if true {
                        var v22 int32 = 2147483647
                        var v23 uint16 = 42
                        _ = v22
                        _ = v23
                    }
                case "":
                    fmt.Println(v6, fi20)
                default:
                    v24 := GenS2{F0: ((("hello" + "hello") + ("hello" + "a")) + "foo bar")}
                    _ = v24
                }
                if true {
                    var v25 uint8 = 255
                    v19 = float64(0.0)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    delete(v1, ("\n\t" + "foo bar"))
                    var v26 uint64 = 4294967295
                    if ((1 <= 42) && (v3 == 42)) {
                        v26 += 65
                    } else {
                        v27 := v1[(("" + ("foo bar" + "hello")) + "hello")]
                        switch fi21 {
                        case 32767:
                            v28 := GenS3{F0: 2, F1: v25}
                            _ = v28
                        case 366:
                            v26 += 4294967295
                        default:
                            fmt.Println(v5, v5)
                        }
                        var v29 uint32 = 295
                        _ = v27
                        _ = v29
                    }
                    _ = v25
                    _ = v26
                }
                v4 += (v4 ^ 1)
                _ = fi21
            }
            for fi30 := 0; fi30 < 5; fi30++ {
                v19 += (0.1 + v19)
                fmt.Println(v19, v6, v6)
                _ = fi30
            }
            var v31 uint16 = 255
            _ = fi20
            _ = v31
        }
    case 502:
        v5 = v5
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        v1["foo bar"] = "\n\t"
        var v33 int16 = (v5 & 32767)
        _ = fi32
        _ = v33
    }
    if ((42 < 219) && (false || true)) {
        if (!false) {
            v5 = 923
        } else {
            if ((v19 + v19) != 0.0) {
                fmt.Println(v4, v5, v3)
                v34 := GenS3{F0: 1, F1: 26}
                v19 -= (0.0 / v19)
                _ = v34
            }
            if false {
                var v35 string = ("\n\t" + (("hello" + "a") + "foo bar"))
                v36 := GenS3{F0: 315, F1: 155}
                _ = v35
                _ = v36
            }
            if (false || false) {
                if len(v2) > 0 {
                    v2[0] = 0
                }
                if true {
                    var v37 int64 = 9223372036854775807
                    for fi38 := 0; fi38 < 14; fi38++ {
                        fmt.Println(v4, v5, fi38)
                        switch "\n\t" {
                        case "a":
                            v39 := GenS4{F0: "foo bar", F1: 2}
                            _ = v39
                        case "":
                            var v40 string = ("foo bar" + "")
                            _ = v40
                        case "hello":
                            if true {
                                var v41 int32 = 104
                                fmt.Println(v6, v3, v4)
                                var v42 uint8 = 225
                                _ = v41
                                _ = v42
                            } else {
                                fmt.Println(v19)
                                var v43 uint8 = 62
                                for fi44 := 0; fi44 < 10; fi44++ {
                                    fmt.Println(v6, v4, v43)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v45 string = "\n\t"
                                    _ = fi44
                                    _ = v45
                                }
                                _ = v43
                            }
                        default:
                            for fi46 := 0; fi46 < 7; fi46++ {
                                if false {
                                    if (42 >= 115) {
                                        if false {
                                            if false {
                                                v6 -= (3.14159 + v6)
                                            }
                                            v37 = (v37 & (v37 % ((0) | 1)))
                                            v3 = v3
                                        }
                                        fmt.Println(v5)
                                        switch ("" + "\n\t") {
                                        case "\n\t":
                                            v47 := GenS3{F0: 32767, F1: 127}
                                            _ = v47
                                        case "hello":
                                            v48 := GenS5{F0: v5}
                                            _ = v48
                                        case "":
                                            if false {
                                                v2 = append(v2, fi46)
                                                fmt.Println(v37)
                                                switch (("hello" + ("a" + "")) + "a") {
                                                case "hello":
                                                    switch fi38 {
                                                    case -9223372036854775808:
                                                        if (true || true) {
                                                            v5 = (v5 << ((12) & 63))
                                                            v49 := GenS6{F0: "a", F1: v3, F2: (!(2147483648 != 1))}
                                                            fmt.Println(v19, v19, v37)
                                                            _ = v49
                                                        }
                                                    }
                                                case "":
                                                    _ = len(v2)
                                                case "foo bar":
                                                    v50 := GenS7{F0: false, F1: 0, F2: 128}
                                                    _ = v50
                                                default:
                                                    v5 += 1
                                                }
                                                v37 += v37
                                                var v51 float64 = (((v19 + v19) * (((float64(v6) / (v19 + v19)) + v19) * (v19 / v19))) + v19)
                                                _ = v51
                                            }
                                        default:
                                            fmt.Println(v4, v3)
                                            v3 += (v3 / ((v3) | 1))
                                        }
                                    } else {
                                        switch fi46 {
                                        case 32768:
                                            v52 := GenS8{F0: (v4 - (v4 << ((v4) & 63)))}
                                            _ = v52
                                        default:
                                            var v53 float64 = (float64((v6 / float32(v19))) + v19)
                                            _ = v53
                                        }
                                    }
                                    v54 := GenS2{F0: ""}
                                    v2 = append(v2, fi38)
                                    _ = v54
                                } else {
                                    var v55 int8 = (v3 & 57)
                                    v3 += -90
                                    v56 := GenS8{F0: v4}
                                    _ = v55
                                    _ = v56
                                }
                                var v57 uint = (v4 - 1)
                                _ = fi46
                                _ = v57
                            }
                        }
                        v58 := GenS9{F0: 42, F1: (v4 % ((v4) | 1)), F2: 126}
                        _ = fi38
                        _ = v58
                    }
                    v59 := GenS10{F0: 4294967295, F1: v5, F2: 1}
                    _ = v37
                    _ = v59
                } else {
                    for fi60 := 0; fi60 < 20; fi60++ {
                        if (true || true) {
                            fmt.Println(v6)
                            for fi61 := 0; fi61 < 16; fi61++ {
                                var v62 bool = false
                                v63 := GenS3{F0: 248, F1: 2}
                                _ = fi61
                                _ = v62
                                _ = v63
                            }
                            switch fi60 {
                            case 9223372036854775807:
                                v4 = (v4 % ((v4) | 1))
                            case 48:
                                for fi64 := 0; fi64 < 9; fi64++ {
                                    fmt.Println(v19, v5, fi60)
                                    switch "hello" {
                                    case "a":
                                        switch ("foo bar" + ("a" + "\n\t")) {
                                        case "":
                                            var v65 uint64 = 42
                                            _ = v65
                                        }
                                    case "\n\t":
                                        _ = len(v2)
                                    }
                                    if false {
                                        switch (fi60 - 9223372036854775807) {
                                        case 42:
                                            v5 = v5
                                        default:
                                            var v66 float32 = v6
                                            var v67 int64 = 0
                                            _ = v66
                                            _ = v67
                                        }
                                        switch "" {
                                        case "hello":
                                            v68 := GenS6{F0: (("hello" + "foo bar") + ("foo bar" + "a")), F1: (v3 * (v3 >> (((^(v4))) & 63))), F2: ((v3 | (-(v3))) != v3)}
                                            _ = v68
                                        case "foo bar":
                                            v19 += v19
                                        case "a":
                                            v6 -= v6
                                        }
                                        var v69 int = (fi64 | fi64)
                                        v19 = float64(v6)
                                        var v70 uint64 = 2
                                        _ = v69
                                        _ = v70
                                    }
                                    _ = fi64
                                }
                            case -1:
                                v5 += v5
                            default:
                                for fi71 := 0; fi71 < 14; fi71++ {
                                    if len(v2) > 0 {
                                        v2[0] = int(v5)
                                    }
                                    _ = fi71
                                }
                            }
                        } else {
                            var v72 uint16 = 537
                            _ = v72
                        }
                        _ = fi60
                    }
                }
            }
        }
        if false {
            v73 := GenS1{F0: 42}
            _ = v73
        }
        var v74 int16 = v5
        _ = v74
    } else {
        v2 = append(v2, -475)
        var v75 uint64 = 788
        var v76 string = (("foo bar" + "hello") + ("\n\t" + "\n\t"))
        _ = v75
        _ = v76
    }
    switch "" {
    case "\n\t":
        switch 32767 {
        case 42:
            v77 := GenS2{F0: "a"}
            _ = v77
        case 0:
            v78 := GenS5{F0: (v5 & 1)}
            _ = v78
        case 305:
            var v79 uint32 = 0
            _ = v79
        }
    case "foo bar":
        switch ("foo bar" + ("" + ("hello" + "hello"))) {
        case "hello":
            var v80 uint16 = 255
            _ = v80
        case "foo bar":
            var v81 uint64 = 508
            _ = v81
        case "":
            var v82 int32 = 0
            _ = v82
        default:
            var v83 int16 = v5
            _ = v83
        }
    case "":
        if false {
            v84 := GenS11{F0: v3, F1: 0}
            if (0 <= 1) {
                if (-1 > 1) {
                    for fi85 := 0; fi85 < 18; fi85++ {
                        v3 = (v3 << ((v4) & 63))
                        _ = fi85
                    }
                    var v86 int8 = v3
                    _ = v86
                } else {
                    _ = len(v1)
                    for fi87 := 0; fi87 < 11; fi87++ {
                        var v88 uint8 = 255
                        var v89 int64 = 9223372036854775807
                        _ = len(v2)
                        _ = fi87
                        _ = v88
                        _ = v89
                    }
                    v5 += -32768
                }
                v6 = (v6 + (v6 * 3.14159))
                var v90 uint8 = 127
                _ = v90
            } else {
                v91 := GenS2{F0: "\n\t"}
                _ = v91
            }
            for fi92 := 0; fi92 < 2; fi92++ {
                switch (("hello" + ("hello" + "a")) + "a") {
                case "a":
                    fmt.Println(v3, v3, fi92)
                case "foo bar":
                    v3 = (v3 << (((v4 + (v4 << (((v4 / ((v4) | 1))) & 63)))) & 63))
                case "":
                    _ = len(v1)
                }
                _ = fi92
            }
            _ = v84
        } else {
            fmt.Println(v6, v3, v4)
            v19 -= float64(v6)
        }
    }
    if ((!false) && true) {
        if (!(((!(42 <= 4294967295)) && false) && true)) {
            for fi93 := 0; fi93 < 4; fi93++ {
                v3 += 127
                if (("hello" + "hello") < (("hello" + ("hello" + ("\n\t" + ""))) + ("foo bar" + "\n\t"))) {
                    v4 -= (^(v4))
                    switch ("foo bar" + "") {
                    case "\n\t":
                        var v94 float32 = v6
                        _ = v94
                    case "hello":
                        if (986 >= 122) {
                            for fi95 := 0; fi95 < 6; fi95++ {
                                v96 := GenS1{F0: -410}
                                var v97 uint16 = 65535
                                _ = fi95
                                _ = v96
                                _ = v97
                            }
                        } else {
                            v1["\n\t"] = ""
                            var v98 uint16 = 0
                            for fi99 := 0; fi99 < 11; fi99++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                delete(v1, (((("\n\t" + "") + ("hello" + "")) + "foo bar") + "hello"))
                                _ = fi99
                            }
                            _ = v98
                        }
                    case "a":
                        var v100 int32 = 2147483647
                        _ = v100
                    }
                    switch "a" {
                    case "hello":
                        v19 += v19
                    default:
                        fmt.Println(v6, v6, v19)
                    }
                } else {
                    for fi101 := 0; fi101 < 16; fi101++ {
                        v102 := GenS12{F0: (^(v5)), F1: v4, F2: 4294967295}
                        _ = fi101
                        _ = v102
                    }
                    for fi103 := 0; fi103 < 12; fi103++ {
                        v19 += (v19 + v19)
                        var v104 int16 = (v5 >> (((^(v4))) & 63))
                        v1["hello"] = ""
                        _ = fi103
                        _ = v104
                    }
                }
                v105 := GenS13{F0: 4294967295, F1: "a"}
                _ = fi93
                _ = v105
            }
            var v106 int64 = -9223372036854775808
            for fi107 := 0; fi107 < 10; fi107++ {
                v19 += (v19 - v19)
                if ((64 != 42) || (fi107 <= (fi107 << ((v4) & 63)))) {
                    if (249 >= 1) {
                        if true {
                            v6 = (v6 / ((v6 - v6) / ((v6 - v6) + (1e10 + (float32(1e10) * ((v6 / v6) / v6))))))
                            v108 := GenS7{F0: (96 <= 1), F1: 482, F2: 128}
                            _ = v108
                        }
                    }
                    v106 -= (-(v106))
                }
                _ = fi107
            }
            _ = v106
        } else {
            var v109 int8 = v3
            switch "\n\t" {
            case "hello":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "\n\t":
                v3 -= 70
            case "a":
                for fi110 := 0; fi110 < 6; fi110++ {
                    v111 := GenS12{F0: (v5 >> (((v4 << ((v4) & 63))) & 63)), F1: (v4 << ((v4) & 63)), F2: 4294967295}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch ("" + "") {
                    case "hello":
                        v4 = 175
                    }
                    _ = fi110
                    _ = v111
                }
            }
            for fi112 := 0; fi112 < 16; fi112++ {
                v3 = (v3 ^ (-(v3)))
                fmt.Println(v5, v5)
                for fi113 := 0; fi113 < 18; fi113++ {
                    v114 := v1["a"]
                    _ = fi113
                    _ = v114
                }
                _ = fi112
            }
            _ = v109
        }
        if ((389 > 759) && ((v19 / ((v19 - (v19 - v19)) + (float64((v6 / v6)) + -2.25))) < float64(float32(-2.25)))) {
            fmt.Println(v3)
            v5 -= v5
            fmt.Println(v3, v19, v6)
        }
    } else {
        v115 := GenS8{F0: 446}
        v6 = v6
        _ = v115
    }
    for fi116 := 0; fi116 < 5; fi116++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v117 int16 = (v5 / ((v5) | 1))
        v6 = 0.0
        _ = fi116
        _ = v117
    }
    v6 -= (3.14159 / v6)
    for fi118 := 0; fi118 < 4; fi118++ {
        var v119 int64 = -204
        var v120 bool = true
        _ = len(v2)
        _ = fi118
        _ = v119
        _ = v120
    }
    var v121 bool = (true && ((v4 ^ (^(v4))) <= (^(v4))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v121:", v121)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v121
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
