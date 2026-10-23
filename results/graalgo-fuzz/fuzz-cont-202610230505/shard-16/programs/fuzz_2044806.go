package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 bool; F1 int16; F2 int64 }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 uint8; F1 int16 }
type GenS8 struct { F0 uint64; F1 int8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint32; F1 int; F2 uint32 }
type GenS11 struct { F0 int8; F1 string; F2 int32 }
type GenS12 struct { F0 bool; F1 int16; F2 string }
type GenS13 struct { F0 int16; F1 int16; F2 uint64 }
type GenS14 struct { F0 int64; F1 bool }
type GenS15 struct { F0 uint; F1 uint16 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 uint64; F1 uint8; F2 int }
type GenS18 struct { F0 int; F1 int; F2 bool }
type GenS19 struct { F0 uint64 }

func main() {
    v1 := []int{9223372036854775807, 128}
    var v2 uint32 = 1
    v3 := make(map[int]bool)
    var v4 uint16 = 32767
    var v5 float32 = (1.5 + (0.0 + 3.14159))
    switch 32768 {
    case -9223372036854775808:
        if (42 > -110) {
            if (0 <= -53) {
                if (42 == 2) {
                    switch ((("foo bar" + "") + "") + ("hello" + "foo bar")) {
                    case "a":
                        v4 -= (v4 << ((255) & 63))
                    case "\n\t":
                        v6 := GenS1{F0: uint64(v2), F1: -9223372036854775808, F2: v4}
                        _ = v6
                    }
                    for fi7 := 0; fi7 < 5; fi7++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v2 = v2
                        _ = fi7
                    }
                } else {
                    fmt.Println(v4)
                    for fi8 := 0; fi8 < 13; fi8++ {
                        v9 := GenS2{F0: (1e10 == float32(0.0)), F1: v4}
                        fmt.Println(v2, v4, fi8)
                        _ = fi8
                        _ = v9
                    }
                    v10 := GenS3{F0: -9223372036854775808, F1: 1}
                    _ = v10
                }
                for fi11 := 0; fi11 < 14; fi11++ {
                    for fi12 := 0; fi12 < 19; fi12++ {
                        if (!false) {
                            v5 -= (float32(3.14159) * (((0.0 + (0.1 / 1.5)) - (3.14159 / ((float32((3.14159 - (0.1 + 1e10))) / ((1e10 - 0.1) / 3.14159)) - 1.5))) / v5))
                            var v13 float32 = v5
                            _ = len(v1)
                            _ = v13
                        } else {
                            v14 := GenS4{F0: 42, F1: -1}
                            _ = v14
                        }
                        fmt.Println(v5, fi12)
                        v15 := GenS4{F0: 127, F1: -603}
                        v4 += (v4 % ((760) | 1))
                        switch (("a" + "foo bar") + ("foo bar" + "\n\t")) {
                        case "":
                            v4 = 32767
                        case "\n\t":
                            v16 := v3[(fi12 * (fi12 & (fi12 << ((711) & 63))))]
                            _ = v16
                        case "a":
                            var v17 uint64 = 4294967295
                            _ = v17
                        default:
                            var v18 string = ("hello" + "")
                            _ = v18
                        }
                        _ = fi12
                        _ = v15
                    }
                    v4 = (v4 * v4)
                    v5 = (v5 * v5)
                    _ = fi11
                }
            } else {
                switch 42 {
                case -1:
                    _ = len(v1)
                case 122:
                    v4 = v4
                case 732:
                    fmt.Println(v4, v2)
                default:
                    for fi19 := 0; fi19 < 15; fi19++ {
                        switch fi19 {
                        case 1:
                            switch "hello" {
                            case "\n\t":
                                var v20 bool = (((0 >= -1) && true) && (false || true))
                                _ = v20
                            case "":
                                if len(v1) > 0 {
                                    v1[0] = fi19
                                }
                            default:
                                v4 = v4
                            }
                        case 0:
                            if true {
                                if (0 >= 0) {
                                    v21 := GenS4{F0: uint8(v2), F1: 0}
                                    v22 := GenS2{F0: false, F1: (v4 | v4)}
                                    _ = v21
                                    _ = v22
                                }
                                v2 += 0
                                v23 := GenS4{F0: 1, F1: 556}
                                _ = v23
                            } else {
                                v4 = 0
                                fmt.Println(v4, v2)
                            }
                        }
                        _ = fi19
                    }
                }
                switch ("hello" + "\n\t") {
                case "":
                    v24 := GenS3{F0: 1, F1: -1}
                    _ = v24
                default:
                    v3[4294967295] = true
                }
            }
            for fi25 := 0; fi25 < 13; fi25++ {
                var v26 bool = (!true)
                if v26 {
                    var v27 uint = 65535
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 5; fi28++ {
                        switch ((("foo bar" + ("a" + "a")) + "\n\t") + "hello") {
                        case "a":
                            v29 := GenS2{F0: ((v2 / ((2147483647) | 1)) < (v2 & v2)), F1: (v4 - (v4 - 597))}
                            _ = v29
                        case "hello":
                            var v30 uint32 = v2
                            _ = v30
                        default:
                            v2 -= 4294967295
                        }
                        switch "foo bar" {
                        case "a":
                            switch (fi28 - fi28) {
                            case 32768:
                                fmt.Println(v5, v2, fi25)
                            }
                        case "hello":
                            fmt.Println(fi28, v26, v2)
                        default:
                            v31 := GenS5{F0: (!true), F1: 1, F2: 9223372036854775807}
                            _ = v31
                        }
                        _ = fi28
                    }
                    v32 := GenS6{F0: v2, F1: int16(v4)}
                    fmt.Println(v5)
                    _ = v32
                }
                _ = fi25
                _ = v26
            }
        } else {
            v33 := GenS7{F0: 3, F1: -32768}
            _ = v33
        }
    case 42:
        fmt.Println(v5, v4, v4)
    default:
        delete(v3, 9223372036854775807)
    }
    var v34 uint32 = 2
    var v35 string = ""
    v34 -= (v2 - v2)
    for fi36 := 0; fi36 < 2; fi36++ {
        _ = len(v3)
        var v37 string = "\n\t"
        v2 -= (^(v34))
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 14; fi38++ {
        var v39 int32 = 1
        v4 += v4
        _ = fi38
        _ = v39
    }
    v3[-1] = false
    var v40 int = -719
    var v41 uint = 2
    var v42 uint32 = v34
    fmt.Println(v42)
    v41 += v41
    for fi43 := 0; fi43 < 13; fi43++ {
        v1 = append(v1, v40)
        if (2 <= 18446744073709551615) {
            v44 := GenS8{F0: 305, F1: 100}
            _ = v44
        }
        _ = fi43
    }
    for fi45 := 0; fi45 < 16; fi45++ {
        for fi46 := 0; fi46 < 14; fi46++ {
            if true {
                v47 := GenS9{F0: -582}
                if true {
                    switch 9223372036854775807 {
                    case 127:
                        if true {
                            if (!(1 > 2147483647)) {
                                v40 -= (v40 - v40)
                                if (fi45 >= fi45) {
                                    var v48 int64 = 9223372036854775807
                                    var v49 uint32 = (v2 | v34)
                                    fmt.Println(v40, v34)
                                    _ = v48
                                    _ = v49
                                } else {
                                    fmt.Println(v34, v35)
                                    var v50 uint16 = v4
                                    fmt.Println(v4, v40)
                                    _ = v50
                                }
                                switch 1 {
                                case 865:
                                    if false {
                                        if (!(false || ((!false) || false))) {
                                            fmt.Println(v35, v34, v5)
                                            v42 += (^(v34))
                                        }
                                        if true {
                                            if (v35 >= v35) {
                                                v51 := GenS4{F0: 2, F1: -2147483648}
                                                var v52 uint32 = 2147483648
                                                fmt.Println(v5, v52, v42)
                                                _ = v51
                                                _ = v52
                                            } else {
                                                v4 = (v4 + (v4 * v4))
                                            }
                                        }
                                    } else {
                                        v53 := GenS3{F0: v40, F1: 42}
                                        var v54 uint16 = v4
                                        _ = v53
                                        _ = v54
                                    }
                                    switch "\n\t" {
                                    case "foo bar":
                                        v34 = 42
                                    case "":
                                        switch v40 {
                                        case 673:
                                            var v55 uint8 = 127
                                            _ = v55
                                        case 2147483647:
                                            v2 += (v2 + v34)
                                            v4 -= 32767
                                        default:
                                            switch v35 {
                                            case "a":
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case "":
                                                v35 = v35
                                            }
                                        }
                                    case "a":
                                        switch v35 {
                                        case "a":
                                            var v56 float32 = v5
                                            _ = v56
                                        case "foo bar":
                                            var v57 uint16 = (v4 & v4)
                                            _ = v57
                                        default:
                                            switch v35 {
                                            case "a":
                                                v42 = uint32(v4)
                                            case "":
                                                var v58 int = (^(v40))
                                                _ = v58
                                            default:
                                                switch v35 {
                                                case "hello":
                                                    fmt.Println(v5, v2)
                                                case "a":
                                                    v40 += fi46
                                                    if (!(v42 <= (v2 / ((2) | 1)))) {
                                                        switch v35 {
                                                        case "\n\t":
                                                            _ = len(v3)
                                                        case "":
                                                            if true {
                                                                v59 := GenS10{F0: v34, F1: fi45, F2: (^(v42))}
                                                                var v60 int = 2147483647
                                                                delete(v3, -1)
                                                                _ = v59
                                                                _ = v60
                                                            } else {
                                                                v41 = (v41 & v41)
                                                            }
                                                        }
                                                        switch "a" {
                                                        case "\n\t":
                                                            if (255 <= v41) {
                                                                var v61 int16 = 42
                                                                _ = v61
                                                            }
                                                        default:
                                                            if false {
                                                                v3[486] = (474 <= 1)
                                                                var v62 uint64 = 4294967295
                                                                _ = v62
                                                            }
                                                        }
                                                        var v63 int = v40
                                                        var v64 int = fi45
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    if true {
                                        var v65 uint = v41
                                        var v66 int64 = 0
                                        _ = v65
                                        _ = v66
                                    }
                                case 127:
                                    if (!true) {
                                        v67 := GenS11{F0: -60, F1: ("\n\t" + v35), F2: 1}
                                        var v68 uint16 = (^(v4))
                                        _ = v67
                                        _ = v68
                                    }
                                case 128:
                                    v69 := GenS12{F0: true, F1: -86, F2: v35}
                                    _ = v69
                                }
                                _ = len(v3)
                            }
                        } else {
                            switch fi46 {
                            case 42:
                                switch ("" + (v35 + "hello")) {
                                case "foo bar":
                                    if (443 < 42) {
                                        _ = len(v1)
                                        v70 := GenS10{F0: (v34 ^ (v42 % (((v34 << ((v41) & 63))) | 1))), F1: v40, F2: v34}
                                        v71 := GenS9{F0: 0}
                                        _ = v70
                                        _ = v71
                                    }
                                case "":
                                    _ = len(v3)
                                case "hello":
                                    fmt.Println(v5, v2, v42)
                                default:
                                    var v72 uint32 = 42
                                    _ = v72
                                }
                            }
                            if len(v1) > 0 {
                                v1[0] = (fi46 % ((2147483647) | 1))
                            }
                            v35 = (v35 + "")
                            var v73 int = (v40 * (fi45 * (fi45 >> (((v41 / ((18446744073709551615) | 1))) & 63))))
                            _ = v73
                        }
                        v1 = append(v1, fi46)
                    case 42:
                        v40 -= fi45
                        v74 := GenS13{F0: 32767, F1: 569, F2: 42}
                        _ = v74
                    case 128:
                        if (2147483647 == 911) {
                            v75 := GenS14{F0: 1, F1: (fi46 == 1)}
                            v42 += uint32(v41)
                            _ = v75
                        }
                        v41 -= (v41 % (((^(v41))) | 1))
                    default:
                        var v76 float32 = v5
                        _ = v76
                    }
                    fmt.Println(v42)
                    if false {
                        fmt.Println(fi46)
                    } else {
                        var v77 uint32 = v42
                        switch fi45 {
                        case 127:
                            if (1 > -1) {
                                v35 = (v35 + (v35 + v35))
                                _ = len(v3)
                                _ = len(v3)
                            }
                        case 825:
                            switch "a" {
                            case "a":
                                v77 = (v34 % (((v2 + (v77 * (v2 % ((364) | 1))))) | 1))
                            default:
                                if len(v1) > 0 {
                                    v1[0] = 32767
                                }
                            }
                        case 128:
                            var v78 int32 = 0
                            _ = v78
                        default:
                            v79 := GenS15{F0: v41, F1: v4}
                            _ = v79
                        }
                        v80 := GenS4{F0: 128, F1: -1}
                        _ = v77
                        _ = v80
                    }
                }
                v41 = 255
                _ = len(v3)
                switch fi45 {
                case -932:
                    switch v35 {
                    case "foo bar":
                        delete(v3, -108)
                    case "":
                        switch "" {
                        case "a":
                            if true {
                                if (-543 != 788) {
                                    v35 = ((("\n\t" + v35) + v35) + "foo bar")
                                    if true {
                                        v41 = v41
                                        if false {
                                            switch v35 {
                                            case "\n\t":
                                                var v81 uint32 = (v42 << ((v41) & 63))
                                                _ = v81
                                            case "a":
                                                fmt.Println(v42)
                                            default:
                                                var v82 int16 = 642
                                                switch ("\n\t" + (v35 + ("a" + ((v35 + "foo bar") + (("" + (v35 + v35)) + ""))))) {
                                                case "a":
                                                    v82 += v82
                                                case "foo bar":
                                                    v4 = 2
                                                case "\n\t":
                                                    v2 += v42
                                                }
                                                fmt.Println(fi46, fi45)
                                                switch (fi46 + v40) {
                                                case 2147483647:
                                                    fmt.Println(v2)
                                                    var v83 int = 0
                                                    _ = v83
                                                case 128:
                                                    var v84 uint16 = (v4 * v4)
                                                    _ = v84
                                                default:
                                                    switch v35 {
                                                    case "hello":
                                                        v85 := GenS16{F0: (v41 | v41)}
                                                        _ = v85
                                                    case "\n\t":
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v5 += v5
                                                        case "hello":
                                                            if false {
                                                                v86 := v3[v40]
                                                                _ = v86
                                                            }
                                                        case "":
                                                            delete(v3, fi46)
                                                        }
                                                    case "foo bar":
                                                        if (2147483647 != -1) {
                                                            if false {
                                                                if len(v1) > 0 {
                                                                    v1[0] = (fi45 | (^(fi45)))
                                                                }
                                                            } else {
                                                                var v87 float64 = 0.1
                                                                _ = v87
                                                            }
                                                            v4 = (v4 / (((v4 << ((v41) & 63))) | 1))
                                                            fmt.Println(v41)
                                                            var v88 float32 = v5
                                                            var v89 int8 = 1
                                                            _ = v88
                                                            _ = v89
                                                        } else {
                                                            v5 = v5
                                                        }
                                                    default:
                                                        v5 += v5
                                                    }
                                                }
                                                fmt.Println(v5)
                                                _ = v82
                                            }
                                        }
                                        fmt.Println(v34)
                                    } else {
                                        fmt.Println(v34)
                                    }
                                    if true {
                                        var v90 uint8 = 128
                                        v91 := GenS10{F0: v2, F1: (fi45 + fi45), F2: (v2 << ((4294967295) & 63))}
                                        _ = v90
                                        _ = v91
                                    } else {
                                        v34 = v42
                                        v92 := GenS12{F0: ((1 > 752) || true), F1: -508, F2: "foo bar"}
                                        var v93 uint = 954
                                        fmt.Println(fi45, v34, v2)
                                        if false {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v3[v40] = (false && (true && false))
                                        } else {
                                            v94 := GenS17{F0: 2, F1: 195, F2: (fi45 ^ int(v4))}
                                            _ = v94
                                        }
                                        delete(v3, 9223372036854775807)
                                        var v95 int16 = 1
                                        _ = len(v3)
                                        _ = v92
                                        _ = v93
                                        _ = v95
                                    }
                                }
                                v2 += v2
                            }
                        case "":
                            v96 := v3[0]
                            _ = v96
                        default:
                            v3[32768] = false
                        }
                    default:
                        var v97 uint64 = 789
                        _ = v97
                    }
                default:
                    v42 = v42
                }
                if ((1 < 112) || true) {
                    var v98 int64 = -815
                    switch v35 {
                    case "\n\t":
                        v99 := GenS13{F0: -918, F1: 4, F2: uint64(v4)}
                        _ = v99
                    case "a":
                        if len(v1) > 0 {
                            v1[0] = v40
                        }
                    case "foo bar":
                        var v100 uint8 = 42
                        _ = v100
                    default:
                        var v101 uint8 = 9
                        _ = v101
                    }
                    fmt.Println(v2)
                    _ = v98
                } else {
                    fmt.Println(v41)
                    if len(v1) > 0 {
                        v1[0] = 128
                    }
                    v2 -= 2
                }
                v41 -= 65535
                _ = v47
            }
            _ = fi46
        }
        v42 = (v2 | v2)
        _ = fi45
    }
    for fi102 := 0; fi102 < 16; fi102++ {
        for fi103 := 0; fi103 < 2; fi103++ {
            v104 := GenS10{F0: (v42 - (v2 ^ v2)), F1: v40, F2: v42}
            _ = fi103
            _ = v104
        }
        v105 := GenS18{F0: fi102, F1: v40, F2: ((v41 ^ (^(v41))) > v41)}
        v106 := GenS19{F0: 0}
        _ = fi102
        _ = v105
        _ = v106
    }
    if (!(!(false || (!false)))) {
        for fi107 := 0; fi107 < 5; fi107++ {
            _ = len(v3)
            v42 = 887
            _ = fi107
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
