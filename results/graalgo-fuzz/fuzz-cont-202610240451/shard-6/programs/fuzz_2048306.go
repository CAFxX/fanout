package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint32; F2 int }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int16; F1 uint32 }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 uint64; F1 int32 }
type GenS10 struct { F0 int64; F1 int32 }
type GenS11 struct { F0 int; F1 int32; F2 uint }
type GenS12 struct { F0 int64; F1 uint; F2 uint64 }
type GenS13 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 bool; F1 uint8 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 int32; F1 uint16 }
type GenS18 struct { F0 uint; F1 int32; F2 int8 }
type GenS19 struct { F0 bool; F1 int }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 string; F1 int16 }
type GenS22 struct { F0 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = 1e10
    v3 := make(map[string]uint32)
    var v4 uint16 = 2
    var v5 int = -19
    v6 := GenS1{F0: ((("a" + ("a" + ("foo bar" + ""))) + (((("" + ("\n\t" + "\n\t")) + "\n\t") + "a") + "hello")) + "\n\t")}
    v5 = v5
    if (!false) {
        v5 = 127
        delete(v3, "hello")
    } else {
        var v7 uint16 = 0
        _ = v7
    }
    fmt.Println(v2, v2, v2)
    if false {
        v8 := GenS2{F0: 677, F1: 2147483648, F2: v5}
        _ = v8
    } else {
        switch "\n\t" {
        case "a":
            v4 = 221
        case "hello":
            var v9 int8 = 100
            _ = v9
        case "":
            switch ("hello" + "a") {
            case "":
                for fi10 := 0; fi10 < 9; fi10++ {
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v12 := GenS3{F0: 374, F1: 116}
                        var v13 int32 = 42
                        fmt.Println(fi11)
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    v4 -= 255
                    _ = fi10
                }
            case "foo bar":
                v5 += v5
            case "hello":
                v1[-9223372036854775808] = 1
            }
        default:
            switch "" {
            case "":
                v14 := GenS4{F0: v4, F1: 42}
                _ = v14
            case "hello":
                v5 += (v5 + (-(v5)))
            default:
                v1[(v5 / (((^(v5))) | 1))] = 1
            }
        }
        switch (-(v5)) {
        case 2147483647:
            fmt.Println(v2, v2, v2)
        case -1:
            v4 = (v4 << ((65535) & 63))
        case 298:
            if false {
                fmt.Println(v5, v5, v2)
                var v15 uint8 = 1
                var v16 bool = true
                _ = v15
                _ = v16
            }
        }
        if false {
            fmt.Println(v4, v5)
            var v17 float32 = float32(float64(float32(3.14159)))
            v5 = (v5 ^ v5)
            _ = v17
        } else {
            v18 := GenS3{F0: 1, F1: 920}
            v19 := GenS5{F0: "foo bar"}
            _ = v18
            _ = v19
        }
    }
    if (42 != 125) {
        v20 := []int{(^(v5)), v5}
        var v21 int = v5
        v20 = append(v20, v21)
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        if true {
            var v23 int16 = 0
            _ = v23
        }
        v24 := GenS6{F0: (v5 & fi22)}
        v25 := GenS7{F0: 42, F1: 2147483647}
        _ = fi22
        _ = v24
        _ = v25
    }
    v26 := GenS6{F0: 0}
    var v27 int64 = 141
    v5 += v5
    fmt.Println(v5)
    if (0.0 != 3.14159) {
        if (v27 > (v27 >> ((255) & 63))) {
            v2 = (v2 + (float32(((3.14159 + -2.25) / float64((float32((((float64(float32(-2.25)) - (-2.25 + 0.0)) / 1.5) / float64(v2))) + v2)))) + v2))
            if false {
                if true {
                    var v28 uint16 = (v4 % ((v4) | 1))
                    v29 := []int{(v5 << ((439) & 63)), (^(v5)), v5, (v5 - v5)}
                    v2 = ((float32(-2.25) - (0.0 / (1e10 / float32(-2.25)))) + v2)
                    _ = v28
                    _ = v29
                } else {
                    var v30 string = "a"
                    _ = v30
                }
                fmt.Println(v27, v5, v4)
            }
            for fi31 := 0; fi31 < 19; fi31++ {
                var v32 bool = (0.1 == 0.0)
                fmt.Println(fi31)
                if ("" <= "\n\t") {
                    var v33 uint8 = 1
                    v2 = (3.14159 - (float32(1.5) / (v2 / ((v2 + 0.1) / 1e10))))
                    for fi34 := 0; fi34 < 14; fi34++ {
                        switch "a" {
                        case "\n\t":
                            switch "hello" {
                            case "foo bar":
                                switch "hello" {
                                case "a":
                                    if (!((v4 ^ v4) < (v4 << ((42) & 63)))) {
                                        var v35 float64 = -2.25
                                        fmt.Println(v27)
                                        fmt.Println(fi34)
                                        _ = v35
                                    } else {
                                        var v36 float32 = (v2 + (v2 - (((3.14159 / float32((1.5 / (((0.1 * -2.25) + (1e10 - (0.0 / 3.14159))) + 3.14159)))) - v2) + 1.5)))
                                        _ = v36
                                    }
                                case "hello":
                                    var v37 string = "\n\t"
                                    _ = v37
                                default:
                                    v38 := GenS8{F0: v32, F1: -492}
                                    _ = v38
                                }
                            }
                        }
                        if v32 {
                            v39 := GenS9{F0: 2, F1: -403}
                            fmt.Println(v27, v4, v27)
                            _ = len(v3)
                            _ = v39
                        }
                        if v32 {
                            v40 := []int{(^(fi34))}
                            if v32 {
                                v5 += 576
                                if ((!(!(2147483647 == 2147483647))) && (!v32)) {
                                    if ("\n\t" == "") {
                                        var v41 uint8 = 2
                                        fmt.Println(v33, v5, v4)
                                        v33 += 30
                                        _ = v41
                                    } else {
                                        v5 -= 32767
                                        if len(v40) >= 2 {
                                            v40 = v40[1:len(v40)]
                                        }
                                        switch fi34 {
                                        case -842:
                                            if (0.1 < (((-2.25 * (0.1 - (((((0.1 + -2.25) + ((float64((-2.25 / float32(((0.1 + 0.0) / 0.1)))) / 3.14159) / 0.1)) - float64(v2)) / 1.5) * 1.5))) - 1.5) + 3.14159)) {
                                                v5 = fi31
                                                if (false && (!(-137 != 1000))) {
                                                    var v42 int16 = 32767
                                                    v43 := GenS10{F0: (^(v27)), F1: -2147483648}
                                                    _ = v42
                                                    _ = v43
                                                }
                                                var v44 uint8 = (v33 * (^(v33)))
                                                v32 = v32
                                                fmt.Println(v2, fi31)
                                                v44 += v44
                                                var v45 string = ("foo bar" + "foo bar")
                                                _ = v44
                                                _ = v45
                                            } else {
                                                if len(v40) > 0 {
                                                    v40[0] = (fi34 % (((v5 / ((fi34) | 1))) | 1))
                                                }
                                            }
                                        case 0:
                                            v5 = int(v4)
                                        default:
                                            v46 := v3["a"]
                                            _ = v46
                                        }
                                    }
                                }
                                if v32 {
                                    if len(v40) >= 2 {
                                        v40 = v40[1:len(v40)]
                                    }
                                    v5 += 32768
                                } else {
                                    if v32 {
                                        v47 := GenS11{F0: (^(v5)), F1: -374, F2: 982}
                                        fmt.Println(v2, v5, v33)
                                        _ = v47
                                    }
                                    var v48 uint16 = v4
                                    v1[(v5 << ((257) & 63))] = 2
                                    _ = v48
                                }
                                var v49 bool = (!false)
                                _ = v49
                            }
                            _ = v40
                        }
                        _ = fi34
                    }
                    _ = v33
                }
                _ = fi31
                _ = v32
            }
        } else {
            fmt.Println(v27, v2)
            var v50 uint64 = 538
            if false {
                switch "a" {
                case "\n\t":
                    if ((v5 | v5) <= v5) {
                        v50 += v50
                        v50 += 18446744073709551615
                        v51 := GenS12{F0: v27, F1: 606, F2: 42}
                        _ = v51
                    } else {
                        for fi52 := 0; fi52 < 2; fi52++ {
                            v53 := GenS13{F0: -7, F1: (v4 - v4), F2: (v50 & 4294967295)}
                            _ = fi52
                            _ = v53
                        }
                    }
                default:
                    v2 -= 1e10
                }
                fmt.Println(v4, v4)
                for fi54 := 0; fi54 < 18; fi54++ {
                    v5 = 9223372036854775807
                    switch "\n\t" {
                    case "\n\t":
                        v55 := []int{}
                        _ = v55
                    case "":
                        v56 := []int{}
                        _ = v56
                    default:
                        var v57 int64 = v27
                        _ = v57
                    }
                    _ = fi54
                }
            } else {
                v58 := []int{(v5 << ((899) & 63)), v5}
                v27 = v27
                _ = v58
            }
            _ = v50
        }
        for fi59 := 0; fi59 < 4; fi59++ {
            for fi60 := 0; fi60 < 1; fi60++ {
                var v61 bool = false
                _ = fi60
                _ = v61
            }
            v62 := v1[-2147483648]
            v2 = ((v2 + float32((0.1 + 3.14159))) / 1.5)
            _ = fi59
            _ = v62
        }
        v27 -= (v27 | (v27 & (v27 >> ((65535) & 63))))
    } else {
        v27 += 1
    }
    for fi63 := 0; fi63 < 10; fi63++ {
        fmt.Println(fi63)
        v5 -= (v5 / ((-863) | 1))
        _ = fi63
    }
    v64 := GenS8{F0: true, F1: (v5 >> ((42) & 63))}
    for fi65 := 0; fi65 < 6; fi65++ {
        v4 = (v4 ^ v4)
        v27 -= v27
        if (false || false) {
            v5 = v5
            v4 = v4
        } else {
            if false {
                for fi66 := 0; fi66 < 17; fi66++ {
                    v27 += v27
                    delete(v3, ("\n\t" + "foo bar"))
                    _ = fi66
                }
                if true {
                    switch 1 {
                    case 32768:
                        var v67 int32 = 2147483647
                        _ = v67
                    case 1:
                        for fi68 := 0; fi68 < 1; fi68++ {
                            if (!true) {
                                v69 := GenS1{F0: "\n\t"}
                                if true {
                                    var v70 int8 = 0
                                    if false {
                                        var v71 int16 = -32768
                                        switch "" {
                                        case "":
                                            var v72 int64 = v27
                                            _ = v72
                                        case "foo bar":
                                            fmt.Println(v5, fi68, v4)
                                        }
                                        _ = v71
                                    } else {
                                        if false {
                                            v5 += fi65
                                            v5 -= 2147483647
                                        } else {
                                            switch ("a" + ("hello" + "\n\t")) {
                                            case "a":
                                                v73 := GenS14{F0: -270}
                                                _ = v73
                                            case "hello":
                                                v74 := []int{(fi65 + (fi68 ^ v5)), fi65}
                                                var v75 float64 = 1.5
                                                _ = v74
                                                _ = v75
                                            default:
                                                if true {
                                                    var v76 float64 = float64(v2)
                                                    _ = v76
                                                } else {
                                                    var v77 string = "a"
                                                    v4 = v4
                                                    v78 := GenS15{F0: true, F1: 0}
                                                    _ = v77
                                                    _ = v78
                                                }
                                            }
                                            var v79 bool = (32767 == -1)
                                            v4 = (v4 / (((^(v4))) | 1))
                                            _ = v79
                                        }
                                        var v80 int32 = 42
                                        _ = v80
                                    }
                                    var v81 int16 = 42
                                    var v82 bool = (!false)
                                    _ = v70
                                    _ = v81
                                    _ = v82
                                } else {
                                    var v83 uint16 = v4
                                    var v84 bool = (false && false)
                                    _ = v83
                                    _ = v84
                                }
                                switch "\n\t" {
                                case "hello":
                                    v27 += (v27 % ((v27) | 1))
                                    if true {
                                        v85 := GenS16{F0: 255}
                                        var v86 int16 = 42
                                        if false {
                                            v87 := GenS17{F0: 1, F1: (v4 + 0)}
                                            _ = v87
                                        }
                                        fmt.Println(v4, v4)
                                        _ = v85
                                        _ = v86
                                    } else {
                                        v4 -= v4
                                        fmt.Println(v2, v5)
                                    }
                                    var v88 int8 = -15
                                    _ = v88
                                }
                                v2 = float32(-2.25)
                                _ = v69
                            }
                            switch 1 {
                            case 128:
                                var v89 bool = false
                                _ = v89
                            default:
                                v90 := []int{v5, v5}
                                _ = v90
                            }
                            _ = fi68
                        }
                    }
                    fmt.Println(v27, v5)
                    switch v5 {
                    case 127:
                        if ((false && false) && true) {
                            if true {
                                switch 765 {
                                case 719:
                                    var v91 int = (v5 | -1)
                                    _ = v91
                                case 128:
                                    fmt.Println(fi65, fi65, fi65)
                                case -852:
                                    v92 := GenS18{F0: 72, F1: 447, F2: -1}
                                    _ = v92
                                }
                                var v93 uint32 = 1
                                _ = v93
                            } else {
                                v94 := GenS19{F0: (false && (true && (true && true))), F1: (v5 ^ (v5 & 32768))}
                                var v95 string = ("a" + "foo bar")
                                v96 := GenS20{F0: (127 >= 255)}
                                _ = v94
                                _ = v95
                                _ = v96
                            }
                            if true {
                                for fi97 := 0; fi97 < 20; fi97++ {
                                    v98 := GenS2{F0: 0, F1: 2, F2: (fi97 >> ((0) & 63))}
                                    var v99 uint8 = 128
                                    _ = fi97
                                    _ = v98
                                    _ = v99
                                }
                            } else {
                                v100 := GenS3{F0: (^(v4)), F1: 42}
                                var v101 int8 = 100
                                _ = v100
                                _ = v101
                            }
                            delete(v3, "hello")
                        } else {
                            v2 -= v2
                        }
                    case 0:
                        fmt.Println(v4, v2)
                    }
                }
                for fi102 := 0; fi102 < 16; fi102++ {
                    if ((v4 > v4) || false) {
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(fi102, v4, v2)
                        case "hello":
                            switch (-(fi102)) {
                            case -834:
                                var v103 int8 = 118
                                _ = v103
                            case 32768:
                                v104 := []int{}
                                _ = v104
                            default:
                                v105 := []int{(fi102 << ((255) & 63)), fi102}
                                _ = v105
                            }
                        case "":
                            delete(v1, fi65)
                        }
                        v106 := GenS18{F0: 1, F1: 2147483647, F2: 0}
                        _ = v106
                    }
                    v107 := GenS21{F0: "foo bar", F1: 0}
                    var v108 int16 = 42
                    _ = fi102
                    _ = v107
                    _ = v108
                }
            }
            if (v5 != v5) {
                v109 := []int{}
                var v110 float32 = v2
                _ = v109
                _ = v110
            } else {
                fmt.Println(v27, v2, v27)
                var v111 string = ""
                v112 := GenS22{F0: 65535}
                _ = v111
                _ = v112
            }
        }
        _ = fi65
    }
    v113 := GenS19{F0: (4294967295 != 808), F1: (-(v5))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v113
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
