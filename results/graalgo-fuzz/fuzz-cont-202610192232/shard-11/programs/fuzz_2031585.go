package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int; F1 string; F2 uint }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int; F1 uint64; F2 bool }
type GenS5 struct { F0 uint32; F1 uint; F2 int16 }
type GenS6 struct { F0 uint16; F1 int32; F2 int64 }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 bool; F1 int32; F2 string }
type GenS10 struct { F0 int32; F1 int32 }
type GenS11 struct { F0 bool; F1 int; F2 int8 }
type GenS12 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS13 struct { F0 int16; F1 int; F2 uint8 }
type GenS14 struct { F0 uint64; F1 uint8 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int; F1 int8 }
func genFn0() int16 {
    var v7 uint16 = 816
    v7 -= (^(v7))
    v7 += 737
    v8 := []int{}
    var v9 int16 = 268
    _ = v7
    _ = v8
    _ = v9
    return -354
}
func genFn1() bool {
    var v10 int16 = 1
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 float64 = 0.0
        fmt.Println(v12, v10, fi11)
        for fi13 := 0; fi13 < 20; fi13++ {
            v12 += v12
            v12 = (0.0 + (((float64(((1.5 * -2.25) + (((0.1 + (float32((0.1 - float64(-2.25))) + 0.1)) / float32(v12)) / 1e10))) - v12) * float64(3.14159)) / 0.1))
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    if (false || true) {
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 int32 = 1
            if (!((!false) || (true && true))) {
                var v16 bool = false
                fmt.Println(fi14, v15)
                for fi17 := 0; fi17 < 17; fi17++ {
                    v18 := GenS1{F0: -79, F1: 1}
                    v10 -= 32767
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
            v10 -= (v10 ^ 1)
            _ = fi14
            _ = v15
        }
    } else {
        var v19 int = -702
        var v20 uint8 = 144
        _ = v19
        _ = v20
    }
    _ = v10
    return (((0.0 / (1.5 - float64((0.1 - 3.14159)))) / 1e10) > -2.25)
}

func main() {
    var v1 int = -965
    v2 := []int{v1, (v1 | (v1 / ((v1) | 1))), (v1 + (^(v1)))}
    var v3 int8 = -1
    var v4 string = ("a" + "hello")
    v5 := []uint32{208, 4294967295}
    v6 := make(map[string]bool)
    for fi21 := 0; fi21 < 18; fi21++ {
        fmt.Println(v3)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi22 := 0; fi22 < 6; fi22++ {
            switch v4 {
            case "":
                v1 = (fi21 ^ 643)
            case "hello":
                _ = len(v6)
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi22
        }
        _ = fi21
    }
    if true {
        fmt.Println(v3, v1)
    } else {
        v3 -= (v3 ^ v3)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v23 := GenS2{F0: (v1 >> ((584) & 63)), F1: ("" + v4), F2: 18446744073709551615}
    for fi24 := 0; fi24 < 7; fi24++ {
        v4 = v4
        _ = fi24
    }
    if (!false) {
        switch (v1 / ((9223372036854775807) | 1)) {
        case 4294967295:
            if (!(((0.1 - 0.0) - -2.25) > 1e10)) {
                v25 := v6[(((v4 + (v4 + v4)) + (((("" + (v4 + v4)) + (v4 + "\n\t")) + "\n\t") + v4)) + v4)]
                if (((!true) || (883 != 817)) || v25) {
                    for fi26 := 0; fi26 < 14; fi26++ {
                        var v27 string = (v4 + v4)
                        if v25 {
                            if ((653 != v1) || (32767 >= 0)) {
                                v1 = v1
                                if true {
                                    for fi28 := 0; fi28 < 6; fi28++ {
                                        v2 = append(v2, -2147483648)
                                        _ = fi28
                                    }
                                    v25 = v25
                                }
                                if (-1 >= -107) {
                                    if (1e10 < 0.1) {
                                        if v25 {
                                            var v29 bool = (!v25)
                                            v29 = true
                                            v29 = (!true)
                                            _ = v29
                                        } else {
                                            switch v27 {
                                            case "a":
                                                v30 := GenS2{F0: v1, F1: v27, F2: 2}
                                                _ = v30
                                            default:
                                                fmt.Println(v27)
                                            }
                                            fmt.Println(v25, v1, fi26)
                                            var v31 uint32 = 2
                                            _ = v31
                                        }
                                    } else {
                                        if (!false) {
                                            if len(v2) > 0 {
                                                v2[0] = v1
                                            }
                                            var v32 string = v4
                                            _ = v32
                                        }
                                    }
                                } else {
                                    var v33 uint16 = 662
                                    _ = v33
                                }
                            }
                            for fi34 := 0; fi34 < 19; fi34++ {
                                v35 := GenS1{F0: v3, F1: (v3 - (v3 ^ v3))}
                                _ = fi34
                                _ = v35
                            }
                            switch (v27 + ("\n\t" + v27)) {
                            case "foo bar":
                                var v36 bool = false
                                _ = v36
                            }
                        }
                        if v25 {
                            v37 := v6[v27]
                            var v38 uint32 = 1
                            _ = v37
                            _ = v38
                        } else {
                            v3 = 0
                            v39 := GenS3{F0: 0, F1: 704, F2: 2}
                            _ = v39
                        }
                        _ = fi26
                        _ = v27
                    }
                    if true {
                        var v40 bool = (-9223372036854775808 <= -9223372036854775808)
                        if (false || v25) {
                            v3 = (v3 / ((v3) | 1))
                        } else {
                            _ = len(v6)
                            v2 = append(v2, v1)
                        }
                        var v41 float32 = float32(((1.5 * -2.25) / 1.0))
                        _ = v40
                        _ = v41
                    } else {
                        var v42 int64 = 0
                        if (genFn0() <= -261) {
                            for fi43 := 0; fi43 < 16; fi43++ {
                                if v25 {
                                    switch -1 {
                                    case 4294967295:
                                        fmt.Println(v42, fi43, fi43)
                                    case -383:
                                        switch ("hello" + "a") {
                                        case "a":
                                            var v44 uint64 = 0
                                            _ = v44
                                        default:
                                            var v45 float64 = 0.0
                                            _ = v45
                                        }
                                    default:
                                        for fi46 := 0; fi46 < 14; fi46++ {
                                            v47 := GenS1{F0: 100, F1: v3}
                                            var v48 bool = (v25 || (2 < 1))
                                            fmt.Println(v3)
                                            _ = fi46
                                            _ = v47
                                            _ = v48
                                        }
                                    }
                                    fmt.Println(fi43)
                                } else {
                                    var v49 uint = 584
                                    _ = v49
                                }
                                var v50 uint = 4294967295
                                _ = fi43
                                _ = v50
                            }
                            v42 -= (v42 % ((-892) | 1))
                            switch v1 {
                            case -574:
                                v4 = v4
                            case -9223372036854775808:
                                var v51 string = "a"
                                _ = v51
                            case 32768:
                                v42 = 0
                            default:
                                v52 := GenS4{F0: (v1 - v1), F1: 0, F2: (v25 || false)}
                                _ = v52
                            }
                        } else {
                            for fi53 := 0; fi53 < 10; fi53++ {
                                if (-32768 < 32767) {
                                    switch v4 {
                                    case "":
                                        delete(v6, "")
                                    case "\n\t":
                                        for fi54 := 0; fi54 < 5; fi54++ {
                                            if (!(2 >= 1)) {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v55 := GenS5{F0: 931, F1: 255, F2: genFn0()}
                                                _ = v55
                                            }
                                            _ = fi54
                                        }
                                    default:
                                        var v56 float64 = -2.25
                                        _ = v56
                                    }
                                    _ = len(v5)
                                }
                                _ = fi53
                            }
                            fmt.Println(v25)
                        }
                        var v57 uint = 255
                        _ = v42
                        _ = v57
                    }
                    var v58 bool = (197 != 2)
                    _ = v58
                }
                _ = v25
            }
        }
        if false {
            v2 = append(v2, v1)
        } else {
            for fi59 := 0; fi59 < 13; fi59++ {
                v4 = v4
                v60 := GenS6{F0: 78, F1: 865, F2: 42}
                if ((true && (!true)) || false) {
                    switch fi59 {
                    case 4294967295:
                        fmt.Println(v1, v1, fi59)
                    case -1:
                        v4 = v4
                    default:
                        fmt.Println(v4, v1, v1)
                    }
                    var v61 float64 = ((1.5 - (float64((1e10 + (-2.25 + 1.5))) / 1.0)) / 1.5)
                    for fi62 := 0; fi62 < 1; fi62++ {
                        v6[(v4 + ("foo bar" + "hello"))] = false
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if (((!(v3 != 0)) || true) || true) {
                            v63 := GenS7{F0: 647, F1: true}
                            if ((!true) && true) {
                                v4 = "foo bar"
                            }
                            switch fi62 {
                            case 42:
                                fmt.Println(v4, fi59)
                            case 695:
                                v64 := v6[(v4 + "")]
                                _ = v64
                            default:
                                v3 -= (v3 % ((0) | 1))
                            }
                            v4 = (v4 + (v4 + (v4 + v4)))
                            _ = len(v2)
                            _ = v63
                        } else {
                            var v65 uint16 = 255
                            _ = v65
                        }
                        if true {
                            v3 = 42
                            if true {
                                v3 -= (v3 ^ (v3 ^ (v3 + v3)))
                            }
                        }
                        var v66 uint8 = 128
                        _ = fi62
                        _ = v66
                    }
                    _ = v61
                }
                _ = fi59
                _ = v60
            }
            var v67 uint32 = 65535
            for fi68 := 0; fi68 < 10; fi68++ {
                v3 = v3
                for fi69 := 0; fi69 < 18; fi69++ {
                    switch 9223372036854775807 {
                    case 4294967295:
                        if (true || ((fi69 | fi69) == (fi68 / ((fi68) | 1)))) {
                            var v70 uint = 778
                            _ = len(v6)
                            _ = v70
                        }
                    case -777:
                        v71 := GenS8{F0: (^(v67))}
                        _ = v71
                    default:
                        v4 = (v4 + (v4 + "a"))
                    }
                    v72 := GenS9{F0: (-1 < v3), F1: 42, F2: "a"}
                    if true {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v73 uint = 255
                        _ = v73
                    } else {
                        var v74 uint32 = 2
                        v3 = (-(v3))
                        v75 := GenS10{F0: -2147483648, F1: 42}
                        _ = v74
                        _ = v75
                    }
                    fmt.Println(fi68, fi68, v1)
                    var v76 string = v4
                    if true {
                        if false {
                            v4 = (((v76 + v76) + v76) + v76)
                            _ = len(v6)
                        }
                        v77 := GenS9{F0: ((0.1 - 1.5) > 0.1), F1: -2147483648, F2: "a"}
                        fmt.Println(fi69)
                        _ = v77
                    } else {
                        fmt.Println(v67, v3, v1)
                    }
                    _ = fi69
                    _ = v72
                    _ = v76
                }
                v67 += v67
                _ = fi68
            }
            _ = v67
        }
    }
    _ = len(v5)
    switch "" {
    case "hello":
        if (32767 <= -570) {
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                v78 := GenS2{F0: 127, F1: v4, F2: 107}
                switch (v1 << ((865) & 63)) {
                case 616:
                    v4 = "a"
                case 32768:
                    if (!false) {
                        if (true && false) {
                            for fi79 := 0; fi79 < 5; fi79++ {
                                if true {
                                    for fi80 := 0; fi80 < 20; fi80++ {
                                        switch v4 {
                                        case "a":
                                            fmt.Println(v4, v3)
                                        case "foo bar":
                                            v81 := v6[v4]
                                            _ = v81
                                        }
                                        _ = fi80
                                    }
                                    var v82 float64 = float64(1e10)
                                    _ = v82
                                } else {
                                    if (!((((0.1 - float32(1e10)) * (-2.25 + 0.1)) / float32(3.14159)) >= 0.1)) {
                                        if true {
                                            v83 := GenS10{F0: 1, F1: -1}
                                            delete(v6, v4)
                                            v84 := GenS5{F0: 2147483648, F1: 255, F2: genFn0()}
                                            _ = v83
                                            _ = v84
                                        } else {
                                            fmt.Println(v3)
                                            var v85 float64 = -2.25
                                            _ = v85
                                        }
                                    }
                                    v4 = "hello"
                                }
                                var v86 float64 = -2.25
                                _ = fi79
                                _ = v86
                            }
                            v1 = (v1 * (^(v1)))
                            var v87 uint32 = 0
                            _ = v87
                        } else {
                            if (v3 == (v3 * v3)) {
                                v88 := GenS11{F0: true, F1: (v1 / ((870) | 1)), F2: v3}
                                v89 := GenS12{F0: 128, F1: 795, F2: 9223372036854775807}
                                for fi90 := 0; fi90 < 1; fi90++ {
                                    var v91 string = "hello"
                                    switch (v91 + "") {
                                    case "hello":
                                        var v92 int16 = -38
                                        _ = v92
                                    }
                                    _ = fi90
                                    _ = v91
                                }
                                _ = v88
                                _ = v89
                            } else {
                                if true {
                                    var v93 float32 = 0.1
                                    _ = v93
                                }
                            }
                            switch (v1 % ((v1) | 1)) {
                            case 128:
                                for fi94 := 0; fi94 < 13; fi94++ {
                                    var v95 uint = 255
                                    _ = fi94
                                    _ = v95
                                }
                            default:
                                if false {
                                    if (1e10 > (0.1 + (0.1 / 0.1))) {
                                        v96 := GenS13{F0: -32768, F1: (^(v1)), F2: 0}
                                        v1 += 32768
                                        _ = v96
                                    } else {
                                        fmt.Println(v1, v4)
                                        switch ((v4 + (v4 + ("\n\t" + "foo bar"))) + v4) {
                                        case "foo bar":
                                            var v97 float32 = 3.14159
                                            _ = v97
                                        case "":
                                            v98 := GenS8{F0: 242}
                                            _ = v98
                                        default:
                                            v3 -= (-(v3))
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            v99 := GenS2{F0: (v1 % ((v1) | 1)), F1: "a", F2: 65535}
                                            _ = v99
                                        case "a":
                                            for fi100 := 0; fi100 < 8; fi100++ {
                                                switch v4 {
                                                case "":
                                                    for fi101 := 0; fi101 < 6; fi101++ {
                                                        v1 = 645
                                                        fmt.Println(v3, v4, fi101)
                                                        _ = fi101
                                                    }
                                                case "a":
                                                    var v102 uint = 255
                                                    _ = v102
                                                default:
                                                    v4 = v4
                                                }
                                                fmt.Println(v3, v4)
                                                _ = fi100
                                            }
                                        }
                                    }
                                } else {
                                    for fi103 := 0; fi103 < 7; fi103++ {
                                        v104 := GenS14{F0: 18446744073709551615, F1: 0}
                                        _ = fi103
                                        _ = v104
                                    }
                                }
                            }
                            if ((false && ((42 == -250) && (true || true))) && (1 != 1)) {
                                v105 := GenS15{F0: 42}
                                fmt.Println(v3, v3)
                                _ = v105
                            } else {
                                fmt.Println(v1)
                                var v106 int8 = (v3 | (v3 & v3))
                                var v107 uint8 = 128
                                _ = v106
                                _ = v107
                            }
                        }
                    }
                case -280:
                    var v108 uint32 = 266
                    _ = v108
                }
                v3 -= (^(v3))
                _ = v78
            }
            _ = len(v6)
            var v109 uint16 = 32767
            _ = v109
        } else {
            switch (v1 - v1) {
            case 1:
                switch v4 {
                case "hello":
                    v110 := GenS16{F0: 627, F1: 42}
                    _ = v110
                default:
                    v4 = "a"
                }
            case 2147483647:
                switch v4 {
                case "a":
                    var v111 float64 = 0.0
                    _ = v111
                case "hello":
                    v112 := GenS11{F0: (!false), F1: v1, F2: (-(v3))}
                    _ = v112
                default:
                    for fi113 := 0; fi113 < 15; fi113++ {
                        for fi114 := 0; fi114 < 5; fi114++ {
                            fmt.Println(v1, v3)
                            if (("foo bar" + "\n\t") > ("" + "")) {
                                v4 = v4
                                var v115 uint = 65535
                                var v116 uint64 = 652
                                _ = v115
                                _ = v116
                            }
                            if true {
                                fmt.Println(fi113, v3)
                            }
                            _ = fi114
                        }
                        v4 = v4
                        var v117 uint16 = 32767
                        _ = fi113
                        _ = v117
                    }
                }
            case 32767:
                for fi118 := 0; fi118 < 17; fi118++ {
                    _ = len(v6)
                    for fi119 := 0; fi119 < 6; fi119++ {
                        v120 := v6[(v4 + v4)]
                        var v121 float32 = 1.5
                        _ = fi119
                        _ = v120
                        _ = v121
                    }
                    _ = fi118
                }
            }
        }
    }
    for fi122 := 0; fi122 < 17; fi122++ {
        if len(v2) > 0 {
            v2[0] = -443
        }
        var v123 int32 = -323
        _ = fi122
        _ = v123
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
