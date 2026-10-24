package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int }

func main() {
    var v1 int16 = 548
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    v4 := []int{-9223372036854775808, 1, 635}
    var v5 uint = 4294967295
    var v6 float32 = 0.1
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int32 = -2147483648
        var v9 int32 = v8
        _ = fi7
        _ = v8
        _ = v9
    }
    v1 += v1
    for fi10 := 0; fi10 < 8; fi10++ {
        v6 -= v6
        var v11 float64 = (-2.25 / 1.5)
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 float32 = 3.14159
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    var v14 int8 = int8(v5)
    var v15 int32 = 42
    var v16 string = "foo bar"
    v4 = append(v4, -459)
    if ((v1 << ((v5) & 63)) != (v1 * v1)) {
        v17 := GenS1{F0: (v14 & v14), F1: (-(v14)), F2: 32768}
        _ = v17
    } else {
        switch "" {
        case "":
            if ((!false) || ((true || (true && true)) || true)) {
                v18 := GenS1{F0: (v14 >> ((v5) & 63)), F1: (^(v14)), F2: -1}
                var v19 bool = true
                _ = v18
                _ = v19
            } else {
                v16 = v16
                if false {
                    v14 = int8(v1)
                    var v20 uint32 = 65535
                    fmt.Println(v1)
                    _ = v20
                }
            }
        default:
            var v21 uint32 = 65535
            _ = v21
        }
        var v22 float64 = (1e10 + (1.5 / 1e10))
        _ = v22
    }
    v1 += v1
    v5 -= 182
    v15 = (v15 ^ 2147483647)
    if len(v4) > 0 {
        v4[0] = 42
    }
    var v23 int16 = v1
    var v24 int64 = -1
    var v25 int8 = 127
    var v26 uint64 = 209
    var v27 int16 = (v1 + 0)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
