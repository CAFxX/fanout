package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 string; F2 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 int; F2 uint64 }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 int16; F1 bool; F2 uint32 }
func genFn0(p0 int32) uint64 {
    fmt.Println(p0)
    v6 := []int{-2147483648, 128, 2147483647}
    _ = len(v6)
    _ = p0
    _ = v6
    return 2
}
func genFn1(p0 uint8, p1 string, p2 int16) uint {
    fmt.Println(p2, p0)
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int8 = -128
        var v9 uint16 = 2
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        v10 := []int{919, 4294967295}
        p0 = p0
        var v11 float64 = 1.5
        _ = v10
        _ = v11
    }
    if (true || (9223372036854775807 != -2147483648)) {
        p1 = p1
        if (421 > 2147483648) {
            p0 = 2
            for fi12 := 0; fi12 < 3; fi12++ {
                v13 := GenS1{F0: true}
                _ = fi12
                _ = v13
            }
            var v14 uint64 = 18446744073709551615
            _ = v14
        } else {
            var v15 int8 = -1
            v16 := GenS1{F0: true}
            _ = v15
            _ = v16
        }
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint32 = 2
    v2 := []string{"\n\t"}
    var v3 int8 = 115
    v4 := make(map[string]int)
    var v5 int8 = (v3 << ((804) & 63))
    if false {
        var v17 string = ((((((("hello" + ("foo bar" + "\n\t")) + "a") + "\n\t") + "foo bar") + ("" + ("" + "\n\t"))) + "foo bar") + "hello")
        _ = v17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        fmt.Println(v5)
        if (42 != -518) {
            v19 := GenS2{F0: 1, F1: "\n\t", F2: 1}
            v20 := GenS2{F0: -32768, F1: ("\n\t" + "\n\t"), F2: -1}
            if false {
                for fi21 := 0; fi21 < 2; fi21++ {
                    v1 = (^(v1))
                    fmt.Println(v5, v1)
                    _ = fi21
                }
            } else {
                var v22 float64 = (0.0 - float64((1e10 - 0.1)))
                _ = v22
            }
            _ = v19
            _ = v20
        } else {
            v5 += (v3 + v3)
            var v23 float64 = 3.14159
            _ = v23
        }
        _ = fi18
    }
    v1 -= (v1 + (v1 + 2))
    if false {
        if (!true) {
            var v24 uint32 = (v1 << ((18446744073709551615) & 63))
            if false {
                var v25 uint8 = 0
                for fi26 := 0; fi26 < 7; fi26++ {
                    v27 := GenS2{F0: 1, F1: "hello", F2: -400}
                    var v28 uint8 = v25
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                switch 0 {
                case -2147483648:
                    v29 := GenS1{F0: false}
                    _ = v29
                case -603:
                    v25 = (v25 - 2)
                case -600:
                    var v30 string = ("foo bar" + (("\n\t" + "a") + "foo bar"))
                    _ = v30
                default:
                    if false {
                        if true {
                            if false {
                                var v31 bool = false
                                _ = v31
                            }
                            if ((v5 % (((v5 * 127)) | 1)) <= (v3 / (((v5 >> ((65535) & 63))) | 1))) {
                                v32 := GenS2{F0: -111, F1: "foo bar", F2: 400}
                                var v33 uint16 = 42
                                v24 += 65535
                                _ = v32
                                _ = v33
                            } else {
                                var v34 uint32 = 42
                                var v35 uint64 = 232
                                _ = v34
                                _ = v35
                            }
                            if len(v2) > 0 {
                                v2[0] = ("foo bar" + "hello")
                            }
                        } else {
                            fmt.Println(v25, v5)
                            if false {
                                v5 = v3
                                if false {
                                    v36 := GenS3{F0: 4294967295}
                                    _ = v36
                                } else {
                                    if true {
                                        v5 -= 100
                                        if len(v2) > 0 {
                                            v2[0] = "foo bar"
                                        }
                                    } else {
                                        for fi37 := 0; fi37 < 13; fi37++ {
                                            if false {
                                                var v38 uint64 = 876
                                                if len(v2) > 0 {
                                                    v2[0] = ("hello" + ("\n\t" + (("" + ("\n\t" + "")) + "a")))
                                                }
                                                for fi39 := 0; fi39 < 3; fi39++ {
                                                    var v40 int = fi39
                                                    v2 = append(v2, ((("foo bar" + ("" + "a")) + ("a" + "foo bar")) + ""))
                                                    _ = fi39
                                                    _ = v40
                                                }
                                                _ = v38
                                            }
                                            delete(v4, ("" + ""))
                                            _ = fi37
                                        }
                                        v3 -= (v5 * (v5 * v3))
                                    }
                                    if (42 >= genFn0(-1)) {
                                        if false {
                                            v3 += v3
                                            var v41 string = "a"
                                            switch "foo bar" {
                                            case "":
                                                for fi42 := 0; fi42 < 8; fi42++ {
                                                    var v43 bool = false
                                                    if v43 {
                                                        v5 -= (v5 << ((2) & 63))
                                                        v44 := GenS2{F0: -896, F1: (v41 + v41), F2: 533}
                                                        _ = v44
                                                    } else {
                                                        v5 += v3
                                                        var v45 int32 = 259
                                                        v46 := GenS4{F0: 1}
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                    var v47 uint32 = v24
                                                    _ = fi42
                                                    _ = v43
                                                    _ = v47
                                                }
                                            default:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            _ = v41
                                        }
                                        v3 = v5
                                        if (("foo bar" > "a") && ((v24 | (v24 ^ (v1 + (v24 ^ v24)))) > v1)) {
                                            fmt.Println(v24, v3, v25)
                                            var v48 int8 = -108
                                            _ = v48
                                        }
                                    }
                                }
                            }
                        }
                        fmt.Println(v5, v5, v24)
                        var v49 uint64 = 4294967295
                        _ = v49
                    } else {
                        v5 = v3
                        if (4294967295 == 42) {
                            v5 = v3
                        } else {
                            if (true && true) {
                                delete(v4, "foo bar")
                                var v50 bool = false
                                var v51 int16 = 854
                                _ = v50
                                _ = v51
                            }
                            v24 += (v1 >> ((2) & 63))
                            switch "a" {
                            case "":
                                v52 := v4["\n\t"]
                                _ = v52
                            case "hello":
                                var v53 uint = uint(v25)
                                _ = v53
                            case "foo bar":
                                if false {
                                    for fi54 := 0; fi54 < 14; fi54++ {
                                        for fi55 := 0; fi55 < 4; fi55++ {
                                            switch fi54 {
                                            case 709:
                                                v3 = (v3 | v3)
                                            default:
                                                var v56 uint = 54
                                                _ = v56
                                            }
                                            switch ("" + (("hello" + "hello") + ((("a" + "a") + "foo bar") + ""))) {
                                            case "hello":
                                                v24 = (v1 | (^(v1)))
                                            case "a":
                                                var v57 int16 = -947
                                                _ = v57
                                            case "\n\t":
                                                v58 := GenS2{F0: int16(fi54), F1: "", F2: -100}
                                                _ = v58
                                            default:
                                                v5 = v5
                                            }
                                            var v59 float32 = ((1e10 + (((1.5 / ((-2.25 / 1.0) + (-2.25 * 1e10))) - 1.5) / 1.5)) * (((3.14159 - 1.5) / -2.25) * -2.25))
                                            _ = fi55
                                            _ = v59
                                        }
                                        var v60 uint = 1
                                        _ = fi54
                                        _ = v60
                                    }
                                }
                            }
                        }
                        v25 = 236
                    }
                }
                _ = v25
            }
            v24 -= 0
            _ = v24
        }
    } else {
        v1 = (^(v1))
        var v61 uint32 = v1
        _ = v61
    }
    fmt.Println(v3, v5)
    v3 = (v5 - (v3 | (v5 << ((0) & 63))))
    for fi62 := 0; fi62 < 14; fi62++ {
        for fi63 := 0; fi63 < 8; fi63++ {
            _ = len(v2)
            _ = fi63
        }
        _ = fi62
    }
    v2 = append(v2, "foo bar")
    var v64 int = -2147483648
    if false {
        if ((uint16(v1) <= 788) || true) {
            if (42 < 18446744073709551615) {
                v64 -= v64
                switch 4294967295 {
                case -712:
                    if ((!true) || false) {
                        if true {
                            for fi65 := 0; fi65 < 12; fi65++ {
                                if (false && (v3 > (v3 - (v5 | 100)))) {
                                    switch (-(fi65)) {
                                    case 127:
                                        fmt.Println(v5, v3, v3)
                                    default:
                                        var v66 string = ((("" + "foo bar") + "hello") + (("\n\t" + "a") + "\n\t"))
                                        _ = v66
                                    }
                                    v64 = (fi65 & 32768)
                                    v67 := GenS4{F0: 42}
                                    _ = v67
                                }
                                v68 := GenS2{F0: 0, F1: "\n\t", F2: -1}
                                var v69 uint64 = 624
                                _ = fi65
                                _ = v68
                                _ = v69
                            }
                        } else {
                            v64 = v64
                            for fi70 := 0; fi70 < 5; fi70++ {
                                v2 = append(v2, "hello")
                                var v71 int8 = 97
                                _ = fi70
                                _ = v71
                            }
                            var v72 uint32 = v1
                            _ = v72
                        }
                    }
                case 127:
                    var v73 float32 = ((1e10 / 1.0) + ((-2.25 * -2.25) / (1.5 - 1e10)))
                    _ = v73
                }
                for fi74 := 0; fi74 < 2; fi74++ {
                    for fi75 := 0; fi75 < 8; fi75++ {
                        if (-1 != -14) {
                            v3 = (v5 % (((^(v3))) | 1))
                            v5 = v3
                            if (((0.1 * float64(3.14159)) / (0.1 * -2.25)) > 0.1) {
                                v76 := v4[("" + "\n\t")]
                                var v77 int32 = 583
                                _ = v76
                                _ = v77
                            }
                        }
                        if true {
                            v78 := GenS4{F0: 65535}
                            var v79 int = (fi74 + 2147483647)
                            var v80 int64 = 9223372036854775807
                            _ = v78
                            _ = v79
                            _ = v80
                        } else {
                            v5 -= 1
                            v3 -= v5
                        }
                        _ = len(v2)
                        _ = fi75
                    }
                    v81 := GenS5{F0: 4294967295, F1: -54, F2: 8}
                    _ = fi74
                    _ = v81
                }
            }
            if (!("" < "foo bar")) {
                v82 := GenS6{F0: v1, F1: -289}
                v83 := GenS3{F0: 1}
                switch ("a" + "") {
                case "hello":
                    v2 = append(v2, "")
                case "\n\t":
                    if (!false) {
                        v1 -= (v1 << ((0) & 63))
                        v1 -= (^(v1))
                        if false {
                            var v84 int16 = 839
                            _ = v84
                        }
                    } else {
                        if (v64 < (v64 ^ v64)) {
                            v3 = 1
                        }
                        v3 = (^(v5))
                    }
                case "foo bar":
                    v85 := GenS7{F0: 0, F1: false, F2: (v1 & (v1 << ((2) & 63)))}
                    _ = v85
                default:
                    delete(v4, "foo bar")
                }
                _ = v82
                _ = v83
            }
            for fi86 := 0; fi86 < 10; fi86++ {
                var v87 string = "\n\t"
                if true {
                    if true {
                        v5 -= 100
                    } else {
                        fmt.Println(v64, v1)
                        fmt.Println(fi86)
                        v88 := GenS7{F0: 0, F1: true, F2: (v1 >> ((79) & 63))}
                        _ = v88
                    }
                } else {
                    fmt.Println(v3, v1)
                    if (true && ((((-1 < 1) && false) && (false || true)) && false)) {
                        fmt.Println(v3, v3)
                        if (1e10 != 0.1) {
                            fmt.Println(v5)
                            switch v87 {
                            case "foo bar":
                                fmt.Println(v3, v3, fi86)
                            case "a":
                                v5 += 0
                            case "\n\t":
                                if (!true) {
                                    if true {
                                        var v89 int = fi86
                                        _ = v89
                                    }
                                } else {
                                    v5 = v3
                                }
                            }
                        } else {
                            fmt.Println(v87, v3)
                            var v90 uint64 = 1
                            _ = v90
                        }
                    }
                    var v91 uint64 = uint64(v1)
                    _ = v91
                }
                _ = fi86
                _ = v87
            }
        }
        for fi92 := 0; fi92 < 2; fi92++ {
            var v93 uint16 = 65535
            var v94 uint = 4294967295
            v95 := GenS3{F0: 1}
            _ = fi92
            _ = v93
            _ = v94
            _ = v95
        }
    }
    delete(v4, "")
    for fi96 := 0; fi96 < 13; fi96++ {
        v5 = v3
        v2 = append(v2, "")
        _ = fi96
    }
    var v97 string = (("a" + (((("hello" + "hello") + "\n\t") + (("\n\t" + (("" + "foo bar") + "hello")) + "\n\t")) + "\n\t")) + "")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v97
}
