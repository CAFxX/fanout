package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint16; F2 int16 }
type GenS4 struct { F0 int; F1 uint8; F2 uint16 }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 string; F1 uint32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint32; F1 uint16; F2 int64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 2
    var v3 uint8 = 219
    var v4 string = "foo bar"
    if (!false) {
        var v5 int64 = -1
        _ = v5
    }
    var v6 uint32 = 0
    var v7 float32 = 0.1
    for fi8 := 0; fi8 < 1; fi8++ {
        switch "foo bar" {
        case "":
            switch (fi8 + (fi8 >> ((1) & 63))) {
            case -1:
                v9 := GenS1{F0: 128}
                _ = v9
            }
        case "hello":
            for fi10 := 0; fi10 < 14; fi10++ {
                if (((v3 * (v3 * (v3 / ((v3) | 1)))) != v3) || true) {
                    v11 := GenS2{F0: 42}
                    v2 -= v2
                    v7 = (((0.1 / (v7 + (((-2.25 * 1e10) / v7) / (1e10 * 1e10)))) - (((v7 - (float32(-2.25) + v7)) - v7) / (0.0 * (v7 + v7)))) - v7)
                    _ = v11
                }
                _ = fi10
            }
        case "foo bar":
            v12 := []int{(fi8 + fi8)}
            _ = v12
        default:
            var v13 bool = false
            _ = v13
        }
        v14 := []int{}
        _ = fi8
        _ = v14
    }
    _ = len(v1)
    v3 -= (v3 + 42)
    var v15 uint = 18446744073709551615
    for fi16 := 0; fi16 < 5; fi16++ {
        v2 = (^(v2))
        fmt.Println(v15, v3)
        _ = fi16
    }
    var v17 float32 = (v7 + 3.14159)
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(v7, v2)
        if (!true) {
            if (v6 != v6) {
                switch fi18 {
                case 107:
                    v6 += 825
                case 228:
                    var v19 string = v4
                    _ = v19
                case 1:
                    if (26 < (v15 + (v15 - (v15 ^ v15)))) {
                        v15 += 2
                        v3 -= 1
                    }
                }
                for fi20 := 0; fi20 < 1; fi20++ {
                    v21 := GenS3{F0: 0, F1: 744, F2: -703}
                    switch 32768 {
                    case 1:
                        var v22 int32 = 42
                        _ = v22
                    case 492:
                        v23 := GenS2{F0: uint(v6)}
                        _ = v23
                    }
                    _ = fi20
                    _ = v21
                }
            } else {
                v15 = v15
                switch (fi18 << (((v15 | uint(v2))) & 63)) {
                case 674:
                    var v24 uint64 = (v2 + v2)
                    _ = v24
                }
                var v25 int = fi18
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 12; fi26++ {
            if false {
                v27 := v1[v6]
                if (!true) {
                    _ = len(v1)
                } else {
                    var v28 float64 = 3.14159
                    v29 := []int{(fi18 - (^(fi26))), (fi26 & fi18), v27}
                    if ((true && true) && false) {
                        fmt.Println(v4)
                    } else {
                        switch fi26 {
                        case 32768:
                            v30 := GenS1{F0: (^(v3))}
                            _ = v30
                        case 127:
                            var v31 float32 = (v17 + 3.14159)
                            _ = v31
                        }
                        if false {
                            v6 -= v6
                            var v32 float64 = (v28 - ((v28 + -2.25) / 0.1))
                            if (253 <= 1) {
                                var v33 int16 = 1
                                var v34 int64 = 42
                                if (127 != 127) {
                                    if true {
                                        var v35 int8 = -9
                                        var v36 int32 = 42
                                        _ = v35
                                        _ = v36
                                    }
                                } else {
                                    if true {
                                        var v37 int8 = 1
                                        var v38 uint64 = v2
                                        _ = v37
                                        _ = v38
                                    }
                                }
                                if true {
                                    var v39 string = v4
                                    var v40 uint64 = (v2 + uint64(v3))
                                    if false {
                                        v41 := GenS4{F0: (fi26 >> ((v15) & 63)), F1: (v3 | (v3 << ((v15) & 63))), F2: 477}
                                        v2 += v2
                                        _ = v41
                                    } else {
                                        _ = len(v29)
                                        if false {
                                            fmt.Println(v2, v33)
                                        } else {
                                            v15 = (v15 + (v15 >> ((v15) & 63)))
                                            if ((v34 + (v34 >> (((^(v15))) & 63))) <= v34) {
                                                switch fi26 {
                                                case -9223372036854775808:
                                                    v40 = 2
                                                }
                                                var v42 uint = (^(v15))
                                                if true {
                                                    fmt.Println(v40, v33)
                                                } else {
                                                    var v43 uint64 = v40
                                                    v34 = v34
                                                    _ = v43
                                                }
                                                _ = v42
                                            } else {
                                                v29 = append(v29, fi26)
                                                var v44 int = v27
                                                if false {
                                                    v6 -= (v6 | v6)
                                                    v17 = v17
                                                } else {
                                                    v45 := GenS5{F0: 32767, F1: 0, F2: v3}
                                                    var v46 uint8 = v3
                                                    _ = v45
                                                    _ = v46
                                                }
                                                _ = v44
                                            }
                                        }
                                        if true {
                                            if len(v29) >= 2 {
                                                v29 = v29[1:len(v29)]
                                            }
                                        } else {
                                            v28 = v28
                                            if true {
                                                v47 := GenS6{F0: (v33 << ((v15) & 63)), F1: 42}
                                                var v48 uint16 = 635
                                                v27 = (fi26 / (((fi18 & (fi26 >> ((42) & 63)))) | 1))
                                                _ = v47
                                                _ = v48
                                            } else {
                                                switch fi18 {
                                                case 2147483647:
                                                    v15 = (v15 ^ uint(v6))
                                                case 842:
                                                    fmt.Println(v32)
                                                case 32767:
                                                    v7 = v7
                                                }
                                                v49 := GenS6{F0: -338, F1: (v3 & 159)}
                                                _ = v49
                                            }
                                            var v50 int16 = (v33 * v33)
                                            _ = v50
                                        }
                                        if true {
                                            v6 -= (v6 >> ((v15) & 63))
                                        }
                                    }
                                    var v51 float32 = 3.14159
                                    v33 = (v33 & v33)
                                    _ = v39
                                    _ = v40
                                    _ = v51
                                }
                                _ = v33
                                _ = v34
                            } else {
                                v1[(^(v6))] = (-(v27))
                                if ((float64(1.5) - v28) != v32) {
                                    if false {
                                        fmt.Println(v28, v28, v4)
                                        if true {
                                            var v52 int8 = 38
                                            var v53 uint32 = v6
                                            var v54 uint = (v15 << ((255) & 63))
                                            v29 = append(v29, fi26)
                                            if (1 >= -890) {
                                                v7 = ((1e10 / 1.5) - (v17 + 1e10))
                                            } else {
                                                var v55 int = -738
                                                var v56 uint16 = 65535
                                                _ = v55
                                                _ = v56
                                            }
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                        } else {
                                            v29 = append(v29, (fi18 << ((255) & 63)))
                                            var v57 float64 = (0.0 * 3.14159)
                                            _ = v57
                                        }
                                        if (int64(v3) == 1) {
                                            if len(v29) > 0 {
                                                v29[0] = v27
                                            }
                                        } else {
                                            var v58 int8 = 0
                                            _ = v58
                                        }
                                    }
                                    if (v2 >= 229) {
                                        v27 = (^(fi18))
                                        var v59 string = v4
                                        _ = v59
                                    } else {
                                        var v60 float64 = v32
                                        fmt.Println(v60, v15, v3)
                                        v32 = float64(1.5)
                                        var v61 string = ((v4 + v4) + (v4 + v4))
                                        _ = v60
                                        _ = v61
                                    }
                                    fmt.Println(v32, v2)
                                    v62 := GenS7{F0: 0}
                                    _ = v62
                                }
                            }
                            _ = v32
                        }
                    }
                    _ = v28
                    _ = v29
                }
                delete(v1, v6)
                _ = v27
            }
            if (v7 >= float32(((0.1 - 3.14159) + 1e10))) {
                var v63 uint8 = uint8(v6)
                if (false && true) {
                    v64 := v1[v6]
                    v65 := GenS4{F0: fi26, F1: (v63 - (^(v63))), F2: uint16(v15)}
                    _ = v64
                    _ = v65
                } else {
                    v66 := GenS8{F0: v4, F1: (v6 | 2147483647)}
                    _ = v66
                }
                v15 -= (v15 & v15)
                _ = v63
            }
            v67 := GenS6{F0: 6, F1: (v3 - 64)}
            if (!true) {
                v6 = uint32(v2)
                switch "a" {
                case "foo bar":
                    var v68 uint8 = v3
                    _ = v68
                case "a":
                    v2 += (v2 ^ uint64(v3))
                default:
                    var v69 bool = true
                    _ = v69
                }
            } else {
                switch ("" + v4) {
                case "a":
                    v70 := []int{fi26, fi18, fi26}
                    _ = v70
                case "hello":
                    switch (fi26 << ((855) & 63)) {
                    case -9223372036854775808:
                        var v71 float32 = (v7 / v7)
                        _ = v71
                    default:
                        var v72 uint32 = v6
                        _ = v72
                    }
                case "":
                    delete(v1, (v6 - (v6 ^ v6)))
                default:
                    v73 := GenS9{F0: 654}
                    _ = v73
                }
                var v74 int8 = 100
                _ = v74
            }
            v15 = v15
            _ = fi26
            _ = v67
        }
        _ = fi18
    }
    v75 := GenS10{F0: 647, F1: 2, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v75
}
