package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint; F1 int8 }

func main() {
    v1 := []uint32{273, 2147483647, 4294967295}
    var v2 uint32 = 4294967295
    v3 := []uint32{2, (v2 * v2), 0, (v2 ^ (v2 % (((^(v2))) | 1)))}
    var v4 int16 = 0
    fmt.Println(v4)
    if (false || true) {
        var v5 uint64 = 2
        if false {
            switch 127 {
            case 913:
                var v6 uint8 = 67
                _ = v6
            case 9223372036854775807:
                v7 := GenS1{F0: -2147483648, F1: -1, F2: 163}
                _ = v7
            }
            fmt.Println(v4)
        }
        v8 := GenS1{F0: -1, F1: 52, F2: 128}
        _ = v5
        _ = v8
    } else {
        v4 -= (v4 % ((v4) | 1))
        fmt.Println(v2, v2)
    }
    var v9 int16 = v4
    switch 32768 {
    case 127:
        var v10 int64 = 0
        _ = v10
    case 32768:
        v11 := GenS1{F0: 0, F1: 42, F2: 1}
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            fmt.Println(v4, v2, v2)
            v9 += (v4 ^ v4)
            _ = fi13
        }
        _ = fi12
    }
    var v14 float64 = (((-2.25 - ((((1e10 + float64((float32(float64(((-2.25 - 3.14159) * ((3.14159 - 1.5) + 0.1)))) * float32(1.5)))) + ((0.0 * (float64(3.14159) + (0.0 / -2.25))) - (1e10 + 0.0))) * ((0.1 + 3.14159) - (-2.25 + (-2.25 + 0.1)))) - (0.1 - ((0.1 + 1.5) / (-2.25 / float64(float32((1e10 / 0.1)))))))) - (-2.25 * (((0.1 * 0.0) + (1.5 + (3.14159 + float64(float32(1.5))))) * -2.25))) * float64(((1.5 / ((float32(-2.25) - 1e10) - 1.5)) - 3.14159)))
    switch "" {
    case "hello":
        switch "foo bar" {
        case "":
            var v15 uint8 = 127
            _ = v15
        case "hello":
            v2 = (v2 ^ (v2 ^ (v2 * (v2 % (((^(v2))) | 1)))))
        case "\n\t":
            v16 := GenS2{F0: 2, F1: 125}
            _ = v16
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch -110 {
    case -9223372036854775808:
        if len(v3) > 0 {
            v3[0] = v2
        }
    case -2147483648:
        var v17 bool = true
        _ = v17
    case 2147483647:
        fmt.Println(v9)
    default:
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 uint = 18446744073709551615
            var v20 uint = (v19 ^ v19)
            var v21 int8 = 42
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
    }
    v2 = (v2 + v2)
    v22 := GenS1{F0: -1, F1: 42, F2: 71}
    fmt.Println(v2)
    if ((42 > 32768) || true) {
        for fi23 := 0; fi23 < 4; fi23++ {
            v4 = 0
            _ = fi23
        }
    }
    v4 += (-(v4))
    v2 = (v2 / (((v2 & 2)) | 1))
    v2 = (v2 + 982)
    var v24 int16 = v9
    v14 = (v14 - float64(3.14159))
    for fi25 := 0; fi25 < 13; fi25++ {
        v26 := GenS2{F0: 2, F1: 1}
        _ = fi25
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
