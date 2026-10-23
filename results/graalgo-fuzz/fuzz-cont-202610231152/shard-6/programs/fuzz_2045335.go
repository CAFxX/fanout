package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int; F1 int64; F2 int16 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 int16; F2 int64 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int16; F1 uint8 }
type GenS9 struct { F0 int8; F1 uint8; F2 int8 }
type GenS10 struct { F0 int8; F1 uint16 }

func main() {
    v1 := []string{"hello"}
    var v2 bool = (!(-2147483648 > 2147483647))
    var v3 float32 = 1.5
    var v4 float64 = (0.0 - 1e10)
    var v5 int = 127
    v6 := []string{("hello" + "hello"), "", "foo bar", (("hello" + ("\n\t" + "foo bar")) + "\n\t")}
    switch (v5 % ((v5) | 1)) {
    case 128:
        if (2 <= 2) {
            switch (("" + ("foo bar" + ("hello" + "foo bar"))) + "") {
            case "":
                v2 = false
            case "\n\t":
                if (!(72 > -59)) {
                    switch "hello" {
                    case "\n\t":
                        var v7 float32 = v3
                        _ = v7
                    case "foo bar":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    default:
                        for fi8 := 0; fi8 < 1; fi8++ {
                            var v9 float32 = v3
                            v10 := GenS1{F0: 127, F1: -1}
                            fmt.Println(fi8, v9)
                            _ = fi8
                            _ = v10
                            _ = v9
                        }
                    }
                }
            }
            var v11 int16 = 62
            _ = v11
        }
    case 4294967295:
        if (!(42 > -675)) {
            var v12 uint16 = 0
            _ = v12
        } else {
            switch ("foo bar" + ("a" + "\n\t")) {
            case "a":
                for fi13 := 0; fi13 < 19; fi13++ {
                    if len(v1) > 0 {
                        v1[0] = ("\n\t" + "foo bar")
                    }
                    switch (v5 / (((v5 ^ (v5 & (^(v5))))) | 1)) {
                    case -957:
                        v14 := GenS2{F0: (fi13 << ((42) & 63)), F1: 251, F2: 1000}
                        _ = v14
                    default:
                        _ = len(v1)
                    }
                    _ = fi13
                }
            default:
                if true {
                    switch ("foo bar" + "foo bar") {
                    case "a":
                        fmt.Println(v4)
                    case "foo bar":
                        for fi15 := 0; fi15 < 16; fi15++ {
                            var v16 uint8 = 2
                            v2 = false
                            if v2 {
                                for fi17 := 0; fi17 < 5; fi17++ {
                                    var v18 float32 = float32(1.5)
                                    if v2 {
                                        var v19 int = (v5 - (v5 & -1))
                                        v16 -= v16
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v19
                                    } else {
                                        var v20 string = ("" + ("" + "foo bar"))
                                        fmt.Println(fi15, v16, v2)
                                        v21 := GenS1{F0: v16, F1: -648}
                                        _ = v20
                                        _ = v21
                                    }
                                    switch int(v16) {
                                    case 861:
                                        v5 = -9223372036854775808
                                    case 127:
                                        v22 := GenS2{F0: (v5 + 4294967295), F1: -1, F2: 1}
                                        _ = v22
                                    case -250:
                                        if v2 {
                                            v23 := GenS3{F0: 255, F1: 65535}
                                            fmt.Println(v16)
                                            _ = v23
                                        }
                                    }
                                    _ = fi17
                                    _ = v18
                                }
                                fmt.Println(v16, fi15)
                            } else {
                                switch "a" {
                                case "\n\t":
                                    switch (fi15 << ((42) & 63)) {
                                    case 9223372036854775807:
                                        if (v3 == float32(((float64(v3) - 0.1) + (v4 + ((float64((((v3 / (v3 + float32(v4))) / float32(v4)) + (v3 + 3.14159))) + v4) + v4))))) {
                                            if (v2 && v2) {
                                                fmt.Println(fi15, v3, v3)
                                                v16 += v16
                                                v2 = (-2.25 <= -2.25)
                                            }
                                            var v24 uint64 = 42
                                            _ = v24
                                        } else {
                                            v25 := GenS1{F0: v16, F1: 0}
                                            var v26 bool = v2
                                            var v27 float64 = 3.14159
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                        }
                                    case 1:
                                        v1 = append(v1, "foo bar")
                                    case 127:
                                        for fi28 := 0; fi28 < 15; fi28++ {
                                            v3 = 1e10
                                            switch 32768 {
                                            case -526:
                                                var v29 uint16 = 32767
                                                _ = v29
                                            case 0:
                                                switch (fi15 >> ((2) & 63)) {
                                                case 32767:
                                                    if v2 {
                                                        fmt.Println(v16, v5, v4)
                                                        fmt.Println(v5, fi15)
                                                        if len(v6) > 0 {
                                                            v6[0] = "a"
                                                        }
                                                    } else {
                                                        v5 += (^(fi15))
                                                        if false {
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                v16 += 42
                                                            case "\n\t":
                                                                var v30 uint = 2
                                                                _ = v30
                                                            }
                                                            var v31 bool = v2
                                                            var v32 uint8 = 128
                                                            _ = v31
                                                            _ = v32
                                                        }
                                                        v3 += float32(v4)
                                                    }
                                                default:
                                                    switch (fi15 & 4294967295) {
                                                    case 4294967295:
                                                        _ = len(v6)
                                                    case 127:
                                                        fmt.Println(fi28)
                                                    default:
                                                        v33 := GenS4{F0: 18446744073709551615}
                                                        _ = v33
                                                    }
                                                }
                                            case -301:
                                                fmt.Println(v16, fi15)
                                            default:
                                                if v2 {
                                                    v34 := GenS1{F0: v16, F1: 1000}
                                                    fmt.Println(v5, v3)
                                                    _ = v34
                                                }
                                            }
                                            _ = fi28
                                        }
                                    }
                                case "hello":
                                    for fi35 := 0; fi35 < 13; fi35++ {
                                        v36 := GenS2{F0: 127, F1: 744, F2: 1000}
                                        _ = fi35
                                        _ = v36
                                    }
                                case "a":
                                    var v37 int16 = 42
                                    _ = v37
                                }
                                if v2 {
                                    if len(v1) > 0 {
                                        v1[0] = "foo bar"
                                    }
                                    fmt.Println(v5, v3, v2)
                                }
                            }
                            _ = fi15
                            _ = v16
                        }
                    default:
                        switch v5 {
                        case 42:
                            if v2 {
                                _ = len(v1)
                                if (0 > 9223372036854775807) {
                                    for fi38 := 0; fi38 < 10; fi38++ {
                                        for fi39 := 0; fi39 < 14; fi39++ {
                                            if false {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                fmt.Println(fi38)
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            _ = fi39
                                        }
                                        if v2 {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v4 -= v4
                                        }
                                        _ = fi38
                                    }
                                    for fi40 := 0; fi40 < 1; fi40++ {
                                        if false {
                                            switch 128 {
                                            case 1:
                                                v4 = v4
                                            }
                                            for fi41 := 0; fi41 < 19; fi41++ {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    if (v2 && v2) {
                                                        switch -262 {
                                                        case 0:
                                                            switch 1 {
                                                            case -804:
                                                                v1 = append(v1, ("foo bar" + "foo bar"))
                                                            case -2147483648:
                                                                fmt.Println(v5, fi40)
                                                            default:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        default:
                                                            if false {
                                                                var v42 uint = 255
                                                                _ = len(v6)
                                                                fmt.Println(v42)
                                                                _ = v42
                                                            }
                                                        }
                                                        switch (^(v5)) {
                                                        case -9223372036854775808:
                                                            v4 = 0.1
                                                        }
                                                    } else {
                                                        v2 = false
                                                        var v43 uint8 = 127
                                                        if (((-2.25 / v3) + v3) <= (0.1 - v3)) {
                                                            var v44 uint8 = v43
                                                            v43 = (v43 ^ (v43 >> ((42) & 63)))
                                                            if (fi41 >= (fi40 * (fi40 ^ (fi40 % ((fi41) | 1))))) {
                                                                switch -36 {
                                                                case -1:
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        if len(v6) > 0 {
                                                                            v6[0] = "hello"
                                                                        }
                                                                    case "hello":
                                                                        switch fi40 {
                                                                        case 2147483647:
                                                                            var v45 int64 = 0
                                                                            _ = v45
                                                                        case 32767:
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                            var v46 int64 = -1
                                                                            v4 -= float64(float32(1e10))
                                                                            _ = v46
                                                                        }
                                                                    case "foo bar":
                                                                        var v47 int8 = 42
                                                                        if ((4294967295 <= 951) || v2) {
                                                                            _ = len(v6)
                                                                            switch (fi41 >> ((973) & 63)) {
                                                                            case -438:
                                                                                v44 += 255
                                                                            default:
                                                                                _ = len(v6)
                                                                                var v48 uint = 239
                                                                                _ = v48
                                                                            }
                                                                            var v49 bool = (v47 >= v47)
                                                                            _ = v49
                                                                        }
                                                                        var v50 int16 = -73
                                                                        _ = v47
                                                                        _ = v50
                                                                    }
                                                                case -628:
                                                                    v51 := GenS5{F0: 42, F1: -1, F2: -1}
                                                                    _ = v51
                                                                default:
                                                                    if true {
                                                                        if (false && v2) {
                                                                            v52 := GenS2{F0: 0, F1: -1, F2: -1}
                                                                            _ = v52
                                                                        }
                                                                        var v53 int8 = 42
                                                                        if v2 {
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                        } else {
                                                                            var v54 uint = 732
                                                                            _ = v54
                                                                        }
                                                                        _ = v53
                                                                    }
                                                                }
                                                            } else {
                                                                if v2 {
                                                                    var v55 string = ("a" + "a")
                                                                    fmt.Println(v43, fi41, v3)
                                                                    _ = v55
                                                                }
                                                            }
                                                            _ = v44
                                                        }
                                                        _ = v43
                                                    }
                                                }
                                                var v56 int8 = -1
                                                var v57 uint16 = 295
                                                v58 := GenS5{F0: 65535, F1: 42, F2: 0}
                                                v57 = (v57 % ((v57) | 1))
                                                v2 = (uint(v57) != 59)
                                                _ = fi41
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            }
                                        } else {
                                            v4 = (3.14159 - -2.25)
                                        }
                                        v6 = append(v6, "foo bar")
                                        for fi59 := 0; fi59 < 20; fi59++ {
                                            v3 = (0.1 - (((v3 + v3) + 3.14159) * ((0.0 + v3) / (1.5 + (float32(1.5) + -2.25)))))
                                            var v60 int8 = 42
                                            switch fi59 {
                                            case 9223372036854775807:
                                                switch fi40 {
                                                case 32767:
                                                    v2 = v2
                                                case 9223372036854775807:
                                                    switch "\n\t" {
                                                    case "":
                                                        fmt.Println(fi59, v5)
                                                    }
                                                case 128:
                                                    v4 += float64(v3)
                                                }
                                            }
                                            _ = fi59
                                            _ = v60
                                        }
                                        _ = fi40
                                    }
                                    for fi61 := 0; fi61 < 17; fi61++ {
                                        fmt.Println(v2)
                                        v4 += v4
                                        _ = fi61
                                    }
                                }
                                v6 = append(v6, "hello")
                            } else {
                                for fi62 := 0; fi62 < 13; fi62++ {
                                    if v2 {
                                        if v2 {
                                            v63 := GenS2{F0: (fi62 % ((32767) | 1)), F1: -1, F2: -1}
                                            _ = v63
                                        }
                                        v4 = v4
                                        for fi64 := 0; fi64 < 11; fi64++ {
                                            v65 := GenS2{F0: (-(fi62)), F1: 167, F2: 251}
                                            _ = fi64
                                            _ = v65
                                        }
                                    }
                                    _ = fi62
                                }
                            }
                        case 4294967295:
                            fmt.Println(v2, v4, v4)
                        default:
                            for fi66 := 0; fi66 < 1; fi66++ {
                                for fi67 := 0; fi67 < 20; fi67++ {
                                    var v68 bool = (v3 < v3)
                                    _ = fi67
                                    _ = v68
                                }
                                if true {
                                    fmt.Println(v3)
                                    var v69 uint = 0
                                    fmt.Println(v69)
                                    _ = v69
                                }
                                v70 := GenS2{F0: fi66, F1: 0, F2: 32767}
                                _ = fi66
                                _ = v70
                            }
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                } else {
                    if (!true) {
                        if v2 {
                            v2 = (float32(v4) <= 1e10)
                            v3 -= v3
                        }
                        v71 := GenS5{F0: 2147483647, F1: -980, F2: -315}
                        v72 := GenS3{F0: 42, F1: 1}
                        _ = v71
                        _ = v72
                    } else {
                        if false {
                            var v73 uint32 = 2
                            if (32 > -9223372036854775808) {
                                var v74 int64 = -9223372036854775808
                                if v2 {
                                    v75 := GenS5{F0: 2, F1: 17, F2: 1}
                                    _ = v75
                                } else {
                                    v1 = append(v1, ("" + ("a" + "foo bar")))
                                }
                                _ = v74
                            }
                            _ = v73
                        } else {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v76 float32 = v3
                            if (true || false) {
                                switch "hello" {
                                case "foo bar":
                                    switch v5 {
                                    case 1:
                                        var v77 uint16 = 255
                                        _ = v77
                                    case 4294967295:
                                        fmt.Println(v3, v76)
                                    case 766:
                                        if len(v1) > 0 {
                                            v1[0] = "\n\t"
                                        }
                                    default:
                                        v78 := GenS6{F0: 42, F1: 18446744073709551615}
                                        _ = v78
                                    }
                                case "\n\t":
                                    v5 = v5
                                }
                                if (false || true) {
                                    switch -223 {
                                    case 127:
                                        var v79 uint8 = 208
                                        _ = v79
                                    case -9223372036854775808:
                                        if len(v6) > 0 {
                                            v6[0] = ""
                                        }
                                    case 642:
                                        _ = len(v1)
                                    default:
                                        v3 = (v3 / ((1e10 * v76) - v3))
                                    }
                                    var v80 uint32 = 363
                                    _ = v80
                                }
                                v81 := GenS7{F0: (127 >= -121)}
                                _ = v81
                            }
                            _ = v76
                        }
                        var v82 int = (v5 << ((406) & 63))
                        v83 := GenS4{F0: 4294967295}
                        _ = v82
                        _ = v83
                    }
                }
            }
            v84 := GenS8{F0: -578, F1: 2}
            if ((355 <= -9223372036854775808) || false) {
                for fi85 := 0; fi85 < 19; fi85++ {
                    v86 := GenS4{F0: 18446744073709551615}
                    v4 = v4
                    v1 = append(v1, "hello")
                    _ = fi85
                    _ = v86
                }
                for fi87 := 0; fi87 < 2; fi87++ {
                    if len(v6) > 0 {
                        v6[0] = "\n\t"
                    }
                    if true {
                        if v2 {
                            var v88 uint16 = 927
                            var v89 bool = ((v5 >> ((4294967295) & 63)) < (v5 / ((-435) | 1)))
                            var v90 uint64 = 1
                            _ = v88
                            _ = v89
                            _ = v90
                        }
                        if (740 <= 42) {
                            var v91 int8 = 1
                            var v92 uint32 = 524
                            if (!v2) {
                                v2 = v2
                                fmt.Println(v2)
                                v91 = 42
                            } else {
                                var v93 bool = false
                                var v94 string = ("\n\t" + "foo bar")
                                _ = v93
                                _ = v94
                            }
                            _ = v91
                            _ = v92
                        }
                        v4 += float64((v3 * ((v3 + v3) * v3)))
                    } else {
                        switch "hello" {
                        case "":
                            v5 = v5
                        case "a":
                            if ((v4 / v4) < (v4 + float64((v3 + 0.0)))) {
                                v1 = append(v1, ("a" + ((("\n\t" + "\n\t") + ("foo bar" + "foo bar")) + (("a" + "\n\t") + ("\n\t" + (("foo bar" + "a") + "a"))))))
                                var v95 float64 = 0.0
                                v2 = (v2 && v2)
                                _ = v95
                            }
                        default:
                            for fi96 := 0; fi96 < 19; fi96++ {
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                                if (v5 < fi87) {
                                    var v97 int64 = -9223372036854775808
                                    v1 = append(v1, "\n\t")
                                    v5 -= -1
                                    _ = v97
                                }
                                _ = fi96
                            }
                        }
                        v2 = false
                    }
                    var v98 int64 = 0
                    _ = fi87
                    _ = v98
                }
                v5 = (v5 | -2147483648)
            } else {
                v5 -= (^(v5))
                v5 -= v5
                if (32767 <= 374) {
                    for fi99 := 0; fi99 < 14; fi99++ {
                        for fi100 := 0; fi100 < 18; fi100++ {
                            if ((166 == 42) && (!v2)) {
                                switch ("" + "hello") {
                                case "a":
                                    switch "" {
                                    case "a":
                                        fmt.Println(v3)
                                    default:
                                        switch fi100 {
                                        case -9223372036854775808:
                                            if (18446744073709551615 < 2) {
                                                var v101 uint64 = 2
                                                _ = len(v1)
                                                v3 += (3.14159 / ((1e10 * v3) * 3.14159))
                                                _ = v101
                                            } else {
                                                if true {
                                                    var v102 uint32 = 65535
                                                    _ = len(v6)
                                                    _ = v102
                                                }
                                            }
                                        default:
                                            var v103 uint8 = 127
                                            _ = v103
                                        }
                                    }
                                case "foo bar":
                                    v2 = v2
                                case "\n\t":
                                    fmt.Println(v2)
                                default:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    fmt.Println(fi99, v3)
                                    switch ("a" + ("a" + ("" + "a"))) {
                                    case "\n\t":
                                        v3 -= v3
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                }
                            }
                            _ = fi100
                        }
                        var v104 int32 = -2147483648
                        for fi105 := 0; fi105 < 11; fi105++ {
                            if len(v1) > 0 {
                                v1[0] = ("a" + "foo bar")
                            }
                            _ = fi105
                        }
                        _ = fi99
                        _ = v104
                    }
                    fmt.Println(v4)
                }
            }
            _ = v84
        }
    default:
        v6 = append(v6, "a")
    }
    if true {
        for fi106 := 0; fi106 < 7; fi106++ {
            switch (v5 << ((236) & 63)) {
            case 375:
                switch -2147483648 {
                case -2147483648:
                    if v2 {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if false {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if (!(!(169 > 893))) {
                                if true {
                                    if len(v6) > 0 {
                                        v6[0] = ("\n\t" + ("\n\t" + "hello"))
                                    }
                                    fmt.Println(fi106)
                                }
                            } else {
                                if v2 {
                                    var v107 uint64 = 2
                                    var v108 int16 = 886
                                    _ = v107
                                    _ = v108
                                }
                                v1 = append(v1, "\n\t")
                                var v109 int16 = -32768
                                _ = v109
                            }
                            var v110 int32 = 0
                            _ = v110
                        }
                    }
                }
            case -423:
                var v111 int16 = -251
                _ = v111
            case -1:
                if v2 {
                    v2 = true
                } else {
                    v2 = (-893 <= 2147483647)
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                }
            default:
                v112 := GenS9{F0: 1, F1: 1, F2: 127}
                _ = v112
            }
            _ = fi106
        }
        if true {
            _ = len(v1)
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch (v5 - -527) {
    case 128:
        fmt.Println(v5, v2, v3)
    case -9223372036854775808:
        v5 -= -448
    default:
        for fi113 := 0; fi113 < 2; fi113++ {
            for fi114 := 0; fi114 < 5; fi114++ {
                if len(v6) > 0 {
                    v6[0] = "foo bar"
                }
                var v115 int8 = -14
                _ = fi114
                _ = v115
            }
            _ = fi113
        }
    }
    if true {
        if v2 {
            v4 -= v4
        } else {
            if (18446744073709551615 == 604) {
                var v116 string = "hello"
                v2 = true
                var v117 uint = 42
                _ = v116
                _ = v117
            }
            if true {
                v1 = append(v1, "foo bar")
            } else {
                v3 = v3
                for fi118 := 0; fi118 < 19; fi118++ {
                    v1 = append(v1, "hello")
                    fmt.Println(v5, v5)
                    switch "\n\t" {
                    case "hello":
                        v6 = append(v6, "\n\t")
                    case "\n\t":
                        v2 = (v4 != (1.5 / (v4 * v4)))
                    case "foo bar":
                        v1 = append(v1, "a")
                    default:
                        v119 := GenS5{F0: 288, F1: 272, F2: 9223372036854775807}
                        _ = v119
                    }
                    _ = fi118
                }
                v2 = v2
            }
            v1 = append(v1, "")
        }
        v1 = append(v1, ((("foo bar" + "foo bar") + "") + ("" + "a")))
    } else {
        v3 = v3
    }
    var v120 int8 = 42
    if false {
        fmt.Println(v5)
    } else {
        v2 = v2
        switch (v5 | v5) {
        case 2147483647:
            v4 -= -2.25
        case 198:
            switch (("a" + "foo bar") + (("\n\t" + "foo bar") + "hello")) {
            case "hello":
                for fi121 := 0; fi121 < 19; fi121++ {
                    var v122 float32 = (0.0 - v3)
                    _ = fi121
                    _ = v122
                }
            default:
                var v123 int32 = 1
                _ = v123
            }
        default:
            var v124 int32 = 223
            _ = v124
        }
    }
    if v2 {
        for fi125 := 0; fi125 < 10; fi125++ {
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                v120 = (v120 / ((v120) | 1))
            case "hello":
                for fi126 := 0; fi126 < 2; fi126++ {
                    v120 += 1
                    if v2 {
                        v4 += (v4 / (float64(float32(3.14159)) / 1.5))
                    }
                    _ = fi126
                }
            }
            _ = fi125
        }
        v127 := GenS10{F0: v120, F1: 65535}
        if false {
            var v128 bool = (0 == 0)
            fmt.Println(v2, v3, v120)
            _ = v128
        } else {
            v129 := GenS7{F0: false}
            _ = v129
        }
        _ = v127
    }
    v120 += (v120 ^ (v120 * (v120 + 98)))
    if len(v6) > 0 {
        v6[0] = "a"
    }
    v6 = append(v6, "")
    var v130 string = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v120:", v120)
    fmt.Println("v130:", v130)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v120
    _ = v130
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
