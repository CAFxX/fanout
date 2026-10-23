package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
func genFn0(p0 int64) int8 {
    var v5 float32 = (0.1 / 1e10)
    switch "" {
    case "":
        fmt.Println(p0, p0)
    case "\n\t":
        var v6 uint16 = 42
        _ = v6
    case "a":
        if (p0 < p0) {
            var v7 uint64 = 195
            v7 -= (v7 + 42)
            v8 := GenS1{F0: (p0 ^ -879), F1: 113}
            _ = v7
            _ = v8
        }
    }
    _ = p0
    _ = v5
    return 1
}

func main() {
    var v1 int32 = 846
    var v2 float32 = 1.5
    v3 := make(map[uint32]string)
    var v4 uint16 = 42
    var v9 uint32 = 4294967295
    var v10 uint64 = 4294967295
    var v11 float32 = v2
    if (128 != 42) {
        var v12 int32 = (^(v1))
        var v13 uint16 = (^(v4))
        v12 = (v12 / (((v12 >> ((42) & 63))) | 1))
        _ = v12
        _ = v13
    } else {
        fmt.Println(v11, v10)
        v10 -= (^(v10))
        v14 := GenS1{F0: -54, F1: genFn0(191)}
        _ = v14
    }
    v15 := GenS1{F0: 42, F1: genFn0(42)}
    fmt.Println(v1, v1, v9)
    fmt.Println(v2, v9, v2)
    v1 = -2147483648
    for fi16 := 0; fi16 < 20; fi16++ {
        v17 := GenS1{F0: 42, F1: 42}
        var v18 uint32 = v9
        v19 := v3[(v9 ^ 814)]
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    switch -2147483648 {
    case 405:
        switch ("\n\t" + "\n\t") {
        case "hello":
            var v20 uint64 = 201
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        delete(v3, v9)
        _ = fi21
    }
    v22 := GenS1{F0: 949, F1: 38}
    var v23 int64 = 930
    var v24 int32 = v1
    for fi25 := 0; fi25 < 17; fi25++ {
        if false {
            v26 := GenS1{F0: (v23 * v23), F1: 42}
            v2 += ((((((float32(float64(v2)) + float32(-2.25)) * 0.0) + (0.1 - 0.0)) - -2.25) / (3.14159 * v11)) / 1.0)
            _ = v26
        }
        var v27 int64 = (v23 & (v23 + v23))
        _ = fi25
        _ = v27
    }
    _ = len(v3)
    var v28 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v9
}
