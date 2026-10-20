package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint; F1 uint8; F2 string }
type GenS4 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint; F1 uint32; F2 string }
func genFn0(p0 bool, p1 bool, p2 bool) int8 {
    var v6 int = 422
    if p2 {
        var v7 uint8 = 127
        var v8 int16 = -32768
        if ((0 < 75) && (!(!p0))) {
            for fi9 := 0; fi9 < 19; fi9++ {
                fmt.Println(fi9, v7, v8)
                var v10 int32 = 42
                _ = fi9
                _ = v10
            }
        }
        _ = v7
        _ = v8
    } else {
        for fi11 := 0; fi11 < 12; fi11++ {
            if false {
                for fi12 := 0; fi12 < 8; fi12++ {
                    p0 = p1
                    p1 = ((fi12 <= fi11) || true)
                    _ = fi12
                }
            } else {
                fmt.Println(p2, fi11, p1)
                p0 = false
                fmt.Println(p1, p0)
            }
            v13 := []int{}
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 bool = p1
                if len(v13) > 0 {
                    v13[0] = 42
                }
                _ = fi14
                _ = v15
            }
            _ = fi11
            _ = v13
        }
        v16 := []int{v6, -9223372036854775808}
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}
func genFn1() uint32 {
    if false {
        var v17 int = 807
        if false {
            switch ("hello" + (("" + "") + "a")) {
            case "hello":
                if true {
                    v18 := GenS1{F0: 42, F1: 0}
                    _ = v18
                }
            case "":
                v19 := []int{v17, -155, v17, (-(v17))}
                _ = v19
            case "a":
                if (true && (!true)) {
                    var v20 int = v17
                    switch "foo bar" {
                    case "\n\t":
                        v21 := GenS1{F0: 255, F1: 858}
                        _ = v21
                    default:
                        var v22 bool = true
                        _ = v22
                    }
                    _ = v20
                }
            }
        }
        var v23 bool = true
        _ = v17
        _ = v23
    }
    v24 := GenS1{F0: 2, F1: -1}
    if (true || false) {
        v25 := GenS2{F0: (((("a" + "hello") + ("foo bar" + ("" + "\n\t"))) + "\n\t") + "foo bar"), F1: 0}
        _ = v25
    } else {
        for fi26 := 0; fi26 < 7; fi26++ {
            var v27 uint = 218
            switch (fi26 | (fi26 * (fi26 & (^(fi26))))) {
            case 1:
                for fi28 := 0; fi28 < 8; fi28++ {
                    v27 -= (v27 % (((v27 / ((0) | 1))) | 1))
                    _ = fi28
                }
            case 436:
                fmt.Println(fi26, fi26)
            case 127:
                switch (fi26 - (^(fi26))) {
                case 142:
                    var v29 int32 = 0
                    _ = v29
                case 32767:
                    if false {
                        fmt.Println(fi26, v27)
                    } else {
                        v30 := []int{-1, (fi26 >> ((v27) & 63)), fi26}
                        var v31 float64 = 0.0
                        _ = v30
                        _ = v31
                    }
                case -182:
                    for fi32 := 0; fi32 < 11; fi32++ {
                        switch -746 {
                        case -31:
                            v33 := []int{}
                            _ = v33
                        default:
                            if true {
                                switch fi32 {
                                case 128:
                                    v34 := GenS3{F0: (v27 / ((42) | 1)), F1: 255, F2: ("\n\t" + "\n\t")}
                                    _ = v34
                                }
                                v35 := GenS4{F0: 2, F1: 1, F2: 0}
                                _ = v35
                            } else {
                                v36 := GenS5{F0: true}
                                v27 = (v27 % ((65535) | 1))
                                var v37 uint16 = 32767
                                _ = v36
                                _ = v37
                            }
                        }
                        v27 -= (v27 + 408)
                        fmt.Println(fi32, v27)
                        _ = fi32
                    }
                default:
                    v38 := []int{}
                    _ = v38
                }
            }
            var v39 int32 = 490
            _ = fi26
            _ = v27
            _ = v39
        }
        var v40 float32 = 0.0
        _ = v40
    }
    _ = v24
    return 2147483648
}

func main() {
    var v1 int64 = -1
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    var v4 string = "hello"
    v5 := make(map[int]string)
    if (118 > 589) {
        if (!(!false)) {
            switch 901 {
            case -2147483648:
                v41 := GenS1{F0: 2, F1: 42}
                _ = v41
            case 0:
                var v42 uint64 = 4294967295
                _ = v42
            }
            for fi43 := 0; fi43 < 15; fi43++ {
                v1 = (v1 - v1)
                for fi44 := 0; fi44 < 20; fi44++ {
                    var v45 int32 = -1
                    fmt.Println(v4)
                    v5[(fi44 << ((424) & 63))] = v4
                    _ = fi44
                    _ = v45
                }
                _ = fi43
            }
        } else {
            v1 = (v1 ^ v1)
            if false {
                fmt.Println(v1)
                var v46 int = -9223372036854775808
                _ = v46
            } else {
                v47 := []int{}
                _ = v47
            }
            var v48 int8 = 1
            _ = v48
        }
        var v49 uint64 = 20
        _ = v49
    } else {
        var v50 float64 = 1e10
        var v51 uint32 = 2147483647
        _ = v50
        _ = v51
    }
    _ = len(v5)
    var v52 string = "\n\t"
    var v53 uint32 = 345
    if true {
        if (-1000 <= 32767) {
            if false {
                switch "hello" {
                case "":
                    v54 := v3[-851]
                    _ = v54
                }
                var v55 uint8 = 127
                _ = v55
            } else {
                var v56 float32 = float32(3.14159)
                if (v56 > -2.25) {
                    var v57 string = ""
                    if true {
                        for fi58 := 0; fi58 < 11; fi58++ {
                            var v59 int64 = -720
                            if (255 != 391) {
                                fmt.Println(v1, v53)
                                if (genFn1() < 2) {
                                    v60 := []int{}
                                    _ = v60
                                }
                                v4 = v52
                            } else {
                                v57 = "foo bar"
                                v61 := []int{fi58, fi58}
                                delete(v3, (-(fi58)))
                                _ = v61
                            }
                            fmt.Println(v52, v57)
                            _ = fi58
                            _ = v59
                        }
                    } else {
                        fmt.Println(v53, v57)
                        var v62 bool = true
                        if true {
                            for fi63 := 0; fi63 < 6; fi63++ {
                                for fi64 := 0; fi64 < 10; fi64++ {
                                    v65 := GenS3{F0: 15, F1: 128, F2: "a"}
                                    _ = fi64
                                    _ = v65
                                }
                                _ = fi63
                            }
                            v66 := []int{-294, 127}
                            v53 = genFn1()
                            _ = v66
                        } else {
                            switch -2147483648 {
                            case -1:
                                if (!true) {
                                    v62 = (42 != -581)
                                    v67 := GenS5{F0: (false || v62)}
                                    var v68 float32 = v56
                                    _ = v67
                                    _ = v68
                                }
                            case 42:
                                fmt.Println(v62, v57, v4)
                            case 128:
                                switch "hello" {
                                case "foo bar":
                                    v69 := []int{0, 346, 2147483647, 9223372036854775807}
                                    _ = v69
                                case "a":
                                    switch v52 {
                                    case "foo bar":
                                        if (false && true) {
                                            v70 := GenS6{F0: -733}
                                            if true {
                                                v56 -= (((v56 - (v56 + 3.14159)) + float32(float64(0.0))) + v56)
                                                for fi71 := 0; fi71 < 19; fi71++ {
                                                    v72 := []int{(fi71 % (((fi71 % (((fi71 | fi71)) | 1))) | 1)), 0, (fi71 >> ((4294967295) & 63))}
                                                    _ = fi71
                                                    _ = v72
                                                }
                                                if (2147483647 < -2147483648) {
                                                    switch -1 {
                                                    case 654:
                                                        for fi73 := 0; fi73 < 4; fi73++ {
                                                            var v74 uint16 = 683
                                                            for fi75 := 0; fi75 < 13; fi75++ {
                                                                v53 = v53
                                                                v76 := GenS7{F0: 127}
                                                                _ = fi75
                                                                _ = v76
                                                            }
                                                            _ = fi73
                                                            _ = v74
                                                        }
                                                    case 127:
                                                        v53 = (^(v53))
                                                    }
                                                    v1 -= -589
                                                    var v77 uint32 = (v53 << ((255) & 63))
                                                    _ = v77
                                                }
                                            }
                                            _ = v70
                                        }
                                    case "\n\t":
                                        v78 := GenS4{F0: 2, F1: 100, F2: 2}
                                        _ = v78
                                    default:
                                        switch ((v57 + (v4 + "")) + (("\n\t" + v57) + "a")) {
                                        case "hello":
                                            v79 := GenS4{F0: 810, F1: 78, F2: 42}
                                            _ = v79
                                        default:
                                            if true {
                                                var v80 string = ("a" + v4)
                                                var v81 int64 = (v1 % ((1) | 1))
                                                _ = v80
                                                _ = v81
                                            } else {
                                                fmt.Println(v62, v56, v4)
                                                v82 := GenS8{F0: -1}
                                                _ = len(v3)
                                                _ = v82
                                            }
                                        }
                                    }
                                case "":
                                    v52 = v52
                                default:
                                    for fi83 := 0; fi83 < 1; fi83++ {
                                        if (18446744073709551615 == 4294967295) {
                                            v84 := GenS9{F0: 411, F1: (v53 / ((genFn1()) | 1)), F2: "hello"}
                                            switch "\n\t" {
                                            case "\n\t":
                                                v85 := []int{127}
                                                _ = v85
                                            case "foo bar":
                                                v86 := []int{-834, -1, (fi83 - -2147483648), 32767}
                                                _ = v86
                                            case "":
                                                var v87 int = fi83
                                                _ = v87
                                            }
                                            if ((115 > 2147483647) || true) {
                                                v57 = v4
                                            }
                                            _ = v84
                                        }
                                        fmt.Println(v62, v52, v4)
                                        v52 = "a"
                                        _ = fi83
                                    }
                                }
                            default:
                                _ = len(v5)
                            }
                            v88 := []int{988, 127, 760}
                            v56 = v56
                            _ = v88
                        }
                        _ = v62
                    }
                    v89 := []int{9223372036854775807}
                    _ = v57
                    _ = v89
                } else {
                    for fi90 := 0; fi90 < 17; fi90++ {
                        v52 = "a"
                        v91 := []int{605, fi90, fi90, 2147483647}
                        _ = fi90
                        _ = v91
                    }
                    v56 -= v56
                }
                _ = v56
            }
            fmt.Println(v1, v4)
            var v92 int8 = 1
            _ = v92
        }
    }
    var v93 int64 = (v1 | int64(v53))
    switch 4294967295 {
    case -2147483648:
        var v94 uint8 = 255
        _ = v94
    case 4294967295:
        var v95 uint8 = 127
        _ = v95
    default:
        var v96 int = -993
        _ = v96
    }
    if true {
        switch "a" {
        case "a":
            for fi97 := 0; fi97 < 17; fi97++ {
                for fi98 := 0; fi98 < 9; fi98++ {
                    v53 -= v53
                    _ = fi98
                }
                _ = fi97
            }
        case "hello":
            v93 = v93
        }
    }
    fmt.Println(v53, v4, v4)
    v53 += v53
    var v99 uint = 18446744073709551615
    v52 = v4
    var v100 int = -67
    v52 = (v4 + v4)
    v93 = (v1 % ((v1) | 1))
    switch (v100 ^ (v100 / ((v100) | 1))) {
    case 4294967295:
        switch 32767 {
        case 127:
            v101 := GenS2{F0: v52, F1: 0}
            _ = v101
        case -606:
            var v102 int = v100
            _ = v102
        case 9223372036854775807:
            var v103 int = (v100 * (-(v100)))
            _ = v103
        }
    case 9223372036854775807:
        v93 = v1
    case -1:
        fmt.Println(v99)
    }
    v93 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v93:", v93)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v93
    _ = v99
}
