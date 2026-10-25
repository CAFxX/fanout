package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int64; F2 int }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 bool; F1 uint16 }

func main() {
    var v1 uint = 42
    v2 := make(map[uint32]uint32)
    var v3 int = -9223372036854775808
    v4 := []uint32{2, 533, 1}
    var v5 string = ""
    v6 := GenS1{F0: 381, F1: 18446744073709551615}
    v7 := GenS1{F0: 151, F1: 397}
    var v8 int = -1
    var v9 int64 = 36
    v8 = (-(v3))
    switch v8 {
    case 2147483647:
        v5 = v5
    case 780:
        var v10 int16 = int16(v9)
        _ = v10
    }
    if len(v4) > 0 {
        v4[0] = 950
    }
    if (true && false) {
        var v11 int32 = 42
        _ = v11
    } else {
        switch v5 {
        case "\n\t":
            v3 -= (v3 + (v3 - int(v9)))
        default:
            var v12 float64 = -2.25
            _ = v12
        }
        delete(v2, 364)
        v9 = (v9 + 42)
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := v2[305]
        for fi15 := 0; fi15 < 11; fi15++ {
            v16 := v2[(^(v14))]
            v17 := GenS1{F0: 258, F1: 254}
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi13
        _ = v14
    }
    var v18 uint64 = 0
    v19 := GenS1{F0: v18, F1: (v18 << (((v1 + (v1 | 0))) & 63))}
    for fi20 := 0; fi20 < 3; fi20++ {
        v3 = fi20
        _ = fi20
    }
    v8 = v8
    var v21 uint = (v1 / (((v1 ^ 310)) | 1))
    for fi22 := 0; fi22 < 19; fi22++ {
        if (true || false) {
            v23 := GenS1{F0: v18, F1: v18}
            switch ((v5 + v5) + v5) {
            case "":
                switch ("foo bar" + "foo bar") {
                case "a":
                    v24 := v2[173]
                    _ = v24
                case "foo bar":
                    var v25 int = (v8 / ((32768) | 1))
                    _ = v25
                default:
                    v26 := GenS1{F0: 1, F1: v18}
                    _ = v26
                }
            case "hello":
                var v27 uint64 = 0
                _ = v27
            default:
                v28 := GenS2{F0: 42, F1: (v9 % (((^(v9))) | 1)), F2: v8}
                _ = v28
            }
            for fi29 := 0; fi29 < 11; fi29++ {
                v3 += v3
                v30 := GenS3{F0: 42, F1: 554}
                if ((v9 ^ v9) >= v9) {
                    switch 472 {
                    case 32768:
                        var v31 uint8 = 1
                        _ = v31
                    case 282:
                        v18 = (v18 & (v18 % ((v18) | 1)))
                    default:
                        var v32 bool = (false || (float32(1e10) < float32(-2.25)))
                        _ = v32
                    }
                    if false {
                        v33 := GenS1{F0: (v18 & v18), F1: (^(v18))}
                        _ = v33
                    } else {
                        _ = len(v4)
                        if false {
                            if len(v4) > 0 {
                                v4[0] = 4294967295
                            }
                        } else {
                            v5 = "foo bar"
                            switch v5 {
                            case "\n\t":
                                if false {
                                    if (42 <= 65535) {
                                        var v34 int32 = -163
                                        v3 -= fi29
                                        _ = v34
                                    } else {
                                        var v35 int = (fi29 | -2147483648)
                                        v36 := GenS4{F0: (v21 & v1), F1: (v9 * (v9 >> ((v21) & 63)))}
                                        _ = v35
                                        _ = v36
                                    }
                                    fmt.Println(v1, v3, fi29)
                                    v18 += (v18 % ((975) | 1))
                                } else {
                                    var v37 string = "a"
                                    if false {
                                        v38 := GenS5{F0: (false || true), F1: 42}
                                        _ = v38
                                    } else {
                                        var v39 int = -991
                                        _ = v39
                                    }
                                    if true {
                                        fmt.Println(v3)
                                    } else {
                                        if false {
                                            fmt.Println(v9, v8, v18)
                                            v40 := GenS5{F0: true, F1: 255}
                                            var v41 int = (v3 + v3)
                                            switch (((v5 + (v5 + (v5 + v5))) + "hello") + v37) {
                                            case "foo bar":
                                                v42 := GenS4{F0: v1, F1: v9}
                                                _ = v42
                                            default:
                                                fmt.Println(v41)
                                            }
                                            v1 = v1
                                            _ = v40
                                            _ = v41
                                        }
                                        v43 := GenS5{F0: (!false), F1: 32767}
                                        switch int(v18) {
                                        case 968:
                                            v44 := GenS5{F0: false, F1: 467}
                                            _ = v44
                                        case 308:
                                            v45 := GenS2{F0: 711, F1: v9, F2: -370}
                                            _ = v45
                                        case 910:
                                            var v46 int32 = -824
                                            _ = v46
                                        }
                                        _ = v43
                                    }
                                    var v47 uint = v21
                                    var v48 uint64 = (^(v18))
                                    _ = v37
                                    _ = v47
                                    _ = v48
                                }
                            }
                        }
                    }
                    var v49 uint8 = 140
                    _ = v49
                } else {
                    if true {
                        v5 = (((v5 + v5) + "foo bar") + v5)
                    }
                    v18 = (v18 ^ v18)
                    var v50 int16 = 1
                    _ = v50
                }
                if (520 > 2) {
                    fmt.Println(v5, fi29)
                    fmt.Println(v9, v8)
                    v4 = append(v4, 2147483648)
                }
                var v51 int16 = -32768
                _ = fi29
                _ = v30
                _ = v51
            }
            _ = v23
        }
        if true {
            var v52 uint32 = 42
            fmt.Println(v52)
            var v53 uint16 = 297
            _ = v52
            _ = v53
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
