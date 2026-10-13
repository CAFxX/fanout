package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 bool }
func genFn0(p0 int32, p1 uint8, p2 int32) int {
    if (79 != 1) {
        for fi6 := 0; fi6 < 15; fi6++ {
            if false {
                var v7 int64 = -1
                var v8 uint32 = 541
                p2 = p2
                _ = v7
                _ = v8
            }
            v9 := GenS1{F0: (!(!(!(368 >= 42)))), F1: 283}
            _ = fi6
            _ = v9
        }
        v10 := []int{-924, 89, 794}
        _ = v10
    }
    fmt.Println(p0, p0, p2)
    if true {
        v11 := []int{857, -1, 9223372036854775807}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint = 0
    var v2 float32 = (3.14159 / (3.14159 + 1e10))
    var v3 uint = (v1 / (((v1 << ((18446744073709551615) & 63))) | 1))
    v4 := make(map[int]int)
    var v5 int16 = 227
    var v12 uint32 = uint32(v3)
    v13 := GenS2{F0: (42 == 128)}
    v1 = v1
    fmt.Println(v3, v12, v1)
    v14 := GenS1{F0: (!true), F1: 200}
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 uint64 = 1
        for fi17 := 0; fi17 < 2; fi17++ {
            _ = len(v4)
            if (!false) {
                v18 := v4[(fi17 * (fi15 + (fi15 % ((genFn0(42, 255, -735)) | 1))))]
                var v19 int32 = 1
                fmt.Println(v5, v2, v3)
                _ = v18
                _ = v19
            }
            switch "\n\t" {
            case "a":
                v20 := []int{fi15, (fi17 & genFn0(809, 0, -2147483648)), (fi17 * genFn0(-73, 128, -2147483648))}
                _ = v20
            }
            var v21 uint16 = 2
            var v22 int = fi17
            v1 = (v1 | (v1 ^ (v3 - uint(v16))))
            _ = fi17
            _ = v21
            _ = v22
        }
        _ = fi15
        _ = v16
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v3)
    case "":
        v4[genFn0(1, 0, 1)] = -9223372036854775808
    case "\n\t":
        var v23 int8 = -128
        _ = v23
    }
    if true {
        var v24 string = "foo bar"
        fmt.Println(v2, v1)
        _ = v24
    }
    v5 = int16(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
