package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint32; F2 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 uint32 }
func genFn0(p0 int16, p1 string, p2 int32) uint {
    p1 = p1
    v7 := []int{-1}
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS1{F0: 4294967295, F1: (fi8 | (fi8 + (fi8 - fi8))), F2: -87}
        fmt.Println(p0, p1, fi8)
        if false {
            fmt.Println(p0, fi8, p2)
            switch ("foo bar" + "hello") {
            case "hello":
                fmt.Println(p2)
            case "a":
                for fi10 := 0; fi10 < 16; fi10++ {
                    if true {
                        if (362 > 499) {
                            p1 = (p1 + "\n\t")
                            if false {
                                v7 = append(v7, (fi8 % ((fi8) | 1)))
                            } else {
                                p2 = (p2 & (p2 + -2147483648))
                            }
                        }
                        p1 = ((p1 + (p1 + "hello")) + ("\n\t" + ("foo bar" + "")))
                        var v11 int16 = (p0 / ((p0) | 1))
                        _ = v11
                    } else {
                        p0 = (-(p0))
                        if false {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            switch (("" + p1) + p1) {
                            case "":
                                v12 := GenS2{F0: 65535}
                                _ = v12
                            case "hello":
                                fmt.Println(p1, p0)
                            case "foo bar":
                                if true {
                                    var v13 uint16 = 111
                                    _ = v13
                                }
                                var v14 int8 = 94
                                _ = v14
                            default:
                                p1 = p1
                            }
                            if false {
                                if len(v7) > 0 {
                                    v7[0] = fi10
                                }
                            }
                        }
                        switch (fi10 | fi8) {
                        case -123:
                            if true {
                                p0 -= p0
                            }
                            switch p1 {
                            case "hello":
                                var v15 uint32 = 42
                                _ = v15
                            case "":
                                p2 = 0
                            }
                        case 315:
                            var v16 int = (fi8 ^ (fi10 << ((18446744073709551615) & 63)))
                            _ = v16
                        }
                        if true {
                            fmt.Println(p2, p1, p2)
                            if false {
                                switch 32767 {
                                case 140:
                                    v17 := GenS3{F0: 0, F1: 0, F2: false}
                                    var v18 int64 = 9223372036854775807
                                    v19 := GenS2{F0: 831}
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                case -9223372036854775808:
                                    v7 = append(v7, (fi8 / ((-360) | 1)))
                                case -891:
                                    var v20 int8 = 0
                                    _ = v20
                                default:
                                    v21 := GenS2{F0: 242}
                                    p1 = p1
                                    if false {
                                        _ = len(v7)
                                        _ = len(v7)
                                    }
                                    var v22 int16 = (p0 << ((368) & 63))
                                    var v23 int = 32767
                                    var v24 int64 = -9223372036854775808
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                                v25 := GenS4{F0: false}
                                p1 = ("" + p1)
                                v7 = append(v7, 42)
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                _ = v25
                            }
                            switch (p1 + p1) {
                            case "":
                                fmt.Println(p2)
                            default:
                                var v26 uint = 0
                                _ = v26
                            }
                        } else {
                            if len(v7) > 0 {
                                v7[0] = (fi10 >> ((927) & 63))
                            }
                            if true {
                                var v27 int8 = -1
                                var v28 string = "a"
                                _ = v27
                                _ = v28
                            }
                        }
                        p2 = (p2 ^ (p2 * -512))
                    }
                    p1 = "foo bar"
                    _ = fi10
                }
            default:
                p2 -= p2
            }
        }
        _ = fi8
        _ = v9
    }
    if true {
        fmt.Println(p1, p0)
        fmt.Println(p2, p2, p0)
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 int8 = 127
            var v31 uint8 = 172
            if len(v7) > 0 {
                v7[0] = fi29
            }
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        switch fi32 {
        case -9:
            var v33 int64 = -1
            _ = v33
        default:
            var v34 int16 = -32768
            _ = v34
        }
        p0 = p0
        var v35 float32 = ((0.0 * 3.14159) + (float32(float64((1e10 - 1.5))) / 0.1))
        _ = fi32
        _ = v35
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 826
}

func main() {
    var v1 string = ""
    var v2 uint8 = 128
    var v3 float32 = 1e10
    v4 := []uint32{219, 2147483648, 65535}
    v5 := make(map[int]int)
    v6 := make(map[int]bool)
    var v36 uint32 = 2147483647
    fmt.Println(v36, v36)
    for fi37 := 0; fi37 < 6; fi37++ {
        fmt.Println(fi37, v2)
        fmt.Println(v1, v2)
        var v38 int16 = -494
        _ = fi37
        _ = v38
    }
    v39 := GenS5{F0: (v36 - v36), F1: (v1 + "a")}
    v3 = v3
    fmt.Println(v1)
    if (0 == 177) {
        for fi40 := 0; fi40 < 9; fi40++ {
            if true {
                if true {
                    v41 := GenS1{F0: 2, F1: (fi40 / ((717) | 1)), F2: -1}
                    _ = v41
                }
                fmt.Println(v36, v1)
                v2 = v2
            }
            if true {
                _ = len(v6)
            }
            if false {
                for fi42 := 0; fi42 < 11; fi42++ {
                    switch ("hello" + "") {
                    case "hello":
                        switch "" {
                        case "hello":
                            v3 += (((v3 - -2.25) * float32((3.14159 + 3.14159))) - (v3 * v3))
                        case "foo bar":
                            v36 = 65535
                        case "\n\t":
                            v4 = append(v4, uint32(v2))
                        default:
                            if false {
                                var v43 int16 = -1
                                v44 := GenS1{F0: 636, F1: 4294967295, F2: 0}
                                if ((fi40 & (-(fi40))) >= fi42) {
                                    switch v1 {
                                    case "hello":
                                        v45 := GenS4{F0: (127 == 50)}
                                        _ = v45
                                    default:
                                        var v46 int16 = 1
                                        _ = v46
                                    }
                                } else {
                                    v43 -= v43
                                    switch fi40 {
                                    case -9223372036854775808:
                                        v1 = ""
                                    }
                                }
                                _ = v43
                                _ = v44
                            }
                            var v47 uint8 = (v2 ^ 127)
                            var v48 int = fi40
                            v48 += 9223372036854775807
                            _ = v47
                            _ = v48
                        }
                        switch fi40 {
                        case 127:
                            var v49 int64 = 42
                            _ = v49
                        }
                    case "":
                        v36 -= 183
                        v36 -= v36
                        if false {
                            v4 = append(v4, (v36 - 785))
                        } else {
                            if ((0 <= 42) && ((-2.25 == float64(((float32(1e10) + v3) - v3))) && false)) {
                                v50 := GenS6{F0: -611}
                                var v51 uint64 = 4294967295
                                _ = v50
                                _ = v51
                            }
                            v52 := GenS7{F0: (fi40 ^ (fi42 * fi40)), F1: (v36 / ((1) | 1))}
                            v36 -= v36
                            if len(v4) > 0 {
                                v4[0] = v36
                            }
                            v53 := GenS6{F0: 0}
                            _ = v52
                            _ = v53
                        }
                    }
                    v54 := GenS4{F0: (100 <= 42)}
                    v55 := GenS3{F0: 42, F1: v36, F2: false}
                    var v56 uint16 = 0
                    _ = fi42
                    _ = v54
                    _ = v55
                    _ = v56
                }
                for fi57 := 0; fi57 < 13; fi57++ {
                    fmt.Println(v1, v3)
                    var v58 bool = false
                    _ = fi57
                    _ = v58
                }
                if (true && false) {
                    var v59 int8 = -42
                    v4 = append(v4, 42)
                    v2 += v2
                    _ = v59
                }
            }
            _ = fi40
        }
    } else {
        var v60 uint32 = v36
        v36 = 65535
        _ = v60
    }
    for fi61 := 0; fi61 < 8; fi61++ {
        var v62 int8 = 100
        _ = fi61
        _ = v62
    }
    var v63 uint = 18446744073709551615
    for fi64 := 0; fi64 < 5; fi64++ {
        fmt.Println(v2)
        _ = fi64
    }
    fmt.Println(v3, v63, v2)
    _ = len(v4)
    var v65 uint = v63
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v63:", v63)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v65
}
