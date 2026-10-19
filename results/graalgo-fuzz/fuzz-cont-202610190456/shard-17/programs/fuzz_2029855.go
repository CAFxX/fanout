package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS5 struct { F0 int; F1 int16; F2 uint }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 uint8; F1 int64; F2 string }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS10 struct { F0 string; F1 int32 }
type GenS11 struct { F0 int32; F1 int32 }
type GenS12 struct { F0 uint16; F1 int32 }
type GenS13 struct { F0 string; F1 int64; F2 uint }
type GenS14 struct { F0 int32; F1 uint16; F2 uint }
type GenS15 struct { F0 uint64; F1 int; F2 string }
type GenS16 struct { F0 bool; F1 uint8 }
type GenS17 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS18 struct { F0 int32; F1 bool }
type GenS19 struct { F0 int64; F1 int8; F2 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 313
    v3 := []uint32{}
    var v4 bool = false
    var v5 int32 = 687
    v4 = v4
    for fi6 := 0; fi6 < 17; fi6++ {
        v5 = (v5 - 1)
        _ = fi6
    }
    fmt.Println(v4, v4)
    switch "foo bar" {
    case "\n\t":
        for fi7 := 0; fi7 < 8; fi7++ {
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 int = fi7
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    case "hello":
        switch 1 {
        case 406:
            v10 := GenS1{F0: 154, F1: 718, F2: -2147483648}
            _ = v10
        }
    default:
        var v11 uint32 = 42
        _ = v11
    }
    switch "hello" {
    case "hello":
        fmt.Println(v2, v2, v2)
    case "":
        if v4 {
            fmt.Println(v2)
            switch "a" {
            case "a":
                delete(v1, "a")
            }
            switch "a" {
            case "":
                v2 -= 42
            default:
                for fi12 := 0; fi12 < 14; fi12++ {
                    v13 := v1["\n\t"]
                    _ = fi12
                    _ = v13
                }
            }
        }
    case "a":
        if false {
            v4 = (!v4)
            var v14 uint32 = 2147483647
            _ = v14
        } else {
            var v15 uint64 = 2
            _ = v15
        }
    }
    if (false || v4) {
        for fi16 := 0; fi16 < 3; fi16++ {
            v5 = (-(v2))
            v5 += 0
            _ = fi16
        }
        if (696 != 928) {
            var v17 int8 = 25
            if (!(1 == 0)) {
                if true {
                    if v4 {
                        var v18 uint8 = 42
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 4; fi19++ {
                        v2 = v2
                        _ = fi19
                    }
                }
            }
            _ = v17
        } else {
            v2 = v2
        }
        for fi20 := 0; fi20 < 1; fi20++ {
            if (0 >= 42) {
                v21 := GenS1{F0: 255, F1: 0, F2: v2}
                switch (fi20 ^ 32767) {
                case -998:
                    switch fi20 {
                    case -41:
                        for fi22 := 0; fi22 < 8; fi22++ {
                            v23 := GenS1{F0: 88, F1: 245, F2: 2147483647}
                            _ = fi22
                            _ = v23
                        }
                    default:
                        for fi24 := 0; fi24 < 13; fi24++ {
                            v5 += (v5 ^ v2)
                            v25 := GenS2{F0: (fi20 != (-(fi24)))}
                            _ = fi24
                            _ = v25
                        }
                    }
                default:
                    if v4 {
                        switch (fi20 | fi20) {
                        case 32767:
                            v5 = (v5 * (-(v2)))
                        case 128:
                            var v26 uint = 42
                            _ = v26
                        case -268:
                            if (1 != 556) {
                                for fi27 := 0; fi27 < 12; fi27++ {
                                    switch (("\n\t" + "") + "a") {
                                    case "\n\t":
                                        fmt.Println(v4)
                                    case "a":
                                        v4 = true
                                    case "hello":
                                        v28 := GenS1{F0: 42, F1: 2147483647, F2: 1}
                                        var v29 int = (fi27 / ((fi20) | 1))
                                        _ = v28
                                        _ = v29
                                    default:
                                        switch fi27 {
                                        case 4294967295:
                                            if (!((!v4) && v4)) {
                                                var v30 int = -2147483648
                                                _ = len(v3)
                                                if true {
                                                    v31 := GenS1{F0: 104, F1: 4294967295, F2: (v2 ^ (v2 - 2147483647))}
                                                    var v32 float32 = float32(0.1)
                                                    _ = v31
                                                    _ = v32
                                                } else {
                                                    var v33 int = fi27
                                                    _ = v33
                                                }
                                                v4 = v4
                                                if len(v3) > 0 {
                                                    v3[0] = 787
                                                }
                                                _ = v30
                                            } else {
                                                var v34 int16 = 1
                                                _ = v34
                                            }
                                        case 56:
                                            v35 := GenS2{F0: v4}
                                            _ = v35
                                        case 42:
                                            v4 = false
                                        default:
                                            switch (fi27 * fi20) {
                                            case 2147483647:
                                                v2 = v5
                                            case 1:
                                                var v36 string = ("hello" + ("\n\t" + "foo bar"))
                                                _ = v36
                                            }
                                            var v37 int = 2147483647
                                            _ = v37
                                        }
                                        switch 0 {
                                        case 2147483647:
                                            v5 -= v5
                                        case 42:
                                            switch fi20 {
                                            case 127:
                                                switch (-(fi27)) {
                                                case 569:
                                                    var v38 string = ""
                                                    _ = v38
                                                case 128:
                                                    v39 := GenS3{F0: (fi20 % ((fi20) | 1))}
                                                    _ = v39
                                                case -2147483648:
                                                    v3 = append(v3, 1)
                                                default:
                                                    v5 -= (v5 & v5)
                                                }
                                            }
                                        default:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        switch -2147483648 {
                                        case -395:
                                            v40 := GenS4{F0: 865, F1: 42, F2: 0}
                                            _ = v40
                                        case 4294967295:
                                            var v41 int8 = -121
                                            _ = v41
                                        case 217:
                                            var v42 int16 = 0
                                            _ = v42
                                        default:
                                            var v43 int16 = 1
                                            _ = v43
                                        }
                                        switch "hello" {
                                        case "a":
                                            v5 -= (v5 * (v2 / (((v2 * 42)) | 1)))
                                        case "":
                                            if v4 {
                                                if len(v3) > 0 {
                                                    v3[0] = 837
                                                }
                                                if v4 {
                                                    var v44 uint = 65535
                                                    v44 += 2
                                                    _ = v44
                                                } else {
                                                    if ((865 != 985) && true) {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v5 += v5
                                                        case "foo bar":
                                                            var v45 string = ("\n\t" + "hello")
                                                            _ = v45
                                                        case "":
                                                            var v46 uint = 2
                                                            _ = v46
                                                        }
                                                        v5 = 1
                                                    } else {
                                                        switch "" {
                                                        case "foo bar":
                                                            switch "a" {
                                                            case "a":
                                                                v47 := GenS4{F0: 635, F1: 42, F2: v2}
                                                                _ = v47
                                                            case "\n\t":
                                                                switch ("\n\t" + "\n\t") {
                                                                case "foo bar":
                                                                    v48 := GenS5{F0: fi27, F1: -1, F2: 65535}
                                                                    _ = v48
                                                                case "\n\t":
                                                                    var v49 uint16 = 1
                                                                    _ = v49
                                                                case "":
                                                                    var v50 uint = 4294967295
                                                                    _ = v50
                                                                default:
                                                                    v2 += -86
                                                                }
                                                                if (-636 >= -1) {
                                                                    v51 := GenS5{F0: fi27, F1: -96, F2: 784}
                                                                    var v52 int64 = 1
                                                                    if v4 {
                                                                        v5 = (v2 & v2)
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 2147483647
                                                                        }
                                                                    }
                                                                    _ = v51
                                                                    _ = v52
                                                                } else {
                                                                    var v53 uint = 2
                                                                    var v54 uint16 = 930
                                                                    fmt.Println(v4, fi27)
                                                                    _ = v53
                                                                    _ = v54
                                                                }
                                                            case "":
                                                                if ((!(127 > 1)) || (!(!v4))) {
                                                                    v55 := v1["foo bar"]
                                                                    var v56 uint8 = 128
                                                                    _ = v55
                                                                    _ = v56
                                                                }
                                                            default:
                                                                if true {
                                                                    switch -9223372036854775808 {
                                                                    case -1:
                                                                        var v57 bool = v4
                                                                        _ = v57
                                                                    case -796:
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            if (!(!true)) {
                                                                                v58 := GenS2{F0: true}
                                                                                var v59 uint = 700
                                                                                var v60 bool = ((v59 * v59) == v59)
                                                                                v60 = ((4294967295 > 42) && (-128 != -26))
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = 2
                                                                                }
                                                                                if v60 {
                                                                                    v61 := GenS6{F0: int8(fi27), F1: (0 < 4294967295)}
                                                                                    var v62 int32 = (v2 << (((v59 - (v59 % (((v59 ^ 852)) | 1)))) & 63))
                                                                                    var v63 int16 = 1
                                                                                    _ = v61
                                                                                    _ = v62
                                                                                    _ = v63
                                                                                }
                                                                                var v64 float32 = 1.5
                                                                                _ = v58
                                                                                _ = v59
                                                                                _ = v60
                                                                                _ = v64
                                                                            }
                                                                        case "":
                                                                            var v65 uint64 = 42
                                                                            _ = v65
                                                                        }
                                                                    case 42:
                                                                        var v66 bool = false
                                                                        fmt.Println(fi20, v4)
                                                                        if (-345 != v5) {
                                                                            var v67 float64 = -2.25
                                                                            switch (("" + "a") + "\n\t") {
                                                                            case "\n\t":
                                                                                fmt.Println(fi27, v5, v4)
                                                                            case "foo bar":
                                                                                fmt.Println(v67)
                                                                            }
                                                                            _ = v67
                                                                        }
                                                                        _ = v66
                                                                    }
                                                                    v4 = false
                                                                    v4 = (-32768 > 42)
                                                                } else {
                                                                    v2 = (v2 * (^(v2)))
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                            }
                                                        case "hello":
                                                            if false {
                                                                switch int(v2) {
                                                                case 32767:
                                                                    if ((42 >= 32767) || (v2 != v2)) {
                                                                        v5 = v5
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    }
                                                                    v4 = (1000 < -32768)
                                                                    v4 = v4
                                                                    v4 = v4
                                                                case 2147483647:
                                                                    var v68 uint16 = 65535
                                                                    _ = v68
                                                                case -9223372036854775808:
                                                                    v4 = (!v4)
                                                                default:
                                                                    switch (^(fi27)) {
                                                                    case 1:
                                                                        if ((v2 & int32(fi20)) < 42) {
                                                                            v5 += (v5 % (((v2 ^ (v2 - (v2 ^ (v5 << ((18446744073709551615) & 63)))))) | 1))
                                                                            fmt.Println(v4)
                                                                        }
                                                                    case -594:
                                                                        if v4 {
                                                                            if ((255 == 581) && (("\n\t" + "") <= "hello")) {
                                                                                switch ("\n\t" + ("" + ("hello" + ((((("hello" + ("a" + "a")) + ("\n\t" + "\n\t")) + "hello") + "a") + "hello")))) {
                                                                                case "foo bar":
                                                                                    v69 := GenS4{F0: 733, F1: 2, F2: v2}
                                                                                    v3 = append(v3, 1)
                                                                                    _ = v69
                                                                                case "\n\t":
                                                                                    v1["foo bar"] = (18446744073709551615 < 65535)
                                                                                case "a":
                                                                                    v70 := GenS7{F0: 28, F1: 9223372036854775807, F2: "hello"}
                                                                                    _ = v70
                                                                                default:
                                                                                    if (255 == 0) {
                                                                                        var v71 int16 = -32768
                                                                                        v71 = 32767
                                                                                        _ = v71
                                                                                    } else {
                                                                                        delete(v1, "hello")
                                                                                    }
                                                                                }
                                                                                var v72 uint64 = 4294967295
                                                                                var v73 uint32 = 909
                                                                                _ = v72
                                                                                _ = v73
                                                                            } else {
                                                                                var v74 uint32 = 42
                                                                                var v75 int32 = (v5 * v2)
                                                                                _ = v74
                                                                                _ = v75
                                                                            }
                                                                            if v4 {
                                                                                v2 = v5
                                                                                var v76 bool = false
                                                                                v5 = v2
                                                                                _ = v76
                                                                            } else {
                                                                                switch fi27 {
                                                                                case -1:
                                                                                    fmt.Println(v2, fi20)
                                                                                }
                                                                            }
                                                                        } else {
                                                                            if (1 < 42) {
                                                                                v77 := GenS8{F0: 100}
                                                                                if (0 == 42) {
                                                                                    var v78 string = "\n\t"
                                                                                    v5 = (v5 - (v2 ^ (v2 >> ((65) & 63))))
                                                                                    v5 = (v2 | (v2 >> ((782) & 63)))
                                                                                    _ = v78
                                                                                } else {
                                                                                    v79 := GenS3{F0: (fi27 + (fi27 % (((^(fi20))) | 1)))}
                                                                                    _ = v79
                                                                                }
                                                                                switch "hello" {
                                                                                case "a":
                                                                                    var v80 uint8 = 127
                                                                                    _ = v80
                                                                                default:
                                                                                    if (4294967295 == 42) {
                                                                                        delete(v1, "a")
                                                                                    }
                                                                                }
                                                                                _ = v77
                                                                            } else {
                                                                                v4 = v4
                                                                                fmt.Println(v2, v4)
                                                                                var v81 int8 = 30
                                                                                _ = v81
                                                                            }
                                                                            v4 = false
                                                                            switch ("\n\t" + "\n\t") {
                                                                            case "foo bar":
                                                                                var v82 uint = 0
                                                                                _ = v82
                                                                            case "":
                                                                                var v83 uint = 667
                                                                                _ = v83
                                                                            case "\n\t":
                                                                                if v4 {
                                                                                    v5 += v2
                                                                                    fmt.Println(fi20, fi20, fi20)
                                                                                    var v84 float64 = -2.25
                                                                                    _ = v84
                                                                                } else {
                                                                                    _ = len(v3)
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                if (!(65535 == 42)) {
                                                                    v85 := GenS9{F0: 0, F1: 1, F2: 362}
                                                                    _ = v85
                                                                }
                                                                v86 := GenS2{F0: v4}
                                                                _ = v86
                                                            } else {
                                                                var v87 uint8 = 128
                                                                _ = v87
                                                            }
                                                        default:
                                                            if false {
                                                                v2 += (v2 | v2)
                                                                v4 = (!v4)
                                                                switch (-(fi20)) {
                                                                case 128:
                                                                    v4 = v4
                                                                case 42:
                                                                    v88 := GenS3{F0: (-(fi27))}
                                                                    _ = v88
                                                                case 9223372036854775807:
                                                                    v4 = v4
                                                                default:
                                                                    v4 = v4
                                                                }
                                                            } else {
                                                                v3 = append(v3, 0)
                                                            }
                                                        }
                                                    }
                                                }
                                                var v89 int64 = -9223372036854775808
                                                if false {
                                                    v90 := GenS3{F0: fi20}
                                                    var v91 bool = (!v4)
                                                    _ = v90
                                                    _ = v91
                                                }
                                                if false {
                                                    switch ("a" + "\n\t") {
                                                    case "hello":
                                                        v5 = v5
                                                        var v92 uint32 = 2147483648
                                                        v93 := GenS10{F0: "\n\t", F1: v2}
                                                        _ = v92
                                                        _ = v93
                                                    default:
                                                        delete(v1, "")
                                                    }
                                                    v2 += (v5 | (v5 % (((v2 & v2)) | 1)))
                                                    if true {
                                                        v89 = -604
                                                        if ((("\n\t" + "a") + "\n\t") != "hello") {
                                                            v4 = (!v4)
                                                            var v94 int16 = 464
                                                            _ = v94
                                                        } else {
                                                            if (float64(-2.25) == (1e10 * (-2.25 / (0.0 - 3.14159)))) {
                                                                v2 -= v5
                                                                var v95 uint8 = 150
                                                                _ = v95
                                                            }
                                                            v2 -= (v5 * v5)
                                                            if (!(2 < 486)) {
                                                                _ = len(v1)
                                                                if false {
                                                                    switch (fi27 % (((^(fi27))) | 1)) {
                                                                    case -386:
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            v2 = -788
                                                                        case "hello":
                                                                            fmt.Println(v89)
                                                                        }
                                                                        v96 := GenS3{F0: fi20}
                                                                        v97 := GenS11{F0: v2, F1: v5}
                                                                        _ = v96
                                                                        _ = v97
                                                                    case 127:
                                                                        _ = len(v3)
                                                                    case 247:
                                                                        v5 += (v2 / ((v5) | 1))
                                                                    }
                                                                    if ((348 == 18446744073709551615) || (!v4)) {
                                                                        var v98 uint16 = 2
                                                                        switch 988 {
                                                                        case 689:
                                                                            v3 = append(v3, 2147483647)
                                                                        case -803:
                                                                            switch fi20 {
                                                                            case 4294967295:
                                                                                v5 = v5
                                                                            }
                                                                            v4 = true
                                                                        case 128:
                                                                            if (v4 && v4) {
                                                                                v2 -= int32(fi20)
                                                                                fmt.Println(v2)
                                                                                v89 -= (-(v89))
                                                                            }
                                                                        }
                                                                        switch fi20 {
                                                                        case 0:
                                                                            v89 = (v89 + v89)
                                                                        case 1:
                                                                            v2 += (v2 * v5)
                                                                        case 799:
                                                                            v89 = v89
                                                                        default:
                                                                            var v99 uint = 797
                                                                            _ = v99
                                                                        }
                                                                        _ = v98
                                                                    } else {
                                                                        v100 := GenS9{F0: 4294967295, F1: 556, F2: 1}
                                                                        fmt.Println(v2, v4)
                                                                        _ = v100
                                                                    }
                                                                } else {
                                                                    var v101 uint = 619
                                                                    if false {
                                                                        v101 = v101
                                                                        if ((v2 + 2147483647) < 2147483647) {
                                                                            fmt.Println(fi27, fi20, fi27)
                                                                            v102 := GenS12{F0: 89, F1: (v2 % (((v2 & v2)) | 1))}
                                                                            v3 = append(v3, uint32(v101))
                                                                            _ = v102
                                                                        } else {
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                var v103 uint32 = 4294967295
                                                                                _ = v103
                                                                            default:
                                                                                var v104 int32 = 42
                                                                                _ = v104
                                                                            }
                                                                            var v105 uint64 = 202
                                                                            _ = v105
                                                                        }
                                                                        var v106 uint64 = 42
                                                                        _ = v106
                                                                    } else {
                                                                        if false {
                                                                            var v107 int8 = 0
                                                                            _ = v107
                                                                        }
                                                                    }
                                                                    _ = v101
                                                                }
                                                            } else {
                                                                if v4 {
                                                                    fmt.Println(v4, v2)
                                                                    if (v89 > 42) {
                                                                        switch "a" {
                                                                        case "a":
                                                                            v5 = (v2 << ((1) & 63))
                                                                        case "":
                                                                            _ = len(v3)
                                                                        case "hello":
                                                                            v2 -= -727
                                                                        default:
                                                                            var v108 int32 = -324
                                                                            _ = v108
                                                                        }
                                                                        switch fi27 {
                                                                        case 310:
                                                                            var v109 float32 = 0.1
                                                                            _ = v109
                                                                        case -1:
                                                                            var v110 uint16 = 258
                                                                            _ = v110
                                                                        default:
                                                                            switch fi27 {
                                                                            case 32767:
                                                                                switch ("\n\t" + "\n\t") {
                                                                                case "\n\t":
                                                                                    if v4 {
                                                                                        switch "hello" {
                                                                                        case "\n\t":
                                                                                            v2 += v5
                                                                                        case "foo bar":
                                                                                            v111 := GenS13{F0: ("hello" + "foo bar"), F1: (v89 & v89), F2: 2}
                                                                                            var v112 int16 = 370
                                                                                            _ = v111
                                                                                            _ = v112
                                                                                        case "a":
                                                                                            switch fi27 {
                                                                                            case -131:
                                                                                                v4 = false
                                                                                            case 1:
                                                                                                v4 = (!(!(18446744073709551615 <= 42)))
                                                                                            }
                                                                                        }
                                                                                        switch (fi20 / ((fi27) | 1)) {
                                                                                        case 42:
                                                                                            v113 := GenS14{F0: (^(v5)), F1: 0, F2: 18446744073709551615}
                                                                                            _ = v113
                                                                                        case -9223372036854775808:
                                                                                            fmt.Println(v5, v89, fi20)
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        v2 = v5
                                                                                    default:
                                                                                        v114 := GenS15{F0: 343, F1: (fi27 ^ 127), F2: "a"}
                                                                                        _ = v114
                                                                                    }
                                                                                    var v115 int8 = -37
                                                                                    _ = v115
                                                                                }
                                                                            default:
                                                                                v116 := GenS14{F0: 581, F1: 840, F2: 42}
                                                                                _ = v116
                                                                            }
                                                                        }
                                                                    } else {
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            if v4 {
                                                                                v1[("hello" + ("\n\t" + ((("\n\t" + "a") + "\n\t") + ("foo bar" + "foo bar"))))] = v4
                                                                                var v117 bool = v4
                                                                                _ = len(v1)
                                                                                var v118 bool = (!v4)
                                                                                _ = v117
                                                                                _ = v118
                                                                            }
                                                                        case "a":
                                                                            v4 = false
                                                                        }
                                                                        v119 := GenS16{F0: false, F1: 42}
                                                                        v120 := GenS6{F0: -33, F1: v4}
                                                                        var v121 uint8 = 2
                                                                        v2 -= v2
                                                                        _ = v119
                                                                        _ = v120
                                                                        _ = v121
                                                                    }
                                                                }
                                                            }
                                                            if v4 {
                                                                v4 = (0.0 >= 0.0)
                                                                v2 = v2
                                                                var v122 uint32 = 0
                                                                v2 = -1
                                                                _ = v122
                                                            } else {
                                                                var v123 uint16 = 255
                                                                _ = v123
                                                            }
                                                            switch (fi20 ^ (fi20 + (fi20 % ((-961) | 1)))) {
                                                            case 2147483647:
                                                                if true {
                                                                    switch (fi27 & fi27) {
                                                                    case -188:
                                                                        var v124 uint16 = 0
                                                                        _ = v124
                                                                    case -528:
                                                                        var v125 int8 = 67
                                                                        _ = v125
                                                                    case -210:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 0
                                                                        }
                                                                    default:
                                                                        var v126 uint8 = 128
                                                                        _ = v126
                                                                    }
                                                                    fmt.Println(v89)
                                                                    var v127 int16 = -32768
                                                                    _ = v127
                                                                }
                                                            default:
                                                                var v128 uint32 = 2147483647
                                                                _ = v128
                                                            }
                                                        }
                                                        v2 -= 1000
                                                    }
                                                    v129 := GenS7{F0: 255, F1: (-(v89)), F2: ""}
                                                    switch fi20 {
                                                    case -351:
                                                        _ = len(v1)
                                                    case -467:
                                                        v5 = v2
                                                    case 32767:
                                                        v89 += (v89 - 42)
                                                    default:
                                                        var v130 uint = 0
                                                        _ = v130
                                                    }
                                                    if v4 {
                                                        v4 = (0 >= 1)
                                                        var v131 float32 = -2.25
                                                        _ = v131
                                                    } else {
                                                        v132 := GenS14{F0: (-(v2)), F1: 65535, F2: 637}
                                                        fmt.Println(fi20, fi27)
                                                        _ = v132
                                                    }
                                                    _ = v129
                                                } else {
                                                    delete(v1, ((("hello" + ("" + ("foo bar" + ("\n\t" + (("foo bar" + ("a" + "a")) + ((((("foo bar" + "\n\t") + "foo bar") + "hello") + (("\n\t" + "a") + ("\n\t" + ((("\n\t" + "") + "\n\t") + (((("foo bar" + "\n\t") + "\n\t") + "") + "a"))))) + (("hello" + "a") + "a"))))))) + ("hello" + (("\n\t" + "a") + ("" + ((("hello" + "hello") + "\n\t") + "\n\t"))))) + ""))
                                                }
                                                v89 = v89
                                                if false {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                } else {
                                                    switch fi27 {
                                                    case 1:
                                                        if v4 {
                                                            switch fi20 {
                                                            case 9223372036854775807:
                                                                if v4 {
                                                                    v4 = (v4 && v4)
                                                                    var v133 uint16 = 65535
                                                                    var v134 uint64 = 81
                                                                    v134 += v134
                                                                    _ = v133
                                                                    _ = v134
                                                                } else {
                                                                    if v4 {
                                                                        if true {
                                                                            v135 := GenS17{F0: 88, F1: (v89 - (v89 >> ((18446744073709551615) & 63))), F2: 2147483647}
                                                                            switch 4294967295 {
                                                                            case 2147483647:
                                                                                fmt.Println(v4)
                                                                            case -100:
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    var v136 uint = 4294967295
                                                                                    _ = v136
                                                                                case "":
                                                                                    var v137 int32 = 1
                                                                                    _ = v137
                                                                                default:
                                                                                    _ = len(v3)
                                                                                }
                                                                            case 302:
                                                                                var v138 bool = v4
                                                                                _ = v138
                                                                            }
                                                                            _ = v135
                                                                        } else {
                                                                            switch (fi27 & 1) {
                                                                            case 0:
                                                                                switch 32768 {
                                                                                case 458:
                                                                                    v139 := GenS10{F0: "a", F1: v2}
                                                                                    _ = v139
                                                                                case 42:
                                                                                    fmt.Println(fi20)
                                                                                default:
                                                                                    switch (fi27 << ((4294967295) & 63)) {
                                                                                    case 9223372036854775807:
                                                                                        v140 := GenS18{F0: 0, F1: (!false)}
                                                                                        _ = v140
                                                                                    }
                                                                                }
                                                                            default:
                                                                                var v141 string = "\n\t"
                                                                                _ = v141
                                                                            }
                                                                            var v142 string = "foo bar"
                                                                            _ = v142
                                                                        }
                                                                        if v4 {
                                                                            var v143 bool = (32767 == -604)
                                                                            _ = v143
                                                                        } else {
                                                                            var v144 float32 = (3.14159 - 1e10)
                                                                            if (!false) {
                                                                                fmt.Println(v89, fi20, fi27)
                                                                            } else {
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    v5 = (v2 % ((1) | 1))
                                                                                case "a":
                                                                                    v145 := GenS10{F0: "\n\t", F1: (v2 + (v2 % (((v2 % (((v5 | v2)) | 1))) | 1)))}
                                                                                    _ = v145
                                                                                case "":
                                                                                    fmt.Println(v144, v4)
                                                                                default:
                                                                                    if v4 {
                                                                                        _ = len(v1)
                                                                                        var v146 int = (fi27 * (fi20 / ((326) | 1)))
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = 2147483648
                                                                                        }
                                                                                        _ = v146
                                                                                    }
                                                                                }
                                                                                v89 += v89
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            }
                                                                            _ = v144
                                                                        }
                                                                        var v147 int = (-(fi20))
                                                                        _ = v147
                                                                    }
                                                                }
                                                            case 32767:
                                                                fmt.Println(v5, fi20)
                                                            default:
                                                                v89 += 324
                                                            }
                                                            v2 -= 2147483647
                                                        } else {
                                                            v1[("\n\t" + ("" + "foo bar"))] = v4
                                                            var v148 string = ("a" + ("a" + "\n\t"))
                                                            var v149 uint16 = 255
                                                            _ = v148
                                                            _ = v149
                                                        }
                                                    case 32768:
                                                        _ = len(v1)
                                                    }
                                                }
                                                if (((!((v2 - v2) == v5)) && (267 <= 1)) || v4) {
                                                    var v150 int = (^(fi20))
                                                    _ = v150
                                                } else {
                                                    var v151 uint32 = 1
                                                    var v152 uint64 = 469
                                                    _ = v151
                                                    _ = v152
                                                }
                                                _ = v89
                                            }
                                        case "\n\t":
                                            fmt.Println(fi20, v4)
                                        default:
                                            v4 = (!(!(!v4)))
                                            switch 381 {
                                            case 128:
                                                var v153 int = (^(fi20))
                                                _ = v153
                                            default:
                                                switch "foo bar" {
                                                case "":
                                                    switch 0 {
                                                    case 0:
                                                        fmt.Println(fi27, v5, fi20)
                                                    default:
                                                        var v154 int = (-(fi27))
                                                        _ = v154
                                                    }
                                                }
                                            }
                                            switch (^(fi27)) {
                                            case -464:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            case 0:
                                                v5 = (-(v2))
                                            case -1:
                                                var v155 uint64 = 178
                                                _ = v155
                                            }
                                            v156 := GenS19{F0: 9223372036854775807, F1: 49, F2: 1}
                                            _ = v156
                                        }
                                    }
                                    v157 := GenS9{F0: 2, F1: 18446744073709551615, F2: 1}
                                    _ = fi27
                                    _ = v157
                                }
                                var v158 uint8 = 2
                                _ = v158
                            } else {
                                delete(v1, "a")
                            }
                        }
                        for fi159 := 0; fi159 < 6; fi159++ {
                            switch "hello" {
                            case "foo bar":
                                var v160 int16 = -206
                                _ = v160
                            case "a":
                                v161 := v1["a"]
                                _ = v161
                            default:
                                if (!(!v4)) {
                                    var v162 bool = (!v4)
                                    v163 := GenS1{F0: 2, F1: 127, F2: (v5 - (v2 + v2))}
                                    _ = v162
                                    _ = v163
                                } else {
                                    switch "" {
                                    case "a":
                                        var v164 uint64 = 42
                                        _ = v164
                                    }
                                }
                            }
                            var v165 bool = v4
                            var v166 bool = (!(false && v165))
                            _ = fi159
                            _ = v165
                            _ = v166
                        }
                    }
                }
                var v167 bool = false
                _ = v167
                _ = v21
            }
            _ = fi20
        }
    }
    _ = len(v3)
    fmt.Println(v4, v4, v4)
    for fi168 := 0; fi168 < 5; fi168++ {
        var v169 int32 = (v5 * v5)
        _ = fi168
        _ = v169
    }
    if v4 {
        for fi170 := 0; fi170 < 18; fi170++ {
            if ((!(!(!(true && v4)))) && ((v5 & (v5 / (((v5 ^ (v5 * (-(v2))))) | 1))) > v2)) {
                v171 := GenS4{F0: 272, F1: 65535, F2: -2147483648}
                var v172 uint32 = 0
                _ = len(v3)
                _ = v171
                _ = v172
            }
            v2 -= -544
            v3 = append(v3, 65535)
            _ = fi170
        }
        var v173 int16 = 909
        _ = v173
    }
    var v174 uint16 = 773
    delete(v1, "\n\t")
    v174 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v174:", v174)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v174
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
