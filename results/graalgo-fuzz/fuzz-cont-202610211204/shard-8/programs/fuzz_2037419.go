package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS4 struct { F0 int8; F1 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    var v3 uint16 = 430
    v4 := make(map[uint32]int)
    v5 := []uint32{2147483648}
    v6 := make(map[string]string)
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 string = "\n\t"
        _ = fi7
        _ = v8
    }
    delete(v6, "\n\t")
    if ((0.1 - 1e10) >= (1e10 + (-2.25 + (((1.5 - -2.25) * 0.0) - float32((0.1 - float64(1.5))))))) {
        var v9 int16 = 570
        delete(v2, -1)
        if (0 <= -1) {
            var v10 int = 128
            var v11 int = v10
            v12 := GenS1{F0: (v10 % ((2147483647) | 1))}
            _ = v10
            _ = v11
            _ = v12
        }
        _ = v9
    }
    v13 := v6[((("hello" + "foo bar") + "hello") + "a")]
    v3 -= 0
    switch (v13 + (v13 + "\n\t")) {
    case "foo bar":
        switch -789 {
        case 127:
            if false {
                for fi14 := 0; fi14 < 16; fi14++ {
                    v4[2] = (fi14 - fi14)
                    _ = fi14
                }
            } else {
                var v15 int16 = 1000
                if (!(((0.1 * ((0.1 * 1e10) - (-2.25 - (-2.25 * -2.25)))) / 1e10) <= (((3.14159 * float64(1.5)) / (0.1 / 1.5)) / (1.5 * 0.1)))) {
                    for fi16 := 0; fi16 < 6; fi16++ {
                        for fi17 := 0; fi17 < 11; fi17++ {
                            if (-1 > 0) {
                                var v18 float64 = ((-2.25 + 3.14159) * 3.14159)
                                if false {
                                    v3 = v3
                                    var v19 float32 = (float32(v18) * -2.25)
                                    v19 -= float32(v18)
                                    _ = v19
                                } else {
                                    v20 := GenS2{F0: v13}
                                    _ = v20
                                }
                                _ = v18
                            }
                            var v21 uint = 0
                            _ = fi17
                            _ = v21
                        }
                        _ = fi16
                    }
                    var v22 uint16 = (v3 << ((301) & 63))
                    _ = v22
                }
                _ = v15
            }
        default:
            if (v13 == v13) {
                fmt.Println(v3, v13)
                if len(v5) > 0 {
                    v5[0] = 42
                }
                var v23 uint16 = (v3 >> ((76) & 63))
                _ = v23
            } else {
                for fi24 := 0; fi24 < 4; fi24++ {
                    if (fi24 >= 288) {
                        v25 := GenS3{F0: 2, F1: -9223372036854775808, F2: 255}
                        var v26 int8 = 30
                        _ = v25
                        _ = v26
                    } else {
                        _ = len(v5)
                        v27 := GenS1{F0: fi24}
                        v3 -= v3
                        _ = v27
                    }
                    _ = len(v5)
                    v28 := v2[fi24]
                    _ = fi24
                    _ = v28
                }
                var v29 uint64 = 416
                _ = v29
            }
        }
    case "\n\t":
        v5 = append(v5, 4294967295)
    case "a":
        var v30 int64 = 9223372036854775807
        _ = v30
    default:
        var v31 string = v13
        _ = v31
    }
    _ = len(v5)
    if true {
        var v32 uint16 = 2
        var v33 int8 = 42
        var v34 bool = false
        _ = v32
        _ = v33
        _ = v34
    } else {
        v35 := GenS4{F0: 127, F1: -594}
        _ = v35
    }
    if true {
        fmt.Println(v3)
        var v36 int = 32767
        var v37 string = "\n\t"
        _ = v36
        _ = v37
    } else {
        for fi38 := 0; fi38 < 11; fi38++ {
            for fi39 := 0; fi39 < 14; fi39++ {
                var v40 uint16 = v3
                var v41 uint32 = 42
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = fi38
        }
        switch v13 {
        case "a":
            v3 += (v3 | 584)
        case "":
            for fi42 := 0; fi42 < 8; fi42++ {
                var v43 int = fi42
                for fi44 := 0; fi44 < 17; fi44++ {
                    var v45 float64 = -2.25
                    _ = fi44
                    _ = v45
                }
                _ = fi42
                _ = v43
            }
        default:
            var v46 float64 = 1.5
            _ = v46
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
