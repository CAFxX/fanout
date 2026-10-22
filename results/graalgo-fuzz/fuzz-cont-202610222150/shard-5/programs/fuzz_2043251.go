package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int; F1 int32; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 uint64; F2 int32 }
type GenS6 struct { F0 bool; F1 uint8 }
func genFn0(p0 int8, p1 int32, p2 uint16) uint16 {
    v7 := GenS1{F0: 65535, F1: 932, F2: (p1 * (p1 ^ p1))}
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 bool = false
        _ = fi8
        _ = v9
    }
    p1 = (p1 + (p1 | p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p2
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 881
    var v3 int16 = 32767
    v4 := make(map[string]uint32)
    var v5 uint = 255
    v6 := make(map[string]uint32)
    for fi10 := 0; fi10 < 16; fi10++ {
        if false {
            var v11 int32 = 1
            v2 -= (v2 + 160)
            _ = v11
        } else {
            for fi12 := 0; fi12 < 13; fi12++ {
                switch "foo bar" {
                case "\n\t":
                    v13 := GenS1{F0: 42, F1: 1, F2: -1}
                    _ = v13
                case "hello":
                    if false {
                        if (137 >= 128) {
                            if (1.5 > 3.14159) {
                                if false {
                                    v14 := GenS2{F0: (v3 - (v3 ^ (v3 ^ int16(fi12)))), F1: 2147483647}
                                    v5 += 0
                                    _ = v14
                                } else {
                                    v3 += (v3 & v3)
                                    fmt.Println(v5, v2)
                                    var v15 int16 = v3
                                    _ = v15
                                }
                                v16 := GenS2{F0: (v3 & 42), F1: 6}
                                fmt.Println(fi12, v2, v2)
                                _ = v16
                            }
                            v3 -= -573
                            if (false && ((fi12 * 9223372036854775807) > 9223372036854775807)) {
                                delete(v1, "\n\t")
                                var v17 int = (^(fi10))
                                _ = v17
                            }
                        }
                        v5 = (v5 & v5)
                        v18 := GenS1{F0: 42, F1: -9223372036854775808, F2: -522}
                        _ = v18
                    } else {
                        var v19 uint64 = (v2 & v2)
                        v20 := GenS1{F0: 42, F1: -1, F2: -1}
                        _ = v19
                        _ = v20
                    }
                case "":
                    switch "foo bar" {
                    case "hello":
                        var v21 uint64 = v2
                        _ = v21
                    }
                default:
                    fmt.Println(v3, fi12, fi12)
                    v22 := []int{(fi10 / ((fi12) | 1))}
                    var v23 uint16 = 0
                    _ = v22
                    _ = v23
                }
                var v24 string = (("hello" + "foo bar") + "")
                _ = fi12
                _ = v24
            }
            if (v3 != v3) {
                if (65535 <= 209) {
                    v2 += (^(v2))
                } else {
                    var v25 float32 = 0.0
                    _ = v25
                }
            } else {
                v26 := v6["a"]
                _ = v26
            }
            v3 += 75
        }
        var v27 uint = v5
        _ = fi10
        _ = v27
    }
    var v28 int16 = v3
    if ((true && true) && (((("foo bar" + "hello") + (("" + "\n\t") + (("hello" + "a") + "hello"))) + "\n\t") <= "\n\t")) {
        v2 -= (v2 | (v2 % (((v2 / (((v2 >> ((v5) & 63))) | 1))) | 1)))
        if (v5 < v5) {
            var v29 string = ("" + "\n\t")
            v30 := GenS2{F0: (v28 >> ((18446744073709551615) & 63)), F1: 132}
            _ = v29
            _ = v30
        }
        v31 := GenS3{F0: 128, F1: 42, F2: 0}
        _ = v31
    }
    var v32 int = int(v3)
    if ((v28 & -32768) != (v3 >> (((v5 & v5)) & 63))) {
        var v33 int = (v32 ^ 128)
        for fi34 := 0; fi34 < 19; fi34++ {
            v33 += (v32 >> ((v5) & 63))
            _ = fi34
        }
        _ = v33
    } else {
        var v35 float64 = float64(float32(1.5))
        v36 := []int{v32, (-(v32)), v32}
        fmt.Println(v2)
        _ = v35
        _ = v36
    }
    if (42 == 649) {
        v37 := GenS3{F0: (v32 ^ v32), F1: -1, F2: uint16(v2)}
        if true {
            fmt.Println(v5, v28)
            var v38 bool = false
            _ = v38
        }
        _ = v37
    }
    if (uint(v2) < (v5 / (((^(v5))) | 1))) {
        v39 := GenS2{F0: v3, F1: 1}
        _ = v39
    }
    var v40 uint = v5
    v41 := []int{(v32 & v32), (-(v32)), (v32 + v32), v32}
    v32 += v32
    for fi42 := 0; fi42 < 5; fi42++ {
        switch v32 {
        case -2147483648:
            v43 := GenS4{F0: -209}
            _ = v43
        case 128:
            switch v32 {
            case 189:
                if true {
                    switch ("foo bar" + ("" + ("a" + "\n\t"))) {
                    case "\n\t":
                        fmt.Println(v32)
                    case "a":
                        var v44 bool = true
                        _ = v44
                    case "":
                        switch v32 {
                        case 555:
                            for fi45 := 0; fi45 < 10; fi45++ {
                                if false {
                                    if false {
                                        v46 := GenS5{F0: 65535, F1: 994, F2: -1}
                                        v40 += (v5 | v40)
                                        v3 -= v28
                                        _ = v46
                                    }
                                    if len(v41) >= 2 {
                                        v41 = v41[1:len(v41)]
                                    }
                                    v2 += v2
                                    if false {
                                        v2 += uint64(v40)
                                    }
                                } else {
                                    v3 = (v3 % (((v3 + (v3 ^ (v28 ^ v28)))) | 1))
                                    fmt.Println(v40)
                                    var v47 int64 = -3
                                    _ = v47
                                }
                                var v48 uint64 = 0
                                _ = fi45
                                _ = v48
                            }
                        default:
                            var v49 bool = false
                            _ = v49
                        }
                    default:
                        v50 := GenS1{F0: 4294967295, F1: 1, F2: 90}
                        _ = v50
                    }
                    delete(v4, "")
                    for fi51 := 0; fi51 < 12; fi51++ {
                        v28 = v3
                        var v52 int32 = -2147483648
                        var v53 string = (("\n\t" + (("a" + "a") + "a")) + "")
                        var v54 uint16 = 830
                        _ = fi51
                        _ = v52
                        _ = v53
                        _ = v54
                    }
                } else {
                    v55 := GenS2{F0: (v3 | 819), F1: -2147483648}
                    var v56 string = (("" + "a") + ("hello" + ""))
                    fmt.Println(v2, v5, v28)
                    _ = v55
                    _ = v56
                }
            case 128:
                v3 = (v3 & (v28 & (v28 >> (((^(v40))) & 63))))
            case -9223372036854775808:
                var v57 int = v32
                _ = v57
            default:
                var v58 uint64 = (v2 + 392)
                _ = v58
            }
        case 42:
            if false {
                var v59 uint16 = 708
                v32 -= 798
                v1[(("hello" + "hello") + "hello")] = (-32 < 42)
                _ = v59
            } else {
                v60 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: 732}
                if true {
                    v61 := GenS6{F0: true, F1: 112}
                    var v62 float32 = (float32(((-2.25 + 1e10) / (1.5 * -2.25))) * float32((0.1 + (0.0 + (float64((((1.5 / (float32((0.1 * 3.14159)) / float32(float64(-2.25)))) + (0.0 * (-2.25 * (-2.25 - 3.14159)))) / float32((1e10 - ((float64(3.14159) - (0.0 - (-2.25 + float64(((1.5 / 1.5) / 1.5))))) + float64(float32(1e10))))))) - (0.0 * (1.5 * 3.14159)))))))
                    _ = v61
                    _ = v62
                }
                v5 -= v5
                _ = v60
            }
        }
        _ = fi42
    }
    v2 += v2
    var v63 bool = ((1e10 >= 0.1) || (true && ((v32 + v32) >= -9223372036854775808)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v63
}
