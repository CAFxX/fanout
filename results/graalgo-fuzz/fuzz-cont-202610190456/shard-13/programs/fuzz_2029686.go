package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 int8; F1 int8; F2 uint64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int16; F1 bool; F2 string }
func genFn0(p0 string, p1 uint64, p2 uint16) uint16 {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint8 = 42
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 uint16 = (^(p2))
            var v10 int8 = 127
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 bool = false
        _ = fi6
        _ = v11
        _ = v7
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        switch (fi12 ^ 1) {
        case 32768:
            if false {
                var v13 uint8 = 127
                for fi14 := 0; fi14 < 13; fi14++ {
                    if true {
                        v15 := GenS1{F0: uint(p1)}
                        v16 := []int{}
                        v13 = v13
                        switch (fi12 + fi14) {
                        case 1:
                            if len(v16) > 0 {
                                v16[0] = fi12
                            }
                        case 32768:
                            p0 = "hello"
                        default:
                            v17 := GenS1{F0: 18446744073709551615}
                            _ = v17
                        }
                        fmt.Println(fi12)
                        _ = v15
                        _ = v16
                    }
                    p2 += p2
                    var v18 int64 = -9223372036854775808
                    _ = fi14
                    _ = v18
                }
                var v19 uint8 = (v13 ^ 78)
                _ = v13
                _ = v19
            }
        case -1:
            if true {
                p1 = 2
            }
        }
        for fi20 := 0; fi20 < 3; fi20++ {
            p1 = 883
            var v21 int = (fi20 >> ((977) & 63))
            _ = fi20
            _ = v21
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 uint32 = 219
    var v2 float64 = 0.1
    var v3 int16 = -32768
    var v4 uint32 = 121
    var v5 uint64 = 803
    fmt.Println(v1)
    v22 := []int{2147483647, 9223372036854775807, -151}
    if true {
        v3 -= 0
        v23 := GenS2{F0: true, F1: 128}
        _ = v23
    } else {
        if (true && ((true || true) || (v2 > (float64(1e10) + float64(3.14159))))) {
            v3 = (v3 & (v3 ^ (v3 - (v3 >> ((255) & 63)))))
            var v24 float32 = (((1.5 * 3.14159) - 1.5) * 3.14159)
            _ = v24
        }
        var v25 float64 = 3.14159
        switch (("foo bar" + "a") + "hello") {
        case "":
            if true {
                fmt.Println(v5)
            } else {
                for fi26 := 0; fi26 < 16; fi26++ {
                    _ = len(v22)
                    v4 += v4
                    _ = fi26
                }
                v22 = append(v22, 42)
                fmt.Println(v4, v5)
            }
        case "hello":
            _ = len(v22)
        case "\n\t":
            var v27 uint = uint(v4)
            _ = v27
        default:
            v3 = 1
        }
        _ = v25
    }
    switch ("hello" + "hello") {
    case "\n\t":
        if true {
            switch -2147483648 {
            case -2147483648:
                for fi28 := 0; fi28 < 16; fi28++ {
                    if (false && (((!((false || true) && (!true))) && false) && true)) {
                        for fi29 := 0; fi29 < 13; fi29++ {
                            v5 -= v5
                            _ = fi29
                        }
                    }
                    var v30 uint = 18446744073709551615
                    switch 2147483647 {
                    case 9223372036854775807:
                        var v31 int8 = 0
                        _ = v31
                    default:
                        v2 = float64(-2.25)
                    }
                    _ = fi28
                    _ = v30
                }
            case 67:
                var v32 uint8 = 127
                _ = v32
            case 42:
                fmt.Println(v4, v1)
            }
            for fi33 := 0; fi33 < 10; fi33++ {
                fmt.Println(v5, v1)
                _ = fi33
            }
        }
    case "foo bar":
        for fi34 := 0; fi34 < 2; fi34++ {
            var v35 uint8 = 51
            var v36 string = "\n\t"
            v3 += (v3 * v3)
            _ = fi34
            _ = v35
            _ = v36
        }
    case "":
        v37 := GenS1{F0: 389}
        _ = v37
    default:
        var v38 int16 = (-(v3))
        _ = v38
    }
    for fi39 := 0; fi39 < 3; fi39++ {
        v2 = 0.0
        _ = fi39
    }
    v22 = append(v22, 32767)
    var v40 uint32 = 191
    var v41 uint32 = v1
    for fi42 := 0; fi42 < 11; fi42++ {
        _ = len(v22)
        v43 := GenS1{F0: 113}
        _ = fi42
        _ = v43
    }
    if (!false) {
        var v44 float64 = (((0.0 / 3.14159) / v2) / float64(3.14159))
        _ = v44
    } else {
        if len(v22) > 0 {
            v22[0] = 32767
        }
        for fi45 := 0; fi45 < 18; fi45++ {
            for fi46 := 0; fi46 < 17; fi46++ {
                if false {
                    v47 := GenS3{F0: 0, F1: v3, F2: 128}
                    _ = v47
                } else {
                    var v48 uint = 407
                    switch (fi45 * (^(fi46))) {
                    case 128:
                        if false {
                            v49 := GenS1{F0: 1}
                            fmt.Println(fi45, v4, v1)
                            v50 := GenS1{F0: v48}
                            _ = v49
                            _ = v50
                        } else {
                            switch ("hello" + (("\n\t" + ("a" + "hello")) + "foo bar")) {
                            case "hello":
                                v51 := GenS4{F0: 0, F1: 0}
                                _ = v51
                            default:
                                switch 1 {
                                case 128:
                                    v52 := GenS5{F0: 0, F1: 255, F2: -2147483648}
                                    _ = v52
                                case -59:
                                    var v53 int32 = 42
                                    _ = v53
                                default:
                                    if len(v22) > 0 {
                                        v22[0] = (fi46 % (((^(fi46))) | 1))
                                    }
                                    if (false || false) {
                                        var v54 int = 128
                                        _ = v54
                                    } else {
                                        v40 = v41
                                        v2 = v2
                                    }
                                }
                            }
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                        }
                    default:
                        var v55 uint32 = v1
                        _ = v55
                    }
                    _ = v48
                }
                var v56 uint64 = v5
                v2 -= 3.14159
                _ = fi46
                _ = v56
            }
            var v57 uint8 = 98
            for fi58 := 0; fi58 < 20; fi58++ {
                v1 = (^(v4))
                var v59 int8 = 100
                switch ((("\n\t" + "") + ("foo bar" + ("hello" + "hello"))) + "a") {
                case "\n\t":
                    v4 = v41
                case "":
                    v3 -= (v3 * v3)
                case "hello":
                    v40 -= (v1 ^ (^(v4)))
                default:
                    v59 = -1
                }
                var v60 uint64 = (v5 & v5)
                _ = fi58
                _ = v59
                _ = v60
            }
            _ = fi45
            _ = v57
        }
    }
    if ((0.1 <= v2) && (true && true)) {
        fmt.Println(v2, v5, v3)
        v61 := GenS1{F0: 142}
        _ = v61
    } else {
        switch -9223372036854775808 {
        case 9223372036854775807:
            v41 += 2147483647
        }
        v22 = append(v22, 9223372036854775807)
    }
    for fi62 := 0; fi62 < 9; fi62++ {
        if (float64((0.1 / 0.1)) <= v2) {
            switch (fi62 | fi62) {
            case -9223372036854775808:
                if len(v22) > 0 {
                    v22[0] = (-(fi62))
                }
            default:
                if (0.1 < (0.1 - (1e10 * 3.14159))) {
                    if true {
                        v63 := GenS6{F0: 0, F1: v3}
                        if ((0 == 2) || (v5 <= v5)) {
                            for fi64 := 0; fi64 < 4; fi64++ {
                                if len(v22) > 0 {
                                    v22[0] = fi64
                                }
                                _ = fi64
                            }
                            for fi65 := 0; fi65 < 1; fi65++ {
                                switch ("foo bar" + ("hello" + "\n\t")) {
                                case "\n\t":
                                    var v66 int32 = 2147483647
                                    _ = v66
                                case "":
                                    v3 -= -33
                                }
                                if false {
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v67 float64 = ((float64((-2.25 - 0.1)) * v2) * ((1.5 * (v2 / ((3.14159 / ((v2 / 1e10) - v2)) - 1e10))) * v2))
                                        _ = v67
                                    default:
                                        if len(v22) > 0 {
                                            v22[0] = fi65
                                        }
                                        v41 -= (v1 % ((v1) | 1))
                                        v22 = append(v22, fi65)
                                    }
                                    v3 = -1
                                    if true {
                                        var v68 int32 = 1
                                        var v69 int8 = 64
                                        if true {
                                            var v70 float32 = ((((1.5 + 3.14159) / -2.25) / 1e10) - ((3.14159 + 1.5) / (3.14159 / -2.25)))
                                            var v71 uint = 2
                                            var v72 float64 = v2
                                            var v73 float32 = v70
                                            _ = v70
                                            _ = v71
                                            _ = v72
                                            _ = v73
                                        } else {
                                            var v74 uint8 = 127
                                            fmt.Println(v4)
                                            v75 := GenS2{F0: (!true), F1: 42}
                                            _ = v74
                                            _ = v75
                                        }
                                        _ = v68
                                        _ = v69
                                    } else {
                                        var v76 string = "\n\t"
                                        fmt.Println(v4)
                                        fmt.Println(v1)
                                        _ = v76
                                    }
                                    v22 = append(v22, (^(fi62)))
                                } else {
                                    switch 42 {
                                    case 577:
                                        var v77 int16 = v3
                                        _ = v77
                                    case 515:
                                        v41 = v40
                                    }
                                    var v78 int8 = -32
                                    _ = len(v22)
                                    _ = v78
                                }
                                _ = fi65
                            }
                        }
                        _ = v63
                    } else {
                        switch fi62 {
                        case 42:
                            var v79 bool = (!true)
                            _ = v79
                        case 685:
                            if ((v1 != 0) || ((v1 * (v4 << ((1) & 63))) >= (v1 % ((v40) | 1)))) {
                                fmt.Println(v1, v1)
                                for fi80 := 0; fi80 < 19; fi80++ {
                                    if (-1 > -808) {
                                        fmt.Println(v4, v1, v3)
                                        var v81 bool = (false || false)
                                        if len(v22) > 0 {
                                            v22[0] = (fi80 + (fi80 >> ((42) & 63)))
                                        }
                                        _ = v81
                                    } else {
                                        v22 = append(v22, (fi62 & -2147483648))
                                        var v82 string = ("\n\t" + "foo bar")
                                        if len(v22) > 0 {
                                            v22[0] = (fi62 | fi80)
                                        }
                                        _ = v82
                                    }
                                    var v83 uint64 = v5
                                    if false {
                                        switch "a" {
                                        case "foo bar":
                                            var v84 float32 = (3.14159 / ((3.14159 * 1e10) + float32(0.0)))
                                            _ = v84
                                        }
                                        if false {
                                            var v85 uint8 = 127
                                            v22 = append(v22, (fi80 % ((fi62) | 1)))
                                            _ = v85
                                        }
                                        var v86 int = (^(fi62))
                                        _ = v86
                                    } else {
                                        v87 := GenS2{F0: (1000 == (v3 / ((42) | 1))), F1: 1}
                                        switch 4294967295 {
                                        case 4294967295:
                                            fmt.Println(v3, v4, fi80)
                                        case -2147483648:
                                            if (150 != 381) {
                                                v88 := GenS5{F0: 32767, F1: 62, F2: 2147483647}
                                                switch 1 {
                                                case -218:
                                                    var v89 float64 = (v2 / v2)
                                                    _ = v89
                                                }
                                                _ = v88
                                            } else {
                                                var v90 uint32 = 0
                                                _ = v90
                                            }
                                        default:
                                            v91 := GenS6{F0: 9223372036854775807, F1: 32767}
                                            _ = v91
                                        }
                                        if false {
                                            var v92 string = ("hello" + "foo bar")
                                            var v93 int8 = 127
                                            v1 += v1
                                            _ = v92
                                            _ = v93
                                        } else {
                                            v94 := GenS1{F0: 963}
                                            _ = v94
                                        }
                                        _ = v87
                                    }
                                    _ = fi80
                                    _ = v83
                                }
                                if true {
                                    switch ("\n\t" + "") {
                                    case "hello":
                                        for fi95 := 0; fi95 < 3; fi95++ {
                                            switch ("hello" + ("a" + "hello")) {
                                            case "":
                                                var v96 uint8 = 179
                                                _ = v96
                                            case "foo bar":
                                                if (!(true || false)) {
                                                    switch fi62 {
                                                    case 1:
                                                        v97 := GenS7{F0: -128, F1: -20, F2: 612}
                                                        _ = v97
                                                    case 732:
                                                        var v98 float64 = float64((0.1 * -2.25))
                                                        _ = v98
                                                    case -1:
                                                        if true {
                                                            var v99 uint16 = 42
                                                            _ = v99
                                                        }
                                                    default:
                                                        if (!((0.1 + ((3.14159 / float32(v2)) - float32(1.5))) != (0.1 * float32(float64((0.0 - 0.1)))))) {
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                _ = len(v22)
                                                            case "foo bar":
                                                                if false {
                                                                    if (true && true) {
                                                                        fmt.Println(v5, v40)
                                                                        v100 := GenS8{F0: true}
                                                                        _ = len(v22)
                                                                        _ = v100
                                                                    } else {
                                                                        fmt.Println(v3)
                                                                        var v101 uint16 = 42
                                                                        _ = v101
                                                                    }
                                                                    v41 = 276
                                                                } else {
                                                                    v5 += v5
                                                                }
                                                            }
                                                        }
                                                    }
                                                    var v102 float32 = (-2.25 + 3.14159)
                                                    var v103 int = (-(fi95))
                                                    if false {
                                                        if (0 < 241) {
                                                            switch 32767 {
                                                            case 73:
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v104 bool = false
                                                                    _ = v104
                                                                }
                                                            case -221:
                                                                v5 += v5
                                                            case 696:
                                                                var v105 string = ("\n\t" + (("\n\t" + "hello") + "foo bar"))
                                                                _ = v105
                                                            default:
                                                                fmt.Println(v4, v41)
                                                            }
                                                            if len(v22) >= 2 {
                                                                v22 = v22[1:len(v22)]
                                                            }
                                                        } else {
                                                            switch 9223372036854775807 {
                                                            case -1:
                                                                var v106 bool = false
                                                                _ = v106
                                                            case -2147483648:
                                                                if (("hello" + "hello") != "\n\t") {
                                                                    var v107 string = "foo bar"
                                                                    _ = v107
                                                                }
                                                            case 0:
                                                                var v108 string = "a"
                                                                _ = v108
                                                            }
                                                            if true {
                                                                if (fi62 >= fi62) {
                                                                    if (0 > fi62) {
                                                                        if len(v22) > 0 {
                                                                            v22[0] = (v103 * 9223372036854775807)
                                                                        }
                                                                        if (-30 == -128) {
                                                                            v102 -= 0.0
                                                                            if false {
                                                                                v4 = 347
                                                                                var v109 float32 = (v102 - 1e10)
                                                                                fmt.Println(v109, v5, v109)
                                                                                _ = v109
                                                                            }
                                                                        }
                                                                        v3 = v3
                                                                    } else {
                                                                        var v110 int32 = 955
                                                                        v110 = (^(v110))
                                                                        if false {
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                _ = len(v22)
                                                                                v22 = append(v22, -9223372036854775808)
                                                                                switch ("" + "foo bar") {
                                                                                case "\n\t":
                                                                                    fmt.Println(v2, fi95)
                                                                                }
                                                                            }
                                                                            v111 := GenS6{F0: -1, F1: v3}
                                                                            var v112 uint8 = 134
                                                                            _ = v111
                                                                            _ = v112
                                                                        }
                                                                        _ = v110
                                                                    }
                                                                    v103 -= 949
                                                                    var v113 string = "hello"
                                                                    _ = v113
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        if len(v22) >= 2 {
                                                                            v22 = v22[1:len(v22)]
                                                                        }
                                                                    case "foo bar":
                                                                        v114 := GenS9{F0: 9223372036854775807}
                                                                        _ = v114
                                                                    default:
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            v115 := GenS5{F0: 1, F1: 1, F2: 2147483647}
                                                                            _ = v115
                                                                        case "foo bar":
                                                                            v40 += (v41 << ((65535) & 63))
                                                                        }
                                                                    }
                                                                    fmt.Println(fi95)
                                                                }
                                                            }
                                                        }
                                                        var v116 int32 = 0
                                                        var v117 uint64 = 1
                                                        var v118 int = (fi95 << ((4294967295) & 63))
                                                        _ = v116
                                                        _ = v117
                                                        _ = v118
                                                    } else {
                                                        if false {
                                                            var v119 int32 = 558
                                                            fmt.Println(v5)
                                                            fmt.Println(v3, fi62)
                                                            _ = v119
                                                        }
                                                        v120 := GenS9{F0: -255}
                                                        _ = v120
                                                    }
                                                    _ = v102
                                                    _ = v103
                                                } else {
                                                    var v121 uint16 = 65535
                                                    _ = v121
                                                }
                                            }
                                            if (!(!true)) {
                                                v122 := GenS9{F0: -1}
                                                var v123 uint16 = 1
                                                if (0 == 255) {
                                                    if (("" + ((("\n\t" + ("\n\t" + ("" + "hello"))) + "foo bar") + ("a" + "hello"))) == ("foo bar" + (("hello" + "hello") + ("" + "a")))) {
                                                        if (1.5 >= (1e10 + 0.1)) {
                                                            v124 := GenS2{F0: ("a" > "foo bar"), F1: 42}
                                                            _ = v124
                                                        } else {
                                                            if len(v22) > 0 {
                                                                v22[0] = (^(fi62))
                                                            }
                                                            var v125 float64 = (float64(3.14159) * v2)
                                                            fmt.Println(v1)
                                                            v1 = (^(v41))
                                                            _ = v125
                                                        }
                                                        fmt.Println(v1, v4)
                                                    }
                                                } else {
                                                    v4 += uint32(v5)
                                                    var v126 float64 = (0.1 - 3.14159)
                                                    v127 := GenS8{F0: true}
                                                    _ = v126
                                                    _ = v127
                                                }
                                                switch (("hello" + "a") + "\n\t") {
                                                case "hello":
                                                    var v128 uint32 = v41
                                                    _ = v128
                                                }
                                                _ = v122
                                                _ = v123
                                            } else {
                                                switch (((("foo bar" + "") + "") + "foo bar") + "foo bar") {
                                                case "a":
                                                    switch (("hello" + "a") + ("" + "\n\t")) {
                                                    case "a":
                                                        v1 -= v1
                                                    case "\n\t":
                                                        switch (("a" + "hello") + ("a" + ("hello" + "foo bar"))) {
                                                        case "a":
                                                            switch fi62 {
                                                            case 0:
                                                                switch "foo bar" {
                                                                case "":
                                                                    var v129 uint8 = 255
                                                                    _ = v129
                                                                }
                                                            case 2147483647:
                                                                v130 := GenS10{F0: 255}
                                                                _ = v130
                                                            case 4294967295:
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v41 -= (v41 - v4)
                                                                case "hello":
                                                                    fmt.Println(fi62)
                                                                case "":
                                                                    fmt.Println(v3, v41)
                                                                }
                                                            }
                                                        default:
                                                            if len(v22) > 0 {
                                                                v22[0] = (fi62 ^ (fi95 ^ (-(fi62))))
                                                            }
                                                        }
                                                    }
                                                case "hello":
                                                    if len(v22) >= 2 {
                                                        v22 = v22[1:len(v22)]
                                                    }
                                                default:
                                                    v22 = append(v22, -2147483648)
                                                }
                                                var v131 string = "foo bar"
                                                _ = v131
                                            }
                                            _ = fi95
                                        }
                                    case "foo bar":
                                        var v132 uint = 18446744073709551615
                                        _ = v132
                                    default:
                                        switch fi62 {
                                        case 2147483647:
                                            var v133 uint64 = (v5 >> ((818) & 63))
                                            _ = v133
                                        case 42:
                                            if ((2 <= 4294967295) && false) {
                                                v134 := GenS10{F0: 18446744073709551615}
                                                if true {
                                                    var v135 int32 = 1
                                                    for fi136 := 0; fi136 < 16; fi136++ {
                                                        if len(v22) >= 2 {
                                                            v22 = v22[1:len(v22)]
                                                        }
                                                        if (v3 > (-(v3))) {
                                                            var v137 uint32 = (v41 - v4)
                                                            var v138 uint8 = 1
                                                            v139 := GenS8{F0: (1 == -947)}
                                                            _ = v137
                                                            _ = v138
                                                            _ = v139
                                                        }
                                                        v40 += 2147483648
                                                        _ = fi136
                                                    }
                                                    _ = v135
                                                } else {
                                                    var v140 uint = 2
                                                    v141 := GenS5{F0: 558, F1: 47, F2: -550}
                                                    v5 -= (v5 * v5)
                                                    _ = v140
                                                    _ = v141
                                                }
                                                _ = v134
                                            } else {
                                                v142 := GenS3{F0: 79, F1: 296, F2: 127}
                                                v5 -= v5
                                                _ = v142
                                            }
                                        case -9223372036854775808:
                                            var v143 uint16 = 0
                                            _ = v143
                                        default:
                                            switch ("foo bar" + "\n\t") {
                                            case "a":
                                                fmt.Println(v1)
                                            case "hello":
                                                var v144 uint16 = 42
                                                _ = v144
                                            case "":
                                                switch "hello" {
                                                case "hello":
                                                    v2 = v2
                                                case "":
                                                    fmt.Println(v4)
                                                case "foo bar":
                                                    var v145 bool = (!((949 != 0) && true))
                                                    _ = v145
                                                }
                                            default:
                                                v146 := GenS6{F0: 467, F1: (-(v3))}
                                                _ = v146
                                            }
                                        }
                                    }
                                    for fi147 := 0; fi147 < 10; fi147++ {
                                        var v148 int8 = 127
                                        var v149 uint16 = genFn0((("\n\t" + "hello") + "\n\t"), uint64(v1), 243)
                                        switch (fi147 / (((fi62 * (fi62 << ((803) & 63)))) | 1)) {
                                        case -136:
                                            v148 = -41
                                        case 32767:
                                            var v150 uint = 255
                                            _ = v150
                                        case 128:
                                            v151 := GenS7{F0: v148, F1: (v148 % (((^(v148))) | 1)), F2: v5}
                                            _ = v151
                                        default:
                                            switch (-(fi62)) {
                                            case 1:
                                                var v152 int32 = -1
                                                _ = v152
                                            case -1:
                                                v41 += (v4 - 65535)
                                            default:
                                                var v153 uint32 = (v4 ^ v4)
                                                _ = v153
                                            }
                                        }
                                        var v154 uint = 127
                                        v40 = v40
                                        _ = fi147
                                        _ = v148
                                        _ = v149
                                        _ = v154
                                    }
                                }
                            }
                        case 128:
                            v2 = v2
                        }
                    }
                    switch (fi62 + 2147483647) {
                    case 523:
                        for fi155 := 0; fi155 < 11; fi155++ {
                            v2 = 1e10
                            _ = fi155
                        }
                    case -1:
                        var v156 int = 9223372036854775807
                        _ = v156
                    case 9223372036854775807:
                        if true {
                            fmt.Println(v4)
                        }
                    default:
                        if true {
                            fmt.Println(v40, v4)
                            switch fi62 {
                            case 548:
                                var v157 uint = 42
                                _ = v157
                            case 32767:
                                fmt.Println(v1, v1)
                            default:
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                            }
                            fmt.Println(v41, v5)
                        }
                    }
                } else {
                    var v158 uint8 = 96
                    var v159 int = 127
                    _ = v158
                    _ = v159
                }
            }
        } else {
            for fi160 := 0; fi160 < 6; fi160++ {
                switch "a" {
                case "hello":
                    v161 := GenS4{F0: 518, F1: 2147483647}
                    _ = v161
                }
                var v162 uint = 42
                _ = fi160
                _ = v162
            }
            if false {
                v163 := GenS6{F0: 0, F1: 1}
                for fi164 := 0; fi164 < 14; fi164++ {
                    fmt.Println(v4, v4, v1)
                    v165 := GenS4{F0: 42, F1: 1}
                    _ = fi164
                    _ = v165
                }
                _ = v163
            }
        }
        _ = fi62
    }
    v4 = (v4 + (^(v4)))
    if len(v22) > 0 {
        v22[0] = -259
    }
    v166 := GenS3{F0: 127, F1: int16(v40), F2: 0}
    fmt.Println(v41, v40, v3)
    v167 := GenS11{F0: (v3 * v3), F1: (false && (true && (!true))), F2: ""}
    v22 = append(v22, 437)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v166
    _ = v167
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
