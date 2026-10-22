package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 string; F1 uint8 }
func genFn0(p0 int16) uint {
    var v7 float32 = 3.14159
    if false {
        if true {
            v7 -= 3.14159
            v8 := []int{-9223372036854775808}
            _ = v8
        }
        if true {
            var v9 uint8 = 24
            fmt.Println(p0, v7, v7)
            _ = v9
        } else {
            v7 = v7
            for fi10 := 0; fi10 < 13; fi10++ {
                v11 := GenS1{F0: 65535, F1: 1, F2: -9223372036854775808}
                switch (("hello" + (("hello" + "hello") + "hello")) + "foo bar") {
                case "foo bar":
                    v12 := GenS1{F0: 899, F1: 42, F2: -1}
                    _ = v12
                case "\n\t":
                    v7 += (v7 / v7)
                }
                var v13 uint = 873
                _ = fi10
                _ = v11
                _ = v13
            }
            if (!(-242 == 0)) {
                var v14 int = -1
                v7 -= (v7 / v7)
                _ = v14
            }
        }
    } else {
        p0 = 32767
        var v15 bool = (1000 <= (p0 & p0))
        var v16 int16 = (p0 | (p0 & p0))
        _ = v15
        _ = v16
    }
    v17 := GenS2{F0: 42}
    v18 := []int{0}
    for fi19 := 0; fi19 < 5; fi19++ {
        if (false && true) {
            switch 9223372036854775807 {
            case 9223372036854775807:
                v20 := GenS2{F0: 65535}
                _ = v20
            case 32768:
                p0 = p0
            case -279:
                switch ("foo bar" + "\n\t") {
                case "\n\t":
                    v18 = append(v18, fi19)
                }
            }
            v7 = (v7 / -2.25)
            v21 := GenS1{F0: 1, F1: int64(fi19), F2: 1}
            _ = v21
        } else {
            switch fi19 {
            case 86:
                v22 := GenS2{F0: 1}
                _ = v22
            }
            switch (fi19 & (fi19 | (^(fi19)))) {
            case 32768:
                fmt.Println(fi19, fi19)
            }
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        }
        var v23 int16 = int16(fi19)
        _ = fi19
        _ = v23
    }
    _ = p0
    _ = v17
    _ = v18
    _ = v7
    return 269
}
func genFn1() int8 {
    var v24 bool = true
    fmt.Println(v24)
    if v24 {
        var v25 string = ((("a" + "") + ((("foo bar" + "a") + "a") + ("" + ""))) + ((("" + "") + ("hello" + ("\n\t" + "hello"))) + "a"))
        for fi26 := 0; fi26 < 1; fi26++ {
            v27 := GenS1{F0: 1, F1: 42, F2: -1}
            _ = fi26
            _ = v27
        }
        for fi28 := 0; fi28 < 1; fi28++ {
            var v29 int8 = 127
            fmt.Println(fi28)
            _ = fi28
            _ = v29
        }
        _ = v25
    }
    _ = v24
    return -128
}

func main() {
    var v1 int8 = 0
    var v2 int8 = -128
    v3 := make(map[uint32]bool)
    var v4 uint16 = 255
    v5 := make(map[string]string)
    v6 := make(map[int]int)
    if (false || (!false)) {
        v4 = v4
    } else {
        v30 := []int{4294967295}
        _ = v30
    }
    if (false || true) {
        var v31 bool = false
        _ = len(v6)
        v1 += (v1 & int8(v4))
        _ = v31
    }
    v32 := GenS2{F0: 18446744073709551615}
    v33 := GenS1{F0: (v4 | 201), F1: 1, F2: -349}
    var v34 uint32 = 2147483647
    if true {
        for fi35 := 0; fi35 < 5; fi35++ {
            var v36 int = (-(fi35))
            v37 := GenS2{F0: uint(v34)}
            fmt.Println(v1)
            _ = fi35
            _ = v36
            _ = v37
        }
        if (!true) {
            if (!((1.5 * 1.5) >= ((0.1 - 3.14159) / 1e10))) {
                fmt.Println(v2, v4, v4)
                if false {
                    v38 := v5["foo bar"]
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 5; fi39++ {
                        v1 = 0
                        v3[v34] = true
                        switch "a" {
                        case "a":
                            v40 := GenS1{F0: v4, F1: 9223372036854775807, F2: 0}
                            _ = v40
                        }
                        _ = fi39
                    }
                    if false {
                        if ((0.0 + -2.25) < float64(float32((1e10 + 3.14159)))) {
                            fmt.Println(v4, v34)
                        }
                        switch "" {
                        case "":
                            for fi41 := 0; fi41 < 16; fi41++ {
                                fmt.Println(fi41, v1)
                                var v42 float32 = 0.1
                                var v43 bool = (!(v4 < (v4 & (^(v4)))))
                                _ = fi41
                                _ = v42
                                _ = v43
                            }
                        case "a":
                            for fi44 := 0; fi44 < 18; fi44++ {
                                v45 := GenS1{F0: (^(v4)), F1: int64(fi44), F2: 1}
                                v2 -= v1
                                _ = fi44
                                _ = v45
                            }
                        default:
                            v34 -= (v34 | (v34 % ((0) | 1)))
                        }
                        _ = len(v3)
                    } else {
                        v2 = (v1 * v2)
                        var v46 uint64 = 18446744073709551615
                        v1 = (^(v1))
                        _ = v46
                    }
                    if (-423 >= 0) {
                        v47 := []int{}
                        switch 4294967295 {
                        case 32768:
                            v48 := v6[2147483647]
                            _ = v48
                        case 1:
                            if false {
                                if (9223372036854775807 > 0) {
                                    switch "a" {
                                    case "a":
                                        v2 += 1
                                    case "\n\t":
                                        if false {
                                            fmt.Println(v34, v2)
                                            if (440 == uint16(v34)) {
                                                _ = len(v47)
                                            } else {
                                                if true {
                                                    var v49 int = 2147483647
                                                    fmt.Println(v4, v2)
                                                    _ = v49
                                                }
                                            }
                                            if (-2147483648 <= 1) {
                                                for fi50 := 0; fi50 < 15; fi50++ {
                                                    for fi51 := 0; fi51 < 13; fi51++ {
                                                        _ = len(v47)
                                                        _ = fi51
                                                    }
                                                    if len(v47) >= 2 {
                                                        v47 = v47[1:len(v47)]
                                                    }
                                                    _ = fi50
                                                }
                                            }
                                        } else {
                                            v1 += v2
                                            for fi52 := 0; fi52 < 1; fi52++ {
                                                for fi53 := 0; fi53 < 12; fi53++ {
                                                    if true {
                                                        var v54 int64 = 30
                                                        if (0.1 != (float32(float64(float32((3.14159 * 1e10)))) / 1e10)) {
                                                            v55 := GenS3{F0: 1000, F1: 750}
                                                            _ = v55
                                                        } else {
                                                            var v56 float32 = float32(-2.25)
                                                            if (false && (2147483647 >= 42)) {
                                                                if (!(false || false)) {
                                                                    var v57 uint64 = 0
                                                                    v2 = (v2 - (v1 | (v2 >> ((genFn0(1)) & 63))))
                                                                    v1 = v2
                                                                    _ = v57
                                                                }
                                                                delete(v3, (v34 / ((v34) | 1)))
                                                            } else {
                                                                switch fi53 {
                                                                case 2147483647:
                                                                    if true {
                                                                        v56 = v56
                                                                        if len(v47) > 0 {
                                                                            v47[0] = -812
                                                                        }
                                                                        v54 -= (v54 & (v54 + v54))
                                                                    }
                                                                default:
                                                                    v2 = genFn1()
                                                                    _ = len(v5)
                                                                    v34 = v34
                                                                    v58 := GenS3{F0: 1000, F1: (v54 / ((v54) | 1))}
                                                                    _ = v58
                                                                }
                                                                var v59 uint8 = 251
                                                                v60 := GenS4{F0: 1000, F1: int8(v59)}
                                                                _ = v59
                                                                _ = v60
                                                            }
                                                            v56 -= v56
                                                            _ = v56
                                                        }
                                                        v1 += v2
                                                        v34 += (v34 + 287)
                                                        _ = v54
                                                    }
                                                    _ = fi53
                                                }
                                                _ = fi52
                                            }
                                            fmt.Println(v2)
                                        }
                                    case "":
                                        fmt.Println(v1)
                                    default:
                                        for fi61 := 0; fi61 < 9; fi61++ {
                                            if false {
                                                v62 := GenS3{F0: -32768, F1: 1}
                                                v1 += -1
                                                _ = v62
                                            }
                                            v34 -= v34
                                            var v63 int = fi61
                                            _ = fi61
                                            _ = v63
                                        }
                                    }
                                    v1 = -97
                                } else {
                                    var v64 uint64 = 0
                                    var v65 float32 = 0.1
                                    _ = v64
                                    _ = v65
                                }
                                var v66 string = "a"
                                v67 := GenS4{F0: -32768, F1: v1}
                                _ = v66
                                _ = v67
                            }
                        case 42:
                            var v68 uint16 = v4
                            _ = v68
                        default:
                            v34 -= (v34 / ((v34) | 1))
                        }
                        _ = v47
                    }
                }
                if true {
                    delete(v5, ("a" + "\n\t"))
                    var v69 uint64 = 2
                    var v70 int64 = -259
                    _ = v69
                    _ = v70
                } else {
                    v71 := GenS1{F0: (v4 << ((18446744073709551615) & 63)), F1: 0, F2: -1}
                    _ = v71
                }
            } else {
                v72 := GenS1{F0: (v4 / ((1) | 1)), F1: 42, F2: -220}
                _ = v72
            }
        }
        v73 := GenS2{F0: 0}
        _ = v73
    }
    if false {
        var v74 uint16 = v4
        var v75 int8 = genFn1()
        _ = v74
        _ = v75
    }
    if ((v1 - (v2 << ((18446744073709551615) & 63))) >= 0) {
        for fi76 := 0; fi76 < 8; fi76++ {
            if true {
                var v77 int = fi76
                var v78 uint = 42
                _ = v77
                _ = v78
            } else {
                fmt.Println(v2, fi76, v4)
                v5[("hello" + (("hello" + "a") + ("hello" + "foo bar")))] = "a"
                v79 := GenS4{F0: 91, F1: (v1 / ((genFn1()) | 1))}
                _ = v79
            }
            v80 := GenS2{F0: uint(v34)}
            _ = fi76
            _ = v80
        }
        var v81 int64 = 9223372036854775807
        _ = v81
    }
    for fi82 := 0; fi82 < 13; fi82++ {
        var v83 uint64 = 42
        switch (fi82 + 1) {
        case 128:
            var v84 uint8 = 156
            _ = v84
        case 859:
            if (fi82 == fi82) {
                v85 := GenS5{F0: uint64(v34), F1: (false || true), F2: v4}
                var v86 int8 = (v1 ^ 1)
                _ = v85
                _ = v86
            } else {
                v87 := GenS6{F0: 136, F1: -1}
                v83 -= 69
                v83 = v83
                _ = v87
            }
        case -2147483648:
            v88 := GenS7{F0: ("hello" + "\n\t"), F1: 128}
            _ = v88
        }
        _ = fi82
        _ = v83
    }
    delete(v6, -9223372036854775808)
    v4 += 42
    var v89 float64 = ((0.0 / 3.14159) * (1.5 / 1.0))
    var v90 float32 = (0.1 * 3.14159)
    if ((true && false) && (true && (0 >= 2))) {
        for fi91 := 0; fi91 < 14; fi91++ {
            var v92 uint8 = 127
            for fi93 := 0; fi93 < 14; fi93++ {
                switch "foo bar" {
                case "a":
                    var v94 int32 = 42
                    _ = v94
                case "":
                    v95 := []int{0, -2147483648}
                    _ = v95
                case "hello":
                    v96 := v3[2]
                    _ = v96
                }
                v97 := []int{fi91, (fi93 / ((32767) | 1))}
                _ = fi93
                _ = v97
            }
            fmt.Println(v4)
            _ = fi91
            _ = v92
        }
        v5["\n\t"] = ("hello" + "hello")
    } else {
        v2 -= v1
        v98 := []int{}
        _ = v98
    }
    for fi99 := 0; fi99 < 20; fi99++ {
        fmt.Println(v90)
        _ = fi99
    }
    v100 := v5[(((("" + (("hello" + ("a" + "foo bar")) + "\n\t")) + "") + "a") + "\n\t")]
    fmt.Println(v4, v90)
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v89
    _ = v90
}
