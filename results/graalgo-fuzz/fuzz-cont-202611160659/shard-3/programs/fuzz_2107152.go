package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int8; F2 uint }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8; F1 int; F2 int32 }
func genFn0(p0 uint16, p1 string, p2 uint32) uint64 {
    var v6 int64 = -1
    v7 := []int{-135, -2147483648}
    var v8 float32 = 1.5
    var v9 int8 = 1
    if ((p2 ^ 65535) > 65535) {
        var v10 string = ("foo bar" + p1)
        var v11 int8 = (v9 | (v9 - (v9 ^ -1)))
        _ = v10
        _ = v11
    } else {
        var v12 int64 = v6
        if len(v7) > 0 {
            v7[0] = 32768
        }
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := []string{"hello"}
    var v2 uint = 513
    v3 := []int{}
    var v4 int32 = 0
    var v5 int16 = 1000
    for fi13 := 0; fi13 < 6; fi13++ {
        fmt.Println(fi13, v5, v2)
        v4 = -414
        v14 := GenS1{F0: (fi13 - (fi13 - 42))}
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: 1, F1: 1}
    var v16 uint64 = 939
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    fmt.Println(v4)
    _ = len(v1)
    var v17 int16 = v5
    v17 -= v5
    var v18 uint32 = 240
    switch 42 {
    case 32768:
        _ = len(v1)
    case -1:
        var v19 bool = true
        _ = v19
    }
    var v20 int16 = (v17 | 0)
    if ((v2 % (((v2 / (((v2 - v2)) | 1))) | 1)) >= 0) {
        v21 := GenS2{F0: 255, F1: 32767}
        if (3.14159 == 3.14159) {
            var v22 bool = ("\n\t" > "hello")
            if v22 {
                fmt.Println(v5)
                if v22 {
                    v4 = (v4 & 1)
                    if (32767 <= 1) {
                        if ("foo bar" > "\n\t") {
                            v23 := GenS2{F0: 1, F1: 32767}
                            var v24 string = "hello"
                            if true {
                                var v25 uint32 = 769
                                if (1 > 128) {
                                    if (!false) {
                                        for fi26 := 0; fi26 < 1; fi26++ {
                                            v27 := GenS1{F0: -9223372036854775808}
                                            _ = fi26
                                            _ = v27
                                        }
                                        if ((1.5 > (1e10 + 3.14159)) || (v4 != (v4 >> (((v2 << ((1) & 63))) & 63)))) {
                                            if v22 {
                                                switch 29 {
                                                case 128:
                                                    var v28 uint8 = 255
                                                    _ = v28
                                                }
                                                for fi29 := 0; fi29 < 5; fi29++ {
                                                    switch v24 {
                                                    case "\n\t":
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v24 = ("\n\t" + ("\n\t" + (("hello" + v24) + v24)))
                                                        case "foo bar":
                                                            if (0 <= 42) {
                                                                switch (("" + v24) + (v24 + "hello")) {
                                                                case "a":
                                                                    if false {
                                                                        v30 := GenS1{F0: (fi29 % ((fi29) | 1))}
                                                                        var v31 int64 = -9223372036854775808
                                                                        if (v16 <= (v16 << ((18446744073709551615) & 63))) {
                                                                            v25 += v18
                                                                            for fi32 := 0; fi32 < 14; fi32++ {
                                                                                var v33 string = (v24 + (((((("\n\t" + v24) + v24) + (v24 + ("a" + v24))) + (((v24 + v24) + v24) + v24)) + (("hello" + v24) + "a")) + v24))
                                                                                if (0 >= 32767) {
                                                                                    var v34 uint16 = 294
                                                                                    switch "a" {
                                                                                    case "a":
                                                                                        if ((((v16 + v16) >= 198) || ((v4 | (v4 / ((v4) | 1))) != v4)) && v22) {
                                                                                            var v35 int8 = 1
                                                                                            _ = v35
                                                                                        }
                                                                                    case "":
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = fi32
                                                                                        }
                                                                                    case "\n\t":
                                                                                        var v36 uint32 = v18
                                                                                        _ = v36
                                                                                    }
                                                                                    _ = v34
                                                                                }
                                                                                var v37 uint8 = uint8(v2)
                                                                                v37 = 55
                                                                                v4 += (v4 * -618)
                                                                                _ = fi32
                                                                                _ = v33
                                                                                _ = v37
                                                                            }
                                                                        } else {
                                                                            for fi38 := 0; fi38 < 14; fi38++ {
                                                                                switch v24 {
                                                                                case "a":
                                                                                    if (!(-436 == 42)) {
                                                                                        v39 := GenS2{F0: 255, F1: 32767}
                                                                                        _ = v39
                                                                                    } else {
                                                                                        v4 -= (v4 | -1)
                                                                                        if v22 {
                                                                                            var v40 bool = (!(v22 && v22))
                                                                                            var v41 uint64 = v16
                                                                                            var v42 uint32 = v18
                                                                                            _ = v40
                                                                                            _ = v41
                                                                                            _ = v42
                                                                                        } else {
                                                                                            if true {
                                                                                                if (v22 && ((v4 % ((v4) | 1)) < v4)) {
                                                                                                    v18 -= v18
                                                                                                }
                                                                                                fmt.Println(v31)
                                                                                                var v43 string = (v24 + "a")
                                                                                                if (true && v22) {
                                                                                                    v44 := GenS3{F0: (v16 - 18446744073709551615)}
                                                                                                    if false {
                                                                                                        if v22 {
                                                                                                            if (43 <= v16) {
                                                                                                                switch v24 {
                                                                                                                case "hello":
                                                                                                                    if (100 >= 0) {
                                                                                                                        if len(v3) >= 2 {
                                                                                                                            v3 = v3[1:len(v3)]
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        v45 := GenS3{F0: (v16 + v16)}
                                                                                                                        fmt.Println(v31, fi38, v31)
                                                                                                                        _ = v45
                                                                                                                    }
                                                                                                                case "a":
                                                                                                                    v46 := GenS4{F0: v22, F1: 76, F2: 341}
                                                                                                                    _ = v46
                                                                                                                }
                                                                                                                switch fi38 {
                                                                                                                case 32768:
                                                                                                                    var v47 int8 = -102
                                                                                                                    _ = v47
                                                                                                                }
                                                                                                            } else {
                                                                                                                fmt.Println(v4, v43, v16)
                                                                                                                var v48 float32 = ((-2.25 - (0.0 * 0.0)) / 1.0)
                                                                                                                _ = v48
                                                                                                            }
                                                                                                            fmt.Println(v20)
                                                                                                        } else {
                                                                                                            v25 = (v25 ^ 42)
                                                                                                        }
                                                                                                    }
                                                                                                    v4 -= -2147483648
                                                                                                    _ = v44
                                                                                                }
                                                                                                if (0 < -128) {
                                                                                                    v43 = v24
                                                                                                    v3 = append(v3, (fi38 * fi38))
                                                                                                    if len(v3) >= 2 {
                                                                                                        v3 = v3[1:len(v3)]
                                                                                                    }
                                                                                                    var v49 uint64 = uint64(v18)
                                                                                                    _ = v49
                                                                                                } else {
                                                                                                    var v50 int16 = v17
                                                                                                    var v51 int16 = (v5 ^ (v17 % ((1000) | 1)))
                                                                                                    _ = v50
                                                                                                    _ = v51
                                                                                                }
                                                                                                _ = v43
                                                                                            }
                                                                                            if v22 {
                                                                                                v20 += 943
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = v24
                                                                                                }
                                                                                                if v22 {
                                                                                                    v24 = v24
                                                                                                }
                                                                                                v52 := GenS3{F0: v16}
                                                                                                switch fi38 {
                                                                                                case -9223372036854775808:
                                                                                                    var v53 uint16 = 756
                                                                                                    _ = v53
                                                                                                case 4294967295:
                                                                                                    var v54 float64 = (1.5 - ((3.14159 - 1.5) * 0.0))
                                                                                                    _ = v54
                                                                                                default:
                                                                                                    var v55 float64 = (-2.25 / (-2.25 + 0.0))
                                                                                                    _ = v55
                                                                                                }
                                                                                                var v56 int64 = (-(v31))
                                                                                                _ = v52
                                                                                                _ = v56
                                                                                            } else {
                                                                                                v25 -= 206
                                                                                                if false {
                                                                                                    v24 = v24
                                                                                                    if (127 == 31) {
                                                                                                        var v57 uint8 = 121
                                                                                                        _ = v57
                                                                                                    }
                                                                                                    if (v4 == v4) {
                                                                                                        var v58 uint = v2
                                                                                                        v25 = (v18 % ((v18) | 1))
                                                                                                        _ = v58
                                                                                                    } else {
                                                                                                        _ = len(v1)
                                                                                                        v16 = (v16 - genFn0(uint16(v16), v24, v18))
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v59 := GenS3{F0: v16}
                                                                                            _ = v59
                                                                                        }
                                                                                    }
                                                                                case "foo bar":
                                                                                    if v22 {
                                                                                        var v60 uint16 = 2
                                                                                        _ = v60
                                                                                    }
                                                                                    v61 := GenS3{F0: (^(v16))}
                                                                                    _ = v61
                                                                                case "\n\t":
                                                                                    _ = len(v1)
                                                                                }
                                                                                if v22 {
                                                                                    var v62 int32 = v4
                                                                                    _ = len(v1)
                                                                                    v22 = v22
                                                                                    _ = v62
                                                                                }
                                                                                var v63 float64 = 0.1
                                                                                if (127 <= 58) {
                                                                                    v2 -= (v2 | v2)
                                                                                } else {
                                                                                    var v64 int8 = -64
                                                                                    _ = v64
                                                                                }
                                                                                _ = fi38
                                                                                _ = v63
                                                                            }
                                                                            var v65 int64 = int64(v5)
                                                                            v2 -= v2
                                                                            _ = v65
                                                                        }
                                                                        _ = v30
                                                                        _ = v31
                                                                    }
                                                                case "":
                                                                    for fi66 := 0; fi66 < 8; fi66++ {
                                                                        v67 := GenS3{F0: (^(v16))}
                                                                        if len(v1) > 0 {
                                                                            v1[0] = v24
                                                                        }
                                                                        var v68 int = fi29
                                                                        _ = fi66
                                                                        _ = v67
                                                                        _ = v68
                                                                    }
                                                                }
                                                            }
                                                        case "":
                                                            _ = len(v1)
                                                        default:
                                                            if v22 {
                                                                if false {
                                                                    v5 = (v17 / ((32767) | 1))
                                                                    v3 = append(v3, (-(fi29)))
                                                                } else {
                                                                    v18 = 264
                                                                    v16 = (v16 - v16)
                                                                }
                                                            }
                                                        }
                                                    case "foo bar":
                                                        if v22 {
                                                            for fi69 := 0; fi69 < 3; fi69++ {
                                                                if ((0.1 == (1e10 * -2.25)) && ((v5 - (v20 - (v5 | (v17 & (v17 ^ v17))))) >= (v20 / ((-361) | 1)))) {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = v24
                                                                    }
                                                                }
                                                                _ = fi69
                                                            }
                                                            if v22 {
                                                                v70 := GenS4{F0: (false || (127 != int64(v25))), F1: 0, F2: 159}
                                                                _ = v70
                                                            }
                                                            v24 = (v24 + "foo bar")
                                                        }
                                                    default:
                                                        if v22 {
                                                            v71 := GenS1{F0: (fi29 + (fi29 + fi29))}
                                                            _ = v71
                                                        } else {
                                                            var v72 string = v24
                                                            v73 := GenS5{F0: 92, F1: v4}
                                                            for fi74 := 0; fi74 < 20; fi74++ {
                                                                fmt.Println(v5, fi74)
                                                                _ = fi74
                                                            }
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                    }
                                                    for fi75 := 0; fi75 < 18; fi75++ {
                                                        v22 = (v24 != "foo bar")
                                                        _ = fi75
                                                    }
                                                    _ = fi29
                                                }
                                            }
                                        }
                                        for fi76 := 0; fi76 < 14; fi76++ {
                                            fmt.Println(v25, v16)
                                            _ = fi76
                                        }
                                    }
                                    v77 := GenS1{F0: -9223372036854775808}
                                    v78 := GenS6{F0: -9223372036854775808}
                                    _ = v77
                                    _ = v78
                                }
                                var v79 uint16 = uint16(v2)
                                _ = v25
                                _ = v79
                            }
                            _ = v23
                            _ = v24
                        }
                        if (v22 && false) {
                            v3 = append(v3, 528)
                        }
                    } else {
                        var v80 uint16 = 184
                        _ = v80
                    }
                    for fi81 := 0; fi81 < 11; fi81++ {
                        for fi82 := 0; fi82 < 3; fi82++ {
                            var v83 int = fi81
                            var v84 int32 = 1
                            _ = fi82
                            _ = v83
                            _ = v84
                        }
                        v85 := GenS7{F0: 101, F1: fi81, F2: v4}
                        _ = fi81
                        _ = v85
                    }
                }
                if (32767 != 248) {
                    fmt.Println(v5, v22, v2)
                    var v86 uint = v2
                    if v22 {
                        v2 -= (v86 * v86)
                        if ((-128 >= -112) && ((!v22) || v22)) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v87 int8 = 1
                            v2 -= v2
                            _ = v87
                        }
                        var v88 uint32 = 4294967295
                        _ = v88
                    }
                    _ = v86
                } else {
                    v3 = append(v3, 150)
                }
            }
            v89 := GenS3{F0: (v16 % ((9) | 1))}
            _ = v22
            _ = v89
        } else {
            for fi90 := 0; fi90 < 6; fi90++ {
                v5 = v5
                v5 += (v17 + (^(v20)))
                _ = fi90
            }
            v4 += v4
        }
        _ = v21
    } else {
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        v91 := GenS2{F0: 2, F1: 0}
        var v92 uint8 = 128
        _ = v91
        _ = v92
    }
    var v93 string = ""
    if false {
        var v94 int = -478
        fmt.Println(v93, v4)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v94
    }
    for fi95 := 0; fi95 < 19; fi95++ {
        v96 := GenS1{F0: fi95}
        _ = fi95
        _ = v96
    }
    if (255 == 127) {
        v93 = (v93 + ((v93 + (((v93 + v93) + ("hello" + "hello")) + "a")) + "hello"))
        if (!(false || true)) {
            if len(v1) > 0 {
                v1[0] = (v93 + v93)
            }
            var v97 uint32 = 42
            _ = v97
        } else {
            var v98 string = "a"
            v16 = genFn0(512, v98, (v18 & (v18 >> (((v2 % ((v2) | 1))) & 63))))
            _ = v98
        }
    } else {
        if true {
            if ((v20 ^ -32768) == (v20 % ((v17) | 1))) {
                var v99 uint = 42
                _ = v99
            }
        }
        switch v93 {
        case "a":
            var v100 float32 = (1.5 + float32(1.5))
            _ = v100
        default:
            if (!false) {
                switch v93 {
                case "\n\t":
                    var v101 int16 = (v17 | 32767)
                    _ = v101
                case "a":
                    v102 := GenS4{F0: (214 != uint8(v18)), F1: 1, F2: v2}
                    _ = v102
                case "hello":
                    switch -977 {
                    case -470:
                        var v103 uint32 = v18
                        _ = v103
                    case 230:
                        var v104 uint8 = 173
                        _ = v104
                    }
                }
            } else {
                switch int(v20) {
                case -9223372036854775808:
                    v105 := GenS6{F0: 4294967295}
                    _ = v105
                case 32768:
                    v106 := GenS6{F0: 42}
                    _ = v106
                case -31:
                    v4 = v4
                default:
                    v107 := GenS2{F0: 42, F1: 1}
                    _ = v107
                }
            }
        }
    }
    v17 -= v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v93
}
