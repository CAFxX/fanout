package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 uint32; F2 int16 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 int; F1 string; F2 int8 }
type GenS9 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int64; F1 uint16 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint32; F1 bool }
type GenS15 struct { F0 bool; F1 string; F2 uint16 }
type GenS16 struct { F0 string; F1 int64; F2 string }
func genFn0(p0 uint16, p1 uint, p2 uint32) bool {
    v4 := GenS1{F0: 128, F1: 1}
    v5 := GenS1{F0: 192, F1: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (!(!true))
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    var v6 int8 = 100
    for fi7 := 0; fi7 < 11; fi7++ {
        if (float64(3.14159) != (0.1 - 0.0)) {
            if (54 == 0) {
                v8 := v3[(fi7 >> ((0) & 63))]
                v6 = v6
                _ = v8
            } else {
                for fi9 := 0; fi9 < 6; fi9++ {
                    var v10 bool = true
                    _ = fi9
                    _ = v10
                }
            }
            if (!(65535 > 803)) {
                if true {
                    v6 += 100
                    v11 := []int{(fi7 + (fi7 << ((1) & 63))), -925, -1, 127}
                    _ = v11
                }
                for fi12 := 0; fi12 < 3; fi12++ {
                    switch 569 {
                    case 42:
                        if true {
                            v13 := GenS1{F0: 110, F1: 42}
                            switch "" {
                            case "hello":
                                if false {
                                    var v14 uint16 = 255
                                    fmt.Println(v6)
                                    v15 := GenS1{F0: uint8(v14), F1: v6}
                                    _ = v14
                                    _ = v15
                                }
                            case "foo bar":
                                if true {
                                    switch 9223372036854775807 {
                                    case 4294967295:
                                        switch fi12 {
                                        case 32768:
                                            switch "foo bar" {
                                            case "hello":
                                                if true {
                                                    switch (("" + "foo bar") + "foo bar") {
                                                    case "\n\t":
                                                        v16 := []int{(fi12 / (((fi12 * 968)) | 1)), 32768}
                                                        _ = v16
                                                    case "":
                                                        v1 += v1
                                                    }
                                                }
                                            case "":
                                                if true {
                                                    if (1 > 1) {
                                                        var v17 float64 = (0.0 - (1.5 * 3.14159))
                                                        _ = v17
                                                    }
                                                    switch (fi12 >> ((255) & 63)) {
                                                    case 127:
                                                        v18 := []int{0, (^(fi12))}
                                                        _ = v18
                                                    case 9223372036854775807:
                                                        v1 -= -1
                                                    }
                                                    v1 += 42
                                                }
                                            default:
                                                v6 += v6
                                            }
                                        case 128:
                                            if true {
                                                var v19 int16 = (v1 % ((v1) | 1))
                                                _ = v19
                                            } else {
                                                var v20 uint16 = 2
                                                var v21 int = fi7
                                                if (v1 == v1) {
                                                    v22 := GenS2{F0: 128, F1: (!(!false))}
                                                    _ = v22
                                                } else {
                                                    fmt.Println(fi12, fi7, v6)
                                                }
                                                _ = v20
                                                _ = v21
                                            }
                                        case 127:
                                            var v23 int8 = (^(v6))
                                            if true {
                                                if (255 == 42) {
                                                    fmt.Println(v1, v23)
                                                    switch 2147483647 {
                                                    case 497:
                                                        if false {
                                                            switch (fi7 + fi12) {
                                                            case 32768:
                                                                delete(v3, fi7)
                                                            case -592:
                                                                v24 := GenS2{F0: (fi12 * (fi7 << ((65535) & 63))), F1: true}
                                                                _ = v24
                                                            case 1:
                                                                fmt.Println(fi7, fi7, v1)
                                                            }
                                                            v25 := GenS1{F0: 2, F1: v23}
                                                            v26 := []int{-1, (-(fi12))}
                                                            v27 := GenS3{F0: 128, F1: fi12}
                                                            var v28 uint32 = 477
                                                            _ = v25
                                                            _ = v26
                                                            _ = v27
                                                            _ = v28
                                                        } else {
                                                            v29 := GenS3{F0: 2, F1: fi12}
                                                            _ = v29
                                                        }
                                                    default:
                                                        v23 += 1
                                                    }
                                                    v6 = v6
                                                }
                                                v30 := GenS4{F0: 1}
                                                v31 := GenS5{F0: -619, F1: 1, F2: (v1 ^ (v1 / ((v1) | 1)))}
                                                _ = v30
                                                _ = v31
                                            } else {
                                                v32 := GenS2{F0: fi7, F1: (v1 > (v1 >> ((65535) & 63)))}
                                                v33 := []int{(fi7 | (fi12 << ((2) & 63))), fi12, (fi7 | 42)}
                                                _ = v32
                                                _ = v33
                                            }
                                            v34 := GenS6{F0: 1, F1: 149}
                                            _ = v23
                                            _ = v34
                                        default:
                                            var v35 uint8 = 127
                                            _ = v35
                                        }
                                    case -2147483648:
                                        v1 += int16(v6)
                                    }
                                    if true {
                                        var v36 uint8 = 66
                                        _ = v36
                                    }
                                    var v37 int = (fi7 / ((fi12) | 1))
                                    _ = v37
                                } else {
                                    v38 := GenS5{F0: 42, F1: 2, F2: (-(v1))}
                                    v1 += v1
                                    var v39 int = fi7
                                    switch "a" {
                                    case "\n\t":
                                        v40 := GenS6{F0: 2, F1: -1}
                                        _ = v40
                                    default:
                                        switch ("a" + ("a" + "")) {
                                        case "a":
                                            var v41 bool = false
                                            _ = v41
                                        case "\n\t":
                                            fmt.Println(fi12)
                                        }
                                    }
                                    _ = v38
                                    _ = v39
                                }
                            case "a":
                                fmt.Println(v6)
                                fmt.Println(fi12, v1)
                            }
                            _ = v13
                        } else {
                            v2[9223372036854775807] = 2147483648
                        }
                    case 884:
                        var v42 uint64 = 2
                        _ = v42
                    }
                    v1 = 177
                    _ = fi12
                }
                v6 += (v6 | 42)
            }
            fmt.Println(v6, fi7)
        }
        if (0.0 != 0.0) {
            fmt.Println(v6)
        }
        for fi43 := 0; fi43 < 15; fi43++ {
            v1 -= v1
            _ = fi43
        }
        _ = fi7
    }
    switch "" {
    case "a":
        v44 := v2[42]
        _ = v44
    }
    for fi45 := 0; fi45 < 4; fi45++ {
        var v46 float64 = (float64((-2.25 / (1.5 - 1e10))) * 0.0)
        _ = fi45
        _ = v46
    }
    var v47 uint = 65535
    v48 := GenS7{F0: 944, F1: int32(v6)}
    switch "\n\t" {
    case "\n\t":
        v6 += (v6 - v6)
    case "hello":
        v49 := GenS5{F0: 531, F1: 710, F2: (v1 & -976)}
        _ = v49
    case "foo bar":
        switch "hello" {
        case "hello":
            var v50 string = "foo bar"
            _ = v50
        case "a":
            v51 := GenS4{F0: 2}
            _ = v51
        default:
            var v52 float64 = (float64((0.0 + ((1e10 * (1e10 - 0.1)) * (0.0 - 1e10)))) / 1e10)
            _ = v52
        }
    default:
        v47 -= 1
    }
    v53 := []int{726, -9223372036854775808}
    fmt.Println(v6)
    for fi54 := 0; fi54 < 13; fi54++ {
        switch (("foo bar" + "foo bar") + "") {
        case "a":
            switch -356 {
            case 32767:
                v1 = 32767
            case -381:
                for fi55 := 0; fi55 < 18; fi55++ {
                    v56 := GenS8{F0: (-(fi54)), F1: "hello", F2: (^(v6))}
                    _ = fi55
                    _ = v56
                }
            }
        case "hello":
            var v57 uint16 = 32767
            _ = v57
        }
        _ = fi54
    }
    fmt.Println(v47, v47)
    if (0 <= 1) {
        v2[-179] = 2147483648
    }
    for fi58 := 0; fi58 < 1; fi58++ {
        switch "\n\t" {
        case "":
            for fi59 := 0; fi59 < 7; fi59++ {
                if (!(!false)) {
                    v47 = v47
                    fmt.Println(v1, fi59, v47)
                    fmt.Println(v47, fi58)
                } else {
                    var v60 uint = (v47 - v47)
                    if false {
                        delete(v2, (fi58 | 702))
                        switch ("hello" + "hello") {
                        case "":
                            var v61 int16 = 0
                            _ = v61
                        case "foo bar":
                            if (128 <= -566) {
                                v62 := GenS5{F0: 124, F1: 360, F2: -1}
                                if (2 <= 42) {
                                    var v63 float32 = 3.14159
                                    _ = v63
                                }
                                _ = v62
                            } else {
                                v64 := GenS1{F0: 1, F1: (-(v6))}
                                switch "foo bar" {
                                case "hello":
                                    switch 128 {
                                    case 4294967295:
                                        var v65 int32 = int32(v1)
                                        _ = v65
                                    default:
                                        var v66 int = -9223372036854775808
                                        _ = v66
                                    }
                                default:
                                    v67 := GenS3{F0: 167, F1: (fi58 << (((v47 & v47)) & 63))}
                                    _ = v67
                                }
                                _ = v64
                            }
                        }
                    }
                    var v68 uint64 = 809
                    _ = v60
                    _ = v68
                }
                _ = fi59
            }
        case "foo bar":
            v1 = (v1 / ((-32768) | 1))
        case "\n\t":
            if true {
                switch "a" {
                case "hello":
                    var v69 float64 = 0.0
                    _ = v69
                case "foo bar":
                    for fi70 := 0; fi70 < 15; fi70++ {
                        delete(v2, (fi70 - fi58))
                        switch (fi58 >> ((v47) & 63)) {
                        case 0:
                            v71 := v3[fi58]
                            _ = v71
                        }
                        v1 = v1
                        _ = fi70
                    }
                case "\n\t":
                    var v72 uint16 = 1
                    _ = v72
                }
                fmt.Println(fi58, v1, v47)
            }
        }
        v73 := GenS4{F0: 65535}
        for fi74 := 0; fi74 < 20; fi74++ {
            var v75 int64 = 9223372036854775807
            if (-132 > 0) {
                if len(v53) >= 2 {
                    v53 = v53[1:len(v53)]
                }
            } else {
                _ = len(v53)
                switch (fi58 % (((fi58 & 42)) | 1)) {
                case 600:
                    fmt.Println(v47, v6)
                case -310:
                    v76 := GenS9{F0: 18446744073709551615, F1: 42, F2: 523}
                    _ = v76
                default:
                    fmt.Println(fi58, v75)
                }
            }
            if (!true) {
                var v77 float32 = ((3.14159 / 1.0) - ((0.0 * (-2.25 - 1.5)) + -2.25))
                v1 += v1
                fmt.Println(v47, fi58)
                _ = v77
            } else {
                switch "" {
                case "\n\t":
                    var v78 int = 127
                    _ = v78
                }
                v79 := GenS10{F0: (-(v1))}
                if false {
                    var v80 bool = false
                    if v80 {
                        var v81 int32 = -2147483648
                        v81 = (v81 / ((759) | 1))
                        v82 := GenS11{F0: 42}
                        _ = v81
                        _ = v82
                    } else {
                        v75 = int64(v6)
                        v1 = v1
                        if true {
                            v1 = 42
                        } else {
                            switch "foo bar" {
                            case "\n\t":
                                var v83 uint16 = 42
                                _ = v83
                            }
                            if ((fi58 << ((v47) & 63)) != fi74) {
                                switch "\n\t" {
                                case "a":
                                    fmt.Println(v80, v1)
                                }
                                fmt.Println(v75, v75, v47)
                                if v80 {
                                    if v80 {
                                        var v84 int8 = -47
                                        var v85 int8 = 0
                                        v86 := GenS1{F0: 128, F1: (v85 ^ v6)}
                                        _ = v84
                                        _ = v85
                                        _ = v86
                                    } else {
                                        if true {
                                            v80 = false
                                            switch "hello" {
                                            case "":
                                                v80 = (!true)
                                                v80 = (!(42 == 403))
                                                var v87 string = ""
                                                _ = v87
                                            }
                                        } else {
                                            v88 := GenS12{F0: v75, F1: 807}
                                            var v89 uint32 = 2
                                            v80 = true
                                            if false {
                                                v80 = v80
                                                v53 = append(v53, int(v6))
                                                if len(v53) > 0 {
                                                    v53[0] = (fi58 | int(v1))
                                                }
                                            }
                                            _ = v88
                                            _ = v89
                                        }
                                        var v90 int8 = (v6 << ((v47) & 63))
                                        if v80 {
                                            if (v47 >= 255) {
                                                v91 := v2[(fi58 * 3)]
                                                if ((v75 >> (((v47 | v47)) & 63)) != v75) {
                                                    v92 := v3[(fi74 + -184)]
                                                    var v93 int64 = v75
                                                    var v94 int = fi74
                                                    _ = v92
                                                    _ = v93
                                                    _ = v94
                                                } else {
                                                    if v80 {
                                                        fmt.Println(v91, v6, v75)
                                                    } else {
                                                        fmt.Println(fi74)
                                                        fmt.Println(v6)
                                                        if v80 {
                                                            var v95 int64 = (-(v75))
                                                            _ = v95
                                                        }
                                                    }
                                                    var v96 int = -2147483648
                                                    _ = v96
                                                }
                                                fmt.Println(v1)
                                                _ = v91
                                            } else {
                                                delete(v2, fi58)
                                            }
                                            var v97 int16 = v1
                                            fmt.Println(v80)
                                            _ = v97
                                        }
                                        _ = v90
                                    }
                                    var v98 bool = (!(2147483647 == 2147483648))
                                    _ = v98
                                } else {
                                    if false {
                                        v99 := GenS13{F0: -128}
                                        v100 := GenS2{F0: (^(fi58)), F1: (1.5 != float32((0.0 - 0.0)))}
                                        v101 := GenS12{F0: v75, F1: 2}
                                        var v102 int = int(v1)
                                        v103 := GenS14{F0: 65535, F1: (874 > 32767)}
                                        _ = v100
                                        _ = v101
                                        _ = v102
                                        _ = v103
                                        _ = v99
                                    } else {
                                        v80 = true
                                        v47 = v47
                                    }
                                    if ((-490 == -2147483648) && false) {
                                        switch ("hello" + (("\n\t" + "foo bar") + (("foo bar" + ("a" + ("" + ("foo bar" + "a")))) + ("" + "foo bar")))) {
                                        case "":
                                            var v104 int = fi58
                                            _ = v104
                                        case "hello":
                                            v6 -= v6
                                        case "foo bar":
                                            if true {
                                                if v80 {
                                                    var v105 uint8 = 0
                                                    v106 := GenS12{F0: 716, F1: uint16(v47)}
                                                    if (!(false && (201 <= (v47 & 65535)))) {
                                                        switch fi74 {
                                                        case 127:
                                                            switch ("hello" + ("a" + "")) {
                                                            case "a":
                                                                v105 -= v105
                                                            case "foo bar":
                                                                var v107 int64 = int64(v47)
                                                                _ = v107
                                                            case "":
                                                                fmt.Println(v47, v75)
                                                                switch ("" + (("\n\t" + "a") + "")) {
                                                                case "":
                                                                    _ = len(v3)
                                                                case "foo bar":
                                                                    v6 = (^(v6))
                                                                case "\n\t":
                                                                    v75 += 484
                                                                }
                                                                v108 := GenS4{F0: 0}
                                                                _ = v108
                                                            default:
                                                                if true {
                                                                    if true {
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            var v109 float64 = 0.1
                                                                            _ = v109
                                                                        case "":
                                                                            var v110 uint64 = 1
                                                                            _ = v110
                                                                        }
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        v3[0] = 0
                                                                    }
                                                                    v47 += v47
                                                                }
                                                            }
                                                        default:
                                                            v111 := GenS15{F0: true, F1: "\n\t", F2: 255}
                                                            _ = v111
                                                        }
                                                        v112 := GenS16{F0: "a", F1: v75, F2: ("\n\t" + "foo bar")}
                                                        _ = v112
                                                    }
                                                    _ = v105
                                                    _ = v106
                                                }
                                                fmt.Println(v75, fi74, v80)
                                                v113 := GenS12{F0: v75, F1: 288}
                                                _ = v113
                                            } else {
                                                if (!(((0.0 * (1.5 - -2.25)) * 0.0) > -2.25)) {
                                                    var v114 int = fi74
                                                    if len(v53) > 0 {
                                                        v53[0] = (fi74 | v114)
                                                    }
                                                    var v115 float64 = 1e10
                                                    _ = v114
                                                    _ = v115
                                                } else {
                                                    switch (fi74 * 127) {
                                                    case 32767:
                                                        fmt.Println(fi74, v1, v80)
                                                    case 367:
                                                        var v116 uint32 = 386
                                                        _ = v116
                                                    }
                                                }
                                            }
                                        default:
                                            if (128 != 2) {
                                                if false {
                                                    fmt.Println(v47)
                                                    if (-33 >= -814) {
                                                        var v117 uint64 = 888
                                                        var v118 uint = 18446744073709551615
                                                        v119 := GenS5{F0: 42, F1: 55, F2: (v1 >> (((v47 / ((42) | 1))) & 63))}
                                                        _ = v117
                                                        _ = v118
                                                        _ = v119
                                                    }
                                                    v75 -= int64(v6)
                                                } else {
                                                    fmt.Println(v1, fi74, fi58)
                                                    v80 = true
                                                }
                                                if (("foo bar" != "\n\t") && (562 <= -615)) {
                                                    v120 := GenS3{F0: 127, F1: fi58}
                                                    _ = len(v53)
                                                    v80 = (2 < 128)
                                                    _ = v120
                                                }
                                                v121 := GenS12{F0: (^(v75)), F1: 360}
                                                v6 -= (v6 / ((v6) | 1))
                                                _ = v121
                                            }
                                        }
                                        v1 = 0
                                    }
                                }
                            }
                        }
                    }
                    _ = v80
                } else {
                    v6 += v6
                    fmt.Println(fi58, v6)
                    switch "\n\t" {
                    case "\n\t":
                        v75 += (v75 / ((0) | 1))
                        v1 = v1
                    }
                }
                _ = v79
            }
            _ = fi74
            _ = v75
        }
        _ = fi58
        _ = v73
    }
    var v122 int32 = 42
    switch (("" + ("a" + (("foo bar" + "foo bar") + ""))) + "hello") {
    case "hello":
        v6 = v6
    case "":
        v47 += (v47 << ((v47) & 63))
    case "a":
        v123 := GenS1{F0: 2, F1: (v6 & (v6 >> ((255) & 63)))}
        _ = v123
    default:
        if (994 >= 2) {
            var v124 float64 = (-2.25 - ((1e10 + ((1e10 + ((float64((1e10 / 1.5)) / 0.1) * -2.25)) - float64(0.1))) / float64((((0.1 * ((0.1 / ((3.14159 / ((1.5 + (1.5 / 1.0)) - (1.5 - -2.25))) / 3.14159)) / (float32((3.14159 * 0.1)) / float32(1.5)))) * (-2.25 / ((0.1 / 1e10) - 0.1))) * 0.1))))
            v47 = v47
            _ = v124
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v122:", v122)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v122
    _ = v2
    _ = v3
    _ = v47
    _ = v48
    _ = v53
    _ = v6
}
