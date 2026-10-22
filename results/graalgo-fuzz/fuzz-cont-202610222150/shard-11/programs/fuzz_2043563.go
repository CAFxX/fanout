package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 int16 }

func main() {
    var v1 string = "\n\t"
    var v2 int = -1
    v3 := make(map[int]string)
    var v4 int64 = -1
    v5 := []uint32{42, 4294967295}
    if len(v5) > 0 {
        v5[0] = 304
    }
    _ = len(v5)
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint32 = 0
        if true {
            var v8 uint = 2
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    if false {
        v9 := GenS1{F0: v4, F1: 520}
        v10 := GenS1{F0: (v4 * v4), F1: 832}
        _ = v10
        _ = v9
    } else {
        fmt.Println(v1)
        var v11 int64 = v4
        _ = v11
    }
    if false {
        for fi12 := 0; fi12 < 9; fi12++ {
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 int32 = 0
                v15 := GenS2{F0: (v4 & int64(v14))}
                _ = fi13
                _ = v14
                _ = v15
            }
            v1 = v1
            v16 := GenS2{F0: (v4 + 9223372036854775807)}
            _ = fi12
            _ = v16
        }
        var v17 uint16 = 1
        v18 := GenS2{F0: v4}
        _ = v17
        _ = v18
    }
    var v19 int32 = 42
    switch v2 {
    case 4294967295:
        _ = len(v3)
    case 2147483647:
        switch "" {
        case "":
            fmt.Println(v19)
        case "a":
            v20 := GenS2{F0: 350}
            _ = v20
        case "foo bar":
            v5 = append(v5, 2147483648)
        default:
            if (((1e10 / (0.1 - (1.5 - (3.14159 - 3.14159)))) / 1.0) > (((-2.25 * -2.25) + 1e10) + float64((((3.14159 - 0.1) - 1.5) / ((0.1 - ((1.5 * 1.5) / ((1e10 - float32(0.1)) * (0.1 * ((1.5 - ((float32(0.1) - float32(-2.25)) - 0.0)) + 0.1))))) - (float32(1.5) * -2.25)))))) {
                if true {
                    for fi21 := 0; fi21 < 16; fi21++ {
                        v2 -= (fi21 & (v2 >> ((255) & 63)))
                        _ = fi21
                    }
                    var v22 bool = (2 <= 32767)
                    _ = v22
                }
            } else {
                var v23 uint = 4294967295
                var v24 float32 = (0.1 - -2.25)
                _ = v23
                _ = v24
            }
        }
    default:
        for fi25 := 0; fi25 < 13; fi25++ {
            switch v1 {
            case "hello":
                _ = len(v5)
            }
            _ = fi25
        }
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 uint64 = 18446744073709551615
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 uint = 65535
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi28
            _ = v29
        }
        fmt.Println(fi26, v27)
        _ = fi26
        _ = v27
    }
    fmt.Println(v19, v2, v1)
    v30 := GenS3{F0: 127, F1: 32767}
    if (!true) {
        if len(v5) > 0 {
            v5[0] = 1
        }
        v2 -= v2
    } else {
        var v31 int32 = (v19 >> ((899) & 63))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v31
    }
    v1 = v1
    var v32 float64 = 1.5
    fmt.Println(v19)
    var v33 int = (-(v2))
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
