package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int64; F1 int8; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint }
func genFn0() uint16 {
    switch "foo bar" {
    case "a":
        var v7 uint8 = 127
        _ = v7
    case "hello":
        var v8 int = 42
        _ = v8
    case "\n\t":
        var v9 uint32 = 65535
        _ = v9
    default:
        switch 127 {
        case 32767:
            for fi10 := 0; fi10 < 6; fi10++ {
                v11 := GenS1{F0: 42, F1: -119}
                var v12 int16 = 42
                fmt.Println(v12)
                _ = fi10
                _ = v11
                _ = v12
            }
        case 1:
            if (!true) {
                v13 := GenS1{F0: 191, F1: 100}
                for fi14 := 0; fi14 < 20; fi14++ {
                    if true {
                        switch "a" {
                        case "a":
                            if (!true) {
                                v15 := GenS2{F0: ("" + "a"), F1: 100}
                                _ = v15
                            } else {
                                var v16 int64 = -9223372036854775808
                                _ = v16
                            }
                        case "\n\t":
                            var v17 int = fi14
                            _ = v17
                        default:
                            v18 := []int{(fi14 * fi14)}
                            _ = v18
                        }
                        v19 := []int{-2147483648, (-(fi14)), (fi14 / ((912) | 1))}
                        _ = v19
                    }
                    _ = fi14
                }
                switch 32768 {
                case -141:
                    v20 := []int{4294967295, 32767, 678, -515}
                    _ = v20
                case 784:
                    if ("a" < ("foo bar" + ("\n\t" + ("a" + "foo bar")))) {
                        v21 := []int{1, -857}
                        var v22 int16 = 0
                        _ = v21
                        _ = v22
                    } else {
                        if (!false) {
                            if ((!false) || (true && false)) {
                                var v23 uint32 = 535
                                v23 += 669
                                _ = v23
                            }
                            for fi24 := 0; fi24 < 13; fi24++ {
                                switch 128 {
                                case 64:
                                    for fi25 := 0; fi25 < 8; fi25++ {
                                        var v26 int32 = 42
                                        var v27 uint16 = 2
                                        _ = fi25
                                        _ = v26
                                        _ = v27
                                    }
                                case 42:
                                    v28 := GenS1{F0: -2147483648, F1: -80}
                                    _ = v28
                                case 13:
                                    var v29 int = (fi24 % ((fi24) | 1))
                                    _ = v29
                                default:
                                    var v30 uint16 = 32767
                                    _ = v30
                                }
                                _ = fi24
                            }
                            var v31 int64 = 42
                            _ = v31
                        }
                    }
                case -2147483648:
                    if (!true) {
                        var v32 uint8 = 2
                        _ = v32
                    } else {
                        var v33 int32 = 1
                        _ = v33
                    }
                default:
                    if (1 == 1) {
                        for fi34 := 0; fi34 < 1; fi34++ {
                            for fi35 := 0; fi35 < 13; fi35++ {
                                v36 := GenS3{F0: 0, F1: -107, F2: 189}
                                _ = fi35
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 2; fi37++ {
                                var v38 int = (fi37 | -284)
                                _ = fi37
                                _ = v38
                            }
                            var v39 int = (fi34 + (fi34 - 128))
                            _ = fi34
                            _ = v39
                        }
                    }
                }
                _ = v13
            }
        }
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        if false {
            fmt.Println(fi40)
            switch "hello" {
            case "hello":
                var v41 uint = 18446744073709551615
                _ = v41
            default:
                for fi42 := 0; fi42 < 12; fi42++ {
                    var v43 uint = 800
                    _ = fi42
                    _ = v43
                }
            }
            switch fi40 {
            case 4294967295:
                if true {
                    var v44 int32 = 0
                    v45 := []int{(fi40 | 32767)}
                    _ = v44
                    _ = v45
                }
            case 9223372036854775807:
                var v46 float32 = (0.0 / 1e10)
                _ = v46
            case 2147483647:
                for fi47 := 0; fi47 < 19; fi47++ {
                    var v48 int8 = -128
                    _ = fi47
                    _ = v48
                }
            default:
                var v49 uint16 = 255
                _ = v49
            }
        }
        for fi50 := 0; fi50 < 3; fi50++ {
            if (((!(false && true)) || (((false || (42 >= 127)) && false) || false)) && false) {
                var v51 int8 = 42
                var v52 uint32 = 2
                fmt.Println(fi50, fi40)
                var v53 bool = (((3.14159 - 1.5) + ((float32(float64(((1e10 + float32((0.1 * -2.25))) * 1.5))) - 1e10) / 1.5)) != 3.14159)
                fmt.Println(v52, fi50)
                v51 -= v51
                _ = v51
                _ = v52
                _ = v53
            }
            var v54 uint64 = 1
            _ = fi50
            _ = v54
        }
        _ = fi40
    }
    switch "a" {
    case "foo bar":
        var v55 uint = 158
        _ = v55
    case "a":
        var v56 float64 = 0.0
        _ = v56
    }
    var v57 int64 = 0
    var v58 int16 = -490
    _ = v57
    _ = v58
    return 496
}
func genFn1(p0 uint8, p1 uint16) int64 {
    v59 := []int{409, 45, -9223372036854775808, 42}
    if (!(!false)) {
        var v60 uint = 65535
        v60 -= (v60 >> (((v60 - uint(p1))) & 63))
        fmt.Println(p0, p1)
        _ = v60
    }
    var v61 float32 = 3.14159
    _ = p0
    _ = p1
    _ = v59
    _ = v61
    return 42
}

func main() {
    var v1 uint32 = 2
    var v2 string = ""
    v3 := make(map[string]uint32)
    var v4 uint = 763
    var v5 string = ""
    v6 := make(map[uint32]string)
    v62 := GenS4{F0: 32768}
    v63 := GenS3{F0: 9223372036854775807, F1: 87, F2: 128}
    v4 = v4
    var v64 string = ("\n\t" + v2)
    if true {
        var v65 uint64 = 4294967295
        _ = v65
    } else {
        switch 128 {
        case 246:
            v64 = ""
        case 303:
            v66 := GenS3{F0: 0, F1: -128, F2: 255}
            _ = v66
        case 9223372036854775807:
            for fi67 := 0; fi67 < 11; fi67++ {
                var v68 int = fi67
                switch v5 {
                case "hello":
                    for fi69 := 0; fi69 < 19; fi69++ {
                        switch (fi67 % ((-1) | 1)) {
                        case 335:
                            if (1e10 <= ((-2.25 + (3.14159 + -2.25)) * -2.25)) {
                                v70 := GenS1{F0: 1, F1: 127}
                                v71 := GenS4{F0: (fi69 | (fi69 | fi67))}
                                _ = v70
                                _ = v71
                            }
                        case 4294967295:
                            v1 += v1
                        default:
                            v72 := GenS2{F0: v5, F1: 1}
                            _ = v72
                        }
                        fmt.Println(v64, fi67, v1)
                        v73 := GenS1{F0: -1, F1: 42}
                        _ = fi69
                        _ = v73
                    }
                case "a":
                    var v74 int = 4294967295
                    _ = v74
                case "\n\t":
                    fmt.Println(v2, v68)
                default:
                    if ((!(100 < 127)) && false) {
                        var v75 float32 = 1.5
                        v6[(v1 % ((v1) | 1))] = ("a" + "\n\t")
                        _ = v75
                    } else {
                        var v76 int32 = 42
                        v4 = v4
                        v77 := GenS3{F0: -9223372036854775808, F1: 0, F2: 127}
                        _ = v76
                        _ = v77
                    }
                }
                for fi78 := 0; fi78 < 19; fi78++ {
                    var v79 int16 = 1
                    _ = fi78
                    _ = v79
                }
                _ = fi67
                _ = v68
            }
        }
        v4 = uint(v1)
    }
    fmt.Println(v64)
    v64 = "a"
    v80 := GenS5{F0: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
    _ = v64
    _ = v80
}
