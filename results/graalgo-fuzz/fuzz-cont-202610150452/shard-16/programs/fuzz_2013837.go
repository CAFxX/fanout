package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int8 }
type GenS2 struct { F0 int16; F1 int; F2 int64 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS4 struct { F0 int16 }
func genFn0() int32 {
    var v7 string = ("a" + (("\n\t" + ("\n\t" + "")) + "foo bar"))
    var v8 uint16 = 544
    _ = v7
    _ = v8
    return 336
}
func genFn1(p0 int64, p1 int64, p2 uint8) string {
    var v9 int = -1
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return "foo bar"
}

func main() {
    var v1 int32 = -471
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    var v4 bool = ((float32(3.14159) + float32(0.0)) >= float32(1.5))
    var v5 uint64 = 4294967295
    var v6 uint16 = 1
    v10 := GenS1{F0: 4294967295, F1: 42, F2: 0}
    var v11 uint = 886
    switch 164 {
    case -1:
        if v4 {
            v11 = 42
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 uint64 = v5
                _ = fi12
                _ = v13
            }
        }
    case -357:
        if v4 {
            v14 := GenS2{F0: 1000, F1: 32768, F2: -1}
            var v15 uint8 = 42
            _ = v14
            _ = v15
        }
    }
    v16 := GenS1{F0: 2147483647, F1: v5, F2: -128}
    v17 := GenS2{F0: 0, F1: -9223372036854775808, F2: 42}
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS1{F0: int(v6), F1: v5, F2: 42}
        fmt.Println(fi18, v11)
        _ = fi18
        _ = v19
    }
    var v20 float64 = 1.5
    if ((v11 % (((v11 % (((v11 << ((v11) & 63))) | 1))) | 1)) != v11) {
        for fi21 := 0; fi21 < 2; fi21++ {
            v22 := GenS1{F0: (-(fi21)), F1: v5, F2: -1}
            _ = fi21
            _ = v22
        }
        v23 := GenS3{F0: 663, F1: (0 >= 71), F2: 2147483647}
        _ = v23
    }
    var v24 int8 = 100
    var v25 bool = (!v4)
    var v26 int16 = 902
    v4 = (!v25)
    v27 := GenS4{F0: 1000}
    if (0.0 > (1e10 / -2.25)) {
        switch 42 {
        case -852:
            v24 += v24
        default:
            switch "hello" {
            case "a":
                var v28 int8 = (v24 / ((v24) | 1))
                _ = v28
            default:
                var v29 float64 = v20
                _ = v29
            }
        }
        fmt.Println(v24, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
