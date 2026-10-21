package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 bool }

func main() {
    var v1 uint = 65535
    v2 := make(map[int]bool)
    var v3 uint8 = 46
    var v4 uint64 = 2
    var v5 float32 = (3.14159 - -2.25)
    var v6 int8 = 0
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 int64 = 0
            v2[fi8] = ((!(v4 > (^(v4)))) && false)
            if ((v4 ^ uint64(v3)) != v4) {
                fmt.Println(v5, fi8, fi7)
            }
            var v10 int = (^(fi7))
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    if false {
        var v11 int = -2147483648
        v12 := []int{(v11 >> (((v1 * (^(v1)))) & 63)), -2147483648, -913}
        v11 -= -9223372036854775808
        _ = v11
        _ = v12
    } else {
        if ((v4 > v4) && (65535 <= 2)) {
            v13 := GenS1{F0: 76}
            _ = v13
        }
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 bool = (float64(((v5 - 0.1) / v5)) < 0.1)
            if (false && v15) {
                v15 = true
            }
            v16 := []int{}
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    v17 := []int{-2147483648, -2147483648}
    v5 = 1e10
    if (true && (float64(((v5 - float32(-2.25)) + (1e10 * ((-2.25 - v5) - (float32(((0.0 - ((3.14159 - 0.1) * 0.1)) / 1.0)) / (v5 - ((v5 + float32(1.5)) / v5))))))) < ((1e10 / 0.1) / 0.1))) {
        var v18 uint16 = 1
        _ = v18
    } else {
        var v19 uint16 = 42
        fmt.Println(v6, v6, v3)
        _ = v19
    }
    var v20 uint32 = 669
    _ = len(v17)
    for fi21 := 0; fi21 < 14; fi21++ {
        switch "foo bar" {
        case "hello":
            var v22 int16 = -32768
            _ = v22
        default:
            if (!true) {
                v3 -= (^(v3))
                var v23 bool = true
                _ = v23
            } else {
                v20 = (^(v20))
                if false {
                    v20 = uint32(v4)
                    if false {
                        for fi24 := 0; fi24 < 8; fi24++ {
                            if true {
                                v3 -= v3
                                var v25 float64 = ((0.1 - ((3.14159 + 0.0) / 1.5)) / 0.1)
                                v26 := GenS1{F0: 856}
                                _ = v25
                                _ = v26
                            } else {
                                fmt.Println(fi21)
                                fmt.Println(fi24, v1, v3)
                                switch (fi21 >> ((v1) & 63)) {
                                case -99:
                                    var v27 float64 = (0.0 * 1.5)
                                    _ = v27
                                case 127:
                                    if (!false) {
                                        v28 := GenS2{F0: 0, F1: 65535, F2: v3}
                                        v29 := GenS3{F0: 886, F1: 1}
                                        _ = v28
                                        _ = v29
                                    }
                                }
                                fmt.Println(v20, fi24, v1)
                            }
                            _ = fi24
                        }
                        v30 := GenS3{F0: (v1 + (v1 % ((42) | 1))), F1: -1}
                        _ = v30
                    }
                    var v31 int8 = v6
                    _ = v31
                }
            }
        }
        switch fi21 {
        case -790:
            for fi32 := 0; fi32 < 15; fi32++ {
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                v33 := GenS4{F0: 32767}
                _ = fi32
                _ = v33
            }
        }
        _ = fi21
    }
    v34 := GenS3{F0: 2, F1: 829}
    for fi35 := 0; fi35 < 2; fi35++ {
        if (v6 < (v6 >> ((v1) & 63))) {
            v1 -= 18446744073709551615
            switch "a" {
            case "":
                v3 -= (v3 % (((^(v3))) | 1))
            case "hello":
                for fi36 := 0; fi36 < 13; fi36++ {
                    var v37 int32 = -2147483648
                    _ = fi36
                    _ = v37
                }
            default:
                v38 := v2[fi35]
                _ = v38
            }
            v39 := GenS5{F0: (v1 % ((664) | 1)), F1: (!((v3 | (v3 + 77)) <= v3))}
            _ = v39
        }
        for fi40 := 0; fi40 < 1; fi40++ {
            if (42 < 42) {
                v41 := GenS3{F0: v1, F1: 1000}
                _ = v41
            } else {
                var v42 int8 = 5
                delete(v2, (fi35 / (((fi35 % ((fi35) | 1))) | 1)))
                v4 = (v4 ^ v4)
                _ = v42
            }
            v43 := v2[(fi35 << (((^(v1))) & 63))]
            v4 = 42
            _ = fi40
            _ = v43
        }
        _ = fi35
    }
    var v44 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
