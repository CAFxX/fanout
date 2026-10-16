package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint64 }
func genFn0() int64 {
    var v5 int16 = 106
    switch "" {
    case "":
        switch "\n\t" {
        case "foo bar":
            for fi6 := 0; fi6 < 9; fi6++ {
                v7 := GenS1{F0: (2 == 1)}
                v8 := []int{(^(fi6)), fi6, 1}
                var v9 uint32 = 42
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
        default:
            v10 := GenS2{F0: 1}
            _ = v10
        }
    default:
        v11 := []int{384}
        _ = v11
    }
    v12 := GenS1{F0: (-284 > 42)}
    v13 := GenS2{F0: 1}
    var v14 string = "hello"
    _ = v12
    _ = v13
    _ = v14
    _ = v5
    return 1
}

func main() {
    var v1 uint16 = 233
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS2{F0: 662}
        v1 += v1
        v17 := []int{}
        _ = fi15
        _ = v16
        _ = v17
    }
    v1 = (^(v1))
    fmt.Println(v1)
    v1 += 608
    v1 += v1
    delete(v2, 2147483647)
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 int16 = 1
        if (!(v19 >= 32767)) {
            v20 := []int{(^(fi18)), fi18}
            v1 = v1
            _ = v20
        } else {
            v19 = -666
        }
        _ = len(v4)
        _ = fi18
        _ = v19
    }
    v21 := []int{-9223372036854775808, 913, 128, -249}
    v22 := GenS3{F0: 86, F1: 1, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
