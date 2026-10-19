package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint16; F2 string }
type GenS4 struct { F0 uint32; F1 int8; F2 int }
type GenS5 struct { F0 string; F1 uint8; F2 uint64 }
type GenS6 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int }
func genFn0(p0 string) int32 {
    v7 := GenS1{F0: 65535, F1: 593}
    v8 := []int{42, 454, 125, -457}
    _ = p0
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 int16 = -32768
    var v2 uint32 = 2147483647
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    var v5 int64 = -1
    v6 := make(map[int]int)
    var v9 uint = 56
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS1{F0: v9, F1: (fi10 - fi10)}
        v12 := GenS1{F0: v9, F1: (fi10 & (fi10 + fi10))}
        if false {
            if false {
                switch fi10 {
                case -9223372036854775808:
                    v13 := GenS1{F0: v9, F1: (fi10 / ((32767) | 1))}
                    _ = v13
                default:
                    var v14 uint64 = 42
                    _ = v14
                }
            }
        } else {
            var v15 uint8 = 127
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 string = "a"
                switch (-(fi10)) {
                case 42:
                    v2 -= (^(v2))
                case 4294967295:
                    if (0 < -128) {
                        if (!true) {
                            var v18 int = fi10
                            _ = v18
                        }
                        v19 := []int{}
                        _ = v19
                    }
                case 32768:
                    v20 := []int{(fi10 - (fi10 + fi10)), (fi16 * fi16), -2, fi16}
                    _ = v20
                default:
                    if (true && false) {
                        v21 := GenS2{F0: 42}
                        var v22 uint8 = (v15 << (((v9 | 18446744073709551615)) & 63))
                        _ = v21
                        _ = v22
                    }
                }
                _ = fi16
                _ = v17
            }
            v23 := GenS3{F0: 0, F1: 222, F2: "hello"}
            _ = v15
            _ = v23
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v5)
    fmt.Println(v5)
    switch (("a" + ((("" + "\n\t") + "") + (("\n\t" + "a") + "hello"))) + "a") {
    case "foo bar":
        fmt.Println(v2, v9, v9)
    case "a":
        switch ("" + "\n\t") {
        case "foo bar":
            if true {
                if true {
                    v24 := GenS2{F0: 0}
                    v25 := GenS4{F0: (v2 << ((v9) & 63)), F1: -1, F2: -2147483648}
                    var v26 int8 = -65
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
        default:
            for fi27 := 0; fi27 < 1; fi27++ {
                for fi28 := 0; fi28 < 16; fi28++ {
                    v29 := GenS4{F0: 625, F1: 0, F2: fi28}
                    var v30 bool = (-753 != -453)
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v2 += (v2 << (((v9 << ((v9) & 63))) & 63))
                _ = fi27
            }
        }
    case "":
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 uint32 = (v2 * v2)
            v2 -= (v2 & 454)
            _ = fi31
            _ = v32
        }
    default:
        v33 := GenS3{F0: 100, F1: 159, F2: "foo bar"}
        _ = v33
    }
    switch (("foo bar" + "hello") + "a") {
    case "hello":
        switch 1 {
        case 2147483647:
            v5 = (v5 - v5)
        case -477:
            fmt.Println(v5)
        case 990:
            if ("hello" < "foo bar") {
                v2 = 422
                v5 = (v5 & 1)
            }
        default:
            v9 = v9
        }
    case "\n\t":
        var v34 float64 = (3.14159 / 3.14159)
        _ = v34
    }
    switch 42 {
    case 127:
        for fi35 := 0; fi35 < 8; fi35++ {
            switch "" {
            case "foo bar":
                var v36 float32 = (3.14159 + 1e10)
                _ = v36
            case "hello":
                for fi37 := 0; fi37 < 14; fi37++ {
                    v5 = v5
                    fmt.Println(v2, v1, v9)
                    _ = fi37
                }
            case "\n\t":
                for fi38 := 0; fi38 < 16; fi38++ {
                    var v39 int = fi35
                    if ("\n\t" <= "") {
                        v40 := GenS3{F0: -128, F1: 255, F2: "a"}
                        if true {
                            if (0 > genFn0("\n\t")) {
                                switch "\n\t" {
                                case "a":
                                    var v41 string = "hello"
                                    _ = v41
                                }
                            } else {
                                if (1 >= v2) {
                                    if (!(!true)) {
                                        v1 = v1
                                    } else {
                                        switch (fi38 << (((v9 - v9)) & 63)) {
                                        case 32768:
                                            _ = len(v6)
                                        case 0:
                                            v5 -= 422
                                        default:
                                            v42 := v4["\n\t"]
                                            _ = v42
                                        }
                                        v43 := []int{}
                                        fmt.Println(v2, v9)
                                        _ = v43
                                    }
                                    v3[(fi35 + fi35)] = (fi35 == (fi35 * fi35))
                                    v44 := GenS4{F0: (v2 | 328), F1: -60, F2: fi35}
                                    _ = v44
                                } else {
                                    switch (-(fi38)) {
                                    case 128:
                                        var v45 float32 = float32(-2.25)
                                        _ = v45
                                    case 347:
                                        v9 -= (v9 & (v9 * v9))
                                    default:
                                        switch (fi35 | (v39 >> (((v9 - 0)) & 63))) {
                                        case -1:
                                            var v46 int32 = 1
                                            _ = v46
                                        case -9223372036854775808:
                                            if false {
                                                fmt.Println(v5, v5, v1)
                                                var v47 int = 968
                                                v2 -= v2
                                                _ = v47
                                            } else {
                                                switch fi38 {
                                                case 42:
                                                    if true {
                                                        v48 := GenS1{F0: 4294967295, F1: (^(fi35))}
                                                        _ = v48
                                                    } else {
                                                        switch (^(v39)) {
                                                        case 32767:
                                                            var v49 float32 = ((1e10 + -2.25) - (3.14159 + float32(0.0)))
                                                            v50 := GenS4{F0: uint32(v9), F1: 114, F2: v39}
                                                            fmt.Println(fi35, fi35)
                                                            v51 := GenS5{F0: "hello", F1: 1, F2: 42}
                                                            _ = v49
                                                            _ = v50
                                                            _ = v51
                                                        case -375:
                                                            var v52 int32 = -2147483648
                                                            _ = v52
                                                        }
                                                        v2 -= v2
                                                        var v53 uint = (v9 | v9)
                                                        _ = v53
                                                    }
                                                    var v54 bool = false
                                                    var v55 int16 = -1
                                                    _ = v54
                                                    _ = v55
                                                case 4294967295:
                                                    if false {
                                                        var v56 float32 = 0.1
                                                        fmt.Println(fi38)
                                                        _ = v56
                                                    } else {
                                                        v39 = 32768
                                                        if true {
                                                            v9 = (v9 - (v9 & 65535))
                                                        } else {
                                                            var v57 float32 = 3.14159
                                                            fmt.Println(fi38, fi35)
                                                            _ = v57
                                                        }
                                                        v58 := GenS2{F0: (^(v2))}
                                                        _ = v58
                                                    }
                                                    v1 -= -897
                                                case 103:
                                                    var v59 float32 = 1.5
                                                    _ = v59
                                                }
                                                var v60 uint16 = 64
                                                var v61 uint8 = 198
                                                _ = v60
                                                _ = v61
                                            }
                                        case 2147483647:
                                            fmt.Println(v1, v9, v2)
                                        default:
                                            v9 += v9
                                        }
                                    }
                                }
                            }
                        } else {
                            switch ("a" + ("hello" + (((("hello" + "") + "\n\t") + ("\n\t" + "foo bar")) + "foo bar"))) {
                            case "":
                                fmt.Println(fi38, v5, v2)
                            case "\n\t":
                                v9 = v9
                            }
                            var v62 int = fi38
                            fmt.Println(v2, fi38)
                            _ = v62
                        }
                        v63 := GenS2{F0: v2}
                        _ = v40
                        _ = v63
                    } else {
                        if (!false) {
                            v9 -= v9
                            switch "hello" {
                            case "foo bar":
                                v64 := GenS3{F0: 42, F1: 711, F2: ("" + "hello")}
                                _ = v64
                            default:
                                v9 -= (v9 / (((v9 | (v9 * 2))) | 1))
                                var v65 int32 = -2147483648
                                _ = v65
                            }
                            v39 = int(v9)
                            v66 := v3[fi38]
                            v67 := GenS1{F0: v9, F1: (fi38 ^ fi35)}
                            v1 -= v1
                            _ = v66
                            _ = v67
                        }
                        var v68 uint16 = 1
                        fmt.Println(v68, fi38, v1)
                        _ = v68
                    }
                    if true {
                        if ((((((0.0 / 3.14159) - float64(1e10)) * 1.5) / 1e10) - float64(0.0)) < ((((0.0 * 1.5) / (1.5 + 1e10)) - (1e10 * -2.25)) * 1e10)) {
                            var v69 string = ("\n\t" + ("\n\t" + "hello"))
                            var v70 uint16 = uint16(v2)
                            if false {
                                fmt.Println(v70, v2, v69)
                            } else {
                                v39 = fi35
                                var v71 uint = 255
                                v72 := GenS6{F0: v2, F1: (^(v5)), F2: 255}
                                _ = v71
                                _ = v72
                            }
                            _ = v69
                            _ = v70
                        } else {
                            fmt.Println(v9)
                            switch fi38 {
                            case -3:
                                fmt.Println(v9, v9)
                            default:
                                if (v5 > v5) {
                                    v73 := []int{}
                                    var v74 uint32 = (v2 & v2)
                                    _ = v73
                                    _ = v74
                                } else {
                                    switch ("foo bar" + "") {
                                    case "a":
                                        var v75 int8 = -128
                                        _ = v75
                                    case "":
                                        var v76 int8 = -89
                                        _ = v76
                                    }
                                    v77 := GenS6{F0: (^(v2)), F1: -9223372036854775808, F2: 0}
                                    switch (fi38 - (^(fi38))) {
                                    case 2147483647:
                                        switch "\n\t" {
                                        case "foo bar":
                                            fmt.Println(v2)
                                        default:
                                            switch (fi38 << ((v9) & 63)) {
                                            case -1:
                                                var v78 uint64 = 4294967295
                                                _ = v78
                                            }
                                            v79 := GenS7{F0: 558}
                                            _ = v79
                                        }
                                    }
                                    _ = v77
                                }
                            }
                        }
                        v5 = -9223372036854775808
                    }
                    _ = fi38
                    _ = v39
                }
            }
            switch (fi35 % (((fi35 + 9223372036854775807)) | 1)) {
            case -1:
                switch "\n\t" {
                case "":
                    v2 = (v2 >> (((v9 & v9)) & 63))
                case "hello":
                    v9 = 1
                case "\n\t":
                    var v80 int16 = 32767
                    _ = v80
                default:
                    if (false || true) {
                        switch (fi35 & 0) {
                        case 127:
                            var v81 uint8 = 127
                            _ = v81
                        case -1:
                            var v82 bool = (!false)
                            _ = v82
                        }
                        if (1 == -2147483648) {
                            for fi83 := 0; fi83 < 16; fi83++ {
                                switch "\n\t" {
                                case "\n\t":
                                    if (4294967295 != fi83) {
                                        var v84 bool = true
                                        var v85 uint64 = 4294967295
                                        v86 := GenS8{F0: 1}
                                        v2 -= (v2 * uint32(v85))
                                        var v87 int8 = 100
                                        _ = v84
                                        _ = v85
                                        _ = v86
                                        _ = v87
                                    }
                                }
                                _ = fi83
                            }
                            fmt.Println(v2, v5)
                        } else {
                            for fi88 := 0; fi88 < 12; fi88++ {
                                v2 = v2
                                _ = fi88
                            }
                            v89 := v4[(("hello" + "a") + "a")]
                            _ = v89
                        }
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            v90 := []int{(fi35 / ((42) | 1)), (fi35 << ((18446744073709551615) & 63))}
                            _ = v90
                        case "\n\t":
                            for fi91 := 0; fi91 < 15; fi91++ {
                                if true {
                                    var v92 int = int(v1)
                                    v93 := GenS5{F0: "", F1: 77, F2: 214}
                                    _ = v92
                                    _ = v93
                                } else {
                                    v94 := GenS6{F0: (v2 - (v2 * (v2 | v2))), F1: (-(v5)), F2: 228}
                                    v1 += v1
                                    _ = v94
                                }
                                switch "a" {
                                case "a":
                                    v95 := []int{}
                                    _ = v95
                                case "foo bar":
                                    var v96 int64 = (v5 ^ 840)
                                    _ = v96
                                case "":
                                    var v97 uint = 4294967295
                                    _ = v97
                                default:
                                    v9 -= 255
                                }
                                _ = fi91
                            }
                        case "a":
                            var v98 int16 = (v1 % (((v1 * 42)) | 1))
                            _ = v98
                        }
                    }
                }
            default:
                if (458 >= 0) {
                    var v99 float32 = -2.25
                    _ = v99
                }
            }
            _ = fi35
        }
    case -1:
        var v100 uint64 = 2
        _ = v100
    }
    v1 = (^(v1))
    var v101 int8 = -1
    v102 := GenS9{F0: 128}
    var v103 float32 = ((0.0 / 1.5) + 0.1)
    v104 := v4["\n\t"]
    v9 = v9
    var v105 float32 = 1e10
    delete(v6, -2147483648)
    if (0 > v1) {
        for fi106 := 0; fi106 < 7; fi106++ {
            var v107 uint16 = 65535
            switch (fi106 * (fi106 / (((fi106 << (((v9 | v9)) & 63))) | 1))) {
            case -2147483648:
                var v108 uint32 = (^(v2))
                _ = v108
            default:
                fmt.Println(v2, v1, v1)
            }
            _ = fi106
            _ = v107
        }
        _ = len(v4)
    } else {
        v101 = 100
        fmt.Println(v101, v101, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v101
    _ = v102
    _ = v103
    _ = v104
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
