package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int64 }
type GenS2 struct { F0 int32; F1 uint64 }
func genFn0() uint32 {
    var v7 bool = true
    var v8 float64 = (1e10 * (((((0.1 - (1.5 - 0.1)) + 0.0) / float64(0.1)) / 0.1) + 3.14159))
    var v9 uint64 = 18446744073709551615
    _ = v7
    _ = v8
    _ = v9
    return 2147483648
}

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 1
    v4 := make(map[string]string)
    v5 := make(map[int]int)
    var v6 uint64 = 1
    var v10 uint8 = 0
    v10 += 180
    switch "" {
    case "hello":
        v11 := []int{}
        _ = v11
    }
    var v12 uint64 = v6
    v1 = (v1 | -182)
    var v13 int32 = -328
    switch 143 {
    case -972:
        v14 := GenS1{F0: v10, F1: 301, F2: 0}
        _ = v14
    case -9223372036854775808:
        v15 := GenS1{F0: v10, F1: 2, F2: -9223372036854775808}
        _ = v15
    case -2147483648:
        for fi16 := 0; fi16 < 3; fi16++ {
            v10 -= (^(v10))
            _ = fi16
        }
    }
    v10 = 27
    fmt.Println(v13, v3, v3)
    switch "a" {
    case "foo bar":
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 bool = false
            _ = fi17
            _ = v18
        }
    default:
        var v19 uint = 581
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        v6 += 4294967295
        for fi21 := 0; fi21 < 10; fi21++ {
            fmt.Println(fi20, v6)
            v22 := GenS2{F0: v13, F1: v6}
            v1 += int16(v10)
            _ = fi21
            _ = v22
        }
        if (v1 < v1) {
            var v23 int64 = 42
            var v24 int64 = 0
            fmt.Println(v6, v24, v10)
            _ = v23
            _ = v24
        } else {
            var v25 float32 = -2.25
            _ = len(v4)
            var v26 int = 127
            _ = v25
            _ = v26
        }
        _ = fi20
    }
    var v27 float32 = -2.25
    v10 = (v10 & 23)
    v28 := GenS1{F0: (v10 + v10), F1: 65535, F2: 857}
    v29 := GenS2{F0: v13, F1: v6}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
