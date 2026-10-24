package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint }
type GenS2 struct { F0 int; F1 int64 }

func main() {
    v1 := []uint32{0, 2147483647, 847, 0}
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    var v4 uint32 = 2147483647
    var v5 uint = 129
    var v6 int = 153
    v4 -= v4
    if true {
        var v7 uint32 = (v4 + v4)
        fmt.Println(v4, v4)
        _ = v7
    }
    fmt.Println(v5)
    fmt.Println(v5, v6)
    v8 := GenS1{F0: 0, F1: (v4 / ((v4) | 1)), F2: (v5 ^ v5)}
    var v9 uint8 = 1
    _ = len(v2)
    var v10 uint8 = (v9 % ((v9) | 1))
    for fi11 := 0; fi11 < 20; fi11++ {
        fmt.Println(fi11)
        switch -243 {
        case 32767:
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 uint64 = 1
                v14 := GenS2{F0: fi11, F1: -1}
                _ = fi12
                _ = v13
                _ = v14
            }
        default:
            delete(v3, 32767)
        }
        if len(v1) > 0 {
            v1[0] = v4
        }
        _ = fi11
    }
    _ = len(v3)
    fmt.Println(v9, v6, v10)
    var v15 bool = false
    fmt.Println(v9, v10, v10)
    switch -590 {
    case -2147483648:
        if (524 > v6) {
            var v16 bool = (!false)
            v15 = (v15 || false)
            _ = v16
        } else {
            switch ("\n\t" + "\n\t") {
            case "":
                if (3.14159 < (0.0 + ((1e10 + 0.0) - ((1e10 - 1e10) - 3.14159)))) {
                    v17 := v3[(^(v6))]
                    switch 42 {
                    case -808:
                        var v18 float64 = (float64(0.1) / 0.1)
                        _ = v18
                    case 4294967295:
                        v5 = (v5 << (((v5 + v5)) & 63))
                    }
                    _ = v17
                } else {
                    v6 = v6
                }
            case "hello":
                var v19 int8 = -1
                _ = v19
            case "\n\t":
                v10 = (v9 / (((v10 >> ((65535) & 63))) | 1))
            default:
                switch v6 {
                case 609:
                    var v20 uint16 = 462
                    _ = v20
                }
            }
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 uint32 = (^(v4))
                _ = fi21
                _ = v22
            }
            v23 := GenS1{F0: 42, F1: v4, F2: v5}
            _ = v23
        }
    default:
        for fi24 := 0; fi24 < 20; fi24++ {
            v5 += v5
            v4 -= v4
            if v15 {
                fmt.Println(v6, v6, v5)
            } else {
                v9 -= 0
                var v25 bool = (100 < -128)
                _ = v25
            }
            _ = fi24
        }
    }
    v26 := GenS1{F0: 134, F1: (v4 * v4), F2: v5}
    v15 = v15
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
