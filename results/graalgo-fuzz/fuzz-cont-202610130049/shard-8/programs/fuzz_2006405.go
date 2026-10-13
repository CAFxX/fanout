package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 string; F1 uint; F2 int8 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16; F1 string; F2 uint32 }
type GenS8 struct { F0 int16; F1 uint16 }
type GenS9 struct { F0 int16; F1 uint32; F2 string }
type GenS10 struct { F0 uint16; F1 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 0
    v3 := []uint32{}
    var v4 uint32 = 3
    v4 = (v4 % (((v4 ^ v4)) | 1))
    var v5 int32 = 0
    if (36 != 2) {
        var v6 string = ("foo bar" + (("foo bar" + (("foo bar" + ("hello" + ("\n\t" + "\n\t"))) + ("" + ("hello" + "hello")))) + "a"))
        v5 -= (v5 * (v5 | v5))
        _ = v6
    } else {
        if (14 != 127) {
            if (float32(-2.25) <= float32((0.1 / (((0.1 + (0.0 * -2.25)) * 0.1) - ((3.14159 / 1.5) * 0.1))))) {
                v1[(v4 ^ (v4 % (((^(v4))) | 1)))] = false
            } else {
                switch 128 {
                case -95:
                    v4 += (^(v4))
                case 32768:
                    if true {
                        fmt.Println(v5, v5, v5)
                        if false {
                            var v7 int16 = 0
                            _ = v7
                        } else {
                            v8 := GenS1{F0: "\n\t", F1: 1}
                            _ = v8
                        }
                    } else {
                        for fi9 := 0; fi9 < 20; fi9++ {
                            v2 -= v2
                            if (!true) {
                                var v10 float32 = 1.5
                                if false {
                                    v11 := GenS1{F0: "a", F1: 1}
                                    for fi12 := 0; fi12 < 17; fi12++ {
                                        if false {
                                            v13 := GenS1{F0: "a", F1: 4294967295}
                                            var v14 int8 = -128
                                            var v15 int32 = v5
                                            _ = v13
                                            _ = v14
                                            _ = v15
                                        }
                                        _ = fi12
                                    }
                                    _ = v11
                                } else {
                                    for fi16 := 0; fi16 < 9; fi16++ {
                                        v4 = (^(v4))
                                        var v17 int16 = 0
                                        var v18 int = fi9
                                        _ = fi16
                                        _ = v17
                                        _ = v18
                                    }
                                }
                                _ = v10
                            } else {
                                var v19 uint8 = 1
                                _ = v19
                            }
                            var v20 int = fi9
                            _ = fi9
                            _ = v20
                        }
                        _ = len(v3)
                    }
                default:
                    if true {
                        fmt.Println(v4, v4, v2)
                        v2 -= v2
                    } else {
                        v4 = (v4 | v4)
                        v5 += (^(v5))
                        switch (("a" + ("foo bar" + "\n\t")) + "hello") {
                        case "\n\t":
                            fmt.Println(v4, v4, v2)
                        case "foo bar":
                            if (!(!false)) {
                                v2 = (v2 & 65535)
                                fmt.Println(v5)
                                var v21 uint64 = 1
                                _ = v21
                            }
                        }
                    }
                }
                var v22 string = "a"
                _ = v22
            }
        }
        if (!true) {
            if true {
                v23 := GenS1{F0: ("a" + "\n\t"), F1: 4294967295}
                _ = v23
            } else {
                for fi24 := 0; fi24 < 12; fi24++ {
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                    fmt.Println(fi24, fi24)
                    _ = fi24
                }
            }
            v4 = v4
        } else {
            var v25 uint16 = 876
            v26 := GenS2{F0: "", F1: 2, F2: 100}
            if false {
                var v27 uint64 = 18446744073709551615
                _ = v27
            }
            _ = v25
            _ = v26
        }
        v2 = uint16(v4)
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        if (true && false) {
            fmt.Println(v5)
            if len(v3) > 0 {
                v3[0] = (v4 % ((v4) | 1))
            }
            switch ("" + ("" + "a")) {
            case "a":
                switch (^(fi28)) {
                case 0:
                    for fi29 := 0; fi29 < 18; fi29++ {
                        var v30 int8 = 42
                        v30 += 100
                        if true {
                            v5 += (v5 % ((v5) | 1))
                        }
                        _ = fi29
                        _ = v30
                    }
                case -1:
                    v31 := GenS2{F0: "hello", F1: 435, F2: 1}
                    _ = v31
                case 819:
                    var v32 int = (fi28 >> ((2) & 63))
                    _ = v32
                }
            case "hello":
                var v33 uint = 42
                _ = v33
            default:
                fmt.Println(v2, v4, v2)
            }
        } else {
            var v34 string = ("foo bar" + "hello")
            for fi35 := 0; fi35 < 4; fi35++ {
                var v36 uint8 = 2
                var v37 string = (v34 + v34)
                v5 -= (v5 >> ((4294967295) & 63))
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = v34
        }
        _ = fi28
    }
    var v38 float64 = (1.5 / 0.1)
    v38 = ((float64(0.1) / 1.0) + (v38 + (v38 / (((v38 + v38) * 1.5) / v38))))
    v5 = (-(v5))
    v4 = (v4 * v4)
    switch ("hello" + "") {
    case "\n\t":
        var v39 int8 = -1
        _ = v39
    case "hello":
        fmt.Println(v2, v5, v38)
    default:
        v40 := GenS3{F0: 127, F1: (v5 % ((v5) | 1))}
        _ = v40
    }
    switch "" {
    case "foo bar":
        v41 := GenS1{F0: "foo bar", F1: 4294967295}
        _ = v41
    case "hello":
        v4 = 4294967295
    case "a":
        for fi42 := 0; fi42 < 11; fi42++ {
            var v43 int = (fi42 * 412)
            v44 := GenS4{F0: "\n\t", F1: v5}
            v45 := GenS3{F0: -1, F1: v5}
            _ = fi42
            _ = v43
            _ = v44
            _ = v45
        }
    default:
        v46 := GenS5{F0: 804}
        _ = v46
    }
    v47 := GenS6{F0: v4}
    if len(v3) > 0 {
        v3[0] = v4
    }
    var v48 uint8 = 42
    for fi49 := 0; fi49 < 8; fi49++ {
        v2 = (v2 * (v2 & 42))
        _ = fi49
    }
    v50 := GenS4{F0: ("hello" + "foo bar"), F1: (^(v5))}
    if false {
        for fi51 := 0; fi51 < 14; fi51++ {
            if (743 > 0) {
                if (((!((^(v48)) == v48)) && false) || false) {
                    var v52 int32 = v5
                    v5 = (v5 - (v5 & (v52 - v5)))
                    _ = v52
                }
                switch (fi51 % ((-2147483648) | 1)) {
                case 127:
                    var v53 float32 = 1e10
                    _ = v53
                case 42:
                    switch fi51 {
                    case 733:
                        for fi54 := 0; fi54 < 4; fi54++ {
                            v48 += (^(v48))
                            if (("foo bar" + ("\n\t" + (("" + ("hello" + ("" + ("foo bar" + (("a" + (("a" + "hello") + "a")) + "a"))))) + ("" + "hello")))) != "\n\t") {
                                switch "" {
                                case "foo bar":
                                    v55 := v1[(v4 / (((v4 * 2)) | 1))]
                                    _ = v55
                                case "":
                                    v5 -= (v5 - 1)
                                    if ((v4 | v4) < v4) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        switch 42 {
                                        case 1:
                                            v2 = (^(v2))
                                        case -2147483648:
                                            v56 := GenS7{F0: v2, F1: "", F2: v4}
                                            _ = v56
                                        case 41:
                                            var v57 int = fi54
                                            _ = v57
                                        }
                                    }
                                    var v58 uint64 = 0
                                    _ = v58
                                case "hello":
                                    v38 += v38
                                default:
                                    var v59 int8 = -1
                                    _ = v59
                                }
                            } else {
                                if true {
                                    if (v5 > 354) {
                                        if (false || ((v48 & (v48 / (((v48 << ((0) & 63))) | 1))) <= (v48 & 2))) {
                                            v60 := GenS8{F0: 1, F1: v2}
                                            var v61 int8 = 127
                                            v62 := GenS7{F0: (^(v2)), F1: "hello", F2: 2147483648}
                                            var v63 bool = (v4 >= 2)
                                            if (v38 <= v38) {
                                                var v64 uint8 = (v48 * v48)
                                                if len(v3) > 0 {
                                                    v3[0] = v4
                                                }
                                                v5 = 42
                                                _ = v64
                                            }
                                            _ = v60
                                            _ = v61
                                            _ = v62
                                            _ = v63
                                        }
                                        fmt.Println(fi51, v48, v5)
                                        var v65 uint32 = (v4 ^ v4)
                                        _ = v65
                                    } else {
                                        v38 -= -2.25
                                        var v66 uint = 2
                                        _ = v66
                                    }
                                    v67 := GenS8{F0: 0, F1: 65535}
                                    var v68 int16 = -32768
                                    var v69 uint16 = uint16(v4)
                                    if (int8(v68) > 1) {
                                        if (!false) {
                                            v70 := GenS6{F0: (v4 | (v4 * v4))}
                                            _ = v70
                                        }
                                        v71 := GenS9{F0: v68, F1: (v4 & v4), F2: ("hello" + "")}
                                        v2 += (^(v69))
                                        fmt.Println(v38, v38, v68)
                                        var v72 float64 = 1.5
                                        _ = v71
                                        _ = v72
                                    }
                                    _ = v67
                                    _ = v68
                                    _ = v69
                                }
                                if len(v3) > 0 {
                                    v3[0] = v4
                                }
                            }
                            v2 = v2
                            _ = fi54
                        }
                    case 127:
                        var v73 uint = 2
                        _ = v73
                    }
                case 4294967295:
                    if (!(9223372036854775807 > int64(v2))) {
                        v48 -= uint8(v2)
                        var v74 uint = 18446744073709551615
                        var v75 int8 = 100
                        _ = v74
                        _ = v75
                    } else {
                        if false {
                            delete(v1, 65535)
                            _ = len(v1)
                            if len(v3) > 0 {
                                v3[0] = v4
                            }
                        }
                        if false {
                            _ = len(v3)
                            for fi76 := 0; fi76 < 2; fi76++ {
                                _ = len(v1)
                                fmt.Println(v38)
                                _ = fi76
                            }
                            v48 += (v48 << ((65535) & 63))
                        } else {
                            v1[(^(v4))] = ("a" >= "foo bar")
                            v2 = v2
                        }
                    }
                }
                v5 -= v5
            } else {
                if ((fi51 / (((-(fi51))) | 1)) < (fi51 & (fi51 | 127))) {
                    v77 := GenS10{F0: v2, F1: 0}
                    delete(v1, (v4 + v4))
                    _ = v77
                }
            }
            v4 = (^(v4))
            _ = fi51
        }
        v3 = append(v3, (v4 * v4))
        switch "" {
        case "\n\t":
            for fi78 := 0; fi78 < 11; fi78++ {
                fmt.Println(v2, v4, v4)
                _ = fi78
            }
        case "a":
            _ = len(v3)
        case "":
            v79 := GenS4{F0: "", F1: (^(v5))}
            _ = v79
        default:
            v4 = 4294967295
        }
    }
    var v80 uint8 = 0
    v81 := GenS1{F0: "a", F1: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v50
    _ = v80
    _ = v81
}
