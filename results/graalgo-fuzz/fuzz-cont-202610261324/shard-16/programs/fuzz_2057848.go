package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 string; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 uint; F2 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint8; F1 int; F2 int16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint; F1 uint8 }
type GenS11 struct { F0 uint8; F1 uint16 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int64; F1 int32 }
type GenS14 struct { F0 string; F1 uint64 }
type GenS15 struct { F0 uint64; F1 uint32 }
type GenS16 struct { F0 int16; F1 int8 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]string)
    var v3 float64 = (1.5 / 3.14159)
    var v4 uint32 = 831
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]bool)
    var v7 uint16 = 65535
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(v7, v7)
        var v9 uint8 = 43
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 string = "\n\t"
            v6[v11] = true
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    v3 -= (-2.25 * v3)
    if (false || (0 == 255)) {
        v3 += 1e10
        v12 := GenS1{F0: v4, F1: 2147483647, F2: -1}
        _ = v12
    }
    v13 := GenS2{F0: false}
    if (127 <= 127) {
        _ = len(v2)
        v14 := GenS1{F0: (v4 + v4), F1: 1, F2: 0}
        _ = v14
    }
    v3 = v3
    fmt.Println(v7)
    for fi15 := 0; fi15 < 15; fi15++ {
        if (false && false) {
            if true {
                v1 = (v1 - 18446744073709551615)
            } else {
                var v16 float32 = -2.25
                fmt.Println(v4, v1)
                _ = v16
            }
            v7 = (^(v7))
            fmt.Println(v7, v4)
        } else {
            if false {
                for fi17 := 0; fi17 < 3; fi17++ {
                    v18 := []int{fi17, (^(fi15)), 224}
                    v3 = 0.1
                    _ = fi17
                    _ = v18
                }
                var v19 uint = v1
                _ = v19
            } else {
                switch ("a" + (("foo bar" + "hello") + "")) {
                case "":
                    switch fi15 {
                    case -15:
                        var v20 int8 = 18
                        _ = v20
                    case 32767:
                        var v21 int8 = 42
                        _ = v21
                    default:
                        if ((v1 >> (((v1 ^ v1)) & 63)) == v1) {
                            var v22 uint16 = 415
                            fmt.Println(v3, v4, v1)
                            _ = v22
                        }
                    }
                case "a":
                    for fi23 := 0; fi23 < 6; fi23++ {
                        var v24 int16 = -1
                        var v25 int32 = 191
                        if (int(v7) > fi23) {
                            v26 := GenS2{F0: false}
                            switch ("foo bar" + "foo bar") {
                            case "a":
                                if (!true) {
                                    v27 := GenS3{F0: 42, F1: "\n\t", F2: int16(fi23)}
                                    v28 := GenS4{F0: 461, F1: 1, F2: 62}
                                    var v29 uint32 = v4
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                            case "foo bar":
                                fmt.Println(v1, fi15)
                                var v30 uint = (v1 | (v1 / (((v1 >> ((18446744073709551615) & 63))) | 1)))
                                _ = v30
                            case "\n\t":
                                var v31 bool = (!(false && (true && false)))
                                _ = v31
                            }
                            v4 = v4
                            _ = v26
                        }
                        v25 = v25
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                case "foo bar":
                    var v32 int8 = 1
                    _ = v32
                default:
                    v7 = 909
                }
                var v33 int = fi15
                _ = v33
            }
            v34 := GenS5{F0: (v1 + 1)}
            v1 += v1
            _ = v34
        }
        for fi35 := 0; fi35 < 4; fi35++ {
            switch "\n\t" {
            case "\n\t":
                v6["\n\t"] = false
            case "":
                v36 := GenS5{F0: (v1 | 0)}
                switch -493 {
                case 771:
                    switch ("a" + "hello") {
                    case "\n\t":
                        v37 := GenS6{F0: 127, F1: v1, F2: ""}
                        _ = v37
                    default:
                        v1 = (v1 - (v1 | (v1 >> (((v1 / ((1) | 1))) & 63))))
                    }
                case 4294967295:
                    fmt.Println(v1, fi15)
                case 32767:
                    fmt.Println(v4, v3, v4)
                default:
                    if false {
                        v7 = (v7 << (((^(v1))) & 63))
                        v38 := []int{}
                        v4 -= 2147483647
                        _ = v38
                    }
                }
                _ = v36
            default:
                if false {
                    if false {
                        fmt.Println(v4, v7)
                        if true {
                            v3 -= float64(3.14159)
                            v3 = float64((0.0 - float32((1.5 * 1.5))))
                        }
                        v1 = (v1 & (v1 << ((v1) & 63)))
                    }
                }
                switch "" {
                case "":
                    v39 := GenS7{F0: 1}
                    _ = v39
                case "hello":
                    var v40 string = "foo bar"
                    _ = v40
                case "foo bar":
                    var v41 uint8 = 1
                    _ = v41
                }
                var v42 float32 = (0.0 / -2.25)
                v43 := GenS8{F0: 254, F1: 128, F2: 1}
                var v44 float64 = v3
                _ = v42
                _ = v43
                _ = v44
            }
            var v45 bool = ((1.5 + 1e10) >= -2.25)
            _ = fi35
            _ = v45
        }
        _ = fi15
    }
    fmt.Println(v4, v3)
    if false {
        v46 := GenS9{F0: v4}
        switch "" {
        case "hello":
            v1 -= 18446744073709551615
        case "\n\t":
            delete(v6, ("hello" + "hello"))
        case "foo bar":
            if (v4 > v4) {
                var v47 float32 = (0.0 + (((1.5 + ((0.1 + (-2.25 / 1.0)) / (0.1 * (1e10 / (1.5 * 1e10))))) - float32(v3)) / 1e10))
                for fi48 := 0; fi48 < 2; fi48++ {
                    switch ("" + "") {
                    case "":
                        switch (fi48 ^ fi48) {
                        case 0:
                            var v49 int32 = 42
                            _ = v49
                        case 127:
                            for fi50 := 0; fi50 < 18; fi50++ {
                                v3 += 1.5
                                _ = fi50
                            }
                        case -1:
                            _ = len(v5)
                        default:
                            switch -737 {
                            case -971:
                                v7 -= v7
                            case 42:
                                fmt.Println(v1, v1, v47)
                            case 9223372036854775807:
                                v7 = (v7 % (((v7 % (((^(v7))) | 1))) | 1))
                            }
                        }
                    default:
                        var v51 int32 = 88
                        _ = v51
                    }
                    switch (fi48 / ((-500) | 1)) {
                    case 9223372036854775807:
                        var v52 int16 = 1
                        _ = v52
                    default:
                        fmt.Println(fi48)
                    }
                    _ = fi48
                }
                _ = v47
            } else {
                var v53 uint64 = 2
                _ = v53
            }
        default:
            switch "a" {
            case "a":
                for fi54 := 0; fi54 < 3; fi54++ {
                    switch (fi54 / ((32768) | 1)) {
                    case 32768:
                        var v55 uint8 = 127
                        _ = v55
                    case 4294967295:
                        v56 := []int{(fi54 + fi54), 936}
                        _ = v56
                    default:
                        v57 := []int{fi54, int(v1)}
                        _ = v57
                    }
                    v58 := []int{(-(fi54)), fi54, (fi54 * 127)}
                    if false {
                        var v59 uint16 = 0
                        _ = v59
                    } else {
                        v7 = (v7 >> (((v1 & v1)) & 63))
                        if (!(false && true)) {
                            v3 -= (float64(float32(float64(((1e10 - (-2.25 / -2.25)) / float32(v3))))) - ((3.14159 * v3) / 0.1))
                        }
                    }
                    _ = fi54
                    _ = v58
                }
            case "\n\t":
                for fi60 := 0; fi60 < 14; fi60++ {
                    v61 := GenS4{F0: 607, F1: 669, F2: 240}
                    v3 = 1.5
                    for fi62 := 0; fi62 < 13; fi62++ {
                        var v63 int16 = 488
                        _ = fi62
                        _ = v63
                    }
                    _ = fi60
                    _ = v61
                }
            case "hello":
                v64 := GenS10{F0: 0, F1: 0}
                _ = v64
            }
        }
        if ((^(v4)) == v4) {
            if false {
                var v65 string = "foo bar"
                for fi66 := 0; fi66 < 5; fi66++ {
                    v67 := GenS2{F0: true}
                    for fi68 := 0; fi68 < 14; fi68++ {
                        if ((-75 != 10) || true) {
                            v69 := GenS5{F0: v1}
                            if (true && false) {
                                var v70 int16 = -1
                                if (498 >= v7) {
                                    v71 := GenS2{F0: (!(true && (-765 > 119)))}
                                    _ = v71
                                }
                                v7 -= 32767
                                _ = v70
                            } else {
                                var v72 string = ((v65 + v65) + (v65 + v65))
                                _ = v72
                            }
                            v73 := []int{(fi66 * (^(fi68)))}
                            _ = v69
                            _ = v73
                        } else {
                            v74 := GenS11{F0: 0, F1: v7}
                            v75 := GenS12{F0: 2}
                            _ = v74
                            _ = v75
                        }
                        if ((v3 - float64((3.14159 / 1.5))) != v3) {
                            var v76 float64 = (v3 * -2.25)
                            if (true || (364 < 1)) {
                                if true {
                                    var v77 uint16 = (v7 | (^(v7)))
                                    var v78 uint32 = v4
                                    _ = v77
                                    _ = v78
                                }
                                var v79 uint32 = (^(v4))
                                if true {
                                    var v80 uint8 = 128
                                    switch fi68 {
                                    case 2147483647:
                                        v81 := GenS2{F0: true}
                                        _ = v81
                                    case 809:
                                        v80 = v80
                                    case 32767:
                                        var v82 uint8 = 228
                                        _ = v82
                                    default:
                                        fmt.Println(v79)
                                    }
                                    _ = v80
                                } else {
                                    fmt.Println(v76, fi66, v4)
                                    v7 += v7
                                }
                                _ = v79
                            } else {
                                fmt.Println(v1)
                            }
                            _ = v76
                        }
                        _ = len(v6)
                        _ = fi68
                    }
                    switch (((v65 + v65) + v65) + ((v65 + ((v65 + v65) + "")) + (v65 + ""))) {
                    case "\n\t":
                        v65 = v65
                    }
                    _ = fi66
                    _ = v67
                }
                if true {
                    if (127 >= 42) {
                        v83 := v2[-9223372036854775808]
                        _ = v83
                    }
                    v3 -= (1e10 + ((v3 + (0.0 / ((((float64(1.5) - (3.14159 * v3)) + 0.1) - v3) - v3))) - v3))
                } else {
                    if ((!false) || true) {
                        v4 = v4
                        v84 := GenS13{F0: 42, F1: 1}
                        _ = v84
                    }
                }
                _ = v65
            }
        } else {
            fmt.Println(v4)
            for fi85 := 0; fi85 < 17; fi85++ {
                var v86 uint64 = 4294967295
                v5[v4] = (v4 % ((v4) | 1))
                v87 := GenS4{F0: (v86 % ((18446744073709551615) | 1)), F1: uint16(v1), F2: 2}
                _ = fi85
                _ = v86
                _ = v87
            }
        }
        _ = v46
    } else {
        for fi88 := 0; fi88 < 6; fi88++ {
            for fi89 := 0; fi89 < 5; fi89++ {
                v90 := GenS10{F0: v1, F1: 255}
                if (!false) {
                    v7 = (v7 & (^(v7)))
                    v91 := GenS9{F0: (^(v4))}
                    fmt.Println(fi89, v4, fi88)
                    _ = v91
                } else {
                    var v92 uint8 = 142
                    _ = v92
                }
                fmt.Println(v1, fi89)
                _ = fi89
                _ = v90
            }
            if true {
                v7 -= (v7 << (((v1 >> ((v1) & 63))) & 63))
                v93 := GenS14{F0: "a", F1: 262}
                fmt.Println(v3, v1, fi88)
                _ = v93
            } else {
                switch ("a" + ("\n\t" + "foo bar")) {
                case "a":
                    if true {
                        v4 += v4
                        var v94 float32 = (((0.1 - 3.14159) - ((float32(((v3 / v3) / 3.14159)) / 0.1) + 3.14159)) - float32(v3))
                        var v95 int8 = 127
                        _ = v94
                        _ = v95
                    }
                case "foo bar":
                    v96 := GenS15{F0: 759, F1: (v4 / ((1) | 1))}
                    _ = v96
                case "\n\t":
                    for fi97 := 0; fi97 < 5; fi97++ {
                        v3 -= (float64((1.5 - 3.14159)) / v3)
                        var v98 uint64 = 4294967295
                        _ = fi97
                        _ = v98
                    }
                default:
                    v99 := GenS16{F0: 1, F1: -128}
                    _ = v99
                }
            }
            _ = fi88
        }
        switch "" {
        case "hello":
            v100 := GenS10{F0: v1, F1: 128}
            _ = v100
        default:
            var v101 int32 = 42
            _ = v101
        }
        v102 := v2[127]
        _ = v102
    }
    v1 = (^(v1))
    var v103 bool = false
    var v104 bool = true
    var v105 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v103
    _ = v104
    _ = v105
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
