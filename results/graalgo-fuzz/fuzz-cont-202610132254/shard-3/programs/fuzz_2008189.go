package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 bool; F1 int; F2 uint }
type GenS5 struct { F0 int; F1 uint32; F2 uint }
type GenS6 struct { F0 int32; F1 uint; F2 int8 }
type GenS7 struct { F0 int32; F1 int8; F2 int16 }
func genFn0(p0 uint) string {
    for fi5 := 0; fi5 < 8; fi5++ {
        fmt.Println(p0, p0)
        _ = fi5
    }
    p0 = 42
    if false {
        v6 := []int{128, -390}
        switch "\n\t" {
        case "":
            switch 128 {
            case -112:
                var v7 int32 = -1
                _ = v7
            }
        case "foo bar":
            fmt.Println(p0)
        case "\n\t":
            if (65535 <= 817) {
                v8 := GenS1{F0: 834}
                if false {
                    _ = len(v6)
                } else {
                    v9 := GenS1{F0: 2147483648}
                    for fi10 := 0; fi10 < 16; fi10++ {
                        p0 = 553
                        v6 = append(v6, (fi10 | fi10))
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi10
                    }
                    _ = v9
                }
                var v11 int64 = 0
                _ = v11
                _ = v8
            } else {
                for fi12 := 0; fi12 < 15; fi12++ {
                    var v13 int16 = 1000
                    p0 = (^(p0))
                    var v14 int8 = 0
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                v15 := GenS1{F0: uint32(p0)}
                p0 = p0
                _ = v15
            }
        default:
            var v16 int16 = -32768
            _ = v16
        }
        _ = v6
    }
    var v17 uint = p0
    var v18 uint64 = 18446744073709551615
    _ = p0
    _ = v17
    _ = v18
    return "\n\t"
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    var v4 float64 = -2.25
    var v19 int32 = -598
    var v20 uint = 42
    switch "a" {
    case "hello":
        v4 = float64(((1e10 / (0.1 * 0.1)) + (float32((3.14159 * 1e10)) + -2.25)))
    case "a":
        var v21 uint32 = 2
        _ = v21
    case "\n\t":
        var v22 float32 = 1.5
        _ = v22
    default:
        v23 := GenS1{F0: 2}
        _ = v23
    }
    switch "foo bar" {
    case "a":
        var v24 uint64 = (v1 * 1)
        _ = v24
    case "":
        v25 := GenS2{F0: 1, F1: v19}
        _ = v25
    case "hello":
        switch 42 {
        case -497:
            v19 = 1
        case 727:
            v2["\n\t"] = "foo bar"
        case 898:
            for fi26 := 0; fi26 < 16; fi26++ {
                for fi27 := 0; fi27 < 13; fi27++ {
                    v28 := []int{(fi27 / (((fi27 - (-(fi26)))) | 1))}
                    v19 -= v19
                    fmt.Println(v19)
                    _ = fi27
                    _ = v28
                }
                var v29 int16 = 1
                _ = fi26
                _ = v29
            }
        default:
            v1 += v1
        }
    default:
        fmt.Println(v1, v1)
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        v1 += v1
        _ = fi30
    }
    switch 2147483647 {
    case 42:
        for fi31 := 0; fi31 < 19; fi31++ {
            var v32 int = fi31
            _ = fi31
            _ = v32
        }
    case 401:
        if (!false) {
            var v33 uint8 = 255
            v34 := []int{32768, 42, 127, 9223372036854775807}
            var v35 uint16 = 65535
            _ = v33
            _ = v34
            _ = v35
        } else {
            v36 := GenS2{F0: -1, F1: v19}
            v37 := GenS1{F0: 202}
            for fi38 := 0; fi38 < 15; fi38++ {
                switch fi38 {
                case -563:
                    var v39 bool = false
                    _ = v39
                case 9223372036854775807:
                    for fi40 := 0; fi40 < 4; fi40++ {
                        if ((false || true) || (701 != (v19 << ((v20) & 63)))) {
                            if ((false || (42 <= 1000)) || (v4 == v4)) {
                                var v41 uint = (v20 ^ (v20 / ((v20) | 1)))
                                _ = v41
                            } else {
                                var v42 float32 = (3.14159 / float32(v4))
                                _ = len(v2)
                                _ = v42
                            }
                            var v43 uint16 = 48
                            var v44 int16 = -1
                            if (-32768 <= (v44 & 0)) {
                                v45 := GenS3{F0: v1, F1: 9223372036854775807}
                                _ = v45
                            }
                            var v46 float64 = ((3.14159 - 0.1) + (v4 / (3.14159 * v4)))
                            _ = v43
                            _ = v44
                            _ = v46
                        }
                        _ = fi40
                    }
                case 127:
                    if (1.5 <= (0.1 / (float32(v4) - (3.14159 / float32(((3.14159 / v4) + (float64((-2.25 / -2.25)) - (v4 * ((v4 + float64(-2.25)) + v4))))))))) {
                        v47 := []int{(fi38 % ((9223372036854775807) | 1)), 128, 32768}
                        var v48 int = (fi38 ^ fi38)
                        _ = v47
                        _ = v48
                    }
                }
                var v49 bool = true
                v50 := v2[("" + ("a" + "hello"))]
                _ = fi38
                _ = v49
                _ = v50
            }
            _ = v36
            _ = v37
        }
    default:
        fmt.Println(v20, v19)
    }
    if ((2 > 65535) && ("hello" == "")) {
        fmt.Println(v19, v4)
    } else {
        v51 := GenS4{F0: true, F1: -363, F2: v20}
        _ = len(v2)
        switch 32767 {
        case 2147483647:
            v52 := []int{-9223372036854775808, 1}
            _ = v52
        default:
            if (!false) {
                fmt.Println(v1, v1, v19)
                v53 := GenS1{F0: 1}
                var v54 uint32 = 42
                _ = v53
                _ = v54
            } else {
                var v55 int64 = -1
                var v56 int8 = 1
                _ = v55
                _ = v56
            }
        }
        _ = v51
    }
    v2["\n\t"] = "\n\t"
    if (-9223372036854775808 < 0) {
        v57 := GenS1{F0: 42}
        v58 := GenS5{F0: 74, F1: 2147483647, F2: v20}
        _ = v57
        _ = v58
    } else {
        v4 = float64((0.0 * 0.0))
    }
    if (331 >= 406) {
        for fi59 := 0; fi59 < 9; fi59++ {
            v60 := GenS1{F0: 2147483647}
            for fi61 := 0; fi61 < 17; fi61++ {
                v62 := GenS6{F0: (v19 - (-(v19))), F1: 811, F2: 0}
                _ = fi61
                _ = v62
            }
            _ = fi59
            _ = v60
        }
        if (v19 != (^(v19))) {
            if true {
                var v63 int8 = -96
                v4 = 0.1
                _ = v63
            } else {
                for fi64 := 0; fi64 < 16; fi64++ {
                    var v65 uint64 = (v1 + (^(v1)))
                    var v66 float32 = (0.0 * (float32(v4) + 0.0))
                    _ = fi64
                    _ = v65
                    _ = v66
                }
            }
            _ = len(v3)
            for fi67 := 0; fi67 < 5; fi67++ {
                v68 := []int{688, (fi67 | (fi67 ^ (fi67 + 128))), 42, fi67}
                _ = fi67
                _ = v68
            }
        } else {
            if (!(748 == 2147483648)) {
                var v69 int8 = 0
                _ = v69
            }
            v1 -= v1
        }
        switch 2147483647 {
        case 0:
            v20 -= v20
        }
    }
    fmt.Println(v1, v19)
    v70 := GenS6{F0: (v19 | (v19 ^ (v19 / ((v19) | 1)))), F1: v20, F2: -1}
    v19 += (v19 * (v19 / ((v19) | 1)))
    var v71 uint = 255
    for fi72 := 0; fi72 < 12; fi72++ {
        for fi73 := 0; fi73 < 16; fi73++ {
            var v74 int16 = 1
            var v75 int32 = v19
            v76 := []int{(fi73 >> (((v71 * (v71 ^ v20))) & 63)), (fi73 >> (((v20 / ((v71) | 1))) & 63)), fi72}
            switch fi73 {
            case -618:
                var v77 int64 = 9223372036854775807
                _ = v77
            }
            fmt.Println(v74)
            _ = fi73
            _ = v74
            _ = v75
            _ = v76
        }
        var v78 float32 = 0.0
        _ = fi72
        _ = v78
    }
    v79 := GenS7{F0: (^(v19)), F1: -128, F2: 823}
    switch -1 {
    case 9223372036854775807:
        v80 := GenS6{F0: (v19 ^ (v19 * v19)), F1: v71, F2: 127}
        _ = v80
    case 2147483647:
        switch 127 {
        case 4294967295:
            v19 = 250
        case 0:
            v81 := v2["hello"]
            _ = v81
        case 32767:
            for fi82 := 0; fi82 < 7; fi82++ {
                if false {
                    if (0 >= -1) {
                        v83 := GenS2{F0: 221, F1: v19}
                        _ = v83
                    }
                }
                fmt.Println(v4)
                _ = fi82
            }
        default:
            fmt.Println(v71, v4)
        }
    case 400:
        var v84 uint8 = 255
        _ = v84
    }
    v1 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v70
    _ = v71
    _ = v79
}
