package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 uint8; F1 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int16; F1 int64; F2 int }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 bool; F1 uint }
type GenS12 struct { F0 uint8 }
func genFn0(p0 bool, p1 uint16) int16 {
    v6 := GenS1{F0: 42, F1: 4294967295, F2: -737}
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{(fi7 / ((fi7) | 1)), (fi7 / ((fi7) | 1))}
        if (255 >= 2) {
            switch fi7 {
            case 127:
                v9 := GenS2{F0: p1, F1: (p1 ^ 136), F2: (false && p0)}
                _ = v9
            case -2147483648:
                p1 += (p1 << ((255) & 63))
            }
            switch "" {
            case "hello":
                if false {
                    v10 := GenS2{F0: 2, F1: (p1 >> ((2) & 63)), F2: (!p0)}
                    _ = v10
                } else {
                    var v11 float64 = 1.5
                    var v12 string = ("foo bar" + "")
                    fmt.Println(v12, p1)
                    _ = v11
                    _ = v12
                }
            default:
                v13 := GenS3{F0: (^(p1))}
                _ = v13
            }
            for fi14 := 0; fi14 < 3; fi14++ {
                p0 = false
                var v15 int = 128
                v16 := GenS2{F0: (p1 ^ 1), F1: 1, F2: (false || (-641 >= -870))}
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        _ = fi7
        _ = v8
    }
    v17 := []int{}
    if len(v17) > 0 {
        v17[0] = -403
    }
    p0 = (!false)
    _ = p0
    _ = p1
    _ = v17
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = ""
    var v3 int = 128
    var v4 uint16 = 1
    v5 := make(map[int]uint32)
    for fi18 := 0; fi18 < 17; fi18++ {
        switch v2 {
        case "":
            if (v4 < v4) {
                var v19 uint8 = 174
                _ = v19
            } else {
                v20 := GenS2{F0: v4, F1: (v4 << ((255) & 63)), F2: (0 < int64(v4))}
                if (!false) {
                    if false {
                        var v21 bool = false
                        var v22 string = v2
                        _ = v21
                        _ = v22
                    }
                }
                v23 := v1[(fi18 ^ fi18)]
                _ = v20
                _ = v23
            }
        }
        var v24 int32 = 0
        _ = fi18
        _ = v24
    }
    v25 := GenS2{F0: 970, F1: v4, F2: (false && (1000 > genFn0(false, 0)))}
    v26 := GenS4{F0: 1, F1: 100}
    fmt.Println(v2, v4)
    var v27 int16 = 257
    fmt.Println(v27, v2)
    v28 := GenS3{F0: (^(v4))}
    for fi29 := 0; fi29 < 12; fi29++ {
        switch v2 {
        case "a":
            for fi30 := 0; fi30 < 9; fi30++ {
                switch int(v4) {
                case 32767:
                    if (false || ((v4 << ((2) & 63)) < (^(v4)))) {
                        if true {
                            switch v2 {
                            case "a":
                                var v31 int16 = 42
                                _ = v31
                            case "foo bar":
                                switch 814 {
                                case 9223372036854775807:
                                    v32 := GenS5{F0: 1}
                                    _ = v32
                                case 42:
                                    v33 := GenS1{F0: 4294967295, F1: 4294967295, F2: -9223372036854775808}
                                    _ = v33
                                case 32768:
                                    var v34 int32 = 42
                                    _ = v34
                                default:
                                    v35 := GenS6{F0: (fi29 % ((fi29) | 1)), F1: 255}
                                    _ = v35
                                }
                            case "":
                                v2 = v2
                                var v36 int64 = 0
                                v2 = v2
                                if true {
                                    v27 -= (v27 / (((v27 - v27)) | 1))
                                    _ = len(v1)
                                }
                                v37 := GenS7{F0: 224, F1: v36}
                                switch (fi30 % ((v3) | 1)) {
                                case -276:
                                    if (!true) {
                                        switch v2 {
                                        case "\n\t":
                                            v27 = 1000
                                        default:
                                            var v38 float64 = (-2.25 / (1.5 * 3.14159))
                                            _ = v38
                                        }
                                    } else {
                                        v39 := GenS1{F0: 1, F1: 255, F2: -9223372036854775808}
                                        switch 9223372036854775807 {
                                        case -282:
                                            fmt.Println(v36, fi29)
                                            if true {
                                                var v40 uint32 = 1
                                                if false {
                                                    var v41 float32 = float32(0.0)
                                                    v3 += int(v27)
                                                    _ = v41
                                                }
                                                _ = v40
                                            } else {
                                                v2 = (v2 + v2)
                                            }
                                            if ("a" < (v2 + v2)) {
                                                v42 := GenS8{F0: 0}
                                                _ = v42
                                            } else {
                                                v4 += 2
                                            }
                                        default:
                                            var v43 uint32 = 267
                                            _ = v43
                                        }
                                        v44 := []int{}
                                        _ = v39
                                        _ = v44
                                    }
                                case -2147483648:
                                    delete(v5, (fi29 % (((-(fi29))) | 1)))
                                }
                                _ = v36
                                _ = v37
                            }
                        }
                        if false {
                            v45 := []int{0, 999, (v3 % ((fi30) | 1)), fi30}
                            v46 := GenS5{F0: 0}
                            _ = v45
                            _ = v46
                        } else {
                            if (!false) {
                                v47 := []int{(fi30 + fi29)}
                                var v48 int32 = 2147483647
                                fmt.Println(v2)
                                var v49 int8 = -75
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                            var v50 uint16 = 2
                            _ = v50
                        }
                    }
                }
                _ = fi30
            }
        case "hello":
            switch 32767 {
            case -857:
                var v51 uint16 = 32767
                _ = v51
            default:
                switch 32767 {
                case -9223372036854775808:
                    fmt.Println(fi29, v27, v27)
                default:
                    if false {
                        v3 = (v3 * (fi29 - 2147483647))
                        v52 := []int{42, v3}
                        v27 = genFn0(((v4 ^ v4) != v4), (v4 & v4))
                        _ = v52
                    }
                }
            }
        }
        if ((fi29 | v3) < v3) {
            if false {
                if false {
                    if false {
                        for fi53 := 0; fi53 < 20; fi53++ {
                            var v54 uint32 = 213
                            v2 = "foo bar"
                            _ = fi53
                            _ = v54
                        }
                    }
                    if false {
                        v55 := GenS6{F0: -2147483648, F1: 4294967295}
                        if true {
                            switch "a" {
                            case "\n\t":
                                v56 := GenS3{F0: v4}
                                _ = v56
                            case "foo bar":
                                v5[(fi29 % (((fi29 % ((4294967295) | 1))) | 1))] = 2147483648
                            case "":
                                v57 := []int{fi29, v3}
                                _ = v57
                            default:
                                for fi58 := 0; fi58 < 7; fi58++ {
                                    v59 := []int{}
                                    v4 += v4
                                    v3 -= 2147483647
                                    switch "hello" {
                                    case "hello":
                                        var v60 int32 = 1
                                        _ = v60
                                    case "foo bar":
                                        v61 := GenS9{F0: -1, F1: 0, F2: fi29}
                                        _ = v61
                                    case "a":
                                        fmt.Println(v2)
                                    }
                                    if (true || false) {
                                        var v62 int8 = 100
                                        switch v2 {
                                        case "\n\t":
                                            v63 := GenS5{F0: -2147483648}
                                            var v64 float64 = float64(float32(0.1))
                                            _ = v63
                                            _ = v64
                                        case "foo bar":
                                            var v65 int = 0
                                            _ = v65
                                        case "a":
                                            v66 := GenS8{F0: 65535}
                                            _ = v66
                                        }
                                        v5[(fi29 & fi29)] = 2147483648
                                        _ = v62
                                    }
                                    if false {
                                        v4 -= (v4 % (((v4 * (v4 + v4))) | 1))
                                        switch (fi58 + 0) {
                                        case 587:
                                            v3 += (-(fi58))
                                        case 127:
                                            if false {
                                                v67 := GenS8{F0: v4}
                                                _ = v67
                                            } else {
                                                fmt.Println(v3, fi58)
                                                v3 = v3
                                                if (true && (int64(v4) != -9223372036854775808)) {
                                                    if true {
                                                        switch (("\n\t" + ((v2 + v2) + (("" + v2) + v2))) + (v2 + v2)) {
                                                        case "":
                                                            v4 = v4
                                                            v68 := GenS10{F0: 127}
                                                            fmt.Println(v2, v3)
                                                            if (4294967295 > 0) {
                                                                var v69 uint64 = 1
                                                                v3 += fi29
                                                                _ = v69
                                                            } else {
                                                                v2 = v2
                                                                switch fi58 {
                                                                case 32767:
                                                                    v70 := GenS1{F0: 42, F1: 42, F2: 0}
                                                                    _ = v70
                                                                }
                                                                v27 = (v27 << ((49) & 63))
                                                            }
                                                            _ = v68
                                                        case "hello":
                                                            v5[1] = 4294967295
                                                        case "a":
                                                            v71 := GenS2{F0: (v4 % ((v4) | 1)), F1: (v4 * v4), F2: true}
                                                            _ = v71
                                                        default:
                                                            v3 -= v3
                                                            if (false && (!(0 <= 65535))) {
                                                                switch (v2 + v2) {
                                                                case "foo bar":
                                                                    if (784 < 717) {
                                                                        fmt.Println(v3, v2, v27)
                                                                        v27 = v27
                                                                        v2 = v2
                                                                    }
                                                                case "\n\t":
                                                                    v2 = (v2 + v2)
                                                                }
                                                                var v72 uint8 = 0
                                                                _ = v72
                                                            }
                                                            fmt.Println(v2, fi29)
                                                            var v73 bool = ((v4 | (v4 + 35)) < 443)
                                                            var v74 float64 = -2.25
                                                            var v75 uint16 = v4
                                                            _ = v73
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                        v76 := GenS11{F0: true, F1: 803}
                                                        if false {
                                                            v77 := GenS10{F0: 82}
                                                            v78 := GenS12{F0: 42}
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                        if false {
                                                            var v79 uint16 = v4
                                                            _ = v79
                                                        } else {
                                                            var v80 uint16 = 255
                                                            var v81 int16 = (-(v27))
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                        var v82 int = (fi58 - -1)
                                                        _ = v76
                                                        _ = v82
                                                    } else {
                                                        var v83 int8 = 100
                                                        var v84 uint16 = (v4 | 32767)
                                                        var v85 uint32 = 267
                                                        _ = v83
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                    v86 := GenS6{F0: v3, F1: uint(v4)}
                                                    v59 = append(v59, (fi58 | (fi58 % ((128) | 1))))
                                                    _ = v86
                                                } else {
                                                    if (!true) {
                                                        switch 1 {
                                                        case 9223372036854775807:
                                                            var v87 int32 = 2147483647
                                                            _ = v87
                                                        case 32768:
                                                            fmt.Println(v2)
                                                        default:
                                                            fmt.Println(v4, v4, v3)
                                                        }
                                                    } else {
                                                        v4 += (^(v4))
                                                        switch (v2 + "foo bar") {
                                                        case "foo bar":
                                                            v3 = fi29
                                                        case "\n\t":
                                                            switch (-(v3)) {
                                                            case -1:
                                                                v88 := GenS3{F0: (^(v4))}
                                                                _ = v88
                                                            default:
                                                                var v89 bool = false
                                                                _ = v89
                                                            }
                                                        }
                                                        v3 -= fi58
                                                    }
                                                    if true {
                                                        v2 = (v2 + ("" + v2))
                                                        fmt.Println(fi29)
                                                        var v90 string = v2
                                                        _ = v90
                                                    }
                                                    v91 := GenS2{F0: (v4 - v4), F1: (v4 + v4), F2: true}
                                                    var v92 string = (v2 + (v2 + v2))
                                                    var v93 uint = 18446744073709551615
                                                    _ = v91
                                                    _ = v92
                                                    _ = v93
                                                }
                                            }
                                        default:
                                            var v94 int16 = -714
                                            _ = v94
                                        }
                                    }
                                    v27 = (v27 & v27)
                                    _ = len(v59)
                                    var v95 float32 = 1e10
                                    if false {
                                        var v96 float32 = (v95 + (float32(-2.25) + v95))
                                        var v97 int16 = genFn0((true || false), v4)
                                        _ = v96
                                        _ = v97
                                    }
                                    _ = fi58
                                    _ = v59
                                    _ = v95
                                }
                            }
                            fmt.Println(v4, fi29, v3)
                        } else {
                            v98 := GenS7{F0: 255, F1: 42}
                            var v99 int = (v3 & fi29)
                            _ = v98
                            _ = v99
                        }
                        var v100 float32 = (1e10 - (0.1 + 3.14159))
                        _ = v100
                        _ = v55
                    } else {
                        for fi101 := 0; fi101 < 8; fi101++ {
                            if false {
                                v102 := GenS6{F0: fi29, F1: 18446744073709551615}
                                var v103 float32 = (float32(3.14159) / float32((3.14159 + 1.5)))
                                _ = v102
                                _ = v103
                            } else {
                                if ((true || false) && true) {
                                    v2 = (v2 + (v2 + "hello"))
                                }
                            }
                            var v104 uint8 = 1
                            _ = fi101
                            _ = v104
                        }
                        v105 := []int{}
                        fmt.Println(v27, fi29, v2)
                        _ = v105
                    }
                    v106 := GenS3{F0: v4}
                    _ = v106
                } else {
                    switch v3 {
                    case -94:
                        var v107 int32 = 1
                        _ = v107
                    }
                }
            }
            if (1.5 > ((0.0 + 1.5) / float64((0.1 + -2.25)))) {
                v4 = (v4 & 1)
            }
        }
        _ = fi29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
