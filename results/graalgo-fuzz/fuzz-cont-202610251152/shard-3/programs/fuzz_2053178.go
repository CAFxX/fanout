package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int8 }
func genFn0() int {
    if (((float64(0.0) + (0.0 - float64((0.0 * 0.0)))) - (3.14159 * 0.1)) >= (((1e10 - 1.5) + 0.1) - (1.5 + -2.25))) {
        var v6 int8 = 42
        _ = v6
    } else {
        v7 := GenS1{F0: 2}
        var v8 string = "foo bar"
        fmt.Println(v8)
        _ = v7
        _ = v8
    }
    switch ("a" + "hello") {
    case "\n\t":
        switch -2147483648 {
        case 2147483647:
            var v9 int = 32767
            _ = v9
        }
    }
    var v10 float32 = 1.5
    _ = v10
    return 128
}

func main() {
    v1 := []int{9223372036854775807}
    var v2 uint8 = 2
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    var v5 bool = (65535 >= 18446744073709551615)
    _ = len(v3)
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 int8 = 54
        v2 -= v2
        _ = fi11
        _ = v12
    }
    v2 = (^(v2))
    var v13 uint8 = v2
    var v14 string = "\n\t"
    v14 = v14
    switch "a" {
    case "foo bar":
        v5 = (65535 <= 2)
    case "":
        var v15 float32 = 0.0
        _ = v15
    default:
        for fi16 := 0; fi16 < 10; fi16++ {
            v14 = v14
            var v17 bool = true
            var v18 uint8 = 127
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    v19 := GenS1{F0: 4294967295}
    if (((v13 < 212) || v5) && (!(!(((genFn0() <= 32768) || (255 > 35)) || v5)))) {
        _ = len(v4)
        if true {
            if (-2.25 != (1.5 / (0.1 / (1e10 / 3.14159)))) {
                for fi20 := 0; fi20 < 12; fi20++ {
                    v2 = v13
                    var v21 uint8 = (v2 % ((v2) | 1))
                    _ = fi20
                    _ = v21
                }
                var v22 int16 = 1
                v23 := GenS2{F0: -576, F1: -75}
                _ = v22
                _ = v23
            } else {
                var v24 int64 = -1
                _ = v24
            }
            fmt.Println(v13)
            v5 = (!(221 == 895))
        }
    } else {
        var v25 float32 = float32((0.1 / 3.14159))
        v26 := GenS2{F0: 0, F1: 42}
        fmt.Println(v13)
        _ = v25
        _ = v26
    }
    _ = len(v4)
    var v27 string = v14
    for fi28 := 0; fi28 < 3; fi28++ {
        _ = len(v3)
        _ = fi28
    }
    v13 -= v2
    for fi29 := 0; fi29 < 12; fi29++ {
        v2 -= v13
        v5 = (v5 || v5)
        _ = fi29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
