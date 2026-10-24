package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint8 }
type GenS2 struct { F0 int64; F1 int32; F2 uint64 }
type GenS3 struct { F0 int; F1 uint8 }

func main() {
    v1 := []int{0, 1, 887}
    var v2 uint16 = 2
    v3 := make(map[uint32]uint32)
    v4 := []uint32{633, 267}
    var v5 uint = 0
    var v6 int16 = -32768
    switch ("hello" + "") {
    case "a":
        var v7 uint8 = 2
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        if ("a" < "\n\t") {
            var v9 bool = true
            fmt.Println(fi8, v2, v6)
            var v10 bool = (!(0.1 < float64(float32(0.1))))
            _ = v10
            _ = v9
        } else {
            if ((1.5 / (0.1 + 1e10)) > 1e10) {
                v6 -= v6
            } else {
                fmt.Println(v5)
                for fi11 := 0; fi11 < 13; fi11++ {
                    if true {
                        fmt.Println(v6, v6)
                    } else {
                        switch -2147483648 {
                        case 127:
                            switch "a" {
                            case "\n\t":
                                v2 += 705
                            case "hello":
                                v12 := v3[4294967295]
                                _ = v12
                            case "":
                                fmt.Println(fi8, v5, fi11)
                            default:
                                v5 = uint(v2)
                            }
                        case 176:
                            v2 += (^(v2))
                        }
                        if true {
                            switch ("hello" + "foo bar") {
                            case "foo bar":
                                var v13 string = "foo bar"
                                fmt.Println(v5, fi11, v6)
                                _ = v13
                            case "a":
                                v3[2147483647] = 0
                            default:
                                v6 = v6
                            }
                            if len(v1) > 0 {
                                v1[0] = (fi8 * (fi11 * fi8))
                            }
                            switch (fi11 - fi11) {
                            case 386:
                                switch 128 {
                                case 4294967295:
                                    if (127 == 128) {
                                        v14 := GenS1{F0: (v6 ^ -32768), F1: int32(v6), F2: 0}
                                        _ = v14
                                    } else {
                                        var v15 uint8 = 127
                                        _ = v15
                                    }
                                case -1:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case 1:
                                    var v16 float32 = float32(3.14159)
                                    _ = v16
                                default:
                                    if (true || (!true)) {
                                        v2 += 2
                                        fmt.Println(v2)
                                    }
                                }
                            default:
                                switch 580 {
                                case -2147483648:
                                    _ = len(v4)
                                }
                            }
                            var v17 int16 = 1
                            if true {
                                if false {
                                    if (28 >= 42) {
                                        var v18 float32 = float32((0.0 * 1e10))
                                        _ = v18
                                    } else {
                                        delete(v3, 2147483647)
                                    }
                                } else {
                                    v5 += 214
                                    var v19 int64 = 1
                                    v20 := GenS1{F0: (v6 % ((-1) | 1)), F1: int32(v2), F2: 52}
                                    _ = v19
                                    _ = v20
                                }
                            } else {
                                var v21 uint = 255
                                _ = v21
                            }
                            _ = v17
                        } else {
                            v6 = (v6 * (v6 << ((96) & 63)))
                            if false {
                                fmt.Println(v2, fi11, v5)
                            } else {
                                v2 = v2
                                if true {
                                    v6 -= (v6 % ((v6) | 1))
                                    v22 := GenS1{F0: v6, F1: -896, F2: 84}
                                    var v23 uint8 = 2
                                    _ = v22
                                    _ = v23
                                } else {
                                    if true {
                                        v5 = (v5 + v5)
                                        v24 := GenS1{F0: (-(v6)), F1: 0, F2: 128}
                                        _ = v24
                                    }
                                }
                                if (936 >= 782) {
                                    var v25 float64 = float64((-2.25 / ((-2.25 / (((1.5 + 1.5) / 1.0) + -2.25)) / 0.1)))
                                    _ = v25
                                }
                            }
                        }
                        v26 := GenS2{F0: 0, F1: 1, F2: 720}
                        _ = v26
                    }
                    fmt.Println(v6, v2, v6)
                    v27 := GenS1{F0: (^(v6)), F1: -1, F2: 42}
                    var v28 int32 = 42
                    _ = fi11
                    _ = v27
                    _ = v28
                }
            }
            var v29 float64 = 0.1
            var v30 uint8 = 186
            _ = v29
            _ = v30
        }
        fmt.Println(v6)
        var v31 int = (fi8 & (fi8 % ((fi8) | 1)))
        _ = fi8
        _ = v31
    }
    v32 := GenS2{F0: 1, F1: 2147483647, F2: 54}
    fmt.Println(v5)
    if ((v5 + 65535) == (^(v5))) {
        v2 += (^(v2))
    }
    v4 = append(v4, 42)
    var v33 uint64 = 0
    switch "foo bar" {
    case "\n\t":
        v34 := GenS3{F0: -9223372036854775808, F1: 128}
        _ = v34
    case "hello":
        v6 += (v6 ^ v6)
    case "foo bar":
        var v35 uint = (^(v5))
        _ = v35
    default:
        fmt.Println(v2, v2, v33)
    }
    var v36 int64 = -9223372036854775808
    if true {
        v5 = uint(v2)
        for fi37 := 0; fi37 < 5; fi37++ {
            var v38 uint8 = 81
            if true {
                v5 = v5
                _ = len(v1)
                if (true || (false || (!false))) {
                    for fi39 := 0; fi39 < 16; fi39++ {
                        v38 -= v38
                        delete(v3, 65535)
                        switch (fi37 >> (((v5 + v5)) & 63)) {
                        case -1:
                            var v40 int32 = 1
                            _ = v40
                        default:
                            v5 += (^(v5))
                            v33 -= (v33 * v33)
                            var v41 uint16 = (v2 | (^(v2)))
                            _ = v41
                        }
                        _ = fi39
                    }
                } else {
                    v42 := GenS3{F0: fi37, F1: (^(v38))}
                    _ = v42
                }
            }
            v38 += uint8(v2)
            _ = fi37
            _ = v38
        }
    } else {
        var v43 float32 = 1e10
        var v44 float32 = (-2.25 + (v43 + v43))
        _ = v43
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
