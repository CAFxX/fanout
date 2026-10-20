package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint }
type GenS6 struct { F0 int8 }

func main() {
    v1 := []string{("foo bar" + ("hello" + "a")), "hello", "a"}
    var v2 bool = (-32768 > -1)
    var v3 uint8 = 1
    var v4 int16 = -1
    v5 := make(map[uint32]uint32)
    var v6 int = 32767
    switch (^(v6)) {
    case 4294967295:
        var v7 int16 = v4
        _ = v7
    case 9223372036854775807:
        v8 := v5[65535]
        _ = v8
    default:
        var v9 uint16 = 65535
        _ = v9
    }
    _ = len(v5)
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 uint = 42
        switch 4294967295 {
        case -2147483648:
            v12 := GenS1{F0: -1, F1: 42}
            _ = v12
        case 68:
            if v2 {
                if true {
                    if (!(false || true)) {
                        for fi13 := 0; fi13 < 11; fi13++ {
                            fmt.Println(v4, fi10, v6)
                            _ = fi13
                        }
                        if (!v2) {
                            switch ("foo bar" + "foo bar") {
                            case "foo bar":
                                if (false || true) {
                                    switch "" {
                                    case "a":
                                        v2 = (v11 == (v11 | (v11 << ((v11) & 63))))
                                    case "hello":
                                        for fi14 := 0; fi14 < 16; fi14++ {
                                            var v15 uint16 = 1
                                            switch v6 {
                                            case -9223372036854775808:
                                                var v16 string = "hello"
                                                _ = v16
                                            default:
                                                switch (fi14 / (((-(fi14))) | 1)) {
                                                case 9223372036854775807:
                                                    var v17 int16 = v4
                                                    _ = v17
                                                case 32768:
                                                    v18 := v5[0]
                                                    _ = v18
                                                }
                                            }
                                            fmt.Println(v2, v11, v11)
                                            var v19 int32 = 1
                                            var v20 string = "a"
                                            _ = fi14
                                            _ = v15
                                            _ = v19
                                            _ = v20
                                        }
                                    case "foo bar":
                                        var v21 int16 = v4
                                        _ = v21
                                    default:
                                        for fi22 := 0; fi22 < 11; fi22++ {
                                            v23 := GenS1{F0: 3, F1: -182}
                                            v3 += v3
                                            if v2 {
                                                v24 := GenS2{F0: 18446744073709551615}
                                                _ = v24
                                            }
                                            _ = fi22
                                            _ = v23
                                        }
                                    }
                                } else {
                                    var v25 int = 42
                                    if ((-(v4)) < v4) {
                                        for fi26 := 0; fi26 < 10; fi26++ {
                                            v2 = (1.5 == float32((((3.14159 * (1.5 + 0.1)) + ((float64((0.1 / 1.5)) / 0.1) * 0.1)) / 1.0)))
                                            var v27 uint64 = 595
                                            _ = fi26
                                            _ = v27
                                        }
                                        switch "" {
                                        case "foo bar":
                                            v28 := GenS1{F0: 1, F1: 0}
                                            _ = v28
                                        }
                                        switch (v6 ^ (fi10 / (((v25 >> (((^(v11))) & 63))) | 1))) {
                                        case -135:
                                            v5[76] = 979
                                        default:
                                            v29 := GenS3{F0: v3, F1: (v11 / ((v11) | 1))}
                                            _ = v29
                                        }
                                    } else {
                                        v30 := GenS1{F0: int64(fi10), F1: -179}
                                        _ = v30
                                    }
                                    var v31 int64 = -1
                                    _ = v25
                                    _ = v31
                                }
                            }
                            var v32 uint = v11
                            _ = v32
                        }
                    } else {
                        if (0 == 32767) {
                            var v33 float32 = 1e10
                            if ((!true) && false) {
                                if v2 {
                                    fmt.Println(v11, v11, v3)
                                }
                                v5[2147483647] = 0
                            }
                            _ = v33
                        }
                    }
                    for fi34 := 0; fi34 < 10; fi34++ {
                        v35 := GenS4{F0: 2, F1: ((0 != 65535) && true), F2: v3}
                        _ = fi34
                        _ = v35
                    }
                    v6 = 2147483647
                }
                var v36 bool = (-74 != 1)
                var v37 uint = (v11 << ((v11) & 63))
                _ = v36
                _ = v37
            } else {
                if true {
                    for fi38 := 0; fi38 < 13; fi38++ {
                        if true {
                            v39 := GenS3{F0: 255, F1: (v11 + (v11 - (v11 - v11)))}
                            v4 += v4
                            delete(v5, 2)
                            _ = v39
                        }
                        _ = fi38
                    }
                } else {
                    v40 := GenS5{F0: 2147483647, F1: 39, F2: 0}
                    _ = v40
                }
                for fi41 := 0; fi41 < 12; fi41++ {
                    v1 = append(v1, "")
                    fmt.Println(v6)
                    _ = fi41
                }
                if false {
                    v4 -= v4
                    for fi42 := 0; fi42 < 15; fi42++ {
                        v43 := v5[0]
                        v44 := GenS2{F0: 18446744073709551615}
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                    var v45 uint32 = 724
                    _ = v45
                } else {
                    switch "a" {
                    case "hello":
                        if false {
                            if ((!false) && true) {
                                var v46 uint16 = 423
                                _ = v46
                            }
                        } else {
                            switch -9223372036854775808 {
                            case 4294967295:
                                for fi47 := 0; fi47 < 5; fi47++ {
                                    _ = len(v1)
                                    if (-9223372036854775808 > 42) {
                                        if (true || false) {
                                            v2 = false
                                            var v48 uint = v11
                                            _ = v48
                                        }
                                        var v49 bool = v2
                                        fmt.Println(v6, fi47)
                                        v2 = (v49 && true)
                                        fmt.Println(v11, v11, fi10)
                                        _ = v49
                                    } else {
                                        v3 -= v3
                                        fmt.Println(v4)
                                        var v50 uint8 = (v3 ^ (v3 - (v3 << ((v11) & 63))))
                                        var v51 uint32 = 562
                                        var v52 float32 = 1e10
                                        fmt.Println(v4)
                                        var v53 float32 = v52
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    }
                                    fmt.Println(v6, fi47, v3)
                                    _ = fi47
                                }
                            case 534:
                                fmt.Println(v3, fi10, fi10)
                            }
                            v54 := GenS4{F0: 255, F1: (false && v2), F2: v3}
                            _ = v54
                        }
                    default:
                        var v55 int64 = 834
                        _ = v55
                    }
                    switch 32767 {
                    case 790:
                        var v56 float64 = (0.0 / 1e10)
                        _ = v56
                    case -9223372036854775808:
                        for fi57 := 0; fi57 < 3; fi57++ {
                            v58 := GenS1{F0: 833, F1: -9223372036854775808}
                            v4 -= (v4 >> ((255) & 63))
                            _ = fi57
                            _ = v58
                        }
                    default:
                        fmt.Println(v4)
                    }
                }
            }
        case -1:
            var v59 bool = v2
            _ = v59
        default:
            if (0 != 316) {
                switch (fi10 >> ((0) & 63)) {
                case -147:
                    fmt.Println(v11)
                case -228:
                    var v60 int32 = 0
                    _ = v60
                }
                fmt.Println(v4, v3)
            } else {
                v61 := GenS6{F0: -28}
                v2 = ((v6 / ((fi10) | 1)) == (fi10 * fi10))
                _ = v61
            }
        }
        _ = fi10
        _ = v11
    }
    var v62 int64 = 42
    for fi63 := 0; fi63 < 19; fi63++ {
        _ = len(v1)
        _ = fi63
    }
    if false {
        if v2 {
            fmt.Println(v2, v2, v62)
            v64 := GenS4{F0: 247, F1: true, F2: (v3 + (v3 - (v3 >> ((720) & 63))))}
            for fi65 := 0; fi65 < 3; fi65++ {
                v62 = (-(v62))
                v6 = fi65
                v4 -= (^(v4))
                _ = fi65
            }
            _ = v64
        }
    }
    var v66 uint32 = 42
    switch v6 {
    case 2147483647:
        if true {
            v67 := GenS2{F0: 824}
            fmt.Println(v4)
            if v2 {
                switch 0 {
                case -1:
                    v6 = (v6 & (v6 ^ v6))
                default:
                    var v68 int32 = -2147483648
                    _ = v68
                }
                if v2 {
                    var v69 float64 = -2.25
                    _ = v69
                }
                v70 := GenS2{F0: uint64(v66)}
                _ = v70
            } else {
                switch "foo bar" {
                case "hello":
                    switch "hello" {
                    case "a":
                        var v71 uint64 = 2
                        _ = v71
                    case "foo bar":
                        if (v2 && v2) {
                            if (0 < 127) {
                                v62 -= (-(v62))
                                for fi72 := 0; fi72 < 10; fi72++ {
                                    for fi73 := 0; fi73 < 7; fi73++ {
                                        fmt.Println(v4, v4)
                                        var v74 string = "foo bar"
                                        v3 = v3
                                        _ = fi73
                                        _ = v74
                                    }
                                    v62 -= v62
                                    _ = fi72
                                }
                                fmt.Println(v4, v3, v4)
                            }
                            for fi75 := 0; fi75 < 4; fi75++ {
                                var v76 int32 = 42
                                var v77 float64 = float64(-2.25)
                                _ = fi75
                                _ = v76
                                _ = v77
                            }
                        } else {
                            v2 = v2
                            var v78 uint16 = 0
                            _ = v78
                        }
                    case "\n\t":
                        var v79 int16 = v4
                        _ = v79
                    }
                }
            }
            _ = v67
        } else {
            v62 += (v62 ^ v62)
        }
    case 4294967295:
        v6 = v6
    case 32767:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v6, v2)
    var v80 uint16 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v66:", v66)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v66
    _ = v80
}
