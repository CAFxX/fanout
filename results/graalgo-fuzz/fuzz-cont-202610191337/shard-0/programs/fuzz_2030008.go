package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 uint64; F1 int8; F2 uint }
type GenS6 struct { F0 uint16 }
func genFn0() int {
    var v5 float32 = -2.25
    if true {
        switch ("hello" + "foo bar") {
        case "hello":
            v6 := []int{}
            _ = v6
        case "a":
            switch -904 {
            case 417:
                v7 := []int{}
                _ = v7
            }
        default:
            for fi8 := 0; fi8 < 12; fi8++ {
                v9 := []int{4294967295, 32767, 0}
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                _ = fi8
                _ = v9
            }
        }
        var v10 int16 = -115
        _ = v10
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 float32 = v5
            v5 = (1e10 * 1.5)
            _ = fi11
            _ = v12
        }
        v13 := []int{-2147483648, 351, 794}
        switch "hello" {
        case "hello":
            switch "a" {
            case "hello":
                var v14 int32 = 268
                _ = v14
            }
        default:
            var v15 int = -532
            _ = v15
        }
        _ = v13
    }
    _ = v5
    return -9223372036854775808
}

func main() {
    var v1 bool = (!false)
    var v2 uint64 = 916
    v3 := []int{127}
    v4 := []string{"hello", "a", ("foo bar" + "foo bar"), "foo bar"}
    v16 := GenS1{F0: v1}
    if ((float64(1.5) > ((3.14159 / 1.0) - -2.25)) && v1) {
        var v17 int32 = -2147483648
        _ = v17
    }
    v1 = (621 > 255)
    switch (("\n\t" + "foo bar") + "hello") {
    case "\n\t":
        var v18 int32 = 42
        _ = v18
    case "":
        v19 := GenS1{F0: (132 >= 1)}
        _ = v19
    case "hello":
        v2 = v2
    default:
        v1 = v1
    }
    fmt.Println(v2)
    fmt.Println(v2, v1)
    fmt.Println(v1, v1)
    fmt.Println(v1)
    if true {
        v20 := GenS1{F0: v1}
        fmt.Println(v2)
        switch -9223372036854775808 {
        case 9223372036854775807:
            v21 := GenS2{F0: -128, F1: -1, F2: 0}
            _ = v21
        case 32767:
            v22 := GenS3{F0: 2147483647}
            _ = v22
        }
        _ = v20
    } else {
        switch 687 {
        case 9223372036854775807:
            if v1 {
                switch 32768 {
                case 42:
                    var v23 int32 = 0
                    _ = v23
                case 2147483647:
                    v1 = (!(!(58 != 1)))
                case 32767:
                    var v24 string = "a"
                    _ = v24
                default:
                    switch "a" {
                    case "hello":
                        switch genFn0() {
                        case -9223372036854775808:
                            var v25 float64 = 3.14159
                            _ = v25
                        case -1:
                            v2 -= (v2 * 18446744073709551615)
                        }
                    default:
                        var v26 string = "a"
                        _ = v26
                    }
                }
            } else {
                switch 4294967295 {
                case 789:
                    v1 = (3.14159 < ((float64(3.14159) * 1e10) + 1.5))
                case -692:
                    var v27 uint8 = 128
                    _ = v27
                case 31:
                    v1 = ((0 == -32768) && (false || true))
                default:
                    var v28 float64 = (1.5 + 1.5)
                    _ = v28
                }
            }
        default:
            if (!(false && false)) {
                for fi29 := 0; fi29 < 4; fi29++ {
                    for fi30 := 0; fi30 < 5; fi30++ {
                        var v31 uint = 990
                        var v32 string = ""
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    if (42 >= -1) {
                        v2 += 67
                        fmt.Println(v2, fi29, v2)
                        switch (("a" + "\n\t") + "a") {
                        case "":
                            var v33 bool = v1
                            _ = v33
                        case "foo bar":
                            switch "hello" {
                            case "a":
                                if true {
                                    switch 128 {
                                    case 791:
                                        v2 += 247
                                    }
                                }
                            case "\n\t":
                                var v34 int16 = int16(fi29)
                                _ = v34
                            case "hello":
                                var v35 uint = 65535
                                _ = v35
                            default:
                                if v1 {
                                    for fi36 := 0; fi36 < 5; fi36++ {
                                        var v37 float64 = float64(1.5)
                                        v1 = (79 == 127)
                                        v2 = 592
                                        _ = fi36
                                        _ = v37
                                    }
                                    switch fi29 {
                                    case -956:
                                        var v38 int8 = 127
                                        _ = v38
                                    case 1:
                                        fmt.Println(v2)
                                    default:
                                        switch (fi29 | fi29) {
                                        case -2147483648:
                                            for fi39 := 0; fi39 < 17; fi39++ {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                var v40 bool = (127 != 1)
                                                _ = fi39
                                                _ = v40
                                            }
                                        case -9223372036854775808:
                                            fmt.Println(fi29)
                                        }
                                    }
                                }
                            }
                        default:
                            for fi41 := 0; fi41 < 17; fi41++ {
                                v4 = append(v4, "foo bar")
                                fmt.Println(fi29, fi29)
                                _ = fi41
                            }
                        }
                    } else {
                        switch "hello" {
                        case "a":
                            for fi42 := 0; fi42 < 10; fi42++ {
                                v1 = true
                                var v43 bool = (0.0 <= 1.5)
                                _ = fi42
                                _ = v43
                            }
                        default:
                            var v44 int16 = -1
                            _ = v44
                        }
                        switch fi29 {
                        case 9223372036854775807:
                            v2 += (v2 * 300)
                        default:
                            if (1.5 != 3.14159) {
                                if ((fi29 / ((fi29) | 1)) >= 0) {
                                    _ = len(v3)
                                }
                            } else {
                                switch genFn0() {
                                case -191:
                                    if v1 {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        switch -961 {
                                        case 128:
                                            for fi45 := 0; fi45 < 3; fi45++ {
                                                v46 := GenS2{F0: 34, F1: -116, F2: 24}
                                                var v47 float32 = float32(1.5)
                                                _ = fi45
                                                _ = v46
                                                _ = v47
                                            }
                                        }
                                    } else {
                                        var v48 float64 = (3.14159 / 3.14159)
                                        var v49 uint8 = 133
                                        switch ("hello" + (("hello" + "foo bar") + ((("\n\t" + "foo bar") + ("\n\t" + (("hello" + ("a" + "foo bar")) + ("" + "a")))) + ""))) {
                                        case "a":
                                            var v50 uint8 = v49
                                            _ = v50
                                        case "hello":
                                            v49 += 0
                                        case "foo bar":
                                            v49 = (^(v49))
                                        default:
                                            if v1 {
                                                if ((!v1) || (!true)) {
                                                    fmt.Println(v49, v2)
                                                    v48 -= (v48 + v48)
                                                }
                                                for fi51 := 0; fi51 < 16; fi51++ {
                                                    v2 -= 42
                                                    v2 -= (v2 >> ((4294967295) & 63))
                                                    _ = fi51
                                                }
                                                v48 += (((v48 * (v48 + float64(float32(v48)))) / (-2.25 / ((v48 + v48) / v48))) + -2.25)
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = (^(fi29))
                                                }
                                            }
                                        }
                                        _ = v48
                                        _ = v49
                                    }
                                default:
                                    v1 = (42 >= -1)
                                }
                                var v52 uint64 = 4294967295
                                var v53 uint16 = 255
                                _ = v52
                                _ = v53
                            }
                        }
                        if v1 {
                            var v54 string = ("" + "")
                            _ = v54
                        } else {
                            switch "a" {
                            case "\n\t":
                                for fi55 := 0; fi55 < 9; fi55++ {
                                    var v56 uint64 = v2
                                    _ = fi55
                                    _ = v56
                                }
                            case "foo bar":
                                var v57 bool = v1
                                _ = v57
                            case "hello":
                                v2 += (^(v2))
                            default:
                                var v58 uint32 = 4294967295
                                _ = v58
                            }
                        }
                    }
                    fmt.Println(v2)
                    _ = fi29
                }
                for fi59 := 0; fi59 < 17; fi59++ {
                    switch "" {
                    case "hello":
                        for fi60 := 0; fi60 < 8; fi60++ {
                            switch (-(fi60)) {
                            case 2147483647:
                                var v61 string = "foo bar"
                                _ = v61
                            case 32767:
                                if v1 {
                                    v2 += 23
                                    if (v1 && v1) {
                                        switch ((("hello" + "") + "hello") + ("foo bar" + ("foo bar" + "a"))) {
                                        case "":
                                            var v62 float32 = 3.14159
                                            _ = v62
                                        default:
                                            v63 := GenS4{F0: 1000, F1: ("" + "")}
                                            _ = v63
                                        }
                                        v1 = (-1 > -374)
                                        var v64 int32 = 0
                                        _ = v64
                                    } else {
                                        if (!v1) {
                                            v1 = (882 != 1)
                                        } else {
                                            var v65 int16 = 1000
                                            v65 = (-(v65))
                                            v65 += 32767
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = v65
                                        }
                                        var v66 uint16 = 32767
                                        var v67 bool = false
                                        _ = v66
                                        _ = v67
                                    }
                                    v68 := GenS5{F0: v2, F1: 126, F2: 111}
                                    _ = v68
                                } else {
                                    var v69 uint32 = 42
                                    fmt.Println(v2)
                                    _ = len(v3)
                                    _ = v69
                                }
                            case 725:
                                if true {
                                    v1 = (v2 >= 0)
                                }
                            }
                            _ = fi60
                        }
                    case "\n\t":
                        for fi70 := 0; fi70 < 16; fi70++ {
                            if false {
                                if (128 >= 108) {
                                    switch ("" + (("a" + "hello") + "foo bar")) {
                                    case "hello":
                                        var v71 int = (fi70 / ((fi70) | 1))
                                        _ = v71
                                    case "foo bar":
                                        var v72 int64 = 0
                                        _ = v72
                                    case "":
                                        v1 = v1
                                    default:
                                        switch genFn0() {
                                        case 128:
                                            if ((v1 || false) || (820 != 1)) {
                                                if (0 > 2147483647) {
                                                    v2 = 301
                                                }
                                                v1 = true
                                                v2 = 18446744073709551615
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = "a"
                                                }
                                                v73 := GenS3{F0: -2147483648}
                                                var v74 uint16 = 450
                                                _ = v73
                                                _ = v74
                                            }
                                        default:
                                            var v75 string = "hello"
                                            _ = v75
                                        }
                                    }
                                    var v76 uint64 = (^(v2))
                                    _ = v76
                                }
                                var v77 int8 = 0
                                _ = v77
                            }
                            v3 = append(v3, genFn0())
                            fmt.Println(v1)
                            _ = fi70
                        }
                    case "":
                        if v1 {
                            var v78 float32 = (1e10 - 0.0)
                            switch "hello" {
                            case "foo bar":
                                if (v1 && (uint8(v2) == 1)) {
                                    switch fi59 {
                                    case 949:
                                        v3 = append(v3, (^(fi59)))
                                    }
                                    if false {
                                        for fi79 := 0; fi79 < 15; fi79++ {
                                            switch fi79 {
                                            case 32768:
                                                v78 = (float32(1e10) * ((v78 / v78) * 0.0))
                                            case 9223372036854775807:
                                                var v80 int16 = -199
                                                _ = v80
                                            }
                                            fmt.Println(v78, fi59, v1)
                                            _ = fi79
                                        }
                                        var v81 int32 = 2147483647
                                        _ = v81
                                    }
                                } else {
                                    if (v1 && true) {
                                        if true {
                                            v82 := GenS6{F0: 98}
                                            _ = v82
                                        }
                                    } else {
                                        v4 = append(v4, "foo bar")
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = ((((("hello" + "") + "foo bar") + "\n\t") + "hello") + "hello")
                                    }
                                    v78 = v78
                                }
                            case "":
                                fmt.Println(v2, v78, v78)
                            case "hello":
                                v83 := GenS5{F0: (v2 >> ((2) & 63)), F1: -1, F2: uint(v2)}
                                _ = v83
                            default:
                                fmt.Println(fi59)
                            }
                            _ = v78
                        }
                    }
                    _ = fi59
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                v1 = v1
            }
        }
    }
    v1 = (!v1)
    v84 := GenS3{F0: -19}
    v2 -= (v2 + 484)
    for fi85 := 0; fi85 < 17; fi85++ {
        v2 -= (v2 - v2)
        if ("" <= "\n\t") {
            fmt.Println(v1)
        } else {
            var v86 int = 813
            v87 := GenS6{F0: 0}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v86
            _ = v87
        }
        _ = fi85
    }
    v88 := GenS6{F0: 1}
    v89 := GenS1{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v84
    _ = v88
    _ = v89
}
