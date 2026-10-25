package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 float64 = 1e10
    var v2 uint8 = 53
    var v3 uint16 = 0
    var v4 int16 = 548
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]int)
    v3 = (v3 % ((v3) | 1))
    v7 := GenS1{F0: 735, F1: (-(v4))}
    var v8 int = 128
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS1{F0: 686, F1: (v4 % ((v4) | 1))}
        v4 = (^(v4))
        _ = fi9
        _ = v10
    }
    var v11 bool = false
    v4 = int16(v3)
    fmt.Println(v8, v11, v11)
    var v12 int64 = 756
    v13 := v6[65535]
    var v14 int64 = -637
    v15 := GenS2{F0: v4, F1: (42 <= int32(v3))}
    v16 := GenS3{F0: (v4 ^ (v4 << ((4294967295) & 63)))}
    v4 = 1
    switch 32767 {
    case 2147483647:
        v17 := GenS3{F0: v4}
        _ = v17
    case 128:
        if false {
            v18 := GenS1{F0: 4294967295, F1: v4}
            var v19 int16 = int16(v13)
            var v20 int8 = 92
            _ = v18
            _ = v19
            _ = v20
        } else {
            for fi21 := 0; fi21 < 18; fi21++ {
                var v22 string = ("a" + "foo bar")
                v23 := GenS2{F0: (^(v4)), F1: (!false)}
                if (((v14 / ((v12) | 1)) > v12) && (65535 >= 223)) {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        switch -652 {
                        case -249:
                            switch v8 {
                            case -1:
                                switch v22 {
                                case "a":
                                    v11 = (4294967295 == (fi21 << ((65535) & 63)))
                                default:
                                    var v25 int16 = 1
                                    _ = v25
                                }
                            default:
                                var v26 int = (fi21 | fi21)
                                _ = v26
                            }
                        default:
                            v4 = 920
                            var v27 int64 = (^(v12))
                            v12 -= (v27 % ((v27) | 1))
                            _ = v27
                        }
                        var v28 uint64 = 42
                        _ = fi24
                        _ = v28
                    }
                    fmt.Println(v14, v4)
                    var v29 uint64 = 18446744073709551615
                    _ = v29
                }
                _ = fi21
                _ = v22
                _ = v23
            }
        }
    case -2147483648:
        var v30 int = v8
        _ = v30
    default:
        if v11 {
            for fi31 := 0; fi31 < 12; fi31++ {
                v13 = (v8 & (^(v13)))
                if ((-417 == -1) || ((v11 && (23 <= 534)) || v11)) {
                    switch "\n\t" {
                    case "a":
                        v32 := GenS2{F0: 1, F1: (2 > 609)}
                        _ = v32
                    default:
                        v33 := GenS2{F0: -32768, F1: (((0.0 + v1) - 0.0) > v1)}
                        _ = v33
                    }
                    delete(v5, 447)
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        for fi34 := 0; fi34 < 16; fi34++ {
                            fmt.Println(v2, v14, v13)
                            switch (v13 * fi31) {
                            case -1:
                                _ = len(v6)
                            case 783:
                                v1 = v1
                            default:
                                var v35 uint16 = 0
                                _ = v35
                            }
                            switch "hello" {
                            case "a":
                                switch (v13 << ((4294967295) & 63)) {
                                case 128:
                                    var v36 int64 = (v14 / (((^(v12))) | 1))
                                    _ = v36
                                }
                            default:
                                v37 := GenS4{F0: (-(v4))}
                                _ = v37
                            }
                            _ = fi34
                        }
                    default:
                        for fi38 := 0; fi38 < 20; fi38++ {
                            v3 = (v3 ^ (v3 * 1))
                            _ = fi38
                        }
                    }
                }
                v39 := GenS4{F0: (v4 % ((32767) | 1))}
                _ = fi31
                _ = v39
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
