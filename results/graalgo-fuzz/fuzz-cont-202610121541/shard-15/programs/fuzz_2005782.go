package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int; F1 uint8; F2 uint8 }

func main() {
    var v1 bool = true
    var v2 int = -1
    var v3 int32 = 2147483647
    v4 := make(map[uint32]string)
    v5 := make(map[int]string)
    v6 := make(map[int]bool)
    var v7 int = (v2 - 678)
    if true {
        v8 := GenS1{F0: v1, F1: 2147483648}
        v9 := GenS1{F0: (899 > v3), F1: 42}
        v7 = (-(v7))
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 float64 = float64(0.1)
            _ = fi10
            _ = v11
        }
    }
    fmt.Println(v3, v7)
    if v1 {
        v12 := GenS1{F0: v1, F1: 42}
        _ = v12
    } else {
        if (0 > -1) {
            var v13 int8 = 5
            v14 := []int{v7}
            fmt.Println(v3, v1, v13)
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            v16 := GenS1{F0: v1, F1: 1}
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 uint64 = 0
            for fi19 := 0; fi19 < 1; fi19++ {
                v18 = v18
                _ = fi19
            }
            _ = fi17
            _ = v18
        }
    }
    var v20 uint8 = 127
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 int64 = 42
        fmt.Println(v1, v3, v3)
        _ = fi21
        _ = v22
    }
    fmt.Println(v2)
    v23 := GenS1{F0: (!false), F1: 329}
    if (!(!v1)) {
        v24 := GenS1{F0: v1, F1: 683}
        v25 := GenS2{F0: (v7 + v7), F1: (v20 + 128), F2: 0}
        _ = v24
        _ = v25
    }
    var v26 uint8 = 73
    var v27 int8 = 0
    switch -9223372036854775808 {
    case -2147483648:
        var v28 int = (v2 * (-(v7)))
        _ = v28
    case 80:
        var v29 int8 = v27
        _ = v29
    case -353:
        v7 += -9223372036854775808
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        v31 := GenS2{F0: 802, F1: v26, F2: v20}
        var v32 int32 = (v3 | -688)
        var v33 int16 = 42
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    v2 += 42
    var v34 uint16 = 94
    v35 := GenS1{F0: (v3 <= 2147483647), F1: 179}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
