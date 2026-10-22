package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int32; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string; F1 uint16 }
func genFn0() uint {
    var v4 uint32 = 2147483647
    switch "foo bar" {
    case "a":
        for fi5 := 0; fi5 < 3; fi5++ {
            v4 = 528
            v4 -= 826
            _ = fi5
        }
    case "foo bar":
        if false {
            var v6 uint = 65535
            _ = v6
        }
    default:
        switch "\n\t" {
        case "":
            v4 = (v4 >> ((2) & 63))
        case "\n\t":
            for fi7 := 0; fi7 < 5; fi7++ {
                v8 := GenS1{F0: 255}
                v4 = v4
                v9 := GenS1{F0: 1}
                _ = fi7
                _ = v8
                _ = v9
            }
        case "a":
            if ((!(!true)) || false) {
                fmt.Println(v4)
                var v10 int8 = 68
                v11 := GenS2{F0: 657, F1: -251, F2: -1}
                _ = v10
                _ = v11
            }
        default:
            v4 = v4
        }
    }
    switch ("" + "\n\t") {
    case "foo bar":
        v4 = v4
    case "a":
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := []int{fi12, (fi12 ^ fi12), int(v4)}
            var v14 string = "hello"
            v4 = 0
            _ = fi12
            _ = v13
            _ = v14
        }
    case "\n\t":
        var v15 float64 = (0.1 + float64((float32(0.1) - 0.0)))
        _ = v15
    default:
        if (-1 >= 42) {
            if true {
                var v16 bool = false
                var v17 int = 214
                _ = v16
                _ = v17
            }
        }
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 int8 = 0
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 uint16 = 255
            var v22 float32 = -2.25
            var v23 int32 = -578
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        _ = fi18
        _ = v19
    }
    _ = v4
    return 2
}

func main() {
    v1 := []uint32{931, 110, 42, 0}
    var v2 uint16 = 842
    var v3 int64 = 1
    v2 = 32767
    if true {
        switch -9223372036854775808 {
        case 32768:
            if true {
                var v24 uint64 = 2
                v2 = (^(v2))
                v25 := GenS3{F0: (v2 >> ((uint(v24)) & 63)), F1: 1, F2: (^(v3))}
                _ = v24
                _ = v25
            } else {
                if false {
                    switch "foo bar" {
                    case "hello":
                        v3 = (^(v3))
                    }
                    switch 539 {
                    case -704:
                        if true {
                            v3 = v3
                            v3 = (-(v3))
                        } else {
                            v3 = (v3 % ((v3) | 1))
                            v2 = (v2 ^ v2)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    case -2147483648:
                        v26 := GenS1{F0: 2}
                        _ = v26
                    case 476:
                        fmt.Println(v2, v3)
                    default:
                        if false {
                            if false {
                                for fi27 := 0; fi27 < 2; fi27++ {
                                    _ = len(v1)
                                    fmt.Println(fi27, fi27, v2)
                                    _ = fi27
                                }
                                if true {
                                    for fi28 := 0; fi28 < 7; fi28++ {
                                        fmt.Println(v2, v2, fi28)
                                        switch 42 {
                                        case 128:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case 0:
                                            var v29 uint = genFn0()
                                            _ = v29
                                        case 810:
                                            switch (fi28 - (fi28 | -301)) {
                                            case -1:
                                                var v30 float64 = 1.5
                                                _ = v30
                                            case 9223372036854775807:
                                                var v31 uint16 = v2
                                                _ = v31
                                            case 2147483647:
                                                var v32 int16 = 42
                                                _ = v32
                                            default:
                                                switch "hello" {
                                                case "hello":
                                                    var v33 uint16 = v2
                                                    _ = v33
                                                case "a":
                                                    v1 = append(v1, 8)
                                                default:
                                                    if (1 < 65535) {
                                                        v2 += v2
                                                        var v34 uint16 = 1
                                                        var v35 uint16 = (v34 * 255)
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                }
                                            }
                                        default:
                                            switch "" {
                                            case "a":
                                                if false {
                                                    for fi36 := 0; fi36 < 18; fi36++ {
                                                        var v37 uint8 = 255
                                                        _ = fi36
                                                        _ = v37
                                                    }
                                                } else {
                                                    fmt.Println(v2, fi28)
                                                }
                                            case "hello":
                                                v3 = 18
                                            case "":
                                                var v38 uint8 = 138
                                                _ = v38
                                            default:
                                                for fi39 := 0; fi39 < 13; fi39++ {
                                                    _ = len(v1)
                                                    _ = fi39
                                                }
                                            }
                                        }
                                        v2 = 71
                                        _ = fi28
                                    }
                                } else {
                                    v3 = -736
                                }
                                if true {
                                    v3 += v3
                                } else {
                                    for fi40 := 0; fi40 < 2; fi40++ {
                                        v41 := GenS1{F0: 166}
                                        var v42 bool = (261 >= (fi40 & fi40))
                                        v1 = append(v1, 111)
                                        _ = fi40
                                        _ = v41
                                        _ = v42
                                    }
                                }
                            } else {
                                for fi43 := 0; fi43 < 4; fi43++ {
                                    for fi44 := 0; fi44 < 2; fi44++ {
                                        fmt.Println(fi44, fi44, fi44)
                                        v45 := GenS4{F0: 255, F1: 191}
                                        fmt.Println(fi44, v2)
                                        _ = fi44
                                        _ = v45
                                    }
                                    _ = fi43
                                }
                                var v46 int64 = v3
                                _ = v46
                            }
                            if (1 <= 2147483647) {
                                _ = len(v1)
                                var v47 float64 = (1e10 * 1e10)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v47
                            } else {
                                if (v3 != 9223372036854775807) {
                                    v2 += 65535
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = 198
                                    }
                                }
                                v2 -= (v2 * v2)
                                v3 += v3
                            }
                            if len(v1) > 0 {
                                v1[0] = 2147483648
                            }
                        } else {
                            v2 -= 1
                        }
                    }
                    v48 := GenS5{F0: 65535, F1: 594}
                    _ = v48
                }
                v3 = 1
            }
        case 127:
            var v49 int8 = 0
            _ = v49
        case 128:
            v2 += (v2 | v2)
        default:
            fmt.Println(v2, v3)
        }
        if (!false) {
            if len(v1) > 0 {
                v1[0] = 65535
            }
            var v50 int = -2147483648
            var v51 float32 = (0.1 - 1.5)
            _ = v50
            _ = v51
        } else {
            v2 = v2
            if false {
                v52 := GenS6{F0: 2, F1: 0}
                for fi53 := 0; fi53 < 3; fi53++ {
                    v3 = 0
                    v3 = (v3 | -9223372036854775808)
                    _ = fi53
                }
                _ = v52
            }
        }
    } else {
        v1 = append(v1, 1)
        v54 := GenS3{F0: 2, F1: 0, F2: v3}
        for fi55 := 0; fi55 < 18; fi55++ {
            v56 := GenS7{F0: (fi55 % (((^(fi55))) | 1))}
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
            var v57 uint32 = 0
            _ = fi55
            _ = v56
            _ = v57
        }
        _ = v54
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v58 uint = 131
    if false {
        var v59 uint = (v58 + 0)
        var v60 int = 128
        var v61 uint32 = 42
        _ = v59
        _ = v60
        _ = v61
    } else {
        var v62 bool = (((0.0 - 3.14159) < -2.25) && (!(22 > 127)))
        for fi63 := 0; fi63 < 4; fi63++ {
            var v64 string = "foo bar"
            switch fi63 {
            case -685:
                v3 = 9223372036854775807
            }
            _ = fi63
            _ = v64
        }
        _ = v62
    }
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch "hello" {
        case "foo bar":
            fmt.Println(v3, v3, v58)
        case "\n\t":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "a":
            for fi65 := 0; fi65 < 1; fi65++ {
                var v66 uint32 = uint32(v2)
                var v67 bool = true
                _ = fi65
                _ = v66
                _ = v67
            }
        default:
            if (v3 <= v3) {
                var v68 int16 = -806
                var v69 int32 = 42
                var v70 uint8 = 1
                _ = v68
                _ = v69
                _ = v70
            }
        }
    }
    for fi71 := 0; fi71 < 17; fi71++ {
        if (2147483647 > -2147483648) {
            switch (fi71 | (^(fi71))) {
            case 4294967295:
                v58 -= v58
            case 1:
                var v72 uint8 = 128
                _ = v72
            case -2147483648:
                fmt.Println(v2)
            default:
                fmt.Println(v3, v2, v58)
            }
        }
        switch "a" {
        case "a":
            _ = len(v1)
        case "":
            v2 += (v2 - uint16(v58))
        case "\n\t":
            for fi73 := 0; fi73 < 15; fi73++ {
                if len(v1) > 0 {
                    v1[0] = 415
                }
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    if (!true) {
                        v3 += -1
                        if true {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            if len(v1) > 0 {
                                v1[0] = 2
                            }
                            v74 := GenS7{F0: fi73}
                            v1 = append(v1, 775)
                            _ = v74
                        }
                        v1 = append(v1, 2147483648)
                    }
                case "":
                    v2 -= v2
                }
                v58 -= v58
                _ = fi73
            }
        }
        var v75 uint64 = 1
        _ = fi71
        _ = v75
    }
    v76 := GenS8{F0: "hello", F1: (v2 << (((v58 & genFn0())) & 63))}
    if false {
        for fi77 := 0; fi77 < 7; fi77++ {
            v2 = (v2 << ((v58) & 63))
            var v78 bool = (21 != 42)
            if (950 >= (v3 / ((v3) | 1))) {
                if (!(!(v3 == -510))) {
                    v2 = (v2 % (((v2 | 255)) | 1))
                }
            }
            _ = fi77
            _ = v78
        }
        var v79 uint = v58
        _ = v79
    } else {
        v58 = 255
    }
    v2 = (v2 & v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v2, v3, v58)
    default:
        v2 = v2
    }
    if len(v1) > 0 {
        v1[0] = 65535
    }
    var v80 uint8 = 88
    var v81 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v58:", v58)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v58
    _ = v76
    _ = v80
    _ = v81
}
