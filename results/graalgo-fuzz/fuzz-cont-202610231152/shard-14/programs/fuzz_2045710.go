package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint64 }

func main() {
    v1 := []string{"hello", ("" + "\n\t")}
    v2 := make(map[int]bool)
    v3 := make(map[uint32]bool)
    var v4 int64 = 1
    v4 = 42
    if false {
        for fi5 := 0; fi5 < 19; fi5++ {
            v4 = v4
            v4 -= (-(v4))
            fmt.Println(fi5)
            _ = fi5
        }
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 float32 = 1.5
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := GenS1{F0: 1000, F1: 2}
                if (!(((^(v4)) == (v4 ^ (v4 / ((v4) | 1)))) || false)) {
                    v10 := GenS1{F0: -32768, F1: 369}
                    fmt.Println(v4)
                    v11 := GenS2{F0: 32767, F1: 18446744073709551615}
                    _ = v10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
            var v12 bool = true
            _ = fi6
            _ = v12
            _ = v7
        }
        fmt.Println(v4)
    }
    switch "\n\t" {
    case "":
        switch 42 {
        case 837:
            if ((-32768 != 32767) || false) {
                switch -1 {
                case 33:
                    v13 := GenS1{F0: 42, F1: 32767}
                    _ = v13
                }
                v4 = (v4 % ((-9223372036854775808) | 1))
            } else {
                v4 += v4
                v4 = -754
            }
        }
    }
    if (1 <= -2147483648) {
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := GenS1{F0: -457, F1: 890}
            var v16 uint16 = 321
            _ = fi14
            _ = v15
            _ = v16
        }
        v4 = v4
        var v17 float64 = 0.0
        _ = v17
    } else {
        var v18 uint16 = 65535
        var v19 int32 = 42
        _ = v18
        _ = v19
    }
    var v20 bool = true
    var v21 int = 0
    v22 := GenS1{F0: 95, F1: 90}
    v1 = append(v1, ("" + ("" + "\n\t")))
    v21 = v21
    fmt.Println(v21, v20)
    var v23 int64 = 1
    if (545 <= 511) {
        fmt.Println(v21)
    } else {
        var v24 uint32 = 1
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
