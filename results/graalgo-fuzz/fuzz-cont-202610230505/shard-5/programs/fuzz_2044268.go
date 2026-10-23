package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint8; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint; F1 int8 }
func genFn0(p0 int16, p1 uint16) int {
    switch 39 {
    case -1:
        p0 += p0
    default:
        var v6 int32 = 1
        _ = v6
    }
    p1 = 2
    var v7 int8 = 127
    _ = p0
    _ = p1
    _ = v7
    return 127
}
func genFn1(p0 uint16) int32 {
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := GenS1{F0: 837}
            switch (-(fi8)) {
            case 966:
                p0 = (p0 % (((p0 | 900)) | 1))
            case 0:
                var v11 int8 = 83
                _ = v11
            case 4294967295:
                var v12 int = (fi8 ^ (fi9 % ((-635) | 1)))
                _ = v12
            default:
                p0 = 255
            }
            p0 -= 460
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v13 uint8 = 127
    _ = p0
    _ = v13
    return -2147483648
}

func main() {
    var v1 uint32 = 40
    var v2 float64 = -2.25
    var v3 uint64 = 42
    var v4 bool = false
    var v5 uint8 = 162
    v3 = v3
    if (42 > v5) {
        switch "foo bar" {
        case "a":
            var v14 uint32 = v1
            _ = v14
        }
        var v15 float32 = float32(v2)
        fmt.Println(v1)
        _ = v15
    } else {
        if true {
            var v16 int = -1
            _ = v16
        }
    }
    if ((true || ((v4 && (1.5 < float32(float64(3.14159)))) || v4)) && ((536 < 582) || v4)) {
        v17 := GenS1{F0: 0}
        for fi18 := 0; fi18 < 16; fi18++ {
            for fi19 := 0; fi19 < 13; fi19++ {
                v20 := []int{}
                _ = fi19
                _ = v20
            }
            if (-1 <= -798) {
                for fi21 := 0; fi21 < 6; fi21++ {
                    v4 = v4
                    _ = fi21
                }
            } else {
                v22 := GenS1{F0: 2}
                _ = v22
            }
            fmt.Println(v2)
            _ = fi18
        }
        v23 := []int{-2147483648}
        _ = v17
        _ = v23
    } else {
        if false {
            for fi24 := 0; fi24 < 6; fi24++ {
                for fi25 := 0; fi25 < 8; fi25++ {
                    fmt.Println(fi24, v2)
                    _ = fi25
                }
                var v26 int8 = -37
                v27 := []int{}
                _ = fi24
                _ = v26
                _ = v27
            }
            fmt.Println(v2)
            v1 = (v1 + (v1 >> ((643) & 63)))
        } else {
            if v4 {
                v28 := GenS2{F0: (!v4), F1: v5, F2: (("foo bar" + "") + "\n\t")}
                v29 := GenS1{F0: 42}
                _ = v28
                _ = v29
            } else {
                v2 += v2
            }
            if v4 {
                if v4 {
                    var v30 int32 = -1
                    _ = v30
                }
                v31 := GenS1{F0: 584}
                _ = v31
            }
            v32 := GenS3{F0: "\n\t"}
            _ = v32
        }
        if (0 > 0) {
            v4 = v4
            v33 := []int{-248}
            _ = v33
        } else {
            v2 = ((v2 - 0.1) - (float64(float32(-2.25)) / v2))
            fmt.Println(v3, v1, v1)
        }
        var v34 uint8 = v5
        _ = v34
    }
    v1 += v1
    for fi35 := 0; fi35 < 19; fi35++ {
        v36 := []int{fi35, -204, -2147483648}
        v36 = append(v36, fi35)
        var v37 string = ""
        _ = fi35
        _ = v36
        _ = v37
    }
    if (!(!false)) {
        v38 := GenS4{F0: 0, F1: -1}
        fmt.Println(v4, v4, v3)
        _ = v38
    } else {
        var v39 uint32 = v1
        if false {
            v39 = uint32(v5)
            switch 42 {
            case 127:
                v40 := []int{}
                _ = v40
            case -39:
                v41 := GenS5{F0: 65535, F1: 1}
                _ = v41
            }
            v3 = v3
        } else {
            v42 := GenS4{F0: -9223372036854775808, F1: 745}
            var v43 int = 679
            for fi44 := 0; fi44 < 3; fi44++ {
                var v45 float64 = v2
                for fi46 := 0; fi46 < 14; fi46++ {
                    var v47 int32 = 0
                    var v48 int32 = (v47 & (v47 / ((v47) | 1)))
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                v49 := []int{}
                _ = fi44
                _ = v45
                _ = v49
            }
            _ = v42
            _ = v43
        }
        _ = v39
    }
    v5 = v5
    fmt.Println(v4)
    v1 += (v1 / ((v1) | 1))
    var v50 int16 = 42
    var v51 uint64 = (v3 | v3)
    switch "a" {
    case "hello":
        var v52 int16 = -1
        _ = v52
    default:
        if (-1 != 1) {
            var v53 uint16 = 32767
            switch 1 {
            case 128:
                if v4 {
                    if (v3 > v51) {
                        fmt.Println(v50, v53)
                    } else {
                        v2 += 1e10
                        var v54 int = -480
                        _ = v54
                    }
                    v4 = (true || (v4 || (1 < 42)))
                } else {
                    fmt.Println(v5, v1)
                    for fi55 := 0; fi55 < 9; fi55++ {
                        if (false && false) {
                            v51 = (^(v3))
                            for fi56 := 0; fi56 < 6; fi56++ {
                                if ((v5 * (^(v5))) > 255) {
                                    switch "\n\t" {
                                    case "a":
                                        if v4 {
                                            var v57 int8 = 100
                                            fmt.Println(v4, v3)
                                            _ = v57
                                        } else {
                                            v58 := GenS3{F0: "a"}
                                            fmt.Println(fi56, v5, v1)
                                            _ = v58
                                        }
                                    case "hello":
                                        if ((^(v3)) == v51) {
                                            v51 = 267
                                        }
                                        v59 := []int{-995, (fi56 % ((fi56) | 1)), int(v50), fi56}
                                        if false {
                                            var v60 int = (fi55 | (^(fi56)))
                                            _ = v60
                                        }
                                        v1 -= (v1 + 2)
                                        _ = v59
                                    }
                                    var v61 uint16 = (v53 + v53)
                                    v62 := []int{}
                                    _ = v61
                                    _ = v62
                                }
                                _ = fi56
                            }
                        } else {
                            switch "hello" {
                            case "\n\t":
                                v63 := []int{fi55, int(v50), (^(fi55)), genFn0((v50 * 42), 0)}
                                _ = v63
                            default:
                                switch "hello" {
                                case "hello":
                                    v64 := []int{(fi55 >> ((18446744073709551615) & 63)), 1, fi55, fi55}
                                    _ = v64
                                case "":
                                    if (v4 && v4) {
                                        if v4 {
                                            var v65 bool = v4
                                            for fi66 := 0; fi66 < 8; fi66++ {
                                                switch "hello" {
                                                case "a":
                                                    var v67 bool = v65
                                                    _ = v67
                                                }
                                                _ = fi66
                                            }
                                            var v68 int32 = int32(v5)
                                            _ = v65
                                            _ = v68
                                        } else {
                                            v69 := []int{genFn0((^(v50)), 0), fi55, genFn0((v50 & 0), (v53 % ((v53) | 1)))}
                                            if len(v69) > 0 {
                                                v69[0] = 32768
                                            }
                                            fmt.Println(v1)
                                            _ = v69
                                        }
                                        var v70 uint64 = (v51 / (((^(v51))) | 1))
                                        _ = v70
                                    } else {
                                        v71 := GenS6{F0: 0, F1: -38}
                                        switch "a" {
                                        case "foo bar":
                                            v72 := GenS4{F0: 605, F1: 894}
                                            _ = v72
                                        case "hello":
                                            if (34 <= v53) {
                                                var v73 int = (fi55 - (fi55 | genFn0(1, (v53 - v53))))
                                                for fi74 := 0; fi74 < 8; fi74++ {
                                                    v75 := GenS2{F0: v4, F1: v5, F2: "hello"}
                                                    v76 := GenS4{F0: -1, F1: -1}
                                                    var v77 uint32 = v1
                                                    _ = fi74
                                                    _ = v75
                                                    _ = v76
                                                    _ = v77
                                                }
                                                for fi78 := 0; fi78 < 18; fi78++ {
                                                    if v4 {
                                                        var v79 uint8 = 0
                                                        _ = v79
                                                    }
                                                    fmt.Println(v2, v51)
                                                    v80 := GenS5{F0: v53, F1: v53}
                                                    if (1 > 0) {
                                                        v81 := GenS6{F0: 4294967295, F1: -1}
                                                        v51 += (v51 * v3)
                                                        _ = v81
                                                    }
                                                    v82 := []int{(-(fi55)), -1, fi55, 32768}
                                                    v53 -= 0
                                                    _ = fi78
                                                    _ = v80
                                                    _ = v82
                                                }
                                                _ = v73
                                            }
                                        case "\n\t":
                                            v2 -= (v2 / ((v2 + v2) * float64(0.1)))
                                        }
                                        _ = v71
                                    }
                                }
                            }
                            v83 := GenS3{F0: "hello"}
                            for fi84 := 0; fi84 < 15; fi84++ {
                                v51 = 916
                                switch fi55 {
                                case 128:
                                    fmt.Println(v51, v50)
                                case 885:
                                    var v85 int8 = -128
                                    v86 := GenS5{F0: v53, F1: 2}
                                    if ((v50 / (((v50 >> ((71) & 63))) | 1)) != int16(v5)) {
                                        v87 := GenS6{F0: 434, F1: v85}
                                        fmt.Println(v50, v3)
                                        _ = v87
                                    } else {
                                        if ((v3 > (v3 ^ (v3 + v3))) || v4) {
                                            if v4 {
                                                var v88 int = fi55
                                                _ = v88
                                            }
                                        }
                                    }
                                    _ = v85
                                    _ = v86
                                default:
                                    var v89 bool = (18446744073709551615 < 1)
                                    _ = v89
                                }
                                _ = fi84
                            }
                            _ = v83
                        }
                        v90 := GenS3{F0: "\n\t"}
                        v91 := []int{}
                        _ = fi55
                        _ = v90
                        _ = v91
                    }
                }
            case -1:
                var v92 float64 = v2
                _ = v92
            default:
                if (!(v3 != (v3 ^ v3))) {
                    var v93 uint32 = v1
                    _ = v93
                } else {
                    if ((v53 ^ v53) != (v53 % (((^(v53))) | 1))) {
                        v94 := []int{-556}
                        _ = v94
                    } else {
                        v95 := []int{}
                        var v96 int32 = 0
                        v97 := GenS2{F0: false, F1: uint8(v3), F2: "hello"}
                        _ = v95
                        _ = v96
                        _ = v97
                    }
                    switch "hello" {
                    case "a":
                        var v98 int32 = -678
                        _ = v98
                    case "\n\t":
                        var v99 uint64 = 18446744073709551615
                        _ = v99
                    default:
                        var v100 uint16 = 738
                        _ = v100
                    }
                }
            }
            _ = v53
        } else {
            var v101 int64 = -9223372036854775808
            if (float64(3.14159) <= (((v2 - (v2 + v2)) - (((v2 / v2) + v2) * 0.0)) - 0.1)) {
                for fi102 := 0; fi102 < 19; fi102++ {
                    for fi103 := 0; fi103 < 18; fi103++ {
                        switch "hello" {
                        case "":
                            fmt.Println(v101, v51)
                        case "a":
                            fmt.Println(v50)
                        case "\n\t":
                            var v104 bool = (42 != 255)
                            _ = v104
                        default:
                            v50 = int16(v101)
                        }
                        _ = fi103
                    }
                    v105 := GenS1{F0: 861}
                    _ = fi102
                    _ = v105
                }
            } else {
                v106 := []int{4294967295, 886}
                _ = v106
            }
            v51 += v51
            _ = v101
        }
    }
    v107 := []int{-660}
    for fi108 := 0; fi108 < 4; fi108++ {
        v51 = (v3 << ((45) & 63))
        v5 += v5
        switch "a" {
        case "":
            var v109 string = ("\n\t" + "")
            _ = v109
        case "a":
            if v4 {
                fmt.Println(v4, v51, v2)
            } else {
                var v110 int64 = 42
                var v111 int8 = -11
                _ = v110
                _ = v111
            }
        default:
            v3 = (v51 >> ((uint(v1)) & 63))
        }
        _ = fi108
    }
    _ = len(v107)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v107:", len(v107), v107)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
}
