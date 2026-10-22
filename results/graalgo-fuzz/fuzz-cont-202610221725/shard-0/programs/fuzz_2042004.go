package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 int8; F1 string }
type GenS8 struct { F0 uint64; F1 int8 }
type GenS9 struct { F0 uint64; F1 uint32; F2 int }
type GenS10 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint16; F1 string; F2 bool }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 int64; F1 uint16 }
func genFn0(p0 int, p1 int32) string {
    fmt.Println(p0)
    p0 += p0
    v6 := GenS1{F0: 1, F1: 160}
    _ = p0
    _ = p1
    _ = v6
    return "\n\t"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 0.1
    v3 := []int{798}
    v4 := make(map[uint32]bool)
    v5 := make(map[string]bool)
    var v7 int = 856
    var v8 bool = false
    v3 = append(v3, -72)
    switch ("" + "hello") {
    case "a":
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 string = ""
            _ = fi9
            _ = v10
        }
    }
    fmt.Println(v2, v2, v2)
    v11 := v5[""]
    fmt.Println(v11, v8)
    v8 = v11
    if len(v3) > 0 {
        v3[0] = v7
    }
    v12 := GenS2{F0: 128}
    switch "" {
    case "a":
        var v13 uint64 = 673
        _ = v13
    }
    var v14 string = "a"
    var v15 uint = 42
    switch v7 {
    case -9223372036854775808:
        var v16 uint16 = 42
        _ = v16
    case 127:
        v17 := GenS3{F0: 1}
        _ = v17
    case -1:
        v18 := GenS4{F0: 239}
        _ = v18
    default:
        fmt.Println(v15, v14, v15)
    }
    var v19 float32 = v2
    if (!(128 < 42)) {
        if ((false || v8) || false) {
            for fi20 := 0; fi20 < 6; fi20++ {
                v8 = v8
                _ = fi20
            }
        } else {
            var v21 int32 = 42
            for fi22 := 0; fi22 < 13; fi22++ {
                for fi23 := 0; fi23 < 5; fi23++ {
                    switch "foo bar" {
                    case "a":
                        v19 = v19
                    case "":
                        v24 := GenS2{F0: 42}
                        _ = v24
                    case "foo bar":
                        if (v8 && (3.14159 < ((-2.25 / (v2 + float32((1.5 * 1.5)))) / 1.5))) {
                            if (1e10 < ((float64(v19) / 3.14159) + -2.25)) {
                                if (v8 || true) {
                                    var v25 int = -1
                                    switch -855 {
                                    case -938:
                                        v26 := GenS3{F0: 1}
                                        _ = v26
                                    case -686:
                                        delete(v4, 2)
                                    case 2147483647:
                                        var v27 uint16 = 525
                                        _ = v27
                                    default:
                                        fmt.Println(fi22, fi23)
                                    }
                                    _ = len(v3)
                                    _ = v25
                                } else {
                                    if (int8(fi23) != -11) {
                                        fmt.Println(v21, v14)
                                    }
                                    var v28 int16 = -1
                                    _ = v28
                                }
                                v29 := GenS4{F0: 270}
                                _ = v29
                            }
                            var v30 string = ((v14 + v14) + (v14 + "\n\t"))
                            _ = v30
                        } else {
                            v31 := GenS4{F0: 4294967295}
                            v8 = (v14 > "hello")
                            _ = v31
                        }
                        if false {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        v32 := GenS5{F0: 676}
                        _ = v32
                    default:
                        if false {
                            fmt.Println(v7)
                            delete(v1, v14)
                            if (!(v11 || true)) {
                                v21 = v21
                                if v11 {
                                    v15 = (v15 | (^(v15)))
                                } else {
                                    switch v7 {
                                    case 2147483647:
                                        var v33 uint16 = 122
                                        _ = v33
                                    case 992:
                                        v34 := GenS3{F0: 0}
                                        _ = v34
                                    case -1:
                                        switch v14 {
                                        case "\n\t":
                                            switch (v7 - (v7 * 128)) {
                                            case -930:
                                                fmt.Println(v2)
                                            default:
                                                v35 := v4[1]
                                                _ = v35
                                            }
                                        default:
                                            v36 := v1[v14]
                                            _ = v36
                                        }
                                    default:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    var v37 int = fi22
                                    _ = v37
                                }
                            } else {
                                if v8 {
                                    var v38 uint32 = 4294967295
                                    v21 = v21
                                    _ = v38
                                }
                                fmt.Println(v14, v2)
                                fmt.Println(fi23, fi22)
                            }
                        }
                    }
                    if v11 {
                        if (-9223372036854775808 != -65) {
                            v19 += float32((float64(v19) + (((3.14159 / float64((1e10 * v19))) / 3.14159) - ((1.5 * 0.0) + 0.0))))
                            v39 := GenS1{F0: -32768, F1: v15}
                            _ = v39
                        } else {
                            switch (fi22 << (((^(v15))) & 63)) {
                            case 127:
                                var v40 float64 = ((0.1 / 1e10) / 1.0)
                                _ = v40
                            case -2147483648:
                                if true {
                                    if (2 != 2) {
                                        var v41 int32 = v21
                                        _ = v41
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 42
                                    }
                                    var v42 uint8 = 0
                                    _ = v42
                                } else {
                                    v43 := GenS4{F0: 250}
                                    switch v14 {
                                    case "\n\t":
                                        v44 := GenS1{F0: 0, F1: (v15 << ((v15) & 63))}
                                        _ = v44
                                    case "foo bar":
                                        if (127 != -1) {
                                            v45 := GenS4{F0: 1}
                                            if true {
                                                v21 += v21
                                            }
                                            _ = v45
                                        }
                                    case "a":
                                        v15 = (v15 + v15)
                                    default:
                                        var v46 string = v14
                                        _ = v46
                                    }
                                    _ = v43
                                }
                            case -365:
                                if v11 {
                                    if (111 <= -118) {
                                        var v47 uint64 = 229
                                        if (v11 || true) {
                                            v48 := GenS6{F0: false, F1: 42}
                                            v47 = (^(v47))
                                            _ = v48
                                        }
                                        _ = v47
                                    } else {
                                        if (!false) {
                                            v2 = -2.25
                                            v2 = (float32((-2.25 - ((3.14159 - float64((v2 - 0.1))) - (((1.5 / (1.5 / (1.5 / (3.14159 + float64((float32(0.0) - (0.0 * (v2 + (0.0 - v2))))))))) + -2.25) + ((1e10 + -2.25) * 0.0))))) * v2)
                                            var v49 int8 = -1
                                            _ = v49
                                        } else {
                                            var v50 int32 = (^(v21))
                                            if true {
                                                fmt.Println(v14, fi22, fi23)
                                                v51 := GenS5{F0: 65535}
                                                _ = v51
                                            } else {
                                                fmt.Println(fi22)
                                                v52 := GenS7{F0: 73, F1: (v14 + ("hello" + "hello"))}
                                                switch (v14 + v14) {
                                                case "hello":
                                                    if v11 {
                                                        fmt.Println(fi22)
                                                        fmt.Println(v19, v2)
                                                        fmt.Println(v8)
                                                    } else {
                                                        if (!(-379 < 1)) {
                                                            v14 = "\n\t"
                                                            _ = len(v4)
                                                        }
                                                        var v53 int64 = 9223372036854775807
                                                        _ = v53
                                                    }
                                                case "a":
                                                    v54 := GenS2{F0: 33}
                                                    _ = v54
                                                }
                                                if (v8 && false) {
                                                    v15 = (v15 | v15)
                                                } else {
                                                    var v55 uint32 = 65535
                                                    fmt.Println(v21)
                                                    v11 = v11
                                                    _ = v55
                                                }
                                                var v56 uint32 = 926
                                                v7 = fi23
                                                switch v14 {
                                                case "":
                                                    v50 += (v50 << ((v15) & 63))
                                                case "hello":
                                                    var v57 int = (^(v7))
                                                    _ = v57
                                                case "\n\t":
                                                    switch ((v14 + (v14 + (v14 + ""))) + v14) {
                                                    case "":
                                                        if ((v11 || (!(!(2147483647 < (v21 / ((v21) | 1)))))) && ((^(v56)) < v56)) {
                                                            var v58 int = (fi23 << ((v15) & 63))
                                                            var v59 int32 = 0
                                                            v14 = v14
                                                            _ = v58
                                                            _ = v59
                                                        } else {
                                                            fmt.Println(v8)
                                                            if ("foo bar" > "hello") {
                                                                fmt.Println(v50, v11)
                                                                if false {
                                                                    if (0.1 == (((1e10 - 0.0) * 0.0) + 1.5)) {
                                                                        switch v14 {
                                                                        case "\n\t":
                                                                            v60 := GenS1{F0: int16(v56), F1: v15}
                                                                            fmt.Println(v11, v11, v14)
                                                                            _ = v60
                                                                        case "a":
                                                                            var v61 int = (fi23 - v7)
                                                                            v2 = ((0.1 / -2.25) / (1e10 + v2))
                                                                            v21 = 453
                                                                            _ = v61
                                                                        case "foo bar":
                                                                            if v8 {
                                                                                v62 := GenS6{F0: (!v11), F1: int16(v15)}
                                                                                v3 = append(v3, (fi23 + int(v15)))
                                                                                v11 = true
                                                                                _ = v62
                                                                            } else {
                                                                                var v63 uint16 = 689
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = 127
                                                                                }
                                                                                _ = v63
                                                                            }
                                                                        }
                                                                        _ = len(v3)
                                                                        v56 += (v56 + (v56 + (v56 << (((v15 << (((v15 * 236)) & 63))) & 63))))
                                                                    }
                                                                    if (v2 == v2) {
                                                                        _ = len(v3)
                                                                        var v64 uint16 = 2
                                                                        v14 = ("" + v14)
                                                                        _ = v64
                                                                    }
                                                                }
                                                                if (32767 <= 187) {
                                                                    if v8 {
                                                                        switch (-(fi23)) {
                                                                        case -254:
                                                                            var v65 uint8 = 10
                                                                            _ = v65
                                                                        case 42:
                                                                            v7 -= fi23
                                                                            var v66 float32 = v2
                                                                            _ = v66
                                                                        case -1:
                                                                            var v67 int = fi23
                                                                            _ = v67
                                                                        }
                                                                    } else {
                                                                        switch (v14 + "\n\t") {
                                                                        case "a":
                                                                            if false {
                                                                                v68 := v4[(v56 & v56)]
                                                                                var v69 int16 = 1000
                                                                                v70 := GenS8{F0: 0, F1: 1}
                                                                                _ = v68
                                                                                _ = v69
                                                                                _ = v70
                                                                            } else {
                                                                                v11 = true
                                                                            }
                                                                        default:
                                                                            var v71 uint8 = 0
                                                                            fmt.Println(fi23, v19)
                                                                            var v72 int64 = 0
                                                                            _ = v71
                                                                            _ = v72
                                                                        }
                                                                        fmt.Println(v19)
                                                                    }
                                                                    if v11 {
                                                                        v19 = v2
                                                                        v56 += 1
                                                                    } else {
                                                                        v3 = append(v3, (fi23 | fi23))
                                                                        switch (v14 + (v14 + "hello")) {
                                                                        case "\n\t":
                                                                            var v73 uint = v15
                                                                            _ = v73
                                                                        case "foo bar":
                                                                            v74 := GenS8{F0: 1, F1: 42}
                                                                            _ = v74
                                                                        default:
                                                                            v75 := v1[v14]
                                                                            _ = v75
                                                                        }
                                                                        fmt.Println(v8)
                                                                        switch fi23 {
                                                                        case -585:
                                                                            v76 := GenS1{F0: 32767, F1: (^(v15))}
                                                                            _ = v76
                                                                        case -1:
                                                                            _ = len(v3)
                                                                        case 128:
                                                                            var v77 int = (v7 % ((-547) | 1))
                                                                            _ = v77
                                                                        }
                                                                        var v78 int = fi23
                                                                        _ = v78
                                                                    }
                                                                    if len(v3) > 0 {
                                                                        v3[0] = v7
                                                                    }
                                                                    var v79 int32 = v21
                                                                    var v80 int8 = -1
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                            }
                                                            var v81 int64 = -9223372036854775808
                                                            _ = v81
                                                        }
                                                    case "\n\t":
                                                        v11 = v8
                                                    case "hello":
                                                        var v82 uint8 = 100
                                                        _ = v82
                                                    default:
                                                        fmt.Println(v15)
                                                    }
                                                default:
                                                    var v83 int16 = 607
                                                    _ = v83
                                                }
                                                var v84 uint32 = v56
                                                if (128 < 42) {
                                                    switch (-(fi22)) {
                                                    case 128:
                                                        fmt.Println(fi23)
                                                    case -206:
                                                        v14 = "\n\t"
                                                    case 4294967295:
                                                        v85 := GenS6{F0: false, F1: 983}
                                                        _ = v85
                                                    }
                                                    if (158 > 201) {
                                                        v14 = v14
                                                        v86 := GenS9{F0: 1, F1: 1, F2: fi23}
                                                        if (v8 && true) {
                                                            var v87 uint8 = 255
                                                            var v88 float64 = -2.25
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                        _ = v86
                                                    } else {
                                                        fmt.Println(fi22, fi22)
                                                        var v89 int16 = 42
                                                        if (0 > 255) {
                                                            v90 := GenS10{F0: 958, F1: (^(v84)), F2: 42}
                                                            v21 += v50
                                                            v91 := GenS11{F0: (v15 << ((255) & 63))}
                                                            _ = v90
                                                            _ = v91
                                                        }
                                                        _ = v89
                                                    }
                                                    var v92 int = (fi23 + fi22)
                                                    _ = v92
                                                } else {
                                                    v93 := GenS12{F0: 914, F1: ((v14 + v14) + ((v14 + v14) + v14)), F2: ((v15 / ((255) | 1)) < 995)}
                                                    fmt.Println(v50, v21)
                                                    if (true || false) {
                                                        fmt.Println(v15, v21)
                                                    }
                                                    _ = v93
                                                }
                                                _ = v52
                                                _ = v56
                                                _ = v84
                                            }
                                            _ = v50
                                        }
                                        var v94 uint = v15
                                        var v95 uint = (v15 / (((^(v15))) | 1))
                                        _ = v94
                                        _ = v95
                                    }
                                    var v96 uint16 = 255
                                    _ = v96
                                } else {
                                    v8 = v8
                                    var v97 float32 = v2
                                    _ = len(v1)
                                    v14 = ((v14 + (v14 + v14)) + (v14 + v14))
                                    _ = v97
                                }
                            default:
                                _ = len(v3)
                            }
                            delete(v4, 377)
                            fmt.Println(v7)
                        }
                    }
                    var v98 int = -873
                    _ = fi23
                    _ = v98
                }
                _ = fi22
            }
            _ = v21
        }
        for fi99 := 0; fi99 < 19; fi99++ {
            switch fi99 {
            case 32768:
                var v100 bool = (!(-34 <= 1))
                _ = v100
            case 42:
                v101 := GenS13{F0: 0}
                _ = v101
            }
            _ = fi99
        }
    } else {
        fmt.Println(v8)
        v3 = append(v3, (-(v7)))
        v102 := GenS14{F0: -746, F1: 42}
        _ = v102
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
