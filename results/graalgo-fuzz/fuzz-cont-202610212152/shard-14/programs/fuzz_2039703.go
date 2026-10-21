package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 int; F1 bool; F2 int64 }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 int64; F1 uint16; F2 int }
type GenS8 struct { F0 int; F1 int; F2 int16 }
type GenS9 struct { F0 uint8; F1 int16; F2 bool }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint32; F1 uint16 }
type GenS12 struct { F0 int16; F1 int16 }

func main() {
    var v1 int64 = 719
    var v2 float64 = (((1.5 + -2.25) * 3.14159) - (float64(0.0) / 1.0))
    var v3 float64 = (-2.25 * 1e10)
    var v4 uint64 = 4294967295
    var v5 string = (("hello" + ("foo bar" + "a")) + "foo bar")
    v6 := GenS1{F0: v4}
    var v7 uint64 = v4
    v5 = ""
    for fi8 := 0; fi8 < 20; fi8++ {
        v3 = v2
        v9 := []int{fi8, (fi8 + 32767), fi8}
        var v10 int32 = 1
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int8 = 100
    v12 := GenS1{F0: (v4 / ((v4) | 1))}
    v13 := GenS1{F0: 42}
    v14 := GenS2{F0: 1}
    var v15 uint = 0
    switch v5 {
    case "a":
        switch 4294967295 {
        case -2147483648:
            v16 := []int{-921, 4294967295, 746, 32768}
            _ = v16
        case 969:
            v11 = v11
        case 32767:
            var v17 bool = (((1 <= 0) && ((v1 % ((0) | 1)) < (v1 * -9223372036854775808))) && (!(42 == 0)))
            _ = v17
        default:
            if ((v5 + (v5 + (v5 + "\n\t"))) != "\n\t") {
                switch 0 {
                case 0:
                    fmt.Println(v1, v7)
                case -314:
                    var v18 uint32 = 2147483648
                    _ = v18
                case -177:
                    v19 := GenS2{F0: 9223372036854775807}
                    _ = v19
                }
            }
        }
    case "\n\t":
        var v20 string = "\n\t"
        _ = v20
    case "foo bar":
        switch "foo bar" {
        case "hello":
            fmt.Println(v15)
        case "a":
            switch "" {
            case "foo bar":
                v5 = "\n\t"
            case "":
                v15 -= v15
            case "\n\t":
                fmt.Println(v3)
            }
        case "foo bar":
            v21 := GenS3{F0: "a", F1: 155}
            _ = v21
        }
    default:
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 bool = false
            _ = fi22
            _ = v23
        }
    }
    var v24 bool = false
    switch 42 {
    case 4294967295:
        for fi25 := 0; fi25 < 10; fi25++ {
            switch fi25 {
            case 42:
                v26 := GenS2{F0: (fi25 & fi25)}
                _ = v26
            case 32768:
                if (2147483648 == 2147483648) {
                    for fi27 := 0; fi27 < 5; fi27++ {
                        var v28 int64 = -1
                        _ = fi27
                        _ = v28
                    }
                    if (!v24) {
                        v29 := []int{-199, (fi25 << ((630) & 63)), 128, int(v11)}
                        if v24 {
                            fmt.Println(v2, v24)
                            for fi30 := 0; fi30 < 14; fi30++ {
                                switch fi25 {
                                case -2147483648:
                                    if v24 {
                                        v31 := GenS4{F0: 42, F1: uint8(v4)}
                                        _ = v31
                                    } else {
                                        v32 := GenS5{F0: (fi30 & (^(fi30))), F1: v24, F2: v1}
                                        _ = v32
                                    }
                                case -167:
                                    fmt.Println(v15, v11, fi25)
                                case 636:
                                    switch v5 {
                                    case "a":
                                        v15 = (v15 + v15)
                                    default:
                                        var v33 float64 = (3.14159 - ((v2 / v2) * 1e10))
                                        _ = v33
                                    }
                                default:
                                    v34 := GenS1{F0: v7}
                                    _ = v34
                                }
                                _ = fi30
                            }
                            v4 = (v7 % ((v4) | 1))
                        } else {
                            v2 += v2
                            v4 -= 18446744073709551615
                            v15 -= v15
                        }
                        _ = v29
                    }
                } else {
                    switch "hello" {
                    case "a":
                        var v35 uint16 = 0
                        _ = v35
                    }
                    var v36 string = ""
                    switch fi25 {
                    case 32768:
                        for fi37 := 0; fi37 < 6; fi37++ {
                            v38 := GenS6{F0: ("" + ("foo bar" + (("a" + v5) + v5))), F1: 181}
                            var v39 uint = v15
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    }
                    _ = v36
                }
            case 127:
                switch 2147483647 {
                case 127:
                    v40 := GenS7{F0: 9223372036854775807, F1: 1, F2: fi25}
                    _ = v40
                default:
                    v2 = v2
                }
            default:
                v41 := GenS2{F0: (fi25 & fi25)}
                _ = v41
            }
            v42 := []int{fi25, int(v4), (fi25 << ((v15) & 63))}
            var v43 uint16 = 1
            _ = fi25
            _ = v42
            _ = v43
        }
    case -354:
        v44 := GenS2{F0: 679}
        _ = v44
    default:
        v7 += (^(v7))
    }
    var v45 uint64 = v4
    for fi46 := 0; fi46 < 3; fi46++ {
        if (170 >= 127) {
            fmt.Println(v15)
        } else {
            fmt.Println(v15)
            fmt.Println(v11, v5, v7)
        }
        fmt.Println(v45)
        _ = fi46
    }
    for fi47 := 0; fi47 < 12; fi47++ {
        for fi48 := 0; fi48 < 1; fi48++ {
            var v49 int = (^(fi48))
            var v50 uint16 = 42
            if (2147483647 < 0) {
                var v51 float64 = float64(3.14159)
                v52 := []int{fi48}
                _ = v51
                _ = v52
            } else {
                v53 := GenS2{F0: 127}
                v54 := GenS8{F0: (fi47 / ((541) | 1)), F1: (fi48 - (fi48 - fi48)), F2: -32768}
                v55 := []int{(fi48 & fi47), (v49 % (((fi48 | 706)) | 1))}
                _ = v53
                _ = v54
                _ = v55
            }
            _ = fi48
            _ = v49
            _ = v50
        }
        _ = fi47
    }
    switch int(v15) {
    case 128:
        switch "" {
        case "\n\t":
            var v56 float32 = 0.1
            _ = v56
        case "hello":
            switch 198 {
            case -126:
                v57 := GenS9{F0: 255, F1: 1000, F2: false}
                _ = v57
            default:
                v58 := []int{-176, 923, -1, -2147483648}
                _ = v58
            }
        case "a":
            switch 895 {
            case 32767:
                if (150 >= 65535) {
                    if v24 {
                        var v59 int16 = 0
                        _ = v59
                    }
                    for fi60 := 0; fi60 < 7; fi60++ {
                        if false {
                            if (v1 <= v1) {
                                fmt.Println(v2, v4)
                                var v61 int32 = -1
                                _ = v61
                            } else {
                                var v62 int64 = 1
                                var v63 float64 = float64(((((0.1 * (float32(v3) * 3.14159)) - -2.25) + 0.0) - ((1.5 + 0.0) * 0.1)))
                                v45 -= (^(v4))
                                _ = v62
                                _ = v63
                            }
                        }
                        _ = fi60
                    }
                }
            case 42:
                v7 += v4
            }
        default:
            switch 127 {
            case -1:
                for fi64 := 0; fi64 < 2; fi64++ {
                    v4 -= 4294967295
                    for fi65 := 0; fi65 < 13; fi65++ {
                        v66 := GenS10{F0: v15}
                        if v24 {
                            var v67 float32 = (1.5 + (1e10 / 0.1))
                            v68 := []int{127}
                            v2 = 0.1
                            var v69 int16 = 1
                            if len(v68) > 0 {
                                v68[0] = (fi65 - fi65)
                            }
                            _ = v67
                            _ = v68
                            _ = v69
                        } else {
                            v70 := []int{fi65, fi64}
                            if len(v70) >= 2 {
                                v70 = v70[1:len(v70)]
                            }
                            if (v24 || (!(v1 <= int64(v45)))) {
                                v45 = v7
                            } else {
                                var v71 uint8 = 127
                                var v72 uint32 = 770
                                _ = v71
                                _ = v72
                            }
                            _ = v70
                        }
                        if (0 > uint32(v15)) {
                            if v24 {
                                v73 := []int{fi64, (fi64 % (((-(fi64))) | 1)), (fi65 | fi64)}
                                fmt.Println(v1, v11, fi65)
                                _ = v73
                            }
                            fmt.Println(v11, v11)
                            fmt.Println(fi65, v15, v3)
                        } else {
                            v74 := []int{744, 32767, (fi64 | 128)}
                            var v75 float64 = float64(0.0)
                            fmt.Println(v3)
                            switch 1 {
                            case 4294967295:
                                var v76 uint32 = 680
                                _ = v76
                            }
                            v7 = 505
                            if (v24 || ((!true) || ((fi64 & (fi65 + (fi64 | -2147483648))) != 0))) {
                                fmt.Println(v5)
                                fmt.Println(v4)
                            }
                            if (false || (!v24)) {
                                v1 = v1
                            } else {
                                if false {
                                    var v77 bool = v24
                                    _ = v77
                                } else {
                                    v78 := GenS8{F0: fi64, F1: (fi64 << (((v15 & (v15 - v15))) & 63)), F2: 744}
                                    fmt.Println(v5, v1)
                                    _ = v78
                                }
                            }
                            var v79 float32 = 0.1
                            v80 := GenS11{F0: uint32(v15), F1: 255}
                            var v81 uint64 = (v4 * (v7 - 300))
                            _ = v74
                            _ = v75
                            _ = v79
                            _ = v80
                            _ = v81
                        }
                        _ = fi65
                        _ = v66
                    }
                    v82 := GenS12{F0: 32767, F1: 304}
                    _ = fi64
                    _ = v82
                }
            default:
                var v83 int16 = 0
                _ = v83
            }
        }
    case 127:
        v4 = (v4 / ((v7) | 1))
    case 32768:
        v7 = v7
    }
    for fi84 := 0; fi84 < 3; fi84++ {
        var v85 int8 = v11
        _ = fi84
        _ = v85
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v7
}
