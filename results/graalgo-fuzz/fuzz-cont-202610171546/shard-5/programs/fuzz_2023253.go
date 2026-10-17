package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string; F1 int16 }
type GenS10 struct { F0 int32; F1 int64 }
type GenS11 struct { F0 int64; F1 bool; F2 uint32 }
type GenS12 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint64, p2 int32) bool {
    fmt.Println(p1, p1)
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    return (true || false)
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[uint32]bool)
    var v3 int64 = 42
    var v4 int = -4
    var v5 uint = 42
    v6 := make(map[string]uint32)
    switch ("hello" + ("a" + ("foo bar" + "\n\t"))) {
    case "\n\t":
        switch "foo bar" {
        case "hello":
            if false {
                switch v4 {
                case 42:
                    v7 := []int{(v4 << (((v5 << ((v5) & 63))) & 63)), (v4 << ((347) & 63))}
                    _ = v7
                case -9223372036854775808:
                    var v8 int16 = 42
                    _ = v8
                }
                v3 = (v3 ^ -9223372036854775808)
            }
        }
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS1{F0: -103}
        var v11 float32 = v1
        v3 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v4, v5, v5)
    var v12 uint = v5
    fmt.Println(v12)
    _ = len(v6)
    var v13 int64 = (v3 ^ v3)
    if (-515 <= -198) {
        for fi14 := 0; fi14 < 10; fi14++ {
            fmt.Println(v1)
            var v15 int = (fi14 * (v4 << ((v5) & 63)))
            _ = fi14
            _ = v15
        }
        switch ((((("foo bar" + "\n\t") + ((((("" + "foo bar") + "a") + ("" + "a")) + ("" + "a")) + "")) + "a") + "foo bar") + "a") {
        case "a":
            switch 2147483647 {
            case 42:
                for fi16 := 0; fi16 < 19; fi16++ {
                    for fi17 := 0; fi17 < 20; fi17++ {
                        v18 := []int{}
                        var v19 int64 = (v13 ^ v13)
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    v12 = 42
                    _ = fi16
                }
            case -9223372036854775808:
                v3 = v3
            }
        case "foo bar":
            fmt.Println(v12, v4, v5)
        case "":
            switch ("a" + ("\n\t" + ("" + "\n\t"))) {
            case "hello":
                v20 := v6[("\n\t" + "hello")]
                _ = v20
            case "\n\t":
                for fi21 := 0; fi21 < 14; fi21++ {
                    switch "a" {
                    case "\n\t":
                        if true {
                            if true {
                                for fi22 := 0; fi22 < 11; fi22++ {
                                    v12 = 4294967295
                                    if true {
                                        switch "a" {
                                        case "a":
                                            if false {
                                                v12 += v5
                                                v5 = (v12 * 973)
                                            }
                                        }
                                    }
                                    _ = fi22
                                }
                                v3 = v13
                                for fi23 := 0; fi23 < 6; fi23++ {
                                    v24 := []int{fi21, (fi21 * v4), (-(fi21)), -960}
                                    var v25 float64 = 0.0
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                }
                            } else {
                                v13 = (v3 - v13)
                                switch "a" {
                                case "hello":
                                    switch fi21 {
                                    case -9223372036854775808:
                                        v26 := GenS1{F0: 92}
                                        _ = v26
                                    }
                                case "a":
                                    var v27 int8 = 0
                                    _ = v27
                                default:
                                    switch (v4 / (((v4 - fi21)) | 1)) {
                                    case 32767:
                                        v12 += (v12 - (v5 + (v12 * v5)))
                                    case 9223372036854775807:
                                        if false {
                                            for fi28 := 0; fi28 < 15; fi28++ {
                                                switch ("a" + "") {
                                                case "a":
                                                    if (float64((v1 * 0.1)) == float64((float32(0.1) / ((float32(0.1) - 3.14159) / v1)))) {
                                                        fmt.Println(v5)
                                                    } else {
                                                        v6["foo bar"] = 2147483648
                                                        if (((42 >= 2) || false) && true) {
                                                            v29 := GenS1{F0: 100}
                                                            if (34 == 128) {
                                                                var v30 float64 = float64((v1 - v1))
                                                                var v31 int = (fi28 >> ((v5) & 63))
                                                                _ = v30
                                                                _ = v31
                                                            }
                                                            switch fi21 {
                                                            case 2147483647:
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v4 += 32767
                                                                }
                                                            case 128:
                                                                v6[(("\n\t" + (("" + "foo bar") + ("foo bar" + "hello"))) + ((((((((("hello" + ("\n\t" + ("foo bar" + (("hello" + "\n\t") + ("foo bar" + ("" + "")))))) + "\n\t") + ("hello" + "hello")) + (("\n\t" + ("" + "hello")) + "hello")) + "a") + "hello") + ("a" + "hello")) + ("\n\t" + (("" + "hello") + ("a" + (("" + "a") + "foo bar"))))) + ("a" + (("hello" + ("a" + ("foo bar" + ("a" + "\n\t")))) + "hello"))))] = 185
                                                            case 42:
                                                                var v32 float64 = float64(float32(float64(0.0)))
                                                                _ = v32
                                                            default:
                                                                var v33 int32 = 0
                                                                v4 -= -9223372036854775808
                                                                _ = v33
                                                            }
                                                            _ = v29
                                                        } else {
                                                            v1 += (v1 / (float32(3.14159) * (1e10 / float32((0.0 + 0.1)))))
                                                            v34 := GenS2{F0: (^(v12)), F1: 255, F2: 4294967295}
                                                            _ = v34
                                                        }
                                                        v35 := []int{}
                                                        _ = v35
                                                    }
                                                case "":
                                                    if (!false) {
                                                        v2[1] = false
                                                        fmt.Println(v1, fi28)
                                                        if ((v5 / (((v5 / ((2) | 1))) | 1)) < (v12 - v12)) {
                                                            if true {
                                                                var v36 uint = 42
                                                                v37 := v2[4294967295]
                                                                if v37 {
                                                                    if (42 != -2147483648) {
                                                                        var v38 int16 = 1000
                                                                        _ = v38
                                                                    } else {
                                                                        v39 := GenS3{F0: ((1 <= -128) || v37)}
                                                                        var v40 float64 = 0.1
                                                                        var v41 uint32 = 791
                                                                        _ = v39
                                                                        _ = v40
                                                                        _ = v41
                                                                    }
                                                                    v37 = ("\n\t" <= "a")
                                                                } else {
                                                                    switch fi28 {
                                                                    case 688:
                                                                        v5 += (v36 >> ((0) & 63))
                                                                    case 2147483647:
                                                                        var v42 uint32 = 42
                                                                        v13 -= (^(v13))
                                                                        _ = v42
                                                                    }
                                                                    v43 := v6["\n\t"]
                                                                    _ = v43
                                                                }
                                                                v44 := GenS1{F0: 100}
                                                                _ = v36
                                                                _ = v37
                                                                _ = v44
                                                            } else {
                                                                if (true && (32767 <= 46)) {
                                                                    v6["\n\t"] = 0
                                                                }
                                                                var v45 bool = (2147483648 == 1)
                                                                _ = v45
                                                            }
                                                            v46 := GenS1{F0: 0}
                                                            var v47 float32 = float32(3.14159)
                                                            v2[42] = (((v1 * 1.5) <= v47) && (!(!false)))
                                                            _ = v46
                                                            _ = v47
                                                        } else {
                                                            v3 = v3
                                                        }
                                                    } else {
                                                        var v48 int32 = 42
                                                        _ = v48
                                                    }
                                                    var v49 int = (v4 >> ((838) & 63))
                                                    _ = v49
                                                default:
                                                    var v50 int32 = -861
                                                    _ = v50
                                                }
                                                _ = len(v6)
                                                v51 := GenS1{F0: 100}
                                                _ = fi28
                                                _ = v51
                                            }
                                            for fi52 := 0; fi52 < 2; fi52++ {
                                                var v53 int8 = 127
                                                _ = fi52
                                                _ = v53
                                            }
                                            v3 += v13
                                        }
                                    default:
                                        var v54 float32 = (-2.25 - (v1 / (float32(1e10) / v1)))
                                        _ = v54
                                    }
                                }
                            }
                            switch v4 {
                            case 407:
                                for fi55 := 0; fi55 < 11; fi55++ {
                                    _ = len(v2)
                                    _ = fi55
                                }
                            case 9223372036854775807:
                                fmt.Println(fi21)
                            case -2147483648:
                                for fi56 := 0; fi56 < 5; fi56++ {
                                    v5 -= 335
                                    _ = fi56
                                }
                            default:
                                v2[2147483647] = (!(127 >= 1000))
                            }
                            if (32767 >= 411) {
                                if true {
                                    v57 := []int{v4}
                                    v12 -= v12
                                    var v58 float64 = (3.14159 * 3.14159)
                                    _ = v57
                                    _ = v58
                                }
                                var v59 float32 = v1
                                _ = v59
                            } else {
                                var v60 uint16 = 2
                                v6["hello"] = 0
                                var v61 uint = (v12 & v12)
                                _ = v60
                                _ = v61
                            }
                        } else {
                            var v62 uint = v5
                            v63 := []int{fi21, 2147483647, (fi21 >> (((v62 + (^(v12)))) & 63)), 128}
                            switch v4 {
                            case 128:
                                if true {
                                    v3 = (v3 | (-(v13)))
                                    v12 -= (^(v62))
                                    v64 := GenS2{F0: (v62 + (v5 / (((v12 % (((^(v5))) | 1))) | 1))), F1: 2, F2: v12}
                                    _ = v64
                                } else {
                                    fmt.Println(v1, v4, v4)
                                    for fi65 := 0; fi65 < 18; fi65++ {
                                        var v66 uint64 = 42
                                        if (((((((v1 + (0.1 - 0.0)) * 1e10) / (float32((0.1 * 1.5)) + v1)) / v1) - v1) * (0.1 - v1)) > v1) {
                                            if false {
                                                switch (fi21 | fi65) {
                                                case 717:
                                                    fmt.Println(v3, v66, v62)
                                                case -715:
                                                    fmt.Println(v66, v3, v1)
                                                case -599:
                                                    var v67 bool = (false && (-315 >= 1))
                                                    _ = v67
                                                }
                                                v4 = (v4 << ((v12) & 63))
                                                var v68 uint = 42
                                                _ = v68
                                            }
                                            switch 0 {
                                            case 42:
                                                v66 -= 48
                                            case -9223372036854775808:
                                                switch "hello" {
                                                case "\n\t":
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        if (127 <= 2) {
                                                            if true {
                                                                fmt.Println(fi65, v13, v66)
                                                                v5 -= (^(v5))
                                                            } else {
                                                                if (v12 == v5) {
                                                                    v69 := GenS4{F0: int(v66)}
                                                                    var v70 int = fi65
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                                switch "hello" {
                                                                case "a":
                                                                    v13 -= v3
                                                                case "\n\t":
                                                                    delete(v2, 42)
                                                                case "":
                                                                    switch (fi65 + -156) {
                                                                    case -878:
                                                                        v71 := v2[2147483647]
                                                                        _ = v71
                                                                    default:
                                                                        var v72 int8 = 1
                                                                        _ = v72
                                                                    }
                                                                default:
                                                                    v73 := GenS3{F0: true}
                                                                    _ = v73
                                                                }
                                                            }
                                                        } else {
                                                            switch fi65 {
                                                            case -296:
                                                                v3 = (v13 >> ((v5) & 63))
                                                            case 2147483647:
                                                                v12 -= 75
                                                                v74 := GenS3{F0: true}
                                                                v63 = append(v63, (fi65 * fi65))
                                                                _ = v74
                                                            case -9223372036854775808:
                                                                if true {
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        fmt.Println(fi21)
                                                                    case "a":
                                                                        if len(v63) >= 2 {
                                                                            v63 = v63[1:len(v63)]
                                                                        }
                                                                    case "":
                                                                        v75 := GenS4{F0: 128}
                                                                        _ = v75
                                                                    }
                                                                    v76 := GenS5{F0: -1, F1: 255}
                                                                    _ = v76
                                                                } else {
                                                                    switch int(v5) {
                                                                    case -174:
                                                                        switch (fi21 * 127) {
                                                                        case -2147483648:
                                                                            var v77 float64 = ((-2.25 - (-2.25 + (3.14159 - float64(3.14159)))) - 1.5)
                                                                            _ = v77
                                                                        }
                                                                        var v78 string = ""
                                                                        v79 := GenS6{F0: 147}
                                                                        _ = v78
                                                                        _ = v79
                                                                    }
                                                                }
                                                            default:
                                                                switch "hello" {
                                                                case "hello":
                                                                    v3 -= (v13 + int64(fi21))
                                                                case "a":
                                                                    fmt.Println(v5, v66)
                                                                default:
                                                                    v62 += 255
                                                                }
                                                            }
                                                            v12 = 2
                                                            var v80 float32 = 1e10
                                                            _ = v80
                                                        }
                                                    case "":
                                                        if true {
                                                            if (!true) {
                                                                switch -135 {
                                                                case 958:
                                                                    v81 := GenS7{F0: 0}
                                                                    _ = v81
                                                                case 2147483647:
                                                                    v82 := GenS5{F0: 312, F1: 1}
                                                                    _ = v82
                                                                case 4294967295:
                                                                    var v83 float64 = 1e10
                                                                    _ = v83
                                                                }
                                                            } else {
                                                                fmt.Println(v13)
                                                            }
                                                        } else {
                                                            var v84 float64 = (3.14159 * float64(v1))
                                                            _ = v84
                                                        }
                                                    default:
                                                        var v85 int16 = 434
                                                        _ = v85
                                                    }
                                                case "foo bar":
                                                    v62 -= v5
                                                    v63 = append(v63, (fi65 % (((fi65 - 9223372036854775807)) | 1)))
                                                    fmt.Println(v13, v3, v66)
                                                }
                                            case 32767:
                                                v86 := GenS8{F0: 2147483647}
                                                _ = v86
                                            }
                                            v4 = v4
                                        } else {
                                            var v87 int8 = -1
                                            _ = v87
                                        }
                                        switch "hello" {
                                        case "foo bar":
                                            switch "" {
                                            case "":
                                                _ = len(v63)
                                            case "foo bar":
                                                if true {
                                                    switch "\n\t" {
                                                    case "":
                                                        if len(v63) >= 2 {
                                                            v63 = v63[1:len(v63)]
                                                        }
                                                        v88 := GenS6{F0: 128}
                                                        if true {
                                                            var v89 float64 = (0.0 / -2.25)
                                                            var v90 int32 = -1
                                                            _ = v89
                                                            _ = v90
                                                        } else {
                                                            v66 = v66
                                                        }
                                                        _ = v88
                                                    case "a":
                                                        if false {
                                                            v13 += 42
                                                            v5 -= (v5 | uint(v66))
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v4, v62, fi21)
                                                    v91 := GenS2{F0: v12, F1: 42, F2: v5}
                                                    fmt.Println(fi21, v5)
                                                    _ = v91
                                                }
                                            default:
                                                v6["a"] = 4294967295
                                            }
                                            v4 = (v4 + fi21)
                                            v6[(("foo bar" + (("\n\t" + "a") + ("a" + "a"))) + "hello")] = 536
                                            if true {
                                                if len(v63) > 0 {
                                                    v63[0] = 388
                                                }
                                                switch fi21 {
                                                case 0:
                                                    v6[("hello" + "hello")] = 243
                                                default:
                                                    v13 -= 42
                                                }
                                                switch "foo bar" {
                                                case "":
                                                    if len(v63) > 0 {
                                                        v63[0] = v4
                                                    }
                                                case "\n\t":
                                                    fmt.Println(v5, v62)
                                                }
                                            }
                                        case "\n\t":
                                            _ = len(v63)
                                        case "":
                                            switch "foo bar" {
                                            case "foo bar":
                                                v92 := GenS4{F0: 523}
                                                _ = v92
                                            case "hello":
                                                var v93 float64 = (float64(v1) - (1.5 - 1.5))
                                                _ = v93
                                            default:
                                                v94 := GenS8{F0: -163}
                                                _ = v94
                                            }
                                        }
                                        v3 -= (v13 << ((42) & 63))
                                        v66 = (v66 >> ((v62) & 63))
                                        _ = fi65
                                        _ = v66
                                    }
                                }
                            case 1:
                                _ = len(v6)
                            case -510:
                                v95 := GenS9{F0: "a", F1: 617}
                                _ = v95
                            }
                            _ = v62
                            _ = v63
                        }
                    default:
                        if true {
                            switch (v4 | fi21) {
                            case 32767:
                                v2[328] = (true && ((1 != 1) || (v5 < 255)))
                            default:
                                for fi96 := 0; fi96 < 11; fi96++ {
                                    fmt.Println(v1)
                                    v97 := v2[618]
                                    v98 := []int{}
                                    _ = fi96
                                    _ = v97
                                    _ = v98
                                }
                            }
                            for fi99 := 0; fi99 < 12; fi99++ {
                                var v100 int16 = -24
                                fmt.Println(v5, v3, fi99)
                                _ = fi99
                                _ = v100
                            }
                        } else {
                            v101 := GenS6{F0: 2}
                            _ = v101
                        }
                    }
                    if false {
                        var v102 float64 = 0.0
                        _ = v102
                    }
                    var v103 int = (fi21 | v4)
                    _ = fi21
                    _ = v103
                }
            default:
                v104 := []int{}
                _ = v104
            }
        }
        v3 = v13
    } else {
        v105 := v6["a"]
        var v106 uint = (v12 >> ((v5) & 63))
        _ = v105
        _ = v106
    }
    v107 := GenS1{F0: 42}
    v12 -= (v5 >> ((v12) & 63))
    v1 = (((((float32(0.0) * (v1 - v1)) - (float32((3.14159 + 1.5)) / (((v1 / (1.5 / v1)) / v1) / 1.0))) - ((1.5 / 0.1) * (((0.1 + 3.14159) - (0.0 / ((-2.25 * (float32(1e10) + ((1.5 / v1) / 1e10))) - (v1 - 3.14159)))) - v1))) * (v1 / -2.25)) * 1e10)
    if ((false || (1.5 <= 1.5)) && false) {
        if false {
            switch (v4 << (((^(v5))) & 63)) {
            case 32768:
                var v108 int16 = 55
                _ = v108
            default:
                fmt.Println(v3, v4)
            }
            switch "hello" {
            case "":
                var v109 int32 = 42
                _ = v109
            case "hello":
                if true {
                    v110 := GenS10{F0: -2147483648, F1: (v3 - (v3 | v13))}
                    _ = v110
                } else {
                    v1 -= v1
                    v111 := v2[637]
                    _ = v111
                }
            default:
                var v112 float32 = -2.25
                _ = v112
            }
            v113 := v2[65535]
            _ = v113
        }
        fmt.Println(v4, v5, v3)
        for fi114 := 0; fi114 < 20; fi114++ {
            var v115 string = "a"
            for fi116 := 0; fi116 < 9; fi116++ {
                v117 := []int{fi114}
                if (!(0 >= 1)) {
                    if true {
                        v3 = (v13 * (v3 * 0))
                        var v118 int = (fi114 | (^(v4)))
                        v117 = append(v117, (v4 | fi114))
                        switch fi116 {
                        case 127:
                            var v119 int = 0
                            _ = v119
                        }
                        _ = v118
                    } else {
                        if true {
                            fmt.Println(v3, v13, v3)
                            switch (v4 >> ((v5) & 63)) {
                            case 32768:
                                switch v115 {
                                case "a":
                                    v120 := GenS11{F0: (v13 >> (((^(v12))) & 63)), F1: (v3 == (-(v3))), F2: 65535}
                                    _ = v120
                                case "foo bar":
                                    v121 := GenS7{F0: -55}
                                    _ = v121
                                case "\n\t":
                                    v12 = v5
                                    var v122 int8 = -128
                                    _ = v122
                                default:
                                    var v123 bool = (true && false)
                                    _ = v123
                                }
                            case 899:
                                var v124 int64 = (v13 - v13)
                                _ = v124
                            default:
                                v125 := GenS10{F0: -1, F1: v3}
                                _ = v125
                            }
                        }
                    }
                    if true {
                        fmt.Println(fi116)
                        var v126 int = 32768
                        _ = v126
                    }
                    v1 = v1
                    v4 -= (fi116 ^ fi116)
                } else {
                    var v127 float32 = 0.0
                    _ = v127
                }
                _ = fi116
                _ = v117
            }
            _ = fi114
            _ = v115
        }
    } else {
        v5 = 2
    }
    for fi128 := 0; fi128 < 19; fi128++ {
        var v129 uint8 = 40
        _ = fi128
        _ = v129
    }
    for fi130 := 0; fi130 < 5; fi130++ {
        v131 := GenS12{F0: 1}
        switch 0 {
        case 2147483647:
            for fi132 := 0; fi132 < 7; fi132++ {
                v133 := GenS3{F0: (((!false) && (uint32(v12) == 1)) || false)}
                _ = fi132
                _ = v133
            }
        }
        _ = fi130
        _ = v131
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v107
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
