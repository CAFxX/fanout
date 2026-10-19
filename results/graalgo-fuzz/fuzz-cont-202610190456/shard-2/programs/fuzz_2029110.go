package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 int8; F2 uint8 }

func main() {
    var v1 int16 = -690
    v2 := []string{"foo bar", ""}
    var v3 int = 42
    var v4 float32 = (float32(1e10) + -2.25)
    var v5 uint64 = 553
    var v6 string = ""
    v1 = (v1 / ((v1) | 1))
    v3 = (^(v3))
    if (2 >= 0) {
        switch v6 {
        case "\n\t":
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 uint8 = 0
                _ = fi7
                _ = v8
            }
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = "a"
            }
        case "a":
            if true {
                var v9 uint8 = 166
                _ = v9
            }
        default:
            var v10 uint64 = 2
            _ = v10
        }
        for fi11 := 0; fi11 < 19; fi11++ {
            v2 = append(v2, (v6 + "foo bar"))
            _ = fi11
        }
    } else {
        var v12 bool = (!(53 <= 89))
        switch v3 {
        case 127:
            v6 = v6
        case 1:
            var v13 int32 = 0
            _ = v13
        case -423:
            _ = len(v2)
        default:
            if v12 {
                v3 = (v3 >> ((1) & 63))
            }
        }
        _ = v12
    }
    v6 = v6
    var v14 uint = 18446744073709551615
    var v15 int16 = (^(v1))
    v14 = (v14 << (((^(v14))) & 63))
    switch (("a" + v6) + ("\n\t" + (v6 + v6))) {
    case "a":
        fmt.Println(v6, v4, v6)
    case "foo bar":
        var v16 float32 = (((-2.25 / 1e10) / (0.0 - v4)) + -2.25)
        _ = v16
    default:
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := GenS1{F0: v5, F1: 42}
            fmt.Println(v15, v4, v1)
            _ = fi17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 string = ""
            var v22 uint32 = 688
            _ = fi20
            _ = v21
            _ = v22
        }
        v5 = 92
        switch fi19 {
        case 42:
            v3 -= v3
        case -2147483648:
            v14 -= 0
        case 32767:
            v3 = v3
        default:
            v23 := GenS2{F0: false}
            _ = v23
        }
        _ = fi19
    }
    v14 += v14
    fmt.Println(v5, v15)
    v24 := GenS3{F0: 42, F1: 102, F2: 226}
    v1 = (v1 & -1)
    fmt.Println(v1, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
