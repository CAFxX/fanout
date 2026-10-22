package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }

func main() {
    var v1 float64 = 0.1
    var v2 uint64 = 2
    v3 := make(map[uint32]bool)
    var v4 string = "foo bar"
    var v5 bool = (!(170 == 42))
    var v6 int32 = -1
    v7 := GenS1{F0: 65535, F1: -72}
    v8 := GenS1{F0: 32767, F1: -50}
    v9 := []int{-9223372036854775808, -259, 42, 206}
    _ = len(v9)
    v4 = v4
    var v10 uint32 = 42
    _ = len(v3)
    for fi11 := 0; fi11 < 8; fi11++ {
        if len(v9) > 0 {
            v9[0] = (fi11 / ((fi11) | 1))
        }
        _ = fi11
    }
    delete(v3, (v10 % (((v10 ^ v10)) | 1)))
    var v12 int64 = -1
    if v5 {
        switch 0 {
        case -2147483648:
            v2 = v2
        }
        if (v5 || (((v10 * (v10 * 1)) >= v10) || ((v10 >= v10) || (42 != 0)))) {
            if v5 {
                var v13 int16 = 1
                _ = v13
            }
        }
        v12 = -405
    } else {
        v9 = append(v9, -2147483648)
    }
    var v14 string = (((("hello" + "") + v4) + v4) + v4)
    var v15 int32 = (^(v6))
    fmt.Println(v6, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
