package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int }
type GenS2 struct { F0 int; F1 int16; F2 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS7 struct { F0 bool; F1 int32; F2 uint16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int; F1 int16 }
type GenS10 struct { F0 bool; F1 int; F2 int16 }
func genFn0(p0 int8, p1 uint, p2 bool) bool {
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            p1 = (^(p1))
            switch ("foo bar" + ("\n\t" + "\n\t")) {
            case "\n\t":
                fmt.Println(p1)
            case "hello":
                switch (fi8 & (fi8 + fi8)) {
                case 13:
                    var v9 uint16 = 65535
                    _ = v9
                }
                switch (fi7 | fi8) {
                case -665:
                    if true {
                        if (p1 <= 422) {
                            p1 = p1
                        } else {
                            v10 := []int{(fi7 & int(p1)), (fi7 * -1)}
                            _ = v10
                        }
                    }
                }
            }
            _ = fi8
        }
        v11 := GenS1{F0: 0, F1: 939, F2: (^(fi7))}
        _ = fi7
        _ = v11
    }
    var v12 int = 128
    switch v12 {
    case 650:
        fmt.Println(p2, p0, p1)
    case 32767:
        if true {
            var v13 uint8 = 0
            _ = v13
        } else {
            for fi14 := 0; fi14 < 20; fi14++ {
                v15 := GenS2{F0: -2147483648, F1: 1, F2: (p0 & (p0 | -90))}
                v16 := GenS1{F0: 32767, F1: -32768, F2: 1}
                v17 := []int{v12, 550, fi14, (v12 - -223)}
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
            v18 := GenS3{F0: 0}
            v19 := GenS2{F0: (v12 | -1), F1: 1, F2: 127}
            _ = v18
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        v21 := []int{}
        if len(v21) > 0 {
            v21[0] = (v12 | v12)
        }
        p2 = false
        _ = fi20
        _ = v21
    }
    switch (^(v12)) {
    case 147:
        v22 := GenS4{F0: 856}
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return ((v12 ^ v12) > -1)
}

func main() {
    var v1 float64 = float64((-2.25 / 3.14159))
    var v2 string = ((("hello" + "a") + "") + "foo bar")
    var v3 int = -800
    v4 := make(map[uint32]bool)
    var v5 int8 = -128
    var v6 int8 = 42
    v23 := []int{-224, 32767, v3, (^(v3))}
    v24 := GenS2{F0: int(v6), F1: -1, F2: 53}
    _ = len(v23)
    var v25 int64 = 9223372036854775807
    v26 := GenS5{F0: 933}
    switch v3 {
    case 32768:
        v27 := GenS3{F0: 42}
        _ = v27
    case -9223372036854775808:
        switch (v2 + v2) {
        case "a":
            switch -2147483648 {
            case 4294967295:
                var v28 int32 = -400
                _ = v28
            case 128:
                delete(v4, 65535)
            default:
                if false {
                    for fi29 := 0; fi29 < 18; fi29++ {
                        v30 := GenS1{F0: 0, F1: -32768, F2: -9223372036854775808}
                        for fi31 := 0; fi31 < 20; fi31++ {
                            v32 := GenS3{F0: 0}
                            _ = len(v23)
                            v33 := GenS4{F0: 4294967295}
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 19; fi34++ {
                            fmt.Println(fi34, v6, v3)
                            _ = fi34
                        }
                        _ = fi29
                        _ = v30
                    }
                    v3 = 2147483647
                }
            }
        case "foo bar":
            for fi35 := 0; fi35 < 7; fi35++ {
                for fi36 := 0; fi36 < 10; fi36++ {
                    var v37 uint8 = 2
                    switch fi35 {
                    case -100:
                        delete(v4, 2147483648)
                    case -416:
                        v37 = v37
                    case 127:
                        v2 = v2
                    }
                    _ = fi36
                    _ = v37
                }
                _ = fi35
            }
        }
    }
    if (100 < -111) {
        for fi38 := 0; fi38 < 14; fi38++ {
            if false {
                var v39 int = (fi38 ^ (v3 ^ fi38))
                var v40 int16 = 1
                _ = v39
                _ = v40
            }
            if true {
                var v41 int16 = -1
                if (fi38 >= 42) {
                    v42 := GenS6{F0: 320, F1: 2147483647, F2: 86}
                    _ = v42
                }
                var v43 string = v2
                _ = v41
                _ = v43
            } else {
                v1 -= (1.5 + (float64(((float32((1e10 + (v1 / (1.5 * v1)))) + float32(v1)) * 1.5)) * (v1 - 0.0)))
                for fi44 := 0; fi44 < 1; fi44++ {
                    v1 += v1
                    switch v2 {
                    case "a":
                        v45 := GenS7{F0: false, F1: -2147483648, F2: 427}
                        _ = v45
                    default:
                        v5 += v6
                    }
                    _ = fi44
                }
                v46 := GenS6{F0: 0, F1: -2147483648, F2: 76}
                _ = v46
            }
            _ = fi38
        }
        if ((4294967295 != 2) && (1.5 == 0.0)) {
            var v47 int8 = (v5 | (v6 / ((v5) | 1)))
            v48 := GenS8{F0: 32767}
            if true {
                var v49 uint8 = 127
                v50 := GenS5{F0: v3}
                _ = v49
                _ = v50
            } else {
                var v51 uint = 4294967295
                if (0.1 <= 1e10) {
                    if true {
                        var v52 float64 = float64(0.1)
                        for fi53 := 0; fi53 < 1; fi53++ {
                            fmt.Println(v2, v47)
                            _ = len(v4)
                            v52 -= 1e10
                            _ = fi53
                        }
                        _ = v52
                    }
                    var v54 string = v2
                    if (false || (true || (!false))) {
                        var v55 uint32 = 2
                        v2 = "\n\t"
                        fmt.Println(v3, v6, v55)
                        _ = v55
                    } else {
                        v56 := GenS9{F0: (v3 ^ v3), F1: 0}
                        v51 = 665
                        switch v54 {
                        case "a":
                            for fi57 := 0; fi57 < 12; fi57++ {
                                if (388 == 42) {
                                    delete(v4, 2147483647)
                                    v6 = (v5 ^ (-(v47)))
                                    for fi58 := 0; fi58 < 6; fi58++ {
                                        var v59 uint64 = 42
                                        _ = fi58
                                        _ = v59
                                    }
                                }
                                var v60 uint = (v51 >> ((0) & 63))
                                _ = fi57
                                _ = v60
                            }
                        case "\n\t":
                            switch -2147483648 {
                            case 32768:
                                if (int64(v51) <= v25) {
                                    v6 -= -128
                                    var v61 float64 = float64(3.14159)
                                    _ = v61
                                } else {
                                    if true {
                                        v4[2] = (2147483647 == 2147483647)
                                        for fi62 := 0; fi62 < 4; fi62++ {
                                            if true {
                                                if (127 <= 42) {
                                                    for fi63 := 0; fi63 < 19; fi63++ {
                                                        v54 = (v54 + v54)
                                                        _ = fi63
                                                    }
                                                    v64 := GenS8{F0: 42}
                                                    _ = v64
                                                }
                                                var v65 int64 = (v25 & 720)
                                                _ = v65
                                            }
                                            var v66 int32 = -727
                                            _ = fi62
                                            _ = v66
                                        }
                                        v67 := GenS10{F0: false, F1: (v3 ^ (v3 - (v3 | 4294967295))), F2: -32768}
                                        _ = v67
                                    } else {
                                        var v68 float32 = (0.1 / 1.5)
                                        _ = v68
                                    }
                                    v69 := GenS10{F0: false, F1: (v3 + v3), F2: 138}
                                    _ = v69
                                }
                            case -9223372036854775808:
                                v1 -= float64(3.14159)
                            case 134:
                                v54 = "a"
                            }
                        case "":
                            v3 = v3
                        }
                        _ = v56
                    }
                    _ = v54
                } else {
                    for fi70 := 0; fi70 < 18; fi70++ {
                        for fi71 := 0; fi71 < 6; fi71++ {
                            switch v2 {
                            case "a":
                                v1 = (float64(1.5) / 3.14159)
                            case "":
                                v2 = "a"
                            default:
                                v72 := GenS7{F0: (0.0 != 0.0), F1: 0, F2: 1}
                                _ = v72
                            }
                            _ = fi71
                        }
                        fmt.Println(fi70)
                        v2 = ("\n\t" + (((v2 + v2) + v2) + (v2 + v2)))
                        _ = fi70
                    }
                    _ = len(v4)
                    var v73 int64 = v25
                    _ = v73
                }
                _ = v51
            }
            _ = v47
            _ = v48
        } else {
            var v74 uint16 = 656
            if (v74 <= 521) {
                fmt.Println(v5, v6, v2)
                for fi75 := 0; fi75 < 3; fi75++ {
                    var v76 int = v3
                    v77 := v4[1]
                    _ = fi75
                    _ = v76
                    _ = v77
                }
                for fi78 := 0; fi78 < 17; fi78++ {
                    var v79 uint8 = 1
                    _ = fi78
                    _ = v79
                }
            }
            _ = v74
        }
    } else {
        v2 = "hello"
        v80 := v4[2]
        var v81 uint8 = 2
        _ = v80
        _ = v81
    }
    if false {
        for fi82 := 0; fi82 < 18; fi82++ {
            v4[529] = false
            v23 = append(v23, fi82)
            for fi83 := 0; fi83 < 2; fi83++ {
                fmt.Println(v25)
                switch (fi82 >> ((0) & 63)) {
                case 9223372036854775807:
                    fmt.Println(v3)
                case 42:
                    if (false && (255 >= 0)) {
                        var v84 int16 = 0
                        v84 += (^(v84))
                        _ = v84
                    }
                    var v85 float64 = v1
                    v86 := GenS2{F0: fi83, F1: 0, F2: v6}
                    _ = v85
                    _ = v86
                case 128:
                    var v87 uint64 = 4294967295
                    switch (v2 + ((v2 + v2) + v2)) {
                    case "a":
                        _ = len(v4)
                    case "":
                        v5 += (v6 * (v5 + 114))
                    }
                    _ = v87
                default:
                    var v88 int8 = v6
                    _ = v88
                }
                _ = fi83
            }
            _ = fi82
        }
        v5 = v6
    }
    fmt.Println(v5, v5, v25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
