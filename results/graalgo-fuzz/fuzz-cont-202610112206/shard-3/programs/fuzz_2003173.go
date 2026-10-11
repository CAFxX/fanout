package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint; F1 uint; F2 int32 }
type GenS6 struct { F0 uint32; F1 uint; F2 int32 }
type GenS7 struct { F0 uint32; F1 int16; F2 int }
type GenS8 struct { F0 uint; F1 string; F2 int }
type GenS9 struct { F0 uint; F1 string }
type GenS10 struct { F0 int32; F1 int16; F2 int16 }
type GenS11 struct { F0 int32; F1 uint; F2 uint }
type GenS12 struct { F0 uint64; F1 uint16; F2 string }
type GenS13 struct { F0 bool; F1 uint8 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 string; F1 uint8 }
type GenS18 struct { F0 int16; F1 uint8; F2 int32 }
type GenS19 struct { F0 int; F1 int32; F2 bool }
type GenS20 struct { F0 uint64; F1 uint32 }
type GenS21 struct { F0 int16 }
type GenS22 struct { F0 int; F1 int64 }
type GenS23 struct { F0 bool; F1 int32 }
type GenS24 struct { F0 uint; F1 uint64 }
type GenS25 struct { F0 uint16 }
type GenS26 struct { F0 uint32; F1 int8 }
type GenS27 struct { F0 uint16 }
type GenS28 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS29 struct { F0 uint16; F1 uint }
type GenS30 struct { F0 int32; F1 uint64; F2 int8 }
type GenS31 struct { F0 int8; F1 int16; F2 string }
type GenS32 struct { F0 uint; F1 uint }
type GenS33 struct { F0 int64 }
func genFn0() uint {
    if true {
        if false {
            var v4 int32 = -2147483648
            _ = v4
        } else {
            v5 := []int{1, -2147483648, 128, 42}
            _ = v5
        }
        v6 := GenS1{F0: 4294967295}
        _ = v6
    } else {
        var v7 int8 = 42
        var v8 float64 = 3.14159
        if (42 != 1) {
            if ((!true) || true) {
                switch -2147483648 {
                case 15:
                    for fi9 := 0; fi9 < 2; fi9++ {
                        var v10 uint64 = 42
                        _ = fi9
                        _ = v10
                    }
                case -391:
                    var v11 int8 = (v7 | (v7 % (((v7 + 18)) | 1)))
                    _ = v11
                default:
                    v12 := []int{-2147483648, 0, 32767}
                    _ = v12
                }
            } else {
                v13 := []int{0}
                fmt.Println(v8, v7)
                for fi14 := 0; fi14 < 18; fi14++ {
                    fmt.Println(fi14, fi14, v7)
                    var v15 string = ("a" + "\n\t")
                    var v16 string = v15
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                _ = v13
            }
            if (!true) {
                var v17 int32 = 0
                v18 := []int{127, 4294967295, int(v7), 4294967295}
                v19 := GenS1{F0: 766}
                _ = v17
                _ = v18
                _ = v19
            }
            v7 -= v7
        }
        _ = v7
        _ = v8
    }
    var v20 uint8 = 140
    if (-44 == 119) {
        v21 := GenS1{F0: 1}
        _ = v21
    }
    v20 -= (v20 >> ((875) & 63))
    var v22 uint = 0
    _ = v20
    _ = v22
    return v22
}

func main() {
    v1 := []string{("a" + "foo bar"), ("foo bar" + "hello"), "a", ("a" + ("foo bar" + "\n\t"))}
    v2 := make(map[int]int)
    var v3 int16 = 312
    _ = len(v1)
    for fi23 := 0; fi23 < 18; fi23++ {
        v24 := GenS2{F0: ("a" + ("\n\t" + "foo bar"))}
        var v25 uint8 = 218
        var v26 string = "a"
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    delete(v2, 0)
    v27 := v2[-1]
    switch ("a" + ("" + "a")) {
    case "\n\t":
        fmt.Println(v27)
    default:
        if len(v1) > 0 {
            v1[0] = (("hello" + "\n\t") + "a")
        }
    }
    v3 -= v3
    v27 += v27
    switch "hello" {
    case "hello":
        var v28 float64 = (1e10 + (float64((3.14159 + -2.25)) * (0.0 - ((-2.25 - 1.5) + ((((1.5 / (1e10 / (float64(0.1) * (1.5 + -2.25)))) - (((0.0 - (float64(1.5) - 1e10)) / 1.5) + (float64(1e10) / ((-2.25 / 3.14159) * 1.5)))) / 1e10) * 0.1)))))
        _ = v28
    case "\n\t":
        if false {
            if false {
                fmt.Println(v3, v27)
                switch (v27 & v27) {
                case 127:
                    for fi29 := 0; fi29 < 7; fi29++ {
                        for fi30 := 0; fi30 < 3; fi30++ {
                            if (((true && (!true)) && ((42 != 127) && true)) || (704 == genFn0())) {
                                v31 := GenS2{F0: "a"}
                                if (((-2.25 * 1e10) != 1e10) || false) {
                                    var v32 int8 = 70
                                    _ = v32
                                }
                                v27 -= (-(fi30))
                                _ = v31
                            } else {
                                fmt.Println(v27, v27, v3)
                                var v33 uint32 = 623
                                _ = v33
                            }
                            switch (fi29 & v27) {
                            case -620:
                                v3 = (-(v3))
                            case -517:
                                if true {
                                    if true {
                                        v27 = (-(fi30))
                                        if ((((1.5 * float32(-2.25)) / float32(1e10)) > 1.5) || true) {
                                            v34 := GenS3{F0: (fi29 << ((944) & 63))}
                                            v27 += 42
                                            _ = v34
                                        }
                                    }
                                    var v35 bool = true
                                    _ = v35
                                }
                            default:
                                var v36 uint = 1
                                _ = v36
                            }
                            _ = len(v2)
                            _ = fi30
                        }
                        _ = fi29
                    }
                case -392:
                    var v37 int16 = v3
                    _ = v37
                case 176:
                    v3 -= -1
                default:
                    for fi38 := 0; fi38 < 2; fi38++ {
                        for fi39 := 0; fi39 < 14; fi39++ {
                            if true {
                                var v40 uint64 = 42
                                v41 := GenS4{F0: (("hello" + "a") + "foo bar"), F1: 42}
                                v42 := GenS2{F0: ""}
                                _ = v40
                                _ = v41
                                _ = v42
                            } else {
                                v27 = (fi38 | 128)
                                v1 = append(v1, ("a" + "\n\t"))
                                if (42 < 1) {
                                    v43 := v2[471]
                                    switch "" {
                                    case "":
                                        v44 := GenS5{F0: 0, F1: 389, F2: 42}
                                        var v45 float64 = 1e10
                                        if false {
                                            if (12 == 42) {
                                                v1 = append(v1, "foo bar")
                                                var v46 uint32 = 0
                                                v47 := GenS6{F0: v46, F1: 65535, F2: 1}
                                                _ = v46
                                                _ = v47
                                            } else {
                                                switch -1 {
                                                case 466:
                                                    var v48 float32 = -2.25
                                                    _ = v48
                                                case 9223372036854775807:
                                                    v3 -= v3
                                                case 32768:
                                                    if true {
                                                        if ((-2147483648 <= -756) && (!false)) {
                                                            if false {
                                                                switch ("a" + "a") {
                                                                case "":
                                                                    v49 := GenS6{F0: 2, F1: 18446744073709551615, F2: 2147483647}
                                                                    _ = v49
                                                                case "foo bar":
                                                                    var v50 int8 = -128
                                                                    _ = v50
                                                                }
                                                            }
                                                            v1 = append(v1, (((("hello" + (("hello" + "") + "foo bar")) + "\n\t") + "") + "\n\t"))
                                                            v3 = 42
                                                            if (18446744073709551615 > genFn0()) {
                                                                v51 := GenS4{F0: (((("foo bar" + "\n\t") + "foo bar") + ("\n\t" + "\n\t")) + "hello"), F1: 1}
                                                                _ = v51
                                                            }
                                                            v3 = (v3 / (((v3 % (((v3 << ((1) & 63))) | 1))) | 1))
                                                        } else {
                                                            fmt.Println(fi39, fi38)
                                                            v2[v27] = (fi38 + (v27 + (fi39 - (fi38 % ((v43) | 1)))))
                                                        }
                                                        v2[(fi39 / (((v43 | (fi39 * (-(v27))))) | 1))] = v43
                                                        switch (^(v43)) {
                                                        case 42:
                                                            if (v45 >= v45) {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                fmt.Println(fi39, v27, fi38)
                                                                v3 -= (v3 + (v3 | (v3 | v3)))
                                                                var v52 string = "hello"
                                                                v53 := GenS5{F0: 57, F1: 0, F2: 96}
                                                                _ = v52
                                                                _ = v53
                                                            }
                                                        case -862:
                                                            v43 = 42
                                                        case -2147483648:
                                                            _ = len(v1)
                                                        }
                                                    }
                                                default:
                                                    var v54 int64 = 273
                                                    v45 += v45
                                                    _ = v54
                                                }
                                            }
                                            switch ("" + "hello") {
                                            case "a":
                                                if (42 != 74) {
                                                    if (true && true) {
                                                        v45 = float64((float32(0.0) - 0.1))
                                                        if (0 >= 9223372036854775807) {
                                                            switch "" {
                                                            case "":
                                                                v45 = float64(-2.25)
                                                            case "foo bar":
                                                                fmt.Println(fi38, v43)
                                                            case "a":
                                                                var v55 uint32 = 0
                                                                _ = v55
                                                            }
                                                            switch "hello" {
                                                            case "foo bar":
                                                                var v56 uint8 = 190
                                                                _ = v56
                                                            case "a":
                                                                if false {
                                                                    v27 = fi39
                                                                    v2[128] = v43
                                                                    var v57 uint8 = 1
                                                                    _ = v57
                                                                }
                                                            case "\n\t":
                                                                switch (("foo bar" + "foo bar") + "hello") {
                                                                case "":
                                                                    switch int(v3) {
                                                                    case -1:
                                                                        if (18446744073709551615 != 0) {
                                                                            v58 := GenS7{F0: 874, F1: 1, F2: (fi39 % ((2147483647) | 1))}
                                                                            var v59 float64 = v45
                                                                            v43 = (fi39 - fi39)
                                                                            _ = v58
                                                                            _ = v59
                                                                        }
                                                                    case -9223372036854775808:
                                                                        switch "" {
                                                                        case "":
                                                                            fmt.Println(fi38, v3, v27)
                                                                            v43 = fi38
                                                                        case "foo bar":
                                                                            switch fi38 {
                                                                            case 2147483647:
                                                                                switch ((("\n\t" + "a") + "hello") + ("" + "hello")) {
                                                                                case "foo bar":
                                                                                    if ((true && ((127 < 40) || true)) && (true && (false || (0 < -128)))) {
                                                                                        if (true || true) {
                                                                                            if ((-125 != 0) || ((!(-9223372036854775808 == 0)) && false)) {
                                                                                                switch (^(fi39)) {
                                                                                                case 32768:
                                                                                                    fmt.Println(fi39, v43, v27)
                                                                                                default:
                                                                                                    if true {
                                                                                                        _ = len(v2)
                                                                                                        v43 = fi39
                                                                                                    }
                                                                                                }
                                                                                                var v60 uint8 = 0
                                                                                                var v61 bool = ("foo bar" > "a")
                                                                                                _ = v60
                                                                                                _ = v61
                                                                                            } else {
                                                                                                switch "\n\t" {
                                                                                                case "a":
                                                                                                    switch "\n\t" {
                                                                                                    case "":
                                                                                                        v62 := GenS1{F0: 2}
                                                                                                        _ = v62
                                                                                                    case "hello":
                                                                                                        v63 := GenS3{F0: v43}
                                                                                                        _ = v63
                                                                                                    case "foo bar":
                                                                                                        v1 = append(v1, "")
                                                                                                    default:
                                                                                                        var v64 int32 = 101
                                                                                                        _ = v64
                                                                                                    }
                                                                                                default:
                                                                                                    v65 := GenS8{F0: 18446744073709551615, F1: "foo bar", F2: 32767}
                                                                                                    _ = v65
                                                                                                }
                                                                                                v43 += fi39
                                                                                            }
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                            v66 := GenS8{F0: 0, F1: (("foo bar" + "\n\t") + "foo bar"), F2: (^(fi39))}
                                                                                            _ = v66
                                                                                        }
                                                                                        var v67 int64 = 0
                                                                                        _ = v67
                                                                                    } else {
                                                                                        var v68 uint16 = 32767
                                                                                        _ = v68
                                                                                    }
                                                                                case "hello":
                                                                                    if (42 > 1) {
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = ("\n\t" + "\n\t")
                                                                                        }
                                                                                    }
                                                                                case "":
                                                                                    var v69 uint32 = 1
                                                                                    _ = v69
                                                                                default:
                                                                                    var v70 float64 = ((v45 / (0.0 + v45)) + v45)
                                                                                    _ = v70
                                                                                }
                                                                            case -397:
                                                                                if (0 == 32767) {
                                                                                    v71 := v2[(v43 & -353)]
                                                                                    _ = v71
                                                                                }
                                                                            default:
                                                                                var v72 int64 = -1
                                                                                _ = v72
                                                                            }
                                                                        }
                                                                    case -943:
                                                                        switch (fi38 >> ((237) & 63)) {
                                                                        case 128:
                                                                            var v73 float64 = (v45 + 1e10)
                                                                            v74 := GenS8{F0: 18446744073709551615, F1: "\n\t", F2: fi39}
                                                                            fmt.Println(v27, v43)
                                                                            _ = v73
                                                                            _ = v74
                                                                        case -9223372036854775808:
                                                                            if (!(482 >= 68)) {
                                                                                switch -13 {
                                                                                case 2147483647:
                                                                                    fmt.Println(fi39, v43)
                                                                                case 768:
                                                                                    var v75 int8 = 1
                                                                                    _ = v75
                                                                                case -2147483648:
                                                                                    v76 := GenS4{F0: "a", F1: -1}
                                                                                    _ = v76
                                                                                default:
                                                                                    fmt.Println(v45, v43)
                                                                                }
                                                                                _ = len(v1)
                                                                                var v77 int64 = 9223372036854775807
                                                                                if (42 < (v43 / (((fi39 - (fi39 / (((fi39 * fi38)) | 1)))) | 1))) {
                                                                                    if (v3 > (v3 / ((int16(v77)) | 1))) {
                                                                                        if true {
                                                                                            var v78 int64 = v77
                                                                                            _ = v78
                                                                                        } else {
                                                                                            delete(v2, (v43 % ((1) | 1)))
                                                                                            v79 := GenS1{F0: 165}
                                                                                            _ = v79
                                                                                        }
                                                                                        v80 := GenS4{F0: (("\n\t" + "foo bar") + "hello"), F1: 42}
                                                                                        v43 -= 2147483647
                                                                                        _ = v80
                                                                                    } else {
                                                                                        var v81 uint = 358
                                                                                        v82 := v2[(^(v27))]
                                                                                        _ = v81
                                                                                        _ = v82
                                                                                    }
                                                                                    fmt.Println(v3, fi38, fi38)
                                                                                }
                                                                                v83 := GenS9{F0: 255, F1: "\n\t"}
                                                                                _ = v77
                                                                                _ = v83
                                                                            } else {
                                                                                switch fi38 {
                                                                                case 9223372036854775807:
                                                                                    var v84 string = (((("\n\t" + "hello") + "") + "foo bar") + "")
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = v84
                                                                                    }
                                                                                    _ = len(v1)
                                                                                    _ = v84
                                                                                default:
                                                                                    delete(v2, (v27 % ((fi39) | 1)))
                                                                                }
                                                                                fmt.Println(v27)
                                                                                v27 = v43
                                                                            }
                                                                        default:
                                                                            v85 := GenS1{F0: 1}
                                                                            _ = v85
                                                                        }
                                                                    default:
                                                                        delete(v2, (v43 & (-(v27))))
                                                                    }
                                                                default:
                                                                    fmt.Println(v43)
                                                                }
                                                                v86 := GenS2{F0: ("a" + (("foo bar" + ("\n\t" + (("\n\t" + "") + ("" + ("foo bar" + "hello"))))) + "\n\t"))}
                                                                _ = v86
                                                            default:
                                                                if len(v1) > 0 {
                                                                    v1[0] = "hello"
                                                                }
                                                            }
                                                        } else {
                                                            fmt.Println(v3)
                                                        }
                                                        switch ("" + "") {
                                                        case "foo bar":
                                                            _ = len(v2)
                                                        default:
                                                            v27 -= (v43 / ((-352) | 1))
                                                        }
                                                    }
                                                } else {
                                                    var v87 int32 = 343
                                                    var v88 int = -581
                                                    if ((v88 << ((693) & 63)) >= 2147483647) {
                                                        switch v88 {
                                                        case -2147483648:
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v89 := GenS3{F0: 2147483647}
                                                                if ((1.5 + 0.0) == 0.0) {
                                                                    var v90 uint16 = 0
                                                                    fmt.Println(v27)
                                                                    _ = v90
                                                                }
                                                                _ = v89
                                                            case "foo bar":
                                                                v87 = 1
                                                            case "hello":
                                                                switch v88 {
                                                                case -1:
                                                                    if false {
                                                                        v45 += 1.5
                                                                        if (!((-126 < -64) || false)) {
                                                                            var v91 int = fi39
                                                                            _ = v91
                                                                        } else {
                                                                            v92 := GenS10{F0: (v87 & 42), F1: (^(v3)), F2: (v3 & v3)}
                                                                            fmt.Println(v43, v27, fi38)
                                                                            if false {
                                                                                var v93 uint8 = 42
                                                                                if (2 >= 1) {
                                                                                    if false {
                                                                                        var v94 uint = 4294967295
                                                                                        _ = v94
                                                                                    }
                                                                                    v43 = (-(fi38))
                                                                                    _ = len(v2)
                                                                                } else {
                                                                                    if false {
                                                                                        if false {
                                                                                            switch ("\n\t" + "a") {
                                                                                            case "":
                                                                                                switch v88 {
                                                                                                case 32768:
                                                                                                    fmt.Println(v43)
                                                                                                case -4:
                                                                                                    v88 = 687
                                                                                                case 1:
                                                                                                    switch v43 {
                                                                                                    case 0:
                                                                                                        v45 -= (v45 + (((-2.25 + float64(0.1)) + float64(3.14159)) * (((v45 + v45) + 3.14159) + 0.0)))
                                                                                                    }
                                                                                                }
                                                                                            case "a":
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                            }
                                                                                            v95 := GenS4{F0: "\n\t", F1: -1}
                                                                                            _ = v95
                                                                                        }
                                                                                    }
                                                                                    var v96 int32 = v87
                                                                                    _ = v96
                                                                                }
                                                                                var v97 int64 = -602
                                                                                _ = v93
                                                                                _ = v97
                                                                            } else {
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    if true {
                                                                                        v3 = v3
                                                                                    } else {
                                                                                        v98 := GenS11{F0: (v87 >> ((334) & 63)), F1: 794, F2: 255}
                                                                                        _ = v98
                                                                                    }
                                                                                default:
                                                                                    v99 := v2[v43]
                                                                                    _ = v99
                                                                                }
                                                                                if (false && (!true)) {
                                                                                    if true {
                                                                                        fmt.Println(fi39, fi38)
                                                                                    } else {
                                                                                        v3 += v3
                                                                                        if (-563 <= (v88 << ((4294967295) & 63))) {
                                                                                            switch (^(v43)) {
                                                                                            case -371:
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                            default:
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = ("" + "foo bar")
                                                                                                }
                                                                                                switch fi39 {
                                                                                                case 0:
                                                                                                    v88 -= (v88 - (^(fi39)))
                                                                                                case -69:
                                                                                                    v2[-2147483648] = 647
                                                                                                default:
                                                                                                    var v100 uint = 65535
                                                                                                    var v101 bool = true
                                                                                                    _ = v100
                                                                                                    _ = v101
                                                                                                }
                                                                                                if (127 != 0) {
                                                                                                    v45 += (((v45 + (0.1 * ((float64(-2.25) - v45) - v45))) - (float64(3.14159) - 0.1)) - v45)
                                                                                                    if false {
                                                                                                        var v102 uint32 = 0
                                                                                                        _ = len(v2)
                                                                                                        if true {
                                                                                                            v103 := GenS11{F0: v87, F1: 42, F2: 555}
                                                                                                            if (2 > 1) {
                                                                                                                v104 := GenS11{F0: (v87 & (-(v87))), F1: 18446744073709551615, F2: 131}
                                                                                                                if len(v1) > 0 {
                                                                                                                    v1[0] = (("a" + "\n\t") + "hello")
                                                                                                                }
                                                                                                                delete(v2, 704)
                                                                                                                _ = v104
                                                                                                            }
                                                                                                            _ = v103
                                                                                                        } else {
                                                                                                            var v105 bool = true
                                                                                                            switch 127 {
                                                                                                            case 0:
                                                                                                                fmt.Println(v88)
                                                                                                            }
                                                                                                            fmt.Println(v102, v45)
                                                                                                            _ = v105
                                                                                                        }
                                                                                                        _ = v102
                                                                                                    } else {
                                                                                                        var v106 float32 = 1e10
                                                                                                        _ = v106
                                                                                                    }
                                                                                                    v87 += v87
                                                                                                }
                                                                                            }
                                                                                            v43 -= v88
                                                                                        }
                                                                                        v87 += (v87 % (((-(v87))) | 1))
                                                                                    }
                                                                                    var v107 uint64 = 937
                                                                                    _ = v107
                                                                                } else {
                                                                                    var v108 int64 = -9223372036854775808
                                                                                    var v109 uint = 65535
                                                                                    v110 := GenS12{F0: 18446744073709551615, F1: 692, F2: ("a" + "foo bar")}
                                                                                    _ = v108
                                                                                    _ = v109
                                                                                    _ = v110
                                                                                }
                                                                                v111 := GenS3{F0: v27}
                                                                                if (false || false) {
                                                                                    if true {
                                                                                        if (646 >= 2) {
                                                                                            var v112 int64 = -9223372036854775808
                                                                                            _ = v112
                                                                                        } else {
                                                                                            switch v27 {
                                                                                            case 0:
                                                                                                var v113 uint8 = 2
                                                                                                _ = v113
                                                                                            default:
                                                                                                if true {
                                                                                                    var v114 float32 = 1.5
                                                                                                    v115 := GenS13{F0: (!true), F1: 1}
                                                                                                    var v116 int32 = v87
                                                                                                    v45 = 3.14159
                                                                                                    _ = v114
                                                                                                    _ = v115
                                                                                                    _ = v116
                                                                                                } else {
                                                                                                    var v117 int8 = 1
                                                                                                    _ = v117
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v45, v88)
                                                                                            var v118 uint = 343
                                                                                            _ = v118
                                                                                        }
                                                                                        switch ("\n\t" + (("\n\t" + "a") + ("foo bar" + "foo bar"))) {
                                                                                        case "foo bar":
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = (("" + "\n\t") + "a")
                                                                                            }
                                                                                        default:
                                                                                            v87 -= (v87 - v87)
                                                                                        }
                                                                                        v2[(fi39 * -2147483648)] = (v88 << ((586) & 63))
                                                                                    } else {
                                                                                        _ = len(v1)
                                                                                    }
                                                                                    var v119 int16 = -661
                                                                                    _ = v119
                                                                                } else {
                                                                                    var v120 string = "hello"
                                                                                    switch v43 {
                                                                                    case 0:
                                                                                        var v121 uint16 = 605
                                                                                        _ = v121
                                                                                    default:
                                                                                        v27 = 32767
                                                                                    }
                                                                                    _ = v120
                                                                                }
                                                                                v43 -= (v43 % (((v43 ^ v88)) | 1))
                                                                                v3 = 1
                                                                                if (664 <= 1) {
                                                                                    var v122 string = ""
                                                                                    switch (v43 ^ (^(v88))) {
                                                                                    case 1:
                                                                                        v87 -= 1
                                                                                        var v123 uint16 = 464
                                                                                        _ = v123
                                                                                    case 712:
                                                                                        v2[(-(v27))] = v27
                                                                                    case 2147483647:
                                                                                        if false {
                                                                                            if ((v27 == 128) && true) {
                                                                                                if false {
                                                                                                    v88 = (v88 >> ((42) & 63))
                                                                                                    var v124 int64 = 450
                                                                                                    _ = v124
                                                                                                } else {
                                                                                                    delete(v2, (fi38 | 4294967295))
                                                                                                    v125 := GenS9{F0: 4294967295, F1: v122}
                                                                                                    _ = v125
                                                                                                }
                                                                                                var v126 bool = false
                                                                                                _ = v126
                                                                                            }
                                                                                            switch -488 {
                                                                                            case 4294967295:
                                                                                                if false {
                                                                                                    v43 += fi38
                                                                                                    v127 := GenS7{F0: 375, F1: -32768, F2: (v27 + (^(fi39)))}
                                                                                                    _ = v127
                                                                                                }
                                                                                            case 32767:
                                                                                                switch (fi38 ^ fi38) {
                                                                                                case 114:
                                                                                                    fmt.Println(fi38, fi38)
                                                                                                case 9223372036854775807:
                                                                                                    delete(v2, 4294967295)
                                                                                                    v2[fi39] = (^(v27))
                                                                                                    switch v122 {
                                                                                                    case "hello":
                                                                                                        v87 += v87
                                                                                                    default:
                                                                                                        if len(v1) > 0 {
                                                                                                            v1[0] = ("hello" + v122)
                                                                                                        }
                                                                                                    }
                                                                                                case -9223372036854775808:
                                                                                                    if false {
                                                                                                        switch 679 {
                                                                                                        case 1:
                                                                                                            var v128 bool = false
                                                                                                            _ = v128
                                                                                                        default:
                                                                                                            v129 := GenS5{F0: 76, F1: 65535, F2: v87}
                                                                                                            _ = v129
                                                                                                        }
                                                                                                    } else {
                                                                                                        switch -306 {
                                                                                                        case -9223372036854775808:
                                                                                                            v43 = (fi39 / ((0) | 1))
                                                                                                        case -515:
                                                                                                            v45 = 0.1
                                                                                                        case 1:
                                                                                                            v130 := v2[9223372036854775807]
                                                                                                            _ = v130
                                                                                                        default:
                                                                                                            var v131 uint8 = 113
                                                                                                            _ = v131
                                                                                                        }
                                                                                                        v132 := GenS1{F0: 1}
                                                                                                        _ = v132
                                                                                                    }
                                                                                                default:
                                                                                                    delete(v2, -2147483648)
                                                                                                }
                                                                                            case 882:
                                                                                                fmt.Println(v45)
                                                                                            default:
                                                                                                if true {
                                                                                                    fmt.Println(v43, fi38, v45)
                                                                                                    v87 += (v87 >> ((genFn0()) & 63))
                                                                                                } else {
                                                                                                    v1 = append(v1, "a")
                                                                                                    if len(v1) >= 2 {
                                                                                                        v1 = v1[1:len(v1)]
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            var v133 uint16 = 493
                                                                                            if (255 <= 86) {
                                                                                                _ = len(v2)
                                                                                                _ = len(v1)
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = "hello"
                                                                                                }
                                                                                            } else {
                                                                                                var v134 uint32 = 2147483648
                                                                                                fmt.Println(v27, v88, v45)
                                                                                                var v135 uint64 = uint64(v134)
                                                                                                switch v122 {
                                                                                                case "a":
                                                                                                    var v136 uint32 = v134
                                                                                                    _ = v136
                                                                                                default:
                                                                                                    switch v122 {
                                                                                                    case "foo bar":
                                                                                                        v87 -= v87
                                                                                                    default:
                                                                                                        var v137 string = ("foo bar" + v122)
                                                                                                        _ = v137
                                                                                                    }
                                                                                                }
                                                                                                _ = v134
                                                                                                _ = v135
                                                                                            }
                                                                                            var v138 uint = 311
                                                                                            v133 -= v133
                                                                                            v122 = "\n\t"
                                                                                            _ = v133
                                                                                            _ = v138
                                                                                        }
                                                                                        v3 -= v3
                                                                                        v122 = ((v122 + ("" + (("hello" + ("" + v122)) + ((v122 + "") + "")))) + v122)
                                                                                    }
                                                                                    delete(v2, 9223372036854775807)
                                                                                    _ = v122
                                                                                } else {
                                                                                    if (true && true) {
                                                                                        fmt.Println(v45, fi39)
                                                                                        v88 = 162
                                                                                    } else {
                                                                                        if (!false) {
                                                                                            var v139 uint32 = 2147483648
                                                                                            _ = v139
                                                                                        } else {
                                                                                            v43 = 1
                                                                                            switch v27 {
                                                                                            case 4294967295:
                                                                                                switch "\n\t" {
                                                                                                case "a":
                                                                                                    var v140 uint32 = 57
                                                                                                    v87 = v87
                                                                                                    _ = v140
                                                                                                case "foo bar":
                                                                                                    fmt.Println(fi39, v87, v43)
                                                                                                }
                                                                                            case 128:
                                                                                                if ((!(!true)) || ((0.0 - 0.1) < ((-2.25 - 1e10) + float32((v45 * 0.1))))) {
                                                                                                    fmt.Println(v3, v3)
                                                                                                    v3 += (-(v3))
                                                                                                    v141 := GenS13{F0: (!(((!true) || ("" > ("foo bar" + "hello"))) && (103 <= 13))), F1: 1}
                                                                                                    _ = v141
                                                                                                }
                                                                                            default:
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                            }
                                                                                            v45 += -2.25
                                                                                        }
                                                                                        v142 := GenS14{F0: v87}
                                                                                        v87 = (^(v87))
                                                                                        _ = v142
                                                                                    }
                                                                                    var v143 bool = false
                                                                                    v3 -= 42
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        v87 = -2147483648
                                                                                    default:
                                                                                        var v144 float32 = 0.0
                                                                                        _ = v144
                                                                                    }
                                                                                    _ = v143
                                                                                }
                                                                                v145 := GenS15{F0: (fi38 | (^(v88)))}
                                                                                _ = v111
                                                                                _ = v145
                                                                            }
                                                                            _ = v92
                                                                        }
                                                                        if true {
                                                                            if ((v87 << ((863) & 63)) == v87) {
                                                                                switch 614 {
                                                                                case -865:
                                                                                    var v146 float64 = -2.25
                                                                                    _ = v146
                                                                                case 248:
                                                                                    switch (v88 ^ v88) {
                                                                                    case -9223372036854775808:
                                                                                        switch (v88 / (((v27 ^ -2147483648)) | 1)) {
                                                                                        case -9223372036854775808:
                                                                                            switch (fi38 / ((v88) | 1)) {
                                                                                            case 32768:
                                                                                                if (((float64((3.14159 / 3.14159)) + (v45 + -2.25)) + 0.0) >= (-2.25 - v45)) {
                                                                                                    if (!(!((32767 < 65535) || true))) {
                                                                                                        var v147 uint64 = 1
                                                                                                        switch "foo bar" {
                                                                                                        case "a":
                                                                                                            var v148 uint = 869
                                                                                                            _ = v148
                                                                                                        default:
                                                                                                            v149 := GenS16{F0: 2}
                                                                                                            fmt.Println(v87, v147, v3)
                                                                                                            _ = v149
                                                                                                        }
                                                                                                        switch "" {
                                                                                                        case "\n\t":
                                                                                                            v150 := GenS17{F0: "\n\t", F1: 148}
                                                                                                            _ = v150
                                                                                                        case "foo bar":
                                                                                                            var v151 int64 = 42
                                                                                                            _ = v151
                                                                                                        }
                                                                                                        var v152 uint8 = 128
                                                                                                        v153 := GenS18{F0: (^(v3)), F1: v152, F2: (v87 & (^(v87)))}
                                                                                                        _ = v147
                                                                                                        _ = v152
                                                                                                        _ = v153
                                                                                                    } else {
                                                                                                        var v154 uint16 = 255
                                                                                                        v154 -= v154
                                                                                                        v154 = v154
                                                                                                        _ = v154
                                                                                                    }
                                                                                                }
                                                                                            case 42:
                                                                                                _ = len(v2)
                                                                                                v155 := GenS10{F0: (v87 - 2147483647), F1: 42, F2: (v3 ^ (v3 + v3))}
                                                                                                v27 -= (fi39 & fi39)
                                                                                                _ = v155
                                                                                            case 652:
                                                                                                fmt.Println(v45, v43, v3)
                                                                                            }
                                                                                        }
                                                                                    case 808:
                                                                                        if false {
                                                                                            if (false && false) {
                                                                                                v88 += (^(fi38))
                                                                                                v45 -= float64(1.5)
                                                                                                fmt.Println(fi39, v27)
                                                                                            } else {
                                                                                                _ = len(v1)
                                                                                                if true {
                                                                                                    var v156 uint8 = 42
                                                                                                    var v157 int8 = 100
                                                                                                    if len(v1) > 0 {
                                                                                                        v1[0] = "\n\t"
                                                                                                    }
                                                                                                    _ = v156
                                                                                                    _ = v157
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            switch ("foo bar" + "foo bar") {
                                                                                            case "":
                                                                                                v3 = v3
                                                                                            case "a":
                                                                                                if false {
                                                                                                    if true {
                                                                                                        v27 += (^(v43))
                                                                                                        fmt.Println(v27, v87, v45)
                                                                                                    }
                                                                                                    if true {
                                                                                                        var v158 bool = false
                                                                                                        _ = v158
                                                                                                    } else {
                                                                                                        v159 := GenS3{F0: -1}
                                                                                                        _ = v159
                                                                                                    }
                                                                                                    switch (v88 & (v27 & v43)) {
                                                                                                    case -9223372036854775808:
                                                                                                        if false {
                                                                                                            v1 = append(v1, "hello")
                                                                                                        } else {
                                                                                                            if (-25 > 127) {
                                                                                                                var v160 int = 379
                                                                                                                var v161 uint8 = 1
                                                                                                                v162 := GenS10{F0: v87, F1: v3, F2: 32767}
                                                                                                                fmt.Println(fi38, v27)
                                                                                                                _ = v160
                                                                                                                _ = v161
                                                                                                                _ = v162
                                                                                                            } else {
                                                                                                                v27 = (v43 >> ((82) & 63))
                                                                                                                v163 := GenS16{F0: 170}
                                                                                                                _ = v163
                                                                                                            }
                                                                                                            v88 = fi39
                                                                                                        }
                                                                                                    case -709:
                                                                                                        var v164 int = (fi39 & (v27 + (fi39 - (fi38 + fi39))))
                                                                                                        _ = v164
                                                                                                    default:
                                                                                                        v27 = v43
                                                                                                    }
                                                                                                }
                                                                                            case "foo bar":
                                                                                                v2[v27] = 32767
                                                                                            default:
                                                                                                v87 -= (v87 - -860)
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = "foo bar"
                                                                                                }
                                                                                            }
                                                                                            var v165 int8 = -103
                                                                                            _ = v165
                                                                                        }
                                                                                    case -1:
                                                                                        if ("\n\t" < ("a" + ("" + "\n\t"))) {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = "foo bar"
                                                                                            }
                                                                                            switch "\n\t" {
                                                                                            case "hello":
                                                                                                var v166 int64 = -9223372036854775808
                                                                                                _ = len(v1)
                                                                                                _ = len(v2)
                                                                                                _ = v166
                                                                                            case "foo bar":
                                                                                                fmt.Println(v27)
                                                                                            default:
                                                                                                if (!(-2.25 >= ((1e10 / -2.25) - 1.5))) {
                                                                                                    _ = len(v2)
                                                                                                    fmt.Println(fi38, v3, v87)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    var v167 int = v43
                                                                                    _ = v167
                                                                                }
                                                                                if (127 == 127) {
                                                                                    switch v43 {
                                                                                    case -1:
                                                                                        var v168 int8 = 1
                                                                                        _ = v168
                                                                                    case 32767:
                                                                                        var v169 int8 = 42
                                                                                        _ = v169
                                                                                    default:
                                                                                        v3 += v3
                                                                                    }
                                                                                    if false {
                                                                                        var v170 int = fi39
                                                                                        _ = v170
                                                                                    }
                                                                                }
                                                                                v45 = (((v45 / 1.5) - 0.1) * 0.0)
                                                                                var v171 uint = 42
                                                                                _ = v171
                                                                            }
                                                                        } else {
                                                                            var v172 float64 = (v45 + float64(1e10))
                                                                            _ = v172
                                                                        }
                                                                    }
                                                                }
                                                            default:
                                                                v173 := GenS17{F0: "", F1: 156}
                                                                _ = v173
                                                            }
                                                        }
                                                    } else {
                                                        v1 = append(v1, ("" + "hello"))
                                                        switch ("a" + "a") {
                                                        case "\n\t":
                                                            v27 -= (fi38 - -1)
                                                        case "foo bar":
                                                            var v174 int8 = -128
                                                            v175 := GenS19{F0: 9223372036854775807, F1: v87, F2: ((!false) && (2 != 2))}
                                                            delete(v2, (fi38 % (((fi38 * (fi39 % (((fi38 / (((fi38 >> ((4294967295) & 63))) | 1))) | 1)))) | 1)))
                                                            _ = v174
                                                            _ = v175
                                                        case "a":
                                                            v27 -= fi38
                                                        default:
                                                            v176 := GenS16{F0: 16}
                                                            _ = v176
                                                        }
                                                        switch "hello" {
                                                        case "\n\t":
                                                            var v177 uint = genFn0()
                                                            _ = v177
                                                        case "hello":
                                                            var v178 int32 = (^(v87))
                                                            _ = v178
                                                        default:
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                    }
                                                    _ = v87
                                                    _ = v88
                                                }
                                            case "":
                                                switch (-(v27)) {
                                                case -361:
                                                    v43 = fi39
                                                default:
                                                    var v179 string = "\n\t"
                                                    _ = v179
                                                }
                                            default:
                                                var v180 string = ("foo bar" + ("foo bar" + "foo bar"))
                                                _ = v180
                                            }
                                            if true {
                                                v43 += (fi38 & (fi39 % (((v43 % ((v27) | 1))) | 1)))
                                            } else {
                                                v181 := GenS17{F0: "a", F1: 78}
                                                v1 = append(v1, "a")
                                                var v182 bool = true
                                                _ = v181
                                                _ = v182
                                            }
                                        } else {
                                            fmt.Println(fi38)
                                            v27 = (v43 >> ((1) & 63))
                                        }
                                        delete(v2, -2147483648)
                                        _ = v44
                                        _ = v45
                                    }
                                    var v183 uint = 57
                                    v3 = v3
                                    if len(v1) > 0 {
                                        v1[0] = "foo bar"
                                    }
                                    _ = v183
                                    _ = v43
                                } else {
                                    var v184 uint8 = 2
                                    _ = v184
                                }
                                var v185 uint64 = 4294967295
                                v3 += (v3 | (-(v3)))
                                _ = v185
                            }
                            v3 += v3
                            var v186 int = (-(fi39))
                            _ = fi39
                            _ = v186
                        }
                        _ = fi38
                    }
                }
                if (61 < 255) {
                    var v187 int16 = -1
                    _ = v187
                }
            } else {
                v188 := GenS20{F0: 1, F1: 1}
                for fi189 := 0; fi189 < 9; fi189++ {
                    for fi190 := 0; fi190 < 16; fi190++ {
                        v3 = v3
                        delete(v2, 422)
                        if (fi189 > (v27 + (-(fi189)))) {
                            if false {
                                v27 += 1
                                fmt.Println(fi190, v3)
                            } else {
                                v191 := GenS17{F0: "foo bar", F1: 255}
                                fmt.Println(v27, fi190)
                                v3 = v3
                                _ = v191
                            }
                            if true {
                                fmt.Println(v27)
                                var v192 int8 = 127
                                _ = v192
                            } else {
                                var v193 uint32 = 2
                                if len(v1) > 0 {
                                    v1[0] = "a"
                                }
                                fmt.Println(fi190)
                                _ = v193
                            }
                            fmt.Println(v27, fi189, v3)
                        }
                        _ = fi190
                    }
                    var v194 uint16 = 42
                    _ = fi189
                    _ = v194
                }
                v3 = (v3 & v3)
                _ = v188
            }
            if true {
                v27 -= 494
            }
        } else {
            v195 := GenS21{F0: v3}
            _ = v195
        }
    }
    switch (("foo bar" + "foo bar") + "foo bar") {
    case "foo bar":
        for fi196 := 0; fi196 < 6; fi196++ {
            if true {
                delete(v2, fi196)
                var v197 uint8 = 127
                v197 += v197
                _ = v197
            }
            v198 := GenS9{F0: 0, F1: (("hello" + ((("foo bar" + "foo bar") + (("a" + ("foo bar" + ("" + "foo bar"))) + ("\n\t" + "a"))) + "")) + "hello")}
            for fi199 := 0; fi199 < 20; fi199++ {
                if true {
                    v200 := GenS22{F0: v27, F1: -9223372036854775808}
                    if true {
                        v1 = append(v1, "a")
                        _ = len(v2)
                    } else {
                        switch v27 {
                        case -795:
                            if (1e10 != (float32((1e10 - 1.5)) + (((float32((((-2.25 - -2.25) * float64(float32(((1.5 - float64(float32(3.14159))) - 3.14159)))) - 0.1)) * 1e10) / 1.0) * (0.1 * ((-2.25 - float32(0.1)) - 3.14159))))) {
                                var v201 uint = genFn0()
                                switch -180 {
                                case 930:
                                    var v202 bool = (42 > 0)
                                    _ = v202
                                default:
                                    v3 = 1
                                }
                                _ = v201
                            } else {
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v3, v27, fi196)
                                case "hello":
                                    v27 -= v27
                                }
                                v27 -= -2147483648
                                fmt.Println(fi199, v27, fi199)
                                fmt.Println(fi196)
                            }
                        case 32767:
                            var v203 uint16 = 546
                            _ = v203
                        case 2147483647:
                            if (!(65535 == 42)) {
                                switch 148 {
                                case 128:
                                    var v204 int8 = -1
                                    _ = v204
                                case -1:
                                    if false {
                                        var v205 float32 = 3.14159
                                        v205 = float32(1e10)
                                        _ = v205
                                    } else {
                                        v2[(fi196 % (((v27 % ((-9223372036854775808) | 1))) | 1))] = fi196
                                        if true {
                                            switch fi196 {
                                            case -738:
                                                v1 = append(v1, "")
                                            case 4294967295:
                                                delete(v2, 4294967295)
                                            case -109:
                                                if len(v1) > 0 {
                                                    v1[0] = ""
                                                }
                                            }
                                            v27 = fi196
                                        }
                                        switch v27 {
                                        case 127:
                                            v2[(-(fi196))] = v27
                                        default:
                                            var v206 int8 = 127
                                            _ = v206
                                        }
                                    }
                                }
                                fmt.Println(v3)
                                v3 = (v3 / ((0) | 1))
                            }
                        default:
                            switch "" {
                            case "a":
                                var v207 int64 = 0
                                _ = v207
                            case "foo bar":
                                v3 -= v3
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                    }
                    v208 := GenS23{F0: false, F1: 757}
                    v3 = (v3 | -1)
                    if (!false) {
                        switch (^(fi199)) {
                        case -866:
                            v27 -= fi199
                        case -9223372036854775808:
                            var v209 int64 = 42
                            _ = v209
                        case -363:
                            if false {
                                v210 := GenS24{F0: 0, F1: 2}
                                switch v27 {
                                case -166:
                                    if false {
                                        v27 -= (^(fi199))
                                        if ((42 > 2) && (!true)) {
                                            v3 = v3
                                            v27 = (v27 / ((fi199) | 1))
                                            v3 = v3
                                        }
                                    }
                                case 32767:
                                    if len(v1) > 0 {
                                        v1[0] = "\n\t"
                                    }
                                }
                                if false {
                                    if (4294967295 > 807) {
                                        switch -2147483648 {
                                        case 671:
                                            fmt.Println(fi199, v27, fi199)
                                        case 22:
                                            var v211 float32 = 1e10
                                            _ = v211
                                        case -9223372036854775808:
                                            if (v3 == v3) {
                                                switch fi199 {
                                                case 9223372036854775807:
                                                    v27 += fi196
                                                default:
                                                    v3 -= (v3 >> ((235) & 63))
                                                }
                                            } else {
                                                fmt.Println(v27, v27, v3)
                                            }
                                        }
                                    }
                                }
                                _ = v210
                            } else {
                                switch fi199 {
                                case 9223372036854775807:
                                    v3 += v3
                                case -2147483648:
                                    v3 -= 1
                                case -1:
                                    fmt.Println(fi196, fi196)
                                }
                                if false {
                                    var v212 uint64 = 398
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                    fmt.Println(v3, v27, v27)
                                    _ = v212
                                }
                                v27 += (fi196 / (((v27 >> ((4294967295) & 63))) | 1))
                                var v213 uint64 = 42
                                _ = v213
                            }
                        default:
                            var v214 uint64 = 742
                            if true {
                                var v215 float64 = float64(((0.0 / (float32((-2.25 / -2.25)) * 0.1)) - 3.14159))
                                var v216 uint64 = v214
                                var v217 string = "a"
                                v216 -= 443
                                _ = v215
                                _ = v216
                                _ = v217
                            }
                            var v218 uint = 65535
                            _ = v214
                            _ = v218
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v27 += v27
                        switch ("\n\t" + "a") {
                        case "foo bar":
                            if false {
                                v219 := GenS25{F0: 940}
                                v27 += (v27 / ((fi199) | 1))
                                _ = v219
                            }
                        }
                    } else {
                        var v220 uint32 = 322
                        _ = v220
                    }
                    _ = v200
                    _ = v208
                }
                v221 := GenS9{F0: 65535, F1: "\n\t"}
                var v222 string = ("\n\t" + "\n\t")
                var v223 string = v222
                var v224 uint16 = 0
                _ = fi199
                _ = v221
                _ = v222
                _ = v223
                _ = v224
            }
            _ = fi196
            _ = v198
        }
    case "":
        v225 := GenS26{F0: 42, F1: -1}
        _ = v225
    default:
        if (v27 != v27) {
            if false {
                v1 = append(v1, "")
                v1 = append(v1, "")
            }
        } else {
            v226 := GenS4{F0: "a", F1: 1}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if ((1.5 / (0.0 + float32(1.5))) != 1.5) {
                fmt.Println(v3, v3)
                for fi227 := 0; fi227 < 17; fi227++ {
                    v27 += fi227
                    var v228 int = fi227
                    switch 479 {
                    case 0:
                        if (!false) {
                            var v229 uint8 = 2
                            for fi230 := 0; fi230 < 18; fi230++ {
                                if true {
                                    if len(v1) > 0 {
                                        v1[0] = ("\n\t" + "")
                                    }
                                    if (231 != 737) {
                                        v231 := GenS27{F0: 65535}
                                        var v232 float64 = 3.14159
                                        _ = v231
                                        _ = v232
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v233 := GenS28{F0: 425, F1: -104, F2: 32767}
                                        v229 += v229
                                        v234 := GenS29{F0: 42, F1: 815}
                                        _ = v233
                                        _ = v234
                                    }
                                }
                                var v235 uint8 = 113
                                v229 += v235
                                _ = fi230
                                _ = v235
                            }
                            v3 += (-(v3))
                            _ = v229
                        } else {
                            v228 -= 127
                            _ = len(v2)
                            for fi236 := 0; fi236 < 6; fi236++ {
                                var v237 uint32 = 204
                                var v238 uint32 = 24
                                _ = fi236
                                _ = v237
                                _ = v238
                            }
                        }
                    default:
                        for fi239 := 0; fi239 < 1; fi239++ {
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                            switch "" {
                            case "a":
                                var v240 int8 = 127
                                _ = v240
                            case "":
                                v3 = (v3 & (-(v3)))
                            case "hello":
                                v2[(fi239 << ((2) & 63))] = 9223372036854775807
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v241 int = fi227
                            _ = fi239
                            _ = v241
                        }
                    }
                    _ = fi227
                    _ = v228
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                if true {
                    for fi242 := 0; fi242 < 13; fi242++ {
                        fmt.Println(fi242, v3, v3)
                        _ = fi242
                    }
                    v243 := GenS4{F0: ("\n\t" + "\n\t"), F1: -1}
                    _ = v243
                }
                if true {
                    for fi244 := 0; fi244 < 16; fi244++ {
                        v245 := GenS24{F0: 2, F1: 42}
                        if true {
                            var v246 uint32 = 314
                            for fi247 := 0; fi247 < 19; fi247++ {
                                if false {
                                    switch "" {
                                    case "\n\t":
                                        v246 += 780
                                    case "":
                                        if ((0.0 * (-2.25 / 1.0)) > (1e10 - (3.14159 + (3.14159 / 1e10)))) {
                                            var v248 bool = false
                                            switch "hello" {
                                            case "a":
                                                v249 := GenS7{F0: (v246 % ((2147483648) | 1)), F1: (v3 ^ v3), F2: v27}
                                                _ = v249
                                            case "hello":
                                                var v250 bool = v248
                                                _ = v250
                                            case "":
                                                if (18446744073709551615 > 573) {
                                                    v248 = (0 >= 1)
                                                    var v251 float64 = 1e10
                                                    if false {
                                                        var v252 int16 = v3
                                                        switch "" {
                                                        case "hello":
                                                            if v248 {
                                                                var v253 int = (fi247 % ((v27) | 1))
                                                                _ = v253
                                                            } else {
                                                                if true {
                                                                    var v254 int16 = 1
                                                                    if v248 {
                                                                        v248 = false
                                                                        if v248 {
                                                                            v27 -= (fi247 & v27)
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                _ = len(v1)
                                                                            case "hello":
                                                                                var v255 float32 = 1e10
                                                                                _ = v255
                                                                            default:
                                                                                var v256 uint16 = 255
                                                                                _ = v256
                                                                            }
                                                                            var v257 int64 = -9223372036854775808
                                                                            _ = v257
                                                                        }
                                                                    }
                                                                    _ = v254
                                                                } else {
                                                                    fmt.Println(v252)
                                                                    v258 := GenS30{F0: 42, F1: 279, F2: 1}
                                                                    v259 := GenS28{F0: 180, F1: (v3 * v3), F2: 0}
                                                                    _ = v258
                                                                    _ = v259
                                                                }
                                                            }
                                                        case "foo bar":
                                                            v246 = v246
                                                            v260 := GenS31{F0: 0, F1: (v252 >> ((4294967295) & 63)), F2: "a"}
                                                            v246 = 2147483648
                                                            _ = v260
                                                        default:
                                                            var v261 float32 = 0.1
                                                            _ = v261
                                                        }
                                                        _ = len(v1)
                                                        _ = v252
                                                    } else {
                                                        delete(v2, fi247)
                                                        var v262 string = "hello"
                                                        if false {
                                                            switch "hello" {
                                                            case "a":
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            case "\n\t":
                                                                switch 32768 {
                                                                case 4294967295:
                                                                    if v248 {
                                                                        var v263 float64 = 1.5
                                                                        v264 := GenS9{F0: genFn0(), F1: "foo bar"}
                                                                        fmt.Println(v27)
                                                                        _ = v263
                                                                        _ = v264
                                                                    } else {
                                                                        v262 = v262
                                                                        fmt.Println(fi244, v3, v3)
                                                                    }
                                                                case 0:
                                                                    if (4294967295 >= v27) {
                                                                        if v248 {
                                                                            var v265 uint = 89
                                                                            _ = v265
                                                                        }
                                                                    }
                                                                default:
                                                                    var v266 uint64 = 18446744073709551615
                                                                    _ = v266
                                                                }
                                                            default:
                                                                if len(v1) > 0 {
                                                                    v1[0] = "hello"
                                                                }
                                                            }
                                                            v267 := GenS10{F0: 622, F1: (v3 - (v3 % ((v3) | 1))), F2: v3}
                                                            v268 := GenS1{F0: 2}
                                                            _ = v267
                                                            _ = v268
                                                        } else {
                                                            v251 = float64(float32(((v251 + float64(((0.1 - 1e10) + (-2.25 + 3.14159)))) * ((3.14159 - 0.0) + 1e10))))
                                                            v269 := GenS32{F0: 42, F1: 387}
                                                            var v270 int16 = -1
                                                            _ = v269
                                                            _ = v270
                                                        }
                                                        _ = v262
                                                    }
                                                    _ = v251
                                                } else {
                                                    var v271 int = v27
                                                    v246 -= v246
                                                    _ = v271
                                                }
                                            default:
                                                if (!v248) {
                                                    switch "" {
                                                    case "a":
                                                        v27 += (fi244 << ((530) & 63))
                                                    case "":
                                                        fmt.Println(v27)
                                                    case "foo bar":
                                                        var v272 uint64 = 135
                                                        _ = v272
                                                    default:
                                                        var v273 float32 = float32(1.5)
                                                        _ = v273
                                                    }
                                                    if v248 {
                                                        if (0.1 > (1e10 / (0.1 * (3.14159 * (float32(float64(((0.0 - 1.5) - (1e10 + -2.25)))) * 1e10))))) {
                                                            var v274 uint8 = 15
                                                            v248 = true
                                                            _ = v274
                                                        }
                                                    } else {
                                                        if v248 {
                                                            if len(v1) > 0 {
                                                                v1[0] = "a"
                                                            }
                                                            v246 -= (^(v246))
                                                        }
                                                    }
                                                    v3 = v3
                                                } else {
                                                    v248 = (v248 && v248)
                                                    v3 = (v3 << ((0) & 63))
                                                }
                                            }
                                            _ = v248
                                        }
                                    case "foo bar":
                                        v246 -= (v246 * (^(v246)))
                                    default:
                                        if false {
                                            var v275 int64 = 149
                                            var v276 uint8 = 0
                                            _ = v275
                                            _ = v276
                                        } else {
                                            v27 = fi244
                                            v246 = v246
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    }
                                }
                                v277 := GenS25{F0: 881}
                                v278 := v2[(v27 % ((int(v3)) | 1))]
                                _ = fi247
                                _ = v277
                                _ = v278
                            }
                            _ = v246
                        }
                        _ = fi244
                        _ = v245
                    }
                    if ("hello" >= ("foo bar" + "\n\t")) {
                        v27 -= 32768
                    }
                    if (!false) {
                        _ = len(v2)
                    } else {
                        var v279 int16 = (v3 + (v3 & v3))
                        _ = v279
                    }
                } else {
                    switch (("a" + ((("a" + ((("hello" + (("\n\t" + "foo bar") + "a")) + "a") + "\n\t")) + ("a" + ("\n\t" + ("a" + "a")))) + (("a" + "a") + "\n\t"))) + ("foo bar" + "foo bar")) {
                    case "foo bar":
                        var v280 string = "\n\t"
                        _ = v280
                    default:
                        v3 += (v3 | v3)
                    }
                }
            }
            _ = v226
        }
    }
    var v281 int8 = -107
    if ((v281 << ((0) & 63)) > v281) {
        var v282 uint64 = 924
        v283 := GenS33{F0: -948}
        var v284 string = "\n\t"
        _ = v282
        _ = v283
        _ = v284
    }
    _ = len(v1)
    v3 += v3
    v27 = -129
    if (false && true) {
        for fi285 := 0; fi285 < 1; fi285++ {
            var v286 int8 = v281
            for fi287 := 0; fi287 < 6; fi287++ {
                if len(v1) > 0 {
                    v1[0] = ("" + (("hello" + "foo bar") + "a"))
                }
                v288 := GenS23{F0: false, F1: 42}
                var v289 int = (v27 * 127)
                if ((v286 | 94) > v286) {
                    switch 71 {
                    case -211:
                        var v290 float32 = (3.14159 + (float32((0.1 - 1.5)) + 1e10))
                        _ = v290
                    }
                }
                delete(v2, v27)
                _ = fi287
                _ = v288
                _ = v289
            }
            v291 := GenS28{F0: 2147483647, F1: 156, F2: 656}
            _ = fi285
            _ = v286
            _ = v291
        }
        v281 += (v281 + v281)
    }
    for fi292 := 0; fi292 < 17; fi292++ {
        fmt.Println(v3, v27)
        var v293 int64 = -510
        var v294 uint16 = 504
        _ = fi292
        _ = v293
        _ = v294
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v281:", v281)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v281
    _ = v3
}
