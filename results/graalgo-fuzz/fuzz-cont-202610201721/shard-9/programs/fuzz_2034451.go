package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 int; F1 int; F2 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint64; F1 string }
type GenS10 struct { F0 int16; F1 int8 }
type GenS11 struct { F0 bool }
func genFn0(p0 string, p1 uint16, p2 int64) uint64 {
    var v6 uint64 = 42
    v7 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 816
}
func genFn1() uint16 {
    var v8 uint32 = 272
    var v9 int32 = 42
    v10 := GenS1{F0: 42}
    _ = v10
    _ = v8
    _ = v9
    return 255
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    v4 := []string{}
    var v5 float32 = 1e10
    var v11 uint8 = 42
    var v12 uint = 255
    v4 = append(v4, ("a" + (((("hello" + ("a" + ("" + ("foo bar" + ("foo bar" + ("\n\t" + "\n\t")))))) + "a") + "\n\t") + "hello")))
    if true {
        if (v5 >= float32(float64(v5))) {
            var v13 uint16 = 351
            _ = v13
        }
    } else {
        if true {
            v14 := GenS1{F0: 9223372036854775807}
            _ = v14
        } else {
            var v15 int8 = -128
            _ = v15
        }
        delete(v1, 65535)
    }
    if (9223372036854775807 > 0) {
        fmt.Println(v11, v5)
        v12 = (v12 | (^(v12)))
        for fi16 := 0; fi16 < 18; fi16++ {
            if (v5 <= float32(((1e10 - float64(((v5 + v5) - v5))) - (0.0 + 3.14159)))) {
                if ((v11 < v11) && false) {
                    if (("" + "a") != "\n\t") {
                        if true {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v12 += (v12 << ((v12) & 63))
                            switch ("" + ("" + "\n\t")) {
                            case "foo bar":
                                switch fi16 {
                                case 2147483647:
                                    v11 = v11
                                case 4294967295:
                                    v17 := GenS2{F0: (fi16 ^ (^(fi16)))}
                                    _ = v17
                                case 777:
                                    v18 := GenS3{F0: v11, F1: (v11 ^ 0), F2: 0}
                                    _ = v18
                                }
                            case "a":
                                var v19 uint = (v12 ^ (v12 >> ((v12) & 63)))
                                _ = v19
                            default:
                                var v20 string = "hello"
                                _ = v20
                            }
                        }
                        v21 := GenS4{F0: 0, F1: 65535}
                        _ = v21
                    }
                }
                v22 := GenS3{F0: v11, F1: v11, F2: 213}
                _ = v22
            }
            v12 = 902
            var v23 uint = 2
            _ = fi16
            _ = v23
        }
    } else {
        switch 32767 {
        case -2147483648:
            v5 = v5
        case 42:
            v5 = v5
        case 127:
            for fi24 := 0; fi24 < 17; fi24++ {
                v11 += 212
                _ = fi24
            }
        default:
            var v25 int64 = 0
            _ = v25
        }
    }
    v2[65535] = (genFn1() >= 0)
    for fi26 := 0; fi26 < 12; fi26++ {
        for fi27 := 0; fi27 < 15; fi27++ {
            var v28 int64 = -1
            var v29 uint64 = 798
            if ((fi26 - fi26) != -222) {
                if (v11 < v11) {
                    if (979 != 536) {
                        if (false && false) {
                            if true {
                                v30 := GenS1{F0: 42}
                                v31 := GenS4{F0: 65535, F1: 2147483647}
                                _ = v30
                                _ = v31
                            } else {
                                var v32 int32 = 1
                                _ = v32
                            }
                            v5 += float32((((0.1 / 0.1) / 1.5) + -2.25))
                            _ = len(v4)
                            var v33 int = fi27
                            _ = v33
                        }
                        if (-2.25 < 1.5) {
                            v12 = 620
                            var v34 int16 = 1
                            _ = v34
                        }
                        var v35 uint64 = v29
                        fmt.Println(v28)
                        _ = v35
                    }
                    v36 := GenS1{F0: 127}
                    _ = v36
                } else {
                    v37 := GenS5{F0: 42, F1: 2147483647}
                    if ((2 == 652) || true) {
                        v38 := GenS2{F0: (fi26 & (-(fi27)))}
                        _ = v38
                    } else {
                        v28 = v28
                    }
                    v5 -= v5
                    _ = v37
                }
                if (29 <= v11) {
                    var v39 int64 = (-(v28))
                    if false {
                        if true {
                            v40 := GenS3{F0: v11, F1: 255, F2: 656}
                            v41 := v3["foo bar"]
                            if len(v4) > 0 {
                                v4[0] = "\n\t"
                            }
                            _ = v40
                            _ = v41
                        } else {
                            v39 = v39
                            switch (fi27 << (((v12 + 4294967295)) & 63)) {
                            case 127:
                                v42 := GenS6{F0: (fi26 ^ -435), F1: (fi26 * (fi27 & 42)), F2: (v11 | (^(v11)))}
                                _ = v42
                            case 32768:
                                v43 := GenS7{F0: (^(v12))}
                                _ = v43
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                        v44 := GenS4{F0: 255, F1: 65535}
                        _ = v44
                    } else {
                        var v45 int32 = -460
                        switch (fi27 << ((v12) & 63)) {
                        case -768:
                            v11 = v11
                        }
                        switch ((("a" + ("\n\t" + ("a" + "a"))) + "") + "a") {
                        case "hello":
                            if (!(v39 == int64(fi26))) {
                                fmt.Println(v5, v5)
                                var v46 int32 = (v45 & v45)
                                v47 := GenS8{F0: ("foo bar" + "a")}
                                _ = v46
                                _ = v47
                            } else {
                                v29 += v29
                            }
                        case "\n\t":
                            if ((v12 ^ v12) < (v12 >> ((v12) & 63))) {
                                delete(v2, 65535)
                            }
                        case "a":
                            switch (fi27 + (fi27 >> ((v12) & 63))) {
                            case 42:
                                v45 -= v45
                            case 4294967295:
                                v12 = v12
                            default:
                                v48 := GenS6{F0: (fi27 * fi26), F1: (fi26 & (fi26 >> ((559) & 63))), F2: 1}
                                _ = v48
                            }
                        default:
                            fmt.Println(v12, v12, v29)
                        }
                        _ = v45
                    }
                    v49 := GenS9{F0: 0, F1: "\n\t"}
                    _ = v39
                    _ = v49
                } else {
                    v12 = 934
                    v12 -= uint(v11)
                    if len(v4) > 0 {
                        v4[0] = ("\n\t" + "foo bar")
                    }
                }
                var v50 float32 = 3.14159
                _ = v50
            }
            if true {
                var v51 int64 = v28
                var v52 uint = 0
                var v53 bool = ((int32(fi27) > -932) && (42 == 0))
                _ = v51
                _ = v52
                _ = v53
            } else {
                v5 = float32(0.0)
                v54 := GenS10{F0: 228, F1: 42}
                _ = v54
            }
            if false {
                v4 = append(v4, "")
                _ = len(v4)
            } else {
                v5 = v5
                v29 = genFn0("\n\t", uint16(v12), v28)
                v29 = v29
                var v55 uint32 = 2
                v28 -= 553
                fmt.Println(v28)
                if true {
                    v56 := GenS8{F0: "\n\t"}
                    v5 += 1e10
                    _ = v56
                } else {
                    var v57 uint8 = v11
                    switch -1 {
                    case 0:
                        v58 := GenS6{F0: -2147483648, F1: (^(fi27)), F2: uint8(v55)}
                        _ = len(v1)
                        switch int(v28) {
                        case -9223372036854775808:
                            v29 = v29
                        default:
                            v5 -= 3.14159
                        }
                        _ = v58
                    default:
                        var v59 int16 = 1
                        var v60 uint8 = 84
                        var v61 int = fi27
                        _ = v59
                        _ = v60
                        _ = v61
                    }
                    _ = v57
                }
                _ = v55
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        v11 = (v11 & uint8(v12))
        v12 -= (^(v12))
        _ = fi26
    }
    for fi62 := 0; fi62 < 12; fi62++ {
        _ = len(v4)
        v11 = 2
        _ = fi62
    }
    v63 := GenS5{F0: uint8(v12), F1: 1}
    v11 = v11
    if (163 < 303) {
        var v64 uint8 = v11
        switch ("\n\t" + "foo bar") {
        case "hello":
            for fi65 := 0; fi65 < 6; fi65++ {
                v5 += 0.0
                _ = fi65
            }
        case "a":
            v66 := GenS2{F0: 97}
            _ = v66
        default:
            if ("foo bar" != "hello") {
                switch 127 {
                case -670:
                    var v67 float64 = 0.0
                    _ = v67
                case -2147483648:
                    var v68 float32 = (v5 - ((v5 * v5) * (float32(-2.25) + -2.25)))
                    _ = v68
                default:
                    v69 := GenS11{F0: true}
                    _ = v69
                }
            }
        }
        if (!(v11 < v64)) {
            fmt.Println(v11, v11, v64)
        }
        _ = v64
    } else {
        for fi70 := 0; fi70 < 15; fi70++ {
            if false {
                for fi71 := 0; fi71 < 4; fi71++ {
                    var v72 int32 = 2147483647
                    _ = fi71
                    _ = v72
                }
                var v73 string = (("hello" + "\n\t") + ("hello" + ""))
                _ = v73
            }
            var v74 bool = true
            _ = fi70
            _ = v74
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
}
