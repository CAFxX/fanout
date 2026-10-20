package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint16; F1 int32; F2 int64 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 int16; F1 int; F2 int32 }
func genFn0(p0 uint, p1 uint32) bool {
    p1 += (p1 - uint32(p0))
    v6 := GenS1{F0: (p0 ^ p0), F1: 1000}
    _ = p0
    _ = p1
    _ = v6
    return true
}
func genFn1() uint {
    var v7 string = "a"
    fmt.Println(v7)
    if true {
        fmt.Println(v7)
        if (((0.0 + 1e10) / 3.14159) == (1e10 + (((1e10 + (-2.25 + ((3.14159 + -2.25) / (0.1 / 1.0)))) / -2.25) + (1e10 + 0.1)))) {
            if (3.14159 <= float32(0.1)) {
                var v8 int8 = -106
                _ = v8
            } else {
                for fi9 := 0; fi9 < 8; fi9++ {
                    var v10 int = (-(fi9))
                    _ = fi9
                    _ = v10
                }
                fmt.Println(v7)
                for fi11 := 0; fi11 < 2; fi11++ {
                    var v12 uint = 2
                    fmt.Println(v7, v7, fi11)
                    _ = fi11
                    _ = v12
                }
            }
            v7 = "hello"
            var v13 int64 = 128
            _ = v13
        } else {
            var v14 bool = false
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 int = (fi15 | fi15)
                _ = fi15
                _ = v16
            }
            switch 42 {
            case 29:
                v14 = true
            case -516:
                var v17 bool = (!(-277 >= -1))
                _ = v17
            case -218:
                fmt.Println(v7, v14)
            }
            _ = v14
        }
        if (128 != 128) {
            v7 = (v7 + (v7 + v7))
            v7 = v7
        } else {
            if ((v7 + (v7 + v7)) >= (v7 + v7)) {
                v18 := GenS1{F0: 255, F1: 0}
                v19 := GenS1{F0: 42, F1: -1}
                v7 = "foo bar"
                _ = v18
                _ = v19
            } else {
                var v20 uint32 = 341
                v20 += (v20 * 65535)
                for fi21 := 0; fi21 < 7; fi21++ {
                    v22 := []int{(fi21 + fi21)}
                    if (false && false) {
                        var v23 uint16 = 0
                        for fi24 := 0; fi24 < 16; fi24++ {
                            if len(v22) > 0 {
                                v22[0] = 127
                            }
                            var v25 int = fi21
                            _ = fi24
                            _ = v25
                        }
                        _ = v23
                    } else {
                        for fi26 := 0; fi26 < 13; fi26++ {
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                            switch (fi26 >> ((1) & 63)) {
                            case 752:
                                v27 := GenS1{F0: 495, F1: -32768}
                                _ = v27
                            case 42:
                                var v28 int32 = 2147483647
                                _ = v28
                            default:
                                fmt.Println(fi21)
                            }
                            _ = fi26
                        }
                        var v29 int = -1
                        if len(v22) > 0 {
                            v22[0] = fi21
                        }
                        _ = v29
                    }
                    _ = fi21
                    _ = v22
                }
                _ = v20
            }
            v7 = v7
        }
    }
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 9223372036854775807
    var v3 uint16 = 53
    v4 := []int{v2, v2, v2}
    var v5 int = (v2 / ((v2) | 1))
    var v30 uint64 = 398
    v31 := GenS1{F0: 457, F1: 1000}
    for fi32 := 0; fi32 < 16; fi32++ {
        switch ("hello" + ("" + "hello")) {
        case "a":
            var v33 int8 = -1
            _ = v33
        case "hello":
            var v34 int32 = -655
            _ = v34
        default:
            for fi35 := 0; fi35 < 15; fi35++ {
                var v36 int32 = 628
                v37 := GenS2{F0: v3, F1: (-(v36)), F2: 997}
                fmt.Println(v36, v36)
                _ = fi35
                _ = v36
                _ = v37
            }
        }
        var v38 uint32 = uint32(v30)
        _ = fi32
        _ = v38
    }
    var v39 uint8 = 42
    delete(v1, "")
    if (!((v30 - (v30 * (v30 + (^(v30))))) == (^(v30)))) {
        switch (("foo bar" + ("hello" + ((("" + (("foo bar" + "hello") + ("a" + ("foo bar" + ("hello" + "a"))))) + ("\n\t" + ("\n\t" + ""))) + ""))) + ("a" + "")) {
        case "a":
            v40 := GenS2{F0: v3, F1: -2147483648, F2: 0}
            _ = v40
        case "\n\t":
            switch v2 {
            case -9223372036854775808:
                var v41 int8 = 100
                _ = v41
            }
        case "":
            if true {
                fmt.Println(v2, v39, v3)
                v39 = 0
                var v42 int8 = -128
                _ = v42
            }
        }
    } else {
        v2 = v5
    }
    _ = len(v1)
    var v43 string = ("" + "foo bar")
    if true {
        for fi44 := 0; fi44 < 14; fi44++ {
            var v45 uint = 737
            if (1 <= 2) {
                if false {
                    if false {
                        if true {
                            v2 = fi44
                        } else {
                            var v46 int64 = 42
                            v47 := GenS3{F0: v46, F1: (^(v30))}
                            v48 := GenS4{F0: v30, F1: v43}
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                    } else {
                        var v49 float64 = float64((-2.25 - ((1.5 + -2.25) - -2.25)))
                        if (42 > 16) {
                            switch (v2 * 187) {
                            case 32767:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case 871:
                                var v50 bool = false
                                _ = v50
                            case 32768:
                                switch 2147483647 {
                                case 1:
                                    for fi51 := 0; fi51 < 20; fi51++ {
                                        v39 -= 1
                                        _ = fi51
                                    }
                                case 32767:
                                    switch v43 {
                                    case "hello":
                                        fmt.Println(v3, fi44, v2)
                                    }
                                default:
                                    switch v2 {
                                    case 127:
                                        var v52 int = (-(fi44))
                                        _ = v52
                                    case 2147483647:
                                        if false {
                                            var v53 float64 = v49
                                            _ = v53
                                        }
                                    default:
                                        fmt.Println(v49, v39, fi44)
                                    }
                                }
                            default:
                                v49 = v49
                            }
                            var v54 float64 = v49
                            var v55 int = (v2 << ((4294967295) & 63))
                            _ = v54
                            _ = v55
                        }
                        v56 := GenS4{F0: (v30 | v30), F1: v43}
                        _ = v49
                        _ = v56
                    }
                } else {
                    var v57 string = v43
                    _ = v57
                }
                if (0.0 != 3.14159) {
                    for fi58 := 0; fi58 < 9; fi58++ {
                        v59 := GenS4{F0: (v30 ^ v30), F1: ((v43 + v43) + ("foo bar" + "hello"))}
                        _ = fi58
                        _ = v59
                    }
                    delete(v1, ("a" + ((v43 + "foo bar") + v43)))
                    fmt.Println(v39)
                }
                for fi60 := 0; fi60 < 10; fi60++ {
                    v45 -= 284
                    var v61 uint16 = 316
                    fmt.Println(fi44)
                    v3 = 65535
                    switch ((v43 + v43) + "") {
                    case "":
                        v62 := GenS5{F0: -1, F1: v2, F2: 42}
                        v5 = (v2 >> ((1) & 63))
                        _ = v62
                    }
                    _ = fi60
                    _ = v61
                }
            }
            _ = fi44
            _ = v45
        }
        v3 -= 299
    } else {
        var v63 uint16 = v3
        _ = v63
    }
    var v64 uint64 = v30
    if (42 <= 9223372036854775807) {
        if true {
            fmt.Println(v39, v2, v64)
            var v65 int32 = -1
            var v66 uint16 = (^(v3))
            _ = v65
            _ = v66
        } else {
            fmt.Println(v3)
            v67 := GenS2{F0: uint16(v39), F1: -821, F2: 9223372036854775807}
            _ = v67
        }
        var v68 float32 = (float32(0.1) * 0.0)
        _ = v68
    }
    var v69 uint = 1
    for fi70 := 0; fi70 < 17; fi70++ {
        for fi71 := 0; fi71 < 11; fi71++ {
            v43 = (v43 + "hello")
            switch fi71 {
            case -815:
                v64 += (v30 << ((0) & 63))
            case 32768:
                var v72 uint64 = (v64 * 0)
                _ = v72
            case 2147483647:
                var v73 uint8 = v39
                _ = v73
            default:
                var v74 int32 = -35
                _ = v74
            }
            _ = fi71
        }
        _ = fi70
    }
    var v75 float64 = 3.14159
    switch "" {
    case "hello":
        fmt.Println(v39, v64, v69)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v69:", v69)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v43
    _ = v5
    _ = v64
    _ = v69
    _ = v75
}
