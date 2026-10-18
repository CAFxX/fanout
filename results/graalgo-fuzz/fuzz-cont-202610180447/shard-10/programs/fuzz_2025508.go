package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 int; F1 uint16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint8; F1 uint; F2 int8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int16; F1 int32; F2 int }
type GenS13 struct { F0 int64; F1 int32; F2 int8 }
func genFn0(p0 int64, p1 int) int8 {
    var v4 float32 = -2.25
    p0 = p0
    fmt.Println(p1, v4, p1)
    v5 := GenS1{F0: (^(p0)), F1: 42, F2: 2}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 100
}
func genFn1(p0 string, p1 uint, p2 uint) int64 {
    p0 = ("hello" + ((("\n\t" + p0) + "hello") + "\n\t"))
    v6 := GenS1{F0: 9223372036854775807, F1: 32767, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[string]uint32)
    var v3 int = 128
    v7 := []int{-2147483648}
    v8 := GenS1{F0: 42, F1: 1, F2: 4294967295}
    fmt.Println(v3, v1)
    v3 += (v3 * v3)
    switch (v3 / (((^(v3))) | 1)) {
    case 0:
        if ((-1 == 1) && (2147483647 == -10)) {
            var v9 bool = true
            v3 = v3
            v10 := GenS1{F0: -1, F1: -494, F2: 504}
            _ = v10
            _ = v9
        }
    case 4294967295:
        v11 := GenS2{F0: 125, F1: 592, F2: 135}
        _ = v11
    default:
        if true {
            fmt.Println(v1)
            if (-2147483648 <= 2147483647) {
                v3 = -159
                v3 = v3
                fmt.Println(v3, v1)
            }
            if true {
                delete(v2, "")
            }
        } else {
            var v12 int32 = -2147483648
            v13 := GenS1{F0: -798, F1: 42, F2: 669}
            for fi14 := 0; fi14 < 13; fi14++ {
                if true {
                    var v15 int16 = -1
                    _ = v15
                } else {
                    if false {
                        if false {
                            if (((((1.5 * (0.0 + (float32(v1) * ((3.14159 * ((float32((1e10 + float64(((0.1 + 0.1) + float32(1.5))))) - -2.25) / 3.14159)) + 0.1)))) * (((((((float32(1.5) / 1e10) + (-2.25 * 1.5)) - float32(float64((1.5 + 0.1)))) / 1.5) / 1.0) * 0.1) - -2.25)) * (1e10 / 3.14159)) <= (3.14159 * 3.14159)) || true) {
                                switch "a" {
                                case "":
                                    switch "\n\t" {
                                    case "\n\t":
                                        for fi16 := 0; fi16 < 1; fi16++ {
                                            if (false && (36 < 4294967295)) {
                                                if (genFn1("a", 826, 42) == genFn1("", 64, 557)) {
                                                    var v17 string = "foo bar"
                                                    var v18 float64 = v1
                                                    fmt.Println(v1, fi14, v3)
                                                    switch 923 {
                                                    case 42:
                                                        v2[(("hello" + (((("a" + "\n\t") + (v17 + (v17 + (v17 + "\n\t")))) + v17) + "")) + v17)] = 4294967295
                                                    case 9223372036854775807:
                                                        v19 := GenS2{F0: 16, F1: 110, F2: 0}
                                                        _ = v19
                                                    default:
                                                        if len(v7) > 0 {
                                                            v7[0] = fi16
                                                        }
                                                    }
                                                    _ = v17
                                                    _ = v18
                                                }
                                            }
                                            v20 := GenS3{F0: v12, F1: 0}
                                            _ = fi16
                                            _ = v20
                                        }
                                    default:
                                        for fi21 := 0; fi21 < 8; fi21++ {
                                            var v22 uint16 = 42
                                            switch "a" {
                                            case "\n\t":
                                                var v23 int = fi14
                                                _ = v23
                                            }
                                            _ = fi21
                                            _ = v22
                                        }
                                    }
                                case "foo bar":
                                    var v24 uint = 4294967295
                                    _ = v24
                                default:
                                    var v25 uint32 = 387
                                    _ = v25
                                }
                                delete(v2, ("" + "a"))
                            } else {
                                v12 -= v12
                                var v26 uint32 = 42
                                _ = v26
                            }
                            fmt.Println(v12, v1)
                            if len(v7) > 0 {
                                v7[0] = 616
                            }
                        }
                    }
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v28 := GenS2{F0: 255, F1: 1, F2: 196}
                        switch fi27 {
                        case 128:
                            v1 = v1
                        case -9223372036854775808:
                            v12 = v12
                            v29 := GenS2{F0: 205, F1: 875, F2: 21}
                            v1 = v1
                            v30 := GenS3{F0: 1, F1: 42}
                            _ = v29
                            _ = v30
                        case -1:
                            var v31 uint64 = 4294967295
                            _ = v31
                        }
                        fmt.Println(v3)
                        _ = fi27
                        _ = v28
                    }
                    fmt.Println(v3, v1, v1)
                }
                for fi32 := 0; fi32 < 8; fi32++ {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    v33 := GenS4{F0: -132, F1: fi14}
                    _ = fi32
                    _ = v33
                }
                var v34 int64 = -9223372036854775808
                _ = fi14
                _ = v34
            }
            _ = v12
            _ = v13
        }
    }
    v35 := GenS5{F0: 1, F1: 100}
    v7 = append(v7, (-(v3)))
    v36 := GenS6{F0: v3, F1: 65535}
    if (42 < 0) {
        if (true || true) {
            v1 = 1e10
            var v37 float64 = v1
            switch ("a" + ("a" + "\n\t")) {
            case "a":
                v37 = v1
            case "\n\t":
                for fi38 := 0; fi38 < 13; fi38++ {
                    for fi39 := 0; fi39 < 17; fi39++ {
                        if true {
                            if (!("" > "\n\t")) {
                                v40 := GenS4{F0: 734, F1: (^(fi39))}
                                if true {
                                    var v41 bool = (false || true)
                                    _ = v41
                                }
                                _ = v40
                            } else {
                                v1 = -2.25
                                v42 := GenS6{F0: fi38, F1: 1}
                                var v43 int = v3
                                _ = v42
                                _ = v43
                            }
                        } else {
                            var v44 string = ("hello" + ("hello" + ("\n\t" + ((("\n\t" + "\n\t") + ("a" + "hello")) + (("hello" + "\n\t") + "")))))
                            _ = v44
                        }
                        switch "hello" {
                        case "":
                            var v45 int64 = genFn1("\n\t", 457, 1)
                            _ = v45
                        case "\n\t":
                            switch (v3 * fi38) {
                            case 2147483647:
                                switch 32768 {
                                case 0:
                                    _ = len(v2)
                                }
                            }
                        case "a":
                            switch fi38 {
                            case 1:
                                v37 = float64(1e10)
                            case 431:
                                v46 := GenS7{F0: 2}
                                _ = v46
                            default:
                                var v47 string = "foo bar"
                                _ = v47
                            }
                        default:
                            v3 += 127
                        }
                        _ = fi39
                    }
                    _ = fi38
                }
            default:
                fmt.Println(v3)
            }
            _ = v37
        } else {
            fmt.Println(v1, v1)
            switch "foo bar" {
            case "hello":
                v1 += v1
            case "\n\t":
                _ = len(v2)
            case "foo bar":
                var v48 int16 = 94
                _ = v48
            default:
                fmt.Println(v3, v1)
            }
        }
        var v49 string = ((("" + "foo bar") + "a") + ("" + (("hello" + "foo bar") + "foo bar")))
        v2[((v49 + (v49 + "hello")) + v49)] = 893
        _ = v49
    } else {
        v1 = ((float64(0.1) * 3.14159) + v1)
        v1 += (3.14159 + (v1 - v1))
    }
    v3 = 1
    if true {
        if true {
            if false {
                v1 -= v1
                var v50 float32 = -2.25
                _ = v50
            }
            if false {
                var v51 int8 = 0
                _ = v51
            } else {
                if false {
                    v52 := GenS1{F0: genFn1("\n\t", 885, 206), F1: 188, F2: 774}
                    _ = v52
                } else {
                    if true {
                        var v53 float64 = 3.14159
                        _ = v53
                    }
                }
            }
            if false {
                for fi54 := 0; fi54 < 6; fi54++ {
                    v55 := GenS8{F0: 448}
                    v1 = v1
                    _ = fi54
                    _ = v55
                }
                for fi56 := 0; fi56 < 2; fi56++ {
                    var v57 uint8 = 128
                    var v58 uint16 = 414
                    if true {
                        if false {
                            v3 = (fi56 % (((v3 % (((v3 << ((1) & 63))) | 1))) | 1))
                        } else {
                            _ = len(v2)
                        }
                        var v59 uint32 = 42
                        _ = v59
                    }
                    _ = fi56
                    _ = v57
                    _ = v58
                }
                v1 = (v1 / ((0.0 + v1) * v1))
            } else {
                if (-235 > -394) {
                    var v60 int = v3
                    v2[("" + "hello")] = 65535
                    _ = v60
                } else {
                    var v61 uint = 1
                    _ = v61
                }
                v62 := GenS7{F0: 65535}
                _ = v62
            }
        } else {
            var v63 int32 = 1
            _ = v63
        }
        if false {
            for fi64 := 0; fi64 < 11; fi64++ {
                var v65 int64 = -334
                if ((v65 * v65) <= (v65 >> ((2) & 63))) {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    v1 -= (0.1 * float64(float32(((v1 + ((float64(-2.25) * float64(float32(float64(3.14159)))) + (v1 / float64(-2.25)))) * (1.5 * v1)))))
                    v1 += v1
                } else {
                    switch 42 {
                    case 128:
                        var v66 float64 = v1
                        _ = v66
                    case 1:
                        for fi67 := 0; fi67 < 13; fi67++ {
                            fmt.Println(v3, fi64)
                            _ = fi67
                        }
                    case 16:
                        _ = len(v7)
                    default:
                        if ((255 <= 2) || true) {
                            var v68 uint64 = 120
                            _ = v68
                        }
                    }
                }
                _ = fi64
                _ = v65
            }
            switch 32768 {
            case 233:
                if (4294967295 <= 0) {
                    if false {
                        if (-2.25 >= 1e10) {
                            if false {
                                v1 = v1
                            }
                            v1 = v1
                            var v69 uint64 = 0
                            _ = v69
                        } else {
                            var v70 int8 = 13
                            v3 = v3
                            fmt.Println(v1, v1)
                            _ = v70
                        }
                        var v71 int = v3
                        _ = v71
                    } else {
                        v1 = (v1 - (v1 * -2.25))
                        switch "foo bar" {
                        case "\n\t":
                            for fi72 := 0; fi72 < 9; fi72++ {
                                _ = len(v7)
                                fmt.Println(v1, v3)
                                v73 := GenS3{F0: 0, F1: 32767}
                                _ = fi72
                                _ = v73
                            }
                        case "":
                            v7 = append(v7, v3)
                        case "foo bar":
                            var v74 uint16 = 255
                            _ = v74
                        }
                        if len(v7) > 0 {
                            v7[0] = -1
                        }
                    }
                    var v75 string = "a"
                    for fi76 := 0; fi76 < 16; fi76++ {
                        switch 128 {
                        case 967:
                            fmt.Println(v75, fi76)
                        case -597:
                            delete(v2, v75)
                        case 32768:
                            var v77 float32 = 0.1
                            _ = v77
                        default:
                            v1 = 0.0
                        }
                        _ = fi76
                    }
                    _ = v75
                }
            }
            if false {
                v78 := GenS9{F0: 516}
                switch "\n\t" {
                case "foo bar":
                    var v79 int16 = 585
                    _ = v79
                case "":
                    switch "foo bar" {
                    case "":
                        for fi80 := 0; fi80 < 13; fi80++ {
                            var v81 int = (v3 / ((v3) | 1))
                            switch ((("hello" + ((("a" + ("a" + "")) + "") + ("a" + "hello"))) + "hello") + ("\n\t" + "foo bar")) {
                            case "\n\t":
                                if false {
                                    _ = len(v2)
                                    var v82 uint64 = 4294967295
                                    _ = v82
                                }
                            case "hello":
                                v3 = v81
                            default:
                                var v83 uint16 = 0
                                _ = v83
                            }
                            _ = fi80
                            _ = v81
                        }
                    }
                }
                _ = v78
            } else {
                var v84 int64 = -941
                _ = v84
            }
        } else {
            v3 = v3
        }
        v1 += (v1 / 1.5)
    } else {
        v85 := GenS10{F0: 128, F1: 2, F2: -16}
        var v86 int8 = 42
        for fi87 := 0; fi87 < 4; fi87++ {
            for fi88 := 0; fi88 < 15; fi88++ {
                v89 := v2["hello"]
                _ = fi88
                _ = v89
            }
            _ = fi87
        }
        _ = v85
        _ = v86
    }
    for fi90 := 0; fi90 < 1; fi90++ {
        fmt.Println(v3, v3, v1)
        fmt.Println(v1)
        if (3.14159 <= ((v1 - (v1 * v1)) * v1)) {
            if (-1 != -66) {
                v91 := GenS3{F0: 1, F1: 1}
                _ = v91
            }
            switch -52 {
            case -1:
                var v92 bool = (true || (false || (2 >= 724)))
                _ = v92
            }
        } else {
            if false {
                var v93 string = "hello"
                if (197 <= -1) {
                    for fi94 := 0; fi94 < 20; fi94++ {
                        v93 = v93
                        v3 -= (fi90 << ((1) & 63))
                        v3 -= v3
                        v93 = (("a" + v93) + "")
                        _ = fi94
                    }
                } else {
                    fmt.Println(v93, v3)
                    v95 := GenS9{F0: 0}
                    _ = v95
                }
                var v96 uint32 = 151
                _ = v93
                _ = v96
            }
            var v97 string = ""
            var v98 float32 = (1.5 * (0.0 * 1e10))
            _ = v97
            _ = v98
        }
        _ = fi90
    }
    fmt.Println(v3)
    switch (("hello" + "") + ("hello" + ("a" + "a"))) {
    case "foo bar":
        v1 -= 1e10
    case "":
        if ((1 == 221) || ("foo bar" <= "")) {
            if true {
                if (v3 >= (-(v3))) {
                    v3 -= v3
                    v99 := GenS11{F0: 1}
                    _ = v99
                } else {
                    fmt.Println(v1, v3)
                    for fi100 := 0; fi100 < 14; fi100++ {
                        if (!(-606 <= 9223372036854775807)) {
                            switch ("a" + ("foo bar" + "a")) {
                            case "a":
                                v3 += (v3 - (v3 * (fi100 / ((v3) | 1))))
                            case "\n\t":
                                var v101 uint16 = 32767
                                _ = v101
                            }
                            for fi102 := 0; fi102 < 14; fi102++ {
                                if true {
                                    switch fi102 {
                                    case 0:
                                        if (!true) {
                                            fmt.Println(v1, fi100, v1)
                                            var v103 string = ("foo bar" + "")
                                            var v104 bool = true
                                            _ = v103
                                            _ = v104
                                        } else {
                                            var v105 float32 = (0.0 - 1e10)
                                            _ = v105
                                        }
                                    case -770:
                                        v106 := GenS6{F0: v3, F1: 611}
                                        _ = v106
                                    }
                                    fmt.Println(fi102)
                                    var v107 int32 = -2147483648
                                    _ = v107
                                }
                                v108 := GenS1{F0: 0, F1: 0, F2: 42}
                                v109 := GenS12{F0: 562, F1: 2147483647, F2: fi102}
                                switch fi100 {
                                case 9223372036854775807:
                                    v2["hello"] = 614
                                case -442:
                                    v7 = append(v7, v3)
                                default:
                                    v7 = append(v7, (fi102 & fi100))
                                }
                                _ = fi102
                                _ = v108
                                _ = v109
                            }
                        }
                        _ = fi100
                    }
                }
            }
        }
    }
    fmt.Println(v1, v3)
    v1 += v1
    v110 := GenS13{F0: genFn1("hello", 167, 472), F1: 1, F2: 100}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v110
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v7
    _ = v8
}
