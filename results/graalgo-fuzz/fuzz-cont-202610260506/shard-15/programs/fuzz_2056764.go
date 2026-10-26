package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint; F2 uint64 }

func main() {
    v1 := []uint32{2147483648, 4294967295, 317}
    v2 := []uint32{2147483647, 1, 0}
    v3 := make(map[int]string)
    var v4 uint = 42
    var v5 uint8 = 1
    var v6 uint16 = 98
    v7 := GenS1{F0: 127}
    switch "" {
    case "hello":
        fmt.Println(v5, v6)
    case "foo bar":
        v8 := GenS2{F0: v5, F1: v5, F2: 255}
        _ = v8
    }
    _ = len(v2)
    v9 := GenS1{F0: 100}
    v10 := GenS2{F0: v5, F1: v5, F2: (v4 % ((v4) | 1))}
    for fi11 := 0; fi11 < 14; fi11++ {
        v6 = v6
        v3[128] = "a"
        _ = fi11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        if true {
            if false {
                var v13 uint8 = (v5 & (v5 - (v5 | 0)))
                switch "a" {
                case "foo bar":
                    for fi14 := 0; fi14 < 14; fi14++ {
                        if (!(true && ((((1.5 + ((0.0 - -2.25) - (1.5 / 1.0))) + (1.5 / 0.1)) * 1e10) > 3.14159))) {
                            v3[(fi12 ^ (-(fi12)))] = "hello"
                            if false {
                                var v15 float32 = (float32((0.1 - 3.14159)) / (float32(1e10) - 0.0))
                                _ = v15
                            }
                        }
                        v6 += v6
                        fmt.Println(v6, v4, fi12)
                        if len(v2) > 0 {
                            v2[0] = 527
                        }
                        _ = fi14
                    }
                default:
                    if (v6 <= 444) {
                        var v16 uint32 = 4294967295
                        v17 := GenS3{F0: v16}
                        _ = v16
                        _ = v17
                    } else {
                        for fi18 := 0; fi18 < 2; fi18++ {
                            switch "foo bar" {
                            case "":
                                if (!(-2.25 <= 1e10)) {
                                    var v19 int16 = 42
                                    v2 = append(v2, 0)
                                    _ = v19
                                }
                            }
                            v6 -= (v6 << ((136) & 63))
                            var v20 uint32 = 0
                            _ = fi18
                            _ = v20
                        }
                    }
                }
                v21 := GenS4{F0: fi12, F1: v4, F2: 4294967295}
                _ = v13
                _ = v21
            }
            var v22 bool = (2147483648 > 798)
            _ = v22
        } else {
            fmt.Println(fi12, v4)
        }
        _ = fi12
    }
    v5 -= (v5 & (v5 ^ v5))
    v6 += 152
    v1 = append(v1, uint32(v5))
    fmt.Println(v4, v6)
    fmt.Println(v6, v6, v6)
    v6 += v6
    if true {
        v23 := GenS3{F0: 730}
        _ = v23
    } else {
        switch 0 {
        case 486:
            var v24 float64 = float64(0.0)
            _ = v24
        case -76:
            delete(v3, 0)
        case 4294967295:
            fmt.Println(v4, v6)
        }
        for fi25 := 0; fi25 < 14; fi25++ {
            var v26 uint16 = v6
            var v27 string = "\n\t"
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
