package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
func genFn0() bool {
    var v7 int32 = 42
    var v8 int = 42
    fmt.Println(v8, v8)
    var v9 float64 = 0.1
    _ = v7
    _ = v8
    _ = v9
    return (436 < 65535)
}

func main() {
    v1 := []int{0, 128, 127, 32767}
    v2 := []string{"hello", "foo bar"}
    var v3 uint = 4294967295
    var v4 int16 = -454
    v5 := make(map[int]bool)
    var v6 int16 = 42
    if false {
        switch 9223372036854775807 {
        case -92:
            var v10 uint64 = 977
            _ = v10
        case -646:
            v11 := v5[0]
            _ = v11
        default:
            var v12 uint16 = 698
            _ = v12
        }
        var v13 string = "foo bar"
        fmt.Println(v4, v4, v13)
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v3 -= v3
        _ = fi14
    }
    var v15 int16 = (^(v6))
    var v16 float32 = 3.14159
    v17 := GenS1{F0: 917, F1: 18446744073709551615}
    switch 42 {
    case 32768:
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 int16 = -1
            var v20 uint64 = 48
            _ = fi18
            _ = v19
            _ = v20
        }
    case -1:
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 int8 = 12
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 int64 = -1
                v25 := GenS1{F0: 4294967295, F1: 4294967295}
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi21
            _ = v22
        }
    }
    v4 = 1
    delete(v5, -2147483648)
    var v26 int64 = 1
    v26 = v26
    fmt.Println(v3, v4, v26)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
