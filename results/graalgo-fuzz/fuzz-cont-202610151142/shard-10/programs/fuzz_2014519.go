package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint64 }
type GenS2 struct { F0 int; F1 int; F2 uint }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int16 }
func genFn0(p0 uint32) uint {
    p0 = (p0 + (^(p0)))
    var v4 uint32 = (p0 & 226)
    switch "" {
    case "a":
        v5 := GenS1{F0: 9223372036854775807, F1: -9, F2: 18446744073709551615}
        _ = v5
    default:
        if (!(!true)) {
            if (0.0 < 0.1) {
                p0 = (v4 * 2)
                v6 := GenS2{F0: -1, F1: -446, F2: 571}
                p0 = v4
                _ = v6
            }
            if (v4 > 386) {
                if (false || true) {
                    v7 := []int{128, -1, -461}
                    v8 := GenS1{F0: -1, F1: -11, F2: 18446744073709551615}
                    var v9 uint16 = 29
                    _ = v7
                    _ = v8
                    _ = v9
                } else {
                    var v10 int = 9223372036854775807
                    _ = v10
                }
                v11 := GenS1{F0: -914, F1: -56, F2: 843}
                v4 -= (^(v4))
                _ = v11
            } else {
                p0 -= p0
                for fi12 := 0; fi12 < 13; fi12++ {
                    var v13 uint16 = 32767
                    _ = fi12
                    _ = v13
                }
                v14 := []int{}
                _ = v14
            }
        } else {
            p0 = v4
        }
    }
    if false {
        var v15 uint64 = 678
        _ = v15
    }
    _ = p0
    _ = v4
    return 82
}
func genFn1(p0 int8, p1 uint16) int16 {
    switch 2147483647 {
    case 184:
        p1 = (p1 * p1)
    case 0:
        p0 = p0
    case 128:
        for fi16 := 0; fi16 < 14; fi16++ {
            p1 += p1
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := []int{(fi17 / (((fi17 << ((1) & 63))) | 1)), (fi17 - (fi17 + (fi16 / ((2147483647) | 1)))), fi17, fi17}
                var v19 int32 = -911
                if len(v18) > 0 {
                    v18[0] = -2147483648
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            switch fi16 {
            case -308:
                v20 := []int{(fi16 >> ((18446744073709551615) & 63)), (fi16 ^ fi16), fi16}
                _ = v20
            case 32767:
                v21 := []int{(^(fi16)), fi16, (fi16 + fi16)}
                _ = v21
            case 42:
                v22 := GenS2{F0: 9223372036854775807, F1: fi16, F2: 1}
                _ = v22
            default:
                var v23 uint8 = 137
                _ = v23
            }
            _ = fi16
        }
    default:
        v24 := GenS2{F0: 42, F1: 737, F2: 191}
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        if true {
            p1 += p1
            v26 := []int{-275, fi25, 413, -228}
            _ = v26
        } else {
            v27 := GenS3{F0: p0, F1: 1}
            p0 = 1
            _ = v27
        }
        var v28 uint16 = 2
        _ = fi25
        _ = v28
    }
    p0 += (p0 + (p0 + 42))
    var v29 int = 4294967295
    v30 := []int{1}
    _ = p0
    _ = p1
    _ = v29
    _ = v30
    return 42
}

func main() {
    var v1 int = -2147483648
    var v2 uint8 = 2
    var v3 int8 = 0
    v31 := GenS4{F0: 0, F1: 32767, F2: (v2 / ((v2) | 1))}
    v32 := []int{(v1 / ((425) | 1)), (v1 % ((1) | 1)), (-(v1))}
    v3 = v3
    v33 := GenS3{F0: (v3 >> ((1) & 63)), F1: 2}
    for fi34 := 0; fi34 < 17; fi34++ {
        if false {
            for fi35 := 0; fi35 < 5; fi35++ {
                v3 -= 90
                _ = fi35
            }
            for fi36 := 0; fi36 < 6; fi36++ {
                switch v1 {
                case -463:
                    fmt.Println(fi34, v1)
                default:
                    v37 := GenS5{F0: 42, F1: (-2.25 > -2.25)}
                    _ = v37
                }
                v1 = -494
                _ = fi36
            }
            v1 = 4294967295
        } else {
            v2 += (^(v2))
        }
        var v38 int = fi34
        _ = fi34
        _ = v38
    }
    if true {
        for fi39 := 0; fi39 < 3; fi39++ {
            var v40 int8 = v3
            v41 := GenS6{F0: -1}
            if false {
                v2 = 0
            } else {
                var v42 int = v1
                v43 := GenS7{F0: 18446744073709551615}
                var v44 int64 = 9223372036854775807
                _ = v42
                _ = v43
                _ = v44
            }
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    fmt.Println(v1, v3)
    v2 = v2
    var v45 uint8 = (v2 ^ (^(v2)))
    if len(v32) > 0 {
        v32[0] = v1
    }
    var v46 uint8 = (v2 * (v2 ^ 255))
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    v47 := GenS8{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v45
    _ = v46
    _ = v47
}
