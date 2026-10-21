package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint16; F2 int }
type GenS3 struct { F0 int8; F1 string; F2 int64 }
type GenS4 struct { F0 uint; F1 int; F2 uint64 }
type GenS5 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint32, p1 uint8, p2 uint16) int16 {
    v7 := []int{32768}
    p0 = 1
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: (p1 / (((p1 - p1)) | 1)), F1: uint32(p1)}
        fmt.Println(fi8, p0)
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -95
}

func main() {
    var v1 bool = (0 <= 0)
    v2 := []uint32{233, 4294967295, 42, 2147483647}
    var v3 uint16 = 658
    var v4 bool = false
    var v5 int8 = 1
    v6 := make(map[uint32]uint32)
    var v10 int32 = -898
    v11 := v6[65535]
    if (true && v4) {
        v12 := GenS2{F0: 65535, F1: (^(v3)), F2: 42}
        v4 = (uint32(v3) == (^(v11)))
        for fi13 := 0; fi13 < 11; fi13++ {
            v2 = append(v2, 2147483648)
            _ = fi13
        }
        _ = v12
    } else {
        switch 32767 {
        case 9223372036854775807:
            for fi14 := 0; fi14 < 7; fi14++ {
                fmt.Println(v1)
                if v4 {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = fi14
            }
        case -2147483648:
            v3 += v3
        case 1:
            v15 := GenS1{F0: 128, F1: v11}
            _ = v15
        default:
            if true {
                fmt.Println(v10, v1)
            } else {
                var v16 uint32 = 315
                for fi17 := 0; fi17 < 6; fi17++ {
                    if len(v2) > 0 {
                        v2[0] = 944
                    }
                    var v18 int8 = v5
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
        }
        if (v4 && v1) {
            var v19 float64 = (1.5 - 3.14159)
            v20 := GenS1{F0: 1, F1: (v11 * (v11 / (((v11 + 65535)) | 1)))}
            _ = v19
            _ = v20
        } else {
            v1 = (2 >= 0)
        }
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 int16 = 32767
        _ = fi21
        _ = v22
    }
    if (!(!(!false))) {
        for fi23 := 0; fi23 < 9; fi23++ {
            switch "foo bar" {
            case "hello":
                var v24 uint64 = 798
                _ = v24
            case "a":
                v25 := GenS3{F0: 42, F1: "\n\t", F2: -1}
                _ = v25
            case "\n\t":
                switch "" {
                case "":
                    var v26 bool = (-2.25 == 1.5)
                    _ = v26
                default:
                    var v27 uint16 = v3
                    _ = v27
                }
            }
            v28 := GenS3{F0: (v5 & -1), F1: ("foo bar" + "hello"), F2: -905}
            _ = fi23
            _ = v28
        }
        if (4294967295 >= 1) {
            var v29 uint32 = v11
            fmt.Println(v5, v3)
            v30 := v6[v11]
            _ = v29
            _ = v30
        } else {
            var v31 int16 = -1
            if v1 {
                var v32 uint8 = 1
                for fi33 := 0; fi33 < 14; fi33++ {
                    fmt.Println(v11, v4)
                    v31 -= 1
                    var v34 int = (fi33 & fi33)
                    _ = fi33
                    _ = v34
                }
                var v35 uint8 = v32
                _ = v32
                _ = v35
            } else {
                if ((v31 >= v31) || v4) {
                    for fi36 := 0; fi36 < 17; fi36++ {
                        var v37 uint = 65535
                        for fi38 := 0; fi38 < 10; fi38++ {
                            var v39 uint = v37
                            _ = fi38
                            _ = v39
                        }
                        var v40 int8 = 1
                        _ = fi36
                        _ = v37
                        _ = v40
                    }
                    if v4 {
                        v31 = genFn0(v11, 127, (v3 >> ((2) & 63)))
                        var v41 uint32 = 152
                        var v42 int64 = -236
                        _ = v41
                        _ = v42
                    } else {
                        v11 = v11
                    }
                }
                var v43 float64 = -2.25
                _ = v43
            }
            for fi44 := 0; fi44 < 19; fi44++ {
                if true {
                    fmt.Println(v1, v1, v3)
                    var v45 int = fi44
                    _ = v45
                }
                for fi46 := 0; fi46 < 17; fi46++ {
                    v11 += v11
                    switch "foo bar" {
                    case "foo bar":
                        v11 += (^(v11))
                    }
                    var v47 uint8 = 52
                    _ = fi46
                    _ = v47
                }
                _ = len(v2)
                _ = fi44
            }
            _ = v31
        }
        if true {
            v48 := GenS4{F0: 255, F1: 2147483647, F2: 42}
            _ = v48
        }
    }
    v49 := GenS1{F0: 128, F1: 414}
    var v50 float32 = 1e10
    v51 := v6[v11]
    var v52 uint16 = (^(v3))
    _ = len(v2)
    var v53 int8 = 42
    var v54 int32 = 42
    if false {
        switch "a" {
        case "":
            switch ("" + ("a" + "\n\t")) {
            case "hello":
                fmt.Println(v52)
            default:
                switch 2147483647 {
                case -1000:
                    for fi55 := 0; fi55 < 10; fi55++ {
                        switch (fi55 / (((fi55 ^ (-(fi55)))) | 1)) {
                        case 214:
                            var v56 uint64 = 4294967295
                            _ = v56
                        case -1:
                            for fi57 := 0; fi57 < 7; fi57++ {
                                v6[v51] = v11
                                v58 := GenS1{F0: 42, F1: 4294967295}
                                _ = fi57
                                _ = v58
                            }
                        case -9223372036854775808:
                            switch "a" {
                            case "":
                                v59 := GenS4{F0: 327, F1: -1, F2: 42}
                                _ = v59
                            case "hello":
                                switch ("a" + ("foo bar" + "foo bar")) {
                                case "":
                                    if ((v10 ^ v54) < -1) {
                                        switch "" {
                                        case "":
                                            var v60 float32 = float32(((-2.25 + (3.14159 / -2.25)) - 3.14159))
                                            _ = v60
                                        case "\n\t":
                                            v1 = false
                                        case "foo bar":
                                            v5 -= v5
                                        }
                                        var v61 int64 = 0
                                        for fi62 := 0; fi62 < 19; fi62++ {
                                            if (v53 >= (v5 - (v53 * 127))) {
                                                var v63 uint = 257
                                                _ = v63
                                            }
                                            delete(v6, (v11 + v51))
                                            _ = fi62
                                        }
                                        _ = v61
                                    } else {
                                        if false {
                                            v2 = append(v2, v51)
                                        } else {
                                            if (-1 > 1) {
                                                v64 := v6[v51]
                                                _ = len(v6)
                                                _ = v64
                                            }
                                        }
                                    }
                                }
                            case "foo bar":
                                v50 -= v50
                            default:
                                v65 := GenS4{F0: 869, F1: fi55, F2: uint64(v3)}
                                _ = v65
                            }
                        default:
                            v10 += -629
                        }
                        v53 = (v53 | (v53 << ((42) & 63)))
                        _ = fi55
                    }
                case 388:
                    switch 2147483647 {
                    case -704:
                        var v66 uint16 = v52
                        _ = v66
                    case -1:
                        v10 += v54
                    case -9223372036854775808:
                        if v4 {
                            var v67 bool = false
                            for fi68 := 0; fi68 < 15; fi68++ {
                                fmt.Println(fi68, v50, v11)
                                _ = fi68
                            }
                            _ = v67
                        } else {
                            switch -900 {
                            case -1:
                                var v69 string = ("foo bar" + "")
                                _ = v69
                            }
                            v53 = -63
                        }
                    default:
                        for fi70 := 0; fi70 < 14; fi70++ {
                            _ = len(v6)
                            var v71 uint = 2
                            fmt.Println(v50)
                            _ = fi70
                            _ = v71
                        }
                    }
                case -15:
                    var v72 int64 = -9223372036854775808
                    _ = v72
                default:
                    if false {
                        var v73 int16 = 0
                        var v74 int8 = -60
                        v1 = v4
                        _ = v73
                        _ = v74
                    } else {
                        for fi75 := 0; fi75 < 15; fi75++ {
                            v4 = v1
                            _ = fi75
                        }
                        var v76 uint64 = 1
                        for fi77 := 0; fi77 < 7; fi77++ {
                            var v78 int32 = (v10 | (v10 + v54))
                            v51 += (v51 << ((2) & 63))
                            for fi79 := 0; fi79 < 14; fi79++ {
                                fmt.Println(v4, fi79)
                                fmt.Println(v52, v76)
                                _ = fi79
                            }
                            _ = fi77
                            _ = v78
                        }
                        _ = v76
                    }
                }
            }
        case "a":
            v51 = (v11 / (((^(v51))) | 1))
        case "hello":
            v80 := GenS3{F0: (v53 - 42), F1: "hello", F2: -477}
            _ = v80
        default:
            for fi81 := 0; fi81 < 16; fi81++ {
                if true {
                    var v82 bool = (v4 || (!(false || (787 == 1))))
                    v83 := v6[v51]
                    _ = v82
                    _ = v83
                } else {
                    v84 := GenS3{F0: v5, F1: ("hello" + (("foo bar" + "hello") + "")), F2: -1}
                    if len(v2) > 0 {
                        v2[0] = v51
                    }
                    _ = v84
                }
                switch fi81 {
                case -9223372036854775808:
                    v85 := GenS4{F0: 189, F1: (fi81 % (((fi81 & fi81)) | 1)), F2: 4294967295}
                    _ = v85
                case -642:
                    if ((1.5 * (float64((1.5 - v50)) / float64(v50))) <= 1.5) {
                        var v86 int = (fi81 / ((2147483647) | 1))
                        v10 = (v10 & (v54 >> ((914) & 63)))
                        _ = v86
                    }
                case 2147483647:
                    for fi87 := 0; fi87 < 3; fi87++ {
                        var v88 bool = true
                        switch 4294967295 {
                        case 703:
                            v89 := GenS3{F0: (v5 * v53), F1: "a", F2: -9223372036854775808}
                            _ = v89
                        }
                        v88 = v1
                        _ = fi87
                        _ = v88
                    }
                default:
                    var v90 int = fi81
                    _ = v90
                }
                _ = fi81
            }
        }
        for fi91 := 0; fi91 < 12; fi91++ {
            v92 := GenS5{F0: v4, F1: (^(v52)), F2: 1}
            v93 := v6[(v11 >> ((420) & 63))]
            var v94 string = ("foo bar" + "foo bar")
            _ = fi91
            _ = v92
            _ = v93
            _ = v94
        }
    }
    if true {
        v95 := GenS6{F0: 255, F1: 42}
        var v96 int16 = -32768
        v3 -= v52
        _ = v95
        _ = v96
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
