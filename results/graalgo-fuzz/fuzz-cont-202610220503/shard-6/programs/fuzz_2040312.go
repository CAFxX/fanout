package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 int8; F2 string }
func genFn0(p0 uint) uint64 {
    switch "\n\t" {
    case "":
        for fi6 := 0; fi6 < 5; fi6++ {
            fmt.Println(fi6)
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := GenS1{F0: 0}
                v9 := []int{(fi6 | (fi7 >> ((p0) & 63))), 42}
                if len(v9) > 0 {
                    v9[0] = fi7
                }
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
    case "a":
        p0 = p0
    case "foo bar":
        var v10 uint8 = 127
        _ = v10
    }
    p0 = 1
    p0 += p0
    _ = p0
    return 4294967295
}
func genFn1() int64 {
    var v11 string = "\n\t"
    v11 = (((v11 + v11) + v11) + ("hello" + "a"))
    switch 4294967295 {
    case 1:
        switch ("a" + "foo bar") {
        case "foo bar":
            v11 = "hello"
        default:
            fmt.Println(v11)
        }
    case 42:
        switch 757 {
        case 0:
            v11 = v11
        default:
            for fi12 := 0; fi12 < 8; fi12++ {
                v13 := []int{(fi12 | (fi12 % ((-441) | 1))), fi12, fi12, -9223372036854775808}
                var v14 int32 = 42
                v14 = (v14 & v14)
                _ = fi12
                _ = v13
                _ = v14
            }
        }
    default:
        v15 := GenS2{F0: 127}
        _ = v15
    }
    v11 = v11
    v16 := GenS3{F0: "", F1: -31, F2: 1}
    _ = v11
    _ = v16
    return 0
}

func main() {
    var v1 int16 = 366
    var v2 int = 9223372036854775807
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    var v5 int = (v2 % (((v2 & -995)) | 1))
    if ("" < ("" + "")) {
        for fi17 := 0; fi17 < 14; fi17++ {
            for fi18 := 0; fi18 < 17; fi18++ {
                switch (^(fi17)) {
                case 32768:
                    var v19 uint64 = 1
                    _ = v19
                case -9223372036854775808:
                    if false {
                        fmt.Println(fi17, fi18, v5)
                        v20 := v3[fi18]
                        var v21 uint = 65535
                        _ = v20
                        _ = v21
                    }
                case 4294967295:
                    v22 := GenS3{F0: "", F1: 0, F2: 65535}
                    _ = v22
                default:
                    if (0 > 1) {
                        v23 := v4[fi17]
                        _ = v23
                    }
                }
                _ = fi18
            }
            for fi24 := 0; fi24 < 7; fi24++ {
                if (255 == 330) {
                    var v25 float32 = (1e10 / ((((3.14159 * float32(1.5)) - ((0.1 + 0.1) / (-2.25 + 0.1))) * 3.14159) - -2.25))
                    v2 = (fi24 | fi24)
                    _ = v25
                }
                var v26 int = 127
                v27 := GenS1{F0: 42}
                _ = fi24
                _ = v26
                _ = v27
            }
            _ = fi17
        }
        fmt.Println(v1)
        var v28 int64 = 42
        _ = v28
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        v30 := GenS3{F0: "a", F1: -67, F2: 563}
        _ = fi29
        _ = v30
    }
    if true {
        switch "foo bar" {
        case "\n\t":
            v31 := GenS1{F0: 255}
            _ = v31
        case "foo bar":
            fmt.Println(v5, v1, v1)
        case "a":
            v2 -= v5
        default:
            switch ("\n\t" + "a") {
            case "\n\t":
                v2 -= v2
            }
        }
        var v32 uint8 = 1
        fmt.Println(v2, v2, v2)
        _ = v32
    } else {
        if false {
            fmt.Println(v1, v5, v5)
            var v33 uint = 293
            switch v5 {
            case 32767:
                for fi34 := 0; fi34 < 20; fi34++ {
                    v1 -= (v1 >> ((v33) & 63))
                    for fi35 := 0; fi35 < 13; fi35++ {
                        if (-192 != 42) {
                            switch (v2 % ((1) | 1)) {
                            case 42:
                                var v36 int32 = -2
                                _ = v36
                            default:
                                var v37 int64 = genFn1()
                                _ = v37
                            }
                            if false {
                                switch ("hello" + "foo bar") {
                                case "foo bar":
                                    fmt.Println(v2)
                                default:
                                    v38 := GenS4{F0: 18446744073709551615, F1: 1, F2: 65535}
                                    _ = v38
                                }
                            } else {
                                switch fi34 {
                                case -723:
                                    fmt.Println(fi34, v2, v33)
                                default:
                                    fmt.Println(v2)
                                    v39 := GenS1{F0: 255}
                                    _ = v39
                                }
                                v40 := v4[v5]
                                _ = v40
                            }
                            v41 := GenS4{F0: 18446744073709551615, F1: 42, F2: 0}
                            _ = v41
                        }
                        var v42 string = "\n\t"
                        _ = fi35
                        _ = v42
                    }
                    _ = fi34
                }
            }
            _ = v33
        } else {
            delete(v3, v5)
            v5 += v5
        }
        v43 := GenS4{F0: 42, F1: -445, F2: 255}
        _ = v43
    }
    switch "\n\t" {
    case "hello":
        v44 := GenS5{F0: 4294967295}
        _ = v44
    default:
        if true {
            if ((true || true) || ("hello" <= ("a" + "a"))) {
                v4[v5] = 989
            } else {
                var v45 float32 = 1.5
                switch "foo bar" {
                case "":
                    var v46 uint8 = 127
                    _ = v46
                case "foo bar":
                    switch "foo bar" {
                    case "":
                        v1 += (v1 - (v1 >> ((54) & 63)))
                    default:
                        if (false && true) {
                            var v47 int8 = 100
                            _ = v47
                        }
                    }
                }
                for fi48 := 0; fi48 < 19; fi48++ {
                    var v49 uint = 42
                    _ = fi48
                    _ = v49
                }
                _ = v45
            }
            for fi50 := 0; fi50 < 5; fi50++ {
                if (false && (127 < -128)) {
                    v5 = (fi50 & (v2 * int(v1)))
                    if false {
                        for fi51 := 0; fi51 < 15; fi51++ {
                            switch "" {
                            case "hello":
                                if true {
                                    var v52 string = ("hello" + ("a" + "foo bar"))
                                    var v53 uint64 = 0
                                    _ = v52
                                    _ = v53
                                } else {
                                    v54 := GenS6{F0: "hello"}
                                    _ = v54
                                }
                            case "foo bar":
                                if false {
                                    fmt.Println(v5)
                                    var v55 uint32 = 4294967295
                                    switch ("foo bar" + "") {
                                    case "":
                                        var v56 int16 = 54
                                        var v57 bool = false
                                        var v58 uint16 = 65535
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                    }
                                    v55 = (v55 - v55)
                                    _ = v55
                                }
                            case "a":
                                var v59 int16 = v1
                                _ = v59
                            default:
                                _ = len(v3)
                            }
                            fmt.Println(fi51)
                            _ = fi51
                        }
                        switch 32768 {
                        case 441:
                            for fi60 := 0; fi60 < 11; fi60++ {
                                v61 := GenS4{F0: 18446744073709551615, F1: -740, F2: 32767}
                                switch "" {
                                case "a":
                                    v2 = v5
                                }
                                _ = fi60
                                _ = v61
                            }
                        case 1:
                            switch "foo bar" {
                            case "hello":
                                var v62 float32 = 3.14159
                                _ = v62
                            default:
                                v3[(v5 / (((v2 << ((360) & 63))) | 1))] = ("hello" + "a")
                            }
                        default:
                            var v63 int16 = (v1 << ((2) & 63))
                            _ = v63
                        }
                    } else {
                        v64 := []int{}
                        if true {
                            var v65 float32 = ((float32(((((-2.25 + float64(3.14159)) + 1e10) + 3.14159) - float64((float32((0.0 + ((((0.1 / 1.0) - ((-2.25 - 1e10) / 1.0)) / (1e10 / 1.0)) + ((0.0 * 0.0) * 1.5)))) * (0.1 * -2.25))))) / ((1.5 / 1e10) * float32((((((0.1 + 0.0) * -2.25) - ((0.1 * 1.5) - 1e10)) + float64(1e10)) / -2.25)))) + ((3.14159 * 1e10) + -2.25))
                            _ = v65
                        } else {
                            v66 := GenS7{F0: 18446744073709551615}
                            v5 = v2
                            _ = v66
                        }
                        _ = v64
                    }
                    v67 := []int{}
                    _ = v67
                } else {
                    delete(v4, 1)
                }
                if true {
                    v2 = (v2 + (v5 / ((2147483647) | 1)))
                } else {
                    switch "a" {
                    case "":
                        v68 := v3[32767]
                        _ = v68
                    case "hello":
                        if (1 >= 111) {
                            switch "hello" {
                            case "a":
                                fmt.Println(v5)
                            case "\n\t":
                                v69 := []int{(v2 / (((v5 * 9223372036854775807)) | 1))}
                                _ = v69
                            }
                            v2 += fi50
                            if true {
                                fmt.Println(fi50)
                            }
                        } else {
                            if (!false) {
                                var v70 int32 = -2147483648
                                v71 := []int{(v2 & (fi50 ^ v5))}
                                _ = v70
                                _ = v71
                            }
                            var v72 int32 = 42
                            _ = v72
                        }
                    case "\n\t":
                        v5 += v2
                    default:
                        switch ("foo bar" + ("\n\t" + "hello")) {
                        case "a":
                            if (false && true) {
                                if (-1 > (fi50 ^ v2)) {
                                    var v73 int32 = 221
                                    v2 = (v5 - 32767)
                                    v5 += (fi50 / ((966) | 1))
                                    _ = v73
                                }
                            } else {
                                if false {
                                    var v74 int64 = -1
                                    _ = v74
                                }
                            }
                        case "hello":
                            var v75 uint64 = 4294967295
                            _ = v75
                        }
                    }
                    switch fi50 {
                    case 127:
                        for fi76 := 0; fi76 < 15; fi76++ {
                            v4[(v5 | 4294967295)] = 4294967295
                            var v77 int32 = 1
                            if false {
                                if true {
                                    var v78 int = (fi76 - 42)
                                    v3[int(v1)] = ""
                                    _ = v78
                                }
                                v79 := GenS8{F0: v5, F1: -108, F2: "a"}
                                _ = v79
                            }
                            _ = fi76
                            _ = v77
                        }
                    default:
                        switch 128 {
                        case -502:
                            delete(v3, fi50)
                        default:
                            delete(v3, fi50)
                        }
                    }
                    v1 -= v1
                }
                _ = fi50
            }
        } else {
            fmt.Println(v1, v5, v2)
            fmt.Println(v5, v5)
        }
    }
    delete(v3, (v2 << ((255) & 63)))
    var v80 uint64 = genFn0(255)
    if false {
        if (!true) {
            var v81 int16 = (v1 - v1)
            for fi82 := 0; fi82 < 20; fi82++ {
                var v83 bool = false
                switch fi82 {
                case 1:
                    v2 -= v2
                case 32768:
                    fmt.Println(v81, v1)
                default:
                    fmt.Println(fi82, v2, v81)
                }
                _ = fi82
                _ = v83
            }
            _ = v81
        }
        v1 += v1
    }
    if (!true) {
        switch 127 {
        case -35:
            var v84 bool = true
            _ = v84
        case -946:
            delete(v4, (v2 >> ((145) & 63)))
        default:
            var v85 float64 = 1.5
            _ = v85
        }
    } else {
        for fi86 := 0; fi86 < 13; fi86++ {
            var v87 int16 = 0
            var v88 float64 = (((3.14159 / (1.5 + 0.1)) * (0.0 / (0.1 + 3.14159))) - 0.1)
            _ = fi86
            _ = v87
            _ = v88
        }
    }
    switch ((("" + "foo bar") + ((("" + "a") + "") + "hello")) + ("hello" + ("\n\t" + ("foo bar" + "a")))) {
    case "foo bar":
        for fi89 := 0; fi89 < 16; fi89++ {
            var v90 uint = 4294967295
            for fi91 := 0; fi91 < 14; fi91++ {
                if false {
                    if (480 != 2) {
                        fmt.Println(v1, fi91)
                        var v92 int8 = 1
                        v93 := GenS1{F0: v90}
                        _ = v92
                        _ = v93
                    } else {
                        fmt.Println(v80, v2, v2)
                    }
                }
                var v94 string = "\n\t"
                v1 = v1
                _ = fi91
                _ = v94
            }
            for fi95 := 0; fi95 < 15; fi95++ {
                if true {
                    v96 := GenS6{F0: ("hello" + "")}
                    fmt.Println(v2, v2, fi89)
                    v97 := []int{}
                    _ = v96
                    _ = v97
                }
                v98 := GenS6{F0: ("" + ("a" + "a"))}
                fmt.Println(v80, v80)
                if false {
                    v99 := []int{(v5 + (fi95 % (((fi95 ^ 1)) | 1))), v5, fi95, (fi89 + (fi95 + fi89))}
                    fmt.Println(fi89, v5)
                    _ = v99
                }
                _ = fi95
                _ = v98
            }
            _ = fi89
            _ = v90
        }
    default:
        v100 := []int{(^(v5)), (v2 ^ (v2 - v2)), 1, 42}
        _ = v100
    }
    if (!(-93 == 1)) {
        for fi101 := 0; fi101 < 17; fi101++ {
            if (true || (!((9223372036854775807 > 889) && (genFn0(255) <= v80)))) {
                fmt.Println(v1, v1)
                for fi102 := 0; fi102 < 18; fi102++ {
                    var v103 int64 = 0
                    if (true || false) {
                        fmt.Println(v2)
                        if true {
                            v104 := GenS6{F0: "hello"}
                            _ = v104
                        } else {
                            var v105 string = "foo bar"
                            var v106 int64 = (v103 >> ((65535) & 63))
                            fmt.Println(v80, v80, v106)
                            var v107 float32 = 1e10
                            _ = v105
                            _ = v106
                            _ = v107
                        }
                        v5 = 128
                    }
                    _ = fi102
                    _ = v103
                }
            } else {
                var v108 int64 = 42
                _ = v108
            }
            _ = fi101
        }
        v80 = (v80 + v80)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v80
}
