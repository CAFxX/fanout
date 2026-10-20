package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 string }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 uint }
func genFn0() int {
    var v7 uint16 = 42
    var v8 uint32 = 2147483648
    v9 := GenS1{F0: (967 > 2), F1: "foo bar", F2: "\n\t"}
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}
func genFn1(p0 bool) bool {
    var v10 float64 = 0.1
    var v11 uint8 = 210
    _ = p0
    _ = v10
    _ = v11
    return false
}

func main() {
    var v1 bool = true
    var v2 string = "a"
    v3 := make(map[uint32]uint32)
    v4 := []uint32{143, 65535, 59, 2147483647}
    v5 := make(map[uint32]string)
    var v6 uint = 255
    var v12 int = -34
    var v13 float32 = (3.14159 / float32(-2.25))
    v14 := GenS1{F0: v1, F1: v2, F2: v2}
    for fi15 := 0; fi15 < 18; fi15++ {
        switch (v2 + ("\n\t" + ("hello" + (v2 + v2)))) {
        case "":
            var v16 float64 = -2.25
            _ = v16
        case "\n\t":
            fmt.Println(v6, v6, v13)
        }
        if v1 {
            for fi17 := 0; fi17 < 18; fi17++ {
                v12 += (v12 ^ genFn0())
                _ = fi17
            }
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := GenS2{F0: -660, F1: -1}
                if v1 {
                    v20 := GenS1{F0: (2147483648 != 2147483648), F1: v2, F2: (v2 + v2)}
                    _ = v20
                } else {
                    if len(v4) > 0 {
                        v4[0] = 612
                    }
                }
                switch (("" + "") + "foo bar") {
                case "foo bar":
                    if v1 {
                        var v21 float32 = v13
                        v22 := GenS3{F0: -1, F1: 0}
                        v1 = v1
                        v13 -= v21
                        var v23 int32 = 0
                        v12 += fi18
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        var v24 uint8 = 0
                        var v25 int8 = 127
                        _ = v24
                        _ = v25
                    }
                    fmt.Println(v13, fi15)
                    if v1 {
                        v4 = append(v4, 913)
                    }
                case "":
                    v26 := GenS1{F0: v1, F1: (("a" + "\n\t") + ""), F2: ("hello" + v2)}
                    _ = v26
                case "\n\t":
                    v27 := GenS3{F0: 0, F1: 1}
                    v28 := GenS4{F0: 16, F1: -595}
                    _ = v27
                    _ = v28
                default:
                    if ((!((18446744073709551615 != 614) && (!false))) && v1) {
                        var v29 int16 = 42
                        _ = len(v5)
                        var v30 uint64 = 18446744073709551615
                        _ = v29
                        _ = v30
                    }
                }
                _ = fi18
                _ = v19
            }
            for fi31 := 0; fi31 < 5; fi31++ {
                v5[2147483648] = v2
                fmt.Println(fi15)
                if v1 {
                    v13 -= v13
                } else {
                    fmt.Println(fi15)
                    v32 := GenS5{F0: 0, F1: 2}
                    v33 := GenS2{F0: 32767, F1: 1}
                    _ = v32
                    _ = v33
                }
                v34 := GenS5{F0: 42, F1: 878}
                var v35 bool = ((104 != 0) && false)
                _ = fi31
                _ = v34
                _ = v35
            }
        } else {
            for fi36 := 0; fi36 < 8; fi36++ {
                var v37 uint16 = 1
                switch 2147483647 {
                case 128:
                    if v1 {
                        v12 = (fi15 >> ((v6) & 63))
                    } else {
                        v2 = ""
                        var v38 float64 = (0.1 - (0.1 + 1.5))
                        _ = v38
                    }
                case -1:
                    if v1 {
                        var v39 uint = v6
                        if v1 {
                            var v40 int32 = -737
                            _ = v40
                        }
                        switch "hello" {
                        case "a":
                            var v41 bool = (1 != 41)
                            _ = v41
                        case "hello":
                            var v42 int = (fi36 >> ((4294967295) & 63))
                            _ = v42
                        }
                        _ = v39
                    }
                    switch (fi36 ^ (fi15 ^ (v12 + (fi15 ^ (fi15 | v12))))) {
                    case 9223372036854775807:
                        _ = len(v5)
                    case -581:
                        v2 = v2
                    case -99:
                        v43 := GenS4{F0: 124, F1: 2147483647}
                        _ = v43
                    }
                default:
                    switch v2 {
                    case "hello":
                        var v44 uint64 = 18446744073709551615
                        _ = v44
                    case "\n\t":
                        if v1 {
                            if true {
                                if ((3.14159 * float64(1e10)) == ((1e10 - (1e10 * ((0.1 - (0.0 + (0.1 + 0.0))) / 1e10))) + (-2.25 + 1.5))) {
                                    v1 = v1
                                    switch fi15 {
                                    case -2147483648:
                                        var v45 int8 = 100
                                        _ = v45
                                    case 0:
                                        switch ("\n\t" + v2) {
                                        case "a":
                                            v1 = (!v1)
                                        case "hello":
                                            var v46 float32 = 0.1
                                            _ = v46
                                        }
                                    case 149:
                                        v3[624] = 219
                                    }
                                } else {
                                    fmt.Println(v6)
                                    if ((0 != -9223372036854775808) || (680 >= 2147483647)) {
                                        var v47 uint64 = 1
                                        _ = v47
                                    } else {
                                        v2 = v2
                                        switch (v12 << (((v6 + (v6 | 640))) & 63)) {
                                        case 32767:
                                            delete(v5, 2)
                                        default:
                                            v13 = v13
                                        }
                                        fmt.Println(v1)
                                    }
                                    var v48 bool = (v1 || true)
                                    _ = v48
                                }
                                if v1 {
                                    switch v12 {
                                    case 32767:
                                        fmt.Println(v37)
                                    case -639:
                                        v49 := GenS2{F0: 1000, F1: 42}
                                        _ = v49
                                    }
                                    fmt.Println(v13, fi36, v13)
                                    v50 := GenS1{F0: v1, F1: (v2 + v2), F2: v2}
                                    _ = v50
                                }
                            }
                            v4 = append(v4, 635)
                        }
                    }
                }
                fmt.Println(v37)
                _ = fi36
                _ = v37
            }
            switch 238 {
            case 42:
                for fi51 := 0; fi51 < 12; fi51++ {
                    var v52 string = v2
                    if (!(!(((v2 + v2) + v2) < v52))) {
                        v53 := GenS6{F0: 1}
                        v1 = v1
                        v13 -= v13
                        _ = v53
                    }
                    _ = len(v4)
                    _ = fi51
                    _ = v52
                }
            case 522:
                for fi54 := 0; fi54 < 14; fi54++ {
                    switch fi54 {
                    case 127:
                        v2 = "hello"
                    case 697:
                        _ = len(v4)
                    case 4294967295:
                        v12 = v12
                    }
                    _ = fi54
                }
            case 2147483647:
                v55 := GenS5{F0: 2, F1: 189}
                _ = v55
            }
        }
        if (!v1) {
            fmt.Println(v2, v13, fi15)
            if (588 == 135) {
                switch v2 {
                case "\n\t":
                    var v56 int32 = 1
                    _ = v56
                default:
                    var v57 float64 = 0.1
                    _ = v57
                }
                v6 -= v6
                v1 = v1
            } else {
                var v58 int32 = 42
                _ = v58
            }
            for fi59 := 0; fi59 < 2; fi59++ {
                var v60 int8 = -1
                if v1 {
                    var v61 int8 = (v60 >> (((v6 / ((v6) | 1))) & 63))
                    if v1 {
                        v62 := GenS7{F0: (fi15 | genFn0()), F1: 0}
                        v1 = v1
                        delete(v3, 961)
                        v63 := GenS3{F0: -1, F1: -1}
                        switch "" {
                        case "\n\t":
                            v64 := GenS5{F0: 2, F1: 1}
                            _ = v64
                        case "hello":
                            fmt.Println(fi15, v61, v2)
                        default:
                            v65 := GenS4{F0: 91, F1: 42}
                            _ = v65
                        }
                        _ = v62
                        _ = v63
                    } else {
                        v61 = (v60 % ((100) | 1))
                        var v66 uint8 = 128
                        var v67 int = v12
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v68 := GenS1{F0: (v1 || (v1 && true)), F1: v2, F2: (v2 + v2)}
                        _ = v66
                        _ = v67
                        _ = v68
                    }
                    switch (v12 ^ (v12 & 0)) {
                    case -927:
                        var v69 float32 = v13
                        _ = v69
                    default:
                        v70 := GenS2{F0: -32768, F1: 0}
                        _ = v70
                    }
                    _ = v61
                }
                _ = len(v5)
                v13 = 1.5
                _ = fi59
                _ = v60
            }
        }
        _ = fi15
    }
    for fi71 := 0; fi71 < 17; fi71++ {
        var v72 uint = 709
        v73 := GenS8{F0: 680}
        v1 = false
        _ = fi71
        _ = v72
        _ = v73
    }
    delete(v5, 2)
    var v74 int64 = 42
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "\n\t" {
    case "\n\t":
        for fi75 := 0; fi75 < 7; fi75++ {
            var v76 int16 = -243
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi75
            _ = v76
        }
    case "hello":
        var v77 int = (v12 * v12)
        _ = v77
    case "foo bar":
        for fi78 := 0; fi78 < 9; fi78++ {
            var v79 int = (fi78 >> ((v6) & 63))
            switch (v2 + v2) {
            case "\n\t":
                v80 := v3[0]
                _ = v80
            case "foo bar":
                var v81 float64 = 3.14159
                _ = v81
            case "":
                v79 = v12
            }
            if len(v4) > 0 {
                v4[0] = 1
            }
            _ = fi78
            _ = v79
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
}
