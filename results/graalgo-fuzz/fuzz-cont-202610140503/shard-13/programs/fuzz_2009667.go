package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint8 = 127
    var v2 int32 = -1
    var v3 int8 = -128
    var v4 int64 = 9223372036854775807
    if (!(true || false)) {
        var v5 int32 = v2
        switch 42 {
        case 240:
            v6 := GenS1{F0: ((v1 ^ v1) <= 255), F1: 0, F2: ("foo bar" + "")}
            _ = v6
        case 621:
            fmt.Println(v4)
        case 32767:
            var v7 bool = (1000 == 255)
            _ = v7
        default:
            v4 += v4
        }
        _ = v5
    }
    var v8 string = "foo bar"
    v3 += v3
    var v9 uint32 = 519
    if (1 != 18446744073709551615) {
        var v10 int32 = (v2 / ((0) | 1))
        if (uint64(v1) > 2) {
            v11 := []int{0, 32767, 9223372036854775807}
            _ = v11
        } else {
            var v12 float64 = 1e10
            fmt.Println(v3)
            v4 = (v4 & v4)
            _ = v12
        }
        _ = v10
    } else {
        var v13 int = 685
        _ = v13
    }
    fmt.Println(v9, v9, v4)
    v14 := GenS2{F0: int16(v2)}
    fmt.Println(v8, v4)
    v15 := GenS1{F0: false, F1: 559, F2: (v8 + v8)}
    for fi16 := 0; fi16 < 3; fi16++ {
        v8 = "\n\t"
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
