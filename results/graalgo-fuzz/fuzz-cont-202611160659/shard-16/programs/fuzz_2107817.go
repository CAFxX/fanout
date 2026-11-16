package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int8; F2 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int16; F2 int32 }
type GenS5 struct { F0 int64; F1 uint32; F2 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64; F1 int16 }
type GenS8 struct { F0 uint; F1 uint16 }
type GenS9 struct { F0 bool; F1 int8; F2 int64 }
type GenS10 struct { F0 int32; F1 string; F2 int8 }
type GenS11 struct { F0 bool; F1 uint; F2 uint32 }

func main() {
    var v1 int8 = -1
    v2 := []int{568, 127, 0}
    v3 := make(map[uint32]uint32)
    v4 := []string{"hello", "a", (("a" + ("\n\t" + "")) + ("\n\t" + ("a" + ("foo bar" + "foo bar")))), "\n\t"}
    v5 := []uint32{1, 2147483647, 2147483647, 2}
    var v6 uint8 = 127
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 int16 = -85
        v9 := v3[42]
        var v10 uint64 = 559
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 int64 = -91
        v13 := GenS1{F0: int32(v6)}
        if (false || (18446744073709551615 < 2)) {
            var v14 string = "hello"
            v12 = int64(v6)
            _ = v14
        } else {
            var v15 int64 = v12
            if false {
                var v16 float32 = 0.1
                if (("a" + "") == "\n\t") {
                    v17 := GenS1{F0: -976}
                    for fi18 := 0; fi18 < 6; fi18++ {
                        v19 := GenS2{F0: 0, F1: v1, F2: fi18}
                        _ = fi18
                        _ = v19
                    }
                    fmt.Println(v16, v1)
                    _ = v17
                } else {
                    v15 = v15
                    delete(v3, 107)
                    if false {
                        fmt.Println(v15, v15, v16)
                    }
                }
                for fi20 := 0; fi20 < 16; fi20++ {
                    v15 = v15
                    var v21 bool = true
                    var v22 string = "a"
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = v16
            } else {
                switch ("foo bar" + "hello") {
                case "":
                    if (0 != 42) {
                        v23 := GenS1{F0: -2147483648}
                        for fi24 := 0; fi24 < 10; fi24++ {
                            switch (fi24 & (-(fi11))) {
                            case -775:
                                v25 := GenS2{F0: 4294967295, F1: v1, F2: fi11}
                                _ = v25
                            case 1:
                                fmt.Println(v12)
                            case 32767:
                                var v26 float32 = 0.1
                                _ = v26
                            }
                            if (((!((("hello" + "") + "hello") <= ((("\n\t" + "foo bar") + "") + "a"))) || (true || false)) && (868 > 2)) {
                                v5 = append(v5, 548)
                                var v27 uint = 18446744073709551615
                                _ = v27
                            }
                            var v28 uint16 = 255
                            _ = fi24
                            _ = v28
                        }
                        var v29 uint = 45
                        _ = v23
                        _ = v29
                    }
                case "\n\t":
                    for fi30 := 0; fi30 < 2; fi30++ {
                        v31 := GenS3{F0: (923 != 255)}
                        _ = fi30
                        _ = v31
                    }
                case "hello":
                    v1 -= -1
                default:
                    var v32 int64 = 42
                    _ = v32
                }
                for fi33 := 0; fi33 < 7; fi33++ {
                    v34 := GenS3{F0: (!(!true))}
                    _ = fi33
                    _ = v34
                }
                switch ("foo bar" + "foo bar") {
                case "":
                    switch "a" {
                    case "":
                        var v35 int32 = int32(v15)
                        _ = v35
                    case "\n\t":
                        v12 = 613
                    case "hello":
                        v1 = v1
                    default:
                        switch ("\n\t" + "hello") {
                        case "":
                            v36 := GenS4{F0: (-(fi11)), F1: 1000, F2: 1}
                            _ = v36
                        case "a":
                            if true {
                                for fi37 := 0; fi37 < 16; fi37++ {
                                    var v38 int16 = -1
                                    v39 := GenS5{F0: -9223372036854775808, F1: uint32(v6), F2: 1}
                                    _ = fi37
                                    _ = v38
                                    _ = v39
                                }
                            }
                        case "hello":
                            var v40 uint16 = 65535
                            _ = v40
                        default:
                            _ = len(v3)
                        }
                    }
                case "hello":
                    v15 = (-(v15))
                }
            }
            _ = v15
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi41 := 0; fi41 < 6; fi41++ {
        switch "" {
        case "a":
            var v42 int32 = -1
            _ = v42
        case "hello":
            if false {
                var v43 int64 = -9223372036854775808
                var v44 uint32 = uint32(v6)
                v45 := GenS6{F0: 2}
                _ = v43
                _ = v44
                _ = v45
            }
        case "foo bar":
            v3[65535] = 2147483648
        default:
            if ((v1 * (-(v1))) >= v1) {
                var v46 bool = false
                if true {
                    v47 := GenS7{F0: 0, F1: -331}
                    v1 = (v1 & (v1 * (v1 % ((0) | 1))))
                    _ = v47
                } else {
                    v6 = (v6 / ((1) | 1))
                    if true {
                        var v48 int = (fi41 | (fi41 | fi41))
                        _ = v48
                    }
                    fmt.Println(v6)
                }
                _ = v46
            } else {
                switch (fi41 - (fi41 / ((fi41) | 1))) {
                case 198:
                    v2 = append(v2, fi41)
                case 419:
                    var v49 int = (fi41 ^ 4294967295)
                    _ = v49
                }
                fmt.Println(fi41, v1)
            }
        }
        _ = fi41
    }
    var v50 int32 = -631
    delete(v3, 934)
    if true {
        fmt.Println(v50)
        v1 -= -99
    } else {
        for fi51 := 0; fi51 < 1; fi51++ {
            v52 := GenS8{F0: 939, F1: 65535}
            for fi53 := 0; fi53 < 4; fi53++ {
                switch 999 {
                case 32768:
                    var v54 uint16 = 65535
                    _ = v54
                default:
                    switch (fi53 & fi53) {
                    case 2147483647:
                        v6 -= v6
                    }
                }
                v6 = (v6 & (v6 ^ v6))
                v3[uint32(v6)] = 1
                v50 += (v50 + (v50 >> ((65535) & 63)))
                _ = fi53
            }
            v55 := GenS9{F0: (0.0 <= (1e10 - (3.14159 * 0.1))), F1: v1, F2: -1}
            _ = fi51
            _ = v52
            _ = v55
        }
    }
    v50 = (v50 * (v50 * v50))
    for fi56 := 0; fi56 < 18; fi56++ {
        var v57 bool = (!false)
        v58 := GenS10{F0: (v50 & -462), F1: "\n\t", F2: v1}
        _ = fi56
        _ = v57
        _ = v58
    }
    v59 := GenS8{F0: 924, F1: 152}
    for fi60 := 0; fi60 < 11; fi60++ {
        for fi61 := 0; fi61 < 17; fi61++ {
            var v62 float64 = 3.14159
            _ = fi61
            _ = v62
        }
        _ = fi60
    }
    var v63 int8 = (-(v1))
    var v64 int8 = -73
    for fi65 := 0; fi65 < 11; fi65++ {
        var v66 uint32 = 4294967295
        _ = fi65
        _ = v66
    }
    for fi67 := 0; fi67 < 2; fi67++ {
        if ((3.14159 > 1e10) && ((false && (!(9223372036854775807 < -133))) && (1e10 < (1.5 - (3.14159 / 0.1))))) {
            var v68 float64 = 3.14159
            var v69 float64 = v68
            for fi70 := 0; fi70 < 5; fi70++ {
                _ = len(v5)
                if len(v5) > 0 {
                    v5[0] = 42
                }
                _ = fi70
            }
            _ = v68
            _ = v69
        } else {
            for fi71 := 0; fi71 < 10; fi71++ {
                if (!(true || (32767 == -1))) {
                    _ = len(v3)
                    v72 := GenS11{F0: (!((!((^(v6)) >= v6)) && false)), F1: 1, F2: 2147483648}
                    _ = v72
                } else {
                    v50 -= (v50 & v50)
                    if (42 == 65535) {
                        _ = len(v3)
                    }
                    fmt.Println(v50, fi71)
                    if len(v4) > 0 {
                        v4[0] = "hello"
                    }
                }
                fmt.Println(v6, fi67)
                _ = fi71
            }
        }
        var v73 uint = 255
        _ = fi67
        _ = v73
    }
    var v74 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v59
    _ = v6
    _ = v63
    _ = v64
    _ = v74
}
