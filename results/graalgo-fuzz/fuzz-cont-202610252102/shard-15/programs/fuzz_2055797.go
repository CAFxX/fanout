package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint }
type GenS3 struct { F0 int16; F1 uint64; F2 uint32 }

func main() {
    var v1 int32 = 0
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    var v5 float64 = (-2.25 + 1e10)
    var v6 int = 42
    if true {
        if false {
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 float64 = (v5 - v5)
                v9 := []int{(v6 + fi7)}
                _ = len(v4)
                _ = fi7
                _ = v8
                _ = v9
            }
            v2[v6] = false
        } else {
            if true {
                var v10 float32 = 1.5
                switch "" {
                case "foo bar":
                    for fi11 := 0; fi11 < 20; fi11++ {
                        var v12 int64 = 1
                        v13 := GenS1{F0: 0, F1: 0, F2: 632}
                        fmt.Println(v1, v6, v5)
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                default:
                    v5 = 1e10
                }
                _ = v10
            }
            var v14 uint = 882
            _ = v14
        }
        switch v6 {
        case -9223372036854775808:
            for fi15 := 0; fi15 < 1; fi15++ {
                v16 := GenS2{F0: 42, F1: 4294967295, F2: 1}
                var v17 int16 = -188
                _ = fi15
                _ = v16
                _ = v17
            }
        case 405:
            fmt.Println(v5, v1, v1)
        case -341:
            fmt.Println(v6, v5)
        }
        for fi18 := 0; fi18 < 12; fi18++ {
            switch v6 {
            case -9223372036854775808:
                v19 := GenS1{F0: 711, F1: 0, F2: 1}
                _ = v19
            case 127:
                var v20 uint32 = 348
                _ = v20
            case 332:
                var v21 int32 = -311
                _ = v21
            }
            var v22 float64 = (((v5 / 3.14159) - float64((-2.25 * (float32(v5) * 1e10)))) * v5)
            _ = fi18
            _ = v22
        }
    }
    v23 := GenS2{F0: 1, F1: 2147483647, F2: 672}
    if true {
        v6 = 127
        v24 := GenS3{F0: -1, F1: 18446744073709551615, F2: 2}
        _ = v24
    }
    fmt.Println(v1, v1, v5)
    var v25 uint8 = 255
    v26 := GenS2{F0: 1, F1: 0, F2: 18446744073709551615}
    fmt.Println(v6, v25, v6)
    if (!(true || ((!(((-2.25 / (float64(0.1) / v5)) / v5) < v5)) && false))) {
        var v27 uint64 = 703
        fmt.Println(v5, v6, v1)
        var v28 int8 = 89
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
