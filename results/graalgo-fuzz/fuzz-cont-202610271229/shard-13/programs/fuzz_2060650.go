package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint8 }

func main() {
    var v1 uint64 = 895
    var v2 int8 = 0
    var v3 uint = 390
    v4 := []int{}
    if (3.14159 == (0.0 - (((float32(1.5) + (-2.25 * float32(3.14159))) + 0.1) - 1.5))) {
        var v5 int32 = 0
        _ = v5
    }
    v6 := GenS1{F0: -9223372036854775808, F1: (v1 % ((v1) | 1)), F2: 2147483648}
    v1 = (v1 + v1)
    for fi7 := 0; fi7 < 12; fi7++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch "" {
        case "a":
            v3 -= (v3 / (((v3 | 255)) | 1))
        case "hello":
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v2, v2, v1)
            }
        case "\n\t":
            var v8 uint64 = 4294967295
            _ = v8
        default:
            switch "foo bar" {
            case "":
                fmt.Println(v1, v3, fi7)
            case "hello":
                if len(v4) > 0 {
                    v4[0] = 143
                }
            case "\n\t":
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 int16 = -1
                    v3 = v3
                    if len(v4) > 0 {
                        v4[0] = (fi7 / ((fi7) | 1))
                    }
                    _ = fi9
                    _ = v10
                }
            }
        }
        var v11 uint16 = 42
        _ = fi7
        _ = v11
    }
    v1 = v1
    if false {
        v12 := GenS1{F0: int(v2), F1: (v1 % (((v1 ^ 331)) | 1)), F2: 2147483648}
        var v13 uint16 = 66
        _ = v12
        _ = v13
    } else {
        v3 = 777
        var v14 bool = false
        v15 := GenS1{F0: 1, F1: (v1 << (((v3 >> (((v3 ^ (^(v3)))) & 63))) & 63)), F2: 42}
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 int64 = 1
            var v19 int8 = (^(v2))
            if (0 > v18) {
                v1 = (v1 % ((v1) | 1))
                switch (fi17 >> ((771) & 63)) {
                case 735:
                    switch fi16 {
                    case 555:
                        v18 -= (v18 & -1)
                    default:
                        fmt.Println(fi17)
                    }
                case 128:
                    v18 = (v18 * (v18 - -9223372036854775808))
                case 32768:
                    v20 := GenS2{F0: 1, F1: 1}
                    _ = v20
                default:
                    var v21 int16 = 252
                    _ = v21
                }
                fmt.Println(v3, v1)
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    v3 = (v3 | v3)
    v1 = (v1 >> (((v3 << (((^(v3))) & 63))) & 63))
    v4 = append(v4, 142)
    fmt.Println(v3)
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        var v22 uint16 = 361
        _ = v22
    }
    v23 := GenS2{F0: int16(v2), F1: uint8(v1)}
    var v24 uint64 = (v1 - (v1 << (((v3 % ((v3) | 1))) & 63)))
    v1 = 962
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v6
}
