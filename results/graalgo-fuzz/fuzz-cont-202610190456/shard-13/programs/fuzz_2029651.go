package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 bool; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32; F1 uint32 }
type GenS9 struct { F0 int32; F1 int }
type GenS10 struct { F0 int64; F1 bool; F2 uint }
type GenS11 struct { F0 uint32; F1 int8; F2 string }
type GenS12 struct { F0 uint8; F1 string }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 bool; F1 uint64 }
func genFn0(p0 int, p1 int64, p2 int8) uint16 {
    fmt.Println(p1, p0, p1)
    fmt.Println(p2, p2, p0)
    p1 = 869
    p1 += (p1 & p1)
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    v1 := []int{32768, 32767, 32768}
    var v2 int8 = 100
    var v3 uint = 255
    var v4 int64 = 332
    v5 := make(map[string]int)
    var v6 uint = 0
    switch 32767 {
    case 520:
        switch 729 {
        case -568:
            for fi7 := 0; fi7 < 19; fi7++ {
                v4 -= -1
                _ = fi7
            }
        case 32767:
            v3 = v6
        case 970:
            if ((!false) || ((112 <= 128) && false)) {
                fmt.Println(v2)
                for fi8 := 0; fi8 < 20; fi8++ {
                    switch ("foo bar" + "a") {
                    case "\n\t":
                        var v9 int = -1
                        _ = v9
                    case "hello":
                        fmt.Println(fi8, fi8)
                    }
                    _ = fi8
                }
                for fi10 := 0; fi10 < 10; fi10++ {
                    for fi11 := 0; fi11 < 8; fi11++ {
                        var v12 int16 = 893
                        v13 := GenS1{F0: ((("hello" + "a") + (("a" + "\n\t") + "foo bar")) + "a")}
                        var v14 int8 = (v2 + v2)
                        if true {
                            delete(v5, ("foo bar" + (("a" + "a") + "a")))
                            switch fi11 {
                            case -1:
                                if true {
                                    v4 = v4
                                    switch "a" {
                                    case "a":
                                        v5["hello"] = 127
                                    default:
                                        switch ("" + "hello") {
                                        case "a":
                                            if ("" <= ("\n\t" + "a")) {
                                                v12 = int16(v2)
                                                v15 := v5["foo bar"]
                                                _ = v15
                                            }
                                        case "foo bar":
                                            v5["hello"] = (fi11 >> (((^(v6))) & 63))
                                        default:
                                            v16 := GenS1{F0: "\n\t"}
                                            var v17 float32 = ((0.0 * 1.5) * 0.1)
                                            v14 -= v14
                                            _ = v16
                                            _ = v17
                                        }
                                    }
                                    v3 -= (^(v3))
                                }
                            default:
                                if ((v4 % ((42) | 1)) != (v4 - v4)) {
                                    var v18 int16 = (v12 | (v12 ^ (v12 ^ (v12 + (v12 >> (((v3 & (v3 ^ (^(v6))))) & 63))))))
                                    var v19 uint64 = 88
                                    var v20 string = "a"
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                } else {
                                    v6 = (v6 / ((v3) | 1))
                                    v4 = (v4 >> (((^(v3))) & 63))
                                    fmt.Println(v12)
                                }
                                switch fi11 {
                                case -1:
                                    switch (fi11 % ((32767) | 1)) {
                                    case 128:
                                        if false {
                                            if ("a" == "a") {
                                                v6 += (v3 << (((v6 - v3)) & 63))
                                                _ = len(v5)
                                            } else {
                                                switch (fi10 >> ((v6) & 63)) {
                                                case 1:
                                                    v5[("a" + "hello")] = fi10
                                                }
                                                if true {
                                                    v2 = 0
                                                    v5["a"] = (fi10 & (fi11 + int(v12)))
                                                } else {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            }
                                            var v21 uint32 = 4294967295
                                            _ = v21
                                        }
                                    case -23:
                                        if len(v1) > 0 {
                                            v1[0] = 9223372036854775807
                                        }
                                    default:
                                        if (!(616 > 4294967295)) {
                                            if false {
                                                switch "" {
                                                case "hello":
                                                    v22 := GenS2{F0: (v2 & (-(v2))), F1: false, F2: 2}
                                                    _ = v22
                                                default:
                                                    _ = len(v1)
                                                }
                                            }
                                        }
                                    }
                                case 127:
                                    v6 += (v6 + (^(v3)))
                                case 629:
                                    if ("hello" < "foo bar") {
                                        var v23 int8 = (v14 - v14)
                                        var v24 uint64 = uint64(v3)
                                        _ = v23
                                        _ = v24
                                    }
                                }
                            }
                            fmt.Println(fi10, v14)
                        }
                        var v25 int32 = -388
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                        _ = v25
                    }
                    _ = fi10
                }
            } else {
                v6 -= 2
                for fi26 := 0; fi26 < 4; fi26++ {
                    var v27 uint8 = 1
                    var v28 int16 = -231
                    v29 := GenS3{F0: (v28 | 0)}
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
                if true {
                    var v30 int16 = 0
                    if false {
                        switch 127 {
                        case -507:
                            var v31 int = 32768
                            _ = v31
                        case 9223372036854775807:
                            v32 := GenS3{F0: v30}
                            _ = v32
                        }
                    } else {
                        if true {
                            var v33 uint = v3
                            delete(v5, "a")
                            _ = v33
                        } else {
                            v34 := GenS3{F0: int16(v2)}
                            var v35 bool = (1 >= v6)
                            _ = v34
                            _ = v35
                        }
                        for fi36 := 0; fi36 < 6; fi36++ {
                            if ((123 <= 127) || (4294967295 > 2147483648)) {
                                if ((9223372036854775807 > -1) || false) {
                                    var v37 string = ("foo bar" + "foo bar")
                                    fmt.Println(v3, v3)
                                    var v38 uint = v6
                                    _ = v37
                                    _ = v38
                                } else {
                                    for fi39 := 0; fi39 < 13; fi39++ {
                                        switch "\n\t" {
                                        case "":
                                            fmt.Println(v3, fi36)
                                        default:
                                            v1 = append(v1, (-(fi39)))
                                        }
                                        var v40 uint64 = 2
                                        var v41 uint32 = 2147483648
                                        _ = fi39
                                        _ = v40
                                        _ = v41
                                    }
                                    switch "" {
                                    case "":
                                        for fi42 := 0; fi42 < 17; fi42++ {
                                            v5[("" + "hello")] = 32767
                                            _ = fi42
                                        }
                                    case "hello":
                                        fmt.Println(v6, v3, fi36)
                                    default:
                                        var v43 float64 = 3.14159
                                        _ = v43
                                    }
                                }
                                if false {
                                    if (1 != 42) {
                                        for fi44 := 0; fi44 < 4; fi44++ {
                                            if (((("hello" + "\n\t") + ("" + (("" + "hello") + ("foo bar" + "hello")))) + "") >= "") {
                                                if true {
                                                    v5[""] = 4294967295
                                                }
                                            }
                                            _ = fi44
                                        }
                                    } else {
                                        var v45 float32 = 1.5
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v45
                                    }
                                    v46 := v5[""]
                                    v30 = (v30 & v30)
                                    _ = v46
                                } else {
                                    v5["\n\t"] = fi36
                                }
                            } else {
                                v4 += (v4 ^ (v4 >> ((v6) & 63)))
                                v5[(("a" + "a") + "hello")] = (fi36 ^ (^(fi36)))
                                var v47 uint32 = 65535
                                _ = v47
                            }
                            _ = fi36
                        }
                        for fi48 := 0; fi48 < 16; fi48++ {
                            v4 -= 655
                            _ = fi48
                        }
                    }
                    if false {
                        if ((v2 > (v2 + -1)) && ((1 >= -1) || false)) {
                            if (v4 == 42) {
                                for fi49 := 0; fi49 < 11; fi49++ {
                                    switch ("a" + (("\n\t" + "hello") + (("" + "hello") + "foo bar"))) {
                                    case "":
                                        var v50 uint16 = 32767
                                        _ = v50
                                    case "foo bar":
                                        v30 = v30
                                    case "a":
                                        var v51 bool = false
                                        _ = v51
                                    }
                                    v4 -= v4
                                    for fi52 := 0; fi52 < 7; fi52++ {
                                        v4 = (v4 / ((-1) | 1))
                                        _ = fi52
                                    }
                                    _ = fi49
                                }
                            }
                        }
                        var v53 uint = (v6 | 2)
                        for fi54 := 0; fi54 < 5; fi54++ {
                            switch (("a" + "hello") + "\n\t") {
                            case "a":
                                if ((1 != -2147483648) || false) {
                                    v6 -= v6
                                    for fi55 := 0; fi55 < 4; fi55++ {
                                        v56 := GenS2{F0: 0, F1: (392 > 0), F2: 42}
                                        var v57 uint16 = 0
                                        _ = fi55
                                        _ = v56
                                        _ = v57
                                    }
                                }
                            case "hello":
                                if true {
                                    if ((9223372036854775807 >= v4) && true) {
                                        v58 := GenS1{F0: "foo bar"}
                                        for fi59 := 0; fi59 < 15; fi59++ {
                                            switch -1 {
                                            case 9223372036854775807:
                                                delete(v5, "a")
                                                switch ("a" + (("\n\t" + "") + "")) {
                                                case "":
                                                    fmt.Println(v2)
                                                case "foo bar":
                                                    fmt.Println(v3, v53, fi54)
                                                case "a":
                                                    var v60 float32 = (0.1 + ((-2.25 - 3.14159) * (float32(3.14159) + -2.25)))
                                                    _ = v60
                                                }
                                            case 127:
                                                v4 += 384
                                            default:
                                                fmt.Println(v30)
                                            }
                                            _ = fi59
                                        }
                                        switch "a" {
                                        case "\n\t":
                                            v30 -= 376
                                        case "":
                                            if true {
                                                for fi61 := 0; fi61 < 14; fi61++ {
                                                    var v62 int16 = -1
                                                    _ = fi61
                                                    _ = v62
                                                }
                                            }
                                        default:
                                            v3 = v3
                                        }
                                        _ = v58
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = fi54
                                        }
                                    }
                                    var v63 uint = v3
                                    _ = v63
                                }
                            default:
                                v6 += v3
                            }
                            _ = fi54
                        }
                        _ = v53
                    }
                    _ = v30
                }
            }
        default:
            v6 = 4294967295
        }
    case 2147483647:
        delete(v5, "a")
    case 0:
        delete(v5, "foo bar")
    }
    v64 := GenS4{F0: 0, F1: v6}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2)
    if true {
        v65 := GenS2{F0: 1, F1: false, F2: 255}
        v66 := GenS2{F0: 127, F1: true, F2: 884}
        if false {
            if true {
                fmt.Println(v3)
            }
            for fi67 := 0; fi67 < 9; fi67++ {
                var v68 int32 = 1
                for fi69 := 0; fi69 < 18; fi69++ {
                    var v70 int16 = 237
                    _ = fi69
                    _ = v70
                }
                _ = fi67
                _ = v68
            }
        } else {
            v71 := v5[("foo bar" + ("foo bar" + ""))]
            switch (v71 / (((-(v71))) | 1)) {
            case 127:
                fmt.Println(v3)
            default:
                var v72 int = int(v6)
                _ = v72
            }
            switch ((("hello" + "\n\t") + ((("\n\t" + ("\n\t" + "")) + "hello") + "a")) + "\n\t") {
            case "\n\t":
                var v73 float32 = 3.14159
                _ = v73
            case "a":
                if (1e10 <= 1e10) {
                    for fi74 := 0; fi74 < 17; fi74++ {
                        v4 = -1
                        delete(v5, "a")
                        for fi75 := 0; fi75 < 17; fi75++ {
                            v3 = v6
                            _ = len(v5)
                            v76 := GenS1{F0: "\n\t"}
                            fmt.Println(v6)
                            switch ("hello" + "\n\t") {
                            case "a":
                                v77 := GenS5{F0: 127}
                                _ = v77
                            case "foo bar":
                                var v78 float32 = 0.1
                                _ = v78
                            case "\n\t":
                                fmt.Println(v3)
                            }
                            _ = fi75
                            _ = v76
                        }
                        _ = fi74
                    }
                    v3 = v3
                    if (2147483647 >= v71) {
                        if true {
                            switch (v71 / ((-676) | 1)) {
                            case 128:
                                v79 := GenS1{F0: ("" + "\n\t")}
                                _ = v79
                            default:
                                if false {
                                    v80 := GenS4{F0: 825, F1: 2}
                                    _ = v80
                                } else {
                                    v81 := GenS6{F0: 42}
                                    var v82 uint32 = 65535
                                    for fi83 := 0; fi83 < 10; fi83++ {
                                        if false {
                                            for fi84 := 0; fi84 < 16; fi84++ {
                                                v85 := GenS1{F0: (("" + ((("a" + "") + ("" + ("hello" + "a"))) + "")) + (((("foo bar" + (((("a" + "hello") + ("a" + "foo bar")) + "") + "a")) + "\n\t") + "foo bar") + "hello"))}
                                                v71 = 735
                                                _ = fi84
                                                _ = v85
                                            }
                                            if false {
                                                v86 := GenS7{F0: 0}
                                                var v87 uint16 = 42
                                                _ = v86
                                                _ = v87
                                            }
                                        }
                                        _ = fi83
                                    }
                                    _ = v81
                                    _ = v82
                                }
                            }
                            v88 := GenS8{F0: 798, F1: 65535}
                            _ = v88
                        } else {
                            var v89 uint = v6
                            _ = v89
                        }
                        var v90 uint = v3
                        for fi91 := 0; fi91 < 11; fi91++ {
                            v90 = (^(v90))
                            var v92 uint32 = 1
                            v4 += (v4 << ((v90) & 63))
                            _ = fi91
                            _ = v92
                        }
                        _ = v90
                    } else {
                        v2 -= (v2 ^ v2)
                    }
                }
            case "hello":
                for fi93 := 0; fi93 < 12; fi93++ {
                    for fi94 := 0; fi94 < 15; fi94++ {
                        var v95 int16 = -936
                        if true {
                            if len(v1) > 0 {
                                v1[0] = 32768
                            }
                            if (3.14159 != (float32(3.14159) / 1.5)) {
                                delete(v5, ("foo bar" + ""))
                                switch (fi94 << (((v3 - v3)) & 63)) {
                                case -321:
                                    delete(v5, (((("hello" + ("a" + "")) + (("foo bar" + "\n\t") + "\n\t")) + (("hello" + "") + "foo bar")) + "foo bar"))
                                default:
                                    fmt.Println(v71, v6)
                                }
                            } else {
                                if true {
                                    var v96 bool = (-2147483648 == 0)
                                    delete(v5, "\n\t")
                                    _ = v96
                                }
                                var v97 float64 = 1e10
                                _ = v97
                            }
                        } else {
                            v6 += (v6 % (((v3 % ((v3) | 1))) | 1))
                            if (-2147483648 < 740) {
                                if ((!false) && ((244 >= 35) && (((0 <= 1) || true) || true))) {
                                    v2 -= v2
                                    v98 := v5["a"]
                                    if true {
                                        var v99 int16 = 42
                                        var v100 uint64 = 42
                                        v3 = v3
                                        _ = v100
                                        _ = v99
                                    } else {
                                        v101 := GenS4{F0: 422, F1: 644}
                                        _ = v101
                                    }
                                    _ = v98
                                } else {
                                    v95 = (v95 % ((int16(v2)) | 1))
                                    switch fi94 {
                                    case -725:
                                        v6 -= (v3 + (v3 + 65535))
                                    case -1:
                                        v102 := GenS9{F0: 42, F1: 2147483647}
                                        _ = v102
                                    default:
                                        if (!true) {
                                            v71 = 635
                                            v4 -= 895
                                            switch "foo bar" {
                                            case "foo bar":
                                                switch fi94 {
                                                case -162:
                                                    v3 = (v6 << ((42) & 63))
                                                default:
                                                    if (!(42 < int64(v95))) {
                                                        switch (("" + "") + ("foo bar" + "\n\t")) {
                                                        case "":
                                                            v103 := v5[(("hello" + ("hello" + "")) + "")]
                                                            v5["hello"] = (v103 | 32767)
                                                            _ = v103
                                                        case "hello":
                                                            switch (("\n\t" + ("\n\t" + "a")) + ("a" + "hello")) {
                                                            case "":
                                                                if (333 > 4294967295) {
                                                                    v71 = fi93
                                                                    v5["hello"] = int(v2)
                                                                    if false {
                                                                        _ = len(v5)
                                                                        var v104 int8 = (v2 - 1)
                                                                        switch ("\n\t" + ("" + "hello")) {
                                                                        case "foo bar":
                                                                            var v105 int16 = 0
                                                                            _ = v105
                                                                        case "\n\t":
                                                                            v3 -= v6
                                                                        default:
                                                                            v1 = append(v1, fi94)
                                                                        }
                                                                        _ = v104
                                                                    }
                                                                } else {
                                                                    var v106 int16 = v95
                                                                    _ = v106
                                                                }
                                                                v1 = append(v1, (fi93 << ((v3) & 63)))
                                                                v107 := GenS10{F0: (v4 + v4), F1: false, F2: (v6 % (((v6 >> ((255) & 63))) | 1))}
                                                                _ = v107
                                                            case "hello":
                                                                switch 1 {
                                                                case 32767:
                                                                    delete(v5, "a")
                                                                case 1:
                                                                    var v108 bool = false
                                                                    _ = v108
                                                                case -1:
                                                                    if true {
                                                                        if (65535 == 2147483647) {
                                                                            v109 := GenS10{F0: -9223372036854775808, F1: (2 <= 42), F2: v6}
                                                                            _ = v109
                                                                        }
                                                                        v3 -= (v3 << (((^(v3))) & 63))
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                var v110 int16 = (v95 % ((v95) | 1))
                                                                _ = v110
                                                            default:
                                                                v111 := GenS3{F0: v95}
                                                                _ = v111
                                                            }
                                                        case "a":
                                                            v112 := GenS5{F0: 42}
                                                            v113 := GenS9{F0: 1, F1: -2147483648}
                                                            v3 += (^(v3))
                                                            _ = v112
                                                            _ = v113
                                                        default:
                                                            var v114 int8 = int8(v6)
                                                            switch ("\n\t" + "foo bar") {
                                                            case "\n\t":
                                                                if true {
                                                                    var v115 int = (^(v71))
                                                                    v114 = (v114 + v2)
                                                                    if false {
                                                                        var v116 uint64 = 42
                                                                        _ = v116
                                                                    } else {
                                                                        v117 := GenS11{F0: 1, F1: v2, F2: ""}
                                                                        _ = len(v1)
                                                                        _ = v117
                                                                    }
                                                                    _ = v115
                                                                }
                                                            default:
                                                                var v118 float32 = 3.14159
                                                                _ = v118
                                                            }
                                                            var v119 int = fi93
                                                            var v120 uint8 = 2
                                                            _ = v114
                                                            _ = v119
                                                            _ = v120
                                                        }
                                                        if true {
                                                            var v121 float32 = 0.0
                                                            _ = v121
                                                        }
                                                        var v122 string = "\n\t"
                                                        _ = v122
                                                    } else {
                                                        v123 := GenS11{F0: 233, F1: v2, F2: ("" + "")}
                                                        _ = v123
                                                    }
                                                    v124 := GenS1{F0: ("hello" + (("a" + "") + "\n\t"))}
                                                    _ = v124
                                                }
                                                v125 := GenS12{F0: 45, F1: ""}
                                                var v126 uint8 = 2
                                                _ = v125
                                                _ = v126
                                            }
                                        } else {
                                            v6 -= (v6 + (v3 * v3))
                                            v4 -= (v4 | 1)
                                            switch (fi94 % ((fi93) | 1)) {
                                            case 9223372036854775807:
                                                v4 = -818
                                            default:
                                                v127 := GenS13{F0: 2147483647}
                                                _ = v127
                                            }
                                        }
                                    }
                                    var v128 uint32 = 2
                                    _ = v128
                                }
                                v71 = -2147483648
                                v129 := GenS5{F0: int(v4)}
                                _ = v129
                            }
                            if (!true) {
                                var v130 int16 = 42
                                v6 += (v6 % ((v3) | 1))
                                _ = v130
                            }
                        }
                        _ = fi94
                        _ = v95
                    }
                    v2 -= (v2 / (((v2 / ((v2) | 1))) | 1))
                    _ = fi93
                }
            default:
                if true {
                    var v131 uint64 = 176
                    _ = v131
                }
            }
            _ = v71
        }
        _ = v65
        _ = v66
    }
    var v132 uint8 = 2
    var v133 int64 = -9223372036854775808
    v2 = 127
    var v134 bool = (((1e10 + ((((3.14159 * -2.25) + (((1.5 / (-2.25 + -2.25)) * 1e10) / ((1.5 + (1e10 * -2.25)) + -2.25))) - 0.0) / 1e10)) != (0.1 - 0.0)) && false)
    v132 -= (v132 | (v132 + v132))
    v135 := GenS14{F0: (0.0 <= 0.1), F1: 4294967295}
    if (!(!false)) {
        if v134 {
            if (!true) {
                for fi136 := 0; fi136 < 4; fi136++ {
                    if (float32(0.1) < float32(((1.5 * 1.5) - ((0.0 * 0.1) + (float64((3.14159 / 1.0)) / 1.0))))) {
                        for fi137 := 0; fi137 < 13; fi137++ {
                            if v134 {
                                v138 := GenS4{F0: 2, F1: v3}
                                _ = v138
                            } else {
                                fmt.Println(v4, fi137, fi136)
                                delete(v5, ("foo bar" + "hello"))
                            }
                            var v139 int = (fi137 % ((fi137) | 1))
                            _ = fi137
                            _ = v139
                        }
                    } else {
                        var v140 int16 = -1
                        v141 := GenS9{F0: -481, F1: 2147483647}
                        _ = v140
                        _ = v141
                    }
                    var v142 uint = (v6 & v3)
                    var v143 uint32 = 2147483648
                    _ = fi136
                    _ = v142
                    _ = v143
                }
                v144 := GenS5{F0: 9223372036854775807}
                _ = v144
            } else {
                var v145 int8 = 122
                _ = v145
            }
            fmt.Println(v4, v3)
        } else {
            v146 := GenS12{F0: v132, F1: "hello"}
            v134 = v134
            for fi147 := 0; fi147 < 8; fi147++ {
                v148 := GenS3{F0: 68}
                for fi149 := 0; fi149 < 17; fi149++ {
                    v134 = v134
                    _ = len(v1)
                    _ = fi149
                }
                switch 697 {
                case 745:
                    v132 = (v132 >> ((v3) & 63))
                case 127:
                    switch ("hello" + "\n\t") {
                    case "\n\t":
                        var v150 int64 = (v4 | (-(v133)))
                        _ = v150
                    case "":
                        v6 = (v6 & v3)
                    }
                }
                _ = fi147
                _ = v148
            }
            _ = v146
        }
        _ = len(v1)
        _ = len(v5)
    }
    var v151 int16 = -286
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v132:", v132)
    fmt.Println("v133:", v133)
    fmt.Println("v134:", v134)
    fmt.Println("v151:", v151)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v132
    _ = v133
    _ = v134
    _ = v135
    _ = v151
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
