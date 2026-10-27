package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint32 }

func main() {
    var v1 bool = (42 <= 128)
    var v2 uint32 = 4294967295
    var v3 int16 = 1000
    var v4 float64 = -2.25
    var v5 uint32 = (v2 - v2)
    var v6 string = "\n\t"
    var v7 float64 = v4
    for fi8 := 0; fi8 < 19; fi8++ {
        if (v3 >= (v3 & (v3 ^ 463))) {
            for fi9 := 0; fi9 < 6; fi9++ {
                v4 = ((1.5 - v4) * v7)
                v4 -= (((v7 + 0.0) * v7) * v7)
                _ = fi9
            }
            v1 = (-125 != 100)
        }
        fmt.Println(fi8, v6, v3)
        _ = fi8
    }
    var v10 uint32 = v2
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS1{F0: false, F1: 0, F2: v5}
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: (-912 != 127), F1: 42, F2: (v10 + v10)}
    v1 = ((!false) && false)
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 int64 = 147
        for fi16 := 0; fi16 < 14; fi16++ {
            fmt.Println(v5)
            v2 = v10
            if ((!true) || true) {
                v10 = (^(v5))
                v3 += (v3 % ((v3) | 1))
                v17 := GenS1{F0: (true && false), F1: 0, F2: v10}
                _ = v17
            } else {
                if (979 == 2) {
                    v5 -= v2
                    var v18 uint = 469
                    _ = v18
                }
                switch "foo bar" {
                case "a":
                    v5 -= v10
                case "hello":
                    var v19 uint8 = 128
                    _ = v19
                case "\n\t":
                    fmt.Println(v4)
                }
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
