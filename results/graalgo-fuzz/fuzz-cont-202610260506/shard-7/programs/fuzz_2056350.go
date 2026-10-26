package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 uint64; F1 string; F2 int }
type GenS7 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS8 struct { F0 bool }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 0.0
    var v3 bool = true
    v4 := make(map[uint32]bool)
    v5 := GenS1{F0: 128}
    v6 := GenS1{F0: -52}
    v2 = 0.1
    for fi7 := 0; fi7 < 1; fi7++ {
        switch fi7 {
        case 31:
            var v8 uint32 = 65535
            _ = v8
        default:
            switch fi7 {
            case 697:
                v9 := GenS1{F0: fi7}
                _ = v9
            case 4294967295:
                switch -2147483648 {
                case 9223372036854775807:
                    v2 = (3.14159 * 1e10)
                case 128:
                    delete(v1, fi7)
                case 1:
                    if v3 {
                        var v10 int = fi7
                        _ = v10
                    }
                }
            case -419:
                for fi11 := 0; fi11 < 6; fi11++ {
                    var v12 float64 = v2
                    _ = fi11
                    _ = v12
                }
            }
        }
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 uint16 = 32767
            v2 += v2
            var v15 int32 = 2147483647
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi7
    }
    v3 = v3
    var v16 int8 = -3
    switch -2147483648 {
    case -9223372036854775808:
        if (75 > 42) {
            v17 := GenS2{F0: (v16 / (((v16 & v16)) | 1))}
            if v3 {
                switch 0 {
                case -203:
                    var v18 float64 = 0.1
                    _ = v18
                case 4294967295:
                    v3 = true
                case 961:
                    var v19 int8 = (v16 % ((42) | 1))
                    _ = v19
                default:
                    if v3 {
                        for fi20 := 0; fi20 < 4; fi20++ {
                            var v21 bool = true
                            for fi22 := 0; fi22 < 4; fi22++ {
                                fmt.Println(v3, v2)
                                switch (("\n\t" + "\n\t") + "\n\t") {
                                case "\n\t":
                                    switch fi22 {
                                    case 1:
                                        delete(v4, 2147483647)
                                    default:
                                        var v23 float64 = v2
                                        _ = v23
                                    }
                                case "hello":
                                    fmt.Println(fi22, fi22)
                                default:
                                    switch (fi22 & (^(fi22))) {
                                    case -9223372036854775808:
                                        _ = len(v1)
                                    case -1:
                                        var v24 float32 = (((0.0 - 1e10) * 1.5) + 3.14159)
                                        _ = v24
                                    default:
                                        var v25 uint64 = 1
                                        _ = v25
                                    }
                                }
                                fmt.Println(fi20, fi20)
                                _ = fi22
                            }
                            _ = fi20
                            _ = v21
                        }
                        var v26 float32 = 1e10
                        _ = v26
                    } else {
                        switch "" {
                        case "a":
                            v3 = true
                        case "\n\t":
                            v16 = v16
                        case "foo bar":
                            if v3 {
                                if true {
                                    v3 = (282 < 0)
                                }
                                for fi27 := 0; fi27 < 14; fi27++ {
                                    fmt.Println(v2)
                                    v28 := []int{-220, (fi27 * fi27), (fi27 | fi27), fi27}
                                    _ = fi27
                                    _ = v28
                                }
                                switch 0 {
                                case 9223372036854775807:
                                    var v29 uint16 = 65535
                                    _ = v29
                                default:
                                    var v30 float32 = (float32(((1.5 + 3.14159) + (3.14159 * v2))) + 0.1)
                                    _ = v30
                                }
                            } else {
                                if (!(v3 || (((1e10 - (float32(v2) / 3.14159)) / 0.1) > 1.5))) {
                                    var v31 int32 = 2147483647
                                    _ = v31
                                } else {
                                    var v32 uint16 = 726
                                    for fi33 := 0; fi33 < 13; fi33++ {
                                        v1[(fi33 << ((406) & 63))] = fi33
                                        var v34 uint16 = v32
                                        switch "foo bar" {
                                        case "\n\t":
                                            for fi35 := 0; fi35 < 9; fi35++ {
                                                if true {
                                                    var v36 string = ""
                                                    _ = v36
                                                }
                                                v3 = v3
                                                _ = fi35
                                            }
                                        case "a":
                                            v37 := GenS3{F0: v16, F1: 255}
                                            _ = v37
                                        }
                                        _ = fi33
                                        _ = v34
                                    }
                                    _ = v32
                                }
                                v38 := GenS3{F0: 1, F1: 854}
                                _ = v38
                            }
                        }
                        if v3 {
                            if v3 {
                                v3 = v3
                                switch "\n\t" {
                                case "foo bar":
                                    _ = len(v1)
                                case "a":
                                    v16 += v16
                                case "\n\t":
                                    v39 := []int{-2147483648, -2147483648}
                                    _ = v39
                                }
                            } else {
                                fmt.Println(v3)
                                var v40 int8 = v16
                                _ = v40
                            }
                            var v41 float64 = ((1e10 + (0.1 / float64(1.5))) - v2)
                            switch 423 {
                            case -1:
                                var v42 uint16 = 922
                                _ = v42
                            }
                            _ = v41
                        }
                        v2 = 1.5
                    }
                }
                v4[116] = (!(!v3))
                v2 = 0.0
            } else {
                for fi43 := 0; fi43 < 15; fi43++ {
                    fmt.Println(v3)
                    _ = fi43
                }
                var v44 int = -801
                var v45 int64 = 9223372036854775807
                _ = v44
                _ = v45
            }
            _ = v17
        } else {
            switch "foo bar" {
            case "":
                for fi46 := 0; fi46 < 2; fi46++ {
                    switch (fi46 / ((fi46) | 1)) {
                    case -704:
                        fmt.Println(v2)
                    }
                    if v3 {
                        if true {
                            fmt.Println(v2, v2)
                        } else {
                            v47 := []int{fi46, 4294967295}
                            _ = v47
                        }
                        var v48 uint64 = 1
                        if (0 != -730) {
                            if v3 {
                                for fi49 := 0; fi49 < 4; fi49++ {
                                    v16 = v16
                                    fmt.Println(v2)
                                    v3 = ((v2 - v2) < (v2 * (((float64((1e10 + (-2.25 - (1.5 / (3.14159 * (-2.25 + 3.14159)))))) - v2) + (1e10 + float64(0.0))) / v2)))
                                    _ = fi49
                                }
                            }
                            switch ("\n\t" + "foo bar") {
                            case "":
                                var v50 float64 = (0.1 / v2)
                                _ = v50
                            case "\n\t":
                                v51 := GenS2{F0: v16}
                                _ = v51
                            case "a":
                                if (255 >= 0) {
                                    for fi52 := 0; fi52 < 12; fi52++ {
                                        v2 = 1e10
                                        _ = fi52
                                    }
                                    if v3 {
                                        switch int(v48) {
                                        case 32767:
                                            v3 = (2 > 161)
                                        case -1:
                                            var v53 string = "\n\t"
                                            _ = v53
                                        default:
                                            for fi54 := 0; fi54 < 6; fi54++ {
                                                _ = len(v1)
                                                if true {
                                                    var v55 int16 = 1000
                                                    _ = v55
                                                }
                                                var v56 int32 = -1
                                                fmt.Println(v48, fi54, v16)
                                                var v57 int = (fi46 / ((fi46) | 1))
                                                _ = fi54
                                                _ = v56
                                                _ = v57
                                            }
                                        }
                                        fmt.Println(v3)
                                        v48 = (v48 & v48)
                                    }
                                }
                            default:
                                v16 = 42
                            }
                        }
                        _ = v48
                    }
                    v58 := []int{fi46, -1, 9223372036854775807, (-(fi46))}
                    _ = fi46
                    _ = v58
                }
            case "\n\t":
                v59 := GenS4{F0: 291, F1: v3}
                _ = v59
            default:
                switch "foo bar" {
                case "a":
                    for fi60 := 0; fi60 < 3; fi60++ {
                        v61 := v1[(-(fi60))]
                        _ = fi60
                        _ = v61
                    }
                case "hello":
                    switch "\n\t" {
                    case "hello":
                        for fi62 := 0; fi62 < 16; fi62++ {
                            if (v3 || (42 >= -9223372036854775808)) {
                                switch fi62 {
                                case 128:
                                    for fi63 := 0; fi63 < 20; fi63++ {
                                        fmt.Println(v2, v2)
                                        _ = fi63
                                    }
                                }
                                v64 := GenS4{F0: 4294967295, F1: (10 != 255)}
                                if true {
                                    var v65 uint64 = 4294967295
                                    v2 += 0.0
                                    switch "" {
                                    case "a":
                                        if (-1 <= 1000) {
                                            v3 = (9223372036854775807 != 0)
                                            for fi66 := 0; fi66 < 4; fi66++ {
                                                v67 := GenS5{F0: 42, F1: 1}
                                                switch "foo bar" {
                                                case "":
                                                    v16 += (^(v16))
                                                }
                                                var v68 float32 = (3.14159 / 1.0)
                                                v16 = (v16 * v16)
                                                v69 := GenS3{F0: (v16 / ((v16) | 1)), F1: 2}
                                                var v70 uint8 = 1
                                                fmt.Println(fi66, v70, v16)
                                                var v71 int = 1
                                                var v72 bool = v3
                                                _ = fi66
                                                _ = v67
                                                _ = v68
                                                _ = v69
                                                _ = v70
                                                _ = v71
                                                _ = v72
                                            }
                                        } else {
                                            var v73 int32 = 2147483647
                                            if true {
                                                v74 := GenS4{F0: fi62, F1: (!true)}
                                                _ = len(v4)
                                                for fi75 := 0; fi75 < 5; fi75++ {
                                                    switch (fi62 + fi75) {
                                                    case 981:
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v76 := []int{int(v16)}
                                                                _ = v76
                                                            case "hello":
                                                                v77 := GenS5{F0: 2147483647, F1: -1}
                                                                _ = v77
                                                            case "\n\t":
                                                                v78 := []int{fi75}
                                                                _ = v78
                                                            default:
                                                                switch (-(fi62)) {
                                                                case 127:
                                                                    v73 -= (v73 * (v73 ^ (^(v73))))
                                                                default:
                                                                    var v79 int = (fi62 - (fi62 / (((fi62 ^ 1)) | 1)))
                                                                    _ = v79
                                                                }
                                                            }
                                                        case "a":
                                                            var v80 float64 = v2
                                                            _ = v80
                                                        }
                                                    case -595:
                                                        var v81 bool = v3
                                                        var v82 uint64 = 202
                                                        _ = v81
                                                        _ = v82
                                                    case 500:
                                                        var v83 int32 = (v73 + (v73 | -2147483648))
                                                        _ = v83
                                                    }
                                                    v84 := []int{fi75, (-(fi75)), (-(fi62))}
                                                    v2 += (-2.25 / (float64((0.1 * float32(((0.1 - (0.0 * v2)) + (1.5 * ((float64(0.0) * v2) / v2)))))) / (v2 + (v2 - 3.14159))))
                                                    _ = fi75
                                                    _ = v84
                                                }
                                                _ = v74
                                            }
                                            var v85 int32 = (v73 >> ((65535) & 63))
                                            _ = v73
                                            _ = v85
                                        }
                                    }
                                    _ = v65
                                }
                                _ = v64
                            } else {
                                if v3 {
                                    var v86 float64 = (-2.25 + (float64((1e10 + 3.14159)) + 1.5))
                                    for fi87 := 0; fi87 < 20; fi87++ {
                                        v88 := v1[fi87]
                                        v89 := GenS6{F0: 42, F1: "hello", F2: fi62}
                                        _ = fi87
                                        _ = v88
                                        _ = v89
                                    }
                                    _ = v86
                                }
                                fmt.Println(v16, v3, v16)
                                var v90 bool = (-2.25 != float32(v2))
                                _ = v90
                            }
                            _ = fi62
                        }
                    default:
                        v91 := GenS1{F0: 128}
                        _ = v91
                    }
                case "":
                    v1[4294967295] = 32768
                }
            }
            v92 := GenS7{F0: v3, F1: 2, F2: 335}
            _ = v92
        }
    default:
        if true {
            if true {
                fmt.Println(v3, v2)
                for fi93 := 0; fi93 < 2; fi93++ {
                    v3 = v3
                    switch fi93 {
                    case 2147483647:
                        v2 = v2
                    }
                    _ = fi93
                }
                var v94 int16 = 1000
                _ = v94
            } else {
                var v95 uint16 = 255
                var v96 uint64 = 4294967295
                _ = v95
                _ = v96
            }
            v3 = v3
        } else {
            if (9 != 42) {
                _ = len(v4)
            } else {
                var v97 int32 = 172
                var v98 int = 4294967295
                if (v3 && v3) {
                    for fi99 := 0; fi99 < 3; fi99++ {
                        var v100 int64 = 970
                        _ = fi99
                        _ = v100
                    }
                    v101 := GenS2{F0: 0}
                    for fi102 := 0; fi102 < 3; fi102++ {
                        var v103 uint8 = 128
                        _ = len(v1)
                        _ = fi102
                        _ = v103
                    }
                    _ = v101
                }
                _ = v97
                _ = v98
            }
        }
    }
    v16 = 1
    var v104 int = -264
    v2 = v2
    v105 := []int{(v104 & v104)}
    for fi106 := 0; fi106 < 19; fi106++ {
        v16 -= (v16 / ((-128) | 1))
        _ = fi106
    }
    var v107 uint32 = 0
    v108 := v1[(-(v104))]
    switch 32768 {
    case 32768:
        for fi109 := 0; fi109 < 3; fi109++ {
            v110 := GenS8{F0: ((v108 / ((v104) | 1)) != 128)}
            v108 = (v104 / (((v104 & (v104 * fi109))) | 1))
            var v111 string = ("\n\t" + ("" + "\n\t"))
            _ = fi109
            _ = v110
            _ = v111
        }
    default:
        var v112 int64 = 42
        _ = v112
    }
    switch v104 {
    case 9223372036854775807:
        v113 := GenS2{F0: v16}
        _ = v113
    case 127:
        if false {
            v3 = false
            v114 := GenS5{F0: 2147483648, F1: -9223372036854775808}
            _ = v114
        }
    }
    var v115 bool = (!v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v104:", v104)
    fmt.Println("slicelen v105:", len(v105), v105)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("v115:", v115)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v104
    _ = v105
    _ = v107
    _ = v108
    _ = v115
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
