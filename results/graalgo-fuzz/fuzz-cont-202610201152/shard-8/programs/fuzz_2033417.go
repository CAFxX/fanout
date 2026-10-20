package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 int16 }
func genFn0(p0 int8, p1 int32, p2 int8) string {
    fmt.Println(p2)
    fmt.Println(p0, p2, p0)
    fmt.Println(p1, p0)
    for fi4 := 0; fi4 < 5; fi4++ {
        for fi5 := 0; fi5 < 14; fi5++ {
            var v6 int8 = 100
            _ = fi5
            _ = v6
        }
        p1 = -1
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint16 = 32767
            var v9 float64 = -2.25
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi4
    }
    p2 -= p0
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}
func genFn1(p0 int16, p1 int8) int {
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint16 = 65535
        var v12 uint16 = (^(v11))
        if ((0.1 + float32(1.5)) >= (1e10 + (0.1 / (0.0 + 0.1)))) {
            var v13 uint64 = uint64(v12)
            _ = v13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    p0 = -32768
    if (1 == 1) {
        v14 := []int{-2147483648, 32768, -696}
        fmt.Println(p0, p1)
        _ = v14
    }
    var v15 int32 = 1
    switch -556 {
    case 9223372036854775807:
        v16 := []int{127, 0, 0, 128}
        _ = v16
    case 4294967295:
        if false {
            var v17 uint32 = 2147483647
            _ = v17
        }
    case -1:
        if (1e10 > 3.14159) {
            var v18 int64 = 1
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = v15
    return 1
}

func main() {
    var v1 uint = 42
    var v2 int8 = -76
    var v3 float64 = ((-2.25 * 0.1) * 0.0)
    var v19 int64 = 0
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 int32 = -1
        fmt.Println(v1, v21)
        _ = fi20
        _ = v21
    }
    var v22 int32 = 1
    v23 := []int{4294967295, 1, 127}
    fmt.Println(v19)
    switch (("a" + "foo bar") + (("a" + "\n\t") + "\n\t")) {
    case "hello":
        if (42 <= 2147483648) {
            v3 -= v3
            v1 = (v1 | 1)
            v1 = (v1 % ((v1) | 1))
        } else {
            if false {
                v24 := GenS1{F0: (v1 ^ v1), F1: 127}
                _ = v24
            }
            for fi25 := 0; fi25 < 18; fi25++ {
                switch "" {
                case "\n\t":
                    var v26 int8 = (v2 & int8(v22))
                    _ = v26
                case "foo bar":
                    var v27 string = "hello"
                    _ = v27
                case "":
                    _ = len(v23)
                default:
                    if false {
                        var v28 float32 = -2.25
                        v29 := GenS1{F0: 304, F1: int8(v1)}
                        _ = v28
                        _ = v29
                    }
                }
                var v30 int32 = 124
                _ = fi25
                _ = v30
            }
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
        }
    case "a":
        for fi31 := 0; fi31 < 18; fi31++ {
            v32 := GenS2{F0: "", F1: 42}
            v22 = v22
            _ = fi31
            _ = v32
        }
    case "":
        v22 = 146
    default:
        var v33 int16 = -32768
        _ = v33
    }
    if true {
        v34 := GenS2{F0: "a", F1: 2}
        for fi35 := 0; fi35 < 15; fi35++ {
            v1 += (v1 - 65535)
            v36 := GenS1{F0: v1, F1: (v2 / (((^(v2))) | 1))}
            _ = fi35
            _ = v36
        }
        v37 := GenS1{F0: 293, F1: v2}
        _ = v34
        _ = v37
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        v39 := GenS2{F0: "a", F1: 468}
        fmt.Println(fi38, v19)
        v23 = append(v23, -540)
        _ = fi38
        _ = v39
    }
    var v40 int = 0
    if true {
        v41 := GenS3{F0: 98}
        _ = v41
    } else {
        v42 := GenS1{F0: (^(v1)), F1: v2}
        v23 = append(v23, v40)
        _ = v42
    }
    if (!(!true)) {
        var v43 string = ""
        if true {
            v22 = -765
            for fi44 := 0; fi44 < 14; fi44++ {
                if ((v1 ^ v1) < v1) {
                    for fi45 := 0; fi45 < 15; fi45++ {
                        switch (fi44 / ((fi44) | 1)) {
                        case 42:
                            fmt.Println(v43, fi44)
                        case 1:
                            fmt.Println(v2, v3, v43)
                        case 717:
                            var v46 int64 = (v19 << (((v1 ^ 2)) & 63))
                            _ = v46
                        default:
                            if true {
                                if ((true || (true && (0 == 0))) || (626 != 4294967295)) {
                                    var v47 int16 = -32768
                                    v40 -= v40
                                    _ = v47
                                } else {
                                    var v48 float32 = 1.5
                                    _ = v48
                                }
                            } else {
                                if false {
                                    v3 += v3
                                    var v49 int16 = 526
                                    var v50 uint32 = 1
                                    _ = v49
                                    _ = v50
                                }
                                var v51 float64 = v3
                                _ = v51
                            }
                        }
                        v1 = (v1 ^ (v1 >> (((v1 ^ v1)) & 63)))
                        _ = fi45
                    }
                } else {
                    var v52 string = v43
                    for fi53 := 0; fi53 < 20; fi53++ {
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                        v3 = (v3 * float64((float32((1.5 * float64((1e10 / (1.5 / 1.0))))) + 1.5)))
                        v40 = -653
                        _ = fi53
                    }
                    _ = v52
                }
                v43 = v43
                _ = fi44
            }
            v54 := GenS2{F0: v43, F1: 255}
            _ = v54
        }
        if false {
            v22 += (v22 * 0)
            v55 := GenS1{F0: (v1 ^ 0), F1: v2}
            for fi56 := 0; fi56 < 12; fi56++ {
                for fi57 := 0; fi57 < 15; fi57++ {
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                    v58 := GenS2{F0: (("hello" + "hello") + v43), F1: 1}
                    _ = fi57
                    _ = v58
                }
                if false {
                    switch fi56 {
                    case 4294967295:
                        var v59 int8 = (v2 + v2)
                        _ = v59
                    case 1:
                        if (false || false) {
                            var v60 uint16 = uint16(v1)
                            _ = v60
                        }
                    case 344:
                        v2 -= (v2 / ((v2) | 1))
                    }
                }
                v23 = append(v23, genFn1(-32768, (v2 ^ (v2 | int8(v1)))))
                _ = fi56
            }
            _ = v55
        }
        _ = v43
    } else {
        for fi61 := 0; fi61 < 6; fi61++ {
            v2 = -114
            for fi62 := 0; fi62 < 7; fi62++ {
                if ((v3 <= v3) && true) {
                    v19 = v19
                    if (("\n\t" + ("foo bar" + ("" + "foo bar"))) < "hello") {
                        v40 -= (v40 * fi61)
                    } else {
                        var v63 float32 = (3.14159 / 1e10)
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                        _ = v63
                    }
                    var v64 int16 = -310
                    var v65 float32 = 3.14159
                    _ = v64
                    _ = v65
                }
                v66 := GenS3{F0: 2}
                _ = fi62
                _ = v66
            }
            if true {
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                var v67 bool = true
                switch ("a" + ("foo bar" + "foo bar")) {
                case "foo bar":
                    for fi68 := 0; fi68 < 12; fi68++ {
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                        var v69 int8 = (v2 << (((v1 >> ((v1) & 63))) & 63))
                        v2 -= 1
                        _ = fi68
                        _ = v69
                    }
                default:
                    if v67 {
                        v70 := GenS2{F0: "foo bar", F1: 0}
                        var v71 string = ""
                        for fi72 := 0; fi72 < 8; fi72++ {
                            var v73 float32 = 0.1
                            _ = fi72
                            _ = v73
                        }
                        _ = v70
                        _ = v71
                    } else {
                        v74 := GenS3{F0: 283}
                        v75 := GenS1{F0: (v1 << ((v1) & 63)), F1: v2}
                        for fi76 := 0; fi76 < 10; fi76++ {
                            if (((-2.25 / -2.25) + 1.5) <= 3.14159) {
                                var v77 uint64 = 2
                                _ = v77
                            }
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                            v22 += (v22 | v22)
                            _ = fi76
                        }
                        _ = v74
                        _ = v75
                    }
                }
                _ = v67
            }
            _ = fi61
        }
    }
    if len(v23) >= 2 {
        v23 = v23[1:len(v23)]
    }
    var v78 uint32 = 65535
    switch genFn1(1000, v2) {
    case 32768:
        v79 := GenS4{F0: (v2 / ((v2) | 1))}
        _ = v79
    case 127:
        for fi80 := 0; fi80 < 10; fi80++ {
            switch "\n\t" {
            case "\n\t":
                var v81 float64 = v3
                _ = v81
            case "foo bar":
                v82 := GenS1{F0: v1, F1: 127}
                _ = v82
            case "a":
                v3 = v3
            }
            var v83 int8 = (v2 * v2)
            _ = fi80
            _ = v83
        }
    default:
        v2 += (v2 ^ v2)
    }
    switch "a" {
    case "":
        fmt.Println(v3, v2, v40)
    case "hello":
        if ((!(true && false)) && true) {
            v84 := GenS5{F0: 802, F1: v2}
            for fi85 := 0; fi85 < 19; fi85++ {
                v86 := GenS2{F0: ("" + ""), F1: 65535}
                _ = fi85
                _ = v86
            }
            v78 = v78
            _ = v84
        }
    case "\n\t":
        for fi87 := 0; fi87 < 8; fi87++ {
            for fi88 := 0; fi88 < 5; fi88++ {
                v19 = (v19 % ((42) | 1))
                _ = fi88
            }
            v1 -= 4294967295
            switch ("foo bar" + "a") {
            case "":
                v23 = append(v23, v40)
            default:
                v89 := GenS6{F0: 42}
                _ = v89
            }
            _ = fi87
        }
    default:
        var v90 uint64 = 0
        _ = v90
    }
    var v91 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v78:", v78)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v40
    _ = v78
    _ = v91
}
