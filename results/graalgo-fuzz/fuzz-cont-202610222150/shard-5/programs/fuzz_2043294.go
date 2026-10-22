package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 int8; F2 int32 }
type GenS7 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS8 struct { F0 uint32; F1 int16 }
type GenS9 struct { F0 uint64; F1 uint8; F2 uint }
type GenS10 struct { F0 uint32; F1 bool; F2 bool }
func genFn0(p0 int64, p1 bool) uint16 {
    for fi7 := 0; fi7 < 7; fi7++ {
        if p1 {
            v8 := []int{}
            switch int(p0) {
            case 622:
                var v9 uint64 = 0
                _ = v9
            case -498:
                p0 -= p0
            }
            _ = v8
        }
        _ = fi7
    }
    v10 := []int{}
    for fi11 := 0; fi11 < 11; fi11++ {
        if len(v10) > 0 {
            v10[0] = fi11
        }
        fmt.Println(p1, fi11, p0)
        _ = fi11
    }
    var v12 int32 = 1
    var v13 int8 = -1
    _ = p0
    _ = p1
    _ = v10
    _ = v12
    _ = v13
    return 42
}
func genFn1(p0 int8, p1 int32) uint32 {
    var v14 uint32 = 65535
    var v15 uint16 = 42
    v16 := GenS1{F0: 1, F1: false}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return (v14 & v14)
}

func main() {
    v1 := []uint32{278, 506, 996}
    var v2 uint32 = 57
    var v3 uint8 = 127
    var v4 float32 = 3.14159
    var v5 uint = 4294967295
    v6 := []uint32{0}
    v17 := GenS1{F0: 4294967295, F1: true}
    switch -483 {
    case -1:
        if false {
            var v18 int16 = 1000
            _ = v18
        } else {
            var v19 float32 = (v4 / 1.0)
            switch 128 {
            case 0:
                v19 = (v4 * v19)
            case 32768:
                var v20 string = ("" + "hello")
                _ = v20
            }
            _ = v19
        }
    case 42:
        if (430 > int(v2)) {
            if (!false) {
                var v21 bool = (!true)
                v22 := GenS1{F0: 0, F1: (0 >= -649)}
                if false {
                    var v23 string = ""
                    _ = v23
                }
                _ = v21
                _ = v22
            }
            for fi24 := 0; fi24 < 13; fi24++ {
                _ = len(v6)
                v2 = uint32(v3)
                if (false || (-549 >= -2147483648)) {
                    var v25 int8 = 127
                    var v26 uint16 = 42
                    if ((v25 ^ 0) >= (v25 * -128)) {
                        var v27 uint = v5
                        _ = v27
                    }
                    _ = v25
                    _ = v26
                } else {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v28 float64 = float64(v4)
                    _ = v28
                }
                _ = fi24
            }
            fmt.Println(v3, v3)
        } else {
            var v29 uint = (v5 % ((v5) | 1))
            var v30 uint = (^(v29))
            _ = v29
            _ = v30
        }
    }
    switch 127 {
    case 4294967295:
        switch ((("\n\t" + "a") + ("\n\t" + "a")) + "") {
        case "":
            var v31 uint8 = 0
            _ = v31
        default:
            for fi32 := 0; fi32 < 2; fi32++ {
                for fi33 := 0; fi33 < 20; fi33++ {
                    switch 738 {
                    case 42:
                        if (-1 <= 398) {
                            var v34 float64 = 0.0
                            _ = v34
                        }
                    case -9223372036854775808:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if false {
                            v35 := GenS2{F0: "a"}
                            var v36 int64 = 0
                            _ = v35
                            _ = v36
                        } else {
                            switch ((("\n\t" + ("a" + ("foo bar" + "\n\t"))) + (("" + "\n\t") + "foo bar")) + "\n\t") {
                            case "foo bar":
                                v5 = (^(v5))
                                v37 := GenS2{F0: "hello"}
                                _ = v37
                            }
                            v38 := GenS1{F0: 2, F1: (("\n\t" != "\n\t") && true)}
                            v5 += 0
                            if false {
                                var v39 int16 = -649
                                var v40 int64 = -9223372036854775808
                                switch ("a" + ("foo bar" + "foo bar")) {
                                case "foo bar":
                                    var v41 int = -9223372036854775808
                                    _ = v41
                                }
                                _ = v39
                                _ = v40
                            }
                            var v42 uint = (^(v5))
                            _ = v38
                            _ = v42
                        }
                    case 2147483647:
                        v43 := GenS1{F0: 194, F1: true}
                        _ = v43
                    }
                    v2 -= (^(v2))
                    _ = len(v6)
                    _ = fi33
                }
                for fi44 := 0; fi44 < 12; fi44++ {
                    v6 = append(v6, (v2 / (((v2 ^ (v2 % ((uint32(v3)) | 1)))) | 1)))
                    _ = fi44
                }
                _ = fi32
            }
        }
    case 2147483647:
        if true {
            if true {
                if false {
                    fmt.Println(v5, v5, v5)
                    v2 = (v2 * 2147483647)
                } else {
                    v2 += v2
                }
                switch "foo bar" {
                case "\n\t":
                    var v45 uint16 = 65535
                    _ = v45
                }
            }
        } else {
            fmt.Println(v4, v5)
            v3 += v3
        }
    default:
        for fi46 := 0; fi46 < 10; fi46++ {
            switch "hello" {
            case "hello":
                var v47 string = "hello"
                _ = v47
            case "\n\t":
                v48 := GenS1{F0: 1, F1: (true || false)}
                _ = v48
            case "a":
                _ = len(v1)
            }
            var v49 float32 = v4
            fmt.Println(v2, v2)
            _ = fi46
            _ = v49
        }
    }
    if ("a" != "\n\t") {
        if false {
            v5 = v5
            if ((true && (!(!true))) && false) {
                var v50 bool = (-99 == 1)
                _ = v50
            }
            var v51 float32 = 0.0
            _ = v51
        } else {
            switch "a" {
            case "hello":
                for fi52 := 0; fi52 < 9; fi52++ {
                    switch "foo bar" {
                    case "\n\t":
                        if len(v6) > 0 {
                            v6[0] = v2
                        }
                    case "hello":
                        if (1 != -555) {
                            v53 := GenS2{F0: ("foo bar" + "hello")}
                            _ = v53
                        } else {
                            v2 = v2
                            v54 := GenS3{F0: 0, F1: 0}
                            _ = v54
                        }
                    case "a":
                        if ((1 == genFn0(-763, (428 <= 2))) && true) {
                            _ = len(v6)
                            fmt.Println(v5, fi52, v2)
                            var v55 uint8 = v3
                            _ = v55
                        }
                    }
                    _ = len(v1)
                    if len(v1) > 0 {
                        v1[0] = (^(v2))
                    }
                    _ = fi52
                }
            case "foo bar":
                if ((^(v5)) < (^(v5))) {
                    if true {
                        if false {
                            var v56 float32 = v4
                            fmt.Println(v2, v2)
                            var v57 int32 = 296
                            _ = v56
                            _ = v57
                        } else {
                            v58 := GenS4{F0: 812}
                            _ = v58
                        }
                        fmt.Println(v5)
                        v59 := GenS1{F0: 4294967295, F1: false}
                        _ = v59
                    }
                } else {
                    v60 := GenS1{F0: 2, F1: (float32((0.1 * (3.14159 + -2.25))) != (1e10 * 0.1))}
                    _ = len(v1)
                    if (4294967295 == 127) {
                        if len(v6) > 0 {
                            v6[0] = 2
                        }
                        v61 := GenS4{F0: 255}
                        v62 := GenS5{F0: uint(v3)}
                        _ = v61
                        _ = v62
                    } else {
                        switch 127 {
                        case 743:
                            v63 := GenS6{F0: false, F1: 100, F2: -2147483648}
                            _ = v63
                        case 4294967295:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                    }
                    _ = v60
                }
            case "\n\t":
                for fi64 := 0; fi64 < 17; fi64++ {
                    for fi65 := 0; fi65 < 6; fi65++ {
                        fmt.Println(fi64)
                        v2 -= genFn1(100, int32(v5))
                        fmt.Println(v5, v2)
                        var v66 uint16 = 561
                        v67 := GenS5{F0: (^(v5))}
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                    v2 -= 889
                    _ = fi64
                }
            }
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                v68 := GenS3{F0: 599, F1: 32767}
                _ = v68
            case "a":
                v2 = 65535
            case "\n\t":
                for fi69 := 0; fi69 < 10; fi69++ {
                    var v70 uint32 = 2147483647
                    if len(v6) > 0 {
                        v6[0] = (v2 % ((genFn1(int8(v70), int32(v5))) | 1))
                    }
                    _ = fi69
                    _ = v70
                }
            default:
                fmt.Println(v5)
            }
            var v71 string = "foo bar"
            _ = v71
        }
        for fi72 := 0; fi72 < 11; fi72++ {
            v3 += v3
            fmt.Println(v4, v5, fi72)
            if (false || (9223372036854775807 <= 9223372036854775807)) {
                v1 = append(v1, genFn1(0, -110))
                var v73 int8 = 99
                for fi74 := 0; fi74 < 6; fi74++ {
                    v75 := GenS7{F0: 527, F1: 1, F2: v2}
                    _ = len(v6)
                    var v76 int8 = int8(fi72)
                    fmt.Println(v76, v2, fi74)
                    _ = len(v6)
                    _ = len(v6)
                    _ = fi74
                    _ = v75
                    _ = v76
                }
                _ = v73
            }
            _ = fi72
        }
    }
    if (0 > 1) {
        if false {
            var v77 uint16 = genFn0(int64(v3), true)
            v77 += 32767
            _ = v77
        }
        if (true || false) {
            _ = len(v6)
            switch ("a" + ("\n\t" + "foo bar")) {
            case "hello":
                v2 = (v2 ^ (v2 * (v2 & uint32(v5))))
            default:
                v78 := GenS2{F0: "hello"}
                _ = v78
            }
            if len(v6) > 0 {
                v6[0] = (v2 + 649)
            }
        }
    } else {
        var v79 int8 = -1
        var v80 int8 = v79
        v5 = v5
        _ = v79
        _ = v80
    }
    v81 := GenS8{F0: 930, F1: -1}
    var v82 uint64 = 18446744073709551615
    v83 := GenS5{F0: (v5 / ((1) | 1))}
    for fi84 := 0; fi84 < 8; fi84++ {
        v85 := GenS8{F0: genFn1(-1, 0), F1: 32767}
        _ = fi84
        _ = v85
    }
    v2 -= (v2 | (v2 | v2))
    switch 32768 {
    case 0:
        fmt.Println(v3, v82, v5)
    }
    if (!false) {
        v3 = (v3 - v3)
        var v86 int64 = 1
        var v87 uint = uint(v2)
        _ = v86
        _ = v87
    }
    if (!(true && (-536 >= 239))) {
        var v88 uint = v5
        _ = v88
    }
    v4 = v4
    for fi89 := 0; fi89 < 3; fi89++ {
        for fi90 := 0; fi90 < 16; fi90++ {
            var v91 int8 = 127
            _ = fi90
            _ = v91
        }
        fmt.Println(v4, v5, v5)
        _ = fi89
    }
    var v92 int = -9223372036854775808
    v93 := GenS9{F0: (v82 << (((v5 << (((v5 ^ v5)) & 63))) & 63)), F1: 0, F2: v5}
    if ((v3 - v3) != (^(v3))) {
        if true {
            fmt.Println(v82, v5, v5)
            fmt.Println(v82, v3)
            for fi94 := 0; fi94 < 2; fi94++ {
                fmt.Println(v5, v4)
                v95 := GenS4{F0: 0}
                _ = fi94
                _ = v95
            }
        } else {
            switch 0 {
            case -2147483648:
                fmt.Println(v5)
            case 32767:
                for fi96 := 0; fi96 < 2; fi96++ {
                    v3 += v3
                    _ = len(v6)
                    for fi97 := 0; fi97 < 7; fi97++ {
                        v4 = v4
                        _ = fi97
                    }
                    _ = fi96
                }
            case 862:
                var v98 uint64 = (v82 | 4294967295)
                _ = v98
            default:
                for fi99 := 0; fi99 < 8; fi99++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    for fi100 := 0; fi100 < 12; fi100++ {
                        fmt.Println(v82, v5)
                        _ = fi100
                    }
                    _ = fi99
                }
            }
            v101 := GenS3{F0: 1, F1: int(v2)}
            _ = v101
        }
        fmt.Println(v82, v2, v4)
    } else {
        v6 = append(v6, v2)
        v102 := GenS10{F0: (v2 ^ 2), F1: (!(-2147483648 >= -707)), F2: false}
        fmt.Println(v2, v5, v82)
        _ = v102
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v82:", v82)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v82
    _ = v83
    _ = v92
    _ = v93
}
