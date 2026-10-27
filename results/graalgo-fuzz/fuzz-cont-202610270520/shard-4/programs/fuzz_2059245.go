package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 int32; F2 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 int; F2 uint }
type GenS9 struct { F0 uint64; F1 int16; F2 int8 }
type GenS10 struct { F0 uint8; F1 int64; F2 int16 }
func genFn0(p0 int32, p1 bool) uint32 {
    for fi6 := 0; fi6 < 5; fi6++ {
        p0 = (p0 * p0)
        fmt.Println(fi6, fi6, fi6)
        _ = fi6
    }
    var v7 uint16 = 0
    v8 := []int{127, 9223372036854775807, 4294967295, 32767}
    v7 = v7
    if (!p1) {
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 uint16 = (v7 & (^(v7)))
            if len(v8) > 0 {
                v8[0] = 32768
            }
            var v11 float32 = 0.1
            _ = fi9
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 12; fi12++ {
            if (!(0 < 0)) {
                if len(v8) > 0 {
                    v8[0] = fi12
                }
                fmt.Println(p0, v7)
                var v13 string = "hello"
                _ = v13
            } else {
                if (!(p0 != (^(p0)))) {
                    v7 -= (v7 % (((^(v7))) | 1))
                    if (p1 || false) {
                        v14 := GenS1{F0: 2147483647, F1: 42}
                        for fi15 := 0; fi15 < 20; fi15++ {
                            if (!p1) {
                                p1 = p1
                            } else {
                                switch fi15 {
                                case 0:
                                    var v16 int8 = 0
                                    _ = v16
                                default:
                                    p0 += 113
                                }
                                _ = len(v8)
                            }
                            _ = len(v8)
                            _ = fi15
                        }
                        _ = v14
                    } else {
                        for fi17 := 0; fi17 < 20; fi17++ {
                            v18 := GenS1{F0: 2147483648, F1: 42}
                            _ = fi17
                            _ = v18
                        }
                        v19 := GenS2{F0: (p0 - p0), F1: 2}
                        _ = v19
                    }
                    if (44 == 39) {
                        fmt.Println(v7, p0)
                    }
                }
                var v20 bool = p1
                for fi21 := 0; fi21 < 5; fi21++ {
                    _ = len(v8)
                    switch "" {
                    case "":
                        v22 := GenS3{F0: int8(fi12)}
                        _ = v22
                    case "a":
                        p0 -= (-(p0))
                    default:
                        var v23 uint = 65535
                        _ = v23
                    }
                    _ = fi21
                }
                _ = v20
            }
            var v24 uint = 42
            _ = fi12
            _ = v24
        }
    } else {
        for fi25 := 0; fi25 < 15; fi25++ {
            for fi26 := 0; fi26 < 11; fi26++ {
                if false {
                    if (3.14159 <= 1e10) {
                        if (!p1) {
                            p1 = (v7 <= 42)
                        } else {
                            v27 := GenS1{F0: 257, F1: 65535}
                            _ = v27
                        }
                        p1 = true
                        p1 = p1
                        fmt.Println(p1)
                    } else {
                        var v28 uint64 = 42
                        v29 := GenS1{F0: 2147483647, F1: 42}
                        _ = v28
                        _ = v29
                    }
                    if (1 < 255) {
                        switch "hello" {
                        case "\n\t":
                            var v30 int32 = (p0 << ((1) & 63))
                            _ = v30
                        default:
                            if len(v8) > 0 {
                                v8[0] = (fi26 & fi25)
                            }
                            if (42 < 146) {
                                var v31 int = 9223372036854775807
                                p1 = (!(!p1))
                                fmt.Println(fi25)
                                _ = v31
                            }
                            switch 127 {
                            case 32768:
                                switch (fi25 ^ -858) {
                                case 127:
                                    v7 = v7
                                case 350:
                                    fmt.Println(fi25, v7)
                                }
                            case 1:
                                var v32 uint16 = v7
                                _ = v32
                            }
                        }
                        v33 := GenS3{F0: -1}
                        _ = v33
                    } else {
                        p0 += 285
                        v34 := GenS4{F0: 127, F1: 4294967295, F2: v7}
                        switch "hello" {
                        case "a":
                            if (!true) {
                                var v35 uint32 = 2
                                v8 = append(v8, fi25)
                                _ = v35
                            }
                        default:
                            v36 := GenS2{F0: (p0 << ((18446744073709551615) & 63)), F1: 2}
                            _ = v36
                        }
                        _ = v34
                    }
                }
                _ = fi26
            }
            _ = fi25
        }
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 878
}
func genFn1(p0 uint64) uint {
    var v37 uint32 = 65535
    v38 := GenS5{F0: 1}
    v39 := GenS3{F0: 100}
    _ = p0
    _ = v37
    _ = v38
    _ = v39
    return 255
}

func main() {
    v1 := []string{(("" + ("foo bar" + "hello")) + "foo bar"), ""}
    v2 := make(map[string]int)
    var v3 uint = 65535
    v4 := make(map[uint32]string)
    var v5 int32 = -1
    if false {
        fmt.Println(v3, v5)
        v40 := GenS5{F0: -359}
        _ = len(v2)
        _ = v40
    } else {
        v41 := GenS2{F0: int32(v3), F1: 255}
        _ = v41
    }
    v3 += 255
    v5 += -254
    switch (("a" + "hello") + "hello") {
    case "a":
        var v42 uint = (v3 * (^(v3)))
        _ = v42
    case "hello":
        for fi43 := 0; fi43 < 4; fi43++ {
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            v3 = 42
            _ = fi43
        }
    case "\n\t":
        switch 4294967295 {
        case 42:
            for fi44 := 0; fi44 < 3; fi44++ {
                var v45 uint16 = 1
                _ = fi44
                _ = v45
            }
        default:
            if false {
                fmt.Println(v5)
                for fi46 := 0; fi46 < 3; fi46++ {
                    v5 = v5
                    v47 := GenS5{F0: 1}
                    fmt.Println(v3, v3)
                    _ = fi46
                    _ = v47
                }
                if false {
                    var v48 float64 = 0.1
                    _ = v48
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            } else {
                v3 = v3
                var v49 int8 = -128
                var v50 uint16 = uint16(v3)
                _ = v49
                _ = v50
            }
        }
    default:
        var v51 int32 = v5
        _ = v51
    }
    v3 -= genFn1(1)
    if (!(!(true || (42 > 25)))) {
        v52 := GenS3{F0: -128}
        var v53 uint64 = 18446744073709551615
        _ = v52
        _ = v53
    } else {
        var v54 string = ("a" + "hello")
        _ = v54
    }
    if ((-(v5)) >= (v5 + v5)) {
        for fi55 := 0; fi55 < 2; fi55++ {
            var v56 bool = true
            _ = fi55
            _ = v56
        }
        _ = len(v1)
        v3 -= (v3 / ((42) | 1))
    } else {
        if (!true) {
            if false {
                var v57 int64 = 0
                _ = v57
            }
            v5 = v5
        }
        for fi58 := 0; fi58 < 12; fi58++ {
            v59 := GenS2{F0: (v5 * v5), F1: 238}
            _ = fi58
            _ = v59
        }
    }
    v5 += (v5 / ((v5) | 1))
    var v60 float32 = (0.0 + 1.5)
    if false {
        for fi61 := 0; fi61 < 9; fi61++ {
            if false {
                var v62 int32 = v5
                _ = v62
            } else {
                if true {
                    var v63 int16 = int16(fi61)
                    v64 := GenS6{F0: (1 >= (v3 * (^(v3)))), F1: (v5 | (v5 / ((179) | 1))), F2: -1}
                    if ((("foo bar" + "foo bar") + (("foo bar" + "foo bar") + "foo bar")) >= ("foo bar" + "")) {
                        _ = len(v4)
                        var v65 uint = (v3 * (^(v3)))
                        _ = v65
                    } else {
                        for fi66 := 0; fi66 < 3; fi66++ {
                            if false {
                                var v67 float64 = float64(v60)
                                var v68 uint16 = 790
                                fmt.Println(v63, v3, v5)
                                _ = v67
                                _ = v68
                            }
                            _ = fi66
                        }
                    }
                    _ = v63
                    _ = v64
                } else {
                    var v69 bool = false
                    for fi70 := 0; fi70 < 12; fi70++ {
                        v71 := v4[genFn0((-(v5)), (v69 && v69))]
                        if (genFn0(v5, true) != 2) {
                            switch fi70 {
                            case 671:
                                v5 = -2147483648
                                if (0 < -53) {
                                    v72 := GenS6{F0: (0 == 2), F1: v5, F2: -128}
                                    _ = v72
                                }
                            case 50:
                                var v73 uint32 = 2147483647
                                _ = v73
                            }
                            var v74 bool = v69
                            _ = v74
                        } else {
                            v75 := GenS3{F0: -20}
                            fmt.Println(fi61, fi61)
                            _ = v75
                        }
                        v3 -= (v3 % (((v3 >> (((v3 % ((583) | 1))) & 63))) | 1))
                        _ = fi70
                        _ = v71
                    }
                    for fi76 := 0; fi76 < 15; fi76++ {
                        if v69 {
                            v5 = 42
                            if v69 {
                                delete(v2, "foo bar")
                                switch "\n\t" {
                                case "":
                                    var v77 string = ""
                                    _ = v77
                                case "foo bar":
                                    if (!true) {
                                        if (("hello" + "foo bar") <= "foo bar") {
                                            fmt.Println(v3, v3)
                                            var v78 int16 = 236
                                            _ = v78
                                        }
                                        v79 := GenS7{F0: int16(v3)}
                                        var v80 uint16 = 219
                                        fmt.Println(fi76, fi61, fi76)
                                        _ = v79
                                        _ = v80
                                    }
                                case "a":
                                    var v81 uint8 = 80
                                    v60 = v60
                                    switch "hello" {
                                    case "":
                                        var v82 bool = (1 > 127)
                                        _ = v82
                                    }
                                    _ = v81
                                }
                            }
                            v3 = v3
                        } else {
                            v83 := GenS6{F0: (255 <= 22), F1: v5, F2: 46}
                            var v84 int64 = 48
                            _ = v83
                            _ = v84
                        }
                        if (1 != 114) {
                            v3 += (v3 & v3)
                            if true {
                                var v85 uint = 255
                                if false {
                                    switch ("\n\t" + "a") {
                                    case "hello":
                                        v1 = append(v1, "")
                                    case "foo bar":
                                        var v86 uint32 = 2
                                        _ = v86
                                    case "a":
                                        if ((!((false || false) || (-76 > 100))) || (834 != genFn0(v5, (fi76 == fi61)))) {
                                            fmt.Println(v5, v85, v3)
                                            v3 = v3
                                        } else {
                                            v5 -= 2147483647
                                            if v69 {
                                                switch (fi76 | fi76) {
                                                case 4294967295:
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v5 -= v5
                                                    case "hello":
                                                        v87 := GenS8{F0: 32767, F1: (fi61 | (fi61 * fi61)), F2: v3}
                                                        _ = v87
                                                    default:
                                                        v85 += 18446744073709551615
                                                    }
                                                default:
                                                    var v88 uint32 = 153
                                                    _ = v88
                                                }
                                                v89 := GenS9{F0: 0, F1: 42, F2: 100}
                                                var v90 int = int(v3)
                                                _ = v89
                                                _ = v90
                                            }
                                        }
                                    }
                                    v5 = (-(v5))
                                    switch "foo bar" {
                                    case "foo bar":
                                        fmt.Println(v69, v3)
                                    }
                                    v91 := v4[145]
                                    v69 = false
                                    fmt.Println(fi61)
                                    _ = v91
                                }
                                _ = v85
                            }
                            if v69 {
                                fmt.Println(v3, v3)
                                if true {
                                    v92 := GenS10{F0: 42, F1: 42, F2: 32767}
                                    v3 += (^(v3))
                                    _ = v92
                                }
                                v5 -= (v5 | (v5 & v5))
                            }
                            v93 := GenS9{F0: 876, F1: int16(v5), F2: -1}
                            _ = v93
                        } else {
                            fmt.Println(v3, fi61, v3)
                        }
                        var v94 string = "hello"
                        _ = fi76
                        _ = v94
                    }
                    _ = v69
                }
            }
            _ = fi61
        }
    }
    v95 := GenS5{F0: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v95
}
