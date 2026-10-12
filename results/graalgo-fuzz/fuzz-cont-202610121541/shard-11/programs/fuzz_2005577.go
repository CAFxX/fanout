package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint32; F1 int16; F2 int }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 int8; F1 uint8; F2 int8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int64; F1 int8 }
type GenS10 struct { F0 string; F1 uint; F2 uint8 }
type GenS11 struct { F0 int32; F1 int32 }
func genFn0(p0 int64, p1 int16) string {
    v6 := GenS1{F0: 4294967295, F1: 100, F2: 255}
    fmt.Println(p0, p1)
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint8 = 1
        switch (-(fi7)) {
        case -2147483648:
            var v9 int8 = 1
            _ = v9
        case 4294967295:
            for fi10 := 0; fi10 < 2; fi10++ {
                if (!false) {
                    var v11 float32 = 1e10
                    fmt.Println(fi7, v11, v8)
                    var v12 bool = true
                    _ = v11
                    _ = v12
                }
                switch "a" {
                case "hello":
                    var v13 float64 = 0.1
                    _ = v13
                case "":
                    if ((0.0 - 0.0) < ((0.1 * -2.25) + 3.14159)) {
                        var v14 uint64 = 368
                        var v15 uint32 = 2147483647
                        var v16 int16 = p1
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        var v17 uint = 255
                        _ = v17
                    }
                case "a":
                    switch ("\n\t" + "\n\t") {
                    case "\n\t":
                        var v18 float32 = -2.25
                        _ = v18
                    case "hello":
                        switch fi10 {
                        case -127:
                            v19 := GenS2{F0: false, F1: (p1 + p1)}
                            _ = v19
                        default:
                            v20 := []int{fi7}
                            _ = v20
                        }
                    case "a":
                        v21 := []int{(^(fi7)), fi7}
                        _ = v21
                    }
                default:
                    switch 32768 {
                    case 127:
                        p0 -= (p0 << ((1) & 63))
                        var v22 int64 = (p0 | p0)
                        _ = v22
                    case -2147483648:
                        fmt.Println(v8)
                    case 254:
                        p0 += (^(p0))
                    default:
                        switch (-(fi10)) {
                        case -2147483648:
                            v23 := GenS1{F0: 2, F1: -1, F2: 42}
                            _ = v23
                        default:
                            switch ((("a" + "hello") + "") + (((("foo bar" + ("foo bar" + "hello")) + ((("hello" + "foo bar") + "") + "\n\t")) + ("a" + "\n\t")) + "\n\t")) {
                            case "\n\t":
                                var v24 int64 = p0
                                _ = v24
                            }
                        }
                    }
                }
                fmt.Println(p1, fi10)
                var v25 int16 = p1
                var v26 int32 = 0
                if (!((((!((-(fi7)) == fi7)) && (!false)) || false) || ((!(255 >= 1)) && (2 != (v8 | 42))))) {
                    var v27 float64 = ((0.1 / (3.14159 + 1e10)) / 1e10)
                    var v28 uint16 = 2
                    v29 := GenS1{F0: 0, F1: 100, F2: 902}
                    var v30 string = "\n\t"
                    var v31 int = -2147483648
                    v27 = (v27 - ((((0.0 / (v27 * v27)) - v27) - 1e10) * float64(3.14159)))
                    _ = v27
                    _ = v28
                    _ = v29
                    _ = v30
                    _ = v31
                }
                if (!((v26 ^ int32(v25)) < v26)) {
                    v32 := GenS3{F0: 632, F1: int16(v26), F2: 1}
                    switch 2147483647 {
                    case -759:
                        var v33 bool = (((v25 * int16(v26)) <= (-(v25))) || (3.14159 == ((-2.25 * (0.1 * 1e10)) + 1.5)))
                        _ = v33
                    case 9223372036854775807:
                        v25 += 697
                        v34 := GenS4{F0: 42, F1: 0}
                        _ = v34
                    case 32767:
                        v8 = (^(v8))
                    }
                    if false {
                        if (true && ((^(v8)) != v8)) {
                            var v35 uint8 = v8
                            switch (fi10 >> ((42) & 63)) {
                            case 2147483647:
                                if (!((((fi7 << ((255) & 63)) >= fi10) && (((2 <= 0) && false) || (v26 != (^(v26))))) || ((!(127 == 103)) && false))) {
                                    var v36 int32 = v26
                                    _ = v36
                                }
                            case 32768:
                                var v37 bool = false
                                _ = v37
                            case -1:
                                v38 := []int{(fi10 << ((4294967295) & 63)), 1}
                                _ = v38
                            }
                            switch "a" {
                            case "\n\t":
                                v26 += v26
                            case "foo bar":
                                if true {
                                    v35 = (v35 * v8)
                                    var v39 uint64 = 42
                                    var v40 string = ("\n\t" + "a")
                                    _ = v39
                                    _ = v40
                                } else {
                                    v41 := GenS5{F0: true, F1: 534}
                                    var v42 int64 = (p0 | (p0 & p0))
                                    if true {
                                        v42 = p0
                                        v8 = (v8 + v35)
                                        v26 += (v26 / ((v26) | 1))
                                    } else {
                                        var v43 uint16 = 1
                                        v44 := GenS6{F0: 100, F1: 39, F2: 27}
                                        switch (("a" + ("" + "a")) + "hello") {
                                        case "hello":
                                            v45 := GenS5{F0: true, F1: (p1 >> ((2) & 63))}
                                            _ = v45
                                        case "\n\t":
                                            var v46 int = (fi7 | (fi7 | fi10))
                                            _ = v46
                                        default:
                                            v26 = v26
                                            p0 += p0
                                            v47 := []int{(fi10 + (fi7 * fi10))}
                                            _ = v47
                                        }
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = v41
                                    _ = v42
                                }
                            default:
                                v48 := []int{}
                                _ = v48
                            }
                            _ = v35
                        }
                    }
                    _ = v32
                }
                _ = fi10
                _ = v25
                _ = v26
            }
        case 42:
            v49 := GenS5{F0: ((1 >= 1) && (!true)), F1: p1}
            _ = v49
        }
        if (32767 < 387) {
            var v50 int32 = -2147483648
            v51 := GenS7{F0: 42}
            _ = v50
            _ = v51
        }
        _ = fi7
        _ = v8
    }
    if (false && true) {
        switch -2147483648 {
        case -224:
            fmt.Println(p0, p0)
        case 9223372036854775807:
            if (42 != 153) {
                p1 -= p1
                var v52 float64 = 1.5
                var v53 float64 = (float64(3.14159) + -2.25)
                _ = v52
                _ = v53
            }
        default:
            for fi54 := 0; fi54 < 12; fi54++ {
                if (!false) {
                    v55 := []int{(fi54 ^ (fi54 << ((2) & 63)))}
                    fmt.Println(p0, p0, fi54)
                    for fi56 := 0; fi56 < 9; fi56++ {
                        var v57 uint32 = 2147483648
                        _ = fi56
                        _ = v57
                    }
                    _ = v55
                }
                if true {
                    var v58 uint32 = 2147483647
                    for fi59 := 0; fi59 < 2; fi59++ {
                        var v60 string = ""
                        p1 -= -32768
                        _ = fi59
                        _ = v60
                    }
                    _ = v58
                }
                _ = fi54
            }
        }
        v61 := []int{127, 42}
        for fi62 := 0; fi62 < 6; fi62++ {
            fmt.Println(p1, p1)
            if true {
                v63 := GenS4{F0: 4294967295, F1: 985}
                _ = v63
            }
            var v64 uint64 = 18446744073709551615
            _ = fi62
            _ = v64
        }
        _ = v61
    }
    v65 := GenS8{F0: 42}
    _ = p0
    _ = p1
    _ = v6
    _ = v65
    return "foo bar"
}

func main() {
    var v1 int8 = 100
    v2 := []string{"hello"}
    v3 := make(map[uint32]string)
    var v4 int64 = 1
    var v5 int32 = -2147483648
    v66 := GenS4{F0: 65535, F1: 976}
    if ((512 < 42) || (!true)) {
        v2 = append(v2, "hello")
        var v67 uint32 = 4294967295
        var v68 int32 = (v5 >> ((1) & 63))
        _ = v67
        _ = v68
    } else {
        var v69 int64 = 952
        v4 += v69
        _ = v69
    }
    fmt.Println(v5)
    v1 -= v1
    fmt.Println(v5, v1)
    v4 -= (v4 / (((v4 % (((v4 * v4)) | 1))) | 1))
    for fi70 := 0; fi70 < 8; fi70++ {
        switch ("a" + "") {
        case "a":
            var v71 float64 = ((0.1 + 0.0) * (1.5 * ((1.5 + 0.1) / 3.14159)))
            _ = v71
        }
        if true {
            for fi72 := 0; fi72 < 14; fi72++ {
                var v73 float64 = 0.1
                v5 += (v5 | (v5 & int32(v4)))
                switch fi72 {
                case -9223372036854775808:
                    if (true || (false && (-32768 < 868))) {
                        if ((1.5 * 0.0) != 0.0) {
                            v74 := v3[733]
                            if (int32(v1) == (v5 * 2147483647)) {
                                var v75 uint16 = 850
                                var v76 int32 = -896
                                switch fi70 {
                                case 908:
                                    fmt.Println(fi72, v4, v5)
                                default:
                                    v77 := GenS1{F0: 65535, F1: (v1 & v1), F2: 0}
                                    if (false && true) {
                                        var v78 int16 = 26
                                        _ = len(v3)
                                        _ = v78
                                    }
                                    fmt.Println(v76, v4, v75)
                                    v79 := GenS9{F0: (v4 ^ (-(v4))), F1: v1}
                                    fmt.Println(v75, v1, v76)
                                    _ = v77
                                    _ = v79
                                }
                                _ = v75
                                _ = v76
                            } else {
                                var v80 uint8 = 117
                                v81 := GenS6{F0: 42, F1: v80, F2: (^(v1))}
                                fmt.Println(v5, v73, v80)
                                _ = v80
                                _ = v81
                            }
                            _ = v74
                        }
                        v73 += (v73 / (v73 / float64(1e10)))
                    } else {
                        v82 := GenS10{F0: "a", F1: 616, F2: 209}
                        if (!(454 > 0)) {
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                            v73 = v73
                            v1 -= v1
                            fmt.Println(v73, fi70, v1)
                        } else {
                            v83 := GenS11{F0: 2147483647, F1: (v5 & v5)}
                            _ = v83
                        }
                        _ = v82
                    }
                case -1:
                    fmt.Println(v4)
                }
                _ = fi72
                _ = v73
            }
        } else {
            v84 := GenS9{F0: (v4 & v4), F1: (-(v1))}
            var v85 int64 = (v4 - v4)
            _ = v84
            _ = v85
        }
        _ = fi70
    }
    var v86 uint64 = 42
    var v87 uint = 2
    v1 -= (v1 % (((-(v1))) | 1))
    fmt.Println(v86)
    v86 += 42
    fmt.Println(v4, v1)
    var v88 float32 = -2.25
    v1 += (v1 * (v1 + (v1 << (((v87 + (^(v87)))) & 63))))
    var v89 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v86
    _ = v87
    _ = v88
    _ = v89
}
