package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }

func main() {
    var v1 float32 = ((0.0 / 0.1) - ((3.14159 / -2.25) * 0.0))
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    var v4 float32 = ((0.0 + (v1 - (3.14159 - v1))) * 3.14159)
    var v5 uint64 = 0
    fmt.Println(v4, v4)
    _ = len(v3)
    if false {
        var v6 uint32 = 42
        _ = v6
    } else {
        var v7 uint16 = 101
        _ = v7
    }
    var v8 uint = 2
    var v9 uint64 = (v5 / (((v5 % ((v5) | 1))) | 1))
    var v10 string = ""
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 int32 = -2147483648
            v13 -= v13
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v5 = v9
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := GenS1{F0: -674, F1: -24}
        v16 := []int{1}
        switch "hello" {
        case "hello":
            switch (fi14 % ((fi14) | 1)) {
            case -753:
                v10 = "a"
            case 128:
                v17 := v2[0]
                _ = v17
            case -12:
                v18 := GenS1{F0: fi14, F1: 0}
                _ = v18
            }
        case "\n\t":
            var v19 uint32 = 695
            _ = v19
        default:
            var v20 int64 = 186
            _ = v20
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        if (!(65535 <= 296)) {
            v22 := GenS1{F0: 127, F1: -1}
            var v23 uint = 42
            switch "a" {
            case "foo bar":
                v1 = 1e10
            case "\n\t":
                var v24 int = (fi21 / ((1) | 1))
                _ = v24
            case "":
                var v25 string = v10
                _ = v25
            default:
                var v26 int32 = 0
                _ = v26
            }
            _ = v22
            _ = v23
        }
        _ = fi21
    }
    v1 -= (float32((0.1 / 1e10)) - float32(0.1))
    v9 += v9
    for fi27 := 0; fi27 < 17; fi27++ {
        v5 = v9
        _ = fi27
    }
    if true {
        v4 = v1
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        fmt.Println(v10, v4)
        _ = fi28
    }
    v29 := GenS1{F0: 127, F1: 42}
    fmt.Println(v9, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
