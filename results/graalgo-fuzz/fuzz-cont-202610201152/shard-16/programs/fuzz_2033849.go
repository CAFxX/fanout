package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int64 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int64; F1 int; F2 int8 }

func main() {
    var v1 float64 = -2.25
    v2 := []uint32{2, 2, 2, 2147483647}
    v3 := make(map[int]uint32)
    var v4 float64 = (v1 + 0.0)
    v5 := []int{-9223372036854775808}
    var v6 uint16 = 255
    v7 := GenS1{F0: "", F1: 1, F2: 9223372036854775807}
    var v8 string = "foo bar"
    v9 := GenS1{F0: ((((("foo bar" + v8) + (v8 + "foo bar")) + "\n\t") + ("a" + v8)) + ((v8 + ("hello" + v8)) + ((v8 + v8) + v8))), F1: 127, F2: -9223372036854775808}
    var v10 uint8 = 2
    var v11 string = v8
    var v12 bool = ((127 < v10) && (-1 == -1))
    for fi13 := 0; fi13 < 10; fi13++ {
        switch 2147483647 {
        case 32767:
            v14 := GenS2{F0: 1, F1: ((^(v10)) == (v10 | (^(v10))))}
            _ = v14
        }
        if (v10 == v10) {
            var v15 int64 = 1
            _ = v15
        }
        _ = fi13
    }
    var v16 uint64 = 653
    if len(v5) > 0 {
        v5[0] = 32767
    }
    switch ((v11 + (((v11 + v11) + "foo bar") + v8)) + v8) {
    case "foo bar":
        v17 := GenS3{F0: 782, F1: -1, F2: -13}
        _ = v17
    }
    var v18 uint32 = 660
    v19 := GenS2{F0: (v18 | (v18 >> ((0) & 63))), F1: true}
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := GenS1{F0: v8, F1: 42, F2: 1}
        switch fi20 {
        case 0:
            v5 = append(v5, 0)
        case 4294967295:
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 int16 = 1000
                _ = fi22
                _ = v23
            }
        case -895:
            fmt.Println(v11, v12, v4)
        default:
            var v24 string = (v11 + v8)
            _ = v24
        }
        if (!false) {
            v10 -= (v10 >> ((uint(v18)) & 63))
        } else {
            var v25 float32 = ((float32(v1) * 1.5) / (1.5 * 1e10))
            _ = v25
        }
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
