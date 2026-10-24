package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint }
func genFn0(p0 int64, p1 uint64) int64 {
    if true {
        var v4 int32 = -841
        var v5 uint32 = 0
        _ = v4
        _ = v5
    }
    v6 := GenS1{F0: 0, F1: 255}
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 42
    var v3 uint32 = (^(v2))
    if (-9223372036854775808 == genFn0(-1, 4294967295)) {
        v7 := GenS1{F0: -1, F1: 40}
        _ = v7
    } else {
        var v8 float32 = 0.0
        var v9 float64 = 0.0
        _ = v8
        _ = v9
    }
    var v10 int16 = 547
    switch 128 {
    case 127:
        var v11 int64 = genFn0(42, 133)
        _ = v11
    case 4294967295:
        switch "a" {
        case "\n\t":
            for fi12 := 0; fi12 < 13; fi12++ {
                v10 += (v10 + v10)
                v13 := GenS2{F0: 4294967295, F1: 65535, F2: 367}
                v14 := GenS2{F0: uint64(v2), F1: uint16(v3), F2: 2}
                _ = fi12
                _ = v13
                _ = v14
            }
        case "":
            var v15 int16 = 32767
            _ = v15
        case "a":
            v16 := GenS1{F0: (v10 / ((v10) | 1)), F1: 1}
            _ = v16
        }
    }
    fmt.Println(v3, v3)
    v2 -= v3
    v17 := GenS1{F0: v10, F1: 255}
    var v18 uint8 = 127
    var v19 int16 = -32768
    v20 := v1["foo bar"]
    v21 := GenS1{F0: v10, F1: (v18 - (^(v18)))}
    v18 = v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
