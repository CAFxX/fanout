package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int32, p1 int32, p2 uint8) bool {
    v5 := GenS1{F0: 42, F1: uint(p2), F2: (p1 >> ((920) & 63))}
    if (true || (!((p2 & (p2 % ((p2) | 1))) > p2))) {
        v6 := GenS1{F0: 18446744073709551615, F1: 1, F2: p0}
        var v7 int8 = 0
        switch -2147483648 {
        case 0:
            var v8 uint = 42
            _ = v8
        case 4294967295:
            v9 := []int{4294967295}
            _ = v9
        default:
            var v10 int64 = 42
            _ = v10
        }
        _ = v6
        _ = v7
    } else {
        var v11 float64 = 0.1
        fmt.Println(p1, p1)
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        fmt.Println(fi12, p1, p0)
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (((("foo bar" + ("" + "a")) + "") > "a") || true)
}
func genFn1(p0 uint, p1 uint64) uint {
    v13 := GenS2{F0: 26}
    p0 += p0
    _ = p0
    _ = p1
    _ = v13
    return p0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]string)
    var v3 bool = (!(!true))
    var v4 float64 = 0.0
    v14 := []int{32768}
    _ = len(v2)
    fmt.Println(v4)
    if ((3.14159 / 1.0) >= (1e10 - 0.0)) {
        if len(v14) > 0 {
            v14[0] = 9223372036854775807
        }
        v3 = (259 >= -1)
    } else {
        var v15 bool = v3
        _ = v15
    }
    if (32767 <= -437) {
        for fi16 := 0; fi16 < 19; fi16++ {
            if false {
                if (18446744073709551615 <= 255) {
                    var v17 uint = 18446744073709551615
                    switch "a" {
                    case "":
                        var v18 int = (fi16 % ((-755) | 1))
                        _ = v18
                    case "foo bar":
                        switch "a" {
                        case "a":
                            v19 := GenS2{F0: int8(v17)}
                            _ = v19
                        case "":
                            if (9223372036854775807 < fi16) {
                                v20 := GenS2{F0: 42}
                                _ = v20
                            } else {
                                switch fi16 {
                                case 0:
                                    if true {
                                        var v21 float64 = v4
                                        v22 := v2[fi16]
                                        v3 = v3
                                        _ = v21
                                        _ = v22
                                    }
                                case 9223372036854775807:
                                    var v23 uint16 = 0
                                    _ = v23
                                }
                                fmt.Println(v3, v4, v3)
                                v4 = (3.14159 * 0.1)
                            }
                        case "foo bar":
                            if (((("hello" + "\n\t") + "hello") + ("" + "hello")) != "a") {
                                if true {
                                    v3 = v3
                                    if len(v14) > 0 {
                                        v14[0] = 128
                                    }
                                }
                                if (((0.0 / 1.0) >= (0.1 / 1e10)) && (true && false)) {
                                    if v3 {
                                        var v24 uint16 = 457
                                        v25 := GenS1{F0: 4294967295, F1: (v17 % (((v17 % ((genFn1(v17, 42)) | 1))) | 1)), F2: 594}
                                        _ = v24
                                        _ = v25
                                    }
                                    v14 = append(v14, (fi16 - fi16))
                                }
                                var v26 string = ""
                                _ = v26
                            } else {
                                for fi27 := 0; fi27 < 10; fi27++ {
                                    if v3 {
                                        var v28 uint16 = 255
                                        switch -106 {
                                        case -1:
                                            var v29 uint8 = uint8(v28)
                                            _ = v29
                                        case -12:
                                            var v30 float64 = v4
                                            _ = v30
                                        case 727:
                                            if ((false && (v3 || false)) || v3) {
                                                v31 := GenS1{F0: 902, F1: (v17 + 769), F2: -1}
                                                _ = v31
                                            }
                                        default:
                                            if (v28 >= 65535) {
                                                v32 := GenS2{F0: 100}
                                                if false {
                                                    v17 -= (v17 + genFn1(v17, 1))
                                                }
                                                var v33 uint = genFn1((v17 * (v17 * (^(v17)))), 302)
                                                _ = len(v2)
                                                var v34 uint64 = 4294967295
                                                v34 = (v34 & v34)
                                                v14 = append(v14, fi16)
                                                _ = v32
                                                _ = v33
                                                _ = v34
                                            }
                                        }
                                        v3 = v3
                                        _ = v28
                                    }
                                    if v3 {
                                        var v35 float64 = ((v4 * float64(3.14159)) + float64(1.5))
                                        var v36 uint8 = 0
                                        _ = v35
                                        _ = v36
                                    } else {
                                        if ((327 < 1) || ((!v3) || (v3 || (false || (fi16 >= 4294967295))))) {
                                            var v37 int32 = 195
                                            v2[(-(fi27))] = ("foo bar" + ("\n\t" + "hello"))
                                            _ = v37
                                        } else {
                                            _ = len(v14)
                                            var v38 float64 = v4
                                            _ = v38
                                        }
                                        if len(v14) > 0 {
                                            v14[0] = fi27
                                        }
                                        if (!("\n\t" == "\n\t")) {
                                            if false {
                                                var v39 bool = (!(0 != 4294967295))
                                                v39 = v3
                                                v40 := GenS1{F0: 0, F1: 327, F2: -2147483648}
                                                v14 = append(v14, (fi16 - fi16))
                                                v41 := GenS2{F0: 127}
                                                _ = v39
                                                _ = v40
                                                _ = v41
                                            } else {
                                                var v42 float32 = 0.1
                                                switch "hello" {
                                                case "foo bar":
                                                    v43 := GenS2{F0: int8(v17)}
                                                    switch fi27 {
                                                    case 128:
                                                        var v44 int16 = 32767
                                                        fmt.Println(v3)
                                                        _ = v44
                                                    default:
                                                        fmt.Println(fi16)
                                                    }
                                                    if false {
                                                        var v45 string = ("hello" + ("foo bar" + "a"))
                                                        _ = len(v2)
                                                        v2[fi16] = "hello"
                                                        _ = v45
                                                    } else {
                                                        v46 := GenS2{F0: -128}
                                                        v42 = (v42 / (v42 + v42))
                                                        _ = v46
                                                    }
                                                    delete(v1, 1)
                                                    _ = v43
                                                case "a":
                                                    var v47 int = -9223372036854775808
                                                    _ = v47
                                                default:
                                                    v1[1] = (fi16 + 2147483647)
                                                }
                                                var v48 float64 = (v4 / (0.0 / (0.0 * (((v4 + v4) + (1.5 + 1e10)) - v4))))
                                                v48 -= (v48 / 0.1)
                                                var v49 bool = v3
                                                _ = v42
                                                _ = v48
                                                _ = v49
                                            }
                                        }
                                    }
                                    fmt.Println(fi27)
                                    v50 := GenS1{F0: 0, F1: (v17 << ((18446744073709551615) & 63)), F2: 345}
                                    _ = fi27
                                    _ = v50
                                }
                                v3 = v3
                                var v51 string = "hello"
                                _ = v51
                            }
                        default:
                            v52 := GenS2{F0: -37}
                            _ = v52
                        }
                    case "\n\t":
                        var v53 int32 = -2147483648
                        _ = v53
                    default:
                        for fi54 := 0; fi54 < 16; fi54++ {
                            switch 42 {
                            case 1:
                                v17 = v17
                                var v55 int64 = 106
                                var v56 int64 = -9223372036854775808
                                _ = v55
                                _ = v56
                            case -9223372036854775808:
                                var v57 uint32 = 4294967295
                                _ = v57
                            }
                            var v58 uint64 = 0
                            if false {
                                _ = len(v14)
                            } else {
                                v3 = (("hello" != "foo bar") || v3)
                                fmt.Println(v4, v58, fi16)
                            }
                            _ = fi54
                            _ = v58
                        }
                    }
                    switch 4294967295 {
                    case -554:
                        var v59 uint16 = 32767
                        _ = v59
                    case 0:
                        var v60 uint8 = 85
                        _ = v60
                    case 128:
                        v61 := GenS1{F0: 435, F1: genFn1((v17 * (v17 & genFn1((v17 >> ((v17) & 63)), 18446744073709551615))), uint64(v17)), F2: 1}
                        _ = v61
                    }
                    _ = v17
                } else {
                    var v62 int32 = 2147483647
                    _ = v62
                }
            }
            _ = fi16
        }
    } else {
        v3 = v3
        switch "" {
        case "hello":
            var v63 uint8 = 42
            _ = v63
        default:
            var v64 uint64 = 4294967295
            _ = v64
        }
    }
    v65 := GenS1{F0: 42, F1: 4294967295, F2: -2147483648}
    v3 = true
    if ((!(2147483648 <= 2)) && (((0.0 * 3.14159) / 1e10) <= (0.0 * (((1e10 - 3.14159) + 0.0) * 3.14159)))) {
        fmt.Println(v3, v4)
        var v66 uint8 = 42
        for fi67 := 0; fi67 < 15; fi67++ {
            v66 = 127
            switch 127 {
            case 2147483647:
                fmt.Println(v66)
            case 68:
                switch fi67 {
                case 128:
                    if v3 {
                        if (!v3) {
                            switch -1 {
                            case -213:
                                switch "a" {
                                case "\n\t":
                                    var v68 int = (fi67 >> ((642) & 63))
                                    _ = v68
                                case "foo bar":
                                    var v69 string = "\n\t"
                                    _ = v69
                                default:
                                    v3 = v3
                                }
                            case 128:
                                v66 += (^(v66))
                            }
                            fmt.Println(fi67, v3)
                        }
                        switch "\n\t" {
                        case "\n\t":
                            v4 = v4
                        case "a":
                            v70 := GenS2{F0: 100}
                            _ = v70
                        default:
                            var v71 uint = 65535
                            _ = v71
                        }
                    }
                case 42:
                    v72 := GenS3{F0: 1}
                    _ = v72
                default:
                    v73 := GenS4{F0: 2, F1: fi67}
                    _ = v73
                }
            case 0:
                for fi74 := 0; fi74 < 16; fi74++ {
                    if (-2.25 != v4) {
                        if ((!(247 == 2147483648)) || v3) {
                            if v3 {
                                var v75 uint16 = 2
                                var v76 uint16 = (v75 | v75)
                                _ = v75
                                _ = v76
                            } else {
                                var v77 int = (fi67 * 67)
                                var v78 uint = 4294967295
                                _ = v77
                                _ = v78
                            }
                            v79 := v1[187]
                            _ = v79
                        }
                    }
                    v66 = (v66 << ((42) & 63))
                    fmt.Println(v3, v4)
                    var v80 float64 = v4
                    v81 := GenS1{F0: 4294967295, F1: 2, F2: 325}
                    delete(v1, 208)
                    _ = fi74
                    _ = v80
                    _ = v81
                }
            }
            for fi82 := 0; fi82 < 10; fi82++ {
                v2[fi82] = "a"
                _ = fi82
            }
            _ = fi67
        }
        _ = v66
    }
    var v83 int16 = -1
    var v84 float32 = 0.0
    var v85 uint8 = 2
    v86 := GenS3{F0: 0}
    if v3 {
        for fi87 := 0; fi87 < 1; fi87++ {
            if true {
                var v88 string = ("foo bar" + "")
                _ = v88
            } else {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                v2[(fi87 % ((fi87) | 1))] = "foo bar"
            }
            v89 := GenS5{F0: 1}
            _ = fi87
            _ = v89
        }
        if (16 > 100) {
            v85 = v85
            for fi90 := 0; fi90 < 9; fi90++ {
                switch ("a" + "a") {
                case "":
                    switch "\n\t" {
                    case "hello":
                        v85 -= (^(v85))
                    case "a":
                        v83 = (-(v83))
                    default:
                        v91 := GenS2{F0: -128}
                        _ = v91
                    }
                }
                _ = fi90
            }
            for fi92 := 0; fi92 < 15; fi92++ {
                for fi93 := 0; fi93 < 5; fi93++ {
                    var v94 uint16 = 1
                    fmt.Println(fi92)
                    if (18446744073709551615 == 0) {
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        fmt.Println(v3, v84, fi92)
                        if v3 {
                            var v95 int8 = 1
                            var v96 int16 = (v83 << ((22) & 63))
                            switch fi92 {
                            case 32767:
                                var v97 int16 = v96
                                _ = v97
                            default:
                                v98 := GenS6{F0: 1}
                                _ = v98
                            }
                            _ = v95
                            _ = v96
                        }
                    }
                    _ = fi93
                    _ = v94
                }
                v99 := v2[(^(fi92))]
                fmt.Println(fi92, v99)
                _ = fi92
                _ = v99
            }
        }
    }
    if (128 >= v85) {
        switch 32767 {
        case 4294967295:
            v100 := GenS3{F0: 9223372036854775807}
            _ = v100
        }
    }
    if (((v85 >= (^(v85))) || false) && (!(359 >= 0))) {
        var v101 int32 = 367
        var v102 int16 = -1
        _ = v101
        _ = v102
    }
    for fi103 := 0; fi103 < 5; fi103++ {
        if len(v14) > 0 {
            v14[0] = (^(fi103))
        }
        switch "a" {
        case "hello":
            switch fi103 {
            case 630:
                var v104 int64 = 812
                _ = v104
            default:
                v85 -= 224
            }
        default:
            var v105 uint32 = 65535
            _ = v105
        }
        _ = fi103
    }
    var v106 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v106:", v106)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v106
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v65
    _ = v83
    _ = v84
    _ = v85
    _ = v86
}
