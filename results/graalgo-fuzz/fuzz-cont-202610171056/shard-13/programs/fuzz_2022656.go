package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int64; F2 int8 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 uint = 740
    var v2 int64 = -129
    var v3 int8 = 54
    var v4 int16 = 42
    v5 := GenS1{F0: (v2 * (v2 & (v2 & (^(v2))))), F1: (v4 & v4)}
    v6 := GenS1{F0: -668, F1: 32767}
    fmt.Println(v3, v2)
    v7 := []int{}
    if (668 != 42) {
        var v8 int = -2147483648
        var v9 int8 = 57
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 0, F1: 1}
    if true {
        var v11 uint64 = 0
        fmt.Println(v3)
        v7 = append(v7, 42)
        _ = v11
    }
    if len(v7) > 0 {
        v7[0] = 816
    }
    fmt.Println(v1, v3, v2)
    for fi12 := 0; fi12 < 9; fi12++ {
        fmt.Println(v1, v3)
        v13 := GenS2{F0: 0, F1: 4294967295}
        _ = fi12
        _ = v13
    }
    v14 := GenS3{F0: "foo bar"}
    fmt.Println(v1, v3)
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 int16 = (v4 - (^(v4)))
        fmt.Println(v4, fi15)
        switch "hello" {
        case "a":
            for fi17 := 0; fi17 < 10; fi17++ {
                v18 := GenS4{F0: true, F1: v2, F2: (v3 * 1)}
                var v19 string = "hello"
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        _ = fi15
        _ = v16
    }
    v20 := GenS5{F0: 442}
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 int8 = v3
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
