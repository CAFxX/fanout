package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint }
func genFn0(p0 bool) int32 {
    var v5 uint32 = 2
    if (!(1e10 < 0.0)) {
        v6 := GenS1{F0: 1, F1: 127, F2: -738}
        p0 = true
        _ = v6
    }
    var v7 string = (("a" + "a") + "a")
    fmt.Println(v5, v5, p0)
    _ = p0
    _ = v5
    _ = v7
    return 42
}
func genFn1(p0 int16, p1 string, p2 uint64) uint64 {
    var v8 int8 = 1
    if true {
        v9 := []int{523, 1}
        var v10 int32 = 858
        _ = v10
        _ = v9
    } else {
        switch (((p1 + (p1 + (p1 + "a"))) + ("a" + (p1 + ("\n\t" + p1)))) + p1) {
        case "a":
            if (genFn0((!(true && true))) <= 477) {
                var v11 int = 128
                var v12 bool = false
                _ = v11
                _ = v12
            } else {
                var v13 float32 = 3.14159
                for fi14 := 0; fi14 < 10; fi14++ {
                    if (0 != 1) {
                        v15 := []int{fi14}
                        v8 += -1
                        _ = v15
                    }
                    _ = fi14
                }
                _ = v13
            }
        case "":
            v16 := []int{0, 4294967295, 621, 128}
            _ = v16
        }
        if (!false) {
            fmt.Println(p1, v8, p2)
        }
        v17 := []int{9223372036854775807, -2147483648}
        _ = v17
    }
    switch p1 {
    case "a":
        var v18 bool = (true && (false || true))
        _ = v18
    case "hello":
        switch 32767 {
        case -564:
            p0 -= -80
        }
    case "\n\t":
        var v19 string = p1
        _ = v19
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p2
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{2147483648, 0}
    v3 := make(map[uint32]bool)
    var v4 uint = 727
    v20 := GenS1{F0: 42, F1: 47, F2: 32767}
    var v21 bool = false
    for fi22 := 0; fi22 < 17; fi22++ {
        for fi23 := 0; fi23 < 15; fi23++ {
            v4 += (v4 / ((2) | 1))
            v24 := GenS2{F0: (18446744073709551615 == 0)}
            _ = fi23
            _ = v24
        }
        var v25 bool = (0.0 > -2.25)
        switch "hello" {
        case "":
            var v26 float64 = 0.0
            _ = v26
        case "a":
            for fi27 := 0; fi27 < 18; fi27++ {
                var v28 float64 = ((float64((-2.25 - (0.1 / (0.0 + 1e10)))) + float64(1.5)) - float64(float32(((3.14159 * (1e10 * 3.14159)) / 1e10))))
                _ = fi27
                _ = v28
            }
        default:
            if v21 {
                switch "a" {
                case "\n\t":
                    v29 := GenS2{F0: v25}
                    _ = v29
                default:
                    for fi30 := 0; fi30 < 8; fi30++ {
                        if v25 {
                            _ = len(v2)
                        }
                        fmt.Println(v25, fi30)
                        _ = fi30
                    }
                }
                v31 := GenS3{F0: int16(v4)}
                _ = v31
            } else {
                if v21 {
                    v21 = (0 < 65535)
                } else {
                    v32 := GenS2{F0: ((v25 && v25) || v21)}
                    v21 = v21
                    _ = v32
                }
            }
        }
        _ = fi22
        _ = v25
    }
    if v21 {
        var v33 string = (("a" + "foo bar") + "foo bar")
        var v34 uint16 = 2
        _ = len(v2)
        _ = v33
        _ = v34
    }
    var v35 bool = (!v21)
    fmt.Println(v4, v35, v4)
    v36 := GenS4{F0: 1}
    var v37 uint16 = 65535
    for fi38 := 0; fi38 < 19; fi38++ {
        v2 = append(v2, 1)
        fmt.Println(fi38, fi38)
        _ = fi38
    }
    switch 32768 {
    case 4294967295:
        var v39 float32 = (((0.1 * (0.1 + 0.0)) - (1e10 / ((((1e10 / 1.0) / (3.14159 / 1.0)) - 0.1) / 3.14159))) + 0.1)
        _ = v39
    case 2147483647:
        if false {
            var v40 uint32 = 42
            _ = v40
        } else {
            _ = len(v2)
            switch ("" + ("" + "\n\t")) {
            case "\n\t":
                var v41 uint = (v4 ^ (v4 * v4))
                _ = v41
            case "hello":
                fmt.Println(v4)
            default:
                var v42 int64 = 1
                _ = v42
            }
            var v43 float64 = (3.14159 * ((1.5 / (0.0 + 3.14159)) - -2.25))
            _ = v43
        }
    case -1:
        switch (("a" + "") + (("" + "a") + "a")) {
        case "hello":
            switch 9223372036854775807 {
            case 4294967295:
                switch ("a" + "") {
                case "hello":
                    for fi44 := 0; fi44 < 16; fi44++ {
                        v21 = (1 >= -81)
                        v4 = (v4 << (((v4 + v4)) & 63))
                        v45 := GenS4{F0: 140}
                        _ = fi44
                        _ = v45
                    }
                case "":
                    v35 = true
                default:
                    v46 := GenS3{F0: 690}
                    _ = v46
                }
            default:
                var v47 string = ("a" + "a")
                _ = v47
            }
        case "foo bar":
            var v48 float64 = (0.1 - (((1e10 / 1.0) + -2.25) / ((-2.25 + 1.5) * 1.5)))
            _ = v48
        case "":
            var v49 uint16 = v37
            _ = v49
        default:
            v35 = v35
        }
    default:
        switch "" {
        case "":
            switch 42 {
            case 127:
                var v50 int64 = 0
                _ = v50
            case 9223372036854775807:
                var v51 float32 = 0.1
                _ = v51
            case 0:
                v52 := GenS4{F0: 0}
                _ = v52
            default:
                var v53 int16 = -1
                _ = v53
            }
        case "a":
            switch ("foo bar" + ("hello" + ((("hello" + "a") + ("\n\t" + "hello")) + ""))) {
            case "foo bar":
                var v54 int = -9223372036854775808
                _ = v54
            default:
                if (0 < 0) {
                    fmt.Println(v21, v21, v35)
                } else {
                    fmt.Println(v21, v35, v37)
                    v55 := GenS1{F0: 742, F1: 154, F2: 1}
                    _ = v55
                }
            }
        default:
            if true {
                v56 := GenS4{F0: 0}
                _ = v56
            } else {
                v37 -= (v37 - 0)
            }
        }
    }
    for fi57 := 0; fi57 < 5; fi57++ {
        v21 = (-1 == -1)
        var v58 uint32 = 4294967295
        _ = fi57
        _ = v58
    }
    var v59 uint8 = 128
    if true {
        if v21 {
            v37 -= v37
            v60 := GenS2{F0: v21}
            fmt.Println(v21, v37, v35)
            _ = v60
        } else {
            var v61 uint16 = (v37 >> ((v4) & 63))
            _ = v61
        }
        v62 := GenS5{F0: v4}
        _ = v62
    } else {
        if len(v2) > 0 {
            v2[0] = 0
        }
        switch ("foo bar" + ("a" + "a")) {
        case "a":
            var v63 uint32 = 195
            _ = v63
        }
    }
    for fi64 := 0; fi64 < 20; fi64++ {
        v35 = (4294967295 != v4)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (false || true) {
            v21 = v21
            for fi65 := 0; fi65 < 12; fi65++ {
                var v66 uint = v4
                if (!(false && (0 < 42))) {
                    v67 := GenS4{F0: 1}
                    _ = v67
                } else {
                    if (0 == 2147483647) {
                        switch "hello" {
                        case "hello":
                            var v68 int32 = genFn0(((float64(((3.14159 * (float32(3.14159) * 0.1)) - (float32((float64(float32(1.5)) - 0.0)) + 1e10))) + 0.1) >= (3.14159 + 0.1)))
                            _ = v68
                        case "a":
                            v59 -= v59
                        default:
                            var v69 uint32 = 943
                            _ = v69
                        }
                    }
                    v70 := GenS1{F0: 843, F1: (v59 ^ (^(v59))), F2: -32768}
                    _ = v70
                }
                _ = fi65
                _ = v66
            }
            for fi71 := 0; fi71 < 18; fi71++ {
                if (9223372036854775807 >= 364) {
                    if (v35 || v21) {
                        var v72 uint16 = (v37 ^ v37)
                        _ = v72
                    }
                    switch "foo bar" {
                    case "foo bar":
                        var v73 int = (fi71 % ((int(v4)) | 1))
                        _ = v73
                    case "\n\t":
                        v4 += v4
                    default:
                        var v74 int64 = 448
                        _ = v74
                    }
                    var v75 int = fi64
                    _ = v75
                }
                v2 = append(v2, uint32(v37))
                _ = fi71
            }
        } else {
            var v76 uint32 = 65535
            _ = v76
        }
        _ = fi64
    }
    var v77 uint = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v59
    _ = v77
}
