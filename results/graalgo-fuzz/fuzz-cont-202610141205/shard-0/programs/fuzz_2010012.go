package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int16; F2 uint8 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int8; F1 bool; F2 uint8 }
type GenS5 struct { F0 bool; F1 int8; F2 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 bool; F2 int16 }
type GenS8 struct { F0 uint64; F1 int16; F2 int64 }
type GenS9 struct { F0 uint64; F1 bool }
type GenS10 struct { F0 uint32; F1 int8 }
type GenS11 struct { F0 uint8; F1 int32 }
func genFn0() int32 {
    for fi5 := 0; fi5 < 11; fi5++ {
        if true {
            var v6 int16 = 1
            _ = v6
        }
        v7 := GenS1{F0: (!true)}
        var v8 uint64 = 477
        _ = fi5
        _ = v7
        _ = v8
    }
    var v9 int = -9223372036854775808
    for fi10 := 0; fi10 < 13; fi10++ {
        switch "a" {
        case "\n\t":
            var v11 int16 = -32768
            _ = v11
        case "hello":
            var v12 float32 = 3.14159
            _ = v12
        case "foo bar":
            v13 := GenS1{F0: (2 < 127)}
            _ = v13
        default:
            v9 = fi10
        }
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := []int{}
            v9 = v9
            _ = fi14
            _ = v15
        }
        _ = fi10
    }
    if false {
        v16 := GenS2{F0: 42, F1: 0, F2: 255}
        v17 := GenS1{F0: true}
        for fi18 := 0; fi18 < 9; fi18++ {
            v19 := GenS3{F0: -1, F1: 0}
            if ((922 <= 2) || (!(!true))) {
                switch -2147483648 {
                case -62:
                    if ((-1 != 963) && (1 >= 227)) {
                        v9 -= (v9 * (fi18 % ((fi18) | 1)))
                    } else {
                        for fi20 := 0; fi20 < 8; fi20++ {
                            if (!((0 <= 150) && true)) {
                                v9 = 127
                                var v21 uint32 = 4294967295
                                if true {
                                    v22 := GenS4{F0: int8(fi20), F1: (!true), F2: 1}
                                    var v23 uint8 = 39
                                    _ = v22
                                    _ = v23
                                }
                                _ = v21
                            }
                            switch 128 {
                            case 0:
                                v9 = fi18
                            case 127:
                                v24 := GenS3{F0: 702, F1: -1}
                                _ = v24
                            case 1:
                                var v25 uint64 = 4294967295
                                _ = v25
                            default:
                                if (!false) {
                                    var v26 int16 = 1
                                    _ = v26
                                }
                            }
                            switch (("foo bar" + "") + "foo bar") {
                            case "hello":
                                v9 = fi18
                            }
                            v27 := GenS5{F0: (false || (252 >= 4294967295)), F1: 59, F2: "hello"}
                            _ = fi20
                            _ = v27
                        }
                        fmt.Println(fi18, fi18)
                        fmt.Println(v9)
                    }
                case -9223372036854775808:
                    for fi28 := 0; fi28 < 14; fi28++ {
                        if true {
                            v29 := []int{}
                            if (0 < 89) {
                                v30 := GenS5{F0: true, F1: -31, F2: "a"}
                                switch "" {
                                case "foo bar":
                                    v9 += (-(v9))
                                    var v31 int8 = 42
                                    v31 = v31
                                    _ = v31
                                case "\n\t":
                                    if (!false) {
                                        switch "hello" {
                                        case "a":
                                            v32 := GenS6{F0: 42}
                                            _ = v32
                                        default:
                                            if (42 != 18446744073709551615) {
                                                var v33 uint32 = 42
                                                v33 = 2147483648
                                                var v34 uint16 = uint16(v33)
                                                _ = v33
                                                _ = v34
                                            }
                                        }
                                        var v35 float64 = 1e10
                                        var v36 uint16 = 0
                                        v37 := GenS1{F0: (-1 >= 1)}
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    } else {
                                        _ = len(v29)
                                        var v38 uint = 877
                                        _ = v38
                                    }
                                case "hello":
                                    v9 = (-(fi28))
                                }
                                fmt.Println(v9)
                                _ = v30
                            } else {
                                v9 = -9223372036854775808
                                var v39 uint = 923
                                if (1e10 >= (3.14159 - 1.5)) {
                                    switch ("hello" + (("" + "\n\t") + "\n\t")) {
                                    case "":
                                        v39 += v39
                                    case "a":
                                        if false {
                                            var v40 float64 = float64(0.1)
                                            _ = v40
                                        }
                                    case "foo bar":
                                        if true {
                                            if len(v29) >= 2 {
                                                v29 = v29[1:len(v29)]
                                            }
                                        }
                                    }
                                    v41 := GenS7{F0: 128, F1: false, F2: -306}
                                    if (true || (0 <= 0)) {
                                        switch fi28 {
                                        case 1:
                                            v29 = append(v29, (v9 % ((fi28) | 1)))
                                        }
                                        switch "\n\t" {
                                        case "hello":
                                            v9 -= v9
                                        case "\n\t":
                                            var v42 uint = v39
                                            _ = v42
                                        }
                                    } else {
                                        var v43 uint8 = 128
                                        v43 = (^(v43))
                                        _ = v43
                                    }
                                    _ = v41
                                }
                                _ = v39
                            }
                            v9 += (fi18 | v9)
                            var v44 uint64 = 614
                            v45 := GenS6{F0: 0}
                            _ = v29
                            _ = v44
                            _ = v45
                        }
                        _ = fi28
                    }
                case 424:
                    var v46 int8 = -1
                    _ = v46
                }
                if (-586 < 995) {
                    var v47 float32 = 0.0
                    _ = v47
                } else {
                    var v48 uint32 = 2147483648
                    _ = v48
                }
                for fi49 := 0; fi49 < 15; fi49++ {
                    v9 -= fi49
                    if (42 >= 65535) {
                        fmt.Println(v9, fi18)
                        fmt.Println(v9)
                    }
                    v50 := []int{fi18, fi18, (fi18 - (-(fi49)))}
                    _ = fi49
                    _ = v50
                }
            }
            v9 += (v9 ^ v9)
            _ = fi18
            _ = v19
        }
        _ = v16
        _ = v17
    }
    _ = v9
    return -506
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int = (v1 * (v1 * v1))
    v3 := []int{(^(v1)), v1}
    var v4 string = "foo bar"
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = ("hello" + v4)
    var v51 int64 = -463
    if false {
        if false {
            v51 += v51
            if (0.0 != (-2.25 + 3.14159)) {
                for fi52 := 0; fi52 < 6; fi52++ {
                    switch (v1 + 6) {
                    case 128:
                        var v53 float32 = (-2.25 / float32(float64(3.14159)))
                        _ = v53
                    case 0:
                        if true {
                            if true {
                                v2 -= (v1 | fi52)
                                v51 -= (v51 ^ (v51 << ((65535) & 63)))
                                var v54 int64 = 1
                                _ = v54
                            }
                            v2 = v2
                            v55 := GenS8{F0: 42, F1: 1000, F2: v51}
                            _ = v55
                        } else {
                            if len(v3) > 0 {
                                v3[0] = v2
                            }
                            if (420 < 42) {
                                v56 := GenS4{F0: -1, F1: (!true), F2: 19}
                                v1 -= v2
                                switch 156 {
                                case -1:
                                    var v57 int = (v1 >> ((255) & 63))
                                    _ = v57
                                }
                                _ = v56
                            } else {
                                fmt.Println(v4, v4)
                                if len(v3) > 0 {
                                    v3[0] = fi52
                                }
                            }
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                for fi58 := 0; fi58 < 6; fi58++ {
                                    v59 := GenS9{F0: 2, F1: (3.14159 <= (1e10 * (float32((1e10 / 3.14159)) + 0.1)))}
                                    var v60 string = "foo bar"
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                }
                                v51 = (v51 ^ (^(v51)))
                            }
                        }
                    case -9223372036854775808:
                        switch v4 {
                        case "hello":
                            v4 = "foo bar"
                        }
                    default:
                        var v61 float64 = 1e10
                        _ = v61
                    }
                    var v62 uint64 = 4294967295
                    fmt.Println(v4, v51, v62)
                    _ = fi52
                    _ = v62
                }
                v1 += v2
                var v63 int16 = -465
                _ = v63
            } else {
                var v64 float32 = ((0.1 / 3.14159) + -2.25)
                v3 = append(v3, v1)
                _ = v64
            }
            if true {
                if (!(!true)) {
                    switch -527 {
                    case 4294967295:
                        if (251 != 235) {
                            v65 := GenS2{F0: -637, F1: 0, F2: 76}
                            for fi66 := 0; fi66 < 4; fi66++ {
                                var v67 uint = 670
                                if ((-1 < 42) && true) {
                                    if true {
                                        for fi68 := 0; fi68 < 9; fi68++ {
                                            var v69 string = v4
                                            var v70 int32 = -2147483648
                                            _ = fi68
                                            _ = v69
                                            _ = v70
                                        }
                                        switch v4 {
                                        case "foo bar":
                                            for fi71 := 0; fi71 < 7; fi71++ {
                                                v67 += 16
                                                _ = fi71
                                            }
                                        case "":
                                            v1 = v2
                                        default:
                                            var v72 uint = (v67 >> (((v67 & (v67 >> (((^(v67))) & 63)))) & 63))
                                            _ = v72
                                        }
                                        fmt.Println(fi66, fi66)
                                    }
                                    v1 -= -154
                                    for fi73 := 0; fi73 < 17; fi73++ {
                                        var v74 bool = false
                                        v75 := GenS10{F0: 668, F1: 63}
                                        _ = fi73
                                        _ = v74
                                        _ = v75
                                    }
                                }
                                _ = fi66
                                _ = v67
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v65
                        } else {
                            var v76 int16 = 0
                            _ = v76
                        }
                    }
                    fmt.Println(v1, v1)
                } else {
                    fmt.Println(v1, v51)
                    var v77 bool = (float64(0.1) <= (3.14159 - 3.14159))
                    _ = v77
                }
            } else {
                if true {
                    var v78 float32 = (0.1 / 3.14159)
                    v4 = "a"
                    _ = v78
                } else {
                    var v79 int8 = -128
                    if false {
                        var v80 uint8 = 255
                        v81 := GenS9{F0: 186, F1: false}
                        _ = v80
                        _ = v81
                    }
                    v4 = v4
                    _ = v79
                }
            }
        }
        var v82 string = ("\n\t" + "a")
        if (32767 > (v2 ^ 1)) {
            v2 -= (v1 & v2)
            v51 = v51
            v51 = v51
        } else {
            switch v2 {
            case -2147483648:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 980:
                if (((((1e10 - (1.5 - (3.14159 / (1.5 / (((0.1 / 1.5) * (1e10 * ((1e10 * (((3.14159 / 1e10) - ((0.0 - float64((0.1 - 0.0))) - 0.0)) + 0.0)) + (1.5 / 1.0)))) - 3.14159))))) * (0.1 - 0.0)) / 1.0) / 0.1) == 0.1) {
                    v1 = (v2 & v1)
                    var v83 uint64 = 42
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v83
                } else {
                    v84 := GenS7{F0: 178, F1: true, F2: 89}
                    for fi85 := 0; fi85 < 8; fi85++ {
                        var v86 bool = (!true)
                        _ = fi85
                        _ = v86
                    }
                    _ = v84
                }
            }
        }
        _ = v82
    } else {
        fmt.Println(v2)
    }
    for fi87 := 0; fi87 < 18; fi87++ {
        v2 += (v2 >> ((0) & 63))
        v88 := GenS8{F0: 4294967295, F1: 32767, F2: (v51 + v51)}
        var v89 uint8 = 73
        _ = fi87
        _ = v88
        _ = v89
    }
    if len(v3) > 0 {
        v3[0] = (^(v2))
    }
    switch (v1 * v2) {
    case 127:
        if true {
            if false {
                v90 := GenS9{F0: 4294967295, F1: (18446744073709551615 == 65535)}
                _ = v90
            }
        }
    case -9223372036854775808:
        v1 -= v2
    default:
        v2 = 0
    }
    for fi91 := 0; fi91 < 16; fi91++ {
        if len(v3) > 0 {
            v3[0] = -803
        }
        switch v4 {
        case "":
            _ = len(v3)
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
        default:
            var v92 uint8 = 42
            _ = v92
        }
        _ = fi91
    }
    v51 += (v51 | v51)
    var v93 int32 = -2147483648
    switch v4 {
    case "foo bar":
        _ = len(v3)
    case "a":
        if false {
            v4 = ("a" + "a")
            if ((915 >= 42) || (-2147483648 <= -619)) {
                _ = len(v3)
                switch 32768 {
                case 2147483647:
                    fmt.Println(v51, v51, v93)
                case 9223372036854775807:
                    var v94 uint = 2
                    _ = v94
                default:
                    v93 = 0
                }
                switch v4 {
                case "hello":
                    v1 = (v1 ^ v2)
                case "a":
                    v51 -= 1
                }
            }
            fmt.Println(v2)
        }
    default:
        var v95 int64 = -1
        _ = v95
    }
    fmt.Println(v4, v51, v51)
    v96 := GenS11{F0: 1, F1: (v93 - v93)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v93
    _ = v96
}
