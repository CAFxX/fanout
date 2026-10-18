package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint; F1 uint32 }
type GenS8 struct { F0 int64; F1 int16; F2 int16 }
func genFn0(p0 string, p1 uint64) int8 {
    p0 = ((p0 + ((p0 + p0) + (p0 + ((("a" + p0) + ("\n\t" + ("foo bar" + "\n\t"))) + p0)))) + "")
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 int8 = 42
        var v8 uint64 = (p1 % ((1) | 1))
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: 2147483647, F1: -607}
        if (0.1 <= 0.0) {
            for fi11 := 0; fi11 < 16; fi11++ {
                p1 -= p1
                fmt.Println(fi9, p0)
                _ = fi11
            }
        } else {
            for fi12 := 0; fi12 < 15; fi12++ {
                switch p0 {
                case "foo bar":
                    var v13 float32 = 0.0
                    _ = v13
                default:
                    switch (("foo bar" + p0) + ((p0 + p0) + "")) {
                    case "\n\t":
                        p0 = p0
                    }
                }
                _ = fi12
            }
            var v14 uint = 255
            _ = v14
        }
        var v15 uint16 = 563
        _ = fi9
        _ = v10
        _ = v15
    }
    _ = p0
    _ = p1
    return 100
}
func genFn1(p0 string) int64 {
    p0 = "a"
    switch p0 {
    case "foo bar":
        if ((false && ((!false) && false)) || true) {
            p0 = (p0 + "")
        }
    default:
        p0 = ("foo bar" + p0)
    }
    _ = p0
    return -1
}

func main() {
    var v1 uint8 = 255
    var v2 float32 = 1e10
    v3 := make(map[uint32]int)
    v4 := []string{"\n\t", "a", "foo bar", ("\n\t" + "")}
    var v5 int16 = 32767
    fmt.Println(v1, v1)
    v16 := GenS1{F0: 42, F1: 87}
    fmt.Println(v2)
    delete(v3, 0)
    if (false && true) {
        v17 := GenS1{F0: 556, F1: -9223372036854775808}
        var v18 uint64 = 724
        var v19 int32 = 42
        _ = v17
        _ = v18
        _ = v19
    } else {
        for fi20 := 0; fi20 < 8; fi20++ {
            v2 = v2
            for fi21 := 0; fi21 < 9; fi21++ {
                var v22 string = "foo bar"
                v23 := GenS2{F0: 1, F1: v1, F2: 1}
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 string = "\n\t"
            _ = fi20
            _ = v24
        }
    }
    if false {
        v4 = append(v4, "\n\t")
    } else {
        v5 -= (v5 / (((v5 - v5)) | 1))
    }
    v5 += 0
    for fi25 := 0; fi25 < 1; fi25++ {
        v26 := GenS3{F0: 100}
        for fi27 := 0; fi27 < 12; fi27++ {
            switch fi25 {
            case 9223372036854775807:
                if false {
                    v4 = append(v4, "hello")
                    if false {
                        v5 -= v5
                        v28 := GenS4{F0: 2147483648, F1: 2147483648}
                        fmt.Println(fi27)
                        v29 := GenS5{F0: 859}
                        if (!true) {
                            var v30 float64 = 3.14159
                            _ = v30
                        } else {
                            switch "hello" {
                            case "hello":
                                v4 = append(v4, ((("a" + ((((("hello" + "") + ("hello" + "a")) + ("\n\t" + "")) + "\n\t") + "")) + (("\n\t" + ((("hello" + ("hello" + "hello")) + "") + "hello")) + "foo bar")) + (("hello" + "") + ("\n\t" + ("hello" + ((((("" + "hello") + ("" + "a")) + "foo bar") + "\n\t") + "foo bar"))))))
                            case "":
                                v2 = v2
                            }
                        }
                        var v31 uint64 = 920
                        _ = v28
                        _ = v29
                        _ = v31
                    } else {
                        var v32 string = ((("foo bar" + "\n\t") + "") + "")
                        _ = v32
                    }
                    v2 -= v2
                } else {
                    if (fi27 != 0) {
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        var v33 string = "a"
                        v34 := GenS1{F0: 2147483648, F1: 32768}
                        _ = v33
                        _ = v34
                    } else {
                        var v35 int32 = 0
                        var v36 uint16 = 32767
                        var v37 int = (fi25 - fi27)
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    var v38 int = (fi27 * (fi27 >> ((18446744073709551615) & 63)))
                    _ = v38
                }
            case 4294967295:
                if ((true || (false || (!false))) || true) {
                    var v39 uint32 = 924
                    _ = v39
                }
            case -49:
                if ("foo bar" == ("foo bar" + ("hello" + "foo bar"))) {
                    if true {
                        var v40 int32 = 42
                        var v41 uint = 2
                        fmt.Println(v41)
                        if len(v4) > 0 {
                            v4[0] = (((("" + "\n\t") + ("" + ("" + ""))) + ("hello" + "a")) + "foo bar")
                        }
                        _ = v40
                        _ = v41
                    }
                    var v42 int = -916
                    v43 := GenS6{F0: 1}
                    v44 := GenS6{F0: 42}
                    var v45 uint64 = 1
                    _ = v42
                    _ = v43
                    _ = v44
                    _ = v45
                }
            default:
                switch (-(fi27)) {
                case 32768:
                    switch (fi27 + (fi27 | (-(fi25)))) {
                    case 32768:
                        v1 = (v1 * v1)
                    default:
                        var v46 uint32 = uint32(v1)
                        _ = v46
                    }
                case -470:
                    v47 := GenS2{F0: 1, F1: v1, F2: 4294967295}
                    _ = v47
                case 42:
                    var v48 bool = (706 != 2)
                    if (19 == 0) {
                        switch ("" + "a") {
                        case "a":
                            var v49 int16 = v5
                            _ = v49
                        case "\n\t":
                            v2 -= ((((1.5 * v2) - float32(3.14159)) / v2) + (-2.25 + (0.1 * float32((1e10 - -2.25)))))
                        case "foo bar":
                            var v50 uint32 = 42
                            _ = v50
                        }
                    }
                    fmt.Println(fi25, v2)
                    _ = v48
                }
            }
            _ = fi27
        }
        if (-92 <= 0) {
            v5 = v5
        } else {
            var v51 uint16 = 0
            if false {
                v52 := GenS2{F0: v51, F1: 42, F2: 255}
                _ = v52
            }
            _ = v51
        }
        _ = fi25
        _ = v26
    }
    for fi53 := 0; fi53 < 8; fi53++ {
        var v54 int16 = v5
        v55 := GenS4{F0: 259, F1: 2147483647}
        _ = fi53
        _ = v54
        _ = v55
    }
    v5 = v5
    fmt.Println(v2, v1, v1)
    var v56 bool = ((1 > (v5 - (v5 | (^(v5))))) && (!true))
    v57 := GenS7{F0: 1, F1: uint32(v1)}
    if v56 {
        for fi58 := 0; fi58 < 19; fi58++ {
            if (!(!v56)) {
                _ = len(v4)
                v5 = v5
                v59 := GenS1{F0: 42, F1: fi58}
                _ = v59
            }
            _ = fi58
        }
    }
    v1 = 208
    for fi60 := 0; fi60 < 1; fi60++ {
        for fi61 := 0; fi61 < 20; fi61++ {
            v62 := GenS3{F0: -1}
            _ = len(v3)
            delete(v3, 4294967295)
            _ = fi61
            _ = v62
        }
        v5 = v5
        if true {
            v5 = (v5 ^ 684)
        } else {
            if (255 != 206) {
                switch fi60 {
                case 128:
                    if v56 {
                        for fi63 := 0; fi63 < 2; fi63++ {
                            v1 -= 1
                            _ = fi63
                        }
                    } else {
                        v64 := GenS4{F0: 2, F1: 588}
                        _ = len(v4)
                        if true {
                            for fi65 := 0; fi65 < 9; fi65++ {
                                switch (^(fi60)) {
                                case -173:
                                    var v66 bool = v56
                                    _ = v66
                                }
                                fmt.Println(fi60)
                                var v67 uint32 = 213
                                _ = fi65
                                _ = v67
                            }
                            var v68 uint8 = v1
                            for fi69 := 0; fi69 < 11; fi69++ {
                                var v70 float32 = ((1e10 - v2) / v2)
                                var v71 uint16 = 397
                                v72 := GenS4{F0: 0, F1: 362}
                                _ = fi69
                                _ = v70
                                _ = v71
                                _ = v72
                            }
                            _ = v68
                        } else {
                            for fi73 := 0; fi73 < 12; fi73++ {
                                if len(v4) > 0 {
                                    v4[0] = ((("a" + "foo bar") + "a") + "a")
                                }
                                _ = fi73
                            }
                            var v74 int = (fi60 | (fi60 / (((fi60 % (((fi60 % ((-878) | 1))) | 1))) | 1)))
                            _ = len(v3)
                            _ = v74
                        }
                        _ = v64
                    }
                case 9223372036854775807:
                    _ = len(v4)
                }
                v75 := GenS8{F0: 42, F1: (v5 % ((v5) | 1)), F2: (^(v5))}
                var v76 uint = 42
                _ = v75
                _ = v76
            } else {
                switch fi60 {
                case -762:
                    v77 := GenS6{F0: 42}
                    _ = v77
                case 2147483647:
                    if len(v4) > 0 {
                        v4[0] = ("" + "\n\t")
                    }
                default:
                    var v78 uint32 = 819
                    _ = v78
                }
                var v79 uint = 0
                var v80 bool = ((!v56) || true)
                _ = v79
                _ = v80
            }
            v5 += v5
            v81 := GenS3{F0: -46}
            _ = v81
        }
        _ = fi60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
}
