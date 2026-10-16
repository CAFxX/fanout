package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int16; F2 int8 }
type GenS4 struct { F0 int16; F1 string; F2 uint16 }
func genFn0() int32 {
    switch "" {
    case "\n\t":
        if false {
            var v6 int64 = 42
            _ = v6
        }
    case "a":
        var v7 int16 = -985
        _ = v7
    default:
        v8 := []int{}
        _ = v8
    }
    var v9 uint8 = 127
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS1{F0: ("hello" + ("" + (((("\n\t" + "foo bar") + "a") + "a") + "a")))}
        for fi12 := 0; fi12 < 18; fi12++ {
            fmt.Println(fi12)
            _ = fi12
        }
        if false {
            var v13 int32 = -2147483648
            _ = v13
        } else {
            v14 := []int{}
            if (2 > 1) {
                if false {
                    v9 = v9
                    v15 := GenS2{F0: v9, F1: v9}
                    var v16 int = (-(fi10))
                    _ = v15
                    _ = v16
                }
                var v17 uint16 = 1
                _ = v17
            }
            _ = v14
        }
        _ = fi10
        _ = v11
    }
    v9 = (v9 >> ((0) & 63))
    _ = v9
    return 2147483647
}

func main() {
    var v1 uint16 = 42
    var v2 int32 = -1
    var v3 float64 = -2.25
    var v4 uint32 = 468
    v5 := make(map[int]int)
    v18 := GenS3{F0: (v1 & 32767), F1: 314, F2: -47}
    v19 := v5[2147483647]
    v19 = -9223372036854775808
    for fi20 := 0; fi20 < 15; fi20++ {
        v21 := []int{-2147483648, v19, (fi20 << ((515) & 63)), -1}
        fmt.Println(v1, v4)
        _ = fi20
        _ = v21
    }
    var v22 float32 = (float32(v3) + (((float32(v3) * (0.0 / (0.1 * -2.25))) / float32(float64((float32(v3) - -2.25)))) - float32(float64(0.1))))
    var v23 uint = 342
    var v24 float64 = -2.25
    var v25 bool = true
    v26 := []int{(v19 << ((2) & 63))}
    if (2 == v4) {
        fmt.Println(v4, v25)
    }
    v4 -= (v4 / ((v4) | 1))
    var v27 int = (v19 / (((v19 << ((v23) & 63))) | 1))
    v4 = (v4 & (v4 >> (((v23 / ((v23) | 1))) & 63)))
    v28 := GenS4{F0: 42, F1: "", F2: (^(v1))}
    for fi29 := 0; fi29 < 4; fi29++ {
        for fi30 := 0; fi30 < 9; fi30++ {
            var v31 uint8 = 0
            _ = fi30
            _ = v31
        }
        for fi32 := 0; fi32 < 7; fi32++ {
            if len(v26) > 0 {
                v26[0] = (v19 >> ((v23) & 63))
            }
            var v33 float64 = float64(v22)
            switch 127 {
            case 95:
                v34 := GenS3{F0: (v1 << ((v23) & 63)), F1: 847, F2: -128}
                _ = v34
            }
            _ = fi32
            _ = v33
        }
        _ = fi29
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        var v36 uint = v23
        _ = fi35
        _ = v36
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        v19 = -2147483648
        v19 -= (v19 ^ v19)
        _ = fi37
    }
    if false {
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
