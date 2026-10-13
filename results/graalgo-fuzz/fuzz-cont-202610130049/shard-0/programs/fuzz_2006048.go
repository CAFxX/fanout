package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int8 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint; F1 int64; F2 int16 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int32; F2 int64 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint16; F1 string; F2 int }
type GenS11 struct { F0 int; F1 int8 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int8; F1 int64 }
type GenS14 struct { F0 uint8 }
func genFn0(p0 uint8, p1 uint16) int16 {
    if true {
        p0 -= p0
        v5 := []int{42, 2147483647}
        _ = v5
    }
    if false {
        switch ("\n\t" + "hello") {
        case "\n\t":
            if ((708 > uint64(p1)) || true) {
                switch "a" {
                case "":
                    var v6 uint32 = 2147483648
                    _ = v6
                case "\n\t":
                    switch "\n\t" {
                    case "":
                        if (42 < 42) {
                            var v7 uint = 2
                            _ = v7
                        }
                    default:
                        fmt.Println(p1)
                    }
                case "hello":
                    fmt.Println(p0)
                default:
                    if (908 > 42) {
                        if false {
                            v8 := []int{32768, 0, -2147483648}
                            _ = v8
                        } else {
                            for fi9 := 0; fi9 < 11; fi9++ {
                                v10 := GenS1{F0: 0, F1: -1, F2: 88}
                                var v11 int32 = -2147483648
                                _ = fi9
                                _ = v10
                                _ = v11
                            }
                            if false {
                                switch "foo bar" {
                                case "a":
                                    var v12 float32 = 0.1
                                    _ = v12
                                case "hello":
                                    p0 += (p0 % ((p0) | 1))
                                }
                                v13 := GenS2{F0: 65535, F1: 42}
                                switch 1 {
                                case -390:
                                    if ((42 <= 2147483647) || false) {
                                        p0 += p0
                                    }
                                case 32768:
                                    var v14 uint16 = (p1 - 42)
                                    _ = v14
                                }
                                _ = v13
                            } else {
                                var v15 string = (("foo bar" + (("foo bar" + "") + "")) + (("\n\t" + "hello") + "\n\t"))
                                _ = v15
                            }
                        }
                        var v16 int64 = 0
                        fmt.Println(p1, v16, v16)
                        _ = v16
                    }
                }
                fmt.Println(p0)
                var v17 uint16 = p1
                _ = v17
            } else {
                for fi18 := 0; fi18 < 18; fi18++ {
                    if (255 > 673) {
                        var v19 int16 = 0
                        v20 := []int{fi18, fi18, fi18, fi18}
                        _ = v19
                        _ = v20
                    }
                    _ = fi18
                }
                if true {
                    v21 := GenS3{F0: 42, F1: -775, F2: 0}
                    if ((p1 >= p1) && true) {
                        v22 := GenS4{F0: 1, F1: 32767, F2: (p0 % (((p0 / (((p0 & p0)) | 1))) | 1))}
                        if (false && false) {
                            p0 = p0
                        }
                        _ = v22
                    } else {
                        p1 = (p1 + (p1 ^ 65535))
                        var v23 uint8 = p0
                        v24 := []int{}
                        _ = v23
                        _ = v24
                    }
                    v25 := GenS5{F0: (p1 + 32767), F1: -1, F2: 0}
                    _ = v21
                    _ = v25
                } else {
                    var v26 int16 = -901
                    _ = v26
                }
                for fi27 := 0; fi27 < 2; fi27++ {
                    v28 := GenS6{F0: 0, F1: p0}
                    var v29 uint32 = 65535
                    v29 += (v29 ^ (v29 >> ((42) & 63)))
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
        case "a":
            for fi30 := 0; fi30 < 8; fi30++ {
                for fi31 := 0; fi31 < 6; fi31++ {
                    p1 -= (p1 / ((p1) | 1))
                    v32 := []int{fi31}
                    p1 = p1
                    _ = fi31
                    _ = v32
                }
                var v33 uint64 = 521
                var v34 float64 = (1.5 * ((float64(0.1) * 1.5) / 1.0))
                _ = fi30
                _ = v33
                _ = v34
            }
        default:
            v35 := GenS4{F0: -140, F1: (^(p1)), F2: (p0 << ((65535) & 63))}
            _ = v35
        }
        p1 -= p1
    }
    p1 -= p1
    p0 += 106
    _ = p0
    _ = p1
    return 32767
}
func genFn1(p0 int64, p1 string, p2 string) string {
    fmt.Println(p1)
    p1 = ((p1 + "") + "foo bar")
    v36 := GenS7{F0: 398, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v36
    return "hello"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = ((1.5 * (0.0 - 1.5)) / ((float64((0.0 * 0.1)) - 0.1) - ((0.1 + (3.14159 * float64(-2.25))) * (1e10 + (1.5 + (1e10 + 0.0))))))
    v3 := make(map[int]string)
    var v4 string = ("foo bar" + "\n\t")
    var v37 int = 577
    fmt.Println(v2, v37, v2)
    var v38 uint8 = 2
    var v39 bool = (4294967295 == 42)
    var v40 int16 = 42
    var v41 string = "foo bar"
    v42 := GenS8{F0: -492}
    switch v37 {
    case -9223372036854775808:
        var v43 int64 = 0
        _ = v43
    case 32768:
        if (2147483647 < -1) {
            var v44 uint8 = 255
            _ = v44
        } else {
            for fi45 := 0; fi45 < 2; fi45++ {
                var v46 bool = true
                _ = fi45
                _ = v46
            }
        }
    case 32767:
        var v47 int32 = 42
        _ = v47
    default:
        v48 := GenS9{F0: v38}
        _ = v48
    }
    if (0.1 <= float32(float64(((3.14159 + ((((((float32((v2 * v2)) * (1.5 - (1.5 * (-2.25 - 0.0)))) / (0.1 * 0.1)) + float32(1e10)) - 1.5) * (3.14159 + (0.1 * float32((v2 / -2.25))))) - 1.5)) * 1e10)))) {
        if true {
            var v49 uint = 964
            v49 = (v49 ^ 65535)
            if v39 {
                if false {
                    var v50 uint8 = v38
                    _ = len(v3)
                    _ = v50
                }
                v51 := GenS10{F0: 2, F1: ((("foo bar" + v4) + "") + v4), F2: (-(v37))}
                _ = v51
            }
            _ = v49
        }
        switch v37 {
        case 546:
            var v52 uint = 2
            _ = v52
        case -84:
            var v53 int8 = 100
            _ = v53
        case 32767:
            switch (v4 + (((v4 + v41) + ("" + "")) + "")) {
            case "hello":
                v54 := []int{(v37 / ((v37) | 1))}
                _ = v54
            default:
                fmt.Println(v4)
            }
        default:
            if v39 {
                var v55 float64 = 1e10
                for fi56 := 0; fi56 < 6; fi56++ {
                    for fi57 := 0; fi57 < 17; fi57++ {
                        _ = len(v3)
                        if (2 > 738) {
                            switch (^(fi56)) {
                            case 910:
                                v58 := GenS11{F0: 42, F1: 100}
                                _ = v58
                            case 32767:
                                v59 := []int{(fi56 & 128)}
                                _ = v59
                            case 32768:
                                v2 = (1.5 * float64(3.14159))
                            default:
                                if ("" >= v4) {
                                    fmt.Println(v41)
                                    switch fi57 {
                                    case -1:
                                        if (v38 == (v38 | v38)) {
                                            fmt.Println(v40, v41, v41)
                                        } else {
                                            v40 -= (v40 << ((532) & 63))
                                            fmt.Println(v40, v2)
                                            v60 := GenS4{F0: (v40 << ((743) & 63)), F1: 42, F2: v38}
                                            _ = v60
                                        }
                                    case -834:
                                        v61 := GenS12{F0: -484}
                                        _ = v61
                                    case 9223372036854775807:
                                        var v62 int64 = -9223372036854775808
                                        _ = v62
                                    }
                                } else {
                                    fmt.Println(v4, v40, v38)
                                }
                                fmt.Println(v37, fi57, v41)
                            }
                            var v63 uint16 = 65535
                            var v64 float64 = 3.14159
                            _ = v63
                            _ = v64
                        }
                        _ = fi57
                    }
                    v38 = v38
                    v40 = 1
                    _ = fi56
                }
                var v65 uint32 = 2147483647
                _ = v55
                _ = v65
            }
        }
        v3[-904] = (v41 + ("foo bar" + v41))
    } else {
        v66 := GenS9{F0: v38}
        var v67 uint32 = 2147483647
        _ = v66
        _ = v67
    }
    var v68 uint32 = 302
    fmt.Println(v39)
    fmt.Println(v4, v39)
    fmt.Println(v39, v41)
    fmt.Println(v68, v41, v38)
    v69 := []int{v37, (-(v37)), v37, -1}
    v40 -= 0
    if v39 {
        fmt.Println(v40, v38)
    } else {
        if (1 == 55) {
            fmt.Println(v38, v68)
            if (!(112 < 0)) {
                switch v37 {
                case 32767:
                    var v70 uint32 = uint32(v38)
                    _ = v70
                case 2147483647:
                    if v39 {
                        fmt.Println(v2, v39)
                        var v71 int32 = 0
                        _ = v71
                    }
                case 951:
                    var v72 float64 = ((((v2 + float64((0.0 / -2.25))) - (v2 + (float64((((-2.25 - 0.0) - (0.1 + 1e10)) / 1.0)) - (v2 - v2)))) + -2.25) + ((0.0 + v2) * v2))
                    _ = v72
                default:
                    var v73 int16 = (v40 >> ((445) & 63))
                    _ = v73
                }
                v74 := v1[v68]
                if len(v69) >= 2 {
                    v69 = v69[1:len(v69)]
                }
                _ = v74
            } else {
                switch v37 {
                case 127:
                    var v75 uint16 = 1
                    _ = v75
                case -873:
                    var v76 uint64 = 938
                    _ = v76
                case -1:
                    fmt.Println(v68, v37)
                default:
                    var v77 int64 = 9223372036854775807
                    _ = v77
                }
                var v78 int8 = 100
                v79 := GenS2{F0: 1, F1: 1}
                _ = v78
                _ = v79
            }
            if v39 {
                var v80 uint64 = 2
                for fi81 := 0; fi81 < 20; fi81++ {
                    if (4294967295 < v80) {
                        _ = len(v1)
                    }
                    v4 = "\n\t"
                    v80 = uint64(v38)
                    _ = fi81
                }
                if false {
                    for fi82 := 0; fi82 < 18; fi82++ {
                        for fi83 := 0; fi83 < 6; fi83++ {
                            switch (fi82 & fi83) {
                            case 450:
                                if false {
                                    if (127 <= 42) {
                                        var v84 uint = 65535
                                        _ = v84
                                    }
                                    switch (v37 % (((v37 | fi82)) | 1)) {
                                    case 4294967295:
                                        if v39 {
                                            var v85 uint64 = 42
                                            var v86 bool = (272 <= -2147483648)
                                            v68 += v68
                                            _ = v85
                                            _ = v86
                                        } else {
                                            if (!v39) {
                                                fmt.Println(v4, fi82)
                                                if (!((v68 ^ v68) < (v68 * v68))) {
                                                    fmt.Println(fi83, v41, v39)
                                                    if len(v69) >= 2 {
                                                        v69 = v69[1:len(v69)]
                                                    }
                                                    if (v41 < "hello") {
                                                        v41 = (v4 + "a")
                                                        var v87 uint16 = 0
                                                        var v88 string = v41
                                                        var v89 int32 = -1
                                                        _ = v87
                                                        _ = v88
                                                        _ = v89
                                                    }
                                                } else {
                                                    var v90 int32 = -2147483648
                                                    _ = v90
                                                }
                                                v41 = v41
                                                v91 := GenS13{F0: 1, F1: 0}
                                                _ = v91
                                            }
                                        }
                                    case -638:
                                        fmt.Println(v38, fi82)
                                    case -1:
                                        fmt.Println(fi83, v38)
                                    default:
                                        v37 = (fi83 & 4294967295)
                                    }
                                    v38 -= v38
                                } else {
                                    var v92 uint64 = 0
                                    v2 -= float64(0.0)
                                    v2 = (v2 / v2)
                                    _ = v92
                                }
                            case 32768:
                                fmt.Println(v41, v2, v80)
                            }
                            var v93 int = v37
                            if true {
                                var v94 uint8 = (v38 - uint8(v80))
                                if ((float32(float64(((float32(float64(((-2.25 + -2.25) + float32((v2 * ((v2 - 1e10) + (1e10 - 0.1))))))) + (1e10 + 1e10)) - float32(((v2 * float64(1e10)) - v2))))) == 1.5) || ((fi82 % ((fi83) | 1)) >= (v93 * (v37 & -10)))) {
                                    _ = len(v1)
                                    fmt.Println(v39, v2, v4)
                                }
                                _ = v94
                            }
                            _ = fi83
                            _ = v93
                        }
                        fmt.Println(v41)
                        _ = fi82
                    }
                    if (107 == -128) {
                        if v39 {
                            v95 := GenS14{F0: 242}
                            _ = v95
                        }
                    } else {
                        var v96 int32 = 42
                        var v97 uint32 = (v68 - v68)
                        _ = v96
                        _ = v97
                    }
                }
                _ = v80
            } else {
                _ = len(v3)
                var v98 int64 = -1
                _ = v98
            }
        } else {
            if v39 {
                v1[v68] = v68
                v4 = ((v41 + ("a" + v4)) + v4)
            }
        }
    }
    fmt.Println(v41, v68, v39)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v68:", v68)
    fmt.Println("slicelen v69:", len(v69), v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v68
    _ = v69
}
