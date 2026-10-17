package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int16; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 int16, p1 string) uint32 {
    v7 := []int{0, 896}
    v8 := GenS1{F0: 42}
    if true {
        switch p1 {
        case "hello":
            p1 = (("hello" + "foo bar") + "hello")
        case "\n\t":
            for fi9 := 0; fi9 < 15; fi9++ {
                fmt.Println(p1)
                _ = fi9
            }
        default:
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 bool = false
                switch "hello" {
                case "":
                    v11 = (42 < (fi10 + fi10))
                case "hello":
                    if v11 {
                        fmt.Println(v11, fi10)
                        var v12 float32 = (-2.25 * 3.14159)
                        if v11 {
                            var v13 uint64 = 172
                            switch -747 {
                            case 9223372036854775807:
                                var v14 int = -316
                                _ = v14
                            case 957:
                                fmt.Println(p0, v12, p0)
                            case 32768:
                                for fi15 := 0; fi15 < 12; fi15++ {
                                    var v16 uint = 255
                                    if v11 {
                                        fmt.Println(v11, fi15)
                                        if v11 {
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                            var v17 uint = (v16 | 784)
                                            v18 := GenS1{F0: (v13 & (v13 & v13))}
                                            _ = v17
                                            _ = v18
                                        }
                                    }
                                    var v19 int = (fi15 << ((v16) & 63))
                                    _ = fi15
                                    _ = v16
                                    _ = v19
                                }
                            default:
                                var v20 float64 = float64(1e10)
                                _ = v20
                            }
                            _ = v13
                        } else {
                            if len(v7) > 0 {
                                v7[0] = (fi10 ^ (-(fi10)))
                            }
                            v21 := GenS2{F0: 125, F1: (p0 << ((207) & 63)), F2: 183}
                            _ = v21
                        }
                        _ = v12
                    }
                }
                _ = fi10
                _ = v11
            }
        }
        var v22 uint8 = 69
        _ = v22
    } else {
        if len(v7) > 0 {
            v7[0] = 127
        }
        v7 = append(v7, 1)
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int32, p1 uint, p2 uint16) uint16 {
    var v23 bool = ("a" < "foo bar")
    switch "hello" {
    case "foo bar":
        v24 := []int{224, 32767}
        _ = v24
    case "":
        v25 := []int{1, 42}
        _ = v25
    case "\n\t":
        fmt.Println(v23, p1, p1)
    default:
        fmt.Println(p1)
    }
    p1 = p1
    for fi26 := 0; fi26 < 13; fi26++ {
        fmt.Println(p2, p2)
        var v27 float32 = float32((0.1 * 1e10))
        _ = fi26
        _ = v27
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return p2
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]bool)
    var v3 uint64 = 735
    var v4 int8 = 119
    var v5 bool = false
    var v6 int32 = -2147483648
    v28 := GenS1{F0: 329}
    fmt.Println(v4)
    if ((^(v3)) < (v3 & 42)) {
        for fi29 := 0; fi29 < 4; fi29++ {
            var v30 uint64 = 2
            v31 := GenS3{F0: true, F1: 18446744073709551615}
            v5 = v5
            _ = fi29
            _ = v30
            _ = v31
        }
        switch (((v1 + v1) + v1) + v1) {
        case "hello":
            _ = len(v2)
        case "\n\t":
            v32 := GenS3{F0: v5, F1: v3}
            _ = v32
        case "foo bar":
            if true {
                if false {
                    for fi33 := 0; fi33 < 19; fi33++ {
                        v34 := []int{fi33, (fi33 % (((fi33 | 264)) | 1)), fi33, (^(fi33))}
                        _ = fi33
                        _ = v34
                    }
                    v35 := v2[704]
                    var v36 int32 = (v6 / (((v6 * v6)) | 1))
                    _ = v35
                    _ = v36
                }
            } else {
                if v5 {
                    switch 0 {
                    case 9223372036854775807:
                        if v5 {
                            delete(v2, 175)
                            for fi37 := 0; fi37 < 13; fi37++ {
                                var v38 uint = 881
                                v39 := GenS2{F0: -1, F1: -902, F2: 42}
                                if (-244 > 13) {
                                    switch (fi37 ^ fi37) {
                                    case -201:
                                        var v40 int8 = v4
                                        _ = v40
                                    case -623:
                                        for fi41 := 0; fi41 < 15; fi41++ {
                                            v38 -= (^(v38))
                                            switch -1 {
                                            case 42:
                                                var v42 float64 = 1.5
                                                _ = v42
                                            case -1:
                                                var v43 float32 = float32(0.0)
                                                _ = v43
                                            }
                                            _ = fi41
                                        }
                                    default:
                                        v6 = int32(v4)
                                    }
                                }
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                            switch 4294967295 {
                            case -2:
                                fmt.Println(v5, v5, v4)
                            case 32767:
                                v4 = v4
                            case 32768:
                                if true {
                                    v4 += v4
                                }
                            default:
                                for fi44 := 0; fi44 < 1; fi44++ {
                                    switch "" {
                                    case "hello":
                                        switch fi44 {
                                        case 32767:
                                            switch "a" {
                                            case "\n\t":
                                                v45 := GenS2{F0: v4, F1: 1, F2: 191}
                                                _ = v45
                                            default:
                                                var v46 int32 = v6
                                                _ = v46
                                            }
                                        case -646:
                                            v3 += 42
                                        default:
                                            if (v3 < (v3 % (((^(v3))) | 1))) {
                                                v47 := GenS1{F0: v3}
                                                fmt.Println(v6, v1, v6)
                                                v48 := GenS4{F0: 65535}
                                                _ = v47
                                                _ = v48
                                            }
                                        }
                                    case "foo bar":
                                        var v49 uint8 = 128
                                        _ = v49
                                    default:
                                        v50 := GenS2{F0: v4, F1: -32768, F2: 128}
                                        _ = v50
                                    }
                                    _ = fi44
                                }
                            }
                        }
                    case -1:
                        switch 32768 {
                        case -9223372036854775808:
                            var v51 uint = 0
                            _ = v51
                        case 32767:
                            var v52 bool = v5
                            _ = v52
                        }
                    }
                } else {
                    switch 127 {
                    case -239:
                        v53 := GenS5{F0: 4294967295, F1: 255}
                        _ = v53
                    case 4294967295:
                        for fi54 := 0; fi54 < 2; fi54++ {
                            var v55 uint64 = v3
                            switch "hello" {
                            case "a":
                                v6 = (v6 | v6)
                            case "":
                                v56 := []int{fi54, -2147483648, (^(fi54)), fi54}
                                _ = v56
                            case "hello":
                                if v5 {
                                    if (((false || true) || true) || v5) {
                                        for fi57 := 0; fi57 < 17; fi57++ {
                                            var v58 uint16 = 42
                                            var v59 int32 = (^(v6))
                                            _ = fi57
                                            _ = v58
                                            _ = v59
                                        }
                                    }
                                    v55 = (v3 - (v3 % ((18446744073709551615) | 1)))
                                    for fi60 := 0; fi60 < 19; fi60++ {
                                        if ((v6 + (v6 % (((v6 * v6)) | 1))) > 1) {
                                            v61 := []int{fi54, 895, fi60, (fi54 / ((fi60) | 1))}
                                            var v62 uint16 = 591
                                            var v63 float32 = 3.14159
                                            v55 -= v55
                                            v62 = v62
                                            _ = v61
                                            _ = v62
                                            _ = v63
                                        } else {
                                            switch "a" {
                                            case "hello":
                                                _ = len(v2)
                                            }
                                            var v64 bool = v5
                                            _ = v64
                                        }
                                        var v65 uint16 = 314
                                        if false {
                                            v5 = (false || true)
                                            if ((v4 >> ((709) & 63)) == (v4 | v4)) {
                                                delete(v2, 695)
                                                var v66 string = v1
                                                _ = v66
                                            } else {
                                                if (255 <= 4294967295) {
                                                    if v5 {
                                                        var v67 int64 = -1
                                                        _ = v67
                                                    } else {
                                                        var v68 float32 = 0.0
                                                        if true {
                                                            v69 := GenS6{F0: 65535}
                                                            _ = v69
                                                        } else {
                                                            var v70 int8 = (^(v4))
                                                            fmt.Println(fi60)
                                                            v1 = ((v1 + v1) + v1)
                                                            _ = v70
                                                        }
                                                        _ = v68
                                                    }
                                                }
                                            }
                                            var v71 uint64 = (v3 << ((2) & 63))
                                            _ = v71
                                        }
                                        _ = fi60
                                        _ = v65
                                    }
                                } else {
                                    var v72 uint16 = 42
                                    switch (fi54 + 949) {
                                    case 0:
                                        if v5 {
                                            _ = len(v2)
                                            if (!false) {
                                                var v73 int = (fi54 * fi54)
                                                _ = v73
                                            }
                                            v4 = -29
                                        } else {
                                            v6 = (v6 >> ((126) & 63))
                                        }
                                    }
                                    _ = v72
                                }
                            default:
                                v55 += 0
                            }
                            _ = fi54
                            _ = v55
                        }
                    default:
                        switch 42 {
                        case 32768:
                            fmt.Println(v4, v1)
                        case 190:
                            v5 = v5
                        }
                    }
                }
                v5 = true
            }
        default:
            switch v1 {
            case "foo bar":
                v74 := []int{0, 42, 127}
                _ = v74
            case "hello":
                var v75 int8 = -59
                _ = v75
            }
        }
        v76 := []int{-479, -781}
        _ = v76
    } else {
        if (!v5) {
            v3 += v3
            fmt.Println(v1, v3, v5)
            v77 := []int{-241, 32767}
            _ = v77
        }
        var v78 float32 = 1e10
        _ = v78
    }
    fmt.Println(v6, v1)
    if v5 {
        switch v1 {
        case "hello":
            if (v5 || v5) {
                var v79 uint = 344
                var v80 float64 = (1.5 - (1.5 / -2.25))
                _ = v79
                _ = v80
            }
        case "foo bar":
            v81 := []int{177, 42, 9223372036854775807, 0}
            _ = v81
        case "\n\t":
            fmt.Println(v5, v5)
        }
        v3 = (v3 | v3)
        v1 = v1
    } else {
        fmt.Println(v3, v6)
        v4 -= 23
        for fi82 := 0; fi82 < 2; fi82++ {
            v3 = 0
            fmt.Println(v3, v4)
            v2[65535] = (v5 && false)
            _ = fi82
        }
    }
    var v83 uint64 = 327
    v6 += v6
    var v84 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v83
    _ = v84
}
