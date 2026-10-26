package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 bool; F2 int64 }
type GenS5 struct { F0 int8; F1 int64 }
func genFn0() uint32 {
    v4 := GenS1{F0: 42, F1: -2147483648, F2: 886}
    if (true && true) {
        for fi5 := 0; fi5 < 9; fi5++ {
            fmt.Println(fi5)
            _ = fi5
        }
        var v6 uint64 = 42
        _ = v6
    }
    var v7 int16 = -217
    var v8 int64 = -9223372036854775808
    for fi9 := 0; fi9 < 16; fi9++ {
        switch "" {
        case "":
            var v10 string = "\n\t"
            _ = v10
        case "foo bar":
            var v11 uint16 = 0
            _ = v11
        }
        _ = fi9
    }
    _ = v4
    _ = v7
    _ = v8
    return 2147483648
}
func genFn1() int16 {
    var v12 int16 = 223
    switch 0 {
    case 0:
        for fi13 := 0; fi13 < 7; fi13++ {
            switch "hello" {
            case "a":
                v12 = v12
            default:
                for fi14 := 0; fi14 < 18; fi14++ {
                    var v15 int32 = -530
                    v12 = (v12 ^ 42)
                    v12 = (-(v12))
                    v12 -= (v12 * 32767)
                    _ = fi14
                    _ = v15
                }
            }
            _ = fi13
        }
    default:
        v16 := GenS2{F0: (v12 * v12)}
        _ = v16
    }
    switch 800 {
    case 127:
        var v17 float32 = -2.25
        _ = v17
    case 128:
        switch "" {
        case "a":
            v12 -= v12
        case "":
            var v18 uint32 = genFn0()
            _ = v18
        }
    case -2147483648:
        fmt.Println(v12)
    default:
        v19 := []int{245, 42, 63}
        _ = v19
    }
    _ = v12
    return 1
}

func main() {
    var v1 int32 = -311
    var v2 bool = true
    var v3 uint = 4294967295
    v20 := []int{624, -737}
    v1 = (v1 / (((v1 * (v1 & v1))) | 1))
    var v21 int32 = (^(v1))
    if true {
        if v2 {
            v3 = 18446744073709551615
            _ = len(v20)
        } else {
            v22 := GenS3{F0: v1}
            var v23 uint32 = 2
            var v24 uint64 = 280
            _ = v22
            _ = v23
            _ = v24
        }
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        var v25 int = -9223372036854775808
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        _ = len(v20)
        v21 = v21
        v27 := GenS3{F0: -814}
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        v3 = v3
        v1 = (v21 + (v1 - v21))
        for fi29 := 0; fi29 < 5; fi29++ {
            _ = len(v20)
            v2 = (!(-128 > -1))
            if (0 <= 4294967295) {
                v20 = append(v20, 2147483647)
                v30 := GenS1{F0: 4294967295, F1: v1, F2: 0}
                v21 = int32(fi28)
                _ = v30
            }
            var v31 int = 32767
            _ = fi29
            _ = v31
        }
        _ = fi28
    }
    if (!(0.1 >= float64(0.1))) {
        v3 = (v3 << ((397) & 63))
        _ = len(v20)
    }
    v3 = (v3 / (((v3 + (v3 % (((v3 % ((4294967295) | 1))) | 1)))) | 1))
    var v32 bool = false
    v1 = 1
    fmt.Println(v2)
    v33 := GenS2{F0: 846}
    for fi34 := 0; fi34 < 20; fi34++ {
        v35 := GenS4{F0: -1, F1: (((0.1 * -2.25) == 0.0) || true), F2: -517}
        v3 += (v3 % ((301) | 1))
        fmt.Println(fi34, v1, v2)
        _ = fi34
        _ = v35
    }
    v1 = (v21 / (((^(v1))) | 1))
    if (-9223372036854775808 >= -830) {
        for fi36 := 0; fi36 < 20; fi36++ {
            v21 = (v1 % ((v21) | 1))
            var v37 uint16 = 32767
            for fi38 := 0; fi38 < 5; fi38++ {
                v39 := GenS5{F0: 0, F1: 42}
                v40 := GenS1{F0: 18446744073709551615, F1: (v21 ^ (-(v21))), F2: 2147483648}
                if true {
                    switch fi36 {
                    case -508:
                        if len(v20) > 0 {
                            v20[0] = fi38
                        }
                    case 2147483647:
                        v3 = v3
                    case -631:
                        var v41 int = fi36
                        _ = v41
                    default:
                        v1 += v1
                    }
                    v37 = (v37 | v37)
                }
                _ = fi38
                _ = v39
                _ = v40
            }
            _ = fi36
            _ = v37
        }
        v42 := GenS4{F0: v1, F1: true, F2: 0}
        var v43 string = ""
        _ = v42
        _ = v43
    }
    v21 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v33
}
