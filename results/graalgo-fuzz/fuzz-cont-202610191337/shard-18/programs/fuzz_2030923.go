package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint8; F1 int; F2 int8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 uint; F2 uint32 }
type GenS8 struct { F0 string; F1 uint8; F2 int32 }
type GenS9 struct { F0 int64; F1 uint8; F2 string }
type GenS10 struct { F0 uint8; F1 int8; F2 string }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 string; F1 int; F2 bool }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 bool; F1 uint32 }
type GenS15 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS16 struct { F0 string; F1 bool }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS19 struct { F0 uint32; F1 int64 }

func main() {
    var v1 int32 = 42
    var v2 uint64 = 1
    var v3 float64 = 1.5
    var v4 int16 = 42
    for fi5 := 0; fi5 < 3; fi5++ {
        v4 = v4
        v6 := GenS1{F0: 1}
        v1 = (v1 << ((2) & 63))
        _ = fi5
        _ = v6
    }
    var v7 uint64 = v2
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 string = ("\n\t" + "foo bar")
        v10 := GenS1{F0: (v1 / ((v1) | 1))}
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        fmt.Println(v7, v7)
        _ = fi11
    }
    var v12 uint16 = 926
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := []int{(fi13 - 128)}
        _ = fi13
        _ = v14
    }
    fmt.Println(v1)
    switch "a" {
    case "foo bar":
        if false {
            v2 -= 0
            if (!false) {
                for fi15 := 0; fi15 < 8; fi15++ {
                    var v16 uint64 = (v2 << ((1) & 63))
                    _ = fi15
                    _ = v16
                }
            } else {
                v17 := []int{32768, 32768, -2147483648, 42}
                for fi18 := 0; fi18 < 19; fi18++ {
                    v12 += 42
                    v19 := GenS2{F0: v4}
                    var v20 uint64 = (^(v2))
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = v17
            }
        }
    case "\n\t":
        if (0 == 127) {
            var v21 int64 = -837
            for fi22 := 0; fi22 < 19; fi22++ {
                if (false || true) {
                    if false {
                        fmt.Println(v2, v21, v1)
                        if (-1 >= -103) {
                            switch (fi22 | fi22) {
                            case -2147483648:
                                var v23 uint32 = 4294967295
                                _ = v23
                            default:
                                fmt.Println(v21)
                            }
                            switch ("hello" + "") {
                            case "":
                                v3 = 1.5
                            }
                            if true {
                                if (int32(fi22) >= -645) {
                                    var v24 int8 = 119
                                    _ = v24
                                }
                            }
                        } else {
                            fmt.Println(v1)
                            if ((^(v21)) >= (v21 >> ((65535) & 63))) {
                                if true {
                                    switch (fi22 / ((1) | 1)) {
                                    case 2147483647:
                                        switch (("foo bar" + "") + "") {
                                        case "\n\t":
                                            for fi25 := 0; fi25 < 3; fi25++ {
                                                switch fi22 {
                                                case -657:
                                                    switch ("" + "\n\t") {
                                                    case "foo bar":
                                                        v26 := GenS2{F0: (v4 * v4)}
                                                        _ = v26
                                                    }
                                                }
                                                if false {
                                                    switch (fi25 & 127) {
                                                    case 151:
                                                        v27 := []int{}
                                                        _ = v27
                                                    case -2147483648:
                                                        var v28 uint16 = (v12 * 926)
                                                        _ = v28
                                                    }
                                                    switch ("\n\t" + ("a" + "\n\t")) {
                                                    case "foo bar":
                                                        fmt.Println(v3, v12, v7)
                                                    case "hello":
                                                        if ((!false) && (42 == 0)) {
                                                            var v29 int = -379
                                                            v30 := GenS2{F0: v4}
                                                            var v31 uint32 = 4294967295
                                                            _ = v29
                                                            _ = v30
                                                            _ = v31
                                                        } else {
                                                            fmt.Println(v4)
                                                        }
                                                    case "":
                                                        switch (fi22 ^ (fi25 + fi25)) {
                                                        case 42:
                                                            switch -993 {
                                                            case 788:
                                                                v1 = (v1 - (^(v1)))
                                                            }
                                                        case 128:
                                                            v32 := GenS2{F0: (^(v4))}
                                                            _ = v32
                                                        case 32768:
                                                            v12 = v12
                                                        default:
                                                            v1 -= -81
                                                            var v33 float64 = -2.25
                                                            _ = v33
                                                        }
                                                    default:
                                                        v34 := []int{(fi22 & fi25), (fi25 & (fi25 ^ (fi22 / ((fi22) | 1)))), (fi25 >> ((18446744073709551615) & 63)), (fi25 ^ (fi22 + fi25))}
                                                        v2 = v7
                                                        _ = v34
                                                    }
                                                    v2 = 2
                                                    if (true || (100 > -28)) {
                                                        v35 := GenS3{F0: fi22}
                                                        v36 := GenS3{F0: -2147483648}
                                                        _ = v35
                                                        _ = v36
                                                    } else {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            var v37 uint64 = v2
                                                            _ = v37
                                                        case "":
                                                            var v38 float64 = v3
                                                            _ = v38
                                                        default:
                                                            v7 = v7
                                                            fmt.Println(v2)
                                                            var v39 float64 = (float64((float32(3.14159) + 1e10)) + ((v3 - v3) + 0.0))
                                                            v40 := []int{9223372036854775807, (fi25 / (((fi22 - (fi25 | 370))) | 1))}
                                                            var v41 string = "a"
                                                            _ = v39
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                        v7 += (^(v7))
                                                        var v42 float64 = (((1.5 - (float64(0.1) - v3)) - (1e10 + v3)) / -2.25)
                                                        v12 = 255
                                                        var v43 float32 = ((float32(((1e10 * v42) * v3)) - (0.1 + (((float32(v3) * (((((0.0 * (1.5 * 0.0)) + 0.1) + (-2.25 / 3.14159)) - (1e10 - 0.1)) - ((0.1 / 1.0) * (0.1 - (0.1 - (0.0 * 1.5)))))) * 0.1) * 3.14159))) * -2.25)
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    if false {
                                                        v44 := []int{}
                                                        _ = v44
                                                    } else {
                                                        switch fi22 {
                                                        case 128:
                                                            fmt.Println(v7, v3, fi25)
                                                        }
                                                        v4 += (v4 & (v4 + (-(v4))))
                                                        v45 := []int{fi25, (-(fi22)), fi25}
                                                        _ = v45
                                                    }
                                                } else {
                                                    v46 := GenS4{F0: 0, F1: 33}
                                                    _ = v46
                                                }
                                                _ = fi25
                                            }
                                        case "hello":
                                            switch (("a" + ("foo bar" + "")) + ("a" + "a")) {
                                            case "foo bar":
                                                var v47 int32 = 2147483647
                                                _ = v47
                                            case "a":
                                                v48 := GenS3{F0: fi22}
                                                _ = v48
                                            case "\n\t":
                                                v2 = uint64(v12)
                                            }
                                        case "foo bar":
                                            var v49 int16 = 643
                                            _ = v49
                                        default:
                                            v1 = -138
                                        }
                                    }
                                    if false {
                                        for fi50 := 0; fi50 < 14; fi50++ {
                                            v2 -= v7
                                            v21 += v21
                                            _ = fi50
                                        }
                                    } else {
                                        if (true || false) {
                                            for fi51 := 0; fi51 < 19; fi51++ {
                                                var v52 float64 = 3.14159
                                                var v53 int8 = 127
                                                _ = fi51
                                                _ = v52
                                                _ = v53
                                            }
                                        }
                                        v54 := []int{(fi22 ^ (fi22 / ((fi22) | 1))), fi22, (^(fi22)), (fi22 - fi22)}
                                        fmt.Println(v3)
                                        _ = v54
                                    }
                                }
                                for fi55 := 0; fi55 < 9; fi55++ {
                                    if (((3.14159 * 0.0) * 1.5) < 1.5) {
                                        if true {
                                            var v56 int32 = 2147483647
                                            _ = v56
                                        } else {
                                            fmt.Println(v7, fi22, v3)
                                            if false {
                                                v57 := GenS5{F0: 0, F1: (fi55 << ((254) & 63)), F2: 1}
                                                var v58 float32 = 3.14159
                                                var v59 uint = 0
                                                _ = v57
                                                _ = v58
                                                _ = v59
                                            }
                                            if false {
                                                if (!((("foo bar" + ("" + "\n\t")) >= (("hello" + "") + "hello")) && (646 == 65535))) {
                                                    if true {
                                                        fmt.Println(v4, v7)
                                                        var v60 uint32 = 761
                                                        var v61 uint32 = (v60 / ((v60) | 1))
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                }
                                                v3 -= float64(0.1)
                                                v62 := []int{(-(fi22)), (-(fi55))}
                                                _ = v62
                                            }
                                        }
                                        if false {
                                            fmt.Println(v12)
                                            v63 := GenS2{F0: (v4 ^ v4)}
                                            v64 := GenS6{F0: (v2 - (v2 * v7))}
                                            fmt.Println(v3, v12, fi22)
                                            switch fi22 {
                                            case 1:
                                                v65 := GenS4{F0: (v7 | 0), F1: 127}
                                                _ = v65
                                            case 42:
                                                v3 = -2.25
                                            case -875:
                                                var v66 uint64 = v7
                                                _ = v66
                                            }
                                            var v67 uint64 = v7
                                            v7 = v2
                                            v68 := GenS6{F0: v2}
                                            switch "a" {
                                            case "a":
                                                switch fi55 {
                                                case 760:
                                                    fmt.Println(v67)
                                                    var v69 uint16 = (v12 | (v12 | (v12 & 1)))
                                                    switch ("" + ("hello" + "hello")) {
                                                    case "":
                                                        var v70 int16 = (^(v4))
                                                        _ = v70
                                                    case "hello":
                                                        var v71 uint8 = 164
                                                        _ = v71
                                                    case "\n\t":
                                                        var v72 float64 = (v3 - (v3 * v3))
                                                        _ = v72
                                                    default:
                                                        v21 += 1
                                                        v73 := []int{fi22, fi22}
                                                        _ = v73
                                                    }
                                                    _ = v69
                                                case 731:
                                                    var v74 uint = 342
                                                    _ = v74
                                                case -166:
                                                    fmt.Println(v67, fi55)
                                                    var v75 uint64 = (^(v2))
                                                    var v76 uint = 299
                                                    _ = v75
                                                    _ = v76
                                                }
                                            case "":
                                                switch (("hello" + "a") + "") {
                                                case "foo bar":
                                                    if true {
                                                        switch (fi55 / (((fi22 & fi22)) | 1)) {
                                                        case 860:
                                                            v77 := GenS7{F0: (v4 / ((v4) | 1)), F1: 42, F2: 2147483648}
                                                            _ = v77
                                                        }
                                                        var v78 uint32 = 4294967295
                                                        _ = v78
                                                    } else {
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v79 uint64 = (v2 / ((15) | 1))
                                                            _ = v79
                                                        case "foo bar":
                                                            v3 = 0.1
                                                        }
                                                    }
                                                case "\n\t":
                                                    v80 := []int{(^(fi22)), (^(fi55)), (fi22 % ((fi22) | 1))}
                                                    _ = v80
                                                default:
                                                    fmt.Println(v67)
                                                }
                                            case "\n\t":
                                                var v81 int = fi55
                                                _ = v81
                                            default:
                                                v82 := GenS8{F0: "foo bar", F1: 2, F2: -1}
                                                _ = v82
                                            }
                                            v83 := GenS1{F0: (v1 >> ((647) & 63))}
                                            switch ("\n\t" + "hello") {
                                            case "a":
                                                switch (fi22 - fi55) {
                                                case 127:
                                                    v84 := GenS1{F0: v1}
                                                    _ = v84
                                                case -476:
                                                    fmt.Println(v1, v7)
                                                case 32768:
                                                    var v85 uint = 1
                                                    _ = v85
                                                }
                                            case "foo bar":
                                                v3 = 1.5
                                            case "hello":
                                                if false {
                                                    fmt.Println(fi55, v1)
                                                    v86 := []int{}
                                                    var v87 string = ("hello" + "a")
                                                    _ = v86
                                                    _ = v87
                                                } else {
                                                    var v88 int = (^(fi55))
                                                    if (!(v1 >= (v1 % ((v1) | 1)))) {
                                                        var v89 uint = 0
                                                        switch v88 {
                                                        case 17:
                                                            v90 := GenS1{F0: (v1 ^ (v1 >> ((v89) & 63)))}
                                                            _ = v90
                                                        case -1:
                                                            if (v21 == 608) {
                                                                v1 = int32(v12)
                                                            }
                                                        }
                                                        v91 := []int{(fi22 << ((v89) & 63)), -2147483648}
                                                        _ = v89
                                                        _ = v91
                                                    } else {
                                                        var v92 bool = false
                                                        v93 := GenS1{F0: 1}
                                                        fmt.Println(v88, v67)
                                                        _ = v92
                                                        _ = v93
                                                    }
                                                    fmt.Println(v21, v67, fi55)
                                                    _ = v88
                                                }
                                            default:
                                                v94 := []int{(fi55 & (fi55 % (((fi55 + -9223372036854775808)) | 1))), fi22}
                                                _ = v94
                                            }
                                            _ = v63
                                            _ = v64
                                            _ = v67
                                            _ = v68
                                            _ = v83
                                        }
                                        switch "hello" {
                                        case "hello":
                                            if false {
                                                if (((false && (255 < v12)) && (v1 >= (^(v1)))) && (!false)) {
                                                    if true {
                                                        v3 = 1.5
                                                        v4 -= 32767
                                                    }
                                                    v95 := GenS8{F0: "foo bar", F1: 15, F2: (v1 * v1)}
                                                    if false {
                                                        v96 := []int{fi55, fi22}
                                                        v1 -= v1
                                                        _ = v96
                                                    }
                                                    _ = v95
                                                }
                                                fmt.Println(v4, v7, fi22)
                                                v97 := []int{32768}
                                                _ = v97
                                            }
                                        default:
                                            switch 127 {
                                            case 2147483647:
                                                if ((3.14159 / (0.1 - 0.0)) != 0.1) {
                                                    fmt.Println(v4, v4)
                                                    v4 += (v4 + int16(v21))
                                                }
                                            default:
                                                v98 := []int{(^(fi55)), fi55}
                                                _ = v98
                                            }
                                        }
                                    }
                                    fmt.Println(v21, fi55, v2)
                                    v99 := []int{}
                                    _ = fi55
                                    _ = v99
                                }
                                if (1 >= 255) {
                                    switch -300 {
                                    case 9223372036854775807:
                                        var v100 float32 = -2.25
                                        _ = v100
                                    default:
                                        if false {
                                            var v101 string = "foo bar"
                                            _ = v101
                                        } else {
                                            v4 = -32768
                                            var v102 float64 = ((float64(0.1) / 1.0) * -2.25)
                                            v102 = (float64(((float32(v3) / 3.14159) / 1.0)) - v3)
                                            _ = v102
                                        }
                                    }
                                    var v103 int8 = 100
                                    _ = v103
                                } else {
                                    v7 -= (^(v2))
                                    for fi104 := 0; fi104 < 4; fi104++ {
                                        var v105 string = ("a" + ("foo bar" + "\n\t"))
                                        fmt.Println(v3, v105, fi22)
                                        _ = fi104
                                        _ = v105
                                    }
                                    switch 127 {
                                    case 127:
                                        switch ("" + ("a" + ("foo bar" + ("hello" + "a")))) {
                                        case "\n\t":
                                            switch (((("" + ("" + (("foo bar" + "foo bar") + (("a" + "\n\t") + "a")))) + "\n\t") + "") + (("a" + "a") + "hello")) {
                                            case "a":
                                                v106 := GenS3{F0: (fi22 - 666)}
                                                _ = v106
                                            case "":
                                                var v107 float32 = -2.25
                                                _ = v107
                                            default:
                                                if (true && true) {
                                                    switch (fi22 & 32768) {
                                                    case -9223372036854775808:
                                                        v108 := GenS3{F0: fi22}
                                                        _ = v108
                                                    case 42:
                                                        var v109 uint = 2
                                                        _ = v109
                                                    default:
                                                        v110 := GenS5{F0: uint8(v12), F1: int(v4), F2: 0}
                                                        _ = v110
                                                    }
                                                    for fi111 := 0; fi111 < 5; fi111++ {
                                                        var v112 string = "a"
                                                        v113 := GenS9{F0: v21, F1: 128, F2: (v112 + ("foo bar" + ("" + v112)))}
                                                        v3 += (((v3 / (v3 / (v3 + 3.14159))) + 1.5) - v3)
                                                        v114 := []int{(fi22 ^ (fi111 * (fi22 - 2147483647))), 1, 42, fi22}
                                                        if (true && true) {
                                                            v114 = append(v114, (-(fi22)))
                                                            var v115 uint = 2
                                                            if (!true) {
                                                                var v116 uint16 = 255
                                                                var v117 int8 = -26
                                                                if len(v114) >= 2 {
                                                                    v114 = v114[1:len(v114)]
                                                                }
                                                                v4 += (v4 >> ((0) & 63))
                                                                v118 := GenS4{F0: (v7 & (v2 * v2)), F1: 2}
                                                                _ = v116
                                                                _ = v117
                                                                _ = v118
                                                            } else {
                                                                if false {
                                                                    if (194 >= 2) {
                                                                        v4 += (v4 / ((v4) | 1))
                                                                        v114 = append(v114, 127)
                                                                        if (true || true) {
                                                                            v2 += (^(v2))
                                                                        } else {
                                                                            v114 = append(v114, fi22)
                                                                            v119 := GenS10{F0: 144, F1: 1, F2: v112}
                                                                            _ = v119
                                                                        }
                                                                    } else {
                                                                        v120 := GenS11{F0: 32}
                                                                        if false {
                                                                            fmt.Println(v1)
                                                                            _ = len(v114)
                                                                            switch 0 {
                                                                            case 2147483647:
                                                                                v121 := GenS7{F0: v4, F1: v115, F2: 65535}
                                                                                _ = v121
                                                                            case -858:
                                                                                fmt.Println(fi22, v4)
                                                                            case 890:
                                                                                v7 = v7
                                                                            default:
                                                                                if (!(!(4294967295 > fi111))) {
                                                                                    fmt.Println(v21)
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v122 float64 = v3
                                                                            if (!(42 != 53)) {
                                                                                v114 = append(v114, fi22)
                                                                                _ = len(v114)
                                                                                v123 := GenS6{F0: (v7 - v2)}
                                                                                v124 := GenS12{F0: ("" + v112), F1: (-(fi22)), F2: true}
                                                                                _ = v123
                                                                                _ = v124
                                                                            }
                                                                            if false {
                                                                                var v125 bool = (!false)
                                                                                if (!v125) {
                                                                                    var v126 float32 = 1.5
                                                                                    if (42 == 65535) {
                                                                                        v2 = v2
                                                                                        v12 = (v12 / ((v12) | 1))
                                                                                        var v127 int = 1
                                                                                        _ = v127
                                                                                    } else {
                                                                                        if len(v114) > 0 {
                                                                                            v114[0] = (-(fi22))
                                                                                        }
                                                                                        fmt.Println(fi111)
                                                                                        var v128 int32 = v1
                                                                                        _ = v128
                                                                                    }
                                                                                    v129 := GenS13{F0: 2147483648}
                                                                                    _ = v126
                                                                                    _ = v129
                                                                                }
                                                                                if (v1 == int32(fi22)) {
                                                                                    if v125 {
                                                                                        fmt.Println(v4, v2)
                                                                                        if v125 {
                                                                                            v12 = v12
                                                                                            v122 = v3
                                                                                        }
                                                                                        var v130 string = "foo bar"
                                                                                        _ = v130
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(fi22)
                                                                                }
                                                                                v131 := GenS2{F0: v4}
                                                                                _ = v125
                                                                                _ = v131
                                                                            } else {
                                                                                _ = len(v114)
                                                                                switch fi111 {
                                                                                case 127:
                                                                                    if false {
                                                                                        if (!true) {
                                                                                            fmt.Println(v122, v115)
                                                                                            if false {
                                                                                                switch (-(fi22)) {
                                                                                                case -9223372036854775808:
                                                                                                    var v132 int64 = (^(v21))
                                                                                                    _ = v132
                                                                                                }
                                                                                            } else {
                                                                                                if len(v114) >= 2 {
                                                                                                    v114 = v114[1:len(v114)]
                                                                                                }
                                                                                                switch (fi22 * 128) {
                                                                                                case 2147483647:
                                                                                                    var v133 uint = 90
                                                                                                    _ = v133
                                                                                                case 14:
                                                                                                    v1 += v1
                                                                                                default:
                                                                                                    switch (fi111 / ((fi111) | 1)) {
                                                                                                    case 9223372036854775807:
                                                                                                        v134 := GenS5{F0: 1, F1: (fi22 * -314), F2: 0}
                                                                                                        _ = v134
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(fi22, v115, v112)
                                                                                            }
                                                                                            fmt.Println(v12, v122)
                                                                                        }
                                                                                        v7 = 2
                                                                                    } else {
                                                                                        switch (-(fi111)) {
                                                                                        case 2147483647:
                                                                                            v112 = (v112 + "hello")
                                                                                        case 4294967295:
                                                                                            if true {
                                                                                                fmt.Println(v12, v122)
                                                                                                if len(v114) >= 2 {
                                                                                                    v114 = v114[1:len(v114)]
                                                                                                }
                                                                                                switch v112 {
                                                                                                case "hello":
                                                                                                    _ = len(v114)
                                                                                                default:
                                                                                                    fmt.Println(v112)
                                                                                                }
                                                                                                v135 := GenS14{F0: true, F1: 1}
                                                                                                if len(v114) > 0 {
                                                                                                    v114[0] = 42
                                                                                                }
                                                                                                _ = v135
                                                                                            } else {
                                                                                                v136 := GenS14{F0: ((true && true) && true), F1: 65535}
                                                                                                _ = v136
                                                                                            }
                                                                                        default:
                                                                                            v122 = (1.5 - 1.5)
                                                                                        }
                                                                                        if len(v114) > 0 {
                                                                                            v114[0] = fi22
                                                                                        }
                                                                                    }
                                                                                case 32767:
                                                                                    v137 := GenS8{F0: v112, F1: 232, F2: v1}
                                                                                    _ = v137
                                                                                default:
                                                                                    if len(v114) >= 2 {
                                                                                        v114 = v114[1:len(v114)]
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v122
                                                                        }
                                                                        fmt.Println(v12, v12, v4)
                                                                        if false {
                                                                            if false {
                                                                                v4 += int16(v21)
                                                                                v12 -= 805
                                                                                fmt.Println(v115, v12)
                                                                            }
                                                                            if (false || (!(!false))) {
                                                                                var v138 int16 = v4
                                                                                v139 := GenS15{F0: 245, F1: (-(v21)), F2: 2147483648}
                                                                                var v140 uint = (v115 | (v115 % (((^(v115))) | 1)))
                                                                                fmt.Println(v12, v12)
                                                                                _ = v138
                                                                                _ = v139
                                                                                _ = v140
                                                                            }
                                                                        }
                                                                        var v141 int = (fi111 + (fi111 + (-(fi111))))
                                                                        fmt.Println(fi111)
                                                                        _ = v120
                                                                        _ = v141
                                                                    }
                                                                }
                                                                var v142 uint64 = (^(v2))
                                                                v143 := GenS3{F0: fi111}
                                                                if (false && false) {
                                                                    v21 += -1
                                                                    if (!false) {
                                                                        var v144 uint8 = 2
                                                                        switch "a" {
                                                                        case "":
                                                                            var v145 uint64 = 0
                                                                            _ = v145
                                                                        default:
                                                                            switch ((("" + v112) + (v112 + (v112 + "foo bar"))) + ("a" + v112)) {
                                                                            case "foo bar":
                                                                                var v146 int8 = 1
                                                                                _ = v146
                                                                            case "a":
                                                                                if false {
                                                                                    if len(v114) > 0 {
                                                                                        v114[0] = (-(fi111))
                                                                                    }
                                                                                    v4 = int16(v12)
                                                                                    var v147 float64 = float64(-2.25)
                                                                                    _ = v147
                                                                                } else {
                                                                                    var v148 uint8 = v144
                                                                                    if len(v114) >= 2 {
                                                                                        v114 = v114[1:len(v114)]
                                                                                    }
                                                                                    switch fi111 {
                                                                                    case 42:
                                                                                        if true {
                                                                                            v7 += v7
                                                                                        }
                                                                                    case 1:
                                                                                        _ = len(v114)
                                                                                    default:
                                                                                        v149 := GenS16{F0: ("\n\t" + (v112 + v112)), F1: (v142 < v7)}
                                                                                        _ = v149
                                                                                    }
                                                                                    _ = v148
                                                                                }
                                                                            case "\n\t":
                                                                                v150 := GenS17{F0: 67}
                                                                                _ = v150
                                                                            default:
                                                                                if false {
                                                                                    var v151 bool = (!true)
                                                                                    _ = v151
                                                                                } else {
                                                                                    if len(v114) > 0 {
                                                                                        v114[0] = fi22
                                                                                    }
                                                                                    var v152 int8 = int8(v21)
                                                                                    if (true && false) {
                                                                                        if len(v114) >= 2 {
                                                                                            v114 = v114[1:len(v114)]
                                                                                        }
                                                                                        v115 = (v115 / ((v115) | 1))
                                                                                    } else {
                                                                                        v152 = (v152 | 62)
                                                                                    }
                                                                                    var v153 float64 = -2.25
                                                                                    var v154 uint64 = 42
                                                                                    _ = v152
                                                                                    _ = v153
                                                                                    _ = v154
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v144
                                                                    } else {
                                                                        fmt.Println(v7, v7, v21)
                                                                        var v155 int32 = -1
                                                                        switch ((v112 + ("" + "a")) + (((v112 + "hello") + v112) + v112)) {
                                                                        case "\n\t":
                                                                            var v156 uint16 = (^(v12))
                                                                            _ = v156
                                                                        case "a":
                                                                            var v157 string = v112
                                                                            _ = v157
                                                                        }
                                                                        _ = v155
                                                                    }
                                                                    var v158 int32 = v1
                                                                    _ = v158
                                                                }
                                                                switch (fi111 ^ fi22) {
                                                                case -904:
                                                                    var v159 uint8 = 255
                                                                    _ = v159
                                                                }
                                                                _ = v142
                                                                _ = v143
                                                            }
                                                            _ = v115
                                                        }
                                                        _ = fi111
                                                        _ = v112
                                                        _ = v113
                                                        _ = v114
                                                    }
                                                    v160 := []int{-1, fi22, fi22, (fi22 & (fi22 * (fi22 ^ int(v12))))}
                                                    _ = v160
                                                }
                                            }
                                        case "foo bar":
                                            v161 := []int{(fi22 * fi22), fi22, (fi22 * 19), fi22}
                                            _ = v161
                                        default:
                                            for fi162 := 0; fi162 < 18; fi162++ {
                                                var v163 uint = 65535
                                                _ = fi162
                                                _ = v163
                                            }
                                        }
                                    default:
                                        var v164 float32 = 0.0
                                        _ = v164
                                    }
                                }
                            }
                            var v165 uint8 = 42
                            _ = v165
                        }
                    }
                    for fi166 := 0; fi166 < 20; fi166++ {
                        switch (fi22 + (fi166 & 32768)) {
                        case -1:
                            fmt.Println(v21, v7)
                        case 0:
                            if true {
                                if (((v3 + (((float64(float32(v3)) + 3.14159) + v3) + (v3 / v3))) >= v3) || true) {
                                    v12 += (v12 << ((1) & 63))
                                    if ((v3 + (1e10 / 1.5)) < v3) {
                                        v21 = v21
                                        v7 += (v7 - (v2 & (v2 - v7)))
                                        v12 -= v12
                                    } else {
                                        if true {
                                            if (!(-2147483648 != (-(v1)))) {
                                                v3 = (3.14159 + -2.25)
                                            } else {
                                                v167 := []int{fi166, fi166, (^(fi166))}
                                                _ = v167
                                            }
                                            var v168 int8 = 1
                                            _ = v168
                                        } else {
                                            v169 := []int{fi22, (fi22 + fi22)}
                                            var v170 string = "foo bar"
                                            _ = v169
                                            _ = v170
                                        }
                                        var v171 bool = (("hello" + ("foo bar" + "a")) != ("a" + ""))
                                        switch (fi22 % (((fi22 ^ fi22)) | 1)) {
                                        case 9223372036854775807:
                                            var v172 bool = false
                                            _ = v172
                                        case 42:
                                            v12 = v12
                                        default:
                                            switch fi166 {
                                            case -1:
                                                v4 = (^(v4))
                                            case 42:
                                                var v173 bool = (0 < v12)
                                                _ = v173
                                            case 748:
                                                v174 := []int{36, -1, fi166}
                                                _ = v174
                                            }
                                        }
                                        v3 = (v3 / (v3 / v3))
                                        _ = v171
                                    }
                                } else {
                                    v7 -= v7
                                    v175 := GenS18{F0: (1 >= v1), F1: 42, F2: 1}
                                    var v176 uint8 = 127
                                    var v177 uint = 1
                                    _ = v175
                                    _ = v176
                                    _ = v177
                                }
                                v178 := []int{}
                                v7 -= uint64(v12)
                                _ = v178
                            }
                        }
                        var v179 int8 = 42
                        var v180 int32 = 42
                        _ = fi166
                        _ = v179
                        _ = v180
                    }
                } else {
                    var v181 int = -71
                    _ = v181
                }
                _ = fi22
            }
            _ = v21
        }
    default:
        v182 := []int{-9223372036854775808, 9223372036854775807, -795}
        _ = v182
    }
    if true {
        var v183 uint16 = (v12 ^ (v12 ^ (^(v12))))
        _ = v183
    }
    v184 := GenS19{F0: 129, F1: 9223372036854775807}
    v185 := GenS12{F0: "\n\t", F1: 32768, F2: ((((false || true) && true) || (-2.25 >= (0.0 + (float32(float64(-2.25)) - float32(((float64(0.1) + (v3 / float64(float32(3.14159)))) - 0.1)))))) || ((4294967295 == 0) || (65535 <= 4294967295)))}
    v186 := []int{-455, 32768, 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v186:", len(v186), v186)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v184
    _ = v185
    _ = v186
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
