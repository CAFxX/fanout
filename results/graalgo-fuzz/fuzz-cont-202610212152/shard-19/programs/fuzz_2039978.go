package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 bool; F1 int; F2 string }
type GenS6 struct { F0 bool; F1 bool; F2 int16 }
type GenS7 struct { F0 int; F1 uint; F2 uint32 }
type GenS8 struct { F0 bool; F1 uint8 }
type GenS9 struct { F0 int64; F1 int; F2 int64 }
func genFn0(p0 int) uint16 {
    var v7 int8 = -64
    p0 = -750
    var v8 float32 = ((1.5 * 1.5) / -2.25)
    switch "hello" {
    case "foo bar":
        if (1 <= -1) {
            var v9 uint32 = 1
            if (-64 == (^(v7))) {
                switch "\n\t" {
                case "":
                    v10 := GenS1{F0: v7, F1: "foo bar"}
                    _ = v10
                case "hello":
                    v11 := []int{(-(p0)), (p0 & (p0 << ((547) & 63)))}
                    _ = v11
                default:
                    v8 = v8
                }
            } else {
                v12 := GenS1{F0: -50, F1: ""}
                v13 := GenS1{F0: 42, F1: "foo bar"}
                v14 := []int{-195}
                _ = v12
                _ = v13
                _ = v14
            }
            var v15 int16 = 42
            _ = v15
            _ = v9
        }
    case "\n\t":
        v16 := []int{}
        _ = v16
    }
    _ = p0
    _ = v7
    _ = v8
    return 255
}
func genFn1() uint16 {
    var v17 uint = 1
    v17 += (v17 - 4294967295)
    var v18 float32 = (float32(3.14159) + 1e10)
    v19 := GenS2{F0: v17}
    v17 += 65535
    _ = v17
    _ = v18
    _ = v19
    return uint16(v17)
}

func main() {
    var v1 bool = (17 == 1)
    v2 := []string{"hello", "", "foo bar", "\n\t"}
    var v3 uint8 = 2
    v4 := make(map[uint32]int)
    var v5 uint64 = 1
    v6 := make(map[uint32]int)
    fmt.Println(v1, v5)
    var v20 int32 = -1
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 uint16 = 0
        for fi23 := 0; fi23 < 1; fi23++ {
            switch "a" {
            case "\n\t":
                v24 := GenS3{F0: 4294967295, F1: (v22 ^ v22)}
                _ = v24
            case "a":
                v1 = v1
            case "":
                if (100 <= 100) {
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    var v25 uint8 = v3
                    switch "a" {
                    case "":
                        if (-900 < 32767) {
                            switch ("" + "hello") {
                            case "a":
                                _ = len(v4)
                            }
                        } else {
                            if len(v2) > 0 {
                                v2[0] = (("" + "\n\t") + ((("a" + ((("a" + "") + ("\n\t" + "\n\t")) + "hello")) + (("foo bar" + "") + "\n\t")) + "hello"))
                            }
                            if v1 {
                                var v26 int16 = 607
                                v27 := GenS3{F0: 2, F1: uint16(v5)}
                                v28 := GenS4{F0: 4, F1: 38}
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                            switch 0 {
                            case -1:
                                v25 = v25
                            case 462:
                                _ = len(v6)
                            default:
                                if v1 {
                                    v1 = (("a" + "foo bar") == (("a" + "") + "hello"))
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v29 := GenS5{F0: false, F1: (^(fi21)), F2: ""}
                                    _ = v29
                                }
                            }
                        }
                    default:
                        v30 := v4[601]
                        _ = v30
                    }
                    _ = v25
                } else {
                    var v31 uint8 = (v3 - (v3 % ((42) | 1)))
                    v31 += (^(v3))
                    _ = v31
                }
            }
            v1 = (2147483647 < 218)
            v32 := GenS2{F0: 4294967295}
            _ = fi23
            _ = v32
        }
        _ = fi21
        _ = v22
    }
    if len(v2) > 0 {
        v2[0] = ("hello" + "hello")
    }
    if (0.0 != 1.5) {
        var v33 int16 = 0
        if (-128 < 100) {
            v34 := v4[1]
            for fi35 := 0; fi35 < 5; fi35++ {
                if (!(v1 || v1)) {
                    if true {
                        v33 = 788
                    } else {
                        for fi36 := 0; fi36 < 19; fi36++ {
                            var v37 uint32 = 1
                            v38 := GenS6{F0: true, F1: ((true && false) && ((^(v37)) > (v37 / (((v37 * 2147483648)) | 1)))), F2: (v33 ^ 288)}
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        v2 = append(v2, ("hello" + (("hello" + (("hello" + "a") + "hello")) + ("foo bar" + (((("a" + "hello") + "foo bar") + "hello") + "\n\t")))))
                        v34 = 127
                    }
                } else {
                    fmt.Println(v5)
                    switch "hello" {
                    case "foo bar":
                        for fi39 := 0; fi39 < 6; fi39++ {
                            v1 = false
                            if v1 {
                                v40 := GenS5{F0: (uint32(v5) < 2147483647), F1: 128, F2: (("foo bar" + "a") + ("hello" + ("foo bar" + "hello")))}
                                v33 = (v33 % ((973) | 1))
                                v33 = (v33 >> ((2) & 63))
                                _ = v40
                            }
                            _ = fi39
                        }
                    default:
                        if ((2147483647 == (v20 | 33)) || v1) {
                            v33 = (v33 | v33)
                            var v41 bool = v1
                            if (!(-960 <= (v33 + v33))) {
                                var v42 uint8 = v3
                                _ = v42
                            } else {
                                if (true || v1) {
                                    for fi43 := 0; fi43 < 14; fi43++ {
                                        v34 += fi43
                                        if v1 {
                                            switch "" {
                                            case "hello":
                                                var v44 int32 = 350
                                                fmt.Println(v1)
                                                var v45 float32 = float32(float64((0.0 + 3.14159)))
                                                _ = v44
                                                _ = v45
                                            case "foo bar":
                                                if v41 {
                                                    fmt.Println(v20, fi35, fi35)
                                                    v46 := v6[4294967295]
                                                    _ = v46
                                                } else {
                                                    var v47 int = fi43
                                                    v20 = (v20 - int32(v5))
                                                    _ = v47
                                                }
                                            case "":
                                                var v48 uint16 = uint16(v5)
                                                _ = v48
                                            }
                                        }
                                        _ = fi43
                                    }
                                    var v49 float32 = float32(3.14159)
                                    switch "\n\t" {
                                    case "foo bar":
                                        v5 = v5
                                    case "a":
                                        switch "a" {
                                        case "hello":
                                            for fi50 := 0; fi50 < 8; fi50++ {
                                                v6[2147483647] = fi50
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v51 int8 = -96
                                                _ = fi50
                                                _ = v51
                                            }
                                        default:
                                            v52 := GenS7{F0: v34, F1: 0, F2: 291}
                                            _ = v52
                                        }
                                    }
                                    _ = v49
                                }
                                v6[2] = -9223372036854775808
                                v53 := GenS6{F0: v1, F1: false, F2: -1}
                                _ = v53
                            }
                            _ = v41
                        }
                    }
                }
                v54 := GenS6{F0: (!v1), F1: v1, F2: (v33 * v33)}
                v34 -= (fi35 | 2147483647)
                _ = fi35
                _ = v54
            }
            _ = v34
        }
        _ = v33
    }
    var v55 uint64 = v5
    if (!v1) {
        for fi56 := 0; fi56 < 2; fi56++ {
            var v57 int32 = -1
            if ((v57 | 0) >= v20) {
                _ = len(v4)
                fmt.Println(v3, v57, v3)
            }
            _ = fi56
            _ = v57
        }
    } else {
        var v58 uint8 = (v3 % (((v3 ^ (^(v3)))) | 1))
        _ = v58
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v59 := GenS8{F0: true, F1: (v3 | uint8(v5))}
    if v1 {
        v5 += (v55 - v5)
    } else {
        var v60 uint8 = v3
        var v61 float32 = 0.0
        var v62 int64 = -9223372036854775808
        _ = v60
        _ = v61
        _ = v62
    }
    var v63 float32 = (((1.5 + 1.5) - 1.5) * 1e10)
    var v64 bool = v1
    v64 = (true && false)
    v65 := GenS6{F0: v64, F1: (!(v3 != v3)), F2: 42}
    if v1 {
        for fi66 := 0; fi66 < 12; fi66++ {
            var v67 int32 = (v20 >> ((255) & 63))
            switch fi66 {
            case 32768:
                _ = len(v2)
            default:
                var v68 int8 = -58
                _ = v68
            }
            for fi69 := 0; fi69 < 11; fi69++ {
                var v70 int8 = 42
                delete(v6, 158)
                switch (fi69 >> ((544) & 63)) {
                case 127:
                    fmt.Println(v55, v67, fi66)
                case -9223372036854775808:
                    fmt.Println(v70)
                    fmt.Println(v55, v1)
                    if (!v64) {
                        v71 := GenS1{F0: (v70 + v70), F1: ""}
                        var v72 uint = 0
                        v72 = v72
                        var v73 uint32 = 65535
                        if v64 {
                            var v74 float64 = (0.1 - (((3.14159 * 0.1) + float64(v63)) + 1e10))
                            switch (fi66 % ((fi69) | 1)) {
                            case -9223372036854775808:
                                var v75 int = int(v73)
                                _ = v75
                            case 127:
                                v63 -= 0.0
                            default:
                                _ = len(v6)
                            }
                            _ = v74
                        } else {
                            if ((((float64((v63 + (3.14159 / 1.0))) * 3.14159) / 3.14159) - 1e10) == 0.0) {
                                switch (fi66 | fi69) {
                                case -2147483648:
                                    v1 = ((v67 >> ((0) & 63)) == v67)
                                default:
                                    switch "" {
                                    case "hello":
                                        var v76 uint64 = (^(v55))
                                        _ = v76
                                    default:
                                        fmt.Println(fi66, v3, v20)
                                    }
                                    fmt.Println(v73, v70)
                                    if (v64 || v64) {
                                        var v77 uint16 = 2
                                        _ = v77
                                    }
                                }
                                v64 = false
                                fmt.Println(fi66, fi69, v73)
                            }
                            var v78 bool = ((v70 << ((v72) & 63)) < v70)
                            _ = v78
                        }
                        var v79 string = ""
                        _ = v71
                        _ = v72
                        _ = v73
                        _ = v79
                    }
                case 789:
                    if v1 {
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v5)
                        default:
                            v1 = v64
                        }
                        _ = len(v4)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v80 int64 = 1
                        _ = v80
                    } else {
                        var v81 uint = uint(v3)
                        _ = v81
                    }
                }
                _ = fi69
                _ = v70
            }
            _ = fi66
            _ = v67
        }
        for fi82 := 0; fi82 < 6; fi82++ {
            for fi83 := 0; fi83 < 17; fi83++ {
                switch "foo bar" {
                case "":
                    v5 -= (v55 * (v55 + v5))
                case "hello":
                    v84 := GenS4{F0: 127, F1: int8(v20)}
                    _ = v84
                }
                _ = fi83
            }
            v4[4294967295] = 32768
            _ = fi82
        }
    } else {
        v85 := GenS9{F0: 184, F1: 4294967295, F2: 9223372036854775807}
        v1 = false
        switch 127 {
        case 127:
            var v86 int = -2147483648
            _ = v86
        case 9223372036854775807:
            var v87 float64 = 1e10
            _ = v87
        default:
            fmt.Println(v20, v1, v63)
        }
        _ = v85
    }
    fmt.Println(v5, v55)
    fmt.Println(v3, v63)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v59
    _ = v6
    _ = v63
    _ = v64
    _ = v65
}
