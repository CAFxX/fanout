package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS6 struct { F0 bool; F1 uint32; F2 int8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool; F1 uint64; F2 int16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint16; F1 int64; F2 int8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int16; F1 int16; F2 int32 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 bool; F1 int64 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS17 struct { F0 int64; F1 bool; F2 int32 }
type GenS18 struct { F0 int; F1 string; F2 uint8 }
type GenS19 struct { F0 int16; F1 uint32 }
type GenS20 struct { F0 int; F1 uint16; F2 bool }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 int32; F1 int8 }
type GenS23 struct { F0 int32; F1 int32 }
type GenS24 struct { F0 uint16; F1 bool; F2 uint64 }
func genFn0() uint16 {
    if false {
        if ((!false) || false) {
            var v6 int64 = -230
            _ = v6
        } else {
            v7 := []int{-501}
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 string = ("hello" + "\n\t")
        switch (fi8 * fi8) {
        case 127:
            var v10 int64 = -1
            _ = v10
        }
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 int32 = 1
            if false {
                v9 = v9
                if true {
                    v13 := []int{-336, int(v12)}
                    if true {
                        switch v9 {
                        case "a":
                            if true {
                                v14 := GenS1{F0: 2147483648, F1: 65535, F2: 2}
                                var v15 bool = true
                                _ = v14
                                _ = v15
                            } else {
                                switch fi11 {
                                case 2147483647:
                                    fmt.Println(fi11)
                                }
                                v9 = v9
                            }
                        case "hello":
                            v9 = ""
                        }
                    }
                    _ = v13
                }
                v12 = v12
            }
            _ = fi11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    var v16 int8 = 16
    for fi17 := 0; fi17 < 5; fi17++ {
        fmt.Println(v16, fi17)
        switch ("\n\t" + "") {
        case "foo bar":
            var v18 string = "hello"
            _ = v18
        case "":
            for fi19 := 0; fi19 < 2; fi19++ {
                var v20 uint16 = 42
                _ = fi19
                _ = v20
            }
        }
        switch "foo bar" {
        case "":
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 string = ("" + ("a" + "hello"))
                _ = fi21
                _ = v22
            }
        case "\n\t":
            if ((-(fi17)) > (fi17 / ((2147483647) | 1))) {
                v16 = 127
                var v23 int64 = 1
                v24 := GenS1{F0: 316, F1: 2147483647, F2: 879}
                _ = v23
                _ = v24
            }
        }
        _ = fi17
    }
    _ = v16
    return 65535
}
func genFn1(p0 uint32, p1 int64) uint {
    v25 := GenS1{F0: (^(p0)), F1: p0, F2: 434}
    v26 := []int{32768, 9223372036854775807, -2147483648}
    for fi27 := 0; fi27 < 17; fi27++ {
        p1 = p1
        _ = fi27
    }
    _ = p0
    _ = p1
    _ = v25
    _ = v26
    return 255
}

func main() {
    v1 := []uint32{1, 2}
    v2 := make(map[uint32]string)
    var v3 float64 = 1.5
    v4 := make(map[string]int)
    v5 := []string{"hello", "foo bar", "\n\t"}
    switch "a" {
    case "hello":
        v28 := GenS2{F0: false, F1: 1}
        _ = v28
    case "":
        v3 = v3
    case "foo bar":
        if true {
            v29 := GenS3{F0: ("" + "foo bar"), F1: 2}
            if (-795 >= -1) {
                fmt.Println(v3)
            }
            _ = v29
        } else {
            for fi30 := 0; fi30 < 6; fi30++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi30
            }
            for fi31 := 0; fi31 < 10; fi31++ {
                v3 -= 1e10
                var v32 float64 = float64(1e10)
                _ = fi31
                _ = v32
            }
            v3 += -2.25
        }
    }
    var v33 uint32 = 65535
    v3 = v3
    if true {
        v3 = v3
    } else {
        v34 := GenS2{F0: (255 > 313), F1: 0}
        v35 := GenS2{F0: false, F1: 113}
        if (127 != 128) {
            if (127 != 109) {
                for fi36 := 0; fi36 < 9; fi36++ {
                    for fi37 := 0; fi37 < 15; fi37++ {
                        v38 := GenS2{F0: (!false), F1: 0}
                        var v39 int16 = -141
                        v40 := GenS3{F0: ("foo bar" + ("\n\t" + "hello")), F1: 255}
                        _ = fi37
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                    _ = fi36
                }
                v3 += 3.14159
                if false {
                    if true {
                        fmt.Println(v3)
                        var v41 float32 = -2.25
                        v4[("foo bar" + (((("a" + "foo bar") + ("a" + "\n\t")) + "hello") + "\n\t"))] = 4294967295
                        _ = v41
                    } else {
                        switch "" {
                        case "":
                            var v42 uint64 = 4294967295
                            _ = v42
                        case "hello":
                            var v43 float64 = 0.0
                            _ = v43
                        case "a":
                            switch "hello" {
                            case "\n\t":
                                fmt.Println(v3, v33)
                            case "":
                                v44 := GenS3{F0: ("a" + ("hello" + ("" + "hello"))), F1: 133}
                                _ = v44
                            case "hello":
                                v3 = v3
                            }
                        }
                        var v45 string = "\n\t"
                        _ = v45
                    }
                }
            } else {
                var v46 uint64 = 699
                _ = v46
            }
            v47 := v4[("hello" + "hello")]
            _ = v47
        } else {
            fmt.Println(v33, v3)
            if (1 > 42) {
                var v48 int = -616
                for fi49 := 0; fi49 < 10; fi49++ {
                    if ((v3 / float64(0.1)) == (float64(1.5) / 3.14159)) {
                        if false {
                            v50 := GenS3{F0: "", F1: 1}
                            _ = v50
                        } else {
                            for fi51 := 0; fi51 < 17; fi51++ {
                                switch fi51 {
                                case 42:
                                    v52 := GenS2{F0: true, F1: 128}
                                    _ = v52
                                }
                                _ = fi51
                            }
                            var v53 bool = (2 >= 2)
                            if false {
                                var v54 int8 = 0
                                if len(v5) > 0 {
                                    v5[0] = (("\n\t" + "foo bar") + ("" + "\n\t"))
                                }
                                _ = v54
                            }
                            _ = v53
                        }
                    }
                    for fi55 := 0; fi55 < 9; fi55++ {
                        if true {
                            v3 -= 0.1
                        } else {
                            v33 -= v33
                            v56 := GenS2{F0: ((42 < -603) || false), F1: 42}
                            _ = v56
                        }
                        if ((42 > -1) && false) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v57 int16 = 738
                            fmt.Println(v48, fi49)
                            _ = v57
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v58 int = v48
                        _ = fi55
                        _ = v58
                    }
                    for fi59 := 0; fi59 < 16; fi59++ {
                        switch "\n\t" {
                        case "a":
                            v3 = -2.25
                        case "":
                            if false {
                                var v60 uint32 = (v33 * 2)
                                if (4294967295 >= 18446744073709551615) {
                                    v61 := v4[((("foo bar" + (("\n\t" + "a") + ("\n\t" + "foo bar"))) + "") + "a")]
                                    var v62 float32 = 0.1
                                    _ = v61
                                    _ = v62
                                } else {
                                    switch (-(fi59)) {
                                    case 969:
                                        if (42 <= 843) {
                                            fmt.Println(v48, v3)
                                            switch "" {
                                            case "\n\t":
                                                var v63 int16 = 42
                                                if (127 < 61) {
                                                    if true {
                                                        var v64 int64 = 645
                                                        _ = v64
                                                    } else {
                                                        v48 = int(v63)
                                                        var v65 float32 = float32((((0.1 * v3) / (v3 / 1.5)) + 3.14159))
                                                        fmt.Println(v48)
                                                        _ = v65
                                                    }
                                                    switch (fi49 + v48) {
                                                    case -669:
                                                        _ = len(v1)
                                                    case 32768:
                                                        var v66 int = fi59
                                                        _ = v66
                                                    case 1:
                                                        v67 := GenS3{F0: "foo bar", F1: 1}
                                                        _ = v67
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                } else {
                                                    v5 = append(v5, "\n\t")
                                                    if (239 == 255) {
                                                        v68 := GenS1{F0: (v60 << ((68) & 63)), F1: (^(v60)), F2: uint64(v33)}
                                                        if (false || false) {
                                                            switch (v48 >> ((255) & 63)) {
                                                            case -9223372036854775808:
                                                                var v69 int32 = 42
                                                                _ = v69
                                                            case -397:
                                                                switch 736 {
                                                                case 42:
                                                                    v33 = v60
                                                                case -687:
                                                                    var v70 int32 = -154
                                                                    _ = v70
                                                                }
                                                            default:
                                                                fmt.Println(fi49)
                                                                switch "\n\t" {
                                                                case "a":
                                                                    v33 = v60
                                                                case "":
                                                                    var v71 uint64 = 4294967295
                                                                    _ = v71
                                                                case "foo bar":
                                                                    if false {
                                                                        v72 := v4[(("hello" + "hello") + "hello")]
                                                                        _ = v72
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            var v73 float64 = v3
                                                            v74 := GenS4{F0: 128, F1: (^(fi59))}
                                                            if (((true || ((((float32(3.14159) + 0.1) + 3.14159) / ((-2.25 + float32(v73)) - 1.5)) != -2.25)) || true) || false) {
                                                                fmt.Println(fi49)
                                                                switch fi49 {
                                                                case -67:
                                                                    fmt.Println(v33)
                                                                case -1:
                                                                    switch 656 {
                                                                    case 1:
                                                                        v75 := GenS5{F0: v60, F1: v33, F2: -1}
                                                                        _ = v75
                                                                    case 9223372036854775807:
                                                                        v73 -= v73
                                                                    case 32767:
                                                                        var v76 int64 = -9223372036854775808
                                                                        _ = v76
                                                                    }
                                                                case 128:
                                                                    fmt.Println(v63, v48, v3)
                                                                default:
                                                                    switch ("foo bar" + "a") {
                                                                    case "foo bar":
                                                                        if false {
                                                                            fmt.Println(v48)
                                                                            var v77 int = (fi49 * (fi49 | 874))
                                                                            _ = len(v5)
                                                                            _ = v77
                                                                        }
                                                                    case "\n\t":
                                                                        var v78 int8 = -54
                                                                        _ = v78
                                                                    case "a":
                                                                        var v79 uint32 = v33
                                                                        _ = v79
                                                                    default:
                                                                        v63 += (v63 + v63)
                                                                    }
                                                                }
                                                            } else {
                                                                v63 -= (v63 / (((-(v63))) | 1))
                                                                if false {
                                                                    v60 += (^(v60))
                                                                    var v80 int32 = -2147483648
                                                                    _ = v80
                                                                } else {
                                                                    var v81 int64 = -1
                                                                    v2[(v60 + v60)] = "hello"
                                                                    switch fi59 {
                                                                    case 32768:
                                                                        fmt.Println(fi49)
                                                                        fmt.Println(v73, v48)
                                                                        switch int(v63) {
                                                                        case 32767:
                                                                            var v82 bool = (false || false)
                                                                            _ = v82
                                                                        case 128:
                                                                            if true {
                                                                                _ = len(v2)
                                                                                v83 := GenS1{F0: (^(v60)), F1: (v33 >> ((255) & 63)), F2: 42}
                                                                                _ = v83
                                                                            } else {
                                                                                if ((2 >= 406) && (255 == 103)) {
                                                                                    if (false && (!true)) {
                                                                                        v63 += v63
                                                                                        if (42 != 813) {
                                                                                            v73 = (((float64((0.1 + 0.0)) - -2.25) + (float64(3.14159) + v3)) * (((0.1 * 1e10) / v3) + (v73 + v73)))
                                                                                            fmt.Println(v3)
                                                                                        } else {
                                                                                            if true {
                                                                                                var v84 int8 = 127
                                                                                                var v85 uint32 = (v60 * 65535)
                                                                                                switch "\n\t" {
                                                                                                case "":
                                                                                                    _ = len(v2)
                                                                                                    fmt.Println(v73, v48, v60)
                                                                                                default:
                                                                                                    v1 = append(v1, (^(v60)))
                                                                                                }
                                                                                                _ = v84
                                                                                                _ = v85
                                                                                            } else {
                                                                                                if true {
                                                                                                    switch (-(fi59)) {
                                                                                                    case -1:
                                                                                                        if len(v1) >= 2 {
                                                                                                            v1 = v1[1:len(v1)]
                                                                                                        }
                                                                                                    default:
                                                                                                        if true {
                                                                                                            switch "\n\t" {
                                                                                                            case "a":
                                                                                                                var v86 string = ""
                                                                                                                _ = v86
                                                                                                            default:
                                                                                                                v33 = v33
                                                                                                            }
                                                                                                            v87 := GenS6{F0: false, F1: v60, F2: int8(v63)}
                                                                                                            v48 += (fi49 & (fi49 + -749))
                                                                                                            _ = v87
                                                                                                        } else {
                                                                                                            var v88 string = "hello"
                                                                                                            switch (v88 + v88) {
                                                                                                            case "":
                                                                                                                v33 -= v60
                                                                                                            }
                                                                                                            _ = v88
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v63)
                                                                                                } else {
                                                                                                    v89 := GenS7{F0: true}
                                                                                                    if true {
                                                                                                        if len(v5) > 0 {
                                                                                                            v5[0] = ("" + "hello")
                                                                                                        }
                                                                                                        if true {
                                                                                                            v33 += (^(v60))
                                                                                                            v3 = v3
                                                                                                        } else {
                                                                                                            fmt.Println(v33)
                                                                                                            if len(v5) > 0 {
                                                                                                                v5[0] = ""
                                                                                                            }
                                                                                                            fmt.Println(v48)
                                                                                                        }
                                                                                                        switch fi49 {
                                                                                                        case 127:
                                                                                                            v48 += (v48 % ((fi59) | 1))
                                                                                                        case 9223372036854775807:
                                                                                                            var v90 int = (fi59 % ((127) | 1))
                                                                                                            _ = v90
                                                                                                        default:
                                                                                                            var v91 uint = 42
                                                                                                            _ = v91
                                                                                                        }
                                                                                                        var v92 string = "hello"
                                                                                                        _ = v92
                                                                                                    }
                                                                                                    _ = v89
                                                                                                }
                                                                                            }
                                                                                            var v93 uint16 = 2
                                                                                            fmt.Println(v48, v48, v73)
                                                                                            v94 := GenS8{F0: true, F1: 42, F2: 42}
                                                                                            v73 += (v73 / (1e10 - v73))
                                                                                            _ = v93
                                                                                            _ = v94
                                                                                        }
                                                                                    } else {
                                                                                        var v95 string = "hello"
                                                                                        _ = v95
                                                                                    }
                                                                                } else {
                                                                                    switch "hello" {
                                                                                    case "a":
                                                                                        var v96 uint = 255
                                                                                        _ = v96
                                                                                    case "":
                                                                                        v60 = v60
                                                                                    case "foo bar":
                                                                                        var v97 uint8 = 0
                                                                                        _ = v97
                                                                                    }
                                                                                    v1 = append(v1, (^(v33)))
                                                                                }
                                                                                v98 := GenS7{F0: false}
                                                                                var v99 int8 = -1
                                                                                _ = v98
                                                                                _ = v99
                                                                            }
                                                                        default:
                                                                            v73 -= v3
                                                                        }
                                                                    case 128:
                                                                        v100 := GenS6{F0: ((int(v81) != 72) && false), F1: v33, F2: 0}
                                                                        _ = v100
                                                                    }
                                                                    _ = v81
                                                                }
                                                                var v101 int8 = 55
                                                                if true {
                                                                    v102 := GenS9{F0: (^(v33))}
                                                                    _ = v102
                                                                } else {
                                                                    var v103 float64 = v3
                                                                    _ = v103
                                                                }
                                                                _ = v101
                                                            }
                                                            _ = v73
                                                            _ = v74
                                                        }
                                                        _ = v68
                                                    }
                                                    if false {
                                                        v104 := GenS1{F0: v33, F1: (^(v33)), F2: 42}
                                                        delete(v4, (("" + "\n\t") + ("" + "\n\t")))
                                                        v60 = v60
                                                        if ("\n\t" >= ("" + "\n\t")) {
                                                            v3 -= v3
                                                            if false {
                                                                if true {
                                                                    var v105 int16 = (v63 % ((-799) | 1))
                                                                    _ = v105
                                                                }
                                                            } else {
                                                                var v106 uint16 = 342
                                                                v107 := GenS3{F0: "foo bar", F1: 4294967295}
                                                                v108 := GenS3{F0: "hello", F1: 65535}
                                                                _ = v106
                                                                _ = v107
                                                                _ = v108
                                                            }
                                                            fmt.Println(v48, v60)
                                                            var v109 uint32 = 2147483648
                                                            v3 += float64(0.1)
                                                            _ = v109
                                                        } else {
                                                            if true {
                                                                if false {
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        var v110 int16 = (v63 << ((65535) & 63))
                                                                        _ = v110
                                                                    }
                                                                    switch (v48 - fi59) {
                                                                    case 127:
                                                                        if true {
                                                                            fmt.Println(fi49, v33)
                                                                            v111 := GenS10{F0: 255, F1: 0, F2: 42}
                                                                            _ = v111
                                                                        } else {
                                                                            if (-111 > -1) {
                                                                                switch (("hello" + ("hello" + "")) + ("foo bar" + ("\n\t" + "a"))) {
                                                                                case "hello":
                                                                                    v2[(v33 ^ v60)] = ""
                                                                                case "":
                                                                                    v5 = append(v5, ("a" + ""))
                                                                                default:
                                                                                    fmt.Println(fi49, fi59)
                                                                                }
                                                                            }
                                                                            fmt.Println(fi49)
                                                                        }
                                                                    case 2147483647:
                                                                        if false {
                                                                            v3 += v3
                                                                            var v112 int16 = v63
                                                                            var v113 uint = 18446744073709551615
                                                                            _ = v112
                                                                            _ = v113
                                                                        } else {
                                                                            v114 := GenS8{F0: ((v48 ^ fi59) < v48), F1: 2, F2: (v63 + v63)}
                                                                            switch ("\n\t" + "\n\t") {
                                                                            case "foo bar":
                                                                                v115 := GenS11{F0: fi59}
                                                                                _ = v115
                                                                            case "\n\t":
                                                                                fmt.Println(fi49, fi59, v33)
                                                                            case "":
                                                                                v116 := v2[(v33 | (v33 & (v33 + v60)))]
                                                                                var v117 int32 = 2147483647
                                                                                _ = v116
                                                                                _ = v117
                                                                            default:
                                                                                delete(v2, v60)
                                                                            }
                                                                            var v118 uint = 69
                                                                            _ = v114
                                                                            _ = v118
                                                                        }
                                                                    case -1:
                                                                        var v119 uint = 1
                                                                        _ = v119
                                                                    }
                                                                    if (!true) {
                                                                        v63 = (v63 * (v63 & v63))
                                                                    }
                                                                    if (-2.25 > 0.0) {
                                                                        fmt.Println(v48)
                                                                        var v120 float64 = v3
                                                                        var v121 string = "hello"
                                                                        _ = v120
                                                                        _ = v121
                                                                    }
                                                                } else {
                                                                    v48 = (v48 >> ((4294967295) & 63))
                                                                }
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v122 bool = (true || false)
                                                                    _ = v122
                                                                case "\n\t":
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        fmt.Println(fi59, v3, v60)
                                                                    case "":
                                                                        v33 = (v60 % (((^(v60))) | 1))
                                                                    default:
                                                                        v33 += v60
                                                                    }
                                                                }
                                                                if true {
                                                                    switch fi49 {
                                                                    case 0:
                                                                        v48 += (fi59 ^ 4294967295)
                                                                    case -9223372036854775808:
                                                                        switch "a" {
                                                                        case "hello":
                                                                            v33 -= v33
                                                                        case "":
                                                                            v4[("hello" + "foo bar")] = fi49
                                                                        case "foo bar":
                                                                            v3 += v3
                                                                        default:
                                                                            if false {
                                                                                v4[("" + "a")] = v48
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = v33
                                                                                }
                                                                                if (v3 == (v3 * (v3 + ((v3 / v3) - v3)))) {
                                                                                    var v123 int8 = -122
                                                                                    v33 = v33
                                                                                    _ = v123
                                                                                } else {
                                                                                    if false {
                                                                                        var v124 uint16 = 42
                                                                                        _ = v124
                                                                                    } else {
                                                                                        v33 = (v33 - 4294967295)
                                                                                        v3 = float64(3.14159)
                                                                                        switch fi49 {
                                                                                        case -1:
                                                                                            switch (fi49 ^ (-(fi49))) {
                                                                                            case 4294967295:
                                                                                                v48 = v48
                                                                                                var v125 int64 = 914
                                                                                                _ = v125
                                                                                            case -9223372036854775808:
                                                                                                if true {
                                                                                                    if len(v5) >= 2 {
                                                                                                        v5 = v5[1:len(v5)]
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                if false {
                                                                                                    switch "hello" {
                                                                                                    case "\n\t":
                                                                                                        var v126 uint64 = 729
                                                                                                        _ = v126
                                                                                                    case "foo bar":
                                                                                                        if true {
                                                                                                            v127 := GenS4{F0: 128, F1: (v48 + fi49)}
                                                                                                            fmt.Println(v3, v63, fi49)
                                                                                                            v48 -= 488
                                                                                                            _ = v127
                                                                                                        } else {
                                                                                                            v63 += 1000
                                                                                                            if false {
                                                                                                                var v128 int = fi59
                                                                                                                fmt.Println(v60, v63)
                                                                                                                var v129 uint64 = 613
                                                                                                                _ = v128
                                                                                                                _ = v129
                                                                                                            }
                                                                                                            if true {
                                                                                                                v33 += v33
                                                                                                            }
                                                                                                        }
                                                                                                    case "a":
                                                                                                        v3 -= v3
                                                                                                    default:
                                                                                                        v5 = append(v5, "\n\t")
                                                                                                    }
                                                                                                    v60 -= v33
                                                                                                    var v130 int8 = -63
                                                                                                    _ = v130
                                                                                                }
                                                                                            }
                                                                                            v33 = (v60 / ((v33) | 1))
                                                                                            v60 = (v33 + v60)
                                                                                            if (32767 >= 841) {
                                                                                                var v131 int32 = -790
                                                                                                var v132 int8 = -46
                                                                                                _ = v131
                                                                                                _ = v132
                                                                                            } else {
                                                                                                v3 = (v3 + 1.5)
                                                                                                if false {
                                                                                                    if false {
                                                                                                        var v133 uint16 = 831
                                                                                                        fmt.Println(v133, v63)
                                                                                                        if len(v1) >= 2 {
                                                                                                            v1 = v1[1:len(v1)]
                                                                                                        }
                                                                                                        if false {
                                                                                                            var v134 float32 = 0.0
                                                                                                            var v135 int = (fi59 % ((fi49) | 1))
                                                                                                            _ = v134
                                                                                                            _ = v135
                                                                                                        } else {
                                                                                                            v63 = -81
                                                                                                            v3 -= v3
                                                                                                            if (v63 <= v63) {
                                                                                                                switch (fi49 ^ 42) {
                                                                                                                case -357:
                                                                                                                    fmt.Println(fi49, v63, fi49)
                                                                                                                default:
                                                                                                                    var v136 float32 = ((0.1 / 1e10) + -2.25)
                                                                                                                    _ = v136
                                                                                                                }
                                                                                                            } else {
                                                                                                                v3 = float64(float32(v3))
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v133
                                                                                                    } else {
                                                                                                        _ = len(v2)
                                                                                                        v33 += v60
                                                                                                    }
                                                                                                    v137 := GenS9{F0: v33}
                                                                                                    _ = v137
                                                                                                }
                                                                                            }
                                                                                            v48 = fi59
                                                                                            v48 -= (fi59 << ((65535) & 63))
                                                                                        case 1:
                                                                                            _ = len(v2)
                                                                                        }
                                                                                    }
                                                                                    if (false && ((("a" + "a") <= (((("\n\t" + "\n\t") + "hello") + (("\n\t" + "a") + ("" + "a"))) + (((("\n\t" + "hello") + (("\n\t" + "foo bar") + "a")) + ("" + "a")) + "a"))) || true)) {
                                                                                        if (true || (true || true)) {
                                                                                            switch (("foo bar" + "") + "foo bar") {
                                                                                            case "a":
                                                                                                var v138 int = fi59
                                                                                                var v139 int = v48
                                                                                                _ = v138
                                                                                                _ = v139
                                                                                            default:
                                                                                                var v140 uint32 = 42
                                                                                                _ = v140
                                                                                            }
                                                                                            fmt.Println(v48, v3)
                                                                                            var v141 string = ("a" + "hello")
                                                                                            _ = v141
                                                                                        }
                                                                                        var v142 bool = false
                                                                                        _ = v142
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    case 1:
                                                                        if true {
                                                                            fmt.Println(fi49, v48)
                                                                        }
                                                                        if (-1 <= 42) {
                                                                            fmt.Println(v63)
                                                                            if false {
                                                                                v63 = v63
                                                                                switch (fi49 & -9223372036854775808) {
                                                                                case 537:
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                case 2147483647:
                                                                                    if true {
                                                                                        fmt.Println(v63, fi59)
                                                                                    }
                                                                                case 9223372036854775807:
                                                                                    var v143 uint = 460
                                                                                    _ = v143
                                                                                default:
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        var v144 int = (fi49 / ((128) | 1))
                                                                                        _ = v144
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        v4[(("foo bar" + ("" + ("hello" + "\n\t"))) + "foo bar")] = fi59
                                                                        if true {
                                                                            v60 = v60
                                                                            var v145 uint16 = 0
                                                                            v63 = (^(v63))
                                                                            if false {
                                                                                v146 := GenS4{F0: uint8(v145), F1: fi49}
                                                                                var v147 bool = (100 >= 100)
                                                                                _ = v146
                                                                                _ = v147
                                                                            }
                                                                            _ = v145
                                                                        }
                                                                        var v148 int8 = int8(fi49)
                                                                        var v149 uint = 65535
                                                                        var v150 int64 = 791
                                                                        _ = v148
                                                                        _ = v149
                                                                        _ = v150
                                                                    }
                                                                } else {
                                                                    v1 = append(v1, (v60 + v33))
                                                                    if ((((("\n\t" + "") + "") + ("\n\t" + ("hello" + "\n\t"))) + ("a" + "hello")) != "\n\t") {
                                                                        if true {
                                                                            fmt.Println(fi49, v3, v33)
                                                                            switch ((((("a" + ("\n\t" + "\n\t")) + ("\n\t" + ("hello" + ("a" + ("a" + ("hello" + "")))))) + ("a" + "hello")) + "") + "") {
                                                                            case "foo bar":
                                                                                if true {
                                                                                    v151 := GenS12{F0: (v63 % ((1000) | 1)), F1: v63, F2: -2147483648}
                                                                                    _ = v151
                                                                                } else {
                                                                                    switch fi49 {
                                                                                    case -396:
                                                                                        if (-364 <= (v63 | 0)) {
                                                                                            fmt.Println(v33, v48)
                                                                                            v152 := GenS12{F0: v63, F1: int16(v48), F2: -1}
                                                                                            _ = v152
                                                                                        } else {
                                                                                            fmt.Println(v33, v48, fi49)
                                                                                            fmt.Println(v33)
                                                                                        }
                                                                                    case 9223372036854775807:
                                                                                        v63 -= (v63 << ((genFn1(4294967295, 0)) & 63))
                                                                                    }
                                                                                    var v153 uint64 = 18446744073709551615
                                                                                    _ = v153
                                                                                }
                                                                            }
                                                                            switch (fi59 & v48) {
                                                                            case 1:
                                                                                v154 := GenS4{F0: 151, F1: (v48 * -2147483648)}
                                                                                _ = v154
                                                                            default:
                                                                                if false {
                                                                                    fmt.Println(v60, fi59, v3)
                                                                                    var v155 int = -2147483648
                                                                                    v156 := GenS10{F0: 9, F1: int64(v155), F2: 125}
                                                                                    _ = v155
                                                                                    _ = v156
                                                                                } else {
                                                                                    v157 := GenS13{F0: (!((-2.25 * 0.0) < 1.5))}
                                                                                    if true {
                                                                                        if false {
                                                                                            v48 += -722
                                                                                            v60 = v33
                                                                                        }
                                                                                    } else {
                                                                                        v158 := GenS14{F0: true, F1: 42}
                                                                                        if true {
                                                                                            v159 := GenS3{F0: "", F1: 442}
                                                                                            _ = v159
                                                                                        }
                                                                                        var v160 uint64 = 752
                                                                                        _ = v158
                                                                                        _ = v160
                                                                                    }
                                                                                    v33 -= v33
                                                                                    _ = v157
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v161 := GenS15{F0: 320}
                                                                            switch fi49 {
                                                                            case -2147483648:
                                                                                var v162 int = (v48 | v48)
                                                                                _ = v162
                                                                            case 128:
                                                                                var v163 uint64 = 4294967295
                                                                                _ = v163
                                                                            default:
                                                                                switch 42 {
                                                                                case 42:
                                                                                    if ((985 <= 2) || (v60 == v60)) {
                                                                                        fmt.Println(fi59, v48, v60)
                                                                                    } else {
                                                                                        v48 = fi49
                                                                                    }
                                                                                default:
                                                                                    v164 := GenS6{F0: (1 >= 1), F1: (v60 ^ (v33 | (^(v60)))), F2: 100}
                                                                                    _ = v164
                                                                                }
                                                                            }
                                                                            _ = v161
                                                                        }
                                                                        v165 := GenS16{F0: 1, F1: (v60 % ((v33) | 1)), F2: -2147483648}
                                                                        _ = v165
                                                                    } else {
                                                                        var v166 string = "a"
                                                                        v48 = (v48 << ((931) & 63))
                                                                        fmt.Println(v48, fi49)
                                                                        _ = v166
                                                                    }
                                                                    v63 = v63
                                                                }
                                                            }
                                                            var v167 float32 = 1.5
                                                            v60 = v60
                                                            fmt.Println(v3, v167)
                                                            _ = v167
                                                        }
                                                        _ = v104
                                                    } else {
                                                        fmt.Println(v60, v33)
                                                        v168 := GenS17{F0: 0, F1: true, F2: int32(v63)}
                                                        _ = v168
                                                    }
                                                }
                                                _ = v63
                                            case "":
                                                v60 -= (v60 << ((2) & 63))
                                            }
                                            switch -2147483648 {
                                            case 0:
                                                v3 = v3
                                            default:
                                                v60 = (v33 % (((v60 + v33)) | 1))
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v169 := GenS12{F0: -1, F1: 0, F2: 1}
                                                fmt.Println(v3, fi59)
                                                _ = v169
                                            }
                                            var v170 uint32 = v33
                                            _ = v170
                                        }
                                        v3 = (float64(0.1) / v3)
                                        var v171 uint = 2
                                        _ = v171
                                    default:
                                        var v172 float32 = ((-2.25 - ((3.14159 + (-2.25 * 1e10)) - 1.5)) / float32(v3))
                                        _ = v172
                                    }
                                }
                                v173 := GenS18{F0: (-(fi59)), F1: "", F2: 255}
                                delete(v4, "hello")
                                _ = v173
                                _ = v60
                            } else {
                                if false {
                                    var v174 string = ("a" + "a")
                                    switch fi49 {
                                    case 448:
                                        fmt.Println(v3)
                                        v175 := GenS19{F0: 1, F1: v33}
                                        fmt.Println(fi59, fi59)
                                        _ = v175
                                    case 42:
                                        var v176 uint32 = (^(v33))
                                        _ = v176
                                    case -411:
                                        v177 := GenS4{F0: 42, F1: -452}
                                        _ = v177
                                    default:
                                        var v178 uint64 = 4294967295
                                        _ = v178
                                    }
                                    if false {
                                        v179 := GenS20{F0: (fi59 + fi59), F1: 0, F2: (255 <= 255)}
                                        v174 = v174
                                        v180 := GenS21{F0: 127}
                                        _ = v179
                                        _ = v180
                                    } else {
                                        var v181 int32 = 2147483647
                                        _ = v181
                                    }
                                    _ = v174
                                } else {
                                    _ = len(v2)
                                    if (!(-1 < -860)) {
                                        switch "hello" {
                                        case "":
                                            v182 := GenS22{F0: 0, F1: -1}
                                            v48 = fi49
                                            switch fi59 {
                                            case 9223372036854775807:
                                                v33 += (v33 % ((v33) | 1))
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = 65535
                                            }
                                            _ = v182
                                        default:
                                            v3 += v3
                                            _ = len(v5)
                                        }
                                        var v183 uint = 65535
                                        _ = v183
                                    } else {
                                        var v184 float32 = ((float32((1e10 / (v3 - float64(float32(v3))))) - -2.25) / 1.0)
                                        var v185 bool = (1 != v33)
                                        _ = v184
                                        _ = v185
                                    }
                                    switch 294 {
                                    case 1:
                                        var v186 uint32 = v33
                                        _ = v186
                                    case -9223372036854775808:
                                        var v187 int8 = 77
                                        _ = v187
                                    case 32768:
                                        v188 := GenS20{F0: fi49, F1: 65535, F2: false}
                                        _ = v188
                                    }
                                }
                                var v189 int16 = 32767
                                if false {
                                    v48 -= fi49
                                }
                                _ = v189
                            }
                        }
                        var v190 uint16 = 92
                        _ = fi59
                        _ = v190
                    }
                    _ = fi49
                }
                _ = v48
            }
            fmt.Println(v3, v33)
        }
        _ = v34
        _ = v35
    }
    v3 += (v3 / float64(((float32(v3) / 3.14159) + 3.14159)))
    if true {
        var v191 float32 = 0.1
        if false {
            switch "foo bar" {
            case "":
                v192 := GenS16{F0: genFn0(), F1: 2147483647, F2: -1}
                _ = v192
            case "hello":
                v193 := v2[(v33 / ((v33) | 1))]
                _ = v193
            case "\n\t":
                switch "" {
                case "\n\t":
                    v33 -= (v33 ^ v33)
                }
            }
            for fi194 := 0; fi194 < 4; fi194++ {
                if (((!(100 == 100)) || (v191 != 3.14159)) || ((18446744073709551615 != 4294967295) || (1 != 176))) {
                    v195 := GenS23{F0: 204, F1: 1}
                    _ = v195
                }
                _ = fi194
            }
        }
        _ = v191
    }
    if false {
        if true {
            switch 0 {
            case 32768:
                fmt.Println(v3)
            case -882:
                fmt.Println(v3, v3)
            case 975:
                for fi196 := 0; fi196 < 11; fi196++ {
                    var v197 uint8 = 255
                    _ = fi196
                    _ = v197
                }
            }
        } else {
            fmt.Println(v3, v3)
        }
        v198 := GenS21{F0: -110}
        v3 = ((3.14159 * float64(1e10)) / 1e10)
        _ = v198
    } else {
        if false {
            v199 := GenS12{F0: 42, F1: 1, F2: 697}
            _ = v199
        } else {
            var v200 int8 = -79
            for fi201 := 0; fi201 < 10; fi201++ {
                v202 := GenS13{F0: true}
                v200 -= 42
                _ = fi201
                _ = v202
            }
            _ = v200
        }
    }
    var v203 string = "foo bar"
    if true {
        switch 2147483647 {
        case -766:
            switch 1 {
            case -9223372036854775808:
                for fi204 := 0; fi204 < 1; fi204++ {
                    v3 = v3
                    _ = fi204
                }
            }
        case -740:
            var v205 int64 = -9223372036854775808
            _ = v205
        }
        if (449 > 255) {
            switch -2147483648 {
            case 701:
                switch v203 {
                case "foo bar":
                    if (3.14159 <= (v3 / (3.14159 * (1e10 + 0.1)))) {
                        for fi206 := 0; fi206 < 20; fi206++ {
                            var v207 float64 = v3
                            _ = fi206
                            _ = v207
                        }
                    }
                default:
                    fmt.Println(v3, v3)
                }
            case -2147483648:
                v3 += 1e10
            default:
                v203 = v203
            }
            v203 = (v203 + v203)
            v208 := GenS24{F0: genFn0(), F1: (-32768 == 0), F2: 1}
            _ = v208
        } else {
            v203 = v203
            v209 := v4[v203]
            _ = v209
        }
        v33 -= (v33 | (v33 * v33))
    } else {
        v203 = "foo bar"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v203:", v203)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v203
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
