package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 int16; F2 uint }
type GenS6 struct { F0 string; F1 uint64; F2 int32 }
type GenS7 struct { F0 int; F1 uint8; F2 uint8 }
type GenS8 struct { F0 int64 }
func genFn0(p0 int8) uint8 {
    fmt.Println(p0)
    p0 = (^(p0))
    _ = p0
    return 1
}
func genFn1() int16 {
    var v4 uint = 0
    v4 += (v4 << (((v4 * v4)) & 63))
    _ = v4
    return -961
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    var v5 int32 = -1
    var v6 uint16 = 42
    var v7 int16 = 215
    v7 += -333
    if ((308 == 0) && false) {
        var v8 string = "foo bar"
        _ = v8
    } else {
        for fi9 := 0; fi9 < 18; fi9++ {
            v5 = v5
            v10 := []int{42, fi9, -9223372036854775808, (fi9 + (fi9 - 951))}
            v5 = v5
            _ = fi9
            _ = v10
        }
        var v11 int32 = (v5 / ((v5) | 1))
        _ = v11
    }
    var v12 uint = 528
    v12 += 42
    var v13 int = 32768
    switch "hello" {
    case "a":
        switch "" {
        case "\n\t":
            if true {
                if false {
                    var v14 int = 32768
                    if (!false) {
                        fmt.Println(v1, v12)
                        if false {
                            fmt.Println(v13, v13)
                            if (v6 != (v6 - (v6 / ((155) | 1)))) {
                                v3[1] = "\n\t"
                                v14 = 127
                            } else {
                                v15 := GenS1{F0: 1}
                                for fi16 := 0; fi16 < 5; fi16++ {
                                    switch v14 {
                                    case 679:
                                        for fi17 := 0; fi17 < 9; fi17++ {
                                            v18 := GenS2{F0: (("hello" + "hello") + "\n\t")}
                                            _ = fi17
                                            _ = v18
                                        }
                                    }
                                    v19 := GenS2{F0: ("\n\t" + "\n\t")}
                                    _ = fi16
                                    _ = v19
                                }
                                _ = v15
                            }
                            var v20 int32 = v5
                            _ = v20
                        } else {
                            switch ("foo bar" + "\n\t") {
                            case "hello":
                                _ = len(v3)
                            case "foo bar":
                                if true {
                                    var v21 string = "hello"
                                    var v22 uint = 255
                                    v12 -= (^(v12))
                                    _ = v21
                                    _ = v22
                                }
                            case "\n\t":
                                var v23 uint = (v12 ^ v12)
                                _ = v23
                            default:
                                v14 -= (^(v13))
                            }
                        }
                    } else {
                        v24 := []int{(v14 / ((32768) | 1)), (^(v14)), (v14 & v13), v13}
                        _ = v24
                    }
                    _ = v14
                } else {
                    v5 += (^(v5))
                    for fi25 := 0; fi25 < 13; fi25++ {
                        fmt.Println(v13, v1)
                        v5 -= (v5 % (((v5 % (((v5 + v5)) | 1))) | 1))
                        fmt.Println(v12, v1)
                        _ = fi25
                    }
                    v5 = (-(v5))
                }
            } else {
                v13 -= 32767
                var v26 float64 = v1
                _ = v26
            }
        }
    case "\n\t":
        var v27 uint16 = (^(v6))
        _ = v27
    }
    var v28 int16 = 1000
    for fi29 := 0; fi29 < 16; fi29++ {
        if false {
            for fi30 := 0; fi30 < 16; fi30++ {
                v12 -= (^(v12))
                _ = fi30
            }
            var v31 float64 = v1
            _ = v31
        }
        v32 := GenS3{F0: 255}
        if false {
            var v33 uint = v12
            if false {
                switch "a" {
                case "foo bar":
                    v1 = 0.1
                case "\n\t":
                    var v34 uint8 = 42
                    _ = v34
                case "a":
                    switch "hello" {
                    case "foo bar":
                        v28 -= genFn1()
                    case "\n\t":
                        var v35 int = v13
                        _ = v35
                    case "hello":
                        switch (v13 / (((fi29 / ((v13) | 1))) | 1)) {
                        case 410:
                            v36 := GenS2{F0: "hello"}
                            _ = v36
                        default:
                            if (!false) {
                                var v37 uint16 = v6
                                fmt.Println(v5, v12)
                                switch -985 {
                                case -374:
                                    v38 := []int{}
                                    _ = v38
                                case -1:
                                    if false {
                                        v39 := GenS2{F0: ("hello" + (("foo bar" + "foo bar") + ""))}
                                        for fi40 := 0; fi40 < 15; fi40++ {
                                            v7 = (v7 + (v7 | genFn1()))
                                            _ = fi40
                                        }
                                        v37 += 42
                                        _ = v39
                                    }
                                case 42:
                                    switch fi29 {
                                    case 42:
                                        if (!(((("hello" + "foo bar") <= "") || false) && true)) {
                                            switch v13 {
                                            case 42:
                                                var v41 uint16 = (v6 + (v6 << (((^(v12))) & 63)))
                                                _ = v41
                                            default:
                                                if false {
                                                    var v42 int8 = 127
                                                    _ = v42
                                                }
                                            }
                                            var v43 uint64 = 4294967295
                                            _ = v43
                                        } else {
                                            v6 = (v6 - v37)
                                            fmt.Println(fi29, v7, v7)
                                            fmt.Println(v13, v33)
                                        }
                                    case 32767:
                                        v28 = genFn1()
                                    default:
                                        v3[(v13 + (v13 ^ v13))] = "hello"
                                    }
                                }
                                _ = v37
                            }
                        }
                    }
                default:
                    for fi44 := 0; fi44 < 9; fi44++ {
                        var v45 float64 = v1
                        var v46 string = "a"
                        switch fi29 {
                        case 32767:
                            v6 = v6
                        case 127:
                            switch ("hello" + v46) {
                            case "":
                                v33 = 2
                                fmt.Println(v12, v45)
                            case "foo bar":
                                v47 := GenS2{F0: "hello"}
                                _ = v47
                            case "a":
                                v13 = fi44
                                v46 = v46
                                fmt.Println(v45, v33)
                            }
                        case 2147483647:
                            switch v46 {
                            case "a":
                                if true {
                                    if false {
                                        if false {
                                            fmt.Println(v46, v5)
                                            if false {
                                                var v48 uint64 = 42
                                                switch "a" {
                                                case "a":
                                                    v1 += v45
                                                case "\n\t":
                                                    v2[2] = int(v7)
                                                }
                                                if true {
                                                    fmt.Println(v7, v46, v7)
                                                    v7 += (-(v7))
                                                }
                                                var v49 uint8 = 127
                                                _ = v48
                                                _ = v49
                                            } else {
                                                if true {
                                                    v50 := GenS4{F0: true}
                                                    if (false && (393 == 2)) {
                                                        _ = len(v3)
                                                        switch 4294967295 {
                                                        case 127:
                                                            var v51 int = fi44
                                                            _ = v51
                                                        default:
                                                            switch (v13 * v13) {
                                                            case 32767:
                                                                v12 = v12
                                                            case 977:
                                                                v52 := []int{(fi29 % ((4294967295) | 1)), fi44, fi29, (fi44 % (((v13 ^ -566)) | 1))}
                                                                _ = v52
                                                            case 1:
                                                                var v53 uint32 = 879
                                                                if false {
                                                                    if false {
                                                                        _ = len(v3)
                                                                    } else {
                                                                        v54 := GenS4{F0: false}
                                                                        _ = v54
                                                                    }
                                                                    v45 = v45
                                                                    v46 = ("foo bar" + (((v46 + "foo bar") + v46) + ("a" + "\n\t")))
                                                                    var v55 int32 = (v5 - 42)
                                                                    _ = v55
                                                                }
                                                                _ = v53
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(v45)
                                                    _ = v50
                                                }
                                            }
                                            v46 = ("a" + v46)
                                            switch v46 {
                                            case "\n\t":
                                                var v56 bool = true
                                                _ = v56
                                            default:
                                                fmt.Println(v5)
                                            }
                                        } else {
                                            v7 = genFn1()
                                            v57 := GenS3{F0: genFn0(86)}
                                            var v58 int = (v13 / (((v13 - fi44)) | 1))
                                            _ = v57
                                            _ = v58
                                        }
                                        v33 += 42
                                    }
                                    if ((^(v6)) >= (v6 << ((65535) & 63))) {
                                        v33 = v12
                                    }
                                }
                                var v59 uint = 817
                                _ = v59
                            default:
                                if (false && (1 >= v6)) {
                                    var v60 uint32 = 2
                                    _ = v60
                                }
                            }
                        default:
                            if false {
                                if (false && ((v28 == (v7 - genFn1())) || (!true))) {
                                    if (((51 != 42) && (true && (457 > 932))) && (2147483647 == 2147483647)) {
                                        var v61 uint32 = 2147483648
                                        _ = v61
                                    } else {
                                        var v62 int64 = 1
                                        _ = v62
                                    }
                                }
                            } else {
                                switch "hello" {
                                case "":
                                    switch 886 {
                                    case 32767:
                                        if (true && true) {
                                            var v63 string = v46
                                            _ = v63
                                        } else {
                                            var v64 int8 = int8(v5)
                                            v12 = v33
                                            v65 := GenS2{F0: v46}
                                            _ = v64
                                            _ = v65
                                        }
                                    case 128:
                                        v66 := GenS5{F0: (v12 - (v12 + 2)), F1: (v7 / (((v7 << (((v33 + (^(v12)))) & 63))) | 1)), F2: (v12 | (v33 / ((v33) | 1)))}
                                        _ = v66
                                    default:
                                        var v67 bool = (3.14159 > 3.14159)
                                        _ = v67
                                    }
                                    if true {
                                        var v68 uint32 = 2
                                        fmt.Println(v7, fi44, v13)
                                        _ = v68
                                    }
                                    if false {
                                        switch v13 {
                                        case -466:
                                            v7 -= -1
                                        case 238:
                                            v69 := []int{fi29, fi44, v13}
                                            _ = v69
                                        case 1:
                                            fmt.Println(v7, v13)
                                        default:
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v70 int32 = 0
                                                _ = v70
                                            case "":
                                                v33 = (^(v33))
                                            case "a":
                                                if true {
                                                    v71 := GenS2{F0: ("foo bar" + v46)}
                                                    if false {
                                                        if false {
                                                            if (42 > 100) {
                                                                switch v46 {
                                                                case "foo bar":
                                                                    if (((-2.25 + 0.1) * 3.14159) <= (-2.25 * (0.0 - (3.14159 + 0.1)))) {
                                                                        v72 := GenS6{F0: "\n\t", F1: 18446744073709551615, F2: 2147483647}
                                                                        v28 = (v28 % ((945) | 1))
                                                                        _ = v72
                                                                    }
                                                                case "\n\t":
                                                                    fmt.Println(fi44, v46, v12)
                                                                }
                                                                switch fi44 {
                                                                case 9223372036854775807:
                                                                    var v73 int = (v13 + int(v6))
                                                                    _ = v73
                                                                case 820:
                                                                    var v74 float32 = 0.0
                                                                    _ = v74
                                                                case 1:
                                                                    var v75 float64 = ((v45 * (v45 * v45)) / (v45 / float64(1e10)))
                                                                    _ = v75
                                                                }
                                                            } else {
                                                                if true {
                                                                    v76 := []int{(fi29 ^ (v13 + fi29)), -125, -6}
                                                                    _ = v76
                                                                }
                                                                v77 := GenS6{F0: (v46 + (("\n\t" + v46) + "a")), F1: 83, F2: (v5 | v5)}
                                                                v33 = (v33 * v33)
                                                                v78 := []int{fi44, -2147483648, (fi44 * -9223372036854775808), -9223372036854775808}
                                                                var v79 uint16 = v6
                                                                _ = v77
                                                                _ = v78
                                                                _ = v79
                                                            }
                                                            v80 := []int{}
                                                            _ = v80
                                                        } else {
                                                            v6 = (v6 | (v6 ^ 718))
                                                            v81 := []int{fi29, (-(v13)), (fi29 % ((fi29) | 1)), fi44}
                                                            if false {
                                                                v82 := GenS5{F0: 329, F1: genFn1(), F2: v12}
                                                                _ = len(v81)
                                                                _ = len(v3)
                                                                fmt.Println(fi29, v13, v33)
                                                                v5 = v5
                                                                _ = v82
                                                            }
                                                            _ = v81
                                                        }
                                                        var v83 int64 = 42
                                                        _ = v83
                                                    } else {
                                                        var v84 uint = (v33 << ((v12) & 63))
                                                        if false {
                                                            var v85 uint16 = 269
                                                            v86 := GenS7{F0: (v13 & 277), F1: 127, F2: 67}
                                                            fmt.Println(v33, v6, v85)
                                                            var v87 uint8 = 199
                                                            if (v85 <= v6) {
                                                                fmt.Println(v84, v45)
                                                                var v88 bool = false
                                                                var v89 int = int(v6)
                                                                _ = v88
                                                                _ = v89
                                                            } else {
                                                                var v90 bool = true
                                                                v84 += 18446744073709551615
                                                                fmt.Println(v90, v87)
                                                                var v91 bool = v90
                                                                _ = v90
                                                                _ = v91
                                                            }
                                                            _ = v85
                                                            _ = v86
                                                            _ = v87
                                                        } else {
                                                            v92 := GenS8{F0: 0}
                                                            _ = v92
                                                        }
                                                        _ = v84
                                                    }
                                                    _ = v71
                                                }
                                            default:
                                                if (-2.25 <= 1e10) {
                                                    v93 := []int{(fi29 % ((fi29) | 1)), (fi29 + (^(v13))), fi29}
                                                    _ = v93
                                                }
                                            }
                                        }
                                        delete(v2, 2147483647)
                                    }
                                }
                            }
                        }
                        if (genFn0(42) >= genFn0(32)) {
                            v7 += (v7 << (((v12 + (v33 / ((v12) | 1)))) & 63))
                            v7 += v7
                        }
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                }
                var v94 int32 = -943
                _ = v94
            }
            _ = v33
        } else {
            v12 = v12
            switch "" {
            case "hello":
                if true {
                    v12 += 601
                    v1 -= float64(float32(v1))
                }
            case "":
                v13 -= v13
            case "a":
                var v95 uint32 = 563
                _ = v95
            }
        }
        _ = fi29
        _ = v32
    }
    var v96 uint = (v12 * (v12 & 4294967295))
    var v97 int32 = (v5 - (v5 % ((v5) | 1)))
    var v98 int8 = -30
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v96
    _ = v97
    _ = v98
}
