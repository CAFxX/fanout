package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint32; F1 int64; F2 string }
type GenS4 struct { F0 uint; F1 uint32; F2 int32 }
type GenS5 struct { F0 int; F1 uint; F2 int32 }
type GenS6 struct { F0 uint16; F1 int16; F2 uint }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]bool)
    var v3 int = -907
    v4 := make(map[uint32]string)
    v5 := make(map[int]uint32)
    switch (v3 / ((v3) | 1)) {
    case 32767:
        v6 := v2[uint32(v1)]
        _ = v6
    case -257:
        if ((v1 != 255) || (!true)) {
            for fi7 := 0; fi7 < 3; fi7++ {
                fmt.Println(fi7)
                v8 := GenS1{F0: 2147483648, F1: false}
                _ = fi7
                _ = v8
            }
            switch "\n\t" {
            case "hello":
                switch v3 {
                case 0:
                    var v9 int64 = -1
                    _ = v9
                case 963:
                    v10 := GenS2{F0: -1, F1: 42}
                    _ = v10
                case 32768:
                    v11 := GenS3{F0: 394, F1: 42, F2: "hello"}
                    _ = v11
                default:
                    for fi12 := 0; fi12 < 4; fi12++ {
                        v1 -= 0
                        v3 = fi12
                        _ = fi12
                    }
                }
            case "a":
                switch "hello" {
                case "\n\t":
                    v13 := []int{}
                    _ = v13
                case "foo bar":
                    v3 = (v3 & 127)
                case "a":
                    if (2147483647 != -2147483648) {
                        var v14 uint8 = 42
                        _ = v14
                    }
                default:
                    v15 := GenS1{F0: 2147483647, F1: (false || false)}
                    _ = v15
                }
            default:
                if true {
                    v16 := GenS4{F0: 42, F1: 65535, F2: -1}
                    var v17 uint = (v1 | (^(v1)))
                    _ = v16
                    _ = v17
                }
            }
        } else {
            var v18 int32 = 2147483647
            var v19 uint8 = 4
            _ = v18
            _ = v19
        }
    case 539:
        for fi20 := 0; fi20 < 2; fi20++ {
            for fi21 := 0; fi21 < 13; fi21++ {
                var v22 int8 = -76
                _ = fi21
                _ = v22
            }
            var v23 bool = ((1.5 + 0.0) == 1e10)
            if v23 {
                v23 = (!(741 > 215))
                if v23 {
                    fmt.Println(v3, v1)
                    var v24 int32 = 0
                    v25 := GenS5{F0: v3, F1: 2, F2: (v24 | 2147483647)}
                    _ = v24
                    _ = v25
                } else {
                    var v26 uint16 = 1
                    v27 := []int{fi20, (v3 - (v3 * (fi20 % ((v3) | 1))))}
                    _ = v26
                    _ = v27
                }
            } else {
                v1 = (v1 >> ((1) & 63))
            }
            _ = fi20
            _ = v23
        }
    default:
        v28 := []int{v3}
        _ = v28
    }
    v1 = (^(v1))
    if false {
        for fi29 := 0; fi29 < 4; fi29++ {
            if true {
                for fi30 := 0; fi30 < 17; fi30++ {
                    v3 = -48
                    _ = fi30
                }
                for fi31 := 0; fi31 < 10; fi31++ {
                    v5[875] = 2147483647
                    _ = fi31
                }
                v3 += v3
            } else {
                v32 := GenS1{F0: 65535, F1: (42 <= -732)}
                v1 = v1
                _ = v32
            }
            if false {
                fmt.Println(fi29)
            }
            _ = fi29
        }
        for fi33 := 0; fi33 < 14; fi33++ {
            v5[2147483647] = 731
            if true {
                for fi34 := 0; fi34 < 5; fi34++ {
                    switch fi33 {
                    case -2147483648:
                        var v35 float32 = float32(0.0)
                        _ = v35
                    case 128:
                        fmt.Println(fi33)
                    case 32767:
                        switch (fi34 * fi34) {
                        case -2147483648:
                            var v36 bool = ((((-2.25 / 1.0) * 0.1) * (1e10 / 1.0)) == (0.1 * 3.14159))
                            _ = v36
                        case -292:
                            v3 = fi34
                        case 9223372036854775807:
                            var v37 string = ""
                            _ = v37
                        default:
                            v38 := GenS4{F0: v1, F1: 2, F2: -1}
                            _ = v38
                        }
                    default:
                        v1 += (v1 - v1)
                    }
                    fmt.Println(fi34)
                    _ = fi34
                }
            }
            v3 = (v3 << ((4294967295) & 63))
            _ = fi33
        }
    } else {
        for fi39 := 0; fi39 < 18; fi39++ {
            var v40 uint = v1
            delete(v4, 0)
            v41 := GenS1{F0: uint32(v40), F1: (false || false)}
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    switch v3 {
    case 32767:
        if true {
            v42 := GenS6{F0: 255, F1: 32767, F2: (v1 - v1)}
            _ = v42
        } else {
            v3 = v3
            for fi43 := 0; fi43 < 3; fi43++ {
                v44 := GenS2{F0: -399, F1: -255}
                switch v3 {
                case 272:
                    if (!((!false) && (!true))) {
                        var v45 int16 = 1
                        _ = v45
                    } else {
                        for fi46 := 0; fi46 < 16; fi46++ {
                            if true {
                                v47 := GenS2{F0: -9223372036854775808, F1: -32768}
                                var v48 uint8 = 42
                                fmt.Println(v3)
                                _ = v47
                                _ = v48
                            } else {
                                v3 = fi43
                            }
                            switch "a" {
                            case "":
                                if (!(-397 == 1)) {
                                    var v49 int32 = -914
                                    v1 -= 1
                                    _ = v49
                                }
                            case "\n\t":
                                v3 -= -2147483648
                            default:
                                fmt.Println(fi43)
                            }
                            var v50 uint16 = 42
                            _ = fi46
                            _ = v50
                        }
                        var v51 float32 = 0.1
                        _ = v51
                    }
                default:
                    v1 = (v1 ^ v1)
                }
                _ = fi43
                _ = v44
            }
        }
    case 2147483647:
        v3 = 42
    case 32768:
        for fi52 := 0; fi52 < 5; fi52++ {
            var v53 uint64 = 158
            _ = fi52
            _ = v53
        }
    default:
        v1 = (v1 / (((v1 & (v1 << (((v1 & 0)) & 63)))) | 1))
    }
    v54 := GenS2{F0: -1, F1: 42}
    switch "foo bar" {
    case "\n\t":
        for fi55 := 0; fi55 < 2; fi55++ {
            for fi56 := 0; fi56 < 14; fi56++ {
                delete(v5, (fi55 / ((v3) | 1)))
                var v57 int64 = -9223372036854775808
                _ = fi56
                _ = v57
            }
            _ = fi55
        }
    case "":
        var v58 int8 = -128
        _ = v58
    case "a":
        var v59 int8 = -13
        _ = v59
    }
    var v60 uint64 = 341
    v1 = (v1 ^ (^(v1)))
    v61 := []int{v3, int(v60)}
    if (42 <= v60) {
        var v62 uint32 = 0
        v62 = (v62 - (^(v62)))
        v60 = uint64(v1)
        _ = v62
    }
    for fi63 := 0; fi63 < 13; fi63++ {
        if true {
            var v64 uint8 = 127
            var v65 float32 = (float32((-2.25 / 1.0)) / float32(0.1))
            var v66 int8 = -1
            _ = v64
            _ = v65
            _ = v66
        } else {
            switch "hello" {
            case "hello":
                v1 = (v1 / ((v1) | 1))
            case "a":
                fmt.Println(v60, v1, v1)
            }
            v1 = v1
        }
        var v67 uint8 = 1
        _ = fi63
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v61:", len(v61), v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v60
    _ = v61
}
