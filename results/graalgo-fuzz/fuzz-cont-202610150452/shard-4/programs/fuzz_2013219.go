package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 bool; F2 int64 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]int)
    var v3 int64 = -9223372036854775808
    switch 32767 {
    case 127:
        fmt.Println(v3, v3)
    default:
        fmt.Println(v3)
    }
    switch ("hello" + (("\n\t" + "a") + ("hello" + "hello"))) {
    case "a":
        v1 += 4294967295
    case "\n\t":
        for fi4 := 0; fi4 < 1; fi4++ {
            v5 := GenS1{F0: 264, F1: (v3 | v3), F2: 1}
            if false {
                v1 = (^(v1))
            } else {
                switch "a" {
                case "hello":
                    var v6 uint64 = uint64(v1)
                    _ = v6
                case "a":
                    for fi7 := 0; fi7 < 9; fi7++ {
                        fmt.Println(fi7, v3, v3)
                        var v8 int8 = -128
                        switch (^(fi4)) {
                        case 4294967295:
                            switch "\n\t" {
                            case "hello":
                                var v9 int16 = -32768
                                _ = v9
                            }
                        case 0:
                            v3 -= (v3 ^ 0)
                            if (!(65535 != 2)) {
                                switch (("a" + "\n\t") + ("\n\t" + ("" + "foo bar"))) {
                                case "hello":
                                    var v10 uint16 = 2
                                    _ = v10
                                case "":
                                    var v11 uint64 = 42
                                    _ = v11
                                default:
                                    fmt.Println(fi7)
                                }
                                var v12 uint8 = 127
                                switch "foo bar" {
                                case "hello":
                                    fmt.Println(v8, v12)
                                    var v13 bool = true
                                    _ = v13
                                case "":
                                    var v14 int64 = v3
                                    _ = v14
                                case "foo bar":
                                    if (961 != -698) {
                                        v15 := []int{fi4, 1}
                                        if (v3 > v3) {
                                            v8 += v8
                                            v16 := GenS2{F0: ("\n\t" + "hello")}
                                            v8 = (v8 * (v8 >> ((2) & 63)))
                                            _ = v16
                                        }
                                        fmt.Println(v3)
                                        _ = v15
                                    }
                                default:
                                    v8 = (v8 + v8)
                                }
                                _ = v12
                            }
                            var v17 uint8 = 42
                            _ = v17
                        case 32767:
                            delete(v2, 460)
                            v18 := GenS3{F0: v3}
                            _ = v18
                        default:
                            v19 := GenS4{F0: 128, F1: (v1 < (v1 | v1)), F2: (-(v3))}
                            _ = v19
                        }
                        v20 := []int{(fi7 << ((0) & 63)), fi7, (fi7 << ((0) & 63))}
                        _ = fi7
                        _ = v20
                        _ = v8
                    }
                case "foo bar":
                    v1 -= 2
                default:
                    for fi21 := 0; fi21 < 4; fi21++ {
                        v3 = 9223372036854775807
                        _ = fi21
                    }
                }
                var v22 float32 = float32(0.1)
                _ = v22
            }
            _ = fi4
            _ = v5
        }
    default:
        if false {
            switch "hello" {
            case "a":
                for fi23 := 0; fi23 < 9; fi23++ {
                    var v24 int = (fi23 * 42)
                    v25 := GenS5{F0: 32767}
                    switch "a" {
                    case "foo bar":
                        v3 += v3
                    case "":
                        v24 = 42
                    default:
                        v1 += 4294967295
                    }
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        }
    }
    switch 32768 {
    case 2147483647:
        v26 := GenS5{F0: 65535}
        _ = v26
    case 280:
        var v27 uint16 = 132
        _ = v27
    default:
        v3 = v3
    }
    if (((v3 <= v3) && false) && (!true)) {
        v28 := GenS2{F0: ("a" + "foo bar")}
        var v29 int16 = 1
        fmt.Println(v29, v1, v29)
        _ = v28
        _ = v29
    }
    if (((-2.25 / (1e10 + -2.25)) + -2.25) < 0.1) {
        var v30 int64 = v3
        var v31 float64 = float64((0.1 + float32(1e10)))
        v2[(v1 << ((556) & 63))] = -9223372036854775808
        _ = v30
        _ = v31
    }
    var v32 uint64 = 836
    if (1000 < -828) {
        for fi33 := 0; fi33 < 13; fi33++ {
            v34 := v2[(v1 - (v1 ^ v1))]
            _ = fi33
            _ = v34
        }
    }
    v35 := []int{2147483647}
    v36 := GenS5{F0: 32767}
    if ((("foo bar" + "foo bar") + "a") >= ("foo bar" + "a")) {
        v3 += (v3 << ((255) & 63))
    }
    v37 := GenS5{F0: 65535}
    if (100 <= 42) {
        v3 -= int64(v1)
        v3 -= v3
        v2[v1] = 4294967295
    } else {
        var v38 float32 = ((float32(0.0) - (((1e10 - 0.0) - (-2.25 - 1e10)) * 1.5)) + ((3.14159 / ((0.1 + 0.0) * 0.1)) * (1e10 / 3.14159)))
        fmt.Println(v1, v1)
        _ = v38
    }
    delete(v2, 216)
    _ = len(v2)
    for fi39 := 0; fi39 < 5; fi39++ {
        v32 = (^(v32))
        fmt.Println(v32, v3)
        v3 -= v3
        _ = fi39
    }
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v35:", len(v35), v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v37
}
