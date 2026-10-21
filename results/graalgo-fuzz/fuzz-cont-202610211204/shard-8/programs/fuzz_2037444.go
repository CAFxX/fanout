package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int32; F2 int16 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 int8 }
type GenS7 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS8 struct { F0 int8; F1 int32; F2 uint }
type GenS9 struct { F0 uint32; F1 uint16 }
func genFn0(p0 bool) uint64 {
    v6 := []int{887, 32768, 9223372036854775807}
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: 255, F1: 127, F2: 1}
        var v9 uint8 = 255
        v10 := GenS1{F0: 310, F1: 100, F2: 42}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := GenS1{F0: 358, F1: 100, F2: 489}
    _ = p0
    _ = v11
    _ = v6
    return 2
}
func genFn1(p0 int32, p1 uint8, p2 bool) string {
    var v12 uint = 65535
    for fi13 := 0; fi13 < 14; fi13++ {
        p2 = p2
        v14 := []int{}
        var v15 string = "foo bar"
        _ = fi13
        _ = v14
        _ = v15
    }
    p2 = (!false)
    var v16 string = ("\n\t" + "hello")
    p2 = (42 < 752)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v16
    return "hello"
}

func main() {
    var v1 uint16 = 2
    var v2 int64 = 282
    v3 := make(map[int]bool)
    var v4 float32 = 1e10
    v5 := make(map[uint32]string)
    var v17 bool = (v1 <= (v1 >> ((0) & 63)))
    switch "hello" {
    case "a":
        _ = len(v3)
    case "\n\t":
        switch ("a" + "") {
        case "\n\t":
            v1 -= (v1 - v1)
        case "hello":
            var v18 int16 = 1
            _ = v18
        case "a":
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 uint = 73
                _ = fi19
                _ = v20
            }
        }
    case "foo bar":
        v2 += (-(v2))
    default:
        switch 4294967295 {
        case -742:
            switch ("" + "foo bar") {
            case "foo bar":
                switch "\n\t" {
                case "":
                    v17 = v17
                case "\n\t":
                    switch "\n\t" {
                    case "\n\t":
                        v21 := GenS2{F0: 2, F1: 19}
                        _ = v21
                    default:
                        var v22 bool = v17
                        _ = v22
                    }
                default:
                    var v23 string = "hello"
                    _ = v23
                }
            }
        case 1:
            for fi24 := 0; fi24 < 4; fi24++ {
                var v25 int = fi24
                var v26 uint8 = 127
                _ = fi24
                _ = v25
                _ = v26
            }
        default:
            v27 := []int{9223372036854775807}
            _ = v27
        }
    }
    v28 := []int{-723, 0, 42, -2147483648}
    for fi29 := 0; fi29 < 19; fi29++ {
        for fi30 := 0; fi30 < 20; fi30++ {
            v17 = v17
            _ = len(v5)
            _ = len(v3)
            _ = fi30
        }
        v2 = -1
        switch "\n\t" {
        case "\n\t":
            switch (^(fi29)) {
            case 531:
                fmt.Println(v1)
            case 127:
                if (!v17) {
                    _ = len(v28)
                    var v31 int8 = -49
                    switch "" {
                    case "a":
                        v4 -= ((v4 - -2.25) * v4)
                    case "":
                        var v32 int = 32767
                        _ = v32
                    }
                    _ = v31
                } else {
                    if ((-(fi29)) == fi29) {
                        switch "a" {
                        case "\n\t":
                            v2 = -716
                        case "":
                            v1 -= v1
                        }
                        v4 = v4
                        switch (fi29 + (fi29 ^ (fi29 / ((fi29) | 1)))) {
                        case 9223372036854775807:
                            v33 := GenS3{F0: 1, F1: -2147483648, F2: -1}
                            _ = v33
                        }
                    } else {
                        var v34 uint = 1
                        _ = v34
                    }
                    fmt.Println(fi29)
                }
            case 364:
                var v35 int32 = -1
                _ = v35
            }
        default:
            var v36 uint32 = 0
            _ = v36
        }
        _ = fi29
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        switch 128 {
        case -9223372036854775808:
            for fi38 := 0; fi38 < 18; fi38++ {
                fmt.Println(v17, v1, fi38)
                v4 -= (3.14159 * v4)
                _ = fi38
            }
        case 9223372036854775807:
            for fi39 := 0; fi39 < 18; fi39++ {
                if (!false) {
                    var v40 uint = 255
                    var v41 uint64 = 2
                    _ = v40
                    _ = v41
                } else {
                    v42 := GenS2{F0: 108, F1: 2}
                    _ = v42
                }
                v43 := GenS2{F0: 18446744073709551615, F1: 128}
                v4 -= 0.1
                _ = fi39
                _ = v43
            }
        case 127:
            for fi44 := 0; fi44 < 4; fi44++ {
                var v45 int = fi37
                var v46 uint32 = 4294967295
                _ = fi44
                _ = v45
                _ = v46
            }
        }
        _ = fi37
    }
    switch "\n\t" {
    case "a":
        v2 -= (v2 & (v2 / ((v2) | 1)))
    case "foo bar":
        if v17 {
            var v47 int = -2
            v48 := GenS3{F0: 1, F1: 535, F2: 1000}
            _ = v47
            _ = v48
        }
    default:
        v1 = 0
    }
    for fi49 := 0; fi49 < 10; fi49++ {
        var v50 bool = (127 != -1)
        _ = fi49
        _ = v50
    }
    v17 = ("\n\t" != ("a" + ("a" + (("" + "a") + "foo bar"))))
    if (54 != -1) {
        for fi51 := 0; fi51 < 14; fi51++ {
            var v52 uint16 = (^(v1))
            _ = fi51
            _ = v52
        }
        if true {
            switch "hello" {
            case "hello":
                var v53 string = ("hello" + "a")
                _ = v53
            }
        } else {
            for fi54 := 0; fi54 < 6; fi54++ {
                v4 += v4
                _ = fi54
            }
            v1 += 350
            v17 = (false && v17)
        }
    } else {
        v55 := GenS2{F0: 65535, F1: 1}
        v56 := GenS1{F0: (v1 << ((0) & 63)), F1: 100, F2: 2}
        v4 = ((1.5 + (v4 + 0.0)) + 3.14159)
        _ = v55
        _ = v56
    }
    v4 -= float32(-2.25)
    var v57 int = 1
    if (!v17) {
        v4 += (v4 + v4)
        v58 := GenS3{F0: 1, F1: -854, F2: 639}
        if (-1 < 1) {
            if v17 {
                switch (^(v57)) {
                case 9223372036854775807:
                    var v59 uint64 = 1
                    _ = v59
                }
                fmt.Println(v4, v4)
            } else {
                for fi60 := 0; fi60 < 16; fi60++ {
                    switch "hello" {
                    case "hello":
                        for fi61 := 0; fi61 < 18; fi61++ {
                            v62 := GenS1{F0: v1, F1: -128, F2: 0}
                            v63 := GenS4{F0: (1 < -1), F1: 127}
                            switch "hello" {
                            case "a":
                                if v17 {
                                    fmt.Println(v2)
                                    fmt.Println(v17, v4, fi61)
                                    v57 = -909
                                } else {
                                    if v17 {
                                        var v64 uint32 = 65535
                                        fmt.Println(v1)
                                        _ = v64
                                    } else {
                                        v5[2147483647] = "a"
                                        var v65 int32 = 119
                                        _ = v65
                                    }
                                    fmt.Println(fi60)
                                }
                            case "\n\t":
                                v66 := v5[2147483648]
                                _ = v66
                            default:
                                v28 = append(v28, (fi60 / ((v57) | 1)))
                            }
                            v57 = (v57 ^ fi61)
                            v67 := GenS5{F0: (v2 << ((255) & 63))}
                            _ = fi61
                            _ = v62
                            _ = v63
                            _ = v67
                        }
                    default:
                        v57 += -1
                    }
                    switch "foo bar" {
                    case "foo bar":
                        if (v17 && true) {
                            if (4294967295 < 889) {
                                if (-1 >= 1000) {
                                    var v68 uint64 = 2
                                    v69 := GenS3{F0: (v68 % ((v68) | 1)), F1: -2147483648, F2: 0}
                                    var v70 bool = (v17 || v17)
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                }
                                for fi71 := 0; fi71 < 2; fi71++ {
                                    v2 = (v2 >> ((5) & 63))
                                    switch "" {
                                    case "":
                                        _ = len(v5)
                                    default:
                                        var v72 uint64 = 18446744073709551615
                                        _ = v72
                                    }
                                    _ = fi71
                                }
                                var v73 bool = (genFn0((0.1 == v4)) >= genFn0((v4 <= (float32((3.14159 / 0.1)) * 0.0))))
                                _ = v73
                            } else {
                                fmt.Println(v4, v4, v1)
                                v74 := GenS6{F0: true, F1: 1}
                                _ = v74
                            }
                            var v75 float32 = ((v4 + 1.5) / ((0.1 - float32(-2.25)) * 3.14159))
                            v76 := GenS4{F0: v17, F1: 100}
                            _ = v75
                            _ = v76
                        } else {
                            v77 := GenS7{F0: 42, F1: 1, F2: 488}
                            v17 = false
                            switch fi60 {
                            case -766:
                                v78 := GenS8{F0: 42, F1: -630, F2: 181}
                                _ = v78
                            default:
                                fmt.Println(v4, fi60, v57)
                            }
                            _ = v77
                        }
                    case "a":
                        var v79 int64 = -1
                        _ = v79
                    default:
                        if (!(!true)) {
                            v2 = (v2 + v2)
                        } else {
                            switch "foo bar" {
                            case "":
                                fmt.Println(v1)
                            case "foo bar":
                                var v80 int64 = v2
                                _ = v80
                            default:
                                var v81 bool = (!v17)
                                _ = v81
                            }
                            if len(v28) >= 2 {
                                v28 = v28[1:len(v28)]
                            }
                        }
                    }
                    var v82 int = 0
                    _ = fi60
                    _ = v82
                }
                v4 -= (v4 - (float32((0.0 * 1e10)) + (1e10 - (v4 - v4))))
                v83 := GenS3{F0: uint64(v1), F1: -2147483648, F2: 32767}
                _ = v83
            }
        } else {
            var v84 uint = 66
            var v85 int64 = v2
            _ = v84
            _ = v85
        }
        _ = v58
    }
    v57 = 128
    v86 := GenS9{F0: 2147483648, F1: (v1 & (v1 & 736))}
    var v87 uint8 = 255
    for fi88 := 0; fi88 < 11; fi88++ {
        v89 := GenS2{F0: 65535, F1: (v87 + 1)}
        _ = fi88
        _ = v89
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v86
    _ = v87
}
