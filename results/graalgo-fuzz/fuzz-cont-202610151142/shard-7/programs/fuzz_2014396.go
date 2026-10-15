package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int32 = -2147483648
    var v3 int = 32768
    v3 -= v3
    switch v3 {
    case -476:
        var v4 float32 = 0.0
        _ = v4
    case -699:
        fmt.Println(v3, v1, v2)
    default:
        v2 = v2
    }
    v5 := []int{(v3 << ((42) & 63)), (^(v3)), v3, (v3 & v3)}
    var v6 uint8 = 22
    v7 := GenS1{F0: (!true), F1: v1}
    var v8 uint8 = (v6 % (((v6 % ((160) | 1))) | 1))
    if len(v5) > 0 {
        v5[0] = v3
    }
    v9 := GenS2{F0: ""}
    v10 := GenS3{F0: -32768}
    v11 := GenS4{F0: (true && ((("foo bar" + "") + "") < "hello"))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
