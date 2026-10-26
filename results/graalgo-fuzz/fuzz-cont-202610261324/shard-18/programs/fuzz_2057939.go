package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int32; F2 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 428
    v3 := []int{4294967295, -2147483648}
    var v4 string = ("foo bar" + "a")
    var v5 uint = 42
    var v6 uint16 = 1
    fmt.Println(v4, v2)
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(v5, v5)
        _ = fi7
    }
    var v8 float64 = 1e10
    var v9 float64 = 1.5
    for fi10 := 0; fi10 < 3; fi10++ {
        v9 = ((0.0 + float64(0.1)) / 1.5)
        if true {
            v11 := GenS1{F0: 0, F1: 42}
            var v12 uint8 = 42
            var v13 uint8 = v12
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS2{F0: 2147483647}
        _ = fi14
        _ = v15
    }
    switch 4294967295 {
    case 755:
        var v16 int64 = -9223372036854775808
        _ = v16
    default:
        fmt.Println(v5, v5, v5)
    }
    var v17 int8 = 42
    var v18 uint64 = 4294967295
    fmt.Println(v17)
    v19 := GenS3{F0: 0, F1: -967, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
