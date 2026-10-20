package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int }
type GenS2 struct { F0 uint64; F1 int32; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool; F1 int8; F2 string }
type GenS8 struct { F0 bool; F1 bool; F2 int16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int }
func genFn0() int16 {
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: 42, F1: 726, F2: 42}
        var v8 string = "hello"
        var v9 float32 = ((3.14159 - (1e10 / float32(1.5))) / float32(1.5))
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 float64 = 1e10
    if (2 != 128) {
        var v11 bool = (65535 > 2)
        var v12 float64 = 1.5
        if (v11 && v11) {
            if v11 {
                var v13 uint16 = 32767
                if (!(474 <= 856)) {
                    v14 := []int{-232, 1}
                    _ = v14
                } else {
                    if (2147483648 <= 779) {
                        v15 := []int{127}
                        _ = v15
                    }
                    v16 := GenS1{F0: 0, F1: (^(v13)), F2: -442}
                    switch 127 {
                    case -9223372036854775808:
                        var v17 uint16 = v13
                        _ = v17
                    case 5:
                        v18 := GenS1{F0: 0, F1: v13, F2: 1}
                        _ = v18
                    case 9223372036854775807:
                        var v19 uint = 502
                        _ = v19
                    default:
                        v11 = true
                    }
                    _ = v16
                }
                _ = v13
            }
            v10 = v10
            switch ("foo bar" + ((("\n\t" + "") + ("foo bar" + ((("foo bar" + "a") + (("\n\t" + "\n\t") + ("a" + ""))) + "hello"))) + "")) {
            case "":
                v20 := GenS2{F0: 42, F1: 0, F2: -342}
                _ = v20
            case "hello":
                if (("a" + "a") <= "\n\t") {
                    switch "a" {
                    case "hello":
                        v11 = true
                    case "foo bar":
                        var v21 uint16 = 456
                        _ = v21
                    case "":
                        v22 := []int{857, 4294967295, 1}
                        _ = v22
                    default:
                        if (((((v12 - 1.5) * (v10 - v12)) / (0.0 - (v10 - (v12 + 1e10)))) * v10) == v10) {
                            for fi23 := 0; fi23 < 17; fi23++ {
                                var v24 string = "hello"
                                v11 = v11
                                _ = fi23
                                _ = v24
                            }
                            if v11 {
                                v25 := []int{32767, 1, -2147483648}
                                if len(v25) >= 2 {
                                    v25 = v25[1:len(v25)]
                                }
                                for fi26 := 0; fi26 < 9; fi26++ {
                                    var v27 uint8 = 52
                                    if (0 < -773) {
                                        var v28 string = "hello"
                                        v29 := GenS2{F0: 1, F1: 42, F2: fi26}
                                        _ = v28
                                        _ = v29
                                    } else {
                                        for fi30 := 0; fi30 < 18; fi30++ {
                                            v31 := GenS2{F0: 287, F1: 0, F2: fi30}
                                            switch "foo bar" {
                                            case "a":
                                                var v32 int16 = 1
                                                _ = v32
                                            case "hello":
                                                v33 := GenS3{F0: 42}
                                                _ = v33
                                            case "foo bar":
                                                switch "a" {
                                                case "foo bar":
                                                    var v34 uint32 = 1
                                                    _ = v34
                                                }
                                            default:
                                                var v35 string = "a"
                                                _ = v35
                                            }
                                            var v36 int32 = -2147483648
                                            _ = fi30
                                            _ = v31
                                            _ = v36
                                        }
                                        if ((v27 - v27) != 42) {
                                            var v37 int8 = -128
                                            switch (fi26 << ((65535) & 63)) {
                                            case 0:
                                                if true {
                                                    var v38 float64 = 0.0
                                                    _ = v38
                                                }
                                            case 820:
                                                v39 := GenS4{F0: fi26, F1: 1, F2: 4294967295}
                                                _ = v39
                                            case 42:
                                                v40 := GenS5{F0: 357, F1: 2}
                                                _ = v40
                                            default:
                                                if ((v27 - (^(v27))) == (v27 >> ((135) & 63))) {
                                                    for fi41 := 0; fi41 < 8; fi41++ {
                                                        if (v11 || v11) {
                                                            v25 = append(v25, fi26)
                                                            fmt.Println(v12, v11, v27)
                                                            v37 = -128
                                                        }
                                                        var v42 float32 = 1.5
                                                        v43 := GenS6{F0: v27}
                                                        _ = fi41
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    if (-1 == -1) {
                                                        var v44 uint = 0
                                                        if v11 {
                                                            var v45 uint32 = 2147483647
                                                            if (!v11) {
                                                                if v11 {
                                                                    v25 = append(v25, (fi26 & int(v27)))
                                                                    v37 = v37
                                                                    v10 -= v12
                                                                }
                                                                if (uint64(v44) != 4294967295) {
                                                                    v46 := GenS7{F0: (((42 >= 42) && (32767 <= 0)) || (v11 || true)), F1: 127, F2: ("a" + "")}
                                                                    _ = v46
                                                                } else {
                                                                    if (v44 == v44) {
                                                                        var v47 float32 = float32((v12 + v10))
                                                                        _ = v47
                                                                    } else {
                                                                        var v48 float32 = -2.25
                                                                        if v11 {
                                                                            v48 -= (v48 / v48)
                                                                            v49 := GenS2{F0: 671, F1: -1, F2: fi26}
                                                                            for fi50 := 0; fi50 < 13; fi50++ {
                                                                                v45 -= v45
                                                                                _ = fi50
                                                                            }
                                                                            _ = v49
                                                                        } else {
                                                                            var v51 uint8 = (v27 << ((65535) & 63))
                                                                            _ = v51
                                                                        }
                                                                        _ = v48
                                                                    }
                                                                    var v52 int8 = 100
                                                                    _ = v52
                                                                }
                                                                var v53 uint8 = v27
                                                                _ = v53
                                                            }
                                                            _ = v45
                                                        } else {
                                                            var v54 uint8 = 255
                                                            var v55 uint8 = (v27 / ((v54) | 1))
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                        _ = v44
                                                    } else {
                                                        if v11 {
                                                            v37 = -89
                                                            v56 := GenS7{F0: (366 == 0), F1: (v37 | (v37 % (((v37 % ((100) | 1))) | 1))), F2: ("" + "")}
                                                            if len(v25) >= 2 {
                                                                v25 = v25[1:len(v25)]
                                                            }
                                                            _ = v56
                                                        }
                                                        switch (fi26 | fi26) {
                                                        case -542:
                                                            switch int(v37) {
                                                            case -970:
                                                                fmt.Println(v37)
                                                            case 2147483647:
                                                                var v57 uint32 = 807
                                                                _ = v57
                                                            default:
                                                                v58 := GenS7{F0: false, F1: (-(v37)), F2: ""}
                                                                _ = v58
                                                            }
                                                        default:
                                                            v27 = (v27 * (v27 / ((v27) | 1)))
                                                        }
                                                        if len(v25) >= 2 {
                                                            v25 = v25[1:len(v25)]
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v27)
                                                    if len(v25) >= 2 {
                                                        v25 = v25[1:len(v25)]
                                                    }
                                                }
                                            }
                                            _ = v37
                                        }
                                    }
                                    _ = fi26
                                    _ = v27
                                }
                                _ = v25
                            }
                            var v59 uint64 = 2
                            _ = v59
                        }
                    }
                } else {
                    for fi60 := 0; fi60 < 15; fi60++ {
                        v11 = v11
                        var v61 int32 = 0
                        for fi62 := 0; fi62 < 12; fi62++ {
                            fmt.Println(v61, v11, v10)
                            var v63 string = ("hello" + "hello")
                            _ = fi62
                            _ = v63
                        }
                        _ = fi60
                        _ = v61
                    }
                    var v64 uint64 = 913
                    _ = v64
                }
            case "foo bar":
                for fi65 := 0; fi65 < 17; fi65++ {
                    for fi66 := 0; fi66 < 5; fi66++ {
                        fmt.Println(fi66, v12)
                        var v67 string = ("\n\t" + "a")
                        _ = fi66
                        _ = v67
                    }
                    fmt.Println(v11)
                    _ = fi65
                }
            default:
                fmt.Println(v12, v10)
            }
        } else {
            switch 273 {
            case 32767:
                for fi68 := 0; fi68 < 1; fi68++ {
                    var v69 uint = 4294967295
                    switch (fi68 + (fi68 * -2147483648)) {
                    case 1:
                        var v70 string = "foo bar"
                        _ = v70
                    case -2147483648:
                        v71 := []int{(fi68 | -135), 32767, (-(fi68))}
                        _ = v71
                    }
                    _ = fi68
                    _ = v69
                }
            case 173:
                var v72 float64 = 1e10
                _ = v72
            default:
                v73 := []int{1}
                _ = v73
            }
        }
        _ = v11
        _ = v12
    } else {
        var v74 float64 = v10
        for fi75 := 0; fi75 < 1; fi75++ {
            v76 := GenS8{F0: false, F1: false, F2: 27}
            _ = fi75
            _ = v76
        }
        v77 := GenS9{F0: 65535}
        _ = v74
        _ = v77
    }
    _ = v10
    return 1
}
func genFn1(p0 bool, p1 uint32, p2 int8) int16 {
    p0 = false
    p0 = true
    v78 := GenS5{F0: 102, F1: 1}
    fmt.Println(p1, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v78
    return genFn0()
}

func main() {
    var v1 int64 = 241
    v2 := make(map[int]string)
    v3 := []string{}
    var v4 uint32 = 939
    var v5 int64 = v1
    if (42 <= int8(v5)) {
        var v79 int = 128
        v80 := GenS3{F0: 0}
        v5 -= v1
        _ = v79
        _ = v80
    } else {
        fmt.Println(v5, v4)
        v5 = v1
        var v81 int32 = -811
        _ = v81
    }
    for fi82 := 0; fi82 < 14; fi82++ {
        var v83 string = "foo bar"
        v84 := GenS10{F0: ((-2147483648 == -717) && true)}
        _ = fi82
        _ = v83
        _ = v84
    }
    _ = len(v3)
    var v85 uint16 = 453
    _ = len(v3)
    v86 := v2[812]
    v85 = (v85 ^ (v85 % (((v85 % ((v85) | 1))) | 1)))
    switch -1 {
    case 359:
        v5 = 0
    case -217:
        for fi87 := 0; fi87 < 9; fi87++ {
            _ = len(v2)
            v4 += v4
            _ = fi87
        }
    default:
        for fi88 := 0; fi88 < 6; fi88++ {
            v86 = v86
            _ = fi88
        }
    }
    if true {
        v4 = 1
        switch 128 {
        case 43:
            v89 := GenS3{F0: 42}
            _ = v89
        case -1:
            v5 += 9223372036854775807
        case -2147483648:
            var v90 uint64 = 2
            _ = v90
        }
        if len(v3) > 0 {
            v3[0] = v86
        }
    } else {
        v91 := GenS11{F0: 9223372036854775807}
        _ = v91
    }
    if false {
        v85 = 1
    }
    var v92 uint = 520
    v85 += (v85 - (v85 / (((^(v85))) | 1)))
    v1 -= v1
    var v93 int32 = 424
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v85
    _ = v86
    _ = v92
    _ = v93
}
