package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int64 }
type GenS2 struct { F0 int; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 bool; F1 bool }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int16; F1 int32; F2 int }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 bool; F1 uint8; F2 int32 }
type GenS12 struct { F0 uint; F1 int; F2 int64 }
func genFn0(p0 string) string {
    if true {
        var v4 int64 = -848
        if true {
            v4 -= (-(v4))
        }
        _ = v4
    }
    v5 := []int{127, -574, 42, -2147483648}
    if true {
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := GenS1{F0: 42, F1: 1000, F2: 719}
            if false {
                if (true && (!false)) {
                    v8 := GenS2{F0: fi6, F1: 2, F2: 255}
                    if (0 == 2) {
                        for fi9 := 0; fi9 < 13; fi9++ {
                            v5 = append(v5, (fi9 << ((255) & 63)))
                            _ = fi9
                        }
                        v5 = append(v5, (fi6 * fi6))
                        var v10 float32 = 3.14159
                        _ = v10
                    }
                    p0 = "\n\t"
                    _ = v8
                } else {
                    p0 = p0
                }
                fmt.Println(fi6, fi6)
                var v11 float32 = (((float32(0.1) / 1.5) / 3.14159) - (0.0 - 3.14159))
                _ = v11
            }
            switch -433 {
            case 127:
                v12 := GenS1{F0: 2, F1: -837, F2: 1}
                _ = v12
            default:
                p0 = p0
            }
            _ = fi6
            _ = v7
        }
    } else {
        if true {
            p0 = "foo bar"
            v5 = append(v5, -762)
        }
        var v13 uint64 = 18446744073709551615
        _ = v13
    }
    _ = p0
    _ = v5
    return p0
}
func genFn1(p0 uint8) uint {
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(fi14, p0)
        p0 += p0
        p0 -= p0
        _ = fi14
    }
    v15 := []int{4294967295}
    for fi16 := 0; fi16 < 1; fi16++ {
        fmt.Println(p0)
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := GenS3{F0: 0}
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    _ = p0
    _ = v15
    return 18446744073709551615
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]bool)
    var v3 uint8 = 51
    switch v1 {
    case "\n\t":
        if (42 == 576) {
            v19 := []int{9223372036854775807, 9223372036854775807, 32768}
            var v20 float64 = float64(1e10)
            _ = v19
            _ = v20
        } else {
            v1 = v1
            var v21 int32 = 173
            v1 = (((v1 + v1) + ("a" + (v1 + v1))) + ("" + (("hello" + v1) + "a")))
            _ = v21
        }
    }
    if true {
        switch "" {
        case "a":
            v22 := []int{-716, -920}
            _ = v22
        default:
            var v23 uint = genFn1(0)
            _ = v23
        }
        v24 := []int{2147483647}
        _ = v24
    } else {
        if false {
            switch 0 {
            case 42:
                for fi25 := 0; fi25 < 7; fi25++ {
                    if true {
                        v3 += v3
                        v26 := GenS1{F0: v3, F1: 42, F2: 0}
                        if ((!true) && (854 == 2147483648)) {
                            var v27 int64 = -9223372036854775808
                            _ = v27
                        } else {
                            v1 = "hello"
                            delete(v2, 1)
                            if (true && (!(1e10 != (((-2.25 / 0.1) * (float64(-2.25) * (-2.25 + (((((3.14159 / 1.5) - float64(((float32(float64(0.0)) / ((1.5 / (1.5 + 1.5)) / (3.14159 + 1e10))) * (float32(float64(-2.25)) * (((((3.14159 - ((0.0 - ((1e10 - (-2.25 - (float32(1.5) / (0.1 + 0.1)))) + -2.25)) - 0.0)) + 3.14159) * float32(1.5)) + (float32((3.14159 * 0.0)) * 1e10)) * 1.5))))) - float64(3.14159)) - 0.1) + 1e10)))) * (float64((float32(1e10) / -2.25)) * 1e10))))) {
                                v1 = v1
                                switch v1 {
                                case "\n\t":
                                    v3 -= 127
                                case "foo bar":
                                    v28 := GenS2{F0: fi25, F1: 0, F2: v3}
                                    _ = v28
                                case "a":
                                    switch "a" {
                                    case "hello":
                                        if (!false) {
                                            switch v1 {
                                            case "":
                                                v1 = ((((v1 + ((v1 + "hello") + (v1 + (("\n\t" + v1) + "a")))) + v1) + v1) + (v1 + v1))
                                            case "a":
                                                v29 := GenS4{F0: 724, F1: 670, F2: 42}
                                                _ = v29
                                            }
                                            v1 = (v1 + "\n\t")
                                        } else {
                                            var v30 uint32 = 259
                                            _ = v30
                                        }
                                    default:
                                        if true {
                                            v31 := GenS5{F0: 42, F1: 1}
                                            _ = v31
                                        } else {
                                            delete(v2, 0)
                                            for fi32 := 0; fi32 < 8; fi32++ {
                                                switch (fi25 << ((613) & 63)) {
                                                case 0:
                                                    if true {
                                                        _ = len(v2)
                                                    } else {
                                                        v33 := GenS4{F0: 1, F1: 2, F2: 902}
                                                        var v34 uint64 = 389
                                                        switch (v1 + v1) {
                                                        case "":
                                                            v35 := GenS6{F0: "a", F1: v3}
                                                            _ = v35
                                                        case "foo bar":
                                                            v2[166] = (false || false)
                                                        }
                                                        if ((v3 & (v3 >> ((255) & 63))) != (v3 * v3)) {
                                                            if (-1 > 1) {
                                                                fmt.Println(fi25, fi25, fi32)
                                                            } else {
                                                                switch "a" {
                                                                case "":
                                                                    v36 := GenS7{F0: true, F1: true}
                                                                    _ = v36
                                                                case "foo bar":
                                                                    var v37 uint = 0
                                                                    _ = v37
                                                                case "\n\t":
                                                                    var v38 uint = 255
                                                                    _ = v38
                                                                }
                                                            }
                                                        }
                                                        var v39 int = -2147483648
                                                        var v40 uint32 = 42
                                                        _ = v33
                                                        _ = v34
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                }
                                                switch (fi25 * (fi32 + -2147483648)) {
                                                case -18:
                                                    v3 = (v3 >> ((255) & 63))
                                                case -9223372036854775808:
                                                    if (-9223372036854775808 <= 0) {
                                                        fmt.Println(fi25)
                                                        var v41 string = ("" + "foo bar")
                                                        var v42 uint = 18446744073709551615
                                                        _ = v41
                                                        _ = v42
                                                    } else {
                                                        v43 := v2[878]
                                                        if (!v43) {
                                                            v44 := GenS8{F0: 30}
                                                            if true {
                                                                fmt.Println(v43, fi25)
                                                                var v45 int8 = 1
                                                                var v46 int = -9223372036854775808
                                                                _ = v45
                                                                _ = v46
                                                            }
                                                            if v43 {
                                                                if ((1e10 + float32(0.1)) == 1.5) {
                                                                    fmt.Println(v1)
                                                                    var v47 float64 = 1.5
                                                                    v1 = "hello"
                                                                    var v48 int = 128
                                                                    v47 += ((((-2.25 - v47) + (0.0 - float64(((-2.25 - 0.1) * (float32(v47) - 1e10))))) * -2.25) + float64((((3.14159 - -2.25) - (3.14159 / float32(1.5))) - 0.0)))
                                                                    _ = v47
                                                                    _ = v48
                                                                }
                                                                var v49 int16 = -1
                                                                _ = v49
                                                            } else {
                                                                var v50 int64 = 0
                                                                _ = v50
                                                            }
                                                            var v51 float32 = 0.1
                                                            _ = v44
                                                            _ = v51
                                                        } else {
                                                            var v52 float64 = float64(-2.25)
                                                            _ = v52
                                                        }
                                                        _ = v43
                                                    }
                                                case 32767:
                                                    v1 = ("foo bar" + (("\n\t" + "") + "\n\t"))
                                                default:
                                                    v53 := []int{}
                                                    _ = v53
                                                }
                                                fmt.Println(fi25, fi32)
                                                _ = fi32
                                            }
                                        }
                                    }
                                default:
                                    v3 = (v3 % ((42) | 1))
                                }
                            }
                        }
                        _ = v26
                    }
                    fmt.Println(v1, v3)
                    _ = fi25
                }
            case -691:
                for fi54 := 0; fi54 < 1; fi54++ {
                    var v55 float32 = 1e10
                    _ = fi54
                    _ = v55
                }
            default:
                fmt.Println(v1)
            }
            var v56 float32 = 3.14159
            _ = v56
        } else {
            fmt.Println(v3, v3)
            if false {
                for fi57 := 0; fi57 < 16; fi57++ {
                    switch (v1 + "") {
                    case "\n\t":
                        v1 = v1
                    case "":
                        var v58 int16 = 1
                        _ = v58
                    }
                    var v59 string = ("\n\t" + v1)
                    var v60 string = "foo bar"
                    _ = fi57
                    _ = v59
                    _ = v60
                }
                v1 = "foo bar"
                v61 := []int{9223372036854775807, 127}
                _ = v61
            }
        }
        if true {
            fmt.Println(v3, v3)
        } else {
            if true {
                switch "\n\t" {
                case "":
                    var v62 uint32 = 4294967295
                    _ = v62
                default:
                    for fi63 := 0; fi63 < 13; fi63++ {
                        var v64 uint = 65535
                        var v65 uint = (v64 | 415)
                        if (!(!true)) {
                            v65 += 0
                            var v66 bool = true
                            for fi67 := 0; fi67 < 18; fi67++ {
                                switch 128 {
                                case 9223372036854775807:
                                    var v68 uint = (v65 * (^(v64)))
                                    _ = v68
                                default:
                                    var v69 uint16 = 433
                                    _ = v69
                                }
                                fmt.Println(v1, fi67, v65)
                                _ = len(v2)
                                v3 -= (v3 >> ((v64) & 63))
                                _ = fi67
                            }
                            _ = v66
                        }
                        _ = fi63
                        _ = v64
                        _ = v65
                    }
                }
                _ = len(v2)
                for fi70 := 0; fi70 < 5; fi70++ {
                    var v71 int32 = 42
                    v1 = v1
                    _ = fi70
                    _ = v71
                }
            }
            for fi72 := 0; fi72 < 19; fi72++ {
                v2[4294967295] = false
                v3 -= v3
                v73 := []int{(fi72 & fi72), (fi72 + fi72)}
                _ = fi72
                _ = v73
            }
            v74 := GenS9{F0: 1000, F1: 206, F2: 127}
            _ = v74
        }
        var v75 int32 = 2147483647
        _ = v75
    }
    fmt.Println(v3, v3)
    var v76 float32 = 1e10
    if (true || false) {
        v3 -= (v3 - (v3 / (((v3 - (v3 << ((42) & 63)))) | 1)))
        var v77 int64 = 363
        _ = v77
    }
    _ = len(v2)
    var v78 float32 = (v76 / (-2.25 - (v76 - v76)))
    v1 = v1
    for fi79 := 0; fi79 < 16; fi79++ {
        var v80 uint16 = 65535
        _ = fi79
        _ = v80
    }
    for fi81 := 0; fi81 < 2; fi81++ {
        for fi82 := 0; fi82 < 8; fi82++ {
            var v83 uint16 = 175
            if (true && false) {
                v78 += float32(float64(v78))
                var v84 uint = 18446744073709551615
                _ = len(v2)
                _ = v84
            }
            fmt.Println(v3)
            var v85 uint32 = 883
            var v86 uint64 = 18446744073709551615
            _ = fi82
            _ = v83
            _ = v85
            _ = v86
        }
        var v87 uint8 = (v3 ^ 71)
        v78 = v76
        _ = fi81
        _ = v87
    }
    if ((float64((((3.14159 * (float32(1.5) + v78)) - v76) / v78)) * 0.0) > 1.5) {
        v78 += (-2.25 + (((0.0 - v78) / 1e10) * (v78 * v76)))
        if true {
            fmt.Println(v76, v1, v1)
        }
        v76 = v78
    }
    if (777 != 42) {
        v88 := GenS10{F0: 1}
        if false {
            if true {
                for fi89 := 0; fi89 < 5; fi89++ {
                    var v90 int8 = 127
                    v90 -= v90
                    _ = fi89
                    _ = v90
                }
                var v91 uint16 = 32767
                _ = v91
            } else {
                v92 := GenS11{F0: (4294967295 < 0), F1: v3, F2: -1}
                _ = v92
            }
            for fi93 := 0; fi93 < 19; fi93++ {
                var v94 float32 = (((-2.25 + 1.5) + ((0.0 - v76) + v76)) + (((v78 * v78) - v76) * 1e10))
                _ = fi93
                _ = v94
            }
        }
        fmt.Println(v3, v78, v1)
        _ = v88
    } else {
        v95 := GenS12{F0: 82, F1: 1, F2: 42}
        _ = v95
    }
    fmt.Println(v76, v76)
    if (!false) {
        v76 -= v76
    } else {
        switch v1 {
        case "hello":
            switch "hello" {
            case "":
                switch -1 {
                case 139:
                    for fi96 := 0; fi96 < 2; fi96++ {
                        v76 += v76
                        if true {
                            _ = len(v2)
                            v97 := GenS4{F0: 1000, F1: 18446744073709551615, F2: -799}
                            _ = v97
                        }
                        _ = fi96
                    }
                default:
                    if false {
                        v76 -= float32(0.1)
                        v1 = v1
                    } else {
                        var v98 uint8 = v3
                        _ = v98
                    }
                }
            default:
                v99 := []int{4294967295, -9223372036854775808, -381}
                _ = v99
            }
        }
    }
    v100 := v2[65535]
    for fi101 := 0; fi101 < 11; fi101++ {
        var v102 float32 = 0.0
        if true {
            v76 = v102
        } else {
            var v103 uint16 = 2
            _ = v103
        }
        _ = fi101
        _ = v102
    }
    if false {
        v3 -= v3
    } else {
        var v104 float64 = 1e10
        _ = v104
    }
    if false {
        var v105 int64 = 9223372036854775807
        var v106 int = -626
        _ = v105
        _ = v106
    }
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v76
    _ = v78
}
