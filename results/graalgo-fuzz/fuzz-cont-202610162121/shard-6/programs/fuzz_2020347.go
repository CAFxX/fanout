package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 int16; F2 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64; F1 int8; F2 string }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint16; F1 int64 }
type GenS11 struct { F0 uint8 }
func genFn0(p0 int64) int {
    if true {
        p0 += p0
        fmt.Println(p0)
    } else {
        if (!((false || false) || true)) {
            v6 := GenS1{F0: 236, F1: 1000}
            v7 := []int{32768, 2147483647, 311}
            var v8 int8 = 127
            _ = v6
            _ = v7
            _ = v8
        }
        fmt.Println(p0)
        if true {
            for fi9 := 0; fi9 < 11; fi9++ {
                fmt.Println(fi9, fi9)
                v10 := []int{(fi9 >> ((0) & 63)), fi9}
                _ = fi9
                _ = v10
            }
            if true {
                for fi11 := 0; fi11 < 16; fi11++ {
                    v12 := []int{(fi11 + fi11), (fi11 - fi11)}
                    var v13 int16 = 32767
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                for fi14 := 0; fi14 < 12; fi14++ {
                    if true {
                        var v15 int16 = 1000
                        p0 = p0
                        p0 -= (p0 % ((p0) | 1))
                        _ = v15
                    }
                    fmt.Println(fi14)
                    _ = fi14
                }
                if true {
                    switch "a" {
                    case "hello":
                        switch 7 {
                        case 2147483647:
                            fmt.Println(p0)
                        case 128:
                            fmt.Println(p0)
                        }
                    case "a":
                        if false {
                            if true {
                                switch 32767 {
                                case 42:
                                    var v16 float32 = 3.14159
                                    _ = v16
                                }
                            } else {
                                v17 := []int{-707, 883}
                                var v18 uint16 = 42
                                _ = v17
                                _ = v18
                            }
                            var v19 float32 = ((((float32(0.0) * (-2.25 + (float32(float64(0.1)) / 1.0))) - float32(0.0)) - (1.5 / 1.0)) + (1e10 / (-2.25 - 3.14159)))
                            _ = v19
                        } else {
                            v20 := []int{40, 979, 1, 637}
                            p0 = p0
                            _ = v20
                        }
                    case "":
                        if false {
                            switch ((("hello" + "a") + "") + "\n\t") {
                            case "foo bar":
                                v21 := []int{32768, 32768, 1, -766}
                                _ = v21
                            case "\n\t":
                                var v22 uint8 = 0
                                _ = v22
                            default:
                                v23 := GenS2{F0: 1}
                                _ = v23
                            }
                        } else {
                            v24 := []int{}
                            v25 := GenS3{F0: 672}
                            if false {
                                if false {
                                    p0 += p0
                                }
                                if len(v24) > 0 {
                                    v24[0] = 798
                                }
                            }
                            _ = v24
                            _ = v25
                        }
                    default:
                        fmt.Println(p0)
                    }
                } else {
                    p0 += -1
                    var v26 uint = 508
                    fmt.Println(p0, v26)
                    _ = v26
                }
            }
        } else {
            if (false || false) {
                v27 := []int{}
                fmt.Println(p0)
                _ = v27
            } else {
                var v28 int16 = -32768
                _ = v28
            }
            v29 := GenS4{F0: 104}
            fmt.Println(p0)
            _ = v29
        }
    }
    v30 := []int{4294967295, 128, 867, 344}
    v30 = append(v30, 0)
    _ = p0
    _ = v30
    return 1
}

func main() {
    var v1 float64 = 3.14159
    var v2 string = "a"
    var v3 bool = (3.14159 > v1)
    var v4 uint32 = 235
    v5 := make(map[uint32]int)
    var v31 uint32 = v4
    fmt.Println(v1, v31)
    if v3 {
        fmt.Println(v4)
        if (true || (!true)) {
            var v32 uint16 = 162
            v33 := GenS1{F0: 2, F1: 42}
            _ = v32
            _ = v33
        } else {
            if ((("" + (v2 + ((v2 + "hello") + ((v2 + (v2 + v2)) + v2)))) + "\n\t") < ((("" + ((v2 + "") + "a")) + (v2 + v2)) + (((v2 + "a") + ("hello" + (v2 + v2))) + ""))) {
                if v3 {
                    fmt.Println(v1, v4, v3)
                    var v34 int32 = -1
                    _ = v34
                }
                v31 -= (v4 % ((2147483648) | 1))
            }
        }
        for fi35 := 0; fi35 < 1; fi35++ {
            var v36 bool = (32 <= 9223372036854775807)
            v36 = (v31 >= v4)
            _ = fi35
            _ = v36
        }
    } else {
        v1 = ((v1 - 3.14159) / (((0.0 + -2.25) - (1.5 - 0.0)) - v1))
        switch ((((v2 + "") + "") + "") + (v2 + v2)) {
        case "":
            if v3 {
                fmt.Println(v1)
                var v37 int64 = -475
                _ = v37
            }
        }
    }
    v38 := []int{-1}
    _ = len(v5)
    var v39 float32 = (-2.25 / 1.5)
    if v3 {
        if ((-605 <= 42) && (963 > 9223372036854775807)) {
            switch v2 {
            case "foo bar":
                var v40 uint16 = 795
                _ = v40
            case "":
                fmt.Println(v2)
            case "\n\t":
                v31 = (^(v31))
            default:
                v2 = v2
            }
            var v41 int16 = 1
            var v42 float32 = 1e10
            _ = v41
            _ = v42
        } else {
            for fi43 := 0; fi43 < 15; fi43++ {
                var v44 uint64 = 4294967295
                _ = fi43
                _ = v44
            }
            var v45 uint64 = 18446744073709551615
            _ = v45
        }
        if (2 != 3) {
            if len(v38) > 0 {
                v38[0] = 42
            }
            if true {
                if (v3 && (-2147483648 != 0)) {
                    var v46 uint8 = 2
                    _ = v46
                }
                v1 += 3.14159
            }
            switch v2 {
            case "a":
                var v47 int8 = -4
                _ = v47
            }
        } else {
            var v48 uint32 = (v4 * (v31 * 0))
            _ = v48
        }
        if v3 {
            var v49 int64 = 42
            switch ("foo bar" + v2) {
            case "a":
                if (0 <= 255) {
                    v5[(v4 * (v31 / ((v4) | 1)))] = -9223372036854775808
                    if (0 < 1) {
                        if (668 == 165) {
                            var v50 int8 = 1
                            if len(v38) >= 2 {
                                v38 = v38[1:len(v38)]
                            }
                            v50 = v50
                            _ = v50
                        }
                        v51 := GenS3{F0: 4294967295}
                        fmt.Println(v2, v4, v3)
                        _ = v51
                    } else {
                        var v52 uint16 = 65535
                        for fi53 := 0; fi53 < 8; fi53++ {
                            if false {
                                v5[(v31 % (((v31 - v4)) | 1))] = (-(fi53))
                                var v54 uint64 = 4294967295
                                _ = v54
                            } else {
                                for fi55 := 0; fi55 < 3; fi55++ {
                                    v56 := v5[4294967295]
                                    switch v2 {
                                    case "":
                                        _ = len(v5)
                                    case "foo bar":
                                        v49 -= (v49 / ((v49) | 1))
                                    }
                                    var v57 int64 = v49
                                    _ = fi55
                                    _ = v56
                                    _ = v57
                                }
                                if (v3 && v3) {
                                    var v58 int64 = v49
                                    _ = len(v38)
                                    v5[v31] = fi53
                                    _ = v58
                                }
                            }
                            if v3 {
                                var v59 int64 = v49
                                for fi60 := 0; fi60 < 12; fi60++ {
                                    var v61 int8 = 127
                                    var v62 float64 = v1
                                    _ = fi60
                                    _ = v61
                                    _ = v62
                                }
                                switch genFn0((v59 | v59)) {
                                case -2147483648:
                                    for fi63 := 0; fi63 < 19; fi63++ {
                                        if v3 {
                                            v31 += v4
                                            if v3 {
                                                var v64 int = fi63
                                                if (v2 < v2) {
                                                    if (221 <= 24) {
                                                        fmt.Println(v39, v1)
                                                        _ = len(v38)
                                                    }
                                                } else {
                                                    fmt.Println(v39, v4)
                                                }
                                                _ = v64
                                            }
                                        }
                                        v65 := v5[(^(v4))]
                                        var v66 uint = 0
                                        _ = fi63
                                        _ = v65
                                        _ = v66
                                    }
                                default:
                                    for fi67 := 0; fi67 < 7; fi67++ {
                                        _ = len(v5)
                                        _ = fi67
                                    }
                                }
                                _ = v59
                            }
                            _ = fi53
                        }
                        _ = v52
                    }
                }
            default:
                var v68 uint16 = 0
                _ = v68
            }
            _ = v49
        } else {
            switch 1 {
            case 1:
                fmt.Println(v31, v2)
            case -2147483648:
                if true {
                    var v69 uint16 = 255
                    fmt.Println(v69)
                    var v70 int16 = 1
                    _ = v69
                    _ = v70
                }
            default:
                v71 := GenS5{F0: -190, F1: 1, F2: 391}
                _ = v71
            }
            fmt.Println(v1)
            for fi72 := 0; fi72 < 18; fi72++ {
                var v73 int8 = 49
                switch (fi72 % ((2147483647) | 1)) {
                case 125:
                    var v74 int32 = -1
                    _ = v74
                default:
                    v75 := GenS4{F0: 170}
                    _ = v75
                }
                if false {
                    v73 = (v73 + (v73 + v73))
                    var v76 uint16 = 563
                    v77 := GenS6{F0: 1}
                    _ = v76
                    _ = v77
                } else {
                    v78 := GenS4{F0: 65535}
                    v79 := GenS7{F0: 0, F1: v73, F2: v2}
                    _ = v78
                    _ = v79
                }
                _ = fi72
                _ = v73
            }
        }
    }
    if true {
        _ = len(v38)
    }
    var v80 float32 = float32(((float64((1e10 + float32(1e10))) + ((float64(v39) + (float64(1e10) + (3.14159 * v1))) - 3.14159)) + (float64(((float32((1.5 + 1e10)) / (3.14159 + (((((((v39 / v39) - (-2.25 + (1.5 / (v39 + 0.0)))) * v39) / v39) * 0.1) + (0.1 / v39)) - float32(v1)))) - 0.0)) - (((-2.25 - ((v1 - (v1 - (float64((v39 * 1.5)) - 3.14159))) + (v1 + (v1 - -2.25)))) + -2.25) * v1))))
    if ((v31 - v31) == v31) {
        for fi81 := 0; fi81 < 17; fi81++ {
            var v82 uint64 = 4294967295
            for fi83 := 0; fi83 < 8; fi83++ {
                if (true && (!(v4 < (v31 >> ((4294967295) & 63))))) {
                    var v84 string = v2
                    var v85 uint32 = v4
                    _ = v84
                    _ = v85
                }
                if true {
                    switch (v2 + v2) {
                    case "hello":
                        var v86 bool = v3
                        _ = v86
                    case "a":
                        v87 := GenS3{F0: 0}
                        _ = v87
                    case "":
                        v88 := v5[uint32(v82)]
                        _ = v88
                    }
                    v39 -= v39
                    v82 += 4294967295
                }
                _ = fi83
            }
            _ = fi81
            _ = v82
        }
    }
    if (42 <= 324) {
        var v89 uint = 0
        var v90 uint64 = 433
        v91 := GenS8{F0: 1}
        _ = v89
        _ = v90
        _ = v91
    }
    var v92 uint8 = 69
    if (255 == v92) {
        if v3 {
            var v93 int8 = -1
            _ = v93
        }
        v3 = (v92 == v92)
        fmt.Println(v3, v2, v1)
    }
    for fi94 := 0; fi94 < 17; fi94++ {
        for fi95 := 0; fi95 < 11; fi95++ {
            if (v1 <= float64(v80)) {
                v96 := GenS9{F0: v2}
                _ = v96
            } else {
                switch fi94 {
                case 304:
                    if (!v3) {
                        var v97 uint64 = 2
                        _ = v97
                    } else {
                        fmt.Println(v3)
                        switch (fi95 ^ fi95) {
                        case 42:
                            v39 = v80
                        }
                        v31 -= v4
                    }
                }
                var v98 bool = (1000 == 842)
                if (!(v4 <= v31)) {
                    v98 = true
                    v5[v31] = (fi95 % ((fi95) | 1))
                } else {
                    if len(v38) >= 2 {
                        v38 = v38[1:len(v38)]
                    }
                    v31 -= (^(v4))
                }
                _ = v98
            }
            v4 = 0
            _ = fi95
        }
        _ = fi94
    }
    if len(v38) > 0 {
        v38[0] = 9223372036854775807
    }
    switch -770 {
    case 32767:
        v5[(v31 << ((255) & 63))] = genFn0(1)
    case -441:
        v39 -= float32((v1 + (0.0 + 0.0)))
    default:
        for fi99 := 0; fi99 < 13; fi99++ {
            fmt.Println(v2, v1)
            if ((v92 << ((0) & 63)) != v92) {
                v100 := GenS3{F0: 9223372036854775807}
                v4 -= 2147483648
                v101 := GenS10{F0: uint16(v92), F1: 9223372036854775807}
                _ = v100
                _ = v101
            } else {
                v31 = 42
                for fi102 := 0; fi102 < 4; fi102++ {
                    v2 = v2
                    var v103 int32 = -2147483648
                    v104 := GenS11{F0: (v92 * v92)}
                    _ = fi102
                    _ = v103
                    _ = v104
                }
            }
            var v105 uint = 65535
            _ = fi99
            _ = v105
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v80:", v80)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v80
    _ = v92
}
