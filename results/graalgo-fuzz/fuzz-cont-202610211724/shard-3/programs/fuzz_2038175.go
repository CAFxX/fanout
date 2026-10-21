package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int16 }

func main() {
    var v1 uint8 = 2
    var v2 int32 = -2147483648
    var v3 uint = 851
    var v4 uint16 = 1
    v5 := make(map[int]string)
    v6 := GenS1{F0: ((-790 >= -181) && false), F1: 1, F2: 2}
    v7 := []int{}
    v3 += 1
    v3 += 849
    switch 32768 {
    case -9223372036854775808:
        if true {
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 int64 = -1
                var v10 bool = (42 > 1)
                _ = fi8
                _ = v10
                _ = v9
            }
            fmt.Println(v4, v3, v4)
            if true {
                fmt.Println(v2, v4)
            } else {
                v11 := GenS1{F0: false, F1: 100, F2: 18446744073709551615}
                if (!((v3 >= v3) && (0.1 <= ((1e10 * (1e10 / (1.5 + (1.5 - 1e10)))) * float32((0.0 + float64(1e10))))))) {
                    v12 := GenS1{F0: (true || true), F1: 95, F2: 42}
                    fmt.Println(v2, v2, v3)
                    _ = v12
                } else {
                    var v13 float32 = 1e10
                    v13 += v13
                    _ = v13
                }
                _ = v11
            }
        } else {
            for fi14 := 0; fi14 < 7; fi14++ {
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 uint = (v3 - (v3 - v3))
                    if (((v2 >> (((v3 ^ (^(v3)))) & 63)) > v2) && true) {
                        var v17 uint = v16
                        v2 = v2
                        v5[fi14] = "foo bar"
                        _ = v17
                    }
                    _ = fi15
                    _ = v16
                }
                _ = fi14
            }
            v4 = 0
            for fi18 := 0; fi18 < 12; fi18++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                if (v2 >= -921) {
                    v19 := GenS1{F0: false, F1: 100, F2: 4294967295}
                    for fi20 := 0; fi20 < 6; fi20++ {
                        if true {
                            v21 := v5[(^(fi20))]
                            var v22 int16 = 32767
                            var v23 int = (fi18 * -1)
                            v2 = v2
                            var v24 int64 = -346
                            _ = v21
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            var v25 float32 = ((1e10 + 1e10) * 3.14159)
                            fmt.Println(v2)
                            _ = v25
                        }
                        if true {
                            v1 -= v1
                            v1 -= (v1 - (v1 / (((^(v1))) | 1)))
                            var v26 int64 = 9223372036854775807
                            _ = v26
                        } else {
                            v3 = (v3 << ((65535) & 63))
                            var v27 int64 = -1
                            fmt.Println(v27, v1, v4)
                            _ = v27
                        }
                        var v28 int = (fi18 ^ -585)
                        _ = fi20
                        _ = v28
                    }
                    _ = v19
                } else {
                    v1 = 2
                    var v29 uint8 = uint8(v3)
                    _ = v29
                }
                _ = fi18
            }
        }
    case 9223372036854775807:
        if false {
            if false {
                switch "" {
                case "hello":
                    var v30 uint16 = (v4 << ((65535) & 63))
                    _ = v30
                default:
                    v31 := GenS1{F0: (0 < uint64(v4)), F1: 116, F2: 0}
                    _ = v31
                }
                v32 := GenS2{F0: (v2 - 2147483647), F1: 619}
                _ = v32
            } else {
                v33 := GenS3{F0: v1, F1: -1}
                v1 = 2
                v4 += v4
                _ = v33
            }
            v1 -= (v1 >> (((^(v3))) & 63))
        } else {
            v34 := GenS2{F0: (v2 / ((v2) | 1)), F1: 0}
            v1 = 38
            switch "a" {
            case "foo bar":
                var v35 int8 = -1
                _ = v35
            case "a":
                if (!((v2 - v2) == v2)) {
                    v2 -= v2
                    fmt.Println(v1, v2, v3)
                    v1 -= 109
                } else {
                    switch 676 {
                    case 836:
                        v1 = v1
                    case 0:
                        v4 = v4
                    case 128:
                        var v36 int32 = 2147483647
                        _ = v36
                    }
                    v37 := GenS2{F0: 0, F1: 433}
                    _ = v37
                }
            default:
                var v38 float32 = float32(1e10)
                _ = v38
            }
            _ = v34
        }
    }
    _ = len(v7)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v39 uint64 = 42
    switch ("hello" + "hello") {
    case "\n\t":
        v7 = append(v7, -398)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
