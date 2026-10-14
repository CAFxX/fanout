package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 int32; F2 uint }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 uint16; F2 int16 }
type GenS9 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS10 struct { F0 int8; F1 uint; F2 int64 }
type GenS11 struct { F0 uint64; F1 uint16; F2 bool }
type GenS12 struct { F0 int32; F1 uint }
type GenS13 struct { F0 uint16; F1 string }
func genFn0(p0 int, p1 int32, p2 int16) uint {
    if false {
        v5 := GenS1{F0: 255}
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := []int{}
            if false {
                switch "\n\t" {
                case "foo bar":
                    if len(v7) > 0 {
                        v7[0] = p0
                    }
                case "\n\t":
                    p0 = p0
                case "hello":
                    for fi8 := 0; fi8 < 10; fi8++ {
                        v9 := GenS2{F0: "foo bar", F1: 867}
                        _ = fi8
                        _ = v9
                    }
                }
            }
            var v10 uint32 = 2147483647
            _ = fi6
            _ = v10
            _ = v7
        }
        var v11 int = (p0 << ((0) & 63))
        _ = v11
        _ = v5
    }
    if (true || false) {
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int = (fi12 / ((p0) | 1))
            _ = fi12
            _ = v13
        }
        v14 := GenS2{F0: "", F1: 0}
        _ = v14
    }
    var v15 uint32 = 2147483647
    v16 := []int{(p0 + 32768)}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 816
}
func genFn1(p0 uint8, p1 uint16, p2 bool) uint64 {
    for fi17 := 0; fi17 < 17; fi17++ {
        fmt.Println(fi17, fi17)
        v18 := GenS3{F0: 410}
        _ = fi17
        _ = v18
    }
    var v19 float64 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 848
}

func main() {
    var v1 float64 = 1.5
    v2 := []int{673}
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 2
    var v20 uint = 2
    if (((("\n\t" + ("a" + (("hello" + "") + "\n\t"))) + "") + "hello") >= "a") {
        for fi21 := 0; fi21 < 9; fi21++ {
            _ = len(v2)
            if false {
                var v22 uint8 = 255
                _ = v22
            }
            _ = fi21
        }
        if ((!false) || true) {
            v23 := GenS1{F0: 42}
            _ = v23
        }
    }
    switch "hello" {
    case "hello":
        if (0 <= 0) {
            if len(v2) > 0 {
                v2[0] = 350
            }
            for fi24 := 0; fi24 < 15; fi24++ {
                for fi25 := 0; fi25 < 10; fi25++ {
                    v26 := GenS4{F0: 0, F1: -2147483648, F2: 42}
                    fmt.Println(v4)
                    var v27 int64 = -325
                    v1 -= v1
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                _ = fi24
            }
            for fi28 := 0; fi28 < 3; fi28++ {
                if len(v2) > 0 {
                    v2[0] = 32767
                }
                _ = fi28
            }
        }
    case "":
        switch "hello" {
        case "\n\t":
            v29 := GenS5{F0: 0, F1: 2147483647}
            _ = v29
        default:
            v20 -= 0
        }
    default:
        var v30 uint16 = 65535
        _ = v30
    }
    v1 = v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        if (0.1 >= v1) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi31 := 0; fi31 < 19; fi31++ {
                if false {
                    switch (fi31 + 629) {
                    case 127:
                        v32 := GenS1{F0: 375}
                        _ = v32
                    case -1:
                        v33 := GenS6{F0: 1}
                        _ = v33
                    }
                    var v34 int32 = 42
                    var v35 float64 = float64(3.14159)
                    _ = v34
                    _ = v35
                } else {
                    switch (fi31 / ((32768) | 1)) {
                    case 4294967295:
                        delete(v3, (v4 << (((^(v20))) & 63)))
                    case -1:
                        fmt.Println(v20, v1, v20)
                    default:
                        var v36 int8 = 42
                        _ = v36
                    }
                }
                switch (("hello" + "") + "foo bar") {
                case "":
                    var v37 uint32 = 2147483647
                    _ = v37
                case "a":
                    var v38 int16 = 1000
                    _ = v38
                case "hello":
                    fmt.Println(v4)
                default:
                    v3[(v4 ^ (v4 - v4))] = v4
                }
                for fi39 := 0; fi39 < 9; fi39++ {
                    var v40 int8 = 23
                    v20 = genFn0((^(fi39)), 42, -32768)
                    switch (fi39 << ((v20) & 63)) {
                    case 4294967295:
                        v20 -= 1
                    case 621:
                        v20 -= (v20 % (((v20 / ((255) | 1))) | 1))
                    }
                    v20 += 2
                    _ = fi39
                    _ = v40
                }
                _ = fi31
            }
            if (uint64(v20) >= 0) {
                for fi41 := 0; fi41 < 6; fi41++ {
                    var v42 float32 = 3.14159
                    v43 := GenS7{F0: -1}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            } else {
                for fi44 := 0; fi44 < 11; fi44++ {
                    v20 = uint(v4)
                    if true {
                        for fi45 := 0; fi45 < 7; fi45++ {
                            fmt.Println(fi45, v1)
                            if false {
                                v46 := GenS8{F0: fi45, F1: 1, F2: 1}
                                if (false || true) {
                                    fmt.Println(fi45)
                                    if (fi44 != fi45) {
                                        v1 += float64((0.0 * float32(v1)))
                                    } else {
                                        _ = len(v3)
                                        var v47 int8 = 54
                                        _ = v47
                                    }
                                } else {
                                    if true {
                                        var v48 int8 = 1
                                        _ = v48
                                    } else {
                                        if false {
                                            switch "foo bar" {
                                            case "\n\t":
                                                if ((-121 < -534) || (1000 < -1)) {
                                                    v49 := v3[143]
                                                    v50 := GenS1{F0: 0}
                                                    _ = v49
                                                    _ = v50
                                                } else {
                                                    v2 = append(v2, fi44)
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                            case "hello":
                                                v51 := GenS1{F0: 859}
                                                var v52 uint = v20
                                                switch "hello" {
                                                case "\n\t":
                                                    v52 = (v20 ^ v52)
                                                case "foo bar":
                                                    var v53 uint64 = 42
                                                    _ = v53
                                                case "a":
                                                    v54 := GenS5{F0: 535, F1: -1}
                                                    _ = v54
                                                default:
                                                    var v55 uint = v20
                                                    _ = v55
                                                }
                                                _ = v51
                                                _ = v52
                                            }
                                            var v56 uint16 = 1
                                            _ = v56
                                        } else {
                                            switch 4294967295 {
                                            case 1:
                                                v4 = (v4 * v4)
                                            case -2147483648:
                                                v57 := GenS9{F0: 831, F1: 49, F2: 32767}
                                                _ = v57
                                            case -458:
                                                if ((-181 != -206) && (-824 != -1)) {
                                                    v1 = 0.1
                                                } else {
                                                    var v58 uint64 = 1
                                                    v59 := GenS10{F0: 127, F1: (v20 / (((v20 >> (((v20 / (((v20 << ((v20) & 63))) | 1))) & 63))) | 1)), F2: -9223372036854775808}
                                                    _ = v58
                                                    _ = v59
                                                }
                                            default:
                                                v20 = genFn0((^(fi45)), 259, 32767)
                                            }
                                        }
                                        var v60 int16 = 1
                                        _ = v60
                                    }
                                }
                                _ = v46
                            } else {
                                fmt.Println(fi44)
                            }
                            _ = fi45
                        }
                        for fi61 := 0; fi61 < 7; fi61++ {
                            fmt.Println(v20, v1)
                            if true {
                                if true {
                                    if (100 >= 42) {
                                        var v62 int8 = 100
                                        _ = v62
                                    }
                                } else {
                                    v3[(v4 - v4)] = v4
                                    var v63 int16 = 1
                                    v64 := GenS11{F0: 42, F1: 264, F2: false}
                                    _ = v63
                                    _ = v64
                                }
                                var v65 int16 = 1000
                                _ = v65
                            }
                            _ = fi61
                        }
                    }
                    _ = fi44
                }
            }
        } else {
            v20 += (v20 ^ (v20 | v20))
            v2 = append(v2, -9223372036854775808)
        }
        if (!(v1 >= 0.0)) {
            if true {
                var v66 int64 = -653
                var v67 int = 4294967295
                v68 := GenS12{F0: -1, F1: v20}
                _ = v66
                _ = v67
                _ = v68
            } else {
                v69 := GenS13{F0: 255, F1: (("\n\t" + "hello") + "hello")}
                _ = v69
            }
            var v70 float32 = 0.0
            _ = v70
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        var v71 int16 = 1
        _ = v71
    }
    for fi72 := 0; fi72 < 18; fi72++ {
        var v73 int = 9223372036854775807
        v74 := v3[(^(v4))]
        var v75 int64 = -9223372036854775808
        _ = fi72
        _ = v73
        _ = v74
        _ = v75
    }
    switch 2147483647 {
    case 42:
        v4 += (v4 << (((v20 >> ((42) & 63))) & 63))
    default:
        v76 := GenS2{F0: "a", F1: -655}
        _ = v76
    }
    v77 := v3[(v4 >> (((v20 & v20)) & 63))]
    fmt.Println(v1)
    fmt.Println(v1, v77, v20)
    v78 := GenS6{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v77
    _ = v78
}
