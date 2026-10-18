package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 int; F1 uint32; F2 string }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int16; F1 int }
func genFn0(p0 uint64, p1 int16) uint {
    v5 := []int{-166, 229}
    _ = len(v5)
    p1 += p1
    if true {
        p1 = (p1 >> ((0) & 63))
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        p0 += 0
        for fi7 := 0; fi7 < 3; fi7++ {
            p1 -= -327
            _ = len(v5)
            var v8 uint64 = 773
            var v9 uint8 = 241
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1(p0 int16) int8 {
    v10 := GenS1{F0: 255, F1: "foo bar", F2: 32767}
    var v11 int8 = 76
    if (true || (117 == 1)) {
        fmt.Println(v11, v11)
    } else {
        fmt.Println(v11, p0)
    }
    fmt.Println(p0)
    if false {
        v12 := GenS1{F0: 893, F1: "\n\t", F2: 0}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 6; fi13++ {
            fmt.Println(p0)
            p0 = (p0 & int16(fi13))
            _ = fi13
        }
    }
    _ = p0
    _ = v10
    _ = v11
    return v11
}

func main() {
    var v1 bool = (false || (910 >= -1))
    v2 := make(map[uint32]int)
    var v3 uint16 = 0
    var v4 float32 = float32(0.0)
    switch 839 {
    case 128:
        var v14 uint64 = 151
        _ = v14
    case 1:
        switch "\n\t" {
        case "\n\t":
            v2[2147483648] = 32768
        default:
            fmt.Println(v4, v1)
        }
    case 729:
        var v15 int16 = 1000
        _ = v15
    }
    v1 = (v1 && (0.0 != float32(3.14159)))
    fmt.Println(v4, v3, v3)
    if true {
        if v1 {
            switch "foo bar" {
            case "foo bar":
                for fi16 := 0; fi16 < 13; fi16++ {
                    v1 = false
                    var v17 int64 = -9223372036854775808
                    _ = fi16
                    _ = v17
                }
            case "\n\t":
                for fi18 := 0; fi18 < 9; fi18++ {
                    var v19 int16 = 0
                    switch 9223372036854775807 {
                    case 32768:
                        v4 = (v4 / v4)
                    case 127:
                        v20 := GenS1{F0: 2, F1: "\n\t", F2: 153}
                        _ = v20
                    }
                    _ = fi18
                    _ = v19
                }
            case "a":
                switch ((("\n\t" + "a") + "\n\t") + "\n\t") {
                case "hello":
                    for fi21 := 0; fi21 < 10; fi21++ {
                        var v22 float64 = 0.0
                        _ = fi21
                        _ = v22
                    }
                case "foo bar":
                    for fi23 := 0; fi23 < 20; fi23++ {
                        v24 := GenS2{F0: 4294967295}
                        _ = fi23
                        _ = v24
                    }
                default:
                    if (32768 < 2147483647) {
                        fmt.Println(v3, v1)
                        if false {
                            var v25 bool = false
                            _ = v25
                        } else {
                            if v1 {
                                switch "\n\t" {
                                case "hello":
                                    v26 := []int{32767, 156}
                                    _ = v26
                                }
                                for fi27 := 0; fi27 < 6; fi27++ {
                                    var v28 int32 = 629
                                    v4 -= v4
                                    v29 := GenS2{F0: 4294967295}
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                }
                                for fi30 := 0; fi30 < 1; fi30++ {
                                    v3 = v3
                                    _ = fi30
                                }
                            } else {
                                v3 -= v3
                                for fi31 := 0; fi31 < 9; fi31++ {
                                    var v32 float64 = -2.25
                                    var v33 int8 = 42
                                    _ = fi31
                                    _ = v32
                                    _ = v33
                                }
                            }
                        }
                        v34 := GenS1{F0: v3, F1: ("\n\t" + "\n\t"), F2: 32767}
                        _ = v34
                    } else {
                        fmt.Println(v1, v4)
                    }
                }
            }
            v35 := []int{-9223372036854775808, 127, 42}
            v3 -= v3
            _ = v35
        } else {
            for fi36 := 0; fi36 < 19; fi36++ {
                v37 := GenS1{F0: (v3 | v3), F1: ("\n\t" + (("\n\t" + ("a" + "a")) + "hello")), F2: v3}
                if v1 {
                    if (!((-2147483648 < 2147483647) && (-1 >= int32(fi36)))) {
                        v1 = true
                    }
                    fmt.Println(fi36, fi36)
                    var v38 bool = (v3 != v3)
                    _ = v38
                }
                fmt.Println(v1, v3, v3)
                _ = fi36
                _ = v37
            }
            var v39 uint = 65535
            _ = v39
        }
        v4 = (v4 + float32(3.14159))
        if v1 {
            v1 = false
            var v40 uint8 = 40
            _ = v40
        } else {
            fmt.Println(v3, v3, v1)
            v4 = float32((0.0 + 0.1))
            v1 = v1
        }
    } else {
        var v41 int32 = 42
        var v42 int8 = 0
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 2; fi43++ {
        v44 := GenS3{F0: -92, F1: 1, F2: 32767}
        _ = fi43
        _ = v44
    }
    fmt.Println(v1, v1)
    fmt.Println(v1, v3)
    v45 := []int{}
    if false {
        for fi46 := 0; fi46 < 16; fi46++ {
            var v47 int16 = 32767
            if (-2.25 >= -2.25) {
                _ = len(v2)
                for fi48 := 0; fi48 < 10; fi48++ {
                    if true {
                        v4 = 3.14159
                        var v49 float64 = 0.0
                        _ = v49
                    }
                    v3 -= (v3 + (^(v3)))
                    _ = fi48
                }
            } else {
                fmt.Println(v3)
                fmt.Println(v1, v1)
                v50 := GenS1{F0: v3, F1: ("hello" + "hello"), F2: v3}
                _ = v50
            }
            _ = len(v2)
            _ = fi46
            _ = v47
        }
        var v51 uint8 = 0
        _ = v51
    }
    v4 += float32(-2.25)
    if (!true) {
        if (2 != 1) {
            _ = len(v45)
            v52 := v2[1]
            for fi53 := 0; fi53 < 8; fi53++ {
                v52 = (^(fi53))
                if (v1 || false) {
                    _ = len(v45)
                    v3 = (v3 | (^(v3)))
                    for fi54 := 0; fi54 < 14; fi54++ {
                        _ = len(v45)
                        if len(v45) >= 2 {
                            v45 = v45[1:len(v45)]
                        }
                        v3 -= (v3 >> ((2) & 63))
                        _ = fi54
                    }
                }
                _ = fi53
            }
            _ = v52
        }
    } else {
        for fi55 := 0; fi55 < 11; fi55++ {
            if (true || (42 != v3)) {
                if (585 >= 42) {
                    v56 := GenS3{F0: 42, F1: 223, F2: -1}
                    for fi57 := 0; fi57 < 6; fi57++ {
                        if ((v3 >> ((genFn0(1, 351)) & 63)) < 32767) {
                            var v58 uint = 654
                            switch "hello" {
                            case "a":
                                if len(v45) > 0 {
                                    v45[0] = fi57
                                }
                                v59 := GenS1{F0: (^(v3)), F1: "foo bar", F2: 0}
                                v60 := GenS4{F0: 255}
                                if v1 {
                                    var v61 int8 = 127
                                    fmt.Println(fi57)
                                    if true {
                                        switch (fi57 << ((v58) & 63)) {
                                        case 2147483647:
                                            if v1 {
                                                if (!((fi55 + (fi57 | (-(fi55)))) == (fi57 | (fi57 << ((65535) & 63))))) {
                                                    switch fi57 {
                                                    case 32768:
                                                        v62 := GenS2{F0: 65535}
                                                        _ = v62
                                                    case -489:
                                                        v58 = 0
                                                    case 32767:
                                                        v63 := GenS1{F0: (v3 + (^(v3))), F1: ("a" + ("foo bar" + "a")), F2: v3}
                                                        _ = v63
                                                    }
                                                    fmt.Println(v4)
                                                    v58 = 255
                                                }
                                                var v64 float32 = 0.1
                                                var v65 string = (("a" + "hello") + "hello")
                                                _ = v64
                                                _ = v65
                                            } else {
                                                v66 := GenS1{F0: (^(v3)), F1: "foo bar", F2: 255}
                                                _ = v66
                                            }
                                        default:
                                            if (!v1) {
                                                var v67 float32 = v4
                                                switch "a" {
                                                case "a":
                                                    var v68 int16 = -1
                                                    _ = v68
                                                }
                                                _ = v67
                                            }
                                        }
                                        fmt.Println(fi57, v58, v58)
                                        switch "hello" {
                                        case "\n\t":
                                            var v69 uint32 = 42
                                            _ = v69
                                        case "a":
                                            var v70 uint32 = 2147483648
                                            _ = v70
                                        case "foo bar":
                                            var v71 bool = (v1 || v1)
                                            _ = v71
                                        default:
                                            v72 := GenS3{F0: (v61 + v61), F1: 1, F2: 32767}
                                            _ = v72
                                        }
                                        var v73 int16 = 42
                                        var v74 float64 = 1.5
                                        if len(v45) >= 2 {
                                            v45 = v45[1:len(v45)]
                                        }
                                        v74 = (float64(((v4 + float32(v74)) * 0.1)) + v74)
                                        _ = v73
                                        _ = v74
                                    } else {
                                        fmt.Println(v58, v58)
                                        if true {
                                            switch (^(fi55)) {
                                            case 317:
                                                if (v1 || (0 >= 32767)) {
                                                    v4 += v4
                                                    v1 = (127 <= 758)
                                                } else {
                                                    v75 := GenS4{F0: 73}
                                                    var v76 float32 = (v4 / v4)
                                                    _ = v75
                                                    _ = v76
                                                }
                                            }
                                        } else {
                                            fmt.Println(v3, v4, v1)
                                            var v77 int16 = 42
                                            _ = v77
                                        }
                                        switch fi55 {
                                        case -9223372036854775808:
                                            var v78 int16 = -32768
                                            _ = v78
                                        case 539:
                                            var v79 float32 = float32((3.14159 - float64(v4)))
                                            _ = v79
                                        default:
                                            var v80 int = 990
                                            _ = v80
                                        }
                                    }
                                    fmt.Println(v3, fi55, v58)
                                    v81 := GenS5{F0: 42, F1: 4294967295}
                                    _ = v61
                                    _ = v81
                                }
                                switch fi57 {
                                case 4294967295:
                                    var v82 int64 = 1
                                    _ = v82
                                case -225:
                                    var v83 string = "a"
                                    _ = v83
                                case 127:
                                    if v1 {
                                        v84 := v2[2147483648]
                                        _ = v84
                                    }
                                }
                                _ = v59
                                _ = v60
                            case "":
                                if v1 {
                                    if (((v3 >> (((v58 % ((412) | 1))) & 63)) < v3) && false) {
                                        v85 := GenS1{F0: v3, F1: "", F2: v3}
                                        v86 := GenS3{F0: 1, F1: 90, F2: 0}
                                        _ = v85
                                        _ = v86
                                    } else {
                                        if len(v45) > 0 {
                                            v45[0] = (fi57 & fi55)
                                        }
                                        if v1 {
                                            v58 = 793
                                            switch "hello" {
                                            case "a":
                                                var v87 uint = 850
                                                fmt.Println(fi57)
                                                v3 += 306
                                                _ = v87
                                            default:
                                                var v88 int64 = 42
                                                _ = v88
                                            }
                                            fmt.Println(v3)
                                            v58 += 338
                                            if len(v45) > 0 {
                                                v45[0] = (fi55 * fi57)
                                            }
                                        } else {
                                            if v1 {
                                                v89 := GenS1{F0: (v3 & 2), F1: ("hello" + "hello"), F2: (v3 >> (((v58 * 65535)) & 63))}
                                                _ = v89
                                            } else {
                                                var v90 uint = v58
                                                v1 = (143 <= 1)
                                                fmt.Println(v90, v58, v3)
                                                v91 := GenS6{F0: 42, F1: 42, F2: ""}
                                                var v92 int64 = -1
                                                if v1 {
                                                    v93 := GenS7{F0: -32768, F1: 0}
                                                    v1 = false
                                                    if true {
                                                        v1 = true
                                                        var v94 uint = (^(v90))
                                                        var v95 uint = v94
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                    _ = v93
                                                }
                                                var v96 int8 = -1
                                                v97 := GenS2{F0: 333}
                                                if false {
                                                    var v98 bool = v1
                                                    _ = v98
                                                }
                                                _ = v90
                                                _ = v91
                                                _ = v92
                                                _ = v96
                                                _ = v97
                                            }
                                        }
                                        v58 = 739
                                    }
                                    v3 -= (v3 + (v3 * v3))
                                    var v99 float64 = 3.14159
                                    _ = v99
                                } else {
                                    fmt.Println(v4, v4)
                                    var v100 int64 = 1
                                    _ = v100
                                }
                                switch 127 {
                                case -2147483648:
                                    fmt.Println(v58)
                                default:
                                    fmt.Println(v4)
                                }
                                var v101 uint32 = 995
                                v3 -= (^(v3))
                                var v102 int = fi57
                                _ = v101
                                _ = v102
                            }
                            _ = len(v45)
                            _ = v58
                        }
                        v3 -= (v3 & v3)
                        _ = fi57
                    }
                    v3 += (v3 / ((2) | 1))
                    _ = v56
                }
                if v1 {
                    v45 = append(v45, (fi55 >> ((1) & 63)))
                    if (false || (3.14159 >= 3.14159)) {
                        v4 += (v4 + (3.14159 * (float32(-2.25) * v4)))
                    } else {
                        var v103 uint32 = 65535
                        _ = v103
                    }
                    v2[0] = 1
                } else {
                    v104 := GenS8{F0: 127}
                    v4 = v4
                    _ = v104
                }
                switch "\n\t" {
                case "a":
                    fmt.Println(v1, v3, v4)
                case "":
                    var v105 bool = (((v1 || (!(0 == -2147483648))) || true) && v1)
                    _ = v105
                default:
                    v4 -= v4
                }
            } else {
                var v106 uint16 = v3
                if (0 <= 2) {
                    switch (fi55 % (((fi55 / ((fi55) | 1))) | 1)) {
                    case 0:
                        v107 := v2[2147483647]
                        _ = v107
                    case 2147483647:
                        if false {
                            var v108 int32 = 2147483647
                            _ = v108
                        } else {
                            for fi109 := 0; fi109 < 9; fi109++ {
                                if v1 {
                                    v1 = (42 <= 409)
                                }
                                _ = fi109
                            }
                            v110 := GenS6{F0: (fi55 + (fi55 % ((fi55) | 1))), F1: 2147483648, F2: ""}
                            _ = v110
                        }
                    default:
                        v1 = false
                    }
                    for fi111 := 0; fi111 < 9; fi111++ {
                        delete(v2, 42)
                        v1 = (!true)
                        var v112 uint32 = 65535
                        _ = fi111
                        _ = v112
                    }
                }
                _ = v106
            }
            v113 := GenS9{F0: -917, F1: fi55}
            _ = fi55
            _ = v113
        }
        fmt.Println(v3)
    }
    for fi114 := 0; fi114 < 7; fi114++ {
        fmt.Println(v1, v1, v3)
        _ = fi114
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
}
