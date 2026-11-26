package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int32 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int32, p1 uint64) int16 {
    p1 -= p1
    for fi4 := 0; fi4 < 11; fi4++ {
        var v5 uint64 = (p1 | (p1 % ((p1) | 1)))
        _ = fi4
        _ = v5
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint16 = 789
    var v2 uint8 = 255
    v3 := []uint32{4294967295, 2147483647, 705}
    _ = len(v3)
    if false {
        if true {
            var v6 int16 = 0
            if len(v3) > 0 {
                v3[0] = uint32(v2)
            }
            _ = v6
        } else {
            var v7 uint8 = v2
            var v8 uint16 = v1
            if ("" != ("" + "a")) {
                var v9 uint32 = 443
                _ = v9
            }
            _ = v7
            _ = v8
        }
        v1 += (v1 / ((42) | 1))
    } else {
        var v10 uint16 = (v1 * (^(v1)))
        fmt.Println(v2, v1, v2)
        _ = v10
    }
    if (128 <= 42) {
        for fi11 := 0; fi11 < 13; fi11++ {
            v2 = v2
            _ = fi11
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 int32 = 0
            fmt.Println(v1)
            if false {
                v14 := GenS1{F0: 108, F1: 625, F2: v13}
                var v15 int16 = 0
                _ = v14
                _ = v15
            } else {
                var v16 int = fi12
                _ = v16
            }
            _ = fi12
            _ = v13
        }
    }
    switch ("hello" + "hello") {
    case "\n\t":
        if (((v2 / ((42) | 1)) <= v2) && (false || false)) {
            if false {
                for fi17 := 0; fi17 < 11; fi17++ {
                    v3 = append(v3, 2147483648)
                    fmt.Println(v2, v1, fi17)
                    _ = fi17
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    case "":
        var v18 float64 = 3.14159
        _ = v18
    case "hello":
        v19 := GenS1{F0: -27, F1: 0, F2: 624}
        _ = v19
    default:
        for fi20 := 0; fi20 < 19; fi20++ {
            if (0.1 <= -2.25) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                for fi21 := 0; fi21 < 11; fi21++ {
                    v1 = 255
                    v22 := GenS1{F0: 127, F1: (fi21 | 32767), F2: 1}
                    if ((971 > 1) || (false && true)) {
                        _ = len(v3)
                        var v23 float64 = 1.5
                        switch (fi21 + -2147483648) {
                        case -453:
                            switch (-(fi20)) {
                            case 32767:
                                v1 += v1
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        }
                        switch fi20 {
                        case 42:
                            v23 = v23
                        case 2147483647:
                            if (-78 == 1) {
                                var v24 string = ""
                                _ = v24
                            } else {
                                v25 := GenS2{F0: -623, F1: (v2 | v2)}
                                _ = v25
                            }
                        case 32767:
                            v26 := GenS1{F0: -32, F1: fi20, F2: 1}
                            _ = v26
                        default:
                            if (!true) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if true {
                                    v27 := GenS1{F0: 100, F1: fi20, F2: 1}
                                    if true {
                                        fmt.Println(fi21, v1, fi20)
                                    }
                                    _ = v27
                                }
                                var v28 uint64 = 4294967295
                                _ = v28
                            }
                        }
                        _ = v23
                    }
                    _ = fi21
                    _ = v22
                }
            } else {
                var v29 int64 = 702
                v1 = v1
                v2 -= v2
                _ = v29
            }
            var v30 bool = (v2 == (v2 + v2))
            v1 = v1
            _ = fi20
            _ = v30
        }
    }
    switch 0 {
    case -1:
        switch "a" {
        case "":
            fmt.Println(v2, v2)
        case "\n\t":
            v2 = v2
        }
    default:
        switch "\n\t" {
        case "hello":
            if (false && (-1 >= 0)) {
                v3 = append(v3, 162)
                fmt.Println(v2, v2)
            }
        }
    }
    v1 -= 42
    v2 = 201
    v1 = (v1 / (((v1 % (((v1 << ((255) & 63))) | 1))) | 1))
    if len(v3) > 0 {
        v3[0] = 2
    }
    _ = len(v3)
    var v31 bool = (-9223372036854775808 == -1)
    var v32 int32 = -752
    v33 := GenS3{F0: (v2 / (((v2 * uint8(v1))) | 1))}
    var v34 uint32 = 709
    v35 := GenS4{F0: int32(v2)}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
