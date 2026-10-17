package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint32 }
type GenS2 struct { F0 int; F1 int64; F2 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 4294967295
    var v3 float64 = (1e10 - 0.1)
    var v4 int32 = -2147483648
    v5 := make(map[int]uint32)
    var v6 uint = 470
    var v7 int8 = 1
    var v8 uint32 = 442
    fmt.Println(v2)
    fmt.Println(v4, v7)
    for fi9 := 0; fi9 < 9; fi9++ {
        switch fi9 {
        case 128:
            v10 := GenS1{F0: -1, F1: (-(v7)), F2: 2147483647}
            _ = v10
        case 969:
            delete(v5, fi9)
        case -135:
            var v11 uint8 = 177
            _ = v11
        default:
            v12 := GenS1{F0: int16(v8), F1: 0, F2: (v8 >> (((v6 >> ((v6) & 63))) & 63))}
            _ = v12
        }
        v13 := GenS2{F0: (fi9 << ((v6) & 63)), F1: 1, F2: (v6 / ((42) | 1))}
        _ = fi9
        _ = v13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 int8 = v7
        for fi16 := 0; fi16 < 14; fi16++ {
            if false {
                v17 := GenS2{F0: fi14, F1: 9223372036854775807, F2: (v2 ^ (v6 << (((v6 << ((4294967295) & 63))) & 63)))}
                var v18 uint = (v6 & 2)
                switch "" {
                case "\n\t":
                    if true {
                        switch (("" + "foo bar") + ("foo bar" + "hello")) {
                        case "":
                            var v19 int8 = (v15 << ((449) & 63))
                            _ = v19
                        case "hello":
                            fmt.Println(v2, v3)
                        case "foo bar":
                            if true {
                                switch fi14 {
                                case 42:
                                    var v20 uint32 = (v8 + (v8 >> (((v6 - 993)) & 63)))
                                    _ = v20
                                case -2147483648:
                                    v6 = v2
                                default:
                                    v18 += 2
                                }
                            } else {
                                v4 = (v4 % (((-(v4))) | 1))
                                switch "\n\t" {
                                case "":
                                    var v21 int32 = v4
                                    _ = v21
                                case "a":
                                    var v22 int8 = -64
                                    _ = v22
                                }
                                v23 := GenS2{F0: (-(fi16)), F1: 9223372036854775807, F2: 18446744073709551615}
                                _ = v23
                            }
                        default:
                            if true {
                                fmt.Println(fi16, v18, v3)
                                v4 = v4
                            }
                        }
                        if (v4 >= (-(v4))) {
                            var v24 int64 = 1
                            v25 := []int{fi14, 32768}
                            _ = v24
                            _ = v25
                        }
                    } else {
                        fmt.Println(v18)
                        var v26 uint8 = 42
                        switch fi16 {
                        case 2147483647:
                            var v27 bool = ("hello" <= ("foo bar" + ""))
                            _ = v27
                        case 32768:
                            var v28 uint32 = v8
                            _ = v28
                        case -373:
                            if false {
                                v29 := []int{(fi16 / (((fi16 >> ((v2) & 63))) | 1)), int(v8), 42}
                                v26 += uint8(v2)
                                if false {
                                    switch "hello" {
                                    case "foo bar":
                                        v30 := GenS1{F0: 42, F1: (v7 << ((0) & 63)), F2: (v8 & (v8 / (((^(v8))) | 1)))}
                                        _ = v30
                                    }
                                }
                                _ = v29
                            }
                        default:
                            v7 += v15
                        }
                        var v31 int16 = 431
                        _ = v26
                        _ = v31
                    }
                default:
                    v32 := GenS2{F0: (fi14 % ((fi16) | 1)), F1: 9223372036854775807, F2: (v6 - (v6 + 255))}
                    _ = v32
                }
                _ = v17
                _ = v18
            }
            var v33 uint32 = v8
            _ = fi16
            _ = v33
        }
        var v34 uint32 = (v8 ^ 0)
        _ = fi14
        _ = v15
        _ = v34
    }
    var v35 int8 = 8
    for fi36 := 0; fi36 < 1; fi36++ {
        fmt.Println(v6, v6, v6)
        _ = fi36
    }
    var v37 string = "foo bar"
    var v38 uint = 42
    if (42 >= -44) {
        switch ("foo bar" + (v37 + ("" + ("foo bar" + v37)))) {
        case "hello":
            for fi39 := 0; fi39 < 3; fi39++ {
                v40 := []int{9223372036854775807, (fi39 ^ (^(fi39))), (fi39 >> (((^(v6))) & 63)), fi39}
                fmt.Println(v38, v38, v37)
                _ = fi39
                _ = v40
            }
        }
    } else {
        var v41 int32 = (v4 % ((v4) | 1))
        _ = v41
    }
    var v42 uint32 = v8
    var v43 uint32 = v42
    fmt.Println(v8, v4, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
