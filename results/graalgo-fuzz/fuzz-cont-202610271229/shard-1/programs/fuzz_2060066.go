package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8 }
func genFn0(p0 int16, p1 int16, p2 int) uint {
    var v6 float32 = 1.5
    p2 -= p2
    fmt.Println(p2, p1, p2)
    v7 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 uint32, p1 uint64) uint8 {
    var v8 uint16 = 42
    if (float64(0.0) < (float64(-2.25) - 1.5)) {
        v9 := GenS1{F0: 4294967295}
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 float32 = 0.1
            switch ("" + "\n\t") {
            case "":
                if (v8 == 1) {
                    fmt.Println(v8, p1)
                }
            case "foo bar":
                p0 = 4294967295
            case "a":
                var v12 float64 = ((((0.0 * 1.5) * 1.5) - 0.1) + 1e10)
                _ = v12
            }
            v11 = ((((1e10 + v11) + 1e10) + (v11 - (float32((float64(v11) * (0.1 - -2.25))) / 1e10))) - v11)
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        p1 += 18446744073709551615
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            if (false && true) {
                var v15 int32 = 0
                _ = v15
            } else {
                v16 := GenS1{F0: 4294967295}
                _ = v16
            }
            if false {
                var v17 float64 = (0.0 / 1.5)
                switch fi14 {
                case -711:
                    v18 := GenS1{F0: 312}
                    _ = v18
                case 4294967295:
                    var v19 int = (fi13 - (-(fi13)))
                    _ = v19
                default:
                    var v20 uint8 = 128
                    _ = v20
                }
                _ = v17
            } else {
                v21 := []int{(-(fi13))}
                switch -427 {
                case 906:
                    var v22 uint8 = 122
                    _ = v22
                }
                if (-53 < -2147483648) {
                    if len(v21) > 0 {
                        v21[0] = (fi13 / ((fi13) | 1))
                    }
                    v8 += 65535
                } else {
                    if len(v21) > 0 {
                        v21[0] = fi14
                    }
                    if (7 >= 112) {
                        var v23 string = "hello"
                        switch (("" + "") + v23) {
                        case "hello":
                            var v24 uint = 2
                            _ = v24
                        case "\n\t":
                            fmt.Println(p0)
                        case "foo bar":
                            var v25 float64 = float64(float32(((((1e10 * -2.25) * 1e10) - 1e10) * (0.1 / 1.0))))
                            _ = v25
                        }
                        _ = v23
                    } else {
                        var v26 int64 = -1
                        _ = v26
                    }
                    v27 := GenS2{F0: p0}
                    _ = v27
                }
                _ = v21
            }
            _ = fi14
        }
        p0 -= 4294967295
        switch "\n\t" {
        case "a":
            for fi28 := 0; fi28 < 2; fi28++ {
                fmt.Println(fi28, v8)
                p0 -= (p0 - 874)
                _ = fi28
            }
        default:
            fmt.Println(p1)
        }
        _ = fi13
    }
    switch "a" {
    case "\n\t":
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 uint64 = p1
            if (-1 != 2147483647) {
                if (true || (139 < 0)) {
                    v8 = 0
                    p1 += 4294967295
                    v8 = (^(v8))
                }
                v31 := GenS1{F0: 2}
                v32 := []int{4294967295, fi29, (fi29 + (fi29 / (((fi29 - (fi29 >> ((genFn0(0, 1, (fi29 ^ fi29))) & 63)))) | 1)))}
                _ = v31
                _ = v32
            }
            v33 := []int{fi29, -1, -1, fi29}
            _ = fi29
            _ = v30
            _ = v33
        }
    case "a":
        p1 = (p1 << ((18446744073709551615) & 63))
    default:
        v34 := []int{128, -1, -505}
        _ = v34
    }
    v35 := []int{1, 4294967295, 4294967295}
    _ = p0
    _ = p1
    _ = v35
    _ = v8
    return 127
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 4294967295
    var v3 uint8 = 127
    v4 := make(map[uint32]int)
    v5 := make(map[string]int)
    var v36 uint16 = 517
    for fi37 := 0; fi37 < 17; fi37++ {
        var v38 uint8 = genFn1(uint32(v3), 42)
        var v39 int = fi37
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 uint8 = uint8(v36)
    if true {
        var v41 int = -9223372036854775808
        v4[(v2 ^ (v2 & 852))] = v41
        if true {
            switch v41 {
            case -983:
                v41 -= (v41 ^ v41)
            case -975:
                var v42 uint16 = (v36 * v36)
                _ = v42
            }
            for fi43 := 0; fi43 < 9; fi43++ {
                delete(v1, "hello")
                v44 := GenS3{F0: fi43, F1: v2, F2: 869}
                _ = fi43
                _ = v44
            }
        }
        _ = v41
    } else {
        fmt.Println(v40, v2, v36)
        v40 = genFn1(v2, 2)
    }
    for fi45 := 0; fi45 < 15; fi45++ {
        var v46 string = ""
        if ((v46 + v46) > (v46 + "")) {
            var v47 uint16 = (v36 | 32767)
            var v48 uint64 = 314
            _ = v47
            _ = v48
        } else {
            var v49 int8 = 55
            if ((v3 >> ((uint(v36)) & 63)) >= v3) {
                v50 := []int{(fi45 | -445)}
                _ = v50
            }
            if (4294967295 != (v2 >> ((genFn0(-1, 1, (fi45 << ((2) & 63)))) & 63))) {
                var v51 int32 = -313
                var v52 int16 = -1
                _ = v51
                _ = v52
            } else {
                fmt.Println(v49, v49, v36)
            }
            _ = v49
        }
        fmt.Println(v2)
        _ = fi45
        _ = v46
    }
    var v53 int8 = 1
    switch ("a" + "foo bar") {
    case "hello":
        v54 := []int{-9223372036854775808, 32767}
        _ = v54
    case "foo bar":
        var v55 int64 = 0
        _ = v55
    case "a":
        switch "foo bar" {
        case "\n\t":
            if (("\n\t" + "foo bar") > ("a" + "")) {
                var v56 float32 = -2.25
                _ = v56
            } else {
                fmt.Println(v3)
            }
        case "foo bar":
            if false {
                for fi57 := 0; fi57 < 3; fi57++ {
                    var v58 int8 = (v53 >> ((1) & 63))
                    _ = fi57
                    _ = v58
                }
                for fi59 := 0; fi59 < 16; fi59++ {
                    v60 := GenS4{F0: (^(fi59))}
                    _ = fi59
                    _ = v60
                }
                var v61 uint32 = uint32(v36)
                _ = v61
            } else {
                v53 -= v53
                var v62 uint16 = (v36 % ((v36) | 1))
                _ = v62
            }
        default:
            switch -231 {
            case -68:
                v53 += (v53 - v53)
            case 2147483647:
                var v63 int32 = -1
                _ = v63
            case -2147483648:
                v3 = v3
            default:
                if (v36 < (v36 * v36)) {
                    var v64 uint64 = 658
                    if true {
                        v65 := GenS4{F0: 576}
                        if (-1 != -816) {
                            if false {
                                for fi66 := 0; fi66 < 20; fi66++ {
                                    var v67 float32 = -2.25
                                    fmt.Println(fi66, v64)
                                    _ = fi66
                                    _ = v67
                                }
                                switch "\n\t" {
                                case "a":
                                    _ = len(v1)
                                default:
                                    v68 := v1[""]
                                    _ = v68
                                }
                                v2 = v2
                            } else {
                                var v69 bool = false
                                var v70 uint16 = v36
                                var v71 uint = 18446744073709551615
                                _ = v69
                                _ = v70
                                _ = v71
                            }
                        } else {
                            fmt.Println(v3, v3, v40)
                            v72 := GenS1{F0: 42}
                            _ = v72
                        }
                        v1["hello"] = (1000 >= 1)
                        _ = v65
                    } else {
                        var v73 float32 = ((0.1 + 1.5) - (0.0 / 0.1))
                        delete(v5, "")
                        fmt.Println(v40)
                        _ = v73
                    }
                    v74 := []int{int(v3)}
                    _ = v64
                    _ = v74
                }
            }
        }
    }
    var v75 uint16 = (v36 >> ((42) & 63))
    v75 = (v75 % ((v36) | 1))
    v75 = 1
    switch (("" + ("" + "a")) + "\n\t") {
    case "a":
        fmt.Println(v53, v53)
    case "foo bar":
        if true {
            switch ("a" + (("foo bar" + "") + "")) {
            case "foo bar":
                if (!(-850 == 1000)) {
                    switch -52 {
                    case -58:
                        v75 = (v75 | (v75 & 0))
                    case 123:
                        delete(v4, v2)
                    default:
                        var v76 int = 32768
                        _ = v76
                    }
                    for fi77 := 0; fi77 < 12; fi77++ {
                        var v78 int64 = -673
                        v79 := []int{}
                        v80 := GenS3{F0: fi77, F1: v2, F2: 2}
                        _ = fi77
                        _ = v78
                        _ = v79
                        _ = v80
                    }
                    var v81 int = 128
                    _ = v81
                } else {
                    v2 += (v2 * (v2 << ((4294967295) & 63)))
                    switch 2147483647 {
                    case 32768:
                        v53 = (-(v53))
                    default:
                        v82 := []int{}
                        _ = v82
                    }
                    v75 += v36
                }
            case "":
                v83 := GenS1{F0: 1}
                _ = v83
            case "a":
                var v84 float32 = 3.14159
                _ = v84
            default:
                v2 = v2
            }
        }
    }
    v75 -= (^(v75))
    v1[("\n\t" + ("hello" + "foo bar"))] = (42 > -129)
    v85 := GenS1{F0: 4294967295}
    v86 := GenS3{F0: 127, F1: v2, F2: 1}
    if (true && (748 != -9223372036854775808)) {
        v87 := []int{9223372036854775807}
        var v88 uint64 = 2
        var v89 string = "foo bar"
        _ = v87
        _ = v88
        _ = v89
    }
    v90 := []int{}
    v91 := GenS5{F0: v53}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v75:", v75)
    fmt.Println("slicelen v90:", len(v90), v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v53
    _ = v75
    _ = v85
    _ = v86
    _ = v90
    _ = v91
}
