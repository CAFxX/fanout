package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint }
type GenS5 struct { F0 int32; F1 bool; F2 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 uint8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint; F1 int8 }
type GenS13 struct { F0 int32; F1 uint16; F2 int }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint; F1 int; F2 int32 }
func genFn0(p0 uint8) int8 {
    var v5 uint64 = 378
    p0 = 196
    switch 32767 {
    case 909:
        if ((^(p0)) == 215) {
            v5 = v5
            var v6 uint8 = p0
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := []int{fi7, int(v5)}
        p0 = 193
        v5 = (^(v5))
        _ = fi7
        _ = v8
    }
    fmt.Println(p0, v5)
    _ = p0
    _ = v5
    return 42
}
func genFn1() uint8 {
    v9 := GenS1{F0: -128, F1: 591}
    var v10 int64 = 1
    if false {
        v11 := []int{-9223372036854775808, 32768, 42, 622}
        _ = v11
    }
    v10 -= (-(v10))
    var v12 uint16 = 42
    _ = v10
    _ = v12
    _ = v9
    return 0
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]uint32)
    var v3 uint8 = v1
    var v4 uint8 = v3
    if true {
        var v13 float32 = float32(1.5)
        for fi14 := 0; fi14 < 7; fi14++ {
            v2["a"] = 955
            if (genFn1() == v4) {
                fmt.Println(v4, v1, fi14)
                v1 = v4
                var v15 float64 = (((1e10 * (float64((v13 / v13)) - ((-2.25 - 1e10) * 0.1))) * 1e10) + 3.14159)
                _ = v15
            }
            _ = fi14
        }
        if false {
            var v16 uint16 = 65535
            if (42 != -407) {
                var v17 int16 = 1
                v18 := GenS2{F0: 42}
                v16 = 434
                _ = v17
                _ = v18
            }
            for fi19 := 0; fi19 < 16; fi19++ {
                if (false && true) {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        v21 := GenS3{F0: 42}
                        v3 = 127
                        if true {
                            switch "foo bar" {
                            case "a":
                                _ = len(v2)
                            case "foo bar":
                                var v22 uint16 = 165
                                _ = v22
                            default:
                                delete(v2, "hello")
                            }
                            var v23 int16 = 1000
                            var v24 uint8 = v4
                            delete(v2, "hello")
                            _ = v23
                            _ = v24
                        } else {
                            var v25 float64 = ((0.1 * (1.5 * 1e10)) * 0.1)
                            var v26 int32 = 42
                            if (!true) {
                                if (!true) {
                                    v27 := GenS2{F0: 919}
                                    v28 := GenS3{F0: 92}
                                    _ = v27
                                    _ = v28
                                } else {
                                    var v29 float64 = v25
                                    switch -233 {
                                    case 0:
                                        if (-2147483648 < (^(v26))) {
                                            if false {
                                                var v30 int32 = v26
                                                _ = v30
                                            }
                                        } else {
                                            var v31 int8 = int8(v4)
                                            _ = v31
                                        }
                                    case -9223372036854775808:
                                        var v32 uint8 = (v3 * v3)
                                        _ = v32
                                    case -1:
                                        v33 := []int{128, (fi20 + fi20), 32767, 128}
                                        _ = v33
                                    }
                                    v34 := []int{}
                                    _ = v29
                                    _ = v34
                                }
                                var v35 int8 = 1
                                v36 := GenS1{F0: (^(v35)), F1: (^(v16))}
                                _ = v35
                                _ = v36
                            } else {
                                var v37 int16 = -1
                                v38 := GenS3{F0: 4294967295}
                                _ = v37
                                _ = v38
                            }
                            _ = v25
                            _ = v26
                        }
                        _ = fi20
                        _ = v21
                    }
                }
                _ = fi19
            }
            _ = v16
        }
        _ = v13
    } else {
        if false {
            for fi39 := 0; fi39 < 4; fi39++ {
                fmt.Println(v3, v1, v3)
                for fi40 := 0; fi40 < 14; fi40++ {
                    v3 += (^(v3))
                    switch ("" + ("\n\t" + "\n\t")) {
                    case "hello":
                        v41 := GenS4{F0: false, F1: v3, F2: 42}
                        _ = v41
                    case "\n\t":
                        v42 := GenS1{F0: 100, F1: 849}
                        _ = v42
                    }
                    switch "a" {
                    case "hello":
                        if ((-1 >= 42) || true) {
                            var v43 int32 = 1
                            _ = v43
                        } else {
                            if true {
                                fmt.Println(fi39)
                                if false {
                                    v3 = v1
                                    var v44 uint32 = 2147483648
                                    _ = v44
                                }
                                if ((-2.25 + (3.14159 / 1.0)) < 3.14159) {
                                    var v45 int8 = -19
                                    v45 = -1
                                    if true {
                                        var v46 int8 = (v45 / ((genFn0(v1)) | 1))
                                        v47 := GenS2{F0: 957}
                                        _ = v46
                                        _ = v47
                                    }
                                    _ = v45
                                } else {
                                    switch fi39 {
                                    case -9223372036854775808:
                                        switch (-(fi39)) {
                                        case 0:
                                            fmt.Println(v3, fi39, fi40)
                                        case -9223372036854775808:
                                            var v48 int64 = -9223372036854775808
                                            _ = v48
                                        case 9223372036854775807:
                                            v49 := GenS4{F0: (true || false), F1: v1, F2: 483}
                                            _ = v49
                                        }
                                    default:
                                        switch (("a" + "") + ("a" + "foo bar")) {
                                        case "":
                                            fmt.Println(fi40)
                                        case "\n\t":
                                            v1 -= (v4 ^ (v1 - 110))
                                            if (!(!true)) {
                                                if ("hello" != ((("a" + ("hello" + "a")) + ("\n\t" + "foo bar")) + "\n\t")) {
                                                    var v50 string = ("\n\t" + "")
                                                    switch fi40 {
                                                    case -2147483648:
                                                        v50 = ("foo bar" + "\n\t")
                                                    default:
                                                        v51 := []int{fi40, (fi40 << ((833) & 63))}
                                                        _ = v51
                                                    }
                                                    if (v50 >= "\n\t") {
                                                        if true {
                                                            var v52 float64 = 1.5
                                                            v50 = ""
                                                            var v53 int8 = -59
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                        switch 1 {
                                                        case 445:
                                                            if (255 == 42) {
                                                                v1 += (v1 % ((v4) | 1))
                                                                fmt.Println(v50, v4, fi40)
                                                            } else {
                                                                switch (^(fi39)) {
                                                                case 128:
                                                                    v54 := GenS5{F0: -2147483648, F1: (554 != 0), F2: 0}
                                                                    _ = v54
                                                                }
                                                                var v55 int64 = -814
                                                                v3 += (v1 << ((733) & 63))
                                                                _ = v55
                                                            }
                                                            var v56 int = 87
                                                            _ = v56
                                                        case 2147483647:
                                                            switch 988 {
                                                            case 9223372036854775807:
                                                                v1 = (^(v3))
                                                            case 42:
                                                                v50 = v50
                                                            default:
                                                                switch fi39 {
                                                                case -550:
                                                                    var v57 int32 = -2147483648
                                                                    _ = v57
                                                                }
                                                            }
                                                            fmt.Println(v3)
                                                            if (false || (18446744073709551615 != 2)) {
                                                                v50 = v50
                                                                _ = len(v2)
                                                                v4 = v3
                                                            }
                                                        case -2147483648:
                                                            delete(v2, (("hello" + (v50 + (((v50 + (v50 + ("hello" + ("" + v50)))) + "hello") + (v50 + v50)))) + v50))
                                                        }
                                                        delete(v2, (("hello" + ("\n\t" + v50)) + v50))
                                                        var v58 int32 = -827
                                                        _ = v58
                                                    }
                                                    v59 := GenS6{F0: 0}
                                                    var v60 uint = 1
                                                    _ = v50
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        switch ((("a" + (("" + ("foo bar" + ("a" + ""))) + "")) + ((("" + "\n\t") + "a") + "")) + "\n\t") {
                                                        case "":
                                                            switch -149 {
                                                            case -2147483648:
                                                                v4 = v3
                                                            case 127:
                                                                var v61 uint = 1
                                                                _ = v61
                                                            case 128:
                                                                v62 := v2[(("hello" + "hello") + ("hello" + "a"))]
                                                                _ = v62
                                                            default:
                                                                if (!false) {
                                                                    v63 := GenS5{F0: -1, F1: ((!((genFn0(v1) < 1) || false)) && false), F2: 1000}
                                                                    fmt.Println(v4, v1, v3)
                                                                    _ = v63
                                                                }
                                                            }
                                                        case "a":
                                                            var v64 bool = true
                                                            _ = v64
                                                        case "\n\t":
                                                            v4 = (v4 * genFn1())
                                                        default:
                                                            v65 := GenS7{F0: (v3 - 2), F1: (v1 * (v4 | v1))}
                                                            if false {
                                                                var v66 int8 = genFn0(v1)
                                                                var v67 string = "foo bar"
                                                                v3 += 2
                                                                _ = v66
                                                                _ = v67
                                                            }
                                                            _ = v65
                                                        }
                                                    }
                                                    v4 -= v3
                                                    fmt.Println(v4)
                                                }
                                                var v68 int = -633
                                                switch fi40 {
                                                case 4294967295:
                                                    v2["\n\t"] = 4294967295
                                                default:
                                                    v69 := GenS8{F0: 1}
                                                    _ = v69
                                                }
                                                _ = v68
                                            }
                                        case "hello":
                                            v70 := GenS1{F0: 127, F1: 42}
                                            _ = v70
                                        }
                                    }
                                }
                            }
                            var v71 uint32 = 0
                            _ = v71
                        }
                    case "foo bar":
                        if false {
                            fmt.Println(fi40, fi40, v4)
                            v4 = v1
                            var v72 float64 = float64(3.14159)
                            var v73 int64 = -663
                            var v74 uint64 = 4294967295
                            _ = v72
                            _ = v73
                            _ = v74
                        }
                    case "\n\t":
                        if false {
                            if false {
                                var v75 int8 = 127
                                _ = v75
                            }
                        } else {
                            fmt.Println(v1, v3, fi39)
                            v2[("\n\t" + "\n\t")] = uint32(v4)
                            fmt.Println(fi40, fi39)
                        }
                        var v76 float64 = float64((((0.0 - 0.1) / (-2.25 + (float32(0.0) / float32(-2.25)))) / 3.14159))
                        _ = v76
                    default:
                        v4 = (v4 >> ((uint(v1)) & 63))
                    }
                    delete(v2, "")
                    v77 := []int{fi40, (fi39 << ((2) & 63)), fi40}
                    v4 += 1
                    _ = fi40
                    _ = v77
                }
                v2["\n\t"] = 1
                _ = fi39
            }
            v78 := GenS9{F0: genFn0((v1 | v3))}
            _ = v78
        }
        v79 := GenS1{F0: 1, F1: 2}
        v80 := GenS10{F0: false}
        _ = v79
        _ = v80
    }
    switch 32767 {
    case -1:
        fmt.Println(v4)
    default:
        v81 := GenS11{F0: 42}
        _ = v81
    }
    for fi82 := 0; fi82 < 17; fi82++ {
        var v83 uint64 = 907
        _ = fi82
        _ = v83
    }
    fmt.Println(v4)
    if true {
        fmt.Println(v3, v3, v4)
        switch "foo bar" {
        case "foo bar":
            for fi84 := 0; fi84 < 16; fi84++ {
                for fi85 := 0; fi85 < 7; fi85++ {
                    v86 := []int{fi84}
                    _ = fi85
                    _ = v86
                }
                v3 += (^(v1))
                fmt.Println(v4, v3)
                _ = fi84
            }
        case "hello":
            v1 = (v4 & (v4 & genFn1()))
        case "\n\t":
            if false {
                var v87 bool = (float64(0.1) == 1e10)
                _ = v87
            } else {
                if true {
                    if (128 <= 2147483647) {
                        if true {
                            v1 = v1
                            var v88 uint64 = 42
                            v2[("hello" + "hello")] = 242
                            _ = v88
                        } else {
                            v89 := []int{-1, 4294967295, -384, 4294967295}
                            if len(v89) >= 2 {
                                v89 = v89[1:len(v89)]
                            }
                            _ = v89
                        }
                    }
                    var v90 uint32 = 2147483647
                    _ = v90
                } else {
                    if true {
                        var v91 uint16 = 555
                        v1 = v4
                        var v92 string = ("hello" + ((("foo bar" + ("" + "\n\t")) + "a") + ("a" + "foo bar")))
                        _ = v91
                        _ = v92
                    } else {
                        for fi93 := 0; fi93 < 6; fi93++ {
                            v94 := GenS12{F0: 255, F1: 100}
                            if false {
                                v4 = v3
                                if (2 >= 2) {
                                    v4 = genFn1()
                                }
                                v1 += 128
                            }
                            _ = fi93
                            _ = v94
                        }
                    }
                    for fi95 := 0; fi95 < 12; fi95++ {
                        var v96 int = 128
                        _ = fi95
                        _ = v96
                    }
                }
                v97 := []int{-9223372036854775808}
                v4 += 180
                _ = v97
            }
        default:
            v4 = v3
        }
        if false {
            var v98 uint8 = 194
            v99 := []int{-1, -1}
            v100 := GenS3{F0: 0}
            _ = v100
            _ = v98
            _ = v99
        } else {
            var v101 uint8 = (v1 / ((v3) | 1))
            _ = v101
        }
    } else {
        var v102 int32 = 1
        delete(v2, "")
        v103 := []int{-584, 128, 32767}
        _ = v102
        _ = v103
    }
    for fi104 := 0; fi104 < 14; fi104++ {
        fmt.Println(v1, v4, v1)
        _ = fi104
    }
    if true {
        if (("\n\t" + "hello") <= "a") {
            var v105 float64 = 3.14159
            _ = v105
        } else {
            delete(v2, "a")
            v4 += v4
            v106 := GenS13{F0: -774, F1: 398, F2: 4294967295}
            _ = v106
        }
        fmt.Println(v3)
        v1 -= (^(v4))
    } else {
        v1 = v3
    }
    switch 2147483647 {
    case 0:
        var v107 int8 = 116
        _ = v107
    case 127:
        v108 := GenS7{F0: 204, F1: v1}
        _ = v108
    default:
        v109 := GenS2{F0: 820}
        _ = v109
    }
    fmt.Println(v1, v4)
    v110 := v2["a"]
    v111 := GenS14{F0: v110}
    switch -1 {
    case -794:
        switch 4294967295 {
        case 1:
            var v112 int8 = -1
            _ = v112
        case -422:
            fmt.Println(v1)
        }
    case 2147483647:
        v113 := GenS15{F0: 42, F1: 2147483647, F2: 1}
        _ = v113
    case -829:
        v114 := GenS7{F0: 128, F1: genFn1()}
        _ = v114
    default:
        v115 := GenS6{F0: -1}
        _ = v115
    }
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v110
    _ = v111
    _ = v2
    _ = v3
    _ = v4
}
