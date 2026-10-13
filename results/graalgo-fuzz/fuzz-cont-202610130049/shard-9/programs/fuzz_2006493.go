package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int; F1 int8; F2 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 uint; F1 string; F2 uint32 }
type GenS7 struct { F0 string; F1 int; F2 int32 }
type GenS8 struct { F0 uint64; F1 string; F2 int }

func main() {
    var v1 string = ""
    var v2 bool = (!false)
    v3 := []int{-9223372036854775808}
    if true {
        if (42 < 4294967295) {
            if (((float64(1.5) * -2.25) - (1.5 - 0.1)) != 0.0) {
                v2 = v2
                switch "foo bar" {
                case "hello":
                    switch 711 {
                    case -629:
                        var v4 int8 = 100
                        _ = v4
                    }
                default:
                    for fi5 := 0; fi5 < 15; fi5++ {
                        v2 = true
                        if v2 {
                            fmt.Println(v1, v2)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = fi5
                    }
                }
                var v6 uint64 = 0
                _ = v6
            } else {
                var v7 uint64 = 236
                v1 = "foo bar"
                _ = v7
            }
            v2 = (42 > -9223372036854775808)
        } else {
            var v8 int8 = 0
            var v9 uint32 = 0
            if (42 > -9223372036854775808) {
                v2 = true
            } else {
                var v10 int64 = 180
                v11 := GenS1{F0: 90, F1: ""}
                var v12 float32 = -2.25
                _ = v10
                _ = v11
                _ = v12
            }
            _ = v8
            _ = v9
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := GenS2{F0: 65535, F1: 92, F2: 1}
        fmt.Println(fi13)
        _ = fi13
        _ = v14
    }
    if v2 {
        fmt.Println(v2, v1)
        var v15 uint32 = 53
        var v16 string = (v1 + v1)
        _ = v15
        _ = v16
    }
    v1 = v1
    if (!(32767 != 74)) {
        fmt.Println(v2, v1)
    }
    if (65535 > 32767) {
        v2 = (32767 > 255)
        var v17 uint64 = 2
        _ = v17
    } else {
        for fi18 := 0; fi18 < 1; fi18++ {
            if (-46 == 32767) {
                for fi19 := 0; fi19 < 7; fi19++ {
                    fmt.Println(v1)
                    v20 := GenS2{F0: 605, F1: 146, F2: 1}
                    _ = fi19
                    _ = v20
                }
                v2 = ((128 <= 203) || ((0.1 / 0.1) != 3.14159))
                switch "\n\t" {
                case "hello":
                    switch (v1 + v1) {
                    case "hello":
                        v2 = (42 > -2147483648)
                    case "\n\t":
                        if true {
                            for fi21 := 0; fi21 < 10; fi21++ {
                                if (v2 || v2) {
                                    fmt.Println(v1, v1)
                                    var v22 int16 = 1000
                                    var v23 bool = (true && ((true || v2) && (v2 && ("foo bar" <= "foo bar"))))
                                    var v24 uint = 746
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                                fmt.Println(v2, fi18)
                                fmt.Println(v2, fi21)
                                var v25 uint16 = 0
                                _ = fi21
                                _ = v25
                            }
                            v26 := GenS1{F0: 15, F1: ((v1 + "hello") + "foo bar")}
                            if (955 <= 2) {
                                fmt.Println(v2, v2)
                            } else {
                                for fi27 := 0; fi27 < 5; fi27++ {
                                    if false {
                                        v1 = v1
                                        switch fi27 {
                                        case 291:
                                            var v28 int8 = -1
                                            v1 = v1
                                            v1 = (v1 + v1)
                                            _ = v28
                                        case 4294967295:
                                            switch fi27 {
                                            case 127:
                                                v2 = (((65535 >= 65535) && (1 > 0)) || ("hello" < v1))
                                            case 1:
                                                v29 := GenS1{F0: 0, F1: v1}
                                                _ = v29
                                            case 42:
                                                var v30 string = v1
                                                _ = v30
                                            default:
                                                v31 := GenS3{F0: (^(fi18)), F1: 100, F2: 42}
                                                if v2 {
                                                    switch 127 {
                                                    case 4294967295:
                                                        var v32 int8 = 42
                                                        _ = v32
                                                    case 127:
                                                        switch v1 {
                                                        case "foo bar":
                                                            if len(v3) > 0 {
                                                                v3[0] = fi18
                                                            }
                                                        }
                                                        if (255 <= 255) {
                                                            var v33 string = ((v1 + v1) + ("foo bar" + (v1 + "\n\t")))
                                                            v1 = v1
                                                            _ = v33
                                                        } else {
                                                            v34 := GenS1{F0: 183, F1: "a"}
                                                            fmt.Println(v2, v2)
                                                            var v35 uint16 = 32767
                                                            v2 = v2
                                                            if (!(!(0.0 <= -2.25))) {
                                                                v2 = v2
                                                                v36 := GenS4{F0: v1}
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                _ = v36
                                                            }
                                                            var v37 uint64 = 1
                                                            v38 := GenS2{F0: 4294967295, F1: 2, F2: v35}
                                                            if (!(v2 || v2)) {
                                                                var v39 int32 = 0
                                                                if v2 {
                                                                    fmt.Println(fi18, v35, v1)
                                                                    if len(v3) > 0 {
                                                                        v3[0] = -975
                                                                    }
                                                                    v2 = (!v2)
                                                                    var v40 int16 = int16(v39)
                                                                    _ = v40
                                                                }
                                                                v39 = v39
                                                                _ = v39
                                                            } else {
                                                                var v41 int64 = 1
                                                                _ = v41
                                                            }
                                                            v35 -= 1
                                                            _ = v34
                                                            _ = v35
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                    }
                                                    v1 = "\n\t"
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                _ = v31
                                            }
                                        }
                                    }
                                    var v42 float32 = (0.1 - -2.25)
                                    _ = fi27
                                    _ = v42
                                }
                                v43 := GenS3{F0: (fi18 & fi18), F1: -128, F2: 848}
                                var v44 uint8 = 1
                                _ = v43
                                _ = v44
                            }
                            _ = v26
                        } else {
                            v45 := GenS5{F0: -32768, F1: 255}
                            for fi46 := 0; fi46 < 4; fi46++ {
                                fmt.Println(fi46)
                                var v47 int32 = -1
                                _ = fi46
                                _ = v47
                            }
                            for fi48 := 0; fi48 < 9; fi48++ {
                                var v49 uint32 = 1
                                fmt.Println(fi18)
                                var v50 uint8 = 42
                                switch (fi48 - fi18) {
                                case -785:
                                    v1 = ((v1 + "foo bar") + (v1 + ("" + "hello")))
                                default:
                                    fmt.Println(v50, v49)
                                }
                                _ = fi48
                                _ = v49
                                _ = v50
                            }
                            _ = v45
                        }
                    }
                case "":
                    var v51 uint = 18446744073709551615
                    _ = v51
                default:
                    var v52 uint16 = 255
                    _ = v52
                }
            } else {
                v1 = ("\n\t" + (v1 + ((v1 + v1) + v1)))
                switch (fi18 * (fi18 & -170)) {
                case 704:
                    if (-1 <= 0) {
                        switch (^(fi18)) {
                        case 9223372036854775807:
                            var v53 int8 = 100
                            _ = v53
                        case 1:
                            v54 := GenS5{F0: -426, F1: 1}
                            _ = v54
                        }
                        var v55 int64 = 1
                        switch (fi18 ^ (fi18 * 128)) {
                        case 680:
                            var v56 int8 = 1
                            _ = v56
                        case 968:
                            var v57 uint8 = 255
                            _ = v57
                        }
                        _ = v55
                    }
                case -523:
                    var v58 float32 = 3.14159
                    _ = v58
                default:
                    switch fi18 {
                    case 4294967295:
                        switch (fi18 & fi18) {
                        case -9223372036854775808:
                            v1 = ("\n\t" + v1)
                        case -564:
                            switch (fi18 - fi18) {
                            case -9223372036854775808:
                                v1 = v1
                            }
                        default:
                            var v59 int8 = int8(fi18)
                            _ = v59
                        }
                    default:
                        if ((true && (65535 == 679)) || v2) {
                            v2 = v2
                        } else {
                            var v60 int8 = -128
                            v2 = ((((!(!(!true))) || (!v2)) || (1 > 173)) && (9223372036854775807 != -451))
                            v1 = (v1 + v1)
                            _ = v60
                        }
                    }
                }
                var v61 int8 = 100
                _ = v61
            }
            if v2 {
                for fi62 := 0; fi62 < 17; fi62++ {
                    if v2 {
                        v63 := GenS6{F0: 4294967295, F1: ("" + v1), F2: 1}
                        var v64 uint32 = 1
                        switch (-(fi18)) {
                        case -1:
                            switch (v1 + v1) {
                            case "a":
                                v2 = ((v64 ^ v64) >= v64)
                            case "\n\t":
                                v2 = true
                            }
                        case -978:
                            if (!v2) {
                                v65 := GenS5{F0: -52, F1: 1}
                                _ = v65
                            } else {
                                v66 := GenS5{F0: 1000, F1: 0}
                                var v67 uint8 = 2
                                v68 := GenS2{F0: 18446744073709551615, F1: v67, F2: 42}
                                _ = v66
                                _ = v67
                                _ = v68
                            }
                        default:
                            var v69 int8 = 100
                            switch "a" {
                            case "hello":
                                var v70 int32 = int32(v69)
                                _ = v70
                            case "\n\t":
                                fmt.Println(fi18)
                                v2 = (!v2)
                            }
                            var v71 int32 = -2147483648
                            _ = v69
                            _ = v71
                        }
                        _ = v63
                        _ = v64
                    }
                    fmt.Println(fi18, v2)
                    _ = fi62
                }
                v1 = ""
            } else {
                fmt.Println(fi18, fi18)
                for fi72 := 0; fi72 < 13; fi72++ {
                    if (v2 || (23 == 100)) {
                        fmt.Println(fi18, v1, fi18)
                        v2 = (!false)
                        switch (v1 + v1) {
                        case "a":
                            v2 = (!(127 != 1))
                        default:
                            if false {
                                fmt.Println(fi72, v2)
                                if v2 {
                                    v73 := GenS6{F0: 0, F1: "foo bar", F2: 42}
                                    var v74 int64 = 182
                                    _ = v73
                                    _ = v74
                                } else {
                                    var v75 float32 = 3.14159
                                    _ = v75
                                }
                            }
                        }
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi72
                }
            }
            for fi76 := 0; fi76 < 17; fi76++ {
                fmt.Println(fi18)
                v3 = append(v3, (fi76 | fi18))
                var v77 uint64 = 1
                _ = fi76
                _ = v77
            }
            _ = fi18
        }
        v1 = v1
        v1 = v1
    }
    for fi78 := 0; fi78 < 13; fi78++ {
        if (v2 && false) {
            v2 = v2
        } else {
            if (!(1 > 168)) {
                fmt.Println(v2)
                switch (fi78 << ((255) & 63)) {
                case 620:
                    var v79 uint16 = 255
                    _ = v79
                case 42:
                    var v80 float32 = (float32(0.1) / 1e10)
                    _ = v80
                }
                switch fi78 {
                case -653:
                    var v81 string = (("a" + ((v1 + (v1 + v1)) + "")) + ("hello" + "hello"))
                    _ = v81
                case -2147483648:
                    v82 := GenS7{F0: v1, F1: (fi78 >> ((0) & 63)), F2: 42}
                    _ = v82
                case 375:
                    v2 = (v2 || v2)
                }
            }
            for fi83 := 0; fi83 < 4; fi83++ {
                v1 = v1
                var v84 uint16 = 1
                v1 = "foo bar"
                _ = fi83
                _ = v84
            }
        }
        var v85 int64 = 9223372036854775807
        v1 = "hello"
        _ = fi78
        _ = v85
    }
    for fi86 := 0; fi86 < 7; fi86++ {
        switch fi86 {
        case 0:
            for fi87 := 0; fi87 < 9; fi87++ {
                v1 = "\n\t"
                v2 = true
                var v88 int = fi87
                v89 := GenS6{F0: 1, F1: ((v1 + (v1 + ((("" + ("hello" + "foo bar")) + (v1 + ("" + "foo bar"))) + v1))) + v1), F2: 303}
                if (42 > 446) {
                    v88 = (v88 & -2147483648)
                    v90 := GenS5{F0: 0, F1: 1}
                    _ = v90
                }
                if (66 > 127) {
                    switch (v1 + "") {
                    case "foo bar":
                        if (!v2) {
                            var v91 uint32 = 65535
                            switch (v88 * (fi86 | fi86)) {
                            case 32768:
                                v91 = 0
                                var v92 int64 = 0
                                if (-1 <= 32767) {
                                    v91 = v91
                                } else {
                                    v93 := GenS3{F0: 127, F1: -1, F2: 1}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v94 uint64 = 2
                                    _ = v93
                                    _ = v94
                                }
                                _ = v92
                            }
                            v95 := GenS4{F0: "\n\t"}
                            _ = v91
                            _ = v95
                        }
                    case "a":
                        fmt.Println(fi86, v1, v88)
                    }
                    v2 = (-872 > 32767)
                } else {
                    if v2 {
                        var v96 uint8 = 255
                        _ = v96
                    }
                    _ = len(v3)
                    if len(v3) > 0 {
                        v3[0] = (fi87 & fi87)
                    }
                }
                _ = fi87
                _ = v88
                _ = v89
            }
        case 1:
            switch (fi86 % ((-26) | 1)) {
            case -9223372036854775808:
                if v2 {
                    if (1 <= 0) {
                        for fi97 := 0; fi97 < 4; fi97++ {
                            v1 = (v1 + v1)
                            v98 := GenS8{F0: 18446744073709551615, F1: "hello", F2: fi97}
                            v2 = (false || (1 > 728))
                            _ = len(v3)
                            var v99 float32 = (float32(1.5) + (1.5 - (0.0 - ((float32(0.0) / -2.25) + (3.14159 - -2.25)))))
                            fmt.Println(fi97, fi97)
                            _ = fi97
                            _ = v98
                            _ = v99
                        }
                    }
                }
            case -793:
                var v100 uint8 = 2
                _ = v100
            case 4294967295:
                if v2 {
                    switch fi86 {
                    case 9223372036854775807:
                        var v101 string = ("" + "\n\t")
                        _ = v101
                    default:
                        var v102 int16 = 934
                        _ = v102
                    }
                }
            }
        case 127:
            fmt.Println(fi86, v1)
        }
        v2 = ((3.14159 + 1.5) >= 0.1)
        v2 = (v2 || true)
        _ = fi86
    }
    switch v1 {
    case "\n\t":
        for fi103 := 0; fi103 < 7; fi103++ {
            switch ((v1 + v1) + "") {
            case "hello":
                var v104 uint64 = 205
                _ = v104
            case "\n\t":
                fmt.Println(v1, v2, v2)
            default:
                var v105 int = fi103
                _ = v105
            }
            _ = fi103
        }
    case "":
        switch 1 {
        case 9223372036854775807:
            for fi106 := 0; fi106 < 18; fi106++ {
                var v107 float32 = float32(-2.25)
                v2 = (722 <= 42)
                var v108 bool = true
                _ = fi106
                _ = v107
                _ = v108
            }
        case 42:
            if (!(3.14159 > float64(1e10))) {
                v2 = v2
            }
        case -1:
            for fi109 := 0; fi109 < 20; fi109++ {
                switch fi109 {
                case -800:
                    _ = len(v3)
                case 481:
                    v1 = v1
                }
                var v110 float32 = -2.25
                var v111 uint = 255
                _ = fi109
                _ = v110
                _ = v111
            }
        }
    default:
        _ = len(v3)
    }
    v1 = v1
    for fi112 := 0; fi112 < 8; fi112++ {
        v113 := GenS5{F0: 1000, F1: 1}
        var v114 int16 = 1000
        v1 = v1
        _ = fi112
        _ = v113
        _ = v114
    }
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    var v115 uint32 = 42
    var v116 int32 = -620
    fmt.Println("v1:", v1)
    fmt.Println("v115:", v115)
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v115
    _ = v116
    _ = v2
    _ = v3
}
