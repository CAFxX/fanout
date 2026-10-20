package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int64 }

func main() {
    var v1 uint64 = 0
    v2 := make(map[string]uint32)
    var v3 float64 = float64(-2.25)
    var v4 uint32 = 4294967295
    var v5 uint8 = 0
    if (!((v4 % ((120) | 1)) < (v4 * 42))) {
        var v6 uint64 = v1
        v4 -= uint32(v6)
        _ = v6
    } else {
        v7 := GenS1{F0: -50, F1: (("\n\t" + "") + "a"), F2: 729}
        v1 = v1
        v8 := GenS1{F0: 44, F1: "", F2: 2}
        _ = v7
        _ = v8
    }
    var v9 float32 = ((float32(1e10) - -2.25) + (3.14159 + (0.1 - 3.14159)))
    if true {
        v10 := GenS1{F0: -75, F1: "foo bar", F2: 2}
        _ = v10
    }
    if false {
        v11 := GenS1{F0: -1, F1: "foo bar", F2: 457}
        v3 = float64(1e10)
        _ = v11
    } else {
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := GenS1{F0: 100, F1: "foo bar", F2: 99}
            v14 := GenS2{F0: 115}
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 string = "hello"
        _ = v15
    }
    var v16 bool = (false || ((42 < -740) && (531 < 845)))
    v17 := GenS2{F0: -81}
    var v18 int16 = 1
    fmt.Println(v3, v1, v5)
    if v16 {
        fmt.Println(v18)
        if v16 {
            v19 := GenS1{F0: 127, F1: "\n\t", F2: 628}
            _ = v19
        } else {
            v20 := GenS2{F0: -1}
            _ = v20
        }
    }
    v21 := []int{-501, 1}
    var v22 int64 = 1
    v21 = append(v21, 11)
    fmt.Println(v18)
    _ = len(v2)
    var v23 uint64 = v1
    v24 := GenS3{F0: 1, F1: (v22 % ((9223372036854775807) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
