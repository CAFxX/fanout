package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 bool; F1 uint32; F2 string }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 bool; F1 bool }
type GenS10 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS11 struct { F0 int32; F1 int16; F2 uint64 }
type GenS12 struct { F0 string; F1 uint8 }
func genFn0(p0 uint32, p1 uint8, p2 int) uint16 {
    fmt.Println(p2, p0)
    v5 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 255
}
func genFn1(p0 int) bool {
    if (true && true) {
        if (42 != -1) {
            v6 := GenS2{F0: 76, F1: genFn0(42, 255, (p0 & 346))}
            fmt.Println(p0)
            _ = v6
        } else {
            v7 := GenS2{F0: 42, F1: 1}
            var v8 string = "\n\t"
            v9 := GenS3{F0: 1}
            _ = v7
            _ = v8
            _ = v9
        }
        var v10 int8 = 0
        _ = v10
    }
    if (!true) {
        var v11 bool = true
        p0 = 2147483647
        _ = v11
    } else {
        var v12 int64 = 0
        var v13 float32 = (0.0 + (0.1 / 1.5))
        _ = v12
        _ = v13
    }
    var v14 int64 = 42
    v15 := GenS2{F0: 109, F1: 32767}
    _ = p0
    _ = v14
    _ = v15
    return true
}

func main() {
    var v1 int8 = -71
    v2 := []int{32767}
    v3 := make(map[string]uint32)
    var v4 uint8 = 1
    for fi16 := 0; fi16 < 12; fi16++ {
        v1 -= v1
        var v17 int32 = -1
        switch "a" {
        case "\n\t":
            var v18 string = "hello"
            _ = v18
        case "foo bar":
            v19 := GenS4{F0: v1}
            _ = v19
        case "hello":
            fmt.Println(fi16)
        default:
            fmt.Println(v4)
        }
        _ = fi16
        _ = v17
    }
    v20 := GenS1{F0: 824}
    fmt.Println(v4)
    v2 = append(v2, 2147483647)
    for fi21 := 0; fi21 < 17; fi21++ {
        delete(v3, "\n\t")
        v22 := GenS3{F0: -1}
        var v23 uint8 = v4
        _ = fi21
        _ = v22
        _ = v23
    }
    switch 61 {
    case -2147483648:
        for fi24 := 0; fi24 < 18; fi24++ {
            if true {
                v1 += v1
                if true {
                    var v25 bool = false
                    if v25 {
                        switch (("hello" + "\n\t") + "hello") {
                        case "hello":
                            for fi26 := 0; fi26 < 1; fi26++ {
                                var v27 int16 = 1000
                                v2 = append(v2, 32768)
                                v28 := GenS2{F0: v1, F1: genFn0(695, (^(v4)), fi24)}
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                        case "a":
                            fmt.Println(v1, fi24, v1)
                        case "\n\t":
                            v29 := GenS4{F0: v1}
                            _ = v29
                        }
                    } else {
                        var v30 uint64 = uint64(v4)
                        _ = v30
                    }
                    _ = v25
                }
                fmt.Println(fi24)
            } else {
                v31 := v3[""]
                var v32 uint8 = uint8(v31)
                _ = v31
                _ = v32
            }
            var v33 uint32 = 2147483647
            _ = fi24
            _ = v33
        }
    }
    v34 := GenS5{F0: 42, F1: 32767}
    v35 := GenS6{F0: (42 != v4), F1: 65535, F2: ("" + "\n\t")}
    switch (("a" + "a") + (("" + (("\n\t" + (("foo bar" + (("\n\t" + "foo bar") + "\n\t")) + ("foo bar" + (((("hello" + ("\n\t" + "a")) + "a") + "hello") + "a")))) + "\n\t")) + "\n\t")) {
    case "\n\t":
        for fi36 := 0; fi36 < 14; fi36++ {
            v37 := GenS7{F0: v4, F1: 1}
            v1 += (v1 & v1)
            _ = fi36
            _ = v37
        }
    case "foo bar":
        var v38 float64 = 1e10
        _ = v38
    }
    if ((!(false && false)) && false) {
        v1 = (v1 / ((v1) | 1))
    }
    var v39 int16 = -166
    v4 = (v4 + v4)
    v1 += v1
    switch -875 {
    case 4294967295:
        var v40 uint64 = 42
        _ = v40
    case -320:
        if false {
            v4 -= 2
            var v41 uint64 = 42
            if (v39 < v39) {
                var v42 int32 = 0
                if true {
                    for fi43 := 0; fi43 < 17; fi43++ {
                        v44 := GenS4{F0: v1}
                        _ = fi43
                        _ = v44
                    }
                    switch -600 {
                    case -398:
                        var v45 string = ""
                        _ = v45
                    default:
                        var v46 uint = 255
                        _ = v46
                    }
                }
                var v47 uint16 = 32767
                _ = v42
                _ = v47
            } else {
                fmt.Println(v41)
                v48 := GenS1{F0: 427}
                v49 := GenS8{F0: 700}
                _ = v48
                _ = v49
            }
            _ = v41
        } else {
            var v50 uint64 = 18446744073709551615
            var v51 uint32 = 171
            _ = v50
            _ = v51
        }
    case 453:
        v3["\n\t"] = 714
    default:
        for fi52 := 0; fi52 < 1; fi52++ {
            v53 := GenS9{F0: true, F1: false}
            switch ("hello" + "") {
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "\n\t":
                if false {
                    var v54 int8 = v1
                    if (true || ("foo bar" <= "foo bar")) {
                        if false {
                            fmt.Println(fi52, v54)
                        }
                        switch ("" + "hello") {
                        case "foo bar":
                            if true {
                                switch (("\n\t" + ("a" + "\n\t")) + "hello") {
                                case "a":
                                    if true {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v55 int = fi52
                                        _ = v55
                                    }
                                case "\n\t":
                                    if len(v2) > 0 {
                                        v2[0] = fi52
                                    }
                                default:
                                    if (!true) {
                                        var v56 uint32 = 2147483647
                                        _ = v56
                                    }
                                }
                            } else {
                                delete(v3, "a")
                            }
                        case "\n\t":
                            switch "hello" {
                            case "":
                                delete(v3, "foo bar")
                            case "hello":
                                v57 := GenS8{F0: 1}
                                _ = v57
                            default:
                                for fi58 := 0; fi58 < 11; fi58++ {
                                    fmt.Println(v1, v39, fi58)
                                    v4 = 43
                                    var v59 int16 = 32767
                                    if ((false && false) && (!true)) {
                                        if (!(fi58 != (fi58 ^ 42))) {
                                            var v60 int = (fi52 % (((-(fi52))) | 1))
                                            if len(v2) > 0 {
                                                v2[0] = int(v39)
                                            }
                                            v54 -= (^(v1))
                                            _ = v60
                                        }
                                        var v61 int64 = -988
                                        v59 = v39
                                        _ = v61
                                    } else {
                                        v39 -= v59
                                    }
                                    var v62 float32 = 1e10
                                    _ = fi58
                                    _ = v59
                                    _ = v62
                                }
                            }
                        case "a":
                            v54 += v1
                        default:
                            for fi63 := 0; fi63 < 10; fi63++ {
                                var v64 uint64 = 18446744073709551615
                                v54 = v1
                                v54 = (v54 / ((v54) | 1))
                                _ = fi63
                                _ = v64
                            }
                        }
                    }
                    fmt.Println(fi52, v39, v1)
                    _ = v54
                } else {
                    fmt.Println(v4, v39, v1)
                }
            case "hello":
                for fi65 := 0; fi65 < 11; fi65++ {
                    switch 0 {
                    case 32767:
                        var v66 int32 = int32(fi52)
                        _ = v66
                    case -1:
                        if (false || (((1.5 == ((float32((float64(((0.0 + 0.1) - ((1.5 / float32(3.14159)) * float32(((float64(0.0) - (float64(1e10) - -2.25)) / 3.14159))))) - 3.14159)) - 0.0) / ((((0.0 - 1e10) - 0.0) + 0.0) * ((((1.5 / 0.1) - 0.1) / ((1e10 * (3.14159 * float32(((0.1 - (0.0 - ((0.1 - 3.14159) * 0.0))) + 0.1)))) / 0.1)) + float32(0.1))))) && ("foo bar" <= "")) || false)) {
                            _ = len(v2)
                        }
                    default:
                        v39 = 1000
                    }
                    _ = fi65
                }
            }
            _ = fi52
            _ = v53
        }
    }
    if false {
        v1 = 100
    }
    v4 += 2
    if (9223372036854775807 < 42) {
        var v67 uint64 = 0
        if (v1 != 1) {
            v1 = -1
            fmt.Println(v1, v67)
            for fi68 := 0; fi68 < 15; fi68++ {
                switch ("" + "a") {
                case "":
                    switch (fi68 % ((128) | 1)) {
                    case -305:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case 401:
                        var v69 float32 = -2.25
                        _ = v69
                    case 247:
                        v70 := GenS10{F0: (v67 ^ (v67 - (v67 / ((v67) | 1)))), F1: 1, F2: -1}
                        _ = v70
                    }
                case "\n\t":
                    if true {
                        if (true && (94 <= (^(v1)))) {
                            for fi71 := 0; fi71 < 11; fi71++ {
                                var v72 string = "hello"
                                fmt.Println(v1, v4)
                                var v73 string = (((("" + "") + (v72 + "")) + ((v72 + (v72 + (("\n\t" + (v72 + v72)) + "a"))) + v72)) + (v72 + "a"))
                                _ = fi71
                                _ = v72
                                _ = v73
                            }
                            var v74 float32 = (1.5 / 0.1)
                            _ = v74
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if (1 >= (v1 / (((v1 & (v1 & v1))) | 1))) {
                                for fi75 := 0; fi75 < 3; fi75++ {
                                    v67 = (v67 - v67)
                                    _ = fi75
                                }
                                v67 = (v67 | (v67 ^ 0))
                            } else {
                                v39 = 0
                                _ = len(v2)
                                fmt.Println(v1, v1)
                            }
                            fmt.Println(v67)
                        }
                        for fi76 := 0; fi76 < 15; fi76++ {
                            v77 := GenS11{F0: 1, F1: 0, F2: v67}
                            switch fi76 {
                            case 587:
                                if true {
                                    v39 += (-(v39))
                                    v67 = v67
                                }
                            default:
                                switch "a" {
                                case "hello":
                                    if len(v2) > 0 {
                                        v2[0] = fi68
                                    }
                                case "":
                                    var v78 int = (fi68 ^ fi76)
                                    _ = v78
                                case "\n\t":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            }
                            _ = fi76
                            _ = v77
                        }
                        v67 += (^(v67))
                    }
                case "a":
                    var v79 int = (fi68 % ((fi68) | 1))
                    _ = v79
                default:
                    v39 += (v39 % ((v39) | 1))
                }
                v4 = 2
                for fi80 := 0; fi80 < 15; fi80++ {
                    var v81 int16 = int16(v4)
                    v82 := GenS5{F0: fi68, F1: genFn0(0, (v4 & v4), (fi80 % (((-(fi80))) | 1)))}
                    var v83 int32 = -1
                    _ = fi80
                    _ = v81
                    _ = v82
                    _ = v83
                }
                _ = fi68
            }
        }
        _ = v67
    } else {
        if true {
            var v84 float64 = float64(3.14159)
            _ = v84
        }
        v85 := GenS12{F0: "hello", F1: 128}
        v1 += (v1 + v1)
        _ = v85
    }
    v3[""] = 335
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v4
}
