package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0() int64 {
    v6 := GenS1{F0: -1, F1: 128, F2: 2}
    for fi7 := 0; fi7 < 5; fi7++ {
        if false {
            v8 := GenS2{F0: 9223372036854775807, F1: 2147483647}
            var v9 int32 = -239
            fmt.Println(fi7, fi7)
            _ = v8
            _ = v9
        }
        var v10 float32 = 0.0
        v10 = float32(1e10)
        _ = fi7
        _ = v10
    }
    switch 1 {
    case 128:
        v11 := []int{2147483647}
        _ = v11
    default:
        var v12 uint8 = 0
        _ = v12
    }
    var v13 bool = false
    _ = v13
    _ = v6
    return 253
}

func main() {
    var v1 int32 = -623
    var v2 int = 2147483647
    var v3 int64 = -9223372036854775808
    v4 := []string{("\n\t" + "a"), ("\n\t" + ""), "hello", "hello"}
    v5 := make(map[int]int)
    v14 := GenS2{F0: (v3 | v3), F1: v1}
    v2 = -420
    v15 := GenS3{F0: (v3 ^ (v3 - 9223372036854775807)), F1: 32767}
    if false {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int8 = -1
            _ = fi16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        fmt.Println(v2)
        _ = fi18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := v5[(v2 + v2)]
        _ = fi19
        _ = v20
    }
    v2 += 330
    var v21 uint = 761
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
