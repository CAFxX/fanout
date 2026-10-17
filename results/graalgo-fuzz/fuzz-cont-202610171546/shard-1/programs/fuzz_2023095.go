package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS8 struct { F0 uint8; F1 bool }
type GenS9 struct { F0 bool; F1 bool }
type GenS10 struct { F0 int16; F1 int32 }
type GenS11 struct { F0 int32; F1 uint8 }
type GenS12 struct { F0 uint32; F1 int64 }
func genFn0(p0 int32) int16 {
    fmt.Println(p0)
    p0 = p0
    p0 += -1
    v7 := []int{-9223372036854775808, 4294967295, 127}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = v7
    return -32768
}
func genFn1(p0 string) bool {
    switch -415 {
    case 128:
        if true {
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 int8 = 28
                fmt.Println(fi8)
                if (!false) {
                    for fi10 := 0; fi10 < 10; fi10++ {
                        var v11 float64 = float64(float32(1.5))
                        v12 := GenS1{F0: 127, F1: 53}
                        v13 := GenS2{F0: 1}
                        var v14 string = ("hello" + p0)
                        _ = fi10
                        _ = v11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                } else {
                    if (2 > 213) {
                        if (262 != 758) {
                            v15 := GenS3{F0: 2, F1: (((((v9 & (v9 ^ v9)) == (-(v9))) || (true && ((19 >= 18446744073709551615) && (!(true || (!true)))))) && false) || (2 != 18446744073709551615))}
                            _ = v15
                        } else {
                            for fi16 := 0; fi16 < 5; fi16++ {
                                v9 = 69
                                _ = fi16
                            }
                        }
                        v17 := []int{128, 1}
                        p0 = "hello"
                        _ = v17
                    }
                    p0 = "\n\t"
                }
                _ = fi8
                _ = v9
            }
            v18 := GenS2{F0: 0}
            if (808 >= 492) {
                v19 := []int{127, -2147483648, -442}
                p0 = (p0 + (p0 + p0))
                p0 = "a"
                _ = v19
            } else {
                var v20 uint32 = 4294967295
                fmt.Println(v20, p0)
                _ = v20
            }
            _ = v18
        } else {
            var v21 uint8 = 1
            p0 = p0
            fmt.Println(v21, p0)
            _ = v21
        }
    case 460:
        switch p0 {
        case "hello":
            switch 656 {
            case -289:
                for fi22 := 0; fi22 < 8; fi22++ {
                    v23 := []int{(fi22 - fi22), (fi22 << ((2) & 63))}
                    _ = fi22
                    _ = v23
                }
            default:
                v24 := []int{834, 2147483647, 4294967295}
                _ = v24
            }
        default:
            p0 = "foo bar"
        }
    default:
        var v25 bool = false
        _ = v25
    }
    var v26 uint64 = 18446744073709551615
    if ((v26 == 235) && (!(("" + "a") >= p0))) {
        if ((372 == -849) || false) {
            v27 := []int{4294967295, 9223372036854775807}
            _ = v27
        } else {
            switch 42 {
            case 2147483647:
                switch (p0 + "") {
                case "hello":
                    fmt.Println(p0, p0)
                case "foo bar":
                    if (-1 >= 232) {
                        v26 = (v26 ^ 42)
                        if (2147483647 >= 679) {
                            var v28 uint = 418
                            var v29 string = "\n\t"
                            if (genFn0(0) != -32768) {
                                p0 = p0
                            } else {
                                v26 += 2
                                v26 = (v26 & (v26 | 42))
                                if false {
                                    p0 = ("" + ((p0 + ((v29 + "hello") + v29)) + p0))
                                } else {
                                    for fi30 := 0; fi30 < 11; fi30++ {
                                        var v31 int64 = 423
                                        if false {
                                            p0 = (p0 + (p0 + v29))
                                            v32 := []int{fi30, fi30, fi30}
                                            var v33 int = int(v28)
                                            _ = v32
                                            _ = v33
                                        } else {
                                            var v34 uint = 42
                                            v28 = uint(v26)
                                            _ = v34
                                        }
                                        v35 := GenS2{F0: genFn0(42)}
                                        _ = fi30
                                        _ = v31
                                        _ = v35
                                    }
                                }
                            }
                            _ = v28
                            _ = v29
                        }
                    } else {
                        var v36 uint8 = 255
                        _ = v36
                    }
                default:
                    v37 := GenS1{F0: 1, F1: -1}
                    _ = v37
                }
            case 32768:
                var v38 uint64 = v26
                _ = v38
            case 1:
                v39 := GenS4{F0: 1, F1: 0}
                _ = v39
            }
        }
        var v40 int = 42
        _ = v40
    }
    fmt.Println(p0)
    if true {
        v41 := GenS2{F0: 755}
        _ = v41
    }
    _ = p0
    _ = v26
    return false
}

func main() {
    v1 := []int{}
    var v2 string = (("foo bar" + "") + "hello")
    var v3 uint8 = 127
    var v4 string = (v2 + (v2 + ("" + v2)))
    v5 := []uint32{2147483647, 1}
    v6 := []uint32{}
    if (1e10 <= 1e10) {
        v3 -= v3
        fmt.Println(v4, v4, v4)
        if ((!true) || (127 == -103)) {
            var v42 int = 42
            var v43 int = (v42 - (-(v42)))
            v2 = "a"
            _ = v42
            _ = v43
        } else {
            for fi44 := 0; fi44 < 2; fi44++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi44
            }
            for fi45 := 0; fi45 < 13; fi45++ {
                for fi46 := 0; fi46 < 6; fi46++ {
                    v1 = append(v1, (fi45 >> ((1) & 63)))
                    switch fi46 {
                    case 0:
                        v47 := GenS5{F0: ((("a" + v4) + ("foo bar" + v4)) + v4)}
                        _ = v47
                    case 42:
                        fmt.Println(v2)
                    default:
                        var v48 uint8 = v3
                        _ = v48
                    }
                    var v49 bool = false
                    v49 = (-253 >= 0)
                    _ = fi46
                    _ = v49
                }
                var v50 float64 = 0.1
                _ = len(v6)
                _ = fi45
                _ = v50
            }
        }
    }
    var v51 int16 = 1000
    for fi52 := 0; fi52 < 11; fi52++ {
        for fi53 := 0; fi53 < 6; fi53++ {
            v4 = (("hello" + "hello") + v4)
            v2 = v4
            switch fi53 {
            case 715:
                var v54 uint32 = 0
                _ = v54
            case 32768:
                v55 := GenS6{F0: 0, F1: 255}
                _ = v55
            case 127:
                v3 += v3
            }
            if (42 != 459) {
                if true {
                    if false {
                        var v56 uint16 = 42
                        v3 -= v3
                        v3 -= 255
                        switch v4 {
                        case "a":
                            var v57 int16 = (-(v51))
                            _ = v57
                        case "":
                            fmt.Println(fi52, v4, v4)
                        case "foo bar":
                            _ = len(v1)
                        default:
                            v58 := GenS7{F0: 65535, F1: 13, F2: genFn0(845)}
                            _ = v58
                        }
                        if (!false) {
                            fmt.Println(fi52, v3)
                        }
                        _ = v56
                    } else {
                        v59 := GenS1{F0: -14, F1: 0}
                        _ = v59
                    }
                    if len(v1) > 0 {
                        v1[0] = 42
                    }
                    v51 = (^(v51))
                    switch (fi52 & (fi52 & fi52)) {
                    case 9223372036854775807:
                        var v60 float64 = float64(float32((((0.0 - float64(3.14159)) + (3.14159 * (0.1 / 1.5))) / 3.14159)))
                        _ = v60
                    case 32767:
                        var v61 float64 = 0.1
                        _ = v61
                    case -2147483648:
                        var v62 uint64 = 0
                        _ = v62
                    default:
                        switch (("a" + v2) + "foo bar") {
                        case "":
                            var v63 string = (v2 + v4)
                            _ = v63
                        default:
                            v64 := GenS8{F0: (v3 >> ((18446744073709551615) & 63)), F1: ((1 < 1) || (!false))}
                            _ = v64
                        }
                    }
                    if (false && (128 != fi53)) {
                        v1 = append(v1, fi52)
                        v65 := GenS5{F0: (v2 + "hello")}
                        v66 := GenS9{F0: (false || (v3 < (^(v3)))), F1: true}
                        _ = v65
                        _ = v66
                    }
                } else {
                    switch fi52 {
                    case 576:
                        _ = len(v6)
                    case 128:
                        fmt.Println(v4)
                    case 0:
                        v51 = v51
                    default:
                        if (0 <= 10) {
                            if true {
                                if len(v5) > 0 {
                                    v5[0] = 723
                                }
                            }
                        } else {
                            v2 = "\n\t"
                            v2 = v4
                        }
                    }
                    switch 4294967295 {
                    case -2147483648:
                        if false {
                            var v67 int16 = (v51 * genFn0(0))
                            v68 := GenS10{F0: genFn0(42), F1: 129}
                            v67 += (^(v67))
                            _ = v67
                            _ = v68
                        }
                    case 32768:
                        v69 := GenS3{F0: 2, F1: true}
                        _ = v69
                    case 1:
                        v51 += -1
                    }
                }
                fmt.Println(v4)
            } else {
                v5 = append(v5, 145)
                var v70 uint64 = 356
                var v71 uint64 = v70
                _ = v70
                _ = v71
            }
            _ = fi53
        }
        switch (fi52 | fi52) {
        case -486:
            for fi72 := 0; fi72 < 10; fi72++ {
                switch v4 {
                case "":
                    v2 = ("a" + v2)
                case "a":
                    if true {
                        if (65535 >= 0) {
                            var v73 bool = true
                            v73 = v73
                            _ = v73
                        }
                    }
                case "hello":
                    v2 = v4
                }
                if ((570 == (-(v51))) && false) {
                    if true {
                        if (-1 == -1) {
                            var v74 uint = 42
                            v74 = v74
                            if true {
                                v75 := GenS11{F0: 42, F1: v3}
                                _ = v75
                            }
                            _ = v74
                        }
                        var v76 int32 = 2147483647
                        var v77 uint64 = 0
                        _ = v76
                        _ = v77
                    } else {
                        v2 = v2
                        var v78 int8 = -107
                        var v79 uint = 65535
                        _ = v78
                        _ = v79
                    }
                    var v80 int32 = 2147483647
                    if false {
                        if false {
                            fmt.Println(v51, v4)
                        }
                    }
                    _ = v80
                } else {
                    v51 += (-(v51))
                    if len(v5) > 0 {
                        v5[0] = 42
                    }
                }
                _ = fi72
            }
        case 289:
            var v81 float32 = (1.5 + 0.0)
            _ = v81
        }
        _ = fi52
    }
    var v82 int64 = -1
    switch 1 {
    case 897:
        v83 := GenS8{F0: v3, F1: true}
        _ = v83
    case -9223372036854775808:
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            for fi84 := 0; fi84 < 3; fi84++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi84
            }
        }
    default:
        for fi85 := 0; fi85 < 3; fi85++ {
            v3 += (v3 << ((527) & 63))
            _ = fi85
        }
    }
    var v86 int16 = 0
    fmt.Println(v86)
    v4 = "\n\t"
    v87 := GenS2{F0: (v51 - -648)}
    v88 := GenS2{F0: (v51 * (^(v86)))}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi89 := 0; fi89 < 1; fi89++ {
        v51 += v86
        var v90 bool = false
        _ = fi89
        _ = v90
    }
    switch "foo bar" {
    case "":
        var v91 uint8 = (v3 * (^(v3)))
        _ = v91
    case "\n\t":
        for fi92 := 0; fi92 < 15; fi92++ {
            for fi93 := 0; fi93 < 1; fi93++ {
                switch (("a" + v4) + v4) {
                case "a":
                    var v94 int8 = -1
                    _ = v94
                case "foo bar":
                    v51 -= (v51 & v86)
                default:
                    v95 := GenS10{F0: v86, F1: 42}
                    _ = v95
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi93
            }
            _ = fi92
        }
    }
    for fi96 := 0; fi96 < 17; fi96++ {
        v97 := GenS12{F0: 65535, F1: 42}
        _ = fi96
        _ = v97
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v82:", v82)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v82
    _ = v86
    _ = v87
    _ = v88
}
