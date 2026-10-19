package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint8; F1 int; F2 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 uint64; F1 int; F2 int16 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint16; F1 uint16; F2 int8 }
func genFn0() uint32 {
    var v4 uint16 = 32767
    if ((32767 != int(v4)) && (1.5 <= 1.5)) {
        fmt.Println(v4)
    }
    v5 := GenS1{F0: 2147483647, F1: (22 != -1)}
    if ((3.14159 * 1e10) <= (1.5 * 3.14159)) {
        for fi6 := 0; fi6 < 20; fi6++ {
            v4 -= (v4 & 32767)
            _ = fi6
        }
    } else {
        var v7 float32 = 1e10
        if true {
            switch "\n\t" {
            case "foo bar":
                var v8 string = ("" + "a")
                _ = v8
            }
        } else {
            if false {
                fmt.Println(v4, v4)
            } else {
                var v9 int = -9223372036854775808
                var v10 float32 = (float32((0.1 - 1.5)) - -2.25)
                _ = v10
                _ = v9
            }
        }
        _ = v7
    }
    v11 := GenS1{F0: -398, F1: true}
    _ = v11
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 uint8, p1 int8) uint {
    if true {
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 string = (("\n\t" + "a") + ("hello" + "\n\t"))
            v14 := GenS2{F0: 479, F1: false}
            v15 := []int{706}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        var v16 int16 = -32768
        _ = v16
    }
    v17 := GenS1{F0: 886, F1: (p1 == p1)}
    v18 := GenS2{F0: 2, F1: false}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 2
}

func main() {
    v1 := []uint32{0}
    v2 := []uint32{}
    v3 := []uint32{2, 2, 4294967295, 0}
    switch "foo bar" {
    case "a":
        v2 = append(v2, 39)
    case "":
        v19 := GenS3{F0: 42, F1: -838, F2: 255}
        _ = v19
    case "foo bar":
        if false {
            var v20 int8 = 100
            fmt.Println(v20)
            _ = v20
        } else {
            v21 := GenS4{F0: 2}
            var v22 uint16 = 42
            v22 = v22
            _ = v21
            _ = v22
        }
    }
    var v23 bool = false
    var v24 uint8 = 174
    if false {
        var v25 uint64 = 1
        var v26 int = 192
        _ = v25
        _ = v26
    } else {
        var v27 uint32 = 65535
        v27 += genFn0()
        switch "" {
        case "hello":
            switch ("hello" + "foo bar") {
            case "hello":
                if true {
                    v27 += uint32(v24)
                    switch ("foo bar" + "") {
                    case "":
                        var v28 uint = 255
                        _ = v28
                    case "hello":
                        _ = len(v3)
                    case "a":
                        if (1.5 > 3.14159) {
                            var v29 uint = 0
                            v30 := GenS4{F0: v29}
                            _ = v29
                            _ = v30
                        } else {
                            var v31 uint16 = 286
                            v2 = append(v2, (v27 << ((genFn1((^(v24)), 1)) & 63)))
                            _ = v31
                        }
                    }
                    if (v23 || true) {
                        v32 := GenS5{F0: 18446744073709551615}
                        _ = v32
                    }
                } else {
                    for fi33 := 0; fi33 < 19; fi33++ {
                        var v34 int32 = 42
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi33
                        _ = v34
                    }
                    v1 = append(v1, (^(v27)))
                }
            }
        default:
            v35 := GenS6{F0: (v27 & (v27 % (((v27 * v27)) | 1)))}
            _ = v35
        }
        _ = v27
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        v37 := GenS6{F0: 918}
        var v38 uint = 784
        _ = fi36
        _ = v37
        _ = v38
    }
    if true {
        fmt.Println(v24)
        switch -943 {
        case 42:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -9223372036854775808:
            if (!v23) {
                var v39 int = 128
                var v40 uint32 = 65535
                var v41 uint64 = 18446744073709551615
                _ = v39
                _ = v40
                _ = v41
            } else {
                for fi42 := 0; fi42 < 10; fi42++ {
                    var v43 uint8 = v24
                    _ = len(v2)
                    switch "a" {
                    case "hello":
                        var v44 bool = v23
                        _ = v44
                    case "a":
                        v45 := GenS7{F0: 89, F1: 65535}
                        _ = v45
                    default:
                        for fi46 := 0; fi46 < 6; fi46++ {
                            v24 = (v24 >> ((255) & 63))
                            _ = fi46
                        }
                    }
                    _ = fi42
                    _ = v43
                }
                if true {
                    var v47 uint32 = 65535
                    if (42 <= 163) {
                        fmt.Println(v23, v24)
                    } else {
                        var v48 bool = v23
                        var v49 bool = (!true)
                        _ = v48
                        _ = v49
                    }
                    var v50 float32 = (-2.25 - (0.1 * 1.5))
                    _ = v47
                    _ = v50
                } else {
                    var v51 int32 = 0
                    _ = v51
                }
            }
        }
    } else {
        var v52 bool = true
        switch 32767 {
        case 452:
            for fi53 := 0; fi53 < 5; fi53++ {
                if v23 {
                    v24 = (^(v24))
                }
                v52 = (255 > 2)
                switch 195 {
                case -1:
                    fmt.Println(v24, v52, fi53)
                case 937:
                    switch ("" + "hello") {
                    case "foo bar":
                        switch fi53 {
                        case 32767:
                            if v52 {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v54 float64 = 1.5
                                _ = v54
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = 1
                                }
                                for fi55 := 0; fi55 < 14; fi55++ {
                                    v56 := GenS8{F0: (!(!(!false))), F1: (-(fi55))}
                                    _ = fi55
                                    _ = v56
                                }
                                v24 = (v24 >> ((0) & 63))
                            }
                        default:
                            if len(v3) > 0 {
                                v3[0] = 875
                            }
                        }
                    case "\n\t":
                        v57 := GenS9{F0: 42, F1: fi53, F2: -771}
                        _ = v57
                    case "hello":
                        _ = len(v1)
                    default:
                        fmt.Println(v24, v52)
                    }
                }
                _ = fi53
            }
        case -9223372036854775808:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 127:
            switch "hello" {
            case "":
                v58 := GenS1{F0: -416, F1: v23}
                _ = v58
            case "hello":
                if (0 != 65535) {
                    for fi59 := 0; fi59 < 6; fi59++ {
                        v60 := GenS10{F0: "a"}
                        v52 = (!(((v24 ^ 247) <= v24) || (18446744073709551615 < 4294967295)))
                        _ = fi59
                        _ = v60
                    }
                    for fi61 := 0; fi61 < 9; fi61++ {
                        v24 -= 81
                        if v52 {
                            switch "\n\t" {
                            case "":
                                if len(v2) > 0 {
                                    v2[0] = 4294967295
                                }
                            case "\n\t":
                                if (816 != 42) {
                                    fmt.Println(v52, v23)
                                }
                            default:
                                for fi62 := 0; fi62 < 12; fi62++ {
                                    _ = len(v3)
                                    v23 = (128 != (v24 + v24))
                                    _ = fi62
                                }
                            }
                        } else {
                            fmt.Println(v52, v23, v52)
                            var v63 int16 = 1
                            v52 = (0 < 1)
                            _ = v63
                        }
                        var v64 uint64 = 2
                        _ = fi61
                        _ = v64
                    }
                }
            case "\n\t":
                var v65 int16 = 32767
                _ = v65
            default:
                var v66 uint16 = 574
                _ = v66
            }
        default:
            var v67 uint16 = 241
            _ = v67
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v52
    }
    v23 = v23
    if (819 >= 4294967295) {
        var v68 int = 32768
        _ = v68
    }
    v69 := GenS11{F0: 1}
    v70 := GenS11{F0: 18446744073709551615}
    if v23 {
        for fi71 := 0; fi71 < 19; fi71++ {
            v72 := GenS12{F0: 255}
            v23 = (-2147483648 >= 0)
            _ = fi71
            _ = v72
        }
        v23 = ((42 <= 2) && v23)
    } else {
        v24 = (v24 & v24)
        var v73 float32 = float32(1e10)
        for fi74 := 0; fi74 < 20; fi74++ {
            for fi75 := 0; fi75 < 1; fi75++ {
                switch "\n\t" {
                case "hello":
                    if (!(!(v73 >= 0.1))) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v73)
                        if len(v1) > 0 {
                            v1[0] = 384
                        }
                        v23 = v23
                        switch fi75 {
                        case -702:
                            if (false && (1e10 < 1.5)) {
                                v76 := GenS12{F0: 32767}
                                fmt.Println(fi74)
                                var v77 uint16 = 1
                                _ = v76
                                _ = v77
                            } else {
                                switch ("\n\t" + "hello") {
                                case "foo bar":
                                    v23 = (v23 || (!v23))
                                case "a":
                                    v73 = (v73 / float32(1e10))
                                case "hello":
                                    if ((!true) && (0 != 0)) {
                                        fmt.Println(fi74, v23)
                                        var v78 int = fi75
                                        _ = v78
                                    }
                                default:
                                    var v79 uint = 42
                                    var v80 uint16 = 32767
                                    _ = v79
                                    _ = v80
                                }
                                if (18446744073709551615 < 18446744073709551615) {
                                    var v81 uint16 = 1
                                    if (v23 && (!v23)) {
                                        switch fi74 {
                                        case -2147483648:
                                            switch (fi74 | fi75) {
                                            case 42:
                                                var v82 uint64 = 4294967295
                                                _ = v82
                                            default:
                                                v83 := GenS13{F0: v81, F1: v81, F2: 127}
                                                _ = v83
                                            }
                                            v24 = v24
                                            fmt.Println(fi74)
                                        case -9223372036854775808:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case 32767:
                                            v84 := GenS4{F0: genFn1((v24 - 127), 100)}
                                            _ = v84
                                        }
                                    }
                                    fmt.Println(v23)
                                    _ = v81
                                } else {
                                    v73 += 1e10
                                    var v85 float64 = 0.0
                                    _ = v85
                                }
                            }
                        case -9223372036854775808:
                            var v86 int8 = 127
                            _ = v86
                        case -664:
                            v24 = v24
                        }
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        switch ("hello" + (("foo bar" + ("hello" + "foo bar")) + "hello")) {
                        case "hello":
                            if (42 < 869) {
                                var v87 int32 = 0
                                var v88 uint64 = 4294967295
                                _ = v87
                                _ = v88
                            }
                        case "":
                            var v89 float32 = v73
                            _ = v89
                        case "\n\t":
                            _ = len(v1)
                        }
                        if (false || ((-226 <= 0) && v23)) {
                            var v90 uint8 = (v24 % ((v24) | 1))
                            var v91 uint16 = 2
                            _ = v90
                            _ = v91
                        } else {
                            var v92 uint32 = 42
                            v24 -= (^(v24))
                            fmt.Println(fi74, v23)
                            var v93 int64 = -465
                            if (genFn1((v24 / ((v24) | 1)), -1) >= 4294967295) {
                                var v94 int8 = 1
                                _ = v94
                            }
                            _ = v92
                            _ = v93
                        }
                        if v23 {
                            var v95 uint8 = 2
                            v24 -= (v24 % (((^(v24))) | 1))
                            _ = v95
                        }
                        v24 -= v24
                        v23 = (v24 == (v24 / ((128) | 1)))
                        v24 = v24
                    }
                case "foo bar":
                    if (!v23) {
                        v96 := GenS8{F0: false, F1: fi74}
                        _ = v96
                    }
                default:
                    var v97 float64 = 3.14159
                    _ = v97
                }
                var v98 uint = uint(v24)
                _ = fi75
                _ = v98
            }
            v23 = false
            switch (fi74 + 32768) {
            case 42:
                var v99 int = (-(fi74))
                _ = v99
            case 9223372036854775807:
                var v100 uint8 = (^(v24))
                _ = v100
            case 32767:
                v101 := GenS4{F0: 1}
                _ = v101
            }
            _ = fi74
        }
        _ = v73
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v69
    _ = v70
}
