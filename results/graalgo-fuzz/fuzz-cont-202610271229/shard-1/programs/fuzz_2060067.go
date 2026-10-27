package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 431
    var v3 string = "foo bar"
    v4 := []int{2147483647, -381, 127}
    fmt.Println(v2, v2)
    var v5 int16 = 442
    for fi6 := 0; fi6 < 15; fi6++ {
        switch (^(fi6)) {
        case -1:
            if (("" + v3) < "hello") {
                v1[973] = fi6
                v1[2147483648] = 128
                var v7 int = fi6
                _ = v7
            }
        case -286:
            v8 := GenS1{F0: 4294967295, F1: (!(255 <= 146)), F2: 42}
            _ = v8
        case 1:
            for fi9 := 0; fi9 < 11; fi9++ {
                switch fi6 {
                case 4294967295:
                    v10 := GenS2{F0: 623}
                    _ = v10
                default:
                    var v11 uint64 = uint64(v2)
                    _ = v11
                }
                _ = fi9
            }
        default:
            for fi12 := 0; fi12 < 17; fi12++ {
                v3 = (("\n\t" + "a") + v3)
                if true {
                    v5 -= 0
                    var v13 int8 = -1
                    v3 = v3
                    _ = v13
                }
                _ = fi12
            }
        }
        fmt.Println(v3)
        _ = fi6
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        v5 = 428
        if (!(-231 == -1)) {
            for fi15 := 0; fi15 < 15; fi15++ {
                delete(v1, 2)
                var v16 float32 = 0.0
                _ = fi15
                _ = v16
            }
            v17 := GenS2{F0: 0}
            _ = v17
        } else {
            var v18 uint8 = 0
            v19 := GenS1{F0: 792, F1: ((fi14 << ((4294967295) & 63)) > fi14), F2: -128}
            _ = v18
            _ = v19
        }
        _ = len(v4)
        _ = fi14
    }
    switch 32767 {
    case 32768:
        switch 2147483647 {
        case 281:
            fmt.Println(v3)
        case 778:
            switch 128 {
            case -2147483648:
                if true {
                    v20 := GenS3{F0: -1}
                    _ = v20
                }
            default:
                var v21 uint16 = 42
                _ = v21
            }
        default:
            for fi22 := 0; fi22 < 16; fi22++ {
                var v23 uint16 = 2
                fmt.Println(v5)
                fmt.Println(v2, v2, v5)
                _ = fi22
                _ = v23
            }
        }
    case 946:
        var v24 uint16 = 2
        _ = v24
    case 42:
        fmt.Println(v3, v5, v2)
    default:
        v3 = v3
    }
    if len(v4) > 0 {
        v4[0] = 156
    }
    v5 -= v5
    v5 = 42
    delete(v1, uint32(v2))
    fmt.Println(v3, v5, v3)
    for fi25 := 0; fi25 < 15; fi25++ {
        fmt.Println(v5, v2)
        fmt.Println(fi25)
        delete(v1, 2)
        _ = fi25
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        v2 = 18446744073709551615
        if true {
            var v27 string = v3
            v5 -= (v5 * (v5 & v5))
            _ = v27
        } else {
            if len(v4) > 0 {
                v4[0] = -9223372036854775808
            }
            if len(v4) > 0 {
                v4[0] = (fi26 | (fi26 * (fi26 * fi26)))
            }
        }
        v28 := GenS1{F0: 0, F1: false, F2: 100}
        _ = fi26
        _ = v28
    }
    v5 += (v5 + v5)
    v29 := GenS1{F0: 42, F1: false, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
