package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint32; F1 int64 }
func genFn0(p0 uint) string {
    var v7 int = 128
    switch (v7 << ((p0) & 63)) {
    case 128:
        p0 = p0
    default:
        if false {
            var v8 int32 = 430
            v9 := GenS1{F0: -9223372036854775808, F1: ("foo bar" + "a")}
            for fi10 := 0; fi10 < 14; fi10++ {
                if (uint64(p0) >= 771) {
                    v11 := []int{fi10, 1, (fi10 ^ fi10)}
                    if false {
                        if false {
                            if false {
                                v8 = (v8 & v8)
                                fmt.Println(v8, v7)
                                for fi12 := 0; fi12 < 9; fi12++ {
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    _ = len(v11)
                                    _ = fi12
                                }
                            } else {
                                v13 := GenS1{F0: 0, F1: "foo bar"}
                                var v14 int = (v7 ^ -1)
                                _ = v13
                                _ = v14
                            }
                            v15 := GenS1{F0: 681, F1: (("hello" + ((("foo bar" + "") + ("\n\t" + ("" + ((("" + "hello") + "\n\t") + ("foo bar" + ((("a" + "hello") + ("hello" + ("" + "hello"))) + "a")))))) + "foo bar")) + (((("hello" + "\n\t") + "\n\t") + ("foo bar" + ("\n\t" + "foo bar"))) + ""))}
                            v16 := GenS2{F0: 338, F1: 38}
                            _ = v15
                            _ = v16
                        } else {
                            var v17 string = "foo bar"
                            v7 += fi10
                            if (-333 <= -1) {
                                v18 := GenS1{F0: fi10, F1: ("" + "hello")}
                                v19 := GenS1{F0: 127, F1: v17}
                                _ = v18
                                _ = v19
                            }
                            _ = v17
                        }
                        var v20 int8 = 1
                        var v21 uint16 = 2
                        _ = v20
                        _ = v21
                    } else {
                        var v22 int = -1
                        _ = v22
                    }
                    _ = v11
                } else {
                    v23 := GenS2{F0: 255, F1: 1}
                    var v24 float32 = 1.5
                    _ = v23
                    _ = v24
                }
                if false {
                    v8 = 42
                    v25 := []int{(^(fi10)), fi10}
                    var v26 int64 = 9223372036854775807
                    _ = v25
                    _ = v26
                } else {
                    v7 += (fi10 - (fi10 / (((^(v7))) | 1)))
                }
                for fi27 := 0; fi27 < 3; fi27++ {
                    v28 := GenS3{F0: 18446744073709551615}
                    v29 := GenS3{F0: 42}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = fi10
            }
            _ = v8
            _ = v9
        } else {
            var v30 string = "a"
            var v31 uint = (^(p0))
            _ = v30
            _ = v31
        }
    }
    if (uint8(p0) < 126) {
        if false {
            if (true || true) {
                v32 := GenS3{F0: 18446744073709551615}
                v33 := []int{}
                _ = v32
                _ = v33
            } else {
                var v34 uint16 = 255
                v35 := []int{}
                _ = len(v35)
                _ = v34
                _ = v35
            }
        }
    } else {
        var v36 int64 = -608
        if true {
            if false {
                var v37 int32 = 42
                v38 := GenS4{F0: ("foo bar" >= ("hello" + "hello")), F1: 17}
                v37 = v37
                _ = v37
                _ = v38
            } else {
                if true {
                    v39 := GenS2{F0: 0, F1: 4294967295}
                    _ = v39
                } else {
                    if (32767 != 255) {
                        if (false || (v7 == (^(v7)))) {
                            v40 := []int{}
                            if ("\n\t" == "hello") {
                                p0 -= 1
                                v7 -= v7
                                var v41 int8 = -1
                                _ = v41
                            }
                            v7 -= (v7 - v7)
                            _ = v40
                        }
                    }
                    switch v7 {
                    case 823:
                        fmt.Println(p0)
                    case 42:
                        p0 += p0
                    case 2147483647:
                        var v42 int16 = 42
                        _ = v42
                    }
                    var v43 float32 = float32(1.5)
                    _ = v43
                }
                var v44 int8 = 127
                p0 = (p0 ^ (p0 << ((p0) & 63)))
                _ = v44
            }
        }
        _ = v36
    }
    v45 := GenS5{F0: 314, F1: 9223372036854775807}
    v46 := []int{2147483647, (v7 >> ((p0) & 63)), (v7 ^ 380)}
    _ = p0
    _ = v45
    _ = v46
    _ = v7
    return "hello"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 42
    v3 := []int{-466, -62, -2147483648, 128}
    v4 := []int{26, -9223372036854775808, 0, 9223372036854775807}
    v5 := make(map[uint32]uint32)
    var v6 uint16 = 604
    var v47 uint16 = v6
    v6 = 42
    _ = len(v3)
    v47 = (v47 % ((v47) | 1))
    _ = len(v1)
    switch "\n\t" {
    case "":
        switch 1 {
        case -1:
            fmt.Println(v47, v6, v2)
        default:
            v2 = (v2 ^ v2)
        }
    case "foo bar":
        if false {
            switch "" {
            case "foo bar":
                if false {
                    for fi48 := 0; fi48 < 15; fi48++ {
                        switch ("a" + "a") {
                        case "":
                            v49 := GenS3{F0: 18446744073709551615}
                            _ = v49
                        }
                        fmt.Println(v6)
                        v50 := GenS3{F0: 47}
                        _ = fi48
                        _ = v50
                    }
                    var v51 bool = (42 != 0)
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                    _ = v51
                }
            case "":
                if false {
                    for fi52 := 0; fi52 < 7; fi52++ {
                        fmt.Println(v6)
                        if (!(1 > 42)) {
                            var v53 int = 4294967295
                            var v54 int8 = 42
                            _ = v53
                            _ = v54
                        } else {
                            for fi55 := 0; fi55 < 7; fi55++ {
                                switch ("hello" + ("" + ("hello" + ""))) {
                                case "\n\t":
                                    if false {
                                        v56 := GenS5{F0: 2147483648, F1: 873}
                                        _ = len(v3)
                                        fmt.Println(fi55, v47)
                                        _ = v56
                                    }
                                    fmt.Println(v6)
                                case "":
                                    if true {
                                        var v57 string = ("hello" + "a")
                                        var v58 int64 = 0
                                        v4 = append(v4, (fi55 | fi55))
                                        _ = v57
                                        _ = v58
                                    }
                                case "foo bar":
                                    v2 -= 42
                                default:
                                    switch "" {
                                    case "hello":
                                        v59 := GenS3{F0: 4294967295}
                                        _ = v59
                                    case "a":
                                        var v60 uint16 = v47
                                        _ = v60
                                    default:
                                        if (0.0 >= 0.1) {
                                            var v61 int8 = 42
                                            _ = v61
                                        } else {
                                            var v62 int32 = -162
                                            var v63 int32 = int32(fi52)
                                            switch ("hello" + "hello") {
                                            case "a":
                                                var v64 int32 = (v63 / ((-803) | 1))
                                                _ = v64
                                            case "":
                                                v6 -= (v6 * 65535)
                                                if (!false) {
                                                    switch "a" {
                                                    case "hello":
                                                        var v65 bool = ((0.1 * ((0.1 + float64(0.0)) / 1e10)) == 0.0)
                                                        _ = v65
                                                    case "\n\t":
                                                        switch (^(fi52)) {
                                                        case -2147483648:
                                                            v6 += v47
                                                        }
                                                        if (0.1 == 0.0) {
                                                            fmt.Println(v63, v47)
                                                            if len(v3) > 0 {
                                                                v3[0] = fi52
                                                            }
                                                            v62 = v63
                                                        } else {
                                                            if ((0.1 * 1e10) != 0.1) {
                                                                v63 -= int32(fi55)
                                                                fmt.Println(v62, v2, v2)
                                                                v47 += (v47 % ((v6) | 1))
                                                                switch 9223372036854775807 {
                                                                case 2147483647:
                                                                    switch fi52 {
                                                                    case 939:
                                                                        switch (("a" + "") + ("hello" + "")) {
                                                                        case "a":
                                                                            var v66 uint = 412
                                                                            _ = v66
                                                                        case "\n\t":
                                                                            var v67 int32 = 42
                                                                            _ = v67
                                                                        case "":
                                                                            var v68 int8 = 127
                                                                            _ = v68
                                                                        default:
                                                                            if (!false) {
                                                                                v47 = (v47 + (v6 - (v6 ^ 32767)))
                                                                                var v69 uint32 = 616
                                                                                var v70 uint32 = (v69 - (v69 + (v69 << ((697) & 63))))
                                                                                _ = v69
                                                                                _ = v70
                                                                            } else {
                                                                                v6 -= v47
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                v62 = v63
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    if true {
                                                                        fmt.Println(fi55)
                                                                        var v71 uint64 = 193
                                                                        v72 := GenS5{F0: uint32(v71), F1: 270}
                                                                        if (1 <= 65535) {
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                switch "" {
                                                                                case "":
                                                                                    var v73 int64 = -9223372036854775808
                                                                                    _ = v73
                                                                                }
                                                                            case "hello":
                                                                                _ = len(v5)
                                                                            case "a":
                                                                                switch (fi52 & (fi55 / ((-2147483648) | 1))) {
                                                                                case 9223372036854775807:
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                case 134:
                                                                                    fmt.Println(v71)
                                                                                default:
                                                                                    fmt.Println(v71, v6, fi52)
                                                                                    var v74 int8 = 1
                                                                                    var v75 float32 = -2.25
                                                                                    _ = v74
                                                                                    _ = v75
                                                                                }
                                                                            default:
                                                                                v6 = v6
                                                                            }
                                                                            var v76 uint8 = uint8(v71)
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            _ = v76
                                                                        } else {
                                                                            v62 = v63
                                                                            v6 += v6
                                                                            v71 -= (^(v71))
                                                                        }
                                                                        fmt.Println(v6, v6)
                                                                        _ = v71
                                                                        _ = v72
                                                                    }
                                                                }
                                                            }
                                                            var v77 uint = 0
                                                            _ = v77
                                                        }
                                                        var v78 bool = (int64(v63) >= -9223372036854775808)
                                                        _ = v78
                                                    }
                                                    v6 = 795
                                                } else {
                                                    fmt.Println(v2)
                                                    v79 := GenS4{F0: true, F1: 42}
                                                    _ = v79
                                                }
                                                var v80 uint8 = v2
                                                _ = v80
                                            case "\n\t":
                                                v6 -= v47
                                            }
                                            _ = v62
                                            _ = v63
                                        }
                                    }
                                }
                                _ = fi55
                            }
                            if true {
                                switch fi52 {
                                case -11:
                                    fmt.Println(v47, v6, fi52)
                                default:
                                    _ = len(v4)
                                }
                                v6 = uint16(v2)
                                v2 = 42
                            }
                            fmt.Println(fi52, v6)
                        }
                        _ = fi52
                    }
                    var v81 int = -958
                    v82 := GenS1{F0: (v81 << ((2) & 63)), F1: ("hello" + "")}
                    _ = v81
                    _ = v82
                } else {
                    switch 281 {
                    case 1:
                        fmt.Println(v6, v6)
                    default:
                        var v83 float64 = 1e10
                        _ = v83
                    }
                }
            case "hello":
                var v84 bool = true
                _ = v84
            }
        }
    case "a":
        var v85 int = -9223372036854775808
        _ = v85
    }
    _ = len(v3)
    if false {
        switch 739 {
        case 240:
            v47 = v47
        case 2147483647:
            var v86 int16 = -391
            _ = v86
        case 0:
            for fi87 := 0; fi87 < 15; fi87++ {
                var v88 string = (("foo bar" + ("a" + "foo bar")) + "a")
                for fi89 := 0; fi89 < 19; fi89++ {
                    fmt.Println(fi87)
                    v90 := GenS3{F0: 0}
                    fmt.Println(v2)
                    v2 += v2
                    _ = fi89
                    _ = v90
                }
                var v91 uint64 = 4294967295
                _ = fi87
                _ = v88
                _ = v91
            }
        default:
            var v92 uint32 = 916
            _ = v92
        }
    } else {
        v4 = append(v4, -1)
        v3 = append(v3, -166)
    }
    switch "hello" {
    case "hello":
        fmt.Println(v6)
    case "":
        if len(v3) > 0 {
            v3[0] = 32767
        }
    case "\n\t":
        v93 := GenS3{F0: 50}
        _ = v93
    default:
        if false {
            for fi94 := 0; fi94 < 16; fi94++ {
                v47 += (v47 & v47)
                _ = fi94
            }
            v6 = (^(v47))
        } else {
            v6 += v6
        }
    }
    delete(v1, "foo bar")
    fmt.Println(v2, v6)
    if (true || (2 < 954)) {
        if (42 < 42) {
            var v95 uint8 = v2
            switch "a" {
            case "a":
                v3 = append(v3, -155)
            case "foo bar":
                for fi96 := 0; fi96 < 2; fi96++ {
                    for fi97 := 0; fi97 < 7; fi97++ {
                        v98 := GenS4{F0: (18446744073709551615 > 18446744073709551615), F1: 2147483647}
                        _ = fi97
                        _ = v98
                    }
                    v95 = (^(v2))
                    _ = fi96
                }
            case "hello":
                var v99 int32 = -2147483648
                _ = v99
            }
            fmt.Println(v6)
            _ = v95
        }
        delete(v5, 805)
        for fi100 := 0; fi100 < 13; fi100++ {
            fmt.Println(fi100, v6)
            if false {
                var v101 uint64 = 2
                v47 = (v47 / (((^(v47))) | 1))
                for fi102 := 0; fi102 < 4; fi102++ {
                    var v103 int = (fi100 ^ fi102)
                    var v104 int32 = 1
                    _ = fi102
                    _ = v103
                    _ = v104
                }
                _ = v101
            }
            v6 = 275
            _ = fi100
        }
    }
    var v105 uint8 = v2
    var v106 int = -1
    switch "hello" {
    case "foo bar":
        for fi107 := 0; fi107 < 5; fi107++ {
            v47 += (v6 / ((v6) | 1))
            v108 := GenS4{F0: false, F1: 4294967295}
            v109 := GenS3{F0: 18446744073709551615}
            _ = fi107
            _ = v108
            _ = v109
        }
    case "a":
        v110 := GenS2{F0: v47, F1: 103}
        _ = v110
    case "hello":
        v106 += (^(v106))
    default:
        _ = len(v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v105
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
