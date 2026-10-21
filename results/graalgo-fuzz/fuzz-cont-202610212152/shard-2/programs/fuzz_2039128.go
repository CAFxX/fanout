package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 string; F2 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int; F2 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint; F1 bool; F2 int32 }
type GenS9 struct { F0 int; F1 uint8 }
type GenS10 struct { F0 int64; F1 int; F2 uint8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int64; F1 int8; F2 uint16 }
type GenS13 struct { F0 uint; F1 uint8; F2 uint16 }

func main() {
    var v1 float32 = 3.14159
    var v2 int32 = -1
    v3 := make(map[int]uint32)
    _ = len(v3)
    switch "\n\t" {
    case "hello":
        var v4 float32 = v1
        _ = v4
    default:
        v2 = (v2 | v2)
    }
    if (true || true) {
        var v5 uint8 = 2
        switch "" {
        case "":
            var v6 bool = true
            _ = v6
        case "a":
            v7 := GenS1{F0: 9223372036854775807}
            _ = v7
        default:
            switch 32768 {
            case -2147483648:
                var v8 string = "foo bar"
                _ = v8
            case 32768:
                if (!false) {
                    for fi9 := 0; fi9 < 4; fi9++ {
                        v5 = (^(v5))
                        _ = fi9
                    }
                    fmt.Println(v2, v5)
                } else {
                    if ((false || ((("foo bar" >= "") && false) || true)) && true) {
                        v10 := GenS1{F0: 128}
                        for fi11 := 0; fi11 < 7; fi11++ {
                            var v12 int64 = 0
                            if (v1 != (float32(((-2.25 + ((1e10 / 0.1) * ((3.14159 / -2.25) + 1.5))) - 0.0)) / (((v1 - 1.5) - 1.5) - v1))) {
                                fmt.Println(v5, fi11, v2)
                                v13 := []int{fi11, 128}
                                _ = v13
                            }
                            _ = fi11
                            _ = v12
                        }
                        _ = v10
                    }
                    v5 = v5
                    switch "hello" {
                    case "a":
                        if (-32768 != 32767) {
                            v14 := []int{879, 1, 0, 42}
                            _ = v14
                        } else {
                            fmt.Println(v5, v1)
                        }
                    default:
                        v15 := v3[32767]
                        _ = v15
                    }
                }
            default:
                var v16 int32 = -1
                _ = v16
            }
        }
        var v17 uint16 = 537
        _ = v17
        _ = v5
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 int = (fi18 << ((65535) & 63))
        delete(v3, fi18)
        for fi20 := 0; fi20 < 5; fi20++ {
            v2 += v2
            switch (("foo bar" + "") + "\n\t") {
            case "foo bar":
                v21 := []int{fi18, fi18, fi18, (fi18 % ((32768) | 1))}
                _ = v21
            default:
                var v22 uint16 = 743
                _ = v22
            }
            _ = len(v3)
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        v1 = (v1 * v1)
        v24 := GenS1{F0: fi23}
        v25 := GenS1{F0: (-(fi23))}
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 uint64 = 961
        v1 = ((1.5 - (v1 + ((v1 * v1) / float32(3.14159)))) / (float32(float64(((v1 + (float32(0.0) - float32(1e10))) * v1))) - float32((3.14159 - (-2.25 + (float64(v1) / (1.5 * (1e10 + float64(v1)))))))))
        var v28 int8 = -1
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v1, v1)
    switch ("hello" + "") {
    case "":
        for fi29 := 0; fi29 < 18; fi29++ {
            v30 := GenS1{F0: 218}
            v31 := []int{fi29}
            _ = fi29
            _ = v30
            _ = v31
        }
    case "a":
        for fi32 := 0; fi32 < 4; fi32++ {
            var v33 uint32 = 319
            v34 := v3[-651]
            delete(v3, -1)
            _ = fi32
            _ = v33
            _ = v34
        }
    default:
        v35 := GenS1{F0: 2147483647}
        _ = v35
    }
    if true {
        switch 9223372036854775807 {
        case 42:
            if false {
                v36 := GenS2{F0: 9223372036854775807}
                var v37 uint8 = 160
                var v38 string = "a"
                _ = v36
                _ = v37
                _ = v38
            }
        case -1:
            if true {
                v2 = v2
                fmt.Println(v1, v1)
            } else {
                if (0 == 1) {
                    if true {
                        v39 := GenS1{F0: 1}
                        for fi40 := 0; fi40 < 8; fi40++ {
                            v1 = v1
                            _ = fi40
                        }
                        _ = v39
                    } else {
                        v41 := GenS2{F0: 32768}
                        fmt.Println(v2)
                        if (!(((^(v2)) != 42) || (42 > 2))) {
                            delete(v3, 4294967295)
                            for fi42 := 0; fi42 < 14; fi42++ {
                                fmt.Println(fi42, fi42, v2)
                                _ = fi42
                            }
                        }
                        _ = v41
                    }
                } else {
                    v1 -= ((v1 * 0.1) + v1)
                    var v43 uint32 = 0
                    _ = v43
                }
            }
        default:
            var v44 int8 = 1
            _ = v44
        }
        _ = len(v3)
    } else {
        for fi45 := 0; fi45 < 15; fi45++ {
            switch "" {
            case "hello":
                for fi46 := 0; fi46 < 18; fi46++ {
                    v1 -= 0.1
                    if true {
                        _ = len(v3)
                        var v47 int64 = -9223372036854775808
                        switch -2147483648 {
                        case 0:
                            v1 = 3.14159
                        case 42:
                            switch (fi45 + -1) {
                            case 4294967295:
                                if true {
                                    var v48 int32 = (v2 + v2)
                                    var v49 uint16 = 28
                                    _ = v48
                                    _ = v49
                                } else {
                                    fmt.Println(v47)
                                    v50 := GenS3{F0: (!(float64(v1) == 1e10)), F1: "a", F2: 2}
                                    v51 := GenS4{F0: 127}
                                    _ = v50
                                    _ = v51
                                }
                            case -2147483648:
                                v1 = 1e10
                                switch 405 {
                                case 902:
                                    if true {
                                        v52 := GenS5{F0: (-(fi45)), F1: 127, F2: 42}
                                        if true {
                                            if (1 <= 1) {
                                                v2 -= 42
                                            }
                                            var v53 bool = (v1 < v1)
                                            delete(v3, (fi46 << ((1) & 63)))
                                            if true {
                                                if false {
                                                    v54 := []int{(fi45 >> ((921) & 63)), fi46}
                                                    fmt.Println(v2, fi46)
                                                    switch (fi46 >> ((2) & 63)) {
                                                    case 32768:
                                                        if v53 {
                                                            v55 := GenS6{F0: 127}
                                                            if len(v54) >= 2 {
                                                                v54 = v54[1:len(v54)]
                                                            }
                                                            _ = v55
                                                        } else {
                                                            v1 += ((v1 / v1) * v1)
                                                        }
                                                    }
                                                    _ = v54
                                                } else {
                                                    if v53 {
                                                        v56 := v3[1]
                                                        v57 := GenS7{F0: -164}
                                                        v58 := []int{(fi46 >> ((255) & 63)), 2147483647, (fi45 + (-(fi45))), 4294967295}
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                    } else {
                                                        var v59 int8 = -1
                                                        v60 := []int{fi46, fi45, fi46, -2147483648}
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                }
                                                switch fi45 {
                                                case -9223372036854775808:
                                                    var v61 float32 = 1.5
                                                    _ = v61
                                                case -2147483648:
                                                    if (!true) {
                                                        var v62 bool = (v2 < v2)
                                                        var v63 int64 = -9223372036854775808
                                                        var v64 string = ("foo bar" + ("" + "foo bar"))
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                    switch ("\n\t" + "") {
                                                    case "hello":
                                                        v53 = false
                                                    case "\n\t":
                                                        v47 = (v47 | -274)
                                                    case "foo bar":
                                                        v53 = (128 < 1)
                                                    default:
                                                        fmt.Println(fi46)
                                                    }
                                                    v2 = v2
                                                    var v65 uint32 = 834
                                                    _ = v65
                                                default:
                                                    fmt.Println(v1)
                                                }
                                            }
                                            _ = v53
                                        }
                                        v66 := GenS1{F0: fi46}
                                        _ = v52
                                        _ = v66
                                    }
                                    if true {
                                        if false {
                                            var v67 int32 = (^(v2))
                                            var v68 uint = 42
                                            v68 += (v68 >> ((42) & 63))
                                            _ = v67
                                            _ = v68
                                        } else {
                                            v69 := v3[fi46]
                                            var v70 float32 = v1
                                            fmt.Println(v69, v70)
                                            _ = v69
                                            _ = v70
                                        }
                                        var v71 string = "foo bar"
                                        v72 := GenS8{F0: 665, F1: true, F2: (^(v2))}
                                        _ = v71
                                        _ = v72
                                    }
                                case -9223372036854775808:
                                    v2 = v2
                                default:
                                    v73 := GenS9{F0: (fi46 >> ((984) & 63)), F1: 2}
                                    var v74 float32 = ((v1 * -2.25) * float32(((3.14159 - (0.1 / -2.25)) + (1.5 + 0.1))))
                                    v75 := []int{fi46, fi45, fi46}
                                    fmt.Println(v47, v47, v74)
                                    v74 -= v1
                                    _ = v73
                                    _ = v74
                                    _ = v75
                                }
                            case -1:
                                switch ("foo bar" + "a") {
                                case "foo bar":
                                    v47 += (v47 | v47)
                                case "hello":
                                    v2 = (v2 / (((v2 - (v2 + 701))) | 1))
                                case "\n\t":
                                    v76 := GenS10{F0: (-(v47)), F1: (fi46 * (fi46 / ((fi46) | 1))), F2: 2}
                                    _ = v76
                                }
                            default:
                                v77 := GenS4{F0: 160}
                                switch (^(fi46)) {
                                case 32767:
                                    var v78 uint32 = 2147483648
                                    _ = v78
                                default:
                                    var v79 float64 = 0.0
                                    _ = v79
                                }
                                _ = v77
                            }
                            fmt.Println(fi45, v1, v2)
                            v2 = (v2 * v2)
                        case -9223372036854775808:
                            var v80 uint32 = 799
                            _ = v80
                        }
                        v81 := GenS7{F0: 42}
                        _ = v47
                        _ = v81
                    }
                    if (255 != 32767) {
                        var v82 string = "foo bar"
                        _ = v82
                    }
                    _ = fi46
                }
            default:
                v1 -= 3.14159
            }
            for fi83 := 0; fi83 < 2; fi83++ {
                var v84 uint64 = 867
                v85 := GenS10{F0: 1, F1: (^(fi45)), F2: 42}
                v86 := []int{}
                _ = fi83
                _ = v84
                _ = v85
                _ = v86
            }
            if (!(18446744073709551615 <= 1)) {
                v87 := GenS9{F0: -2147483648, F1: 1}
                for fi88 := 0; fi88 < 1; fi88++ {
                    var v89 int = (fi45 & (fi45 & fi88))
                    _ = fi88
                    _ = v89
                }
                _ = v87
            }
            _ = fi45
        }
        v90 := GenS11{F0: 127}
        switch 2147483647 {
        case -9223372036854775808:
            for fi91 := 0; fi91 < 9; fi91++ {
                var v92 uint64 = 159
                _ = len(v3)
                _ = fi91
                _ = v92
            }
        case 430:
            if true {
                v93 := []int{2147483647, -9223372036854775808, int(v2), 9223372036854775807}
                v2 = v2
                if false {
                    for fi94 := 0; fi94 < 3; fi94++ {
                        if (true || ((fi94 + (-(fi94))) != 32768)) {
                            var v95 int64 = 42
                            _ = v95
                        } else {
                            v96 := GenS9{F0: (fi94 % ((fi94) | 1)), F1: 1}
                            if len(v93) > 0 {
                                v93[0] = (fi94 | (fi94 & (fi94 / (((-(fi94))) | 1))))
                            }
                            var v97 uint = 593
                            _ = v96
                            _ = v97
                        }
                        fmt.Println(v1)
                        _ = fi94
                    }
                }
                _ = v93
            } else {
                v98 := v3[9223372036854775807]
                _ = v98
            }
        case 32767:
            if true {
                v99 := GenS12{F0: -1, F1: 127, F2: 2}
                _ = v99
            }
        }
        _ = v90
    }
    if false {
        v2 = v2
        var v100 int64 = -9223372036854775808
        _ = v100
    }
    fmt.Println(v1, v2)
    if false {
        v1 += (1e10 * float32(0.1))
        if (!(42 > 127)) {
            fmt.Println(v1, v2)
            for fi101 := 0; fi101 < 10; fi101++ {
                v102 := GenS2{F0: 32768}
                _ = fi101
                _ = v102
            }
            _ = len(v3)
        }
    }
    switch 1 {
    case 660:
        for fi103 := 0; fi103 < 8; fi103++ {
            var v104 bool = (((int16(fi103) != 0) && false) || true)
            var v105 int8 = 0
            switch "foo bar" {
            case "a":
                switch fi103 {
                case -631:
                    v106 := GenS13{F0: 2, F1: 127, F2: 255}
                    _ = v106
                case 127:
                    v1 = v1
                default:
                    v2 = (v2 * v2)
                }
            default:
                var v107 int16 = 1000
                _ = v107
            }
            _ = fi103
            _ = v104
            _ = v105
        }
    case 2147483647:
        v1 = 0.1
    }
    fmt.Println(v1, v2)
    var v108 uint16 = 876
    if false {
        fmt.Println(v108, v108)
    }
    v2 -= (v2 >> ((255) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v108
    _ = v2
    _ = v3
}
