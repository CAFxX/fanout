package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint8 }

func main() {
    var v1 float64 = 0.1
    var v2 int8 = 0
    v3 := make(map[int]bool)
    v4 := make(map[int]string)
    v5 := make(map[string]int)
    if (42 != 19) {
        var v6 int16 = 42
        if (264 <= 0) {
            var v7 float64 = (1.5 / 0.1)
            var v8 string = "\n\t"
            _ = v7
            _ = v8
        } else {
            fmt.Println(v1, v2, v2)
        }
        v1 -= ((v1 - (3.14159 / (v1 / v1))) + ((3.14159 - float64(1e10)) * (((float64(3.14159) - 3.14159) * 3.14159) * v1)))
        _ = v6
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 int8 = v2
            if (!false) {
                v12 := GenS1{F0: 2147483648, F1: 65535, F2: 0}
                var v13 int8 = 0
                var v14 int8 = (v11 >> ((0) & 63))
                v15 := []int{128, fi10}
                _ = v12
                _ = v13
                _ = v14
                _ = v15
            } else {
                v16 := []int{2147483647}
                _ = v16
            }
            _ = fi10
            _ = v11
        }
        v17 := v4[fi9]
        for fi18 := 0; fi18 < 5; fi18++ {
            v19 := v3[fi18]
            var v20 int16 = 0
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi9
        _ = v17
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        fmt.Println(v1, fi21)
        _ = fi21
    }
    var v22 int8 = v2
    var v23 int64 = 1
    switch "hello" {
    case "foo bar":
        v24 := []int{24, 128, 193, 2147483647}
        _ = v24
    case "hello":
        var v25 string = "a"
        _ = v25
    default:
        v26 := GenS1{F0: 707, F1: 1, F2: 67}
        _ = v26
    }
    var v27 uint32 = 2
    for fi28 := 0; fi28 < 16; fi28++ {
        if (35 >= 2) {
            var v29 float32 = (0.0 - float32(v1))
            v22 = 0
            _ = v29
        }
        _ = fi28
    }
    if (!false) {
        delete(v3, 4294967295)
        v30 := GenS1{F0: (v27 - (v27 * v27)), F1: (^(v27)), F2: 191}
        fmt.Println(v22, v27, v27)
        _ = v30
    }
    switch 0 {
    case -586:
        if false {
            var v31 uint8 = 128
            _ = v31
        }
    case 2147483647:
        v4[-9223372036854775808] = ""
    case 127:
        var v32 float64 = 0.1
        _ = v32
    default:
        v2 = (v2 << ((545) & 63))
    }
    v22 = (v2 ^ 100)
    v2 -= v22
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
