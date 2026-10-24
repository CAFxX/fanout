package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 bool; F2 int }

func main() {
    var v1 int = 480
    var v2 int32 = -1
    var v3 uint32 = 77
    v4 := make(map[uint32]string)
    var v5 uint8 = 55
    var v6 int = (v1 << ((4294967295) & 63))
    v1 -= v6
    fmt.Println(v5)
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := v4[4294967295]
        _ = fi7
        _ = v8
    }
    v3 += 1
    v9 := []int{}
    v3 = (^(v3))
    if (v3 > (v3 << ((1) & 63))) {
        if true {
            for fi10 := 0; fi10 < 14; fi10++ {
                switch -32 {
                case 0:
                    v3 = uint32(v5)
                case 9223372036854775807:
                    if false {
                        var v11 float64 = ((1.5 / 1.0) + 0.0)
                        var v12 int64 = 1
                        _ = v11
                        _ = v12
                    }
                default:
                    switch (((("foo bar" + (("a" + "foo bar") + "")) + "\n\t") + "\n\t") + "a") {
                    case "a":
                        v13 := GenS1{F0: 490, F1: 531, F2: ((!(false && true)) || (!true))}
                        _ = v13
                    case "foo bar":
                        if true {
                            if ((v5 >> ((4294967295) & 63)) <= (v5 >> ((0) & 63))) {
                                _ = len(v9)
                                var v14 int16 = 32767
                                v2 = (v2 >> ((1) & 63))
                                _ = v14
                            } else {
                                delete(v4, 2147483647)
                                v5 = 0
                                if false {
                                    switch "" {
                                    case "hello":
                                        if true {
                                            if false {
                                                switch "" {
                                                case "hello":
                                                    if len(v9) >= 2 {
                                                        v9 = v9[1:len(v9)]
                                                    }
                                                default:
                                                    v2 += 1
                                                }
                                                var v15 uint64 = 169
                                                var v16 uint8 = v5
                                                _ = v15
                                                _ = v16
                                            }
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                        } else {
                                            fmt.Println(v2)
                                        }
                                    case "foo bar":
                                        for fi17 := 0; fi17 < 7; fi17++ {
                                            v1 = v1
                                            var v18 int32 = v2
                                            _ = fi17
                                            _ = v18
                                        }
                                    case "\n\t":
                                        v19 := GenS2{F0: (^(v2))}
                                        _ = v19
                                    default:
                                        switch v1 {
                                        case 127:
                                            var v20 int32 = -2147483648
                                            _ = v20
                                        case -2147483648:
                                            _ = len(v9)
                                        }
                                    }
                                } else {
                                    v3 = (v3 ^ 2147483648)
                                    v5 = v5
                                    fmt.Println(fi10, v1)
                                }
                            }
                            fmt.Println(v5, v3, v3)
                            v21 := GenS1{F0: 0, F1: 42, F2: false}
                            _ = v21
                        }
                    case "\n\t":
                        _ = len(v9)
                    }
                }
                v6 += (fi10 % ((fi10) | 1))
                var v22 int16 = 42
                _ = fi10
                _ = v22
            }
        } else {
            v23 := GenS3{F0: 32767, F1: (127 == -97), F2: v1}
            _ = v23
        }
        v6 -= 9223372036854775807
    } else {
        v6 -= v6
        if false {
            fmt.Println(v6)
            for fi24 := 0; fi24 < 15; fi24++ {
                fmt.Println(v1, v6, fi24)
                v5 += (v5 % ((v5) | 1))
                var v25 int64 = -1
                _ = fi24
                _ = v25
            }
            v2 = (v2 ^ (-(v2)))
        } else {
            v5 = (v5 | v5)
        }
    }
    if ((!((v6 % (((v1 / ((42) | 1))) | 1)) < (-(v6)))) || false) {
        _ = len(v9)
        var v26 int8 = -1
        var v27 uint8 = (v5 & 42)
        _ = v26
        _ = v27
    }
    v2 += v2
    if len(v9) > 0 {
        v9[0] = -1
    }
    if (624 != 1000) {
        if false {
            _ = len(v4)
            var v28 int = (v6 * v1)
            if (false && false) {
                if (1.5 > 0.1) {
                    v6 = v28
                } else {
                    if false {
                        v28 -= (v6 + int(v2))
                    }
                    v5 = v5
                }
                v9 = append(v9, (v28 * v28))
            }
            _ = v28
        }
    } else {
        var v29 int = (v1 >> ((18446744073709551615) & 63))
        var v30 uint = 717
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
