package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int16; F1 int64; F2 string }
type GenS5 struct { F0 uint64; F1 int32; F2 bool }
func genFn0(p0 int, p1 uint8, p2 string) bool {
    if false {
        v7 := GenS1{F0: p1, F1: 1, F2: -229}
        p1 = p1
        _ = v7
    }
    var v8 int32 = 42
    p1 = (p1 << ((0) & 63))
    if (2 > 1) {
        var v9 int32 = 42
        _ = v9
    } else {
        fmt.Println(p1)
    }
    if (1 >= 32767) {
        v10 := GenS2{F0: true, F1: (p1 + p1)}
        var v11 int16 = 1000
        fmt.Println(p1)
        _ = v10
        _ = v11
    } else {
        if true {
            var v12 bool = true
            _ = v12
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(fi13, p1)
            var v14 int16 = 32767
            switch p2 {
            case "a":
                p1 = 13
            }
            _ = fi13
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return true
}

func main() {
    var v1 string = ""
    var v2 int8 = 127
    var v3 bool = false
    v4 := make(map[uint32]string)
    var v5 float64 = 1.5
    v6 := make(map[uint32]bool)
    switch "" {
    case "foo bar":
        var v15 int = 42
        _ = v15
    case "":
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := GenS1{F0: 128, F1: 5, F2: 562}
            fmt.Println(v3, v2, v3)
            _ = fi16
            _ = v17
        }
    case "\n\t":
        if (883 > -1) {
            if (false && v3) {
                switch 127 {
                case 127:
                    fmt.Println(v1, v2, v3)
                default:
                    if v3 {
                        v3 = (1 > 1000)
                        var v18 int8 = (v2 & -1)
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 3; fi19++ {
                            if v3 {
                                var v20 uint = 976
                                v5 = (v5 / ((v5 - v5) / v5))
                                _ = v20
                            } else {
                                v21 := GenS3{F0: (fi19 % ((905) | 1)), F1: 127}
                                _ = v21
                            }
                            fmt.Println(v3, v3)
                            if (100 <= (v2 + (v2 >> ((153) & 63)))) {
                                v22 := []int{(fi19 % (((fi19 * fi19)) | 1)), fi19}
                                if v3 {
                                    for fi23 := 0; fi23 < 2; fi23++ {
                                        _ = len(v22)
                                        _ = fi23
                                    }
                                    v24 := GenS3{F0: (fi19 / (((fi19 - (fi19 << ((18446744073709551615) & 63)))) | 1)), F1: 2}
                                    switch (fi19 % ((4294967295) | 1)) {
                                    case 913:
                                        var v25 uint8 = 185
                                        _ = v25
                                    case 1:
                                        switch (fi19 * fi19) {
                                        case 2147483647:
                                            var v26 uint32 = 0
                                            _ = v26
                                        case 339:
                                            var v27 int = (fi19 & (^(fi19)))
                                            _ = v27
                                        default:
                                            v28 := GenS1{F0: 255, F1: 1, F2: -1}
                                            _ = v28
                                        }
                                    case 414:
                                        v29 := GenS1{F0: 127, F1: 232, F2: -240}
                                        _ = v29
                                    }
                                    _ = v24
                                } else {
                                    v5 = v5
                                    v30 := GenS1{F0: 42, F1: 2, F2: 1}
                                    var v31 uint8 = 2
                                    _ = v30
                                    _ = v31
                                }
                                var v32 uint8 = 2
                                _ = v22
                                _ = v32
                            }
                            _ = fi19
                        }
                        v33 := []int{-278, -9223372036854775808}
                        _ = len(v33)
                        _ = v33
                    }
                }
                var v34 uint16 = 472
                _ = v34
            }
            var v35 float64 = v5
            _ = v35
        } else {
            switch -2147483648 {
            case 615:
                var v36 int8 = v2
                _ = v36
            default:
                v37 := []int{-1, 32767, 42}
                _ = v37
            }
            v5 -= (float64(0.1) - v5)
            v38 := GenS4{F0: 0, F1: 966, F2: v1}
            _ = v38
        }
    }
    var v39 uint32 = 2147483647
    var v40 int16 = 302
    var v41 float32 = 0.1
    var v42 int16 = (-(v40))
    switch "hello" {
    case "foo bar":
        switch 9223372036854775807 {
        case 32768:
            v41 = v41
        case 9223372036854775807:
            v42 -= v40
        }
    case "a":
        v5 -= v5
    }
    var v43 uint32 = v39
    var v44 bool = v3
    v45 := []int{0, -843}
    if (v3 || (true || (!(!v3)))) {
        v1 = v1
        if (!false) {
            if len(v45) > 0 {
                v45[0] = -1
            }
            fmt.Println(v39, v39, v39)
        } else {
            for fi46 := 0; fi46 < 17; fi46++ {
                var v47 uint8 = 222
                if (971 <= 179) {
                    var v48 uint16 = 861
                    v49 := GenS1{F0: (v47 / ((v47) | 1)), F1: (v48 & uint16(v43)), F2: v42}
                    _ = v48
                    _ = v49
                }
                var v50 uint64 = 42
                _ = fi46
                _ = v47
                _ = v50
            }
            if v44 {
                if v3 {
                    for fi51 := 0; fi51 < 13; fi51++ {
                        var v52 int64 = 0
                        for fi53 := 0; fi53 < 1; fi53++ {
                            fmt.Println(v42, v5)
                            fmt.Println(v41, v3)
                            _ = fi53
                        }
                        switch fi51 {
                        case 4294967295:
                            if (false || (!(v3 && (319 != 502)))) {
                                switch ("" + v1) {
                                case "":
                                    if (1 <= 1) {
                                        var v54 uint = 0
                                        _ = v54
                                    }
                                default:
                                    for fi55 := 0; fi55 < 7; fi55++ {
                                        var v56 float32 = v41
                                        switch int(v2) {
                                        case 4294967295:
                                            var v57 uint8 = 214
                                            _ = v57
                                        case 2147483647:
                                            v45 = append(v45, fi51)
                                        case -9223372036854775808:
                                            if (42 >= -167) {
                                                v1 = "a"
                                                var v58 int32 = 209
                                                if ((^(fi51)) >= 128) {
                                                    v59 := v6[1]
                                                    if (v42 < 42) {
                                                        v60 := GenS2{F0: ((((1.5 / v56) / float32((((v5 / v5) * v5) - v5))) + (v41 / 1e10)) < (v41 + -2.25)), F1: 148}
                                                        v61 := GenS1{F0: 96, F1: 65535, F2: v42}
                                                        _ = v60
                                                        _ = v61
                                                    } else {
                                                        v62 := GenS1{F0: 177, F1: 32767, F2: (v40 * int16(v52))}
                                                        fmt.Println(v43, v44)
                                                        var v63 int8 = v2
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                    v64 := GenS5{F0: 702, F1: 42, F2: false}
                                                    v65 := GenS4{F0: v40, F1: v52, F2: v1}
                                                    v66 := GenS3{F0: 128, F1: 127}
                                                    if (0 <= v2) {
                                                        switch (fi55 - 182) {
                                                        case 506:
                                                            if true {
                                                                v42 = (v40 ^ v40)
                                                            } else {
                                                                var v67 int8 = v2
                                                                _ = v67
                                                            }
                                                        case -9223372036854775808:
                                                            var v68 uint64 = 4294967295
                                                            _ = v68
                                                        default:
                                                            v5 = float64((0.1 + float32(v5)))
                                                        }
                                                        v5 = v5
                                                        var v69 float64 = v5
                                                        _ = v69
                                                    } else {
                                                        v5 += float64((v56 * 3.14159))
                                                    }
                                                    switch (-(fi55)) {
                                                    case 127:
                                                        var v70 float64 = -2.25
                                                        _ = v70
                                                    }
                                                    _ = v59
                                                    _ = v64
                                                    _ = v65
                                                    _ = v66
                                                } else {
                                                    v6[v43] = v44
                                                    var v71 bool = true
                                                    _ = v71
                                                }
                                                _ = v58
                                            } else {
                                                v6[v43] = (42 == 367)
                                                var v72 int16 = v42
                                                fmt.Println(v42, fi55)
                                                v3 = false
                                                _ = v72
                                            }
                                        default:
                                            if ((v2 | 127) != (v2 - v2)) {
                                                var v73 uint32 = v43
                                                _ = v73
                                            }
                                        }
                                        v4[0] = "a"
                                        _ = fi55
                                        _ = v56
                                    }
                                }
                            }
                        case 1:
                            fmt.Println(v39, v43)
                        default:
                            v3 = true
                        }
                        _ = fi51
                        _ = v52
                    }
                    switch 127 {
                    case -494:
                        if (((0 <= 2) || (!(18446744073709551615 > 1))) || v3) {
                            var v74 float64 = (((0.1 / 1e10) - 0.1) + ((0.1 + 1e10) + v5))
                            for fi75 := 0; fi75 < 20; fi75++ {
                                var v76 float64 = v74
                                _ = fi75
                                _ = v76
                            }
                            _ = v74
                        }
                    case 9223372036854775807:
                        if (3.14159 == (v41 / float32(1.5))) {
                            v77 := v4[(v43 | (v43 - (v39 ^ (v39 >> ((390) & 63)))))]
                            v5 = (v5 * float64(0.0))
                            if len(v45) > 0 {
                                v45[0] = 2147483647
                            }
                            _ = v77
                        }
                    default:
                        v40 = v42
                    }
                    v78 := v4[42]
                    _ = v78
                }
                v1 = (v1 + v1)
                var v79 int8 = (v2 * -77)
                _ = v79
            } else {
                var v80 int16 = (v42 >> ((42) & 63))
                _ = v80
            }
            for fi81 := 0; fi81 < 7; fi81++ {
                v1 = ("\n\t" + "")
                _ = fi81
            }
        }
    } else {
        v40 = (v42 << ((716) & 63))
        for fi82 := 0; fi82 < 3; fi82++ {
            for fi83 := 0; fi83 < 20; fi83++ {
                var v84 float32 = -2.25
                var v85 float64 = v5
                _ = fi83
                _ = v84
                _ = v85
            }
            _ = fi82
        }
    }
    fmt.Println(v43, v40)
    _ = len(v4)
    v86 := GenS3{F0: 812, F1: 128}
    if (-1 == 644) {
        if (v3 || v3) {
            if (!false) {
                var v87 bool = v3
                v1 = (v1 + v1)
                for fi88 := 0; fi88 < 20; fi88++ {
                    fmt.Println(fi88, v2)
                    _ = fi88
                }
                _ = v87
            }
            var v89 float32 = float32(float64(3.14159))
            v40 -= v40
            _ = v89
        }
        v5 = (v5 / v5)
        v90 := GenS5{F0: 42, F1: 42, F2: (553 != 65535)}
        _ = v90
    }
    v45 = append(v45, 32767)
    var v91 uint64 = 86
    var v92 int32 = 1
    var v93 int64 = -652
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v86
    _ = v91
    _ = v92
    _ = v93
}
