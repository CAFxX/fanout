package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int32; F1 int16; F2 int64 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 bool }
func genFn0(p0 int, p1 uint) uint16 {
    if true {
        v7 := GenS1{F0: p1, F1: 2147483647}
        var v8 string = ("\n\t" + "foo bar")
        switch (v8 + (v8 + v8)) {
        case "hello":
            var v9 uint8 = 2
            _ = v9
        default:
            switch v8 {
            case "hello":
                v8 = ("foo bar" + v8)
            case "foo bar":
                v10 := GenS1{F0: p1, F1: 2147483648}
                _ = v10
            }
        }
        _ = v7
        _ = v8
    } else {
        var v11 uint32 = 755
        _ = v11
    }
    if true {
        var v12 float64 = (((1.5 - (1.5 + ((0.1 / float64(1.5)) + (0.1 + 3.14159)))) / (1e10 * ((3.14159 - (0.0 * 1e10)) - 0.0))) / 0.1)
        for fi13 := 0; fi13 < 20; fi13++ {
            v14 := []int{}
            p1 -= 279
            _ = fi13
            _ = v14
        }
        v15 := []int{(p0 & p0), 983, 810, (-(p0))}
        _ = v12
        _ = v15
    } else {
        switch "a" {
        case "foo bar":
            for fi16 := 0; fi16 < 7; fi16++ {
                v17 := GenS1{F0: p1, F1: 0}
                _ = fi16
                _ = v17
            }
        case "a":
            p0 = 128
        }
    }
    _ = p0
    _ = p1
    return 32767
}

func main() {
    v1 := []int{9223372036854775807, -2147483648}
    var v2 uint8 = 239
    var v3 int32 = 42
    v4 := make(map[int]int)
    var v5 uint = 65535
    var v6 int16 = 54
    switch "hello" {
    case "foo bar":
        for fi18 := 0; fi18 < 15; fi18++ {
            switch (fi18 ^ (fi18 << ((uint(v2)) & 63))) {
            case -1:
                switch ("hello" + "") {
                case "\n\t":
                    for fi19 := 0; fi19 < 8; fi19++ {
                        v3 += (v3 >> ((v5) & 63))
                        v5 = v5
                        v20 := v4[42]
                        _ = fi19
                        _ = v20
                    }
                case "hello":
                    var v21 uint64 = 0
                    _ = v21
                case "foo bar":
                    if false {
                        for fi22 := 0; fi22 < 3; fi22++ {
                            var v23 int = fi18
                            _ = fi22
                            _ = v23
                        }
                        var v24 int8 = -1
                        switch fi18 {
                        case -1:
                            var v25 int16 = 0
                            _ = v25
                        case 9223372036854775807:
                            switch 1 {
                            case -392:
                                if false {
                                    var v26 float64 = (0.0 - 0.1)
                                    for fi27 := 0; fi27 < 3; fi27++ {
                                        if (0 < -715) {
                                            if len(v1) > 0 {
                                                v1[0] = fi18
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v28 string = "\n\t"
                                            _ = v28
                                        }
                                        _ = fi27
                                    }
                                    fmt.Println(v2, v3, v26)
                                    _ = v26
                                } else {
                                    switch 127 {
                                    case 0:
                                        for fi29 := 0; fi29 < 17; fi29++ {
                                            var v30 int16 = int16(v24)
                                            _ = fi29
                                            _ = v30
                                        }
                                    case 127:
                                        v31 := GenS2{F0: 0, F1: (v6 * v6), F2: 465}
                                        _ = v31
                                    }
                                    v2 -= v2
                                }
                            case -9223372036854775808:
                                var v32 uint64 = 18446744073709551615
                                _ = v32
                            case 395:
                                v6 = (^(v6))
                            default:
                                var v33 float64 = (0.1 * (float64(1.5) - (-2.25 + ((-2.25 * (1e10 / -2.25)) + -2.25))))
                                _ = v33
                            }
                        default:
                            for fi34 := 0; fi34 < 15; fi34++ {
                                var v35 float64 = 3.14159
                                v36 := GenS1{F0: v5, F1: 2147483648}
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                        }
                        _ = v24
                    } else {
                        var v37 float32 = 3.14159
                        var v38 float64 = 3.14159
                        v6 = -307
                        _ = v37
                        _ = v38
                    }
                default:
                    fmt.Println(fi18, v5, fi18)
                }
            case 32768:
                switch "a" {
                case "hello":
                    if (3.14159 < 0.0) {
                        v2 = (v2 >> ((4294967295) & 63))
                        if true {
                            var v39 int32 = (v3 | (v3 + -731))
                            var v40 int8 = int8(v3)
                            switch ((("hello" + "\n\t") + "\n\t") + "\n\t") {
                            case "foo bar":
                                v39 = (v3 * v3)
                            default:
                                v41 := GenS2{F0: int32(fi18), F1: 0, F2: -775}
                                _ = v41
                            }
                            _ = v39
                            _ = v40
                        } else {
                            for fi42 := 0; fi42 < 10; fi42++ {
                                var v43 int32 = v3
                                var v44 uint8 = uint8(v5)
                                v3 += (v43 << ((v5) & 63))
                                _ = fi42
                                _ = v43
                                _ = v44
                            }
                            switch fi18 {
                            case 42:
                                var v45 uint32 = 42
                                _ = v45
                            case 32768:
                                v5 += (v5 * v5)
                            default:
                                v5 = 255
                            }
                            fmt.Println(v3, v6)
                        }
                        v3 = 42
                    }
                }
            case 115:
                var v46 int64 = 222
                _ = v46
            default:
                if (v2 < v2) {
                    switch fi18 {
                    case 2147483647:
                        v5 -= v5
                    }
                    if (!true) {
                        fmt.Println(fi18, v2)
                        fmt.Println(v2, v6, v2)
                    }
                    for fi47 := 0; fi47 < 2; fi47++ {
                        var v48 float32 = 3.14159
                        _ = fi47
                        _ = v48
                    }
                }
            }
            _ = fi18
        }
    default:
        v2 = 190
    }
    for fi49 := 0; fi49 < 18; fi49++ {
        for fi50 := 0; fi50 < 3; fi50++ {
            if (false || true) {
                v51 := GenS2{F0: (v3 % ((v3) | 1)), F1: (v6 + v6), F2: -443}
                _ = v51
            }
            _ = fi50
        }
        _ = fi49
    }
    for fi52 := 0; fi52 < 10; fi52++ {
        delete(v4, 32767)
        for fi53 := 0; fi53 < 1; fi53++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v54 int32 = v3
            v6 -= (v6 ^ (v6 | -1))
            _ = fi53
            _ = v54
        }
        for fi55 := 0; fi55 < 12; fi55++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if false {
                var v56 int64 = -9223372036854775808
                var v57 int16 = 1
                v58 := GenS3{F0: 756, F1: 2}
                var v59 int16 = 991
                if false {
                    var v60 string = ("a" + "hello")
                    _ = v60
                }
                _ = v56
                _ = v57
                _ = v58
                _ = v59
            }
            switch -704 {
            case 0:
                switch fi52 {
                case 4294967295:
                    var v61 int8 = 42
                    if true {
                        var v62 float64 = 0.1
                        var v63 uint = (v5 & uint(v2))
                        var v64 int = fi55
                        _ = v62
                        _ = v63
                        _ = v64
                    } else {
                        var v65 int64 = 9223372036854775807
                        _ = v65
                    }
                    v6 = v6
                    _ = v61
                case 32767:
                    var v66 int64 = -213
                    _ = v66
                }
            case 32768:
                var v67 uint = (v5 * v5)
                _ = v67
            case -642:
                switch -1 {
                case -972:
                    v6 += v6
                case 2147483647:
                    if (!(!true)) {
                        v68 := GenS3{F0: 2, F1: (v5 / ((v5) | 1))}
                        v6 += (v6 * (v6 - v6))
                        _ = v68
                    } else {
                        fmt.Println(v2, v5)
                    }
                }
            default:
                v2 = (v2 - (v2 | v2))
            }
            _ = fi55
        }
        _ = fi52
    }
    v69 := GenS4{F0: (v2 + (v2 >> (((v5 * (v5 << (((v5 + 703)) & 63)))) & 63))), F1: 1}
    if (1e10 < ((3.14159 - ((0.0 - ((float64(1.5) - 3.14159) / (3.14159 + (3.14159 + 0.1)))) - 0.1)) - -2.25)) {
        var v70 float32 = (0.0 - (1.5 / (-2.25 / 1e10)))
        _ = v70
    }
    if (1 >= 4294967295) {
        fmt.Println(v5)
    } else {
        var v71 uint64 = 4294967295
        _ = v71
    }
    switch -810 {
    case 32768:
        delete(v4, 128)
    }
    v72 := v4[-9223372036854775808]
    if (42 < 4294967295) {
        v73 := GenS5{F0: (((100 <= -1) && (1 <= uint32(v2))) && (!true))}
        _ = v73
    } else {
        _ = len(v1)
    }
    v2 = v2
    v72 = (v72 * (v72 & v72))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v72
}
