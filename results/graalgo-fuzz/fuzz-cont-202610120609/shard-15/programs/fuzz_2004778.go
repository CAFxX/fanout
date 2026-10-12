package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 uint; F1 uint32; F2 int }
type GenS7 struct { F0 uint16; F1 int64; F2 uint }
type GenS8 struct { F0 int8; F1 int16; F2 int16 }
type GenS9 struct { F0 uint8; F1 bool }
type GenS10 struct { F0 int8; F1 uint8 }
type GenS11 struct { F0 int64; F1 uint16; F2 int8 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int16; F1 uint8; F2 int64 }
type GenS14 struct { F0 uint32; F1 int8 }
func genFn0(p0 string, p1 int8, p2 uint32) uint {
    v6 := GenS1{F0: 1, F1: 1000, F2: 0}
    var v7 int64 = 1
    if true {
        var v8 uint8 = 0
        fmt.Println(p0, v8)
        var v9 int16 = 32767
        _ = v8
        _ = v9
    }
    var v10 int8 = (p1 % (((p1 ^ -36)) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1() int16 {
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 uint16 = 860
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 17; fi14++ {
            if true {
                var v15 int8 = int8(fi14)
                v15 += (^(v15))
                _ = v15
            }
            var v16 string = "\n\t"
            _ = fi14
            _ = v16
        }
        var v17 uint64 = 505
        fmt.Println(v17, v17)
        _ = fi13
        _ = v17
    }
    var v18 uint = 1
    var v19 int32 = -1
    _ = v18
    _ = v19
    return -626
}

func main() {
    var v1 float32 = (-2.25 * (float32(-2.25) + 0.1))
    var v2 int16 = 0
    v3 := []int{763, 59, -112, 127}
    v4 := make(map[int]bool)
    var v5 uint64 = 18446744073709551615
    for fi20 := 0; fi20 < 2; fi20++ {
        if (fi20 < fi20) {
            v2 += v2
            var v21 int = 42
            switch "a" {
            case "a":
                var v22 uint64 = v5
                _ = v22
            case "":
                for fi23 := 0; fi23 < 13; fi23++ {
                    v24 := GenS2{F0: 1, F1: 2147483648}
                    switch (v21 ^ 589) {
                    case -9223372036854775808:
                        var v25 int8 = 42
                        _ = v25
                    case 0:
                        fmt.Println(v1)
                        var v26 float64 = 1.5
                        _ = v26
                    case 32768:
                        v27 := GenS3{F0: v5, F1: 0, F2: -1}
                        _ = v27
                    }
                    _ = fi23
                    _ = v24
                }
            case "hello":
                fmt.Println(v1, v2, fi20)
            }
            _ = v21
        }
        v28 := GenS4{F0: (fi20 * fi20), F1: 2147483648}
        _ = fi20
        _ = v28
    }
    if len(v3) > 0 {
        v3[0] = -1
    }
    v29 := GenS5{F0: 1000, F1: 207}
    switch "hello" {
    case "foo bar":
        var v30 int64 = -809
        _ = v30
    default:
        if false {
            v5 += v5
            switch 32768 {
            case 128:
                if false {
                    _ = len(v3)
                    for fi31 := 0; fi31 < 3; fi31++ {
                        v3 = append(v3, fi31)
                        if false {
                            var v32 uint8 = 128
                            _ = v32
                        } else {
                            if (255 == 255) {
                                switch ("" + "a") {
                                case "foo bar":
                                    if (false && (false || false)) {
                                        if true {
                                            _ = len(v4)
                                        } else {
                                            var v33 int32 = 42
                                            v1 -= float32(1e10)
                                            _ = v33
                                        }
                                        if true {
                                            fmt.Println(v5, v1, v5)
                                            v34 := GenS3{F0: 18446744073709551615, F1: 0, F2: 0}
                                            for fi35 := 0; fi35 < 17; fi35++ {
                                                if ((v5 & (v5 >> ((478) & 63))) >= v5) {
                                                    if false {
                                                        v2 += int16(fi35)
                                                        var v36 int = fi31
                                                        _ = v36
                                                    }
                                                    v1 -= v1
                                                    var v37 int = fi31
                                                    _ = v37
                                                } else {
                                                    if true {
                                                        var v38 int32 = -1
                                                        _ = v38
                                                    } else {
                                                        _ = len(v4)
                                                        v5 = 4294967295
                                                        switch fi35 {
                                                        case 4294967295:
                                                            var v39 int8 = 1
                                                            _ = v39
                                                        }
                                                        v2 = (v2 & v2)
                                                        var v40 uint16 = 34
                                                        _ = v40
                                                    }
                                                    if (297 > 1) {
                                                        if true {
                                                            var v41 float32 = v1
                                                            fmt.Println(fi35, fi35, v41)
                                                            _ = v41
                                                        }
                                                        v42 := GenS6{F0: 719, F1: 820, F2: 127}
                                                        v2 = genFn1()
                                                        if true {
                                                            var v43 uint = 0
                                                            var v44 int64 = -1
                                                            if false {
                                                                v45 := GenS7{F0: 0, F1: (v44 & v44), F2: v43}
                                                                _ = v45
                                                            }
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                        _ = v42
                                                    } else {
                                                        switch (("" + (("a" + "hello") + "\n\t")) + ("\n\t" + "\n\t")) {
                                                        case "foo bar":
                                                            fmt.Println(v5, v2)
                                                        case "a":
                                                            v46 := GenS8{F0: -117, F1: (^(v2)), F2: (v2 | v2)}
                                                            _ = v46
                                                        case "":
                                                            v47 := GenS9{F0: 2, F1: true}
                                                            _ = v47
                                                        default:
                                                            var v48 uint16 = 2
                                                            fmt.Println(v5, fi31)
                                                            var v49 int64 = 9223372036854775807
                                                            _ = v48
                                                            _ = v49
                                                        }
                                                        fmt.Println(v2, fi35)
                                                        switch fi35 {
                                                        case 65:
                                                            switch "" {
                                                            case "a":
                                                                fmt.Println(v2)
                                                            case "":
                                                                v50 := GenS10{F0: 100, F1: 127}
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                _ = v50
                                                            case "\n\t":
                                                                v51 := GenS11{F0: -9223372036854775808, F1: 286, F2: 42}
                                                                _ = v51
                                                            default:
                                                                v1 += v1
                                                            }
                                                        case 4294967295:
                                                            v52 := GenS4{F0: (fi35 << ((255) & 63)), F1: 4294967295}
                                                            var v53 int8 = 1
                                                            _ = v52
                                                            _ = v53
                                                        default:
                                                            v54 := GenS3{F0: (^(v5)), F1: 2147483648, F2: 9223372036854775807}
                                                            _ = v54
                                                        }
                                                    }
                                                    v3 = append(v3, (fi35 + fi35))
                                                    v5 = (v5 & (v5 + v5))
                                                    if true {
                                                        v1 = v1
                                                        if true {
                                                            switch (fi35 - (fi31 - -568)) {
                                                            case 32767:
                                                                v55 := GenS5{F0: (v2 ^ (v2 + (^(v2)))), F1: 127}
                                                                _ = v55
                                                            case 42:
                                                                v56 := GenS11{F0: -1, F1: 2, F2: 0}
                                                                _ = v56
                                                            }
                                                            v57 := GenS2{F0: 255, F1: 1}
                                                            var v58 int64 = -86
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                        v59 := GenS8{F0: -1, F1: v2, F2: (v2 & v2)}
                                                        if (((123 < 127) || (!true)) && true) {
                                                            v4[(fi35 >> ((356) & 63))] = true
                                                        }
                                                        _ = v59
                                                    }
                                                }
                                                _ = fi35
                                            }
                                            _ = v34
                                        }
                                    } else {
                                        v60 := GenS10{F0: -128, F1: 127}
                                        v61 := GenS5{F0: (v2 - -1), F1: 0}
                                        _ = v60
                                        _ = v61
                                    }
                                case "":
                                    fmt.Println(v1)
                                case "\n\t":
                                    _ = len(v3)
                                }
                                v1 -= (v1 * 3.14159)
                            } else {
                                v5 += (v5 >> ((18446744073709551615) & 63))
                                v1 = (((v1 - v1) * (1e10 / v1)) - -2.25)
                            }
                        }
                        _ = fi31
                    }
                    if true {
                        switch -1 {
                        case 32767:
                            if (v5 < 592) {
                                fmt.Println(v1, v5)
                            }
                        }
                    }
                }
            case 32767:
                fmt.Println(v5, v1, v1)
            case -9223372036854775808:
                _ = len(v3)
            default:
                var v62 uint8 = 167
                _ = v62
            }
        }
    }
    v63 := GenS12{F0: ((((v1 * -2.25) / v1) - v1) <= v1)}
    v5 += (v5 ^ 42)
    var v64 uint = 18446744073709551615
    switch -122 {
    case -1:
        var v65 int = 32768
        _ = v65
    case 42:
        var v66 uint32 = 2147483648
        _ = v66
    case 32767:
        if false {
            v67 := GenS13{F0: genFn1(), F1: 62, F2: -771}
            v2 -= v2
            var v68 string = "a"
            _ = v67
            _ = v68
        }
    default:
        for fi69 := 0; fi69 < 8; fi69++ {
            v2 += v2
            v3 = append(v3, fi69)
            if ("\n\t" >= "\n\t") {
                v70 := GenS10{F0: -128, F1: 42}
                v71 := GenS14{F0: 1, F1: 100}
                v72 := GenS4{F0: fi69, F1: uint32(v5)}
                _ = v70
                _ = v71
                _ = v72
            } else {
                var v73 int32 = 42
                fmt.Println(v64, v5, fi69)
                _ = v73
            }
            _ = fi69
        }
    }
    switch "hello" {
    case "":
        v1 -= ((v1 + (v1 + v1)) / float32((3.14159 / 1e10)))
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "\n\t":
        v74 := GenS7{F0: 65535, F1: 492, F2: 34}
        _ = v74
    }
    v64 = v64
    v75 := GenS14{F0: uint32(v64), F1: 100}
    var v76 uint32 = 2147483647
    v77 := GenS13{F0: v2, F1: 255, F2: 1}
    for fi78 := 0; fi78 < 11; fi78++ {
        if (0 >= 2147483647) {
            v2 = -625
        }
        v79 := GenS11{F0: 150, F1: 231, F2: 103}
        _ = fi78
        _ = v79
    }
    var v80 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v76:", v76)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v63
    _ = v64
    _ = v75
    _ = v76
    _ = v77
    _ = v80
}
