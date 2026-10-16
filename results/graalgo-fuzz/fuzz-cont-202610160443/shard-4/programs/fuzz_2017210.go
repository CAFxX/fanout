package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 int32; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 uint64; F2 bool }
type GenS6 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS7 struct { F0 uint }

func main() {
    var v1 int32 = -431
    v2 := []string{"", "a", ("" + "a"), "foo bar"}
    var v3 float32 = -2.25
    v4 := make(map[int]string)
    var v5 string = ("foo bar" + "\n\t")
    if (42 < 0) {
        if false {
            v6 := v4[42]
            v3 = float32(1.5)
            v4[-2147483648] = v6
            _ = v6
        }
        var v7 int64 = -1
        _ = v7
    } else {
        for fi8 := 0; fi8 < 8; fi8++ {
            v3 = v3
            _ = fi8
        }
    }
    switch 534 {
    case -1:
        switch 2147483647 {
        case 388:
            delete(v4, 957)
        case 9223372036854775807:
            v1 += (v1 / ((v1) | 1))
        }
    case 42:
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS1{F0: (v1 & v1)}
            switch (fi9 | fi9) {
            case -9223372036854775808:
                for fi11 := 0; fi11 < 3; fi11++ {
                    var v12 uint64 = 552
                    v12 += v12
                    _ = fi11
                    _ = v12
                }
            case 612:
                v2 = append(v2, (v5 + "foo bar"))
            case -2147483648:
                for fi13 := 0; fi13 < 5; fi13++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if false {
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                        var v14 uint8 = 46
                        _ = v14
                    } else {
                        if false {
                            var v15 int = (fi9 << ((42) & 63))
                            v16 := GenS1{F0: (v1 / ((int32(fi9)) | 1))}
                            fmt.Println(v5)
                            fmt.Println(v1, fi9, v15)
                            v17 := GenS1{F0: v1}
                            v2 = append(v2, "foo bar")
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                    }
                    _ = fi13
                }
            }
            _ = fi9
            _ = v10
        }
    case 0:
        _ = len(v2)
    }
    var v18 uint32 = 0
    v19 := v4[42]
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 16; fi21++ {
            v22 := GenS1{F0: v1}
            if ((true || (true && false)) || (!(!(!true)))) {
                var v23 int16 = 32767
                _ = v23
            } else {
                v24 := GenS1{F0: (v1 + (v1 % (((v1 ^ v1)) | 1)))}
                var v25 float32 = float32(0.1)
                _ = v24
                _ = v25
            }
            if false {
                if (false && (!(!false))) {
                    if false {
                        v26 := GenS1{F0: (v1 >> ((4294967295) & 63))}
                        switch ((((v5 + "a") + "hello") + "\n\t") + (v5 + (v5 + "a"))) {
                        case "":
                            v18 += v18
                        case "foo bar":
                            var v27 uint64 = 768
                            _ = v27
                        case "hello":
                            if true {
                                var v28 uint8 = 127
                                var v29 float64 = 0.0
                                v29 -= v29
                                _ = v28
                                _ = v29
                            } else {
                                switch v5 {
                                case "":
                                    var v30 string = (("" + "a") + v5)
                                    _ = v30
                                case "a":
                                    v2 = append(v2, "foo bar")
                                case "foo bar":
                                    if true {
                                        if (2 != 4294967295) {
                                            v31 := GenS1{F0: (v1 / ((2147483647) | 1))}
                                            switch fi20 {
                                            case 1:
                                                if (178 > -229) {
                                                    if (0 == 2) {
                                                        v2 = append(v2, (v19 + v5))
                                                        v18 = v18
                                                        v19 = v5
                                                        var v32 uint16 = 65535
                                                        v1 += 42
                                                        _ = v32
                                                    }
                                                    fmt.Println(fi20, v18, fi20)
                                                    v19 = "foo bar"
                                                } else {
                                                    if false {
                                                        v33 := GenS2{F0: (fi20 == (^(fi21))), F1: (v1 + 1), F2: 42}
                                                        switch v5 {
                                                        case "hello":
                                                            fmt.Println(fi20)
                                                        case "":
                                                            if true {
                                                                v19 = v5
                                                            } else {
                                                                v34 := GenS2{F0: false, F1: v1, F2: (v18 | (v18 ^ (v18 >> ((255) & 63))))}
                                                                fmt.Println(v19, v1, fi21)
                                                                _ = v34
                                                            }
                                                        }
                                                        _ = v33
                                                    }
                                                    switch fi20 {
                                                    case -466:
                                                        var v35 int8 = 57
                                                        fmt.Println(v35)
                                                        _ = v35
                                                    case -2147483648:
                                                        if true {
                                                            if ((v18 - 1) > v18) {
                                                                switch 2147483647 {
                                                                case -9223372036854775808:
                                                                    v18 = v18
                                                                }
                                                                v19 = v5
                                                            } else {
                                                                fmt.Println(v3, fi20)
                                                                var v36 uint64 = 42
                                                                v1 -= v1
                                                                _ = v36
                                                            }
                                                            var v37 float32 = (float32((0.0 / float64(float32(-2.25)))) - 0.1)
                                                            v1 -= (-(v1))
                                                            _ = v37
                                                        }
                                                    }
                                                    v5 = v5
                                                }
                                            }
                                            var v38 int8 = -128
                                            switch fi20 {
                                            case 127:
                                                v39 := GenS1{F0: (^(v1))}
                                                _ = v39
                                            case 1:
                                                var v40 uint64 = 685
                                                _ = v40
                                            case 726:
                                                var v41 int32 = v1
                                                _ = v41
                                            default:
                                                switch v5 {
                                                case "\n\t":
                                                    if true {
                                                        v42 := GenS3{F0: v38, F1: 1, F2: 2}
                                                        _ = v42
                                                    } else {
                                                        if false {
                                                            var v43 uint16 = 32767
                                                            var v44 uint16 = v43
                                                            _ = v43
                                                            _ = v44
                                                        } else {
                                                            var v45 bool = true
                                                            _ = v45
                                                        }
                                                    }
                                                case "":
                                                    if false {
                                                        v46 := GenS4{F0: 2147483647}
                                                        _ = v46
                                                    }
                                                default:
                                                    if (!false) {
                                                        var v47 uint32 = v18
                                                        v4[fi21] = "hello"
                                                        v18 -= (v47 >> ((0) & 63))
                                                        _ = v47
                                                    } else {
                                                        v2 = append(v2, "\n\t")
                                                        delete(v4, 32767)
                                                        v3 += (v3 / v3)
                                                    }
                                                }
                                                v4[fi21] = v5
                                                if (!true) {
                                                    v48 := GenS4{F0: (fi20 << ((18446744073709551615) & 63))}
                                                    delete(v4, -67)
                                                    _ = v48
                                                } else {
                                                    var v49 int32 = (v1 | v1)
                                                    _ = v49
                                                }
                                            }
                                            _ = v31
                                            _ = v38
                                        }
                                        v18 -= 65535
                                    }
                                }
                                v1 = 42
                            }
                        default:
                            if false {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v50 string = ("hello" + v5)
                                v51 := v4[(-(fi20))]
                                _ = v50
                                _ = v51
                            } else {
                                switch "\n\t" {
                                case "a":
                                    fmt.Println(fi20, v1)
                                case "foo bar":
                                    var v52 int32 = 42
                                    _ = v52
                                case "\n\t":
                                    if true {
                                        fmt.Println(fi21, fi21)
                                        var v53 int8 = 1
                                        _ = v53
                                    } else {
                                        var v54 int8 = 42
                                        _ = v54
                                    }
                                default:
                                    if (uint16(v18) < 255) {
                                        v3 += v3
                                    }
                                }
                            }
                        }
                        _ = v26
                    }
                } else {
                    fmt.Println(fi21)
                }
                switch "a" {
                case "\n\t":
                    var v55 uint32 = (^(v18))
                    _ = v55
                case "a":
                    var v56 int32 = (v1 % ((-862) | 1))
                    var v57 int8 = -128
                    var v58 float32 = (v3 + -2.25)
                    v59 := GenS5{F0: 32767, F1: 42, F2: (9223372036854775807 > 215)}
                    _ = v56
                    _ = v57
                    _ = v58
                    _ = v59
                default:
                    v2 = append(v2, v19)
                }
            } else {
                if (((v5 + (("foo bar" + "") + "\n\t")) + v19) == (("foo bar" + ("foo bar" + ("foo bar" + "foo bar"))) + v19)) {
                    v60 := GenS1{F0: -2147483648}
                    switch fi20 {
                    case -2147483648:
                        var v61 float32 = ((((v3 * v3) * 3.14159) * (1.5 * v3)) * v3)
                        _ = v61
                    }
                    fmt.Println(v18, v1, v1)
                    v5 = v19
                    _ = v60
                } else {
                    if true {
                        v1 -= v1
                    }
                    var v62 uint64 = 42
                    var v63 int = 2147483647
                    _ = v62
                    _ = v63
                }
                switch 128 {
                case 458:
                    var v64 int16 = -710
                    _ = v64
                }
                if true {
                    v19 = v5
                }
            }
            _ = fi21
            _ = v22
        }
        var v65 int8 = 0
        v66 := GenS3{F0: v65, F1: 2, F2: 65535}
        _ = fi20
        _ = v65
        _ = v66
    }
    fmt.Println(v3, v18, v5)
    var v67 bool = (v1 < (-(v1)))
    v68 := GenS6{F0: (-907 > 2147483647), F1: 65535, F2: 42}
    v69 := v4[-436]
    v2 = append(v2, (v5 + v19))
    var v70 int64 = 1
    for fi71 := 0; fi71 < 18; fi71++ {
        for fi72 := 0; fi72 < 13; fi72++ {
            v67 = false
            _ = fi72
        }
        v73 := GenS7{F0: 272}
        for fi74 := 0; fi74 < 6; fi74++ {
            v75 := v4[(fi71 >> ((471) & 63))]
            fmt.Println(fi74, v67, v18)
            v75 = v5
            _ = fi74
            _ = v75
        }
        _ = fi71
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v67
    _ = v68
    _ = v69
    _ = v70
}
