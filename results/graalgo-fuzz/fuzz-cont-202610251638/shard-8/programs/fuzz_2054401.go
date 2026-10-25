package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -2147483648
    var v3 int8 = -1
    var v4 uint32 = 438
    var v5 uint = 42
    v4 -= v4
    if false {
        fmt.Println(v3, v3)
    }
    var v6 uint8 = 127
    v7 := GenS1{F0: 128, F1: 884, F2: v3}
    v8 := GenS1{F0: -1, F1: 32767, F2: (v3 | 0)}
    v9 := []int{42, -2147483648, 192}
    v5 -= v5
    for fi10 := 0; fi10 < 12; fi10++ {
        v2 += int32(v6)
        switch "" {
        case "foo bar":
            switch (fi10 % (((fi10 * fi10)) | 1)) {
            case 42:
                if (1 < 990) {
                    _ = len(v1)
                }
            default:
                v11 := GenS2{F0: 0}
                _ = v11
            }
        case "":
            var v12 bool = (2 <= 1)
            _ = v12
        default:
            fmt.Println(v4)
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            if (false && ((0.0 / (3.14159 - float64(((-2.25 * ((((0.1 + 1.5) / 1.5) / -2.25) + 1.5)) - ((1e10 / (((1e10 - 3.14159) - (0.1 * 1.5)) * (0.0 + 0.1))) - (float32(-2.25) - float32(-2.25))))))) >= 1e10)) {
                if (false && (9223372036854775807 > (fi13 - int(v4)))) {
                    var v14 int8 = (v3 ^ (v3 - int8(v5)))
                    v3 -= (v14 >> (((v5 | (v5 & 42))) & 63))
                    _ = v14
                }
                if ((true && true) || (false || (float32(3.14159) == 1.5))) {
                    var v15 uint = v5
                    var v16 int64 = 0
                    var v17 int64 = (v16 - (v16 - 0))
                    _ = v15
                    _ = v16
                    _ = v17
                }
            }
            _ = fi13
        }
        _ = fi10
    }
    var v18 uint16 = 255
    v19 := GenS3{F0: v5}
    if (!(-2147483648 < -1)) {
        v20 := GenS2{F0: (v18 & (v18 << (((v5 + (v5 >> ((v5) & 63)))) & 63)))}
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 int32 = (v2 / (((v2 & v2)) | 1))
            delete(v1, uint32(v6))
            _ = fi21
            _ = v22
        }
        v23 := GenS4{F0: (v5 & 4294967295)}
        _ = v20
        _ = v23
    } else {
        if (((v4 * 2) <= 4294967295) && (false || true)) {
            v3 -= (v3 - (-(v3)))
            for fi24 := 0; fi24 < 13; fi24++ {
                var v25 uint = v5
                v2 = 2147483647
                _ = fi24
                _ = v25
            }
            v1[v4] = "\n\t"
        } else {
            if ((-1 != int16(v18)) || ((1.5 + 0.0) <= 1.5)) {
                var v26 uint8 = v6
                var v27 int16 = int16(v3)
                _ = v26
                _ = v27
            }
        }
        for fi28 := 0; fi28 < 4; fi28++ {
            switch "" {
            case "\n\t":
                var v29 float64 = 0.0
                _ = v29
            case "foo bar":
                v18 += (v18 - (v18 & 0))
            case "a":
                fmt.Println(fi28)
            default:
                switch ("hello" + "\n\t") {
                case "foo bar":
                    var v30 string = ("foo bar" + "")
                    _ = v30
                case "a":
                    v31 := GenS4{F0: (v5 << (((^(v5))) & 63))}
                    _ = v31
                case "hello":
                    v3 = 1
                default:
                    v9 = append(v9, fi28)
                }
            }
            var v32 int8 = (v3 & (v3 ^ 79))
            _ = fi28
            _ = v32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
