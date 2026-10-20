package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int8 }

func main() {
    var v1 uint16 = 255
    var v2 uint16 = (v1 + (^(v1)))
    v3 := make(map[uint32]int)
    var v4 int8 = -1
    var v5 int = 9223372036854775807
    var v6 uint64 = 4294967295
    fmt.Println(v1)
    var v7 int16 = 638
    v6 = v6
    delete(v3, 42)
    if (!((v6 & 0) != 989)) {
        var v8 uint64 = 0
        _ = v8
    }
    v9 := GenS1{F0: 2}
    v10 := []int{v5, -416, (^(v5)), v5}
    v11 := GenS1{F0: 2}
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 float64 = -2.25
        v14 := GenS2{F0: 52}
        var v15 int64 = 996
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v6 = 0
    v4 = (v4 / ((v4) | 1))
    v6 = v6
    if false {
        var v16 bool = true
        var v17 float32 = (((-2.25 / -2.25) / 3.14159) + (0.0 - 0.1))
        _ = v16
        _ = v17
    }
    if true {
        if ((("a" + "foo bar") + "foo bar") != "\n\t") {
            var v18 float64 = -2.25
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            var v19 float32 = ((float32(v18) + float32(v18)) / 1e10)
            _ = v18
            _ = v19
        } else {
            v20 := GenS3{F0: (v4 + v4), F1: (v4 + int8(v1))}
            for fi21 := 0; fi21 < 10; fi21++ {
                fmt.Println(v2)
                _ = fi21
            }
            v3[0] = v5
            _ = v20
        }
    }
    if len(v10) > 0 {
        v10[0] = -9223372036854775808
    }
    var v22 uint64 = v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
