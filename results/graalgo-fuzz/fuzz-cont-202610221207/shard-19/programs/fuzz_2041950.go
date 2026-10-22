package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 string; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 uint8; F1 int32; F2 bool }
type GenS8 struct { F0 int; F1 uint16 }
type GenS9 struct { F0 uint16; F1 int16 }
type GenS10 struct { F0 uint32; F1 int8; F2 uint }
func genFn0(p0 int32, p1 int64) int64 {
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := []int{}
        if true {
            p1 = 377
        } else {
            for fi7 := 0; fi7 < 14; fi7++ {
                fmt.Println(fi7, fi5)
                var v8 uint32 = 42
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 15; fi9++ {
                fmt.Println(fi5, fi5)
                _ = fi9
            }
            if (162 < 32767) {
                switch 0 {
                case 803:
                    fmt.Println(p1, p1, p0)
                }
            }
        }
        _ = fi5
        _ = v6
    }
    fmt.Println(p0, p1)
    p1 = 42
    _ = p0
    _ = p1
    return p1
}
func genFn1(p0 uint8, p1 int) int32 {
    p0 = p0
    var v10 uint64 = 18446744073709551615
    v10 = v10
    _ = p0
    _ = p1
    _ = v10
    return 42
}

func main() {
    var v1 float64 = 1e10
    v2 := []int{0, -2147483648, 128}
    v3 := []int{-2147483648, 2147483647, 32768}
    v4 := make(map[uint32]bool)
    var v11 uint32 = 2147483647
    if len(v3) > 0 {
        v3[0] = -698
    }
    switch ("a" + ("hello" + "")) {
    case "a":
        v11 = v11
    case "":
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 uint16 = 693
            _ = fi12
            _ = v13
        }
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        fmt.Println(v11)
    }
    v3 = append(v3, 1)
    fmt.Println(v11, v1)
    var v14 bool = true
    v2 = append(v2, 127)
    switch 127 {
    case 9223372036854775807:
        if v14 {
            fmt.Println(v14, v11)
            for fi15 := 0; fi15 < 13; fi15++ {
                v16 := GenS1{F0: uint(v11), F1: 1}
                v4[0] = false
                var v17 uint16 = 32767
                _ = fi15
                _ = v16
                _ = v17
            }
        } else {
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 int32 = -1
                if (!v14) {
                    switch fi18 {
                    case -2147483648:
                        v20 := v4[v11]
                        _ = v20
                    default:
                        _ = len(v3)
                    }
                    v19 += v19
                } else {
                    v14 = (42 < 42)
                    if v14 {
                        v4[(^(v11))] = (1 >= -624)
                        v11 = (v11 - v11)
                    }
                }
                if (!true) {
                    for fi21 := 0; fi21 < 19; fi21++ {
                        v11 = v11
                        var v22 uint64 = 1
                        if false {
                            var v23 uint = 116
                            var v24 int32 = v19
                            _ = v23
                            _ = v24
                        } else {
                            if ((fi21 >> ((18446744073709551615) & 63)) <= 42) {
                                var v25 string = "hello"
                                v25 = v25
                                fmt.Println(v19)
                                var v26 int64 = -428
                                _ = v25
                                _ = v26
                            }
                            var v27 float32 = ((((0.0 * 1e10) - -2.25) / 1.0) / (0.1 + 1e10))
                            if (v14 && v14) {
                                v2 = append(v2, fi18)
                                v28 := GenS2{F0: "a", F1: 717, F2: 1}
                                _ = v28
                            } else {
                                switch "foo bar" {
                                case "hello":
                                    _ = len(v4)
                                case "a":
                                    v27 = ((((float32(float64(v27)) + -2.25) * float32(0.1)) * 0.1) / ((v27 + float32(v1)) / ((v27 + (3.14159 + 1e10)) * (float32(float64(((v27 * v27) * v27))) + 3.14159))))
                                case "foo bar":
                                    v14 = ((^(v22)) == (v22 & 42))
                                default:
                                    fmt.Println(v14, v1)
                                }
                                if v14 {
                                    v29 := GenS3{F0: 1, F1: (v11 | 65535)}
                                    _ = v29
                                } else {
                                    if v14 {
                                        if (v27 < float32(float64(v27))) {
                                            var v30 float32 = v27
                                            v1 += 1.5
                                            var v31 int8 = -59
                                            _ = v30
                                            _ = v31
                                        } else {
                                            var v32 int64 = genFn0((v19 ^ genFn1(42, fi21)), -9223372036854775808)
                                            var v33 int16 = -938
                                            var v34 int8 = int8(fi18)
                                            if v14 {
                                                if (false && ((v19 >> ((uint(v11)) & 63)) <= (v19 % (((v19 + (^(v19)))) | 1)))) {
                                                    v22 -= v22
                                                    if v14 {
                                                        v34 = (v34 & 94)
                                                        v33 = (v33 * v33)
                                                    } else {
                                                        delete(v4, uint32(v22))
                                                    }
                                                    switch fi18 {
                                                    case -2147483648:
                                                        v35 := GenS2{F0: "foo bar", F1: v33, F2: 127}
                                                        v36 := GenS3{F0: 65535, F1: v11}
                                                        var v37 int8 = (^(v34))
                                                        _ = v35
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                    v33 -= (v33 - v33)
                                                    v38 := GenS4{F0: uint8(v22), F1: (-1 < (v34 ^ (-(v34))))}
                                                    _ = v38
                                                }
                                                switch fi18 {
                                                case -105:
                                                    var v39 uint8 = 42
                                                    _ = v39
                                                default:
                                                    var v40 uint = 65535
                                                    _ = v40
                                                }
                                                if v14 {
                                                    switch (fi21 - (fi21 & (fi21 >> ((2) & 63)))) {
                                                    case 177:
                                                        v41 := GenS5{F0: (v33 & 568), F1: 1}
                                                        _ = v41
                                                    }
                                                } else {
                                                    v42 := GenS6{F0: int8(v22), F1: 1}
                                                    v11 = v11
                                                    _ = v42
                                                }
                                                v43 := GenS1{F0: 42, F1: 0}
                                                v44 := GenS7{F0: 127, F1: (v19 ^ v19), F2: (!(-368 > 829))}
                                                _ = v43
                                                _ = v44
                                            } else {
                                                _ = len(v2)
                                                v45 := GenS4{F0: 248, F1: true}
                                                _ = v45
                                            }
                                            _ = v32
                                            _ = v33
                                            _ = v34
                                        }
                                        if (fi21 == fi21) {
                                            v27 = ((v27 / v27) / 1.5)
                                            if true {
                                                v27 += -2.25
                                                fmt.Println(fi21, fi21, v1)
                                                if v14 {
                                                    fmt.Println(fi18, fi18)
                                                    v19 -= (-(v19))
                                                    v22 = 4294967295
                                                }
                                            } else {
                                                v11 = 90
                                                fmt.Println(fi21, v11)
                                            }
                                        } else {
                                            switch "hello" {
                                            case "":
                                                if (973 < 1) {
                                                    var v46 uint64 = (v22 << ((0) & 63))
                                                    _ = v46
                                                } else {
                                                    var v47 uint64 = v22
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v48 int16 = 648
                                                        _ = v48
                                                    }
                                                    _ = v47
                                                }
                                            default:
                                                v49 := GenS8{F0: (fi21 & (fi21 >> ((1) & 63))), F1: 522}
                                                _ = v49
                                            }
                                            var v50 int32 = v19
                                            _ = v50
                                        }
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        fmt.Println(fi21)
                                    case "hello":
                                        v51 := v4[v11]
                                        _ = v51
                                    }
                                    var v52 uint8 = uint8(v22)
                                    _ = v52
                                }
                                v53 := GenS2{F0: "", F1: 736, F2: 127}
                                _ = v53
                            }
                            _ = v27
                        }
                        _ = fi21
                        _ = v22
                    }
                } else {
                    var v54 int8 = -119
                    _ = v54
                }
                _ = fi18
                _ = v19
            }
            var v55 int32 = -1
            _ = v55
        }
    case 0:
        v11 = v11
    case 2147483647:
        v56 := GenS9{F0: 65535, F1: 42}
        _ = v56
    }
    v1 += v1
    if ((255 <= 42) || v14) {
        var v57 int = 127
        v57 = -70
        switch v57 {
        case 9223372036854775807:
            for fi58 := 0; fi58 < 10; fi58++ {
                v59 := GenS10{F0: (v11 % (((v11 << ((4294967295) & 63))) | 1)), F1: 127, F2: 524}
                fmt.Println(v11)
                _ = fi58
                _ = v59
            }
        }
        _ = v57
    } else {
        v1 -= v1
    }
    var v60 uint16 = 65535
    fmt.Println(v1)
    if (0.1 == 0.1) {
        for fi61 := 0; fi61 < 12; fi61++ {
            var v62 uint8 = 85
            var v63 uint32 = (^(v11))
            switch "a" {
            case "\n\t":
                v64 := GenS5{F0: -32768, F1: v60}
                _ = v64
            case "":
                v3 = append(v3, (fi61 >> ((1) & 63)))
            case "hello":
                switch (fi61 << ((uint(v11)) & 63)) {
                case -363:
                    var v65 uint64 = 340
                    _ = v65
                case 32767:
                    v66 := GenS8{F0: fi61, F1: (v60 / ((v60) | 1))}
                    _ = v66
                default:
                    var v67 int32 = -2147483648
                    _ = v67
                }
            }
            _ = fi61
            _ = v62
            _ = v63
        }
        _ = len(v2)
        fmt.Println(v1)
    } else {
        var v68 int64 = genFn0(352, 696)
        v69 := GenS4{F0: 113, F1: true}
        _ = v68
        _ = v69
    }
    v60 = 255
    v11 -= (v11 / (((^(v11))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v60
}
