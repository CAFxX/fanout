package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 string; F1 uint32 }

func main() {
    var v1 int32 = 663
    var v2 uint8 = 145
    var v3 uint16 = 54
    v4 := GenS1{F0: 0, F1: 9223372036854775807}
    var v5 int64 = -141
    v6 := GenS1{F0: -1, F1: 128}
    v1 += -1
    fmt.Println(v5, v1)
    fmt.Println(v3, v5, v5)
    v7 := []int{501, -2147483648, 0, -404}
    v5 -= (v5 * v5)
    var v8 int16 = 1
    var v9 string = "foo bar"
    v10 := GenS2{F0: (v9 + v9), F1: 19}
    fmt.Println(v3, v1)
    var v11 int = 42
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(v11, v3)
        _ = fi12
    }
    if true {
        var v13 bool = (!((v11 & 42) != v11))
        if false {
            if (v9 < (("a" + v9) + v9)) {
                v13 = v13
                var v14 uint32 = 257
                fmt.Println(v5)
                _ = v14
            } else {
                if len(v7) > 0 {
                    v7[0] = v11
                }
            }
            if true {
                v7 = append(v7, (v11 * v11))
                v15 := GenS2{F0: (v9 + v9), F1: 402}
                _ = v15
            } else {
                fmt.Println(v13, v11)
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
        var v16 int64 = -1
        _ = v13
        _ = v16
    } else {
        var v17 uint64 = 4294967295
        if len(v7) > 0 {
            v7[0] = v11
        }
        v9 = v9
        _ = v17
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
