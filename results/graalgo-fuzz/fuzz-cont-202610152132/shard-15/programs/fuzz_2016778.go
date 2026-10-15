package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint16; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint8; F1 bool; F2 string }
func genFn0(p0 uint, p1 uint8) int64 {
    v7 := GenS1{F0: 42}
    switch 0 {
    case 127:
        for fi8 := 0; fi8 < 16; fi8++ {
            for fi9 := 0; fi9 < 6; fi9++ {
                switch "hello" {
                case "a":
                    if (2 == 889) {
                        p1 = 42
                        if (false || true) {
                            var v10 bool = (70 != (p1 >> ((p0) & 63)))
                            _ = v10
                        }
                        var v11 string = ""
                        _ = v11
                    } else {
                        var v12 int32 = 0
                        if (4294967295 <= uint64(p1)) {
                            v13 := []int{(^(fi8)), (fi8 * fi9)}
                            _ = v13
                        } else {
                            v14 := GenS1{F0: 42}
                            _ = v14
                        }
                        v15 := GenS1{F0: 1}
                        _ = v12
                        _ = v15
                    }
                case "foo bar":
                    fmt.Println(fi8, fi9)
                case "\n\t":
                    switch "hello" {
                    case "a":
                        if (!(32767 < 2)) {
                            var v16 uint8 = p1
                            var v17 uint = (^(p0))
                            v18 := GenS1{F0: 912}
                            _ = v16
                            _ = v17
                            _ = v18
                        } else {
                            if ((^(fi9)) == (fi9 << (((p0 << ((p0) & 63))) & 63))) {
                                if ("a" < "hello") {
                                    fmt.Println(p0, p0, fi8)
                                }
                            }
                            v19 := []int{fi8, (fi8 & (-(fi8))), fi8}
                            p0 = 500
                            _ = v19
                        }
                        if false {
                            if (true && (!((788 <= int64(fi8)) || true))) {
                                v20 := GenS1{F0: -1}
                                _ = v20
                            } else {
                                p1 = (p1 << ((65535) & 63))
                                var v21 int32 = 1
                                _ = v21
                            }
                            v22 := []int{-579, (fi8 - (fi8 % (((fi9 | fi9)) | 1))), (fi8 % (((fi9 ^ 497)) | 1))}
                            p1 -= (p1 << ((p0) & 63))
                            _ = v22
                        }
                        var v23 uint64 = 0
                        _ = v23
                    case "":
                        v24 := GenS2{F0: true, F1: 0, F2: (!(true || (65535 < 65535)))}
                        _ = v24
                    case "\n\t":
                        if false {
                            p1 = p1
                        } else {
                            if (2147483647 < -396) {
                                v25 := GenS2{F0: false, F1: 65535, F2: true}
                                var v26 uint8 = uint8(p0)
                                _ = v25
                                _ = v26
                            }
                            var v27 int16 = int16(p0)
                            v27 += -1
                            _ = v27
                        }
                    default:
                        p1 = (^(p1))
                    }
                default:
                    if false {
                        fmt.Println(p0)
                        if false {
                            p0 = (p0 + (^(p0)))
                        } else {
                            if true {
                                v28 := []int{491}
                                if false {
                                    if (42 != 127) {
                                        if false {
                                            var v29 int32 = 0
                                            v30 := GenS3{F0: 273}
                                            _ = v29
                                            _ = v30
                                        } else {
                                            switch 577 {
                                            case -9223372036854775808:
                                                if false {
                                                    if len(v28) >= 2 {
                                                        v28 = v28[1:len(v28)]
                                                    }
                                                }
                                            case 935:
                                                if true {
                                                    var v31 string = (("foo bar" + "\n\t") + "hello")
                                                    v28 = append(v28, 42)
                                                    _ = v31
                                                }
                                            default:
                                                var v32 float32 = float32(1.5)
                                                _ = v32
                                            }
                                        }
                                    } else {
                                        switch (^(fi9)) {
                                        case -9223372036854775808:
                                            fmt.Println(fi8, fi8, fi9)
                                        case -739:
                                            v28 = append(v28, 0)
                                        default:
                                            p1 += p1
                                        }
                                        p0 = (p0 % ((p0) | 1))
                                    }
                                    p0 -= p0
                                }
                                var v33 int8 = 42
                                _ = v28
                                _ = v33
                            }
                            v34 := GenS1{F0: 318}
                            var v35 string = ("foo bar" + "\n\t")
                            switch fi8 {
                            case 128:
                                if true {
                                    v36 := []int{}
                                    if len(v36) > 0 {
                                        v36[0] = (fi8 ^ fi8)
                                    }
                                    _ = v36
                                } else {
                                    var v37 uint = (p0 >> (((p0 / ((p0) | 1))) & 63))
                                    v38 := GenS2{F0: (v35 <= v35), F1: 0, F2: true}
                                    _ = v37
                                    _ = v38
                                }
                            case 0:
                                var v39 float32 = -2.25
                                _ = v39
                            }
                            v35 = (("\n\t" + (("" + v35) + (("hello" + ("\n\t" + v35)) + (v35 + (v35 + v35))))) + v35)
                            var v40 float64 = (0.0 * 0.1)
                            var v41 int64 = 1
                            _ = v34
                            _ = v35
                            _ = v40
                            _ = v41
                        }
                        v42 := GenS2{F0: false, F1: 32767, F2: true}
                        _ = v42
                    } else {
                        p0 = (p0 >> (((p0 - p0)) & 63))
                        if (1 >= 428) {
                            var v43 float32 = (float32(0.0) + -2.25)
                            var v44 string = ""
                            p1 = (p1 | p1)
                            fmt.Println(p1)
                            var v45 int = (fi9 - (-(fi9)))
                            _ = v43
                            _ = v44
                            _ = v45
                        } else {
                            v46 := GenS1{F0: -2147483648}
                            v47 := []int{(fi9 & fi8), (fi9 >> ((p0) & 63)), (fi8 << ((65535) & 63)), (fi8 >> ((p0) & 63))}
                            _ = v46
                            _ = v47
                        }
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(fi8)
                        case "":
                            p0 = 255
                        default:
                            if (3.14159 >= 3.14159) {
                                v48 := []int{-1, -2147483648}
                                if len(v48) >= 2 {
                                    v48 = v48[1:len(v48)]
                                }
                                _ = v48
                            }
                        }
                    }
                }
                _ = fi9
            }
            p0 = 4294967295
            _ = fi8
        }
    default:
        v49 := []int{}
        _ = v49
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 uint8, p1 uint16) uint32 {
    switch (("foo bar" + ("hello" + ("\n\t" + (("a" + ("" + ("hello" + "hello"))) + "hello")))) + ("\n\t" + "foo bar")) {
    case "":
        var v50 int = 2147483647
        _ = v50
    default:
        switch -9223372036854775808 {
        case 128:
            v51 := []int{1, -965, -1}
            _ = v51
        case 32767:
            p0 = 70
        }
    }
    fmt.Println(p0, p1)
    v52 := GenS3{F0: 721}
    var v53 int32 = -403
    _ = p0
    _ = p1
    _ = v52
    _ = v53
    return 42
}

func main() {
    v1 := []int{1, -1}
    var v2 uint16 = 547
    var v3 int8 = 127
    v4 := make(map[uint32]int)
    var v5 string = (("foo bar" + "a") + "foo bar")
    v6 := make(map[uint32]int)
    v5 = (v5 + "\n\t")
    var v54 int32 = 1
    v5 = (v5 + v5)
    _ = len(v4)
    if (v2 > (v2 - (v2 % ((v2) | 1)))) {
        v55 := GenS4{F0: -871, F1: 2, F2: v2}
        var v56 int16 = 0
        _ = v55
        _ = v56
    }
    v54 -= v54
    var v57 float64 = 0.0
    var v58 float64 = v57
    switch 32767 {
    case 725:
        if len(v1) > 0 {
            v1[0] = 127
        }
    }
    for fi59 := 0; fi59 < 19; fi59++ {
        for fi60 := 0; fi60 < 13; fi60++ {
            var v61 float32 = 0.0
            if (0 > 4294967295) {
                var v62 uint = 0
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v62
            }
            var v63 uint64 = 2
            v64 := v6[65535]
            var v65 int16 = 744
            _ = fi60
            _ = v61
            _ = v63
            _ = v64
            _ = v65
        }
        _ = fi59
    }
    var v66 bool = true
    if (4294967295 < 1) {
        for fi67 := 0; fi67 < 20; fi67++ {
            for fi68 := 0; fi68 < 12; fi68++ {
                var v69 bool = (!v66)
                _ = fi68
                _ = v69
            }
            _ = fi67
        }
        switch (v5 + ((v5 + v5) + ("hello" + "hello"))) {
        case "foo bar":
            fmt.Println(v2, v2, v58)
        case "":
            if true {
                v54 = -1
                v70 := GenS1{F0: int32(v3)}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v70
            }
        default:
            v66 = (!false)
        }
        if v66 {
            for fi71 := 0; fi71 < 12; fi71++ {
                if (!v66) {
                    v66 = v66
                    var v72 string = ((v5 + ((v5 + ("foo bar" + ((v5 + v5) + v5))) + v5)) + "hello")
                    _ = v72
                } else {
                    v2 = (^(v2))
                }
                _ = fi71
            }
            if len(v1) > 0 {
                v1[0] = 128
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = 1
    }
    if false {
        if v66 {
            if (((((v57 + v57) + v58) + (v57 * 1.5)) / 1.0) >= 1e10) {
                switch -1 {
                case 371:
                    v57 = v57
                case 857:
                    var v73 float64 = v57
                    _ = v73
                case -651:
                    v74 := GenS2{F0: true, F1: 2, F2: (!true)}
                    _ = v74
                }
                var v75 string = ""
                _ = v75
            }
            var v76 bool = ((1.5 * 0.0) != (((((float32((float64(1.5) + v57)) + 1.5) * 1e10) + (1.5 / float32(3.14159))) - ((-2.25 * float32((float64(1.5) - float64(float32(0.0))))) / 1e10)) - (-2.25 + -2.25)))
            _ = v76
        }
        v77 := GenS5{F0: 218, F1: ((!false) || v66)}
        _ = v77
    } else {
        for fi78 := 0; fi78 < 11; fi78++ {
            for fi79 := 0; fi79 < 16; fi79++ {
                v2 += (v2 | (v2 ^ 640))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v80 uint16 = 255
                _ = fi79
                _ = v80
            }
            var v81 uint64 = 265
            _ = fi78
            _ = v81
        }
        fmt.Println(v54)
        if v66 {
            if true {
                fmt.Println(v54, v3, v3)
            }
        }
    }
    fmt.Println(v3, v57)
    v82 := GenS4{F0: -32768, F1: 18446744073709551615, F2: (v2 & v2)}
    if false {
        var v83 uint32 = 2
        v66 = true
        if v66 {
            v84 := GenS6{F0: -558}
            v85 := GenS7{F0: 132, F1: true, F2: v5}
            _ = v84
            _ = v85
        }
        _ = v83
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v57
    _ = v58
    _ = v6
    _ = v66
    _ = v82
}
