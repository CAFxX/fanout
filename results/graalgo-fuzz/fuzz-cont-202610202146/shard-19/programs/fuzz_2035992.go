package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 bool; F1 int8 }

func main() {
    var v1 float32 = 1.5
    v2 := []uint32{}
    var v3 uint8 = 5
    var v4 int = -378
    var v5 int64 = -9223372036854775808
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: v5, F1: 2}
        if ((42 < 539) || (17 > 2)) {
            var v8 uint32 = 2
            var v9 float32 = float32(-2.25)
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: 1, F1: 18446744073709551615}
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    v1 = v1
    v4 = 9223372036854775807
    var v11 bool = (216 <= v3)
    v12 := GenS2{F0: ((-2.25 + 0.1) != (3.14159 + 1e10)), F1: 1}
    if (!false) {
        v13 := GenS2{F0: (!true), F1: 0}
        v1 += 1e10
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 bool = (v11 || true)
            var v16 uint32 = 42
            var v17 string = ""
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = v13
    }
    v3 += (v3 % ((v3) | 1))
    for fi18 := 0; fi18 < 7; fi18++ {
        v11 = ((1e10 + v1) <= v1)
        var v19 uint16 = 46
        var v20 int = 32767
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
