package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int16; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int32 }
func genFn0() int8 {
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint8 = 42
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 210
    switch "\n\t" {
    case "hello":
        v9 := GenS1{F0: 4294967295}
        _ = v9
    case "a":
        v10 := GenS2{F0: 368, F1: 550, F2: 127}
        _ = v10
    }
    var v11 uint32 = 478
    _ = v11
    _ = v8
    return -6
}
func genFn1(p0 int, p1 uint64) uint64 {
    v12 := GenS1{F0: p1}
    switch ("a" + "foo bar") {
    case "":
        switch p0 {
        case 0:
            var v13 bool = true
            _ = v13
        case 128:
            p0 += (^(p0))
        default:
            v14 := GenS3{F0: 128, F1: uint8(p1), F2: 234}
            _ = v14
        }
    }
    if ((-1 == 1) || true) {
        v15 := GenS4{F0: -1, F1: p1, F2: 127}
        if (!("\n\t" == ("\n\t" + "\n\t"))) {
            v16 := GenS5{F0: -32768}
            switch "foo bar" {
            case "hello":
                v17 := []int{}
                _ = v17
            default:
                v18 := GenS6{F0: -372, F1: (p0 | p0)}
                _ = v18
            }
            v19 := GenS5{F0: 1000}
            _ = v16
            _ = v19
        }
        p1 = p1
        _ = v15
    } else {
        var v20 uint16 = 42
        for fi21 := 0; fi21 < 6; fi21++ {
            v22 := []int{}
            p1 -= p1
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    if true {
        v23 := GenS6{F0: -2147483648, F1: p0}
        p1 = p1
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v12
    return (p1 ^ 18446744073709551615)
}

func main() {
    var v1 uint8 = 243
    var v2 uint64 = 225
    v3 := []string{"hello"}
    var v4 uint16 = 172
    v5 := make(map[uint32]uint32)
    if (false && (v2 >= 1)) {
        v24 := v5[0]
        v1 -= uint8(v24)
        var v25 float32 = 0.0
        _ = v24
        _ = v25
    } else {
        v1 -= v1
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        v27 := GenS1{F0: (^(v2))}
        switch "a" {
        case "hello":
            fmt.Println(v2, v4, v2)
        case "a":
            if (true && ((true || (!(!false))) && true)) {
                v1 += uint8(v2)
                switch ("foo bar" + ("foo bar" + "")) {
                case "\n\t":
                    v4 += 1
                case "":
                    for fi28 := 0; fi28 < 18; fi28++ {
                        switch 767 {
                        case 0:
                            v4 -= v4
                        default:
                            switch (fi28 * (fi28 | fi28)) {
                            case -9223372036854775808:
                                var v29 int32 = 1
                                _ = v29
                            default:
                                v2 = v2
                            }
                        }
                        v2 = 2
                        v4 = (v4 % (((^(v4))) | 1))
                        _ = fi28
                    }
                }
                var v30 uint64 = v2
                _ = v30
            } else {
                var v31 bool = (-418 <= 0)
                switch fi26 {
                case 4294967295:
                    v1 += (v1 ^ (v1 % ((255) | 1)))
                default:
                    if (-111 >= 9223372036854775807) {
                        fmt.Println(v1)
                        for fi32 := 0; fi32 < 10; fi32++ {
                            if false {
                                if (2147483648 <= 334) {
                                    var v33 bool = v31
                                    _ = len(v5)
                                    v5[2147483647] = 251
                                    _ = v33
                                }
                                var v34 float64 = 1e10
                                _ = v34
                            } else {
                                v2 = genFn1(-546, 42)
                                v4 = 18
                            }
                            _ = fi32
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        switch ("hello" + "a") {
                        case "":
                            for fi35 := 0; fi35 < 12; fi35++ {
                                if true {
                                    switch (fi26 ^ -111) {
                                    case 806:
                                        v5[561] = 4294967295
                                    case 1:
                                        fmt.Println(v2, fi26, fi26)
                                    case 127:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    default:
                                        v3 = append(v3, "hello")
                                    }
                                    v31 = v31
                                    v36 := GenS1{F0: 744}
                                    _ = v36
                                }
                                v31 = ((v31 || (65535 < 553)) && v31)
                                _ = fi35
                            }
                        case "\n\t":
                            var v37 int16 = int16(v1)
                            _ = v37
                        default:
                            v38 := GenS5{F0: 918}
                            _ = v38
                        }
                        var v39 uint8 = (v1 % (((v1 - (v1 << ((2) & 63)))) | 1))
                        switch 2147483647 {
                        case 9223372036854775807:
                            v40 := GenS7{F0: "foo bar"}
                            _ = v40
                        case -1:
                            if v31 {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v4 = (v4 & (^(v4)))
                            }
                        case 42:
                            v39 += 0
                        }
                        _ = v39
                    }
                }
                v31 = false
                _ = v31
            }
        default:
            switch (fi26 ^ 127) {
            case 32767:
                var v41 bool = true
                _ = v41
            default:
                v4 -= (^(v4))
            }
        }
        _ = fi26
        _ = v27
    }
    v42 := GenS8{F0: 42}
    var v43 uint32 = 1
    fmt.Println(v1)
    fmt.Println(v1, v43, v4)
    for fi44 := 0; fi44 < 5; fi44++ {
        v43 = 758
        _ = fi44
    }
    fmt.Println(v1, v2, v4)
    for fi45 := 0; fi45 < 1; fi45++ {
        if (!true) {
            switch "a" {
            case "a":
                v46 := GenS2{F0: 0, F1: 1, F2: genFn0()}
                _ = v46
            default:
                for fi47 := 0; fi47 < 14; fi47++ {
                    v5[v43] = v43
                    _ = fi47
                }
            }
        } else {
            v4 -= 73
        }
        v48 := v5[(v43 * uint32(v1))]
        v49 := GenS1{F0: genFn1(fi45, (v2 >> ((1) & 63)))}
        _ = fi45
        _ = v48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
