package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 bool }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 int64; F2 bool }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 uint; F1 uint16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS12 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 int16 = 1000
    var v2 uint32 = 0
    v3 := make(map[string]bool)
    var v4 uint32 = v2
    v4 -= v2
    var v5 int16 = (^(v1))
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(v4, fi6, v2)
        fmt.Println(v1, v2)
        v4 += 42
        _ = fi6
    }
    if (true || false) {
        for fi7 := 0; fi7 < 12; fi7++ {
            delete(v3, "\n\t")
            _ = fi7
        }
    } else {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 bool = (("\n\t" + (("" + "a") + "a")) != "")
            _ = fi8
            _ = v9
        }
        v2 += v4
        switch "hello" {
        case "\n\t":
            v10 := GenS1{F0: 973, F1: -9223372036854775808, F2: (422 == (v5 - (v1 * (^(v5)))))}
            _ = v10
        case "hello":
            var v11 int8 = -47
            _ = v11
        case "a":
            v2 = 4294967295
        default:
            for fi12 := 0; fi12 < 20; fi12++ {
                for fi13 := 0; fi13 < 4; fi13++ {
                    if true {
                        if false {
                            if false {
                                var v14 float32 = (float32(1.5) + 1e10)
                                _ = v14
                            } else {
                                var v15 int8 = 100
                                _ = v15
                            }
                        }
                    } else {
                        switch fi12 {
                        case 42:
                            var v16 float64 = 0.1
                            _ = v16
                        case -734:
                            var v17 int64 = -1
                            _ = v17
                        case -189:
                            switch "a" {
                            case "\n\t":
                                var v18 int16 = -685
                                _ = v18
                            case "":
                                v1 = v5
                            case "foo bar":
                                if true {
                                    v1 = v1
                                } else {
                                    v19 := GenS1{F0: 255, F1: fi12, F2: false}
                                    v20 := v3["foo bar"]
                                    v21 := GenS2{F0: 22, F1: v1}
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                            default:
                                if true {
                                    fmt.Println(v5, v5, v1)
                                    v22 := []int{32767, -9223372036854775808, fi12, -588}
                                    _ = v22
                                } else {
                                    v4 = (^(v2))
                                    _ = len(v3)
                                }
                            }
                        default:
                            if true {
                                var v23 string = "\n\t"
                                if false {
                                    _ = len(v3)
                                    v2 = (v4 | v4)
                                    var v24 int8 = 100
                                    _ = v24
                                }
                                fmt.Println(v1, fi13)
                                _ = len(v3)
                                if false {
                                    var v25 bool = (191 <= 128)
                                    var v26 uint64 = 594
                                    v27 := v3[v23]
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                } else {
                                    fmt.Println(v2, v2, v2)
                                }
                                if true {
                                    if ((-9223372036854775808 <= int64(v1)) || (false || (int32(v1) != -1))) {
                                        var v28 uint = 2
                                        switch (fi13 - (-(fi13))) {
                                        case 898:
                                            switch (fi13 << ((v28) & 63)) {
                                            case 993:
                                                v29 := GenS3{F0: 18446744073709551615}
                                                _ = v29
                                            case 9223372036854775807:
                                                var v30 float32 = 1e10
                                                _ = v30
                                            case 32767:
                                                v31 := GenS4{F0: 1}
                                                _ = v31
                                            }
                                        case 128:
                                            v32 := v3[v23]
                                            _ = v32
                                        }
                                        _ = v28
                                    }
                                }
                                _ = v23
                            } else {
                                var v33 uint64 = 1
                                var v34 string = "foo bar"
                                _ = v33
                                _ = v34
                            }
                        }
                        _ = len(v3)
                    }
                    switch fi13 {
                    case -2147483648:
                        if ((-822 != -9223372036854775808) || (false && true)) {
                            fmt.Println(v2, v2, v4)
                        }
                        v35 := []int{fi13, (fi13 << ((4294967295) & 63))}
                        _ = v35
                    case 2147483647:
                        if true {
                            v4 -= (v2 & 120)
                            v36 := []int{}
                            fmt.Println(v2, v4)
                            _ = v36
                        }
                    case 32768:
                        if false {
                            v5 = (v1 & v1)
                            v37 := GenS2{F0: 18446744073709551615, F1: v1}
                            _ = v37
                        } else {
                            v2 = v2
                            v38 := []int{(fi12 ^ fi13)}
                            v1 += v1
                            _ = v38
                        }
                    default:
                        var v39 uint16 = 0
                        fmt.Println(fi13)
                        v5 = v5
                        _ = v39
                    }
                    _ = fi13
                }
                v4 += v4
                _ = fi12
            }
        }
    }
    v5 = v5
    var v40 uint8 = 1
    for fi41 := 0; fi41 < 14; fi41++ {
        for fi42 := 0; fi42 < 9; fi42++ {
            v43 := v3["a"]
            var v44 uint32 = 0
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = fi41
    }
    v2 = (v4 - (v4 * 2))
    fmt.Println(v5, v4, v5)
    for fi45 := 0; fi45 < 7; fi45++ {
        var v46 uint = 65535
        _ = fi45
        _ = v46
    }
    var v47 int16 = v5
    if false {
        fmt.Println(v40, v40)
        if false {
            fmt.Println(v5, v4)
        }
        if true {
            fmt.Println(v4, v5)
        }
    } else {
        switch "" {
        case "hello":
            for fi48 := 0; fi48 < 4; fi48++ {
                if ((286 >= 65535) || true) {
                    switch "hello" {
                    case "":
                        v49 := []int{2147483647, (fi48 & 42)}
                        _ = v49
                    case "hello":
                        for fi50 := 0; fi50 < 3; fi50++ {
                            _ = len(v3)
                            v5 = (v5 + (v5 * 32767))
                            v4 += (v2 >> ((355) & 63))
                            _ = fi50
                        }
                    default:
                        v51 := []int{}
                        _ = v51
                    }
                } else {
                    v52 := GenS5{F0: 255}
                    fmt.Println(v40, v5)
                    switch (-(fi48)) {
                    case 909:
                        v1 += v47
                    case 32767:
                        for fi53 := 0; fi53 < 12; fi53++ {
                            v40 = (v40 & (^(v40)))
                            fmt.Println(v40, v1, v5)
                            _ = fi53
                        }
                    case -214:
                        var v54 uint64 = 0
                        _ = v54
                    default:
                        for fi55 := 0; fi55 < 3; fi55++ {
                            v4 += (v2 % ((2) | 1))
                            var v56 int16 = v5
                            if true {
                                switch ("a" + "hello") {
                                case "":
                                    var v57 int32 = 42
                                    _ = v57
                                case "hello":
                                    v58 := []int{127, 128, fi55, fi48}
                                    _ = v58
                                case "a":
                                    var v59 int = (fi55 >> ((42) & 63))
                                    _ = v59
                                }
                                v4 = v2
                                var v60 float32 = 3.14159
                                _ = v60
                            }
                            if false {
                                _ = len(v3)
                                fmt.Println(v40, v4)
                                if ((1 <= 4294967295) || (15 < -128)) {
                                    switch ("" + "") {
                                    case "hello":
                                        switch fi48 {
                                        case 128:
                                            v47 = (v1 / ((v56) | 1))
                                        case 1:
                                            var v61 string = ""
                                            _ = v61
                                        case -9223372036854775808:
                                            if (!(818 <= 2)) {
                                                var v62 float32 = float32(float64(3.14159))
                                                v62 -= -2.25
                                                v63 := GenS1{F0: 822, F1: fi48, F2: (!(!true))}
                                                _ = v62
                                                _ = v63
                                            } else {
                                                switch ("" + "foo bar") {
                                                case "":
                                                    fmt.Println(v56, fi55, v1)
                                                case "a":
                                                    switch ("foo bar" + "hello") {
                                                    case "hello":
                                                        if ((v4 != (v4 << ((671) & 63))) && true) {
                                                            if true {
                                                                v64 := v3[""]
                                                                if v64 {
                                                                    if (18446744073709551615 == 0) {
                                                                        var v65 string = "a"
                                                                        var v66 uint32 = v4
                                                                        if false {
                                                                            if ((!(v4 != (v4 & 0))) || ("" < ("foo bar" + (v65 + v65)))) {
                                                                                fmt.Println(v56, v66, v65)
                                                                                var v67 float32 = 3.14159
                                                                                fmt.Println(v66)
                                                                                v40 += (v40 + v40)
                                                                                _ = v67
                                                                            } else {
                                                                                v47 = (v47 + v1)
                                                                                v2 = v4
                                                                                fmt.Println(v40, v1)
                                                                            }
                                                                            fmt.Println(fi55, v66, fi48)
                                                                            v68 := GenS6{F0: (v1 / (((v1 & (-(v5)))) | 1)), F1: -9223372036854775808, F2: false}
                                                                            if v64 {
                                                                                if ((^(v66)) >= 2147483648) {
                                                                                    if (1e10 < -2.25) {
                                                                                        if true {
                                                                                            v65 = v65
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v66, v64, v4)
                                                                                        v69 := GenS7{F0: 0, F1: 1}
                                                                                        _ = v69
                                                                                    }
                                                                                    fmt.Println(v5, v40)
                                                                                    fmt.Println(v5, v65, v56)
                                                                                } else {
                                                                                    var v70 float32 = (1e10 * 0.1)
                                                                                    _ = v70
                                                                                }
                                                                                if ((0.1 * 0.0) > ((((0.1 - ((0.1 + 0.1) / ((0.1 + 0.0) * float64(1.5)))) + 0.1) + 0.0) + (0.0 / 1.5))) {
                                                                                    if (v64 && true) {
                                                                                        var v71 uint8 = 255
                                                                                        var v72 uint = 65535
                                                                                        _ = v71
                                                                                        _ = v72
                                                                                    }
                                                                                    var v73 uint = 4294967295
                                                                                    _ = v73
                                                                                }
                                                                                fmt.Println(v56)
                                                                                var v74 string = (((v65 + v65) + v65) + "foo bar")
                                                                                fmt.Println(fi55)
                                                                                _ = v74
                                                                            } else {
                                                                                delete(v3, v65)
                                                                                if v64 {
                                                                                    v75 := GenS6{F0: (v47 % ((v56) | 1)), F1: -9223372036854775808, F2: v64}
                                                                                    v76 := GenS3{F0: 42}
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                }
                                                                                if (!(416 <= 42)) {
                                                                                    var v77 uint8 = (^(v40))
                                                                                    v78 := GenS8{F0: 390, F1: 7}
                                                                                    _ = v77
                                                                                    _ = v78
                                                                                }
                                                                            }
                                                                            _ = len(v3)
                                                                            fmt.Println(fi55)
                                                                            _ = v68
                                                                        }
                                                                        _ = v65
                                                                        _ = v66
                                                                    } else {
                                                                        var v79 int8 = -128
                                                                        var v80 float64 = 3.14159
                                                                        fmt.Println(v79)
                                                                        _ = v79
                                                                        _ = v80
                                                                    }
                                                                    v81 := []int{fi48, 354, (^(fi55))}
                                                                    _ = v81
                                                                } else {
                                                                    switch (fi48 << ((88) & 63)) {
                                                                    case 0:
                                                                        v82 := GenS5{F0: v40}
                                                                        _ = v82
                                                                    case 127:
                                                                        if (!v64) {
                                                                            v4 = (v4 & (v2 ^ v2))
                                                                            fmt.Println(v4, fi48)
                                                                        } else {
                                                                            v83 := GenS9{F0: (fi48 / ((1) | 1))}
                                                                            _ = v83
                                                                        }
                                                                    default:
                                                                        switch fi55 {
                                                                        case 127:
                                                                            if (913 != 42) {
                                                                                var v84 int = (fi48 & (^(fi48)))
                                                                                fmt.Println(fi48)
                                                                                v3["a"] = (2 < 309)
                                                                                _ = v84
                                                                            } else {
                                                                                var v85 float32 = 1.5
                                                                                _ = v85
                                                                            }
                                                                        case 32768:
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                var v86 int32 = 42
                                                                                _ = v86
                                                                            }
                                                                            v2 += (^(v2))
                                                                            v4 += (v2 | v4)
                                                                        case 9223372036854775807:
                                                                            switch ("foo bar" + "") {
                                                                            case "":
                                                                                if v64 {
                                                                                    v87 := v3[("foo bar" + "\n\t")]
                                                                                    v56 = v5
                                                                                    _ = v87
                                                                                }
                                                                            default:
                                                                                if (42 < -1) {
                                                                                    switch fi55 {
                                                                                    case -255:
                                                                                        v47 -= (v56 & 1)
                                                                                    case -210:
                                                                                        v88 := []int{fi55, fi48}
                                                                                        v5 -= v1
                                                                                        _ = v88
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    switch ("foo bar" + "\n\t") {
                                                                    case "":
                                                                        v64 = ("" >= "a")
                                                                    case "foo bar":
                                                                        v56 += 32767
                                                                    default:
                                                                        if v64 {
                                                                            if v64 {
                                                                                if v64 {
                                                                                    v5 += -1
                                                                                    var v89 bool = (0.0 >= 0.1)
                                                                                    v5 = (v5 >> ((30) & 63))
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        var v90 uint = 42
                                                                                        _ = v90
                                                                                    case "hello":
                                                                                        v1 = 724
                                                                                    }
                                                                                    var v91 uint32 = (v4 | (v2 / ((uint32(v40)) | 1)))
                                                                                    _ = v89
                                                                                    _ = v91
                                                                                } else {
                                                                                    v2 += v4
                                                                                    v5 = int16(v40)
                                                                                }
                                                                                var v92 uint8 = 0
                                                                                v93 := GenS1{F0: 1, F1: fi48, F2: (true || false)}
                                                                                if true {
                                                                                    var v94 uint = uint(v40)
                                                                                    if v64 {
                                                                                        var v95 int64 = 42
                                                                                        _ = v95
                                                                                    }
                                                                                    _ = v94
                                                                                } else {
                                                                                    fmt.Println(v4, v47)
                                                                                    v5 += (-(v5))
                                                                                }
                                                                                fmt.Println(v64)
                                                                                _ = v92
                                                                                _ = v93
                                                                            } else {
                                                                                v56 -= (v56 ^ 0)
                                                                                v47 += (v5 + v1)
                                                                                v96 := []int{32767, fi48}
                                                                                switch fi48 {
                                                                                case 0:
                                                                                    var v97 float64 = 1.5
                                                                                    _ = v97
                                                                                case 459:
                                                                                    v96 = append(v96, fi48)
                                                                                }
                                                                                _ = v96
                                                                            }
                                                                            v4 = 2147483647
                                                                            if false {
                                                                                delete(v3, ("foo bar" + ("foo bar" + ("a" + "a"))))
                                                                                fmt.Println(fi48)
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v56, v1)
                                                                            fmt.Println(v2, v2, v1)
                                                                        }
                                                                    }
                                                                }
                                                                v98 := GenS10{F0: -1}
                                                                _ = v64
                                                                _ = v98
                                                            } else {
                                                                if false {
                                                                    v99 := GenS11{F0: 378, F1: v40, F2: 18446744073709551615}
                                                                    _ = v99
                                                                } else {
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        var v100 string = "foo bar"
                                                                        _ = v100
                                                                    case "a":
                                                                        v47 = v56
                                                                    }
                                                                }
                                                            }
                                                            if true {
                                                                var v101 float64 = (1e10 - 0.0)
                                                                v101 += (v101 * 3.14159)
                                                                _ = v101
                                                            } else {
                                                                v102 := []int{int(v40), fi48, 0}
                                                                switch "\n\t" {
                                                                case "":
                                                                    var v103 uint8 = 127
                                                                    _ = v103
                                                                }
                                                                _ = len(v102)
                                                                var v104 bool = ((4294967295 >= 2) && false)
                                                                v40 = v40
                                                                _ = v102
                                                                _ = v104
                                                            }
                                                            var v105 uint16 = 883
                                                            _ = v105
                                                        } else {
                                                            fmt.Println(v40)
                                                            var v106 uint8 = (v40 - 179)
                                                            _ = v106
                                                        }
                                                    }
                                                default:
                                                    v107 := []int{}
                                                    _ = v107
                                                }
                                                v108 := GenS6{F0: 1000, F1: 389, F2: true}
                                                var v109 uint8 = uint8(v4)
                                                _ = v108
                                                _ = v109
                                            }
                                        }
                                    case "":
                                        v5 = v47
                                    case "foo bar":
                                        if true {
                                            if (1 == -1) {
                                                fmt.Println(fi48, v56)
                                                var v110 uint = 65535
                                                v111 := GenS12{F0: v2, F1: 119}
                                                _ = v110
                                                _ = v111
                                            }
                                            var v112 int32 = 1
                                            _ = v112
                                        }
                                        var v113 int32 = int32(v40)
                                        _ = v113
                                    }
                                    _ = len(v3)
                                    v3[""] = true
                                } else {
                                    var v114 float32 = (((1e10 + (0.1 + (float32((0.0 * 1.5)) + 1.5))) - 0.0) + ((1e10 * 0.1) * (float32(float64(0.0)) * (1e10 / 1.5))))
                                    v5 -= 680
                                    _ = v114
                                }
                            }
                            _ = fi55
                            _ = v56
                        }
                    }
                    _ = v52
                }
                v4 = v2
                var v115 float32 = (1e10 + ((float32((3.14159 * float64(0.1))) * (1.5 * 1.5)) / 1.0))
                _ = fi48
                _ = v115
            }
        case "":
            var v116 string = "foo bar"
            _ = v116
        }
    }
    for fi117 := 0; fi117 < 8; fi117++ {
        fmt.Println(v4, v2)
        v118 := []int{(fi117 ^ (fi117 * (fi117 * 32767)))}
        _ = fi117
        _ = v118
    }
    var v119 uint32 = 42
    fmt.Println(v47)
    if (true && (-128 >= 1)) {
        v40 -= 79
        for fi120 := 0; fi120 < 3; fi120++ {
            v4 = v119
            fmt.Println(v5, v5, v2)
            var v121 int8 = 1
            _ = fi120
            _ = v121
        }
    } else {
        for fi122 := 0; fi122 < 8; fi122++ {
            var v123 float32 = ((0.1 - 1e10) / 1e10)
            _ = fi122
            _ = v123
        }
    }
    var v124 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v119
    _ = v124
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v47
    _ = v5
}
