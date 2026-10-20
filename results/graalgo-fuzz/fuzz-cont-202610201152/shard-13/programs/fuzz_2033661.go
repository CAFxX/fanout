package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint64; F1 uint8 }
func genFn0() uint {
    var v7 uint16 = 0
    var v8 int16 = 32767
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 42
    var v3 int8 = 28
    var v4 int32 = -2147483648
    var v5 int16 = 42
    v6 := []uint32{2, (v2 & v2)}
    v4 -= (-(v4))
    if true {
        var v9 bool = true
        if len(v6) > 0 {
            v6[0] = (v2 ^ 2)
        }
        _ = v9
    } else {
        var v10 uint = 750
        _ = v10
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            v13 := GenS1{F0: ((!(true || true)) || true), F1: genFn0()}
            switch 323 {
            case -817:
                if (true && (false || true)) {
                    _ = len(v6)
                    v14 := GenS2{F0: 255}
                    _ = v14
                }
            }
            if false {
                v15 := GenS3{F0: 646, F1: 0, F2: 18446744073709551615}
                delete(v1, "")
                _ = v15
            }
            _ = fi12
            _ = v13
        }
        var v16 int64 = 0
        if ((!true) && true) {
            switch "foo bar" {
            case "\n\t":
                v6 = append(v6, v2)
            default:
                var v17 int32 = v4
                _ = v17
            }
            if len(v6) > 0 {
                v6[0] = v2
            }
            fmt.Println(v5)
        }
        _ = fi11
        _ = v16
    }
    var v18 int16 = (v5 + -1)
    for fi19 := 0; fi19 < 17; fi19++ {
        v20 := GenS4{F0: 4294967295, F1: 103, F2: 2147483647}
        _ = fi19
        _ = v20
    }
    v6 = append(v6, v2)
    var v21 int64 = -9223372036854775808
    v22 := GenS4{F0: 0, F1: 1, F2: (v4 ^ v4)}
    var v23 int32 = v4
    switch -1 {
    case 128:
        fmt.Println(v4, v18, v21)
    default:
        for fi24 := 0; fi24 < 16; fi24++ {
            v25 := GenS3{F0: 2147483647, F1: 283, F2: 1}
            v23 = (v23 / ((v4) | 1))
            v21 = v21
            _ = fi24
            _ = v25
        }
    }
    v18 -= -147
    var v26 int8 = v3
    var v27 int16 = v18
    v4 += v23
    v28 := GenS5{F0: 38, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
