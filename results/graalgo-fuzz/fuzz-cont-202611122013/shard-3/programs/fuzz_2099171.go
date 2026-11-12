package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint = 2
    var v2 int32 = -748
    var v3 float64 = 3.14159
    v4 := []int{0, 42, -9223372036854775808}
    v5 := GenS1{F0: 134, F1: 2147483648, F2: 2147483647}
    if len(v4) > 0 {
        v4[0] = -784
    }
    switch 650 {
    case 32767:
        if true {
            var v6 int64 = 9223372036854775807
            var v7 float32 = ((1e10 - 0.0) + 1e10)
            _ = v6
            _ = v7
        }
    case 128:
        v8 := GenS2{F0: 32767, F1: 21}
        _ = v8
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if len(v4) > 0 {
        v4[0] = 118
    }
    v2 = 853
    fmt.Println(v2)
    v9 := GenS3{F0: 0, F1: 42, F2: 974}
    var v10 int32 = (v2 + v2)
    switch "" {
    case "hello":
        v1 -= (v1 % ((880) | 1))
    case "\n\t":
        switch 332 {
        case 4294967295:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case -31:
            var v11 int = -2147483648
            _ = v11
        default:
            _ = len(v4)
        }
    case "foo bar":
        v12 := GenS4{F0: 100}
        _ = v12
    default:
        switch -966 {
        case 1:
            for fi13 := 0; fi13 < 16; fi13++ {
                if (v10 >= (v2 + (v2 + v10))) {
                    fmt.Println(fi13)
                    var v14 uint16 = 2
                    v15 := GenS3{F0: 2147483648, F1: 42, F2: 1}
                    _ = v14
                    _ = v15
                }
                _ = fi13
            }
        default:
            v2 -= (v2 & -59)
        }
    }
    var v16 uint16 = 0
    var v17 int32 = (v10 ^ 42)
    if (!true) {
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 bool = ((v17 << ((v1) & 63)) >= v10)
            _ = fi18
            _ = v19
        }
    } else {
        v17 += (v17 | (v17 ^ -2147483648))
        var v20 int32 = v2
        _ = v20
    }
    fmt.Println(v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
