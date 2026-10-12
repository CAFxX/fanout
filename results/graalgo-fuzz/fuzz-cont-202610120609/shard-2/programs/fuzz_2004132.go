package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int64 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 bool; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int16; F1 uint64; F2 uint }
type GenS9 struct { F0 uint; F1 bool; F2 int }
type GenS10 struct { F0 string; F1 uint }
type GenS11 struct { F0 bool; F1 uint16 }
type GenS12 struct { F0 int16; F1 uint8 }
type GenS13 struct { F0 uint; F1 int }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 uint16; F1 int32; F2 int }
type GenS16 struct { F0 int8; F1 string; F2 int8 }
func genFn0(p0 bool, p1 string) string {
    p1 = p1
    var v7 string = p1
    var v8 uint8 = 128
    var v9 uint64 = 42
    fmt.Println(p0, v8)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (v7 + p1)
}
func genFn1() int {
    var v10 int8 = 42
    v10 -= (v10 - -87)
    v11 := []int{-943, -720, 32767}
    v12 := GenS1{F0: -128, F1: 0, F2: 0}
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int16 = 32767
        _ = fi13
        _ = v14
    }
    _ = v10
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 float64 = 1.5
    var v2 int = 0
    v3 := []int{}
    var v4 uint8 = 255
    v5 := make(map[int]uint32)
    v6 := make(map[string]string)
    v15 := GenS1{F0: 0, F1: 9223372036854775807, F2: 9223372036854775807}
    fmt.Println(v2, v1)
    v16 := GenS1{F0: 42, F1: 350, F2: int64(v2)}
    v2 = v2
    v17 := GenS1{F0: -1, F1: -9223372036854775808, F2: 163}
    fmt.Println(v2)
    for fi18 := 0; fi18 < 6; fi18++ {
        switch 1 {
        case 471:
            if false {
                var v19 int16 = -32768
                v19 -= (v19 * (v19 ^ 42))
                _ = v19
            } else {
                if true {
                    v4 = (v4 | (^(v4)))
                } else {
                    v20 := GenS2{F0: 2147483647, F1: 608}
                    v2 = v2
                    _ = v20
                }
                if true {
                    v1 -= (v1 + float64(((-2.25 + 3.14159) - 0.1)))
                } else {
                    switch (v2 >> ((42) & 63)) {
                    case -151:
                        v2 += (fi18 - genFn1())
                    case -1:
                        if false {
                            v1 = 1.5
                            switch 127 {
                            case -1:
                                var v21 float32 = -2.25
                                _ = v21
                            }
                        }
                    case 2147483647:
                        if (("\n\t" + ("\n\t" + ("hello" + ""))) <= "hello") {
                            switch ("\n\t" + "a") {
                            case "":
                                switch "a" {
                                case "":
                                    fmt.Println(v2, v2, v2)
                                case "hello":
                                    for fi22 := 0; fi22 < 14; fi22++ {
                                        if true {
                                            _ = len(v3)
                                            v23 := GenS3{F0: 18446744073709551615}
                                            v24 := GenS4{F0: -1, F1: (42 == 1), F2: 2}
                                            v3 = append(v3, (-(fi22)))
                                            var v25 float64 = v1
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        } else {
                                            var v26 int = (fi22 + fi22)
                                            if (-2147483648 != -322) {
                                                switch ("\n\t" + (("" + "foo bar") + "")) {
                                                case "":
                                                    if (false && (!(!((!(false && (255 == 4294967295))) || (!(1.5 == float32(-2.25))))))) {
                                                        switch (("foo bar" + ("" + (((("hello" + "a") + ("a" + (("" + "hello") + "\n\t"))) + "") + (("foo bar" + "foo bar") + (("" + "a") + (("foo bar" + "foo bar") + "hello")))))) + ("" + "a")) {
                                                        case "":
                                                            var v27 uint16 = 255
                                                            _ = v27
                                                        case "\n\t":
                                                            switch int(v4) {
                                                            case 2147483647:
                                                                if false {
                                                                    var v28 uint = 255
                                                                    if ((!(false && true)) && true) {
                                                                        if false {
                                                                            switch fi18 {
                                                                            case 2147483647:
                                                                                var v29 int32 = 733
                                                                                _ = v29
                                                                            case 128:
                                                                                var v30 uint = v28
                                                                                _ = v30
                                                                            }
                                                                            v1 = 0.0
                                                                            v1 += float64(1.5)
                                                                        } else {
                                                                            fmt.Println(fi18, v26, fi22)
                                                                            v31 := GenS4{F0: 0, F1: false, F2: 0}
                                                                            var v32 int16 = 32767
                                                                            switch v2 {
                                                                            case 2147483647:
                                                                                var v33 int16 = (v32 * v32)
                                                                                _ = v33
                                                                            default:
                                                                                var v34 uint = (v28 / ((v28) | 1))
                                                                                _ = v34
                                                                            }
                                                                            v35 := v5[(fi22 | v2)]
                                                                            _ = v31
                                                                            _ = v32
                                                                            _ = v35
                                                                        }
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 9223372036854775807
                                                                        }
                                                                    }
                                                                    _ = v28
                                                                }
                                                            case 63:
                                                                var v36 float32 = 1e10
                                                                _ = v36
                                                            case -9223372036854775808:
                                                                v37 := GenS5{F0: (!true)}
                                                                _ = v37
                                                            }
                                                        case "foo bar":
                                                            v1 = v1
                                                        default:
                                                            v38 := GenS3{F0: 350}
                                                            _ = v38
                                                        }
                                                        if (false || ((false && false) || true)) {
                                                            v5[fi18] = 395
                                                        } else {
                                                            var v39 bool = ((-351 >= fi18) && (!(!(true || (2 > 0)))))
                                                            fmt.Println(v26)
                                                            fmt.Println(fi18, fi18, v4)
                                                            _ = v39
                                                        }
                                                        v40 := GenS5{F0: ("\n\t" != "\n\t")}
                                                        _ = v40
                                                    } else {
                                                        var v41 int8 = -128
                                                        fmt.Println(v2, v4, v4)
                                                        var v42 int32 = 771
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                case "hello":
                                                    var v43 float32 = ((((3.14159 / 1e10) - (1.5 + 0.1)) * 0.1) + (1.5 + 1e10))
                                                    _ = v43
                                                case "foo bar":
                                                    var v44 string = (((("\n\t" + "foo bar") + "foo bar") + "hello") + "foo bar")
                                                    _ = v44
                                                default:
                                                    v2 -= v2
                                                }
                                                v45 := v6["hello"]
                                                switch v45 {
                                                case "foo bar":
                                                    v46 := GenS6{F0: 2147483647, F1: v2}
                                                    _ = v46
                                                default:
                                                    var v47 int = fi22
                                                    if false {
                                                        v48 := v6[v45]
                                                        fmt.Println(v4)
                                                        v2 -= (fi22 & v47)
                                                        _ = v48
                                                    }
                                                    if (v45 < "a") {
                                                        var v49 string = ((v45 + v45) + v45)
                                                        var v50 string = "foo bar"
                                                        _ = v49
                                                        _ = v50
                                                    } else {
                                                        v51 := GenS7{F0: 137, F1: (v26 ^ fi18)}
                                                        v52 := v6[v45]
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                    v53 := v6[v45]
                                                    v47 -= (v47 + 42)
                                                    _ = v47
                                                    _ = v53
                                                }
                                                _ = v45
                                            }
                                            v54 := GenS8{F0: 35, F1: 7, F2: 18446744073709551615}
                                            _ = v26
                                            _ = v54
                                        }
                                        _ = fi22
                                    }
                                case "\n\t":
                                    var v55 string = ("\n\t" + "hello")
                                    _ = v55
                                default:
                                    for fi56 := 0; fi56 < 7; fi56++ {
                                        v4 -= (v4 * (^(v4)))
                                        _ = fi56
                                    }
                                }
                            }
                            var v57 uint16 = 42
                            if false {
                                if (42 > 18446744073709551615) {
                                    var v58 uint8 = v4
                                    if (!(!(692 != v57))) {
                                        if ((float64((((1e10 / 1e10) + -2.25) / ((3.14159 - float32(((1e10 + ((float64((1e10 - 1.5)) * v1) * float64(-2.25))) - (float64(0.0) + v1)))) * (float32(v1) / float32(0.1))))) + (((3.14159 / float64((0.1 - (0.0 / 3.14159)))) + 3.14159) * -2.25)) > (((1.5 * (-2.25 * v1)) / (3.14159 + v1)) - float64(1.5))) {
                                            if true {
                                                var v59 uint = 18446744073709551615
                                                switch "" {
                                                case "a":
                                                    if (-509 != genFn1()) {
                                                        v4 = v4
                                                        for fi60 := 0; fi60 < 10; fi60++ {
                                                            var v61 int16 = 42
                                                            v57 = v57
                                                            if ((v58 + (v58 ^ 2)) < 42) {
                                                                if ((float64(3.14159) / ((v1 * 3.14159) / 1.0)) <= 1.5) {
                                                                    v62 := GenS6{F0: -2147483648, F1: (fi18 - (^(v2)))}
                                                                    var v63 int = genFn1()
                                                                    _ = v62
                                                                    _ = v63
                                                                }
                                                            } else {
                                                                v1 += 1e10
                                                                v2 = v2
                                                            }
                                                            if true {
                                                                var v64 float32 = float32(((1e10 - 1.5) + ((v1 - v1) * 0.1)))
                                                                if (v61 > v61) {
                                                                    var v65 bool = (true && ((((v64 * (3.14159 + ((float32(float64(0.0)) + v64) / float32(-2.25)))) - -2.25) / v64) != ((v64 + -2.25) + v64)))
                                                                    _ = v65
                                                                }
                                                                _ = v64
                                                            } else {
                                                                if true {
                                                                    v66 := v5[(fi18 / ((2147483647) | 1))]
                                                                    v59 += (v59 / ((4294967295) | 1))
                                                                    if true {
                                                                        v67 := GenS7{F0: 2147483648, F1: (fi18 | 128)}
                                                                        if false {
                                                                            var v68 float32 = ((1.5 * (1.5 + (-2.25 * -2.25))) * (0.0 / 0.1))
                                                                            _ = v68
                                                                        }
                                                                        if (!true) {
                                                                            v3 = append(v3, int(v4))
                                                                            switch fi18 {
                                                                            case 501:
                                                                                if false {
                                                                                    if (v59 != v59) {
                                                                                        var v69 float64 = float64(1e10)
                                                                                        var v70 uint64 = 4294967295
                                                                                        _ = v69
                                                                                        _ = v70
                                                                                    }
                                                                                }
                                                                            case -9223372036854775808:
                                                                                v71 := GenS8{F0: (v61 * 42), F1: 919, F2: 4294967295}
                                                                                _ = v71
                                                                            default:
                                                                                v72 := GenS1{F0: 127, F1: 9223372036854775807, F2: int64(v66)}
                                                                                _ = v72
                                                                            }
                                                                            fmt.Println(v61, v2)
                                                                            var v73 uint8 = (v4 >> ((uint(v66)) & 63))
                                                                            if (v66 < v66) {
                                                                                if (542 >= 759) {
                                                                                    switch (-(fi60)) {
                                                                                    case -200:
                                                                                        fmt.Println(fi60)
                                                                                    case 607:
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            var v74 string = "a"
                                                                                            _ = v74
                                                                                        case "a":
                                                                                            v73 += (v73 ^ v58)
                                                                                        case "foo bar":
                                                                                            v73 -= uint8(v66)
                                                                                        default:
                                                                                            if (4294967295 > 0) {
                                                                                                if true {
                                                                                                    v75 := GenS9{F0: (v59 / ((uint(v66)) | 1)), F1: true, F2: (-(fi60))}
                                                                                                    if ((-2.25 * 0.0) > 3.14159) {
                                                                                                        switch (v2 / ((-341) | 1)) {
                                                                                                        case 909:
                                                                                                            fmt.Println(v73, v59, fi18)
                                                                                                        case 333:
                                                                                                            var v76 bool = (false || (42 != 1))
                                                                                                            _ = v76
                                                                                                        }
                                                                                                        fmt.Println(v1, v59, v59)
                                                                                                        var v77 int8 = 40
                                                                                                        _ = v77
                                                                                                    }
                                                                                                    var v78 int8 = 35
                                                                                                    var v79 int64 = int64(v66)
                                                                                                    fmt.Println(v2, v59, fi18)
                                                                                                    _ = v75
                                                                                                    _ = v78
                                                                                                    _ = v79
                                                                                                }
                                                                                                if (278 != 9223372036854775807) {
                                                                                                    fmt.Println(v73)
                                                                                                    if true {
                                                                                                        if ((v73 >> (((v59 | 255)) & 63)) >= 2) {
                                                                                                            fmt.Println(v66, fi60, v61)
                                                                                                            switch ("\n\t" + "hello") {
                                                                                                            case "":
                                                                                                                v80 := GenS10{F0: ("\n\t" + ("a" + "hello")), F1: (v59 << (((v59 & v59)) & 63))}
                                                                                                                _ = v80
                                                                                                            case "\n\t":
                                                                                                                var v81 string = ("hello" + "hello")
                                                                                                                _ = v81
                                                                                                            default:
                                                                                                                var v82 uint8 = 154
                                                                                                                _ = v82
                                                                                                            }
                                                                                                        }
                                                                                                        v83 := GenS11{F0: (("" + (("\n\t" + (("\n\t" + "hello") + "\n\t")) + "\n\t")) <= "foo bar"), F1: v57}
                                                                                                        v84 := GenS10{F0: "\n\t", F1: 255}
                                                                                                        _ = v83
                                                                                                        _ = v84
                                                                                                    } else {
                                                                                                        v57 -= v57
                                                                                                    }
                                                                                                    fmt.Println(v73, v1, v59)
                                                                                                    _ = len(v3)
                                                                                                    _ = len(v6)
                                                                                                    if (v59 <= (v59 / ((v59) | 1))) {
                                                                                                        v66 += (^(v66))
                                                                                                        v66 = v66
                                                                                                        var v85 uint16 = 0
                                                                                                        v86 := GenS4{F0: 9223372036854775807, F1: ((^(v61)) >= v61), F2: 477}
                                                                                                        _ = v85
                                                                                                        _ = v86
                                                                                                    }
                                                                                                }
                                                                                                switch "\n\t" {
                                                                                                case "hello":
                                                                                                    fmt.Println(v73)
                                                                                                case "\n\t":
                                                                                                    v3 = append(v3, (fi60 >> ((4294967295) & 63)))
                                                                                                case "":
                                                                                                    var v87 int64 = 717
                                                                                                    _ = v87
                                                                                                default:
                                                                                                    var v88 float64 = v1
                                                                                                    _ = v88
                                                                                                }
                                                                                                var v89 uint = 18446744073709551615
                                                                                                var v90 int16 = v61
                                                                                                v91 := GenS1{F0: -42, F1: 410, F2: 9223372036854775807}
                                                                                                switch (v2 - (fi60 ^ 2147483647)) {
                                                                                                case 387:
                                                                                                    var v92 uint64 = 2
                                                                                                    _ = v92
                                                                                                case 42:
                                                                                                    v89 = v89
                                                                                                case -9223372036854775808:
                                                                                                    delete(v6, "")
                                                                                                default:
                                                                                                    switch (("hello" + "foo bar") + ("hello" + "\n\t")) {
                                                                                                    case "hello":
                                                                                                        switch (v2 >> ((255) & 63)) {
                                                                                                        case 4294967295:
                                                                                                            var v93 uint64 = 4294967295
                                                                                                            _ = v93
                                                                                                        case 1:
                                                                                                            v94 := GenS11{F0: true, F1: 32767}
                                                                                                            _ = v94
                                                                                                        case 0:
                                                                                                            if false {
                                                                                                                fmt.Println(fi60, v66, v1)
                                                                                                                fmt.Println(v59, v59, v90)
                                                                                                                v3 = append(v3, fi18)
                                                                                                            }
                                                                                                        }
                                                                                                    case "":
                                                                                                        fmt.Println(v66, v66)
                                                                                                    default:
                                                                                                        var v95 int16 = (-(v61))
                                                                                                        _ = v95
                                                                                                    }
                                                                                                }
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = -1
                                                                                                }
                                                                                                var v96 int16 = v61
                                                                                                v97 := GenS12{F0: (v90 / (((v96 / ((int16(v59)) | 1))) | 1)), F1: v4}
                                                                                                _ = len(v3)
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                                v98 := GenS13{F0: 817, F1: -867}
                                                                                                fmt.Println(v57)
                                                                                                switch (v2 | int(v66)) {
                                                                                                case 734:
                                                                                                    var v99 int8 = 0
                                                                                                    _ = v99
                                                                                                case 9223372036854775807:
                                                                                                    var v100 float64 = (1e10 * v1)
                                                                                                    _ = v100
                                                                                                case 1:
                                                                                                    var v101 int32 = -788
                                                                                                    v3 = append(v3, 752)
                                                                                                    if false {
                                                                                                        v102 := GenS9{F0: v89, F1: false, F2: fi18}
                                                                                                        if (true || ((v96 & (v96 & -127)) != v96)) {
                                                                                                            delete(v5, fi60)
                                                                                                        } else {
                                                                                                            var v103 int32 = v101
                                                                                                            _ = v103
                                                                                                        }
                                                                                                        if true {
                                                                                                            fmt.Println(fi60, v89)
                                                                                                            v66 = v66
                                                                                                        }
                                                                                                        _ = v102
                                                                                                    }
                                                                                                    _ = v101
                                                                                                default:
                                                                                                    var v104 int32 = 42
                                                                                                    _ = v104
                                                                                                }
                                                                                                fmt.Println(fi60)
                                                                                                _ = v89
                                                                                                _ = v90
                                                                                                _ = v91
                                                                                                _ = v96
                                                                                                _ = v97
                                                                                                _ = v98
                                                                                            }
                                                                                        }
                                                                                    case 42:
                                                                                        v61 = v61
                                                                                    }
                                                                                }
                                                                                if (true || (0 <= 0)) {
                                                                                    if (v59 <= v59) {
                                                                                        var v105 int32 = 1
                                                                                        v57 += 1
                                                                                        _ = v105
                                                                                    } else {
                                                                                        if true {
                                                                                            v3 = append(v3, (fi18 << ((uint(v57)) & 63)))
                                                                                        } else {
                                                                                            v1 -= -2.25
                                                                                        }
                                                                                        v106 := v6[(("a" + "foo bar") + "\n\t")]
                                                                                        _ = v106
                                                                                    }
                                                                                    fmt.Println(v57)
                                                                                }
                                                                                var v107 float32 = 1.5
                                                                                _ = v107
                                                                            }
                                                                            _ = v73
                                                                        }
                                                                        _ = v67
                                                                    }
                                                                    fmt.Println(v58)
                                                                    _ = v66
                                                                }
                                                                delete(v5, (v2 << ((v59) & 63)))
                                                                fmt.Println(fi18, fi60, v58)
                                                                var v108 uint8 = v58
                                                                if len(v3) > 0 {
                                                                    v3[0] = (v2 | genFn1())
                                                                }
                                                                fmt.Println(v4)
                                                                v58 += (v58 | v108)
                                                                _ = v108
                                                            }
                                                            var v109 int16 = v61
                                                            v110 := GenS4{F0: 42, F1: false, F2: 42}
                                                            v58 = v4
                                                            var v111 int32 = -2147483648
                                                            _ = fi60
                                                            _ = v109
                                                            _ = v110
                                                            _ = v111
                                                            _ = v61
                                                        }
                                                        fmt.Println(v2, v58)
                                                    } else {
                                                        for fi112 := 0; fi112 < 17; fi112++ {
                                                            v113 := GenS6{F0: -1, F1: (v2 | (fi112 << ((3) & 63)))}
                                                            v114 := GenS7{F0: 4294967295, F1: fi18}
                                                            _ = fi112
                                                            _ = v113
                                                            _ = v114
                                                        }
                                                    }
                                                case "":
                                                    if false {
                                                        v6[""] = "hello"
                                                    }
                                                default:
                                                    v1 -= ((v1 / 1.5) * 0.0)
                                                }
                                                switch genFn1() {
                                                case 2147483647:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case 1:
                                                    v1 = 0.1
                                                default:
                                                    if ((v1 == 3.14159) || (-1 < 1)) {
                                                        v6[""] = "a"
                                                    } else {
                                                        switch (("\n\t" + "") + ("foo bar" + (("" + "\n\t") + "hello"))) {
                                                        case "":
                                                            for fi115 := 0; fi115 < 16; fi115++ {
                                                                fmt.Println(fi115, v4)
                                                                var v116 uint16 = (v57 >> ((v59) & 63))
                                                                v59 += v59
                                                                _ = fi115
                                                                _ = v116
                                                            }
                                                        case "\n\t":
                                                            var v117 float32 = float32(v1)
                                                            _ = v117
                                                        case "hello":
                                                            var v118 int64 = -631
                                                            _ = v118
                                                        default:
                                                            var v119 bool = (false || (4294967295 >= 42))
                                                            _ = v119
                                                        }
                                                        var v120 int8 = -128
                                                        _ = v120
                                                    }
                                                }
                                                _ = v59
                                            }
                                            for fi121 := 0; fi121 < 4; fi121++ {
                                                var v122 uint16 = v57
                                                _ = fi121
                                                _ = v122
                                            }
                                        } else {
                                            fmt.Println(v58, v1)
                                        }
                                        var v123 uint64 = 686
                                        _ = v123
                                    } else {
                                        v57 += v57
                                        switch (fi18 | fi18) {
                                        case 42:
                                            v124 := GenS14{F0: -128}
                                            _ = v124
                                        }
                                        v57 = (v57 % ((v57) | 1))
                                    }
                                    _ = v58
                                }
                                v57 = 255
                                switch fi18 {
                                case -2147483648:
                                    for fi125 := 0; fi125 < 15; fi125++ {
                                        delete(v5, (v2 | genFn1()))
                                        _ = fi125
                                    }
                                }
                            }
                            _ = v57
                        } else {
                            delete(v6, ("\n\t" + "a"))
                            if false {
                                var v126 uint16 = 32767
                                var v127 int = genFn1()
                                _ = v126
                                _ = v127
                            } else {
                                var v128 uint16 = 22
                                _ = len(v5)
                                var v129 int = 2147483647
                                _ = v128
                                _ = v129
                            }
                            _ = len(v3)
                        }
                    default:
                        var v130 int32 = 644
                        _ = v130
                    }
                }
                var v131 uint8 = v4
                _ = v131
            }
        default:
            var v132 int64 = 1
            _ = v132
        }
        _ = fi18
    }
    if true {
        if true {
            v133 := GenS15{F0: 0, F1: 294, F2: 839}
            v134 := GenS7{F0: 1, F1: (v2 / (((v2 * v2)) | 1))}
            _ = v133
            _ = v134
        } else {
            fmt.Println(v4, v1)
        }
    }
    if false {
        var v135 int16 = 32767
        var v136 uint64 = 638
        _ = v135
        _ = v136
    }
    v137 := GenS16{F0: -1, F1: "a", F2: 42}
    fmt.Println(v2, v4)
    var v138 float64 = v1
    var v139 float64 = v138
    v2 = (v2 % ((v2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v138:", v138)
    fmt.Println("v139:", v139)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v137
    _ = v138
    _ = v139
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
