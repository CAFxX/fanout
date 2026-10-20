package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 float32 = ((1e10 + 1.5) * 3.14159)
    var v4 uint64 = 227
    var v5 float64 = (0.1 + (0.0 - 1e10))
    var v6 int16 = -71
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: -1, F1: 0}
        v9 := GenS2{F0: 647, F1: 2147483647}
        _ = fi7
        _ = v8
        _ = v9
    }
    v2[4294967295] = 1
    v4 = 1
    for fi10 := 0; fi10 < 11; fi10++ {
        if false {
            v11 := GenS3{F0: -9223372036854775808}
            var v12 int8 = 1
            switch ("foo bar" + ("hello" + "a")) {
            case "a":
                fmt.Println(v12, v12)
            default:
                v13 := GenS4{F0: (fi10 * 313), F1: -9223372036854775808}
                _ = v13
            }
            _ = v11
            _ = v12
        } else {
            for fi14 := 0; fi14 < 9; fi14++ {
                var v15 uint = 8
                _ = fi14
                _ = v15
            }
            var v16 int16 = (v6 ^ (v6 << ((1) & 63)))
            fmt.Println(v3, v3)
            _ = v16
        }
        v17 := GenS5{F0: fi10}
        v5 = float64(0.0)
        _ = fi10
        _ = v17
    }
    v6 = 1000
    var v18 string = "a"
    fmt.Println(v3, v3)
    v19 := []int{147, 2147483647}
    v20 := GenS6{F0: 51}
    for fi21 := 0; fi21 < 18; fi21++ {
        v22 := v2[2147483648]
        v23 := GenS1{F0: -954, F1: (v4 - (v4 | v4))}
        _ = fi21
        _ = v22
        _ = v23
    }
    if len(v19) > 0 {
        v19[0] = 0
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
