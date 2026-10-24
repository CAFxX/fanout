package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint8; F1 string; F2 int32 }
type GenS3 struct { F0 int; F1 bool; F2 string }
type GenS4 struct { F0 uint8; F1 uint32 }
func genFn0(p0 uint, p1 int8) int64 {
    var v4 int32 = 0
    if true {
        var v5 uint8 = 2
        fmt.Println(p0, p1, p0)
        _ = v5
    }
    _ = p0
    _ = p1
    _ = v4
    return 467
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 102
    var v3 bool = true
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 int = 128
            v9 := GenS1{F0: 4294967295, F1: 255, F2: 2}
            fmt.Println(v2, v3, v3)
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    _ = len(v1)
    if v3 {
        v10 := GenS1{F0: 42, F1: 2, F2: 856}
        _ = v10
    } else {
        if (true || true) {
            v11 := GenS2{F0: (v2 ^ (^(v2))), F1: "hello", F2: 42}
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            switch fi13 {
            case -2147483648:
                v3 = v3
            case -819:
                fmt.Println(v2, v3, v3)
            case 32767:
                if true {
                    fmt.Println(v3, fi13)
                    v1 = append(v1, 65535)
                    v1 = append(v1, 2147483647)
                }
            default:
                var v14 int = fi12
                _ = v14
            }
            switch "\n\t" {
            case "a":
                v1 = append(v1, 332)
            default:
                switch "a" {
                case "foo bar":
                    var v15 int64 = 721
                    _ = v15
                case "\n\t":
                    fmt.Println(fi13)
                case "hello":
                    fmt.Println(v2, v2, v2)
                }
            }
            if v3 {
                fmt.Println(fi12, v2)
            } else {
                fmt.Println(v2, fi13, fi13)
                v3 = (true || (!v3))
            }
            if v3 {
                fmt.Println(fi13)
                v2 = v2
                if true {
                    if (32767 <= 32767) {
                        if (986 > 18446744073709551615) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v16 int = (fi12 | -642)
                            v1 = append(v1, 1)
                            fmt.Println(fi12, fi12)
                            v3 = (754 > 1)
                            _ = v16
                        } else {
                            _ = len(v1)
                        }
                        var v17 int16 = int16(fi13)
                        _ = v17
                    }
                    if ((!((v2 & 128) >= (v2 / (((v2 - (v2 - v2))) | 1)))) || true) {
                        var v18 float64 = 0.0
                        if len(v1) > 0 {
                            v1[0] = 557
                        }
                        fmt.Println(fi12, fi12, v2)
                        _ = v18
                    } else {
                        switch "" {
                        case "hello":
                            v1 = append(v1, 2147483647)
                        case "":
                            v19 := GenS1{F0: 0, F1: 18446744073709551615, F2: 583}
                            _ = v19
                        case "a":
                            v2 = (v2 ^ v2)
                        default:
                            v20 := GenS3{F0: fi13, F1: true, F2: "\n\t"}
                            _ = v20
                        }
                    }
                    var v21 float32 = 0.0
                    v21 -= v21
                    _ = v21
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            v2 -= (v2 / ((1) | 1))
            _ = fi13
        }
        v22 := GenS4{F0: (v2 ^ (^(v2))), F1: 594}
        var v23 string = "a"
        _ = fi12
        _ = v22
        _ = v23
    }
    var v24 uint = 400
    _ = len(v1)
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 float32 = 3.14159
        _ = fi25
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
}
