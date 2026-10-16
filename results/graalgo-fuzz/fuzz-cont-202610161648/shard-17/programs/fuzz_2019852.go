package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 uint32; F2 int16 }
type GenS8 struct { F0 int64; F1 int32; F2 int16 }
type GenS9 struct { F0 uint; F1 int64 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int64; F1 uint8; F2 bool }
type GenS14 struct { F0 int; F1 int16 }
type GenS15 struct { F0 int64 }

func main() {
    var v1 float64 = (0.1 * 1e10)
    v2 := []uint32{2, 3, 65535, 741}
    v3 := []uint32{0, 42}
    v4 := make(map[uint32]uint32)
    v5 := []uint32{4294967295, 2147483647}
    v6 := make(map[int]uint32)
    if true {
        v7 := GenS1{F0: "", F1: -9223372036854775808}
        v1 = v1
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := GenS2{F0: 127}
            if false {
                v10 := GenS3{F0: 575}
                for fi11 := 0; fi11 < 12; fi11++ {
                    if false {
                        v1 -= float64(3.14159)
                    }
                    var v12 bool = false
                    _ = len(v3)
                    _ = fi11
                    _ = v12
                }
                _ = v10
            } else {
                var v13 uint16 = 32767
                for fi14 := 0; fi14 < 3; fi14++ {
                    if false {
                        if true {
                            if (!false) {
                                if false {
                                    v15 := GenS4{F0: 638, F1: -128}
                                    switch (-(fi8)) {
                                    case 4294967295:
                                        v16 := GenS4{F0: 0, F1: -81}
                                        _ = v16
                                    default:
                                        var v17 uint32 = 1
                                        _ = v17
                                    }
                                    _ = v15
                                }
                                var v18 uint32 = 2147483648
                                _ = v18
                            } else {
                                v1 += (float64(1.5) * (v1 * v1))
                                v19 := GenS1{F0: (("hello" + (("\n\t" + "a") + "")) + ("hello" + "\n\t")), F1: fi8}
                                _ = v19
                            }
                            switch "hello" {
                            case "":
                                v13 += (v13 - (v13 / (((v13 + (v13 & 1))) | 1)))
                                fmt.Println(v13, v13, fi14)
                            case "hello":
                                v20 := v4[523]
                                _ = v20
                            default:
                                switch ("hello" + "a") {
                                case "":
                                    v21 := GenS5{F0: 0, F1: int16(v13), F2: 1}
                                    switch int(v13) {
                                    case 127:
                                        var v22 int8 = 42
                                        _ = v22
                                    case 2147483647:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    _ = v21
                                case "a":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "\n\t":
                                    if (1000 < 918) {
                                        if (0 < 0) {
                                            var v23 bool = true
                                            if true {
                                                v24 := GenS6{F0: -952}
                                                _ = v24
                                            } else {
                                                switch ("\n\t" + "\n\t") {
                                                case "":
                                                    v23 = v23
                                                    if true {
                                                        if v23 {
                                                            v1 = (1e10 - 3.14159)
                                                            switch (("a" + (("\n\t" + ("" + "a")) + ("hello" + ""))) + ((("\n\t" + ("" + "foo bar")) + "foo bar") + ("" + ""))) {
                                                            case "":
                                                                v25 := GenS6{F0: 252}
                                                                _ = v25
                                                            case "\n\t":
                                                                v23 = false
                                                            default:
                                                                v26 := GenS7{F0: 0, F1: 0, F2: 0}
                                                                _ = v26
                                                            }
                                                            switch fi14 {
                                                            case -9223372036854775808:
                                                                v13 = v13
                                                            default:
                                                                var v27 uint64 = 1
                                                                _ = v27
                                                            }
                                                            v28 := GenS7{F0: 9223372036854775807, F1: 42, F2: -944}
                                                            v29 := GenS6{F0: 1}
                                                            _ = v28
                                                            _ = v29
                                                        }
                                                        var v30 uint = 1
                                                        _ = v30
                                                    }
                                                    var v31 float64 = (v1 - (v1 - v1))
                                                    _ = v31
                                                case "hello":
                                                    v32 := GenS2{F0: 42}
                                                    _ = v32
                                                case "a":
                                                    var v33 int32 = -583
                                                    _ = v33
                                                }
                                                _ = len(v5)
                                            }
                                            _ = v23
                                        }
                                    } else {
                                        if ((2 == 590) && true) {
                                            var v34 int32 = 2147483647
                                            if true {
                                                fmt.Println(fi8)
                                                v1 -= v1
                                                if (float32(v1) > 0.1) {
                                                    v35 := GenS1{F0: "\n\t", F1: fi8}
                                                    v1 = v1
                                                    _ = v35
                                                }
                                            }
                                            if true {
                                                var v36 int16 = 1
                                                _ = v36
                                            }
                                            _ = v34
                                        }
                                    }
                                }
                            }
                        }
                    } else {
                        switch fi14 {
                        case 712:
                            if false {
                                v13 = (v13 - (v13 * 2))
                                var v37 string = "hello"
                                v38 := GenS8{F0: 0, F1: 2147483647, F2: 42}
                                _ = v37
                                _ = v38
                            }
                        case -1:
                            _ = len(v3)
                        case 0:
                            if true {
                                v1 -= v1
                            } else {
                                if false {
                                    v1 = (float64(1e10) - 0.1)
                                    v1 += 1e10
                                    if true {
                                        var v39 int16 = 73
                                        switch ("\n\t" + "hello") {
                                        case "foo bar":
                                            var v40 uint = 18446744073709551615
                                            _ = v40
                                        }
                                        _ = v39
                                    } else {
                                        var v41 int16 = -32768
                                        var v42 int16 = v41
                                        _ = v41
                                        _ = v42
                                    }
                                } else {
                                    if false {
                                        v43 := GenS3{F0: v13}
                                        v2 = append(v2, 2147483648)
                                        _ = v43
                                    }
                                }
                                if true {
                                    v44 := GenS2{F0: 128}
                                    v45 := v4[2147483648]
                                    _ = v44
                                    _ = v45
                                }
                                v1 = v1
                            }
                        default:
                            var v46 bool = (!false)
                            _ = v46
                        }
                        v47 := GenS9{F0: 975, F1: -9223372036854775808}
                        if len(v3) > 0 {
                            v3[0] = 331
                        }
                        _ = v47
                    }
                    _ = fi14
                }
                _ = v13
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    var v48 int32 = 0
    v49 := GenS6{F0: 0}
    switch 2147483647 {
    case 192:
        v50 := GenS9{F0: 255, F1: -1}
        _ = v50
    case 128:
        var v51 uint = 878
        _ = v51
    }
    switch 608 {
    case 127:
        v48 = -2147483648
    case 458:
        v1 -= 3.14159
    case 32768:
        switch 210 {
        case 896:
            switch 2147483647 {
            case 462:
                fmt.Println(v48, v1)
            case 752:
                v48 = (v48 / ((0) | 1))
            case -477:
                if (1 <= 18446744073709551615) {
                    v48 += v48
                    if (!(2 != 18446744073709551615)) {
                        var v52 int32 = 42
                        v48 = v52
                        var v53 uint16 = 65535
                        _ = v52
                        _ = v53
                    } else {
                        if (42 != 1) {
                            var v54 int32 = (v48 | (v48 ^ 749))
                            for fi55 := 0; fi55 < 12; fi55++ {
                                var v56 int32 = -1
                                _ = fi55
                                _ = v56
                            }
                            var v57 int16 = -1
                            _ = v54
                            _ = v57
                        }
                        v48 = -2147483648
                    }
                    var v58 int = 127
                    _ = v58
                } else {
                    v59 := GenS6{F0: 9223372036854775807}
                    _ = v59
                }
            default:
                switch "hello" {
                case "foo bar":
                    var v60 int64 = 42
                    _ = v60
                case "a":
                    for fi61 := 0; fi61 < 3; fi61++ {
                        for fi62 := 0; fi62 < 7; fi62++ {
                            switch "\n\t" {
                            case "foo bar":
                                if len(v2) > 0 {
                                    v2[0] = 42
                                }
                            }
                            _ = fi62
                        }
                        v1 -= ((v1 / float64(-2.25)) / (1e10 / v1))
                        var v63 uint16 = 32767
                        _ = fi61
                        _ = v63
                    }
                case "hello":
                    v64 := GenS7{F0: -1, F1: 2147483648, F2: 1000}
                    _ = v64
                default:
                    v65 := GenS9{F0: 18446744073709551615, F1: 1}
                    _ = v65
                }
            }
        case 32768:
            if (v48 != 2147483647) {
                var v66 uint8 = 88
                v1 = v1
                _ = v66
            } else {
                var v67 int16 = -32768
                if true {
                    v68 := GenS10{F0: "hello"}
                    if true {
                        _ = len(v6)
                        for fi69 := 0; fi69 < 17; fi69++ {
                            fmt.Println(v1, v67)
                            switch ("a" + "foo bar") {
                            case "hello":
                                var v70 float32 = (3.14159 + (1e10 - -2.25))
                                _ = v70
                            case "":
                                v71 := GenS11{F0: 4294967295, F1: 73, F2: v67}
                                _ = v71
                            default:
                                v72 := GenS12{F0: 127}
                                _ = v72
                            }
                            _ = fi69
                        }
                    } else {
                        var v73 uint8 = 255
                        _ = v73
                    }
                    _ = v68
                }
                _ = v67
            }
        }
    default:
        fmt.Println(v48, v1)
    }
    var v74 float32 = (1.5 - (1.5 - (((float32(float64(-2.25)) - 3.14159) - (1.5 / 3.14159)) + 3.14159)))
    switch -752 {
    case 173:
        v48 = (v48 + v48)
    case 9223372036854775807:
        if (9223372036854775807 != 32768) {
            v48 = 2147483647
        } else {
            if len(v2) > 0 {
                v2[0] = 308
            }
            if false {
                _ = len(v4)
                _ = len(v6)
            } else {
                v1 -= v1
            }
        }
    case -9223372036854775808:
        fmt.Println(v1, v48)
    }
    var v75 uint64 = 338
    if (!(false && true)) {
        fmt.Println(v74, v75)
        v1 -= (float64(((v74 - (((v74 * 0.0) + v74) * ((((v74 + v74) * (0.1 / (float32(1.5) / 0.1))) * ((float32(v1) * float32(-2.25)) / v74)) / (0.1 / v74)))) + v74)) + -2.25)
        for fi76 := 0; fi76 < 17; fi76++ {
            fmt.Println(v75, v75, v48)
            v75 += (v75 & 18446744073709551615)
            _ = fi76
        }
    } else {
        for fi77 := 0; fi77 < 6; fi77++ {
            v75 += (^(v75))
            _ = fi77
        }
        if true {
            if len(v3) > 0 {
                v3[0] = 42
            }
            v78 := v6[42]
            v79 := GenS5{F0: 32767, F1: 1000, F2: v75}
            _ = v78
            _ = v79
        }
        for fi80 := 0; fi80 < 7; fi80++ {
            v75 -= (v75 << ((483) & 63))
            var v81 string = "hello"
            if true {
                v82 := GenS2{F0: 2}
                for fi83 := 0; fi83 < 2; fi83++ {
                    var v84 uint64 = 4294967295
                    v85 := GenS13{F0: 42, F1: 45, F2: (719 != 18446744073709551615)}
                    v86 := GenS14{F0: (fi80 & fi83), F1: -70}
                    _ = fi83
                    _ = v84
                    _ = v85
                    _ = v86
                }
                v74 += v74
                _ = v82
            }
            _ = fi80
            _ = v81
        }
    }
    switch 400 {
    case 909:
        if true {
            fmt.Println(v74, v1)
        } else {
            _ = len(v4)
        }
    case 42:
        fmt.Println(v48, v74)
    default:
        var v87 bool = false
        _ = v87
    }
    v75 = (v75 & v75)
    delete(v6, 32767)
    for fi88 := 0; fi88 < 10; fi88++ {
        var v89 uint = 65535
        for fi90 := 0; fi90 < 12; fi90++ {
            var v91 float32 = v74
            _ = fi90
            _ = v91
        }
        switch fi88 {
        case 598:
            if true {
                var v92 uint = 18446744073709551615
                var v93 float32 = 0.0
                _ = v92
                _ = v93
            } else {
                if (-4 <= 712) {
                    var v94 uint8 = 42
                    _ = v94
                } else {
                    var v95 string = ("a" + "a")
                    fmt.Println(v48)
                    _ = v95
                }
            }
        case 4294967295:
            v74 += v74
        }
        _ = fi88
        _ = v89
    }
    var v96 int8 = 42
    var v97 float32 = float32(((float64(v74) * (v1 / (((((1e10 / (float64(v74) - 1e10)) + v1) + ((((float64(v74) - ((((((v1 + v1) - v1) / (v1 + (-2.25 - ((0.1 / v1) - v1)))) - v1) / v1) * 0.1)) * v1) / 3.14159) + (v1 + (1.5 - 0.1)))) * v1) - (v1 * (((v1 / v1) / float64((v74 / v74))) + ((v1 + (1.5 + ((v1 * v1) * (v1 * v1)))) * v1)))))) - v1))
    for fi98 := 0; fi98 < 19; fi98++ {
        var v99 int64 = 1
        _ = fi98
        _ = v99
    }
    v100 := GenS15{F0: int64(v48)}
    if false {
        var v101 int = 2147483647
        _ = v101
    } else {
        for fi102 := 0; fi102 < 19; fi102++ {
            var v103 int8 = v96
            _ = fi102
            _ = v103
        }
        if (42 != -978) {
            v104 := GenS7{F0: 42, F1: 1, F2: -644}
            for fi105 := 0; fi105 < 9; fi105++ {
                var v106 int8 = 1
                var v107 int32 = (v48 * (v48 << ((0) & 63)))
                switch fi105 {
                case 209:
                    v75 = (v75 << ((255) & 63))
                case 32768:
                    var v108 float32 = (v97 - (0.0 * ((v97 + v97) / 1e10)))
                    _ = v108
                case 9223372036854775807:
                    for fi109 := 0; fi109 < 17; fi109++ {
                        fmt.Println(fi109, fi105, v107)
                        var v110 int16 = -32768
                        var v111 bool = false
                        _ = fi109
                        _ = v110
                        _ = v111
                    }
                }
                _ = fi105
                _ = v106
                _ = v107
            }
            v112 := GenS7{F0: 1, F1: 633, F2: 32767}
            _ = v104
            _ = v112
        } else {
            v48 += v48
        }
        delete(v6, 42)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v74
    _ = v75
    _ = v96
    _ = v97
}
