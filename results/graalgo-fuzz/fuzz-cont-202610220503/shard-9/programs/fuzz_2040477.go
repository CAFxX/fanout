package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 bool; F2 uint64 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32; F1 uint; F2 int64 }
type GenS9 struct { F0 uint64 }

func main() {
    var v1 uint16 = 1
    var v2 float32 = -2.25
    var v3 int8 = 127
    v4 := GenS1{F0: v3}
    v3 = v3
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 uint8 = 42
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint32 = 2
            if false {
                fmt.Println(v3, v1, fi5)
                if false {
                    if false {
                        v9 := GenS2{F0: 1}
                        if (1 == 272) {
                            if ((0 >= -1) && true) {
                                v10 := []int{-2147483648, (fi5 - 4294967295)}
                                var v11 string = "foo bar"
                                _ = v10
                                _ = v11
                            } else {
                                v12 := []int{fi5, fi7, -9223372036854775808, (fi7 ^ fi7)}
                                _ = v12
                            }
                            v13 := []int{}
                            var v14 int64 = -1
                            _ = v13
                            _ = v14
                        } else {
                            if (false || true) {
                                fmt.Println(fi5, v2)
                                var v15 int64 = -1
                                _ = v15
                            } else {
                                v6 = v6
                                v16 := GenS3{F0: 367, F1: (-1 >= 840), F2: 2}
                                _ = v16
                            }
                        }
                        v17 := GenS4{F0: ((42 >= 65535) || (v2 >= (0.1 - float32(0.0)))), F1: -702}
                        _ = v17
                        _ = v9
                    }
                    fmt.Println(fi7)
                    var v18 uint8 = 42
                    _ = v18
                } else {
                    if false {
                        v19 := GenS5{F0: 1, F1: 2147483647, F2: 0}
                        if ((253 <= (fi5 + (-(fi7)))) || (0 == 4294967295)) {
                            fmt.Println(v8, v8, fi7)
                            v2 = (0.0 / ((3.14159 * (v2 - float32((float64(v2) / 1e10)))) / v2))
                            var v20 int = 127
                            if (42 <= 0) {
                                v21 := []int{(fi5 << ((42) & 63))}
                                fmt.Println(v1, v6, fi5)
                                _ = v21
                            }
                            var v22 float32 = float32(-2.25)
                            v1 += (v1 << ((902) & 63))
                            _ = v20
                            _ = v22
                        }
                        _ = v19
                    } else {
                        v23 := []int{(fi7 & -2147483648)}
                        var v24 int16 = 688
                        if true {
                            v24 += 1
                            _ = len(v23)
                            v25 := GenS6{F0: v24, F1: 487}
                            var v26 int64 = 9223372036854775807
                            var v27 uint64 = 4294967295
                            _ = v25
                            _ = v26
                            _ = v27
                        } else {
                            var v28 uint8 = (v6 & v6)
                            v29 := GenS7{F0: (v28 | (^(v28)))}
                            _ = v28
                            _ = v29
                        }
                        _ = v23
                        _ = v24
                    }
                    var v30 uint64 = 42
                    v31 := []int{32768}
                    _ = v30
                    _ = v31
                }
            } else {
                v6 += uint8(v1)
                var v32 uint16 = (v1 + (^(v1)))
                var v33 float64 = float64((0.0 * (((0.0 / ((1.5 + -2.25) * (1e10 / (1.5 / float32((1e10 + float64(v2))))))) / 3.14159) - float32(3.14159))))
                _ = v32
                _ = v33
            }
            _ = fi7
            _ = v8
        }
        v34 := []int{fi5, fi5}
        _ = fi5
        _ = v34
        _ = v6
    }
    if false {
        v35 := GenS8{F0: 0, F1: 4294967295, F2: -9223372036854775808}
        _ = v35
    } else {
        v2 -= v2
        switch 9223372036854775807 {
        case -551:
            v3 = v3
        case 9223372036854775807:
            v36 := []int{9223372036854775807, 1}
            _ = v36
        default:
            fmt.Println(v2, v1)
        }
        switch 814 {
        case -2147483648:
            v3 = v3
        default:
            if (!false) {
                v37 := GenS7{F0: 214}
                _ = v37
            } else {
                for fi38 := 0; fi38 < 19; fi38++ {
                    var v39 uint8 = 1
                    for fi40 := 0; fi40 < 3; fi40++ {
                        v41 := []int{(fi38 >> ((151) & 63)), fi40}
                        fmt.Println(fi40)
                        if false {
                            switch fi38 {
                            case -967:
                                v3 = int8(v1)
                            }
                            var v42 int64 = 912
                            v41 = append(v41, (fi40 >> ((0) & 63)))
                            _ = v42
                        }
                        _ = fi40
                        _ = v41
                    }
                    _ = fi38
                    _ = v39
                }
                switch "foo bar" {
                case "":
                    var v43 string = "foo bar"
                    _ = v43
                }
            }
        }
    }
    var v44 int64 = -1
    v3 += v3
    fmt.Println(v1)
    if true {
        if false {
            v45 := []int{127, -629, 128}
            _ = v45
        }
        switch 32768 {
        case 32768:
            if false {
                if (!(-600 != -2147483648)) {
                    v46 := GenS1{F0: v3}
                    v1 -= (v1 >> ((890) & 63))
                    v47 := GenS7{F0: 231}
                    _ = v46
                    _ = v47
                } else {
                    if (!((v1 - 1) > (v1 * v1))) {
                        v1 += 255
                        var v48 float32 = v2
                        switch 9223372036854775807 {
                        case -230:
                            v49 := []int{}
                            _ = v49
                        default:
                            for fi50 := 0; fi50 < 13; fi50++ {
                                fmt.Println(v3)
                                v44 = 0
                                _ = fi50
                            }
                        }
                        _ = v48
                    } else {
                        switch "a" {
                        case "":
                            if (true && ((v2 - v2) >= v2)) {
                                for fi51 := 0; fi51 < 6; fi51++ {
                                    v2 += v2
                                    _ = fi51
                                }
                                var v52 uint64 = 0
                                if true {
                                    var v53 uint32 = 858
                                    if (0 > 4294967295) {
                                        var v54 string = ("\n\t" + "foo bar")
                                        v55 := []int{-1, 4294967295, 9223372036854775807, int(v53)}
                                        _ = v54
                                        _ = v55
                                    }
                                    fmt.Println(v53)
                                    _ = v53
                                }
                                _ = v52
                            }
                        case "\n\t":
                            if (!true) {
                                v1 = (v1 & 481)
                                v1 += (v1 + v1)
                                v3 += v3
                            } else {
                                v1 = 232
                                v2 -= v2
                                v3 -= (v3 % ((v3) | 1))
                            }
                        }
                        fmt.Println(v2, v2, v3)
                        var v56 int8 = (^(v3))
                        _ = v56
                    }
                }
                switch 32767 {
                case -418:
                    v3 -= -1
                case 4294967295:
                    if true {
                        v57 := []int{int(v3), 0, 821, 42}
                        if len(v57) > 0 {
                            v57[0] = 9223372036854775807
                        }
                        switch -918 {
                        case 1:
                            fmt.Println(v2, v3, v1)
                        default:
                            var v58 int8 = v3
                            _ = v58
                        }
                        _ = v57
                    }
                case 42:
                    fmt.Println(v3, v44)
                default:
                    v3 += (v3 + (-(v3)))
                }
                if false {
                    fmt.Println(v1, v2, v2)
                    v59 := []int{-139}
                    for fi60 := 0; fi60 < 2; fi60++ {
                        var v61 int = (^(fi60))
                        _ = fi60
                        _ = v61
                    }
                    _ = v59
                } else {
                    for fi62 := 0; fi62 < 11; fi62++ {
                        v63 := []int{(fi62 & fi62), fi62, fi62}
                        _ = fi62
                        _ = v63
                    }
                }
            }
        case 67:
            v2 -= v2
        case 9223372036854775807:
            for fi64 := 0; fi64 < 14; fi64++ {
                var v65 uint = 4294967295
                if true {
                    v3 = 1
                    for fi66 := 0; fi66 < 3; fi66++ {
                        if (-2.25 >= (v2 + 0.0)) {
                            v67 := GenS8{F0: 42, F1: 64, F2: v44}
                            _ = v67
                        } else {
                            v68 := []int{(fi66 / (((-(fi66))) | 1)), -128, fi66, fi66}
                            v3 += (v3 | v3)
                            _ = v68
                        }
                        var v69 int = fi66
                        _ = fi66
                        _ = v69
                    }
                    var v70 uint16 = (v1 | 2)
                    _ = v70
                }
                _ = fi64
                _ = v65
            }
        default:
            for fi71 := 0; fi71 < 12; fi71++ {
                v72 := GenS6{F0: -875, F1: 4294967295}
                if true {
                    var v73 uint32 = 292
                    _ = v73
                } else {
                    var v74 uint32 = 732
                    _ = v74
                }
                _ = fi71
                _ = v72
            }
        }
    } else {
        v3 += (v3 | v3)
        var v75 uint32 = 424
        var v76 int32 = 234
        _ = v75
        _ = v76
    }
    for fi77 := 0; fi77 < 8; fi77++ {
        for fi78 := 0; fi78 < 14; fi78++ {
            v2 -= v2
            var v79 int8 = -128
            var v80 int8 = (-(v3))
            _ = fi78
            _ = v79
            _ = v80
        }
        v81 := GenS1{F0: (v3 >> ((1) & 63))}
        _ = fi77
        _ = v81
    }
    for fi82 := 0; fi82 < 4; fi82++ {
        if true {
            var v83 int16 = 0
            switch (fi82 * fi82) {
            case -492:
                v84 := GenS9{F0: 298}
                _ = v84
            }
            _ = v83
        }
        if ((!(!(!true))) || false) {
            v85 := []int{32767, (-(fi82)), (fi82 & fi82), fi82}
            v44 = (v44 << ((65535) & 63))
            _ = v85
        }
        if true {
            v86 := GenS9{F0: 0}
            fmt.Println(fi82, v2)
            for fi87 := 0; fi87 < 20; fi87++ {
                v88 := []int{-2147483648, fi82, fi87, fi87}
                _ = fi87
                _ = v88
            }
            _ = v86
        } else {
            v89 := []int{}
            _ = v89
        }
        _ = fi82
    }
    for fi90 := 0; fi90 < 19; fi90++ {
        fmt.Println(v2)
        fmt.Println(v44, v3, v3)
        if false {
            for fi91 := 0; fi91 < 18; fi91++ {
                fmt.Println(fi91, v44)
                _ = fi91
            }
            for fi92 := 0; fi92 < 11; fi92++ {
                v93 := []int{fi90, fi90, (fi90 & fi92), (fi92 & -2147483648)}
                _ = fi92
                _ = v93
            }
        }
        _ = fi90
    }
    var v94 float32 = (0.0 + (v2 + -2.25))
    var v95 float32 = (float32((1e10 - 0.1)) * 3.14159)
    for fi96 := 0; fi96 < 10; fi96++ {
        v97 := []int{fi96, (-(fi96)), (fi96 + (fi96 | -550))}
        var v98 uint64 = 18446744073709551615
        v95 -= (v2 / (1.5 - v95))
        _ = fi96
        _ = v97
        _ = v98
    }
    switch 1 {
    case 1:
        v99 := GenS2{F0: 42}
        _ = v99
    default:
        v2 -= float32(1.5)
    }
    switch 32767 {
    case -60:
        for fi100 := 0; fi100 < 17; fi100++ {
            var v101 uint = 4294967295
            for fi102 := 0; fi102 < 5; fi102++ {
                v94 = (-2.25 * (1.5 + (v95 * v95)))
                _ = fi102
            }
            _ = fi100
            _ = v101
        }
    case 4294967295:
        var v103 uint64 = 2
        _ = v103
    }
    v3 = (-(v3))
    v95 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v94
    _ = v95
}
