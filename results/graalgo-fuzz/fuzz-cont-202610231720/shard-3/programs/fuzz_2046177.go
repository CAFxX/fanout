package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int8; F2 string }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 int64 }

func main() {
    v1 := []int{-923}
    v2 := make(map[uint32]bool)
    var v3 int64 = 0
    v4 := []int{0, 127, 128}
    var v5 bool = false
    v6 := make(map[string]int)
    v7 := GenS1{F0: 1000}
    for fi8 := 0; fi8 < 1; fi8++ {
        v5 = v5
        v4 = append(v4, (fi8 * -554))
        _ = fi8
    }
    v9 := GenS2{F0: 42, F1: -36, F2: ""}
    if false {
        v10 := GenS3{F0: (v3 % ((v3) | 1)), F1: 284}
        v4 = append(v4, 1)
        v11 := GenS3{F0: v3, F1: 110}
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        v3 -= v3
        _ = fi12
    }
    var v13 string = ("\n\t" + ((("foo bar" + "foo bar") + "a") + ("hello" + ("\n\t" + "a"))))
    v14 := GenS4{F0: (-645 == 906)}
    for fi15 := 0; fi15 < 9; fi15++ {
        _ = len(v4)
        v16 := GenS5{F0: 1, F1: (^(v3))}
        v3 = v3
        _ = fi15
        _ = v16
    }
    var v17 float32 = float32(-2.25)
    switch ("a" + ((v13 + "") + "")) {
    case "a":
        fmt.Println(v13)
    default:
        v18 := GenS1{F0: 1}
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
