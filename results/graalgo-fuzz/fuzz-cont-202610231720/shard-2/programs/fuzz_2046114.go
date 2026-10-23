package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint32; F2 string }
type GenS3 struct { F0 int32; F1 int16; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 bool; F1 uint32; F2 uint }
type GenS9 struct { F0 bool; F1 string; F2 int32 }
func genFn0(p0 string, p1 uint) uint64 {
    fmt.Println(p1, p0)
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 float32 = (0.1 - (float32(3.14159) + (0.0 - 1.5)))
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 int8, p1 int8) uint64 {
    p0 = 42
    var v9 int16 = -32768
    var v10 bool = false
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 bool = false
    var v2 float64 = 1e10
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    var v5 uint16 = 1
    var v6 float64 = (1e10 - v2)
    fmt.Println(v2, v5)
    fmt.Println(v2)
    fmt.Println(v2, v6)
    switch ("a" + "\n\t") {
    case "":
        if v1 {
            if false {
                for fi11 := 0; fi11 < 15; fi11++ {
                    for fi12 := 0; fi12 < 3; fi12++ {
                        fmt.Println(v1)
                        _ = fi12
                    }
                    switch (fi11 - fi11) {
                    case 32768:
                        v13 := GenS1{F0: 690, F1: 42, F2: 255}
                        _ = v13
                    case -361:
                        if false {
                            v6 += v2
                            if (1 >= -128) {
                                v14 := GenS1{F0: 508, F1: 0, F2: 127}
                                _ = v14
                            }
                        } else {
                            var v15 int32 = 0
                            var v16 int = (fi11 * fi11)
                            if ((^(v16)) <= v16) {
                                v17 := []int{(fi11 % (((-(v16))) | 1)), fi11}
                                for fi18 := 0; fi18 < 4; fi18++ {
                                    switch (fi18 + (fi18 / (((fi11 ^ (v16 / (((v16 + 2147483647)) | 1)))) | 1))) {
                                    case 32767:
                                        switch "a" {
                                        case "foo bar":
                                            v19 := GenS2{F0: v15, F1: 22, F2: "a"}
                                            _ = v19
                                        }
                                    case 4294967295:
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                    default:
                                        v5 -= v5
                                    }
                                    _ = fi18
                                }
                                _ = v17
                            }
                            _ = v15
                            _ = v16
                        }
                    case -2147483648:
                        if v1 {
                            v2 = v6
                        } else {
                            v20 := GenS2{F0: 42, F1: 370, F2: "hello"}
                            _ = v20
                        }
                    }
                    var v21 float64 = v2
                    _ = fi11
                    _ = v21
                }
            }
        }
    case "\n\t":
        v22 := GenS1{F0: 0, F1: 1, F2: 255}
        _ = v22
    case "a":
        v6 -= float64((1e10 + (((1.5 - 1e10) * -2.25) * 3.14159)))
    }
    if (2147483648 == 2147483648) {
        for fi23 := 0; fi23 < 18; fi23++ {
            v24 := v3[fi23]
            if (!v1) {
                fmt.Println(v2, v5, v6)
            } else {
                v25 := GenS1{F0: 1, F1: -128, F2: 42}
                _ = v25
            }
            _ = fi23
            _ = v24
        }
        v6 -= ((float64(-2.25) + 0.1) - (0.0 - v6))
        var v26 uint8 = 254
        _ = v26
    } else {
        if v1 {
            v27 := []int{443, 128, 128}
            if (false && (96 == 42)) {
                switch ("hello" + ("" + ("\n\t" + "a"))) {
                case "hello":
                    v2 -= 1e10
                case "\n\t":
                    v3[32767] = 9223372036854775807
                default:
                    v4["\n\t"] = 407
                }
                var v28 uint = 42
                _ = v28
            } else {
                v5 -= 2
            }
            _ = v27
        } else {
            var v29 int32 = -63
            for fi30 := 0; fi30 < 13; fi30++ {
                if (0 > 257) {
                    v2 -= (-2.25 / -2.25)
                }
                for fi31 := 0; fi31 < 20; fi31++ {
                    v32 := GenS3{F0: (v29 - v29), F1: -452, F2: 543}
                    if v1 {
                        var v33 int16 = 1
                        _ = v33
                    }
                    fmt.Println(v29)
                    v34 := GenS4{F0: 42, F1: (v5 - (v5 - (v5 >> ((0) & 63))))}
                    var v35 bool = v1
                    _ = fi31
                    _ = v32
                    _ = v34
                    _ = v35
                }
                _ = fi30
            }
            fmt.Println(v6)
            _ = v29
        }
    }
    var v36 int16 = 32767
    v37 := GenS5{F0: 2, F1: 1000, F2: 42}
    if (!(!v1)) {
        for fi38 := 0; fi38 < 1; fi38++ {
            var v39 uint32 = 2147483647
            v1 = v1
            v6 -= 0.1
            _ = fi38
            _ = v39
        }
    }
    v1 = (65535 < 2147483648)
    v1 = false
    if true {
        if v1 {
            for fi40 := 0; fi40 < 15; fi40++ {
                for fi41 := 0; fi41 < 6; fi41++ {
                    var v42 uint16 = (v5 + (^(v5)))
                    v43 := []int{fi40, (fi41 - fi40), 127, fi40}
                    fmt.Println(v2)
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                v5 = 454
                var v44 int64 = 1
                _ = fi40
                _ = v44
            }
            if v1 {
                var v45 uint16 = 1
                _ = v45
            } else {
                switch "foo bar" {
                case "a":
                    var v46 int = -2147483648
                    _ = v46
                case "":
                    v47 := []int{1, 32768, 127, 0}
                    _ = v47
                case "hello":
                    switch "a" {
                    case "a":
                        if v1 {
                            v48 := GenS6{F0: (^(v5)), F1: 58}
                            _ = v48
                        }
                    case "\n\t":
                        var v49 int8 = 127
                        _ = v49
                    case "hello":
                        v36 += 42
                    default:
                        var v50 bool = ((!v1) && ("" > "foo bar"))
                        _ = v50
                    }
                default:
                    switch "" {
                    case "hello":
                        var v51 uint8 = 23
                        _ = v51
                    case "\n\t":
                        fmt.Println(v6, v1, v5)
                    case "foo bar":
                        var v52 string = "foo bar"
                        _ = v52
                    default:
                        v53 := GenS7{F0: -866, F1: 1}
                        _ = v53
                    }
                }
            }
            var v54 uint8 = 128
            _ = v54
        }
        v55 := []int{32768, 142}
        fmt.Println(v5)
        _ = v55
    } else {
        v1 = true
        v36 = v36
        for fi56 := 0; fi56 < 17; fi56++ {
            switch fi56 {
            case 659:
                for fi57 := 0; fi57 < 12; fi57++ {
                    var v58 uint16 = 255
                    v59 := v3[fi57]
                    _ = fi57
                    _ = v58
                    _ = v59
                }
            case 201:
                if false {
                    switch "hello" {
                    case "":
                        if v1 {
                            var v60 uint16 = v5
                            v60 = (v5 + (v5 - v60))
                            _ = v60
                        }
                    case "hello":
                        var v61 int32 = 0
                        _ = v61
                    }
                } else {
                    if (!(((("\n\t" + "") + ("a" + "")) + "a") >= "\n\t")) {
                        var v62 float64 = (0.1 - float64(0.0))
                        _ = v62
                    } else {
                        v6 += v6
                    }
                    switch fi56 {
                    case 936:
                        switch 9223372036854775807 {
                        case 1:
                            for fi63 := 0; fi63 < 3; fi63++ {
                                v64 := v4["a"]
                                switch ("\n\t" + "\n\t") {
                                case "a":
                                    var v65 int8 = 1
                                    _ = v65
                                case "\n\t":
                                    var v66 int8 = -122
                                    _ = v66
                                case "hello":
                                    if v1 {
                                        var v67 int32 = -1
                                        fmt.Println(fi63, v1, v6)
                                        var v68 float64 = 1.5
                                        _ = v67
                                        _ = v68
                                    } else {
                                        if (v2 >= -2.25) {
                                            v2 -= v2
                                            var v69 float32 = float32((1.5 / (float64(3.14159) * 3.14159)))
                                            if v1 {
                                                if (int8(v5) != -92) {
                                                    v6 -= 3.14159
                                                    v6 = (1.5 * (0.0 * v6))
                                                    v69 -= v69
                                                } else {
                                                    if (!(v1 && (1 > 42))) {
                                                        switch "\n\t" {
                                                        case "":
                                                            var v70 uint = 255
                                                            _ = v70
                                                        default:
                                                            v5 += v5
                                                        }
                                                        v1 = v1
                                                        var v71 uint16 = (v5 / ((v5) | 1))
                                                        _ = v71
                                                    }
                                                    v36 -= v36
                                                    v5 -= v5
                                                    v72 := []int{(fi63 + -812), fi56, 331}
                                                    v72 = append(v72, fi56)
                                                    _ = v72
                                                }
                                                var v73 int = fi56
                                                _ = v73
                                            }
                                            _ = v69
                                        }
                                        v74 := GenS8{F0: v1, F1: 42, F2: 12}
                                        _ = v74
                                    }
                                default:
                                    _ = len(v3)
                                }
                                _ = fi63
                                _ = v64
                            }
                        case 198:
                            fmt.Println(fi56, v36)
                        case -27:
                            if (!(!(0.0 >= float32(float64(3.14159))))) {
                                switch (fi56 - (fi56 << ((42) & 63))) {
                                case 32768:
                                    var v75 uint64 = 346
                                    _ = v75
                                case -1:
                                    var v76 uint16 = (v5 ^ (v5 % ((1) | 1)))
                                    _ = v76
                                default:
                                    if (0 > 57) {
                                        for fi77 := 0; fi77 < 11; fi77++ {
                                            v2 -= v2
                                            var v78 uint = 42
                                            var v79 int = (fi77 * (^(fi77)))
                                            var v80 float64 = 0.0
                                            v2 -= (v2 + v80)
                                            var v81 uint16 = 458
                                            var v82 uint = (^(v78))
                                            v6 += v80
                                            _ = fi77
                                            _ = v78
                                            _ = v79
                                            _ = v80
                                            _ = v81
                                            _ = v82
                                        }
                                        var v83 uint = 0
                                        var v84 uint8 = 1
                                        _ = v83
                                        _ = v84
                                    }
                                }
                                for fi85 := 0; fi85 < 12; fi85++ {
                                    v2 = v2
                                    _ = len(v3)
                                    _ = fi85
                                }
                                for fi86 := 0; fi86 < 14; fi86++ {
                                    switch "hello" {
                                    case "a":
                                        var v87 float32 = float32(0.1)
                                        _ = v87
                                    case "foo bar":
                                        if true {
                                            switch (fi86 ^ (-(fi86))) {
                                            case 1:
                                                if v1 {
                                                    var v88 int8 = 0
                                                    _ = v88
                                                }
                                            default:
                                                var v89 int = fi86
                                                _ = v89
                                            }
                                            v90 := GenS9{F0: (v5 != v5), F1: ("foo bar" + ""), F2: 0}
                                            var v91 uint16 = 515
                                            _ = v90
                                            _ = v91
                                        }
                                    case "hello":
                                        var v92 int64 = 746
                                        _ = v92
                                    }
                                    v36 = v36
                                    switch 32767 {
                                    case 390:
                                        v93 := GenS3{F0: -615, F1: (-(v36)), F2: genFn0(("hello" + ("a" + "foo bar")), 255)}
                                        _ = v93
                                    default:
                                        v1 = (false || (-2.25 <= -2.25))
                                    }
                                    _ = fi86
                                }
                            } else {
                                v6 += (v2 + v2)
                                var v94 int32 = 768
                                _ = v94
                            }
                        }
                    case 32768:
                        v95 := GenS8{F0: v1, F1: 42, F2: 42}
                        _ = v95
                    default:
                        v96 := GenS2{F0: 254, F1: 42, F2: ""}
                        _ = v96
                    }
                }
            default:
                var v97 int = (^(fi56))
                _ = v97
            }
            v5 = 258
            _ = fi56
        }
    }
    if ((!("hello" > "\n\t")) || v1) {
        v98 := []int{-2147483648}
        switch "\n\t" {
        case "":
            v99 := GenS6{F0: v5, F1: 171}
            _ = v99
        }
        v5 = (v5 | 751)
        _ = v98
    } else {
        var v100 bool = true
        _ = v100
    }
    for fi101 := 0; fi101 < 9; fi101++ {
        if v1 {
            var v102 uint64 = 0
            _ = v102
        }
        _ = fi101
    }
    for fi103 := 0; fi103 < 6; fi103++ {
        var v104 int16 = v36
        fmt.Println(v104, fi103, fi103)
        for fi105 := 0; fi105 < 3; fi105++ {
            v104 += v36
            v106 := []int{fi105}
            _ = fi105
            _ = v106
        }
        _ = fi103
        _ = v104
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
