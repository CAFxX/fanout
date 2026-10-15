package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int8 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []int{-2147483648, 4294967295}
    var v2 uint32 = 4294967295
    v3 := make(map[string]int)
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    v6 := make(map[string]int)
    delete(v4, "hello")
    switch (("" + "foo bar") + "a") {
    case "":
        switch ("a" + "foo bar") {
        case "foo bar":
            for fi7 := 0; fi7 < 9; fi7++ {
                var v8 float64 = float64(1.5)
                var v9 uint64 = 1
                var v10 int = (^(fi7))
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
        default:
            _ = len(v1)
        }
    case "a":
        _ = len(v5)
    case "\n\t":
        v2 = 0
    }
    var v11 float64 = -2.25
    v11 = v11
    _ = len(v3)
    var v12 uint16 = 640
    var v13 int8 = -1
    fmt.Println(v13, v13)
    if len(v1) > 0 {
        v1[0] = 127
    }
    switch -1 {
    case 128:
        v12 = v12
    case -477:
        v2 = 4294967295
    }
    fmt.Println(v12, v11, v13)
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 int64 = -936
            _ = fi15
            _ = v16
        }
        v12 -= (v12 ^ (v12 ^ (^(v12))))
        _ = fi14
    }
    var v17 int16 = -1
    var v18 uint = 968
    switch 0 {
    case 128:
        if false {
            v19 := GenS1{F0: 2, F1: 9223372036854775807, F2: 12}
            v13 -= v13
            v2 = (^(v2))
            _ = v19
        }
    case 608:
        v13 -= v13
    case 127:
        v20 := GenS2{F0: 183}
        _ = v20
    }
    switch 1 {
    case 9223372036854775807:
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 int64 = -9223372036854775808
            v12 = 42
            v12 = (v12 | v12)
            _ = fi21
            _ = v22
        }
    case 891:
        v12 = 0
    default:
        var v23 bool = true
        _ = v23
    }
    fmt.Println(v18, v18, v18)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
