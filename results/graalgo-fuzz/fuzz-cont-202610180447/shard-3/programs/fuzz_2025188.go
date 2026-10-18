package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int }
type GenS2 struct { F0 bool; F1 bool; F2 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 uint; F2 int16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32; F1 int8; F2 int64 }
type GenS8 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS9 struct { F0 int64; F1 int8 }
type GenS10 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint8, p1 int64, p2 bool) bool {
    var v4 int64 = p1
    if false {
        for fi5 := 0; fi5 < 18; fi5++ {
            switch (fi5 % ((1) | 1)) {
            case 42:
                p0 = (p0 << ((230) & 63))
            case 128:
                v6 := GenS1{F0: 73, F1: 2147483647, F2: fi5}
                _ = v6
            case -1:
                v7 := []int{(fi5 * fi5)}
                _ = v7
            default:
                for fi8 := 0; fi8 < 7; fi8++ {
                    switch ("foo bar" + "") {
                    case "":
                        if (!(68 != 1)) {
                            v9 := GenS2{F0: p2, F1: (!(!true)), F2: -1}
                            _ = v9
                        } else {
                            v10 := []int{(fi8 << ((255) & 63))}
                            v11 := GenS3{F0: 127}
                            _ = v10
                            _ = v11
                        }
                    case "a":
                        switch (^(fi5)) {
                        case 42:
                            v4 = p1
                        default:
                            if (true || ((p2 && false) && true)) {
                                if ((!p2) || (!p2)) {
                                    p2 = p2
                                    switch "a" {
                                    case "foo bar":
                                        fmt.Println(v4, v4)
                                    case "":
                                        v12 := GenS4{F0: 1}
                                        _ = v12
                                    case "a":
                                        switch (^(fi5)) {
                                        case 0:
                                            fmt.Println(fi8, p1, fi5)
                                        case -2147483648:
                                            v13 := GenS1{F0: 1, F1: 4294967295, F2: fi5}
                                            _ = v13
                                        }
                                    default:
                                        var v14 float64 = float64(0.0)
                                        _ = v14
                                    }
                                }
                                v4 -= 0
                                if (!true) {
                                    var v15 uint32 = 0
                                    switch "a" {
                                    case "a":
                                        if false {
                                            v16 := GenS5{F0: 4294967295, F1: 5, F2: 42}
                                            if ((3.14159 + ((((((0.1 - (1e10 / 1.5)) + float32(0.1)) + -2.25) - 0.0) * 0.1) / 1.0)) != -2.25) {
                                                p1 -= (p1 * p1)
                                                fmt.Println(v4, p1, fi5)
                                                v17 := GenS2{F0: (v15 > 9), F1: ((-887 < v4) || p2), F2: 0}
                                                _ = v17
                                            } else {
                                                v15 -= (v15 / ((2147483648) | 1))
                                                var v18 int64 = -1
                                                v19 := []int{fi5}
                                                if true {
                                                    var v20 bool = (p2 && (!true))
                                                    if true {
                                                        var v21 int16 = 42
                                                        _ = v21
                                                    } else {
                                                        var v22 uint16 = 2
                                                        p2 = p2
                                                        _ = v22
                                                    }
                                                    if p2 {
                                                        p1 = v18
                                                        var v23 float32 = float32(3.14159)
                                                        _ = v23
                                                    } else {
                                                        _ = len(v19)
                                                    }
                                                    _ = v20
                                                } else {
                                                    var v24 float32 = float32((0.0 * 0.1))
                                                    if len(v19) >= 2 {
                                                        v19 = v19[1:len(v19)]
                                                    }
                                                    _ = v24
                                                }
                                                _ = v18
                                                _ = v19
                                            }
                                            fmt.Println(p1)
                                            _ = v16
                                        }
                                    }
                                    var v25 bool = p2
                                    fmt.Println(v15, p1, p2)
                                    _ = v15
                                    _ = v25
                                }
                                if (0 == -627) {
                                    switch "" {
                                    case "\n\t":
                                        v26 := GenS5{F0: fi5, F1: 18446744073709551615, F2: -1}
                                        _ = v26
                                    default:
                                        var v27 float64 = (-2.25 * 1.5)
                                        _ = v27
                                    }
                                } else {
                                    switch "" {
                                    case "":
                                        var v28 uint = 18446744073709551615
                                        var v29 uint16 = 42
                                        v30 := GenS3{F0: 1}
                                        if p2 {
                                            v4 = (v4 << ((65535) & 63))
                                            if (!(p0 <= uint8(v28))) {
                                                v31 := GenS3{F0: fi5}
                                                p1 -= (v4 << ((1) & 63))
                                                _ = v31
                                            }
                                            if p2 {
                                                v32 := []int{fi5}
                                                _ = v32
                                            } else {
                                                v29 -= (v29 & 2)
                                                v28 -= (^(v28))
                                                var v33 uint16 = 32767
                                                _ = v33
                                            }
                                            var v34 int8 = 0
                                            v29 += (v29 * 1)
                                            _ = v34
                                        } else {
                                            fmt.Println(v4, v29, p2)
                                            var v35 uint32 = 4294967295
                                            var v36 float64 = (1e10 / 1.5)
                                            _ = v35
                                            _ = v36
                                        }
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    default:
                                        v37 := GenS6{F0: "hello"}
                                        _ = v37
                                    }
                                    p1 += (v4 | p1)
                                    var v38 int16 = 32767
                                    _ = v38
                                }
                                var v39 int8 = 1
                                v40 := GenS7{F0: 0, F1: -128, F2: (p1 % ((42) | 1))}
                                _ = v39
                                _ = v40
                            }
                        }
                    case "hello":
                        v4 = 783
                    default:
                        v41 := GenS2{F0: (true && (4294967295 >= 4294967295)), F1: (2147483647 >= -38), F2: -290}
                        _ = v41
                    }
                    var v42 int8 = 42
                    var v43 int8 = v42
                    _ = fi8
                    _ = v42
                    _ = v43
                }
            }
            var v44 float64 = 1e10
            _ = fi5
            _ = v44
        }
    } else {
        v45 := []int{81, 128, -9223372036854775808}
        for fi46 := 0; fi46 < 20; fi46++ {
            v47 := GenS5{F0: (fi46 & (fi46 | -985)), F1: 4294967295, F2: 108}
            p1 += v4
            _ = fi46
            _ = v47
        }
        _ = v45
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (!p2)
}
func genFn1() bool {
    if true {
        var v48 uint64 = 42
        _ = v48
    } else {
        v49 := GenS2{F0: true, F1: true, F2: 0}
        var v50 float64 = 0.0
        _ = v49
        _ = v50
    }
    var v51 int32 = 42
    _ = v51
    return false
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 47
    var v3 int = 128
    if true {
        switch (v3 ^ -620) {
        case -9223372036854775808:
            v3 = v3
        case -794:
            v3 = -9223372036854775808
        case 1:
            switch "hello" {
            case "a":
                var v52 uint = 18446744073709551615
                _ = v52
            case "hello":
                if false {
                    switch "" {
                    case "":
                        var v53 uint32 = 1
                        _ = v53
                    case "\n\t":
                        fmt.Println(v3)
                    default:
                        fmt.Println(v3)
                    }
                    var v54 int32 = 777
                    _ = v54
                }
            }
        }
    } else {
        for fi55 := 0; fi55 < 6; fi55++ {
            for fi56 := 0; fi56 < 3; fi56++ {
                var v57 uint16 = 32767
                var v58 uint32 = 354
                fmt.Println(v2, fi56)
                _ = fi56
                _ = v57
                _ = v58
            }
            v59 := v1["a"]
            _ = fi55
            _ = v59
        }
        if true {
            var v60 int = 32767
            for fi61 := 0; fi61 < 20; fi61++ {
                var v62 uint = 65535
                var v63 float64 = 1.5
                if false {
                    v64 := GenS1{F0: 21, F1: 518, F2: (v60 - (^(v3)))}
                    if (false && true) {
                        if (-602 < 668) {
                            for fi65 := 0; fi65 < 1; fi65++ {
                                v2 = 239
                                _ = fi65
                            }
                            v1[("hello" + "\n\t")] = -426
                            v63 += 3.14159
                        } else {
                            for fi66 := 0; fi66 < 1; fi66++ {
                                if false {
                                    var v67 uint16 = 1
                                    _ = v67
                                } else {
                                    v68 := GenS5{F0: (^(v3)), F1: (v62 >> ((2) & 63)), F2: -949}
                                    _ = v68
                                }
                                _ = fi66
                            }
                            switch (fi61 << ((1) & 63)) {
                            case -340:
                                for fi69 := 0; fi69 < 4; fi69++ {
                                    v62 += (v62 % (((v62 % (((v62 | (v62 ^ (v62 - 2)))) | 1))) | 1))
                                    v3 = (v3 << (((v62 - v62)) & 63))
                                    _ = fi69
                                }
                            case 963:
                                fmt.Println(v60, v63, v62)
                            }
                            if (!true) {
                                v3 -= fi61
                            }
                        }
                    } else {
                        v2 = v2
                        v70 := GenS4{F0: -128}
                        _ = v70
                    }
                    _ = v64
                }
                _ = fi61
                _ = v62
                _ = v63
            }
            if (900 < 0) {
                v1["hello"] = (-(v3))
                var v71 uint64 = 0
                var v72 uint = 1
                _ = v71
                _ = v72
            } else {
                v3 += (v3 | v3)
                if true {
                    var v73 uint64 = 1
                    if (v2 < 0) {
                        switch "foo bar" {
                        case "hello":
                            if (true || (65535 < 1)) {
                                v74 := GenS8{F0: 1, F1: 4294967295, F2: 2147483647}
                                _ = v74
                            }
                        case "":
                            for fi75 := 0; fi75 < 2; fi75++ {
                                fmt.Println(fi75)
                                v3 += fi75
                                _ = fi75
                            }
                        case "\n\t":
                            if (2147483647 == 2) {
                                var v76 uint32 = 42
                                _ = v76
                            }
                        default:
                            for fi77 := 0; fi77 < 11; fi77++ {
                                for fi78 := 0; fi78 < 13; fi78++ {
                                    v73 = (v73 << ((1) & 63))
                                    _ = fi78
                                }
                                v79 := []int{v60, (v3 % (((v3 >> ((4294967295) & 63))) | 1))}
                                _ = fi77
                                _ = v79
                            }
                        }
                        var v80 uint16 = 812
                        _ = v80
                    }
                    _ = v73
                }
                v2 = v2
            }
            _ = v60
        }
        switch (-(v3)) {
        case -1:
            v81 := v1["a"]
            _ = v81
        case 2147483647:
            v3 += (v3 ^ v3)
        case 774:
            fmt.Println(v3, v3)
        }
    }
    var v82 uint8 = v2
    for fi83 := 0; fi83 < 8; fi83++ {
        if true {
            var v84 int8 = 125
            for fi85 := 0; fi85 < 7; fi85++ {
                var v86 uint16 = 2
                _ = len(v1)
                v82 = (v82 * (v82 >> ((42) & 63)))
                fmt.Println(v2, fi85)
                v87 := GenS3{F0: fi85}
                fmt.Println(v86, fi85)
                _ = fi85
                _ = v86
                _ = v87
            }
            v88 := GenS9{F0: 9223372036854775807, F1: (v84 | -1)}
            _ = v84
            _ = v88
        } else {
            var v89 bool = (true || false)
            v3 -= fi83
            _ = v89
        }
        var v90 int32 = int32(v82)
        switch "hello" {
        case "\n\t":
            var v91 uint32 = 1
            _ = v91
        }
        _ = fi83
        _ = v90
    }
    var v92 int = v3
    if (42 <= -1) {
        var v93 uint64 = 865
        _ = v93
    } else {
        v94 := []int{86}
        for fi95 := 0; fi95 < 6; fi95++ {
            var v96 string = ("\n\t" + (("\n\t" + ("" + (("hello" + "\n\t") + "hello"))) + ""))
            for fi97 := 0; fi97 < 9; fi97++ {
                var v98 int8 = 127
                _ = fi97
                _ = v98
            }
            _ = len(v1)
            _ = fi95
            _ = v96
        }
        var v99 uint16 = 476
        _ = v94
        _ = v99
    }
    var v100 bool = (true || true)
    v101 := GenS10{F0: 127, F1: 517}
    switch (^(v92)) {
    case 4294967295:
        var v102 float32 = 1.5
        _ = v102
    default:
        v103 := GenS4{F0: -75}
        _ = v103
    }
    fmt.Println(v2, v82)
    if v100 {
        if (42 < 0) {
            v3 = (^(v92))
            v1["hello"] = (v3 >> ((42) & 63))
            switch v92 {
            case -93:
                _ = len(v1)
            }
        }
        v104 := []int{v3, (v92 >> ((371) & 63)), -9223372036854775808}
        _ = v104
    } else {
        for fi105 := 0; fi105 < 16; fi105++ {
            v3 = v92
            _ = fi105
        }
        var v106 int32 = -1
        _ = v106
    }
    var v107 uint = 588
    var v108 float64 = (3.14159 + 0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v82:", v82)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v100
    _ = v101
    _ = v107
    _ = v108
    _ = v2
    _ = v3
    _ = v82
    _ = v92
}
