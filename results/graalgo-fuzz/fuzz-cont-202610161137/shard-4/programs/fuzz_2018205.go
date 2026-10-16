package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int16 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 int; F1 uint64 }
func genFn0(p0 string, p1 uint64, p2 uint) string {
    for fi4 := 0; fi4 < 16; fi4++ {
        p0 = "\n\t"
        v5 := GenS1{F0: true, F1: false, F2: -32768}
        _ = fi4
        _ = v5
    }
    p0 = (p0 + p0)
    var v6 int = 1
    if true {
        var v7 float32 = -2.25
        var v8 int16 = -261
        if (723 == 2) {
            v7 = -2.25
            var v9 string = "hello"
            v10 := GenS1{F0: true, F1: (uint(p1) > 141), F2: (-(v8))}
            _ = v10
            _ = v9
        }
        _ = v7
        _ = v8
    }
    if true {
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 uint8 = 0
            if true {
                if false {
                    var v13 int32 = 1
                    _ = v13
                } else {
                    var v14 int32 = 449
                    _ = v14
                }
            }
            _ = fi11
            _ = v12
        }
        if ((^(p2)) != (p2 ^ 1)) {
            p0 = (((((("hello" + (p0 + (("a" + "") + p0))) + "\n\t") + p0) + p0) + p0) + p0)
        }
        switch "hello" {
        case "foo bar":
            fmt.Println(p0)
        case "hello":
            v15 := GenS1{F0: true, F1: (-2147483648 >= 12), F2: -32768}
            _ = v15
        case "\n\t":
            v16 := GenS1{F0: (p1 < (p1 & (p1 & 4294967295))), F1: true, F2: 32767}
            _ = v16
        }
    } else {
        if false {
            var v17 uint16 = 175
            v18 := GenS1{F0: (!false), F1: ((!((!false) && false)) || (32767 <= 418)), F2: 1000}
            for fi19 := 0; fi19 < 10; fi19++ {
                for fi20 := 0; fi20 < 1; fi20++ {
                    if false {
                        var v21 uint64 = 18446744073709551615
                        v22 := GenS1{F0: false, F1: ((p1 * 981) <= (v21 * 812)), F2: -631}
                        _ = v21
                        _ = v22
                    } else {
                        var v23 int8 = 42
                        if false {
                            v24 := []int{(-(fi19)), 0, (fi20 >> (((p2 & p2)) & 63))}
                            var v25 bool = true
                            if (v25 && v25) {
                                if v25 {
                                    switch p0 {
                                    case "foo bar":
                                        v25 = v25
                                    case "":
                                        switch fi20 {
                                        case 2147483647:
                                            fmt.Println(p1)
                                        case -311:
                                            v24 = append(v24, -2147483648)
                                        default:
                                            var v26 uint = (p2 ^ (^(p2)))
                                            _ = v26
                                        }
                                    }
                                    var v27 uint16 = (v17 ^ (v17 / (((v17 % (((^(v17))) | 1))) | 1)))
                                    var v28 uint8 = 42
                                    var v29 uint16 = (v17 >> ((p2) & 63))
                                    fmt.Println(v29, v29)
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                            }
                            _ = v24
                            _ = v25
                        }
                        p1 = p1
                        _ = v23
                    }
                    _ = fi20
                }
                for fi30 := 0; fi30 < 12; fi30++ {
                    if false {
                        if false {
                            var v31 string = p0
                            v31 = p0
                            v32 := []int{(v6 << ((p2) & 63)), fi30, v6}
                            _ = v31
                            _ = v32
                        }
                        v33 := []int{int(p2), fi19}
                        fmt.Println(v17, p2, v6)
                        _ = v33
                    } else {
                        v34 := []int{}
                        var v35 int8 = -27
                        v36 := GenS2{F0: (^(fi30)), F1: (187 >= 726)}
                        _ = len(v34)
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    _ = fi30
                }
                p0 = ((((("\n\t" + p0) + (p0 + "a")) + "\n\t") + "a") + p0)
                _ = fi19
            }
            _ = v17
            _ = v18
        }
        for fi37 := 0; fi37 < 10; fi37++ {
            if (((3.14159 + 0.1) - 1.5) < 1.5) {
                v6 = (fi37 + fi37)
                var v38 int = v6
                _ = v38
            } else {
                switch v6 {
                case 32768:
                    fmt.Println(fi37)
                case 2147483647:
                    fmt.Println(v6, v6, p0)
                case 699:
                    v39 := []int{(v6 & fi37), 4294967295, v6}
                    _ = v39
                default:
                    if false {
                        v40 := GenS1{F0: false, F1: (p0 > "\n\t"), F2: -843}
                        _ = v40
                    }
                }
                var v41 int16 = 513
                fmt.Println(p1, v41)
                _ = v41
            }
            _ = fi37
        }
        p1 += p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "a"
}
func genFn1(p0 int) int16 {
    switch "a" {
    case "hello":
        var v42 string = ("a" + ("a" + "foo bar"))
        _ = v42
    case "":
        fmt.Println(p0)
    case "\n\t":
        fmt.Println(p0)
    default:
        p0 -= -2147483648
    }
    for fi43 := 0; fi43 < 6; fi43++ {
        fmt.Println(fi43, p0)
        if false {
            fmt.Println(fi43)
        } else {
            switch "" {
            case "":
                v44 := GenS3{F0: 2, F1: 1, F2: 4294967295}
                _ = v44
            default:
                p0 = fi43
            }
        }
        _ = fi43
    }
    _ = p0
    return 244
}

func main() {
    var v1 int64 = 201
    var v2 uint8 = 128
    var v3 int16 = 42
    var v45 int16 = (v3 / (((-(v3))) | 1))
    switch (("\n\t" + (("\n\t" + "a") + "foo bar")) + "hello") {
    case "foo bar":
        v2 -= 1
    case "a":
        for fi46 := 0; fi46 < 11; fi46++ {
            if (-110 != -1) {
                var v47 uint32 = 65535
                _ = v47
            }
            _ = fi46
        }
    }
    var v48 uint = 150
    if false {
        v49 := GenS2{F0: 128, F1: true}
        v50 := []int{350, 4294967295}
        _ = v49
        _ = v50
    } else {
        var v51 uint8 = 42
        var v52 uint32 = 1
        _ = v51
        _ = v52
    }
    for fi53 := 0; fi53 < 5; fi53++ {
        var v54 bool = (!false)
        v55 := GenS1{F0: v54, F1: false, F2: (v45 >> ((v48) & 63))}
        var v56 uint64 = 1
        _ = fi53
        _ = v54
        _ = v55
        _ = v56
    }
    v45 = 0
    switch -2147483648 {
    case 2147483647:
        v48 -= v48
    case 1:
        for fi57 := 0; fi57 < 3; fi57++ {
            if false {
                if (42 >= 42) {
                    for fi58 := 0; fi58 < 6; fi58++ {
                        fmt.Println(v2, v48)
                        fmt.Println(fi58, v2, fi58)
                        v59 := GenS2{F0: fi58, F1: (!(fi58 > 128))}
                        _ = fi58
                        _ = v59
                    }
                }
                if (int32(v48) > 0) {
                    var v60 uint8 = v2
                    _ = v60
                }
                v61 := GenS2{F0: (fi57 & -121), F1: (0.1 == 0.0)}
                _ = v61
            } else {
                v3 = v45
                if (!false) {
                    v3 = -32768
                    for fi62 := 0; fi62 < 11; fi62++ {
                        v48 += v48
                        v2 -= 42
                        v63 := []int{1}
                        fmt.Println(v1)
                        switch fi62 {
                        case 128:
                            _ = len(v63)
                        case -2147483648:
                            if true {
                                var v64 uint32 = 65535
                                _ = v64
                            }
                        default:
                            v48 = (^(v48))
                        }
                        _ = fi62
                        _ = v63
                    }
                } else {
                    var v65 bool = true
                    _ = v65
                }
            }
            for fi66 := 0; fi66 < 5; fi66++ {
                v45 += (v45 / (((v45 << ((v48) & 63))) | 1))
                var v67 int64 = (-(v1))
                _ = fi66
                _ = v67
            }
            switch (fi57 * (fi57 - (fi57 % ((fi57) | 1)))) {
            case 767:
                v68 := []int{fi57, fi57, fi57, fi57}
                _ = v68
            default:
                switch ((("hello" + "\n\t") + "") + "\n\t") {
                case "":
                    v48 = 4294967295
                case "foo bar":
                    switch (^(fi57)) {
                    case 32768:
                        v69 := []int{128, fi57, -9223372036854775808}
                        _ = v69
                    case -726:
                        v70 := GenS3{F0: 42, F1: 166, F2: 1}
                        _ = v70
                    default:
                        if (!false) {
                            v1 += v1
                            v71 := []int{fi57, int(v2), fi57}
                            _ = v71
                        }
                    }
                case "hello":
                    for fi72 := 0; fi72 < 11; fi72++ {
                        var v73 uint16 = uint16(v2)
                        v3 += (v3 << ((v48) & 63))
                        v74 := GenS4{F0: v2}
                        v75 := GenS5{F0: genFn1((fi57 << ((v48) & 63))), F1: true}
                        _ = fi72
                        _ = v73
                        _ = v74
                        _ = v75
                    }
                default:
                    for fi76 := 0; fi76 < 16; fi76++ {
                        if false {
                            v77 := []int{fi76, (fi76 * -2147483648)}
                            v78 := GenS1{F0: false, F1: (971 == 0), F2: v3}
                            if true {
                                v79 := GenS6{F0: 0, F1: "\n\t"}
                                switch 32767 {
                                case -9223372036854775808:
                                    v80 := GenS4{F0: (^(v2))}
                                    _ = v80
                                case 42:
                                    var v81 float32 = (3.14159 - 1e10)
                                    _ = v81
                                case -2147483648:
                                    if len(v77) > 0 {
                                        v77[0] = -1
                                    }
                                }
                                fmt.Println(fi76, v1)
                                _ = v79
                            } else {
                                if (!false) {
                                    var v82 bool = (v3 <= -32768)
                                    var v83 bool = true
                                    _ = v82
                                    _ = v83
                                } else {
                                    if false {
                                        switch 918 {
                                        case -429:
                                            v2 -= 106
                                            _ = len(v77)
                                            v45 = (v45 ^ (v3 & v45))
                                            v48 = v48
                                            if false {
                                                fmt.Println(v48)
                                            } else {
                                                var v84 float64 = float64(1.5)
                                                v45 = genFn1(128)
                                                if len(v77) > 0 {
                                                    v77[0] = fi57
                                                }
                                                _ = v84
                                            }
                                        default:
                                            v85 := GenS1{F0: (-2.25 < ((((1.5 / (1e10 / 1.5)) - 0.0) - 1.5) + 1.5)), F1: true, F2: (v45 >> ((v48) & 63))}
                                            _ = v85
                                        }
                                        v77 = append(v77, (fi57 << ((65535) & 63)))
                                        v45 = v3
                                        v3 -= genFn1(fi57)
                                    }
                                    if false {
                                        switch "" {
                                        case "foo bar":
                                            switch "a" {
                                            case "\n\t":
                                                var v86 uint16 = 891
                                                _ = v86
                                            default:
                                                v45 = v3
                                            }
                                        case "a":
                                            if (int8(v1) == -128) {
                                                var v87 uint16 = 65535
                                                _ = v87
                                            }
                                        case "\n\t":
                                            v88 := GenS7{F0: v2, F1: 4294967295, F2: -1}
                                            _ = v88
                                        default:
                                            if true {
                                                v2 -= v2
                                                v77 = append(v77, (^(fi76)))
                                                v89 := GenS8{F0: 794, F1: 4294967295}
                                                _ = v89
                                            }
                                        }
                                        var v90 float32 = (float32(1.5) / 3.14159)
                                        if len(v77) >= 2 {
                                            v77 = v77[1:len(v77)]
                                        }
                                        _ = v90
                                    } else {
                                        switch fi76 {
                                        case 9223372036854775807:
                                            v91 := GenS8{F0: 0, F1: 0}
                                            _ = v91
                                        }
                                        v45 -= -32768
                                    }
                                }
                                v45 = v45
                                v77 = append(v77, fi76)
                            }
                            _ = v77
                            _ = v78
                        }
                        _ = fi76
                    }
                }
            }
            _ = fi57
        }
    case -2147483648:
        for fi92 := 0; fi92 < 17; fi92++ {
            var v93 uint64 = 863
            if true {
                var v94 float64 = (-2.25 + float64(0.0))
                var v95 int8 = 127
                v96 := GenS9{F0: fi92, F1: v93}
                _ = v94
                _ = v95
                _ = v96
            } else {
                var v97 uint64 = 18446744073709551615
                var v98 string = "\n\t"
                _ = v97
                _ = v98
            }
            _ = fi92
            _ = v93
        }
    default:
        var v99 uint8 = (v2 & v2)
        _ = v99
    }
    v100 := []int{336, 477, -331, 1}
    switch ("\n\t" + ((("" + "\n\t") + "\n\t") + "\n\t")) {
    case "\n\t":
        switch "hello" {
        case "foo bar":
            v1 = 42
        case "a":
            fmt.Println(v1, v3)
        case "hello":
            v48 -= v48
        }
    default:
        for fi101 := 0; fi101 < 10; fi101++ {
            for fi102 := 0; fi102 < 1; fi102++ {
                if (1 > 0) {
                    var v103 uint32 = 2147483647
                    v100 = append(v100, fi102)
                    fmt.Println(fi101)
                    _ = v103
                } else {
                    if true {
                        fmt.Println(v2)
                        var v104 float32 = float32(3.14159)
                        if (0 == -9223372036854775808) {
                            var v105 uint64 = 0
                            _ = v105
                        }
                        var v106 float32 = v104
                        if false {
                            switch (fi102 / (((fi102 * (fi101 | (fi102 % ((fi101) | 1))))) | 1)) {
                            case -361:
                                if len(v100) >= 2 {
                                    v100 = v100[1:len(v100)]
                                }
                                v1 -= -9223372036854775808
                                var v107 int64 = -1
                                _ = v107
                            default:
                                fmt.Println(v48)
                            }
                            fmt.Println(v106)
                        } else {
                            v3 -= v45
                            var v108 bool = ((1.5 != (v104 / v104)) && (!(16 <= -64)))
                            _ = v108
                        }
                        var v109 uint32 = 843
                        if len(v100) > 0 {
                            v100[0] = (fi102 % (((fi101 / (((fi101 ^ (-(fi102)))) | 1))) | 1))
                        }
                        _ = v104
                        _ = v106
                        _ = v109
                    } else {
                        v1 += v1
                    }
                    if true {
                        var v110 uint64 = 2
                        fmt.Println(v110, v48)
                        _ = v110
                    }
                    var v111 float64 = 1e10
                    v1 = (v1 >> ((v48) & 63))
                    _ = v111
                }
                v1 += v1
                var v112 uint64 = 75
                fmt.Println(v112)
                var v113 int8 = 85
                _ = len(v100)
                v113 = v113
                _ = fi102
                _ = v112
                _ = v113
            }
            var v114 int8 = 100
            _ = fi101
            _ = v114
        }
    }
    for fi115 := 0; fi115 < 2; fi115++ {
        v3 = (v45 % ((v3) | 1))
        _ = fi115
    }
    var v116 float32 = float32(-2.25)
    var v117 bool = false
    for fi118 := 0; fi118 < 9; fi118++ {
        v100 = append(v100, (^(fi118)))
        if len(v100) >= 2 {
            v100 = v100[1:len(v100)]
        }
        _ = fi118
    }
    v45 = -1
    _ = len(v100)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v100:", len(v100), v100)
    fmt.Println("v116:", v116)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v100
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v45
    _ = v48
}
