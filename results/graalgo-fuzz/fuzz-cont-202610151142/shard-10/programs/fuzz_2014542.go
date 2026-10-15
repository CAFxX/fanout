package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 bool; F1 int64; F2 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 int; F2 int8 }
type GenS8 struct { F0 int32; F1 uint8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 string; F1 int8 }
type GenS11 struct { F0 string; F1 int16 }
type GenS12 struct { F0 bool; F1 string }
func genFn0(p0 int64, p1 uint) uint8 {
    var v7 float64 = 3.14159
    v8 := GenS1{F0: 0, F1: 128, F2: 4294967295}
    var v9 float64 = (v7 * float64(0.0))
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 float64 = 1.5
        var v12 uint = 18446744073709551615
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := []int{42, -833}
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]string)
    v3 := make(map[uint32]uint32)
    var v4 int32 = 1
    v5 := make(map[int]int)
    v6 := make(map[int]bool)
    if (v1 == v1) {
        var v14 uint8 = 188
        if false {
            for fi15 := 0; fi15 < 12; fi15++ {
                v16 := []int{(fi15 % (((fi15 - fi15)) | 1)), (fi15 & (^(fi15))), fi15, fi15}
                _ = fi15
                _ = v16
            }
            var v17 uint64 = 42
            if true {
                if false {
                    var v18 int64 = 872
                    if (1.5 > (1e10 * 1e10)) {
                        var v19 int16 = -32768
                        _ = v19
                    }
                    _ = v18
                } else {
                    switch 886 {
                    case 128:
                        switch 2147483647 {
                        case 29:
                            v17 += (v17 ^ v17)
                        }
                    }
                }
                v20 := GenS1{F0: 0, F1: 355, F2: 771}
                switch "hello" {
                case "":
                    if true {
                        var v21 float64 = ((((0.0 - float64(0.1)) / 1.0) - 1e10) * 0.1)
                        var v22 uint16 = 1
                        _ = v21
                        _ = v22
                    }
                case "a":
                    var v23 int64 = v1
                    _ = v23
                case "\n\t":
                    v24 := GenS2{F0: v4, F1: 4294967295}
                    _ = v24
                }
                _ = v20
            }
            _ = v17
        }
        _ = v14
    } else {
        fmt.Println(v4, v4)
        for fi25 := 0; fi25 < 11; fi25++ {
            for fi26 := 0; fi26 < 19; fi26++ {
                switch "hello" {
                case "":
                    v1 = v1
                case "hello":
                    var v27 int = 32768
                    _ = v27
                default:
                    if (0 >= 18446744073709551615) {
                        var v28 float32 = 1e10
                        v29 := GenS3{F0: 119, F1: (fi25 * 1)}
                        v4 += v4
                        v4 = 42
                        _ = v28
                        _ = v29
                    } else {
                        var v30 uint64 = 18446744073709551615
                        if ((1.5 <= (0.1 + (1.5 + float64(-2.25)))) && (2 == 466)) {
                            if false {
                                v31 := GenS2{F0: v4, F1: 255}
                                v4 = (v4 << ((255) & 63))
                                fmt.Println(v1)
                                _ = v31
                            }
                            if false {
                                var v32 uint = 782
                                _ = v32
                            }
                            if (2147483648 <= 2147483648) {
                                if false {
                                    v33 := GenS2{F0: v4, F1: 42}
                                    if false {
                                        switch 271 {
                                        case 0:
                                            v34 := GenS4{F0: uint8(v30), F1: ("foo bar" + "")}
                                            _ = v34
                                        case -756:
                                            v1 -= v1
                                        }
                                        var v35 string = ""
                                        _ = v35
                                    } else {
                                        _ = len(v2)
                                    }
                                    v36 := GenS5{F0: (3.14159 <= (((float32((-2.25 + float64(0.1))) / 1.5) * (((-2.25 - 1.5) / 1.0) * 0.1)) - (1e10 * ((3.14159 + (0.0 / (float32(float64(-2.25)) - 1.5))) * 1.5)))), F1: 322, F2: 42}
                                    switch (fi26 - fi26) {
                                    case -9223372036854775808:
                                        v37 := GenS6{F0: 42}
                                        _ = v37
                                    }
                                    var v38 bool = false
                                    v4 = (v4 | (v4 << ((0) & 63)))
                                    v39 := []int{(fi25 * fi25), (^(fi26))}
                                    var v40 uint16 = 255
                                    _ = v33
                                    _ = v36
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                } else {
                                    v30 -= 314
                                    var v41 uint16 = 65535
                                    _ = v41
                                }
                                v42 := GenS7{F0: (^(v30)), F1: (fi26 >> ((1) & 63)), F2: 1}
                                _ = v42
                            } else {
                                if false {
                                    v43 := []int{-444, 0, fi26, fi25}
                                    v1 = 587
                                    _ = v43
                                }
                                v30 += (v30 / (((v30 - v30)) | 1))
                                v44 := []int{(fi25 << ((124) & 63)), (fi25 + fi25), (fi25 >> ((438) & 63)), fi26}
                                v45 := GenS8{F0: (^(v4)), F1: 176}
                                _ = v44
                                _ = v45
                            }
                        }
                        fmt.Println(v4, fi25, fi25)
                        _ = v30
                    }
                }
                v4 = -973
                _ = fi26
            }
            fmt.Println(v1, v4)
            _ = fi25
        }
        v4 += (v4 / ((v4) | 1))
    }
    var v46 int = 420
    var v47 uint32 = 42
    for fi48 := 0; fi48 < 19; fi48++ {
        v49 := []int{fi48, fi48}
        v4 = (-(v4))
        if true {
            for fi50 := 0; fi50 < 19; fi50++ {
                if true {
                    v46 = 4294967295
                }
                _ = fi50
            }
            if (1 != 1) {
                for fi51 := 0; fi51 < 12; fi51++ {
                    var v52 int = 921
                    v53 := GenS9{F0: (fi51 ^ 871)}
                    _ = fi51
                    _ = v52
                    _ = v53
                }
                for fi54 := 0; fi54 < 5; fi54++ {
                    v55 := GenS6{F0: 0}
                    var v56 uint16 = 0
                    if ((v1 / (((v1 / (((^(v1))) | 1))) | 1)) > (-(v1))) {
                        var v57 uint8 = genFn0(v1, 42)
                        switch v46 {
                        case 579:
                            if (((v56 ^ 0) >= (^(v56))) && true) {
                                var v58 float64 = 1e10
                                v59 := GenS2{F0: 762, F1: 255}
                                _ = v58
                                _ = v59
                            }
                        case 128:
                            switch ("a" + "\n\t") {
                            case "hello":
                                v57 += v57
                            }
                        }
                        var v60 int32 = v4
                        _ = v57
                        _ = v60
                    }
                    _ = fi54
                    _ = v55
                    _ = v56
                }
                var v61 int16 = 0
                _ = v61
            } else {
                v46 -= v46
                var v62 int64 = (v1 | v1)
                v63 := GenS2{F0: v4, F1: 255}
                _ = v62
                _ = v63
            }
            v64 := GenS10{F0: ("foo bar" + "a"), F1: 1}
            _ = v64
        } else {
            var v65 int8 = 1
            for fi66 := 0; fi66 < 12; fi66++ {
                v67 := v3[(v47 - v47)]
                var v68 float32 = (1.5 + 3.14159)
                _ = fi66
                _ = v67
                _ = v68
            }
            v6[(v46 + (fi48 * (fi48 | v46)))] = false
            _ = v65
        }
        _ = fi48
        _ = v49
    }
    v69 := []int{(v46 ^ v46), (v46 * v46), v46}
    for fi70 := 0; fi70 < 4; fi70++ {
        var v71 uint32 = (^(v47))
        for fi72 := 0; fi72 < 8; fi72++ {
            if false {
                fmt.Println(v4)
                if ((-323 == -32768) && (-1 > 1)) {
                    var v73 string = "\n\t"
                    _ = v73
                }
                fmt.Println(fi72, v46)
                v74 := GenS11{F0: "foo bar", F1: 1}
                _ = v74
            } else {
                if true {
                    delete(v6, fi72)
                    var v75 int16 = 452
                    _ = v75
                } else {
                    v76 := v6[(^(fi70))]
                    if ((v76 && v76) && true) {
                        if len(v69) >= 2 {
                            v69 = v69[1:len(v69)]
                        }
                        var v77 float64 = (1.5 / 1.0)
                        switch "a" {
                        case "hello":
                            switch (fi72 >> ((18446744073709551615) & 63)) {
                            case -987:
                                v76 = v76
                            }
                        case "foo bar":
                            v46 = (fi72 + fi70)
                        case "\n\t":
                            v78 := v6[(fi72 % (((fi70 ^ 532)) | 1))]
                            _ = v78
                        default:
                            v47 -= v47
                        }
                        _ = v77
                    }
                    _ = v76
                }
            }
            _ = fi72
        }
        _ = fi70
        _ = v71
    }
    var v79 float64 = float64(3.14159)
    v80 := GenS3{F0: 32767, F1: v46}
    var v81 uint64 = 18446744073709551615
    var v82 uint16 = 65535
    if true {
        if true {
            v46 += (v46 >> ((uint(v81)) & 63))
            var v83 uint64 = (v81 * v81)
            v47 -= (v47 >> ((42) & 63))
            _ = v83
        } else {
            v81 = 0
            v84 := GenS11{F0: (("\n\t" + "hello") + "foo bar"), F1: -622}
            if false {
                if (uint8(v47) <= 2) {
                    v82 -= (v82 / ((255) | 1))
                } else {
                    if (v1 == v1) {
                        if false {
                            v85 := GenS12{F0: false, F1: ("a" + (("\n\t" + "foo bar") + "\n\t"))}
                            _ = v85
                        } else {
                            if false {
                                v82 -= uint16(v47)
                                if ((v1 == int64(v82)) && true) {
                                    var v86 int = v46
                                    _ = v86
                                }
                            }
                            _ = len(v69)
                            var v87 uint16 = (v82 >> ((574) & 63))
                            _ = v87
                        }
                        v1 += -776
                        var v88 uint64 = (v81 ^ v81)
                        _ = v88
                    } else {
                        switch "foo bar" {
                        case "":
                            v89 := GenS9{F0: 2147483647}
                            _ = v89
                        }
                        _ = len(v5)
                        var v90 float64 = (1.5 * v79)
                        _ = v90
                    }
                }
                if len(v69) >= 2 {
                    v69 = v69[1:len(v69)]
                }
                if ((!(v4 == int32(v82))) || true) {
                    var v91 int8 = -1
                    switch (v46 >> ((255) & 63)) {
                    case 792:
                        var v92 float64 = v79
                        _ = v92
                    case -1:
                        var v93 uint16 = 0
                        _ = v93
                    case -9223372036854775808:
                        v82 += v82
                    }
                    _ = v91
                } else {
                    fmt.Println(v46, v1)
                    v94 := GenS9{F0: (^(v46))}
                    v95 := GenS7{F0: 1, F1: (^(v46)), F2: -128}
                    _ = v94
                    _ = v95
                }
            } else {
                if (9223372036854775807 != (v46 & (v46 % ((v46) | 1)))) {
                    v96 := GenS3{F0: (v82 << ((121) & 63)), F1: (^(v46))}
                    for fi97 := 0; fi97 < 2; fi97++ {
                        v46 += fi97
                        for fi98 := 0; fi98 < 6; fi98++ {
                            v46 = int(v81)
                            switch "" {
                            case "foo bar":
                                var v99 int16 = -1
                                _ = v99
                            case "":
                                var v100 float32 = 1e10
                                _ = v100
                            case "\n\t":
                                if false {
                                    if false {
                                        v1 = v1
                                        var v101 bool = (true || true)
                                        _ = v101
                                    }
                                    v81 = (v81 * uint64(v82))
                                    var v102 int8 = 1
                                    var v103 uint16 = 0
                                    _ = v102
                                    _ = v103
                                }
                            default:
                                v46 -= v46
                            }
                            switch ("hello" + "hello") {
                            case "a":
                                var v104 int32 = (v4 & v4)
                                _ = v104
                            }
                            switch fi98 {
                            case -2147483648:
                                v105 := GenS9{F0: int(v4)}
                                _ = v105
                            case -360:
                                if ((!false) || (631 < (v81 * v81))) {
                                    if (int64(fi97) <= (v1 | v1)) {
                                        var v106 int64 = v1
                                        _ = v106
                                    }
                                    if len(v69) > 0 {
                                        v69[0] = v46
                                    }
                                }
                            }
                            v79 += (1e10 / v79)
                            _ = fi98
                        }
                        _ = fi97
                    }
                    _ = v96
                }
                switch (v46 << ((541) & 63)) {
                case 32767:
                    fmt.Println(v47, v81, v82)
                }
                fmt.Println(v79, v82)
            }
            _ = v84
        }
        var v107 uint = 2
        var v108 uint64 = 220
        _ = v107
        _ = v108
    } else {
        if len(v69) >= 2 {
            v69 = v69[1:len(v69)]
        }
    }
    v1 += 9223372036854775807
    for fi109 := 0; fi109 < 14; fi109++ {
        switch fi109 {
        case 182:
            switch "a" {
            case "foo bar":
                var v110 string = "foo bar"
                _ = v110
            case "\n\t":
                if true {
                    fmt.Println(v4)
                }
            }
        case 0:
            fmt.Println(v79)
        case 32768:
            var v111 int = (v46 & fi109)
            _ = v111
        default:
            var v112 uint32 = 65535
            _ = v112
        }
        fmt.Println(v79, fi109)
        _ = fi109
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v69:", len(v69), v69)
    fmt.Println("v79:", v79)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v69
    _ = v79
    _ = v80
    _ = v81
    _ = v82
}
