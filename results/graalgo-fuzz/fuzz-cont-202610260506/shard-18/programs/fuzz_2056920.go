package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 int8; F1 int; F2 uint16 }
type GenS5 struct { F0 int64; F1 int32; F2 int32 }
type GenS6 struct { F0 bool; F1 string; F2 uint8 }
type GenS7 struct { F0 uint; F1 int32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int; F1 int }
type GenS10 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint16; F1 int8; F2 int8 }
type GenS13 struct { F0 string; F1 int16 }
type GenS14 struct { F0 int32; F1 int }
type GenS15 struct { F0 uint32; F1 uint64 }
type GenS16 struct { F0 int8; F1 string }
type GenS17 struct { F0 int64 }

func main() {
    var v1 uint8 = 255
    var v2 int16 = -1
    v3 := []uint32{65535, 42}
    var v4 uint32 = 0
    var v5 uint8 = v1
    v2 -= 42
    fmt.Println(v2)
    for fi6 := 0; fi6 < 11; fi6++ {
        if len(v3) > 0 {
            v3[0] = v4
        }
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 float32 = (float32(float64(3.14159)) - ((float32((3.14159 * float64(float32(0.1)))) - float32(-2.25)) - float32((float64(0.1) / ((0.1 + 1.5) + float64(0.0))))))
            var v9 uint = 1
            switch ("" + "") {
            case "foo bar":
                switch "a" {
                case "hello":
                    v5 -= 128
                }
            default:
                var v10 int16 = v2
                _ = v10
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    if true {
        var v11 float64 = (0.0 * 1.5)
        _ = len(v3)
        var v12 uint = 255
        _ = v11
        _ = v12
    }
    v2 = v2
    switch 0 {
    case 128:
        fmt.Println(v4)
    case -294:
        v13 := GenS1{F0: (v4 % ((v4) | 1))}
        _ = v13
    default:
        v14 := GenS2{F0: uint32(v1)}
        _ = v14
    }
    var v15 int = 32768
    v2 += (-(v2))
    if false {
        if ("\n\t" <= "foo bar") {
            var v16 uint16 = 603
            _ = v16
        }
        var v17 uint = 65535
        fmt.Println(v2, v1)
        _ = v17
    }
    var v18 uint = 1
    _ = len(v3)
    switch "a" {
    case "\n\t":
        var v19 float32 = -2.25
        _ = v19
    case "foo bar":
        if false {
            fmt.Println(v18, v4)
            switch ("\n\t" + "a") {
            case "a":
                var v20 float64 = -2.25
                _ = v20
            case "foo bar":
                var v21 int16 = 32767
                _ = v21
            case "\n\t":
                if false {
                    _ = len(v3)
                } else {
                    for fi22 := 0; fi22 < 12; fi22++ {
                        var v23 int = (fi22 - -2147483648)
                        for fi24 := 0; fi24 < 5; fi24++ {
                            v25 := GenS3{F0: v4, F1: v1}
                            fmt.Println(fi22)
                            _ = fi24
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 3; fi26++ {
                            var v27 uint8 = 40
                            if ((v27 + (^(v5))) >= (v27 ^ v1)) {
                                var v28 int32 = -683
                                switch (v23 + fi22) {
                                case 9223372036854775807:
                                    v4 = (^(v4))
                                case 32767:
                                    _ = len(v3)
                                }
                                var v29 uint = (^(v18))
                                _ = v28
                                _ = v29
                            } else {
                                v30 := GenS1{F0: v4}
                                switch ("" + "\n\t") {
                                case "a":
                                    v1 -= (v1 & 43)
                                }
                                _ = v30
                            }
                            if len(v3) > 0 {
                                v3[0] = (v4 - 344)
                            }
                            _ = fi26
                            _ = v27
                        }
                        _ = fi22
                        _ = v23
                    }
                    switch "hello" {
                    case "foo bar":
                        v31 := GenS3{F0: 469, F1: (v5 << ((v18) & 63))}
                        _ = v31
                    case "a":
                        v2 = 1000
                    default:
                        var v32 int32 = 0
                        _ = v32
                    }
                }
            }
            var v33 int = v15
            _ = v33
        } else {
            var v34 float32 = 0.1
            switch "foo bar" {
            case "foo bar":
                for fi35 := 0; fi35 < 1; fi35++ {
                    var v36 uint8 = 1
                    var v37 uint = v18
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            case "":
                v38 := GenS1{F0: v4}
                _ = v38
            case "hello":
                var v39 uint8 = v1
                _ = v39
            default:
                if (!true) {
                    for fi40 := 0; fi40 < 4; fi40++ {
                        for fi41 := 0; fi41 < 7; fi41++ {
                            switch (fi40 & (fi40 % ((128) | 1))) {
                            case 32768:
                                if (fi41 == 2147483647) {
                                    if (!((v18 > 766) || true)) {
                                        v5 = v1
                                        var v42 uint32 = v4
                                        _ = len(v3)
                                        _ = v42
                                    }
                                    switch 9223372036854775807 {
                                    case 2147483647:
                                        v43 := GenS4{F0: -128, F1: (fi40 - (fi41 & (-(fi41)))), F2: 255}
                                        _ = v43
                                    case -1:
                                        v4 = v4
                                    case 936:
                                        _ = len(v3)
                                    default:
                                        var v44 uint64 = 2
                                        _ = v44
                                    }
                                    if false {
                                        v45 := GenS5{F0: int64(v1), F1: 0, F2: 2147483647}
                                        _ = v45
                                    }
                                    switch (fi40 % (((v15 + v15)) | 1)) {
                                    case -2147483648:
                                        _ = len(v3)
                                    case 128:
                                        v34 = (-2.25 - v34)
                                    case 819:
                                        v46 := GenS4{F0: 0, F1: 32767, F2: 706}
                                        _ = v46
                                    default:
                                        var v47 int64 = -234
                                        _ = v47
                                    }
                                }
                            case -1:
                                v48 := GenS6{F0: true, F1: (("foo bar" + "foo bar") + "hello"), F2: (v1 ^ (v5 << ((v18) & 63)))}
                                _ = v48
                            case 4294967295:
                                switch "hello" {
                                case "a":
                                    var v49 string = ""
                                    _ = v49
                                }
                            default:
                                var v50 bool = (!(false && (false && (42 < 127))))
                                _ = v50
                            }
                            _ = fi41
                        }
                        var v51 uint = v18
                        _ = fi40
                        _ = v51
                    }
                    switch ("" + ("a" + "foo bar")) {
                    case "hello":
                        v2 = 32767
                    }
                    var v52 float64 = 1.5
                    _ = v52
                } else {
                    for fi53 := 0; fi53 < 15; fi53++ {
                        if true {
                            v54 := GenS5{F0: 1, F1: -2147483648, F2: -1}
                            fmt.Println(v5)
                            _ = v54
                        } else {
                            if (!false) {
                                v55 := GenS3{F0: v4, F1: (^(v5))}
                                if (true || (true || (42 <= -1))) {
                                    v3 = append(v3, (v4 * 2147483647))
                                    if false {
                                        if ((false && (((1 >= 1) || true) && true)) || true) {
                                            if true {
                                                for fi56 := 0; fi56 < 11; fi56++ {
                                                    fmt.Println(v34)
                                                    _ = fi56
                                                }
                                                if false {
                                                    switch ("a" + ("" + (("\n\t" + "a") + "foo bar"))) {
                                                    case "foo bar":
                                                        var v57 uint = 2
                                                        _ = v57
                                                    case "a":
                                                        fmt.Println(v1)
                                                    case "":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                } else {
                                                    v4 = 0
                                                }
                                                fmt.Println(v1, v5, v34)
                                            } else {
                                                if false {
                                                    for fi58 := 0; fi58 < 19; fi58++ {
                                                        if len(v3) > 0 {
                                                            v3[0] = 2147483648
                                                        }
                                                        v18 += 530
                                                        _ = fi58
                                                    }
                                                    switch "" {
                                                    case "hello":
                                                        var v59 uint32 = (v4 - (v4 / ((v4) | 1)))
                                                        _ = v59
                                                    }
                                                } else {
                                                    fmt.Println(v1, v18)
                                                    switch 1 {
                                                    case 780:
                                                        for fi60 := 0; fi60 < 8; fi60++ {
                                                            v34 = v34
                                                            if (1.5 == float32((0.1 / (1e10 / 1e10)))) {
                                                                var v61 uint = (v18 & v18)
                                                                var v62 float64 = (-2.25 - (((float64(float32((float64((v34 + v34)) * ((-2.25 + 0.0) / 1e10)))) - 1.5) - float64(v34)) + ((0.0 / 1.5) + 0.1)))
                                                                _ = v61
                                                                _ = v62
                                                            } else {
                                                                if false {
                                                                    var v63 float32 = (float32(0.1) + float32(0.1))
                                                                    _ = len(v3)
                                                                    v63 = v34
                                                                    if true {
                                                                        var v64 uint = (v18 >> ((v18) & 63))
                                                                        _ = v64
                                                                    } else {
                                                                        fmt.Println(v18, v2)
                                                                        switch ("foo bar" + "a") {
                                                                        case "":
                                                                            v65 := GenS6{F0: false, F1: ("\n\t" + "a"), F2: (v1 ^ 56)}
                                                                            _ = v65
                                                                        case "hello":
                                                                            v66 := GenS6{F0: ((true && (0.1 != v34)) && true), F1: (("hello" + "") + "a"), F2: (v5 % ((v1) | 1))}
                                                                            _ = v66
                                                                        }
                                                                        v67 := GenS6{F0: ((!false) || false), F1: "foo bar", F2: v5}
                                                                        _ = v67
                                                                    }
                                                                    _ = v63
                                                                }
                                                            }
                                                            _ = fi60
                                                        }
                                                    case 57:
                                                        v5 = (^(v5))
                                                    default:
                                                        for fi68 := 0; fi68 < 12; fi68++ {
                                                            v5 -= (v5 + (v1 << (((v18 | (^(v18)))) & 63)))
                                                            if (false && false) {
                                                                v3 = append(v3, (^(v4)))
                                                                switch 1 {
                                                                case -426:
                                                                    v34 = -2.25
                                                                }
                                                                v3 = append(v3, v4)
                                                            }
                                                            var v69 string = "a"
                                                            _ = fi68
                                                            _ = v69
                                                        }
                                                    }
                                                }
                                                v70 := GenS5{F0: 254, F1: 1, F2: 2147483647}
                                                switch ("a" + "hello") {
                                                case "hello":
                                                    for fi71 := 0; fi71 < 3; fi71++ {
                                                        var v72 bool = (12 < -1)
                                                        fmt.Println(v18, v15)
                                                        switch fi53 {
                                                        case 127:
                                                            _ = len(v3)
                                                        case 1:
                                                            switch (v15 + (^(v15))) {
                                                            case -467:
                                                                if v72 {
                                                                    v73 := GenS4{F0: -9, F1: v15, F2: 255}
                                                                    switch "hello" {
                                                                    case "":
                                                                        var v74 bool = true
                                                                        fmt.Println(v5)
                                                                        v75 := GenS4{F0: -99, F1: v15, F2: 32767}
                                                                        _ = v74
                                                                        _ = v75
                                                                    case "\n\t":
                                                                        if (!(v2 == (v2 + v2))) {
                                                                            var v76 float64 = (float64(float32((float64(v34) + (0.1 / 1.0)))) - 3.14159)
                                                                            if len(v3) > 0 {
                                                                                v3[0] = uint32(v18)
                                                                            }
                                                                            _ = v76
                                                                        }
                                                                    case "hello":
                                                                        v2 -= v2
                                                                        fmt.Println(v18, v18)
                                                                    default:
                                                                        v1 -= (v5 ^ (^(v5)))
                                                                    }
                                                                    _ = v73
                                                                } else {
                                                                    var v77 uint16 = 0
                                                                    v78 := GenS3{F0: (v4 + (v4 + (v4 + (v4 & (v4 ^ (v4 * 1)))))), F1: v5}
                                                                    _ = v77
                                                                    _ = v78
                                                                }
                                                            case 952:
                                                                if false {
                                                                    var v79 int16 = (-(v2))
                                                                    v72 = v72
                                                                    fmt.Println(v2)
                                                                    _ = v79
                                                                } else {
                                                                    switch (fi53 / ((fi71) | 1)) {
                                                                    case -1:
                                                                        fmt.Println(fi71)
                                                                    case 831:
                                                                        var v80 uint16 = 625
                                                                        _ = v80
                                                                    case 2147483647:
                                                                        if v72 {
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                v4 = (v4 ^ v4)
                                                                                if (682 != 9223372036854775807) {
                                                                                    var v81 int16 = 817
                                                                                    _ = v81
                                                                                } else {
                                                                                    switch ("hello" + "\n\t") {
                                                                                    case "hello":
                                                                                        v82 := GenS2{F0: 2147483648}
                                                                                        _ = v82
                                                                                    case "":
                                                                                        var v83 int = fi53
                                                                                        _ = v83
                                                                                    }
                                                                                    v84 := GenS3{F0: 458, F1: (^(v5))}
                                                                                    v2 += 32767
                                                                                    _ = v84
                                                                                }
                                                                                var v85 uint32 = (v4 & (v4 | (v4 / ((v4) | 1))))
                                                                                _ = v85
                                                                            default:
                                                                                var v86 uint = (v18 | 65535)
                                                                                _ = v86
                                                                            }
                                                                        } else {
                                                                            if true {
                                                                                if (true && (v72 || v72)) {
                                                                                    switch ("" + ("\n\t" + ("" + ((("hello" + ("" + "hello")) + "\n\t") + "a")))) {
                                                                                    case "\n\t":
                                                                                        v15 -= 1
                                                                                    default:
                                                                                        v87 := GenS2{F0: v4}
                                                                                        _ = v87
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(fi53, v18)
                                                                                    switch ("a" + (("hello" + ("a" + "a")) + ("\n\t" + "foo bar"))) {
                                                                                    case "hello":
                                                                                        if ((9223372036854775807 <= -9223372036854775808) && (!true)) {
                                                                                            switch v15 {
                                                                                            case 2147483647:
                                                                                                fmt.Println(v2)
                                                                                            case -2147483648:
                                                                                                var v88 int32 = 2147483647
                                                                                                _ = v88
                                                                                            }
                                                                                            v89 := GenS2{F0: v4}
                                                                                            _ = v89
                                                                                        }
                                                                                    case "foo bar":
                                                                                        switch ("hello" + ("foo bar" + (("foo bar" + "\n\t") + (("hello" + ("foo bar" + ("foo bar" + "foo bar"))) + ("a" + ""))))) {
                                                                                        case "a":
                                                                                            var v90 int8 = 1
                                                                                            _ = v90
                                                                                        case "":
                                                                                            v18 -= v18
                                                                                        case "hello":
                                                                                            v4 -= uint32(v18)
                                                                                        default:
                                                                                            v2 += v2
                                                                                            v4 += (v4 >> ((0) & 63))
                                                                                            if v72 {
                                                                                                switch "a" {
                                                                                                case "foo bar":
                                                                                                    v2 -= (v2 - 469)
                                                                                                case "a":
                                                                                                    v15 = v15
                                                                                                case "":
                                                                                                    var v91 uint16 = 65535
                                                                                                    v34 += 3.14159
                                                                                                    _ = v91
                                                                                                default:
                                                                                                    var v92 int64 = 1
                                                                                                    _ = v92
                                                                                                }
                                                                                                if v72 {
                                                                                                    v5 = v1
                                                                                                }
                                                                                                var v93 uint32 = v4
                                                                                                _ = v93
                                                                                            } else {
                                                                                                fmt.Println(v72, v5, v4)
                                                                                            }
                                                                                        }
                                                                                    case "a":
                                                                                        if (uint16(v1) == 2) {
                                                                                            fmt.Println(fi53, v15, v4)
                                                                                            switch fi53 {
                                                                                            case 715:
                                                                                                v1 = v1
                                                                                            default:
                                                                                                if (fi53 == fi71) {
                                                                                                    if (v5 <= (v5 + 1)) {
                                                                                                        var v94 uint8 = 98
                                                                                                        _ = v94
                                                                                                    }
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = (v4 | 2147483647)
                                                                                                    }
                                                                                                    v1 -= 127
                                                                                                } else {
                                                                                                    if (!false) {
                                                                                                        var v95 int8 = int8(v2)
                                                                                                        if false {
                                                                                                            switch fi71 {
                                                                                                            case 271:
                                                                                                                var v96 int = fi71
                                                                                                                _ = v96
                                                                                                            }
                                                                                                        }
                                                                                                        v97 := GenS7{F0: 1, F1: -2147483648}
                                                                                                        _ = v95
                                                                                                        _ = v97
                                                                                                    } else {
                                                                                                        var v98 int8 = int8(v5)
                                                                                                        if false {
                                                                                                            v99 := GenS8{F0: uint64(v5)}
                                                                                                            var v100 uint = (v18 << ((255) & 63))
                                                                                                            _ = v100
                                                                                                            _ = v99
                                                                                                        }
                                                                                                        switch ("foo bar" + "\n\t") {
                                                                                                        case "":
                                                                                                            fmt.Println(v18)
                                                                                                        case "hello":
                                                                                                            switch "a" {
                                                                                                            case "\n\t":
                                                                                                                v72 = (v72 || v72)
                                                                                                            case "":
                                                                                                                var v101 bool = (-9223372036854775808 > 42)
                                                                                                                _ = v101
                                                                                                            }
                                                                                                        case "\n\t":
                                                                                                            if (992 <= (v18 << (((v18 ^ v18)) & 63))) {
                                                                                                                switch "\n\t" {
                                                                                                                case "hello":
                                                                                                                    switch (-(v15)) {
                                                                                                                    case 282:
                                                                                                                        if (!v72) {
                                                                                                                            switch (^(fi71)) {
                                                                                                                            case 9223372036854775807:
                                                                                                                                v102 := GenS9{F0: (fi71 - fi71), F1: -9223372036854775808}
                                                                                                                                _ = v102
                                                                                                                            case -154:
                                                                                                                                var v103 int8 = v98
                                                                                                                                _ = len(v3)
                                                                                                                                _ = v103
                                                                                                                            default:
                                                                                                                                if (-1 != 0) {
                                                                                                                                    var v104 float64 = -2.25
                                                                                                                                    var v105 float64 = 0.1
                                                                                                                                    _ = v104
                                                                                                                                    _ = v105
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v15, v98)
                                                                                                                                    v106 := GenS4{F0: 42, F1: fi53, F2: 255}
                                                                                                                                    _ = len(v3)
                                                                                                                                    if len(v3) > 0 {
                                                                                                                                        v3[0] = (v4 / ((v4) | 1))
                                                                                                                                    }
                                                                                                                                    var v107 string = ("hello" + "foo bar")
                                                                                                                                    _ = v106
                                                                                                                                    _ = v107
                                                                                                                                }
                                                                                                                            }
                                                                                                                            if ((v5 != (v5 + (v5 * (v1 ^ 2)))) && (0 <= 4294967295)) {
                                                                                                                                v2 = v2
                                                                                                                            } else {
                                                                                                                                var v108 int64 = 0
                                                                                                                                _ = v108
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            v34 += (v34 / ((float32(float64(-2.25)) - v34) * ((1.5 + v34) + float32((-2.25 + (-2.25 * (0.0 + 1e10)))))))
                                                                                                                            var v109 float64 = (float64(((-2.25 * v34) / (v34 / v34))) - 0.1)
                                                                                                                            var v110 uint32 = 2147483648
                                                                                                                            _ = v109
                                                                                                                            _ = v110
                                                                                                                        }
                                                                                                                        v98 += v98
                                                                                                                    default:
                                                                                                                        v3 = append(v3, 842)
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    if (488 < 18446744073709551615) {
                                                                                                                        v111 := GenS10{F0: 128, F1: 0, F2: 1}
                                                                                                                        if (v18 == (v18 * 18446744073709551615)) {
                                                                                                                            fmt.Println(fi71, v72)
                                                                                                                            var v112 uint = (v18 & 2)
                                                                                                                            switch "foo bar" {
                                                                                                                            case "":
                                                                                                                                fmt.Println(v15)
                                                                                                                            case "hello":
                                                                                                                                if (548 != v18) {
                                                                                                                                    v113 := GenS4{F0: (v98 - (v98 - (v98 / (((v98 ^ 42)) | 1)))), F1: (fi53 / (((fi53 | (^(fi53)))) | 1)), F2: 255}
                                                                                                                                    _ = v113
                                                                                                                                } else {
                                                                                                                                    var v114 uint32 = uint32(v18)
                                                                                                                                    fmt.Println(v34, v18)
                                                                                                                                    _ = len(v3)
                                                                                                                                    _ = v114
                                                                                                                                }
                                                                                                                            }
                                                                                                                            v34 += v34
                                                                                                                            _ = v112
                                                                                                                        }
                                                                                                                        _ = v111
                                                                                                                    } else {
                                                                                                                        var v115 int32 = -1
                                                                                                                        var v116 float32 = 0.1
                                                                                                                        v5 += (^(v5))
                                                                                                                        _ = v115
                                                                                                                        _ = v116
                                                                                                                    }
                                                                                                                }
                                                                                                                var v117 float64 = 0.1
                                                                                                                if v72 {
                                                                                                                    v118 := GenS2{F0: v4}
                                                                                                                    if (42 >= -478) {
                                                                                                                        v119 := GenS11{F0: (v18 / ((152) | 1))}
                                                                                                                        v3 = append(v3, (^(v4)))
                                                                                                                        var v120 bool = (255 == 528)
                                                                                                                        _ = v119
                                                                                                                        _ = v120
                                                                                                                    }
                                                                                                                    v4 = (v4 ^ v4)
                                                                                                                    _ = v118
                                                                                                                }
                                                                                                                _ = v117
                                                                                                            } else {
                                                                                                                if v72 {
                                                                                                                    v3 = append(v3, (v4 >> ((65535) & 63)))
                                                                                                                }
                                                                                                                v121 := GenS8{F0: 0}
                                                                                                                var v122 uint = 2
                                                                                                                _ = v121
                                                                                                                _ = v122
                                                                                                            }
                                                                                                        default:
                                                                                                            if v72 {
                                                                                                                v123 := GenS12{F0: 1, F1: (v98 ^ v98), F2: v98}
                                                                                                                _ = v123
                                                                                                            } else {
                                                                                                                v98 += (v98 + (-(v98)))
                                                                                                                v98 += (v98 % (((-(v98))) | 1))
                                                                                                                var v124 uint = (v18 << (((v18 + (v18 >> ((uint(v1)) & 63)))) & 63))
                                                                                                                _ = v124
                                                                                                            }
                                                                                                        }
                                                                                                        v125 := GenS5{F0: 42, F1: -1, F2: -1}
                                                                                                        _ = v125
                                                                                                        _ = v98
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v126 := GenS13{F0: "hello", F1: 32767}
                                                                                            var v127 uint32 = v4
                                                                                            var v128 uint64 = 2
                                                                                            _ = v126
                                                                                            _ = v127
                                                                                            _ = v128
                                                                                        } else {
                                                                                            var v129 int = v15
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            _ = v129
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                        var v130 int16 = (v2 * -32768)
                                                                                        _ = v130
                                                                                    }
                                                                                }
                                                                                var v131 int64 = -1
                                                                                _ = len(v3)
                                                                                _ = v131
                                                                            } else {
                                                                                v3 = append(v3, (v4 % (((v4 / (((v4 - (v4 / ((v4) | 1)))) | 1))) | 1)))
                                                                                v15 = (fi71 * v15)
                                                                                switch (-(fi71)) {
                                                                                case -1:
                                                                                    var v132 int = (-(v15))
                                                                                    _ = v132
                                                                                case 46:
                                                                                    v133 := GenS14{F0: -1, F1: fi53}
                                                                                    _ = v133
                                                                                case -9223372036854775808:
                                                                                    var v134 uint32 = uint32(v18)
                                                                                    v135 := GenS2{F0: v4}
                                                                                    _ = len(v3)
                                                                                    _ = v134
                                                                                    _ = v135
                                                                                default:
                                                                                    var v136 int64 = 42
                                                                                    _ = v136
                                                                                }
                                                                            }
                                                                            switch ("a" + "\n\t") {
                                                                            case "":
                                                                                var v137 int8 = 127
                                                                                _ = v137
                                                                            }
                                                                        }
                                                                    }
                                                                    switch (fi71 >> ((2) & 63)) {
                                                                    case 9223372036854775807:
                                                                        if v72 {
                                                                            v4 = (^(v4))
                                                                            var v138 int8 = 42
                                                                            v139 := GenS9{F0: (fi71 / (((-(v15))) | 1)), F1: v15}
                                                                            v2 = 32767
                                                                            _ = v138
                                                                            _ = v139
                                                                        } else {
                                                                            var v140 int = fi71
                                                                            switch "hello" {
                                                                            case "a":
                                                                                var v141 float32 = v34
                                                                                var v142 int16 = v2
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                v3 = append(v3, (v4 << ((uint(v5)) & 63)))
                                                                                _ = v141
                                                                                _ = v142
                                                                            default:
                                                                                if (v4 > 4294967295) {
                                                                                    v15 -= 0
                                                                                } else {
                                                                                    fmt.Println(fi71, v2, fi53)
                                                                                }
                                                                            }
                                                                            switch "" {
                                                                            case "":
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            case "hello":
                                                                                v72 = v72
                                                                            case "\n\t":
                                                                                _ = len(v3)
                                                                            }
                                                                            _ = v140
                                                                        }
                                                                        _ = len(v3)
                                                                    case 922:
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    }
                                                                    v1 = 1
                                                                }
                                                            case -401:
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            default:
                                                                var v143 int16 = 1000
                                                                _ = v143
                                                            }
                                                        default:
                                                            v144 := GenS12{F0: 703, F1: -4, F2: int8(v1)}
                                                            _ = v144
                                                        }
                                                        _ = fi71
                                                        _ = v72
                                                    }
                                                case "\n\t":
                                                    v15 = -1
                                                }
                                                _ = v70
                                            }
                                        }
                                    } else {
                                        if (!true) {
                                            fmt.Println(v34, v34)
                                            var v145 uint16 = 32767
                                            _ = v145
                                        }
                                        if (!(v4 > uint32(v5))) {
                                            for fi146 := 0; fi146 < 6; fi146++ {
                                                if true {
                                                    switch fi53 {
                                                    case -19:
                                                        _ = len(v3)
                                                    case 266:
                                                        if false {
                                                            fmt.Println(v1, v1, fi53)
                                                        } else {
                                                            if ((v5 + v1) >= 128) {
                                                                if (!true) {
                                                                    var v147 float32 = (v34 * -2.25)
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 0
                                                                    }
                                                                    if false {
                                                                        var v148 uint8 = (v1 + (v5 << (((v18 + v18)) & 63)))
                                                                        switch (("" + "foo bar") + "hello") {
                                                                        case "\n\t":
                                                                            v34 = 1e10
                                                                        default:
                                                                            v15 = -818
                                                                        }
                                                                        if true {
                                                                            var v149 int = (fi146 & v15)
                                                                            _ = v149
                                                                        } else {
                                                                            var v150 int32 = 0
                                                                            var v151 int16 = v2
                                                                            v2 -= 1
                                                                            _ = v150
                                                                            _ = v151
                                                                        }
                                                                        switch v15 {
                                                                        case 498:
                                                                            var v152 int32 = 0
                                                                            _ = v152
                                                                        case 2147483647:
                                                                            v153 := GenS7{F0: (^(v18)), F1: -798}
                                                                            var v154 uint8 = v5
                                                                            _ = v153
                                                                            _ = v154
                                                                        }
                                                                        _ = v148
                                                                    } else {
                                                                        var v155 int8 = -128
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        var v156 int16 = v2
                                                                        _ = v155
                                                                        _ = v156
                                                                    }
                                                                    v157 := GenS15{F0: (^(v4)), F1: 1}
                                                                    if (686 != 387) {
                                                                        v3 = append(v3, v4)
                                                                        v3 = append(v3, (v4 << ((0) & 63)))
                                                                    } else {
                                                                        var v158 uint = v18
                                                                        if true {
                                                                            if ((^(v2)) > 606) {
                                                                                fmt.Println(v4, v158, fi53)
                                                                            }
                                                                            v147 = v34
                                                                            if false {
                                                                                fmt.Println(v18)
                                                                                fmt.Println(v1, v15, v158)
                                                                                switch (fi53 - (fi146 & 1)) {
                                                                                case -970:
                                                                                    if (1 >= 534) {
                                                                                        var v159 bool = true
                                                                                        fmt.Println(v147, fi53, v158)
                                                                                        _ = v159
                                                                                    } else {
                                                                                        if false {
                                                                                            v160 := GenS5{F0: 0, F1: 42, F2: 2147483647}
                                                                                            _ = v160
                                                                                        }
                                                                                        var v161 int8 = 49
                                                                                        _ = v161
                                                                                    }
                                                                                }
                                                                                if false {
                                                                                    v5 = v1
                                                                                    v162 := GenS3{F0: 852, F1: 128}
                                                                                    _ = v162
                                                                                } else {
                                                                                    var v163 int32 = -1
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    if true {
                                                                                        fmt.Println(v158, v147, v18)
                                                                                        if (true && true) {
                                                                                            v147 += v34
                                                                                            fmt.Println(v18, v34)
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = (^(v4))
                                                                                            }
                                                                                            v163 -= (v163 % ((269) | 1))
                                                                                        } else {
                                                                                            var v164 float64 = -2.25
                                                                                            if (42 >= 1) {
                                                                                                v1 = (v5 * (v1 | v5))
                                                                                                fmt.Println(fi53, fi146, fi53)
                                                                                            }
                                                                                            fmt.Println(v163)
                                                                                            _ = v164
                                                                                        }
                                                                                        v158 -= v158
                                                                                    }
                                                                                    var v165 int16 = -32768
                                                                                    var v166 uint8 = 42
                                                                                    var v167 uint16 = 255
                                                                                    _ = v163
                                                                                    _ = v165
                                                                                    _ = v166
                                                                                    _ = v167
                                                                                }
                                                                            }
                                                                        }
                                                                        fmt.Println(v1)
                                                                        _ = v158
                                                                    }
                                                                    _ = v147
                                                                    _ = v157
                                                                }
                                                                if false {
                                                                    var v168 string = "foo bar"
                                                                    _ = v168
                                                                }
                                                                fmt.Println(v4)
                                                            }
                                                            v15 += int(v18)
                                                        }
                                                    case 394:
                                                        if ((v34 - (3.14159 * ((((1e10 + 1e10) - float32((0.0 - 3.14159))) - 0.1) - (3.14159 / v34)))) > 0.1) {
                                                            switch -730 {
                                                            case 2147483647:
                                                                fmt.Println(v18, fi53)
                                                            case 0:
                                                                switch (v15 >> ((1) & 63)) {
                                                                case 128:
                                                                    fmt.Println(v5, v5, v15)
                                                                case 32768:
                                                                    v169 := GenS7{F0: 892, F1: -224}
                                                                    _ = v169
                                                                default:
                                                                    v170 := GenS15{F0: (v4 * 977), F1: 18446744073709551615}
                                                                    _ = v170
                                                                }
                                                            case -1:
                                                                if (!((v2 % ((1) | 1)) <= (v2 >> (((^(v18))) & 63)))) {
                                                                    var v171 float64 = 1.5
                                                                    var v172 uint = (v18 / (((v18 & 543)) | 1))
                                                                    if ((v4 >> ((v18) & 63)) == (v4 * (v4 + v4))) {
                                                                        var v173 int = -436
                                                                        _ = v173
                                                                    }
                                                                    fmt.Println(v172, v34, v171)
                                                                    fmt.Println(v4)
                                                                    _ = len(v3)
                                                                    _ = v171
                                                                    _ = v172
                                                                }
                                                            }
                                                            if true {
                                                                v5 += v1
                                                            }
                                                            switch fi53 {
                                                            case -740:
                                                                if true {
                                                                    var v174 uint16 = 42
                                                                    v175 := GenS16{F0: 127, F1: "\n\t"}
                                                                    switch 346 {
                                                                    case 423:
                                                                        var v176 int = (v15 % ((v15) | 1))
                                                                        _ = v176
                                                                    case 32768:
                                                                        var v177 int = fi53
                                                                        _ = v177
                                                                    }
                                                                    _ = v174
                                                                    _ = v175
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        var v178 uint64 = 1
                                                        fmt.Println(v1, v15, v18)
                                                        if false {
                                                            var v179 uint32 = (v4 >> ((v18) & 63))
                                                            _ = v179
                                                        }
                                                        _ = v178
                                                    }
                                                    fmt.Println(v34, fi53)
                                                    var v180 uint64 = 286
                                                    _ = v180
                                                } else {
                                                    v18 = (^(v18))
                                                }
                                                var v181 uint = v18
                                                _ = fi146
                                                _ = v181
                                            }
                                            v1 = (v5 + v1)
                                        } else {
                                            v182 := GenS7{F0: (v18 + v18), F1: 883}
                                            var v183 int16 = v2
                                            _ = v182
                                            _ = v183
                                        }
                                        _ = len(v3)
                                    }
                                } else {
                                    fmt.Println(v4)
                                }
                                if true {
                                    if true {
                                        v34 = v34
                                    } else {
                                        _ = len(v3)
                                        if (int(v4) < (fi53 * -2147483648)) {
                                            _ = len(v3)
                                            if len(v3) > 0 {
                                                v3[0] = (v4 & (^(v4)))
                                            }
                                        } else {
                                            var v184 int16 = v2
                                            var v185 string = ""
                                            _ = v184
                                            _ = v185
                                        }
                                    }
                                } else {
                                    if false {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    } else {
                                        var v186 uint64 = 18446744073709551615
                                        _ = v186
                                    }
                                    var v187 int = (fi53 - -227)
                                    _ = v187
                                }
                                _ = v55
                            }
                            v3 = append(v3, v4)
                            if (v4 < (v4 >> ((42) & 63))) {
                                v188 := GenS1{F0: v4}
                                switch (fi53 & v15) {
                                case 147:
                                    switch ("foo bar" + ("a" + ("" + ""))) {
                                    case "foo bar":
                                        var v189 int32 = -1
                                        _ = v189
                                    }
                                case 32768:
                                    fmt.Println(v18)
                                }
                                _ = v188
                            }
                        }
                        if false {
                            if true {
                                switch "foo bar" {
                                case "hello":
                                    var v190 int = (-(v15))
                                    _ = v190
                                case "foo bar":
                                    if false {
                                        v5 = v1
                                    } else {
                                        var v191 bool = true
                                        _ = v191
                                    }
                                }
                            }
                            if false {
                                if false {
                                    fmt.Println(v15, v1)
                                } else {
                                    if false {
                                        _ = len(v3)
                                        v1 -= 2
                                        switch "\n\t" {
                                        case "":
                                            if len(v3) > 0 {
                                                v3[0] = (v4 & 0)
                                            }
                                        default:
                                            var v192 string = "\n\t"
                                            _ = v192
                                        }
                                    } else {
                                        for fi193 := 0; fi193 < 10; fi193++ {
                                            var v194 bool = true
                                            switch (("" + ("" + "\n\t")) + "hello") {
                                            case "a":
                                                var v195 uint64 = 810
                                                _ = v195
                                            case "\n\t":
                                                fmt.Println(fi193)
                                            case "foo bar":
                                                v4 = 2147483648
                                            }
                                            _ = fi193
                                            _ = v194
                                        }
                                        var v196 uint64 = 643
                                        _ = v196
                                    }
                                    var v197 int = int(v2)
                                    switch "" {
                                    case "hello":
                                        if true {
                                            var v198 int = 0
                                            var v199 uint64 = 42
                                            _ = v198
                                            _ = v199
                                        }
                                    case "":
                                        v200 := GenS17{F0: 1}
                                        _ = v200
                                    }
                                    _ = v197
                                }
                            } else {
                                var v201 string = "foo bar"
                                v202 := GenS6{F0: true, F1: "hello", F2: v1}
                                v201 = v201
                                _ = v201
                                _ = v202
                            }
                        } else {
                            v3 = append(v3, 2147483647)
                        }
                        _ = fi53
                    }
                    var v203 uint16 = 42
                    _ = v203
                }
            }
            _ = v34
        }
    default:
        v18 -= v18
    }
    v204 := GenS17{F0: -1}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "hello" {
    case "foo bar":
        v3 = append(v3, v4)
    case "":
        v3 = append(v3, uint32(v18))
    }
    var v205 bool = (0 < (v4 % ((956) | 1)))
    v2 = (v2 & 32767)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v205:", v205)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v204
    _ = v205
    _ = v3
    _ = v4
    _ = v5
}
