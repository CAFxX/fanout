package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint8; F1 bool; F2 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8; F1 int16; F2 uint }
type GenS6 struct { F0 int; F1 uint32; F2 int64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int64; F1 string; F2 string }
type GenS9 struct { F0 uint32; F1 uint32 }
type GenS10 struct { F0 int16; F1 uint16 }
type GenS11 struct { F0 int; F1 int64; F2 string }
type GenS12 struct { F0 uint32; F1 uint8 }
type GenS13 struct { F0 string; F1 int }
type GenS14 struct { F0 int; F1 bool }
func genFn0(p0 int64, p1 uint) bool {
    var v5 float64 = 0.0
    v6 := GenS1{F0: 0}
    if true {
        var v7 uint16 = 2
        _ = v7
    } else {
        p0 += -9223372036854775808
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (!(42 != 42))
}
func genFn1(p0 int, p1 uint32, p2 uint) uint {
    if (true && (!(true && true))) {
        switch (-(p0)) {
        case -511:
            var v8 uint8 = 0
            _ = v8
        case 4294967295:
            v9 := []int{(-(p0)), -38, (-(p0))}
            _ = v9
        default:
            p2 += (p2 - (p2 | p2))
        }
        v10 := GenS1{F0: 2}
        var v11 uint64 = 95
        _ = v10
        _ = v11
    } else {
        v12 := []int{p0, p0, (p0 & p0)}
        for fi13 := 0; fi13 < 6; fi13++ {
            for fi14 := 0; fi14 < 18; fi14++ {
                p0 = int(p1)
                if true {
                    if (9223372036854775807 >= 252) {
                        v15 := GenS1{F0: 127}
                        p1 -= p1
                        switch (-(fi13)) {
                        case 384:
                            p1 = p1
                        case -1:
                            fmt.Println(p1, p0, fi13)
                        case 217:
                            switch (("a" + "\n\t") + ("hello" + "")) {
                            case "hello":
                                v16 := GenS1{F0: 42}
                                _ = v16
                            case "a":
                                v17 := GenS2{F0: -952, F1: (float32((float64(float32(0.1)) * 1e10)) < 1e10)}
                                _ = v17
                            case "foo bar":
                                var v18 uint32 = (^(p1))
                                _ = v18
                            }
                        }
                        _ = v15
                    } else {
                        p2 += 42
                        if (42 >= 255) {
                            if false {
                                var v19 float32 = 0.1
                                v20 := GenS2{F0: -2147483648, F1: true}
                                v21 := GenS2{F0: 0, F1: (false && true)}
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            if false {
                                fmt.Println(p0, p0, fi14)
                                if (-74 <= int8(p2)) {
                                    v22 := GenS3{F0: 128, F1: false, F2: (fi14 & (fi13 / (((fi13 | fi13)) | 1)))}
                                    v23 := GenS2{F0: -1, F1: (true || false)}
                                    if false {
                                        var v24 string = "hello"
                                        v24 = v24
                                        _ = v24
                                    }
                                    _ = v22
                                    _ = v23
                                } else {
                                    v25 := GenS3{F0: 255, F1: true, F2: (-(fi14))}
                                    switch p0 {
                                    case -560:
                                        var v26 int8 = 100
                                        _ = v26
                                    case -133:
                                        if true {
                                            switch ("\n\t" + ("a" + ("foo bar" + "a"))) {
                                            case "foo bar":
                                                p2 = 0
                                                v27 := GenS4{F0: -914}
                                                p2 += 561
                                                _ = v27
                                            case "":
                                                v28 := GenS5{F0: -1, F1: -274, F2: 65535}
                                                _ = v28
                                            case "hello":
                                                if len(v12) > 0 {
                                                    v12[0] = int(p2)
                                                }
                                            }
                                        } else {
                                            _ = len(v12)
                                            v12 = append(v12, fi13)
                                        }
                                    default:
                                        p2 -= (p2 / (((p2 * 4294967295)) | 1))
                                    }
                                    if false {
                                        fmt.Println(p1, fi13, fi14)
                                        if true {
                                            var v29 uint = (p2 % ((p2) | 1))
                                            p0 -= (-(fi13))
                                            _ = v29
                                        }
                                        v30 := GenS3{F0: 238, F1: true, F2: -592}
                                        _ = v30
                                    }
                                    var v31 string = ("hello" + ("" + "hello"))
                                    p1 = (p1 * p1)
                                    fmt.Println(p1, p0)
                                    fmt.Println(fi14, v31, fi14)
                                    p0 = fi14
                                    var v32 int = 42
                                    _ = v25
                                    _ = v31
                                    _ = v32
                                }
                                v33 := GenS6{F0: fi14, F1: (^(p1)), F2: -9223372036854775808}
                                _ = v33
                            } else {
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            }
                            fmt.Println(p0)
                        }
                    }
                    p1 -= (^(p1))
                    var v34 int32 = -1
                    _ = v34
                } else {
                    v35 := GenS1{F0: 157}
                    if false {
                        var v36 int64 = 1
                        v37 := GenS4{F0: 1000}
                        var v38 int16 = 1
                        v39 := GenS7{F0: 777}
                        if len(v12) > 0 {
                            v12[0] = (-(p0))
                        }
                        _ = v36
                        _ = v37
                        _ = v38
                        _ = v39
                    } else {
                        var v40 float64 = -2.25
                        var v41 int64 = 42
                        if true {
                            switch fi13 {
                            case 4294967295:
                                _ = len(v12)
                            case 132:
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                if false {
                                    if len(v12) > 0 {
                                        v12[0] = fi13
                                    }
                                }
                            case 625:
                                p2 = p2
                            }
                            var v42 float32 = ((((0.1 / -2.25) * 0.1) * -2.25) / float32((1e10 * v40)))
                            var v43 uint = p2
                            switch "hello" {
                            case "hello":
                                if true {
                                    if len(v12) > 0 {
                                        v12[0] = 127
                                    }
                                    fmt.Println(v43)
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    if false {
                                        if (2 == 2) {
                                            p2 -= 4294967295
                                        }
                                        v43 = 4294967295
                                        switch fi13 {
                                        case 128:
                                            switch (fi13 % (((p0 + (p0 >> (((^(v43))) & 63)))) | 1)) {
                                            case 819:
                                                var v44 int8 = 42
                                                _ = v44
                                            case 42:
                                                v41 = -9223372036854775808
                                            case -778:
                                                p1 = (p1 / (((p1 % ((p1) | 1))) | 1))
                                            default:
                                                var v45 uint32 = (^(p1))
                                                _ = v45
                                            }
                                        case -477:
                                            switch ("hello" + "foo bar") {
                                            case "\n\t":
                                                v46 := GenS2{F0: 0, F1: true}
                                                fmt.Println(v40, p1, v41)
                                                _ = v46
                                            }
                                        default:
                                            v47 := GenS5{F0: 127, F1: -33, F2: v43}
                                            _ = v47
                                        }
                                    } else {
                                        var v48 bool = ((v43 ^ (^(p2))) > (v43 * (p2 - v43)))
                                        _ = v48
                                    }
                                    if (0.0 < v42) {
                                        _ = len(v12)
                                    } else {
                                        var v49 string = "a"
                                        if true {
                                            v50 := GenS8{F0: 9223372036854775807, F1: (v49 + (((v49 + "hello") + v49) + (((("\n\t" + "foo bar") + ("foo bar" + (((v49 + (("foo bar" + ("\n\t" + v49)) + v49)) + "") + (v49 + ("foo bar" + v49))))) + "\n\t") + v49))), F2: v49}
                                            if true {
                                                v40 -= float64(v42)
                                            }
                                            _ = v50
                                        }
                                        if (42 > -596) {
                                            var v51 uint32 = (^(p1))
                                            _ = v51
                                        } else {
                                            v52 := GenS9{F0: p1, F1: (p1 ^ p1)}
                                            switch fi14 {
                                            case 42:
                                                var v53 float32 = 0.1
                                                _ = v53
                                            case 4294967295:
                                                v49 = v49
                                                switch fi14 {
                                                case 128:
                                                    switch (v49 + "a") {
                                                    case "a":
                                                        if len(v12) >= 2 {
                                                            v12 = v12[1:len(v12)]
                                                        }
                                                    case "\n\t":
                                                        if len(v12) > 0 {
                                                            v12[0] = fi14
                                                        }
                                                    }
                                                case 4294967295:
                                                    v42 += 1e10
                                                }
                                                var v54 uint32 = 42
                                                _ = v54
                                            case 0:
                                                v43 += 1
                                                p2 = v43
                                                var v55 uint16 = 1
                                                _ = v55
                                            default:
                                                if len(v12) > 0 {
                                                    v12[0] = fi14
                                                }
                                            }
                                            _ = v52
                                        }
                                        _ = v49
                                    }
                                }
                            }
                            var v56 int64 = v41
                            _ = v42
                            _ = v43
                            _ = v56
                        } else {
                            var v57 int16 = 0
                            var v58 uint = 42
                            var v59 bool = true
                            _ = v57
                            _ = v58
                            _ = v59
                        }
                        if (255 == p2) {
                            fmt.Println(v40, fi14)
                            fmt.Println(v41, p1)
                        }
                        if (true || false) {
                            switch (fi13 + -1) {
                            case -143:
                                fmt.Println(v40, p0)
                            }
                            v40 = (float64(float32(1.5)) - ((v40 - (v40 - float64(0.1))) - v40))
                        }
                        _ = v40
                        _ = v41
                    }
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    _ = v35
                }
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = fi14
            }
            _ = fi13
        }
        var v60 int64 = 0
        _ = v12
        _ = v60
    }
    var v61 int8 = 107
    var v62 float32 = 1e10
    for fi63 := 0; fi63 < 10; fi63++ {
        fmt.Println(v61)
        _ = fi63
    }
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v61
    _ = v62
    return (p2 >> ((p2) & 63))
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    v3 := []int{}
    var v4 bool = (2 <= 4294967295)
    v4 = v4
    v64 := GenS8{F0: 909, F1: "hello", F2: (("hello" + "") + "\n\t")}
    var v65 uint = 90
    if (32767 > 1) {
        var v66 uint8 = uint8(v65)
        v66 = 42
        v67 := GenS10{F0: -1, F1: 27}
        _ = v66
        _ = v67
    } else {
        var v68 float32 = 1e10
        _ = v68
    }
    var v69 uint16 = 0
    var v70 int16 = int16(v65)
    fmt.Println(v65, v65, v65)
    v70 = v70
    v71 := GenS11{F0: 207, F1: 42, F2: "hello"}
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    fmt.Println(v4, v70, v70)
    if (-95 > -1) {
        if (1 != 1) {
            if true {
                for fi72 := 0; fi72 < 14; fi72++ {
                    fmt.Println(v65, v65)
                    _ = fi72
                }
                v4 = v4
                v4 = (!(v4 && (true && v4)))
            } else {
                var v73 int = 117
                if v4 {
                    var v74 float32 = float32(3.14159)
                    v69 = v69
                    switch ("" + "\n\t") {
                    case "\n\t":
                        var v75 uint = 18446744073709551615
                        _ = v75
                    }
                    _ = v74
                }
                for fi76 := 0; fi76 < 17; fi76++ {
                    for fi77 := 0; fi77 < 15; fi77++ {
                        v78 := GenS12{F0: 2147483647, F1: 87}
                        v79 := GenS6{F0: v73, F1: 2147483647, F2: 62}
                        var v80 uint = (v65 >> (((v65 >> ((232) & 63))) & 63))
                        _ = fi77
                        _ = v78
                        _ = v79
                        _ = v80
                    }
                    v73 = int(v70)
                    if v4 {
                        switch "" {
                        case "":
                            if v4 {
                                v65 -= 18446744073709551615
                            }
                        case "a":
                            if v4 {
                                fmt.Println(v4, v69)
                            }
                        case "foo bar":
                            var v81 float32 = -2.25
                            _ = v81
                        }
                    } else {
                        v69 += (v69 + 255)
                        if false {
                            if len(v3) > 0 {
                                v3[0] = fi76
                            }
                            for fi82 := 0; fi82 < 3; fi82++ {
                                v69 = (^(v69))
                                v73 = (fi82 + 128)
                                fmt.Println(v65)
                                var v83 uint32 = 949
                                var v84 int64 = 9223372036854775807
                                _ = fi82
                                _ = v83
                                _ = v84
                            }
                        }
                        var v85 float64 = 1.5
                        _ = v85
                    }
                    _ = fi76
                }
                _ = v73
            }
        }
        if false {
            for fi86 := 0; fi86 < 11; fi86++ {
                v87 := GenS2{F0: 1, F1: (!v4)}
                v88 := GenS13{F0: "foo bar", F1: (fi86 >> ((genFn1((fi86 / (((^(fi86))) | 1)), 65535, (v65 % ((uint(v69)) | 1)))) & 63))}
                switch "foo bar" {
                case "foo bar":
                    var v89 int8 = -1
                    _ = v89
                default:
                    switch "foo bar" {
                    case "hello":
                        if ((945 != 4294967295) && v4) {
                            v69 = 987
                            v70 -= 42
                        } else {
                            v3 = append(v3, fi86)
                        }
                    case "a":
                        for fi90 := 0; fi90 < 15; fi90++ {
                            var v91 uint8 = 2
                            switch "foo bar" {
                            case "foo bar":
                                v70 -= (v70 - 42)
                            case "":
                                var v92 float32 = 1e10
                                _ = v92
                            case "hello":
                                v93 := GenS1{F0: 1}
                                _ = v93
                            }
                            v4 = v4
                            _ = fi90
                            _ = v91
                        }
                    }
                }
                _ = fi86
                _ = v87
                _ = v88
            }
            if (true && (!((v65 << (((^(v65))) & 63)) <= 487))) {
                v94 := GenS9{F0: 2, F1: 42}
                _ = v94
            } else {
                var v95 int32 = 1
                if (!(!(3.14159 > 1e10))) {
                    v65 += 1
                }
                switch -771 {
                case -331:
                    fmt.Println(v65, v65)
                case 994:
                    for fi96 := 0; fi96 < 14; fi96++ {
                        if ("hello" < "\n\t") {
                            var v97 float32 = 1e10
                            v98 := GenS14{F0: (^(fi96)), F1: (false && true)}
                            v3 = append(v3, fi96)
                            _ = v97
                            _ = v98
                        }
                        var v99 uint64 = 42
                        _ = fi96
                        _ = v99
                    }
                default:
                    v95 = (v95 * v95)
                }
                _ = v95
            }
        } else {
            switch 128 {
            case 4294967295:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                fmt.Println(v65, v70, v4)
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v65:", v65)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v65
    _ = v69
    _ = v70
    _ = v71
}
