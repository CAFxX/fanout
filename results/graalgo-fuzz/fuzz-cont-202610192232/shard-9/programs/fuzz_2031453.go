package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 int32; F2 uint16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool; F1 bool }
type GenS11 struct { F0 bool; F1 string }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 bool; F1 string; F2 int8 }
type GenS15 struct { F0 int16 }
func genFn0(p0 int32, p1 uint8) int8 {
    p0 = 42
    switch 128 {
    case 848:
        p1 -= (p1 & 1)
    default:
        v5 := []int{0}
        _ = v5
    }
    var v6 int8 = 100
    for fi7 := 0; fi7 < 15; fi7++ {
        if true {
            v8 := []int{(fi7 + (-(fi7))), (fi7 + (-(fi7)))}
            _ = v8
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(p0, fi7, fi9)
            var v10 float64 = (-2.25 + 1e10)
            fmt.Println(fi7, p1)
            _ = fi9
            _ = v10
        }
        _ = fi7
    }
    var v11 float32 = 0.0
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return (v6 + 42)
}
func genFn1() string {
    v12 := []int{}
    if (false || (false && (!false))) {
        if false {
            switch "" {
            case "hello":
                if true {
                    switch ("\n\t" + "hello") {
                    case "\n\t":
                        if true {
                            switch -889 {
                            case -1:
                                if (true && ((float64(-2.25) / -2.25) > -2.25)) {
                                    for fi13 := 0; fi13 < 16; fi13++ {
                                        if len(v12) > 0 {
                                            v12[0] = 128
                                        }
                                        v14 := GenS1{F0: 1, F1: 1}
                                        _ = fi13
                                        _ = v14
                                    }
                                } else {
                                    for fi15 := 0; fi15 < 15; fi15++ {
                                        var v16 int16 = 1000
                                        _ = fi15
                                        _ = v16
                                    }
                                }
                            case 2147483647:
                                for fi17 := 0; fi17 < 14; fi17++ {
                                    if true {
                                        v12 = append(v12, (fi17 % ((fi17) | 1)))
                                    } else {
                                        for fi18 := 0; fi18 < 18; fi18++ {
                                            v12 = append(v12, (fi18 | (fi17 | (fi17 * fi18))))
                                            _ = fi18
                                        }
                                        if ((((389 <= 18446744073709551615) || false) || true) || true) {
                                            var v19 uint16 = 0
                                            for fi20 := 0; fi20 < 13; fi20++ {
                                                var v21 bool = true
                                                if ((v21 || v21) && (-9223372036854775808 < 9223372036854775807)) {
                                                    v19 = (^(v19))
                                                } else {
                                                    fmt.Println(v19)
                                                    v22 := GenS1{F0: 332, F1: v19}
                                                    _ = v22
                                                }
                                                v23 := GenS2{F0: "", F1: "a"}
                                                _ = fi20
                                                _ = v21
                                                _ = v23
                                            }
                                            _ = len(v12)
                                            _ = v19
                                        }
                                        var v24 int8 = -128
                                        _ = v24
                                    }
                                    var v25 float32 = 3.14159
                                    v26 := GenS3{F0: -465, F1: 18446744073709551615}
                                    _ = fi17
                                    _ = v25
                                    _ = v26
                                }
                            case -2147483648:
                                var v27 int32 = -2147483648
                                _ = v27
                            default:
                                switch "foo bar" {
                                case "":
                                    switch ("" + "") {
                                    case "":
                                        for fi28 := 0; fi28 < 18; fi28++ {
                                            var v29 int8 = 1
                                            _ = fi28
                                            _ = v29
                                        }
                                    case "hello":
                                        if (291 < 248) {
                                            var v30 int = 32767
                                            var v31 float32 = -2.25
                                            var v32 float64 = (3.14159 - 0.1)
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                        } else {
                                            if (3.14159 != 0.1) {
                                                var v33 int = 112
                                                if true {
                                                    var v34 uint64 = 1
                                                    v35 := GenS1{F0: 0, F1: 989}
                                                    _ = v34
                                                    _ = v35
                                                }
                                                _ = v33
                                            }
                                        }
                                    case "\n\t":
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                    }
                                case "\n\t":
                                    var v36 string = "foo bar"
                                    _ = v36
                                default:
                                    if ((float64(0.0) + 1e10) > (0.1 - ((((3.14159 * 0.0) / float64(((1e10 + -2.25) / 1.5))) - -2.25) / 1.0))) {
                                        var v37 bool = (!false)
                                        v37 = v37
                                        _ = v37
                                    } else {
                                        var v38 float32 = -2.25
                                        if len(v12) > 0 {
                                            v12[0] = 32768
                                        }
                                        _ = v38
                                    }
                                }
                            }
                            if true {
                                var v39 uint = 42
                                fmt.Println(v39)
                                _ = v39
                            }
                            for fi40 := 0; fi40 < 14; fi40++ {
                                var v41 float32 = 3.14159
                                v12 = append(v12, (fi40 + fi40))
                                v42 := GenS3{F0: -472, F1: 91}
                                _ = fi40
                                _ = v41
                                _ = v42
                            }
                        }
                    case "hello":
                        for fi43 := 0; fi43 < 6; fi43++ {
                            v44 := GenS4{F0: -855}
                            if (3.14159 == 1e10) {
                                v45 := GenS4{F0: -633}
                                var v46 int32 = 2147483647
                                _ = v45
                                _ = v46
                            }
                            _ = fi43
                            _ = v44
                        }
                    default:
                        if len(v12) > 0 {
                            v12[0] = 2147483647
                        }
                    }
                    switch -75 {
                    case 165:
                        if true {
                            for fi47 := 0; fi47 < 18; fi47++ {
                                v12 = append(v12, -2147483648)
                                v48 := GenS5{F0: (("" + "") + "foo bar"), F1: 0}
                                if ((-2.25 < ((1e10 * -2.25) * (3.14159 + (-2.25 + 1.5)))) && (!((1.5 + 0.1) < ((-2.25 * 1.5) * float64(1e10))))) {
                                    if false {
                                        for fi49 := 0; fi49 < 18; fi49++ {
                                            var v50 uint32 = 2147483647
                                            v50 -= (v50 & 1)
                                            _ = fi49
                                            _ = v50
                                        }
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                    } else {
                                        var v51 int64 = 839
                                        _ = v51
                                    }
                                    for fi52 := 0; fi52 < 8; fi52++ {
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        v53 := GenS2{F0: "\n\t", F1: "\n\t"}
                                        if false {
                                            switch "hello" {
                                            case "a":
                                                if (!(0.0 > -2.25)) {
                                                    fmt.Println(fi47, fi52)
                                                } else {
                                                    if len(v12) >= 2 {
                                                        v12 = v12[1:len(v12)]
                                                    }
                                                    if false {
                                                        if true {
                                                            v54 := GenS4{F0: -1}
                                                            _ = v54
                                                        }
                                                        v12 = append(v12, 32768)
                                                        var v55 uint16 = 255
                                                        switch (fi52 ^ fi52) {
                                                        case -54:
                                                            v56 := GenS6{F0: ("" + ("" + "\n\t"))}
                                                            fmt.Println(v55, v55, fi52)
                                                            _ = v56
                                                        case 32767:
                                                            if ((-(fi47)) >= (fi52 + fi52)) {
                                                                if (-1 >= 273) {
                                                                    _ = len(v12)
                                                                }
                                                                var v57 uint64 = 330
                                                                v58 := GenS5{F0: ((("\n\t" + "a") + "\n\t") + ("" + (("hello" + ("hello" + "\n\t")) + (("foo bar" + "hello") + ("" + "a"))))), F1: (v57 << ((802) & 63))}
                                                                var v59 int16 = -32768
                                                                _ = v57
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                        case -1:
                                                            if false {
                                                                v55 = (v55 >> ((18446744073709551615) & 63))
                                                                if len(v12) > 0 {
                                                                    v12[0] = fi52
                                                                }
                                                            } else {
                                                                _ = len(v12)
                                                                if false {
                                                                    switch (fi52 * (^(fi47))) {
                                                                    case 114:
                                                                        v55 = v55
                                                                    case 0:
                                                                        v12 = append(v12, (fi52 - fi52))
                                                                    case 2147483647:
                                                                        var v60 uint8 = 128
                                                                        _ = v60
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        if true {
                                                            v55 = (v55 << ((4294967295) & 63))
                                                            v12 = append(v12, fi47)
                                                            var v61 int16 = -846
                                                            v61 -= (-(v61))
                                                            var v62 uint8 = 128
                                                            switch "" {
                                                            case "foo bar":
                                                                fmt.Println(fi47)
                                                            case "hello":
                                                                if false {
                                                                    fmt.Println(fi47, v62, fi52)
                                                                } else {
                                                                    var v63 float64 = (1.5 * (0.1 * float64(0.0)))
                                                                    _ = v63
                                                                }
                                                            }
                                                            var v64 string = ""
                                                            if len(v12) > 0 {
                                                                v12[0] = (-(fi47))
                                                            }
                                                            _ = v61
                                                            _ = v62
                                                            _ = v64
                                                        }
                                                        _ = v55
                                                    }
                                                }
                                            case "\n\t":
                                                var v65 uint64 = 4294967295
                                                v65 -= v65
                                                var v66 uint = 0
                                                _ = v65
                                                _ = v66
                                            default:
                                                var v67 int16 = -32768
                                                _ = v67
                                            }
                                            var v68 uint8 = 255
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            _ = v68
                                        } else {
                                            var v69 uint64 = 18446744073709551615
                                            var v70 int8 = 0
                                            switch "" {
                                            case "\n\t":
                                                v12 = append(v12, 340)
                                            case "":
                                                v12 = append(v12, fi47)
                                            default:
                                                v69 = v69
                                            }
                                            var v71 int16 = 42
                                            v12 = append(v12, 1)
                                            _ = v69
                                            _ = v70
                                            _ = v71
                                        }
                                        _ = fi52
                                        _ = v53
                                    }
                                    if (false && false) {
                                        switch (fi47 ^ (fi47 >> ((138) & 63))) {
                                        case 128:
                                            var v72 bool = false
                                            _ = v72
                                        case -1:
                                            var v73 float32 = 3.14159
                                            _ = v73
                                        case 9223372036854775807:
                                            if true {
                                                _ = len(v12)
                                                if len(v12) >= 2 {
                                                    v12 = v12[1:len(v12)]
                                                }
                                            } else {
                                                var v74 int16 = 1000
                                                v12 = append(v12, (-(fi47)))
                                                v74 = v74
                                                _ = v74
                                            }
                                        default:
                                            var v75 float32 = 0.0
                                            _ = v75
                                        }
                                    } else {
                                        if (42 != 525) {
                                            if false {
                                                switch (("a" + "") + ((("hello" + "foo bar") + "a") + "hello")) {
                                                case "a":
                                                    v76 := GenS6{F0: "foo bar"}
                                                    _ = v76
                                                default:
                                                    var v77 float64 = -2.25
                                                    _ = v77
                                                }
                                                _ = len(v12)
                                            } else {
                                                for fi78 := 0; fi78 < 3; fi78++ {
                                                    v79 := GenS7{F0: -870, F1: 1, F2: 886}
                                                    switch (("hello" + "foo bar") + "a") {
                                                    case "":
                                                        if true {
                                                            if false {
                                                                var v80 int8 = 127
                                                                v80 = (v80 + genFn0(-2147483648, 128))
                                                                v80 = 1
                                                                _ = v80
                                                            } else {
                                                                var v81 int = fi78
                                                                if len(v12) >= 2 {
                                                                    v12 = v12[1:len(v12)]
                                                                }
                                                                _ = v81
                                                            }
                                                        } else {
                                                            var v82 uint32 = 2147483648
                                                            _ = len(v12)
                                                            _ = len(v12)
                                                            _ = v82
                                                        }
                                                    case "a":
                                                        fmt.Println(fi78)
                                                    }
                                                    _ = fi78
                                                    _ = v79
                                                }
                                                fmt.Println(fi47)
                                            }
                                            _ = len(v12)
                                        }
                                        var v83 uint8 = 128
                                        var v84 bool = (!false)
                                        _ = v83
                                        _ = v84
                                    }
                                } else {
                                    var v85 uint64 = 4294967295
                                    v85 += (v85 / ((v85) | 1))
                                    for fi86 := 0; fi86 < 5; fi86++ {
                                        if (127 <= (^(fi86))) {
                                            var v87 float64 = float64(-2.25)
                                            v87 = (-2.25 / 1.0)
                                            var v88 int16 = 1
                                            _ = v87
                                            _ = v88
                                        } else {
                                            var v89 uint8 = 156
                                            var v90 int16 = -32768
                                            v90 = 0
                                            _ = v89
                                            _ = v90
                                        }
                                        _ = fi86
                                    }
                                    _ = v85
                                }
                                _ = fi47
                                _ = v48
                            }
                            if true {
                                for fi91 := 0; fi91 < 20; fi91++ {
                                    fmt.Println(fi91)
                                    _ = fi91
                                }
                                var v92 string = "hello"
                                switch ((((v92 + "foo bar") + ("hello" + ("hello" + ""))) + "") + v92) {
                                case "a":
                                    fmt.Println(v92)
                                case "":
                                    if false {
                                        var v93 string = v92
                                        _ = v93
                                    } else {
                                        fmt.Println(v92)
                                        switch -2147483648 {
                                        case 4294967295:
                                            var v94 int32 = 1
                                            _ = v94
                                        default:
                                            if ((1 > -1) && (183 < 2)) {
                                                if len(v12) > 0 {
                                                    v12[0] = 4294967295
                                                }
                                                switch (("" + "\n\t") + "hello") {
                                                case "":
                                                    switch 127 {
                                                    case 32768:
                                                        v95 := GenS7{F0: 1, F1: 0, F2: 0}
                                                        _ = v95
                                                    case -2147483648:
                                                        if false {
                                                            var v96 int64 = 9223372036854775807
                                                            _ = v96
                                                        }
                                                    }
                                                default:
                                                    switch v92 {
                                                    case "a":
                                                        var v97 float64 = -2.25
                                                        _ = v97
                                                    case "\n\t":
                                                        var v98 uint16 = 32767
                                                        _ = v98
                                                    case "foo bar":
                                                        v92 = ("a" + v92)
                                                    default:
                                                        v12 = append(v12, -547)
                                                    }
                                                }
                                                var v99 int = -387
                                                _ = v99
                                            } else {
                                                v12 = append(v12, -537)
                                                for fi100 := 0; fi100 < 7; fi100++ {
                                                    for fi101 := 0; fi101 < 9; fi101++ {
                                                        if ((42 > 2) && true) {
                                                            switch v92 {
                                                            case "a":
                                                                if true {
                                                                    v12 = append(v12, (fi100 >> ((711) & 63)))
                                                                    if true {
                                                                        var v102 int16 = 32767
                                                                        var v103 uint8 = 71
                                                                        _ = v102
                                                                        _ = v103
                                                                    } else {
                                                                        var v104 string = v92
                                                                        v92 = "foo bar"
                                                                        if false {
                                                                            switch fi100 {
                                                                            case 0:
                                                                                var v105 bool = (!true)
                                                                                _ = v105
                                                                            case 1:
                                                                                var v106 string = (v92 + "foo bar")
                                                                                _ = v106
                                                                            case 321:
                                                                                v104 = v92
                                                                            default:
                                                                                if (75 <= 1) {
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        if len(v12) > 0 {
                                                                                            v12[0] = (fi101 ^ -2147483648)
                                                                                        }
                                                                                    case "\n\t":
                                                                                        switch fi100 {
                                                                                        case -48:
                                                                                            var v107 bool = false
                                                                                            _ = v107
                                                                                        case 9223372036854775807:
                                                                                            if false {
                                                                                                if (false && (0.0 > (3.14159 * ((1e10 + 0.0) * 1.5)))) {
                                                                                                    var v108 int16 = 1
                                                                                                    _ = v108
                                                                                                } else {
                                                                                                    v12 = append(v12, (fi100 >> ((2) & 63)))
                                                                                                    v92 = (v92 + v92)
                                                                                                    v92 = v92
                                                                                                }
                                                                                                switch v92 {
                                                                                                case "a":
                                                                                                    var v109 float64 = 3.14159
                                                                                                    _ = v109
                                                                                                case "hello":
                                                                                                    v104 = "foo bar"
                                                                                                case "":
                                                                                                    var v110 int16 = 1
                                                                                                    _ = v110
                                                                                                }
                                                                                            }
                                                                                        case -797:
                                                                                            if len(v12) > 0 {
                                                                                                v12[0] = (fi100 ^ 32767)
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(v104)
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    if (!(!true)) {
                                                                                        var v111 int = (fi100 / ((-1) | 1))
                                                                                        var v112 bool = (!(42 < 42))
                                                                                        _ = v111
                                                                                        _ = v112
                                                                                    }
                                                                                    fmt.Println(fi100, fi101)
                                                                                }
                                                                            }
                                                                            switch (("foo bar" + ("\n\t" + (v92 + v92))) + "a") {
                                                                            case "foo bar":
                                                                                v104 = (v104 + v92)
                                                                            default:
                                                                                _ = len(v12)
                                                                            }
                                                                            if true {
                                                                                v12 = append(v12, 32768)
                                                                                if true {
                                                                                    if false {
                                                                                        var v113 int16 = -1
                                                                                        v113 = -1
                                                                                        _ = v113
                                                                                    } else {
                                                                                        var v114 uint16 = 696
                                                                                        if false {
                                                                                            v115 := GenS7{F0: 98, F1: -2147483648, F2: (v114 / (((^(v114))) | 1))}
                                                                                            _ = v115
                                                                                        } else {
                                                                                            v104 = v104
                                                                                            v12 = append(v12, -2147483648)
                                                                                        }
                                                                                        _ = v114
                                                                                    }
                                                                                    v92 = (v104 + "a")
                                                                                }
                                                                                var v116 bool = (true || false)
                                                                                _ = v116
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v92, v92)
                                                                        }
                                                                        _ = v104
                                                                    }
                                                                    switch (((v92 + (("foo bar" + "a") + v92)) + v92) + (v92 + v92)) {
                                                                    case "\n\t":
                                                                        switch ((((v92 + v92) + ("" + v92)) + v92) + ("hello" + "hello")) {
                                                                        case "\n\t":
                                                                            v117 := GenS2{F0: (("" + (v92 + v92)) + v92), F1: ((v92 + (v92 + "")) + "")}
                                                                            _ = v117
                                                                        default:
                                                                            v92 = (v92 + "")
                                                                        }
                                                                    case "a":
                                                                        fmt.Println(v92, fi101, v92)
                                                                    }
                                                                    if (779 != 2147483647) {
                                                                        v118 := GenS8{F0: 29}
                                                                        v92 = v92
                                                                        var v119 int32 = 179
                                                                        _ = v118
                                                                        _ = v119
                                                                    } else {
                                                                        v92 = v92
                                                                    }
                                                                }
                                                            }
                                                            v120 := GenS2{F0: "foo bar", F1: v92}
                                                            if (714 == 712) {
                                                                var v121 int8 = 1
                                                                v92 = v92
                                                                _ = v121
                                                            } else {
                                                                switch 32768 {
                                                                case -891:
                                                                    switch v92 {
                                                                    case "hello":
                                                                        fmt.Println(fi101, v92, fi101)
                                                                    case "\n\t":
                                                                        v92 = ((v92 + v92) + v92)
                                                                    case "foo bar":
                                                                        if true {
                                                                            v92 = "foo bar"
                                                                        } else {
                                                                            switch v92 {
                                                                            case "hello":
                                                                                fmt.Println(v92, fi101)
                                                                            case "\n\t":
                                                                                v92 = (v92 + v92)
                                                                            default:
                                                                                if (-1 < -2147483648) {
                                                                                    if len(v12) > 0 {
                                                                                        v12[0] = fi101
                                                                                    }
                                                                                    fmt.Println(fi100)
                                                                                }
                                                                            }
                                                                            if len(v12) >= 2 {
                                                                                v12 = v12[1:len(v12)]
                                                                            }
                                                                        }
                                                                    default:
                                                                        var v122 string = v92
                                                                        _ = v122
                                                                    }
                                                                default:
                                                                    v92 = v92
                                                                }
                                                                if false {
                                                                    v92 = v92
                                                                    if (false || false) {
                                                                        fmt.Println(fi101, v92)
                                                                    } else {
                                                                        if (-2147483648 <= 42) {
                                                                            v123 := GenS4{F0: 1}
                                                                            _ = v123
                                                                        } else {
                                                                            var v124 bool = false
                                                                            var v125 float64 = 1e10
                                                                            var v126 string = v92
                                                                            var v127 int64 = -1
                                                                            _ = v124
                                                                            _ = v125
                                                                            _ = v126
                                                                            _ = v127
                                                                        }
                                                                    }
                                                                    var v128 float64 = 1.5
                                                                    _ = v128
                                                                }
                                                            }
                                                            _ = v120
                                                        }
                                                        _ = fi101
                                                    }
                                                    var v129 bool = (false && false)
                                                    var v130 float64 = (1e10 * -2.25)
                                                    _ = fi100
                                                    _ = v129
                                                    _ = v130
                                                }
                                                if false {
                                                    var v131 int8 = -43
                                                    v12 = append(v12, 938)
                                                    _ = v131
                                                }
                                            }
                                        }
                                        if false {
                                            for fi132 := 0; fi132 < 13; fi132++ {
                                                switch ("" + ("a" + "")) {
                                                case "hello":
                                                    var v133 int16 = -322
                                                    _ = v133
                                                case "":
                                                    if false {
                                                        for fi134 := 0; fi134 < 17; fi134++ {
                                                            if len(v12) > 0 {
                                                                v12[0] = fi134
                                                            }
                                                            _ = fi134
                                                        }
                                                        var v135 uint32 = 1
                                                        _ = v135
                                                    } else {
                                                        if (333 <= 770) {
                                                            if false {
                                                                if (!true) {
                                                                    var v136 uint8 = 105
                                                                    fmt.Println(v92, v92, v92)
                                                                    _ = v136
                                                                }
                                                                fmt.Println(v92)
                                                                for fi137 := 0; fi137 < 18; fi137++ {
                                                                    if ((!(!(2 != 643))) || true) {
                                                                        v92 = v92
                                                                        _ = len(v12)
                                                                        var v138 uint16 = 777
                                                                        _ = v138
                                                                    }
                                                                    _ = fi137
                                                                }
                                                            }
                                                            for fi139 := 0; fi139 < 3; fi139++ {
                                                                v140 := GenS8{F0: 0}
                                                                var v141 bool = true
                                                                var v142 uint8 = 2
                                                                _ = fi139
                                                                _ = v140
                                                                _ = v141
                                                                _ = v142
                                                            }
                                                        } else {
                                                            v143 := GenS9{F0: 2}
                                                            _ = v143
                                                        }
                                                        for fi144 := 0; fi144 < 3; fi144++ {
                                                            switch "" {
                                                            case "hello":
                                                                var v145 uint32 = 491
                                                                _ = v145
                                                            case "\n\t":
                                                                v12 = append(v12, fi132)
                                                            default:
                                                                if len(v12) > 0 {
                                                                    v12[0] = 4294967295
                                                                }
                                                            }
                                                            var v146 int16 = 1000
                                                            _ = fi144
                                                            _ = v146
                                                        }
                                                    }
                                                case "\n\t":
                                                    var v147 int32 = 1
                                                    _ = v147
                                                }
                                                v148 := GenS9{F0: 4294967295}
                                                _ = fi132
                                                _ = v148
                                            }
                                            v92 = "\n\t"
                                            switch 839 {
                                            case 1:
                                                var v149 bool = (true && (0 >= 1))
                                                _ = v149
                                            case 0:
                                                var v150 bool = (1 >= 1)
                                                _ = v150
                                            default:
                                                if true {
                                                    _ = len(v12)
                                                    for fi151 := 0; fi151 < 2; fi151++ {
                                                        var v152 uint32 = 11
                                                        fmt.Println(v92)
                                                        v152 = (v152 >> ((42) & 63))
                                                        _ = fi151
                                                        _ = v152
                                                    }
                                                }
                                            }
                                        } else {
                                            if false {
                                                var v153 uint32 = 263
                                                v153 = v153
                                                if false {
                                                    var v154 int64 = -1
                                                    _ = v154
                                                }
                                                _ = v153
                                            } else {
                                                if (!false) {
                                                    if (!(!(4294967295 == 65535))) {
                                                        v155 := GenS7{F0: -1, F1: -2147483648, F2: 65535}
                                                        switch 42 {
                                                        case 566:
                                                            v92 = (v92 + (v92 + (v92 + v92)))
                                                        case -804:
                                                            if (2147483647 >= 2147483647) {
                                                                if (v92 >= ((v92 + "\n\t") + "a")) {
                                                                    _ = len(v12)
                                                                }
                                                                v92 = v92
                                                            }
                                                        }
                                                        _ = v155
                                                    }
                                                    if true {
                                                        if len(v12) > 0 {
                                                            v12[0] = 1
                                                        }
                                                        v92 = ("" + v92)
                                                    }
                                                }
                                            }
                                            v92 = (v92 + (v92 + v92))
                                            switch v92 {
                                            case "\n\t":
                                                v92 = ((v92 + v92) + v92)
                                            }
                                        }
                                    }
                                case "\n\t":
                                    fmt.Println(v92)
                                default:
                                    switch 9223372036854775807 {
                                    case 42:
                                        v156 := GenS2{F0: v92, F1: ""}
                                        _ = v156
                                    }
                                }
                                _ = v92
                            }
                            var v157 uint64 = 0
                            _ = v157
                        } else {
                            if true {
                                if len(v12) > 0 {
                                    v12[0] = -666
                                }
                                var v158 int16 = -32768
                                _ = v158
                            }
                        }
                    case 32768:
                        v159 := GenS10{F0: false, F1: true}
                        _ = v159
                    case 9223372036854775807:
                        _ = len(v12)
                    }
                } else {
                    var v160 string = "foo bar"
                    if false {
                        switch v160 {
                        case "\n\t":
                            v161 := GenS6{F0: v160}
                            _ = v161
                        }
                        var v162 float32 = 0.1
                        _ = v162
                    } else {
                        if true {
                            v12 = append(v12, 42)
                            v160 = v160
                            v160 = "\n\t"
                        }
                        var v163 uint16 = 2
                        _ = v163
                    }
                    v164 := GenS11{F0: (!true), F1: v160}
                    _ = v160
                    _ = v164
                }
            }
        }
        var v165 int32 = -1
        fmt.Println(v165)
        _ = v165
    }
    if (-75 >= 1000) {
        for fi166 := 0; fi166 < 7; fi166++ {
            var v167 uint64 = 1
            switch ("a" + "foo bar") {
            case "foo bar":
                if len(v12) > 0 {
                    v12[0] = fi166
                }
            case "\n\t":
                for fi168 := 0; fi168 < 12; fi168++ {
                    switch (fi166 / (((fi168 % (((fi168 % ((fi166) | 1))) | 1))) | 1)) {
                    case -1:
                        v169 := GenS3{F0: 683, F1: (v167 + v167)}
                        _ = v169
                    case -843:
                        if (438 < 0) {
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            fmt.Println(v167, v167)
                        }
                    }
                    var v170 int16 = -32768
                    _ = fi168
                    _ = v170
                }
            }
            _ = fi166
            _ = v167
        }
        v171 := GenS6{F0: "foo bar"}
        var v172 string = "\n\t"
        _ = v171
        _ = v172
    } else {
        v173 := GenS7{F0: -349, F1: -2147483648, F2: 42}
        v174 := GenS12{F0: 0}
        _ = v173
        _ = v174
    }
    _ = v12
    return "\n\t"
}

func main() {
    v1 := []uint32{2147483648, 65535}
    var v2 uint32 = 2
    var v3 int64 = 1
    var v4 int64 = v3
    var v175 int16 = 1
    var v176 int8 = -128
    fmt.Println(v3, v4, v176)
    if len(v1) > 0 {
        v1[0] = v2
    }
    v177 := GenS13{F0: (v3 % (((v3 - (v4 - -9223372036854775808))) | 1))}
    if true {
        for fi178 := 0; fi178 < 19; fi178++ {
            var v179 uint64 = 18446744073709551615
            v2 = v2
            v180 := GenS9{F0: 1}
            _ = fi178
            _ = v179
            _ = v180
        }
        var v181 float64 = (-2.25 / ((0.1 - 1e10) / 0.1))
        v182 := GenS14{F0: true, F1: "foo bar", F2: genFn0(-2147483648, 2)}
        _ = v181
        _ = v182
    } else {
        var v183 string = "\n\t"
        _ = v183
    }
    var v184 uint32 = 932
    for fi185 := 0; fi185 < 15; fi185++ {
        if true {
            var v186 bool = (-731 > v3)
            switch 4294967295 {
            case 32767:
                v184 += v2
            case -1:
                _ = len(v1)
            default:
                for fi187 := 0; fi187 < 11; fi187++ {
                    var v188 int = 0
                    switch "foo bar" {
                    case "":
                        v189 := GenS10{F0: (2 <= 65535), F1: v186}
                        _ = v189
                    case "foo bar":
                        var v190 float64 = (3.14159 - ((1.5 + 0.0) - 3.14159))
                        _ = v190
                    case "hello":
                        switch ("a" + ("hello" + "hello")) {
                        case "foo bar":
                            v1 = append(v1, v184)
                        case "\n\t":
                            if (((0.0 + 3.14159) * 3.14159) < (float32(float64(-2.25)) * (3.14159 - -2.25))) {
                                switch (^(fi187)) {
                                case -454:
                                    v191 := GenS11{F0: false, F1: "\n\t"}
                                    _ = v191
                                default:
                                    if len(v1) > 0 {
                                        v1[0] = v2
                                    }
                                }
                                v3 = 9223372036854775807
                            }
                        }
                        if (126 > 0) {
                            if (!v186) {
                                v3 = v3
                                if len(v1) > 0 {
                                    v1[0] = (v2 & (v184 / (((v184 ^ v2)) | 1)))
                                }
                                var v192 int16 = (v175 - (v175 >> ((93) & 63)))
                                if v186 {
                                    v193 := GenS11{F0: (-2.25 > 0.1), F1: ("foo bar" + ("foo bar" + "foo bar"))}
                                    fmt.Println(v186, v4, v186)
                                    _ = v193
                                }
                                var v194 float64 = 0.0
                                _ = v192
                                _ = v194
                            }
                        } else {
                            v186 = (v4 <= v4)
                            v195 := GenS14{F0: v186, F1: "\n\t", F2: (v176 * 14)}
                            var v196 uint64 = 2
                            _ = v195
                            _ = v196
                        }
                    }
                    _ = fi187
                    _ = v188
                }
            }
            v197 := GenS15{F0: v175}
            _ = v186
            _ = v197
        } else {
            var v198 uint = 18446744073709551615
            _ = v198
        }
        _ = fi185
    }
    if ((false || true) || false) {
        var v199 int16 = v175
        _ = v199
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v175:", v175)
    fmt.Println("v176:", v176)
    fmt.Println("v184:", v184)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v175
    _ = v176
    _ = v177
    _ = v184
    _ = v2
    _ = v3
    _ = v4
}
