package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 bool; F2 uint }

func main() {
    var v1 uint = 2
    var v2 bool = false
    var v3 int64 = -9223372036854775808
    var v4 float64 = float64(1e10)
    v1 = v1
    if v2 {
        for fi5 := 0; fi5 < 19; fi5++ {
            v4 = 3.14159
            for fi6 := 0; fi6 < 3; fi6++ {
                var v7 float32 = 0.1
                var v8 string = "\n\t"
                var v9 int16 = 42
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            _ = fi5
        }
        v10 := GenS1{F0: -506}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            v2 = (0 > 105)
            _ = fi11
        }
        var v12 uint = (v1 >> ((v1) & 63))
        _ = v12
    }
    v3 += v3
    if false {
        v13 := []int{790, -109, 9223372036854775807}
        _ = v13
    }
    v14 := []int{144, 701, 9223372036854775807}
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint64 = 27
        v14 = append(v14, (fi15 ^ 4294967295))
        _ = fi15
        _ = v16
    }
    var v17 int = -709
    v18 := GenS2{F0: 0, F1: (((((("a" + "") + "a") + "foo bar") + "foo bar") != "hello") || true), F2: v1}
    switch v17 {
    case 128:
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    case -192:
        _ = len(v14)
    case -2147483648:
        v4 = ((1.5 + float64(((float32(((v4 - 1e10) - (3.14159 - float64(((float32(float64(0.1)) / 1e10) - float32((1e10 + v4))))))) - 0.0) * 0.1))) * v4)
    }
    fmt.Println(v3, v2, v3)
    v19 := GenS2{F0: 65535, F1: true, F2: 4294967295}
    _ = len(v14)
    v4 += float64(1.5)
    var v20 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
