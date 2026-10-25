package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 bool; F1 int; F2 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 int16; F1 int32; F2 bool }
type GenS10 struct { F0 int16; F1 uint32; F2 int64 }
type GenS11 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 1
    v4 := []int{-9223372036854775808}
    switch "hello" {
    case "a":
        switch 787 {
        case -1:
            v3 = v3
        case 4294967295:
            if (!false) {
                v3 -= v3
                v5 := GenS1{F0: 42, F1: "a"}
                v1[65535] = (("a" + "hello") + "")
                _ = v5
            }
        }
    case "hello":
        if true {
            v3 += (v3 + v3)
        }
    case "foo bar":
        switch "\n\t" {
        case "\n\t":
            v3 = 0
        default:
            for fi6 := 0; fi6 < 17; fi6++ {
                v7 := v1[65535]
                var v8 int16 = -359
                _ = fi6
                _ = v7
                _ = v8
            }
        }
    default:
        var v9 int32 = -5
        _ = v9
    }
    var v10 int64 = -1
    if (65535 == 32767) {
        fmt.Println(v3, v3)
        delete(v2, 632)
        if true {
            if true {
                fmt.Println(v10)
                var v11 string = "\n\t"
                _ = v11
            } else {
                var v12 string = ""
                _ = v12
            }
            switch 42 {
            case 4294967295:
                v13 := GenS1{F0: -2147483648, F1: "foo bar"}
                _ = v13
            default:
                if (false || false) {
                    var v14 uint = 1
                    v3 += v3
                    _ = v14
                } else {
                    v10 = v10
                    var v15 int64 = -237
                    v15 -= v15
                    _ = v15
                }
            }
            v3 = v3
        } else {
            for fi16 := 0; fi16 < 2; fi16++ {
                if (945 > 862) {
                    for fi17 := 0; fi17 < 8; fi17++ {
                        fmt.Println(v10)
                        _ = fi17
                    }
                    var v18 int8 = -24
                    _ = v18
                }
                if (42 > 893) {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        if false {
                            v3 = (v3 / (((v3 >> ((18446744073709551615) & 63))) | 1))
                        }
                        var v20 int32 = 206
                        _ = fi19
                        _ = v20
                    }
                    v10 += v10
                }
                var v21 string = "foo bar"
                _ = fi16
                _ = v21
            }
        }
    } else {
        if false {
            if (v3 != (^(v3))) {
                if (true || false) {
                    v10 = v10
                }
            }
            v22 := GenS1{F0: 756, F1: ""}
            _ = v22
        } else {
            fmt.Println(v3, v10)
            var v23 bool = (255 > 128)
            _ = v23
        }
    }
    var v24 uint8 = v3
    if len(v4) > 0 {
        v4[0] = 32768
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        for fi26 := 0; fi26 < 15; fi26++ {
            var v27 float64 = ((0.1 * -2.25) + 0.1)
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    switch "foo bar" {
    case "\n\t":
        for fi28 := 0; fi28 < 18; fi28++ {
            v10 -= (v10 + v10)
            _ = fi28
        }
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        if (255 <= 32767) {
            _ = len(v2)
            v4 = append(v4, (fi29 % ((42) | 1)))
            fmt.Println(v3)
        }
        if len(v4) > 0 {
            v4[0] = 127
        }
        _ = fi29
    }
    delete(v1, 42)
    if len(v4) > 0 {
        v4[0] = -257
    }
    if ((!false) && ((((1e10 * ((float32(1e10) * (0.0 - float32((float64(0.1) / (0.1 / (float64(((-2.25 / 1.0) * 1e10)) * (1.5 * 1e10))))))) * float32(float64(1.5)))) + 3.14159) * 0.1) != (0.1 + (float32(-2.25) * 3.14159)))) {
        v30 := GenS1{F0: 0, F1: "hello"}
        for fi31 := 0; fi31 < 7; fi31++ {
            if len(v4) > 0 {
                v4[0] = (fi31 | (fi31 / ((fi31) | 1)))
            }
            if true {
                v32 := GenS2{F0: ((!false) && (18446744073709551615 <= 2)), F1: (fi31 ^ (fi31 / (((fi31 >> ((1) & 63))) | 1))), F2: "hello"}
                if true {
                    fmt.Println(fi31)
                    switch "hello" {
                    case "\n\t":
                        v33 := GenS2{F0: true, F1: fi31, F2: "a"}
                        _ = v33
                    case "a":
                        v3 = (v24 % (((^(v3))) | 1))
                    default:
                        var v34 uint16 = 118
                        _ = v34
                    }
                    fmt.Println(v10)
                } else {
                    if ((true || true) && ("" < "foo bar")) {
                        v35 := GenS3{F0: -1}
                        switch 32767 {
                        case 155:
                            _ = len(v4)
                        }
                        _ = v35
                    }
                }
                _ = v32
            }
            for fi36 := 0; fi36 < 2; fi36++ {
                v37 := GenS4{F0: 1000, F1: v3, F2: 223}
                _ = fi36
                _ = v37
            }
            _ = fi31
        }
        _ = v30
    }
    for fi38 := 0; fi38 < 5; fi38++ {
        if true {
            v24 = 127
            v39 := GenS5{F0: -1}
            _ = v39
        } else {
            v24 = (^(v3))
            for fi40 := 0; fi40 < 8; fi40++ {
                v41 := GenS4{F0: 302, F1: 20, F2: 65535}
                var v42 int64 = 0
                if true {
                    v43 := GenS6{F0: 1}
                    v44 := GenS1{F0: 703, F1: ("hello" + "\n\t")}
                    switch "foo bar" {
                    case "\n\t":
                        v45 := GenS2{F0: false, F1: fi38, F2: ("a" + "")}
                        _ = v45
                    case "a":
                        v24 = v24
                    }
                    _ = v43
                    _ = v44
                } else {
                    v24 -= (v3 & 201)
                    v46 := GenS5{F0: -32768}
                    _ = v46
                }
                _ = fi40
                _ = v41
                _ = v42
            }
            if false {
                v10 = (v10 << ((4294967295) & 63))
                if true {
                    if false {
                        var v47 uint = 805
                        v47 = (v47 * v47)
                        v10 = v10
                        _ = v47
                    } else {
                        if (2 > 310) {
                            delete(v2, 636)
                            v24 = (v3 - (v24 & v3))
                        } else {
                            v4 = append(v4, fi38)
                            if (1 != 672) {
                                v10 = v10
                                var v48 float64 = 1.5
                                if (1 != -1) {
                                    if (!true) {
                                        if true {
                                            v3 = 0
                                            if true {
                                                var v49 string = "hello"
                                                fmt.Println(v49, v49, fi38)
                                                _ = v49
                                            }
                                        } else {
                                            switch fi38 {
                                            case 408:
                                                if false {
                                                    var v50 string = ((("foo bar" + "foo bar") + (("" + ("a" + (("a" + "a") + "a"))) + ("hello" + ("\n\t" + ("" + (("\n\t" + "\n\t") + (("foo bar" + "") + (("a" + "") + (("foo bar" + "") + "hello"))))))))) + "\n\t")
                                                    for fi51 := 0; fi51 < 3; fi51++ {
                                                        var v52 uint8 = (v3 >> ((71) & 63))
                                                        v53 := GenS7{F0: v50}
                                                        if len(v4) > 0 {
                                                            v4[0] = fi38
                                                        }
                                                        _ = fi51
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                    _ = v50
                                                }
                                            case 9223372036854775807:
                                                var v54 uint64 = 18446744073709551615
                                                _ = v54
                                            default:
                                                v48 = (float64((1e10 - 1.5)) * v48)
                                            }
                                            var v55 uint64 = 42
                                            v24 -= v24
                                            _ = v55
                                        }
                                        v10 += v10
                                        fmt.Println(v10, fi38, v24)
                                    } else {
                                        v56 := GenS1{F0: -2147483648, F1: "\n\t"}
                                        fmt.Println(fi38)
                                        if (false && (((fi38 >> ((491) & 63)) != -48) || false)) {
                                            v48 = ((v48 - v48) * (float64(3.14159) * (v48 + v48)))
                                        } else {
                                            for fi57 := 0; fi57 < 10; fi57++ {
                                                fmt.Println(v10, v10, v3)
                                                fmt.Println(fi38, v48)
                                                _ = fi57
                                            }
                                        }
                                        _ = v56
                                    }
                                    if (-1 >= 42) {
                                        var v58 uint8 = v3
                                        switch fi38 {
                                        case 32768:
                                            v59 := GenS5{F0: 0}
                                            _ = v59
                                        case 40:
                                            if ((true || (-2147483648 < 7)) || true) {
                                                var v60 uint64 = 211
                                                var v61 uint64 = v60
                                                _ = v60
                                                _ = v61
                                            }
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            if (2147483647 > 1) {
                                                if true {
                                                    v62 := GenS8{F0: (1 < 2147483648), F1: 0}
                                                    _ = v62
                                                } else {
                                                    if (fi38 > 42) {
                                                        var v63 int16 = 32767
                                                        _ = v63
                                                    }
                                                    v64 := GenS9{F0: int16(fi38), F1: -341, F2: false}
                                                    _ = v64
                                                }
                                            } else {
                                                _ = len(v4)
                                                var v65 uint32 = 2147483647
                                                _ = v65
                                            }
                                        case "":
                                            fmt.Println(v24, v48)
                                        case "\n\t":
                                            v3 = v3
                                        default:
                                            v66 := GenS10{F0: -18, F1: 0, F2: (v10 | 0)}
                                            _ = v66
                                        }
                                        _ = v58
                                    } else {
                                        var v67 int8 = -41
                                        _ = len(v1)
                                        v48 = 0.1
                                        _ = v67
                                    }
                                } else {
                                    var v68 int = (fi38 / ((fi38) | 1))
                                    _ = v68
                                }
                                _ = v48
                            }
                            _ = len(v4)
                        }
                        switch "foo bar" {
                        case "\n\t":
                            for fi69 := 0; fi69 < 13; fi69++ {
                                fmt.Println(v3, fi38, fi38)
                                _ = fi69
                            }
                        case "":
                            switch (fi38 << ((255) & 63)) {
                            case 4294967295:
                                fmt.Println(v3)
                            case 42:
                                var v70 uint = 0
                                _ = v70
                            default:
                                v24 = v3
                            }
                        }
                        var v71 uint16 = 0
                        _ = v71
                    }
                    v72 := GenS2{F0: true, F1: fi38, F2: (("a" + "foo bar") + (("" + ("" + (("hello" + ("\n\t" + (("a" + "foo bar") + "foo bar"))) + "a"))) + "hello"))}
                    if (2 == 65535) {
                        var v73 uint64 = 620
                        if false {
                            for fi74 := 0; fi74 < 2; fi74++ {
                                v75 := GenS9{F0: 42, F1: -1, F2: false}
                                _ = fi74
                                _ = v75
                            }
                        } else {
                            var v76 int = (fi38 << ((uint(v24)) & 63))
                            if (true || true) {
                                v76 -= fi38
                            } else {
                                v77 := GenS9{F0: 1, F1: 2147483647, F2: (fi38 < (fi38 / ((v76) | 1)))}
                                _ = v77
                            }
                            _ = v76
                        }
                        var v78 uint16 = 2
                        _ = v73
                        _ = v78
                    }
                    _ = v72
                }
            }
        }
        v24 += (v24 & (v3 / ((v3) | 1)))
        _ = fi38
    }
    _ = len(v2)
    switch -401 {
    case 32768:
        v10 -= (-(v10))
    case 128:
        var v79 int = -1
        _ = v79
    case 4294967295:
        v80 := GenS11{F0: -978}
        _ = v80
    }
    v81 := GenS5{F0: 42}
    var v82 int32 = int32(v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v81
    _ = v82
}
