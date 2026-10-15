package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int8; F1 string; F2 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint }
func genFn0(p0 int16, p1 bool) int64 {
    for fi4 := 0; fi4 < 9; fi4++ {
        if (-94 >= p0) {
            v5 := GenS1{F0: 42}
            v6 := GenS1{F0: -495}
            p0 = p0
            _ = v5
            _ = v6
        }
        if (p1 && (2147483647 < 1)) {
            for fi7 := 0; fi7 < 1; fi7++ {
                switch "a" {
                case "\n\t":
                    v8 := []int{}
                    _ = v8
                case "":
                    fmt.Println(fi7, p0, fi4)
                case "hello":
                    var v9 int16 = (-(p0))
                    _ = v9
                }
                p0 += (p0 + (-(p0)))
                p0 -= 1000
                _ = fi7
            }
            v10 := GenS2{F0: 12, F1: "\n\t"}
            v11 := []int{fi4, 2147483647, fi4, (fi4 / ((fi4) | 1))}
            _ = v10
            _ = v11
        } else {
            var v12 uint16 = 1
            _ = v12
        }
        _ = fi4
    }
    var v13 float64 = 0.0
    for fi14 := 0; fi14 < 11; fi14++ {
        if p1 {
            v15 := []int{fi14}
            fmt.Println(p1)
            _ = v15
        } else {
            var v16 float32 = 1e10
            _ = v16
        }
        fmt.Println(p1)
        p1 = false
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = v13
    return 66
}
func genFn1(p0 int, p1 uint16, p2 int) uint64 {
    v17 := GenS1{F0: -2147483648}
    v18 := []int{(p0 + 9223372036854775807), 9223372036854775807, p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return 42
}

func main() {
    var v1 uint = 466
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    switch 127 {
    case 42:
        v19 := GenS3{F0: 122, F1: "foo bar", F2: -9223372036854775808}
        _ = v19
    case -210:
        for fi20 := 0; fi20 < 16; fi20++ {
            switch "\n\t" {
            case "\n\t":
                if false {
                    v1 -= 0
                    v21 := v2[(-(fi20))]
                    _ = v21
                }
            default:
                var v22 int = (fi20 | -2147483648)
                _ = v22
            }
            fmt.Println(fi20, v1)
            for fi23 := 0; fi23 < 8; fi23++ {
                var v24 int = fi20
                switch fi23 {
                case -1:
                    v25 := GenS2{F0: -6, F1: "a"}
                    _ = v25
                case 9223372036854775807:
                    var v26 int16 = 1
                    fmt.Println(v26, fi20, v1)
                    v26 = v26
                    var v27 uint32 = 2147483647
                    var v28 int64 = 9223372036854775807
                    var v29 string = ("a" + "\n\t")
                    v30 := GenS1{F0: 0}
                    _ = v26
                    _ = v27
                    _ = v28
                    _ = v29
                    _ = v30
                default:
                    var v31 uint32 = 598
                    _ = v31
                }
                v24 += fi23
                _ = fi23
                _ = v24
            }
            _ = fi20
        }
    case 32768:
        v32 := GenS1{F0: 488}
        _ = v32
    default:
        v1 = (v1 >> ((v1) & 63))
    }
    v33 := []int{543}
    _ = len(v3)
    if (float64((0.0 * ((0.1 + 1e10) * float32((1e10 - (0.0 - 1.5)))))) >= 1.5) {
        v34 := GenS3{F0: 42, F1: "a", F2: 0}
        _ = v34
    } else {
        fmt.Println(v1)
        v1 = 2
        for fi35 := 0; fi35 < 1; fi35++ {
            v36 := GenS1{F0: 2147483647}
            var v37 int = 128
            _ = fi35
            _ = v36
            _ = v37
        }
    }
    var v38 uint64 = 958
    for fi39 := 0; fi39 < 18; fi39++ {
        var v40 uint16 = 598
        for fi41 := 0; fi41 < 13; fi41++ {
            switch "hello" {
            case "hello":
                if (0 == (v38 ^ v38)) {
                    v40 += 65535
                    if (false && true) {
                        v33 = append(v33, (fi41 % (((fi41 ^ fi39)) | 1)))
                    }
                    v2[fi41] = (fi39 >> ((v1) & 63))
                } else {
                    v40 -= v40
                }
            case "":
                if len(v33) > 0 {
                    v33[0] = fi41
                }
            case "foo bar":
                if (677 != 471) {
                    var v42 uint = 18446744073709551615
                    if (false || (false || (648 >= genFn1((^(fi41)), (v40 * v40), (fi39 ^ fi39))))) {
                        var v43 uint16 = (v40 * (v40 + (v40 >> (((^(v1))) & 63))))
                        _ = v43
                    } else {
                        var v44 int16 = 0
                        _ = len(v2)
                        if ((fi41 * 0) == fi39) {
                            if (false && true) {
                                var v45 uint64 = (^(v38))
                                _ = v45
                            } else {
                                v1 = v42
                                v46 := GenS1{F0: 2147483647}
                                switch fi41 {
                                case 42:
                                    fmt.Println(v1)
                                case 128:
                                    var v47 uint64 = (v38 & (^(v38)))
                                    _ = v47
                                }
                                _ = v46
                            }
                        } else {
                            v38 = 591
                            if true {
                                v42 -= (v1 / (((v1 * v1)) | 1))
                            } else {
                                var v48 uint32 = 2
                                _ = v48
                            }
                        }
                        _ = v44
                    }
                    _ = v42
                }
            }
            if (true && (false && true)) {
                var v49 float32 = float32(0.1)
                v50 := GenS2{F0: 100, F1: ("a" + ("" + ""))}
                var v51 int16 = 915
                _ = v49
                _ = v50
                _ = v51
            }
            _ = fi41
        }
        var v52 uint16 = v40
        _ = fi39
        _ = v40
        _ = v52
    }
    v1 = v1
    v38 = v38
    if ((2 != 65535) || (true && true)) {
        if (uint32(v38) != 2147483647) {
            for fi53 := 0; fi53 < 5; fi53++ {
                v3[("\n\t" + "foo bar")] = -765
                var v54 string = "foo bar"
                _ = fi53
                _ = v54
            }
        }
        v55 := GenS4{F0: -1}
        switch 128 {
        case -2147483648:
            v1 = 213
        }
        _ = v55
    } else {
        v56 := GenS5{F0: (v1 & v1)}
        for fi57 := 0; fi57 < 14; fi57++ {
            if len(v33) > 0 {
                v33[0] = fi57
            }
            v58 := GenS5{F0: 2}
            _ = fi57
            _ = v58
        }
        fmt.Println(v1)
        _ = v56
    }
    if len(v33) >= 2 {
        v33 = v33[1:len(v33)]
    }
    fmt.Println(v38, v38)
    switch -686 {
    case 235:
        fmt.Println(v1, v38)
    case 632:
        v38 += genFn1(177, 1, -917)
    case -2147483648:
        for fi59 := 0; fi59 < 9; fi59++ {
            v60 := GenS3{F0: -1, F1: "foo bar", F2: fi59}
            _ = fi59
            _ = v60
        }
    default:
        v38 -= 0
    }
    for fi61 := 0; fi61 < 4; fi61++ {
        _ = len(v3)
        v1 = 673
        _ = fi61
    }
    var v62 int32 = -2147483648
    var v63 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v38:", v38)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v62
    _ = v63
}
