package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 string; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int64; F1 int; F2 uint }
type GenS5 struct { F0 int; F1 uint; F2 bool }
type GenS6 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS7 struct { F0 int8; F1 int16; F2 string }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint; F1 uint8; F2 bool }
type GenS10 struct { F0 uint32; F1 int8; F2 uint }
type GenS11 struct { F0 string; F1 uint8; F2 uint8 }
type GenS12 struct { F0 uint16; F1 int32; F2 bool }
type GenS13 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS14 struct { F0 int64; F1 uint }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 string; F1 string }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int32; F1 int }
func genFn0() uint8 {
    var v6 bool = true
    var v7 uint = 787
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint8 = 80
    var v2 float32 = 0.0
    var v3 float64 = (3.14159 / (-2.25 / 3.14159))
    var v4 float32 = v2
    v5 := make(map[uint32]uint32)
    fmt.Println(v2)
    var v8 uint64 = 894
    switch ("\n\t" + "hello") {
    case "\n\t":
        fmt.Println(v1, v3, v1)
    default:
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 uint32 = uint32(v1)
            var v11 uint = 4294967295
            v12 := GenS1{F0: 1, F1: "foo bar"}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        if (-1 != 42) {
            for fi14 := 0; fi14 < 16; fi14++ {
                v4 = 0.0
                var v15 float32 = v2
                var v16 string = ("hello" + "")
                _ = fi14
                _ = v15
                _ = v16
            }
            v4 += v2
        } else {
            v17 := []int{fi13, 126, int(v8)}
            v2 = v4
            _ = v17
        }
        var v18 int8 = -1
        v4 = (float32(0.1) - 1e10)
        _ = fi13
        _ = v18
    }
    v3 = v3
    fmt.Println(v3, v4, v4)
    var v19 float32 = (v2 * 0.0)
    var v20 uint16 = 1
    var v21 uint32 = 65535
    if ((("foo bar" + "hello") + "hello") != "") {
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := GenS2{F0: "foo bar", F1: (^(v8)), F2: genFn0()}
            v19 -= -2.25
            fmt.Println(fi22, v20, v19)
            _ = fi22
            _ = v23
        }
        v24 := GenS2{F0: "", F1: v8, F2: (v1 & (^(v1)))}
        _ = v24
    } else {
        var v25 uint16 = v20
        fmt.Println(v21, v21, v19)
        _ = v25
    }
    if (!true) {
        if false {
            for fi26 := 0; fi26 < 9; fi26++ {
                v2 += 1.5
                v27 := []int{fi26}
                for fi28 := 0; fi28 < 20; fi28++ {
                    if true {
                        if ("\n\t" < ("foo bar" + "a")) {
                            switch 67 {
                            case 186:
                                _ = len(v5)
                            default:
                                v29 := GenS2{F0: (("foo bar" + "") + "\n\t"), F1: (v8 / (((v8 | v8)) | 1)), F2: 11}
                                _ = v29
                            }
                            switch "\n\t" {
                            case "hello":
                                fmt.Println(v1, v19)
                            case "a":
                                if len(v27) >= 2 {
                                    v27 = v27[1:len(v27)]
                                }
                            case "foo bar":
                                v30 := GenS3{F0: 0, F1: ("hello" + "")}
                                _ = v30
                            }
                            fmt.Println(fi28)
                        } else {
                            _ = len(v5)
                            if len(v27) > 0 {
                                v27[0] = (fi28 * fi26)
                            }
                        }
                        if (1 < (v21 / ((v21) | 1))) {
                            v31 := GenS2{F0: "", F1: v8, F2: v1}
                            switch fi28 {
                            case 32767:
                                v4 = ((v4 + (1.5 - 0.1)) * v19)
                            case 241:
                                v32 := GenS4{F0: 9223372036854775807, F1: (fi26 ^ -712), F2: 0}
                                _ = v32
                            }
                            v8 -= v8
                            v8 -= 274
                            _ = v31
                        } else {
                            v21 = (v21 - (v21 * uint32(v1)))
                            v4 = -2.25
                            var v33 int32 = -7
                            _ = v33
                        }
                        v1 = genFn0()
                        switch fi26 {
                        case 128:
                            switch (fi28 ^ fi26) {
                            case 155:
                                if false {
                                    if len(v27) > 0 {
                                        v27[0] = fi26
                                    }
                                    var v34 uint16 = 42
                                    switch (^(fi28)) {
                                    case 9223372036854775807:
                                        delete(v5, v21)
                                        v21 = uint32(v1)
                                    case 32768:
                                        switch "a" {
                                        case "a":
                                            v4 = (0.1 * -2.25)
                                        default:
                                            v35 := v5[675]
                                            _ = v35
                                        }
                                    default:
                                        v1 = 226
                                    }
                                    switch fi28 {
                                    case 642:
                                        var v36 int32 = -754
                                        _ = v36
                                    case 139:
                                        var v37 uint64 = 182
                                        _ = v37
                                    }
                                    _ = v34
                                }
                            case 0:
                                if false {
                                    var v38 int8 = -124
                                    if (v1 <= (^(v1))) {
                                        v3 += (1e10 + (v3 * ((((v3 / v3) / v3) + v3) + v3)))
                                        switch 4294967295 {
                                        case 4294967295:
                                            switch (("a" + "hello") + "a") {
                                            case "":
                                                fmt.Println(v19, fi26)
                                                switch -88 {
                                                case -42:
                                                    v39 := GenS2{F0: "a", F1: (v8 / (((v8 << ((0) & 63))) | 1)), F2: v1}
                                                    v8 = (v8 * (v8 & 1))
                                                    _ = v39
                                                }
                                                v40 := GenS5{F0: fi28, F1: 2, F2: (v8 != v8)}
                                                _ = v40
                                            case "foo bar":
                                                var v41 uint8 = uint8(v21)
                                                _ = v41
                                            case "hello":
                                                v4 -= 1.5
                                            }
                                        case -525:
                                            var v42 int8 = (v38 + -128)
                                            _ = v42
                                        default:
                                            switch fi28 {
                                            case 621:
                                                if false {
                                                    if true {
                                                        var v43 uint32 = (v21 | (v21 << ((652) & 63)))
                                                        v3 += v3
                                                        v8 = (v8 % ((uint64(v20)) | 1))
                                                        var v44 int16 = 469
                                                        if (!((fi26 << ((0) & 63)) >= fi28)) {
                                                            v3 = (v3 - (v3 + (v3 * (v3 + v3))))
                                                            v45 := GenS6{F0: (v8 & (^(v8))), F1: v8, F2: (v44 | 0)}
                                                            var v46 bool = true
                                                            v47 := GenS1{F0: -995, F1: ""}
                                                            _ = v45
                                                            _ = v46
                                                            _ = v47
                                                        } else {
                                                            var v48 int = fi26
                                                            _ = v48
                                                        }
                                                        fmt.Println(fi28, v43, v2)
                                                        _ = v43
                                                        _ = v44
                                                    } else {
                                                        if ((!(!(1 >= 42))) && ((v38 * 56) != -1)) {
                                                            v38 = v38
                                                            if (!false) {
                                                                if (!(fi28 <= fi28)) {
                                                                    switch ((("hello" + "a") + "hello") + "\n\t") {
                                                                    case "hello":
                                                                        v21 = v21
                                                                    case "":
                                                                        v1 += 255
                                                                    case "foo bar":
                                                                        _ = len(v27)
                                                                    }
                                                                } else {
                                                                    if len(v27) >= 2 {
                                                                        v27 = v27[1:len(v27)]
                                                                    }
                                                                }
                                                            } else {
                                                                var v49 bool = (!false)
                                                                switch "hello" {
                                                                case "":
                                                                    _ = len(v5)
                                                                }
                                                                v19 = 0.1
                                                                _ = v49
                                                            }
                                                        }
                                                        var v50 uint32 = (v21 % (((v21 - v21)) | 1))
                                                        _ = v50
                                                    }
                                                    v51 := GenS1{F0: -833, F1: ("\n\t" + "")}
                                                    if len(v27) >= 2 {
                                                        v27 = v27[1:len(v27)]
                                                    }
                                                    v21 = v21
                                                    _ = v51
                                                }
                                            case 4294967295:
                                                v4 = v2
                                            case 9223372036854775807:
                                                switch fi28 {
                                                case -2147483648:
                                                    if false {
                                                        v52 := GenS4{F0: 1, F1: 523, F2: 154}
                                                        var v53 int64 = -442
                                                        v54 := v5[(v21 * v21)]
                                                        _ = v52
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                case 9223372036854775807:
                                                    var v55 int32 = -2147483648
                                                    _ = v55
                                                case -842:
                                                    v21 += 4294967295
                                                }
                                            }
                                        }
                                    } else {
                                        var v56 int32 = 2147483647
                                        _ = v56
                                    }
                                    switch (-(fi28)) {
                                    case 1:
                                        fmt.Println(v19, v2)
                                    case 4294967295:
                                        var v57 int64 = -686
                                        _ = v57
                                    case 42:
                                        _ = len(v27)
                                    }
                                    _ = v38
                                } else {
                                    v58 := GenS3{F0: 42, F1: (("hello" + ("a" + "foo bar")) + ("foo bar" + ("" + "")))}
                                    _ = v58
                                }
                            case 127:
                                v27 = append(v27, (fi28 << ((675) & 63)))
                            }
                        case 42:
                            fmt.Println(fi28)
                        default:
                            v20 -= (v20 + 42)
                        }
                        var v59 bool = (-2.25 > v2)
                        _ = v59
                    }
                    _ = len(v27)
                    if (false || (false && true)) {
                        var v60 float32 = v4
                        if (true && false) {
                            fmt.Println(v3)
                            v61 := GenS1{F0: -1, F1: ((("a" + "a") + (("a" + "\n\t") + "a")) + ("a" + ("" + "hello")))}
                            _ = v61
                        } else {
                            if ((^(v21)) < 2147483647) {
                                v19 -= 0.0
                                switch int(v1) {
                                case -9223372036854775808:
                                    fmt.Println(v2, fi26)
                                case 997:
                                    var v62 uint16 = v20
                                    _ = v62
                                default:
                                    switch ("foo bar" + "foo bar") {
                                    case "hello":
                                        v21 += 2
                                    case "\n\t":
                                        var v63 uint32 = (v21 << ((0) & 63))
                                        _ = v63
                                    }
                                }
                                v5[796] = v21
                                var v64 uint16 = v20
                                _ = v64
                            } else {
                                fmt.Println(v60)
                            }
                            v4 = v19
                        }
                        _ = v60
                    }
                    _ = fi28
                }
                _ = fi26
                _ = v27
            }
            fmt.Println(v21, v2, v1)
        } else {
            if false {
                fmt.Println(v19)
                switch ("" + "a") {
                case "a":
                    v65 := v5[v21]
                    _ = v65
                case "\n\t":
                    var v66 int8 = int8(v8)
                    _ = v66
                case "foo bar":
                    if true {
                        if ((-34 <= 127) && true) {
                            fmt.Println(v19, v3, v4)
                            v67 := GenS6{F0: v8, F1: v8, F2: 32767}
                            if (v20 == (v20 + (v20 * (v20 >> ((65535) & 63))))) {
                                v68 := []int{42, 4294967295}
                                var v69 uint8 = genFn0()
                                _ = v68
                                _ = v69
                            } else {
                                var v70 float32 = v2
                                for fi71 := 0; fi71 < 15; fi71++ {
                                    var v72 uint64 = (v8 ^ v8)
                                    v73 := GenS7{F0: -36, F1: 736, F2: ("a" + "")}
                                    _ = fi71
                                    _ = v72
                                    _ = v73
                                }
                                if true {
                                    switch "a" {
                                    case "a":
                                        for fi74 := 0; fi74 < 1; fi74++ {
                                            if (!true) {
                                                var v75 uint8 = (v1 % ((v1) | 1))
                                                _ = v75
                                            } else {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v19 = float32(0.1)
                                                case "hello":
                                                    v76 := []int{}
                                                    _ = v76
                                                default:
                                                    v70 = ((v2 / float32(-2.25)) / ((-2.25 + (v70 + v2)) / v70))
                                                }
                                                v1 = uint8(v8)
                                                v19 = -2.25
                                            }
                                            _ = fi74
                                        }
                                    default:
                                        v77 := GenS3{F0: 42, F1: ("\n\t" + "hello")}
                                        _ = v77
                                    }
                                } else {
                                    var v78 string = "\n\t"
                                    for fi79 := 0; fi79 < 10; fi79++ {
                                        v21 -= (^(v21))
                                        v1 += 2
                                        _ = fi79
                                    }
                                    var v80 uint32 = uint32(v20)
                                    _ = v78
                                    _ = v80
                                }
                                _ = v70
                            }
                            _ = v67
                        } else {
                            var v81 int16 = 0
                            _ = v81
                        }
                        var v82 float64 = v3
                        _ = v82
                    }
                }
                v83 := GenS1{F0: 42, F1: "a"}
                _ = v83
            }
        }
        for fi84 := 0; fi84 < 4; fi84++ {
            v85 := v5[uint32(v20)]
            if (!(v3 != v3)) {
                switch ("\n\t" + "foo bar") {
                case "":
                    if true {
                        fmt.Println(v85)
                    } else {
                        if true {
                            for fi86 := 0; fi86 < 14; fi86++ {
                                v5[v85] = v85
                                v19 = ((v2 / 0.1) - float32((((v3 + float64(0.1)) * v3) + v3)))
                                var v87 int = (fi84 / (((fi84 >> ((2) & 63))) | 1))
                                _ = fi86
                                _ = v87
                            }
                            for fi88 := 0; fi88 < 1; fi88++ {
                                fmt.Println(fi88)
                                if (42 > 504) {
                                    v89 := GenS8{F0: (-32768 <= -500)}
                                    v8 -= (v8 | v8)
                                    switch ("" + ("hello" + ("foo bar" + ("" + "")))) {
                                    case "a":
                                        var v90 bool = (255 < 0)
                                        _ = v90
                                    case "foo bar":
                                        v91 := []int{int(v20)}
                                        _ = v91
                                    default:
                                        delete(v5, (^(v21)))
                                    }
                                    var v92 uint16 = (v20 / ((v20) | 1))
                                    _ = v89
                                    _ = v92
                                } else {
                                    if true {
                                        var v93 uint64 = v8
                                        switch (-(fi88)) {
                                        case 32768:
                                            v1 = v1
                                        case 128:
                                            v94 := []int{(fi84 >> ((409) & 63)), fi88}
                                            _ = v94
                                        }
                                        fmt.Println(v4, v8, v4)
                                        _ = v93
                                    } else {
                                        v95 := GenS9{F0: 544, F1: (^(v1)), F2: false}
                                        _ = v95
                                    }
                                    v3 -= -2.25
                                }
                                var v96 uint32 = (v85 & 4294967295)
                                v97 := []int{}
                                _ = fi88
                                _ = v96
                                _ = v97
                            }
                            for fi98 := 0; fi98 < 7; fi98++ {
                                v2 += v4
                                _ = fi98
                            }
                        }
                        switch (^(fi84)) {
                        case 1:
                            for fi99 := 0; fi99 < 20; fi99++ {
                                switch "hello" {
                                case "hello":
                                    switch ("\n\t" + "a") {
                                    case "":
                                        fmt.Println(v1, v21)
                                    default:
                                        if false {
                                            v100 := []int{-880, int(v20), (fi99 ^ fi99), (fi99 / (((fi84 / ((fi84) | 1))) | 1))}
                                            switch (-(fi84)) {
                                            case -41:
                                                v4 -= (v19 + v4)
                                            case 4294967295:
                                                v8 = v8
                                            default:
                                                v19 += v4
                                            }
                                            _ = v100
                                        }
                                    }
                                    switch fi99 {
                                    case 1:
                                        switch ("\n\t" + (("\n\t" + "") + ("\n\t" + ("foo bar" + "a")))) {
                                        case "":
                                            if false {
                                                switch int(v20) {
                                                case 42:
                                                    v101 := []int{fi84}
                                                    _ = v101
                                                }
                                                v102 := GenS10{F0: v85, F1: 42, F2: 0}
                                                v103 := GenS11{F0: "foo bar", F1: v1, F2: v1}
                                                _ = v102
                                                _ = v103
                                            } else {
                                                v104 := GenS7{F0: -1, F1: 335, F2: ((("a" + "") + "") + "hello")}
                                                var v105 int8 = -6
                                                _ = v104
                                                _ = v105
                                            }
                                        case "a":
                                            var v106 int16 = 1
                                            v107 := GenS6{F0: 18446744073709551615, F1: 466, F2: -344}
                                            fmt.Println(v8)
                                            _ = v106
                                            _ = v107
                                        default:
                                            switch ("\n\t" + "") {
                                            case "\n\t":
                                                v4 = v19
                                            default:
                                                v108 := v5[4294967295]
                                                _ = v108
                                            }
                                        }
                                    case -2147483648:
                                        var v109 uint64 = (v8 & v8)
                                        _ = v109
                                    }
                                    if (true && false) {
                                        v110 := v5[v85]
                                        v111 := []int{(^(fi84))}
                                        _ = v110
                                        _ = v111
                                    }
                                case "\n\t":
                                    switch "" {
                                    case "a":
                                        var v112 bool = (0 == 42)
                                        _ = v112
                                    default:
                                        if ((v8 + 301) <= (v8 + (v8 % ((v8) | 1)))) {
                                            switch (("foo bar" + "") + "\n\t") {
                                            case "foo bar":
                                                v113 := []int{(^(fi99)), 0, fi84, 445}
                                                _ = v113
                                            case "\n\t":
                                                v19 -= v4
                                            case "hello":
                                                _ = len(v5)
                                            default:
                                                v20 += v20
                                            }
                                        }
                                    }
                                    switch (fi99 >> ((862) & 63)) {
                                    case 0:
                                        if ((^(v8)) <= (v8 / ((uint64(v20)) | 1))) {
                                            if ((2 < uint64(v20)) && true) {
                                                v2 = v2
                                            }
                                        }
                                    case -2147483648:
                                        fmt.Println(v20, v4, v21)
                                        if ((fi84 - 1) < fi99) {
                                            if true {
                                                v5[v21] = 1
                                            } else {
                                                if false {
                                                    if true {
                                                        v21 = 681
                                                        var v114 uint8 = genFn0()
                                                        switch fi84 {
                                                        case -2147483648:
                                                            v115 := GenS12{F0: (^(v20)), F1: 663, F2: false}
                                                            _ = v115
                                                        case 2147483647:
                                                            var v116 bool = false
                                                            _ = v116
                                                        case 42:
                                                            var v117 uint8 = v1
                                                            _ = v117
                                                        default:
                                                            if false {
                                                                v85 = 42
                                                                _ = len(v5)
                                                                v118 := GenS2{F0: "", F1: 42, F2: (v114 ^ (v1 % (((^(v1))) | 1)))}
                                                                _ = v118
                                                            } else {
                                                                if (2147483647 < 407) {
                                                                    v119 := GenS8{F0: (true || false)}
                                                                    var v120 int32 = 2147483647
                                                                    v121 := []int{}
                                                                    _ = v119
                                                                    _ = v120
                                                                    _ = v121
                                                                }
                                                                v122 := GenS7{F0: 42, F1: int16(v85), F2: "\n\t"}
                                                                _ = v122
                                                            }
                                                        }
                                                        v2 += v2
                                                        _ = v114
                                                    }
                                                    v20 += (v20 + 32767)
                                                    v5[v21] = v85
                                                    fmt.Println(v19)
                                                } else {
                                                    var v123 string = ""
                                                    _ = v123
                                                }
                                                if (1 == 1) {
                                                    delete(v5, (v85 % ((858) | 1)))
                                                    v124 := GenS11{F0: "\n\t", F1: 42, F2: (v1 + genFn0())}
                                                    v1 += 255
                                                    v1 = v1
                                                    _ = v124
                                                } else {
                                                    if true {
                                                        if false {
                                                            var v125 int32 = -988
                                                            _ = v125
                                                        } else {
                                                            var v126 int16 = int16(v20)
                                                            _ = v126
                                                        }
                                                        _ = len(v5)
                                                        fmt.Println(v85, v8, fi84)
                                                        if true {
                                                            v5[(v85 * (^(v21)))] = v21
                                                            v3 = 0.0
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v127 := []int{(fi84 | (fi84 * (fi84 / (((fi84 + -968)) | 1)))), (fi99 >> ((906) & 63)), 9223372036854775807, (fi99 >> ((0) & 63))}
                                                                _ = v127
                                                            default:
                                                                if false {
                                                                    var v128 int16 = -100
                                                                    if (1000 <= v128) {
                                                                        v4 = v2
                                                                        var v129 bool = (float32(v3) < v2)
                                                                        _ = v129
                                                                    }
                                                                    _ = v128
                                                                }
                                                            }
                                                            fmt.Println(v2)
                                                        } else {
                                                            if false {
                                                                v130 := GenS10{F0: 4294967295, F1: 1, F2: 18446744073709551615}
                                                                var v131 uint8 = (v1 * (v1 >> ((65535) & 63)))
                                                                v132 := v5[2147483647]
                                                                _ = v130
                                                                _ = v131
                                                                _ = v132
                                                            } else {
                                                                switch -1 {
                                                                case -1:
                                                                    v133 := GenS6{F0: 193, F1: v8, F2: 0}
                                                                    _ = v133
                                                                default:
                                                                    v134 := GenS10{F0: 932, F1: 42, F2: 2}
                                                                    _ = v134
                                                                }
                                                            }
                                                            var v135 int32 = -2147483648
                                                            _ = v135
                                                        }
                                                    } else {
                                                        if true {
                                                            fmt.Println(v1)
                                                            v136 := []int{}
                                                            if false {
                                                                fmt.Println(v2, v3)
                                                            } else {
                                                                v85 += v85
                                                            }
                                                            if ((("foo bar" + ("foo bar" + ("a" + "foo bar"))) == "foo bar") || true) {
                                                                switch (fi84 | 2147483647) {
                                                                case 9223372036854775807:
                                                                    var v137 uint8 = v1
                                                                    _ = v137
                                                                default:
                                                                    v4 -= (v4 / -2.25)
                                                                }
                                                                if (true || true) {
                                                                    var v138 uint16 = (^(v20))
                                                                    var v139 int32 = 1
                                                                    _ = v138
                                                                    _ = v139
                                                                }
                                                            } else {
                                                                v140 := GenS7{F0: 127, F1: -1, F2: "a"}
                                                                var v141 float64 = (float64(v19) * v3)
                                                                _ = v140
                                                                _ = v141
                                                            }
                                                            _ = v136
                                                        }
                                                        if true {
                                                            v142 := GenS10{F0: 42, F1: int8(v8), F2: 2}
                                                            switch fi99 {
                                                            case 42:
                                                                fmt.Println(v1)
                                                            default:
                                                                if true {
                                                                    _ = len(v5)
                                                                    var v143 int64 = 0
                                                                    _ = v143
                                                                }
                                                            }
                                                            _ = v142
                                                        } else {
                                                            v144 := []int{-1}
                                                            delete(v5, (v85 | (v21 ^ (v21 << ((uint(v1)) & 63)))))
                                                            v19 = v2
                                                            _ = v144
                                                        }
                                                        switch fi99 {
                                                        case 9223372036854775807:
                                                            var v145 uint32 = v21
                                                            _ = v145
                                                        case 4294967295:
                                                            var v146 uint16 = (v20 - v20)
                                                            _ = v146
                                                        case 127:
                                                            v147 := []int{(fi99 - (fi99 ^ fi84)), fi84, int(v21)}
                                                            _ = v147
                                                        }
                                                        var v148 string = "a"
                                                        v149 := []int{}
                                                        fmt.Println(v19, v8, v20)
                                                        _ = v148
                                                        _ = v149
                                                    }
                                                }
                                                if false {
                                                    v150 := GenS13{F0: (^(v1)), F1: -32768, F2: (v20 * v20)}
                                                    v21 = (^(v21))
                                                    if (((v19 - (v4 + 1.5)) < v4) && true) {
                                                        fmt.Println(v8, v8, v2)
                                                        if (1.5 != float32(v3)) {
                                                            v151 := GenS11{F0: "\n\t", F1: v1, F2: (v1 | (v1 >> ((2) & 63)))}
                                                            v20 = v20
                                                            _ = v151
                                                        }
                                                        switch 0 {
                                                        case 128:
                                                            if false {
                                                                var v152 uint32 = (^(v85))
                                                                var v153 uint8 = (^(v1))
                                                                _ = v152
                                                                _ = v153
                                                            } else {
                                                                if (!(v21 > v85)) {
                                                                    var v154 int8 = -128
                                                                    _ = v154
                                                                }
                                                                v155 := GenS1{F0: -9223372036854775808, F1: ""}
                                                                var v156 int16 = 32767
                                                                _ = v155
                                                                _ = v156
                                                            }
                                                        case -2147483648:
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v157 := GenS7{F0: -128, F1: -32768, F2: "hello"}
                                                                var v158 int32 = 42
                                                                _ = v157
                                                                _ = v158
                                                            case "foo bar":
                                                                var v159 uint8 = 159
                                                                _ = v159
                                                            case "hello":
                                                                v160 := GenS14{F0: -9223372036854775808, F1: 0}
                                                                _ = v160
                                                            default:
                                                                if ("hello" < "\n\t") {
                                                                    var v161 string = ("a" + "\n\t")
                                                                    var v162 uint = 46
                                                                    _ = v161
                                                                    _ = v162
                                                                }
                                                                v19 += (v19 * 0.1)
                                                                v21 = 0
                                                            }
                                                        case 32767:
                                                            v2 -= v19
                                                            v2 += float32((0.1 / v3))
                                                            var v163 float32 = (v4 + ((0.1 + (v2 * v2)) * ((v19 * float32(1e10)) / float32(1.5))))
                                                            _ = v163
                                                        default:
                                                            if (!(!(2 >= 600))) {
                                                                var v164 uint8 = 60
                                                                v165 := []int{fi99, (fi84 | fi99), 0}
                                                                if len(v165) >= 2 {
                                                                    v165 = v165[1:len(v165)]
                                                                }
                                                                var v166 uint16 = (^(v20))
                                                                if ((42 != int8(v85)) && (42 <= 0)) {
                                                                    fmt.Println(v3)
                                                                    var v167 uint32 = (v21 / ((1) | 1))
                                                                    if true {
                                                                        v21 += 4294967295
                                                                        fmt.Println(v85)
                                                                        var v168 float32 = -2.25
                                                                        _ = v168
                                                                    }
                                                                    _ = v167
                                                                } else {
                                                                    var v169 bool = false
                                                                    v1 = (v164 % ((v164) | 1))
                                                                    v21 = v21
                                                                    v170 := GenS15{F0: 65535}
                                                                    var v171 uint64 = v8
                                                                    _ = v169
                                                                    _ = v170
                                                                    _ = v171
                                                                }
                                                                fmt.Println(fi84)
                                                                var v172 uint32 = (v85 & v85)
                                                                if ("" != "\n\t") {
                                                                    fmt.Println(v4, v21, fi84)
                                                                }
                                                                _ = v164
                                                                _ = v165
                                                                _ = v166
                                                                _ = v172
                                                            } else {
                                                                if (uint64(v85) == 2) {
                                                                    if ((219 != 1) || ((v1 - genFn0()) < v1)) {
                                                                        _ = len(v5)
                                                                        switch -999 {
                                                                        case -1:
                                                                            var v173 bool = true
                                                                            _ = v173
                                                                        case -687:
                                                                            v8 -= (v8 - v8)
                                                                        case 0:
                                                                            var v174 string = "\n\t"
                                                                            _ = v174
                                                                        default:
                                                                            var v175 uint8 = v1
                                                                            _ = v175
                                                                        }
                                                                        fmt.Println(v3)
                                                                    } else {
                                                                        v21 = 2147483647
                                                                    }
                                                                    var v176 int = fi84
                                                                    var v177 int16 = 42
                                                                    _ = v176
                                                                    _ = v177
                                                                }
                                                                var v178 int32 = -2147483648
                                                                _ = v178
                                                            }
                                                        }
                                                    } else {
                                                        var v179 float64 = (v3 - -2.25)
                                                        v180 := GenS5{F0: (-(fi84)), F1: 766, F2: true}
                                                        v181 := GenS5{F0: -2147483648, F1: uint(v20), F2: (0 < 916)}
                                                        _ = v179
                                                        _ = v180
                                                        _ = v181
                                                    }
                                                    switch (-(fi99)) {
                                                    case -1:
                                                        v182 := []int{fi99, fi84, 127, fi84}
                                                        _ = v182
                                                    case -129:
                                                        v183 := GenS14{F0: 42, F1: 255}
                                                        _ = v183
                                                    default:
                                                        var v184 uint16 = v20
                                                        _ = v184
                                                    }
                                                    if (false || true) {
                                                        switch (fi84 >> ((42) & 63)) {
                                                        case -184:
                                                            v185 := GenS11{F0: ("" + (("a" + "a") + ("" + ("" + "\n\t")))), F1: v1, F2: genFn0()}
                                                            _ = v185
                                                        }
                                                        delete(v5, (v85 / ((2147483648) | 1)))
                                                    }
                                                    _ = v150
                                                } else {
                                                    v20 += v20
                                                    var v186 int8 = 1
                                                    if false {
                                                        v2 -= v2
                                                        v187 := []int{-676, fi99, int(v20)}
                                                        v188 := GenS16{F0: (("hello" + "a") + ""), F1: "foo bar"}
                                                        _ = v187
                                                        _ = v188
                                                    } else {
                                                        v189 := GenS17{F0: (v20 * v20)}
                                                        _ = len(v5)
                                                        v5[123] = v85
                                                        _ = v189
                                                    }
                                                    _ = v186
                                                }
                                            }
                                        } else {
                                            v190 := GenS12{F0: (v20 & (v20 - (v20 & v20))), F1: 645, F2: ("" > "a")}
                                            v191 := GenS11{F0: "hello", F1: 2, F2: (v1 & uint8(v20))}
                                            _ = v190
                                            _ = v191
                                        }
                                        v192 := []int{int(v1), (fi99 ^ (fi99 - fi84))}
                                        _ = v192
                                    }
                                    v19 = ((v4 * float32((-2.25 + (-2.25 + float64((float32(3.14159) - v2)))))) + 0.1)
                                case "":
                                    var v193 bool = true
                                    v193 = v193
                                    _ = v193
                                }
                                v194 := GenS18{F0: -1, F1: int(v8)}
                                _ = fi99
                                _ = v194
                            }
                        case -2147483648:
                            var v195 float64 = v3
                            _ = v195
                        case 128:
                            if true {
                                var v196 uint8 = (^(v1))
                                v20 -= 0
                                _ = v196
                            } else {
                                fmt.Println(v85, v3)
                                delete(v5, v21)
                                if ((!(true && false)) && ((!true) || (-547 != -426))) {
                                    v3 = v3
                                } else {
                                    v20 = (^(v20))
                                    switch (fi84 | fi84) {
                                    case 0:
                                        v19 -= 1e10
                                    case -890:
                                        v5[v21] = 1
                                    case 420:
                                        for fi197 := 0; fi197 < 20; fi197++ {
                                            switch ("a" + "foo bar") {
                                            case "\n\t":
                                                v4 = v19
                                            }
                                            _ = fi197
                                        }
                                    default:
                                        v198 := v5[(v85 * v85)]
                                        _ = v198
                                    }
                                }
                            }
                        }
                    }
                default:
                    var v199 int8 = 100
                    _ = v199
                }
                fmt.Println(v19)
                v2 -= v19
            }
            _ = fi84
            _ = v85
        }
    } else {
        v200 := v5[v21]
        switch (("" + ("a" + "\n\t")) + "") {
        case "a":
            var v201 int16 = 1000
            _ = v201
        case "\n\t":
            var v202 uint32 = (^(v21))
            _ = v202
        case "hello":
            v19 = 0.0
        default:
            v203 := []int{0}
            _ = v203
        }
        var v204 int32 = 550
        _ = v200
        _ = v204
    }
    var v205 float32 = v19
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v205:", v205)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v205
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
