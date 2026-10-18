package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint; F2 int8 }
type GenS3 struct { F0 uint64; F1 int8 }
func genFn0(p0 int16, p1 int8, p2 uint8) uint64 {
    var v6 int8 = (p1 << ((65535) & 63))
    fmt.Println(p0, p2)
    v7 := []int{}
    var v8 uint8 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 816
}

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    var v4 int64 = 516
    v5 := make(map[int]int)
    v1 = v1
    var v9 string = "foo bar"
    var v10 bool = (-292 < 0)
    var v11 string = v9
    if (14 <= 100) {
        if true {
            var v12 float32 = 1e10
            _ = v12
        }
        var v13 int = 2147483647
        for fi14 := 0; fi14 < 19; fi14++ {
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 uint = 65535
                v9 = (("" + v11) + "hello")
                var v17 int = fi15
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi14
        }
        _ = v13
    } else {
        var v18 int64 = v4
        fmt.Println(v9, v9)
        _ = v18
    }
    v19 := GenS1{F0: 1, F1: 0}
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 int64 = 9223372036854775807
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := GenS2{F0: genFn0(42, -71, 1), F1: 2, F2: 46}
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    v24 := []int{-665}
    switch 128 {
    case 33:
        v9 = v9
    case 128:
        if len(v24) > 0 {
            v24[0] = -114
        }
    case 320:
        v25 := GenS3{F0: genFn0(42, 100, 175), F1: 105}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
