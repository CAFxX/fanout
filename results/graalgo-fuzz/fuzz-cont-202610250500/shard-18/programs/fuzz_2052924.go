package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int32; F1 int32; F2 int64 }
type GenS3 struct { F0 int32; F1 string; F2 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 uint8; F2 uint }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS8 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 int; F1 int16; F2 bool }
func genFn0(p0 int16, p1 int32) bool {
    if true {
        var v6 uint64 = 631
        var v7 int8 = -1
        _ = v6
        _ = v7
    } else {
        if false {
            p0 = p0
            p0 = p0
        } else {
            p0 = (p0 | p0)
        }
    }
    var v8 uint16 = 0
    _ = p0
    _ = p1
    _ = v8
    return (!(false && (true && true)))
}
func genFn1(p0 uint32, p1 uint8) bool {
    switch "\n\t" {
    case "foo bar":
        if (false || true) {
            var v9 int8 = 1
            v10 := GenS1{F0: (p1 >> ((65535) & 63)), F1: 32767}
            p0 = (p0 & p0)
            _ = v10
            _ = v9
        }
    }
    switch "" {
    case "foo bar":
        var v11 string = (("hello" + "") + (("hello" + "hello") + ("hello" + "foo bar")))
        _ = v11
    default:
        v12 := []int{933}
        _ = v12
    }
    var v13 uint8 = (p1 - 1)
    p0 = (^(p0))
    switch 42 {
    case -2:
        if false {
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 uint8 = (v13 >> ((255) & 63))
                if (65535 <= 133) {
                    v16 := []int{}
                    var v17 string = "a"
                    _ = v16
                    _ = v17
                } else {
                    var v18 uint8 = 2
                    if (false || false) {
                        if (4294967295 != 42) {
                            v15 = (v15 * v15)
                        }
                    } else {
                        v19 := []int{(^(fi14)), (fi14 - (fi14 % (((fi14 * 0)) | 1)))}
                        _ = v19
                    }
                    v20 := GenS2{F0: -724, F1: -423, F2: -823}
                    _ = v18
                    _ = v20
                }
                _ = fi14
                _ = v15
            }
        }
    case 9223372036854775807:
        if ("a" != "") {
            var v21 uint16 = 32767
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = v13
    return true
}

func main() {
    var v1 int64 = 0
    var v2 uint64 = 0
    v3 := []string{}
    v4 := make(map[int]uint32)
    var v5 string = "a"
    switch -2147483648 {
    case -2147483648:
        var v22 int32 = 962
        _ = v22
    case 294:
        for fi23 := 0; fi23 < 3; fi23++ {
            v24 := GenS3{F0: -2147483648, F1: v5, F2: 42}
            fmt.Println(v5, v5, v5)
            delete(v4, fi23)
            _ = fi23
            _ = v24
        }
    case 1:
        if (42 == (v2 & (v2 ^ (^(v2))))) {
            var v25 uint8 = 255
            _ = v25
        }
    }
    if (18446744073709551615 >= 4294967295) {
        v26 := GenS4{F0: 4294967295}
        _ = v26
    }
    switch -942 {
    case -391:
        v1 = v1
    case 128:
        if true {
            switch v5 {
            case "foo bar":
                switch 128 {
                case -344:
                    v27 := GenS1{F0: 188, F1: 0}
                    _ = v27
                default:
                    if (false && false) {
                        v28 := GenS3{F0: 1, F1: v5, F2: -127}
                        v2 += (^(v2))
                        _ = len(v4)
                        _ = v28
                    }
                }
            case "\n\t":
                switch "\n\t" {
                case "foo bar":
                    v29 := GenS5{F0: v2, F1: 130, F2: 0}
                    _ = v29
                default:
                    if true {
                        fmt.Println(v1, v5, v1)
                        fmt.Println(v5, v5)
                    }
                }
            case "a":
                v1 = (v1 & v1)
            default:
                var v30 uint32 = 65535
                _ = v30
            }
            for fi31 := 0; fi31 < 12; fi31++ {
                v5 = v5
                if (fi31 > (fi31 & fi31)) {
                    v4[(fi31 + fi31)] = 42
                    for fi32 := 0; fi32 < 11; fi32++ {
                        v5 = "hello"
                        v33 := GenS4{F0: (fi32 + 899)}
                        _ = fi32
                        _ = v33
                    }
                    var v34 uint64 = (^(v2))
                    _ = v34
                } else {
                    v1 -= (v1 ^ v1)
                    var v35 int16 = 205
                    _ = v35
                }
                v5 = "\n\t"
                _ = fi31
            }
            v2 += (v2 - (v2 ^ 0))
        }
    case 9223372036854775807:
        fmt.Println(v5, v2)
    default:
        var v36 uint8 = 169
        _ = v36
    }
    switch "\n\t" {
    case "foo bar":
        if false {
            v37 := GenS1{F0: 255, F1: 1000}
            v38 := GenS6{F0: 212, F1: (v1 - 875)}
            switch -733 {
            case 9223372036854775807:
                if (-128 <= -128) {
                    v1 = -9223372036854775808
                    v5 = (v5 + v5)
                    switch v5 {
                    case "foo bar":
                        fmt.Println(v1)
                    default:
                        for fi39 := 0; fi39 < 9; fi39++ {
                            v1 = -1
                            var v40 float32 = float32(1e10)
                            _ = fi39
                            _ = v40
                        }
                    }
                } else {
                    var v41 float64 = (0.0 + float64((3.14159 + 0.1)))
                    for fi42 := 0; fi42 < 13; fi42++ {
                        fmt.Println(v41)
                        _ = fi42
                    }
                    _ = v41
                }
            case 32767:
                v43 := GenS7{F0: 151, F1: 0, F2: -1}
                _ = v43
            }
            _ = v37
            _ = v38
        } else {
            fmt.Println(v2)
            v5 = v5
            var v44 uint32 = 4294967295
            _ = v44
        }
    case "\n\t":
        var v45 uint32 = 2147483648
        _ = v45
    }
    switch -986 {
    case -1:
        if (!false) {
            v4[-9223372036854775808] = 0
            v46 := GenS8{F0: 127, F1: -128, F2: 255}
            _ = v46
        } else {
            var v47 float64 = -2.25
            _ = v47
        }
    }
    v48 := GenS8{F0: 114, F1: -1, F2: 2}
    if true {
        if (v1 == v1) {
            v2 -= (v2 % ((v2) | 1))
            v1 -= (v1 >> ((18446744073709551615) & 63))
        } else {
            var v49 int32 = 440
            var v50 bool = false
            _ = v49
            _ = v50
        }
        v51 := GenS9{F0: v5, F1: -2147483648}
        v2 = v2
        _ = v51
    }
    v5 = v5
    for fi52 := 0; fi52 < 8; fi52++ {
        for fi53 := 0; fi53 < 11; fi53++ {
            var v54 int32 = -2147483648
            v54 += v54
            var v55 int64 = v1
            var v56 uint32 = 779
            v56 = 65535
            _ = fi53
            _ = v54
            _ = v55
            _ = v56
        }
        _ = fi52
    }
    var v57 bool = (!(!(1 == v2)))
    for fi58 := 0; fi58 < 19; fi58++ {
        if v57 {
            v57 = true
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi58
    }
    var v59 bool = (137 > 20)
    var v60 int64 = (^(v1))
    for fi61 := 0; fi61 < 12; fi61++ {
        if (((0 > 388) && v59) && (0 < 1)) {
            for fi62 := 0; fi62 < 12; fi62++ {
                switch v5 {
                case "hello":
                    fmt.Println(v57)
                default:
                    var v63 uint64 = v2
                    _ = v63
                }
                v1 = (^(v1))
                var v64 uint32 = 4294967295
                v59 = (!(246 >= -672))
                v57 = v59
                var v65 int8 = 9
                var v66 int16 = 1
                _ = fi62
                _ = v64
                _ = v65
                _ = v66
            }
            v67 := GenS10{F0: fi61, F1: 859, F2: true}
            _ = v67
        }
        for fi68 := 0; fi68 < 8; fi68++ {
            var v69 float32 = 1e10
            v3 = append(v3, ((v5 + (v5 + v5)) + ("a" + (v5 + v5))))
            _ = fi68
            _ = v69
        }
        _ = fi61
    }
    var v70 int64 = 1
    switch -455 {
    case -1:
        var v71 int16 = 0
        _ = v71
    case -9223372036854775808:
        v1 += v1
    }
    for fi72 := 0; fi72 < 11; fi72++ {
        _ = len(v3)
        v73 := v4[(^(fi72))]
        var v74 int8 = 127
        _ = fi72
        _ = v73
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v57
    _ = v59
    _ = v60
    _ = v70
}
