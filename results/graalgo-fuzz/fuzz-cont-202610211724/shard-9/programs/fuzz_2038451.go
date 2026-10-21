package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int8 }

func main() {
    var v1 float32 = 1.5
    var v2 bool = (0 == 0)
    var v3 uint64 = 176
    var v4 float64 = float64(3.14159)
    v5 := make(map[int]uint32)
    v4 = ((float64((1e10 * v1)) + -2.25) + ((float64(float32(v4)) / float64(0.1)) + v4))
    v3 += v3
    v3 = (v3 % ((v3) | 1))
    v1 -= v1
    var v6 int = 42
    switch "foo bar" {
    case "":
        if v2 {
            v7 := GenS1{F0: 2147483647}
            _ = v7
        }
    default:
        if (!(255 <= 0)) {
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 uint32 = 920
                v10 := GenS2{F0: (^(v9)), F1: -844}
                var v11 uint64 = (v3 << ((589) & 63))
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            v2 = v2
            for fi12 := 0; fi12 < 3; fi12++ {
                for fi13 := 0; fi13 < 9; fi13++ {
                    if false {
                        if (int16(fi12) > 32767) {
                            v14 := GenS1{F0: -2147483648}
                            v15 := v5[fi12]
                            v16 := []int{}
                            _ = v14
                            _ = v15
                            _ = v16
                        } else {
                            v2 = v2
                        }
                        v1 -= v1
                        fmt.Println(v6)
                    } else {
                        switch "foo bar" {
                        case "hello":
                            if ((!v2) || false) {
                                fmt.Println(v6)
                            }
                        case "foo bar":
                            var v17 uint64 = v3
                            _ = v17
                        case "\n\t":
                            if (3.14159 != 0.0) {
                                var v18 uint8 = 1
                                _ = v18
                            }
                        default:
                            if v2 {
                                var v19 uint16 = 42
                                var v20 float32 = (((v1 + (-2.25 * v1)) / v1) + v1)
                                _ = v19
                                _ = v20
                            }
                            if (false || (2147483648 != 65535)) {
                                v21 := []int{v6, (fi12 / ((-660) | 1))}
                                _ = v21
                            } else {
                                v22 := GenS2{F0: 2, F1: -32768}
                                v1 = v1
                                _ = v22
                            }
                            fmt.Println(v4, fi13)
                        }
                        v1 = float32((v4 * (v4 / (-2.25 * (((-2.25 * -2.25) / (v4 / (3.14159 * v4))) * v4)))))
                        v23 := []int{(fi13 & fi13)}
                        _ = v23
                    }
                    _ = fi13
                }
                _ = fi12
            }
        }
    }
    v24 := GenS3{F0: -128, F1: uint16(v3)}
    var v25 uint16 = 74
    v26 := []int{}
    switch (v6 + (v6 - 0)) {
    case -1:
        var v27 uint = 0
        _ = v27
    case 32768:
        var v28 int64 = 42
        _ = v28
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        if ((v25 | 65535) < (v25 % (((v25 % (((v25 & 719)) | 1))) | 1))) {
            for fi30 := 0; fi30 < 15; fi30++ {
                switch "foo bar" {
                case "hello":
                    fmt.Println(v25, v4)
                default:
                    v25 += (v25 * 32767)
                }
                if (v1 >= ((1.5 + (0.0 / 1e10)) + (((v1 - v1) + (v1 + v1)) + (float32(v4) + (v1 + ((float32(v4) + v1) / float32((v4 * v4)))))))) {
                    v1 -= (v1 - 0.0)
                    v31 := GenS4{F0: 42, F1: 82}
                    _ = v31
                } else {
                    switch v6 {
                    case -9223372036854775808:
                        v2 = v2
                    default:
                        var v32 int16 = 906
                        _ = v32
                    }
                    v33 := GenS4{F0: -32768, F1: 100}
                    _ = v33
                }
                switch "\n\t" {
                case "\n\t":
                    if (v2 && (1.5 >= float32(v4))) {
                        var v34 int16 = 32767
                        if ((v3 >= 18446744073709551615) || (int16(v25) >= 32767)) {
                            if len(v26) >= 2 {
                                v26 = v26[1:len(v26)]
                            }
                        }
                        v4 -= v4
                        delete(v5, (v6 ^ 127))
                        v4 -= 1e10
                        _ = v34
                    } else {
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        v25 = 255
                        switch "\n\t" {
                        case "foo bar":
                            switch (("foo bar" + (("a" + ("hello" + "a")) + "a")) + ((("foo bar" + (("a" + "\n\t") + (("" + "a") + "hello"))) + ("hello" + "\n\t")) + "foo bar")) {
                            case "\n\t":
                                var v35 uint16 = (v25 ^ 91)
                                _ = v35
                            case "hello":
                                v36 := GenS5{F0: 127}
                                _ = v36
                            }
                        case "a":
                            if len(v26) > 0 {
                                v26[0] = fi30
                            }
                            v37 := v5[32768]
                            v37 = (v37 + 326)
                            _ = v37
                        }
                    }
                case "":
                    v2 = (v1 != 0.1)
                }
                _ = fi30
            }
        } else {
            _ = len(v26)
            var v38 int8 = 127
            for fi39 := 0; fi39 < 12; fi39++ {
                v3 += 18446744073709551615
                var v40 uint32 = 1
                _ = len(v5)
                _ = fi39
                _ = v40
            }
            _ = v38
        }
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
        var v41 uint16 = v25
        _ = fi29
        _ = v41
    }
    for fi42 := 0; fi42 < 6; fi42++ {
        v1 += 1.5
        _ = fi42
    }
    if true {
        v6 = 32767
        if true {
            for fi43 := 0; fi43 < 7; fi43++ {
                for fi44 := 0; fi44 < 14; fi44++ {
                    switch fi43 {
                    case 9223372036854775807:
                        v45 := GenS1{F0: 42}
                        _ = v45
                    case -607:
                        var v46 uint16 = (v25 << ((42) & 63))
                        _ = v46
                    case 546:
                        fmt.Println(v4, v4)
                    }
                    fmt.Println(v4, v2, v6)
                    _ = fi44
                }
                _ = fi43
            }
        } else {
            v1 += v1
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
            if (284 <= 181) {
                if (238 < (v25 ^ v25)) {
                    var v47 uint64 = 4294967295
                    _ = v47
                } else {
                    v3 += v3
                    fmt.Println(v3, v3, v6)
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
