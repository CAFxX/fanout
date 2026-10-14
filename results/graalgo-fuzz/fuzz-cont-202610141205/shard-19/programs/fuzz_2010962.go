package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string; F1 uint32; F2 int16 }
type GenS9 struct { F0 int8; F1 string }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS13 struct { F0 int; F1 uint16; F2 uint8 }
type GenS14 struct { F0 int; F1 uint64 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int; F1 uint64; F2 uint32 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 uint8 }
type GenS19 struct { F0 int32; F1 uint; F2 uint }
type GenS20 struct { F0 string; F1 int32; F2 uint8 }
type GenS21 struct { F0 uint16; F1 string; F2 string }
type GenS22 struct { F0 uint8; F1 int; F2 bool }
type GenS23 struct { F0 uint32; F1 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    var v3 int8 = 100
    fmt.Println(v3)
    v4 := GenS1{F0: ((("" + "") + "") + ("\n\t" + "\n\t")), F1: 128, F2: "foo bar"}
    var v5 uint = 0
    fmt.Println(v5, v5)
    v5 = v5
    v6 := []int{32767}
    var v7 float32 = 1.5
    v7 -= v7
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            switch (fi9 >> ((v5) & 63)) {
            case -9223372036854775808:
                if len(v6) > 0 {
                    v6[0] = int(v3)
                }
            case 128:
                if (1 <= -797) {
                    v10 := GenS1{F0: (("a" + "a") + ((("" + ("foo bar" + "foo bar")) + "foo bar") + ("\n\t" + "hello"))), F1: 255, F2: "a"}
                    v11 := v2[""]
                    var v12 int16 = 32767
                    _ = v10
                    _ = v11
                    _ = v12
                }
                if (false && false) {
                    v7 = (float32(float64((v7 * v7))) + -2.25)
                    v5 = (^(v5))
                    var v13 int64 = 9223372036854775807
                    v7 = -2.25
                    if false {
                        var v14 uint8 = 0
                        fmt.Println(v14)
                        if false {
                            fmt.Println(fi9, fi9)
                            if (4294967295 != 0) {
                                v5 = (v5 & (v5 % ((v5) | 1)))
                                if len(v6) > 0 {
                                    v6[0] = 713
                                }
                            }
                            v5 += 255
                        }
                        var v15 uint16 = 1
                        if true {
                            v7 = 0.1
                            v7 = v7
                        } else {
                            if false {
                                var v16 bool = false
                                var v17 int = 42
                                _ = v16
                                _ = v17
                            } else {
                                var v18 uint16 = 418
                                _ = v18
                            }
                            v3 -= (v3 >> ((v5) & 63))
                        }
                        _ = v14
                        _ = v15
                    }
                    if ((0 >= 42) || false) {
                        v5 = v5
                    }
                    _ = v13
                } else {
                    var v19 int64 = 185
                    if len(v6) > 0 {
                        v6[0] = (fi9 | (fi8 | (fi9 ^ (^(fi9)))))
                    }
                    v19 -= v19
                    _ = v19
                }
                v7 = (1e10 * v7)
            }
            var v20 int8 = v3
            _ = fi9
            _ = v20
        }
        _ = fi8
    }
    var v21 float64 = 1.5
    v2["foo bar"] = 310
    switch "\n\t" {
    case "\n\t":
        switch "\n\t" {
        case "":
            var v22 uint16 = 65535
            _ = v22
        case "foo bar":
            if false {
                var v23 uint8 = 1
                _ = v23
            }
        case "hello":
            switch 509 {
            case 128:
                var v24 int8 = v3
                _ = v24
            case 9223372036854775807:
                fmt.Println(v5)
            case 42:
                for fi25 := 0; fi25 < 1; fi25++ {
                    if false {
                        v3 = -5
                        for fi26 := 0; fi26 < 9; fi26++ {
                            if ((("" == "") && (true && (false && ((1 == 4294967295) || true)))) || ((2 >= 42) && true)) {
                                v3 = -1
                            } else {
                                var v27 uint64 = 18446744073709551615
                                _ = v27
                            }
                            _ = fi26
                        }
                    }
                    if len(v6) > 0 {
                        v6[0] = (fi25 >> (((^(v5))) & 63))
                    }
                    var v28 int16 = -32768
                    _ = fi25
                    _ = v28
                }
            }
        default:
            if false {
                v29 := GenS1{F0: "foo bar", F1: 42, F2: ("a" + "foo bar")}
                v21 = (float64(((float32((v21 - 0.1)) * (float32(float64(0.0)) + ((((((0.1 - v7) - (float32((((1e10 / float64((float32(((1e10 + v21) * v21)) - float32(1e10)))) / (0.0 + float64(0.1))) - 0.0)) / v7)) * v7) / v7) / v7) + (v7 + (v7 * v7))))) + 3.14159)) - (1.5 - v21))
                _ = v29
            } else {
                if true {
                    if (!(v7 > ((-2.25 / v7) / v7))) {
                        v21 -= -2.25
                        for fi30 := 0; fi30 < 7; fi30++ {
                            v31 := GenS2{F0: 1}
                            v21 += 1.5
                            var v32 string = ("\n\t" + (("foo bar" + ("a" + "foo bar")) + ("foo bar" + ("\n\t" + "foo bar"))))
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        v33 := GenS1{F0: "a", F1: 183, F2: "hello"}
                        _ = v33
                    } else {
                        for fi34 := 0; fi34 < 13; fi34++ {
                            if ((!(false || (!(276 < v5)))) || true) {
                                switch ("foo bar" + ("a" + "")) {
                                case "hello":
                                    for fi35 := 0; fi35 < 15; fi35++ {
                                        var v36 int8 = 100
                                        fmt.Println(fi35, fi35)
                                        v37 := GenS1{F0: "", F1: 1, F2: ("\n\t" + "a")}
                                        switch -2147483648 {
                                        case 9223372036854775807:
                                            var v38 int = (fi34 % ((fi35) | 1))
                                            _ = v38
                                        case -673:
                                            if true {
                                                var v39 uint32 = 1
                                                if (v21 > float64(0.0)) {
                                                    delete(v1, 65535)
                                                    fmt.Println(v21, v5)
                                                }
                                                v40 := v2["foo bar"]
                                                var v41 uint32 = (v40 & v39)
                                                var v42 uint16 = 783
                                                _ = v39
                                                _ = v40
                                                _ = v41
                                                _ = v42
                                            } else {
                                                var v43 int64 = 0
                                                var v44 int64 = v43
                                                _ = v43
                                                _ = v44
                                            }
                                        default:
                                            var v45 bool = (v21 > (1.5 + v21))
                                            _ = v45
                                        }
                                        var v46 int16 = -1
                                        var v47 uint = v5
                                        var v48 uint64 = 18446744073709551615
                                        v47 -= (v47 & v47)
                                        fmt.Println(v47, v5, v21)
                                        _ = fi35
                                        _ = v36
                                        _ = v37
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    }
                                case "foo bar":
                                    if len(v6) > 0 {
                                        v6[0] = fi34
                                    }
                                case "a":
                                    v3 += (v3 + 47)
                                }
                            } else {
                                var v49 int16 = 42
                                _ = v49
                            }
                            _ = fi34
                        }
                    }
                }
                var v50 string = "\n\t"
                if false {
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        switch (v50 + "") {
                        case "a":
                            if (v7 > (v7 * (-2.25 * v7))) {
                                switch "a" {
                                case "\n\t":
                                    var v51 int16 = 42
                                    _ = v51
                                case "":
                                    _ = len(v2)
                                }
                                v52 := GenS3{F0: -405, F1: 128}
                                v5 += v5
                                _ = v52
                            }
                        default:
                            v3 = -128
                        }
                    case "a":
                        var v53 uint8 = 0
                        _ = v53
                    case "":
                        v54 := GenS4{F0: (v50 + v50), F1: v50}
                        _ = v54
                    default:
                        for fi55 := 0; fi55 < 2; fi55++ {
                            v21 = v21
                            _ = fi55
                        }
                    }
                    var v56 int16 = 32767
                    switch "" {
                    case "foo bar":
                        fmt.Println(v5, v7)
                    case "hello":
                        fmt.Println(v50, v56, v50)
                    case "a":
                        if len(v6) > 0 {
                            v6[0] = -1
                        }
                    }
                    _ = v56
                } else {
                    if true {
                        for fi57 := 0; fi57 < 13; fi57++ {
                            fmt.Println(fi57, v5, v7)
                            _ = fi57
                        }
                        switch 42 {
                        case 0:
                            var v58 uint64 = 18446744073709551615
                            _ = v58
                        }
                    } else {
                        var v59 int8 = (v3 >> ((v5) & 63))
                        _ = v59
                    }
                }
                _ = v50
            }
        }
    }
    switch "hello" {
    case "":
        var v60 uint32 = 1
        _ = v60
    case "a":
        for fi61 := 0; fi61 < 17; fi61++ {
            for fi62 := 0; fi62 < 20; fi62++ {
                v7 -= float32(float64((float32(float64((-2.25 / v7))) - float32(float64(0.0)))))
                if false {
                    var v63 float64 = float64(v7)
                    if len(v6) > 0 {
                        v6[0] = fi62
                    }
                    if (1000 < 1000) {
                        var v64 uint = (^(v5))
                        switch 32768 {
                        case 2147483647:
                            var v65 uint = 2
                            _ = v65
                        case -185:
                            if (1 < 0) {
                                _ = len(v6)
                                v66 := GenS3{F0: (fi62 | fi61), F1: 42}
                                _ = v66
                            }
                        case 32768:
                            var v67 int8 = 100
                            _ = v67
                        }
                        switch "" {
                        case "":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v68 := GenS1{F0: ("\n\t" + (("a" + ("a" + "\n\t")) + ("\n\t" + "a"))), F1: 2, F2: "foo bar"}
                            _ = v68
                        case "foo bar":
                            v63 += (-2.25 - v63)
                        case "\n\t":
                            var v69 int64 = 19
                            _ = v69
                        default:
                            var v70 int = (fi61 * (fi61 << (((v64 % ((18446744073709551615) | 1))) & 63)))
                            _ = v70
                        }
                        _ = v64
                    } else {
                        switch ("" + "") {
                        case "\n\t":
                            switch ("foo bar" + "") {
                            case "foo bar":
                                if (!(float64(((v7 * v7) - v7)) >= (v21 * float64(((1.5 + v7) * v7))))) {
                                    v71 := GenS5{F0: 255}
                                    if (v5 >= (v5 / (((^(v5))) | 1))) {
                                        if true {
                                            var v72 int8 = (v3 | (^(v3)))
                                            if false {
                                                if true {
                                                    var v73 bool = false
                                                    if (1 != 490) {
                                                        v63 += float64((float32((v21 / (-2.25 + float64(((3.14159 / (v7 * (v7 - v7))) / ((v7 - v7) / v7)))))) - 1e10))
                                                    }
                                                    _ = v73
                                                } else {
                                                    if false {
                                                        if true {
                                                            var v74 int16 = 890
                                                            _ = v74
                                                        } else {
                                                            _ = len(v6)
                                                            var v75 uint8 = 35
                                                            var v76 int32 = 2147483647
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            _ = len(v6)
                                                        case "a":
                                                            if (true || false) {
                                                                fmt.Println(fi61)
                                                                var v77 int = fi62
                                                                v72 = v3
                                                                _ = v77
                                                            } else {
                                                                v78 := GenS4{F0: "\n\t", F1: "hello"}
                                                                _ = v78
                                                            }
                                                        case "":
                                                            var v79 uint32 = 2147483648
                                                            switch int(v79) {
                                                            case -277:
                                                                v6 = append(v6, fi61)
                                                            case 0:
                                                                if (!true) {
                                                                    var v80 string = (((("hello" + ((("" + (("hello" + "") + "a")) + "") + "")) + ("\n\t" + "a")) + ("" + ("hello" + "foo bar"))) + "hello")
                                                                    v3 = 1
                                                                    var v81 int64 = -9223372036854775808
                                                                    v82 := GenS2{F0: -1}
                                                                    v1[(^(v79))] = ("\n\t" + ((v80 + "a") + ""))
                                                                    _ = v80
                                                                    _ = v81
                                                                    _ = v82
                                                                }
                                                                fmt.Println(fi61, v3)
                                                            case 1:
                                                                if true {
                                                                    v83 := GenS6{F0: 42}
                                                                    var v84 uint8 = 0
                                                                    v72 = 11
                                                                    _ = v83
                                                                    _ = v84
                                                                }
                                                            }
                                                            v85 := GenS7{F0: 0}
                                                            v79 -= (v79 + (v79 % ((v79) | 1)))
                                                            _ = v79
                                                            _ = v85
                                                        default:
                                                            var v86 int32 = 0
                                                            _ = v86
                                                        }
                                                        v87 := GenS7{F0: 0}
                                                        v5 = (v5 % (((^(v5))) | 1))
                                                        _ = v87
                                                    } else {
                                                        v88 := GenS8{F0: "foo bar", F1: 640, F2: 1}
                                                        if (0 < -1) {
                                                            delete(v1, 42)
                                                            _ = len(v1)
                                                            v3 -= int8(v5)
                                                            v1[2147483647] = ""
                                                            switch "foo bar" {
                                                            case "a":
                                                                switch -9223372036854775808 {
                                                                case 127:
                                                                    switch fi62 {
                                                                    case 900:
                                                                        _ = len(v6)
                                                                    }
                                                                default:
                                                                    v7 = v7
                                                                }
                                                                switch ("hello" + "a") {
                                                                case "\n\t":
                                                                    fmt.Println(v21)
                                                                case "hello":
                                                                    v63 = v21
                                                                case "foo bar":
                                                                    v89 := GenS1{F0: "hello", F1: 127, F2: (("\n\t" + ("a" + "hello")) + "hello")}
                                                                    if true {
                                                                        if false {
                                                                            var v90 uint16 = 7
                                                                            _ = v90
                                                                        } else {
                                                                            fmt.Println(fi62, v21, v7)
                                                                        }
                                                                        var v91 string = ""
                                                                        if (true || true) {
                                                                            v92 := GenS5{F0: (v5 - v5)}
                                                                            var v93 uint16 = 0
                                                                            var v94 uint8 = 42
                                                                            var v95 uint32 = 927
                                                                            v95 = (^(v95))
                                                                            _ = v92
                                                                            _ = v93
                                                                            _ = v94
                                                                            _ = v95
                                                                        } else {
                                                                            switch (fi61 << ((v5) & 63)) {
                                                                            case -1:
                                                                                var v96 uint64 = 1
                                                                                _ = v96
                                                                            case 42:
                                                                                var v97 uint16 = 1
                                                                                _ = v97
                                                                            }
                                                                        }
                                                                        _ = v91
                                                                    }
                                                                    fmt.Println(v5, v63)
                                                                    _ = v89
                                                                default:
                                                                    if false {
                                                                        var v98 uint64 = 18446744073709551615
                                                                        switch "a" {
                                                                        case "a":
                                                                            v5 -= (v5 | (v5 - (v5 >> (((v5 ^ (v5 % ((v5) | 1)))) & 63))))
                                                                        }
                                                                        _ = v98
                                                                    } else {
                                                                        switch ("a" + "hello") {
                                                                        case "\n\t":
                                                                            switch "a" {
                                                                            case "a":
                                                                                if true {
                                                                                    v99 := GenS9{F0: (v3 + v72), F1: "\n\t"}
                                                                                    _ = v99
                                                                                }
                                                                            case "foo bar":
                                                                                fmt.Println(fi61)
                                                                            case "":
                                                                                var v100 uint64 = 42
                                                                                _ = v100
                                                                            default:
                                                                                var v101 uint64 = 2
                                                                                _ = v101
                                                                            }
                                                                        case "foo bar":
                                                                            fmt.Println(v7, v3, v5)
                                                                            var v102 float32 = float32(3.14159)
                                                                            _ = v102
                                                                        }
                                                                        delete(v2, ("a" + ("\n\t" + "foo bar")))
                                                                        v63 = (v63 / ((v63 + v63) * -2.25))
                                                                    }
                                                                }
                                                            case "hello":
                                                                if ((false && true) && true) {
                                                                    if true {
                                                                        var v103 int16 = 310
                                                                        _ = v103
                                                                    } else {
                                                                        v104 := GenS3{F0: 824, F1: 46}
                                                                        if (1 != 2) {
                                                                            if (!false) {
                                                                                var v105 uint64 = 0
                                                                                v106 := GenS10{F0: -35}
                                                                                _ = v105
                                                                                _ = v106
                                                                            } else {
                                                                                v107 := GenS11{F0: "foo bar"}
                                                                                v7 = 1e10
                                                                                fmt.Println(v5)
                                                                                _ = v107
                                                                            }
                                                                        }
                                                                        v108 := GenS10{F0: -1}
                                                                        _ = v104
                                                                        _ = v108
                                                                    }
                                                                } else {
                                                                    var v109 bool = true
                                                                    _ = v109
                                                                }
                                                            default:
                                                                var v110 int = 2147483647
                                                                _ = v110
                                                            }
                                                            v2[(("hello" + "foo bar") + "hello")] = 42
                                                            v21 = v21
                                                            var v111 int8 = -1
                                                            var v112 int = 728
                                                            switch 42 {
                                                            case 1:
                                                                v7 = (0.0 - float32(v63))
                                                            default:
                                                                switch fi61 {
                                                                case 42:
                                                                    v111 += v111
                                                                    switch "hello" {
                                                                    case "":
                                                                        v21 = v21
                                                                    case "\n\t":
                                                                        v113 := GenS12{F0: -247, F1: 223, F2: 255}
                                                                        _ = v113
                                                                    }
                                                                case 9223372036854775807:
                                                                    v6 = append(v6, v112)
                                                                case -9223372036854775808:
                                                                    if true {
                                                                        v111 = v111
                                                                    }
                                                                default:
                                                                    if ((false || (true || (!false))) || false) {
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            _ = len(v1)
                                                                        case "a":
                                                                            v114 := GenS13{F0: (fi62 + fi62), F1: 11, F2: 210}
                                                                            _ = v114
                                                                        case "":
                                                                            fmt.Println(fi61)
                                                                        default:
                                                                            switch "a" {
                                                                            case "a":
                                                                                v115 := GenS14{F0: v112, F1: 2}
                                                                                _ = v115
                                                                            case "foo bar":
                                                                                _ = len(v1)
                                                                            default:
                                                                                var v116 uint32 = uint32(v5)
                                                                                _ = v116
                                                                            }
                                                                        }
                                                                        v117 := GenS4{F0: "a", F1: ("hello" + "\n\t")}
                                                                        var v118 int = int(v111)
                                                                        _ = v117
                                                                        _ = v118
                                                                    }
                                                                }
                                                            }
                                                            if (127 == 42) {
                                                                v7 += 0.0
                                                                if true {
                                                                    v119 := GenS10{F0: 517}
                                                                    _ = v119
                                                                }
                                                                var v120 uint16 = 1
                                                                _ = v120
                                                            } else {
                                                                fmt.Println(v112, v72, v63)
                                                                if ((2147483647 >= int32(v5)) || true) {
                                                                    switch ((("foo bar" + "hello") + "a") + "foo bar") {
                                                                    case "foo bar":
                                                                        delete(v1, 193)
                                                                    case "a":
                                                                        delete(v1, 42)
                                                                    }
                                                                    v121 := GenS1{F0: "\n\t", F1: 0, F2: ""}
                                                                    _ = v121
                                                                } else {
                                                                    v111 = 0
                                                                    v122 := GenS1{F0: "", F1: 125, F2: "\n\t"}
                                                                    fmt.Println(v111, v5)
                                                                    _ = v122
                                                                }
                                                                var v123 int32 = 1
                                                                v7 += 1e10
                                                                v124 := GenS14{F0: 2147483647, F1: 4294967295}
                                                                switch ("" + ("" + "hello")) {
                                                                case "hello":
                                                                    v7 += v7
                                                                }
                                                                _ = v123
                                                                _ = v124
                                                            }
                                                            _ = v111
                                                            _ = v112
                                                        } else {
                                                            fmt.Println(v7)
                                                            var v125 float64 = 1e10
                                                            var v126 string = (("foo bar" + ("hello" + ((("hello" + "foo bar") + "hello") + "a"))) + "hello")
                                                            switch (fi62 - fi62) {
                                                            case 0:
                                                                if ((!(!false)) && true) {
                                                                    v127 := GenS14{F0: (fi62 >> ((v5) & 63)), F1: 355}
                                                                    _ = v127
                                                                }
                                                            case 42:
                                                                fmt.Println(fi62)
                                                            case 32768:
                                                                v5 = (v5 * 42)
                                                            }
                                                            _ = v125
                                                            _ = v126
                                                        }
                                                        var v128 uint = (v5 | (v5 ^ 18446744073709551615))
                                                        _ = v128
                                                        _ = v88
                                                    }
                                                }
                                            } else {
                                                if false {
                                                    v129 := GenS13{F0: fi61, F1: 32767, F2: 128}
                                                    _ = v129
                                                }
                                                v2[""] = 0
                                            }
                                            switch ("hello" + "foo bar") {
                                            case "foo bar":
                                                v130 := GenS2{F0: -2147483648}
                                                _ = v130
                                            case "a":
                                                var v131 uint16 = 65535
                                                _ = v131
                                            }
                                            _ = v72
                                        }
                                        var v132 uint16 = 1
                                        var v133 int64 = 370
                                        _ = v132
                                        _ = v133
                                    } else {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        if true {
                                            v3 += v3
                                            switch "\n\t" {
                                            case "\n\t":
                                                if false {
                                                    if false {
                                                        fmt.Println(v63, fi62, v5)
                                                    } else {
                                                        fmt.Println(fi62, v63, v21)
                                                        var v134 uint = 255
                                                        _ = v134
                                                    }
                                                } else {
                                                    v7 -= v7
                                                    v3 += (v3 >> ((v5) & 63))
                                                    v5 += 0
                                                }
                                            }
                                            if false {
                                                var v135 int8 = (v3 % ((v3) | 1))
                                                switch fi62 {
                                                case 32768:
                                                    v136 := GenS14{F0: int(v5), F1: 0}
                                                    _ = v136
                                                case 128:
                                                    v63 = v63
                                                case 1:
                                                    v137 := GenS15{F0: 2}
                                                    _ = v137
                                                }
                                                var v138 bool = false
                                                fmt.Println(v5)
                                                _ = v135
                                                _ = v138
                                            } else {
                                                switch -2147483648 {
                                                case -1:
                                                    if (1 <= 1) {
                                                        if (1 >= 0) {
                                                            var v139 float32 = (float32((v63 - v21)) + float32((float64(-2.25) / 3.14159)))
                                                            _ = v139
                                                        }
                                                        var v140 bool = false
                                                        var v141 float64 = float64(1e10)
                                                        var v142 string = "\n\t"
                                                        v140 = (32767 == 207)
                                                        switch v142 {
                                                        case "":
                                                            switch -273 {
                                                            case 127:
                                                                var v143 float32 = (0.1 * -2.25)
                                                                _ = v143
                                                            case 128:
                                                                var v144 int16 = -1
                                                                _ = v144
                                                            case 0:
                                                                switch (^(fi62)) {
                                                                case 42:
                                                                    v140 = true
                                                                case -844:
                                                                    v145 := v1[uint32(v5)]
                                                                    _ = v145
                                                                }
                                                            default:
                                                                fmt.Println(v63, v21, v140)
                                                            }
                                                        default:
                                                            v146 := GenS16{F0: fi61, F1: 4294967295, F2: 42}
                                                            _ = v146
                                                        }
                                                        _ = v140
                                                        _ = v141
                                                        _ = v142
                                                    }
                                                case -9223372036854775808:
                                                    if false {
                                                        var v147 int64 = -761
                                                        var v148 uint8 = 1
                                                        v149 := GenS9{F0: v3, F1: "a"}
                                                        _ = v147
                                                        _ = v148
                                                        _ = v149
                                                    } else {
                                                        if false {
                                                            var v150 uint16 = 32767
                                                            v151 := GenS7{F0: (v3 / (((v3 / ((42) | 1))) | 1))}
                                                            _ = v150
                                                            _ = v151
                                                        }
                                                        switch fi62 {
                                                        case 713:
                                                            var v152 uint64 = 0
                                                            _ = v152
                                                        case 2147483647:
                                                            var v153 int32 = 0
                                                            if true {
                                                                v153 += v153
                                                            } else {
                                                                _ = len(v6)
                                                                switch 32767 {
                                                                case 91:
                                                                    var v154 uint32 = 2147483647
                                                                    _ = v154
                                                                default:
                                                                    v155 := GenS17{F0: 2}
                                                                    _ = v155
                                                                }
                                                                v3 = -1
                                                                switch (fi61 | fi61) {
                                                                case -134:
                                                                    var v156 int64 = int64(v3)
                                                                    _ = v156
                                                                case 0:
                                                                    fmt.Println(v5, fi62, fi62)
                                                                }
                                                            }
                                                            var v157 uint32 = 4
                                                            if true {
                                                                v3 -= v3
                                                                switch ("a" + "foo bar") {
                                                                case "hello":
                                                                    var v158 string = "a"
                                                                    _ = v158
                                                                default:
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                }
                                                                if false {
                                                                    v159 := GenS18{F0: 255}
                                                                    var v160 float64 = -2.25
                                                                    v161 := GenS6{F0: 1}
                                                                    _ = v159
                                                                    _ = v160
                                                                    _ = v161
                                                                }
                                                            }
                                                            var v162 int32 = v153
                                                            _ = v153
                                                            _ = v157
                                                            _ = v162
                                                        }
                                                    }
                                                case 2147483647:
                                                    v63 = 1e10
                                                default:
                                                    v21 = v63
                                                }
                                                fmt.Println(v21, v3)
                                            }
                                        } else {
                                            if true {
                                                switch (fi61 & (fi61 & (^(fi62)))) {
                                                case 830:
                                                    switch ("" + "a") {
                                                    case "\n\t":
                                                        _ = len(v2)
                                                    case "":
                                                        _ = len(v6)
                                                    }
                                                case 127:
                                                    v21 = v63
                                                    var v163 int8 = -49
                                                    _ = v163
                                                default:
                                                    var v164 float64 = 0.1
                                                    _ = v164
                                                }
                                                var v165 int8 = (v3 ^ (v3 ^ v3))
                                                _ = v165
                                            } else {
                                                var v166 float64 = v63
                                                if len(v6) > 0 {
                                                    v6[0] = (fi62 | fi62)
                                                }
                                                var v167 int32 = -1
                                                if (v167 > int32(v3)) {
                                                    var v168 float64 = (v166 * v21)
                                                    v169 := GenS3{F0: (fi62 - (-(fi62))), F1: 128}
                                                    delete(v2, "foo bar")
                                                    switch fi61 {
                                                    case 197:
                                                        switch fi62 {
                                                        case 2147483647:
                                                            if false {
                                                                v1[1] = ("foo bar" + "foo bar")
                                                                if (!true) {
                                                                    switch fi61 {
                                                                    case 387:
                                                                        if ("\n\t" <= "a") {
                                                                            if (2 != 42) {
                                                                                fmt.Println(v167)
                                                                                var v170 float32 = v7
                                                                                v3 = v3
                                                                                var v171 uint32 = 42
                                                                                _ = v170
                                                                                _ = v171
                                                                            } else {
                                                                                if ((v167 << (((v5 | (v5 - (^(v5))))) & 63)) == (v167 & v167)) {
                                                                                    if len(v6) > 0 {
                                                                                        v6[0] = -2147483648
                                                                                    }
                                                                                }
                                                                                _ = len(v6)
                                                                                if true {
                                                                                    if false {
                                                                                        if (v7 > v7) {
                                                                                            _ = len(v2)
                                                                                            var v172 bool = (v63 < (v166 * v63))
                                                                                            v167 += v167
                                                                                            _ = v172
                                                                                        } else {
                                                                                            if (18446744073709551615 > v5) {
                                                                                                var v173 uint = (v5 ^ v5)
                                                                                                _ = v173
                                                                                            } else {
                                                                                                var v174 uint32 = 2
                                                                                                if true {
                                                                                                    var v175 string = "hello"
                                                                                                    _ = len(v2)
                                                                                                    var v176 uint = 65535
                                                                                                    _ = v175
                                                                                                    _ = v176
                                                                                                }
                                                                                                _ = v174
                                                                                            }
                                                                                            var v177 uint64 = uint64(v5)
                                                                                            _ = v177
                                                                                        }
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            switch fi62 {
                                                                                            case 42:
                                                                                                var v178 uint8 = 128
                                                                                                _ = v178
                                                                                            case -411:
                                                                                                fmt.Println(v168)
                                                                                            }
                                                                                        case "a":
                                                                                            var v179 uint8 = 0
                                                                                            if (v3 <= v3) {
                                                                                                fmt.Println(v7, v166)
                                                                                                var v180 int8 = (^(v3))
                                                                                                _ = v180
                                                                                            }
                                                                                            _ = v179
                                                                                        case "":
                                                                                            switch fi61 {
                                                                                            case 795:
                                                                                                var v181 uint32 = 980
                                                                                                _ = v181
                                                                                            case -741:
                                                                                                fmt.Println(v21, v21)
                                                                                            case -1:
                                                                                                if false {
                                                                                                    var v182 int8 = v3
                                                                                                    switch (("" + ("\n\t" + ("foo bar" + (("hello" + "") + ((("foo bar" + "") + "") + "a"))))) + ("\n\t" + "")) {
                                                                                                    case "\n\t":
                                                                                                        if true {
                                                                                                            var v183 int = (fi61 * (fi61 + fi62))
                                                                                                            v166 += -2.25
                                                                                                            switch v183 {
                                                                                                            case 673:
                                                                                                                switch "a" {
                                                                                                                case "hello":
                                                                                                                    fmt.Println(v7, v167)
                                                                                                                default:
                                                                                                                    v184 := GenS3{F0: v183, F1: 203}
                                                                                                                    _ = v184
                                                                                                                }
                                                                                                            default:
                                                                                                                fmt.Println(v182, v5)
                                                                                                            }
                                                                                                            _ = v183
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        v5 += 2
                                                                                                    }
                                                                                                    if (1 >= -394) {
                                                                                                        delete(v2, "foo bar")
                                                                                                        v185 := GenS19{F0: 850, F1: v5, F2: v5}
                                                                                                        _ = v185
                                                                                                    } else {
                                                                                                        v5 += v5
                                                                                                    }
                                                                                                    _ = v182
                                                                                                }
                                                                                            default:
                                                                                                switch "foo bar" {
                                                                                                case "a":
                                                                                                    v63 = v166
                                                                                                    var v186 int32 = (v167 ^ 2147483647)
                                                                                                    _ = v186
                                                                                                case "foo bar":
                                                                                                    v3 = -110
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            v167 = -2147483648
                                                                                        }
                                                                                        var v187 int = fi62
                                                                                        _ = v187
                                                                                    }
                                                                                    v6 = append(v6, (fi61 & 32767))
                                                                                }
                                                                                v21 = float64((float32(v166) + v7))
                                                                                switch ("" + "foo bar") {
                                                                                case "a":
                                                                                    v6 = append(v6, 0)
                                                                                case "foo bar":
                                                                                    fmt.Println(v21, v168, v167)
                                                                                }
                                                                            }
                                                                        }
                                                                    case 128:
                                                                        v5 -= (v5 * (v5 % ((v5) | 1)))
                                                                    case -2147483648:
                                                                        switch 782 {
                                                                        case -421:
                                                                            v167 = v167
                                                                        case 32767:
                                                                            v3 = v3
                                                                        }
                                                                    }
                                                                }
                                                                v188 := GenS20{F0: ("" + (("a" + "foo bar") + "")), F1: (v167 * -1), F2: 42}
                                                                _ = v188
                                                            } else {
                                                                var v189 string = ""
                                                                _ = v189
                                                            }
                                                        }
                                                    case 0:
                                                        fmt.Println(v166, fi61, v63)
                                                    case 9223372036854775807:
                                                        v190 := GenS21{F0: 65535, F1: ("" + ((("foo bar" + "foo bar") + "hello") + ("\n\t" + "foo bar"))), F2: "hello"}
                                                        _ = v190
                                                    }
                                                    v21 += v63
                                                    _ = v168
                                                    _ = v169
                                                } else {
                                                    v191 := GenS11{F0: "foo bar"}
                                                    if len(v6) > 0 {
                                                        v6[0] = (^(fi62))
                                                    }
                                                    var v192 int = fi61
                                                    _ = v191
                                                    _ = v192
                                                }
                                                _ = v166
                                                _ = v167
                                            }
                                        }
                                        switch 128 {
                                        case 32768:
                                            var v193 uint64 = 4294967295
                                            _ = v193
                                        case 32767:
                                            var v194 uint64 = 2
                                            _ = v194
                                        case -9223372036854775808:
                                            var v195 int64 = 711
                                            _ = v195
                                        default:
                                            switch ("\n\t" + (("foo bar" + ("a" + "hello")) + "\n\t")) {
                                            case "hello":
                                                if false {
                                                    var v196 int8 = (v3 + (v3 | -93))
                                                    var v197 string = ((("hello" + "foo bar") + ("\n\t" + "foo bar")) + "\n\t")
                                                    _ = v196
                                                    _ = v197
                                                } else {
                                                    _ = len(v1)
                                                    v3 = (v3 + (v3 & (v3 | (v3 % (((v3 | 0)) | 1)))))
                                                }
                                            case "\n\t":
                                                if (!(true || (!false))) {
                                                    switch (fi61 / ((fi61) | 1)) {
                                                    case -595:
                                                        var v198 uint = (v5 + v5)
                                                        _ = v198
                                                    case 127:
                                                        v199 := GenS19{F0: 2147483647, F1: (v5 & v5), F2: (v5 + 869)}
                                                        _ = v199
                                                    }
                                                    v63 += (v21 / float64((0.0 - -2.25)))
                                                }
                                            case "foo bar":
                                                fmt.Println(v3, fi61)
                                            default:
                                                delete(v2, "a")
                                                var v200 uint8 = 255
                                                _ = v200
                                            }
                                        }
                                    }
                                    _ = v71
                                }
                            }
                        case "":
                            fmt.Println(v21, v21, fi61)
                        case "hello":
                            v7 = 3.14159
                        }
                        var v201 uint = (^(v5))
                        _ = v201
                    }
                    _ = v63
                } else {
                    if (2147483647 == 2147483647) {
                        fmt.Println(v21, v5, v5)
                        if (v3 == v3) {
                            switch "a" {
                            case "\n\t":
                                fmt.Println(fi61, v7)
                            case "foo bar":
                                if (false && (true && false)) {
                                    v202 := GenS7{F0: (-(v3))}
                                    if (-373 > 42) {
                                        v5 -= v5
                                    } else {
                                        switch (fi62 % (((fi61 / ((fi61) | 1))) | 1)) {
                                        case -2147483648:
                                            var v203 bool = (v7 != -2.25)
                                            _ = v203
                                        }
                                        var v204 int8 = v3
                                        _ = v204
                                    }
                                    var v205 int32 = int32(v5)
                                    _ = v202
                                    _ = v205
                                } else {
                                    v7 = float32(v21)
                                }
                            }
                        } else {
                            var v206 int32 = 2147483647
                            var v207 uint = 325
                            if (false || false) {
                                switch ("hello" + "a") {
                                case "a":
                                    var v208 int64 = -818
                                    _ = v208
                                case "":
                                    v21 += float64(3.14159)
                                }
                            }
                            _ = v206
                            _ = v207
                        }
                    }
                    switch 32767 {
                    case 1:
                        _ = len(v6)
                    case 42:
                        var v209 string = "\n\t"
                        _ = v209
                    case -875:
                        switch 127 {
                        case 35:
                            var v210 uint8 = 25
                            _ = v210
                        }
                    default:
                        switch fi61 {
                        case 32768:
                            v211 := GenS6{F0: 2}
                            v21 = (v21 - (((float64(0.1) * v21) - ((v21 / (-2.25 / v21)) * (float64(v7) - (1.5 / v21)))) - (((1.5 - (3.14159 - -2.25)) - (v21 / (0.0 + v21))) * v21)))
                            v212 := GenS14{F0: fi62, F1: 18446744073709551615}
                            if ((((fi62 / (((fi62 << ((154) & 63))) | 1)) == fi61) && true) && true) {
                                var v213 uint = v5
                                v214 := GenS2{F0: 0}
                                v213 = v5
                                _ = v213
                                _ = v214
                            } else {
                                v215 := GenS22{F0: 255, F1: fi62, F2: (v5 <= 427)}
                                _ = v215
                            }
                            if true {
                                if (42 != 65535) {
                                    v216 := GenS14{F0: (fi62 << (((v5 ^ v5)) & 63)), F1: 18446744073709551615}
                                    fmt.Println(fi61, v7)
                                    fmt.Println(v7, fi62, v5)
                                    _ = v216
                                }
                                v217 := GenS23{F0: 2147483647, F1: -1}
                                v5 = (v5 ^ v5)
                                _ = v217
                            } else {
                                v218 := GenS21{F0: 255, F1: "hello", F2: "hello"}
                                _ = v218
                            }
                            _ = v211
                            _ = v212
                        case 127:
                            if (true || true) {
                                v7 += v7
                            }
                        }
                    }
                }
                v6 = append(v6, (fi62 / (((fi61 >> ((v5) & 63))) | 1)))
                _ = fi62
            }
            _ = fi61
        }
    }
    switch "foo bar" {
    case "hello":
        v5 -= 42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
