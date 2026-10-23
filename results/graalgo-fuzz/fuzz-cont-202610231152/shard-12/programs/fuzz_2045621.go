package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int16; F2 int8 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int16; F1 string; F2 uint16 }
func genFn0(p0 uint64) uint16 {
    if false {
        p0 = p0
    }
    for fi4 := 0; fi4 < 4; fi4++ {
        p0 += p0
        v5 := []int{(fi4 ^ fi4), (^(fi4)), 128}
        if len(v5) > 0 {
            v5[0] = fi4
        }
        _ = fi4
        _ = v5
    }
    switch -438 {
    case 32768:
        v6 := []int{}
        _ = v6
    case 9223372036854775807:
        if (-1 > 2147483647) {
            switch 127 {
            case 1:
                v7 := GenS1{F0: 453, F1: 1}
                _ = v7
            case -2147483648:
                if (!true) {
                    p0 = 273
                }
            }
            v8 := []int{9223372036854775807}
            _ = v8
        }
    case 4294967295:
        switch "" {
        case "":
            switch -939 {
            case 32768:
                var v9 uint32 = 0
                _ = v9
            case 4294967295:
                for fi10 := 0; fi10 < 8; fi10++ {
                    fmt.Println(fi10, fi10)
                    v11 := GenS2{F0: 0}
                    _ = fi10
                    _ = v11
                }
            case 832:
                var v12 int16 = -1
                _ = v12
            }
        default:
            p0 += 42
        }
    default:
        v13 := GenS1{F0: 584, F1: 128}
        _ = v13
    }
    var v14 string = "\n\t"
    for fi15 := 0; fi15 < 20; fi15++ {
        v14 = v14
        _ = fi15
    }
    _ = p0
    _ = v14
    return 0
}

func main() {
    v1 := []int{330}
    var v2 int64 = 184
    var v3 int16 = 1
    if false {
        v3 += v3
        for fi16 := 0; fi16 < 18; fi16++ {
            switch "\n\t" {
            case "hello":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                fmt.Println(fi16, fi16)
            }
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 int32 = 181
                var v19 uint8 = 126
                _ = fi17
                _ = v18
                _ = v19
            }
            fmt.Println(v3, fi16, fi16)
            _ = fi16
        }
        var v20 int64 = v2
        _ = v20
    } else {
        v3 = (v3 >> ((18446744073709551615) & 63))
    }
    var v21 uint16 = 255
    fmt.Println(v21, v3)
    v22 := GenS3{F0: int32(v3), F1: v3, F2: 42}
    var v23 int8 = -128
    switch 1 {
    case -619:
        if true {
            v23 = v23
            v21 -= v21
        }
    case 4294967295:
        var v24 int32 = 1
        _ = v24
    default:
        v25 := GenS2{F0: v23}
        _ = v25
    }
    fmt.Println(v23)
    if false {
        for fi26 := 0; fi26 < 6; fi26++ {
            switch 236 {
            case 128:
                if len(v1) > 0 {
                    v1[0] = fi26
                }
            }
            var v27 uint8 = 42
            _ = fi26
            _ = v27
        }
        if len(v1) > 0 {
            v1[0] = 960
        }
    }
    switch 9223372036854775807 {
    case 2147483647:
        v2 = (v2 & (v2 / ((v2) | 1)))
    case -2147483648:
        if true {
            var v28 string = ("\n\t" + ("foo bar" + "\n\t"))
            _ = v28
        }
    default:
        v23 += -17
    }
    var v29 uint = 4294967295
    switch "hello" {
    case "":
        v3 -= v3
    default:
        if len(v1) > 0 {
            v1[0] = int(v29)
        }
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        if ("\n\t" < "hello") {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        if (1e10 > 3.14159) {
            v1 = append(v1, fi30)
        } else {
            var v31 uint = v29
            switch "hello" {
            case "hello":
                if (v29 == (^(v31))) {
                    for fi32 := 0; fi32 < 9; fi32++ {
                        switch "" {
                        case "\n\t":
                            var v33 uint = v31
                            _ = v33
                        }
                        v34 := GenS3{F0: 535, F1: (v3 << ((1) & 63)), F2: (v23 >> (((v31 >> (((^(v29))) & 63))) & 63))}
                        _ = fi32
                        _ = v34
                    }
                }
            default:
                for fi35 := 0; fi35 < 11; fi35++ {
                    v31 -= 426
                    v1 = append(v1, -1)
                    _ = fi35
                }
            }
            v36 := GenS4{F0: 18446744073709551615, F1: 978}
            _ = v31
            _ = v36
        }
        switch "" {
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "hello":
            v23 = v23
        case "\n\t":
            v3 = 0
        }
        _ = fi30
    }
    v37 := GenS5{F0: (v3 & (^(v3))), F1: "", F2: (v21 % (((v21 / (((v21 << ((v29) & 63))) | 1))) | 1))}
    fmt.Println(v3, v3)
    if false {
        if ((v21 | (^(v21))) != 0) {
            fmt.Println(v23, v29)
            v21 = genFn0(1)
        } else {
            for fi38 := 0; fi38 < 6; fi38++ {
                _ = len(v1)
                _ = fi38
            }
            fmt.Println(v2, v29, v2)
        }
        var v39 int8 = (v23 << ((v29) & 63))
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v37
}
