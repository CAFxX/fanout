package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int64; F1 bool; F2 uint }
type GenS5 struct { F0 int8; F1 int8 }
func genFn0() uint64 {
    v7 := GenS1{F0: 2, F1: 65535, F2: 9223372036854775807}
    v8 := GenS1{F0: 255, F1: 0, F2: 420}
    _ = v7
    _ = v8
    return 666
}
func genFn1(p0 int, p1 uint64, p2 uint) int32 {
    v9 := GenS2{F0: -413, F1: -9223372036854775808}
    for fi10 := 0; fi10 < 7; fi10++ {
        p1 += p1
        for fi11 := 0; fi11 < 12; fi11++ {
            v12 := []int{(p0 | p0)}
            if (true && ((4 > -612) && true)) {
                switch "hello" {
                case "\n\t":
                    v12 = append(v12, (fi11 & fi10))
                case "foo bar":
                    p0 -= (fi11 / (((fi11 / ((0) | 1))) | 1))
                case "":
                    v12 = append(v12, p0)
                default:
                    var v13 uint8 = 0
                    _ = v13
                }
                var v14 uint = 255
                var v15 uint64 = uint64(v14)
                _ = v14
                _ = v15
            }
            var v16 int16 = 0
            _ = fi11
            _ = v12
            _ = v16
        }
        p0 = -9223372036854775808
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2147483647
}

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]bool)
    v3 := []uint32{1, 65535, 2147483648}
    var v4 uint8 = 42
    v5 := make(map[uint32]int)
    var v6 float32 = -2.25
    v5[647] = -288
    var v17 uint8 = v4
    if false {
        for fi18 := 0; fi18 < 18; fi18++ {
            if (18446744073709551615 <= 1) {
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(fi18, v1)
                case "":
                    fmt.Println(v17, v17, v17)
                }
                fmt.Println(v1)
                var v19 int16 = int16(v1)
                _ = v19
            } else {
                fmt.Println(v4, fi18, v6)
            }
            if (v1 < 0) {
                v5[42] = fi18
                v6 -= 0.0
                fmt.Println(v6, v4, fi18)
            }
            for fi20 := 0; fi20 < 6; fi20++ {
                v6 = -2.25
                var v21 uint32 = 1
                var v22 int32 = -452
                v21 -= (v21 % ((65535) | 1))
                if (!(78 < 255)) {
                    v23 := GenS2{F0: -32768, F1: 48}
                    if false {
                        if false {
                            var v24 uint32 = v21
                            _ = v24
                        } else {
                            var v25 uint8 = (^(v17))
                            switch "hello" {
                            case "hello":
                                var v26 int = fi18
                                _ = v26
                            case "":
                                v17 = (v25 % ((v4) | 1))
                            default:
                                fmt.Println(v4, v22)
                            }
                            _ = v25
                        }
                    }
                    v1 = v1
                    _ = v23
                } else {
                    if (2 == 42) {
                        v22 = 1
                        v6 -= v6
                        switch fi20 {
                        case 128:
                            var v27 int8 = (v1 | v1)
                            _ = v27
                        case 42:
                            if (!(1e10 != ((1.5 + 1.5) * (0.1 * (0.1 / 1.0))))) {
                                v28 := GenS1{F0: 65535, F1: 624, F2: 2147483647}
                                fmt.Println(v6)
                                _ = v28
                            } else {
                                v29 := GenS2{F0: 1000, F1: fi18}
                                var v30 uint16 = 0
                                v31 := GenS3{F0: v1, F1: 32767}
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        default:
                            switch 860 {
                            case 42:
                                if (false || (!false)) {
                                    if ((((0.1 / 1.0) * 1e10) < float64((v6 / (v6 + (((v6 - -2.25) / v6) + float32(0.1)))))) || true) {
                                        if false {
                                            switch fi18 {
                                            case 128:
                                                v6 -= (v6 * 0.1)
                                            case 4294967295:
                                                var v32 float32 = (0.1 / v6)
                                                _ = v32
                                            }
                                        }
                                    } else {
                                        v22 -= genFn1(fi20, 1, 0)
                                        v1 = v1
                                        var v33 uint32 = 649
                                        _ = v33
                                    }
                                    var v34 string = "\n\t"
                                    _ = v34
                                }
                            case 32768:
                                v1 += 80
                            case -912:
                                v1 = -67
                            }
                        }
                    }
                    v35 := GenS1{F0: 2, F1: 4294967295, F2: (fi20 & (fi18 << ((274) & 63)))}
                    _ = v35
                }
                fmt.Println(v4)
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = fi18
        }
    } else {
        if len(v3) > 0 {
            v3[0] = 737
        }
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        v37 := GenS4{F0: -1, F1: true, F2: 427}
        _ = fi36
        _ = v37
    }
    var v38 int16 = 0
    v38 -= 32767
    fmt.Println(v38, v38)
    var v39 string = "\n\t"
    v40 := GenS4{F0: 0, F1: true, F2: 1}
    var v41 string = v39
    _ = len(v3)
    for fi42 := 0; fi42 < 17; fi42++ {
        var v43 int32 = 1
        delete(v5, 4294967295)
        _ = fi42
        _ = v43
    }
    switch 984 {
    case -2147483648:
        v44 := GenS5{F0: v1, F1: v1}
        _ = v44
    default:
        var v45 float64 = 1e10
        _ = v45
    }
    v38 += -619
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
