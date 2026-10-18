package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 bool; F2 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int32; F1 string; F2 uint8 }
type GenS9 struct { F0 string; F1 string; F2 int16 }

func main() {
    var v1 bool = (!false)
    var v2 bool = v1
    var v3 bool = ("\n\t" != ("" + "\n\t"))
    for fi4 := 0; fi4 < 8; fi4++ {
        switch (fi4 | fi4) {
        case -2147483648:
            switch "" {
            case "a":
                var v5 int32 = 206
                _ = v5
            case "hello":
                fmt.Println(v2, v3, v1)
            case "":
                fmt.Println(v1, v2)
            }
        case -450:
            v1 = (1 >= -9223372036854775808)
        case 42:
            v6 := GenS1{F0: 1}
            _ = v6
        default:
            v7 := []int{4294967295}
            _ = v7
        }
        switch fi4 {
        case -9223372036854775808:
            fmt.Println(v2, v1, v3)
        default:
            for fi8 := 0; fi8 < 19; fi8++ {
                v1 = (!(!(227 != 67)))
                if ((19 > 13) || v3) {
                    v9 := GenS1{F0: 42}
                    _ = v9
                } else {
                    var v10 int8 = 42
                    _ = v10
                }
                v11 := []int{fi8, (fi4 ^ (fi8 << ((36) & 63)))}
                _ = fi8
                _ = v11
            }
        }
        switch (fi4 & fi4) {
        case 127:
            switch (^(fi4)) {
            case 128:
                v12 := []int{}
                _ = v12
            }
        case 42:
            v13 := []int{fi4, fi4}
            _ = v13
        default:
            v2 = (!v2)
        }
        _ = fi4
    }
    v14 := GenS1{F0: 255}
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS1{F0: 0}
        _ = fi15
        _ = v16
    }
    switch "hello" {
    case "a":
        for fi17 := 0; fi17 < 10; fi17++ {
            if false {
                switch "\n\t" {
                case "":
                    v18 := GenS1{F0: 0}
                    _ = v18
                case "foo bar":
                    v19 := GenS1{F0: 139}
                    _ = v19
                case "\n\t":
                    switch "" {
                    case "a":
                        fmt.Println(v2, v3, v2)
                    case "foo bar":
                        v2 = v3
                    case "":
                        v20 := GenS1{F0: 2}
                        _ = v20
                    }
                }
                switch "a" {
                case "\n\t":
                    if (-1 > -1) {
                        var v21 string = ("\n\t" + "foo bar")
                        _ = v21
                    } else {
                        fmt.Println(v3, v2, v2)
                        v22 := []int{fi17, (fi17 - (fi17 >> ((255) & 63))), 273}
                        for fi23 := 0; fi23 < 20; fi23++ {
                            var v24 int8 = -1
                            var v25 uint16 = 887
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        _ = v22
                    }
                }
            }
            v26 := GenS1{F0: 0}
            _ = fi17
            _ = v26
        }
    case "hello":
        v27 := []int{1, 32767, 425, 558}
        _ = v27
    case "":
        v28 := GenS1{F0: 23}
        _ = v28
    default:
        v29 := []int{42, -2147483648}
        _ = v29
    }
    v3 = (255 != 255)
    v30 := GenS2{F0: 168, F1: 100}
    switch "a" {
    case "a":
        fmt.Println(v2, v2, v2)
    }
    fmt.Println(v3)
    for fi31 := 0; fi31 < 3; fi31++ {
        fmt.Println(v3, v1)
        for fi32 := 0; fi32 < 14; fi32++ {
            var v33 int16 = 42
            if v2 {
                v1 = (fi31 >= fi31)
                var v34 int = (fi31 | -1)
                switch (v34 | 42) {
                case 32768:
                    var v35 uint = 1
                    _ = v35
                case -599:
                    v36 := []int{}
                    _ = v36
                case 4294967295:
                    var v37 string = (("\n\t" + "foo bar") + "\n\t")
                    _ = v37
                default:
                    if (("a" + (("a" + "foo bar") + ((("\n\t" + ((("" + "foo bar") + (("\n\t" + ("hello" + "hello")) + "a")) + (("\n\t" + "\n\t") + "\n\t"))) + ("a" + "")) + ""))) == "hello") {
                        var v38 uint32 = 65535
                        _ = v38
                    } else {
                        var v39 int8 = -128
                        _ = v39
                    }
                }
                _ = v34
            }
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        v1 = (("" + ((("\n\t" + (("foo bar" + "hello") + ("" + ("a" + "a")))) + "hello") + "")) <= ("hello" + "a"))
        if v3 {
            switch ("a" + "\n\t") {
            case "hello":
                for fi41 := 0; fi41 < 10; fi41++ {
                    var v42 bool = true
                    v43 := GenS2{F0: 255, F1: 0}
                    var v44 uint16 = 2
                    if true {
                        if false {
                            if (100 > 0) {
                                v2 = (("foo bar" + "foo bar") <= "hello")
                                v1 = (!(!v1))
                                if (float32(((-2.25 + (1e10 / float64(0.1))) * 1.5)) < (1e10 + ((0.1 * ((float32(float64(0.1)) * (-2.25 / ((1e10 * 0.0) + 3.14159))) + 0.1)) * 3.14159))) {
                                    fmt.Println(v1, fi40, v44)
                                    v45 := GenS2{F0: 255, F1: 1}
                                    var v46 bool = v3
                                    fmt.Println(v42, v44)
                                    if v46 {
                                        v2 = (!v2)
                                        v46 = (!false)
                                        var v47 string = "a"
                                        _ = v47
                                    }
                                    var v48 uint = 564
                                    v42 = true
                                    v49 := []int{}
                                    _ = v45
                                    _ = v46
                                    _ = v48
                                    _ = v49
                                }
                            } else {
                                v50 := []int{fi40}
                                var v51 uint64 = 2
                                _ = v50
                                _ = v51
                            }
                        }
                        v2 = v2
                    }
                    _ = fi41
                    _ = v42
                    _ = v43
                    _ = v44
                }
            case "a":
                for fi52 := 0; fi52 < 7; fi52++ {
                    if true {
                        v3 = false
                        v53 := []int{fi40, fi52, 9223372036854775807, (fi40 * fi52)}
                        if (true && true) {
                            fmt.Println(v2, v3)
                            v2 = (!(-2.25 <= -2.25))
                            switch ((((((("\n\t" + (("foo bar" + "\n\t") + "hello")) + "hello") + "\n\t") + "hello") + "foo bar") + ("\n\t" + "")) + "hello") {
                            case "":
                                if v2 {
                                    var v54 int8 = 1
                                    if false {
                                        switch -2147483648 {
                                        case -839:
                                            v1 = v3
                                        case 32768:
                                            v53 = append(v53, fi52)
                                        case 4294967295:
                                            var v55 uint32 = 42
                                            _ = v55
                                        }
                                        var v56 int = (fi52 & 127)
                                        _ = v56
                                    } else {
                                        switch "" {
                                        case "hello":
                                            v54 = (v54 & (v54 << ((821) & 63)))
                                        }
                                        var v57 bool = (!true)
                                        _ = v57
                                    }
                                    if (-2147483648 != -109) {
                                        v1 = (v3 && (v2 && (true || (false || ("a" < "a")))))
                                        if (v3 || (v54 == 1)) {
                                            var v58 int32 = 2147483647
                                            v3 = (18446744073709551615 == 18446744073709551615)
                                            var v59 uint16 = 2
                                            _ = v58
                                            _ = v59
                                        } else {
                                            v60 := GenS3{F0: 2}
                                            _ = v60
                                        }
                                    } else {
                                        if (((0.1 + (1e10 / 1.5)) + 0.1) == 1e10) {
                                            var v61 uint16 = 565
                                            v53 = append(v53, fi52)
                                            if false {
                                                if len(v53) >= 2 {
                                                    v53 = v53[1:len(v53)]
                                                }
                                                if v1 {
                                                    v3 = (v61 <= (v61 >> ((65535) & 63)))
                                                    switch "a" {
                                                    case "a":
                                                        v3 = (((((0.1 - 1e10) * float64((float32((-2.25 + (1.5 + 0.0))) + (0.1 + (float32(1.5) - 1.5))))) >= 1.5) && true) || (2 >= 42))
                                                    default:
                                                        var v62 int8 = (v54 ^ v54)
                                                        v2 = v3
                                                        v61 = v61
                                                        _ = v62
                                                    }
                                                    v63 := GenS4{F0: 0, F1: 2}
                                                    _ = v63
                                                } else {
                                                    if ((0.1 * (float64(3.14159) * float64(((1e10 + (3.14159 + (3.14159 / (float32(((0.1 / 1.5) / 1.0)) + 1e10)))) / (0.1 / 1.5))))) == 1.5) {
                                                        v1 = true
                                                        v64 := GenS4{F0: 255, F1: 1}
                                                        switch (("foo bar" + (("foo bar" + "") + ("foo bar" + ("" + "a")))) + ("hello" + "\n\t")) {
                                                        case "":
                                                            var v65 uint = 2
                                                            if v1 {
                                                                v54 -= v54
                                                                switch (-(fi40)) {
                                                                case 2147483647:
                                                                    var v66 uint16 = (v61 << ((0) & 63))
                                                                    _ = v66
                                                                }
                                                            }
                                                            switch fi40 {
                                                            case -358:
                                                                var v67 float64 = ((3.14159 * (0.1 * 1e10)) - (1e10 * (-2.25 * (float64(((1.5 / ((float32(0.1) - 1e10) + (1e10 * 1e10))) - 1e10)) + 0.1))))
                                                                _ = v67
                                                            case 0:
                                                                if v1 {
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        v68 := GenS5{F0: -979}
                                                                        _ = v68
                                                                    }
                                                                } else {
                                                                    if len(v53) >= 2 {
                                                                        v53 = v53[1:len(v53)]
                                                                    }
                                                                }
                                                            default:
                                                                fmt.Println(v61, fi52, fi40)
                                                            }
                                                            v69 := GenS3{F0: (^(v61))}
                                                            switch fi40 {
                                                            case -2147483648:
                                                                var v70 int32 = -1
                                                                _ = v70
                                                            case 445:
                                                                v71 := GenS3{F0: 1}
                                                                _ = v71
                                                            default:
                                                                var v72 int32 = 0
                                                                _ = v72
                                                            }
                                                            _ = v65
                                                            _ = v69
                                                        case "\n\t":
                                                            v54 = 42
                                                        default:
                                                            if (255 >= 81) {
                                                                v3 = v1
                                                                v73 := GenS2{F0: 128, F1: -1}
                                                                var v74 int64 = 9223372036854775807
                                                                _ = v73
                                                                _ = v74
                                                            }
                                                        }
                                                        var v75 uint16 = v61
                                                        var v76 uint8 = uint8(v75)
                                                        _ = v64
                                                        _ = v75
                                                        _ = v76
                                                    } else {
                                                        switch (fi40 | (-(fi52))) {
                                                        case -293:
                                                            v77 := GenS6{F0: 65535, F1: true, F2: 0}
                                                            _ = v77
                                                        case 455:
                                                            var v78 int8 = -1
                                                            _ = v78
                                                        }
                                                        _ = len(v53)
                                                    }
                                                }
                                                v54 -= 100
                                            }
                                            _ = v61
                                        }
                                        switch fi52 {
                                        case -716:
                                            v79 := GenS5{F0: 1}
                                            _ = v79
                                        default:
                                            v80 := GenS7{F0: fi52}
                                            _ = v80
                                        }
                                    }
                                    if len(v53) >= 2 {
                                        v53 = v53[1:len(v53)]
                                    }
                                    v54 = (-(v54))
                                    _ = v54
                                }
                            case "\n\t":
                                v81 := GenS5{F0: 32767}
                                _ = v81
                            case "hello":
                                fmt.Println(v3)
                            default:
                                if true {
                                    var v82 int8 = 1
                                    if (0 == 896) {
                                        v83 := GenS1{F0: 42}
                                        if len(v53) >= 2 {
                                            v53 = v53[1:len(v53)]
                                        }
                                        v53 = append(v53, (fi52 ^ 32768))
                                        _ = v83
                                    }
                                    if len(v53) > 0 {
                                        v53[0] = 0
                                    }
                                    if true {
                                        var v84 uint64 = 2
                                        v85 := GenS1{F0: 112}
                                        v53 = append(v53, (^(fi52)))
                                        _ = v84
                                        _ = v85
                                    }
                                    _ = v82
                                } else {
                                    switch ("a" + "a") {
                                    case "hello":
                                        v3 = (((!v2) && v2) || (false && (100 >= -128)))
                                    case "\n\t":
                                        v86 := GenS1{F0: 180}
                                        _ = v86
                                    default:
                                        var v87 int16 = 242
                                        _ = v87
                                    }
                                }
                            }
                            v88 := GenS7{F0: (fi52 / (((^(fi40))) | 1))}
                            v2 = v2
                            if (9223372036854775807 >= 9223372036854775807) {
                                v3 = (42 < 127)
                                switch fi52 {
                                case 42:
                                    switch "foo bar" {
                                    case "hello":
                                        var v89 int8 = 42
                                        _ = v89
                                    }
                                default:
                                    var v90 float64 = (0.1 - (1.5 * 3.14159))
                                    _ = v90
                                }
                                var v91 string = ("\n\t" + ("foo bar" + "a"))
                                v1 = (v2 || v1)
                                var v92 uint8 = 42
                                _ = v91
                                _ = v92
                            } else {
                                v1 = v1
                                v53 = append(v53, -661)
                                v2 = (!v2)
                            }
                            if len(v53) >= 2 {
                                v53 = v53[1:len(v53)]
                            }
                            _ = v88
                        } else {
                            switch fi40 {
                            case 4294967295:
                                var v93 int32 = 950
                                _ = v93
                            }
                            v1 = false
                        }
                        _ = v53
                    }
                    switch fi40 {
                    case 32767:
                        var v94 int32 = 2147483647
                        _ = v94
                    case 32768:
                        fmt.Println(fi52, fi52, v1)
                    default:
                        var v95 int32 = 1
                        _ = v95
                    }
                    v96 := []int{4294967295, fi52, (-(fi52)), fi40}
                    v2 = (1 > 859)
                    var v97 bool = v3
                    _ = fi52
                    _ = v96
                    _ = v97
                }
            default:
                switch (fi40 | fi40) {
                case -1:
                    for fi98 := 0; fi98 < 5; fi98++ {
                        if (!v3) {
                            switch "a" {
                            case "":
                                var v99 float32 = (float32(-2.25) - 1e10)
                                _ = v99
                            case "hello":
                                var v100 uint64 = 18446744073709551615
                                _ = v100
                            case "a":
                                fmt.Println(v2)
                            }
                            v3 = v2
                            v3 = v3
                        } else {
                            switch 42 {
                            case 991:
                                var v101 uint32 = 2
                                _ = v101
                            case -2147483648:
                                switch 1 {
                                case 127:
                                    var v102 int64 = 42
                                    _ = v102
                                case 146:
                                    var v103 uint32 = 42
                                    _ = v103
                                case 174:
                                    v104 := []int{fi98, 128}
                                    _ = v104
                                default:
                                    fmt.Println(fi40, fi40)
                                }
                            case -317:
                                v105 := []int{}
                                _ = v105
                            default:
                                switch "foo bar" {
                                case "":
                                    fmt.Println(v1)
                                }
                            }
                            if v3 {
                                if (1 < 117) {
                                    var v106 uint16 = 1
                                    _ = v106
                                }
                                v3 = (!v1)
                            } else {
                                var v107 int16 = 0
                                var v108 bool = (18446744073709551615 >= 4294967295)
                                v109 := []int{}
                                if len(v109) >= 2 {
                                    v109 = v109[1:len(v109)]
                                }
                                _ = v107
                                _ = v108
                                _ = v109
                            }
                        }
                        _ = fi98
                    }
                case 9223372036854775807:
                    if true {
                        v110 := []int{877, -9223372036854775808, 32768}
                        switch "\n\t" {
                        case "a":
                            v1 = (v2 || (!(!v2)))
                        case "foo bar":
                            fmt.Println(fi40)
                        case "hello":
                            for fi111 := 0; fi111 < 18; fi111++ {
                                if v2 {
                                    if len(v110) > 0 {
                                        v110[0] = fi111
                                    }
                                } else {
                                    fmt.Println(fi40, fi40, fi111)
                                }
                                _ = fi111
                            }
                        }
                        _ = v110
                    }
                case 182:
                    v112 := GenS8{F0: -2147483648, F1: "\n\t", F2: 42}
                    _ = v112
                }
            }
            v113 := GenS1{F0: 255}
            _ = v113
        } else {
            v114 := GenS9{F0: "foo bar", F1: "\n\t", F2: 42}
            _ = v114
        }
        _ = fi40
    }
    var v115 uint = 172
    if (881 < 4294967295) {
        var v116 int64 = 641
        fmt.Println(v115, v2)
        var v117 uint16 = 347
        _ = v116
        _ = v117
    } else {
        for fi118 := 0; fi118 < 17; fi118++ {
            var v119 uint64 = 4294967295
            _ = fi118
            _ = v119
        }
    }
    fmt.Println(v1, v1)
    var v120 uint = 42
    for fi121 := 0; fi121 < 14; fi121++ {
        v120 = v115
        var v122 int = fi121
        switch "\n\t" {
        case "hello":
            v2 = false
        }
        _ = fi121
        _ = v122
    }
    if v1 {
        v123 := GenS4{F0: 3, F1: 4294967295}
        var v124 int8 = 1
        v124 = (v124 * v124)
        _ = v123
        _ = v124
    }
    fmt.Println("v1:", v1)
    fmt.Println("v115:", v115)
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v115
    _ = v120
    _ = v14
    _ = v2
    _ = v3
    _ = v30
}
