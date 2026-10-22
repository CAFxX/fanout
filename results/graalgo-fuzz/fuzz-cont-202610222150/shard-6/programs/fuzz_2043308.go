package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 string; F1 bool; F2 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 42
    var v3 int64 = -520
    var v4 uint = 2
    v5 := make(map[int]uint32)
    var v6 float64 = 1.5
    v3 += v3
    if true {
        for fi7 := 0; fi7 < 10; fi7++ {
            v4 = 521
            _ = fi7
        }
        v8 := GenS1{F0: 658, F1: v4}
        var v9 uint64 = 1
        _ = v8
        _ = v9
    } else {
        v10 := v1[626]
        var v11 int64 = v3
        fmt.Println(v11)
        _ = v10
        _ = v11
    }
    v3 = -272
    var v12 uint64 = 2
    v2 = (v2 % (((-(v2))) | 1))
    v3 = v3
    v13 := GenS2{F0: 660, F1: v3}
    v14 := GenS3{F0: "foo bar", F1: ("" != ((("\n\t" + "\n\t") + "") + ("a" + ((("a" + "\n\t") + "hello") + "hello")))), F2: 0}
    v6 -= v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
