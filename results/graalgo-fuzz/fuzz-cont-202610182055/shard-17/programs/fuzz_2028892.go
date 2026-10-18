package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 bool; F1 uint16 }

func main() {
    var v1 uint8 = 2
    var v2 uint8 = 4
    var v3 string = ""
    var v4 int32 = -2147483648
    v5 := make(map[uint32]bool)
    var v6 uint8 = (v1 ^ 127)
    v6 = (v6 % (((v1 % ((v6) | 1))) | 1))
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(v1, v4)
        var v8 uint32 = 157
        v9 := []int{32767, fi7, (fi7 - fi7), -2147483648}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 0, F1: 9223372036854775807}
    v3 = ""
    switch -2147483648 {
    case 1:
        var v11 bool = ((-2.25 + -2.25) == (-2.25 * float32(0.1)))
        _ = v11
    case 32767:
        fmt.Println(v1)
    case 127:
        var v12 int32 = (v4 ^ (v4 ^ v4))
        _ = v12
    default:
        v13 := GenS1{F0: 66, F1: -919}
        _ = v13
    }
    if true {
        var v14 float32 = 1e10
        var v15 uint16 = 160
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := GenS1{F0: 127, F1: 340}
            v18 := GenS1{F0: -19, F1: -1}
            var v19 float64 = 0.1
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        _ = v14
        _ = v15
    } else {
        v20 := []int{-2147483648, 0, 32767}
        var v21 uint = 2
        var v22 float32 = ((1e10 - -2.25) * (float32(3.14159) * 1.5))
        _ = v20
        _ = v21
        _ = v22
    }
    v2 = v6
    var v23 uint = 2
    var v24 uint16 = 160
    v25 := GenS2{F0: (true && false), F1: (v24 * 1)}
    var v26 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
