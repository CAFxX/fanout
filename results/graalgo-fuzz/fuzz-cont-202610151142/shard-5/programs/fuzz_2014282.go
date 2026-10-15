package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 int16; F2 uint64 }
func genFn0(p0 bool, p1 string) uint {
    var v5 float64 = 0.1
    for fi6 := 0; fi6 < 5; fi6++ {
        switch fi6 {
        case 425:
            var v7 int8 = -125
            _ = v7
        case 128:
            var v8 bool = true
            _ = v8
        case 850:
            fmt.Println(p1, fi6, p0)
        default:
            v9 := GenS1{F0: p1}
            _ = v9
        }
        var v10 uint32 = 42
        var v11 bool = p0
        _ = fi6
        _ = v10
        _ = v11
    }
    var v12 string = (("foo bar" + "foo bar") + (p1 + p1))
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(v5)
        switch ("hello" + p1) {
        case "":
            v14 := GenS1{F0: (v12 + (v12 + ("\n\t" + p1)))}
            _ = v14
        case "hello":
            var v15 string = (((((p1 + "") + (p1 + (("a" + ((v12 + v12) + ("hello" + ("\n\t" + ("hello" + ""))))) + p1))) + "hello") + "") + "foo bar")
            _ = v15
        case "foo bar":
            for fi16 := 0; fi16 < 11; fi16++ {
                if (0 <= 576) {
                    v17 := []int{-9223372036854775808, (fi13 / (((fi13 % ((fi16) | 1))) | 1)), fi13}
                    _ = v17
                }
                _ = fi16
            }
        }
        switch fi13 {
        case 0:
            p0 = false
        case 9223372036854775807:
            v18 := []int{0, (fi13 + (^(fi13))), fi13, (fi13 % (((fi13 + (fi13 / ((fi13) | 1)))) | 1))}
            _ = v18
        case 1:
            var v19 uint64 = 1
            _ = v19
        }
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v5
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 2
    v3 := []int{-89, 74}
    var v4 int16 = 1
    var v20 string = ("foo bar" + "\n\t")
    var v21 int64 = 272
    var v22 string = v20
    v23 := GenS1{F0: v20}
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 float32 = 1.5
        for fi26 := 0; fi26 < 11; fi26++ {
            if ((3.14159 + -2.25) <= v25) {
                var v27 int16 = (v4 & int16(fi24))
                var v28 bool = false
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v27
                _ = v28
            }
            var v29 int64 = 788
            _ = fi26
            _ = v29
        }
        v2 = 42
        _ = fi24
        _ = v25
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        v31 := GenS2{F0: (v21 * v21)}
        _ = fi30
        _ = v31
    }
    var v32 float32 = (3.14159 + (3.14159 * ((((float32(1.5) / (((1.5 - (1e10 * (0.1 / 1.5))) / 1.0) / float32(float64((0.1 * (float32(((1e10 * 1.5) + 0.1)) * -2.25)))))) * (1e10 - float32(1e10))) - ((0.0 / (-2.25 + float32(-2.25))) * -2.25)) - 1.5)))
    fmt.Println(v4, v22)
    var v33 int64 = v21
    if false {
        for fi34 := 0; fi34 < 18; fi34++ {
            if (0.0 < 1e10) {
                fmt.Println(v22, v33)
            }
            _ = fi34
        }
    }
    var v35 float32 = (v32 - float32(0.0))
    if (1 != 42) {
        if (42 <= 42) {
            fmt.Println(v22, v4, v21)
            switch 76 {
            case -2147483648:
                v36 := GenS2{F0: v33}
                _ = v36
            case 32767:
                v21 -= 42
            default:
                v37 := GenS3{F0: v22}
                _ = v37
            }
        }
        for fi38 := 0; fi38 < 1; fi38++ {
            var v39 uint64 = v2
            if true {
                var v40 uint32 = 2
                var v41 bool = false
                var v42 float64 = 0.0
                _ = v40
                _ = v41
                _ = v42
            }
            v39 = v2
            _ = fi38
            _ = v39
        }
        fmt.Println(v2, v35, v22)
    }
    v3 = append(v3, int(v4))
    for fi43 := 0; fi43 < 13; fi43++ {
        switch v20 {
        case "hello":
            for fi44 := 0; fi44 < 10; fi44++ {
                delete(v1, v22)
                _ = fi44
            }
        }
        switch ((v22 + ((((v20 + (v20 + ("\n\t" + v20))) + v22) + ((v22 + v22) + v22)) + v20)) + "hello") {
        case "hello":
            if false {
                if true {
                    fmt.Println(v2, v35)
                    v21 = (v21 | (-(v21)))
                    if len(v3) > 0 {
                        v3[0] = fi43
                    }
                }
            } else {
                switch -1 {
                case -2147483648:
                    for fi45 := 0; fi45 < 11; fi45++ {
                        var v46 uint64 = (v2 | (v2 / ((568) | 1)))
                        v20 = "\n\t"
                        var v47 uint8 = 254
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                default:
                    for fi48 := 0; fi48 < 8; fi48++ {
                        var v49 float32 = 1e10
                        var v50 uint64 = v2
                        if (false || true) {
                            if false {
                                var v51 string = v22
                                _ = v51
                            } else {
                                var v52 float64 = (1e10 / (1.5 - 0.1))
                                if ((!true) || false) {
                                    var v53 int = (-(fi43))
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v54 := GenS1{F0: v20}
                                    _ = v53
                                    _ = v54
                                }
                                _ = v52
                            }
                            if (127 <= 0) {
                                if (0 != (v50 & (v2 | v2))) {
                                    if len(v3) > 0 {
                                        v3[0] = (^(fi43))
                                    }
                                    v55 := GenS3{F0: ("a" + v20)}
                                    _ = v55
                                } else {
                                    var v56 int16 = 1
                                    var v57 int16 = 182
                                    if true {
                                        var v58 uint16 = 32767
                                        var v59 uint = 42
                                        var v60 int8 = -94
                                        var v61 uint16 = v58
                                        v62 := GenS2{F0: -102}
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if (-1 == 2147483647) {
                                            v35 -= ((-2.25 / 3.14159) * v32)
                                            var v63 uint64 = v50
                                            _ = v63
                                        }
                                        _ = v58
                                        _ = v59
                                        _ = v60
                                        _ = v61
                                        _ = v62
                                    } else {
                                        delete(v1, "")
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = len(v3)
                                    }
                                    var v64 uint32 = 2147483647
                                    _ = len(v3)
                                    switch "" {
                                    case "\n\t":
                                        var v65 float32 = v32
                                        _ = v65
                                    case "a":
                                        var v66 float32 = 0.0
                                        v67 := GenS2{F0: 152}
                                        _ = v66
                                        _ = v67
                                    case "hello":
                                        v20 = v22
                                        var v68 uint16 = 2
                                        _ = v68
                                    default:
                                        var v69 int8 = 42
                                        _ = v69
                                    }
                                    fmt.Println(v33, v50, v32)
                                    _ = v56
                                    _ = v57
                                    _ = v64
                                }
                                var v70 int32 = -947
                                v35 = (0.0 * (v35 / v32))
                                switch fi48 {
                                case -629:
                                    v35 = v35
                                }
                                if ((false || (true && (((v20 + "a") != v20) || false))) || false) {
                                    if ((!(fi48 < 42)) && false) {
                                        var v71 uint = 1
                                        fmt.Println(fi43)
                                        _ = v71
                                    }
                                }
                                var v72 uint = genFn0(((v4 == v4) && true), (v20 + v20))
                                _ = v70
                                _ = v72
                            } else {
                                var v73 int32 = -2147483648
                                v20 = v20
                                var v74 int32 = v73
                                if (v74 != (v73 >> ((2) & 63))) {
                                    var v75 uint32 = 774
                                    v76 := GenS3{F0: "a"}
                                    if true {
                                        var v77 int16 = 42
                                        v20 = "hello"
                                        fmt.Println(v77, v75)
                                        v21 -= 0
                                        _ = v77
                                    }
                                    _ = v75
                                    _ = v76
                                }
                                _ = v73
                                _ = v74
                            }
                        } else {
                            switch "a" {
                            case "":
                                _ = len(v1)
                            case "hello":
                                var v78 uint8 = 128
                                _ = v78
                            case "\n\t":
                                if true {
                                    v79 := GenS4{F0: (-(v21)), F1: (^(v4)), F2: v2}
                                    _ = len(v3)
                                    var v80 int16 = v4
                                    _ = v79
                                    _ = v80
                                }
                                v1[v22] = fi43
                            }
                            _ = len(v1)
                        }
                        v81 := GenS1{F0: ((v22 + "") + "hello")}
                        switch 42 {
                        case 628:
                            if ((v20 + (v20 + "hello")) < v22) {
                                var v82 string = v20
                                if ((128 == 127) || (true || (!false))) {
                                    fmt.Println(fi48, v21)
                                    var v83 uint8 = 255
                                    if true {
                                        fmt.Println(fi48, v35)
                                        if len(v3) > 0 {
                                            v3[0] = fi43
                                        }
                                    }
                                    _ = v83
                                } else {
                                    if (true && (!((!true) || (v50 > 552)))) {
                                        v3 = append(v3, fi43)
                                    } else {
                                        var v84 int64 = v21
                                        _ = len(v3)
                                        _ = v84
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = (fi43 | fi48)
                                    }
                                    v82 = (("" + "hello") + v20)
                                    if true {
                                        _ = len(v3)
                                    } else {
                                        var v85 uint16 = 65535
                                        _ = v85
                                    }
                                    var v86 uint8 = 0
                                    _ = v86
                                }
                                v20 = ((v20 + "") + ((v20 + "") + (v20 + v82)))
                                _ = v82
                            }
                        case 296:
                            v20 = (v20 + v22)
                        }
                        var v87 int64 = (v21 + int64(fi48))
                        _ = fi48
                        _ = v49
                        _ = v50
                        _ = v81
                        _ = v87
                    }
                }
                if false {
                    switch fi43 {
                    case 128:
                        v20 = (v20 + ((v22 + "a") + "\n\t"))
                    case 323:
                        v22 = "foo bar"
                    case -531:
                        for fi88 := 0; fi88 < 16; fi88++ {
                            var v89 int64 = v33
                            var v90 uint64 = (v2 << ((126) & 63))
                            fmt.Println(v22, v21, v32)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi88
                            _ = v89
                            _ = v90
                        }
                    }
                    switch ("a" + ((v22 + (v22 + "a")) + "foo bar")) {
                    case "\n\t":
                        v91 := GenS3{F0: (("hello" + "hello") + (("hello" + (v20 + v20)) + v20))}
                        _ = v91
                    default:
                        switch (fi43 * fi43) {
                        case 1:
                            v20 = ((v22 + v22) + v22)
                        }
                    }
                    _ = len(v3)
                } else {
                    for fi92 := 0; fi92 < 2; fi92++ {
                        var v93 uint = 42
                        var v94 uint64 = (v2 + (v2 >> ((172) & 63)))
                        v95 := GenS1{F0: (v20 + "foo bar")}
                        _ = fi92
                        _ = v93
                        _ = v94
                        _ = v95
                    }
                    for fi96 := 0; fi96 < 14; fi96++ {
                        v97 := v1[("hello" + (v22 + (v20 + "hello")))]
                        _ = len(v3)
                        v98 := GenS3{F0: "a"}
                        _ = fi96
                        _ = v97
                        _ = v98
                    }
                }
                var v99 uint32 = 112
                _ = v99
            }
        }
        _ = fi43
    }
    var v100 int8 = 0
    if (!true) {
        fmt.Println(v22, v32, v21)
        fmt.Println(v20, v21)
        var v101 float32 = v35
        _ = v101
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v100
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v4
}
