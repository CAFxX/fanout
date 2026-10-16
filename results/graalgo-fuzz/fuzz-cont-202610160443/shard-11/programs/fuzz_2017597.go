package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS2 struct { F0 bool; F1 bool; F2 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 uint; F1 uint; F2 uint32 }
type GenS6 struct { F0 string; F1 int32; F2 int16 }
type GenS7 struct { F0 int32; F1 int16; F2 int32 }
type GenS8 struct { F0 string; F1 uint16 }
type GenS9 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int32; F1 uint64 }
type GenS12 struct { F0 uint64; F1 int64; F2 int8 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 int32; F1 int16; F2 string }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 uint; F1 int; F2 uint16 }
type GenS17 struct { F0 uint32; F1 uint8 }
type GenS18 struct { F0 int64 }
func genFn0(p0 int32) int8 {
    var v7 int = 812
    v7 += (v7 + 0)
    var v8 int = 0
    if false {
        fmt.Println(v8)
        var v9 int32 = 2147483647
        _ = v9
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = v7
    _ = v8
    return -128
}

func main() {
    var v1 int = 42
    var v2 string = "\n\t"
    var v3 int32 = 2147483647
    v4 := []string{"hello"}
    v5 := make(map[uint32]bool)
    var v6 uint8 = 251
    if true {
        if false {
            var v10 int16 = 42
            v6 += (^(v6))
            _ = v10
        } else {
            switch (v2 + v2) {
            case "\n\t":
                v5[13] = (true && ((3.14159 / 3.14159) >= (3.14159 - (float64(-2.25) * 0.1))))
            }
            if false {
                fmt.Println(v6, v3)
            }
        }
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 float64 = float64(0.0)
            _ = fi11
            _ = v12
        }
    }
    if true {
        var v13 int32 = 855
        _ = v13
    }
    if true {
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS1{F0: (^(v6)), F1: 0, F2: 1}
            var v16 bool = (1e10 == (1.5 / (1.5 / -2.25)))
            _ = fi14
            _ = v15
            _ = v16
        }
        fmt.Println(v6)
    } else {
        v3 = (v3 - (v3 - (v3 + (v3 + v3))))
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 int32 = (v3 * v3)
            if (0 >= 0) {
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                if (65535 <= 4294967295) {
                    var v19 int64 = 42
                    if (1 < 0) {
                        var v20 uint = 42
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 4; fi21++ {
                            switch v2 {
                            case "hello":
                                switch ("" + (v2 + "a")) {
                                case "foo bar":
                                    v22 := GenS2{F0: (float32((((1e10 * 0.1) * float64(0.0)) * 1.5)) != 0.1), F1: false, F2: 0}
                                    _ = v22
                                default:
                                    v23 := GenS1{F0: 42, F1: 255, F2: 2}
                                    _ = v23
                                }
                            case "":
                                var v24 int64 = 1
                                _ = v24
                            }
                            _ = fi21
                        }
                        var v25 bool = (((((3.14159 * 1e10) + 1e10) * float64(1e10)) + (((3.14159 / 1.0) - (0.0 + (1e10 - (0.0 - ((0.0 - 1.5) * ((-2.25 / 1.0) - 0.1)))))) / float64(1.5))) <= -2.25)
                        switch (^(v1)) {
                        case 712:
                            if (515 < 42) {
                                fmt.Println(v1, fi17, fi17)
                                var v26 float32 = 3.14159
                                switch (fi17 - (fi17 ^ v1)) {
                                case 2147483647:
                                    if v25 {
                                        v27 := GenS3{F0: v25}
                                        _ = v27
                                    } else {
                                        _ = len(v5)
                                        for fi28 := 0; fi28 < 15; fi28++ {
                                            var v29 uint16 = 344
                                            fmt.Println(v19, v1)
                                            v26 += (3.14159 - ((v26 / float32(0.1)) / v26))
                                            _ = fi28
                                            _ = v29
                                        }
                                    }
                                default:
                                    var v30 bool = (true && ((!false) && ((1.5 - 1.5) > float64((1e10 + 0.0)))))
                                    _ = v30
                                }
                                _ = v26
                            }
                        case 4294967295:
                            v2 = v2
                        }
                        _ = v25
                    }
                    for fi31 := 0; fi31 < 15; fi31++ {
                        fmt.Println(v3, v19, v19)
                        var v32 float64 = (float64(float32((1e10 + 1e10))) + float64((1e10 - 0.1)))
                        _ = fi31
                        _ = v32
                    }
                    _ = v19
                } else {
                    for fi33 := 0; fi33 < 10; fi33++ {
                        fmt.Println(v18)
                        _ = len(v4)
                        _ = fi33
                    }
                    var v34 int = (^(v1))
                    _ = v34
                }
                v6 += (^(v6))
            }
            _ = fi17
            _ = v18
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 = v2
    if (true && false) {
        if (65535 <= 1) {
            _ = len(v5)
            v6 = v6
            if true {
                var v35 bool = true
                _ = v35
            } else {
                v6 = v6
                switch (v1 >> ((642) & 63)) {
                case 4294967295:
                    var v36 int16 = 1
                    _ = v36
                case -97:
                    var v37 uint16 = 619
                    _ = v37
                case -1:
                    var v38 uint16 = 473
                    _ = v38
                }
                v2 = v2
            }
        } else {
            var v39 int = (v1 - 4294967295)
            var v40 uint16 = 2
            switch v2 {
            case "foo bar":
                if len(v4) > 0 {
                    v4[0] = v2
                }
            case "\n\t":
                fmt.Println(v1, v1)
            default:
                if (false || false) {
                    if (1 > genFn0(-939)) {
                        if false {
                            v3 = v3
                            if (!false) {
                                if (13 != 47) {
                                    if (!(!true)) {
                                        if (true && true) {
                                            switch (v2 + "hello") {
                                            case "a":
                                                switch (v2 + v2) {
                                                case "a":
                                                    if false {
                                                        v41 := GenS1{F0: 255, F1: 425, F2: v40}
                                                        v39 += v1
                                                        _ = v41
                                                    } else {
                                                        v4 = append(v4, (v2 + ((v2 + "\n\t") + ("foo bar" + (v2 + v2)))))
                                                        _ = len(v5)
                                                        if false {
                                                            _ = len(v4)
                                                            var v42 string = v2
                                                            _ = v42
                                                        }
                                                    }
                                                case "hello":
                                                    var v43 int = v39
                                                    _ = v43
                                                case "\n\t":
                                                    v3 = v3
                                                default:
                                                    if false {
                                                        v3 = (^(v3))
                                                    } else {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        var v44 int8 = -106
                                                        _ = v44
                                                    }
                                                }
                                            case "\n\t":
                                                v45 := GenS1{F0: v6, F1: 99, F2: (v40 / (((^(v40))) | 1))}
                                                _ = v45
                                            case "":
                                                v39 = (v39 + v1)
                                            default:
                                                for fi46 := 0; fi46 < 11; fi46++ {
                                                    v5[0] = true
                                                    var v47 uint32 = 65535
                                                    fmt.Println(fi46)
                                                    _ = fi46
                                                    _ = v47
                                                }
                                            }
                                            v6 = 42
                                        } else {
                                            v48 := GenS1{F0: v6, F1: 255, F2: v40}
                                            if true {
                                                for fi49 := 0; fi49 < 2; fi49++ {
                                                    for fi50 := 0; fi50 < 11; fi50++ {
                                                        var v51 int32 = 1
                                                        v39 -= -9223372036854775808
                                                        v52 := GenS4{F0: -115, F1: (v3 % (((v3 + v3)) | 1))}
                                                        v2 = v2
                                                        v53 := GenS3{F0: true}
                                                        switch ("hello" + v2) {
                                                        case "foo bar":
                                                            v54 := GenS3{F0: (true && true)}
                                                            _ = v54
                                                        case "hello":
                                                            v6 = (v6 ^ v6)
                                                        case "\n\t":
                                                            v5[42] = false
                                                        default:
                                                            v55 := GenS3{F0: true}
                                                            _ = v55
                                                        }
                                                        switch (fi49 << ((2) & 63)) {
                                                        case 141:
                                                            var v56 bool = (1e10 < 0.1)
                                                            fmt.Println(v3)
                                                            var v57 int32 = (v51 << ((715) & 63))
                                                            _ = v56
                                                            _ = v57
                                                        case 127:
                                                            fmt.Println(v3, fi49, v2)
                                                        case 856:
                                                            var v58 int64 = -753
                                                            _ = v58
                                                        }
                                                        _ = fi50
                                                        _ = v51
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                    if (("\n\t" + ((v2 + ("foo bar" + "\n\t")) + v2)) <= v2) {
                                                        v1 = v1
                                                        v59 := GenS5{F0: 818, F1: 65535, F2: 378}
                                                        v1 = 127
                                                        _ = v59
                                                    }
                                                    _ = fi49
                                                }
                                                fmt.Println(v6)
                                            }
                                            _ = v48
                                        }
                                    } else {
                                        for fi60 := 0; fi60 < 13; fi60++ {
                                            v61 := GenS6{F0: v2, F1: 1, F2: 1000}
                                            for fi62 := 0; fi62 < 3; fi62++ {
                                                fmt.Println(v39, v40, v39)
                                                _ = fi62
                                            }
                                            _ = fi60
                                            _ = v61
                                        }
                                    }
                                    if (729 > 1) {
                                        if false {
                                            var v63 uint16 = 294
                                            v64 := GenS3{F0: (genFn0(-207) == -103)}
                                            for fi65 := 0; fi65 < 19; fi65++ {
                                                v39 += fi65
                                                for fi66 := 0; fi66 < 5; fi66++ {
                                                    v4 = append(v4, (v2 + ((v2 + "") + v2)))
                                                    v67 := GenS1{F0: v6, F1: 217, F2: (v63 << ((42) & 63))}
                                                    _ = fi66
                                                    _ = v67
                                                }
                                                _ = fi65
                                            }
                                            _ = v63
                                            _ = v64
                                        } else {
                                            v40 -= (v40 + (v40 - v40))
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        var v68 uint8 = v6
                                        _ = v68
                                    } else {
                                        v69 := GenS4{F0: 42, F1: v3}
                                        var v70 int16 = 42
                                        _ = v69
                                        _ = v70
                                    }
                                }
                            } else {
                                var v71 uint64 = 4294967295
                                _ = v71
                            }
                        }
                        if false {
                            v1 -= (v39 << ((18446744073709551615) & 63))
                            v39 -= 986
                            var v72 uint32 = 758
                            _ = v72
                        } else {
                            v4 = append(v4, "foo bar")
                            var v73 int32 = (^(v3))
                            if false {
                                if false {
                                    for fi74 := 0; fi74 < 10; fi74++ {
                                        if len(v4) > 0 {
                                            v4[0] = "hello"
                                        }
                                        v1 = (v1 & (v1 ^ v39))
                                        if (3.14159 < (0.0 - 0.0)) {
                                            for fi75 := 0; fi75 < 3; fi75++ {
                                                var v76 int32 = int32(v6)
                                                var v77 uint32 = 42
                                                var v78 int16 = -32768
                                                _ = fi75
                                                _ = v76
                                                _ = v77
                                                _ = v78
                                            }
                                            switch (^(fi74)) {
                                            case 32767:
                                                var v79 float32 = ((3.14159 * 1e10) * -2.25)
                                                _ = v79
                                            case 0:
                                                _ = len(v4)
                                            case 32768:
                                                for fi80 := 0; fi80 < 16; fi80++ {
                                                    fmt.Println(v6, v3)
                                                    _ = fi80
                                                }
                                            default:
                                                for fi81 := 0; fi81 < 16; fi81++ {
                                                    var v82 float32 = float32((((3.14159 / 0.1) * -2.25) * (1.5 * 0.1)))
                                                    v3 -= v73
                                                    _ = fi81
                                                    _ = v82
                                                }
                                            }
                                            fmt.Println(fi74, v2)
                                        }
                                        _ = fi74
                                    }
                                    if (true && ((1 != 476) && false)) {
                                        var v83 uint32 = 2147483647
                                        _ = v83
                                    } else {
                                        if false {
                                            var v84 uint8 = (v6 % ((v6) | 1))
                                            _ = v84
                                        }
                                        v39 = v39
                                        switch v2 {
                                        case "":
                                            if true {
                                                _ = len(v5)
                                                v85 := GenS7{F0: v73, F1: 0, F2: (v73 * v3)}
                                                _ = v85
                                            } else {
                                                v6 -= v6
                                                var v86 string = "foo bar"
                                                var v87 uint16 = v40
                                                _ = v86
                                                _ = v87
                                            }
                                        }
                                    }
                                } else {
                                    for fi88 := 0; fi88 < 17; fi88++ {
                                        v89 := GenS8{F0: ("\n\t" + ""), F1: (^(v40))}
                                        _ = fi88
                                        _ = v89
                                    }
                                }
                            } else {
                                for fi90 := 0; fi90 < 14; fi90++ {
                                    var v91 uint8 = (v6 << ((93) & 63))
                                    _ = fi90
                                    _ = v91
                                }
                                delete(v5, 42)
                            }
                            _ = v73
                        }
                        var v92 int16 = -939
                        _ = v92
                    } else {
                        var v93 uint = 65535
                        if (((("a" + (("a" + "hello") + v2)) + (v2 + v2)) <= v2) && (576 >= 301)) {
                            switch (v39 + (v1 >> (((v93 + v93)) & 63))) {
                            case 4294967295:
                                var v94 bool = false
                                _ = v94
                            case 32767:
                                var v95 uint8 = v6
                                _ = v95
                            case -504:
                                var v96 string = (((v2 + v2) + (v2 + v2)) + v2)
                                _ = v96
                            default:
                                v6 -= v6
                            }
                            if (0.1 < 3.14159) {
                                for fi97 := 0; fi97 < 3; fi97++ {
                                    var v98 int16 = -301
                                    _ = fi97
                                    _ = v98
                                }
                            } else {
                                if true {
                                    switch int(v40) {
                                    case -159:
                                        v99 := GenS5{F0: 0, F1: v93, F2: 4294967295}
                                        _ = v99
                                    case 2147483647:
                                        fmt.Println(v2, v2)
                                    default:
                                        if (!false) {
                                            if (!true) {
                                                if (!(false && true)) {
                                                    v40 = (v40 & v40)
                                                    for fi100 := 0; fi100 < 17; fi100++ {
                                                        if len(v4) > 0 {
                                                            v4[0] = v2
                                                        }
                                                        _ = fi100
                                                    }
                                                }
                                                if (v93 > v93) {
                                                    for fi101 := 0; fi101 < 12; fi101++ {
                                                        fmt.Println(v39)
                                                        _ = fi101
                                                    }
                                                    v2 = ((v2 + (("\n\t" + "") + "")) + ("foo bar" + ("foo bar" + "hello")))
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                } else {
                                                    switch (v1 * (v39 % (((v1 >> ((v93) & 63))) | 1))) {
                                                    case 128:
                                                        if ((1 != 631) && (false && true)) {
                                                            var v102 bool = (-32768 < -1)
                                                            switch 127 {
                                                            case 1:
                                                                var v103 bool = (v40 == v40)
                                                                _ = v103
                                                            }
                                                            _ = v102
                                                        } else {
                                                            switch int(v6) {
                                                            case -704:
                                                                if false {
                                                                    v6 = v6
                                                                } else {
                                                                    var v104 int32 = (^(v3))
                                                                    if true {
                                                                        fmt.Println(v3)
                                                                        fmt.Println(v6)
                                                                        fmt.Println(v6)
                                                                    } else {
                                                                        var v105 bool = (0 >= 965)
                                                                        var v106 int32 = v104
                                                                        switch (v1 - -2147483648) {
                                                                        case -1:
                                                                            var v107 int = 128
                                                                            _ = v107
                                                                        default:
                                                                            var v108 int32 = (v106 >> ((v93) & 63))
                                                                            _ = v108
                                                                        }
                                                                        _ = v105
                                                                        _ = v106
                                                                    }
                                                                    _ = v104
                                                                }
                                                            case -413:
                                                                for fi109 := 0; fi109 < 20; fi109++ {
                                                                    for fi110 := 0; fi110 < 12; fi110++ {
                                                                        switch v2 {
                                                                        case "a":
                                                                            var v111 int16 = 0
                                                                            _ = v111
                                                                        case "hello":
                                                                            fmt.Println(v93, v6, v3)
                                                                        default:
                                                                            v40 -= 2
                                                                        }
                                                                        _ = fi110
                                                                    }
                                                                    fmt.Println(v2, v93, v39)
                                                                    _ = fi109
                                                                }
                                                            case -1:
                                                                for fi112 := 0; fi112 < 6; fi112++ {
                                                                    v6 += v6
                                                                    switch v39 {
                                                                    case 127:
                                                                        v113 := GenS9{F0: 650, F1: (v6 / ((v6) | 1)), F2: -128}
                                                                        _ = v113
                                                                    case 4294967295:
                                                                        for fi114 := 0; fi114 < 4; fi114++ {
                                                                            fmt.Println(v93, v6)
                                                                            var v115 uint8 = v6
                                                                            _ = fi114
                                                                            _ = v115
                                                                        }
                                                                    case 221:
                                                                        var v116 uint32 = 287
                                                                        _ = v116
                                                                    }
                                                                    _ = fi112
                                                                }
                                                            default:
                                                                _ = len(v5)
                                                            }
                                                            v40 = 42
                                                            if (0.0 == -2.25) {
                                                                var v117 int8 = genFn0((v3 >> (((v93 - (v93 ^ uint(v40)))) & 63)))
                                                                v118 := GenS10{F0: 42}
                                                                var v119 float64 = 3.14159
                                                                _ = v117
                                                                _ = v118
                                                                _ = v119
                                                            }
                                                        }
                                                    default:
                                                        v5[2] = true
                                                    }
                                                    v5[2] = (!(false || true))
                                                }
                                            }
                                            fmt.Println(v6, v39, v39)
                                        } else {
                                            for fi120 := 0; fi120 < 9; fi120++ {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v3 -= (v3 * (v3 / ((v3) | 1)))
                                                case "":
                                                    v40 = uint16(v93)
                                                case "a":
                                                    v4 = append(v4, ("" + v2))
                                                default:
                                                    if true {
                                                        v1 = v1
                                                    } else {
                                                        v39 += v39
                                                        v121 := v5[65535]
                                                        _ = v121
                                                    }
                                                }
                                                for fi122 := 0; fi122 < 7; fi122++ {
                                                    var v123 int = (fi122 + (-(v1)))
                                                    _ = fi122
                                                    _ = v123
                                                }
                                                if (uint32(v93) == 1) {
                                                    v6 -= v6
                                                    v124 := GenS11{F0: v3, F1: 996}
                                                    _ = v124
                                                }
                                                _ = fi120
                                            }
                                            v93 = (v93 ^ (v93 << ((4294967295) & 63)))
                                            v2 = "\n\t"
                                        }
                                    }
                                } else {
                                    for fi125 := 0; fi125 < 9; fi125++ {
                                        _ = len(v5)
                                        fmt.Println(v2)
                                        _ = fi125
                                    }
                                    var v126 float32 = (1e10 - (0.0 - (float32((0.1 + (1e10 + float64(0.1)))) - -2.25)))
                                    _ = v126
                                }
                                if false {
                                    var v127 int64 = 1
                                    _ = v127
                                }
                                v1 = (v39 >> ((v93) & 63))
                            }
                        } else {
                            var v128 uint32 = 0
                            var v129 uint16 = v40
                            _ = v128
                            _ = v129
                        }
                        var v130 string = ("foo bar" + v2)
                        _ = v130
                        _ = v93
                    }
                    if len(v4) > 0 {
                        v4[0] = (v2 + v2)
                    }
                    v2 = (v2 + ("\n\t" + "hello"))
                } else {
                    v2 = v2
                    fmt.Println(v40, v2, v2)
                }
            }
            _ = v39
            _ = v40
        }
        v4 = append(v4, v2)
    } else {
        for fi131 := 0; fi131 < 7; fi131++ {
            fmt.Println(v6, v1, fi131)
            _ = fi131
        }
        for fi132 := 0; fi132 < 9; fi132++ {
            for fi133 := 0; fi133 < 19; fi133++ {
                var v134 uint32 = 782
                var v135 int = fi133
                _ = fi133
                _ = v134
                _ = v135
            }
            _ = fi132
        }
        if ((106 == 350) || true) {
            var v136 uint32 = 16
            v137 := GenS12{F0: 537, F1: -1, F2: 42}
            _ = v136
            _ = v137
        } else {
            for fi138 := 0; fi138 < 20; fi138++ {
                v1 = 1
                var v139 uint8 = 0
                _ = fi138
                _ = v139
            }
            fmt.Println(v2, v3)
            for fi140 := 0; fi140 < 9; fi140++ {
                for fi141 := 0; fi141 < 12; fi141++ {
                    if (!false) {
                        _ = len(v5)
                        var v142 int16 = 1
                        if false {
                            v6 = 119
                            if false {
                                fmt.Println(v6, fi140)
                                if ((416 >= 32767) && false) {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    switch v2 {
                                    case "hello":
                                        if true {
                                            _ = len(v4)
                                            var v143 string = "\n\t"
                                            if true {
                                                fmt.Println(v6, v143)
                                                switch ((v2 + "hello") + "\n\t") {
                                                case "a":
                                                    if false {
                                                        var v144 uint = 840
                                                        v145 := GenS5{F0: 994, F1: v144, F2: 2147483648}
                                                        _ = v144
                                                        _ = v145
                                                    }
                                                default:
                                                    delete(v5, 4294967295)
                                                }
                                            } else {
                                                v2 = "\n\t"
                                                fmt.Println(v143, v142, v2)
                                            }
                                            _ = v143
                                        } else {
                                            v146 := GenS13{F0: v142}
                                            v147 := v5[65535]
                                            fmt.Println(v1, v2, v142)
                                            switch v2 {
                                            case "foo bar":
                                                if false {
                                                    if false {
                                                        fmt.Println(v142, v1, v2)
                                                    }
                                                    var v148 uint8 = (v6 >> ((111) & 63))
                                                    v5[2] = ((42 >= 0) || true)
                                                    _ = v148
                                                }
                                            }
                                            _ = v146
                                            _ = v147
                                        }
                                    default:
                                        var v149 int = (^(v1))
                                        _ = v149
                                    }
                                } else {
                                    fmt.Println(v1)
                                    if (((v2 + ((("foo bar" + v2) + v2) + (v2 + v2))) + "") > v2) {
                                        var v150 float32 = ((((float32(((float64(0.1) - 0.0) + -2.25)) * 1.5) / 1e10) + float32(0.1)) / -2.25)
                                        var v151 int8 = 0
                                        if (4294967295 <= 65535) {
                                            var v152 uint = 65535
                                            if (v152 >= 18446744073709551615) {
                                                if (!(!((v151 ^ genFn0(int32(v6))) == v151))) {
                                                    v153 := GenS14{F0: v3, F1: v142, F2: v2}
                                                    _ = v153
                                                }
                                            }
                                            _ = v152
                                        } else {
                                            v1 += (-(fi141))
                                            var v154 uint32 = 2147483647
                                            _ = v154
                                        }
                                        _ = v150
                                        _ = v151
                                    } else {
                                        _ = len(v4)
                                        delete(v5, 65535)
                                        v155 := GenS15{F0: false}
                                        _ = v155
                                    }
                                    var v156 int8 = 0
                                    _ = v156
                                }
                                var v157 int16 = v142
                                v158 := GenS13{F0: 1000}
                                var v159 int16 = v157
                                v1 = -9223372036854775808
                                v157 = 42
                                _ = v157
                                _ = v158
                                _ = v159
                            } else {
                                switch (fi141 >> ((1) & 63)) {
                                case 0:
                                    var v160 int32 = (v3 * 42)
                                    _ = v160
                                case -9223372036854775808:
                                    fmt.Println(v1, v3, v2)
                                case 389:
                                    if ((!(v2 <= v2)) && true) {
                                        var v161 int32 = (v3 << ((255) & 63))
                                        if len(v4) > 0 {
                                            v4[0] = "foo bar"
                                        }
                                        if (v161 == -75) {
                                            switch (v2 + (v2 + "hello")) {
                                            case "foo bar":
                                                var v162 float64 = 1.5
                                                _ = v162
                                            }
                                        }
                                        _ = v161
                                    }
                                default:
                                    fmt.Println(v6, v2, fi140)
                                }
                                if (42 < 4294967295) {
                                    fmt.Println(v3)
                                    v163 := GenS16{F0: 255, F1: 9223372036854775807, F2: 0}
                                    _ = v163
                                }
                                v1 += fi141
                                v6 -= (^(v6))
                            }
                            var v164 int64 = -527
                            _ = v164
                        }
                        _ = v142
                    }
                    var v165 uint8 = (v6 >> ((29) & 63))
                    v1 -= 32767
                    _ = fi141
                    _ = v165
                }
                if (0 <= 18446744073709551615) {
                    if true {
                        if true {
                            switch (v2 + "hello") {
                            case "hello":
                                var v166 int32 = v3
                                _ = v166
                            default:
                                v167 := GenS17{F0: 786, F1: 171}
                                _ = v167
                            }
                        }
                    }
                }
                _ = fi140
            }
        }
    }
    var v168 uint8 = (v6 + (v6 - (^(v6))))
    var v169 int8 = -128
    v2 = ("a" + (v2 + (v2 + ("foo bar" + ((v2 + v2) + (v2 + v2))))))
    if false {
        var v170 uint8 = 128
        switch (^(v1)) {
        case -1:
            for fi171 := 0; fi171 < 5; fi171++ {
                var v172 int8 = (v169 + 1)
                v6 += (v170 * v170)
                _ = fi171
                _ = v172
            }
        case 4294967295:
            var v173 int32 = (v3 >> ((18446744073709551615) & 63))
            _ = v173
        }
        _ = v170
    } else {
        for fi174 := 0; fi174 < 18; fi174++ {
            v175 := GenS18{F0: -415}
            if (!(v2 > "a")) {
                v2 = "a"
            } else {
                for fi176 := 0; fi176 < 1; fi176++ {
                    v1 -= fi174
                    var v177 int8 = -1
                    _ = fi176
                    _ = v177
                }
            }
            _ = fi174
            _ = v175
        }
    }
    if ("" == v2) {
        v178 := GenS3{F0: false}
        _ = v178
    }
    fmt.Println(v169, v169, v1)
    var v179 float32 = 3.14159
    fmt.Println(v2, v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v168:", v168)
    fmt.Println("v169:", v169)
    fmt.Println("v179:", v179)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v168
    _ = v169
    _ = v179
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
