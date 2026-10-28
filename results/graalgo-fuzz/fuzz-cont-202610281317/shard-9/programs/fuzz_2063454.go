package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 string }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int8; F1 bool; F2 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 int32; F2 uint64 }
type GenS7 struct { F0 int16; F1 bool }
type GenS8 struct { F0 string; F1 string; F2 uint16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint64; F1 int32 }
type GenS11 struct { F0 uint16 }
func genFn0(p0 bool, p1 int32, p2 int) uint64 {
    var v4 int8 = -128
    v5 := GenS1{F0: 294, F1: 645, F2: "hello"}
    v6 := []int{(p2 - p2)}
    var v7 uint8 = 252
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{65535, 0}
    var v3 string = "foo bar"
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: 18446744073709551615, F1: 42, F2: "\n\t"}
        for fi10 := 0; fi10 < 20; fi10++ {
            v3 = v3
            fmt.Println(fi10)
            if true {
                v3 = v3
                switch v3 {
                case "foo bar":
                    var v11 int = fi8
                    _ = v11
                case "hello":
                    var v12 uint8 = 42
                    _ = v12
                }
            }
            if true {
                var v13 float32 = (1e10 - (-2.25 + float32((((3.14159 - 3.14159) / 3.14159) + 0.0))))
                _ = len(v1)
                if false {
                    var v14 float32 = v13
                    if false {
                        fmt.Println(v3, v14, v13)
                    }
                    var v15 float32 = (float32(0.0) - v14)
                    _ = v14
                    _ = v15
                }
                var v16 int32 = 244
                _ = v13
                _ = v16
            } else {
                fmt.Println(fi10)
            }
            v3 = v3
            _ = fi10
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            switch v3 {
            case "":
                switch fi8 {
                case -9223372036854775808:
                    var v18 string = ("hello" + ((("a" + v3) + (v3 + "foo bar")) + v3))
                    _ = v18
                default:
                    fmt.Println(fi8, v3)
                }
            default:
                v3 = v3
            }
            v19 := GenS1{F0: 0, F1: 42, F2: "\n\t"}
            _ = fi17
            _ = v19
        }
        _ = fi8
        _ = v9
    }
    if false {
        switch -2147483648 {
        case 42:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case -2147483648:
            delete(v1, "foo bar")
        }
    } else {
        for fi20 := 0; fi20 < 10; fi20++ {
            v2 = append(v2, 65535)
            v3 = v3
            for fi21 := 0; fi21 < 4; fi21++ {
                switch 627 {
                case 317:
                    v3 = ("hello" + v3)
                case -2147483648:
                    if false {
                        if true {
                            var v22 int16 = 0
                            var v23 int64 = 1
                            v24 := v1["\n\t"]
                            fmt.Println(v24, v23, fi20)
                            var v25 uint64 = 0
                            v24 = (v24 + v3)
                            fmt.Println(v22, v23, v23)
                            var v26 float32 = 1.5
                            v25 -= v25
                            if (!true) {
                                var v27 uint16 = 536
                                _ = v27
                            }
                            switch 127 {
                            case 550:
                                switch ((("a" + v3) + ("a" + v24)) + "\n\t") {
                                case "\n\t":
                                    switch "a" {
                                    case "hello":
                                        fmt.Println(v24, v24, fi20)
                                    case "\n\t":
                                        if false {
                                            var v28 uint16 = 931
                                            v29 := GenS2{F0: v25, F1: v24}
                                            fmt.Println(v22, v24)
                                            _ = v28
                                            _ = v29
                                        }
                                    default:
                                        v30 := GenS3{F0: -128, F1: 195}
                                        _ = v30
                                    }
                                case "hello":
                                    switch fi21 {
                                    case 1:
                                        _ = len(v1)
                                        switch v24 {
                                        case "\n\t":
                                            v26 += -2.25
                                        case "a":
                                            v23 = (v23 % ((v23) | 1))
                                        default:
                                            if ((v23 > (v23 - v23)) || true) {
                                                if true {
                                                    v31 := GenS3{F0: 100, F1: 4294967295}
                                                    var v32 uint = 760
                                                    if len(v2) > 0 {
                                                        v2[0] = 4294967295
                                                    }
                                                    v3 = v3
                                                    v22 -= v22
                                                    _ = v31
                                                    _ = v32
                                                }
                                                _ = len(v2)
                                                var v33 bool = (true || true)
                                                _ = v33
                                            }
                                        }
                                    case -2147483648:
                                        if true {
                                            var v34 uint = 18446744073709551615
                                            if len(v2) > 0 {
                                                v2[0] = 2
                                            }
                                            v23 += (v23 | -831)
                                            _ = v34
                                        }
                                    case -1:
                                        v35 := GenS4{F0: 127, F1: false, F2: -338}
                                        _ = v35
                                    }
                                case "foo bar":
                                    fmt.Println(v3, v23, v3)
                                default:
                                    v36 := GenS1{F0: (^(v25)), F1: (^(v23)), F2: ("" + ("a" + v3))}
                                    _ = v36
                                }
                            case 4294967295:
                                v22 -= (-(v22))
                            }
                            var v37 uint64 = v25
                            v38 := GenS4{F0: 100, F1: false, F2: (fi20 ^ 2147483647)}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v22 -= (v22 ^ v22)
                            _ = v22
                            _ = v23
                            _ = v24
                            _ = v25
                            _ = v26
                            _ = v37
                            _ = v38
                        }
                        if (false && true) {
                            v3 = v3
                        } else {
                            v39 := GenS4{F0: -83, F1: ((false && ((false && false) && true)) && (((v3 + v3) + (v3 + ((v3 + v3) + v3))) <= "a")), F2: fi21}
                            switch v3 {
                            case "a":
                                v3 = "hello"
                                var v40 int16 = 32767
                                var v41 int8 = 127
                                _ = v40
                                _ = v41
                            }
                            v3 = v3
                            _ = v39
                        }
                        v1[v3] = v3
                        var v42 int32 = -221
                        _ = v42
                    }
                default:
                    v3 = v3
                }
                var v43 int16 = 0
                v44 := GenS5{F0: -9223372036854775808}
                _ = fi21
                _ = v43
                _ = v44
            }
            _ = fi20
        }
        if (42 <= 1) {
            v3 = "a"
            _ = len(v2)
            var v45 int32 = 795
            _ = v45
        }
        var v46 uint8 = 128
        _ = v46
    }
    v3 = "\n\t"
    if false {
        if true {
            v3 = "a"
            switch 0 {
            case -9223372036854775808:
                var v47 int = 782
                _ = v47
            case -2147483648:
                for fi48 := 0; fi48 < 5; fi48++ {
                    if ((!(!(!(2 < 1)))) && (1 <= 607)) {
                        var v49 string = v3
                        if (false || (fi48 >= fi48)) {
                            v50 := GenS6{F0: (true && (1 < 725)), F1: -1, F2: 4294967295}
                            if false {
                                var v51 int = -2147483648
                                v51 = fi48
                                _ = v51
                            }
                            _ = v50
                        }
                        v52 := GenS6{F0: (!false), F1: 1, F2: 42}
                        _ = v49
                        _ = v52
                    } else {
                        v3 = ("\n\t" + ("foo bar" + v3))
                        if (211 >= 116) {
                            v3 = "foo bar"
                            var v53 float32 = float32((((1e10 / 1e10) * -2.25) / float64(((float32(3.14159) * -2.25) * float32(float64(0.1))))))
                            switch (fi48 << ((763) & 63)) {
                            case -1:
                                v53 = float32((1e10 + (3.14159 - 1e10)))
                            case 42:
                                if (-159 <= 42) {
                                    v54 := GenS4{F0: 0, F1: true, F2: fi48}
                                    v53 = v53
                                    fmt.Println(v53, fi48)
                                    _ = v54
                                } else {
                                    v55 := GenS3{F0: -128, F1: 0}
                                    v56 := GenS7{F0: int16(fi48), F1: false}
                                    fmt.Println(v3, v53, v3)
                                    _ = v55
                                    _ = v56
                                }
                            case 1:
                                var v57 int8 = int8(fi48)
                                _ = v57
                            }
                            _ = v53
                        }
                        for fi58 := 0; fi58 < 3; fi58++ {
                            var v59 uint16 = 1
                            _ = fi58
                            _ = v59
                        }
                    }
                    if false {
                        for fi60 := 0; fi60 < 4; fi60++ {
                            switch fi60 {
                            case 4294967295:
                                var v61 int16 = -32768
                                _ = v61
                            case 297:
                                v3 = (("hello" + "a") + ((v3 + (v3 + v3)) + (v3 + v3)))
                            default:
                                if false {
                                    fmt.Println(fi60, fi60)
                                } else {
                                    if (!true) {
                                        v62 := GenS3{F0: 1, F1: 45}
                                        _ = v62
                                    }
                                    v3 = v3
                                    var v63 float32 = ((1e10 * ((1.5 * 1.5) - -2.25)) - (1.5 - ((3.14159 * 0.1) * 3.14159)))
                                    _ = v63
                                }
                            }
                            fmt.Println(fi48)
                            _ = fi60
                        }
                    }
                    _ = fi48
                }
            case 9223372036854775807:
                switch 127 {
                case 32767:
                    v3 = "hello"
                default:
                    if (!true) {
                        delete(v1, ("foo bar" + (v3 + v3)))
                    } else {
                        v64 := GenS7{F0: 42, F1: (!true)}
                        if (0 == 65535) {
                            fmt.Println(v3)
                        } else {
                            if (((81 <= 585) || (-560 <= 9223372036854775807)) || true) {
                                v65 := GenS8{F0: ((v3 + v3) + "a"), F1: v3, F2: 255}
                                if false {
                                    fmt.Println(v3)
                                    for fi66 := 0; fi66 < 18; fi66++ {
                                        v67 := GenS9{F0: -1}
                                        _ = fi66
                                        _ = v67
                                    }
                                } else {
                                    switch 439 {
                                    case -2147483648:
                                        v3 = v3
                                    case 641:
                                        fmt.Println(v3)
                                    default:
                                        switch 556 {
                                        case -889:
                                            v3 = v3
                                        case -1:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case 32768:
                                            fmt.Println(v3)
                                        }
                                    }
                                    v1["foo bar"] = v3
                                }
                                _ = v65
                            } else {
                                switch v3 {
                                case "foo bar":
                                    var v68 uint = 1
                                    _ = v68
                                case "\n\t":
                                    fmt.Println(v3)
                                case "a":
                                    var v69 string = v3
                                    _ = v69
                                }
                            }
                            _ = len(v1)
                        }
                        var v70 uint16 = 880
                        _ = v64
                        _ = v70
                    }
                }
            default:
                for fi71 := 0; fi71 < 18; fi71++ {
                    v3 = v3
                    if (!true) {
                        var v72 uint16 = 255
                        v3 = v3
                        _ = v72
                    }
                    if false {
                        if (2147483648 != 2147483647) {
                            fmt.Println(v3, v3)
                            if (0 == 128) {
                                delete(v1, v3)
                            }
                            if false {
                                fmt.Println(fi71, v3)
                                v73 := GenS7{F0: -32768, F1: false}
                                _ = v73
                            }
                        } else {
                            var v74 bool = false
                            v75 := GenS5{F0: 9223372036854775807}
                            fmt.Println(v74, v3)
                            _ = v74
                            _ = v75
                        }
                    } else {
                        delete(v1, v3)
                        v3 = ""
                    }
                    _ = fi71
                }
            }
        } else {
            for fi76 := 0; fi76 < 1; fi76++ {
                _ = len(v2)
                v3 = ((((v3 + v3) + v3) + "hello") + ("a" + "hello"))
                var v77 float32 = 0.1
                _ = fi76
                _ = v77
            }
        }
        v78 := GenS7{F0: -1, F1: (!(!false))}
        _ = v78
    } else {
        if ((1.5 >= 1e10) && (-32768 <= -1)) {
            fmt.Println(v3)
            var v79 int64 = 42
            v79 = (v79 >> ((2) & 63))
            _ = v79
        } else {
            v3 = "\n\t"
        }
        for fi80 := 0; fi80 < 7; fi80++ {
            fmt.Println(fi80, fi80)
            v3 = v3
            switch fi80 {
            case 42:
                for fi81 := 0; fi81 < 12; fi81++ {
                    var v82 uint8 = 2
                    _ = fi81
                    _ = v82
                }
            case 219:
                var v83 int16 = 1000
                _ = v83
            }
            _ = fi80
        }
    }
    v3 = "a"
    var v84 int = 4294967295
    v85 := GenS8{F0: "\n\t", F1: v3, F2: 223}
    if true {
        for fi86 := 0; fi86 < 5; fi86++ {
            var v87 string = v3
            fmt.Println(v87, v87, v84)
            _ = fi86
            _ = v87
        }
    } else {
        if true {
            v88 := GenS9{F0: -1}
            _ = v88
        }
        var v89 string = v3
        v89 = v3
        _ = v89
    }
    if (false && false) {
        v2 = append(v2, 1)
        v3 = (v3 + v3)
    } else {
        delete(v1, "")
        v84 = v84
        for fi90 := 0; fi90 < 7; fi90++ {
            if (190 < 957) {
                if false {
                    if false {
                        if false {
                            var v91 uint16 = 65535
                            _ = v91
                        }
                        var v92 uint16 = 2
                        delete(v1, (v3 + "\n\t"))
                        _ = v92
                    }
                    for fi93 := 0; fi93 < 5; fi93++ {
                        switch fi93 {
                        case 4294967295:
                            v84 -= fi90
                        }
                        fmt.Println(fi93, v3, v84)
                        fmt.Println(fi93, fi93)
                        _ = fi93
                    }
                    v94 := v1[((v3 + v3) + v3)]
                    _ = v94
                } else {
                    if true {
                        switch (^(fi90)) {
                        case -209:
                            if (v3 == (v3 + v3)) {
                                var v95 uint64 = 278
                                _ = v95
                            }
                        }
                    }
                    var v96 uint32 = 2147483648
                    _ = v96
                }
            }
            if true {
                for fi97 := 0; fi97 < 15; fi97++ {
                    if false {
                        v98 := GenS3{F0: -128, F1: 535}
                        _ = v98
                    } else {
                        delete(v1, "a")
                        var v99 uint8 = 127
                        _ = len(v2)
                        _ = v99
                    }
                    if (v84 <= fi97) {
                        switch (fi97 ^ v84) {
                        case 32767:
                            fmt.Println(fi97)
                        case 42:
                            if (-268 <= 0) {
                                if ((false || true) || true) {
                                    var v100 int16 = -121
                                    switch "" {
                                    case "":
                                        if true {
                                            var v101 bool = true
                                            var v102 float64 = (1.5 - 1.5)
                                            _ = len(v2)
                                            _ = v101
                                            _ = v102
                                        } else {
                                            v103 := GenS9{F0: -458}
                                            _ = v103
                                        }
                                    case "hello":
                                        fmt.Println(v84, v3)
                                    default:
                                        var v104 uint16 = 0
                                        _ = v104
                                    }
                                    _ = v100
                                } else {
                                    if (!false) {
                                        v84 -= -584
                                    } else {
                                        var v105 int8 = -22
                                        _ = v105
                                    }
                                }
                                if true {
                                    v84 = v84
                                }
                                var v106 float32 = ((1.5 + ((1.5 / 1e10) - 0.1)) / 0.1)
                                _ = v106
                            }
                        case 2147483647:
                            switch "a" {
                            case "foo bar":
                                var v107 float64 = (0.1 / (0.0 - float64((1.5 / (-2.25 + 1e10)))))
                                _ = v107
                            case "hello":
                                if len(v2) > 0 {
                                    v2[0] = 2147483648
                                }
                            }
                            var v108 bool = true
                            _ = v108
                        default:
                            v1[(v3 + v3)] = v3
                        }
                    } else {
                        if (true && false) {
                            _ = len(v2)
                            v84 -= -1
                        } else {
                            var v109 uint64 = 18446744073709551615
                            var v110 uint32 = 2147483647
                            v3 = v3
                            _ = v109
                            _ = v110
                        }
                    }
                    var v111 int8 = 1
                    _ = fi97
                    _ = v111
                }
            }
            _ = fi90
        }
    }
    if len(v2) > 0 {
        v2[0] = 42
    }
    switch v3 {
    case "foo bar":
        var v112 uint16 = 255
        _ = v112
    case "hello":
        if false {
            var v113 int8 = -128
            for fi114 := 0; fi114 < 1; fi114++ {
                v2 = append(v2, 2147483648)
                _ = fi114
            }
            _ = v113
        } else {
            fmt.Println(v84, v84)
            switch v3 {
            case "hello":
                v115 := GenS10{F0: 172, F1: -938}
                _ = v115
            case "\n\t":
                if true {
                    var v116 uint32 = 0
                    _ = v116
                } else {
                    if false {
                        v2 = append(v2, 652)
                    } else {
                        _ = len(v1)
                    }
                    v3 = (v3 + v3)
                    var v117 int32 = 0
                    _ = v117
                }
            default:
                var v118 int = (-(v84))
                _ = v118
            }
        }
    }
    for fi119 := 0; fi119 < 5; fi119++ {
        var v120 int16 = 1
        switch "hello" {
        case "hello":
            v121 := GenS4{F0: -73, F1: false, F2: (fi119 + fi119)}
            _ = v121
        case "":
            _ = len(v2)
        case "foo bar":
            var v122 uint16 = 903
            _ = v122
        }
        for fi123 := 0; fi123 < 1; fi123++ {
            fmt.Println(v3, v3)
            _ = fi123
        }
        _ = fi119
        _ = v120
    }
    v124 := GenS11{F0: 42}
    var v125 uint16 = 484
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v125:", v125)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v124
    _ = v125
    _ = v2
    _ = v3
    _ = v84
    _ = v85
}
