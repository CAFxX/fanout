package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int32 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint16; F1 int64; F2 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int32; F1 uint16; F2 int32 }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 int8; F1 uint64 }
type GenS9 struct { F0 uint16; F1 string }
type GenS10 struct { F0 uint32; F1 string; F2 int64 }
type GenS11 struct { F0 uint8; F1 uint }
type GenS12 struct { F0 bool; F1 uint8 }
func genFn0(p0 int32, p1 uint64, p2 bool) int32 {
    var v5 float32 = 3.14159
    v6 := GenS1{F0: 65535, F1: p2, F2: (p0 >> ((255) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p0
}
func genFn1(p0 uint8, p1 uint16) int32 {
    if ("hello" != "\n\t") {
        var v7 bool = false
        _ = v7
    } else {
        p1 = p1
        var v8 uint = 65535
        _ = v8
    }
    var v9 uint16 = 2
    _ = p0
    _ = p1
    _ = v9
    return int32(p0)
}

func main() {
    var v1 uint64 = 584
    var v2 bool = (true || (9223372036854775807 >= -9223372036854775808))
    var v3 uint8 = 213
    v4 := make(map[uint32]int)
    for fi10 := 0; fi10 < 8; fi10++ {
        if (fi10 <= 238) {
            switch (fi10 + (fi10 << ((42) & 63))) {
            case 0:
                v3 -= (v3 * 42)
            case 2147483647:
                switch 223 {
                case 9223372036854775807:
                    var v11 int8 = -1
                    _ = v11
                case 127:
                    v12 := []int{(fi10 - (fi10 / (((fi10 & (^(fi10)))) | 1))), (fi10 & (^(fi10))), int(v1), (fi10 << ((18446744073709551615) & 63))}
                    _ = v12
                case 32768:
                    v13 := []int{fi10, fi10, 128}
                    _ = v13
                default:
                    _ = len(v4)
                }
            case 42:
                if v2 {
                    var v14 int16 = -32768
                    _ = v14
                }
            }
            switch -1 {
            case 689:
                if ((42 >= 1) && (float32(1e10) != float32(0.1))) {
                    if (false && v2) {
                        v2 = (0 >= 0)
                        fmt.Println(v1, fi10)
                        for fi15 := 0; fi15 < 12; fi15++ {
                            if true {
                                var v16 int = 128
                                _ = v16
                            } else {
                                switch fi15 {
                                case 9223372036854775807:
                                    v17 := []int{fi10, fi15}
                                    _ = v17
                                case 42:
                                    v3 += v3
                                case 0:
                                    var v18 uint16 = 958
                                    _ = v18
                                default:
                                    v19 := []int{fi15}
                                    v20 := GenS1{F0: 2147483648, F1: v2, F2: -672}
                                    _ = v19
                                    _ = v20
                                }
                                if true {
                                    if v2 {
                                        if (9223372036854775807 > 0) {
                                            var v21 uint64 = 4294967295
                                            var v22 uint64 = (v1 >> ((0) & 63))
                                            _ = v21
                                            _ = v22
                                        }
                                    } else {
                                        var v23 uint = 1
                                        var v24 uint = v23
                                        v23 -= v23
                                        _ = v23
                                        _ = v24
                                    }
                                    v1 -= v1
                                } else {
                                    var v25 int32 = genFn0(-1, 972, v2)
                                    fmt.Println(v2, v3)
                                    var v26 bool = (v25 > (v25 - 416))
                                    delete(v4, 2147483648)
                                    _ = v25
                                    _ = v26
                                }
                            }
                            var v27 int64 = 1
                            _ = fi15
                            _ = v27
                        }
                    } else {
                        v28 := GenS2{F0: (-2.25 <= 0.1), F1: 2147483647}
                        for fi29 := 0; fi29 < 16; fi29++ {
                            fmt.Println(v3)
                            switch fi10 {
                            case 1:
                                var v30 uint64 = (v1 ^ v1)
                                _ = v30
                            case -2147483648:
                                v1 += 515
                            case -778:
                                fmt.Println(v1, fi10, v2)
                            }
                            if v2 {
                                if (835 != 167) {
                                    if v2 {
                                        delete(v4, 0)
                                        v31 := GenS1{F0: uint32(v1), F1: v2, F2: 2147483647}
                                        _ = v31
                                    }
                                    var v32 float32 = (((0.1 / 1e10) * -2.25) * ((1.5 * ((float32(((3.14159 * (((float64(1e10) + ((0.1 + (0.1 + 0.0)) + 1e10)) / 1e10) - 1.5)) - (0.1 * 0.1))) / 1e10) / float32(1e10))) * 1.5))
                                    v32 += 3.14159
                                    _ = v32
                                } else {
                                    fmt.Println(v2)
                                    var v33 int64 = 42
                                    switch fi29 {
                                    case 606:
                                        switch fi29 {
                                        case 128:
                                            var v34 int = (fi10 & (fi10 / ((fi29) | 1)))
                                            _ = v34
                                        case 32767:
                                            if (0.0 > (1.5 / 1.5)) {
                                                if (!true) {
                                                    var v35 float32 = float32(-2.25)
                                                    fmt.Println(v1, v33, fi10)
                                                    v36 := GenS1{F0: 1, F1: v2, F2: 42}
                                                    _ = v35
                                                    _ = v36
                                                }
                                                fmt.Println(v33, v33)
                                                v37 := []int{fi29, fi10, fi29, fi29}
                                                _ = v37
                                            }
                                        default:
                                            fmt.Println(fi10, v33, v2)
                                        }
                                    case -2147483648:
                                        v38 := GenS3{F0: 519, F1: (-(v33)), F2: 153}
                                        _ = v38
                                    default:
                                        var v39 int32 = genFn0(-1, v1, true)
                                        _ = v39
                                    }
                                    _ = v33
                                }
                                fmt.Println(v3, v3, fi10)
                                v40 := []int{(fi29 % ((fi29) | 1))}
                                _ = v40
                            } else {
                                v3 = 128
                                _ = len(v4)
                                switch "foo bar" {
                                case "":
                                    v41 := []int{(fi10 - 640), (fi29 + (fi10 % ((fi10) | 1))), fi29, fi29}
                                    _ = v41
                                case "foo bar":
                                    if v2 {
                                        var v42 float32 = (0.0 + (0.0 * 3.14159))
                                        if v2 {
                                            v43 := []int{fi29, int(v3), fi29}
                                            var v44 uint64 = v1
                                            v45 := v4[42]
                                            _ = v43
                                            _ = v44
                                            _ = v45
                                        }
                                        _ = v42
                                    } else {
                                        var v46 int64 = 440
                                        v47 := GenS2{F0: false, F1: 2147483647}
                                        _ = v46
                                        _ = v47
                                    }
                                case "hello":
                                    v48 := []int{fi29, 945, (fi10 - 32768)}
                                    _ = v48
                                default:
                                    switch "" {
                                    case "\n\t":
                                        var v49 uint64 = (v1 % (((v1 << ((255) & 63))) | 1))
                                        _ = v49
                                    case "hello":
                                        var v50 uint = 18446744073709551615
                                        _ = v50
                                    case "a":
                                        v3 = v3
                                    }
                                }
                                v51 := []int{fi10}
                                if true {
                                    var v52 float64 = 0.0
                                    v53 := GenS1{F0: 0, F1: v2, F2: 0}
                                    var v54 float32 = ((-2.25 + 0.1) - (3.14159 + float32((-2.25 * v52))))
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                }
                                _ = v51
                            }
                            var v55 float64 = 1.5
                            v4[0] = fi10
                            _ = fi29
                            _ = v55
                        }
                        _ = v28
                    }
                } else {
                    if ((((v2 && v2) || v2) && false) && (fi10 < fi10)) {
                        v56 := []int{9223372036854775807}
                        v1 = v1
                        v3 = 128
                        _ = v56
                    }
                }
            default:
                switch 42 {
                case -435:
                    for fi57 := 0; fi57 < 5; fi57++ {
                        v1 = uint64(v3)
                        _ = fi57
                    }
                default:
                    if v2 {
                        if (2147483648 == 598) {
                            if true {
                                v58 := GenS4{F0: 1}
                                v59 := GenS5{F0: (!(!(489 != 0))), F1: 57}
                                v60 := GenS6{F0: 2147483647, F1: 240, F2: 0}
                                _ = v58
                                _ = v59
                                _ = v60
                            } else {
                                switch ((("foo bar" + (("foo bar" + ("" + "")) + ("foo bar" + ("foo bar" + "hello")))) + "") + "foo bar") {
                                case "":
                                    v61 := GenS5{F0: false, F1: v1}
                                    _ = v61
                                case "foo bar":
                                    var v62 uint8 = v3
                                    _ = v62
                                case "hello":
                                    fmt.Println(fi10)
                                }
                                var v63 int8 = 19
                                v3 += v3
                                _ = v63
                            }
                            v1 += (v1 ^ (^(v1)))
                        } else {
                            v2 = (fi10 >= (fi10 | (fi10 | fi10)))
                            delete(v4, 2)
                            v64 := GenS7{F0: (v2 && v2), F1: 65535}
                            _ = v64
                        }
                        v3 -= v3
                    } else {
                        if ((1e10 / ((1e10 - float64(0.1)) + (0.1 / -2.25))) == 1.5) {
                            switch "foo bar" {
                            case "\n\t":
                                v2 = true
                            case "":
                                fmt.Println(v1, v2)
                            case "foo bar":
                                var v65 uint = 0
                                _ = v65
                            default:
                                if (673 <= -32768) {
                                    switch ((("hello" + "foo bar") + "\n\t") + "\n\t") {
                                    case "foo bar":
                                        fmt.Println(v1)
                                    }
                                    if (v2 || (("foo bar" + "") >= ("" + ("hello" + "hello")))) {
                                        var v66 uint32 = 65535
                                        _ = v66
                                    }
                                    var v67 uint = 2
                                    _ = v67
                                }
                            }
                            switch "foo bar" {
                            case "":
                                v68 := []int{fi10, fi10, int(v3), fi10}
                                _ = v68
                            case "\n\t":
                                v1 += v1
                            }
                            var v69 string = "a"
                            _ = v69
                        } else {
                            var v70 uint = 4294967295
                            _ = v70
                        }
                    }
                }
            }
            v4[698] = 1
        }
        _ = fi10
    }
    switch "foo bar" {
    case "":
        if (6 >= 4294967295) {
            for fi71 := 0; fi71 < 4; fi71++ {
                var v72 float64 = 0.1
                _ = fi71
                _ = v72
            }
        } else {
            fmt.Println(v2, v2, v3)
            v1 -= (v1 ^ (v1 ^ (v1 | (v1 + 0))))
            fmt.Println(v3, v2, v2)
        }
    case "\n\t":
        v2 = v2
    case "a":
        var v73 float32 = 0.1
        _ = v73
    default:
        if (-1 >= 1000) {
            for fi74 := 0; fi74 < 20; fi74++ {
                var v75 float32 = 0.0
                var v76 bool = v2
                switch "a" {
                case "a":
                    v77 := GenS8{F0: -117, F1: uint64(v3)}
                    _ = v77
                case "hello":
                    v78 := GenS2{F0: v2, F1: -949}
                    _ = v78
                case "":
                    if v2 {
                        if (v3 >= v3) {
                            v3 = uint8(v1)
                            var v79 float32 = (0.0 / v75)
                            v3 += v3
                            _ = v79
                        } else {
                            for fi80 := 0; fi80 < 16; fi80++ {
                                var v81 float64 = 1e10
                                _ = fi80
                                _ = v81
                            }
                        }
                    }
                default:
                    for fi82 := 0; fi82 < 17; fi82++ {
                        v2 = (!v2)
                        _ = fi82
                    }
                }
                _ = fi74
                _ = v75
                _ = v76
            }
        } else {
            var v83 int32 = 0
            v84 := []int{}
            var v85 uint16 = 0
            _ = v83
            _ = v84
            _ = v85
        }
    }
    if (100 >= -37) {
        v86 := GenS6{F0: 0, F1: 522, F2: -1}
        v87 := v4[1]
        _ = v86
        _ = v87
    }
    var v88 uint = 0
    for fi89 := 0; fi89 < 6; fi89++ {
        v90 := []int{(fi89 / (((fi89 * -9223372036854775808)) | 1)), fi89}
        v90 = append(v90, fi89)
        v1 = 0
        _ = fi89
        _ = v90
    }
    if (!(42 != 376)) {
        v2 = (9223372036854775807 >= 127)
        if v2 {
            switch 2147483647 {
            case -9223372036854775808:
                if v2 {
                    v91 := GenS2{F0: (!((^(v1)) >= 4294967295)), F1: -2147483648}
                    var v92 int64 = 1
                    v88 -= (v88 + 4294967295)
                    _ = v91
                    _ = v92
                } else {
                    v88 = (v88 % ((42) | 1))
                }
            default:
                var v93 string = ("foo bar" + "\n\t")
                _ = v93
            }
        } else {
            v1 = 2
            var v94 int8 = -128
            v95 := GenS9{F0: 902, F1: "\n\t"}
            _ = v94
            _ = v95
        }
        v1 = 354
    } else {
        v96 := GenS10{F0: 2147483647, F1: "a", F2: 9223372036854775807}
        v97 := GenS11{F0: 21, F1: (v88 * v88)}
        var v98 uint64 = (v1 * 1)
        _ = v96
        _ = v97
        _ = v98
    }
    v99 := GenS12{F0: (536 <= -2147483648), F1: 128}
    var v100 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v88:", v88)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v88
    _ = v99
}
