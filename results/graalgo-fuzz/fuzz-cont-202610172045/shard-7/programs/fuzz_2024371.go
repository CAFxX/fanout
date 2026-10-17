package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 string }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int; F1 int8; F2 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    v3 := []uint32{}
    var v4 uint16 = 517
    var v5 uint16 = (v4 / ((v4) | 1))
    var v6 uint16 = 32767
    var v7 uint = 6
    for fi8 := 0; fi8 < 19; fi8++ {
        v6 += v6
        if true {
            var v9 string = "a"
            _ = v9
        } else {
            switch "" {
            case "":
                v4 = v6
            case "\n\t":
                if true {
                    v10 := GenS1{F0: -1, F1: 1, F2: (("foo bar" + "") + "foo bar")}
                    _ = v10
                }
            case "a":
                v6 -= v5
            }
            v6 += 255
        }
        _ = fi8
    }
    v5 += v4
    for fi11 := 0; fi11 < 10; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            if false {
                v5 -= v5
                v5 += (v6 >> (((v7 + (v7 >> (((v7 ^ (v7 << (((v7 ^ (v7 & (v7 / ((v7) | 1))))) & 63)))) & 63)))) & 63))
            }
            v13 := GenS2{F0: 774, F1: 42}
            v5 = (v4 << ((v7) & 63))
            _ = fi12
            _ = v13
        }
        fmt.Println(v4, v7)
        _ = fi11
    }
    fmt.Println(v6, v6, v4)
    v6 -= 2
    switch 128 {
    case -285:
        if false {
            if (!(-2.25 > (float32(-2.25) + 1.5))) {
                for fi14 := 0; fi14 < 8; fi14++ {
                    v4 -= 561
                    v15 := GenS1{F0: -1, F1: 9223372036854775807, F2: ""}
                    if (!(false && true)) {
                        var v16 float32 = (((float32(-2.25) + 1e10) + 0.1) / 0.1)
                        var v17 uint = v7
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                }
            }
            v4 += (v4 + v5)
            v5 -= v4
        }
    case -241:
        for fi18 := 0; fi18 < 16; fi18++ {
            for fi19 := 0; fi19 < 14; fi19++ {
                v7 += (v7 * v7)
                v5 = 1
                _ = fi19
            }
            v7 = (^(v7))
            _ = fi18
        }
    default:
        var v20 int = -259
        _ = v20
    }
    v21 := GenS2{F0: 0, F1: 42}
    v22 := GenS2{F0: 372, F1: -580}
    v1[-2147483648] = "foo bar"
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 uint8 = uint8(v7)
        _ = fi23
        _ = v24
    }
    var v25 uint8 = 2
    for fi26 := 0; fi26 < 1; fi26++ {
        fmt.Println(v6)
        _ = fi26
    }
    v27 := GenS3{F0: 42, F1: 30, F2: -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
