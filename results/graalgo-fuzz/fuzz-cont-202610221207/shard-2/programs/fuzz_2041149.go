package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint }
type GenS2 struct { F0 int8; F1 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    v3 := []int{707, 32768, 127}
    var v4 int32 = 2147483647
    var v5 int16 = -896
    v6 := make(map[int]int)
    v5 = (-(v5))
    if (4294967295 != 2) {
        switch -1 {
        case -1:
            if (1 >= 0) {
                v5 = -32768
            } else {
                var v7 uint32 = 286
                _ = v7
            }
        case 128:
            v8 := GenS1{F0: ((-345 != 127) || (v5 < v5)), F1: 32768, F2: 18446744073709551615}
            _ = v8
        case 42:
            _ = len(v1)
        default:
            v1[4294967295] = -9223372036854775808
        }
        v9 := v6[-506]
        for fi10 := 0; fi10 < 2; fi10++ {
            v11 := GenS1{F0: false, F1: v9, F2: 23}
            v9 = v9
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        switch ("\n\t" + (("" + "hello") + "a")) {
        case "foo bar":
            fmt.Println(v5, v4)
        }
    }
    var v12 uint = 65535
    v12 += v12
    if len(v3) > 0 {
        v3[0] = -167
    }
    var v13 uint32 = 4294967295
    v5 = 32767
    var v14 uint = (v12 | v12)
    fmt.Println(v14, v4)
    if false {
        fmt.Println(v12)
    }
    _ = len(v3)
    v1[(^(v13))] = 4294967295
    v15 := GenS1{F0: true, F1: -2147483648, F2: v12}
    switch "hello" {
    case "a":
        v16 := GenS2{F0: 42, F1: -2147483648}
        _ = v16
    case "\n\t":
        var v17 float32 = -2.25
        _ = v17
    }
    v14 = v12
    var v18 int64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
