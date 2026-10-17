package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string; F1 int }
type GenS9 struct { F0 uint32; F1 string; F2 string }
type GenS10 struct { F0 bool; F1 int16; F2 int64 }
type GenS11 struct { F0 string; F1 int32 }
type GenS12 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint, p1 bool, p2 int64) int8 {
    v7 := []int{9223372036854775807}
    switch "\n\t" {
    case "\n\t":
        p2 += -141
    case "foo bar":
        var v8 int16 = -1
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -1
}
func genFn1() string {
    if false {
        var v9 uint16 = 0
        _ = v9
    } else {
        var v10 string = "foo bar"
        v11 := GenS1{F0: 65535, F1: -39, F2: 28}
        _ = v10
        _ = v11
    }
    var v12 uint64 = 0
    for fi13 := 0; fi13 < 4; fi13++ {
        if ((!false) && true) {
            if false {
                v12 -= v12
            }
            var v14 string = "\n\t"
            if true {
                v15 := []int{fi13, 1}
                var v16 string = (v14 + ((v14 + v14) + "hello"))
                v17 := GenS2{F0: genFn0(255, (0.1 <= (0.0 + float32(3.14159))), 42)}
                _ = v15
                _ = v16
                _ = v17
            }
            _ = v14
        }
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 uint32 = 46
            v19 = (v19 - (v19 / ((v19) | 1)))
            if false {
                if ((!false) || true) {
                    switch "a" {
                    case "\n\t":
                        var v20 bool = ((1 != -1) && (true && true))
                        _ = v20
                    case "a":
                        v21 := GenS3{F0: ("hello" + ("\n\t" + "\n\t")), F1: 42}
                        _ = v21
                    default:
                        v22 := []int{}
                        _ = v22
                    }
                } else {
                    var v23 float32 = 0.0
                    _ = v23
                }
            } else {
                v12 = 1
                if (true && (!(!(143 != 42)))) {
                    fmt.Println(v19, fi18)
                } else {
                    switch fi13 {
                    case 32767:
                        fmt.Println(v12, v12, v12)
                        v24 := GenS4{F0: 42}
                        var v25 int = fi13
                        _ = v24
                        _ = v25
                    case 0:
                        v26 := GenS4{F0: 32767}
                        _ = v26
                    }
                    v27 := GenS3{F0: "foo bar", F1: 122}
                    _ = v27
                }
            }
            _ = fi18
            _ = v19
        }
        _ = fi13
    }
    _ = v12
    return "hello"
}

func main() {
    var v1 uint8 = 40
    var v2 uint = uint(v1)
    var v3 uint8 = (v1 | (v1 - 255))
    v4 := make(map[int]int)
    var v5 float64 = 1e10
    v6 := []uint32{163, 4294967295, 2147483647}
    switch -302 {
    case 732:
        var v28 int32 = 0
        _ = v28
    case -2147483648:
        v3 += 2
    case 32767:
        if (-1 < 0) {
            if ((4294967295 > 69) && ((true || true) && true)) {
                if (float32(v5) > 3.14159) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                } else {
                    switch 357 {
                    case 977:
                        for fi29 := 0; fi29 < 4; fi29++ {
                            fmt.Println(v1, v3)
                            var v30 uint32 = 2
                            _ = fi29
                            _ = v30
                        }
                    case 1:
                        var v31 int = 32768
                        _ = v31
                    case 2147483647:
                        for fi32 := 0; fi32 < 13; fi32++ {
                            switch (("foo bar" + "\n\t") + ("a" + "\n\t")) {
                            case "foo bar":
                                if (!false) {
                                    var v33 int32 = -1
                                    switch "a" {
                                    case "foo bar":
                                        v3 -= (v1 ^ 128)
                                    default:
                                        if true {
                                            var v34 uint32 = 2147483648
                                            for fi35 := 0; fi35 < 5; fi35++ {
                                                var v36 float32 = (3.14159 - 0.1)
                                                if (!(1 <= v33)) {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v37 := v4[fi32]
                                                        _ = v37
                                                    default:
                                                        var v38 string = "a"
                                                        _ = v38
                                                    }
                                                } else {
                                                    if true {
                                                        var v39 float32 = v36
                                                        if (true || false) {
                                                            v36 += 0.1
                                                            var v40 uint16 = 65535
                                                            v6 = append(v6, 65535)
                                                            _ = v40
                                                        }
                                                        _ = len(v4)
                                                        _ = v39
                                                    }
                                                    v3 = (v1 % ((v1) | 1))
                                                }
                                                _ = len(v6)
                                                v41 := GenS3{F0: "hello", F1: 120}
                                                v2 += (v2 + v2)
                                                switch "foo bar" {
                                                case "a":
                                                    var v42 float64 = v5
                                                    _ = v42
                                                }
                                                switch "hello" {
                                                case "foo bar":
                                                    v34 -= v34
                                                case "\n\t":
                                                    switch ("foo bar" + "\n\t") {
                                                    case "hello":
                                                        if (!(true || true)) {
                                                            var v43 uint16 = 0
                                                            _ = v43
                                                        }
                                                    case "foo bar":
                                                        if true {
                                                            switch ("" + "hello") {
                                                            case "a":
                                                                v2 -= v2
                                                            default:
                                                                v44 := GenS3{F0: "a", F1: genFn0((v2 & 42), true, -429)}
                                                                _ = v44
                                                            }
                                                            _ = len(v4)
                                                        }
                                                    }
                                                }
                                                _ = fi35
                                                _ = v36
                                                _ = v41
                                            }
                                            var v45 float32 = (1e10 + 3.14159)
                                            _ = v34
                                            _ = v45
                                        } else {
                                            var v46 int8 = 23
                                            _ = v46
                                        }
                                    }
                                    _ = v33
                                }
                            case "a":
                                v47 := GenS1{F0: 2, F1: 0, F2: 100}
                                _ = v47
                            case "\n\t":
                                fmt.Println(v3, v5, fi32)
                            default:
                                for fi48 := 0; fi48 < 1; fi48++ {
                                    _ = len(v4)
                                    v49 := GenS5{F0: 0}
                                    _ = fi48
                                    _ = v49
                                }
                            }
                            _ = fi32
                        }
                    default:
                        for fi50 := 0; fi50 < 12; fi50++ {
                            v6 = append(v6, 42)
                            _ = fi50
                        }
                    }
                    var v51 int8 = -128
                    _ = v51
                }
                v2 += 587
                var v52 uint32 = 0
                _ = v52
            } else {
                var v53 bool = true
                fmt.Println(v2)
                var v54 string = "foo bar"
                _ = v53
                _ = v54
            }
        }
    default:
        switch "\n\t" {
        case "":
            v3 += (v3 % (((v3 / (((v3 & 42)) | 1))) | 1))
        case "\n\t":
            for fi55 := 0; fi55 < 9; fi55++ {
                var v56 uint16 = 255
                var v57 uint64 = 285
                _ = fi55
                _ = v56
                _ = v57
            }
        case "hello":
            v6 = append(v6, 65535)
        default:
            for fi58 := 0; fi58 < 14; fi58++ {
                if false {
                    v59 := GenS1{F0: 32767, F1: 100, F2: -1}
                    var v60 int64 = 1
                    var v61 uint8 = 26
                    _ = v59
                    _ = v60
                    _ = v61
                }
                if (!false) {
                    v3 = v3
                    if true {
                        switch fi58 {
                        case 42:
                            v2 += (v2 / (((v2 & (v2 + (v2 - 65535)))) | 1))
                        case 128:
                            switch "\n\t" {
                            case "hello":
                                var v62 uint64 = 42
                                _ = v62
                            case "\n\t":
                                _ = len(v4)
                            }
                        case 32768:
                            v5 = v5
                        }
                        var v63 uint8 = (v3 >> ((v2) & 63))
                        _ = v63
                    } else {
                        v64 := GenS2{F0: 65}
                        v65 := GenS4{F0: 1}
                        _ = v64
                        _ = v65
                    }
                    v66 := GenS6{F0: -311, F1: 128}
                    _ = v66
                }
                _ = fi58
            }
        }
    }
    var v67 uint8 = (v3 ^ 128)
    v68 := GenS7{F0: -1}
    var v69 bool = false
    for fi70 := 0; fi70 < 3; fi70++ {
        v3 -= v1
        fmt.Println(v5)
        for fi71 := 0; fi71 < 3; fi71++ {
            v72 := v4[fi70]
            if v69 {
                var v73 uint64 = 1
                v3 = v3
                _ = v73
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = len(v4)
                v74 := GenS3{F0: "hello", F1: 100}
                _ = len(v6)
                v75 := GenS8{F0: "a", F1: (v72 << ((65535) & 63))}
                _ = v74
                _ = v75
            }
            if len(v6) > 0 {
                v6[0] = 14
            }
            _ = fi71
            _ = v72
        }
        _ = fi70
    }
    switch -41 {
    case -2147483648:
        if false {
            for fi76 := 0; fi76 < 18; fi76++ {
                var v77 string = ("\n\t" + "a")
                v78 := GenS9{F0: 2147483648, F1: "", F2: v77}
                if v69 {
                    for fi79 := 0; fi79 < 7; fi79++ {
                        v80 := v4[32768]
                        v81 := GenS4{F0: 0}
                        _ = fi79
                        _ = v80
                        _ = v81
                    }
                }
                _ = fi76
                _ = v77
                _ = v78
            }
        }
    }
    _ = len(v6)
    var v82 bool = false
    if v82 {
        var v83 uint32 = 2147483647
        var v84 int16 = -32768
        switch -580 {
        case 128:
            if v82 {
                v84 += (-(v84))
                if (-1 == 0) {
                    fmt.Println(v69, v84, v84)
                    v82 = v69
                    var v85 uint32 = v83
                    _ = v85
                } else {
                    var v86 uint16 = 65535
                    var v87 uint8 = v67
                    for fi88 := 0; fi88 < 2; fi88++ {
                        for fi89 := 0; fi89 < 4; fi89++ {
                            var v90 uint32 = v83
                            _ = fi89
                            _ = v90
                        }
                        var v91 uint16 = (v86 >> (((v2 + (v2 - v2))) & 63))
                        _ = fi88
                        _ = v91
                    }
                    _ = v86
                    _ = v87
                }
                if len(v6) > 0 {
                    v6[0] = (^(v83))
                }
            } else {
                switch ("foo bar" + "a") {
                case "a":
                    fmt.Println(v5, v2)
                }
            }
        case 116:
            fmt.Println(v67, v84)
        }
        _ = v83
        _ = v84
    } else {
        for fi92 := 0; fi92 < 1; fi92++ {
            var v93 uint8 = (^(v1))
            v1 = (^(v1))
            _ = len(v6)
            _ = fi92
            _ = v93
        }
        if v69 {
            v94 := GenS1{F0: 42, F1: -1, F2: genFn0(v2, (-795 < int(v2)), 587)}
            if (784 >= -1) {
                fmt.Println(v3)
                if v82 {
                    v82 = false
                    v5 = 0.0
                    v95 := GenS10{F0: v69, F1: -78, F2: 578}
                    _ = v95
                }
                for fi96 := 0; fi96 < 20; fi96++ {
                    if ((v69 || v69) || (v5 > (v5 * 1e10))) {
                        v97 := GenS11{F0: ("\n\t" + "a"), F1: 738}
                        v3 = 255
                        v69 = false
                        _ = v97
                    }
                    _ = fi96
                }
            } else {
                v98 := GenS10{F0: false, F1: -596, F2: 42}
                _ = v98
            }
            _ = v94
        } else {
            for fi99 := 0; fi99 < 2; fi99++ {
                var v100 float32 = float32(v5)
                v2 = v2
                _ = fi99
                _ = v100
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        for fi101 := 0; fi101 < 20; fi101++ {
            if v82 {
                v4[-174] = (fi101 >> ((v2) & 63))
                v102 := v4[fi101]
                v103 := GenS12{F0: v82, F1: 255}
                _ = v102
                _ = v103
            }
            _ = fi101
        }
    }
    _ = len(v6)
    var v104 uint8 = (v1 + (v67 & v67))
    for fi105 := 0; fi105 < 16; fi105++ {
        v82 = v69
        if (!v82) {
            v3 = (v1 | (v3 * v3))
        }
        v2 += v2
        _ = fi105
    }
    v5 -= v5
    for fi106 := 0; fi106 < 17; fi106++ {
        var v107 uint8 = (v3 + v67)
        v2 += uint(v3)
        _ = fi106
        _ = v107
    }
    var v108 uint16 = 42
    v3 -= v3
    var v109 int64 = 425
    v69 = (2147483647 >= -302)
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v108:", v108)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v67:", v67)
    fmt.Println("v69:", v69)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v104
    _ = v108
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
    _ = v69
    _ = v82
}
