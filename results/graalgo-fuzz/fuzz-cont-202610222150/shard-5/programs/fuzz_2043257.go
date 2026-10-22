package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int; F2 uint8 }
type GenS3 struct { F0 int; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int8; F2 uint }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 string; F1 string }
type GenS9 struct { F0 uint32; F1 bool; F2 int16 }
type GenS10 struct { F0 int; F1 int8; F2 uint64 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint32 }
func genFn0(p0 string) int8 {
    switch p0 {
    case "\n\t":
        p0 = "a"
    case "a":
        var v6 float32 = ((1e10 / 1.5) / 0.1)
        _ = v6
    case "hello":
        if false {
            v7 := []int{-1, 32767, 593}
            for fi8 := 0; fi8 < 12; fi8++ {
                p0 = (("" + (p0 + ("\n\t" + p0))) + (p0 + p0))
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = fi8
            }
            var v9 float32 = -2.25
            _ = v7
            _ = v9
        } else {
            v10 := GenS1{F0: -9223372036854775808, F1: 127}
            switch -2147483648 {
            case -897:
                fmt.Println(p0)
            case 2147483647:
                fmt.Println(p0)
            default:
                var v11 uint = 65535
                _ = v11
            }
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 int64 = 42
                v14 := []int{}
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v10
        }
    }
    v15 := GenS2{F0: 0, F1: -609, F2: 0}
    _ = p0
    _ = v15
    return 100
}
func genFn1(p0 string, p1 uint8) uint32 {
    for fi16 := 0; fi16 < 3; fi16++ {
        v17 := []int{(fi16 << ((2) & 63)), fi16, 32767}
        var v18 uint8 = p1
        p1 = (^(v18))
        _ = fi16
        _ = v17
        _ = v18
    }
    p1 = (p1 | (p1 ^ (p1 << ((2) & 63))))
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 0
    var v3 string = "foo bar"
    v4 := []int{9223372036854775807}
    v5 := make(map[uint32]uint32)
    v19 := v5[0]
    var v20 string = "hello"
    var v21 int64 = v2
    switch "foo bar" {
    case "\n\t":
        v2 += 1
    case "foo bar":
        if false {
            if true {
                v22 := GenS1{F0: 32767, F1: 2}
                v23 := GenS2{F0: v19, F1: 4294967295, F2: 183}
                v19 = (v19 / ((genFn1("a", 194)) | 1))
                _ = v22
                _ = v23
            } else {
                var v24 bool = true
                _ = v24
            }
        } else {
            v25 := GenS3{F0: 9223372036854775807, F1: true, F2: 255}
            if (1 != -1) {
                if (!(1 < -1)) {
                    v26 := GenS2{F0: v19, F1: 1, F2: 250}
                    if (32767 != 999) {
                        switch 32767 {
                        case 32768:
                            var v27 int16 = 495
                            _ = v27
                        case 32767:
                            v28 := GenS4{F0: 123, F1: 1, F2: 18446744073709551615}
                            _ = v28
                        }
                        var v29 int8 = 1
                        _ = v29
                    }
                    v20 = v3
                    _ = v26
                }
                _ = len(v5)
                v30 := v5[(^(v19))]
                _ = v30
            }
            fmt.Println(v19)
            _ = v25
        }
    }
    fmt.Println(v19)
    for fi31 := 0; fi31 < 3; fi31++ {
        if (false || true) {
            fmt.Println(v2, v19, v3)
            var v32 float32 = 1e10
            _ = v32
        } else {
            v33 := GenS2{F0: 765, F1: (fi31 / ((fi31) | 1)), F2: 255}
            if false {
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
            }
            if false {
                for fi34 := 0; fi34 < 13; fi34++ {
                    var v35 float64 = 1.5
                    if (0.1 == ((1.5 / -2.25) * 3.14159)) {
                        switch v20 {
                        case "foo bar":
                            if (!false) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                    }
                    v36 := GenS5{F0: int32(v19), F1: "\n\t"}
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                if true {
                    v20 = v3
                    v37 := GenS4{F0: 229, F1: 100, F2: 512}
                    for fi38 := 0; fi38 < 19; fi38++ {
                        v20 = v3
                        var v39 uint8 = uint8(v19)
                        if ((!(0 > 536)) || true) {
                            if (!(v3 < v3)) {
                                var v40 uint32 = v19
                                var v41 uint64 = 18446744073709551615
                                var v42 bool = ((-1 <= -32768) || (-27 != -1))
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                            v43 := GenS6{F0: 117}
                            if (0 >= 446) {
                                var v44 uint = 18446744073709551615
                                _ = v44
                            } else {
                                var v45 bool = (false && true)
                                v46 := GenS7{F0: v3, F1: fi31}
                                v21 += 1
                                _ = v45
                                _ = v46
                            }
                            v5[uint32(v39)] = 0
                            if (826 != 18446744073709551615) {
                                var v47 uint8 = v39
                                var v48 int64 = (-(v21))
                                _ = v47
                                _ = v48
                            } else {
                                if false {
                                    if true {
                                        switch v3 {
                                        case "a":
                                            v20 = (v20 + v3)
                                            v49 := GenS8{F0: (("" + "") + "\n\t"), F1: (v20 + "hello")}
                                            _ = v49
                                        case "hello":
                                            switch (v20 + "\n\t") {
                                            case "":
                                                fmt.Println(v3, fi38)
                                            }
                                        case "\n\t":
                                            switch v3 {
                                            case "hello":
                                                if (3.14159 != float64(1.5)) {
                                                    if (1.5 <= (-2.25 + (float64((-2.25 + (0.0 * 1.5))) + ((1e10 - 0.1) - (3.14159 - (float64(1e10) / 1e10)))))) {
                                                        if len(v4) > 0 {
                                                            v4[0] = fi31
                                                        }
                                                        var v50 float32 = 1.5
                                                        fmt.Println(v3, v2)
                                                        _ = v50
                                                    } else {
                                                        fmt.Println(v2, fi31)
                                                        v51 := GenS9{F0: genFn1(v3, 2), F1: false, F2: 32767}
                                                        if (!((1000 <= 1) || false)) {
                                                            v19 = genFn1(v20, (v39 >> ((855) & 63)))
                                                            if (18446744073709551615 <= 4294967295) {
                                                                v5[(v19 + v19)] = genFn1(v20, (v39 << ((575) & 63)))
                                                                v2 -= (v2 / (((v21 ^ v2)) | 1))
                                                                v3 = v3
                                                            }
                                                            if true {
                                                                if false {
                                                                    v52 := v1[fi38]
                                                                    var v53 bool = false
                                                                    _ = v52
                                                                    _ = v53
                                                                } else {
                                                                    v54 := GenS10{F0: fi31, F1: -86, F2: uint64(v19)}
                                                                    v55 := GenS3{F0: fi38, F1: true, F2: v39}
                                                                    v19 += (v19 >> ((65535) & 63))
                                                                    _ = v54
                                                                    _ = v55
                                                                }
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                if false {
                                                                    if false {
                                                                        v20 = (((v20 + v20) + (v20 + "a")) + v3)
                                                                    }
                                                                }
                                                                var v56 int = fi38
                                                                v57 := GenS11{F0: fi31}
                                                                _ = v56
                                                                _ = v57
                                                            } else {
                                                                v3 = v3
                                                                var v58 uint = 255
                                                                var v59 bool = ((((1.5 + (((-2.25 / float32(float64(float32(0.1)))) * (0.0 / (float32(float64(float32((3.14159 * float64(1e10))))) * -2.25))) * ((0.0 / 0.1) / 1.0))) - -2.25) / float32(3.14159)) < ((1.5 + 0.0) / -2.25))
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                        }
                                                        _ = v51
                                                    }
                                                } else {
                                                    var v60 float64 = (((0.1 * float64((0.0 * 1e10))) + 3.14159) - 1.5)
                                                    var v61 uint8 = v39
                                                    fmt.Println(v20)
                                                    _ = v60
                                                    _ = v61
                                                }
                                            default:
                                                var v62 string = v20
                                                _ = v62
                                            }
                                        default:
                                            var v63 int8 = -52
                                            _ = v63
                                        }
                                        delete(v1, fi31)
                                        v39 = (v39 >> ((1) & 63))
                                        v39 = (^(v39))
                                    }
                                }
                                var v64 uint8 = 1
                                if false {
                                    var v65 int16 = int16(v2)
                                    _ = v65
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                var v66 uint64 = 0
                                if (4294967295 > v66) {
                                    v66 = 42
                                }
                                v67 := GenS12{F0: -297}
                                _ = v64
                                _ = v66
                                _ = v67
                            }
                            _ = v43
                        } else {
                            switch ((v20 + (v20 + (v3 + ""))) + (v3 + v20)) {
                            case "\n\t":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case "a":
                                if true {
                                    v20 = v20
                                    v2 = 9223372036854775807
                                } else {
                                    fmt.Println(v21, v20)
                                    var v68 uint32 = (^(v19))
                                    if false {
                                        v3 = v3
                                        v69 := GenS1{F0: fi38, F1: 127}
                                        _ = v69
                                    }
                                    switch fi38 {
                                    case -9223372036854775808:
                                        v70 := GenS13{F0: -280}
                                        _ = v70
                                    case 9223372036854775807:
                                        if true {
                                            switch (v20 + ((v20 + ((v20 + v3) + v3)) + ("foo bar" + v3))) {
                                            case "hello":
                                                v20 = (("hello" + v20) + v3)
                                            case "\n\t":
                                                v39 = (v39 * (v39 ^ 255))
                                            default:
                                                v71 := GenS12{F0: 1}
                                                _ = v71
                                            }
                                        }
                                    }
                                    _ = v68
                                }
                            }
                        }
                        fmt.Println(fi31)
                        var v72 int32 = 715
                        _ = fi38
                        _ = v39
                        _ = v72
                    }
                    _ = v37
                } else {
                    v21 = v2
                    v73 := GenS11{F0: fi31}
                    _ = v73
                }
                v74 := GenS6{F0: 713}
                _ = v74
            }
            _ = v33
        }
        _ = fi31
    }
    for fi75 := 0; fi75 < 17; fi75++ {
        switch (((v3 + ((v20 + "foo bar") + ("foo bar" + (("\n\t" + "a") + v20)))) + ((v3 + "") + v3)) + v20) {
        case "foo bar":
            v76 := GenS14{F0: 127}
            _ = v76
        case "a":
            var v77 bool = false
            _ = v77
        case "":
            fmt.Println(v19)
        default:
            var v78 float32 = 3.14159
            _ = v78
        }
        for fi79 := 0; fi79 < 11; fi79++ {
            var v80 int = (fi75 & fi75)
            fmt.Println(v2)
            _ = fi79
            _ = v80
        }
        fmt.Println(fi75)
        _ = fi75
    }
    switch -213 {
    case 9223372036854775807:
        v19 -= 2147483647
    case 42:
        fmt.Println(v19, v3, v20)
    default:
        v5[(v19 >> ((0) & 63))] = 65535
    }
    for fi81 := 0; fi81 < 14; fi81++ {
        switch "" {
        case "":
            switch int(v19) {
            case 0:
                if true {
                    var v82 uint16 = 664
                    _ = v82
                } else {
                    var v83 int64 = 9223372036854775807
                    v2 = (^(v21))
                    _ = v83
                }
            case 900:
                switch fi81 {
                case 571:
                    var v84 uint32 = v19
                    _ = v84
                case 1:
                    for fi85 := 0; fi85 < 8; fi85++ {
                        var v86 int32 = -1
                        v86 += (v86 % ((v86) | 1))
                        v87 := GenS11{F0: (fi85 & fi85)}
                        _ = fi85
                        _ = v86
                        _ = v87
                    }
                case 2147483647:
                    v21 = v21
                default:
                    v88 := GenS15{F0: 42}
                    _ = v88
                }
            case -1:
                v19 += (v19 ^ genFn1("foo bar", 42))
            }
        default:
            fmt.Println(v2, v21)
        }
        _ = fi81
    }
    for fi89 := 0; fi89 < 20; fi89++ {
        v20 = v20
        v21 -= v2
        _ = fi89
    }
    var v90 string = v20
    if false {
        var v91 float64 = 0.0
        _ = v91
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v90:", v90)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v90
}
