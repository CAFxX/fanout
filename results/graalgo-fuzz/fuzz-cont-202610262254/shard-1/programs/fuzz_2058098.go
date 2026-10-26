package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int }
type GenS2 struct { F0 int; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS9 struct { F0 int16; F1 int }
type GenS10 struct { F0 int32; F1 string }
type GenS11 struct { F0 int16; F1 int64 }
type GenS12 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int64; F1 int16 }
type GenS15 struct { F0 uint32; F1 int64; F2 int16 }
type GenS16 struct { F0 bool; F1 int64; F2 int32 }

func main() {
    v1 := []uint32{1}
    v2 := make(map[uint32]string)
    var v3 uint32 = 65535
    var v4 int16 = 289
    v5 := []uint32{}
    switch "\n\t" {
    case "\n\t":
        v3 = (v3 % ((2) | 1))
    case "foo bar":
        var v6 int32 = 1
        _ = v6
    case "a":
        var v7 int64 = 42
        _ = v7
    default:
        for fi8 := 0; fi8 < 15; fi8++ {
            if (32767 > 255) {
                _ = len(v2)
                if true {
                    for fi9 := 0; fi9 < 19; fi9++ {
                        var v10 string = ""
                        _ = fi9
                        _ = v10
                    }
                    if false {
                        for fi11 := 0; fi11 < 1; fi11++ {
                            v3 = (v3 * v3)
                            v12 := GenS1{F0: "hello", F1: fi11, F2: fi8}
                            _ = fi11
                            _ = v12
                        }
                    } else {
                        for fi13 := 0; fi13 < 20; fi13++ {
                            v3 += (v3 | (v3 * v3))
                            if (41 >= 127) {
                                var v14 uint = 65535
                                _ = v14
                            }
                            _ = fi13
                        }
                    }
                    v4 = (v4 ^ (v4 - (v4 & -1)))
                }
                for fi15 := 0; fi15 < 2; fi15++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    fmt.Println(v4, fi8)
                    _ = fi15
                }
            } else {
                var v16 uint32 = 2147483647
                v17 := GenS2{F0: fi8, F1: 65535, F2: 93}
                _ = v16
                _ = v17
            }
            for fi18 := 0; fi18 < 8; fi18++ {
                v19 := GenS3{F0: 65535}
                v20 := GenS4{F0: 1, F1: 42}
                _ = fi18
                _ = v19
                _ = v20
            }
            var v21 float32 = (0.1 + -2.25)
            _ = fi8
            _ = v21
        }
    }
    _ = len(v2)
    v4 -= -503
    var v22 uint32 = v3
    var v23 uint32 = (v22 ^ (^(v3)))
    v24 := GenS2{F0: 32767, F1: 601, F2: uint16(v22)}
    for fi25 := 0; fi25 < 9; fi25++ {
        if ((!true) && true) {
            var v26 uint16 = 255
            _ = v26
        } else {
            var v27 uint = 42
            _ = v27
        }
        v28 := GenS4{F0: 18446744073709551615, F1: 1}
        _ = fi25
        _ = v28
    }
    switch -9223372036854775808 {
    case 1:
        if true {
            switch -926 {
            case 42:
                if false {
                    v22 -= v3
                } else {
                    if true {
                        v29 := GenS5{F0: "foo bar"}
                        v3 = v3
                        for fi30 := 0; fi30 < 18; fi30++ {
                            if true {
                                switch "\n\t" {
                                case "hello":
                                    v2[4294967295] = "\n\t"
                                case "foo bar":
                                    v23 -= (v23 & (v22 << ((511) & 63)))
                                }
                                var v31 int8 = -1
                                for fi32 := 0; fi32 < 7; fi32++ {
                                    v3 = v3
                                    switch (fi32 >> ((862) & 63)) {
                                    case 1:
                                        v31 = (-(v31))
                                    case -9223372036854775808:
                                        var v33 float32 = -2.25
                                        _ = v33
                                    }
                                    _ = fi32
                                }
                                _ = v31
                            } else {
                                delete(v2, v23)
                                switch 9223372036854775807 {
                                case -4:
                                    if (-1 > 1) {
                                        var v34 int8 = 1
                                        var v35 uint8 = 236
                                        fmt.Println(v34, v23)
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v36 := v2[v22]
                                        v23 -= v23
                                        _ = v36
                                    }
                                default:
                                    switch "a" {
                                    case "hello":
                                        v4 += -606
                                    case "\n\t":
                                        var v37 int64 = -1
                                        _ = v37
                                    case "":
                                        v38 := GenS6{F0: v23}
                                        _ = v38
                                    default:
                                        var v39 int32 = -1
                                        _ = v39
                                    }
                                }
                            }
                            _ = fi30
                        }
                        _ = v29
                    }
                    v2[(v23 - 0)] = "a"
                    switch 2147483647 {
                    case 32768:
                        v40 := v2[(v3 | 4294967295)]
                        _ = v40
                    case 9223372036854775807:
                        if ((1 < 42) || (401 > 2)) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            for fi41 := 0; fi41 < 13; fi41++ {
                                for fi42 := 0; fi42 < 14; fi42++ {
                                    v43 := GenS7{F0: (v4 * (v4 / ((v4) | 1))), F1: -1}
                                    var v44 float32 = 1.5
                                    if ((54 >= -128) && (true || false)) {
                                        var v45 int = 128
                                        v46 := GenS8{F0: 2, F1: 127, F2: 563}
                                        _ = v45
                                        _ = v46
                                    } else {
                                        var v47 int64 = -283
                                        v22 = 0
                                        _ = v47
                                    }
                                    delete(v2, (v23 << ((42) & 63)))
                                    v2[(v23 * 0)] = "\n\t"
                                    if (!false) {
                                        var v48 bool = true
                                        _ = v48
                                    } else {
                                        switch "\n\t" {
                                        case "hello":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case "":
                                            v49 := GenS3{F0: 617}
                                            _ = v49
                                        case "a":
                                            var v50 float32 = (v44 / (0.1 - (float32(-2.25) / (v44 - 1e10))))
                                            _ = v50
                                        default:
                                            fmt.Println(v44, fi41)
                                        }
                                        v51 := GenS6{F0: v23}
                                        if true {
                                            v3 = (v22 * v22)
                                        }
                                        var v52 float32 = (3.14159 * (1.5 / (v44 + 0.0)))
                                        v5 = append(v5, v23)
                                        v2[(v3 | 2147483648)] = (("" + "") + ("hello" + ""))
                                        _ = v51
                                        _ = v52
                                    }
                                    _ = fi42
                                    _ = v43
                                    _ = v44
                                }
                                _ = fi41
                            }
                            if (false || true) {
                                for fi53 := 0; fi53 < 14; fi53++ {
                                    delete(v2, (v3 * (v22 ^ v23)))
                                    _ = fi53
                                }
                            } else {
                                v54 := GenS9{F0: v4, F1: -987}
                                _ = v54
                            }
                        }
                    default:
                        fmt.Println(v22, v3, v3)
                    }
                }
            case -863:
                fmt.Println(v23, v22)
            }
            v22 += (v3 - (^(v22)))
            var v55 float32 = float32(-2.25)
            _ = v55
        } else {
            if true {
                v4 = (-(v4))
                v56 := v2[(v3 + (^(v22)))]
                _ = v56
            }
            if (false || (0 < 89)) {
                v57 := GenS10{F0: -1, F1: "\n\t"}
                _ = v57
            }
        }
    case -2147483648:
        var v58 uint16 = 1
        _ = v58
    case -796:
        v1 = append(v1, 42)
    default:
        for fi59 := 0; fi59 < 8; fi59++ {
            for fi60 := 0; fi60 < 14; fi60++ {
                var v61 string = "foo bar"
                if false {
                    v4 -= (-(v4))
                } else {
                    var v62 uint = 18446744073709551615
                    v23 += (v22 - 4294967295)
                    delete(v2, v22)
                    _ = v62
                }
                _ = fi60
                _ = v61
            }
            for fi63 := 0; fi63 < 10; fi63++ {
                v23 = (v3 - v3)
                v4 -= 738
                _ = fi63
            }
            var v64 int64 = 380
            _ = fi59
            _ = v64
        }
    }
    if true {
        if false {
            for fi65 := 0; fi65 < 8; fi65++ {
                if (false || true) {
                    for fi66 := 0; fi66 < 20; fi66++ {
                        v67 := v2[(v23 + v22)]
                        if (false || false) {
                            var v68 uint64 = 103
                            v2[(v23 & v23)] = v67
                            fmt.Println(fi66, fi65)
                            _ = v68
                        } else {
                            v69 := GenS11{F0: -1, F1: 1}
                            switch v67 {
                            case "hello":
                                var v70 uint = 255
                                _ = v70
                            case "\n\t":
                                var v71 int = fi66
                                _ = v71
                            case "a":
                                fmt.Println(fi65, v23)
                                var v72 uint64 = 2
                                _ = v72
                            default:
                                var v73 int32 = 0
                                _ = v73
                            }
                            _ = v69
                        }
                        _ = fi66
                        _ = v67
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi74 := 0; fi74 < 3; fi74++ {
                        v22 = 4294967295
                        _ = fi74
                    }
                }
                var v75 uint16 = 65535
                _ = fi65
                _ = v75
            }
            if false {
                v23 = v3
                v3 = v23
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    v23 = v3
    for fi76 := 0; fi76 < 19; fi76++ {
        v23 = 65535
        _ = fi76
    }
    switch 1 {
    case 9223372036854775807:
        switch 1 {
        case 1:
            delete(v2, 2147483647)
        }
    case 4294967295:
        for fi77 := 0; fi77 < 5; fi77++ {
            var v78 bool = true
            if v78 {
                switch fi77 {
                case 32768:
                    var v79 uint8 = 181
                    _ = v79
                default:
                    for fi80 := 0; fi80 < 18; fi80++ {
                        var v81 int = 128
                        if (0.0 >= 0.0) {
                            var v82 float32 = ((1.5 - 0.1) + 0.1)
                            if v78 {
                                if false {
                                    v3 -= 659
                                    fmt.Println(v81, v4, fi77)
                                    v23 += v23
                                } else {
                                    var v83 int32 = -165
                                    switch (fi80 + 9223372036854775807) {
                                    case -9223372036854775808:
                                        switch ("" + "") {
                                        case "\n\t":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case "":
                                            if v78 {
                                                var v84 string = (("hello" + "hello") + (("\n\t" + "foo bar") + "\n\t"))
                                                v85 := GenS12{F0: 4294967295, F1: 128, F2: v4}
                                                switch v84 {
                                                case "":
                                                    if (-9223372036854775808 >= -1) {
                                                        _ = len(v2)
                                                        v3 = v22
                                                    }
                                                case "a":
                                                    v81 = fi77
                                                case "foo bar":
                                                    v86 := GenS2{F0: fi77, F1: 979, F2: 255}
                                                    _ = v86
                                                }
                                                _ = v84
                                                _ = v85
                                            } else {
                                                var v87 float64 = 3.14159
                                                _ = v87
                                            }
                                        case "a":
                                            if len(v1) > 0 {
                                                v1[0] = (v22 * v3)
                                            }
                                        default:
                                            v81 -= int(v4)
                                            var v88 string = (("\n\t" + "") + "")
                                            _ = v88
                                        }
                                    case 892:
                                        v89 := GenS11{F0: 42, F1: 0}
                                        if v78 {
                                            var v90 uint64 = 796
                                            var v91 uint16 = 42
                                            switch (-(v81)) {
                                            case 42:
                                                if true {
                                                    v92 := GenS1{F0: "hello", F1: (fi80 * (-(fi80))), F2: 297}
                                                    _ = v92
                                                } else {
                                                    var v93 int = (fi80 ^ (fi80 & v81))
                                                    var v94 uint8 = uint8(v91)
                                                    _ = v93
                                                    _ = v94
                                                }
                                            }
                                            _ = v90
                                            _ = v91
                                        }
                                        if v78 {
                                            v2[2] = ("" + "\n\t")
                                            if (2 > (v3 | 286)) {
                                                v78 = (!false)
                                                v5 = append(v5, 42)
                                            }
                                            var v95 float32 = v82
                                            switch fi80 {
                                            case 751:
                                                v96 := GenS12{F0: 4294967295, F1: 1, F2: v4}
                                                _ = v96
                                            case 32767:
                                                v97 := GenS13{F0: v3}
                                                _ = v97
                                            }
                                            v78 = v78
                                            _ = v95
                                        }
                                        v78 = v78
                                        _ = v89
                                    case -2147483648:
                                        var v98 float64 = (0.0 + (float64(float32(float64(0.1))) * (-2.25 - float64(float32(-2.25)))))
                                        _ = v98
                                    }
                                    if (("hello" + "") < ("hello" + ("hello" + ("hello" + "hello")))) {
                                        v99 := GenS14{F0: 499, F1: v4}
                                        _ = v99
                                    } else {
                                        switch ("a" + "\n\t") {
                                        case "":
                                            delete(v2, v23)
                                        default:
                                            v100 := GenS15{F0: (v23 & (^(v22))), F1: 1, F2: v4}
                                            _ = v100
                                        }
                                        v3 += v22
                                    }
                                    _ = v83
                                }
                                var v101 uint16 = 32767
                                _ = v101
                            } else {
                                v82 = -2.25
                                var v102 uint = 4294967295
                                v23 -= (v23 / (((v22 * v22)) | 1))
                                _ = v102
                            }
                            _ = v82
                        }
                        var v103 float64 = (0.0 / -2.25)
                        v104 := GenS16{F0: false, F1: -9223372036854775808, F2: 42}
                        _ = fi80
                        _ = v103
                        _ = v104
                        _ = v81
                    }
                }
                var v105 bool = v78
                _ = v105
            }
            _ = fi77
            _ = v78
        }
    case 1:
        if (-2.25 > 3.14159) {
            var v106 int8 = -1
            v2[(v23 + v3)] = ("a" + ("hello" + "hello"))
            _ = v106
        }
    }
    v23 -= (v3 >> ((0) & 63))
    v107 := GenS14{F0: 0, F1: 670}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v107
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
