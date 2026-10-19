package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int16; F1 int; F2 int16 }
type GenS4 struct { F0 uint32; F1 uint8; F2 bool }
type GenS5 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS6 struct { F0 bool; F1 uint; F2 string }
type GenS7 struct { F0 int8; F1 int; F2 uint8 }

func main() {
    var v1 int64 = 42
    var v2 string = ""
    var v3 int = 880
    v4 := make(map[string]bool)
    v5 := make(map[string]string)
    var v6 uint64 = 0
    v7 := []int{}
    switch v2 {
    case "hello":
        switch ("foo bar" + "") {
        case "foo bar":
            if (!(417 == 53)) {
                for fi8 := 0; fi8 < 6; fi8++ {
                    v9 := GenS1{F0: 32767, F1: 2147483648}
                    for fi10 := 0; fi10 < 4; fi10++ {
                        if true {
                            var v11 int16 = 1
                            delete(v5, ("" + (v2 + v2)))
                            switch v2 {
                            case "foo bar":
                                switch (v3 + v3) {
                                case -9223372036854775808:
                                    v12 := GenS1{F0: (v11 + int16(v6)), F1: 1}
                                    _ = v12
                                }
                            default:
                                var v13 int = v3
                                _ = v13
                            }
                            if ((!true) || true) {
                                var v14 int = (fi8 << ((206) & 63))
                                _ = v14
                            }
                            _ = v11
                        }
                        _ = fi10
                    }
                    _ = fi8
                    _ = v9
                }
                if true {
                    var v15 int8 = 100
                    v4[("a" + v2)] = (-731 != -556)
                    _ = v15
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        v16 := GenS1{F0: -32768, F1: 0}
                        _ = v16
                    case "hello":
                        var v17 int = -9223372036854775808
                        _ = v17
                    default:
                        v2 = ""
                    }
                }
            } else {
                if true {
                    v18 := GenS2{F0: true, F1: 258}
                    _ = v18
                } else {
                    v6 += v6
                    var v19 uint16 = 1
                    var v20 string = (v2 + v2)
                    _ = v19
                    _ = v20
                }
                v3 -= (-(v3))
            }
        default:
            var v21 uint = 18446744073709551615
            _ = v21
        }
    case "a":
        var v22 uint64 = v6
        _ = v22
    case "":
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 int64 = v1
            _ = fi23
            _ = v24
        }
    default:
        v3 = v3
    }
    var v25 int32 = 0
    v2 = "hello"
    if false {
        v5[v2] = v2
    } else {
        var v26 int16 = -1
        delete(v4, (("\n\t" + v2) + (v2 + "a")))
        _ = v26
    }
    if (1 != 1000) {
        v2 = "\n\t"
        v27 := GenS3{F0: 833, F1: -482, F2: 1000}
        switch (v3 & v3) {
        case 0:
            var v28 float32 = ((3.14159 * 1.5) + float32(0.0))
            _ = v28
        case -512:
            var v29 float64 = (0.0 * (3.14159 * (1.5 - 1.5)))
            _ = v29
        default:
            v30 := GenS3{F0: 1000, F1: (v3 | v3), F2: 863}
            _ = v30
        }
        _ = v27
    } else {
        var v31 uint32 = 2147483647
        if ((1.5 - (1.5 / 1e10)) > 3.14159) {
            var v32 uint64 = 4294967295
            switch v3 {
            case 42:
                v33 := GenS4{F0: 121, F1: 255, F2: (0.1 != float64(3.14159))}
                _ = v33
            case 32767:
                var v34 int8 = 89
                _ = v34
            case 2147483647:
                if (((0.0 * 1e10) <= -2.25) || false) {
                    fmt.Println(v2, v3)
                } else {
                    for fi35 := 0; fi35 < 1; fi35++ {
                        if (false && true) {
                            for fi36 := 0; fi36 < 11; fi36++ {
                                delete(v4, "\n\t")
                                fmt.Println(v3, v25)
                                v37 := GenS5{F0: 149, F1: 283, F2: (v31 + v31)}
                                _ = fi36
                                _ = v37
                            }
                            v38 := GenS6{F0: true, F1: 638, F2: v2}
                            _ = v38
                        }
                        _ = fi35
                    }
                }
            }
            _ = v32
        }
        var v39 int32 = (v25 ^ (v25 | (v25 + (v25 + v25))))
        _ = v31
        _ = v39
    }
    v2 = v2
    v40 := GenS3{F0: 32767, F1: 4294967295, F2: 0}
    v41 := GenS7{F0: -28, F1: (v3 % ((173) | 1)), F2: 230}
    for fi42 := 0; fi42 < 7; fi42++ {
        var v43 uint16 = 42
        v3 -= (v3 / (((^(fi42))) | 1))
        var v44 uint8 = 88
        _ = fi42
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
