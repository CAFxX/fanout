package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint16; F2 string }
type GenS4 struct { F0 string; F1 uint32; F2 int64 }
type GenS5 struct { F0 int32; F1 uint; F2 string }
type GenS6 struct { F0 uint; F1 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int16; F1 int64 }
type GenS9 struct { F0 bool; F1 bool }
type GenS10 struct { F0 uint32; F1 int16; F2 uint32 }
func genFn0(p0 uint16, p1 int, p2 int16) int32 {
    var v6 uint16 = (p0 + (p0 + (p0 | (^(p0)))))
    p2 -= (p2 / ((p2) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -255
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 1
    var v3 int32 = 42
    var v4 uint8 = 127
    v5 := []uint32{42, 4294967295, uint32(v4)}
    if false {
        v5 = append(v5, 42)
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 int16 = (-(v2))
                switch "\n\t" {
                case "a":
                    switch (fi8 % ((42) | 1)) {
                    case -9223372036854775808:
                        v3 -= 0
                    case 128:
                        var v10 uint8 = v4
                        _ = v10
                    case 1:
                        v11 := GenS1{F0: 2}
                        if (202 != 4294967295) {
                            fmt.Println(v3, v9)
                        } else {
                            _ = len(v1)
                            v2 -= (v2 / (((^(v2))) | 1))
                        }
                        _ = v11
                    }
                case "hello":
                    v9 = v2
                    switch fi8 {
                    case 32768:
                        if ((!false) && (1 < 127)) {
                            v12 := GenS1{F0: 1}
                            _ = v12
                        } else {
                            if true {
                                _ = len(v5)
                            }
                            if true {
                                if ("hello" > "") {
                                    var v13 float32 = (-2.25 + 1.5)
                                    v9 += (v2 * (-(v2)))
                                    _ = v13
                                }
                                if (1 <= 42) {
                                    var v14 string = ("foo bar" + "hello")
                                    var v15 float64 = (float64(((0.0 / 1.0) * 1e10)) * float64(3.14159))
                                    var v16 uint = 2
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                } else {
                                    var v17 uint = 1
                                    var v18 float64 = 3.14159
                                    if true {
                                        var v19 bool = false
                                        _ = v19
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    var v20 uint = v17
                                    var v21 uint32 = 564
                                    _ = v17
                                    _ = v18
                                    _ = v20
                                    _ = v21
                                }
                                var v22 uint8 = 255
                                var v23 int16 = v2
                                _ = v22
                                _ = v23
                            } else {
                                if true {
                                    var v24 int = (fi8 << ((1) & 63))
                                    v4 = v4
                                    switch (-(v24)) {
                                    case -1:
                                        v25 := GenS2{F0: v9, F1: 0}
                                        delete(v1, ("\n\t" + "\n\t"))
                                        if true {
                                            var v26 float32 = 1e10
                                            var v27 int = 32767
                                            _ = v26
                                            _ = v27
                                        } else {
                                            switch "hello" {
                                            case "foo bar":
                                                var v28 uint32 = 434
                                                _ = v28
                                            case "hello":
                                                v29 := GenS3{F0: 0, F1: 65535, F2: "a"}
                                                _ = v29
                                            }
                                        }
                                        _ = v25
                                    case 894:
                                        var v30 float32 = -2.25
                                        _ = v30
                                    case 476:
                                        fmt.Println(fi7, fi8)
                                    default:
                                        fmt.Println(v24, v24)
                                    }
                                    _ = v24
                                } else {
                                    v3 = v3
                                    switch (fi7 & (-(fi8))) {
                                    case -9223372036854775808:
                                        var v31 int8 = 0
                                        _ = v31
                                    case 32768:
                                        if true {
                                            v5 = append(v5, 2147483647)
                                        } else {
                                            var v32 uint16 = 2
                                            if false {
                                                if ((((((0.0 / 1e10) - ((float64(3.14159) * float64(0.0)) + 1.5)) + 1e10) / (1.5 * float64(3.14159))) >= -2.25) && false) {
                                                    var v33 bool = false
                                                    v5 = append(v5, 2147483647)
                                                    v2 = 32767
                                                    _ = v33
                                                } else {
                                                    v34 := GenS2{F0: int16(v4), F1: 265}
                                                    _ = v34
                                                }
                                            } else {
                                                var v35 bool = true
                                                v2 += 0
                                                fmt.Println(v2, v4)
                                                _ = v35
                                            }
                                            switch "foo bar" {
                                            case "a":
                                                if false {
                                                    fmt.Println(v3, v32, v32)
                                                    switch (fi8 >> ((255) & 63)) {
                                                    case 127:
                                                        v2 += v9
                                                    case 42:
                                                        if false {
                                                            if true {
                                                                var v36 uint8 = (^(v4))
                                                                var v37 bool = (1000 > 42)
                                                                fmt.Println(v36, v4)
                                                                _ = v36
                                                                _ = v37
                                                            } else {
                                                                v4 -= v4
                                                            }
                                                            fmt.Println(fi7)
                                                            fmt.Println(v4)
                                                        } else {
                                                            v32 = v32
                                                            switch "foo bar" {
                                                            case "a":
                                                                if true {
                                                                    if ((v4 & v4) >= v4) {
                                                                        if true {
                                                                            fmt.Println(fi8, v32)
                                                                            var v38 int32 = (v3 & (v3 | (v3 - 596)))
                                                                            v39 := GenS4{F0: ("hello" + ""), F1: 1, F2: 0}
                                                                            _ = v38
                                                                            _ = v39
                                                                        } else {
                                                                            v40 := GenS5{F0: genFn0((^(v32)), 597, 1), F1: 27, F2: "foo bar"}
                                                                            _ = v40
                                                                        }
                                                                        if true {
                                                                            v2 = v9
                                                                        }
                                                                    }
                                                                    var v41 uint8 = (v4 * (v4 - 1))
                                                                    var v42 bool = true
                                                                    _ = v41
                                                                    _ = v42
                                                                } else {
                                                                    v43 := GenS6{F0: 4294967295, F1: fi8}
                                                                    var v44 int32 = v3
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                            }
                                                            if (42 == 9223372036854775807) {
                                                                _ = len(v1)
                                                                if ((v2 / ((-430) | 1)) != (v2 ^ (v9 - v9))) {
                                                                    var v45 float64 = -2.25
                                                                    _ = v45
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        v46 := GenS7{F0: 2}
                                                        _ = v46
                                                    }
                                                    v3 = 42
                                                }
                                            case "":
                                                v3 -= int32(v4)
                                            }
                                            _ = v32
                                        }
                                    default:
                                        fmt.Println(v4, v9, v2)
                                        _ = len(v1)
                                        v3 = (v3 - (v3 + (v3 % (((v3 ^ (v3 ^ v3))) | 1))))
                                    }
                                    v1["\n\t"] = false
                                }
                                v9 += -214
                                switch "a" {
                                case "\n\t":
                                    v47 := GenS1{F0: 942}
                                    _ = v47
                                case "a":
                                    var v48 uint8 = v4
                                    _ = v48
                                case "":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                default:
                                    fmt.Println(fi7, v2)
                                    var v49 uint16 = 255
                                    _ = v49
                                }
                            }
                            delete(v1, "hello")
                        }
                    }
                    if len(v5) > 0 {
                        v5[0] = 1
                    }
                    if (!(0 >= 2)) {
                        v2 += (v2 + v9)
                    } else {
                        v4 += v4
                        var v50 uint16 = 837
                        v51 := GenS7{F0: 2147483647}
                        if false {
                            v2 = v2
                            var v52 float64 = -2.25
                            _ = v52
                        } else {
                            v53 := GenS8{F0: -2, F1: 220}
                            _ = v53
                        }
                        _ = v50
                        _ = v51
                    }
                }
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        fmt.Println(v2, v2)
    }
    if (42 <= -9223372036854775808) {
        for fi54 := 0; fi54 < 1; fi54++ {
            v55 := GenS4{F0: ("\n\t" + "foo bar"), F1: 400, F2: -912}
            _ = fi54
            _ = v55
        }
    }
    v3 -= v3
    var v56 int32 = v3
    _ = len(v5)
    var v57 int64 = 42
    if (4294967295 != 240) {
        fmt.Println(v4, v57)
        for fi58 := 0; fi58 < 19; fi58++ {
            v3 -= genFn0(65535, fi58, 1)
            delete(v1, "a")
            _ = fi58
        }
        v56 = (v56 / (((v56 | 1)) | 1))
    }
    v59 := v1["\n\t"]
    var v60 int8 = -1
    if v59 {
        v1["foo bar"] = (!(false || (v57 < 345)))
        for fi61 := 0; fi61 < 16; fi61++ {
            v62 := GenS7{F0: 4294967295}
            var v63 bool = (0 == 255)
            var v64 uint16 = 419
            _ = fi61
            _ = v62
            _ = v63
            _ = v64
        }
    } else {
        for fi65 := 0; fi65 < 16; fi65++ {
            v66 := GenS9{F0: v59, F1: false}
            _ = fi65
            _ = v66
        }
        switch 9223372036854775807 {
        case -791:
            v2 += v2
        case 256:
            var v67 float64 = 0.0
            _ = v67
        }
    }
    v68 := GenS1{F0: 644}
    var v69 float64 = -2.25
    for fi70 := 0; fi70 < 1; fi70++ {
        if v59 {
            for fi71 := 0; fi71 < 1; fi71++ {
                if v59 {
                    fmt.Println(v59, v60)
                    v59 = (-25 != v60)
                    v72 := v1[""]
                    _ = v72
                }
                v60 -= -46
                _ = fi71
            }
            _ = len(v1)
            v73 := GenS7{F0: 856}
            _ = v73
        } else {
            var v74 uint = 4294967295
            _ = v74
        }
        for fi75 := 0; fi75 < 7; fi75++ {
            var v76 string = "hello"
            _ = fi75
            _ = v76
        }
        v77 := GenS10{F0: 1, F1: v2, F2: 1}
        _ = fi70
        _ = v77
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v59
    _ = v60
    _ = v68
    _ = v69
}
