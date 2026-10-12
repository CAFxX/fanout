package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 int64 }
func genFn0(p0 int8, p1 int32, p2 int32) string {
    var v4 uint64 = 2
    v4 = (v4 >> ((18446744073709551615) & 63))
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 bool = (0 == (v4 ^ v4))
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        p1 -= (p1 ^ p2)
        var v8 int8 = (p0 + -27)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return ("\n\t" + "\n\t")
}
func genFn1(p0 bool, p1 int8, p2 uint16) int16 {
    if true {
        v9 := GenS1{F0: (-2.25 > ((1.5 + ((-2.25 - (-2.25 * (((-2.25 * 0.0) - 0.0) + ((3.14159 - -2.25) - float64((1.5 / 1.0)))))) * 0.1)) + 3.14159)), F1: (p2 % (((p2 * (p2 * p2))) | 1)), F2: 4294967295}
        var v10 uint32 = 530
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 int = fi11
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 uint32 = 450
    v2 := make(map[int]uint32)
    var v3 int64 = 0
    fmt.Println(v1, v3)
    switch -30 {
    case 127:
        var v13 float32 = -2.25
        _ = v13
    case 849:
        _ = len(v2)
    case 128:
        var v14 uint64 = 2
        _ = v14
    }
    if (0 != 0) {
        v15 := GenS2{F0: 123}
        _ = v15
    } else {
        fmt.Println(v1, v1)
        v16 := GenS1{F0: true, F1: 0, F2: 0}
        _ = v16
    }
    if (!false) {
        v1 -= 2147483647
        delete(v2, 9223372036854775807)
    } else {
        for fi17 := 0; fi17 < 18; fi17++ {
            if (5 < 42) {
                v18 := v2[(-(fi17))]
                var v19 uint16 = 0
                if (0 < uint8(v18)) {
                    v1 = (^(v1))
                    v3 = -9223372036854775808
                    fmt.Println(v1, v18)
                } else {
                    v20 := GenS3{F0: fi17}
                    _ = v20
                }
                _ = v18
                _ = v19
            } else {
                v1 = v1
                v1 += v1
            }
            v2[fi17] = v1
            if (777 > 113) {
                var v21 uint8 = 42
                for fi22 := 0; fi22 < 15; fi22++ {
                    v23 := GenS1{F0: false, F1: 65535, F2: (fi22 % (((fi17 / ((fi17) | 1))) | 1))}
                    v21 += (v21 % ((2) | 1))
                    _ = fi22
                    _ = v23
                }
                _ = v21
            }
            _ = fi17
        }
        for fi24 := 0; fi24 < 2; fi24++ {
            v3 -= (v3 * -9223372036854775808)
            _ = fi24
        }
        switch ("a" + "foo bar") {
        case "\n\t":
            switch -1 {
            case 128:
                v3 -= (v3 * v3)
            case 538:
                var v25 int16 = 1
                _ = v25
            default:
                v26 := GenS4{F0: -32768, F1: v3}
                _ = v26
            }
        default:
            v27 := GenS2{F0: 294}
            _ = v27
        }
    }
    v28 := GenS4{F0: 1000, F1: 121}
    for fi29 := 0; fi29 < 20; fi29++ {
        for fi30 := 0; fi30 < 18; fi30++ {
            if (925 < 257) {
                var v31 uint16 = 16
                _ = v31
            }
            v32 := []int{(fi29 ^ 467), (fi29 ^ (fi30 ^ fi30)), (fi29 << ((2) & 63)), fi29}
            v33 := GenS4{F0: 180, F1: 9223372036854775807}
            _ = fi30
            _ = v32
            _ = v33
        }
        v34 := GenS1{F0: false, F1: 0, F2: (fi29 & (fi29 & fi29))}
        fmt.Println(v3)
        _ = fi29
        _ = v34
    }
    switch "\n\t" {
    case "hello":
        for fi35 := 0; fi35 < 13; fi35++ {
            for fi36 := 0; fi36 < 1; fi36++ {
                if (!false) {
                    fmt.Println(v1, v1, fi36)
                    v3 += 9223372036854775807
                    v2[fi36] = v1
                    v3 = 0
                    var v37 bool = true
                    _ = v37
                }
                _ = fi36
            }
            _ = fi35
        }
    default:
        fmt.Println(v1, v3)
    }
    var v38 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v38
}
