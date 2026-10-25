package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS2 struct { F0 int32; F1 int; F2 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 uint64 }
func genFn0(p0 uint8, p1 string, p2 uint) int {
    v7 := GenS1{F0: 0, F1: 0, F2: 937}
    var v8 uint8 = (^(p0))
    var v9 uint64 = 0
    p0 = v8
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint, p1 bool) int32 {
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := []int{(fi10 & 2147483647), fi10, (fi10 / ((fi10) | 1)), fi10}
        p1 = (((fi10 / ((fi10) | 1)) <= fi10) && (0.0 <= 1e10))
        for fi12 := 0; fi12 < 12; fi12++ {
            switch ("a" + "a") {
            case "foo bar":
                v13 := GenS1{F0: 1, F1: -9223372036854775808, F2: 32767}
                _ = v13
            default:
                fmt.Println(p0, p1, p0)
            }
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            if ("hello" > ("" + (("hello" + "hello") + "foo bar"))) {
                if (234 > 1) {
                    v14 := GenS2{F0: 640, F1: (fi12 % ((fi10) | 1)), F2: 42}
                    var v15 uint32 = 4294967295
                    _ = v14
                    _ = v15
                } else {
                    v11 = append(v11, fi10)
                }
                p0 = 1
                var v16 uint32 = 291
                _ = v16
            }
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    if (!(0 < 2)) {
        var v17 int32 = 1
        _ = v17
    } else {
        p1 = ("\n\t" == "foo bar")
        var v18 int32 = -1
        _ = v18
    }
    switch ("a" + ("" + "foo bar")) {
    case "":
        v19 := []int{29, -66, -9223372036854775808}
        _ = v19
    case "a":
        p1 = ((2147483647 == 2) && ((100 >= -97) && (42 != 42)))
    case "\n\t":
        p0 -= (p0 % ((p0) | 1))
    default:
        p0 = p0
    }
    v20 := []int{128, 0, 306}
    _ = p0
    _ = p1
    _ = v20
    return 42
}

func main() {
    var v1 bool = (941 >= 2)
    var v2 uint64 = 1
    v3 := []uint32{1}
    v4 := make(map[uint32]uint32)
    v5 := []uint32{977, 282, 1}
    v6 := make(map[int]bool)
    if v1 {
        for fi21 := 0; fi21 < 14; fi21++ {
            if ((("hello" + "foo bar") + (("hello" + ("hello" + (("foo bar" + "foo bar") + "\n\t"))) + ("" + ""))) <= (("" + "") + "hello")) {
                v22 := GenS1{F0: v2, F1: -1, F2: 65535}
                _ = v22
            } else {
                var v23 float64 = 1.5
                v2 = v2
                _ = v23
            }
            v1 = false
            var v24 float64 = ((3.14159 - (((1e10 * (-2.25 + 0.1)) + (1.5 + 1.5)) + ((3.14159 - 0.1) + (float64((float32((1.5 - 1.5)) * 3.14159)) - -2.25)))) + 0.0)
            _ = fi21
            _ = v24
        }
        if (!(!v1)) {
            v2 += 2
        }
        if ((255 == 255) || (true || ((-571 >= 432) || (1.5 >= 3.14159)))) {
            var v25 int16 = 42
            _ = v25
        }
    }
    var v26 float64 = ((0.0 + 0.1) + 0.1)
    if (100 == 100) {
        switch -2147483648 {
        case 32768:
            var v27 bool = (-2.25 < (((0.0 / v26) + v26) + float64((3.14159 * (1.5 / 1.0)))))
            _ = v27
        case -535:
            v1 = (2147483647 != 42)
        case 42:
            if v1 {
                v1 = (953 != 32767)
                v1 = (65535 > 2147483648)
            } else {
                if (!v1) {
                    var v28 float32 = 0.1
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    for fi29 := 0; fi29 < 20; fi29++ {
                        v2 = 73
                        switch (fi29 - fi29) {
                        case 32767:
                            v30 := GenS1{F0: (v2 & (v2 << ((809) & 63))), F1: 1, F2: 0}
                            _ = v30
                        case 1:
                            v28 += (v28 - v28)
                        case -242:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        _ = fi29
                    }
                    _ = v28
                } else {
                    v31 := GenS2{F0: 0, F1: genFn0(65, "hello", 4294967295), F2: 42}
                    var v32 float32 = ((0.0 - (-2.25 + 3.14159)) * 1e10)
                    for fi33 := 0; fi33 < 16; fi33++ {
                        if (0 != 1) {
                            var v34 bool = v1
                            v35 := GenS1{F0: 1, F1: 758, F2: 0}
                            if (154 < 4294967295) {
                                v36 := GenS2{F0: -2147483648, F1: genFn0(128, "\n\t", 4294967295), F2: -128}
                                var v37 uint = 18446744073709551615
                                _ = v36
                                _ = v37
                            } else {
                                if v34 {
                                    v32 -= v32
                                } else {
                                    v38 := GenS3{F0: 0}
                                    var v39 uint = 255
                                    _ = v38
                                    _ = v39
                                }
                            }
                            _ = v34
                            _ = v35
                        } else {
                            var v40 uint16 = 8
                            v2 += v2
                            var v41 float32 = v32
                            _ = v40
                            _ = v41
                        }
                        if (128 >= 1) {
                            v26 -= v26
                            if true {
                                var v42 bool = (162 > 2)
                                if v1 {
                                    if v42 {
                                        var v43 string = (("\n\t" + "foo bar") + ("foo bar" + ("foo bar" + "\n\t")))
                                        v44 := v4[0]
                                        switch fi33 {
                                        case 2147483647:
                                            if v1 {
                                                if v1 {
                                                    var v45 uint8 = 0
                                                    for fi46 := 0; fi46 < 5; fi46++ {
                                                        v47 := GenS1{F0: v2, F1: -1, F2: 255}
                                                        v44 = v44
                                                        v43 = ""
                                                        _ = fi46
                                                        _ = v47
                                                    }
                                                    _ = v45
                                                }
                                                v32 += float32(v26)
                                            }
                                        case -2147483648:
                                            for fi48 := 0; fi48 < 9; fi48++ {
                                                v1 = (uint64(v44) < 42)
                                                _ = fi48
                                            }
                                        }
                                        _ = v43
                                        _ = v44
                                    } else {
                                        for fi49 := 0; fi49 < 16; fi49++ {
                                            v3 = append(v3, 950)
                                            _ = fi49
                                        }
                                        var v50 int32 = 1
                                        switch "\n\t" {
                                        case "a":
                                            v51 := GenS1{F0: v2, F1: 0, F2: 32767}
                                            _ = v51
                                        }
                                        _ = v50
                                    }
                                    switch (fi33 * 4294967295) {
                                    case 431:
                                        if v42 {
                                            var v52 float64 = v26
                                            _ = v52
                                        }
                                    case -980:
                                        var v53 string = ("\n\t" + (("hello" + ("a" + "a")) + "a"))
                                        _ = v53
                                    default:
                                        for fi54 := 0; fi54 < 17; fi54++ {
                                            v55 := GenS2{F0: -313, F1: fi33, F2: 119}
                                            _ = fi54
                                            _ = v55
                                        }
                                    }
                                }
                                _ = v42
                            }
                            if (((63 <= 1) && ((("a" + "") + ("\n\t" + "foo bar")) == "")) && ((v2 - (v2 - (v2 & v2))) > (v2 * v2))) {
                                v56 := GenS1{F0: (v2 & 0), F1: -9223372036854775808, F2: 32767}
                                v57 := GenS4{F0: (^(v2)), F1: 100}
                                v2 -= (v2 << ((4294967295) & 63))
                                _ = v56
                                _ = v57
                            }
                        } else {
                            if v1 {
                                v1 = v1
                                v58 := GenS5{F0: 255}
                                _ = v58
                            } else {
                                v59 := GenS4{F0: v2, F1: 98}
                                fmt.Println(v2)
                                _ = v59
                            }
                            v26 += 1.5
                            var v60 int16 = -301
                            _ = v60
                        }
                        if v1 {
                            if v1 {
                                v61 := GenS3{F0: -515}
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v1 = false
                                _ = v61
                            } else {
                                v2 += (^(v2))
                                var v62 int64 = -9223372036854775808
                                _ = v62
                            }
                        } else {
                            fmt.Println(fi33, v1, v1)
                            if (42 == 1000) {
                                var v63 uint8 = 1
                                v64 := GenS6{F0: 42, F1: 0}
                                v32 -= (((v32 - v32) + (((-2.25 * -2.25) + 1e10) / ((v32 + v32) + (3.14159 * v32)))) + -2.25)
                                _ = v63
                                _ = v64
                            }
                        }
                        _ = fi33
                    }
                    _ = v31
                    _ = v32
                }
                var v65 uint = 1
                _ = v65
            }
        default:
            if true {
                var v66 float32 = 3.14159
                var v67 int32 = 8
                _ = v66
                _ = v67
            } else {
                v26 = (((float64(0.0) + v26) - (v26 + (v26 - 0.0))) * (0.1 / v26))
                fmt.Println(v26, v26, v1)
                if (1 >= 428) {
                    if (127 != 42) {
                        v5 = append(v5, 839)
                    } else {
                        v1 = (-1 == 647)
                        _ = len(v5)
                    }
                    for fi68 := 0; fi68 < 9; fi68++ {
                        delete(v4, 2147483648)
                        if (!(876 < 1)) {
                            v69 := v4[170]
                            _ = v69
                        }
                        _ = fi68
                    }
                } else {
                    for fi70 := 0; fi70 < 2; fi70++ {
                        v71 := v4[826]
                        v2 = (v2 * 4294967295)
                        var v72 int64 = 9223372036854775807
                        _ = fi70
                        _ = v71
                        _ = v72
                    }
                    for fi73 := 0; fi73 < 10; fi73++ {
                        v2 -= 4294967295
                        _ = fi73
                    }
                    switch "" {
                    case "foo bar":
                        for fi74 := 0; fi74 < 19; fi74++ {
                            for fi75 := 0; fi75 < 13; fi75++ {
                                v76 := GenS4{F0: 42, F1: -76}
                                v2 = v2
                                _ = fi75
                                _ = v76
                            }
                            v77 := GenS7{F0: -434, F1: genFn1(0, v1)}
                            v78 := GenS7{F0: 0, F1: 42}
                            _ = fi74
                            _ = v77
                            _ = v78
                        }
                    }
                }
            }
        }
        switch (("\n\t" + "hello") + "\n\t") {
        case "hello":
            for fi79 := 0; fi79 < 20; fi79++ {
                for fi80 := 0; fi80 < 2; fi80++ {
                    _ = len(v3)
                    v81 := GenS8{F0: (v2 ^ v2)}
                    var v82 uint8 = 207
                    v26 = v26
                    switch "" {
                    case "a":
                        v83 := GenS7{F0: -995, F1: -2147483648}
                        _ = v83
                    }
                    fmt.Println(v2, v26, v82)
                    _ = fi80
                    _ = v81
                    _ = v82
                }
                _ = fi79
            }
        case "\n\t":
            if (v2 != (v2 % ((v2) | 1))) {
                for fi84 := 0; fi84 < 7; fi84++ {
                    switch 127 {
                    case -276:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    default:
                        var v85 float64 = (0.1 * -2.25)
                        _ = v85
                    }
                    for fi86 := 0; fi86 < 6; fi86++ {
                        v3 = append(v3, 65535)
                        _ = fi86
                    }
                    _ = fi84
                }
            } else {
                var v87 float64 = v26
                _ = v87
            }
        case "":
            var v88 float64 = 1e10
            _ = v88
        default:
            for fi89 := 0; fi89 < 20; fi89++ {
                v1 = v1
                _ = fi89
            }
        }
    } else {
        var v90 uint32 = 65535
        for fi91 := 0; fi91 < 14; fi91++ {
            v2 = (v2 - 4294967295)
            v92 := v6[genFn0(128, "a", 355)]
            _ = len(v6)
            _ = fi91
            _ = v92
        }
        _ = v90
    }
    v93 := GenS3{F0: -479}
    var v94 int = 42
    v95 := GenS6{F0: 42, F1: 986}
    var v96 int = -701
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v94:", v94)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v93
    _ = v94
    _ = v95
    _ = v96
}
