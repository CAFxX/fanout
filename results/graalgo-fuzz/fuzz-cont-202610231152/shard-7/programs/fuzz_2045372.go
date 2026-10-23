package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint8 }
func genFn0(p0 int64) int32 {
    p0 = -9223372036854775808
    fmt.Println(p0)
    _ = p0
    return 0
}

func main() {
    var v1 int16 = 672
    v2 := []string{"a", "foo bar", ""}
    var v3 string = ""
    v4 := make(map[uint32]int)
    switch v3 {
    case "hello":
        var v5 int8 = 1
        _ = v5
    case "foo bar":
        v3 = v3
    case "a":
        if (!(-1 == 0)) {
            for fi6 := 0; fi6 < 15; fi6++ {
                v1 -= -1
                v7 := GenS1{F0: (v1 << ((2) & 63))}
                _ = fi6
                _ = v7
            }
            var v8 uint64 = 73
            _ = v8
        } else {
            for fi9 := 0; fi9 < 9; fi9++ {
                fmt.Println(v3)
                v10 := GenS1{F0: v1}
                _ = fi9
                _ = v10
            }
            switch -968 {
            case -9223372036854775808:
                var v11 float32 = -2.25
                _ = v11
            default:
                fmt.Println(v1)
            }
            v12 := GenS2{F0: -250, F1: 315}
            _ = v12
        }
    default:
        v13 := GenS3{F0: 646, F1: 2147483647, F2: 42}
        _ = v13
    }
    var v14 uint16 = 1
    v14 -= (^(v14))
    var v15 int64 = 172
    for fi16 := 0; fi16 < 3; fi16++ {
        v14 = v14
        _ = fi16
    }
    v17 := GenS2{F0: -9223372036854775808, F1: 0}
    _ = len(v4)
    v15 += (v15 | (v15 * (^(v15))))
    fmt.Println(v15, v15, v3)
    var v18 int = -711
    var v19 int64 = v15
    v2 = append(v2, "foo bar")
    var v20 float64 = float64(1.5)
    var v21 uint8 = 128
    switch v18 {
    case 517:
        var v22 uint64 = 1
        _ = v22
    case 127:
        _ = len(v2)
    }
    v18 -= (v18 + -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
