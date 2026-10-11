package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int16 }
type GenS2 struct { F0 uint; F1 int16; F2 uint }
type GenS3 struct { F0 int16; F1 bool; F2 bool }
type GenS4 struct { F0 int8; F1 bool; F2 int64 }
type GenS5 struct { F0 bool; F1 int64; F2 string }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int8; F1 string }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint32; F1 uint }
type GenS13 struct { F0 uint; F1 uint32 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 string }
func genFn0(p0 uint8, p1 uint8, p2 bool) uint64 {
    if ("\n\t" > ("" + "hello")) {
        var v5 int64 = -1
        var v6 int32 = 1
        fmt.Println(p2, p0)
        _ = v5
        _ = v6
    }
    fmt.Println(p1, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 uint32, p1 bool) uint {
    var v7 float64 = 3.14159
    var v8 int = 0
    switch "\n\t" {
    case "a":
        v9 := []int{}
        _ = v9
    case "\n\t":
        var v10 uint64 = 18446744073709551615
        _ = v10
    default:
        var v11 uint32 = (p0 >> ((2) & 63))
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        v7 -= float64(3.14159)
        if true {
            if (false || p1) {
                var v13 float64 = 1e10
                p0 -= (p0 << ((255) & 63))
                v14 := GenS1{F0: 642, F1: "foo bar", F2: 0}
                _ = v13
                _ = v14
            } else {
                v7 -= v7
                v15 := GenS1{F0: -1, F1: "\n\t", F2: 42}
                switch "" {
                case "\n\t":
                    if (!p1) {
                        v16 := GenS2{F0: 811, F1: 1000, F2: 65535}
                        _ = v16
                    } else {
                        fmt.Println(v7)
                    }
                default:
                    var v17 int16 = -1
                    _ = v17
                }
                _ = v15
            }
        } else {
            var v18 bool = p1
            _ = v18
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 string = "a"
    v2 := []int{-1, 0, -9223372036854775808, 4294967295}
    v3 := []string{"foo bar", "\n\t", (v1 + ("foo bar" + v1))}
    var v4 bool = false
    v19 := GenS3{F0: 32767, F1: false, F2: (42 > 698)}
    for fi20 := 0; fi20 < 10; fi20++ {
        if true {
            v4 = (100 == 38)
            v21 := GenS4{F0: -1, F1: v4, F2: 1}
            for fi22 := 0; fi22 < 3; fi22++ {
                fmt.Println(fi22)
                v23 := GenS2{F0: 633, F1: 0, F2: 42}
                var v24 uint8 = 1
                v25 := GenS5{F0: ((1e10 != float64(3.14159)) || (534 >= 2)), F1: 42, F2: v1}
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = v21
        } else {
            v26 := GenS4{F0: 100, F1: (2147483648 != 4294967295), F2: 42}
            if v4 {
                v27 := GenS6{F0: 42, F1: v1}
                v28 := GenS2{F0: 772, F1: 32767, F2: 4294967295}
                switch "a" {
                case "":
                    v29 := GenS7{F0: 103, F1: "hello"}
                    _ = v29
                default:
                    switch (fi20 | 4294967295) {
                    case 0:
                        v1 = (v1 + "a")
                    case 42:
                        v1 = (v1 + v1)
                    case 32768:
                        fmt.Println(v1)
                    }
                }
                _ = v27
                _ = v28
            }
            _ = v26
        }
        v4 = (v4 || ((fi20 % (((fi20 >> ((346) & 63))) | 1)) < (fi20 * fi20)))
        fmt.Println(v4, v1)
        _ = fi20
    }
    _ = len(v2)
    v30 := GenS8{F0: -91}
    var v31 uint32 = 0
    fmt.Println(v1)
    if (v4 && ((genFn1(v31, (42 >= -1)) <= 4294967295) || v4)) {
        v32 := GenS9{F0: 65535}
        var v33 float64 = (1e10 / float64(float32(-2.25)))
        for fi34 := 0; fi34 < 3; fi34++ {
            v31 += (v31 - 1)
            _ = fi34
        }
        _ = v32
        _ = v33
    } else {
        fmt.Println(v4)
        for fi35 := 0; fi35 < 12; fi35++ {
            v4 = v4
            v36 := GenS5{F0: true, F1: 9223372036854775807, F2: v1}
            for fi37 := 0; fi37 < 9; fi37++ {
                var v38 uint32 = (^(v31))
                _ = fi37
                _ = v38
            }
            _ = fi35
            _ = v36
        }
        v39 := GenS3{F0: 42, F1: ((0.0 / 0.1) >= -2.25), F2: true}
        _ = v39
    }
    v31 -= (v31 << ((0) & 63))
    if (65535 == v31) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi40 := 0; fi40 < 10; fi40++ {
            v41 := GenS9{F0: 255}
            _ = fi40
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        fmt.Println(v1)
        _ = fi42
    }
    for fi43 := 0; fi43 < 6; fi43++ {
        fmt.Println(v1, v31, v1)
        var v44 uint = 4294967295
        for fi45 := 0; fi45 < 8; fi45++ {
            switch (fi45 >> ((v44) & 63)) {
            case -548:
                v46 := GenS6{F0: 103, F1: ("" + v1)}
                _ = v46
            case 0:
                v47 := GenS10{F0: -9223372036854775808}
                _ = v47
            default:
                if v4 {
                    switch (fi43 | fi45) {
                    case 9223372036854775807:
                        if v4 {
                            var v48 uint16 = 2
                            var v49 int16 = 752
                            _ = v48
                            _ = v49
                        }
                    default:
                        var v50 uint8 = 128
                        _ = v50
                    }
                } else {
                    v2 = append(v2, (fi45 | (fi43 & (fi43 * fi43))))
                    var v51 string = (v1 + "hello")
                    _ = v51
                }
            }
            var v52 uint64 = 1
            switch ((v1 + v1) + "a") {
            case "hello":
                v53 := GenS11{F0: 55}
                v52 -= (v52 | v52)
                _ = v53
            case "foo bar":
                var v54 float64 = -2.25
                if (v4 || (358 <= 255)) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(fi43)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = v54
            }
            _ = fi45
            _ = v52
        }
        _ = fi43
        _ = v44
    }
    _ = len(v2)
    if (v4 || v4) {
        switch ("\n\t" + "foo bar") {
        case "":
            var v55 string = ""
            _ = v55
        case "foo bar":
            v56 := GenS12{F0: v31, F1: 4294967295}
            _ = v56
        case "a":
            _ = len(v3)
        default:
            fmt.Println(v4)
        }
        var v57 uint32 = (v31 - (v31 / ((v31) | 1)))
        switch (v1 + (((v1 + "foo bar") + (("" + v1) + v1)) + "foo bar")) {
        case "":
            v4 = (1 != 255)
        case "foo bar":
            v57 += v57
        }
        _ = v57
    } else {
        v4 = (float32((1e10 - 0.0)) >= (0.1 / 1.5))
    }
    for fi58 := 0; fi58 < 19; fi58++ {
        v4 = false
        _ = fi58
    }
    for fi59 := 0; fi59 < 6; fi59++ {
        v4 = ((false && (!v4)) || false)
        for fi60 := 0; fi60 < 16; fi60++ {
            v4 = v4
            switch fi59 {
            case 368:
                fmt.Println(v1, v1)
            case 42:
                v4 = false
            }
            if v4 {
                v61 := GenS13{F0: 4294967295, F1: (v31 - 0)}
                fmt.Println(fi60, fi59)
                if v4 {
                    v62 := GenS14{F0: v4}
                    switch (v1 + v1) {
                    case "":
                        var v63 bool = (-128 > 1)
                        _ = v63
                    case "foo bar":
                        var v64 string = "\n\t"
                        _ = v64
                    case "\n\t":
                        v1 = v1
                    default:
                        v65 := GenS15{F0: 104, F1: 2147483647, F2: 4294967295}
                        _ = v65
                    }
                    _ = v62
                }
                fmt.Println(fi60, v4)
                if (false || v4) {
                    switch v1 {
                    case "a":
                        v4 = true
                    case "foo bar":
                        var v66 uint32 = v31
                        var v67 string = (((((v1 + (v1 + v1)) + "foo bar") + v1) + ("" + "")) + ((v1 + "foo bar") + "hello"))
                        if v4 {
                            switch fi59 {
                            case 98:
                                var v68 int = fi60
                                _ = v68
                            case 127:
                                var v69 uint = 1
                                _ = v69
                            }
                        } else {
                            v67 = v1
                            if ((0 > 42) && v4) {
                                if ("\n\t" <= v1) {
                                    v2 = append(v2, (^(fi60)))
                                } else {
                                    if (255 < 255) {
                                        v2 = append(v2, (^(fi60)))
                                        v66 -= 575
                                        v70 := GenS16{F0: v1}
                                        _ = v70
                                    }
                                }
                            } else {
                                var v71 int32 = 1
                                v72 := GenS17{F0: v1}
                                _ = v71
                                _ = v72
                            }
                        }
                        var v73 string = ((v67 + v1) + v1)
                        _ = v66
                        _ = v67
                        _ = v73
                    case "hello":
                        v74 := GenS5{F0: v4, F1: -502, F2: v1}
                        _ = v74
                    }
                    var v75 uint64 = 1
                    _ = v75
                }
                _ = v61
            }
            _ = fi60
        }
        _ = fi59
    }
    for fi76 := 0; fi76 < 14; fi76++ {
        var v77 uint32 = v31
        _ = len(v2)
        v77 += (v31 % (((^(v77))) | 1))
        _ = fi76
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
