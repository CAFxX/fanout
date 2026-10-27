package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int8; F2 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 int8; F1 bool; F2 string }
type GenS9 struct { F0 int64; F1 uint; F2 int8 }
type GenS10 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS11 struct { F0 uint; F1 uint8 }
type GenS12 struct { F0 int8; F1 int16 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 uint8; F1 uint64 }
type GenS17 struct { F0 int32; F1 uint16 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 int16; F1 uint8 }
type GenS20 struct { F0 int8; F1 uint8 }
type GenS21 struct { F0 uint64; F1 int; F2 int32 }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 bool; F1 int }
type GenS24 struct { F0 bool; F1 uint64; F2 int32 }
func genFn0(p0 uint32, p1 uint8, p2 uint64) int32 {
    fmt.Println(p0, p1, p0)
    v7 := GenS1{F0: 603}
    var v8 uint32 = (p0 << ((65535) & 63))
    v9 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -190
}
func genFn1(p0 int16) int32 {
    var v10 int = 32767
    if ((4294967295 >= 882) && false) {
        v11 := GenS3{F0: 100, F1: 100, F2: 218}
        var v12 float32 = 0.1
        switch (v10 - 32767) {
        case 42:
            switch 533 {
            case 724:
                v13 := []int{}
                _ = v13
            case 866:
                var v14 string = ("" + "\n\t")
                _ = v14
            case 4294967295:
                fmt.Println(p0)
            default:
                var v15 int32 = -1
                _ = v15
            }
        }
        _ = v11
        _ = v12
    } else {
        v16 := GenS1{F0: -497}
        if false {
            var v17 float64 = 1.5
            v18 := GenS4{F0: 315}
            for fi19 := 0; fi19 < 14; fi19++ {
                v20 := []int{4294967295, fi19, -350, 0}
                _ = fi19
                _ = v20
            }
            _ = v17
            _ = v18
        } else {
            fmt.Println(p0)
        }
        _ = v16
    }
    fmt.Println(v10, v10)
    fmt.Println(p0)
    v10 = (v10 + (v10 >> ((18446744073709551615) & 63)))
    _ = p0
    _ = v10
    return -2147483648
}

func main() {
    v1 := []uint32{2, 53, 2147483647, 65535}
    v2 := []uint32{629, 1, 4294967295}
    v3 := make(map[string]string)
    var v4 int16 = 808
    var v5 float32 = 3.14159
    v6 := make(map[int]bool)
    if false {
        var v21 uint8 = 1
        v4 = (-(v4))
        if (false || (true && true)) {
            for fi22 := 0; fi22 < 13; fi22++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi22
            }
        }
        _ = v21
    }
    v23 := GenS4{F0: -1}
    var v24 int = 4294967295
    v4 += v4
    var v25 float32 = v5
    var v26 int = (v24 >> ((735) & 63))
    for fi27 := 0; fi27 < 12; fi27++ {
        fmt.Println(v4, v25)
        for fi28 := 0; fi28 < 9; fi28++ {
            if false {
                switch "hello" {
                case "\n\t":
                    switch "a" {
                    case "foo bar":
                        switch 32768 {
                        case 893:
                            var v29 uint8 = 128
                            _ = v29
                        default:
                            v30 := GenS4{F0: -2147483648}
                            _ = v30
                        }
                    case "":
                        v31 := GenS5{F0: 42, F1: 42, F2: 4294967295}
                        _ = v31
                    }
                case "a":
                    switch int(v4) {
                    case 4294967295:
                        v26 = v26
                    default:
                        fmt.Println(fi27, v24, v5)
                    }
                default:
                    v32 := GenS5{F0: -684, F1: 4294967295, F2: 18446744073709551615}
                    _ = v32
                }
                v33 := GenS6{F0: 42}
                if (636 != -1) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch "\n\t" {
                    case "":
                        var v34 bool = false
                        _ = v34
                    }
                    if false {
                        v35 := GenS4{F0: -1}
                        v4 = v4
                        fmt.Println(v5, fi28)
                        _ = v35
                    }
                } else {
                    if (-1 >= -829) {
                        v4 = v4
                        fmt.Println(v24, v5, v5)
                    }
                    if true {
                        var v36 uint8 = 0
                        _ = v36
                    } else {
                        var v37 float32 = 3.14159
                        if true {
                            var v38 uint8 = 2
                            var v39 float64 = 1e10
                            switch "hello" {
                            case "a":
                                if (-201 != 2147483647) {
                                    switch (("hello" + ("foo bar" + ("foo bar" + ("\n\t" + "foo bar")))) + "\n\t") {
                                    case "\n\t":
                                        v40 := GenS3{F0: -128, F1: 42, F2: v38}
                                        _ = v40
                                    default:
                                        if (98 != 0) {
                                            switch fi28 {
                                            case -1:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            default:
                                                v39 -= v39
                                            }
                                        }
                                    }
                                }
                            }
                            _ = v38
                            _ = v39
                        } else {
                            switch "foo bar" {
                            case "hello":
                                v41 := GenS1{F0: 1}
                                _ = v41
                            }
                            if false {
                                v42 := GenS7{F0: (518 <= 4294967295), F1: 255}
                                if (18446744073709551615 >= 18446744073709551615) {
                                    v24 = fi28
                                    v2 = append(v2, 65535)
                                }
                                _ = v42
                            }
                        }
                        v43 := GenS8{F0: 86, F1: false, F2: "\n\t"}
                        _ = v37
                        _ = v43
                    }
                }
                var v44 uint8 = 42
                v45 := GenS9{F0: -1, F1: 219, F2: 23}
                var v46 uint64 = 841
                _ = v33
                _ = v44
                _ = v45
                _ = v46
            }
            var v47 int16 = v4
            if false {
                v48 := GenS10{F0: 304, F1: 2147483648, F2: 339}
                v26 -= (v24 % ((fi28) | 1))
                _ = v48
            }
            _ = fi28
            _ = v47
        }
        for fi49 := 0; fi49 < 14; fi49++ {
            fmt.Println(v26, v5, v26)
            if false {
                fmt.Println(v26)
            }
            _ = len(v3)
            fmt.Println(fi49, v5, v5)
            v26 -= 42
            _ = fi49
        }
        _ = fi27
    }
    var v50 int32 = 0
    switch (v26 % (((-(v26))) | 1)) {
    case -1:
        switch ("" + (("" + "hello") + "\n\t")) {
        case "foo bar":
            for fi51 := 0; fi51 < 8; fi51++ {
                if false {
                    switch (("" + ("foo bar" + ("\n\t" + "hello"))) + "foo bar") {
                    case "foo bar":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "":
                        for fi52 := 0; fi52 < 14; fi52++ {
                            var v53 int16 = (v4 | v4)
                            var v54 uint16 = 833
                            var v55 int8 = 16
                            var v56 int16 = v53
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi52
                            _ = v53
                            _ = v54
                            _ = v55
                            _ = v56
                        }
                    case "\n\t":
                        switch "a" {
                        case "":
                            v57 := GenS8{F0: 42, F1: false, F2: "foo bar"}
                            _ = v57
                        case "hello":
                            if (42 != 42) {
                                for fi58 := 0; fi58 < 5; fi58++ {
                                    v24 -= v26
                                    v59 := GenS11{F0: 0, F1: 127}
                                    if false {
                                        v4 = (v4 & (v4 % (((-(v4))) | 1)))
                                        v25 = ((v5 - v25) + (v25 - v25))
                                    }
                                    var v60 uint8 = 2
                                    v2 = append(v2, 65535)
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                }
                            } else {
                                if ((!(-2.25 > v5)) || true) {
                                    if true {
                                        switch -2147483648 {
                                        case 302:
                                            fmt.Println(v5, v5, v50)
                                        default:
                                            for fi61 := 0; fi61 < 8; fi61++ {
                                                v62 := v3[("" + "\n\t")]
                                                if true {
                                                    switch v62 {
                                                    case "\n\t":
                                                        if true {
                                                            if ((0.1 - v5) == v5) {
                                                                var v63 float32 = (-2.25 * v25)
                                                                var v64 uint32 = 42
                                                                _ = v63
                                                                _ = v64
                                                            } else {
                                                                v4 -= v4
                                                            }
                                                            v65 := GenS7{F0: (!(((float32(((-2.25 / (0.0 - (1.5 / (1.5 - float64((v25 * v5)))))) + 3.14159)) > 0.1) && true) && false)), F1: 32767}
                                                            var v66 float64 = 0.0
                                                            v50 = -1
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                switch (v62 + v62) {
                                                                case "foo bar":
                                                                    var v67 string = v62
                                                                    _ = v67
                                                                }
                                                            case "a":
                                                                v2 = append(v2, 861)
                                                            case "hello":
                                                                fmt.Println(v66)
                                                            default:
                                                                fmt.Println(v5)
                                                            }
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                    case "a":
                                                        var v68 uint64 = 18446744073709551615
                                                        _ = v68
                                                    case "hello":
                                                        fmt.Println(v24, v25)
                                                    }
                                                    var v69 string = ("a" + (("\n\t" + v62) + ((v62 + v62) + (v62 + "foo bar"))))
                                                    _ = v69
                                                }
                                                var v70 bool = false
                                                _ = fi61
                                                _ = v62
                                                _ = v70
                                            }
                                        }
                                        switch (v24 ^ (fi51 | fi51)) {
                                        case -656:
                                            v71 := v6[896]
                                            _ = v71
                                        case -446:
                                            for fi72 := 0; fi72 < 7; fi72++ {
                                                var v73 int64 = 0
                                                switch v24 {
                                                case 622:
                                                    switch fi72 {
                                                    case 140:
                                                        if false {
                                                            var v74 uint16 = 32767
                                                            v50 = (v50 * v50)
                                                            _ = v74
                                                        }
                                                    }
                                                case 127:
                                                    if ((3.14159 < ((v5 - (v5 + (float32(float64((0.0 - 3.14159))) + 1e10))) * (v5 - v25))) && false) {
                                                        var v75 int8 = -12
                                                        _ = v75
                                                    } else {
                                                        if true {
                                                            fmt.Println(v26, v26)
                                                            fmt.Println(v4, v73, fi72)
                                                            _ = len(v6)
                                                        }
                                                        var v76 uint = 1
                                                        var v77 uint16 = 255
                                                        v24 = v26
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                case -827:
                                                    if len(v1) > 0 {
                                                        v1[0] = 466
                                                    }
                                                default:
                                                    v3["hello"] = "\n\t"
                                                }
                                                _ = fi72
                                                _ = v73
                                            }
                                        case -9223372036854775808:
                                            var v78 float64 = 1.5
                                            _ = v78
                                        }
                                        for fi79 := 0; fi79 < 10; fi79++ {
                                            if true {
                                                if true {
                                                    if true {
                                                        var v80 string = "a"
                                                        v81 := GenS2{F0: int64(v50)}
                                                        _ = v80
                                                        _ = v81
                                                    }
                                                    switch (fi79 % ((fi79) | 1)) {
                                                    case -9223372036854775808:
                                                        if true {
                                                            v82 := GenS11{F0: 780, F1: 126}
                                                            _ = v82
                                                        } else {
                                                            v83 := GenS12{F0: 1, F1: v4}
                                                            _ = v83
                                                        }
                                                    case 222:
                                                        v24 = fi51
                                                        var v84 bool = true
                                                        _ = v84
                                                    }
                                                    switch int(v4) {
                                                    case 127:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    case 32768:
                                                        switch "" {
                                                        case "hello":
                                                            var v85 int16 = (v4 / ((v4) | 1))
                                                            switch fi79 {
                                                            case 156:
                                                                if (("hello" + "foo bar") > "\n\t") {
                                                                    var v86 int8 = 122
                                                                    var v87 int = (fi79 << ((0) & 63))
                                                                    switch ("foo bar" + "foo bar") {
                                                                    case "hello":
                                                                        delete(v3, ("a" + ("a" + "a")))
                                                                    default:
                                                                        v26 = v26
                                                                    }
                                                                    _ = v86
                                                                    _ = v87
                                                                } else {
                                                                    var v88 int = -2147483648
                                                                    fmt.Println(v85, v88, v5)
                                                                    fmt.Println(v5, fi51, v25)
                                                                    _ = v88
                                                                }
                                                            case 4294967295:
                                                                v89 := GenS2{F0: 9223372036854775807}
                                                                _ = v89
                                                            }
                                                            fmt.Println(v25, v4, fi51)
                                                            _ = v85
                                                        case "a":
                                                            var v90 uint = 4294967295
                                                            _ = v90
                                                        default:
                                                            if ((264 > 620) || (0 > -9223372036854775808)) {
                                                                v91 := GenS1{F0: -903}
                                                                v5 += (((1e10 / -2.25) * 0.0) * (-2.25 - ((0.0 * (v25 / 3.14159)) / 1e10)))
                                                                _ = v91
                                                            }
                                                        }
                                                    case 32767:
                                                        v92 := GenS13{F0: genFn1(v4)}
                                                        _ = v92
                                                    default:
                                                        var v93 uint = 18446744073709551615
                                                        _ = v93
                                                    }
                                                } else {
                                                    v94 := v3["\n\t"]
                                                    _ = v94
                                                }
                                            }
                                            _ = len(v6)
                                            fmt.Println(v26, v5, v25)
                                            if false {
                                                fmt.Println(v26, v5)
                                                fmt.Println(v25, v4, v25)
                                                if true {
                                                    var v95 int32 = (v50 + (v50 % ((v50) | 1)))
                                                    _ = v95
                                                }
                                            }
                                            v96 := GenS7{F0: true, F1: 1}
                                            _ = fi79
                                            _ = v96
                                        }
                                    } else {
                                        v97 := GenS14{F0: -961}
                                        v25 = -2.25
                                        v26 = v26
                                        _ = v97
                                    }
                                    switch fi51 {
                                    case -1:
                                        var v98 string = ""
                                        _ = v98
                                    default:
                                        if true {
                                            fmt.Println(v24, v4, v5)
                                        } else {
                                            var v99 float32 = 3.14159
                                            _ = v99
                                        }
                                    }
                                } else {
                                    switch (("a" + (("" + ("hello" + "a")) + (("hello" + "a") + ""))) + "\n\t") {
                                    case "a":
                                        for fi100 := 0; fi100 < 5; fi100++ {
                                            var v101 uint = 2
                                            switch (^(v24)) {
                                            case 9223372036854775807:
                                                fmt.Println(v25)
                                            default:
                                                if (!(((("hello" + "") + ("a" + "a")) != "\n\t") && (460 <= 0))) {
                                                    if (true || false) {
                                                        _ = len(v1)
                                                        var v102 int32 = (v50 % (((v50 / ((754) | 1))) | 1))
                                                        _ = v102
                                                    } else {
                                                        switch "hello" {
                                                        case "a":
                                                            if false {
                                                                v50 = (-(v50))
                                                                switch (-(v26)) {
                                                                case -2147483648:
                                                                    _ = len(v2)
                                                                case 42:
                                                                    v4 -= v4
                                                                }
                                                                var v103 int16 = -816
                                                                switch (fi51 ^ (v26 - fi51)) {
                                                                case 127:
                                                                    v25 = 1e10
                                                                case 9223372036854775807:
                                                                    fmt.Println(v4, v24, v5)
                                                                case -348:
                                                                    var v104 uint8 = 0
                                                                    _ = v104
                                                                default:
                                                                    var v105 uint16 = 0
                                                                    _ = v105
                                                                }
                                                                _ = v103
                                                            } else {
                                                                switch fi51 {
                                                                case 179:
                                                                    switch ((("hello" + "\n\t") + "") + ("a" + (("hello" + "\n\t") + "a"))) {
                                                                    case "\n\t":
                                                                        var v106 bool = (v50 < v50)
                                                                        _ = v106
                                                                    case "foo bar":
                                                                        switch (fi51 % ((int(v4)) | 1)) {
                                                                        case 760:
                                                                            switch ("a" + (("foo bar" + ((("foo bar" + "hello") + ((("" + ("a" + "a")) + "hello") + ("" + "\n\t"))) + "a")) + "")) {
                                                                            case "hello":
                                                                                v4 = (v4 - v4)
                                                                            }
                                                                        case 127:
                                                                            v50 = (v50 << (((v101 * v101)) & 63))
                                                                        }
                                                                    case "a":
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        if (true && (0.0 <= 1e10)) {
                                                                            fmt.Println(fi100, fi51)
                                                                            if true {
                                                                                switch v24 {
                                                                                case 2147483647:
                                                                                    var v107 string = "a"
                                                                                    var v108 float64 = 3.14159
                                                                                    fmt.Println(fi51)
                                                                                    if false {
                                                                                        switch v107 {
                                                                                        case "\n\t":
                                                                                            v107 = v107
                                                                                        case "foo bar":
                                                                                            var v109 bool = (true && false)
                                                                                            _ = v109
                                                                                        }
                                                                                        var v110 bool = true
                                                                                        _ = v110
                                                                                    }
                                                                                    v5 = 1e10
                                                                                    _ = v107
                                                                                    _ = v108
                                                                                case 627:
                                                                                    v111 := GenS3{F0: -23, F1: 127, F2: 83}
                                                                                    _ = v111
                                                                                }
                                                                                fmt.Println(v50, fi51, fi51)
                                                                                var v112 int32 = -89
                                                                                _ = v112
                                                                            }
                                                                            fmt.Println(v101, v24)
                                                                            var v113 int8 = int8(v24)
                                                                            _ = v113
                                                                        }
                                                                    default:
                                                                        v2 = append(v2, 1)
                                                                    }
                                                                case 2147483647:
                                                                    v114 := GenS8{F0: 100, F1: false, F2: ""}
                                                                    _ = v114
                                                                case -950:
                                                                    v115 := GenS15{F0: (^(v101))}
                                                                    _ = v115
                                                                default:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 2147483647
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            v25 += v5
                                                        }
                                                        v6[(fi100 ^ v26)] = true
                                                    }
                                                    switch (v24 + v26) {
                                                    case -1:
                                                        v101 += v101
                                                    default:
                                                        var v116 uint64 = 4294967295
                                                        _ = v116
                                                    }
                                                    if ((true && ((0.0 >= (3.14159 + 0.1)) && (false || false))) || true) {
                                                        v117 := GenS16{F0: 128, F1: 918}
                                                        _ = v117
                                                    }
                                                } else {
                                                    _ = len(v1)
                                                }
                                            }
                                            if true {
                                                v2 = append(v2, 4294967295)
                                                v26 -= (v26 ^ (v24 * v26))
                                                v118 := GenS17{F0: v50, F1: 846}
                                                _ = v118
                                            }
                                            _ = fi100
                                            _ = v101
                                        }
                                    case "\n\t":
                                        switch (v26 - (fi51 * v26)) {
                                        case 283:
                                            if (("\n\t" + ("hello" + ("\n\t" + ((("" + "foo bar") + "\n\t") + ("" + ("" + ("a" + ("a" + "a")))))))) == "a") {
                                                for fi119 := 0; fi119 < 13; fi119++ {
                                                    var v120 int8 = 118
                                                    v121 := v3["\n\t"]
                                                    _ = fi119
                                                    _ = v120
                                                    _ = v121
                                                }
                                                for fi122 := 0; fi122 < 4; fi122++ {
                                                    if false {
                                                        v4 = v4
                                                    } else {
                                                        v4 = 42
                                                        v25 -= 1e10
                                                    }
                                                    _ = fi122
                                                }
                                                if true {
                                                    v50 += v50
                                                }
                                            } else {
                                                for fi123 := 0; fi123 < 2; fi123++ {
                                                    switch ("a" + "hello") {
                                                    case "a":
                                                        switch 4294967295 {
                                                        case -599:
                                                            v124 := GenS18{F0: 793}
                                                            _ = v124
                                                        case 128:
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v25 = v5
                                                            }
                                                        default:
                                                            v125 := GenS17{F0: v50, F1: 0}
                                                            _ = v125
                                                        }
                                                    default:
                                                        v24 = (fi51 % ((v26) | 1))
                                                    }
                                                    _ = fi123
                                                }
                                            }
                                        case -2147483648:
                                            for fi126 := 0; fi126 < 3; fi126++ {
                                                var v127 int32 = v50
                                                v128 := GenS19{F0: v4, F1: 198}
                                                v5 += (float32(0.1) / (v25 / float32(3.14159)))
                                                if len(v1) > 0 {
                                                    v1[0] = 2
                                                }
                                                v50 += 2147483647
                                                _ = fi126
                                                _ = v127
                                                _ = v128
                                            }
                                        }
                                    }
                                    switch "" {
                                    case "hello":
                                        var v129 int16 = (v4 | 0)
                                        _ = v129
                                    case "\n\t":
                                        for fi130 := 0; fi130 < 6; fi130++ {
                                            var v131 string = "a"
                                            var v132 float64 = ((float64(0.1) / 1.0) * -2.25)
                                            _ = fi130
                                            _ = v131
                                            _ = v132
                                        }
                                    case "a":
                                        v133 := GenS20{F0: -115, F1: 0}
                                        _ = v133
                                    }
                                }
                            }
                        default:
                            if true {
                                for fi134 := 0; fi134 < 20; fi134++ {
                                    if false {
                                        var v135 string = "hello"
                                        v136 := GenS7{F0: false, F1: 691}
                                        v50 -= (v50 / (((v50 - (v50 / ((v50) | 1)))) | 1))
                                        _ = v135
                                        _ = v136
                                    }
                                    _ = fi134
                                }
                            } else {
                                v4 = v4
                                v5 += (v5 * v5)
                                v25 -= v25
                            }
                        }
                    default:
                        fmt.Println(v26)
                    }
                    for fi137 := 0; fi137 < 4; fi137++ {
                        if len(v1) > 0 {
                            v1[0] = 360
                        }
                        v138 := GenS21{F0: 839, F1: fi51, F2: 42}
                        _ = fi137
                        _ = v138
                    }
                }
                _ = len(v3)
                v139 := v3["foo bar"]
                _ = fi51
                _ = v139
            }
        case "hello":
            var v140 int = (v24 << ((65535) & 63))
            _ = v140
        }
    }
    var v141 uint8 = 121
    if true {
        v2 = append(v2, 42)
        if (((0.1 / 1.0) != -2.25) && (false && true)) {
            v25 += v25
            if (v141 < 42) {
                for fi142 := 0; fi142 < 4; fi142++ {
                    switch (v24 << ((65535) & 63)) {
                    case 1:
                        var v143 bool = ((^(v24)) == (v26 & 0))
                        _ = v143
                    }
                    if (true || false) {
                        for fi144 := 0; fi144 < 8; fi144++ {
                            v145 := GenS21{F0: 4294967295, F1: (v26 & v26), F2: (v50 / ((genFn1((v4 << ((65535) & 63)))) | 1))}
                            var v146 string = "a"
                            var v147 float64 = (3.14159 - (float64((float32((1e10 + 3.14159)) / v25)) + 1e10))
                            _ = fi144
                            _ = v145
                            _ = v146
                            _ = v147
                        }
                    } else {
                        var v148 bool = (!true)
                        var v149 uint8 = 127
                        if (!(!true)) {
                            v24 -= v26
                            if false {
                                v25 -= v25
                                fmt.Println(v141, v141, v24)
                                switch fi142 {
                                case 127:
                                    for fi150 := 0; fi150 < 18; fi150++ {
                                        v5 = ((float32(3.14159) - (((v25 - (((float32((3.14159 - 0.1)) / v5) / (v25 + ((v25 / (v5 + 0.1)) / 3.14159))) - float32(1e10))) + v5) + ((3.14159 / ((1e10 / 1.0) - float32((1.5 - 1.5)))) * (float32(-2.25) * (v25 - 1.5))))) - v5)
                                        var v151 uint16 = 24
                                        _ = fi150
                                        _ = v151
                                    }
                                case -232:
                                    fmt.Println(v149, v25)
                                case 703:
                                    if false {
                                        if ((v5 * 3.14159) < -2.25) {
                                            var v152 int = (v24 ^ 9223372036854775807)
                                            switch -9223372036854775808 {
                                            case 450:
                                                if false {
                                                    var v153 int64 = -1
                                                    for fi154 := 0; fi154 < 20; fi154++ {
                                                        fmt.Println(v24, v152, fi142)
                                                        v155 := GenS18{F0: 1}
                                                        v5 = (0.0 - v5)
                                                        _ = fi154
                                                        _ = v155
                                                    }
                                                    _ = v153
                                                } else {
                                                    var v156 string = "\n\t"
                                                    v156 = v156
                                                    _ = v156
                                                }
                                            case 32767:
                                                switch (("hello" + "hello") + "hello") {
                                                case "a":
                                                    if v148 {
                                                        fmt.Println(v148, v152, v149)
                                                        v141 = (v141 % (((v141 & (v141 - (^(v149))))) | 1))
                                                        v5 += (v25 + ((((float32(0.1) - (v5 / 1.5)) - 1e10) / float32((1e10 + 1.5))) - (v5 * (v25 * v5))))
                                                    }
                                                default:
                                                    for fi157 := 0; fi157 < 15; fi157++ {
                                                        if v148 {
                                                            var v158 int16 = (v4 & int16(v149))
                                                            _ = v158
                                                        } else {
                                                            var v159 int64 = -1
                                                            var v160 string = ""
                                                            _ = v159
                                                            _ = v160
                                                        }
                                                        switch (^(v26)) {
                                                        case 32768:
                                                            if (true && (("\n\t" != ("\n\t" + (("" + "\n\t") + "foo bar"))) && v148)) {
                                                                var v161 bool = false
                                                                v141 = v149
                                                                _ = v161
                                                            }
                                                        default:
                                                            if len(v1) > 0 {
                                                                v1[0] = 589
                                                            }
                                                            var v162 string = ("hello" + "foo bar")
                                                            switch "a" {
                                                            case "":
                                                                var v163 uint64 = 2
                                                                if (!(18 < 100)) {
                                                                    if ((^(v4)) > v4) {
                                                                        fmt.Println(v149)
                                                                    }
                                                                }
                                                                if v148 {
                                                                    switch v162 {
                                                                    case "\n\t":
                                                                        var v164 int8 = -1
                                                                        _ = v164
                                                                    case "foo bar":
                                                                        var v165 string = v162
                                                                        _ = v165
                                                                    case "hello":
                                                                        var v166 int32 = v50
                                                                        _ = v166
                                                                    default:
                                                                        delete(v3, "a")
                                                                    }
                                                                }
                                                                if v148 {
                                                                    if true {
                                                                        if v148 {
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 1
                                                                                }
                                                                            case "hello":
                                                                                fmt.Println(v141, v50, fi142)
                                                                            case "a":
                                                                                v163 += v163
                                                                                v1 = append(v1, 4294967295)
                                                                            default:
                                                                                v4 += 1
                                                                            }
                                                                        }
                                                                        var v167 string = "foo bar"
                                                                        if v148 {
                                                                            switch v162 {
                                                                            case "\n\t":
                                                                                var v168 string = "\n\t"
                                                                                _ = v168
                                                                            default:
                                                                                v169 := v3[v167]
                                                                                _ = v169
                                                                            }
                                                                        }
                                                                        _ = v167
                                                                    }
                                                                } else {
                                                                    v162 = v162
                                                                }
                                                                _ = v163
                                                            case "foo bar":
                                                                var v170 uint32 = 2147483647
                                                                _ = v170
                                                            default:
                                                                v171 := GenS4{F0: (v50 - v50)}
                                                                _ = v171
                                                            }
                                                            var v172 uint16 = 1
                                                            var v173 int64 = 1
                                                            _ = v162
                                                            _ = v172
                                                            _ = v173
                                                        }
                                                        if (("a" + "hello") == "foo bar") {
                                                            v174 := GenS9{F0: -1, F1: 255, F2: 100}
                                                            var v175 string = "\n\t"
                                                            _ = v174
                                                            _ = v175
                                                        } else {
                                                            v4 = 32767
                                                        }
                                                        var v176 uint16 = 0
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v176 += v176
                                                        default:
                                                            switch "hello" {
                                                            case "a":
                                                                switch ("a" + "\n\t") {
                                                                case "foo bar":
                                                                    v5 -= ((v25 / 1.5) - (v25 / 1.0))
                                                                case "\n\t":
                                                                    var v177 uint16 = (v176 & (v176 + v176))
                                                                    _ = v177
                                                                }
                                                            case "foo bar":
                                                                var v178 uint64 = 2
                                                                _ = v178
                                                            default:
                                                                if v148 {
                                                                    switch 42 {
                                                                    case 127:
                                                                        v152 = int(v149)
                                                                    case 32767:
                                                                        if v148 {
                                                                            var v179 bool = true
                                                                            if (9223372036854775807 < 42) {
                                                                                var v180 string = "\n\t"
                                                                                _ = v180
                                                                            } else {
                                                                                var v181 int = v26
                                                                                _ = v181
                                                                            }
                                                                            v3["foo bar"] = "foo bar"
                                                                            v25 += float32(3.14159)
                                                                            if true {
                                                                                var v182 float32 = 1e10
                                                                                v183 := GenS10{F0: 0, F1: 753, F2: 0}
                                                                                _ = v182
                                                                                _ = v183
                                                                            } else {
                                                                                var v184 int = fi157
                                                                                v5 = v25
                                                                                var v185 uint8 = 2
                                                                                v24 += fi142
                                                                                v186 := GenS22{F0: v148}
                                                                                _ = v184
                                                                                _ = v185
                                                                                _ = v186
                                                                            }
                                                                            var v187 uint = 65535
                                                                            var v188 uint = v187
                                                                            _ = v179
                                                                            _ = v187
                                                                            _ = v188
                                                                        }
                                                                        _ = len(v1)
                                                                        v3[""] = "a"
                                                                    default:
                                                                        v189 := v6[(v152 - 4294967295)]
                                                                        _ = v189
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = fi157
                                                        _ = v176
                                                    }
                                                }
                                            case -403:
                                                switch v26 {
                                                case 0:
                                                    v149 -= (v141 | 2)
                                                case 32767:
                                                    var v190 float64 = 1e10
                                                    _ = v190
                                                default:
                                                    fmt.Println(v24, v148)
                                                }
                                            default:
                                                _ = len(v3)
                                            }
                                            for fi191 := 0; fi191 < 12; fi191++ {
                                                v26 = (fi142 << ((18446744073709551615) & 63))
                                                var v192 bool = (v148 || (836 > (-(v4))))
                                                var v193 uint16 = 255
                                                _ = fi191
                                                _ = v192
                                                _ = v193
                                            }
                                            _ = v152
                                        }
                                        fmt.Println(v148, v149, v25)
                                        v194 := GenS23{F0: (!v148), F1: -1}
                                        _ = v194
                                    }
                                }
                            } else {
                                if false {
                                    switch fi142 {
                                    case -9223372036854775808:
                                        fmt.Println(v5, v26, v50)
                                    case 1:
                                        v195 := GenS4{F0: (v50 ^ 448)}
                                        _ = v195
                                    }
                                    delete(v6, (fi142 << ((917) & 63)))
                                    var v196 int = (-(v26))
                                    _ = v196
                                } else {
                                    if (32767 != 42) {
                                        var v197 float32 = v25
                                        v197 += v5
                                        _ = v197
                                    } else {
                                        v198 := GenS9{F0: 1, F1: 1, F2: -1}
                                        _ = v198
                                    }
                                }
                                for fi199 := 0; fi199 < 9; fi199++ {
                                    if len(v1) > 0 {
                                        v1[0] = 481
                                    }
                                    var v200 uint16 = 647
                                    _ = fi199
                                    _ = v200
                                }
                                if len(v1) > 0 {
                                    v1[0] = 662
                                }
                            }
                            fmt.Println(v26, v26)
                        } else {
                            var v201 uint16 = 211
                            _ = v201
                        }
                        _ = v148
                        _ = v149
                    }
                    for fi202 := 0; fi202 < 13; fi202++ {
                        v3["hello"] = (("\n\t" + ("" + "foo bar")) + ("a" + "hello"))
                        v4 -= (-(v4))
                        var v203 uint16 = 341
                        _ = fi202
                        _ = v203
                    }
                    _ = fi142
                }
            }
            v25 = (3.14159 + float32((0.0 / 1e10)))
        }
        switch (v26 << ((995) & 63)) {
        case 0:
            v204 := GenS9{F0: -1, F1: 918, F2: 42}
            _ = v204
        case 32768:
            for fi205 := 0; fi205 < 4; fi205++ {
                var v206 int64 = 42
                _ = fi205
                _ = v206
            }
        case -449:
            var v207 float64 = (0.1 * 3.14159)
            _ = v207
        }
    } else {
        v26 += (v26 | v24)
        for fi208 := 0; fi208 < 10; fi208++ {
            if (false || false) {
                v209 := GenS5{F0: 9223372036854775807, F1: 624, F2: 4294967295}
                v210 := GenS4{F0: (v50 | v50)}
                var v211 int32 = (^(v50))
                _ = v209
                _ = v210
                _ = v211
            } else {
                var v212 int8 = 100
                _ = v212
            }
            switch "foo bar" {
            case "\n\t":
                switch "foo bar" {
                case "":
                    var v213 float32 = 1e10
                    _ = v213
                case "\n\t":
                    fmt.Println(v50, v26, v25)
                default:
                    if true {
                        v26 -= v24
                        fmt.Println(v26, v141, fi208)
                        for fi214 := 0; fi214 < 9; fi214++ {
                            var v215 uint32 = 42
                            _ = fi214
                            _ = v215
                        }
                    } else {
                        switch ("hello" + "hello") {
                        case "a":
                            _ = len(v6)
                        case "\n\t":
                            v216 := GenS3{F0: -22, F1: 42, F2: 255}
                            _ = v216
                        default:
                            v217 := GenS24{F0: (false && (true && false)), F1: 42, F2: 925}
                            _ = v217
                        }
                        v24 += v26
                        v218 := GenS6{F0: -930}
                        _ = v218
                    }
                }
            case "a":
                v50 = (v50 >> ((1) & 63))
            }
            v5 = v5
            _ = fi208
        }
    }
    for fi219 := 0; fi219 < 18; fi219++ {
        var v220 int64 = int64(v141)
        _ = fi219
        _ = v220
    }
    if false {
        v141 += (v141 * (v141 - v141))
    } else {
        var v221 int8 = -1
        _ = v221
    }
    if false {
        var v222 float32 = v25
        for fi223 := 0; fi223 < 19; fi223++ {
            if ((v222 / v222) < (((float32(float64(float32(-2.25))) / (v222 * v5)) - (float32((0.1 - 1e10)) * (v5 / 1.5))) * 0.0)) {
                v50 -= -2147483648
                switch "hello" {
                case "":
                    v224 := GenS20{F0: 100, F1: 1}
                    _ = v224
                case "hello":
                    if (!false) {
                        for fi225 := 0; fi225 < 6; fi225++ {
                            v26 = (v26 >> ((255) & 63))
                            _ = fi225
                        }
                        var v226 float64 = 1.5
                        v6[(v24 / ((2147483647) | 1))] = ((-(fi223)) > fi223)
                        _ = v226
                    }
                default:
                    switch 1 {
                    case 40:
                        switch fi223 {
                        case 729:
                            v227 := GenS10{F0: 848, F1: 0, F2: 0}
                            _ = v227
                        case 201:
                            v228 := GenS18{F0: -679}
                            _ = v228
                        case 9223372036854775807:
                            for fi229 := 0; fi229 < 7; fi229++ {
                                fmt.Println(fi229, v50)
                                v26 = (v26 % ((v26) | 1))
                                _ = fi229
                            }
                        }
                    }
                }
            } else {
                v3[((("a" + "a") + "foo bar") + ("" + "\n\t"))] = ("a" + ("hello" + (("a" + "") + (("a" + "foo bar") + ""))))
            }
            _ = fi223
        }
        _ = v222
    }
    v25 = v25
    var v230 int16 = int16(v50)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v141:", v141)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v230:", v230)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v141
    _ = v2
    _ = v23
    _ = v230
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
