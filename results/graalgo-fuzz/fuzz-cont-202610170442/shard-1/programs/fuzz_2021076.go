package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int32; F1 bool; F2 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS8 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int16; F1 int16; F2 int8 }
func genFn0(p0 bool) uint64 {
    if true {
        switch -122 {
        case 42:
            fmt.Println(p0)
        case 4294967295:
            v7 := GenS1{F0: 0, F1: 1000}
            _ = v7
        case 32768:
            v8 := GenS2{F0: 2147483647, F1: (-9223372036854775808 < 0), F2: -58}
            _ = v8
        default:
            var v9 bool = false
            _ = v9
        }
        if ((!p0) && (p0 || p0)) {
            p0 = p0
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := GenS3{F0: 2}
                if (52 > 0) {
                    fmt.Println(fi10)
                    for fi12 := 0; fi12 < 13; fi12++ {
                        v13 := []int{fi10, fi12, (fi10 / (((-(fi10))) | 1)), fi12}
                        _ = fi12
                        _ = v13
                    }
                } else {
                    var v14 float64 = ((-2.25 * -2.25) + float64(float32(1e10)))
                    var v15 string = "a"
                    if (!true) {
                        v16 := GenS4{F0: 0, F1: 255, F2: int64(fi10)}
                        _ = v16
                    } else {
                        var v17 int16 = 942
                        var v18 string = "\n\t"
                        v19 := []int{fi10, fi10}
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = v14
                    _ = v15
                }
                _ = fi10
                _ = v11
            }
        } else {
            var v20 int = 32767
            v21 := GenS2{F0: 1, F1: (p0 || (32767 != 32767)), F2: -102}
            _ = v20
            _ = v21
        }
    } else {
        for fi22 := 0; fi22 < 19; fi22++ {
            fmt.Println(p0, fi22)
            for fi23 := 0; fi23 < 5; fi23++ {
                var v24 uint64 = 588
                var v25 uint = 255
                if (!true) {
                    v26 := GenS5{F0: (v25 | 2), F1: 255}
                    switch -29 {
                    case -2147483648:
                        var v27 uint32 = 42
                        _ = v27
                    case 4294967295:
                        var v28 uint64 = (v24 >> ((v25) & 63))
                        _ = v28
                    }
                    fmt.Println(fi23, p0, v25)
                    v29 := GenS2{F0: -2147483648, F1: (false || true), F2: 127}
                    _ = v26
                    _ = v29
                } else {
                    v30 := []int{661, (fi22 % (((fi23 ^ 735)) | 1)), (fi23 & fi23)}
                    switch fi22 {
                    case 32767:
                        v31 := GenS6{F0: v24}
                        _ = v31
                    case 987:
                        p0 = (((0.1 + (1.5 - (0.0 * 1e10))) + float32(0.0)) == 3.14159)
                    case 4294967295:
                        var v32 uint = (v25 | (v25 / ((2) | 1)))
                        _ = v32
                    default:
                        _ = len(v30)
                    }
                    v24 -= 42
                    if len(v30) > 0 {
                        v30[0] = (fi22 << ((42) & 63))
                    }
                    var v33 int32 = -2147483648
                    switch "" {
                    case "\n\t":
                        if ("" <= "a") {
                            v34 := GenS1{F0: 154, F1: -1}
                            p0 = (p0 || p0)
                            v35 := GenS1{F0: 1, F1: 0}
                            v30 = append(v30, (fi23 / ((32768) | 1)))
                            v36 := GenS4{F0: (v24 - 612), F1: uint16(v25), F2: -666}
                            fmt.Println(v33, v33, v25)
                            if len(v30) > 0 {
                                v30[0] = 1
                            }
                            switch fi23 {
                            case 1:
                                var v37 int = 1
                                _ = v37
                            case 127:
                                v25 -= v25
                            case 0:
                                p0 = true
                            default:
                                if (p0 || p0) {
                                    var v38 float64 = -2.25
                                    var v39 int = (-(fi22))
                                    _ = v38
                                    _ = v39
                                } else {
                                    fmt.Println(fi22, fi23)
                                    if true {
                                        if ("\n\t" < "foo bar") {
                                            var v40 uint32 = 983
                                            if p0 {
                                                var v41 int32 = v33
                                                switch -713 {
                                                case 955:
                                                    var v42 float64 = (0.1 + 0.1)
                                                    var v43 uint8 = 0
                                                    var v44 uint64 = v24
                                                    _ = v42
                                                    _ = v43
                                                    _ = v44
                                                }
                                                var v45 float64 = (1.5 / 1e10)
                                                _ = v41
                                                _ = v45
                                            } else {
                                                v25 -= (v25 << ((v25) & 63))
                                                _ = len(v30)
                                            }
                                            _ = v40
                                        }
                                        var v46 int64 = -1
                                        v33 = v33
                                        _ = v46
                                    }
                                }
                            }
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                    default:
                        fmt.Println(fi22, fi23)
                    }
                    _ = v30
                    _ = v33
                }
                var v47 uint32 = 0
                _ = fi23
                _ = v24
                _ = v25
                _ = v47
            }
            switch (fi22 << ((875) & 63)) {
            case 42:
                for fi48 := 0; fi48 < 5; fi48++ {
                    switch "" {
                    case "foo bar":
                        v49 := []int{fi22, (fi48 - (fi48 * fi22))}
                        _ = v49
                    case "a":
                        v50 := GenS3{F0: 17}
                        _ = v50
                    case "":
                        if true {
                            fmt.Println(fi48, fi22)
                            p0 = ((255 == 0) || p0)
                            p0 = false
                        } else {
                            v51 := []int{(fi48 * (fi48 * fi22))}
                            _ = v51
                        }
                    }
                    _ = fi48
                }
            default:
                var v52 int64 = 0
                _ = v52
            }
            _ = fi22
        }
        fmt.Println(p0)
        p0 = (-1 < 2147483647)
    }
    fmt.Println(p0)
    v53 := []int{42}
    v54 := GenS5{F0: 65535, F1: 65535}
    _ = p0
    _ = v53
    _ = v54
    return 481
}

func main() {
    v1 := []uint32{4294967295}
    var v2 uint64 = 0
    var v3 string = ("" + ("foo bar" + "\n\t"))
    var v4 uint8 = 126
    v5 := []int{32768}
    v6 := make(map[uint32]uint32)
    var v55 uint64 = 1
    fmt.Println(v55, v3, v4)
    var v56 bool = false
    v57 := GenS5{F0: 42, F1: 32767}
    v3 = "\n\t"
    if (!v56) {
        delete(v6, uint32(v2))
        v58 := GenS1{F0: v4, F1: -506}
        var v59 string = (v3 + (v3 + v3))
        _ = v58
        _ = v59
    }
    if v56 {
        v60 := v6[138]
        v61 := GenS1{F0: (v4 ^ v4), F1: -1}
        _ = v60
        _ = v61
    }
    var v62 int64 = -9223372036854775808
    if false {
        fmt.Println(v55)
        var v63 float64 = -2.25
        for fi64 := 0; fi64 < 9; fi64++ {
            switch fi64 {
            case 0:
                v65 := GenS7{F0: 255, F1: -32768, F2: genFn0(false)}
                _ = v65
            case -9223372036854775808:
                if (42 == 4294967295) {
                    if len(v5) > 0 {
                        v5[0] = fi64
                    }
                } else {
                    var v66 int16 = -846
                    _ = v66
                }
            case 32768:
                switch -1 {
                case -122:
                    v67 := v6[2147483648]
                    _ = v67
                case -2147483648:
                    var v68 string = ((("foo bar" + "a") + "foo bar") + (((v3 + "\n\t") + ("foo bar" + ("" + ""))) + v3))
                    _ = v68
                case -592:
                    v69 := GenS2{F0: -1, F1: (0 != 2147483647), F2: -128}
                    _ = v69
                default:
                    for fi70 := 0; fi70 < 8; fi70++ {
                        v71 := GenS3{F0: v4}
                        v72 := GenS1{F0: (^(v4)), F1: 32767}
                        if (!(!v56)) {
                            var v73 int = (fi64 - (fi70 >> ((255) & 63)))
                            _ = v73
                        } else {
                            v74 := GenS3{F0: (v4 - v4)}
                            _ = v74
                        }
                        _ = fi70
                        _ = v71
                        _ = v72
                    }
                }
            }
            for fi75 := 0; fi75 < 17; fi75++ {
                if (int32(fi64) > 42) {
                    fmt.Println(v56, v63)
                    var v76 float32 = 0.0
                    _ = v76
                } else {
                    var v77 float64 = v63
                    if v56 {
                        switch "foo bar" {
                        case "a":
                            var v78 float32 = (0.0 * 0.0)
                            _ = v78
                        default:
                            v56 = (326 < uint32(v2))
                        }
                        v55 = (v2 - (v2 - (v2 & (v2 & (v2 & 4294967295)))))
                    } else {
                        var v79 uint32 = 860
                        _ = v79
                    }
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    _ = v77
                }
                v2 -= genFn0((fi75 != 665))
                _ = fi75
            }
            _ = len(v6)
            _ = fi64
        }
        _ = v63
    } else {
        var v80 uint32 = 0
        if v56 {
            if false {
                if (!v56) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    for fi81 := 0; fi81 < 11; fi81++ {
                        v82 := GenS8{F0: (v55 % ((v55) | 1)), F1: 0, F2: (v80 % ((v80) | 1))}
                        switch "hello" {
                        case "hello":
                            if false {
                                for fi83 := 0; fi83 < 9; fi83++ {
                                    var v84 int8 = 42
                                    var v85 uint32 = 2
                                    _ = fi83
                                    _ = v84
                                    _ = v85
                                }
                            }
                        case "\n\t":
                            if ((v4 | 17) > v4) {
                                var v86 int32 = -684
                                if false {
                                    for fi87 := 0; fi87 < 15; fi87++ {
                                        v62 = v62
                                        if (v56 && true) {
                                            var v88 uint16 = 1
                                            v89 := GenS9{F0: 674}
                                            _ = v88
                                            _ = v89
                                        } else {
                                            if ((v80 | (^(v80))) == (v80 / ((v80) | 1))) {
                                                var v90 float32 = (-2.25 + 0.0)
                                                _ = v90
                                            } else {
                                                v80 = (v80 * (^(v80)))
                                            }
                                        }
                                        var v91 uint16 = 65535
                                        _ = fi87
                                        _ = v91
                                    }
                                    v55 += (v55 >> ((42) & 63))
                                    for fi92 := 0; fi92 < 19; fi92++ {
                                        v80 -= v80
                                        _ = fi92
                                    }
                                } else {
                                    switch (fi81 % ((fi81) | 1)) {
                                    case -304:
                                        var v93 float64 = 0.0
                                        _ = v93
                                    case 32768:
                                        v4 = v4
                                    }
                                }
                                if len(v1) > 0 {
                                    v1[0] = uint32(v55)
                                }
                                _ = v86
                            }
                        }
                        delete(v6, (^(v80)))
                        _ = fi81
                        _ = v82
                    }
                    var v94 int8 = 127
                    _ = v94
                } else {
                    v2 = genFn0((0 <= 1))
                    var v95 int8 = 127
                    _ = v95
                }
                for fi96 := 0; fi96 < 2; fi96++ {
                    switch "\n\t" {
                    case "\n\t":
                        for fi97 := 0; fi97 < 1; fi97++ {
                            v3 = v3
                            var v98 int8 = -33
                            _ = fi97
                            _ = v98
                        }
                    default:
                        for fi99 := 0; fi99 < 3; fi99++ {
                            var v100 uint = 65535
                            _ = fi99
                            _ = v100
                        }
                    }
                    var v101 int8 = -128
                    for fi102 := 0; fi102 < 16; fi102++ {
                        v103 := GenS10{F0: -1, F1: 1000, F2: v101}
                        switch (v3 + "") {
                        case "hello":
                            v104 := GenS5{F0: uint(v80), F1: 255}
                            _ = v104
                        default:
                            delete(v6, (v80 >> ((58) & 63)))
                        }
                        _ = fi102
                        _ = v103
                    }
                    _ = fi96
                    _ = v101
                }
            } else {
                v105 := GenS1{F0: 1, F1: 0}
                var v106 int16 = -1
                v107 := GenS1{F0: (v4 - 1), F1: -32768}
                _ = v105
                _ = v106
                _ = v107
            }
        }
        for fi108 := 0; fi108 < 2; fi108++ {
            if (!(v56 && ((fi108 | (fi108 - 4294967295)) != fi108))) {
                v109 := GenS3{F0: (^(v4))}
                _ = v109
            } else {
                var v110 uint = uint(v80)
                _ = v110
            }
            _ = fi108
        }
        _ = v80
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v6
    _ = v62
}
