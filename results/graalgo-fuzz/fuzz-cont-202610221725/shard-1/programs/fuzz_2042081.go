package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 bool; F2 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 int64; F2 int64 }
type GenS8 struct { F0 uint64; F1 int8 }
type GenS9 struct { F0 int64; F1 int16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint8; F1 uint16 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 int32; F1 int64; F2 uint32 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 int16; F1 int }
type GenS20 struct { F0 uint32; F1 uint32 }
type GenS21 struct { F0 bool; F1 int; F2 uint16 }
type GenS22 struct { F0 string; F1 int }
type GenS23 struct { F0 uint64; F1 bool }
type GenS24 struct { F0 string }
type GenS25 struct { F0 uint64; F1 int32; F2 int64 }
type GenS26 struct { F0 int32; F1 int32; F2 string }
type GenS27 struct { F0 uint32 }
type GenS28 struct { F0 uint64; F1 uint }
type GenS29 struct { F0 uint32 }

func main() {
    var v1 uint = 661
    var v2 string = "foo bar"
    v3 := make(map[string]bool)
    v4 := make(map[uint32]string)
    switch "hello" {
    case "a":
        switch ((v2 + "a") + (("a" + ("a" + v2)) + "\n\t")) {
        case "hello":
            v5 := GenS1{F0: "foo bar", F1: 591, F2: -1}
            _ = v5
        case "foo bar":
            v6 := v3[v2]
            _ = v6
        default:
            var v7 uint16 = 997
            _ = v7
        }
    case "\n\t":
        switch 1 {
        case -381:
            var v8 float32 = (-2.25 / float32((3.14159 * 3.14159)))
            _ = v8
        case 42:
            v2 = "a"
        }
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := GenS2{F0: 599}
        v11 := GenS2{F0: 193}
        _ = fi9
        _ = v10
        _ = v11
    }
    if true {
        v12 := GenS3{F0: -590}
        fmt.Println(v1, v1)
        _ = v12
    } else {
        fmt.Println(v1)
        for fi13 := 0; fi13 < 9; fi13++ {
            switch v2 {
            case "foo bar":
                if (!false) {
                    v1 -= v1
                    v14 := GenS1{F0: ("" + "hello"), F1: 143, F2: -9223372036854775808}
                    switch (v2 + v2) {
                    case "":
                        fmt.Println(v1)
                    case "foo bar":
                        v15 := GenS4{F0: "foo bar", F1: ((!true) && (false && false)), F2: int16(v1)}
                        _ = v15
                    case "a":
                        fmt.Println(v1, v1)
                    }
                    _ = v14
                } else {
                    v2 = v2
                    v3[v2] = (2 >= 65535)
                    v1 = v1
                }
            case "a":
                for fi16 := 0; fi16 < 6; fi16++ {
                    v17 := GenS5{F0: 0}
                    var v18 int16 = 32767
                    switch fi16 {
                    case 9223372036854775807:
                        v18 = (^(v18))
                    case 127:
                        var v19 int = -504
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            default:
                if (1 == 1) {
                    for fi20 := 0; fi20 < 18; fi20++ {
                        v2 = (v2 + v2)
                        _ = fi20
                    }
                    v21 := GenS2{F0: 2}
                    _ = v21
                } else {
                    var v22 string = "a"
                    switch (v22 + (((v2 + "hello") + (((("foo bar" + "\n\t") + v2) + "") + (v22 + v2))) + v22)) {
                    case "a":
                        switch fi13 {
                        case 64:
                            for fi23 := 0; fi23 < 4; fi23++ {
                                var v24 int64 = 1
                                _ = fi23
                                _ = v24
                            }
                        case 54:
                            var v25 int16 = -692
                            _ = v25
                        case -9223372036854775808:
                            for fi26 := 0; fi26 < 18; fi26++ {
                                switch fi13 {
                                case -2147483648:
                                    fmt.Println(fi26, v1)
                                default:
                                    if true {
                                        var v27 string = "a"
                                        v3[v2] = (1 > 42)
                                        if true {
                                            var v28 int32 = -1
                                            var v29 uint = (v1 >> ((4294967295) & 63))
                                            v29 = 2
                                            _ = v28
                                            _ = v29
                                        }
                                        fmt.Println(fi26, v27)
                                        _ = v27
                                    }
                                }
                                var v30 uint16 = 65535
                                var v31 uint32 = 2147483647
                                v30 += 869
                                _ = fi26
                                _ = v30
                                _ = v31
                            }
                        }
                    case "hello":
                        for fi32 := 0; fi32 < 17; fi32++ {
                            v2 = (v22 + ("hello" + "\n\t"))
                            _ = fi32
                        }
                    case "foo bar":
                        v3[v22] = (false && (true && ((v22 + ("\n\t" + v22)) >= v2)))
                    default:
                        v2 = v2
                    }
                    _ = v22
                }
            }
            _ = fi13
        }
    }
    var v33 bool = (1000 == 1000)
    for fi34 := 0; fi34 < 18; fi34++ {
        for fi35 := 0; fi35 < 1; fi35++ {
            var v36 bool = ((1 >= 65535) || false)
            _ = fi35
            _ = v36
        }
        fmt.Println(v33, v2)
        _ = fi34
    }
    v37 := GenS3{F0: 42}
    switch 4294967295 {
    case 0:
        for fi38 := 0; fi38 < 20; fi38++ {
            fmt.Println(v33, v1)
            _ = fi38
        }
    case 628:
        v39 := []int{521, 128, -9223372036854775808, 0}
        _ = v39
    }
    v1 -= v1
    v40 := GenS4{F0: "foo bar", F1: v33, F2: 1}
    switch 42 {
    case -2147483648:
        v1 = 18446744073709551615
    case 868:
        for fi41 := 0; fi41 < 10; fi41++ {
            var v42 int = fi41
            for fi43 := 0; fi43 < 3; fi43++ {
                var v44 int = -899
                v45 := GenS6{F0: 42}
                v46 := []int{v42, fi41, (v44 | -742)}
                _ = fi43
                _ = v44
                _ = v45
                _ = v46
            }
            fmt.Println(fi41, v42, v42)
            _ = fi41
            _ = v42
        }
    case -901:
        if true {
            for fi47 := 0; fi47 < 8; fi47++ {
                switch v2 {
                case "\n\t":
                    var v48 int8 = 58
                    _ = v48
                case "hello":
                    fmt.Println(v33)
                default:
                    for fi49 := 0; fi49 < 2; fi49++ {
                        if true {
                            v1 -= 100
                        } else {
                            var v50 uint64 = 42
                            v51 := GenS3{F0: -91}
                            if ((42 < -2147483648) && (v33 || false)) {
                                if (v33 || (98 >= 4294967295)) {
                                    var v52 uint8 = 135
                                    switch fi49 {
                                    case 1:
                                        v53 := v3[v2]
                                        _ = v53
                                    case 2147483647:
                                        v54 := []int{fi47, (fi47 + -900), -877, 32768}
                                        _ = v54
                                    case -9223372036854775808:
                                        var v55 int32 = -2147483648
                                        _ = v55
                                    }
                                    _ = v52
                                } else {
                                    v56 := []int{}
                                    _ = v56
                                }
                                switch v2 {
                                case "a":
                                    var v57 int64 = -1
                                    _ = v57
                                case "\n\t":
                                    switch ("foo bar" + "foo bar") {
                                    case "\n\t":
                                        if v33 {
                                            var v58 int16 = 860
                                            _ = v58
                                        } else {
                                            v33 = (((1e10 - 3.14159) / (-2.25 - (((1e10 / (((0.0 - (((0.1 * 0.0) - ((-2.25 / (3.14159 * ((1.5 / 3.14159) * 1.5))) / (-2.25 * (-2.25 / float32((1e10 - 0.0)))))) - (float32(3.14159) - (1.5 * float32((0.1 * (((0.1 - 0.0) * float64(float32((3.14159 * 0.1)))) / 3.14159))))))) * 0.1) * 1e10)) / ((0.1 * (0.1 * (-2.25 / 0.1))) - 0.0)) * (0.0 - ((1e10 + float32(0.1)) * 1e10))))) != -2.25)
                                        }
                                    case "a":
                                        v3[("hello" + (v2 + (("\n\t" + v2) + v2)))] = v33
                                    default:
                                        var v59 uint32 = uint32(v1)
                                        _ = v59
                                    }
                                case "hello":
                                    var v60 bool = ((v1 * (v1 << ((0) & 63))) != 220)
                                    _ = v60
                                default:
                                    if v33 {
                                        if ((v50 >> (((v1 * 1)) & 63)) < (^(v50))) {
                                            var v61 int8 = -128
                                            var v62 uint32 = 623
                                            _ = v61
                                            _ = v62
                                        }
                                        if false {
                                            if true {
                                                switch (((v2 + "foo bar") + v2) + "a") {
                                                case "\n\t":
                                                    if false {
                                                        v63 := GenS7{F0: 2, F1: 9223372036854775807, F2: 221}
                                                        _ = v63
                                                    } else {
                                                        v33 = v33
                                                    }
                                                default:
                                                    if true {
                                                        var v64 int8 = 80
                                                        _ = v64
                                                    } else {
                                                        v50 = 4294967295
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v1)
                                                if true {
                                                    switch (fi47 | fi49) {
                                                    case -932:
                                                        switch -416 {
                                                        case 4294967295:
                                                            v2 = ""
                                                        }
                                                    case 157:
                                                        v2 = v2
                                                    case 32767:
                                                        switch (((v2 + "a") + "") + v2) {
                                                        case "hello":
                                                            v1 += v1
                                                        default:
                                                            var v65 int32 = 0
                                                            _ = v65
                                                        }
                                                    }
                                                    if (int64(v1) <= -1) {
                                                        v50 = v50
                                                    }
                                                } else {
                                                    fmt.Println(fi47, fi47)
                                                }
                                                var v66 uint8 = 169
                                                _ = v66
                                            }
                                            switch (fi49 % (((fi49 - (-(fi47)))) | 1)) {
                                            case 2147483647:
                                                switch ((v2 + v2) + v2) {
                                                case "\n\t":
                                                    if v33 {
                                                        v33 = (65535 > 177)
                                                        var v67 int32 = -860
                                                        _ = v67
                                                    } else {
                                                        var v68 int32 = 2147483647
                                                        var v69 bool = (!(v33 || (!v33)))
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                case "hello":
                                                    var v70 uint = (v1 % ((765) | 1))
                                                    _ = v70
                                                case "a":
                                                    if false {
                                                        if v33 {
                                                            switch v2 {
                                                            case "hello":
                                                                if false {
                                                                    var v71 uint = v1
                                                                    if (255 > 1) {
                                                                        v71 = (v1 ^ v1)
                                                                    } else {
                                                                        if (285 >= (fi49 - fi47)) {
                                                                            switch (fi49 / ((32768) | 1)) {
                                                                            case 402:
                                                                                var v72 float32 = -2.25
                                                                                var v73 int32 = 0
                                                                                switch fi47 {
                                                                                case 0:
                                                                                    var v74 int64 = 9223372036854775807
                                                                                    _ = v74
                                                                                case 32767:
                                                                                    var v75 uint16 = 195
                                                                                    v3[v2] = (42 < -1)
                                                                                    _ = v75
                                                                                }
                                                                                _ = v72
                                                                                _ = v73
                                                                            case 2147483647:
                                                                                var v76 float32 = 0.1
                                                                                if (((3.14159 - (float32(-2.25) - v76)) / (1e10 + -2.25)) <= v76) {
                                                                                    _ = len(v3)
                                                                                    if (1e10 > 3.14159) {
                                                                                        if (false || ((!v33) && true)) {
                                                                                            v2 = "hello"
                                                                                            var v77 int8 = 42
                                                                                            _ = v77
                                                                                        } else {
                                                                                            v78 := []int{(fi49 >> (((v1 | 429)) & 63))}
                                                                                            if v33 {
                                                                                                var v79 string = ((v2 + "a") + "")
                                                                                                _ = v79
                                                                                            } else {
                                                                                                v80 := GenS8{F0: 2, F1: 100}
                                                                                                _ = v80
                                                                                            }
                                                                                            if v33 {
                                                                                                if (32767 != -32768) {
                                                                                                    fmt.Println(v1, v2)
                                                                                                    if (1e10 <= (3.14159 - float64(1.5))) {
                                                                                                        if (1 == 1) {
                                                                                                            v50 += 575
                                                                                                            v81 := GenS5{F0: v50}
                                                                                                            _ = v81
                                                                                                        } else {
                                                                                                            fmt.Println(v76, fi49, fi49)
                                                                                                        }
                                                                                                        v2 = "a"
                                                                                                    }
                                                                                                    if v33 {
                                                                                                        v33 = true
                                                                                                        if false {
                                                                                                            if (4294967295 < v50) {
                                                                                                                var v82 bool = v33
                                                                                                                if len(v78) >= 2 {
                                                                                                                    v78 = v78[1:len(v78)]
                                                                                                                }
                                                                                                                v71 = v71
                                                                                                                _ = v82
                                                                                                            }
                                                                                                            delete(v4, 2147483648)
                                                                                                        }
                                                                                                        if v33 {
                                                                                                            if false {
                                                                                                                switch "hello" {
                                                                                                                case "foo bar":
                                                                                                                    if (!false) {
                                                                                                                        v76 = v76
                                                                                                                        v83 := GenS9{F0: 42, F1: 32767}
                                                                                                                        fmt.Println(v76, v71)
                                                                                                                        v76 = v76
                                                                                                                        v3[v2] = (true || (v33 || v33))
                                                                                                                        _ = v83
                                                                                                                    }
                                                                                                                default:
                                                                                                                    var v84 int32 = 0
                                                                                                                    _ = v84
                                                                                                                }
                                                                                                                v76 += v76
                                                                                                                v50 = v50
                                                                                                                var v85 uint64 = (v50 % (((v50 | (^(v50)))) | 1))
                                                                                                                fmt.Println(v76, v50, fi47)
                                                                                                                v1 -= (v1 + (^(v1)))
                                                                                                                _ = v85
                                                                                                            } else {
                                                                                                                var v86 uint32 = 570
                                                                                                                _ = v86
                                                                                                            }
                                                                                                            v71 -= v71
                                                                                                        } else {
                                                                                                            v78 = append(v78, (-(fi49)))
                                                                                                            fmt.Println(v1)
                                                                                                            var v87 int = (fi49 ^ (fi47 & 9223372036854775807))
                                                                                                            _ = v87
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                var v88 int64 = 0
                                                                                                switch (fi49 / ((fi47) | 1)) {
                                                                                                case 42:
                                                                                                    fmt.Println(v2, v71)
                                                                                                case 0:
                                                                                                    fmt.Println(v88)
                                                                                                    switch fi47 {
                                                                                                    case 274:
                                                                                                        switch fi47 {
                                                                                                        case 32768:
                                                                                                            switch (fi47 - fi47) {
                                                                                                            case -992:
                                                                                                                v33 = (!(!false))
                                                                                                            default:
                                                                                                                if len(v78) > 0 {
                                                                                                                    v78[0] = 128
                                                                                                                }
                                                                                                            }
                                                                                                        case 1:
                                                                                                            v71 = (v71 * (^(v71)))
                                                                                                        case -9223372036854775808:
                                                                                                            v88 = (-(v88))
                                                                                                        }
                                                                                                    case 127:
                                                                                                        var v89 uint64 = v50
                                                                                                        _ = v89
                                                                                                    case -1:
                                                                                                        if true {
                                                                                                            var v90 uint16 = 42
                                                                                                            var v91 bool = v33
                                                                                                            fmt.Println(v50, v76)
                                                                                                            _ = v90
                                                                                                            _ = v91
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                v88 = (v88 << ((v71) & 63))
                                                                                                _ = v88
                                                                                            } else {
                                                                                                fmt.Println(v33, v2)
                                                                                                var v92 bool = v33
                                                                                                _ = v92
                                                                                            }
                                                                                            _ = v78
                                                                                        }
                                                                                        var v93 uint = 652
                                                                                        _ = v93
                                                                                    } else {
                                                                                        if (1 > 1) {
                                                                                            if true {
                                                                                                v94 := GenS9{F0: 9223372036854775807, F1: 32767}
                                                                                                v95 := GenS10{F0: 4294967295}
                                                                                                var v96 int32 = 42
                                                                                                v97 := []int{fi47, fi47, -429}
                                                                                                _ = v94
                                                                                                _ = v95
                                                                                                _ = v96
                                                                                                _ = v97
                                                                                            }
                                                                                        } else {
                                                                                            var v98 uint = 1
                                                                                            switch v2 {
                                                                                            case "a":
                                                                                                if v33 {
                                                                                                    var v99 int16 = -1
                                                                                                    v100 := GenS3{F0: v99}
                                                                                                    v101 := []int{fi47, fi49, fi49}
                                                                                                    var v102 uint8 = uint8(v98)
                                                                                                    if true {
                                                                                                        v1 = (^(v71))
                                                                                                    }
                                                                                                    _ = v100
                                                                                                    _ = v101
                                                                                                    _ = v102
                                                                                                    _ = v99
                                                                                                } else {
                                                                                                    v3[("hello" + (("a" + (v2 + v2)) + (v2 + v2)))] = true
                                                                                                    var v103 uint8 = 42
                                                                                                    _ = v103
                                                                                                }
                                                                                            case "hello":
                                                                                                _ = len(v3)
                                                                                            }
                                                                                            _ = v98
                                                                                        }
                                                                                        v50 += (^(v50))
                                                                                        if (!true) {
                                                                                            v104 := []int{(fi47 - fi49), 128, fi47}
                                                                                            _ = v104
                                                                                        } else {
                                                                                            v105 := GenS11{F0: 2147483648}
                                                                                            _ = v105
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v76
                                                                            }
                                                                        } else {
                                                                            var v106 uint32 = 989
                                                                            _ = v106
                                                                        }
                                                                        v107 := []int{1, fi47, fi47, (fi49 ^ 127)}
                                                                        _ = v107
                                                                    }
                                                                    _ = v71
                                                                }
                                                            case "a":
                                                                var v108 int32 = -2147483648
                                                                _ = v108
                                                            case "\n\t":
                                                                v33 = ((0 == 99) || (v2 >= v2))
                                                            }
                                                        } else {
                                                            if false {
                                                                var v109 uint = (^(v1))
                                                                v50 = v50
                                                                switch -422 {
                                                                case 256:
                                                                    v110 := []int{(fi49 % ((fi47) | 1)), (-(fi49)), 127, (fi47 % ((fi49) | 1))}
                                                                    _ = v110
                                                                case -9223372036854775808:
                                                                    fmt.Println(v109, v2, v1)
                                                                case 32767:
                                                                    if true {
                                                                        var v111 string = ""
                                                                        switch fi47 {
                                                                        case -74:
                                                                            v2 = (v111 + v2)
                                                                            fmt.Println(v1, v109)
                                                                            v50 = (v50 ^ (^(v50)))
                                                                            var v112 int = (-(fi49))
                                                                            if (!(v33 && (42 != 2147483648))) {
                                                                                if v33 {
                                                                                    if (v33 || false) {
                                                                                        v111 = ""
                                                                                        if (false || (v33 && (577 != uint16(v109)))) {
                                                                                            switch 0 {
                                                                                            case 203:
                                                                                                fmt.Println(v112, fi49, fi49)
                                                                                            case -85:
                                                                                                v1 -= (v1 / ((255) | 1))
                                                                                            }
                                                                                            if v33 {
                                                                                                var v113 uint16 = 2
                                                                                                _ = v113
                                                                                            }
                                                                                            v50 += (v50 * v50)
                                                                                        } else {
                                                                                            v2 = "foo bar"
                                                                                            if v33 {
                                                                                                fmt.Println(v50)
                                                                                                var v114 uint = 4294967295
                                                                                                if v33 {
                                                                                                    switch v112 {
                                                                                                    case 0:
                                                                                                        v115 := GenS7{F0: 411, F1: 199, F2: 9223372036854775807}
                                                                                                        _ = v115
                                                                                                    case 2147483647:
                                                                                                        var v116 int8 = 127
                                                                                                        switch fi47 {
                                                                                                        case -921:
                                                                                                            v117 := GenS3{F0: 1000}
                                                                                                            _ = v117
                                                                                                        case 1:
                                                                                                            if (-2.25 <= (-2.25 - ((0.0 + (1e10 + 1e10)) - 3.14159))) {
                                                                                                                v118 := []int{fi47, (fi49 >> (((v109 + v109)) & 63)), v112, int(v116)}
                                                                                                                v118 = append(v118, fi49)
                                                                                                                _ = v118
                                                                                                            }
                                                                                                        case 32767:
                                                                                                            if v33 {
                                                                                                                v119 := GenS12{F0: 2, F1: 65535, F2: 255}
                                                                                                                v120 := GenS10{F0: 42}
                                                                                                                var v121 uint64 = (v50 % ((v50) | 1))
                                                                                                                if (352 > 2) {
                                                                                                                    var v122 int = 4294967295
                                                                                                                    v123 := GenS7{F0: 1, F1: 437, F2: -1}
                                                                                                                    v124 := GenS13{F0: -397}
                                                                                                                    _ = v122
                                                                                                                    _ = v123
                                                                                                                    _ = v124
                                                                                                                } else {
                                                                                                                    v1 += 974
                                                                                                                    v125 := GenS14{F0: 1, F1: 0}
                                                                                                                    v126 := GenS15{F0: 809}
                                                                                                                    _ = v125
                                                                                                                    _ = v126
                                                                                                                }
                                                                                                                v127 := GenS16{F0: 1, F1: -9223372036854775808, F2: 65535}
                                                                                                                _ = v119
                                                                                                                _ = v120
                                                                                                                _ = v121
                                                                                                                _ = v127
                                                                                                            } else {
                                                                                                                fmt.Println(v2)
                                                                                                            }
                                                                                                        default:
                                                                                                            delete(v3, v2)
                                                                                                        }
                                                                                                        _ = v116
                                                                                                    case -1:
                                                                                                        if v33 {
                                                                                                            var v128 float64 = float64(float32(1e10))
                                                                                                            var v129 bool = v33
                                                                                                            _ = v128
                                                                                                            _ = v129
                                                                                                        }
                                                                                                        if v33 {
                                                                                                            var v130 int16 = -1
                                                                                                            var v131 float64 = ((3.14159 - 1e10) / (0.0 + (((-2.25 * 0.1) / (1.5 / (1.5 + ((1e10 + 0.1) / 1.5)))) / 1.0)))
                                                                                                            _ = v130
                                                                                                            _ = v131
                                                                                                        }
                                                                                                    default:
                                                                                                        var v132 float64 = 1.5
                                                                                                        var v133 uint16 = 65535
                                                                                                        v134 := GenS17{F0: 42}
                                                                                                        _ = v132
                                                                                                        _ = v133
                                                                                                        _ = v134
                                                                                                    }
                                                                                                    if (!(v33 || (true || (0 >= 32767)))) {
                                                                                                        v135 := v4[42]
                                                                                                        var v136 int8 = -1
                                                                                                        var v137 int16 = -594
                                                                                                        _ = v135
                                                                                                        _ = v136
                                                                                                        _ = v137
                                                                                                    }
                                                                                                    var v138 int64 = -445
                                                                                                    var v139 float64 = 3.14159
                                                                                                    _ = v138
                                                                                                    _ = v139
                                                                                                } else {
                                                                                                    v140 := GenS14{F0: 255, F1: 618}
                                                                                                    _ = v140
                                                                                                }
                                                                                                _ = v114
                                                                                            } else {
                                                                                                fmt.Println(v50, v112, v112)
                                                                                                var v141 int = (fi47 - 32767)
                                                                                                _ = v141
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        delete(v4, 335)
                                                                                    }
                                                                                    var v142 int32 = -1
                                                                                    v142 = (-(v142))
                                                                                    _ = v142
                                                                                } else {
                                                                                    if (1 <= -1) {
                                                                                        var v143 float32 = ((float32(float64(1.5)) * ((0.1 - (0.1 + 1.5)) / 1e10)) + (float32(1.5) * (0.1 * ((-2.25 + 1e10) + (3.14159 - ((1e10 * (-2.25 * 1e10)) + (3.14159 * 0.0)))))))
                                                                                        v144 := []int{(fi47 | (fi47 >> (((v1 + v1)) & 63))), (fi49 & (fi47 << ((42) & 63)))}
                                                                                        _ = v143
                                                                                        _ = v144
                                                                                    } else {
                                                                                        if false {
                                                                                            v3[v111] = (-9223372036854775808 > 42)
                                                                                            var v145 string = (v111 + ("foo bar" + v2))
                                                                                            v1 += v109
                                                                                            _ = v145
                                                                                        }
                                                                                        if (!v33) {
                                                                                            var v146 string = ((v2 + v2) + ("foo bar" + ""))
                                                                                            _ = v146
                                                                                        }
                                                                                        switch v112 {
                                                                                        case 0:
                                                                                            if (-2147483648 > 584) {
                                                                                                fmt.Println(v111)
                                                                                            }
                                                                                        }
                                                                                        if true {
                                                                                            var v147 int8 = 100
                                                                                            _ = v147
                                                                                        }
                                                                                        v148 := []int{fi47, 0, (-(fi49))}
                                                                                        _ = v148
                                                                                    }
                                                                                    var v149 int32 = -2147483648
                                                                                    var v150 uint64 = v50
                                                                                    _ = v149
                                                                                    _ = v150
                                                                                }
                                                                            } else {
                                                                                v151 := []int{}
                                                                                if (2 <= 696) {
                                                                                    switch 274 {
                                                                                    case 0:
                                                                                        var v152 bool = (true || v33)
                                                                                        _ = v152
                                                                                    case 32768:
                                                                                        v111 = "\n\t"
                                                                                    case 128:
                                                                                        fmt.Println(v111, v2)
                                                                                    default:
                                                                                        v153 := GenS18{F0: 42}
                                                                                        _ = v153
                                                                                    }
                                                                                    if len(v151) >= 2 {
                                                                                        v151 = v151[1:len(v151)]
                                                                                    }
                                                                                    v154 := GenS19{F0: int16(v1), F1: 0}
                                                                                    v155 := GenS5{F0: v50}
                                                                                    var v156 uint32 = 2
                                                                                    _ = v154
                                                                                    _ = v155
                                                                                    _ = v156
                                                                                } else {
                                                                                    var v157 bool = (51 <= 65535)
                                                                                    var v158 int64 = -1
                                                                                    _ = v157
                                                                                    _ = v158
                                                                                }
                                                                                v112 = fi49
                                                                                v159 := v4[99]
                                                                                _ = v151
                                                                                _ = v159
                                                                            }
                                                                            v160 := GenS2{F0: (v50 % ((v50) | 1))}
                                                                            _ = v112
                                                                            _ = v160
                                                                        default:
                                                                            delete(v3, v2)
                                                                        }
                                                                        var v161 bool = (v33 && (v1 < (v1 + v1)))
                                                                        _ = v111
                                                                        _ = v161
                                                                    } else {
                                                                        fmt.Println(v1)
                                                                        fmt.Println(fi47, v50)
                                                                    }
                                                                }
                                                                var v162 int8 = 127
                                                                fmt.Println(v2, fi49, fi49)
                                                                _ = v109
                                                                _ = v162
                                                            }
                                                        }
                                                    } else {
                                                        v2 = ("\n\t" + v2)
                                                    }
                                                    delete(v4, 0)
                                                    v163 := GenS20{F0: 313, F1: 65535}
                                                    _ = v163
                                                default:
                                                    if v33 {
                                                        var v164 uint32 = 2147483648
                                                        _ = v164
                                                    }
                                                }
                                            case -9223372036854775808:
                                                v3[(((v2 + "a") + v2) + "foo bar")] = true
                                            }
                                        } else {
                                            var v165 uint64 = v50
                                            if (!(!("\n\t" != "a"))) {
                                                var v166 int8 = 127
                                                if v33 {
                                                    if (!(0 >= 1000)) {
                                                        var v167 uint64 = (v165 / ((v165) | 1))
                                                        _ = v167
                                                    } else {
                                                        switch "a" {
                                                        case "\n\t":
                                                            v166 = v166
                                                        case "hello":
                                                            fmt.Println(fi49)
                                                        }
                                                        v4[2] = "foo bar"
                                                        if (0.0 >= 0.1) {
                                                            var v168 int8 = (v166 % (((^(v166))) | 1))
                                                            fmt.Println(fi49, v168)
                                                            var v169 int = -372
                                                            var v170 bool = false
                                                            _ = v168
                                                            _ = v169
                                                            _ = v170
                                                        }
                                                    }
                                                } else {
                                                    var v171 uint64 = 18446744073709551615
                                                    _ = v171
                                                }
                                                v166 = (v166 % ((v166) | 1))
                                                var v172 uint32 = 2147483648
                                                v173 := GenS21{F0: true, F1: int(v1), F2: 255}
                                                _ = v166
                                                _ = v172
                                                _ = v173
                                            } else {
                                                v174 := []int{(fi47 & (^(fi47)))}
                                                v50 -= 2
                                                fmt.Println(v33, v50)
                                                _ = v174
                                            }
                                            v175 := GenS22{F0: ((v2 + (v2 + v2)) + v2), F1: fi49}
                                            _ = v165
                                            _ = v175
                                        }
                                    }
                                }
                            } else {
                                v50 -= (v50 / ((734) | 1))
                            }
                            _ = v50
                            _ = v51
                        }
                        _ = fi49
                    }
                }
                for fi176 := 0; fi176 < 1; fi176++ {
                    v177 := GenS23{F0: 42, F1: v33}
                    _ = fi176
                    _ = v177
                }
                v178 := []int{fi47, 127, (^(fi47))}
                _ = fi47
                _ = v178
            }
        } else {
            v2 = ""
        }
    }
    if v33 {
        switch v2 {
        case "a":
            var v179 bool = v33
            _ = v179
        case "foo bar":
            v180 := []int{}
            _ = v180
        case "":
            var v181 int16 = 198
            _ = v181
        }
        switch v2 {
        case "":
            v182 := []int{32768, 15}
            _ = v182
        case "hello":
            switch -330 {
            case 32767:
                switch (((v2 + ((v2 + "hello") + "hello")) + (v2 + "hello")) + ("a" + (v2 + "\n\t"))) {
                case "":
                    if true {
                        if false {
                            switch 2147483647 {
                            case -1:
                                var v183 uint = (v1 * v1)
                                _ = v183
                            case 16:
                                var v184 string = (v2 + "hello")
                                _ = v184
                            case 32767:
                                v185 := GenS4{F0: "", F1: (!(1 >= 42)), F2: 0}
                                _ = v185
                            }
                            var v186 uint8 = 186
                            _ = v186
                        } else {
                            var v187 int32 = 2147483647
                            v2 = (("foo bar" + "\n\t") + v2)
                            _ = v187
                        }
                        switch "" {
                        case "foo bar":
                            v188 := GenS24{F0: v2}
                            _ = v188
                        case "\n\t":
                            for fi189 := 0; fi189 < 7; fi189++ {
                                var v190 string = "foo bar"
                                v191 := GenS3{F0: 566}
                                fmt.Println(v190, v2, fi189)
                                _ = fi189
                                _ = v190
                                _ = v191
                            }
                        case "hello":
                            var v192 int8 = 85
                            _ = v192
                        }
                    }
                case "foo bar":
                    for fi193 := 0; fi193 < 20; fi193++ {
                        v33 = (1 <= -1)
                        _ = fi193
                    }
                }
            case 0:
                v1 = (v1 << ((v1) & 63))
            }
        }
    } else {
        v194 := GenS25{F0: 59, F1: -717, F2: 42}
        var v195 int32 = int32(v1)
        var v196 int64 = 1
        _ = v194
        _ = v195
        _ = v196
    }
    v33 = (3.14159 != 0.1)
    if v33 {
        var v197 int16 = -628
        v2 = v2
        _ = v197
    }
    switch -1 {
    case 556:
        v33 = (13 < 52)
    case 2147483647:
        for fi198 := 0; fi198 < 15; fi198++ {
            v33 = (fi198 <= fi198)
            for fi199 := 0; fi199 < 13; fi199++ {
                var v200 string = v2
                v201 := []int{}
                _ = fi199
                _ = v200
                _ = v201
            }
            _ = fi198
        }
    case 363:
        fmt.Println(v2)
    default:
        var v202 uint32 = 2147483647
        _ = v202
    }
    var v203 int16 = -1
    if (9223372036854775807 != 45) {
        if (!(0 >= 2147483648)) {
            for fi204 := 0; fi204 < 3; fi204++ {
                v33 = v33
                _ = fi204
            }
            v205 := GenS17{F0: -52}
            if (213 > 42) {
                if v33 {
                    switch -1 {
                    case 32768:
                        var v206 int64 = -9223372036854775808
                        _ = v206
                    case -100:
                        for fi207 := 0; fi207 < 1; fi207++ {
                            v208 := []int{(fi207 ^ (fi207 - fi207)), (fi207 / ((fi207) | 1)), (^(fi207))}
                            v209 := GenS6{F0: 2}
                            _ = fi207
                            _ = v208
                            _ = v209
                        }
                    case -216:
                        delete(v4, 122)
                    }
                    var v210 float32 = 0.0
                    _ = v210
                } else {
                    var v211 int64 = 1
                    switch int(v203) {
                    case -4:
                        if ((v1 >= v1) || (false && (743 != 930))) {
                            if false {
                                v33 = v33
                                if (!false) {
                                    switch v2 {
                                    case "":
                                        if (!v33) {
                                            v211 = 9223372036854775807
                                            var v212 uint8 = 1
                                            var v213 float32 = 1.5
                                            _ = v212
                                            _ = v213
                                        }
                                    case "foo bar":
                                        if (!v33) {
                                            if (v33 || v33) {
                                                if v33 {
                                                    if v33 {
                                                        if (574 >= int32(v211)) {
                                                            v2 = v2
                                                            if (398 >= 127) {
                                                                v214 := v4[497]
                                                                _ = v214
                                                            } else {
                                                                for fi215 := 0; fi215 < 10; fi215++ {
                                                                    v216 := []int{4294967295, fi215, fi215}
                                                                    var v217 float64 = 0.1
                                                                    if false {
                                                                        if true {
                                                                            fmt.Println(fi215, v1, v33)
                                                                            if ((v33 || (v33 && (v33 || true))) && (3.14159 > float32(float64(0.0)))) {
                                                                                _ = len(v3)
                                                                                switch (fi215 | (fi215 ^ -1)) {
                                                                                case -161:
                                                                                    var v218 uint8 = 223
                                                                                    _ = v218
                                                                                }
                                                                            } else {
                                                                                if (2147483647 <= 2147483648) {
                                                                                    switch v2 {
                                                                                    case "foo bar":
                                                                                        fmt.Println(v1)
                                                                                    case "hello":
                                                                                        var v219 uint8 = 1
                                                                                        _ = v219
                                                                                    default:
                                                                                        v220 := GenS17{F0: 0}
                                                                                        _ = v220
                                                                                    }
                                                                                    var v221 int8 = -128
                                                                                    _ = v221
                                                                                } else {
                                                                                    var v222 int32 = 0
                                                                                    var v223 uint = (v1 ^ 459)
                                                                                    _ = v222
                                                                                    _ = v223
                                                                                }
                                                                                for fi224 := 0; fi224 < 7; fi224++ {
                                                                                    v225 := GenS15{F0: 2147483647}
                                                                                    v33 = (v33 || (v217 > 0.0))
                                                                                    fmt.Println(v211, v211)
                                                                                    var v226 bool = ("foo bar" > v2)
                                                                                    _ = fi224
                                                                                    _ = v225
                                                                                    _ = v226
                                                                                }
                                                                                for fi227 := 0; fi227 < 7; fi227++ {
                                                                                    fmt.Println(v211, v2)
                                                                                    switch "a" {
                                                                                    case "\n\t":
                                                                                        v211 = (v211 >> ((65535) & 63))
                                                                                    case "a":
                                                                                        v33 = (false && v33)
                                                                                    case "":
                                                                                        switch (fi227 % ((fi227) | 1)) {
                                                                                        case -2147483648:
                                                                                            if (("foo bar" <= v2) || ((v203 >= 0) || true)) {
                                                                                                v228 := GenS15{F0: 42}
                                                                                                var v229 int = (fi215 >> ((v1) & 63))
                                                                                                if true {
                                                                                                    var v230 uint64 = 1
                                                                                                    v33 = (uint16(v1) <= 2)
                                                                                                    fmt.Println(v33, v33)
                                                                                                    _ = v230
                                                                                                } else {
                                                                                                    v217 = 0.0
                                                                                                }
                                                                                                _ = v228
                                                                                                _ = v229
                                                                                            } else {
                                                                                                if v33 {
                                                                                                    v2 = v2
                                                                                                    v231 := GenS26{F0: -160, F1: -179, F2: v2}
                                                                                                    _ = v231
                                                                                                }
                                                                                                if false {
                                                                                                    if v33 {
                                                                                                        v203 = 32767
                                                                                                        v2 = ("a" + v2)
                                                                                                        v232 := GenS1{F0: v2, F1: 0, F2: 42}
                                                                                                        var v233 int16 = -32768
                                                                                                        if (!(!true)) {
                                                                                                            if v33 {
                                                                                                                v234 := GenS18{F0: 2}
                                                                                                                if ((!((v211 & 190) != -9223372036854775808)) && v33) {
                                                                                                                    var v235 bool = v33
                                                                                                                    _ = v235
                                                                                                                } else {
                                                                                                                    var v236 uint32 = 663
                                                                                                                    var v237 int32 = 804
                                                                                                                    var v238 float64 = v217
                                                                                                                    _ = v236
                                                                                                                    _ = v237
                                                                                                                    _ = v238
                                                                                                                }
                                                                                                                _ = v234
                                                                                                            } else {
                                                                                                                v239 := GenS14{F0: 128, F1: 65535}
                                                                                                                if len(v216) >= 2 {
                                                                                                                    v216 = v216[1:len(v216)]
                                                                                                                }
                                                                                                                v240 := GenS27{F0: 65535}
                                                                                                                v203 = v203
                                                                                                                _ = v239
                                                                                                                _ = v240
                                                                                                            }
                                                                                                            fmt.Println(v203, fi227, v1)
                                                                                                            if false {
                                                                                                                switch fi215 {
                                                                                                                case 9223372036854775807:
                                                                                                                    if (2 == 85) {
                                                                                                                        v233 -= v203
                                                                                                                        var v241 int32 = -339
                                                                                                                        var v242 uint32 = 209
                                                                                                                        if len(v216) > 0 {
                                                                                                                            v216[0] = fi227
                                                                                                                        }
                                                                                                                        v216 = append(v216, (fi215 / ((fi215) | 1)))
                                                                                                                        _ = len(v216)
                                                                                                                        _ = v241
                                                                                                                        _ = v242
                                                                                                                    }
                                                                                                                case 94:
                                                                                                                    v217 = float64(((0.1 + (1e10 * 3.14159)) / float32(((v217 - (v217 - ((((v217 + float64(float32(v217))) + (v217 / ((v217 / ((float64(0.1) * (0.0 - (1e10 / (1e10 * float64(3.14159))))) * v217)) + v217))) * v217) / (v217 * (float64(-2.25) + v217))))) / 1e10))))
                                                                                                                default:
                                                                                                                    var v243 int16 = 42
                                                                                                                    _ = v243
                                                                                                                }
                                                                                                                var v244 uint16 = 32767
                                                                                                                _ = v244
                                                                                                            }
                                                                                                        }
                                                                                                        v245 := v4[42]
                                                                                                        _ = v232
                                                                                                        _ = v233
                                                                                                        _ = v245
                                                                                                    }
                                                                                                } else {
                                                                                                    switch (v2 + ("a" + (v2 + v2))) {
                                                                                                    case "":
                                                                                                        if (85 >= 0) {
                                                                                                            var v246 uint64 = 2
                                                                                                            _ = v246
                                                                                                        }
                                                                                                    case "\n\t":
                                                                                                        if true {
                                                                                                            fmt.Println(v33, v2, fi227)
                                                                                                            if (v217 > v217) {
                                                                                                                if ((^(v1)) <= (^(v1))) {
                                                                                                                    v247 := v4[2147483647]
                                                                                                                    _ = v247
                                                                                                                }
                                                                                                                var v248 uint32 = 2147483647
                                                                                                                _ = v248
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                if len(v216) > 0 {
                                                                                                    v216[0] = fi227
                                                                                                }
                                                                                            }
                                                                                        case 0:
                                                                                            v249 := GenS15{F0: 2147483648}
                                                                                            _ = v249
                                                                                        case 42:
                                                                                            var v250 string = v2
                                                                                            _ = v250
                                                                                        default:
                                                                                            delete(v3, v2)
                                                                                            if true {
                                                                                                delete(v4, 613)
                                                                                                switch v2 {
                                                                                                case "":
                                                                                                    v2 = "foo bar"
                                                                                                case "\n\t":
                                                                                                    var v251 int16 = (^(v203))
                                                                                                    _ = v251
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        v252 := GenS28{F0: 4294967295, F1: (v1 >> ((v1) & 63))}
                                                                                        _ = v252
                                                                                    }
                                                                                    if v33 {
                                                                                        v2 = ""
                                                                                        v203 = (v203 * (v203 ^ 42))
                                                                                    } else {
                                                                                        v33 = ((42 > 42) || ((v2 == ("a" + "\n\t")) && true))
                                                                                        var v253 string = v2
                                                                                        var v254 float64 = float64(-2.25)
                                                                                        fmt.Println(fi215, fi215)
                                                                                        var v255 int8 = 42
                                                                                        v253 = v253
                                                                                        _ = len(v4)
                                                                                        _ = v253
                                                                                        _ = v254
                                                                                        _ = v255
                                                                                    }
                                                                                    _ = fi227
                                                                                }
                                                                            }
                                                                            if (!(3.14159 == 0.0)) {
                                                                                if true {
                                                                                    for fi256 := 0; fi256 < 8; fi256++ {
                                                                                        _ = len(v216)
                                                                                        if v33 {
                                                                                            fmt.Println(v211, fi256)
                                                                                            v1 += 560
                                                                                        }
                                                                                        _ = fi256
                                                                                    }
                                                                                }
                                                                                fmt.Println(fi215)
                                                                            } else {
                                                                                var v257 uint32 = 1
                                                                                var v258 uint = 42
                                                                                _ = v257
                                                                                _ = v258
                                                                            }
                                                                        } else {
                                                                            v259 := GenS27{F0: 0}
                                                                            fmt.Println(v33, v217)
                                                                            _ = v259
                                                                        }
                                                                        var v260 string = ("a" + v2)
                                                                        var v261 uint16 = 32767
                                                                        _ = v260
                                                                        _ = v261
                                                                    } else {
                                                                        v262 := GenS22{F0: ("a" + "hello"), F1: 2147483647}
                                                                        _ = v262
                                                                    }
                                                                    _ = fi215
                                                                    _ = v216
                                                                    _ = v217
                                                                }
                                                            }
                                                        }
                                                        v263 := []int{-110}
                                                        _ = v263
                                                    } else {
                                                        var v264 uint32 = 2
                                                        _ = v264
                                                    }
                                                } else {
                                                    if v33 {
                                                        var v265 int16 = -740
                                                        v3["foo bar"] = false
                                                        _ = v265
                                                    } else {
                                                        switch 128 {
                                                        case -250:
                                                            v2 = (((v2 + "foo bar") + (v2 + v2)) + ("a" + ""))
                                                        case 412:
                                                            if ((0.1 / 1.0) > ((1e10 / -2.25) / (3.14159 - ((0.1 * 3.14159) / -2.25)))) {
                                                                var v266 uint8 = 128
                                                                delete(v4, 4294967295)
                                                                for fi267 := 0; fi267 < 1; fi267++ {
                                                                    v211 = (v211 >> (((v1 * 42)) & 63))
                                                                    _ = fi267
                                                                }
                                                                _ = v266
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(v33, v33, v33)
                                                    fmt.Println(v1)
                                                }
                                                if v33 {
                                                    v1 -= (^(v1))
                                                }
                                            }
                                            fmt.Println(v33)
                                        } else {
                                            for fi268 := 0; fi268 < 19; fi268++ {
                                                var v269 float64 = 0.0
                                                if v33 {
                                                    _ = len(v3)
                                                }
                                                v270 := GenS29{F0: 65535}
                                                _ = fi268
                                                _ = v269
                                                _ = v270
                                            }
                                            var v271 int32 = -492
                                            fmt.Println(v271)
                                            _ = v271
                                        }
                                    case "hello":
                                        var v272 int64 = 0
                                        _ = v272
                                    }
                                    v273 := []int{-741, 42, -9223372036854775808}
                                    _ = v273
                                }
                            } else {
                                v33 = false
                                v4[2147483647] = "foo bar"
                                v274 := GenS17{F0: 127}
                                _ = v274
                            }
                            if v33 {
                                v203 -= 0
                                var v275 int = -419
                                _ = v275
                            }
                            v3[v2] = v33
                        } else {
                            if v33 {
                                fmt.Println(v211, v211, v1)
                                var v276 string = (v2 + (v2 + ((v2 + ((("\n\t" + v2) + v2) + v2)) + v2)))
                                _ = v276
                            }
                            v203 = (v203 >> ((v1) & 63))
                            if (0 != 42) {
                                fmt.Println(v203)
                                if true {
                                    v4[42] = (v2 + v2)
                                } else {
                                    fmt.Println(v33, v33)
                                }
                                fmt.Println(v211, v1)
                            }
                        }
                    case 793:
                        v33 = true
                    case 32767:
                        for fi277 := 0; fi277 < 9; fi277++ {
                            fmt.Println(v211, v211)
                            if true {
                                var v278 uint16 = 255
                                v211 = 42
                                _ = v278
                            }
                            switch v2 {
                            case "":
                                _ = len(v3)
                            default:
                                v279 := GenS24{F0: (v2 + "\n\t")}
                                _ = v279
                            }
                            _ = fi277
                        }
                    default:
                        var v280 string = v2
                        _ = v280
                    }
                    _ = v211
                }
                delete(v4, 2147483647)
            } else {
                for fi281 := 0; fi281 < 2; fi281++ {
                    var v282 uint8 = 117
                    fmt.Println(v2, v203, v203)
                    v2 = (v2 + v2)
                    _ = fi281
                    _ = v282
                }
                v33 = v33
            }
            _ = v205
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v203:", v203)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v203
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v40
}
