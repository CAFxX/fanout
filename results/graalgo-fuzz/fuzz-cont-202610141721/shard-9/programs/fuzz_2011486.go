package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint; F2 string }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int32; F1 int; F2 uint64 }
type GenS10 struct { F0 bool; F1 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (42 < 42)
    var v3 float32 = 1.5
    var v4 uint32 = 0
    v5 := make(map[string]bool)
    v6 := []int{128, 16, 2147483647}
    if v2 {
        if true {
            fmt.Println(v4, v2)
        } else {
            v3 -= 1.5
            v7 := GenS1{F0: -128}
            _ = v7
        }
    } else {
        v8 := GenS1{F0: 127}
        for fi9 := 0; fi9 < 19; fi9++ {
            fmt.Println(fi9, v3, v3)
            if v2 {
                if len(v6) > 0 {
                    v6[0] = (^(fi9))
                }
                v4 += (v4 / (((v4 ^ (v4 / ((v4) | 1)))) | 1))
            } else {
                var v10 int = fi9
                for fi11 := 0; fi11 < 6; fi11++ {
                    v10 = fi11
                    var v12 uint32 = v4
                    if (0 <= 0) {
                        if (1 >= 1) {
                            v10 += v10
                            fmt.Println(fi11)
                            var v13 float64 = 1e10
                            _ = v13
                        }
                    }
                    _ = fi11
                    _ = v12
                }
                if v2 {
                    v14 := GenS2{F0: (42 < -604), F1: 4294967295, F2: ("hello" + "hello")}
                    var v15 bool = (!(42 < -32768))
                    for fi16 := 0; fi16 < 19; fi16++ {
                        v4 -= 851
                        _ = fi16
                    }
                    _ = v14
                    _ = v15
                }
                _ = v10
            }
            _ = fi9
        }
        if false {
            if v2 {
                if true {
                    for fi17 := 0; fi17 < 19; fi17++ {
                        var v18 int8 = -128
                        _ = fi17
                        _ = v18
                    }
                    v1[("\n\t" + "")] = v4
                    if len(v6) > 0 {
                        v6[0] = 639
                    }
                } else {
                    fmt.Println(v2, v3, v2)
                }
                switch -530 {
                case 659:
                    if v2 {
                        var v19 int32 = 0
                        _ = v19
                    } else {
                        fmt.Println(v3)
                        if true {
                            var v20 uint16 = 736
                            fmt.Println(v3, v4)
                            _ = v20
                        } else {
                            var v21 uint32 = v4
                            v22 := v5["foo bar"]
                            _ = v21
                            _ = v22
                        }
                    }
                case 128:
                    v23 := GenS3{F0: (342 <= 2147483647), F1: 9223372036854775807}
                    _ = v23
                }
                v24 := GenS2{F0: v2, F1: 2, F2: "foo bar"}
                _ = v24
            }
        }
        _ = v8
    }
    v25 := GenS1{F0: -1}
    v26 := GenS2{F0: (!v2), F1: 65535, F2: "foo bar"}
    var v27 int = -943
    v3 -= v3
    fmt.Println(v3, v27)
    switch v27 {
    case -482:
        for fi28 := 0; fi28 < 4; fi28++ {
            var v29 uint = 18446744073709551615
            _ = fi28
            _ = v29
        }
    case 32768:
        v27 += -470
    case -863:
        _ = len(v6)
    default:
        var v30 int32 = int32(v27)
        _ = v30
    }
    if (!(164 < 127)) {
        for fi31 := 0; fi31 < 8; fi31++ {
            var v32 uint16 = 65535
            _ = fi31
            _ = v32
        }
        v33 := GenS1{F0: 32}
        _ = v33
    }
    v34 := GenS4{F0: 42, F1: false, F2: 255}
    fmt.Println(v2, v4)
    for fi35 := 0; fi35 < 7; fi35++ {
        switch (v27 - v27) {
        case 9223372036854775807:
            var v36 uint16 = 42
            _ = v36
        case 42:
            for fi37 := 0; fi37 < 8; fi37++ {
                switch ("a" + "") {
                case "\n\t":
                    var v38 float32 = 1e10
                    _ = v38
                default:
                    v39 := GenS2{F0: false, F1: 2, F2: ""}
                    _ = v39
                }
                if (!(v2 && v2)) {
                    var v40 string = "a"
                    _ = v40
                }
                if (0 > 255) {
                    switch "" {
                    case "":
                        switch (fi37 & (fi37 % ((842) | 1))) {
                        case 4294967295:
                            v41 := v1["a"]
                            _ = v41
                        case -140:
                            var v42 uint64 = 42
                            _ = v42
                        case 127:
                            v27 = 128
                        }
                    case "foo bar":
                        v6 = append(v6, -9223372036854775808)
                    }
                }
                _ = fi37
            }
        }
        switch 510 {
        case -1:
            fmt.Println(fi35)
        case 325:
            var v43 bool = (v2 && true)
            _ = v43
        default:
            if len(v6) > 0 {
                v6[0] = 283
            }
        }
        v4 = (^(v4))
        _ = fi35
    }
    var v44 string = "\n\t"
    if ((v44 > ((v44 + "hello") + "foo bar")) || true) {
        var v45 uint32 = v4
        for fi46 := 0; fi46 < 4; fi46++ {
            for fi47 := 0; fi47 < 11; fi47++ {
                fmt.Println(v2, v3)
                var v48 float64 = 0.1
                _ = len(v5)
                _ = fi47
                _ = v48
            }
            v49 := GenS2{F0: (("" + "") != "a"), F1: 520, F2: v44}
            v50 := GenS5{F0: 598, F1: 794}
            _ = fi46
            _ = v49
            _ = v50
        }
        _ = v45
    } else {
        if (true && v2) {
            v44 = v44
            if (!v2) {
                if v2 {
                    if (!true) {
                        v5[v44] = v2
                        var v51 uint8 = 45
                        _ = v51
                    } else {
                        for fi52 := 0; fi52 < 1; fi52++ {
                            v4 += v4
                            if true {
                                var v53 bool = (!v2)
                                _ = v53
                            }
                            _ = fi52
                        }
                        v2 = ((v44 + v44) == (v44 + v44))
                        for fi54 := 0; fi54 < 10; fi54++ {
                            for fi55 := 0; fi55 < 10; fi55++ {
                                if (v2 && false) {
                                    fmt.Println(v2, fi54)
                                    var v56 float64 = (0.1 - 0.0)
                                    var v57 float64 = v56
                                    _ = v56
                                    _ = v57
                                }
                                _ = fi55
                            }
                            if ("\n\t" != v44) {
                                switch v44 {
                                case "\n\t":
                                    for fi58 := 0; fi58 < 5; fi58++ {
                                        var v59 int32 = 42
                                        v2 = ((v2 && (891 == 42)) || (v2 || (3.14159 == 1.5)))
                                        _ = fi58
                                        _ = v59
                                    }
                                }
                            }
                            var v60 int64 = 1
                            _ = fi54
                            _ = v60
                        }
                    }
                    v61 := GenS6{F0: 42, F1: 2}
                    _ = v61
                } else {
                    for fi62 := 0; fi62 < 15; fi62++ {
                        if v2 {
                            if len(v6) > 0 {
                                v6[0] = -2147483648
                            }
                            if (1.5 < 3.14159) {
                                if true {
                                    switch (("foo bar" + "foo bar") + (v44 + (v44 + (v44 + v44)))) {
                                    case "\n\t":
                                        var v63 uint = 65535
                                        _ = v63
                                    case "hello":
                                        v64 := GenS1{F0: 42}
                                        _ = v64
                                    case "foo bar":
                                        v65 := v1[v44]
                                        _ = v65
                                    }
                                } else {
                                    fmt.Println(fi62, v4, fi62)
                                }
                                for fi66 := 0; fi66 < 7; fi66++ {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = fi66
                                }
                                for fi67 := 0; fi67 < 20; fi67++ {
                                    v2 = v2
                                    fmt.Println(v4, fi67, v3)
                                    v68 := v1[v44]
                                    _ = fi67
                                    _ = v68
                                }
                            } else {
                                v3 = ((((float32(float64(-2.25)) / v3) + v3) / (float32(0.1) + 0.1)) / v3)
                                v27 = (v27 / (((v27 + (fi62 * fi62))) | 1))
                                var v69 int8 = 0
                                _ = v69
                            }
                            switch (v27 + 9223372036854775807) {
                            case 773:
                                var v70 int16 = -769
                                _ = v70
                            }
                        } else {
                            var v71 float32 = v3
                            _ = v71
                        }
                        for fi72 := 0; fi72 < 3; fi72++ {
                            if (!false) {
                                switch 1 {
                                case -9223372036854775808:
                                    if false {
                                        v2 = true
                                    } else {
                                        v73 := GenS6{F0: 1, F1: 0}
                                        if true {
                                            if (!((false || ((9223372036854775807 < v27) && (1 <= -91))) && (false || true))) {
                                                v6 = append(v6, fi62)
                                                if v2 {
                                                    v4 -= (v4 | v4)
                                                    var v74 int8 = 1
                                                    v75 := GenS7{F0: 42, F1: v74, F2: 1}
                                                    _ = v74
                                                    _ = v75
                                                } else {
                                                    v2 = false
                                                    if true {
                                                        if v2 {
                                                            if (0.0 == 0.0) {
                                                                switch 127 {
                                                                case -2147483648:
                                                                    v76 := v1[(v44 + v44)]
                                                                    _ = v76
                                                                case 32768:
                                                                    v77 := GenS8{F0: 94}
                                                                    _ = v77
                                                                }
                                                            }
                                                            delete(v5, (("hello" + v44) + (v44 + v44)))
                                                        }
                                                    }
                                                    var v78 string = "\n\t"
                                                    _ = v78
                                                }
                                                v4 = v4
                                            } else {
                                                var v79 int = (v27 ^ (v27 + (fi62 + 0)))
                                                _ = v79
                                            }
                                            v3 = v3
                                        } else {
                                            _ = len(v5)
                                            v80 := GenS1{F0: 1}
                                            v81 := GenS9{F0: -2147483648, F1: fi62, F2: 1}
                                            _ = v80
                                            _ = v81
                                        }
                                        _ = v73
                                    }
                                default:
                                    v82 := GenS9{F0: 976, F1: (v27 + (v27 - 737)), F2: 566}
                                    _ = v82
                                }
                                if (v3 >= ((v3 / v3) / v3)) {
                                    if ("foo bar" <= ((v44 + v44) + "")) {
                                        switch fi72 {
                                        case 4294967295:
                                            switch "a" {
                                            case "foo bar":
                                                var v83 int16 = -32768
                                                _ = v83
                                            case "hello":
                                                fmt.Println(fi72, v3, v4)
                                            case "a":
                                                if (219 < int(v4)) {
                                                    if true {
                                                        v6 = append(v6, (v27 >> ((2) & 63)))
                                                        v84 := GenS4{F0: 0, F1: v2, F2: 65535}
                                                        v3 -= v3
                                                        _ = v84
                                                    }
                                                    var v85 int32 = 2147483647
                                                    var v86 float64 = (float64(float32((((float64(v3) / 1.5) / float64((v3 / 0.1))) + (1e10 + 0.1)))) - (((1e10 / (-2.25 * (((1e10 - 0.0) * 1.5) * 1.5))) + ((((3.14159 + 1.5) - (0.0 - (1e10 * 1e10))) - (1.5 - 0.1)) * 0.0)) + ((1.5 / float64(3.14159)) / 0.1)))
                                                    _ = v85
                                                    _ = v86
                                                } else {
                                                    v3 += 0.0
                                                    if ((float64(v3) / (float64(v3) / (0.1 - 3.14159))) != (0.0 - (((1e10 * 0.1) - float64(v3)) / ((1.5 - float64(((v3 - (v3 / 1.0)) + v3))) / (0.0 + 1.5))))) {
                                                        v87 := GenS8{F0: 216}
                                                        v4 -= v4
                                                        if (!true) {
                                                            var v88 int32 = 0
                                                            _ = v88
                                                        }
                                                        _ = v87
                                                    }
                                                    if (4294967295 > 0) {
                                                        fmt.Println(v44, fi62, v3)
                                                        fmt.Println(v3)
                                                        if v2 {
                                                            v3 = ((3.14159 * (v3 + (3.14159 + v3))) * v3)
                                                        }
                                                    }
                                                    var v89 float32 = v3
                                                    fmt.Println(v4)
                                                    _ = v89
                                                }
                                            default:
                                                var v90 int64 = -1
                                                _ = v90
                                            }
                                        default:
                                            var v91 float64 = 1e10
                                            _ = v91
                                        }
                                    }
                                    switch (fi72 % (((fi62 * v27)) | 1)) {
                                    case 528:
                                        var v92 int = fi62
                                        _ = v92
                                    case -687:
                                        delete(v1, ("hello" + "a"))
                                    }
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                v4 += 2
                            } else {
                                var v93 uint16 = 319
                                _ = v93
                            }
                            v94 := GenS4{F0: 7, F1: v2, F2: 255}
                            var v95 int16 = 1000
                            _ = fi72
                            _ = v94
                            _ = v95
                        }
                        _ = fi62
                    }
                }
                for fi96 := 0; fi96 < 17; fi96++ {
                    v97 := GenS10{F0: false, F1: 219}
                    if v2 {
                        switch v44 {
                        case "\n\t":
                            for fi98 := 0; fi98 < 11; fi98++ {
                                _ = len(v1)
                                switch v44 {
                                case "\n\t":
                                    var v99 uint32 = v4
                                    _ = v99
                                }
                                var v100 uint32 = (^(v4))
                                _ = fi98
                                _ = v100
                            }
                        case "a":
                            if (1 > 37) {
                                if (42 == 414) {
                                    if ((("hello" + "") + (v44 + v44)) == v44) {
                                        var v101 int32 = -2147483648
                                        v44 = (v44 + "hello")
                                        _ = v101
                                    } else {
                                        var v102 uint = 309
                                        _ = v102
                                    }
                                }
                                var v103 uint32 = (v4 & 4294967295)
                                v104 := GenS7{F0: 2, F1: 42, F2: 255}
                                _ = v103
                                _ = v104
                            }
                        case "":
                            if (-985 >= -869) {
                                if v2 {
                                    fmt.Println(v27, v2)
                                } else {
                                    for fi105 := 0; fi105 < 1; fi105++ {
                                        fmt.Println(v27, v27)
                                        if v2 {
                                            v106 := v1[("\n\t" + (v44 + v44))]
                                            v2 = v2
                                            _ = v106
                                        }
                                        if ((!(v2 && v2)) || true) {
                                            var v107 string = v44
                                            v3 += float32(1.5)
                                            var v108 int32 = -2147483648
                                            var v109 uint64 = 643
                                            v27 += fi105
                                            _ = len(v6)
                                            switch "\n\t" {
                                            case "a":
                                                delete(v1, "")
                                            case "":
                                                v108 -= v108
                                            case "\n\t":
                                                var v110 int16 = 0
                                                _ = v110
                                            default:
                                                v109 = (^(v109))
                                            }
                                            _ = v107
                                            _ = v108
                                            _ = v109
                                        } else {
                                            var v111 uint32 = v4
                                            var v112 int16 = -474
                                            if false {
                                                v4 = (v111 - v111)
                                                switch fi96 {
                                                case 128:
                                                    fmt.Println(v27, v44, fi105)
                                                case 556:
                                                    if (18446744073709551615 >= 536) {
                                                        var v113 uint32 = (v4 & v111)
                                                        v111 = v113
                                                        v114 := GenS1{F0: -128}
                                                        if v2 {
                                                            v115 := GenS5{F0: 65535, F1: 142}
                                                            _ = v115
                                                        } else {
                                                            v27 = 309
                                                            if false {
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                switch "hello" {
                                                                case "hello":
                                                                    if ((1e10 * float32(-2.25)) < v3) {
                                                                        if (true || (2 >= 65535)) {
                                                                            fmt.Println(fi96)
                                                                        }
                                                                    }
                                                                case "\n\t":
                                                                    v4 = v4
                                                                case "a":
                                                                    switch (v27 + (v27 + 127)) {
                                                                    case 170:
                                                                        v3 = ((v3 - (v3 - ((0.0 * v3) - 0.1))) + v3)
                                                                    case 850:
                                                                        switch v44 {
                                                                        case "hello":
                                                                            var v116 float64 = 0.1
                                                                            _ = v116
                                                                        case "":
                                                                            delete(v5, v44)
                                                                        case "a":
                                                                            v1[(("" + v44) + (v44 + v44))] = 42
                                                                        }
                                                                    case 32768:
                                                                        _ = len(v6)
                                                                    default:
                                                                        if (0 <= -1) {
                                                                            switch (fi105 >> ((65535) & 63)) {
                                                                            case 332:
                                                                                var v117 uint16 = 358
                                                                                _ = v117
                                                                            case 2147483647:
                                                                                v5[v44] = false
                                                                            default:
                                                                                if len(v6) >= 2 {
                                                                                    v6 = v6[1:len(v6)]
                                                                                }
                                                                            }
                                                                            var v118 int = (fi96 + fi96)
                                                                            v27 = v118
                                                                            _ = v118
                                                                        }
                                                                    }
                                                                }
                                                                switch v44 {
                                                                case "":
                                                                    var v119 int8 = 127
                                                                    _ = v119
                                                                case "foo bar":
                                                                    v4 = v4
                                                                default:
                                                                    v113 = v111
                                                                }
                                                            } else {
                                                                var v120 uint64 = 18446744073709551615
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                v112 = (v112 >> ((350) & 63))
                                                                v44 = (v44 + v44)
                                                                _ = v120
                                                            }
                                                        }
                                                        _ = v113
                                                        _ = v114
                                                    }
                                                default:
                                                    var v121 uint16 = 0
                                                    _ = v121
                                                }
                                            }
                                            _ = v111
                                            _ = v112
                                        }
                                        _ = fi105
                                    }
                                    v44 = "hello"
                                }
                            }
                        default:
                            v6 = append(v6, fi96)
                        }
                        var v122 int = fi96
                        v44 = (("foo bar" + v44) + v44)
                        _ = v122
                    } else {
                        var v123 uint = 0
                        _ = v123
                    }
                    _ = fi96
                    _ = v97
                }
                v6 = append(v6, (v27 << ((607) & 63)))
            } else {
                v124 := v5[v44]
                v2 = v2
                _ = v124
            }
        }
        v3 += v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
