package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
func genFn0() uint {
    var v4 int = 2147483647
    var v5 float64 = 3.14159
    v6 := []int{}
    _ = v4
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 int16 = 42
    v2 := []int{128, 32768, 1, 32767}
    var v3 uint32 = 2147483647
    v7 := GenS1{F0: 42, F1: -9223372036854775808}
    v3 += (v3 & (v3 - (^(v3))))
    if (-1 > v1) {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int16 = v1
            v9 = 1000
            _ = fi8
            _ = v9
        }
        v1 -= (v1 % (((v1 << ((255) & 63))) | 1))
        fmt.Println(v1, v1)
    } else {
        var v10 bool = true
        _ = v10
    }
    switch ("foo bar" + ("" + "")) {
    case "foo bar":
        fmt.Println(v1)
    case "hello":
        v3 = 135
    case "a":
        v2 = append(v2, 0)
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(fi11, v3, fi11)
        var v12 uint32 = (v3 << ((65535) & 63))
        _ = fi11
        _ = v12
    }
    var v13 float32 = (-2.25 / float32(float64(0.1)))
    v1 = (v1 << ((2) & 63))
    var v14 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
}
