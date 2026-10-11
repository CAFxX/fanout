package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 uint8; F1 int16; F2 int32 }
type GenS8 struct { F0 int; F1 int8; F2 int32 }
type GenS9 struct { F0 uint32; F1 uint16; F2 uint }
type GenS10 struct { F0 uint64 }
func genFn0(p0 int8, p1 int8) string {
    var v7 int16 = 1000
    if true {
        var v8 uint16 = 653
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(p0)
            v7 -= -688
            _ = fi9
        }
        for fi10 := 0; fi10 < 12; fi10++ {
            p1 -= p0
            var v11 int = fi10
            switch (-(v11)) {
            case -1:
                for fi12 := 0; fi12 < 5; fi12++ {
                    if (42 >= (v8 ^ v8)) {
                        v13 := GenS1{F0: 32767}
                        v7 -= v7
                        _ = v13
                    }
                    v14 := GenS2{F0: (p1 >> ((0) & 63))}
                    _ = fi12
                    _ = v14
                }
            case 32767:
                v15 := []int{(fi10 & 32768), -1, fi10}
                _ = v15
            case 484:
                var v16 uint64 = uint64(v8)
                _ = v16
            default:
                var v17 uint8 = 154
                _ = v17
            }
            _ = fi10
            _ = v11
        }
        _ = v8
    }
    switch -1 {
    case 127:
        var v18 uint = 201
        _ = v18
    case -857:
        if (!(437 > 1)) {
            var v19 bool = false
            _ = v19
        } else {
            v7 = v7
        }
    default:
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 string = "hello"
            fmt.Println(fi20, v21, v7)
            v7 = (v7 + 389)
            _ = fi20
            _ = v21
        }
    }
    switch -338 {
    case 128:
        v22 := []int{-369}
        _ = v22
    default:
        if false {
            v23 := GenS2{F0: p1}
            if (823 < 2) {
                v24 := GenS3{F0: 0}
                p1 += p1
                _ = v24
            } else {
                p0 -= p1
            }
            for fi25 := 0; fi25 < 11; fi25++ {
                for fi26 := 0; fi26 < 18; fi26++ {
                    v27 := []int{-2147483648, 128, 848, fi25}
                    v28 := GenS3{F0: 42}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = fi25
            }
            _ = v23
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return ("\n\t" + "")
}

func main() {
    var v1 string = "hello"
    var v2 int16 = -940
    var v3 int64 = 9223372036854775807
    var v4 uint64 = 762
    var v5 bool = true
    var v6 int32 = 0
    var v29 int32 = v6
    v2 = v2
    v30 := GenS2{F0: 100}
    v3 = v3
    v31 := []int{}
    for fi32 := 0; fi32 < 15; fi32++ {
        v2 -= 32767
        _ = fi32
    }
    v4 -= 0
    for fi33 := 0; fi33 < 1; fi33++ {
        for fi34 := 0; fi34 < 15; fi34++ {
            if false {
                v35 := GenS2{F0: int8(v2)}
                _ = v35
            } else {
                var v36 int32 = v29
                var v37 uint64 = v4
                v38 := GenS4{F0: v6, F1: 42}
                _ = v36
                _ = v37
                _ = v38
            }
            fmt.Println(v5, fi34, v3)
            _ = fi34
        }
        _ = fi33
    }
    var v39 float32 = (0.0 - (0.0 * (((3.14159 / float32(0.1)) * 0.0) - -2.25)))
    var v40 uint16 = 32767
    var v41 uint = 255
    if false {
        var v42 bool = true
        _ = v42
    } else {
        if false {
            if v5 {
                var v43 uint = (^(v41))
                _ = v43
            } else {
                var v44 uint = (^(v41))
                _ = v44
            }
        } else {
            v3 += (v3 * 183)
            v45 := GenS3{F0: 0}
            _ = v45
        }
        var v46 bool = v5
        _ = v46
    }
    fmt.Println(v4, v1)
    v47 := GenS3{F0: (v41 + 226)}
    if ((v3 / ((42) | 1)) != v3) {
        for fi48 := 0; fi48 < 14; fi48++ {
            v3 = int64(v41)
            _ = fi48
        }
    } else {
        for fi49 := 0; fi49 < 7; fi49++ {
            v6 = v6
            if len(v31) > 0 {
                v31[0] = fi49
            }
            _ = fi49
        }
        if v5 {
            switch (v1 + v1) {
            case "\n\t":
                var v50 uint = 198
                _ = v50
            default:
                v51 := GenS3{F0: 2}
                _ = v51
            }
            v52 := GenS5{F0: ("foo bar" + v1)}
            if (v40 <= (v40 - (^(v40)))) {
                v1 = ("a" + (v1 + v1))
                switch v1 {
                case "a":
                    v6 += -697
                case "foo bar":
                    var v53 string = "hello"
                    _ = v53
                }
                v6 -= v6
            }
            _ = v52
        }
    }
    switch (("" + v1) + ((("foo bar" + v1) + v1) + v1)) {
    case "\n\t":
        v31 = append(v31, 42)
    case "":
        if (v5 && true) {
            if ((v2 + v2) >= v2) {
                for fi54 := 0; fi54 < 11; fi54++ {
                    switch (fi54 % ((9223372036854775807) | 1)) {
                    case 402:
                        var v55 int64 = (v3 & (v3 % ((1) | 1)))
                        _ = v55
                    case 32767:
                        var v56 uint64 = v4
                        _ = v56
                    default:
                        if (v5 && true) {
                            _ = len(v31)
                            if len(v31) >= 2 {
                                v31 = v31[1:len(v31)]
                            }
                            v57 := GenS6{F0: 42, F1: (false || (-59 >= (^(v3))))}
                            _ = v57
                        } else {
                            switch fi54 {
                            case 27:
                                _ = len(v31)
                            default:
                                v39 = v39
                            }
                            if (!(!(42 <= v40))) {
                                fmt.Println(v39, v41)
                                switch v1 {
                                case "\n\t":
                                    var v58 float32 = (v39 + ((((1e10 - v39) + v39) / 1e10) - v39))
                                    _ = v58
                                case "a":
                                    if true {
                                        for fi59 := 0; fi59 < 13; fi59++ {
                                            if (false || true) {
                                                var v60 uint16 = (v40 >> ((1) & 63))
                                                v61 := GenS7{F0: 255, F1: v2, F2: int32(v3)}
                                                _ = v60
                                                _ = v61
                                            } else {
                                                v40 -= (v40 - v40)
                                            }
                                            v62 := GenS8{F0: -694, F1: int8(v3), F2: v6}
                                            var v63 int64 = 42
                                            _ = fi59
                                            _ = v62
                                            _ = v63
                                        }
                                    } else {
                                        var v64 int32 = -104
                                        var v65 bool = ((v1 + v1) != v1)
                                        if v65 {
                                            var v66 uint = (^(v41))
                                            var v67 float32 = (3.14159 - (float32(0.1) / float32(1e10)))
                                            _ = v66
                                            _ = v67
                                        }
                                        _ = v64
                                        _ = v65
                                    }
                                default:
                                    var v68 string = ("foo bar" + (v1 + v1))
                                    _ = v68
                                }
                                if ((1 >= 854) || v5) {
                                    var v69 string = ("foo bar" + "a")
                                    if len(v31) >= 2 {
                                        v31 = v31[1:len(v31)]
                                    }
                                    fmt.Println(v2, v41, v2)
                                    _ = v69
                                } else {
                                    var v70 uint16 = 0
                                    var v71 uint8 = 194
                                    v3 -= (v3 - 9223372036854775807)
                                    _ = v70
                                    _ = v71
                                }
                            } else {
                                var v72 float32 = (1.5 * float32(3.14159))
                                _ = v72
                            }
                            switch v1 {
                            case "\n\t":
                                var v73 bool = (v1 > v1)
                                _ = v73
                            case "hello":
                                var v74 float32 = v39
                                _ = v74
                            case "foo bar":
                                if v5 {
                                    for fi75 := 0; fi75 < 14; fi75++ {
                                        if v5 {
                                            var v76 int8 = 1
                                            switch (v1 + "\n\t") {
                                            case "hello":
                                                v77 := GenS9{F0: 2147483648, F1: v40, F2: 687}
                                                _ = v77
                                            case "":
                                                var v78 int64 = v3
                                                _ = v78
                                            }
                                            v4 += 56
                                            var v79 bool = v5
                                            var v80 int64 = -9223372036854775808
                                            if len(v31) >= 2 {
                                                v31 = v31[1:len(v31)]
                                            }
                                            _ = v76
                                            _ = v79
                                            _ = v80
                                        }
                                        v3 = v3
                                        _ = fi75
                                    }
                                }
                            }
                        }
                    }
                    for fi81 := 0; fi81 < 14; fi81++ {
                        v82 := GenS10{F0: v4}
                        v31 = append(v31, fi81)
                        _ = fi81
                        _ = v82
                    }
                    v41 -= (v41 * v41)
                    _ = fi54
                }
                _ = len(v31)
                v39 += v39
            } else {
                v29 = v6
                var v83 uint32 = 234
                _ = v83
            }
        }
    case "hello":
        if len(v31) >= 2 {
            v31 = v31[1:len(v31)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v47
    _ = v5
    _ = v6
}
