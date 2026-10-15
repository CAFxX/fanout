package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint8; F1 bool; F2 int8 }
type GenS6 struct { F0 int64; F1 bool; F2 string }
type GenS7 struct { F0 bool; F1 string; F2 int32 }
type GenS8 struct { F0 bool; F1 int8; F2 uint32 }
type GenS9 struct { F0 uint32; F1 int32; F2 string }
type GenS10 struct { F0 int64; F1 string }
type GenS11 struct { F0 uint8; F1 uint16; F2 uint }
type GenS12 struct { F0 bool }

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    v4 := []uint32{2, 0, 1, 727}
    v5 := make(map[uint32]string)
    var v6 float32 = -2.25
    switch 32767 {
    case 9223372036854775807:
        v6 = v6
    case 2147483647:
        v7 := GenS1{F0: 55, F1: true}
        _ = v7
    default:
        var v8 int64 = 9223372036854775807
        _ = v8
    }
    if true {
        switch 4294967295 {
        case 9223372036854775807:
            var v9 int32 = 1
            _ = v9
        case 607:
            fmt.Println(v6)
        default:
            if true {
                var v10 uint8 = 127
                _ = v10
            } else {
                for fi11 := 0; fi11 < 6; fi11++ {
                    for fi12 := 0; fi12 < 3; fi12++ {
                        v6 = v6
                        var v13 int8 = 86
                        var v14 float32 = ((3.14159 / -2.25) / (v6 + 0.0))
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = fi11
                }
            }
        }
        var v15 int16 = 835
        var v16 float32 = (v6 + v6)
        _ = v15
        _ = v16
    }
    v6 -= (1.5 - float32(float64(0.1)))
    var v17 int64 = -978
    if false {
        _ = len(v1)
        if (107 >= 2) {
            v17 = v17
        }
        switch "hello" {
        case "":
            for fi18 := 0; fi18 < 14; fi18++ {
                v6 -= 1.5
                if true {
                    v19 := GenS1{F0: 0, F1: (v17 < (v17 >> ((341) & 63)))}
                    v6 = float32(1e10)
                    for fi20 := 0; fi20 < 8; fi20++ {
                        if false {
                            v17 = 1
                            var v21 uint8 = 255
                            _ = v21
                        } else {
                            if (!false) {
                                if (677 < 2147483647) {
                                    v2[""] = 2147483647
                                    switch 32767 {
                                    case 637:
                                        v6 = v6
                                    case 50:
                                        if true {
                                            v17 = v17
                                        } else {
                                            var v22 uint = 0
                                            _ = v22
                                        }
                                        v23 := GenS2{F0: 165, F1: 4294967295, F2: 4294967295}
                                        v6 = 1.5
                                        _ = v23
                                    default:
                                        v24 := GenS3{F0: (v17 - (v17 & v17)), F1: int8(fi20)}
                                        var v25 int16 = -588
                                        var v26 int64 = (-(v17))
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    }
                                }
                                if (!(246 < 2)) {
                                    var v27 int8 = 42
                                    if (true && true) {
                                        if len(v1) > 0 {
                                            v1[0] = "hello"
                                        }
                                    } else {
                                        v17 -= v17
                                        if ("\n\t" >= "") {
                                            fmt.Println(v27, fi18, v27)
                                            fmt.Println(v17, fi18)
                                        } else {
                                            v28 := GenS4{F0: "hello", F1: 2}
                                            v29 := GenS2{F0: 1, F1: 112, F2: 255}
                                            switch ("a" + (("hello" + ("" + "\n\t")) + "foo bar")) {
                                            case "foo bar":
                                                var v30 int8 = (-(v27))
                                                _ = v30
                                            }
                                            _ = v28
                                            _ = v29
                                        }
                                        fmt.Println(v17)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        fmt.Println(fi20, v27, v27)
                                    }
                                    _ = v27
                                } else {
                                    v17 = (v17 & (v17 ^ (v17 * -1)))
                                }
                            } else {
                                delete(v5, 2147483647)
                                _ = len(v1)
                            }
                            _ = len(v1)
                        }
                        if true {
                            fmt.Println(fi20, v6)
                            delete(v2, "")
                            if (!("a" == "hello")) {
                                v1 = append(v1, "")
                            } else {
                                delete(v5, 65535)
                                var v31 int16 = -656
                                _ = v31
                            }
                        } else {
                            var v32 uint32 = 2147483647
                            fmt.Println(v17, fi18)
                            _ = v32
                        }
                        if false {
                            v33 := GenS3{F0: (-(v17)), F1: -112}
                            _ = v33
                        }
                        _ = fi20
                    }
                    _ = v19
                } else {
                    _ = len(v5)
                }
                for fi34 := 0; fi34 < 16; fi34++ {
                    if (43 > 79) {
                        fmt.Println(v6, fi34)
                        v35 := GenS4{F0: ("a" + ("" + ("a" + ((("" + "a") + ("\n\t" + ("\n\t" + (("hello" + "") + ("hello" + "a"))))) + "")))), F1: 65535}
                        v17 -= (v17 ^ (v17 - -1))
                        _ = v35
                    }
                    delete(v2, ("" + ("hello" + "foo bar")))
                    if (42 < 65535) {
                        if true {
                            if (true || false) {
                                v36 := GenS5{F0: 127, F1: false, F2: 0}
                                _ = v36
                            } else {
                                var v37 float64 = -2.25
                                v6 = v6
                                _ = v37
                            }
                            if (1.5 != float32((0.1 * 1e10))) {
                                var v38 int64 = (v17 - v17)
                                v38 += -1
                                _ = v38
                            } else {
                                v6 += float32((1e10 * (3.14159 * -2.25)))
                                var v39 uint = 255
                                _ = v39
                            }
                        }
                        var v40 float32 = 3.14159
                        _ = v40
                    }
                    v41 := GenS6{F0: -981, F1: (-1 >= 1), F2: "\n\t"}
                    _ = fi34
                    _ = v41
                }
                _ = fi18
            }
        case "hello":
            v42 := GenS6{F0: 0, F1: false, F2: ""}
            _ = v42
        case "a":
            if (true && (false || true)) {
                var v43 bool = ((212 >= -1) && true)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v17 -= v17
                _ = v43
            }
        default:
            v6 = 1e10
        }
    }
    delete(v2, "foo bar")
    fmt.Println(v6)
    for fi44 := 0; fi44 < 11; fi44++ {
        switch ("foo bar" + "foo bar") {
        case "a":
            var v45 string = "a"
            _ = v45
        case "":
            var v46 int64 = (v17 % (((v17 % ((9223372036854775807) | 1))) | 1))
            _ = v46
        case "foo bar":
            fmt.Println(fi44, fi44)
        }
        _ = fi44
    }
    v6 = v6
    var v47 bool = false
    v48 := v5[0]
    switch 0 {
    case 515:
        v49 := GenS1{F0: 247, F1: (!v47)}
        _ = v49
    case 32768:
        for fi50 := 0; fi50 < 6; fi50++ {
            if v47 {
                if v47 {
                    for fi51 := 0; fi51 < 13; fi51++ {
                        switch 1 {
                        case -991:
                            if v47 {
                                var v52 int = fi51
                                if (!(1 >= -2147483648)) {
                                    v6 = (((-2.25 + float32(0.1)) - 1e10) - ((((float32((((((1e10 / -2.25) + 1e10) * 0.0) + -2.25) * 3.14159)) * v6) - (-2.25 + (v6 / ((v6 / v6) + 0.0)))) - v6) / (3.14159 * 1.5)))
                                } else {
                                    var v53 int16 = int16(v17)
                                    var v54 int = fi51
                                    fmt.Println(fi50, fi50)
                                    _ = v53
                                    _ = v54
                                }
                                if (-552 < 168) {
                                    v47 = (0 >= 115)
                                    var v55 int32 = -1
                                    var v56 string = v48
                                    switch v52 {
                                    case 62:
                                        if v47 {
                                            if v47 {
                                                switch v48 {
                                                case "a":
                                                    v6 -= (0.1 / v6)
                                                case "\n\t":
                                                    var v57 string = "hello"
                                                    _ = v57
                                                default:
                                                    v6 += ((3.14159 - float32(1e10)) * 0.1)
                                                }
                                            }
                                        }
                                    default:
                                        if v47 {
                                            v56 = v56
                                            v58 := GenS7{F0: v47, F1: (v48 + "hello"), F2: v55}
                                            _ = v58
                                        } else {
                                            switch v52 {
                                            case 0:
                                                fmt.Println(v56)
                                            case -9223372036854775808:
                                                v5[785] = (v56 + v56)
                                            default:
                                                fmt.Println(v55, v56)
                                            }
                                            v52 += (v52 * (fi51 / ((v52) | 1)))
                                            v59 := GenS1{F0: 125, F1: v47}
                                            _ = v59
                                        }
                                    }
                                    _ = v55
                                    _ = v56
                                }
                                _ = v52
                            }
                        case -9223372036854775808:
                            var v60 int32 = 42
                            _ = v60
                        }
                        v61 := GenS5{F0: 55, F1: false, F2: 42}
                        _ = fi51
                        _ = v61
                    }
                    if true {
                        if len(v4) > 0 {
                            v4[0] = 985
                        }
                        if (888 <= (fi50 | fi50)) {
                            v62 := GenS8{F0: (834 == 711), F1: 52, F2: 2147483648}
                            var v63 string = v48
                            v48 = "\n\t"
                            _ = v62
                            _ = v63
                        }
                        delete(v5, 1)
                    } else {
                        for fi64 := 0; fi64 < 1; fi64++ {
                            fmt.Println(v17, v48)
                            var v65 float32 = 1e10
                            _ = fi64
                            _ = v65
                        }
                        if v47 {
                            v48 = ""
                        } else {
                            v66 := GenS9{F0: 688, F1: 527, F2: ((v48 + v48) + v48)}
                            var v67 uint64 = 400
                            _ = v66
                            _ = v67
                        }
                        if v47 {
                            if (32767 < 42) {
                                v17 -= v17
                                var v68 bool = v47
                                var v69 int32 = -346
                                _ = v68
                                _ = v69
                            } else {
                                v70 := GenS1{F0: 157, F1: v47}
                                v47 = v47
                                _ = v70
                            }
                        }
                    }
                }
            } else {
                for fi71 := 0; fi71 < 7; fi71++ {
                    v6 = v6
                    _ = fi71
                }
            }
            fmt.Println(v48, v6, v6)
            _ = fi50
        }
    case 42:
        switch "" {
        case "foo bar":
            if v47 {
                if (v6 >= -2.25) {
                    v72 := GenS7{F0: v47, F1: v48, F2: 2147483647}
                    switch 128 {
                    case 1:
                        var v73 float32 = (v6 * float32(0.1))
                        _ = v73
                    case 32768:
                        switch "hello" {
                        case "a":
                            if len(v4) > 0 {
                                v4[0] = 2147483648
                            }
                        }
                    case 4294967295:
                        v74 := GenS8{F0: true, F1: 1, F2: 0}
                        _ = v74
                    }
                    var v75 string = "foo bar"
                    _ = v72
                    _ = v75
                }
                for fi76 := 0; fi76 < 14; fi76++ {
                    v77 := v2[""]
                    if v47 {
                        fmt.Println(v47, fi76, v48)
                        for fi78 := 0; fi78 < 18; fi78++ {
                            if ("foo bar" <= ("a" + "\n\t")) {
                                v79 := v2[(v48 + v48)]
                                fmt.Println(v47, v48, fi78)
                                _ = v79
                            } else {
                                v48 = "hello"
                            }
                            v80 := GenS10{F0: v17, F1: v48}
                            switch (("foo bar" + "hello") + (("\n\t" + (v48 + v48)) + "hello")) {
                            case "hello":
                                var v81 string = v48
                                _ = v81
                            case "a":
                                var v82 float64 = ((0.0 / -2.25) + -2.25)
                                _ = v82
                            }
                            v47 = (!(v6 != v6))
                            v83 := GenS2{F0: v77, F1: 0, F2: 0}
                            _ = fi78
                            _ = v80
                            _ = v83
                        }
                        if (-2.25 > float64(float32(3.14159))) {
                            if ((v6 - ((v6 - (((((0.0 - 1e10) + -2.25) * v6) * v6) / v6)) / v6)) > v6) {
                                v84 := GenS11{F0: 128, F1: 2, F2: 0}
                                v85 := GenS12{F0: (42 < 0)}
                                for fi86 := 0; fi86 < 7; fi86++ {
                                    var v87 int = 42
                                    var v88 int = v87
                                    switch (fi76 * fi86) {
                                    case 42:
                                        var v89 int = (-(v88))
                                        _ = v89
                                    case -784:
                                        var v90 uint16 = 255
                                        _ = v90
                                    default:
                                        var v91 bool = ("" == v48)
                                        _ = v91
                                    }
                                    _ = fi86
                                    _ = v87
                                    _ = v88
                                }
                                _ = v84
                                _ = v85
                            } else {
                                v2[("hello" + ("foo bar" + (v48 + (v48 + v48))))] = v77
                                var v92 uint8 = 127
                                v93 := GenS4{F0: "hello", F1: 2}
                                _ = v92
                                _ = v93
                            }
                        }
                    } else {
                        switch ((v48 + v48) + "a") {
                        case "foo bar":
                            v94 := GenS1{F0: 42, F1: (false || v47)}
                            _ = v94
                        default:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v77 -= 65535
                        var v95 int32 = -120
                        _ = v95
                    }
                    fmt.Println(v17, v6)
                    _ = fi76
                    _ = v77
                }
            } else {
                if (!(v48 > v48)) {
                    for fi96 := 0; fi96 < 3; fi96++ {
                        v3[4294967295] = ((-(fi96)) <= (fi96 - fi96))
                        _ = fi96
                    }
                    var v97 uint = 703
                    if true {
                        var v98 int = -428
                        var v99 int32 = -1
                        var v100 bool = (32767 <= 318)
                        _ = v100
                        _ = v98
                        _ = v99
                    }
                    _ = v97
                } else {
                    v6 += v6
                    for fi101 := 0; fi101 < 18; fi101++ {
                        fmt.Println(v6, v47)
                        _ = fi101
                    }
                    v48 = v48
                }
                v102 := GenS5{F0: 128, F1: v47, F2: 1}
                _ = v102
            }
        }
    }
    v17 = v17
    for fi103 := 0; fi103 < 9; fi103++ {
        fmt.Println(v48)
        _ = fi103
    }
    if (false && (255 >= 4294967295)) {
        for fi104 := 0; fi104 < 5; fi104++ {
            v6 += float32(3.14159)
            v105 := GenS7{F0: (255 > 42), F1: (v48 + v48), F2: int32(fi104)}
            for fi106 := 0; fi106 < 16; fi106++ {
                v107 := GenS7{F0: v47, F1: "a", F2: -2147483648}
                _ = fi106
                _ = v107
            }
            _ = fi104
            _ = v105
        }
        var v108 float64 = (1.5 - -2.25)
        var v109 float64 = (v108 + 0.0)
        _ = v108
        _ = v109
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
