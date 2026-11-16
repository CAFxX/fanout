package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint; F1 int32; F2 bool }
type GenS5 struct { F0 int8; F1 int64 }
func genFn0(p0 uint, p1 int16, p2 uint) int {
    var v7 uint16 = 1
    for fi8 := 0; fi8 < 1; fi8++ {
        p0 -= 0
        _ = fi8
    }
    fmt.Println(p1, p0)
    p2 += (p2 & 466)
    v9 := GenS1{F0: 448, F1: 0, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return 2147483647
}
func genFn1(p0 int32, p1 uint64) uint {
    v10 := GenS1{F0: 255, F1: p0, F2: (p0 + 940)}
    for fi11 := 0; fi11 < 12; fi11++ {
        p0 = p0
        for fi12 := 0; fi12 < 10; fi12++ {
            if (3.14159 < 3.14159) {
                var v13 int = -2147483648
                v14 := []int{(^(fi11)), fi11, fi12}
                _ = v13
                _ = v14
            }
            v15 := GenS1{F0: 4294967295, F1: (p0 * (p0 % (((p0 % ((int32(p1)) | 1))) | 1))), F2: (p0 / (((p0 * (p0 + (p0 ^ 0)))) | 1))}
            p1 -= p1
            v16 := []int{}
            _ = fi12
            _ = v15
            _ = v16
        }
        fmt.Println(p0, p1)
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    return 18446744073709551615
}

func main() {
    var v1 int16 = -32768
    var v2 uint = 0
    var v3 int16 = -889
    var v4 bool = true
    var v5 int64 = -9223372036854775808
    var v6 uint = (v2 & (v2 >> ((v2) & 63)))
    v17 := GenS1{F0: (v2 % ((v6) | 1)), F1: -37, F2: 2147483647}
    for fi18 := 0; fi18 < 18; fi18++ {
        v6 = genFn1(1, 388)
        var v19 uint = v6
        _ = fi18
        _ = v19
    }
    var v20 uint32 = 239
    var v21 int32 = -1
    switch ("foo bar" + "foo bar") {
    case "":
        if true {
            var v22 int64 = (v5 % ((-9223372036854775808) | 1))
            _ = v22
        } else {
            if v4 {
                v2 = v6
                v6 -= genFn1(2147483647, 2)
            } else {
                switch 171 {
                case 1:
                    var v23 float32 = 0.0
                    _ = v23
                case 128:
                    if (!(-2.25 < -2.25)) {
                        switch "foo bar" {
                        case "foo bar":
                            v6 = 2
                        default:
                            for fi24 := 0; fi24 < 16; fi24++ {
                                if true {
                                    if (true || (1.5 >= float64((-2.25 - 1e10)))) {
                                        var v25 bool = ((0 > 42) && ((v5 - -9223372036854775808) < 827))
                                        if v4 {
                                            if ((^(v2)) > v6) {
                                                v6 += v2
                                                v26 := GenS1{F0: 0, F1: v21, F2: v21}
                                                _ = v26
                                            } else {
                                                if (!v4) {
                                                    for fi27 := 0; fi27 < 12; fi27++ {
                                                        v28 := []int{fi27, (fi27 / ((fi24) | 1)), 151, fi27}
                                                        v29 := GenS1{F0: v2, F1: (v21 | 2147483647), F2: v21}
                                                        _ = fi27
                                                        _ = v28
                                                        _ = v29
                                                    }
                                                    v25 = false
                                                } else {
                                                    var v30 int8 = 1
                                                    _ = v30
                                                }
                                                v31 := GenS1{F0: v6, F1: (v21 ^ -932), F2: int32(fi24)}
                                                v32 := GenS1{F0: (v2 >> (((v6 | 7)) & 63)), F1: 954, F2: v21}
                                                _ = v31
                                                _ = v32
                                            }
                                        } else {
                                            var v33 string = "\n\t"
                                            _ = v33
                                        }
                                        _ = v25
                                    } else {
                                        if v4 {
                                            for fi34 := 0; fi34 < 17; fi34++ {
                                                switch "a" {
                                                case "hello":
                                                    var v35 uint = v6
                                                    _ = v35
                                                }
                                                v36 := []int{fi34, fi24, (fi34 - fi24)}
                                                if v4 {
                                                    if false {
                                                        if true {
                                                            fmt.Println(v5)
                                                            var v37 int16 = (v3 & (v1 & (v3 - 42)))
                                                            v5 += v5
                                                            v1 = (v3 & (v1 % (((v3 << ((708) & 63))) | 1)))
                                                            _ = v37
                                                        } else {
                                                            var v38 int64 = (v5 ^ v5)
                                                            _ = v38
                                                        }
                                                        if len(v36) >= 2 {
                                                            v36 = v36[1:len(v36)]
                                                        }
                                                        if true {
                                                            if (1 < 42) {
                                                                if v4 {
                                                                    var v39 float64 = float64(1.5)
                                                                    v1 -= int16(v20)
                                                                    _ = v39
                                                                }
                                                                v2 += (v6 % ((v6) | 1))
                                                                var v40 uint = (v2 + (v6 >> ((v6) & 63)))
                                                                var v41 int32 = 42
                                                                _ = v40
                                                                _ = v41
                                                            }
                                                            switch "a" {
                                                            case "":
                                                                v42 := GenS1{F0: (v6 - v2), F1: (^(v21)), F2: (v21 + v21)}
                                                                _ = v42
                                                            default:
                                                                var v43 uint16 = 2
                                                                _ = v43
                                                            }
                                                            switch (fi24 << ((genFn1(v21, 4294967295)) & 63)) {
                                                            case 112:
                                                                fmt.Println(fi34, v3)
                                                            case 32768:
                                                                switch (fi24 << ((2) & 63)) {
                                                                case -317:
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        switch ("foo bar" + ("a" + ("hello" + "hello"))) {
                                                                        case "\n\t":
                                                                            if (226 > 207) {
                                                                                fmt.Println(v1, fi24)
                                                                                if (float64(1.5) <= (((float64(-2.25) - 0.1) * (((3.14159 - 0.1) / 3.14159) - ((1e10 / 1.5) - ((3.14159 + 0.1) * float64(float32(1e10)))))) - float64(0.1))) {
                                                                                    if len(v36) >= 2 {
                                                                                        v36 = v36[1:len(v36)]
                                                                                    }
                                                                                    var v44 uint8 = 235
                                                                                    if false {
                                                                                        switch "a" {
                                                                                        case "foo bar":
                                                                                            v6 = v2
                                                                                        case "":
                                                                                            v3 += 581
                                                                                        }
                                                                                        var v45 uint8 = 206
                                                                                        _ = v45
                                                                                    } else {
                                                                                        v46 := GenS2{F0: 302}
                                                                                        var v47 uint16 = 255
                                                                                        v48 := GenS1{F0: v2, F1: v21, F2: v21}
                                                                                        v49 := GenS1{F0: (^(v2)), F1: v21, F2: -105}
                                                                                        _ = v46
                                                                                        _ = v47
                                                                                        _ = v48
                                                                                        _ = v49
                                                                                    }
                                                                                    _ = v44
                                                                                } else {
                                                                                    if ((fi24 != -2147483648) && (-310 > (v21 * (-(v21))))) {
                                                                                        switch ("" + "a") {
                                                                                        case "foo bar":
                                                                                            if v4 {
                                                                                                v4 = (4294967295 > 1)
                                                                                                switch fi24 {
                                                                                                case -977:
                                                                                                    var v50 int64 = int64(fi24)
                                                                                                    _ = v50
                                                                                                default:
                                                                                                    v2 = (v2 / ((v6) | 1))
                                                                                                }
                                                                                            }
                                                                                        case "\n\t":
                                                                                            v4 = v4
                                                                                        default:
                                                                                            fmt.Println(v5, v2, v6)
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v21 = (-(v21))
                                                                            }
                                                                        case "":
                                                                            v36 = append(v36, (-(fi34)))
                                                                        case "foo bar":
                                                                            if false {
                                                                                if true {
                                                                                    v6 += 929
                                                                                    v3 = (v3 % (((^(v3))) | 1))
                                                                                    v36 = append(v36, fi34)
                                                                                } else {
                                                                                    fmt.Println(v3)
                                                                                }
                                                                                var v51 string = ""
                                                                                v52 := GenS3{F0: 1, F1: 0}
                                                                                _ = v51
                                                                                _ = v52
                                                                            }
                                                                        }
                                                                    default:
                                                                        v6 -= (v2 ^ genFn1((v21 & 962), 624))
                                                                    }
                                                                case 1:
                                                                    var v53 int8 = 45
                                                                    _ = v53
                                                                }
                                                            }
                                                        } else {
                                                            fmt.Println(v6, v1, v2)
                                                        }
                                                        var v54 uint = (v6 ^ (v2 >> ((65535) & 63)))
                                                        _ = v54
                                                    }
                                                    var v55 uint32 = 2
                                                    _ = v55
                                                }
                                                fmt.Println(v4, v6, fi34)
                                                _ = fi34
                                                _ = v36
                                            }
                                            v1 -= v1
                                        }
                                        v56 := GenS2{F0: 32767}
                                        _ = v56
                                    }
                                }
                                v57 := GenS2{F0: 946}
                                if v4 {
                                    switch "hello" {
                                    case "hello":
                                        for fi58 := 0; fi58 < 3; fi58++ {
                                            var v59 uint8 = 82
                                            _ = fi58
                                            _ = v59
                                        }
                                    case "foo bar":
                                        for fi60 := 0; fi60 < 7; fi60++ {
                                            var v61 uint16 = 42
                                            _ = fi60
                                            _ = v61
                                        }
                                    default:
                                        for fi62 := 0; fi62 < 5; fi62++ {
                                            v21 -= (v21 + (-(v21)))
                                            v63 := []int{257, fi24, int(v21), 42}
                                            if len(v63) >= 2 {
                                                v63 = v63[1:len(v63)]
                                            }
                                            var v64 float32 = float32(0.1)
                                            v63 = append(v63, fi62)
                                            _ = fi62
                                            _ = v63
                                            _ = v64
                                        }
                                    }
                                    for fi65 := 0; fi65 < 18; fi65++ {
                                        var v66 uint32 = uint32(v2)
                                        _ = fi65
                                        _ = v66
                                    }
                                } else {
                                    fmt.Println(fi24)
                                }
                                _ = fi24
                                _ = v57
                            }
                        }
                        v67 := GenS4{F0: genFn1((v21 * v21), uint64(v20)), F1: v21, F2: (!v4)}
                        for fi68 := 0; fi68 < 10; fi68++ {
                            v20 = (^(v20))
                            _ = fi68
                        }
                        _ = v67
                    }
                case 42:
                    v5 += v5
                }
                var v69 bool = (v4 && ((3.14159 - (((1.5 - ((0.1 / 1.0) * -2.25)) * 0.0) * 3.14159)) <= float64((-2.25 * -2.25))))
                var v70 int64 = v5
                _ = v69
                _ = v70
            }
            v1 = 1
        }
    case "\n\t":
        v71 := []int{genFn0((v6 % ((1) | 1)), (v1 - 1000), 4294967295), 127, int(v20), -2147483648}
        _ = v71
    }
    var v72 float32 = (1.5 + 1.5)
    var v73 uint16 = 65535
    for fi74 := 0; fi74 < 9; fi74++ {
        if v4 {
            var v75 bool = ((1.5 / float64(-2.25)) <= 0.1)
            v6 = (v2 / ((genFn1(42, 2)) | 1))
            if false {
                var v76 uint64 = 42
                for fi77 := 0; fi77 < 7; fi77++ {
                    v78 := GenS5{F0: 42, F1: v5}
                    if (!(!(1e10 >= 0.1))) {
                        var v79 uint64 = (^(v76))
                        _ = v79
                    }
                    _ = fi77
                    _ = v78
                }
                for fi80 := 0; fi80 < 5; fi80++ {
                    v81 := GenS4{F0: 2, F1: -766, F2: v4}
                    v20 = 1
                    _ = fi80
                    _ = v81
                }
                _ = v76
            }
            _ = v75
        }
        _ = fi74
    }
    var v82 int64 = v5
    for fi83 := 0; fi83 < 15; fi83++ {
        var v84 uint16 = (^(v73))
        var v85 float64 = 0.0
        var v86 uint64 = 73
        _ = fi83
        _ = v84
        _ = v85
        _ = v86
    }
    switch ("" + "a") {
    case "\n\t":
        var v87 uint8 = 84
        _ = v87
    case "foo bar":
        switch 2147483647 {
        case 127:
            v4 = v4
        }
    case "a":
        v3 = (v3 / ((v3) | 1))
    default:
        var v88 uint16 = (^(v73))
        _ = v88
    }
    v89 := GenS4{F0: (^(v6)), F1: (v21 + v21), F2: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
    _ = v73
    _ = v82
    _ = v89
}
