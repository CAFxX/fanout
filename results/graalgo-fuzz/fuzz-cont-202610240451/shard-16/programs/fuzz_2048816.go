package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint8; F1 string }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = ("" + ("" + "hello"))
    var v3 int32 = 2147483647
    v4 := make(map[int]bool)
    var v5 uint8 = 0
    v6 := []uint32{0, 65535, 0}
    v5 += (v5 * 0)
    v3 -= -2147483648
    v3 = v3
    v7 := v4[0]
    for fi8 := 0; fi8 < 10; fi8++ {
        if (!(858 < 0)) {
            var v9 int = (fi8 * (fi8 + 2147483647))
            v7 = v7
            _ = v9
        }
        v3 += 1
        _ = fi8
    }
    v7 = v7
    if true {
        switch ((v2 + "\n\t") + "\n\t") {
        case "\n\t":
            v3 = (v3 >> ((4294967295) & 63))
        case "":
            if false {
                v7 = v7
                for fi10 := 0; fi10 < 1; fi10++ {
                    var v11 int = -492
                    _ = fi10
                    _ = v11
                }
            } else {
                v12 := GenS1{F0: v5, F1: 2, F2: 32767}
                v3 -= 356
                _ = v12
            }
        case "foo bar":
            v5 = v5
        }
    }
    if true {
        v5 -= v5
        v13 := GenS1{F0: (v5 ^ (v5 ^ 255)), F1: 828, F2: 42}
        if v7 {
            switch v2 {
            case "a":
                if (3.14159 < 3.14159) {
                    var v14 uint64 = 0
                    _ = len(v4)
                    _ = v14
                } else {
                    v2 = ""
                }
            default:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            var v15 uint8 = v5
            _ = v15
        } else {
            if (9223372036854775807 <= 0) {
                if v7 {
                    v16 := GenS2{F0: v2}
                    _ = v16
                } else {
                    v5 -= v5
                    for fi17 := 0; fi17 < 14; fi17++ {
                        if ((!v7) && false) {
                            v7 = (!true)
                            delete(v1, 65535)
                        } else {
                            var v18 uint32 = 2
                            if v7 {
                                v2 = v2
                                v3 = -224
                                v19 := GenS1{F0: v5, F1: (v18 - v18), F2: 65535}
                                _ = v19
                            }
                            for fi20 := 0; fi20 < 6; fi20++ {
                                var v21 int16 = 1
                                v7 = (42 <= 1)
                                if v7 {
                                    v3 += 802
                                }
                                _ = fi20
                                _ = v21
                            }
                            _ = v18
                        }
                        v5 -= (v5 ^ (v5 ^ (v5 / ((v5) | 1))))
                        _ = fi17
                    }
                    v22 := v1[4294967295]
                    _ = v22
                }
                v23 := GenS3{F0: v2, F1: 0}
                _ = v23
            }
            v24 := GenS2{F0: v2}
            if v7 {
                if true {
                    if v7 {
                        var v25 bool = ((float64(3.14159) - (1e10 - 1.5)) >= 1.5)
                        fmt.Println(v7, v7)
                        for fi26 := 0; fi26 < 6; fi26++ {
                            fmt.Println(v2, v2, v25)
                            for fi27 := 0; fi27 < 2; fi27++ {
                                var v28 uint = 2
                                if v7 {
                                    var v29 uint64 = 18446744073709551615
                                    var v30 uint64 = (^(v29))
                                    if v25 {
                                        if len(v6) > 0 {
                                            v6[0] = 2147483648
                                        }
                                        var v31 int32 = (v3 & (v3 + v3))
                                        _ = v31
                                    }
                                    _ = v29
                                    _ = v30
                                } else {
                                    v32 := GenS4{F0: v5, F1: v2}
                                    v7 = true
                                    if v25 {
                                        var v33 int8 = -124
                                        _ = v33
                                    }
                                    if (v7 && true) {
                                        v3 = (v3 % (((v3 * -2147483648)) | 1))
                                        v34 := GenS1{F0: (^(v5)), F1: 65535, F2: 1}
                                        var v35 int64 = 540
                                        var v36 int16 = -297
                                        var v37 int = 4294967295
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    } else {
                                        var v38 uint64 = 2
                                        _ = v38
                                    }
                                    _ = v32
                                }
                                _ = fi27
                                _ = v28
                            }
                            _ = fi26
                        }
                        _ = v25
                    } else {
                        var v39 uint8 = (^(v5))
                        _ = v39
                    }
                    switch 161 {
                    case 32767:
                        v40 := GenS3{F0: v2, F1: 9223372036854775807}
                        _ = v40
                    default:
                        v2 = v2
                    }
                }
            }
            _ = v24
        }
        _ = v13
    } else {
        if v7 {
            var v41 int16 = 1
            _ = len(v1)
            _ = v41
        }
        v3 = -983
        for fi42 := 0; fi42 < 9; fi42++ {
            var v43 uint32 = 759
            switch ((("\n\t" + (v2 + "a")) + v2) + v2) {
            case "a":
                v5 += (v5 << ((2) & 63))
            case "foo bar":
                v43 = v43
            }
            fmt.Println(v43, v3, v5)
            _ = fi42
            _ = v43
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
