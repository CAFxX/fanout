package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int64 = 42
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        fmt.Println(fi9)
        var v10 int32 = 2147483647
        v11 := GenS1{F0: 1, F1: 42}
        _ = fi9
        _ = v10
        _ = v11
    }
    if true {
        var v12 float64 = 0.0
        _ = v12
    } else {
        if (189 <= 2) {
            var v13 int32 = 988
            var v14 uint64 = 1
            var v15 float32 = 0.0
            _ = v13
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 uint16 = 0
                v18 := GenS1{F0: 18446744073709551615, F1: 255}
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        v19 := []int{127, 0, 127, 32767}
        for fi20 := 0; fi20 < 20; fi20++ {
            v21 := GenS1{F0: 0, F1: 1}
            fmt.Println(fi20)
            var v22 string = ("" + ("hello" + "a"))
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = v19
    }
    var v23 uint = 255
    var v24 uint32 = 147
    _ = v23
    _ = v24
    return 4294967295
}

func main() {
    var v1 string = "hello"
    var v2 string = v1
    var v3 int64 = -93
    v4 := []uint32{2147483647, 42, 880}
    var v5 bool = ((-2.25 / 0.1) >= -2.25)
    var v6 float32 = (float32((0.0 / (1.5 - (1e10 * 0.1)))) * 1e10)
    var v25 int16 = 0
    var v26 int = -653
    fmt.Println(v1)
    fmt.Println(v3, v1, v1)
    v3 = 0
    v27 := GenS1{F0: 527, F1: 2}
    if (!(100 == 5)) {
        var v28 float32 = ((1.5 - (v6 / v6)) - v6)
        _ = v28
    } else {
        v2 = v2
        var v29 uint = 42
        _ = v29
    }
    var v30 int = 42
    v6 = v6
    var v31 bool = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
