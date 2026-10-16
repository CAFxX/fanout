package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int8; F1 int8; F2 int }
func genFn0() int8 {
    var v5 int16 = -32768
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint = 255
        for fi8 := 0; fi8 < 9; fi8++ {
            if (18 >= 42) {
                var v9 float32 = float32(0.0)
                fmt.Println(fi8)
                _ = v9
            } else {
                fmt.Println(v5)
            }
            if (1 > 2147483647) {
                v7 -= v7
                if (2147483647 >= 788) {
                    var v10 uint8 = 255
                    v5 = v5
                    v11 := GenS1{F0: (^(fi8)), F1: 460}
                    _ = v10
                    _ = v11
                } else {
                    switch (fi8 % ((fi8) | 1)) {
                    case 2147483647:
                        var v12 string = "a"
                        _ = v12
                    case 32767:
                        v5 += v5
                    case 904:
                        var v13 float32 = ((3.14159 / 1e10) / -2.25)
                        _ = v13
                    default:
                        fmt.Println(v5, fi6, v5)
                    }
                    switch fi8 {
                    case -248:
                        var v14 int8 = 42
                        _ = v14
                    case 422:
                        v7 = v7
                    case 32767:
                        v15 := GenS1{F0: fi6, F1: 9223372036854775807}
                        _ = v15
                    }
                }
            }
            _ = fi8
        }
        for fi16 := 0; fi16 < 15; fi16++ {
            fmt.Println(fi16, fi16)
            v7 = v7
            var v17 float64 = ((0.1 / (float64(((((float32((0.0 * (1.5 * 0.0))) / ((1.5 * (1.5 * 0.1)) / (1.5 / 1.0))) - 0.1) + (3.14159 / 1e10)) * -2.25)) / (1.5 - float64(0.0)))) * 0.0)
            v5 = (-(v5))
            _ = fi16
            _ = v17
        }
        _ = fi6
        _ = v7
    }
    v5 -= (v5 & (v5 / (((v5 | 32767)) | 1)))
    _ = v5
    return 0
}
func genFn1(p0 int64, p1 uint64, p2 int64) uint8 {
    v18 := GenS1{F0: 32767, F1: 9223372036854775807}
    p0 += (p0 + (p0 + 82))
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 uint8 = 2
        var v21 int = 128
        var v22 string = "hello"
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 176
}

func main() {
    var v1 uint8 = 128
    var v2 int32 = 1
    v3 := []uint32{717}
    var v4 int32 = (v2 - v2)
    v23 := GenS1{F0: 32768, F1: 1}
    fmt.Println(v4, v4)
    v3 = append(v3, 65535)
    v24 := GenS2{F0: 37, F1: 1, F2: 687}
    for fi25 := 0; fi25 < 11; fi25++ {
        v26 := GenS1{F0: fi25, F1: 0}
        switch "\n\t" {
        case "foo bar":
            if (false && false) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v27 string = "foo bar"
                switch (fi25 ^ 2147483647) {
                case 376:
                    if len(v3) > 0 {
                        v3[0] = 64
                    }
                case 0:
                    v1 += 60
                case 336:
                    var v28 int16 = 531
                    _ = v28
                }
                _ = v27
            } else {
                if (100 > 127) {
                    for fi29 := 0; fi29 < 7; fi29++ {
                        if true {
                            fmt.Println(v2, fi25)
                            var v30 int = 42
                            _ = v30
                        }
                        var v31 uint64 = 655
                        _ = fi29
                        _ = v31
                    }
                } else {
                    fmt.Println(v2, fi25)
                    if true {
                        if true {
                            var v32 bool = false
                            _ = v32
                        }
                        if (!true) {
                            v4 = (v2 ^ v2)
                            var v33 float32 = 1e10
                            switch (-(fi25)) {
                            case 9223372036854775807:
                                var v34 float64 = 0.0
                                _ = v34
                            }
                            _ = v33
                        } else {
                            var v35 uint32 = 107
                            _ = v35
                        }
                    }
                    v4 = 285
                }
            }
        case "\n\t":
            var v36 bool = true
            _ = v36
        case "a":
            var v37 string = "a"
            _ = v37
        default:
            v2 = (v2 >> ((4294967295) & 63))
        }
        v4 += v4
        _ = fi25
        _ = v26
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 += (^(v1))
    var v38 int = -360
    var v39 float32 = -2.25
    fmt.Println(v4)
    for fi40 := 0; fi40 < 16; fi40++ {
        for fi41 := 0; fi41 < 19; fi41++ {
            if (1000 >= int16(fi40)) {
                v4 = (v2 | v2)
                v1 = v1
            } else {
                v42 := GenS2{F0: 71, F1: 0, F2: fi40}
                _ = v42
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi41
        }
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
