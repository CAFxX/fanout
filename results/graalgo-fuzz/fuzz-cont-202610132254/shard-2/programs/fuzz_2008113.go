package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 uint64; F1 string }
func genFn0(p0 int32, p1 uint16, p2 int64) uint32 {
    var v6 int = 302
    var v7 int32 = (p0 - (p0 << ((725) & 63)))
    if (2 != 65535) {
        for fi8 := 0; fi8 < 9; fi8++ {
            v6 += v6
            _ = fi8
        }
        v9 := []int{v6, (v6 - v6)}
        var v10 uint8 = 42
        _ = v10
        _ = v9
    } else {
        v6 -= v6
        fmt.Println(p2, p1)
    }
    var v11 bool = (false || true)
    if (true || (!v11)) {
        if (!true) {
            fmt.Println(v11, v11)
            for fi12 := 0; fi12 < 5; fi12++ {
                switch "" {
                case "foo bar":
                    v13 := []int{}
                    _ = v13
                default:
                    switch fi12 {
                    case 0:
                        switch fi12 {
                        case 42:
                            p1 += p1
                        case 127:
                            v11 = (false && (!true))
                        }
                    case 32768:
                        if (128 < 0) {
                            var v14 int8 = 64
                            var v15 uint = 255
                            _ = v14
                            _ = v15
                        } else {
                            switch "" {
                            case "hello":
                                v16 := []int{fi12, (fi12 ^ v6), 2147483647, (v6 ^ v6)}
                                _ = v16
                            case "foo bar":
                                var v17 uint64 = 2
                                _ = v17
                            default:
                                for fi18 := 0; fi18 < 19; fi18++ {
                                    fmt.Println(v11, fi12, v7)
                                    if (v11 || v11) {
                                        v19 := []int{fi18, (fi12 << ((1) & 63)), (v6 * (fi18 / ((v6) | 1)))}
                                        _ = v19
                                    }
                                    v20 := GenS1{F0: 2}
                                    _ = fi18
                                    _ = v20
                                }
                            }
                        }
                    default:
                        v21 := []int{-870, (v6 - v6), -973}
                        _ = v21
                    }
                }
                v22 := GenS2{F0: 512}
                _ = fi12
                _ = v22
            }
            v7 += 1
        } else {
            switch v6 {
            case -9223372036854775808:
                switch "" {
                case "\n\t":
                    v6 = v6
                case "hello":
                    for fi23 := 0; fi23 < 14; fi23++ {
                        v24 := GenS2{F0: uint64(p1)}
                        for fi25 := 0; fi25 < 13; fi25++ {
                            switch "hello" {
                            case "\n\t":
                                p1 = p1
                            default:
                                var v26 uint32 = 2147483647
                                _ = v26
                            }
                            _ = fi25
                        }
                        _ = fi23
                        _ = v24
                    }
                case "a":
                    p2 -= p2
                default:
                    v27 := []int{-2147483648, 9223372036854775807, (v6 >> ((255) & 63)), (v6 + v6)}
                    _ = v27
                }
            default:
                p0 = v7
            }
            v28 := []int{(v6 / ((42) | 1)), v6, (v6 + (^(v6))), v6}
            p1 -= 0
            _ = v28
        }
    } else {
        var v29 float64 = 0.1
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    _ = v7
    return 2147483648
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 947
    var v3 uint = 457
    v4 := []uint32{}
    v5 := make(map[int]int)
    for fi30 := 0; fi30 < 11; fi30++ {
        v2 = v2
        fmt.Println(fi30)
        var v31 int = (fi30 + (-(fi30)))
        _ = fi30
        _ = v31
    }
    switch "foo bar" {
    case "a":
        for fi32 := 0; fi32 < 15; fi32++ {
            var v33 bool = (true || ((1e10 - 3.14159) < 3.14159))
            v3 += v3
            for fi34 := 0; fi34 < 10; fi34++ {
                var v35 int16 = 577
                v2 += (v2 - v2)
                if (65535 < v3) {
                    var v36 int32 = 686
                    var v37 uint16 = 951
                    if true {
                        fmt.Println(v2, v33, v33)
                        switch "a" {
                        case "\n\t":
                            var v38 uint32 = v2
                            _ = v38
                        case "hello":
                            var v39 string = "\n\t"
                            _ = v39
                        case "a":
                            v2 = v2
                        default:
                            if true {
                                v35 = v35
                                v37 = v37
                                v3 = v3
                            } else {
                                var v40 int8 = 0
                                fmt.Println(fi34, v3, v36)
                                var v41 bool = ((0.0 * 1.5) != 0.0)
                                _ = v40
                                _ = v41
                            }
                            var v42 uint32 = (v2 << ((v3) & 63))
                            _ = v42
                        }
                        v43 := GenS2{F0: 18446744073709551615}
                        v33 = (v33 || (!(!("foo bar" < "\n\t"))))
                        _ = v43
                    } else {
                        var v44 bool = (!(uint16(v3) <= (^(v37))))
                        var v45 int8 = 100
                        v46 := GenS3{F0: 0}
                        _ = v44
                        _ = v45
                        _ = v46
                    }
                    v37 += v37
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v36
                    _ = v37
                } else {
                    v47 := GenS1{F0: (v3 / ((v3) | 1))}
                    var v48 float64 = 1.5
                    v49 := GenS3{F0: 0}
                    _ = v47
                    _ = v48
                    _ = v49
                }
                _ = fi34
                _ = v35
            }
            _ = fi32
            _ = v33
        }
    default:
        var v50 int16 = -171
        _ = v50
    }
    switch "foo bar" {
    case "foo bar":
        v51 := GenS2{F0: 0}
        _ = v51
    case "a":
        var v52 int8 = -128
        _ = v52
    case "hello":
        var v53 int8 = 42
        _ = v53
    }
    fmt.Println(v2, v2)
    if false {
        if true {
            if false {
                var v54 string = "\n\t"
                for fi55 := 0; fi55 < 10; fi55++ {
                    var v56 float32 = 0.0
                    v4 = append(v4, (v2 << ((4294967295) & 63)))
                    _ = fi55
                    _ = v56
                }
                switch 32767 {
                case -484:
                    switch v54 {
                    case "hello":
                        var v57 uint16 = 3
                        _ = v57
                    case "foo bar":
                        for fi58 := 0; fi58 < 14; fi58++ {
                            v59 := GenS3{F0: 0}
                            var v60 uint16 = 2
                            for fi61 := 0; fi61 < 3; fi61++ {
                                var v62 uint16 = v60
                                v3 += 1
                                _ = fi61
                                _ = v62
                            }
                            _ = fi58
                            _ = v59
                            _ = v60
                        }
                    default:
                        for fi63 := 0; fi63 < 18; fi63++ {
                            var v64 bool = false
                            for fi65 := 0; fi65 < 8; fi65++ {
                                var v66 int32 = 1
                                _ = fi65
                                _ = v66
                            }
                            _ = fi63
                            _ = v64
                        }
                    }
                default:
                    switch v54 {
                    case "\n\t":
                        v67 := v5[32767]
                        _ = v67
                    }
                }
                _ = v54
            }
            v68 := GenS4{F0: 2, F1: 42}
            _ = v68
        }
        v3 -= v3
    }
    v3 += v3
    for fi69 := 0; fi69 < 4; fi69++ {
        var v70 float32 = -2.25
        _ = fi69
        _ = v70
    }
    v71 := GenS5{F0: 2, F1: ""}
    delete(v1, "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v71
}
