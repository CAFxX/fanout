package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int32 }
type GenS2 struct { F0 bool; F1 uint64; F2 int8 }
func genFn0(p0 int16) uint32 {
    var v4 bool = false
    var v5 uint = 289
    _ = p0
    _ = v4
    _ = v5
    return 571
}

func main() {
    var v1 int64 = -217
    v2 := make(map[int]int)
    var v3 uint = 765
    var v6 int32 = -877
    if (!(("a" + "a") <= ("" + "foo bar"))) {
        var v7 float32 = 0.0
        v3 -= (v3 % (((^(v3))) | 1))
        _ = v7
    }
    v6 -= (v6 % (((v6 - 2147483647)) | 1))
    v6 = v6
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 int16 = 32767
        var v10 int32 = v6
        v9 = v9
        _ = fi8
        _ = v10
        _ = v9
    }
    v6 += -240
    switch "a" {
    case "foo bar":
        var v11 uint8 = 103
        _ = v11
    case "\n\t":
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 int64 = -9223372036854775808
            v1 -= -9223372036854775808
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := []int{}
        fmt.Println(v3, v6, v6)
        for fi16 := 0; fi16 < 11; fi16++ {
            v17 := GenS1{F0: -1, F1: (false && false), F2: (v6 & v6)}
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    v18 := v2[128]
    var v19 bool = (!((17 != 1) || (!true)))
    v1 += (v1 % ((-1) | 1))
    v20 := GenS2{F0: (2147483647 == 65535), F1: 1, F2: 42}
    var v21 uint = (v3 + v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v6
}
