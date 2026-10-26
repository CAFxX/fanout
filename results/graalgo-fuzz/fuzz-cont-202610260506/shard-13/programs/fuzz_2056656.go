package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int; F1 int16; F2 int }
type GenS9 struct { F0 int64 }
func genFn0(p0 string, p1 int64) bool {
    v7 := []int{}
    fmt.Println(p1)
    fmt.Println(p1)
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint16 = 0
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return (false || true)
}
func genFn1(p0 uint8) uint8 {
    fmt.Println(p0)
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int = (^(fi10))
        if true {
            fmt.Println(fi10)
            var v12 uint64 = 1
            _ = v12
        } else {
            if false {
                p0 = (p0 / ((p0) | 1))
                if false {
                    p0 += 195
                    p0 -= (p0 + 110)
                }
            }
        }
        p0 += (p0 + p0)
        _ = fi10
        _ = v11
    }
    switch "hello" {
    case "foo bar":
        p0 = 153
    case "":
        p0 = (p0 >> ((560) & 63))
    case "\n\t":
        p0 = p0
    default:
        var v13 int64 = -9223372036854775808
        _ = v13
    }
    _ = p0
    return p0
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{"foo bar", "", "\n\t"}
    v3 := make(map[string]bool)
    var v4 bool = false
    var v5 float32 = 0.1
    var v6 uint64 = 42
    if v4 {
        var v14 float32 = v5
        v5 = float32((-2.25 - (((0.1 / 1.0) / (-2.25 + (0.1 / float64(v5)))) / 0.1)))
        var v15 uint64 = v6
        _ = v14
        _ = v15
    } else {
        delete(v1, 32768)
    }
    var v16 int64 = 0
    if true {
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 string = "hello"
            var v19 uint16 = 2
            _ = fi17
            _ = v18
            _ = v19
        }
    } else {
        var v20 int = -2147483648
        v21 := GenS1{F0: (v16 ^ (v16 % (((v16 / ((v16) | 1))) | 1))), F1: ("hello" + "\n\t"), F2: 826}
        var v22 float64 = -2.25
        _ = v20
        _ = v21
        _ = v22
    }
    fmt.Println(v4)
    _ = len(v2)
    var v23 bool = v4
    fmt.Println(v23, v16)
    delete(v3, ("a" + "hello"))
    for fi24 := 0; fi24 < 19; fi24++ {
        for fi25 := 0; fi25 < 17; fi25++ {
            v16 += (v16 & 559)
            _ = fi25
        }
        var v26 uint8 = 127
        if ((!(!false)) || v4) {
            var v27 string = ""
            switch (v27 + v27) {
            case "a":
                var v28 int32 = 0
                _ = v28
            case "":
                v6 -= uint64(v26)
            case "hello":
                if (!true) {
                    for fi29 := 0; fi29 < 16; fi29++ {
                        switch fi24 {
                        case 0:
                            v4 = v4
                        case 9223372036854775807:
                            fmt.Println(v26, fi24, v23)
                        case 42:
                            v6 = (v6 | (v6 >> ((18446744073709551615) & 63)))
                        }
                        delete(v3, ("\n\t" + ""))
                        v6 = 18446744073709551615
                        switch (fi29 & (fi29 << ((18446744073709551615) & 63))) {
                        case -603:
                            var v30 uint = 1
                            _ = v30
                        case 176:
                            v31 := GenS1{F0: v16, F1: (v27 + v27), F2: v6}
                            v27 = ("hello" + v27)
                            _ = v31
                        default:
                            v32 := GenS2{F0: 65535, F1: 255}
                            _ = v32
                        }
                        var v33 int = (-(fi29))
                        _ = fi29
                        _ = v33
                    }
                    var v34 float64 = (1e10 + (3.14159 + ((3.14159 - ((3.14159 + -2.25) - float64((0.1 / v5)))) / 1.5)))
                    v4 = (false || (255 >= 65535))
                    _ = v34
                } else {
                    fmt.Println(v27)
                }
            default:
                for fi35 := 0; fi35 < 9; fi35++ {
                    if v23 {
                        if v4 {
                            if (!(v26 >= (^(v26)))) {
                                switch "a" {
                                case "a":
                                    v36 := GenS2{F0: 602, F1: 576}
                                    _ = v36
                                case "foo bar":
                                    fmt.Println(v5)
                                case "\n\t":
                                    switch "hello" {
                                    case "":
                                        if ((v16 % ((v16) | 1)) == v16) {
                                            if (304 >= 293) {
                                                if false {
                                                    v27 = "\n\t"
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    if (0.0 <= (1.5 - -2.25)) {
                                                        v6 = 42
                                                        v4 = ((^(v6)) >= v6)
                                                    }
                                                }
                                                var v37 int16 = 0
                                                v5 += v5
                                                _ = v37
                                            }
                                            v38 := GenS1{F0: v16, F1: v27, F2: v6}
                                            switch fi24 {
                                            case 127:
                                                fmt.Println(v5, fi24)
                                            case 939:
                                                v2 = append(v2, v27)
                                            default:
                                                if len(v2) > 0 {
                                                    v2[0] = v27
                                                }
                                            }
                                            _ = v38
                                        }
                                    case "foo bar":
                                        switch v27 {
                                        case "foo bar":
                                            switch v27 {
                                            case "a":
                                                v26 -= 253
                                            case "foo bar":
                                                v39 := GenS2{F0: 65535, F1: 42}
                                                _ = v39
                                            default:
                                                fmt.Println(v27)
                                            }
                                        case "\n\t":
                                            v4 = ((v23 || ((v16 * (v16 ^ v16)) <= v16)) || (549 != v6))
                                        }
                                    case "hello":
                                        _ = len(v2)
                                    default:
                                        var v40 float32 = -2.25
                                        _ = v40
                                    }
                                default:
                                    if (100 > 42) {
                                        fmt.Println(v5, v27, fi24)
                                    } else {
                                        fmt.Println(v23)
                                        v41 := GenS3{F0: 739, F1: 2}
                                        v23 = ((-(fi24)) <= (fi35 + (-(fi35))))
                                        _ = v41
                                    }
                                }
                                if true {
                                    switch v27 {
                                    case "":
                                        var v42 bool = (true && v23)
                                        _ = v42
                                    }
                                    var v43 int16 = 0
                                    var v44 int16 = (^(v43))
                                    _ = v43
                                    _ = v44
                                } else {
                                    if (fi35 < fi35) {
                                        fmt.Println(v6, v6)
                                        fmt.Println(v27, fi24)
                                        delete(v3, ((v27 + v27) + v27))
                                        v27 = v27
                                        if (255 >= 100) {
                                            fmt.Println(v26)
                                            fmt.Println(v6, v4, v4)
                                            fmt.Println(v16)
                                            if (922 == 42) {
                                                switch "a" {
                                                case "":
                                                    v45 := GenS4{F0: 488}
                                                    _ = v45
                                                case "hello":
                                                    v1[(fi24 + (fi24 | (fi24 / (((fi24 >> ((65535) & 63))) | 1))))] = v27
                                                }
                                                _ = len(v3)
                                            } else {
                                                v16 = v16
                                            }
                                            v6 = (v6 * (^(v6)))
                                            var v46 string = "foo bar"
                                            var v47 float32 = v5
                                            if v23 {
                                                var v48 uint = 91
                                                _ = v48
                                            }
                                            _ = v46
                                            _ = v47
                                        } else {
                                            var v49 int8 = 42
                                            fmt.Println(fi35, v27)
                                            _ = v49
                                        }
                                        var v50 bool = v4
                                        delete(v1, -647)
                                        v6 += 1
                                        _ = v50
                                    }
                                }
                                var v51 uint = 4294967295
                                _ = v51
                            } else {
                                var v52 uint8 = (v26 | (v26 ^ (v26 & (v26 + v26))))
                                v5 += -2.25
                                v4 = ((1 <= 2) && (true || true))
                                _ = v52
                            }
                            var v53 uint8 = 2
                            _ = v53
                        }
                        v6 += (v6 ^ (v6 % ((386) | 1)))
                        _ = len(v3)
                    } else {
                        if (v23 && v4) {
                            if v4 {
                                v54 := GenS5{F0: v27, F1: 65535}
                                var v55 uint16 = 768
                                _ = v54
                                _ = v55
                            } else {
                                v56 := GenS5{F0: v27, F1: 245}
                                fmt.Println(fi35)
                                if (true || v4) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if (-1 >= 0) {
                                        switch fi24 {
                                        case 32767:
                                            fmt.Println(v5, v5, v26)
                                        case -133:
                                            if (v27 <= v27) {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                fmt.Println(v5, v6, v27)
                                                fmt.Println(fi24, v26)
                                            } else {
                                                var v57 uint8 = (^(v26))
                                                if true {
                                                    var v58 int = (^(fi24))
                                                    _ = v58
                                                } else {
                                                    v59 := GenS6{F0: 42}
                                                    if (1 <= 32767) {
                                                        v4 = true
                                                    }
                                                    if v4 {
                                                        fmt.Println(v27)
                                                        v27 = "hello"
                                                        var v60 int8 = -128
                                                        _ = v60
                                                    }
                                                    if (!v4) {
                                                        v2 = append(v2, ("foo bar" + v27))
                                                        v57 += v57
                                                        v23 = false
                                                    }
                                                    _ = v59
                                                }
                                                _ = v57
                                            }
                                        }
                                        fmt.Println(v4, v26)
                                    }
                                }
                                var v61 int16 = 1000
                                _ = v56
                                _ = v61
                            }
                        }
                        v4 = (-2.25 != 0.1)
                        v62 := GenS7{F0: 65535}
                        _ = v62
                    }
                    fmt.Println(v27, v16)
                    v16 = v16
                    _ = fi35
                }
            }
            var v63 int = -2147483648
            _ = v27
            _ = v63
        } else {
            fmt.Println(v6, v16)
            v2 = append(v2, "")
            switch "" {
            case "a":
                var v64 uint = 2
                _ = v64
            case "foo bar":
                v65 := GenS8{F0: 32768, F1: 0, F2: (fi24 << ((18446744073709551615) & 63))}
                _ = v65
            case "":
                if true {
                    switch ("foo bar" + ("hello" + "foo bar")) {
                    case "foo bar":
                        v66 := GenS1{F0: (^(v16)), F1: "a", F2: v6}
                        _ = v66
                    case "hello":
                        fmt.Println(v6)
                    default:
                        if v4 {
                            v1[(fi24 % ((fi24) | 1))] = "\n\t"
                            var v67 float64 = (-2.25 / -2.25)
                            _ = v67
                        } else {
                            if (!false) {
                                var v68 float32 = (1.5 / float32(0.1))
                                _ = v68
                            } else {
                                v6 = (v6 & (v6 % ((1) | 1)))
                            }
                            v69 := GenS2{F0: 32767, F1: 104}
                            _ = v69
                        }
                    }
                }
            default:
                v70 := GenS9{F0: (v16 + v16)}
                _ = v70
            }
        }
        _ = fi24
        _ = v26
    }
    fmt.Println(v16, v4)
    for fi71 := 0; fi71 < 18; fi71++ {
        switch "hello" {
        case "foo bar":
            v72 := GenS3{F0: 65535, F1: 32767}
            _ = v72
        }
        _ = fi71
    }
    fmt.Println(v16, v5, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
