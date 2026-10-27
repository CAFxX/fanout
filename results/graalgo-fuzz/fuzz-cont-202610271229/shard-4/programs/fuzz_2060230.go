package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int64; F1 int32; F2 int }
type GenS5 struct { F0 bool; F1 bool; F2 int16 }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 uint64; F1 int16 }
func genFn0(p0 string, p1 uint64, p2 uint32) int {
    var v6 uint32 = p2
    for fi7 := 0; fi7 < 4; fi7++ {
        v6 = v6
        _ = fi7
    }
    if false {
        if false {
            p2 = v6
        }
        for fi8 := 0; fi8 < 3; fi8++ {
            if false {
                var v9 float32 = float32(3.14159)
                _ = v9
            } else {
                switch fi8 {
                case 127:
                    if (true && (true || ((0.1 - (0.1 - ((0.0 - float32(3.14159)) - 1.5))) == (0.0 - float32(1e10))))) {
                        var v10 uint8 = 38
                        _ = v10
                    }
                case 4294967295:
                    var v11 int = 32767
                    _ = v11
                default:
                    var v12 uint = 0
                    _ = v12
                }
                var v13 uint = 0
                v14 := GenS1{F0: (v13 - (v13 | (v13 | v13))), F1: v13, F2: v13}
                _ = v13
                _ = v14
            }
            v6 -= (p2 + (p2 | v6))
            _ = fi8
        }
        if (!true) {
            var v15 float64 = 0.1
            v16 := []int{32767, -1}
            _ = v15
            _ = v16
        }
    }
    v17 := GenS2{F0: 127, F1: p2}
    var v18 float32 = (-2.25 - -2.25)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v6
    return -1
}
func genFn1(p0 int8, p1 uint) int16 {
    var v19 bool = ((p0 - p0) > 100)
    for fi20 := 0; fi20 < 15; fi20++ {
        if true {
            v21 := GenS2{F0: 30, F1: 434}
            if (-1 < 0) {
                v22 := GenS2{F0: p0, F1: 65535}
                var v23 uint8 = 255
                var v24 int32 = -2147483648
                _ = v22
                _ = v23
                _ = v24
            }
            for fi25 := 0; fi25 < 14; fi25++ {
                if (!v19) {
                    if ((v19 || false) || v19) {
                        fmt.Println(v19, p1, p0)
                    }
                    if v19 {
                        v26 := GenS2{F0: p0, F1: 42}
                        var v27 int32 = -396
                        var v28 uint64 = 4294967295
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                }
                _ = fi25
            }
            _ = v21
        }
        _ = fi20
    }
    if (4294967295 >= 42) {
        switch ("hello" + "") {
        case "a":
            switch "\n\t" {
            case "a":
                if false {
                    for fi29 := 0; fi29 < 17; fi29++ {
                        for fi30 := 0; fi30 < 7; fi30++ {
                            v31 := GenS3{F0: (fi30 >> (((^(p1))) & 63)), F1: 42, F2: 4294967295}
                            _ = fi30
                            _ = v31
                        }
                        v32 := GenS1{F0: p1, F1: (^(p1)), F2: p1}
                        _ = fi29
                        _ = v32
                    }
                }
            case "foo bar":
                switch "hello" {
                case "hello":
                    v19 = true
                case "":
                    v19 = (42 != 2)
                case "foo bar":
                    p0 -= p0
                }
            case "\n\t":
                v33 := []int{2147483647, 33, genFn0((("hello" + "foo bar") + ""), 2, 4294967295), 128}
                _ = v33
            }
        case "foo bar":
            if (!(!(p1 != (p1 / (((^(p1))) | 1))))) {
                fmt.Println(p1, p0, p0)
                v19 = true
            } else {
                for fi34 := 0; fi34 < 16; fi34++ {
                    v35 := GenS3{F0: genFn0(("foo bar" + "\n\t"), 246, 65535), F1: 513, F2: 42}
                    var v36 float64 = 0.1
                    if v19 {
                        var v37 uint16 = 116
                        _ = v37
                    } else {
                        switch ("foo bar" + ("hello" + "a")) {
                        case "hello":
                            var v38 int16 = 32767
                            _ = v38
                        case "a":
                            for fi39 := 0; fi39 < 3; fi39++ {
                                v36 = -2.25
                                _ = fi39
                            }
                        default:
                            v40 := []int{fi34}
                            _ = v40
                        }
                        switch genFn0("hello", 2, 0) {
                        case 1:
                            var v41 uint16 = 0
                            _ = v41
                        case 32767:
                            if (!false) {
                                v19 = v19
                                switch "\n\t" {
                                case "a":
                                    fmt.Println(p0, p0)
                                case "hello":
                                    v36 = 0.1
                                default:
                                    var v42 int16 = 32767
                                    _ = v42
                                }
                            }
                        }
                    }
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
        }
    } else {
        for fi43 := 0; fi43 < 12; fi43++ {
            if true {
                var v44 uint8 = uint8(p1)
                v45 := []int{32767, fi43, (-(fi43))}
                _ = v44
                _ = v45
            }
            v46 := GenS2{F0: p0, F1: 0}
            for fi47 := 0; fi47 < 1; fi47++ {
                switch "" {
                case "":
                    fmt.Println(v19)
                case "a":
                    switch (fi47 & genFn0("a", 1, 2147483647)) {
                    case -999:
                        switch "" {
                        case "\n\t":
                            v48 := GenS1{F0: (p1 & p1), F1: p1, F2: p1}
                            _ = v48
                        }
                        v49 := []int{}
                        _ = v49
                    }
                case "foo bar":
                    fmt.Println(fi43, p0, p0)
                default:
                    p1 -= (^(p1))
                }
                _ = fi47
            }
            _ = fi43
            _ = v46
        }
        v19 = v19
        var v50 uint = (^(p1))
        _ = v50
    }
    _ = p0
    _ = p1
    _ = v19
    return -389
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"foo bar"}
    var v3 float64 = -2.25
    var v4 float32 = 1e10
    var v5 uint16 = 32767
    for fi51 := 0; fi51 < 15; fi51++ {
        fmt.Println(v3, v4)
        if (166 >= 0) {
            if true {
                if ((genFn0("foo bar", 18446744073709551615, 322) >= genFn0("hello", 4294967295, 65535)) || (false || (("" + (("foo bar" + ((("" + "") + "hello") + "hello")) + "a")) == "hello"))) {
                    if (468 <= 337) {
                        if false {
                            var v52 int = 9223372036854775807
                            for fi53 := 0; fi53 < 19; fi53++ {
                                v54 := GenS4{F0: 42, F1: -943, F2: fi51}
                                v3 -= 1e10
                                _ = fi53
                                _ = v54
                            }
                            var v55 bool = ((^(v52)) != (v52 + (fi51 | (v52 & (fi51 ^ (fi51 ^ (fi51 << ((42) & 63))))))))
                            _ = v52
                            _ = v55
                        } else {
                            var v56 int = -2147483648
                            v2 = append(v2, "")
                            for fi57 := 0; fi57 < 18; fi57++ {
                                v58 := GenS3{F0: -2147483648, F1: 18446744073709551615, F2: 2}
                                var v59 int8 = 100
                                _ = fi57
                                _ = v58
                                _ = v59
                            }
                            _ = v56
                        }
                    }
                    var v60 int64 = 811
                    for fi61 := 0; fi61 < 6; fi61++ {
                        v4 = v4
                        _ = fi61
                    }
                    _ = v60
                }
                var v62 bool = (("" > ("foo bar" + (("" + "\n\t") + (("\n\t" + "a") + "foo bar")))) && false)
                _ = v62
            } else {
                var v63 int = (fi51 + fi51)
                _ = v63
            }
            for fi64 := 0; fi64 < 19; fi64++ {
                _ = len(v1)
                if false {
                    var v65 uint16 = 897
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v65
                }
                switch -1 {
                case 32768:
                    v66 := v1[234]
                    _ = v66
                case 9223372036854775807:
                    var v67 uint32 = 0
                    _ = v67
                default:
                    var v68 float32 = v4
                    _ = v68
                }
                _ = fi64
            }
        } else {
            if false {
                fmt.Println(fi51, v3)
                switch "a" {
                case "\n\t":
                    switch "" {
                    case "\n\t":
                        for fi69 := 0; fi69 < 4; fi69++ {
                            if (2 >= 139) {
                                if ((false || false) || true) {
                                    var v70 int = -1
                                    _ = v70
                                } else {
                                    v71 := GenS5{F0: false, F1: (524 <= 969), F2: 42}
                                    v3 += float64(-2.25)
                                    v72 := GenS3{F0: genFn0(("a" + ("foo bar" + "hello")), 40, 669), F1: 4294967295, F2: 494}
                                    _ = v71
                                    _ = v72
                                }
                                v73 := GenS3{F0: (fi69 - fi51), F1: 42, F2: 2}
                                if ((("" + "a") > ("hello" + "hello")) || (757 == 2)) {
                                    var v74 float32 = v4
                                    _ = v74
                                }
                                _ = v73
                            }
                            switch (^(fi51)) {
                            case 0:
                                v3 += float64(((v4 + (0.0 / (v4 * v4))) - v4))
                            default:
                                var v75 int8 = 1
                                _ = v75
                            }
                            var v76 int32 = 42
                            _ = fi69
                            _ = v76
                        }
                    }
                case "a":
                    for fi77 := 0; fi77 < 4; fi77++ {
                        v3 += (((3.14159 - v3) / v3) * (v3 + v3))
                        delete(v1, 42)
                        _ = fi77
                    }
                case "hello":
                    if false {
                        _ = len(v2)
                        fmt.Println(v5, v4, fi51)
                        var v78 float64 = v3
                        _ = v78
                    } else {
                        v79 := v1[2147483648]
                        var v80 uint32 = 394
                        v79 = (v79 + "foo bar")
                        _ = v79
                        _ = v80
                    }
                }
            }
            v3 = v3
            for fi81 := 0; fi81 < 1; fi81++ {
                v5 += (v5 << ((373) & 63))
                v3 = v3
                _ = fi81
            }
        }
        var v82 float32 = v4
        _ = fi51
        _ = v82
    }
    switch -652 {
    case -634:
        if (484 > 1) {
            v5 -= v5
        }
    case -2147483648:
        var v83 uint16 = 42
        _ = v83
    }
    var v84 uint16 = (v5 % (((v5 & (v5 - 255))) | 1))
    if (0 == -839) {
        v84 += v5
    }
    if (!false) {
        var v85 float32 = (v4 + (((v4 + (v4 + 0.1)) + (v4 + 0.0)) * v4))
        v86 := GenS1{F0: 2, F1: 4294967295, F2: 967}
        v4 += 1.5
        _ = v85
        _ = v86
    }
    var v87 int32 = 363
    if false {
        v88 := v1[42]
        _ = v88
    }
    fmt.Println(v87)
    for fi89 := 0; fi89 < 11; fi89++ {
        switch "a" {
        case "a":
            v4 = v4
        case "\n\t":
            if false {
                v3 -= (-2.25 * float64(float32((((v3 - (((v3 + 3.14159) / v3) + v3)) - (0.1 - 1.5)) + v3))))
                fmt.Println(v3, v84, fi89)
                var v90 float32 = float32((((v3 + (v3 / float64(v4))) / v3) + (-2.25 - v3)))
                _ = v90
            }
        }
        _ = fi89
    }
    for fi91 := 0; fi91 < 4; fi91++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v92 uint8 = 1
        _ = fi91
        _ = v92
    }
    if false {
        v5 = (v84 ^ v84)
        switch -9223372036854775808 {
        case 535:
            v93 := GenS6{F0: 9223372036854775807, F1: ("" + "\n\t")}
            _ = v93
        case -600:
            v4 -= -2.25
        default:
            v94 := GenS7{F0: v87, F1: (v84 * v5)}
            _ = v94
        }
    } else {
        for fi95 := 0; fi95 < 18; fi95++ {
            var v96 float64 = v3
            _ = fi95
            _ = v96
        }
        switch "" {
        case "a":
            var v97 uint64 = 1
            _ = v97
        case "\n\t":
            fmt.Println(v84, v84)
        }
        switch -2147483648 {
        case -1:
            var v98 float32 = v4
            _ = v98
        case -674:
            for fi99 := 0; fi99 < 16; fi99++ {
                var v100 uint64 = 776
                if false {
                    for fi101 := 0; fi101 < 5; fi101++ {
                        v3 = v3
                        v5 = v5
                        var v102 uint64 = 0
                        _ = fi101
                        _ = v102
                    }
                } else {
                    switch ("a" + "\n\t") {
                    case "hello":
                        v100 = (v100 >> ((665) & 63))
                    case "\n\t":
                        if false {
                            for fi103 := 0; fi103 < 15; fi103++ {
                                var v104 uint16 = v84
                                v84 = 905
                                _ = fi103
                                _ = v104
                            }
                        }
                    case "foo bar":
                        if (!(v100 == (v100 / ((uint64(v84)) | 1)))) {
                            v84 += (v84 % (((v84 / (((v5 | v84)) | 1))) | 1))
                        } else {
                            v100 += v100
                            var v105 int = fi99
                            fmt.Println(v100, v100, v5)
                            _ = v105
                        }
                    }
                    for fi106 := 0; fi106 < 6; fi106++ {
                        fmt.Println(v100, v100, fi106)
                        v107 := GenS6{F0: 9223372036854775807, F1: "foo bar"}
                        var v108 int32 = (v87 >> ((42) & 63))
                        _ = fi106
                        _ = v107
                        _ = v108
                    }
                }
                if (!((!(!false)) || (153 <= (fi99 / (((fi99 % (((fi99 << ((416) & 63))) | 1))) | 1))))) {
                    v109 := GenS1{F0: 1, F1: 278, F2: 596}
                    _ = v109
                }
                _ = fi99
                _ = v100
            }
        default:
            switch 0 {
            case 847:
                v110 := GenS8{F0: 4294967295, F1: 32767}
                _ = v110
            case -669:
                switch 2147483647 {
                case 0:
                    v87 = v87
                }
            default:
                for fi111 := 0; fi111 < 10; fi111++ {
                    var v112 bool = false
                    _ = fi111
                    _ = v112
                }
            }
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v84:", v84)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v84
    _ = v87
}
