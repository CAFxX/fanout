package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS2 struct { F0 int32; F1 int; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 int16 }

func main() {
    v1 := []uint32{679}
    var v2 bool = (244 < 0)
    var v3 int = -9223372036854775808
    var v4 uint32 = 2
    v5 := []uint32{v4, 1}
    v6 := []uint32{v4, (v4 - v4)}
    var v7 uint32 = (v4 + v4)
    for fi8 := 0; fi8 < 15; fi8++ {
        v2 = (v2 && (2 >= 111))
        var v9 uint16 = 746
        v10 := GenS1{F0: 2147483647, F1: 0, F2: (v9 + v9)}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: 160, F1: v3, F2: (v7 ^ (^(v7)))}
    for fi12 := 0; fi12 < 19; fi12++ {
        if false {
            var v13 float64 = float64(0.1)
            switch "a" {
            case "a":
                if (!(!false)) {
                    switch (-(fi12)) {
                    case 127:
                        v3 = (^(v3))
                    case 42:
                        if v2 {
                            v6 = append(v6, v7)
                        }
                    }
                } else {
                    _ = len(v1)
                    var v14 bool = (42 == 558)
                    _ = v14
                }
            case "foo bar":
                v7 = v4
            default:
                if ((float64((3.14159 / 1e10)) - 3.14159) > 1e10) {
                    fmt.Println(fi12, v3)
                }
            }
            for fi15 := 0; fi15 < 6; fi15++ {
                v7 += (v7 << ((2) & 63))
                if ((255 != 0) && false) {
                    fmt.Println(fi12)
                    switch (v3 / ((fi15) | 1)) {
                    case 127:
                        switch 2147483647 {
                        case 4294967295:
                            fmt.Println(fi15)
                        case -9223372036854775808:
                            v16 := GenS1{F0: 2147483647, F1: 0, F2: 1}
                            _ = v16
                        case 498:
                            v4 = (v7 >> ((255) & 63))
                        default:
                            switch "a" {
                            case "\n\t":
                                v3 += fi15
                            case "foo bar":
                                v13 += ((v13 + 1e10) / v13)
                                v5 = append(v5, v4)
                            default:
                                v3 += -358
                            }
                            if v2 {
                                var v17 int8 = 70
                                v4 = (v7 + (v4 << ((4294967295) & 63)))
                                _ = v17
                            }
                            v3 += 1
                        }
                    case 778:
                        v18 := GenS2{F0: -1, F1: v3, F2: 0}
                        _ = v18
                    case 2147483647:
                        switch "hello" {
                        case "":
                            v7 += v4
                        default:
                            v4 = (v4 - v7)
                        }
                    }
                }
                var v19 string = "a"
                _ = fi15
                _ = v19
            }
            _ = v13
        }
        switch "" {
        case "a":
            var v20 int8 = -128
            _ = v20
        }
        _ = fi12
    }
    v2 = v2
    switch ((("hello" + "foo bar") + (("" + "hello") + "")) + ("hello" + "a")) {
    case "hello":
        var v21 float32 = (float32((((0.1 + 0.0) / -2.25) / 1.0)) - 1e10)
        _ = v21
    case "foo bar":
        v22 := GenS3{F0: uint8(v7), F1: 1000}
        _ = v22
    default:
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 float32 = (0.1 / 1.0)
            _ = len(v5)
            _ = fi23
            _ = v24
        }
    }
    v25 := GenS4{F0: v2, F1: 2}
    for fi26 := 0; fi26 < 1; fi26++ {
        fmt.Println(fi26, v2)
        var v27 int64 = 9223372036854775807
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        fmt.Println(v4, v7, v2)
        var v29 float64 = (3.14159 / (1e10 / ((3.14159 - ((((0.1 - -2.25) + (0.1 / (1e10 * float64(0.1)))) * 1e10) / 1.5)) + 0.0)))
        v30 := GenS3{F0: 1, F1: 691}
        _ = fi28
        _ = v29
        _ = v30
    }
    v31 := GenS5{F0: -907}
    v32 := GenS3{F0: 127, F1: 1000}
    v6 = append(v6, (v4 ^ 65535))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
