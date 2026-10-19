package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int64 }
func genFn0() uint64 {
    if (2 >= 4294967295) {
        var v4 float64 = float64((((1e10 / ((float32(-2.25) - (3.14159 * float32(1.5))) * (0.1 + -2.25))) * float32(1.5)) / -2.25))
        _ = v4
    }
    var v5 uint16 = 255
    v5 = (v5 / ((v5) | 1))
    v6 := GenS1{F0: 1, F1: (v5 & v5), F2: 523}
    v5 = (v5 + 2)
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int32 = 1
    v2 := []string{"foo bar"}
    var v3 uint8 = 128
    v1 -= v1
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int8 = -128
        fmt.Println(fi7)
        fmt.Println(v3, v3, v3)
        _ = fi7
        _ = v8
    }
    var v9 int64 = -197
    for fi10 := 0; fi10 < 6; fi10++ {
        v2 = append(v2, "\n\t")
        _ = fi10
    }
    if true {
        var v11 uint64 = 1
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS1{F0: (-(v9)), F1: 32767, F2: (v9 | 0)}
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    if (-468 <= 42) {
        if ("" != ("" + "\n\t")) {
            for fi14 := 0; fi14 < 19; fi14++ {
                var v15 int = (fi14 - fi14)
                var v16 int8 = 100
                for fi17 := 0; fi17 < 1; fi17++ {
                    fmt.Println(fi17, fi17)
                    var v18 string = "\n\t"
                    v19 := GenS1{F0: 1, F1: 32767, F2: (v9 / ((v9) | 1))}
                    _ = len(v2)
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            v20 := GenS1{F0: 125, F1: 42, F2: (v9 / (((v9 << ((0) & 63))) | 1))}
            _ = v20
        }
        v3 = 127
    }
    v21 := GenS1{F0: v9, F1: 65535, F2: v9}
    fmt.Println(v9, v1, v9)
    if true {
        fmt.Println(v3)
        var v22 int64 = (v9 | 1)
        _ = v22
    }
    _ = len(v2)
    v9 += (v9 / (((v9 * (v9 << ((42) & 63)))) | 1))
    v3 = (v3 | (^(v3)))
    var v23 uint = 42
    v1 = (v1 | v1)
    v24 := GenS1{F0: (-(v9)), F1: 0, F2: -804}
    v23 = (v23 + (^(v23)))
    v2 = append(v2, "")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v9
}
