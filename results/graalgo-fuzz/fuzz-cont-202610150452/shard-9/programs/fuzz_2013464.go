package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int64; F1 int8; F2 int32 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 int16 }
type GenS8 struct { F0 int64; F1 uint32 }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 int; F1 uint64; F2 int8 }
type GenS11 struct { F0 int32; F1 int16 }
type GenS12 struct { F0 string; F1 int8 }
type GenS13 struct { F0 uint32; F1 uint32; F2 string }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int32; F1 string }
type GenS16 struct { F0 int64; F1 int64 }
type GenS17 struct { F0 uint; F1 uint; F2 int64 }
type GenS18 struct { F0 uint64 }
type GenS19 struct { F0 int }
type GenS20 struct { F0 uint16 }
type GenS21 struct { F0 uint; F1 int32 }
type GenS22 struct { F0 int8; F1 int8 }
type GenS23 struct { F0 int16 }
type GenS24 struct { F0 int8 }
type GenS25 struct { F0 int32 }
type GenS26 struct { F0 uint32; F1 bool }
type GenS27 struct { F0 int64; F1 int32 }
type GenS28 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS29 struct { F0 int64; F1 int8 }
type GenS30 struct { F0 uint16; F1 int16; F2 int32 }
type GenS31 struct { F0 uint8; F1 int32 }
type GenS32 struct { F0 int32; F1 uint }
func genFn0(p0 int64, p1 uint16) uint32 {
    p1 = (p1 * (p1 ^ (p1 - (p1 & 32767))))
    var v5 uint64 = 0
    switch "hello" {
    case "":
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 uint64 = (v5 + 0)
            p0 = p0
            _ = fi6
            _ = v7
        }
    }
    switch 2147483647 {
    case 128:
        v8 := GenS1{F0: 128}
        _ = v8
    case 2147483647:
        v9 := GenS1{F0: 2}
        _ = v9
    case -674:
        v5 = 0
    default:
        v10 := []int{}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1(p0 int16, p1 uint8, p2 uint) uint32 {
    var v11 uint32 = 2
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 uint16 = 1
            switch (-(fi12)) {
            case -9223372036854775808:
                fmt.Println(p2, v11)
            case 9223372036854775807:
                v15 := GenS2{F0: 2147483647, F1: -2147483648}
                _ = v15
            case 128:
                var v16 float64 = -2.25
                _ = v16
            }
            p2 = (p2 | (^(p2)))
            fmt.Println(fi13, v11)
            _ = fi13
            _ = v14
        }
        var v17 int8 = 100
        _ = fi12
        _ = v17
    }
    if true {
        var v18 int = int(v11)
        _ = v18
    } else {
        var v19 bool = true
        var v20 uint64 = 1
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v21
    return 65535
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    if (-59 > -32768) {
        var v22 int64 = 5
        var v23 bool = false
        var v24 float64 = -2.25
        _ = v22
        _ = v23
        _ = v24
    }
    var v25 float64 = (1e10 / ((1.5 / 1.0) * (-2.25 + (((0.1 * ((((float64(3.14159) / 1e10) * ((1e10 - (-2.25 - 1e10)) * -2.25)) * float64(3.14159)) - -2.25)) / 1.0) * (3.14159 + (0.0 * (-2.25 * float64(3.14159))))))))
    if ("\n\t" != "hello") {
        if ((2147483648 < 4294967295) || (42 == 2)) {
            for fi26 := 0; fi26 < 3; fi26++ {
                if (((v25 == v25) && (42 != -2147483648)) || (!false)) {
                    var v27 float64 = 1e10
                    v27 = v27
                    switch 859 {
                    case -2147483648:
                        v25 -= v27
                    case 87:
                        switch (^(fi26)) {
                        case 9223372036854775807:
                            var v28 int16 = 1000
                            _ = v28
                        case -9223372036854775808:
                            for fi29 := 0; fi29 < 9; fi29++ {
                                if (true && (-732 < 32767)) {
                                    var v30 bool = true
                                    if (v30 && v30) {
                                        var v31 uint16 = 32767
                                        fmt.Println(v30, fi29, fi26)
                                        v32 := GenS3{F0: 1, F1: -1, F2: 81}
                                        _ = v31
                                        _ = v32
                                    } else {
                                        v33 := GenS3{F0: 36, F1: -42, F2: -1}
                                        _ = v33
                                    }
                                    _ = v30
                                }
                                _ = fi29
                            }
                        case 260:
                            switch ("hello" + "hello") {
                            case "":
                                switch "hello" {
                                case "hello":
                                    var v34 float64 = (v27 + -2.25)
                                    _ = v34
                                case "foo bar":
                                    if (-2.25 != (3.14159 / 1.5)) {
                                        v25 += (((float64((1.5 + ((1e10 / 1.5) + 3.14159))) * (float64(0.0) + float64(((0.0 + 1e10) + (1e10 * 1e10))))) / 1e10) * v25)
                                        for fi35 := 0; fi35 < 6; fi35++ {
                                            fmt.Println(fi35)
                                            var v36 uint16 = 978
                                            if true {
                                                switch (fi26 ^ 9223372036854775807) {
                                                case 42:
                                                    var v37 uint = 307
                                                    _ = v37
                                                case -874:
                                                    var v38 int64 = 842
                                                    _ = v38
                                                case 104:
                                                    v39 := v4[480]
                                                    _ = v39
                                                default:
                                                    switch "a" {
                                                    case "a":
                                                        var v40 string = "a"
                                                        _ = v40
                                                    case "\n\t":
                                                        v25 = v25
                                                    default:
                                                        v41 := GenS4{F0: uint(v36), F1: 1}
                                                        v27 += v27
                                                        switch 42 {
                                                        case 32767:
                                                            v25 = v27
                                                        case -9223372036854775808:
                                                            v4[426] = (false || (false || true))
                                                        case 42:
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                v42 := []int{(-(fi26)), fi35}
                                                                _ = v42
                                                            case "":
                                                                switch -684 {
                                                                case 32768:
                                                                    var v43 int = fi26
                                                                    _ = v43
                                                                case -73:
                                                                    if ((v36 / (((v36 * 65535)) | 1)) < (v36 & v36)) {
                                                                        fmt.Println(fi35, fi26)
                                                                        if false {
                                                                            v44 := GenS5{F0: -1, F1: int64(v36)}
                                                                            if (726 <= 2) {
                                                                                switch fi26 {
                                                                                case 638:
                                                                                    v45 := GenS6{F0: 1}
                                                                                    _ = v45
                                                                                case 128:
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        var v46 int8 = 52
                                                                                        _ = v46
                                                                                    case "":
                                                                                        v47 := v1[2]
                                                                                        _ = v47
                                                                                    }
                                                                                case 32767:
                                                                                    v48 := []int{}
                                                                                    _ = v48
                                                                                }
                                                                            }
                                                                            _ = v44
                                                                        } else {
                                                                            v49 := v3[genFn0(0, v36)]
                                                                            _ = v49
                                                                        }
                                                                    } else {
                                                                        v25 = v25
                                                                        if (("a" + ("hello" + "foo bar")) >= "a") {
                                                                            v36 += (^(v36))
                                                                        } else {
                                                                            v50 := GenS2{F0: 612, F1: -2147483648}
                                                                            var v51 int16 = -32768
                                                                            _ = v50
                                                                            _ = v51
                                                                        }
                                                                    }
                                                                case 32767:
                                                                    v52 := []int{fi35, (fi35 ^ (fi35 + fi26)), (fi26 | 615)}
                                                                    _ = v52
                                                                }
                                                            default:
                                                                var v53 float32 = 0.1
                                                                _ = v53
                                                            }
                                                        default:
                                                            v36 += (v36 % (((^(v36))) | 1))
                                                        }
                                                        fmt.Println(fi26, v36)
                                                        _ = v41
                                                    }
                                                }
                                            }
                                            _ = fi35
                                            _ = v36
                                        }
                                        v54 := []int{(fi26 + fi26), (fi26 % ((fi26) | 1)), (fi26 << ((1) & 63)), 42}
                                        _ = v54
                                    } else {
                                        var v55 uint = 65535
                                        _ = v55
                                    }
                                default:
                                    for fi56 := 0; fi56 < 8; fi56++ {
                                        var v57 bool = (3.14159 <= (1e10 + 1e10))
                                        v58 := v1[745]
                                        _ = fi56
                                        _ = v57
                                        _ = v58
                                    }
                                }
                            case "hello":
                                var v59 bool = ((0 == -1) || true)
                                _ = v59
                            }
                        }
                    case -863:
                        switch fi26 {
                        case 127:
                            for fi60 := 0; fi60 < 3; fi60++ {
                                v61 := GenS7{F0: "\n\t", F1: 482}
                                v27 -= 0.0
                                _ = fi60
                                _ = v61
                            }
                        default:
                            var v62 uint8 = 2
                            _ = v62
                        }
                    default:
                        var v63 uint16 = 0
                        _ = v63
                    }
                    _ = v27
                } else {
                    v25 = (0.1 - ((0.1 + (v25 / 1e10)) / float64((1.5 * 3.14159))))
                    for fi64 := 0; fi64 < 11; fi64++ {
                        v65 := []int{(fi26 / ((fi26) | 1)), fi64, fi26, -1}
                        v25 += (float64(1.5) - float64(-2.25))
                        v66 := GenS7{F0: ("hello" + "hello"), F1: 0}
                        if (false && (false && true)) {
                            if false {
                                v25 += v25
                                switch (fi64 ^ fi26) {
                                case -610:
                                    fmt.Println(v25)
                                case 1:
                                    v65 = append(v65, fi26)
                                default:
                                    v25 -= ((v25 / 3.14159) * (v25 * v25))
                                }
                                switch "" {
                                case "":
                                    switch "foo bar" {
                                    case "a":
                                        switch -399 {
                                        case 544:
                                            fmt.Println(v25, v25)
                                        case 657:
                                            var v67 bool = (10 < 50)
                                            _ = v67
                                        case 567:
                                            switch (fi64 * fi26) {
                                            case 9223372036854775807:
                                                v25 -= v25
                                            case 70:
                                                fmt.Println(fi26, v25)
                                            case -135:
                                                var v68 string = "hello"
                                                _ = v68
                                            default:
                                                _ = len(v2)
                                            }
                                            v65 = append(v65, fi26)
                                            if (fi26 >= 42) {
                                                var v69 int = (fi26 * fi64)
                                                var v70 string = ""
                                                var v71 bool = (false || (!(!(fi26 == fi64))))
                                                v72 := GenS8{F0: 791, F1: genFn1(1, 128, 1)}
                                                if (((("\n\t" + v70) != v70) || (4294967295 < 270)) && v71) {
                                                    switch 9223372036854775807 {
                                                    case 307:
                                                        _ = len(v3)
                                                    case 0:
                                                        var v73 int8 = 70
                                                        _ = v73
                                                    case -9223372036854775808:
                                                        fmt.Println(v70, v71, v69)
                                                    }
                                                    v71 = false
                                                    v74 := GenS7{F0: v70, F1: 1}
                                                    var v75 int16 = -32768
                                                    _ = v74
                                                    _ = v75
                                                } else {
                                                    var v76 uint8 = 42
                                                    var v77 int32 = 2147483647
                                                    var v78 uint64 = 0
                                                    _ = v76
                                                    _ = v77
                                                    _ = v78
                                                }
                                                if (2147483647 > 4294967295) {
                                                    fmt.Println(v69)
                                                    v79 := GenS6{F0: 42}
                                                    v71 = v71
                                                    _ = v79
                                                }
                                                v80 := GenS1{F0: 1}
                                                _ = v69
                                                _ = v70
                                                _ = v71
                                                _ = v72
                                                _ = v80
                                            }
                                        default:
                                            if (!(284 < 0)) {
                                                v25 = (v25 - (v25 + v25))
                                                v81 := GenS2{F0: -314, F1: 0}
                                                _ = v81
                                            } else {
                                                v25 -= 1e10
                                            }
                                        }
                                    case "foo bar":
                                        v82 := GenS9{F0: -1, F1: -1}
                                        _ = v82
                                    case "\n\t":
                                        v25 -= float64(((float32(1.5) * 1e10) * (0.1 - float32(float64((0.1 * (((0.1 / 1e10) - (1e10 - -2.25)) * 3.14159)))))))
                                    }
                                case "\n\t":
                                    v83 := GenS3{F0: 9223372036854775807, F1: -77, F2: -595}
                                    _ = v83
                                }
                            }
                            fmt.Println(v25, fi26, fi64)
                        } else {
                            switch fi26 {
                            case 9223372036854775807:
                                v25 = ((v25 / -2.25) / v25)
                                if (0.0 >= 0.0) {
                                    v25 = 3.14159
                                    var v84 float64 = v25
                                    _ = v84
                                } else {
                                    fmt.Println(fi64)
                                    switch -506 {
                                    case -2147483648:
                                        v85 := GenS10{F0: (fi26 << ((65535) & 63)), F1: 628, F2: -128}
                                        _ = v85
                                    case 11:
                                        switch "hello" {
                                        case "foo bar":
                                            var v86 float64 = float64(float32(v25))
                                            _ = v86
                                        case "hello":
                                            v87 := GenS11{F0: -1, F1: -1}
                                            _ = v87
                                        case "\n\t":
                                            if false {
                                                var v88 uint8 = 0
                                                fmt.Println(fi26, fi64, v25)
                                                _ = v88
                                            } else {
                                                var v89 uint16 = 2
                                                _ = v89
                                            }
                                        }
                                    case -119:
                                        delete(v1, 2)
                                    }
                                }
                                if false {
                                    fmt.Println(fi26)
                                    fmt.Println(fi26)
                                } else {
                                    if ((((false && (179 <= 42)) || ((42 >= 0) && false)) || (false || ((!(2 < 18446744073709551615)) && false))) || true) {
                                        v25 = -2.25
                                        switch "" {
                                        case "":
                                            var v90 int8 = 0
                                            _ = v90
                                        case "foo bar":
                                            v91 := v1[698]
                                            _ = v91
                                        }
                                        _ = len(v4)
                                        v25 -= (((float64(1e10) - 0.1) - -2.25) * v25)
                                    }
                                    var v92 int8 = 127
                                    var v93 int = (fi26 - fi64)
                                    _ = v92
                                    _ = v93
                                }
                            case 0:
                                switch (fi64 << ((42) & 63)) {
                                case 1:
                                    var v94 bool = (72 != -739)
                                    _ = v94
                                case -1:
                                    var v95 int16 = -600
                                    _ = v95
                                case 863:
                                    switch (fi64 << ((2) & 63)) {
                                    case 127:
                                        fmt.Println(fi64, fi64)
                                    }
                                default:
                                    var v96 float32 = 0.1
                                    if false {
                                        if len(v65) > 0 {
                                            v65[0] = 177
                                        }
                                    } else {
                                        v97 := GenS8{F0: 42, F1: 65535}
                                        _ = v97
                                    }
                                    delete(v4, 2147483647)
                                    _ = v96
                                }
                            case -2147483648:
                                v25 -= 1e10
                            }
                            v25 = ((0.0 / v25) / 1.0)
                        }
                        if (!false) {
                            switch (fi64 * (fi64 - (-(fi26)))) {
                            case -80:
                                fmt.Println(fi64)
                            case 42:
                                var v98 float32 = 1e10
                                _ = v98
                            case -568:
                                v25 += (v25 / 0.1)
                            default:
                                var v99 uint64 = 18446744073709551615
                                _ = v99
                            }
                        } else {
                            if (-1 >= 2147483647) {
                                v25 = (3.14159 * 0.1)
                                if (float32(1.5) <= -2.25) {
                                    switch fi26 {
                                    case -9223372036854775808:
                                        var v100 uint = 0
                                        _ = v100
                                    case 4294967295:
                                        v101 := GenS6{F0: 180}
                                        _ = v101
                                    }
                                    fmt.Println(v25, fi26)
                                    v25 -= v25
                                } else {
                                    v25 += ((float64((-2.25 * ((0.1 + 1e10) * float32(-2.25)))) / 1e10) / (v25 / 1e10))
                                    fmt.Println(fi26)
                                    fmt.Println(v25, v25, fi64)
                                }
                                _ = len(v1)
                            } else {
                                v25 += (0.1 / 0.1)
                                v102 := GenS8{F0: -1, F1: 65535}
                                _ = v102
                            }
                            v25 = (-2.25 + -2.25)
                        }
                        if false {
                            _ = len(v65)
                            v103 := GenS3{F0: 654, F1: 0, F2: 155}
                            v25 = 3.14159
                            _ = v103
                        } else {
                            v25 = (v25 / (1.5 / v25))
                            var v104 uint8 = 127
                            v105 := GenS10{F0: (fi26 >> ((255) & 63)), F1: 18446744073709551615, F2: 36}
                            _ = v104
                            _ = v105
                        }
                        switch "\n\t" {
                        case "":
                            var v106 bool = true
                            _ = v106
                        case "a":
                            switch (fi64 | 32768) {
                            case 2147483647:
                                if false {
                                    v25 += v25
                                }
                            case 513:
                                switch (("foo bar" + "a") + "\n\t") {
                                case "a":
                                    v107 := GenS6{F0: 127}
                                    _ = v107
                                default:
                                    v108 := GenS2{F0: int32(fi26), F1: 783}
                                    _ = v108
                                }
                            case -36:
                                var v109 int16 = 279
                                _ = v109
                            default:
                                v110 := GenS10{F0: (fi64 ^ 32768), F1: 4294967295, F2: 0}
                                _ = v110
                            }
                        case "foo bar":
                            v111 := GenS7{F0: "a", F1: -446}
                            _ = v111
                        }
                        if true {
                            v25 -= float64(0.1)
                            fmt.Println(fi26, fi64, fi26)
                        }
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                }
                fmt.Println(fi26, v25)
                if (true || ((-128 >= 1) && false)) {
                    fmt.Println(fi26, v25)
                } else {
                    if (!(42 == 272)) {
                        if (128 < 208) {
                            v112 := []int{fi26, (fi26 ^ 727)}
                            _ = v112
                        } else {
                            var v113 int16 = -999
                            _ = v113
                        }
                    }
                    for fi114 := 0; fi114 < 2; fi114++ {
                        v25 -= 0.1
                        _ = fi114
                    }
                }
                _ = fi26
            }
            v115 := GenS5{F0: 42, F1: 838}
            fmt.Println(v25)
            _ = v115
        } else {
            var v116 int32 = 516
            if false {
                v117 := GenS12{F0: ("\n\t" + (("a" + "") + "foo bar")), F1: 1}
                switch ("\n\t" + "\n\t") {
                case "a":
                    fmt.Println(v25, v25)
                default:
                    if ((0 > -5) || false) {
                        switch 2147483647 {
                        case 0:
                            v118 := GenS2{F0: v116, F1: v116}
                            _ = v118
                        case -179:
                            var v119 bool = (1 <= 0)
                            _ = v119
                        case -239:
                            var v120 int8 = 42
                            _ = v120
                        default:
                            for fi121 := 0; fi121 < 11; fi121++ {
                                for fi122 := 0; fi122 < 5; fi122++ {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v25 -= v25
                                    case "hello":
                                        switch 42 {
                                        case 74:
                                            var v123 uint = 454
                                            _ = v123
                                        case 4294967295:
                                            var v124 float32 = 1e10
                                            _ = v124
                                        default:
                                            v116 = 694
                                        }
                                    default:
                                        switch "foo bar" {
                                        case "foo bar":
                                            v125 := []int{fi122, (fi121 - 788), fi121}
                                            _ = v125
                                        case "\n\t":
                                            v25 = (v25 / 1.0)
                                        case "a":
                                            v126 := GenS13{F0: 42, F1: 613, F2: ""}
                                            _ = v126
                                        default:
                                            var v127 bool = true
                                            var v128 string = "a"
                                            _ = v127
                                            _ = v128
                                        }
                                    }
                                    v116 += (v116 - (v116 >> ((17) & 63)))
                                    fmt.Println(fi121, fi121)
                                    _ = fi122
                                }
                                for fi129 := 0; fi129 < 7; fi129++ {
                                    fmt.Println(fi129)
                                    if false {
                                        v25 += float64(0.1)
                                        if false {
                                            v130 := []int{-9223372036854775808, (fi121 / ((fi121) | 1)), (-(fi121))}
                                            v25 += -2.25
                                            if (0.0 >= v25) {
                                                v4[2147483648] = (v25 == float64(3.14159))
                                            } else {
                                                v25 -= (v25 - (3.14159 * (1e10 * v25)))
                                                v131 := GenS14{F0: 128}
                                                _ = v131
                                            }
                                            _ = v130
                                        }
                                    } else {
                                        if (-1 > -922) {
                                            var v132 uint64 = 18446744073709551615
                                            fmt.Println(v116)
                                            _ = v132
                                        }
                                        v133 := GenS5{F0: int64(fi121), F1: 977}
                                        v134 := []int{fi129, fi121, (-(fi129))}
                                        var v135 bool = (!true)
                                        switch "hello" {
                                        case "\n\t":
                                            switch (fi129 - fi129) {
                                            case -479:
                                                fmt.Println(v116, fi121, v135)
                                            case 32768:
                                                _ = len(v134)
                                            default:
                                                v25 = (0.1 - (v25 - v25))
                                            }
                                        case "":
                                            v25 = (float64((0.0 + 1e10)) + 0.1)
                                        case "hello":
                                            if v135 {
                                                v25 -= v25
                                                v136 := GenS1{F0: 255}
                                                _ = v136
                                            } else {
                                                var v137 bool = false
                                                var v138 uint16 = 0
                                                v139 := GenS15{F0: int32(fi129), F1: "\n\t"}
                                                _ = v137
                                                _ = v138
                                                _ = v139
                                            }
                                        default:
                                            v140 := GenS4{F0: 504, F1: (-(v116))}
                                            _ = v140
                                        }
                                        _ = v133
                                        _ = v134
                                        _ = v135
                                    }
                                    _ = fi129
                                }
                                if true {
                                    for fi141 := 0; fi141 < 2; fi141++ {
                                        if false {
                                            if true {
                                                delete(v4, 187)
                                                var v142 int16 = 32767
                                                _ = v142
                                            } else {
                                                var v143 uint8 = 128
                                                v1[47] = 4294967295
                                                if false {
                                                    var v144 int16 = 1000
                                                    if true {
                                                        if true {
                                                            var v145 bool = (v116 != v116)
                                                            var v146 float64 = v25
                                                            if (!v145) {
                                                                v143 += (v143 * (v143 >> ((0) & 63)))
                                                                if false {
                                                                    v4[2] = ((0 == int64(v144)) && v145)
                                                                    v147 := GenS16{F0: 98, F1: 0}
                                                                    v25 -= ((1.5 + 1.5) * v146)
                                                                    _ = v147
                                                                }
                                                                if (42 > 2) {
                                                                    v148 := GenS17{F0: 4294967295, F1: 0, F2: 42}
                                                                    switch 1 {
                                                                    case 32767:
                                                                        v143 = 0
                                                                    case 1:
                                                                        var v149 int16 = (v144 ^ (v144 * (v144 + (v144 - v144))))
                                                                        _ = v149
                                                                    }
                                                                    if (!(!(float64((0.0 + (0.0 - (-2.25 * (3.14159 - float32(1.5)))))) <= -2.25))) {
                                                                        v150 := GenS18{F0: 2}
                                                                        if (0 < 2) {
                                                                            v151 := []int{32768, 1, fi121}
                                                                            if v145 {
                                                                                var v152 uint16 = 65535
                                                                                var v153 int8 = 0
                                                                                var v154 int64 = 1
                                                                                _ = v152
                                                                                _ = v153
                                                                                _ = v154
                                                                            }
                                                                            var v155 int32 = (v116 % ((999) | 1))
                                                                            _ = v151
                                                                            _ = v155
                                                                        }
                                                                        _ = v150
                                                                    }
                                                                    _ = v148
                                                                }
                                                            }
                                                            _ = v145
                                                            _ = v146
                                                        } else {
                                                            var v156 uint16 = 813
                                                            v25 -= v25
                                                            _ = v156
                                                        }
                                                        v157 := GenS19{F0: fi121}
                                                        _ = v157
                                                    } else {
                                                        v116 = v116
                                                        switch 127 {
                                                        case 128:
                                                            delete(v3, 1)
                                                        case -780:
                                                            v158 := GenS1{F0: 28}
                                                            _ = v158
                                                        case 127:
                                                            var v159 uint = 4294967295
                                                            _ = v159
                                                        default:
                                                            v25 = 3.14159
                                                        }
                                                        v160 := GenS20{F0: 32767}
                                                        _ = v160
                                                    }
                                                    var v161 uint16 = 32767
                                                    fmt.Println(v144, v25, fi141)
                                                    v144 += v144
                                                    _ = v144
                                                    _ = v161
                                                } else {
                                                    v162 := GenS12{F0: "\n\t", F1: -1}
                                                    v163 := GenS13{F0: 0, F1: 207, F2: ("hello" + (("a" + "\n\t") + ""))}
                                                    _ = v162
                                                    _ = v163
                                                }
                                                _ = v143
                                            }
                                            if (v116 < 886) {
                                                var v164 bool = false
                                                _ = v164
                                            } else {
                                                if false {
                                                    var v165 uint = 255
                                                    if true {
                                                        var v166 uint = v165
                                                        v167 := []int{}
                                                        fmt.Println(v166, v25, v25)
                                                        var v168 uint16 = 255
                                                        fmt.Println(v168)
                                                        if ((v116 >> ((v165) & 63)) == (-(v116))) {
                                                            if (v168 > v168) {
                                                                if false {
                                                                    v116 = 84
                                                                    v165 -= v165
                                                                    v168 = (v168 >> ((82) & 63))
                                                                }
                                                            } else {
                                                                v25 = -2.25
                                                                v169 := v2[54]
                                                                var v170 uint8 = 216
                                                                _ = v169
                                                                _ = v170
                                                            }
                                                            var v171 uint = (^(v166))
                                                            _ = v171
                                                        } else {
                                                            fmt.Println(fi141)
                                                            var v172 string = ("a" + "")
                                                            if (!(!(98 <= 0))) {
                                                                v172 = (v172 + "foo bar")
                                                                v25 = v25
                                                                v25 -= v25
                                                            } else {
                                                                var v173 string = ("foo bar" + v172)
                                                                _ = v173
                                                            }
                                                            v116 = v116
                                                            if (false && false) {
                                                                if true {
                                                                    var v174 int16 = 42
                                                                    if (-885 > (v116 * v116)) {
                                                                        var v175 uint = v165
                                                                        _ = v175
                                                                    } else {
                                                                        v2[2147483647] = (-951 >= int64(v174))
                                                                        v176 := GenS15{F0: v116, F1: (v172 + v172)}
                                                                        _ = v176
                                                                    }
                                                                    var v177 string = v172
                                                                    switch ("foo bar" + "") {
                                                                    case "hello":
                                                                        var v178 int64 = -9223372036854775808
                                                                        _ = v178
                                                                    case "a":
                                                                        var v179 int32 = -1
                                                                        var v180 uint16 = (v168 * (^(v168)))
                                                                        _ = v179
                                                                        _ = v180
                                                                    case "foo bar":
                                                                        var v181 uint64 = 1
                                                                        _ = v181
                                                                    default:
                                                                        fmt.Println(v25)
                                                                    }
                                                                    _ = v174
                                                                    _ = v177
                                                                } else {
                                                                    if false {
                                                                        if false {
                                                                            if len(v167) >= 2 {
                                                                                v167 = v167[1:len(v167)]
                                                                            }
                                                                            fmt.Println(v165, v166, v116)
                                                                            v1[42] = 4294967295
                                                                            _ = len(v167)
                                                                            switch ("foo bar" + "\n\t") {
                                                                            case "foo bar":
                                                                                fmt.Println(v172, v165, fi121)
                                                                            case "":
                                                                                v168 -= (^(v168))
                                                                            default:
                                                                                v168 = v168
                                                                            }
                                                                        } else {
                                                                            var v182 bool = false
                                                                            switch v172 {
                                                                            case "\n\t":
                                                                                var v183 uint = 42
                                                                                _ = v183
                                                                            }
                                                                            fmt.Println(v172, fi121, v25)
                                                                            _ = v182
                                                                        }
                                                                        var v184 uint = (v166 | (^(v165)))
                                                                        v184 = v166
                                                                        _ = v184
                                                                    } else {
                                                                        v185 := v2[2]
                                                                        var v186 int16 = 273
                                                                        _ = v185
                                                                        _ = v186
                                                                    }
                                                                    var v187 uint64 = 42
                                                                    v165 += (v165 % ((v165) | 1))
                                                                    _ = v187
                                                                }
                                                                switch (^(fi141)) {
                                                                case -842:
                                                                    var v188 int64 = -670
                                                                    _ = v188
                                                                default:
                                                                    var v189 int32 = (v116 << (((v166 + v165)) & 63))
                                                                    _ = v189
                                                                }
                                                            } else {
                                                                if false {
                                                                    switch v172 {
                                                                    case "a":
                                                                        var v190 float64 = (v25 * v25)
                                                                        _ = v190
                                                                    case "\n\t":
                                                                        _ = len(v2)
                                                                    case "hello":
                                                                        switch fi141 {
                                                                        case 128:
                                                                            var v191 float32 = -2.25
                                                                            _ = v191
                                                                        case 673:
                                                                            var v192 float32 = 1e10
                                                                            _ = v192
                                                                        default:
                                                                            if ((fi141 % ((fi121) | 1)) >= fi121) {
                                                                                fmt.Println(fi141, v166, v116)
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                switch fi121 {
                                                                case 4294967295:
                                                                    v193 := GenS21{F0: v166, F1: v116}
                                                                    _ = v193
                                                                default:
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        v194 := GenS14{F0: 127}
                                                                        _ = v194
                                                                    }
                                                                    switch "" {
                                                                    case "hello":
                                                                        v25 -= v25
                                                                    case "foo bar":
                                                                        v165 -= 65535
                                                                    default:
                                                                        switch fi121 {
                                                                        case 127:
                                                                            var v195 int64 = int64(v166)
                                                                            _ = v195
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            _ = v172
                                                        }
                                                        _ = v166
                                                        _ = v167
                                                        _ = v168
                                                    }
                                                    _ = v165
                                                }
                                                v196 := GenS5{F0: 42, F1: 9223372036854775807}
                                                if false {
                                                    var v197 uint = 605
                                                    switch ("a" + "") {
                                                    case "\n\t":
                                                        delete(v1, 2)
                                                    case "hello":
                                                        var v198 int8 = 0
                                                        _ = v198
                                                    case "a":
                                                        switch fi141 {
                                                        case -6:
                                                            v199 := []int{fi141, fi121}
                                                            _ = v199
                                                        case 127:
                                                            v116 += (v116 >> (((v197 << ((42) & 63))) & 63))
                                                        }
                                                    }
                                                    if false {
                                                        v25 += (v25 - v25)
                                                        v200 := GenS22{F0: 12, F1: 0}
                                                        var v201 bool = true
                                                        _ = v200
                                                        _ = v201
                                                    }
                                                    var v202 int32 = 0
                                                    _ = v197
                                                    _ = v202
                                                }
                                                v116 += v116
                                                _ = v196
                                            }
                                        }
                                        v203 := GenS23{F0: -322}
                                        if (1 >= 699) {
                                            v116 -= (-(v116))
                                            v204 := []int{(fi121 / ((fi121) | 1))}
                                            var v205 uint = 1
                                            _ = v204
                                            _ = v205
                                        } else {
                                            switch ("foo bar" + "hello") {
                                            case "hello":
                                                v206 := []int{(fi121 >> ((255) & 63)), 42, fi121}
                                                if true {
                                                    if true {
                                                        if false {
                                                            v25 = (3.14159 / 1.5)
                                                            v206 = append(v206, fi141)
                                                        } else {
                                                            if (-42 == 0) {
                                                                fmt.Println(fi141, v25)
                                                                if false {
                                                                    v206 = append(v206, (fi121 - -1))
                                                                    if len(v206) > 0 {
                                                                        v206[0] = (fi121 * fi141)
                                                                    }
                                                                    v25 -= v25
                                                                }
                                                                fmt.Println(fi121)
                                                            }
                                                            var v207 int64 = 102
                                                            v25 -= 0.1
                                                            _ = v207
                                                        }
                                                        switch (("" + "foo bar") + "a") {
                                                        case "hello":
                                                            fmt.Println(fi141)
                                                        case "":
                                                            v25 -= 0.0
                                                        }
                                                    }
                                                    fmt.Println(fi141)
                                                }
                                                _ = v206
                                            }
                                            fmt.Println(fi121, v25, fi121)
                                            if (0 > 2) {
                                                if false {
                                                    v208 := GenS24{F0: int8(fi121)}
                                                    v25 += v25
                                                    v209 := GenS25{F0: (v116 & (v116 ^ (^(v116))))}
                                                    v210 := GenS26{F0: 389, F1: true}
                                                    _ = v208
                                                    _ = v209
                                                    _ = v210
                                                } else {
                                                    fmt.Println(fi141, fi141, fi121)
                                                    var v211 float64 = (v25 * 3.14159)
                                                    var v212 uint = 65535
                                                    _ = v211
                                                    _ = v212
                                                }
                                                _ = len(v4)
                                                var v213 string = ((("a" + "") + "hello") + "")
                                                _ = v213
                                            } else {
                                                v116 = v116
                                            }
                                            v25 = v25
                                            var v214 int64 = -9223372036854775808
                                            _ = v214
                                        }
                                        _ = fi141
                                        _ = v203
                                    }
                                    v215 := []int{fi121, (fi121 - (^(fi121)))}
                                    v216 := GenS23{F0: -687}
                                    _ = v215
                                    _ = v216
                                }
                                _ = fi121
                            }
                        }
                        var v217 bool = true
                        fmt.Println(v217)
                        _ = v217
                    } else {
                        for fi218 := 0; fi218 < 7; fi218++ {
                            v4[213] = (v25 == (v25 * v25))
                            var v219 int = (fi218 / (((fi218 + 42)) | 1))
                            if false {
                                var v220 uint8 = 128
                                v25 -= (0.1 + 1e10)
                                switch "\n\t" {
                                case "":
                                    v221 := []int{}
                                    _ = v221
                                case "foo bar":
                                    if false {
                                        for fi222 := 0; fi222 < 13; fi222++ {
                                            fmt.Println(fi218, fi222, v25)
                                            delete(v1, 0)
                                            _ = fi222
                                        }
                                        v223 := GenS27{F0: -9223372036854775808, F1: (v116 + (v116 & (v116 + (v116 | v116))))}
                                        _ = v223
                                    }
                                case "a":
                                    if false {
                                        for fi224 := 0; fi224 < 6; fi224++ {
                                            var v225 uint8 = v220
                                            v225 = v220
                                            _ = fi224
                                            _ = v225
                                        }
                                    }
                                }
                                _ = v220
                            } else {
                                for fi226 := 0; fi226 < 9; fi226++ {
                                    var v227 uint32 = 4294967295
                                    _ = fi226
                                    _ = v227
                                }
                                v25 -= (v25 - v25)
                                v228 := GenS5{F0: -815, F1: 147}
                                _ = v228
                            }
                            _ = fi218
                            _ = v219
                        }
                        switch (("a" + (("a" + ("foo bar" + "\n\t")) + ((("a" + ("" + "a")) + "") + "foo bar"))) + "") {
                        case "hello":
                            if true {
                                var v229 uint16 = 65535
                                v1[684] = 4294967295
                                _ = v229
                            }
                        case "\n\t":
                            fmt.Println(v25, v25)
                        }
                    }
                }
                v230 := GenS28{F0: 226, F1: 156, F2: 196}
                _ = v117
                _ = v230
            } else {
                var v231 float32 = ((float32(3.14159) + ((3.14159 - float32((((((float64(((1.5 + 0.1) - 0.1)) * 1e10) - v25) / float64(0.1)) / (3.14159 / v25)) / float64(1.5)))) - 0.0)) + 0.0)
                _ = v231
            }
            _ = v116
        }
        if true {
            var v232 uint16 = 1
            var v233 float32 = 0.1
            _ = v232
            _ = v233
        } else {
            switch ((("a" + "hello") + ("" + "hello")) + "hello") {
            case "":
                var v234 uint8 = 255
                _ = v234
            case "foo bar":
                v235 := GenS19{F0: 128}
                _ = v235
            case "\n\t":
                switch 618 {
                case 32767:
                    for fi236 := 0; fi236 < 14; fi236++ {
                        v237 := GenS17{F0: 4294967295, F1: 2, F2: 9223372036854775807}
                        var v238 uint64 = 4294967295
                        _ = fi236
                        _ = v237
                        _ = v238
                    }
                case 18:
                    for fi239 := 0; fi239 < 14; fi239++ {
                        for fi240 := 0; fi240 < 10; fi240++ {
                            switch "foo bar" {
                            case "\n\t":
                                var v241 float32 = -2.25
                                _ = v241
                            default:
                                v242 := []int{(fi240 - fi240), fi239}
                                _ = v242
                            }
                            var v243 uint8 = 128
                            var v244 string = "\n\t"
                            _ = fi240
                            _ = v243
                            _ = v244
                        }
                        _ = fi239
                    }
                case 317:
                    fmt.Println(v25)
                default:
                    v245 := GenS29{F0: -1, F1: 42}
                    _ = v245
                }
            }
            v25 = v25
            for fi246 := 0; fi246 < 7; fi246++ {
                switch fi246 {
                case -424:
                    var v247 uint64 = 1
                    _ = v247
                case 968:
                    var v248 int8 = 42
                    _ = v248
                case 128:
                    v249 := GenS30{F0: 255, F1: -1, F2: 1}
                    _ = v249
                default:
                    var v250 string = ""
                    _ = v250
                }
                for fi251 := 0; fi251 < 11; fi251++ {
                    var v252 float32 = float32((v25 / v25))
                    var v253 uint64 = 18446744073709551615
                    v253 = (^(v253))
                    _ = fi251
                    _ = v252
                    _ = v253
                }
                var v254 float64 = 0.0
                _ = fi246
                _ = v254
            }
        }
        delete(v4, 1)
    } else {
        switch (("a" + (("" + ("\n\t" + "a")) + "")) + "") {
        case "\n\t":
            v255 := GenS31{F0: 255, F1: 42}
            _ = v255
        }
        var v256 bool = (true || (true && (v25 >= v25)))
        fmt.Println(v256)
        _ = v256
    }
    v25 = (1e10 - float64(float32(v25)))
    for fi257 := 0; fi257 < 11; fi257++ {
        var v258 uint8 = 1
        v259 := []int{fi257, (fi257 | (^(fi257))), (fi257 * (-(fi257))), fi257}
        _ = fi257
        _ = v258
        _ = v259
    }
    v25 -= -2.25
    var v260 int16 = 242
    v261 := GenS31{F0: 26, F1: 2147483647}
    v25 = -2.25
    switch 2147483647 {
    case 32767:
        switch ("a" + ("foo bar" + "a")) {
        case "hello":
            v25 = (v25 - float64(0.0))
        case "a":
            v25 -= v25
        }
    case -9223372036854775808:
        if false {
            fmt.Println(v260, v260)
        }
    }
    v260 = v260
    switch 9223372036854775807 {
    case -1:
        v260 = 564
    case -622:
        v260 += v260
    }
    switch -891 {
    case -2147483648:
        switch "" {
        case "hello":
            var v262 int16 = 735
            _ = v262
        case "\n\t":
            fmt.Println(v260, v260)
        }
    }
    for fi263 := 0; fi263 < 1; fi263++ {
        var v264 float64 = float64(3.14159)
        switch (fi263 + 32767) {
        case 820:
            switch ("a" + "hello") {
            case "\n\t":
                var v265 float64 = 1.5
                _ = v265
            default:
                fmt.Println(v264, fi263, v260)
            }
        }
        for fi266 := 0; fi266 < 12; fi266++ {
            var v267 uint64 = 0
            _ = fi266
            _ = v267
        }
        _ = fi263
        _ = v264
    }
    v260 = v260
    v260 -= v260
    switch -9223372036854775808 {
    case 847:
        if true {
            var v268 int8 = 42
            if true {
                for fi269 := 0; fi269 < 19; fi269++ {
                    if (232 == 1) {
                        var v270 uint32 = 2
                        var v271 uint64 = 421
                        _ = v270
                        _ = v271
                    } else {
                        v272 := []int{}
                        v273 := GenS20{F0: 255}
                        _ = v272
                        _ = v273
                    }
                    var v274 uint = 1
                    _ = fi269
                    _ = v274
                }
                fmt.Println(v25, v268)
                switch 4294967295 {
                case 1:
                    v275 := GenS32{F0: -296, F1: 65535}
                    _ = v275
                case 42:
                    var v276 uint64 = 18446744073709551615
                    _ = v276
                case -2147483648:
                    v260 = -1
                }
            }
            _ = v268
        } else {
            var v277 int32 = 42
            switch -2147483648 {
            case 546:
                switch 456 {
                case 643:
                    var v278 int = -2147483648
                    _ = v278
                case 127:
                    var v279 int8 = 0
                    _ = v279
                }
            case 128:
                for fi280 := 0; fi280 < 12; fi280++ {
                    var v281 int = (fi280 - (fi280 % ((-99) | 1)))
                    v282 := v4[2147483648]
                    _ = fi280
                    _ = v281
                    _ = v282
                }
            default:
                var v283 string = ((("a" + ((("\n\t" + "\n\t") + ("foo bar" + ("a" + (((((("a" + "foo bar") + ("hello" + "a")) + "a") + (((("hello" + "foo bar") + "hello") + ("a" + (("a" + (("\n\t" + "\n\t") + "hello")) + ("hello" + "")))) + ("\n\t" + ""))) + "\n\t") + (((("hello" + "hello") + "foo bar") + "a") + "\n\t"))))) + "")) + "foo bar") + "hello")
                _ = v283
            }
            fmt.Println(v25)
            _ = v277
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v260:", v260)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v260
    _ = v261
    _ = v3
    _ = v4
}
