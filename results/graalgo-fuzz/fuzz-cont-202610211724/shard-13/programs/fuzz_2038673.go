package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 int64; F2 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 uint16; F2 int }
type GenS9 struct { F0 int8; F1 uint8; F2 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    var v3 uint64 = 57
    var v4 float32 = ((0.1 - 3.14159) * (-2.25 / 3.14159))
    var v5 uint64 = (^(v3))
    var v6 int = 127
    var v7 uint16 = 0
    switch v6 {
    case -2147483648:
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 uint = 4294967295
            v10 := GenS1{F0: 42}
            var v11 uint64 = (^(v5))
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    case 4294967295:
        for fi12 := 0; fi12 < 17; fi12++ {
            switch "\n\t" {
            case "foo bar":
                v6 -= int(v3)
            case "hello":
                for fi13 := 0; fi13 < 5; fi13++ {
                    if (!(!false)) {
                        v14 := GenS2{F0: 42, F1: v7, F2: (v5 ^ v5)}
                        v15 := []int{v6, (v6 >> ((255) & 63))}
                        v4 -= v4
                        _ = v14
                        _ = v15
                    }
                    v7 = v7
                    _ = fi13
                }
            case "\n\t":
                if false {
                    var v16 uint16 = (v7 * v7)
                    switch "hello" {
                    case "\n\t":
                        v17 := GenS3{F0: 2147483647}
                        _ = v17
                    default:
                        if (2 < 167) {
                            var v18 uint16 = (v7 - (^(v7)))
                            var v19 int8 = 1
                            _ = v18
                            _ = v19
                        } else {
                            var v20 uint32 = 65535
                            var v21 int8 = 100
                            var v22 int16 = 1
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    }
                    _ = v16
                }
            }
            _ = fi12
        }
    default:
        v23 := []int{4294967295, v6}
        _ = v23
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        v7 += 2
        _ = fi24
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        v4 = (v4 + 0.1)
        _ = fi25
    }
    v5 += v3
    v4 = (v4 / float32(((float64(1e10) / 0.1) * 1e10)))
    if false {
        for fi26 := 0; fi26 < 11; fi26++ {
            for fi27 := 0; fi27 < 18; fi27++ {
                if false {
                    if false {
                        var v28 int8 = 109
                        _ = v28
                    } else {
                        var v29 int = 0
                        switch "foo bar" {
                        case "":
                            if (true && (42 <= -206)) {
                                v3 -= v5
                                var v30 uint16 = (v7 * 2)
                                _ = v30
                            } else {
                                if (!("hello" == "foo bar")) {
                                    v5 -= v5
                                } else {
                                    var v31 int8 = -128
                                    _ = v31
                                }
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v5)
                                case "":
                                    var v32 int = (v29 + (v6 ^ (-(fi27))))
                                    _ = v32
                                }
                            }
                        case "hello":
                            var v33 int8 = -1
                            _ = v33
                        case "foo bar":
                            v34 := []int{}
                            _ = v34
                        default:
                            if false {
                                fmt.Println(v5, fi26, v6)
                                v6 -= -241
                            } else {
                                if (!false) {
                                    var v35 float64 = (0.1 * (3.14159 * ((0.0 - (1e10 / -2.25)) - 3.14159)))
                                    switch "hello" {
                                    case "":
                                        v36 := GenS2{F0: 594, F1: 32767, F2: (v3 + v5)}
                                        _ = v36
                                    default:
                                        if true {
                                            switch "" {
                                            case "\n\t":
                                                v6 -= (v6 >> ((1) & 63))
                                            case "":
                                                v3 -= 2
                                            case "hello":
                                                var v37 int32 = -2147483648
                                                _ = v37
                                            }
                                        }
                                    }
                                    _ = v35
                                }
                            }
                        }
                        _ = v29
                    }
                    var v38 uint16 = v7
                    v7 = v7
                    _ = v38
                } else {
                    fmt.Println(v6, v3, v4)
                    var v39 float32 = v4
                    var v40 float32 = v4
                    var v41 int32 = -539
                    v42 := v1[2]
                    v43 := GenS4{F0: 941}
                    _ = v39
                    _ = v40
                    _ = v41
                    _ = v42
                    _ = v43
                }
                _ = fi27
            }
            var v44 int8 = -128
            _ = fi26
            _ = v44
        }
    } else {
        for fi45 := 0; fi45 < 8; fi45++ {
            v46 := GenS1{F0: 999}
            v47 := []int{(fi45 - (v6 + (v6 * v6))), (-(fi45))}
            v48 := GenS2{F0: 2147483647, F1: (v7 / ((v7) | 1)), F2: (v3 ^ v5)}
            _ = fi45
            _ = v46
            _ = v47
            _ = v48
        }
        switch ("foo bar" + "a") {
        case "a":
            var v49 float32 = 3.14159
            _ = v49
        default:
            v50 := GenS4{F0: v7}
            _ = v50
        }
    }
    var v51 uint = 0
    v52 := GenS5{F0: (v51 + v51)}
    if true {
        var v53 uint64 = (v3 % ((v3) | 1))
        _ = v53
    }
    var v54 uint8 = 128
    v6 = (v6 << (((^(v51))) & 63))
    v55 := GenS6{F0: (751 >= 2147483647), F1: int64(v54), F2: 0}
    switch (v6 >> ((v51) & 63)) {
    case 127:
        switch "foo bar" {
        case "a":
            if false {
                switch "foo bar" {
                case "":
                    v7 += (v7 | 0)
                case "foo bar":
                    v3 -= v3
                case "\n\t":
                    v56 := []int{}
                    _ = v56
                default:
                    if true {
                        for fi57 := 0; fi57 < 16; fi57++ {
                            var v58 bool = (((((1e10 * ((0.1 - ((3.14159 - (float64(float32((0.0 / (3.14159 / 0.1)))) - (3.14159 / float64(v4)))) + (-2.25 / 1.5))) + (float64(float32(-2.25)) - -2.25))) != 1.5) && (700 >= 1)) && true) || true)
                            switch "" {
                            case "\n\t":
                                fmt.Println(v51)
                            case "foo bar":
                                var v59 float32 = v4
                                _ = v59
                            case "a":
                                var v60 int8 = -61
                                _ = v60
                            }
                            _ = fi57
                            _ = v58
                        }
                        switch "\n\t" {
                        case "a":
                            var v61 uint16 = 1
                            _ = v61
                        }
                        for fi62 := 0; fi62 < 15; fi62++ {
                            for fi63 := 0; fi63 < 1; fi63++ {
                                if (!true) {
                                    v64 := GenS7{F0: -1}
                                    _ = v64
                                } else {
                                    if true {
                                        v65 := GenS5{F0: (v51 + uint(v7))}
                                        _ = v65
                                    }
                                    if false {
                                        v51 -= (v51 & v51)
                                        v5 -= v3
                                        var v66 uint64 = v3
                                        _ = v66
                                    }
                                    if false {
                                        v4 = 0.0
                                        var v67 int64 = 9223372036854775807
                                        v5 -= (v3 & 1)
                                        _ = v67
                                    }
                                }
                                fmt.Println(fi62, v5)
                                v5 += v3
                                if false {
                                    v7 += uint16(v5)
                                    if (0.1 >= 0.1) {
                                        if (127 == 32) {
                                            v68 := GenS8{F0: 42, F1: (v7 / (((v7 + (v7 << (((v51 << (((v51 | v51)) & 63))) & 63)))) | 1)), F2: 9223372036854775807}
                                            v69 := []int{(^(v6)), v6, (^(fi62))}
                                            _ = v68
                                            _ = v69
                                        } else {
                                            v70 := GenS7{F0: int16(v3)}
                                            v54 += v54
                                            _ = v70
                                        }
                                        v71 := GenS8{F0: 0, F1: v7, F2: fi63}
                                        _ = v71
                                    } else {
                                        v5 -= (v3 & (^(v3)))
                                        v51 = (v51 * v51)
                                    }
                                } else {
                                    fmt.Println(v51)
                                }
                                _ = len(v1)
                                var v72 float32 = v4
                                _ = fi63
                                _ = v72
                            }
                            v6 += (^(v6))
                            var v73 bool = true
                            _ = fi62
                            _ = v73
                        }
                    } else {
                        var v74 int16 = 42
                        v75 := GenS9{F0: 0, F1: v54, F2: v74}
                        v76 := []int{(v6 % ((v6) | 1)), v6, v6}
                        _ = v74
                        _ = v75
                        _ = v76
                    }
                }
                var v77 int32 = 0
                v4 += 1e10
                _ = v77
            }
        case "\n\t":
            var v78 string = "hello"
            _ = v78
        case "hello":
            if false {
                v51 = v51
                var v79 int = 32767
                _ = v79
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v54
    _ = v55
    _ = v6
    _ = v7
}
