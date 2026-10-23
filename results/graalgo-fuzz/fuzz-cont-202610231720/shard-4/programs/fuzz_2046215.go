package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint64; F2 int }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 string; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64; F1 int16 }
type GenS10 struct { F0 int32; F1 int }
type GenS11 struct { F0 int8; F1 int16 }
type GenS12 struct { F0 int }
func genFn0() int64 {
    v6 := GenS1{F0: 432, F1: 571, F2: 1}
    switch ("foo bar" + "foo bar") {
    case "":
        if (-9223372036854775808 != 0) {
            v7 := []int{-2147483648, 2147483647, -843}
            switch "foo bar" {
            case "foo bar":
                if (false || true) {
                    v7 = append(v7, -9223372036854775808)
                    var v8 int16 = -362
                    _ = v8
                } else {
                    var v9 int64 = 604
                    _ = v9
                }
            case "":
                if true {
                    var v10 uint16 = 54
                    _ = v10
                }
            }
            _ = v7
        } else {
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 uint16 = 1
                fmt.Println(v12, fi11)
                for fi13 := 0; fi13 < 7; fi13++ {
                    v12 = (v12 * v12)
                    switch "a" {
                    case "a":
                        switch "" {
                        case "a":
                            var v14 uint16 = 255
                            _ = v14
                        case "":
                            v15 := GenS2{F0: -114, F1: 983, F2: 0}
                            _ = v15
                        }
                    case "hello":
                        if (189 <= 42) {
                            var v16 uint64 = 4294967295
                            _ = v16
                        }
                    }
                    v17 := []int{fi11, fi11}
                    _ = fi13
                    _ = v17
                }
                _ = fi11
                _ = v12
            }
            v18 := []int{}
            var v19 int64 = -712
            _ = v18
            _ = v19
        }
    case "\n\t":
        for fi20 := 0; fi20 < 1; fi20++ {
            v21 := GenS3{F0: true, F1: 0, F2: (fi20 - fi20)}
            switch (fi20 & (fi20 + (^(fi20)))) {
            case 591:
                var v22 float32 = (1.5 * 0.0)
                _ = v22
            case 32767:
                var v23 uint64 = 42
                _ = v23
            case 9223372036854775807:
                switch fi20 {
                case -1:
                    var v24 bool = false
                    _ = v24
                case 32768:
                    fmt.Println(fi20)
                }
            default:
                var v25 uint = 983
                _ = v25
            }
            switch fi20 {
            case 935:
                v26 := []int{fi20, (fi20 | (fi20 / ((fi20) | 1)))}
                _ = v26
            }
            _ = fi20
            _ = v21
        }
    case "a":
        switch "hello" {
        case "hello":
            if true {
                var v27 uint64 = 1
                for fi28 := 0; fi28 < 18; fi28++ {
                    v27 = v27
                    for fi29 := 0; fi29 < 12; fi29++ {
                        var v30 int = -687
                        _ = fi29
                        _ = v30
                    }
                    _ = fi28
                }
                _ = v27
            }
        case "a":
            var v31 float32 = 1e10
            _ = v31
        case "\n\t":
            v32 := GenS2{F0: 0, F1: 590, F2: 255}
            _ = v32
        }
    }
    var v33 uint8 = 69
    v34 := GenS4{F0: v33, F1: 0}
    switch "" {
    case "":
        v33 += v33
    case "\n\t":
        v35 := []int{2147483647}
        _ = v35
    case "foo bar":
        var v36 uint = 244
        _ = v36
    }
    _ = v33
    _ = v34
    _ = v6
    return 42
}
func genFn1(p0 bool) uint8 {
    if (false || false) {
        v37 := GenS3{F0: p0, F1: 4294967295, F2: 527}
        _ = v37
    } else {
        v38 := GenS5{F0: "hello", F1: 1000, F2: 2}
        _ = v38
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        v40 := []int{(^(fi39)), (fi39 & (^(fi39))), -2147483648, -2147483648}
        for fi41 := 0; fi41 < 9; fi41++ {
            if len(v40) >= 2 {
                v40 = v40[1:len(v40)]
            }
            p0 = (4294967295 == 486)
            if (genFn0() > 9223372036854775807) {
                var v42 uint16 = 924
                if p0 {
                    v43 := GenS2{F0: 42, F1: 0, F2: 127}
                    v44 := GenS6{F0: 67, F1: (v42 * (v42 * v42)), F2: 2}
                    v42 += (v42 ^ 210)
                    _ = v43
                    _ = v44
                }
                v45 := GenS7{F0: ((fi41 % ((fi41) | 1)) == (fi39 - 32768)), F1: v42}
                _ = v42
                _ = v45
            } else {
                p0 = (1 < -32768)
                if len(v40) >= 2 {
                    v40 = v40[1:len(v40)]
                }
            }
            _ = fi41
        }
        _ = fi39
        _ = v40
    }
    var v46 uint64 = 42
    _ = p0
    _ = v46
    return 208
}

func main() {
    var v1 uint64 = 2
    var v2 int16 = 0
    var v3 float64 = 3.14159
    var v4 float32 = -2.25
    v5 := make(map[string]int)
    v47 := GenS8{F0: 42}
    fmt.Println(v2, v2)
    v4 = float32(v3)
    if (-32768 > v2) {
        v48 := GenS4{F0: 0, F1: (v2 + v2)}
        fmt.Println(v4, v3, v3)
        _ = v48
    } else {
        switch 1 {
        case 9223372036854775807:
            for fi49 := 0; fi49 < 18; fi49++ {
                var v50 uint = 43
                v51 := GenS9{F0: v1, F1: v2}
                if (true || false) {
                    v3 = 0.0
                    var v52 int32 = -294
                    _ = v52
                } else {
                    v53 := GenS5{F0: "a", F1: (v2 | (v2 + (v2 - v2))), F2: 18446744073709551615}
                    v54 := GenS10{F0: -429, F1: 761}
                    _ = v53
                    _ = v54
                }
                _ = fi49
                _ = v50
                _ = v51
            }
        default:
            v2 -= 1000
        }
        if false {
            for fi55 := 0; fi55 < 12; fi55++ {
                v1 += v1
                _ = fi55
            }
        } else {
            v56 := []int{0, 644}
            var v57 int = 32768
            _ = v56
            _ = v57
        }
        for fi58 := 0; fi58 < 5; fi58++ {
            switch "\n\t" {
            case "hello":
                switch fi58 {
                case 9223372036854775807:
                    if (61 >= -28) {
                        for fi59 := 0; fi59 < 15; fi59++ {
                            v60 := GenS11{F0: 53, F1: int16(fi58)}
                            _ = fi59
                            _ = v60
                        }
                        var v61 int8 = -128
                        _ = v61
                    }
                }
            }
            _ = fi58
        }
    }
    var v62 uint8 = 42
    if true {
        switch 545 {
        case 32767:
            var v63 uint32 = 1
            _ = v63
        default:
            switch "a" {
            case "\n\t":
                var v64 uint16 = 77
                _ = v64
            case "foo bar":
                v65 := []int{9223372036854775807, 1, -189, -479}
                _ = v65
            case "":
                v2 = (v2 & (v2 << ((0) & 63)))
            default:
                if (100 <= -128) {
                    switch "\n\t" {
                    case "foo bar":
                        v5[("hello" + "a")] = 32767
                    case "a":
                        var v66 int32 = -1
                        _ = v66
                    case "hello":
                        if false {
                            for fi67 := 0; fi67 < 16; fi67++ {
                                var v68 uint8 = (v62 & 128)
                                if (((^(v1)) > (^(v1))) && true) {
                                    v1 -= v1
                                } else {
                                    v1 = v1
                                    var v69 int32 = 0
                                    _ = v69
                                }
                                var v70 string = ("" + (("\n\t" + "a") + "\n\t"))
                                _ = fi67
                                _ = v68
                                _ = v70
                            }
                        }
                    default:
                        var v71 uint32 = 515
                        _ = v71
                    }
                }
            }
        }
    } else {
        v72 := GenS12{F0: int(v62)}
        _ = v72
    }
    v73 := GenS2{F0: 948, F1: 485, F2: v62}
    for fi74 := 0; fi74 < 13; fi74++ {
        var v75 uint32 = 838
        _ = fi74
        _ = v75
    }
    v76 := []int{2147483647}
    v77 := GenS3{F0: false, F1: 2, F2: 1}
    fmt.Println(v4)
    v78 := v5["a"]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v62:", v62)
    fmt.Println("slicelen v76:", len(v76), v76)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v62
    _ = v73
    _ = v76
    _ = v77
    _ = v78
}
