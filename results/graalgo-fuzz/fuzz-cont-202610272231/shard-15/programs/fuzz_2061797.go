package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int16; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int; F1 int16; F2 bool }
type GenS6 struct { F0 int; F1 int16; F2 int32 }
type GenS7 struct { F0 uint64; F1 int16 }
type GenS8 struct { F0 int; F1 uint32 }
type GenS9 struct { F0 uint32; F1 uint32 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int8; F1 string; F2 string }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint32; F1 uint32 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 int8; F1 uint16 }
type GenS18 struct { F0 int; F1 int; F2 string }
type GenS19 struct { F0 bool; F1 int8 }
type GenS20 struct { F0 uint32; F1 string; F2 int64 }
type GenS21 struct { F0 uint32; F1 uint16; F2 int }
type GenS22 struct { F0 uint }
type GenS23 struct { F0 int32 }
type GenS24 struct { F0 string; F1 int16 }
func genFn0() int {
    v6 := []int{}
    if true {
        var v7 int8 = 100
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 int32 = -1
            switch (fi8 / ((127) | 1)) {
            case 127:
                v7 = (^(v7))
            }
            switch fi8 {
            case 127:
                v9 -= v9
            case 2147483647:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case 203:
                v7 = v7
            default:
                switch "a" {
                case "hello":
                    var v10 uint32 = 552
                    _ = v10
                }
            }
            _ = fi8
            _ = v9
        }
        switch "foo bar" {
        case "foo bar":
            if true {
                for fi11 := 0; fi11 < 8; fi11++ {
                    if false {
                        for fi12 := 0; fi12 < 9; fi12++ {
                            fmt.Println(fi12, fi11)
                            _ = fi12
                        }
                        v7 += 67
                        v13 := GenS1{F0: 2147483647, F1: 0}
                        _ = v13
                    } else {
                        switch fi11 {
                        case -825:
                            v14 := GenS2{F0: (v7 >> ((116) & 63)), F1: 0}
                            _ = v14
                        case 32767:
                            for fi15 := 0; fi15 < 11; fi15++ {
                                switch (-(fi11)) {
                                case -343:
                                    v7 -= v7
                                }
                                _ = fi15
                            }
                        case 42:
                            for fi16 := 0; fi16 < 12; fi16++ {
                                if true {
                                    switch (fi16 + (fi11 | fi11)) {
                                    case 9223372036854775807:
                                        var v17 int32 = 1
                                        _ = v17
                                    case 127:
                                        if false {
                                            switch ("\n\t" + "") {
                                            case "a":
                                                var v18 int32 = 0
                                                _ = v18
                                            case "":
                                                fmt.Println(v7, fi16, fi16)
                                            case "hello":
                                                v19 := GenS2{F0: (v7 << ((2) & 63)), F1: -9223372036854775808}
                                                _ = v19
                                            default:
                                                if false {
                                                    var v20 string = "a"
                                                    _ = v20
                                                } else {
                                                    fmt.Println(fi16, fi16)
                                                    v7 += 127
                                                }
                                                v7 += -1
                                            }
                                            if ((v7 + v7) != v7) {
                                                switch (fi11 | -978) {
                                                case 513:
                                                    if (370 >= 2) {
                                                        var v21 float32 = 3.14159
                                                        var v22 int32 = 2147483647
                                                        _ = v21
                                                        _ = v22
                                                    }
                                                default:
                                                    v7 = v7
                                                }
                                            } else {
                                                fmt.Println(fi16, fi11, fi11)
                                                switch fi16 {
                                                case 728:
                                                    var v23 string = "foo bar"
                                                    _ = v23
                                                case 127:
                                                    fmt.Println(fi11, fi11)
                                                case 0:
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                }
                                            }
                                        }
                                    case 462:
                                        v7 -= (v7 >> ((0) & 63))
                                    }
                                }
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                var v24 uint8 = 39
                                v25 := GenS3{F0: 1, F1: 32767, F2: 4294967295}
                                v26 := GenS4{F0: 72, F1: (true && (true || (!(1.5 != 0.0))))}
                                if (31 >= 2147483647) {
                                    var v27 uint32 = 1
                                    v24 = v24
                                    switch fi16 {
                                    case 127:
                                        _ = len(v6)
                                    case 4294967295:
                                        var v28 uint16 = 690
                                        _ = v28
                                    }
                                    _ = v27
                                }
                                _ = fi16
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                        }
                        fmt.Println(fi11, v7)
                    }
                    for fi29 := 0; fi29 < 16; fi29++ {
                        switch fi11 {
                        case 42:
                            var v30 uint = 65535
                            _ = v30
                        case -827:
                            v31 := GenS3{F0: 1000, F1: 32767, F2: 4294967295}
                            _ = v31
                        default:
                            var v32 uint8 = 1
                            _ = v32
                        }
                        switch ((("hello" + "\n\t") + ("" + (("a" + "foo bar") + "foo bar"))) + "foo bar") {
                        case "foo bar":
                            switch (fi11 >> ((65535) & 63)) {
                            case 0:
                                v7 += v7
                            case 2147483647:
                                if true {
                                    switch fi29 {
                                    case 1:
                                        v33 := GenS2{F0: (v7 / ((v7) | 1)), F1: 101}
                                        _ = v33
                                    case 708:
                                        v7 -= (v7 * (-(v7)))
                                    case -776:
                                        v7 = (v7 >> ((255) & 63))
                                    }
                                    fmt.Println(fi29, fi29, fi11)
                                    var v34 uint64 = 2
                                    _ = v34
                                } else {
                                    if false {
                                        v35 := GenS5{F0: fi29, F1: 670, F2: true}
                                        var v36 int64 = -1
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v35
                                        _ = v36
                                    } else {
                                        switch "foo bar" {
                                        case "":
                                            v37 := GenS4{F0: 65535, F1: false}
                                            _ = v37
                                        case "a":
                                            var v38 uint8 = 112
                                            _ = v38
                                        }
                                        if (true && (70 <= -862)) {
                                            if true {
                                                fmt.Println(fi29)
                                                v39 := GenS6{F0: (fi29 / ((fi29) | 1)), F1: 948, F2: 97}
                                                _ = v39
                                            } else {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                if ("foo bar" == ("foo bar" + ((("hello" + (((("\n\t" + "a") + (("foo bar" + "a") + "a")) + "a") + "hello")) + "") + ("\n\t" + "\n\t")))) {
                                                    v7 += v7
                                                    v40 := GenS3{F0: 387, F1: 494, F2: 2147483648}
                                                    var v41 string = (("" + "hello") + ("hello" + "hello"))
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    var v42 uint8 = 153
                                                    _ = v40
                                                    _ = v41
                                                    _ = v42
                                                } else {
                                                    v43 := GenS3{F0: 32767, F1: int16(v7), F2: 2147483647}
                                                    _ = v43
                                                }
                                            }
                                            fmt.Println(fi11, fi29, v7)
                                            v44 := GenS5{F0: (fi29 - (fi11 * fi11)), F1: int16(v7), F2: (3.14159 > -2.25)}
                                            switch ("foo bar" + ("hello" + "\n\t")) {
                                            case "hello":
                                                if len(v6) > 0 {
                                                    v6[0] = (fi11 % ((fi11) | 1))
                                                }
                                            default:
                                                if true {
                                                    var v45 int = fi11
                                                    v46 := GenS5{F0: (fi11 & 9223372036854775807), F1: 1000, F2: true}
                                                    v7 += -76
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    v47 := GenS3{F0: 1000, F1: 42, F2: 4294967295}
                                                    var v48 uint8 = 3
                                                    _ = v47
                                                    _ = v48
                                                }
                                            }
                                            _ = v44
                                        }
                                        if ((!(65535 <= 4294967295)) || true) {
                                            if ((!true) || true) {
                                                var v49 uint32 = 945
                                                _ = v49
                                            }
                                            fmt.Println(fi11, fi11, v7)
                                        } else {
                                            v50 := GenS1{F0: 4294967295, F1: 688}
                                            var v51 uint16 = 65535
                                            if (!(v7 <= (v7 * (v7 % ((0) | 1))))) {
                                                var v52 int16 = -370
                                                _ = v52
                                            }
                                            _ = v50
                                            _ = v51
                                        }
                                    }
                                    if (!false) {
                                        fmt.Println(fi11, v7)
                                        if (!(65535 < 2147483647)) {
                                            var v53 uint16 = 853
                                            v54 := GenS4{F0: (v53 * 42), F1: (65535 > 373)}
                                            var v55 uint8 = 42
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        } else {
                                            v7 -= v7
                                            v7 = v7
                                            var v56 int8 = (v7 ^ (-(v7)))
                                            _ = v56
                                        }
                                        if true {
                                            v7 += (^(v7))
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        } else {
                                            var v57 int = fi29
                                            switch "foo bar" {
                                            case "a":
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            default:
                                                v6 = append(v6, (fi11 / (((fi29 | 9223372036854775807)) | 1)))
                                            }
                                            switch v57 {
                                            case 128:
                                                v7 -= (v7 % (((^(v7))) | 1))
                                            case 42:
                                                v7 -= v7
                                            case 2147483647:
                                                fmt.Println(v57, v57)
                                            }
                                            _ = v57
                                        }
                                    } else {
                                        v7 = v7
                                        v7 = (v7 % ((v7) | 1))
                                        if true {
                                            v58 := GenS3{F0: 1000, F1: 1, F2: 2147483648}
                                            switch (fi11 | fi29) {
                                            case 32768:
                                                if true {
                                                    if (v7 > (v7 | 112)) {
                                                        v6 = append(v6, (fi11 % ((32767) | 1)))
                                                        switch ("hello" + "") {
                                                        case "hello":
                                                            switch ("\n\t" + "hello") {
                                                            case "a":
                                                                if true {
                                                                    fmt.Println(fi11, v7)
                                                                    if ("" > ("foo bar" + "\n\t")) {
                                                                        v7 += v7
                                                                        switch (-(fi29)) {
                                                                        case 9223372036854775807:
                                                                            var v59 int16 = -32768
                                                                            v7 += (v7 | v7)
                                                                            if (4294967295 != 2) {
                                                                                switch ("foo bar" + ("a" + "hello")) {
                                                                                case "a":
                                                                                    v59 = (v59 & -319)
                                                                                case "":
                                                                                    v7 += v7
                                                                                }
                                                                                fmt.Println(v7, v7, fi29)
                                                                                var v60 int32 = 0
                                                                                _ = v60
                                                                            }
                                                                            _ = v59
                                                                        }
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            switch ("\n\t" + (("\n\t" + ((("hello" + "foo bar") + "a") + "hello")) + "hello")) {
                                                                            case "\n\t":
                                                                                var v61 uint8 = 127
                                                                                _ = v61
                                                                            }
                                                                        default:
                                                                            switch 9223372036854775807 {
                                                                            case -2147483648:
                                                                                v6 = append(v6, fi29)
                                                                            case 32767:
                                                                                if len(v6) >= 2 {
                                                                                    v6 = v6[1:len(v6)]
                                                                                }
                                                                            case 127:
                                                                                var v62 int64 = -682
                                                                                _ = v62
                                                                            }
                                                                        }
                                                                        _ = len(v6)
                                                                    } else {
                                                                        v7 -= (v7 + (v7 << ((42) & 63)))
                                                                        v7 = (v7 + 100)
                                                                    }
                                                                    v7 = (v7 << ((1) & 63))
                                                                    v63 := GenS7{F0: 239, F1: -314}
                                                                    var v64 int16 = -1
                                                                    var v65 int64 = -9223372036854775808
                                                                    if true {
                                                                        switch (fi11 * fi29) {
                                                                        case -562:
                                                                            v66 := GenS8{F0: 32767, F1: 4294967295}
                                                                            _ = v66
                                                                        case 4294967295:
                                                                            if (false || false) {
                                                                                v67 := GenS9{F0: 2147483647, F1: 2147483647}
                                                                                v65 = (v65 * (^(v65)))
                                                                                switch "\n\t" {
                                                                                case "hello":
                                                                                    v64 -= v64
                                                                                case "foo bar":
                                                                                    var v68 int8 = (v7 - 127)
                                                                                    _ = v68
                                                                                case "":
                                                                                    if false {
                                                                                        var v69 uint16 = 1
                                                                                        switch fi11 {
                                                                                        case -66:
                                                                                            _ = len(v6)
                                                                                            v70 := GenS10{F0: 877}
                                                                                            v64 = (-(v64))
                                                                                            _ = v70
                                                                                        case 4294967295:
                                                                                            var v71 uint = 613
                                                                                            _ = v71
                                                                                        default:
                                                                                            v72 := GenS11{F0: (v7 % (((-(v7))) | 1)), F1: "a", F2: "\n\t"}
                                                                                            _ = v72
                                                                                        }
                                                                                        if len(v6) >= 2 {
                                                                                            v6 = v6[1:len(v6)]
                                                                                        }
                                                                                        if (false && false) {
                                                                                            v6 = append(v6, (fi29 / ((-34) | 1)))
                                                                                            switch (((("foo bar" + "hello") + "\n\t") + "\n\t") + "\n\t") {
                                                                                            case "\n\t":
                                                                                                switch ("\n\t" + "\n\t") {
                                                                                                case "foo bar":
                                                                                                    switch "hello" {
                                                                                                    case "":
                                                                                                        var v73 uint32 = 83
                                                                                                        _ = v73
                                                                                                    }
                                                                                                case "":
                                                                                                    v74 := GenS12{F0: v69}
                                                                                                    _ = v74
                                                                                                case "hello":
                                                                                                    var v75 uint16 = (v69 + (v69 | (v69 & 65535)))
                                                                                                    _ = v75
                                                                                                default:
                                                                                                    v7 = (v7 * 0)
                                                                                                }
                                                                                            case "foo bar":
                                                                                                var v76 float32 = 0.1
                                                                                                _ = v76
                                                                                            }
                                                                                        } else {
                                                                                            v65 = v65
                                                                                            var v77 float32 = 1e10
                                                                                            _ = v77
                                                                                        }
                                                                                        v65 += v65
                                                                                        switch 944 {
                                                                                        case 128:
                                                                                            if (true || true) {
                                                                                                switch "\n\t" {
                                                                                                case "hello":
                                                                                                    var v78 float32 = 3.14159
                                                                                                    fmt.Println(v69)
                                                                                                    fmt.Println(v7, v78, fi11)
                                                                                                    _ = v78
                                                                                                case "a":
                                                                                                    var v79 string = ""
                                                                                                    _ = v79
                                                                                                }
                                                                                                _ = len(v6)
                                                                                            }
                                                                                        case 2147483647:
                                                                                            v80 := GenS13{F0: 889, F1: 42, F2: 42}
                                                                                            v81 := GenS14{F0: 65535}
                                                                                            _ = v80
                                                                                            _ = v81
                                                                                        case 32768:
                                                                                            if ((v65 != v65) || true) {
                                                                                                var v82 int16 = -1
                                                                                                var v83 bool = true
                                                                                                v6 = append(v6, fi11)
                                                                                                _ = v82
                                                                                                _ = v83
                                                                                            } else {
                                                                                                switch "\n\t" {
                                                                                                case "foo bar":
                                                                                                    v84 := GenS5{F0: (fi29 / (((-(fi29))) | 1)), F1: (v64 >> ((0) & 63)), F2: false}
                                                                                                    _ = v84
                                                                                                case "hello":
                                                                                                    fmt.Println(v69, fi29, v65)
                                                                                                default:
                                                                                                    v6 = append(v6, (fi11 / ((fi11) | 1)))
                                                                                                }
                                                                                                v65 = (v65 * 430)
                                                                                                switch fi11 {
                                                                                                case 9223372036854775807:
                                                                                                    _ = len(v6)
                                                                                                    v7 -= v7
                                                                                                    var v85 uint16 = (^(v69))
                                                                                                    var v86 int64 = v65
                                                                                                    _ = v85
                                                                                                    _ = v86
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v69
                                                                                    } else {
                                                                                        v64 = v64
                                                                                        fmt.Println(fi29, v65)
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(v7)
                                                                                }
                                                                                var v87 int8 = v7
                                                                                _ = v67
                                                                                _ = v87
                                                                            }
                                                                        case 127:
                                                                            var v88 int = fi11
                                                                            _ = v88
                                                                        default:
                                                                            v65 = (^(v65))
                                                                        }
                                                                        var v89 float32 = (1.5 - float32((0.0 * 1.5)))
                                                                        if false {
                                                                            if len(v6) > 0 {
                                                                                v6[0] = fi11
                                                                            }
                                                                            fmt.Println(v65, v64)
                                                                        }
                                                                        _ = v89
                                                                    } else {
                                                                        fmt.Println(v64, v7, fi29)
                                                                        v65 -= -9223372036854775808
                                                                        v90 := GenS13{F0: 561, F1: 40, F2: 4294967295}
                                                                        _ = v90
                                                                    }
                                                                    switch (fi11 / ((fi29) | 1)) {
                                                                    case 272:
                                                                        var v91 int16 = (v64 >> ((18446744073709551615) & 63))
                                                                        _ = v91
                                                                    }
                                                                    _ = v63
                                                                    _ = v64
                                                                    _ = v65
                                                                }
                                                            case "":
                                                                var v92 int = (fi29 % ((128) | 1))
                                                                _ = v92
                                                            case "foo bar":
                                                                if len(v6) > 0 {
                                                                    v6[0] = (fi29 / ((2147483647) | 1))
                                                                }
                                                            }
                                                        case "\n\t":
                                                            v7 += (v7 >> ((18446744073709551615) & 63))
                                                        case "":
                                                            if len(v6) > 0 {
                                                                v6[0] = -584
                                                            }
                                                        }
                                                    }
                                                }
                                            case 352:
                                                var v93 bool = false
                                                _ = v93
                                            case 918:
                                                switch (((("" + "") + "hello") + "hello") + ("a" + "")) {
                                                case "a":
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                case "\n\t":
                                                    var v94 float32 = (-2.25 / 1.0)
                                                    _ = v94
                                                }
                                            default:
                                                v95 := GenS13{F0: 1, F1: 2147483647, F2: 354}
                                                _ = v95
                                            }
                                            _ = v58
                                        }
                                    }
                                    v7 -= (v7 / ((1) | 1))
                                    v96 := GenS3{F0: -1, F1: 42, F2: 2}
                                    _ = v96
                                }
                            case 791:
                                v97 := GenS4{F0: 42, F1: (false && true)}
                                _ = v97
                            default:
                                var v98 uint16 = 1
                                _ = v98
                            }
                        case "\n\t":
                            _ = len(v6)
                        default:
                            v7 += 100
                        }
                        _ = fi29
                    }
                    _ = fi11
                }
                for fi99 := 0; fi99 < 5; fi99++ {
                    if (v7 == v7) {
                        var v100 uint8 = 71
                        v100 -= (v100 % (((v100 / ((v100) | 1))) | 1))
                        _ = v100
                    } else {
                        for fi101 := 0; fi101 < 1; fi101++ {
                            var v102 float64 = 0.1
                            v102 = ((v102 + v102) + v102)
                            _ = fi101
                            _ = v102
                        }
                        v103 := GenS12{F0: 0}
                        v7 += (v7 >> ((42) & 63))
                        _ = v103
                    }
                    _ = fi99
                }
            }
        }
        _ = v7
    } else {
        var v104 int8 = 0
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v104
    }
    v105 := GenS12{F0: 65535}
    if ((false && true) || false) {
        v106 := GenS3{F0: 1, F1: 0, F2: 0}
        _ = v106
    } else {
        if false {
            if (!(260 == 2)) {
                var v107 uint32 = 2
                v108 := GenS14{F0: 42}
                v107 -= (^(v107))
                _ = v107
                _ = v108
            }
        } else {
            if ("a" < ("\n\t" + "a")) {
                if false {
                    if false {
                        switch -930 {
                        case 128:
                            var v109 uint32 = 78
                            _ = v109
                        default:
                            if true {
                                var v110 float64 = float64(0.0)
                                var v111 uint = 255
                                var v112 string = ("foo bar" + "foo bar")
                                _ = v110
                                _ = v111
                                _ = v112
                            } else {
                                var v113 int64 = 42
                                if (1000 < -32768) {
                                    var v114 string = ((("\n\t" + ("a" + "foo bar")) + (("a" + "a") + "hello")) + "a")
                                    var v115 int16 = 0
                                    var v116 uint16 = 0
                                    _ = v114
                                    _ = v115
                                    _ = v116
                                }
                                _ = v113
                            }
                        }
                        var v117 int16 = -995
                        fmt.Println(v117)
                        _ = v117
                    } else {
                        _ = len(v6)
                    }
                    for fi118 := 0; fi118 < 4; fi118++ {
                        if true {
                            for fi119 := 0; fi119 < 13; fi119++ {
                                fmt.Println(fi119, fi118)
                                v120 := GenS4{F0: 0, F1: true}
                                _ = fi119
                                _ = v120
                            }
                            v121 := GenS14{F0: 362}
                            var v122 int64 = 9223372036854775807
                            _ = v121
                            _ = v122
                        } else {
                            v123 := GenS14{F0: 42}
                            var v124 uint = 222
                            _ = v123
                            _ = v124
                        }
                        for fi125 := 0; fi125 < 11; fi125++ {
                            var v126 uint32 = 65535
                            if false {
                                var v127 int32 = 82
                                var v128 uint64 = 178
                                switch (fi125 | (fi118 << ((42) & 63))) {
                                case 32768:
                                    switch (fi118 - (fi118 & -63)) {
                                    case 725:
                                        switch ("foo bar" + "\n\t") {
                                        case "foo bar":
                                            var v129 int8 = 100
                                            _ = v129
                                        default:
                                            switch "hello" {
                                            case "":
                                                v130 := GenS12{F0: 32767}
                                                _ = v130
                                            case "\n\t":
                                                v126 = v126
                                            case "a":
                                                v131 := GenS15{F0: v126, F1: v126}
                                                _ = v131
                                            default:
                                                var v132 int32 = (v127 / (((^(v127))) | 1))
                                                _ = v132
                                            }
                                        }
                                    case 42:
                                        if (false && (("\n\t" == "foo bar") && false)) {
                                            if false {
                                                _ = len(v6)
                                                v128 += (v128 / ((v128) | 1))
                                            } else {
                                                var v133 float32 = 0.1
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                _ = v133
                                            }
                                            switch -262 {
                                            case 0:
                                                switch "" {
                                                case "":
                                                    fmt.Println(v126)
                                                case "\n\t":
                                                    v134 := GenS2{F0: 127, F1: 42}
                                                    _ = v134
                                                default:
                                                    var v135 uint16 = 1
                                                    _ = v135
                                                }
                                            case 174:
                                                if (0 <= -1) {
                                                    v127 = (^(v127))
                                                    v127 += 1
                                                    v128 += 1
                                                    v136 := GenS16{F0: int(v128)}
                                                    _ = v136
                                                }
                                            case -758:
                                                var v137 string = "foo bar"
                                                _ = v137
                                            }
                                            var v138 uint32 = 2147483648
                                            _ = v138
                                        } else {
                                            var v139 uint64 = v128
                                            _ = len(v6)
                                            _ = v139
                                        }
                                    case -143:
                                        _ = len(v6)
                                    default:
                                        if true {
                                            v140 := GenS16{F0: fi118}
                                            switch "" {
                                            case "\n\t":
                                                if (-477 > 1) {
                                                    if (uint32(v128) < 304) {
                                                        v141 := GenS8{F0: 2147483647, F1: (v126 << ((1) & 63))}
                                                        switch fi118 {
                                                        case 4294967295:
                                                            fmt.Println(fi125, fi125, v127)
                                                        }
                                                        if ("foo bar" >= "\n\t") {
                                                            fmt.Println(fi125)
                                                        }
                                                        _ = v141
                                                    } else {
                                                        if len(v6) > 0 {
                                                            v6[0] = fi118
                                                        }
                                                    }
                                                }
                                            case "foo bar":
                                                if true {
                                                    fmt.Println(fi125, v127, fi125)
                                                    var v142 uint64 = v128
                                                    if (1000 == 32767) {
                                                        var v143 float32 = 0.0
                                                        if (170 < 18) {
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            v144 := GenS15{F0: v126, F1: v126}
                                                            switch (fi125 + (fi125 ^ 971)) {
                                                            case 1:
                                                                var v145 bool = false
                                                                v6 = append(v6, (fi118 ^ 1))
                                                                v127 = (-(v127))
                                                                v126 = (v126 << ((4294967295) & 63))
                                                                _ = v145
                                                            case 2147483647:
                                                                v143 -= ((v143 - v143) + v143)
                                                            }
                                                            if false {
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                v126 = v126
                                                                fmt.Println(v126, v126)
                                                                var v146 float64 = 1.5
                                                                _ = v146
                                                            } else {
                                                                if false {
                                                                    v147 := GenS8{F0: (fi125 - (fi118 % ((-1) | 1))), F1: 0}
                                                                    _ = v147
                                                                }
                                                                if len(v6) > 0 {
                                                                    v6[0] = -2147483648
                                                                }
                                                            }
                                                            v6 = append(v6, (fi125 >> ((4294967295) & 63)))
                                                            _ = v144
                                                        }
                                                        _ = v143
                                                    }
                                                    _ = v142
                                                } else {
                                                    v148 := GenS17{F0: 118, F1: 985}
                                                    if true {
                                                        v149 := GenS18{F0: 32767, F1: (fi125 | 1), F2: "hello"}
                                                        v126 -= 293
                                                        _ = v149
                                                    }
                                                    _ = len(v6)
                                                    var v150 uint64 = (^(v128))
                                                    if (-850 > -1) {
                                                        v128 = 4294967295
                                                    } else {
                                                        switch fi125 {
                                                        case 128:
                                                            var v151 uint32 = (v126 | (v126 & (v126 * (v126 + (v126 - 2)))))
                                                            _ = v151
                                                        case -89:
                                                            if (!(true || true)) {
                                                                v152 := GenS13{F0: 18446744073709551615, F1: v126, F2: v150}
                                                                var v153 int64 = 0
                                                                v126 = (v126 << ((42) & 63))
                                                                _ = v152
                                                                _ = v153
                                                            }
                                                        default:
                                                            var v154 uint16 = 2
                                                            _ = v154
                                                        }
                                                    }
                                                    _ = v148
                                                    _ = v150
                                                }
                                                _ = len(v6)
                                            }
                                            _ = v140
                                        }
                                    }
                                case -9223372036854775808:
                                    v126 -= v126
                                case -45:
                                    fmt.Println(fi118, v127)
                                }
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if len(v6) > 0 {
                                    v6[0] = fi125
                                }
                                _ = v127
                                _ = v128
                            }
                            _ = fi125
                            _ = v126
                        }
                        fmt.Println(fi118)
                        _ = fi118
                    }
                    v155 := GenS19{F0: false, F1: -11}
                    _ = v155
                } else {
                    var v156 bool = false
                    switch -151 {
                    case 32768:
                        var v157 int8 = 70
                        _ = v157
                    case 9223372036854775807:
                        fmt.Println(v156)
                    case 2147483647:
                        var v158 int64 = 224
                        _ = v158
                    }
                    _ = v156
                }
            } else {
                if false {
                    v159 := GenS9{F0: 1, F1: 2}
                    _ = v159
                } else {
                    var v160 int8 = 100
                    _ = v160
                }
            }
        }
    }
    v161 := GenS10{F0: 4294967295}
    _ = v105
    _ = v161
    _ = v6
    return 4294967295
}

func main() {
    v1 := []uint32{0, 2}
    var v2 int = 113
    v3 := []uint32{571, 2}
    v4 := []int{v2, v2, (v2 ^ (v2 - 4294967295)), v2}
    v5 := make(map[string]string)
    var v162 float32 = 0.1
    v2 = (v2 % ((v2) | 1))
    for fi163 := 0; fi163 < 3; fi163++ {
        v2 += fi163
        v2 += (v2 + genFn0())
        var v164 uint64 = 18446744073709551615
        _ = fi163
        _ = v164
    }
    var v165 int = genFn0()
    var v166 string = ("" + ("" + ("a" + "")))
    fmt.Println(v162, v162, v162)
    var v167 int64 = -105
    var v168 uint = 65535
    v167 += v167
    v169 := GenS20{F0: 257, F1: "foo bar", F2: (v167 - 42)}
    if (!(((-128 > -116) && (true && (117 < 47))) || (509 != (^(v2))))) {
        v170 := GenS18{F0: (v2 * (v165 & v2)), F1: (v165 / ((32767) | 1)), F2: ("" + ((v166 + v166) + v166))}
        fmt.Println(v162, v167, v168)
        _ = v170
    } else {
        if (false || true) {
            v171 := GenS10{F0: 2}
            switch genFn0() {
            case 4294967295:
                switch (v165 ^ v2) {
                case 2147483647:
                    if (1 != 0) {
                        for fi172 := 0; fi172 < 7; fi172++ {
                            var v173 int64 = (v167 << ((v168) & 63))
                            switch v166 {
                            case "a":
                                fmt.Println(v167)
                            case "\n\t":
                                v174 := GenS6{F0: 79, F1: -339, F2: int32(v168)}
                                _ = v174
                            case "hello":
                                v175 := GenS5{F0: fi172, F1: 1000, F2: (!(false && (v162 != 0.0)))}
                                _ = v175
                            }
                            if (false || false) {
                                for fi176 := 0; fi176 < 6; fi176++ {
                                    var v177 uint64 = 42
                                    var v178 float64 = (0.1 + (((float64(0.1) + (float64((v162 - (v162 - 0.1))) * (-2.25 / (1e10 + 0.1)))) + 3.14159) - (0.1 + 1.5)))
                                    _ = fi176
                                    _ = v177
                                    _ = v178
                                }
                            }
                            _ = fi172
                            _ = v173
                        }
                    }
                case 127:
                    for fi179 := 0; fi179 < 16; fi179++ {
                        switch int(v167) {
                        case 128:
                            switch genFn0() {
                            case 4294967295:
                                var v180 string = v166
                                _ = v180
                            }
                        default:
                            v181 := GenS13{F0: v168, F1: 502, F2: 67}
                            _ = v181
                        }
                        _ = fi179
                    }
                case 128:
                    _ = len(v1)
                }
            case 674:
                var v182 float32 = ((((float32((0.1 - float64(v162))) / v162) / 1.0) - (v162 * 3.14159)) - v162)
                _ = v182
            case -924:
                if true {
                    for fi183 := 0; fi183 < 16; fi183++ {
                        for fi184 := 0; fi184 < 11; fi184++ {
                            if true {
                                v5["hello"] = v166
                                v185 := GenS13{F0: 4294967295, F1: 4294967295, F2: 612}
                                _ = v185
                            }
                            fmt.Println(v162, v168)
                            var v186 uint = v168
                            _ = fi184
                            _ = v186
                        }
                        fmt.Println(v168)
                        for fi187 := 0; fi187 < 19; fi187++ {
                            if ((3.14159 - float32((0.0 / -2.25))) >= ((float32(1.5) * v162) - (1.5 + ((v162 / (float32(1.5) + (v162 + 0.1))) / v162)))) {
                                v167 = (v167 & v167)
                            } else {
                                fmt.Println(v168)
                            }
                            if false {
                                var v188 int32 = -2147483648
                                v167 += (v167 / (((v167 * (v167 / ((v167) | 1)))) | 1))
                                if (1 != 1) {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                v189 := GenS15{F0: 2147483648, F1: 42}
                                _ = v188
                                _ = v189
                            } else {
                                var v190 uint8 = 0
                                if false {
                                    v5[""] = v166
                                    if true {
                                        fmt.Println(v2)
                                        v191 := GenS2{F0: 40, F1: v167}
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v191
                                    }
                                    v192 := v5[v166]
                                    _ = v192
                                } else {
                                    v165 += -349
                                    v165 += genFn0()
                                    var v193 uint32 = 461
                                    _ = v193
                                }
                                fmt.Println(v2, v167)
                                if true {
                                    v2 -= -2147483648
                                    fmt.Println(fi183, v190, v167)
                                    if ((0.0 <= (1.5 - 3.14159)) || false) {
                                        v194 := GenS18{F0: (fi183 & (v165 << (((v168 - (^(v168)))) & 63))), F1: v165, F2: ("" + "a")}
                                        var v195 uint = (v168 ^ v168)
                                        v2 -= fi187
                                        v196 := v5[(v166 + (("hello" + v166) + v166))]
                                        v197 := GenS18{F0: fi183, F1: -9223372036854775808, F2: v166}
                                        v4 = append(v4, genFn0())
                                        v198 := GenS14{F0: 127}
                                        _ = v194
                                        _ = v195
                                        _ = v196
                                        _ = v197
                                        _ = v198
                                    }
                                }
                                delete(v5, ((v166 + v166) + "a"))
                                _ = v190
                            }
                            switch "foo bar" {
                            case "\n\t":
                                var v199 uint64 = 0
                                switch v166 {
                                case "\n\t":
                                    switch -472 {
                                    case 108:
                                        fmt.Println(v168)
                                        fmt.Println(v168, v166)
                                    default:
                                        switch v165 {
                                        case 127:
                                            var v200 int = fi183
                                            _ = v200
                                        case -617:
                                            switch -96 {
                                            case 42:
                                                fmt.Println(v162)
                                            case -1:
                                                var v201 int16 = -1
                                                v202 := GenS9{F0: 2, F1: 4294967295}
                                                v162 += 3.14159
                                                _ = v201
                                                _ = v202
                                            }
                                        case -998:
                                            v199 = (v199 + (v199 | (v199 + v199)))
                                        }
                                    }
                                case "foo bar":
                                    v165 -= v165
                                case "":
                                    v203 := GenS21{F0: uint32(v199), F1: 1, F2: 779}
                                    _ = v203
                                }
                                _ = v199
                            case "foo bar":
                                fmt.Println(v166)
                            case "a":
                                if false {
                                    fmt.Println(v165, v162)
                                    v204 := GenS9{F0: 729, F1: 2147483648}
                                    _ = v204
                                } else {
                                    switch (v166 + v166) {
                                    case "\n\t":
                                        v165 = (fi187 | -115)
                                        v2 = (v2 / ((127) | 1))
                                        var v205 float32 = v162
                                        _ = v205
                                    case "hello":
                                        fmt.Println(fi187, v168)
                                    }
                                    fmt.Println(v165, v162)
                                    v206 := v5["\n\t"]
                                    _ = v206
                                }
                            default:
                                if (715 != 2147483648) {
                                    fmt.Println(v165)
                                    var v207 int8 = 100
                                    _ = v207
                                }
                            }
                            _ = fi187
                        }
                        _ = fi183
                    }
                }
            }
            delete(v5, v166)
            _ = v171
        } else {
            v168 = v168
            v167 += v167
        }
        fmt.Println(v168, v2, v165)
        switch "a" {
        case "hello":
            for fi208 := 0; fi208 < 10; fi208++ {
                for fi209 := 0; fi209 < 6; fi209++ {
                    if true {
                        v210 := GenS14{F0: 42}
                        fmt.Println(v168, fi208)
                        var v211 uint16 = 377
                        var v212 string = ("foo bar" + (v166 + "\n\t"))
                        v212 = "a"
                        _ = v210
                        _ = v211
                        _ = v212
                    }
                    _ = fi209
                }
                if false {
                    v213 := GenS22{F0: v168}
                    _ = v213
                }
                for fi214 := 0; fi214 < 3; fi214++ {
                    switch v166 {
                    case "hello":
                        switch fi208 {
                        case -9223372036854775808:
                            v165 = 622
                        case 0:
                            if (!false) {
                                switch int(v167) {
                                case 32768:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case 42:
                                    v215 := v5[""]
                                    _ = v215
                                }
                                if (int16(v167) != -32768) {
                                    var v216 uint8 = 144
                                    _ = v216
                                }
                            } else {
                                if true {
                                    var v217 string = "hello"
                                    _ = v217
                                } else {
                                    if false {
                                        var v218 int16 = 32767
                                        _ = v218
                                    } else {
                                        var v219 float64 = (0.0 / 3.14159)
                                        _ = v219
                                    }
                                }
                            }
                        default:
                            v2 = v2
                        }
                    default:
                        var v220 uint16 = 1
                        _ = v220
                    }
                    _ = fi214
                }
                _ = fi208
            }
        case "foo bar":
            if ((true && (32767 == 255)) || true) {
                for fi221 := 0; fi221 < 6; fi221++ {
                    if (!(v165 > v165)) {
                        for fi222 := 0; fi222 < 2; fi222++ {
                            v162 = (0.1 * (v162 / (1e10 / (0.0 / (((((v162 + v162) - v162) + v162) * 3.14159) - 0.1)))))
                            v165 += 32768
                            _ = fi222
                        }
                    } else {
                        v165 = v2
                        var v223 string = ("a" + ("foo bar" + ("hello" + v166)))
                        _ = v223
                    }
                    var v224 float64 = 0.1
                    _ = fi221
                    _ = v224
                }
                var v225 float32 = v162
                var v226 int = -2147483648
                _ = v225
                _ = v226
            } else {
                v227 := GenS23{F0: 1}
                var v228 uint8 = 127
                _ = v227
                _ = v228
            }
        case "":
            switch genFn0() {
            case -2147483648:
                if true {
                    switch ("hello" + v166) {
                    case "\n\t":
                        for fi229 := 0; fi229 < 11; fi229++ {
                            var v230 int = 32768
                            v165 += v230
                            _ = fi229
                            _ = v230
                        }
                    case "foo bar":
                        for fi231 := 0; fi231 < 10; fi231++ {
                            var v232 int = v165
                            _ = fi231
                            _ = v232
                        }
                    }
                    v233 := GenS24{F0: v166, F1: 0}
                    v234 := GenS20{F0: 1, F1: ("" + (((v166 + v166) + "\n\t") + v166)), F2: (v167 % ((352) | 1))}
                    _ = v233
                    _ = v234
                }
            case 1:
                if true {
                    v235 := GenS20{F0: 754, F1: ("foo bar" + v166), F2: v167}
                    _ = v235
                } else {
                    var v236 string = v166
                    _ = v236
                }
            default:
                fmt.Println(v167, v162)
            }
        }
    }
    if ((!(!true)) && false) {
        var v237 bool = ((((v162 - (v162 + (3.14159 - v162))) / v162) * v162) < 3.14159)
        v238 := GenS12{F0: 32767}
        _ = v237
        _ = v238
    }
    v165 = 2147483647
    v167 = -326
    if (!(((0.0 - ((1.5 / (((0.1 - (1e10 - 1e10)) / (((1.5 + 1.5) * 3.14159) - float64(1e10))) * 1.5)) * 0.0)) / (1e10 + 0.0)) > -2.25)) {
        for fi239 := 0; fi239 < 19; fi239++ {
            v240 := v5[v166]
            for fi241 := 0; fi241 < 18; fi241++ {
                if false {
                    v166 = "\n\t"
                } else {
                    switch v166 {
                    case "":
                        v167 += 744
                    case "a":
                        if len(v3) > 0 {
                            v3[0] = 2
                        }
                    case "hello":
                        v165 = (fi239 + (fi241 / (((^(v2))) | 1)))
                    default:
                        var v242 float64 = 3.14159
                        _ = v242
                    }
                    v167 = (v167 | v167)
                }
                switch (v166 + ("foo bar" + v240)) {
                case "":
                    v165 += (v2 / ((genFn0()) | 1))
                case "a":
                    var v243 float32 = 0.0
                    v244 := GenS1{F0: 4294967295, F1: 799}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v243
                    _ = v244
                case "foo bar":
                    switch v2 {
                    case 32768:
                        switch (fi241 << ((v168) & 63)) {
                        case 4294967295:
                            if false {
                                v2 -= (fi239 % (((v2 << ((v168) & 63))) | 1))
                            }
                        default:
                            v245 := v5[v240]
                            _ = v245
                        }
                    case -1:
                        v166 = v166
                    default:
                        var v246 uint = v168
                        _ = v246
                    }
                }
                var v247 int16 = 0
                if (int16(v168) >= -991) {
                    v167 -= (v167 + v167)
                    v162 -= 1e10
                    v248 := v5[v240]
                    _ = v248
                } else {
                    fmt.Println(v162)
                    _ = len(v5)
                }
                v249 := GenS21{F0: 0, F1: 32767, F2: genFn0()}
                _ = fi241
                _ = v247
                _ = v249
            }
            _ = fi239
            _ = v240
        }
    }
    if (v166 != v166) {
        var v250 int = genFn0()
        var v251 int8 = -1
        for fi252 := 0; fi252 < 15; fi252++ {
            v251 = v251
            _ = fi252
        }
        _ = v250
        _ = v251
    } else {
        if false {
            v1 = append(v1, 2147483648)
            delete(v5, "")
        }
    }
    var v253 int8 = -5
    var v254 uint32 = 464
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v162:", v162)
    fmt.Println("v165:", v165)
    fmt.Println("v166:", v166)
    fmt.Println("v167:", v167)
    fmt.Println("v168:", v168)
    fmt.Println("v2:", v2)
    fmt.Println("v253:", v253)
    fmt.Println("v254:", v254)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v162
    _ = v165
    _ = v166
    _ = v167
    _ = v168
    _ = v169
    _ = v2
    _ = v253
    _ = v254
    _ = v3
    _ = v4
    _ = v5
}
