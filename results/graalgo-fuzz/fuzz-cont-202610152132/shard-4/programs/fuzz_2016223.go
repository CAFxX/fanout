package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 float64 = ((1.5 - 3.14159) * 1e10)
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 uint32 = 576
        var v6 int32 = 1
        var v7 int16 = 1000
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    switch 405 {
    case 329:
        for fi8 := 0; fi8 < 7; fi8++ {
            if (149 >= 300) {
                if (!(!true)) {
                    switch (("" + "hello") + ("\n\t" + "hello")) {
                    case "\n\t":
                        switch (fi8 ^ 127) {
                        case -9223372036854775808:
                            v3 = v3
                        default:
                            v3 = (v3 * v3)
                        }
                    }
                    var v9 uint32 = 1
                    for fi10 := 0; fi10 < 19; fi10++ {
                        var v11 int64 = 1
                        switch ("" + (("foo bar" + "foo bar") + "foo bar")) {
                        case "a":
                            fmt.Println(v11, v3)
                        case "hello":
                            v3 += (0.0 / v3)
                            v11 -= (v11 / ((v11) | 1))
                        default:
                            switch fi8 {
                            case 747:
                                var v12 uint = 49
                                _ = v12
                            case -2147483648:
                                fmt.Println(v11, v11, fi10)
                            case 32768:
                                switch fi8 {
                                case 32768:
                                    v3 -= 1.5
                                case 181:
                                    fmt.Println(v11, v3, fi8)
                                case 12:
                                    fmt.Println(v3)
                                default:
                                    delete(v2, fi8)
                                }
                            }
                        }
                        _ = fi10
                        _ = v11
                    }
                    _ = v9
                }
                delete(v2, (fi8 << ((255) & 63)))
            } else {
                var v13 int16 = -1
                v3 += v3
                _ = v13
            }
            var v14 float64 = (float64(-2.25) / ((float64(float32(0.0)) * (v3 + float64(((((3.14159 + (3.14159 - (0.0 + -2.25))) + 1.5) / 1e10) - 0.1)))) / v3))
            v15 := GenS1{F0: (fi8 << ((65535) & 63))}
            _ = fi8
            _ = v14
            _ = v15
        }
    case 158:
        v16 := v2[-2147483648]
        _ = v16
    case 743:
        v17 := GenS1{F0: 127}
        _ = v17
    }
    if true {
        v18 := []int{32767, -481, 32768}
        switch "" {
        case "foo bar":
            fmt.Println(v3)
        case "":
            v3 -= 0.1
        default:
            var v19 uint64 = 132
            _ = v19
        }
        if (-1 != -32768) {
            v3 = v3
            v20 := v1[284]
            _ = v20
        } else {
            v21 := GenS1{F0: -99}
            v3 += v3
            fmt.Println(v3)
            _ = v21
        }
        _ = v18
    }
    v22 := GenS1{F0: 1}
    fmt.Println(v3)
    v3 = v3
    var v23 int32 = 1
    switch 32767 {
    case -9223372036854775808:
        v23 += (v23 / ((v23) | 1))
    case -1:
        switch 32767 {
        case 9223372036854775807:
            var v24 int = 128
            _ = v24
        case 128:
            v23 = -133
        default:
            var v25 float64 = ((v3 + ((v3 - ((v3 - v3) / 1e10)) + v3)) + (v3 * (3.14159 * v3)))
            _ = v25
        }
    case 127:
        switch "\n\t" {
        case "\n\t":
            for fi26 := 0; fi26 < 12; fi26++ {
                v23 -= v23
                _ = fi26
            }
        case "hello":
            for fi27 := 0; fi27 < 3; fi27++ {
                var v28 uint16 = 2
                v28 += 154
                _ = fi27
                _ = v28
            }
        case "foo bar":
            v29 := GenS2{F0: 0, F1: 42}
            _ = v29
        }
    default:
        var v30 int32 = (v23 / (((v23 & -28)) | 1))
        _ = v30
    }
    v31 := GenS3{F0: -1}
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v31
}
