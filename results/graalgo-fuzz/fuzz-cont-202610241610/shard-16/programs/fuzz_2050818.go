package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 bool; F1 bool; F2 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32; F1 int32; F2 int32 }
type GenS7 struct { F0 uint16; F1 string; F2 uint32 }
type GenS8 struct { F0 uint8; F1 int }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint8 }

func main() {
    var v1 float32 = 0.1
    var v2 float32 = (float32(1.5) + float32(0.1))
    v3 := make(map[uint32]bool)
    var v4 int32 = -2147483648
    v5 := make(map[uint32]string)
    var v6 int64 = 9223372036854775807
    var v7 float32 = (1e10 / 1.5)
    var v8 int8 = 100
    v8 = v8
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 int16 = -61
        v11 := GenS1{F0: -943}
        switch "hello" {
        case "\n\t":
            switch "" {
            case "hello":
                for fi12 := 0; fi12 < 8; fi12++ {
                    var v13 float32 = (((float32(3.14159) / float32(-2.25)) - v1) * (v1 + float32(float64(v1))))
                    var v14 uint64 = 2
                    if (true && false) {
                        v15 := GenS2{F0: (v8 << ((73) & 63)), F1: (v6 & (v6 + (v6 - v6)))}
                        if ((((1.5 / 0.1) / ((float64(0.1) - ((0.1 + 0.0) + float64(v7))) - 3.14159)) / (((3.14159 / float64(((v1 - v7) / v1))) - 0.1) / (((0.0 - 0.1) * 1e10) + -2.25))) == 0.1) {
                            v8 += 100
                        }
                        _ = v15
                    } else {
                        v16 := GenS3{F0: (v14 - 1), F1: (v10 >> ((792) & 63))}
                        _ = v16
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            case "":
                var v17 uint64 = 518
                _ = v17
            }
        case "a":
            switch (fi9 ^ fi9) {
            case 0:
                for fi18 := 0; fi18 < 17; fi18++ {
                    v1 = 0.0
                    v1 -= 0.1
                    v19 := GenS2{F0: v8, F1: -44}
                    _ = fi18
                    _ = v19
                }
            case 4294967295:
                fmt.Println(v4, fi9, v1)
            default:
                if true {
                    if true {
                        v2 -= (v7 * v1)
                    } else {
                        v20 := GenS3{F0: 651, F1: (v10 - (-(v10)))}
                        var v21 uint16 = 0
                        _ = v20
                        _ = v21
                    }
                } else {
                    fmt.Println(v1, v1)
                    for fi22 := 0; fi22 < 15; fi22++ {
                        if true {
                            delete(v3, 4294967295)
                            switch 42 {
                            case 1:
                                v23 := []int{}
                                _ = v23
                            case -1:
                                var v24 uint16 = 390
                                _ = v24
                            }
                        } else {
                            v25 := v5[2147483648]
                            var v26 uint16 = 0
                            _ = v25
                            _ = v26
                        }
                        v4 += (v4 * 2147483647)
                        _ = fi22
                    }
                    if (255 > 128) {
                        fmt.Println(fi9, v8)
                        if (true || true) {
                            var v27 uint64 = 875
                            _ = v27
                        }
                        v28 := []int{}
                        _ = v28
                    }
                }
            }
        case "hello":
            fmt.Println(v2, v4)
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    switch "a" {
    case "\n\t":
        v29 := []int{1}
        _ = v29
    case "":
        v2 += v7
    case "hello":
        if (65535 != 32767) {
            var v30 uint16 = 2
            _ = v30
        }
    }
    if (42 <= 4294967295) {
        switch -1 {
        case 277:
            switch -2147483648 {
            case -164:
                v6 -= v6
            case -877:
                fmt.Println(v2, v4, v4)
            default:
                if true {
                    if (!false) {
                        for fi31 := 0; fi31 < 13; fi31++ {
                            v32 := GenS4{F0: true, F1: true, F2: int64(fi31)}
                            switch 32768 {
                            case 4294967295:
                                v33 := GenS3{F0: 42, F1: -993}
                                _ = v33
                            }
                            delete(v5, 65535)
                            _ = fi31
                            _ = v32
                        }
                    }
                } else {
                    var v34 string = "a"
                    _ = v34
                }
            }
        case -9223372036854775808:
            switch "\n\t" {
            case "\n\t":
                v4 += -2147483648
            case "":
                if false {
                    v7 = v1
                    for fi35 := 0; fi35 < 2; fi35++ {
                        v36 := v3[0]
                        for fi37 := 0; fi37 < 13; fi37++ {
                            switch (fi37 << ((4294967295) & 63)) {
                            case -2147483648:
                                var v38 uint = 903
                                _ = v38
                            default:
                                var v39 uint64 = 1
                                _ = v39
                            }
                            _ = fi37
                        }
                        if (!((v6 >> ((255) & 63)) != v6)) {
                            v5[735] = ((("foo bar" + "") + "foo bar") + ("" + "foo bar"))
                        }
                        _ = fi35
                        _ = v36
                    }
                    var v40 int = 418
                    _ = v40
                } else {
                    var v41 int32 = v4
                    var v42 float64 = float64(v7)
                    if (249 == -400) {
                        for fi43 := 0; fi43 < 13; fi43++ {
                            for fi44 := 0; fi44 < 20; fi44++ {
                                v7 = v2
                                v45 := GenS5{F0: (-(v41))}
                                v46 := []int{}
                                switch fi44 {
                                case -884:
                                    var v47 int32 = 0
                                    _ = v47
                                case -9223372036854775808:
                                    if len(v46) >= 2 {
                                        v46 = v46[1:len(v46)]
                                    }
                                default:
                                    if true {
                                        v48 := GenS1{F0: v6}
                                        v8 += v8
                                        _ = v48
                                    } else {
                                        var v49 bool = (!true)
                                        var v50 int8 = (v8 / (((v8 ^ v8)) | 1))
                                        _ = v49
                                        _ = v50
                                    }
                                }
                                var v51 uint8 = 255
                                _ = fi44
                                _ = v45
                                _ = v46
                                _ = v51
                            }
                            _ = fi43
                        }
                    } else {
                        var v52 int = 353
                        for fi53 := 0; fi53 < 8; fi53++ {
                            v54 := []int{(v52 & (fi53 >> ((255) & 63))), -624}
                            v55 := GenS5{F0: 42}
                            var v56 uint64 = 119
                            _ = fi53
                            _ = v54
                            _ = v55
                            _ = v56
                        }
                        _ = v52
                    }
                    _ = v41
                    _ = v42
                }
            default:
                var v57 int = -565
                _ = v57
            }
        default:
            v4 = (v4 >> ((111) & 63))
        }
        switch -9223372036854775808 {
        case 4294967295:
            fmt.Println(v8, v4, v6)
        default:
            if (!false) {
                v2 = v1
                v4 = v4
            }
        }
        if true {
            if false {
                for fi58 := 0; fi58 < 19; fi58++ {
                    switch "\n\t" {
                    case "\n\t":
                        if (int8(fi58) < (v8 + -1)) {
                            var v59 uint64 = 0
                            fmt.Println(v4, v8)
                            for fi60 := 0; fi60 < 19; fi60++ {
                                switch ("" + "hello") {
                                case "foo bar":
                                    delete(v3, 2147483648)
                                case "hello":
                                    var v61 uint16 = 255
                                    _ = v61
                                case "a":
                                    v62 := GenS2{F0: v8, F1: (v6 & (v6 - (v6 % ((v6) | 1))))}
                                    _ = v62
                                default:
                                    v7 += (1e10 - ((-2.25 - v2) + 0.0))
                                }
                                fmt.Println(v1, v59, v1)
                                fmt.Println(v8, fi58)
                                _ = fi60
                            }
                            _ = v59
                        } else {
                            if true {
                                v2 -= v1
                                var v63 uint = 4294967295
                                var v64 uint8 = 128
                                _ = v63
                                _ = v64
                            } else {
                                switch fi58 {
                                case -989:
                                    v2 = 0.1
                                case 2147483647:
                                    v65 := GenS4{F0: (0 != v8), F1: (true || (20 == 4294967295)), F2: 0}
                                    _ = v65
                                }
                                var v66 int64 = v6
                                _ = v66
                            }
                            _ = len(v5)
                            var v67 int64 = (v6 + -1)
                            _ = v67
                        }
                    case "hello":
                        var v68 uint16 = 2
                        _ = v68
                    case "foo bar":
                        if (v8 >= (v8 | -75)) {
                            for fi69 := 0; fi69 < 20; fi69++ {
                                if true {
                                    v70 := []int{fi69, (fi69 & 127), 331}
                                    var v71 bool = (476 <= 2)
                                    v72 := GenS6{F0: 4294967295, F1: v4, F2: -305}
                                    fmt.Println(v8)
                                    var v73 string = "\n\t"
                                    _ = v70
                                    _ = v71
                                    _ = v72
                                    _ = v73
                                } else {
                                    var v74 uint64 = 79
                                    if false {
                                        if (v74 > v74) {
                                            fmt.Println(v4, v7, fi58)
                                        }
                                    } else {
                                        v75 := []int{(fi69 % (((-(fi58))) | 1)), 128}
                                        if true {
                                            var v76 uint32 = 2147483647
                                            if true {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v77 := GenS7{F0: 42, F1: "foo bar", F2: (v76 + (v76 >> ((42) & 63)))}
                                                    v7 -= (((1e10 * 1.5) + (v7 * (1.5 - v1))) / v1)
                                                    _ = v77
                                                case "\n\t":
                                                    var v78 float64 = 1e10
                                                    _ = v78
                                                case "hello":
                                                    v79 := GenS8{F0: 0, F1: -548}
                                                    _ = v79
                                                }
                                                v6 = (v6 / ((v6) | 1))
                                                v80 := v5[v76]
                                                _ = v80
                                            } else {
                                                fmt.Println(v2)
                                                var v81 float32 = v7
                                                if true {
                                                    v82 := GenS8{F0: 127, F1: 465}
                                                    _ = v82
                                                } else {
                                                    v1 = v2
                                                    if false {
                                                        v75 = append(v75, (fi58 * fi69))
                                                    }
                                                    v76 -= uint32(v74)
                                                }
                                                _ = v81
                                            }
                                            v8 = int8(v4)
                                            _ = v76
                                        } else {
                                            if len(v75) > 0 {
                                                v75[0] = fi69
                                            }
                                            v83 := v3[65535]
                                            _ = v83
                                        }
                                        _ = v75
                                    }
                                    _ = v74
                                }
                                var v84 int32 = (v4 | (^(v4)))
                                if false {
                                    var v85 uint16 = 2
                                    fmt.Println(v2, v84, v8)
                                    _ = v85
                                }
                                fmt.Println(v7, v1, fi58)
                                _ = fi69
                                _ = v84
                            }
                            var v86 int = 9223372036854775807
                            _ = v86
                        } else {
                            v7 = (v2 * v2)
                            var v87 string = "\n\t"
                            _ = v87
                        }
                    }
                    var v88 float64 = 3.14159
                    v89 := GenS9{F0: (v4 * (v4 + v4))}
                    _ = fi58
                    _ = v88
                    _ = v89
                }
                v90 := v3[581]
                _ = v90
            }
        } else {
            v91 := GenS10{F0: 1}
            _ = v91
        }
    }
    v2 = 0.0
    var v92 uint = 2
    v6 = (v6 % (((v6 + (v6 << (((v92 * (v92 % (((v92 + v92)) | 1)))) & 63)))) | 1))
    if true {
        for fi93 := 0; fi93 < 2; fi93++ {
            if false {
                v94 := GenS8{F0: 42, F1: (fi93 & int(v4))}
                v95 := v3[689]
                var v96 uint8 = 127
                _ = v94
                _ = v95
                _ = v96
            }
            _ = fi93
        }
    } else {
        var v97 uint32 = 2
        for fi98 := 0; fi98 < 2; fi98++ {
            fmt.Println(v92, v7, fi98)
            var v99 int64 = -1
            _ = fi98
            _ = v99
        }
        v3[(v97 % (((v97 * v97)) | 1))] = (!true)
        _ = v97
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v92
}
