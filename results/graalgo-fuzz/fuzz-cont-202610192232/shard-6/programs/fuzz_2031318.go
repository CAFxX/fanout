package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 string }
type GenS2 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]int)
    var v3 int8 = -1
    var v4 bool = true
    var v5 int64 = 0
    v1 -= v1
    v6 := GenS1{F0: (^(v1)), F1: 191, F2: "foo bar"}
    v4 = true
    delete(v2, 2147483648)
    v7 := GenS1{F0: 18446744073709551615, F1: 0, F2: ("\n\t" + "hello")}
    var v8 uint16 = 681
    if (false && v4) {
        if (!(668 >= 32767)) {
            for fi9 := 0; fi9 < 12; fi9++ {
                v4 = (!v4)
                for fi10 := 0; fi10 < 1; fi10++ {
                    switch fi9 {
                    case 9223372036854775807:
                        var v11 string = "foo bar"
                        _ = v11
                    case 771:
                        if (235 < 0) {
                            v12 := GenS1{F0: 0, F1: (v8 - 255), F2: "\n\t"}
                            _ = v12
                        }
                        v13 := GenS1{F0: (v1 & 18446744073709551615), F1: 65535, F2: ("hello" + "foo bar")}
                        _ = v13
                    case -1:
                        v14 := []int{(fi9 * fi10), 32768, (^(fi9))}
                        _ = v14
                    default:
                        fmt.Println(v5, fi10)
                    }
                    _ = fi10
                }
                if (!false) {
                    v15 := GenS1{F0: (^(v1)), F1: (^(v8)), F2: "foo bar"}
                    var v16 uint64 = (v1 ^ v1)
                    _ = v15
                    _ = v16
                } else {
                    if v4 {
                        fmt.Println(v4, v8)
                        fmt.Println(v3)
                    }
                    var v17 uint64 = v1
                    _ = v17
                }
                _ = fi9
            }
            switch -873 {
            case -1:
                v18 := GenS2{F0: 42, F1: v1, F2: v8}
                _ = v18
            case 9223372036854775807:
                v19 := GenS1{F0: (v1 - (v1 & 460)), F1: (v8 << ((152) & 63)), F2: ("hello" + "hello")}
                _ = v19
            case -2147483648:
                v20 := GenS2{F0: v3, F1: (^(v1)), F2: v8}
                _ = v20
            default:
                v21 := GenS3{F0: 255}
                _ = v21
            }
            v1 -= 2
        }
        var v22 uint32 = uint32(v1)
        _ = v22
    } else {
        v3 += v3
        var v23 float32 = (((3.14159 * -2.25) * 1.5) * 3.14159)
        v8 -= uint16(v1)
        _ = v23
    }
    v24 := []int{9223372036854775807, 772, int(v8)}
    var v25 uint = 2
    v25 += v25
    if v4 {
        if len(v24) > 0 {
            v24[0] = -1
        }
    }
    fmt.Println(v4, v5, v4)
    for fi26 := 0; fi26 < 5; fi26++ {
        v5 = (v5 & -642)
        _ = fi26
    }
    switch -9223372036854775808 {
    case 127:
        v27 := GenS2{F0: (v3 - -1), F1: (v1 & (v1 % (((v1 % (((^(v1))) | 1))) | 1))), F2: (v8 * (v8 >> ((v25) & 63)))}
        _ = v27
    case -9223372036854775808:
        var v28 int = -9223372036854775808
        _ = v28
    case 646:
        v25 += (^(v25))
    default:
        if (!(0 != 0)) {
            v3 -= v3
            for fi29 := 0; fi29 < 8; fi29++ {
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
                fmt.Println(v4)
                var v30 uint64 = 1
                _ = fi29
                _ = v30
            }
            v3 = v3
        } else {
            for fi31 := 0; fi31 < 13; fi31++ {
                fmt.Println(v1)
                _ = fi31
            }
            _ = len(v24)
            v8 = 65535
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
