package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int64; F1 uint32 }

func main() {
    var v1 uint16 = 1
    v2 := []string{"a"}
    v3 := []int{}
    var v4 int16 = 1
    var v5 int32 = 0
    if ((!((float64((((-2.25 * 1.5) + 1e10) + 1.5)) <= (1.5 / (-2.25 + 3.14159))) && true)) && true) {
        var v6 float32 = 0.1
        _ = v6
    } else {
        v7 := GenS1{F0: 32768}
        v3 = append(v3, 9223372036854775807)
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS2{F0: -9223372036854775808, F1: 9223372036854775807}
        v3 = append(v3, (fi8 % ((fi8) | 1)))
        _ = fi8
        _ = v9
    }
    var v10 int16 = -32768
    var v11 string = "hello"
    switch 9223372036854775807 {
    case -870:
        fmt.Println(v11, v10, v11)
    }
    var v12 float32 = 0.1
    v13 := GenS3{F0: 978, F1: 0}
    v14 := GenS2{F0: 9223372036854775807, F1: -794}
    v15 := GenS4{F0: 3, F1: 1}
    var v16 uint8 = 42
    var v17 uint32 = 0
    var v18 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
