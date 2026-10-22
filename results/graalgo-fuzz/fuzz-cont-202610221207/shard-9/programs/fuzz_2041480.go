package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int8 }
type GenS2 struct { F0 int16; F1 uint64 }

func main() {
    var v1 float64 = float64(-2.25)
    v2 := []string{("foo bar" + ("\n\t" + "a"))}
    v3 := make(map[int]bool)
    var v4 float32 = (0.1 * (1.5 * 1.5))
    var v5 int8 = 100
    v6 := make(map[int]string)
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: 73, F1: 296, F2: v5}
        _ = len(v3)
        _ = fi7
        _ = v8
    }
    v5 += (v5 * 18)
    for fi9 := 0; fi9 < 11; fi9++ {
        fmt.Println(v4, v4)
        _ = fi9
    }
    v1 = ((float64(0.0) + -2.25) - v1)
    switch -2147483648 {
    case 32767:
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int8 = (v5 ^ v5)
            if (!false) {
                var v12 uint8 = 42
                var v13 string = ""
                for fi14 := 0; fi14 < 18; fi14++ {
                    fmt.Println(fi14, v13)
                    v15 := v6[(fi14 + (fi10 >> ((65535) & 63)))]
                    _ = fi14
                    _ = v15
                }
                _ = v12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
    case -2147483648:
        v3[42] = true
    case 128:
        fmt.Println(v5, v5, v1)
    }
    v16 := GenS1{F0: (v5 << ((42) & 63)), F1: 2, F2: (v5 & (v5 << ((1) & 63)))}
    v17 := GenS2{F0: 1, F1: 18446744073709551615}
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
