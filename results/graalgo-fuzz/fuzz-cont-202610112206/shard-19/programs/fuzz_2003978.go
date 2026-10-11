package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 bool; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8; F1 uint64 }
type GenS9 struct { F0 int32; F1 uint16 }
type GenS10 struct { F0 uint64; F1 int32; F2 uint }
type GenS11 struct { F0 int; F1 bool; F2 int8 }
type GenS12 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS13 struct { F0 int8; F1 uint8; F2 uint }
type GenS14 struct { F0 int; F1 uint32 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 int8; F1 uint }
type GenS17 struct { F0 uint32; F1 int32 }
type GenS18 struct { F0 uint8; F1 int8; F2 int32 }
type GenS19 struct { F0 uint; F1 string; F2 int8 }
type GenS20 struct { F0 uint16; F1 bool; F2 int64 }
type GenS21 struct { F0 int32; F1 int32; F2 bool }
type GenS22 struct { F0 uint64; F1 int; F2 uint32 }

func main() {
    var v1 uint32 = 2
    v2 := make(map[uint32]uint32)
    var v3 int32 = 0
    var v4 int32 = 42
    v5 := []uint32{(^(v1)), 2}
    switch "hello" {
    case "\n\t":
        for fi6 := 0; fi6 < 8; fi6++ {
            switch "\n\t" {
            case "\n\t":
                var v7 uint8 = 2
                _ = v7
            }
            switch "\n\t" {
            case "hello":
                switch (fi6 + fi6) {
                case 42:
                    v3 = v4
                default:
                    v8 := GenS1{F0: fi6}
                    _ = v8
                }
            case "":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "\n\t":
                _ = len(v2)
            default:
                for fi9 := 0; fi9 < 19; fi9++ {
                    v1 -= 690
                    _ = fi9
                }
            }
            _ = fi6
        }
    }
    v10 := GenS2{F0: 2, F1: (v4 ^ v4)}
    v1 += (^(v1))
    switch "\n\t" {
    case "foo bar":
        v1 += (v1 / ((v1) | 1))
    case "\n\t":
        fmt.Println(v4, v1)
    case "a":
        v3 = -2147483648
    default:
        switch ("" + "\n\t") {
        case "":
            v1 = (v1 - 2147483648)
        case "a":
            v4 += -834
        }
    }
    v11 := GenS1{F0: -1}
    if false {
        v2[v1] = v1
        v1 = 0
    }
    var v12 int8 = 1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v13 string = ("hello" + "a")
    v14 := GenS3{F0: (1 > 0), F1: 240, F2: 9223372036854775807}
    v3 -= v4
    if (!(v1 == v1)) {
        for fi15 := 0; fi15 < 7; fi15++ {
            fmt.Println(v1)
            switch (v13 + v13) {
            case "":
                var v16 int16 = -1
                _ = v16
            case "hello":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "\n\t":
                if (1.5 != (0.1 * (1.5 / (0.1 - ((float64(0.1) + 0.0) / 0.1))))) {
                    v17 := GenS4{F0: 976, F1: (v13 + (v13 + v13))}
                    v18 := v2[2]
                    _ = v17
                    _ = v18
                }
            }
            if (true || false) {
                v2[v1] = 1
                var v19 int32 = (-(v3))
                if true {
                    v19 -= v19
                }
                _ = v19
            } else {
                switch 4294967295 {
                case 591:
                    if (65535 != v1) {
                        _ = len(v2)
                        var v20 int32 = v4
                        fmt.Println(fi15)
                        _ = v20
                    }
                case 1:
                    for fi21 := 0; fi21 < 15; fi21++ {
                        switch "" {
                        case "":
                            v4 = (-(v4))
                        case "hello":
                            if ((fi21 ^ (-(fi15))) <= fi15) {
                                v22 := GenS5{F0: "hello"}
                                var v23 uint = 255
                                _ = v22
                                _ = v23
                            }
                        }
                        var v24 uint16 = 2
                        switch 2147483647 {
                        case 9223372036854775807:
                            if (!false) {
                                v2[v1] = v1
                                fmt.Println(v24)
                                if (v12 > (v12 & 111)) {
                                    v13 = v13
                                }
                            } else {
                                var v25 bool = (!true)
                                _ = v25
                            }
                        case -519:
                            v4 += v3
                            v13 = ("" + v13)
                            delete(v2, (v1 << ((18446744073709551615) & 63)))
                        case 0:
                            fmt.Println(v12, fi15, v24)
                        }
                        switch v13 {
                        case "hello":
                            if true {
                                var v26 uint32 = (v1 + 2)
                                _ = v26
                            } else {
                                v27 := GenS6{F0: (!false)}
                                v1 = (v1 | (v1 << ((425) & 63)))
                                if false {
                                    var v28 uint16 = (v24 / ((v24) | 1))
                                    _ = v28
                                }
                                _ = v27
                            }
                        case "foo bar":
                            if (979 == v3) {
                                v13 = v13
                                fmt.Println(v4, v24, fi15)
                                v3 = -1
                            }
                        default:
                            v29 := GenS3{F0: true, F1: 114, F2: -835}
                            _ = v29
                        }
                        v1 -= 2147483647
                        fmt.Println(v24, v12)
                        if false {
                            var v30 uint16 = (v24 - (v24 | (v24 >> ((42) & 63))))
                            var v31 uint32 = (v1 + v1)
                            _ = v30
                            _ = v31
                        } else {
                            var v32 bool = false
                            _ = v32
                        }
                        _ = fi21
                        _ = v24
                    }
                default:
                    var v33 uint32 = v1
                    _ = v33
                }
                v2[(v1 << ((950) & 63))] = (v1 >> ((255) & 63))
            }
            _ = fi15
        }
    }
    var v34 int8 = (v12 | 0)
    if true {
        switch 4294967295 {
        case 128:
            v13 = ("foo bar" + "foo bar")
        case 127:
            var v35 int64 = -9223372036854775808
            _ = v35
        }
        var v36 uint16 = 2
        _ = v36
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        for fi38 := 0; fi38 < 9; fi38++ {
            switch (("\n\t" + (v13 + ((v13 + v13) + v13))) + v13) {
            case "hello":
                v12 = (v12 * v34)
            case "\n\t":
                if ("\n\t" == "hello") {
                    v39 := GenS6{F0: (255 == 1)}
                    _ = v39
                }
            case "a":
                var v40 uint8 = 143
                _ = v40
            default:
                fmt.Println(v12, v12)
            }
            _ = fi38
        }
        if (2147483647 > (v3 * -1)) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v3 += (v4 & (^(v3)))
            var v41 float32 = 0.0
            _ = v41
        } else {
            for fi42 := 0; fi42 < 10; fi42++ {
                if (4294967295 == 0) {
                    if false {
                        fmt.Println(fi37, v12, v4)
                        v13 = v13
                        switch "a" {
                        case "":
                            var v43 string = "\n\t"
                            _ = v43
                        case "hello":
                            if true {
                                v12 += 1
                            }
                        default:
                            if true {
                                var v44 float32 = 1e10
                                _ = v44
                            } else {
                                switch "foo bar" {
                                case "hello":
                                    var v45 bool = false
                                    _ = v45
                                case "a":
                                    var v46 float64 = 1e10
                                    _ = v46
                                case "foo bar":
                                    v13 = v13
                                default:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                if (1e10 < float64((((1.5 * (0.0 - (float32(3.14159) + 0.0))) / (1.5 / ((0.0 + ((((-2.25 + 0.0) / ((float32(3.14159) + 0.0) - 1e10)) + -2.25) / 0.1)) * 1e10))) + (0.1 * (float32((0.0 / 1.0)) / ((1.5 * (3.14159 - 0.0)) - (((0.0 / (1e10 + float32((((float64(1e10) - 1e10) / ((1e10 * (1.5 + 0.0)) + (1.5 / float64(1e10)))) / (((3.14159 + 0.0) / 0.1) + ((3.14159 + 3.14159) * (1.5 + 1e10))))))) * (0.0 - 3.14159)) + 0.0))))))) {
                                    v1 += v1
                                    if (false || false) {
                                        if false {
                                            var v47 uint16 = 65535
                                            if false {
                                                var v48 float32 = 0.1
                                                delete(v2, (v1 * v1))
                                                _ = v48
                                            }
                                            fmt.Println(fi42)
                                            _ = v47
                                        } else {
                                            var v49 int16 = -741
                                            _ = v49
                                        }
                                        var v50 int8 = v34
                                        if true {
                                            var v51 int64 = 0
                                            switch "a" {
                                            case "":
                                                v52 := GenS7{F0: (v13 + v13)}
                                                _ = v52
                                            case "a":
                                                if false {
                                                    if (0.1 != ((-2.25 * -2.25) * (1.5 * 1e10))) {
                                                        var v53 uint32 = (v1 & v1)
                                                        _ = v53
                                                    }
                                                    v13 = ("a" + v13)
                                                } else {
                                                    if false {
                                                        var v54 int = fi37
                                                        var v55 string = "hello"
                                                        v56 := GenS8{F0: v12, F1: 4294967295}
                                                        _ = v54
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                    var v57 int32 = v4
                                                    _ = v57
                                                }
                                            default:
                                                if false {
                                                    switch (v13 + (v13 + "a")) {
                                                    case "a":
                                                        switch "" {
                                                        case "hello":
                                                            if true {
                                                                v51 = v51
                                                                v58 := GenS2{F0: 42, F1: (v3 & v3)}
                                                                if true {
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        if false {
                                                                            var v59 bool = false
                                                                            var v60 int = fi37
                                                                            switch 32768 {
                                                                            case -110:
                                                                                v4 -= (v4 ^ v3)
                                                                            case -2147483648:
                                                                                var v61 uint16 = 456
                                                                                _ = v61
                                                                            default:
                                                                                fmt.Println(v12, v3)
                                                                            }
                                                                            _ = v59
                                                                            _ = v60
                                                                        }
                                                                    default:
                                                                        var v62 float64 = (-2.25 - -2.25)
                                                                        _ = v62
                                                                    }
                                                                    v12 = v12
                                                                }
                                                                _ = v58
                                                            }
                                                        case "a":
                                                            if ((v1 << ((363) & 63)) < 65535) {
                                                                if true {
                                                                    var v63 int64 = (v51 | v51)
                                                                    v12 = -1
                                                                    delete(v2, v1)
                                                                    if ((v3 ^ 0) >= v3) {
                                                                        v63 = v51
                                                                        fmt.Println(v63)
                                                                        var v64 int = fi37
                                                                        _ = v64
                                                                    } else {
                                                                        v13 = v13
                                                                    }
                                                                    if (v50 == (v34 | (v34 << ((277) & 63)))) {
                                                                        v65 := GenS9{F0: v3, F1: 0}
                                                                        switch (((v13 + "hello") + "hello") + "\n\t") {
                                                                        case "a":
                                                                            v2[(^(v1))] = v1
                                                                        case "\n\t":
                                                                            var v66 uint = 18446744073709551615
                                                                            _ = v66
                                                                        default:
                                                                            v51 += v63
                                                                        }
                                                                        v67 := GenS3{F0: (0.0 <= (float32(-2.25) - 0.1)), F1: 127, F2: v63}
                                                                        _ = v65
                                                                        _ = v67
                                                                    }
                                                                    v68 := v2[v1]
                                                                    v12 += 127
                                                                    _ = v63
                                                                    _ = v68
                                                                }
                                                                var v69 uint16 = 65535
                                                                v70 := GenS5{F0: ("hello" + "")}
                                                                _ = v69
                                                                _ = v70
                                                            } else {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                v34 += 127
                                                            }
                                                        default:
                                                            v71 := v2[(v1 & 42)]
                                                            _ = v71
                                                        }
                                                    case "foo bar":
                                                        if len(v5) > 0 {
                                                            v5[0] = (v1 - (^(v1)))
                                                        }
                                                    case "hello":
                                                        if ((v1 >> ((4294967295) & 63)) > v1) {
                                                            var v72 float64 = float64(0.0)
                                                            var v73 int32 = 241
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                    default:
                                                        v13 = "hello"
                                                    }
                                                    if len(v5) > 0 {
                                                        v5[0] = v1
                                                    }
                                                    v3 = -1
                                                    switch (-(fi37)) {
                                                    case -991:
                                                        var v74 int32 = v3
                                                        _ = v74
                                                    }
                                                    var v75 uint32 = (v1 << ((0) & 63))
                                                    v12 += v12
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    switch (((v13 + ("" + v13)) + "hello") + ("a" + "hello")) {
                                                    case "foo bar":
                                                        v13 = v13
                                                    case "":
                                                        v76 := GenS2{F0: 770, F1: (v3 + v4)}
                                                        _ = v76
                                                    case "\n\t":
                                                        var v77 int32 = int32(v12)
                                                        _ = v77
                                                    default:
                                                        if (true && false) {
                                                            switch 1 {
                                                            case -375:
                                                                var v78 uint32 = (v1 | 42)
                                                                _ = v78
                                                            case 9223372036854775807:
                                                                var v79 int = fi37
                                                                _ = v79
                                                            default:
                                                                if ((float32(-2.25) != float32(0.1)) || false) {
                                                                    v13 = "a"
                                                                    v50 -= 51
                                                                    v1 = v75
                                                                    fmt.Println(v50, v1)
                                                                }
                                                            }
                                                            switch fi37 {
                                                            case 32767:
                                                                switch (^(fi37)) {
                                                                case 42:
                                                                    v5 = append(v5, 1)
                                                                case 1:
                                                                    var v80 int8 = -3
                                                                    _ = v80
                                                                case 9223372036854775807:
                                                                    v50 -= -128
                                                                }
                                                            case -303:
                                                                switch 42 {
                                                                case -410:
                                                                    var v81 uint64 = 570
                                                                    _ = v81
                                                                case -582:
                                                                    if (42 < 915) {
                                                                        var v82 uint64 = 18446744073709551615
                                                                        _ = v82
                                                                    } else {
                                                                        if false {
                                                                            if (v50 == (v34 / ((v34) | 1))) {
                                                                                v83 := GenS10{F0: 18446744073709551615, F1: v3, F2: 2}
                                                                                _ = v83
                                                                            } else {
                                                                                v34 -= v12
                                                                                if true {
                                                                                    if (2 == 877) {
                                                                                        switch "\n\t" {
                                                                                        case "":
                                                                                            var v84 uint32 = 861
                                                                                            _ = v84
                                                                                        }
                                                                                        var v85 bool = (float64(((((1e10 - ((0.1 - 1.5) * 1.5)) - ((float32(float64((3.14159 - float32(0.0)))) - float32((0.0 + 3.14159))) * -2.25)) * 3.14159) / 1.5)) > float64(0.1))
                                                                                        var v86 float32 = float32(3.14159)
                                                                                        _ = v85
                                                                                        _ = v86
                                                                                    } else {
                                                                                        var v87 float32 = 1e10
                                                                                        var v88 int16 = 0
                                                                                        _ = len(v5)
                                                                                        _ = v87
                                                                                        _ = v88
                                                                                    }
                                                                                }
                                                                                v51 += 42
                                                                                if (v51 <= (v51 + v51)) {
                                                                                    var v89 uint = 0
                                                                                    v89 = v89
                                                                                    _ = v89
                                                                                }
                                                                                var v90 int64 = (^(v51))
                                                                                var v91 uint = 4294967295
                                                                                v4 += (v4 * (v3 | int32(v1)))
                                                                                _ = v90
                                                                                _ = v91
                                                                            }
                                                                            v50 = (v12 & v34)
                                                                            var v92 float32 = 0.1
                                                                            _ = v92
                                                                        }
                                                                        switch ("\n\t" + v13) {
                                                                        case "":
                                                                            switch (fi37 * 127) {
                                                                            case 1:
                                                                                if true {
                                                                                    v93 := GenS9{F0: (v4 & -664), F1: 0}
                                                                                    _ = v93
                                                                                } else {
                                                                                    fmt.Println(v13, v3)
                                                                                }
                                                                            }
                                                                        default:
                                                                            _ = len(v5)
                                                                        }
                                                                        v3 = 1
                                                                        var v94 string = ("hello" + v13)
                                                                        fmt.Println(v94, v13)
                                                                        _ = v94
                                                                    }
                                                                }
                                                                v95 := v2[2147483648]
                                                                _ = v95
                                                            case -192:
                                                                v51 = v51
                                                            default:
                                                                switch v13 {
                                                                case "hello":
                                                                    switch fi42 {
                                                                    case -802:
                                                                        v96 := GenS11{F0: 826, F1: true, F2: -62}
                                                                        _ = v96
                                                                    case 4294967295:
                                                                        if len(v5) > 0 {
                                                                            v5[0] = v75
                                                                        }
                                                                    case 0:
                                                                        var v97 uint32 = v75
                                                                        _ = v97
                                                                    default:
                                                                        v98 := GenS5{F0: ""}
                                                                        _ = v98
                                                                    }
                                                                case "foo bar":
                                                                    var v99 int32 = v4
                                                                    _ = v99
                                                                default:
                                                                    v100 := GenS12{F0: (^(v1)), F1: v51, F2: (v1 / ((v1) | 1))}
                                                                    _ = v100
                                                                }
                                                            }
                                                            var v101 uint32 = (v75 / ((v75) | 1))
                                                            _ = v101
                                                        } else {
                                                            switch (fi37 % ((fi37) | 1)) {
                                                            case 9223372036854775807:
                                                                if len(v5) > 0 {
                                                                    v5[0] = (v1 << ((703) & 63))
                                                                }
                                                            default:
                                                                var v102 float32 = (3.14159 * (((3.14159 / 1.5) - 1.5) * 1e10))
                                                                _ = v102
                                                            }
                                                            v1 = v1
                                                            v75 -= 65535
                                                        }
                                                    }
                                                    _ = v75
                                                } else {
                                                    if ((97 <= 42) && true) {
                                                        var v103 int8 = v34
                                                        v104 := GenS13{F0: v50, F1: 255, F2: uint(v1)}
                                                        var v105 uint = 866
                                                        _ = v103
                                                        _ = v104
                                                        _ = v105
                                                    } else {
                                                        v1 -= 69
                                                    }
                                                    v106 := GenS11{F0: (-(fi42)), F1: ((!((1.5 * 1.5) <= 0.0)) && (!false)), F2: int8(v1)}
                                                    _ = len(v5)
                                                    _ = v106
                                                }
                                            }
                                            v107 := GenS9{F0: -970, F1: 2}
                                            if true {
                                                v3 -= 1
                                                var v108 float32 = 0.0
                                                var v109 uint16 = 32767
                                                if true {
                                                    var v110 uint32 = v1
                                                    v34 = v50
                                                    _ = v110
                                                } else {
                                                    fmt.Println(v3, v12)
                                                    _ = len(v2)
                                                }
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v111 uint = 0
                                                    _ = v111
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = (v1 << ((48) & 63))
                                                }
                                                var v112 uint8 = 255
                                                _ = v108
                                                _ = v109
                                                _ = v112
                                            }
                                            _ = v107
                                            _ = v51
                                        }
                                        _ = v50
                                    } else {
                                        switch ("" + "a") {
                                        case "\n\t":
                                            _ = len(v5)
                                            v113 := GenS14{F0: fi42, F1: 2}
                                            switch "" {
                                            case "a":
                                                v114 := GenS8{F0: v12, F1: 740}
                                                if (!true) {
                                                    v115 := GenS3{F0: ((fi37 * -129) > 9223372036854775807), F1: 27, F2: 9223372036854775807}
                                                    if (127 > 128) {
                                                        var v116 float32 = -2.25
                                                        _ = v116
                                                    }
                                                    _ = v115
                                                } else {
                                                    switch "a" {
                                                    case "":
                                                        if false {
                                                            var v117 int = fi37
                                                            if (!(false && false)) {
                                                                v118 := GenS10{F0: 42, F1: 42, F2: 4294967295}
                                                                if true {
                                                                    if (-123 >= 1) {
                                                                        if (!false) {
                                                                            if false {
                                                                                var v119 int16 = -1
                                                                                switch v13 {
                                                                                case "hello":
                                                                                    var v120 float64 = 0.0
                                                                                    _ = v120
                                                                                case "\n\t":
                                                                                    var v121 uint64 = 690
                                                                                    _ = v121
                                                                                case "foo bar":
                                                                                    var v122 int64 = -9223372036854775808
                                                                                    _ = v122
                                                                                }
                                                                                _ = v119
                                                                            }
                                                                            var v123 uint8 = uint8(v1)
                                                                            fmt.Println(v3)
                                                                            var v124 uint = 65535
                                                                            var v125 uint16 = 2
                                                                            v4 -= v4
                                                                            _ = v123
                                                                            _ = v124
                                                                            _ = v125
                                                                        }
                                                                    } else {
                                                                        if (true || true) {
                                                                            var v126 int32 = 42
                                                                            v127 := GenS15{F0: "\n\t"}
                                                                            _ = v126
                                                                            _ = v127
                                                                        } else {
                                                                            if true {
                                                                                v2[v1] = v1
                                                                                if (!true) {
                                                                                    var v128 float64 = 0.1
                                                                                    _ = v128
                                                                                }
                                                                            } else {
                                                                                v5 = append(v5, v1)
                                                                                fmt.Println(v3, fi37, fi37)
                                                                                if (2 <= 42) {
                                                                                    var v129 int = (-(v117))
                                                                                    _ = v129
                                                                                } else {
                                                                                    var v130 uint = 4294967295
                                                                                    _ = v130
                                                                                }
                                                                                v1 += v1
                                                                                var v131 uint8 = 2
                                                                                v34 = (v34 >> ((255) & 63))
                                                                                _ = v131
                                                                            }
                                                                            v34 += (-(v12))
                                                                            var v132 uint32 = v1
                                                                            _ = v132
                                                                        }
                                                                        if false {
                                                                            v12 -= (v12 + (-(v34)))
                                                                            var v133 int8 = 42
                                                                            if ((324 != 4294967295) && true) {
                                                                                v134 := GenS3{F0: true, F1: 42, F2: -9223372036854775808}
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                                switch "" {
                                                                                case "":
                                                                                    v4 = v3
                                                                                case "a":
                                                                                    var v135 int32 = (-(v4))
                                                                                    _ = v135
                                                                                case "\n\t":
                                                                                    fmt.Println(fi42, v1, v4)
                                                                                    v136 := GenS15{F0: v13}
                                                                                    _ = len(v5)
                                                                                    v133 -= 1
                                                                                    _ = v136
                                                                                }
                                                                                fmt.Println(v117, fi37)
                                                                                _ = v134
                                                                            }
                                                                            _ = v133
                                                                        }
                                                                        v12 = (-(v12))
                                                                        var v137 bool = (1 == 475)
                                                                        v137 = false
                                                                        v138 := GenS12{F0: 1, F1: -459, F2: v1}
                                                                        v34 = 1
                                                                        _ = v137
                                                                        _ = v138
                                                                    }
                                                                    var v139 float64 = 3.14159
                                                                    v139 = 3.14159
                                                                    _ = v139
                                                                } else {
                                                                    if true {
                                                                        var v140 int32 = 42
                                                                        _ = v140
                                                                    } else {
                                                                        v141 := GenS7{F0: v13}
                                                                        var v142 uint32 = 2
                                                                        var v143 uint32 = 492
                                                                        var v144 string = (v13 + "")
                                                                        _ = v141
                                                                        _ = v142
                                                                        _ = v143
                                                                        _ = v144
                                                                    }
                                                                    var v145 float64 = (float64(0.0) * 0.1)
                                                                    v2[v1] = (v1 / ((v1) | 1))
                                                                    v3 += (v4 + (v3 * -437))
                                                                    _ = v145
                                                                }
                                                                v146 := GenS4{F0: 32767, F1: v13}
                                                                _ = v118
                                                                _ = v146
                                                            }
                                                            v4 = -669
                                                            _ = v117
                                                        } else {
                                                            _ = len(v2)
                                                            v1 += 2147483647
                                                            if (-860 == -1) {
                                                                if false {
                                                                    switch (fi42 >> ((65535) & 63)) {
                                                                    case -664:
                                                                        fmt.Println(v1, v12)
                                                                    case 128:
                                                                        if (277 > 0) {
                                                                            v147 := GenS4{F0: 255, F1: v13}
                                                                            v3 = v3
                                                                            v2[v1] = (v1 ^ v1)
                                                                            _ = v147
                                                                        }
                                                                    default:
                                                                        v148 := GenS15{F0: v13}
                                                                        _ = v148
                                                                    }
                                                                    if (false && ((!(127 == 165)) || (((v13 + "a") + (v13 + "hello")) != "foo bar"))) {
                                                                        if false {
                                                                            switch 1 {
                                                                            case 2147483647:
                                                                                v149 := GenS16{F0: 114, F1: 174}
                                                                                _ = v149
                                                                            case 32768:
                                                                                v150 := GenS16{F0: v12, F1: 255}
                                                                                _ = v150
                                                                            }
                                                                            v151 := GenS13{F0: v12, F1: 2, F2: 1}
                                                                            _ = v151
                                                                        } else {
                                                                            if false {
                                                                                switch (^(fi37)) {
                                                                                case -1:
                                                                                    var v152 int8 = 0
                                                                                    _ = v152
                                                                                default:
                                                                                    if ((fi37 ^ 128) < (^(fi37))) {
                                                                                        fmt.Println(v34, v34, v1)
                                                                                        switch fi42 {
                                                                                        case 252:
                                                                                            switch (fi37 >> ((1) & 63)) {
                                                                                            case 4294967295:
                                                                                                v153 := GenS17{F0: 42, F1: v3}
                                                                                                _ = v153
                                                                                            case -334:
                                                                                                var v154 int8 = (-(v12))
                                                                                                _ = v154
                                                                                            default:
                                                                                                fmt.Println(v34)
                                                                                            }
                                                                                        case 32768:
                                                                                            var v155 int32 = (v3 + v3)
                                                                                            _ = v155
                                                                                        }
                                                                                    } else {
                                                                                        v4 += v3
                                                                                    }
                                                                                }
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                            } else {
                                                                                switch (fi42 * fi42) {
                                                                                case 32767:
                                                                                    v156 := GenS18{F0: 2, F1: (v34 ^ (v12 + (v34 >> ((4294967295) & 63)))), F2: v4}
                                                                                    _ = v156
                                                                                default:
                                                                                    v4 = v3
                                                                                    var v157 uint8 = 162
                                                                                    var v158 int64 = 888
                                                                                    _ = v157
                                                                                    _ = v158
                                                                                }
                                                                                var v159 float64 = (3.14159 * 3.14159)
                                                                                var v160 int8 = v12
                                                                                _ = v159
                                                                                _ = v160
                                                                            }
                                                                            var v161 string = (("a" + ((v13 + ((v13 + ((v13 + ((((((v13 + v13) + "hello") + "a") + ("foo bar" + v13)) + v13) + v13)) + v13)) + "foo bar")) + "hello")) + "\n\t")
                                                                            _ = v161
                                                                        }
                                                                        v162 := GenS7{F0: (v13 + "\n\t")}
                                                                        _ = v162
                                                                    }
                                                                }
                                                                var v163 int8 = -128
                                                                var v164 uint = 0
                                                                v165 := GenS3{F0: ((42 == 9223372036854775807) || (fi42 == (fi37 | (fi37 << (((v164 >> ((v164) & 63))) & 63))))), F1: 0, F2: 0}
                                                                _ = v163
                                                                _ = v164
                                                                _ = v165
                                                            }
                                                        }
                                                    case "hello":
                                                        v166 := GenS19{F0: 2, F1: v13, F2: v12}
                                                        _ = v166
                                                    default:
                                                        var v167 float32 = 1e10
                                                        if (!true) {
                                                            v168 := GenS11{F0: (fi42 + (fi37 & fi37)), F1: (1 != 1000), F2: (v34 / (((v12 % ((int8(fi37)) | 1))) | 1))}
                                                            v4 += (v4 & int32(fi42))
                                                            _ = v168
                                                        }
                                                        if false {
                                                            v4 += v4
                                                        } else {
                                                            v4 = (-(v4))
                                                        }
                                                        _ = v167
                                                    }
                                                }
                                                _ = v114
                                            }
                                            _ = v113
                                        default:
                                            v34 += (v12 / (((-(v12))) | 1))
                                        }
                                        v3 = v3
                                    }
                                } else {
                                    delete(v2, v1)
                                    v169 := v2[1]
                                    _ = v169
                                }
                            }
                        }
                    }
                    fmt.Println(fi42, v4, v4)
                    v3 += v3
                } else {
                    v12 = (-(v34))
                }
                switch (v13 + "hello") {
                case "foo bar":
                    v170 := v2[2147483648]
                    _ = v170
                }
                if true {
                    v171 := GenS14{F0: 42, F1: (v1 << ((564) & 63))}
                    v1 -= (v1 * 1)
                    v172 := GenS19{F0: 4294967295, F1: (v13 + v13), F2: -1}
                    _ = v171
                    _ = v172
                } else {
                    if (!((42 > 42) && false)) {
                        fmt.Println(v34, v1)
                        fmt.Println(v34, fi37, fi42)
                    } else {
                        v12 += 1
                    }
                    v173 := GenS20{F0: 65535, F1: (((3.14159 / float32(3.14159)) >= (1e10 - ((1e10 / -2.25) - ((3.14159 - (float32((0.0 + ((float64(-2.25) / 1.5) / 1.0))) + float32(((1.5 + ((((1e10 - 1.5) / ((-2.25 - ((0.0 * 0.0) / 1.0)) - 3.14159)) / ((1.5 * ((3.14159 / -2.25) * (float64((((0.1 + 3.14159) + 1.5) / 1.5)) / 1.0))) / -2.25)) * 1.5)) + 0.0)))) + (-2.25 - (((0.0 - 0.0) / ((1e10 - 0.0) * 3.14159)) / float32(3.14159))))))) || true), F2: 1}
                    v34 = v12
                    fmt.Println(fi37, v13, v34)
                    v174 := GenS21{F0: (-(v3)), F1: (v4 << ((2) & 63)), F2: true}
                    if (4294967295 >= 65535) {
                        if true {
                            var v175 string = ((v13 + ("" + (v13 + v13))) + "a")
                            v2[(^(v1))] = v1
                            v1 += (^(v1))
                            _ = v175
                        } else {
                            _ = len(v5)
                            v176 := GenS22{F0: 2, F1: (fi37 + 390), F2: v1}
                            v4 = v3
                            v13 = ("" + v13)
                            fmt.Println(v1)
                            _ = v176
                        }
                    } else {
                        switch fi42 {
                        case -2147483648:
                            var v177 uint32 = (v1 & v1)
                            _ = v177
                        }
                    }
                    _ = v173
                    _ = v174
                }
                _ = fi42
            }
        }
        var v178 uint8 = 18
        _ = fi37
        _ = v178
    }
    for fi179 := 0; fi179 < 2; fi179++ {
        v180 := GenS14{F0: (fi179 | (fi179 | fi179)), F1: v1}
        _ = fi179
        _ = v180
    }
    _ = len(v2)
    var v181 int32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v181:", v181)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v181
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
