package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 string; F2 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 bool; F2 int8 }
type GenS5 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int16; F1 int }
type GenS9 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS10 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS11 struct { F0 uint }
func genFn0(p0 uint16, p1 int64) int8 {
    var v6 int32 = -2147483648
    var v7 float64 = (0.0 * -2.25)
    var v8 uint = 255
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int32, p1 uint32) int8 {
    var v9 uint = 2
    for fi10 := 0; fi10 < 11; fi10++ {
        if true {
            v11 := GenS1{F0: p0}
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := []int{(fi12 & (fi10 + 1))}
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        _ = fi10
    }
    var v14 int64 = 42
    var v15 bool = (687 >= 42)
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v9
    return 0
}

func main() {
    var v1 float32 = 1.5
    var v2 uint8 = 1
    v3 := make(map[uint32]uint32)
    var v4 float32 = v1
    var v5 int32 = 42
    v4 = v1
    v5 -= v5
    fmt.Println(v4)
    if true {
        v16 := v3[567]
        v4 = ((((float32(1.5) / v1) / v1) + float32(((3.14159 / 1.0) - 3.14159))) + 0.0)
        _ = v16
    } else {
        if (!false) {
            v17 := []int{}
            delete(v3, 4294967295)
            _ = v17
        } else {
            var v18 int64 = 9223372036854775807
            fmt.Println(v5)
            var v19 float32 = v4
            _ = v18
            _ = v19
        }
        if (false || (0 > 2)) {
            var v20 bool = ("a" > ((("\n\t" + "a") + ("\n\t" + ("a" + "a"))) + ("" + ((("a" + "hello") + "hello") + "hello"))))
            v21 := []int{}
            if (v20 || true) {
                var v22 uint8 = (v2 << ((1) & 63))
                v20 = (0 != 619)
                _ = v22
            } else {
                for fi23 := 0; fi23 < 12; fi23++ {
                    v24 := GenS2{F0: 0, F1: ("a" + ("" + "hello")), F2: 32767}
                    v25 := GenS3{F0: -1}
                    v26 := GenS2{F0: 0, F1: ("foo bar" + "a"), F2: 42}
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
            _ = v20
            _ = v21
        }
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        var v28 string = ("hello" + ("a" + ((("foo bar" + ((("" + "") + "\n\t") + "hello")) + "a") + "a")))
        _ = fi27
        _ = v28
    }
    if (!("" < "hello")) {
        v4 = 0.0
        if false {
            fmt.Println(v1, v2, v1)
        } else {
            for fi29 := 0; fi29 < 20; fi29++ {
                var v30 int32 = v5
                v2 = 127
                if false {
                    var v31 uint = 1
                    switch ("" + "hello") {
                    case "foo bar":
                        v32 := GenS4{F0: 18446744073709551615, F1: false, F2: -109}
                        _ = v32
                    default:
                        var v33 int16 = -1
                        _ = v33
                    }
                    _ = v31
                } else {
                    var v34 uint8 = (v2 + (v2 - v2))
                    v34 += v2
                    _ = v34
                }
                _ = fi29
                _ = v30
            }
        }
    } else {
        var v35 uint8 = 128
        if (!false) {
            v5 = (-(v5))
            v2 += (v2 | 42)
            fmt.Println(v2, v5, v1)
        } else {
            for fi36 := 0; fi36 < 15; fi36++ {
                v37 := []int{}
                var v38 int64 = 1
                _ = fi36
                _ = v37
                _ = v38
            }
        }
        _ = v35
    }
    v4 -= 0.0
    if false {
        _ = len(v3)
        delete(v3, 256)
    } else {
        v39 := GenS5{F0: (v2 | (v2 - v2)), F1: 255, F2: uint64(v2)}
        _ = v39
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        switch "a" {
        case "foo bar":
            var v41 uint8 = (^(v2))
            _ = v41
        case "":
            var v42 float64 = (0.1 * -2.25)
            _ = v42
        case "hello":
            var v43 uint64 = 18446744073709551615
            _ = v43
        }
        v44 := GenS4{F0: 18446744073709551615, F1: true, F2: 0}
        var v45 float64 = float64(v4)
        _ = fi40
        _ = v44
        _ = v45
    }
    if (598 != 2) {
        if true {
            v2 = v2
            switch "hello" {
            case "a":
                for fi46 := 0; fi46 < 6; fi46++ {
                    var v47 int16 = 0
                    v48 := GenS6{F0: (v2 * 128), F1: 0, F2: 1}
                    var v49 int32 = (-(v5))
                    _ = fi46
                    _ = v47
                    _ = v48
                    _ = v49
                }
            case "":
                v50 := []int{}
                _ = v50
            }
        }
        _ = len(v3)
        if false {
            switch 9223372036854775807 {
            case 128:
                switch (("" + "a") + "hello") {
                case "hello":
                    v51 := []int{-304, -9223372036854775808}
                    _ = v51
                case "a":
                    switch "\n\t" {
                    case "hello":
                        v52 := GenS7{F0: -9223372036854775808}
                        _ = v52
                    case "a":
                        fmt.Println(v4, v1, v5)
                    case "":
                        var v53 bool = (42 < 65535)
                        _ = v53
                    }
                default:
                    v54 := GenS8{F0: 1000, F1: 9223372036854775807}
                    _ = v54
                }
            case 9223372036854775807:
                if true {
                    _ = len(v3)
                    var v55 uint8 = (^(v2))
                    var v56 int64 = 1
                    _ = v55
                    _ = v56
                } else {
                    v4 = v1
                    v57 := GenS7{F0: 127}
                    _ = v57
                }
            case 32767:
                for fi58 := 0; fi58 < 19; fi58++ {
                    switch fi58 {
                    case 4294967295:
                        v2 = 255
                    case 989:
                        if true {
                            v59 := []int{127, 384, (fi58 << ((18446744073709551615) & 63)), (fi58 & fi58)}
                            _ = v59
                        } else {
                            v5 = v5
                            fmt.Println(v1)
                            var v60 string = "\n\t"
                            _ = v60
                        }
                    default:
                        v1 -= 0.0
                    }
                    for fi61 := 0; fi61 < 10; fi61++ {
                        switch fi58 {
                        case -1:
                            v62 := []int{0, (-(fi61)), -840, (fi58 | 0)}
                            _ = v62
                        case 2147483647:
                            var v63 uint8 = v2
                            _ = v63
                        case -22:
                            if (("foo bar" + ((((("foo bar" + "a") + "hello") + "hello") + "hello") + "\n\t")) <= "a") {
                                switch 128 {
                                case 804:
                                    v4 += (v1 * v1)
                                case 128:
                                    v64 := GenS3{F0: -88}
                                    _ = v64
                                case -358:
                                    v65 := GenS4{F0: 910, F1: false, F2: 49}
                                    _ = v65
                                default:
                                    if (0 >= 128) {
                                        fmt.Println(v5, fi61)
                                    }
                                }
                            } else {
                                switch (fi61 ^ 0) {
                                case 127:
                                    var v66 uint16 = 593
                                    _ = v66
                                }
                                switch (fi58 / (((fi61 >> ((65535) & 63))) | 1)) {
                                case -9223372036854775808:
                                    v67 := GenS5{F0: 255, F1: 224, F2: 0}
                                    _ = v67
                                case 248:
                                    if ((18446744073709551615 != 478) || ((false || ((0.0 == v1) || true)) || false)) {
                                        v5 = (v5 / ((v5) | 1))
                                        var v68 int64 = -1
                                        if false {
                                            v69 := []int{}
                                            _ = v69
                                        } else {
                                            v70 := []int{}
                                            if (580 > 42) {
                                                v1 -= 0.0
                                                v1 = float32(0.0)
                                            } else {
                                                switch -2147483648 {
                                                case 1:
                                                    switch ((("a" + ("foo bar" + ("foo bar" + "\n\t"))) + "a") + (("\n\t" + "hello") + (("" + "foo bar") + "hello"))) {
                                                    case "\n\t":
                                                        if false {
                                                            if true {
                                                                fmt.Println(fi61, v2, fi61)
                                                                v71 := GenS9{F0: 65535, F1: 0, F2: 4294967295}
                                                                if ((v4 + v1) < (3.14159 - float32(((float64(1.5) + 1.5) * float64(v1))))) {
                                                                    var v72 int16 = -660
                                                                    fmt.Println(v2)
                                                                    _ = v72
                                                                }
                                                                _ = v71
                                                            }
                                                            if false {
                                                                v5 = 42
                                                            } else {
                                                                if (261 > 330) {
                                                                    if true {
                                                                        v73 := GenS10{F0: 268, F1: 23, F2: 42}
                                                                        v1 -= v4
                                                                        _ = v73
                                                                    }
                                                                }
                                                            }
                                                            if false {
                                                                if false {
                                                                    if len(v70) >= 2 {
                                                                        v70 = v70[1:len(v70)]
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(fi61, v4, fi61)
                                                    case "a":
                                                        var v74 uint32 = 275
                                                        _ = v74
                                                    }
                                                    fmt.Println(fi61)
                                                    v4 = v1
                                                case 114:
                                                    v4 = v4
                                                default:
                                                    v75 := GenS1{F0: -128}
                                                    _ = v75
                                                }
                                            }
                                            _ = v70
                                        }
                                        _ = v68
                                    } else {
                                        var v76 int32 = (v5 - -2147483648)
                                        v4 = v1
                                        _ = v76
                                    }
                                default:
                                    switch "hello" {
                                    case "\n\t":
                                        v77 := GenS10{F0: uint64(v2), F1: 1, F2: 1}
                                        _ = v77
                                    case "":
                                        var v78 float64 = ((1.5 / (3.14159 - float64(float32(0.0)))) - ((0.0 / 1.5) * ((1e10 - -2.25) * 3.14159)))
                                        _ = v78
                                    case "hello":
                                        fmt.Println(fi61, v2, v1)
                                    default:
                                        switch "hello" {
                                        case "a":
                                            var v79 int64 = -387
                                            _ = v79
                                        case "hello":
                                            v2 = v2
                                        case "foo bar":
                                            v80 := GenS11{F0: 947}
                                            _ = v80
                                        default:
                                            var v81 int32 = 585
                                            _ = v81
                                        }
                                    }
                                }
                            }
                        }
                        fmt.Println(v5, v1)
                        _ = len(v3)
                        switch fi61 {
                        case 1:
                            v1 -= -2.25
                        default:
                            var v82 int32 = v5
                            _ = len(v3)
                            v82 -= v5
                            _ = v82
                        }
                        v1 -= v1
                        _ = fi61
                    }
                    v83 := []int{128, fi58, 127}
                    _ = fi58
                    _ = v83
                }
            }
        }
    } else {
        if true {
            v1 = (v4 * v4)
            var v84 float32 = 1.5
            _ = v84
        }
        switch -9223372036854775808 {
        case 32768:
            v4 = 1.5
        }
    }
    _ = len(v3)
    var v85 int8 = 0
    for fi86 := 0; fi86 < 14; fi86++ {
        for fi87 := 0; fi87 < 2; fi87++ {
            v88 := []int{2147483647, fi87, fi87, 127}
            if false {
                if (!(!true)) {
                    fmt.Println(fi86, fi86)
                    v85 = (v85 * v85)
                }
                _ = len(v88)
                fmt.Println(fi87)
            } else {
                v2 = v2
                var v89 uint = uint(v2)
                var v90 uint = 255
                _ = v89
                _ = v90
            }
            _ = fi87
            _ = v88
        }
        _ = fi86
    }
    var v91 int32 = -1
    for fi92 := 0; fi92 < 5; fi92++ {
        v93 := GenS10{F0: 2, F1: (v85 * v85), F2: 2147483647}
        fmt.Println(v2, v2)
        _ = fi92
        _ = v93
    }
    for fi94 := 0; fi94 < 18; fi94++ {
        if false {
            for fi95 := 0; fi95 < 18; fi95++ {
                if true {
                    var v96 string = "a"
                    v97 := []int{32767, (^(fi94))}
                    _ = v96
                    _ = v97
                } else {
                    var v98 int64 = 0
                    v99 := []int{(fi95 | (fi95 << ((255) & 63))), (fi94 | 128), int(v85), 128}
                    _ = v98
                    _ = v99
                }
                var v100 float64 = 0.0
                _ = fi95
                _ = v100
            }
            fmt.Println(v85)
        }
        v91 -= (v91 / (((v91 + -189)) | 1))
        _ = fi94
    }
    v5 = v91
    v85 = (v85 >> ((255) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v85:", v85)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v85
    _ = v91
}
