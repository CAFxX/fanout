package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 bool; F2 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 bool; F2 int }
type GenS8 struct { F0 int32; F1 uint; F2 bool }
type GenS9 struct { F0 uint64; F1 int8 }
type GenS10 struct { F0 bool; F1 int }
type GenS11 struct { F0 string; F1 uint8; F2 int16 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 bool; F1 uint8 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint; F1 int64; F2 uint8 }
type GenS16 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS17 struct { F0 int64; F1 bool; F2 uint16 }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 int64; F1 int; F2 bool }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    var v3 int64 = 840
    v4 := []string{"\n\t", "", ""}
    var v5 int = 42
    var v6 string = "a"
    var v7 uint = 0
    for fi8 := 0; fi8 < 13; fi8++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v3 = v3
        _ = fi8
    }
    v7 -= 18446744073709551615
    if true {
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS1{F0: 4294967295}
            var v11 string = (v6 + "a")
            _ = fi9
            _ = v10
            _ = v11
        }
        switch "" {
        case "foo bar":
            switch -1 {
            case 0:
                switch v6 {
                case "foo bar":
                    var v12 bool = true
                    _ = v12
                case "\n\t":
                    if true {
                        for fi13 := 0; fi13 < 11; fi13++ {
                            v14 := GenS1{F0: 42}
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v15 int = 1
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        v6 = "hello"
                    } else {
                        v7 = v7
                    }
                case "a":
                    if true {
                        for fi16 := 0; fi16 < 10; fi16++ {
                            for fi17 := 0; fi17 < 16; fi17++ {
                                var v18 int16 = 883
                                var v19 float32 = 0.0
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            v7 = (v7 >> ((v7) & 63))
                            v6 = v6
                            _ = fi16
                        }
                    } else {
                        switch v6 {
                        case "hello":
                            v20 := GenS1{F0: 4294967295}
                            _ = v20
                        default:
                            for fi21 := 0; fi21 < 17; fi21++ {
                                var v22 float64 = 1e10
                                _ = len(v4)
                                fmt.Println(v22)
                                _ = fi21
                                _ = v22
                            }
                        }
                    }
                default:
                    v23 := GenS1{F0: 18446744073709551615}
                    _ = v23
                }
            case 497:
                fmt.Println(v3, v7, v6)
            default:
                var v24 int = (v5 - (^(v5)))
                _ = v24
            }
        case "":
            v3 = 487
        case "a":
            if false {
                for fi25 := 0; fi25 < 7; fi25++ {
                    switch v5 {
                    case 127:
                        v26 := GenS2{F0: 42, F1: int32(fi25)}
                        _ = v26
                    case -232:
                        fmt.Println(fi25, v6, v7)
                    case 128:
                        fmt.Println(v3, v5, v6)
                    default:
                        v27 := GenS3{F0: -2147483648}
                        _ = v27
                    }
                    v3 = (-(v3))
                    _ = fi25
                }
            }
        default:
            fmt.Println(v5)
        }
        fmt.Println(v6, v3)
    } else {
        if (234 > 97) {
            var v28 float64 = ((-2.25 + 0.1) - 0.0)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v28
        } else {
            if (128 < 128) {
                v29 := GenS4{F0: (v3 + (^(v3))), F1: (true || false), F2: -32768}
                _ = v29
            } else {
                if ((!(-861 == 42)) && (-1 >= v3)) {
                    if false {
                        if (!true) {
                            for fi30 := 0; fi30 < 6; fi30++ {
                                for fi31 := 0; fi31 < 19; fi31++ {
                                    v32 := v2[("hello" + v6)]
                                    var v33 int = 127
                                    switch v6 {
                                    case "a":
                                        v34 := GenS3{F0: -2147483648}
                                        _ = v34
                                    case "hello":
                                        v4 = append(v4, "foo bar")
                                    case "foo bar":
                                        var v35 uint16 = 1
                                        _ = v35
                                    }
                                    _ = fi31
                                    _ = v32
                                    _ = v33
                                }
                                var v36 uint64 = 259
                                switch (fi30 ^ v5) {
                                case 32768:
                                    v37 := GenS2{F0: 1, F1: int32(v36)}
                                    _ = v37
                                case 1:
                                    switch int(v3) {
                                    case 127:
                                        var v38 uint32 = uint32(v36)
                                        _ = v38
                                    case 32767:
                                        v7 = (v7 - (v7 | v7))
                                    case -9223372036854775808:
                                        var v39 uint32 = 0
                                        _ = v39
                                    default:
                                        var v40 uint = (v7 & 18446744073709551615)
                                        _ = v40
                                    }
                                }
                                _ = fi30
                                _ = v36
                            }
                            for fi41 := 0; fi41 < 14; fi41++ {
                                var v42 int64 = v3
                                var v43 string = ((v6 + "\n\t") + v6)
                                for fi44 := 0; fi44 < 12; fi44++ {
                                    if (255 <= v7) {
                                        var v45 bool = (178 <= 152)
                                        switch v43 {
                                        case "foo bar":
                                            v45 = (("a" + (v43 + v43)) <= (v43 + v43))
                                        case "\n\t":
                                            if (1000 == 770) {
                                                v1[fi44] = 0
                                                v46 := v1[(fi44 | fi44)]
                                                _ = v46
                                            }
                                            var v47 bool = (v45 && (v45 && v45))
                                            switch v5 {
                                            case 127:
                                                switch (fi41 | fi44) {
                                                case 609:
                                                    v48 := GenS3{F0: 2147483647}
                                                    _ = v48
                                                case -9223372036854775808:
                                                    var v49 float64 = (-2.25 / -2.25)
                                                    _ = v49
                                                default:
                                                    _ = len(v2)
                                                }
                                                var v50 string = v6
                                                if v45 {
                                                    switch v6 {
                                                    case "":
                                                        _ = len(v4)
                                                    default:
                                                        var v51 int8 = 105
                                                        v51 -= (v51 | (v51 ^ (v51 * v51)))
                                                        _ = v51
                                                    }
                                                    fmt.Println(v43, fi44, v6)
                                                } else {
                                                    _ = len(v4)
                                                }
                                                var v52 int16 = 612
                                                v50 = ((v6 + v6) + ("a" + (v50 + v43)))
                                                v53 := GenS5{F0: (!(3.14159 == 1.5))}
                                                _ = v50
                                                _ = v52
                                                _ = v53
                                            case 254:
                                                fmt.Println(v45, fi44)
                                            default:
                                                v54 := GenS1{F0: 240}
                                                _ = v54
                                            }
                                            if v45 {
                                                v55 := v2[(("a" + (v43 + v43)) + v6)]
                                                _ = v55
                                            } else {
                                                var v56 bool = v45
                                                if v47 {
                                                    if v45 {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        v57 := GenS4{F0: v42, F1: (((true || true) && (0 >= -902)) && false), F2: 0}
                                                        if v45 {
                                                            var v58 uint = 255
                                                            _ = v58
                                                        }
                                                        _ = v57
                                                    }
                                                    if true {
                                                        v7 -= v7
                                                        v3 = v42
                                                    } else {
                                                        v3 -= (v3 ^ v42)
                                                    }
                                                    fmt.Println(fi41, v6, fi44)
                                                } else {
                                                    v56 = ((-32768 <= 0) && v47)
                                                    v42 -= int64(v7)
                                                    var v59 uint64 = 0
                                                    _ = v59
                                                }
                                                if (!(true || ((!(2147483648 <= uint32(v7))) || (((42 >= 405) && v47) || false)))) {
                                                    var v60 uint64 = 1
                                                    v4 = append(v4, "hello")
                                                    v43 = "\n\t"
                                                    _ = v60
                                                } else {
                                                    fmt.Println(v7, v47, v47)
                                                    var v61 int32 = -2147483648
                                                    _ = v61
                                                }
                                                _ = v56
                                            }
                                            _ = v47
                                        case "a":
                                            var v62 string = v6
                                            _ = v62
                                        }
                                        switch (v43 + (v43 + ((v6 + "foo bar") + (((v43 + v43) + "a") + v43)))) {
                                        case "hello":
                                            fmt.Println(fi41)
                                        default:
                                            var v63 string = (v6 + (v6 + ("foo bar" + "foo bar")))
                                            _ = v63
                                        }
                                        _ = v45
                                    } else {
                                        var v64 uint8 = 42
                                        _ = v64
                                    }
                                    _ = len(v4)
                                    _ = fi44
                                }
                                _ = fi41
                                _ = v42
                                _ = v43
                            }
                            for fi65 := 0; fi65 < 12; fi65++ {
                                v66 := GenS2{F0: 1, F1: 292}
                                for fi67 := 0; fi67 < 1; fi67++ {
                                    var v68 int16 = 0
                                    _ = fi67
                                    _ = v68
                                }
                                _ = fi65
                                _ = v66
                            }
                        }
                        var v69 int8 = int8(v3)
                        _ = v69
                    } else {
                        v70 := GenS6{F0: 4294967295}
                        v71 := GenS7{F0: 2, F1: true, F2: 9223372036854775807}
                        fmt.Println(v7, v3)
                        _ = v70
                        _ = v71
                    }
                }
                if false {
                    for fi72 := 0; fi72 < 3; fi72++ {
                        delete(v2, v6)
                        for fi73 := 0; fi73 < 2; fi73++ {
                            var v74 uint32 = 2
                            v75 := GenS8{F0: 1, F1: v7, F2: ((4294967295 < 18446744073709551615) && (!true))}
                            delete(v2, v6)
                            _ = fi73
                            _ = v74
                            _ = v75
                        }
                        var v76 float32 = 3.14159
                        _ = fi72
                        _ = v76
                    }
                    for fi77 := 0; fi77 < 17; fi77++ {
                        for fi78 := 0; fi78 < 13; fi78++ {
                            v79 := GenS9{F0: 18446744073709551615, F1: 114}
                            _ = fi78
                            _ = v79
                        }
                        if (255 < v7) {
                            if false {
                                _ = len(v4)
                                v6 = (v6 + ("hello" + v6))
                                v80 := GenS3{F0: 0}
                                _ = v80
                            }
                            v6 = v6
                            v5 -= fi77
                        } else {
                            for fi81 := 0; fi81 < 10; fi81++ {
                                v7 -= 912
                                _ = fi81
                            }
                            switch -689 {
                            case 1:
                                var v82 uint8 = 1
                                _ = v82
                            case 42:
                                v83 := GenS2{F0: 65535, F1: 1}
                                _ = v83
                            default:
                                fmt.Println(v7, v6, v7)
                            }
                        }
                        if (float64(float32(float64(3.14159))) != (float64((float32(1.5) * (-2.25 / float32(1.5)))) / -2.25)) {
                            switch v6 {
                            case "\n\t":
                                var v84 float64 = ((3.14159 / ((float64(0.1) / ((float64((0.1 + (3.14159 / 1.5))) / 1.5) / 1e10)) - 1e10)) / float64((1e10 * 1e10)))
                                _ = v84
                            case "hello":
                                switch v6 {
                                case "\n\t":
                                    if false {
                                        for fi85 := 0; fi85 < 15; fi85++ {
                                            if true {
                                                fmt.Println(fi85)
                                                v7 += v7
                                            } else {
                                                v86 := GenS3{F0: -85}
                                                var v87 uint64 = 381
                                                _ = v86
                                                _ = v87
                                            }
                                            switch v6 {
                                            case "\n\t":
                                                v4 = append(v4, ("\n\t" + (v6 + v6)))
                                            case "":
                                                v6 = v6
                                            case "hello":
                                                var v88 int = (fi85 * 0)
                                                _ = v88
                                            default:
                                                var v89 uint = v7
                                                _ = v89
                                            }
                                            _ = fi85
                                        }
                                        fmt.Println(v6, v5)
                                        for fi90 := 0; fi90 < 12; fi90++ {
                                            if true {
                                                if true {
                                                    var v91 int64 = v3
                                                    if true {
                                                        var v92 bool = true
                                                        var v93 uint = v7
                                                        v94 := v1[-362]
                                                        _ = v92
                                                        _ = v93
                                                        _ = v94
                                                    }
                                                    switch fi90 {
                                                    case -690:
                                                        v95 := GenS5{F0: true}
                                                        _ = v95
                                                    default:
                                                        v7 = 313
                                                    }
                                                    var v96 uint16 = 536
                                                    _ = v91
                                                    _ = v96
                                                } else {
                                                    if ((0.1 < (1e10 * 0.1)) && false) {
                                                        var v97 uint16 = 388
                                                        var v98 uint32 = 4294967295
                                                        _ = v97
                                                        _ = v98
                                                    } else {
                                                        switch v5 {
                                                        case 42:
                                                            _ = len(v4)
                                                        case -944:
                                                            if false {
                                                                v5 = fi77
                                                                if (true && true) {
                                                                    if (0.0 == 3.14159) {
                                                                        switch fi90 {
                                                                        case -870:
                                                                            v6 = v6
                                                                        case 255:
                                                                            if true {
                                                                                if ((0.1 - (3.14159 * 0.1)) == 1e10) {
                                                                                    var v99 uint16 = 2
                                                                                    if false {
                                                                                        var v100 int = fi77
                                                                                        var v101 int16 = 32767
                                                                                        delete(v1, v100)
                                                                                        _ = v100
                                                                                        _ = v101
                                                                                    }
                                                                                    _ = v99
                                                                                } else {
                                                                                    var v102 float32 = float32(float64((1e10 / -2.25)))
                                                                                    v1[(fi77 ^ (-(v5)))] = 2147483648
                                                                                    _ = v102
                                                                                }
                                                                            }
                                                                        case 7:
                                                                            v7 = 938
                                                                        }
                                                                        if (720 != 42) {
                                                                            fmt.Println(v6, v7)
                                                                            var v103 uint64 = uint64(v7)
                                                                            if false {
                                                                                var v104 uint32 = 4294967295
                                                                                fmt.Println(fi77, v5, v104)
                                                                                var v105 uint16 = 2
                                                                                _ = v104
                                                                                _ = v105
                                                                            }
                                                                            switch fi90 {
                                                                            case -9223372036854775808:
                                                                                delete(v1, (v5 ^ int(v7)))
                                                                            case -1:
                                                                                switch fi77 {
                                                                                case 32768:
                                                                                    v6 = "hello"
                                                                                default:
                                                                                    if true {
                                                                                        if (((1.5 <= -2.25) && true) || true) {
                                                                                            switch v5 {
                                                                                            case 0:
                                                                                                v103 = 18446744073709551615
                                                                                            case -2147483648:
                                                                                                var v106 int64 = v3
                                                                                                _ = v106
                                                                                            case 1:
                                                                                                if true {
                                                                                                    v107 := GenS10{F0: ((!(fi90 == 800)) && false), F1: 128}
                                                                                                    v6 = "foo bar"
                                                                                                    v5 = fi90
                                                                                                    _ = v107
                                                                                                } else {
                                                                                                    var v108 float32 = 0.1
                                                                                                    v2[v6] = 0
                                                                                                    _ = v108
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v103 -= (v103 * (v103 / (((v103 ^ (v103 & v103))) | 1)))
                                                                                        }
                                                                                        switch "hello" {
                                                                                        case "":
                                                                                            v7 = 449
                                                                                        case "foo bar":
                                                                                            switch fi77 {
                                                                                            case -9223372036854775808:
                                                                                                switch fi77 {
                                                                                                case 128:
                                                                                                    if (v3 > v3) {
                                                                                                        v109 := GenS11{F0: "foo bar", F1: 85, F2: 1}
                                                                                                        fmt.Println(v103, fi77, v5)
                                                                                                        switch fi90 {
                                                                                                        case 4294967295:
                                                                                                            if (1e10 < float32(0.0)) {
                                                                                                                if ((0 == 1000) || false) {
                                                                                                                    v110 := GenS12{F0: v6}
                                                                                                                    _ = v110
                                                                                                                } else {
                                                                                                                    v111 := GenS7{F0: 4294967295, F1: false, F2: (v5 << ((uint(v103)) & 63))}
                                                                                                                    var v112 bool = (int8(v5) <= -47)
                                                                                                                    v3 = v3
                                                                                                                    _ = v111
                                                                                                                    _ = v112
                                                                                                                }
                                                                                                            }
                                                                                                            switch (fi90 + 737) {
                                                                                                            case -1:
                                                                                                                var v113 uint8 = 42
                                                                                                                fmt.Println(v3)
                                                                                                                v114 := GenS8{F0: -266, F1: 369, F2: (2147483648 <= 852)}
                                                                                                                delete(v1, 593)
                                                                                                                if (false || (2 == 255)) {
                                                                                                                    if (((v6 + "a") + ("hello" + ((("a" + "a") + "hello") + "a"))) < ("foo bar" + "\n\t")) {
                                                                                                                        v2[v6] = 2
                                                                                                                    } else {
                                                                                                                        fmt.Println(fi77, v5)
                                                                                                                        if ((0.1 > ((1e10 + 0.1) * 0.0)) || false) {
                                                                                                                            switch 640 {
                                                                                                                            case -318:
                                                                                                                                var v115 float64 = (0.0 - float64(((0.1 / float32(1e10)) - 0.1)))
                                                                                                                                _ = v115
                                                                                                                            default:
                                                                                                                                v3 = (v3 & (v3 + (v3 << ((633) & 63))))
                                                                                                                            }
                                                                                                                            var v116 int16 = 0
                                                                                                                            v103 -= (v103 * (v103 & v103))
                                                                                                                            v5 = 0
                                                                                                                            var v117 string = v6
                                                                                                                            _ = v116
                                                                                                                            _ = v117
                                                                                                                        }
                                                                                                                        v5 -= fi77
                                                                                                                        v4 = append(v4, ("a" + ((v6 + "\n\t") + v6)))
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v113
                                                                                                                _ = v114
                                                                                                            case -9223372036854775808:
                                                                                                                v5 += fi77
                                                                                                            case -296:
                                                                                                                if len(v4) > 0 {
                                                                                                                    v4[0] = (v6 + v6)
                                                                                                                }
                                                                                                            default:
                                                                                                                fmt.Println(v6, v103)
                                                                                                            }
                                                                                                        case 128:
                                                                                                            v5 += fi90
                                                                                                        case 32768:
                                                                                                            v118 := GenS4{F0: (-(v3)), F1: false, F2: -735}
                                                                                                            _ = v118
                                                                                                        default:
                                                                                                            v5 = 2147483647
                                                                                                        }
                                                                                                        _ = v109
                                                                                                    } else {
                                                                                                        if false {
                                                                                                            switch fi90 {
                                                                                                            case 128:
                                                                                                                if (!false) {
                                                                                                                    v3 = 42
                                                                                                                    v119 := GenS2{F0: 737, F1: 42}
                                                                                                                    _ = len(v4)
                                                                                                                    _ = v119
                                                                                                                } else {
                                                                                                                    v5 = (v5 | 1)
                                                                                                                    var v120 float64 = (1e10 / 3.14159)
                                                                                                                    _ = v120
                                                                                                                }
                                                                                                            case -1:
                                                                                                                switch (v5 << (((^(v7))) & 63)) {
                                                                                                                case -273:
                                                                                                                    if false {
                                                                                                                        v3 = 1
                                                                                                                    } else {
                                                                                                                        v7 -= 918
                                                                                                                    }
                                                                                                                case 760:
                                                                                                                    fmt.Println(v6)
                                                                                                                case -2147483648:
                                                                                                                    switch int(v103) {
                                                                                                                    case 42:
                                                                                                                        var v121 uint8 = 40
                                                                                                                        _ = v121
                                                                                                                    case 0:
                                                                                                                        if (false && true) {
                                                                                                                            if len(v4) >= 2 {
                                                                                                                                v4 = v4[1:len(v4)]
                                                                                                                            }
                                                                                                                            var v122 uint32 = 65535
                                                                                                                            switch fi77 {
                                                                                                                            case 1:
                                                                                                                                fmt.Println(v103)
                                                                                                                            }
                                                                                                                            var v123 string = v6
                                                                                                                            v123 = v123
                                                                                                                            _ = v122
                                                                                                                            _ = v123
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v124 uint32 = 4294967295
                                                                                                                        _ = v124
                                                                                                                    }
                                                                                                                default:
                                                                                                                    fmt.Println(fi90, v5)
                                                                                                                }
                                                                                                            case 127:
                                                                                                                var v125 float32 = 1e10
                                                                                                                _ = v125
                                                                                                            default:
                                                                                                                if (!(!(false || false))) {
                                                                                                                    v126 := GenS6{F0: 2147483648}
                                                                                                                    _ = v126
                                                                                                                } else {
                                                                                                                    var v127 int8 = -63
                                                                                                                    v103 += (v103 / ((v103) | 1))
                                                                                                                    _ = v127
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                case 862:
                                                                                                    var v128 int64 = v3
                                                                                                    _ = v128
                                                                                                }
                                                                                            case 9223372036854775807:
                                                                                                var v129 bool = false
                                                                                                _ = v129
                                                                                            case 4294967295:
                                                                                                v2[v6] = 747
                                                                                            }
                                                                                        }
                                                                                        switch (fi90 >> ((v7) & 63)) {
                                                                                        case -9223372036854775808:
                                                                                            if (!(-61 <= 1)) {
                                                                                                v103 += v103
                                                                                                fmt.Println(v103, fi90, fi90)
                                                                                                var v130 uint16 = 2
                                                                                                _ = v130
                                                                                            }
                                                                                        case -2147483648:
                                                                                            var v131 int = -9223372036854775808
                                                                                            _ = v131
                                                                                        case 0:
                                                                                            if false {
                                                                                                switch fi90 {
                                                                                                case -625:
                                                                                                    v132 := v1[(fi90 / (((fi90 ^ (^(v5)))) | 1))]
                                                                                                    _ = v132
                                                                                                }
                                                                                            } else {
                                                                                                v3 = 0
                                                                                                if true {
                                                                                                    v133 := GenS13{F0: true, F1: 127}
                                                                                                    var v134 bool = (!(2147483647 == int(v7)))
                                                                                                    _ = v133
                                                                                                    _ = v134
                                                                                                } else {
                                                                                                    if len(v4) >= 2 {
                                                                                                        v4 = v4[1:len(v4)]
                                                                                                    }
                                                                                                    switch (fi90 * 0) {
                                                                                                    case -1:
                                                                                                        v6 = (v6 + (v6 + ("\n\t" + ((("foo bar" + "a") + v6) + v6))))
                                                                                                    case -2147483648:
                                                                                                        var v135 uint64 = v103
                                                                                                        _ = v135
                                                                                                    default:
                                                                                                        if true {
                                                                                                            if true {
                                                                                                                var v136 uint64 = 510
                                                                                                                _ = v136
                                                                                                            }
                                                                                                            v6 = ("a" + "foo bar")
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(v7)
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case 428:
                                                                                var v137 uint64 = v103
                                                                                _ = v137
                                                                            }
                                                                            v138 := GenS12{F0: v6}
                                                                            var v139 int16 = 42
                                                                            _ = v103
                                                                            _ = v138
                                                                            _ = v139
                                                                        } else {
                                                                            v7 += v7
                                                                            var v140 uint64 = 68
                                                                            switch v5 {
                                                                            case 894:
                                                                                v2["a"] = 42
                                                                            default:
                                                                                if (false && (!(true && ((((((1.5 - -2.25) <= -2.25) && (!(true || true))) && (-2147483648 > -2147483648)) && ((^(v140)) <= 1)) || false)))) {
                                                                                    v141 := GenS14{F0: (v3 << ((4294967295) & 63))}
                                                                                    var v142 uint = v7
                                                                                    var v143 uint16 = 32767
                                                                                    _ = v141
                                                                                    _ = v142
                                                                                    _ = v143
                                                                                } else {
                                                                                    v7 -= (v7 * 18446744073709551615)
                                                                                }
                                                                            }
                                                                            _ = v140
                                                                        }
                                                                        if ((fi77 * (^(fi90))) < v5) {
                                                                            v6 = v6
                                                                            var v144 string = "hello"
                                                                            switch (v144 + (((((v6 + "foo bar") + v6) + ("foo bar" + v144)) + v144) + v6)) {
                                                                            case "":
                                                                                var v145 uint32 = 42
                                                                                _ = v145
                                                                            case "hello":
                                                                                v7 = 4294967295
                                                                            default:
                                                                                if true {
                                                                                    v5 += v5
                                                                                } else {
                                                                                    v146 := v2[(v6 + v144)]
                                                                                    v147 := GenS13{F0: false, F1: 248}
                                                                                    v3 = v3
                                                                                    _ = v146
                                                                                    _ = v147
                                                                                }
                                                                            }
                                                                            _ = v144
                                                                        }
                                                                        var v148 uint8 = 1
                                                                        _ = len(v1)
                                                                        _ = v148
                                                                    }
                                                                } else {
                                                                    v149 := GenS5{F0: (0 > -81)}
                                                                    fmt.Println(v6, v7)
                                                                    var v150 float32 = 0.1
                                                                    var v151 float32 = 1e10
                                                                    _ = v149
                                                                    _ = v150
                                                                    _ = v151
                                                                }
                                                            }
                                                        case -1:
                                                            fmt.Println(v3)
                                                        }
                                                        switch fi77 {
                                                        case 9223372036854775807:
                                                            var v152 int8 = 42
                                                            _ = v152
                                                        }
                                                        var v153 int32 = 747
                                                        var v154 uint64 = 18446744073709551615
                                                        if false {
                                                            var v155 uint64 = 4294967295
                                                            v156 := GenS3{F0: (v153 * v153)}
                                                            v153 = (v153 - int32(v3))
                                                            _ = v155
                                                            _ = v156
                                                        } else {
                                                            v4 = append(v4, v6)
                                                            v157 := GenS15{F0: 458, F1: (v3 | 0), F2: 90}
                                                            if (!(false || (!true))) {
                                                                var v158 float64 = (((-2.25 - float64(((3.14159 / 1e10) - 0.0))) + ((((1e10 + float64(0.0)) - 0.1) / float64(-2.25)) - 0.0)) / 0.1)
                                                                var v159 float32 = 0.0
                                                                if false {
                                                                    fmt.Println(v6, fi77)
                                                                }
                                                                _ = v158
                                                                _ = v159
                                                            } else {
                                                                v153 += v153
                                                            }
                                                            _ = v157
                                                        }
                                                        fmt.Println(v154, v6)
                                                        _ = v153
                                                        _ = v154
                                                    }
                                                }
                                            } else {
                                                if true {
                                                    fmt.Println(v3)
                                                    if false {
                                                        switch ("a" + "hello") {
                                                        case "":
                                                            v160 := GenS7{F0: 2147483648, F1: false, F2: -2147483648}
                                                            _ = v160
                                                        default:
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        var v161 float32 = float32(((1e10 - -2.25) * 3.14159))
                                                        v6 = "foo bar"
                                                        _ = v161
                                                    }
                                                }
                                                var v162 float32 = (0.0 / 0.1)
                                                var v163 string = (v6 + ((v6 + ("foo bar" + ("" + ((v6 + v6) + (v6 + v6))))) + v6))
                                                _ = v162
                                                _ = v163
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = ("" + v6)
                                            }
                                            if false {
                                                fmt.Println(v6)
                                                v3 -= int64(fi77)
                                                v6 = "foo bar"
                                                if (!false) {
                                                    var v164 float64 = (1e10 / 1.5)
                                                    v165 := GenS10{F0: (18446744073709551615 == 994), F1: fi77}
                                                    if false {
                                                        v7 = (v7 * v7)
                                                        v5 += (v5 << ((v7) & 63))
                                                        var v166 float32 = float32(v164)
                                                        _ = v166
                                                    } else {
                                                        _ = len(v4)
                                                        v164 = v164
                                                    }
                                                    _ = v164
                                                    _ = v165
                                                } else {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    if false {
                                                        v167 := GenS16{F0: 2, F1: 2, F2: 70}
                                                        v7 = (v7 >> ((477) & 63))
                                                        _ = v167
                                                    } else {
                                                        _ = len(v2)
                                                        var v168 bool = (v3 == v3)
                                                        _ = v168
                                                    }
                                                }
                                            }
                                            _ = fi90
                                        }
                                    } else {
                                        v7 -= 42
                                        var v169 int32 = 780
                                        for fi170 := 0; fi170 < 3; fi170++ {
                                            v169 += (^(v169))
                                            v169 -= (^(v169))
                                            fmt.Println(v169, fi170)
                                            _ = fi170
                                        }
                                        _ = v169
                                    }
                                }
                            }
                        }
                        _ = fi77
                    }
                } else {
                    for fi171 := 0; fi171 < 17; fi171++ {
                        v172 := GenS17{F0: (v3 ^ -9223372036854775808), F1: ((((v3 >> ((v7) & 63)) == v3) || (1 <= 197)) && true), F2: 255}
                        var v173 uint16 = 8
                        _ = fi171
                        _ = v172
                        _ = v173
                    }
                    fmt.Println(v5, v3, v3)
                    fmt.Println(v7)
                }
                for fi174 := 0; fi174 < 11; fi174++ {
                    for fi175 := 0; fi175 < 2; fi175++ {
                        var v176 uint32 = 65535
                        v7 = 500
                        _ = fi175
                        _ = v176
                    }
                    fmt.Println(fi174)
                    _ = fi174
                }
            }
            _ = len(v4)
            var v177 int16 = int16(v5)
            _ = v177
        }
        v6 = "a"
        fmt.Println(v5, v5, v5)
    }
    if false {
        if (42 <= 42) {
            v178 := GenS5{F0: false}
            if (65535 == 1) {
                v179 := GenS18{F0: v7}
                _ = v179
            } else {
                if (839 != 2147483647) {
                    v3 = (v3 % ((0) | 1))
                    switch (v5 + (^(v5))) {
                    case 2147483647:
                        v180 := GenS6{F0: 439}
                        _ = v180
                    case 127:
                        switch v5 {
                        case 2147483647:
                            v181 := GenS19{F0: -786}
                            _ = v181
                        case 32767:
                            var v182 uint = v7
                            _ = v182
                        case 32768:
                            for fi183 := 0; fi183 < 11; fi183++ {
                                var v184 bool = (!(false || (false || false)))
                                _ = fi183
                                _ = v184
                            }
                        default:
                            var v185 uint32 = 2147483648
                            _ = v185
                        }
                    default:
                        for fi186 := 0; fi186 < 7; fi186++ {
                            var v187 int = v5
                            _ = fi186
                            _ = v187
                        }
                    }
                }
            }
            _ = v178
        }
    }
    if len(v4) > 0 {
        v4[0] = v6
    }
    v3 = 9223372036854775807
    var v188 float64 = 3.14159
    for fi189 := 0; fi189 < 11; fi189++ {
        v190 := GenS20{F0: v3, F1: -2147483648, F2: (v6 > v6)}
        _ = fi189
        _ = v190
    }
    v191 := GenS1{F0: 738}
    fmt.Println(v5, v5, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v188:", v188)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v188
    _ = v191
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
