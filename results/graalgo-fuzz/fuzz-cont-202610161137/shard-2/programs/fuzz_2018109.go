package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 0
    v3 := []string{"", ("hello" + "\n\t"), "hello", "\n\t"}
    var v4 int16 = 857
    var v5 int = 32768
    var v6 string = "\n\t"
    var v7 float64 = 0.1
    v8 := GenS1{F0: (v5 | (v5 << ((1) & 63))), F1: 9223372036854775807, F2: 42}
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(v5, v7, v4)
        if false {
            v4 += v4
            v10 := GenS1{F0: (^(v5)), F1: -350, F2: 255}
            _ = v10
        } else {
            fmt.Println(fi9)
        }
        _ = fi9
    }
    fmt.Println(v7, v6, v7)
    v11 := GenS2{F0: v6}
    v12 := GenS2{F0: v6}
    var v13 int = v5
    v14 := GenS3{F0: (v6 + "foo bar")}
    if true {
        v15 := GenS1{F0: (-(v5)), F1: 615, F2: 218}
        var v16 int32 = 561
        _ = v15
        _ = v16
    }
    var v17 int16 = v4
    fmt.Println(v7, v17)
    v7 += ((0.1 - 1e10) / 0.1)
    var v18 bool = ((false && (9223372036854775807 > (^(v5)))) && false)
    var v19 uint64 = v2
    var v20 int32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
