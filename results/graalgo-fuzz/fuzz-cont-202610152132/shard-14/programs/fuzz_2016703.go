package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 string; F2 uint16 }
type GenS4 struct { F0 uint; F1 uint32; F2 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int; F1 uint64; F2 uint32 }
type GenS11 struct { F0 int64; F1 uint; F2 uint }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint8; F1 uint8 }
type GenS14 struct { F0 uint8; F1 int; F2 uint16 }
type GenS15 struct { F0 int32; F1 int16 }
type GenS16 struct { F0 int32; F1 int16; F2 uint8 }
func genFn0(p0 uint32, p1 int64, p2 uint16) int16 {
    var v6 int64 = p1
    v7 := GenS1{F0: 165}
    if ((18446744073709551615 <= 18446744073709551615) && ((!(!(1 < 0))) && false)) {
        fmt.Println(p2, p1)
        if true {
            var v8 bool = true
            p1 = v6
            var v9 int16 = -987
            _ = v8
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 7; fi10++ {
            v6 = (p1 + v6)
            var v11 uint16 = p2
            _ = fi10
            _ = v11
        }
    }
    var v12 int8 = 8
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int, p1 uint16, p2 string) string {
    switch p2 {
    case "foo bar":
        switch -9223372036854775808 {
        case 9223372036854775807:
            p1 = (p1 >> ((467) & 63))
        case 32768:
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 float64 = 1e10
                _ = fi13
                _ = v14
            }
        case 42:
            p1 += p1
        }
    case "\n\t":
        switch (((("foo bar" + p2) + (p2 + p2)) + (p2 + (((("foo bar" + (p2 + ((p2 + p2) + "a"))) + p2) + p2) + (p2 + "a")))) + p2) {
        case "a":
            var v15 uint8 = 1
            _ = v15
        default:
            if (-85 != -1) {
                var v16 int32 = 42
                _ = v16
            } else {
                if (!false) {
                    v17 := GenS2{F0: -72}
                    _ = v17
                }
            }
        }
    }
    p0 += -486
    v18 := GenS3{F0: (((0.1 + 1.5) - 1.5) > (float64(0.1) * ((1.5 * -2.25) + float64(float32(-2.25))))), F1: "foo bar", F2: 965}
    var v19 int32 = 2147483647
    p2 = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return (p2 + "hello")
}

func main() {
    var v1 int64 = 433
    var v2 string = "hello"
    v3 := make(map[int]int)
    var v4 uint = 1
    v5 := make(map[uint32]string)
    v20 := GenS4{F0: (v4 / (((v4 / ((2) | 1))) | 1)), F1: 669, F2: ((("hello" + ((v2 + ("" + "hello")) + "hello")) + "hello") + v2)}
    v1 = (v1 * v1)
    switch v2 {
    case "a":
        var v21 int32 = 1
        _ = v21
    default:
        v22 := []int{473, 2147483647}
        _ = v22
    }
    v23 := []int{}
    for fi24 := 0; fi24 < 8; fi24++ {
        _ = len(v23)
        delete(v5, 2)
        var v25 int = fi24
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        if (851 > 432) {
            switch fi26 {
            case 128:
                v1 = int64(fi26)
            case 1:
                switch "hello" {
                case "\n\t":
                    _ = len(v3)
                case "foo bar":
                    if true {
                        for fi27 := 0; fi27 < 20; fi27++ {
                            switch (fi26 + (fi26 + fi26)) {
                            case -9223372036854775808:
                                v2 = ""
                            default:
                                v28 := GenS1{F0: 42}
                                _ = v28
                            }
                            switch (fi26 & (fi26 / ((fi26) | 1))) {
                            case -73:
                                v2 = ""
                            }
                            if (true || true) {
                                v29 := GenS5{F0: 0}
                                var v30 string = (v2 + "")
                                _ = v29
                                _ = v30
                            }
                            _ = fi27
                        }
                        v23 = append(v23, (fi26 * fi26))
                        v31 := v3[fi26]
                        _ = v31
                    } else {
                        v4 = (v4 % ((v4) | 1))
                        if (!true) {
                            v2 = (v2 + v2)
                            for fi32 := 0; fi32 < 19; fi32++ {
                                var v33 int32 = 0
                                _ = fi32
                                _ = v33
                            }
                        }
                        v34 := GenS1{F0: 774}
                        _ = v34
                    }
                case "a":
                    switch (fi26 % (((fi26 % ((fi26) | 1))) | 1)) {
                    case 760:
                        if true {
                            switch (fi26 * fi26) {
                            case 0:
                                var v35 uint = (v4 % (((v4 % (((v4 % (((v4 << ((18446744073709551615) & 63))) | 1))) | 1))) | 1))
                                _ = v35
                            default:
                                switch v2 {
                                case "foo bar":
                                    var v36 int8 = 0
                                    _ = v36
                                case "a":
                                    switch v2 {
                                    case "":
                                        v37 := GenS2{F0: -1}
                                        _ = v37
                                    default:
                                        if (255 <= 0) {
                                            var v38 int = fi26
                                            var v39 uint32 = 4
                                            for fi40 := 0; fi40 < 17; fi40++ {
                                                v1 = 0
                                                v39 += (v39 / ((v39) | 1))
                                                v4 = v4
                                                _ = fi40
                                            }
                                            _ = v38
                                            _ = v39
                                        }
                                    }
                                case "\n\t":
                                    if (((1.5 / 1.0) + 0.1) > 3.14159) {
                                        v1 += 0
                                    }
                                }
                            }
                            var v41 uint64 = 750
                            if true {
                                var v42 uint32 = 2147483647
                                _ = v42
                            }
                            _ = v41
                        } else {
                            _ = len(v3)
                            v1 -= v1
                            v43 := GenS6{F0: 0}
                            _ = v43
                        }
                    case 32767:
                        if (664 > 0) {
                            v2 = v2
                        }
                    case -567:
                        v4 = v4
                    }
                default:
                    for fi44 := 0; fi44 < 9; fi44++ {
                        v45 := GenS2{F0: 100}
                        if true {
                            switch ((("a" + v2) + ("a" + (v2 + v2))) + (v2 + (v2 + v2))) {
                            case "a":
                                if true {
                                    var v46 int64 = (v1 ^ (v1 - (v1 | v1)))
                                    v2 = ("foo bar" + "hello")
                                    _ = v46
                                }
                            case "\n\t":
                                var v47 bool = ((-1 == -904) && false)
                                _ = v47
                            default:
                                if (((18446744073709551615 <= (v4 * v4)) || true) || ((128 != 109) || (!(false || (v2 == "a"))))) {
                                    if len(v23) > 0 {
                                        v23[0] = fi44
                                    }
                                    var v48 int32 = 0
                                    if true {
                                        switch "" {
                                        case "a":
                                            v2 = v2
                                        case "":
                                            _ = len(v23)
                                        case "\n\t":
                                            switch (v2 + (v2 + v2)) {
                                            case "\n\t":
                                                v49 := GenS7{F0: -128}
                                                _ = v49
                                            default:
                                                fmt.Println(fi26)
                                            }
                                        default:
                                            switch 128 {
                                            case 2147483647:
                                                switch int(v1) {
                                                case 0:
                                                    v4 += (v4 << ((v4) & 63))
                                                case 42:
                                                    var v50 uint8 = 77
                                                    _ = v50
                                                case -12:
                                                    if true {
                                                        if false {
                                                            if false {
                                                                if false {
                                                                    var v51 int16 = -46
                                                                    v48 -= v48
                                                                    _ = v51
                                                                }
                                                                var v52 uint64 = 18446744073709551615
                                                                _ = v52
                                                            } else {
                                                                if (!(!(false && (!(((0.0 + -2.25) * 3.14159) >= float32(0.0)))))) {
                                                                    v48 = (v48 >> (((v4 * (v4 & (v4 - 255)))) & 63))
                                                                    fmt.Println(v1, v1, v1)
                                                                }
                                                            }
                                                            v1 = -160
                                                        }
                                                        var v53 uint8 = 25
                                                        var v54 uint16 = 943
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                    if len(v23) >= 2 {
                                                        v23 = v23[1:len(v23)]
                                                    }
                                                    var v55 uint16 = 65535
                                                    _ = v55
                                                }
                                            case -96:
                                                if (false && true) {
                                                    v2 = v2
                                                    if false {
                                                        v1 -= v1
                                                        if (0.0 < -2.25) {
                                                            var v56 uint8 = 127
                                                            var v57 int64 = -357
                                                            _ = v56
                                                            _ = v57
                                                        } else {
                                                            v58 := GenS1{F0: 397}
                                                            delete(v3, (^(fi26)))
                                                            _ = v58
                                                        }
                                                    } else {
                                                        if (!false) {
                                                            if false {
                                                                v1 += v1
                                                            }
                                                        }
                                                        var v59 uint = 4294967295
                                                        _ = v59
                                                    }
                                                    var v60 string = "a"
                                                    _ = v60
                                                }
                                            case 0:
                                                v4 = 185
                                            default:
                                                if (!true) {
                                                    if (781 >= -174) {
                                                        v61 := GenS8{F0: 42}
                                                        _ = v61
                                                    }
                                                }
                                            }
                                        }
                                        if ((0.0 + 0.1) < (float32(-2.25) / (0.1 * 3.14159))) {
                                            fmt.Println(v2, fi44, fi44)
                                            v48 -= (v48 / ((v48) | 1))
                                        } else {
                                            switch fi26 {
                                            case 410:
                                                var v62 int16 = -576
                                                _ = v62
                                            case 602:
                                                var v63 float32 = float32(1.5)
                                                _ = v63
                                            default:
                                                v23 = append(v23, 32768)
                                            }
                                            fmt.Println(fi26)
                                            v23 = append(v23, (fi44 * fi44))
                                        }
                                        if ((!(177 == 42)) && (127 >= 127)) {
                                            fmt.Println(v2)
                                            v1 = v1
                                        }
                                    }
                                    _ = v48
                                } else {
                                    if (1.5 > 3.14159) {
                                        if ((!true) || false) {
                                            if false {
                                                var v64 int64 = (^(v1))
                                                if (!((v4 | 0) > v4)) {
                                                    var v65 uint = 4294967295
                                                    _ = v65
                                                }
                                                v4 = v4
                                                _ = v64
                                            } else {
                                                var v66 int16 = 42
                                                _ = v66
                                            }
                                            var v67 bool = (!(false || (!true)))
                                            v68 := GenS9{F0: 2}
                                            _ = v67
                                            _ = v68
                                        }
                                    }
                                    v69 := GenS4{F0: v4, F1: 42, F2: (v2 + "")}
                                    v70 := GenS10{F0: fi26, F1: 42, F2: 0}
                                    _ = v69
                                    _ = v70
                                }
                            }
                            switch v2 {
                            case "foo bar":
                                v71 := GenS8{F0: 0}
                                _ = v71
                            case "a":
                                switch int(v4) {
                                case -2147483648:
                                    var v72 int64 = (v1 << (((v4 << ((42) & 63))) & 63))
                                    _ = v72
                                case 127:
                                    v73 := GenS7{F0: -23}
                                    _ = v73
                                default:
                                    if false {
                                        var v74 int16 = 32767
                                        _ = v74
                                    } else {
                                        switch fi26 {
                                        case -731:
                                            switch fi26 {
                                            case 803:
                                                v75 := GenS6{F0: 0}
                                                _ = v75
                                            default:
                                                v4 = v4
                                            }
                                        case -1:
                                            v76 := v3[int(v1)]
                                            if (0.0 < 3.14159) {
                                                var v77 bool = false
                                                var v78 uint16 = 65535
                                                v79 := GenS11{F0: (v1 ^ -821), F1: v4, F2: 255}
                                                _ = v77
                                                _ = v78
                                                _ = v79
                                            }
                                            v76 = int(v4)
                                            v1 += 42
                                            if (100 > -58) {
                                                v80 := GenS7{F0: 100}
                                                v5[690] = v2
                                                v81 := GenS2{F0: 125}
                                                _ = v80
                                                _ = v81
                                            }
                                            _ = v76
                                        }
                                    }
                                }
                            case "\n\t":
                                var v82 float32 = (-2.25 + (1e10 - 3.14159))
                                _ = v82
                            }
                            if (v2 == v2) {
                                v83 := GenS1{F0: 1}
                                fmt.Println(v2, fi44, fi26)
                                v2 = (v2 + v2)
                                _ = v83
                            }
                        }
                        if (3.14159 != 3.14159) {
                            _ = len(v3)
                            v84 := GenS12{F0: -1}
                            if (v4 > 42) {
                                if false {
                                    if ((!((!true) && false)) && (("hello" + (v2 + v2)) <= "a")) {
                                        if len(v23) >= 2 {
                                            v23 = v23[1:len(v23)]
                                        }
                                    } else {
                                        v2 = v2
                                        var v85 int = (fi44 | (fi44 * (fi44 / ((244) | 1))))
                                        switch (v85 | fi44) {
                                        case -1:
                                            v2 = v2
                                        case 96:
                                            if false {
                                                v1 -= -503
                                                v2 = v2
                                                v86 := GenS13{F0: 100, F1: 134}
                                                _ = v86
                                            }
                                        case 0:
                                            v1 = (v1 | v1)
                                        }
                                        fmt.Println(v1, v2)
                                        if (((v2 + v2) + v2) != v2) {
                                            var v87 int16 = 32767
                                            _ = v87
                                        } else {
                                            var v88 uint64 = 524
                                            v2 = "\n\t"
                                            var v89 uint = (v4 / ((90) | 1))
                                            _ = v88
                                            _ = v89
                                        }
                                        _ = v85
                                    }
                                    v5[1] = v2
                                }
                            } else {
                                if (!false) {
                                    fmt.Println(v4, v1, fi44)
                                    switch -1 {
                                    case 872:
                                        _ = len(v3)
                                    case 128:
                                        var v90 uint32 = 65535
                                        _ = v90
                                    default:
                                        _ = len(v23)
                                    }
                                }
                                v91 := GenS14{F0: 206, F1: (^(fi26)), F2: 32767}
                                fmt.Println(fi26)
                                _ = v91
                            }
                            _ = v84
                        } else {
                            v4 = (v4 - (v4 * (v4 % ((v4) | 1))))
                            switch "hello" {
                            case "a":
                                _ = len(v3)
                            case "foo bar":
                                if (1 == 2) {
                                    _ = len(v23)
                                    if true {
                                        if true {
                                            v3[(fi26 >> ((v4) & 63))] = fi44
                                            v5[65535] = ("foo bar" + v2)
                                            if false {
                                                var v92 string = ("foo bar" + v2)
                                                _ = v92
                                            }
                                            fmt.Println(fi26, fi26, fi26)
                                            v2 = v2
                                        }
                                        fmt.Println(v4, v4)
                                    }
                                    switch v2 {
                                    case "\n\t":
                                        v2 = ((v2 + (v2 + "foo bar")) + v2)
                                    default:
                                        fmt.Println(v1, v1)
                                        var v93 int64 = (v1 * v1)
                                        _ = v93
                                    }
                                }
                            case "\n\t":
                                v1 = 42
                            }
                            if false {
                                var v94 uint64 = 2
                                fmt.Println(v94, fi44, fi44)
                                v95 := GenS15{F0: 0, F1: genFn0(170, 42, 255)}
                                if (false && true) {
                                    v2 = (v2 + v2)
                                    fmt.Println(v2, fi44)
                                }
                                _ = v94
                                _ = v95
                            } else {
                                var v96 int = fi26
                                if true {
                                    switch v2 {
                                    case "foo bar":
                                        v96 = 128
                                    default:
                                        var v97 uint8 = 0
                                        _ = v97
                                    }
                                    switch (("\n\t" + v2) + ("" + v2)) {
                                    case "hello":
                                        v1 -= (v1 % ((v1) | 1))
                                    case "\n\t":
                                        v4 = v4
                                    }
                                }
                                switch (fi26 - (fi44 * fi26)) {
                                case 684:
                                    v1 = v1
                                case -719:
                                    var v98 int = 42
                                    _ = v98
                                case 1:
                                    var v99 int8 = -1
                                    _ = v99
                                default:
                                    v4 += v4
                                }
                                _ = v96
                            }
                            v4 += 906
                            v100 := GenS16{F0: 2147483647, F1: 42, F2: 255}
                            _ = v100
                        }
                        _ = fi44
                        _ = v45
                    }
                }
            case 544:
                for fi101 := 0; fi101 < 5; fi101++ {
                    v23 = append(v23, fi101)
                    _ = fi101
                }
            }
            v1 += 9223372036854775807
        } else {
            if len(v23) > 0 {
                v23[0] = (^(fi26))
            }
        }
        for fi102 := 0; fi102 < 6; fi102++ {
            v103 := GenS2{F0: 76}
            delete(v5, 112)
            switch (-(fi102)) {
            case 1:
                fmt.Println(v4, fi26, fi102)
                v4 -= (v4 << ((v4) & 63))
                switch fi102 {
                case 9223372036854775807:
                    var v104 string = ("hello" + (v2 + ""))
                    _ = v104
                case 0:
                    v3[(fi102 - fi26)] = 32767
                case 2147483647:
                    var v105 uint8 = 128
                    _ = v105
                default:
                    if true {
                        var v106 int = 177
                        _ = v106
                    } else {
                        v5[266] = "foo bar"
                        _ = len(v23)
                        v4 = (v4 & 255)
                    }
                    v1 += v1
                }
            default:
                v4 = 1
            }
            _ = fi102
            _ = v103
        }
        _ = fi26
    }
    for fi107 := 0; fi107 < 7; fi107++ {
        v108 := v5[1]
        fmt.Println(v4)
        if len(v23) > 0 {
            v23[0] = fi107
        }
        _ = fi107
        _ = v108
    }
    var v109 uint32 = 0
    var v110 int32 = int32(v1)
    switch v2 {
    case "\n\t":
        v1 -= v1
    case "a":
        fmt.Println(v1)
    }
    var v111 int64 = (^(v1))
    var v112 float32 = (-2.25 * 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v109
    _ = v110
    _ = v111
    _ = v112
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
