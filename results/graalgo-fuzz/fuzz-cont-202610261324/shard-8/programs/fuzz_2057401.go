package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 int8; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 int16; F1 int64; F2 uint16 }
type GenS7 struct { F0 uint; F1 int64; F2 uint8 }
type GenS8 struct { F0 string; F1 uint8; F2 string }
func genFn0(p0 string, p1 uint32, p2 int64) string {
    var v5 bool = true
    if true {
        fmt.Println(v5, p0, p0)
        if v5 {
            v6 := []int{32767, 0, 32767, 128}
            for fi7 := 0; fi7 < 10; fi7++ {
                v5 = true
                _ = fi7
            }
            v6 = append(v6, 32768)
            _ = v6
        }
    } else {
        var v8 bool = v5
        v9 := GenS1{F0: 1}
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := GenS2{F0: p0, F1: 1, F2: 42}
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 uint32 = (p1 + p1)
                v14 := []int{}
                _ = fi12
                _ = v13
                _ = v14
            }
            switch (p0 + p0) {
            case "foo bar":
                v15 := GenS1{F0: -2147483648}
                _ = v15
            case "a":
                var v16 int32 = -1
                _ = v16
            case "":
                var v17 bool = (false || true)
                _ = v17
            default:
                for fi18 := 0; fi18 < 14; fi18++ {
                    if v5 {
                        if (v8 || false) {
                            var v19 int32 = -2147483648
                            fmt.Println(fi10, v8, p0)
                            var v20 uint32 = p1
                            _ = v19
                            _ = v20
                        } else {
                            var v21 uint64 = 42
                            switch p0 {
                            case "a":
                                p0 = p0
                            default:
                                var v22 uint = 1
                                _ = v22
                            }
                            var v23 int16 = 688
                            _ = v21
                            _ = v23
                        }
                        fmt.Println(fi10)
                        var v24 int16 = 0
                        _ = v24
                    } else {
                        v25 := GenS3{F0: 421}
                        _ = v25
                    }
                    v26 := []int{(fi18 ^ 2147483647)}
                    if len(v26) > 0 {
                        v26[0] = (fi10 >> ((18446744073709551615) & 63))
                    }
                    if len(v26) > 0 {
                        v26[0] = (fi10 * (fi10 % ((fi10) | 1)))
                    }
                    v27 := GenS4{F0: 209}
                    if len(v26) > 0 {
                        v26[0] = fi10
                    }
                    var v28 bool = ((!(42 > 1)) && (((v8 || v8) || (int8(fi18) < 13)) || true))
                    var v29 int32 = 42
                    switch ((p0 + p0) + p0) {
                    case "a":
                        v29 -= (v29 * (^(v29)))
                    case "\n\t":
                        switch p0 {
                        case "\n\t":
                            if len(v26) > 0 {
                                v26[0] = (fi10 % ((fi18) | 1))
                            }
                        case "hello":
                            p2 -= 0
                        case "a":
                            v30 := GenS3{F0: 32767}
                            _ = v30
                        default:
                            if (((((float64(1e10) - ((float64(1e10) + (0.1 / 0.1)) * 3.14159)) * 3.14159) / (1e10 / (0.1 - 1e10))) >= 1.5) || (74 <= 18446744073709551615)) {
                                var v31 int = (fi10 ^ 32768)
                                fmt.Println(v29, fi18)
                                var v32 int32 = (v29 >> ((0) & 63))
                                _ = v31
                                _ = v32
                            }
                        }
                    }
                    if v28 {
                        var v33 uint16 = 1
                        var v34 uint8 = 96
                        _ = v33
                        _ = v34
                    }
                    switch 408 {
                    case 2147483647:
                        var v35 uint16 = 654
                        _ = v35
                    case 32768:
                        if (42 != 0) {
                            v26 = append(v26, fi10)
                            if true {
                                if len(v26) > 0 {
                                    v26[0] = -1
                                }
                            }
                        } else {
                            p1 = p1
                            fmt.Println(v28, p0, v5)
                            fmt.Println(v29, v29, fi10)
                        }
                    case -562:
                        var v36 float64 = 1e10
                        _ = v36
                    default:
                        if (446 < 114) {
                            var v37 int64 = (p2 | (p2 | p2))
                            _ = v37
                        } else {
                            if true {
                                if ((!(548 > 2)) || v5) {
                                    var v38 string = ("" + (p0 + "foo bar"))
                                    _ = v38
                                } else {
                                    if true {
                                        var v39 float64 = float64(0.0)
                                        var v40 int8 = 42
                                        _ = v39
                                        _ = v40
                                    }
                                    p0 = p0
                                    fmt.Println(p1, v29, fi18)
                                }
                            } else {
                                var v41 uint16 = 821
                                v42 := GenS5{F0: 530, F1: 114}
                                v8 = (!v8)
                                _ = v41
                                _ = v42
                            }
                        }
                    }
                    var v43 int32 = (v29 ^ (v29 & (v29 & (-(v29)))))
                    _ = fi18
                    _ = v26
                    _ = v27
                    _ = v28
                    _ = v29
                    _ = v43
                }
            }
            _ = fi10
            _ = v11
        }
        _ = v8
        _ = v9
    }
    var v44 int16 = -32768
    for fi45 := 0; fi45 < 14; fi45++ {
        if (p1 == p1) {
            for fi46 := 0; fi46 < 18; fi46++ {
                v47 := GenS6{F0: -1, F1: p2, F2: 65535}
                if v5 {
                    v5 = false
                    p1 -= 2147483647
                    v5 = (!false)
                }
                if (128 >= 127) {
                    var v48 uint32 = p1
                    _ = v48
                } else {
                    v49 := []int{fi45, (fi45 % ((fi45) | 1))}
                    var v50 uint8 = 2
                    _ = v49
                    _ = v50
                }
                var v51 uint32 = p1
                p1 += p1
                _ = fi46
                _ = v47
                _ = v51
            }
            v52 := []int{(fi45 << ((255) & 63)), (fi45 & fi45), (fi45 & fi45), fi45}
            _ = v52
        } else {
            var v53 int64 = p2
            if (!((96 == 42) && false)) {
                for fi54 := 0; fi54 < 4; fi54++ {
                    var v55 uint16 = 475
                    p0 = ((p0 + p0) + ((p0 + p0) + p0))
                    _ = fi54
                    _ = v55
                }
                if false {
                    v56 := []int{(fi45 | (fi45 & 775)), fi45, (fi45 >> ((0) & 63)), 127}
                    var v57 uint = 470
                    _ = v56
                    _ = v57
                }
            } else {
                switch (p0 + (p0 + p0)) {
                case "\n\t":
                    switch fi45 {
                    case 298:
                        var v58 int16 = v44
                        _ = v58
                    case 127:
                        if v5 {
                            if (v5 || v5) {
                                if v5 {
                                    v59 := GenS1{F0: 940}
                                    for fi60 := 0; fi60 < 20; fi60++ {
                                        if v5 {
                                            var v61 uint64 = 2
                                            v62 := GenS4{F0: (v61 << ((255) & 63))}
                                            _ = v61
                                            _ = v62
                                        } else {
                                            v63 := []int{-9223372036854775808, (fi60 + fi45)}
                                            if (((-2.25 / 1e10) <= 1.5) && (!v5)) {
                                                v64 := GenS6{F0: -32768, F1: (p2 + (v53 << ((42) & 63))), F2: 265}
                                                switch (^(fi60)) {
                                                case 1:
                                                    if len(v63) > 0 {
                                                        v63[0] = 128
                                                    }
                                                case 128:
                                                    var v65 int = fi45
                                                    _ = v65
                                                }
                                                _ = v64
                                            } else {
                                                fmt.Println(p2, p1)
                                            }
                                            _ = v63
                                        }
                                        _ = fi60
                                    }
                                    _ = v59
                                }
                                v66 := []int{}
                                _ = v66
                            }
                        }
                    default:
                        fmt.Println(p2)
                    }
                case "hello":
                    v67 := []int{fi45, fi45, fi45}
                    _ = v67
                }
                var v68 float32 = 0.1
                fmt.Println(p0, v53, v68)
                _ = v68
            }
            _ = v53
        }
        v69 := []int{fi45, (fi45 >> ((888) & 63))}
        _ = fi45
        _ = v69
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v44
    _ = v5
    return p0
}
func genFn1(p0 int32) uint64 {
    var v70 float64 = 1.5
    for fi71 := 0; fi71 < 12; fi71++ {
        for fi72 := 0; fi72 < 8; fi72++ {
            v70 += v70
            _ = fi72
        }
        v70 = ((0.0 + v70) / 1.5)
        v73 := []int{924, fi71, (fi71 % ((fi71) | 1)), (fi71 + (fi71 * (fi71 << ((0) & 63))))}
        _ = fi71
        _ = v73
    }
    _ = p0
    _ = v70
    return 4294967295
}

func main() {
    var v1 int8 = -30
    var v2 int = -2147483648
    v3 := make(map[string]bool)
    var v4 int32 = 0
    if (!(false || (false && (false || true)))) {
        fmt.Println(v2, v1)
    }
    if ((0.1 + (1e10 + (1.5 / 1e10))) != -2.25) {
        _ = len(v3)
        v1 -= (v1 % ((v1) | 1))
        v4 = v4
    } else {
        var v74 int = 708
        _ = v74
    }
    v4 = (v4 << ((67) & 63))
    for fi75 := 0; fi75 < 7; fi75++ {
        for fi76 := 0; fi76 < 5; fi76++ {
            _ = len(v3)
            v2 = (v2 ^ v2)
            var v77 bool = (4294967295 <= 2)
            _ = fi76
            _ = v77
        }
        v1 -= (v1 << ((839) & 63))
        v78 := []int{1, 276, (v2 ^ (v2 + (-(fi75))))}
        _ = fi75
        _ = v78
    }
    switch "a" {
    case "":
        switch v2 {
        case 9223372036854775807:
            var v79 float32 = (float32(float64((1.5 + 3.14159))) / float32(((((0.1 * 3.14159) - 1.5) * 0.1) / 1e10)))
            _ = v79
        case 645:
            v1 = (v1 ^ (v1 % (((-(v1))) | 1)))
        default:
            var v80 float64 = (float64((1e10 + ((0.1 * 3.14159) * -2.25))) - float64(((0.0 + (0.1 / float32(1.5))) / 1.0)))
            _ = v80
        }
    case "hello":
        var v81 float32 = (0.1 - 3.14159)
        _ = v81
    default:
        var v82 int8 = (v1 + v1)
        _ = v82
    }
    var v83 uint64 = genFn1(v4)
    var v84 uint = 2
    v84 += (v84 * (v84 ^ (v84 - v84)))
    if (188 == v83) {
        var v85 uint8 = 127
        var v86 uint16 = 255
        switch 9223372036854775807 {
        case -1:
            var v87 int = (-(v2))
            _ = v87
        default:
            var v88 uint8 = v85
            _ = v88
        }
        _ = v85
        _ = v86
    } else {
        for fi89 := 0; fi89 < 8; fi89++ {
            for fi90 := 0; fi90 < 8; fi90++ {
                v2 -= fi89
                var v91 uint64 = v83
                _ = fi90
                _ = v91
            }
            v4 = 0
            if false {
                v92 := GenS7{F0: 285, F1: -9223372036854775808, F2: 127}
                if true {
                    var v93 float64 = 1e10
                    _ = v93
                }
                var v94 int8 = (v1 % (((v1 - v1)) | 1))
                _ = v92
                _ = v94
            }
            _ = fi89
        }
    }
    v1 -= v1
    v4 = (-(v4))
    fmt.Println(v1, v1, v4)
    for fi95 := 0; fi95 < 4; fi95++ {
        v1 = v1
        _ = fi95
    }
    switch (v2 * v2) {
    case 1:
        if (-2.25 > 1e10) {
            var v96 string = "hello"
            var v97 float32 = float32(((1e10 - 0.0) / 0.1))
            for fi98 := 0; fi98 < 8; fi98++ {
                v4 -= (v4 - (v4 * (v4 | (v4 / ((86) | 1)))))
                v96 = "foo bar"
                switch (v2 % ((v2) | 1)) {
                case 128:
                    var v99 int8 = v1
                    _ = v99
                case 633:
                    v100 := GenS8{F0: ("\n\t" + (v96 + (v96 + v96))), F1: 140, F2: (("" + v96) + ("hello" + ""))}
                    _ = v100
                default:
                    if (!(-2.25 > (0.0 / 0.1))) {
                        v3[v96] = (65535 >= uint16(v83))
                        if true {
                            var v101 uint32 = 2
                            v102 := GenS5{F0: v83, F1: (v1 | v1)}
                            v97 = float32(((0.1 / (float64(v97) / 0.1)) / -2.25))
                            _ = v101
                            _ = v102
                        } else {
                            v103 := GenS3{F0: -674}
                            _ = v103
                        }
                    }
                }
                _ = fi98
            }
            _ = v96
            _ = v97
        }
    case 42:
        v1 = 0
    default:
        fmt.Println(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v83
    _ = v84
}
