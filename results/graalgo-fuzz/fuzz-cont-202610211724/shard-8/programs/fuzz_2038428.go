package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS6 struct { F0 int; F1 uint16; F2 bool }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint16; F1 int32 }
func genFn0(p0 int8, p1 uint16) int16 {
    switch 0 {
    case 0:
        switch -795 {
        case 32767:
            p1 += p1
        case 9223372036854775807:
            var v7 int = -9223372036854775808
            _ = v7
        default:
            v8 := []int{0, int(p1)}
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 int = (fi9 & fi9)
            _ = fi9
            _ = v10
        }
    }
    v11 := []int{32767, -9223372036854775808, 32767}
    var v12 uint16 = 228
    var v13 uint64 = 0
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return -320
}

func main() {
    var v1 string = "a"
    var v2 uint = 65535
    var v3 float32 = 1e10
    v4 := []string{(v1 + v1), (v1 + ((v1 + "") + ("hello" + ("" + (v1 + (v1 + v1))))))}
    var v5 uint16 = 255
    var v6 int16 = 1
    if ((true || (0 != 2)) || (!true)) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v14 int = -9223372036854775808
        _ = v14
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v5 -= v5
    }
    v4 = append(v4, "\n\t")
    v3 += v3
    switch 32767 {
    case 128:
        if false {
            var v15 uint16 = v5
            _ = v15
        }
    case 1:
        _ = len(v4)
    }
    if false {
        v3 = v3
        if (2147483647 <= -398) {
            var v16 int16 = (v6 - (v6 & v6))
            var v17 uint64 = 42
            _ = v16
            _ = v17
        } else {
            var v18 uint64 = uint64(v5)
            v1 = v1
            switch ("foo bar" + ("a" + "\n\t")) {
            case "foo bar":
                for fi19 := 0; fi19 < 17; fi19++ {
                    v20 := GenS1{F0: (v2 | (v2 << (((v2 >> ((v2) & 63))) & 63))), F1: 127, F2: 1}
                    _ = fi19
                    _ = v20
                }
            case "a":
                for fi21 := 0; fi21 < 3; fi21++ {
                    v22 := GenS2{F0: 9223372036854775807}
                    var v23 int16 = v6
                    for fi24 := 0; fi24 < 7; fi24++ {
                        v5 = uint16(v2)
                        var v25 int32 = 517
                        v25 = -1
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi24
                        _ = v25
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            case "hello":
                if (!(float64(((v3 - float32((0.0 - float64(v3)))) + (3.14159 * 0.0))) > 0.1)) {
                    v18 -= v18
                    var v26 int8 = 127
                    _ = v26
                }
            }
            _ = v18
        }
        var v27 uint64 = 2
        _ = v27
    } else {
        for fi28 := 0; fi28 < 19; fi28++ {
            var v29 uint64 = 2
            v30 := GenS1{F0: (v2 % (((v2 / ((v2) | 1))) | 1)), F1: 1, F2: (v6 * (v6 / ((1) | 1)))}
            var v31 string = (((v1 + ((("" + ("a" + (v1 + (v1 + v1)))) + v1) + "a")) + (v1 + "a")) + "")
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
        }
        if (1.5 == 1e10) {
            var v32 uint = (v2 + v2)
            var v33 uint16 = 2
            var v34 int8 = -37
            _ = v32
            _ = v33
            _ = v34
        }
    }
    v4 = append(v4, (v1 + (v1 + ("" + ((v1 + v1) + (v1 + v1))))))
    v3 = v3
    for fi35 := 0; fi35 < 9; fi35++ {
        if (false && false) {
            var v36 float32 = v3
            _ = v36
        } else {
            v1 = (v1 + v1)
            var v37 float32 = (float32(((1.5 - -2.25) / (float64(0.1) * -2.25))) * (((v3 / (v3 + v3)) + -2.25) / float32(0.1)))
            _ = v37
        }
        if (732 <= 0) {
            switch (^(fi35)) {
            case 1:
                v6 = (^(v6))
            case -641:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                var v38 uint = (^(v2))
                _ = v38
            }
            v1 = "\n\t"
        } else {
            v39 := GenS3{F0: 1, F1: true}
            var v40 uint = 375
            _ = len(v4)
            _ = v39
            _ = v40
        }
        _ = fi35
    }
    if (true || (false && (80 != 42))) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v1 = v1
        if (uint8(v5) > uint8(v2)) {
            fmt.Println(v3, v1, v5)
        } else {
            switch 4294967295 {
            case 32767:
                for fi41 := 0; fi41 < 16; fi41++ {
                    if ((v1 < "\n\t") && true) {
                        var v42 uint8 = uint8(v2)
                        if ((true && false) && (true && false)) {
                            for fi43 := 0; fi43 < 20; fi43++ {
                                var v44 int8 = -128
                                var v45 int64 = 1
                                v1 = v1
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                            v42 = uint8(v2)
                            v46 := GenS4{F0: 332, F1: v1}
                            _ = v46
                        } else {
                            switch fi41 {
                            case 128:
                                v2 = 839
                            default:
                                switch "" {
                                case "hello":
                                    if ((v6 >> (((v2 % ((v2) | 1))) & 63)) <= v6) {
                                        var v47 uint64 = 258
                                        if true {
                                            v6 -= v6
                                            var v48 bool = true
                                            for fi49 := 0; fi49 < 17; fi49++ {
                                                v3 = float32(float64(v3))
                                                var v50 uint8 = 104
                                                _ = fi49
                                                _ = v50
                                            }
                                            _ = v48
                                        }
                                        _ = v47
                                    } else {
                                        var v51 float64 = 1.5
                                        _ = v51
                                    }
                                case "\n\t":
                                    fmt.Println(v2, v6)
                                case "foo bar":
                                    for fi52 := 0; fi52 < 14; fi52++ {
                                        v6 = v6
                                        v53 := GenS1{F0: v2, F1: -128, F2: (v6 >> (((v2 ^ 0)) & 63))}
                                        var v54 uint = v2
                                        _ = fi52
                                        _ = v53
                                        _ = v54
                                    }
                                default:
                                    for fi55 := 0; fi55 < 19; fi55++ {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v56 bool = ((!(!false)) && true)
                                        _ = fi55
                                        _ = v56
                                    }
                                }
                            }
                        }
                        v57 := GenS5{F0: v5, F1: 504, F2: -1}
                        _ = v42
                        _ = v57
                    }
                    _ = fi41
                }
            case 128:
                switch 32768 {
                case 2147483647:
                    switch 1 {
                    case -1:
                        switch v1 {
                        case "":
                            var v58 int64 = 9223372036854775807
                            _ = v58
                        case "foo bar":
                            var v59 string = v1
                            _ = v59
                        }
                    case 42:
                        var v60 int = 38
                        _ = v60
                    case -169:
                        switch -2147483648 {
                        case 113:
                            var v61 uint32 = 873
                            _ = v61
                        case 2147483647:
                            for fi62 := 0; fi62 < 19; fi62++ {
                                if true {
                                    _ = len(v4)
                                    var v63 int8 = 37
                                    v64 := GenS2{F0: 42}
                                    _ = v63
                                    _ = v64
                                } else {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v65 := GenS6{F0: fi62, F1: v5, F2: ((v2 | (v2 + v2)) != v2)}
                                        _ = v65
                                    case "hello":
                                        for fi66 := 0; fi66 < 12; fi66++ {
                                            switch v1 {
                                            case "":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case "a":
                                                var v67 uint32 = 65535
                                                _ = v67
                                            case "foo bar":
                                                if (true || false) {
                                                    if (4294967295 >= 931) {
                                                        fmt.Println(v6)
                                                        v2 = v2
                                                        var v68 bool = ((68 != 125) || (!false))
                                                        _ = v68
                                                    } else {
                                                        if (-1 != 9223372036854775807) {
                                                            v2 = (v2 | v2)
                                                        } else {
                                                            var v69 uint64 = 2
                                                            v70 := GenS7{F0: uint32(v5)}
                                                            _ = v69
                                                            _ = v70
                                                        }
                                                        switch fi62 {
                                                        case 32768:
                                                            v6 += v6
                                                        case 159:
                                                            switch v1 {
                                                            case "hello":
                                                                v6 = v6
                                                            }
                                                        default:
                                                            v3 = ((v3 + v3) - ((v3 - 1.5) + float32(0.0)))
                                                        }
                                                        v5 = 859
                                                    }
                                                    var v71 string = v1
                                                    if (1.5 < 1e10) {
                                                        var v72 uint16 = (v5 * v5)
                                                        v73 := GenS1{F0: (v2 + 42), F1: 100, F2: v6}
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                    _ = v71
                                                }
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = fi66
                                        }
                                    }
                                    switch (v1 + ((v1 + "foo bar") + v1)) {
                                    case "foo bar":
                                        for fi74 := 0; fi74 < 6; fi74++ {
                                            v75 := GenS2{F0: 0}
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = fi74
                                            _ = v75
                                        }
                                    case "":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case "a":
                                        var v76 int = int(v2)
                                        _ = v76
                                    default:
                                        v2 = v2
                                    }
                                }
                                var v77 int = fi62
                                _ = fi62
                                _ = v77
                            }
                        }
                    }
                default:
                    _ = len(v4)
                }
            case 1:
                var v78 uint8 = 42
                _ = v78
            }
            fmt.Println(v5)
        }
    } else {
        v79 := GenS6{F0: 292, F1: uint16(v2), F2: (int(v5) >= 9223372036854775807)}
        switch -162 {
        case -309:
            if false {
                v3 = 3.14159
                v80 := GenS8{F0: (v5 << ((v2) & 63)), F1: -1}
                if len(v4) > 0 {
                    v4[0] = v1
                }
                _ = v80
            } else {
                var v81 int16 = v6
                fmt.Println(v6, v6, v1)
                _ = v81
            }
        }
        switch (v1 + v1) {
        case "hello":
            for fi82 := 0; fi82 < 4; fi82++ {
                var v83 int64 = 271
                var v84 uint = (v2 * v2)
                _ = len(v4)
                _ = fi82
                _ = v83
                _ = v84
            }
        case "":
            var v85 uint32 = uint32(v5)
            _ = v85
        default:
            if true {
                fmt.Println(v5, v6, v5)
                v86 := GenS8{F0: 65535, F1: 2147483647}
                _ = v86
            } else {
                v3 -= 3.14159
                if true {
                    var v87 int32 = -887
                    _ = v87
                } else {
                    var v88 uint32 = 65535
                    v89 := GenS8{F0: v5, F1: 0}
                    _ = v88
                    _ = v89
                }
                v3 = ((-2.25 - v3) / (v3 - ((v3 + -2.25) + 0.1)))
            }
        }
        _ = v79
    }
    switch -5 {
    case 4294967295:
        var v90 bool = true
        _ = v90
    default:
        switch -381 {
        case -565:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 435:
            if len(v4) > 0 {
                v4[0] = (("hello" + ((((v1 + ("foo bar" + v1)) + "foo bar") + v1) + (("a" + v1) + ("" + (((v1 + "\n\t") + v1) + "foo bar"))))) + v1)
            }
        case 0:
            var v91 int8 = 127
            _ = v91
        }
    }
    var v92 float64 = (1e10 * (0.0 - 3.14159))
    v4 = append(v4, (v1 + ((v1 + ("\n\t" + ((((((("hello" + "foo bar") + "hello") + "") + "\n\t") + (v1 + ("foo bar" + v1))) + ((v1 + "hello") + "\n\t")) + v1))) + "hello")))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v92
}
