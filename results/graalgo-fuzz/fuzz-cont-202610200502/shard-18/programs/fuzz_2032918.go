package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 bool; F2 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 bool; F2 uint32 }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 uint; F1 int64; F2 uint8 }
type GenS9 struct { F0 uint8; F1 int }
type GenS10 struct { F0 uint; F1 uint64; F2 string }
type GenS11 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS12 struct { F0 int32; F1 int16 }
type GenS13 struct { F0 int32; F1 uint }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]string)
    var v3 int8 = -128
    var v4 int8 = v3
    v4 += (v4 | -16)
    v1 += -2.25
    if true {
        if (!((!true) && (282 >= 128))) {
            for fi5 := 0; fi5 < 14; fi5++ {
                if false {
                    v6 := v2["hello"]
                    var v7 string = v6
                    if ((v4 | v4) < (v3 / ((42) | 1))) {
                        fmt.Println(v7, v4, v7)
                        var v8 uint8 = 127
                        v4 += v3
                        _ = v8
                    } else {
                        var v9 float32 = 0.0
                        v7 = v6
                        _ = v9
                    }
                    _ = v6
                    _ = v7
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        switch (-(fi5)) {
                        case 32768:
                            if true {
                                switch fi5 {
                                case 846:
                                    for fi10 := 0; fi10 < 20; fi10++ {
                                        v11 := GenS1{F0: 1}
                                        v4 = 13
                                        _ = fi10
                                        _ = v11
                                    }
                                }
                            }
                        case 585:
                            v1 = (float64((0.0 * 1.5)) + 0.0)
                        case -752:
                            v4 = -123
                        }
                    case "":
                        fmt.Println(v4)
                    case "a":
                        var v12 int = (fi5 >> ((18446744073709551615) & 63))
                        _ = v12
                    default:
                        switch "hello" {
                        case "":
                            if false {
                                for fi13 := 0; fi13 < 13; fi13++ {
                                    delete(v2, "")
                                    _ = fi13
                                }
                                var v14 string = ""
                                for fi15 := 0; fi15 < 17; fi15++ {
                                    var v16 int8 = (v3 % ((v3) | 1))
                                    v17 := GenS1{F0: fi15}
                                    if true {
                                        var v18 int8 = 127
                                        if (0 != 0) {
                                            v19 := GenS1{F0: fi15}
                                            fmt.Println(v14)
                                            if true {
                                                if true {
                                                    if true {
                                                        var v20 uint64 = 4294967295
                                                        var v21 float32 = 0.0
                                                        _ = v20
                                                        _ = v21
                                                    } else {
                                                        switch (^(fi15)) {
                                                        case 4294967295:
                                                            var v22 int64 = -1
                                                            var v23 int32 = 42
                                                            _ = v22
                                                            _ = v23
                                                        case 474:
                                                            v4 = (v18 * v3)
                                                        case 0:
                                                            v16 -= v3
                                                        default:
                                                            var v24 uint16 = 898
                                                            _ = v24
                                                        }
                                                    }
                                                    var v25 int64 = 9223372036854775807
                                                    _ = v25
                                                }
                                            } else {
                                                v26 := GenS2{F0: 1000}
                                                var v27 float64 = v1
                                                v27 += 1e10
                                                _ = v26
                                                _ = v27
                                            }
                                            fmt.Println(fi15, v3)
                                            _ = v19
                                        }
                                        _ = v18
                                    } else {
                                        var v28 bool = false
                                        v28 = v28
                                        _ = v28
                                    }
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                }
                                _ = v14
                            } else {
                                var v29 uint = 65535
                                switch fi5 {
                                case 2147483647:
                                    for fi30 := 0; fi30 < 19; fi30++ {
                                        if false {
                                            if true {
                                                v31 := []int{fi30, (fi5 / ((-792) | 1)), fi30}
                                                fmt.Println(v29)
                                                v1 = ((3.14159 + (v1 - 3.14159)) + (0.1 * 1.5))
                                                v32 := GenS1{F0: fi5}
                                                delete(v2, ("hello" + (("a" + ("\n\t" + "hello")) + "foo bar")))
                                                v2[("" + "foo bar")] = ("" + "hello")
                                                _ = v31
                                                _ = v32
                                            } else {
                                                v33 := GenS2{F0: -929}
                                                _ = v33
                                            }
                                            v34 := []int{fi30, 127, 32767}
                                            if (!(0 >= 0)) {
                                                switch fi30 {
                                                case 128:
                                                    if (!false) {
                                                        if false {
                                                            if len(v34) >= 2 {
                                                                v34 = v34[1:len(v34)]
                                                            }
                                                        } else {
                                                            v29 -= v29
                                                            fmt.Println(v29, v1)
                                                        }
                                                    }
                                                default:
                                                    var v35 uint8 = 127
                                                    _ = v35
                                                }
                                            }
                                            _ = v34
                                        } else {
                                            _ = len(v2)
                                            fmt.Println(v29)
                                        }
                                        v36 := v2["hello"]
                                        _ = fi30
                                        _ = v36
                                    }
                                case 42:
                                    for fi37 := 0; fi37 < 14; fi37++ {
                                        var v38 bool = true
                                        var v39 int16 = 400
                                        _ = fi37
                                        _ = v38
                                        _ = v39
                                    }
                                case 1:
                                    for fi40 := 0; fi40 < 4; fi40++ {
                                        v1 = float64(1e10)
                                        delete(v2, "foo bar")
                                        var v41 int64 = 42
                                        _ = fi40
                                        _ = v41
                                    }
                                }
                                _ = v29
                            }
                        case "\n\t":
                            v42 := GenS2{F0: 42}
                            _ = v42
                        case "a":
                            fmt.Println(v1, v3)
                        default:
                            switch fi5 {
                            case -1:
                                v43 := []int{fi5, (fi5 - (fi5 + fi5))}
                                _ = v43
                            default:
                                for fi44 := 0; fi44 < 10; fi44++ {
                                    if (!false) {
                                        if true {
                                            v3 += (v3 ^ 100)
                                            v45 := GenS2{F0: 414}
                                            v46 := GenS2{F0: 42}
                                            _ = v45
                                            _ = v46
                                        } else {
                                            if (154 > -9223372036854775808) {
                                                switch ("a" + (("\n\t" + "") + ("hello" + "foo bar"))) {
                                                case "a":
                                                    switch 272 {
                                                    case -872:
                                                        switch ("a" + "\n\t") {
                                                        case "":
                                                            if false {
                                                                v47 := GenS2{F0: 742}
                                                                fmt.Println(fi44)
                                                                delete(v2, "\n\t")
                                                                _ = v47
                                                            } else {
                                                                var v48 int = 932
                                                                v49 := v2[("hello" + "")]
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                        case "a":
                                                            switch "" {
                                                            case "a":
                                                                v50 := GenS3{F0: "\n\t"}
                                                                _ = v50
                                                            case "hello":
                                                                v51 := GenS4{F0: 919, F1: (0 != 125), F2: 0}
                                                                _ = v51
                                                            case "\n\t":
                                                                switch fi44 {
                                                                case -2147483648:
                                                                    var v52 int32 = -2147483648
                                                                    _ = v52
                                                                case -264:
                                                                    var v53 uint = 1
                                                                    _ = v53
                                                                default:
                                                                    if true {
                                                                        v3 = v3
                                                                        var v54 uint = 1
                                                                        switch (fi44 - (fi44 >> ((1) & 63))) {
                                                                        case 256:
                                                                            var v55 int8 = (v4 * (v3 & (v4 - 100)))
                                                                            _ = v55
                                                                        case 32768:
                                                                            _ = len(v2)
                                                                        }
                                                                        _ = v54
                                                                    }
                                                                }
                                                            }
                                                        case "\n\t":
                                                            v3 = v3
                                                            v56 := []int{fi5, fi5, (fi44 + 42), fi5}
                                                            _ = v56
                                                        default:
                                                            v2["foo bar"] = (("\n\t" + "\n\t") + "foo bar")
                                                        }
                                                    case 4294967295:
                                                        v1 += -2.25
                                                    default:
                                                        if (false || (-32768 != 32767)) {
                                                            v57 := GenS5{F0: v3}
                                                            _ = v57
                                                        }
                                                        var v58 uint32 = 65535
                                                        fmt.Println(v4, fi5)
                                                        _ = v58
                                                    }
                                                case "hello":
                                                    v59 := []int{(fi5 + 793), 42}
                                                    _ = v59
                                                case "":
                                                    v2["hello"] = "\n\t"
                                                default:
                                                    var v60 uint64 = 143
                                                    if false {
                                                        v3 += v4
                                                        v60 -= (v60 + 2)
                                                    } else {
                                                        v60 -= (v60 - v60)
                                                        switch "hello" {
                                                        case "":
                                                            v60 = v60
                                                        case "hello":
                                                            if (true || true) {
                                                                var v61 float64 = float64(-2.25)
                                                                v62 := GenS6{F0: 413, F1: false, F2: 0}
                                                                if ((!false) && ((v60 >> ((528) & 63)) > v60)) {
                                                                    switch fi44 {
                                                                    case 42:
                                                                        var v63 bool = true
                                                                        _ = v63
                                                                    case 834:
                                                                        v64 := GenS2{F0: 0}
                                                                        _ = v64
                                                                    default:
                                                                        v2[(("a" + (((((("a" + ("" + "")) + (("\n\t" + "a") + "a")) + ("hello" + "")) + ("foo bar" + "foo bar")) + "a") + "a")) + (((("a" + "foo bar") + (((("a" + "hello") + ((("" + "\n\t") + "a") + ("" + ("\n\t" + "hello")))) + "foo bar") + ("hello" + ("" + ("hello" + "a"))))) + "a") + ("a" + "foo bar")))] = "\n\t"
                                                                    }
                                                                    fmt.Println(fi5, v61, fi44)
                                                                }
                                                                _ = v61
                                                                _ = v62
                                                            } else {
                                                                var v65 uint16 = 65535
                                                                _ = v65
                                                            }
                                                        case "foo bar":
                                                            switch fi44 {
                                                            case 474:
                                                                fmt.Println(v60, fi44)
                                                            }
                                                            switch ("" + "\n\t") {
                                                            case "\n\t":
                                                                if (false && true) {
                                                                    v66 := []int{42, fi44, (fi5 + fi5), fi5}
                                                                    if len(v66) >= 2 {
                                                                        v66 = v66[1:len(v66)]
                                                                    }
                                                                    if false {
                                                                        v4 += -128
                                                                        v4 = v4
                                                                        fmt.Println(v4, fi44)
                                                                        var v67 int64 = 1
                                                                        _ = v67
                                                                    } else {
                                                                        switch "" {
                                                                        case "hello":
                                                                            v1 = (((1.5 + (v1 - ((v1 / ((v1 * float64(0.1)) / (v1 / v1))) + 1e10))) * 1e10) - float64((((1e10 * 1e10) * (-2.25 / 1e10)) - 0.1)))
                                                                            v2[((("hello" + "foo bar") + "") + "\n\t")] = "\n\t"
                                                                        default:
                                                                            if (!true) {
                                                                                v2[("" + (("a" + "\n\t") + "foo bar"))] = "a"
                                                                                fmt.Println(fi5, fi5, v4)
                                                                            }
                                                                        }
                                                                        if true {
                                                                            v3 += (v3 & (v3 >> ((255) & 63)))
                                                                            switch fi5 {
                                                                            case -1:
                                                                                fmt.Println(v3)
                                                                            case 4294967295:
                                                                                v68 := GenS5{F0: v4}
                                                                                v4 += v4
                                                                                switch ("foo bar" + "a") {
                                                                                case "a":
                                                                                    var v69 float64 = v1
                                                                                    _ = v69
                                                                                case "foo bar":
                                                                                    fmt.Println(v60)
                                                                                }
                                                                                _ = v68
                                                                            default:
                                                                                v1 = 0.1
                                                                            }
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                fmt.Println(v1, v1, v60)
                                                                            case "":
                                                                                fmt.Println(v4, fi5)
                                                                            default:
                                                                                v4 = 47
                                                                            }
                                                                            var v70 float64 = v1
                                                                            if (!(v3 <= (v3 - v3))) {
                                                                                v3 -= 96
                                                                            }
                                                                            switch fi5 {
                                                                            case -9223372036854775808:
                                                                                v60 -= 4294967295
                                                                            case -715:
                                                                                switch fi5 {
                                                                                case 1:
                                                                                    v1 = 0.1
                                                                                    fmt.Println(v4, v1, v70)
                                                                                    if (!false) {
                                                                                        switch "" {
                                                                                        case "a":
                                                                                            var v71 uint8 = 0
                                                                                            _ = v71
                                                                                        case "hello":
                                                                                            if (!false) {
                                                                                                var v72 int16 = 928
                                                                                                fmt.Println(v70)
                                                                                                _ = v72
                                                                                            }
                                                                                        case "":
                                                                                            if (!(1 <= 18446744073709551615)) {
                                                                                                switch fi44 {
                                                                                                case 128:
                                                                                                    if len(v66) >= 2 {
                                                                                                        v66 = v66[1:len(v66)]
                                                                                                    }
                                                                                                case 9223372036854775807:
                                                                                                    if true {
                                                                                                        fmt.Println(v3)
                                                                                                        switch "hello" {
                                                                                                        case "a":
                                                                                                            switch fi44 {
                                                                                                            case 2147483647:
                                                                                                                if (4294967295 >= 475) {
                                                                                                                    if (-3 >= 0) {
                                                                                                                        var v73 int64 = -9223372036854775808
                                                                                                                        _ = v73
                                                                                                                    }
                                                                                                                    var v74 uint16 = 0
                                                                                                                    _ = v74
                                                                                                                } else {
                                                                                                                    fmt.Println(v1, v70)
                                                                                                                }
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            v60 = (v60 * 32)
                                                                                                        case "\n\t":
                                                                                                            switch (fi5 % (((fi44 % ((fi44) | 1))) | 1)) {
                                                                                                            case 704:
                                                                                                                var v75 uint = 42
                                                                                                                _ = v75
                                                                                                            default:
                                                                                                                switch "\n\t" {
                                                                                                                case "foo bar":
                                                                                                                    v60 += (v60 & 0)
                                                                                                                case "a":
                                                                                                                    switch "hello" {
                                                                                                                    case "":
                                                                                                                        if (!(42 >= 1)) {
                                                                                                                            var v76 uint = 0
                                                                                                                            v76 += (v76 - v76)
                                                                                                                            v77 := GenS3{F0: "a"}
                                                                                                                            _ = v76
                                                                                                                            _ = v77
                                                                                                                        }
                                                                                                                    case "hello":
                                                                                                                        if ((fi44 + (fi5 % ((fi44) | 1))) < (fi44 % (((fi5 - fi5)) | 1))) {
                                                                                                                            fmt.Println(v3, v4, v60)
                                                                                                                        } else {
                                                                                                                            if ((!(((false || (true || (128 > 42))) && (true || (false && ((v4 | v3) > -19)))) || false)) || ((v3 * (v4 >> ((65535) & 63))) <= v3)) {
                                                                                                                                if ((150 != 62) && false) {
                                                                                                                                    if (false || false) {
                                                                                                                                        fmt.Println(v60, fi44, fi5)
                                                                                                                                        v2[(("a" + "foo bar") + "")] = ""
                                                                                                                                        var v78 uint = 1
                                                                                                                                        _ = v78
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    switch 9223372036854775807 {
                                                                                                                                    case -992:
                                                                                                                                        v1 += v70
                                                                                                                                    default:
                                                                                                                                        var v79 float64 = v1
                                                                                                                                        _ = v79
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                            if (!true) {
                                                                                                                                switch fi5 {
                                                                                                                                case 2147483647:
                                                                                                                                    if len(v66) > 0 {
                                                                                                                                        v66[0] = (^(fi44))
                                                                                                                                    }
                                                                                                                                    v60 += v60
                                                                                                                                case -220:
                                                                                                                                    v60 = v60
                                                                                                                                }
                                                                                                                                var v80 float32 = (1e10 + float32((1e10 * v70)))
                                                                                                                                var v81 float32 = 1.5
                                                                                                                                _ = v80
                                                                                                                                _ = v81
                                                                                                                            }
                                                                                                                            switch ("" + "") {
                                                                                                                            case "hello":
                                                                                                                                v4 = v4
                                                                                                                            case "":
                                                                                                                                fmt.Println(v70, fi44, v1)
                                                                                                                            }
                                                                                                                        }
                                                                                                                    case "\n\t":
                                                                                                                        var v82 float32 = ((1.5 * 1e10) * 1.5)
                                                                                                                        _ = v82
                                                                                                                    default:
                                                                                                                        v83 := GenS7{F0: ("hello" + "a"), F1: 255}
                                                                                                                        _ = v83
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    fmt.Println(v3)
                                                                                                                default:
                                                                                                                    var v84 uint = 42
                                                                                                                    v60 -= (v60 / (((^(v60))) | 1))
                                                                                                                    v84 += (v84 % ((v84) | 1))
                                                                                                                    fmt.Println(v4, fi5)
                                                                                                                    _ = len(v66)
                                                                                                                    switch (("foo bar" + ((("foo bar" + "a") + "\n\t") + "a")) + "a") {
                                                                                                                    case "":
                                                                                                                        v1 = float64((3.14159 / 1.0))
                                                                                                                    case "foo bar":
                                                                                                                        if true {
                                                                                                                            if true {
                                                                                                                                v4 = v3
                                                                                                                                v85 := GenS4{F0: 42, F1: false, F2: -1}
                                                                                                                                _ = v85
                                                                                                                            }
                                                                                                                            fmt.Println(v3, v60, v1)
                                                                                                                        }
                                                                                                                    case "hello":
                                                                                                                        _ = len(v66)
                                                                                                                    }
                                                                                                                    v60 = (v60 * (^(v60)))
                                                                                                                    _ = v84
                                                                                                                }
                                                                                                            }
                                                                                                            v3 += (v4 + (v3 % ((-96) | 1)))
                                                                                                            if len(v66) >= 2 {
                                                                                                                v66 = v66[1:len(v66)]
                                                                                                            }
                                                                                                            v3 += v4
                                                                                                            fmt.Println(v70, fi44, fi44)
                                                                                                            v60 -= 238
                                                                                                            if len(v66) > 0 {
                                                                                                                v66[0] = -1
                                                                                                            }
                                                                                                            v66 = append(v66, fi5)
                                                                                                        default:
                                                                                                            v86 := GenS8{F0: 1, F1: -766, F2: 0}
                                                                                                            _ = v86
                                                                                                        }
                                                                                                    } else {
                                                                                                        switch fi44 {
                                                                                                        case 1:
                                                                                                            _ = len(v66)
                                                                                                        case 2147483647:
                                                                                                            v60 -= (^(v60))
                                                                                                        default:
                                                                                                            fmt.Println(v3)
                                                                                                            v70 = v70
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    var v87 string = "foo bar"
                                                                                                    _ = v87
                                                                                                }
                                                                                                if (1 < 2) {
                                                                                                    v88 := GenS7{F0: ("\n\t" + "foo bar"), F1: 4294967295}
                                                                                                    v60 = (v60 / (((v60 ^ v60)) | 1))
                                                                                                    _ = v88
                                                                                                } else {
                                                                                                    fmt.Println(v60, fi5)
                                                                                                    var v89 uint8 = 0
                                                                                                    _ = v89
                                                                                                }
                                                                                                if (!(-2.25 == (1e10 / (((v1 * v1) - float64(0.0)) - ((-2.25 + v1) / -2.25))))) {
                                                                                                    v90 := GenS2{F0: 1000}
                                                                                                    fmt.Println(fi5, v70)
                                                                                                    _ = v90
                                                                                                } else {
                                                                                                    if true {
                                                                                                        switch ("" + "") {
                                                                                                        case "foo bar":
                                                                                                            var v91 int8 = (v4 - v3)
                                                                                                            _ = v91
                                                                                                        default:
                                                                                                            switch (^(fi5)) {
                                                                                                            case 76:
                                                                                                                switch 32768 {
                                                                                                                case 29:
                                                                                                                    var v92 uint32 = 4294967295
                                                                                                                    _ = v92
                                                                                                                case 127:
                                                                                                                    fmt.Println(v1)
                                                                                                                case 0:
                                                                                                                    _ = len(v2)
                                                                                                                    var v93 uint8 = 128
                                                                                                                    _ = v93
                                                                                                                }
                                                                                                            case 4294967295:
                                                                                                                fmt.Println(fi5, v70)
                                                                                                            case 2147483647:
                                                                                                                v2["foo bar"] = ("hello" + "\n\t")
                                                                                                            default:
                                                                                                                var v94 int16 = 32767
                                                                                                                _ = v94
                                                                                                            }
                                                                                                        }
                                                                                                        v1 += 0.0
                                                                                                        var v95 int64 = -1
                                                                                                        _ = v95
                                                                                                    } else {
                                                                                                        _ = len(v2)
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            v1 = v1
                                                                                        }
                                                                                    }
                                                                                case 127:
                                                                                    if false {
                                                                                        v96 := GenS9{F0: 72, F1: (fi5 >> ((42) & 63))}
                                                                                        var v97 int16 = int16(fi5)
                                                                                        var v98 float32 = 0.0
                                                                                        _ = v96
                                                                                        _ = v97
                                                                                        _ = v98
                                                                                    }
                                                                                default:
                                                                                    v60 -= 2
                                                                                }
                                                                            default:
                                                                                var v99 float64 = float64(0.0)
                                                                                _ = v99
                                                                            }
                                                                            v100 := GenS4{F0: 0, F1: true, F2: -2147483648}
                                                                            _ = v100
                                                                            _ = v70
                                                                        } else {
                                                                            switch fi44 {
                                                                            case 32767:
                                                                                if (true && (0 < -2147483648)) {
                                                                                    switch "hello" {
                                                                                    case "foo bar":
                                                                                        v4 -= 42
                                                                                    case "":
                                                                                        v66 = append(v66, 127)
                                                                                    case "a":
                                                                                        if false {
                                                                                            if false {
                                                                                                v101 := GenS7{F0: "hello", F1: 255}
                                                                                                _ = v101
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        if false {
                                                                                            var v102 uint64 = (v60 * (v60 & (v60 & (v60 % (((v60 / ((4294967295) | 1))) | 1)))))
                                                                                            _ = v102
                                                                                        } else {
                                                                                            switch (fi5 & (fi44 >> ((4294967295) & 63))) {
                                                                                            case -2147483648:
                                                                                                v103 := v2["a"]
                                                                                                _ = v103
                                                                                            case 644:
                                                                                                v3 = 0
                                                                                            }
                                                                                            switch 42 {
                                                                                            case 9223372036854775807:
                                                                                                v1 = (((v1 + v1) * v1) - v1)
                                                                                            }
                                                                                            v104 := GenS6{F0: 1, F1: (32767 != 1), F2: 0}
                                                                                            _ = v104
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(fi44, fi44, v1)
                                                                                }
                                                                            case 629:
                                                                                fmt.Println(v60, fi44)
                                                                            case -2147483648:
                                                                                if true {
                                                                                    fmt.Println(fi5, fi44, v3)
                                                                                    v105 := GenS2{F0: 590}
                                                                                    _ = v105
                                                                                }
                                                                                v106 := GenS7{F0: "hello", F1: 2}
                                                                                v107 := GenS10{F0: 328, F1: v60, F2: "hello"}
                                                                                _ = v106
                                                                                _ = v107
                                                                            default:
                                                                                v3 += (-(v3))
                                                                            }
                                                                            if ((((1.5 + v1) * ((float64(float32(3.14159)) * 3.14159) * float64(((-2.25 + 3.14159) - (3.14159 / (float32(v1) + 3.14159)))))) >= (1.5 - (1.5 * v1))) && ((1.5 + (1e10 * 1.5)) <= (0.1 * float32((v1 + (v1 - (1.5 - ((v1 / (float64(float32(float64(float32(v1)))) + v1)) + ((v1 / 1.0) - (0.1 + -2.25)))))))))) {
                                                                                switch fi44 {
                                                                                case -64:
                                                                                    var v108 uint64 = v60
                                                                                    _ = v108
                                                                                case 128:
                                                                                    v1 = float64(1e10)
                                                                                }
                                                                                fmt.Println(v4)
                                                                                var v109 float64 = (v1 / 0.1)
                                                                                if len(v66) >= 2 {
                                                                                    v66 = v66[1:len(v66)]
                                                                                }
                                                                                v110 := GenS11{F0: 255, F1: 1, F2: 707}
                                                                                _ = v109
                                                                                _ = v110
                                                                            }
                                                                        }
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            var v111 int32 = 1
                                                                            _ = v111
                                                                        case "":
                                                                            fmt.Println(v3)
                                                                        default:
                                                                            v3 += (v4 ^ 0)
                                                                        }
                                                                    }
                                                                    v60 = v60
                                                                    _ = v66
                                                                } else {
                                                                    v3 += (v4 % ((v3) | 1))
                                                                }
                                                                v112 := []int{(^(fi44))}
                                                                var v113 int8 = -128
                                                                _ = v112
                                                                _ = v113
                                                            case "a":
                                                                v4 += 1
                                                            }
                                                            switch "" {
                                                            case "foo bar":
                                                                if (!true) {
                                                                    if false {
                                                                        v3 = (v3 * 81)
                                                                    }
                                                                } else {
                                                                    var v114 string = "a"
                                                                    switch "" {
                                                                    case "a":
                                                                        var v115 uint8 = uint8(v60)
                                                                        _ = v115
                                                                    }
                                                                    switch v114 {
                                                                    case "a":
                                                                        fmt.Println(v114)
                                                                    case "":
                                                                        if (float32(((3.14159 * ((3.14159 * 1e10) - (float64((((1.5 / 0.1) / ((float32(0.1) + 1e10) + 1.5)) - ((1.5 * 3.14159) / 1.0))) / -2.25))) * v1)) <= 1e10) {
                                                                            v116 := GenS7{F0: (v114 + (v114 + v114)), F1: 18446744073709551615}
                                                                            if false {
                                                                                v117 := GenS7{F0: v114, F1: 0}
                                                                                _ = v117
                                                                            }
                                                                            v118 := []int{fi44}
                                                                            _ = v116
                                                                            _ = v118
                                                                        }
                                                                    }
                                                                    _ = v114
                                                                }
                                                            case "a":
                                                                switch fi5 {
                                                                case -1:
                                                                    v4 += v3
                                                                }
                                                            case "":
                                                                v119 := GenS2{F0: 42}
                                                                _ = v119
                                                            default:
                                                                v120 := GenS12{F0: 2147483647, F1: 392}
                                                                _ = v120
                                                            }
                                                        }
                                                        v1 = 1e10
                                                    }
                                                    v60 -= v60
                                                    _ = v60
                                                }
                                                v4 += v4
                                                v121 := GenS13{F0: -2147483648, F1: 65535}
                                                _ = v121
                                            }
                                        }
                                    }
                                    _ = fi44
                                }
                            }
                        }
                    }
                }
                _ = fi5
            }
        } else {
            v2[("foo bar" + "")] = ("foo bar" + ("hello" + ("a" + "a")))
        }
        v122 := GenS8{F0: 1, F1: -855, F2: 127}
        var v123 uint = 18446744073709551615
        _ = v122
        _ = v123
    }
    v124 := []int{-328, 32768, 32768, 4294967295}
    v1 -= (((v1 * 1e10) * float64(3.14159)) / v1)
    fmt.Println(v3, v1, v1)
    for fi125 := 0; fi125 < 3; fi125++ {
        switch (^(fi125)) {
        case 32768:
            v126 := GenS8{F0: 255, F1: -1, F2: 255}
            _ = v126
        case 1:
            v1 = v1
        }
        _ = fi125
    }
    v3 += 100
    var v127 int16 = 0
    v128 := GenS3{F0: ("" + "hello")}
    var v129 int64 = 9223372036854775807
    v127 += v127
    var v130 string = "foo bar"
    v130 = v130
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v124:", len(v124), v124)
    fmt.Println("v127:", v127)
    fmt.Println("v129:", v129)
    fmt.Println("v130:", v130)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v124
    _ = v127
    _ = v128
    _ = v129
    _ = v130
    _ = v2
    _ = v3
    _ = v4
}
