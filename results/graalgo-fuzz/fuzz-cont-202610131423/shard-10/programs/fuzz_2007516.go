package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16; F1 uint }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    v3 := []int{0, 2147483647, -1}
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := GenS1{F0: 0, F1: "hello", F2: -1}
        var v6 string = "\n\t"
        var v7 uint8 = 2
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint32 = 2
        switch "hello" {
        case "\n\t":
            v9 -= v9
        case "":
            for fi10 := 0; fi10 < 10; fi10++ {
                v2[v9] = v9
                v9 += 2147483647
                _ = fi10
            }
        case "a":
            switch (fi8 << ((65535) & 63)) {
            case 42:
                v9 = 824
            case 128:
                if len(v3) > 0 {
                    v3[0] = (fi8 + -20)
                }
            }
        }
        var v11 string = (((("foo bar" + ("a" + ((("foo bar" + ((("" + "a") + "foo bar") + ("foo bar" + "hello"))) + "hello") + "hello"))) + ("\n\t" + "foo bar")) + "\n\t") + "\n\t")
        _ = fi8
        _ = v11
        _ = v9
    }
    var v12 int64 = -1
    switch 127 {
    case 4294967295:
        v13 := GenS1{F0: 1, F1: (("" + "\n\t") + ("foo bar" + ("foo bar" + "hello"))), F2: 311}
        _ = v13
    case 901:
        var v14 int32 = -510
        _ = v14
    case 42:
        var v15 bool = false
        _ = v15
    }
    var v16 float32 = (0.1 * 0.0)
    for fi17 := 0; fi17 < 8; fi17++ {
        for fi18 := 0; fi18 < 18; fi18++ {
            v19 := GenS1{F0: 105, F1: (("hello" + "") + (("a" + "") + "")), F2: 1}
            var v20 int8 = 100
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    _ = len(v3)
    fmt.Println(v12, v16)
    if (true && true) {
        if false {
            v16 = (float32(1.5) + v16)
        }
    } else {
        v21 := GenS2{F0: 2}
        var v22 int8 = -128
        _ = v21
        _ = v22
    }
    if true {
        switch ("" + "") {
        case "foo bar":
            for fi23 := 0; fi23 < 11; fi23++ {
                if false {
                    var v24 int16 = 1000
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 15; fi25++ {
                        v16 -= ((v16 * (v16 + 1.5)) * v16)
                        fmt.Println(fi25, v12)
                        fmt.Println(v12, v12)
                        _ = fi25
                    }
                    switch -679 {
                    case -309:
                        v26 := GenS3{F0: 42, F1: 1}
                        _ = v26
                    case 4294967295:
                        v16 = (v16 / float32(3.14159))
                    default:
                        v27 := GenS4{F0: 119}
                        _ = v27
                    }
                    var v28 float32 = (v16 + v16)
                    _ = v28
                }
                var v29 uint64 = 4294967295
                _ = fi23
                _ = v29
            }
        case "hello":
            v12 += -9223372036854775808
        case "\n\t":
            v30 := GenS3{F0: 4294967295, F1: 9223372036854775807}
            _ = v30
        }
    } else {
        switch (("" + ("" + "foo bar")) + "hello") {
        case "\n\t":
            var v31 uint32 = 1
            _ = v31
        case "hello":
            fmt.Println(v16, v16)
        case "foo bar":
            _ = len(v3)
        }
    }
    var v32 int8 = 50
    if true {
        v32 -= (v32 / ((1) | 1))
    } else {
        v12 = v12
        if (-1 <= (v12 + (v12 * v12))) {
            for fi33 := 0; fi33 < 19; fi33++ {
                v16 -= v16
                fmt.Println(fi33, fi33)
                switch "foo bar" {
                case "\n\t":
                    for fi34 := 0; fi34 < 11; fi34++ {
                        var v35 int64 = (v12 % (((v12 & (v12 + (v12 - (v12 ^ v12))))) | 1))
                        v16 -= (float32(3.14159) / v16)
                        _ = fi34
                        _ = v35
                    }
                case "foo bar":
                    if false {
                        var v36 uint32 = 4294967295
                        v37 := GenS2{F0: 1}
                        _ = v36
                        _ = v37
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                default:
                    for fi38 := 0; fi38 < 14; fi38++ {
                        var v39 uint64 = 849
                        v40 := GenS1{F0: v32, F1: "a", F2: 386}
                        if false {
                            var v41 uint32 = 350
                            v32 = 100
                            var v42 int = fi38
                            _ = v41
                            _ = v42
                        } else {
                            v32 -= (v32 + 127)
                            switch 2147483647 {
                            case 32768:
                                var v43 int64 = v12
                                _ = v43
                            }
                            switch "foo bar" {
                            case "hello":
                                if false {
                                    _ = len(v1)
                                    var v44 string = "a"
                                    _ = v44
                                }
                                v39 -= v39
                                if (float32(-2.25) < (1.5 - float32((0.0 * float64((float32(0.1) - (((v16 * v16) * (3.14159 * 0.0)) + float32(3.14159)))))))) {
                                    v45 := GenS5{F0: 100, F1: true}
                                    v12 += (v12 | (v12 / ((-239) | 1)))
                                    fmt.Println(v39)
                                    _ = v45
                                } else {
                                    var v46 int8 = v32
                                    var v47 uint32 = 4294967295
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v46
                                    _ = v47
                                }
                            case "\n\t":
                                var v48 float64 = ((1.5 + 1.5) / 0.1)
                                _ = v48
                            case "":
                                if true {
                                    var v49 uint32 = 309
                                    v39 = v39
                                    _ = v49
                                } else {
                                    if false {
                                        switch ("foo bar" + (("foo bar" + "\n\t") + "")) {
                                        case "hello":
                                            var v50 int32 = -1
                                            _ = v50
                                        case "":
                                            var v51 uint8 = 1
                                            var v52 int = (fi38 - fi33)
                                            fmt.Println(v52)
                                            _ = v51
                                            _ = v52
                                        case "foo bar":
                                            v3 = append(v3, (fi33 | (fi38 << ((621) & 63))))
                                        default:
                                            switch (fi38 << ((18446744073709551615) & 63)) {
                                            case 521:
                                                var v53 uint8 = 127
                                                _ = v53
                                            case 9223372036854775807:
                                                if true {
                                                    _ = len(v1)
                                                    v39 = (v39 / ((42) | 1))
                                                    if (18446744073709551615 < 65535) {
                                                        v12 += v12
                                                    }
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = (fi33 + 950)
                                                    }
                                                }
                                            case 128:
                                                var v54 int = fi33
                                                _ = v54
                                            }
                                        }
                                    } else {
                                        v55 := GenS3{F0: (^(v39)), F1: fi33}
                                        v16 = 3.14159
                                        if (false || (!true)) {
                                            if false {
                                                var v56 float64 = (0.0 * 0.0)
                                                _ = v56
                                            } else {
                                                _ = len(v1)
                                                v12 -= (v12 - (v12 * v12))
                                            }
                                            switch "hello" {
                                            case "a":
                                                var v57 uint = 551
                                                _ = v57
                                            case "":
                                                v12 += v12
                                            default:
                                                v12 -= v12
                                            }
                                            fmt.Println(fi33)
                                            var v58 uint16 = 0
                                            _ = v58
                                        }
                                        _ = v55
                                    }
                                    var v59 float64 = 0.0
                                    var v60 int = fi33
                                    _ = v59
                                    _ = v60
                                }
                            default:
                                switch "\n\t" {
                                case "foo bar":
                                    v16 = v16
                                case "hello":
                                    fmt.Println(fi38, v32)
                                    var v61 float32 = (1e10 * 1e10)
                                    if true {
                                        switch (fi33 >> ((uint(v39)) & 63)) {
                                        case 298:
                                            if (!(744 < -208)) {
                                                var v62 int = fi33
                                                v63 := GenS1{F0: 5, F1: ("hello" + "a"), F2: 1}
                                                var v64 float64 = 1e10
                                                _ = v62
                                                _ = v63
                                                _ = v64
                                            }
                                        case 2147483647:
                                            if false {
                                                var v65 uint32 = 559
                                                var v66 int = (fi33 | 128)
                                                v67 := GenS6{F0: (v32 >> ((775) & 63))}
                                                _ = v65
                                                _ = v66
                                                _ = v67
                                            }
                                        case 128:
                                            switch ("\n\t" + (("\n\t" + (("hello" + "") + ((("a" + "foo bar") + "a") + "\n\t"))) + "hello")) {
                                            case "hello":
                                                v61 += (v61 / 1.5)
                                            default:
                                                v68 := GenS5{F0: -1, F1: true}
                                                _ = v68
                                            }
                                        }
                                        var v69 uint64 = 1
                                        var v70 uint16 = 255
                                        var v71 uint = 2
                                        _ = v69
                                        _ = v70
                                        _ = v71
                                    } else {
                                        v39 = (v39 & 1)
                                        if false {
                                            v39 = v39
                                        }
                                        delete(v2, 2147483648)
                                    }
                                    _ = v61
                                default:
                                    if false {
                                        if (!true) {
                                            v12 -= v12
                                            fmt.Println(fi33, fi33, v32)
                                            v72 := GenS7{F0: 1000, F1: 812}
                                            _ = v72
                                        } else {
                                            var v73 uint = 1
                                            var v74 uint = uint(v39)
                                            var v75 int8 = v32
                                            _ = v73
                                            _ = v74
                                            _ = v75
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = (fi38 - (fi38 | (-(fi33))))
                                        }
                                        v12 = (v12 << ((65535) & 63))
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                }
                            }
                        }
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                }
                _ = fi33
            }
            v32 -= v32
            v76 := GenS8{F0: v12}
            _ = v76
        } else {
            switch ("\n\t" + "foo bar") {
            case "":
                if (true && (false && (!false))) {
                    v32 = -1
                    switch ("foo bar" + "foo bar") {
                    case "foo bar":
                        v77 := GenS2{F0: 2147483647}
                        _ = v77
                    }
                } else {
                    fmt.Println(v16, v16)
                    var v78 uint16 = 739
                    for fi79 := 0; fi79 < 14; fi79++ {
                        var v80 string = "\n\t"
                        if true {
                            var v81 uint = 1
                            for fi82 := 0; fi82 < 20; fi82++ {
                                v83 := GenS9{F0: v80, F1: 0}
                                _ = fi82
                                _ = v83
                            }
                            switch 42 {
                            case -820:
                                for fi84 := 0; fi84 < 8; fi84++ {
                                    switch "\n\t" {
                                    case "":
                                        v85 := v1[595]
                                        _ = v85
                                    case "a":
                                        var v86 int8 = -1
                                        _ = v86
                                    default:
                                        v12 += v12
                                        v80 = v80
                                        if (65535 <= (v78 + (v78 | v78))) {
                                            var v87 string = "foo bar"
                                            var v88 uint = (v81 >> ((v81) & 63))
                                            if false {
                                                v78 = (v78 ^ (v78 & uint16(v81)))
                                                v89 := v1[1]
                                                _ = v89
                                            }
                                            _ = v87
                                            _ = v88
                                        } else {
                                            switch (fi84 + fi84) {
                                            case 32767:
                                                v32 -= (v32 / (((-(v32))) | 1))
                                            case -1:
                                                v16 = -2.25
                                            }
                                            v90 := GenS6{F0: (v32 | v32)}
                                            _ = v90
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = fi79
                                        }
                                    }
                                    var v91 int32 = 414
                                    _ = fi84
                                    _ = v91
                                }
                            case -928:
                                switch v80 {
                                case "a":
                                    switch fi79 {
                                    case 752:
                                        if false {
                                            var v92 int32 = -2147483648
                                            _ = v92
                                        } else {
                                            if (2 == 4294967295) {
                                                var v93 int64 = (v12 % (((v12 | v12)) | 1))
                                                switch -2147483648 {
                                                case 177:
                                                    v94 := GenS5{F0: (-(v32)), F1: (32767 >= -258)}
                                                    _ = v94
                                                case 361:
                                                    switch int(v32) {
                                                    case 1:
                                                        delete(v2, 853)
                                                    case -174:
                                                        if true {
                                                            v95 := GenS1{F0: int8(v81), F1: (v80 + "foo bar"), F2: 0}
                                                            var v96 bool = true
                                                            _ = v95
                                                            _ = v96
                                                        } else {
                                                            fmt.Println(v93, v12)
                                                            if false {
                                                                var v97 int16 = 1000
                                                                switch (-(fi79)) {
                                                                case 1:
                                                                    v3 = append(v3, fi79)
                                                                }
                                                                _ = v97
                                                            }
                                                        }
                                                    }
                                                default:
                                                    var v98 float32 = (float32(1e10) - v16)
                                                    _ = v98
                                                }
                                                var v99 int = int(v93)
                                                _ = v93
                                                _ = v99
                                            }
                                            v3 = append(v3, fi79)
                                            if len(v3) > 0 {
                                                v3[0] = (fi79 ^ (^(fi79)))
                                            }
                                        }
                                    case 173:
                                        for fi100 := 0; fi100 < 13; fi100++ {
                                            fmt.Println(v12, fi79, v16)
                                            if true {
                                                fmt.Println(v81, v12)
                                                var v101 int = fi100
                                                _ = v101
                                            }
                                            _ = fi100
                                        }
                                    }
                                case "hello":
                                    for fi102 := 0; fi102 < 11; fi102++ {
                                        var v103 string = v80
                                        if false {
                                            v32 -= 1
                                            fmt.Println(v80, v80, v32)
                                            fmt.Println(v32, v81)
                                        } else {
                                            var v104 float32 = -2.25
                                            _ = v104
                                        }
                                        _ = fi102
                                        _ = v103
                                    }
                                }
                            case 32768:
                                for fi105 := 0; fi105 < 5; fi105++ {
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v106 int32 = 338
                                        _ = v106
                                    } else {
                                        v107 := v2[65535]
                                        v81 -= (v81 / ((1) | 1))
                                        v108 := v1[(v107 - v107)]
                                        _ = v107
                                        _ = v108
                                    }
                                    var v109 uint16 = 42
                                    _ = fi105
                                    _ = v109
                                }
                            }
                            _ = v81
                        }
                        _ = fi79
                        _ = v80
                    }
                    _ = v78
                }
            case "\n\t":
                fmt.Println(v16, v32)
            case "hello":
                fmt.Println(v16, v12)
            default:
                var v110 uint16 = 863
                _ = v110
            }
            _ = len(v3)
            for fi111 := 0; fi111 < 14; fi111++ {
                var v112 int32 = 42
                var v113 int16 = 0
                _ = fi111
                _ = v112
                _ = v113
            }
        }
    }
    switch "a" {
    case "a":
        var v114 float32 = (v16 / 0.1)
        _ = v114
    }
    for fi115 := 0; fi115 < 1; fi115++ {
        var v116 float64 = float64((((1.5 - (1e10 + (v16 * (v16 - v16)))) - -2.25) + (0.0 + v16)))
        v116 -= float64(float32(v116))
        fmt.Println(v16, v32)
        _ = fi115
        _ = v116
    }
    v16 = v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v32
}
