package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 string = "\n\t"
    var v2 uint64 = 4294967295
    var v3 float32 = (float32((3.14159 * float64(0.1))) / (0.0 + (1e10 - -2.25)))
    var v4 int8 = -19
    var v5 uint8 = 0
    var v6 uint32 = 802
    v7 := GenS1{F0: 42, F1: 900}
    fmt.Println(v2, v5, v3)
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 int64 = -1
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := GenS1{F0: 268, F1: 42}
            v12 := []int{(fi8 >> ((255) & 63)), (fi10 / ((fi10) | 1)), -9223372036854775808}
            var v13 float32 = v3
            switch 4294967295 {
            case 9223372036854775807:
                v3 += v3
            }
            v6 += (v6 << ((733) & 63))
            v2 = (v2 + v2)
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        switch (-(fi8)) {
        case -219:
            for fi14 := 0; fi14 < 15; fi14++ {
                v15 := GenS2{F0: (v1 + (v1 + (v1 + ((("" + v1) + (v1 + v1)) + ((v1 + (("hello" + v1) + "hello")) + v1)))))}
                v16 := GenS3{F0: (v6 << ((30) & 63)), F1: (^(v6))}
                _ = fi14
                _ = v15
                _ = v16
            }
        case 127:
            fmt.Println(v5, v1)
        case -2147483648:
            switch "" {
            case "\n\t":
                fmt.Println(v9)
            default:
                v1 = v1
            }
        }
        _ = fi8
        _ = v9
    }
    if false {
        v17 := GenS1{F0: 65535, F1: 740}
        _ = v17
    }
    var v18 string = v1
    var v19 float32 = ((v3 - 1.5) + v3)
    var v20 int8 = v4
    v4 -= (v20 + v20)
    var v21 int8 = 19
    if (!true) {
        if ((-2.25 != ((1.5 - 1.5) - 1e10)) || false) {
            var v22 float64 = float64(v19)
            v18 = v18
            fmt.Println(v4, v18)
            _ = v22
        } else {
            v23 := []int{32767, 191, 9223372036854775807}
            fmt.Println(v2, v18)
            _ = v23
        }
        v4 = 42
    }
    v24 := []int{}
    _ = len(v24)
    if false {
        var v25 string = ("" + "\n\t")
        var v26 int8 = -128
        var v27 float64 = ((((1e10 + ((3.14159 / 1e10) / 1.0)) + 1.5) / ((1e10 * 0.1) - 0.1)) / 1.5)
        _ = v25
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
