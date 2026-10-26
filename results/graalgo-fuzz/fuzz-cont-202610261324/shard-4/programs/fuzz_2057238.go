package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint; F1 uint32; F2 bool }
type GenS5 struct { F0 uint64; F1 int8; F2 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int8; F1 int64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int16; F1 uint64; F2 bool }
type GenS15 struct { F0 bool; F1 int32; F2 uint32 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 int8; F1 uint32 }
type GenS21 struct { F0 uint8; F1 int; F2 uint }
type GenS22 struct { F0 uint16 }
type GenS23 struct { F0 bool }
type GenS24 struct { F0 uint64; F1 uint }
func genFn0(p0 uint64) uint {
    v5 := GenS1{F0: (("hello" == ("foo bar" + "a")) || false)}
    v6 := GenS1{F0: (!((^(p0)) >= (p0 ^ p0)))}
    switch ("foo bar" + "hello") {
    case "\n\t":
        switch ("a" + "") {
        case "\n\t":
            var v7 bool = (1.5 < 3.14159)
            _ = v7
        }
    case "foo bar":
        switch 284 {
        case 4294967295:
            v8 := []int{32768, -576}
            _ = v8
        default:
            v9 := GenS2{F0: (p0 + (p0 % (((p0 / ((38) | 1))) | 1))), F1: 1, F2: -1}
            _ = v9
        }
    case "hello":
        p0 = p0
    default:
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := []int{2147483647}
            switch 32768 {
            case 32768:
                if len(v11) > 0 {
                    v11[0] = (fi10 + 9223372036854775807)
                }
            default:
                p0 = (^(p0))
            }
            p0 += (p0 >> ((18446744073709551615) & 63))
            _ = fi10
            _ = v11
        }
    }
    var v12 uint8 = uint8(p0)
    _ = p0
    _ = v12
    _ = v5
    _ = v6
    return uint(p0)
}

func main() {
    var v1 int8 = 0
    v2 := []string{(("hello" + ("hello" + "hello")) + "foo bar"), "foo bar", "hello"}
    v3 := []string{(("" + ("a" + ("\n\t" + "\n\t"))) + ("hello" + "")), ("\n\t" + ""), ("foo bar" + ("foo bar" + "hello"))}
    v4 := []string{("" + ("foo bar" + "a")), "foo bar", "foo bar"}
    fmt.Println(v1)
    switch 590 {
    case -9223372036854775808:
        var v13 bool = true
        _ = v13
    default:
        switch 127 {
        case 98:
            for fi14 := 0; fi14 < 2; fi14++ {
                for fi15 := 0; fi15 < 13; fi15++ {
                    v1 += v1
                    _ = fi15
                }
                v16 := GenS1{F0: true}
                _ = fi14
                _ = v16
            }
        case 32767:
            for fi17 := 0; fi17 < 1; fi17++ {
                v18 := GenS2{F0: 18446744073709551615, F1: 400, F2: 1}
                _ = len(v3)
                v19 := GenS1{F0: ((fi17 & (fi17 & fi17)) <= fi17)}
                _ = fi17
                _ = v18
                _ = v19
            }
        default:
            v1 += v1
        }
    }
    var v20 uint64 = 296
    for fi21 := 0; fi21 < 8; fi21++ {
        if (!("" > "a")) {
            v20 = (^(v20))
        }
        if len(v2) > 0 {
            v2[0] = ("\n\t" + "")
        }
        _ = fi21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        v23 := GenS3{F0: v1, F1: -1}
        _ = fi22
        _ = v23
    }
    fmt.Println(v1)
    v20 += v20
    v24 := GenS4{F0: genFn0(2), F1: 4294967295, F2: false}
    switch (("foo bar" + (("a" + (("\n\t" + "foo bar") + ("\n\t" + ("hello" + "hello")))) + "foo bar")) + "") {
    case "":
        v20 += 471
    default:
        if true {
            if true {
                var v25 uint8 = 255
                var v26 string = "hello"
                _ = v25
                _ = v26
            } else {
                switch 849 {
                case 606:
                    switch 4294967295 {
                    case 32768:
                        switch "a" {
                        case "\n\t":
                            fmt.Println(v1, v1)
                        }
                    default:
                        if len(v3) > 0 {
                            v3[0] = ""
                        }
                    }
                case 9223372036854775807:
                    fmt.Println(v20, v20)
                case 42:
                    for fi27 := 0; fi27 < 2; fi27++ {
                        switch fi27 {
                        case 802:
                            v28 := GenS2{F0: (v20 / ((v20) | 1)), F1: 0, F2: 1000}
                            _ = v28
                        case -497:
                            var v29 uint32 = 0
                            _ = v29
                        default:
                            fmt.Println(v20)
                        }
                        switch ("hello" + ("\n\t" + "hello")) {
                        case "foo bar":
                            switch "\n\t" {
                            case "\n\t":
                                switch ("\n\t" + "") {
                                case "":
                                    fmt.Println(v1)
                                case "a":
                                    v20 = 1
                                case "hello":
                                    var v30 string = "a"
                                    _ = v30
                                }
                            case "hello":
                                var v31 uint8 = 42
                                _ = v31
                            default:
                                v2 = append(v2, ("\n\t" + "\n\t"))
                            }
                        case "a":
                            switch (fi27 + 804) {
                            case 1:
                                switch "hello" {
                                case "a":
                                    for fi32 := 0; fi32 < 9; fi32++ {
                                        if (((!(-1 != 180)) && ((fi32 % ((int(v20)) | 1)) <= (fi27 ^ (fi27 | 2147483647)))) && (!(((v1 / ((1) | 1)) != (v1 ^ v1)) || (fi32 == fi32)))) {
                                            v20 += (v20 | 2)
                                            var v33 float32 = 0.0
                                            v34 := GenS1{F0: false}
                                            _ = v33
                                            _ = v34
                                        }
                                        var v35 uint8 = 2
                                        if true {
                                            switch fi32 {
                                            case 127:
                                                v36 := GenS5{F0: 42, F1: v1, F2: (!true)}
                                                _ = v36
                                            case -416:
                                                v3 = append(v3, "\n\t")
                                            }
                                            if true {
                                                switch "" {
                                                case "foo bar":
                                                    if true {
                                                        var v37 bool = false
                                                        _ = v37
                                                    } else {
                                                        v2 = append(v2, (("a" + "a") + "foo bar"))
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if false {
                                                            fmt.Println(fi32, v1)
                                                            var v38 uint32 = 0
                                                            _ = v38
                                                        }
                                                        fmt.Println(fi27)
                                                        var v39 bool = (!false)
                                                        _ = v39
                                                    }
                                                case "a":
                                                    var v40 string = "\n\t"
                                                    _ = v40
                                                case "":
                                                    v35 = (^(v35))
                                                default:
                                                    if (fi27 >= (fi32 ^ (fi27 ^ -907))) {
                                                        var v41 int8 = (v1 & 127)
                                                        _ = v41
                                                    }
                                                }
                                            }
                                        } else {
                                            var v42 int8 = 109
                                            if false {
                                                if ((410 > 2) && ((fi27 >> ((genFn0((v20 * 1))) & 63)) > fi27)) {
                                                    fmt.Println(fi32)
                                                }
                                                switch "\n\t" {
                                                case "":
                                                    if (true && (!(!true))) {
                                                        if len(v3) > 0 {
                                                            v3[0] = ("" + ("foo bar" + "\n\t"))
                                                        }
                                                        var v43 int8 = (v42 + 42)
                                                        var v44 uint64 = v20
                                                        var v45 float32 = 0.0
                                                        v2 = append(v2, "")
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                case "\n\t":
                                                    var v46 float32 = ((1e10 / (((0.1 - (1.5 * (float32(1.5) + 0.0))) + 3.14159) * (3.14159 * (((0.0 - 0.0) / 0.1) - 1.5)))) / 3.14159)
                                                    _ = v46
                                                }
                                            } else {
                                                v20 = v20
                                            }
                                            switch (("hello" + "hello") + "") {
                                            case "":
                                                var v47 int32 = 126
                                                _ = v47
                                            }
                                            _ = v42
                                        }
                                        v48 := GenS1{F0: false}
                                        _ = fi32
                                        _ = v35
                                        _ = v48
                                    }
                                default:
                                    fmt.Println(fi27, v1)
                                }
                            case -1:
                                for fi49 := 0; fi49 < 2; fi49++ {
                                    _ = len(v4)
                                    v50 := GenS6{F0: fi49}
                                    _ = fi49
                                    _ = v50
                                }
                            default:
                                v20 -= (v20 * 0)
                            }
                        case "hello":
                            var v51 int = fi27
                            _ = v51
                        default:
                            v52 := GenS2{F0: 18446744073709551615, F1: 324, F2: 42}
                            _ = v52
                        }
                        var v53 uint64 = (v20 + (^(v20)))
                        _ = fi27
                        _ = v53
                    }
                default:
                    if (!(true || true)) {
                        _ = len(v2)
                        var v54 int32 = 2147483647
                        _ = v54
                    }
                }
                var v55 uint = 79
                _ = v55
            }
        }
    }
    fmt.Println(v20, v20)
    if ("a" == ("\n\t" + "a")) {
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        if (42 > 4294967295) {
            var v56 float64 = -2.25
            _ = v56
        } else {
            if false {
                v57 := GenS4{F0: 403, F1: 1, F2: (0 <= 1)}
                var v58 uint8 = 42
                _ = v57
                _ = v58
            } else {
                v59 := GenS7{F0: 1, F1: -9223372036854775808}
                if (2 >= v20) {
                    if true {
                        switch "hello" {
                        case "foo bar":
                            switch "a" {
                            case "\n\t":
                                if true {
                                    if (-2.25 != (1.5 * (1.5 / 1e10))) {
                                        if false {
                                            if true {
                                                if len(v4) > 0 {
                                                    v4[0] = "hello"
                                                }
                                                switch 128 {
                                                case 840:
                                                    v1 = v1
                                                case 127:
                                                    if true {
                                                        v1 += -128
                                                    } else {
                                                        switch -773 {
                                                        case -9223372036854775808:
                                                            v60 := GenS8{F0: (v20 - (v20 | v20))}
                                                            _ = v60
                                                        case 0:
                                                            v20 += v20
                                                        }
                                                    }
                                                case 128:
                                                    var v61 int8 = (v1 - (v1 - -128))
                                                    _ = v61
                                                default:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                v62 := GenS9{F0: (v1 + -128), F1: -1}
                                                _ = v62
                                            }
                                        } else {
                                            var v63 bool = true
                                            _ = v63
                                        }
                                    } else {
                                        v64 := GenS10{F0: 0}
                                        _ = v64
                                    }
                                } else {
                                    switch ("a" + ((("hello" + "foo bar") + "hello") + ("" + "hello"))) {
                                    case "foo bar":
                                        switch 127 {
                                        case 128:
                                            v65 := GenS11{F0: 2}
                                            _ = v65
                                        case 32768:
                                            for fi66 := 0; fi66 < 3; fi66++ {
                                                fmt.Println(v1, v1, fi66)
                                                for fi67 := 0; fi67 < 17; fi67++ {
                                                    v68 := GenS9{F0: (v1 >> ((0) & 63)), F1: -9223372036854775808}
                                                    _ = fi67
                                                    _ = v68
                                                }
                                                if (-9223372036854775808 < 512) {
                                                    var v69 int32 = 1
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    _ = v69
                                                } else {
                                                    for fi70 := 0; fi70 < 11; fi70++ {
                                                        var v71 int16 = 0
                                                        _ = len(v4)
                                                        v4 = append(v4, "foo bar")
                                                        _ = fi70
                                                        _ = v71
                                                    }
                                                }
                                                _ = fi66
                                            }
                                        case 1:
                                            switch 32767 {
                                            case -588:
                                                _ = len(v4)
                                            case 4294967295:
                                                for fi72 := 0; fi72 < 17; fi72++ {
                                                    for fi73 := 0; fi73 < 14; fi73++ {
                                                        v20 -= v20
                                                        v74 := GenS12{F0: 2147483647}
                                                        v1 -= v1
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        v75 := GenS13{F0: ("a" + ("" + ("a" + ("" + "a"))))}
                                                        _ = fi73
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                    v20 -= (^(v20))
                                                    _ = fi72
                                                }
                                            case 32768:
                                                var v76 float64 = float64(1e10)
                                                _ = v76
                                            }
                                        default:
                                            switch "\n\t" {
                                            case "":
                                                for fi77 := 0; fi77 < 8; fi77++ {
                                                    for fi78 := 0; fi78 < 4; fi78++ {
                                                        switch fi77 {
                                                        case 4294967295:
                                                            if true {
                                                                v2 = append(v2, "a")
                                                                if len(v2) > 0 {
                                                                    v2[0] = "a"
                                                                }
                                                            } else {
                                                                var v79 int = fi77
                                                                switch 1 {
                                                                case 32768:
                                                                    v20 = v20
                                                                }
                                                                _ = v79
                                                            }
                                                        case -9223372036854775808:
                                                            v20 = v20
                                                        case 32767:
                                                            var v80 int = (-(fi77))
                                                            _ = v80
                                                        }
                                                        _ = fi78
                                                    }
                                                    v20 -= v20
                                                    _ = fi77
                                                }
                                            default:
                                                if true {
                                                    for fi81 := 0; fi81 < 19; fi81++ {
                                                        if false {
                                                            if (964 < 65535) {
                                                                var v82 uint16 = 1
                                                                for fi83 := 0; fi83 < 11; fi83++ {
                                                                    switch ("\n\t" + "") {
                                                                    case "a":
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        v84 := GenS6{F0: 42}
                                                                        _ = v84
                                                                    case "foo bar":
                                                                        v1 += 0
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        if true {
                                                                            if true {
                                                                                if true {
                                                                                    fmt.Println(fi83, v82, fi81)
                                                                                } else {
                                                                                    v2 = append(v2, "\n\t")
                                                                                    var v85 uint64 = 758
                                                                                    v86 := GenS9{F0: (-(v1)), F1: 1}
                                                                                    _ = v85
                                                                                    _ = v86
                                                                                }
                                                                                switch (fi83 & int(v82)) {
                                                                                case 32767:
                                                                                    v20 += (^(v20))
                                                                                case 212:
                                                                                    var v87 uint64 = 4294967295
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = ("hello" + ("\n\t" + "hello"))
                                                                                    }
                                                                                    _ = v87
                                                                                case -667:
                                                                                    v88 := GenS12{F0: 596}
                                                                                    _ = v88
                                                                                }
                                                                            } else {
                                                                                switch (fi81 | 368) {
                                                                                case 9223372036854775807:
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                default:
                                                                                    v89 := GenS14{F0: 651, F1: 101, F2: ("foo bar" < "")}
                                                                                    _ = v89
                                                                                }
                                                                                if ("\n\t" == "") {
                                                                                    if ((!(((v20 & v20) >= v20) && true)) || false) {
                                                                                        fmt.Println(fi83)
                                                                                    } else {
                                                                                        v2 = append(v2, ("" + ""))
                                                                                    }
                                                                                    fmt.Println(fi81)
                                                                                    if true {
                                                                                        var v90 uint8 = 128
                                                                                        fmt.Println(v20)
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = ("foo bar" + "foo bar")
                                                                                        }
                                                                                        _ = v90
                                                                                    } else {
                                                                                        v20 = 0
                                                                                        v1 = v1
                                                                                        if false {
                                                                                            if false {
                                                                                                v1 -= v1
                                                                                                if (0.1 < ((0.1 - float32(float64((-2.25 - 0.0)))) + (((3.14159 + (3.14159 * 1e10)) + 3.14159) + -2.25))) {
                                                                                                    v20 = v20
                                                                                                    if false {
                                                                                                        fmt.Println(fi83)
                                                                                                    } else {
                                                                                                        var v91 uint8 = 123
                                                                                                        _ = v91
                                                                                                    }
                                                                                                    var v92 int8 = (v1 ^ v1)
                                                                                                    if (false && false) {
                                                                                                        if (313 != 57) {
                                                                                                            fmt.Println(v20, fi83, v1)
                                                                                                            _ = len(v4)
                                                                                                        } else {
                                                                                                            switch fi81 {
                                                                                                            case 32768:
                                                                                                                v93 := GenS15{F0: true, F1: 42, F2: 1}
                                                                                                                _ = v93
                                                                                                            case 4294967295:
                                                                                                                _ = len(v2)
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        if ((!(-916 == 1)) || false) {
                                                                                                            v92 -= 0
                                                                                                            var v94 uint32 = 931
                                                                                                            _ = v94
                                                                                                        } else {
                                                                                                            if (-2.25 == 1.5) {
                                                                                                                v20 = (v20 + v20)
                                                                                                                v82 -= v82
                                                                                                            } else {
                                                                                                                var v95 int64 = -9223372036854775808
                                                                                                                if len(v3) > 0 {
                                                                                                                    v3[0] = "foo bar"
                                                                                                                }
                                                                                                                var v96 uint = 202
                                                                                                                _ = v95
                                                                                                                _ = v96
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(v20, v82, fi81)
                                                                                                    }
                                                                                                    _ = v92
                                                                                                }
                                                                                                var v97 string = ("a" + "a")
                                                                                                _ = v97
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v82 = v82
                                                                                } else {
                                                                                    var v98 uint8 = 81
                                                                                    if true {
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        v20 = v20
                                                                                    } else {
                                                                                        var v99 float32 = (0.1 - (3.14159 + 0.1))
                                                                                        _ = v99
                                                                                    }
                                                                                    v98 = 130
                                                                                    if false {
                                                                                        if false {
                                                                                            fmt.Println(v20)
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = (("hello" + "\n\t") + "hello")
                                                                                            }
                                                                                            fmt.Println(v1, v1, fi81)
                                                                                            fmt.Println(fi83)
                                                                                        } else {
                                                                                            switch "foo bar" {
                                                                                            case "hello":
                                                                                                var v100 uint8 = (v98 * v98)
                                                                                                _ = v100
                                                                                            case "a":
                                                                                                v101 := GenS1{F0: false}
                                                                                                _ = v101
                                                                                            case "foo bar":
                                                                                                v102 := GenS12{F0: 116}
                                                                                                _ = v102
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        v82 += 1
                                                                                        if (335 == int16(v20)) {
                                                                                            switch (fi81 + (fi81 & (^(fi83)))) {
                                                                                            case -2147483648:
                                                                                                var v103 uint16 = v82
                                                                                                _ = v103
                                                                                            case 261:
                                                                                                fmt.Println(fi81, fi81, v20)
                                                                                            default:
                                                                                                v1 = v1
                                                                                            }
                                                                                            if true {
                                                                                                var v104 int32 = 964
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = (("\n\t" + "foo bar") + ("" + "a"))
                                                                                                }
                                                                                                if (!true) {
                                                                                                    fmt.Println(v98)
                                                                                                    var v105 uint16 = (v82 & v82)
                                                                                                    _ = v105
                                                                                                } else {
                                                                                                    switch fi81 {
                                                                                                    case 723:
                                                                                                        var v106 int8 = (v1 / ((v1) | 1))
                                                                                                        if (18446744073709551615 > v20) {
                                                                                                            var v107 bool = true
                                                                                                            _ = v107
                                                                                                        } else {
                                                                                                            v108 := GenS16{F0: 127}
                                                                                                            var v109 uint64 = v20
                                                                                                            _ = v108
                                                                                                            _ = v109
                                                                                                        }
                                                                                                        _ = v106
                                                                                                    default:
                                                                                                        v82 -= v82
                                                                                                    }
                                                                                                }
                                                                                                _ = v104
                                                                                            } else {
                                                                                                v110 := GenS15{F0: (-1 < int32(v20)), F1: -439, F2: 65535}
                                                                                                var v111 int8 = v1
                                                                                                _ = v110
                                                                                                _ = v111
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(v82, fi81)
                                                                                            fmt.Println(v20)
                                                                                            v20 += v20
                                                                                        }
                                                                                        var v112 uint64 = 1
                                                                                        v113 := GenS9{F0: 127, F1: -829}
                                                                                        if ((1e10 + -2.25) < 1.5) {
                                                                                            if true {
                                                                                                if (1000 < 42) {
                                                                                                    switch ("" + ("foo bar" + "hello")) {
                                                                                                    case "foo bar":
                                                                                                        var v114 string = ("" + "hello")
                                                                                                        v4 = append(v4, "hello")
                                                                                                        _ = v114
                                                                                                    default:
                                                                                                        if false {
                                                                                                            switch "hello" {
                                                                                                            case "hello":
                                                                                                                v3 = append(v3, (("\n\t" + "foo bar") + "foo bar"))
                                                                                                            case "\n\t":
                                                                                                                v4 = append(v4, ("a" + ""))
                                                                                                            case "a":
                                                                                                                if (true || false) {
                                                                                                                    switch "foo bar" {
                                                                                                                    case "":
                                                                                                                        var v115 uint64 = 42
                                                                                                                        _ = v115
                                                                                                                    case "hello":
                                                                                                                        var v116 int64 = 0
                                                                                                                        _ = v116
                                                                                                                    case "\n\t":
                                                                                                                        fmt.Println(v98, v98, fi83)
                                                                                                                    default:
                                                                                                                        var v117 string = "hello"
                                                                                                                        _ = v117
                                                                                                                    }
                                                                                                                    v112 -= v20
                                                                                                                } else {
                                                                                                                    var v118 int = 127
                                                                                                                    v4 = append(v4, (("foo bar" + ("" + "")) + "a"))
                                                                                                                    _ = v118
                                                                                                                }
                                                                                                            default:
                                                                                                                var v119 int16 = 1
                                                                                                                _ = v119
                                                                                                            }
                                                                                                            v120 := GenS17{F0: 42}
                                                                                                            switch int(v1) {
                                                                                                            case 32767:
                                                                                                                _ = len(v4)
                                                                                                            case -2147483648:
                                                                                                                _ = len(v2)
                                                                                                            default:
                                                                                                                v121 := GenS18{F0: -1}
                                                                                                                _ = v121
                                                                                                            }
                                                                                                            var v122 int64 = 762
                                                                                                            _ = v120
                                                                                                            _ = v122
                                                                                                        } else {
                                                                                                            v123 := GenS19{F0: ((v1 < (v1 ^ (v1 - (^(v1))))) || (false || true))}
                                                                                                            _ = v123
                                                                                                        }
                                                                                                    }
                                                                                                    if false {
                                                                                                        v82 += 42
                                                                                                        if (int64(v98) == int64(v98)) {
                                                                                                            v82 = 259
                                                                                                            var v124 float32 = (float32(1e10) - 0.0)
                                                                                                            if false {
                                                                                                                var v125 uint16 = v82
                                                                                                                _ = v125
                                                                                                            }
                                                                                                            _ = v124
                                                                                                        }
                                                                                                        var v126 int = (fi83 & fi81)
                                                                                                        _ = v126
                                                                                                    } else {
                                                                                                        _ = len(v3)
                                                                                                    }
                                                                                                }
                                                                                                v127 := GenS19{F0: true}
                                                                                                _ = v127
                                                                                            } else {
                                                                                                v128 := GenS20{F0: v1, F1: 4294967295}
                                                                                                v3 = append(v3, "\n\t")
                                                                                                v82 = (^(v82))
                                                                                                _ = v128
                                                                                            }
                                                                                            var v129 int64 = 555
                                                                                            if false {
                                                                                                v130 := GenS2{F0: (v20 ^ (^(v20))), F1: v82, F2: 32767}
                                                                                                switch "a" {
                                                                                                case "a":
                                                                                                    var v131 int8 = 98
                                                                                                    _ = v131
                                                                                                case "hello":
                                                                                                    var v132 int64 = (v129 / (((-(v129))) | 1))
                                                                                                    _ = v132
                                                                                                default:
                                                                                                    if (("a" + ("\n\t" + (("a" + "foo bar") + "a"))) <= "hello") {
                                                                                                        var v133 int64 = -9223372036854775808
                                                                                                        _ = v133
                                                                                                    }
                                                                                                }
                                                                                                if true {
                                                                                                    v2 = append(v2, "foo bar")
                                                                                                    v20 += (v112 + (^(v112)))
                                                                                                    v134 := GenS8{F0: v112}
                                                                                                    v1 = v1
                                                                                                    _ = v134
                                                                                                }
                                                                                                var v135 uint16 = (^(v82))
                                                                                                fmt.Println(v82, v20)
                                                                                                v82 = (v135 / ((v135) | 1))
                                                                                                _ = v130
                                                                                                _ = v135
                                                                                            } else {
                                                                                                var v136 string = "foo bar"
                                                                                                _ = v136
                                                                                            }
                                                                                            _ = v129
                                                                                        }
                                                                                        _ = v112
                                                                                        _ = v113
                                                                                    }
                                                                                    v20 = v20
                                                                                    _ = v98
                                                                                }
                                                                                v137 := GenS14{F0: 0, F1: (v20 << ((1) & 63)), F2: (v82 == (v82 << ((4294967295) & 63)))}
                                                                                _ = v137
                                                                            }
                                                                            if (2147483648 > 471) {
                                                                                switch fi81 {
                                                                                case 1:
                                                                                    var v138 int32 = -1
                                                                                    _ = v138
                                                                                }
                                                                                var v139 int = fi81
                                                                                var v140 uint16 = v82
                                                                                _ = v139
                                                                                _ = v140
                                                                            } else {
                                                                                v141 := GenS21{F0: 25, F1: int(v20), F2: genFn0((^(v20)))}
                                                                                if (!((v1 % (((v1 % ((127) | 1))) | 1)) == v1)) {
                                                                                    fmt.Println(v82, v20)
                                                                                    var v142 int64 = 1
                                                                                    _ = v142
                                                                                }
                                                                                _ = v141
                                                                            }
                                                                        } else {
                                                                            fmt.Println(fi83, fi83, v20)
                                                                            var v143 float64 = (-2.25 * 1e10)
                                                                            var v144 uint64 = v20
                                                                            _ = v143
                                                                            _ = v144
                                                                        }
                                                                    }
                                                                    _ = fi83
                                                                }
                                                                for fi145 := 0; fi145 < 8; fi145++ {
                                                                    v20 = (v20 / (((v20 >> ((genFn0((v20 % (((v20 >> ((4294967295) & 63))) | 1)))) & 63))) | 1))
                                                                    v82 -= v82
                                                                    _ = fi145
                                                                }
                                                                _ = v82
                                                            }
                                                            var v146 int16 = 32767
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v146
                                                        } else {
                                                            v20 += 18446744073709551615
                                                            v147 := GenS22{F0: 339}
                                                            _ = v147
                                                        }
                                                        _ = fi81
                                                    }
                                                } else {
                                                    v20 = (v20 + v20)
                                                    var v148 string = "hello"
                                                    _ = v148
                                                }
                                            }
                                        }
                                    case "hello":
                                        for fi149 := 0; fi149 < 13; fi149++ {
                                            for fi150 := 0; fi150 < 7; fi150++ {
                                                fmt.Println(v20, fi149, fi149)
                                                var v151 string = "foo bar"
                                                v152 := GenS3{F0: (v1 >> ((261) & 63)), F1: 78}
                                                _ = fi150
                                                _ = v151
                                                _ = v152
                                            }
                                            switch fi149 {
                                            case 127:
                                                if false {
                                                    fmt.Println(v1, v1)
                                                    v1 += (^(v1))
                                                }
                                            }
                                            var v153 float32 = ((((float32(1e10) - -2.25) / float32(-2.25)) / (-2.25 + float32(0.0))) * -2.25)
                                            _ = fi149
                                            _ = v153
                                        }
                                    case "\n\t":
                                        if (32768 != -9223372036854775808) {
                                            v154 := GenS1{F0: true}
                                            if (-660 > 2147483647) {
                                                var v155 int64 = -982
                                                if ((661 > 2147483647) || (true || false)) {
                                                    v4 = append(v4, "a")
                                                    var v156 uint16 = 956
                                                    _ = v156
                                                } else {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v1 += 18
                                                    case "a":
                                                        switch 4294967295 {
                                                        case 128:
                                                            var v157 int32 = 42
                                                            _ = v157
                                                        default:
                                                            var v158 int = -2147483648
                                                            _ = v158
                                                        }
                                                    }
                                                }
                                                _ = v155
                                            }
                                            if (!(753 != 318)) {
                                                fmt.Println(v20)
                                                if (!false) {
                                                    fmt.Println(v20, v1)
                                                    if true {
                                                        if len(v4) > 0 {
                                                            v4[0] = "foo bar"
                                                        }
                                                        var v159 uint = 255
                                                        _ = v159
                                                    }
                                                }
                                                for fi160 := 0; fi160 < 9; fi160++ {
                                                    v161 := GenS23{F0: false}
                                                    for fi162 := 0; fi162 < 20; fi162++ {
                                                        fmt.Println(fi162, fi162, fi160)
                                                        if false {
                                                            v163 := GenS24{F0: v20, F1: 18446744073709551615}
                                                            fmt.Println(v1)
                                                            _ = v163
                                                        }
                                                        _ = fi162
                                                    }
                                                    var v164 float32 = 1.5
                                                    _ = fi160
                                                    _ = v161
                                                    _ = v164
                                                }
                                            }
                                            _ = v154
                                        }
                                    default:
                                        var v165 int16 = 701
                                        _ = v165
                                    }
                                    v166 := GenS2{F0: (v20 | 0), F1: 32767, F2: -32768}
                                    _ = v166
                                }
                            case "":
                                fmt.Println(v1, v1)
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        case "hello":
                            v167 := GenS18{F0: -547}
                            _ = v167
                        default:
                            if true {
                                fmt.Println(v1, v1)
                            }
                        }
                    } else {
                        fmt.Println(v20)
                    }
                }
                _ = v59
            }
        }
        fmt.Println(v1, v20)
    } else {
        v4 = append(v4, ("\n\t" + "\n\t"))
    }
    v168 := GenS16{F0: v1}
    switch "" {
    case "":
        switch 866 {
        case -660:
            for fi169 := 0; fi169 < 1; fi169++ {
                if true {
                    for fi170 := 0; fi170 < 1; fi170++ {
                        if false {
                            var v171 uint32 = uint32(v20)
                            _ = v171
                        }
                        switch 0 {
                        case 633:
                            if (!(!false)) {
                                var v172 uint32 = 2147483647
                                if false {
                                    if false {
                                        v173 := GenS2{F0: (^(v20)), F1: 42, F2: -1}
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v173
                                    } else {
                                        var v174 int32 = -2147483648
                                        _ = v174
                                    }
                                } else {
                                    var v175 uint8 = uint8(v172)
                                    _ = v175
                                }
                                v172 += (v172 % ((v172) | 1))
                                var v176 int8 = (v1 / ((int8(fi169)) | 1))
                                _ = v172
                                _ = v176
                            }
                        case -525:
                            v20 += (v20 << ((477) & 63))
                        }
                        _ = fi170
                    }
                    v20 += v20
                }
                for fi177 := 0; fi177 < 4; fi177++ {
                    var v178 string = (("" + (("hello" + "hello") + "\n\t")) + "foo bar")
                    fmt.Println(v20, v1)
                    v3 = append(v3, (v178 + v178))
                    _ = fi177
                    _ = v178
                }
                _ = fi169
            }
        case 1:
            if len(v2) > 0 {
                v2[0] = "hello"
            }
        }
    default:
        var v179 string = "\n\t"
        _ = v179
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v168
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
}
