package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int64 }
func genFn0() uint8 {
    var v7 int = 0
    v7 = 2147483647
    _ = v7
    return 24
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -32768
    v3 := make(map[int]int)
    v4 := []string{"a", "a"}
    var v5 int16 = (v2 - 1)
    v6 := []string{(("" + "foo bar") + ("\n\t" + "foo bar"))}
    for fi8 := 0; fi8 < 9; fi8++ {
        v5 = v5
        _ = fi8
    }
    var v9 uint = 4294967295
    fmt.Println(v5, v5)
    if (0 <= -128) {
        v2 -= v2
        var v10 int = -1
        _ = v10
    }
    switch ("" + "a") {
    case "foo bar":
        switch 366 {
        case 4294967295:
            fmt.Println(v5, v2)
        case -24:
            fmt.Println(v2)
        }
    case "\n\t":
        v5 -= v2
    case "hello":
        v9 = (^(v9))
    default:
        switch "\n\t" {
        case "":
            var v11 int8 = 100
            _ = v11
        }
    }
    if ((-909 <= int64(v9)) || (127 > 4294967295)) {
        switch "" {
        case "hello":
            fmt.Println(v9, v2, v2)
        case "a":
            switch 329 {
            case 32767:
                var v12 int8 = 117
                _ = v12
            case 4294967295:
                fmt.Println(v9, v9)
            case 552:
                v9 = (v9 % ((65535) | 1))
            }
        }
    }
    switch ("\n\t" + "a") {
    case "a":
        for fi13 := 0; fi13 < 2; fi13++ {
            if (!false) {
                if (!false) {
                    var v14 int = (fi13 + fi13)
                    _ = v14
                }
                fmt.Println(v5)
                v2 -= (v2 << ((v9) & 63))
            } else {
                switch fi13 {
                case 0:
                    for fi15 := 0; fi15 < 6; fi15++ {
                        var v16 uint16 = 1
                        _ = fi15
                        _ = v16
                    }
                case 2147483647:
                    for fi17 := 0; fi17 < 1; fi17++ {
                        var v18 int64 = 0
                        _ = fi17
                        _ = v18
                    }
                case 42:
                    if (-2.25 < 0.1) {
                        v5 -= (v5 ^ 1000)
                    } else {
                        fmt.Println(fi13)
                        switch fi13 {
                        case 4294967295:
                            switch fi13 {
                            case -1:
                                v2 = v5
                            case 1:
                                for fi19 := 0; fi19 < 13; fi19++ {
                                    v2 = (-(v5))
                                    _ = fi19
                                }
                            default:
                                if (1 >= (v9 - v9)) {
                                    v20 := GenS1{F0: 4294967295}
                                    v21 := GenS2{F0: 255}
                                    _ = v20
                                    _ = v21
                                }
                            }
                        case -2147483648:
                            delete(v3, (fi13 ^ 42))
                        case 0:
                            v22 := GenS3{F0: 516, F1: 2, F2: 2147483648}
                            _ = v22
                        }
                        v5 -= v5
                    }
                }
            }
            switch 820 {
            case -1:
                var v23 int8 = 22
                _ = v23
            case 2147483647:
                fmt.Println(v2, v2)
            }
            _ = fi13
        }
    case "\n\t":
        var v24 int8 = 56
        _ = v24
    case "":
        fmt.Println(v2, v9)
    }
    switch -9223372036854775808 {
    case 127:
        v25 := v3[32768]
        _ = v25
    case -610:
        switch ("a" + ("foo bar" + ("foo bar" + ("foo bar" + "")))) {
        case "":
            if false {
                v5 = v2
                for fi26 := 0; fi26 < 9; fi26++ {
                    if true {
                        var v27 int64 = 1
                        v28 := GenS2{F0: (v9 | 1)}
                        _ = v27
                        _ = v28
                    }
                    fmt.Println(fi26, v9)
                    v1[fi26] = fi26
                    _ = fi26
                }
                v9 = v9
            }
        }
    case -1:
        for fi29 := 0; fi29 < 7; fi29++ {
            var v30 uint32 = 1
            var v31 uint32 = (^(v30))
            delete(v1, fi29)
            _ = fi29
            _ = v30
            _ = v31
        }
    default:
        var v32 int16 = v2
        _ = v32
    }
    v5 = -1
    fmt.Println(v2, v5)
    v33 := GenS4{F0: 128, F1: 530}
    for fi34 := 0; fi34 < 13; fi34++ {
        for fi35 := 0; fi35 < 13; fi35++ {
            if ((-1 > 42) && ("foo bar" < "foo bar")) {
                fmt.Println(fi34, v9, fi34)
            }
            if (65535 <= 32767) {
                var v36 float32 = 3.14159
                _ = v36
            } else {
                var v37 uint = (v9 << (((v9 << (((^(v9))) & 63))) & 63))
                _ = v37
            }
            if (-36 >= -1) {
                v38 := v3[-2147483648]
                _ = v38
            } else {
                v2 += (v2 % ((v2) | 1))
            }
            _ = fi35
        }
        for fi39 := 0; fi39 < 15; fi39++ {
            var v40 uint = v9
            var v41 int8 = 100
            v42 := GenS1{F0: 278}
            _ = fi39
            _ = v40
            _ = v41
            _ = v42
        }
        var v43 float64 = ((float64((((1e10 * -2.25) / ((((0.0 / 3.14159) + ((float32(3.14159) - float32(((-2.25 / 0.1) * 1e10))) - 0.0)) * 1e10) + 1.5)) + 0.1)) - 0.1) * -2.25)
        _ = fi34
        _ = v43
    }
    v9 = (v9 << ((v9) & 63))
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
