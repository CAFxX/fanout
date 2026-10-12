package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 0.1
    var v3 int32 = -2147483648
    var v4 float64 = 0.1
    var v5 int = 42
    switch "hello" {
    case "foo bar":
        v3 -= (v3 & v3)
    case "hello":
        fmt.Println(v2, v3)
    default:
        v2 = v2
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: 898, F1: (104 > 0), F2: 68}
        var v8 uint16 = 32767
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 int32 = 821
        v11 := GenS1{F0: 1, F1: (v3 >= (v3 / ((1) | 1))), F2: 609}
        v2 = (1e10 - 1e10)
        _ = fi9
        _ = v10
        _ = v11
    }
    switch "a" {
    case "foo bar":
        var v12 int16 = 42
        _ = v12
    case "a":
        var v13 uint = 0
        _ = v13
    case "hello":
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := GenS1{F0: 1, F1: (1 <= 1), F2: 242}
            var v16 float32 = v2
            fmt.Println(v3, v2, v2)
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    var v17 float64 = v4
    v18 := GenS1{F0: -9223372036854775808, F1: (false || (465 == 949)), F2: 65535}
    v5 -= v5
    var v19 int16 = -32768
    v20 := GenS1{F0: -1, F1: false, F2: 42}
    switch "a" {
    case "a":
        fmt.Println(v17, v17)
    case "foo bar":
        var v21 int = (-(v5))
        _ = v21
    case "\n\t":
        if ((v17 - 1e10) >= (v4 * v17)) {
            v1[v5] = ""
            var v22 string = "hello"
            _ = v22
        } else {
            switch "" {
            case "a":
                var v23 bool = false
                _ = v23
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
