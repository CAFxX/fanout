package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint8; F1 uint32; F2 bool }
type GenS8 struct { F0 uint32; F1 uint }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 int32 }
type GenS11 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS12 struct { F0 bool; F1 uint32; F2 int8 }
type GenS13 struct { F0 int8; F1 bool; F2 uint8 }
func genFn0(p0 int8) int64 {
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint8 = 128
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 1
    if (v8 != v8) {
        var v9 int8 = p0
        v10 := []int{32768}
        var v11 int16 = -327
        _ = v10
        _ = v11
        _ = v9
    } else {
        switch -310 {
        case 9223372036854775807:
            var v12 bool = (255 != 0)
            _ = v12
        case 0:
            if true {
                if false {
                    v13 := GenS1{F0: 34}
                    _ = v13
                }
            } else {
                for fi14 := 0; fi14 < 2; fi14++ {
                    for fi15 := 0; fi15 < 4; fi15++ {
                        fmt.Println(fi14)
                        var v16 uint32 = 466
                        v17 := []int{fi14, (fi15 / ((fi14) | 1)), fi15, 32768}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    p0 += p0
                    _ = fi14
                }
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 bool = (71 >= 100)
                    fmt.Println(p0, v8)
                    _ = fi18
                    _ = v19
                }
                v8 = (v8 >> ((uint(v8)) & 63))
            }
        case -799:
            v8 += v8
        default:
            var v20 uint64 = 295
            _ = v20
        }
    }
    var v21 int64 = 9223372036854775807
    _ = p0
    _ = v21
    _ = v8
    return (v21 ^ v21)
}
func genFn1(p0 uint8, p1 int64, p2 uint8) int8 {
    if false {
        fmt.Println(p2, p1)
        if (3.14159 >= (0.0 - 0.1)) {
            fmt.Println(p0, p2, p2)
            v22 := GenS1{F0: 32767}
            _ = v22
        }
    } else {
        v23 := GenS1{F0: -353}
        p2 = p2
        _ = v23
    }
    switch -2147483648 {
    case 128:
        for fi24 := 0; fi24 < 11; fi24++ {
            fmt.Println(fi24, p0)
            if (false || true) {
                p2 -= 175
            }
            _ = fi24
        }
    case 32768:
        var v25 uint = 1
        _ = v25
    case 32767:
        var v26 int16 = 0
        _ = v26
    }
    var v27 uint = 950
    switch -2147483648 {
    case 4294967295:
        for fi28 := 0; fi28 < 11; fi28++ {
            switch "foo bar" {
            case "hello":
                v29 := []int{fi28, -881, int(p0)}
                _ = v29
            case "foo bar":
                if (-32768 <= -32768) {
                    var v30 bool = true
                    for fi31 := 0; fi31 < 17; fi31++ {
                        v32 := GenS2{F0: p1, F1: 1}
                        _ = fi31
                        _ = v32
                    }
                    _ = v30
                } else {
                    for fi33 := 0; fi33 < 16; fi33++ {
                        v27 = v27
                        v34 := GenS1{F0: 271}
                        switch fi28 {
                        case 2147483647:
                            fmt.Println(p1, fi28, p0)
                        case 32767:
                            if false {
                                var v35 uint = (v27 + 396)
                                _ = v35
                            } else {
                                p2 += p0
                                p2 += p0
                            }
                        }
                        var v36 float32 = 1e10
                        _ = fi33
                        _ = v34
                        _ = v36
                    }
                }
            case "a":
                fmt.Println(v27, p2)
            }
            _ = fi28
        }
    default:
        if (("hello" + "a") == (((("foo bar" + "foo bar") + "\n\t") + "a") + ("\n\t" + ("\n\t" + ("foo bar" + ""))))) {
            var v37 int16 = -1
            p2 -= (^(p2))
            v38 := GenS1{F0: (v37 % ((v37) | 1))}
            _ = v37
            _ = v38
        }
    }
    v39 := GenS3{F0: 460, F1: 4294967295, F2: (p1 >> (((v27 | (^(v27)))) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v39
    return 103
}

func main() {
    var v1 float64 = ((((0.1 * 3.14159) + 1.5) / 1e10) - 0.0)
    var v2 int64 = -254
    var v3 float32 = 1.5
    v4 := make(map[int]uint32)
    var v5 uint64 = 2
    if true {
        var v40 float32 = v3
        var v41 uint32 = 2147483647
        _ = v40
        _ = v41
    } else {
        for fi42 := 0; fi42 < 15; fi42++ {
            v2 -= (v2 + genFn0(0))
            if false {
                switch ((("\n\t" + "a") + "") + "\n\t") {
                case "hello":
                    fmt.Println(v5, fi42, v1)
                case "\n\t":
                    v43 := []int{}
                    _ = v43
                }
                fmt.Println(v2, v3)
                v44 := GenS4{F0: (v2 | (v2 * v2)), F1: 0}
                _ = v44
            } else {
                v5 -= (v5 | 42)
            }
            _ = fi42
        }
        for fi45 := 0; fi45 < 5; fi45++ {
            if true {
                v46 := []int{}
                if ("a" <= "a") {
                    v4[fi45] = 701
                    fmt.Println(fi45, v2, v2)
                } else {
                    v1 += (v1 / (v1 * (float64(0.0) / v1)))
                    for fi47 := 0; fi47 < 7; fi47++ {
                        if (0 > 928) {
                            switch "a" {
                            case "":
                                var v48 float32 = float32(((v1 / (0.0 - 0.1)) - v1))
                                _ = v48
                            }
                            if len(v46) > 0 {
                                v46[0] = fi45
                            }
                        } else {
                            v5 = (v5 / ((v5) | 1))
                            v2 -= 1
                            v4[fi45] = 42
                        }
                        if (0.0 >= v3) {
                            var v49 uint = 18446744073709551615
                            v2 = v2
                            _ = v49
                        }
                        switch (fi45 | fi45) {
                        case 269:
                            if len(v46) > 0 {
                                v46[0] = (fi47 << ((2) & 63))
                            }
                            var v50 float64 = v1
                            fmt.Println(fi47, v2, fi47)
                            _ = v50
                        case 100:
                            fmt.Println(fi45, v2, fi45)
                        }
                        v1 -= 1.5
                        var v51 uint16 = 58
                        _ = fi47
                        _ = v51
                    }
                    var v52 float64 = float64(float32((float64(v3) * (v1 * 0.0))))
                    _ = v52
                }
                v2 = (v2 + v2)
                _ = v46
            }
            v53 := GenS5{F0: 474, F1: 0}
            _ = fi45
            _ = v53
        }
        v54 := GenS5{F0: 177, F1: genFn1(216, (v2 % ((v2) | 1)), 130)}
        _ = v54
    }
    var v55 uint64 = (v5 << ((488) & 63))
    v3 -= v3
    fmt.Println(v55)
    v1 = (((1e10 - (float64(v3) - v1)) / ((1e10 + (v1 / v1)) + float64(1e10))) + v1)
    v56 := []int{}
    v57 := GenS2{F0: v2, F1: 2}
    fmt.Println(v3)
    v58 := GenS1{F0: 1}
    if true {
        var v59 float64 = (((3.14159 / 1e10) + float64(float32(((v1 * v1) / float64(v3))))) - -2.25)
        v60 := v4[32768]
        v1 -= v1
        _ = v59
        _ = v60
    } else {
        v3 -= 3.14159
        v61 := GenS5{F0: 0, F1: -79}
        if len(v56) >= 2 {
            v56 = v56[1:len(v56)]
        }
        _ = v61
    }
    var v62 string = "\n\t"
    for fi63 := 0; fi63 < 16; fi63++ {
        for fi64 := 0; fi64 < 20; fi64++ {
            switch fi64 {
            case -216:
                switch ("a" + "a") {
                case "foo bar":
                    fmt.Println(fi63, fi63)
                case "a":
                    var v65 int64 = (v2 % (((v2 - -9223372036854775808)) | 1))
                    switch fi63 {
                    case 193:
                        v4[fi64] = 936
                    case 127:
                        if false {
                            delete(v4, (fi63 << ((359) & 63)))
                        } else {
                            fmt.Println(v55, fi63)
                            v66 := GenS5{F0: 42, F1: 127}
                            v67 := GenS6{F0: v62}
                            _ = v66
                            _ = v67
                        }
                    case 0:
                        v55 = v5
                    }
                    _ = v65
                case "\n\t":
                    switch v62 {
                    case "a":
                        v68 := GenS7{F0: 0, F1: 0, F2: true}
                        _ = v68
                    case "":
                        v1 = v1
                    case "\n\t":
                        switch -22 {
                        case 834:
                            switch (("a" + v62) + "foo bar") {
                            case "foo bar":
                                var v69 int8 = 127
                                _ = v69
                            case "a":
                                if ((true || true) && (true || (((true && false) && false) && false))) {
                                    v62 = (v62 + "\n\t")
                                }
                            case "hello":
                                switch ("\n\t" + v62) {
                                case "foo bar":
                                    v70 := GenS3{F0: 2147483647, F1: v5, F2: v2}
                                    _ = v70
                                case "":
                                    var v71 uint16 = 1
                                    _ = v71
                                }
                            }
                        }
                    }
                default:
                    var v72 int32 = 1
                    _ = v72
                }
            case 4294967295:
                if (false || (true || true)) {
                    if (true && false) {
                        v5 = (v55 + 18446744073709551615)
                        fmt.Println(fi63)
                        v2 -= 42
                    }
                    v55 -= v5
                    var v73 int8 = 42
                    _ = v73
                }
                if false {
                    if (v55 < 18446744073709551615) {
                        var v74 uint8 = 0
                        _ = v74
                    }
                    if (v3 == 1e10) {
                        if (v3 < (v3 * v3)) {
                            v75 := v4[fi63]
                            v2 += genFn0(-38)
                            var v76 int = (fi63 * (fi64 + 0))
                            _ = v75
                            _ = v76
                        }
                        var v77 float32 = v3
                        fmt.Println(v2, v55, v55)
                        _ = v77
                    } else {
                        if true {
                            fmt.Println(v2, v55)
                            if (255 != 2) {
                                v56 = append(v56, -2147483648)
                                var v78 int32 = 0
                                v62 = ((v62 + v62) + v62)
                                fmt.Println(v2, fi63, v3)
                                _ = v78
                            }
                            if (1 > 1) {
                                fmt.Println(fi64, fi63, v2)
                                v1 -= v1
                                if (false && false) {
                                    if (("foo bar" != (v62 + ((v62 + (v62 + v62)) + v62))) || true) {
                                        if false {
                                            var v79 uint64 = (v55 - v5)
                                            if false {
                                                v3 += float32(v1)
                                                v80 := GenS4{F0: (v2 >> ((18446744073709551615) & 63)), F1: -810}
                                                switch v62 {
                                                case "":
                                                    v5 -= (v55 * 4294967295)
                                                case "hello":
                                                    if (1000 == -35) {
                                                        if (-586 >= 209) {
                                                            var v81 string = (v62 + v62)
                                                            v81 = "a"
                                                            v56 = append(v56, int(v2))
                                                            var v82 int32 = -2147483648
                                                            var v83 int32 = v82
                                                            v84 := GenS8{F0: 392, F1: 255}
                                                            switch fi64 {
                                                            case 128:
                                                                v4[(fi64 + 0)] = uint32(v55)
                                                            default:
                                                                var v85 uint32 = 2147483647
                                                                _ = v85
                                                            }
                                                            _ = v81
                                                            _ = v82
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                        var v86 bool = ((2 != 2147483648) || true)
                                                        _ = v86
                                                    }
                                                default:
                                                    v79 += (v55 + v55)
                                                }
                                                switch "hello" {
                                                case "\n\t":
                                                    if len(v56) >= 2 {
                                                        v56 = v56[1:len(v56)]
                                                    }
                                                case "":
                                                    v5 += 178
                                                case "hello":
                                                    var v87 uint8 = 255
                                                    _ = v87
                                                }
                                                v2 = 0
                                                var v88 uint = 1
                                                v89 := GenS2{F0: (v2 - v2), F1: 9}
                                                _ = v80
                                                _ = v88
                                                _ = v89
                                            }
                                            _ = v79
                                        } else {
                                            if ((42 != 435) || (!((0 == -1) || false))) {
                                                var v90 float32 = 3.14159
                                                _ = v90
                                            }
                                            var v91 float64 = -2.25
                                            v92 := GenS9{F0: (fi63 + fi64)}
                                            _ = v91
                                            _ = v92
                                        }
                                    }
                                    delete(v4, fi64)
                                }
                            }
                        } else {
                            fmt.Println(v2, v62, fi63)
                            var v93 int8 = -128
                            v94 := GenS10{F0: fi64, F1: 1}
                            switch v62 {
                            case "hello":
                                switch ("foo bar" + v62) {
                                case "\n\t":
                                    v55 -= 2
                                default:
                                    fmt.Println(v93, v3, v5)
                                }
                            default:
                                v62 = v62
                            }
                            _ = v93
                            _ = v94
                        }
                    }
                    v95 := GenS11{F0: 97, F1: 2147483648, F2: 127}
                    _ = v95
                }
            case 127:
                var v96 string = "a"
                _ = v96
            }
            if true {
                switch (fi64 & fi64) {
                case -855:
                    var v97 uint8 = 164
                    _ = v97
                case 42:
                    var v98 float32 = v3
                    _ = v98
                default:
                    v62 = "a"
                }
                var v99 int8 = 119
                _ = v99
            }
            v2 += v2
            _ = fi64
        }
        _ = fi63
    }
    for fi100 := 0; fi100 < 3; fi100++ {
        for fi101 := 0; fi101 < 14; fi101++ {
            v55 += (v5 ^ v5)
            if true {
                if (true && (!(v1 == (v1 + (v1 / 3.14159))))) {
                    switch v62 {
                    case "a":
                        switch fi100 {
                        case 4294967295:
                            v102 := GenS7{F0: 68, F1: 1, F2: true}
                            _ = v102
                        case -249:
                            v103 := GenS12{F0: (!false), F1: 0, F2: 100}
                            _ = v103
                        case 127:
                            v104 := GenS9{F0: fi100}
                            _ = v104
                        default:
                            if ((^(v2)) != (^(v2))) {
                                if true {
                                    v105 := GenS2{F0: (^(v2)), F1: 1}
                                    _ = v105
                                }
                                if (42 == 0) {
                                    v106 := GenS13{F0: -66, F1: true, F2: 79}
                                    var v107 int16 = -558
                                    _ = v106
                                    _ = v107
                                }
                            } else {
                                v2 += genFn0(0)
                            }
                        }
                    case "":
                        v62 = ("a" + "hello")
                    }
                    v108 := GenS4{F0: genFn0(127), F1: 315}
                    _ = v108
                } else {
                    v3 -= (v3 - (float32(v1) + -2.25))
                    v109 := GenS5{F0: -344, F1: -3}
                    _ = v109
                }
            }
            _ = fi101
        }
        v110 := GenS12{F0: false, F1: 42, F2: 8}
        v1 -= v1
        _ = fi100
        _ = v110
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v62
}
