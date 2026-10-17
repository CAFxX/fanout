package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 bool }
type GenS2 struct { F0 int8; F1 uint; F2 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint; F2 uint64 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 uint; F1 int16 }
func genFn0(p0 uint, p1 uint8) uint64 {
    for fi5 := 0; fi5 < 10; fi5++ {
        if (0 != 2) {
            var v6 int8 = 1
            v6 = (^(v6))
            for fi7 := 0; fi7 < 1; fi7++ {
                p0 += (p0 * (^(p0)))
                _ = fi7
            }
            _ = v6
        }
        var v8 float32 = (3.14159 + 0.1)
        _ = fi5
        _ = v8
    }
    v9 := GenS1{F0: false, F1: (-2147483648 == -1), F2: true}
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 uint64 = 42
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return 18446744073709551615
}
func genFn1(p0 string, p1 string) string {
    var v12 int16 = 32767
    if false {
        fmt.Println(p1)
    }
    var v13 int64 = -1
    p1 = ("a" + "a")
    v14 := []int{-9223372036854775808}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    return ((p1 + (p1 + ((p1 + (((p0 + "foo bar") + p1) + (("\n\t" + p0) + p0))) + (("hello" + p1) + p0)))) + p0)
}

func main() {
    var v1 float64 = (0.0 * -2.25)
    var v2 uint16 = 742
    v3 := []int{-512, 1}
    v4 := []int{int(v2)}
    switch "hello" {
    case "foo bar":
        var v15 uint = 548
        _ = v15
    case "a":
        v16 := GenS1{F0: false, F1: true, F2: (("" + "hello") >= (("hello" + "\n\t") + "hello"))}
        _ = v16
    }
    if false {
        var v17 bool = true
        _ = v17
    } else {
        var v18 bool = (2 != 18446744073709551615)
        if v18 {
            if v18 {
                v19 := GenS1{F0: (-725 >= -913), F1: (18446744073709551615 >= 18446744073709551615), F2: (!v18)}
                var v20 uint = 255
                _ = v19
                _ = v20
            }
            var v21 uint32 = 42
            v2 += v2
            _ = v21
        } else {
            v22 := GenS1{F0: v18, F1: true, F2: true}
            _ = v22
        }
        _ = v18
    }
    var v23 int8 = -56
    fmt.Println(v2)
    for fi24 := 0; fi24 < 9; fi24++ {
        v23 += (v23 + (^(v23)))
        for fi25 := 0; fi25 < 13; fi25++ {
            if true {
                v26 := GenS1{F0: (2147483647 > 42), F1: false, F2: true}
                _ = v26
            } else {
                fmt.Println(v1, v23, v2)
                fmt.Println(v1, v23, v1)
                if false {
                    if (0 >= 2) {
                        if (true || ((0 >= 255) || (4294967295 != 27))) {
                            fmt.Println(fi25, fi24, v2)
                            v1 -= float64(1e10)
                            v27 := GenS1{F0: (v2 <= v2), F1: (!(!(-32768 > 18))), F2: (0 <= 1)}
                            if (v1 == (3.14159 * (float64(0.1) - -2.25))) {
                                if true {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    switch (fi25 - 32768) {
                                    case 2147483647:
                                        if true {
                                            if len(v4) > 0 {
                                                v4[0] = fi24
                                            }
                                            var v28 uint32 = 0
                                            v2 += (^(v2))
                                            if true {
                                                v23 += (-(v23))
                                                _ = len(v3)
                                            }
                                            _ = v28
                                        } else {
                                            v1 = ((v1 + v1) + ((-2.25 / v1) - v1))
                                            v2 -= (v2 ^ (v2 % ((v2) | 1)))
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v29 float32 = (3.14159 / 3.14159)
                                                _ = v29
                                            case "":
                                                var v30 uint16 = 32767
                                                _ = v30
                                            }
                                            v23 += v23
                                            _ = len(v4)
                                        }
                                    case 0:
                                        v31 := GenS2{F0: (v23 | (^(v23))), F1: 42, F2: 541}
                                        _ = v31
                                    case 539:
                                        switch "foo bar" {
                                        case "hello":
                                            if (42 != 222) {
                                                fmt.Println(fi25, v23, fi24)
                                                v1 = 0.1
                                                if false {
                                                    var v32 int64 = 913
                                                    var v33 bool = (true && ((127 != 197) && true))
                                                    _ = v32
                                                    _ = v33
                                                }
                                            } else {
                                                switch ("" + ("foo bar" + "foo bar")) {
                                                case "hello":
                                                    v34 := GenS3{F0: 1000}
                                                    _ = v34
                                                case "\n\t":
                                                    v2 += v2
                                                case "":
                                                    switch (fi24 | (fi25 | 0)) {
                                                    case 503:
                                                        var v35 int64 = 9223372036854775807
                                                        _ = v35
                                                    default:
                                                        v23 = 1
                                                    }
                                                default:
                                                    var v36 int8 = v23
                                                    _ = v36
                                                }
                                                fmt.Println(v2, v1, fi24)
                                            }
                                        }
                                    }
                                } else {
                                    v23 = v23
                                }
                                if (v2 >= 108) {
                                    if (1.5 == 1.5) {
                                        var v37 int8 = v23
                                        v37 += (v37 + v23)
                                        v38 := GenS4{F0: 255, F1: 65535, F2: 588}
                                        _ = v37
                                        _ = v38
                                    } else {
                                        if true {
                                            v1 += float64((((((float32(v1) / 1.5) + 0.0) - 0.1) * 0.1) - 0.0))
                                            switch "\n\t" {
                                            case "\n\t":
                                                switch "a" {
                                                case "\n\t":
                                                    var v39 uint64 = 18446744073709551615
                                                    _ = v39
                                                case "a":
                                                    fmt.Println(v23, v2)
                                                }
                                            default:
                                                switch "foo bar" {
                                                case "hello":
                                                    var v40 float32 = 1e10
                                                    _ = v40
                                                case "foo bar":
                                                    _ = len(v4)
                                                case "\n\t":
                                                    var v41 float32 = 1.5
                                                    _ = v41
                                                default:
                                                    if true {
                                                        var v42 uint = 42
                                                        var v43 string = ("foo bar" + "foo bar")
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                }
                                            }
                                            var v44 float64 = ((v1 + 1e10) - v1)
                                            _ = v44
                                        } else {
                                            v45 := GenS1{F0: (496 <= 2), F1: ((-22 <= (v23 ^ v23)) || false), F2: true}
                                            _ = v45
                                        }
                                        var v46 uint64 = 2
                                        _ = v46
                                    }
                                    v23 = v23
                                    var v47 uint8 = 1
                                    _ = v47
                                } else {
                                    switch 734 {
                                    case 169:
                                        fmt.Println(fi25, v2)
                                    case 9223372036854775807:
                                        var v48 uint16 = v2
                                        _ = v48
                                    case -9223372036854775808:
                                        v1 -= v1
                                    }
                                    var v49 bool = false
                                    _ = v49
                                }
                            } else {
                                var v50 uint16 = (v2 - 42)
                                switch -9223372036854775808 {
                                case 1:
                                    if len(v3) > 0 {
                                        v3[0] = fi24
                                    }
                                case -858:
                                    if false {
                                        var v51 float32 = 1.5
                                        switch "foo bar" {
                                        case "hello":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        v23 += v23
                                        _ = v51
                                    }
                                case -933:
                                    if true {
                                        if (0 > int64(v23)) {
                                            var v52 bool = true
                                            var v53 int = fi25
                                            if v52 {
                                                fmt.Println(v50, fi25, v1)
                                                v54 := GenS5{F0: v23, F1: v23}
                                                if v52 {
                                                    if (1 > int32(v2)) {
                                                        v52 = true
                                                    } else {
                                                        v1 = v1
                                                        v53 -= fi25
                                                        if (!(2 < 127)) {
                                                            v55 := GenS1{F0: v52, F1: (!(!(v23 == -128))), F2: v52}
                                                            _ = v55
                                                        }
                                                        if (!v52) {
                                                            v56 := GenS5{F0: (-(v23)), F1: -128}
                                                            v57 := GenS6{F0: 2147483647, F1: (^(v53))}
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        v2 = 255
                                                    }
                                                } else {
                                                    var v58 uint8 = 42
                                                    var v59 uint8 = (^(v58))
                                                    if (uint64(v58) != 302) {
                                                        if true {
                                                            v1 = v1
                                                            if len(v4) > 0 {
                                                                v4[0] = v53
                                                            }
                                                            v1 -= v1
                                                            v23 -= (v23 ^ v23)
                                                        }
                                                    }
                                                    _ = v58
                                                    _ = v59
                                                }
                                                _ = v54
                                            }
                                            _ = v52
                                            _ = v53
                                        } else {
                                            var v60 int = (-(fi25))
                                            var v61 int = (fi24 + (v60 / (((v60 / ((fi25) | 1))) | 1)))
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = v60
                                            _ = v61
                                        }
                                        var v62 uint8 = uint8(v2)
                                        _ = v62
                                    } else {
                                        v1 = (v1 * 1.5)
                                    }
                                default:
                                    var v63 float32 = (0.0 / float32((v1 + ((((v1 * v1) * 1.5) / v1) * float64(float32(((v1 + -2.25) / float64(-2.25))))))))
                                    _ = v63
                                }
                                if (!false) {
                                    v64 := GenS5{F0: 1, F1: (v23 & 75)}
                                    var v65 uint8 = 95
                                    _ = v64
                                    _ = v65
                                }
                                _ = v50
                            }
                            _ = v27
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if (!false) {
                            var v66 float64 = -2.25
                            if len(v3) > 0 {
                                v3[0] = (fi24 | fi25)
                            }
                            _ = v66
                        } else {
                            var v67 int8 = v23
                            switch (^(fi24)) {
                            case 32767:
                                fmt.Println(fi24, fi24, v1)
                            default:
                                v2 -= (v2 | (v2 + (v2 | 42)))
                            }
                            _ = v67
                        }
                    } else {
                        _ = len(v3)
                    }
                    fmt.Println(v23)
                }
            }
            _ = fi25
        }
        _ = fi24
    }
    if (1e10 >= 0.0) {
        v1 = 0.1
    }
    var v68 float32 = 1.5
    v69 := GenS2{F0: v23, F1: 51, F2: 42}
    for fi70 := 0; fi70 < 14; fi70++ {
        for fi71 := 0; fi71 < 18; fi71++ {
            var v72 float64 = (0.0 * 0.1)
            _ = fi71
            _ = v72
        }
        v1 = ((v1 + v1) * v1)
        _ = fi70
    }
    if false {
        v68 -= (v68 + v68)
    } else {
        var v73 int8 = v23
        for fi74 := 0; fi74 < 18; fi74++ {
            for fi75 := 0; fi75 < 14; fi75++ {
                switch "a" {
                case "foo bar":
                    switch (fi75 >> ((1) & 63)) {
                    case 4294967295:
                        var v76 uint = 2
                        _ = v76
                    case 32768:
                        v2 = v2
                    case -886:
                        v73 = (v73 >> ((197) & 63))
                    }
                case "\n\t":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                v77 := GenS1{F0: (false || false), F1: (true || false), F2: (65535 != 2147483648)}
                _ = fi75
                _ = v77
            }
            _ = fi74
        }
        var v78 float32 = float32(-2.25)
        _ = v73
        _ = v78
    }
    v79 := GenS1{F0: false, F1: ((592 >= 250) && (!(0 <= 321))), F2: (!false)}
    var v80 int = 740
    var v81 uint = 1
    for fi82 := 0; fi82 < 15; fi82++ {
        v83 := GenS4{F0: (v81 / (((v81 / ((v81) | 1))) | 1)), F1: v81, F2: 0}
        fmt.Println(v80, v81)
        _ = fi82
        _ = v83
    }
    v84 := GenS7{F0: (v81 % ((v81) | 1)), F1: -32768}
    fmt.Println(v80)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v68:", v68)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v68
    _ = v69
    _ = v79
    _ = v80
    _ = v81
    _ = v84
}
