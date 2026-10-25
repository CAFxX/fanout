package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint; F1 uint; F2 int }
func genFn0(p0 uint16, p1 int32, p2 uint8) uint16 {
    p0 = 255
    v7 := GenS1{F0: 359, F1: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 264
}
func genFn1(p0 int64, p1 uint32, p2 uint) int8 {
    switch "hello" {
    case "\n\t":
        var v8 int = -1
        _ = v8
    case "foo bar":
        p1 += (p1 & 65535)
    default:
        if true {
            fmt.Println(p0, p2, p2)
            var v9 int = 4294967295
            _ = v9
        } else {
            v10 := GenS2{F0: (p1 | (p1 ^ p1)), F1: p2}
            v11 := GenS1{F0: 255, F1: 42}
            _ = v10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        switch (fi12 ^ fi12) {
        case 4294967295:
            if false {
                p0 -= -588
            } else {
                switch (fi12 << ((p2) & 63)) {
                case 32767:
                    var v13 int8 = int8(p1)
                    _ = v13
                case -1:
                    switch -169 {
                    case -1:
                        if (!(-53 >= 8)) {
                            p2 -= (^(p2))
                            var v14 int8 = 0
                            v15 := GenS1{F0: 1, F1: 1}
                            _ = v14
                            _ = v15
                        }
                    case 128:
                        var v16 string = "a"
                        _ = v16
                    default:
                        if (0 == p0) {
                            v17 := GenS2{F0: uint32(p2), F1: 78}
                            var v18 uint = p2
                            _ = v17
                            _ = v18
                        }
                    }
                default:
                    p1 += (p1 | (p1 + 2147483647))
                }
            }
        case -835:
            fmt.Println(p1, p2)
        case -9223372036854775808:
            var v19 uint8 = 128
            _ = v19
        }
        var v20 int = (fi12 % ((fi12) | 1))
        _ = fi12
        _ = v20
    }
    v21 := GenS1{F0: 74, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 127
}

func main() {
    var v1 bool = true
    var v2 uint32 = 2
    var v3 uint = 4294967295
    var v4 uint8 = 2
    var v5 uint32 = 114
    v6 := []uint32{v5, 757, v2}
    if len(v6) > 0 {
        v6[0] = (^(v2))
    }
    var v22 uint16 = 255
    var v23 int16 = 32767
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        for fi25 := 0; fi25 < 18; fi25++ {
            v4 = v4
            v26 := GenS1{F0: (v22 ^ v22), F1: (v22 >> ((v3) & 63))}
            switch (fi24 >> ((18446744073709551615) & 63)) {
            case 32767:
                if len(v6) > 0 {
                    v6[0] = v5
                }
            case -638:
                if v1 {
                    var v27 uint32 = (^(v2))
                    v28 := GenS2{F0: 2147483648, F1: v3}
                    fmt.Println(v5)
                    _ = v27
                    _ = v28
                }
            case -9223372036854775808:
                v1 = (v1 || v1)
                v29 := GenS1{F0: (v22 | (v22 ^ v22)), F1: v22}
                _ = v29
            }
            _ = fi25
            _ = v26
        }
        v23 -= v23
        _ = fi24
    }
    v30 := GenS1{F0: v22, F1: v22}
    v31 := GenS2{F0: v5, F1: 2}
    var v32 int32 = -831
    if v1 {
        var v33 bool = (v32 >= (^(v32)))
        if (1 >= -870) {
            fmt.Println(v33, v33, v33)
            var v34 uint = 42
            fmt.Println(v1, v32)
            _ = v34
        } else {
            if len(v6) > 0 {
                v6[0] = (v5 | (v5 * (v5 - v5)))
            }
        }
        v6 = append(v6, (v2 | (v2 & 65535)))
        _ = v33
    } else {
        v1 = (!((v1 || (v4 >= 0)) && (42 < 925)))
    }
    var v35 uint32 = uint32(v3)
    if v1 {
        if (!false) {
            switch "\n\t" {
            case "foo bar":
                v36 := GenS1{F0: genFn0(v22, (v32 >> ((v3) & 63)), v4), F1: genFn0((v22 | (v22 | uint16(v35))), (v32 & (v32 + v32)), (v4 << ((v3) & 63)))}
                _ = v36
            case "\n\t":
                fmt.Println(v1, v32, v1)
            case "a":
                if v1 {
                    v37 := GenS3{F0: v3, F1: v3, F2: 1}
                    _ = v37
                } else {
                    var v38 uint = 65535
                    v32 -= -2147483648
                    _ = v38
                }
            }
        } else {
            for fi39 := 0; fi39 < 10; fi39++ {
                for fi40 := 0; fi40 < 5; fi40++ {
                    _ = len(v6)
                    _ = fi40
                }
                _ = fi39
            }
            fmt.Println(v22, v3)
            var v41 int8 = 42
            _ = v41
        }
    } else {
        for fi42 := 0; fi42 < 14; fi42++ {
            v3 -= (v3 * v3)
            var v43 uint16 = genFn0((v22 / (((v22 % (((v22 | (^(v22)))) | 1))) | 1)), v32, uint8(v3))
            _ = fi42
            _ = v43
        }
        for fi44 := 0; fi44 < 17; fi44++ {
            fmt.Println(v35, v5, v35)
            _ = fi44
        }
    }
    var v45 int8 = -1
    _ = len(v6)
    var v46 int16 = v23
    if v1 {
        _ = len(v6)
        for fi47 := 0; fi47 < 6; fi47++ {
            for fi48 := 0; fi48 < 8; fi48++ {
                fmt.Println(v35)
                if ((1.5 > float32(0.0)) || v1) {
                    v32 = v32
                    if (v23 < 0) {
                        _ = len(v6)
                    } else {
                        v6 = append(v6, v35)
                    }
                    if v1 {
                        fmt.Println(v4)
                        v4 -= 42
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                _ = fi48
            }
            v4 = (^(v4))
            _ = fi47
        }
    }
    v6 = append(v6, uint32(v3))
    fmt.Println(v1)
    v2 = (v5 | (^(v5)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
