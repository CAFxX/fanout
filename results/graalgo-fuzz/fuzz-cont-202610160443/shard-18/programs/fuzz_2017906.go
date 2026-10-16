package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS9 struct { F0 string; F1 uint64 }
type GenS10 struct { F0 uint8; F1 string }
func genFn0(p0 int32, p1 int, p2 int64) uint64 {
    var v7 int32 = -2147483648
    var v8 uint16 = 473
    if false {
        v7 += (v7 & -1)
        for fi9 := 0; fi9 < 4; fi9++ {
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 uint8 = 2
                _ = fi10
                _ = v11
            }
            switch "" {
            case "a":
                var v12 float64 = 3.14159
                _ = v12
            case "foo bar":
                if (v8 == v8) {
                    v13 := GenS1{F0: 561, F1: 127, F2: -80}
                    p2 -= (p2 + 0)
                    for fi14 := 0; fi14 < 15; fi14++ {
                        if false {
                            var v15 int16 = -506
                            var v16 uint16 = 642
                            v17 := []int{fi14}
                            switch p1 {
                            case 32767:
                                fmt.Println(fi9, fi14, p1)
                            }
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        _ = fi14
                    }
                    _ = v13
                }
            case "\n\t":
                var v18 int16 = 42
                _ = v18
            default:
                p1 -= (fi9 / ((fi9) | 1))
            }
            _ = fi9
        }
    } else {
        switch "\n\t" {
        case "a":
            if ((p1 < (p1 / ((1) | 1))) || (float32(-2.25) > 1.5)) {
                v19 := GenS1{F0: 0, F1: 123, F2: 127}
                _ = v19
            } else {
                var v20 bool = ((v8 < v8) && ((v7 >= p0) && (-39 <= -128)))
                _ = v20
            }
        default:
            v8 += (v8 << ((255) & 63))
        }
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 3; fi22++ {
            switch "hello" {
            case "":
                var v23 int = fi21
                _ = v23
            }
            if false {
                var v24 int32 = (-(v7))
                v25 := GenS2{F0: (^(p0)), F1: fi21}
                _ = v24
                _ = v25
            } else {
                v26 := []int{-1, (-(p1)), (p1 << ((4294967295) & 63)), (-(fi22))}
                _ = v26
            }
            var v27 int16 = 59
            _ = fi22
            _ = v27
        }
        var v28 uint = 800
        for fi29 := 0; fi29 < 15; fi29++ {
            v30 := []int{(fi29 & (p1 ^ fi21)), fi29}
            _ = fi29
            _ = v30
        }
        _ = fi21
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}
func genFn1() int64 {
    v31 := GenS1{F0: 18446744073709551615, F1: -128, F2: 48}
    for fi32 := 0; fi32 < 1; fi32++ {
        v33 := GenS2{F0: -579, F1: (fi32 | fi32)}
        var v34 uint64 = 340
        var v35 float32 = 3.14159
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    _ = v31
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 521
    var v3 float64 = 1e10
    var v4 int64 = -976
    v5 := make(map[int]bool)
    var v6 bool = false
    v5[-2147483648] = (42 == 1)
    v5[32768] = ("" != "\n\t")
    v36 := GenS3{F0: (32767 >= 358)}
    v37 := []int{0, 32768, 447, 1}
    if false {
        if (-2147483648 == 32767) {
            if ((751 <= 223) || (1 >= -1)) {
                var v38 uint64 = 18446744073709551615
                _ = v38
            }
            v39 := GenS4{F0: 65535, F1: (v2 / (((v2 ^ 4294967295)) | 1))}
            var v40 bool = v6
            _ = v39
            _ = v40
        }
        for fi41 := 0; fi41 < 20; fi41++ {
            v42 := GenS5{F0: 255}
            v43 := GenS4{F0: 531, F1: v2}
            _ = fi41
            _ = v42
            _ = v43
        }
        delete(v1, (("foo bar" + ("a" + "a")) + "\n\t"))
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        switch (fi44 & (fi44 + (fi44 >> ((42) & 63)))) {
        case -1:
            if v6 {
                delete(v5, fi44)
                switch (^(fi44)) {
                case 0:
                    for fi45 := 0; fi45 < 11; fi45++ {
                        var v46 uint16 = 225
                        v37 = append(v37, -862)
                        _ = fi45
                        _ = v46
                    }
                case 1:
                    var v47 string = ""
                    _ = v47
                }
                delete(v1, "foo bar")
            } else {
                var v48 uint = 922
                _ = v48
            }
        case -2147483648:
            var v49 float64 = 3.14159
            _ = v49
        case 127:
            if ((v2 & (v2 | v2)) != 4294967295) {
                if v6 {
                    var v50 uint = 18446744073709551615
                    v51 := GenS4{F0: 255, F1: (v2 | v2)}
                    var v52 int = 9223372036854775807
                    _ = v50
                    _ = v51
                    _ = v52
                }
                v2 = v2
            } else {
                switch "" {
                case "a":
                    v6 = v6
                case "\n\t":
                    var v53 uint32 = v2
                    _ = v53
                case "":
                    if v6 {
                        var v54 string = ("\n\t" + "\n\t")
                        _ = v54
                    }
                }
                var v55 uint = 4294967295
                v56 := GenS1{F0: (v55 & (v55 ^ 1)), F1: 127, F2: 127}
                _ = v55
                _ = v56
            }
        default:
            v3 = 0.1
        }
        v57 := GenS6{F0: (^(v2)), F1: v4}
        _ = fi44
        _ = v57
    }
    var v58 int = 199
    v59 := GenS1{F0: 949, F1: 42, F2: 42}
    if false {
        delete(v5, (-(v58)))
        for fi60 := 0; fi60 < 12; fi60++ {
            v61 := GenS6{F0: (v2 << ((758) & 63)), F1: 42}
            _ = fi60
            _ = v61
        }
        var v62 uint32 = v2
        _ = v62
    } else {
        if (!v6) {
            fmt.Println(v58)
            v63 := v5[32768]
            v64 := GenS7{F0: (v2 * (^(v2)))}
            _ = v63
            _ = v64
        } else {
            if v6 {
                v65 := v1[("hello" + (("foo bar" + ("foo bar" + "a")) + (("" + ("a" + ("\n\t" + ("" + "hello")))) + ("foo bar" + "\n\t"))))]
                fmt.Println(v6, v65)
                switch (v58 % ((v58) | 1)) {
                case 540:
                    switch (v58 >> ((290) & 63)) {
                    case 0:
                        for fi66 := 0; fi66 < 14; fi66++ {
                            fmt.Println(v58, v3, v4)
                            _ = fi66
                        }
                    }
                case -9223372036854775808:
                    if v6 {
                        var v67 int = v58
                        for fi68 := 0; fi68 < 3; fi68++ {
                            if (true || ((!(540 > 31)) && ((-(v4)) != 9223372036854775807))) {
                                if (714 == 569) {
                                    v2 -= (v65 >> ((42) & 63))
                                    v69 := GenS5{F0: 940}
                                    _ = v69
                                } else {
                                    fmt.Println(v58, v67, v58)
                                    for fi70 := 0; fi70 < 8; fi70++ {
                                        switch "foo bar" {
                                        case "":
                                            v71 := GenS8{F0: 368, F1: 65535, F2: 2}
                                            _ = v71
                                        case "foo bar":
                                            if ((0 == 0) && (733 > (-(fi70)))) {
                                                if false {
                                                    if (v2 == v65) {
                                                        if len(v37) >= 2 {
                                                            v37 = v37[1:len(v37)]
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(fi70, v3)
                                                    var v72 float64 = 1e10
                                                    _ = v72
                                                }
                                                var v73 int8 = 1
                                                v67 -= (fi70 + v67)
                                                fmt.Println(v65, v73, v65)
                                                _ = v73
                                            }
                                        }
                                        if len(v37) > 0 {
                                            v37[0] = (fi68 >> ((2) & 63))
                                        }
                                        _ = fi70
                                    }
                                    for fi74 := 0; fi74 < 20; fi74++ {
                                        delete(v1, ("" + ("\n\t" + "\n\t")))
                                        _ = len(v1)
                                        if len(v37) > 0 {
                                            v37[0] = 42
                                        }
                                        _ = fi74
                                    }
                                }
                                if v6 {
                                    var v75 uint32 = (v2 / (((v65 - v2)) | 1))
                                    _ = v75
                                }
                                var v76 uint16 = 32767
                                _ = v76
                            }
                            _ = len(v37)
                            _ = fi68
                        }
                        var v77 int64 = -9223372036854775808
                        _ = v67
                        _ = v77
                    } else {
                        v78 := GenS5{F0: 65535}
                        _ = v78
                    }
                }
                _ = v65
            } else {
                fmt.Println(v4)
                if v6 {
                    if (682 > 42) {
                        v79 := v5[(v58 + v58)]
                        v80 := GenS9{F0: "a", F1: genFn0(765, v58, v4)}
                        _ = v79
                        _ = v80
                    }
                } else {
                    for fi81 := 0; fi81 < 14; fi81++ {
                        for fi82 := 0; fi82 < 4; fi82++ {
                            v37 = append(v37, (fi82 | 32768))
                            fmt.Println(v6)
                            _ = len(v37)
                            _ = fi82
                        }
                        _ = fi81
                    }
                    v83 := GenS1{F0: 1, F1: -119, F2: -1}
                    v58 -= v58
                    _ = v83
                }
            }
        }
        _ = len(v5)
    }
    v84 := GenS10{F0: 14, F1: "foo bar"}
    var v85 uint8 = 65
    if v6 {
        fmt.Println(v6)
    } else {
        var v86 int8 = 26
        _ = v86
    }
    v6 = ((521 >= 255) || false)
    v85 -= (v85 << ((372) & 63))
    switch v58 {
    case 127:
        for fi87 := 0; fi87 < 2; fi87++ {
            var v88 uint = 95
            _ = fi87
            _ = v88
        }
    case 32767:
        for fi89 := 0; fi89 < 9; fi89++ {
            for fi90 := 0; fi90 < 5; fi90++ {
                fmt.Println(fi89)
                if (false && v6) {
                    v2 = 1
                }
                var v91 uint64 = 0
                _ = fi90
                _ = v91
            }
            var v92 string = ("a" + "")
            _ = fi89
            _ = v92
        }
    case -9223372036854775808:
        v93 := GenS10{F0: 42, F1: ((((("foo bar" + "foo bar") + "") + "foo bar") + "hello") + "foo bar")}
        _ = v93
    default:
        var v94 bool = true
        _ = v94
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v84
    _ = v85
}
