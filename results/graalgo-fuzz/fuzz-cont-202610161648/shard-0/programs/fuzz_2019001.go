package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint64; F2 int8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 4294967295
    var v3 float64 = (float64(0.0) / 0.1)
    v4 := make(map[uint32]string)
    v5 := make(map[int]bool)
    v6 := []uint32{}
    v7 := GenS1{F0: -14, F1: 42}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int = (fi8 ^ -9223372036854775808)
        _ = fi8
        _ = v9
    }
    if (!false) {
        v10 := GenS1{F0: 552, F1: -686}
        var v11 float64 = v3
        var v12 float64 = (v3 * (v3 / v3))
        _ = v10
        _ = v11
        _ = v12
    }
    v3 += 1.5
    var v13 float32 = 0.0
    for fi14 := 0; fi14 < 1; fi14++ {
        if false {
            switch -813 {
            case 127:
                if (2 >= 18446744073709551615) {
                    var v15 float32 = float32(v3)
                    switch "a" {
                    case "a":
                        v15 -= v15
                    case "":
                        v4[65535] = "a"
                    case "foo bar":
                        fmt.Println(v2, v13, v2)
                    }
                    switch fi14 {
                    case 128:
                        v15 = 0.1
                    default:
                        var v16 int8 = 1
                        _ = v16
                    }
                    _ = v15
                } else {
                    v13 -= 1e10
                    v2 += 0
                    for fi17 := 0; fi17 < 2; fi17++ {
                        v2 -= v2
                        _ = fi17
                    }
                }
            case 786:
                for fi18 := 0; fi18 < 16; fi18++ {
                    if (v2 != (v2 + (v2 / ((0) | 1)))) {
                        delete(v1, "hello")
                        var v19 int8 = 100
                        _ = v19
                    } else {
                        v3 = 0.0
                        v20 := GenS1{F0: (fi18 / ((fi18) | 1)), F1: 1}
                        _ = v20
                    }
                    switch (fi14 & 70) {
                    case 32768:
                        var v21 int64 = -670
                        _ = v21
                    case 320:
                        if true {
                            if (-1 > 0) {
                                v2 = 4294967295
                                var v22 float64 = 0.0
                                _ = v22
                            }
                        }
                    case 2147483647:
                        v23 := GenS2{F0: 1, F1: 0}
                        _ = v23
                    }
                    _ = fi18
                }
            }
            v24 := GenS3{F0: ("foo bar" + "a")}
            _ = v24
        }
        _ = fi14
    }
    if (428 <= 0) {
        for fi25 := 0; fi25 < 17; fi25++ {
            v3 += v3
            for fi26 := 0; fi26 < 8; fi26++ {
                _ = len(v1)
                var v27 uint32 = 110
                v3 = (3.14159 + v3)
                _ = fi26
                _ = v27
            }
            switch fi25 {
            case 9223372036854775807:
                v13 -= 1.5
            case 32767:
                var v28 float64 = (1e10 + float64(v13))
                _ = v28
            case 2147483647:
                fmt.Println(fi25, fi25, v13)
            default:
                if true {
                    if false {
                        for fi29 := 0; fi29 < 5; fi29++ {
                            v3 -= 1e10
                            var v30 int64 = 0
                            if false {
                                v31 := GenS4{F0: true, F1: 18446744073709551615, F2: -128}
                                _ = v31
                            }
                            _ = fi29
                            _ = v30
                        }
                        v32 := v4[65535]
                        if false {
                            for fi33 := 0; fi33 < 13; fi33++ {
                                v32 = (((v32 + v32) + v32) + v32)
                                _ = fi33
                            }
                        }
                        _ = v32
                    }
                    for fi34 := 0; fi34 < 15; fi34++ {
                        fmt.Println(v2, v3)
                        _ = fi34
                    }
                    var v35 uint64 = 659
                    _ = v35
                }
            }
            _ = fi25
        }
        var v36 float32 = 0.0
        _ = v36
    }
    if false {
        _ = len(v1)
    }
    var v37 int32 = -1
    v37 = (-(v37))
    var v38 int32 = -2147483648
    if (97 == 696) {
        v3 = float64((v13 * v13))
        for fi39 := 0; fi39 < 12; fi39++ {
            var v40 string = "\n\t"
            _ = fi39
            _ = v40
        }
    } else {
        v3 += 3.14159
        fmt.Println(v2)
        var v41 uint64 = (v2 ^ (v2 / ((v2) | 1)))
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
