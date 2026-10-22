package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS6 struct { F0 int32; F1 bool; F2 bool }

func main() {
    var v1 int8 = -128
    var v2 int32 = 1
    v3 := make(map[int]int)
    var v4 bool = false
    var v5 int8 = (v1 % ((-19) | 1))
    var v6 float32 = float32(3.14159)
    v2 = 42
    v7 := []int{127, 349, 128}
    fmt.Println(v2)
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int32 = (v2 + (v2 >> ((2) & 63)))
        if (!((false && false) || true)) {
            v5 += (v5 ^ (v1 + 42))
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 uint32 = 2147483647
                switch "a" {
                case "\n\t":
                    var v12 float32 = 3.14159
                    _ = v12
                default:
                    v5 = -80
                }
                var v13 uint8 = 42
                if false {
                    var v14 uint32 = v11
                    if len(v7) > 0 {
                        v7[0] = (fi8 / (((fi10 ^ fi8)) | 1))
                    }
                    _ = v14
                } else {
                    if v4 {
                        fmt.Println(v6, fi8, v5)
                    } else {
                        _ = len(v3)
                        _ = len(v7)
                        fmt.Println(v5)
                        var v15 int64 = -126
                        if (((fi8 % (((fi10 & 127)) | 1)) <= fi8) && v4) {
                            switch "a" {
                            case "":
                                _ = len(v3)
                            case "a":
                                fmt.Println(v15, fi10, v9)
                            }
                            v16 := GenS1{F0: 2147483648}
                            var v17 uint32 = (v11 + (v11 * v11))
                            _ = v16
                            _ = v17
                        } else {
                            v3[(fi10 + int(v13))] = fi8
                            _ = len(v3)
                            _ = len(v3)
                        }
                        _ = v15
                    }
                    if v4 {
                        var v18 uint64 = 4294967295
                        fmt.Println(fi8)
                        if ((1e10 > ((0.1 + -2.25) - -2.25)) && (!(341 != -1))) {
                            switch (fi10 - 193) {
                            case 2147483647:
                                v18 = 42
                            case 1:
                                v13 = (v13 & 1)
                            case -9223372036854775808:
                                var v19 int16 = 42
                                _ = v19
                            default:
                                v3[(^(fi8))] = (fi10 & fi8)
                            }
                        } else {
                            v20 := GenS2{F0: true}
                            v9 = 507
                            var v21 int8 = v5
                            _ = v20
                            _ = v21
                        }
                        _ = v18
                    } else {
                        v22 := GenS3{F0: fi8, F1: -647}
                        var v23 int8 = (v5 * (v1 * v5))
                        var v24 int32 = v9
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    if v4 {
                        fmt.Println(v1)
                        switch "" {
                        case "hello":
                            v2 -= (v9 << ((2) & 63))
                        case "":
                            var v25 int64 = 9223372036854775807
                            _ = v25
                        }
                        if (32767 == 1) {
                            var v26 uint64 = 0
                            switch "" {
                            case "foo bar":
                                _ = len(v3)
                            }
                            _ = v26
                        }
                        switch "hello" {
                        case "a":
                            var v27 float32 = float32(-2.25)
                            _ = v27
                        }
                    }
                    var v28 int8 = (v5 | v5)
                    _ = len(v3)
                    _ = v28
                }
                var v29 uint = 18446744073709551615
                _ = fi10
                _ = v11
                _ = v13
                _ = v29
            }
            var v30 float64 = (float64(v6) / 0.1)
            _ = v30
        }
        _ = fi8
        _ = v9
    }
    var v31 int8 = v1
    var v32 int64 = -1
    if false {
        v33 := GenS4{F0: 42, F1: 519}
        for fi34 := 0; fi34 < 19; fi34++ {
            if v4 {
                if true {
                    v2 -= v2
                } else {
                    switch (fi34 - fi34) {
                    case -1:
                        if v4 {
                            v5 -= (v31 | v1)
                        }
                    case 9223372036854775807:
                        fmt.Println(v32, v2)
                    }
                    v35 := GenS5{F0: 865, F1: -985, F2: 1}
                    _ = v35
                }
            }
            for fi36 := 0; fi36 < 2; fi36++ {
                var v37 uint8 = 128
                var v38 uint32 = 0
                if (v4 || (42 > 1)) {
                    switch ("a" + ("a" + "foo bar")) {
                    case "foo bar":
                        v39 := GenS2{F0: (v4 || v4)}
                        _ = v39
                    }
                    if (-2.25 < 0.0) {
                        var v40 uint16 = 42
                        _ = v40
                    }
                } else {
                    v41 := GenS2{F0: v4}
                    var v42 int = fi34
                    v38 = v38
                    _ = v41
                    _ = v42
                }
                _ = fi36
                _ = v37
                _ = v38
            }
            var v43 int64 = 0
            _ = fi34
            _ = v43
        }
        _ = v33
    } else {
        for fi44 := 0; fi44 < 8; fi44++ {
            if ((42 > 809) && false) {
                if len(v7) > 0 {
                    v7[0] = 128
                }
            }
            _ = fi44
        }
        if (v4 && v4) {
            var v45 float32 = (1.5 * 1e10)
            v46 := GenS6{F0: v2, F1: (!true), F2: ((-128 <= (v5 | v31)) || false)}
            _ = v45
            _ = v46
        }
        var v47 int = 1
        _ = v47
    }
    v2 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
