package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int; F2 uint }
type GenS4 struct { F0 uint32; F1 int16; F2 int64 }
type GenS5 struct { F0 int; F1 uint16; F2 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32; F1 uint; F2 int }
type GenS9 struct { F0 int16; F1 uint }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint; F1 int64; F2 uint32 }
type GenS12 struct { F0 uint32; F1 uint8 }
type GenS13 struct { F0 int16; F1 int16; F2 uint64 }
func genFn0(p0 uint64, p1 uint32) int8 {
    var v7 int16 = 0
    var v8 uint64 = 766
    var v9 uint16 = 0
    v10 := []int{9223372036854775807}
    for fi11 := 0; fi11 < 13; fi11++ {
        if ((^(v9)) != (v9 ^ v9)) {
            for fi12 := 0; fi12 < 6; fi12++ {
                switch ("foo bar" + "hello") {
                case "a":
                    if (574 < v9) {
                        switch ("" + "\n\t") {
                        case "hello":
                            p1 += (p1 << ((850) & 63))
                        default:
                            var v13 bool = (false && (0 != 255))
                            _ = v13
                        }
                    }
                case "hello":
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                }
                var v14 int64 = -9223372036854775808
                _ = fi12
                _ = v14
            }
            switch ("foo bar" + ("hello" + "")) {
            case "foo bar":
                var v15 uint32 = p1
                _ = v15
            }
        } else {
            if (!((461 <= uint(v9)) && ((fi11 * fi11) == -2147483648))) {
                fmt.Println(v7, p1)
                if (-1 > 127) {
                    if false {
                        v16 := GenS1{F0: v9, F1: 133}
                        switch "a" {
                        case "":
                            v8 = p0
                        case "foo bar":
                            v10 = append(v10, fi11)
                        }
                        switch "a" {
                        case "foo bar":
                            var v17 int = fi11
                            _ = v17
                        case "a":
                            v18 := GenS2{F0: 1, F1: 127}
                            _ = v18
                        case "":
                            var v19 float64 = -2.25
                            _ = v19
                        }
                        _ = v16
                    }
                } else {
                    for fi20 := 0; fi20 < 7; fi20++ {
                        var v21 uint16 = 1
                        _ = fi20
                        _ = v21
                    }
                }
                for fi22 := 0; fi22 < 9; fi22++ {
                    v7 += v7
                    if len(v10) > 0 {
                        v10[0] = fi11
                    }
                    var v23 float64 = ((0.1 + (3.14159 * 1e10)) + 1e10)
                    if (1 != 0) {
                        v10 = append(v10, 42)
                        if false {
                            v10 = append(v10, fi22)
                            v8 -= p0
                        } else {
                            v24 := GenS3{F0: 0, F1: fi22, F2: uint(p1)}
                            p0 += (^(v8))
                            _ = v24
                        }
                        if len(v10) > 0 {
                            v10[0] = fi11
                        }
                    }
                    p1 = (p1 + (p1 + (p1 ^ (p1 % ((p1) | 1)))))
                    var v25 string = ("foo bar" + "hello")
                    fmt.Println(v8, fi11)
                    _ = fi22
                    _ = v23
                    _ = v25
                }
            } else {
                v26 := GenS3{F0: 255, F1: int(v9), F2: 0}
                var v27 uint = 4294967295
                if ((v9 ^ (v9 ^ v9)) < (v9 << ((uint(v9)) & 63))) {
                    fmt.Println(fi11)
                    var v28 int16 = 1000
                    v27 += 1
                    _ = v28
                }
                _ = v26
                _ = v27
            }
            var v29 uint16 = (^(v9))
            p1 = 2147483648
            _ = v29
        }
        switch 4294967295 {
        case 32768:
            v8 = 0
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 127
}

func main() {
    var v1 float32 = 0.1
    var v2 uint32 = 11
    v3 := make(map[string]uint32)
    var v4 uint32 = v2
    var v5 uint8 = 128
    var v6 float64 = (1.5 + 1.5)
    if false {
        v1 = v1
        var v30 uint64 = 4294967295
        v31 := GenS4{F0: (v4 ^ v4), F1: -99, F2: 0}
        _ = v30
        _ = v31
    } else {
        v6 = v6
        var v32 uint32 = v4
        v3["\n\t"] = (v2 + (v2 + 2))
        _ = v32
    }
    if (int32(v4) > -2147483648) {
        v1 += float32(v6)
        for fi33 := 0; fi33 < 17; fi33++ {
            v5 = (v5 + (v5 | uint8(v4)))
            v34 := GenS5{F0: fi33, F1: 316, F2: v2}
            _ = fi33
            _ = v34
        }
        var v35 float64 = -2.25
        _ = v35
    } else {
        switch "\n\t" {
        case "a":
            v6 -= float64((1.5 + (1.5 * v1)))
        }
        var v36 int = 302
        _ = v36
    }
    var v37 float64 = v6
    for fi38 := 0; fi38 < 13; fi38++ {
        v39 := GenS2{F0: -118, F1: 42}
        v40 := GenS6{F0: 32767}
        _ = fi38
        _ = v39
        _ = v40
    }
    v41 := v3["a"]
    var v42 int32 = 2147483647
    var v43 int16 = 0
    fmt.Println(v43, v6, v41)
    if true {
        fmt.Println(v37)
    } else {
        if ("foo bar" == "\n\t") {
            v2 = uint32(v5)
        } else {
            v43 = -1
            if true {
                if (255 >= 606) {
                    var v44 float32 = -2.25
                    v3["foo bar"] = (v41 & v4)
                    _ = v44
                } else {
                    v43 += v43
                    v45 := []int{-313, 0}
                    v46 := GenS7{F0: int16(v2)}
                    _ = v45
                    _ = v46
                }
                if (1 == -1) {
                    v47 := GenS6{F0: 42}
                    if (-9223372036854775808 <= 9223372036854775807) {
                        var v48 uint32 = (v2 | v4)
                        var v49 uint64 = 309
                        v50 := []int{2147483647, -74, 42}
                        _ = v48
                        _ = v49
                        _ = v50
                    }
                    _ = v47
                }
            }
            v51 := GenS8{F0: v2, F1: 42, F2: 32767}
            _ = v51
        }
        delete(v3, "hello")
        v52 := GenS9{F0: v43, F1: 944}
        _ = v52
    }
    v41 = 4294967295
    switch ("\n\t" + "a") {
    case "foo bar":
        if true {
            var v53 float64 = v37
            fmt.Println(v42)
            _ = v53
        } else {
            v54 := GenS10{F0: int16(v5)}
            if true {
                if (false || false) {
                    v55 := GenS2{F0: (v42 * v42), F1: v5}
                    v4 = v41
                    for fi56 := 0; fi56 < 9; fi56++ {
                        var v57 int = fi56
                        v58 := GenS11{F0: 17, F1: -97, F2: (v41 & (^(v41)))}
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                    _ = v55
                }
                if false {
                    if false {
                        v59 := GenS2{F0: (v42 >> ((0) & 63)), F1: v5}
                        v60 := []int{}
                        _ = v59
                        _ = v60
                    } else {
                        switch int(v43) {
                        case 9223372036854775807:
                            for fi61 := 0; fi61 < 4; fi61++ {
                                if (1 != 694) {
                                    v62 := GenS7{F0: (v43 >> ((0) & 63))}
                                    _ = v62
                                } else {
                                    v63 := GenS12{F0: v2, F1: 65}
                                    v64 := GenS11{F0: 255, F1: -111, F2: v2}
                                    _ = v63
                                    _ = v64
                                }
                                for fi65 := 0; fi65 < 10; fi65++ {
                                    switch 2147483647 {
                                    case -841:
                                        if (127 <= int8(v43)) {
                                            var v66 int32 = v42
                                            _ = v66
                                        }
                                    case 128:
                                        if (v2 <= v4) {
                                            fmt.Println(fi61, v6, v5)
                                            v67 := []int{int(v2), fi65, fi61}
                                            _ = v67
                                        }
                                    }
                                    switch fi65 {
                                    case -493:
                                        v68 := GenS13{F0: (v43 ^ v43), F1: v43, F2: 720}
                                        _ = v68
                                    }
                                    v5 = v5
                                    if (((v1 + 3.14159) >= ((1.5 * float32(0.0)) * v1)) && false) {
                                        var v69 int16 = v43
                                        v4 -= v2
                                        if true {
                                            var v70 bool = (!((495 >= -9223372036854775808) || true))
                                            var v71 uint = 263
                                            v72 := GenS12{F0: v2, F1: 1}
                                            _ = v70
                                            _ = v71
                                            _ = v72
                                        } else {
                                            _ = len(v3)
                                        }
                                        _ = v69
                                    }
                                    v1 -= float32(v37)
                                    _ = fi65
                                }
                                _ = fi61
                            }
                        case -523:
                            v73 := GenS3{F0: v5, F1: 1, F2: 313}
                            _ = v73
                        default:
                            var v74 uint64 = 0
                            _ = v74
                        }
                        switch "" {
                        case "":
                            v75 := GenS1{F0: 32767, F1: 49}
                            _ = v75
                        }
                    }
                    fmt.Println(v6, v6, v4)
                } else {
                    switch 53 {
                    case 32768:
                        v76 := GenS6{F0: 166}
                        _ = v76
                    case -564:
                        v42 = v42
                    }
                    for fi77 := 0; fi77 < 3; fi77++ {
                        if false {
                            if (true && true) {
                                v1 = 0.0
                                var v78 uint64 = 905
                                _ = v78
                            }
                            if (114 < 1) {
                                if (fi77 > (fi77 - (fi77 & (fi77 >> ((18446744073709551615) & 63))))) {
                                    switch (("foo bar" + "a") + "a") {
                                    case "\n\t":
                                        fmt.Println(fi77, v6, v43)
                                    }
                                    v42 = v42
                                } else {
                                    delete(v3, ("hello" + "hello"))
                                    var v79 uint16 = 1
                                    _ = v79
                                }
                            }
                            v80 := GenS13{F0: (v43 + v43), F1: v43, F2: 4294967295}
                            _ = v80
                        }
                        _ = fi77
                    }
                    v4 = 2147483648
                }
            } else {
                v81 := []int{-1}
                v6 = v6
                _ = v81
            }
            for fi82 := 0; fi82 < 5; fi82++ {
                fmt.Println(fi82, fi82)
                _ = fi82
            }
            _ = v54
        }
    case "":
        if false {
            var v83 uint = 65535
            for fi84 := 0; fi84 < 6; fi84++ {
                if (fi84 > (fi84 / ((fi84) | 1))) {
                    v85 := v3["a"]
                    _ = v85
                }
                for fi86 := 0; fi86 < 7; fi86++ {
                    var v87 uint64 = 588
                    v1 = 0.1
                    _ = fi86
                    _ = v87
                }
                _ = fi84
            }
            var v88 int16 = (v43 - v43)
            _ = v83
            _ = v88
        }
    case "a":
        var v89 float64 = (float64(((v1 / (((v1 * v1) / -2.25) - -2.25)) + float32((v37 - v6)))) / v37)
        _ = v89
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
