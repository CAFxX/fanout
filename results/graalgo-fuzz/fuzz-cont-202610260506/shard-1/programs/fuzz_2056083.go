package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int32; F1 int16; F2 uint16 }
type GenS3 struct { F0 int16; F1 string }

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = -1
    v3 := []uint32{2147483647, 4294967295, 65535}
    var v4 int64 = 1
    var v5 string = ("" + "\n\t")
    var v6 uint8 = 182
    switch 1 {
    case -9223372036854775808:
        for fi7 := 0; fi7 < 1; fi7++ {
            if (!((fi7 + fi7) != fi7)) {
                for fi8 := 0; fi8 < 8; fi8++ {
                    v2 -= (v2 * 32767)
                    var v9 uint8 = (^(v6))
                    var v10 float64 = 1e10
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                var v11 uint64 = 1
                _ = v11
            } else {
                for fi12 := 0; fi12 < 1; fi12++ {
                    v5 = ((("hello" + "") + ("hello" + v5)) + (("a" + "") + v5))
                    _ = fi12
                }
            }
            _ = fi7
        }
    case 128:
        delete(v1, 32768)
    default:
        if (1.5 >= -2.25) {
            v5 = v5
            v5 = v5
            if false {
                v6 += (^(v6))
            }
        } else {
            if true {
                if (189 >= 255) {
                    if true {
                        v5 = "hello"
                        switch (("" + v5) + v5) {
                        case "hello":
                            v13 := GenS1{F0: 42, F1: 2, F2: v6}
                            _ = v13
                        case "\n\t":
                            for fi14 := 0; fi14 < 16; fi14++ {
                                v15 := GenS1{F0: -128, F1: v6, F2: (v6 ^ v6)}
                                _ = fi14
                                _ = v15
                            }
                        case "a":
                            if true {
                                v16 := GenS2{F0: 29, F1: (v2 / (((v2 >> ((65535) & 63))) | 1)), F2: 42}
                                var v17 uint8 = v6
                                _ = v16
                                _ = v17
                            }
                        }
                    }
                    fmt.Println(v2)
                }
                if true {
                    _ = len(v3)
                    for fi18 := 0; fi18 < 11; fi18++ {
                        var v19 uint32 = 280
                        _ = fi18
                        _ = v19
                    }
                } else {
                    var v20 float32 = (-2.25 / (1.5 / -2.25))
                    fmt.Println(v4, v4)
                    _ = v20
                }
            } else {
                fmt.Println(v4)
                _ = len(v1)
            }
            var v21 uint8 = (v6 - (^(v6)))
            for fi22 := 0; fi22 < 16; fi22++ {
                v6 += v21
                _ = fi22
            }
            _ = v21
        }
    }
    if (65535 < 1) {
        v5 = ""
        for fi23 := 0; fi23 < 8; fi23++ {
            if false {
                fmt.Println(v4)
                v2 += 1000
                delete(v1, (^(fi23)))
            }
            _ = len(v1)
            fmt.Println(v4, v5)
            _ = fi23
        }
    } else {
        for fi24 := 0; fi24 < 16; fi24++ {
            v6 = (^(v6))
            var v25 uint8 = 208
            var v26 uint16 = 1
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = len(v3)
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        var v28 uint64 = 4294967295
        _ = fi27
        _ = v28
    }
    var v29 int32 = 310
    if true {
        var v30 uint32 = 903
        var v31 float64 = ((0.0 + float64(1.5)) * float64(-2.25))
        _ = v30
        _ = v31
    }
    var v32 int8 = 127
    v32 += 42
    v4 = 532
    fmt.Println(v29, v2, v2)
    var v33 int64 = 267
    var v34 uint16 = 32767
    _ = len(v1)
    switch -9223372036854775808 {
    case 803:
        if (false && (uint32(v6) != 386)) {
            for fi35 := 0; fi35 < 5; fi35++ {
                for fi36 := 0; fi36 < 13; fi36++ {
                    var v37 float64 = float64(0.1)
                    _ = fi36
                    _ = v37
                }
                _ = fi35
            }
            v38 := GenS1{F0: v32, F1: (v6 << ((865) & 63)), F2: (^(v6))}
            fmt.Println(v6)
            _ = v38
        } else {
            if ((0.0 == 3.14159) || true) {
                var v39 int16 = -584
                v40 := v1[4294967295]
                _ = v39
                _ = v40
            }
            switch v5 {
            case "hello":
                fmt.Println(v4, v34, v2)
            }
            var v41 int8 = v32
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 16; fi42++ {
        var v43 int64 = v4
        _ = fi42
        _ = v43
    }
    v3 = append(v3, 380)
    v44 := GenS3{F0: v2, F1: "foo bar"}
    if true {
        var v45 int32 = v29
        v46 := GenS2{F0: (v45 + v29), F1: v2, F2: (v34 << ((2) & 63))}
        v47 := v1[233]
        _ = v45
        _ = v46
        _ = v47
    } else {
        v29 += int32(v34)
        v4 = (v4 << ((1) & 63))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
