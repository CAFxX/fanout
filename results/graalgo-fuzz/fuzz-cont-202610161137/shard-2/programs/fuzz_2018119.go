package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int32; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 bool; F1 int32; F2 uint64 }
type GenS8 struct { F0 string; F1 int8; F2 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int8; F1 uint; F2 int32 }
type GenS11 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS12 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int64; F1 int; F2 uint16 }
type GenS16 struct { F0 int64; F1 int32 }
type GenS17 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 int8; F1 uint }
type GenS20 struct { F0 uint; F1 uint16 }
type GenS21 struct { F0 int16; F1 int32 }
type GenS22 struct { F0 uint; F1 uint16 }
type GenS23 struct { F0 uint16; F1 int; F2 bool }
type GenS24 struct { F0 int8; F1 bool }
type GenS25 struct { F0 int32; F1 int8 }

func main() {
    var v1 uint8 = 128
    var v2 int8 = -128
    var v3 int8 = v2
    v4 := []uint32{2, 556, 554, 65535}
    var v5 uint64 = 4294967295
    var v6 uint16 = 487
    for fi7 := 0; fi7 < 14; fi7++ {
        if false {
            switch fi7 {
            case 10:
                v8 := GenS1{F0: 0, F1: true}
                _ = v8
            case 32767:
                switch (-(fi7)) {
                case 32767:
                    if (303 < 9223372036854775807) {
                        for fi9 := 0; fi9 < 16; fi9++ {
                            if len(v4) > 0 {
                                v4[0] = 4294967295
                            }
                            switch "foo bar" {
                            case "":
                                var v10 int16 = 0
                                _ = v10
                            case "a":
                                if true {
                                    v11 := GenS1{F0: 1, F1: true}
                                    v12 := GenS2{F0: -928, F1: v6, F2: (-(v2))}
                                    _ = v11
                                    _ = v12
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = 775
                                    }
                                    v2 = v2
                                    var v13 uint16 = v6
                                    switch "hello" {
                                    case "foo bar":
                                        v14 := GenS3{F0: 42, F1: int32(v1)}
                                        _ = v14
                                    case "\n\t":
                                        v13 = (v6 << ((482) & 63))
                                    }
                                    _ = v13
                                }
                                var v15 int8 = (v2 - (v2 - 127))
                                _ = v15
                            case "foo bar":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                v16 := GenS4{F0: -1, F1: (v6 ^ 2)}
                                switch (((("a" + "a") + "") + "\n\t") + "a") {
                                case "foo bar":
                                    var v17 int16 = 239
                                    _ = v17
                                case "hello":
                                    switch (fi9 << ((121) & 63)) {
                                    case 781:
                                        switch "hello" {
                                        case "hello":
                                            var v18 uint = 42
                                            _ = v18
                                        }
                                    case 9223372036854775807:
                                        if (1e10 < -2.25) {
                                            var v19 float64 = 1.5
                                            var v20 uint = 2
                                            _ = v19
                                            _ = v20
                                        } else {
                                            var v21 float64 = 3.14159
                                            _ = v21
                                        }
                                    }
                                }
                                _ = v16
                            }
                            _ = fi9
                        }
                        v4 = append(v4, 65535)
                        var v22 int16 = 42
                        _ = v22
                    }
                }
            }
        } else {
            switch "" {
            case "":
                v2 -= -38
            case "\n\t":
                for fi23 := 0; fi23 < 17; fi23++ {
                    if (453 != v5) {
                        var v24 uint64 = (v5 & (v5 << ((0) & 63)))
                        _ = len(v4)
                        _ = v24
                    }
                    _ = fi23
                }
            }
        }
        _ = fi7
    }
    v3 -= v3
    fmt.Println(v6, v2)
    if (4294967295 < 0) {
        v1 -= v1
        if true {
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                switch "" {
                case "\n\t":
                    v6 = (v6 - (v6 * v6))
                case "foo bar":
                    fmt.Println(v6)
                default:
                    v6 = 600
                }
            case "a":
                fmt.Println(v3, v1, v6)
            }
            var v25 uint16 = v6
            switch 128 {
            case 1:
                var v26 uint = 724
                _ = v26
            }
            _ = v25
        }
    }
    if (861 >= -361) {
        v3 = 18
    } else {
        var v27 bool = (-2147483648 <= -2147483648)
        _ = v27
    }
    var v28 float32 = ((0.0 + float32(float64(3.14159))) - 1.5)
    v29 := GenS2{F0: 0, F1: v6, F2: v3}
    switch "a" {
    case "foo bar":
        v6 = (v6 ^ uint16(v5))
    case "a":
        var v30 int16 = 817
        _ = v30
    }
    fmt.Println(v5)
    for fi31 := 0; fi31 < 4; fi31++ {
        for fi32 := 0; fi32 < 19; fi32++ {
            v4 = append(v4, 359)
            var v33 uint8 = 128
            var v34 uint = 2
            _ = fi32
            _ = v33
            _ = v34
        }
        switch "\n\t" {
        case "\n\t":
            v6 = 596
        case "a":
            if true {
                v4 = append(v4, 1)
            }
        case "foo bar":
            v1 += v1
        default:
            var v35 int64 = 422
            _ = v35
        }
        _ = fi31
    }
    fmt.Println(v1)
    for fi36 := 0; fi36 < 14; fi36++ {
        switch -203 {
        case -107:
            if (false || (float32((0.0 * 0.0)) != float32(0.0))) {
                fmt.Println(v1)
            } else {
                v2 = 0
                switch (fi36 % (((fi36 ^ (fi36 * (-(fi36))))) | 1)) {
                case 32768:
                    for fi37 := 0; fi37 < 5; fi37++ {
                        switch ("hello" + (("foo bar" + "a") + "foo bar")) {
                        case "":
                            if true {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            } else {
                                var v38 int8 = (v2 & 1)
                                var v39 float32 = float32(0.0)
                                fmt.Println(v3, v3, fi37)
                                _ = v38
                                _ = v39
                            }
                        default:
                            var v40 uint16 = 1
                            _ = v40
                        }
                        _ = fi37
                    }
                case 42:
                    if false {
                        v41 := GenS5{F0: (v3 << ((65535) & 63)), F1: (true || (0.0 >= ((float64(v28) / (((1e10 + 1.5) - float64((v28 * (0.1 - 0.1)))) - 0.1)) + (1.5 + (3.14159 / 1.5)))))}
                        _ = v41
                    }
                default:
                    if len(v4) > 0 {
                        v4[0] = 2147483647
                    }
                }
                var v42 float32 = v28
                _ = v42
            }
        case 127:
            for fi43 := 0; fi43 < 1; fi43++ {
                v6 += (v6 << ((39) & 63))
                switch "hello" {
                case "\n\t":
                    var v44 uint16 = (v6 - v6)
                    _ = v44
                default:
                    if false {
                        v45 := GenS6{F0: -822, F1: 457}
                        if true {
                            v2 += (-(v2))
                            v46 := GenS4{F0: int32(v2), F1: 593}
                            _ = v46
                        } else {
                            var v47 float64 = 0.0
                            _ = v47
                        }
                        switch "hello" {
                        case "\n\t":
                            if (true || true) {
                                var v48 int8 = v3
                                v49 := GenS7{F0: false, F1: 2147483647, F2: 989}
                                switch ("" + ((("hello" + "a") + "\n\t") + "")) {
                                case "hello":
                                    var v50 int = (^(fi43))
                                    _ = v50
                                default:
                                    var v51 uint = 4294967295
                                    _ = v51
                                }
                                _ = v48
                                _ = v49
                            }
                        case "foo bar":
                            if (-1 < 929) {
                                var v52 uint64 = 18446744073709551615
                                v52 = v5
                                _ = v52
                            }
                        case "a":
                            v4 = append(v4, 4294967295)
                        }
                        v53 := GenS5{F0: v2, F1: false}
                        _ = v45
                        _ = v53
                    }
                    var v54 float32 = v28
                    _ = v54
                }
                if true {
                    v4 = append(v4, 4294967295)
                    v6 = v6
                    v55 := GenS8{F0: ((("\n\t" + ("hello" + (("" + "") + (("a" + "foo bar") + "a")))) + ("a" + "\n\t")) + ""), F1: -128, F2: 238}
                    var v56 float64 = (0.0 * ((3.14159 - (float64((((v28 * (3.14159 / v28)) / 1e10) / v28)) / 1.0)) + 0.1))
                    if false {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        var v57 uint16 = (v6 + v6)
                        switch "foo bar" {
                        case "foo bar":
                            if (42 == (v6 - (v6 << ((1) & 63)))) {
                                if (18446744073709551615 != 65535) {
                                    switch fi43 {
                                    case 32768:
                                        if (((fi36 * (^(fi36))) != (^(fi36))) || (v5 > (v5 | v5))) {
                                            var v58 uint8 = v1
                                            _ = v58
                                        } else {
                                            switch ("foo bar" + "foo bar") {
                                            case "a":
                                                switch fi36 {
                                                case -1:
                                                    if (!(!false)) {
                                                        if (!(!false)) {
                                                            var v59 int8 = (v2 >> ((42) & 63))
                                                            switch (-(fi36)) {
                                                            case 692:
                                                                var v60 int8 = (v3 << ((242) & 63))
                                                                _ = v60
                                                            case 9223372036854775807:
                                                                if false {
                                                                    v2 = v2
                                                                    if true {
                                                                        v4 = append(v4, 2147483648)
                                                                        switch fi43 {
                                                                        case 32767:
                                                                            var v61 int = fi36
                                                                            _ = v61
                                                                        case 0:
                                                                            if (-9223372036854775808 >= 25) {
                                                                                fmt.Println(fi36, v59, v57)
                                                                            }
                                                                        case 42:
                                                                            var v62 int = (-(fi43))
                                                                            _ = v62
                                                                        }
                                                                        if false {
                                                                            var v63 string = ""
                                                                            _ = v63
                                                                        }
                                                                        var v64 uint8 = (v1 << ((uint(v1)) & 63))
                                                                        _ = v64
                                                                    } else {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        v57 = 1
                                                                        var v65 float64 = v56
                                                                        v57 = (v57 % ((uint16(v1)) | 1))
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            v66 := GenS9{F0: "\n\t"}
                                                                            _ = v66
                                                                        case "hello":
                                                                            v56 = float64(float32(v65))
                                                                        case "a":
                                                                            v3 -= v59
                                                                        default:
                                                                            switch (^(fi36)) {
                                                                            case 128:
                                                                                var v67 uint16 = 32767
                                                                                _ = v67
                                                                            default:
                                                                                switch -413 {
                                                                                case 668:
                                                                                    var v68 float64 = -2.25
                                                                                    _ = v68
                                                                                case 128:
                                                                                    v59 = (v59 & v2)
                                                                                case 976:
                                                                                    var v69 string = "a"
                                                                                    _ = v69
                                                                                default:
                                                                                    v59 = (v59 ^ v3)
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v65
                                                                    }
                                                                    v70 := GenS10{F0: v59, F1: 65535, F2: int32(v57)}
                                                                    var v71 float64 = v56
                                                                    _ = v70
                                                                    _ = v71
                                                                }
                                                            default:
                                                                v72 := GenS1{F0: 9223372036854775807, F1: false}
                                                                _ = v72
                                                            }
                                                            _ = v59
                                                        } else {
                                                            var v73 uint32 = 65535
                                                            var v74 uint16 = (v6 << ((18446744073709551615) & 63))
                                                            _ = v73
                                                            _ = v74
                                                        }
                                                        fmt.Println(fi43)
                                                        v6 = 65535
                                                    } else {
                                                        switch "" {
                                                        case "foo bar":
                                                            v75 := GenS11{F0: 255, F1: (v1 / ((42) | 1)), F2: -2147483648}
                                                            _ = v75
                                                        case "":
                                                            fmt.Println(v1)
                                                        default:
                                                            v2 = (v3 - v2)
                                                        }
                                                    }
                                                case -9223372036854775808:
                                                    if (v56 == (float64(v28) - v56)) {
                                                        if (v56 >= 3.14159) {
                                                            var v76 uint64 = (v5 ^ 0)
                                                            if true {
                                                                if true {
                                                                    v77 := GenS12{F0: v76, F1: (v1 / (((v1 | v1)) | 1)), F2: v3}
                                                                    var v78 int8 = v2
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 4294967295
                                                                    }
                                                                    v57 -= (v6 >> ((18446744073709551615) & 63))
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    switch (fi36 << ((255) & 63)) {
                                                                    case 1:
                                                                        v2 = (v3 & int8(v57))
                                                                    case 9223372036854775807:
                                                                        v1 = (v1 - (v1 % ((v1) | 1)))
                                                                    default:
                                                                        v79 := GenS13{F0: fi36}
                                                                        _ = v79
                                                                    }
                                                                    var v80 float32 = 1.5
                                                                    if (true && false) {
                                                                        fmt.Println(v56)
                                                                        switch (fi43 / (((fi36 - (fi36 & int(v1)))) | 1)) {
                                                                        case 0:
                                                                            v81 := GenS12{F0: v76, F1: v1, F2: 11}
                                                                            _ = v81
                                                                        case 127:
                                                                            switch fi43 {
                                                                            case -1:
                                                                                if true {
                                                                                    var v82 int = (fi36 % (((-(fi43))) | 1))
                                                                                    fmt.Println(v57)
                                                                                    switch v82 {
                                                                                    case -1:
                                                                                        fmt.Println(v78, v6, v2)
                                                                                    default:
                                                                                        var v83 uint8 = v1
                                                                                        _ = v83
                                                                                    }
                                                                                    var v84 int = fi36
                                                                                    switch "\n\t" {
                                                                                    case "":
                                                                                        v57 = v6
                                                                                    default:
                                                                                        v85 := GenS14{F0: 203}
                                                                                        _ = v85
                                                                                    }
                                                                                    var v86 uint64 = (v76 | (v76 * v5))
                                                                                    _ = v82
                                                                                    _ = v84
                                                                                    _ = v86
                                                                                } else {
                                                                                    var v87 float64 = (v56 + v56)
                                                                                    if true {
                                                                                        if true {
                                                                                            switch (fi36 / (((fi43 ^ (fi43 / ((1) | 1)))) | 1)) {
                                                                                            case 265:
                                                                                                v80 = 0.0
                                                                                            case 738:
                                                                                                fmt.Println(v57)
                                                                                            }
                                                                                            var v88 int64 = 1
                                                                                            var v89 uint8 = (v1 * (^(v1)))
                                                                                            if false {
                                                                                                switch "" {
                                                                                                case "hello":
                                                                                                    var v90 uint8 = uint8(v76)
                                                                                                    _ = v90
                                                                                                case "\n\t":
                                                                                                    var v91 uint8 = v89
                                                                                                    _ = v91
                                                                                                default:
                                                                                                    v92 := GenS15{F0: v88, F1: -9223372036854775808, F2: 32767}
                                                                                                    _ = v92
                                                                                                }
                                                                                                var v93 string = "a"
                                                                                                _ = v93
                                                                                            }
                                                                                            fmt.Println(v78)
                                                                                            _ = v88
                                                                                            _ = v89
                                                                                        } else {
                                                                                            v57 = (v57 >> ((18446744073709551615) & 63))
                                                                                            if true {
                                                                                                v2 = (v78 % (((v2 >> ((0) & 63))) | 1))
                                                                                            }
                                                                                            var v94 uint32 = 2147483648
                                                                                            switch int(v78) {
                                                                                            case 32768:
                                                                                                _ = len(v4)
                                                                                            case 1:
                                                                                                switch (-(fi36)) {
                                                                                                case -59:
                                                                                                    if (!true) {
                                                                                                        v95 := GenS7{F0: true, F1: 1, F2: (^(v5))}
                                                                                                        v78 -= -128
                                                                                                        var v96 uint = 65535
                                                                                                        var v97 float32 = 0.1
                                                                                                        var v98 float32 = (((v97 / (v80 / ((float32(((3.14159 - v56) * (3.14159 * v56))) + (-2.25 / v97)) - float32(v56)))) - (0.0 + 0.1)) / 3.14159)
                                                                                                        _ = v95
                                                                                                        _ = v96
                                                                                                        _ = v97
                                                                                                        _ = v98
                                                                                                    } else {
                                                                                                        switch "foo bar" {
                                                                                                        case "hello":
                                                                                                            v76 += v5
                                                                                                        default:
                                                                                                            var v99 int16 = -32768
                                                                                                            _ = v99
                                                                                                        }
                                                                                                    }
                                                                                                case -575:
                                                                                                    fmt.Println(v28, v78)
                                                                                                case 2147483647:
                                                                                                    switch (fi43 * fi36) {
                                                                                                    case 128:
                                                                                                        if true {
                                                                                                            v100 := GenS5{F0: (v2 + (v3 * (v2 + (^(v3))))), F1: (false || (false && true))}
                                                                                                            if false {
                                                                                                                if len(v4) > 0 {
                                                                                                                    v4[0] = (v94 % (((v94 & (v94 | 2147483648))) | 1))
                                                                                                                }
                                                                                                            }
                                                                                                            v101 := GenS2{F0: 1, F1: (v57 & 2), F2: int8(v94)}
                                                                                                            if (3.14159 > (v87 / v87)) {
                                                                                                                v80 = 1e10
                                                                                                                if (-2147483648 != 1) {
                                                                                                                    if true {
                                                                                                                        if true {
                                                                                                                            var v102 uint16 = v6
                                                                                                                            _ = v102
                                                                                                                        }
                                                                                                                        var v103 float32 = float32(v56)
                                                                                                                        _ = len(v4)
                                                                                                                        if true {
                                                                                                                            if (-9223372036854775808 > -967) {
                                                                                                                                v87 += ((v87 + (v87 * v56)) / (v87 / 1.0))
                                                                                                                                var v104 uint8 = 137
                                                                                                                                _ = v104
                                                                                                                            }
                                                                                                                            v94 = 511
                                                                                                                        }
                                                                                                                        var v105 uint16 = 42
                                                                                                                        _ = v103
                                                                                                                        _ = v105
                                                                                                                    }
                                                                                                                    if (9223372036854775807 < int64(v57)) {
                                                                                                                        v4 = append(v4, (v94 << ((4294967295) & 63)))
                                                                                                                        v4 = append(v4, (^(v94)))
                                                                                                                        switch "" {
                                                                                                                        case "\n\t":
                                                                                                                            fmt.Println(v76, v6, v76)
                                                                                                                        case "foo bar":
                                                                                                                            var v106 uint64 = v76
                                                                                                                            _ = v106
                                                                                                                        case "a":
                                                                                                                            switch (fi43 - fi36) {
                                                                                                                            case 32768:
                                                                                                                                v107 := GenS13{F0: 2147483647}
                                                                                                                                _ = v107
                                                                                                                            case -72:
                                                                                                                                v108 := GenS3{F0: 397, F1: -2147483648}
                                                                                                                                _ = v108
                                                                                                                            case -9223372036854775808:
                                                                                                                                v78 = v2
                                                                                                                            default:
                                                                                                                                if (v2 <= (v3 % ((int8(v94)) | 1))) {
                                                                                                                                    if true {
                                                                                                                                        v3 = 100
                                                                                                                                        v3 = 127
                                                                                                                                        v5 = v5
                                                                                                                                    } else {
                                                                                                                                        v109 := GenS16{F0: -1, F1: 0}
                                                                                                                                        fmt.Println(v5, v80, v6)
                                                                                                                                        if (!(v76 <= (^(v76)))) {
                                                                                                                                            v87 -= ((v87 - ((v87 * v56) / 1.5)) / float64(v80))
                                                                                                                                            switch (fi43 >> ((341) & 63)) {
                                                                                                                                            case -214:
                                                                                                                                                fmt.Println(v57, v28, v57)
                                                                                                                                            case 42:
                                                                                                                                                switch "a" {
                                                                                                                                                case "hello":
                                                                                                                                                    v110 := GenS7{F0: false, F1: -581, F2: (v76 % ((v5) | 1))}
                                                                                                                                                    _ = v110
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            var v111 uint32 = v94
                                                                                                                                            _ = v111
                                                                                                                                        }
                                                                                                                                        _ = v109
                                                                                                                                    }
                                                                                                                                    switch ("a" + "") {
                                                                                                                                    case "\n\t":
                                                                                                                                        switch 1 {
                                                                                                                                        case 500:
                                                                                                                                            var v112 float32 = (v28 - (((float32((v56 + (1e10 * (v87 - 3.14159)))) + (0.0 - 0.0)) - -2.25) - v28))
                                                                                                                                            _ = v112
                                                                                                                                        case 719:
                                                                                                                                            var v113 uint8 = v1
                                                                                                                                            _ = v113
                                                                                                                                        case 145:
                                                                                                                                            fmt.Println(v5, v87, v5)
                                                                                                                                        default:
                                                                                                                                            if false {
                                                                                                                                                switch (fi36 ^ (fi43 ^ (^(fi36)))) {
                                                                                                                                                case 32767:
                                                                                                                                                    v3 -= v78
                                                                                                                                                case -9223372036854775808:
                                                                                                                                                    v114 := GenS6{F0: 875, F1: 9223372036854775807}
                                                                                                                                                    _ = v114
                                                                                                                                                default:
                                                                                                                                                    if false {
                                                                                                                                                        if len(v4) > 0 {
                                                                                                                                                            v4[0] = 4294967295
                                                                                                                                                        }
                                                                                                                                                        switch "\n\t" {
                                                                                                                                                        case "a":
                                                                                                                                                            var v115 int32 = 154
                                                                                                                                                            _ = v115
                                                                                                                                                        case "hello":
                                                                                                                                                            v78 = v2
                                                                                                                                                            v1 += (^(v1))
                                                                                                                                                            v116 := GenS14{F0: 42}
                                                                                                                                                            _ = v116
                                                                                                                                                        }
                                                                                                                                                    } else {
                                                                                                                                                        var v117 bool = true
                                                                                                                                                        var v118 int32 = -1
                                                                                                                                                        if len(v4) > 0 {
                                                                                                                                                            v4[0] = (v94 + (^(v94)))
                                                                                                                                                        }
                                                                                                                                                        _ = v117
                                                                                                                                                        _ = v118
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                                switch "" {
                                                                                                                                                case "hello":
                                                                                                                                                    var v119 uint64 = v76
                                                                                                                                                    _ = v119
                                                                                                                                                case "a":
                                                                                                                                                    v4 = append(v4, v94)
                                                                                                                                                }
                                                                                                                                                fmt.Println(v57)
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    case "hello":
                                                                                                                                        v4 = append(v4, v94)
                                                                                                                                    case "a":
                                                                                                                                        v28 += ((v28 - 1e10) - v28)
                                                                                                                                    default:
                                                                                                                                        v120 := GenS17{F0: 42, F1: (((v1 * (v1 * 63)) <= 127) && true), F2: (v57 - (^(v57)))}
                                                                                                                                        var v121 int64 = 1
                                                                                                                                        _ = v120
                                                                                                                                        _ = v121
                                                                                                                                    }
                                                                                                                                    v1 = v1
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    if (1.5 > 0.0) {
                                                                                                                        if ((v1 * (v1 / ((v1) | 1))) != 137) {
                                                                                                                            v5 = v76
                                                                                                                            var v122 float64 = float64(v28)
                                                                                                                            _ = v122
                                                                                                                        } else {
                                                                                                                            fmt.Println(v5)
                                                                                                                        }
                                                                                                                        v123 := GenS16{F0: -9223372036854775808, F1: 1}
                                                                                                                        var v124 float32 = v80
                                                                                                                        _ = v123
                                                                                                                        _ = v124
                                                                                                                    }
                                                                                                                    if len(v4) > 0 {
                                                                                                                        v4[0] = (v94 | (v94 ^ (v94 << ((255) & 63))))
                                                                                                                    }
                                                                                                                    v56 = v87
                                                                                                                }
                                                                                                                v125 := GenS18{F0: (^(v57))}
                                                                                                                _ = v125
                                                                                                            }
                                                                                                            switch 0 {
                                                                                                            case 509:
                                                                                                                fmt.Println(v57)
                                                                                                            case 559:
                                                                                                                v5 = 4294967295
                                                                                                            }
                                                                                                            _ = v100
                                                                                                            _ = v101
                                                                                                        }
                                                                                                    case 32768:
                                                                                                        v80 += v80
                                                                                                    case 9223372036854775807:
                                                                                                        v4 = append(v4, v94)
                                                                                                        v87 -= -2.25
                                                                                                        if len(v4) >= 2 {
                                                                                                            v4 = v4[1:len(v4)]
                                                                                                        }
                                                                                                    default:
                                                                                                        if len(v4) >= 2 {
                                                                                                            v4 = v4[1:len(v4)]
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            case -934:
                                                                                                fmt.Println(v80, v3)
                                                                                            }
                                                                                            _ = v94
                                                                                        }
                                                                                        fmt.Println(v78, fi43, v6)
                                                                                    } else {
                                                                                        fmt.Println(fi43, v78)
                                                                                        if len(v4) > 0 {
                                                                                            v4[0] = 2147483648
                                                                                        }
                                                                                        switch -9223372036854775808 {
                                                                                        case -2147483648:
                                                                                            v5 -= v76
                                                                                            fmt.Println(v28)
                                                                                        default:
                                                                                            v2 -= (v2 ^ (v2 * v3))
                                                                                            switch (fi36 + -2147483648) {
                                                                                            case -2147483648:
                                                                                                fmt.Println(v87)
                                                                                            case 42:
                                                                                                v126 := GenS19{F0: v78, F1: 940}
                                                                                                _ = v126
                                                                                            case 32767:
                                                                                                v4 = append(v4, uint32(v5))
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = v87
                                                                                }
                                                                            default:
                                                                                if true {
                                                                                    if (true && ((fi36 & int(v1)) != -2147483648)) {
                                                                                        v127 := GenS20{F0: 1, F1: (v6 / ((0) | 1))}
                                                                                        _ = v127
                                                                                    } else {
                                                                                        fmt.Println(v56, v3, v78)
                                                                                    }
                                                                                    var v128 uint = 265
                                                                                    _ = v128
                                                                                } else {
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        var v129 float32 = ((v80 * 3.14159) / 1e10)
                                                                        _ = v129
                                                                    } else {
                                                                        if false {
                                                                            switch (^(fi36)) {
                                                                            case 127:
                                                                                if (!(!false)) {
                                                                                    if (!true) {
                                                                                        var v130 uint32 = 2
                                                                                        _ = v130
                                                                                    }
                                                                                    if (!(("" + ("a" + ("hello" + "foo bar"))) < ("foo bar" + ("a" + "\n\t")))) {
                                                                                        var v131 bool = (!true)
                                                                                        _ = v131
                                                                                    }
                                                                                    var v132 string = "hello"
                                                                                    _ = v132
                                                                                } else {
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = 480
                                                                                    }
                                                                                }
                                                                            case -822:
                                                                                var v133 uint = 101
                                                                                _ = v133
                                                                            default:
                                                                                var v134 uint64 = (^(v5))
                                                                                _ = v134
                                                                            }
                                                                            var v135 float64 = float64(v80)
                                                                            _ = v135
                                                                        }
                                                                    }
                                                                    _ = v77
                                                                    _ = v78
                                                                    _ = v80
                                                                } else {
                                                                    fmt.Println(v56, v2, v56)
                                                                    v136 := GenS21{F0: -1, F1: 42}
                                                                    fmt.Println(fi36)
                                                                    _ = v136
                                                                }
                                                                v137 := GenS9{F0: "\n\t"}
                                                                _ = v137
                                                            } else {
                                                                if ((("hello" + "foo bar") + "\n\t") == "hello") {
                                                                    v5 += (v76 % ((v76) | 1))
                                                                }
                                                            }
                                                            _ = v76
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    case -612:
                                        v138 := GenS16{F0: 1, F1: -455}
                                        _ = v138
                                    }
                                }
                            } else {
                                v139 := GenS9{F0: ("\n\t" + "\n\t")}
                                v140 := GenS22{F0: 4294967295, F1: (v57 - 65535)}
                                _ = v139
                                _ = v140
                            }
                        case "hello":
                            v1 = (v1 * v1)
                        case "\n\t":
                            v141 := GenS23{F0: v57, F1: fi43, F2: false}
                            _ = v141
                        default:
                            fmt.Println(fi43, v2, fi36)
                        }
                        v57 = v57
                        _ = v57
                    }
                    v3 -= v3
                    v142 := GenS24{F0: 124, F1: true}
                    v6 += v6
                    _ = v142
                    _ = v55
                    _ = v56
                } else {
                    v143 := GenS25{F0: 2147483647, F1: v3}
                    v1 = (v1 + (v1 + v1))
                    _ = v143
                }
                _ = fi43
            }
        case -2147483648:
            v3 -= v2
        }
        var v144 uint8 = (v1 + (v1 * (v1 * (v1 * (v1 >> ((391) & 63))))))
        var v145 uint = 613
        _ = fi36
        _ = v144
        _ = v145
    }
    var v146 string = ("" + "hello")
    var v147 float32 = float32((float64(1e10) * 0.0))
    fmt.Println("v1:", v1)
    fmt.Println("v146:", v146)
    fmt.Println("v147:", v147)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v146
    _ = v147
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
