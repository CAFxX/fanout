package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int16; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 int8; F2 uint }
type GenS6 struct { F0 int8; F1 int64; F2 uint32 }
type GenS7 struct { F0 int64; F1 uint64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int16; F1 bool }
type GenS10 struct { F0 int8; F1 int }
type GenS11 struct { F0 uint8 }
func genFn0(p0 int, p1 int32, p2 int32) int {
    p1 = -169
    v5 := GenS1{F0: 0}
    for fi6 := 0; fi6 < 12; fi6++ {
        switch p0 {
        case 2147483647:
            v7 := []int{}
            _ = v7
        case 1:
            if false {
                v8 := GenS1{F0: 1}
                if ("hello" <= ("\n\t" + (("" + ("\n\t" + "")) + "a"))) {
                    v9 := GenS1{F0: 0}
                    if ((42 < 42) || (!true)) {
                        if (!(1 >= fi6)) {
                            p0 -= 54
                            v10 := []int{p0}
                            _ = v10
                        } else {
                            var v11 string = "foo bar"
                            for fi12 := 0; fi12 < 3; fi12++ {
                                var v13 int = (fi12 | fi6)
                                v14 := []int{v13, (p0 << ((0) & 63))}
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                            _ = v11
                        }
                    }
                    _ = v9
                } else {
                    v15 := []int{(p0 % ((fi6) | 1))}
                    _ = v15
                }
                _ = v8
            } else {
                var v16 uint8 = 127
                p2 = (p1 << ((42) & 63))
                p2 = p2
                _ = v16
            }
        }
        v17 := []int{-541, p0, p0}
        p1 += 42
        _ = fi6
        _ = v17
    }
    if false {
        for fi18 := 0; fi18 < 5; fi18++ {
            switch "foo bar" {
            case "":
                v19 := GenS1{F0: 255}
                _ = v19
            default:
                var v20 int64 = 1
                _ = v20
            }
            var v21 float32 = float32(0.1)
            _ = fi18
            _ = v21
        }
        var v22 int8 = 120
        _ = v22
    } else {
        fmt.Println(p2)
    }
    var v23 int16 = -946
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v5
    return p0
}

func main() {
    var v1 string = ("a" + "foo bar")
    var v2 int8 = 110
    var v3 float64 = float64(3.14159)
    v4 := make(map[string]int)
    v24 := GenS2{F0: 68}
    if (1 != 255) {
        v3 = v3
    }
    var v25 int16 = 757
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 uint32 = 181
        _ = fi26
        _ = v27
    }
    switch 42 {
    case -328:
        v28 := GenS2{F0: 65535}
        _ = v28
    case 293:
        if false {
            if (0.0 >= float32(v3)) {
                v25 = v25
            }
            if true {
                if false {
                    switch ((v1 + ((v1 + v1) + ("" + v1))) + ("foo bar" + (v1 + ((v1 + v1) + v1)))) {
                    case "\n\t":
                        v29 := []int{-1, 9223372036854775807, genFn0(-955, -2147483648, 921), genFn0(2147483647, 1, 1)}
                        _ = v29
                    case "a":
                        v30 := GenS3{F0: 736, F1: (-(v25)), F2: 255}
                        _ = v30
                    case "hello":
                        switch ("\n\t" + ("" + (("hello" + v1) + v1))) {
                        case "":
                            if false {
                                switch 32767 {
                                case 0:
                                    v2 = v2
                                case -636:
                                    fmt.Println(v3)
                                case -258:
                                    var v31 uint64 = 1
                                    _ = v31
                                default:
                                    if (-321 != 4294967295) {
                                        if (42 < 792) {
                                            var v32 uint32 = 4294967295
                                            fmt.Println(v2, v1, v3)
                                            _ = v32
                                        }
                                        if (4294967295 > 65535) {
                                            v33 := []int{genFn0(1, 0, int32(v25)), 127, genFn0(0, 0, -498)}
                                            var v34 uint = 1
                                            v3 = (3.14159 + 0.1)
                                            _ = v33
                                            _ = v34
                                        }
                                    } else {
                                        switch "foo bar" {
                                        case "\n\t":
                                            v3 -= v3
                                        case "foo bar":
                                            v1 = (("a" + (((((((v1 + v1) + "a") + v1) + "\n\t") + "hello") + (v1 + v1)) + v1)) + v1)
                                        default:
                                            v25 += v25
                                        }
                                    }
                                }
                            } else {
                                v35 := GenS4{F0: false}
                                v1 = "hello"
                                v36 := []int{518, 32768, -9223372036854775808}
                                _ = v35
                                _ = v36
                            }
                        case "foo bar":
                            for fi37 := 0; fi37 < 12; fi37++ {
                                for fi38 := 0; fi38 < 2; fi38++ {
                                    v2 -= v2
                                    _ = fi38
                                }
                                var v39 uint32 = 2
                                _ = fi37
                                _ = v39
                            }
                        case "a":
                            fmt.Println(v25, v25)
                        }
                    }
                    var v40 string = ""
                    for fi41 := 0; fi41 < 2; fi41++ {
                        var v42 string = v40
                        _ = fi41
                        _ = v42
                    }
                    _ = v40
                } else {
                    var v43 uint16 = 1
                    _ = v43
                }
                for fi44 := 0; fi44 < 14; fi44++ {
                    switch (v1 + ((v1 + v1) + (v1 + (("hello" + "a") + v1)))) {
                    case "":
                        if true {
                            for fi45 := 0; fi45 < 11; fi45++ {
                                v46 := []int{}
                                v47 := GenS4{F0: (false || (42 >= 0))}
                                _ = fi45
                                _ = v46
                                _ = v47
                            }
                            var v48 uint8 = 255
                            _ = v48
                        }
                    default:
                        if true {
                            for fi49 := 0; fi49 < 4; fi49++ {
                                if ((true && (v3 > v3)) || (-1 > 17)) {
                                    _ = len(v4)
                                    var v50 bool = (2 <= 2)
                                    v1 = v1
                                    v51 := GenS4{F0: v50}
                                    if (!(v3 > v3)) {
                                        var v52 int = fi44
                                        v53 := GenS2{F0: 65535}
                                        var v54 int64 = 1
                                        _ = v52
                                        _ = v53
                                        _ = v54
                                    } else {
                                        v55 := GenS3{F0: 67, F1: 1000, F2: 135}
                                        v3 = v3
                                        fmt.Println(v3, v50, v50)
                                        fmt.Println(v2, v50, v1)
                                        _ = v55
                                    }
                                    _ = v50
                                    _ = v51
                                } else {
                                    if false {
                                        switch fi49 {
                                        case 621:
                                            var v56 uint = 4294967295
                                            _ = v56
                                        }
                                        fmt.Println(v2, fi44, fi49)
                                        var v57 int32 = 42
                                        v57 -= v57
                                        _ = v57
                                    }
                                    if (42 != 255) {
                                        var v58 bool = false
                                        var v59 uint = 106
                                        _ = v58
                                        _ = v59
                                    } else {
                                        v60 := GenS5{F0: 4294967295, F1: (v2 ^ v2), F2: 2}
                                        _ = v60
                                    }
                                }
                                fmt.Println(v2, fi49, v25)
                                fmt.Println(v25, v3, fi44)
                                _ = fi49
                            }
                        } else {
                            var v61 uint = 480
                            fmt.Println(v3, v25)
                            _ = v61
                        }
                    }
                    var v62 uint16 = 2
                    _ = fi44
                    _ = v62
                }
                var v63 uint64 = 80
                _ = v63
            } else {
                v1 = ""
                v2 += -126
                var v64 uint8 = 0
                _ = v64
            }
            for fi65 := 0; fi65 < 2; fi65++ {
                switch 127 {
                case 32767:
                    for fi66 := 0; fi66 < 8; fi66++ {
                        var v67 int8 = (v2 & (-(v2)))
                        var v68 bool = true
                        var v69 int16 = v25
                        delete(v4, ("hello" + "foo bar"))
                        var v70 int16 = v25
                        _ = fi66
                        _ = v67
                        _ = v68
                        _ = v69
                        _ = v70
                    }
                default:
                    v3 = 3.14159
                }
                _ = fi65
            }
        }
    case 9223372036854775807:
        switch 1 {
        case -2147483648:
            for fi71 := 0; fi71 < 11; fi71++ {
                v72 := GenS5{F0: 65535, F1: (v2 ^ v2), F2: 306}
                _ = fi71
                _ = v72
            }
        case 2147483647:
            fmt.Println(v2)
        }
    default:
        switch 9223372036854775807 {
        case 585:
            v1 = v1
        case 128:
            v25 += 515
        case 32768:
            if (!(false || (!(false && ("hello" > (v1 + "a")))))) {
                var v73 uint = 18446744073709551615
                var v74 string = ""
                for fi75 := 0; fi75 < 18; fi75++ {
                    var v76 uint16 = 1
                    for fi77 := 0; fi77 < 6; fi77++ {
                        var v78 int32 = 0
                        _ = fi77
                        _ = v78
                    }
                    v1 = v74
                    _ = fi75
                    _ = v76
                }
                _ = v73
                _ = v74
            }
        default:
            if ((false || false) && true) {
                v4[v1] = -9223372036854775808
                v4[v1] = -9223372036854775808
                var v79 float32 = (((((1e10 - 0.1) - 1.5) / 0.1) - (float32((1.5 - (float64((-2.25 / 3.14159)) * (0.1 + 1.5)))) * float32(-2.25))) + ((float32(0.1) / 3.14159) / ((1.5 - 3.14159) / (0.1 / 1.0))))
                _ = v79
            } else {
                if true {
                    switch 9223372036854775807 {
                    case 42:
                        var v80 string = (v1 + (v1 + v1))
                        _ = v80
                    case 0:
                        var v81 uint = 999
                        _ = v81
                    case 32768:
                        v82 := GenS6{F0: (v2 | 45), F1: 0, F2: 65535}
                        _ = v82
                    }
                    switch v1 {
                    case "a":
                        var v83 uint16 = 505
                        _ = v83
                    case "":
                        var v84 uint = 705
                        _ = v84
                    case "hello":
                        v3 -= ((((3.14159 + v3) + 0.0) / v3) / (v3 + (-2.25 * v3)))
                    }
                } else {
                    fmt.Println(v2, v1, v25)
                    fmt.Println(v1, v2)
                }
            }
        }
    }
    v3 -= (1e10 / (v3 + 3.14159))
    v25 += (v25 >> ((0) & 63))
    var v85 int64 = 57
    v86 := GenS4{F0: true}
    v87 := GenS7{F0: v85, F1: 18446744073709551615}
    if false {
        v88 := []int{}
        v89 := v4[""]
        if true {
            var v90 int16 = (v25 ^ -438)
            v91 := GenS1{F0: 255}
            v92 := v4[v1]
            _ = v90
            _ = v91
            _ = v92
        } else {
            v93 := GenS8{F0: 65535}
            switch (v89 * (v89 | v89)) {
            case 4294967295:
                var v94 string = v1
                _ = v94
            case 2147483647:
                var v95 uint32 = 1
                _ = v95
            case 128:
                var v96 bool = (false || (2 > 128))
                _ = v96
            }
            _ = v93
        }
        _ = v88
        _ = v89
    } else {
        for fi97 := 0; fi97 < 19; fi97++ {
            var v98 uint64 = 42
            _ = fi97
            _ = v98
        }
        fmt.Println(v2, v3)
        v1 = v1
    }
    v99 := GenS9{F0: v25, F1: true}
    v85 = v85
    switch v1 {
    case "\n\t":
        v3 += float64(1.5)
    default:
        v3 = (float64(3.14159) + 0.0)
    }
    fmt.Println(v85, v25, v85)
    switch int(v2) {
    case 32768:
        for fi100 := 0; fi100 < 17; fi100++ {
            switch 698 {
            case -1:
                if false {
                    var v101 uint = 776
                    for fi102 := 0; fi102 < 13; fi102++ {
                        switch fi100 {
                        case 127:
                            v103 := GenS10{F0: (v2 << ((v101) & 63)), F1: -2147483648}
                            _ = v103
                        case -2147483648:
                            switch "hello" {
                            case "":
                                v104 := GenS9{F0: 42, F1: (v101 != (v101 / (((v101 ^ (^(v101)))) | 1)))}
                                _ = v104
                            case "hello":
                                fmt.Println(v1, v85, v1)
                            }
                        default:
                            v105 := GenS1{F0: v101}
                            _ = v105
                        }
                        v106 := v4[v1]
                        v101 = v101
                        v107 := []int{int(v25), 128, (fi102 >> ((4294967295) & 63)), genFn0(fi102, 0, -2147483648)}
                        var v108 uint64 = 859
                        var v109 int = int(v108)
                        _ = fi102
                        _ = v106
                        _ = v107
                        _ = v108
                        _ = v109
                    }
                    if true {
                        if false {
                            var v110 uint64 = 0
                            if (!(!false)) {
                                if ((v3 - v3) < float64(((0.1 * 0.0) + (0.0 + 1.5)))) {
                                    fmt.Println(v25, v101)
                                    v110 -= v110
                                    for fi111 := 0; fi111 < 7; fi111++ {
                                        var v112 bool = (genFn0(fi111, 1, -1) == (fi100 | fi111))
                                        _ = fi111
                                        _ = v112
                                    }
                                } else {
                                    if true {
                                        v4["hello"] = -923
                                        v113 := []int{(fi100 / ((9223372036854775807) | 1))}
                                        _ = v113
                                    }
                                }
                                switch (fi100 & (fi100 | -59)) {
                                case -2147483648:
                                    var v114 int32 = 0
                                    _ = v114
                                case 9223372036854775807:
                                    v2 -= (v2 ^ v2)
                                }
                                v115 := GenS9{F0: v25, F1: true}
                                _ = v115
                            } else {
                                var v116 uint16 = 42
                                v117 := []int{fi100, (-(fi100))}
                                _ = v116
                                _ = v117
                            }
                            v118 := GenS11{F0: 255}
                            _ = v110
                            _ = v118
                        }
                        v3 = (v3 + v3)
                        var v119 int16 = (v25 * 1000)
                        _ = v119
                    }
                    _ = v101
                }
            }
            switch ("" + v1) {
            case "\n\t":
                v25 = v25
            case "hello":
                v2 = int8(v85)
            case "a":
                var v120 bool = (!(!(true || true)))
                _ = v120
            }
            _ = fi100
        }
    case 9223372036854775807:
        fmt.Println(v25, v1)
    case 577:
        var v121 bool = true
        _ = v121
    default:
        var v122 float64 = (((-2.25 * float64((0.1 * ((3.14159 / (3.14159 * 0.1)) + -2.25)))) / (v3 - v3)) + v3)
        _ = v122
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v85
    _ = v86
    _ = v87
    _ = v99
}
