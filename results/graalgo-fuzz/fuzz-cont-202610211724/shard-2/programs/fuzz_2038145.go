package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }

func main() {
    var v1 bool = ((0 != 1) && (4294967295 >= 42))
    var v2 int = 128
    var v3 int64 = 9223372036854775807
    var v4 uint64 = 4294967295
    v3 = (v3 * int64(v2))
    v4 += (v4 / ((901) | 1))
    var v5 bool = v1
    v6 := GenS1{F0: v1, F1: -532}
    v7 := []int{(v2 << ((42) & 63)), 127, -598, v2}
    var v8 int = (v2 % ((v2) | 1))
    var v9 int16 = 42
    var v10 bool = v1
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 int16 = 1000
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v5 = (110 >= 128)
    var v14 float32 = ((float32(0.0) - 3.14159) + 0.1)
    if v1 {
        var v15 uint8 = uint8(v4)
        var v16 uint = 65535
        _ = v15
        _ = v16
    }
    var v17 string = "hello"
    v3 = v3
    if (!false) {
        if true {
            v18 := GenS1{F0: (!v1), F1: -1}
            if len(v7) > 0 {
                v7[0] = v8
            }
            _ = v18
        } else {
            fmt.Println(v1, v5)
            var v19 float64 = -2.25
            _ = v19
        }
    } else {
        if (1 != 0) {
            switch 0 {
            case 195:
                switch v8 {
                case -9223372036854775808:
                    var v20 uint = 18446744073709551615
                    _ = v20
                }
            }
            v9 += v9
        } else {
            var v21 int32 = 0
            fmt.Println(v5, v14)
            _ = v21
        }
    }
    var v22 float32 = -2.25
    switch 127 {
    case 42:
        v23 := GenS1{F0: v10, F1: -98}
        _ = v23
    default:
        for fi24 := 0; fi24 < 6; fi24++ {
            var v25 int = v2
            if ((42 == (v4 * (^(v4)))) || (!false)) {
                v7 = append(v7, (fi24 | v25))
            }
            var v26 uint32 = 251
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    var v27 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v17
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
