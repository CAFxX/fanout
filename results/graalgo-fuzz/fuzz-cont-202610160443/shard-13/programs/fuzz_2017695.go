package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 bool = false
    v2 := []uint32{73, 2147483648, 164, 1}
    v3 := []uint32{0, 855, 2}
    v4 := make(map[int]string)
    var v5 uint32 = 4294967295
    v6 := make(map[uint32]string)
    var v7 int8 = 1
    for fi8 := 0; fi8 < 14; fi8++ {
        switch fi8 {
        case 4294967295:
            v7 = 59
        case 127:
            var v9 float64 = 0.0
            _ = v9
        default:
            v5 = v5
        }
        var v10 uint = 18446744073709551615
        _ = fi8
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v2 = append(v2, 4294967295)
        for fi12 := 0; fi12 < 19; fi12++ {
            if (v1 && v1) {
                v7 = v7
                if len(v3) > 0 {
                    v3[0] = 507
                }
                if len(v2) > 0 {
                    v2[0] = 2147483648
                }
            } else {
                v13 := GenS1{F0: -864, F1: 0, F2: (-(v7))}
                _ = v13
            }
            if (v1 && v1) {
                v14 := GenS2{F0: 2}
                _ = v14
            } else {
                v5 -= (v5 & (v5 | v5))
                var v15 string = "hello"
                _ = v15
            }
            if v1 {
                switch "foo bar" {
                case "hello":
                    var v16 bool = true
                    _ = v16
                case "a":
                    if (721 == (v5 << ((969) & 63))) {
                        var v17 uint32 = (v5 & (v5 ^ (v5 << ((65535) & 63))))
                        fmt.Println(fi12)
                        _ = v17
                    } else {
                        v1 = (0 < 936)
                        if (v5 == 65535) {
                            fmt.Println(v7, fi11, v1)
                            v5 += 4294967295
                            if ((((0 <= uint(v5)) || ((("" + "\n\t") != "a") || true)) || true) && v1) {
                                v5 += (^(v5))
                                v7 -= (v7 ^ v7)
                            }
                        }
                        v18 := GenS1{F0: 954, F1: 18446744073709551615, F2: v7}
                        _ = v18
                    }
                }
                if v1 {
                    fmt.Println(v1, fi11)
                    v5 -= (v5 % ((v5) | 1))
                    switch -1 {
                    case 0:
                        var v19 bool = (!v1)
                        _ = v19
                    case -840:
                        if (1 <= 32767) {
                            v20 := GenS3{F0: 0, F1: false, F2: 2}
                            var v21 uint32 = (^(v5))
                            v3 = append(v3, 42)
                            _ = v20
                            _ = v21
                        }
                    case 127:
                        if false {
                            var v22 int = 42
                            _ = v22
                        } else {
                            v5 += v5
                            var v23 bool = ((!false) && ((255 > uint(v5)) || (675 == 651)))
                            var v24 string = "hello"
                            if true {
                                var v25 uint = 391
                                _ = v25
                            }
                            var v26 int32 = 42
                            if (false || true) {
                                v27 := GenS2{F0: 128}
                                _ = v27
                            } else {
                                fmt.Println(v24)
                            }
                            if len(v3) > 0 {
                                v3[0] = 315
                            }
                            _ = v23
                            _ = v24
                            _ = v26
                        }
                    default:
                        if (!v1) {
                            v28 := GenS2{F0: 165}
                            var v29 uint32 = (v5 << ((2) & 63))
                            var v30 uint64 = 655
                            var v31 int16 = -32768
                            _ = v28
                            _ = v29
                            _ = v30
                            _ = v31
                        } else {
                            var v32 uint = 0
                            var v33 int8 = v7
                            switch (fi11 ^ (fi12 / ((fi12) | 1))) {
                            case -9223372036854775808:
                                v34 := GenS3{F0: 209, F1: (-410 <= -561), F2: 6}
                                _ = v34
                            case 9223372036854775807:
                                v32 = 0
                            case -2147483648:
                                var v35 bool = true
                                _ = v35
                            default:
                                v32 += v32
                                fmt.Println(v5, v32, v5)
                            }
                            _ = v32
                            _ = v33
                        }
                    }
                }
                v36 := GenS3{F0: 42, F1: v1, F2: 42}
                _ = v36
            } else {
                var v37 int32 = 837
                _ = len(v3)
                _ = v37
            }
            _ = fi12
        }
        delete(v4, fi11)
        _ = fi11
    }
    if (-1 != 32767) {
        v7 = (v7 >> ((4294967295) & 63))
        var v38 int = -2147483648
        var v39 int32 = 1
        _ = v38
        _ = v39
    }
    v7 = v7
    var v40 int16 = 0
    fmt.Println(v7, v5)
    v41 := GenS4{F0: 243, F1: 255}
    for fi42 := 0; fi42 < 20; fi42++ {
        _ = len(v3)
        for fi43 := 0; fi43 < 12; fi43++ {
            v44 := GenS2{F0: 128}
            var v45 int64 = -9223372036854775808
            _ = fi43
            _ = v44
            _ = v45
        }
        var v46 int8 = v7
        _ = fi42
        _ = v46
    }
    var v47 string = (("hello" + "foo bar") + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
