package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int16; F1 string; F2 int }
type GenS9 struct { F0 int64; F1 bool; F2 uint }
type GenS10 struct { F0 uint; F1 bool }
type GenS11 struct { F0 uint8; F1 uint64 }
type GenS12 struct { F0 int32; F1 int; F2 bool }
type GenS13 struct { F0 int8; F1 bool; F2 uint }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint; F1 uint16 }
type GenS16 struct { F0 uint16; F1 uint32; F2 int64 }
func genFn0(p0 int16, p1 int32) int8 {
    v5 := []int{2147483647, -1}
    if false {
        var v6 uint = 810
        _ = v6
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = v5
    return -128
}

func main() {
    var v1 uint16 = 608
    var v2 float32 = (((float32(((0.1 * 1.5) * (0.1 + (3.14159 / 1e10)))) / -2.25) + (3.14159 * -2.25)) + (1e10 / -2.25))
    v3 := []string{("a" + ""), "\n\t"}
    var v4 float64 = float64(-2.25)
    v3 = append(v3, "a")
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int8 = 127
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(v2, v1, v4)
            var v10 uint32 = 2147483648
            _ = fi9
            _ = v10
        }
        switch "" {
        case "":
            v2 = 1e10
        case "\n\t":
            var v11 int = fi7
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    v3 = append(v3, ("" + "foo bar"))
    v1 -= 2
    var v12 float64 = v4
    for fi13 := 0; fi13 < 6; fi13++ {
        fmt.Println(fi13, v2)
        if true {
            v14 := GenS1{F0: 4294967295, F1: (^(fi13))}
            fmt.Println(fi13)
            v4 += 1.5
            _ = v14
        } else {
            v4 += (1.5 + 1.5)
            v4 += 0.0
        }
        _ = fi13
    }
    switch 32767 {
    case 127:
        if false {
            if false {
                for fi15 := 0; fi15 < 5; fi15++ {
                    _ = len(v3)
                    _ = fi15
                }
                switch 42 {
                case 1:
                    for fi16 := 0; fi16 < 19; fi16++ {
                        for fi17 := 0; fi17 < 4; fi17++ {
                            v4 += 1e10
                            v2 = v2
                            fmt.Println(v2)
                            _ = fi17
                        }
                        if false {
                            v3 = append(v3, "hello")
                            for fi18 := 0; fi18 < 4; fi18++ {
                                if true {
                                    var v19 float32 = 1e10
                                    _ = v19
                                }
                                _ = fi18
                            }
                            v20 := GenS1{F0: 42, F1: (fi16 << ((61) & 63))}
                            _ = v20
                        }
                        _ = fi16
                    }
                case 9223372036854775807:
                    if true {
                        for fi21 := 0; fi21 < 6; fi21++ {
                            if ((-326 <= (fi21 & 9223372036854775807)) || (!(v4 != (v4 * 3.14159)))) {
                                for fi22 := 0; fi22 < 7; fi22++ {
                                    switch fi21 {
                                    case 32768:
                                        v23 := GenS2{F0: -470, F1: ""}
                                        _ = v23
                                    case 922:
                                        if (!("hello" > "\n\t")) {
                                            v12 = (0.1 - v4)
                                            if (!false) {
                                                v24 := GenS3{F0: -1, F1: 42}
                                                _ = v24
                                            } else {
                                                fmt.Println(fi22, v4)
                                                var v25 float64 = (((v12 * -2.25) * (v12 + 3.14159)) - (1.5 + v4))
                                                if true {
                                                    v2 -= (3.14159 + ((v2 - 0.1) + 1e10))
                                                }
                                                _ = v25
                                            }
                                            var v26 int16 = -32768
                                            _ = v26
                                        } else {
                                            v1 -= v1
                                            v27 := GenS2{F0: -1, F1: "a"}
                                            if (!true) {
                                                var v28 uint32 = 2147483648
                                                _ = v28
                                            }
                                            _ = v27
                                        }
                                    default:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    v29 := GenS4{F0: (!((float32((v4 - 1.5)) > 3.14159) || false))}
                                    v30 := GenS5{F0: 4294967295}
                                    _ = fi22
                                    _ = v29
                                    _ = v30
                                }
                                for fi31 := 0; fi31 < 9; fi31++ {
                                    var v32 float64 = v4
                                    if ((2 >= 42) || true) {
                                        var v33 uint64 = 193
                                        if (127 < 127) {
                                            var v34 uint8 = 127
                                            var v35 uint32 = 2
                                            fmt.Println(v2, v35, v33)
                                            _ = v34
                                            _ = v35
                                        }
                                        v36 := GenS2{F0: -1, F1: (("hello" + "hello") + "\n\t")}
                                        _ = v33
                                        _ = v36
                                    }
                                    if ((!(-9223372036854775808 <= 447)) && ("foo bar" < "")) {
                                        v37 := GenS6{F0: 255, F1: -89}
                                        if false {
                                            v38 := GenS1{F0: 65535, F1: 691}
                                            _ = v38
                                        } else {
                                            var v39 string = ("hello" + "hello")
                                            v4 += v32
                                            v12 -= 0.1
                                            _ = v39
                                        }
                                        _ = v37
                                    } else {
                                        v4 = (0.0 * (v12 * 1e10))
                                        _ = len(v3)
                                        var v40 uint64 = 42
                                        switch "hello" {
                                        case "hello":
                                            switch "hello" {
                                            case "foo bar":
                                                fmt.Println(fi31, v12, v12)
                                            default:
                                                var v41 float32 = (v2 / 3.14159)
                                                _ = v41
                                            }
                                        case "":
                                            fmt.Println(fi21, v12)
                                        case "\n\t":
                                            switch (fi21 & fi21) {
                                            case 0:
                                                v2 = v2
                                            case 32767:
                                                fmt.Println(fi31)
                                            }
                                        default:
                                            var v42 int = fi21
                                            _ = v42
                                        }
                                        _ = v40
                                    }
                                    v12 += v32
                                    v43 := GenS5{F0: 71}
                                    var v44 float64 = (((-2.25 / float64(v2)) * v4) + (v12 + ((float64(float32((0.0 / v12))) / v32) * -2.25)))
                                    var v45 uint = 1
                                    var v46 int16 = 42
                                    _ = fi31
                                    _ = v32
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                }
                            }
                            v47 := GenS4{F0: false}
                            fmt.Println(v2)
                            _ = fi21
                            _ = v47
                        }
                    } else {
                        v48 := GenS4{F0: ((v4 + 1e10) >= ((0.1 / (0.0 - (0.0 + float64(float32(v4))))) + 3.14159))}
                        _ = v48
                    }
                case 127:
                    for fi49 := 0; fi49 < 17; fi49++ {
                        switch 2147483647 {
                        case 32768:
                            for fi50 := 0; fi50 < 7; fi50++ {
                                fmt.Println(fi49, v2)
                                v12 += 1e10
                                _ = fi50
                            }
                        default:
                            for fi51 := 0; fi51 < 10; fi51++ {
                                v3 = append(v3, "\n\t")
                                v2 = v2
                                switch "" {
                                case "\n\t":
                                    switch fi49 {
                                    case -388:
                                        switch (fi51 & -645) {
                                        case 32768:
                                            v52 := GenS7{F0: 42}
                                            _ = v52
                                        case 32767:
                                            v1 = v1
                                        case -1:
                                            var v53 bool = true
                                            _ = v53
                                        default:
                                            v54 := GenS8{F0: -1, F1: "", F2: fi49}
                                            _ = v54
                                        }
                                    case -217:
                                        if len(v3) > 0 {
                                            v3[0] = "a"
                                        }
                                    default:
                                        if false {
                                            if true {
                                                v4 += v12
                                            } else {
                                                fmt.Println(v12, v4)
                                                switch ("foo bar" + "a") {
                                                case "\n\t":
                                                    v1 = 255
                                                case "foo bar":
                                                    var v55 uint32 = 42
                                                    _ = v55
                                                case "a":
                                                    var v56 uint8 = 2
                                                    _ = v56
                                                default:
                                                    switch (("hello" + "") + "\n\t") {
                                                    case "foo bar":
                                                        v57 := GenS8{F0: -596, F1: "", F2: fi51}
                                                        _ = v57
                                                    case "\n\t":
                                                        if len(v3) > 0 {
                                                            v3[0] = "foo bar"
                                                        }
                                                    }
                                                }
                                            }
                                            v58 := GenS5{F0: 42}
                                            _ = v58
                                        } else {
                                            v59 := GenS9{F0: 1, F1: (false && (!true)), F2: 343}
                                            _ = v59
                                        }
                                    }
                                    switch ("" + "") {
                                    case "":
                                        v3 = append(v3, "")
                                    case "hello":
                                        switch "hello" {
                                        case "":
                                            fmt.Println(fi51, v1, v2)
                                        }
                                    case "a":
                                        v1 = (v1 & v1)
                                    }
                                    fmt.Println(v12, v4, v1)
                                case "hello":
                                    var v60 float32 = float32(v12)
                                    _ = v60
                                case "a":
                                    if ((((((v2 + 0.0) * (v2 - float32(float64(v2)))) / (1.5 + float32(1.5))) / v2) <= v2) && true) {
                                        v1 -= (^(v1))
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    } else {
                                        v4 += (v4 * v4)
                                        var v61 float64 = float64(3.14159)
                                        switch "foo bar" {
                                        case "":
                                            fmt.Println(v4)
                                        }
                                        var v62 uint8 = 2
                                        _ = v61
                                        _ = v62
                                    }
                                default:
                                    switch (fi49 - fi51) {
                                    case 127:
                                        v2 += ((((((0.0 + ((v2 + (v2 * v2)) - v2)) + 0.0) + v2) / (float32(((v12 / v4) - (1.5 / ((v12 * (v4 + v4)) * 0.0)))) + v2)) - v2) + ((v2 - (0.0 - v2)) + 0.0))
                                        _ = len(v3)
                                        if ((((!false) || ((v1 * (v1 + v1)) > v1)) && ((true || (!true)) && true)) && true) {
                                            fmt.Println(fi49, fi49, fi49)
                                            switch (fi51 | 9223372036854775807) {
                                            case -754:
                                                v12 = 1e10
                                            case -2147483648:
                                                var v63 int = (fi51 % (((fi51 - (-(fi49)))) | 1))
                                                var v64 int64 = -9223372036854775808
                                                fmt.Println(v4, v12)
                                                fmt.Println(v64, v12, fi51)
                                                var v65 uint64 = 489
                                                _ = v63
                                                _ = v64
                                                _ = v65
                                            }
                                            var v66 uint32 = 42
                                            _ = v66
                                        } else {
                                            if (!(4294967295 <= 251)) {
                                                v2 = (1.5 + v2)
                                                var v67 float32 = ((v2 - (((((1.5 - ((v2 + 0.0) / ((v2 * v2) - (v2 - 1.5)))) * -2.25) + v2) - (v2 + ((v2 + (float32(v12) / 1.0)) + 0.0))) / ((float32(-2.25) / 3.14159) + (3.14159 + (v2 - v2))))) * v2)
                                                if (false || true) {
                                                    v67 += (-2.25 / (float32(v4) - v2))
                                                    if (false || false) {
                                                        if true {
                                                            switch (fi51 * fi51) {
                                                            case 0:
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            default:
                                                                _ = len(v3)
                                                            }
                                                            var v68 uint64 = 24
                                                            _ = v68
                                                        }
                                                        v69 := GenS10{F0: 832, F1: false}
                                                        v12 = v4
                                                        _ = v69
                                                    }
                                                } else {
                                                    v67 = 0.1
                                                }
                                                _ = v67
                                            }
                                            _ = len(v3)
                                            v3 = append(v3, "foo bar")
                                        }
                                        var v70 int32 = 42
                                        _ = v70
                                    case 2147483647:
                                        fmt.Println(fi49, fi51, fi49)
                                    case -471:
                                        if (true || (56 > 400)) {
                                            var v71 int64 = -632
                                            v3 = append(v3, ("a" + "\n\t"))
                                            _ = v71
                                        } else {
                                            v72 := GenS1{F0: 42, F1: fi49}
                                            switch "" {
                                            case "":
                                                if ((false || (true || (2 >= 0))) && false) {
                                                    switch -9223372036854775808 {
                                                    case 9223372036854775807:
                                                        switch ("\n\t" + "hello") {
                                                        case "":
                                                            v4 -= -2.25
                                                        default:
                                                            v73 := GenS7{F0: 220}
                                                            _ = v73
                                                        }
                                                    case 32768:
                                                        var v74 float32 = 0.0
                                                        _ = v74
                                                    default:
                                                        var v75 float64 = (1.5 * v4)
                                                        _ = v75
                                                    }
                                                    fmt.Println(v2, v4, fi51)
                                                    v2 += ((v2 + v2) * v2)
                                                } else {
                                                    if (-162 != -922) {
                                                        var v76 int8 = 100
                                                        v3 = append(v3, ("hello" + "a"))
                                                        fmt.Println(v2, v76)
                                                        _ = v76
                                                    } else {
                                                        var v77 uint8 = uint8(v1)
                                                        if len(v3) > 0 {
                                                            v3[0] = "hello"
                                                        }
                                                        var v78 uint8 = v77
                                                        if (!(!false)) {
                                                            v12 = ((v4 - (v12 + 0.1)) * float64((((float32((float64(float32(1.5)) / 0.1)) * (v2 - 1e10)) + ((v2 + (0.0 - v2)) - v2)) + v2)))
                                                            var v79 uint8 = v78
                                                            _ = v79
                                                        }
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                }
                                            case "foo bar":
                                                var v80 float32 = v2
                                                _ = v80
                                            }
                                            _ = v72
                                        }
                                    default:
                                        var v81 int32 = -184
                                        _ = v81
                                    }
                                }
                                _ = fi51
                            }
                        }
                        var v82 uint64 = 34
                        _ = fi49
                        _ = v82
                    }
                }
                v83 := GenS11{F0: 128, F1: 2}
                _ = v83
            } else {
                var v84 float64 = float64(0.1)
                _ = v84
            }
            v1 = (v1 - 32767)
        } else {
            v12 += v12
        }
    case -178:
        if ((!false) && true) {
            if ((1 < 65535) && false) {
                if (32767 >= -614) {
                    switch "a" {
                    case "":
                        var v85 uint8 = 1
                        _ = v85
                    case "foo bar":
                        fmt.Println(v2, v4, v12)
                    case "\n\t":
                        for fi86 := 0; fi86 < 6; fi86++ {
                            fmt.Println(v12)
                            if true {
                                for fi87 := 0; fi87 < 7; fi87++ {
                                    var v88 int64 = -9223372036854775808
                                    if true {
                                        v1 -= v1
                                        if true {
                                            v89 := GenS8{F0: 32767, F1: "", F2: (^(fi86))}
                                            switch "foo bar" {
                                            case "hello":
                                                v12 += (((float64((float32((float64(-2.25) + (0.0 / v12))) + v2)) + (float64(1e10) / v12)) - ((-2.25 * float64(v2)) / 1.5)) / 3.14159)
                                            case "\n\t":
                                                v1 = 470
                                            default:
                                                v1 = (v1 % ((v1) | 1))
                                            }
                                            _ = v89
                                        } else {
                                            v90 := GenS3{F0: 42, F1: 144}
                                            _ = v90
                                        }
                                    } else {
                                        var v91 int32 = 218
                                        var v92 float32 = v2
                                        v88 += v88
                                        _ = v91
                                        _ = v92
                                    }
                                    _ = fi87
                                    _ = v88
                                }
                                v2 = 3.14159
                                v1 = (v1 >> ((0) & 63))
                            }
                            if false {
                                if false {
                                    if (25 == 42) {
                                        var v93 string = "a"
                                        switch "a" {
                                        case "foo bar":
                                            var v94 bool = true
                                            _ = v94
                                        case "\n\t":
                                            v12 += v4
                                        }
                                        _ = v93
                                    }
                                } else {
                                    v2 -= ((v2 / (((float32(0.1) * v2) - ((float32(((3.14159 + v4) - float64(v2))) / v2) * 1.5)) / 1.5)) / ((v2 + v2) + v2))
                                    fmt.Println(v2)
                                }
                                if ((-1 >= 42) && false) {
                                    v2 = v2
                                    var v95 float64 = v4
                                    v1 += v1
                                    _ = v95
                                } else {
                                    if true {
                                        var v96 uint16 = (^(v1))
                                        for fi97 := 0; fi97 < 7; fi97++ {
                                            v2 += ((float32(0.1) - v2) * v2)
                                            _ = fi97
                                        }
                                        _ = v96
                                    } else {
                                        fmt.Println(v2, v12)
                                        var v98 float32 = v2
                                        _ = v98
                                    }
                                    v3 = append(v3, "\n\t")
                                }
                                var v99 int64 = -477
                                _ = v99
                            }
                            _ = fi86
                        }
                    default:
                        for fi100 := 0; fi100 < 1; fi100++ {
                            switch "foo bar" {
                            case "foo bar":
                                if false {
                                    if true {
                                        v12 -= v4
                                        fmt.Println(v2, fi100, v12)
                                        if ((((!true) && false) && true) && true) {
                                            switch "" {
                                            case "a":
                                                _ = len(v3)
                                            case "hello":
                                                for fi101 := 0; fi101 < 8; fi101++ {
                                                    v12 += v4
                                                    fmt.Println(fi100)
                                                    _ = fi101
                                                }
                                            default:
                                                _ = len(v3)
                                            }
                                            if (42 > 9223372036854775807) {
                                                var v102 int64 = 371
                                                _ = v102
                                            } else {
                                                var v103 uint8 = 42
                                                for fi104 := 0; fi104 < 16; fi104++ {
                                                    if ((!false) && true) {
                                                        if (32767 > 1) {
                                                            if (18446744073709551615 > 18446744073709551615) {
                                                                fmt.Println(v2)
                                                            } else {
                                                                var v105 float64 = (v4 - 0.0)
                                                                v2 -= v2
                                                                _ = v105
                                                            }
                                                            if false {
                                                                if (917 > 830) {
                                                                    if true {
                                                                        if true {
                                                                            switch ((("foo bar" + ((("hello" + "hello") + "hello") + "foo bar")) + "") + "") {
                                                                            case "\n\t":
                                                                                switch "" {
                                                                                case "a":
                                                                                    if ((true && false) || false) {
                                                                                        var v106 float64 = 0.0
                                                                                        _ = v106
                                                                                    } else {
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = ""
                                                                                        }
                                                                                        v1 += v1
                                                                                        if true {
                                                                                            fmt.Println(fi104, v2)
                                                                                            fmt.Println(v103, fi100)
                                                                                        }
                                                                                    }
                                                                                case "hello":
                                                                                    var v107 float64 = v4
                                                                                    _ = v107
                                                                                }
                                                                            case "":
                                                                                var v108 uint16 = v1
                                                                                _ = v108
                                                                            case "a":
                                                                                var v109 float32 = v2
                                                                                _ = v109
                                                                            default:
                                                                                v1 = 42
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v1 -= v1
                                                                        var v110 int64 = 9223372036854775807
                                                                        _ = v110
                                                                    }
                                                                }
                                                            }
                                                            switch (fi104 & fi100) {
                                                            case 128:
                                                                if true {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = "\n\t"
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            var v111 uint32 = 4294967295
                                                            v112 := GenS12{F0: -2147483648, F1: fi100, F2: false}
                                                            _ = v111
                                                            _ = v112
                                                        }
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    } else {
                                                        var v113 int = 4294967295
                                                        _ = v113
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    if true {
                                                        var v114 string = (((("\n\t" + "a") + "\n\t") + "foo bar") + ("\n\t" + "hello"))
                                                        _ = v114
                                                    }
                                                    _ = fi104
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = "foo bar"
                                                }
                                                _ = v103
                                            }
                                            if true {
                                                var v115 string = "a"
                                                v116 := GenS13{F0: 1, F1: false, F2: 65535}
                                                for fi117 := 0; fi117 < 17; fi117++ {
                                                    var v118 int8 = genFn0(933, 0)
                                                    var v119 float32 = ((((1e10 * float32(v4)) * (v2 - v2)) * 1e10) * v2)
                                                    v120 := GenS6{F0: 255, F1: (v118 * v118)}
                                                    switch ((v115 + ((v115 + "\n\t") + ((v115 + "") + ((v115 + "foo bar") + ("hello" + v115))))) + v115) {
                                                    case "":
                                                        var v121 float32 = ((v2 + (v119 * v2)) * v119)
                                                        _ = v121
                                                    case "foo bar":
                                                        v2 += v2
                                                    }
                                                    var v122 uint8 = 127
                                                    _ = fi117
                                                    _ = v118
                                                    _ = v119
                                                    _ = v120
                                                    _ = v122
                                                }
                                                _ = v115
                                                _ = v116
                                            }
                                        }
                                    } else {
                                        fmt.Println(v4, v1)
                                    }
                                }
                            }
                            var v123 int16 = 42
                            _ = fi100
                            _ = v123
                        }
                    }
                }
                if (-1 > -732) {
                    _ = len(v3)
                    var v124 uint = 4
                    v3 = append(v3, "")
                    _ = v124
                }
            }
        }
    case 2147483647:
        var v125 bool = true
        _ = v125
    default:
        if (!false) {
            v126 := GenS3{F0: 32767, F1: 0}
            v3 = append(v3, ("foo bar" + ("" + "")))
            var v127 int64 = 42
            _ = v126
            _ = v127
        } else {
            fmt.Println(v2, v1)
            v2 += (3.14159 * v2)
            var v128 uint16 = 32767
            _ = v128
        }
    }
    if (57 > 0) {
        if true {
            v129 := GenS6{F0: 89, F1: 96}
            _ = v129
        } else {
            fmt.Println(v4, v1, v2)
            v3 = append(v3, "\n\t")
            var v130 int16 = 1
            _ = v130
        }
        if false {
            v1 += v1
            var v131 string = "a"
            v1 = (v1 / (((v1 << ((4294967295) & 63))) | 1))
            _ = v131
        }
    }
    v12 += (0.1 * 3.14159)
    var v132 float64 = v12
    for fi133 := 0; fi133 < 20; fi133++ {
        if len(v3) > 0 {
            v3[0] = ""
        }
        if len(v3) > 0 {
            v3[0] = (("a" + (((("a" + "foo bar") + "") + ("hello" + "")) + "a")) + "hello")
        }
        v2 = (((v2 + v2) + v2) * ((0.1 + (v2 * (0.1 - 0.1))) / v2))
        _ = fi133
    }
    fmt.Println(v12, v132, v132)
    var v134 float64 = v4
    var v135 int64 = 1
    for fi136 := 0; fi136 < 19; fi136++ {
        for fi137 := 0; fi137 < 11; fi137++ {
            var v138 string = "hello"
            v134 += v12
            v3 = append(v3, v138)
            v2 = v2
            _ = fi137
            _ = v138
        }
        v132 = v12
        for fi139 := 0; fi139 < 15; fi139++ {
            var v140 int32 = 42
            v141 := GenS14{F0: ((v1 | (v1 >> ((0) & 63))) < v1)}
            v12 += v132
            v142 := GenS15{F0: 762, F1: (v1 | v1)}
            _ = fi139
            _ = v140
            _ = v141
            _ = v142
        }
        _ = fi136
    }
    v143 := GenS16{F0: (v1 & v1), F1: 42, F2: v135}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v132:", v132)
    fmt.Println("v134:", v134)
    fmt.Println("v135:", v135)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v132
    _ = v134
    _ = v135
    _ = v143
    _ = v2
    _ = v3
    _ = v4
}
