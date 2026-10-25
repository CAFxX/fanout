package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint }
type GenS3 struct { F0 uint16; F1 uint64; F2 string }
type GenS4 struct { F0 int32; F1 uint16; F2 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 uint; F1 int16; F2 uint64 }
type GenS8 struct { F0 uint16; F1 int64 }
func genFn0(p0 int8, p1 int8, p2 int16) int64 {
    v6 := GenS1{F0: p1}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 int = 128
    var v2 uint8 = 42
    var v3 uint16 = 252
    var v4 int8 = 127
    var v5 uint64 = 861
    var v7 float32 = float32(-2.25)
    var v8 uint16 = uint16(v2)
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            if false {
                fmt.Println(v2, v2)
                var v11 float32 = 3.14159
                _ = v11
            } else {
                v12 := GenS1{F0: 0}
                var v13 int32 = 0
                _ = v12
                _ = v13
            }
            if true {
                if (!((v3 + uint16(v2)) != (^(v3)))) {
                    v14 := GenS2{F0: 1, F1: (v5 | 0), F2: 4294967295}
                    v15 := []int{fi10}
                    if (("" != ("hello" + (("\n\t" + ("" + "foo bar")) + ("a" + ("\n\t" + ("a" + "hello")))))) || true) {
                        v7 = ((0.0 - v7) * (v7 + 0.1))
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        if ((65535 < 342) && true) {
                            v3 -= 255
                        } else {
                            _ = len(v15)
                            var v16 float32 = 3.14159
                            var v17 float32 = v7
                            switch "hello" {
                            case "":
                                switch 128 {
                                case -993:
                                    if false {
                                        var v18 float64 = 1.5
                                        fmt.Println(v5, v3, fi10)
                                        var v19 bool = false
                                        _ = v18
                                        _ = v19
                                    } else {
                                        v8 = 0
                                        var v20 int64 = genFn0(42, v4, 32767)
                                        _ = v20
                                    }
                                default:
                                    if true {
                                        var v21 int32 = 0
                                        v21 += (-(v21))
                                        _ = v21
                                    } else {
                                        v16 = 0.0
                                    }
                                }
                            case "hello":
                                v22 := GenS2{F0: -1, F1: v5, F2: 65535}
                                _ = v22
                            case "\n\t":
                                var v23 float32 = v16
                                _ = v23
                            default:
                                fmt.Println(v3, v2)
                            }
                            var v24 int16 = 32767
                            if true {
                                var v25 uint16 = 897
                                _ = v25
                            }
                            v26 := GenS3{F0: v8, F1: (v5 >> ((796) & 63)), F2: "foo bar"}
                            _ = v16
                            _ = v17
                            _ = v24
                            _ = v26
                        }
                    } else {
                        if (uint8(v5) >= v2) {
                            v15 = append(v15, (fi10 % ((9223372036854775807) | 1)))
                            switch ("" + ("foo bar" + "\n\t")) {
                            case "a":
                                v27 := GenS2{F0: -9223372036854775808, F1: (v5 * (v5 / ((v5) | 1))), F2: 2}
                                _ = v27
                            case "foo bar":
                                var v28 int64 = 0
                                _ = v28
                            case "":
                                v29 := GenS4{F0: -866, F1: v3, F2: 66}
                                _ = v29
                            }
                        }
                        v30 := GenS4{F0: -1, F1: (v8 << ((386) & 63)), F2: -2147483648}
                        var v31 uint32 = 571
                        _ = v30
                        _ = v31
                    }
                    _ = v14
                    _ = v15
                }
            }
            var v32 uint64 = (v5 / (((^(v5))) | 1))
            var v33 uint16 = uint16(v2)
            _ = fi10
            _ = v32
            _ = v33
        }
        v3 = 42
        var v34 int16 = -32768
        _ = fi9
        _ = v34
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        var v36 uint = 299
        switch (fi35 & v1) {
        case -286:
            if false {
                for fi37 := 0; fi37 < 10; fi37++ {
                    if true {
                        switch ("a" + "") {
                        case "\n\t":
                            v36 += v36
                        case "hello":
                            if true {
                                if (42 >= (v4 | (v4 + v4))) {
                                    switch "a" {
                                    case "foo bar":
                                        fmt.Println(fi37)
                                    case "\n\t":
                                        if true {
                                            var v38 string = "hello"
                                            if false {
                                                var v39 int16 = 676
                                                _ = v39
                                            }
                                            var v40 int16 = 812
                                            var v41 int16 = 0
                                            _ = v38
                                            _ = v40
                                            _ = v41
                                        } else {
                                            v42 := GenS1{F0: 100}
                                            var v43 int64 = -99
                                            _ = v42
                                            _ = v43
                                        }
                                    default:
                                        var v44 int32 = -844
                                        _ = v44
                                    }
                                }
                                var v45 uint = v36
                                _ = v45
                            } else {
                                var v46 float32 = v7
                                v47 := GenS4{F0: 1, F1: v3, F2: 320}
                                switch "" {
                                case "hello":
                                    v48 := GenS5{F0: 0}
                                    _ = v48
                                case "foo bar":
                                    v49 := []int{v1, (^(v1)), -658, (fi37 / (((v1 ^ (v1 & v1))) | 1))}
                                    _ = v49
                                default:
                                    var v50 int = v1
                                    _ = v50
                                }
                                _ = v46
                                _ = v47
                            }
                        case "":
                            if false {
                                var v51 uint32 = 4294967295
                                v3 += 255
                                if false {
                                    v52 := []int{(-(fi37))}
                                    if len(v52) > 0 {
                                        v52[0] = (fi37 & (v1 >> ((v36) & 63)))
                                    }
                                    if false {
                                        v52 = append(v52, (fi37 >> (((^(v36))) & 63)))
                                    }
                                    _ = v52
                                } else {
                                    var v53 uint8 = v2
                                    switch fi37 {
                                    case 42:
                                        v54 := []int{(-(fi37))}
                                        v51 = (v51 >> ((v36) & 63))
                                        var v55 uint64 = (^(v5))
                                        _ = v54
                                        _ = v55
                                    case 4294967295:
                                        if ((v8 * v3) < v8) {
                                            fmt.Println(v1, v3, v8)
                                            fmt.Println(v2, v53, v4)
                                        }
                                        fmt.Println(v51, fi37, fi37)
                                    }
                                    _ = v53
                                }
                                if true {
                                    var v56 string = "foo bar"
                                    v56 = v56
                                    _ = v56
                                }
                                v57 := []int{(fi35 + -755), 42}
                                _ = v51
                                _ = v57
                            } else {
                                var v58 uint16 = (v8 + 42)
                                _ = v58
                            }
                        }
                    } else {
                        fmt.Println(v2, v8)
                        fmt.Println(v4, fi35)
                        if (true && true) {
                            var v59 string = ("hello" + "\n\t")
                            v60 := GenS4{F0: -575, F1: (v8 - v8), F2: 42}
                            fmt.Println(v59, fi35)
                            _ = v59
                            _ = v60
                        } else {
                            var v61 float64 = 0.0
                            v8 = v8
                            _ = v61
                        }
                    }
                    if (2 == 871) {
                        v62 := GenS6{F0: -897, F1: 2}
                        if ((42 <= int8(v5)) || (uint64(v36) >= v5)) {
                            var v63 int64 = 9223372036854775807
                            fmt.Println(v4)
                            _ = v63
                        } else {
                            v2 += v2
                            fmt.Println(v3, v8)
                            v64 := []int{1, fi35}
                            _ = v64
                        }
                        _ = v62
                    } else {
                        if (true || (!(v2 < (v2 * uint8(v36))))) {
                            var v65 int64 = 283
                            if (((^(v3)) < (v8 >> (((v36 & (v36 - 255))) & 63))) || true) {
                                v2 = v2
                            } else {
                                if true {
                                    v66 := []int{fi37, 32768, fi35}
                                    _ = v66
                                } else {
                                    v67 := []int{(fi35 | v1)}
                                    var v68 float64 = (3.14159 / 1e10)
                                    fmt.Println(fi37, fi35, v7)
                                    _ = v67
                                    _ = v68
                                }
                                v69 := []int{}
                                _ = v69
                            }
                            v70 := []int{v1, fi35, -112}
                            _ = v65
                            _ = v70
                        } else {
                            var v71 string = "a"
                            v5 = (^(v5))
                            v72 := []int{fi37, (fi37 / (((v1 * v1)) | 1))}
                            fmt.Println(v3, v7, fi37)
                            _ = v71
                            _ = v72
                        }
                    }
                    v1 = 9223372036854775807
                    v73 := []int{(v1 | 32768), v1}
                    v74 := GenS6{F0: -9223372036854775808, F1: (^(v2))}
                    if false {
                        var v75 int = 128
                        v76 := GenS5{F0: int32(v2)}
                        _ = v75
                        _ = v76
                    }
                    v5 -= v5
                    v73 = append(v73, (fi35 - (^(fi37))))
                    _ = fi37
                    _ = v73
                    _ = v74
                }
            }
        case 42:
            var v77 int32 = -2147483648
            _ = v77
        }
        _ = fi35
        _ = v36
    }
    v78 := GenS4{F0: -425, F1: uint16(v2), F2: -2147483648}
    v79 := []int{}
    var v80 int16 = 32767
    v79 = append(v79, v1)
    var v81 uint = 255
    v79 = append(v79, 42)
    v7 += 1e10
    v8 -= v3
    v79 = append(v79, 512)
    fmt.Println(v5)
    v79 = append(v79, v1)
    for fi82 := 0; fi82 < 16; fi82++ {
        for fi83 := 0; fi83 < 5; fi83++ {
            fmt.Println(v4, v81, v8)
            if (((v3 << ((v81) & 63)) < (v3 * 32767)) || false) {
                v5 = (^(v5))
            }
            switch (v1 | fi82) {
            case 8:
                v84 := GenS7{F0: (v81 + (v81 - uint(v8))), F1: (v80 % (((v80 % (((v80 | 964)) | 1))) | 1)), F2: (v5 >> (((v81 - uint(v5))) & 63))}
                _ = v84
            case -656:
                if true {
                    var v85 float32 = ((1e10 * 1e10) + v7)
                    fmt.Println(v8, v2, v2)
                    _ = v85
                } else {
                    v3 += 0
                    if (v5 < (v5 / ((4294967295) | 1))) {
                        fmt.Println(v80, v80)
                    } else {
                        if true {
                            if (9223372036854775807 >= v1) {
                                _ = len(v79)
                                var v86 int8 = (^(v4))
                                _ = v86
                            } else {
                                v87 := GenS8{F0: v8, F1: 0}
                                _ = v87
                            }
                        } else {
                            v1 += (fi83 - v1)
                            v5 += (v5 + (v5 + 42))
                            if (-971 < (v80 + -1)) {
                                var v88 uint16 = v8
                                _ = v88
                            } else {
                                v81 += v81
                            }
                        }
                        if len(v79) >= 2 {
                            v79 = v79[1:len(v79)]
                        }
                    }
                    var v89 string = "hello"
                    _ = v89
                }
            }
            var v90 string = "a"
            v7 = ((((v7 * (-2.25 / ((1e10 * (float32(0.1) / v7)) / (0.1 * v7)))) - v7) * float32(1.5)) / 1e10)
            v2 = (v2 ^ 128)
            _ = fi83
            _ = v90
        }
        switch "a" {
        case "foo bar":
            if ((!false) && (-1 >= 9223372036854775807)) {
                if len(v79) >= 2 {
                    v79 = v79[1:len(v79)]
                }
                fmt.Println(v4)
                if (850 >= 235) {
                    if false {
                        var v91 string = ("" + ("" + ("hello" + "\n\t")))
                        _ = v91
                    } else {
                        fmt.Println(v8)
                        v3 = v8
                        for fi92 := 0; fi92 < 4; fi92++ {
                            var v93 int8 = 100
                            var v94 uint8 = (v2 ^ v2)
                            var v95 float64 = ((0.1 + 0.0) * (((1.5 * 0.1) + 1.5) / ((((1e10 * 3.14159) + (-2.25 / 1e10)) + (0.0 + (1.5 / 1.5))) / (1.5 - ((1.5 * (1e10 - 1e10)) / 3.14159)))))
                            _ = fi92
                            _ = v93
                            _ = v94
                            _ = v95
                        }
                    }
                } else {
                    fmt.Println(v5)
                    if len(v79) >= 2 {
                        v79 = v79[1:len(v79)]
                    }
                }
            }
        case "hello":
            fmt.Println(v7)
        }
        _ = fi82
    }
    for fi96 := 0; fi96 < 1; fi96++ {
        v97 := GenS4{F0: -2147483648, F1: (v8 / ((v8) | 1)), F2: -331}
        _ = fi96
        _ = v97
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v79:", len(v79), v79)
    fmt.Println("v8:", v8)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v78
    _ = v79
    _ = v8
    _ = v80
    _ = v81
}
