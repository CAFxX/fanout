package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 uint8; F2 uint32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 bool; F2 int }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int; F1 int8; F2 int }
type GenS10 struct { F0 uint32; F1 int; F2 uint }

func main() {
    v1 := []int{713, 4294967295}
    var v2 bool = true
    v3 := make(map[int]bool)
    v2 = v2
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 string = "hello"
        _ = fi4
        _ = v5
    }
    v2 = false
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v6 int16 = 42
    for fi7 := 0; fi7 < 4; fi7++ {
        v2 = true
        var v8 int32 = -22
        _ = fi7
        _ = v8
    }
    if ("hello" > "") {
        var v9 uint32 = 65535
        fmt.Println(v2, v9)
        _ = v9
    } else {
        if (246 > 201) {
            fmt.Println(v6, v6)
            v1 = append(v1, 2147483647)
            v10 := GenS1{F0: 653, F1: 0}
            _ = v10
        } else {
            if (!true) {
                fmt.Println(v6)
                v2 = v2
                v2 = (2 == 979)
            }
            v11 := GenS1{F0: 4294967295, F1: 4294967295}
            _ = v11
        }
        if v2 {
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 int = (-(fi12))
                if ((false && (!(!false))) && v2) {
                    var v14 int = (^(v13))
                    switch v14 {
                    case -9223372036854775808:
                        fmt.Println(v13)
                    default:
                        if v2 {
                            for fi15 := 0; fi15 < 4; fi15++ {
                                var v16 int16 = -32768
                                v2 = true
                                var v17 int32 = 2147483647
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 15; fi18++ {
                                if v2 {
                                    var v19 int32 = 1
                                    v19 = 55
                                    v20 := GenS2{F0: 42}
                                    _ = v19
                                    _ = v20
                                }
                                switch "" {
                                case "\n\t":
                                    var v21 int32 = -476
                                    _ = v21
                                default:
                                    var v22 int8 = 100
                                    _ = v22
                                }
                                switch 9223372036854775807 {
                                case 9223372036854775807:
                                    var v23 bool = false
                                    _ = v23
                                case 369:
                                    v6 += v6
                                }
                                _ = fi18
                            }
                            v24 := GenS2{F0: 247}
                            _ = v24
                        } else {
                            v25 := GenS1{F0: 305, F1: 1}
                            _ = v25
                        }
                    }
                    for fi26 := 0; fi26 < 1; fi26++ {
                        v14 = fi12
                        v27 := GenS3{F0: -32768}
                        if (966 == 1) {
                            if v2 {
                                v13 = (fi12 & -287)
                                var v28 float32 = 3.14159
                                if true {
                                    switch "" {
                                    case "hello":
                                        v29 := GenS4{F0: (("foo bar" + "hello") + "foo bar"), F1: 149, F2: 2147483647}
                                        _ = v29
                                    case "\n\t":
                                        delete(v3, fi12)
                                        if v2 {
                                            v30 := v3[-1]
                                            v1 = append(v1, -9223372036854775808)
                                            _ = v30
                                        } else {
                                            if (!(!true)) {
                                                v2 = v2
                                                var v31 string = "hello"
                                                _ = v31
                                            }
                                            if v2 {
                                                fmt.Println(v28, fi12, v28)
                                            } else {
                                                var v32 uint16 = 255
                                                if len(v1) > 0 {
                                                    v1[0] = fi26
                                                }
                                                delete(v3, fi12)
                                                _ = v32
                                            }
                                            var v33 int = -572
                                            _ = v33
                                        }
                                    case "foo bar":
                                        if (!v2) {
                                            if v2 {
                                                v14 = (v14 / ((fi12) | 1))
                                                v6 = v6
                                            }
                                            fmt.Println(v6, fi26, v13)
                                            fmt.Println(v2, v28, fi12)
                                            var v34 uint32 = 0
                                            _ = v34
                                        } else {
                                            v3[v14] = v2
                                        }
                                    }
                                    if true {
                                        v35 := GenS2{F0: 1}
                                        _ = v35
                                    }
                                    var v36 int64 = -9223372036854775808
                                    _ = v36
                                } else {
                                    fmt.Println(v2)
                                }
                                v3[v14] = (v2 || (v2 || (-1 > 42)))
                                fmt.Println(v28, v2)
                                v14 = (fi26 & int(v6))
                                var v37 float32 = 1e10
                                if v2 {
                                    v37 = v37
                                    if (v2 || ((v2 || true) || (22 <= 42))) {
                                        fmt.Println(v28, v14)
                                        var v38 int = v14
                                        _ = v38
                                    } else {
                                        if true {
                                            v13 = (^(fi26))
                                            v39 := GenS3{F0: (v6 % ((v6) | 1))}
                                            _ = v39
                                        }
                                    }
                                }
                                if v2 {
                                    v40 := GenS3{F0: (v6 >> ((987) & 63))}
                                    var v41 bool = v2
                                    v42 := v3[v13]
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                }
                                v37 += 0.0
                                _ = v28
                                _ = v37
                            }
                            fmt.Println(fi12, v6, fi26)
                        } else {
                            var v43 uint32 = 65535
                            var v44 int = (fi26 + 4294967295)
                            v44 = (^(v13))
                            switch 1 {
                            case 4294967295:
                                var v45 uint64 = uint64(v43)
                                _ = v45
                            case 127:
                                if v2 {
                                    fmt.Println(v14, fi26, v13)
                                    v1 = append(v1, -2147483648)
                                    var v46 uint64 = 18446744073709551615
                                    _ = v46
                                } else {
                                    v13 += (fi26 & 127)
                                    v43 -= 2147483648
                                    if v2 {
                                        v47 := GenS5{F0: v43}
                                        v2 = v2
                                        v48 := GenS1{F0: (v43 ^ (^(v43))), F1: 2}
                                        _ = v47
                                        _ = v48
                                    }
                                }
                            }
                            v14 -= int(v43)
                            _ = v43
                            _ = v44
                        }
                        _ = fi26
                        _ = v27
                    }
                    _ = v14
                } else {
                    v2 = v2
                    v6 += (v6 << ((903) & 63))
                }
                _ = fi12
                _ = v13
            }
            var v49 int = 4294967295
            if true {
                var v50 int32 = -107
                var v51 uint64 = 18446744073709551615
                _ = v50
                _ = v51
            }
            _ = v49
        } else {
            _ = len(v1)
        }
        for fi52 := 0; fi52 < 1; fi52++ {
            var v53 int = (fi52 + 253)
            _ = fi52
            _ = v53
        }
    }
    v6 += (v6 * v6)
    v2 = (-802 >= 42)
    v54 := GenS6{F0: (!(3.14159 > 3.14159)), F1: v2, F2: 42}
    v2 = false
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!(false || (!(3.14159 < float64((3.14159 * float32(-2.25))))))) {
        var v55 int32 = 2147483647
        _ = v55
    } else {
        switch (("foo bar" + "") + "a") {
        case "":
            for fi56 := 0; fi56 < 17; fi56++ {
                switch "hello" {
                case "a":
                    fmt.Println(fi56, v6, v6)
                case "":
                    if false {
                        for fi57 := 0; fi57 < 13; fi57++ {
                            var v58 int16 = (v6 / (((v6 * 42)) | 1))
                            switch 32768 {
                            case 506:
                                switch 1 {
                                case 127:
                                    _ = len(v3)
                                default:
                                    v2 = false
                                }
                            case 198:
                                switch ((("\n\t" + "a") + ((("" + "a") + "") + "a")) + "\n\t") {
                                case "a":
                                    var v59 string = "a"
                                    _ = v59
                                case "foo bar":
                                    if (946 == 0) {
                                        fmt.Println(fi56, v2)
                                        v2 = v2
                                        if v2 {
                                            _ = len(v1)
                                        } else {
                                            fmt.Println(v2)
                                            switch (((("foo bar" + "foo bar") + "a") + "foo bar") + "a") {
                                            case "foo bar":
                                                v6 = v6
                                            case "hello":
                                                v2 = false
                                            case "a":
                                                v60 := v3[(fi57 * fi57)]
                                                _ = v60
                                            }
                                            var v61 int = fi56
                                            _ = v61
                                        }
                                    }
                                }
                            }
                            _ = fi57
                            _ = v58
                        }
                        fmt.Println(v2, v2)
                        v6 += v6
                    } else {
                        _ = len(v1)
                        switch fi56 {
                        case 4294967295:
                            if v2 {
                                switch 95 {
                                case -388:
                                    v3[fi56] = v2
                                case -9223372036854775808:
                                    v62 := GenS3{F0: (^(v6))}
                                    _ = v62
                                case 9223372036854775807:
                                    if (((1.5 / (3.14159 - 0.1)) != 0.1) || v2) {
                                        if (0 < 4294967295) {
                                            if ((-(fi56)) > fi56) {
                                                v2 = true
                                                var v63 uint8 = 59
                                                _ = v63
                                            } else {
                                                switch (fi56 << ((18446744073709551615) & 63)) {
                                                case 42:
                                                    _ = len(v3)
                                                default:
                                                    var v64 int32 = -806
                                                    _ = v64
                                                }
                                                switch fi56 {
                                                case 32767:
                                                    v65 := GenS7{F0: -43}
                                                    _ = v65
                                                case 1:
                                                    fmt.Println(v2, v2)
                                                case 449:
                                                    fmt.Println(v6, v6, fi56)
                                                }
                                            }
                                            var v66 int16 = (v6 >> ((1) & 63))
                                            _ = v66
                                        } else {
                                            var v67 float64 = ((0.0 + (1.5 + ((3.14159 / 0.1) / 1.0))) - ((1.5 * 1.5) / 3.14159))
                                            _ = v67
                                        }
                                        switch (fi56 / ((128) | 1)) {
                                        case 32768:
                                            switch "a" {
                                            case "":
                                                var v68 float32 = (0.1 / (((1.5 * 0.0) + 0.1) / 1.5))
                                                _ = v68
                                            default:
                                                v2 = v2
                                            }
                                        case 0:
                                            fmt.Println(v2, fi56, v2)
                                        }
                                    } else {
                                        var v69 int64 = 42
                                        for fi70 := 0; fi70 < 9; fi70++ {
                                            switch "foo bar" {
                                            case "hello":
                                                switch ("foo bar" + ("foo bar" + "foo bar")) {
                                                case "foo bar":
                                                    v2 = v2
                                                }
                                            case "\n\t":
                                                v69 += (v69 | v69)
                                            }
                                            if (("foo bar" + "") >= "hello") {
                                                var v71 uint8 = 74
                                                v71 = (v71 ^ (^(v71)))
                                                _ = v71
                                            } else {
                                                if ((v6 * (v6 >> ((65535) & 63))) == (-(v6))) {
                                                    v72 := GenS8{F0: 107}
                                                    _ = v72
                                                } else {
                                                    switch "a" {
                                                    case "\n\t":
                                                        v73 := GenS2{F0: 715}
                                                        var v74 uint8 = 127
                                                        _ = v73
                                                        _ = v74
                                                    default:
                                                        var v75 int64 = (-(v69))
                                                        _ = v75
                                                    }
                                                    var v76 uint16 = 42
                                                    if ("foo bar" == ("foo bar" + "hello")) {
                                                        v76 -= (v76 & (v76 << ((830) & 63)))
                                                        v6 -= (v6 ^ (v6 ^ 1))
                                                        if (!(!v2)) {
                                                            _ = len(v3)
                                                            switch "\n\t" {
                                                            case "a":
                                                                switch 0 {
                                                                case -876:
                                                                    fmt.Println(v2)
                                                                case 4294967295:
                                                                    var v77 int = (fi56 | -2147483648)
                                                                    _ = v77
                                                                case 32768:
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        v69 = 0
                                                                    case "\n\t":
                                                                        switch (fi70 ^ (fi56 ^ 42)) {
                                                                        case 42:
                                                                            v69 += 42
                                                                        case 893:
                                                                            v78 := GenS4{F0: "foo bar", F1: uint8(v76), F2: 86}
                                                                            _ = v78
                                                                        case 4294967295:
                                                                            if false {
                                                                                v79 := GenS5{F0: 2}
                                                                                if (!(float32(float64(3.14159)) > -2.25)) {
                                                                                    if (1.5 == (1e10 * 0.0)) {
                                                                                        fmt.Println(v76, v69)
                                                                                        v3[(-(fi70))] = (0.0 != 1e10)
                                                                                        var v80 uint64 = 42
                                                                                        _ = v80
                                                                                    }
                                                                                    switch "" {
                                                                                    case "a":
                                                                                        v81 := GenS8{F0: 1}
                                                                                        _ = v81
                                                                                    case "\n\t":
                                                                                        var v82 int = fi70
                                                                                        _ = v82
                                                                                    default:
                                                                                        v83 := GenS5{F0: 4294967295}
                                                                                        _ = v83
                                                                                    }
                                                                                }
                                                                                var v84 int = (fi56 | (-(fi56)))
                                                                                _ = v79
                                                                                _ = v84
                                                                            } else {
                                                                                switch fi56 {
                                                                                case 942:
                                                                                    v76 += (v76 + (v76 >> ((117) & 63)))
                                                                                    var v85 int = fi56
                                                                                    var v86 uint8 = 42
                                                                                    if v2 {
                                                                                        var v87 int8 = 42
                                                                                        _ = v87
                                                                                    } else {
                                                                                        fmt.Println(fi56, v2)
                                                                                        if (fi56 > fi56) {
                                                                                            if (!(!v2)) {
                                                                                                if ((((((1.5 / (0.0 + (1.5 - 1e10))) * ((1e10 - -2.25) / 1.0)) * ((((-2.25 / 1e10) - float64(-2.25)) * ((0.1 / 0.1) / 1.5)) + 0.1)) + 0.0) + 1e10) >= 1e10) {
                                                                                                    fmt.Println(fi70, v69)
                                                                                                    v86 = 126
                                                                                                } else {
                                                                                                    _ = len(v3)
                                                                                                    v6 = v6
                                                                                                    var v88 int32 = -2147483648
                                                                                                    _ = v88
                                                                                                }
                                                                                                v89 := v3[fi56]
                                                                                                _ = v89
                                                                                            }
                                                                                            var v90 int = fi70
                                                                                            _ = v90
                                                                                        } else {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = 32768
                                                                                            }
                                                                                        }
                                                                                        var v91 float32 = 1.5
                                                                                        _ = v91
                                                                                    }
                                                                                    _ = v85
                                                                                    _ = v86
                                                                                }
                                                                                v6 -= (v6 >> ((2) & 63))
                                                                            }
                                                                        default:
                                                                            v2 = v2
                                                                        }
                                                                    default:
                                                                        v92 := GenS6{F0: true, F1: ("a" < "hello"), F2: fi70}
                                                                        _ = v92
                                                                    }
                                                                }
                                                            case "":
                                                                var v93 uint16 = v76
                                                                _ = v93
                                                            case "hello":
                                                                fmt.Println(v2, fi56, fi56)
                                                            }
                                                            v6 = v6
                                                        }
                                                    }
                                                    _ = v76
                                                }
                                                v94 := GenS9{F0: (fi70 * (fi56 >> ((409) & 63))), F1: -124, F2: fi70}
                                                v95 := GenS8{F0: 18446744073709551615}
                                                _ = v94
                                                _ = v95
                                            }
                                            _ = fi70
                                        }
                                        _ = v69
                                    }
                                default:
                                    if v2 {
                                        if v2 {
                                            for fi96 := 0; fi96 < 6; fi96++ {
                                                var v97 int8 = 100
                                                var v98 int16 = v6
                                                _ = fi96
                                                _ = v97
                                                _ = v98
                                            }
                                            for fi99 := 0; fi99 < 5; fi99++ {
                                                if (18446744073709551615 >= 42) {
                                                    v6 += 107
                                                    fmt.Println(v6, fi56, v6)
                                                    if false {
                                                        var v100 uint16 = 222
                                                        _ = v100
                                                    }
                                                }
                                                _ = fi99
                                            }
                                        } else {
                                            fmt.Println(v2, v6, v6)
                                            for fi101 := 0; fi101 < 13; fi101++ {
                                                var v102 int64 = -615
                                                _ = fi101
                                                _ = v102
                                            }
                                        }
                                        var v103 uint = 65535
                                        _ = v103
                                    }
                                }
                                delete(v3, 0)
                                for fi104 := 0; fi104 < 8; fi104++ {
                                    fmt.Println(v2, fi104)
                                    _ = fi104
                                }
                            }
                        case 1:
                            var v105 uint32 = 42
                            _ = v105
                        case -711:
                            v1 = append(v1, fi56)
                        }
                    }
                default:
                    for fi106 := 0; fi106 < 5; fi106++ {
                        switch "a" {
                        case "hello":
                            switch fi106 {
                            case -63:
                                var v107 uint16 = 65535
                                _ = v107
                            }
                        default:
                            fmt.Println(fi106, fi56, v6)
                        }
                        if len(v1) > 0 {
                            v1[0] = (fi56 & fi56)
                        }
                        _ = fi106
                    }
                }
                _ = fi56
            }
        case "a":
            v2 = v2
        case "foo bar":
            switch (("\n\t" + "") + ("foo bar" + "hello")) {
            case "a":
                v6 = v6
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = 32768
                }
            }
        default:
            var v108 float64 = 3.14159
            _ = v108
        }
        for fi109 := 0; fi109 < 19; fi109++ {
            var v110 string = ("a" + ("a" + "hello"))
            if (9223372036854775807 < -90) {
                if true {
                    var v111 int32 = -702
                    _ = v111
                }
                if (v2 && false) {
                    v110 = v110
                    var v112 uint32 = 42
                    v112 = (^(v112))
                    _ = v112
                } else {
                    v2 = (!((-701 == 1) && true))
                }
                switch v110 {
                case "a":
                    v113 := GenS8{F0: 808}
                    _ = v113
                case "":
                    v114 := GenS10{F0: 204, F1: fi109, F2: 4294967295}
                    _ = v114
                }
            }
            v2 = (!v2)
            _ = fi109
            _ = v110
        }
        fmt.Println(v2, v6)
    }
    v115 := GenS5{F0: 42}
    _ = len(v3)
    for fi116 := 0; fi116 < 18; fi116++ {
        _ = len(v3)
        _ = fi116
    }
    var v117 string = ""
    v118 := GenS7{F0: 29}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v115
    _ = v117
    _ = v118
    _ = v2
    _ = v3
    _ = v54
    _ = v6
}
