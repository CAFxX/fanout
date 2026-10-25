package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS7 struct { F0 uint16; F1 string; F2 uint }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint32; F1 uint64 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS13 struct { F0 bool; F1 int16 }
type GenS14 struct { F0 uint32; F1 bool }
type GenS15 struct { F0 uint8; F1 uint; F2 int32 }
type GenS16 struct { F0 int8; F1 int64 }
type GenS17 struct { F0 uint; F1 bool }
type GenS18 struct { F0 int16; F1 int64; F2 uint16 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 uint; F1 int64 }
type GenS21 struct { F0 uint16; F1 bool }
type GenS22 struct { F0 int16 }
type GenS23 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS24 struct { F0 int; F1 int64; F2 int }
type GenS25 struct { F0 bool; F1 string }
func genFn0() int {
    v5 := GenS1{F0: -73, F1: 0, F2: -88}
    var v6 uint32 = 437
    var v7 uint64 = 555
    _ = v5
    _ = v6
    _ = v7
    return -181
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"hello", "foo bar", ("hello" + ("hello" + "\n\t"))}
    var v3 int64 = 0
    v4 := make(map[string]uint32)
    for fi8 := 0; fi8 < 2; fi8++ {
        v3 = -9223372036854775808
        switch ("" + ("hello" + "hello")) {
        case "hello":
            if true {
                v1[fi8] = false
            }
        default:
            switch (fi8 >> ((1) & 63)) {
            case 9223372036854775807:
                fmt.Println(fi8)
            case 1:
                if (42 >= 1) {
                    var v9 int16 = 42
                    v9 += v9
                    v9 += (v9 >> ((0) & 63))
                    _ = v9
                }
            default:
                if (-128 >= 31) {
                    switch (fi8 + (fi8 & -454)) {
                    case -2147483648:
                        v3 = 0
                    case 32767:
                        var v10 string = ("foo bar" + "\n\t")
                        _ = v10
                    }
                    v3 = v3
                    v3 += -8
                }
            }
        }
        if (!(-1 >= 1000)) {
            v3 -= v3
            if (0 != -1) {
                if ((!((false || false) || true)) || true) {
                    v3 -= v3
                } else {
                    v11 := GenS1{F0: -527, F1: 1, F2: fi8}
                    v3 -= v3
                    _ = v11
                }
                var v12 int32 = 0
                fmt.Println(v12, fi8)
                _ = v12
            }
        } else {
            if ((fi8 & fi8) <= (fi8 >> ((2) & 63))) {
                for fi13 := 0; fi13 < 9; fi13++ {
                    switch (fi8 & fi8) {
                    case 482:
                        v3 += (v3 / ((v3) | 1))
                    case -1:
                        var v14 int16 = 518
                        v14 = (v14 ^ (v14 - -1))
                        _ = v14
                    default:
                        var v15 string = "foo bar"
                        _ = v15
                    }
                    if true {
                        switch "hello" {
                        case "hello":
                            fmt.Println(fi8, v3)
                        case "foo bar":
                            var v16 float64 = 0.0
                            _ = v16
                        case "":
                            var v17 int = genFn0()
                            _ = v17
                        default:
                            var v18 bool = (255 <= 65535)
                            _ = v18
                        }
                        var v19 float64 = 1.5
                        switch fi13 {
                        case 0:
                            fmt.Println(fi8)
                        }
                        var v20 int16 = 42
                        v21 := GenS1{F0: 42, F1: 902, F2: fi13}
                        fmt.Println(v19, v3)
                        _ = v19
                        _ = v20
                        _ = v21
                    } else {
                        v3 += (-(v3))
                    }
                    _ = fi13
                }
                switch "\n\t" {
                case "hello":
                    v3 = v3
                case "":
                    switch fi8 {
                    case 4294967295:
                        fmt.Println(v3, v3)
                    case 829:
                        var v22 uint32 = 0
                        _ = v22
                    case 1:
                        v3 = (v3 << ((0) & 63))
                    default:
                        var v23 float32 = 1.5
                        _ = v23
                    }
                default:
                    v24 := GenS1{F0: -1, F1: 0, F2: fi8}
                    _ = v24
                }
                v3 -= -1
            }
            var v25 string = ("" + ("foo bar" + ""))
            delete(v1, (fi8 | -1))
            _ = v25
        }
        _ = fi8
    }
    var v26 int64 = v3
    switch 4294967295 {
    case 128:
        v26 += (^(v26))
    }
    v27 := GenS2{F0: 127, F1: 84}
    v28 := GenS2{F0: 127, F1: 2}
    v3 = -500
    v3 -= 42
    v3 += (v26 - (-(v26)))
    if ((!(42 != -566)) && false) {
        var v29 int = 128
        _ = v29
    } else {
        switch "hello" {
        case "\n\t":
            var v30 int = 947
            _ = v30
        case "foo bar":
            if ((176 > 1) && false) {
                var v31 int = genFn0()
                var v32 float32 = (0.0 * -2.25)
                _ = v31
                _ = v32
            } else {
                fmt.Println(v26, v26)
                var v33 bool = (4294967295 >= 42)
                var v34 uint = 593
                _ = v33
                _ = v34
            }
        case "hello":
            for fi35 := 0; fi35 < 18; fi35++ {
                var v36 int32 = 2147483647
                if (-32768 > -903) {
                    if (!true) {
                        for fi37 := 0; fi37 < 15; fi37++ {
                            fmt.Println(fi35, fi35, fi35)
                            var v38 bool = (127 <= 35)
                            v2 = append(v2, "hello")
                            _ = fi37
                            _ = v38
                        }
                    }
                    var v39 int16 = -1
                    for fi40 := 0; fi40 < 2; fi40++ {
                        if (true && false) {
                            fmt.Println(v36, v26)
                            if true {
                                var v41 int64 = (v26 << ((65535) & 63))
                                var v42 string = ""
                                _ = v41
                                _ = v42
                            } else {
                                v43 := GenS3{F0: -422, F1: 42}
                                v44 := GenS3{F0: -65, F1: 1}
                                if false {
                                    if false {
                                        v45 := GenS4{F0: 0, F1: v3}
                                        _ = v45
                                    } else {
                                        if (18446744073709551615 != 982) {
                                            v39 = -1
                                        } else {
                                            v2 = append(v2, ("hello" + ("foo bar" + ("" + "\n\t"))))
                                        }
                                        delete(v1, -9223372036854775808)
                                        v46 := GenS5{F0: 1}
                                        _ = v46
                                    }
                                    delete(v4, "hello")
                                    if true {
                                        switch "hello" {
                                        case "":
                                            if len(v2) > 0 {
                                                v2[0] = ""
                                            }
                                        case "foo bar":
                                            var v47 int64 = 9223372036854775807
                                            _ = v47
                                        default:
                                            var v48 int64 = 9223372036854775807
                                            _ = v48
                                        }
                                        fmt.Println(v26)
                                        var v49 uint16 = 607
                                        if (float64(3.14159) != (((float64(1e10) - (1.5 * 1.5)) / (-2.25 + 1.5)) * -2.25)) {
                                            if (936 >= v26) {
                                                v36 -= (v36 | v36)
                                                var v50 uint64 = 697
                                                _ = len(v4)
                                                _ = v50
                                            } else {
                                                switch "a" {
                                                case "hello":
                                                    if (false && (!true)) {
                                                        v39 = v39
                                                        switch 1 {
                                                        case 0:
                                                            var v51 float64 = (float64((0.1 - (float32(((0.0 * (((-2.25 + 0.1) - 0.0) / 0.1)) / 1.5)) / (float32(((0.1 * 0.1) + (-2.25 + 1.5))) + ((1.5 + ((1e10 + 0.1) + 3.14159)) - 0.1))))) + ((3.14159 - (-2.25 / 0.1)) * 3.14159))
                                                            _ = v51
                                                        case -1:
                                                            var v52 float64 = 1.5
                                                            _ = v52
                                                        case 127:
                                                            var v53 uint32 = 360
                                                            _ = v53
                                                        }
                                                        switch "\n\t" {
                                                        case "":
                                                            if true {
                                                                if true {
                                                                    v54 := GenS6{F0: 1, F1: 1, F2: 147}
                                                                    if (!(true || (!(!(false || true))))) {
                                                                        switch "a" {
                                                                        case "":
                                                                            v3 -= (^(v3))
                                                                        case "a":
                                                                            var v55 float64 = (3.14159 / 1.5)
                                                                            _ = v55
                                                                        }
                                                                        if ((false || true) && true) {
                                                                            if true {
                                                                                var v56 uint8 = 42
                                                                                _ = v56
                                                                            }
                                                                            if len(v2) > 0 {
                                                                                v2[0] = "foo bar"
                                                                            }
                                                                        }
                                                                        var v57 int8 = -128
                                                                        _ = v57
                                                                    } else {
                                                                        if true {
                                                                            if true {
                                                                                switch (fi35 << ((498) & 63)) {
                                                                                case 42:
                                                                                    var v58 int32 = v36
                                                                                    _ = v58
                                                                                }
                                                                                var v59 int16 = (v39 << ((42) & 63))
                                                                                v60 := GenS7{F0: v49, F1: "a", F2: 4294967295}
                                                                                var v61 uint32 = 352
                                                                                v59 = v39
                                                                                _ = v59
                                                                                _ = v60
                                                                                _ = v61
                                                                            } else {
                                                                                if (3.14159 == -2.25) {
                                                                                    var v62 float64 = (1.5 * (float64(3.14159) * 0.0))
                                                                                    _ = v62
                                                                                } else {
                                                                                    fmt.Println(v39)
                                                                                    var v63 uint64 = 18446744073709551615
                                                                                    var v64 int64 = v3
                                                                                    _ = v63
                                                                                    _ = v64
                                                                                }
                                                                                if true {
                                                                                    if false {
                                                                                        var v65 float64 = float64(float32(1e10))
                                                                                        _ = v65
                                                                                    }
                                                                                    var v66 float64 = 1.5
                                                                                    switch ("\n\t" + "foo bar") {
                                                                                    case "foo bar":
                                                                                        fmt.Println(fi40, v36, fi35)
                                                                                    case "":
                                                                                        _ = len(v2)
                                                                                    case "hello":
                                                                                        if false {
                                                                                            switch "a" {
                                                                                            case "":
                                                                                                switch 281 {
                                                                                                case 42:
                                                                                                    delete(v1, fi40)
                                                                                                case -339:
                                                                                                    var v67 uint32 = 2
                                                                                                    _ = v67
                                                                                                case 127:
                                                                                                    var v68 uint64 = 18446744073709551615
                                                                                                    switch fi40 {
                                                                                                    case -2147483648:
                                                                                                        v69 := GenS7{F0: 1, F1: ((("hello" + "hello") + "a") + "foo bar"), F2: 255}
                                                                                                        _ = v69
                                                                                                    }
                                                                                                    var v70 int16 = v39
                                                                                                    var v71 uint = 365
                                                                                                    _ = v68
                                                                                                    _ = v70
                                                                                                    _ = v71
                                                                                                }
                                                                                            case "foo bar":
                                                                                                v66 += (v66 / ((float64((1.5 - 1e10)) + (3.14159 + v66)) * -2.25))
                                                                                            default:
                                                                                                var v72 float32 = 0.0
                                                                                                _ = v72
                                                                                            }
                                                                                            v2 = append(v2, "foo bar")
                                                                                        } else {
                                                                                            v73 := GenS2{F0: 100, F1: 43}
                                                                                            v3 = v26
                                                                                            var v74 float32 = 0.0
                                                                                            switch "foo bar" {
                                                                                            case "foo bar":
                                                                                                fmt.Println(v3, v39)
                                                                                            }
                                                                                            v36 = -2147483648
                                                                                            _ = v73
                                                                                            _ = v74
                                                                                        }
                                                                                    }
                                                                                    _ = v66
                                                                                }
                                                                            }
                                                                        } else {
                                                                            fmt.Println(fi35)
                                                                            v75 := GenS8{F0: 337}
                                                                            _ = v75
                                                                        }
                                                                        var v76 string = "hello"
                                                                        _ = v76
                                                                    }
                                                                    _ = v54
                                                                }
                                                                v2 = append(v2, "foo bar")
                                                            }
                                                        case "\n\t":
                                                            var v77 uint16 = v49
                                                            switch -1 {
                                                            case 0:
                                                                switch "" {
                                                                case "hello":
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        if true {
                                                                            if false {
                                                                                var v78 float64 = 0.1
                                                                                switch (fi35 << ((4294967295) & 63)) {
                                                                                case 32767:
                                                                                    _ = len(v4)
                                                                                case -407:
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                }
                                                                                v79 := GenS4{F0: 42, F1: v3}
                                                                                _ = v78
                                                                                _ = v79
                                                                            }
                                                                            if true {
                                                                                v80 := GenS9{F0: 19, F1: 1}
                                                                                if (2 < 1) {
                                                                                    v77 -= 834
                                                                                    v81 := GenS9{F0: 12, F1: 598}
                                                                                    if (126 == 127) {
                                                                                        switch -770 {
                                                                                        case 0:
                                                                                            switch "hello" {
                                                                                            case "\n\t":
                                                                                                fmt.Println(v36)
                                                                                            }
                                                                                        case 127:
                                                                                            if true {
                                                                                                var v82 int64 = (v3 >> ((255) & 63))
                                                                                                var v83 uint8 = 255
                                                                                                fmt.Println(fi35)
                                                                                                _ = v82
                                                                                                _ = v83
                                                                                            }
                                                                                            if false {
                                                                                                switch "" {
                                                                                                case "":
                                                                                                    switch 32767 {
                                                                                                    case 443:
                                                                                                        var v84 uint32 = 211
                                                                                                        _ = v84
                                                                                                    default:
                                                                                                        v85 := v1[(fi35 >> ((255) & 63))]
                                                                                                        _ = v85
                                                                                                    }
                                                                                                }
                                                                                                if (!(0 != int8(fi35))) {
                                                                                                    switch "a" {
                                                                                                    case "a":
                                                                                                        v3 = v3
                                                                                                    case "foo bar":
                                                                                                        switch "" {
                                                                                                        case "":
                                                                                                            var v86 bool = (true && true)
                                                                                                            _ = v86
                                                                                                        }
                                                                                                    }
                                                                                                    var v87 int64 = (v3 * (v3 >> ((847) & 63)))
                                                                                                    if (v26 < -9223372036854775808) {
                                                                                                        switch "hello" {
                                                                                                        case "a":
                                                                                                            switch fi40 {
                                                                                                            case -564:
                                                                                                                v88 := GenS2{F0: 127, F1: 21}
                                                                                                                _ = v88
                                                                                                            case -2147483648:
                                                                                                                v87 -= v87
                                                                                                            }
                                                                                                        case "":
                                                                                                            var v89 uint = 4294967295
                                                                                                            _ = v89
                                                                                                        default:
                                                                                                            fmt.Println(v49, v77, v36)
                                                                                                        }
                                                                                                        v77 += v49
                                                                                                        v90 := v1[(fi40 ^ (fi35 - int(v36)))]
                                                                                                        var v91 int16 = 0
                                                                                                        switch 2147483647 {
                                                                                                        case 42:
                                                                                                            if (255 == 42) {
                                                                                                                if len(v2) > 0 {
                                                                                                                    v2[0] = ((("a" + "") + "\n\t") + "")
                                                                                                                }
                                                                                                                var v92 bool = v90
                                                                                                                _ = len(v1)
                                                                                                                switch (^(fi35)) {
                                                                                                                case 4294967295:
                                                                                                                    v93 := GenS7{F0: 32767, F1: (("a" + "") + "\n\t"), F2: 418}
                                                                                                                    _ = v93
                                                                                                                }
                                                                                                                if (!(122 != -128)) {
                                                                                                                    var v94 uint64 = 975
                                                                                                                    fmt.Println(v49, v92, v26)
                                                                                                                    _ = v94
                                                                                                                } else {
                                                                                                                    if ((-(v3)) != v87) {
                                                                                                                        var v95 int32 = 31
                                                                                                                        if v92 {
                                                                                                                            fmt.Println(v91)
                                                                                                                            v96 := GenS3{F0: genFn0(), F1: 18446744073709551615}
                                                                                                                            v91 = (^(v91))
                                                                                                                            _ = v96
                                                                                                                        }
                                                                                                                        var v97 uint = 42
                                                                                                                        _ = v95
                                                                                                                        _ = v97
                                                                                                                    }
                                                                                                                    if v90 {
                                                                                                                        _ = len(v1)
                                                                                                                        v39 = (v39 * (v91 >> ((2) & 63)))
                                                                                                                    }
                                                                                                                    if len(v2) > 0 {
                                                                                                                        v2[0] = "a"
                                                                                                                    }
                                                                                                                    var v98 int64 = (v26 >> ((14) & 63))
                                                                                                                    var v99 bool = ((v39 % ((-1) | 1)) >= (v39 + v39))
                                                                                                                    v26 += v98
                                                                                                                    switch (fi40 * (fi35 % (((-(fi40))) | 1))) {
                                                                                                                    case 4294967295:
                                                                                                                        v100 := GenS5{F0: -128}
                                                                                                                        _ = v100
                                                                                                                    case 42:
                                                                                                                        switch fi35 {
                                                                                                                        case 42:
                                                                                                                            if v99 {
                                                                                                                                if v99 {
                                                                                                                                    if (1 <= v77) {
                                                                                                                                        if len(v2) > 0 {
                                                                                                                                            v2[0] = ""
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    if true {
                                                                                                                                        fmt.Println(v92)
                                                                                                                                    }
                                                                                                                                    v2 = append(v2, "")
                                                                                                                                }
                                                                                                                                switch -599 {
                                                                                                                                case 4294967295:
                                                                                                                                    if len(v2) > 0 {
                                                                                                                                        v2[0] = "foo bar"
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    v101 := GenS2{F0: 100, F1: 127}
                                                                                                                                    _ = v101
                                                                                                                                }
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            if (212 != 18446744073709551615) {
                                                                                                                                delete(v4, (("hello" + (("a" + "") + ("hello" + "hello"))) + (("\n\t" + (((("a" + "\n\t") + ("\n\t" + "a")) + "hello") + "\n\t")) + "\n\t")))
                                                                                                                                var v102 bool = false
                                                                                                                                _ = v102
                                                                                                                            }
                                                                                                                            v103 := v1[fi35]
                                                                                                                            fmt.Println(v49)
                                                                                                                            _ = v103
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v104 string = "foo bar"
                                                                                                                        _ = v104
                                                                                                                    }
                                                                                                                    _ = v98
                                                                                                                    _ = v99
                                                                                                                }
                                                                                                                _ = v92
                                                                                                            } else {
                                                                                                                v3 -= 88
                                                                                                                delete(v1, fi35)
                                                                                                                switch "hello" {
                                                                                                                case "a":
                                                                                                                    v105 := GenS10{F0: ((("\n\t" + "foo bar") + "hello") + "hello")}
                                                                                                                    _ = v105
                                                                                                                case "\n\t":
                                                                                                                    if v90 {
                                                                                                                        v91 = (v91 | 1000)
                                                                                                                    } else {
                                                                                                                        v106 := GenS8{F0: 65535}
                                                                                                                        if true {
                                                                                                                            fmt.Println(fi40)
                                                                                                                        } else {
                                                                                                                            var v107 uint32 = 158
                                                                                                                            var v108 uint32 = v107
                                                                                                                            if (255 == 127) {
                                                                                                                                fmt.Println(fi35, v107)
                                                                                                                                var v109 float32 = (-2.25 / float32(float64(1.5)))
                                                                                                                                switch int(v87) {
                                                                                                                                case -1:
                                                                                                                                    v110 := v4[("hello" + "hello")]
                                                                                                                                    _ = v110
                                                                                                                                }
                                                                                                                                _ = v109
                                                                                                                            } else {
                                                                                                                                var v111 int64 = -732
                                                                                                                                _ = v111
                                                                                                                            }
                                                                                                                            _ = v107
                                                                                                                            _ = v108
                                                                                                                        }
                                                                                                                        _ = v106
                                                                                                                    }
                                                                                                                case "hello":
                                                                                                                    if (v39 == 0) {
                                                                                                                        v91 = (v91 + 1000)
                                                                                                                        switch 585 {
                                                                                                                        case 4294967295:
                                                                                                                            v91 = (v39 / ((32767) | 1))
                                                                                                                        case 127:
                                                                                                                            if false {
                                                                                                                                switch (("hello" + ("" + "foo bar")) + "a") {
                                                                                                                                case "foo bar":
                                                                                                                                    switch fi40 {
                                                                                                                                    case 1:
                                                                                                                                        var v112 uint8 = 89
                                                                                                                                        _ = v112
                                                                                                                                    case 4294967295:
                                                                                                                                        var v113 uint16 = 891
                                                                                                                                        _ = v113
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    switch genFn0() {
                                                                                                                                    case -2147483648:
                                                                                                                                        if len(v2) >= 2 {
                                                                                                                                            v2 = v2[1:len(v2)]
                                                                                                                                        }
                                                                                                                                    case 2147483647:
                                                                                                                                        fmt.Println(v77)
                                                                                                                                    default:
                                                                                                                                        if v90 {
                                                                                                                                            switch (fi35 ^ (fi40 ^ (fi35 ^ 360))) {
                                                                                                                                            case -979:
                                                                                                                                                if (v90 || (!false)) {
                                                                                                                                                    if (2 > 4294967295) {
                                                                                                                                                        var v114 string = "foo bar"
                                                                                                                                                        _ = v114
                                                                                                                                                    }
                                                                                                                                                    if false {
                                                                                                                                                        fmt.Println(v36)
                                                                                                                                                        if v90 {
                                                                                                                                                            var v115 string = "a"
                                                                                                                                                            _ = len(v1)
                                                                                                                                                            var v116 bool = true
                                                                                                                                                            _ = v115
                                                                                                                                                            _ = v116
                                                                                                                                                        }
                                                                                                                                                    } else {
                                                                                                                                                        fmt.Println(fi35, v3, v77)
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            case -1:
                                                                                                                                                if (2 == 42) {
                                                                                                                                                    v87 = -9223372036854775808
                                                                                                                                                } else {
                                                                                                                                                    var v117 int8 = 0
                                                                                                                                                    _ = v117
                                                                                                                                                }
                                                                                                                                            default:
                                                                                                                                                fmt.Println(v26)
                                                                                                                                            }
                                                                                                                                            v26 -= (v3 ^ v87)
                                                                                                                                            v118 := GenS5{F0: -128}
                                                                                                                                            var v119 float64 = 0.0
                                                                                                                                            _ = v118
                                                                                                                                            _ = v119
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                if v90 {
                                                                                                                                    var v120 string = ("" + ((("foo bar" + ("\n\t" + "")) + ("hello" + ("hello" + "\n\t"))) + "foo bar"))
                                                                                                                                    if (128 < 127) {
                                                                                                                                        if (1 != (v49 % (((v49 % ((42) | 1))) | 1))) {
                                                                                                                                            if false {
                                                                                                                                                v121 := GenS6{F0: 18446744073709551615, F1: 4294967295, F2: 42}
                                                                                                                                                _ = v121
                                                                                                                                            }
                                                                                                                                            v4[v120] = 65535
                                                                                                                                            var v122 int32 = v36
                                                                                                                                            v36 += (v122 & v36)
                                                                                                                                            fmt.Println(v87, v90)
                                                                                                                                            _ = v122
                                                                                                                                        } else {
                                                                                                                                            if true {
                                                                                                                                                v90 = (36 > 121)
                                                                                                                                            }
                                                                                                                                            var v123 int = (fi35 + 9223372036854775807)
                                                                                                                                            v77 += 0
                                                                                                                                            _ = v123
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        if len(v2) > 0 {
                                                                                                                                            v2[0] = (v120 + ("\n\t" + v120))
                                                                                                                                        }
                                                                                                                                        var v124 uint64 = 4294967295
                                                                                                                                        v87 = (v3 % ((v87) | 1))
                                                                                                                                        var v125 int = int(v36)
                                                                                                                                        _ = len(v4)
                                                                                                                                        _ = v124
                                                                                                                                        _ = v125
                                                                                                                                    }
                                                                                                                                    v77 = (v77 % (((v49 >> ((2) & 63))) | 1))
                                                                                                                                    _ = v120
                                                                                                                                } else {
                                                                                                                                    v36 = (^(v36))
                                                                                                                                }
                                                                                                                            }
                                                                                                                        case 0:
                                                                                                                            switch (fi40 & 250) {
                                                                                                                            case -1:
                                                                                                                                if v90 {
                                                                                                                                    if v90 {
                                                                                                                                        var v126 int64 = 0
                                                                                                                                        v126 += v3
                                                                                                                                        switch ((("\n\t" + ((("" + ("a" + "hello")) + "foo bar") + "")) + "") + "hello") {
                                                                                                                                        case "hello":
                                                                                                                                            switch "hello" {
                                                                                                                                            case "hello":
                                                                                                                                                var v127 float64 = float64(1.5)
                                                                                                                                                _ = v127
                                                                                                                                            case "foo bar":
                                                                                                                                                v91 += v91
                                                                                                                                            case "":
                                                                                                                                                if (42 == 2147483648) {
                                                                                                                                                    if v90 {
                                                                                                                                                        v39 = v39
                                                                                                                                                    } else {
                                                                                                                                                        v91 += v39
                                                                                                                                                        v90 = true
                                                                                                                                                    }
                                                                                                                                                    var v128 uint64 = 608
                                                                                                                                                    _ = v128
                                                                                                                                                } else {
                                                                                                                                                    var v129 uint8 = 255
                                                                                                                                                    _ = v129
                                                                                                                                                }
                                                                                                                                                _ = len(v2)
                                                                                                                                                switch ("" + "foo bar") {
                                                                                                                                                case "hello":
                                                                                                                                                    fmt.Println(v77)
                                                                                                                                                default:
                                                                                                                                                    v90 = v90
                                                                                                                                                }
                                                                                                                                                v130 := GenS11{F0: 0}
                                                                                                                                                _ = v130
                                                                                                                                            }
                                                                                                                                        case "foo bar":
                                                                                                                                            v131 := GenS12{F0: 42, F1: (v90 || v90), F2: 198}
                                                                                                                                            _ = v131
                                                                                                                                        default:
                                                                                                                                            v132 := GenS6{F0: 407, F1: 2147483648, F2: 1}
                                                                                                                                            _ = v132
                                                                                                                                        }
                                                                                                                                        fmt.Println(fi35, v87, v77)
                                                                                                                                        _ = v126
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                if v90 {
                                                                                                                                    var v133 uint = 0
                                                                                                                                    var v134 uint8 = 42
                                                                                                                                    if v90 {
                                                                                                                                        var v135 bool = (1e10 != 3.14159)
                                                                                                                                        _ = v135
                                                                                                                                    } else {
                                                                                                                                        if ((v3 % ((1) | 1)) < -1) {
                                                                                                                                            switch "\n\t" {
                                                                                                                                            case "a":
                                                                                                                                                v136 := GenS13{F0: false, F1: (v91 * v91)}
                                                                                                                                                _ = v136
                                                                                                                                            }
                                                                                                                                            v137 := GenS12{F0: 1, F1: v90, F2: 18446744073709551615}
                                                                                                                                            var v138 int = (fi35 % ((-192) | 1))
                                                                                                                                            var v139 int16 = -1
                                                                                                                                            v138 -= fi35
                                                                                                                                            _ = v137
                                                                                                                                            _ = v138
                                                                                                                                            _ = v139
                                                                                                                                        }
                                                                                                                                        v91 += v39
                                                                                                                                        v49 = 729
                                                                                                                                    }
                                                                                                                                    _ = v133
                                                                                                                                    _ = v134
                                                                                                                                } else {
                                                                                                                                    v140 := GenS1{F0: 0, F1: v77, F2: fi40}
                                                                                                                                    _ = v140
                                                                                                                                }
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            v141 := GenS14{F0: 269, F1: ("a" != ("" + "hello"))}
                                                                                                                            _ = v141
                                                                                                                        }
                                                                                                                        var v142 bool = true
                                                                                                                        _ = v142
                                                                                                                    }
                                                                                                                default:
                                                                                                                    if false {
                                                                                                                        v26 += v26
                                                                                                                        fmt.Println(fi35)
                                                                                                                        v3 = -9223372036854775808
                                                                                                                    }
                                                                                                                    fmt.Println(fi35, v87, v77)
                                                                                                                    var v143 int64 = (v26 * 1)
                                                                                                                    _ = v143
                                                                                                                }
                                                                                                                var v144 string = (("hello" + "hello") + "\n\t")
                                                                                                                _ = v144
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v90
                                                                                                        _ = v91
                                                                                                    } else {
                                                                                                        v145 := GenS15{F0: 0, F1: 42, F2: int32(v39)}
                                                                                                        if true {
                                                                                                            if false {
                                                                                                                v146 := GenS16{F0: -128, F1: (v87 / ((v26) | 1))}
                                                                                                                _ = v146
                                                                                                            }
                                                                                                            v26 += -1
                                                                                                        }
                                                                                                        _ = v145
                                                                                                    }
                                                                                                    switch (fi40 - fi35) {
                                                                                                    case -9223372036854775808:
                                                                                                        var v147 float32 = -2.25
                                                                                                        _ = v147
                                                                                                    case 42:
                                                                                                        var v148 int8 = 42
                                                                                                        _ = v148
                                                                                                    case -78:
                                                                                                        switch "hello" {
                                                                                                        case "\n\t":
                                                                                                            v26 = (-(v3))
                                                                                                        case "hello":
                                                                                                            delete(v1, (^(fi35)))
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v87
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        var v149 uint64 = 632
                                                                                        _ = v149
                                                                                    } else {
                                                                                        if false {
                                                                                            fmt.Println(v36)
                                                                                            if ((fi35 | (fi40 - fi35)) > genFn0()) {
                                                                                                var v150 uint32 = 914
                                                                                                _ = v150
                                                                                            } else {
                                                                                                fmt.Println(fi35, fi40)
                                                                                                v77 -= v49
                                                                                            }
                                                                                        } else {
                                                                                            v49 += (v77 + v77)
                                                                                            switch fi35 {
                                                                                            case 1:
                                                                                                var v151 int16 = v39
                                                                                                _ = v151
                                                                                            case 9223372036854775807:
                                                                                                fmt.Println(v39, v3, v3)
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v26, v49, v49)
                                                                                    }
                                                                                    _ = v81
                                                                                }
                                                                                _ = v80
                                                                            }
                                                                        }
                                                                    }
                                                                case "a":
                                                                    v152 := GenS17{F0: 42, F1: (((v3 ^ -9223372036854775808) < v3) || false)}
                                                                    _ = v152
                                                                case "":
                                                                    v39 -= (-(v39))
                                                                }
                                                            case -2147483648:
                                                                var v153 int64 = (^(v26))
                                                                _ = v153
                                                            case 2147483647:
                                                                v3 = (-(v26))
                                                            }
                                                            v39 += 42
                                                            _ = v77
                                                        }
                                                    } else {
                                                        switch (fi40 + genFn0()) {
                                                        case 1:
                                                            fmt.Println(v39)
                                                        case 9223372036854775807:
                                                            v154 := GenS16{F0: 0, F1: v26}
                                                            _ = v154
                                                        case 196:
                                                            if false {
                                                                fmt.Println(fi40)
                                                                v1[int(v36)] = (4294967295 < 42)
                                                            } else {
                                                                switch (fi40 << ((0) & 63)) {
                                                                case 32768:
                                                                    switch (("" + "a") + "\n\t") {
                                                                    case "foo bar":
                                                                        var v155 int = fi35
                                                                        _ = v155
                                                                    case "hello":
                                                                        if false {
                                                                            switch ("a" + ((("hello" + "hello") + "foo bar") + "a")) {
                                                                            case "":
                                                                                v156 := GenS18{F0: (v39 / (((v39 % ((int16(v3)) | 1))) | 1)), F1: -470, F2: (^(v49))}
                                                                                _ = v156
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        var v157 uint32 = 498
                                                                        _ = v157
                                                                    }
                                                                case -2147483648:
                                                                    v39 -= v39
                                                                default:
                                                                    if ((v36 & v36) < v36) {
                                                                        v49 -= (v49 >> ((2) & 63))
                                                                        var v158 uint8 = 0
                                                                        _ = v158
                                                                    } else {
                                                                        var v159 float64 = 0.1
                                                                        _ = v159
                                                                    }
                                                                }
                                                                var v160 float64 = ((0.1 - 3.14159) - 1.5)
                                                                if (4294967295 == 105) {
                                                                    switch "a" {
                                                                    case "":
                                                                        if len(v2) > 0 {
                                                                            v2[0] = "foo bar"
                                                                        }
                                                                    case "a":
                                                                        var v161 int = 32768
                                                                        if (!true) {
                                                                            v162 := GenS2{F0: -128, F1: 128}
                                                                            var v163 int16 = v39
                                                                            _ = v162
                                                                            _ = v163
                                                                        }
                                                                        _ = v161
                                                                    case "foo bar":
                                                                        v39 += v39
                                                                    }
                                                                    var v164 int64 = 33
                                                                    if (1 <= 2) {
                                                                        if ((-(v36)) < v36) {
                                                                            var v165 int32 = (v36 | v36)
                                                                            _ = v165
                                                                        }
                                                                        if (false || (((0.1 + (-2.25 / (0.0 - 3.14159))) + (3.14159 - 3.14159)) <= 1.5)) {
                                                                            var v166 uint8 = 42
                                                                            _ = v166
                                                                        } else {
                                                                            v36 += (v36 << ((42) & 63))
                                                                        }
                                                                        if true {
                                                                            var v167 float32 = 1e10
                                                                            if (genFn0() != 4294967295) {
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    var v168 float64 = (0.1 / v160)
                                                                                    _ = v168
                                                                                case "\n\t":
                                                                                    if true {
                                                                                        if false {
                                                                                            switch (^(fi35)) {
                                                                                            case -276:
                                                                                                fmt.Println(v49, v26, v3)
                                                                                            }
                                                                                            v167 = v167
                                                                                        } else {
                                                                                            var v169 uint = 90
                                                                                            v3 = 9223372036854775807
                                                                                            _ = v169
                                                                                        }
                                                                                        v26 -= v26
                                                                                        var v170 int32 = v36
                                                                                        _ = v170
                                                                                    }
                                                                                }
                                                                                v171 := GenS19{F0: 127}
                                                                                _ = v171
                                                                            }
                                                                            var v172 int32 = v36
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            _ = v167
                                                                            _ = v172
                                                                        }
                                                                    } else {
                                                                        v36 -= (v36 * v36)
                                                                    }
                                                                    _ = v164
                                                                } else {
                                                                    var v173 float32 = (0.1 / (1e10 - (1.5 - float32(v160))))
                                                                    if (v173 == (v173 - (v173 + ((1.5 / v173) - 1.5)))) {
                                                                        if false {
                                                                            if false {
                                                                                v160 = v160
                                                                            } else {
                                                                                var v174 string = "foo bar"
                                                                                var v175 int = int(v26)
                                                                                var v176 float64 = (0.0 - (0.0 - 1.5))
                                                                                _ = v174
                                                                                _ = v175
                                                                                _ = v176
                                                                            }
                                                                            if true {
                                                                                v177 := v4[(("hello" + "hello") + "a")]
                                                                                _ = len(v4)
                                                                                fmt.Println(v173)
                                                                                _ = v177
                                                                            } else {
                                                                                var v178 int8 = 127
                                                                                var v179 uint32 = 2
                                                                                v173 -= float32(v160)
                                                                                _ = v178
                                                                                _ = v179
                                                                            }
                                                                            v180 := v4["a"]
                                                                            _ = v180
                                                                        }
                                                                        v181 := GenS20{F0: 844, F1: -210}
                                                                        var v182 uint = 42
                                                                        v160 -= ((float64(0.0) - 1e10) / v160)
                                                                        v26 += (v3 % ((v26) | 1))
                                                                        _ = v181
                                                                        _ = v182
                                                                    } else {
                                                                        v36 -= v36
                                                                        if false {
                                                                            if (846 > 415) {
                                                                                v3 -= v26
                                                                                _ = len(v1)
                                                                                v160 -= (v160 * float64(v173))
                                                                            }
                                                                            if len(v2) > 0 {
                                                                                v2[0] = (("hello" + "a") + (("foo bar" + "") + "hello"))
                                                                            }
                                                                        } else {
                                                                            _ = len(v2)
                                                                            if (1 == 2) {
                                                                                switch "a" {
                                                                                case "\n\t":
                                                                                    var v183 uint8 = 0
                                                                                    _ = v183
                                                                                }
                                                                                _ = len(v2)
                                                                                v160 = v160
                                                                            } else {
                                                                                var v184 int8 = -1
                                                                                v160 = v160
                                                                                _ = v184
                                                                            }
                                                                            v49 = (v49 | v49)
                                                                        }
                                                                        v39 = (v39 | (-(v39)))
                                                                    }
                                                                    _ = v173
                                                                }
                                                                _ = v160
                                                            }
                                                        default:
                                                            v185 := v4["hello"]
                                                            _ = v185
                                                        }
                                                    }
                                                case "a":
                                                    if false {
                                                        v186 := GenS5{F0: 1}
                                                        var v187 uint8 = 2
                                                        _ = v186
                                                        _ = v187
                                                    } else {
                                                        v49 += 683
                                                    }
                                                    v49 += 1
                                                    v188 := GenS1{F0: (v36 % ((-1) | 1)), F1: (v49 & v49), F2: fi40}
                                                    _ = v188
                                                case "":
                                                    v3 += v3
                                                }
                                                v189 := GenS21{F0: v49, F1: (!false)}
                                                v3 = (v26 * v3)
                                                _ = v189
                                            }
                                            v26 += (v26 | v26)
                                        } else {
                                            v49 = 2
                                            fmt.Println(v49)
                                            v2 = append(v2, "")
                                        }
                                        _ = v49
                                    }
                                } else {
                                    v190 := v1[fi35]
                                    switch "hello" {
                                    case "a":
                                        fmt.Println(fi35, v190, v36)
                                    default:
                                        if v190 {
                                            var v191 uint = 1
                                            v39 = v39
                                            var v192 int64 = (v26 << (((v191 + v191)) & 63))
                                            var v193 int = (fi40 + fi35)
                                            _ = v191
                                            _ = v192
                                            _ = v193
                                        }
                                    }
                                    v194 := v4["foo bar"]
                                    fmt.Println(v26, v39)
                                    v3 = (v3 * int64(fi40))
                                    fmt.Println(fi40, fi40, v26)
                                    if v190 {
                                        if (((1.5 / (0.1 - (0.0 / 1.0))) + 0.1) < 1.5) {
                                            v1[(fi35 / (((fi40 - fi40)) | 1))] = (!v190)
                                        } else {
                                            v39 -= v39
                                        }
                                        delete(v4, "foo bar")
                                    }
                                    _ = v190
                                    _ = v194
                                }
                                _ = v43
                                _ = v44
                            }
                        } else {
                            var v195 string = ""
                            _ = v195
                        }
                        if (0 == 101) {
                            v39 = (v39 << ((42) & 63))
                        } else {
                            delete(v1, 9223372036854775807)
                        }
                        var v196 bool = false
                        _ = fi40
                        _ = v196
                    }
                    _ = v39
                } else {
                    var v197 uint32 = 0
                    v198 := GenS22{F0: 465}
                    v199 := GenS15{F0: 1, F1: 255, F2: 42}
                    _ = v197
                    _ = v198
                    _ = v199
                }
                _ = fi35
                _ = v36
            }
        }
    }
    if (-128 <= 32) {
        v26 = v26
    } else {
        v200 := GenS23{F0: 0, F1: (2147483647 != 191), F2: 825}
        _ = v200
    }
    for fi201 := 0; fi201 < 20; fi201++ {
        v202 := GenS4{F0: 23, F1: v26}
        _ = len(v2)
        var v203 float32 = float32(0.0)
        _ = fi201
        _ = v202
        _ = v203
    }
    var v204 int8 = -1
    var v205 int16 = -131
    for fi206 := 0; fi206 < 4; fi206++ {
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        delete(v1, fi206)
        _ = fi206
    }
    _ = len(v2)
    switch 336 {
    case 32767:
        var v207 int64 = v3
        _ = v207
    case 9223372036854775807:
        switch "foo bar" {
        case "\n\t":
            if (!(v205 == 436)) {
                var v208 float32 = (0.0 + 0.1)
                for fi209 := 0; fi209 < 2; fi209++ {
                    v205 += (v205 - (v205 * v205))
                    fmt.Println(fi209, v3, v204)
                    _ = fi209
                }
                _ = v208
            }
        default:
            v210 := GenS24{F0: 2147483647, F1: v26, F2: genFn0()}
            _ = v210
        }
    case -2147483648:
        v211 := GenS25{F0: true, F1: ("foo bar" + ("foo bar" + "hello"))}
        _ = v211
    }
    v205 = 846
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v204:", v204)
    fmt.Println("v205:", v205)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v204
    _ = v205
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
