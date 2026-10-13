package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int64; F1 bool }

func main() {
    v1 := []int{32767, -114, 0}
    v2 := make(map[int]int)
    var v3 int32 = -807
    v4 := make(map[int]bool)
    var v5 int16 = -215
    fmt.Println(v5)
    if (2147483647 == 2147483647) {
        fmt.Println(v3)
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v3, v5)
    default:
        if (true || true) {
            v6 := GenS1{F0: 787, F1: -128, F2: 255}
            _ = v6
        } else {
            switch 42 {
            case -775:
                for fi7 := 0; fi7 < 12; fi7++ {
                    v3 += v3
                    var v8 float32 = 0.0
                    _ = fi7
                    _ = v8
                }
            case 128:
                v5 -= v5
            }
            for fi9 := 0; fi9 < 11; fi9++ {
                v3 += v3
                for fi10 := 0; fi10 < 18; fi10++ {
                    v3 = v3
                    _ = fi10
                }
                var v11 uint32 = 2147483647
                _ = fi9
                _ = v11
            }
            if (v3 != (v3 / ((v3) | 1))) {
                switch int(v5) {
                case -9223372036854775808:
                    switch "a" {
                    case "":
                        _ = len(v2)
                    case "hello":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                default:
                    v5 = v5
                }
            } else {
                switch "" {
                case "hello":
                    for fi12 := 0; fi12 < 5; fi12++ {
                        delete(v2, fi12)
                        _ = fi12
                    }
                }
                v5 = v5
                switch -829 {
                case 127:
                    if len(v1) > 0 {
                        v1[0] = 0
                    }
                case -20:
                    v3 = v3
                case -599:
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                default:
                    for fi13 := 0; fi13 < 19; fi13++ {
                        for fi14 := 0; fi14 < 16; fi14++ {
                            switch "hello" {
                            case "\n\t":
                                if (4294967295 < 42) {
                                    switch "foo bar" {
                                    case "a":
                                        var v15 int32 = v3
                                        _ = v15
                                    }
                                    v5 -= (v5 >> ((2) & 63))
                                    var v16 int64 = -76
                                    if len(v1) > 0 {
                                        v1[0] = (fi13 - fi13)
                                    }
                                    _ = v16
                                }
                            case "hello":
                                if (-32768 == v5) {
                                    v17 := GenS1{F0: 2147483648, F1: 127, F2: 32767}
                                    var v18 int16 = -1
                                    switch 32768 {
                                    case -1:
                                        fmt.Println(v18)
                                    case 529:
                                        v19 := GenS2{F0: (v3 * -1), F1: 966, F2: 344}
                                        _ = v19
                                    }
                                    _ = v17
                                    _ = v18
                                }
                            case "foo bar":
                                var v20 float32 = 0.0
                                _ = v20
                            }
                            if (!false) {
                                var v21 int8 = -38
                                var v22 float64 = 0.0
                                _ = len(v1)
                                _ = v21
                                _ = v22
                            }
                            _ = fi14
                        }
                        fmt.Println(v3)
                        _ = fi13
                    }
                }
            }
        }
    }
    v3 = v3
    var v23 float64 = 1e10
    v3 = (v3 | (v3 & v3))
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        var v24 int32 = v3
        _ = v24
    case "a":
        var v25 float32 = ((1e10 * -2.25) / 1.5)
        _ = v25
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if (false || (32767 >= 1)) {
        fmt.Println(v23, v23, v23)
        for fi26 := 0; fi26 < 19; fi26++ {
            var v27 string = ""
            v28 := GenS2{F0: -943, F1: 382, F2: 1}
            v29 := GenS1{F0: 0, F1: 100, F2: 0}
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
    }
    v30 := GenS2{F0: v3, F1: 42, F2: 42}
    v3 = (v3 % ((v3) | 1))
    delete(v4, 32767)
    fmt.Println(v5)
    switch 32768 {
    case -9223372036854775808:
        fmt.Println(v3)
    default:
        var v31 uint16 = 42
        _ = v31
    }
    delete(v4, 2147483647)
    v32 := v2[2147483647]
    v33 := GenS3{F0: 490, F1: ((4294967295 != 42) || ((true || ((1 != 255) && (255 <= 32767))) && ((3.14159 + float64(3.14159)) != ((float64(float32(v23)) - ((3.14159 * ((v23 - (v23 * v23)) * v23)) - 3.14159)) - float64(-2.25)))))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
