package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool; F1 int8 }
func genFn0() int {
    v6 := []int{-9223372036854775808}
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint16 = 0
            _ = fi8
            _ = v9
        }
        switch fi7 {
        case 42:
            if true {
                var v10 uint = 42
                _ = v10
            } else {
                fmt.Println(fi7)
                var v11 float64 = 0.0
                var v12 float32 = 1.5
                _ = v11
                _ = v12
            }
        case 591:
            var v13 uint8 = 128
            _ = v13
        case 383:
            v14 := GenS1{F0: 652, F1: 9223372036854775807}
            _ = v14
        }
        _ = fi7
    }
    v15 := GenS1{F0: 185, F1: 1}
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 uint16 = 294
        v18 := GenS1{F0: 42, F1: 1}
        _ = fi16
        _ = v17
        _ = v18
    }
    _ = v15
    _ = v6
    return 4294967295
}
func genFn1(p0 uint) int32 {
    v19 := GenS2{F0: -907, F1: 42}
    if false {
        switch 128 {
        case -2147483648:
            var v20 bool = (!false)
            _ = v20
        }
        var v21 uint = p0
        _ = v21
    } else {
        for fi22 := 0; fi22 < 7; fi22++ {
            v23 := []int{(-(fi22)), (fi22 / (((fi22 % ((fi22) | 1))) | 1))}
            v24 := GenS1{F0: -1, F1: -9223372036854775808}
            var v25 string = ("hello" + "a")
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        var v26 uint64 = 521
        _ = v26
    }
    var v27 float64 = float64(3.14159)
    v28 := GenS1{F0: 0, F1: 1}
    if true {
        fmt.Println(v27)
    } else {
        var v29 uint32 = 938
        var v30 int = 1
        _ = v29
        _ = v30
    }
    _ = p0
    _ = v19
    _ = v27
    _ = v28
    return 2147483647
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 string = "a"
    var v4 uint64 = 481
    v5 := make(map[uint32]uint32)
    v31 := GenS3{F0: false, F1: false}
    var v32 int16 = -1
    for fi33 := 0; fi33 < 12; fi33++ {
        fmt.Println(v3, fi33, fi33)
        _ = fi33
    }
    v34 := GenS1{F0: -304, F1: 64}
    v4 += 18446744073709551615
    var v35 int8 = 100
    if true {
        v35 = 0
        for fi36 := 0; fi36 < 20; fi36++ {
            v32 = (^(v32))
            if (false || true) {
                v37 := GenS3{F0: false, F1: true}
                _ = v37
            } else {
                var v38 uint64 = v4
                _ = v38
            }
            if false {
                switch v3 {
                case "a":
                    if false {
                        for fi39 := 0; fi39 < 6; fi39++ {
                            if true {
                                fmt.Println(fi39, v32, v3)
                                var v40 float32 = 0.1
                                v35 = v35
                                v32 += -1
                                switch v3 {
                                case "\n\t":
                                    var v41 int64 = 1
                                    _ = v41
                                case "hello":
                                    v42 := v2[genFn0()]
                                    _ = v42
                                case "":
                                    v43 := GenS3{F0: (!((3.14159 / (1.5 / -2.25)) < (0.1 / 3.14159))), F1: (!true)}
                                    fmt.Println(v4)
                                    _ = v43
                                default:
                                    _ = len(v2)
                                }
                                _ = v40
                            } else {
                                v44 := GenS3{F0: (-2.25 <= (0.0 * -2.25)), F1: (fi39 != (fi39 / (((fi36 << ((1) & 63))) | 1)))}
                                _ = v44
                            }
                            _ = fi39
                        }
                    } else {
                        v35 += (v35 << ((4294967295) & 63))
                        var v45 uint8 = 2
                        _ = v45
                    }
                case "":
                    v46 := []int{fi36, fi36}
                    _ = v46
                case "\n\t":
                    fmt.Println(v3, v35, v3)
                default:
                    v47 := GenS4{F0: genFn1(255)}
                    _ = v47
                }
                v48 := GenS5{F0: 2, F1: v35, F2: 0}
                _ = v48
            } else {
                var v49 float64 = float64(1.5)
                switch 2147483647 {
                case 125:
                    v50 := GenS4{F0: genFn1(42)}
                    _ = v50
                case -426:
                    for fi51 := 0; fi51 < 3; fi51++ {
                        v4 = 42
                        v52 := []int{}
                        _ = fi51
                        _ = v52
                    }
                default:
                    v53 := []int{32768, 1, (fi36 ^ (fi36 % (((fi36 % (((fi36 - 628)) | 1))) | 1))), (fi36 % (((fi36 >> ((2) & 63))) | 1))}
                    _ = v53
                }
                _ = v49
            }
            _ = fi36
        }
        v54 := []int{}
        _ = v54
    } else {
        v1[-2147483648] = 1
        if true {
            var v55 bool = (false && (!false))
            _ = v55
        } else {
            for fi56 := 0; fi56 < 8; fi56++ {
                fmt.Println(v3, v32, v4)
                _ = fi56
            }
            var v57 uint8 = 115
            switch "\n\t" {
            case "foo bar":
                v57 += 42
            }
            _ = v57
        }
        fmt.Println(v4)
    }
    v35 += (v35 & v35)
    if (127 < 255) {
        for fi58 := 0; fi58 < 19; fi58++ {
            for fi59 := 0; fi59 < 4; fi59++ {
                switch "foo bar" {
                case "hello":
                    v60 := GenS6{F0: ((v3 + v3) + v3)}
                    _ = v60
                case "a":
                    v4 = 31
                case "\n\t":
                    if (!(12 <= 42)) {
                        v35 = (v35 - v35)
                        switch "\n\t" {
                        case "hello":
                            v61 := GenS7{F0: (true || false), F1: 105}
                            _ = v61
                        default:
                            v35 = (v35 + (v35 % ((42) | 1)))
                        }
                    } else {
                        v62 := GenS4{F0: genFn1(1)}
                        fmt.Println(v32)
                        _ = v62
                    }
                default:
                    v35 = (v35 + int8(v4))
                    var v63 uint = 1
                    _ = v63
                }
                v64 := []int{}
                v35 = 42
                if len(v64) > 0 {
                    v64[0] = fi59
                }
                _ = fi59
                _ = v64
            }
            for fi65 := 0; fi65 < 11; fi65++ {
                fmt.Println(v32, fi58)
                _ = len(v2)
                delete(v5, 65535)
                var v66 uint16 = 42
                var v67 string = "\n\t"
                var v68 int = (fi65 << ((0) & 63))
                _ = fi65
                _ = v66
                _ = v67
                _ = v68
            }
            switch ((v3 + (v3 + "\n\t")) + v3) {
            case "a":
                v4 -= v4
            case "hello":
                var v69 int16 = 258
                _ = v69
            case "\n\t":
                v70 := GenS1{F0: -7, F1: 0}
                _ = v70
            }
            _ = fi58
        }
    } else {
        v32 -= v32
        var v71 uint16 = 941
        v71 = (v71 * 2)
        _ = v71
    }
    for fi72 := 0; fi72 < 2; fi72++ {
        var v73 uint = 329
        _ = fi72
        _ = v73
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
