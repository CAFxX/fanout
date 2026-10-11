package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int }
func genFn0(p0 int32, p1 string, p2 uint8) uint32 {
    v7 := []int{}
    p0 -= -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 65535
}
func genFn1(p0 uint16) int64 {
    p0 = (^(p0))
    v8 := GenS1{F0: 127, F1: 1, F2: -9223372036854775808}
    switch 4294967295 {
    case 222:
        var v9 uint32 = 42
        _ = v9
    }
    _ = p0
    _ = v8
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 100
    var v3 uint16 = 457
    v4 := []int{}
    var v5 int16 = 1
    var v6 float32 = ((-2.25 + 0.1) + -2.25)
    for fi10 := 0; fi10 < 19; fi10++ {
        if len(v4) > 0 {
            v4[0] = fi10
        }
        v11 := GenS1{F0: -108, F1: (^(v5)), F2: fi10}
        _ = fi10
        _ = v11
    }
    fmt.Println(v3)
    var v12 uint16 = (v3 + (^(v3)))
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            fmt.Println(v5)
            if false {
                v2 -= v2
                fmt.Println(v6, v3, v12)
                fmt.Println(v12, v12)
            } else {
                v15 := v1[""]
                _ = v15
            }
            var v16 int64 = 9223372036854775807
            _ = fi14
            _ = v16
        }
        switch "" {
        case "foo bar":
            fmt.Println(v12)
        }
        _ = fi13
    }
    fmt.Println(v6, v3, v6)
    var v17 float32 = v6
    v18 := GenS1{F0: -128, F1: (v5 | 1), F2: -2147483648}
    var v19 int = -2147483648
    fmt.Println(v17, v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
