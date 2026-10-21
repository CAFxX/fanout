package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
func genFn0(p0 int, p1 int16) uint {
    v6 := []int{p0, p0, p0, 42}
    v7 := GenS1{F0: -1, F1: -97}
    v8 := GenS1{F0: 42, F1: 57}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint16 = 62
    var v2 uint = 208
    var v3 int8 = 100
    var v4 int32 = -327
    var v5 uint = v2
    if false {
        for fi9 := 0; fi9 < 1; fi9++ {
            switch "a" {
            case "":
                var v10 float32 = 0.1
                _ = v10
            }
            var v11 int32 = -2147483648
            _ = fi9
            _ = v11
        }
        v12 := GenS1{F0: 0, F1: (v3 * v3)}
        fmt.Println(v3)
        _ = v12
    }
    var v13 float32 = (0.1 * (0.1 + 0.0))
    v14 := []int{int(v5), 4294967295}
    v3 += (-(v3))
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := GenS1{F0: 1, F1: (v3 / (((-(v3))) | 1))}
        _ = fi15
        _ = v16
    }
    if false {
        switch "\n\t" {
        case "foo bar":
            var v17 uint32 = 327
            _ = v17
        case "":
            var v18 int8 = (v3 % (((^(v3))) | 1))
            _ = v18
        case "hello":
            var v19 int8 = (v3 << ((255) & 63))
            _ = v19
        }
    } else {
        var v20 int32 = (v4 + (v4 & (v4 % ((42) | 1))))
        v20 = -2147483648
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = v20
    }
    var v21 uint8 = 2
    fmt.Println(v5, v21)
    var v22 float64 = float64((v13 * (v13 / 1.5)))
    v4 -= -2147483648
    var v23 int = 42
    var v24 float32 = (v13 - v13)
    var v25 uint32 = 159
    var v26 uint8 = (^(v21))
    var v27 int = int(v25)
    if false {
        v28 := GenS1{F0: 42, F1: v3}
        _ = v28
    } else {
        if ((v27 << ((uint(v25)) & 63)) == (-(v23))) {
            for fi29 := 0; fi29 < 6; fi29++ {
                for fi30 := 0; fi30 < 15; fi30++ {
                    fmt.Println(v26, v25, v2)
                    _ = fi30
                }
                _ = fi29
            }
        }
    }
    v21 -= (v21 & 255)
    _ = len(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
