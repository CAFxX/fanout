package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint16; F1 string; F2 string }
type GenS5 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint8; F1 uint32 }
type GenS8 struct { F0 uint32; F1 string; F2 int }
type GenS9 struct { F0 bool; F1 int32; F2 uint32 }
type GenS10 struct { F0 int64; F1 string; F2 uint32 }
type GenS11 struct { F0 int8; F1 int32; F2 int64 }
func genFn0(p0 int64) uint8 {
    v6 := GenS1{F0: 539}
    var v7 uint = 162
    v8 := GenS1{F0: v7}
    var v9 float32 = ((0.0 * 0.0) - (3.14159 - (((0.0 + 3.14159) + 1e10) * ((1e10 / 1.5) / 1.0))))
    if true {
        v10 := GenS2{F0: ((p0 ^ (p0 << ((0) & 63))) == p0), F1: 1}
        if false {
            fmt.Println(v9)
        } else {
            v11 := []int{2147483647, -2147483648}
            fmt.Println(p0)
            var v12 uint8 = 202
            _ = v11
            _ = v12
        }
        _ = v10
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 248
}
func genFn1(p0 int, p1 bool) uint16 {
    if false {
        for fi13 := 0; fi13 < 13; fi13++ {
            if ((78 == 0) || p1) {
                if false {
                    v14 := GenS3{F0: p1, F1: 0}
                    var v15 int16 = -81
                    _ = v14
                    _ = v15
                }
                if p1 {
                    v16 := GenS4{F0: 1, F1: ("\n\t" + "\n\t"), F2: "\n\t"}
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 18; fi17++ {
                        if false {
                            if (p1 || p1) {
                                v18 := GenS5{F0: 2, F1: 1, F2: -1}
                                var v19 uint = 762
                                var v20 int8 = 100
                                _ = v18
                                _ = v19
                                _ = v20
                            } else {
                                var v21 uint = 327
                                var v22 int16 = -32768
                                _ = v21
                                _ = v22
                            }
                        }
                        if ((((!("foo bar" > "a")) || (42 <= 42)) && (!(-2.25 > (0.1 * float64((0.0 + (0.1 * 3.14159))))))) || (773 != 42)) {
                            var v23 uint8 = 0
                            _ = v23
                        } else {
                            v24 := GenS6{F0: 0}
                            var v25 uint8 = 127
                            _ = v24
                            _ = v25
                        }
                        if ((-2.25 <= 0.0) && false) {
                            v26 := GenS6{F0: 4294967295}
                            _ = v26
                        }
                        p1 = p1
                        if p1 {
                            var v27 uint32 = 541
                            switch "\n\t" {
                            case "\n\t":
                                fmt.Println(p1, v27, fi17)
                            case "foo bar":
                                var v28 uint = 1
                                _ = v28
                            default:
                                if (-13 > -128) {
                                    fmt.Println(fi13, fi17)
                                }
                            }
                            _ = v27
                        } else {
                            v29 := []int{(p0 & (fi13 >> ((645) & 63)))}
                            p0 += p0
                            _ = v29
                        }
                        _ = fi17
                    }
                    if (-358 < 9223372036854775807) {
                        p1 = p1
                        fmt.Println(fi13, p1)
                        p1 = false
                    } else {
                        var v30 uint64 = 0
                        _ = v30
                    }
                }
                var v31 uint16 = 300
                _ = v31
            } else {
                v32 := GenS4{F0: 52, F1: "foo bar", F2: ((("foo bar" + "hello") + (((((((("foo bar" + "a") + "a") + (("\n\t" + "hello") + "foo bar")) + "") + ("\n\t" + "hello")) + "\n\t") + "\n\t") + ("hello" + "foo bar"))) + "\n\t")}
                _ = v32
            }
            _ = fi13
        }
        fmt.Println(p1, p1)
    } else {
        var v33 int16 = 904
        if p1 {
            var v34 uint = 0
            v35 := GenS2{F0: p1, F1: 42}
            fmt.Println(v34)
            _ = v34
            _ = v35
        } else {
            v36 := GenS1{F0: 1}
            _ = v36
        }
        _ = v33
    }
    for fi37 := 0; fi37 < 13; fi37++ {
        if p1 {
            fmt.Println(fi37, p1)
            for fi38 := 0; fi38 < 14; fi38++ {
                if p1 {
                    p0 = 0
                } else {
                    fmt.Println(p1)
                    v39 := []int{fi37, 0}
                    _ = v39
                }
                v40 := GenS1{F0: 765}
                p0 = (fi37 % (((p0 + p0)) | 1))
                _ = fi38
                _ = v40
            }
        } else {
            v41 := []int{(^(p0)), -20, p0}
            _ = v41
        }
        v42 := GenS1{F0: 862}
        _ = fi37
        _ = v42
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    v3 := make(map[uint32]string)
    var v4 string = "foo bar"
    v5 := []uint32{42}
    var v43 uint64 = 1
    if false {
        v44 := GenS5{F0: 0, F1: 1, F2: 0}
        _ = v44
    } else {
        fmt.Println(v4, v43)
        fmt.Println(v43, v43)
    }
    var v45 uint32 = 2
    var v46 float64 = 1.5
    v4 = v4
    var v47 int64 = 776
    if len(v5) > 0 {
        v5[0] = v45
    }
    switch 1 {
    case 42:
        for fi48 := 0; fi48 < 12; fi48++ {
            var v49 int16 = 42
            v43 = 2
            _ = fi48
            _ = v49
        }
    case 4294967295:
        if (1.5 <= (float32((-2.25 / 3.14159)) + 1e10)) {
            switch v4 {
            case "a":
                v47 = -1
            case "":
                if (!true) {
                    if false {
                        for fi50 := 0; fi50 < 6; fi50++ {
                            switch (v4 + ((v4 + "") + (v4 + (v4 + v4)))) {
                            case "hello":
                                v47 = (v47 + (-(v47)))
                            default:
                                switch 32768 {
                                case -994:
                                    if (127 != genFn0(v47)) {
                                        var v51 uint = uint(v45)
                                        _ = v51
                                    }
                                }
                            }
                            var v52 int16 = 1
                            for fi53 := 0; fi53 < 5; fi53++ {
                                var v54 int = (fi50 + fi50)
                                _ = fi53
                                _ = v54
                            }
                            _ = fi50
                            _ = v52
                        }
                        if (102 > 255) {
                            fmt.Println(v4, v4)
                            switch (("" + (v4 + "foo bar")) + "\n\t") {
                            case "a":
                                _ = len(v3)
                            default:
                                v55 := GenS2{F0: true, F1: 42}
                                _ = v55
                            }
                        }
                        fmt.Println(v43, v43, v46)
                    } else {
                        for fi56 := 0; fi56 < 10; fi56++ {
                            var v57 int8 = 100
                            _ = fi56
                            _ = v57
                        }
                        switch v4 {
                        case "\n\t":
                            v58 := v3[v45]
                            _ = v58
                        default:
                            v59 := GenS2{F0: (v43 != v43), F1: 1}
                            _ = v59
                        }
                    }
                    if ((-(v47)) >= (v47 + (v47 % ((v47) | 1)))) {
                        v4 = v4
                        if false {
                            v5 = append(v5, v45)
                            if true {
                                if (!false) {
                                    fmt.Println(v43, v43, v47)
                                    for fi60 := 0; fi60 < 14; fi60++ {
                                        v46 = v46
                                        _ = fi60
                                    }
                                    var v61 int8 = 127
                                    _ = v61
                                }
                            } else {
                                if (false || (false || true)) {
                                    for fi62 := 0; fi62 < 18; fi62++ {
                                        for fi63 := 0; fi63 < 8; fi63++ {
                                            if len(v5) > 0 {
                                                v5[0] = v45
                                            }
                                            _ = fi63
                                        }
                                        _ = fi62
                                    }
                                } else {
                                    var v64 uint16 = 602
                                    var v65 int = 42
                                    _ = v64
                                    _ = v65
                                }
                                var v66 bool = (-1 < -956)
                                var v67 string = v4
                                _ = v66
                                _ = v67
                            }
                        } else {
                            for fi68 := 0; fi68 < 2; fi68++ {
                                switch (fi68 - (^(fi68))) {
                                case -400:
                                    var v69 uint = 887
                                    _ = v69
                                case 687:
                                    v45 -= (v45 * v45)
                                case -2147483648:
                                    if true {
                                        if false {
                                            if false {
                                                delete(v1, ("\n\t" + ("hello" + v4)))
                                            } else {
                                                v70 := GenS7{F0: 42, F1: v45}
                                                var v71 uint8 = 1
                                                if false {
                                                    for fi72 := 0; fi72 < 10; fi72++ {
                                                        v73 := GenS7{F0: v71, F1: (v45 + v45)}
                                                        delete(v1, ((v4 + v4) + v4))
                                                        _ = fi72
                                                        _ = v73
                                                    }
                                                    v3[v45] = (v4 + "\n\t")
                                                    switch "foo bar" {
                                                    case "":
                                                        for fi74 := 0; fi74 < 17; fi74++ {
                                                            var v75 int16 = 0
                                                            switch ("\n\t" + "a") {
                                                            case "hello":
                                                                v47 += v47
                                                            case "":
                                                                if false {
                                                                    var v76 uint16 = 429
                                                                    _ = v76
                                                                }
                                                            default:
                                                                switch (fi74 + (fi68 >> ((42) & 63))) {
                                                                case 42:
                                                                    switch v4 {
                                                                    case "a":
                                                                        fmt.Println(v47, v47, v71)
                                                                    case "":
                                                                        v71 -= genFn0((v47 + -1))
                                                                    }
                                                                case -9223372036854775808:
                                                                    v77 := GenS7{F0: v71, F1: v45}
                                                                    if true {
                                                                        var v78 int = fi74
                                                                        _ = v78
                                                                    }
                                                                    _ = v77
                                                                case 32768:
                                                                    v79 := GenS2{F0: (42 <= 18446744073709551615), F1: 127}
                                                                    _ = v79
                                                                default:
                                                                    fmt.Println(fi74, v43)
                                                                }
                                                            }
                                                            switch "hello" {
                                                            case "\n\t":
                                                                if ((!(v71 >= v71)) && (-817 > (^(v47)))) {
                                                                    if true {
                                                                        v47 += (v47 % (((v47 + 42)) | 1))
                                                                        v47 = (v47 & (v47 & v47))
                                                                        var v80 uint16 = 306
                                                                        _ = v80
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                        v1["hello"] = v4
                                                                    } else {
                                                                        switch fi68 {
                                                                        case 1:
                                                                            v43 = (v43 >> ((4294967295) & 63))
                                                                        case 127:
                                                                            if false {
                                                                                v81 := GenS3{F0: (((v71 | (^(v71))) != v71) || ((1 <= (-(v47))) && true)), F1: 4294967295}
                                                                                v47 += v47
                                                                                _ = v81
                                                                            } else {
                                                                                _ = len(v5)
                                                                                v82 := GenS1{F0: 1}
                                                                                v3[2147483647] = v4
                                                                                _ = v82
                                                                            }
                                                                            fmt.Println(v43)
                                                                            _ = len(v5)
                                                                        default:
                                                                            v45 = v45
                                                                        }
                                                                        _ = len(v3)
                                                                        v4 = ("\n\t" + ("\n\t" + "foo bar"))
                                                                    }
                                                                    v43 = (^(v43))
                                                                }
                                                            default:
                                                                switch "" {
                                                                case "hello":
                                                                    v47 = v47
                                                                case "foo bar":
                                                                    switch "" {
                                                                    case "":
                                                                        if false {
                                                                            fmt.Println(v71, v46)
                                                                            v83 := GenS8{F0: (v45 | (v45 << ((2) & 63))), F1: "foo bar", F2: (fi68 & fi74)}
                                                                            _ = v83
                                                                        } else {
                                                                            fmt.Println(v4, v4, v46)
                                                                            switch v4 {
                                                                            case "":
                                                                                var v84 float64 = v46
                                                                                _ = v84
                                                                            case "a":
                                                                                if false {
                                                                                    v3[669] = v4
                                                                                    var v85 int16 = (^(v75))
                                                                                    _ = v85
                                                                                } else {
                                                                                    var v86 int32 = 2147483647
                                                                                    _ = v86
                                                                                }
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = v45
                                                                                }
                                                                                fmt.Println(v46)
                                                                            case "\n\t":
                                                                                v46 = (-2.25 - 3.14159)
                                                                            }
                                                                            v87 := GenS3{F0: (!false), F1: v43}
                                                                            _ = v87
                                                                        }
                                                                    case "foo bar":
                                                                        if ((v43 / (((^(v43))) | 1)) < (^(v43))) {
                                                                            switch (^(fi68)) {
                                                                            case 0:
                                                                                fmt.Println(v71)
                                                                                var v88 int16 = -263
                                                                                switch "\n\t" {
                                                                                case "\n\t":
                                                                                    v46 = 1e10
                                                                                case "foo bar":
                                                                                    if false {
                                                                                        v88 += 42
                                                                                        if (!false) {
                                                                                            if true {
                                                                                                fmt.Println(v75, v75, v71)
                                                                                                fmt.Println(fi74, v46)
                                                                                            }
                                                                                            var v89 float64 = v46
                                                                                            _ = v89
                                                                                        } else {
                                                                                            switch "\n\t" {
                                                                                            case "\n\t":
                                                                                                _ = len(v2)
                                                                                            case "":
                                                                                                v90 := GenS9{F0: true, F1: -340, F2: v45}
                                                                                                _ = v90
                                                                                            case "foo bar":
                                                                                                var v91 float32 = float32(v46)
                                                                                                var v92 uint8 = (v71 % (((v71 >> ((255) & 63))) | 1))
                                                                                                _ = v91
                                                                                                _ = v92
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    delete(v3, 4294967295)
                                                                                }
                                                                                _ = v88
                                                                            }
                                                                            v46 += float64(1.5)
                                                                        } else {
                                                                            v47 = v47
                                                                            v46 += v46
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            var v93 int64 = (v47 << ((1) & 63))
                                                                            var v94 string = (v4 + v4)
                                                                            _ = v93
                                                                            _ = v94
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            _ = fi74
                                                            _ = v75
                                                        }
                                                    case "a":
                                                        for fi95 := 0; fi95 < 8; fi95++ {
                                                            v96 := GenS10{F0: 9223372036854775807, F1: "foo bar", F2: 2147483648}
                                                            _ = fi95
                                                            _ = v96
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v43)
                                                }
                                                _ = v70
                                                _ = v71
                                            }
                                            v43 = 42
                                        }
                                        fmt.Println(v4, v4)
                                        if (genFn0(-1) != 151) {
                                            delete(v3, (v45 * (v45 - v45)))
                                            var v97 uint64 = (v43 << ((2) & 63))
                                            v98 := GenS11{F0: 100, F1: 99, F2: (v47 % (((^(v47))) | 1))}
                                            _ = v97
                                            _ = v98
                                        }
                                    }
                                default:
                                    v47 += 1
                                }
                                var v99 uint8 = 1
                                _ = fi68
                                _ = v99
                            }
                        }
                        v46 += -2.25
                    }
                    switch 32767 {
                    case 128:
                        var v100 uint16 = 652
                        _ = v100
                    case -229:
                        if (!false) {
                            _ = len(v3)
                            for fi101 := 0; fi101 < 19; fi101++ {
                                v46 += ((v46 + (v46 / 0.1)) * v46)
                                v47 = v47
                                _ = fi101
                            }
                        } else {
                            var v102 uint64 = 632
                            _ = v102
                        }
                    case 32767:
                        v47 = v47
                    }
                } else {
                    v5 = append(v5, v45)
                }
            case "\n\t":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            if ((672 != (v45 + (v45 | (v45 >> ((452) & 63))))) || (uint16(v43) <= 2)) {
                var v103 bool = (v45 > v45)
                var v104 string = ((v4 + "\n\t") + (v4 + v4))
                _ = v103
                _ = v104
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
