package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint; F1 int32; F2 string }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int, p1 uint32, p2 int64) uint {
    v4 := []int{(p0 | (p0 % ((-1) | 1)))}
    var v5 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 781
}

func main() {
    var v1 uint = 1
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    v6 := GenS1{F0: (false && (122 != 2)), F1: -9223372036854775808, F2: -493}
    v7 := []int{4294967295, 2147483647}
    v1 += genFn0(4294967295, 42, 9223372036854775807)
    for fi8 := 0; fi8 < 7; fi8++ {
        v1 += (v1 >> ((0) & 63))
        _ = fi8
    }
    if ((2 <= 18446744073709551615) || (true || false)) {
        var v9 float32 = (0.0 * 1.5)
        v10 := v2[2147483647]
        _ = v10
        _ = v9
    }
    switch -122 {
    case 32767:
        v1 -= 2
    case 1:
        fmt.Println(v1)
    case 0:
        if (-144 >= 32768) {
            var v11 int8 = 1
            fmt.Println(v1)
            v12 := v3[("foo bar" + ("foo bar" + "\n\t"))]
            _ = v11
            _ = v12
        } else {
            if false {
                v1 = genFn0(1, 924, -719)
                for fi13 := 0; fi13 < 16; fi13++ {
                    if true {
                        for fi14 := 0; fi14 < 14; fi14++ {
                            if (true || true) {
                                if (((0.1 * 0.0) / (-2.25 + 0.1)) >= float64(0.0)) {
                                    fmt.Println(fi13, fi13)
                                    v15 := GenS2{F0: 42, F1: 0}
                                    _ = v15
                                }
                                fmt.Println(v1, fi13, fi14)
                                v1 += 4294967295
                                v16 := GenS2{F0: 867, F1: -1}
                                _ = v16
                            } else {
                                delete(v3, "hello")
                            }
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = fi14
                        }
                        v1 += v1
                    } else {
                        for fi17 := 0; fi17 < 13; fi17++ {
                            if false {
                                fmt.Println(fi17, v1, fi17)
                            } else {
                                if (!(!false)) {
                                    v1 -= (v1 ^ 255)
                                    fmt.Println(fi17, fi13)
                                    v18 := v3["a"]
                                    _ = v18
                                }
                                if false {
                                    fmt.Println(fi17, fi17, v1)
                                }
                                var v19 uint8 = 1
                                _ = v19
                            }
                            var v20 int16 = 32767
                            _ = fi17
                            _ = v20
                        }
                        for fi21 := 0; fi21 < 1; fi21++ {
                            switch "\n\t" {
                            case "a":
                                var v22 uint64 = 1
                                _ = v22
                            case "hello":
                                switch ("a" + ("foo bar" + "a")) {
                                case "hello":
                                    var v23 uint8 = 42
                                    _ = v23
                                case "":
                                    v1 = (v1 - genFn0((^(fi13)), 1, 9223372036854775807))
                                case "foo bar":
                                    switch (fi21 & 128) {
                                    case 127:
                                        v1 -= genFn0(159, 2147483647, -9223372036854775808)
                                    case 0:
                                        var v24 float64 = 1.5
                                        switch "hello" {
                                        case "foo bar":
                                            if false {
                                                var v25 uint16 = 1
                                                var v26 int64 = -9223372036854775808
                                                v25 = v25
                                                _ = v25
                                                _ = v26
                                            }
                                        case "\n\t":
                                            v27 := GenS1{F0: false, F1: (fi13 % ((fi21) | 1)), F2: 42}
                                            _ = v27
                                        default:
                                            var v28 float64 = (v24 + v24)
                                            _ = v28
                                        }
                                        switch (fi13 ^ fi21) {
                                        case 952:
                                            v24 = v24
                                        case -1:
                                            if (2 != 674) {
                                                v24 = v24
                                                var v29 float64 = float64(1.5)
                                                _ = v29
                                            }
                                        case 886:
                                            if (2147483647 == 891) {
                                                var v30 int8 = -119
                                                switch (fi21 << ((v1) & 63)) {
                                                case -581:
                                                    v31 := GenS1{F0: (32767 < 2), F1: (fi21 & (fi13 | (fi13 << (((v1 & v1)) & 63)))), F2: 1}
                                                    _ = v31
                                                case -150:
                                                    switch "a" {
                                                    case "hello":
                                                        v32 := GenS2{F0: 81, F1: 0}
                                                        _ = v32
                                                    case "\n\t":
                                                        var v33 float64 = ((v24 - (-2.25 / v24)) + ((v24 / (v24 - 1.5)) * ((0.1 / v24) / (v24 + 0.0))))
                                                        _ = v33
                                                    }
                                                    fmt.Println(v24)
                                                    v34 := GenS3{F0: (v1 * (v1 % (((v1 + (v1 | v1))) | 1))), F1: -1, F2: "\n\t"}
                                                    v1 = (^(v1))
                                                    var v35 string = ("a" + "a")
                                                    v1 = genFn0(fi13, 281, 0)
                                                    var v36 int = fi21
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                case 0:
                                                    var v37 int16 = 58
                                                    _ = v37
                                                default:
                                                    var v38 string = "a"
                                                    _ = v38
                                                }
                                                var v39 float64 = float64(0.1)
                                                _ = len(v3)
                                                _ = v30
                                                _ = v39
                                            }
                                        }
                                        _ = v24
                                    case 203:
                                        v1 = v1
                                    default:
                                        fmt.Println(fi13, v1, fi13)
                                    }
                                    if true {
                                        var v40 string = ("" + ("foo bar" + "foo bar"))
                                        v40 = v40
                                        _ = v40
                                    } else {
                                        v1 -= 81
                                        var v41 string = "foo bar"
                                        v42 := v2[2147483648]
                                        _ = v41
                                        _ = v42
                                    }
                                default:
                                    fmt.Println(v1)
                                }
                            }
                            v1 = v1
                            if false {
                                var v43 int64 = 526
                                switch "hello" {
                                case "hello":
                                    var v44 int = fi21
                                    _ = v44
                                }
                                if (255 >= v1) {
                                    fmt.Println(v1)
                                } else {
                                    _ = len(v7)
                                    switch fi21 {
                                    case 550:
                                        if false {
                                            v43 = (^(v43))
                                        }
                                    case -2147483648:
                                        var v45 uint8 = 255
                                        _ = v45
                                    default:
                                        var v46 int64 = v43
                                        _ = v46
                                    }
                                }
                                _ = v43
                            }
                            _ = fi21
                        }
                    }
                    var v47 uint64 = 462
                    for fi48 := 0; fi48 < 19; fi48++ {
                        switch "hello" {
                        case "":
                            v47 = 45
                        case "hello":
                            v49 := v2[0]
                            _ = v49
                        default:
                            v50 := GenS2{F0: 2, F1: 1}
                            _ = v50
                        }
                        delete(v3, "\n\t")
                        _ = fi48
                    }
                    _ = fi13
                    _ = v47
                }
                switch 0 {
                case -669:
                    for fi51 := 0; fi51 < 4; fi51++ {
                        var v52 string = ""
                        _ = fi51
                        _ = v52
                    }
                case 127:
                    if true {
                        v1 += (^(v1))
                    }
                case -1:
                    if len(v7) > 0 {
                        v7[0] = 32768
                    }
                default:
                    if true {
                        fmt.Println(v1)
                        fmt.Println(v1)
                    }
                }
            } else {
                for fi53 := 0; fi53 < 7; fi53++ {
                    v3[("" + ("a" + "\n\t"))] = 2
                    v1 = v1
                    var v54 float64 = -2.25
                    _ = fi53
                    _ = v54
                }
                v55 := GenS3{F0: 2, F1: -107, F2: "foo bar"}
                _ = v55
            }
        }
    }
    var v56 uint16 = 255
    var v57 float32 = 1.5
    if false {
        for fi58 := 0; fi58 < 16; fi58++ {
            for fi59 := 0; fi59 < 6; fi59++ {
                var v60 int32 = 1
                _ = fi59
                _ = v60
            }
            _ = fi58
        }
        var v61 int8 = -84
        _ = v61
    }
    if (false || (false || (!(182 >= 128)))) {
        v57 -= v57
        v62 := GenS4{F0: 1, F1: 1}
        var v63 int16 = 1
        _ = v62
        _ = v63
    }
    _ = len(v7)
    v64 := GenS5{F0: -404}
    fmt.Println(v56, v57, v56)
    switch 4294967295 {
    case -116:
        v65 := GenS3{F0: (v1 * 4294967295), F1: 2147483647, F2: ("foo bar" + "a")}
        _ = v65
    case -412:
        if false {
            if true {
                fmt.Println(v56, v57)
            } else {
                if (-9223372036854775808 == -9223372036854775808) {
                    v66 := GenS2{F0: (^(v56)), F1: 0}
                    switch "" {
                    case "a":
                        v67 := v2[65535]
                        _ = v67
                    }
                    _ = v66
                } else {
                    v57 = (1.5 * v57)
                }
                _ = len(v2)
            }
            for fi68 := 0; fi68 < 4; fi68++ {
                var v69 uint64 = 1
                _ = fi68
                _ = v69
            }
            var v70 uint32 = 807
            _ = v70
        }
    case 32767:
        v2[2147483648] = 42
    default:
        var v71 int = 2147483647
        _ = v71
    }
    v56 -= (v56 >> ((v1) & 63))
    var v72 int64 = -80
    for fi73 := 0; fi73 < 1; fi73++ {
        var v74 uint64 = 0
        v2[uint32(v1)] = (fi73 * fi73)
        _ = fi73
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v56
    _ = v57
    _ = v6
    _ = v64
    _ = v7
    _ = v72
}
