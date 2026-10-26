package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 int16 = 0
    v2 := []int{127}
    var v3 uint8 = 255
    var v4 uint8 = 142
    v5 := []string{"foo bar", ("a" + "a")}
    v6 := make(map[string]string)
    var v7 uint16 = 2
    v8 := GenS1{F0: 467}
    var v9 int = -169
    if true {
        var v10 uint32 = 0
        v3 = 22
        switch (("" + ("\n\t" + "foo bar")) + "a") {
        case "hello":
            for fi11 := 0; fi11 < 16; fi11++ {
                if (true || true) {
                    v10 += v10
                    for fi12 := 0; fi12 < 20; fi12++ {
                        var v13 int32 = 0
                        _ = fi12
                        _ = v13
                    }
                    v1 = 0
                }
                v9 += (-(fi11))
                var v14 float64 = 3.14159
                _ = fi11
                _ = v14
            }
        case "\n\t":
            for fi15 := 0; fi15 < 18; fi15++ {
                v1 = -32768
                if true {
                    if false {
                        if (v7 <= v7) {
                            fmt.Println(v4, v10)
                            fmt.Println(fi15, fi15)
                            var v16 uint32 = (v10 >> ((65535) & 63))
                            _ = v16
                        }
                        fmt.Println(fi15)
                        for fi17 := 0; fi17 < 9; fi17++ {
                            v18 := GenS1{F0: 4294967295}
                            _ = fi17
                            _ = v18
                        }
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                }
                switch "foo bar" {
                case "hello":
                    var v19 int32 = 2147483647
                    _ = v19
                }
                _ = fi15
            }
        default:
            fmt.Println(v10)
        }
        _ = v10
    } else {
        if true {
            fmt.Println(v9, v4)
        } else {
            v2 = append(v2, v9)
            v4 = 112
        }
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 uint32 = 681
            _ = fi20
            _ = v21
        }
        if false {
            if (!(2 > 388)) {
                var v22 bool = (true && (true && false))
                _ = v22
            }
            fmt.Println(v4, v4)
            for fi23 := 0; fi23 < 6; fi23++ {
                if (!(fi23 == (v9 >> ((18446744073709551615) & 63)))) {
                    v7 = (v7 << ((76) & 63))
                    v9 = 139
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                if ((v3 ^ (^(v3))) > (v4 % (((^(v3))) | 1))) {
                    for fi24 := 0; fi24 < 17; fi24++ {
                        var v25 uint64 = 42
                        var v26 uint32 = 132
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    var v27 uint16 = 42
                    if (-128 == 80) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v28 float64 = 0.1
                        _ = v28
                    }
                    _ = v27
                } else {
                    if false {
                        for fi29 := 0; fi29 < 15; fi29++ {
                            v1 = 1
                            delete(v6, ("hello" + "hello"))
                            switch (fi29 - v9) {
                            case 0:
                                v30 := GenS2{F0: false, F1: -9223372036854775808}
                                var v31 uint32 = 2147483648
                                var v32 float32 = 1e10
                                _ = v30
                                _ = v31
                                _ = v32
                            case 1:
                                v6["\n\t"] = "\n\t"
                            }
                            var v33 int16 = (v1 - 1000)
                            _ = fi29
                            _ = v33
                        }
                        if (false || (false && (v1 != (v1 + int16(fi23))))) {
                            var v34 int = 4294967295
                            v35 := GenS1{F0: 1}
                            _ = v34
                            _ = v35
                        } else {
                            var v36 uint64 = 4294967295
                            v1 -= 1000
                            _ = v36
                        }
                    } else {
                        v3 = (v4 + v3)
                    }
                }
                _ = fi23
            }
        }
    }
    v7 += (v7 >> ((42) & 63))
    switch (-(v9)) {
    case 0:
        v1 -= (^(v1))
    default:
        switch 733 {
        case 2147483647:
            fmt.Println(v4, v3)
        case -2147483648:
            var v37 string = ("a" + "foo bar")
            _ = v37
        case 0:
            for fi38 := 0; fi38 < 10; fi38++ {
                var v39 uint64 = 18446744073709551615
                switch "\n\t" {
                case "hello":
                    fmt.Println(v3, v4, v7)
                default:
                    var v40 uint32 = 4294967295
                    _ = v40
                }
                var v41 int64 = 353
                _ = fi38
                _ = v39
                _ = v41
            }
        }
    }
    v3 = v4
    v9 -= (-(v9))
    var v42 uint8 = v4
    var v43 bool = true
    v44 := GenS1{F0: 394}
    for fi45 := 0; fi45 < 2; fi45++ {
        v6[("\n\t" + "hello")] = "a"
        _ = fi45
    }
    v3 = 42
    if len(v5) > 0 {
        v5[0] = ("hello" + "hello")
    }
    for fi46 := 0; fi46 < 1; fi46++ {
        var v47 int64 = 156
        fmt.Println(fi46)
        for fi48 := 0; fi48 < 19; fi48++ {
            switch "a" {
            case "":
                if (!v43) {
                    v49 := GenS2{F0: v43, F1: (v47 << ((255) & 63))}
                    var v50 int16 = (v1 * -32768)
                    _ = v49
                    _ = v50
                }
            case "a":
                v3 -= v4
            case "hello":
                switch "a" {
                case "hello":
                    if ((4294967295 == 42) || false) {
                        if (v43 && (!true)) {
                            fmt.Println(v3, v4)
                            if ((v47 % (((v47 + -1)) | 1)) == v47) {
                                switch "" {
                                case "\n\t":
                                    var v51 uint32 = 2147483647
                                    _ = v51
                                case "foo bar":
                                    fmt.Println(v7, fi46)
                                case "a":
                                    if ((fi48 ^ (fi46 ^ (^(fi48)))) == (-(v9))) {
                                        var v52 int = (-(fi48))
                                        if false {
                                            var v53 uint16 = (^(v7))
                                            var v54 uint32 = 42
                                            if v43 {
                                                if (true && v43) {
                                                    v55 := GenS3{F0: v53, F1: 76}
                                                    v3 = 42
                                                    switch (("hello" + (("hello" + "hello") + ("foo bar" + "foo bar"))) + "foo bar") {
                                                    case "\n\t":
                                                        v42 = v42
                                                    case "hello":
                                                        if v43 {
                                                            var v56 int64 = v47
                                                            v57 := GenS1{F0: 4294967295}
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                    }
                                                    if v43 {
                                                        fmt.Println(fi46, v1, v43)
                                                    }
                                                    v58 := GenS1{F0: 1}
                                                    var v59 int32 = 1
                                                    var v60 uint16 = v7
                                                    _ = v55
                                                    _ = v58
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    v47 += 42
                                                }
                                                if (v43 && false) {
                                                    if (-32768 <= v1) {
                                                        var v61 uint8 = (v42 ^ 127)
                                                        _ = v61
                                                    } else {
                                                        v54 += 334
                                                        switch "" {
                                                        case "a":
                                                            v43 = (uint64(v7) > 4294967295)
                                                        }
                                                        if (42 < 1) {
                                                            v62 := GenS2{F0: false, F1: -9223372036854775808}
                                                            fmt.Println(v4, v4)
                                                            if v43 {
                                                                if v43 {
                                                                    v9 += 127
                                                                    switch v9 {
                                                                    case -63:
                                                                        v63 := GenS3{F0: (v53 | 32767), F1: 42}
                                                                        _ = v63
                                                                    }
                                                                    v7 = v53
                                                                } else {
                                                                    if (((v43 && v43) && v43) && (0 != 0)) {
                                                                        v3 = v42
                                                                    } else {
                                                                        v7 += uint16(v4)
                                                                        v7 = uint16(v42)
                                                                    }
                                                                    switch ("a" + "hello") {
                                                                    case "hello":
                                                                        fmt.Println(v54)
                                                                    case "\n\t":
                                                                        v52 -= 42
                                                                    }
                                                                }
                                                            } else {
                                                                _ = len(v6)
                                                            }
                                                            _ = v62
                                                        }
                                                        if (v54 <= v54) {
                                                            v53 += (v7 << ((18446744073709551615) & 63))
                                                            if v43 {
                                                                v64 := GenS2{F0: true, F1: 836}
                                                                v65 := GenS4{F0: 496}
                                                                v66 := GenS3{F0: 2, F1: 133}
                                                                _ = v64
                                                                _ = v65
                                                                _ = v66
                                                            } else {
                                                                fmt.Println(v43, v54)
                                                                delete(v6, "hello")
                                                                var v67 string = "\n\t"
                                                                var v68 float64 = ((-2.25 - (0.0 - 0.0)) - 3.14159)
                                                                _ = v67
                                                                _ = v68
                                                            }
                                                        } else {
                                                            if len(v5) > 0 {
                                                                v5[0] = "hello"
                                                            }
                                                            v69 := GenS1{F0: 18446744073709551615}
                                                            _ = v69
                                                        }
                                                        if v43 {
                                                            var v70 int16 = v1
                                                            _ = v70
                                                        }
                                                    }
                                                    if ((^(v54)) >= (v54 / (((v54 + 65535)) | 1))) {
                                                        switch (^(v52)) {
                                                        case -265:
                                                            var v71 int8 = 34
                                                            _ = v71
                                                        case 127:
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                v72 := GenS4{F0: 0}
                                                                _ = v72
                                                            case "a":
                                                                v73 := GenS1{F0: 472}
                                                                _ = v73
                                                            default:
                                                                if v43 {
                                                                    if ((v43 || true) && ((float64((3.14159 + -2.25)) * (3.14159 - 0.1)) <= ((((1e10 / float64(float32(3.14159))) / 1.0) / (3.14159 / 1.0)) + (3.14159 * 3.14159)))) {
                                                                        if (float32((float64((-2.25 + -2.25)) + (-2.25 + 1e10))) > -2.25) {
                                                                            var v74 int8 = 1
                                                                            v42 = (v3 - 128)
                                                                            _ = v74
                                                                        }
                                                                        switch 42 {
                                                                        case -1:
                                                                            _ = len(v5)
                                                                        case 724:
                                                                            _ = len(v6)
                                                                        case 0:
                                                                            v42 = 0
                                                                        }
                                                                        var v75 string = "a"
                                                                        _ = v75
                                                                    } else {
                                                                        v4 = v4
                                                                        switch (v9 % (((v9 + (v52 % ((int(v54)) | 1)))) | 1)) {
                                                                        case 32767:
                                                                            var v76 string = "a"
                                                                            _ = v76
                                                                        case 42:
                                                                            if v43 {
                                                                                var v77 int32 = 316
                                                                                v78 := GenS4{F0: (-(v77))}
                                                                                _ = v77
                                                                                _ = v78
                                                                            }
                                                                        case 855:
                                                                            switch v9 {
                                                                            case 2147483647:
                                                                                fmt.Println(v42, v9)
                                                                            case 32768:
                                                                                v42 = v4
                                                                            default:
                                                                                switch ("\n\t" + (("foo bar" + ("hello" + "hello")) + (("\n\t" + "\n\t") + "\n\t"))) {
                                                                                case "\n\t":
                                                                                    if (!v43) {
                                                                                        var v79 int = fi46
                                                                                        v80 := GenS4{F0: 2147483647}
                                                                                        _ = v79
                                                                                        _ = v80
                                                                                    } else {
                                                                                        v54 += v54
                                                                                        var v81 uint16 = v7
                                                                                        v3 = v3
                                                                                        _ = v81
                                                                                    }
                                                                                default:
                                                                                    v82 := GenS4{F0: -2147483648}
                                                                                    _ = v82
                                                                                }
                                                                            }
                                                                        }
                                                                        v83 := GenS5{F0: 563}
                                                                        _ = v83
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        v84 := GenS5{F0: int32(v4)}
                                                        v85 := GenS4{F0: 2147483647}
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                } else {
                                                    delete(v6, "foo bar")
                                                    _ = len(v6)
                                                }
                                            }
                                            _ = v53
                                            _ = v54
                                        } else {
                                            if (127 > v52) {
                                                var v86 uint32 = 4294967295
                                                var v87 int16 = v1
                                                if v43 {
                                                    v88 := GenS1{F0: 18446744073709551615}
                                                    if v43 {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        var v89 int8 = -48
                                                        if v43 {
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v86 = v86
                                                            case "a":
                                                                switch ("\n\t" + "") {
                                                                case "hello":
                                                                    if v43 {
                                                                        v90 := GenS1{F0: 971}
                                                                        var v91 int16 = (v87 % ((v87) | 1))
                                                                        v5 = append(v5, ("" + "foo bar"))
                                                                        _ = v90
                                                                        _ = v91
                                                                    }
                                                                case "":
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        var v92 int16 = v87
                                                                        _ = v92
                                                                    default:
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            _ = len(v6)
                                                                        default:
                                                                            switch (^(v9)) {
                                                                            case 32768:
                                                                                if true {
                                                                                    v43 = v43
                                                                                    v6[("foo bar" + "a")] = ("foo bar" + "")
                                                                                } else {
                                                                                    v3 = (^(v4))
                                                                                    v42 = (v3 << ((4294967295) & 63))
                                                                                }
                                                                            case 127:
                                                                                var v93 int8 = v89
                                                                                _ = v93
                                                                            default:
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    v7 -= (v7 - (^(v7)))
                                                                default:
                                                                    if (((v89 ^ (v89 ^ (v89 | (v89 & v89)))) == (v89 >> ((18446744073709551615) & 63))) && (574 > 9223372036854775807)) {
                                                                        var v94 float64 = (3.14159 - -2.25)
                                                                        v95 := GenS3{F0: (v7 & v7), F1: uint64(v86)}
                                                                        var v96 uint16 = 0
                                                                        _ = v94
                                                                        _ = v95
                                                                        _ = v96
                                                                    } else {
                                                                        if (!v43) {
                                                                            var v97 float32 = (0.1 / (-2.25 * 3.14159))
                                                                            switch fi48 {
                                                                            case 4294967295:
                                                                                v98 := GenS3{F0: 0, F1: 190}
                                                                                _ = v98
                                                                            case -9223372036854775808:
                                                                                _ = len(v5)
                                                                            case 0:
                                                                                var v99 int8 = (v89 / (((v89 & (v89 % ((v89) | 1)))) | 1))
                                                                                _ = v99
                                                                            default:
                                                                                v97 = 0.0
                                                                            }
                                                                            _ = v97
                                                                        }
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    }
                                                                }
                                                            case "":
                                                                if true {
                                                                    var v100 float64 = 3.14159
                                                                    v101 := GenS5{F0: -2147483648}
                                                                    fmt.Println(v86)
                                                                    if (2147483647 > 42) {
                                                                        fmt.Println(v7, v4, v9)
                                                                        _ = len(v6)
                                                                        var v102 uint64 = 4294967295
                                                                        _ = v102
                                                                    }
                                                                    _ = v100
                                                                    _ = v101
                                                                } else {
                                                                    v52 = int(v4)
                                                                    v89 -= 1
                                                                }
                                                            }
                                                            switch (fi48 / ((19) | 1)) {
                                                            case -9223372036854775808:
                                                                var v103 int = v52
                                                                _ = v103
                                                            case 32768:
                                                                fmt.Println(v42, v86)
                                                            case 32767:
                                                                var v104 int64 = (-(v47))
                                                                _ = v104
                                                            }
                                                        }
                                                        switch fi48 {
                                                        case 42:
                                                            var v105 bool = false
                                                            _ = v105
                                                        case 128:
                                                            v7 = v7
                                                        case -323:
                                                            if ((int8(v4) < v89) && true) {
                                                                v106 := GenS3{F0: (v7 ^ v7), F1: 0}
                                                                v107 := GenS3{F0: v7, F1: uint64(v4)}
                                                                _ = v106
                                                                _ = v107
                                                            }
                                                        default:
                                                            _ = len(v2)
                                                        }
                                                        var v108 bool = true
                                                        _ = v108
                                                        _ = v89
                                                    } else {
                                                        switch (v9 & (v52 ^ fi48)) {
                                                        case -254:
                                                            var v109 int64 = v47
                                                            _ = v109
                                                        default:
                                                            var v110 float32 = (1.5 - 3.14159)
                                                            _ = v110
                                                        }
                                                        v111 := GenS5{F0: 975}
                                                        v7 += v7
                                                        _ = v111
                                                    }
                                                    _ = len(v6)
                                                    _ = v88
                                                } else {
                                                    fmt.Println(v52)
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                _ = v86
                                                _ = v87
                                            } else {
                                                fmt.Println(v52, v42, v52)
                                            }
                                        }
                                        v43 = (((!(!v43)) && ((0 >= -480) && (0.0 != 1e10))) || v43)
                                        _ = v52
                                    }
                                }
                            }
                        } else {
                            switch (v9 >> ((42) & 63)) {
                            case 42:
                                v112 := GenS1{F0: 1}
                                _ = v112
                            default:
                                v3 -= (v3 - (v3 & 255))
                            }
                            v113 := GenS1{F0: 42}
                            v42 -= v3
                            switch ("foo bar" + "a") {
                            case "":
                                if v43 {
                                    v3 -= 127
                                    v9 = (fi46 | (^(fi46)))
                                }
                            case "\n\t":
                                v114 := GenS6{F0: 2}
                                _ = v114
                            default:
                                switch ("\n\t" + "hello") {
                                case "foo bar":
                                    switch "hello" {
                                    case "a":
                                        var v115 uint16 = (v7 >> ((633) & 63))
                                        _ = v115
                                    case "foo bar":
                                        if (3.14159 > ((3.14159 * 1.5) / -2.25)) {
                                            if v43 {
                                                v1 -= (v1 - (v1 << ((65535) & 63)))
                                                if (!v43) {
                                                    fmt.Println(v1, v47)
                                                    v42 -= v3
                                                    fmt.Println(v43, v9)
                                                }
                                                fmt.Println(v3, v3)
                                                v6[("" + "a")] = ""
                                                var v116 int = (fi48 - 32768)
                                                _ = v116
                                            } else {
                                                if (false && v43) {
                                                    var v117 int8 = 30
                                                    _ = v117
                                                }
                                                switch "foo bar" {
                                                case "hello":
                                                    var v118 bool = ((0.0 >= 1e10) && ((v1 & v1) > v1))
                                                    _ = v118
                                                }
                                                v119 := GenS4{F0: -552}
                                                _ = v119
                                            }
                                        }
                                    default:
                                        v4 -= v4
                                    }
                                case "":
                                    v120 := GenS4{F0: -36}
                                    _ = v120
                                case "hello":
                                    switch "foo bar" {
                                    case "":
                                        var v121 int32 = int32(v42)
                                        _ = v121
                                    }
                                }
                            }
                            _ = v113
                        }
                    }
                case "\n\t":
                    v9 += (v9 & (-(fi48)))
                default:
                    v3 += (v42 - (v4 ^ 42))
                    if false {
                        fmt.Println(v42, v42, fi48)
                        v122 := GenS1{F0: 2}
                        _ = v122
                    }
                }
            }
            _ = fi48
        }
        _ = fi46
        _ = v47
    }
    v123 := GenS1{F0: 42}
    fmt.Println(v7, v3, v43)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v123
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
