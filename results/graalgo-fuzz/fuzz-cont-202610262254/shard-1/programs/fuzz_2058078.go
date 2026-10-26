package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint8 }
type GenS2 struct { F0 uint; F1 int8; F2 bool }
type GenS3 struct { F0 int; F1 uint64; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 uint; F2 int8 }
type GenS6 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS7 struct { F0 uint32; F1 int64 }
type GenS8 struct { F0 int8; F1 int64; F2 uint64 }
type GenS9 struct { F0 string; F1 uint8 }
type GenS10 struct { F0 int32; F1 string; F2 bool }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint; F1 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = -1
    v3 := []string{(("" + "hello") + "a"), "a", ("foo bar" + (("a" + "") + ("foo bar" + ("foo bar" + "hello"))))}
    v2 = (v2 << ((4294967295) & 63))
    if ((float64((3.14159 + ((1.5 / (1.5 + ((1.5 / 1e10) * float32((1.5 + -2.25))))) / (((3.14159 - 1.5) - 1e10) * ((1e10 + (float32(3.14159) / 0.1)) + 1.5))))) / 1e10) < 1e10) {
        var v4 uint64 = 2
        if true {
            for fi5 := 0; fi5 < 20; fi5++ {
                v2 += (v2 ^ -32768)
                _ = fi5
            }
        } else {
            v4 -= (v4 / ((417) | 1))
            var v6 int16 = -32768
            v2 = (v2 - -586)
            _ = v6
        }
        if len(v3) > 0 {
            v3[0] = "a"
        }
        _ = v4
    } else {
        if (42 >= 2) {
            if false {
                var v7 int = 42
                switch "a" {
                case "a":
                    for fi8 := 0; fi8 < 15; fi8++ {
                        delete(v1, (fi8 / (((v7 + (v7 & v7))) | 1)))
                        var v9 int64 = 270
                        _ = fi8
                        _ = v9
                    }
                case "hello":
                    for fi10 := 0; fi10 < 3; fi10++ {
                        v1[fi10] = 42
                        _ = fi10
                    }
                case "":
                    if ((((("foo bar" + "") != ("" + "\n\t")) && (1 == 2)) && (v2 <= v2)) && (!true)) {
                        for fi11 := 0; fi11 < 11; fi11++ {
                            _ = len(v3)
                            v2 += v2
                            _ = fi11
                        }
                        if false {
                            v12 := GenS1{F0: v7, F1: 0, F2: 180}
                            _ = v12
                        } else {
                            _ = len(v3)
                            if true {
                                v7 = v7
                                var v13 uint8 = 0
                                _ = v13
                            }
                        }
                    } else {
                        switch "a" {
                        case "a":
                            for fi14 := 0; fi14 < 15; fi14++ {
                                delete(v1, (fi14 - v7))
                                _ = fi14
                            }
                        case "foo bar":
                            var v15 int16 = v2
                            _ = v15
                        }
                        var v16 uint32 = 4294967295
                        for fi17 := 0; fi17 < 2; fi17++ {
                            switch fi17 {
                            case -146:
                                switch (-(fi17)) {
                                case 0:
                                    if (0 <= 127) {
                                        _ = len(v1)
                                        if ((v2 ^ (v2 + v2)) >= v2) {
                                            if false {
                                                v2 = (v2 - v2)
                                            }
                                            v7 = (v7 * (fi17 + v7))
                                            var v18 int64 = 9223372036854775807
                                            _ = v18
                                        } else {
                                            var v19 int = fi17
                                            _ = len(v1)
                                            _ = v19
                                        }
                                    } else {
                                        for fi20 := 0; fi20 < 12; fi20++ {
                                            switch fi20 {
                                            case 127:
                                                v21 := GenS1{F0: 127, F1: 0, F2: 202}
                                                _ = v21
                                            case 2147483647:
                                                v2 = (v2 * v2)
                                            case 4294967295:
                                                _ = len(v1)
                                            }
                                            var v22 uint16 = 417
                                            _ = fi20
                                            _ = v22
                                        }
                                        if false {
                                            var v23 int8 = 43
                                            var v24 bool = (true && (-1 != -492))
                                            _ = v23
                                            _ = v24
                                        } else {
                                            v7 = (fi17 & (-(fi17)))
                                        }
                                    }
                                case 9223372036854775807:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case -620:
                                    fmt.Println(v16)
                                default:
                                    v25 := GenS1{F0: 32, F1: (v7 >> ((1) & 63)), F2: 42}
                                    _ = v25
                                }
                            case 334:
                                var v26 uint32 = (^(v16))
                                _ = v26
                            case 9223372036854775807:
                                v16 = v16
                            }
                            _ = fi17
                        }
                        _ = v16
                    }
                default:
                    var v27 int = 42
                    _ = v27
                }
                _ = v7
            }
        } else {
            v28 := GenS1{F0: 282, F1: 42, F2: 128}
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        v1[fi29] = 284
        var v30 uint64 = 692
        v31 := GenS2{F0: 343, F1: -43, F2: (42 > 0)}
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 string = ("" + ("a" + "a"))
    v33 := GenS1{F0: 32767, F1: 127, F2: 0}
    v32 = ((v32 + "a") + "")
    if ((1.5 - ((1e10 - float64((1.5 * 1e10))) + 0.0)) < 0.0) {
        v32 = ""
        var v34 int8 = 127
        fmt.Println(v2, v32)
        _ = v34
    }
    var v35 uint8 = 42
    v35 += (v35 / (((v35 << ((354) & 63))) | 1))
    v1[32768] = 4294967295
    v36 := GenS3{F0: 1, F1: 2, F2: -9223372036854775808}
    if (!true) {
        v37 := GenS3{F0: 32767, F1: 0, F2: -358}
        if false {
            _ = len(v3)
            v38 := GenS3{F0: 9223372036854775807, F1: 607, F2: -1}
            if true {
                switch ((v32 + v32) + ("hello" + v32)) {
                case "foo bar":
                    var v39 string = "foo bar"
                    _ = v39
                }
            }
            _ = v38
        }
        _ = v37
    }
    v32 = ((v32 + (v32 + v32)) + "\n\t")
    v40 := v1[int(v2)]
    if false {
        for fi41 := 0; fi41 < 3; fi41++ {
            switch (fi41 ^ (fi41 % (((-(fi41))) | 1))) {
            case 2147483647:
                for fi42 := 0; fi42 < 14; fi42++ {
                    var v43 int16 = v2
                    switch 753 {
                    case 32767:
                        var v44 string = v32
                        _ = v44
                    case 42:
                        v43 = (v2 ^ (v2 * v2))
                    case 32768:
                        switch (("a" + (v32 + "")) + (v32 + (v32 + ("hello" + v32)))) {
                        case "hello":
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v45 := GenS3{F0: 32768, F1: 0, F2: 311}
                                _ = v45
                            }
                        case "foo bar":
                            v43 += -605
                        }
                    }
                    var v46 string = v32
                    v35 = 42
                    var v47 int8 = -128
                    _ = fi42
                    _ = v43
                    _ = v46
                    _ = v47
                }
            case -817:
                switch (fi41 % (((fi41 >> ((1) & 63))) | 1)) {
                case -61:
                    switch v32 {
                    case "foo bar":
                        if (!true) {
                            v40 = v40
                            v35 = v35
                        } else {
                            var v48 uint = 4294967295
                            _ = v48
                        }
                    case "\n\t":
                        v2 -= int16(fi41)
                    default:
                        v49 := GenS3{F0: 287, F1: uint64(v35), F2: 1}
                        _ = v49
                    }
                case 9223372036854775807:
                    v35 += (v35 ^ 222)
                case 42:
                    v32 = v32
                default:
                    _ = len(v1)
                }
            }
            var v50 bool = true
            _ = fi41
            _ = v50
        }
        if false {
            v51 := GenS4{F0: 2147483647}
            var v52 int64 = 0
            if (969 == (v40 - 212)) {
                for fi53 := 0; fi53 < 18; fi53++ {
                    _ = len(v1)
                    v52 += 0
                    _ = fi53
                }
                fmt.Println(v40)
            }
            _ = v51
            _ = v52
        } else {
            switch ((v32 + "\n\t") + (v32 + "foo bar")) {
            case "":
                switch "" {
                case "a":
                    switch 127 {
                    case 42:
                        if (!(!(540 == -2147483648))) {
                            var v54 uint16 = 108
                            v32 = v32
                            _ = v54
                        } else {
                            _ = len(v1)
                        }
                    case -2147483648:
                        fmt.Println(v2)
                    case 1:
                        if (true && (!(true || ((127 <= 55) || false)))) {
                            for fi55 := 0; fi55 < 14; fi55++ {
                                var v56 int16 = v2
                                var v57 uint64 = 18446744073709551615
                                _ = fi55
                                _ = v56
                                _ = v57
                            }
                        }
                    default:
                        if (-9223372036854775808 <= -870) {
                            _ = len(v1)
                            for fi58 := 0; fi58 < 10; fi58++ {
                                v59 := GenS5{F0: 1, F1: 42, F2: -128}
                                _ = fi58
                                _ = v59
                            }
                            var v60 uint64 = 1
                            _ = v60
                        } else {
                            _ = len(v3)
                            v40 -= uint32(v35)
                        }
                    }
                case "":
                    v40 = v40
                case "\n\t":
                    for fi61 := 0; fi61 < 16; fi61++ {
                        switch (-(fi61)) {
                        case 4294967295:
                            for fi62 := 0; fi62 < 3; fi62++ {
                                if ((float32(0.1) * float32(0.1)) != (3.14159 * 1e10)) {
                                    switch v32 {
                                    case "":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case "a":
                                        if (-236 <= 9223372036854775807) {
                                            v35 -= (v35 & v35)
                                        } else {
                                            if (816 >= 18446744073709551615) {
                                                var v63 float32 = ((((1e10 + ((float32(3.14159) / 0.1) - -2.25)) + (-2.25 - ((0.0 - ((3.14159 * -2.25) + ((3.14159 * 0.0) / ((-2.25 / 1.5) + (float32(3.14159) + 0.0))))) / 0.1))) - (3.14159 / float32((-2.25 / 1e10)))) * 3.14159)
                                                if (v2 <= -1) {
                                                    fmt.Println(v35, v40, v2)
                                                }
                                                _ = len(v3)
                                                _ = v63
                                            }
                                        }
                                    default:
                                        if false {
                                            v40 += v40
                                            fmt.Println(v35, v40)
                                        } else {
                                            v3 = append(v3, v32)
                                        }
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = v32
                                    }
                                    if false {
                                        v3 = append(v3, "hello")
                                        if true {
                                            _ = len(v3)
                                            v35 = v35
                                        }
                                    }
                                } else {
                                    if true {
                                        var v64 int8 = 60
                                        var v65 int32 = 17
                                        _ = v64
                                        _ = v65
                                    } else {
                                        v66 := GenS4{F0: 1}
                                        v2 -= v2
                                        _ = v66
                                    }
                                }
                                v67 := GenS3{F0: 32767, F1: 4294967295, F2: 856}
                                var v68 float64 = float64(3.14159)
                                var v69 uint16 = 32767
                                switch 760 {
                                case 42:
                                    var v70 int16 = v2
                                    _ = v70
                                case 358:
                                    _ = len(v1)
                                case 0:
                                    switch (fi62 - 127) {
                                    case 0:
                                        v69 = v69
                                    case -9223372036854775808:
                                        var v71 uint64 = 0
                                        _ = v71
                                    case 32768:
                                        fmt.Println(fi61, fi61)
                                    default:
                                        if false {
                                            v3 = append(v3, ("foo bar" + ("hello" + v32)))
                                            v72 := GenS6{F0: -1, F1: (v69 * v69), F2: 0}
                                            _ = v72
                                        } else {
                                            fmt.Println(v2, fi61)
                                            var v73 string = "foo bar"
                                            v74 := GenS7{F0: (v40 & (v40 + 0)), F1: 42}
                                            v75 := v1[(fi61 ^ (fi62 ^ -9223372036854775808))]
                                            fmt.Println(fi62)
                                            _ = v73
                                            _ = v74
                                            _ = v75
                                        }
                                    }
                                default:
                                    var v76 float64 = v68
                                    _ = v76
                                }
                                _ = fi62
                                _ = v67
                                _ = v68
                                _ = v69
                            }
                        case 1:
                            v2 = v2
                        case 2147483647:
                            v77 := GenS3{F0: 2147483647, F1: 2, F2: -1}
                            _ = v77
                        }
                        v40 += (v40 + 42)
                        var v78 int32 = -2147483648
                        _ = fi61
                        _ = v78
                    }
                default:
                    for fi79 := 0; fi79 < 16; fi79++ {
                        if (127 == 1) {
                            fmt.Println(v2, v35)
                            for fi80 := 0; fi80 < 8; fi80++ {
                                if (0.1 >= 0.0) {
                                    v2 = 0
                                }
                                v2 += v2
                                var v81 bool = (!true)
                                _ = fi80
                                _ = v81
                            }
                            var v82 float32 = 0.0
                            _ = v82
                        } else {
                            if true {
                                v2 -= (^(v2))
                                v83 := GenS5{F0: 42, F1: 2, F2: 97}
                                _ = v83
                            }
                            var v84 int16 = (v2 % ((v2) | 1))
                            _ = v84
                        }
                        _ = fi79
                    }
                }
            case "\n\t":
                v85 := GenS4{F0: -705}
                _ = v85
            case "foo bar":
                fmt.Println(v32, v35, v32)
            }
            if (630 != int32(v40)) {
                switch (("" + v32) + ((v32 + v32) + (v32 + "\n\t"))) {
                case "\n\t":
                    if true {
                        var v86 uint8 = (v35 - (^(v35)))
                        _ = v86
                    } else {
                        fmt.Println(v2, v2)
                        switch v32 {
                        case "\n\t":
                            for fi87 := 0; fi87 < 7; fi87++ {
                                for fi88 := 0; fi88 < 7; fi88++ {
                                    if (-2.25 <= -2.25) {
                                        v35 = v35
                                    } else {
                                        var v89 uint16 = 65535
                                        var v90 uint16 = (v89 >> ((4294967295) & 63))
                                        _ = v89
                                        _ = v90
                                    }
                                    var v91 int64 = 430
                                    _ = fi88
                                    _ = v91
                                }
                                if (!(-40 == 42)) {
                                    if (false && (true && (0.1 != 0.1))) {
                                        switch 9223372036854775807 {
                                        case 1:
                                            v35 = 180
                                        case 2147483647:
                                            _ = len(v3)
                                        case 4294967295:
                                            switch (fi87 - fi87) {
                                            case 4294967295:
                                                for fi92 := 0; fi92 < 10; fi92++ {
                                                    v32 = "foo bar"
                                                    _ = len(v3)
                                                    _ = fi92
                                                }
                                            case 128:
                                                var v93 bool = (!(int(v35) < (fi87 ^ (fi87 | (fi87 % ((fi87) | 1))))))
                                                _ = v93
                                            case 500:
                                                v94 := v1[(^(fi87))]
                                                _ = v94
                                            default:
                                                var v95 uint32 = v40
                                                _ = v95
                                            }
                                        default:
                                            if (3.14159 == float64(0.0)) {
                                                for fi96 := 0; fi96 < 10; fi96++ {
                                                    switch -497 {
                                                    case 1:
                                                        var v97 string = ("hello" + "a")
                                                        _ = v97
                                                    case -494:
                                                        switch fi96 {
                                                        case 138:
                                                            switch fi96 {
                                                            case 32768:
                                                                if true {
                                                                    fmt.Println(fi87)
                                                                    var v98 float64 = 1.5
                                                                    _ = v98
                                                                } else {
                                                                    var v99 bool = (true || ("foo bar" == "foo bar"))
                                                                    var v100 uint32 = (v40 * (v40 + 2))
                                                                    _ = v100
                                                                    _ = v99
                                                                }
                                                            case -2147483648:
                                                                fmt.Println(v32)
                                                            case 0:
                                                                if len(v3) > 0 {
                                                                    v3[0] = v32
                                                                }
                                                            default:
                                                                v40 += (v40 << ((2) & 63))
                                                            }
                                                        case -630:
                                                            var v101 string = v32
                                                            _ = v101
                                                        case -9223372036854775808:
                                                            var v102 int16 = 493
                                                            _ = v102
                                                        }
                                                    }
                                                    var v103 int64 = 878
                                                    if (!false) {
                                                        fmt.Println(v35)
                                                        if true {
                                                            v32 = v32
                                                            var v104 int16 = (v2 << ((572) & 63))
                                                            fmt.Println(v103, v104)
                                                            _ = v104
                                                        } else {
                                                            if (!false) {
                                                                switch -9223372036854775808 {
                                                                case 4294967295:
                                                                    v2 = -133
                                                                case 0:
                                                                    v105 := GenS5{F0: v103, F1: 65535, F2: -128}
                                                                    _ = v105
                                                                case -613:
                                                                    v40 += v40
                                                                default:
                                                                    fmt.Println(fi87)
                                                                }
                                                            } else {
                                                                var v106 int32 = 42
                                                                v35 += (^(v35))
                                                                v2 -= 1
                                                                _ = v106
                                                            }
                                                            _ = len(v3)
                                                        }
                                                    } else {
                                                        v2 = 1000
                                                        if (v32 != (v32 + ("\n\t" + v32))) {
                                                            fmt.Println(v32, fi96)
                                                            var v107 int = -621
                                                            _ = len(v1)
                                                            fmt.Println(v103)
                                                            v40 -= (v40 + v40)
                                                            v107 = 431
                                                            fmt.Println(v107, v2)
                                                            switch (-(v107)) {
                                                            case 32767:
                                                                var v108 uint32 = v40
                                                                _ = v108
                                                            case 4294967295:
                                                                v109 := GenS8{F0: 1, F1: (v103 + 9223372036854775807), F2: uint64(v40)}
                                                                _ = v109
                                                            default:
                                                                v2 -= v2
                                                            }
                                                            fmt.Println(fi96, fi96)
                                                            v110 := v1[(fi96 % ((v107) | 1))]
                                                            _ = v107
                                                            _ = v110
                                                        } else {
                                                            v103 -= v103
                                                            v35 = 127
                                                            if true {
                                                                v111 := GenS8{F0: -128, F1: (v103 & (v103 - -362)), F2: 42}
                                                                switch fi96 {
                                                                case -386:
                                                                    v112 := GenS3{F0: fi87, F1: 0, F2: -1}
                                                                    _ = v112
                                                                case 620:
                                                                    var v113 uint16 = 32767
                                                                    _ = v113
                                                                case -1:
                                                                    fmt.Println(fi96, v32)
                                                                default:
                                                                    v32 = "\n\t"
                                                                }
                                                                _ = v111
                                                            } else {
                                                                v2 -= v2
                                                                v32 = v32
                                                            }
                                                        }
                                                    }
                                                    _ = fi96
                                                    _ = v103
                                                }
                                                v1[4294967295] = (v40 % ((2147483647) | 1))
                                            }
                                        }
                                        if (true && false) {
                                            var v114 int16 = -32768
                                            fmt.Println(v35, v32, v32)
                                            v35 += (v35 >> ((4294967295) & 63))
                                            _ = v114
                                        }
                                    } else {
                                        var v115 uint8 = v35
                                        switch (v32 + v32) {
                                        case "":
                                            var v116 uint = 65535
                                            _ = v116
                                        }
                                        _ = v115
                                    }
                                    var v117 string = "a"
                                    if false {
                                        switch fi87 {
                                        case -1:
                                            switch fi87 {
                                            case -2147483648:
                                                var v118 string = v32
                                                _ = v118
                                            case 128:
                                                v119 := GenS9{F0: ((((((v32 + v117) + "\n\t") + ("\n\t" + v32)) + (v32 + (v117 + v117))) + "") + v32), F1: (v35 / ((1) | 1))}
                                                _ = v119
                                            default:
                                                var v120 int = fi87
                                                _ = v120
                                            }
                                        case 127:
                                            v121 := v1[(-(fi87))]
                                            _ = v121
                                        case -877:
                                            v2 += 1000
                                        default:
                                            v32 = (v32 + v117)
                                        }
                                        if true {
                                            switch "hello" {
                                            case "hello":
                                                if len(v3) > 0 {
                                                    v3[0] = v32
                                                }
                                            case "a":
                                                if (9223372036854775807 <= -623) {
                                                    fmt.Println(v35)
                                                } else {
                                                    var v122 int = (fi87 ^ 42)
                                                    switch fi87 {
                                                    case 9223372036854775807:
                                                        _ = len(v3)
                                                    case 32768:
                                                        var v123 uint32 = v40
                                                        _ = v123
                                                    case 42:
                                                        if (42 <= 990) {
                                                            switch (v32 + ((("a" + (v117 + v32)) + "\n\t") + v117)) {
                                                            case "a":
                                                                var v124 bool = false
                                                                _ = v124
                                                            case "foo bar":
                                                                switch 128 {
                                                                case 0:
                                                                    var v125 int8 = 1
                                                                    _ = v125
                                                                default:
                                                                    var v126 int64 = 0
                                                                    _ = v126
                                                                }
                                                            }
                                                            if (-292 == (-(v2))) {
                                                                if true {
                                                                    if true {
                                                                        switch fi87 {
                                                                        case 0:
                                                                            for fi127 := 0; fi127 < 20; fi127++ {
                                                                                v32 = v117
                                                                                _ = fi127
                                                                            }
                                                                        default:
                                                                            var v128 bool = ((fi87 ^ 203) < fi87)
                                                                            _ = v128
                                                                        }
                                                                    } else {
                                                                        var v129 uint = 0
                                                                        fmt.Println(v32, v122, v2)
                                                                        v35 += uint8(v40)
                                                                        _ = v129
                                                                    }
                                                                    _ = len(v3)
                                                                    delete(v1, (v122 & (fi87 | (v122 - 128))))
                                                                }
                                                                if (1e10 == -2.25) {
                                                                    var v130 uint = 0
                                                                    if (2 >= uint8(v40)) {
                                                                        v131 := GenS7{F0: v40, F1: 719}
                                                                        _ = v131
                                                                    } else {
                                                                        switch v32 {
                                                                        case "\n\t":
                                                                            _ = len(v3)
                                                                        case "":
                                                                            v2 -= (-(v2))
                                                                        case "foo bar":
                                                                            if (!((v2 + v2) >= v2)) {
                                                                                var v132 int = v122
                                                                                var v133 float32 = (0.1 - 0.1)
                                                                                _ = v132
                                                                                _ = v133
                                                                            } else {
                                                                                v130 -= uint(v40)
                                                                            }
                                                                        }
                                                                        v134 := GenS5{F0: -1, F1: (v130 >> ((v130) & 63)), F2: 94}
                                                                        delete(v1, 128)
                                                                        _ = v134
                                                                    }
                                                                    v135 := GenS6{F0: 430, F1: 2, F2: 2}
                                                                    _ = v130
                                                                    _ = v135
                                                                } else {
                                                                    v117 = (v32 + v117)
                                                                }
                                                                switch (v32 + v32) {
                                                                case "":
                                                                    v1[fi87] = (v40 | (^(v40)))
                                                                case "\n\t":
                                                                    _ = len(v1)
                                                                case "hello":
                                                                    if len(v3) > 0 {
                                                                        v3[0] = v117
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            v136 := v1[-858]
                                                            switch (v32 + v32) {
                                                            case "a":
                                                                switch v117 {
                                                                case "\n\t":
                                                                    for fi137 := 0; fi137 < 3; fi137++ {
                                                                        var v138 uint = 0
                                                                        if (false && (!false)) {
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            switch ("" + ("" + v32)) {
                                                                            case "foo bar":
                                                                                fmt.Println(v35)
                                                                            case "hello":
                                                                                var v139 int16 = (v2 - (v2 / ((v2) | 1)))
                                                                                _ = v139
                                                                            default:
                                                                                var v140 uint16 = 0
                                                                                _ = v140
                                                                            }
                                                                            var v141 float32 = 0.0
                                                                            _ = v141
                                                                        } else {
                                                                            if false {
                                                                                var v142 uint8 = 64
                                                                                _ = v142
                                                                            }
                                                                            if (false || false) {
                                                                                var v143 bool = true
                                                                                var v144 uint64 = 1
                                                                                var v145 int64 = -77
                                                                                _ = v143
                                                                                _ = v144
                                                                                _ = v145
                                                                            } else {
                                                                                var v146 int32 = 557
                                                                                _ = v146
                                                                            }
                                                                            var v147 uint = v138
                                                                            _ = v147
                                                                        }
                                                                        if false {
                                                                            var v148 int32 = 0
                                                                            _ = v148
                                                                        }
                                                                        _ = fi137
                                                                        _ = v138
                                                                    }
                                                                case "a":
                                                                    delete(v1, 127)
                                                                }
                                                            case "":
                                                                var v149 int = (fi87 << ((2) & 63))
                                                                _ = v149
                                                            case "\n\t":
                                                                v150 := GenS1{F0: 4294967295, F1: v122, F2: (v35 | (v35 ^ (v35 | 1)))}
                                                                _ = v150
                                                            default:
                                                                fmt.Println(v136)
                                                            }
                                                            if false {
                                                                fmt.Println(v117)
                                                                fmt.Println(v136)
                                                                fmt.Println(v2, v35, v122)
                                                            } else {
                                                                var v151 int64 = 0
                                                                for fi152 := 0; fi152 < 8; fi152++ {
                                                                    if ((v40 * (v136 / ((2147483647) | 1))) >= v40) {
                                                                        v153 := GenS1{F0: -2147483648, F1: -1, F2: v35}
                                                                        if false {
                                                                            switch (v122 >> ((uint(v40)) & 63)) {
                                                                            case 42:
                                                                                var v154 uint64 = 18446744073709551615
                                                                                _ = v154
                                                                            default:
                                                                                _ = len(v1)
                                                                            }
                                                                            var v155 int8 = int8(v151)
                                                                            _ = v155
                                                                        } else {
                                                                            switch v122 {
                                                                            case 1:
                                                                                v35 = 2
                                                                            case 32767:
                                                                                delete(v1, (^(v122)))
                                                                            case 9223372036854775807:
                                                                                v156 := GenS9{F0: (v117 + ("hello" + (v117 + v32))), F1: (v35 / ((v35) | 1))}
                                                                                _ = v156
                                                                            }
                                                                            var v157 int = (-(fi87))
                                                                            _ = v157
                                                                        }
                                                                        _ = v153
                                                                    }
                                                                    _ = fi152
                                                                }
                                                                for fi158 := 0; fi158 < 8; fi158++ {
                                                                    var v159 uint16 = 812
                                                                    _ = fi158
                                                                    _ = v159
                                                                }
                                                                _ = v151
                                                            }
                                                            _ = v136
                                                        }
                                                    }
                                                    for fi160 := 0; fi160 < 16; fi160++ {
                                                        v122 = (-(v122))
                                                        var v161 int32 = -682
                                                        if false {
                                                            v117 = (((v32 + v117) + ((v117 + (((v32 + v117) + v32) + v117)) + ((("hello" + v32) + v32) + v117))) + ((v117 + "") + ((v117 + (v117 + (v32 + (v117 + ("a" + ("" + v32)))))) + (v32 + ("foo bar" + v32)))))
                                                            fmt.Println(v122, fi160, v117)
                                                        }
                                                        v122 += v122
                                                        var v162 int = fi160
                                                        _ = fi160
                                                        _ = v161
                                                        _ = v162
                                                    }
                                                    _ = v122
                                                }
                                            }
                                        }
                                    } else {
                                        fmt.Println(v2, v40)
                                    }
                                    _ = v117
                                }
                                switch fi87 {
                                case 606:
                                    if true {
                                        _ = len(v3)
                                        var v163 int8 = 54
                                        _ = v163
                                    } else {
                                        v164 := GenS10{F0: -490, F1: v32, F2: ((0.0 + (1.5 - (1e10 - -2.25))) == float64((float32(float64((((float32(0.1) * 1e10) / 1.0) - (float32(0.1) / ((3.14159 - (1e10 * ((3.14159 / 0.1) - 1.5))) / float32(float64(3.14159))))))) * 1.5)))}
                                        v32 = "a"
                                        _ = v164
                                    }
                                case 1:
                                    switch fi87 {
                                    case -9223372036854775808:
                                        var v165 int32 = 42
                                        _ = v165
                                    case 32767:
                                        v166 := GenS3{F0: fi87, F1: 249, F2: 399}
                                        _ = v166
                                    }
                                case 934:
                                    var v167 int64 = 0
                                    _ = v167
                                default:
                                    _ = len(v1)
                                }
                                _ = fi87
                            }
                        }
                        switch "a" {
                        case "":
                            v32 = v32
                        case "foo bar":
                            for fi168 := 0; fi168 < 10; fi168++ {
                                v35 = v35
                                for fi169 := 0; fi169 < 16; fi169++ {
                                    delete(v1, fi168)
                                    _ = fi169
                                }
                                _ = fi168
                            }
                        default:
                            _ = len(v3)
                        }
                    }
                case "a":
                    var v170 int16 = v2
                    _ = v170
                }
                var v171 string = (v32 + ("hello" + v32))
                _ = v171
            }
        }
        if (32767 <= 2) {
            for fi172 := 0; fi172 < 9; fi172++ {
                if (v2 <= int16(fi172)) {
                    v2 = 812
                    var v173 int32 = -2147483648
                    var v174 uint64 = 1
                    _ = v173
                    _ = v174
                }
                v35 -= 137
                _ = fi172
            }
            switch 149 {
            case -1:
                for fi175 := 0; fi175 < 2; fi175++ {
                    v176 := GenS11{F0: v32}
                    switch (("\n\t" + ((((("foo bar" + ((v32 + "foo bar") + "hello")) + v32) + (((v32 + v32) + "") + v32)) + (v32 + ("hello" + "a"))) + v32)) + (((((v32 + "\n\t") + "\n\t") + ((v32 + "foo bar") + "\n\t")) + v32) + ("" + (v32 + "foo bar")))) {
                    case "\n\t":
                        if true {
                            var v177 int8 = 54
                            for fi178 := 0; fi178 < 13; fi178++ {
                                if (v40 < 0) {
                                    v1[32768] = (v40 % (((^(v40))) | 1))
                                    var v179 uint32 = 4294967295
                                    var v180 string = v32
                                    _ = v179
                                    _ = v180
                                } else {
                                    var v181 uint64 = 1
                                    _ = v181
                                }
                                var v182 float64 = 0.0
                                delete(v1, 42)
                                _ = fi178
                                _ = v182
                            }
                            v3 = append(v3, "")
                            _ = v177
                        }
                    case "":
                        v183 := GenS8{F0: 27, F1: -1, F2: 42}
                        _ = v183
                    case "hello":
                        switch int(v2) {
                        case 0:
                            if false {
                                v35 += (v35 << ((4294967295) & 63))
                                switch v32 {
                                case "\n\t":
                                    v40 -= uint32(v35)
                                case "foo bar":
                                    var v184 int64 = -1
                                    _ = v184
                                }
                            } else {
                                switch "" {
                                case "a":
                                    fmt.Println(v32, v32, v2)
                                default:
                                    fmt.Println(fi175)
                                }
                            }
                        case -660:
                            var v185 int8 = 0
                            _ = v185
                        case 1:
                            v40 = (v40 & 2147483648)
                        }
                    default:
                        var v186 float64 = 3.14159
                        _ = v186
                    }
                    _ = fi175
                    _ = v176
                }
            default:
                v2 -= 42
            }
        } else {
            if (!false) {
                v32 = ((v32 + v32) + v32)
                if (!true) {
                    var v187 int8 = -1
                    fmt.Println(v187, v187, v2)
                    _ = v187
                } else {
                    v188 := GenS12{F0: 811, F1: (3.14159 == (3.14159 - 0.1))}
                    for fi189 := 0; fi189 < 9; fi189++ {
                        fmt.Println(fi189, v2, v40)
                        v2 = (v2 >> ((2) & 63))
                        if len(v3) > 0 {
                            v3[0] = v32
                        }
                        _ = fi189
                    }
                    v35 = v35
                    _ = v188
                }
            } else {
                v1[42] = (v40 / ((v40) | 1))
                fmt.Println(v2, v40, v40)
                fmt.Println(v35)
            }
        }
    } else {
        v190 := GenS9{F0: "foo bar", F1: (v35 << ((1) & 63))}
        _ = v190
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v40
}
