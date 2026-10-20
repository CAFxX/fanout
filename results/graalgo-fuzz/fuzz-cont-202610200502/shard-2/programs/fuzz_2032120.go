package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint64; F1 uint8; F2 int16 }

func main() {
    var v1 int64 = 42
    v2 := []int{}
    var v3 uint8 = 2
    var v4 uint8 = v3
    v5 := GenS1{F0: (42 == 0), F1: 42}
    var v6 bool = (!(true || true))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v2 = append(v2, 1)
    if v6 {
        if true {
            v1 += (^(v1))
        } else {
            fmt.Println(v4, v1, v3)
            switch -9223372036854775808 {
            case 128:
                v1 -= 1
            }
        }
        if v6 {
            var v7 int32 = -2147483648
            _ = v7
        }
        v1 -= (-(v1))
    } else {
        var v8 int32 = 1
        _ = v8
    }
    if true {
        var v9 float64 = (-2.25 + 3.14159)
        v2 = append(v2, -2147483648)
        _ = v9
    }
    var v10 int32 = 0
    fmt.Println(v1, v10, v3)
    var v11 int = 0
    var v12 int16 = 1000
    v13 := GenS2{F0: 77, F1: v3, F2: v12}
    var v14 string = (("hello" + ("\n\t" + "a")) + "hello")
    var v15 float64 = 0.0
    v15 -= (3.14159 - v15)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
