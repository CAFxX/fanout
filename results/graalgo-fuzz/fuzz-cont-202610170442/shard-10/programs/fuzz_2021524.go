package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }
type GenS2 struct { F0 int64; F1 int32 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    var v5 int64 = 9223372036854775807
    v6 := []int{}
    v7 := GenS1{F0: 2, F1: -567, F2: 9223372036854775807}
    v5 += v5
    var v8 float32 = (float32((0.1 - (0.1 + float64(1.5)))) / 1e10)
    _ = len(v4)
    if false {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v9 int64 = (v5 - (v5 | v5))
        _ = v9
    }
    var v10 int64 = v5
    if len(v6) > 0 {
        v6[0] = int(v10)
    }
    if true {
        var v11 uint64 = 1
        _ = v11
    } else {
        var v12 int = 677
        _ = v12
    }
    fmt.Println(v10)
    v8 -= 0.1
    v10 += v10
    if (!(!false)) {
        v5 = (v10 + v10)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v13 int16 = 0
        _ = v13
    } else {
        v14 := GenS2{F0: v5, F1: 322}
        _ = v14
    }
    switch "a" {
    case "\n\t":
        var v15 uint16 = 1
        _ = v15
    case "a":
        delete(v4, 42)
    case "foo bar":
        v5 = (v5 - -9223372036854775808)
    }
    v8 = v8
    if (!(3.14159 >= 1e10)) {
        for fi16 := 0; fi16 < 11; fi16++ {
            v10 += v10
            v5 = v5
            _ = fi16
        }
        v10 = v10
        var v17 int = 4294967295
        _ = v17
    }
    var v18 bool = false
    v1[((("" + "hello") + (("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))) + "foo bar")) + "")] = (1.5 >= ((1e10 + ((float64(float32(0.1)) / float64((((float32(-2.25) + v8) * (v8 - (-2.25 - v8))) + v8))) - ((0.1 - (float64(float32(1e10)) + float64(-2.25))) * 0.0))) - 1.5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
