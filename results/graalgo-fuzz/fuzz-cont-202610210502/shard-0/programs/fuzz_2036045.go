package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int16; F1 bool; F2 uint }
type GenS4 struct { F0 uint; F1 uint16 }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 65535
    v3 := []uint32{v2, 0}
    var v4 uint32 = 4294967295
    var v5 float32 = (float32(((1.5 + 3.14159) + 1.5)) / (1e10 * (1.5 / (0.1 * (-2.25 / float32((3.14159 + 0.1)))))))
    var v6 int8 = 12
    for fi7 := 0; fi7 < 20; fi7++ {
        v5 -= v5
        v2 += v2
        _ = fi7
    }
    switch -262 {
    case 32768:
        if (-2.25 > -2.25) {
            for fi8 := 0; fi8 < 19; fi8++ {
                if ((false || true) || true) {
                    if len(v3) > 0 {
                        v3[0] = (v4 << ((18446744073709551615) & 63))
                    }
                    fmt.Println(v5, v5)
                    for fi9 := 0; fi9 < 16; fi9++ {
                        v10 := GenS1{F0: 4294967295, F1: 42, F2: uint8(v4)}
                        if false {
                            fmt.Println(v2, v2, v5)
                            _ = len(v1)
                        } else {
                            switch fi8 {
                            case 4294967295:
                                v2 += v4
                            default:
                                var v11 uint16 = 32767
                                _ = v11
                            }
                        }
                        v2 = v4
                        fmt.Println(fi9, v4, fi9)
                        var v12 int32 = -600
                        _ = fi9
                        _ = v10
                        _ = v12
                    }
                } else {
                    _ = len(v3)
                    v2 = v2
                }
                if false {
                    var v13 uint8 = 2
                    _ = v13
                }
                v14 := v1[fi8]
                _ = fi8
                _ = v14
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                v4 += v4
                _ = fi15
            }
        } else {
            if true {
                switch "a" {
                case "\n\t":
                    switch 139 {
                    case 128:
                        if (false || false) {
                            if (false || true) {
                                _ = len(v3)
                            } else {
                                if true {
                                    fmt.Println(v5)
                                    v16 := v1[2147483647]
                                    var v17 uint64 = 755
                                    _ = v16
                                    _ = v17
                                }
                            }
                        } else {
                            var v18 int64 = -9223372036854775808
                            _ = len(v3)
                            if (255 != 42) {
                                v6 = 36
                                v2 += v2
                                for fi19 := 0; fi19 < 9; fi19++ {
                                    v2 = 2147483648
                                    for fi20 := 0; fi20 < 6; fi20++ {
                                        v6 -= v6
                                        _ = fi20
                                    }
                                    _ = fi19
                                }
                            } else {
                                v18 += (v18 | -1)
                                var v21 float64 = ((0.0 - 1.5) * ((1e10 * 3.14159) / float64(0.1)))
                                _ = v21
                            }
                            _ = v18
                        }
                    case -2147483648:
                        v22 := GenS2{F0: -1, F1: 22}
                        _ = v22
                    }
                default:
                    var v23 int32 = -1
                    _ = v23
                }
                v24 := GenS1{F0: 4294967295, F1: 4294967295, F2: 0}
                _ = v24
            } else {
                fmt.Println(v2, v4, v2)
                for fi25 := 0; fi25 < 19; fi25++ {
                    v26 := GenS3{F0: -1, F1: ((!(fi25 == fi25)) && ((32767 < 116) || true)), F2: 65535}
                    var v27 int64 = -793
                    switch (-(fi25)) {
                    case -884:
                        v28 := v1[fi25]
                        _ = v28
                    case 0:
                        v2 = (v4 & (v2 + (v4 % ((1) | 1))))
                    case -195:
                        v29 := v1[fi25]
                        _ = v29
                    }
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                v5 -= 1.5
            }
        }
    case -2147483648:
        v30 := GenS4{F0: 795, F1: 354}
        _ = v30
    default:
        var v31 string = ("foo bar" + ("\n\t" + "a"))
        _ = v31
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        for fi33 := 0; fi33 < 2; fi33++ {
            v5 -= v5
            _ = fi33
        }
        _ = fi32
    }
    var v34 float64 = (1.5 / 0.1)
    var v35 int16 = -32768
    for fi36 := 0; fi36 < 2; fi36++ {
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            fmt.Println(v2, v2, v4)
            v37 := v1[(^(fi36))]
            var v38 int16 = (-(v35))
            _ = v37
            _ = v38
        }
        _ = fi36
    }
    _ = len(v1)
    v6 = (v6 >> ((760) & 63))
    fmt.Println(v5, v2, v2)
    for fi39 := 0; fi39 < 9; fi39++ {
        _ = len(v1)
        var v40 int16 = (v35 << ((2) & 63))
        if (0 > 1) {
            v4 = v4
            var v41 int64 = -1
            _ = v41
        }
        _ = fi39
        _ = v40
    }
    if false {
        var v42 bool = true
        switch (("hello" + "a") + "a") {
        case "\n\t":
            var v43 uint32 = 2147483647
            _ = v43
        case "hello":
            fmt.Println(v5, v35)
        }
        _ = v42
    } else {
        for fi44 := 0; fi44 < 16; fi44++ {
            v35 -= (-(v35))
            v45 := GenS3{F0: v35, F1: false, F2: 2}
            _ = fi44
            _ = v45
        }
        v1[32767] = "foo bar"
        v46 := GenS4{F0: 2, F1: 255}
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
