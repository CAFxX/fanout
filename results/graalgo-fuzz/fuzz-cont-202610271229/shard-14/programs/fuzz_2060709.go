package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int; F1 uint64; F2 bool }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 int; F1 int32; F2 uint64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{4294967295}
    var v3 uint64 = 753
    var v4 bool = (-1 > -459)
    var v5 uint32 = 0
    var v6 int64 = -1
    for fi7 := 0; fi7 < 13; fi7++ {
        switch ("hello" + "foo bar") {
        case "a":
            if v4 {
                var v8 float64 = (-2.25 * 0.1)
                v8 = 3.14159
                v9 := GenS1{F0: v5}
                _ = v8
                _ = v9
            } else {
                if v4 {
                    for fi10 := 0; fi10 < 14; fi10++ {
                        v4 = (570 > 65535)
                        var v11 int32 = 42
                        _ = fi10
                        _ = v11
                    }
                    if v4 {
                        fmt.Println(v6, v5)
                    } else {
                        var v12 uint64 = 42
                        _ = v12
                    }
                    var v13 int8 = 100
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 6; fi14++ {
                        v4 = false
                        v15 := v1[(fi14 - (-(fi7)))]
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi14
                        _ = v15
                    }
                    if v4 {
                        var v16 uint16 = 748
                        _ = v16
                    } else {
                        switch (-(fi7)) {
                        case 274:
                            switch ("a" + "") {
                            case "\n\t":
                                v17 := GenS1{F0: v5}
                                _ = v17
                            default:
                                if (48 <= 127) {
                                    if (!false) {
                                        v18 := GenS1{F0: (^(v5))}
                                        _ = v18
                                    }
                                    var v19 uint16 = 473
                                    _ = v19
                                } else {
                                    for fi20 := 0; fi20 < 6; fi20++ {
                                        v21 := GenS1{F0: 0}
                                        _ = fi20
                                        _ = v21
                                    }
                                }
                            }
                        case 1:
                            if false {
                                var v22 int = (fi7 & fi7)
                                switch v22 {
                                case -1:
                                    if v4 {
                                        fmt.Println(fi7, v5, v22)
                                    } else {
                                        if true {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        } else {
                                            fmt.Println(v5, v4)
                                        }
                                    }
                                }
                                fmt.Println(v3)
                                _ = v22
                            }
                        case 679:
                            v23 := GenS1{F0: 2147483647}
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 14; fi24++ {
                            fmt.Println(v6)
                            if (!v4) {
                                v1[(fi7 % ((fi7) | 1))] = v5
                                if (!false) {
                                    fmt.Println(fi24, v4, v6)
                                    v3 = v3
                                    v6 -= (v6 & (v6 + v6))
                                }
                            } else {
                                v25 := GenS2{F0: int32(v6)}
                                switch -2147483648 {
                                case 128:
                                    var v26 float32 = (3.14159 * (((3.14159 / 1.5) - (((3.14159 - 0.0) - 1e10) - 1.5)) / (0.0 + (0.1 + 1e10))))
                                    _ = v26
                                case -850:
                                    if len(v2) > 0 {
                                        v2[0] = (v5 - v5)
                                    }
                                case 127:
                                    if (1e10 <= (0.0 / (3.14159 - 0.1))) {
                                        v27 := GenS3{F0: (^(fi24)), F1: 255}
                                        var v28 uint = 255
                                        fmt.Println(v5)
                                        _ = v27
                                        _ = v28
                                    } else {
                                        var v29 uint8 = 0
                                        _ = v29
                                    }
                                default:
                                    fmt.Println(fi7, v6)
                                }
                                v3 += 471
                                _ = v25
                            }
                            v30 := GenS3{F0: (fi24 * fi24), F1: 2}
                            if len(v2) > 0 {
                                v2[0] = 2147483648
                            }
                            v5 += 42
                            var v31 string = "a"
                            v32 := GenS1{F0: v5}
                            _ = fi24
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        if false {
                            switch "" {
                            case "foo bar":
                                var v33 float64 = 0.0
                                _ = v33
                            case "":
                                var v34 uint = 65535
                                _ = v34
                            case "a":
                                var v35 int8 = -1
                                _ = v35
                            }
                            fmt.Println(v3)
                            if (255 <= 1) {
                                v36 := GenS1{F0: v5}
                                _ = v36
                            } else {
                                var v37 uint64 = v3
                                v2 = append(v2, v5)
                                _ = v37
                            }
                        } else {
                            v38 := GenS4{F0: -1, F1: v3, F2: ((fi7 * 1) > (fi7 / ((fi7) | 1)))}
                            for fi39 := 0; fi39 < 18; fi39++ {
                                v6 = (-(v6))
                                if ((4294967295 == 2) || (-192 <= -1)) {
                                    if false {
                                        var v40 int32 = -408
                                        if v4 {
                                            v5 = v5
                                            v40 -= -2147483648
                                            v40 = 883
                                            if false {
                                                v41 := GenS3{F0: fi7, F1: 65535}
                                                _ = v41
                                            } else {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v4 = ("" <= ("\n\t" + "hello"))
                                                case "a":
                                                    if ((^(v5)) == uint32(v3)) {
                                                        var v42 bool = v4
                                                        _ = v42
                                                    } else {
                                                        if v4 {
                                                            var v43 float32 = (1.5 - 1e10)
                                                            var v44 bool = true
                                                            v4 = v4
                                                            _ = v43
                                                            _ = v44
                                                        } else {
                                                            switch int(v3) {
                                                            case -1:
                                                                v2 = append(v2, v5)
                                                            case 32767:
                                                                var v45 int16 = 283
                                                                switch ("a" + ("hello" + ((("" + "\n\t") + "") + ("a" + ("a" + "foo bar"))))) {
                                                                case "a":
                                                                    if (v4 && (-85 <= -77)) {
                                                                        if v4 {
                                                                            var v46 int32 = (v40 | 184)
                                                                            v46 = (v40 + (v46 * v40))
                                                                            _ = v46
                                                                        } else {
                                                                            v6 = v6
                                                                            if ((!true) && (true && v4)) {
                                                                                if ((1.5 == 1e10) && v4) {
                                                                                    var v47 float64 = 1.5
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        v45 = (v45 ^ -30)
                                                                                    case "a":
                                                                                        v4 = false
                                                                                    }
                                                                                    _ = v47
                                                                                } else {
                                                                                    v48 := GenS3{F0: 9223372036854775807, F1: 82}
                                                                                    _ = v48
                                                                                }
                                                                                var v49 int32 = (^(v40))
                                                                                _ = v49
                                                                            }
                                                                            var v50 int32 = (v40 * v40)
                                                                            _ = v50
                                                                        }
                                                                    } else {
                                                                        var v51 int32 = (v40 ^ (v40 - (-(v40))))
                                                                        v45 += (v45 << ((504) & 63))
                                                                        v52 := GenS2{F0: (v51 ^ v51)}
                                                                        if (v4 && v4) {
                                                                            var v53 float64 = (((float64(float32((0.1 - 1.5))) / -2.25) + (-2.25 * (1.5 * 0.0))) / (1e10 + 1e10))
                                                                            fmt.Println(v53, v3)
                                                                            _ = v53
                                                                        } else {
                                                                            v3 = 18446744073709551615
                                                                            var v54 int64 = (v6 / ((9223372036854775807) | 1))
                                                                            var v55 int16 = v45
                                                                            _ = v54
                                                                            _ = v55
                                                                        }
                                                                        _ = v51
                                                                        _ = v52
                                                                    }
                                                                case "\n\t":
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                case "":
                                                                    var v56 int32 = int32(fi39)
                                                                    _ = v56
                                                                default:
                                                                    v4 = false
                                                                }
                                                                v57 := v1[fi7]
                                                                _ = v45
                                                                _ = v57
                                                            case 127:
                                                                v58 := GenS4{F0: fi39, F1: (v3 - (v3 / ((v3) | 1))), F2: true}
                                                                _ = v58
                                                            }
                                                            var v59 uint8 = 127
                                                            v5 -= v5
                                                            _ = v59
                                                        }
                                                        switch 358 {
                                                        case -2147483648:
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                fmt.Println(v40)
                                                            case "":
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    var v60 string = (("hello" + "\n\t") + ("\n\t" + "foo bar"))
                                                                    _ = v60
                                                                case "\n\t":
                                                                    if false {
                                                                        switch (^(fi39)) {
                                                                        case 9223372036854775807:
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        case 1:
                                                                            switch -874 {
                                                                            case 127:
                                                                                v3 = v3
                                                                            }
                                                                        }
                                                                        v61 := GenS5{F0: v4, F1: 42}
                                                                        v40 -= v40
                                                                        fmt.Println(v4, v6)
                                                                        fmt.Println(fi39, v5, fi39)
                                                                        v4 = true
                                                                        _ = v61
                                                                    } else {
                                                                        var v62 uint = 964
                                                                        v6 += -1
                                                                        v6 = v6
                                                                        _ = v62
                                                                    }
                                                                case "a":
                                                                    var v63 uint = 42
                                                                    _ = v63
                                                                default:
                                                                    switch fi39 {
                                                                    case 32767:
                                                                        v4 = (0.1 == 1e10)
                                                                    case 952:
                                                                        delete(v1, fi7)
                                                                        if (v4 || ((v40 / ((v40) | 1)) >= v40)) {
                                                                            v40 += (v40 << ((42) & 63))
                                                                        }
                                                                        fmt.Println(v3)
                                                                    case -1:
                                                                        var v64 string = ((("" + "hello") + (((("hello" + "") + "\n\t") + ("\n\t" + "hello")) + "")) + "")
                                                                        _ = v64
                                                                    }
                                                                }
                                                            }
                                                        case 9223372036854775807:
                                                            v6 += (v6 & int64(fi7))
                                                        default:
                                                            fmt.Println(v40, fi7)
                                                        }
                                                    }
                                                case "foo bar":
                                                    var v65 uint64 = v3
                                                    _ = v65
                                                default:
                                                    if len(v2) > 0 {
                                                        v2[0] = v5
                                                    }
                                                }
                                                switch (fi7 | (fi39 << ((42) & 63))) {
                                                case 30:
                                                    if false {
                                                        var v66 uint32 = (^(v5))
                                                        _ = v66
                                                    } else {
                                                        fmt.Println(v6, v40, v3)
                                                        if (float32(1e10) != ((0.0 * -2.25) - (1e10 - (1.5 / float32(3.14159))))) {
                                                            var v67 bool = v4
                                                            if true {
                                                                if v4 {
                                                                    v6 = (v6 + (v6 * (v6 | (-(v6)))))
                                                                    v68 := v1[(^(fi39))]
                                                                    _ = v68
                                                                } else {
                                                                    var v69 int = fi7
                                                                    _ = v69
                                                                }
                                                                var v70 uint32 = (v5 / ((2147483647) | 1))
                                                                _ = v70
                                                            }
                                                            _ = len(v2)
                                                            _ = v67
                                                        }
                                                        fmt.Println(v5, v40, v3)
                                                    }
                                                }
                                            }
                                        } else {
                                            v1[fi7] = v5
                                            v3 = (^(v3))
                                            fmt.Println(v5, fi39)
                                        }
                                        _ = v40
                                    }
                                    if v4 {
                                        var v71 bool = v4
                                        v6 += v6
                                        fmt.Println(fi7, v4)
                                        _ = v71
                                    }
                                }
                                _ = fi39
                            }
                            for fi72 := 0; fi72 < 14; fi72++ {
                                _ = len(v1)
                                _ = fi72
                            }
                            _ = v38
                        }
                    }
                }
            }
        }
        _ = fi7
    }
    var v73 int64 = (v6 - v6)
    var v74 int = -694
    for fi75 := 0; fi75 < 7; fi75++ {
        v74 = (v74 << ((289) & 63))
        for fi76 := 0; fi76 < 11; fi76++ {
            switch v74 {
            case -1:
                fmt.Println(v3, v5)
            case 127:
                var v77 int64 = v6
                _ = v77
            }
            v5 += v5
            fmt.Println(v4, fi76, v3)
            _ = fi76
        }
        _ = fi75
    }
    for fi78 := 0; fi78 < 17; fi78++ {
        v79 := v1[int(v5)]
        v74 -= fi78
        if v4 {
            var v80 uint8 = 0
            v81 := GenS6{F0: 42, F1: 127}
            v4 = v4
            _ = v80
            _ = v81
        } else {
            v82 := GenS6{F0: 0, F1: -1}
            for fi83 := 0; fi83 < 12; fi83++ {
                if (!false) {
                    var v84 uint8 = 7
                    var v85 int16 = 42
                    _ = v84
                    _ = v85
                }
                v86 := GenS7{F0: (v74 | -1), F1: -1, F2: v3}
                var v87 int8 = 1
                var v88 int64 = (^(v73))
                _ = fi83
                _ = v86
                _ = v87
                _ = v88
            }
            _ = v82
        }
        _ = fi78
        _ = v79
    }
    if (!v4) {
        switch "a" {
        case "a":
            for fi89 := 0; fi89 < 5; fi89++ {
                var v90 int = fi89
                _ = fi89
                _ = v90
            }
        case "":
            var v91 int8 = -86
            _ = v91
        case "\n\t":
            v6 -= (v73 << ((726) & 63))
        }
        v92 := GenS5{F0: v4, F1: -9223372036854775808}
        v6 -= (v73 - 9223372036854775807)
        _ = v92
    }
    var v93 uint8 = 216
    v5 = uint32(v3)
    if (!(true && (v74 > v74))) {
        var v94 int16 = -537
        if true {
            var v95 float32 = 1e10
            _ = v95
        }
        v74 = v74
        _ = v94
    }
    var v96 int64 = 355
    v97 := GenS5{F0: false, F1: 42}
    switch 32767 {
    case 42:
        v98 := GenS4{F0: v74, F1: (v3 & 0), F2: (3.14159 == (0.1 / 1.5))}
        _ = v98
    case -2147483648:
        for fi99 := 0; fi99 < 18; fi99++ {
            v100 := GenS1{F0: (v5 ^ (v5 >> ((383) & 63)))}
            switch "hello" {
            case "\n\t":
                switch fi99 {
                case 806:
                    if v4 {
                        fmt.Println(v4, v4)
                        v101 := GenS3{F0: (fi99 >> ((4294967295) & 63)), F1: 14}
                        _ = v101
                    }
                default:
                    var v102 uint64 = (v3 & 56)
                    _ = v102
                }
            default:
                var v103 float32 = 1.5
                _ = v103
            }
            _ = fi99
            _ = v100
        }
    case -16:
        fmt.Println(v4, v73)
    default:
        v74 = v74
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v93:", v93)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
    _ = v93
    _ = v96
    _ = v97
}
