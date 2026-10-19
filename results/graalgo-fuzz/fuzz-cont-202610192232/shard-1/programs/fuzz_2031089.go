package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS6 struct { F0 uint16; F1 bool; F2 int }
type GenS7 struct { F0 int64 }
func genFn0(p0 uint32) uint8 {
    v6 := GenS1{F0: uint8(p0), F1: -1}
    var v7 string = ("" + "foo bar")
    var v8 uint8 = 255
    v9 := []int{128, -459, 1, 0}
    for fi10 := 0; fi10 < 17; fi10++ {
        switch v7 {
        case "foo bar":
            var v11 uint = 13
            _ = v11
        case "\n\t":
            var v12 float32 = -2.25
            _ = v12
        }
        var v13 int = 32767
        _ = fi10
        _ = v13
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 int64, p1 uint64, p2 int32) int {
    p2 = (p2 | 2147483647)
    v14 := []int{1, 2147483647, 0, 4294967295}
    var v15 float64 = 0.0
    p0 += (p0 >> ((474) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 2147483647
}

func main() {
    var v1 uint8 = 128
    var v2 int16 = -672
    v3 := make(map[int]bool)
    v4 := make(map[string]uint32)
    var v5 uint = 895
    v16 := GenS2{F0: v5, F1: 217, F2: (v1 * (^(v1)))}
    if (0 == 96) {
        if false {
            var v17 uint8 = 255
            _ = v17
        }
        v18 := GenS1{F0: v1, F1: 255}
        _ = v18
    }
    if (false || (("\n\t" + "\n\t") <= "")) {
        var v19 int8 = -1
        _ = v19
    }
    if true {
        if true {
            for fi20 := 0; fi20 < 19; fi20++ {
                v1 = (v1 / (((^(v1))) | 1))
                v21 := []int{(fi20 >> (((v5 * (v5 * (^(v5))))) & 63)), fi20, fi20}
                fmt.Println(fi20, v2)
                _ = fi20
                _ = v21
            }
            v5 = (v5 / ((v5) | 1))
        } else {
            v22 := GenS1{F0: v1, F1: -382}
            v1 = (v1 & v1)
            if (true || false) {
                if true {
                    for fi23 := 0; fi23 < 18; fi23++ {
                        if true {
                            switch genFn1(0, 474, 0) {
                            case 442:
                                for fi24 := 0; fi24 < 14; fi24++ {
                                    v5 = (v5 * 665)
                                    _ = fi24
                                }
                            case -1:
                                switch ("hello" + "foo bar") {
                                case "foo bar":
                                    switch "" {
                                    case "\n\t":
                                        v2 += v2
                                    default:
                                        v25 := []int{genFn1(-9223372036854775808, 856, 970)}
                                        _ = v25
                                    }
                                case "a":
                                    if (245 == 67) {
                                        v26 := []int{genFn1(9223372036854775807, 368, int32(fi23))}
                                        v1 = (v1 + 222)
                                        _ = v26
                                    } else {
                                        v5 = (v5 ^ (v5 + (^(v5))))
                                    }
                                case "\n\t":
                                    var v27 uint32 = 639
                                    _ = v27
                                default:
                                    var v28 int = 4294967295
                                    _ = v28
                                }
                            case 42:
                                v5 -= 42
                            }
                            var v29 uint16 = 2
                            _ = v29
                        }
                        if true {
                            for fi30 := 0; fi30 < 10; fi30++ {
                                v31 := GenS1{F0: 2, F1: 0}
                                var v32 bool = false
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                            v33 := GenS1{F0: (v1 | 19), F1: -32}
                            _ = v33
                        } else {
                            for fi34 := 0; fi34 < 17; fi34++ {
                                if false {
                                    switch "a" {
                                    case "\n\t":
                                        v4[""] = 4294967295
                                    default:
                                        v35 := GenS3{F0: 65535}
                                        _ = v35
                                    }
                                    v5 -= (v5 << ((134) & 63))
                                    var v36 uint64 = 616
                                    v37 := GenS2{F0: (v5 + (^(v5))), F1: 11, F2: v1}
                                    _ = v36
                                    _ = v37
                                } else {
                                    switch 32768 {
                                    case 4294967295:
                                        var v38 uint = v5
                                        _ = v38
                                    case 42:
                                        switch "foo bar" {
                                        case "hello":
                                            var v39 uint16 = 565
                                            _ = v39
                                        default:
                                            if (false && false) {
                                                if false {
                                                    v40 := GenS2{F0: v5, F1: 842, F2: 240}
                                                    _ = v40
                                                } else {
                                                    v2 += (v2 - int16(v5))
                                                    if (-1 > 1) {
                                                        switch genFn1(678, 4294967295, -1) {
                                                        case 376:
                                                            v41 := []int{(fi34 + fi34)}
                                                            _ = v41
                                                        case 692:
                                                            var v42 uint8 = v1
                                                            _ = v42
                                                        }
                                                    }
                                                    v1 -= genFn0(42)
                                                }
                                                v1 -= genFn0(406)
                                            }
                                            v5 -= (v5 >> ((v5) & 63))
                                        }
                                        var v43 int = fi34
                                        _ = v43
                                    case 163:
                                        var v44 uint = 4294967295
                                        _ = v44
                                    }
                                }
                                _ = fi34
                            }
                        }
                        fmt.Println(fi23, v2, v1)
                        _ = fi23
                    }
                } else {
                    fmt.Println(v1)
                    v45 := v3[807]
                    _ = v45
                }
                var v46 uint32 = 890
                switch (("" + (("\n\t" + "") + "a")) + (("foo bar" + "a") + "foo bar")) {
                case "hello":
                    var v47 float32 = (1e10 - 0.0)
                    _ = v47
                case "":
                    var v48 int16 = (v2 | (v2 << ((uint(v1)) & 63)))
                    _ = v48
                case "foo bar":
                    if ("hello" == ("hello" + ("hello" + "hello"))) {
                        v1 += (^(v1))
                        v46 += (v46 % ((v46) | 1))
                    }
                }
                _ = v46
            }
            _ = v22
        }
    }
    fmt.Println(v2)
    for fi49 := 0; fi49 < 7; fi49++ {
        switch (fi49 | (fi49 << (((v5 << ((1) & 63))) & 63))) {
        case 2147483647:
            if true {
                var v50 uint8 = v1
                fmt.Println(v2, v50)
                _ = v50
            }
        case -2147483648:
            v1 = 255
        case 9223372036854775807:
            v1 += genFn0(64)
        }
        var v51 float32 = 3.14159
        _ = fi49
        _ = v51
    }
    if ((v2 ^ 0) >= (v2 + -251)) {
        v52 := GenS4{F0: true}
        switch "foo bar" {
        case "foo bar":
            if ((42 >= 0) || false) {
                var v53 string = ""
                _ = v53
            } else {
                if (!false) {
                    switch "" {
                    case "a":
                        var v54 uint16 = 347
                        _ = v54
                    case "foo bar":
                        if ((42 > -1) || (65535 != 42)) {
                            v1 = v1
                            var v55 uint32 = 405
                            _ = v55
                        }
                    case "":
                        v5 -= v5
                    default:
                        if true {
                            switch "" {
                            case "a":
                                v56 := GenS5{F0: genFn0(453), F1: -128, F2: (v1 + 128)}
                                _ = v56
                            default:
                                switch "\n\t" {
                                case "":
                                    _ = len(v4)
                                case "a":
                                    if true {
                                        var v57 uint64 = 42
                                        _ = v57
                                    }
                                }
                            }
                            var v58 float32 = 3.14159
                            _ = v58
                        } else {
                            v59 := GenS6{F0: 1, F1: (1 != 2147483647), F2: 128}
                            switch "foo bar" {
                            case "\n\t":
                                if (!(35 <= (v5 | (v5 % ((65535) | 1))))) {
                                    v60 := []int{2147483647, 155, 4294967295}
                                    _ = v60
                                } else {
                                    var v61 uint16 = 284
                                    _ = v61
                                }
                            case "foo bar":
                                v62 := []int{127, genFn1(0, 1, 1), 4294967295, 2147483647}
                                _ = v62
                            default:
                                if true {
                                    v5 = v5
                                    v2 -= (v2 | v2)
                                    if true {
                                        v63 := []int{1}
                                        if (false || ((0.0 != 3.14159) || true)) {
                                            var v64 float32 = (0.0 / (1.5 / 1.0))
                                            fmt.Println(v64, v2)
                                            _ = v64
                                        }
                                        v2 = 924
                                        _ = v63
                                    } else {
                                        var v65 int16 = v2
                                        _ = v65
                                    }
                                } else {
                                    if ((2147483647 < -1) && (979 > v5)) {
                                        v66 := v4[("a" + "a")]
                                        if ("\n\t" >= "hello") {
                                            var v67 float32 = float32(3.14159)
                                            var v68 uint64 = 1
                                            _ = v67
                                            _ = v68
                                        } else {
                                            fmt.Println(v1, v5, v2)
                                        }
                                        _ = v66
                                    } else {
                                        v2 = v2
                                    }
                                    v5 -= v5
                                }
                            }
                            _ = v59
                        }
                    }
                } else {
                    _ = len(v3)
                }
                fmt.Println(v1)
            }
        case "\n\t":
            if true {
                v69 := GenS3{F0: 42}
                _ = v69
            } else {
                delete(v4, "\n\t")
                v70 := GenS1{F0: (v1 >> (((v5 >> ((uint(v1)) & 63))) & 63)), F1: 303}
                v5 += (v5 % ((65535) | 1))
                _ = v70
            }
        case "hello":
            fmt.Println(v2, v5)
        default:
            v5 = 1
        }
        for fi71 := 0; fi71 < 7; fi71++ {
            v72 := GenS5{F0: 128, F1: 102, F2: (v1 >> ((411) & 63))}
            if (((42 > -207) || (210 < -939)) && true) {
                fmt.Println(v5, v1, v1)
            }
            _ = fi71
            _ = v72
        }
        _ = v52
    } else {
        v1 -= v1
    }
    switch ("a" + (("foo bar" + "foo bar") + ("" + "a"))) {
    case "\n\t":
        if false {
            var v73 uint64 = 0
            _ = v73
        } else {
            var v74 uint16 = 605
            v75 := []int{genFn1(1, 615, 42), -537, int(v74), 32767}
            _ = v74
            _ = v75
        }
    case "a":
        var v76 uint32 = 65535
        _ = v76
    }
    var v77 float64 = (((0.0 * 3.14159) - -2.25) / 1.5)
    v5 = v5
    fmt.Println(v77, v2, v1)
    for fi78 := 0; fi78 < 17; fi78++ {
        var v79 string = "\n\t"
        _ = fi78
        _ = v79
    }
    fmt.Println(v77, v5, v77)
    var v80 int16 = 0
    v81 := GenS7{F0: -40}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v77:", v77)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
    _ = v80
    _ = v81
}
