package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint32; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 bool }
func genFn0(p0 uint8) bool {
    switch "\n\t" {
    case "foo bar":
        switch ("foo bar" + (("hello" + "\n\t") + "\n\t")) {
        case "\n\t":
            for fi5 := 0; fi5 < 18; fi5++ {
                var v6 float64 = 0.1
                var v7 uint16 = 255
                _ = fi5
                _ = v6
                _ = v7
            }
        }
    case "\n\t":
        fmt.Println(p0)
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 bool = true
        switch ("hello" + ("hello" + ("\n\t" + "foo bar"))) {
        case "a":
            if (v9 || v9) {
                v9 = (p0 >= (p0 << ((1) & 63)))
                v10 := GenS1{F0: 0, F1: 0, F2: 444}
                _ = v10
            } else {
                if (0 <= 42) {
                    v9 = ((fi8 % ((128) | 1)) != 2147483647)
                    v11 := GenS2{F0: -32768, F1: 205, F2: fi8}
                    v9 = true
                    _ = v11
                }
            }
        case "hello":
            v12 := GenS2{F0: 0, F1: 65535, F2: -2147483648}
            _ = v12
        case "foo bar":
            v13 := GenS1{F0: 707, F1: 410, F2: 32767}
            _ = v13
        }
        _ = fi8
        _ = v9
    }
    if (-1 != 26) {
        if false {
            for fi14 := 0; fi14 < 17; fi14++ {
                if (117 == -2) {
                    p0 -= p0
                } else {
                    v15 := []int{(fi14 / (((-(fi14))) | 1)), fi14, fi14}
                    if len(v15) > 0 {
                        v15[0] = 983
                    }
                    _ = v15
                }
                _ = fi14
            }
            var v16 uint8 = p0
            var v17 float64 = 1e10
            _ = v16
            _ = v17
        }
        switch ("foo bar" + "hello") {
        case "foo bar":
            for fi18 := 0; fi18 < 5; fi18++ {
                if (true || false) {
                    v19 := GenS2{F0: 856, F1: 2, F2: 473}
                    if (!(4294967295 < 2)) {
                        fmt.Println(fi18)
                        var v20 int32 = 1
                        _ = v20
                    } else {
                        if (false && (false && false)) {
                            p0 = 127
                            for fi21 := 0; fi21 < 5; fi21++ {
                                var v22 uint16 = 255
                                v23 := GenS3{F0: -570}
                                v22 = 985
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                            p0 = (p0 / ((255) | 1))
                        }
                        p0 += p0
                        for fi24 := 0; fi24 < 15; fi24++ {
                            p0 -= p0
                            _ = fi24
                        }
                    }
                    _ = v19
                }
                var v25 float32 = 1.5
                _ = fi18
                _ = v25
            }
        case "":
            p0 -= p0
        case "\n\t":
            var v26 bool = true
            _ = v26
        default:
            var v27 int32 = 804
            _ = v27
        }
        if (1000 < -859) {
            var v28 int = 128
            _ = v28
        }
    } else {
        for fi29 := 0; fi29 < 3; fi29++ {
            v30 := []int{(fi29 | fi29), fi29, -1}
            var v31 int8 = 42
            _ = fi29
            _ = v30
            _ = v31
        }
        if (-9223372036854775808 == 128) {
            p0 = (p0 / (((p0 * (p0 | (p0 + 128)))) | 1))
        }
    }
    p0 -= (p0 * (p0 | (^(p0))))
    _ = p0
    return (65535 < uint16(p0))
}

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[int]uint32)
    var v3 uint32 = 2147483648
    v4 := make(map[string]bool)
    v32 := GenS4{F0: "", F1: v1}
    var v33 float32 = 1e10
    for fi34 := 0; fi34 < 17; fi34++ {
        v4[("\n\t" + "foo bar")] = true
        _ = fi34
    }
    if (!false) {
        var v35 float32 = v33
        _ = v35
    } else {
        v1 += v1
        var v36 float64 = (1e10 * 0.1)
        _ = v36
    }
    v37 := v2[352]
    v38 := GenS1{F0: v3, F1: 65535, F2: 1}
    var v39 uint8 = 113
    v3 = 4294967295
    switch "\n\t" {
    case "hello":
        for fi40 := 0; fi40 < 13; fi40++ {
            for fi41 := 0; fi41 < 20; fi41++ {
                v33 = v33
                var v42 float32 = ((1.5 * (v33 / v33)) / v33)
                v43 := []int{9223372036854775807, (fi40 ^ (fi40 % ((-845) | 1))), (fi41 * (fi41 % (((fi40 ^ fi40)) | 1)))}
                switch "foo bar" {
                case "foo bar":
                    v42 -= (3.14159 * (((0.1 - v42) / 1e10) + 0.1))
                }
                _ = fi41
                _ = v42
                _ = v43
            }
            var v44 uint16 = 0
            fmt.Println(fi40, v1)
            _ = fi40
            _ = v44
        }
    }
    v45 := GenS5{F0: false}
    v46 := GenS3{F0: 32767}
    if (true || (true && false)) {
        var v47 uint16 = 191
        _ = v47
    } else {
        v37 = 759
    }
    v1 -= v1
    if (32767 == 837) {
        switch 9223372036854775807 {
        case 9223372036854775807:
            switch "\n\t" {
            case "":
                fmt.Println(v1)
            case "a":
                var v48 uint16 = 42
                _ = v48
            case "\n\t":
                v49 := GenS5{F0: false}
                _ = v49
            default:
                v50 := GenS3{F0: int16(v39)}
                _ = v50
            }
        }
        if true {
            var v51 int8 = 1
            v3 = (v37 - (v3 << ((575) & 63)))
            _ = v51
        }
        for fi52 := 0; fi52 < 4; fi52++ {
            var v53 int32 = 2147483647
            fmt.Println(v53)
            for fi54 := 0; fi54 < 2; fi54++ {
                if ((-32768 != int16(v39)) && (((1 != 0) || false) || true)) {
                    if true {
                        if (!(0.0 <= 1.5)) {
                            v55 := []int{}
                            v53 = (^(v53))
                            _ = v55
                        }
                    }
                    fmt.Println(v3, v3, v53)
                } else {
                    v56 := GenS5{F0: (2 < 4294967295)}
                    switch (("a" + (("hello" + "\n\t") + "hello")) + (("hello" + "\n\t") + "\n\t")) {
                    case "a":
                        var v57 string = "\n\t"
                        _ = v57
                    case "hello":
                        switch (^(fi52)) {
                        case -9223372036854775808:
                            if false {
                                var v58 uint = 204
                                if (int8(v39) >= -1) {
                                    v53 = (v53 >> (((v58 + (v58 | 255))) & 63))
                                    v33 += v33
                                    var v59 int64 = -9223372036854775808
                                    _ = v59
                                }
                                v3 = v3
                                _ = v58
                            } else {
                                fmt.Println(v39, v1, v33)
                                v60 := GenS5{F0: (436 == 1)}
                                var v61 uint64 = 221
                                _ = v60
                                _ = v61
                            }
                        default:
                            v62 := []int{0, (^(fi52)), fi54, v1}
                            _ = v62
                        }
                    case "\n\t":
                        switch int(v39) {
                        case 213:
                            var v63 float64 = 1e10
                            _ = v63
                        }
                    }
                    _ = v56
                }
                v3 += (v37 >> ((69) & 63))
                _ = fi54
            }
            _ = fi52
            _ = v53
        }
    } else {
        v64 := GenS2{F0: 69, F1: (v3 ^ (v3 / (((v37 + 4294967295)) | 1))), F2: (v1 + v1)}
        _ = v64
    }
    switch 42 {
    case 128:
        v39 += v39
    case -757:
        v65 := []int{v1, v1, (v1 | v1)}
        _ = v65
    default:
        for fi66 := 0; fi66 < 7; fi66++ {
            fmt.Println(v33)
            var v67 int16 = -500
            _ = fi66
            _ = v67
        }
    }
    v37 -= (v3 | v3)
    v33 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v45
    _ = v46
}
