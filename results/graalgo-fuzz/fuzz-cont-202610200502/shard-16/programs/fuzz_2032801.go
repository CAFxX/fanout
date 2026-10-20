package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64; F1 int16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 bool; F1 uint64; F2 int32 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int16; F1 int }
type GenS16 struct { F0 int64; F1 int; F2 bool }
type GenS17 struct { F0 int8; F1 int16; F2 uint8 }
type GenS18 struct { F0 int; F1 bool; F2 uint64 }
type GenS19 struct { F0 uint16; F1 int; F2 uint16 }
type GenS20 struct { F0 int8; F1 bool }
type GenS21 struct { F0 uint; F1 int32 }
type GenS22 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS23 struct { F0 uint32 }
func genFn0() int32 {
    var v7 int32 = -1
    var v8 int64 = 42
    if true {
        for fi9 := 0; fi9 < 4; fi9++ {
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 float32 = -2.25
                _ = fi10
                _ = v11
            }
            var v12 uint64 = 2
            v13 := []int{fi9, 127}
            _ = fi9
            _ = v12
            _ = v13
        }
        var v14 bool = (!(!(!false)))
        _ = v14
    } else {
        if (375 > 4294967295) {
            v7 = v7
            fmt.Println(v7, v7)
            for fi15 := 0; fi15 < 1; fi15++ {
                for fi16 := 0; fi16 < 2; fi16++ {
                    if (478 != 2) {
                        v17 := []int{(fi15 % ((fi15) | 1))}
                        if true {
                            var v18 bool = false
                            v7 = (v7 * v7)
                            _ = v18
                        } else {
                            v17 = append(v17, (fi16 * fi15))
                            var v19 int = fi16
                            _ = v19
                        }
                        var v20 int64 = v8
                        fmt.Println(v20, fi15)
                        fmt.Println(fi16, v20)
                        _ = v17
                        _ = v20
                    }
                    v7 = (-(v7))
                    var v21 float64 = 1.5
                    _ = fi16
                    _ = v21
                }
                fmt.Println(v7, v7)
                if true {
                    fmt.Println(fi15, v7)
                    for fi22 := 0; fi22 < 20; fi22++ {
                        v23 := []int{(-(fi15)), fi22}
                        v24 := GenS1{F0: 2147483648}
                        v7 += 0
                        if true {
                            var v25 bool = true
                            v23 = append(v23, -301)
                            var v26 int64 = (v8 & (v8 + 0))
                            _ = v25
                            _ = v26
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                } else {
                    v27 := GenS2{F0: 18446744073709551615, F1: ("" + ("\n\t" + "a"))}
                    _ = v27
                }
                _ = fi15
            }
        } else {
            v7 += v7
        }
        var v28 string = "foo bar"
        _ = v28
    }
    _ = v7
    _ = v8
    return (v7 ^ v7)
}
func genFn1(p0 uint, p1 uint, p2 int32) string {
    if ((((32767 != -154) && (100 >= int8(p2))) || (true || (false || (0 <= 0)))) || true) {
        v29 := []int{9223372036854775807, 42, 1, 0}
        if len(v29) >= 2 {
            v29 = v29[1:len(v29)]
        }
        _ = v29
    }
    var v30 string = ("hello" + "hello")
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return (v30 + v30)
}

func main() {
    var v1 int16 = 1
    v2 := []int{32767, -342}
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    v5 := []string{("a" + "hello")}
    var v6 float32 = (((float32(float64((0.0 - 0.0))) - 1e10) - 1e10) + float32((0.0 / 0.1)))
    switch ("\n\t" + (("" + "foo bar") + ("hello" + "\n\t"))) {
    case "\n\t":
        fmt.Println(v1, v1)
    case "a":
        v31 := GenS2{F0: 42, F1: ("hello" + "")}
        _ = v31
    }
    for fi32 := 0; fi32 < 6; fi32++ {
        for fi33 := 0; fi33 < 11; fi33++ {
            fmt.Println(v1, v6, fi33)
            _ = fi33
        }
        _ = fi32
    }
    if (!false) {
        if false {
            var v34 int = -771
            if true {
                if len(v5) > 0 {
                    v5[0] = (("hello" + ("hello" + "\n\t")) + ("hello" + "foo bar"))
                }
                for fi35 := 0; fi35 < 1; fi35++ {
                    v1 = v1
                    v36 := GenS3{F0: "foo bar", F1: 18446744073709551615}
                    var v37 int32 = 2147483647
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                for fi38 := 0; fi38 < 14; fi38++ {
                    var v39 uint64 = 4294967295
                    if len(v5) > 0 {
                        v5[0] = "\n\t"
                    }
                    for fi40 := 0; fi40 < 11; fi40++ {
                        v41 := GenS3{F0: ("hello" + "\n\t"), F1: 795}
                        delete(v3, (fi40 & v34))
                        var v42 uint16 = 1
                        v34 = fi40
                        var v43 uint8 = 2
                        _ = fi40
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    _ = fi38
                    _ = v39
                }
            }
            _ = v34
        }
    } else {
        if (!false) {
            switch -2147483648 {
            case 2147483647:
                for fi44 := 0; fi44 < 14; fi44++ {
                    var v45 int32 = 1
                    fmt.Println(v1)
                    _ = fi44
                    _ = v45
                }
            }
        }
        v1 += v1
        var v46 uint32 = 1
        _ = v46
    }
    v47 := GenS4{F0: 255, F1: -509}
    if (0.1 != -2.25) {
        var v48 int64 = 0
        _ = v48
    }
    for fi49 := 0; fi49 < 18; fi49++ {
        var v50 float32 = 1e10
        v6 = float32(1e10)
        _ = fi49
        _ = v50
    }
    for fi51 := 0; fi51 < 14; fi51++ {
        for fi52 := 0; fi52 < 9; fi52++ {
            v1 = v1
            _ = fi52
        }
        _ = fi51
    }
    var v53 string = "\n\t"
    if true {
        fmt.Println(v1)
        if (!(!(18446744073709551615 >= 2))) {
            if false {
                fmt.Println(v6, v1, v6)
            } else {
                fmt.Println(v1, v53)
                if ((true || false) || false) {
                    var v54 int8 = 0
                    for fi55 := 0; fi55 < 18; fi55++ {
                        if (42 != 4294967295) {
                            if true {
                                v1 = (v1 * 42)
                                v56 := GenS2{F0: 18446744073709551615, F1: ((v53 + "\n\t") + ("\n\t" + (v53 + v53)))}
                                switch fi55 {
                                case 0:
                                    switch v53 {
                                    case "\n\t":
                                        if (42 < 65535) {
                                            if (v54 > (v54 << ((1) & 63))) {
                                                if (-2147483648 <= -1) {
                                                    if true {
                                                        var v57 int64 = 9223372036854775807
                                                        if false {
                                                            var v58 float32 = v6
                                                            switch (^(fi55)) {
                                                            case 0:
                                                                if true {
                                                                    var v59 int16 = 32767
                                                                    var v60 int = fi55
                                                                    _ = len(v4)
                                                                    _ = v59
                                                                    _ = v60
                                                                } else {
                                                                    if (((v57 ^ v57) > (v57 >> ((0) & 63))) && true) {
                                                                        if true {
                                                                            for fi61 := 0; fi61 < 16; fi61++ {
                                                                                if (true || ((!(1 >= 311)) && false)) {
                                                                                    var v62 int8 = v54
                                                                                    v54 = (v62 * int8(fi61))
                                                                                    fmt.Println(fi55)
                                                                                    _ = v62
                                                                                } else {
                                                                                    fmt.Println(v6, v53)
                                                                                    v63 := GenS5{F0: 231}
                                                                                    v57 = v57
                                                                                    _ = v63
                                                                                }
                                                                                v64 := GenS6{F0: -2147483648, F1: fi55}
                                                                                v65 := GenS3{F0: (("" + v53) + "hello"), F1: 0}
                                                                                switch fi61 {
                                                                                case 42:
                                                                                    if false {
                                                                                        var v66 uint64 = 0
                                                                                        _ = v66
                                                                                    }
                                                                                }
                                                                                _ = fi61
                                                                                _ = v64
                                                                                _ = v65
                                                                            }
                                                                        } else {
                                                                            switch v53 {
                                                                            case "\n\t":
                                                                                var v67 int = (fi55 * fi55)
                                                                                _ = v67
                                                                            case "hello":
                                                                                var v68 uint8 = 0
                                                                                _ = v68
                                                                            case "foo bar":
                                                                                v53 = ("" + "")
                                                                            default:
                                                                                if (v57 > -288) {
                                                                                    var v69 int16 = 1
                                                                                    _ = v69
                                                                                } else {
                                                                                    v70 := GenS4{F0: 0, F1: -199}
                                                                                    if true {
                                                                                        switch (fi55 >> ((65535) & 63)) {
                                                                                        case 693:
                                                                                            if (224 >= 2) {
                                                                                                switch fi55 {
                                                                                                case 127:
                                                                                                    var v71 float32 = 1e10
                                                                                                    _ = v71
                                                                                                case 2147483647:
                                                                                                    for fi72 := 0; fi72 < 13; fi72++ {
                                                                                                        v57 = v57
                                                                                                        _ = fi72
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                var v73 uint8 = 42
                                                                                                v3[int(v1)] = true
                                                                                                _ = v73
                                                                                            }
                                                                                        case -232:
                                                                                            switch fi55 {
                                                                                            case 32768:
                                                                                                fmt.Println(v53, v57)
                                                                                            case -2147483648:
                                                                                                for fi74 := 0; fi74 < 19; fi74++ {
                                                                                                    v58 -= 1e10
                                                                                                    switch (((v53 + ("a" + "foo bar")) + "foo bar") + "\n\t") {
                                                                                                    case "foo bar":
                                                                                                        if false {
                                                                                                            switch (fi55 - (fi74 - fi55)) {
                                                                                                            case 42:
                                                                                                                v57 = v57
                                                                                                            case 1:
                                                                                                                switch fi74 {
                                                                                                                case 85:
                                                                                                                    if true {
                                                                                                                        switch (v53 + v53) {
                                                                                                                        case "":
                                                                                                                            var v75 uint16 = 478
                                                                                                                            _ = v75
                                                                                                                        case "a":
                                                                                                                            v58 = ((((3.14159 - (float32((float64(((v6 * (v6 / (-2.25 * v58))) / -2.25)) * -2.25)) / (0.1 - 1e10))) - 1e10) * 0.1) * v58)
                                                                                                                        case "foo bar":
                                                                                                                            v54 += 1
                                                                                                                        }
                                                                                                                        switch (fi74 | fi55) {
                                                                                                                        case -1:
                                                                                                                            var v76 int32 = -1
                                                                                                                            _ = v76
                                                                                                                        }
                                                                                                                        v1 -= -138
                                                                                                                        v77 := GenS6{F0: int32(v57), F1: (fi74 / (((fi55 + 4294967295)) | 1))}
                                                                                                                        _ = v77
                                                                                                                    }
                                                                                                                case 22:
                                                                                                                    if (!true) {
                                                                                                                        if ((1 >= 18446744073709551615) || (!false)) {
                                                                                                                            v57 -= 367
                                                                                                                            v78 := v3[fi74]
                                                                                                                            var v79 uint8 = 0
                                                                                                                            var v80 uint64 = 18446744073709551615
                                                                                                                            _ = v78
                                                                                                                            _ = v79
                                                                                                                            _ = v80
                                                                                                                        } else {
                                                                                                                            var v81 string = ("\n\t" + ("foo bar" + "a"))
                                                                                                                            _ = v81
                                                                                                                        }
                                                                                                                        v6 -= (0.0 - v6)
                                                                                                                        v58 -= (v6 - v6)
                                                                                                                    } else {
                                                                                                                        var v82 int = 435
                                                                                                                        if false {
                                                                                                                            var v83 int8 = v54
                                                                                                                            fmt.Println(v54, fi74, fi74)
                                                                                                                            v58 += v58
                                                                                                                            var v84 bool = (!false)
                                                                                                                            _ = v83
                                                                                                                            _ = v84
                                                                                                                        } else {
                                                                                                                            v85 := GenS6{F0: -2147483648, F1: fi74}
                                                                                                                            _ = v85
                                                                                                                        }
                                                                                                                        _ = v82
                                                                                                                    }
                                                                                                                case 42:
                                                                                                                    var v86 int64 = v57
                                                                                                                    var v87 string = v53
                                                                                                                    _ = v86
                                                                                                                    _ = v87
                                                                                                                }
                                                                                                            case 32767:
                                                                                                                if (368 == 255) {
                                                                                                                    fmt.Println(fi55, v6, fi55)
                                                                                                                } else {
                                                                                                                    fmt.Println(fi74, v58)
                                                                                                                    v88 := GenS7{F0: v54, F1: -32768}
                                                                                                                    _ = v88
                                                                                                                }
                                                                                                            default:
                                                                                                                var v89 int32 = -1
                                                                                                                _ = v89
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        switch 32768 {
                                                                                                        case 0:
                                                                                                            v90 := GenS8{F0: true}
                                                                                                            fmt.Println(v6)
                                                                                                            v58 = v58
                                                                                                            var v91 uint = 65535
                                                                                                            _ = v90
                                                                                                            _ = v91
                                                                                                        case -263:
                                                                                                            var v92 float32 = v6
                                                                                                            _ = v92
                                                                                                        default:
                                                                                                            v58 += 1e10
                                                                                                        }
                                                                                                    }
                                                                                                    _ = fi74
                                                                                                }
                                                                                            default:
                                                                                                if (!false) {
                                                                                                    if false {
                                                                                                        v1 = 775
                                                                                                    } else {
                                                                                                        var v93 int8 = 25
                                                                                                        _ = v93
                                                                                                    }
                                                                                                    fmt.Println(v6, v53, v54)
                                                                                                    fmt.Println(v58)
                                                                                                }
                                                                                            }
                                                                                        case 1:
                                                                                            v58 += 0.0
                                                                                        default:
                                                                                            if len(v5) >= 2 {
                                                                                                v5 = v5[1:len(v5)]
                                                                                            }
                                                                                        }
                                                                                        var v94 int16 = (-(v1))
                                                                                        _ = v94
                                                                                    } else {
                                                                                        var v95 uint32 = 65535
                                                                                        v1 += (v1 << ((0) & 63))
                                                                                        _ = v95
                                                                                    }
                                                                                    var v96 bool = false
                                                                                    _ = v70
                                                                                    _ = v96
                                                                                }
                                                                            }
                                                                            v1 += v1
                                                                            var v97 int8 = v54
                                                                            _ = v97
                                                                        }
                                                                        v6 += (float32((3.14159 + (0.1 * 1e10))) - v6)
                                                                    }
                                                                }
                                                            case -564:
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            case 2147483647:
                                                                switch "foo bar" {
                                                                case "a":
                                                                    if len(v2) > 0 {
                                                                        v2[0] = (fi55 % ((fi55) | 1))
                                                                    }
                                                                case "":
                                                                    for fi98 := 0; fi98 < 5; fi98++ {
                                                                        v1 = (v1 + v1)
                                                                        _ = fi98
                                                                    }
                                                                case "foo bar":
                                                                    var v99 bool = (1 <= 0)
                                                                    _ = v99
                                                                default:
                                                                    switch ((v53 + v53) + (((v53 + v53) + ("hello" + v53)) + "\n\t")) {
                                                                    case "":
                                                                        switch (fi55 >> ((65535) & 63)) {
                                                                        case 128:
                                                                            v100 := GenS9{F0: 1, F1: (v1 - (^(v1)))}
                                                                            _ = v100
                                                                        case 42:
                                                                            for fi101 := 0; fi101 < 17; fi101++ {
                                                                                var v102 bool = true
                                                                                _ = fi101
                                                                                _ = v102
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            _ = v58
                                                        }
                                                        switch ((v53 + v53) + (v53 + v53)) {
                                                        case "hello":
                                                            switch v53 {
                                                            case "a":
                                                                for fi103 := 0; fi103 < 10; fi103++ {
                                                                    fmt.Println(v57, fi55)
                                                                    v5 = append(v5, "foo bar")
                                                                    _ = fi103
                                                                }
                                                            case "hello":
                                                                if true {
                                                                    v57 = v57
                                                                    for fi104 := 0; fi104 < 6; fi104++ {
                                                                        v57 = (v57 / ((v57) | 1))
                                                                        _ = fi104
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        switch -9223372036854775808 {
                                                                        case 0:
                                                                            if ((-2.25 / (float64((1.5 - (v6 * (0.1 / 3.14159)))) / 1.5)) <= ((1.5 / 3.14159) / 1e10)) {
                                                                                v1 = (^(v1))
                                                                                v57 -= v57
                                                                                switch "a" {
                                                                                case "a":
                                                                                    fmt.Println(v57)
                                                                                case "hello":
                                                                                    var v105 uint64 = 4294967295
                                                                                    _ = v105
                                                                                default:
                                                                                    v6 = float32(0.1)
                                                                                }
                                                                            } else {
                                                                                var v106 uint16 = 32767
                                                                                var v107 int32 = 42
                                                                                _ = v106
                                                                                _ = v107
                                                                            }
                                                                        case -170:
                                                                            if true {
                                                                                fmt.Println(fi55, v54)
                                                                                fmt.Println(v1, fi55)
                                                                                if (!(true || ((914 > 0) || (0 < 42)))) {
                                                                                    var v108 uint64 = 555
                                                                                    fmt.Println(fi55, v57, v108)
                                                                                    var v109 string = ("\n\t" + ((v53 + v53) + v53))
                                                                                    _ = v108
                                                                                    _ = v109
                                                                                } else {
                                                                                    v57 += 0
                                                                                }
                                                                            } else {
                                                                                switch fi55 {
                                                                                case 42:
                                                                                    fmt.Println(v1)
                                                                                case 2147483647:
                                                                                    var v110 int = fi55
                                                                                    _ = v110
                                                                                case -1:
                                                                                    for fi111 := 0; fi111 < 16; fi111++ {
                                                                                        if false {
                                                                                            delete(v4, 778)
                                                                                            var v112 string = v53
                                                                                            _ = v112
                                                                                        } else {
                                                                                            if true {
                                                                                                if true {
                                                                                                    if len(v2) > 0 {
                                                                                                        v2[0] = fi111
                                                                                                    }
                                                                                                }
                                                                                                var v113 float64 = -2.25
                                                                                                _ = v113
                                                                                            } else {
                                                                                                if true {
                                                                                                    if true {
                                                                                                        var v114 int32 = 42
                                                                                                        var v115 uint32 = 1
                                                                                                        switch v53 {
                                                                                                        case "foo bar":
                                                                                                            var v116 string = "foo bar"
                                                                                                            _ = v116
                                                                                                        case "":
                                                                                                            v117 := GenS2{F0: uint64(v115), F1: "hello"}
                                                                                                            _ = v117
                                                                                                        case "\n\t":
                                                                                                            var v118 uint8 = 2
                                                                                                            _ = v118
                                                                                                        default:
                                                                                                            if true {
                                                                                                                if true {
                                                                                                                    v114 += (v114 * v114)
                                                                                                                } else {
                                                                                                                    if false {
                                                                                                                        _ = len(v3)
                                                                                                                        v1 -= v1
                                                                                                                        switch (fi111 | (-(fi55))) {
                                                                                                                        case 1:
                                                                                                                            if ((v1 & v1) != (^(v1))) {
                                                                                                                                var v119 int32 = v114
                                                                                                                                if (((v114 / (((^(v119))) | 1)) >= (v119 | genFn0())) || (!false)) {
                                                                                                                                    v1 = (v1 ^ (v1 % ((v1) | 1)))
                                                                                                                                    delete(v4, v115)
                                                                                                                                } else {
                                                                                                                                    v4[4294967295] = (-(fi55))
                                                                                                                                    fmt.Println(v119)
                                                                                                                                    v120 := GenS4{F0: 42, F1: 42}
                                                                                                                                    _ = v120
                                                                                                                                }
                                                                                                                                _ = v119
                                                                                                                            } else {
                                                                                                                                var v121 string = (v53 + v53)
                                                                                                                                v4[(v115 & (v115 ^ (v115 * v115)))] = 184
                                                                                                                                fmt.Println(fi111, v54, v1)
                                                                                                                                v122 := GenS7{F0: (v54 << ((902) & 63)), F1: (v1 + v1)}
                                                                                                                                fmt.Println(v121, v57, fi55)
                                                                                                                                _ = len(v5)
                                                                                                                                switch 713 {
                                                                                                                                case 291:
                                                                                                                                    if false {
                                                                                                                                        fmt.Println(v54)
                                                                                                                                        switch (("\n\t" + ((v53 + "hello") + v121)) + "\n\t") {
                                                                                                                                        case "foo bar":
                                                                                                                                            _ = len(v4)
                                                                                                                                        case "":
                                                                                                                                            delete(v4, (^(v115)))
                                                                                                                                        case "a":
                                                                                                                                            delete(v4, v115)
                                                                                                                                        default:
                                                                                                                                            if (255 >= 1) {
                                                                                                                                                v123 := GenS6{F0: v114, F1: (-(fi111))}
                                                                                                                                                v124 := GenS8{F0: false}
                                                                                                                                                var v125 int = (^(fi55))
                                                                                                                                                var v126 uint = 42
                                                                                                                                                if len(v5) > 0 {
                                                                                                                                                    v5[0] = "a"
                                                                                                                                                }
                                                                                                                                                _ = v123
                                                                                                                                                _ = v124
                                                                                                                                                _ = v125
                                                                                                                                                _ = v126
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                case 133:
                                                                                                                                    fmt.Println(v54, fi111)
                                                                                                                                default:
                                                                                                                                    var v127 int8 = -128
                                                                                                                                    _ = v127
                                                                                                                                }
                                                                                                                                _ = v121
                                                                                                                                _ = v122
                                                                                                                            }
                                                                                                                        case 11:
                                                                                                                            if ((!(false || ((v1 / (((v1 - (^(v1)))) | 1)) < v1))) || (!(65535 <= 1))) {
                                                                                                                                switch (fi55 * -2147483648) {
                                                                                                                                case 93:
                                                                                                                                    v6 = 1.5
                                                                                                                                    fmt.Println(v115, v114, fi55)
                                                                                                                                    var v128 int32 = -2147483648
                                                                                                                                    _ = v128
                                                                                                                                case 42:
                                                                                                                                    v5 = append(v5, ("a" + ""))
                                                                                                                                case 9223372036854775807:
                                                                                                                                    _ = len(v2)
                                                                                                                                default:
                                                                                                                                    delete(v4, (v115 / (((v115 * (v115 & 65535))) | 1)))
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        v115 = (v115 * (v115 >> ((0) & 63)))
                                                                                                                        v129 := v3[(fi111 + (fi111 | -2147483648))]
                                                                                                                        fmt.Println(v114, fi55)
                                                                                                                        _ = v129
                                                                                                                    }
                                                                                                                    fmt.Println(v1)
                                                                                                                }
                                                                                                                switch (-(fi55)) {
                                                                                                                case 9223372036854775807:
                                                                                                                    if (false && false) {
                                                                                                                        switch (^(fi111)) {
                                                                                                                        case 2147483647:
                                                                                                                            switch int(v1) {
                                                                                                                            case 0:
                                                                                                                                fmt.Println(v1, v6, v53)
                                                                                                                            case -9223372036854775808:
                                                                                                                                if len(v2) > 0 {
                                                                                                                                    v2[0] = (fi55 >> ((65535) & 63))
                                                                                                                                }
                                                                                                                            case 128:
                                                                                                                                if false {
                                                                                                                                    v130 := GenS10{F0: 9223372036854775807}
                                                                                                                                    _ = v130
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                var v131 int16 = v1
                                                                                                                                _ = v131
                                                                                                                            }
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        v2 = append(v2, (fi55 / (((fi55 & (fi111 & fi55))) | 1)))
                                                                                                                        if true {
                                                                                                                            var v132 int = 691
                                                                                                                            v133 := GenS11{F0: false, F1: 4294967295, F2: (v114 >> ((2) & 63))}
                                                                                                                            v57 += (v57 * 459)
                                                                                                                            _ = v132
                                                                                                                            _ = v133
                                                                                                                        }
                                                                                                                    }
                                                                                                                case 127:
                                                                                                                    v1 -= -651
                                                                                                                case -2:
                                                                                                                    switch (fi55 * 2147483647) {
                                                                                                                    case 165:
                                                                                                                        var v134 uint = 2
                                                                                                                        _ = v134
                                                                                                                    default:
                                                                                                                        fmt.Println(v114, v114)
                                                                                                                    }
                                                                                                                default:
                                                                                                                    switch (fi111 % ((4294967295) | 1)) {
                                                                                                                    case 2147483647:
                                                                                                                        v135 := GenS12{F0: false}
                                                                                                                        _ = v135
                                                                                                                    }
                                                                                                                }
                                                                                                                if true {
                                                                                                                    fmt.Println(fi55, v6)
                                                                                                                    if ((2 <= 2) || (42 < 2)) {
                                                                                                                        fmt.Println(v1, v53, v115)
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    if (0 >= 0) {
                                                                                                                        v57 -= 0
                                                                                                                    }
                                                                                                                    v115 = 4294967295
                                                                                                                    v6 -= 1.5
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v114
                                                                                                        _ = v115
                                                                                                    } else {
                                                                                                        v136 := v3[fi55]
                                                                                                        if true {
                                                                                                            var v137 bool = v136
                                                                                                            _ = v137
                                                                                                        }
                                                                                                        _ = v136
                                                                                                    }
                                                                                                    _ = len(v2)
                                                                                                    var v138 int16 = 32767
                                                                                                    _ = v138
                                                                                                }
                                                                                            }
                                                                                            switch (fi55 ^ fi111) {
                                                                                            case -9223372036854775808:
                                                                                                var v139 uint32 = 429
                                                                                                _ = v139
                                                                                            case 2147483647:
                                                                                                v140 := GenS13{F0: 1}
                                                                                                _ = v140
                                                                                            default:
                                                                                                if false {
                                                                                                    _ = len(v4)
                                                                                                    fmt.Println(v54)
                                                                                                } else {
                                                                                                    var v141 string = (v53 + v53)
                                                                                                    fmt.Println(v141, v57, v53)
                                                                                                    if true {
                                                                                                        v57 -= 1
                                                                                                        var v142 float64 = 1.5
                                                                                                        _ = v142
                                                                                                    } else {
                                                                                                        fmt.Println(v57)
                                                                                                        var v143 uint64 = 0
                                                                                                        _ = v143
                                                                                                    }
                                                                                                    _ = v141
                                                                                                }
                                                                                                var v144 bool = (0 < 0)
                                                                                                _ = v144
                                                                                            }
                                                                                            if ((false || (!(!(false || false)))) || ((v54 / (((v54 % ((127) | 1))) | 1)) == (v54 / (((v54 & v54)) | 1)))) {
                                                                                                v57 = v57
                                                                                                switch -2147483648 {
                                                                                                case 585:
                                                                                                    delete(v3, fi111)
                                                                                                case 9223372036854775807:
                                                                                                    var v145 uint = 674
                                                                                                    _ = v145
                                                                                                case 32767:
                                                                                                    var v146 string = ((v53 + (("" + (v53 + "a")) + (v53 + v53))) + (v53 + "\n\t"))
                                                                                                    _ = v146
                                                                                                }
                                                                                                fmt.Println(v1, v57)
                                                                                            } else {
                                                                                                if (true || true) {
                                                                                                    v53 = (v53 + (v53 + ((v53 + v53) + ((("a" + v53) + v53) + "hello"))))
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        var v147 float64 = 1.5
                                                                                        switch fi55 {
                                                                                        case -406:
                                                                                            var v148 int = fi111
                                                                                            _ = v148
                                                                                        }
                                                                                        fmt.Println(v1)
                                                                                        var v149 bool = true
                                                                                        _ = fi111
                                                                                        _ = v147
                                                                                        _ = v149
                                                                                    }
                                                                                }
                                                                                var v150 uint64 = 0
                                                                                v150 = 2
                                                                                _ = v150
                                                                            }
                                                                        case 138:
                                                                            var v151 int64 = v57
                                                                            _ = v151
                                                                        }
                                                                    } else {
                                                                        v54 = 49
                                                                        for fi152 := 0; fi152 < 17; fi152++ {
                                                                            var v153 bool = (32767 < fi55)
                                                                            v53 = (v53 + v53)
                                                                            v53 = "foo bar"
                                                                            _ = fi152
                                                                            _ = v153
                                                                        }
                                                                        if ((v54 - v54) == (v54 & (-(v54)))) {
                                                                            v154 := GenS6{F0: genFn0(), F1: 127}
                                                                            _ = v154
                                                                        } else {
                                                                            v57 = v57
                                                                            switch (fi55 ^ (fi55 << ((222) & 63))) {
                                                                            case -271:
                                                                                if true {
                                                                                    var v155 uint8 = 0
                                                                                    _ = v155
                                                                                } else {
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = v53
                                                                                    }
                                                                                }
                                                                            case -9223372036854775808:
                                                                                v156 := v3[fi55]
                                                                                _ = v156
                                                                            case 32767:
                                                                                v57 -= 510
                                                                            }
                                                                            if (!true) {
                                                                                if ((fi55 & (^(fi55))) <= (fi55 ^ (fi55 % (((fi55 + (fi55 | 9223372036854775807))) | 1)))) {
                                                                                    v157 := GenS10{F0: (fi55 | fi55)}
                                                                                    _ = v157
                                                                                } else {
                                                                                    for fi158 := 0; fi158 < 7; fi158++ {
                                                                                        v57 -= (-(v57))
                                                                                        v57 = v57
                                                                                        _ = fi158
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v159 uint32 = 42
                                                                                var v160 int64 = (^(v57))
                                                                                fmt.Println(v53, v1, v54)
                                                                                _ = v159
                                                                                _ = v160
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        case "a":
                                                            v54 = 115
                                                        }
                                                        _ = v57
                                                    } else {
                                                        if (((float64(v6) >= (-2.25 + -2.25)) || true) || true) {
                                                            var v161 uint8 = 2
                                                            v162 := GenS14{F0: fi55}
                                                            _ = v161
                                                            _ = v162
                                                        }
                                                    }
                                                    if (554 < 620) {
                                                        var v163 bool = (3.14159 >= 1.5)
                                                        _ = v163
                                                    } else {
                                                        v5 = append(v5, "hello")
                                                        var v164 float64 = 3.14159
                                                        _ = v164
                                                    }
                                                    fmt.Println(v54, fi55)
                                                }
                                                v165 := GenS8{F0: true}
                                                if (9223372036854775807 < 617) {
                                                    _ = len(v5)
                                                    var v166 int16 = v1
                                                    switch (fi55 % (((fi55 * -1)) | 1)) {
                                                    case 643:
                                                        if (false || (61 < 42)) {
                                                            for fi167 := 0; fi167 < 5; fi167++ {
                                                                if false {
                                                                    _ = len(v4)
                                                                } else {
                                                                    v168 := v3[fi167]
                                                                    _ = v168
                                                                }
                                                                _ = fi167
                                                            }
                                                            for fi169 := 0; fi169 < 17; fi169++ {
                                                                v166 = (v1 ^ 317)
                                                                _ = fi169
                                                            }
                                                        } else {
                                                            v4[199] = fi55
                                                            var v170 int8 = 100
                                                            v6 = ((3.14159 / v6) * v6)
                                                            _ = v170
                                                        }
                                                    case -9223372036854775808:
                                                        var v171 uint32 = 941
                                                        _ = v171
                                                    case 2147483647:
                                                        var v172 int16 = (v1 >> ((124) & 63))
                                                        _ = v172
                                                    }
                                                    _ = v166
                                                } else {
                                                    v6 = float32(0.0)
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                _ = v165
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    default:
                                        if (!((((v53 + "") + v53) + "\n\t") > v53)) {
                                            for fi173 := 0; fi173 < 8; fi173++ {
                                                v174 := GenS4{F0: 42, F1: -730}
                                                fmt.Println(v6, v54)
                                                _ = fi173
                                                _ = v174
                                            }
                                        }
                                    }
                                case -80:
                                    for fi175 := 0; fi175 < 17; fi175++ {
                                        var v176 uint16 = 65535
                                        v176 = (^(v176))
                                        if (!true) {
                                            delete(v3, 1)
                                        } else {
                                            v177 := GenS6{F0: genFn0(), F1: 128}
                                            fmt.Println(v176, v176)
                                            _ = v177
                                        }
                                        _ = fi175
                                        _ = v176
                                    }
                                case 1:
                                    var v178 int16 = 42
                                    _ = v178
                                default:
                                    delete(v4, 610)
                                }
                                _ = v56
                            } else {
                                v179 := GenS11{F0: ((v54 - v54) <= (^(v54))), F1: 42, F2: -2147483648}
                                _ = v179
                            }
                            if (((^(v1)) != v1) && false) {
                                v1 -= v1
                                switch fi55 {
                                case 32767:
                                    if true {
                                        fmt.Println(v54)
                                        switch fi55 {
                                        case 750:
                                            var v180 float32 = (1.5 / float32(float64(v6)))
                                            _ = v180
                                        case -2147483648:
                                            v181 := GenS15{F0: -782, F1: (fi55 - (-(fi55)))}
                                            _ = v181
                                        }
                                    }
                                case -472:
                                    var v182 uint8 = 2
                                    _ = v182
                                default:
                                    v183 := GenS3{F0: "", F1: 42}
                                    _ = v183
                                }
                                if true {
                                    v53 = ((("\n\t" + v53) + (v53 + (v53 + (v53 + (v53 + "a"))))) + v53)
                                    v1 -= (v1 >> ((255) & 63))
                                    for fi184 := 0; fi184 < 1; fi184++ {
                                        v185 := GenS16{F0: -9223372036854775808, F1: fi55, F2: false}
                                        fmt.Println(fi184, v54)
                                        v53 = v53
                                        _ = fi184
                                        _ = v185
                                    }
                                }
                            } else {
                                delete(v3, -1)
                            }
                        }
                        _ = fi55
                    }
                    _ = v54
                }
                var v186 uint32 = 961
                _ = v186
            }
            v6 -= 1.5
        }
    } else {
        fmt.Println(v1, v6)
    }
    v187 := v3[0]
    if (!(682 >= 42)) {
        v188 := GenS17{F0: 1, F1: (v1 << ((1) & 63)), F2: 1}
        for fi189 := 0; fi189 < 18; fi189++ {
            v190 := GenS16{F0: 0, F1: fi189, F2: ((1 > 0) || (60 < 695))}
            _ = fi189
            _ = v190
        }
        v191 := GenS18{F0: 1, F1: true, F2: 822}
        _ = v188
        _ = v191
    } else {
        v187 = ("hello" < ("a" + v53))
        for fi192 := 0; fi192 < 12; fi192++ {
            if true {
                fmt.Println(fi192, fi192)
                v187 = v187
                if true {
                    var v193 int32 = 42
                    _ = v193
                }
            }
            v194 := GenS19{F0: 1, F1: 1, F2: 530}
            _ = fi192
            _ = v194
        }
        v6 = v6
    }
    var v195 uint16 = 111
    var v196 string = v53
    if v187 {
        v187 = v187
        switch 32767 {
        case 1:
            var v197 uint64 = 18446744073709551615
            _ = v197
        case 42:
            v195 += (v195 / (((v195 << ((351) & 63))) | 1))
        case 128:
            switch v196 {
            case "":
                _ = len(v5)
            case "hello":
                fmt.Println(v196)
            case "foo bar":
                fmt.Println(v6)
            default:
                fmt.Println(v1, v196)
            }
        default:
            if false {
                switch v196 {
                case "foo bar":
                    if ((v1 >= -1) && (v6 <= v6)) {
                        v198 := GenS20{F0: -70, F1: v187}
                        delete(v3, 9223372036854775807)
                        _ = v198
                    } else {
                        v199 := GenS21{F0: 1, F1: 42}
                        var v200 int = 127
                        _ = v199
                        _ = v200
                    }
                case "hello":
                    _ = len(v3)
                case "\n\t":
                    var v201 uint16 = (v195 << ((2) & 63))
                    _ = v201
                default:
                    v202 := v3[-441]
                    _ = v202
                }
                for fi203 := 0; fi203 < 3; fi203++ {
                    var v204 int64 = -249
                    v205 := GenS7{F0: 127, F1: (-(v1))}
                    v206 := GenS20{F0: -104, F1: ("a" < v196)}
                    _ = fi203
                    _ = v204
                    _ = v205
                    _ = v206
                }
                var v207 int8 = -1
                _ = v207
            } else {
                if v187 {
                    if v187 {
                        var v208 bool = false
                        if v187 {
                            v195 = (^(v195))
                        } else {
                            if v208 {
                                fmt.Println(v195, v53, v53)
                                switch 0 {
                                case 1:
                                    switch -157 {
                                    case 2147483647:
                                        var v209 uint8 = 16
                                        _ = v209
                                    default:
                                        fmt.Println(v53, v6, v187)
                                    }
                                }
                            } else {
                                switch (v196 + (v196 + "foo bar")) {
                                case "a":
                                    for fi210 := 0; fi210 < 4; fi210++ {
                                        _ = len(v3)
                                        _ = fi210
                                    }
                                }
                                var v211 uint = 0
                                _ = v211
                            }
                            var v212 int = 9223372036854775807
                            _ = v212
                        }
                        for fi213 := 0; fi213 < 15; fi213++ {
                            if (19 != 42) {
                                v214 := GenS16{F0: -336, F1: 2147483647, F2: (100 < 127)}
                                v215 := GenS4{F0: 479, F1: 0}
                                _ = v214
                                _ = v215
                            } else {
                                if false {
                                    if v187 {
                                        var v216 uint8 = 194
                                        v217 := GenS22{F0: uint64(v216), F1: (v195 << ((18446744073709551615) & 63)), F2: 65535}
                                        _ = v216
                                        _ = v217
                                    }
                                    if true {
                                        v218 := GenS6{F0: 0, F1: 13}
                                        _ = v218
                                    }
                                    v187 = v187
                                }
                            }
                            delete(v4, 284)
                            _ = fi213
                        }
                        _ = v208
                    }
                } else {
                    v219 := GenS23{F0: 2}
                    _ = v219
                }
            }
        }
    } else {
        for fi220 := 0; fi220 < 6; fi220++ {
            var v221 uint = 569
            _ = fi220
            _ = v221
        }
        v3[-651] = v187
        fmt.Println(v187, v187)
    }
    _ = len(v3)
    fmt.Println(v187, v195)
    for fi222 := 0; fi222 < 10; fi222++ {
        delete(v4, 42)
        var v223 uint16 = v195
        _ = fi222
        _ = v223
    }
    v196 = v196
    fmt.Println("v1:", v1)
    fmt.Println("v187:", v187)
    fmt.Println("v195:", v195)
    fmt.Println("v196:", v196)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v187
    _ = v195
    _ = v196
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v53
    _ = v6
}
