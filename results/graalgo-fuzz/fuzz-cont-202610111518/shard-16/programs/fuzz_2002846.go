package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint8; F2 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 int8 }
func genFn0(p0 int32, p1 uint32, p2 string) int8 {
    p1 -= 65535
    if (true || (p1 != 1)) {
        fmt.Println(p0)
        for fi7 := 0; fi7 < 17; fi7++ {
            p1 = (^(p1))
            for fi8 := 0; fi8 < 10; fi8++ {
                var v9 int64 = 686
                if (866 <= 2) {
                    v9 -= v9
                }
                fmt.Println(v9, p2, p0)
                v10 := []int{(fi8 | (fi8 >> ((669) & 63))), fi8, fi8, (fi8 + fi7)}
                _ = fi8
                _ = v10
                _ = v9
            }
            switch (fi7 << ((1) & 63)) {
            case -9223372036854775808:
                v11 := GenS1{F0: true, F1: 2}
                _ = v11
            case 4294967295:
                var v12 uint8 = 0
                _ = v12
            }
            _ = fi7
        }
    }
    v13 := GenS1{F0: (18446744073709551615 <= 2), F1: 0}
    v14 := GenS1{F0: true, F1: 0}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 1
}
func genFn1(p0 uint8) uint32 {
    v15 := []int{}
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 uint16 = 908
        if ((false || (!true)) && (false && true)) {
            var v18 uint64 = 42
            v19 := GenS1{F0: false, F1: (^(p0))}
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 18; fi20++ {
                fmt.Println(p0, fi20, p0)
                _ = fi20
            }
            var v21 int32 = 261
            _ = v21
        }
        _ = fi16
        _ = v17
    }
    p0 = p0
    _ = p0
    _ = v15
    return 954
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint32 = 65535
    var v3 uint = 361
    var v4 int32 = -1
    var v5 int64 = 0
    var v6 uint64 = 42
    var v22 int32 = (v4 / ((-767) | 1))
    v23 := GenS2{F0: -333}
    if true {
        v24 := GenS3{F0: "a", F1: 108, F2: 0}
        v3 = v3
        _ = v24
    } else {
        v25 := []int{9223372036854775807}
        for fi26 := 0; fi26 < 13; fi26++ {
            v27 := GenS4{F0: -2147483648}
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    v28 := GenS4{F0: -2147483648}
    switch 32767 {
    case -2147483648:
        v3 = (^(v3))
    case -158:
        v29 := []int{}
        _ = v29
    case 9223372036854775807:
        for fi30 := 0; fi30 < 2; fi30++ {
            v22 -= v22
            v31 := GenS1{F0: ((^(v2)) < v2), F1: 2}
            _ = fi30
            _ = v31
        }
    }
    v32 := []int{0}
    v5 -= v5
    if (((((0.0 + (float32(1.5) + 1.5)) * 0.0) / (0.0 + (1e10 - -2.25))) <= (-2.25 / 3.14159)) && (float32(1e10) != 1.5)) {
        switch 273 {
        case 1:
            v2 += 792
        case 2147483647:
            for fi33 := 0; fi33 < 4; fi33++ {
                v1 = 18446744073709551615
                if len(v32) >= 2 {
                    v32 = v32[1:len(v32)]
                }
                v32 = append(v32, fi33)
                _ = fi33
            }
        case 127:
            v34 := GenS4{F0: (v22 - 42)}
            _ = v34
        default:
            v35 := GenS3{F0: "", F1: 127, F2: 0}
            _ = v35
        }
        v4 = v4
        for fi36 := 0; fi36 < 12; fi36++ {
            switch "a" {
            case "":
                v4 = v4
            case "foo bar":
                v37 := GenS5{F0: (v2 ^ (v2 * v2)), F1: 0}
                _ = v37
            default:
                v32 = append(v32, fi36)
            }
            if true {
                if len(v32) > 0 {
                    v32[0] = fi36
                }
                var v38 float32 = 0.0
                v5 = (-(v5))
                _ = v38
            } else {
                v2 = v2
            }
            var v39 int8 = 0
            _ = fi36
            _ = v39
        }
    } else {
        v4 -= (^(v4))
    }
    v40 := GenS3{F0: "\n\t", F1: 127, F2: 101}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
