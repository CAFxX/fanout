package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int8 }
type GenS2 struct { F0 int16; F1 uint32; F2 int16 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 int64; F2 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool; F1 bool; F2 int }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 uint32; F1 int8 }
type GenS10 struct { F0 uint16; F1 int64; F2 string }
type GenS11 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS12 struct { F0 int64; F1 int32 }
func genFn0(p0 bool, p1 string) uint32 {
    switch ("\n\t" + (p1 + (p1 + "foo bar"))) {
    case "hello":
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 int16 = 0
            var v8 bool = true
            _ = fi6
            _ = v7
            _ = v8
        }
    case "":
        for fi9 := 0; fi9 < 16; fi9++ {
            if p0 {
                p1 = ""
            }
            var v10 uint = 1
            for fi11 := 0; fi11 < 10; fi11++ {
                if false {
                    v12 := GenS1{F0: 2147483647, F1: -32768, F2: 42}
                    _ = v12
                }
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
    case "a":
        p1 = (p1 + "")
    default:
        v13 := []int{127}
        _ = v13
    }
    var v14 uint64 = 2
    v15 := GenS1{F0: 42, F1: -920, F2: -58}
    p0 = (!(114 >= 3))
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 0
}
func genFn1(p0 int32) bool {
    if true {
        v16 := []int{-1, 42, -719}
        fmt.Println(p0)
        p0 = p0
        _ = v16
    }
    v17 := GenS1{F0: (p0 + (p0 << ((255) & 63))), F1: -998, F2: -119}
    fmt.Println(p0)
    p0 = (p0 >> ((255) & 63))
    _ = p0
    _ = v17
    return true
}

func main() {
    v1 := []uint32{470, 65535, 2147483648, 4294967295}
    var v2 string = "a"
    var v3 uint16 = 65535
    var v4 int = 32768
    v5 := make(map[int]bool)
    if (!true) {
        if true {
            if false {
                var v18 int8 = 100
                _ = v18
            }
        }
        delete(v5, (v4 | (v4 ^ (v4 ^ (^(v4))))))
        v3 = (v3 * (v3 & 642))
    } else {
        v19 := GenS1{F0: -524, F1: 159, F2: 81}
        _ = v19
    }
    var v20 bool = (-9223372036854775808 <= -190)
    v21 := GenS2{F0: 1000, F1: 65535, F2: 935}
    switch "hello" {
    case "a":
        if (v20 && (v20 || v20)) {
            fmt.Println(v2, v2, v2)
        }
    case "hello":
        v1 = append(v1, 4294967295)
    default:
        fmt.Println(v20, v2)
    }
    var v22 uint = 4294967295
    var v23 int32 = 0
    switch "hello" {
    case "":
        var v24 int64 = -1
        _ = v24
    }
    var v25 int = v4
    fmt.Println(v20)
    fmt.Println(v2)
    v23 = v23
    for fi26 := 0; fi26 < 12; fi26++ {
        if (1e10 < 0.0) {
            v3 -= (v3 - (^(v3)))
        }
        _ = fi26
    }
    switch (v2 + ("\n\t" + v2)) {
    case "foo bar":
        var v27 int = 1
        _ = v27
    case "hello":
        v5[4294967295] = false
    }
    switch "foo bar" {
    case "hello":
        v23 -= -532
    case "foo bar":
        for fi28 := 0; fi28 < 17; fi28++ {
            var v29 float32 = 1e10
            if (v22 != (v22 ^ v22)) {
                if v20 {
                    if (v3 < v3) {
                        v30 := GenS3{F0: -660, F1: true}
                        switch (v25 << (((^(v22))) & 63)) {
                        case 42:
                            var v31 uint8 = 128
                            _ = v31
                        default:
                            v2 = v2
                        }
                        _ = v30
                    } else {
                        if v20 {
                            v32 := GenS1{F0: 506, F1: 1, F2: 42}
                            _ = v32
                        } else {
                            v20 = v20
                            for fi33 := 0; fi33 < 4; fi33++ {
                                if (!v20) {
                                    if v20 {
                                        var v34 float32 = (float32(1.5) * 3.14159)
                                        v22 -= (v22 % (((v22 - v22)) | 1))
                                        v35 := GenS4{F0: v23}
                                        _ = v34
                                        _ = v35
                                    }
                                    v23 -= 0
                                }
                                if true {
                                    v23 = v23
                                    switch v4 {
                                    case 42:
                                        v23 = v23
                                        var v36 int8 = -128
                                        _ = v36
                                    default:
                                        if (v20 || v20) {
                                            switch v2 {
                                            case "hello":
                                                switch 924 {
                                                case 0:
                                                    var v37 int32 = (-(v23))
                                                    _ = v37
                                                default:
                                                    switch 790 {
                                                    case -9223372036854775808:
                                                        var v38 int64 = -9223372036854775808
                                                        _ = v38
                                                    case 2147483647:
                                                        v20 = (float64(float32((0.1 * (-2.25 * (-2.25 / 3.14159))))) != 1e10)
                                                    }
                                                    var v39 int16 = 754
                                                    v25 += 2147483647
                                                    delete(v5, fi33)
                                                    v4 += fi33
                                                    if false {
                                                        v40 := GenS4{F0: -58}
                                                        fmt.Println(fi28, v39)
                                                        var v41 int8 = 100
                                                        var v42 uint32 = 2147483647
                                                        _ = len(v1)
                                                        _ = v40
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    if ((!((0 == 0) || (false || (true || v20)))) && v20) {
                                                        fmt.Println(v22, v2, v4)
                                                        switch v2 {
                                                        case "":
                                                            fmt.Println(v22)
                                                        case "hello":
                                                            v43 := GenS1{F0: (v23 << (((v22 >> (((v22 + v22)) & 63))) & 63)), F1: (v39 >> ((655) & 63)), F2: -128}
                                                            _ = v43
                                                        case "a":
                                                            v44 := GenS3{F0: (^(v39)), F1: v20}
                                                            switch 127 {
                                                            case -1:
                                                                v2 = ((v2 + v2) + v2)
                                                            case 2147483647:
                                                                fmt.Println(fi33)
                                                            default:
                                                                v3 = (v3 | v3)
                                                            }
                                                            v45 := v5[v4]
                                                            switch v2 {
                                                            case "hello":
                                                                v46 := GenS5{F0: (v22 + (v22 ^ uint(v3))), F1: -1, F2: 42}
                                                                _ = v46
                                                            case "":
                                                                v20 = v20
                                                            }
                                                            _ = v44
                                                            _ = v45
                                                        default:
                                                            switch (-(v25)) {
                                                            case 32767:
                                                                if true {
                                                                    delete(v5, v25)
                                                                    if (1 < v25) {
                                                                        v47 := GenS5{F0: 1, F1: -9223372036854775808, F2: v22}
                                                                        var v48 uint8 = 132
                                                                        _ = v47
                                                                        _ = v48
                                                                    } else {
                                                                        v23 -= v23
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        var v49 uint = v22
                                                                        _ = v49
                                                                    }
                                                                }
                                                            case 2147483647:
                                                                fmt.Println(v4, v25, fi28)
                                                            case 602:
                                                                v39 = v39
                                                            }
                                                        }
                                                        v50 := GenS2{F0: (-(v39)), F1: 169, F2: (^(v39))}
                                                        switch "a" {
                                                        case "":
                                                            if v20 {
                                                                switch (v4 + (v4 | v4)) {
                                                                case 0:
                                                                    if (v20 || v20) {
                                                                        if ((fi33 / ((fi33) | 1)) <= int(v3)) {
                                                                            if (255 > v3) {
                                                                                switch fi28 {
                                                                                case -9223372036854775808:
                                                                                    var v51 bool = v20
                                                                                    var v52 uint = (^(v22))
                                                                                    _ = v51
                                                                                    _ = v52
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    if (!v20) {
                                                                        var v53 int8 = -60
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        _ = v53
                                                                    }
                                                                }
                                                                var v54 float32 = v29
                                                                _ = v54
                                                            } else {
                                                                var v55 uint = v22
                                                                if v20 {
                                                                    if (193 <= v22) {
                                                                        v23 += v23
                                                                        switch (^(v25)) {
                                                                        case -2147483648:
                                                                            fmt.Println(v55, v29, v20)
                                                                        case -534:
                                                                            if false {
                                                                                fmt.Println(v25, v4)
                                                                                var v56 string = v2
                                                                                var v57 float32 = (float32((float64((float32(0.1) * (0.1 + (v29 * 1e10)))) - 0.0)) / 1.5)
                                                                                _ = v56
                                                                                _ = v57
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v58 := GenS3{F0: v39, F1: (!(996 >= 1))}
                                                                        switch 32768 {
                                                                        case -803:
                                                                            v2 = "foo bar"
                                                                        case -745:
                                                                            if (!false) {
                                                                                switch v25 {
                                                                                case 0:
                                                                                    fmt.Println(v55, v20, v39)
                                                                                case 701:
                                                                                    v59 := GenS6{F0: uint8(v3)}
                                                                                    _ = v59
                                                                                case 98:
                                                                                    v22 = v55
                                                                                default:
                                                                                    fmt.Println(v29)
                                                                                }
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = genFn0(true, v2)
                                                                                }
                                                                                fmt.Println(v29, fi33)
                                                                            } else {
                                                                                var v60 int64 = -9223372036854775808
                                                                                _ = v60
                                                                            }
                                                                        }
                                                                        switch v2 {
                                                                        case "foo bar":
                                                                            v3 = v3
                                                                        default:
                                                                            if true {
                                                                                v61 := GenS6{F0: 1}
                                                                                _ = v61
                                                                            }
                                                                        }
                                                                        _ = v58
                                                                    }
                                                                }
                                                                _ = v55
                                                            }
                                                        }
                                                        if (91 >= -23) {
                                                            switch v2 {
                                                            case "foo bar":
                                                                v22 = (v22 ^ (^(v22)))
                                                            case "\n\t":
                                                                var v62 uint64 = 0
                                                                _ = v62
                                                            }
                                                            if true {
                                                                if (-128 == -72) {
                                                                    _ = len(v1)
                                                                    fmt.Println(v29, fi28)
                                                                } else {
                                                                    v5[(fi28 | fi28)] = v20
                                                                    v63 := GenS7{F0: v20, F1: v20, F2: 42}
                                                                    if v20 {
                                                                        v4 = int(v3)
                                                                    } else {
                                                                        v25 += v25
                                                                        if v20 {
                                                                            if len(v1) > 0 {
                                                                                v1[0] = genFn0(v20, "\n\t")
                                                                            }
                                                                            if v20 {
                                                                                var v64 uint8 = 2
                                                                                switch (v25 % ((fi33) | 1)) {
                                                                                case -2147483648:
                                                                                    fmt.Println(v23, v64, v29)
                                                                                case -563:
                                                                                    if (genFn0((2147483647 >= 2147483647), (v2 + v2)) > 42) {
                                                                                        v65 := GenS8{F0: 0, F1: 2}
                                                                                        v66 := v5[fi33]
                                                                                        fmt.Println(fi28, fi33)
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                    }
                                                                                case 1:
                                                                                    v67 := GenS9{F0: 576, F1: 49}
                                                                                    _ = v67
                                                                                }
                                                                                _ = v64
                                                                            } else {
                                                                                v20 = false
                                                                                fmt.Println(fi33, v2, v3)
                                                                                if v20 {
                                                                                    var v68 uint64 = 1
                                                                                    _ = v68
                                                                                } else {
                                                                                    v69 := GenS9{F0: 2147483648, F1: -1}
                                                                                    v70 := v5[v25]
                                                                                    _ = v69
                                                                                    _ = v70
                                                                                }
                                                                                if v20 {
                                                                                    var v71 string = v2
                                                                                    v23 += 1
                                                                                    if (!v20) {
                                                                                        v25 -= 1
                                                                                        v72 := GenS9{F0: genFn0(v20, v71), F1: 68}
                                                                                        v73 := GenS10{F0: 2, F1: -1, F2: (((v2 + "hello") + v2) + v71)}
                                                                                        _ = v72
                                                                                        _ = v73
                                                                                    }
                                                                                    _ = v71
                                                                                }
                                                                            }
                                                                            if v20 {
                                                                                v1 = append(v1, 65535)
                                                                                v22 -= 720
                                                                            }
                                                                        } else {
                                                                            var v74 bool = (!false)
                                                                            v75 := GenS11{F0: 1, F1: 32767, F2: 208}
                                                                            _ = v74
                                                                            _ = v75
                                                                        }
                                                                        v5[v4] = v20
                                                                        var v76 bool = ((^(fi28)) < (fi28 + fi28))
                                                                        _ = v76
                                                                    }
                                                                    _ = v63
                                                                }
                                                                v5[v25] = (v20 || ((v3 << ((v22) & 63)) == (v3 >> ((v22) & 63))))
                                                            } else {
                                                                var v77 int = fi28
                                                                var v78 int64 = int64(v39)
                                                                v3 += (v3 >> ((1) & 63))
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                            if v20 {
                                                                var v79 uint8 = 127
                                                                v29 = v29
                                                                v80 := GenS10{F0: (v3 | 1), F1: int64(v79), F2: v2}
                                                                _ = v79
                                                                _ = v80
                                                            } else {
                                                                var v81 uint32 = 65535
                                                                if false {
                                                                    v82 := GenS2{F0: v39, F1: v81, F2: (v39 | v39)}
                                                                    var v83 bool = (v25 <= (^(fi28)))
                                                                    var v84 int8 = 0
                                                                    _ = v82
                                                                    _ = v83
                                                                    _ = v84
                                                                }
                                                                _ = v81
                                                            }
                                                        }
                                                        _ = v50
                                                    } else {
                                                        fmt.Println(v4, v23)
                                                    }
                                                    var v85 uint16 = 524
                                                    _ = v39
                                                    _ = v85
                                                }
                                            case "\n\t":
                                                var v86 float64 = 1e10
                                                _ = v86
                                            }
                                        } else {
                                            v3 -= (v3 + 1)
                                            v87 := GenS12{F0: 9223372036854775807, F1: v23}
                                            if (-65 == -96) {
                                                fmt.Println(v29, v25)
                                                v22 = v22
                                            }
                                            _ = v87
                                        }
                                    }
                                } else {
                                    var v88 uint64 = 867
                                    var v89 float64 = 1.5
                                    fmt.Println(v2)
                                    _ = v88
                                    _ = v89
                                }
                                var v90 uint16 = v3
                                _ = fi33
                                _ = v90
                            }
                        }
                    }
                } else {
                    v4 -= (^(v25))
                }
            }
            _ = fi28
            _ = v29
        }
    case "\n\t":
        delete(v5, (v4 * (v25 - 128)))
    default:
        if v20 {
            var v91 uint16 = v3
            var v92 int8 = 24
            var v93 uint16 = v91
            _ = v91
            _ = v92
            _ = v93
        } else {
            fmt.Println(v3, v25, v2)
            var v94 float64 = -2.25
            _ = v94
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
