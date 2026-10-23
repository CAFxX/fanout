package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 string; F1 uint64; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 uint8; F2 string }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 uint8; F1 int8; F2 int16 }
type GenS9 struct { F0 uint32; F1 string; F2 int8 }
func genFn0(p0 string) int8 {
    v7 := []int{-25, 4294967295}
    p0 = ""
    p0 = "foo bar"
    _ = p0
    _ = v7
    return 100
}
func genFn1() uint64 {
    switch "\n\t" {
    case "a":
        var v8 int = 32767
        _ = v8
    }
    var v9 int32 = -1
    switch ("foo bar" + "hello") {
    case "foo bar":
        switch 4294967295 {
        case -264:
            v10 := []int{42, 32768, 0, -109}
            _ = v10
        case 1:
            var v11 float32 = 3.14159
            _ = v11
        default:
            switch 982 {
            case 9223372036854775807:
                switch 32767 {
                case -234:
                    if true {
                        if false {
                            v9 = v9
                            v9 = (v9 >> ((2) & 63))
                        } else {
                            fmt.Println(v9)
                            if false {
                                v9 += 666
                            }
                            v9 -= (v9 - (v9 * v9))
                        }
                        fmt.Println(v9)
                    } else {
                        v9 = (v9 % ((0) | 1))
                        var v12 uint8 = 1
                        _ = v12
                    }
                case 127:
                    v9 -= -1
                default:
                    if (3.14159 > 1.5) {
                        v9 -= 2147483647
                        v13 := []int{1}
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        _ = v13
                    }
                }
            case 4294967295:
                v14 := []int{-948, 334}
                _ = v14
            }
        }
    default:
        v9 += v9
    }
    fmt.Println(v9)
    v15 := GenS1{F0: 115, F1: 1}
    _ = v15
    _ = v9
    return 0
}

func main() {
    var v1 int16 = 1
    var v2 int16 = (v1 >> ((94) & 63))
    var v3 int32 = -720
    var v4 uint32 = 65535
    var v5 uint64 = 505
    var v6 float32 = -2.25
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 uint = 0
        _ = fi16
        _ = v17
    }
    fmt.Println(v1, v2, v4)
    var v18 string = "foo bar"
    v19 := []int{-202, 0}
    var v20 int = int(v5)
    _ = len(v19)
    for fi21 := 0; fi21 < 18; fi21++ {
        v18 = v18
        switch (fi21 * fi21) {
        case 32767:
            switch v20 {
            case 134:
                if true {
                    var v22 int8 = genFn0("\n\t")
                    v19 = append(v19, fi21)
                    _ = v22
                } else {
                    _ = len(v19)
                    if (true && (v5 < (v5 >> ((0) & 63)))) {
                        var v23 int64 = 9223372036854775807
                        _ = v23
                    } else {
                        v24 := GenS1{F0: (^(fi21)), F1: 2}
                        _ = v24
                    }
                    if (!(!(0 > 824))) {
                        v6 -= v6
                        if true {
                            v18 = v18
                            for fi25 := 0; fi25 < 5; fi25++ {
                                var v26 int8 = 42
                                _ = fi25
                                _ = v26
                            }
                        } else {
                            v3 = (v3 | (v3 % ((v3) | 1)))
                            switch (fi21 ^ (v20 | int(v3))) {
                            case 0:
                                v19 = append(v19, (fi21 >> ((940) & 63)))
                            case 786:
                                if true {
                                    for fi27 := 0; fi27 < 10; fi27++ {
                                        if (true || (-1 > (v3 + (v3 + v3)))) {
                                            var v28 uint = 0
                                            switch ("a" + v18) {
                                            case "":
                                                if true {
                                                    v1 = 371
                                                }
                                            case "foo bar":
                                                v6 = v6
                                            case "a":
                                                var v29 bool = false
                                                _ = v29
                                            default:
                                                if (v3 == -1) {
                                                    if false {
                                                        v30 := GenS2{F0: v28}
                                                        var v31 bool = true
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                    v28 = v28
                                                    v6 += v6
                                                } else {
                                                    _ = len(v19)
                                                    if (!(255 != 255)) {
                                                        if ((v4 | (^(v4))) < v4) {
                                                            v1 = (v2 % ((v1) | 1))
                                                            var v32 string = "foo bar"
                                                            v20 = -1
                                                            v33 := GenS1{F0: fi21, F1: 18446744073709551615}
                                                            _ = v32
                                                            _ = v33
                                                        } else {
                                                            switch (fi27 % ((int(v5)) | 1)) {
                                                            case -2147483648:
                                                                _ = len(v19)
                                                            case 2147483647:
                                                                fmt.Println(v18, fi21)
                                                            default:
                                                                var v34 bool = (!false)
                                                                _ = v34
                                                            }
                                                        }
                                                        _ = len(v19)
                                                    }
                                                }
                                            }
                                            _ = v28
                                        }
                                        _ = fi27
                                    }
                                } else {
                                    if len(v19) > 0 {
                                        v19[0] = v20
                                    }
                                }
                            case 732:
                                v6 += -2.25
                            default:
                                v1 += v1
                            }
                        }
                    } else {
                        if ((-1 > -132) && ((!(!false)) && false)) {
                            var v35 string = "hello"
                            _ = v35
                        } else {
                            if true {
                                var v36 float32 = float32(-2.25)
                                var v37 string = (v18 + v18)
                                _ = v36
                                _ = v37
                            } else {
                                if false {
                                    _ = len(v19)
                                }
                            }
                            v38 := GenS1{F0: v20, F1: v5}
                            var v39 uint32 = 0
                            _ = v38
                            _ = v39
                        }
                    }
                }
            case 1:
                fmt.Println(v4, v5)
            }
        case -2147483648:
            fmt.Println(v6, v20)
        case 55:
            for fi40 := 0; fi40 < 2; fi40++ {
                if (fi21 > (fi40 - 128)) {
                    if false {
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        var v41 string = ((v18 + "hello") + ((v18 + "hello") + (v18 + ((((v18 + (v18 + (v18 + ((v18 + "\n\t") + ("a" + ""))))) + (v18 + "")) + "\n\t") + v18))))
                        v42 := GenS3{F0: v3, F1: 126}
                        _ = v41
                        _ = v42
                    } else {
                        var v43 int8 = 0
                        switch -2147483648 {
                        case 127:
                            var v44 uint64 = genFn1()
                            _ = v44
                        case 128:
                            fmt.Println(v5, v43)
                        default:
                            var v45 int = (fi21 ^ int(v3))
                            _ = v45
                        }
                        _ = v43
                    }
                    v46 := GenS3{F0: (v3 ^ v3), F1: 1}
                    v6 = v6
                    _ = v46
                } else {
                    _ = len(v19)
                }
                v47 := GenS2{F0: 42}
                v48 := GenS4{F0: v18, F1: (v5 / ((genFn1()) | 1)), F2: (v5 ^ (v5 & v5))}
                _ = fi40
                _ = v47
                _ = v48
            }
        }
        v20 += v20
        _ = fi21
    }
    v49 := GenS3{F0: v3, F1: 127}
    for fi50 := 0; fi50 < 4; fi50++ {
        switch ("" + ("a" + "a")) {
        case "foo bar":
            _ = len(v19)
        case "a":
            if (877 > uint64(v4)) {
                if true {
                    if (v3 > 1) {
                        var v51 int32 = v3
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        _ = v51
                    } else {
                        for fi52 := 0; fi52 < 1; fi52++ {
                            v53 := GenS5{F0: (4294967295 != (v5 - (^(v5))))}
                            _ = fi52
                            _ = v53
                        }
                        for fi54 := 0; fi54 < 3; fi54++ {
                            _ = len(v19)
                            v1 += v2
                            if true {
                                fmt.Println(v1, v5)
                            }
                            _ = fi54
                        }
                    }
                } else {
                    var v55 uint8 = 127
                    fmt.Println(v3, v4)
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    _ = v55
                }
                v6 -= v6
                var v56 bool = true
                _ = v56
            }
        case "hello":
            switch 4294967295 {
            case -9223372036854775808:
                var v57 int = fi50
                _ = v57
            case -1:
                v19 = append(v19, (-(v20)))
            default:
                switch fi50 {
                case -55:
                    if len(v19) > 0 {
                        v19[0] = fi50
                    }
                }
            }
        default:
            switch "\n\t" {
            case "":
                var v58 int32 = -2147483648
                _ = v58
            case "a":
                for fi59 := 0; fi59 < 11; fi59++ {
                    switch -738 {
                    case 128:
                        if (!(!false)) {
                            switch fi50 {
                            case 127:
                                var v60 int16 = v2
                                _ = v60
                            case 32767:
                                v5 -= v5
                            }
                        } else {
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                        }
                    }
                    _ = fi59
                }
            case "\n\t":
                v61 := GenS3{F0: int32(v1), F1: 150}
                _ = v61
            default:
                v62 := GenS6{F0: (v2 - (^(v1))), F1: 42, F2: (v18 + ((v18 + "foo bar") + ("hello" + v18)))}
                _ = v62
            }
        }
        v20 = (fi50 ^ (^(v20)))
        _ = fi50
    }
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    var v63 int = v20
    fmt.Println(v2)
    if (v63 < v20) {
        if (float32(((float64(1.5) + float64(v6)) - (3.14159 - (0.1 / ((float64((v6 - 3.14159)) - (-2.25 / (((1.5 / 3.14159) + 0.1) + 1.5))) / (0.0 + 1.5)))))) > v6) {
            var v64 uint16 = 65535
            v65 := GenS5{F0: (!(1 <= 42))}
            _ = v64
            _ = v65
        } else {
            var v66 uint16 = 65535
            switch (-(v63)) {
            case -2147483648:
                for fi67 := 0; fi67 < 15; fi67++ {
                    v63 += 32767
                    v4 = (v4 + v4)
                    _ = fi67
                }
            default:
                if false {
                    v68 := GenS5{F0: false}
                    _ = v68
                } else {
                    for fi69 := 0; fi69 < 19; fi69++ {
                        if len(v19) > 0 {
                            v19[0] = v63
                        }
                        if true {
                            if (true && false) {
                                v2 += v2
                                for fi70 := 0; fi70 < 2; fi70++ {
                                    v1 = v2
                                    _ = fi70
                                }
                                if len(v19) >= 2 {
                                    v19 = v19[1:len(v19)]
                                }
                            }
                            if (196 <= 316) {
                                v71 := GenS1{F0: 127, F1: 4294967295}
                                _ = v71
                            }
                            var v72 float32 = v6
                            _ = v72
                        } else {
                            if false {
                                if true {
                                    for fi73 := 0; fi73 < 17; fi73++ {
                                        switch v20 {
                                        case 0:
                                            var v74 float64 = float64(0.1)
                                            _ = v74
                                        case -9223372036854775808:
                                            v20 = -851
                                        case -143:
                                            fmt.Println(v6, v4, fi69)
                                        default:
                                            v19 = append(v19, (v20 << ((uint(v66)) & 63)))
                                        }
                                        fmt.Println(v4, v66, fi73)
                                        fmt.Println(v6)
                                        _ = fi73
                                    }
                                    switch (v18 + (v18 + (((("a" + v18) + "a") + v18) + v18))) {
                                    case "foo bar":
                                        var v75 int32 = v3
                                        _ = v75
                                    case "a":
                                        switch v63 {
                                        case 4294967295:
                                            if true {
                                                v5 = (v5 & v5)
                                            }
                                        case 127:
                                            v18 = v18
                                        default:
                                            switch fi69 {
                                            case 127:
                                                switch (("hello" + ("\n\t" + (v18 + "a"))) + v18) {
                                                case "hello":
                                                    var v76 float64 = ((-2.25 - (((-2.25 - 0.1) + -2.25) + (((1.5 + 0.0) * (float64(v6) / 1.5)) - 1e10))) / 1.5)
                                                    _ = v76
                                                case "a":
                                                    for fi77 := 0; fi77 < 8; fi77++ {
                                                        if false {
                                                            switch (v18 + "\n\t") {
                                                            case "hello":
                                                                v4 += 2
                                                            case "\n\t":
                                                                v78 := GenS6{F0: v1, F1: 77, F2: ("foo bar" + v18)}
                                                                _ = v78
                                                            }
                                                            var v79 uint = 42
                                                            v6 = float32(1.5)
                                                            var v80 uint16 = (v66 ^ v66)
                                                            v81 := GenS7{F0: false, F1: ""}
                                                            _ = v79
                                                            _ = v80
                                                            _ = v81
                                                        } else {
                                                            var v82 int32 = (^(v3))
                                                            var v83 uint64 = (v5 / ((v5) | 1))
                                                            v84 := GenS7{F0: false, F1: "a"}
                                                            _ = v82
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                        var v85 uint64 = uint64(v4)
                                                        v5 += (v5 + uint64(v66))
                                                        _ = fi77
                                                        _ = v85
                                                    }
                                                case "":
                                                    v86 := GenS8{F0: 210, F1: 42, F2: 1000}
                                                    _ = v86
                                                default:
                                                    for fi87 := 0; fi87 < 9; fi87++ {
                                                        _ = len(v19)
                                                        _ = fi87
                                                    }
                                                }
                                            case 0:
                                                v6 = ((3.14159 - float32((-2.25 / float64(v6)))) * float32(-2.25))
                                            default:
                                                var v88 uint8 = 127
                                                _ = v88
                                            }
                                        }
                                    case "hello":
                                        var v89 string = ""
                                        _ = v89
                                    default:
                                        if len(v19) >= 2 {
                                            v19 = v19[1:len(v19)]
                                        }
                                    }
                                } else {
                                    fmt.Println(v3, v1)
                                }
                                for fi90 := 0; fi90 < 12; fi90++ {
                                    if true {
                                        switch v18 {
                                        case "hello":
                                            if len(v19) >= 2 {
                                                v19 = v19[1:len(v19)]
                                            }
                                        default:
                                            v5 -= (v5 & v5)
                                        }
                                        var v91 uint16 = uint16(v5)
                                        _ = v91
                                    }
                                    var v92 float64 = 0.1
                                    if true {
                                        v93 := GenS8{F0: 128, F1: 100, F2: (v1 ^ (v2 % ((v1) | 1)))}
                                        _ = v93
                                    } else {
                                        if (0.1 > ((v6 * (v6 / v6)) / ((v6 - 1.5) * (0.1 + 1e10)))) {
                                            var v94 int64 = -9223372036854775808
                                            _ = v94
                                        }
                                        var v95 float64 = 1.5
                                        var v96 int32 = (v3 / ((v3) | 1))
                                        _ = v95
                                        _ = v96
                                    }
                                    _ = fi90
                                    _ = v92
                                }
                                v97 := GenS3{F0: 1, F1: 60}
                                _ = v97
                            }
                            _ = len(v19)
                        }
                        _ = fi69
                    }
                    switch ("a" + "hello") {
                    case "\n\t":
                        var v98 int64 = 1
                        _ = v98
                    case "":
                        v4 -= (v4 ^ v4)
                    case "hello":
                        v19 = append(v19, (v63 ^ v63))
                    default:
                        if (((v66 / ((255) | 1)) < (v66 >> ((1) & 63))) && false) {
                            var v99 uint8 = 160
                            if (!(true || true)) {
                                var v100 uint16 = uint16(v99)
                                if true {
                                    var v101 uint16 = v100
                                    var v102 uint8 = (v99 ^ (v99 | 127))
                                    _ = v101
                                    _ = v102
                                } else {
                                    for fi103 := 0; fi103 < 14; fi103++ {
                                        if false {
                                            v20 = 853
                                        }
                                        _ = fi103
                                    }
                                }
                                _ = v100
                            } else {
                                if false {
                                    v104 := GenS9{F0: v4, F1: v18, F2: 1}
                                    var v105 int = (v63 & v63)
                                    _ = v104
                                    _ = v105
                                }
                                for fi106 := 0; fi106 < 9; fi106++ {
                                    var v107 uint8 = v99
                                    if (3.14159 != 0.0) {
                                        fmt.Println(v99, v5, v66)
                                        var v108 uint = 635
                                        var v109 uint8 = v99
                                        _ = v108
                                        _ = v109
                                    }
                                    _ = fi106
                                    _ = v107
                                }
                                for fi110 := 0; fi110 < 16; fi110++ {
                                    var v111 int = -1
                                    v6 -= (float32(1.5) - ((0.1 + ((3.14159 / v6) * (((0.0 / v6) - (float32(-2.25) * (float32(1.5) / v6))) - (v6 - (1e10 + 1.5))))) - v6))
                                    _ = fi110
                                    _ = v111
                                }
                            }
                            _ = v99
                        }
                    }
                    var v112 float64 = 0.0
                    _ = v112
                }
            }
            _ = v66
        }
    } else {
        var v113 int32 = 0
        _ = v113
    }
    var v114 uint32 = (v4 << ((1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v114:", v114)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v114
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v63
}
