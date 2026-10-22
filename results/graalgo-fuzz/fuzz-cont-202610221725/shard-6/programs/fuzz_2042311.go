package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint32 }

func main() {
    var v1 int = 4294967295
    var v2 float32 = (0.1 - (3.14159 / 1.0))
    var v3 float64 = float64(((v2 * v2) / (1.5 * (v2 * v2))))
    var v4 bool = (!(!false))
    for fi5 := 0; fi5 < 8; fi5++ {
        v1 = 42
        _ = fi5
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint = 4294967295
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 uint = 493
            v10 := GenS1{F0: 2147483647, F1: -2147483648, F2: fi6}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v11 int64 = 1
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := []int{(fi12 + (fi12 + (fi12 * v1))), v1, fi12, v1}
        var v14 int16 = 32767
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    v3 = -2.25
    v15 := GenS1{F0: 0, F1: int32(v11), F2: (v1 & (v1 << ((18446744073709551615) & 63)))}
    v16 := GenS1{F0: 665, F1: 171, F2: v1}
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            if ((v4 || v4) && (true || v4)) {
                var v19 int64 = v11
                if v4 {
                    var v20 int8 = 127
                    _ = v20
                }
                _ = v19
            }
            v3 += v3
            v21 := GenS2{F0: 65535, F1: 42, F2: 0}
            _ = fi18
            _ = v21
        }
        _ = fi17
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 uint8 = 1
        var v24 uint8 = (v23 / ((v23) | 1))
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: -957, F1: 42, F2: (v1 - v1)}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
