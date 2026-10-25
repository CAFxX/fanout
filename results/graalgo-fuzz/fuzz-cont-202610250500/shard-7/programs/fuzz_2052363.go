package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int64 }
type GenS2 struct { F0 int; F1 int32; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 bool; F1 uint16 }

func main() {
    v1 := []uint32{1, 0, 2147483647}
    var v2 uint32 = 42
    var v3 int = -2147483648
    v4 := GenS1{F0: -9223372036854775808, F1: v3, F2: -975}
    var v5 bool = (false && (true || true))
    v5 = v5
    fmt.Println(v2, v5)
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 float64 = -2.25
        _ = fi6
        _ = v7
    }
    _ = len(v1)
    if false {
        var v8 uint64 = 1
        switch v3 {
        case 2147483647:
            v8 -= v8
        case 149:
            v9 := GenS2{F0: v3, F1: -992, F2: int8(v2)}
            _ = v9
        case -439:
            switch (v3 + (^(v3))) {
            case 32767:
                if v5 {
                    v10 := GenS1{F0: (v3 - (-(v3))), F1: v3, F2: -466}
                    if v5 {
                        v2 += v2
                        v8 += 42
                    }
                    var v11 uint64 = v8
                    _ = v10
                    _ = v11
                }
            case 0:
                if v5 {
                    switch v3 {
                    case -160:
                        v3 = -1
                    case -901:
                        for fi12 := 0; fi12 < 20; fi12++ {
                            var v13 int64 = -848
                            if v5 {
                                if (v8 <= (v8 | (v8 / ((v8) | 1)))) {
                                    _ = len(v1)
                                    v14 := GenS1{F0: v3, F1: 32768, F2: v13}
                                    _ = v14
                                }
                            } else {
                                v15 := GenS1{F0: (^(fi12)), F1: fi12, F2: 21}
                                _ = v15
                            }
                            _ = fi12
                            _ = v13
                        }
                    default:
                        var v16 int8 = 42
                        _ = v16
                    }
                } else {
                    for fi17 := 0; fi17 < 16; fi17++ {
                        if v5 {
                            v18 := GenS2{F0: 32768, F1: 0, F2: 100}
                            if v5 {
                                fmt.Println(v2, v8, v2)
                                var v19 string = ""
                                _ = v19
                            } else {
                                v8 = 2
                                for fi20 := 0; fi20 < 13; fi20++ {
                                    v21 := GenS3{F0: (^(v8)), F1: (^(v2)), F2: 2147483647}
                                    _ = fi20
                                    _ = v21
                                }
                            }
                            var v22 bool = v5
                            _ = v18
                            _ = v22
                        }
                        var v23 uint = 272
                        _ = fi17
                        _ = v23
                    }
                }
            case 4294967295:
                for fi24 := 0; fi24 < 2; fi24++ {
                    v25 := GenS1{F0: v3, F1: (v3 % (((v3 << ((0) & 63))) | 1)), F2: 9223372036854775807}
                    v2 -= (v2 ^ v2)
                    _ = fi24
                    _ = v25
                }
            default:
                switch "" {
                case "\n\t":
                    for fi26 := 0; fi26 < 2; fi26++ {
                        fmt.Println(fi26)
                        switch (fi26 & v3) {
                        case 42:
                            var v27 float64 = (0.1 * 1e10)
                            _ = v27
                        case 9223372036854775807:
                            v28 := GenS4{F0: 145, F1: (!v5)}
                            _ = v28
                        case -555:
                            if (-836 >= -620) {
                                var v29 uint32 = (v2 >> ((65535) & 63))
                                for fi30 := 0; fi30 < 17; fi30++ {
                                    fmt.Println(v3, fi26, v2)
                                    v3 = fi26
                                    v2 += (^(v2))
                                    _ = fi30
                                }
                                _ = v29
                            } else {
                                switch "a" {
                                case "":
                                    var v31 uint = 29
                                    _ = v31
                                default:
                                    v1 = append(v1, 42)
                                }
                                v32 := GenS5{F0: v5, F1: 830}
                                if v5 {
                                    v8 = 18446744073709551615
                                } else {
                                    if (uint8(v8) >= 48) {
                                        v1 = append(v1, v2)
                                        switch ("" + ("" + "foo bar")) {
                                        case "":
                                            fmt.Println(v8, v5)
                                        case "foo bar":
                                            switch "foo bar" {
                                            case "a":
                                                var v33 uint16 = 775
                                                _ = v33
                                            case "\n\t":
                                                switch ("\n\t" + "a") {
                                                case "":
                                                    _ = len(v1)
                                                }
                                            default:
                                                for fi34 := 0; fi34 < 12; fi34++ {
                                                    v2 = (v2 ^ 2147483648)
                                                    _ = fi34
                                                }
                                            }
                                        case "\n\t":
                                            if (0.1 > 3.14159) {
                                                var v35 int8 = 88
                                                _ = v35
                                            }
                                        default:
                                            fmt.Println(fi26, v2, v3)
                                        }
                                    } else {
                                        for fi36 := 0; fi36 < 14; fi36++ {
                                            v5 = v5
                                            v3 = -1
                                            var v37 int16 = 0
                                            v5 = false
                                            _ = fi36
                                            _ = v37
                                        }
                                    }
                                    if (1 == v2) {
                                        if (32767 == 42) {
                                            for fi38 := 0; fi38 < 10; fi38++ {
                                                var v39 int16 = 1
                                                _ = fi38
                                                _ = v39
                                            }
                                            v1 = append(v1, 468)
                                            var v40 uint8 = 188
                                            _ = v40
                                        }
                                        v41 := GenS2{F0: (fi26 - (v3 - 128)), F1: 2147483647, F2: 47}
                                        var v42 float64 = 0.0
                                        _ = v41
                                        _ = v42
                                    }
                                }
                                _ = v32
                            }
                        default:
                            fmt.Println(v2, v5)
                        }
                        if (true || true) {
                            switch 570 {
                            case 42:
                                if v5 {
                                    v5 = true
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = 2
                                    }
                                    _ = len(v1)
                                }
                            case 32768:
                                fmt.Println(v2, v5)
                            case 4294967295:
                                var v43 float32 = (-2.25 / 1e10)
                                _ = v43
                            }
                            if len(v1) > 0 {
                                v1[0] = v2
                            }
                        } else {
                            v8 = v8
                        }
                        _ = fi26
                    }
                }
            }
        default:
            fmt.Println(v2)
        }
        _ = v8
    }
    fmt.Println(v2, v5, v3)
    v44 := GenS5{F0: false, F1: 889}
    v5 = (int64(v2) <= 9223372036854775807)
    v3 -= (-(v3))
    for fi45 := 0; fi45 < 3; fi45++ {
        v46 := GenS1{F0: v3, F1: 128, F2: -9223372036854775808}
        _ = fi45
        _ = v46
    }
    v47 := GenS5{F0: (v5 || (v5 && true)), F1: 692}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v47
    _ = v5
}
